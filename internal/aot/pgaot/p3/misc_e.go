package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_EA_get_flat_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L59
	} else {
		goto L60
	}
L2:
	;
	m.G0 = v13 + int32(16)
	return v187
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v187 = int32(base.Ui32(v16) >> (uint(int32(2)) % 32))
	goto L2
L4:
	;
	goto L5
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v19 != 0 {
		v187 = v19
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v22 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v171 = v21 << (uint(int32(3)) % 32)
	if v20 != 0 {
		goto L56
	} else {
		goto L57
	}
L8:
	;
	v161 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v29 = int32(0)
	v31 = int32(0)
	goto L11
L11:
	;
	if v20 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v161 = v154
	goto L7
L13:
	;
	v158 = v31 + int32(1)
	if v158 != v22 {
		v29 = v154
		v31 = v158
		goto L11
	} else {
		goto L55
	}
L14:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v20))))
	if v39 != 0 {
		v154 = v29
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+44)))
	if int32(0) < v40 {
		v133 = v40
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v135 = v29 + v133
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+47)))
	switch v136 - int32(99) {
	case 0:
		v151 = v135
		goto L50
	case 1:
		goto L52
	default:
		goto L51
	case 6:
		goto L53
	}
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v26+v31<<(uint(int32(2))%32))))
	if v40 == int32(-1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v49 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	if v46&int32(3) == int32(0) {
		v97 = v46
		goto L35
	} else {
		goto L36
	}
L23:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+1)))
	if base.Ui32((v53-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v133 = int32(6)
		goto L18
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v49&int32(1) != 0 {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v60 = int32(18)
	if v53&int32(255) == v60 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v66 = v60
	goto L29
L28:
	;
	v66 = int32(2)
	goto L29
L29:
	;
	v133 = v66
	goto L18
L30:
	;
	v133 = int32(base.Ui32(v49) >> (uint(int32(1)) % 32))
	goto L18
L31:
	;
	goto L32
L32:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v133 = int32(base.Ui32(v71) >> (uint(int32(2)) % 32))
	goto L18
L33:
	;
	v133 = v130 + int32(1)
	goto L18
L34:
	;
	v130 = v122 - v46
	goto L33
L35:
	;
	v101 = v97
	goto L44
L36:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v81 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v130 = int32(0)
	goto L33
L38:
	;
	goto L39
L39:
	;
	v86 = v46
	goto L40
L40:
	;
	v90 = v86 + int32(1)
	if v90&int32(3) == int32(0) {
		v97 = v90
		goto L35
	} else {
		goto L42
	}
L41:
	;
	v122 = v90
	goto L34
L42:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v95 != 0 {
		v86 = v90
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v110 = int32(-2139062144)
	if (int32(16843008)-v107|v107)&v110 == v110 {
		v101 = v101 + int32(4)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v116 = v101
	goto L47
L46:
	;
	goto L45
L47:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v120 != 0 {
		v116 = v116 + int32(1)
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v122 = v116
	goto L34
L49:
	;
	goto L48
L50:
	;
	if base.Ui32(int32(1073741824)) <= base.Ui32(v151) {
		goto L1
	} else {
		goto L54
	}
L51:
	;
	v151 = (v135 + int32(1)) & int32(-2)
	goto L50
L52:
	;
	v151 = (v135 + int32(7)) & int32(-8)
	goto L50
L53:
	;
	v151 = (v135 + int32(3)) & int32(-4)
	goto L50
L54:
	;
	v154 = v151
	goto L13
L55:
	;
	goto L12
L56:
	;
	v175 = base.I32_div_s(v22+int32(7), int32(8))
	v181 = v171 + v175 + int32(23)
	goto L58
L57:
	;
	v181 = v171 + int32(23)
	goto L58
L58:
	;
	v184 = v181&int32(-8) + v161
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v184
	v187 = v184
	goto L2
L59:
	;
	return int32(0)
L60:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(1073741823)
	F_errmsg(m, int32(658024), v13)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(494245), int32(275), int32(336871))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L59
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_EOH_get_flat_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	v4 = m.T0[v3].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_EndImplicitTransactionBlock(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
	if v4 == int32(4) {
		*(*int32)(unsafe.Add(mBase, uint32(v3)+24)) = int32(1)
	} else {
	}
	return
}
func F_EvaluateParams(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if l2 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(48)
	return v208
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L13
	} else {
		goto L42
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v20 = v18
	goto L5
L4:
	;
	v20 = int32(0)
	goto L5
L5:
	;
	if v17 == v20 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if v17 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L13
	} else {
		goto L37
	}
L9:
	;
	v208 = int32(0)
	goto L1
L10:
	;
	goto L11
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v26 = F_copyObjectImpl(m, l2)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v85 = F_ExecPrepareExprList(m, v26, l3)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L13
	} else {
		goto L25
	}
L13:
	;
	return int32(0)
L14:
	;
	if v26 == int32(0) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v32 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v33 <= v32 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v38 = v32
	goto L17
L17:
	;
	v48 = v38 << (uint(int32(2)) % 32)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v25+v48)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v52 = v51 + v48
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v55 = F_transformExpr(m, l0, v53, int32(36))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L13
	} else {
		goto L19
	}
L18:
	;
	goto L12
L19:
	;
	v57 = F_exprType(m, v55)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v59 = int32(-1)
	v63 = F_coerce_to_target_type(m, l0, v55, v57, v50, v59, int32(1), int32(2), v59)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	if v63 == int32(0) {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	F_assign_expr_collations(m, l0, v63)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v63
	v71 = v38 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v71 < v72 {
		v38 = v71
		goto L17
	} else {
		goto L24
	}
L24:
	;
	goto L18
L25:
	;
	v87 = F_makeParamList(m, v17)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L13
	} else {
		goto L26
	}
L26:
	;
	if v85 == int32(0) {
		v208 = v87
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v91 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v92 <= v91 {
		v208 = v87
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v97 = v91
	goto L29
L29:
	;
	v109 = v97 << (uint(int32(2)) % 32)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109+v110)))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v109+v25)))
	v117 = v87 + int32(32) + v97*int32(12)
	v118 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v117)+6)) = uint16(v118)
	*(*int32)(unsafe.Add(mBase, uint32(v117)+8)) = v114
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l3)+152))
	if v121 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v208 = v87
	goto L1
L31:
	;
	v124 = F_MakePerTupleExprContext(m, l3)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L13
	} else {
		goto L34
	}
L32:
	;
	v126 = v121
	goto L33
L33:
	;
	v127 = int32(4476144)
	v128 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v130
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	v135 = m.T0[v134].(func(*base.Module, int32, int32, int32) int32)(m, v112, v126, v117+int32(4))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L13
	} else {
		goto L35
	}
L34:
	;
	v126 = v124
	goto L33
L35:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = v135
	v141 = v97 + int32(1)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v141 < v142 {
		v97 = v141
		goto L29
	} else {
		goto L36
	}
L36:
	;
	goto L30
L37:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L13
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l1
	F_errmsg(m, int32(677794), v14+int32(32))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L13
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v17
	F_errdetail(m, int32(631445), v14+int32(16))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L13
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(493302), int32(298), int32(149579))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L13
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L13
	} else {
		goto L43
	}
L43:
	;
	v176 = F_format_type_be(m, v57)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	v178 = F_format_type_be(m, v50)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v38 + int32(1)
	F_errmsg(m, int32(191282), v14)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	F_errhint(m, int32(596351), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v193 = F_exprLocation(m, v192)
	mBase = m.M
	F_parser_errposition(m, l0, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(493302), int32(335), int32(149579))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L13
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecASUpdateTriggers(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v5 == int32(0) {
		return
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+17)))
		if v8 != int32(1) {
			return
		} else {
			v11 = int32(0)
			v18 = F_ExecGetAllUpdatedCols(m, l1, l0)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				F_AfterTriggerSaveEvent(m, l0, l1, v11, v11, int32(2), v11, v11, v11, v11, v18, l2, int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_ExecBRDeleteTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
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
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v17 = F_ExecGetTriggerOldSlot(m, l0, l2)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v22 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = v22
	v30 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+19)) = uint8(v30)
	if l4 == v30 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v15 - int32(-64)
	return v149
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+20)) = int64(55834575290)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v64
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v67 <= int32(0) {
		v132 = int32(1)
		goto L17
	} else {
		goto L18
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(0)
	v41 = F_GetTupleForTrigger(m, l0, l1, l2, l3, int32(3), v17, l8^int32(1), v13+int32(-52), l6, l7)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_ExecForceStoreHeapTuple(m, l4, v17, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L16
	}
L8:
	;
	v55 = F_ExecFetchSlotHeapTuple(m, v17, int32(1), v13+int32(-45))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L15
	}
L9:
	;
	if v41 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if l5 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v149 = int32(0)
	goto L3
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v45 == int32(0) {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v45
	goto L12
L15:
	;
	v60 = v55
	goto L4
L16:
	;
	v60 = l4
	goto L4
L17:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+19)))
	if v137 != int32(1) {
		v149 = v132
		goto L3
	} else {
		goto L36
	}
L18:
	;
	v79 = int32(0)
	goto L19
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v86 = v83 + v79*int32(60)
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+12)))
	if v87&int32(75) != int32(11) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v132 = v120
	goto L17
L21:
	;
	v120 = int32(1)
	v122 = v79 + v120
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v122 < v123 {
		v79 = v122
		goto L19
	} else {
		goto L35
	}
L22:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v93 = int32(0)
	v95 = F_TriggerEnabled(m, l0, l2, v86, v92, v93, v17, v93)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v95 == int32(0) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v17
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v106 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v109 = v106
	goto L27
L26:
	;
	v107 = F_MakePerTupleExprContext(m, l0)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+20))
	v111 = F_ExecCallTriggerFunc(m, v13+int32(-44), v79, v104, v105, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v109 = v107
	goto L27
L29:
	;
	if v111 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v132 = int32(0)
	goto L17
L31:
	;
	goto L32
L32:
	;
	if v60 == v111 {
		goto L21
	} else {
		goto L33
	}
L33:
	;
	F_pfree(m, v111)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L21
L35:
	;
	goto L20
L36:
	;
	F_pfree(m, v60)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v149 = v132
	goto L3
}
func F_ExecBRInsertTriggers(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	v4 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v20 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+60)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v17)+52)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v17)+44)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v17)+36)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = int64(51539607994)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v30
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v33 <= v4 {
		v211 = int32(1)
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L8
	} else {
		goto L51
	}
L2:
	;
	m.G0 = v17 + int32(80)
	return v211
L3:
	;
	v41 = v4
	v44 = v4
	goto L4
L4:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v53 = v50 + v44*int32(60)
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)))
	if v54&int32(71) != int32(7) {
		v193 = v41
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v211 = v202
	goto L2
L6:
	;
	v202 = int32(1)
	v204 = v44 + v202
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v204 < v205 {
		v41 = v193
		v44 = v204
		goto L4
	} else {
		goto L50
	}
L7:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v60 = int32(0)
	v62 = F_TriggerEnabled(m, l0, l1, v53, v59, v60, v60, l2)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	if v62 == int32(0) {
		v193 = v41
		goto L6
	} else {
		goto L10
	}
L10:
	;
	if v41 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v73 = F_ExecFetchSlotHeapTuple(m, l2, int32(1), v17+int32(70))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L8
	} else {
		goto L14
	}
L12:
	;
	v75 = v41
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = l2
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v83 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v75 = v73
	goto L13
L15:
	;
	v86 = v83
	goto L17
L16:
	;
	v84 = F_MakePerTupleExprContext(m, l0)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L8
	} else {
		goto L18
	}
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v88 = F_ExecCallTriggerFunc(m, v17+int32(24), v44, v81, v82, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L8
	} else {
		goto L19
	}
L18:
	;
	v86 = v84
	goto L17
L19:
	;
	if v88 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v92 = int32(0)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+70)))
	if v93 != int32(1) {
		v211 = v92
		goto L2
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v75 == v88 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	F_pfree(m, v75)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	v211 = v92
	goto L2
L25:
	;
	v193 = v88
	goto L6
L26:
	;
	goto L27
L27:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+52))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+16))
	if v101 == int32(0) {
		v163 = v88
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_ExecForceStoreHeapTuple(m, v163, l2, int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L8
	} else {
		goto L40
	}
L29:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+18)))
	if v104 != int32(1) {
		v163 = v88
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v107 <= int32(0) {
		v163 = v88
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v117 = int32(0)
	v119 = v88
	v124 = v107
	goto L32
L32:
	;
	v130 = v117 + int32(1)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117*int32(100)+(v100+int32(110)+v124<<(uint(int32(4))%32))))))
	if v135 != int32(118) {
		v154 = v119
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v163 = v154
	goto L28
L34:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v130 < v155 {
		v117 = v130
		v119 = v154
		v124 = v155
		goto L32
	} else {
		goto L39
	}
L35:
	;
	v138 = F_heap_attisnull(m, v119, v130, v100)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	if v138 != 0 {
		v154 = v119
		goto L34
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(0)
	v143 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+71)) = uint8(v143)
	v152 = F_heap_modify_tuple_by_cols(m, v119, v100, v143, v17+int32(76), v17+int32(72), v17+int32(71))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v154 = v152
	goto L34
L39:
	;
	goto L33
L40:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+16)))
	if v174 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v178 = F_ExecPartitionCheck(m, l1, l2, l0, int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L8
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+70)))
	if v182 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	if v178 == int32(0) {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	F_pfree(m, v75)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L8
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v193 = int32(0)
	goto L6
L49:
	;
	goto L48
L50:
	;
	goto L5
L51:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L8
	} else {
		goto L52
	}
L52:
	;
	F_errmsg(m, int32(437677), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+48))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+68))
	v240 = F_get_namespace_name(m, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L8
	} else {
		goto L54
	}
L54:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v236
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v243 + int32(4)
	F_errdetail(m, int32(642799), v17)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(490070), int32(2530), int32(133276))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecBRUpdateTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int64
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v340 int32
	_ = v340
	v10 = int32(0)
	v18 = m.G0
	v20 = v18 + int32(-64)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v23 = F_ExecGetTriggerOldSlot(m, l0, l2)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+54)) = uint8(v27)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+53)) = uint8(v27)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v27
	v33 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+36)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v20)+28)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v20)+20)) = v33
	v39 = F_ExecUpdateLockMode(m, l0, l2)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l4 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	m.G0 = v20 - int32(-64)
	return v340
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = int64(60129542586)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v109
	v111 = F_ExecGetAllUpdatedCols(m, l2, l0)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L25
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = int32(0)
	v49 = F_GetTupleForTrigger(m, l0, l1, l2, l3, v39, v23, l8^int32(1), v18+int32(-4), l6, l7)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_ExecForceStoreHeapTuple(m, l4, v23, int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L24
	}
L9:
	;
	if v49 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	if v51 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v340 = v10
	goto L4
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v53)+12)) = v51
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+72))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	m.T0[v59].(func(*base.Module, int32))(m, v57)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v98 = F_ExecFetchSlotHeapTuple(m, v23, int32(1), v18+int32(-10))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L23
	}
L16:
	;
	v62 = int32(4476144)
	v63 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v65
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int32)(m, v52+int32(4), v56, int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v63
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
	v77 = v75 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)) = uint16(v77)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v80)
	if v57 != l5 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+32))
	m.T0[v84].(func(*base.Module, int32, int32))(m, l5, v57)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+28))
	m.T0[v88].(func(*base.Module, int32))(m, l5)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	goto L15
L23:
	;
	v105 = v98
	goto L5
L24:
	;
	v105 = l4
	goto L5
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v111
	v114 = int32(1)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if int32(0) < v115 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v126 = int32(0)
	v130 = v10
	goto L29
L27:
	;
	goto L28
L28:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+54)))
	if v321 != int32(1) {
		v340 = v114
		goto L4
	} else {
		goto L78
	}
L29:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v139 = v136 + v126*int32(60)
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+12)))
	if v140&int32(83) != int32(19) {
		v294 = v130
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L28
L31:
	;
	v301 = v126 + int32(1)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v301 < v302 {
		v126 = v301
		v130 = v294
		goto L29
	} else {
		goto L77
	}
L32:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v146 = F_TriggerEnabled(m, l0, l2, v139, v145, v111, v23, l5)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v146 == int32(0) {
		v294 = v130
		goto L31
	} else {
		goto L34
	}
L34:
	;
	if v130 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v155 = F_ExecFetchSlotHeapTuple(m, l5, int32(1), v18+int32(-11))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	v157 = v130
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v139
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v167 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v157 = v155
	goto L37
L39:
	;
	v170 = v167
	goto L41
L40:
	;
	v168 = F_MakePerTupleExprContext(m, l0)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+20))
	v172 = F_ExecCallTriggerFunc(m, v18+int32(-56), v126, v165, v166, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	v170 = v168
	goto L41
L43:
	;
	if v172 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+54)))
	if v176 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	if v172 == v157 {
		goto L53
	} else {
		goto L54
	}
L47:
	;
	F_pfree(m, v105)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v181 = int32(0)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+53)))
	if v182 != int32(1) {
		v340 = v181
		goto L4
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	F_pfree(m, v157)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v340 = v181
	goto L4
L53:
	;
	v294 = v172
	goto L31
L54:
	;
	goto L55
L55:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+52))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+16))
	if v190 == int32(0) {
		v259 = v172
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_ExecForceStoreHeapTuple(m, v259, l5, int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L68
	}
L57:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+18)))
	if v193 != int32(1) {
		v259 = v172
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	if v196 <= int32(0) {
		v259 = v172
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v208 = v196
	v210 = int32(0)
	v212 = v172
	goto L60
L60:
	;
	v222 = v210 + int32(1)
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210*int32(100)+(v189+int32(110)+v208<<(uint(int32(4))%32))))))
	if v227 != int32(118) {
		v246 = v212
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v259 = v246
	goto L56
L62:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	if v222 < v247 {
		v208 = v247
		v210 = v222
		v212 = v246
		goto L60
	} else {
		goto L67
	}
L63:
	;
	v230 = F_heap_attisnull(m, v212, v222, v189)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if v230 != 0 {
		v246 = v212
		goto L62
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = int32(0)
	v235 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+55)) = uint8(v235)
	v244 = F_heap_modify_tuple_by_cols(m, v212, v189, v235, v18+int32(-4), v18+int32(-8), v18+int32(-9))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v246 = v244
	goto L62
L67:
	;
	goto L61
L68:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+54)))
	if v269 != int32(1) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+53)))
	if v277 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	if v105 != v259 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+28))
	m.T0[v274].(func(*base.Module, int32))(m, l5)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	F_pfree(m, v157)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v294 = int32(0)
	goto L31
L76:
	;
	goto L75
L77:
	;
	goto L30
L78:
	;
	F_pfree(m, v105)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v340 = v114
	goto L4
}
func F_ExecBuildHash32Expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v294 int64
	_ = v294
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	v7 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(16)
	m.G0 = v27
	v30 = F_palloc0(m, int32(68))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(380)
	if l2 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v37 = v36
	goto L5
L4:
	;
	v37 = v7
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+40)) = l4
	v39 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+8)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v27))) = v39
	v43 = F_expr_setup_walker(m, l2, v27)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	F_ExecPushExprSetupSteps(m, v30, v27)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if int64(2) <= base.I64_extend_i32_s(v37) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v51 = F_palloc(m, int32(8))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v53 = v7
	goto L10
L10:
	;
	v72 = int32(0)
	v73 = v7
	v76 = int32(86)
	v78 = v7
	v79 = v7
	v80 = int32(87)
	v81 = v7
	v83 = v7
	v88 = v7
	goto L14
L11:
	;
	v53 = v51
	goto L10
L12:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v287 + int32(1)
	v293 = v286 + v287*int32(40)
	v294 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v293))) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v293)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v293)+32)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v293)+28)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v293)+24)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v293)+20)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v293)+16)) = v79
	*(*int64)(unsafe.Add(mBase, uint32(v293)+8)) = v294
	v305 = F_jit_compile_expr(m, v30)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L65
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v284
	v286 = v284
	goto L12
L14:
	;
	v89 = int32(0)
	if l2 == v89 {
		v99 = v89
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
	if v273 != v181 {
		goto L61
	} else {
		goto L62
	}
L16:
	;
	if l1 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v93 <= v73 {
		v99 = int32(0)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v99 = v95 + v73<<(uint(int32(2))%32)
	goto L16
L19:
	;
	goto L15
L20:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0+v107)))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v192 = F_palloc0(m, int32(28))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L34
	}
L21:
	;
	if v78 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v102 <= v73 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	if v99 == int32(0) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v107 = v73 << (uint(int32(2)) % 32)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v109 = v107 + v108
	if v109 != 0 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v30)+36))
	if v181 != 0 {
		goto L19
	} else {
		goto L32
	}
L27:
	;
	v114 = int32(0)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v115 <= v114 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v122 = v114
	goto L29
L29:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v143+v122<<(uint(int32(2))%32))))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v142+v147*int32(40))+28)) = v151
	v154 = v122 + int32(1)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v154 < v155 {
		v122 = v154
		goto L29
	} else {
		goto L31
	}
L30:
	;
	goto L26
L31:
	;
	goto L30
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+36)) = int32(16)
	v185 = F_palloc(m, int32(640))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v284 = v185
	goto L13
L34:
	;
	v195 = F_palloc0(m, int32(28))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_fmgr_info(m, v188, v192)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_ExecInitExprRec(m, v190, v30, v195+int32(20), v195+int32(24))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v205 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v205)
	v207 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v195)+16)) = uint8(v207)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+12)) = v189
	*(*int64)(unsafe.Add(mBase, uint32(v195)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v192
	v213 = base.B2i32(v73 == v37-int32(1))
	if v73 == v37-int32(1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v214 = v30 + int32(8)
	goto L40
L39:
	;
	v214 = v53
	goto L40
L40:
	;
	if v73 == v37-int32(1) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v215 = v30 + int32(5)
	goto L43
L42:
	;
	v215 = v53 + int32(4)
	goto L43
L43:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v73))))
	if v217 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v218 = v80
	goto L46
L45:
	;
	v218 = v76
	goto L46
L46:
	;
	if l5 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v219 = v76
	goto L49
L48:
	;
	v219 = v218
	goto L49
L49:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v30)+36))
	if v221 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
	v244 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v243 + v244
	v249 = v242 + v243*int32(40)
	v250 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v249)+36)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(v249)+32)) = v53
	v253 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v249)+28)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v249)+24)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v249)+20)) = v195
	*(*int32)(unsafe.Add(mBase, uint32(v249)+16)) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v249)+12)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(v249)+8)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v249)+4)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v249))) = v219
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
	v271 = F_lappend_int(m, v78, v268-v244)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L60
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v240
	v242 = v240
	goto L50
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+36)) = int32(16)
	v227 = F_palloc(m, int32(640))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
	if v229 != v221 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v240 = v227
	goto L51
L56:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v242 = v231
	goto L50
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+36)) = v221 << (uint(int32(1)) % 32)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v238 = F_repalloc(m, v235, v221*int32(80))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v240 = v238
	goto L51
L60:
	;
	v72 = v195
	v73 = v73 + v244
	v76 = int32(88)
	v78 = v271
	v79 = v192
	v80 = int32(89)
	v81 = v220
	v83 = v53
	v88 = v253
	goto L14
L61:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v286 = v275
	goto L12
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+36)) = v181 << (uint(int32(1)) % 32)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v282 = F_repalloc(m, v279, v181*int32(80))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v284 = v282
	goto L13
L65:
	;
	if v305 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	F_ExecReadyInterpretedExpr(m, v30)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	m.G0 = v27 + int32(16)
	return v30
L69:
	;
	goto L68
}
func F_ExecCheckPermissions(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	v4 = int32(0)
	if l1 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v73
L2:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	if v61 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v8 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v14 = v4
	goto L5
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16+v14<<(uint(int32(2))%32))))
	v21 = F_ExecCheckOneRelPerms(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v29 = int32(0)
	if l2 == v29 {
		v73 = v29
		goto L1
	} else {
		goto L13
	}
L7:
	;
	return int32(0)
L8:
	;
	if v21 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v26 = v14 + int32(1)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v26 < v27 {
		v14 = v26
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	goto L6
L12:
	;
	goto L2
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v34 = F_get_rel_relkind(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	switch v34 - int32(73) {
	case 0, 32:
		goto L21
	default:
		v45 = int32(41)
		goto L16
	case 10:
		goto L20
	case 29:
		goto L17
	case 36:
		goto L18
	case 45:
		goto L19
	}
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v49 = F_get_rel_name(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L7
	} else {
		goto L22
	}
L16:
	;
	v47 = v45
	goto L15
L17:
	;
	v45 = int32(18)
	goto L16
L18:
	;
	v47 = int32(23)
	goto L15
L19:
	;
	v47 = int32(51)
	goto L15
L20:
	;
	v47 = int32(37)
	goto L15
L21:
	;
	v47 = int32(20)
	goto L15
L22:
	;
	F_aclcheck_error(m, int32(1), v47, v49)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	return int32(0)
L24:
	;
	return int32(1)
L25:
	;
	goto L26
L26:
	;
	v66 = m.T0[v61].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	v73 = v66
	goto L1
}
func F_ExecCheckPlanOutput(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
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
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if l1 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L20
	} else {
		goto L55
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L20
	} else {
		goto L47
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L20
	} else {
		goto L42
	}
L4:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v110 != v118 {
		goto L1
	} else {
		goto L41
	}
L5:
	;
	v110 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v17 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v18 <= v17 {
		v110 = v17
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v23 = v17
	goto L9
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v31 <= v23 {
		goto L3
	} else {
		goto L11
	}
L10:
	;
	v110 = v41
	goto L4
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v23<<(uint(int32(2))%32))))
	v40 = int32(1)
	v41 = v23 + v40
	v45 = v23*int32(100) + (v13 + int32(20) + v31<<(uint(int32(4))%32))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+91)))
	if v46 == v40 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v41 < v108 {
		v23 = v41
		goto L9
	} else {
		goto L40
	}
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v50 == int32(7) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+90)))
	if v75 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+24)))
	if v53 != 0 {
		goto L12
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	return
L21:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errmsg(m, int32(322156), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v41
	F_errdetail(m, int32(632328), v11)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(493597), int32(232), int32(64203))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L20
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
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v76 == int32(7) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v102 = F_exprType(m, v74)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L20
	} else {
		goto L38
	}
L29:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+24)))
	if v79 != 0 {
		goto L12
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L20
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L20
	} else {
		goto L34
	}
L34:
	;
	F_errmsg(m, int32(322156), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L20
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v41
	F_errdetail(m, int32(632258), v11+int32(32))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L20
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(493597), int32(249), int32(64203))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L20
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v45)+68))
	if v102 != v104 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	goto L12
L40:
	;
	goto L10
L41:
	;
	m.G0 = v11 + int32(48)
	return
L42:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L20
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(322156), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L20
	} else {
		goto L44
	}
L44:
	;
	F_errdetail(m, int32(570703), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L20
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(493597), int32(212), int32(64203))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L20
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L20
	} else {
		goto L48
	}
L48:
	;
	F_errmsg(m, int32(322156), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L20
	} else {
		goto L49
	}
L49:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v45)+68))
	v155 = F_format_type_be(m, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L20
	} else {
		goto L50
	}
L50:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v158 = F_exprType(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L20
	} else {
		goto L51
	}
L51:
	;
	v160 = F_format_type_be(m, v158)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L20
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v155
	F_errdetail(m, int32(582419), v11+int32(16))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L20
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(493597), int32(261), int32(64203))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L20
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L20
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(322156), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L20
	} else {
		goto L57
	}
L57:
	;
	F_errdetail(m, int32(570818), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L20
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(493597), int32(268), int32(64203))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L20
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecConstraints(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
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
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+16)))
	if v20 != int32(1) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v15 + int32(48)
	return
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L15
	} else {
		goto L87
	}
L3:
	;
	F_ReportNotNullViolationError(m, l0, l1, l2, v76)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L15
	} else {
		goto L86
	}
L4:
	;
	F_ReportNotNullViolationError(m, l0, l1, l2, v33)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L15
	} else {
		goto L85
	}
L5:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	v91 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+122)))
	if v91 <= int32(0) {
		goto L1
	} else {
		goto L26
	}
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v23 <= int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v28 = int32(1)
	v33 = v28
	v34 = v23
	v35 = v28
	v38 = int32(0)
	goto L8
L8:
	;
	v47 = v18 - int32(80) + v34<<(uint(int32(4))%32) + v33*int32(100)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+86)))
	if v48 != int32(1) {
		v68 = v38
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v68 == int32(0) {
		goto L5
	} else {
		goto L23
	}
L10:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v71 = v35 + int32(1)
	v72 = base.I32_extend16_s(v71)
	if v72 <= v69 {
		v33 = v72
		v34 = v69
		v35 = v71
		v38 = v68
		goto L8
	} else {
		goto L22
	}
L11:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+90)))
	if v51 == int32(118) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v54 = F_lappend_int(m, v38, v33)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v56 < base.I32_extend16_s(v35) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	return
L16:
	;
	v68 = v54
	goto L10
L17:
	;
	F_slot_getsomeattrs_int(m, l1, v33)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v63 = int32(1)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v33-v63))))
	if v65 == v63 {
		goto L4
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v68 = v38
	goto L10
L22:
	;
	goto L9
L23:
	;
	v76 = F_ExecRelGenVirtualNotNull(m, l0, l1, l2, v68)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	if v76 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	goto L5
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+52))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96)+14)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)+48))
	v99 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98)+122)))
	if v97 != v99 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v102 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v106 = int32(4476144)
	v107 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l2)+100))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v109
	v113 = F_palloc0(m, v97<<(uint(int32(2))%32))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L15
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l2)+152))
	if v177 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v113
	if v97 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v119 = int32(0)
	goto L35
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v107
	goto L30
L35:
	;
	v130 = v101 + v119*int32(12)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+8)))
	if v131 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L34
L37:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v135 = F_stringToNode(m, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L15
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v149 = v119 + int32(1)
	if v149 != v97 {
		v119 = v149
		goto L35
	} else {
		goto L43
	}
L40:
	;
	v138 = F_expand_generated_columns_in_expr(m, v135, v94, int32(1))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L15
	} else {
		goto L41
	}
L41:
	;
	v140 = F_ExecPrepareExpr(m, v138, l2)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L15
	} else {
		goto L42
	}
L42:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v142+v119<<(uint(int32(2))%32)))) = v140
	goto L39
L43:
	;
	goto L36
L44:
	;
	v180 = F_MakePerTupleExprContext(m, l2)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L15
	} else {
		goto L47
	}
L45:
	;
	v182 = v177
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+4)) = l1
	if v97 == int32(0) {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	v182 = v180
	goto L46
L48:
	;
	v190 = int32(0)
	goto L49
L49:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v199+v190<<(uint(int32(2))%32))))
	if v203 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v101+v190*int32(12))))
	if v214 == int32(0) {
		goto L1
	} else {
		goto L58
	}
L51:
	;
	goto L50
L52:
	;
	v204 = F_ExecCheck(m, v203, v182)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L15
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v209 = v190 + int32(1)
	if v209 != v97 {
		v190 = v209
		goto L49
	} else {
		goto L57
	}
L55:
	;
	if v204 == int32(0) {
		goto L51
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	goto L1
L58:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v217 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+56))
	v248 = F_ExecBuildSlotValueDescription(m, v247, v243, v245, v244)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L15
	} else {
		goto L75
	}
L60:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+52))
	v222 = F_build_attrmap_by_name_if_req(m, v218, v220, int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L15
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v237 = F_ExecGetInsertedCols(m, l0, l2)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L15
	} else {
		goto L72
	}
L63:
	;
	if v222 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v225 = F_MakeTupleTableSlot(m, v220, int32(1596068))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L15
	} else {
		goto L67
	}
L65:
	;
	v229 = l1
	goto L66
L66:
	;
	v230 = F_ExecGetInsertedCols(m, v217, l2)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L15
	} else {
		goto L69
	}
L67:
	;
	v227 = F_execute_attr_map_slot(m, v222, l1, v225)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L15
	} else {
		goto L68
	}
L68:
	;
	v229 = v227
	goto L66
L69:
	;
	v232 = F_ExecGetUpdatedCols(m, v217, l2)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L15
	} else {
		goto L70
	}
L70:
	;
	v234 = F_bms_union(m, v230, v232)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L15
	} else {
		goto L71
	}
L71:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	v243 = v229
	v244 = v234
	v245 = v220
	v246 = v236
	goto L59
L72:
	;
	v239 = F_ExecGetUpdatedCols(m, l0, l2)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L15
	} else {
		goto L73
	}
L73:
	;
	v241 = F_bms_union(m, v237, v239)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L15
	} else {
		goto L74
	}
L74:
	;
	v243 = l1
	v244 = v241
	v245 = v18
	v246 = v17
	goto L59
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L15
	} else {
		goto L76
	}
L76:
	;
	F_errcode(m, int32(67391682))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L15
	} else {
		goto L77
	}
L77:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v257 + int32(4)
	F_errmsg(m, int32(677375), v15+int32(16))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L15
	} else {
		goto L78
	}
L78:
	;
	if v248 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v248
	F_errdetail(m, int32(582592), v15)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L15
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	F_errtableconstraint(m, v17, v214)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L15
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	F_errfinish(m, int32(491223), int32(2081), int32(119120))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L15
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v94)+48))
	v287 = int32(*(*int16)(unsafe.Add(mBase, uint32(v286)+122)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v286 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v287 - v97
	F_errmsg_internal(m, int32(683192), v15+int32(32))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L15
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(491223), int32(1798), int32(315098))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L15
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecCustomScan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v3 != 0 {
		F_ProcessInterrupts(m)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return int32(0)
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			v10 = m.T0[v9].(func(*base.Module, int32) int32)(m, l0)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		v10 = m.T0[v9].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_ExecFilterJunk(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v12 < v11 {
		F_slot_getsomeattrs_int(m, l1, v11)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
			m.T0[v25].(func(*base.Module, int32))(m, v23)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if int32(0) < v22 {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
					v34 = int32(0)
					for {
						v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18+v34<<(uint(int32(1))%32)))))
						if v46 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v31+v34<<(uint(int32(2))%32)))) = int32(0)
							v68 = int32(1)
						} else {
							v55 = int32(2)
							v59 = v46 - int32(1)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v20+v59<<(uint(v55)%32))))
							*(*int32)(unsafe.Add(mBase, uint32(v31+v34<<(uint(v55)%32)))) = v63
							v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v19))))
							v68 = v66
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v34+v30))) = uint8(v68)
						v71 = v34 + int32(1)
						if v71 != v22 {
							v34 = v71
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
				v84 = v82 & int32(65533)
				*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)) = uint16(v84)
				v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
				*(*uint16)(unsafe.Add(mBase, uint32(v23)+6)) = uint16(v87)
				return v23
			}
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
		m.T0[v25].(func(*base.Module, int32))(m, v23)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			if int32(0) < v22 {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
				v34 = int32(0)
				for {
					v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18+v34<<(uint(int32(1))%32)))))
					if v46 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v31+v34<<(uint(int32(2))%32)))) = int32(0)
						v68 = int32(1)
					} else {
						v55 = int32(2)
						v59 = v46 - int32(1)
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v20+v59<<(uint(v55)%32))))
						*(*int32)(unsafe.Add(mBase, uint32(v31+v34<<(uint(v55)%32)))) = v63
						v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v19))))
						v68 = v66
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v34+v30))) = uint8(v68)
					v71 = v34 + int32(1)
					if v71 != v22 {
						v34 = v71
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
			v84 = v82 & int32(65533)
			*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)) = uint16(v84)
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
			*(*uint16)(unsafe.Add(mBase, uint32(v23)+6)) = uint16(v87)
			return v23
		}
	}
}
func F_ExecGetUpdatedCols(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v8 != 0 {
		v9 = v8
	} else {
		v9 = l0
	}
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v10 == int32(0) {
		v80 = v3
		return v80
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v15+v10<<(uint(int32(2))%32)-int32(4))))
		v22 = F_getRTEPermissionInfo(m, v13, v21)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			if v22 == int32(0) {
				v80 = v3
				return v80
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				if v28 == int32(0) {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
					v80 = v77
					return v80
				} else {
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)))
					if v31 == int32(0) {
						v34 = int32(4476144)
						v35 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v41
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)+48))
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+131)))
						v49 = F_build_attrmap_by_name_if_req(m, v39, v37, (v44^int32(-1))&int32(1))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							if v49 != 0 {
								v51 = F_convert_tuples_by_name_attrmap(m, v39, v37, v49)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v51
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v35
									v56 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v56)
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
									if v63 == int32(0) {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
										v80 = v77
										return v80
									} else {
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
										v68 = F_execute_attr_map_cols(m, v66, v67)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											return v68
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v35
								v56 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v56)
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
								if v63 == int32(0) {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
									v80 = v77
									return v80
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
									v68 = F_execute_attr_map_cols(m, v66, v67)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										return v68
									}
								}
							}
						}
					} else {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
						if v63 == int32(0) {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
							v80 = v77
							return v80
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
							v68 = F_execute_attr_map_cols(m, v66, v67)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								return v68
							}
						}
					}
				}
			}
		}
	}
}
func F_ExecGrantStmt_oids(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v144 int64
	_ = v144
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v181 int64
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int64
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int64
	_ = v331
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int64
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v371 int64
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v448 int32
	_ = v448
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int64
	_ = v478
	var v479 int32
	_ = v479
	var v484 int64
	_ = v484
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v512 int64
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v524 int32
	_ = v524
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	var v565 int64
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v790 int64
	_ = v790
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int64
	_ = v801
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int64
	_ = v939
	var v942 int32
	_ = v942
	var v946 int64
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1107 int32
	_ = v1107
	var v1112 int32
	_ = v1112
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1135 int64
	_ = v1135
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1156 int32
	_ = v1156
	var v1177 int32
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1218 int64
	_ = v1218
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int64
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int64
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1234 int64
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1259 int32
	_ = v1259
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1284 int32
	_ = v1284
	var v1288 int32
	_ = v1288
	var v1292 int32
	_ = v1292
	var v1296 int32
	_ = v1296
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1331 int32
	_ = v1331
	var v1337 int32
	_ = v1337
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1383 int32
	_ = v1383
	var v1387 int32
	_ = v1387
	var v1393 int32
	_ = v1393
	var v1398 int32
	_ = v1398
	var v1403 int32
	_ = v1403
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1412 int64
	_ = v1412
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1443 int32
	_ = v1443
	var v1455 int32
	_ = v1455
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1520 int64
	_ = v1520
	var v1526 int32
	_ = v1526
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1537 int64
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1539 int64
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1544 int32
	_ = v1544
	var v1546 int64
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1562 int32
	_ = v1562
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1574 int32
	_ = v1574
	var v1576 int32
	_ = v1576
	var v1580 int32
	_ = v1580
	var v1586 int32
	_ = v1586
	var v1588 int32
	_ = v1588
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1597 int32
	_ = v1597
	var v1599 int32
	_ = v1599
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1632 int32
	_ = v1632
	var v1636 int32
	_ = v1636
	var v1642 int32
	_ = v1642
	var v1647 int32
	_ = v1647
	var v1652 int32
	_ = v1652
	var v1657 int32
	_ = v1657
	var v1662 int32
	_ = v1662
	var v1667 int32
	_ = v1667
	var v1672 int32
	_ = v1672
	var v1699 int32
	_ = v1699
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1718 int32
	_ = v1718
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1723 int64
	_ = v1723
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1729 int64
	_ = v1729
	var v1731 int64
	_ = v1731
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1736 int32
	_ = v1736
	var v1737 int64
	_ = v1737
	var v1739 int64
	_ = v1739
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1751 int32
	_ = v1751
	var v1754 int32
	_ = v1754
	var v1760 int32
	_ = v1760
	var v1770 int32
	_ = v1770
	var v1784 int32
	_ = v1784
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1825 int32
	_ = v1825
	var v1826 int32
	_ = v1826
	var v1830 int32
	_ = v1830
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1841 int32
	_ = v1841
	v2 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(512)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v31 - int32(9) {
	case 0:
		goto L12
	default:
		goto L11
	case 3, 40:
		goto L2
	case 7:
		goto L3
	case 8:
		goto L4
	case 10, 20, 25:
		goto L5
	case 12:
		goto L6
	case 13:
		goto L7
	case 18:
		goto L10
	case 27:
		goto L8
	case 28, 32:
		goto L13
	case 33:
		goto L9
	}
L1:
	;
	v1699 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L382
L2:
	;
	F_ExecGrant_common(m, l0, int32(1247), int64(256), int32(461))
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L14
	} else {
		goto L381
	}
L3:
	;
	F_ExecGrant_common(m, l0, int32(2328), int64(256), int32(0))
	mBase = m.M
	v1667 = m.ExcPending
	if v1667 != 0 {
		goto L14
	} else {
		goto L380
	}
L4:
	;
	F_ExecGrant_common(m, l0, int32(1417), int64(256), int32(0))
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		goto L14
	} else {
		goto L379
	}
L5:
	;
	F_ExecGrant_common(m, l0, int32(1255), int64(128), int32(0))
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L14
	} else {
		goto L378
	}
L6:
	;
	F_ExecGrant_common(m, l0, int32(2612), int64(256), int32(460))
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		goto L14
	} else {
		goto L377
	}
L7:
	;
	v1409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v1409 != int32(1) {
		goto L333
	} else {
		goto L334
	}
L8:
	;
	F_ExecGrant_common(m, l0, int32(2615), int64(768), int32(0))
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L14
	} else {
		goto L332
	}
L9:
	;
	F_ExecGrant_common(m, l0, int32(1213), int64(512), int32(0))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L14
	} else {
		goto L331
	}
L10:
	;
	v1132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v1132 != int32(1) {
		goto L264
	} else {
		goto L265
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L14
	} else {
		goto L261
	}
L12:
	;
	F_ExecGrant_common(m, l0, int32(1262), int64(3584), int32(0))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L14
	} else {
		goto L260
	}
L13:
	;
	v36 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	v40 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v42 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L14
	} else {
		goto L257
	}
L18:
	;
	F_sequence_close(m, v40, int32(3))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L14
	} else {
		goto L255
	}
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v45 <= int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v70 = v2
	goto L21
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v70<<(uint(int32(2))%32))))
	v80 = F_SearchSysCacheLocked1(m, int32(57), v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L14
	} else {
		goto L30
	}
L22:
	;
	goto L18
L23:
	;
	if v182 < int32(-7) {
		goto L183
	} else {
		goto L184
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L14
	} else {
		goto L180
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L14
	} else {
		goto L175
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L14
	} else {
		goto L170
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L14
	} else {
		goto L167
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L14
	} else {
		goto L163
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L14
	} else {
		goto L159
	}
L30:
	;
	if v80 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+22)))
	v84 = v82 + v83
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+119)))
	switch v85 - int32(99) {
	case 0:
		goto L35
	case 1, 2, 3, 4, 5:
		goto L34
	case 6:
		goto L36
	default:
		goto L37
	}
L32:
	;
	goto L33
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L14
	} else {
		goto L156
	}
L34:
	;
	v131 = base.B2i32(v85 == int32(83))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.B2i32(v131 == int32(0))&base.B2i32(v134 == int32(37)) != 0 {
		goto L29
	} else {
		goto L47
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L14
	} else {
		goto L43
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L14
	} else {
		goto L39
	}
L37:
	;
	if v85 != int32(73) {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+160)) = v84 + int32(4)
	F_errmsg(m, int32(28306), v29+int32(160))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(492197), int32(1823), int32(261753))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L14
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L14
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+176)) = v84 + int32(4)
	F_errmsg(m, int32(365824), v29+int32(176))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L14
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(492197), int32(1830), int32(261753))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L14
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	if v85 == int32(83) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v140 = int64(262)
	goto L50
L49:
	;
	v140 = int64(16511)
	goto L50
L50:
	;
	v141 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v141 == int64(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v144 = v140
	goto L53
L52:
	;
	v144 = v141
	goto L53
L53:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v145 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v146 = v144
	goto L56
L55:
	;
	v146 = v141
	goto L56
L56:
	;
	if v134 != int32(41) {
		v181 = v146
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v182 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84)+120)))
	v184 = v182 + int32(8)
	v187 = F_palloc0(m, v184<<(uint(int32(3))%32))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L14
	} else {
		goto L71
	}
L58:
	;
	if v85 == int32(83) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	if v146&int64(-263) == int64(0) {
		v181 = v146
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v146&int64(-16512) != int64(0) {
		goto L28
	} else {
		goto L70
	}
L62:
	;
	v157 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L14
	} else {
		goto L63
	}
L63:
	;
	if v157 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L14
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v181 = v146 & int64(262)
	goto L57
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+112)) = v84 + int32(4)
	F_errmsg(m, int32(168352), v29+int32(112))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L14
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(492197), int32(1876), int32(261753))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L14
	} else {
		goto L69
	}
L69:
	;
	goto L66
L70:
	;
	v181 = v146
	goto L57
L71:
	;
	v189 = int32(0)
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v190 != 0 {
		v268 = v189
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v84)+80))
	v295 = F_SysCacheGetAttr(m, int32(57), v80, int32(32), v29+int32(507))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L14
	} else {
		goto L89
	}
L73:
	;
	v192 = v181 & int64(39)
	if v192 == int64(0) {
		v268 = v189
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v195 = int32(-6)
	v197 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84)+120)))
	if v197 < v195 {
		v268 = int32(1)
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v205 = v195
	v206 = int32(65530)
	goto L76
L76:
	;
	if v206&int32(65535) == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v268 = v258
	goto L72
L78:
	;
	v258 = int32(1)
	v261 = base.I32_extend16_s(v206 + v258)
	v262 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84)+120)))
	if v261 <= v262 {
		v205 = v261
		v206 = v261
		goto L76
	} else {
		goto L88
	}
L79:
	;
	if base.I32_extend16_s(v206) < int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+119)))
	if v236 == int32(118) {
		goto L78
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v240 = F_SearchSysCache2(m, int32(7), v79, v205)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L14
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	if v240 == int32(0) {
		goto L27
	} else {
		goto L85
	}
L85:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240)+16))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+22)))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244+v245)+91)))
	F_ReleaseCatCache(m, v240)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L14
	} else {
		goto L86
	}
L86:
	;
	if v247 != 0 {
		goto L78
	} else {
		goto L87
	}
L87:
	;
	v252 = v187 + int32(56) + v205<<(uint(int32(3))%32)
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v252)))
	*(*int64)(unsafe.Add(mBase, uint32(v252))) = v253 | v192
	goto L78
L88:
	;
	goto L77
L89:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+507)))
	if v297 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v319 = F_aclcopy(m, v318)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L14
	} else {
		goto L100
	}
L91:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+119)))
	if v303 == int32(83) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	v311 = F_pg_detoast_datum_copy(m, v295)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L14
	} else {
		goto L98
	}
L94:
	;
	v306 = int32(37)
	goto L96
L95:
	;
	v306 = int32(41)
	goto L96
L96:
	;
	v307 = F_acldefault(m, v306, v290)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L14
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+428)) = int32(0)
	v317 = int32(0)
	v318 = v307
	goto L90
L98:
	;
	v315 = F_aclmembers(m, v311, v29+int32(428))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L14
	} else {
		goto L99
	}
L99:
	;
	v317 = v315
	v318 = v311
	goto L90
L100:
	;
	if v181 != int64(0) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v434 == int32(0) {
		v731 = v268
		goto L23
	} else {
		goto L125
	}
L102:
	;
	v328 = F__emscripten_memset_bulkmem(m, v29+int32(288), base.I32_extend8_s(int32(0)), int32(136))
	mBase = m.M
	goto L105
L103:
	;
	goto L104
L104:
	;
	F_UnlockTuple(m, v36, v80+int32(4), int32(7))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L14
	} else {
		goto L124
	}
L105:
	;
	v329 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+464)) = uint16(v329)
	v331 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+456)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v29)+448)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v29)+440)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v29)+432)) = v331
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+272)) = uint16(v329)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+264)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v29)+256)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v29)+248)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v29)+240)) = v331
	v350 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	F_select_best_grantor(m, v350, v181, v318, v290, v29+int32(508), v29+int32(496))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L14
	} else {
		goto L106
	}
L106:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v358 = *(*int64)(unsafe.Add(mBase, uint32(v29)+496))
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v29)+508))
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+119)))
	if v363 == int32(83) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v366 = int32(37)
	goto L109
L108:
	;
	v366 = int32(41)
	goto L109
L109:
	;
	v369 = int32(0)
	v371 = F_restrict_and_check_grant(m, v357, v358, v359, v181, v79, v360, v366, v84+int32(4), v369, v369)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L14
	} else {
		goto L110
	}
L110:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v29)+508))
	v378 = F_merge_acl_with_grant(m, v318, v373, v374, v375, v376, v371, v377, v290)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L14
	} else {
		goto L111
	}
L111:
	;
	v382 = F_aclmembers(m, v378, v29+int32(488))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L14
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+412)) = v378
	v385 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+271)) = uint8(v385)
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
	v394 = F_heap_modify_tuple(m, v80, v387, v29+int32(288), v29+int32(432), v29+int32(240))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L14
	} else {
		goto L113
	}
L113:
	;
	F_CatalogTupleUpdate(m, v36, v394+int32(4), v394)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L14
	} else {
		goto L114
	}
L114:
	;
	F_UnlockTuple(m, v36, v80+int32(4), int32(7))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L14
	} else {
		goto L115
	}
L115:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, _consts[417])))
	if v406 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v29)+428))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v29)+488))
	F_updateAclDependencies(m, int32(1259), v79, int32(0), v290, v317, v419, v382, v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L14
	} else {
		goto L122
	}
L117:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, _consts[418])))
	if v410 != int32(1) {
		goto L116
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	F_recordExtensionInitPrivWorker(m, v79, int32(1259), int32(0), v378)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L14
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	goto L116
L122:
	;
	F_pfree(m, v378)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L14
	} else {
		goto L123
	}
L123:
	;
	goto L101
L124:
	;
	goto L101
L125:
	;
	v437 = int32(0)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v434)+4))
	if v438 <= v437 {
		v731 = v268
		goto L23
	} else {
		goto L126
	}
L126:
	;
	v448 = v437
	goto L127
L127:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v434)+12))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v469+v448<<(uint(int32(2))%32))))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v473)+4))
	if v474 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	v731 = v598
	goto L23
L129:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+119)))
	if v485 != int32(83) {
		v512 = v484
		goto L135
	} else {
		goto L136
	}
L130:
	;
	v484 = int64(39)
	goto L129
L131:
	;
	goto L132
L132:
	;
	v478 = F_string_to_privilege(m, v474)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L14
	} else {
		goto L133
	}
L133:
	;
	if v478&int64(32728) != int64(0) {
		goto L26
	} else {
		goto L134
	}
L134:
	;
	v484 = v478
	goto L129
L135:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v473)+8))
	if v513 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L136:
	;
	if v484&int64(37) == int64(0) {
		v512 = v484
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v494 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L14
	} else {
		goto L138
	}
L138:
	;
	if v494 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L14
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v512 = v484 & int64(2)
	goto L135
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = v84 + int32(4)
	F_errmsg(m, int32(168299), v29-int32(-64))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L14
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(492197), int32(2071), int32(261753))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L14
	} else {
		goto L144
	}
L144:
	;
	goto L141
L145:
	;
	v598 = int32(1)
	v600 = v448 + v598
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v434)+4))
	if v600 < v601 {
		v448 = v600
		goto L127
	} else {
		goto L155
	}
L146:
	;
	v516 = int32(0)
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v513)+4))
	if v517 <= v516 {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v524 = v516
	goto L148
L148:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v513)+12))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v546+v524<<(uint(int32(2))%32))))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v550)+4))
	v552 = F_get_attnum(m, v79, v551)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L14
	} else {
		goto L150
	}
L149:
	;
	goto L145
L150:
	;
	if v552 == int32(0) {
		goto L25
	} else {
		goto L151
	}
L151:
	;
	v558 = base.I32_extend16_s(v552 + int32(7))
	if v558 <= int32(0) {
		goto L24
	} else {
		goto L152
	}
L152:
	;
	if v184 <= v558 {
		goto L24
	} else {
		goto L153
	}
L153:
	;
	v564 = v187 + v558<<(uint(int32(3))%32)
	v565 = *(*int64)(unsafe.Add(mBase, uint32(v564)))
	*(*int64)(unsafe.Add(mBase, uint32(v564))) = v565 | v512
	v569 = v524 + int32(1)
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v513)+4))
	if v569 < v570 {
		v524 = v569
		goto L148
	} else {
		goto L154
	}
L154:
	;
	goto L149
L155:
	;
	goto L128
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v79
	F_errmsg_internal(m, int32(46015), v29+int32(16))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L14
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(492197), int32(1814), int32(261753))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L14
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L14
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+144)) = v84 + int32(4)
	F_errmsg(m, int32(411718), v29+int32(144))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L14
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(492197), int32(1838), int32(261753))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L14
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L14
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+128)) = int32(535536)
	F_errmsg(m, int32(388632), v29+int32(128))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L14
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(492197), int32(1893), int32(261753))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L14
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+100)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v29)+96)) = v205
	F_errmsg_internal(m, int32(46203), v29+int32(96))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L14
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(492197), int32(1621), int32(168204))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L14
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L14
	} else {
		goto L171
	}
L171:
	;
	v680 = F_privilege_to_string(m, v478)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L14
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+80)) = v680
	F_errmsg(m, int32(270855), v29+int32(80))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L14
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(492197), int32(2058), int32(261753))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L14
	} else {
		goto L174
	}
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L175:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L14
	} else {
		goto L176
	}
L176:
	;
	v700 = F_get_rel_name(m, v79)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L14
	} else {
		goto L177
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v700
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v551
	F_errmsg(m, int32(71058), v29+int32(48))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L14
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(492197), int32(1578), int32(168230))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L14
	} else {
		goto L179
	}
L179:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L180:
	;
	F_errmsg_internal(m, int32(398075), int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L14
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(492197), int32(1581), int32(168230))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L14
	} else {
		goto L182
	}
L182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L183:
	;
	F_pfree(m, v319)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L14
	} else {
		goto L250
	}
L184:
	;
	if v731 == int32(0) {
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v759 = int32(0)
	v764 = v759
	v765 = v759
	goto L186
L186:
	;
	v790 = *(*int64)(unsafe.Add(mBase, uint32(v187+v765<<(uint(int32(3))%32))))
	if v790 != int64(0) {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	goto L183
L188:
	;
	v798 = F__emscripten_memset_bulkmem(m, v29+int32(288), base.I32_extend8_s(int32(0)), int32(100))
	mBase = m.M
	goto L191
L189:
	;
	goto L190
L190:
	;
	v1024 = v764 + int32(1)
	v1025 = base.I32_extend16_s(v1024)
	if v1025 < v184 {
		v764 = v1024
		v765 = v1025
		goto L186
	} else {
		goto L249
	}
L191:
	;
	v799 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+456)) = uint8(v799)
	v801 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+448)) = v801
	*(*int64)(unsafe.Add(mBase, uint32(v29)+440)) = v801
	*(*int64)(unsafe.Add(mBase, uint32(v29)+432)) = v801
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+264)) = uint8(v799)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+256)) = v801
	*(*int64)(unsafe.Add(mBase, uint32(v29)+248)) = v801
	*(*int64)(unsafe.Add(mBase, uint32(v29)+240)) = v801
	v815 = int32(7)
	v818 = base.I32_extend16_s(v764 - v815)
	v819 = F_SearchSysCache2(m, v815, v79, v818)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L14
	} else {
		goto L192
	}
L192:
	;
	if v819 == int32(0) {
		goto L17
	} else {
		goto L193
	}
L193:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v819)+16))
	v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v823)+22)))
	v830 = F_SysCacheGetAttr(m, int32(7), v819, int32(22), v29+int32(492))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L14
	} else {
		goto L194
	}
L194:
	;
	v832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+492)))
	if v832 == int32(1) {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v849 = m.G0
	v851 = v849 - int32(16)
	m.G0 = v851
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v847)+16))
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v319)+16))
	v855 = v853 + v854
	if int32(0) <= v855 {
		goto L203
	} else {
		goto L204
	}
L196:
	;
	v837 = F_acldefault(m, int32(6), v290)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L14
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	v841 = F_pg_detoast_datum_copy(m, v830)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L14
	} else {
		goto L200
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = int32(0)
	v847 = v837
	v848 = int32(0)
	goto L195
L200:
	;
	v845 = F_aclmembers(m, v841, v29+int32(488))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L14
	} else {
		goto L201
	}
L201:
	;
	v847 = v841
	v848 = v845
	goto L195
L202:
	;
	v929 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	F_select_best_grantor(m, v929, v790, v862, v290, v29+int32(508), v29+int32(496))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L14
	} else {
		goto L224
	}
L203:
	;
	v861 = v855<<(uint(int32(4))%32) + int32(24)
	v862 = F_palloc0(m, v861)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L14
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L14
	} else {
		goto L221
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v862)+12)) = int32(1033)
	*(*int64)(unsafe.Add(mBase, uint32(v862)+4)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v862))) = v861 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v862)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v862)+16)) = v855
	v875 = v862 + int32(24)
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v319)+8))
	if v876 != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v884 = v876
	goto L209
L208:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	v884 = (v877<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L209
L209:
	;
	v887 = v319 + int32(16)
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v887)))
	v890 = v888 << (uint(int32(4)) % 32)
	if v890 != 0 {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v887)))
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v847)+8))
	if v897 != 0 {
		goto L214
	} else {
		goto L215
	}
L211:
	;
	v891 = F__emscripten_memcpy_bulkmem(m, v875, v884+v319, v890)
	mBase = m.M
	v892 = v891
	goto L213
L212:
	;
	v892 = v875
	goto L213
L213:
	;
	goto L210
L214:
	;
	v905 = v897
	goto L216
L215:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v847)+4))
	v905 = (v898<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L216
L216:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v847)+16))
	v909 = v907 << (uint(int32(4)) % 32)
	if v909 != 0 {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	m.G0 = v851 + int32(16)
	goto L202
L218:
	;
	v910 = F__emscripten_memcpy_bulkmem(m, v892+v893<<(uint(int32(4))%32), v905+v847, v909)
	mBase = m.M
	goto L220
L219:
	;
	goto L220
L220:
	;
	goto L217
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v851))) = v855
	F_errmsg_internal(m, int32(478278), v851)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L14
	} else {
		goto L222
	}
L222:
	;
	F_errfinish(m, int32(492073), int32(433), int32(305309))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L14
	} else {
		goto L223
	}
L223:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L224:
	;
	F_pfree(m, v862)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L14
	} else {
		goto L225
	}
L225:
	;
	v938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v939 = *(*int64)(unsafe.Add(mBase, uint32(v29)+496))
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v29)+508))
	v946 = F_restrict_and_check_grant(m, v938, v939, base.B2i32(v790 == int64(39)), v790, v79, v942, int32(6), v84+int32(4), v818, v823+v824+int32(4))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L14
	} else {
		goto L226
	}
L226:
	;
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v951 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v29)+508))
	v953 = F_merge_acl_with_grant(m, v847, v948, v949, v950, v951, v946, v952, v290)
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L14
	} else {
		goto L227
	}
L227:
	;
	v957 = F_aclmembers(m, v953, v29+int32(484))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L14
	} else {
		goto L228
	}
L228:
	;
	v960 = v953 + int32(16)
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v960)))
	if int32(0) < v961 {
		goto L231
	} else {
		goto L232
	}
L229:
	;
	F_pfree(m, v953)
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L14
	} else {
		goto L247
	}
L230:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v40)+52))
	v979 = F_heap_modify_tuple(m, v819, v972, v29+int32(288), v29+int32(432), v29+int32(240))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L14
	} else {
		goto L235
	}
L231:
	;
	v964 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+261)) = uint8(v964)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+372)) = v953
	goto L230
L232:
	;
	goto L233
L233:
	;
	v967 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+453)) = uint8(v967)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+261)) = uint8(v967)
	v971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+492)))
	if v971 != 0 {
		goto L229
	} else {
		goto L234
	}
L234:
	;
	goto L230
L235:
	;
	F_CatalogTupleUpdate(m, v40, v979+int32(4), v979)
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L14
	} else {
		goto L236
	}
L236:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v960)))
	v987 = int32(*(*uint8)(unsafe.Add(mBase, _consts[417])))
	if v987 == int32(0) {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v29)+488))
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v29)+484))
	F_updateAclDependencies(m, int32(1259), v79, v818, v290, v848, v1002, v957, v1003)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L14
	} else {
		goto L246
	}
L238:
	;
	v991 = int32(*(*uint8)(unsafe.Add(mBase, _consts[418])))
	if v991 != int32(1) {
		goto L237
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	v995 = int32(0)
	if v995 < v985 {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	goto L240
L242:
	;
	v998 = v953
	goto L244
L243:
	;
	v998 = v995
	goto L244
L244:
	;
	F_recordExtensionInitPrivWorker(m, v79, int32(1259), v818, v998)
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L14
	} else {
		goto L245
	}
L245:
	;
	goto L237
L246:
	;
	goto L229
L247:
	;
	F_ReleaseCatCache(m, v819)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L14
	} else {
		goto L248
	}
L248:
	;
	goto L190
L249:
	;
	goto L187
L250:
	;
	F_pfree(m, v187)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L14
	} else {
		goto L251
	}
L251:
	;
	F_ReleaseCatCache(m, v80)
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L14
	} else {
		goto L252
	}
L252:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L14
	} else {
		goto L253
	}
L253:
	;
	v1062 = v70 + int32(1)
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v1062 < v1063 {
		v70 = v1062
		goto L21
	} else {
		goto L254
	}
L254:
	;
	goto L22
L255:
	;
	F_sequence_close(m, v36, int32(3))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L14
	} else {
		goto L256
	}
L256:
	;
	goto L1
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v818
	F_errmsg_internal(m, int32(46203), v29+int32(32))
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L14
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(492197), int32(1668), int32(345307))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L14
	} else {
		goto L259
	}
L259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L260:
	;
	goto L1
L261:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v1122
	F_errmsg_internal(m, int32(478794), v29)
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L14
	} else {
		goto L262
	}
L262:
	;
	F_errfinish(m, int32(492197), int32(645), int32(171652))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L14
	} else {
		goto L263
	}
L263:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L264:
	;
	v1142 = F_table_open(m, int32(6243), int32(3))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L14
	} else {
		goto L267
	}
L265:
	;
	v1135 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v1135 != int64(0) {
		goto L264
	} else {
		goto L266
	}
L266:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(12288)
	goto L264
L267:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1144 == int32(0) {
		goto L269
	} else {
		goto L270
	}
L268:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L14
	} else {
		goto L328
	}
L269:
	;
	F_sequence_close(m, v1142, int32(3))
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L14
	} else {
		goto L327
	}
L270:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1144)+4))
	if v1147 <= int32(0) {
		goto L269
	} else {
		goto L271
	}
L271:
	;
	v1156 = v2
	goto L272
L272:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1144)+12))
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1177+v1156<<(uint(int32(2))%32))))
	v1182 = F_SearchSysCache1(m, int32(44), v1181)
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L14
	} else {
		goto L274
	}
L273:
	;
	goto L269
L274:
	;
	if v1182 == int32(0) {
		goto L268
	} else {
		goto L275
	}
L275:
	;
	v1188 = F_SysCacheGetAttrNotNull(m, int32(44), v1182, int32(2))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L14
	} else {
		goto L276
	}
L276:
	;
	v1190 = F_text_to_cstring(m, v1188)
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L14
	} else {
		goto L277
	}
L277:
	;
	v1196 = F_SysCacheGetAttr(m, int32(44), v1182, int32(3), v29+int32(428))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L14
	} else {
		goto L278
	}
L278:
	;
	v1198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+428)))
	if v1198 == int32(1) {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	v1217 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v1218 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	F_select_best_grantor(m, v1217, v1218, v1214, int32(10), v29+int32(240), v29+int32(432))
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L14
	} else {
		goto L286
	}
L280:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1204 = F_acldefault(m, v1202, int32(10))
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L14
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	v1208 = F_pg_detoast_datum_copy(m, v1196)
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L14
	} else {
		goto L284
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+496)) = int32(0)
	v1214 = v1204
	v1215 = int32(0)
	goto L279
L284:
	;
	v1212 = F_aclmembers(m, v1208, v29+int32(496))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L14
	} else {
		goto L285
	}
L285:
	;
	v1214 = v1208
	v1215 = v1212
	goto L279
L286:
	;
	v1226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v1227 = *(*int64)(unsafe.Add(mBase, uint32(v29)+432))
	v1228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v1229 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v29)+240))
	v1232 = int32(0)
	v1234 = F_restrict_and_check_grant(m, v1226, v1227, v1228, v1229, v1181, v1230, int32(27), v1190, v1232, v1232)
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L14
	} else {
		goto L287
	}
L287:
	;
	v1236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v1237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v29)+240))
	v1242 = F_merge_acl_with_grant(m, v1214, v1236, v1237, v1238, v1239, v1234, v1240, int32(10))
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L14
	} else {
		goto L288
	}
L288:
	;
	v1246 = F_aclmembers(m, v1242, v29+int32(508))
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L14
	} else {
		goto L289
	}
L289:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1250 = F_acldefault(m, v1248, int32(10))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L14
	} else {
		goto L291
	}
L290:
	;
	v1327 = int32(*(*uint8)(unsafe.Add(mBase, _consts[417])))
	if v1327 == int32(0) {
		goto L317
	} else {
		goto L318
	}
L291:
	;
	v1252 = int32(0)
	if v1242 != 0 {
		goto L294
	} else {
		goto L295
	}
L292:
	;
	if v1296 != 0 {
		goto L310
	} else {
		goto L311
	}
L293:
	;
	if v1250 == int32(0) {
		v1292 = v1252
		goto L301
	} else {
		goto L302
	}
L294:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1242)+16))
	if v1254 != 0 {
		goto L293
	} else {
		goto L297
	}
L295:
	;
	goto L296
L296:
	;
	if v1250 == int32(0) {
		goto L298
	} else {
		goto L299
	}
L297:
	;
	goto L296
L298:
	;
	v1296 = int32(1)
	goto L292
L299:
	;
	goto L300
L300:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+16))
	v1296 = base.B2i32(v1259 == int32(0))
	goto L292
L301:
	;
	v1296 = v1292
	goto L292
L302:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+16))
	if v1254 != v1264 {
		v1292 = v1252
		goto L301
	} else {
		goto L303
	}
L303:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1242)+8))
	if v1266 != 0 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1274 = v1266
	goto L306
L305:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1242)+4))
	v1274 = (v1267<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L306
L306:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+8))
	if v1276 != 0 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1284 = v1276
	goto L309
L308:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+4))
	v1284 = (v1277<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L309
L309:
	;
	v1288 = F_memcmp(m, v1274+v1242, v1284+v1250, v1254<<(uint(int32(4))%32))
	mBase = m.M
	v1292 = base.B2i32(v1288 == int32(0))
	goto L301
L310:
	;
	F_CatalogTupleDelete(m, v1142, v1182+int32(4))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L14
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+288)) = int64(0)
	v1303 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+490)) = uint8(v1303)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+488)) = uint16(v1303)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+296)) = v1242
	v1308 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+486)) = uint8(v1308)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+484)) = uint16(v1303)
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+52))
	v1319 = F_heap_modify_tuple(m, v1182, v1312, v29+int32(288), v29+int32(488), v29+int32(484))
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L14
	} else {
		goto L314
	}
L313:
	;
	goto L290
L314:
	;
	F_CatalogTupleUpdate(m, v1142, v1319+int32(4), v1319)
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L14
	} else {
		goto L315
	}
L315:
	;
	goto L290
L316:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v29)+496))
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v29)+508))
	F_updateAclDependencies(m, int32(6243), v1181, int32(0), int32(10), v1215, v1341, v1246, v1342)
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L14
	} else {
		goto L322
	}
L317:
	;
	v1331 = int32(*(*uint8)(unsafe.Add(mBase, _consts[418])))
	if v1331 != int32(1) {
		goto L316
	} else {
		goto L320
	}
L318:
	;
	goto L319
L319:
	;
	F_recordExtensionInitPrivWorker(m, v1181, int32(6243), int32(0), v1242)
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L14
	} else {
		goto L321
	}
L320:
	;
	goto L319
L321:
	;
	goto L316
L322:
	;
	F_ReleaseCatCache(m, v1182)
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L14
	} else {
		goto L323
	}
L323:
	;
	F_pfree(m, v1242)
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L14
	} else {
		goto L324
	}
L324:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L14
	} else {
		goto L325
	}
L325:
	;
	v1352 = v1156 + int32(1)
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v1144)+4))
	if v1352 < v1353 {
		v1156 = v1352
		goto L272
	} else {
		goto L326
	}
L326:
	;
	goto L273
L327:
	;
	goto L1
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+224)) = v1181
	F_errmsg_internal(m, int32(55373), v29+int32(224))
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L14
	} else {
		goto L329
	}
L329:
	;
	F_errfinish(m, int32(492197), int32(2456), int32(215042))
	mBase = m.M
	v1398 = m.ExcPending
	if v1398 != 0 {
		goto L14
	} else {
		goto L330
	}
L330:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L331:
	;
	goto L1
L332:
	;
	goto L1
L333:
	;
	v1419 = F_table_open(m, int32(2995), int32(3))
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L14
	} else {
		goto L336
	}
L334:
	;
	v1412 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v1412 != int64(0) {
		goto L333
	} else {
		goto L335
	}
L335:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(6)
	goto L333
L336:
	;
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1421 == int32(0) {
		goto L338
	} else {
		goto L339
	}
L337:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1636 = m.ExcPending
	if v1636 != 0 {
		goto L14
	} else {
		goto L374
	}
L338:
	;
	F_sequence_close(m, v1419, int32(3))
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L14
	} else {
		goto L373
	}
L339:
	;
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v1421)+4))
	if v1424 <= int32(0) {
		goto L338
	} else {
		goto L340
	}
L340:
	;
	v1443 = v2
	goto L341
L341:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v1421)+12))
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(v1455+v1443<<(uint(int32(2))%32))))
	v1460 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29+int32(248)))) = v1460
	*(*int64)(unsafe.Add(mBase, uint32(v29)+240)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+430)) = uint8(v1460)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+428)) = uint16(v1460)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+494)) = uint8(v1460)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+492)) = uint16(v1460)
	F_ScanKeyInit(m, v29+int32(432), int32(1), int32(3), int32(184), v1459)
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
		goto L14
	} else {
		goto L343
	}
L342:
	;
	goto L338
L343:
	;
	v1480 = int32(1)
	v1485 = F_systable_beginscan(m, v1419, int32(2996), v1480, int32(0), v1480, v29+int32(432))
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L14
	} else {
		goto L344
	}
L344:
	;
	v1487 = F_systable_getnext(m, v1485)
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L14
	} else {
		goto L345
	}
L345:
	;
	if v1487 == int32(0) {
		goto L337
	} else {
		goto L346
	}
L346:
	;
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1487)+16))
	v1492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1491)+22)))
	v1493 = v1491 + v1492
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v1493)+4))
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v1419)+52))
	v1499 = F_heap_getattr_2(m, v1487, int32(3), v1496, v29+int32(507))
	mBase = m.M
	v1500 = m.ExcPending
	if v1500 != 0 {
		goto L14
	} else {
		goto L347
	}
L347:
	;
	v1501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+507)))
	if v1501 == int32(1) {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v1520 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	F_select_best_grantor(m, v1519, v1520, v1516, v1494, v29+int32(508), v29+int32(496))
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L14
	} else {
		goto L355
	}
L349:
	;
	v1506 = F_acldefault(m, int32(22), v1494)
	mBase = m.M
	v1507 = m.ExcPending
	if v1507 != 0 {
		goto L14
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	v1510 = F_pg_detoast_datum_copy(m, v1499)
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L14
	} else {
		goto L353
	}
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = int32(0)
	v1516 = v1506
	v1517 = int32(0)
	goto L348
L353:
	;
	v1514 = F_aclmembers(m, v1510, v29+int32(488))
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L14
	} else {
		goto L354
	}
L354:
	;
	v1516 = v1510
	v1517 = v1514
	goto L348
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+208)) = v1459
	v1534 = F_pg_snprintf(m, v29+int32(288), int32(64), int32(41836), v29+int32(208))
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L14
	} else {
		goto L356
	}
L356:
	;
	v1536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v1537 = *(*int64)(unsafe.Add(mBase, uint32(v29)+496))
	v1538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v1539 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v29)+508))
	v1544 = int32(0)
	v1546 = F_restrict_and_check_grant(m, v1536, v1537, v1538, v1539, v1459, v1540, int32(22), v29+int32(288), v1544, v1544)
	mBase = m.M
	v1547 = m.ExcPending
	if v1547 != 0 {
		goto L14
	} else {
		goto L357
	}
L357:
	;
	v1548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v1549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v29)+508))
	v1553 = F_merge_acl_with_grant(m, v1516, v1548, v1549, v1550, v1551, v1546, v1552, v1494)
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L14
	} else {
		goto L358
	}
L358:
	;
	v1557 = F_aclmembers(m, v1553, v29+int32(484))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L14
	} else {
		goto L359
	}
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v1553
	v1560 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+494)) = uint8(v1560)
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1419)+52))
	v1569 = F_heap_modify_tuple(m, v1487, v1562, v29+int32(240), v29+int32(428), v29+int32(492))
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L14
	} else {
		goto L360
	}
L360:
	;
	F_CatalogTupleUpdate(m, v1419, v1569+int32(4), v1569)
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L14
	} else {
		goto L361
	}
L361:
	;
	v1576 = int32(*(*uint8)(unsafe.Add(mBase, _consts[417])))
	if v1576 == int32(0) {
		goto L363
	} else {
		goto L364
	}
L362:
	;
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(v1493)))
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v29)+488))
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v29)+484))
	F_updateAclDependencies(m, int32(2613), v1588, int32(0), v1494, v1517, v1590, v1557, v1591)
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L14
	} else {
		goto L368
	}
L363:
	;
	v1580 = int32(*(*uint8)(unsafe.Add(mBase, _consts[418])))
	if v1580 != int32(1) {
		goto L362
	} else {
		goto L366
	}
L364:
	;
	goto L365
L365:
	;
	F_recordExtensionInitPrivWorker(m, v1459, int32(2613), int32(0), v1553)
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L14
	} else {
		goto L367
	}
L366:
	;
	goto L365
L367:
	;
	goto L362
L368:
	;
	F_systable_endscan(m, v1485)
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L14
	} else {
		goto L369
	}
L369:
	;
	F_pfree(m, v1553)
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L14
	} else {
		goto L370
	}
L370:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L14
	} else {
		goto L371
	}
L371:
	;
	v1601 = v1443 + int32(1)
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(v1421)+4))
	if v1601 < v1602 {
		v1443 = v1601
		goto L341
	} else {
		goto L372
	}
L372:
	;
	goto L342
L373:
	;
	goto L1
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+192)) = v1459
	F_errmsg_internal(m, int32(41740), v29+int32(192))
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L14
	} else {
		goto L375
	}
L375:
	;
	F_errfinish(m, int32(492197), int32(2316), int32(109259))
	mBase = m.M
	v1647 = m.ExcPending
	if v1647 != 0 {
		goto L14
	} else {
		goto L376
	}
L376:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L377:
	;
	goto L1
L378:
	;
	goto L1
L379:
	;
	goto L1
L380:
	;
	goto L1
L381:
	;
	goto L1
L382:
	;
	if (base.I32_wrap_i64(int64(base.Ui64(int64(4389322341887))>>(uint(base.I64_extend_i32_u(v1699))%64)))|base.B2i32(base.Ui32(int32(42)) < base.Ui32(v1699)))&int32(1) != 0 {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	v1709 = int32(0)
	v1711 = *(*int32)(unsafe.Add(mBase, _consts[419]))
	if v1711 == v1709 {
		goto L386
	} else {
		goto L387
	}
L384:
	;
	goto L385
L385:
	;
	m.G0 = v29 + int32(512)
	return
L386:
	;
	goto L385
L387:
	;
	v1714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1711)+20)))
	if v1714 != 0 {
		goto L386
	} else {
		goto L388
	}
L388:
	;
	v1715 = int32(4476144)
	v1716 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v1711)))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1718
	v1721 = F_palloc(m, int32(40))
	mBase = m.M
	v1722 = m.ExcPending
	if v1722 != 0 {
		goto L14
	} else {
		goto L389
	}
L389:
	;
	v1723 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1721)+32)) = v1723
	v1725 = int32(24)
	v1726 = v1721 + v1725
	v1728 = l0 + v1725
	v1729 = *(*int64)(unsafe.Add(mBase, uint32(v1728)))
	*(*int64)(unsafe.Add(mBase, uint32(v1726))) = v1729
	v1731 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1721)+16)) = v1731
	v1733 = int32(8)
	v1734 = v1721 + v1733
	v1736 = l0 + v1733
	v1737 = *(*int64)(unsafe.Add(mBase, uint32(v1736)))
	*(*int64)(unsafe.Add(mBase, uint32(v1734))) = v1737
	v1739 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v1721))) = v1739
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v1736)))
	v1742 = F_list_copy(m, v1741)
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L14
	} else {
		goto L390
	}
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1734))) = v1742
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1746 = F_list_copy(m, v1745)
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L14
	} else {
		goto L391
	}
L391:
	;
	v1748 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1726))) = v1748
	*(*int32)(unsafe.Add(mBase, uint32(v1721)+28)) = v1746
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1728)))
	if v1751 == v1748 {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	v1825 = F_palloc(m, int32(40))
	mBase = m.M
	v1826 = m.ExcPending
	if v1826 != 0 {
		goto L14
	} else {
		goto L400
	}
L393:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+4))
	if v1754 <= int32(0) {
		goto L392
	} else {
		goto L394
	}
L394:
	;
	v1760 = v1709
	v1770 = int32(0)
	goto L395
L395:
	;
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+12))
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(v1784+v1760<<(uint(int32(2))%32))))
	v1789 = F_copyObjectImpl(m, v1788)
	mBase = m.M
	v1790 = m.ExcPending
	if v1790 != 0 {
		goto L14
	} else {
		goto L397
	}
L396:
	;
	goto L392
L397:
	;
	v1791 = F_lappend(m, v1770, v1789)
	mBase = m.M
	v1792 = m.ExcPending
	if v1792 != 0 {
		goto L14
	} else {
		goto L398
	}
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1721)+24)) = v1791
	v1795 = v1760 + int32(1)
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+4))
	if v1795 < v1796 {
		v1760 = v1795
		v1770 = v1791
		goto L395
	} else {
		goto L399
	}
L399:
	;
	goto L396
L400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1825))) = int32(2)
	v1830 = int32(*(*uint8)(unsafe.Add(mBase, _consts[417])))
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+12)) = v1721
	*(*uint8)(unsafe.Add(mBase, uint32(v1825)+4)) = uint8(v1830)
	*(*int32)(unsafe.Add(mBase, uint32(v1825)+8)) = int32(0)
	v1836 = *(*int32)(unsafe.Add(mBase, _consts[419]))
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+28))
	v1838 = F_lappend(m, v1837, v1825)
	mBase = m.M
	v1839 = m.ExcPending
	if v1839 != 0 {
		goto L14
	} else {
		goto L401
	}
L401:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, _consts[419]))
	*(*int32)(unsafe.Add(mBase, uint32(v1841)+28)) = v1838
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1716
	goto L386
}
func F_ExecIncrementalSort(m *base.Module, l0 int32) int32 {
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
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
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v196 int32
	_ = v196
	var v199 int64
	_ = v199
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v232 int64
	_ = v232
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int64
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int64
	_ = v256
	var v270 int64
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int64
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v317 int64
	_ = v317
	var v318 int64
	_ = v318
	var v320 int32
	_ = v320
	var v321 int64
	_ = v321
	var v323 int64
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int64
	_ = v331
	var v337 int64
	_ = v337
	var v341 int32
	_ = v341
	var v351 int32
	_ = v351
	var v354 int64
	_ = v354
	var v358 int64
	_ = v358
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int64
	_ = v375
	var v376 int64
	_ = v376
	var v379 int64
	_ = v379
	var v382 int64
	_ = v382
	var v383 int64
	_ = v383
	var v386 int64
	_ = v386
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int64
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v408 int64
	_ = v408
	var v409 int64
	_ = v409
	var v411 int32
	_ = v411
	var v412 int64
	_ = v412
	var v414 int64
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int64
	_ = v422
	var v428 int64
	_ = v428
	var v432 int32
	_ = v432
	var v442 int32
	_ = v442
	var v445 int64
	_ = v445
	var v449 int64
	_ = v449
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int64
	_ = v466
	var v467 int64
	_ = v467
	var v470 int64
	_ = v470
	var v473 int64
	_ = v473
	var v474 int64
	_ = v474
	var v477 int64
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v492 int32
	_ = v492
	var v494 int64
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v510 int64
	_ = v510
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int64
	_ = v547
	var v548 int64
	_ = v548
	var v549 int64
	_ = v549
	var v551 int64
	_ = v551
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v581 int64
	_ = v581
	var v586 int32
	_ = v586
	var v587 int64
	_ = v587
	var v588 int64
	_ = v588
	var v591 int64
	_ = v591
	var v594 int64
	_ = v594
	var v595 int64
	_ = v595
	var v598 int64
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v622 int64
	_ = v622
	var v627 int32
	_ = v627
	var v628 int64
	_ = v628
	var v629 int64
	_ = v629
	var v632 int64
	_ = v632
	var v635 int64
	_ = v635
	var v636 int64
	_ = v636
	var v639 int64
	_ = v639
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int64
	_ = v653
	var v654 int64
	_ = v654
	var v655 int64
	_ = v655
	var v657 int64
	_ = v657
	var v658 int64
	_ = v658
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v677 int64
	_ = v677
	var v679 int32
	_ = v679
	var v695 int64
	_ = v695
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v744 int64
	_ = v744
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v757 int64
	_ = v757
	var v758 int64
	_ = v758
	var v760 int32
	_ = v760
	var v761 int64
	_ = v761
	var v763 int64
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v771 int64
	_ = v771
	var v777 int64
	_ = v777
	var v781 int32
	_ = v781
	var v791 int32
	_ = v791
	var v794 int64
	_ = v794
	var v798 int64
	_ = v798
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v815 int64
	_ = v815
	var v816 int64
	_ = v816
	var v819 int64
	_ = v819
	var v822 int64
	_ = v822
	var v823 int64
	_ = v823
	var v826 int64
	_ = v826
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v835 int64
	_ = v835
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v848 int64
	_ = v848
	var v849 int64
	_ = v849
	var v851 int32
	_ = v851
	var v852 int64
	_ = v852
	var v854 int64
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v862 int64
	_ = v862
	var v868 int64
	_ = v868
	var v872 int32
	_ = v872
	var v882 int32
	_ = v882
	var v885 int64
	_ = v885
	var v889 int64
	_ = v889
	var v891 int32
	_ = v891
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v906 int64
	_ = v906
	var v907 int64
	_ = v907
	var v910 int64
	_ = v910
	var v913 int64
	_ = v913
	var v914 int64
	_ = v914
	var v917 int64
	_ = v917
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v930 int32
	_ = v930
	var v933 int64
	_ = v933
	var v934 int64
	_ = v934
	var v935 int64
	_ = v935
	var v937 int64
	_ = v937
	var v942 int32
	_ = v942
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v30&int32(-2) != int32(2) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L4
	} else {
		goto L284
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L4
	} else {
		goto L281
	}
L8:
	;
	m.G0 = v18 + int32(48)
	return v968
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+56))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v60 != 0 {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	if v30 != int32(2) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v38 = v37
	goto L13
L12:
	;
	v38 = v27
	goto L13
L13:
	;
	v41 = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v44 = F_tuplesort_gettupleslot(m, v38, base.B2i32(v29 == int32(1)), v41, v42, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	if v44 != 0 {
		v968 = v42
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)))
	if v46 != 0 {
		v968 = v42
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	if int64(0) < v47 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_switchToPresortedPrefixMode(m, l0)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(0)
	goto L9
L20:
	;
	goto L9
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v29
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v955 != int32(2) {
		goto L277
	} else {
		goto L278
	}
L22:
	;
	v667 = v27
	v677 = int64(0)
	v679 = v60
	goto L24
L23:
	;
	if v27 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	if v679 != int32(1) {
		v942 = v667
		goto L21
	} else {
		goto L197
	}
L25:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v196 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L26:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+96))
	v67 = F_palloc(m, v64*int32(36))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_tuplesort_reset(m, v27)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L46
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v67
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v63)+96))
	if int32(0) < v70 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v77 = int32(0)
	goto L33
L31:
	;
	goto L32
L32:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	v170 = *(*int32)(unsafe.Add(mBase, _consts[523]))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v173 != 0 {
		goto L42
	} else {
		goto L43
	}
L33:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v92 = v89 + v77*int32(36)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v63)+76))
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93+v77<<(uint(int32(1))%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v92)+32)) = uint16(v97)
	v100 = v77 << (uint(int32(2)) % 32)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v63)+80))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v100+v101)))
	v105 = F_get_equality_op_for_ordering_op(m, v103, int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L35
	}
L34:
	;
	goto L32
L35:
	;
	if v105 == int32(0) {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	v109 = F_get_opcode(m, v105)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	if v109 == int32(0) {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	F_fmgr_info_cxt(m, v109, v92, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v118 = F_palloc0(m, int32(36))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92)+28)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v92
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v92)+28))
	v123 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v122)+4)) = v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v92)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+8)) = v123
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v92)+28))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v63)+84))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129+v100)))
	*(*int32)(unsafe.Add(mBase, uint32(v128)+12)) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v92)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+16)) = uint8(v123)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v92)+28))
	v137 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v136)+18)) = uint16(v137)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v92)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v139)+24)) = uint8(v123)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v92)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v142)+32)) = uint8(v123)
	v146 = v77 + int32(1)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v63)+96))
	if v146 < v147 {
		v77 = v146
		goto L33
	} else {
		goto L41
	}
L41:
	;
	goto L34
L42:
	;
	v174 = int32(2)
	goto L44
L43:
	;
	v174 = int32(0)
	goto L44
L44:
	;
	v175 = F_tuplesort_begin_heap(m, v59, v164, v165, v166, v167, v168, v170, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v175
	v183 = v175
	goto L25
L46:
	;
	v183 = v27
	goto L25
L47:
	;
	v199 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v200 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	v201 = v199 - v200
	if v201 <= int64(31) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v237 = int64(32)
	goto L49
L49:
	;
	v238 = int64(0)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v239 == int32(0) {
		v256 = v238
		goto L68
	} else {
		goto L69
	}
L50:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v183)+236))
	if v206 != 0 {
		goto L56
	} else {
		goto L57
	}
L51:
	;
	goto L52
L52:
	;
	v232 = int64(32)
	if v232 <= v201 {
		goto L65
	} else {
		goto L66
	}
L53:
	;
	goto L52
L54:
	;
	goto L53
L55:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v183)+72)) = uint32(v201)
	v215 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v183)+68)) = uint8(v215)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v183)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+24)) = int32(0)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v183)+44))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+32))
	if v221 != 0 {
		goto L62
	} else {
		goto L63
	}
L56:
	;
	if int64(1073741823) < v201 {
		goto L54
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if int64(1073741823) < v201 {
		goto L54
	} else {
		goto L61
	}
L59:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v183)+232))
	if v209 != int32(-1) {
		goto L55
	} else {
		goto L60
	}
L60:
	;
	goto L54
L61:
	;
	goto L55
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+16)) = v221
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v183)+44))
	v224 = v223
	goto L64
L63:
	;
	v224 = v220
	goto L64
L64:
	;
	v225 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v224)+28)) = v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v183)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v227)+32)) = v225
	goto L54
L65:
	;
	v235 = v232
	goto L67
L66:
	;
	v235 = v201
	goto L67
L67:
	;
	v237 = v235
	goto L49
L68:
	;
	v270 = v256
	goto L78
L69:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+4)))
	if v242&int32(2) != 0 {
		v256 = v238
		goto L68
	} else {
		goto L70
	}
L70:
	;
	F_tuplesort_puttupleslot(m, v183, v239)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	v247 = int64(1)
	if v237 == v247 {
		v256 = v247
		goto L68
	} else {
		goto L72
	}
L72:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+8))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+12))
	m.T0[v252].(func(*base.Module, int32))(m, v250)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	v256 = v247
	goto L68
L74:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651)+69)))
	if v652 != 0 {
		goto L190
	} else {
		goto L191
	}
L75:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v618 = m.G0
	v620 = v618 - int32(16)
	m.G0 = v620
	v622 = *(*int64)(unsafe.Add(mBase, uint32(v614)))
	*(*int64)(unsafe.Add(mBase, uint32(v614))) = v622 + int64(1)
	F_tuplesort_get_stats(m, v615, v620)
	mBase = m.M
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v620)+4))
	switch v627 {
	case 0:
		goto L187
	case 1:
		goto L186
	default:
		goto L185
	}
L76:
	;
	v614 = l0 + int32(176)
	goto L75
L77:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v539)+8))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v540)+32))
	m.T0[v541].(func(*base.Module, int32, int32))(m, v539, v276)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L4
	} else {
		goto L163
	}
L78:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v58)+52))
	if v272 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v516)+8))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v517)+12))
	m.T0[v518].(func(*base.Module, int32))(m, v516)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L4
	} else {
		goto L158
	}
L80:
	;
	F_ExecReScan(m, v58)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L4
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v276 = m.T0[v275].(func(*base.Module, int32) int32)(m, v58)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L4
	} else {
		goto L85
	}
L83:
	;
	goto L82
L84:
	;
	if v270 < v237 {
		goto L147
	} else {
		goto L148
	}
L85:
	;
	if v276 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+4)))
	if v278&int32(2) == int32(0) {
		goto L84
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v283 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)) = uint8(v283)
	F_tuplesort_performsort(m, v183)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L4
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v287 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(2)
	v942 = v183
	goto L21
L92:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v290 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v395 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v395 + int64(1)
	v400 = v18 + int32(32)
	v401 = int32(0)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v394)+128))
	if v405 == v401 {
		goto L124
	} else {
		goto L125
	}
L94:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+280)))
	if v293 != int32(1) {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v298 = *(*int32)(unsafe.Add(mBase, _consts[422]))
	v303 = v290 + v298*int32(96) + int32(8)
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v303)))
	*(*int64)(unsafe.Add(mBase, uint32(v303))) = v304 + int64(1)
	v309 = v18 + int32(32)
	v310 = int32(0)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v296)+128))
	if v314 == v310 {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	switch v374 {
	case 0:
		goto L118
	case 1:
		goto L117
	default:
		goto L116
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v309)+4)) = v351
	v354 = *(*int64)(unsafe.Add(mBase, uint32(v296)+112))
	v358 = base.I64_div_s(v354+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+8)) = v358
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v296)+124))
	switch v360 - int32(3) {
	case 0:
		goto L112
	case 1:
		v371 = v360
		goto L109
	case 2:
		goto L111
	default:
		goto L110
	}
L98:
	;
	if v330&int32(255) != base.B2i32(v314 != int32(0)) {
		goto L104
	} else {
		goto L105
	}
L99:
	;
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v296)+96))
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v296)+88))
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+120)))
	v330 = v320
	v331 = v317 - v318
	goto L98
L100:
	;
	goto L101
L101:
	;
	v321 = F_LogicalTapeSetBlocks(m, v314)
	mBase = m.M
	v323 = v321 << (uint(int64(13)) % 64)
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+120)))
	if v324 != 0 {
		v330 = v324
		v331 = v323
		goto L98
	} else {
		goto L102
	}
L102:
	;
	v325 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v296)+120)) = uint8(v325)
	*(*int64)(unsafe.Add(mBase, uint32(v296)+112)) = v323
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v296)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v296)+124)) = v328
	v351 = v310
	goto L97
L103:
	;
	v351 = int32(1)
	goto L97
L104:
	;
	if v330&int32(1) != 0 {
		v351 = v310
		goto L97
	} else {
		goto L108
	}
L105:
	;
	v337 = *(*int64)(unsafe.Add(mBase, uint32(v296)+112))
	if v331 <= v337 {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v296)+120)) = uint8(v330)
	*(*int64)(unsafe.Add(mBase, uint32(v296)+112)) = v331
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v296)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v296)+124)) = v341
	if v330&int32(1) == int32(0) {
		goto L103
	} else {
		goto L107
	}
L107:
	;
	v351 = v310
	goto L97
L108:
	;
	goto L103
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v309))) = v371
	goto L96
L110:
	;
	v371 = int32(0)
	goto L109
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v309))) = int32(8)
	goto L96
L112:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+69)))
	if v365 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v366 = int32(1)
	goto L115
L114:
	;
	v366 = int32(2)
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v309))) = v366
	goto L96
L116:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v303)+40))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v303)+40)) = v390 | v391
	goto L91
L117:
	;
	v382 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	v383 = *(*int64)(unsafe.Add(mBase, uint32(v303)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v303)+32)) = v382 + v383
	v386 = *(*int64)(unsafe.Add(mBase, uint32(v303)+24))
	if v382 <= v386 {
		goto L116
	} else {
		goto L120
	}
L118:
	;
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	v376 = *(*int64)(unsafe.Add(mBase, uint32(v303)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v303)+16)) = v375 + v376
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v303)+8))
	if v375 <= v379 {
		goto L116
	} else {
		goto L119
	}
L119:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v303)+8)) = v375
	goto L116
L120:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v303)+24)) = v382
	goto L116
L121:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	switch v465 {
	case 0:
		goto L143
	case 1:
		goto L142
	default:
		goto L141
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v400)+4)) = v442
	v445 = *(*int64)(unsafe.Add(mBase, uint32(v394)+112))
	v449 = base.I64_div_s(v445+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(v400)+8)) = v449
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v394)+124))
	switch v451 - int32(3) {
	case 0:
		goto L137
	case 1:
		v462 = v451
		goto L134
	case 2:
		goto L136
	default:
		goto L135
	}
L123:
	;
	if v421&int32(255) != base.B2i32(v405 != int32(0)) {
		goto L129
	} else {
		goto L130
	}
L124:
	;
	v408 = *(*int64)(unsafe.Add(mBase, uint32(v394)+96))
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v394)+88))
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+120)))
	v421 = v411
	v422 = v408 - v409
	goto L123
L125:
	;
	goto L126
L126:
	;
	v412 = F_LogicalTapeSetBlocks(m, v405)
	mBase = m.M
	v414 = v412 << (uint(int64(13)) % 64)
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+120)))
	if v415 != 0 {
		v421 = v415
		v422 = v414
		goto L123
	} else {
		goto L127
	}
L127:
	;
	v416 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v394)+120)) = uint8(v416)
	*(*int64)(unsafe.Add(mBase, uint32(v394)+112)) = v414
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v394)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v394)+124)) = v419
	v442 = v401
	goto L122
L128:
	;
	v442 = int32(1)
	goto L122
L129:
	;
	if v421&int32(1) != 0 {
		v442 = v401
		goto L122
	} else {
		goto L133
	}
L130:
	;
	v428 = *(*int64)(unsafe.Add(mBase, uint32(v394)+112))
	if v422 <= v428 {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v394)+120)) = uint8(v421)
	*(*int64)(unsafe.Add(mBase, uint32(v394)+112)) = v422
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v394)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v394)+124)) = v432
	if v421&int32(1) == int32(0) {
		goto L128
	} else {
		goto L132
	}
L132:
	;
	v442 = v401
	goto L122
L133:
	;
	goto L128
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v400))) = v462
	goto L121
L135:
	;
	v462 = int32(0)
	goto L134
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v400))) = int32(8)
	goto L121
L137:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+69)))
	if v456 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v457 = int32(1)
	goto L140
L139:
	;
	v457 = int32(2)
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v400))) = v457
	goto L121
L141:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v481 | v482
	goto L91
L142:
	;
	v473 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	v474 = *(*int64)(unsafe.Add(mBase, uint32(l0)+208))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v473 + v474
	v477 = *(*int64)(unsafe.Add(mBase, uint32(l0)+200))
	if v473 <= v477 {
		goto L141
	} else {
		goto L145
	}
L143:
	;
	v466 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	v467 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v466 + v467
	v470 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	if v466 <= v470 {
		goto L141
	} else {
		goto L144
	}
L144:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v466
	goto L141
L145:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+200)) = v473
	goto L141
L146:
	;
	if v510 < int64(65) {
		v270 = v510
		goto L78
	} else {
		goto L156
	}
L147:
	;
	F_tuplesort_puttupleslot(m, v183, v276)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L4
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v502 = F_isCurrentGroup(m, l0, v501, v276)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L4
	} else {
		goto L153
	}
L150:
	;
	v494 = v270 + int64(1)
	if v494 != v237 {
		v510 = v494
		goto L146
	} else {
		goto L151
	}
L151:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)+8))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+32))
	m.T0[v498].(func(*base.Module, int32, int32))(m, v496, v276)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	v270 = v237
	goto L78
L153:
	;
	if v502 == int32(0) {
		goto L77
	} else {
		goto L154
	}
L154:
	;
	F_tuplesort_puttupleslot(m, v183, v276)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	v510 = v270 + int64(1)
	goto L146
L156:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v513 == int32(2) {
		v270 = v510
		goto L78
	} else {
		goto L157
	}
L157:
	;
	goto L79
L158:
	;
	F_tuplesort_performsort(m, v183)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L4
	} else {
		goto L159
	}
L159:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v523 == int32(0) {
		goto L74
	} else {
		goto L160
	}
L160:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v526 == int32(0) {
		goto L76
	} else {
		goto L161
	}
L161:
	;
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+280)))
	if v529 != int32(1) {
		goto L76
	} else {
		goto L162
	}
L162:
	;
	v533 = *(*int32)(unsafe.Add(mBase, _consts[422]))
	v614 = v526 + v533*int32(96) + int32(8)
	goto L75
L163:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v544 == int32(1) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v547 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v548 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	v549 = v548 + v270
	if v547 < v549 {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	goto L166
L166:
	;
	F_tuplesort_performsort(m, v183)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L4
	} else {
		goto L170
	}
L167:
	;
	v551 = v547
	goto L169
L168:
	;
	v551 = v549
	goto L169
L169:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v551
	goto L166
L170:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v557 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v558 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	goto L173
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(2)
	v942 = v183
	goto L21
L174:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v577 = m.G0
	v579 = v577 - int32(16)
	m.G0 = v579
	v581 = *(*int64)(unsafe.Add(mBase, uint32(v573)))
	*(*int64)(unsafe.Add(mBase, uint32(v573))) = v581 + int64(1)
	F_tuplesort_get_stats(m, v574, v579)
	mBase = m.M
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v579)+4))
	switch v586 {
	case 0:
		goto L181
	case 1:
		goto L180
	default:
		goto L179
	}
L175:
	;
	v573 = l0 + int32(176)
	goto L174
L176:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+280)))
	if v561 != int32(1) {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v565 = *(*int32)(unsafe.Add(mBase, _consts[422]))
	v573 = v558 + v565*int32(96) + int32(8)
	goto L174
L178:
	;
	goto L173
L179:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v573)+40))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v579)))
	*(*int32)(unsafe.Add(mBase, uint32(v573)+40)) = v602 | v603
	m.G0 = v579 + int32(16)
	goto L178
L180:
	;
	v594 = *(*int64)(unsafe.Add(mBase, uint32(v579)+8))
	v595 = *(*int64)(unsafe.Add(mBase, uint32(v573)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v573)+32)) = v594 + v595
	v598 = *(*int64)(unsafe.Add(mBase, uint32(v573)+24))
	if v594 <= v598 {
		goto L179
	} else {
		goto L183
	}
L181:
	;
	v587 = *(*int64)(unsafe.Add(mBase, uint32(v579)+8))
	v588 = *(*int64)(unsafe.Add(mBase, uint32(v573)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v573)+16)) = v587 + v588
	v591 = *(*int64)(unsafe.Add(mBase, uint32(v573)+8))
	if v587 <= v591 {
		goto L179
	} else {
		goto L182
	}
L182:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v573)+8)) = v587
	goto L179
L183:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v573)+24)) = v594
	goto L179
L184:
	;
	goto L74
L185:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v614)+40))
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	*(*int32)(unsafe.Add(mBase, uint32(v614)+40)) = v643 | v644
	m.G0 = v620 + int32(16)
	goto L184
L186:
	;
	v635 = *(*int64)(unsafe.Add(mBase, uint32(v620)+8))
	v636 = *(*int64)(unsafe.Add(mBase, uint32(v614)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v614)+32)) = v635 + v636
	v639 = *(*int64)(unsafe.Add(mBase, uint32(v614)+24))
	if v635 <= v639 {
		goto L185
	} else {
		goto L189
	}
L187:
	;
	v628 = *(*int64)(unsafe.Add(mBase, uint32(v620)+8))
	v629 = *(*int64)(unsafe.Add(mBase, uint32(v614)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v614)+16)) = v628 + v629
	v632 = *(*int64)(unsafe.Add(mBase, uint32(v614)+8))
	if v628 <= v632 {
		goto L185
	} else {
		goto L188
	}
L188:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v614)+8)) = v628
	goto L185
L189:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v614)+24)) = v635
	goto L185
L190:
	;
	v653 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v654 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	v655 = v653 - v654
	if v655 < v510 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	v658 = v510
	goto L192
L192:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v658
	F_switchToPresortedPrefixMode(m, l0)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L4
	} else {
		goto L196
	}
L193:
	;
	v657 = v655
	goto L195
L194:
	;
	v657 = v510
	goto L195
L195:
	;
	v658 = v657
	goto L192
L196:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v667 = v183
	v677 = v658
	v679 = v663
	goto L24
L197:
	;
	v695 = v677
	goto L199
L198:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	F_tuplesort_performsort(m, v724)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L4
	} else {
		goto L217
	}
L199:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v58)+52))
	if v697 != 0 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v718)+8))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v719)+32))
	m.T0[v720].(func(*base.Module, int32, int32))(m, v718, v701)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L4
	} else {
		goto L216
	}
L201:
	;
	F_ExecReScan(m, v58)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L4
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v701 = m.T0[v700].(func(*base.Module, int32) int32)(m, v58)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L4
	} else {
		goto L206
	}
L204:
	;
	goto L203
L205:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v711 = F_isCurrentGroup(m, l0, v710, v701)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L4
	} else {
		goto L211
	}
L206:
	;
	if v701 != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701)+4)))
	if v703&int32(2) == int32(0) {
		goto L205
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	v708 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)) = uint8(v708)
	goto L198
L210:
	;
	goto L209
L211:
	;
	if v711 != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	F_tuplesort_puttupleslot(m, v713, v701)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L4
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	goto L200
L215:
	;
	v695 = v695 + int64(1)
	goto L199
L216:
	;
	goto L198
L217:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v727 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(3)
	v930 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v930 != int32(1) {
		v942 = v667
		goto L21
	} else {
		goto L273
	}
L219:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v730 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v835 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+224)) = v835 + int64(1)
	v840 = v18 + int32(32)
	v841 = int32(0)
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v834)+128))
	if v845 == v841 {
		goto L251
	} else {
		goto L252
	}
L221:
	;
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+280)))
	if v733 != int32(1) {
		goto L220
	} else {
		goto L222
	}
L222:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v738 = *(*int32)(unsafe.Add(mBase, _consts[422]))
	v743 = v730 + v738*int32(96) + int32(56)
	v744 = *(*int64)(unsafe.Add(mBase, uint32(v743)))
	*(*int64)(unsafe.Add(mBase, uint32(v743))) = v744 + int64(1)
	v749 = v18 + int32(32)
	v750 = int32(0)
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v736)+128))
	if v754 == v750 {
		goto L226
	} else {
		goto L227
	}
L223:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	switch v814 {
	case 0:
		goto L245
	case 1:
		goto L244
	default:
		goto L243
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v749)+4)) = v791
	v794 = *(*int64)(unsafe.Add(mBase, uint32(v736)+112))
	v798 = base.I64_div_s(v794+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(v749)+8)) = v798
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v736)+124))
	switch v800 - int32(3) {
	case 0:
		goto L239
	case 1:
		v811 = v800
		goto L236
	case 2:
		goto L238
	default:
		goto L237
	}
L225:
	;
	if v770&int32(255) != base.B2i32(v754 != int32(0)) {
		goto L231
	} else {
		goto L232
	}
L226:
	;
	v757 = *(*int64)(unsafe.Add(mBase, uint32(v736)+96))
	v758 = *(*int64)(unsafe.Add(mBase, uint32(v736)+88))
	v760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v736)+120)))
	v770 = v760
	v771 = v757 - v758
	goto L225
L227:
	;
	goto L228
L228:
	;
	v761 = F_LogicalTapeSetBlocks(m, v754)
	mBase = m.M
	v763 = v761 << (uint(int64(13)) % 64)
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v736)+120)))
	if v764 != 0 {
		v770 = v764
		v771 = v763
		goto L225
	} else {
		goto L229
	}
L229:
	;
	v765 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v736)+120)) = uint8(v765)
	*(*int64)(unsafe.Add(mBase, uint32(v736)+112)) = v763
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v736)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+124)) = v768
	v791 = v750
	goto L224
L230:
	;
	v791 = int32(1)
	goto L224
L231:
	;
	if v770&int32(1) != 0 {
		v791 = v750
		goto L224
	} else {
		goto L235
	}
L232:
	;
	v777 = *(*int64)(unsafe.Add(mBase, uint32(v736)+112))
	if v771 <= v777 {
		goto L231
	} else {
		goto L233
	}
L233:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v736)+120)) = uint8(v770)
	*(*int64)(unsafe.Add(mBase, uint32(v736)+112)) = v771
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v736)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+124)) = v781
	if v770&int32(1) == int32(0) {
		goto L230
	} else {
		goto L234
	}
L234:
	;
	v791 = v750
	goto L224
L235:
	;
	goto L230
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v749))) = v811
	goto L223
L237:
	;
	v811 = int32(0)
	goto L236
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v749))) = int32(8)
	goto L223
L239:
	;
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v736)+69)))
	if v805 != 0 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v806 = int32(1)
	goto L242
L241:
	;
	v806 = int32(2)
	goto L242
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v749))) = v806
	goto L223
L243:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v743)+40))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v743)+40)) = v830 | v831
	goto L218
L244:
	;
	v822 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	v823 = *(*int64)(unsafe.Add(mBase, uint32(v743)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v743)+32)) = v822 + v823
	v826 = *(*int64)(unsafe.Add(mBase, uint32(v743)+24))
	if v822 <= v826 {
		goto L243
	} else {
		goto L247
	}
L245:
	;
	v815 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	v816 = *(*int64)(unsafe.Add(mBase, uint32(v743)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v743)+16)) = v815 + v816
	v819 = *(*int64)(unsafe.Add(mBase, uint32(v743)+8))
	if v815 <= v819 {
		goto L243
	} else {
		goto L246
	}
L246:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v743)+8)) = v815
	goto L243
L247:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v743)+24)) = v822
	goto L243
L248:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	switch v905 {
	case 0:
		goto L270
	case 1:
		goto L269
	default:
		goto L268
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v840)+4)) = v882
	v885 = *(*int64)(unsafe.Add(mBase, uint32(v834)+112))
	v889 = base.I64_div_s(v885+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(v840)+8)) = v889
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v834)+124))
	switch v891 - int32(3) {
	case 0:
		goto L264
	case 1:
		v902 = v891
		goto L261
	case 2:
		goto L263
	default:
		goto L262
	}
L250:
	;
	if v861&int32(255) != base.B2i32(v845 != int32(0)) {
		goto L256
	} else {
		goto L257
	}
L251:
	;
	v848 = *(*int64)(unsafe.Add(mBase, uint32(v834)+96))
	v849 = *(*int64)(unsafe.Add(mBase, uint32(v834)+88))
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v834)+120)))
	v861 = v851
	v862 = v848 - v849
	goto L250
L252:
	;
	goto L253
L253:
	;
	v852 = F_LogicalTapeSetBlocks(m, v845)
	mBase = m.M
	v854 = v852 << (uint(int64(13)) % 64)
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v834)+120)))
	if v855 != 0 {
		v861 = v855
		v862 = v854
		goto L250
	} else {
		goto L254
	}
L254:
	;
	v856 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v834)+120)) = uint8(v856)
	*(*int64)(unsafe.Add(mBase, uint32(v834)+112)) = v854
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v834)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v834)+124)) = v859
	v882 = v841
	goto L249
L255:
	;
	v882 = int32(1)
	goto L249
L256:
	;
	if v861&int32(1) != 0 {
		v882 = v841
		goto L249
	} else {
		goto L260
	}
L257:
	;
	v868 = *(*int64)(unsafe.Add(mBase, uint32(v834)+112))
	if v862 <= v868 {
		goto L256
	} else {
		goto L258
	}
L258:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v834)+120)) = uint8(v861)
	*(*int64)(unsafe.Add(mBase, uint32(v834)+112)) = v862
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v834)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v834)+124)) = v872
	if v861&int32(1) == int32(0) {
		goto L255
	} else {
		goto L259
	}
L259:
	;
	v882 = v841
	goto L249
L260:
	;
	goto L255
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v840))) = v902
	goto L248
L262:
	;
	v902 = int32(0)
	goto L261
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v840))) = int32(8)
	goto L248
L264:
	;
	v896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v834)+69)))
	if v896 != 0 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v897 = int32(1)
	goto L267
L266:
	;
	v897 = int32(2)
	goto L267
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v840))) = v897
	goto L248
L268:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+264)) = v921 | v922
	goto L218
L269:
	;
	v913 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	v914 = *(*int64)(unsafe.Add(mBase, uint32(l0)+256))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+256)) = v913 + v914
	v917 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
	if v913 <= v917 {
		goto L268
	} else {
		goto L272
	}
L270:
	;
	v906 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	v907 = *(*int64)(unsafe.Add(mBase, uint32(l0)+240))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+240)) = v906 + v907
	v910 = *(*int64)(unsafe.Add(mBase, uint32(l0)+232))
	if v906 <= v910 {
		goto L268
	} else {
		goto L271
	}
L271:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+232)) = v906
	goto L268
L272:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v913
	goto L268
L273:
	;
	v933 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v934 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	v935 = v934 + v695
	if v933 < v935 {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v937 = v933
	goto L276
L275:
	;
	v937 = v935
	goto L276
L276:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v937
	v942 = v667
	goto L21
L277:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v959 = v958
	goto L279
L278:
	;
	v959 = v942
	goto L279
L279:
	;
	v962 = int32(0)
	v963 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v965 = F_tuplesort_gettupleslot(m, v959, base.B2i32(v29 == int32(1)), v962, v963, v962)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L4
	} else {
		goto L280
	}
L280:
	;
	v968 = v963
	goto L8
L281:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v63)+80))
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v990+v77<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v994
	F_errmsg_internal(m, int32(42925), v18)
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L4
	} else {
		goto L282
	}
L282:
	;
	F_errfinish(m, int32(487517), int32(186), int32(150464))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L4
	} else {
		goto L283
	}
L283:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v105
	F_errmsg_internal(m, int32(42774), v18+int32(16))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L4
	} else {
		goto L285
	}
L285:
	;
	F_errfinish(m, int32(487517), int32(190), int32(150464))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L4
	} else {
		goto L286
	}
L286:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecInitJunkFilter(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v3 = int32(0)
	v8 = F_ExecCleanTypeFromTL(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if l1 != 0 {
			F_ExecSetSlotDescriptor(m, l1, v8)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				v17 = l1
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				if v18 <= int32(0) {
					v64 = v3
					v67 = F_palloc0(m, int32(20))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v67)+16)) = v17
						*(*int32)(unsafe.Add(mBase, uint32(v67)+12)) = v64
						*(*int32)(unsafe.Add(mBase, uint32(v67)+8)) = v8
						*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(385)
						return v67
					}
				} else {
					v23 = F_palloc(m, v18<<(uint(int32(1))%32))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						if l0 == int32(0) {
							v64 = v23
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v27 <= int32(0) {
								v64 = v23
							} else {
								v33 = int32(0)
								v35 = v3
								for {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v33<<(uint(int32(2))%32))))
									v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+26)))
									if v43 == int32(0) {
										v47 = int32(1)
										v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+8)))
										*(*uint16)(unsafe.Add(mBase, uint32(v23+base.I32_extend16_s(v35)<<(uint(v47)%32)))) = uint16(v50)
										v54 = v35 + v47
									} else {
										v54 = v35
									}
									v56 = v33 + int32(1)
									v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v56 < v57 {
										v33 = v56
										v35 = v54
										continue
									} else {
										break
									}
									break
								}
								v64 = v23
							}
						}
						v67 = F_palloc0(m, int32(20))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v67)+16)) = v17
							*(*int32)(unsafe.Add(mBase, uint32(v67)+12)) = v64
							*(*int32)(unsafe.Add(mBase, uint32(v67)+8)) = v8
							*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(385)
							return v67
						}
					}
				}
			}
		} else {
			v15 = F_MakeSingleTupleTableSlot(m, v8, int32(1596068))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = v15
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				if v18 <= int32(0) {
					v64 = v3
					v67 = F_palloc0(m, int32(20))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v67)+16)) = v17
						*(*int32)(unsafe.Add(mBase, uint32(v67)+12)) = v64
						*(*int32)(unsafe.Add(mBase, uint32(v67)+8)) = v8
						*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(385)
						return v67
					}
				} else {
					v23 = F_palloc(m, v18<<(uint(int32(1))%32))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						if l0 == int32(0) {
							v64 = v23
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v27 <= int32(0) {
								v64 = v23
							} else {
								v33 = int32(0)
								v35 = v3
								for {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v33<<(uint(int32(2))%32))))
									v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+26)))
									if v43 == int32(0) {
										v47 = int32(1)
										v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+8)))
										*(*uint16)(unsafe.Add(mBase, uint32(v23+base.I32_extend16_s(v35)<<(uint(v47)%32)))) = uint16(v50)
										v54 = v35 + v47
									} else {
										v54 = v35
									}
									v56 = v33 + int32(1)
									v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v56 < v57 {
										v33 = v56
										v35 = v54
										continue
									} else {
										break
									}
									break
								}
								v64 = v23
							}
						}
						v67 = F_palloc0(m, int32(20))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v67)+16)) = v17
							*(*int32)(unsafe.Add(mBase, uint32(v67)+12)) = v64
							*(*int32)(unsafe.Add(mBase, uint32(v67)+8)) = v8
							*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(385)
							return v67
						}
					}
				}
			}
		}
	}
}
func F_ExecInitRoutingInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
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
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	v7 = l6
	v10 = int32(4476144)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v13
	v15 = F_ExecGetRootToChildMap(m, l4, l1)
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
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v20 = F_table_slot_create(m, v17, l1+int32(104))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v22 = int32(0)
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+204)) = v22
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l4)+84))
	if v25 == int32(0) {
		v49 = int32(1)
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v22 = v20
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+208)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+104)) = v49
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v55 = v53 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v55 < v57 {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	m.T0[v28].(func(*base.Module, int32, int32))(m, l0, l4)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v35 = v25
	goto L11
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+60))
	if v37 == int32(0) {
		v49 = int32(1)
		goto L7
	} else {
		goto L14
	}
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l4)+84))
	if v32 == int32(0) {
		v49 = int32(1)
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v35 = v32
	goto L11
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	if v41 == int32(0) {
		v49 = int32(1)
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v44 = m.T0[v37].(func(*base.Module, int32) int32)(m, l4)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v49 = v44
	goto L7
L17:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v86 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v85+v53<<(uint(v86)%32)))) = l4
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v90+v53))) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(l3+l5<<(uint(v86)%32))+24)) = v53
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v11
	return
L18:
	;
	if v57 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = int32(8)
	v64 = F_palloc(m, int32(32))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v57 << (uint(int32(1)) % 32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v77 = F_repalloc(m, v74, v57<<(uint(int32(3))%32))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v64
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v68 = F_palloc(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v68
	goto L17
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v77
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v82 = F_repalloc(m, v80, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v82
	goto L17
}
func F_ExecMemoize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 float64
	_ = v50
	var v53 float64
	_ = v53
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v75 int64
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v97 int64
	_ = v97
	var v112 float64
	_ = v112
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v299 int64
	_ = v299
	var v302 int32
	_ = v302
	var v304 int64
	_ = v304
	var v306 int64
	_ = v306
	var v309 int64
	_ = v309
	var v310 int64
	_ = v310
	var v320 int64
	_ = v320
	var v325 int32
	_ = v325
	var v326 int64
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int64
	_ = v335
	var v345 int64
	_ = v345
	var v353 int32
	_ = v353
	var v360 float64
	_ = v360
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int64
	_ = v448
	var v450 int64
	_ = v450
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v592 int32
	_ = v592
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v605 int64
	_ = v605
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v656 int64
	_ = v656
	var v658 int64
	_ = v658
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v691 int64
	_ = v691
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v721 int32
	_ = v721
	var v727 int32
	_ = v727
	var v730 int64
	_ = v730
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int64
	_ = v748
	var v752 int64
	_ = v752
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v766 int64
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v779 int64
	_ = v779
	var v792 int64
	_ = v792
	var v795 int32
	_ = v795
	var v799 int64
	_ = v799
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v844 int32
	_ = v844
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v872 int64
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v896 int32
	_ = v896
	var v904 int64
	_ = v904
	var v905 int64
	_ = v905
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
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
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v1012 int32
	_ = v1012
	var v1024 int64
	_ = v1024
	var v1031 int32
	_ = v1031
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int64
	_ = v1062
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1102 int32
	_ = v1102
	var v1108 int32
	_ = v1108
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1119 int32
	_ = v1119
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1157 int32
	_ = v1157
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1170 int64
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v22 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	F_MemoryContextReset(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	switch v30 - int32(1) {
	case 0:
		goto L12
	case 1:
		goto L8
	case 2:
		goto L9
	case 3:
		goto L10
	case 4:
		v1206 = v2
		goto L7
	default:
		goto L11
	}
L7:
	;
	m.G0 = v18 + int32(16)
	return v1206
L8:
	;
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1192)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v1193
	if v1193 == int32(0) {
		goto L288
	} else {
		goto L289
	}
L9:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+52))
	if v1146 != 0 {
		goto L267
	} else {
		goto L268
	}
L10:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+52))
	if v1126 != 0 {
		goto L256
	} else {
		goto L257
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L4
	} else {
		goto L253
	}
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v33 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+92))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v40 = F_MemoryContextAllocZero(m, v38, int32(32))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+8))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	m.T0[v158].(func(*base.Module, int32))(m, v156)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L56
	}
L16:
	;
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+28)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = v38
	if v37 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L53
	}
L19:
	;
	if base.Ui64(v61) <= base.Ui64(int64(2)) {
		goto L29
	} else {
		goto L30
	}
L20:
	;
	v47 = v37
	goto L22
L21:
	;
	v47 = int32(1024)
	goto L22
L22:
	;
	v50 = base.F64_div(base.F64_convert_i32_u(v47), float64(0.9))
	if base.F64_ge(v50, float64(4.294967296e+09)) != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v53 = float64(4.294967296e+09)
	goto L25
L24:
	;
	v53 = v50
	goto L25
L25:
	;
	if base.F64_lt(v53, float64(1.8446744073709552e+19))&base.F64_ge(v53, float64(0)) != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v59 = base.I64_trunc_f64_u(v53)
	v61 = v59
	goto L19
L27:
	;
	goto L28
L28:
	;
	v61 = int64(0)
	goto L19
L29:
	;
	v64 = int64(2)
	goto L31
L30:
	;
	v64 = v61
	goto L31
L31:
	;
	v65 = int64(1)
	if v64&(v64-v65) == int64(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v75 = v64
	goto L34
L33:
	;
	v75 = v65 << (uint(int64(64)-base.I64_clz(v64)) % 64)
	goto L34
L34:
	;
	if base.Ui64(v75<<(uint(int64(4))%64)) < base.Ui64(int64(2147483647)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v84 = F_MemoryContextAllocExtended(m, v38, base.I32_wrap_i64(v75)<<(uint(int32(4))%32), int32(5))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L50
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v84
	v87 = int64(1)
	if v75&(v75-v87) == int64(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v97 = v75
	goto L41
L40:
	;
	v97 = v87 << (uint(int64(64)-base.I64_clz(v75)) % 64)
	goto L41
L41:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v97<<(uint(int64(4))%64)) {
		goto L18
	} else {
		goto L42
	}
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40))) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = base.I32_wrap_i64(v97) - int32(1)
	v112 = base.F64_mul(base.F64_convert_i64_u(v97), float64(0.9))
	if base.F64_lt(v112, float64(4.294967296e+09))&base.F64_ge(v112, float64(0)) != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v97 == int64(4294967296) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v118 = base.I32_trunc_f64_u(v112)
	v120 = v118
	goto L43
L45:
	;
	goto L46
L46:
	;
	v120 = int32(0)
	goto L43
L47:
	;
	v121 = int32(-85899346)
	goto L49
L48:
	;
	v121 = v120
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v40
	goto L16
L50:
	;
	F_errmsg_internal(m, int32(396342), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(323259), int32(327), int32(337432))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
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
	F_errmsg_internal(m, int32(396342), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(323259), int32(327), int32(337432))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
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
	v161 = int32(4476144)
	v162 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v165
	if v155 <= int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v162
	v271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+4)))
	v273 = v271 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v156)+4)) = uint16(v273)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	*(*uint16)(unsafe.Add(mBase, uint32(v156)+6)) = uint16(v276)
	goto L69
L58:
	;
	v169 = int32(1)
	if v155 != v169 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v178 = v2
	v180 = int32(0)
	goto L62
L60:
	;
	v226 = v2
	goto L61
L61:
	;
	if v155&v169 == int32(0) {
		goto L57
	} else {
		goto L67
	}
L62:
	;
	v192 = v178 << (uint(int32(2)) % 32)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v192+v193)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v156)+20))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v195)+20))
	v199 = m.T0[v198].(func(*base.Module, int32, int32, int32) int32)(m, v195, v164, v196+v178)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L64
	}
L63:
	;
	v226 = v220
	goto L61
L64:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v156)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v201+v192))) = v199
	v205 = v178 | int32(1)
	v207 = v205 << (uint(int32(2)) % 32)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v207+v208)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v156)+20))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v210)+20))
	v214 = m.T0[v213].(func(*base.Module, int32, int32, int32) int32)(m, v210, v164, v211+v205)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v156)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v216+v207))) = v214
	v219 = int32(2)
	v220 = v178 + v219
	v222 = v180 + v219
	if v222 != v155&int32(-2) {
		v178 = v220
		v180 = v222
		goto L62
	} else {
		goto L66
	}
L66:
	;
	goto L63
L67:
	;
	v242 = v226 << (uint(int32(2)) % 32)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v242+v243)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v156)+20))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v245)+20))
	v249 = m.T0[v248].(func(*base.Module, int32, int32, int32) int32)(m, v245, v164, v246+v226)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v156)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v251+v242))) = v249
	goto L57
L69:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v279 = F_MemoizeHash_hash(m, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v278)+16))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v278)+8))
	v285 = v281
	v292 = v282
	goto L71
L71:
	;
	if base.Ui32(v285) <= base.Ui32(v292) {
		goto L79
	} else {
		goto L80
	}
L73:
	;
	v1108 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v278)+16)) = v1108
	v285 = v1108
	v292 = v1102
	goto L71
L74:
	;
	v1206 = int32(0)
	goto L7
L75:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+52))
	if v1044 != 0 {
		goto L232
	} else {
		goto L233
	}
L76:
	;
	v856 = int32(4476144)
	v857 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v859 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v859
	v862 = F_palloc(m, int32(12))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L4
	} else {
		goto L193
	}
L77:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v278)+8))
	v833 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v278)+8)) = v832 + v833
	*(*uint8)(unsafe.Add(mBase, uint32(v820)+12)) = uint8(v833)
	*(*int32)(unsafe.Add(mBase, uint32(v820)+8)) = v279
	*(*int32)(unsafe.Add(mBase, uint32(v820))) = int32(0)
	v844 = v820
	goto L76
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L4
	} else {
		goto L190
	}
L79:
	;
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v278)))
	if v299 == int64(4294967296) {
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v535 = int32(0)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v278)+20))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	v538 = v537 & v279
	v541 = v536 + v538<<(uint(int32(4))%32)
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541)+12)))
	if v542 == v535 {
		v820 = v541
		goto L77
	} else {
		goto L134
	}
L82:
	;
	v302 = int32(0)
	v304 = int64(2)
	v306 = v299 << (uint(int64(1)) % 64)
	if base.Ui64(v306) <= base.Ui64(v304) {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L81
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L4
	} else {
		goto L131
	}
L85:
	;
	v309 = v304
	goto L87
L86:
	;
	v309 = v306
	goto L87
L87:
	;
	v310 = int64(1)
	if v309&(v309-v310) == int64(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v320 = v309
	goto L90
L89:
	;
	v320 = v310 << (uint(int64(64)-base.I64_clz(v309)) % 64)
	goto L90
L90:
	;
	if base.Ui64(v320<<(uint(int64(4))%64)) < base.Ui64(int64(2147483647)) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v278)+20))
	v326 = *(*int64)(unsafe.Add(mBase, uint32(v278)))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v278)+24))
	v332 = F_MemoryContextAllocExtended(m, v327, base.I32_wrap_i64(v320)<<(uint(int32(4))%32), int32(5))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L4
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L4
	} else {
		goto L128
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278)+20)) = v332
	v335 = int64(1)
	if v320&(v320-v335) == int64(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v345 = v320
	goto L97
L96:
	;
	v345 = v335 << (uint(int64(64)-base.I64_clz(v320)) % 64)
	goto L97
L97:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v345<<(uint(int64(4))%64)) {
		goto L84
	} else {
		goto L98
	}
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v278))) = v345
	v353 = base.I32_wrap_i64(v345) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v278)+12)) = v353
	v360 = base.F64_mul(base.F64_convert_i64_u(v345), float64(0.9))
	if base.F64_lt(v360, float64(4.294967296e+09))&base.F64_ge(v360, float64(0)) != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if v345 == int64(4294967296) {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v366 = base.I32_trunc_f64_u(v360)
	v368 = v366
	goto L99
L101:
	;
	goto L102
L102:
	;
	v368 = int32(0)
	goto L99
L103:
	;
	v369 = int32(-85899346)
	goto L105
L104:
	;
	v369 = v368
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278)+16)) = v369
	if v326 != int64(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v375 = v302
	goto L110
L107:
	;
	goto L108
L108:
	;
	F_pfree(m, v325)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L4
	} else {
		goto L127
	}
L109:
	;
	v405 = v402
	v406 = v302
	goto L115
L110:
	;
	v390 = v325 + v375<<(uint(int32(4))%32)
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390)+12)))
	if v391 != int32(1) {
		v402 = v375
		goto L109
	} else {
		goto L112
	}
L111:
	;
	v402 = int32(0)
	goto L109
L112:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v390)+8))
	if v394&v353 == v375 {
		v402 = v375
		goto L109
	} else {
		goto L113
	}
L113:
	;
	v398 = v375 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v398)) < base.Ui64(v326) {
		v375 = v398
		goto L110
	} else {
		goto L114
	}
L114:
	;
	goto L111
L115:
	;
	v420 = v325 + v405<<(uint(int32(4))%32)
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420)+12)))
	if v421 == int32(1) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L108
L117:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v420)+8))
	v427 = v425
	goto L120
L118:
	;
	goto L119
L119:
	;
	v468 = v405 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v468)) < base.Ui64(v326) {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	v441 = v427 & v424
	v446 = v332 + v441<<(uint(int32(4))%32)
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446)+12)))
	if v447 != 0 {
		v427 = v441 + int32(1)
		goto L120
	} else {
		goto L122
	}
L121:
	;
	v448 = *(*int64)(unsafe.Add(mBase, uint32(v420)))
	*(*int64)(unsafe.Add(mBase, uint32(v446))) = v448
	v450 = *(*int64)(unsafe.Add(mBase, uint32(v420)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v446)+8)) = v450
	goto L119
L122:
	;
	goto L121
L123:
	;
	v472 = v468
	goto L125
L124:
	;
	v472 = int32(0)
	goto L125
L125:
	;
	v474 = v406 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v474)) < base.Ui64(v326) {
		v405 = v472
		v406 = v474
		goto L115
	} else {
		goto L126
	}
L126:
	;
	goto L116
L127:
	;
	goto L83
L128:
	;
	F_errmsg_internal(m, int32(396342), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(323259), int32(327), int32(337432))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	F_errmsg_internal(m, int32(396342), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(323259), int32(327), int32(337432))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	v546 = v538
	v548 = v541
	v549 = v537
	v550 = v535
	goto L135
L135:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v548)+8))
	if v560 == v279 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
	v703 = v701 + int32(4)
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v703 != v704 {
		goto L170
	} else {
		goto L171
	}
L137:
	;
	goto L136
L138:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
	v563 = F_MemoizeHash_equal(m, v278, v562)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L4
	} else {
		goto L141
	}
L139:
	;
	v567 = v560
	v568 = v549
	goto L140
L140:
	;
	v569 = v567 & v568
	if base.Ui32(v546) < base.Ui32(v569) {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	if v563 != 0 {
		goto L137
	} else {
		goto L142
	}
L142:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v548)+8))
	v567 = v566
	v568 = v565
	goto L140
L143:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	v573 = v546 + v571
	goto L145
L144:
	;
	v573 = v546
	goto L145
L145:
	;
	v576 = v568 & (v546 + int32(1))
	if base.Ui32(v573-v569) < base.Ui32(v550) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v582 = v536 + v576<<(uint(int32(4))%32)
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582)+12)))
	if v583 != 0 {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	goto L148
L148:
	;
	v686 = v550 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v686) {
		goto L165
	} else {
		goto L166
	}
L149:
	;
	v586 = v576
	v592 = int32(0)
	goto L152
L150:
	;
	v620 = v576
	v623 = v582
	goto L151
L151:
	;
	if v546 != v620 {
		goto L159
	} else {
		goto L160
	}
L152:
	;
	v600 = v592 + int32(1)
	if int32(151) <= v600 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	v620 = v613
	v623 = v616
	goto L151
L154:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v278)+8))
	v605 = *(*int64)(unsafe.Add(mBase, uint32(v278)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v603), base.F64_convert_i64_u(v605)), float64(0.1)) != 0 {
		v1102 = v603
		goto L73
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v613 = (v586 + int32(1)) & v568
	v616 = v536 + v613<<(uint(int32(4))%32)
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v616)+12)))
	if v617 != 0 {
		v586 = v613
		v592 = v600
		goto L152
	} else {
		goto L158
	}
L157:
	;
	goto L156
L158:
	;
	goto L153
L159:
	;
	v636 = v620
	v639 = v623
	goto L162
L160:
	;
	goto L161
L161:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v278)+8))
	v677 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v278)+8)) = v676 + v677
	*(*uint8)(unsafe.Add(mBase, uint32(v548)+12)) = uint8(v677)
	*(*int32)(unsafe.Add(mBase, uint32(v548)+8)) = v279
	*(*int32)(unsafe.Add(mBase, uint32(v548))) = int32(0)
	v844 = v548
	goto L76
L162:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	v652 = v649 & (v636 - int32(1))
	v655 = v536 + v652<<(uint(int32(4))%32)
	v656 = *(*int64)(unsafe.Add(mBase, uint32(v655)))
	*(*int64)(unsafe.Add(mBase, uint32(v639))) = v656
	v658 = *(*int64)(unsafe.Add(mBase, uint32(v655)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v639)+8)) = v658
	if v546 != v652 {
		v636 = v652
		v639 = v655
		goto L162
	} else {
		goto L164
	}
L163:
	;
	goto L161
L164:
	;
	goto L163
L165:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v278)+8))
	v691 = *(*int64)(unsafe.Add(mBase, uint32(v278)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v689), base.F64_convert_i64_u(v691)), float64(0.1)) != 0 {
		v1102 = v689
		goto L73
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v699 = v536 + v576<<(uint(int32(4))%32)
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699)+12)))
	if v700 != 0 {
		v546 = v576
		v548 = v699
		v549 = v568
		v550 = v686
		goto L135
	} else {
		goto L169
	}
L168:
	;
	goto L167
L169:
	;
	v820 = v699
	goto L77
L170:
	;
	v707 = l0 + int32(180)
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v701)+4))
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v701)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v708)+4)) = v709
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v701)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v709))) = v711
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if v713 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	goto L172
L172:
	;
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+13)))
	if v727 == int32(1) {
		goto L177
	} else {
		goto L178
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = l0 + int32(180)
	goto L175
L174:
	;
	goto L175
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v701)+8)) = v707
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v707)))
	*(*int32)(unsafe.Add(mBase, uint32(v701)+4)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v721)+4)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v707))) = v703
	goto L172
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(5)
	goto L74
L177:
	;
	v730 = *(*int64)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+200)) = v730 + int64(1)
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v548)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v548
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v734
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v548)+4))
	if v737 == int32(0) {
		goto L176
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v748 = *(*int64)(unsafe.Add(mBase, uint32(l0)+208))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v748 + int64(1)
	v752 = int64(0)
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v548)+4))
	if v753 != 0 {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(2)
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v548)+4))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v742)))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v746 = F_ExecStoreMinimalTuple(m, v743, v744, int32(0))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	v1206 = v744
	goto L7
L182:
	;
	v756 = v753
	v766 = v752
	goto L185
L183:
	;
	v792 = v752
	goto L184
L184:
	;
	v795 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v548)+4)) = v795
	*(*uint8)(unsafe.Add(mBase, uint32(v548)+13)) = uint8(v795)
	v799 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v799 - v792
	v1031 = v548
	goto L75
L185:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v756)+4))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v756)))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v770)))
	F_pfree(m, v770)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L4
	} else {
		goto L187
	}
L186:
	;
	v792 = v779
	goto L184
L187:
	;
	F_pfree(m, v756)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L4
	} else {
		goto L188
	}
L188:
	;
	v779 = v766 + base.I64_extend_i32_u(v771+int32(8))
	if v769 != 0 {
		v756 = v769
		v766 = v779
		goto L185
	} else {
		goto L189
	}
L189:
	;
	goto L186
L190:
	;
	F_errmsg_internal(m, int32(457558), int32(0))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L4
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(323259), int32(630), int32(308586))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v844))) = v862
	v865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v865)+8))
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v867)+48))
	v869 = m.T0[v868].(func(*base.Module, int32, int32) int32)(m, v865, int32(0))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L4
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v862))) = v869
	v872 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v844)))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v873)))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v874)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v872 + base.I64_extend_i32_u(v875+int32(28))
	v881 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v844)+4)) = v881
	*(*uint8)(unsafe.Add(mBase, uint32(v844)+13)) = uint8(v881)
	v886 = l0 + int32(180)
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v844)))
	v889 = v887 + int32(4)
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if v890 == v881 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v886
	goto L197
L196:
	;
	goto L197
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v887)+8)) = v886
	v896 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	*(*int32)(unsafe.Add(mBase, uint32(v887)+4)) = v896
	*(*int32)(unsafe.Add(mBase, uint32(v896)+4)) = v889
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v889
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v857
	v904 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v905 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	if base.Ui64(v904) <= base.Ui64(v905) {
		v1012 = v844
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v1024 = *(*int64)(unsafe.Add(mBase, uint32(l0)+208))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v1024 + int64(1)
	v1031 = v1012
	goto L75
L199:
	;
	v907 = F_cache_reduce_memory(m, l0, v862)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L4
	} else {
		goto L201
	}
L200:
	;
	v1012 = int32(0)
	goto L198
L201:
	;
	if v907 == int32(0) {
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v844)+12)))
	if v911 == int32(1) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v844)))
	if v914 == v862 {
		v1012 = v844
		goto L198
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v918)+8))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v919)+12))
	m.T0[v920].(func(*base.Module, int32))(m, v918)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L4
	} else {
		goto L207
	}
L206:
	;
	goto L205
L207:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v862)))
	v925 = F_ExecStoreMinimalTuple(m, v923, v917, int32(0))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L4
	} else {
		goto L208
	}
L208:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v917)+12))
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v927)))
	v929 = int32(*(*int16)(unsafe.Add(mBase, uint32(v917)+6)))
	if v929 < v928 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	F_slot_getsomeattrs_int(m, v917, v928)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L4
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v918)+16))
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v917)+16))
	v936 = v916 << (uint(int32(2)) % 32)
	if v936 != 0 {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	goto L211
L213:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v918)+20))
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v917)+20))
	if v916 != 0 {
		goto L218
	} else {
		goto L219
	}
L214:
	;
	v937 = F__emscripten_memcpy_bulkmem(m, v933, v934, v936)
	mBase = m.M
	goto L216
L215:
	;
	goto L216
L216:
	;
	goto L213
L217:
	;
	v943 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v918)+4)))
	v945 = v943 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v918)+4)) = uint16(v945)
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v918)+12))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v947)))
	*(*uint16)(unsafe.Add(mBase, uint32(v918)+6)) = uint16(v948)
	goto L221
L218:
	;
	v941 = F__emscripten_memcpy_bulkmem(m, v939, v940, v916)
	mBase = m.M
	goto L220
L219:
	;
	goto L220
L220:
	;
	goto L217
L221:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v951 = F_MemoizeHash_hash(m, v950)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L4
	} else {
		goto L222
	}
L222:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v950)+20))
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v950)+12))
	v955 = v951 & v954
	v958 = v953 + v955<<(uint(int32(4))%32)
	v959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958)+12)))
	if v959 == int32(0) {
		goto L200
	} else {
		goto L223
	}
L223:
	;
	v964 = v955
	v965 = v958
	v966 = v953
	v967 = v954
	goto L224
L224:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v965)+8))
	if v977 == v951 {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	goto L200
L226:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v965)))
	v980 = F_MemoizeHash_equal(m, v950, v979)
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L4
	} else {
		goto L229
	}
L227:
	;
	v984 = v966
	v985 = v967
	goto L228
L228:
	;
	v988 = v985 & (v964 + int32(1))
	v991 = v984 + v988<<(uint(int32(4))%32)
	v992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v991)+12)))
	if v992 != 0 {
		v964 = v988
		v965 = v991
		v966 = v984
		v967 = v985
		goto L224
	} else {
		goto L231
	}
L229:
	;
	if v980 != 0 {
		v1012 = v965
		goto L198
	} else {
		goto L230
	}
L230:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v950)+12))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v950)+20))
	v984 = v983
	v985 = v982
	goto L228
L231:
	;
	goto L225
L232:
	;
	F_ExecReScan(m, v1043)
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L4
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+12))
	v1048 = m.T0[v1047].(func(*base.Module, int32) int32)(m, v1043)
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L4
	} else {
		goto L237
	}
L235:
	;
	goto L234
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v1031
	if v1031 != 0 {
		goto L247
	} else {
		goto L248
	}
L237:
	;
	if v1048 != 0 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v1050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1048)+4)))
	if v1050&int32(2) == int32(0) {
		goto L236
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	if v1031 != 0 {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	goto L240
L242:
	;
	v1055 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1031)+13)) = uint8(v1055)
	goto L244
L243:
	;
	goto L244
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(5)
	goto L74
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v1070
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1072)+8))
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1073)+32))
	m.T0[v1074].(func(*base.Module, int32, int32))(m, v1072, v1048)
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L4
	} else {
		goto L252
	}
L246:
	;
	v1067 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1031)+13)) = uint8(v1067)
	v1070 = int32(3)
	goto L245
L247:
	;
	v1060 = F_cache_store_tuple(m, l0, v1048)
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L4
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v1062 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+224)) = v1062 + int64(1)
	v1070 = int32(4)
	goto L245
L250:
	;
	if v1060 != 0 {
		goto L246
	} else {
		goto L251
	}
L251:
	;
	goto L249
L252:
	;
	v1206 = v1072
	goto L7
L253:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v1115
	F_errmsg_internal(m, int32(478508), v18)
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L4
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(492917), int32(946), int32(337775))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L4
	} else {
		goto L255
	}
L255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L256:
	;
	F_ExecReScan(m, v1125)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L4
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+12))
	v1130 = m.T0[v1129].(func(*base.Module, int32) int32)(m, v1125)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L4
	} else {
		goto L261
	}
L259:
	;
	goto L258
L260:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1139)+8))
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1140)+32))
	m.T0[v1141].(func(*base.Module, int32, int32))(m, v1139, v1130)
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L4
	} else {
		goto L266
	}
L261:
	;
	if v1130 != 0 {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v1132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1130)+4)))
	if v1132&int32(2) == int32(0) {
		goto L260
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(5)
	v1206 = v2
	goto L7
L265:
	;
	goto L264
L266:
	;
	v1206 = v1139
	goto L7
L267:
	;
	F_ExecReScan(m, v1145)
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L4
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+12))
	v1150 = m.T0[v1149].(func(*base.Module, int32) int32)(m, v1145)
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L4
	} else {
		goto L272
	}
L270:
	;
	goto L269
L271:
	;
	v1161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1144)+13)))
	if v1161 != int32(1) {
		goto L277
	} else {
		goto L278
	}
L272:
	;
	if v1150 != 0 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1150)+4)))
	if v1152&int32(2) == int32(0) {
		goto L271
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	v1157 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1144)+13)) = uint8(v1157)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(5)
	v1206 = v2
	goto L7
L276:
	;
	goto L275
L277:
	;
	v1164 = F_cache_store_tuple(m, l0, v1150)
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L4
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L4
	} else {
		goto L285
	}
L280:
	;
	if v1164 == int32(0) {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(4)
	v1170 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+224)) = v1170 + int64(1)
	goto L283
L282:
	;
	goto L283
L283:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1174)+8))
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1175)+32))
	m.T0[v1176].(func(*base.Module, int32, int32))(m, v1174, v1150)
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L4
	} else {
		goto L284
	}
L284:
	;
	v1206 = v1174
	goto L7
L285:
	;
	F_errmsg_internal(m, int32(346918), int32(0))
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L4
	} else {
		goto L286
	}
L286:
	;
	F_errfinish(m, int32(492917), int32(893), int32(337775))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L4
	} else {
		goto L287
	}
L287:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(5)
	v1206 = v2
	goto L7
L289:
	;
	goto L290
L290:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1193)))
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v1202 = F_ExecStoreMinimalTuple(m, v1199, v1200, int32(0))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L4
	} else {
		goto L291
	}
L291:
	;
	v1206 = v1200
	goto L7
}
func F_ExecNestLoop(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v137 int32
	_ = v137
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 float64
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 float64
	_ = v238
	var v242 int32
	_ = v242
	var v245 float64
	_ = v245
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	F_MemoryContextReset(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	goto L8
L7:
	;
	m.G0 = v17 + int32(16)
	return v285
L8:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v48 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+72))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v256)+16))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+8))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	m.T0[v260].(func(*base.Module, int32))(m, v258)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L4
	} else {
		goto L68
	}
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	if v51 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	if v152 != 0 {
		goto L32
	} else {
		goto L33
	}
L13:
	;
	F_ExecReScan(m, v26)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v54 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, v26)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	if v56 == int32(0) {
		v285 = v54
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
	if v60&int32(2) != 0 {
		v285 = v54
		goto L7
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v56
	v64 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+116)) = uint16(v64)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v29)+88))
	if v66 == v64 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_ExecReScan(m, v25)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L31
	}
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v69 <= int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v74 = v54
	goto L23
L23:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v74<<(uint(int32(2))%32))))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v95 = v86 + v92*int32(12)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v96)+8)))
	v98 = int32(*(*int16)(unsafe.Add(mBase, uint32(v56)+6)))
	if v98 < v97 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L20
L25:
	;
	F_slot_getsomeattrs_int(m, v56, v97)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v103 = v97 - int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103+v104))))
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+8)) = uint8(v106)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v103<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	v115 = F_bms_add_member(m, v114, v92)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+52)) = v115
	v119 = v74 + int32(1)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v119 < v120 {
		v74 = v119
		goto L23
	} else {
		goto L30
	}
L30:
	;
	goto L24
L31:
	;
	goto L12
L32:
	;
	F_ExecReScan(m, v25)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v156 = m.T0[v155].(func(*base.Module, int32) int32)(m, v25)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v156
	if v156 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L9
L38:
	;
	if v28 != 0 {
		goto L51
	} else {
		goto L52
	}
L39:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+4)))
	if v159&int32(2) == int32(0) {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v164 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)) = uint8(v164)
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	if v166 != 0 {
		goto L8
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	switch v167 - int32(1) {
	case 0, 4:
		goto L44
	default:
		goto L8
	}
L44:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v170
	if v27 == int32(0) {
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v174 = int32(4476144)
	v175 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v177
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v182 = m.T0[v181].(func(*base.Module, int32, int32, int32) int32)(m, v27, v30, v17+int32(13))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v175
	if v182 != 0 {
		goto L37
	} else {
		goto L47
	}
L47:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v186 == int32(0) {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	v189 = *(*float64)(unsafe.Add(mBase, uint32(v186)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v186)+248)) = base.F64_add(v189, float64(1))
	goto L8
L49:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	F_MemoryContextReset(m, v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L67
	}
L50:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v242 == int32(0) {
		goto L49
	} else {
		goto L66
	}
L51:
	;
	v193 = int32(4476144)
	v194 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v196
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v201 = m.T0[v200].(func(*base.Module, int32, int32, int32) int32)(m, v28, v30, v17+int32(14))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v209 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)) = uint8(v209)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v211 == int32(5) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v194
	if v201 == int32(0) {
		goto L50
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v214 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)) = uint8(v214)
	goto L8
L57:
	;
	goto L58
L58:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	if v216 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v219 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)) = uint8(v219)
	goto L61
L60:
	;
	goto L61
L61:
	;
	if v27 == int32(0) {
		goto L37
	} else {
		goto L62
	}
L62:
	;
	v223 = int32(4476144)
	v224 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v226
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v231 = m.T0[v230].(func(*base.Module, int32, int32, int32) int32)(m, v27, v30, v17+int32(15))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v224
	if v231 != 0 {
		goto L37
	} else {
		goto L64
	}
L64:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v235 == int32(0) {
		goto L49
	} else {
		goto L65
	}
L65:
	;
	v238 = *(*float64)(unsafe.Add(mBase, uint32(v235)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v235)+248)) = base.F64_add(v238, float64(1))
	goto L49
L66:
	;
	v245 = *(*float64)(unsafe.Add(mBase, uint32(v242)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v242)+240)) = base.F64_add(v245, float64(1))
	goto L49
L67:
	;
	goto L8
L68:
	;
	v263 = int32(4476144)
	v264 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v257)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v266
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v256)+24))
	v272 = m.T0[v271].(func(*base.Module, int32, int32, int32) int32)(m, v256+int32(4), v257, int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v264
	v276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258)+4)))
	v278 = v276 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v258)+4)) = uint16(v278)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v258)+12))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	*(*uint16)(unsafe.Add(mBase, uint32(v258)+6)) = uint16(v281)
	v285 = v258
	goto L7
}
func F_ExecRestrPos(m *base.Module, l0 int32) {
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int64
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int64
	_ = v146
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v14 - int32(394) {
	case 0:
		v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if v218 != 0 {
			F_ExecRestrPos(m, v218)
			mBase = m.M
			v220 = m.ExcPending
			if v220 != 0 {
				return
			} else {
				m.G0 = v12 + int32(16)
				return
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v224 = m.ExcPending
			if v224 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(360163), int32(0))
				mBase = m.M
				v228 = m.ExcPending
				if v228 != 0 {
					return
				} else {
					F_errfinish(m, int32(487726), int32(168), int32(135337))
					mBase = m.M
					v233 = m.ExcPending
					if v233 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v237 = m.ExcPending
		if v237 != 0 {
			return
		} else {
			v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v238
			F_errmsg_internal(m, int32(480638), v12)
			mBase = m.M
			v242 = m.ExcPending
			if v242 != 0 {
				return
			} else {
				F_errfinish(m, int32(492309), int32(405), int32(135375))
				mBase = m.M
				v247 = m.ExcPending
				if v247 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 11:
		v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+156))
		if v249 == int32(0) {
			v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
			F_index_restrpos(m, v286)
			mBase = m.M
			v288 = m.ExcPending
			if v288 != 0 {
				return
			} else {
				m.G0 = v12 + int32(16)
				return
			}
		} else {
			v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+72))
			v255 = v253 - int32(1)
			v257 = v255 << (uint(int32(2)) % 32)
			v258 = *(*int32)(unsafe.Add(mBase, uint32(v249)+16))
			v260 = *(*int32)(unsafe.Add(mBase, uint32(v257+v258)))
			if v260 == int32(0) {
				v263 = *(*int32)(unsafe.Add(mBase, uint32(v249)+36))
				v265 = *(*int32)(unsafe.Add(mBase, uint32(v263+v257)))
				if v265 == int32(0) {
					v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
					F_index_restrpos(m, v286)
					mBase = m.M
					v288 = m.ExcPending
					if v288 != 0 {
						return
					} else {
						m.G0 = v12 + int32(16)
						return
					}
				} else {
					v268 = *(*int32)(unsafe.Add(mBase, uint32(v249)+40))
					v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268+v255))))
					if v270 != 0 {
						m.G0 = v12 + int32(16)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v274 = m.ExcPending
						if v274 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(314617), int32(0))
							mBase = m.M
							v278 = m.ExcPending
							if v278 != 0 {
								return
							} else {
								F_errfinish(m, int32(491339), int32(889), int32(135319))
								mBase = m.M
								v283 = m.ExcPending
								if v283 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			} else {
				v268 = *(*int32)(unsafe.Add(mBase, uint32(v249)+40))
				v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268+v255))))
				if v270 != 0 {
					m.G0 = v12 + int32(16)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v274 = m.ExcPending
					if v274 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(314617), int32(0))
						mBase = m.M
						v278 = m.ExcPending
						if v278 != 0 {
							return
						} else {
							F_errfinish(m, int32(491339), int32(889), int32(135319))
							mBase = m.M
							v283 = m.ExcPending
							if v283 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	case 12:
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+156))
		if v18 == int32(0) {
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
			F_index_restrpos(m, v55)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return
			} else {
				m.G0 = v12 + int32(16)
				return
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+72))
			v24 = v22 - int32(1)
			v26 = v24 << (uint(int32(2)) % 32)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v26+v27)))
			if v29 == int32(0) {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v32+v26)))
				if v34 == int32(0) {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
					F_index_restrpos(m, v55)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						m.G0 = v12 + int32(16)
						return
					}
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v24))))
					if v39 != 0 {
						m.G0 = v12 + int32(16)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(314564), int32(0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								F_errfinish(m, int32(491297), int32(508), int32(135297))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v24))))
				if v39 != 0 {
					m.G0 = v12 + int32(16)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(314564), int32(0))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							F_errfinish(m, int32(491297), int32(508), int32(135297))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	case 25:
		v60 = m.G0
		v62 = v60 - int32(16)
		m.G0 = v62
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
		v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
		if v65 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
					*(*int32)(unsafe.Add(mBase, uint32(v62))) = v76
					F_errmsg(m, int32(135474), v62)
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						F_errfinish(m, int32(491658), int32(156), int32(135356))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			m.T0[v65].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return
			} else {
				m.G0 = v62 + int32(16)
				m.G0 = v12 + int32(16)
				return
			}
		}
	case 30:
		v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
		if v91 != 0 {
			F_tuplestore_copy_read_pointer(m, v91, int32(1), int32(0))
			mBase = m.M
			v95 = m.ExcPending
			if v95 != 0 {
				return
			} else {
				m.G0 = v12 + int32(16)
				return
			}
		} else {
			m.G0 = v12 + int32(16)
			return
		}
	case 32:
		v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)))
		if v96 == int32(1) {
			v99 = int32(4476144)
			v100 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+28))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v103
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v102)+64))
			switch v105 - int32(3) {
			case 0:
				v198 = *(*int32)(unsafe.Add(mBase, uint32(v102)+224))
				*(*int32)(unsafe.Add(mBase, uint32(v102)+204)) = v198
				v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+228)))
				*(*uint8)(unsafe.Add(mBase, uint32(v102)+208)) = uint8(v206)
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v100
				m.G0 = v12 + int32(16)
				return
			case 1:
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v102)+200))
				v109 = *(*int64)(unsafe.Add(mBase, uint32(v102)+216))
				v110 = *(*int32)(unsafe.Add(mBase, uint32(v102)+224))
				v111 = m.G0
				v113 = v111 - int32(16)
				m.G0 = v113
				v115 = *(*int32)(unsafe.Add(mBase, uint32(v108)+40))
				if v115 == int32(0) {
					v118 = *(*int32)(unsafe.Add(mBase, uint32(v108)+44))
					v119 = F_palloc(m, v118)
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v108)+40)) = v119
						*(*int64)(unsafe.Add(mBase, uint32(v108)+52)) = int64(0)
						v124 = *(*int64)(unsafe.Add(mBase, uint32(v108)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v108)+24)) = v124
						v126 = F_ltsReadFillBuffer(m, v108)
						mBase = m.M
						v127 = m.ExcPending
						if v127 != 0 {
							return
						} else {
							v128 = *(*int64)(unsafe.Add(mBase, uint32(v108)+16))
							if v128 == v109 {
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v108)+56))
								v151 = v130
								if v151 < v110 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v175 = m.ExcPending
									if v175 != 0 {
										return
									} else {
										F_errmsg_internal(m, int32(247074), int32(0))
										mBase = m.M
										v179 = m.ExcPending
										if v179 != 0 {
											return
										} else {
											F_errfinish(m, int32(493349), int32(1151), int32(312723))
											mBase = m.M
											v184 = m.ExcPending
											if v184 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v108)+52)) = v110
									m.G0 = v113 + int32(16)
									v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+228)))
									*(*uint8)(unsafe.Add(mBase, uint32(v102)+208)) = uint8(v206)
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v100
									m.G0 = v12 + int32(16)
									return
								}
							} else {
								v131 = *(*int32)(unsafe.Add(mBase, uint32(v108)+40))
								v132 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
								v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
								v134 = F_BufFileSeekBlock(m, v133, v109)
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return
								} else {
									if v134 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v160 = m.ExcPending
										if v160 != 0 {
											return
										} else {
											F_errcode_for_file_access(m)
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v113))) = v109
												F_errmsg(m, int32(382821), v113)
												mBase = m.M
												v166 = m.ExcPending
												if v166 != 0 {
													return
												} else {
													F_errfinish(m, int32(493349), int32(288), int32(313944))
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v136 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
										F_BufFileReadExact(m, v136, v131, int32(8192))
										mBase = m.M
										v139 = m.ExcPending
										if v139 != 0 {
											return
										} else {
											v140 = int32(8176)
											*(*int32)(unsafe.Add(mBase, uint32(v108)+56)) = v140
											*(*int64)(unsafe.Add(mBase, uint32(v108)+16)) = v109
											v143 = *(*int32)(unsafe.Add(mBase, uint32(v108)+40))
											v146 = *(*int64)(unsafe.Add(mBase, uint32(v143)+uint32(_consts[520])))
											*(*int64)(unsafe.Add(mBase, uint32(v108)+24)) = v146
											v151 = v140
											if v151 < v110 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return
												} else {
													F_errmsg_internal(m, int32(247074), int32(0))
													mBase = m.M
													v179 = m.ExcPending
													if v179 != 0 {
														return
													} else {
														F_errfinish(m, int32(493349), int32(1151), int32(312723))
														mBase = m.M
														v184 = m.ExcPending
														if v184 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v108)+52)) = v110
												m.G0 = v113 + int32(16)
												v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+228)))
												*(*uint8)(unsafe.Add(mBase, uint32(v102)+208)) = uint8(v206)
												*(*int32)(unsafe.Add(mBase, _consts[0])) = v100
												m.G0 = v12 + int32(16)
												return
											}
										}
									}
								}
							}
						}
					}
				} else {
					v128 = *(*int64)(unsafe.Add(mBase, uint32(v108)+16))
					if v128 == v109 {
						v130 = *(*int32)(unsafe.Add(mBase, uint32(v108)+56))
						v151 = v130
						if v151 < v110 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v175 = m.ExcPending
							if v175 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(247074), int32(0))
								mBase = m.M
								v179 = m.ExcPending
								if v179 != 0 {
									return
								} else {
									F_errfinish(m, int32(493349), int32(1151), int32(312723))
									mBase = m.M
									v184 = m.ExcPending
									if v184 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v108)+52)) = v110
							m.G0 = v113 + int32(16)
							v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+228)))
							*(*uint8)(unsafe.Add(mBase, uint32(v102)+208)) = uint8(v206)
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v100
							m.G0 = v12 + int32(16)
							return
						}
					} else {
						v131 = *(*int32)(unsafe.Add(mBase, uint32(v108)+40))
						v132 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
						v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
						v134 = F_BufFileSeekBlock(m, v133, v109)
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return
						} else {
							if v134 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v160 = m.ExcPending
								if v160 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v113))) = v109
										F_errmsg(m, int32(382821), v113)
										mBase = m.M
										v166 = m.ExcPending
										if v166 != 0 {
											return
										} else {
											F_errfinish(m, int32(493349), int32(288), int32(313944))
											mBase = m.M
											v171 = m.ExcPending
											if v171 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v136 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
								F_BufFileReadExact(m, v136, v131, int32(8192))
								mBase = m.M
								v139 = m.ExcPending
								if v139 != 0 {
									return
								} else {
									v140 = int32(8176)
									*(*int32)(unsafe.Add(mBase, uint32(v108)+56)) = v140
									*(*int64)(unsafe.Add(mBase, uint32(v108)+16)) = v109
									v143 = *(*int32)(unsafe.Add(mBase, uint32(v108)+40))
									v146 = *(*int64)(unsafe.Add(mBase, uint32(v143)+uint32(_consts[520])))
									*(*int64)(unsafe.Add(mBase, uint32(v108)+24)) = v146
									v151 = v140
									if v151 < v110 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v175 = m.ExcPending
										if v175 != 0 {
											return
										} else {
											F_errmsg_internal(m, int32(247074), int32(0))
											mBase = m.M
											v179 = m.ExcPending
											if v179 != 0 {
												return
											} else {
												F_errfinish(m, int32(493349), int32(1151), int32(312723))
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v108)+52)) = v110
										m.G0 = v113 + int32(16)
										v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+228)))
										*(*uint8)(unsafe.Add(mBase, uint32(v102)+208)) = uint8(v206)
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v100
										m.G0 = v12 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v188 = m.ExcPending
				if v188 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(348080), int32(0))
					mBase = m.M
					v192 = m.ExcPending
					if v192 != 0 {
						return
					} else {
						F_errfinish(m, int32(487438), int32(2485), int32(135161))
						mBase = m.M
						v197 = m.ExcPending
						if v197 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			m.G0 = v12 + int32(16)
			return
		}
	}
}
func F_ExecWorkTableScan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v3 == int32(0) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+92))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7+v9*int32(12))+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v13
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
		F_ExecSetSlotDescriptor(m, v15, v16)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			F_ExecAssignScanProjectionInfo(m, l0)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v26 = F_ExecScan(m, l0, int32(777), int32(778))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					return v26
				}
			}
		}
	} else {
		v26 = F_ExecScan(m, l0, int32(777), int32(778))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			return v26
		}
	}
}
func F_ExecuteGrantStmt(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v481 int32
	_ = v481
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int64
	_ = v529
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v574 int64
	_ = v574
	var v575 int32
	_ = v575
	var v579 int64
	_ = v579
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v605 int32
	_ = v605
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(208)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L9
	} else {
		goto L191
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L9
	} else {
		goto L188
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L9
	} else {
		goto L184
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L9
	} else {
		goto L181
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L9
	} else {
		goto L177
	}
L6:
	;
	v20 = F_get_rolespec_oid(m, v18, int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+64)) = uint8(v25)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	switch v29 {
	case 0:
		goto L14
	case 1:
		goto L13
	default:
		goto L4
	}
L9:
	;
	return
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	if v20 != v23 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+88)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v406
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+96)) = uint8(v421)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+100)) = v423
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v425 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L13:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v252 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L14:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v27 - int32(12) {
	case 0, 37:
		goto L16
	default:
		goto L18
	case 15:
		goto L19
	case 25, 29:
		goto L17
	}
L15:
	;
	v163 = int32(0)
	v168 = v2
	goto L58
L16:
	;
	if v30 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L17:
	;
	if v30 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L18:
	;
	if v30 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	if v30 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v406 = int32(0)
	goto L12
L21:
	;
	goto L22
L22:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if int32(0) < v36 {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v406 = int32(0)
	goto L12
L24:
	;
	v406 = int32(0)
	goto L12
L25:
	;
	goto L26
L26:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v43 <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v406 = int32(0)
	goto L12
L28:
	;
	goto L29
L29:
	;
	v49 = int32(0)
	v50 = v2
	goto L30
L30:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v50<<(uint(int32(2))%32))))
	v68 = int32(0)
	F_get_object_address(m, v16+int32(112), v27, v67, v68, int32(1), v68)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L9
	} else {
		goto L32
	}
L31:
	;
	v406 = v74
	goto L12
L32:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v16)+116))
	v74 = F_lappend_oid(m, v49, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	v77 = v50 + int32(1)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v77 < v78 {
		v49 = v74
		v50 = v77
		goto L30
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	v406 = int32(0)
	goto L12
L36:
	;
	goto L37
L37:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v83 <= int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v406 = int32(0)
	goto L12
L39:
	;
	goto L40
L40:
	;
	v89 = int32(0)
	v90 = v2
	goto L41
L41:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101+v90<<(uint(int32(2))%32))))
	v107 = int32(0)
	v110 = F_RangeVarGetRelidExtended(m, v105, int32(1), v107, v107, v107)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L9
	} else {
		goto L43
	}
L42:
	;
	v406 = v112
	goto L12
L43:
	;
	v112 = F_lappend_oid(m, v89, v110)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	v115 = v90 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v115 < v116 {
		v89 = v112
		v90 = v115
		goto L41
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	v406 = int32(0)
	goto L12
L47:
	;
	goto L48
L48:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v121 <= int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v406 = int32(0)
	goto L12
L50:
	;
	goto L51
L51:
	;
	v127 = int32(0)
	v128 = v2
	goto L52
L52:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+v128<<(uint(int32(2))%32))))
	v146 = F_makeTypeNameFromNameList(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L9
	} else {
		goto L54
	}
L53:
	;
	v406 = v155
	goto L12
L54:
	;
	F_get_object_address(m, v16+int32(112), v27, v146, v16+int32(108), int32(1), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v16)+116))
	v155 = F_lappend_oid(m, v127, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L9
	} else {
		goto L56
	}
L56:
	;
	v158 = v128 + int32(1)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v158 < v159 {
		v127 = v155
		v128 = v158
		goto L52
	} else {
		goto L57
	}
L57:
	;
	goto L53
L58:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175+v168<<(uint(int32(2))%32))))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	v182 = F_ParameterAclLookup(m, v180, int32(1))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L9
	} else {
		goto L61
	}
L59:
	;
	v406 = v247
	goto L12
L60:
	;
	if v240 != 0 {
		goto L78
	} else {
		goto L79
	}
L61:
	;
	if v182 != 0 {
		v240 = v182
		goto L60
	} else {
		goto L62
	}
L62:
	;
	if v25&int32(1) == int32(0) {
		v240 = v182
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v188 = m.G0
	v190 = v188 - int32(16)
	m.G0 = v190
	v192 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v190)+12)) = v192
	*(*uint16)(unsafe.Add(mBase, uint32(v190))) = uint16(v192)
	v199 = F_find_option(m, v180, v192, int32(1), int32(10))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	if v199 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v205 = F_assignable_custom_variable_name(m, v180, int32(0), int32(21))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L9
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v207 = F_convert_GUC_name_for_parameter_acl(m, v180)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L9
	} else {
		goto L69
	}
L68:
	;
	goto L67
L69:
	;
	v211 = F_table_open(m, int32(6243), int32(3))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L70
	}
L70:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v211)+52))
	v216 = F_GetNewOidWithIndex(m, v211, int32(6247), int32(1))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+4)) = v216
	v219 = F_cstring_to_text(m, v207)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L9
	} else {
		goto L72
	}
L72:
	;
	v221 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+2)) = uint8(v221)
	*(*int32)(unsafe.Add(mBase, uint32(v190)+8)) = v219
	v226 = F_heap_form_tuple(m, v213, v190+int32(4), v190)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	F_CatalogTupleInsert(m, v211, v226)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L9
	} else {
		goto L74
	}
L74:
	;
	F_pfree(m, v226)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L9
	} else {
		goto L75
	}
L75:
	;
	F_sequence_close(m, v211, int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	m.G0 = v190 + int32(16)
	F_CommandCounterIncrement(m)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L9
	} else {
		goto L77
	}
L77:
	;
	v240 = v216
	goto L60
L78:
	;
	v245 = F_lappend_oid(m, v163, v240)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L9
	} else {
		goto L81
	}
L79:
	;
	v247 = v163
	goto L80
L80:
	;
	v249 = v168 + int32(1)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v249 < v250 {
		v163 = v247
		v168 = v249
		goto L58
	} else {
		goto L82
	}
L81:
	;
	v247 = v245
	goto L80
L82:
	;
	goto L59
L83:
	;
	v406 = int32(0)
	goto L12
L84:
	;
	goto L85
L85:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	if v256 <= int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v406 = int32(0)
	goto L12
L87:
	;
	goto L88
L88:
	;
	v263 = v27 - int32(19)
	v266 = int32(0)
	v271 = v2
	goto L89
L89:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v278+v271<<(uint(int32(2))%32))))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	v285 = F_LookupExplicitNamespace(m, v283, int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L9
	} else {
		goto L91
	}
L90:
	;
	v406 = v389
	goto L12
L91:
	;
	switch v263 {
	case 0, 10, 15:
		goto L95
	default:
		goto L94
	case 18:
		goto L96
	case 22:
		goto L93
	}
L92:
	;
	v402 = v271 + int32(1)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	if v402 < v403 {
		v266 = v389
		v271 = v402
		goto L89
	} else {
		goto L128
	}
L93:
	;
	v364 = F_getRelationsInNamespace(m, v285, int32(114))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L9
	} else {
		goto L118
	}
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L9
	} else {
		goto L115
	}
L95:
	;
	v294 = int32(3)
	F_ScanKeyInit(m, v16+int32(112), v294, v294, int32(184), v285)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L9
	} else {
		goto L99
	}
L96:
	;
	v288 = F_getRelationsInNamespace(m, v285, int32(83))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L9
	} else {
		goto L97
	}
L97:
	;
	v290 = F_list_concat(m, v266, v288)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L9
	} else {
		goto L98
	}
L98:
	;
	v389 = v290
	goto L92
L99:
	;
	switch v263 {
	case 0:
		v302 = int32(70)
		goto L101
	default:
		v310 = int32(1)
		goto L100
	case 10:
		goto L102
	}
L100:
	;
	v313 = F_table_open(m, int32(1255), int32(1))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L9
	} else {
		goto L104
	}
L101:
	;
	F_ScanKeyInit(m, v16+int32(160), int32(10), int32(3), v302, int32(112))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L9
	} else {
		goto L103
	}
L102:
	;
	v302 = int32(61)
	goto L101
L103:
	;
	v310 = int32(2)
	goto L100
L104:
	;
	v317 = F_table_beginscan_catalog(m, v313, v310, v16+int32(112))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L9
	} else {
		goto L105
	}
L105:
	;
	v320 = v266
	goto L106
L106:
	;
	v332 = F_heap_getnext(m, v317)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L9
	} else {
		goto L108
	}
L107:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)+188))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+12))
	m.T0[v342].(func(*base.Module, int32))(m, v317)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L9
	} else {
		goto L113
	}
L108:
	;
	if v332 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v332)+16))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+22)))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v334+v335)))
	v338 = F_lappend_oid(m, v320, v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L9
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	goto L107
L112:
	;
	v320 = v338
	goto L106
L113:
	;
	F_sequence_close(m, v313, int32(1))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L9
	} else {
		goto L114
	}
L114:
	;
	v389 = v320
	goto L92
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v27
	F_errmsg_internal(m, int32(478794), v16+int32(48))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L9
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(492197), int32(866), int32(172007))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L9
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	v366 = F_list_concat(m, v266, v364)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L9
	} else {
		goto L119
	}
L119:
	;
	v369 = F_getRelationsInNamespace(m, v285, int32(118))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L9
	} else {
		goto L120
	}
L120:
	;
	v371 = F_list_concat(m, v366, v369)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L9
	} else {
		goto L121
	}
L121:
	;
	v374 = F_getRelationsInNamespace(m, v285, int32(109))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L9
	} else {
		goto L122
	}
L122:
	;
	v376 = F_list_concat(m, v371, v374)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L9
	} else {
		goto L123
	}
L123:
	;
	v379 = F_getRelationsInNamespace(m, v285, int32(102))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L9
	} else {
		goto L124
	}
L124:
	;
	v381 = F_list_concat(m, v376, v379)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L9
	} else {
		goto L125
	}
L125:
	;
	v384 = F_getRelationsInNamespace(m, v285, int32(112))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L9
	} else {
		goto L126
	}
L126:
	;
	v386 = F_list_concat(m, v381, v384)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L9
	} else {
		goto L127
	}
L127:
	;
	v389 = v386
	goto L92
L128:
	;
	goto L90
L129:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	switch v481 - int32(9) {
	case 0:
		goto L155
	default:
		goto L142
	case 3:
		goto L154
	case 7:
		goto L145
	case 8:
		goto L144
	case 10:
		goto L153
	case 12:
		goto L152
	case 13:
		goto L151
	case 18:
		goto L143
	case 20:
		goto L149
	case 25:
		goto L148
	case 27:
		goto L150
	case 28:
		goto L141
	case 32:
		v528 = int32(260737)
		v529 = int64(-16768)
		goto L140
	case 33:
		goto L147
	case 40:
		goto L146
	}
L130:
	;
	v428 = int32(0)
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v425)+4))
	if v429 <= v428 {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v434 = v428
	v439 = int32(0)
	goto L132
L132:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v425)+12))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v446+v434<<(uint(int32(2))%32))))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v450)+4))
	if v451 != int32(4) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	goto L129
L134:
	;
	v455 = F_get_rolespec_oid(m, v450, int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L9
	} else {
		goto L137
	}
L135:
	;
	v458 = int32(0)
	goto L136
L136:
	;
	v459 = F_lappend_oid(m, v439, v458)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L9
	} else {
		goto L138
	}
L137:
	;
	v458 = v455
	goto L136
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v459
	v463 = v434 + int32(1)
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v425)+4))
	if v463 < v464 {
		v434 = v463
		v439 = v459
		goto L132
	} else {
		goto L139
	}
L139:
	;
	goto L133
L140:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v530 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L141:
	;
	v528 = int32(411544)
	v529 = int64(-263)
	goto L140
L142:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L9
	} else {
		goto L156
	}
L143:
	;
	v528 = int32(214727)
	v529 = int64(-12289)
	goto L140
L144:
	;
	v528 = int32(212637)
	v529 = int64(-257)
	goto L140
L145:
	;
	v528 = int32(216230)
	v529 = int64(-257)
	goto L140
L146:
	;
	v528 = int32(364154)
	v529 = int64(-257)
	goto L140
L147:
	;
	v528 = int32(414340)
	v529 = int64(-513)
	goto L140
L148:
	;
	v528 = int32(369024)
	v529 = int64(-129)
	goto L140
L149:
	;
	v528 = int32(359669)
	v529 = int64(-129)
	goto L140
L150:
	;
	v528 = int32(500613)
	v529 = int64(-769)
	goto L140
L151:
	;
	v528 = int32(110326)
	v529 = int64(-7)
	goto L140
L152:
	;
	v528 = int32(399120)
	v529 = int64(-257)
	goto L140
L153:
	;
	v528 = int32(249656)
	v529 = int64(-129)
	goto L140
L154:
	;
	v528 = int32(274926)
	v529 = int64(-257)
	goto L140
L155:
	;
	v528 = int32(357874)
	v529 = int64(-3585)
	goto L140
L156:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v514
	F_errmsg_internal(m, int32(478794), v16+int32(16))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L9
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(492197), int32(540), int32(96390))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L9
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_ExecGrantStmt_oids(m, v16-int32(-64))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L9
	} else {
		goto L176
	}
L160:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = int64(0)
	v535 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+76)) = uint8(v535)
	goto L159
L161:
	;
	goto L162
L162:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = int64(0)
	v539 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+76)) = uint8(v539)
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v530)+4))
	if v542 <= v539 {
		goto L159
	} else {
		goto L163
	}
L163:
	;
	v547 = int32(0)
	v552 = v539
	goto L164
L164:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v530)+12))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v559+v547<<(uint(int32(2))%32))))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)+8))
	if v564 != 0 {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	goto L159
L166:
	;
	v586 = v547 + int32(1)
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v530)+4))
	if v586 < v587 {
		v547 = v586
		v552 = v583
		goto L164
	} else {
		goto L175
	}
L167:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v565 != int32(41) {
		goto L3
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v563)+4))
	if v571 == int32(0) {
		goto L2
	} else {
		goto L172
	}
L170:
	;
	v568 = F_lappend(m, v552, v563)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L9
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = v568
	v583 = v568
	goto L166
L172:
	;
	v574 = F_string_to_privilege(m, v571)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L9
	} else {
		goto L173
	}
L173:
	;
	if v574&v529 != int64(0) {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	v579 = *(*int64)(unsafe.Add(mBase, uint32(v16)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = v579 | v574
	v583 = v552
	goto L166
L175:
	;
	goto L165
L176:
	;
	m.G0 = v16 + int32(208)
	return
L177:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L9
	} else {
		goto L178
	}
L178:
	;
	F_errmsg(m, int32(215740), int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L9
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(492197), int32(412), int32(96390))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L9
	} else {
		goto L180
	}
L180:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L181:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v629
	F_errmsg_internal(m, int32(478860), v16)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L9
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(492197), int32(434), int32(96390))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L9
	} else {
		goto L183
	}
L183:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L184:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L9
	} else {
		goto L185
	}
L185:
	;
	F_errmsg(m, int32(141732), int32(0))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L9
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(492197), int32(575), int32(96390))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L9
	} else {
		goto L187
	}
L187:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L188:
	;
	F_errmsg_internal(m, int32(146258), int32(0))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L9
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(492197), int32(581), int32(96390))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L9
	} else {
		goto L190
	}
L190:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L191:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L9
	} else {
		goto L192
	}
L192:
	;
	v675 = F_privilege_to_string(m, v574)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L9
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v675
	F_errmsg(m, v528, v16+int32(32))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L9
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(492197), int32(587), int32(96390))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L9
	} else {
		goto L195
	}
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExtendBufferedRelCommon(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int64
	_ = v141
	var v156 int32
	_ = v156
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int64
	_ = v299
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int64
	_ = v307
	var v309 int64
	_ = v309
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int64
	_ = v324
	var v329 int32
	_ = v329
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int64
	_ = v349
	var v350 int64
	_ = v350
	var v354 int64
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int64
	_ = v364
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v376 int64
	_ = v376
	var v377 int64
	_ = v377
	var v381 int64
	_ = v381
	var v388 int32
	_ = v388
	var v390 int64
	_ = v390
	var v392 int64
	_ = v392
	var v395 int32
	_ = v395
	var v397 int64
	_ = v397
	var v400 int32
	_ = v400
	var v402 int64
	_ = v402
	var v433 int32
	_ = v433
	var v434 int64
	_ = v434
	var v438 int32
	_ = v438
	var v456 int32
	_ = v456
	var v457 int64
	_ = v457
	var v461 int32
	_ = v461
	var v478 int32
	_ = v478
	var v479 int64
	_ = v479
	var v484 int32
	_ = v484
	var v485 int64
	_ = v485
	var v490 int32
	_ = v490
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v509 int64
	_ = v509
	var v510 int64
	_ = v510
	var v514 int64
	_ = v514
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v526 int64
	_ = v526
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v538 int64
	_ = v538
	var v539 int64
	_ = v539
	var v543 int64
	_ = v543
	var v550 int32
	_ = v550
	var v552 int64
	_ = v552
	var v554 int64
	_ = v554
	var v557 int32
	_ = v557
	var v559 int64
	_ = v559
	var v562 int32
	_ = v562
	var v564 int64
	_ = v564
	var v595 int32
	_ = v595
	var v596 int64
	_ = v596
	var v600 int32
	_ = v600
	var v618 int32
	_ = v618
	var v619 int64
	_ = v619
	var v623 int32
	_ = v623
	var v640 int32
	_ = v640
	var v641 int64
	_ = v641
	var v646 int32
	_ = v646
	var v647 int64
	_ = v647
	var v652 int32
	_ = v652
	var v660 int32
	_ = v660
	var v669 int32
	_ = v669
	var v677 int32
	_ = v677
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v792 int32
	_ = v792
	var v794 int64
	_ = v794
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v894 int32
	_ = v894
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v915 int32
	_ = v915
	var v929 int32
	_ = v929
	var v934 int32
	_ = v934
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1043 int32
	_ = v1043
	var v1057 int64
	_ = v1057
	var v1062 int32
	_ = v1062
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1082 int32
	_ = v1082
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1254 int32
	_ = v1254
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1294 int32
	_ = v1294
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1360 int32
	_ = v1360
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1390 int32
	_ = v1390
	var v1393 int32
	_ = v1393
	var v1400 int32
	_ = v1400
	var v1402 int64
	_ = v1402
	var v1404 int32
	_ = v1404
	var v1406 int64
	_ = v1406
	var v1413 int32
	_ = v1413
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1442 int32
	_ = v1442
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1477 int64
	_ = v1477
	var v1478 int64
	_ = v1478
	var v1482 int64
	_ = v1482
	var v1487 int32
	_ = v1487
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1498 int64
	_ = v1498
	var v1502 int32
	_ = v1502
	var v1504 int32
	_ = v1504
	var v1510 int64
	_ = v1510
	var v1511 int64
	_ = v1511
	var v1515 int64
	_ = v1515
	var v1522 int32
	_ = v1522
	var v1524 int64
	_ = v1524
	var v1526 int64
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1531 int64
	_ = v1531
	var v1534 int32
	_ = v1534
	var v1536 int64
	_ = v1536
	var v1561 int32
	_ = v1561
	var v1568 int64
	_ = v1568
	var v1572 int32
	_ = v1572
	var v1584 int32
	_ = v1584
	var v1591 int64
	_ = v1591
	var v1595 int32
	_ = v1595
	var v1606 int32
	_ = v1606
	var v1613 int64
	_ = v1613
	var v1619 int64
	_ = v1619
	var v1624 int32
	_ = v1624
	var v1635 int32
	_ = v1635
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1641 int32
	_ = v1641
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1665 int32
	_ = v1665
	var v1670 int32
	_ = v1670
	var v1694 int32
	_ = v1694
	var v1698 int32
	_ = v1698
	var v1701 int32
	_ = v1701
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1718 int32
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1745 int32
	_ = v1745
	var v1747 int64
	_ = v1747
	var v1759 int32
	_ = v1759
	var v1763 int32
	_ = v1763
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1790 int32
	_ = v1790
	var v1799 int32
	_ = v1799
	var v1803 int32
	_ = v1803
	var v1808 int32
	_ = v1808
	var v1812 int32
	_ = v1812
	var v1815 int32
	_ = v1815
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1823 int32
	_ = v1823
	var v1833 int32
	_ = v1833
	var v1838 int32
	_ = v1838
	v24 = m.G0
	v26 = v24 - int32(144)
	m.G0 = v26
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = l4
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v29 == int32(116) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1812 = m.ExcPending
	if v1812 != 0 {
		goto L11
	} else {
		goto L292
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1781 = m.ExcPending
	if v1781 != 0 {
		goto L11
	} else {
		goto L287
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v1759
	m.G0 = v26 + int32(144)
	return v1763
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v32
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = v34
	v38 = m.G0
	v40 = v38 - int32(128)
	m.G0 = v40
	v43 = *(*int32)(unsafe.Add(mBase, _consts[778]))
	if v43 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L6
L6:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v831 = F_IOContextForStrategy(m, l2)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L11
	} else {
		goto L119
	}
L7:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	v1759 = v828
	v1763 = v138
	goto L3
L8:
	;
	F_InitLocalBuffers(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l4) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	return int32(0)
L12:
	;
	goto L10
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[790]))
	v55 = *(*int32)(unsafe.Add(mBase, _consts[777]))
	v56 = v53 - v55
	if base.Ui32(l4) < base.Ui32(v56) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v60 = l4
	goto L15
L15:
	;
	if v60 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v58 = l4
	goto L18
L17:
	;
	v58 = v56
	goto L18
L18:
	;
	v60 = v58
	goto L15
L19:
	;
	v64 = int32(0)
	goto L22
L20:
	;
	goto L21
L21:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v138 = F_smgrnblocks(m, v137, l1)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L11
	} else {
		goto L27
	}
L22:
	;
	v88 = F_GetLocalVictimBuffer(m)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L11
	} else {
		goto L24
	}
L23:
	;
	goto L21
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6+v64<<(uint(int32(2))%32)))) = v88
	v92 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v95 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v95+(v88^int32(-1))<<(uint(int32(6))%32))+20))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v92+(int32(-2)-v101)<<(uint(int32(2))%32))))
	v110 = F__emscripten_memset_bulkmem(m, v106, base.I32_extend8_s(int32(0)), int32(8192))
	mBase = m.M
	goto L25
L25:
	;
	v112 = v64 + int32(1)
	if v112 != v60 {
		v64 = v112
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	v141 = base.I64_extend_i32_u(v60)
	if base.Ui64(base.I64_extend_i32_u(v138)+v141) <= base.Ui64(int64(4294967293)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v60 != 0 {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	goto L30
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L11
	} else {
		goto L114
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(48)))) = v60
	v792 = int32(4374464)
	v794 = *(*int64)(unsafe.Add(mBase, _consts[59]))
	*(*int64)(unsafe.Add(mBase, _consts[59])) = v794 + v141
	m.G0 = v40 + int32(128)
	goto L7
L32:
	;
	v498 = int32(0)
	v500 = int32(*(*uint8)(unsafe.Add(mBase, _consts[791])))
	v503 = m.G0
	v505 = v503 - int32(16)
	m.G0 = v505
	if v500 != 0 {
		goto L86
	} else {
		goto L87
	}
L33:
	;
	v156 = int32(0)
	goto L36
L34:
	;
	goto L35
L35:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, _consts[791])))
	v343 = m.G0
	v345 = v343 - int32(16)
	m.G0 = v345
	if v340 != 0 {
		goto L64
	} else {
		goto L65
	}
L36:
	;
	v171 = l6 + v156<<(uint(int32(2))%32)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	v174 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v176 = *(*int32)(unsafe.Add(mBase, _consts[170]))
	F_ResourceOwnerEnlarge(m, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L11
	} else {
		goto L38
	}
L37:
	;
	goto L32
L38:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v179
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = v181
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v156 + v138
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v40)+28)) = v183
	v189 = v172 ^ int32(-1)
	v192 = v174 + v189<<(uint(int32(6))%32)
	v194 = *(*int32)(unsafe.Add(mBase, _consts[778]))
	v200 = F_hash_search(m, v194, v40+int32(20), int32(1), v40+int32(19))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+19)))
	if v202 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v337 = v156 + int32(1)
	if v60 != v337 {
		v156 = v337
		goto L36
	} else {
		goto L62
	}
L41:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _consts[776]))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v192)+20))
	v209 = int32(-2) - v208
	v212 = v206 + v209<<(uint(int32(2))%32)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v214 = int32(1)
	v215 = v213 - v214
	*(*int32)(unsafe.Add(mBase, uint32(v212))) = v215
	if v215 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v192)+24))
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v40)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v192))) = v307
	v309 = *(*int64)(unsafe.Add(mBase, uint32(v40)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v192)+8)) = v309
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v40)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v192)+16)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v192)+24)) = v306 | int32(33816576)
	*(*int32)(unsafe.Add(mBase, uint32(v200)+20)) = v189
	v318 = v192 + int32(36)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	goto L57
L44:
	;
	v221 = int32(4392088)
	v223 = *(*int32)(unsafe.Add(mBase, _consts[777]))
	v224 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[777])) = v223 - v224
	v228 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v233 = v228 + v209<<(uint(int32(6))%32) + int32(24)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = v234 - v224
	goto L46
L45:
	;
	goto L46
L46:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _consts[170]))
	F_ResourceOwnerForget(m, v240, v208+v214, int32(1606512))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v200)+20))
	v249 = v245 + v246<<(uint(int32(6))%32)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+24))
	v252 = *(*int32)(unsafe.Add(mBase, _consts[776]))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v249)+20))
	v258 = v252 + (int32(-2)-v254)<<(uint(int32(2))%32)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	if v259 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v262 = int32(4392088)
	v264 = *(*int32)(unsafe.Add(mBase, _consts[777]))
	v265 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[777])) = v264 + v265
	*(*int32)(unsafe.Add(mBase, uint32(v249)+24)) = v250 + v265
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v272 = v271
	goto L50
L49:
	;
	v272 = v259
	goto L50
L50:
	;
	v273 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = v272 + v273
	v277 = *(*int32)(unsafe.Add(mBase, _consts[170]))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v249)+20))
	F_ResourceOwnerRemember(m, v277, v278+v273, int32(1606512))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L11
	} else {
		goto L51
	}
L51:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v249)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v284 + int32(1)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v249)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v249)+24)) = v288 & int32(-16777217)
	v293 = v249 + int32(36)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	goto L52
L52:
	;
	if v294 != int32(-1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v293)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+120)) = v297
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v293)))
	*(*int64)(unsafe.Add(mBase, uint32(v40)+112)) = v299
	F_pgaio_wref_wait(m, v40+int32(112))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L11
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	goto L40
L56:
	;
	goto L55
L57:
	;
	if v319 != int32(-1) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+120)) = v322
	v324 = *(*int64)(unsafe.Add(mBase, uint32(v318)))
	*(*int64)(unsafe.Add(mBase, uint32(v40)+112)) = v324
	F_pgaio_wref_wait(m, v40+int32(112))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L11
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	goto L40
L61:
	;
	goto L60
L62:
	;
	goto L37
L63:
	;
	F_smgrzeroextend(m, v137, l1, v138, v60)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L11
	} else {
		goto L67
	}
L64:
	;
	F___clock_gettime(m, int32(1), v345)
	mBase = m.M
	v349 = int64(*(*int32)(unsafe.Add(mBase, uint32(v345)+8)))
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v345)))
	v354 = v349 + v350*int64(1000000000)
	goto L66
L65:
	;
	v354 = int64(0)
	goto L66
L66:
	;
	m.G0 = v345 + int32(16)
	goto L63
L67:
	;
	v360 = int32(1)
	v364 = int64(0)
	v368 = m.G0
	v370 = v368 - int32(16)
	m.G0 = v370
	if v354 != v364 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L31
L69:
	;
	F___clock_gettime(m, int32(1), v370)
	mBase = m.M
	v376 = int64(*(*int32)(unsafe.Add(mBase, uint32(v370)+8)))
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v370)))
	v381 = v376 + (v377*int64(1000000000) - v354)
	goto L73
L70:
	;
	goto L71
L71:
	;
	v478 = int32(4457440)
	v479 = *(*int64)(unsafe.Add(mBase, _consts[792]))
	*(*int64)(unsafe.Add(mBase, _consts[792])) = v479 + base.I64_extend_i32_u(v360)
	v484 = int32(4456480)
	v485 = *(*int64)(unsafe.Add(mBase, _consts[793]))
	*(*int64)(unsafe.Add(mBase, _consts[793])) = v485 + v364
	F_pgstat_count_backend_io_op(m, v360, int32(3), int32(5), v360, v364)
	mBase = m.M
	v490 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[143])) = uint8(v490)
	*(*uint8)(unsafe.Add(mBase, _consts[187])) = uint8(v490)
	m.G0 = v370 + int32(16)
	goto L68
L72:
	;
	v433 = int32(4458400)
	v434 = *(*int64)(unsafe.Add(mBase, _consts[794]))
	*(*int64)(unsafe.Add(mBase, _consts[794])) = v434 + v381
	v438 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	if base.Ui32(int32(16)) < base.Ui32(v438) {
		goto L82
	} else {
		goto L83
	}
L73:
	;
	goto L74
L74:
	;
	v388 = int32(4455872)
	v390 = *(*int64)(unsafe.Add(mBase, _consts[795]))
	v392 = base.I64_div_s(v381, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[795])) = v390 + v392
	switch v360 {
	case 0:
		goto L78
	case 1:
		goto L77
	default:
		goto L72
	}
L77:
	;
	v400 = int32(4374512)
	v402 = *(*int64)(unsafe.Add(mBase, _consts[65]))
	*(*int64)(unsafe.Add(mBase, _consts[65])) = v402 + v381
	goto L72
L78:
	;
	v395 = int32(4374496)
	v397 = *(*int64)(unsafe.Add(mBase, _consts[63]))
	*(*int64)(unsafe.Add(mBase, _consts[63])) = v397 + v381
	goto L72
L82:
	;
	goto L71
L83:
	;
	if int32(1)<<(uint(v438)%32)&int32(115186) == int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v456 = int32(4455296)
	v457 = *(*int64)(unsafe.Add(mBase, _consts[796]))
	*(*int64)(unsafe.Add(mBase, _consts[796])) = v457 + v381
	v461 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[143])) = uint8(v461)
	*(*uint8)(unsafe.Add(mBase, _consts[191])) = uint8(v461)
	goto L82
L85:
	;
	F_smgrzeroextend(m, v137, l1, v138, v60)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L11
	} else {
		goto L89
	}
L86:
	;
	F___clock_gettime(m, int32(1), v505)
	mBase = m.M
	v509 = int64(*(*int32)(unsafe.Add(mBase, uint32(v505)+8)))
	v510 = *(*int64)(unsafe.Add(mBase, uint32(v505)))
	v514 = v509 + v510*int64(1000000000)
	goto L88
L87:
	;
	v514 = int64(0)
	goto L88
L88:
	;
	m.G0 = v505 + int32(16)
	goto L85
L89:
	;
	v520 = int32(1)
	v526 = base.I64_extend_i32_u(v60 << (uint(int32(13)) % 32))
	v530 = m.G0
	v532 = v530 - int32(16)
	m.G0 = v532
	if v514 != int64(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v660 = int32(1)
	if v60 != v660 {
		goto L107
	} else {
		goto L108
	}
L91:
	;
	F___clock_gettime(m, int32(1), v532)
	mBase = m.M
	v538 = int64(*(*int32)(unsafe.Add(mBase, uint32(v532)+8)))
	v539 = *(*int64)(unsafe.Add(mBase, uint32(v532)))
	v543 = v538 + (v539*int64(1000000000) - v514)
	goto L95
L92:
	;
	goto L93
L93:
	;
	v640 = int32(4457440)
	v641 = *(*int64)(unsafe.Add(mBase, _consts[792]))
	*(*int64)(unsafe.Add(mBase, _consts[792])) = v641 + base.I64_extend_i32_u(v520)
	v646 = int32(4456480)
	v647 = *(*int64)(unsafe.Add(mBase, _consts[793]))
	*(*int64)(unsafe.Add(mBase, _consts[793])) = v647 + v526
	F_pgstat_count_backend_io_op(m, v520, int32(3), int32(5), v520, v526)
	mBase = m.M
	v652 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[143])) = uint8(v652)
	*(*uint8)(unsafe.Add(mBase, _consts[187])) = uint8(v652)
	m.G0 = v532 + int32(16)
	goto L90
L94:
	;
	v595 = int32(4458400)
	v596 = *(*int64)(unsafe.Add(mBase, _consts[794]))
	*(*int64)(unsafe.Add(mBase, _consts[794])) = v596 + v543
	v600 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	if base.Ui32(int32(16)) < base.Ui32(v600) {
		goto L104
	} else {
		goto L105
	}
L95:
	;
	goto L96
L96:
	;
	v550 = int32(4455872)
	v552 = *(*int64)(unsafe.Add(mBase, _consts[795]))
	v554 = base.I64_div_s(v543, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[795])) = v552 + v554
	switch v520 {
	case 0:
		goto L100
	case 1:
		goto L99
	default:
		goto L94
	}
L99:
	;
	v562 = int32(4374512)
	v564 = *(*int64)(unsafe.Add(mBase, _consts[65]))
	*(*int64)(unsafe.Add(mBase, _consts[65])) = v564 + v543
	goto L94
L100:
	;
	v557 = int32(4374496)
	v559 = *(*int64)(unsafe.Add(mBase, _consts[63]))
	*(*int64)(unsafe.Add(mBase, _consts[63])) = v559 + v543
	goto L94
L104:
	;
	goto L93
L105:
	;
	if int32(1)<<(uint(v600)%32)&int32(115186) == int32(0) {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v618 = int32(4455296)
	v619 = *(*int64)(unsafe.Add(mBase, _consts[796]))
	*(*int64)(unsafe.Add(mBase, _consts[796])) = v619 + v543
	v623 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[143])) = uint8(v623)
	*(*uint8)(unsafe.Add(mBase, _consts[191])) = uint8(v623)
	goto L104
L107:
	;
	v669 = v498
	v677 = int32(0)
	goto L110
L108:
	;
	v728 = v498
	goto L109
L109:
	;
	if v60&v660 == int32(0) {
		goto L31
	} else {
		goto L113
	}
L110:
	;
	v690 = int32(4392072)
	v691 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v692 = int32(2)
	v694 = l6 + v669<<(uint(v692)%32)
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v694)))
	v696 = int32(-1)
	v698 = int32(6)
	v701 = int32(24)
	v702 = v691 + (v695^v696)<<(uint(v698)%32) + v701
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v702)))
	v704 = int32(16777216)
	*(*int32)(unsafe.Add(mBase, uint32(v702))) = v703 | v704
	v708 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v694)+4))
	v716 = v708 + (v709^v696)<<(uint(v698)%32) + v701
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v716)))
	*(*int32)(unsafe.Add(mBase, uint32(v716))) = v717 | v704
	v722 = v669 + v692
	v724 = v677 + v692
	if v724 != v60&int32(-2) {
		v669 = v722
		v677 = v724
		goto L110
	} else {
		goto L112
	}
L111:
	;
	v728 = v722
	goto L109
L112:
	;
	goto L111
L113:
	;
	v752 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l6+v728<<(uint(int32(2))%32))))
	v763 = v752 + (v756^int32(-1))<<(uint(int32(6))%32) + int32(24)
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v763)))
	*(*int32)(unsafe.Add(mBase, uint32(v763))) = v764 | int32(16777216)
	goto L31
L114:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L11
	} else {
		goto L115
	}
L115:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	F_GetRelationPath(m, v40+int32(40), v809, v810, v811, v812, l1)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L11
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v40 + int32(40)
	F_errmsg(m, int32(152414), v40)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L11
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(492813), int32(395), int32(310912))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L11
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l4) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v929 = l3 & int32(1)
	if v929 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L121:
	;
	v856 = int32(0)
	goto L135
L122:
	;
	v837 = *(*int32)(unsafe.Add(mBase, _consts[797]))
	v839 = *(*int32)(unsafe.Add(mBase, _consts[787]))
	v842 = v837 - v839 - int32(8)
	if base.Ui32(v842) <= base.Ui32(v837) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	if l4 != 0 {
		v854 = int32(1)
		goto L121
	} else {
		goto L134
	}
L125:
	;
	v845 = v842
	goto L127
L126:
	;
	v845 = int32(0)
	goto L127
L127:
	;
	if base.Ui32(v845) <= base.Ui32(int32(1)) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v848 = int32(1)
	goto L130
L129:
	;
	v848 = v845
	goto L130
L130:
	;
	if base.Ui32(v845) < base.Ui32(l4) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v850 = v848
	goto L133
L132:
	;
	v850 = l4
	goto L133
L133:
	;
	v854 = v850
	goto L121
L134:
	;
	v915 = int32(0)
	goto L120
L135:
	;
	v882 = F_GetVictimBuffer(m, l2, v831)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L11
	} else {
		goto L137
	}
L136:
	;
	v915 = v854
	goto L120
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6+v856<<(uint(int32(2))%32)))) = v882
	v886 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v888 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v888+v882<<(uint(int32(6))%32)-int32(44))))
	v901 = F__emscripten_memset_bulkmem(m, v886+v894<<(uint(int32(13))%32), base.I32_extend8_s(int32(0)), int32(8192))
	mBase = m.M
	goto L138
L138:
	;
	v903 = v856 + int32(1)
	if v903 != v854 {
		v856 = v903
		goto L135
	} else {
		goto L139
	}
L139:
	;
	goto L136
L140:
	;
	F_LockRelationForExtension(m, v830, int32(7))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L11
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	if l3&int32(16) != 0 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	goto L142
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v829+l1<<(uint(int32(2))%32))+20)) = int32(-1)
	goto L146
L145:
	;
	goto L146
L146:
	;
	v942 = F_smgrnblocks(m, v829, l1)
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L11
	} else {
		goto L147
	}
L147:
	;
	if l5 == int32(-1) {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	v1057 = base.I64_extend_i32_u(v1043)
	if base.Ui64(int64(4294967293)) < base.Ui64(v1057+base.I64_extend_i32_u(v942)) {
		goto L1
	} else {
		goto L170
	}
L149:
	;
	v1043 = v915
	goto L148
L150:
	;
	goto L151
L151:
	;
	if base.Ui32(v942) <= base.Ui32(l5) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if base.Ui64(base.I64_extend_i32_u(l5)) < base.Ui64(base.I64_extend_i32_u(v942)+base.I64_extend_i32_u(v915)) {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	v954 = int32(0)
	goto L154
L154:
	;
	if base.Ui32(v954) < base.Ui32(v915) {
		goto L158
	} else {
		goto L159
	}
L155:
	;
	v953 = l5 - v942
	goto L157
L156:
	;
	v953 = v915
	goto L157
L157:
	;
	v954 = v953
	goto L154
L158:
	;
	v956 = v954
	goto L161
L159:
	;
	goto L160
L160:
	;
	if v954 != 0 {
		v1043 = v954
		goto L148
	} else {
		goto L167
	}
L161:
	;
	v980 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(l6+v956<<(uint(int32(2))%32))))
	v987 = v980 + v984<<(uint(int32(6))%32)
	v989 = v987 + int32(-64)
	F_StrategyFreeBuffer(m, v989)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L11
	} else {
		goto L163
	}
L162:
	;
	goto L160
L163:
	;
	v993 = *(*int32)(unsafe.Add(mBase, _consts[170]))
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v987-int32(44))))
	F_ResourceOwnerForget(m, v993, v996+int32(1), int32(1606512))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L11
	} else {
		goto L164
	}
L164:
	;
	F_UnpinBufferNoOwner(m, v989)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L11
	} else {
		goto L165
	}
L165:
	;
	v1005 = v956 + int32(1)
	if v1005 != v915 {
		v956 = v1005
		goto L161
	} else {
		goto L166
	}
L166:
	;
	goto L162
L167:
	;
	v1030 = int32(0)
	if v929 != 0 {
		v1759 = v1030
		v1763 = v942
		goto L3
	} else {
		goto L168
	}
L168:
	;
	F_UnlockRelationForExtension(m, v830, int32(7))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L11
	} else {
		goto L169
	}
L169:
	;
	v1759 = v1030
	v1763 = v942
	goto L3
L170:
	;
	if v1043 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v1062 = int32(-2113667072)
	if v29 == int32(112) {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	goto L173
L173:
	;
	v1468 = int32(*(*uint8)(unsafe.Add(mBase, _consts[791])))
	v1471 = m.G0
	v1473 = v1471 - int32(16)
	m.G0 = v1473
	if v1468 != 0 {
		goto L244
	} else {
		goto L245
	}
L174:
	;
	v1067 = v1062
	goto L176
L175:
	;
	v1067 = int32(33816576)
	goto L176
L176:
	;
	if l1 == int32(3) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v1070 = v1062
	goto L179
L178:
	;
	v1070 = v1067
	goto L179
L179:
	;
	v1082 = int32(0)
	goto L180
L180:
	;
	v1097 = l6 + v1082<<(uint(int32(2))%32)
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1097)))
	v1100 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v1102 = *(*int32)(unsafe.Add(mBase, _consts[170]))
	F_ResourceOwnerEnlarge(m, v1102)
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L11
	} else {
		goto L182
	}
L181:
	;
	goto L173
L182:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L11
	} else {
		goto L183
	}
L183:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v829)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v1107
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v829)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v1109
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v829)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v1082 + v942
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v1111
	v1118 = F_BufTableHashCode(m, v26+int32(52))
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L11
	} else {
		goto L184
	}
L184:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v1128 = v1121 + v1118&int32(127)<<(uint(int32(7))%32) + int32(6912)
	v1130 = F_LWLockAcquire(m, v1128, int32(0))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L11
	} else {
		goto L185
	}
L185:
	;
	v1134 = v1100 + v1098<<(uint(int32(6))%32)
	v1136 = v1134 + int32(-64)
	v1140 = v1134 - int32(44)
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1140)))
	v1142 = F_BufTableInsert(m, v26+int32(52), v1118, v1141)
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L11
	} else {
		goto L187
	}
L186:
	;
	v1442 = v1082 + int32(1)
	if v1442 != v1043 {
		v1082 = v1442
		goto L180
	} else {
		goto L242
	}
L187:
	;
	if int32(0) <= v1142 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v1150 = v1147 + v1142<<(uint(int32(6))%32)
	v1151 = F_PinBuffer(m, v1150, l2)
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L11
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+92)) = int32(227047)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+88)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+84)) = int32(489711)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+80)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+72)) = int64(0)
	v1316 = v1134 - int32(40)
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v1316)))
	v1318 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v1316))) = v1317 | v1318
	if v1317&v1318 != 0 {
		goto L222
	} else {
		goto L223
	}
L191:
	;
	F_LWLockRelease(m, v1128)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L11
	} else {
		goto L192
	}
L192:
	;
	F_StrategyFreeBuffer(m, v1136)
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L11
	} else {
		goto L193
	}
L193:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, _consts[170]))
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1140)))
	F_ResourceOwnerForget(m, v1158, v1159+int32(1), int32(1606512))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L11
	} else {
		goto L194
	}
L194:
	;
	F_UnpinBufferNoOwner(m, v1136)
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L11
	} else {
		goto L195
	}
L195:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1150)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1097))) = v1167 + int32(1)
	if v1151 != 0 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1150)+20))
	v1177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1172+v1173<<(uint(int32(13))%32))+14)))
	if v1177 != 0 {
		goto L2
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	goto L200
L199:
	;
	goto L198
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+92)) = int32(227047)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+88)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+84)) = int32(489711)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+80)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+72)) = int64(0)
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1150)+24))
	v1212 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v1150)+24)) = v1211 | v1212
	if v1211&v1212 != 0 {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	goto L186
L202:
	;
	goto L205
L203:
	;
	v1254 = v1211
	goto L204
L204:
	;
	v1276 = int32(4094604)
	v1277 = *(*int32)(unsafe.Add(mBase, _consts[324]))
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(72))+8))
	if v1279 == int32(0) {
		goto L212
	} else {
		goto L213
	}
L205:
	;
	F_perform_spin_delay(m, v26+int32(72))
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L11
	} else {
		goto L207
	}
L206:
	;
	v1254 = v1244
	goto L204
L207:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1150)+24))
	v1245 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v1150)+24)) = v1244 | v1245
	if v1244&v1245 != 0 {
		goto L205
	} else {
		goto L208
	}
L208:
	;
	goto L206
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1150)+24)) = v1254 & int32(-20971521)
	v1301 = F_StartBufferIO(m, v1150, int32(1), int32(0))
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L11
	} else {
		goto L220
	}
L210:
	;
	goto L209
L211:
	;
	*(*int32)(unsafe.Add(mBase, _consts[324])) = v1294
	goto L210
L212:
	;
	if int32(999) < v1277 {
		goto L210
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	if v1277 < int32(11) {
		goto L210
	} else {
		goto L219
	}
L215:
	;
	v1284 = int32(900)
	if v1284 <= v1277 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v1287 = v1284
	goto L218
L217:
	;
	v1287 = v1277
	goto L218
L218:
	;
	v1294 = v1287 + int32(100)
	goto L211
L219:
	;
	v1294 = v1277 - int32(1)
	goto L211
L220:
	;
	if v1301 == int32(0) {
		goto L200
	} else {
		goto L221
	}
L221:
	;
	goto L201
L222:
	;
	goto L225
L223:
	;
	v1360 = v1317
	goto L224
L224:
	;
	v1382 = int32(4094604)
	v1383 = *(*int32)(unsafe.Add(mBase, _consts[324]))
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(72))+8))
	if v1385 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L225:
	;
	F_perform_spin_delay(m, v26+int32(72))
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L11
	} else {
		goto L227
	}
L226:
	;
	v1360 = v1350
	goto L224
L227:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1316)))
	v1351 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v1316))) = v1350 | v1351
	if v1350&v1351 != 0 {
		goto L225
	} else {
		goto L228
	}
L228:
	;
	goto L226
L229:
	;
	v1402 = *(*int64)(unsafe.Add(mBase, uint32(v26)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v1136))) = v1402
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v1136)+16)) = v1404
	v1406 = *(*int64)(unsafe.Add(mBase, uint32(v26)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v1136)+8)) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v1316))) = v1360&int32(-38010881) | v1070
	F_LWLockRelease(m, v1128)
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L11
	} else {
		goto L240
	}
L230:
	;
	goto L229
L231:
	;
	*(*int32)(unsafe.Add(mBase, _consts[324])) = v1400
	goto L230
L232:
	;
	if int32(999) < v1383 {
		goto L230
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	if v1383 < int32(11) {
		goto L230
	} else {
		goto L239
	}
L235:
	;
	v1390 = int32(900)
	if v1390 <= v1383 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v1393 = v1390
	goto L238
L237:
	;
	v1393 = v1383
	goto L238
L238:
	;
	v1400 = v1393 + int32(100)
	goto L231
L239:
	;
	v1400 = v1383 - int32(1)
	goto L231
L240:
	;
	v1416 = F_StartBufferIO(m, v1136, int32(1), int32(0))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L11
	} else {
		goto L241
	}
L241:
	;
	goto L186
L242:
	;
	goto L181
L243:
	;
	F_smgrzeroextend(m, v829, l1, v942, v1043)
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L11
	} else {
		goto L247
	}
L244:
	;
	F___clock_gettime(m, int32(1), v1473)
	mBase = m.M
	v1477 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1473)+8)))
	v1478 = *(*int64)(unsafe.Add(mBase, uint32(v1473)))
	v1482 = v1477 + v1478*int64(1000000000)
	goto L246
L245:
	;
	v1482 = int64(0)
	goto L246
L246:
	;
	m.G0 = v1473 + int32(16)
	goto L243
L247:
	;
	if v929 == int32(0) {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	F_UnlockRelationForExtension(m, v830, int32(7))
	mBase = m.M
	v1492 = m.ExcPending
	if v1492 != 0 {
		goto L11
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	v1493 = int32(0)
	v1495 = int32(1)
	v1498 = base.I64_extend_i32_u(v1043 << (uint(int32(13)) % 32))
	v1502 = m.G0
	v1504 = v1502 - int32(16)
	m.G0 = v1504
	if v1482 != int64(0) {
		goto L253
	} else {
		goto L254
	}
L251:
	;
	goto L250
L252:
	;
	if v1043 == int32(0) {
		goto L269
	} else {
		goto L270
	}
L253:
	;
	F___clock_gettime(m, int32(1), v1504)
	mBase = m.M
	v1510 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1504)+8)))
	v1511 = *(*int64)(unsafe.Add(mBase, uint32(v1504)))
	v1515 = v1510 + (v1511*int64(1000000000) - v1482)
	goto L257
L254:
	;
	goto L255
L255:
	;
	v1606 = v831 << (uint(int32(6)) % 32)
	v1613 = *(*int64)(unsafe.Add(mBase, uint32(v1606)+uint32(_consts[798])))
	*(*int64)(unsafe.Add(mBase, uint32(v1606)+uint32(_consts[798]))) = v1613 + base.I64_extend_i32_u(v1495)
	v1619 = *(*int64)(unsafe.Add(mBase, uint32(v1606)+uint32(_consts[799])))
	*(*int64)(unsafe.Add(mBase, uint32(v1606)+uint32(_consts[799]))) = v1619 + v1498
	F_pgstat_count_backend_io_op(m, v1493, v831, int32(5), v1495, v1498)
	mBase = m.M
	v1624 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[143])) = uint8(v1624)
	*(*uint8)(unsafe.Add(mBase, _consts[187])) = uint8(v1624)
	m.G0 = v1504 + int32(16)
	goto L252
L256:
	;
	v1561 = v831 << (uint(int32(6)) % 32)
	v1568 = *(*int64)(unsafe.Add(mBase, uint32(v1561)+uint32(_consts[800])))
	*(*int64)(unsafe.Add(mBase, uint32(v1561)+uint32(_consts[800]))) = v1568 + v1515
	v1572 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	if base.Ui32(int32(16)) < base.Ui32(v1572) {
		goto L266
	} else {
		goto L267
	}
L257:
	;
	goto L258
L258:
	;
	v1522 = int32(4455872)
	v1524 = *(*int64)(unsafe.Add(mBase, _consts[795]))
	v1526 = base.I64_div_s(v1515, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[795])) = v1524 + v1526
	switch v1493 {
	case 0:
		goto L262
	case 1:
		goto L261
	default:
		goto L256
	}
L261:
	;
	v1534 = int32(4374512)
	v1536 = *(*int64)(unsafe.Add(mBase, _consts[65]))
	*(*int64)(unsafe.Add(mBase, _consts[65])) = v1536 + v1515
	goto L256
L262:
	;
	v1529 = int32(4374496)
	v1531 = *(*int64)(unsafe.Add(mBase, _consts[63]))
	*(*int64)(unsafe.Add(mBase, _consts[63])) = v1531 + v1515
	goto L256
L266:
	;
	goto L255
L267:
	;
	if int32(1)<<(uint(v1572)%32)&int32(115186) == int32(0) {
		goto L266
	} else {
		goto L268
	}
L268:
	;
	v1584 = v831 << (uint(int32(6)) % 32)
	v1591 = *(*int64)(unsafe.Add(mBase, uint32(v1584)+uint32(_consts[801])))
	*(*int64)(unsafe.Add(mBase, uint32(v1584)+uint32(_consts[801]))) = v1591 + v1515
	v1595 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[143])) = uint8(v1595)
	*(*uint8)(unsafe.Add(mBase, _consts[191])) = uint8(v1595)
	goto L266
L269:
	;
	v1745 = int32(4374432)
	v1747 = *(*int64)(unsafe.Add(mBase, _consts[55]))
	*(*int64)(unsafe.Add(mBase, _consts[55])) = v1747 + v1057
	v1759 = v1043
	v1763 = v942
	goto L3
L270:
	;
	v1635 = v942 + int32(1)
	v1637 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v1641 = v1637 + v1638<<(uint(int32(6))%32)
	if l3&int32(8) == int32(0) {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	v1659 = int32(1)
	v1660 = int32(0)
	F_TerminateBufferIO(m, v1641+int32(-64), v1660, int32(16777216), v1659, v1660)
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L11
	} else {
		goto L277
	}
L272:
	;
	if int32(base.Ui32(l3)>>(uint(int32(5))%32))&base.B2i32(l5 == v1635) == int32(0) {
		goto L271
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	v1657 = F_LWLockAcquire(m, v1641-int32(16), int32(0))
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L11
	} else {
		goto L276
	}
L275:
	;
	goto L274
L276:
	;
	goto L271
L277:
	;
	if v1043 == int32(1) {
		goto L269
	} else {
		goto L278
	}
L278:
	;
	v1670 = v1659
	goto L279
L279:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(l6+v1670<<(uint(int32(2))%32))))
	v1701 = v1694 + v1698<<(uint(int32(6))%32)
	if l3&int32(32) == int32(0) {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	goto L269
L281:
	;
	v1713 = int32(0)
	F_TerminateBufferIO(m, v1701+int32(-64), v1713, int32(16777216), int32(1), v1713)
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L11
	} else {
		goto L285
	}
L282:
	;
	if v1670+v1635 != l5 {
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1711 = F_LWLockAcquire(m, v1701-int32(16), int32(0))
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L11
	} else {
		goto L284
	}
L284:
	;
	goto L281
L285:
	;
	v1720 = v1670 + int32(1)
	if v1720 != v1043 {
		v1670 = v1720
		goto L279
	} else {
		goto L286
	}
L286:
	;
	goto L280
L287:
	;
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1150)+16))
	v1785 = *(*int32)(unsafe.Add(mBase, uint32(v829)+4))
	v1786 = *(*int32)(unsafe.Add(mBase, uint32(v829)))
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v829)+8))
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(v829)+12))
	F_GetRelationPath(m, v26+int32(72), v1785, v1786, v1787, v1788, l1)
	mBase = m.M
	v1790 = m.ExcPending
	if v1790 != 0 {
		goto L11
	} else {
		goto L288
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v1782
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v26 + int32(72)
	F_errmsg(m, int32(685973), v26+int32(32))
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L11
	} else {
		goto L289
	}
L289:
	;
	F_errhint(m, int32(599962), int32(0))
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L11
	} else {
		goto L290
	}
L290:
	;
	F_errfinish(m, int32(489711), int32(2777), int32(446640))
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		goto L11
	} else {
		goto L291
	}
L291:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L292:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L11
	} else {
		goto L293
	}
L293:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v829)+4))
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v829)))
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v829)+8))
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v829)+12))
	F_GetRelationPath(m, v26+int32(72), v1818, v1819, v1820, v1821, l1)
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L11
	} else {
		goto L294
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v26 + int32(72)
	F_errmsg(m, int32(152414), v26+int32(16))
	mBase = m.M
	v1833 = m.ExcPending
	if v1833 != 0 {
		goto L11
	} else {
		goto L295
	}
L295:
	;
	F_errfinish(m, int32(489711), int32(2705), int32(446640))
	mBase = m.M
	v1838 = m.ExcPending
	if v1838 != 0 {
		goto L11
	} else {
		goto L296
	}
L296:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__Exit(m *base.Module, l0 int32) {
	m.Wasi_snapshot_preview1.Proc_exit(m, l0)
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__equalMergeAction(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != v5 {
		v31 = v3
		return v31
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v7 != v8 {
			v31 = v3
			return v31
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			if v10 != v11 {
				v31 = v3
				return v31
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				v15 = F_equal(m, v13, v14)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					if v15 == int32(0) {
						v31 = v3
						return v31
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						v23 = F_equal(m, v21, v22)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							if v23 == int32(0) {
								v31 = v3
								return v31
							} else {
								v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								v29 = F_equal(m, v27, v28)
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return int32(0)
								} else {
									v31 = v29
									return v31
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_ec_clear_derived_clauses(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_list_free(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		if v8 != 0 {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
			F_pfree(m, v9)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				F_pfree(m, v8)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
					return
				}
			}
		} else {
			return
		}
	}
}
func F_elem_contained_by_range_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	if v4 == int32(457) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		v13 = F_find_simplified_clause(m, v7, v11, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v18 = v13
			return v18
		}
	} else {
		v18 = int32(0)
		return v18
	}
}
func F_elog_node_display(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v440 int32
	_ = v440
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = int32(4383768)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _consts[583])))
	v23 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[583])) = uint8(v23)
	F_initStringInfo(m, v18)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_outNode(m, v18, l1)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = v21 & int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[583])) = uint8(v31)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	m.G0 = v18 + int32(16)
	if l2 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_pfree(m, v33)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L105
	}
L5:
	;
	v37 = int32(0)
	v39 = m.G0
	v41 = v39 - int32(208)
	m.G0 = v41
	F_initStringInfo(m, v41+int32(112))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v304 = m.G0
	v306 = v304 - int32(128)
	m.G0 = v306
	F_initStringInfo(m, v306+int32(32))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L76
	}
L8:
	;
	v51 = v37
	v53 = v4
	v55 = v37
	goto L9
L9:
	;
	v58 = int32(0)
	if v51 <= v58 {
		v68 = v58
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if int32(0) < v267 {
		goto L71
	} else {
		goto L72
	}
L11:
	;
	v276 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(128)+v267))) = uint8(v276)
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v270))))
	if v279 != 0 {
		goto L67
	} else {
		goto L68
	}
L12:
	;
	v73 = v51
	v74 = v68
	v75 = v53
	v77 = v55
	goto L16
L13:
	;
	v65 = F__emscripten_memset_bulkmem(m, v41+int32(128), base.I32_extend8_s(int32(32)), v51)
	mBase = m.M
	goto L14
L14:
	;
	if base.Ui32(v51) <= base.Ui32(int32(77)) {
		v68 = v51
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v266 = v51
	v267 = v51
	v268 = v53
	v270 = v55
	goto L11
L16:
	;
	v80 = v33 + v77
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v81 == int32(0) {
		v266 = v73
		v267 = v74
		v268 = v75
		v270 = v77
		goto L11
	} else {
		goto L18
	}
L17:
	;
	v266 = v249
	v267 = v259
	v268 = v251
	v270 = v257
	goto L11
L18:
	;
	v86 = v41 + int32(128) + v74
	*(*uint8)(unsafe.Add(mBase, uint32(v86))) = uint8(v81)
	switch v81 - int32(41) {
	case 0:
		goto L22
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16:
		v249 = v73
		v250 = v74
		v251 = v75
		v253 = v77
		goto L19
	case 17:
		goto L20
	default:
		goto L23
	}
L19:
	;
	v256 = int32(1)
	v257 = v253 + v256
	v259 = v250 + v256
	if v259 < int32(78) {
		v73 = v249
		v74 = v259
		v75 = v251
		v77 = v257
		goto L16
	} else {
		goto L66
	}
L20:
	;
	if v73 != v74 {
		goto L62
	} else {
		goto L63
	}
L21:
	;
	if v73 != v74 {
		goto L47
	} else {
		goto L48
	}
L22:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
	if v151 == int32(41) {
		v249 = v73
		v250 = v74
		v251 = v75
		v253 = v77
		goto L19
	} else {
		goto L42
	}
L23:
	;
	switch v81 - int32(123) {
	case 0:
		goto L21
	default:
		v249 = v73
		v250 = v74
		v251 = v75
		v253 = v77
		goto L19
	case 2:
		goto L24
	}
L24:
	;
	if v73 != v74 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v93 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v86))) = uint8(v93)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v41 + int32(128)
	F_appendStringInfo(m, v41+int32(112), int32(726504), v41+int32(48))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v106 = v41 + int32(128)
	v108 = int32(125)
	*(*uint16)(unsafe.Add(mBase, uint32(v106+v73))) = uint16(v108)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v106
	F_appendStringInfo(m, v41+int32(112), int32(726504), v41+int32(32))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	v120 = int32(60)
	v122 = v75 - int32(1)
	v124 = v122 * int32(3)
	if v120 <= v124 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v127 = v120
	goto L32
L31:
	;
	v127 = v124
	goto L32
L32:
	;
	v129 = base.B2i32(int32(0) < v75)
	if int32(0) < v75 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v130 = v127
	goto L35
L34:
	;
	v130 = v73
	goto L35
L35:
	;
	v139 = v77
	goto L36
L36:
	;
	v143 = v139 + int32(1)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v143))))
	if v145 == int32(32) {
		v139 = v143
		goto L36
	} else {
		goto L38
	}
L37:
	;
	if int32(0) < v75 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	v148 = v122
	goto L41
L40:
	;
	v148 = v75
	goto L41
L41:
	;
	v249 = v130
	v250 = v130 - int32(1)
	v251 = v148
	v253 = v139
	goto L19
L42:
	;
	v155 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v41+v74)+129)) = uint8(v155)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v41 + int32(128)
	F_appendStringInfo(m, v41+int32(112), int32(726504), v41-int32(-64))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v175 = v77
	goto L44
L44:
	;
	v179 = v175 + int32(1)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v179))))
	if v181 == int32(32) {
		v175 = v179
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v249 = v73
	v250 = v73 - int32(1)
	v251 = v75
	v253 = v175
	goto L19
L46:
	;
	goto L45
L47:
	;
	v187 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v86))) = uint8(v187)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v41 + int32(128)
	F_appendStringInfo(m, v41+int32(112), int32(726504), v41+int32(80))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v199 = int32(60)
	v201 = v75 + int32(1)
	v203 = v201 * int32(3)
	if v199 <= v203 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L49
L51:
	;
	v206 = v199
	goto L53
L52:
	;
	v206 = v203
	goto L53
L53:
	;
	if v203 <= int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(128)+v219))) = uint8(v223)
	v249 = v206
	v250 = v219
	v251 = v201
	v253 = v77
	goto L19
L55:
	;
	v219 = int32(0)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v213 = int32(1)
	if v206 <= v213 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v216 = v213
	goto L60
L59:
	;
	v216 = v206
	goto L60
L60:
	;
	v218 = F__emscripten_memset_bulkmem(m, v41+int32(128), base.I32_extend8_s(int32(32)), v216)
	mBase = m.M
	goto L61
L61:
	;
	v219 = v216
	goto L54
L62:
	;
	v229 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v86))) = uint8(v229)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+96)) = v41 + int32(128)
	F_appendStringInfo(m, v41+int32(112), int32(726504), v41+int32(96))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	v243 = int32(58)
	goto L64
L64:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(128)+v73))) = uint8(v243)
	v249 = v73
	v250 = v73
	v251 = v75
	v253 = v77
	goto L19
L65:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	v243 = v241
	goto L64
L66:
	;
	goto L17
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v41 + int32(128)
	F_appendStringInfo(m, v41+int32(112), int32(726504), v41+int32(16))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	goto L10
L70:
	;
	v51 = v266
	v53 = v268
	v55 = v270
	goto L9
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v41 + int32(128)
	F_appendStringInfo(m, v41+int32(112), int32(726504), v41)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v41)+112))
	m.G0 = v41 + int32(208)
	v452 = v300
	goto L4
L74:
	;
	goto L73
L75:
	;
	v452 = v407
	goto L4
L76:
	;
	v319 = v4
	v321 = v4
	goto L77
L77:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v321))))
	if v324 != 0 {
		goto L83
	} else {
		goto L84
	}
L79:
	;
	v425 = int32(0)
	v427 = v306 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v427+v418))) = uint8(v425)
	*(*int32)(unsafe.Add(mBase, uint32(v306)+16)) = v427
	F_appendStringInfo(m, v306+int32(32), int32(726504), v306+int32(16))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L104
	}
L80:
	;
	v418 = int32(78)
	v423 = v321 + int32(2)
	goto L79
L81:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v306)+32))
	m.G0 = v306 + int32(128)
	goto L75
L82:
	;
	v392 = v306 + int32(48)
	v394 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v392+v389))) = uint8(v394)
	*(*int32)(unsafe.Add(mBase, uint32(v306))) = v392
	F_appendStringInfo(m, v306+int32(32), int32(726504), v306)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L103
	}
L83:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v306+int32(48)+v319))) = uint8(v324)
	v329 = int32(1)
	v330 = v321 + v329
	v332 = v319 + v329
	if v332 != int32(78) {
		v319 = v332
		v321 = v330
		goto L77
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if v319 == int32(0) {
		goto L81
	} else {
		goto L102
	}
L86:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v330))))
	if v336 == int32(32) {
		goto L80
	} else {
		goto L87
	}
L87:
	;
	v339 = int32(78)
	if v336 == int32(0) {
		v389 = v339
		goto L82
	} else {
		goto L88
	}
L88:
	;
	v349 = v339
	goto L89
L89:
	;
	v354 = v349 - int32(1)
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354+(v306+int32(48))))))
	if v358 == int32(32) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v418 = v381
	v423 = v380 + v321 - int32(77)
	goto L79
L91:
	;
	goto L90
L92:
	;
	v380 = v349
	v381 = v354
	goto L91
L93:
	;
	goto L94
L94:
	;
	v362 = v349 - int32(2)
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362+(v306+int32(48))))))
	if v366 == int32(32) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v380 = v354
	v381 = v362
	goto L91
L96:
	;
	goto L97
L97:
	;
	if base.Ui32(v362) < base.Ui32(int32(2)) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v418 = int32(78)
	v423 = v330
	goto L79
L99:
	;
	goto L100
L100:
	;
	v373 = v349 - int32(3)
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306+int32(48)+v373))))
	if v377 != int32(32) {
		v349 = v373
		goto L89
	} else {
		goto L101
	}
L101:
	;
	v380 = v362
	v381 = v373
	goto L91
L102:
	;
	v389 = v319
	goto L82
L103:
	;
	goto L81
L104:
	;
	v319 = v425
	v321 = v423
	goto L77
L105:
	;
	v457 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	if v457 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l0
	F_errmsg_internal(m, int32(541124), v14+int32(16))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	F_pfree(m, v452)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L113
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v452
	F_errdetail_internal(m, int32(204393), v14)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(487641), int32(85), int32(26358))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	goto L109
L113:
	;
	m.G0 = v14 + int32(32)
	return
}
func F_enlargeStringInfo(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	if int32(0) <= l1 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if base.Ui32(int32(1073741823)-v12) <= base.Ui32(l1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(1073741823)
					F_errmsg(m, int32(652222), v7+int32(32))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v70
						F_errdetail(m, int32(575624), v7+int32(16))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							F_errfinish(m, int32(490729), int32(364), int32(240032))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		} else {
			v17 = l1 + v12 + int32(1)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v18 < v17 {
				v21 = v18
				for {
					v25 = v21 << (uint(int32(1)) % 32)
					if v25 < v17 {
						v21 = v25
						continue
					} else {
						break
					}
					break
				}
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v28 = int32(1073741823)
				if v28 <= v25 {
					v31 = v28
				} else {
					v31 = v25
				}
				v32 = F_repalloc(m, v27, v31)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v31
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v32
					m.G0 = v7 + int32(48)
					return
				}
			} else {
				m.G0 = v7 + int32(48)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
			F_errmsg_internal(m, int32(478190), v7)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				F_errfinish(m, int32(490729), int32(351), int32(240032))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_err_generic_string(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[1086]))
	if int32(0) <= v11 {
		switch l0 - int32(99) {
		case 0:
			v34 = int32(68)
			v36 = v11 * int32(100)
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+uint32(_consts[1097])))
			v43 = F_MemoryContextStrdup(m, v42, l1)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v36+int32(4469072)+v34))) = v43
				m.G0 = v8 + int32(16)
				return
			}
		case 1:
			v34 = int32(72)
			v36 = v11 * int32(100)
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+uint32(_consts[1097])))
			v43 = F_MemoryContextStrdup(m, v42, l1)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v36+int32(4469072)+v34))) = v43
				m.G0 = v8 + int32(16)
				return
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg_internal(m, int32(482200), v8)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					F_errfinish(m, int32(492666), int32(1559), int32(326516))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 11:
			v34 = int32(76)
			v36 = v11 * int32(100)
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+uint32(_consts[1097])))
			v43 = F_MemoryContextStrdup(m, v42, l1)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v36+int32(4469072)+v34))) = v43
				m.G0 = v8 + int32(16)
				return
			}
		case 16:
			v34 = int32(60)
			v36 = v11 * int32(100)
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+uint32(_consts[1097])))
			v43 = F_MemoryContextStrdup(m, v42, l1)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v36+int32(4469072)+v34))) = v43
				m.G0 = v8 + int32(16)
				return
			}
		case 17:
			v34 = int32(64)
			v36 = v11 * int32(100)
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+uint32(_consts[1097])))
			v43 = F_MemoryContextStrdup(m, v42, l1)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v36+int32(4469072)+v34))) = v43
				m.G0 = v8 + int32(16)
				return
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[1086])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(449112), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				F_errfinish(m, int32(492666), int32(1539), int32(326516))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_errstart(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v260 int32
	_ = v260
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	if l0 < int32(21) {
		v131 = l0
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[14]))
		if v14 != 0 {
			v15 = int32(23)
		} else {
			v15 = l0
		}
		if v15 != int32(21) {
			v31 = v15
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[50]))
			if v19 == int32(0) {
				v31 = int32(22)
			} else {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1084])))
				if v23 != 0 {
					v31 = int32(22)
				} else {
					v26 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1085])))
					if v26 != int32(1) {
						v31 = int32(21)
					} else {
						v31 = int32(22)
					}
				}
			}
		}
		v33 = *(*int32)(unsafe.Add(mBase, _consts[1086]))
		if v33 < int32(0) {
			v131 = v31
		} else {
			v36 = int32(1)
			v38 = v33 + v36
			if v38 <= v36 {
				v41 = v36
			} else {
				v41 = v38
			}
			v43 = v41 & int32(3)
			if v38 < int32(4) {
				v99 = v31
				v101 = int32(0)
			} else {
				v49 = int32(0)
				v51 = v31
				v53 = v49
				v54 = v49
				for {
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v53*int32(100))+uint32(_consts[1087])))
					if v64 < v51 {
						v66 = v51
					} else {
						v66 = v64
					}
					v73 = *(*int32)(unsafe.Add(mBase, uint32((v53|int32(1))*int32(100))+uint32(_consts[1087])))
					if v73 < v66 {
						v75 = v66
					} else {
						v75 = v73
					}
					v82 = *(*int32)(unsafe.Add(mBase, uint32((v53|int32(2))*int32(100))+uint32(_consts[1087])))
					if v82 < v75 {
						v84 = v75
					} else {
						v84 = v82
					}
					v91 = *(*int32)(unsafe.Add(mBase, uint32((v53|int32(3))*int32(100))+uint32(_consts[1087])))
					if v91 < v84 {
						v93 = v84
					} else {
						v93 = v91
					}
					v94 = int32(4)
					v95 = v53 + v94
					v97 = v54 + v94
					if v97 != v41&int32(2147483644) {
						v51 = v93
						v53 = v95
						v54 = v97
						continue
					} else {
						break
					}
					break
				}
				v99 = v93
				v101 = v95
			}
			if v43 == int32(0) {
				v131 = v99
			} else {
				v110 = v99
				v112 = v101
				v115 = int32(0)
				for {
					v123 = *(*int32)(unsafe.Add(mBase, uint32(v112*int32(100))+uint32(_consts[1087])))
					if v123 < v110 {
						v125 = v110
					} else {
						v125 = v123
					}
					v126 = int32(1)
					v129 = v115 + v126
					if v129 != v43 {
						v110 = v125
						v112 = v112 + v126
						v115 = v129
						continue
					} else {
						break
					}
					break
				}
				v131 = v125
			}
		}
	}
	v140 = int32(1)
	v142 = *(*int32)(unsafe.Add(mBase, _consts[161]))
	if base.Ui32(v131-int32(15)) <= base.Ui32(v140) {
		if v142 < int32(22) {
			v157 = v140
		} else {
			v157 = int32(0)
		}
	} else {
		if v131 == int32(20) {
			v157 = int32(0)
		} else {
			if v142 == int32(15) {
				if v131 <= int32(21) {
					v157 = int32(0)
				} else {
					v157 = v140
				}
			} else {
				if v131 < v142 {
					v157 = int32(0)
				} else {
					v157 = v140
				}
			}
		}
	}
	v158 = int32(0)
	if v131 == int32(16) {
		v177 = v158
	} else {
		v162 = *(*int32)(unsafe.Add(mBase, _consts[162]))
		if v162 != int32(2) {
			v177 = v158
		} else {
			v166 = int32(*(*uint8)(unsafe.Add(mBase, _consts[163])))
			if v166 == int32(1) {
				v177 = base.B2i32(int32(20) < v131)
			} else {
				v174 = *(*int32)(unsafe.Add(mBase, _consts[164]))
				v177 = base.B2i32(v131 == int32(17)) | base.B2i32(v174 <= v131)
			}
		}
	}
	v183 = (base.B2i32(int32(20) < v131) | v157 | v177) & int32(1)
	if v183 != 0 {
		v185 = *(*int32)(unsafe.Add(mBase, _consts[1088]))
		if v185 == int32(0) {
			F_write_stderr(m, int32(728267), int32(0))
			mBase = m.M
			v275 = m.ExcPending
			if v275 != 0 {
				return int32(0)
			} else {
				F_pgl_exit(m, int32(2))
				mBase = m.M
				v278 = m.ExcPending
				if v278 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		} else {
			v188 = int32(4469068)
			v190 = *(*int32)(unsafe.Add(mBase, _consts[1089]))
			v192 = v190 + int32(1)
			*(*int32)(unsafe.Add(mBase, _consts[1089])) = v192
			if v131 < int32(21) {
				v212 = v192
				v213 = int32(4094804)
				v215 = *(*int32)(unsafe.Add(mBase, _consts[1086]))
				v217 = v215 + int32(1)
				*(*int32)(unsafe.Add(mBase, _consts[1086])) = v217
				if int32(5) <= v217 {
					*(*int32)(unsafe.Add(mBase, _consts[1086])) = int32(-1)
					F_errstart_cold(m, int32(23), int32(0))
					mBase = m.M
					v285 = m.ExcPending
					if v285 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(457583), int32(0))
						mBase = m.M
						v289 = m.ExcPending
						if v289 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492666), int32(762), int32(12100))
							mBase = m.M
							v294 = m.ExcPending
							if v294 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v221 = int32(100)
					v222 = v217 * v221
					v228 = F__emscripten_memset_bulkmem(m, v222+int32(4469072), base.I32_extend8_s(int32(0)), v221)
					mBase = m.M
					v230 = *(*int32)(unsafe.Add(mBase, _consts[40]))
					if l1 != 0 {
						v234 = l1
					} else {
						v234 = int32(543635)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v222)+uint32(_consts[1090]))) = v234
					*(*int32)(unsafe.Add(mBase, uint32(v222)+uint32(_consts[1091]))) = v234
					*(*uint8)(unsafe.Add(mBase, uint32(v222)+uint32(_consts[1092]))) = uint8(v177)
					*(*uint8)(unsafe.Add(mBase, uint32(v222)+uint32(_consts[1093]))) = uint8(v157)
					*(*int32)(unsafe.Add(mBase, uint32(v228))) = v131
					*(*int32)(unsafe.Add(mBase, uint32(v222)+uint32(_consts[1094]))) = v230
					if int32(21) <= v131 {
						*(*int32)(unsafe.Add(mBase, uint32(v228)+28)) = int32(2600)
					} else {
						if int32(19) <= v131 {
							*(*int32)(unsafe.Add(mBase, uint32(v228)+28)) = int32(64)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v228)+28)) = int32(0)
						}
					}
					v260 = *(*int32)(unsafe.Add(mBase, _consts[1088]))
					*(*int32)(unsafe.Add(mBase, uint32(v228)+96)) = v260
					*(*int32)(unsafe.Add(mBase, _consts[1089])) = v212 - int32(1)
					return v183
				}
			} else {
				if v190 <= int32(0) {
					v212 = v192
					v213 = int32(4094804)
					v215 = *(*int32)(unsafe.Add(mBase, _consts[1086]))
					v217 = v215 + int32(1)
					*(*int32)(unsafe.Add(mBase, _consts[1086])) = v217
					if int32(5) <= v217 {
						*(*int32)(unsafe.Add(mBase, _consts[1086])) = int32(-1)
						F_errstart_cold(m, int32(23), int32(0))
						mBase = m.M
						v285 = m.ExcPending
						if v285 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(457583), int32(0))
							mBase = m.M
							v289 = m.ExcPending
							if v289 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(492666), int32(762), int32(12100))
								mBase = m.M
								v294 = m.ExcPending
								if v294 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v221 = int32(100)
						v222 = v217 * v221
						v228 = F__emscripten_memset_bulkmem(m, v222+int32(4469072), base.I32_extend8_s(int32(0)), v221)
						mBase = m.M
						v230 = *(*int32)(unsafe.Add(mBase, _consts[40]))
						if l1 != 0 {
							v234 = l1
						} else {
							v234 = int32(543635)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v222)+uint32(_consts[1090]))) = v234
						*(*int32)(unsafe.Add(mBase, uint32(v222)+uint32(_consts[1091]))) = v234
						*(*uint8)(unsafe.Add(mBase, uint32(v222)+uint32(_consts[1092]))) = uint8(v177)
						*(*uint8)(unsafe.Add(mBase, uint32(v222)+uint32(_consts[1093]))) = uint8(v157)
						*(*int32)(unsafe.Add(mBase, uint32(v228))) = v131
						*(*int32)(unsafe.Add(mBase, uint32(v222)+uint32(_consts[1094]))) = v230
						if int32(21) <= v131 {
							*(*int32)(unsafe.Add(mBase, uint32(v228)+28)) = int32(2600)
						} else {
							if int32(19) <= v131 {
								*(*int32)(unsafe.Add(mBase, uint32(v228)+28)) = int32(64)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v228)+28)) = int32(0)
							}
						}
						v260 = *(*int32)(unsafe.Add(mBase, _consts[1088]))
						*(*int32)(unsafe.Add(mBase, uint32(v228)+96)) = v260
						*(*int32)(unsafe.Add(mBase, _consts[1089])) = v212 - int32(1)
						return v183
					}
				} else {
					F_MemoryContextReset(m, v185)
					mBase = m.M
					v201 = m.ExcPending
					if v201 != 0 {
						return int32(0)
					} else {
						v203 = *(*int32)(unsafe.Add(mBase, _consts[1089]))
						if v203 < int32(3) {
							v212 = v203
						} else {
							v207 = int32(0)
							*(*int32)(unsafe.Add(mBase, _consts[1095])) = v207
							*(*int32)(unsafe.Add(mBase, _consts[49])) = v207
							v212 = v203
						}
						v213 = int32(4094804)
						v215 = *(*int32)(unsafe.Add(mBase, _consts[1086]))
						v217 = v215 + int32(1)
						*(*int32)(unsafe.Add(mBase, _consts[1086])) = v217
						if int32(5) <= v217 {
							*(*int32)(unsafe.Add(mBase, _consts[1086])) = int32(-1)
							F_errstart_cold(m, int32(23), int32(0))
							mBase = m.M
							v285 = m.ExcPending
							if v285 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(457583), int32(0))
								mBase = m.M
								v289 = m.ExcPending
								if v289 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(492666), int32(762), int32(12100))
									mBase = m.M
									v294 = m.ExcPending
									if v294 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v221 = int32(100)
							v222 = v217 * v221
							v228 = F__emscripten_memset_bulkmem(m, v222+int32(4469072), base.I32_extend8_s(int32(0)), v221)
							mBase = m.M
							v230 = *(*int32)(unsafe.Add(mBase, _consts[40]))
							if l1 != 0 {
								v234 = l1
							} else {
								v234 = int32(543635)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v222)+uint32(_consts[1090]))) = v234
							*(*int32)(unsafe.Add(mBase, uint32(v222)+uint32(_consts[1091]))) = v234
							*(*uint8)(unsafe.Add(mBase, uint32(v222)+uint32(_consts[1092]))) = uint8(v177)
							*(*uint8)(unsafe.Add(mBase, uint32(v222)+uint32(_consts[1093]))) = uint8(v157)
							*(*int32)(unsafe.Add(mBase, uint32(v228))) = v131
							*(*int32)(unsafe.Add(mBase, uint32(v222)+uint32(_consts[1094]))) = v230
							if int32(21) <= v131 {
								*(*int32)(unsafe.Add(mBase, uint32(v228)+28)) = int32(2600)
							} else {
								if int32(19) <= v131 {
									*(*int32)(unsafe.Add(mBase, uint32(v228)+28)) = int32(64)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v228)+28)) = int32(0)
								}
							}
							v260 = *(*int32)(unsafe.Add(mBase, _consts[1088]))
							*(*int32)(unsafe.Add(mBase, uint32(v228)+96)) = v260
							*(*int32)(unsafe.Add(mBase, _consts[1089])) = v212 - int32(1)
							return v183
						}
					}
				}
			}
		}
	} else {
		return v183
	}
}
func F_errtable(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+68))
	v5 = F_get_namespace_name(m, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		F_err_generic_string(m, int32(115), v5)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			F_err_generic_string(m, int32(116), v10+int32(4))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_errtablecol(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	if l1 <= int32(0) {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v23 = F_get_attname(m, v20, base.I32_extend16_s(l1), int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v27 = v23
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
			v31 = F_get_namespace_name(m, v30)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				F_err_generic_string(m, int32(115), v31)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					F_err_generic_string(m, int32(116), v36+int32(4))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						F_err_generic_string(m, int32(99), v27)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		if v8 < l1 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v23 = F_get_attname(m, v20, base.I32_extend16_s(l1), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v27 = v23
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
				v31 = F_get_namespace_name(m, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					F_err_generic_string(m, int32(115), v31)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						F_err_generic_string(m, int32(116), v36+int32(4))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_err_generic_string(m, int32(99), v27)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		} else {
			v27 = v7 + v8<<(uint(int32(4))%32) + l1*int32(100) - int32(76)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
			v31 = F_get_namespace_name(m, v30)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				F_err_generic_string(m, int32(115), v31)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					F_err_generic_string(m, int32(116), v36+int32(4))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						F_err_generic_string(m, int32(99), v27)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	}
}
func F_estimate_multivariate_ndistinct(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v199 int32
	_ = v199
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v418 int32
	_ = v418
	var v428 int32
	_ = v428
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v694 int32
	_ = v694
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v715 int32
	_ = v715
	var v722 float64
	_ = v722
	var v725 int32
	_ = v725
	v5 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v33 = int32(0)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	if v34 == v33 {
		v725 = v33
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v32 = v18 + v19<<(uint(int32(2))%32)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v32 = v25 + v26<<(uint(int32(2))%32) - int32(4)
	goto L1
L5:
	;
	return v725
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v37 <= int32(0) {
		v725 = v33
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v48 = v5
	v50 = v5
	v54 = v5
	v55 = v5
	v56 = v5
	goto L8
L8:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v48<<(uint(int32(2))%32))))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+16)))
	if v63 != int32(100) {
		v209 = v50
		v213 = v54
		v214 = v55
		v215 = v56
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v221 = int32(0)
	if v214 == v221 {
		v725 = v221
		goto L5
	} else {
		goto L44
	}
L10:
	;
	v218 = v48 + int32(1)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v218 < v219 {
		v48 = v218
		v50 = v209
		v54 = v213
		v55 = v214
		v56 = v215
		goto L8
	} else {
		goto L43
	}
L11:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+20)))
	if v66 != v67 {
		v209 = v50
		v213 = v54
		v214 = v55
		v215 = v56
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v69 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v180+v188 < int32(2) {
		v209 = v50
		v213 = v54
		v214 = v55
		v215 = v56
		goto L10
	} else {
		goto L37
	}
L14:
	;
	v72 = int32(0)
	v180 = v72
	v188 = v72
	goto L13
L15:
	;
	goto L16
L16:
	;
	v74 = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v77 <= v74 {
		v180 = v74
		v188 = v74
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v84 = v74
	v90 = v74
	v92 = v74
	goto L18
L18:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97+v90<<(uint(int32(2))%32))))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	if v103 == int32(6) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v180 = v159
	v188 = v167
	goto L13
L20:
	;
	v173 = v90 + int32(1)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v173 < v174 {
		v84 = v159
		v90 = v173
		v92 = v167
		goto L18
	} else {
		goto L36
	}
L21:
	;
	v106 = int32(*(*int16)(unsafe.Add(mBase, uint32(v102)+8)))
	if v106 <= int32(0) {
		v159 = v84
		v167 = v92
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v62)+24))
	if v115 == int32(0) {
		v159 = v84
		v167 = v92
		goto L20
	} else {
		goto L27
	}
L24:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
	v110 = F_bms_is_member(m, v106, v109)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int32(0)
L26:
	;
	v159 = v110 + v84
	v167 = v92
	goto L20
L27:
	;
	v118 = int32(0)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v119 <= v118 {
		v159 = v84
		v167 = v92
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v123 = v118
	goto L29
L29:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v140+v123<<(uint(int32(2))%32))))
	v145 = F_equal(m, v139, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L25
	} else {
		goto L31
	}
L30:
	;
	v159 = v84
	v167 = v92 + int32(1)
	goto L20
L31:
	;
	if v145 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v150 = v123 + int32(1)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v150 < v151 {
		v123 = v150
		goto L29
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	goto L30
L35:
	;
	v159 = v84
	v167 = v92
	goto L20
L36:
	;
	goto L19
L37:
	;
	if v188 <= v50 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if v50 != v188 {
		v209 = v50
		v213 = v54
		v214 = v55
		v215 = v56
		goto L10
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v209 = v188
	v213 = v180
	v214 = v199
	v215 = v62
	goto L10
L41:
	;
	if v180 <= v54 {
		v209 = v50
		v213 = v54
		v214 = v55
		v215 = v56
		goto L10
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	goto L9
L44:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+20)))
	v225 = m.G0
	v227 = v225 - int32(32)
	m.G0 = v227
	v230 = F_SearchSysCache2(m, int32(62), v214, v224)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L25
	} else {
		goto L47
	}
L45:
	;
	if v243 == int32(0) {
		v725 = v221
		goto L5
	} else {
		goto L62
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L25
	} else {
		goto L59
	}
L47:
	;
	if v230 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v236 = F_SysCacheGetAttr(m, int32(62), v230, int32(3), v227+int32(31))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L25
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L25
	} else {
		goto L56
	}
L51:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+31)))
	if v238 == int32(1) {
		goto L46
	} else {
		goto L52
	}
L52:
	;
	v241 = F_pg_detoast_datum_packed(m, v236)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L25
	} else {
		goto L53
	}
L53:
	;
	v243 = F_statext_ndistinct_deserialize(m, v241)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L25
	} else {
		goto L54
	}
L54:
	;
	F_ReleaseCatCache(m, v230)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L25
	} else {
		goto L55
	}
L55:
	;
	m.G0 = v227 + int32(32)
	goto L45
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227))) = v214
	F_errmsg_internal(m, int32(41592), v227)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L25
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(487980), int32(158), int32(458648))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L25
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
	*(*int32)(unsafe.Add(mBase, uint32(v227)+20)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v227)+16)) = int32(100)
	F_errmsg_internal(m, int32(41473), v227+int32(16))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L25
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(487980), int32(165), int32(458648))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L25
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	v282 = int32(0)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v215)+24))
	if v284 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	v286 = int32(16)
	v292 = (v285<<(uint(v286)%32) + int32(65536)) >> (uint(v286) % 32)
	goto L65
L64:
	;
	v292 = v282
	goto L65
L65:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v293 == int32(0) {
		v406 = v282
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	if v418 != 0 {
		goto L91
	} else {
		goto L92
	}
L67:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	if v296 <= int32(0) {
		v406 = v282
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v305 = v282
	v307 = int32(0)
	goto L69
L69:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v293)+12))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v317+v307<<(uint(int32(2))%32))))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	if v323 == int32(6) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v406 = v385
	goto L66
L71:
	;
	v398 = v307 + int32(1)
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	if v398 < v399 {
		v305 = v385
		v307 = v398
		goto L69
	} else {
		goto L89
	}
L72:
	;
	v326 = int32(*(*int16)(unsafe.Add(mBase, uint32(v322)+8)))
	if v326 <= int32(0) {
		v385 = v305
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v215)+24))
	if v338 == int32(0) {
		v385 = v305
		goto L71
	} else {
		goto L79
	}
L75:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v215)+20))
	v330 = F_bms_is_member(m, v326, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L25
	} else {
		goto L76
	}
L76:
	;
	if v330 == int32(0) {
		v385 = v305
		goto L71
	} else {
		goto L77
	}
L77:
	;
	v336 = F_bms_add_member(m, v305, base.I32_extend16_s(v326+v292))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L25
	} else {
		goto L78
	}
L78:
	;
	v385 = v336
	goto L71
L79:
	;
	v341 = int32(0)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
	if v342 <= v341 {
		v385 = v305
		goto L71
	} else {
		goto L80
	}
L80:
	;
	v346 = v341
	goto L81
L81:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v363+v346<<(uint(int32(2))%32))))
	v368 = F_equal(m, v362, v367)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L25
	} else {
		goto L83
	}
L82:
	;
	v385 = v305
	goto L71
L83:
	;
	if v368 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v374 = F_bms_add_member(m, v305, base.I32_extend16_s(v292+(v346^int32(-1))))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L25
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v377 = v346 + int32(1)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
	if v377 < v378 {
		v346 = v377
		goto L81
	} else {
		goto L88
	}
L87:
	;
	v385 = v374
	goto L71
L88:
	;
	goto L82
L89:
	;
	goto L70
L90:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v584 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L91:
	;
	v428 = int32(0)
	goto L94
L92:
	;
	goto L93
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L25
	} else {
		goto L121
	}
L94:
	;
	v441 = v243 + int32(16) + v428<<(uint(int32(4))%32)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)+8))
	v443 = int32(0)
	if v406 == v443 {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	goto L93
L96:
	;
	v551 = v428 + int32(1)
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	if base.Ui32(v551) < base.Ui32(v552) {
		v428 = v551
		goto L94
	} else {
		goto L120
	}
L97:
	;
	if v442 != v478 {
		goto L96
	} else {
		goto L110
	}
L98:
	;
	v478 = int32(0)
	goto L97
L99:
	;
	goto L100
L100:
	;
	v450 = int32(1)
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v406)+4))
	if v451 <= v450 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v454 = v450
	goto L103
L102:
	;
	v454 = v451
	goto L103
L103:
	;
	v458 = int32(0)
	v460 = v443
	goto L104
L104:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v406+int32(8)+v458<<(uint(int32(2))%32))))
	if v466 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v478 = v469
	goto L97
L106:
	;
	v469 = v460 + base.I32_popcnt(v466)
	goto L108
L107:
	;
	v469 = v460
	goto L108
L108:
	;
	v471 = v458 + int32(1)
	if v471 != v454 {
		v458 = v471
		v460 = v469
		goto L104
	} else {
		goto L109
	}
L109:
	;
	goto L105
L110:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v441)+8))
	if int32(0) < v480 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v485 = int32(0)
	goto L114
L112:
	;
	goto L113
L113:
	;
	if v441 != 0 {
		goto L90
	} else {
		goto L119
	}
L114:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v441)+12))
	v505 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v501+v485<<(uint(int32(1))%32)))))
	v508 = F_bms_is_member(m, base.I32_extend16_s(v505+v292), v406)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L25
	} else {
		goto L116
	}
L115:
	;
	goto L113
L116:
	;
	if v508 == int32(0) {
		goto L96
	} else {
		goto L117
	}
L117:
	;
	v513 = v485 + int32(1)
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v441)+8))
	if v513 < v514 {
		v485 = v513
		goto L114
	} else {
		goto L118
	}
L118:
	;
	goto L115
L119:
	;
	goto L96
L120:
	;
	goto L95
L121:
	;
	F_errmsg_internal(m, int32(12199), int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L25
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(489369), int32(4472), int32(108005))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L25
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v715
	v722 = *(*float64)(unsafe.Add(mBase, uint32(v441)))
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = v722
	v725 = int32(1)
	goto L5
L125:
	;
	v715 = int32(0)
	goto L124
L126:
	;
	goto L127
L127:
	;
	v588 = int32(0)
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v584)+4))
	if v589 <= v588 {
		v715 = v588
		goto L124
	} else {
		goto L128
	}
L128:
	;
	v600 = int32(0)
	v604 = v588
	goto L129
L129:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v584)+12))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v610+v600<<(uint(int32(2))%32))))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v615)))
	if v616 == int32(6) {
		goto L133
	} else {
		goto L134
	}
L130:
	;
	v715 = v694
	goto L124
L131:
	;
	v701 = v600 + int32(1)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v584)+4))
	if v701 < v702 {
		v600 = v701
		v604 = v694
		goto L129
	} else {
		goto L147
	}
L132:
	;
	v681 = F_lappend(m, v604, v614)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L25
	} else {
		goto L146
	}
L133:
	;
	v619 = int32(*(*int16)(unsafe.Add(mBase, uint32(v615)+8)))
	if v619 <= int32(0) {
		goto L132
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v215)+24))
	if v628 == int32(0) {
		goto L132
	} else {
		goto L139
	}
L136:
	;
	v624 = F_bms_is_member(m, base.I32_extend16_s(v619+v292), v406)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L25
	} else {
		goto L137
	}
L137:
	;
	if v624 == int32(0) {
		goto L132
	} else {
		goto L138
	}
L138:
	;
	v694 = v604
	goto L131
L139:
	;
	v631 = int32(0)
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v628)+4))
	if v632 <= v631 {
		goto L132
	} else {
		goto L140
	}
L140:
	;
	v636 = v631
	goto L141
L141:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v628)+12))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v653+v636<<(uint(int32(2))%32))))
	v658 = F_equal(m, v652, v657)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L25
	} else {
		goto L143
	}
L142:
	;
	goto L132
L143:
	;
	if v658 != 0 {
		v694 = v604
		goto L131
	} else {
		goto L144
	}
L144:
	;
	v661 = v636 + int32(1)
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v628)+4))
	if v661 < v662 {
		v636 = v661
		goto L141
	} else {
		goto L145
	}
L145:
	;
	goto L142
L146:
	;
	v694 = v681
	goto L131
L147:
	;
	goto L130
}
func F_exec_assign_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
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
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int64
	_ = v140
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	v4 = l3
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)) = uint8(v4)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v14 {
	case 0, 4:
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
		v20 = F_exec_cast_value(m, l0, l2, v11+int32(79), l4, l5, v18, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
			if v22 == int32(1) {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
				if v25 != int32(1) {
					v71 = v20
					v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					if v71 != v73 {
						v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
						v82 = v75
						v83 = int32(1)
						v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+14)))
						F_assign_simple_var(m, l0, l1, v71, v82&v83, (v86|v82^int32(-1))&v83)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							m.G0 = v11 + int32(80)
							return
						}
					} else {
						v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
						v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
						if v77 != 0 {
							v82 = v76
							v83 = int32(1)
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+14)))
							F_assign_simple_var(m, l0, l1, v71, v82&v83, (v86|v82^int32(-1))&v83)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								m.G0 = v11 + int32(80)
								return
							}
						} else {
							if v76&int32(1) == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
								m.G0 = v11 + int32(80)
								return
							} else {
								v82 = v76
								v83 = int32(1)
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+14)))
								F_assign_simple_var(m, l0, l1, v71, v82&v83, (v86|v82^int32(-1))&v83)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return
								} else {
									m.G0 = v11 + int32(80)
									return
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(543647))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v37
							F_errmsg(m, int32(526742), v11+int32(16))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								F_errfinish(m, int32(494530), int32(5089), int32(341285))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+14)))
				if v53 != 0 {
					v71 = v20
					v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					if v71 != v73 {
						v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
						v82 = v75
						v83 = int32(1)
						v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+14)))
						F_assign_simple_var(m, l0, l1, v71, v82&v83, (v86|v82^int32(-1))&v83)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							m.G0 = v11 + int32(80)
							return
						}
					} else {
						v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
						v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
						if v77 != 0 {
							v82 = v76
							v83 = int32(1)
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+14)))
							F_assign_simple_var(m, l0, l1, v71, v82&v83, (v86|v82^int32(-1))&v83)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								m.G0 = v11 + int32(80)
								return
							}
						} else {
							if v76&int32(1) == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
								m.G0 = v11 + int32(80)
								return
							} else {
								v82 = v76
								v83 = int32(1)
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+14)))
								F_assign_simple_var(m, l0, l1, v71, v82&v83, (v86|v82^int32(-1))&v83)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return
								} else {
									m.G0 = v11 + int32(80)
									return
								}
							}
						}
					}
				} else {
					v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+20)))
					if v54 != int32(1) {
						v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+12)))
						v69 = F_datumTransfer(m, v20, int32(0), v68)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							v71 = v69
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
							if v71 != v73 {
								v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
								v82 = v75
								v83 = int32(1)
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+14)))
								F_assign_simple_var(m, l0, l1, v71, v82&v83, (v86|v82^int32(-1))&v83)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return
								} else {
									m.G0 = v11 + int32(80)
									return
								}
							} else {
								v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
								v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
								if v77 != 0 {
									v82 = v76
									v83 = int32(1)
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+14)))
									F_assign_simple_var(m, l0, l1, v71, v82&v83, (v86|v82^int32(-1))&v83)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										m.G0 = v11 + int32(80)
										return
									}
								} else {
									if v76&int32(1) == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
										m.G0 = v11 + int32(80)
										return
									} else {
										v82 = v76
										v83 = int32(1)
										v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+14)))
										F_assign_simple_var(m, l0, l1, v71, v82&v83, (v86|v82^int32(-1))&v83)
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											m.G0 = v11 + int32(80)
											return
										}
									}
								}
							}
						}
					} else {
						v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
						if v57 == int32(1) {
							v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
							if v60 == int32(3) {
								v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+12)))
								v69 = F_datumTransfer(m, v20, int32(0), v68)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									v71 = v69
									v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
									if v71 != v73 {
										v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
										v82 = v75
										v83 = int32(1)
										v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+14)))
										F_assign_simple_var(m, l0, l1, v71, v82&v83, (v86|v82^int32(-1))&v83)
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											m.G0 = v11 + int32(80)
											return
										}
									} else {
										v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
										v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
										if v77 != 0 {
											v82 = v76
											v83 = int32(1)
											v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
											v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+14)))
											F_assign_simple_var(m, l0, l1, v71, v82&v83, (v86|v82^int32(-1))&v83)
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												m.G0 = v11 + int32(80)
												return
											}
										} else {
											if v76&int32(1) == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
												m.G0 = v11 + int32(80)
												return
											} else {
												v82 = v76
												v83 = int32(1)
												v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
												v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+14)))
												F_assign_simple_var(m, l0, l1, v71, v82&v83, (v86|v82^int32(-1))&v83)
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return
												} else {
													m.G0 = v11 + int32(80)
													return
												}
											}
										}
									}
								}
							} else {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
								v65 = F_expand_array(m, v20, v63, int32(0))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									v71 = v65
									v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
									if v71 != v73 {
										v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
										v82 = v75
										v83 = int32(1)
										v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+14)))
										F_assign_simple_var(m, l0, l1, v71, v82&v83, (v86|v82^int32(-1))&v83)
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											m.G0 = v11 + int32(80)
											return
										}
									} else {
										v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
										v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
										if v77 != 0 {
											v82 = v76
											v83 = int32(1)
											v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
											v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+14)))
											F_assign_simple_var(m, l0, l1, v71, v82&v83, (v86|v82^int32(-1))&v83)
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												m.G0 = v11 + int32(80)
												return
											}
										} else {
											if v76&int32(1) == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
												m.G0 = v11 + int32(80)
												return
											} else {
												v82 = v76
												v83 = int32(1)
												v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
												v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+14)))
												F_assign_simple_var(m, l0, l1, v71, v82&v83, (v86|v82^int32(-1))&v83)
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return
												} else {
													m.G0 = v11 + int32(80)
													return
												}
											}
										}
									}
								}
							}
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
							v65 = F_expand_array(m, v20, v63, int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								v71 = v65
								v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
								if v71 != v73 {
									v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
									v82 = v75
									v83 = int32(1)
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+14)))
									F_assign_simple_var(m, l0, l1, v71, v82&v83, (v86|v82^int32(-1))&v83)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										m.G0 = v11 + int32(80)
										return
									}
								} else {
									v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
									v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
									if v77 != 0 {
										v82 = v76
										v83 = int32(1)
										v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+14)))
										F_assign_simple_var(m, l0, l1, v71, v82&v83, (v86|v82^int32(-1))&v83)
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											m.G0 = v11 + int32(80)
											return
										}
									} else {
										if v76&int32(1) == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
											m.G0 = v11 + int32(80)
											return
										} else {
											v82 = v76
											v83 = int32(1)
											v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
											v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+14)))
											F_assign_simple_var(m, l0, l1, v71, v82&v83, (v86|v82^int32(-1))&v83)
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												m.G0 = v11 + int32(80)
												return
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	case 1:
		if v4 != 0 {
			v94 = int32(0)
			F_exec_move_row(m, l0, l1, v94, v94)
			mBase = m.M
			v97 = m.ExcPending
			if v97 != 0 {
				return
			} else {
				m.G0 = v11 + int32(80)
				return
			}
		} else {
			v98 = F_type_is_rowtype(m, l4)
			mBase = m.M
			v99 = m.ExcPending
			if v99 != 0 {
				return
			} else {
				if v98 == int32(0) {
					F_errstart_cold(m, int32(21), int32(543647))
					mBase = m.M
					v193 = m.ExcPending
					if v193 != 0 {
						return
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v196 = m.ExcPending
						if v196 != 0 {
							return
						} else {
							F_errmsg(m, int32(392081), int32(0))
							mBase = m.M
							v201 = m.ExcPending
							if v201 != 0 {
								return
							} else {
								F_errfinish(m, int32(494530), int32(5164), int32(341285))
								mBase = m.M
								v208 = m.ExcPending
								if v208 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					F_exec_move_row_from_datum(m, l0, l1, l2)
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return
					} else {
						m.G0 = v11 + int32(80)
						return
					}
				}
			}
		}
	case 2:
		if v4 != 0 {
			v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
			if v104 == int32(1) {
				F_errstart_cold(m, int32(21), int32(543647))
				mBase = m.M
				v214 = m.ExcPending
				if v214 != 0 {
					return
				} else {
					F_errcode(m, int32(67108994))
					mBase = m.M
					v217 = m.ExcPending
					if v217 != 0 {
						return
					} else {
						v218 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v218
						F_errmsg(m, int32(526742), v11+int32(32))
						mBase = m.M
						v225 = m.ExcPending
						if v225 != 0 {
							return
						} else {
							F_errfinish(m, int32(494530), int32(5184), int32(341285))
							mBase = m.M
							v232 = m.ExcPending
							if v232 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v107 = int32(0)
				F_exec_move_row(m, l0, l1, v107, v107)
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return
				} else {
					m.G0 = v11 + int32(80)
					return
				}
			}
		} else {
			v111 = F_type_is_rowtype(m, l4)
			mBase = m.M
			v112 = m.ExcPending
			if v112 != 0 {
				return
			} else {
				if v111 == int32(0) {
					F_errstart_cold(m, int32(21), int32(543647))
					mBase = m.M
					v238 = m.ExcPending
					if v238 != 0 {
						return
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v241 = m.ExcPending
						if v241 != 0 {
							return
						} else {
							F_errmsg(m, int32(392508), int32(0))
							mBase = m.M
							v246 = m.ExcPending
							if v246 != 0 {
								return
							} else {
								F_errfinish(m, int32(494530), int32(5196), int32(341285))
								mBase = m.M
								v253 = m.ExcPending
								if v253 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					F_exec_move_row_from_datum(m, l0, l1, l2)
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return
					} else {
						m.G0 = v11 + int32(80)
						return
					}
				}
			}
		}
	case 3:
		v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v122 = *(*int32)(unsafe.Add(mBase, uint32(v117+v118<<(uint(int32(2))%32))))
		v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+36))
		if v123 == int32(0) {
			F_instantiate_empty_record_variable(m, l0, v122)
			mBase = m.M
			v127 = m.ExcPending
			if v127 != 0 {
				return
			} else {
				v128 = *(*int32)(unsafe.Add(mBase, uint32(v122)+36))
				v129 = v128
				v130 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
				v131 = *(*int64)(unsafe.Add(mBase, uint32(v129)+48))
				if v130 != v131 {
					v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v136 = F_expanded_record_lookup_field(m, v129, v133, l1+int32(32))
					mBase = m.M
					v137 = m.ExcPending
					if v137 != 0 {
						return
					} else {
						if v136 == int32(0) {
							F_errstart_cold(m, int32(21), int32(543647))
							mBase = m.M
							v259 = m.ExcPending
							if v259 != 0 {
								return
							} else {
								F_errcode(m, int32(50360452))
								mBase = m.M
								v262 = m.ExcPending
								if v262 != 0 {
									return
								} else {
									v263 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
									v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v264
									*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v263
									F_errmsg(m, int32(701381), v11-int32(-64))
									mBase = m.M
									v272 = m.ExcPending
									if v272 != 0 {
										return
									} else {
										F_errfinish(m, int32(494530), int32(5239), int32(341285))
										mBase = m.M
										v279 = m.ExcPending
										if v279 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v140 = *(*int64)(unsafe.Add(mBase, uint32(v129)+48))
							*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v140
							v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
							if v142 <= int32(0) {
								F_errstart_cold(m, int32(21), int32(543647))
								mBase = m.M
								v285 = m.ExcPending
								if v285 != 0 {
									return
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v288 = m.ExcPending
									if v288 != 0 {
										return
									} else {
										v289 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v289
										F_errmsg(m, int32(689649), v11+int32(48))
										mBase = m.M
										v296 = m.ExcPending
										if v296 != 0 {
											return
										} else {
											F_errfinish(m, int32(494530), int32(5248), int32(341285))
											mBase = m.M
											v303 = m.ExcPending
											if v303 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
								v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
								v149 = F_exec_cast_value(m, l0, l2, v11+int32(79), l4, l5, v147, v148)
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return
								} else {
									v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
									v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
									v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
									v156 = int32(1)
									F_expanded_record_set_field_internal(m, v129, v151, v149, v152, (v153^int32(-1))&v156, v156)
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										m.G0 = v11 + int32(80)
										return
									}
								}
							}
						}
					}
				} else {
					v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
					if v142 <= int32(0) {
						F_errstart_cold(m, int32(21), int32(543647))
						mBase = m.M
						v285 = m.ExcPending
						if v285 != 0 {
							return
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v288 = m.ExcPending
							if v288 != 0 {
								return
							} else {
								v289 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v289
								F_errmsg(m, int32(689649), v11+int32(48))
								mBase = m.M
								v296 = m.ExcPending
								if v296 != 0 {
									return
								} else {
									F_errfinish(m, int32(494530), int32(5248), int32(341285))
									mBase = m.M
									v303 = m.ExcPending
									if v303 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
						v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						v149 = F_exec_cast_value(m, l0, l2, v11+int32(79), l4, l5, v147, v148)
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return
						} else {
							v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
							v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
							v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
							v156 = int32(1)
							F_expanded_record_set_field_internal(m, v129, v151, v149, v152, (v153^int32(-1))&v156, v156)
							mBase = m.M
							v160 = m.ExcPending
							if v160 != 0 {
								return
							} else {
								m.G0 = v11 + int32(80)
								return
							}
						}
					}
				}
			}
		} else {
			v129 = v123
			v130 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
			v131 = *(*int64)(unsafe.Add(mBase, uint32(v129)+48))
			if v130 != v131 {
				v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v136 = F_expanded_record_lookup_field(m, v129, v133, l1+int32(32))
				mBase = m.M
				v137 = m.ExcPending
				if v137 != 0 {
					return
				} else {
					if v136 == int32(0) {
						F_errstart_cold(m, int32(21), int32(543647))
						mBase = m.M
						v259 = m.ExcPending
						if v259 != 0 {
							return
						} else {
							F_errcode(m, int32(50360452))
							mBase = m.M
							v262 = m.ExcPending
							if v262 != 0 {
								return
							} else {
								v263 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
								v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v264
								*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v263
								F_errmsg(m, int32(701381), v11-int32(-64))
								mBase = m.M
								v272 = m.ExcPending
								if v272 != 0 {
									return
								} else {
									F_errfinish(m, int32(494530), int32(5239), int32(341285))
									mBase = m.M
									v279 = m.ExcPending
									if v279 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v140 = *(*int64)(unsafe.Add(mBase, uint32(v129)+48))
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v140
						v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						if v142 <= int32(0) {
							F_errstart_cold(m, int32(21), int32(543647))
							mBase = m.M
							v285 = m.ExcPending
							if v285 != 0 {
								return
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v288 = m.ExcPending
								if v288 != 0 {
									return
								} else {
									v289 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v289
									F_errmsg(m, int32(689649), v11+int32(48))
									mBase = m.M
									v296 = m.ExcPending
									if v296 != 0 {
										return
									} else {
										F_errfinish(m, int32(494530), int32(5248), int32(341285))
										mBase = m.M
										v303 = m.ExcPending
										if v303 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
							v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
							v149 = F_exec_cast_value(m, l0, l2, v11+int32(79), l4, l5, v147, v148)
							mBase = m.M
							v150 = m.ExcPending
							if v150 != 0 {
								return
							} else {
								v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
								v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
								v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
								v156 = int32(1)
								F_expanded_record_set_field_internal(m, v129, v151, v149, v152, (v153^int32(-1))&v156, v156)
								mBase = m.M
								v160 = m.ExcPending
								if v160 != 0 {
									return
								} else {
									m.G0 = v11 + int32(80)
									return
								}
							}
						}
					}
				}
			} else {
				v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
				if v142 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(543647))
					mBase = m.M
					v285 = m.ExcPending
					if v285 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v288 = m.ExcPending
						if v288 != 0 {
							return
						} else {
							v289 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v289
							F_errmsg(m, int32(689649), v11+int32(48))
							mBase = m.M
							v296 = m.ExcPending
							if v296 != 0 {
								return
							} else {
								F_errfinish(m, int32(494530), int32(5248), int32(341285))
								mBase = m.M
								v303 = m.ExcPending
								if v303 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
					v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					v149 = F_exec_cast_value(m, l0, l2, v11+int32(79), l4, l5, v147, v148)
					mBase = m.M
					v150 = m.ExcPending
					if v150 != 0 {
						return
					} else {
						v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
						v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
						v156 = int32(1)
						F_expanded_record_set_field_internal(m, v129, v151, v149, v152, (v153^int32(-1))&v156, v156)
						mBase = m.M
						v160 = m.ExcPending
						if v160 != 0 {
							return
						} else {
							m.G0 = v11 + int32(80)
							return
						}
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(543647))
		mBase = m.M
		v166 = m.ExcPending
		if v166 != 0 {
			return
		} else {
			v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v167
			F_errmsg_internal(m, int32(478925), v11)
			mBase = m.M
			v172 = m.ExcPending
			if v172 != 0 {
				return
			} else {
				F_errfinish(m, int32(494530), int32(5266), int32(341285))
				mBase = m.M
				v179 = m.ExcPending
				if v179 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_exec_cast_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
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
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	if l3 == l5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(32)
	return v326
L2:
	;
	if l4 == l6 {
		v326 = l1
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l3
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v32 = F_hash_search(m, v26, v16+int32(16), int32(1), v16+int32(15))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if l6 == int32(-1) {
		v326 = l1
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	return int32(0)
L8:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
	if v36 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	if v62 != 0 {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[1253]))
	v46 = F_hash_search(m, v40, v16+int32(16), int32(1), v16+int32(15))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v61 = v60
	goto L9
L13:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
	if v48 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = int32(0)
	goto L16
L15:
	;
	goto L16
L16:
	;
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v53
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+24)) = uint8(v53)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v46
	v61 = v46
	goto L9
L17:
	;
	if v276 == int32(0) {
		v326 = l1
		goto L1
	} else {
		goto L73
	}
L18:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)))
	if v63 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v78 = int32(4476144)
	v79 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v82
	v85 = F_palloc0(m, int32(16))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L7
	} else {
		goto L25
	}
L21:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v276 = v66
	goto L17
L22:
	;
	goto L23
L23:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)+24))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v62)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
	F_MemoryContextDelete(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+20)) = int32(0)
	goto L20
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+8)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(34)
	v91 = F_get_typcollation(m, l3)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+12)) = v91
	if l3 == int32(705) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v125 = m.G0
	v127 = v125 - int32(16)
	m.G0 = v127
	v133 = m.G0
	v135 = v133 - int32(480)
	m.G0 = v135
	v142 = F__emscripten_memset_bulkmem(m, v135+int32(392), base.I32_extend8_s(int32(0)), int32(88))
	mBase = m.M
	goto L36
L28:
	;
	v106 = F_palloc0(m, int32(24))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L7
	} else {
		goto L33
	}
L29:
	;
	if l3 == int32(2249) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v99 = int32(2)
	v102 = F_coerce_to_target_type(m, int32(0), v85, l3, l5, l6, v99, v99, int32(-1))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	if v102 != 0 {
		v124 = v102
		goto L27
	} else {
		goto L32
	}
L32:
	;
	goto L28
L33:
	;
	v108 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v106)+20)) = v108
	*(*int64)(unsafe.Add(mBase, uint32(v106)+12)) = int64(8589934592)
	*(*int32)(unsafe.Add(mBase, uint32(v106)+8)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = int32(28)
	if l6 == v108 {
		v124 = v106
		goto L27
	} else {
		goto L34
	}
L34:
	;
	v122 = F_coerce_to_target_type(m, int32(0), v106, l5, l5, l6, int32(1), int32(2), int32(-1))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	v124 = v122
	goto L27
L36:
	;
	v143 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v135)+448)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v135)+388)) = int32(266)
	v150 = F__emscripten_memset_bulkmem(m, v135, base.I32_extend8_s(v143), int32(384))
	mBase = m.M
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = int32(267)
	*(*int32)(unsafe.Add(mBase, uint32(v150)+8)) = v150 + int32(388)
	v156 = F_eval_const_expressions(m, v150, v124)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	F_fix_opfuncids(m, v156)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	v160 = F_extract_query_dependencies_walker(m, v156, v150)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v150)+444))
	*(*int32)(unsafe.Add(mBase, uint32(v127+int32(12)))) = v162
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v150)+448))
	*(*int32)(unsafe.Add(mBase, uint32(v127+int32(8)))) = v164
	m.G0 = v150 + int32(480)
	v170 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v175 = F_AllocSetContextCreateInternal(m, v170, int32(267906), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	v177 = int32(4476144)
	v178 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v175
	v182 = F_palloc(m, int32(32))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = int32(838275847)
	v186 = F_copyObjectImpl(m, v156)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	v188 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+8)) = uint8(v188)
	*(*int32)(unsafe.Add(mBase, uint32(v182)+4)) = v186
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	v192 = F_copyObjectImpl(m, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+12)) = v192
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v127)+8))
	v196 = F_copyObjectImpl(m, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+20)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v182)+16)) = v196
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v178
	v203 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v175)+16))
	if v207 != v203 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _consts[1254]))
	if v238 != 0 {
		goto L64
	} else {
		goto L65
	}
L47:
	;
	if v207 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	goto L46
L50:
	;
	if v203 != 0 {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v175)+28))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v175)+24))
	if v212 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v211 == int32(0) {
		goto L50
	} else {
		goto L56
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212)+28)) = v211
	goto L52
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207)+20)) = v211
	goto L52
L56:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v175)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v211)+24)) = v217
	goto L50
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v175)+16)) = v203
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v203)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+28)) = v224
	if v224 != 0 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v175)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v175)+16)) = int32(0)
	goto L49
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v224)+24)) = v175
	goto L62
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+20)) = v175
	goto L46
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+24)) = v245
	v247 = int32(4094776)
	*(*int32)(unsafe.Add(mBase, uint32(v182)+28)) = v247
	v250 = v182 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v245)+4)) = v250
	*(*int32)(unsafe.Add(mBase, _consts[1255])) = v250
	m.G0 = v127 + int32(16)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	if v258 == int32(27) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _consts[1255]))
	v245 = v240
	goto L63
L65:
	;
	goto L66
L66:
	;
	v242 = int32(4094776)
	*(*int32)(unsafe.Add(mBase, _consts[1254])) = v242
	v245 = v242
	goto L63
L67:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	if v262 != v85 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v265 = v257
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v61)+20)) = v182
	v268 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+24)) = uint8(v268)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v268
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v79
	v276 = v265
	goto L17
L70:
	;
	v264 = v257
	goto L72
L71:
	;
	v264 = int32(0)
	goto L72
L72:
	;
	v265 = v264
	goto L69
L73:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+56))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	if v287 != v288 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v311)+40)) = l1
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	*(*uint8)(unsafe.Add(mBase, uint32(v311)+44)) = uint8(v315)
	v317 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+24)) = uint8(v317)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v308)+20))
	v320 = m.T0[v319].(func(*base.Module, int32, int32, int32) int32)(m, v308, v311, l2)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L7
	} else {
		goto L81
	}
L75:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+100))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v299
	v302 = F_ExecInitExpr(m, v276, int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L7
	} else {
		goto L80
	}
L76:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v296 = v291
	goto L75
L77:
	;
	goto L78
L78:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+24)))
	if v294 != 0 {
		v296 = v293
		goto L75
	} else {
		goto L79
	}
L79:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v308 = v295
	v309 = v293
	goto L74
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v287
	v305 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+24)) = uint8(v305)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v302
	v308 = v302
	v309 = v296
	goto L74
L81:
	;
	v322 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+24)) = uint8(v322)
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v309
	v326 = v320
	goto L1
}
func F_exec_move_row_from_datum(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v12 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v10 + int32(32)
	return
L2:
	;
	if v265&int32(5) == int32(0) {
		goto L103
	} else {
		goto L104
	}
L3:
	;
	v145 = int32(4476144)
	v146 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v149
	v151 = F_pg_detoast_datum(m, l2)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L10
	} else {
		goto L63
	}
L4:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if v15&int32(254) != int32(2) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+2))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v21 != int32(2) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v264 = int32(0)
	v265 = v24
	goto L2
L7:
	;
	goto L8
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v20 == v26 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_revalidate_rectypeid(m, l1)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v30 != int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v47 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L13:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if v33 != int32(3) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	if v36 != v37 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v36 != int32(2249) {
		goto L12
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_assign_record_var(m, l0, l1, v20)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+28)))
	if v41&int32(64) != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	goto L1
L21:
	;
	v76 = F_make_expanded_record_for_rec(m, l0, l1, v20)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L10
	} else {
		goto L31
	}
L22:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+28)))
	if v50&int32(1) == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	if v55 != v56 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	if v55 == int32(2249) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	if v60 < int32(0) {
		goto L21
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	v67 = int32(1)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	F_expanded_record_set_tuple(m, v47, v66, v67, (v68^int32(-1))&v67)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L10
	} else {
		goto L30
	}
L28:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v47)+40))
	if v60 != v63 {
		goto L21
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	goto L1
L31:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	if v78&int32(1) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v78&int32(5) != 0 {
		v264 = v76
		v265 = v78
		goto L2
	} else {
		goto L60
	}
L33:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v83 != int32(2249) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	if v83 != v86 {
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	v89 = int32(1)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	F_expanded_record_set_tuple(m, v76, v88, v89, (v90^int32(-1))&v89)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L10
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	if v102 != v98 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v132 != 0 {
		goto L56
	} else {
		goto L57
	}
L40:
	;
	if v102 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	goto L39
L43:
	;
	if v98 != 0 {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v97)+28))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
	if v107 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v106 == int32(0) {
		goto L43
	} else {
		goto L49
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+28)) = v106
	goto L45
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102)+20)) = v106
	goto L45
L49:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v106)+24)) = v112
	goto L43
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+16)) = v98
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+28)) = v119
	if v119 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v97)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+16)) = int32(0)
	goto L42
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119)+24)) = v97
	goto L55
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+20)) = v97
	goto L39
L56:
	;
	F_DeleteExpandedObject(m, v132+int32(12))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L10
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v76
	goto L1
L59:
	;
	goto L58
L60:
	;
	F_deconstruct_expanded_record(m, v76)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	F_assign_record_var(m, l0, l1, v76)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	goto L1
L63:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v146
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v151
	v157 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v157
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+20)) = uint16(v157)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(-1)
	v163 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(base.Ui32(v155) >> (uint(v163) % 32))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v168 != v163 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v255 = F_lookup_rowtype_tupdesc(m, v167, v166)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L10
	} else {
		goto L99
	}
L65:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v171 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if base.B2i32(v192 != int32(2249))&base.B2i32(v167 != v192) != 0 {
		goto L64
	} else {
		goto L75
	}
L67:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v171)+36))
	if v167 != v174 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	if v167 == int32(2249) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	if v166 < int32(0) {
		goto L66
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v184 = int32(1)
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	F_expanded_record_set_tuple(m, v171, v10+int32(12), v184, (v185^int32(-1))&v184)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L10
	} else {
		goto L74
	}
L72:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v171)+40))
	if v166 != v180 {
		goto L66
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	goto L1
L75:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+20))
	v199 = F_make_expanded_record_from_typeid(m, v167, v166, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L10
	} else {
		goto L76
	}
L76:
	;
	v203 = int32(1)
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	F_expanded_record_set_tuple(m, v199, v10+int32(12), v203, (v204^int32(-1))&v203)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L10
	} else {
		goto L77
	}
L77:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v199)+8))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v211)+16))
	if v216 != v212 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v246 != 0 {
		goto L95
	} else {
		goto L96
	}
L79:
	;
	if v216 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	goto L78
L82:
	;
	if v212 != 0 {
		goto L89
	} else {
		goto L90
	}
L83:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v211)+28))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v211)+24))
	if v221 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	if v220 == int32(0) {
		goto L82
	} else {
		goto L88
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v221)+28)) = v220
	goto L84
L86:
	;
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+20)) = v220
	goto L84
L88:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v211)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+24)) = v226
	goto L82
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v211)+16)) = v212
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v212)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v211)+28)) = v233
	if v233 != 0 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v211)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v211)+16)) = int32(0)
	goto L81
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v233)+24)) = v211
	goto L94
L93:
	;
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212)+20)) = v211
	goto L78
L95:
	;
	F_DeleteExpandedObject(m, v246+int32(12))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L10
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v199
	goto L1
L98:
	;
	goto L97
L99:
	;
	F_exec_move_row(m, l0, l1, v10+int32(12), v255)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L10
	} else {
		goto L100
	}
L100:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	if v259 < int32(0) {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_DecrTupleDescRefCount(m, v255)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L10
	} else {
		goto L102
	}
L102:
	;
	goto L1
L103:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	if v272 != 0 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	F_deconstruct_expanded_record(m, v20)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L10
	} else {
		goto L111
	}
L106:
	;
	v275 = v272
	goto L108
L107:
	;
	v273 = F_expanded_record_fetch_tupdesc(m, v20)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L10
	} else {
		goto L109
	}
L108:
	;
	F_exec_move_row(m, l0, l1, int32(0), v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L10
	} else {
		goto L110
	}
L109:
	;
	v275 = v273
	goto L108
L110:
	;
	goto L1
L111:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	if v282 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v285 = v282
	goto L114
L113:
	;
	v283 = F_expanded_record_fetch_tupdesc(m, v20)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L10
	} else {
		goto L115
	}
L114:
	;
	F_exec_move_row_from_fields(m, l0, l1, v264, v280, v281, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L10
	} else {
		goto L116
	}
L115:
	;
	v285 = v283
	goto L114
L116:
	;
	goto L1
}
func F_exec_prepare_plan(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int64
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(4781)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v16
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v20 = m.G0
	v22 = v20 - int32(48)
	m.G0 = v22
	v26 = v11 + int32(16)
	if v26 != 0 {
		v27 = v19
	} else {
		v27 = int32(0)
	}
	if v27 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[493])) = int32(-6)
		v86 = int32(0)
		m.G0 = v22 + int32(48)
		if v86 != 0 {
			F_SPI_keepplan(m, v86)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l1)+48)) = int64(0)
				v97 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v97
				*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v86
				v103 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
				if v103 == v97 {
					v151 = v97
				} else {
					v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
					if v106 != int32(1) {
						v151 = v97
					} else {
						v109 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
						v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
						v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+60))
						if v111 == int32(0) {
							v151 = v97
						} else {
							v114 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
							if v114 != int32(1) {
								v151 = v97
							} else {
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
								if v119 != int32(67) {
									v151 = v97
								} else {
									v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
									if v122 != int32(1) {
										v151 = v97
									} else {
										v125 = *(*int32)(unsafe.Add(mBase, uint32(v118)+52))
										if v125 != 0 {
											v151 = v97
										} else {
											v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+36)))
											if v126 != 0 {
												v151 = v97
											} else {
												v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+37)))
												if v127 != 0 {
													v151 = v97
												} else {
													v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+38)))
													if v128 != 0 {
														v151 = v97
													} else {
														v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+39)))
														if v129 != 0 {
															v151 = v97
														} else {
															v130 = *(*int32)(unsafe.Add(mBase, uint32(v118)+48))
															if v130 != 0 {
																v151 = v97
															} else {
																v131 = *(*int32)(unsafe.Add(mBase, uint32(v118)+60))
																v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
																if v132 != 0 {
																	v151 = v97
																} else {
																	v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
																	if v133 != 0 {
																		v151 = v97
																	} else {
																		v134 = *(*int32)(unsafe.Add(mBase, uint32(v118)+100))
																		if v134 != 0 {
																			v151 = v97
																		} else {
																			v135 = *(*int32)(unsafe.Add(mBase, uint32(v118)+108))
																			if v135 != 0 {
																				v151 = v97
																			} else {
																				v136 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
																				if v136 != 0 {
																					v151 = v97
																				} else {
																					v137 = *(*int32)(unsafe.Add(mBase, uint32(v118)+116))
																					if v137 != 0 {
																						v151 = v97
																					} else {
																						v138 = *(*int32)(unsafe.Add(mBase, uint32(v118)+120))
																						if v138 != 0 {
																							v151 = v97
																						} else {
																							v139 = *(*int32)(unsafe.Add(mBase, uint32(v118)+124))
																							if v139 != 0 {
																								v151 = v97
																							} else {
																								v140 = *(*int32)(unsafe.Add(mBase, uint32(v118)+128))
																								if v140 != 0 {
																									v151 = v97
																								} else {
																									v141 = *(*int32)(unsafe.Add(mBase, uint32(v118)+132))
																									if v141 != 0 {
																										v151 = v97
																									} else {
																										v142 = *(*int32)(unsafe.Add(mBase, uint32(v118)+144))
																										if v142 != 0 {
																											v151 = v97
																										} else {
																											v143 = *(*int32)(unsafe.Add(mBase, uint32(v118)+76))
																											if v143 == int32(0) {
																												v151 = v97
																											} else {
																												v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
																												v151 = base.B2i32(v146 == int32(1))
																											}
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
				if v151 != 0 {
					v152 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
					v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
					v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
					v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
					v156 = int32(4476144)
					v157 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
					v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+20))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v160
					v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
					v163 = F_SPI_plan_get_cached_plan(m, v162)
					mBase = m.M
					v164 = m.ExcPending
					if v164 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v157
						v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						v168 = F_CachedPlanAllowsSimpleValidityCheck(m, v155, v163, v167)
						mBase = m.M
						v169 = m.ExcPending
						if v169 != 0 {
							return
						} else {
							if v168 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+60)) = v163
								*(*int32)(unsafe.Add(mBase, uint32(l1)+56)) = v155
								v173 = *(*int32)(unsafe.Add(mBase, _consts[128]))
								v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+56))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v174
								F_exec_save_simple_expr(m, l1, v163)
								mBase = m.M
								v177 = m.ExcPending
								if v177 != 0 {
									return
								} else {
									v179 = *(*int32)(unsafe.Add(mBase, _consts[170]))
									F_ReleaseCachedPlan(m, v163, v179)
									mBase = m.M
									v181 = m.ExcPending
									if v181 != 0 {
										return
									} else {
										m.G0 = v11 + int32(32)
										return
									}
								}
							} else {
								v179 = *(*int32)(unsafe.Add(mBase, _consts[170]))
								F_ReleaseCachedPlan(m, v163, v179)
								mBase = m.M
								v181 = m.ExcPending
								if v181 != 0 {
									return
								} else {
									m.G0 = v11 + int32(32)
									return
								}
							}
						}
					}
				} else {
					m.G0 = v11 + int32(32)
					return
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(543647))
			mBase = m.M
			v191 = m.ExcPending
			if v191 != 0 {
				return
			} else {
				v192 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v194 = *(*int32)(unsafe.Add(mBase, _consts[493]))
				v195 = F_SPI_result_code_string(m, v194)
				mBase = m.M
				v196 = m.ExcPending
				if v196 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v195
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v192
					F_errmsg_internal(m, int32(203679), v11)
					mBase = m.M
					v201 = m.ExcPending
					if v201 != 0 {
						return
					} else {
						F_errfinish(m, int32(494530), int32(4190), int32(280189))
						mBase = m.M
						v206 = m.ExcPending
						if v206 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		v35 = *(*int32)(unsafe.Add(mBase, _consts[529]))
		if v35 != 0 {
			v37 = *(*int32)(unsafe.Add(mBase, _consts[72]))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
			v40 = *(*int32)(unsafe.Add(mBase, _consts[529]))
			*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v38
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v43
			v45 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v45
			*(*int32)(unsafe.Add(mBase, _consts[493])) = v45
			v50 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v22)+12)) = v50
			*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(569278163)
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v54
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
			*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v50
			*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v56
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v60
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v62
			F__SPI_prepare_plan(m, v19, v22+int32(8))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return
			} else {
				v70 = F__SPI_make_plan_non_temp(m, v22+int32(8))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return
				} else {
					v74 = *(*int32)(unsafe.Add(mBase, _consts[529]))
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v75
					*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = int32(0)
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
					F_MemoryContextReset(m, v79)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						v86 = v70
						m.G0 = v22 + int32(48)
						if v86 != 0 {
							F_SPI_keepplan(m, v86)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l1)+48)) = int64(0)
								v97 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v97
								*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v86
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
								if v103 == v97 {
									v151 = v97
								} else {
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
									if v106 != int32(1) {
										v151 = v97
									} else {
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
										v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+60))
										if v111 == int32(0) {
											v151 = v97
										} else {
											v114 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
											if v114 != int32(1) {
												v151 = v97
											} else {
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
												if v119 != int32(67) {
													v151 = v97
												} else {
													v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
													if v122 != int32(1) {
														v151 = v97
													} else {
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v118)+52))
														if v125 != 0 {
															v151 = v97
														} else {
															v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+36)))
															if v126 != 0 {
																v151 = v97
															} else {
																v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+37)))
																if v127 != 0 {
																	v151 = v97
																} else {
																	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+38)))
																	if v128 != 0 {
																		v151 = v97
																	} else {
																		v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+39)))
																		if v129 != 0 {
																			v151 = v97
																		} else {
																			v130 = *(*int32)(unsafe.Add(mBase, uint32(v118)+48))
																			if v130 != 0 {
																				v151 = v97
																			} else {
																				v131 = *(*int32)(unsafe.Add(mBase, uint32(v118)+60))
																				v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
																				if v132 != 0 {
																					v151 = v97
																				} else {
																					v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
																					if v133 != 0 {
																						v151 = v97
																					} else {
																						v134 = *(*int32)(unsafe.Add(mBase, uint32(v118)+100))
																						if v134 != 0 {
																							v151 = v97
																						} else {
																							v135 = *(*int32)(unsafe.Add(mBase, uint32(v118)+108))
																							if v135 != 0 {
																								v151 = v97
																							} else {
																								v136 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
																								if v136 != 0 {
																									v151 = v97
																								} else {
																									v137 = *(*int32)(unsafe.Add(mBase, uint32(v118)+116))
																									if v137 != 0 {
																										v151 = v97
																									} else {
																										v138 = *(*int32)(unsafe.Add(mBase, uint32(v118)+120))
																										if v138 != 0 {
																											v151 = v97
																										} else {
																											v139 = *(*int32)(unsafe.Add(mBase, uint32(v118)+124))
																											if v139 != 0 {
																												v151 = v97
																											} else {
																												v140 = *(*int32)(unsafe.Add(mBase, uint32(v118)+128))
																												if v140 != 0 {
																													v151 = v97
																												} else {
																													v141 = *(*int32)(unsafe.Add(mBase, uint32(v118)+132))
																													if v141 != 0 {
																														v151 = v97
																													} else {
																														v142 = *(*int32)(unsafe.Add(mBase, uint32(v118)+144))
																														if v142 != 0 {
																															v151 = v97
																														} else {
																															v143 = *(*int32)(unsafe.Add(mBase, uint32(v118)+76))
																															if v143 == int32(0) {
																																v151 = v97
																															} else {
																																v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
																																v151 = base.B2i32(v146 == int32(1))
																															}
																														}
																													}
																												}
																											}
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
								if v151 != 0 {
									v152 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
									v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
									v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
									v156 = int32(4476144)
									v157 = *(*int32)(unsafe.Add(mBase, _consts[0]))
									v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
									v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+20))
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v160
									v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									v163 = F_SPI_plan_get_cached_plan(m, v162)
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v157
										v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
										v168 = F_CachedPlanAllowsSimpleValidityCheck(m, v155, v163, v167)
										mBase = m.M
										v169 = m.ExcPending
										if v169 != 0 {
											return
										} else {
											if v168 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(l1)+60)) = v163
												*(*int32)(unsafe.Add(mBase, uint32(l1)+56)) = v155
												v173 = *(*int32)(unsafe.Add(mBase, _consts[128]))
												v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+56))
												*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v174
												F_exec_save_simple_expr(m, l1, v163)
												mBase = m.M
												v177 = m.ExcPending
												if v177 != 0 {
													return
												} else {
													v179 = *(*int32)(unsafe.Add(mBase, _consts[170]))
													F_ReleaseCachedPlan(m, v163, v179)
													mBase = m.M
													v181 = m.ExcPending
													if v181 != 0 {
														return
													} else {
														m.G0 = v11 + int32(32)
														return
													}
												}
											} else {
												v179 = *(*int32)(unsafe.Add(mBase, _consts[170]))
												F_ReleaseCachedPlan(m, v163, v179)
												mBase = m.M
												v181 = m.ExcPending
												if v181 != 0 {
													return
												} else {
													m.G0 = v11 + int32(32)
													return
												}
											}
										}
									}
								} else {
									m.G0 = v11 + int32(32)
									return
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(543647))
							mBase = m.M
							v191 = m.ExcPending
							if v191 != 0 {
								return
							} else {
								v192 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v194 = *(*int32)(unsafe.Add(mBase, _consts[493]))
								v195 = F_SPI_result_code_string(m, v194)
								mBase = m.M
								v196 = m.ExcPending
								if v196 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v195
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v192
									F_errmsg_internal(m, int32(203679), v11)
									mBase = m.M
									v201 = m.ExcPending
									if v201 != 0 {
										return
									} else {
										F_errfinish(m, int32(494530), int32(4190), int32(280189))
										mBase = m.M
										v206 = m.ExcPending
										if v206 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[493])) = int32(-4)
			v86 = int32(0)
			m.G0 = v22 + int32(48)
			if v86 != 0 {
				F_SPI_keepplan(m, v86)
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l1)+48)) = int64(0)
					v97 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v97
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v86
					v103 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
					if v103 == v97 {
						v151 = v97
					} else {
						v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
						if v106 != int32(1) {
							v151 = v97
						} else {
							v109 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
							v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+60))
							if v111 == int32(0) {
								v151 = v97
							} else {
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
								if v114 != int32(1) {
									v151 = v97
								} else {
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
									if v119 != int32(67) {
										v151 = v97
									} else {
										v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
										if v122 != int32(1) {
											v151 = v97
										} else {
											v125 = *(*int32)(unsafe.Add(mBase, uint32(v118)+52))
											if v125 != 0 {
												v151 = v97
											} else {
												v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+36)))
												if v126 != 0 {
													v151 = v97
												} else {
													v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+37)))
													if v127 != 0 {
														v151 = v97
													} else {
														v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+38)))
														if v128 != 0 {
															v151 = v97
														} else {
															v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+39)))
															if v129 != 0 {
																v151 = v97
															} else {
																v130 = *(*int32)(unsafe.Add(mBase, uint32(v118)+48))
																if v130 != 0 {
																	v151 = v97
																} else {
																	v131 = *(*int32)(unsafe.Add(mBase, uint32(v118)+60))
																	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
																	if v132 != 0 {
																		v151 = v97
																	} else {
																		v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
																		if v133 != 0 {
																			v151 = v97
																		} else {
																			v134 = *(*int32)(unsafe.Add(mBase, uint32(v118)+100))
																			if v134 != 0 {
																				v151 = v97
																			} else {
																				v135 = *(*int32)(unsafe.Add(mBase, uint32(v118)+108))
																				if v135 != 0 {
																					v151 = v97
																				} else {
																					v136 = *(*int32)(unsafe.Add(mBase, uint32(v118)+112))
																					if v136 != 0 {
																						v151 = v97
																					} else {
																						v137 = *(*int32)(unsafe.Add(mBase, uint32(v118)+116))
																						if v137 != 0 {
																							v151 = v97
																						} else {
																							v138 = *(*int32)(unsafe.Add(mBase, uint32(v118)+120))
																							if v138 != 0 {
																								v151 = v97
																							} else {
																								v139 = *(*int32)(unsafe.Add(mBase, uint32(v118)+124))
																								if v139 != 0 {
																									v151 = v97
																								} else {
																									v140 = *(*int32)(unsafe.Add(mBase, uint32(v118)+128))
																									if v140 != 0 {
																										v151 = v97
																									} else {
																										v141 = *(*int32)(unsafe.Add(mBase, uint32(v118)+132))
																										if v141 != 0 {
																											v151 = v97
																										} else {
																											v142 = *(*int32)(unsafe.Add(mBase, uint32(v118)+144))
																											if v142 != 0 {
																												v151 = v97
																											} else {
																												v143 = *(*int32)(unsafe.Add(mBase, uint32(v118)+76))
																												if v143 == int32(0) {
																													v151 = v97
																												} else {
																													v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
																													v151 = base.B2i32(v146 == int32(1))
																												}
																											}
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
					if v151 != 0 {
						v152 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
						v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
						v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
						v156 = int32(4476144)
						v157 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+20))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v160
						v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						v163 = F_SPI_plan_get_cached_plan(m, v162)
						mBase = m.M
						v164 = m.ExcPending
						if v164 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v157
							v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							v168 = F_CachedPlanAllowsSimpleValidityCheck(m, v155, v163, v167)
							mBase = m.M
							v169 = m.ExcPending
							if v169 != 0 {
								return
							} else {
								if v168 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+60)) = v163
									*(*int32)(unsafe.Add(mBase, uint32(l1)+56)) = v155
									v173 = *(*int32)(unsafe.Add(mBase, _consts[128]))
									v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+56))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v174
									F_exec_save_simple_expr(m, l1, v163)
									mBase = m.M
									v177 = m.ExcPending
									if v177 != 0 {
										return
									} else {
										v179 = *(*int32)(unsafe.Add(mBase, _consts[170]))
										F_ReleaseCachedPlan(m, v163, v179)
										mBase = m.M
										v181 = m.ExcPending
										if v181 != 0 {
											return
										} else {
											m.G0 = v11 + int32(32)
											return
										}
									}
								} else {
									v179 = *(*int32)(unsafe.Add(mBase, _consts[170]))
									F_ReleaseCachedPlan(m, v163, v179)
									mBase = m.M
									v181 = m.ExcPending
									if v181 != 0 {
										return
									} else {
										m.G0 = v11 + int32(32)
										return
									}
								}
							}
						}
					} else {
						m.G0 = v11 + int32(32)
						return
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(543647))
				mBase = m.M
				v191 = m.ExcPending
				if v191 != 0 {
					return
				} else {
					v192 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v194 = *(*int32)(unsafe.Add(mBase, _consts[493]))
					v195 = F_SPI_result_code_string(m, v194)
					mBase = m.M
					v196 = m.ExcPending
					if v196 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v195
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v192
						F_errmsg_internal(m, int32(203679), v11)
						mBase = m.M
						v201 = m.ExcPending
						if v201 != 0 {
							return
						} else {
							F_errfinish(m, int32(494530), int32(4190), int32(280189))
							mBase = m.M
							v206 = m.ExcPending
							if v206 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_execconsistent(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v11 == int32(0) {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v30 = (v23<<(uint(int32(3))%32) + int32(23)) & int32(-8)
		v31 = v23
		v32 = l1 + v30
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v32
		v36 = F_ArrayGetNItems(m, v31, l1+int32(16))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v32 + v36<<(uint(int32(2))%32)
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v50 = F_execute(m, l0+v42<<(uint(int32(3))%32), v9+int32(8), int32(0), l2, int32(6790))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				m.G0 = v9 + int32(16)
				return v50
			}
		}
	} else {
		v14 = F_array_contains_nulls(m, l1)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v14 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67108994))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(151167), int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(491898), int32(311), int32(90887))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				if v18 == int32(0) {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v30 = (v23<<(uint(int32(3))%32) + int32(23)) & int32(-8)
					v31 = v23
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v30 = v18
					v31 = v21
				}
				v32 = l1 + v30
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v32
				v36 = F_ArrayGetNItems(m, v31, l1+int32(16))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v32 + v36<<(uint(int32(2))%32)
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v50 = F_execute(m, l0+v42<<(uint(int32(3))%32), v9+int32(8), int32(0), l2, int32(6790))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v50
					}
				}
			}
		}
	}
}
func F_executeLikeRegex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v6 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		if v9 == int32(0) {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v14 = F_cstring_to_text_with_len(m, v12, v13)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v14
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v23&int32(1) != 0 {
					v26 = int32(11)
				} else {
					v26 = int32(3)
				}
				if v23&int32(16) != 0 {
					v60 = v26&int32(8) | int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(l3+int32(4)))) = v60
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					v65 = v63
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
					v73 = F_RE_compile_and_cache(m, v65, v69|int32(16), int32(100))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						v79 = F_palloc(m, v68<<(uint(int32(2))%32)+int32(4))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							v81 = F_pg_mb2wchar_with_len(m, v67, v79, v68)
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								v83 = int32(0)
								v86 = F_RE_wchar_execute(m, v79, v81, v83, v83, v83)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v79)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										v95 = v86
										return v95
									}
								}
							}
						}
					}
				} else {
					if v23&int32(8) == int32(0) {
						v60 = v23<<(uint(int32(5))%32)&int32(192) | v26 ^ int32(64)
						*(*int32)(unsafe.Add(mBase, uint32(l3+int32(4)))) = v60
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						v65 = v63
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
						v73 = F_RE_compile_and_cache(m, v65, v69|int32(16), int32(100))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							v79 = F_palloc(m, v68<<(uint(int32(2))%32)+int32(4))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v81 = F_pg_mb2wchar_with_len(m, v67, v79, v68)
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									v83 = int32(0)
									v86 = F_RE_wchar_execute(m, v79, v81, v83, v83, v83)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v79)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											v95 = v86
											return v95
										}
									}
								}
							}
						}
					} else {
						v45 = F_errsave_start(m, int32(0))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							if v45 != 0 {
								F_errcode(m, int32(1088))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(441057), int32(0))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										F_errsave_finish(m, int32(0), int32(26873), int32(680), int32(155450))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											v63 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
											v65 = v63
											v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											v69 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
											v73 = F_RE_compile_and_cache(m, v65, v69|int32(16), int32(100))
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												v79 = F_palloc(m, v68<<(uint(int32(2))%32)+int32(4))
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int32(0)
												} else {
													v81 = F_pg_mb2wchar_with_len(m, v67, v79, v68)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return int32(0)
													} else {
														v83 = int32(0)
														v86 = F_RE_wchar_execute(m, v79, v81, v83, v83, v83)
														mBase = m.M
														v87 = m.ExcPending
														if v87 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v79)
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return int32(0)
															} else {
																v95 = v86
																return v95
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								v65 = v63
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v69 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
								v73 = F_RE_compile_and_cache(m, v65, v69|int32(16), int32(100))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									v79 = F_palloc(m, v68<<(uint(int32(2))%32)+int32(4))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										v81 = F_pg_mb2wchar_with_len(m, v67, v79, v68)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											v83 = int32(0)
											v86 = F_RE_wchar_execute(m, v79, v81, v83, v83, v83)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												F_pfree(m, v79)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													v95 = v86
													return v95
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v65 = v9
			v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v69 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
			v73 = F_RE_compile_and_cache(m, v65, v69|int32(16), int32(100))
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				v79 = F_palloc(m, v68<<(uint(int32(2))%32)+int32(4))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					v81 = F_pg_mb2wchar_with_len(m, v67, v79, v68)
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int32(0)
					} else {
						v83 = int32(0)
						v86 = F_RE_wchar_execute(m, v79, v81, v83, v83, v83)
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v79)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								v95 = v86
								return v95
							}
						}
					}
				}
			}
		}
	} else {
		v95 = int32(2)
		return v95
	}
}
func F_exp(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v40 int64
	_ = v40
	var v56 float64
	_ = v56
	var v60 float64
	_ = v60
	var v62 int32
	_ = v62
	var v64 float64
	_ = v64
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	var v69 float64
	_ = v69
	var v71 float64
	_ = v71
	var v74 float64
	_ = v74
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v81 float64
	_ = v81
	var v84 float64
	_ = v84
	var v88 float64
	_ = v88
	var v91 float64
	_ = v91
	var v94 int64
	_ = v94
	var v99 int32
	_ = v99
	var v102 float64
	_ = v102
	var v105 float64
	_ = v105
	var v108 int64
	_ = v108
	var v111 int64
	_ = v111
	var v120 float64
	_ = v120
	var v127 float64
	_ = v127
	var v128 float64
	_ = v128
	var v129 float64
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 float64
	_ = v138
	var v141 int32
	_ = v141
	var v145 float64
	_ = v145
	var v146 float64
	_ = v146
	var v147 float64
	_ = v147
	var v156 float64
	_ = v156
	var v159 float64
	_ = v159
	var v162 float64
	_ = v162
	var v169 float64
	_ = v169
	var v171 float64
	_ = v171
	var v181 float64
	_ = v181
	v14 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0))>>(uint(int64(52))%64))) & int32(2047)
	v19 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(5.551115123125783e-17))) >> (uint(int64(52)) % 64)))
	if base.Ui32(v14-v19) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(512)))>>(uint(int64(52))%64)))-v19) {
		v62 = v14
		v64 = *(*float64)(unsafe.Add(mBase, _consts[1310]))
		v67 = *(*float64)(unsafe.Add(mBase, _consts[1311]))
		v68 = base.F64_add(base.F64_mul(l0, v64), v67)
		v69 = base.F64_sub(v68, v67)
		v71 = *(*float64)(unsafe.Add(mBase, _consts[1312]))
		v74 = *(*float64)(unsafe.Add(mBase, _consts[1313]))
		v77 = base.F64_add(base.F64_mul(v69, v71), base.F64_add(base.F64_mul(v69, v74), l0))
		v78 = base.F64_mul(v77, v77)
		v81 = *(*float64)(unsafe.Add(mBase, _consts[1314]))
		v84 = *(*float64)(unsafe.Add(mBase, _consts[1315]))
		v88 = *(*float64)(unsafe.Add(mBase, _consts[1316]))
		v91 = *(*float64)(unsafe.Add(mBase, _consts[1317]))
		v94 = base.I64_reinterpret_f64(v68)
		v99 = base.I32_wrap_i64(v94) << (uint(int32(4)) % 32) & int32(2032)
		v102 = *(*float64)(unsafe.Add(mBase, uint32(v99)+uint32(_consts[1318])))
		v105 = base.F64_add(base.F64_mul(base.F64_mul(v78, v78), base.F64_add(base.F64_mul(v77, v81), v84)), base.F64_add(base.F64_mul(v78, base.F64_add(base.F64_mul(v77, v88), v91)), base.F64_add(v102, v77)))
		v108 = *(*int64)(unsafe.Add(mBase, uint32(v99)+uint32(_consts[1319])))
		v111 = v108 + v94<<(uint(int64(45))%64)
		if v62 == int32(0) {
			if v94&int64(2147483648) == int64(0) {
				v120 = base.F64_reinterpret_i64(v111 - int64(4544132024016830464))
				v169 = base.F64_mul(base.F64_add(base.F64_mul(v120, v105), v120), float64(5.486124068793689e+303))
			} else {
				v127 = base.F64_reinterpret_i64(v111 + int64(4602678819172646912))
				v128 = base.F64_mul(v127, v105)
				v129 = base.F64_add(v128, v127)
				if base.F64_lt(v129, float64(1)) != 0 {
					v133 = m.G0
					v135 = v133 - int32(16)
					*(*int64)(unsafe.Add(mBase, uint32(v135)+8)) = int64(4503599627370496)
					v138 = *(*float64)(unsafe.Add(mBase, uint32(v135)+8))
					v141 = m.G0
					*(*float64)(unsafe.Add(mBase, uint32(v141-int32(16))+8)) = base.F64_mul(v138, float64(2.2250738585072014e-308))
					v145 = float64(0)
					v146 = float64(1)
					v147 = base.F64_add(v129, v146)
					v156 = base.F64_add(base.F64_add(v147, base.F64_add(base.F64_add(v128, base.F64_sub(v127, v129)), base.F64_add(v129, base.F64_sub(v146, v147)))), float64(-1))
					if base.F64_eq(v156, v145) != 0 {
						v159 = v145
					} else {
						v159 = v156
					}
					v162 = v159
				} else {
					v162 = v129
				}
				v169 = base.F64_mul(v162, float64(2.2250738585072014e-308))
			}
			return v169
		} else {
			v171 = base.F64_reinterpret_i64(v111)
			v181 = base.F64_add(base.F64_mul(v171, v105), v171)
			return v181
		}
	} else {
		if base.Ui32(v14) < base.Ui32(v19) {
			return base.F64_add(l0, float64(1))
		} else {
			if base.Ui32(v14) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(1024)))>>(uint(int64(52))%64)))) {
				v62 = int32(0)
				v64 = *(*float64)(unsafe.Add(mBase, _consts[1310]))
				v67 = *(*float64)(unsafe.Add(mBase, _consts[1311]))
				v68 = base.F64_add(base.F64_mul(l0, v64), v67)
				v69 = base.F64_sub(v68, v67)
				v71 = *(*float64)(unsafe.Add(mBase, _consts[1312]))
				v74 = *(*float64)(unsafe.Add(mBase, _consts[1313]))
				v77 = base.F64_add(base.F64_mul(v69, v71), base.F64_add(base.F64_mul(v69, v74), l0))
				v78 = base.F64_mul(v77, v77)
				v81 = *(*float64)(unsafe.Add(mBase, _consts[1314]))
				v84 = *(*float64)(unsafe.Add(mBase, _consts[1315]))
				v88 = *(*float64)(unsafe.Add(mBase, _consts[1316]))
				v91 = *(*float64)(unsafe.Add(mBase, _consts[1317]))
				v94 = base.I64_reinterpret_f64(v68)
				v99 = base.I32_wrap_i64(v94) << (uint(int32(4)) % 32) & int32(2032)
				v102 = *(*float64)(unsafe.Add(mBase, uint32(v99)+uint32(_consts[1318])))
				v105 = base.F64_add(base.F64_mul(base.F64_mul(v78, v78), base.F64_add(base.F64_mul(v77, v81), v84)), base.F64_add(base.F64_mul(v78, base.F64_add(base.F64_mul(v77, v88), v91)), base.F64_add(v102, v77)))
				v108 = *(*int64)(unsafe.Add(mBase, uint32(v99)+uint32(_consts[1319])))
				v111 = v108 + v94<<(uint(int64(45))%64)
				if v62 == int32(0) {
					if v94&int64(2147483648) == int64(0) {
						v120 = base.F64_reinterpret_i64(v111 - int64(4544132024016830464))
						v169 = base.F64_mul(base.F64_add(base.F64_mul(v120, v105), v120), float64(5.486124068793689e+303))
					} else {
						v127 = base.F64_reinterpret_i64(v111 + int64(4602678819172646912))
						v128 = base.F64_mul(v127, v105)
						v129 = base.F64_add(v128, v127)
						if base.F64_lt(v129, float64(1)) != 0 {
							v133 = m.G0
							v135 = v133 - int32(16)
							*(*int64)(unsafe.Add(mBase, uint32(v135)+8)) = int64(4503599627370496)
							v138 = *(*float64)(unsafe.Add(mBase, uint32(v135)+8))
							v141 = m.G0
							*(*float64)(unsafe.Add(mBase, uint32(v141-int32(16))+8)) = base.F64_mul(v138, float64(2.2250738585072014e-308))
							v145 = float64(0)
							v146 = float64(1)
							v147 = base.F64_add(v129, v146)
							v156 = base.F64_add(base.F64_add(v147, base.F64_add(base.F64_add(v128, base.F64_sub(v127, v129)), base.F64_add(v129, base.F64_sub(v146, v147)))), float64(-1))
							if base.F64_eq(v156, v145) != 0 {
								v159 = v145
							} else {
								v159 = v156
							}
							v162 = v159
						} else {
							v162 = v129
						}
						v169 = base.F64_mul(v162, float64(2.2250738585072014e-308))
					}
					return v169
				} else {
					v171 = base.F64_reinterpret_i64(v111)
					v181 = base.F64_add(base.F64_mul(v171, v105), v171)
					return v181
				}
			} else {
				v40 = base.I64_reinterpret_f64(l0)
				if v40 == int64(-4503599627370496) {
					v181 = float64(0)
					return v181
				} else {
					if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(math.Float64frombits(uint64(0x7ff0000000000000))))>>(uint(int64(52))%64)))) <= base.Ui32(v14) {
						return base.F64_add(l0, float64(1))
					} else {
						if v40 < int64(0) {
							v56 = F___math_xflow(m, int32(0), float64(1.2882297539194267e-231))
							mBase = m.M
							return v56
						} else {
							v60 = F___math_xflow(m, int32(0), float64(3.105036184601418e+231))
							mBase = m.M
							return v60
						}
					}
				}
			}
		}
	}
}
func F_expand_single_inheritance_child(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v255 int32
	_ = v255
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v435 int32
	_ = v435
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	v9 = int32(0)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l5)+56))
	v31 = F_palloc0(m, int32(136))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(101)
	goto L4
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v29
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+119)))
	v41 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+21)) = uint8(v40)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+28)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+20)) = uint8(base.B2i32(v40 == int32(112)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v28)+52))
	v51 = F_lappend(m, v50, v36)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v36 = F__emscripten_memcpy_bulkmem(m, v31, l1, int32(136))
	mBase = m.M
	goto L6
L6:
	;
	goto L3
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+52)) = v51
	if v51 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v55 = v54
	goto L10
L9:
	;
	v55 = v9
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v36
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v55
	v59 = m.G0
	v61 = v59 - int32(48)
	m.G0 = v61
	v64 = F_palloc0(m, int32(36))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v337 = F_lappend(m, v336, v64)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L65
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(322)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l5)+56))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l5)+52))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = v80
	v84 = F_palloc0(m, v80<<(uint(int32(1))%32))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = v84
	if int32(0) < v78 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L61
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L57
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L54
	}
L17:
	;
	v89 = int32(20)
	v90 = v79 + v89
	v97 = int32(0)
	v103 = int32(0)
	v107 = v9
	goto L20
L18:
	;
	v255 = v9
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = v255
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v270
	m.G0 = v61 + int32(48)
	goto L11
L20:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v127 = v77 + v89 + v121<<(uint(int32(4))%32) + v97*int32(100)
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+91)))
	if v128 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v255 = v234
	goto L19
L22:
	;
	v240 = v97 + int32(1)
	if v240 != v78 {
		v97 = v240
		v103 = v233
		v107 = v234
		goto L20
	} else {
		goto L53
	}
L23:
	;
	v132 = F_lappend(m, v107, int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v127)+96))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v127)+76))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v127)+68))
	if l5 == l3 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v233 = v103
	v234 = v132
	goto L22
L27:
	;
	v139 = v97 + int32(1)
	v142 = F_makeVar(m, v55, base.I32_extend16_s(v139), v136, v135, v134, int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v151 = v127 + int32(4)
	if v80 <= v103 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v144 = F_lappend(m, v107, v142)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v84+v97<<(uint(int32(1))%32)))) = uint16(v139)
	v233 = v103
	v234 = v144
	goto L22
L32:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v210)+68))
	if v136 != v212 {
		goto L15
	} else {
		goto L48
	}
L33:
	;
	v191 = F_SearchSysCacheAttName(m, v76, v151)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L45
	}
L34:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v159 = v90 + v153<<(uint(int32(4))%32) + v103*int32(100)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+91)))
	if v160 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v162 = v159 + int32(4)
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v166 == int32(0) {
		v185 = v165
		v186 = v166
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v186-v185 == int32(0) {
		v210 = v159
		v211 = v103
		goto L32
	} else {
		goto L44
	}
L37:
	;
	goto L36
L38:
	;
	if v165 != v166 {
		v185 = v165
		v186 = v166
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v170 = v151
	v171 = v162
	goto L40
L40:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+1)))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+1)))
	if v175 == int32(0) {
		v185 = v174
		v186 = v175
		goto L37
	} else {
		goto L42
	}
L41:
	;
	v185 = v174
	v186 = v175
	goto L37
L42:
	;
	v178 = int32(1)
	if v174 == v175 {
		v170 = v170 + v178
		v171 = v171 + v178
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	goto L33
L45:
	;
	if v191 == int32(0) {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191)+16))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+22)))
	v198 = int32(*(*int16)(unsafe.Add(mBase, uint32(v195+v196)+74)))
	F_ReleaseCatCache(m, v191)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v206 = v198 - int32(1)
	v210 = v90 + v201<<(uint(int32(4))%32) + v206*int32(100)
	v211 = v206
	goto L32
L48:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210)+76))
	if v135 != v214 {
		goto L15
	} else {
		goto L49
	}
L49:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v210)+96))
	if v134 != v216 {
		goto L14
	} else {
		goto L50
	}
L50:
	;
	v218 = int32(1)
	v221 = v211 + v218
	v224 = F_makeVar(m, v55, base.I32_extend16_s(v221), v136, v135, v134, int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v226 = F_lappend(m, v107, v224)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v230 = v97 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v211<<(uint(v218)%32)+v84))) = uint16(v230)
	v233 = v221
	v234 = v226
	goto L22
L53:
	;
	goto L21
L54:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v279 + int32(4)
	F_errmsg_internal(m, int32(687081), v61)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(490742), int32(153), int32(73360))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errcode(m, int32(17064068))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+32)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v61)+36)) = v299 + int32(4)
	F_errmsg(m, int32(363975), v61+int32(32))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(490742), int32(166), int32(73360))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
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
	F_errcode(m, int32(17064068))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v61)+20)) = v321 + int32(4)
	F_errmsg(m, int32(260094), v61+int32(16))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(490742), int32(171), int32(73360))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v337
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v341 = F_copyObjectImpl(m, v340)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v341
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l5)+52))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	if int32(0) < v345 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v350)+8))
	v360 = v345
	v361 = int32(0)
	v367 = v41
	goto L70
L68:
	;
	v435 = v41
	goto L69
L69:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	v450 = F_makeAlias(m, v449, v435)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L84
	}
L70:
	;
	v386 = v344 + int32(20) + v360<<(uint(int32(4))%32) + v361*int32(100)
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386)+91)))
	if v387 != 0 {
		v410 = int32(735586)
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v435 = v415
	goto L69
L72:
	;
	v411 = F_pstrdup(m, v410)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L80
	}
L73:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	v392 = int32(*(*int16)(unsafe.Add(mBase, uint32(v388+v361<<(uint(int32(1))%32)))))
	if v392 <= int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v410 = v386 + int32(4)
	goto L72
L75:
	;
	if v351 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v351)+4))
	v397 = v395
	goto L78
L77:
	;
	v397 = int32(0)
	goto L78
L78:
	;
	if v397 < v392 {
		goto L74
	} else {
		goto L79
	}
L79:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v399+v392<<(uint(int32(2))%32)-int32(4))))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	v410 = v406
	goto L72
L80:
	;
	v413 = F_makeString(m, v411)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v415 = F_lappend(m, v367, v413)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v418 = v361 + int32(1)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	if v418 < v419 {
		v360 = v419
		v361 = v418
		v367 = v415
		goto L70
	} else {
		goto L83
	}
L83:
	;
	goto L71
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v450
	v455 = v55 << (uint(int32(2)) % 32)
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v455+v456))) = v36
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v459+v455))) = v64
	if l4 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v463 = F_palloc0(m, int32(36))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v497 = F_bms_is_member(m, l2, v496)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L92
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v463)+4)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v463))) = int32(374)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v463)+8)) = v468
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v463)+12)) = v470
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
	v473 = F_select_rowmark_type(m, v36, v472)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v463)+16)) = v473
	v477 = int32(1) << (uint(v473) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v463)+20)) = v477
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v463)+24)) = v479
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v463)+28)) = v481
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v463)+32)) = uint8(base.B2i32(v483 == int32(112)))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v487 | v477
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v491 = F_lappend(m, v490, v463)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v491
	goto L87
L91:
	;
	return
L92:
	;
	if v497 == int32(0) {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v502 = F_bms_add_member(m, v501, v55)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v502
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+21)))
	if v505 == int32(112) {
		goto L91
	} else {
		goto L95
	}
L95:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v509 = F_bms_add_member(m, v508, v55)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v509
	v515 = int32(0)
	v517 = F_makeVar(m, v55, int32(-6), int32(26), int32(-1), v515, v515)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_add_row_identity_var(m, l0, v517, v55, int32(428861))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_add_row_identity_columns(m, l0, v55, v36, l5)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	goto L91
}
