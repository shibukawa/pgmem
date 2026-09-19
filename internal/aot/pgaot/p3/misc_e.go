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
	var v34 int32
	_ = v34
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
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
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
	v147 = m.ExcPending
	if v147 != 0 {
		goto L42
	} else {
		goto L43
	}
L2:
	;
	m.G0 = v13 + int32(16)
	return v129
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v129 = int32(base.Ui32(v16) >> (uint(int32(2)) % 32))
	goto L2
L4:
	;
	goto L5
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v19 != 0 {
		v129 = v19
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
	v113 = v21 << (uint(int32(3)) % 32)
	v117 = base.I32_div_s(v22+int32(7), int32(8))
	v119 = int32(23)
	if v20 != 0 {
		goto L39
	} else {
		goto L40
	}
L8:
	;
	v103 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v29 = int32(0)
	v34 = int32(0)
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
	v103 = v96
	goto L7
L13:
	;
	v100 = v34 + int32(1)
	if v100 != v22 {
		v29 = v96
		v34 = v100
		goto L11
	} else {
		goto L38
	}
L14:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v34))))
	if v39 != 0 {
		v96 = v29
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
		v75 = v40
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v77 = v29 + v75
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+47)))
	switch v78 - int32(99) {
	case 0:
		v93 = v77
		goto L33
	case 1:
		goto L35
	default:
		goto L34
	case 6:
		goto L36
	}
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v26+v34<<(uint(int32(2))%32))))
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
	v72 = F_strlen(m, v46)
	mBase = m.M
	v75 = v72 + int32(1)
	goto L18
L23:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+1)))
	if base.Ui32((v53-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v75 = int32(6)
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
	if v53 == v60 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v64 = v60
	goto L29
L28:
	;
	v64 = int32(2)
	goto L29
L29:
	;
	v75 = v64
	goto L18
L30:
	;
	v75 = int32(base.Ui32(v49) >> (uint(int32(1)) % 32))
	goto L18
L31:
	;
	goto L32
L32:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v75 = int32(base.Ui32(v69) >> (uint(int32(2)) % 32))
	goto L18
L33:
	;
	if base.Ui32(int32(1073741824)) <= base.Ui32(v93) {
		goto L1
	} else {
		goto L37
	}
L34:
	;
	v93 = (v77 + int32(1)) & int32(-2)
	goto L33
L35:
	;
	v93 = (v77 + int32(7)) & int32(-8)
	goto L33
L36:
	;
	v93 = (v77 + int32(3)) & int32(-4)
	goto L33
L37:
	;
	v96 = v93
	goto L13
L38:
	;
	goto L12
L39:
	;
	v123 = v113 + v117 + v119
	goto L41
L40:
	;
	v123 = v113 + v119
	goto L41
L41:
	;
	v126 = v123&int32(-8) + v103
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v126
	v129 = v126
	goto L2
L42:
	;
	return int32(0)
L43:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(1073741823)
	F_errmsg(m, int32(_a_F_EA_get_flat_size_0), v13)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_EA_get_flat_size_1), int32(275), int32(_a_F_EA_get_flat_size_2))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L42
	} else {
		goto L46
	}
L46:
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
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_EndImplicitTransactionBlock[0]))
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
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
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
	var v149 int32
	_ = v149
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if l2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L13
	} else {
		goto L42
	}
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v20 = v18
	goto L4
L3:
	;
	v20 = int32(0)
	goto L4
L4:
	;
	if v17 == v20 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if v17 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
	} else {
		goto L37
	}
L8:
	;
	m.G0 = v14 + int32(48)
	return v149
L9:
	;
	v149 = int32(0)
	goto L8
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
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v32 <= int32(0) {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v38 = int32(0)
	goto L17
L17:
	;
	v48 = v38 + int32(1)
	v50 = v38 << (uint(int32(2)) % 32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v25+v50)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v54 = v53 + v50
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v57 = F_transformExpr(m, l0, v55, int32(36))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L13
	} else {
		goto L19
	}
L18:
	;
	goto L12
L19:
	;
	v59 = F_exprType(m, v57)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v61 = int32(-1)
	v65 = F_coerce_to_target_type(m, l0, v57, v59, v52, v61, int32(1), int32(2), v61)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	if v65 == int32(0) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_assign_expr_collations(m, l0, v65)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v65
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v48 < v72 {
		v38 = v48
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
		v149 = v87
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v91 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v92 <= v91 {
		v149 = v87
		goto L8
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
	v149 = v87
	goto L8
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
	v127 = int32(_a_F_EvaluateParams_0)
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_EvaluateParams[0]))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_EvaluateParams[0])) = v130
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
	*(*int32)(unsafe.Add(mBase, _c_F_EvaluateParams[0])) = v128
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
	v165 = m.ExcPending
	if v165 != 0 {
		goto L13
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l1
	F_errmsg(m, int32(_a_F_EvaluateParams_1), v14+int32(32))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L13
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v17
	F_errdetail(m, int32(_a_F_EvaluateParams_2), v14+int32(16))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L13
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_EvaluateParams_3), int32(298), int32(_a_F_EvaluateParams_4))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
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
	v190 = m.ExcPending
	if v190 != 0 {
		goto L13
	} else {
		goto L43
	}
L43:
	;
	v191 = F_format_type_be(m, v59)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	v193 = F_format_type_be(m, v52)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v48
	F_errmsg(m, int32(_a_F_EvaluateParams_5), v14)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	F_errhint(m, int32(_a_F_EvaluateParams_6), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v206 = F_exprLocation(m, v205)
	mBase = m.M
	F_parser_errposition(m, l0, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_EvaluateParams_3), int32(335), int32(_a_F_EvaluateParams_4))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
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
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
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
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
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
	return v148
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+20)) = int64(55834575290)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v64
	v66 = int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v67 <= int32(0) {
		v131 = v66
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
	v148 = int32(0)
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
	v61 = v55
	goto L4
L16:
	;
	v61 = l4
	goto L4
L17:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+19)))
	if v136 != int32(1) {
		v148 = v131
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
	v131 = v66
	goto L17
L21:
	;
	v121 = v79 + int32(1)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v121 < v122 {
		v79 = v121
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v61
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
	v131 = int32(0)
	goto L17
L31:
	;
	goto L32
L32:
	;
	if v111 == v61 {
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
	F_pfree(m, v61)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v148 = v131
	goto L3
}
func F_ExecBRInsertTriggers(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
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
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v19 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+60)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v16)+52)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v16)+44)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v16)+36)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = int64(51539607994)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v29
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v32 <= v4 {
		v204 = int32(1)
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L8
	} else {
		goto L51
	}
L2:
	;
	m.G0 = v16 + int32(80)
	return v204
L3:
	;
	v40 = v4
	v45 = v4
	goto L4
L4:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v51 = v48 + v45*int32(60)
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+12)))
	if v52&int32(71) != int32(7) {
		v187 = v40
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v204 = v195
	goto L2
L6:
	;
	v195 = int32(1)
	v197 = v45 + v195
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v197 < v198 {
		v40 = v187
		v45 = v197
		goto L4
	} else {
		goto L50
	}
L7:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v58 = int32(0)
	v60 = F_TriggerEnabled(m, l0, l1, v51, v57, v58, v58, l2)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	if v60 == int32(0) {
		v187 = v40
		goto L6
	} else {
		goto L10
	}
L10:
	;
	if v40 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v71 = F_ExecFetchSlotHeapTuple(m, l2, int32(1), v16+int32(70))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L14
	}
L12:
	;
	v73 = v40
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = l2
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v81 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v73 = v71
	goto L13
L15:
	;
	v84 = v81
	goto L17
L16:
	;
	v82 = F_MakePerTupleExprContext(m, l0)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L8
	} else {
		goto L18
	}
L17:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+20))
	v86 = F_ExecCallTriggerFunc(m, v16+int32(24), v45, v79, v80, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L19
	}
L18:
	;
	v84 = v82
	goto L17
L19:
	;
	if v86 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v90 = int32(0)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+70)))
	if v91 != int32(1) {
		v204 = v90
		goto L2
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v73 == v86 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	F_pfree(m, v73)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	v204 = v90
	goto L2
L25:
	;
	v187 = v86
	goto L6
L26:
	;
	goto L27
L27:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+52))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+16))
	if v99 == int32(0) {
		v158 = v86
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_ExecForceStoreHeapTuple(m, v158, l2, int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L8
	} else {
		goto L40
	}
L29:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+18)))
	if v102 != int32(1) {
		v158 = v86
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v105 = int32(0)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v106 <= v105 {
		v158 = v86
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v113 = v105
	v115 = v86
	v120 = v106
	goto L32
L32:
	;
	v125 = v113 + int32(1)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113*int32(100)+(v98+v120<<(uint(int32(4))%32)))+110)))
	if v130 != int32(118) {
		v149 = v115
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v158 = v149
	goto L28
L34:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v125 < v150 {
		v113 = v125
		v115 = v149
		v120 = v150
		goto L32
	} else {
		goto L39
	}
L35:
	;
	v133 = F_heap_attisnull(m, v115, v125, v98)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	if v133 != 0 {
		v149 = v115
		goto L34
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = int32(0)
	v138 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+71)) = uint8(v138)
	v147 = F_heap_modify_tuple_by_cols(m, v115, v98, v138, v16+int32(76), v16+int32(72), v16+int32(71))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v149 = v147
	goto L34
L39:
	;
	goto L33
L40:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+16)))
	if v168 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v172 = F_ExecPartitionCheck(m, l1, l2, l0, int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L8
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+70)))
	if v176 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	if v172 == int32(0) {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	F_pfree(m, v73)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L8
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v187 = int32(0)
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
	v223 = m.ExcPending
	if v223 != 0 {
		goto L8
	} else {
		goto L52
	}
L52:
	;
	F_errmsg(m, int32(_a_F_ExecBRInsertTriggers_0), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+48))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+68))
	v232 = F_get_namespace_name(m, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L8
	} else {
		goto L54
	}
L54:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v235 + int32(4)
	F_errdetail(m, int32(_a_F_ExecBRInsertTriggers_1), v16)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_ExecBRInsertTriggers_2), int32(2530), int32(_a_F_ExecBRInsertTriggers_3))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int64
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
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
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v337 int32
	_ = v337
	v10 = int32(0)
	v17 = m.G0
	v19 = v17 + int32(-64)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v22 = F_ExecGetTriggerOldSlot(m, l0, l2)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+54)) = uint8(v26)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+53)) = uint8(v26)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v26
	v32 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+36)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v19)+28)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v19)+20)) = v32
	v38 = F_ExecUpdateLockMode(m, l0, l2)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
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
	m.G0 = v19 - int32(-64)
	return v337
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = int64(60129542586)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v110
	v112 = F_ExecGetAllUpdatedCols(m, l2, l0)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L23
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = int32(0)
	v48 = F_GetTupleForTrigger(m, l0, l1, l2, l3, v38, v22, l8^int32(1), v17+int32(-4), l6, l7)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_ExecForceStoreHeapTuple(m, l4, v22, int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L22
	}
L9:
	;
	if v48 == int32(0) {
		v337 = v10
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if v52 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v52
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+72))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	m.T0[v60].(func(*base.Module, int32))(m, v58)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v99 = F_ExecFetchSlotHeapTuple(m, v22, int32(1), v17+int32(-10))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L21
	}
L14:
	;
	v63 = int32(_a_F_ExecBRUpdateTriggers_0)
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_ExecBRUpdateTriggers[0]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecBRUpdateTriggers[0])) = v66
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v72 = m.T0[v71].(func(*base.Module, int32, int32, int32) int32)(m, v53+int32(4), v57, int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecBRUpdateTriggers[0])) = v64
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+4)))
	v78 = v76 & int32(_a_F_ExecBRUpdateTriggers_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v58)+4)) = uint16(v78)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	*(*uint16)(unsafe.Add(mBase, uint32(v58)+6)) = uint16(v81)
	if v58 != l5 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+32))
	m.T0[v85].(func(*base.Module, int32, int32))(m, l5, v58)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+28))
	m.T0[v89].(func(*base.Module, int32))(m, l5)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	goto L13
L21:
	;
	v106 = v99
	goto L5
L22:
	;
	v106 = l4
	goto L5
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v112
	v115 = int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if int32(0) < v116 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v127 = int32(0)
	v130 = v10
	goto L27
L25:
	;
	goto L26
L26:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+54)))
	if v318 != int32(1) {
		v337 = v115
		goto L4
	} else {
		goto L76
	}
L27:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v139 = v136 + v127*int32(60)
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+12)))
	if v140&int32(83) != int32(19) {
		v292 = v130
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L26
L29:
	;
	v299 = v127 + int32(1)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v299 < v300 {
		v127 = v299
		v130 = v292
		goto L27
	} else {
		goto L75
	}
L30:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v146 = F_TriggerEnabled(m, l0, l2, v139, v145, v112, v22, l5)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v146 == int32(0) {
		v292 = v130
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if v130 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v155 = F_ExecFetchSlotHeapTuple(m, l5, int32(1), v17+int32(-11))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	v157 = v130
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v139
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v167 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v157 = v155
	goto L35
L37:
	;
	v170 = v167
	goto L39
L38:
	;
	v168 = F_MakePerTupleExprContext(m, l0)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+20))
	v172 = F_ExecCallTriggerFunc(m, v17+int32(-56), v127, v165, v166, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	v170 = v168
	goto L39
L41:
	;
	if v172 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+54)))
	if v176 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	if v172 == v157 {
		goto L51
	} else {
		goto L52
	}
L45:
	;
	F_pfree(m, v106)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v181 = int32(0)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+53)))
	if v182 != int32(1) {
		v337 = v181
		goto L4
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	F_pfree(m, v157)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v337 = v181
	goto L4
L51:
	;
	v292 = v172
	goto L29
L52:
	;
	goto L53
L53:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+52))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+16))
	if v190 == int32(0) {
		v247 = v172
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_ExecForceStoreHeapTuple(m, v247, l5, int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L66
	}
L55:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+18)))
	if v193 != int32(1) {
		v247 = v172
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v196 = int32(0)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	if v197 <= v196 {
		v247 = v172
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v201 = v172
	v203 = v197
	v208 = v196
	goto L58
L58:
	;
	v219 = v208 + int32(1)
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208*int32(100)+(v189+v203<<(uint(int32(4))%32)))+110)))
	if v224 != int32(118) {
		v243 = v201
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v247 = v243
	goto L54
L60:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	if v219 < v244 {
		v201 = v243
		v203 = v244
		v208 = v219
		goto L58
	} else {
		goto L65
	}
L61:
	;
	v227 = F_heap_attisnull(m, v201, v219, v189)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	if v227 != 0 {
		v243 = v201
		goto L60
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = int32(0)
	v232 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+55)) = uint8(v232)
	v241 = F_heap_modify_tuple_by_cols(m, v201, v189, v232, v17+int32(-4), v17+int32(-8), v17+int32(-9))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v243 = v241
	goto L60
L65:
	;
	goto L59
L66:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+54)))
	if base.B2i32(v265 != int32(1))|base.B2i32(v247 != v106) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)+28))
	m.T0[v273].(func(*base.Module, int32))(m, l5)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+53)))
	if v276 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	F_pfree(m, v157)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v292 = int32(0)
	goto L29
L74:
	;
	goto L73
L75:
	;
	goto L28
L76:
	;
	F_pfree(m, v106)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v337 = v115
	goto L4
}
func F_ExecBuildHash32Expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
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
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v290 int64
	_ = v290
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	v7 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	v29 = F_palloc0(m, int32(68))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(380)
	if l2 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v36 = v35
	goto L5
L4:
	;
	v36 = v7
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = l4
	v38 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = v38
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = v38
	v42 = F_expr_setup_walker(m, l2, v26)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	F_ExecPushExprSetupSteps(m, v29, v26)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if int64(2) <= base.I64_extend_i32_s(v36) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v50 = F_palloc(m, int32(8))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v52 = v7
	goto L10
L10:
	;
	v68 = int32(86)
	v71 = int32(0)
	v72 = v7
	v73 = v7
	v77 = v7
	v78 = v7
	v79 = int32(87)
	v80 = v7
	v86 = v7
	goto L14
L11:
	;
	v52 = v50
	goto L10
L12:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v283 + int32(1)
	v289 = v282 + v283*int32(40)
	v290 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v289)+8)) = v290
	*(*int64)(unsafe.Add(mBase, uint32(v289))) = v290
	*(*int32)(unsafe.Add(mBase, uint32(v289)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v289)+32)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v289)+28)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v289)+24)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v289)+20)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v289)+16)) = v78
	v301 = F_jit_compile_expr(m, v29)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L64
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v280
	v282 = v280
	goto L12
L14:
	;
	v87 = int32(0)
	if l2 == v87 {
		v97 = v87
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v269 != v174 {
		goto L60
	} else {
		goto L61
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
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v91 <= v73 {
		v97 = int32(0)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v97 = v93 + v73<<(uint(int32(2))%32)
	goto L16
L19:
	;
	goto L15
L20:
	;
	v181 = v73 << (uint(int32(2)) % 32)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v105+v181)))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0+v181)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v188 = F_palloc0(m, int32(28))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L33
	}
L21:
	;
	if v77 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(v97 == int32(0))|base.B2i32(v102 <= v73) != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v105 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v174 != 0 {
		goto L19
	} else {
		goto L31
	}
L26:
	;
	v109 = int32(0)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v110 <= v109 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v117 = v109
	goto L28
L28:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137+v117<<(uint(int32(2))%32))))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v136+v141*int32(40))+28)) = v145
	v148 = v117 + int32(1)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v148 < v149 {
		v117 = v148
		goto L28
	} else {
		goto L30
	}
L29:
	;
	goto L25
L30:
	;
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = int32(16)
	v178 = F_palloc(m, int32(640))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v280 = v178
	goto L13
L33:
	;
	v191 = F_palloc0(m, int32(28))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_fmgr_info(m, v185, v188)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_ExecInitExprRec(m, v186, v29, v191+int32(20), v191+int32(24))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v201 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v191)+18)) = uint16(v201)
	v203 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v191)+16)) = uint8(v203)
	*(*int32)(unsafe.Add(mBase, uint32(v191)+12)) = v183
	*(*int64)(unsafe.Add(mBase, uint32(v191)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = v188
	v209 = base.B2i32(v73 == v36-int32(1))
	if v73 == v36-int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v210 = v29 + int32(8)
	goto L39
L38:
	;
	v210 = v52
	goto L39
L39:
	;
	if v73 == v36-int32(1) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v211 = v29 + int32(5)
	goto L42
L41:
	;
	v211 = v52 + int32(4)
	goto L42
L42:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v73))))
	if v213 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v214 = v79
	goto L45
L44:
	;
	v214 = v68
	goto L45
L45:
	;
	if l5 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v215 = v68
	goto L48
L47:
	;
	v215 = v214
	goto L48
L48:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v217 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	v240 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v239 + v240
	v245 = v238 + v239*int32(40)
	v246 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v245)+36)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v245)+32)) = v52
	v249 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v245)+28)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(v245)+24)) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v245)+20)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v245)+16)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v245)+12)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v245)+8)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v245)+4)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = v215
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	v267 = F_lappend_int(m, v77, v264-v240)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L59
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v236
	v238 = v236
	goto L49
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = int32(16)
	v223 = F_palloc(m, int32(640))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v225 != v217 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v236 = v223
	goto L50
L55:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v238 = v227
	goto L49
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v217 << (uint(int32(1)) % 32)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v234 = F_repalloc(m, v231, v217*int32(80))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v236 = v234
	goto L50
L59:
	;
	v68 = int32(88)
	v71 = v191
	v72 = v52
	v73 = v73 + v240
	v77 = v267
	v78 = v188
	v79 = int32(89)
	v80 = v216
	v86 = v249
	goto L14
L60:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v282 = v271
	goto L12
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v174 << (uint(int32(1)) % 32)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v278 = F_repalloc(m, v275, v174*int32(80))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v280 = v278
	goto L13
L64:
	;
	if v301 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	F_ExecReadyInterpretedExpr(m, v29)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	m.G0 = v26 + int32(16)
	return v29
L68:
	;
	goto L67
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
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCheckPermissions[0]))
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
		v45 = int32(20)
		goto L16
	default:
		goto L17
	case 10:
		goto L21
	case 29:
		goto L18
	case 36:
		goto L19
	case 45:
		goto L20
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
	v45 = int32(41)
	goto L16
L18:
	;
	v47 = int32(18)
	goto L15
L19:
	;
	v47 = int32(23)
	goto L15
L20:
	;
	v47 = int32(51)
	goto L15
L21:
	;
	v47 = int32(37)
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
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
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if l1 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L20
	} else {
		goto L55
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L20
	} else {
		goto L47
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L20
	} else {
		goto L42
	}
L4:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v109 != v116 {
		goto L1
	} else {
		goto L41
	}
L5:
	;
	v109 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v16 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v17 <= v16 {
		v109 = v16
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v20 = v16
	goto L9
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v27 <= v20 {
		goto L3
	} else {
		goto L11
	}
L10:
	;
	v109 = v37
	goto L4
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v20<<(uint(int32(2))%32))))
	v36 = int32(1)
	v37 = v20 + v36
	v41 = v20*int32(100) + (v12 + v27<<(uint(int32(4))%32))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+111)))
	if v42 == v36 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v37 < v107 {
		v20 = v37
		goto L9
	} else {
		goto L40
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v46 == int32(7) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v72 = v41 + int32(20)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+90)))
	if v73 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+24)))
	if v49 != 0 {
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
	v53 = m.ExcPending
	if v53 != 0 {
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
	v56 = m.ExcPending
	if v56 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errmsg(m, int32(_a_F_ExecCheckPlanOutput_0), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v37
	F_errdetail(m, int32(_a_F_ExecCheckPlanOutput_1), v10)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_ExecCheckPlanOutput_2), int32(232), int32(_a_F_ExecCheckPlanOutput_3))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
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
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if v74 == int32(7) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v100 = F_exprType(m, v70)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L20
	} else {
		goto L38
	}
L29:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+24)))
	if v77 != 0 {
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
	v81 = m.ExcPending
	if v81 != 0 {
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
	v84 = m.ExcPending
	if v84 != 0 {
		goto L20
	} else {
		goto L34
	}
L34:
	;
	F_errmsg(m, int32(_a_F_ExecCheckPlanOutput_0), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L20
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v37
	F_errdetail(m, int32(_a_F_ExecCheckPlanOutput_4), v10+int32(32))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L20
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_ExecCheckPlanOutput_2), int32(249), int32(_a_F_ExecCheckPlanOutput_3))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
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
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v72)+68))
	if v100 != v102 {
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
	m.G0 = v10 + int32(48)
	return
L42:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L20
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(_a_F_ExecCheckPlanOutput_0), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L20
	} else {
		goto L44
	}
L44:
	;
	F_errdetail(m, int32(_a_F_ExecCheckPlanOutput_5), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L20
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_ExecCheckPlanOutput_2), int32(212), int32(_a_F_ExecCheckPlanOutput_3))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
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
	v147 = m.ExcPending
	if v147 != 0 {
		goto L20
	} else {
		goto L48
	}
L48:
	;
	F_errmsg(m, int32(_a_F_ExecCheckPlanOutput_0), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L20
	} else {
		goto L49
	}
L49:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v72)+68))
	v153 = F_format_type_be(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L20
	} else {
		goto L50
	}
L50:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v156 = F_exprType(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L20
	} else {
		goto L51
	}
L51:
	;
	v158 = F_format_type_be(m, v156)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L20
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v153
	F_errdetail(m, int32(_a_F_ExecCheckPlanOutput_6), v10+int32(16))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L20
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_ExecCheckPlanOutput_2), int32(261), int32(_a_F_ExecCheckPlanOutput_3))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
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
	v179 = m.ExcPending
	if v179 != 0 {
		goto L20
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(_a_F_ExecCheckPlanOutput_0), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L20
	} else {
		goto L57
	}
L57:
	;
	F_errdetail(m, int32(_a_F_ExecCheckPlanOutput_7), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L20
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_ExecCheckPlanOutput_2), int32(268), int32(_a_F_ExecCheckPlanOutput_3))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
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
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+16)))
	if v20 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v319 = m.G0
	v321 = v319 - int32(32)
	m.G0 = v321
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+52))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v332 != 0 {
		goto L89
	} else {
		goto L90
	}
L2:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	v91 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+122)))
	if v91 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v23 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v26 = int32(1)
	v31 = v26
	v32 = int32(0)
	v33 = v23
	v37 = v26
	goto L5
L5:
	;
	v45 = v18 + v33<<(uint(int32(4))%32) + v31*int32(100)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+6)))
	if v46 != int32(1) {
		v68 = v32
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v68 == int32(0) {
		goto L2
	} else {
		goto L20
	}
L7:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v71 = v37 + int32(1)
	v72 = base.I32_extend16_s(v71)
	if v72 <= v69 {
		v31 = v72
		v32 = v68
		v33 = v69
		v37 = v71
		goto L5
	} else {
		goto L19
	}
L8:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45-int32(80))+90)))
	if v51 == int32(118) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v54 = F_lappend_int(m, v32, v31)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v56 < base.I32_extend16_s(v37) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	return
L13:
	;
	v68 = v54
	goto L7
L14:
	;
	F_slot_getsomeattrs_int(m, l1, v31)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v63 = int32(1)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v31-v63))))
	if v65 == v63 {
		v315 = v31
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v68 = v32
	goto L7
L19:
	;
	goto L6
L20:
	;
	v76 = F_ExecRelGenVirtualNotNull(m, l0, l1, l2, v68)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	if v76 != 0 {
		v315 = v76
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L2
L23:
	;
	m.G0 = v15 + int32(48)
	return
L24:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+52))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96)+14)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)+48))
	v99 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98)+122)))
	if v97 == v99 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v102 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L12
	} else {
		goto L85
	}
L28:
	;
	v106 = int32(_a_F_ExecConstraints_0)
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_ExecConstraints[0]))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l2)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecConstraints[0])) = v109
	v113 = F_palloc0(m, v97<<(uint(int32(2))%32))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L12
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
	*(*int32)(unsafe.Add(mBase, _c_F_ExecConstraints[0])) = v107
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
		goto L12
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
		goto L12
	} else {
		goto L41
	}
L41:
	;
	v140 = F_ExecPrepareExpr(m, v138, l2)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L12
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
		goto L12
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
		goto L23
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
		goto L23
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
		goto L12
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
	goto L23
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
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+56))
	v249 = F_ExecBuildSlotValueDescription(m, v248, v244, v246, v245)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L12
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
		goto L12
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
		goto L12
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
	v225 = F_MakeTupleTableSlot(m, v220, int32(_a_F_ExecConstraints_1))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L12
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
		goto L12
	} else {
		goto L69
	}
L67:
	;
	v227 = F_execute_attr_map_slot(m, v222, l1, v225)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L12
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
		goto L12
	} else {
		goto L70
	}
L70:
	;
	v234 = F_bms_union(m, v230, v232)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	v244 = v229
	v245 = v234
	v246 = v220
	v247 = v236
	goto L59
L72:
	;
	v239 = F_ExecGetUpdatedCols(m, l0, l2)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	v241 = F_bms_union(m, v237, v239)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	v244 = l1
	v245 = v241
	v246 = v18
	v247 = v17
	goto L59
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L12
	} else {
		goto L76
	}
L76:
	;
	F_errcode(m, int32(67391682))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L12
	} else {
		goto L77
	}
L77:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v258 + int32(4)
	F_errmsg(m, int32(_a_F_ExecConstraints_2), v15+int32(16))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L12
	} else {
		goto L78
	}
L78:
	;
	if v249 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v249
	F_errdetail(m, int32(_a_F_ExecConstraints_3), v15)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L12
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
	v273 = m.ExcPending
	if v273 != 0 {
		goto L12
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	F_errfinish(m, int32(_a_F_ExecConstraints_4), int32(2081), int32(_a_F_ExecConstraints_5))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L12
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
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v94)+48))
	v284 = int32(*(*int16)(unsafe.Add(mBase, uint32(v283)+122)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v283 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v284 - v97
	F_errmsg_internal(m, int32(_a_F_ExecConstraints_6), v15+int32(32))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L12
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_ExecConstraints_4), int32(1798), int32(_a_F_ExecConstraints_7))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L12
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)+56))
	v363 = F_ExecBuildSlotValueDescription(m, v362, v358, v357, v359)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L12
	} else {
		goto L104
	}
L89:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)+8))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+52))
	v336 = F_build_attrmap_by_name_if_req(m, v324, v334, int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L12
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v351 = F_ExecGetInsertedCols(m, l0, l2)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L12
	} else {
		goto L101
	}
L92:
	;
	if v336 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v339 = F_MakeTupleTableSlot(m, v334, int32(_a_F_ExecConstraints_1))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L12
	} else {
		goto L96
	}
L94:
	;
	v343 = l1
	goto L95
L95:
	;
	v344 = F_ExecGetInsertedCols(m, v332, l2)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L12
	} else {
		goto L98
	}
L96:
	;
	v341 = F_execute_attr_map_slot(m, v336, l1, v339)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L12
	} else {
		goto L97
	}
L97:
	;
	v343 = v341
	goto L95
L98:
	;
	v346 = F_ExecGetUpdatedCols(m, v332, l2)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L12
	} else {
		goto L99
	}
L99:
	;
	v348 = F_bms_union(m, v344, v346)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L12
	} else {
		goto L100
	}
L100:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v332)+8))
	v357 = v334
	v358 = v343
	v359 = v348
	v361 = v350
	goto L88
L101:
	;
	v353 = F_ExecGetUpdatedCols(m, l0, l2)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L12
	} else {
		goto L102
	}
L102:
	;
	v355 = F_bms_union(m, v351, v353)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L12
	} else {
		goto L103
	}
L103:
	;
	v357 = v324
	v358 = l1
	v359 = v355
	v361 = v323
	goto L88
L104:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L12
	} else {
		goto L105
	}
L105:
	;
	F_errcode(m, int32(33575106))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L12
	} else {
		goto L106
	}
L106:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v323)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+16)) = v324 + v325<<(uint(int32(4))%32) + v315*int32(100) - int32(76)
	*(*int32)(unsafe.Add(mBase, uint32(v321)+20)) = v372 + int32(4)
	F_errmsg(m, int32(_a_F_ExecConstraints_8), v321+int32(16))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L12
	} else {
		goto L107
	}
L107:
	;
	if v363 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321))) = v363
	F_errdetail(m, int32(_a_F_ExecConstraints_3), v321)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L12
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	F_errtablecol(m, v323, v315)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L12
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	F_errfinish(m, int32(_a_F_ExecConstraints_4), int32(2219), int32(_a_F_ExecConstraints_9))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L12
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
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
				v84 = v82 & int32(_a_F_ExecFilterJunk_0)
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
			v84 = v82 & int32(_a_F_ExecFilterJunk_0)
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
	var v82 int32
	_ = v82
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v8 != 0 {
		v9 = v8
	} else {
		v9 = l0
	}
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v10 == int32(0) {
		v82 = v3
		return v82
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
				v82 = v3
				return v82
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				if v28 == int32(0) {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
					v82 = v77
					return v82
				} else {
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)))
					if v31 == int32(0) {
						v34 = int32(_a_F_ExecGetUpdatedCols_0)
						v35 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGetUpdatedCols[0]))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
						*(*int32)(unsafe.Add(mBase, _c_F_ExecGetUpdatedCols[0])) = v41
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
									*(*int32)(unsafe.Add(mBase, _c_F_ExecGetUpdatedCols[0])) = v35
									v56 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v56)
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
									if v63 == int32(0) {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
										v82 = v77
										return v82
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
								*(*int32)(unsafe.Add(mBase, _c_F_ExecGetUpdatedCols[0])) = v35
								v56 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v56)
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
								if v63 == int32(0) {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
									v82 = v77
									return v82
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
							v82 = v77
							return v82
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
	var v68 int32
	_ = v68
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
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v179 int64
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int64
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int64
	_ = v251
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
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
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int64
	_ = v329
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int64
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int64
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
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
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v447 int32
	_ = v447
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int64
	_ = v476
	var v477 int32
	_ = v477
	var v482 int64
	_ = v482
	var v487 int32
	_ = v487
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v513 int64
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v566 int32
	_ = v566
	var v567 int64
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v755 int32
	_ = v755
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v795 int64
	_ = v795
	var v800 int32
	_ = v800
	var v805 int64
	_ = v805
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int64
	_ = v943
	var v946 int32
	_ = v946
	var v950 int64
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
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1110 int32
	_ = v1110
	var v1115 int32
	_ = v1115
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1129 int32
	_ = v1129
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1138 int64
	_ = v1138
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1160 int32
	_ = v1160
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1221 int64
	_ = v1221
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int64
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1232 int64
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1237 int64
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1262 int32
	_ = v1262
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1287 int32
	_ = v1287
	var v1291 int32
	_ = v1291
	var v1295 int32
	_ = v1295
	var v1299 int32
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1311 int32
	_ = v1311
	var v1315 int32
	_ = v1315
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1334 int32
	_ = v1334
	var v1342 int32
	_ = v1342
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1388 int32
	_ = v1388
	var v1392 int32
	_ = v1392
	var v1398 int32
	_ = v1398
	var v1403 int32
	_ = v1403
	var v1408 int32
	_ = v1408
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1417 int64
	_ = v1417
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1429 int32
	_ = v1429
	var v1443 int32
	_ = v1443
	var v1458 int32
	_ = v1458
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1476 int32
	_ = v1476
	var v1481 int32
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1521 int64
	_ = v1521
	var v1527 int32
	_ = v1527
	var v1530 int32
	_ = v1530
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1538 int64
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1540 int64
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1545 int64
	_ = v1545
	var v1546 int32
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
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1559 int32
	_ = v1559
	var v1561 int32
	_ = v1561
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1579 int32
	_ = v1579
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1633 int32
	_ = v1633
	var v1637 int32
	_ = v1637
	var v1643 int32
	_ = v1643
	var v1648 int32
	_ = v1648
	var v1653 int32
	_ = v1653
	var v1658 int32
	_ = v1658
	var v1663 int32
	_ = v1663
	var v1668 int32
	_ = v1668
	var v1673 int32
	_ = v1673
	var v1700 int32
	_ = v1700
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1724 int64
	_ = v1724
	var v1726 int64
	_ = v1726
	var v1728 int64
	_ = v1728
	var v1730 int64
	_ = v1730
	var v1732 int64
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1744 int32
	_ = v1744
	var v1747 int32
	_ = v1747
	var v1751 int32
	_ = v1751
	var v1753 int32
	_ = v1753
	var v1777 int32
	_ = v1777
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1823 int32
	_ = v1823
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
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
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L377
L2:
	;
	F_ExecGrant_common(m, l0, int32(1247), int64(256), int32(461))
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L14
	} else {
		goto L376
	}
L3:
	;
	F_ExecGrant_common(m, l0, int32(2328), int64(256), int32(0))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L14
	} else {
		goto L375
	}
L4:
	;
	F_ExecGrant_common(m, l0, int32(1417), int64(256), int32(0))
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L14
	} else {
		goto L374
	}
L5:
	;
	F_ExecGrant_common(m, l0, int32(1255), int64(128), int32(0))
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L14
	} else {
		goto L373
	}
L6:
	;
	F_ExecGrant_common(m, l0, int32(2612), int64(256), int32(460))
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		goto L14
	} else {
		goto L372
	}
L7:
	;
	v1414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v1414 != int32(1) {
		goto L328
	} else {
		goto L329
	}
L8:
	;
	F_ExecGrant_common(m, l0, int32(2615), int64(768), int32(0))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L14
	} else {
		goto L327
	}
L9:
	;
	F_ExecGrant_common(m, l0, int32(1213), int64(512), int32(0))
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L14
	} else {
		goto L326
	}
L10:
	;
	v1135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v1135 != int32(1) {
		goto L259
	} else {
		goto L260
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L14
	} else {
		goto L256
	}
L12:
	;
	F_ExecGrant_common(m, l0, int32(1262), int64(3584), int32(0))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L14
	} else {
		goto L255
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
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L14
	} else {
		goto L252
	}
L18:
	;
	F_relation_close(m, v40, int32(3))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L14
	} else {
		goto L250
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
	v68 = v2
	goto L21
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v68<<(uint(int32(2))%32))))
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
	v755 = int32(0)
	if base.B2i32(v733 == v755)|base.B2i32(v180 < int32(-7)) == v755 {
		goto L181
	} else {
		goto L182
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L14
	} else {
		goto L178
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L14
	} else {
		goto L173
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L14
	} else {
		goto L168
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L14
	} else {
		goto L165
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L14
	} else {
		goto L161
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L14
	} else {
		goto L157
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
	v608 = m.ExcPending
	if v608 != 0 {
		goto L14
	} else {
		goto L154
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
	F_errmsg(m, int32(_a_F_ExecGrantStmt_oids_0), v29+int32(160))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1823), int32(_a_F_ExecGrantStmt_oids_2))
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
	F_errmsg(m, int32(_a_F_ExecGrantStmt_oids_3), v29+int32(176))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L14
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1830), int32(_a_F_ExecGrantStmt_oids_2))
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
		v179 = v146
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v180 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84)+120)))
	v182 = v180 + int32(8)
	v185 = F_palloc0(m, v182<<(uint(int32(3))%32))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
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
		v179 = v146
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
	v155 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L14
	} else {
		goto L63
	}
L63:
	;
	if v155 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L14
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v179 = v146 & int64(262)
	goto L57
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+112)) = v84 + int32(4)
	F_errmsg(m, int32(_a_F_ExecGrantStmt_oids_4), v29+int32(112))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L14
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1876), int32(_a_F_ExecGrantStmt_oids_2))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L14
	} else {
		goto L69
	}
L69:
	;
	goto L66
L70:
	;
	v179 = v146
	goto L57
L71:
	;
	v187 = int32(0)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v188 != 0 {
		v267 = v187
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v84)+80))
	v294 = F_SysCacheGetAttr(m, int32(57), v80, int32(32), v29+int32(507))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L14
	} else {
		goto L89
	}
L73:
	;
	v190 = v179 & int64(39)
	if v190 == int64(0) {
		v267 = v187
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v193 = int32(-6)
	v196 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84)+120)))
	if v196 < v193 {
		v267 = int32(1)
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v201 = int32(_a_F_ExecGrantStmt_oids_5)
	v205 = v193
	goto L76
L76:
	;
	if v201&int32(_a_F_ExecGrantStmt_oids_6) == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v267 = v257
	goto L72
L78:
	;
	v257 = int32(1)
	v260 = base.I32_extend16_s(v201 + v257)
	v261 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84)+120)))
	if v260 <= v261 {
		v201 = v260
		v205 = v260
		goto L76
	} else {
		goto L88
	}
L79:
	;
	if base.I32_extend16_s(v201) < int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+119)))
	if v232 == int32(118) {
		goto L78
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v236 = F_SearchSysCache2(m, int32(7), v79, v205)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L14
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	if v236 == int32(0) {
		goto L27
	} else {
		goto L85
	}
L85:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v236)+16))
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+22)))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240+v241)+91)))
	F_ReleaseCatCache(m, v236)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L14
	} else {
		goto L86
	}
L86:
	;
	if v243 != 0 {
		goto L78
	} else {
		goto L87
	}
L87:
	;
	v250 = v185 + v205<<(uint(int32(3))%32) + int32(56)
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v250)))
	*(*int64)(unsafe.Add(mBase, uint32(v250))) = v251 | v190
	goto L78
L88:
	;
	goto L77
L89:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+507)))
	if v296 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v318 = F_aclcopy(m, v316)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L14
	} else {
		goto L100
	}
L91:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+119)))
	if v302 == int32(83) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	v310 = F_pg_detoast_datum_copy(m, v294)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L14
	} else {
		goto L98
	}
L94:
	;
	v305 = int32(37)
	goto L96
L95:
	;
	v305 = int32(41)
	goto L96
L96:
	;
	v306 = F_acldefault(m, v305, v289)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L14
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+428)) = int32(0)
	v316 = v306
	v317 = int32(0)
	goto L90
L98:
	;
	v314 = F_aclmembers(m, v310, v29+int32(428))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L14
	} else {
		goto L99
	}
L99:
	;
	v316 = v310
	v317 = v314
	goto L90
L100:
	;
	if v179 != int64(0) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v432 == int32(0) {
		v733 = v267
		goto L23
	} else {
		goto L124
	}
L102:
	;
	v323 = v29 + int32(288)
	v324 = int32(0)
	base.MemoryFill(m, v323, v324, int32(136))
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+464)) = uint16(v324)
	v329 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+456)) = v329
	*(*int64)(unsafe.Add(mBase, uint32(v29)+448)) = v329
	*(*int64)(unsafe.Add(mBase, uint32(v29)+440)) = v329
	*(*int64)(unsafe.Add(mBase, uint32(v29)+432)) = v329
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+272)) = uint16(v324)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+264)) = v329
	*(*int64)(unsafe.Add(mBase, uint32(v29)+256)) = v329
	*(*int64)(unsafe.Add(mBase, uint32(v29)+248)) = v329
	*(*int64)(unsafe.Add(mBase, uint32(v29)+240)) = v329
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[0]))
	F_select_best_grantor(m, v348, v179, v316, v289, v29+int32(508), v29+int32(496))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L14
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	F_UnlockTuple(m, v36, v80+int32(4), int32(7))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L14
	} else {
		goto L123
	}
L105:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v29)+496))
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v29)+508))
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+119)))
	if v361 == int32(83) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v364 = int32(37)
	goto L108
L107:
	;
	v364 = int32(41)
	goto L108
L108:
	;
	v367 = int32(0)
	v369 = F_restrict_and_check_grant(m, v355, v356, v357, v179, v79, v358, v364, v84+int32(4), v367, v367)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L14
	} else {
		goto L109
	}
L109:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v29)+508))
	v376 = F_merge_acl_with_grant(m, v316, v371, v372, v373, v374, v369, v375, v289)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L14
	} else {
		goto L110
	}
L110:
	;
	v380 = F_aclmembers(m, v376, v29+int32(488))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L14
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+412)) = v376
	v383 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+271)) = uint8(v383)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
	v390 = F_heap_modify_tuple(m, v80, v385, v323, v29+int32(432), v29+int32(240))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L14
	} else {
		goto L112
	}
L112:
	;
	F_CatalogTupleUpdate(m, v36, v390+int32(4), v390)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L14
	} else {
		goto L113
	}
L113:
	;
	F_UnlockTuple(m, v36, v80+int32(4), int32(7))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L14
	} else {
		goto L114
	}
L114:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[1])))
	if v402 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v29)+428))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v29)+488))
	F_updateAclDependencies(m, int32(1259), v79, int32(0), v289, v317, v417, v380, v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L14
	} else {
		goto L121
	}
L116:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[2])))
	if v406&int32(1) == int32(0) {
		goto L115
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	F_recordExtensionInitPrivWorker(m, v79, int32(1259), int32(0), v376)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L14
	} else {
		goto L120
	}
L119:
	;
	goto L118
L120:
	;
	goto L115
L121:
	;
	F_pfree(m, v376)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L14
	} else {
		goto L122
	}
L122:
	;
	goto L101
L123:
	;
	goto L101
L124:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v432)+4))
	if v435 <= int32(0) {
		v733 = v267
		goto L23
	} else {
		goto L125
	}
L125:
	;
	v447 = int32(0)
	goto L126
L126:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v432)+12))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v467+v447<<(uint(int32(2))%32))))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)+4))
	if v472 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v733 = v600
	goto L23
L128:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+119)))
	if base.B2i32(v482&int64(37) == int64(0))|base.B2i32(v487 != int32(83)) == int32(0) {
		goto L134
	} else {
		goto L135
	}
L129:
	;
	v482 = int64(39)
	goto L128
L130:
	;
	goto L131
L131:
	;
	v476 = F_string_to_privilege(m, v472)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L14
	} else {
		goto L132
	}
L132:
	;
	if v476&int64(32728) != int64(0) {
		goto L26
	} else {
		goto L133
	}
L133:
	;
	v482 = v476
	goto L128
L134:
	;
	v495 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L14
	} else {
		goto L137
	}
L135:
	;
	v513 = v482
	goto L136
L136:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v471)+8))
	if v514 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L137:
	;
	if v495 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L14
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v513 = v482 & int64(2)
	goto L136
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = v84 + int32(4)
	F_errmsg(m, int32(_a_F_ExecGrantStmt_oids_7), v29-int32(-64))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L14
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(2071), int32(_a_F_ExecGrantStmt_oids_2))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L14
	} else {
		goto L143
	}
L143:
	;
	goto L140
L144:
	;
	v600 = int32(1)
	v602 = v447 + v600
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v432)+4))
	if v602 < v603 {
		v447 = v602
		goto L126
	} else {
		goto L153
	}
L145:
	;
	v517 = int32(0)
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	if v518 <= v517 {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v525 = v517
	goto L147
L147:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v514)+12))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v547+v525<<(uint(int32(2))%32))))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)+4))
	v553 = F_get_attnum(m, v79, v552)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L14
	} else {
		goto L149
	}
L148:
	;
	goto L144
L149:
	;
	if v553 == int32(0) {
		goto L25
	} else {
		goto L150
	}
L150:
	;
	v559 = base.I32_extend16_s(v553 + int32(7))
	if base.B2i32(v559 <= int32(0))|base.B2i32(v182 <= v559) != 0 {
		goto L24
	} else {
		goto L151
	}
L151:
	;
	v566 = v185 + v559<<(uint(int32(3))%32)
	v567 = *(*int64)(unsafe.Add(mBase, uint32(v566)))
	*(*int64)(unsafe.Add(mBase, uint32(v566))) = v567 | v513
	v571 = v525 + int32(1)
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	if v571 < v572 {
		v525 = v571
		goto L147
	} else {
		goto L152
	}
L152:
	;
	goto L148
L153:
	;
	goto L127
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v79
	F_errmsg_internal(m, int32(_a_F_ExecGrantStmt_oids_8), v29+int32(16))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L14
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1814), int32(_a_F_ExecGrantStmt_oids_2))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L14
	} else {
		goto L156
	}
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L157:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L14
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+144)) = v84 + int32(4)
	F_errmsg(m, int32(_a_F_ExecGrantStmt_oids_9), v29+int32(144))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L14
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1838), int32(_a_F_ExecGrantStmt_oids_2))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L14
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L14
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+128)) = int32(_a_F_ExecGrantStmt_oids_10)
	F_errmsg(m, int32(_a_F_ExecGrantStmt_oids_11), v29+int32(128))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L14
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1893), int32(_a_F_ExecGrantStmt_oids_2))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L14
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+100)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v29)+96)) = v205
	F_errmsg_internal(m, int32(_a_F_ExecGrantStmt_oids_12), v29+int32(96))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L14
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1621), int32(_a_F_ExecGrantStmt_oids_13))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L14
	} else {
		goto L167
	}
L167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L168:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L14
	} else {
		goto L169
	}
L169:
	;
	v682 = F_privilege_to_string(m, v476)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L14
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+80)) = v682
	F_errmsg(m, int32(_a_F_ExecGrantStmt_oids_14), v29+int32(80))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L14
	} else {
		goto L171
	}
L171:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(2058), int32(_a_F_ExecGrantStmt_oids_2))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L14
	} else {
		goto L172
	}
L172:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L173:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L14
	} else {
		goto L174
	}
L174:
	;
	v702 = F_get_rel_name(m, v79)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L14
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v702
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v552
	F_errmsg(m, int32(_a_F_ExecGrantStmt_oids_15), v29+int32(48))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L14
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1578), int32(_a_F_ExecGrantStmt_oids_16))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L14
	} else {
		goto L177
	}
L177:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L178:
	;
	F_errmsg_internal(m, int32(_a_F_ExecGrantStmt_oids_17), int32(0))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L14
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1581), int32(_a_F_ExecGrantStmt_oids_16))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L14
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
	v764 = int32(0)
	v770 = v764
	v772 = v764
	goto L184
L182:
	;
	goto L183
L183:
	;
	F_pfree(m, v318)
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L14
	} else {
		goto L245
	}
L184:
	;
	v795 = *(*int64)(unsafe.Add(mBase, uint32(v185+v770<<(uint(int32(3))%32))))
	if v795 != int64(0) {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	goto L183
L186:
	;
	v800 = int32(0)
	base.MemoryFill(m, v29+int32(288), v800, int32(100))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+456)) = uint8(v800)
	v805 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+448)) = v805
	*(*int64)(unsafe.Add(mBase, uint32(v29)+440)) = v805
	*(*int64)(unsafe.Add(mBase, uint32(v29)+432)) = v805
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+264)) = uint8(v800)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+256)) = v805
	*(*int64)(unsafe.Add(mBase, uint32(v29)+248)) = v805
	*(*int64)(unsafe.Add(mBase, uint32(v29)+240)) = v805
	v819 = int32(7)
	v822 = base.I32_extend16_s(v772 - v819)
	v823 = F_SearchSysCache2(m, v819, v79, v822)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L14
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v1027 = v772 + int32(1)
	v1028 = base.I32_extend16_s(v1027)
	if v1028 < v182 {
		v770 = v1028
		v772 = v1027
		goto L184
	} else {
		goto L244
	}
L189:
	;
	if v823 == int32(0) {
		goto L17
	} else {
		goto L190
	}
L190:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v823)+16))
	v828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v827)+22)))
	v834 = F_SysCacheGetAttr(m, int32(7), v823, int32(22), v29+int32(492))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L14
	} else {
		goto L191
	}
L191:
	;
	v836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+492)))
	if v836 == int32(1) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v853 = m.G0
	v855 = v853 - int32(16)
	m.G0 = v855
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v851)+16))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v318)+16))
	v859 = v857 + v858
	if int32(0) <= v859 {
		goto L200
	} else {
		goto L201
	}
L193:
	;
	v841 = F_acldefault(m, int32(6), v289)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L14
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v845 = F_pg_detoast_datum_copy(m, v834)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L14
	} else {
		goto L197
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = int32(0)
	v851 = v841
	v852 = int32(0)
	goto L192
L197:
	;
	v849 = F_aclmembers(m, v845, v29+int32(488))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L14
	} else {
		goto L198
	}
L198:
	;
	v851 = v845
	v852 = v849
	goto L192
L199:
	;
	v933 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[0]))
	F_select_best_grantor(m, v933, v795, v866, v289, v29+int32(508), v29+int32(496))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L14
	} else {
		goto L219
	}
L200:
	;
	v865 = v859<<(uint(int32(4))%32) + int32(24)
	v866 = F_palloc0(m, v865)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L14
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L14
	} else {
		goto L216
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v866)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v866)+12)) = int32(1033)
	*(*int64)(unsafe.Add(mBase, uint32(v866)+4)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v866))) = v865 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v866)+16)) = v859
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	if v878 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	v888 = (v881<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L206
L205:
	;
	v888 = v878
	goto L206
L206:
	;
	v890 = v866 + int32(24)
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v318)+16))
	v893 = v891 << (uint(int32(4)) % 32)
	if v893 != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	base.MemoryCopy(m, v890, v888+v318, v893)
	goto L209
L208:
	;
	goto L209
L209:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v851)+8))
	if v896 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v851)+4))
	v906 = (v899<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L212
L211:
	;
	v906 = v896
	goto L212
L212:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v851)+16))
	v909 = v907 << (uint(int32(4)) % 32)
	if v909 != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v318)+16))
	base.MemoryCopy(m, v890+v910<<(uint(int32(4))%32), v851+v906, v909)
	goto L215
L214:
	;
	goto L215
L215:
	;
	m.G0 = v855 + int32(16)
	goto L199
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v855))) = v859
	F_errmsg_internal(m, int32(_a_F_ExecGrantStmt_oids_18), v855)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L14
	} else {
		goto L217
	}
L217:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_19), int32(433), int32(_a_F_ExecGrantStmt_oids_20))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L14
	} else {
		goto L218
	}
L218:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L219:
	;
	F_pfree(m, v866)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L14
	} else {
		goto L220
	}
L220:
	;
	v942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v943 = *(*int64)(unsafe.Add(mBase, uint32(v29)+496))
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v29)+508))
	v950 = F_restrict_and_check_grant(m, v942, v943, base.B2i32(v795 == int64(39)), v795, v79, v946, int32(6), v84+int32(4), v822, v827+v828+int32(4))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L14
	} else {
		goto L221
	}
L221:
	;
	v952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v954 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v29)+508))
	v957 = F_merge_acl_with_grant(m, v851, v952, v953, v954, v955, v950, v956, v289)
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L14
	} else {
		goto L222
	}
L222:
	;
	v961 = F_aclmembers(m, v957, v29+int32(484))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L14
	} else {
		goto L223
	}
L223:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v957)+16))
	if int32(0) < v963 {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	F_pfree(m, v957)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L14
	} else {
		goto L242
	}
L225:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v40)+52))
	v981 = F_heap_modify_tuple(m, v823, v974, v29+int32(288), v29+int32(432), v29+int32(240))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L14
	} else {
		goto L230
	}
L226:
	;
	v966 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+261)) = uint8(v966)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+372)) = v957
	goto L225
L227:
	;
	goto L228
L228:
	;
	v969 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+453)) = uint8(v969)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+261)) = uint8(v969)
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+492)))
	if v973 != 0 {
		goto L224
	} else {
		goto L229
	}
L229:
	;
	goto L225
L230:
	;
	F_CatalogTupleUpdate(m, v40, v981+int32(4), v981)
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L14
	} else {
		goto L231
	}
L231:
	;
	v988 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[1])))
	if v988 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v29)+488))
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v29)+484))
	F_updateAclDependencies(m, int32(1259), v79, v822, v289, v852, v1006, v961, v1007)
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L14
	} else {
		goto L241
	}
L233:
	;
	v992 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[2])))
	if v992&int32(1) == int32(0) {
		goto L232
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	v998 = int32(0)
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v957)+16))
	if v998 < v999 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	goto L235
L237:
	;
	v1002 = v957
	goto L239
L238:
	;
	v1002 = v998
	goto L239
L239:
	;
	F_recordExtensionInitPrivWorker(m, v79, int32(1259), v822, v1002)
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L14
	} else {
		goto L240
	}
L240:
	;
	goto L232
L241:
	;
	goto L224
L242:
	;
	F_ReleaseCatCache(m, v823)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L14
	} else {
		goto L243
	}
L243:
	;
	goto L188
L244:
	;
	goto L185
L245:
	;
	F_pfree(m, v185)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L14
	} else {
		goto L246
	}
L246:
	;
	F_ReleaseCatCache(m, v80)
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L14
	} else {
		goto L247
	}
L247:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L14
	} else {
		goto L248
	}
L248:
	;
	v1065 = v68 + int32(1)
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v1065 < v1066 {
		v68 = v1065
		goto L21
	} else {
		goto L249
	}
L249:
	;
	goto L22
L250:
	;
	F_relation_close(m, v36, int32(3))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L14
	} else {
		goto L251
	}
L251:
	;
	goto L1
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v822
	F_errmsg_internal(m, int32(_a_F_ExecGrantStmt_oids_12), v29+int32(32))
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L14
	} else {
		goto L253
	}
L253:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1668), int32(_a_F_ExecGrantStmt_oids_21))
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L14
	} else {
		goto L254
	}
L254:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L255:
	;
	goto L1
L256:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v1125
	F_errmsg_internal(m, int32(_a_F_ExecGrantStmt_oids_22), v29)
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L14
	} else {
		goto L257
	}
L257:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(645), int32(_a_F_ExecGrantStmt_oids_23))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L14
	} else {
		goto L258
	}
L258:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L259:
	;
	v1145 = F_table_open(m, int32(_a_F_ExecGrantStmt_oids_24), int32(3))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L14
	} else {
		goto L262
	}
L260:
	;
	v1138 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v1138 != int64(0) {
		goto L259
	} else {
		goto L261
	}
L261:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(12288)
	goto L259
L262:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1147 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L14
	} else {
		goto L323
	}
L264:
	;
	F_relation_close(m, v1145, int32(3))
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L14
	} else {
		goto L322
	}
L265:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1147)+4))
	if v1150 <= int32(0) {
		goto L264
	} else {
		goto L266
	}
L266:
	;
	v1160 = v2
	goto L267
L267:
	;
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1147)+12))
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1180+v1160<<(uint(int32(2))%32))))
	v1185 = F_SearchSysCache1(m, int32(44), v1184)
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L14
	} else {
		goto L269
	}
L268:
	;
	goto L264
L269:
	;
	if v1185 == int32(0) {
		goto L263
	} else {
		goto L270
	}
L270:
	;
	v1191 = F_SysCacheGetAttrNotNull(m, int32(44), v1185, int32(2))
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L14
	} else {
		goto L271
	}
L271:
	;
	v1193 = F_text_to_cstring(m, v1191)
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L14
	} else {
		goto L272
	}
L272:
	;
	v1199 = F_SysCacheGetAttr(m, int32(44), v1185, int32(3), v29+int32(428))
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L14
	} else {
		goto L273
	}
L273:
	;
	v1201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+428)))
	if v1201 == int32(1) {
		goto L275
	} else {
		goto L276
	}
L274:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[0]))
	v1221 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	F_select_best_grantor(m, v1220, v1221, v1217, int32(10), v29+int32(240), v29+int32(432))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L14
	} else {
		goto L281
	}
L275:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1207 = F_acldefault(m, v1205, int32(10))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L14
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	v1211 = F_pg_detoast_datum_copy(m, v1199)
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L14
	} else {
		goto L279
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+496)) = int32(0)
	v1217 = v1207
	v1218 = int32(0)
	goto L274
L279:
	;
	v1215 = F_aclmembers(m, v1211, v29+int32(496))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L14
	} else {
		goto L280
	}
L280:
	;
	v1217 = v1211
	v1218 = v1215
	goto L274
L281:
	;
	v1229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v1230 = *(*int64)(unsafe.Add(mBase, uint32(v29)+432))
	v1231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v1232 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v29)+240))
	v1235 = int32(0)
	v1237 = F_restrict_and_check_grant(m, v1229, v1230, v1231, v1232, v1184, v1233, int32(27), v1193, v1235, v1235)
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L14
	} else {
		goto L282
	}
L282:
	;
	v1239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v1240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v29)+240))
	v1245 = F_merge_acl_with_grant(m, v1217, v1239, v1240, v1241, v1242, v1237, v1243, int32(10))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L14
	} else {
		goto L283
	}
L283:
	;
	v1249 = F_aclmembers(m, v1245, v29+int32(508))
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L14
	} else {
		goto L284
	}
L284:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1253 = F_acldefault(m, v1251, int32(10))
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L14
	} else {
		goto L286
	}
L285:
	;
	v1330 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[1])))
	if v1330 == int32(0) {
		goto L312
	} else {
		goto L313
	}
L286:
	;
	v1255 = int32(0)
	if v1245 != 0 {
		goto L289
	} else {
		goto L290
	}
L287:
	;
	if v1299 != 0 {
		goto L305
	} else {
		goto L306
	}
L288:
	;
	if v1253 == int32(0) {
		v1295 = v1255
		goto L296
	} else {
		goto L297
	}
L289:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+16))
	if v1257 != 0 {
		goto L288
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	if v1253 == int32(0) {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	goto L291
L293:
	;
	v1299 = int32(1)
	goto L287
L294:
	;
	goto L295
L295:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1253)+16))
	v1299 = base.B2i32(v1262 == int32(0))
	goto L287
L296:
	;
	v1299 = v1295
	goto L287
L297:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1253)+16))
	if v1257 != v1267 {
		v1295 = v1255
		goto L296
	} else {
		goto L298
	}
L298:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+8))
	if v1269 != 0 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1277 = v1269
	goto L301
L300:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+4))
	v1277 = (v1270<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L301
L301:
	;
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v1253)+8))
	if v1279 != 0 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1287 = v1279
	goto L304
L303:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1253)+4))
	v1287 = (v1280<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L304
L304:
	;
	v1291 = F_memcmp(m, v1277+v1245, v1287+v1253, v1257<<(uint(int32(4))%32))
	mBase = m.M
	v1295 = base.B2i32(v1291 == int32(0))
	goto L296
L305:
	;
	F_simple_heap_delete(m, v1145, v1185+int32(4))
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L14
	} else {
		goto L308
	}
L306:
	;
	goto L307
L307:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+288)) = int64(0)
	v1306 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+490)) = uint8(v1306)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+488)) = uint16(v1306)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+296)) = v1245
	v1311 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+486)) = uint8(v1311)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+484)) = uint16(v1306)
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+52))
	v1322 = F_heap_modify_tuple(m, v1185, v1315, v29+int32(288), v29+int32(488), v29+int32(484))
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L14
	} else {
		goto L309
	}
L308:
	;
	goto L285
L309:
	;
	F_CatalogTupleUpdate(m, v1145, v1322+int32(4), v1322)
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L14
	} else {
		goto L310
	}
L310:
	;
	goto L285
L311:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v29)+496))
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v29)+508))
	F_updateAclDependencies(m, int32(_a_F_ExecGrantStmt_oids_24), v1184, int32(0), int32(10), v1218, v1346, v1249, v1347)
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L14
	} else {
		goto L317
	}
L312:
	;
	v1334 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[2])))
	if v1334&int32(1) == int32(0) {
		goto L311
	} else {
		goto L315
	}
L313:
	;
	goto L314
L314:
	;
	F_recordExtensionInitPrivWorker(m, v1184, int32(_a_F_ExecGrantStmt_oids_24), int32(0), v1245)
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L14
	} else {
		goto L316
	}
L315:
	;
	goto L314
L316:
	;
	goto L311
L317:
	;
	F_ReleaseCatCache(m, v1185)
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L14
	} else {
		goto L318
	}
L318:
	;
	F_pfree(m, v1245)
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L14
	} else {
		goto L319
	}
L319:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L14
	} else {
		goto L320
	}
L320:
	;
	v1357 = v1160 + int32(1)
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v1147)+4))
	if v1357 < v1358 {
		v1160 = v1357
		goto L267
	} else {
		goto L321
	}
L321:
	;
	goto L268
L322:
	;
	goto L1
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+224)) = v1184
	F_errmsg_internal(m, int32(_a_F_ExecGrantStmt_oids_25), v29+int32(224))
	mBase = m.M
	v1398 = m.ExcPending
	if v1398 != 0 {
		goto L14
	} else {
		goto L324
	}
L324:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(2456), int32(_a_F_ExecGrantStmt_oids_26))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L14
	} else {
		goto L325
	}
L325:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L326:
	;
	goto L1
L327:
	;
	goto L1
L328:
	;
	v1424 = F_table_open(m, int32(2995), int32(3))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L14
	} else {
		goto L331
	}
L329:
	;
	v1417 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v1417 != int64(0) {
		goto L328
	} else {
		goto L330
	}
L330:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(6)
	goto L328
L331:
	;
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1426 == int32(0) {
		goto L333
	} else {
		goto L334
	}
L332:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L14
	} else {
		goto L369
	}
L333:
	;
	F_relation_close(m, v1424, int32(3))
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L14
	} else {
		goto L368
	}
L334:
	;
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+4))
	if v1429 <= int32(0) {
		goto L333
	} else {
		goto L335
	}
L335:
	;
	v1443 = v2
	goto L336
L336:
	;
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+12))
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v1458+v1443<<(uint(int32(2))%32))))
	v1463 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v1463
	*(*int64)(unsafe.Add(mBase, uint32(v29)+240)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+430)) = uint8(v1463)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+428)) = uint16(v1463)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+494)) = uint8(v1463)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+492)) = uint16(v1463)
	v1476 = v29 + int32(432)
	F_ScanKeyInit(m, v1476, int32(1), int32(3), int32(184), v1462)
	mBase = m.M
	v1481 = m.ExcPending
	if v1481 != 0 {
		goto L14
	} else {
		goto L338
	}
L337:
	;
	goto L333
L338:
	;
	v1483 = int32(1)
	v1486 = F_systable_beginscan(m, v1424, int32(2996), v1483, int32(0), v1483, v1476)
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L14
	} else {
		goto L339
	}
L339:
	;
	v1488 = F_systable_getnext(m, v1486)
	mBase = m.M
	v1489 = m.ExcPending
	if v1489 != 0 {
		goto L14
	} else {
		goto L340
	}
L340:
	;
	if v1488 == int32(0) {
		goto L332
	} else {
		goto L341
	}
L341:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v1488)+16))
	v1493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1492)+22)))
	v1494 = v1492 + v1493
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v1494)+4))
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v1424)+52))
	v1500 = F_heap_getattr_2(m, v1488, int32(3), v1497, v29+int32(507))
	mBase = m.M
	v1501 = m.ExcPending
	if v1501 != 0 {
		goto L14
	} else {
		goto L342
	}
L342:
	;
	v1502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+507)))
	if v1502 == int32(1) {
		goto L344
	} else {
		goto L345
	}
L343:
	;
	v1520 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[0]))
	v1521 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	F_select_best_grantor(m, v1520, v1521, v1517, v1495, v29+int32(508), v29+int32(496))
	mBase = m.M
	v1527 = m.ExcPending
	if v1527 != 0 {
		goto L14
	} else {
		goto L350
	}
L344:
	;
	v1507 = F_acldefault(m, int32(22), v1495)
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L14
	} else {
		goto L347
	}
L345:
	;
	goto L346
L346:
	;
	v1511 = F_pg_detoast_datum_copy(m, v1500)
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L14
	} else {
		goto L348
	}
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+488)) = int32(0)
	v1517 = v1507
	v1518 = int32(0)
	goto L343
L348:
	;
	v1515 = F_aclmembers(m, v1511, v29+int32(488))
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L14
	} else {
		goto L349
	}
L349:
	;
	v1517 = v1511
	v1518 = v1515
	goto L343
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+208)) = v1462
	v1530 = v29 + int32(288)
	v1535 = F_pg_snprintf(m, v1530, int32(64), int32(_a_F_ExecGrantStmt_oids_27), v29+int32(208))
	mBase = m.M
	v1536 = m.ExcPending
	if v1536 != 0 {
		goto L14
	} else {
		goto L351
	}
L351:
	;
	v1537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v1538 = *(*int64)(unsafe.Add(mBase, uint32(v29)+496))
	v1539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v1540 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v29)+508))
	v1543 = int32(0)
	v1545 = F_restrict_and_check_grant(m, v1537, v1538, v1539, v1540, v1462, v1541, int32(22), v1530, v1543, v1543)
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L14
	} else {
		goto L352
	}
L352:
	;
	v1547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v1548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v29)+508))
	v1552 = F_merge_acl_with_grant(m, v1517, v1547, v1548, v1549, v1550, v1545, v1551, v1495)
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L14
	} else {
		goto L353
	}
L353:
	;
	v1556 = F_aclmembers(m, v1552, v29+int32(484))
	mBase = m.M
	v1557 = m.ExcPending
	if v1557 != 0 {
		goto L14
	} else {
		goto L354
	}
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v1552
	v1559 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+494)) = uint8(v1559)
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v1424)+52))
	v1568 = F_heap_modify_tuple(m, v1488, v1561, v29+int32(240), v29+int32(428), v29+int32(492))
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L14
	} else {
		goto L355
	}
L355:
	;
	F_CatalogTupleUpdate(m, v1424, v1568+int32(4), v1568)
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L14
	} else {
		goto L356
	}
L356:
	;
	v1575 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[1])))
	if v1575 == int32(0) {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v1494)))
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v29)+488))
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(v29)+484))
	F_updateAclDependencies(m, int32(2613), v1589, int32(0), v1495, v1518, v1591, v1556, v1592)
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L14
	} else {
		goto L363
	}
L358:
	;
	v1579 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[2])))
	if v1579&int32(1) == int32(0) {
		goto L357
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	F_recordExtensionInitPrivWorker(m, v1462, int32(2613), int32(0), v1552)
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L14
	} else {
		goto L362
	}
L361:
	;
	goto L360
L362:
	;
	goto L357
L363:
	;
	F_systable_endscan(m, v1486)
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L14
	} else {
		goto L364
	}
L364:
	;
	F_pfree(m, v1552)
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L14
	} else {
		goto L365
	}
L365:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L14
	} else {
		goto L366
	}
L366:
	;
	v1602 = v1443 + int32(1)
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+4))
	if v1602 < v1603 {
		v1443 = v1602
		goto L336
	} else {
		goto L367
	}
L367:
	;
	goto L337
L368:
	;
	goto L1
L369:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+192)) = v1462
	F_errmsg_internal(m, int32(_a_F_ExecGrantStmt_oids_28), v29+int32(192))
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		goto L14
	} else {
		goto L370
	}
L370:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(2316), int32(_a_F_ExecGrantStmt_oids_29))
	mBase = m.M
	v1648 = m.ExcPending
	if v1648 != 0 {
		goto L14
	} else {
		goto L371
	}
L371:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L372:
	;
	goto L1
L373:
	;
	goto L1
L374:
	;
	goto L1
L375:
	;
	goto L1
L376:
	;
	goto L1
L377:
	;
	if (base.I32_wrap_i64(int64(base.Ui64(int64(4389322341887))>>(uint(base.I64_extend_i32_u(v1700))%64)))|base.B2i32(base.Ui32(int32(42)) < base.Ui32(v1700)))&int32(1) != 0 {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	v1710 = int32(0)
	v1712 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[3]))
	if v1712 == v1710 {
		goto L381
	} else {
		goto L382
	}
L379:
	;
	goto L380
L380:
	;
	m.G0 = v29 + int32(512)
	return
L381:
	;
	goto L380
L382:
	;
	v1715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1712)+20)))
	if v1715 != 0 {
		goto L381
	} else {
		goto L383
	}
L383:
	;
	v1716 = int32(_a_F_ExecGrantStmt_oids_30)
	v1717 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[4]))
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1712)))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[4])) = v1719
	v1722 = F_palloc(m, int32(40))
	mBase = m.M
	v1723 = m.ExcPending
	if v1723 != 0 {
		goto L14
	} else {
		goto L384
	}
L384:
	;
	v1724 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1722)+32)) = v1724
	v1726 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1722)+24)) = v1726
	v1728 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1722)+16)) = v1728
	v1730 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1722)+8)) = v1730
	v1732 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v1722))) = v1732
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1735 = F_list_copy(m, v1734)
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L14
	} else {
		goto L385
	}
L385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+8)) = v1735
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1739 = F_list_copy(m, v1738)
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L14
	} else {
		goto L386
	}
L386:
	;
	v1741 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+24)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+28)) = v1739
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1744 == v1741 {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v1818 = F_palloc(m, int32(40))
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L14
	} else {
		goto L395
	}
L388:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(v1744)+4))
	if v1747 <= int32(0) {
		goto L387
	} else {
		goto L389
	}
L389:
	;
	v1751 = int32(0)
	v1753 = v1710
	goto L390
L390:
	;
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(v1744)+12))
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(v1777+v1751<<(uint(int32(2))%32))))
	v1782 = F_copyObjectImpl(m, v1781)
	mBase = m.M
	v1783 = m.ExcPending
	if v1783 != 0 {
		goto L14
	} else {
		goto L392
	}
L391:
	;
	goto L387
L392:
	;
	v1784 = F_lappend(m, v1753, v1782)
	mBase = m.M
	v1785 = m.ExcPending
	if v1785 != 0 {
		goto L14
	} else {
		goto L393
	}
L393:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1722)+24)) = v1784
	v1788 = v1751 + int32(1)
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(v1744)+4))
	if v1788 < v1789 {
		v1751 = v1788
		v1753 = v1784
		goto L390
	} else {
		goto L394
	}
L394:
	;
	goto L391
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1818))) = int32(2)
	v1823 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[1])))
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+12)) = v1722
	*(*uint8)(unsafe.Add(mBase, uint32(v1818)+4)) = uint8(v1823)
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+8)) = int32(0)
	v1829 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[3]))
	v1830 = *(*int32)(unsafe.Add(mBase, uint32(v1829)+28))
	v1831 = F_lappend(m, v1830, v1818)
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L14
	} else {
		goto L396
	}
L396:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v1834)+28)) = v1831
	*(*int32)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[4])) = v1717
	goto L381
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v197 int32
	_ = v197
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v202 int64
	_ = v202
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v233 int64
	_ = v233
	var v236 int64
	_ = v236
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int64
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int64
	_ = v257
	var v271 int64
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int64
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v321 int32
	_ = v321
	var v322 int64
	_ = v322
	var v324 int64
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int64
	_ = v333
	var v339 int64
	_ = v339
	var v343 int32
	_ = v343
	var v353 int32
	_ = v353
	var v356 int64
	_ = v356
	var v360 int64
	_ = v360
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int64
	_ = v377
	var v378 int64
	_ = v378
	var v381 int64
	_ = v381
	var v384 int64
	_ = v384
	var v385 int64
	_ = v385
	var v388 int64
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int64
	_ = v397
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v410 int64
	_ = v410
	var v411 int64
	_ = v411
	var v413 int32
	_ = v413
	var v414 int64
	_ = v414
	var v416 int64
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int64
	_ = v425
	var v431 int64
	_ = v431
	var v435 int32
	_ = v435
	var v445 int32
	_ = v445
	var v448 int64
	_ = v448
	var v452 int64
	_ = v452
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int64
	_ = v469
	var v470 int64
	_ = v470
	var v473 int64
	_ = v473
	var v476 int64
	_ = v476
	var v477 int64
	_ = v477
	var v480 int64
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v496 int64
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v512 int64
	_ = v512
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int64
	_ = v549
	var v550 int64
	_ = v550
	var v551 int64
	_ = v551
	var v553 int64
	_ = v553
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v583 int64
	_ = v583
	var v588 int32
	_ = v588
	var v589 int64
	_ = v589
	var v590 int64
	_ = v590
	var v593 int64
	_ = v593
	var v596 int64
	_ = v596
	var v597 int64
	_ = v597
	var v600 int64
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int64
	_ = v624
	var v629 int32
	_ = v629
	var v630 int64
	_ = v630
	var v631 int64
	_ = v631
	var v634 int64
	_ = v634
	var v637 int64
	_ = v637
	var v638 int64
	_ = v638
	var v641 int64
	_ = v641
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int64
	_ = v655
	var v656 int64
	_ = v656
	var v657 int64
	_ = v657
	var v659 int64
	_ = v659
	var v660 int64
	_ = v660
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v679 int64
	_ = v679
	var v681 int32
	_ = v681
	var v697 int64
	_ = v697
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int64
	_ = v746
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v759 int64
	_ = v759
	var v760 int64
	_ = v760
	var v762 int32
	_ = v762
	var v763 int64
	_ = v763
	var v765 int64
	_ = v765
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int64
	_ = v774
	var v780 int64
	_ = v780
	var v784 int32
	_ = v784
	var v794 int32
	_ = v794
	var v797 int64
	_ = v797
	var v801 int64
	_ = v801
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int64
	_ = v820
	var v821 int64
	_ = v821
	var v825 int32
	_ = v825
	var v826 int64
	_ = v826
	var v830 int32
	_ = v830
	var v831 int64
	_ = v831
	var v832 int64
	_ = v832
	var v836 int32
	_ = v836
	var v837 int64
	_ = v837
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v849 int64
	_ = v849
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v862 int64
	_ = v862
	var v863 int64
	_ = v863
	var v865 int32
	_ = v865
	var v866 int64
	_ = v866
	var v868 int64
	_ = v868
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v877 int64
	_ = v877
	var v883 int64
	_ = v883
	var v887 int32
	_ = v887
	var v897 int32
	_ = v897
	var v900 int64
	_ = v900
	var v904 int64
	_ = v904
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v921 int64
	_ = v921
	var v922 int64
	_ = v922
	var v925 int64
	_ = v925
	var v928 int64
	_ = v928
	var v929 int64
	_ = v929
	var v932 int64
	_ = v932
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v946 int32
	_ = v946
	var v949 int64
	_ = v949
	var v950 int64
	_ = v950
	var v951 int64
	_ = v951
	var v953 int64
	_ = v953
	var v959 int32
	_ = v959
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1029 int32
	_ = v1029
	var v1034 int32
	_ = v1034
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_ExecIncrementalSort[0]))
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
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L4
	} else {
		goto L281
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L4
	} else {
		goto L278
	}
L8:
	;
	m.G0 = v18 + int32(48)
	return v984
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+56))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v59 != 0 {
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
		v984 = v42
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)))
	if v46 != 0 {
		v984 = v42
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
	v971 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v971 != int32(2) {
		goto L274
	} else {
		goto L275
	}
L22:
	;
	v670 = v27
	v679 = int64(0)
	v681 = v59
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
	if v681 != int32(1) {
		v959 = v670
		goto L21
	} else {
		goto L194
	}
L25:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v197 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L26:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+96))
	v66 = F_palloc(m, v63*int32(36))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
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
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L43
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v66
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v62)+96))
	if int32(0) < v69 {
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
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_ExecIncrementalSort[1]))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	v176 = F_tuplesort_begin_heap(m, v58, v163, v164, v165, v166, v167, v169, int32(0), v171<<(uint(int32(1))%32)&int32(2))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L42
	}
L33:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v91 = v88 + v77*int32(36)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v62)+76))
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92+v77<<(uint(int32(1))%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v91)+32)) = uint16(v96)
	v99 = v77 << (uint(int32(2)) % 32)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v62)+80))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99+v100)))
	v104 = F_get_equality_op_for_ordering_op(m, v102, int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L35
	}
L34:
	;
	goto L32
L35:
	;
	if v104 == int32(0) {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	v108 = F_get_opcode(m, v104)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	if v108 == int32(0) {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_ExecIncrementalSort[2]))
	F_fmgr_info_cxt(m, v108, v91, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v117 = F_palloc0(m, int32(36))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = v91
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v122 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = v122
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+8)) = v122
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v62)+84))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v128+v99)))
	*(*int32)(unsafe.Add(mBase, uint32(v127)+12)) = v130
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v132)+16)) = uint8(v122)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v136 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v135)+18)) = uint16(v136)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v138)+24)) = uint8(v122)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+32)) = uint8(v122)
	v145 = v77 + int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v62)+96))
	if v145 < v146 {
		v77 = v145
		goto L33
	} else {
		goto L41
	}
L41:
	;
	goto L34
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v176
	v185 = v176
	goto L25
L43:
	;
	v185 = v27
	goto L25
L44:
	;
	v200 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v201 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	v202 = v200 - v201
	if v202 <= int64(31) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v238 = int64(32)
	goto L46
L46:
	;
	v239 = int64(0)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v240 == int32(0) {
		v257 = v239
		goto L65
	} else {
		goto L66
	}
L47:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v185)+236))
	if v207 != 0 {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	goto L49
L49:
	;
	v233 = int64(32)
	if v233 <= v202 {
		goto L62
	} else {
		goto L63
	}
L50:
	;
	goto L49
L51:
	;
	goto L50
L52:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v185)+72)) = uint32(v202)
	v216 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v185)+68)) = uint8(v216)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v185)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v218)+24)) = int32(0)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v185)+44))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+32))
	if v222 != 0 {
		goto L59
	} else {
		goto L60
	}
L53:
	;
	if int64(1073741823) < v202 {
		goto L51
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if int64(1073741823) < v202 {
		goto L51
	} else {
		goto L58
	}
L56:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v185)+232))
	if v210 != int32(-1) {
		goto L52
	} else {
		goto L57
	}
L57:
	;
	goto L51
L58:
	;
	goto L52
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v221)+16)) = v222
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v185)+44))
	v225 = v224
	goto L61
L60:
	;
	v225 = v221
	goto L61
L61:
	;
	v226 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v225)+28)) = v226
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v185)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+32)) = v226
	goto L51
L62:
	;
	v236 = v233
	goto L64
L63:
	;
	v236 = v202
	goto L64
L64:
	;
	v238 = v236
	goto L46
L65:
	;
	v271 = v257
	goto L75
L66:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+4)))
	if v243&int32(2) != 0 {
		v257 = v239
		goto L65
	} else {
		goto L67
	}
L67:
	;
	F_tuplesort_puttupleslot(m, v185, v240)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v248 = int64(1)
	if v238 == v248 {
		v257 = v248
		goto L65
	} else {
		goto L69
	}
L69:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+8))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
	m.T0[v253].(func(*base.Module, int32))(m, v251)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v257 = v248
	goto L65
L71:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653)+69)))
	if v654 != 0 {
		goto L187
	} else {
		goto L188
	}
L72:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v620 = m.G0
	v622 = v620 - int32(16)
	m.G0 = v622
	v624 = *(*int64)(unsafe.Add(mBase, uint32(v616)))
	*(*int64)(unsafe.Add(mBase, uint32(v616))) = v624 + int64(1)
	F_tuplesort_get_stats(m, v617, v622)
	mBase = m.M
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v622)+4))
	switch v629 {
	case 0:
		goto L184
	case 1:
		goto L183
	default:
		goto L182
	}
L73:
	;
	v616 = l0 + int32(176)
	goto L72
L74:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)+8))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)+32))
	m.T0[v543].(func(*base.Module, int32, int32))(m, v541, v277)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L4
	} else {
		goto L160
	}
L75:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v57)+52))
	if v273 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v518)+8))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v519)+12))
	m.T0[v520].(func(*base.Module, int32))(m, v518)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L4
	} else {
		goto L155
	}
L77:
	;
	F_ExecReScan(m, v57)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v277 = m.T0[v276].(func(*base.Module, int32) int32)(m, v57)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L4
	} else {
		goto L82
	}
L80:
	;
	goto L79
L81:
	;
	if v271 < v238 {
		goto L144
	} else {
		goto L145
	}
L82:
	;
	if v277 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277)+4)))
	if v279&int32(2) == int32(0) {
		goto L81
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v284 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)) = uint8(v284)
	F_tuplesort_performsort(m, v185)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L87
	}
L86:
	;
	goto L85
L87:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v288 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(2)
	v959 = v185
	goto L21
L89:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v291 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v397 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v397 + int64(1)
	v402 = v18 + int32(32)
	v403 = int32(0)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v396)+128))
	if v407 == v403 {
		goto L121
	} else {
		goto L122
	}
L91:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+280)))
	if v294 != int32(1) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_ExecIncrementalSort[3]))
	v304 = v291 + v299*int32(96) + int32(8)
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v304)))
	*(*int64)(unsafe.Add(mBase, uint32(v304))) = v305 + int64(1)
	v310 = v18 + int32(32)
	v311 = int32(0)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v297)+128))
	if v315 == v311 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	switch v376 {
	case 0:
		goto L115
	case 1:
		goto L114
	default:
		goto L113
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v310)+4)) = v353
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v297)+112))
	v360 = base.I64_div_s(v356+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(v310)+8)) = v360
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v297)+124))
	switch v362 - int32(3) {
	case 0:
		goto L109
	case 1:
		v373 = v362
		goto L106
	case 2:
		goto L108
	default:
		goto L107
	}
L95:
	;
	if v332&int32(255) != base.B2i32(v315 != int32(0)) {
		goto L101
	} else {
		goto L102
	}
L96:
	;
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v297)+96))
	v319 = *(*int64)(unsafe.Add(mBase, uint32(v297)+88))
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+120)))
	v332 = v321
	v333 = v318 - v319
	goto L95
L97:
	;
	goto L98
L98:
	;
	v322 = F_LogicalTapeSetBlocks(m, v315)
	mBase = m.M
	v324 = v322 << (uint(int64(13)) % 64)
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+120)))
	if v326 != 0 {
		v332 = int32(1)
		v333 = v324
		goto L95
	} else {
		goto L99
	}
L99:
	;
	v327 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v297)+120)) = uint8(v327)
	*(*int64)(unsafe.Add(mBase, uint32(v297)+112)) = v324
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v297)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v297)+124)) = v330
	v353 = v311
	goto L94
L100:
	;
	v353 = int32(1)
	goto L94
L101:
	;
	if v332&int32(1) != 0 {
		v353 = v311
		goto L94
	} else {
		goto L105
	}
L102:
	;
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v297)+112))
	if v333 <= v339 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v297)+120)) = uint8(v332)
	*(*int64)(unsafe.Add(mBase, uint32(v297)+112)) = v333
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v297)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v297)+124)) = v343
	if v332&int32(1) == int32(0) {
		goto L100
	} else {
		goto L104
	}
L104:
	;
	v353 = v311
	goto L94
L105:
	;
	goto L100
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v310))) = v373
	goto L93
L107:
	;
	v373 = int32(0)
	goto L106
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v310))) = int32(8)
	goto L93
L109:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+69)))
	if v367 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v368 = int32(1)
	goto L112
L111:
	;
	v368 = int32(2)
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v310))) = v368
	goto L93
L113:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v304)+40))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v304)+40)) = v392 | v393
	goto L88
L114:
	;
	v384 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	v385 = *(*int64)(unsafe.Add(mBase, uint32(v304)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v304)+32)) = v384 + v385
	v388 = *(*int64)(unsafe.Add(mBase, uint32(v304)+24))
	if v384 <= v388 {
		goto L113
	} else {
		goto L117
	}
L115:
	;
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	v378 = *(*int64)(unsafe.Add(mBase, uint32(v304)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v304)+16)) = v377 + v378
	v381 = *(*int64)(unsafe.Add(mBase, uint32(v304)+8))
	if v377 <= v381 {
		goto L113
	} else {
		goto L116
	}
L116:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v304)+8)) = v377
	goto L113
L117:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v304)+24)) = v384
	goto L113
L118:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	switch v468 {
	case 0:
		goto L140
	case 1:
		goto L139
	default:
		goto L138
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v402)+4)) = v445
	v448 = *(*int64)(unsafe.Add(mBase, uint32(v396)+112))
	v452 = base.I64_div_s(v448+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(v402)+8)) = v452
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v396)+124))
	switch v454 - int32(3) {
	case 0:
		goto L134
	case 1:
		v465 = v454
		goto L131
	case 2:
		goto L133
	default:
		goto L132
	}
L120:
	;
	if v424&int32(255) != base.B2i32(v407 != int32(0)) {
		goto L126
	} else {
		goto L127
	}
L121:
	;
	v410 = *(*int64)(unsafe.Add(mBase, uint32(v396)+96))
	v411 = *(*int64)(unsafe.Add(mBase, uint32(v396)+88))
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+120)))
	v424 = v413
	v425 = v410 - v411
	goto L120
L122:
	;
	goto L123
L123:
	;
	v414 = F_LogicalTapeSetBlocks(m, v407)
	mBase = m.M
	v416 = v414 << (uint(int64(13)) % 64)
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+120)))
	if v418 != 0 {
		v424 = int32(1)
		v425 = v416
		goto L120
	} else {
		goto L124
	}
L124:
	;
	v419 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v396)+120)) = uint8(v419)
	*(*int64)(unsafe.Add(mBase, uint32(v396)+112)) = v416
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v396)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v396)+124)) = v422
	v445 = v403
	goto L119
L125:
	;
	v445 = int32(1)
	goto L119
L126:
	;
	if v424&int32(1) != 0 {
		v445 = v403
		goto L119
	} else {
		goto L130
	}
L127:
	;
	v431 = *(*int64)(unsafe.Add(mBase, uint32(v396)+112))
	if v425 <= v431 {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v396)+120)) = uint8(v424)
	*(*int64)(unsafe.Add(mBase, uint32(v396)+112)) = v425
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v396)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v396)+124)) = v435
	if v424&int32(1) == int32(0) {
		goto L125
	} else {
		goto L129
	}
L129:
	;
	v445 = v403
	goto L119
L130:
	;
	goto L125
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v402))) = v465
	goto L118
L132:
	;
	v465 = int32(0)
	goto L131
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v402))) = int32(8)
	goto L118
L134:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+69)))
	if v459 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v460 = int32(1)
	goto L137
L136:
	;
	v460 = int32(2)
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v402))) = v460
	goto L118
L138:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v484 | v485
	goto L88
L139:
	;
	v476 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	v477 = *(*int64)(unsafe.Add(mBase, uint32(l0)+208))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v476 + v477
	v480 = *(*int64)(unsafe.Add(mBase, uint32(l0)+200))
	if v476 <= v480 {
		goto L138
	} else {
		goto L142
	}
L140:
	;
	v469 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	v470 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v469 + v470
	v473 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	if v469 <= v473 {
		goto L138
	} else {
		goto L141
	}
L141:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v469
	goto L138
L142:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+200)) = v476
	goto L138
L143:
	;
	if v512 < int64(65) {
		v271 = v512
		goto L75
	} else {
		goto L153
	}
L144:
	;
	F_tuplesort_puttupleslot(m, v185, v277)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L4
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v504 = F_isCurrentGroup(m, l0, v503, v277)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L4
	} else {
		goto L150
	}
L147:
	;
	v496 = v271 + int64(1)
	if v496 != v238 {
		v512 = v496
		goto L143
	} else {
		goto L148
	}
L148:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)+8))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v499)+32))
	m.T0[v500].(func(*base.Module, int32, int32))(m, v498, v277)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	v271 = v238
	goto L75
L150:
	;
	if v504 == int32(0) {
		goto L74
	} else {
		goto L151
	}
L151:
	;
	F_tuplesort_puttupleslot(m, v185, v277)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	v512 = v271 + int64(1)
	goto L143
L153:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v515 == int32(2) {
		v271 = v512
		goto L75
	} else {
		goto L154
	}
L154:
	;
	goto L76
L155:
	;
	F_tuplesort_performsort(m, v185)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v525 == int32(0) {
		goto L71
	} else {
		goto L157
	}
L157:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v528 == int32(0) {
		goto L73
	} else {
		goto L158
	}
L158:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+280)))
	if v531 != int32(1) {
		goto L73
	} else {
		goto L159
	}
L159:
	;
	v535 = *(*int32)(unsafe.Add(mBase, _c_F_ExecIncrementalSort[3]))
	v616 = v528 + v535*int32(96) + int32(8)
	goto L72
L160:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v546 == int32(1) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v549 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v550 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	v551 = v550 + v271
	if v549 < v551 {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	goto L163
L163:
	;
	F_tuplesort_performsort(m, v185)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L4
	} else {
		goto L167
	}
L164:
	;
	v553 = v549
	goto L166
L165:
	;
	v553 = v551
	goto L166
L166:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v553
	goto L163
L167:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v559 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v560 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L169:
	;
	goto L170
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(2)
	v959 = v185
	goto L21
L171:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v579 = m.G0
	v581 = v579 - int32(16)
	m.G0 = v581
	v583 = *(*int64)(unsafe.Add(mBase, uint32(v575)))
	*(*int64)(unsafe.Add(mBase, uint32(v575))) = v583 + int64(1)
	F_tuplesort_get_stats(m, v576, v581)
	mBase = m.M
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v581)+4))
	switch v588 {
	case 0:
		goto L178
	case 1:
		goto L177
	default:
		goto L176
	}
L172:
	;
	v575 = l0 + int32(176)
	goto L171
L173:
	;
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+280)))
	if v563 != int32(1) {
		goto L172
	} else {
		goto L174
	}
L174:
	;
	v567 = *(*int32)(unsafe.Add(mBase, _c_F_ExecIncrementalSort[3]))
	v575 = v560 + v567*int32(96) + int32(8)
	goto L171
L175:
	;
	goto L170
L176:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v575)+40))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v581)))
	*(*int32)(unsafe.Add(mBase, uint32(v575)+40)) = v604 | v605
	m.G0 = v581 + int32(16)
	goto L175
L177:
	;
	v596 = *(*int64)(unsafe.Add(mBase, uint32(v581)+8))
	v597 = *(*int64)(unsafe.Add(mBase, uint32(v575)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v575)+32)) = v596 + v597
	v600 = *(*int64)(unsafe.Add(mBase, uint32(v575)+24))
	if v596 <= v600 {
		goto L176
	} else {
		goto L180
	}
L178:
	;
	v589 = *(*int64)(unsafe.Add(mBase, uint32(v581)+8))
	v590 = *(*int64)(unsafe.Add(mBase, uint32(v575)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v575)+16)) = v589 + v590
	v593 = *(*int64)(unsafe.Add(mBase, uint32(v575)+8))
	if v589 <= v593 {
		goto L176
	} else {
		goto L179
	}
L179:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v575)+8)) = v589
	goto L176
L180:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v575)+24)) = v596
	goto L176
L181:
	;
	goto L71
L182:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v616)+40))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v622)))
	*(*int32)(unsafe.Add(mBase, uint32(v616)+40)) = v645 | v646
	m.G0 = v622 + int32(16)
	goto L181
L183:
	;
	v637 = *(*int64)(unsafe.Add(mBase, uint32(v622)+8))
	v638 = *(*int64)(unsafe.Add(mBase, uint32(v616)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v616)+32)) = v637 + v638
	v641 = *(*int64)(unsafe.Add(mBase, uint32(v616)+24))
	if v637 <= v641 {
		goto L182
	} else {
		goto L186
	}
L184:
	;
	v630 = *(*int64)(unsafe.Add(mBase, uint32(v622)+8))
	v631 = *(*int64)(unsafe.Add(mBase, uint32(v616)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v616)+16)) = v630 + v631
	v634 = *(*int64)(unsafe.Add(mBase, uint32(v616)+8))
	if v630 <= v634 {
		goto L182
	} else {
		goto L185
	}
L185:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v616)+8)) = v630
	goto L182
L186:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v616)+24)) = v637
	goto L182
L187:
	;
	v655 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v656 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	v657 = v655 - v656
	if v657 < v512 {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	v660 = v512
	goto L189
L189:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v660
	F_switchToPresortedPrefixMode(m, l0)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L4
	} else {
		goto L193
	}
L190:
	;
	v659 = v657
	goto L192
L191:
	;
	v659 = v512
	goto L192
L192:
	;
	v660 = v659
	goto L189
L193:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v670 = v185
	v679 = v660
	v681 = v665
	goto L24
L194:
	;
	v697 = v679
	goto L196
L195:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	F_tuplesort_performsort(m, v726)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L4
	} else {
		goto L214
	}
L196:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v57)+52))
	if v699 != 0 {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v720)+8))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v721)+32))
	m.T0[v722].(func(*base.Module, int32, int32))(m, v720, v703)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L4
	} else {
		goto L213
	}
L198:
	;
	F_ExecReScan(m, v57)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L4
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v703 = m.T0[v702].(func(*base.Module, int32) int32)(m, v57)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L4
	} else {
		goto L203
	}
L201:
	;
	goto L200
L202:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v713 = F_isCurrentGroup(m, l0, v712, v703)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L4
	} else {
		goto L208
	}
L203:
	;
	if v703 != 0 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703)+4)))
	if v705&int32(2) == int32(0) {
		goto L202
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	v710 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)) = uint8(v710)
	goto L195
L207:
	;
	goto L206
L208:
	;
	if v713 != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	F_tuplesort_puttupleslot(m, v715, v703)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L4
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	goto L197
L212:
	;
	v697 = v697 + int64(1)
	goto L196
L213:
	;
	goto L195
L214:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v729 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(3)
	v946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v946 != int32(1) {
		v959 = v670
		goto L21
	} else {
		goto L270
	}
L216:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v732 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v849 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+224)) = v849 + int64(1)
	v854 = v18 + int32(32)
	v855 = int32(0)
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v848)+128))
	if v859 == v855 {
		goto L248
	} else {
		goto L249
	}
L218:
	;
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+280)))
	if v735 != int32(1) {
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v740 = *(*int32)(unsafe.Add(mBase, _c_F_ExecIncrementalSort[3]))
	v743 = v732 + v740*int32(96)
	v745 = v743 + int32(56)
	v746 = *(*int64)(unsafe.Add(mBase, uint32(v745)))
	*(*int64)(unsafe.Add(mBase, uint32(v745))) = v746 + int64(1)
	v751 = v18 + int32(32)
	v752 = int32(0)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v738)+128))
	if v756 == v752 {
		goto L223
	} else {
		goto L224
	}
L220:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	switch v817 {
	case 0:
		goto L242
	case 1:
		goto L241
	default:
		goto L240
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v751)+4)) = v794
	v797 = *(*int64)(unsafe.Add(mBase, uint32(v738)+112))
	v801 = base.I64_div_s(v797+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(v751)+8)) = v801
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v738)+124))
	switch v803 - int32(3) {
	case 0:
		goto L236
	case 1:
		v814 = v803
		goto L233
	case 2:
		goto L235
	default:
		goto L234
	}
L222:
	;
	if v773&int32(255) != base.B2i32(v756 != int32(0)) {
		goto L228
	} else {
		goto L229
	}
L223:
	;
	v759 = *(*int64)(unsafe.Add(mBase, uint32(v738)+96))
	v760 = *(*int64)(unsafe.Add(mBase, uint32(v738)+88))
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v738)+120)))
	v773 = v762
	v774 = v759 - v760
	goto L222
L224:
	;
	goto L225
L225:
	;
	v763 = F_LogicalTapeSetBlocks(m, v756)
	mBase = m.M
	v765 = v763 << (uint(int64(13)) % 64)
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v738)+120)))
	if v767 != 0 {
		v773 = int32(1)
		v774 = v765
		goto L222
	} else {
		goto L226
	}
L226:
	;
	v768 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v738)+120)) = uint8(v768)
	*(*int64)(unsafe.Add(mBase, uint32(v738)+112)) = v765
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v738)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v738)+124)) = v771
	v794 = v752
	goto L221
L227:
	;
	v794 = int32(1)
	goto L221
L228:
	;
	if v773&int32(1) != 0 {
		v794 = v752
		goto L221
	} else {
		goto L232
	}
L229:
	;
	v780 = *(*int64)(unsafe.Add(mBase, uint32(v738)+112))
	if v774 <= v780 {
		goto L228
	} else {
		goto L230
	}
L230:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v738)+120)) = uint8(v773)
	*(*int64)(unsafe.Add(mBase, uint32(v738)+112)) = v774
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v738)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v738)+124)) = v784
	if v773&int32(1) == int32(0) {
		goto L227
	} else {
		goto L231
	}
L231:
	;
	v794 = v752
	goto L221
L232:
	;
	goto L227
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v751))) = v814
	goto L220
L234:
	;
	v814 = int32(0)
	goto L233
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v751))) = int32(8)
	goto L220
L236:
	;
	v808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v738)+69)))
	if v808 != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v809 = int32(1)
	goto L239
L238:
	;
	v809 = int32(2)
	goto L239
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v751))) = v809
	goto L220
L240:
	;
	v843 = v743 + int32(96)
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v843)))
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v843))) = v844 | v845
	goto L215
L241:
	;
	v830 = v743 + int32(88)
	v831 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	v832 = *(*int64)(unsafe.Add(mBase, uint32(v830)))
	*(*int64)(unsafe.Add(mBase, uint32(v830))) = v831 + v832
	v836 = v743 + int32(80)
	v837 = *(*int64)(unsafe.Add(mBase, uint32(v836)))
	if v831 <= v837 {
		goto L240
	} else {
		goto L244
	}
L242:
	;
	v819 = v743 + int32(72)
	v820 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	v821 = *(*int64)(unsafe.Add(mBase, uint32(v819)))
	*(*int64)(unsafe.Add(mBase, uint32(v819))) = v820 + v821
	v825 = v743 - int32(-64)
	v826 = *(*int64)(unsafe.Add(mBase, uint32(v825)))
	if v820 <= v826 {
		goto L240
	} else {
		goto L243
	}
L243:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v825))) = v820
	goto L240
L244:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v836))) = v831
	goto L240
L245:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	switch v920 {
	case 0:
		goto L267
	case 1:
		goto L266
	default:
		goto L265
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v854)+4)) = v897
	v900 = *(*int64)(unsafe.Add(mBase, uint32(v848)+112))
	v904 = base.I64_div_s(v900+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(v854)+8)) = v904
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v848)+124))
	switch v906 - int32(3) {
	case 0:
		goto L261
	case 1:
		v917 = v906
		goto L258
	case 2:
		goto L260
	default:
		goto L259
	}
L247:
	;
	if v876&int32(255) != base.B2i32(v859 != int32(0)) {
		goto L253
	} else {
		goto L254
	}
L248:
	;
	v862 = *(*int64)(unsafe.Add(mBase, uint32(v848)+96))
	v863 = *(*int64)(unsafe.Add(mBase, uint32(v848)+88))
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v848)+120)))
	v876 = v865
	v877 = v862 - v863
	goto L247
L249:
	;
	goto L250
L250:
	;
	v866 = F_LogicalTapeSetBlocks(m, v859)
	mBase = m.M
	v868 = v866 << (uint(int64(13)) % 64)
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v848)+120)))
	if v870 != 0 {
		v876 = int32(1)
		v877 = v868
		goto L247
	} else {
		goto L251
	}
L251:
	;
	v871 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v848)+120)) = uint8(v871)
	*(*int64)(unsafe.Add(mBase, uint32(v848)+112)) = v868
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v848)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v848)+124)) = v874
	v897 = v855
	goto L246
L252:
	;
	v897 = int32(1)
	goto L246
L253:
	;
	if v876&int32(1) != 0 {
		v897 = v855
		goto L246
	} else {
		goto L257
	}
L254:
	;
	v883 = *(*int64)(unsafe.Add(mBase, uint32(v848)+112))
	if v877 <= v883 {
		goto L253
	} else {
		goto L255
	}
L255:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v848)+120)) = uint8(v876)
	*(*int64)(unsafe.Add(mBase, uint32(v848)+112)) = v877
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v848)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v848)+124)) = v887
	if v876&int32(1) == int32(0) {
		goto L252
	} else {
		goto L256
	}
L256:
	;
	v897 = v855
	goto L246
L257:
	;
	goto L252
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v854))) = v917
	goto L245
L259:
	;
	v917 = int32(0)
	goto L258
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v854))) = int32(8)
	goto L245
L261:
	;
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v848)+69)))
	if v911 != 0 {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v912 = int32(1)
	goto L264
L263:
	;
	v912 = int32(2)
	goto L264
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v854))) = v912
	goto L245
L265:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+264)) = v936 | v937
	goto L215
L266:
	;
	v928 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	v929 = *(*int64)(unsafe.Add(mBase, uint32(l0)+256))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+256)) = v928 + v929
	v932 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
	if v928 <= v932 {
		goto L265
	} else {
		goto L269
	}
L267:
	;
	v921 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	v922 = *(*int64)(unsafe.Add(mBase, uint32(l0)+240))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+240)) = v921 + v922
	v925 = *(*int64)(unsafe.Add(mBase, uint32(l0)+232))
	if v921 <= v925 {
		goto L265
	} else {
		goto L268
	}
L268:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+232)) = v921
	goto L265
L269:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v928
	goto L265
L270:
	;
	v949 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v950 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	v951 = v950 + v697
	if v949 < v951 {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v953 = v949
	goto L273
L272:
	;
	v953 = v951
	goto L273
L273:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v953
	v959 = v670
	goto L21
L274:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v975 = v974
	goto L276
L275:
	;
	v975 = v959
	goto L276
L276:
	;
	v978 = int32(0)
	v979 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v981 = F_tuplesort_gettupleslot(m, v975, base.B2i32(v29 == int32(1)), v978, v979, v978)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L4
	} else {
		goto L277
	}
L277:
	;
	v984 = v979
	goto L8
L278:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v62)+80))
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v1006+v77<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v1010
	F_errmsg_internal(m, int32(_a_F_ExecIncrementalSort_0), v18)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L4
	} else {
		goto L279
	}
L279:
	;
	F_errfinish(m, int32(_a_F_ExecIncrementalSort_1), int32(186), int32(_a_F_ExecIncrementalSort_2))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L4
	} else {
		goto L280
	}
L280:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v104
	F_errmsg_internal(m, int32(_a_F_ExecIncrementalSort_3), v18+int32(16))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L4
	} else {
		goto L282
	}
L282:
	;
	F_errfinish(m, int32(_a_F_ExecIncrementalSort_1), int32(190), int32(_a_F_ExecIncrementalSort_2))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L4
	} else {
		goto L283
	}
L283:
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
			v15 = F_MakeTupleTableSlot(m, v8, int32(_a_F_ExecInitJunkFilter_0))
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
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
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	v7 = l6
	v10 = int32(_a_F_ExecInitRoutingInfo_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInitRoutingInfo[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInitRoutingInfo[0])) = v13
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
	v23 = int32(0)
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+204)) = v23
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l4)+84))
	if v26 == int32(0) {
		v50 = int32(1)
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v23 = v20
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+208)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+104)) = v50
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v56 = v54 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v56 < v58 {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+76))
	if v29 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	m.T0[v29].(func(*base.Module, int32, int32))(m, l0, l4)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v36 = v26
	goto L11
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+60))
	if v38 == int32(0) {
		v50 = int32(1)
		goto L7
	} else {
		goto L14
	}
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l4)+84))
	if v33 == int32(0) {
		v50 = int32(1)
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v36 = v33
	goto L11
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+56))
	if v42 == int32(0) {
		v50 = int32(1)
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v45 = m.T0[v38].(func(*base.Module, int32) int32)(m, l4)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v50 = v45
	goto L7
L17:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v87 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v86+v54<<(uint(v87)%32)))) = l4
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v91+v54))) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(l3+l5<<(uint(v87)%32))+24)) = v54
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInitRoutingInfo[0])) = v11
	return
L18:
	;
	if v58 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = int32(8)
	v65 = F_palloc(m, int32(32))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v58 << (uint(int32(1)) % 32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v78 = F_repalloc(m, v75, v58<<(uint(int32(3))%32))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v65
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v69 = F_palloc(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v69
	goto L17
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v78
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v83 = F_repalloc(m, v81, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v83
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
	var v54 int64
	_ = v54
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v68 int64
	_ = v68
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v90 int64
	_ = v90
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
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
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
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
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int64
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v357 int64
	_ = v357
	var v360 int32
	_ = v360
	var v362 int64
	_ = v362
	var v364 int64
	_ = v364
	var v367 int64
	_ = v367
	var v368 int64
	_ = v368
	var v378 int64
	_ = v378
	var v383 int32
	_ = v383
	var v384 int64
	_ = v384
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int64
	_ = v393
	var v403 int64
	_ = v403
	var v411 int32
	_ = v411
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v484 int32
	_ = v484
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int64
	_ = v499
	var v501 int64
	_ = v501
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v614 int64
	_ = v614
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int64
	_ = v664
	var v666 int64
	_ = v666
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v690 int64
	_ = v690
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v739 int32
	_ = v739
	var v742 int64
	_ = v742
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int64
	_ = v760
	var v764 int64
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v778 int64
	_ = v778
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v791 int64
	_ = v791
	var v804 int64
	_ = v804
	var v807 int32
	_ = v807
	var v811 int64
	_ = v811
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int64
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v891 int32
	_ = v891
	var v899 int64
	_ = v899
	var v900 int64
	_ = v900
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v1004 int32
	_ = v1004
	var v1017 int64
	_ = v1017
	var v1023 int32
	_ = v1023
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int64
	_ = v1053
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1101 int32
	_ = v1101
	var v1106 int32
	_ = v1106
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1150 int32
	_ = v1150
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMemoize[0]))
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
		goto L17
	case 1:
		goto L16
	case 2:
		goto L15
	case 3:
		goto L14
	case 4:
		v1106 = v2
		goto L8
	default:
		goto L13
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L4
	} else {
		goto L270
	}
L8:
	;
	m.G0 = v18 + int32(16)
	return v1106
L9:
	;
	v342 = v339
	goto L100
L10:
	;
	v339 = int32(0)
	goto L9
L11:
	;
	v339 = int32(1)
	goto L9
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L4
	} else {
		goto L97
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L4
	} else {
		goto L94
	}
L14:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+52))
	if v292 != 0 {
		goto L83
	} else {
		goto L84
	}
L15:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+52))
	if v258 != 0 {
		goto L67
	} else {
		goto L68
	}
L16:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v245
	if v245 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L17:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v33 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
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
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	m.T0[v118].(func(*base.Module, int32))(m, v116)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L47
	}
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+28)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = v38
	if v37 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L7
L24:
	;
	v47 = v37
	goto L26
L25:
	;
	v47 = int32(1024)
	goto L26
L26:
	;
	v50 = base.F64_div(base.F64_convert_i32_u(v47), float64(0.9))
	if base.F64_ge(v50, float64(4.294967296e+09)) != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v53 = float64(4.294967296e+09)
	goto L29
L28:
	;
	v53 = v50
	goto L29
L29:
	;
	v54 = base.I64_trunc_sat_f64_u(v53)
	if base.Ui64(v54) <= base.Ui64(int64(2)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v57 = int64(2)
	goto L32
L31:
	;
	v57 = v54
	goto L32
L32:
	;
	v58 = int64(1)
	if v57&(v57-v58) == int64(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v68 = v57
	goto L35
L34:
	;
	v68 = v58 << (uint(int64(64)-base.I64_clz(v57)) % 64)
	goto L35
L35:
	;
	if base.Ui64(v68<<(uint(int64(4))%64)) < base.Ui64(int64(2147483647)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v77 = F_MemoryContextAllocExtended(m, v38, base.I32_wrap_i64(v68)<<(uint(int32(4))%32), int32(5))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L7
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v77
	v80 = int64(1)
	if v68&(v68-v80) == int64(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v90 = v68
	goto L42
L41:
	;
	v90 = v80 << (uint(int64(64)-base.I64_clz(v68)) % 64)
	goto L42
L42:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v90<<(uint(int64(4))%64)) {
		goto L23
	} else {
		goto L43
	}
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40))) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = base.I32_wrap_i64(v90) - int32(1)
	if v90 == int64(4294967296) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v107 = int32(-85899346)
	goto L46
L45:
	;
	v107 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v90), float64(0.9)))
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v40
	goto L21
L47:
	;
	v121 = int32(_a_F_ExecMemoize_0)
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMemoize[1]))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMemoize[1])) = v125
	if v115 <= int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMemoize[1])) = v122
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116)+4)))
	v233 = v231 & int32(_a_F_ExecMemoize_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v116)+4)) = uint16(v233)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)) = uint16(v236)
	goto L60
L49:
	;
	if v115 != int32(1) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v138 = v2
	v139 = int32(0)
	goto L53
L51:
	;
	v188 = v2
	goto L52
L52:
	;
	v202 = v188 << (uint(int32(2)) % 32)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v202+v203)))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v116)+20))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v205)+20))
	v209 = m.T0[v208].(func(*base.Module, int32, int32, int32) int32)(m, v205, v124, v206+v188)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L4
	} else {
		goto L59
	}
L53:
	;
	v152 = v138 << (uint(int32(2)) % 32)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v152+v153)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v116)+20))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v155)+20))
	v159 = m.T0[v158].(func(*base.Module, int32, int32, int32) int32)(m, v155, v124, v156+v138)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L55
	}
L54:
	;
	if v115&int32(1) == int32(0) {
		goto L48
	} else {
		goto L58
	}
L55:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v161+v152))) = v159
	v165 = v138 | int32(1)
	v167 = v165 << (uint(int32(2)) % 32)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v167+v168)))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v116)+20))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v170)+20))
	v174 = m.T0[v173].(func(*base.Module, int32, int32, int32) int32)(m, v170, v124, v171+v165)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v176+v167))) = v174
	v179 = int32(2)
	v180 = v138 + v179
	v182 = v139 + v179
	if v182 != v115&int32(-2) {
		v138 = v180
		v139 = v182
		goto L53
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	v188 = v180
	goto L52
L59:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v211+v202))) = v209
	goto L48
L60:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v239 = F_MemoizeHash_hash(m, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v238)+16))
	if base.Ui32(v242) <= base.Ui32(v241) {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	goto L11
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(5)
	v1106 = v2
	goto L8
L64:
	;
	goto L65
L65:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v254 = F_ExecStoreMinimalTuple(m, v251, v252, int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	v1106 = v252
	goto L8
L67:
	;
	F_ExecReScan(m, v257)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v257)+12))
	v262 = m.T0[v261].(func(*base.Module, int32) int32)(m, v257)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+13)))
	if v273 == int32(1) {
		goto L12
	} else {
		goto L77
	}
L72:
	;
	if v262 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+4)))
	if v264&int32(2) == int32(0) {
		goto L71
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v269 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v256)+13)) = uint8(v269)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(5)
	v1106 = v2
	goto L8
L76:
	;
	goto L75
L77:
	;
	v276 = F_cache_store_tuple(m, l0, v262)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	if v276 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(4)
	v282 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+224)) = v282 + int64(1)
	goto L81
L80:
	;
	goto L81
L81:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+8))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+32))
	m.T0[v288].(func(*base.Module, int32, int32))(m, v286, v262)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v1106 = v286
	goto L8
L83:
	;
	F_ExecReScan(m, v291)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L4
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v291)+12))
	v296 = m.T0[v295].(func(*base.Module, int32) int32)(m, v291)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L4
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+8))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)+32))
	m.T0[v307].(func(*base.Module, int32, int32))(m, v305, v296)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L4
	} else {
		goto L93
	}
L88:
	;
	if v296 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+4)))
	if v298&int32(2) == int32(0) {
		goto L87
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(5)
	v1106 = v2
	goto L8
L92:
	;
	goto L91
L93:
	;
	v1106 = v305
	goto L8
L94:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v314
	F_errmsg_internal(m, int32(_a_F_ExecMemoize_2), v18)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_ExecMemoize_3), int32(946), int32(_a_F_ExecMemoize_4))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	F_errmsg_internal(m, int32(_a_F_ExecMemoize_5), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_ExecMemoize_3), int32(893), int32(_a_F_ExecMemoize_4))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	if v342 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L102:
	;
	v1101 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v238)+16)) = v1101
	v342 = v1101
	goto L100
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(5)
	v1106 = int32(0)
	goto L8
L104:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1036)+52))
	if v1037 != 0 {
		goto L249
	} else {
		goto L250
	}
L105:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v843 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v238)+8)) = v842 + v843
	*(*uint8)(unsafe.Add(mBase, uint32(v829)+12)) = uint8(v843)
	*(*int32)(unsafe.Add(mBase, uint32(v829)+8)) = v239
	*(*int32)(unsafe.Add(mBase, uint32(v829))) = int32(0)
	v851 = int32(_a_F_ExecMemoize_0)
	v852 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMemoize[1]))
	v854 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMemoize[1])) = v854
	v857 = F_palloc(m, int32(12))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L4
	} else {
		goto L212
	}
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L4
	} else {
		goto L209
	}
L107:
	;
	v357 = *(*int64)(unsafe.Add(mBase, uint32(v238)))
	if v357 == int64(4294967296) {
		goto L106
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v238)+20))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v548 = v547 & v239
	v551 = v546 + v548<<(uint(int32(4))%32)
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+12)))
	if v552 != 0 {
		goto L152
	} else {
		goto L153
	}
L110:
	;
	v360 = int32(0)
	v362 = int64(2)
	v364 = v357 << (uint(int64(1)) % 64)
	if base.Ui64(v364) <= base.Ui64(v362) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v342 = int32(1)
	goto L100
L112:
	;
	v367 = v362
	goto L114
L113:
	;
	v367 = v364
	goto L114
L114:
	;
	v368 = int64(1)
	if v367&(v367-v368) == int64(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v378 = v367
	goto L117
L116:
	;
	v378 = v368 << (uint(int64(64)-base.I64_clz(v367)) % 64)
	goto L117
L117:
	;
	if base.Ui64(v378<<(uint(int64(4))%64)) < base.Ui64(int64(2147483647)) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v238)+20))
	v384 = *(*int64)(unsafe.Add(mBase, uint32(v238)))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v238)+24))
	v390 = F_MemoryContextAllocExtended(m, v385, base.I32_wrap_i64(v378)<<(uint(int32(4))%32), int32(5))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L4
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	goto L7
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v390
	v393 = int64(1)
	if v378&(v378-v393) == int64(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v403 = v378
	goto L124
L123:
	;
	v403 = v393 << (uint(int64(64)-base.I64_clz(v378)) % 64)
	goto L124
L124:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v403<<(uint(int64(4))%64)) {
		goto L7
	} else {
		goto L125
	}
L125:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v238))) = v403
	v411 = base.I32_wrap_i64(v403) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = v411
	if v403 == int64(4294967296) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v420 = int32(-85899346)
	goto L128
L127:
	;
	v420 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v403), float64(0.9)))
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+16)) = v420
	if v384 != int64(0) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v427 = v360
	goto L133
L130:
	;
	goto L131
L131:
	;
	F_pfree(m, v383)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L4
	} else {
		goto L150
	}
L132:
	;
	v456 = v360
	v457 = v453
	goto L138
L133:
	;
	v441 = v383 + v427<<(uint(int32(4))%32)
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441)+12)))
	if v442 != int32(1) {
		v453 = v427
		goto L132
	} else {
		goto L135
	}
L134:
	;
	v453 = int32(0)
	goto L132
L135:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v441)+8))
	if v445&v411 == v427 {
		v453 = v427
		goto L132
	} else {
		goto L136
	}
L136:
	;
	v449 = v427 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v449)) < base.Ui64(v384) {
		v427 = v449
		goto L133
	} else {
		goto L137
	}
L137:
	;
	goto L134
L138:
	;
	v471 = v383 + v457<<(uint(int32(4))%32)
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471)+12)))
	if v472 == int32(1) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	goto L131
L140:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v471)+8))
	v484 = v476
	goto L143
L141:
	;
	goto L142
L142:
	;
	v519 = v457 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v519)) < base.Ui64(v384) {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	v492 = v484 & v475
	v497 = v390 + v492<<(uint(int32(4))%32)
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+12)))
	if v498 != 0 {
		v484 = v492 + int32(1)
		goto L143
	} else {
		goto L145
	}
L144:
	;
	v499 = *(*int64)(unsafe.Add(mBase, uint32(v471)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v497)+8)) = v499
	v501 = *(*int64)(unsafe.Add(mBase, uint32(v471)))
	*(*int64)(unsafe.Add(mBase, uint32(v497))) = v501
	goto L142
L145:
	;
	goto L144
L146:
	;
	v523 = v519
	goto L148
L147:
	;
	v523 = int32(0)
	goto L148
L148:
	;
	v525 = v456 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v525)) < base.Ui64(v384) {
		v456 = v525
		v457 = v523
		goto L138
	} else {
		goto L149
	}
L149:
	;
	goto L139
L150:
	;
	goto L111
L151:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
	v716 = v714 + int32(4)
	v717 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v716 != v717 {
		goto L189
	} else {
		goto L190
	}
L152:
	;
	v556 = v551
	v557 = v548
	v558 = int32(0)
	v561 = v547
	goto L155
L153:
	;
	v701 = v551
	goto L154
L154:
	;
	v829 = v701
	goto L105
L155:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v556)+8))
	if v569 == v239 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v701 = v697
	goto L154
L157:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
	v572 = F_MemoizeHash_equal(m, v238, v571)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L4
	} else {
		goto L160
	}
L158:
	;
	v576 = v569
	v577 = v561
	goto L159
L159:
	;
	v578 = v576 & v577
	if base.Ui32(v557) < base.Ui32(v578) {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	if v572 != 0 {
		goto L151
	} else {
		goto L161
	}
L161:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v556)+8))
	v576 = v575
	v577 = v574
	goto L159
L162:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v582 = v557 + v580
	goto L164
L163:
	;
	v582 = v557
	goto L164
L164:
	;
	v585 = v577 & (v557 + int32(1))
	if base.Ui32(v582-v578) < base.Ui32(v558) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v590 = v546 + v585<<(uint(int32(4))%32)
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590)+12)))
	if v591 != 0 {
		goto L168
	} else {
		goto L169
	}
L166:
	;
	goto L167
L167:
	;
	v685 = v558 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v685) {
		goto L184
	} else {
		goto L185
	}
L168:
	;
	v594 = v585
	v598 = int32(0)
	goto L171
L169:
	;
	v627 = v585
	v630 = v590
	goto L170
L170:
	;
	if v627 != v557 {
		goto L178
	} else {
		goto L179
	}
L171:
	;
	v609 = v598 + int32(1)
	if int32(151) <= v609 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v627 = v621
	v630 = v624
	goto L170
L173:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v614 = *(*int64)(unsafe.Add(mBase, uint32(v238)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v612), base.F64_convert_i64_u(v614)), float64(0.1)) != 0 {
		goto L102
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v621 = (v594 + int32(1)) & v577
	v624 = v546 + v621<<(uint(int32(4))%32)
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624)+12)))
	if v625 != 0 {
		v594 = v621
		v598 = v609
		goto L171
	} else {
		goto L177
	}
L176:
	;
	goto L175
L177:
	;
	goto L172
L178:
	;
	v643 = v627
	v646 = v630
	goto L181
L179:
	;
	goto L180
L180:
	;
	v829 = v556
	goto L105
L181:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v660 = v657 & (v643 - int32(1))
	v663 = v546 + v660<<(uint(int32(4))%32)
	v664 = *(*int64)(unsafe.Add(mBase, uint32(v663)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v646)+8)) = v664
	v666 = *(*int64)(unsafe.Add(mBase, uint32(v663)))
	*(*int64)(unsafe.Add(mBase, uint32(v646))) = v666
	if v660 != v557 {
		v643 = v660
		v646 = v663
		goto L181
	} else {
		goto L183
	}
L182:
	;
	goto L180
L183:
	;
	goto L182
L184:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v690 = *(*int64)(unsafe.Add(mBase, uint32(v238)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v688), base.F64_convert_i64_u(v690)), float64(0.1)) != 0 {
		goto L102
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v697 = v546 + v585<<(uint(int32(4))%32)
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v697)+12)))
	if v698 != 0 {
		v556 = v697
		v557 = v585
		v558 = v685
		v561 = v577
		goto L155
	} else {
		goto L188
	}
L187:
	;
	goto L186
L188:
	;
	goto L156
L189:
	;
	v720 = l0 + int32(180)
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v714)+4))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v714)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v721)+4)) = v722
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v714)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v722))) = v724
	v726 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if v726 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L190:
	;
	goto L191
L191:
	;
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+13)))
	if v739 == int32(1) {
		goto L196
	} else {
		goto L197
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v720
	goto L194
L193:
	;
	goto L194
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v714)+8)) = v720
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v720)))
	*(*int32)(unsafe.Add(mBase, uint32(v714)+4)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v732)+4)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v720))) = v716
	goto L191
L195:
	;
	goto L103
L196:
	;
	v742 = *(*int64)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+200)) = v742 + int64(1)
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v556)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v556
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v746
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v556)+4))
	if v749 == int32(0) {
		goto L195
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	v760 = *(*int64)(unsafe.Add(mBase, uint32(l0)+208))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v760 + int64(1)
	v764 = int64(0)
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v556)+4))
	if v765 != 0 {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(2)
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v556)+4))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v754)))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v758 = F_ExecStoreMinimalTuple(m, v755, v756, int32(0))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	v1106 = v756
	goto L8
L201:
	;
	v767 = v765
	v778 = v764
	goto L204
L202:
	;
	v804 = v764
	goto L203
L203:
	;
	v807 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v556)+4)) = v807
	*(*uint8)(unsafe.Add(mBase, uint32(v556)+13)) = uint8(v807)
	v811 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v811 - v804
	v1023 = v556
	goto L104
L204:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v767)+4))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v767)))
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v782)))
	F_pfree(m, v782)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L4
	} else {
		goto L206
	}
L205:
	;
	v804 = v791
	goto L203
L206:
	;
	F_pfree(m, v767)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L4
	} else {
		goto L207
	}
L207:
	;
	v791 = v778 + base.I64_extend_i32_u(v783+int32(8))
	if v781 != 0 {
		v767 = v781
		v778 = v791
		goto L204
	} else {
		goto L208
	}
L208:
	;
	goto L205
L209:
	;
	F_errmsg_internal(m, int32(_a_F_ExecMemoize_6), int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L4
	} else {
		goto L210
	}
L210:
	;
	F_errfinish(m, int32(_a_F_ExecMemoize_7), int32(630), int32(_a_F_ExecMemoize_8))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L4
	} else {
		goto L211
	}
L211:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v829))) = v857
	v860 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v860)+8))
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v862)+48))
	v864 = m.T0[v863].(func(*base.Module, int32, int32) int32)(m, v860, int32(0))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L4
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v857))) = v864
	v867 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v829)))
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v869)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v867 + base.I64_extend_i32_u(v870+int32(28))
	v876 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v829)+4)) = v876
	*(*uint8)(unsafe.Add(mBase, uint32(v829)+13)) = uint8(v876)
	v881 = l0 + int32(180)
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v829)))
	v884 = v882 + int32(4)
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if v885 == v876 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v881
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v881
	goto L216
L215:
	;
	goto L216
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v882)+8)) = v881
	v891 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	*(*int32)(unsafe.Add(mBase, uint32(v882)+4)) = v891
	*(*int32)(unsafe.Add(mBase, uint32(v891)+4)) = v884
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v884
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMemoize[1])) = v852
	v899 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v900 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	if base.Ui64(v899) <= base.Ui64(v900) {
		v1004 = v829
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v1017 = *(*int64)(unsafe.Add(mBase, uint32(l0)+208))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v1017 + int64(1)
	v1023 = v1004
	goto L104
L218:
	;
	v902 = F_cache_reduce_memory(m, l0, v857)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L4
	} else {
		goto L220
	}
L219:
	;
	v1004 = int32(0)
	goto L217
L220:
	;
	if v902 == int32(0) {
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829)+12)))
	if v906 == int32(1) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v829)))
	if v909 == v857 {
		v1004 = v829
		goto L217
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v913)+8))
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v914)+12))
	m.T0[v915].(func(*base.Module, int32))(m, v913)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L4
	} else {
		goto L226
	}
L225:
	;
	goto L224
L226:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v857)))
	v920 = F_ExecStoreMinimalTuple(m, v918, v912, int32(0))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L4
	} else {
		goto L227
	}
L227:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v912)+12))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v922)))
	v924 = int32(*(*int16)(unsafe.Add(mBase, uint32(v912)+6)))
	if v924 < v923 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	F_slot_getsomeattrs_int(m, v912, v923)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L4
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v929 = v911 << (uint(int32(2)) % 32)
	if v929 != 0 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	goto L230
L232:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v913)+16))
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v912)+16))
	base.MemoryCopy(m, v930, v931, v929)
	goto L234
L233:
	;
	goto L234
L234:
	;
	if v911 != 0 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v913)+20))
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v912)+20))
	base.MemoryCopy(m, v933, v934, v911)
	goto L237
L236:
	;
	goto L237
L237:
	;
	v936 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v913)+4)))
	v938 = v936 & int32(_a_F_ExecMemoize_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v913)+4)) = uint16(v938)
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v913)+12))
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v940)))
	*(*uint16)(unsafe.Add(mBase, uint32(v913)+6)) = uint16(v941)
	goto L238
L238:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v944 = F_MemoizeHash_hash(m, v943)
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L4
	} else {
		goto L239
	}
L239:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v943)+20))
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v943)+12))
	v948 = v944 & v947
	v951 = v946 + v948<<(uint(int32(4))%32)
	v952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v951)+12)))
	if v952 == int32(0) {
		goto L219
	} else {
		goto L240
	}
L240:
	;
	v956 = v948
	v957 = v951
	v958 = v946
	v962 = v947
	goto L241
L241:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v957)+8))
	if v970 == v944 {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	goto L219
L243:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v957)))
	v973 = F_MemoizeHash_equal(m, v943, v972)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L4
	} else {
		goto L246
	}
L244:
	;
	v977 = v958
	v978 = v962
	goto L245
L245:
	;
	v981 = v978 & (v956 + int32(1))
	v984 = v977 + v981<<(uint(int32(4))%32)
	v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v984)+12)))
	if v985 != 0 {
		v956 = v981
		v957 = v984
		v958 = v977
		v962 = v978
		goto L241
	} else {
		goto L248
	}
L246:
	;
	if v973 != 0 {
		v1004 = v957
		goto L217
	} else {
		goto L247
	}
L247:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v943)+12))
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v943)+20))
	v977 = v976
	v978 = v975
	goto L245
L248:
	;
	goto L242
L249:
	;
	F_ExecReScan(m, v1036)
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L4
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1036)+12))
	v1041 = m.T0[v1040].(func(*base.Module, int32) int32)(m, v1036)
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L4
	} else {
		goto L254
	}
L252:
	;
	goto L251
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v1023
	if v1023 != 0 {
		goto L264
	} else {
		goto L265
	}
L254:
	;
	if v1041 != 0 {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1041)+4)))
	if v1043&int32(2) == int32(0) {
		goto L253
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	if v1023 != 0 {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	goto L257
L259:
	;
	v1048 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1023)+13)) = uint8(v1048)
	goto L261
L260:
	;
	goto L261
L261:
	;
	goto L103
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v1061
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+8))
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+32))
	m.T0[v1065].(func(*base.Module, int32, int32))(m, v1063, v1041)
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L4
	} else {
		goto L269
	}
L263:
	;
	v1058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1023)+13)) = uint8(v1058)
	v1061 = int32(3)
	goto L262
L264:
	;
	v1051 = F_cache_store_tuple(m, l0, v1041)
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L4
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v1053 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+224)) = v1053 + int64(1)
	v1061 = int32(4)
	goto L262
L267:
	;
	if v1051 != 0 {
		goto L263
	} else {
		goto L268
	}
L268:
	;
	goto L266
L269:
	;
	v1106 = v1063
	goto L8
L270:
	;
	F_errmsg_internal(m, int32(_a_F_ExecMemoize_9), int32(0))
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L4
	} else {
		goto L271
	}
L271:
	;
	F_errfinish(m, int32(_a_F_ExecMemoize_7), int32(327), int32(_a_F_ExecMemoize_10))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L4
	} else {
		goto L272
	}
L272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	var v73 int32
	_ = v73
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
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 float64
	_ = v237
	var v241 int32
	_ = v241
	var v244 float64
	_ = v244
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
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
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[0]))
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
	return v281
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
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+72))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v253)+16))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+8))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
	m.T0[v257].(func(*base.Module, int32))(m, v255)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
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
		v281 = v54
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
	if v60&int32(2) != 0 {
		v281 = v54
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
	v73 = v54
	goto L23
L23:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v73<<(uint(int32(2))%32))))
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
	v119 = v73 + int32(1)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v119 < v120 {
		v73 = v119
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
	v174 = int32(_a_F_ExecNestLoop_0)
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1]))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1])) = v177
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
	*(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1])) = v175
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
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	F_MemoryContextReset(m, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L67
	}
L50:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v241 == int32(0) {
		goto L49
	} else {
		goto L66
	}
L51:
	;
	v193 = int32(_a_F_ExecNestLoop_0)
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1]))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1])) = v196
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
	v208 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)) = uint8(v208)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v210 == int32(5) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1])) = v194
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
	v213 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)) = uint8(v213)
	goto L8
L57:
	;
	goto L58
L58:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	if v215 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v218 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)) = uint8(v218)
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
	v222 = int32(_a_F_ExecNestLoop_0)
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1]))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1])) = v225
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v230 = m.T0[v229].(func(*base.Module, int32, int32, int32) int32)(m, v27, v30, v17+int32(15))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1])) = v223
	if v230 != 0 {
		goto L37
	} else {
		goto L64
	}
L64:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v234 == int32(0) {
		goto L49
	} else {
		goto L65
	}
L65:
	;
	v237 = *(*float64)(unsafe.Add(mBase, uint32(v234)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v234)+248)) = base.F64_add(v237, float64(1))
	goto L49
L66:
	;
	v244 = *(*float64)(unsafe.Add(mBase, uint32(v241)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v241)+240)) = base.F64_add(v244, float64(1))
	goto L49
L67:
	;
	goto L8
L68:
	;
	v260 = int32(_a_F_ExecNestLoop_0)
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1]))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v254)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1])) = v263
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v253)+24))
	v269 = m.T0[v268].(func(*base.Module, int32, int32, int32) int32)(m, v253+int32(4), v254, int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1])) = v261
	v273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v255)+4)))
	v275 = v273 & int32(_a_F_ExecNestLoop_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v255)+4)) = uint16(v275)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	*(*uint16)(unsafe.Add(mBase, uint32(v255)+6)) = uint16(v278)
	v281 = v255
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
	var v144 int64
	_ = v144
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v14 - int32(394) {
	case 0:
		v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if v216 != 0 {
			F_ExecRestrPos(m, v216)
			mBase = m.M
			v218 = m.ExcPending
			if v218 != 0 {
				return
			} else {
				m.G0 = v12 + int32(16)
				return
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v222 = m.ExcPending
			if v222 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_ExecRestrPos_0), int32(0))
				mBase = m.M
				v226 = m.ExcPending
				if v226 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_ExecRestrPos_1), int32(168), int32(_a_F_ExecRestrPos_2))
					mBase = m.M
					v231 = m.ExcPending
					if v231 != 0 {
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
		v235 = m.ExcPending
		if v235 != 0 {
			return
		} else {
			v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v236
			F_errmsg_internal(m, int32(_a_F_ExecRestrPos_3), v12)
			mBase = m.M
			v240 = m.ExcPending
			if v240 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_ExecRestrPos_4), int32(405), int32(_a_F_ExecRestrPos_5))
				mBase = m.M
				v245 = m.ExcPending
				if v245 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 11:
		v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+156))
		if v247 == int32(0) {
			v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
			F_index_restrpos(m, v284)
			mBase = m.M
			v286 = m.ExcPending
			if v286 != 0 {
				return
			} else {
				m.G0 = v12 + int32(16)
				return
			}
		} else {
			v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+72))
			v253 = v251 - int32(1)
			v255 = v253 << (uint(int32(2)) % 32)
			v256 = *(*int32)(unsafe.Add(mBase, uint32(v247)+16))
			v258 = *(*int32)(unsafe.Add(mBase, uint32(v255+v256)))
			if v258 == int32(0) {
				v261 = *(*int32)(unsafe.Add(mBase, uint32(v247)+36))
				v263 = *(*int32)(unsafe.Add(mBase, uint32(v261+v255)))
				if v263 == int32(0) {
					v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
					F_index_restrpos(m, v284)
					mBase = m.M
					v286 = m.ExcPending
					if v286 != 0 {
						return
					} else {
						m.G0 = v12 + int32(16)
						return
					}
				} else {
					v266 = *(*int32)(unsafe.Add(mBase, uint32(v247)+40))
					v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266+v253))))
					if v268 != 0 {
						m.G0 = v12 + int32(16)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v272 = m.ExcPending
						if v272 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_ExecRestrPos_6), int32(0))
							mBase = m.M
							v276 = m.ExcPending
							if v276 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ExecRestrPos_7), int32(889), int32(_a_F_ExecRestrPos_8))
								mBase = m.M
								v281 = m.ExcPending
								if v281 != 0 {
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
				v266 = *(*int32)(unsafe.Add(mBase, uint32(v247)+40))
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266+v253))))
				if v268 != 0 {
					m.G0 = v12 + int32(16)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v272 = m.ExcPending
					if v272 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_ExecRestrPos_6), int32(0))
						mBase = m.M
						v276 = m.ExcPending
						if v276 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecRestrPos_7), int32(889), int32(_a_F_ExecRestrPos_8))
							mBase = m.M
							v281 = m.ExcPending
							if v281 != 0 {
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
							F_errmsg_internal(m, int32(_a_F_ExecRestrPos_9), int32(0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ExecRestrPos_10), int32(508), int32(_a_F_ExecRestrPos_11))
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
						F_errmsg_internal(m, int32(_a_F_ExecRestrPos_9), int32(0))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecRestrPos_10), int32(508), int32(_a_F_ExecRestrPos_11))
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
					F_errmsg(m, int32(_a_F_ExecRestrPos_12), v62)
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ExecRestrPos_13), int32(156), int32(_a_F_ExecRestrPos_14))
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
			v99 = int32(_a_F_ExecRestrPos_15)
			v100 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRestrPos[0]))
			v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+28))
			*(*int32)(unsafe.Add(mBase, _c_F_ExecRestrPos[0])) = v103
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v102)+64))
			switch v105 - int32(3) {
			case 0:
				v196 = *(*int32)(unsafe.Add(mBase, uint32(v102)+224))
				*(*int32)(unsafe.Add(mBase, uint32(v102)+204)) = v196
				v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+228)))
				*(*uint8)(unsafe.Add(mBase, uint32(v102)+208)) = uint8(v204)
				*(*int32)(unsafe.Add(mBase, _c_F_ExecRestrPos[0])) = v100
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
								v149 = v130
								if v149 < v110 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v173 = m.ExcPending
									if v173 != 0 {
										return
									} else {
										F_errmsg_internal(m, int32(_a_F_ExecRestrPos_16), int32(0))
										mBase = m.M
										v177 = m.ExcPending
										if v177 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ExecRestrPos_17), int32(1151), int32(_a_F_ExecRestrPos_18))
											mBase = m.M
											v182 = m.ExcPending
											if v182 != 0 {
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
									v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+228)))
									*(*uint8)(unsafe.Add(mBase, uint32(v102)+208)) = uint8(v204)
									*(*int32)(unsafe.Add(mBase, _c_F_ExecRestrPos[0])) = v100
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
										v158 = m.ExcPending
										if v158 != 0 {
											return
										} else {
											F_errcode_for_file_access(m)
											mBase = m.M
											v160 = m.ExcPending
											if v160 != 0 {
												return
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v113))) = v109
												F_errmsg(m, int32(_a_F_ExecRestrPos_19), v113)
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_ExecRestrPos_17), int32(288), int32(_a_F_ExecRestrPos_20))
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
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
										F_BufFileReadExact(m, v136, v131, int32(_a_F_ExecRestrPos_21))
										mBase = m.M
										v139 = m.ExcPending
										if v139 != 0 {
											return
										} else {
											v140 = int32(_a_F_ExecRestrPos_22)
											*(*int32)(unsafe.Add(mBase, uint32(v108)+56)) = v140
											*(*int64)(unsafe.Add(mBase, uint32(v108)+16)) = v109
											v143 = *(*int32)(unsafe.Add(mBase, uint32(v108)+40))
											v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)+uint32(_c_F_ExecRestrPos[1])))
											*(*int64)(unsafe.Add(mBase, uint32(v108)+24)) = v144
											v149 = v140
											if v149 < v110 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v173 = m.ExcPending
												if v173 != 0 {
													return
												} else {
													F_errmsg_internal(m, int32(_a_F_ExecRestrPos_16), int32(0))
													mBase = m.M
													v177 = m.ExcPending
													if v177 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ExecRestrPos_17), int32(1151), int32(_a_F_ExecRestrPos_18))
														mBase = m.M
														v182 = m.ExcPending
														if v182 != 0 {
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
												v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+228)))
												*(*uint8)(unsafe.Add(mBase, uint32(v102)+208)) = uint8(v204)
												*(*int32)(unsafe.Add(mBase, _c_F_ExecRestrPos[0])) = v100
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
						v149 = v130
						if v149 < v110 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v173 = m.ExcPending
							if v173 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F_ExecRestrPos_16), int32(0))
								mBase = m.M
								v177 = m.ExcPending
								if v177 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ExecRestrPos_17), int32(1151), int32(_a_F_ExecRestrPos_18))
									mBase = m.M
									v182 = m.ExcPending
									if v182 != 0 {
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
							v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+228)))
							*(*uint8)(unsafe.Add(mBase, uint32(v102)+208)) = uint8(v204)
							*(*int32)(unsafe.Add(mBase, _c_F_ExecRestrPos[0])) = v100
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
								v158 = m.ExcPending
								if v158 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v113))) = v109
										F_errmsg(m, int32(_a_F_ExecRestrPos_19), v113)
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ExecRestrPos_17), int32(288), int32(_a_F_ExecRestrPos_20))
											mBase = m.M
											v169 = m.ExcPending
											if v169 != 0 {
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
								F_BufFileReadExact(m, v136, v131, int32(_a_F_ExecRestrPos_21))
								mBase = m.M
								v139 = m.ExcPending
								if v139 != 0 {
									return
								} else {
									v140 = int32(_a_F_ExecRestrPos_22)
									*(*int32)(unsafe.Add(mBase, uint32(v108)+56)) = v140
									*(*int64)(unsafe.Add(mBase, uint32(v108)+16)) = v109
									v143 = *(*int32)(unsafe.Add(mBase, uint32(v108)+40))
									v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)+uint32(_c_F_ExecRestrPos[1])))
									*(*int64)(unsafe.Add(mBase, uint32(v108)+24)) = v144
									v149 = v140
									if v149 < v110 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v173 = m.ExcPending
										if v173 != 0 {
											return
										} else {
											F_errmsg_internal(m, int32(_a_F_ExecRestrPos_16), int32(0))
											mBase = m.M
											v177 = m.ExcPending
											if v177 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_ExecRestrPos_17), int32(1151), int32(_a_F_ExecRestrPos_18))
												mBase = m.M
												v182 = m.ExcPending
												if v182 != 0 {
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
										v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+228)))
										*(*uint8)(unsafe.Add(mBase, uint32(v102)+208)) = uint8(v204)
										*(*int32)(unsafe.Add(mBase, _c_F_ExecRestrPos[0])) = v100
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
				v186 = m.ExcPending
				if v186 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_ExecRestrPos_23), int32(0))
					mBase = m.M
					v190 = m.ExcPending
					if v190 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ExecRestrPos_24), int32(2485), int32(_a_F_ExecRestrPos_25))
						mBase = m.M
						v195 = m.ExcPending
						if v195 != 0 {
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
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
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
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
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
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
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
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v480 int32
	_ = v480
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int64
	_ = v528
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int64
	_ = v573
	var v574 int32
	_ = v574
	var v578 int64
	_ = v578
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v604 int32
	_ = v604
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
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
	v670 = m.ExcPending
	if v670 != 0 {
		goto L9
	} else {
		goto L184
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L9
	} else {
		goto L181
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L9
	} else {
		goto L177
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L9
	} else {
		goto L174
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L9
	} else {
		goto L170
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
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteGrantStmt[0]))
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
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+96)) = uint8(v420)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+100)) = v422
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v424 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L13:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v252 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L14:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v27 - int32(12) {
	case 0, 37:
		goto L15
	default:
		goto L17
	case 15:
		goto L18
	case 25, 29:
		goto L16
	}
L15:
	;
	if v30 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L16:
	;
	if v30 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L17:
	;
	if v30 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L18:
	;
	if v30 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v406 = int32(0)
	goto L12
L20:
	;
	goto L21
L21:
	;
	v36 = int32(0)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v37 <= v36 {
		v406 = v36
		goto L12
	} else {
		goto L22
	}
L22:
	;
	v42 = v36
	v46 = v2
	goto L23
L23:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v46<<(uint(int32(2))%32))))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v60 = F_ParameterAclLookup(m, v58, int32(1))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L9
	} else {
		goto L25
	}
L24:
	;
	v406 = v128
	goto L12
L25:
	;
	v64 = int32(0)
	if v60|base.B2i32(v25&int32(1) == v64) == v64 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v69 = m.G0
	v71 = v69 - int32(16)
	m.G0 = v71
	v73 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = v73
	*(*uint16)(unsafe.Add(mBase, uint32(v71))) = uint16(v73)
	v80 = F_find_option(m, v58, v73, int32(1), int32(10))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L29
	}
L27:
	;
	v121 = v60
	goto L28
L28:
	;
	if v121 != 0 {
		goto L43
	} else {
		goto L44
	}
L29:
	;
	if v80 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v86 = F_assignable_custom_variable_name(m, v58, int32(0), int32(21))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L9
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v88 = F_convert_GUC_name_for_parameter_acl(m, v58)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L9
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	v92 = F_table_open(m, int32(_a_F_ExecuteGrantStmt_0), int32(3))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v92)+52))
	v97 = F_GetNewOidWithIndex(m, v92, int32(_a_F_ExecuteGrantStmt_1), int32(1))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = v97
	v100 = F_cstring_to_text(m, v88)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	v102 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v71)+2)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v100
	v107 = F_heap_form_tuple(m, v94, v71+int32(4), v71)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	F_CatalogTupleInsert(m, v92, v107)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	F_pfree(m, v107)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	F_relation_close(m, v92, int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	m.G0 = v71 + int32(16)
	F_CommandCounterIncrement(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L9
	} else {
		goto L42
	}
L42:
	;
	v121 = v97
	goto L28
L43:
	;
	v126 = F_lappend_oid(m, v42, v121)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L9
	} else {
		goto L46
	}
L44:
	;
	v128 = v42
	goto L45
L45:
	;
	v130 = v46 + int32(1)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v130 < v131 {
		v42 = v128
		v46 = v130
		goto L23
	} else {
		goto L47
	}
L46:
	;
	v128 = v126
	goto L45
L47:
	;
	goto L24
L48:
	;
	v406 = int32(0)
	goto L12
L49:
	;
	goto L50
L50:
	;
	v136 = int32(0)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v137 <= v136 {
		v406 = v136
		goto L12
	} else {
		goto L51
	}
L51:
	;
	v141 = v2
	v142 = v136
	goto L52
L52:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155+v141<<(uint(int32(2))%32))))
	v160 = int32(0)
	F_get_object_address(m, v16+int32(112), v27, v159, v160, int32(1), v160)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L9
	} else {
		goto L54
	}
L53:
	;
	v406 = v166
	goto L12
L54:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v16)+116))
	v166 = F_lappend_oid(m, v142, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	v169 = v141 + int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v169 < v170 {
		v141 = v169
		v142 = v166
		goto L52
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v406 = int32(0)
	goto L12
L58:
	;
	goto L59
L59:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v175 <= int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v406 = int32(0)
	goto L12
L61:
	;
	goto L62
L62:
	;
	v181 = v2
	v182 = int32(0)
	goto L63
L63:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v193+v181<<(uint(int32(2))%32))))
	v199 = int32(0)
	v202 = F_RangeVarGetRelidExtended(m, v197, int32(1), v199, v199, v199)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L9
	} else {
		goto L65
	}
L64:
	;
	v406 = v204
	goto L12
L65:
	;
	v204 = F_lappend_oid(m, v182, v202)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	v207 = v181 + int32(1)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v207 < v208 {
		v181 = v207
		v182 = v204
		goto L63
	} else {
		goto L67
	}
L67:
	;
	goto L64
L68:
	;
	v406 = int32(0)
	goto L12
L69:
	;
	goto L70
L70:
	;
	v213 = int32(0)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v214 <= v213 {
		v406 = v213
		goto L12
	} else {
		goto L71
	}
L71:
	;
	v218 = v2
	v219 = v213
	goto L72
L72:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v232+v218<<(uint(int32(2))%32))))
	v237 = F_makeTypeNameFromNameList(m, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L9
	} else {
		goto L74
	}
L73:
	;
	v406 = v246
	goto L12
L74:
	;
	F_get_object_address(m, v16+int32(112), v27, v237, v16+int32(108), int32(1), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L9
	} else {
		goto L75
	}
L75:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v16)+116))
	v246 = F_lappend_oid(m, v219, v245)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	v249 = v218 + int32(1)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v249 < v250 {
		v218 = v249
		v219 = v246
		goto L72
	} else {
		goto L77
	}
L77:
	;
	goto L73
L78:
	;
	v406 = int32(0)
	goto L12
L79:
	;
	goto L80
L80:
	;
	v256 = int32(0)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	if v257 <= v256 {
		v406 = v256
		goto L12
	} else {
		goto L81
	}
L81:
	;
	v263 = v27 - int32(19)
	v266 = v256
	v270 = v2
	goto L82
L82:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v277+v270<<(uint(int32(2))%32))))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	v284 = F_LookupExplicitNamespace(m, v282, int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L9
	} else {
		goto L84
	}
L83:
	;
	v406 = v389
	goto L12
L84:
	;
	switch v263 {
	case 0, 10, 15:
		goto L88
	default:
		goto L87
	case 18:
		goto L89
	case 22:
		goto L86
	}
L85:
	;
	v401 = v270 + int32(1)
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	if v401 < v402 {
		v266 = v389
		v270 = v401
		goto L82
	} else {
		goto L121
	}
L86:
	;
	v363 = F_getRelationsInNamespace(m, v284, int32(114))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L9
	} else {
		goto L111
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L9
	} else {
		goto L108
	}
L88:
	;
	v293 = int32(3)
	F_ScanKeyInit(m, v16+int32(112), v293, v293, int32(184), v284)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L9
	} else {
		goto L92
	}
L89:
	;
	v287 = F_getRelationsInNamespace(m, v284, int32(83))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L9
	} else {
		goto L90
	}
L90:
	;
	v289 = F_list_concat(m, v266, v287)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L9
	} else {
		goto L91
	}
L91:
	;
	v389 = v289
	goto L85
L92:
	;
	switch v263 {
	case 0:
		v301 = int32(70)
		goto L94
	default:
		v309 = int32(1)
		goto L93
	case 10:
		goto L95
	}
L93:
	;
	v312 = F_table_open(m, int32(1255), int32(1))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L9
	} else {
		goto L97
	}
L94:
	;
	F_ScanKeyInit(m, v16+int32(160), int32(10), int32(3), v301, int32(112))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L9
	} else {
		goto L96
	}
L95:
	;
	v301 = int32(61)
	goto L94
L96:
	;
	v309 = int32(2)
	goto L93
L97:
	;
	v316 = F_table_beginscan_catalog(m, v312, v309, v16+int32(112))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L9
	} else {
		goto L98
	}
L98:
	;
	v320 = v266
	goto L99
L99:
	;
	v331 = F_heap_getnext(m, v316)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L9
	} else {
		goto L101
	}
L100:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)+188))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)+12))
	m.T0[v341].(func(*base.Module, int32))(m, v316)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L9
	} else {
		goto L106
	}
L101:
	;
	if v331 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v331)+16))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+22)))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v333+v334)))
	v337 = F_lappend_oid(m, v320, v336)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L9
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	goto L100
L105:
	;
	v320 = v337
	goto L99
L106:
	;
	F_relation_close(m, v312, int32(1))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L9
	} else {
		goto L107
	}
L107:
	;
	v389 = v320
	goto L85
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v27
	F_errmsg_internal(m, int32(_a_F_ExecuteGrantStmt_2), v16+int32(48))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L9
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(_a_F_ExecuteGrantStmt_3), int32(866), int32(_a_F_ExecuteGrantStmt_4))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L9
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	v365 = F_list_concat(m, v266, v363)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L9
	} else {
		goto L112
	}
L112:
	;
	v368 = F_getRelationsInNamespace(m, v284, int32(118))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L9
	} else {
		goto L113
	}
L113:
	;
	v370 = F_list_concat(m, v365, v368)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L9
	} else {
		goto L114
	}
L114:
	;
	v373 = F_getRelationsInNamespace(m, v284, int32(109))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L9
	} else {
		goto L115
	}
L115:
	;
	v375 = F_list_concat(m, v370, v373)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L9
	} else {
		goto L116
	}
L116:
	;
	v378 = F_getRelationsInNamespace(m, v284, int32(102))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L9
	} else {
		goto L117
	}
L117:
	;
	v380 = F_list_concat(m, v375, v378)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L9
	} else {
		goto L118
	}
L118:
	;
	v383 = F_getRelationsInNamespace(m, v284, int32(112))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L9
	} else {
		goto L119
	}
L119:
	;
	v385 = F_list_concat(m, v380, v383)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L9
	} else {
		goto L120
	}
L120:
	;
	v389 = v385
	goto L85
L121:
	;
	goto L83
L122:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	switch v480 - int32(9) {
	case 0:
		goto L148
	default:
		goto L135
	case 3:
		goto L147
	case 7:
		goto L138
	case 8:
		goto L137
	case 10:
		goto L146
	case 12:
		goto L145
	case 13:
		goto L144
	case 18:
		goto L136
	case 20:
		goto L142
	case 25:
		goto L141
	case 27:
		goto L143
	case 28:
		goto L134
	case 32:
		v527 = int32(_a_F_ExecuteGrantStmt_5)
		v528 = int64(-16768)
		goto L133
	case 33:
		goto L140
	case 40:
		goto L139
	}
L123:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v424)+4))
	if v427 <= int32(0) {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v430 = int32(0)
	v434 = v430
	v438 = v430
	goto L125
L125:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v424)+12))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v445+v434<<(uint(int32(2))%32))))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v450 != int32(4) {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	goto L122
L127:
	;
	v454 = F_get_rolespec_oid(m, v449, int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L9
	} else {
		goto L130
	}
L128:
	;
	v457 = int32(0)
	goto L129
L129:
	;
	v458 = F_lappend_oid(m, v438, v457)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L9
	} else {
		goto L131
	}
L130:
	;
	v457 = v454
	goto L129
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v458
	v462 = v434 + int32(1)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v424)+4))
	if v462 < v463 {
		v434 = v462
		v438 = v458
		goto L125
	} else {
		goto L132
	}
L132:
	;
	goto L126
L133:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v529 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L134:
	;
	v527 = int32(_a_F_ExecuteGrantStmt_6)
	v528 = int64(-263)
	goto L133
L135:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L9
	} else {
		goto L149
	}
L136:
	;
	v527 = int32(_a_F_ExecuteGrantStmt_7)
	v528 = int64(-12289)
	goto L133
L137:
	;
	v527 = int32(_a_F_ExecuteGrantStmt_8)
	v528 = int64(-257)
	goto L133
L138:
	;
	v527 = int32(_a_F_ExecuteGrantStmt_9)
	v528 = int64(-257)
	goto L133
L139:
	;
	v527 = int32(_a_F_ExecuteGrantStmt_10)
	v528 = int64(-257)
	goto L133
L140:
	;
	v527 = int32(_a_F_ExecuteGrantStmt_11)
	v528 = int64(-513)
	goto L133
L141:
	;
	v527 = int32(_a_F_ExecuteGrantStmt_12)
	v528 = int64(-129)
	goto L133
L142:
	;
	v527 = int32(_a_F_ExecuteGrantStmt_13)
	v528 = int64(-129)
	goto L133
L143:
	;
	v527 = int32(_a_F_ExecuteGrantStmt_14)
	v528 = int64(-769)
	goto L133
L144:
	;
	v527 = int32(_a_F_ExecuteGrantStmt_15)
	v528 = int64(-7)
	goto L133
L145:
	;
	v527 = int32(_a_F_ExecuteGrantStmt_16)
	v528 = int64(-257)
	goto L133
L146:
	;
	v527 = int32(_a_F_ExecuteGrantStmt_17)
	v528 = int64(-129)
	goto L133
L147:
	;
	v527 = int32(_a_F_ExecuteGrantStmt_18)
	v528 = int64(-257)
	goto L133
L148:
	;
	v527 = int32(_a_F_ExecuteGrantStmt_19)
	v528 = int64(-3585)
	goto L133
L149:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v513
	F_errmsg_internal(m, int32(_a_F_ExecuteGrantStmt_2), v16+int32(16))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L9
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_ExecuteGrantStmt_3), int32(540), int32(_a_F_ExecuteGrantStmt_20))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L9
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	F_ExecGrantStmt_oids(m, v16-int32(-64))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L9
	} else {
		goto L169
	}
L153:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = int64(0)
	v534 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+76)) = uint8(v534)
	goto L152
L154:
	;
	goto L155
L155:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = int64(0)
	v538 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+76)) = uint8(v538)
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v529)+4))
	if v541 <= v538 {
		goto L152
	} else {
		goto L156
	}
L156:
	;
	v547 = v538
	v551 = int32(0)
	goto L157
L157:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v529)+12))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v558+v547<<(uint(int32(2))%32))))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v562)+8))
	if v563 != 0 {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	goto L152
L159:
	;
	v585 = v547 + int32(1)
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v529)+4))
	if v585 < v586 {
		v547 = v585
		v551 = v582
		goto L157
	} else {
		goto L168
	}
L160:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v564 != int32(41) {
		goto L3
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v562)+4))
	if v570 == int32(0) {
		goto L2
	} else {
		goto L165
	}
L163:
	;
	v567 = F_lappend(m, v551, v562)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L9
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = v567
	v582 = v567
	goto L159
L165:
	;
	v573 = F_string_to_privilege(m, v570)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L9
	} else {
		goto L166
	}
L166:
	;
	if v573&v528 != int64(0) {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v578 = *(*int64)(unsafe.Add(mBase, uint32(v16)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = v578 | v573
	v582 = v551
	goto L159
L168:
	;
	goto L158
L169:
	;
	m.G0 = v16 + int32(208)
	return
L170:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L9
	} else {
		goto L171
	}
L171:
	;
	F_errmsg(m, int32(_a_F_ExecuteGrantStmt_21), int32(0))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L9
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_ExecuteGrantStmt_3), int32(412), int32(_a_F_ExecuteGrantStmt_20))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L9
	} else {
		goto L173
	}
L173:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L174:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v628
	F_errmsg_internal(m, int32(_a_F_ExecuteGrantStmt_22), v16)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L9
	} else {
		goto L175
	}
L175:
	;
	F_errfinish(m, int32(_a_F_ExecuteGrantStmt_3), int32(434), int32(_a_F_ExecuteGrantStmt_20))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L9
	} else {
		goto L176
	}
L176:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L177:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L9
	} else {
		goto L178
	}
L178:
	;
	F_errmsg(m, int32(_a_F_ExecuteGrantStmt_23), int32(0))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L9
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(_a_F_ExecuteGrantStmt_3), int32(575), int32(_a_F_ExecuteGrantStmt_20))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_ExecuteGrantStmt_24), int32(0))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L9
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(_a_F_ExecuteGrantStmt_3), int32(581), int32(_a_F_ExecuteGrantStmt_20))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
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
	v673 = m.ExcPending
	if v673 != 0 {
		goto L9
	} else {
		goto L185
	}
L185:
	;
	v674 = F_privilege_to_string(m, v573)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L9
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v674
	F_errmsg(m, v527, v16+int32(32))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L9
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(_a_F_ExecuteGrantStmt_3), int32(587), int32(_a_F_ExecuteGrantStmt_20))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L9
	} else {
		goto L188
	}
L188:
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
	var v111 int32
	_ = v111
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int64
	_ = v140
	var v155 int32
	_ = v155
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int64
	_ = v289
	var v291 int64
	_ = v291
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int64
	_ = v311
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int64
	_ = v331
	var v332 int64
	_ = v332
	var v336 int64
	_ = v336
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int64
	_ = v346
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v358 int64
	_ = v358
	var v359 int64
	_ = v359
	var v363 int64
	_ = v363
	var v370 int32
	_ = v370
	var v372 int64
	_ = v372
	var v374 int64
	_ = v374
	var v377 int32
	_ = v377
	var v379 int64
	_ = v379
	var v382 int32
	_ = v382
	var v384 int64
	_ = v384
	var v413 int32
	_ = v413
	var v414 int64
	_ = v414
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v430 int64
	_ = v430
	var v434 int32
	_ = v434
	var v450 int32
	_ = v450
	var v451 int64
	_ = v451
	var v455 int64
	_ = v455
	var v460 int32
	_ = v460
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v478 int64
	_ = v478
	var v479 int64
	_ = v479
	var v483 int64
	_ = v483
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v495 int64
	_ = v495
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v507 int64
	_ = v507
	var v508 int64
	_ = v508
	var v512 int64
	_ = v512
	var v519 int32
	_ = v519
	var v521 int64
	_ = v521
	var v523 int64
	_ = v523
	var v526 int32
	_ = v526
	var v528 int64
	_ = v528
	var v531 int32
	_ = v531
	var v533 int64
	_ = v533
	var v562 int32
	_ = v562
	var v563 int64
	_ = v563
	var v567 int32
	_ = v567
	var v574 int32
	_ = v574
	var v579 int64
	_ = v579
	var v583 int32
	_ = v583
	var v599 int32
	_ = v599
	var v600 int64
	_ = v600
	var v604 int64
	_ = v604
	var v609 int32
	_ = v609
	var v617 int32
	_ = v617
	var v627 int32
	_ = v627
	var v636 int32
	_ = v636
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v684 int32
	_ = v684
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v744 int32
	_ = v744
	var v746 int64
	_ = v746
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v844 int32
	_ = v844
	var v852 int32
	_ = v852
	var v865 int32
	_ = v865
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v993 int32
	_ = v993
	var v1007 int64
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1033 int32
	_ = v1033
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
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
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1159 int32
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1200 int32
	_ = v1200
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1240 int32
	_ = v1240
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1257 int32
	_ = v1257
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1304 int32
	_ = v1304
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1348 int64
	_ = v1348
	var v1350 int64
	_ = v1350
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1386 int32
	_ = v1386
	var v1412 int32
	_ = v1412
	var v1415 int32
	_ = v1415
	var v1417 int32
	_ = v1417
	var v1421 int64
	_ = v1421
	var v1422 int64
	_ = v1422
	var v1426 int64
	_ = v1426
	var v1431 int32
	_ = v1431
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1442 int64
	_ = v1442
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1454 int64
	_ = v1454
	var v1455 int64
	_ = v1455
	var v1459 int64
	_ = v1459
	var v1466 int32
	_ = v1466
	var v1468 int64
	_ = v1468
	var v1470 int64
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1475 int64
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1480 int64
	_ = v1480
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1510 int64
	_ = v1510
	var v1514 int32
	_ = v1514
	var v1526 int64
	_ = v1526
	var v1530 int32
	_ = v1530
	var v1542 int32
	_ = v1542
	var v1547 int64
	_ = v1547
	var v1551 int64
	_ = v1551
	var v1556 int32
	_ = v1556
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1579 int32
	_ = v1579
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1594 int32
	_ = v1594
	var v1599 int32
	_ = v1599
	var v1623 int32
	_ = v1623
	var v1627 int32
	_ = v1627
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1650 int32
	_ = v1650
	var v1652 int32
	_ = v1652
	var v1677 int32
	_ = v1677
	var v1679 int64
	_ = v1679
	var v1691 int32
	_ = v1691
	var v1695 int32
	_ = v1695
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1722 int32
	_ = v1722
	var v1729 int32
	_ = v1729
	var v1733 int32
	_ = v1733
	var v1738 int32
	_ = v1738
	var v1742 int32
	_ = v1742
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1753 int32
	_ = v1753
	var v1761 int32
	_ = v1761
	var v1766 int32
	_ = v1766
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
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L11
	} else {
		goto L283
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L11
	} else {
		goto L278
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v1691
	m.G0 = v26 + int32(144)
	return v1695
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v32
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = v34
	v38 = m.G0
	v40 = v38 - int32(112)
	m.G0 = v40
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[0]))
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
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v781 = F_IOContextForStrategy(m, l2)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L11
	} else {
		goto L113
	}
L7:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	v1691 = v778
	v1695 = v137
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
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[1]))
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[2]))
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
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v137 = F_smgrnblocks(m, v136, l1)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L11
	} else {
		goto L26
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
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[3]))
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[4]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v95+(v88^int32(-1))<<(uint(int32(6))%32))+20))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v92+(int32(-2)-v101)<<(uint(int32(2))%32))))
	base.MemoryFill(m, v106, int32(0), int32(_a_F_ExtendBufferedRelCommon_0))
	v111 = v64 + int32(1)
	if v111 != v60 {
		v64 = v111
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v140 = base.I64_extend_i32_u(v60)
	if base.Ui64(base.I64_extend_i32_u(v137)+v140) <= base.Ui64(int64(4294967293)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v60 != 0 {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	goto L29
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L11
	} else {
		goto L108
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(48)))) = v60
	v744 = int32(_a_F_ExtendBufferedRelCommon_1)
	v746 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[5]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[5])) = v746 + v140
	m.G0 = v40 + int32(112)
	goto L7
L31:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[6])))
	v472 = m.G0
	v474 = v472 - int32(16)
	m.G0 = v474
	if v469 != 0 {
		goto L80
	} else {
		goto L81
	}
L32:
	;
	v155 = int32(0)
	goto L35
L33:
	;
	goto L34
L34:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[6])))
	v325 = m.G0
	v327 = v325 - int32(16)
	m.G0 = v327
	if v322 != 0 {
		goto L58
	} else {
		goto L59
	}
L35:
	;
	v169 = l6 + v155<<(uint(int32(2))%32)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[4]))
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[7]))
	F_ResourceOwnerEnlarge(m, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L11
	} else {
		goto L37
	}
L36:
	;
	goto L31
L37:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = v177
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+28)) = v179
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+40)) = v137 + v155
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v181
	v187 = v170 ^ int32(-1)
	v190 = v172 + v187<<(uint(int32(6))%32)
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[0]))
	v198 = F_hash_search(m, v192, v40+int32(24), int32(1), v40+int32(23))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+23)))
	if v200 == int32(1) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v305 = v299 + int32(36)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	goto L51
L40:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[8]))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v190)+20))
	v207 = int32(-2) - v206
	v210 = v204 + v207<<(uint(int32(2))%32)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v213 = v211 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v213
	if v213 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v190)+24))
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v40)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v190))) = v289
	v291 = *(*int64)(unsafe.Add(mBase, uint32(v40)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v190)+8)) = v291
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v190)+16)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v190)+24)) = v288 | int32(33816576)
	*(*int32)(unsafe.Add(mBase, uint32(v198)+20)) = v187
	v299 = v190
	goto L39
L43:
	;
	v217 = int32(_a_F_ExtendBufferedRelCommon_2)
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[2]))
	v220 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[2])) = v219 - v220
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[4]))
	v227 = v224 + v207<<(uint(int32(6))%32)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v227)+24)) = v228 - v220
	goto L45
L44:
	;
	goto L45
L45:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[7]))
	F_ResourceOwnerForget(m, v234, v206+int32(1), int32(_a_F_ExtendBufferedRelCommon_3))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L11
	} else {
		goto L46
	}
L46:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[4]))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	v245 = v241 + v242<<(uint(int32(6))%32)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+24))
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[8]))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v245)+20))
	v254 = v248 + (int32(-2)-v250)<<(uint(int32(2))%32)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	if v255 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v258 = int32(_a_F_ExtendBufferedRelCommon_2)
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[2]))
	v261 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[2])) = v260 + v261
	*(*int32)(unsafe.Add(mBase, uint32(v245)+24)) = v246 + v261
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	v268 = v267
	goto L49
L48:
	;
	v268 = v255
	goto L49
L49:
	;
	v269 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v254))) = v268 + v269
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[7]))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v245)+20))
	F_ResourceOwnerRemember(m, v273, v274+v269, int32(_a_F_ExtendBufferedRelCommon_3))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L11
	} else {
		goto L50
	}
L50:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v245)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = v280 + int32(1)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v245)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v245)+24)) = v284 & int32(-16777217)
	v299 = v245
	goto L39
L51:
	;
	if v306 != int32(-1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v305)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+104)) = v309
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v305)))
	*(*int64)(unsafe.Add(mBase, uint32(v40)+96)) = v311
	F_pgaio_wref_wait(m, v40+int32(96))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L11
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v319 = v155 + int32(1)
	if v60 != v319 {
		v155 = v319
		goto L35
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	goto L36
L57:
	;
	F_smgrzeroextend(m, v136, l1, v137, v60)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L11
	} else {
		goto L61
	}
L58:
	;
	F___clock_gettime(m, int32(1), v327)
	mBase = m.M
	v331 = int64(*(*int32)(unsafe.Add(mBase, uint32(v327)+8)))
	v332 = *(*int64)(unsafe.Add(mBase, uint32(v327)))
	v336 = v331 + v332*int64(1000000000)
	goto L60
L59:
	;
	v336 = int64(0)
	goto L60
L60:
	;
	m.G0 = v327 + int32(16)
	goto L57
L61:
	;
	v342 = int32(1)
	v346 = int64(0)
	v350 = m.G0
	v352 = v350 - int32(16)
	m.G0 = v352
	if v336 != v346 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L30
L63:
	;
	F___clock_gettime(m, int32(1), v352)
	mBase = m.M
	v358 = int64(*(*int32)(unsafe.Add(mBase, uint32(v352)+8)))
	v359 = *(*int64)(unsafe.Add(mBase, uint32(v352)))
	v363 = v358 + (v359*int64(1000000000) - v336)
	goto L67
L64:
	;
	goto L65
L65:
	;
	v450 = int32(552)
	v451 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[9]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[9])) = v451 + base.I64_extend_i32_u(v342)
	v455 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[10]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[10])) = v455 + v346
	F_pgstat_count_backend_io_op(m, v342, int32(3), int32(5), v342, v346)
	mBase = m.M
	v460 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[11])) = uint8(v460)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[12])) = uint8(v460)
	m.G0 = v352 + int32(16)
	goto L62
L66:
	;
	v413 = int32(552)
	v414 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[13]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[13])) = v414 + v363
	v418 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[14]))
	v425 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v418))|base.B2i32(int32(1)<<(uint(v418)%32)&int32(_a_F_ExtendBufferedRelCommon_4) == v425) == v425 {
		goto L76
	} else {
		goto L77
	}
L67:
	;
	goto L68
L68:
	;
	v370 = int32(_a_F_ExtendBufferedRelCommon_5)
	v372 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[15]))
	v374 = base.I64_div_s(v363, int64(1000))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[15])) = v372 + v374
	switch v342 {
	case 0:
		goto L72
	case 1:
		goto L71
	default:
		goto L66
	}
L71:
	;
	v382 = int32(_a_F_ExtendBufferedRelCommon_6)
	v384 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[16]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[16])) = v384 + v363
	goto L66
L72:
	;
	v377 = int32(_a_F_ExtendBufferedRelCommon_7)
	v379 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[17]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[17])) = v379 + v363
	goto L66
L76:
	;
	v430 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[18]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[18])) = v430 + v363
	v434 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[11])) = uint8(v434)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[19])) = uint8(v434)
	goto L78
L77:
	;
	goto L78
L78:
	;
	goto L65
L79:
	;
	F_smgrzeroextend(m, v136, l1, v137, v60)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L11
	} else {
		goto L83
	}
L80:
	;
	F___clock_gettime(m, int32(1), v474)
	mBase = m.M
	v478 = int64(*(*int32)(unsafe.Add(mBase, uint32(v474)+8)))
	v479 = *(*int64)(unsafe.Add(mBase, uint32(v474)))
	v483 = v478 + v479*int64(1000000000)
	goto L82
L81:
	;
	v483 = int64(0)
	goto L82
L82:
	;
	m.G0 = v474 + int32(16)
	goto L79
L83:
	;
	v489 = int32(1)
	v495 = base.I64_extend_i32_u(v60 << (uint(int32(13)) % 32))
	v499 = m.G0
	v501 = v499 - int32(16)
	m.G0 = v501
	if v483 != int64(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v617 = int32(0)
	if v60 != int32(1) {
		goto L101
	} else {
		goto L102
	}
L85:
	;
	F___clock_gettime(m, int32(1), v501)
	mBase = m.M
	v507 = int64(*(*int32)(unsafe.Add(mBase, uint32(v501)+8)))
	v508 = *(*int64)(unsafe.Add(mBase, uint32(v501)))
	v512 = v507 + (v508*int64(1000000000) - v483)
	goto L89
L86:
	;
	goto L87
L87:
	;
	v599 = int32(552)
	v600 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[9]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[9])) = v600 + base.I64_extend_i32_u(v489)
	v604 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[10]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[10])) = v604 + v495
	F_pgstat_count_backend_io_op(m, v489, int32(3), int32(5), v489, v495)
	mBase = m.M
	v609 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[11])) = uint8(v609)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[12])) = uint8(v609)
	m.G0 = v501 + int32(16)
	goto L84
L88:
	;
	v562 = int32(552)
	v563 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[13]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[13])) = v563 + v512
	v567 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[14]))
	v574 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v567))|base.B2i32(int32(1)<<(uint(v567)%32)&int32(_a_F_ExtendBufferedRelCommon_4) == v574) == v574 {
		goto L98
	} else {
		goto L99
	}
L89:
	;
	goto L90
L90:
	;
	v519 = int32(_a_F_ExtendBufferedRelCommon_5)
	v521 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[15]))
	v523 = base.I64_div_s(v512, int64(1000))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[15])) = v521 + v523
	switch v489 {
	case 0:
		goto L94
	case 1:
		goto L93
	default:
		goto L88
	}
L93:
	;
	v531 = int32(_a_F_ExtendBufferedRelCommon_6)
	v533 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[16]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[16])) = v533 + v512
	goto L88
L94:
	;
	v526 = int32(_a_F_ExtendBufferedRelCommon_7)
	v528 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[17]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[17])) = v528 + v512
	goto L88
L98:
	;
	v579 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[18]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[18])) = v579 + v512
	v583 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[11])) = uint8(v583)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[19])) = uint8(v583)
	goto L100
L99:
	;
	goto L100
L100:
	;
	goto L87
L101:
	;
	v627 = v617
	v636 = int32(0)
	goto L104
L102:
	;
	v684 = v617
	goto L103
L103:
	;
	v706 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[4]))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l6+v684<<(uint(int32(2))%32))))
	v715 = v706 + (v710^int32(-1))<<(uint(int32(6))%32)
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v715)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v715)+24)) = v716 | int32(16777216)
	goto L30
L104:
	;
	v648 = int32(_a_F_ExtendBufferedRelCommon_8)
	v649 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[4]))
	v650 = int32(2)
	v652 = l6 + v627<<(uint(v650)%32)
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v652)))
	v654 = int32(-1)
	v656 = int32(6)
	v658 = v649 + (v653^v654)<<(uint(v656)%32)
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v658)+24))
	v660 = int32(16777216)
	*(*int32)(unsafe.Add(mBase, uint32(v658)+24)) = v659 | v660
	v664 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[4]))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v652)+4))
	v670 = v664 + (v665^v654)<<(uint(v656)%32)
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v670)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+24)) = v671 | v660
	v676 = v627 + v650
	v678 = v636 + v650
	if v678 != v60&int32(-2) {
		v627 = v676
		v636 = v678
		goto L104
	} else {
		goto L106
	}
L105:
	;
	if v60&int32(1) == int32(0) {
		goto L30
	} else {
		goto L107
	}
L106:
	;
	goto L105
L107:
	;
	v684 = v676
	goto L103
L108:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L11
	} else {
		goto L109
	}
L109:
	;
	v760 = v40 + int32(24)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	F_GetRelationPath(m, v760, v761, v762, v763, v764, l1)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L11
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v760
	F_errmsg(m, int32(_a_F_ExtendBufferedRelCommon_9), v40)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L11
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_ExtendBufferedRelCommon_10), int32(395), int32(_a_F_ExtendBufferedRelCommon_11))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L11
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l4) {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v878 = l3 & int32(1)
	if v878 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L115:
	;
	v806 = int32(0)
	goto L129
L116:
	;
	v787 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[20]))
	v789 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[21]))
	v792 = v787 - v789 - int32(8)
	if base.Ui32(v792) <= base.Ui32(v787) {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	goto L118
L118:
	;
	if l4 != 0 {
		v804 = int32(1)
		goto L115
	} else {
		goto L128
	}
L119:
	;
	v795 = v792
	goto L121
L120:
	;
	v795 = int32(0)
	goto L121
L121:
	;
	if base.Ui32(v795) <= base.Ui32(int32(1)) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v798 = int32(1)
	goto L124
L123:
	;
	v798 = v795
	goto L124
L124:
	;
	if base.Ui32(v795) < base.Ui32(l4) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v800 = v798
	goto L127
L126:
	;
	v800 = l4
	goto L127
L127:
	;
	v804 = v800
	goto L115
L128:
	;
	v865 = int32(0)
	goto L114
L129:
	;
	v832 = F_GetVictimBuffer(m, l2, v781)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L11
	} else {
		goto L131
	}
L130:
	;
	v865 = v804
	goto L114
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6+v806<<(uint(int32(2))%32)))) = v832
	v836 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[22]))
	v838 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[23]))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v838+v832<<(uint(int32(6))%32)-int32(44))))
	base.MemoryFill(m, v836+v844<<(uint(int32(13))%32), int32(0), int32(_a_F_ExtendBufferedRelCommon_0))
	v852 = v806 + int32(1)
	if v852 != v804 {
		v806 = v852
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	F_LockRelationForExtension(m, v780, int32(7))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L11
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	if l3&int32(16) != 0 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	goto L135
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v779+l1<<(uint(int32(2))%32))+20)) = int32(-1)
	goto L139
L138:
	;
	goto L139
L139:
	;
	v891 = F_smgrnblocks(m, v779, l1)
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L11
	} else {
		goto L140
	}
L140:
	;
	if l5 == int32(-1) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v1007 = base.I64_extend_i32_u(v993)
	if base.Ui64(int64(4294967293)) < base.Ui64(v1007+base.I64_extend_i32_u(v891)) {
		goto L1
	} else {
		goto L163
	}
L142:
	;
	v993 = v865
	goto L141
L143:
	;
	goto L144
L144:
	;
	if base.Ui64(base.I64_extend_i32_u(l5)) < base.Ui64(base.I64_extend_i32_u(v891)+base.I64_extend_i32_u(v865)) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v901 = l5 - v891
	goto L147
L146:
	;
	v901 = v865
	goto L147
L147:
	;
	if base.Ui32(v891) <= base.Ui32(l5) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v904 = v901
	goto L150
L149:
	;
	v904 = int32(0)
	goto L150
L150:
	;
	if base.Ui32(v904) < base.Ui32(v865) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v906 = v904
	goto L154
L152:
	;
	goto L153
L153:
	;
	if v904 != 0 {
		v993 = v904
		goto L141
	} else {
		goto L160
	}
L154:
	;
	v930 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[23]))
	v934 = *(*int32)(unsafe.Add(mBase, uint32(l6+v906<<(uint(int32(2))%32))))
	v937 = v930 + v934<<(uint(int32(6))%32)
	v939 = v937 + int32(-64)
	F_StrategyFreeBuffer(m, v939)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L11
	} else {
		goto L156
	}
L155:
	;
	goto L153
L156:
	;
	v943 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[7]))
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v937-int32(44))))
	F_ResourceOwnerForget(m, v943, v946+int32(1), int32(_a_F_ExtendBufferedRelCommon_3))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L11
	} else {
		goto L157
	}
L157:
	;
	F_UnpinBufferNoOwner(m, v939)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L11
	} else {
		goto L158
	}
L158:
	;
	v955 = v906 + int32(1)
	if v955 != v865 {
		v906 = v955
		goto L154
	} else {
		goto L159
	}
L159:
	;
	goto L155
L160:
	;
	v980 = int32(0)
	if v878 != 0 {
		v1691 = v980
		v1695 = v891
		goto L3
	} else {
		goto L161
	}
L161:
	;
	F_UnlockRelationForExtension(m, v780, int32(7))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L11
	} else {
		goto L162
	}
L162:
	;
	v1691 = v980
	v1695 = v891
	goto L3
L163:
	;
	if v993 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v1012 = int32(-2113667072)
	if v29 == int32(112) {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	goto L166
L166:
	;
	v1412 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[6])))
	v1415 = m.G0
	v1417 = v1415 - int32(16)
	m.G0 = v1417
	if v1412 != 0 {
		goto L237
	} else {
		goto L238
	}
L167:
	;
	v1017 = v1012
	goto L169
L168:
	;
	v1017 = int32(33816576)
	goto L169
L169:
	;
	if l1 == int32(3) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v1020 = v1012
	goto L172
L171:
	;
	v1020 = v1017
	goto L172
L172:
	;
	v1033 = int32(0)
	goto L173
L173:
	;
	v1047 = l6 + v1033<<(uint(int32(2))%32)
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v1047)))
	v1050 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[23]))
	v1052 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[7]))
	F_ResourceOwnerEnlarge(m, v1052)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L11
	} else {
		goto L175
	}
L174:
	;
	goto L166
L175:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L11
	} else {
		goto L176
	}
L176:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v779)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v1057
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v779)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v1059
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v779)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v1033 + v891
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v1061
	v1067 = v26 + int32(52)
	v1068 = F_BufTableHashCode(m, v1067)
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L11
	} else {
		goto L177
	}
L177:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[24]))
	v1078 = v1071 + v1068&int32(127)<<(uint(int32(7))%32) + int32(_a_F_ExtendBufferedRelCommon_12)
	v1080 = F_LWLockAcquire(m, v1078, int32(0))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L11
	} else {
		goto L178
	}
L178:
	;
	v1084 = v1050 + v1048<<(uint(int32(6))%32)
	v1086 = v1084 + int32(-64)
	v1088 = v1084 - int32(44)
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1088)))
	v1090 = F_BufTableInsert(m, v1067, v1068, v1089)
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L11
	} else {
		goto L180
	}
L179:
	;
	v1386 = v1033 + int32(1)
	if v1386 != v993 {
		v1033 = v1386
		goto L173
	} else {
		goto L235
	}
L180:
	;
	if int32(0) <= v1090 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[23]))
	v1098 = v1095 + v1090<<(uint(int32(6))%32)
	v1099 = F_PinBuffer(m, v1098, l2)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L11
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+92)) = int32(_a_F_ExtendBufferedRelCommon_13)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+88)) = int32(_a_F_ExtendBufferedRelCommon_14)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+84)) = int32(_a_F_ExtendBufferedRelCommon_15)
	v1257 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+80)) = v1257
	*(*int64)(unsafe.Add(mBase, uint32(v26)+72)) = int64(0)
	v1262 = v1084 - int32(40)
	v1263 = int32(_a_F_ExtendBufferedRelCommon_16)
	v1265 = base.AtomicRmwOr32(m, v1262, v1257, v1263)
	if v1265&v1263 != 0 {
		goto L215
	} else {
		goto L216
	}
L184:
	;
	F_LWLockRelease(m, v1078)
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L11
	} else {
		goto L185
	}
L185:
	;
	F_StrategyFreeBuffer(m, v1086)
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L11
	} else {
		goto L186
	}
L186:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[7]))
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v1088)))
	F_ResourceOwnerForget(m, v1106, v1107+int32(1), int32(_a_F_ExtendBufferedRelCommon_3))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L11
	} else {
		goto L187
	}
L187:
	;
	F_UnpinBufferNoOwner(m, v1086)
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L11
	} else {
		goto L188
	}
L188:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1047))) = v1115 + int32(1)
	if v1099 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[22]))
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+20))
	v1125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1120+v1121<<(uint(int32(13))%32))+14)))
	if v1125 != 0 {
		goto L2
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	goto L193
L192:
	;
	goto L191
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+92)) = int32(_a_F_ExtendBufferedRelCommon_13)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+88)) = int32(_a_F_ExtendBufferedRelCommon_14)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+84)) = int32(_a_F_ExtendBufferedRelCommon_15)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+80)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+72)) = int64(0)
	v1159 = int32(_a_F_ExtendBufferedRelCommon_16)
	v1161 = base.AtomicRmwOr32(m, v1098, int32(24), v1159)
	if v1161&v1159 != 0 {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	goto L179
L195:
	;
	goto L198
L196:
	;
	v1200 = v1161
	goto L197
L197:
	;
	v1222 = int32(_a_F_ExtendBufferedRelCommon_17)
	v1223 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[25]))
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(72))+8))
	if v1225 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L198:
	;
	F_perform_spin_delay(m, v26+int32(72))
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L11
	} else {
		goto L200
	}
L199:
	;
	v1200 = v1193
	goto L197
L200:
	;
	v1191 = int32(_a_F_ExtendBufferedRelCommon_16)
	v1193 = base.AtomicRmwOr32(m, v1098, int32(24), v1191)
	if v1193&v1191 != 0 {
		goto L198
	} else {
		goto L201
	}
L201:
	;
	goto L199
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1098)+24)) = v1200 & int32(-20971521)
	v1247 = F_StartBufferIO(m, v1098, int32(1), int32(0))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L11
	} else {
		goto L213
	}
L203:
	;
	goto L202
L204:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[25])) = v1240
	goto L203
L205:
	;
	if int32(999) < v1223 {
		goto L203
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	if v1223 < int32(11) {
		goto L203
	} else {
		goto L212
	}
L208:
	;
	v1230 = int32(900)
	if v1230 <= v1223 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v1233 = v1230
	goto L211
L210:
	;
	v1233 = v1223
	goto L211
L211:
	;
	v1240 = v1233 + int32(100)
	goto L204
L212:
	;
	v1240 = v1223 - int32(1)
	goto L204
L213:
	;
	if v1247 == int32(0) {
		goto L193
	} else {
		goto L214
	}
L214:
	;
	goto L194
L215:
	;
	goto L218
L216:
	;
	v1304 = v1265
	goto L217
L217:
	;
	v1326 = int32(_a_F_ExtendBufferedRelCommon_17)
	v1327 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[25]))
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(72))+8))
	if v1329 == int32(0) {
		goto L225
	} else {
		goto L226
	}
L218:
	;
	F_perform_spin_delay(m, v26+int32(72))
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L11
	} else {
		goto L220
	}
L219:
	;
	v1304 = v1297
	goto L217
L220:
	;
	v1295 = int32(_a_F_ExtendBufferedRelCommon_16)
	v1297 = base.AtomicRmwOr32(m, v1262, int32(0), v1295)
	if v1297&v1295 != 0 {
		goto L218
	} else {
		goto L221
	}
L221:
	;
	goto L219
L222:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v1086)+16)) = v1346
	v1348 = *(*int64)(unsafe.Add(mBase, uint32(v26)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v1086)+8)) = v1348
	v1350 = *(*int64)(unsafe.Add(mBase, uint32(v26)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v1086))) = v1350
	*(*int32)(unsafe.Add(mBase, uint32(v1262))) = v1304&int32(-38010881) | v1020
	F_LWLockRelease(m, v1078)
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L11
	} else {
		goto L233
	}
L223:
	;
	goto L222
L224:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[25])) = v1344
	goto L223
L225:
	;
	if int32(999) < v1327 {
		goto L223
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	if v1327 < int32(11) {
		goto L223
	} else {
		goto L232
	}
L228:
	;
	v1334 = int32(900)
	if v1334 <= v1327 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1337 = v1334
	goto L231
L230:
	;
	v1337 = v1327
	goto L231
L231:
	;
	v1344 = v1337 + int32(100)
	goto L224
L232:
	;
	v1344 = v1327 - int32(1)
	goto L224
L233:
	;
	v1360 = F_StartBufferIO(m, v1086, int32(1), int32(0))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L11
	} else {
		goto L234
	}
L234:
	;
	goto L179
L235:
	;
	goto L174
L236:
	;
	F_smgrzeroextend(m, v779, l1, v891, v993)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L11
	} else {
		goto L240
	}
L237:
	;
	F___clock_gettime(m, int32(1), v1417)
	mBase = m.M
	v1421 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1417)+8)))
	v1422 = *(*int64)(unsafe.Add(mBase, uint32(v1417)))
	v1426 = v1421 + v1422*int64(1000000000)
	goto L239
L238:
	;
	v1426 = int64(0)
	goto L239
L239:
	;
	m.G0 = v1417 + int32(16)
	goto L236
L240:
	;
	if v878 == int32(0) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	F_UnlockRelationForExtension(m, v780, int32(7))
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L11
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v1437 = int32(0)
	v1439 = int32(1)
	v1442 = base.I64_extend_i32_u(v993 << (uint(int32(13)) % 32))
	v1446 = m.G0
	v1448 = v1446 - int32(16)
	m.G0 = v1448
	if v1426 != int64(0) {
		goto L246
	} else {
		goto L247
	}
L244:
	;
	goto L243
L245:
	;
	if v993 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L246:
	;
	F___clock_gettime(m, int32(1), v1448)
	mBase = m.M
	v1454 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1448)+8)))
	v1455 = *(*int64)(unsafe.Add(mBase, uint32(v1448)))
	v1459 = v1454 + (v1455*int64(1000000000) - v1426)
	goto L250
L247:
	;
	goto L248
L248:
	;
	v1542 = v781 << (uint(int32(6)) % 32)
	v1547 = *(*int64)(unsafe.Add(mBase, uint32(v1542)+uint32(_c_F_ExtendBufferedRelCommon[26])))
	*(*int64)(unsafe.Add(mBase, uint32(v1542)+uint32(_c_F_ExtendBufferedRelCommon[26]))) = v1547 + base.I64_extend_i32_u(v1439)
	v1551 = *(*int64)(unsafe.Add(mBase, uint32(v1542)+uint32(_c_F_ExtendBufferedRelCommon[27])))
	*(*int64)(unsafe.Add(mBase, uint32(v1542)+uint32(_c_F_ExtendBufferedRelCommon[27]))) = v1551 + v1442
	F_pgstat_count_backend_io_op(m, v1437, v781, int32(5), v1439, v1442)
	mBase = m.M
	v1556 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[11])) = uint8(v1556)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[12])) = uint8(v1556)
	m.G0 = v1448 + int32(16)
	goto L245
L249:
	;
	v1503 = int32(0)
	v1505 = v781 << (uint(int32(6)) % 32)
	v1510 = *(*int64)(unsafe.Add(mBase, uint32(v1505)+uint32(_c_F_ExtendBufferedRelCommon[28])))
	*(*int64)(unsafe.Add(mBase, uint32(v1505)+uint32(_c_F_ExtendBufferedRelCommon[28]))) = v1510 + v1459
	v1514 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[14]))
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v1514))|base.B2i32(int32(1)<<(uint(v1514)%32)&int32(_a_F_ExtendBufferedRelCommon_4) == v1503) == v1503 {
		goto L259
	} else {
		goto L260
	}
L250:
	;
	goto L251
L251:
	;
	v1466 = int32(_a_F_ExtendBufferedRelCommon_5)
	v1468 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[15]))
	v1470 = base.I64_div_s(v1459, int64(1000))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[15])) = v1468 + v1470
	switch v1437 {
	case 0:
		goto L255
	case 1:
		goto L254
	default:
		goto L249
	}
L254:
	;
	v1478 = int32(_a_F_ExtendBufferedRelCommon_6)
	v1480 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[16]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[16])) = v1480 + v1459
	goto L249
L255:
	;
	v1473 = int32(_a_F_ExtendBufferedRelCommon_7)
	v1475 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[17]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[17])) = v1475 + v1459
	goto L249
L259:
	;
	v1526 = *(*int64)(unsafe.Add(mBase, uint32(v1505)+uint32(_c_F_ExtendBufferedRelCommon[29])))
	*(*int64)(unsafe.Add(mBase, uint32(v1505)+uint32(_c_F_ExtendBufferedRelCommon[29]))) = v1526 + v1459
	v1530 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[11])) = uint8(v1530)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[19])) = uint8(v1530)
	goto L261
L260:
	;
	goto L261
L261:
	;
	goto L248
L262:
	;
	v1677 = int32(_a_F_ExtendBufferedRelCommon_18)
	v1679 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[30]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[30])) = v1679 + v1007
	v1691 = v993
	v1695 = v891
	goto L3
L263:
	;
	v1567 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[23]))
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v1571 = v1567 + v1568<<(uint(int32(6))%32)
	v1579 = v891 + int32(1)
	if l3&int32(8)|int32(base.Ui32(l3)>>(uint(int32(5))%32))&base.B2i32(v1579 == l5) != 0 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v1586 = F_LWLockAcquire(m, v1571-int32(16), int32(0))
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L11
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v1588 = int32(1)
	v1589 = int32(0)
	F_TerminateBufferIO(m, v1571+int32(-64), v1589, int32(16777216), v1588, v1589)
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L11
	} else {
		goto L268
	}
L267:
	;
	goto L266
L268:
	;
	if v993 == int32(1) {
		goto L262
	} else {
		goto L269
	}
L269:
	;
	v1599 = v1588
	goto L270
L270:
	;
	v1623 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[23]))
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(l6+v1599<<(uint(int32(2))%32))))
	v1630 = v1623 + v1627<<(uint(int32(6))%32)
	v1633 = int32(0)
	if base.B2i32(l3&int32(32) == v1633)|base.B2i32(v1599+v1579 != l5) == v1633 {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	goto L262
L272:
	;
	v1643 = F_LWLockAcquire(m, v1630-int32(16), int32(0))
	mBase = m.M
	v1644 = m.ExcPending
	if v1644 != 0 {
		goto L11
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	v1645 = int32(0)
	F_TerminateBufferIO(m, v1630+int32(-64), v1645, int32(16777216), int32(1), v1645)
	mBase = m.M
	v1650 = m.ExcPending
	if v1650 != 0 {
		goto L11
	} else {
		goto L276
	}
L275:
	;
	goto L274
L276:
	;
	v1652 = v1599 + int32(1)
	if v1652 != v993 {
		v1599 = v1652
		goto L270
	} else {
		goto L277
	}
L277:
	;
	goto L271
L278:
	;
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+16))
	v1716 = v26 + int32(72)
	v1717 = *(*int32)(unsafe.Add(mBase, uint32(v779)+4))
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v779)))
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v779)+8))
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v779)+12))
	F_GetRelationPath(m, v1716, v1717, v1718, v1719, v1720, l1)
	mBase = m.M
	v1722 = m.ExcPending
	if v1722 != 0 {
		goto L11
	} else {
		goto L279
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v1714
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v1716
	F_errmsg(m, int32(_a_F_ExtendBufferedRelCommon_19), v26+int32(32))
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		goto L11
	} else {
		goto L280
	}
L280:
	;
	F_errhint(m, int32(_a_F_ExtendBufferedRelCommon_20), int32(0))
	mBase = m.M
	v1733 = m.ExcPending
	if v1733 != 0 {
		goto L11
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(_a_F_ExtendBufferedRelCommon_15), int32(2777), int32(_a_F_ExtendBufferedRelCommon_21))
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L11
	} else {
		goto L282
	}
L282:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L283:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L11
	} else {
		goto L284
	}
L284:
	;
	v1747 = v26 + int32(72)
	v1748 = *(*int32)(unsafe.Add(mBase, uint32(v779)+4))
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v779)))
	v1750 = *(*int32)(unsafe.Add(mBase, uint32(v779)+8))
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v779)+12))
	F_GetRelationPath(m, v1747, v1748, v1749, v1750, v1751, l1)
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L11
	} else {
		goto L285
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v1747
	F_errmsg(m, int32(_a_F_ExtendBufferedRelCommon_9), v26+int32(16))
	mBase = m.M
	v1761 = m.ExcPending
	if v1761 != 0 {
		goto L11
	} else {
		goto L286
	}
L286:
	;
	F_errfinish(m, int32(_a_F_ExtendBufferedRelCommon_15), int32(2705), int32(_a_F_ExtendBufferedRelCommon_21))
	mBase = m.M
	v1766 = m.ExcPending
	if v1766 != 0 {
		goto L11
	} else {
		goto L287
	}
L287:
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
func F_ean13_out(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13929(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_ean2string(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v118 int32
	_ = v118
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
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
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
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
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v599 int32
	_ = v599
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v666 int32
	_ = v666
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
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
	var v876 int32
	_ = v876
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v931 int32
	_ = v931
	var v939 int32
	_ = v939
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v986 int32
	_ = v986
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1008 int32
	_ = v1008
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1053 int32
	_ = v1053
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1085 int32
	_ = v1085
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1130 int32
	_ = v1130
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1163 int32
	_ = v1163
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1190 int32
	_ = v1190
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1227 int32
	_ = v1227
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1274 int32
	_ = v1274
	var v1280 int32
	_ = v1280
	var v1299 int32
	_ = v1299
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1312 int32
	_ = v1312
	var v1318 int32
	_ = v1318
	var v1323 int32
	_ = v1323
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	v16 = int64(base.Ui64(l0) >> (uint(int64(1)) % 64))
	if base.Ui64(l0) <= base.Ui64(int64(19999999999999)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+18)) = uint8(v19)
	v21 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+15)) = uint8(v21)
	if l0&int64(1) == int64(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v16
	v1299 = v13 + int32(32)
	v1304 = F_pg_snprintf(m, v1299, int32(64), int32(_a_F_ean2string_0), v13+int32(16))
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L296
	} else {
		goto L297
	}
L4:
	;
	v29 = v19
	goto L6
L5:
	;
	v29 = int32(33)
	goto L6
L6:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)) = uint8(v29)
	v32 = base.I64_rem_u_s(v16, int64(10))
	v35 = base.I32_wrap_i64(v32) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v35)
	v38 = l1 + int32(15)
	if base.Ui64(int64(20)) <= base.Ui64(l0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v102 = int32(3)
	v104 = l1 + v102
	v107 = int32(0)
	v118 = int32(*(*int8)(unsafe.Add(mBase, uint32(v104))))
	goto L21
L8:
	;
	v42 = base.I64_div_u_s(l0, int64(20))
	v44 = v42
	v47 = int32(1)
	v49 = v38
	goto L11
L9:
	;
	v76 = int32(0)
	v78 = v38
	goto L10
L10:
	;
	v84 = int32(13) - v76
	if v84 == int32(0) {
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v55 = v49 - int32(1)
	v56 = int64(10)
	v57 = base.I64_div_u_s(v44, v56)
	v63 = base.I32_wrap_i64(v44-v57*v56) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v63)
	if base.Ui64(v44) < base.Ui64(v56) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if base.Ui32(int32(12)) < base.Ui32(v47) {
		goto L7
	} else {
		goto L16
	}
L13:
	;
	goto L12
L14:
	;
	v68 = v47 + int32(1)
	if v68 != int32(14) {
		v44 = v57
		v47 = v68
		v49 = v55
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L7
L16:
	;
	v76 = v47
	v78 = v55
	goto L10
L17:
	;
	base.MemoryFill(m, v76+v78-int32(13), int32(48), v84)
	goto L7
L18:
	;
	if l2 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L19:
	;
	v923 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v916))) = uint8(v923)
	v931 = v919
	goto L18
L20:
	;
	if v377 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L21:
	;
	goto L23
L23:
	;
	goto L25
L25:
	;
	goto L26
L26:
	;
	v170 = int32(_a_F_ean2string_1) + v118<<(uint(int32(3))%32)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v170-int32(380))))
	v174 = int32(1)
	v177 = int32(base.Ui32(v173+v174) >> (uint(v174) % 32))
	if v177 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v377 = int32(0)
	goto L20
L34:
	;
	goto L35
L35:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v170-int32(384))))
	v186 = v183 - int32(1)
	v187 = v186 + v177
	v189 = v187 << (uint(int32(3)) % 32)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v189)+uint32(_c_F_ean2string[0])))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v189)+uint32(_c_F_ean2string[1])))
	v196 = v104
	v197 = v192
	v198 = v191
	v200 = v107
	v202 = v187
	v203 = v177
	v204 = v186
	v205 = v107
	v206 = v183 + v173
	goto L36
L36:
	;
	v208 = int32(*(*int8)(unsafe.Add(mBase, uint32(v196))))
	if v200&int32(1) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v377 = int32(0)
	goto L20
L38:
	;
	if v355 != 0 {
		v196 = v349
		v197 = v360
		v198 = v350
		v200 = v352
		v202 = v354
		v203 = v355
		v204 = v356
		v205 = v357
		v206 = v358
		goto L36
	} else {
		goto L74
	}
L39:
	;
	v247 = v200 | base.B2i32(base.I32_extend8_s(v243) < v208)
	v248 = int32(1)
	v251 = base.B2i32(v208 < v244) | v205
	if v247&v248&(v251&v248) != 0 {
		goto L56
	} else {
		goto L57
	}
L40:
	;
	v229 = (v226 | v200) & int32(1)
	if v229 != 0 {
		goto L49
	} else {
		goto L50
	}
L41:
	;
	v214 = int32(*(*int8)(unsafe.Add(mBase, uint32(v197))))
	if v208 < v214 {
		v226 = int32(0)
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if v205&int32(1) != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	v218 = int32(*(*int8)(unsafe.Add(mBase, uint32(v198))))
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	v243 = v219
	v244 = v218
	goto L39
L46:
	;
	goto L47
L47:
	;
	v220 = int32(*(*int8)(unsafe.Add(mBase, uint32(v197))))
	v221 = int32(*(*int8)(unsafe.Add(mBase, uint32(v198))))
	if v208 <= v221 {
		v243 = v220
		v244 = v221
		goto L39
	} else {
		goto L48
	}
L48:
	;
	v226 = base.B2i32(v220 <= v208)
	goto L40
L49:
	;
	v230 = v206
	goto L51
L50:
	;
	v230 = v202
	goto L51
L51:
	;
	if v229 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v231 = v202
	goto L54
L53:
	;
	v231 = v204
	goto L54
L54:
	;
	v234 = int32(base.Ui32(v230-v231) >> (uint(int32(1)) % 32))
	v235 = v234 + v231
	v237 = v235 << (uint(int32(3)) % 32)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v237)+uint32(_c_F_ean2string[0])))
	v240 = int32(0)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v237)+uint32(_c_F_ean2string[1])))
	v349 = v104
	v350 = v239
	v352 = v240
	v354 = v235
	v355 = v234
	v356 = v231
	v357 = v240
	v358 = v230
	v360 = v242
	goto L38
L55:
	;
	v345 = int32(2)
	v349 = v264
	v350 = v198 + v345
	v352 = v247
	v354 = v202
	v355 = v203
	v356 = v204
	v357 = v251
	v358 = v206
	v360 = v197 + v345
	goto L38
L56:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v202<<(uint(int32(3))%32))+uint32(_c_F_ean2string[1])))
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282))))
	if v283 != 0 {
		goto L62
	} else {
		goto L63
	}
L57:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+1)))
	if v255 == int32(0) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v259 = v198 + int32(1)
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if v260 == int32(0) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v264 = v196 + int32(1)
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	if v265 == int32(0) {
		goto L56
	} else {
		goto L60
	}
L60:
	;
	if base.Ui32(int32(10)) <= base.Ui32((v255-int32(48))&int32(255)) {
		goto L55
	} else {
		goto L61
	}
L61:
	;
	v349 = v264
	v350 = v259
	v352 = v247
	v354 = v202
	v355 = v203
	v356 = v204
	v357 = v251
	v358 = v206
	v360 = v197 + int32(1)
	goto L38
L62:
	;
	v285 = l1
	v286 = v104
	v288 = v283
	v289 = v282
	v292 = int32(0)
	goto L65
L63:
	;
	v325 = l1
	v326 = v104
	v340 = int32(1)
	goto L64
L64:
	;
	v341 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v325))) = uint8(v341)
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326))))
	*(*uint8)(unsafe.Add(mBase, uint32(v325)+1)) = uint8(v343)
	v377 = v340
	goto L20
L65:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	if v300 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v325 = v316
	v326 = v317
	v340 = v321 + int32(1)
	goto L64
L67:
	;
	v301 = int32(45)
	v305 = base.B2i32(v288&int32(255) != v301)
	if v288&int32(255) != v301 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v316 = v285
	v317 = v286
	v321 = v292
	goto L69
L69:
	;
	goto L66
L70:
	;
	v306 = v300
	goto L72
L71:
	;
	v306 = v301
	goto L72
L72:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v285))) = uint8(v306)
	v308 = int32(1)
	v309 = v292 + v308
	v311 = v285 + v308
	v312 = v286 + v305
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+1)))
	if v313 != 0 {
		v285 = v311
		v286 = v312
		v288 = v313
		v289 = v289 + v308
		v292 = v309
		goto L65
	} else {
		goto L73
	}
L73:
	;
	v316 = v311
	v317 = v312
	v321 = v309
	goto L69
L74:
	;
	goto L37
L75:
	;
	v380 = int32(0)
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v381 == v380 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	v400 = int32(_a_F_ean2string_2)
	if v377 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L78:
	;
	v916 = l1
	v919 = v380
	goto L19
L79:
	;
	goto L80
L80:
	;
	v387 = l1
	v388 = v381
	v389 = v104
	goto L81
L81:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v387))) = uint8(v388)
	v395 = int32(1)
	v396 = v387 + v395
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+1)))
	if v397 != 0 {
		v387 = v396
		v388 = v397
		v389 = v389 + v395
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v916 = v396
	v919 = v380
	goto L19
L83:
	;
	goto L82
L84:
	;
	v620 = l1 + v377
	v622 = v620 + int32(2)
	v623 = int32(0)
	v634 = int32(*(*int8)(unsafe.Add(mBase, uint32(v622))))
	if v616 != 0 {
		goto L159
	} else {
		goto L160
	}
L85:
	;
	if v445 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L86:
	;
	v445 = int32(0)
	goto L85
L87:
	;
	goto L88
L88:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ean2string[2])))
	if v406 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v407 = v400
	v408 = l1
	v409 = v377
	v410 = v406
	goto L93
L90:
	;
	v433 = l1
	v437 = int32(0)
	goto L91
L91:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433))))
	v445 = v437 - v438
	goto L85
L92:
	;
	v433 = v428
	v437 = v430
	goto L91
L93:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408))))
	if base.B2i32(v410 != v412)|base.B2i32(v412 == int32(0)) != 0 {
		v428 = v408
		v430 = v410
		goto L92
	} else {
		goto L95
	}
L94:
	;
	v428 = v422
	v430 = int32(0)
	goto L92
L95:
	;
	v418 = v409 - int32(1)
	if v418 == int32(0) {
		v428 = v408
		v430 = v410
		goto L92
	} else {
		goto L96
	}
L96:
	;
	v421 = int32(1)
	v422 = v408 + v421
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407)+1)))
	if v423 != 0 {
		v407 = v407 + v421
		v408 = v422
		v409 = v418
		v410 = v423
		goto L93
	} else {
		goto L97
	}
L97:
	;
	goto L94
L98:
	;
	v616 = int32(_a_F_ean2string_3)
	v618 = v102
	v619 = int32(_a_F_ean2string_4)
	goto L84
L99:
	;
	goto L100
L100:
	;
	v450 = int32(_a_F_ean2string_5)
	if v377 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	if v495 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L102:
	;
	v495 = int32(0)
	goto L101
L103:
	;
	goto L104
L104:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ean2string[3])))
	if v456 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v457 = v450
	v458 = l1
	v459 = v377
	v460 = v456
	goto L109
L106:
	;
	v483 = l1
	v487 = int32(0)
	goto L107
L107:
	;
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
	v495 = v487 - v488
	goto L101
L108:
	;
	v483 = v478
	v487 = v480
	goto L107
L109:
	;
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458))))
	if base.B2i32(v460 != v462)|base.B2i32(v462 == int32(0)) != 0 {
		v478 = v458
		v480 = v460
		goto L108
	} else {
		goto L111
	}
L110:
	;
	v478 = v472
	v480 = int32(0)
	goto L108
L111:
	;
	v468 = v459 - int32(1)
	if v468 == int32(0) {
		v478 = v458
		v480 = v460
		goto L108
	} else {
		goto L112
	}
L112:
	;
	v471 = int32(1)
	v472 = v458 + v471
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457)+1)))
	if v473 != 0 {
		v457 = v457 + v471
		v458 = v472
		v459 = v468
		v460 = v473
		goto L109
	} else {
		goto L113
	}
L113:
	;
	goto L110
L114:
	;
	v616 = int32(_a_F_ean2string_6)
	v618 = int32(5)
	v619 = int32(_a_F_ean2string_7)
	goto L84
L115:
	;
	goto L116
L116:
	;
	v501 = int32(_a_F_ean2string_8)
	v503 = v377 + int32(1)
	if v503 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	if v548 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L118:
	;
	v548 = int32(0)
	goto L117
L119:
	;
	goto L120
L120:
	;
	v509 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ean2string[4])))
	if v509 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v510 = v501
	v511 = l1
	v512 = v503
	v513 = v509
	goto L125
L122:
	;
	v536 = l1
	v540 = int32(0)
	goto L123
L123:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v536))))
	v548 = v540 - v541
	goto L117
L124:
	;
	v536 = v531
	v540 = v533
	goto L123
L125:
	;
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511))))
	if base.B2i32(v513 != v515)|base.B2i32(v515 == int32(0)) != 0 {
		v531 = v511
		v533 = v513
		goto L124
	} else {
		goto L127
	}
L126:
	;
	v531 = v525
	v533 = int32(0)
	goto L124
L127:
	;
	v521 = v512 - int32(1)
	if v521 == int32(0) {
		v531 = v511
		v533 = v513
		goto L124
	} else {
		goto L128
	}
L128:
	;
	v524 = int32(1)
	v525 = v511 + v524
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+1)))
	if v526 != 0 {
		v510 = v510 + v524
		v511 = v525
		v512 = v521
		v513 = v526
		goto L125
	} else {
		goto L129
	}
L129:
	;
	goto L126
L130:
	;
	v616 = int32(_a_F_ean2string_9)
	v618 = int32(4)
	v619 = int32(_a_F_ean2string_10)
	goto L84
L131:
	;
	goto L132
L132:
	;
	v554 = int32(_a_F_ean2string_11)
	if v377 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	if v599 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L134:
	;
	v599 = int32(0)
	goto L133
L135:
	;
	goto L136
L136:
	;
	v560 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ean2string[5])))
	if v560 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v561 = v554
	v562 = l1
	v563 = v377
	v564 = v560
	goto L141
L138:
	;
	v587 = l1
	v591 = int32(0)
	goto L139
L139:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587))))
	v599 = v591 - v592
	goto L133
L140:
	;
	v587 = v582
	v591 = v584
	goto L139
L141:
	;
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562))))
	if base.B2i32(v564 != v566)|base.B2i32(v566 == int32(0)) != 0 {
		v582 = v562
		v584 = v564
		goto L140
	} else {
		goto L143
	}
L142:
	;
	v582 = v576
	v584 = int32(0)
	goto L140
L143:
	;
	v572 = v563 - int32(1)
	if v572 == int32(0) {
		v582 = v562
		v584 = v564
		goto L140
	} else {
		goto L144
	}
L144:
	;
	v575 = int32(1)
	v576 = v562 + v575
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561)+1)))
	if v577 != 0 {
		v561 = v561 + v575
		v562 = v576
		v563 = v572
		v564 = v577
		goto L141
	} else {
		goto L145
	}
L145:
	;
	goto L142
L146:
	;
	v616 = int32(_a_F_ean2string_12)
	v618 = v102
	v619 = int32(_a_F_ean2string_13)
	goto L84
L147:
	;
	goto L148
L148:
	;
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v608 = base.B2i32(v606 == int32(48))
	if v606 == int32(48) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v609 = int32(_a_F_ean2string_14)
	goto L151
L150:
	;
	v609 = int32(0)
	goto L151
L151:
	;
	if v606 == int32(48) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v612 = int32(6)
	goto L154
L153:
	;
	v612 = int32(2)
	goto L154
L154:
	;
	if v606 == int32(48) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v615 = int32(_a_F_ean2string_15)
	goto L157
L156:
	;
	v615 = int32(0)
	goto L157
L157:
	;
	v616 = v609
	v618 = v612
	v619 = v615
	goto L84
L158:
	;
	if v893 != 0 {
		v931 = v618
		goto L18
	} else {
		goto L213
	}
L159:
	;
	v636 = v619
	goto L161
L160:
	;
	v636 = v623
	goto L161
L161:
	;
	if v636 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	if v634 != 0 {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	goto L164
L164:
	;
	v686 = v616 + v634<<(uint(int32(3))%32)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v686-int32(380))))
	v690 = int32(1)
	v693 = int32(base.Ui32(v689+v690) >> (uint(v690) % 32))
	if v693 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L165:
	;
	v640 = v620
	v641 = v622
	v643 = int32(0)
	v644 = v634
	goto L168
L166:
	;
	v666 = v620
	v681 = int32(1)
	goto L167
L167:
	;
	v682 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v666))) = uint8(v682)
	v893 = v681
	goto L158
L168:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v640))) = uint8(v644)
	v656 = int32(1)
	v659 = v640 + v656
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+1)))
	if v660 != 0 {
		v640 = v659
		v641 = v641 + v656
		v643 = v643 + v656
		v644 = v660
		goto L168
	} else {
		goto L170
	}
L169:
	;
	v666 = v659
	v681 = v643 + int32(2)
	goto L167
L170:
	;
	goto L169
L171:
	;
	v893 = int32(0)
	goto L158
L172:
	;
	goto L173
L173:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v686-int32(384))))
	v702 = v699 - int32(1)
	v703 = v702 + v693
	v706 = v619 + v703<<(uint(int32(3))%32)
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v706)+4))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	v712 = v622
	v713 = v708
	v714 = v707
	v716 = v623
	v718 = v703
	v719 = v693
	v720 = v702
	v721 = v623
	v722 = v699 + v689
	goto L174
L174:
	;
	v724 = int32(*(*int8)(unsafe.Add(mBase, uint32(v712))))
	if v716&int32(1) == int32(0) {
		goto L179
	} else {
		goto L180
	}
L175:
	;
	v893 = int32(0)
	goto L158
L176:
	;
	if v871 != 0 {
		v712 = v865
		v713 = v876
		v714 = v866
		v716 = v868
		v718 = v870
		v719 = v871
		v720 = v872
		v721 = v873
		v722 = v874
		goto L174
	} else {
		goto L212
	}
L177:
	;
	v763 = v716 | base.B2i32(base.I32_extend8_s(v759) < v724)
	v764 = int32(1)
	v767 = base.B2i32(v724 < v760) | v721
	if v763&v764&(v767&v764) != 0 {
		goto L194
	} else {
		goto L195
	}
L178:
	;
	v745 = (v742 | v716) & int32(1)
	if v745 != 0 {
		goto L187
	} else {
		goto L188
	}
L179:
	;
	v730 = int32(*(*int8)(unsafe.Add(mBase, uint32(v713))))
	if v724 < v730 {
		v742 = int32(0)
		goto L178
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	if v721&int32(1) != 0 {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	goto L181
L183:
	;
	v734 = int32(*(*int8)(unsafe.Add(mBase, uint32(v714))))
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713))))
	v759 = v735
	v760 = v734
	goto L177
L184:
	;
	goto L185
L185:
	;
	v736 = int32(*(*int8)(unsafe.Add(mBase, uint32(v713))))
	v737 = int32(*(*int8)(unsafe.Add(mBase, uint32(v714))))
	if v724 <= v737 {
		v759 = v736
		v760 = v737
		goto L177
	} else {
		goto L186
	}
L186:
	;
	v742 = base.B2i32(v736 <= v724)
	goto L178
L187:
	;
	v746 = v722
	goto L189
L188:
	;
	v746 = v718
	goto L189
L189:
	;
	if v745 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v747 = v718
	goto L192
L191:
	;
	v747 = v720
	goto L192
L192:
	;
	v750 = int32(base.Ui32(v746-v747) >> (uint(int32(1)) % 32))
	v751 = v750 + v747
	v754 = v619 + v751<<(uint(int32(3))%32)
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v754)+4))
	v756 = int32(0)
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v754)))
	v865 = v622
	v866 = v755
	v868 = v756
	v870 = v751
	v871 = v750
	v872 = v747
	v873 = v756
	v874 = v746
	v876 = v758
	goto L176
L193:
	;
	v861 = int32(2)
	v865 = v780
	v866 = v714 + v861
	v868 = v763
	v870 = v718
	v871 = v719
	v872 = v720
	v873 = v767
	v874 = v722
	v876 = v713 + v861
	goto L176
L194:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v619+v718<<(uint(int32(3))%32))))
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798))))
	if v799 != 0 {
		goto L200
	} else {
		goto L201
	}
L195:
	;
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713)+1)))
	if v771 == int32(0) {
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v775 = v714 + int32(1)
	v776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775))))
	if v776 == int32(0) {
		goto L194
	} else {
		goto L197
	}
L197:
	;
	v780 = v712 + int32(1)
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v780))))
	if v781 == int32(0) {
		goto L194
	} else {
		goto L198
	}
L198:
	;
	if base.Ui32(int32(10)) <= base.Ui32((v771-int32(48))&int32(255)) {
		goto L193
	} else {
		goto L199
	}
L199:
	;
	v865 = v780
	v866 = v775
	v868 = v763
	v870 = v718
	v871 = v719
	v872 = v720
	v873 = v767
	v874 = v722
	v876 = v713 + int32(1)
	goto L176
L200:
	;
	v801 = v620
	v802 = v622
	v804 = v799
	v805 = v798
	v808 = int32(0)
	goto L203
L201:
	;
	v841 = v620
	v842 = v622
	v856 = int32(1)
	goto L202
L202:
	;
	v857 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v841))) = uint8(v857)
	v859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842))))
	*(*uint8)(unsafe.Add(mBase, uint32(v841)+1)) = uint8(v859)
	v893 = v856
	goto L158
L203:
	;
	v816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v802))))
	if v816 != 0 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v841 = v832
	v842 = v833
	v856 = v837 + int32(1)
	goto L202
L205:
	;
	v817 = int32(45)
	v821 = base.B2i32(v804&int32(255) != v817)
	if v804&int32(255) != v817 {
		goto L208
	} else {
		goto L209
	}
L206:
	;
	v832 = v801
	v833 = v802
	v837 = v808
	goto L207
L207:
	;
	goto L204
L208:
	;
	v822 = v816
	goto L210
L209:
	;
	v822 = v817
	goto L210
L210:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v801))) = uint8(v822)
	v824 = int32(1)
	v825 = v808 + v824
	v827 = v801 + v824
	v828 = v802 + v821
	v829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v805)+1)))
	if v829 != 0 {
		v801 = v827
		v802 = v828
		v804 = v829
		v805 = v805 + v824
		v808 = v825
		goto L203
	} else {
		goto L211
	}
L211:
	;
	v832 = v827
	v833 = v828
	v837 = v825
	goto L207
L212:
	;
	goto L175
L213:
	;
	v894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v622))))
	if v894 == int32(0) {
		v916 = v620
		v919 = v618
		goto L19
	} else {
		goto L214
	}
L214:
	;
	v900 = v620
	v901 = v894
	v902 = v622
	goto L215
L215:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v900))) = uint8(v901)
	v908 = int32(1)
	v909 = v900 + v908
	v910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902)+1)))
	if v910 != 0 {
		v900 = v909
		v901 = v910
		v902 = v902 + v908
		goto L215
	} else {
		goto L217
	}
L216:
	;
	v916 = v909
	v919 = v618
	goto L19
L217:
	;
	goto L216
L218:
	;
	m.G0 = v13 + int32(96)
	return
L219:
	;
	switch v931 - int32(3) {
	case 0:
		goto L223
	case 1:
		goto L222
	case 2:
		goto L221
	case 3:
		goto L220
	default:
		goto L218
	}
L220:
	;
	v1244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v1244 != 0 {
		goto L287
	} else {
		goto L288
	}
L221:
	;
	v1141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v1141 != 0 {
		goto L268
	} else {
		goto L269
	}
L222:
	;
	v1108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v1108 != 0 {
		goto L262
	} else {
		goto L263
	}
L223:
	;
	v939 = int32(_a_F_ean2string_2)
	goto L226
L224:
	;
	if v977-v978 != 0 {
		goto L218
	} else {
		goto L237
	}
L226:
	;
	goto L227
L227:
	;
	v946 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ean2string[2])))
	if v946 != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v947 = v939
	v948 = l1
	v949 = int32(4)
	v950 = v946
	goto L232
L229:
	;
	v973 = l1
	v977 = int32(0)
	goto L230
L230:
	;
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v973))))
	goto L224
L231:
	;
	v973 = v968
	v977 = v970
	goto L230
L232:
	;
	v952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948))))
	if base.B2i32(v950 != v952)|base.B2i32(v952 == int32(0)) != 0 {
		v968 = v948
		v970 = v950
		goto L231
	} else {
		goto L234
	}
L233:
	;
	v968 = v962
	v970 = int32(0)
	goto L231
L234:
	;
	v958 = v949 - int32(1)
	if v958 == int32(0) {
		v968 = v948
		v970 = v950
		goto L231
	} else {
		goto L235
	}
L235:
	;
	v961 = int32(1)
	v962 = v948 + v961
	v963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v947)+1)))
	if v963 != 0 {
		v947 = v947 + v961
		v948 = v962
		v949 = v958
		v950 = v963
		goto L232
	} else {
		goto L236
	}
L236:
	;
	goto L233
L237:
	;
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v986 != 0 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v992 = l1
	v993 = l1 + int32(4)
	v994 = v986
	goto L241
L239:
	;
	v1008 = l1
	goto L240
L240:
	;
	v1015 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1008))) = uint8(v1015)
	v1018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v1018 == v1015 {
		goto L245
	} else {
		goto L246
	}
L241:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v992))) = uint8(v994)
	v1000 = int32(1)
	v1001 = v992 + v1000
	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v993)+1)))
	if v1002 != 0 {
		v992 = v1001
		v993 = v993 + v1000
		v994 = v1002
		goto L241
	} else {
		goto L243
	}
L242:
	;
	v1008 = v1001
	goto L240
L243:
	;
	goto L242
L244:
	;
	v1080 = F_strlen(m, l1)
	mBase = m.M
	v1085 = v1080 + l1
	goto L256
L245:
	;
	v1079 = int32(0)
	goto L244
L246:
	;
	v1025 = int32(10)
	v1026 = v1015
	v1027 = l1
	v1028 = v1018
	goto L247
L247:
	;
	v1035 = (v1028 - int32(48)) & int32(255)
	v1039 = base.B2i32(base.Ui32(v1035) < base.Ui32(int32(10)))
	if base.Ui32(v1035) < base.Ui32(int32(10)) {
		goto L250
	} else {
		goto L251
	}
L248:
	;
	v1053 = base.I32_rem_u_s(v1041, int32(11))
	if v1053 == int32(0) {
		goto L245
	} else {
		goto L255
	}
L249:
	;
	goto L248
L250:
	;
	v1040 = v1025 * v1035
	goto L252
L251:
	;
	v1040 = int32(0)
	goto L252
L252:
	;
	v1041 = v1040 + v1026
	v1042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1027)+1)))
	if v1042 == int32(0) {
		goto L249
	} else {
		goto L253
	}
L253:
	;
	v1045 = int32(1)
	v1047 = v1025 - v1039
	if base.Ui32(v1045) < base.Ui32(v1047) {
		v1025 = v1047
		v1026 = v1041
		v1027 = v1027 + v1045
		v1028 = v1042
		goto L247
	} else {
		goto L254
	}
L254:
	;
	goto L249
L255:
	;
	v1079 = int32(11) - v1053
	goto L244
L256:
	;
	v1093 = v1085 - int32(1)
	v1094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1093))))
	if base.Ui32((v1094-int32(58))&int32(255)) < base.Ui32(int32(246)) {
		v1085 = v1093
		goto L256
	} else {
		goto L258
	}
L257:
	;
	if v1079 == int32(10) {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	goto L257
L259:
	;
	v1106 = int32(88)
	goto L261
L260:
	;
	v1106 = v1079 | int32(48)
	goto L261
L261:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1093))) = uint8(v1106)
	goto L218
L262:
	;
	v1114 = l1
	v1115 = l1 + int32(4)
	v1116 = v1108
	goto L265
L263:
	;
	v1130 = l1
	goto L264
L264:
	;
	v1137 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1130))) = uint8(v1137)
	v1139 = int32(77)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v1139)
	goto L218
L265:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1114))) = uint8(v1116)
	v1122 = int32(1)
	v1123 = v1114 + v1122
	v1124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1115)+1)))
	if v1124 != 0 {
		v1114 = v1123
		v1115 = v1115 + v1122
		v1116 = v1124
		goto L265
	} else {
		goto L267
	}
L266:
	;
	v1130 = v1123
	goto L264
L267:
	;
	goto L266
L268:
	;
	v1147 = l1
	v1148 = l1 + int32(4)
	v1149 = v1141
	goto L271
L269:
	;
	v1163 = l1
	goto L270
L270:
	;
	v1170 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1163))) = uint8(v1170)
	v1173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v1173 == v1170 {
		v1227 = v1170
		goto L275
	} else {
		goto L276
	}
L271:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1147))) = uint8(v1149)
	v1155 = int32(1)
	v1156 = v1147 + v1155
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1148)+1)))
	if v1157 != 0 {
		v1147 = v1156
		v1148 = v1148 + v1155
		v1149 = v1157
		goto L271
	} else {
		goto L273
	}
L272:
	;
	v1163 = v1156
	goto L270
L273:
	;
	goto L272
L274:
	;
	v1241 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)) = uint8(v1241)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v1240)
	goto L218
L275:
	;
	v1240 = v1227 | int32(48)
	goto L274
L276:
	;
	v1180 = int32(8)
	v1181 = v1170
	v1182 = l1
	v1183 = v1173
	goto L277
L277:
	;
	v1190 = (v1183 - int32(48)) & int32(255)
	v1194 = base.B2i32(base.Ui32(v1190) < base.Ui32(int32(10)))
	if base.Ui32(v1190) < base.Ui32(int32(10)) {
		goto L280
	} else {
		goto L281
	}
L278:
	;
	v1207 = int32(0)
	v1209 = base.I32_rem_u_s(v1196, int32(11))
	if v1209 == v1207 {
		v1227 = v1207
		goto L275
	} else {
		goto L285
	}
L279:
	;
	goto L278
L280:
	;
	v1195 = v1180 * v1190
	goto L282
L281:
	;
	v1195 = int32(0)
	goto L282
L282:
	;
	v1196 = v1195 + v1181
	v1197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1182)+1)))
	if v1197 == int32(0) {
		goto L279
	} else {
		goto L283
	}
L283:
	;
	v1200 = int32(1)
	v1202 = v1180 - v1194
	if base.Ui32(v1200) < base.Ui32(v1202) {
		v1180 = v1202
		v1181 = v1196
		v1182 = v1182 + v1200
		v1183 = v1197
		goto L277
	} else {
		goto L284
	}
L284:
	;
	goto L279
L285:
	;
	if v1209 == int32(1) {
		v1240 = int32(88)
		goto L274
	} else {
		goto L286
	}
L286:
	;
	v1227 = int32(11) - v1209
	goto L275
L287:
	;
	v1250 = v1244
	v1251 = l1
	v1252 = l1 + int32(1)
	goto L290
L288:
	;
	v1274 = l1
	goto L289
L289:
	;
	v1280 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1274))) = uint8(v1280)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v1280)
	goto L218
L290:
	;
	if base.Ui32((v1250-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	v1274 = v1266
	goto L289
L292:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1251))) = uint8(v1250)
	v1266 = v1251 + int32(1)
	goto L294
L293:
	;
	v1266 = v1251
	goto L294
L294:
	;
	v1267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1252)+1)))
	if v1267 != 0 {
		v1250 = v1267
		v1251 = v1266
		v1252 = v1252 + int32(1)
		goto L290
	} else {
		goto L295
	}
L295:
	;
	goto L291
L296:
	;
	return
L297:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L296
	} else {
		goto L298
	}
L298:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L296
	} else {
		goto L299
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(_a_F_ean2string_16)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v1299
	F_errmsg(m, int32(_a_F_ean2string_17), v13)
	mBase = m.M
	v1318 = m.ExcPending
	if v1318 != 0 {
		goto L296
	} else {
		goto L300
	}
L300:
	;
	F_errfinish(m, int32(_a_F_ean2string_18), int32(674), int32(_a_F_ean2string_19))
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L296
	} else {
		goto L301
	}
L301:
	;
	base.Wasm_trap_unreachable()
	for {
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
	var v2 int32
	_ = v2
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	if v3 == int32(457) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v2)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		v12 = F_find_simplified_clause(m, v6, v10, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v18 = v12
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
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v429 int32
	_ = v429
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = int32(_a_F_elog_node_display_0)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_elog_node_display[0])))
	v23 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_elog_node_display[0])) = uint8(v23)
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
	*(*uint8)(unsafe.Add(mBase, _c_F_elog_node_display[0])) = uint8(v31)
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
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L107
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
	v297 = m.G0
	v299 = v297 - int32(128)
	m.G0 = v299
	F_initStringInfo(m, v299+int32(32))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L78
	}
L8:
	;
	v48 = v37
	v52 = v37
	v53 = v4
	goto L9
L9:
	;
	v58 = int32(0)
	if v52 <= v58 {
		v67 = v58
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if int32(0) < v261 {
		goto L73
	} else {
		goto L74
	}
L11:
	;
	v269 = v41 + int32(128)
	v271 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v269+v261))) = uint8(v271)
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258+v33))))
	if v274 != 0 {
		goto L69
	} else {
		goto L70
	}
L12:
	;
	v69 = v48
	v72 = v67
	v73 = v52
	v74 = v53
	goto L18
L13:
	;
	if v52 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	base.MemoryFill(m, v41+int32(128), int32(32), v52)
	goto L16
L15:
	;
	goto L16
L16:
	;
	if base.Ui32(v52) <= base.Ui32(int32(77)) {
		v67 = v52
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v258 = v48
	v261 = v52
	v262 = v52
	v263 = v53
	goto L11
L18:
	;
	v79 = v69 + v33
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v80 == int32(0) {
		v258 = v69
		v261 = v72
		v262 = v73
		v263 = v74
		goto L11
	} else {
		goto L20
	}
L19:
	;
	v258 = v252
	v261 = v254
	v262 = v245
	v263 = v246
	goto L11
L20:
	;
	v85 = v41 + int32(128) + v72
	*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v80)
	switch v80 - int32(41) {
	case 0:
		goto L24
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16:
		v241 = v69
		v244 = v72
		v245 = v73
		v246 = v74
		goto L21
	case 17:
		goto L22
	default:
		goto L25
	}
L21:
	;
	v251 = int32(1)
	v252 = v241 + v251
	v254 = v244 + v251
	if v254 < int32(78) {
		v69 = v252
		v72 = v254
		v73 = v245
		v74 = v246
		goto L18
	} else {
		goto L68
	}
L22:
	;
	v223 = v41 + int32(128)
	if v72 != v73 {
		goto L64
	} else {
		goto L65
	}
L23:
	;
	if v72 != v73 {
		goto L49
	} else {
		goto L50
	}
L24:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	if v148 == int32(41) {
		v241 = v69
		v244 = v72
		v245 = v73
		v246 = v74
		goto L21
	} else {
		goto L44
	}
L25:
	;
	switch v80 - int32(123) {
	case 0:
		goto L23
	default:
		v241 = v69
		v244 = v72
		v245 = v73
		v246 = v74
		goto L21
	case 2:
		goto L26
	}
L26:
	;
	if v72 != v73 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v92 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v92)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v41 + int32(128)
	F_appendStringInfo(m, v41+int32(112), int32(_a_F_elog_node_display_1), v41+int32(48))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v105 = v41 + int32(128)
	v107 = int32(125)
	*(*uint16)(unsafe.Add(mBase, uint32(v105+v73))) = uint16(v107)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v105
	F_appendStringInfo(m, v41+int32(112), int32(_a_F_elog_node_display_1), v41+int32(32))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v118 = v69
	goto L32
L32:
	;
	v129 = v118 + int32(1)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v129))))
	if v131 == int32(32) {
		v118 = v129
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v135 = v74 - int32(1)
	v137 = base.B2i32(int32(0) < v74)
	if int32(0) < v74 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	v138 = v135
	goto L37
L36:
	;
	v138 = v74
	goto L37
L37:
	;
	v139 = int32(60)
	v141 = v135 * int32(3)
	if v139 <= v141 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v144 = v139
	goto L40
L39:
	;
	v144 = v141
	goto L40
L40:
	;
	if int32(0) < v74 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v145 = v144
	goto L43
L42:
	;
	v145 = v73
	goto L43
L43:
	;
	v241 = v118
	v244 = v145 - int32(1)
	v245 = v145
	v246 = v138
	goto L21
L44:
	;
	v151 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)) = uint8(v151)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v41 + int32(128)
	F_appendStringInfo(m, v41+int32(112), int32(_a_F_elog_node_display_1), v41-int32(-64))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v164 = v69
	goto L46
L46:
	;
	v175 = v164 + int32(1)
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v175))))
	if v177 == int32(32) {
		v164 = v175
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v241 = v164
	v244 = v73 - int32(1)
	v245 = v73
	v246 = v74
	goto L21
L48:
	;
	goto L47
L49:
	;
	v183 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v183)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v41 + int32(128)
	F_appendStringInfo(m, v41+int32(112), int32(_a_F_elog_node_display_1), v41+int32(80))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v195 = int32(60)
	v197 = v74 + int32(1)
	v199 = v197 * int32(3)
	if v195 <= v199 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L51
L53:
	;
	v202 = v195
	goto L55
L54:
	;
	v202 = v199
	goto L55
L55:
	;
	if v199 <= int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(128)+v216))) = uint8(v220)
	v241 = v69
	v244 = v216
	v245 = v202
	v246 = v197
	goto L21
L57:
	;
	v216 = int32(0)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v206 = int32(1)
	if v202 <= v206 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v209 = v206
	goto L62
L61:
	;
	v209 = v202
	goto L62
L62:
	;
	if v209 == int32(0) {
		v216 = v209
		goto L56
	} else {
		goto L63
	}
L63:
	;
	base.MemoryFill(m, v41+int32(128), int32(32), v209)
	v216 = v209
	goto L56
L64:
	;
	v226 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v226)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+96)) = v223
	F_appendStringInfo(m, v41+int32(112), int32(_a_F_elog_node_display_1), v41+int32(96))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	v238 = int32(58)
	goto L66
L66:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v223+v73))) = uint8(v238)
	v241 = v69
	v244 = v73
	v245 = v73
	v246 = v74
	goto L21
L67:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	v238 = v236
	goto L66
L68:
	;
	goto L19
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v269
	F_appendStringInfo(m, v41+int32(112), int32(_a_F_elog_node_display_1), v41+int32(16))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	goto L10
L72:
	;
	v48 = v258
	v52 = v262
	v53 = v263
	goto L9
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v41 + int32(128)
	F_appendStringInfo(m, v41+int32(112), int32(_a_F_elog_node_display_1), v41)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v41)+112))
	m.G0 = v41 + int32(208)
	v441 = v293
	goto L4
L76:
	;
	goto L75
L77:
	;
	v441 = v398
	goto L4
L78:
	;
	v312 = v4
	v313 = v4
	goto L79
L79:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v313))))
	if v317 != 0 {
		goto L85
	} else {
		goto L86
	}
L81:
	;
	v416 = int32(0)
	v418 = v299 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v418+v410))) = uint8(v416)
	*(*int32)(unsafe.Add(mBase, uint32(v299)+16)) = v418
	F_appendStringInfo(m, v299+int32(32), int32(_a_F_elog_node_display_1), v299+int32(16))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L106
	}
L82:
	;
	v410 = int32(78)
	v413 = v313 + int32(2)
	goto L81
L83:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v299)+32))
	m.G0 = v299 + int32(128)
	goto L77
L84:
	;
	v385 = v299 + int32(48)
	v387 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v385+v382))) = uint8(v387)
	*(*int32)(unsafe.Add(mBase, uint32(v299))) = v385
	F_appendStringInfo(m, v299+int32(32), int32(_a_F_elog_node_display_1), v299)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L105
	}
L85:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v299+int32(48)+v312))) = uint8(v317)
	v322 = int32(1)
	v323 = v313 + v322
	v325 = v312 + v322
	if v325 != int32(78) {
		v312 = v325
		v313 = v323
		goto L79
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	if v312 == int32(0) {
		goto L83
	} else {
		goto L104
	}
L88:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323+v33))))
	if v329 == int32(32) {
		goto L82
	} else {
		goto L89
	}
L89:
	;
	v332 = int32(78)
	if v329 == int32(0) {
		v382 = v332
		goto L84
	} else {
		goto L90
	}
L90:
	;
	v342 = v332
	goto L91
L91:
	;
	v347 = v342 - int32(1)
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347+(v299+int32(48))))))
	if v351 == int32(32) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v410 = v374
	v413 = v373 + v313 - int32(77)
	goto L81
L93:
	;
	goto L92
L94:
	;
	v373 = v342
	v374 = v347
	goto L93
L95:
	;
	goto L96
L96:
	;
	v355 = v342 - int32(2)
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355+(v299+int32(48))))))
	if v359 == int32(32) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v373 = v347
	v374 = v355
	goto L93
L98:
	;
	goto L99
L99:
	;
	if v342 < int32(4) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v410 = int32(78)
	v413 = v323
	goto L81
L101:
	;
	goto L102
L102:
	;
	v366 = v342 - int32(3)
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299+int32(48)+v366))))
	if v370 != int32(32) {
		v342 = v366
		goto L91
	} else {
		goto L103
	}
L103:
	;
	v373 = v355
	v374 = v366
	goto L93
L104:
	;
	v382 = v312
	goto L84
L105:
	;
	goto L83
L106:
	;
	v312 = v416
	v313 = v413
	goto L79
L107:
	;
	v446 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	if v446 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_elog_node_display_2), v14+int32(16))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	F_pfree(m, v441)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L115
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v441
	F_errdetail_internal(m, int32(_a_F_elog_node_display_3), v14)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_elog_node_display_4), int32(85), int32(_a_F_elog_node_display_5))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	goto L111
L115:
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
					F_errmsg(m, int32(_a_F_enlargeStringInfo_0), v7+int32(32))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v70
						F_errdetail(m, int32(_a_F_enlargeStringInfo_1), v7+int32(16))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_enlargeStringInfo_2), int32(364), int32(_a_F_enlargeStringInfo_3))
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
			F_errmsg_internal(m, int32(_a_F_enlargeStringInfo_4), v7)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_enlargeStringInfo_2), int32(351), int32(_a_F_enlargeStringInfo_3))
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_err_generic_string[0]))
	if int32(0) <= v11 {
		switch l0 - int32(99) {
		case 0:
			v34 = int32(68)
			v36 = v11 * int32(100)
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+uint32(_c_F_err_generic_string[1])))
			v41 = F_MemoryContextStrdup(m, v40, l1)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v36+int32(_a_F_err_generic_string_0)+v34))) = v41
				m.G0 = v8 + int32(16)
				return
			}
		case 1:
			v34 = int32(72)
			v36 = v11 * int32(100)
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+uint32(_c_F_err_generic_string[1])))
			v41 = F_MemoryContextStrdup(m, v40, l1)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v36+int32(_a_F_err_generic_string_0)+v34))) = v41
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
				F_errmsg_internal(m, int32(_a_F_err_generic_string_1), v8)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_err_generic_string_2), int32(1559), int32(_a_F_err_generic_string_3))
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
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+uint32(_c_F_err_generic_string[1])))
			v41 = F_MemoryContextStrdup(m, v40, l1)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v36+int32(_a_F_err_generic_string_0)+v34))) = v41
				m.G0 = v8 + int32(16)
				return
			}
		case 16:
			v34 = int32(60)
			v36 = v11 * int32(100)
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+uint32(_c_F_err_generic_string[1])))
			v41 = F_MemoryContextStrdup(m, v40, l1)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v36+int32(_a_F_err_generic_string_0)+v34))) = v41
				m.G0 = v8 + int32(16)
				return
			}
		case 17:
			v34 = int32(64)
			v36 = v11 * int32(100)
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+uint32(_c_F_err_generic_string[1])))
			v41 = F_MemoryContextStrdup(m, v40, l1)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v36+int32(_a_F_err_generic_string_0)+v34))) = v41
				m.G0 = v8 + int32(16)
				return
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_err_generic_string[0])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_err_generic_string_4), int32(0))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_err_generic_string_2), int32(1539), int32(_a_F_err_generic_string_3))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
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
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
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
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v241 int32
	_ = v241
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	v3 = int32(0)
	if l0 < int32(21) {
		v112 = l0
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[0]))
		if v14 != 0 {
			v15 = int32(23)
		} else {
			v15 = l0
		}
		if v15 == int32(21) {
			v18 = int32(22)
			v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_errstart[1])))
			if v22&int32(1) != 0 {
				v25 = v18
			} else {
				v25 = int32(21)
			}
			v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_errstart[2])))
			if v27&int32(1) != 0 {
				v30 = v18
			} else {
				v30 = v25
			}
			v33 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[3]))
			if v33 != 0 {
				v34 = v30
			} else {
				v34 = int32(22)
			}
			v35 = v34
		} else {
			v35 = v15
		}
		v37 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[4]))
		if v37 < int32(0) {
			v112 = v35
		} else {
			v40 = int32(1)
			v42 = v37 + v40
			if v42 <= v40 {
				v45 = v40
			} else {
				v45 = v42
			}
			v47 = v45 & int32(3)
			if v42 < int32(4) {
				v84 = v35
				v87 = int32(0)
				v93 = v84
				v96 = v87
				v101 = v3
				for {
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v96*int32(100))+uint32(_c_F_errstart[5])))
					if v104 < v93 {
						v106 = v93
					} else {
						v106 = v104
					}
					v107 = int32(1)
					v110 = v101 + v107
					if v110 != v47 {
						v93 = v106
						v96 = v96 + v107
						v101 = v110
						continue
					} else {
						break
					}
					break
				}
				v112 = v106
			} else {
				v54 = v35
				v57 = int32(0)
				v61 = v3
				for {
					v64 = v57 * int32(100)
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+uint32(_c_F_errstart[5])))
					if v65 < v54 {
						v67 = v54
					} else {
						v67 = v65
					}
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+uint32(_c_F_errstart[6])))
					if v68 < v67 {
						v70 = v67
					} else {
						v70 = v68
					}
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)+uint32(_c_F_errstart[7])))
					if v71 < v70 {
						v73 = v70
					} else {
						v73 = v71
					}
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v64)+uint32(_c_F_errstart[8])))
					if v74 < v73 {
						v76 = v73
					} else {
						v76 = v74
					}
					v77 = int32(4)
					v78 = v57 + v77
					v80 = v61 + v77
					if v80 != v45&int32(2147483644) {
						v54 = v76
						v57 = v78
						v61 = v80
						continue
					} else {
						break
					}
					break
				}
				if v47 == int32(0) {
					v112 = v76
				} else {
					v84 = v76
					v87 = v78
					v93 = v84
					v96 = v87
					v101 = v3
					for {
						v104 = *(*int32)(unsafe.Add(mBase, uint32(v96*int32(100))+uint32(_c_F_errstart[5])))
						if v104 < v93 {
							v106 = v93
						} else {
							v106 = v104
						}
						v107 = int32(1)
						v110 = v101 + v107
						if v110 != v47 {
							v93 = v106
							v96 = v96 + v107
							v101 = v110
							continue
						} else {
							break
						}
						break
					}
					v112 = v106
				}
			}
		}
	}
	v121 = int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[9]))
	if base.Ui32(v112-int32(15)) <= base.Ui32(v121) {
		if v123 < int32(22) {
			v138 = v121
		} else {
			v138 = int32(0)
		}
	} else {
		if v112 == int32(20) {
			v138 = int32(0)
		} else {
			if v123 == int32(15) {
				if v112 <= int32(21) {
					v138 = int32(0)
				} else {
					v138 = v121
				}
			} else {
				if v112 < v123 {
					v138 = int32(0)
				} else {
					v138 = v121
				}
			}
		}
	}
	v139 = int32(0)
	if v112 == int32(16) {
		v158 = v139
	} else {
		v143 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[10]))
		if v143 != int32(2) {
			v158 = v139
		} else {
			v147 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_errstart[11])))
			if v147 == int32(1) {
				v158 = base.B2i32(int32(20) < v112)
			} else {
				v155 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[12]))
				v158 = base.B2i32(v112 == int32(17)) | base.B2i32(v155 <= v112)
			}
		}
	}
	v164 = (base.B2i32(int32(20) < v112) | v138 | v158) & int32(1)
	if v164 != 0 {
		v166 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[13]))
		if v166 == int32(0) {
			F_write_stderr(m, int32(_a_F_errstart_0), int32(0))
			mBase = m.M
			v256 = m.ExcPending
			if v256 != 0 {
				return int32(0)
			} else {
				F_pgl_exit(m, int32(2))
				mBase = m.M
				v259 = m.ExcPending
				if v259 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		} else {
			v169 = int32(_a_F_errstart_1)
			v171 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[14]))
			v173 = v171 + int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_errstart[14])) = v173
			if base.B2i32(v112 < int32(21))|base.B2i32(v171 <= int32(0)) != 0 {
				v194 = v173
				v195 = int32(_a_F_errstart_2)
				v197 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[4]))
				v199 = v197 + int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_errstart[4])) = v199
				if int32(5) <= v199 {
					*(*int32)(unsafe.Add(mBase, _c_F_errstart[4])) = int32(-1)
					F_errstart_cold(m, int32(23), int32(0))
					mBase = m.M
					v266 = m.ExcPending
					if v266 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_errstart_3), int32(0))
						mBase = m.M
						v270 = m.ExcPending
						if v270 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_errstart_4), int32(762), int32(_a_F_errstart_5))
							mBase = m.M
							v275 = m.ExcPending
							if v275 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v203 = int32(100)
					v204 = v199 * v203
					base.MemoryFill(m, v204+int32(_a_F_errstart_6), int32(0), v203)
					v211 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[15]))
					if l1 != 0 {
						v215 = l1
					} else {
						v215 = int32(_a_F_errstart_7)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[16]))) = v215
					*(*int32)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[17]))) = v215
					*(*uint8)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[18]))) = uint8(v158)
					*(*uint8)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[19]))) = uint8(v138)
					*(*int32)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[5]))) = v112
					*(*int32)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[20]))) = v211
					if int32(21) <= v112 {
						*(*int32)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[21]))) = int32(2600)
					} else {
						if int32(19) <= v112 {
							*(*int32)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[21]))) = int32(64)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[21]))) = int32(0)
						}
					}
					v241 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[13]))
					*(*int32)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[22]))) = v241
					*(*int32)(unsafe.Add(mBase, _c_F_errstart[14])) = v194 - int32(1)
					return v164
				}
			} else {
				F_MemoryContextReset(m, v166)
				mBase = m.M
				v183 = m.ExcPending
				if v183 != 0 {
					return int32(0)
				} else {
					v185 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[14]))
					if v185 < int32(3) {
						v194 = v185
					} else {
						v189 = int32(0)
						*(*int32)(unsafe.Add(mBase, _c_F_errstart[23])) = v189
						*(*int32)(unsafe.Add(mBase, _c_F_errstart[24])) = v189
						v194 = v185
					}
					v195 = int32(_a_F_errstart_2)
					v197 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[4]))
					v199 = v197 + int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_errstart[4])) = v199
					if int32(5) <= v199 {
						*(*int32)(unsafe.Add(mBase, _c_F_errstart[4])) = int32(-1)
						F_errstart_cold(m, int32(23), int32(0))
						mBase = m.M
						v266 = m.ExcPending
						if v266 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_errstart_3), int32(0))
							mBase = m.M
							v270 = m.ExcPending
							if v270 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_errstart_4), int32(762), int32(_a_F_errstart_5))
								mBase = m.M
								v275 = m.ExcPending
								if v275 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v203 = int32(100)
						v204 = v199 * v203
						base.MemoryFill(m, v204+int32(_a_F_errstart_6), int32(0), v203)
						v211 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[15]))
						if l1 != 0 {
							v215 = l1
						} else {
							v215 = int32(_a_F_errstart_7)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[16]))) = v215
						*(*int32)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[17]))) = v215
						*(*uint8)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[18]))) = uint8(v158)
						*(*uint8)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[19]))) = uint8(v138)
						*(*int32)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[5]))) = v112
						*(*int32)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[20]))) = v211
						if int32(21) <= v112 {
							*(*int32)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[21]))) = int32(2600)
						} else {
							if int32(19) <= v112 {
								*(*int32)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[21]))) = int32(64)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[21]))) = int32(0)
							}
						}
						v241 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[13]))
						*(*int32)(unsafe.Add(mBase, uint32(v204)+uint32(_c_F_errstart[22]))) = v241
						*(*int32)(unsafe.Add(mBase, _c_F_errstart[14])) = v194 - int32(1)
						return v164
					}
				}
			}
		}
	} else {
		return v164
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
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
	var v161 int32
	_ = v161
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v391 int32
	_ = v391
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v412 int32
	_ = v412
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v600 int32
	_ = v600
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v725 float64
	_ = v725
	var v728 int32
	_ = v728
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
		v728 = v33
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
	return v728
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v37 <= int32(0) {
		v728 = v33
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v46 = v5
	v48 = v5
	v52 = v5
	v53 = v5
	v54 = v5
	goto L8
L8:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v52<<(uint(int32(2))%32))))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+16)))
	if v63 != int32(100) {
		v208 = v46
		v210 = v48
		v215 = v53
		v216 = v54
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v224 = int32(0)
	if v215 == v224 {
		v728 = v224
		goto L5
	} else {
		goto L39
	}
L10:
	;
	v221 = v52 + int32(1)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v221 < v222 {
		v46 = v208
		v48 = v210
		v52 = v221
		v53 = v215
		v54 = v216
		goto L8
	} else {
		goto L38
	}
L11:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+20)))
	if v66 != v67 {
		v208 = v46
		v210 = v48
		v215 = v53
		v216 = v54
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
	if (base.B2i32(v46 != v182)|base.B2i32(v180 <= v48))&base.B2i32(v182 <= v46)|base.B2i32(v180+v182 < int32(2)) != 0 {
		v208 = v46
		v210 = v48
		v215 = v53
		v216 = v54
		goto L10
	} else {
		goto L37
	}
L14:
	;
	v72 = int32(0)
	v180 = v72
	v182 = v72
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
		v182 = v74
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v84 = v74
	v86 = v74
	v90 = v74
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
	v182 = v161
	goto L13
L20:
	;
	v173 = v90 + int32(1)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v173 < v174 {
		v84 = v159
		v86 = v161
		v90 = v173
		goto L18
	} else {
		goto L36
	}
L21:
	;
	v106 = int32(*(*int16)(unsafe.Add(mBase, uint32(v102)+8)))
	if v106 <= int32(0) {
		v159 = v84
		v161 = v86
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
		v161 = v86
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
	v161 = v86
	goto L20
L27:
	;
	v118 = int32(0)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v119 <= v118 {
		v159 = v84
		v161 = v86
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
	v161 = v86 + int32(1)
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
	v161 = v86
	goto L20
L36:
	;
	goto L19
L37:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v208 = v182
	v210 = v180
	v215 = v202
	v216 = v62
	goto L10
L38:
	;
	goto L9
L39:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+20)))
	v228 = m.G0
	v230 = v228 - int32(32)
	m.G0 = v230
	v233 = F_SearchSysCache2(m, int32(62), v215, v227)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L25
	} else {
		goto L42
	}
L40:
	;
	if v246 == int32(0) {
		v728 = v224
		goto L5
	} else {
		goto L57
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L25
	} else {
		goto L54
	}
L42:
	;
	if v233 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v239 = F_SysCacheGetAttr(m, int32(62), v233, int32(3), v230+int32(31))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L25
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L25
	} else {
		goto L51
	}
L46:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+31)))
	if v241 == int32(1) {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v244 = F_pg_detoast_datum_packed(m, v239)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L25
	} else {
		goto L48
	}
L48:
	;
	v246 = F_statext_ndistinct_deserialize(m, v244)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L25
	} else {
		goto L49
	}
L49:
	;
	F_ReleaseCatCache(m, v233)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L25
	} else {
		goto L50
	}
L50:
	;
	m.G0 = v230 + int32(32)
	goto L40
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230))) = v215
	F_errmsg_internal(m, int32(_a_F_estimate_multivariate_ndistinct_0), v230)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L25
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_estimate_multivariate_ndistinct_1), int32(158), int32(_a_F_estimate_multivariate_ndistinct_2))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L25
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230)+20)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v230)+16)) = int32(100)
	F_errmsg_internal(m, int32(_a_F_estimate_multivariate_ndistinct_3), v230+int32(16))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L25
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_estimate_multivariate_ndistinct_1), int32(165), int32(_a_F_estimate_multivariate_ndistinct_2))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L25
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
	v285 = int32(0)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v216)+24))
	if v287 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	v289 = int32(16)
	v295 = (v288<<(uint(v289)%32) + int32(_a_F_estimate_multivariate_ndistinct_4)) >> (uint(v289) % 32)
	goto L60
L59:
	;
	v295 = v285
	goto L60
L60:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v296 == int32(0) {
		v412 = v285
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v246)+8))
	if v421 != 0 {
		goto L86
	} else {
		goto L87
	}
L62:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v296)+4))
	if v299 <= int32(0) {
		v412 = v285
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v311 = v285
	v313 = int32(0)
	goto L64
L64:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v296)+12))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v320+v313<<(uint(int32(2))%32))))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	if v326 != int32(6) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v412 = v391
	goto L61
L66:
	;
	v401 = v313 + int32(1)
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v296)+4))
	if v401 < v402 {
		v311 = v391
		v313 = v401
		goto L64
	} else {
		goto L84
	}
L67:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v216)+24))
	if v329 == int32(0) {
		v391 = v311
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v371 = int32(*(*int16)(unsafe.Add(mBase, uint32(v325)+8)))
	if v371 <= int32(0) {
		v391 = v311
		goto L66
	} else {
		goto L80
	}
L70:
	;
	v332 = int32(0)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	if v333 <= v332 {
		v391 = v311
		goto L66
	} else {
		goto L71
	}
L71:
	;
	v337 = v332
	goto L72
L72:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v354+v337<<(uint(int32(2))%32))))
	v359 = F_equal(m, v353, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L25
	} else {
		goto L74
	}
L73:
	;
	v391 = v311
	goto L66
L74:
	;
	if v359 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v365 = F_bms_add_member(m, v311, base.I32_extend16_s(v295+(v337^int32(-1))))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L25
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v368 = v337 + int32(1)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	if v368 < v369 {
		v337 = v368
		goto L72
	} else {
		goto L79
	}
L78:
	;
	v391 = v365
	goto L66
L79:
	;
	goto L73
L80:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v216)+20))
	v375 = F_bms_is_member(m, v371, v374)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L25
	} else {
		goto L81
	}
L81:
	;
	if v375 == int32(0) {
		v391 = v311
		goto L66
	} else {
		goto L82
	}
L82:
	;
	v381 = F_bms_add_member(m, v311, base.I32_extend16_s(v371+v295))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L25
	} else {
		goto L83
	}
L83:
	;
	v391 = v381
	goto L66
L84:
	;
	goto L65
L85:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v587 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L86:
	;
	v425 = int32(0)
	goto L89
L87:
	;
	goto L88
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L25
	} else {
		goto L113
	}
L89:
	;
	v444 = v246 + int32(16) + v425<<(uint(int32(4))%32)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v444)+8))
	v446 = int32(0)
	if v412 == v446 {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	goto L88
L91:
	;
	v537 = v425 + int32(1)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v246)+8))
	if base.Ui32(v537) < base.Ui32(v538) {
		v425 = v537
		goto L89
	} else {
		goto L112
	}
L92:
	;
	if v445 != v481 {
		goto L91
	} else {
		goto L105
	}
L93:
	;
	v481 = int32(0)
	goto L92
L94:
	;
	goto L95
L95:
	;
	v453 = int32(1)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v412)+4))
	if v454 <= v453 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v457 = v453
	goto L98
L97:
	;
	v457 = v454
	goto L98
L98:
	;
	v461 = int32(0)
	v463 = v446
	goto L99
L99:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v412+int32(8)+v461<<(uint(int32(2))%32))))
	if v469 != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v481 = v472
	goto L92
L101:
	;
	v472 = v463 + base.I32_popcnt(v469)
	goto L103
L102:
	;
	v472 = v463
	goto L103
L103:
	;
	v474 = v461 + int32(1)
	if v474 != v457 {
		v461 = v474
		v463 = v472
		goto L99
	} else {
		goto L104
	}
L104:
	;
	goto L100
L105:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v444)+8))
	if v483 <= int32(0) {
		goto L85
	} else {
		goto L106
	}
L106:
	;
	v488 = int32(0)
	goto L107
L107:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v444)+12))
	v508 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v504+v488<<(uint(int32(1))%32)))))
	v511 = F_bms_is_member(m, base.I32_extend16_s(v508+v295), v412)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L25
	} else {
		goto L109
	}
L108:
	;
	goto L85
L109:
	;
	if v511 == int32(0) {
		goto L91
	} else {
		goto L110
	}
L110:
	;
	v516 = v488 + int32(1)
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v444)+8))
	if v516 < v517 {
		v488 = v516
		goto L107
	} else {
		goto L111
	}
L111:
	;
	goto L108
L112:
	;
	goto L90
L113:
	;
	F_errmsg_internal(m, int32(_a_F_estimate_multivariate_ndistinct_5), int32(0))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L25
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_estimate_multivariate_ndistinct_6), int32(_a_F_estimate_multivariate_ndistinct_7), int32(_a_F_estimate_multivariate_ndistinct_8))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L25
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
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v711
	v725 = *(*float64)(unsafe.Add(mBase, uint32(v444)))
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = v725
	v728 = int32(1)
	goto L5
L117:
	;
	v711 = int32(0)
	goto L116
L118:
	;
	goto L119
L119:
	;
	v591 = int32(0)
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v587)+4))
	if v592 <= v591 {
		v711 = v591
		goto L116
	} else {
		goto L120
	}
L120:
	;
	v600 = v591
	v607 = int32(0)
	goto L121
L121:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v587)+12))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v613+v607<<(uint(int32(2))%32))))
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v617)))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v618)))
	if v619 == int32(6) {
		goto L125
	} else {
		goto L126
	}
L122:
	;
	v711 = v690
	goto L116
L123:
	;
	v704 = v607 + int32(1)
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v587)+4))
	if v704 < v705 {
		v600 = v690
		v607 = v704
		goto L121
	} else {
		goto L139
	}
L124:
	;
	v684 = F_lappend(m, v600, v617)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L25
	} else {
		goto L138
	}
L125:
	;
	v622 = int32(*(*int16)(unsafe.Add(mBase, uint32(v618)+8)))
	if v622 <= int32(0) {
		goto L124
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v216)+24))
	if v631 == int32(0) {
		goto L124
	} else {
		goto L131
	}
L128:
	;
	v627 = F_bms_is_member(m, base.I32_extend16_s(v622+v295), v412)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L25
	} else {
		goto L129
	}
L129:
	;
	if v627 == int32(0) {
		goto L124
	} else {
		goto L130
	}
L130:
	;
	v690 = v600
	goto L123
L131:
	;
	v634 = int32(0)
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v631)+4))
	if v635 <= v634 {
		goto L124
	} else {
		goto L132
	}
L132:
	;
	v639 = v634
	goto L133
L133:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v617)))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v631)+12))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v656+v639<<(uint(int32(2))%32))))
	v661 = F_equal(m, v655, v660)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L25
	} else {
		goto L135
	}
L134:
	;
	goto L124
L135:
	;
	if v661 != 0 {
		v690 = v600
		goto L123
	} else {
		goto L136
	}
L136:
	;
	v664 = v639 + int32(1)
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v631)+4))
	if v664 < v665 {
		v639 = v664
		goto L133
	} else {
		goto L137
	}
L137:
	;
	goto L134
L138:
	;
	v690 = v684
	goto L123
L139:
	;
	goto L122
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
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
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int64
	_ = v131
	var v133 int32
	_ = v133
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
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
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
					v67 = v20
					v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					if v69 != v67 {
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+14)))
						F_assign_simple_var(m, l0, l1, v67, v68&int32(1), base.B2i32(v79|v68 == int32(0)))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return
						} else {
							m.G0 = v11 + int32(80)
							return
						}
					} else {
						v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
						if v71 != 0 {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+14)))
							F_assign_simple_var(m, l0, l1, v67, v68&int32(1), base.B2i32(v79|v68 == int32(0)))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return
							} else {
								m.G0 = v11 + int32(80)
								return
							}
						} else {
							if v68&int32(1) == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
								m.G0 = v11 + int32(80)
								return
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+14)))
								F_assign_simple_var(m, l0, l1, v67, v68&int32(1), base.B2i32(v79|v68 == int32(0)))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return
								} else {
									m.G0 = v11 + int32(80)
									return
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v35
							F_errmsg(m, int32(_a_F_exec_assign_value_1), v11+int32(16))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_3), int32(_a_F_exec_assign_value_4))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
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
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+14)))
				if v48 != 0 {
					v67 = v20
					v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					if v69 != v67 {
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+14)))
						F_assign_simple_var(m, l0, l1, v67, v68&int32(1), base.B2i32(v79|v68 == int32(0)))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return
						} else {
							m.G0 = v11 + int32(80)
							return
						}
					} else {
						v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
						if v71 != 0 {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+14)))
							F_assign_simple_var(m, l0, l1, v67, v68&int32(1), base.B2i32(v79|v68 == int32(0)))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return
							} else {
								m.G0 = v11 + int32(80)
								return
							}
						} else {
							if v68&int32(1) == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
								m.G0 = v11 + int32(80)
								return
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+14)))
								F_assign_simple_var(m, l0, l1, v67, v68&int32(1), base.B2i32(v79|v68 == int32(0)))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return
								} else {
									m.G0 = v11 + int32(80)
									return
								}
							}
						}
					}
				} else {
					v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+20)))
					if v49 != int32(1) {
						v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+12)))
						v64 = F_datumTransfer(m, v20, int32(0), v63)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							v67 = v64
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
							if v69 != v67 {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+14)))
								F_assign_simple_var(m, l0, l1, v67, v68&int32(1), base.B2i32(v79|v68 == int32(0)))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return
								} else {
									m.G0 = v11 + int32(80)
									return
								}
							} else {
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
								if v71 != 0 {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+14)))
									F_assign_simple_var(m, l0, l1, v67, v68&int32(1), base.B2i32(v79|v68 == int32(0)))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return
									} else {
										m.G0 = v11 + int32(80)
										return
									}
								} else {
									if v68&int32(1) == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
										m.G0 = v11 + int32(80)
										return
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+14)))
										F_assign_simple_var(m, l0, l1, v67, v68&int32(1), base.B2i32(v79|v68 == int32(0)))
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
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
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
						if v52 == int32(1) {
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
							if v55 == int32(3) {
								v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+12)))
								v64 = F_datumTransfer(m, v20, int32(0), v63)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									v67 = v64
									v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
									v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
									if v69 != v67 {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+14)))
										F_assign_simple_var(m, l0, l1, v67, v68&int32(1), base.B2i32(v79|v68 == int32(0)))
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return
										} else {
											m.G0 = v11 + int32(80)
											return
										}
									} else {
										v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
										if v71 != 0 {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
											v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+14)))
											F_assign_simple_var(m, l0, l1, v67, v68&int32(1), base.B2i32(v79|v68 == int32(0)))
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return
											} else {
												m.G0 = v11 + int32(80)
												return
											}
										} else {
											if v68&int32(1) == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
												m.G0 = v11 + int32(80)
												return
											} else {
												v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
												v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+14)))
												F_assign_simple_var(m, l0, l1, v67, v68&int32(1), base.B2i32(v79|v68 == int32(0)))
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
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
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
								v60 = F_expand_array(m, v20, v58, int32(0))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									v67 = v60
									v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
									v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
									if v69 != v67 {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+14)))
										F_assign_simple_var(m, l0, l1, v67, v68&int32(1), base.B2i32(v79|v68 == int32(0)))
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return
										} else {
											m.G0 = v11 + int32(80)
											return
										}
									} else {
										v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
										if v71 != 0 {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
											v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+14)))
											F_assign_simple_var(m, l0, l1, v67, v68&int32(1), base.B2i32(v79|v68 == int32(0)))
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return
											} else {
												m.G0 = v11 + int32(80)
												return
											}
										} else {
											if v68&int32(1) == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
												m.G0 = v11 + int32(80)
												return
											} else {
												v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
												v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+14)))
												F_assign_simple_var(m, l0, l1, v67, v68&int32(1), base.B2i32(v79|v68 == int32(0)))
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
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
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
							v60 = F_expand_array(m, v20, v58, int32(0))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								v67 = v60
								v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
								v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
								if v69 != v67 {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+14)))
									F_assign_simple_var(m, l0, l1, v67, v68&int32(1), base.B2i32(v79|v68 == int32(0)))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return
									} else {
										m.G0 = v11 + int32(80)
										return
									}
								} else {
									v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
									if v71 != 0 {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+14)))
										F_assign_simple_var(m, l0, l1, v67, v68&int32(1), base.B2i32(v79|v68 == int32(0)))
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return
										} else {
											m.G0 = v11 + int32(80)
											return
										}
									} else {
										if v68&int32(1) == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
											m.G0 = v11 + int32(80)
											return
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
											v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+14)))
											F_assign_simple_var(m, l0, l1, v67, v68&int32(1), base.B2i32(v79|v68 == int32(0)))
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
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
			v85 = int32(0)
			F_exec_move_row(m, l0, l1, v85, v85)
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return
			} else {
				m.G0 = v11 + int32(80)
				return
			}
		} else {
			v89 = F_type_is_rowtype(m, l4)
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return
			} else {
				if v89 == int32(0) {
					F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
					mBase = m.M
					v177 = m.ExcPending
					if v177 != 0 {
						return
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v180 = m.ExcPending
						if v180 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_exec_assign_value_5), int32(0))
							mBase = m.M
							v184 = m.ExcPending
							if v184 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_6), int32(_a_F_exec_assign_value_4))
								mBase = m.M
								v189 = m.ExcPending
								if v189 != 0 {
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
					v94 = m.ExcPending
					if v94 != 0 {
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
			v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
			if v95 == int32(1) {
				F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
				mBase = m.M
				v193 = m.ExcPending
				if v193 != 0 {
					return
				} else {
					F_errcode(m, int32(67108994))
					mBase = m.M
					v196 = m.ExcPending
					if v196 != 0 {
						return
					} else {
						v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v197
						F_errmsg(m, int32(_a_F_exec_assign_value_1), v11+int32(32))
						mBase = m.M
						v203 = m.ExcPending
						if v203 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_7), int32(_a_F_exec_assign_value_4))
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
				v98 = int32(0)
				F_exec_move_row(m, l0, l1, v98, v98)
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return
				} else {
					m.G0 = v11 + int32(80)
					return
				}
			}
		} else {
			v102 = F_type_is_rowtype(m, l4)
			mBase = m.M
			v103 = m.ExcPending
			if v103 != 0 {
				return
			} else {
				if v102 == int32(0) {
					F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
					mBase = m.M
					v212 = m.ExcPending
					if v212 != 0 {
						return
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v215 = m.ExcPending
						if v215 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_exec_assign_value_8), int32(0))
							mBase = m.M
							v219 = m.ExcPending
							if v219 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_9), int32(_a_F_exec_assign_value_4))
								mBase = m.M
								v224 = m.ExcPending
								if v224 != 0 {
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
					v107 = m.ExcPending
					if v107 != 0 {
						return
					} else {
						m.G0 = v11 + int32(80)
						return
					}
				}
			}
		}
	case 3:
		v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v113 = *(*int32)(unsafe.Add(mBase, uint32(v108+v109<<(uint(int32(2))%32))))
		v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+36))
		if v114 == int32(0) {
			F_instantiate_empty_record_variable(m, l0, v113)
			mBase = m.M
			v118 = m.ExcPending
			if v118 != 0 {
				return
			} else {
				v119 = *(*int32)(unsafe.Add(mBase, uint32(v113)+36))
				v120 = v119
				v121 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
				v122 = *(*int64)(unsafe.Add(mBase, uint32(v120)+48))
				if v121 != v122 {
					v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v127 = F_expanded_record_lookup_field(m, v120, v124, l1+int32(32))
					mBase = m.M
					v128 = m.ExcPending
					if v128 != 0 {
						return
					} else {
						if v127 == int32(0) {
							F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
							mBase = m.M
							v228 = m.ExcPending
							if v228 != 0 {
								return
							} else {
								F_errcode(m, int32(50360452))
								mBase = m.M
								v231 = m.ExcPending
								if v231 != 0 {
									return
								} else {
									v232 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
									v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v233
									*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v232
									F_errmsg(m, int32(_a_F_exec_assign_value_10), v11-int32(-64))
									mBase = m.M
									v240 = m.ExcPending
									if v240 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_11), int32(_a_F_exec_assign_value_4))
										mBase = m.M
										v245 = m.ExcPending
										if v245 != 0 {
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
							v131 = *(*int64)(unsafe.Add(mBase, uint32(v120)+48))
							*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v131
							v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
							if v133 <= int32(0) {
								F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
								mBase = m.M
								v249 = m.ExcPending
								if v249 != 0 {
									return
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v252 = m.ExcPending
									if v252 != 0 {
										return
									} else {
										v253 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v253
										F_errmsg(m, int32(_a_F_exec_assign_value_12), v11+int32(48))
										mBase = m.M
										v259 = m.ExcPending
										if v259 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_13), int32(_a_F_exec_assign_value_4))
											mBase = m.M
											v264 = m.ExcPending
											if v264 != 0 {
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
								v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
								v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
								v140 = F_exec_cast_value(m, l0, l2, v11+int32(79), l4, l5, v138, v139)
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return
								} else {
									v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
									v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
									v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
									v147 = int32(1)
									F_expanded_record_set_field_internal(m, v120, v142, v140, v143, (v144^int32(-1))&v147, v147)
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
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
					v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
					if v133 <= int32(0) {
						F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
						mBase = m.M
						v249 = m.ExcPending
						if v249 != 0 {
							return
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v252 = m.ExcPending
							if v252 != 0 {
								return
							} else {
								v253 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v253
								F_errmsg(m, int32(_a_F_exec_assign_value_12), v11+int32(48))
								mBase = m.M
								v259 = m.ExcPending
								if v259 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_13), int32(_a_F_exec_assign_value_4))
									mBase = m.M
									v264 = m.ExcPending
									if v264 != 0 {
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
						v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
						v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						v140 = F_exec_cast_value(m, l0, l2, v11+int32(79), l4, l5, v138, v139)
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return
						} else {
							v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
							v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
							v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
							v147 = int32(1)
							F_expanded_record_set_field_internal(m, v120, v142, v140, v143, (v144^int32(-1))&v147, v147)
							mBase = m.M
							v151 = m.ExcPending
							if v151 != 0 {
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
			v120 = v114
			v121 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
			v122 = *(*int64)(unsafe.Add(mBase, uint32(v120)+48))
			if v121 != v122 {
				v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v127 = F_expanded_record_lookup_field(m, v120, v124, l1+int32(32))
				mBase = m.M
				v128 = m.ExcPending
				if v128 != 0 {
					return
				} else {
					if v127 == int32(0) {
						F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
						mBase = m.M
						v228 = m.ExcPending
						if v228 != 0 {
							return
						} else {
							F_errcode(m, int32(50360452))
							mBase = m.M
							v231 = m.ExcPending
							if v231 != 0 {
								return
							} else {
								v232 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
								v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v233
								*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v232
								F_errmsg(m, int32(_a_F_exec_assign_value_10), v11-int32(-64))
								mBase = m.M
								v240 = m.ExcPending
								if v240 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_11), int32(_a_F_exec_assign_value_4))
									mBase = m.M
									v245 = m.ExcPending
									if v245 != 0 {
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
						v131 = *(*int64)(unsafe.Add(mBase, uint32(v120)+48))
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v131
						v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						if v133 <= int32(0) {
							F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
							mBase = m.M
							v249 = m.ExcPending
							if v249 != 0 {
								return
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v252 = m.ExcPending
								if v252 != 0 {
									return
								} else {
									v253 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v253
									F_errmsg(m, int32(_a_F_exec_assign_value_12), v11+int32(48))
									mBase = m.M
									v259 = m.ExcPending
									if v259 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_13), int32(_a_F_exec_assign_value_4))
										mBase = m.M
										v264 = m.ExcPending
										if v264 != 0 {
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
							v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
							v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
							v140 = F_exec_cast_value(m, l0, l2, v11+int32(79), l4, l5, v138, v139)
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
								v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
								v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
								v147 = int32(1)
								F_expanded_record_set_field_internal(m, v120, v142, v140, v143, (v144^int32(-1))&v147, v147)
								mBase = m.M
								v151 = m.ExcPending
								if v151 != 0 {
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
				v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
				if v133 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
					mBase = m.M
					v249 = m.ExcPending
					if v249 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v252 = m.ExcPending
						if v252 != 0 {
							return
						} else {
							v253 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v253
							F_errmsg(m, int32(_a_F_exec_assign_value_12), v11+int32(48))
							mBase = m.M
							v259 = m.ExcPending
							if v259 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_13), int32(_a_F_exec_assign_value_4))
								mBase = m.M
								v264 = m.ExcPending
								if v264 != 0 {
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
					v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
					v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					v140 = F_exec_cast_value(m, l0, l2, v11+int32(79), l4, l5, v138, v139)
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return
					} else {
						v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
						v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
						v147 = int32(1)
						F_expanded_record_set_field_internal(m, v120, v142, v140, v143, (v144^int32(-1))&v147, v147)
						mBase = m.M
						v151 = m.ExcPending
						if v151 != 0 {
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
		F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
		mBase = m.M
		v155 = m.ExcPending
		if v155 != 0 {
			return
		} else {
			v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v156
			F_errmsg_internal(m, int32(_a_F_exec_assign_value_14), v11)
			mBase = m.M
			v160 = m.ExcPending
			if v160 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_15), int32(_a_F_exec_assign_value_4))
				mBase = m.M
				v165 = m.ExcPending
				if v165 != 0 {
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
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
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	if (base.B2i32(l4 == l6)|base.B2i32(l6 == int32(-1)))&base.B2i32(l3 == l5) != 0 {
		v325 = l1
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v16 + int32(32)
	return v325
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l3
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v30 = v16 + int32(16)
	v33 = v16 + int32(15)
	v34 = F_hash_search(m, v28, v30, int32(1), v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
	if v38 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	if v60 != 0 {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[0]))
	v44 = F_hash_search(m, v42, v30, int32(1), v33)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v59 = v58
	goto L5
L9:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
	if v46 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = int32(0)
	goto L12
L11:
	;
	goto L12
L12:
	;
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v51
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+24)) = uint8(v51)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v44
	v59 = v44
	goto L5
L13:
	;
	if v274 == int32(0) {
		v325 = l1
		goto L1
	} else {
		goto L67
	}
L14:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+8)))
	if v61 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v76 = int32(_a_F_exec_cast_value_0)
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2])) = v80
	v83 = F_palloc0(m, int32(16))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L3
	} else {
		goto L21
	}
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v274 = v64
	goto L13
L18:
	;
	goto L19
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)+24))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v60)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v60)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	F_MemoryContextDelete(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = int32(0)
	goto L16
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(34)
	v89 = F_get_typcollation(m, l3)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = v89
	if base.B2i32(l3 == int32(705))|base.B2i32(l3 == int32(2249)) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v126 = m.G0
	v128 = v126 - int32(16)
	m.G0 = v128
	v134 = m.G0
	v136 = v134 - int32(480)
	m.G0 = v136
	v140 = int32(0)
	base.MemoryFill(m, v136+int32(392), v140, int32(88))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+448)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v136)+388)) = int32(266)
	base.MemoryFill(m, v136, v140, int32(384))
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = int32(267)
	*(*int32)(unsafe.Add(mBase, uint32(v136)+8)) = v136 + int32(388)
	v155 = F_eval_const_expressions(m, v136, v125)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L3
	} else {
		goto L32
	}
L24:
	;
	v100 = int32(2)
	v103 = F_coerce_to_target_type(m, int32(0), v83, l3, l5, l6, v100, v100, int32(-1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L3
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v107 = F_palloc0(m, int32(24))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L29
	}
L27:
	;
	if v103 != 0 {
		v125 = v103
		goto L23
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v109 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v107)+20)) = v109
	*(*int64)(unsafe.Add(mBase, uint32(v107)+12)) = int64(8589934592)
	*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = int32(28)
	if l6 == v109 {
		v125 = v107
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v123 = F_coerce_to_target_type(m, int32(0), v107, l5, l5, l6, int32(1), int32(2), int32(-1))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	v125 = v123
	goto L23
L32:
	;
	F_fix_opfuncids(m, v155)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	v159 = F_extract_query_dependencies_walker(m, v155, v136)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v136)+444))
	*(*int32)(unsafe.Add(mBase, uint32(v128+int32(12)))) = v161
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v136)+448))
	*(*int32)(unsafe.Add(mBase, uint32(v128+int32(8)))) = v163
	m.G0 = v136 + int32(480)
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2]))
	v174 = F_AllocSetContextCreateInternal(m, v169, int32(_a_F_exec_cast_value_1), int32(0), int32(1024), int32(_a_F_exec_cast_value_2))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v176 = int32(_a_F_exec_cast_value_0)
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2])) = v174
	v181 = F_palloc(m, int32(32))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = int32(838275847)
	v185 = F_copyObjectImpl(m, v155)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v187 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v181)+8)) = uint8(v187)
	*(*int32)(unsafe.Add(mBase, uint32(v181)+4)) = v185
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v191 = F_copyObjectImpl(m, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181)+12)) = v191
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	v195 = F_copyObjectImpl(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181)+20)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v181)+16)) = v195
	*(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2])) = v177
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[3]))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v174)+16))
	if v206 != v202 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[4]))
	if v236 != 0 {
		goto L58
	} else {
		goto L59
	}
L41:
	;
	if v206 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	goto L40
L44:
	;
	if v202 != 0 {
		goto L51
	} else {
		goto L52
	}
L45:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v174)+28))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v174)+24))
	if v211 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v210 == int32(0) {
		goto L44
	} else {
		goto L50
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+28)) = v210
	goto L46
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206)+20)) = v210
	goto L46
L50:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v174)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v210)+24)) = v216
	goto L44
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v174)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v174)+16)) = v202
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v202)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v174)+28)) = v223
	if v223 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v174)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v174)+16)) = int32(0)
	goto L43
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+24)) = v174
	goto L56
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+20)) = v174
	goto L40
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181)+24)) = v243
	v245 = int32(_a_F_exec_cast_value_3)
	*(*int32)(unsafe.Add(mBase, uint32(v181)+28)) = v245
	v248 = v181 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v243)+4)) = v248
	*(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[5])) = v248
	m.G0 = v128 + int32(16)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	if v256 == int32(27) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[5]))
	v243 = v238
	goto L57
L59:
	;
	goto L60
L60:
	;
	v240 = int32(_a_F_exec_cast_value_3)
	*(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[4])) = v240
	v243 = v240
	goto L57
L61:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	if v260 != v83 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v263 = v255
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+16)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = v181
	v266 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+24)) = uint8(v266)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v266
	*(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2])) = v77
	v274 = v263
	goto L13
L64:
	;
	v262 = v255
	goto L66
L65:
	;
	v262 = int32(0)
	goto L66
L66:
	;
	v263 = v262
	goto L63
L67:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[1]))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+56))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v34)+28))
	if v285 != v286 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2])) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v309)+40)) = l1
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	*(*uint8)(unsafe.Add(mBase, uint32(v309)+44)) = uint8(v313)
	v315 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+24)) = uint8(v315)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v306)+20))
	v318 = m.T0[v317].(func(*base.Module, int32, int32, int32) int32)(m, v306, v309, l2)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L3
	} else {
		goto L75
	}
L69:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2])) = v297
	v300 = F_ExecInitExpr(m, v274, int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L3
	} else {
		goto L74
	}
L70:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2]))
	v294 = v289
	goto L69
L71:
	;
	goto L72
L72:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2]))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+24)))
	if v292 != 0 {
		v294 = v291
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	v306 = v293
	v307 = v291
	goto L68
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v285
	v303 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+24)) = uint8(v303)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v300
	v306 = v300
	v307 = v294
	goto L68
L75:
	;
	v320 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+24)) = uint8(v320)
	*(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2])) = v307
	v325 = v318
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
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
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
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
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
	if v263&int32(5) == int32(0) {
		goto L103
	} else {
		goto L104
	}
L3:
	;
	v144 = int32(_a_F_exec_move_row_from_datum_0)
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_exec_move_row_from_datum[0]))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_move_row_from_datum[0])) = v148
	v150 = F_pg_detoast_datum(m, l2)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
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
	v262 = int32(0)
	v263 = v24
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
		v262 = v76
		v263 = v78
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
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v131 != 0 {
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
	F_DeleteExpandedObject(m, v131+int32(12))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
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
	v141 = m.ExcPending
	if v141 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	F_assign_record_var(m, l0, l1, v76)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	goto L1
L63:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_move_row_from_datum[0])) = v145
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v150
	v156 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v156
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+20)) = uint16(v156)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(-1)
	v162 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(base.Ui32(v154) >> (uint(v162) % 32))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v167 != v162 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v253 = F_lookup_rowtype_tupdesc(m, v166, v165)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L10
	} else {
		goto L99
	}
L65:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v170 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if base.B2i32(v191 != int32(2249))&base.B2i32(v166 != v191) != 0 {
		goto L64
	} else {
		goto L75
	}
L67:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v170)+36))
	if v166 != v173 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	if v166 == int32(2249) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	if v165 < int32(0) {
		goto L66
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v183 = int32(1)
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	F_expanded_record_set_tuple(m, v170, v10+int32(12), v183, (v184^int32(-1))&v183)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L10
	} else {
		goto L74
	}
L72:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v170)+40))
	if v165 != v179 {
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
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+20))
	v198 = F_make_expanded_record_from_typeid(m, v166, v165, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L10
	} else {
		goto L76
	}
L76:
	;
	v202 = int32(1)
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	F_expanded_record_set_tuple(m, v198, v10+int32(12), v202, (v203^int32(-1))&v202)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L10
	} else {
		goto L77
	}
L77:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v198)+8))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v210)+16))
	if v215 != v211 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v244 != 0 {
		goto L95
	} else {
		goto L96
	}
L79:
	;
	if v215 == int32(0) {
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
	if v211 != 0 {
		goto L89
	} else {
		goto L90
	}
L83:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v210)+28))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v210)+24))
	if v220 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	if v219 == int32(0) {
		goto L82
	} else {
		goto L88
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+28)) = v219
	goto L84
L86:
	;
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215)+20)) = v219
	goto L84
L88:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v210)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v219)+24)) = v225
	goto L82
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v210)+16)) = v211
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v211)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v210)+28)) = v232
	if v232 != 0 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v210)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v210)+16)) = int32(0)
	goto L81
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+24)) = v210
	goto L94
L93:
	;
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+20)) = v210
	goto L78
L95:
	;
	F_DeleteExpandedObject(m, v244+int32(12))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L10
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v198
	goto L1
L98:
	;
	goto L97
L99:
	;
	F_exec_move_row(m, l0, l1, v10+int32(12), v253)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L10
	} else {
		goto L100
	}
L100:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v253)+12))
	if v257 < int32(0) {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_DecrTupleDescRefCount(m, v253)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L10
	} else {
		goto L102
	}
L102:
	;
	goto L1
L103:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	if v270 != 0 {
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
	v277 = m.ExcPending
	if v277 != 0 {
		goto L10
	} else {
		goto L111
	}
L106:
	;
	v273 = v270
	goto L108
L107:
	;
	v271 = F_expanded_record_fetch_tupdesc(m, v20)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L10
	} else {
		goto L109
	}
L108:
	;
	F_exec_move_row(m, l0, l1, int32(0), v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L10
	} else {
		goto L110
	}
L109:
	;
	v273 = v271
	goto L108
L110:
	;
	goto L1
L111:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	if v280 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v283 = v280
	goto L114
L113:
	;
	v281 = F_expanded_record_fetch_tupdesc(m, v20)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L10
	} else {
		goto L115
	}
L114:
	;
	F_exec_move_row_from_fields(m, l0, l1, v262, v278, v279, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L10
	} else {
		goto L116
	}
L115:
	;
	v283 = v281
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
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
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
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
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
	var v144 int32
	_ = v144
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
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
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
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_exec_prepare_plan_0)
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
		*(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[0])) = int32(-6)
		v84 = int32(0)
		m.G0 = v22 + int32(48)
		if v84 != 0 {
			F_SPI_keepplan(m, v84)
			mBase = m.M
			v92 = m.ExcPending
			if v92 != 0 {
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l1)+48)) = int64(0)
				v95 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v95
				*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v84
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
				if v101 == v95 {
					v149 = v95
				} else {
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
					if v104 != int32(1) {
						v149 = v95
					} else {
						v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
						v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
						v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+60))
						if v109 == int32(0) {
							v149 = v95
						} else {
							v112 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
							if v112 != int32(1) {
								v149 = v95
							} else {
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
								if v117 != int32(67) {
									v149 = v95
								} else {
									v120 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
									if v120 != int32(1) {
										v149 = v95
									} else {
										v123 = *(*int32)(unsafe.Add(mBase, uint32(v116)+52))
										if v123 != 0 {
											v149 = v95
										} else {
											v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+36)))
											if v124 != 0 {
												v149 = v95
											} else {
												v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+37)))
												if v125 != 0 {
													v149 = v95
												} else {
													v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+38)))
													if v126 != 0 {
														v149 = v95
													} else {
														v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+39)))
														if v127 != 0 {
															v149 = v95
														} else {
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v116)+48))
															if v128 != 0 {
																v149 = v95
															} else {
																v129 = *(*int32)(unsafe.Add(mBase, uint32(v116)+60))
																v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
																if v130 != 0 {
																	v149 = v95
																} else {
																	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
																	if v131 != 0 {
																		v149 = v95
																	} else {
																		v132 = *(*int32)(unsafe.Add(mBase, uint32(v116)+100))
																		if v132 != 0 {
																			v149 = v95
																		} else {
																			v133 = *(*int32)(unsafe.Add(mBase, uint32(v116)+108))
																			if v133 != 0 {
																				v149 = v95
																			} else {
																				v134 = *(*int32)(unsafe.Add(mBase, uint32(v116)+112))
																				if v134 != 0 {
																					v149 = v95
																				} else {
																					v135 = *(*int32)(unsafe.Add(mBase, uint32(v116)+116))
																					if v135 != 0 {
																						v149 = v95
																					} else {
																						v136 = *(*int32)(unsafe.Add(mBase, uint32(v116)+120))
																						if v136 != 0 {
																							v149 = v95
																						} else {
																							v137 = *(*int32)(unsafe.Add(mBase, uint32(v116)+124))
																							if v137 != 0 {
																								v149 = v95
																							} else {
																								v138 = *(*int32)(unsafe.Add(mBase, uint32(v116)+128))
																								if v138 != 0 {
																									v149 = v95
																								} else {
																									v139 = *(*int32)(unsafe.Add(mBase, uint32(v116)+132))
																									if v139 != 0 {
																										v149 = v95
																									} else {
																										v140 = *(*int32)(unsafe.Add(mBase, uint32(v116)+144))
																										if v140 != 0 {
																											v149 = v95
																										} else {
																											v141 = *(*int32)(unsafe.Add(mBase, uint32(v116)+76))
																											if v141 == int32(0) {
																												v149 = v95
																											} else {
																												v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
																												v149 = base.B2i32(v144 == int32(1))
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
				if v149 != 0 {
					v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
					v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
					v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
					v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
					v154 = int32(_a_F_exec_prepare_plan_1)
					v155 = *(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[1]))
					v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
					v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+20))
					*(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[1])) = v158
					v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
					v161 = F_SPI_plan_get_cached_plan(m, v160)
					mBase = m.M
					v162 = m.ExcPending
					if v162 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[1])) = v155
						v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						v166 = F_CachedPlanAllowsSimpleValidityCheck(m, v153, v161, v165)
						mBase = m.M
						v167 = m.ExcPending
						if v167 != 0 {
							return
						} else {
							if v166 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+60)) = v161
								*(*int32)(unsafe.Add(mBase, uint32(l1)+56)) = v153
								v171 = *(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[2]))
								v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+56))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v172
								F_exec_save_simple_expr(m, l1, v161)
								mBase = m.M
								v175 = m.ExcPending
								if v175 != 0 {
									return
								} else {
									v177 = *(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[3]))
									F_ReleaseCachedPlan(m, v161, v177)
									mBase = m.M
									v179 = m.ExcPending
									if v179 != 0 {
										return
									} else {
										m.G0 = v11 + int32(32)
										return
									}
								}
							} else {
								v177 = *(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[3]))
								F_ReleaseCachedPlan(m, v161, v177)
								mBase = m.M
								v179 = m.ExcPending
								if v179 != 0 {
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
			F_errstart_cold(m, int32(21), int32(_a_F_exec_prepare_plan_2))
			mBase = m.M
			v189 = m.ExcPending
			if v189 != 0 {
				return
			} else {
				v190 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v192 = *(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[0]))
				v193 = F_SPI_result_code_string(m, v192)
				mBase = m.M
				v194 = m.ExcPending
				if v194 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v193
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v190
					F_errmsg_internal(m, int32(_a_F_exec_prepare_plan_3), v11)
					mBase = m.M
					v199 = m.ExcPending
					if v199 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_exec_prepare_plan_4), int32(_a_F_exec_prepare_plan_5), int32(_a_F_exec_prepare_plan_6))
						mBase = m.M
						v204 = m.ExcPending
						if v204 != 0 {
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
		v35 = *(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[4]))
		if v35 != 0 {
			v37 = *(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[5]))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
			v40 = *(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[4]))
			*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v38
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[1])) = v43
			v46 = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[0])) = v46
			v48 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v22)+12)) = v48
			*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v46
			*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(569278163)
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v54
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
			*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v48
			*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v56
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v60
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v62
			v65 = v22 + int32(8)
			F__SPI_prepare_plan(m, v19, v65)
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return
			} else {
				v68 = F__SPI_make_plan_non_temp(m, v65)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return
				} else {
					v72 = *(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[4]))
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
					*(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[1])) = v73
					*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = int32(0)
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+24))
					F_MemoryContextReset(m, v77)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return
					} else {
						v84 = v68
						m.G0 = v22 + int32(48)
						if v84 != 0 {
							F_SPI_keepplan(m, v84)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l1)+48)) = int64(0)
								v95 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v95
								*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v84
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
								if v101 == v95 {
									v149 = v95
								} else {
									v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
									if v104 != int32(1) {
										v149 = v95
									} else {
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+60))
										if v109 == int32(0) {
											v149 = v95
										} else {
											v112 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
											if v112 != int32(1) {
												v149 = v95
											} else {
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
												if v117 != int32(67) {
													v149 = v95
												} else {
													v120 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
													if v120 != int32(1) {
														v149 = v95
													} else {
														v123 = *(*int32)(unsafe.Add(mBase, uint32(v116)+52))
														if v123 != 0 {
															v149 = v95
														} else {
															v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+36)))
															if v124 != 0 {
																v149 = v95
															} else {
																v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+37)))
																if v125 != 0 {
																	v149 = v95
																} else {
																	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+38)))
																	if v126 != 0 {
																		v149 = v95
																	} else {
																		v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+39)))
																		if v127 != 0 {
																			v149 = v95
																		} else {
																			v128 = *(*int32)(unsafe.Add(mBase, uint32(v116)+48))
																			if v128 != 0 {
																				v149 = v95
																			} else {
																				v129 = *(*int32)(unsafe.Add(mBase, uint32(v116)+60))
																				v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
																				if v130 != 0 {
																					v149 = v95
																				} else {
																					v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
																					if v131 != 0 {
																						v149 = v95
																					} else {
																						v132 = *(*int32)(unsafe.Add(mBase, uint32(v116)+100))
																						if v132 != 0 {
																							v149 = v95
																						} else {
																							v133 = *(*int32)(unsafe.Add(mBase, uint32(v116)+108))
																							if v133 != 0 {
																								v149 = v95
																							} else {
																								v134 = *(*int32)(unsafe.Add(mBase, uint32(v116)+112))
																								if v134 != 0 {
																									v149 = v95
																								} else {
																									v135 = *(*int32)(unsafe.Add(mBase, uint32(v116)+116))
																									if v135 != 0 {
																										v149 = v95
																									} else {
																										v136 = *(*int32)(unsafe.Add(mBase, uint32(v116)+120))
																										if v136 != 0 {
																											v149 = v95
																										} else {
																											v137 = *(*int32)(unsafe.Add(mBase, uint32(v116)+124))
																											if v137 != 0 {
																												v149 = v95
																											} else {
																												v138 = *(*int32)(unsafe.Add(mBase, uint32(v116)+128))
																												if v138 != 0 {
																													v149 = v95
																												} else {
																													v139 = *(*int32)(unsafe.Add(mBase, uint32(v116)+132))
																													if v139 != 0 {
																														v149 = v95
																													} else {
																														v140 = *(*int32)(unsafe.Add(mBase, uint32(v116)+144))
																														if v140 != 0 {
																															v149 = v95
																														} else {
																															v141 = *(*int32)(unsafe.Add(mBase, uint32(v116)+76))
																															if v141 == int32(0) {
																																v149 = v95
																															} else {
																																v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
																																v149 = base.B2i32(v144 == int32(1))
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
								if v149 != 0 {
									v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
									v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
									v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
									v154 = int32(_a_F_exec_prepare_plan_1)
									v155 = *(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[1]))
									v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
									v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+20))
									*(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[1])) = v158
									v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									v161 = F_SPI_plan_get_cached_plan(m, v160)
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[1])) = v155
										v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
										v166 = F_CachedPlanAllowsSimpleValidityCheck(m, v153, v161, v165)
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return
										} else {
											if v166 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(l1)+60)) = v161
												*(*int32)(unsafe.Add(mBase, uint32(l1)+56)) = v153
												v171 = *(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[2]))
												v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+56))
												*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v172
												F_exec_save_simple_expr(m, l1, v161)
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return
												} else {
													v177 = *(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[3]))
													F_ReleaseCachedPlan(m, v161, v177)
													mBase = m.M
													v179 = m.ExcPending
													if v179 != 0 {
														return
													} else {
														m.G0 = v11 + int32(32)
														return
													}
												}
											} else {
												v177 = *(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[3]))
												F_ReleaseCachedPlan(m, v161, v177)
												mBase = m.M
												v179 = m.ExcPending
												if v179 != 0 {
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
							F_errstart_cold(m, int32(21), int32(_a_F_exec_prepare_plan_2))
							mBase = m.M
							v189 = m.ExcPending
							if v189 != 0 {
								return
							} else {
								v190 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v192 = *(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[0]))
								v193 = F_SPI_result_code_string(m, v192)
								mBase = m.M
								v194 = m.ExcPending
								if v194 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v193
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v190
									F_errmsg_internal(m, int32(_a_F_exec_prepare_plan_3), v11)
									mBase = m.M
									v199 = m.ExcPending
									if v199 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_exec_prepare_plan_4), int32(_a_F_exec_prepare_plan_5), int32(_a_F_exec_prepare_plan_6))
										mBase = m.M
										v204 = m.ExcPending
										if v204 != 0 {
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
			*(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[0])) = int32(-4)
			v84 = int32(0)
			m.G0 = v22 + int32(48)
			if v84 != 0 {
				F_SPI_keepplan(m, v84)
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l1)+48)) = int64(0)
					v95 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v95
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v84
					v101 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
					if v101 == v95 {
						v149 = v95
					} else {
						v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
						if v104 != int32(1) {
							v149 = v95
						} else {
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
							v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
							v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+60))
							if v109 == int32(0) {
								v149 = v95
							} else {
								v112 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
								if v112 != int32(1) {
									v149 = v95
								} else {
									v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
									v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
									if v117 != int32(67) {
										v149 = v95
									} else {
										v120 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
										if v120 != int32(1) {
											v149 = v95
										} else {
											v123 = *(*int32)(unsafe.Add(mBase, uint32(v116)+52))
											if v123 != 0 {
												v149 = v95
											} else {
												v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+36)))
												if v124 != 0 {
													v149 = v95
												} else {
													v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+37)))
													if v125 != 0 {
														v149 = v95
													} else {
														v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+38)))
														if v126 != 0 {
															v149 = v95
														} else {
															v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+39)))
															if v127 != 0 {
																v149 = v95
															} else {
																v128 = *(*int32)(unsafe.Add(mBase, uint32(v116)+48))
																if v128 != 0 {
																	v149 = v95
																} else {
																	v129 = *(*int32)(unsafe.Add(mBase, uint32(v116)+60))
																	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
																	if v130 != 0 {
																		v149 = v95
																	} else {
																		v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
																		if v131 != 0 {
																			v149 = v95
																		} else {
																			v132 = *(*int32)(unsafe.Add(mBase, uint32(v116)+100))
																			if v132 != 0 {
																				v149 = v95
																			} else {
																				v133 = *(*int32)(unsafe.Add(mBase, uint32(v116)+108))
																				if v133 != 0 {
																					v149 = v95
																				} else {
																					v134 = *(*int32)(unsafe.Add(mBase, uint32(v116)+112))
																					if v134 != 0 {
																						v149 = v95
																					} else {
																						v135 = *(*int32)(unsafe.Add(mBase, uint32(v116)+116))
																						if v135 != 0 {
																							v149 = v95
																						} else {
																							v136 = *(*int32)(unsafe.Add(mBase, uint32(v116)+120))
																							if v136 != 0 {
																								v149 = v95
																							} else {
																								v137 = *(*int32)(unsafe.Add(mBase, uint32(v116)+124))
																								if v137 != 0 {
																									v149 = v95
																								} else {
																									v138 = *(*int32)(unsafe.Add(mBase, uint32(v116)+128))
																									if v138 != 0 {
																										v149 = v95
																									} else {
																										v139 = *(*int32)(unsafe.Add(mBase, uint32(v116)+132))
																										if v139 != 0 {
																											v149 = v95
																										} else {
																											v140 = *(*int32)(unsafe.Add(mBase, uint32(v116)+144))
																											if v140 != 0 {
																												v149 = v95
																											} else {
																												v141 = *(*int32)(unsafe.Add(mBase, uint32(v116)+76))
																												if v141 == int32(0) {
																													v149 = v95
																												} else {
																													v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
																													v149 = base.B2i32(v144 == int32(1))
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
					if v149 != 0 {
						v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
						v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
						v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
						v154 = int32(_a_F_exec_prepare_plan_1)
						v155 = *(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[1]))
						v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+20))
						*(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[1])) = v158
						v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						v161 = F_SPI_plan_get_cached_plan(m, v160)
						mBase = m.M
						v162 = m.ExcPending
						if v162 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[1])) = v155
							v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							v166 = F_CachedPlanAllowsSimpleValidityCheck(m, v153, v161, v165)
							mBase = m.M
							v167 = m.ExcPending
							if v167 != 0 {
								return
							} else {
								if v166 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+60)) = v161
									*(*int32)(unsafe.Add(mBase, uint32(l1)+56)) = v153
									v171 = *(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[2]))
									v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+56))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v172
									F_exec_save_simple_expr(m, l1, v161)
									mBase = m.M
									v175 = m.ExcPending
									if v175 != 0 {
										return
									} else {
										v177 = *(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[3]))
										F_ReleaseCachedPlan(m, v161, v177)
										mBase = m.M
										v179 = m.ExcPending
										if v179 != 0 {
											return
										} else {
											m.G0 = v11 + int32(32)
											return
										}
									}
								} else {
									v177 = *(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[3]))
									F_ReleaseCachedPlan(m, v161, v177)
									mBase = m.M
									v179 = m.ExcPending
									if v179 != 0 {
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
				F_errstart_cold(m, int32(21), int32(_a_F_exec_prepare_plan_2))
				mBase = m.M
				v189 = m.ExcPending
				if v189 != 0 {
					return
				} else {
					v190 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v192 = *(*int32)(unsafe.Add(mBase, _c_F_exec_prepare_plan[0]))
					v193 = F_SPI_result_code_string(m, v192)
					mBase = m.M
					v194 = m.ExcPending
					if v194 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v193
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v190
						F_errmsg_internal(m, int32(_a_F_exec_prepare_plan_3), v11)
						mBase = m.M
						v199 = m.ExcPending
						if v199 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_exec_prepare_plan_4), int32(_a_F_exec_prepare_plan_5), int32(_a_F_exec_prepare_plan_6))
							mBase = m.M
							v204 = m.ExcPending
							if v204 != 0 {
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
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
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
		v36 = F_ArrayGetNItemsSafe(m, v31, l1+int32(16))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v32 + v36<<(uint(int32(2))%32)
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v50 = F_execute(m, l0+v42<<(uint(int32(3))%32), v9+int32(8), int32(0), l2, int32(_a_F_execconsistent_0))
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
						F_errmsg(m, int32(_a_F_execconsistent_1), int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_execconsistent_2), int32(311), int32(_a_F_execconsistent_3))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
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
				v36 = F_ArrayGetNItemsSafe(m, v31, l1+int32(16))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v32 + v36<<(uint(int32(2))%32)
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v50 = F_execute(m, l0+v42<<(uint(int32(3))%32), v9+int32(8), int32(0), l2, int32(_a_F_execconsistent_0))
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
									F_errmsg(m, int32(_a_F_executeLikeRegex_0), int32(0))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										F_errsave_finish(m, int32(0), int32(_a_F_executeLikeRegex_1), int32(680), int32(_a_F_executeLikeRegex_2))
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
	var v100 float64
	_ = v100
	var v103 float64
	_ = v103
	var v104 int64
	_ = v104
	var v107 int64
	_ = v107
	var v116 float64
	_ = v116
	var v123 float64
	_ = v123
	var v124 float64
	_ = v124
	var v125 float64
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 float64
	_ = v134
	var v137 int32
	_ = v137
	var v141 float64
	_ = v141
	var v142 float64
	_ = v142
	var v143 float64
	_ = v143
	var v152 float64
	_ = v152
	var v155 float64
	_ = v155
	var v157 float64
	_ = v157
	var v164 float64
	_ = v164
	var v166 float64
	_ = v166
	var v176 float64
	_ = v176
	v14 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0))>>(uint(int64(52))%64))) & int32(2047)
	v19 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(5.551115123125783e-17))) >> (uint(int64(52)) % 64)))
	if base.Ui32(v14-v19) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(512)))>>(uint(int64(52))%64)))-v19) {
		v62 = v14
		v64 = *(*float64)(unsafe.Add(mBase, _c_F_exp[0]))
		v67 = *(*float64)(unsafe.Add(mBase, _c_F_exp[1]))
		v68 = base.F64_add(base.F64_mul(l0, v64), v67)
		v69 = base.F64_sub(v68, v67)
		v71 = *(*float64)(unsafe.Add(mBase, _c_F_exp[2]))
		v74 = *(*float64)(unsafe.Add(mBase, _c_F_exp[3]))
		v77 = base.F64_add(base.F64_mul(v69, v71), base.F64_add(base.F64_mul(v69, v74), l0))
		v78 = base.F64_mul(v77, v77)
		v81 = *(*float64)(unsafe.Add(mBase, _c_F_exp[4]))
		v84 = *(*float64)(unsafe.Add(mBase, _c_F_exp[5]))
		v88 = *(*float64)(unsafe.Add(mBase, _c_F_exp[6]))
		v91 = *(*float64)(unsafe.Add(mBase, _c_F_exp[7]))
		v94 = base.I64_reinterpret_f64(v68)
		v99 = base.I32_wrap_i64(v94) << (uint(int32(4)) % 32) & int32(2032)
		v100 = *(*float64)(unsafe.Add(mBase, uint32(v99)+uint32(_c_F_exp[8])))
		v103 = base.F64_add(base.F64_mul(base.F64_mul(v78, v78), base.F64_add(base.F64_mul(v77, v81), v84)), base.F64_add(base.F64_mul(v78, base.F64_add(base.F64_mul(v77, v88), v91)), base.F64_add(v100, v77)))
		v104 = *(*int64)(unsafe.Add(mBase, uint32(v99)+uint32(_c_F_exp[9])))
		v107 = v104 + v94<<(uint(int64(45))%64)
		if v62 == int32(0) {
			if v94&int64(2147483648) == int64(0) {
				v116 = base.F64_reinterpret_i64(v107 - int64(4544132024016830464))
				v164 = base.F64_mul(base.F64_add(base.F64_mul(v116, v103), v116), float64(5.486124068793689e+303))
			} else {
				v123 = base.F64_reinterpret_i64(v107 + int64(4602678819172646912))
				v124 = base.F64_mul(v123, v103)
				v125 = base.F64_add(v124, v123)
				if base.F64_lt(v125, float64(1)) != 0 {
					v129 = m.G0
					v131 = v129 - int32(16)
					*(*int64)(unsafe.Add(mBase, uint32(v131)+8)) = int64(4503599627370496)
					v134 = *(*float64)(unsafe.Add(mBase, uint32(v131)+8))
					v137 = m.G0
					*(*float64)(unsafe.Add(mBase, uint32(v137-int32(16))+8)) = base.F64_mul(v134, float64(2.2250738585072014e-308))
					v141 = float64(0)
					v142 = float64(1)
					v143 = base.F64_add(v125, v142)
					v152 = base.F64_add(base.F64_add(v143, base.F64_add(base.F64_add(v124, base.F64_sub(v123, v125)), base.F64_add(v125, base.F64_sub(v142, v143)))), float64(-1))
					if base.F64_eq(v152, v141) != 0 {
						v155 = v141
					} else {
						v155 = v152
					}
					v157 = v155
				} else {
					v157 = v125
				}
				v164 = base.F64_mul(v157, float64(2.2250738585072014e-308))
			}
			return v164
		} else {
			v166 = base.F64_reinterpret_i64(v107)
			v176 = base.F64_add(base.F64_mul(v166, v103), v166)
			return v176
		}
	} else {
		if base.Ui32(v14) < base.Ui32(v19) {
			return base.F64_add(l0, float64(1))
		} else {
			if base.Ui32(v14) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(1024)))>>(uint(int64(52))%64)))) {
				v62 = int32(0)
				v64 = *(*float64)(unsafe.Add(mBase, _c_F_exp[0]))
				v67 = *(*float64)(unsafe.Add(mBase, _c_F_exp[1]))
				v68 = base.F64_add(base.F64_mul(l0, v64), v67)
				v69 = base.F64_sub(v68, v67)
				v71 = *(*float64)(unsafe.Add(mBase, _c_F_exp[2]))
				v74 = *(*float64)(unsafe.Add(mBase, _c_F_exp[3]))
				v77 = base.F64_add(base.F64_mul(v69, v71), base.F64_add(base.F64_mul(v69, v74), l0))
				v78 = base.F64_mul(v77, v77)
				v81 = *(*float64)(unsafe.Add(mBase, _c_F_exp[4]))
				v84 = *(*float64)(unsafe.Add(mBase, _c_F_exp[5]))
				v88 = *(*float64)(unsafe.Add(mBase, _c_F_exp[6]))
				v91 = *(*float64)(unsafe.Add(mBase, _c_F_exp[7]))
				v94 = base.I64_reinterpret_f64(v68)
				v99 = base.I32_wrap_i64(v94) << (uint(int32(4)) % 32) & int32(2032)
				v100 = *(*float64)(unsafe.Add(mBase, uint32(v99)+uint32(_c_F_exp[8])))
				v103 = base.F64_add(base.F64_mul(base.F64_mul(v78, v78), base.F64_add(base.F64_mul(v77, v81), v84)), base.F64_add(base.F64_mul(v78, base.F64_add(base.F64_mul(v77, v88), v91)), base.F64_add(v100, v77)))
				v104 = *(*int64)(unsafe.Add(mBase, uint32(v99)+uint32(_c_F_exp[9])))
				v107 = v104 + v94<<(uint(int64(45))%64)
				if v62 == int32(0) {
					if v94&int64(2147483648) == int64(0) {
						v116 = base.F64_reinterpret_i64(v107 - int64(4544132024016830464))
						v164 = base.F64_mul(base.F64_add(base.F64_mul(v116, v103), v116), float64(5.486124068793689e+303))
					} else {
						v123 = base.F64_reinterpret_i64(v107 + int64(4602678819172646912))
						v124 = base.F64_mul(v123, v103)
						v125 = base.F64_add(v124, v123)
						if base.F64_lt(v125, float64(1)) != 0 {
							v129 = m.G0
							v131 = v129 - int32(16)
							*(*int64)(unsafe.Add(mBase, uint32(v131)+8)) = int64(4503599627370496)
							v134 = *(*float64)(unsafe.Add(mBase, uint32(v131)+8))
							v137 = m.G0
							*(*float64)(unsafe.Add(mBase, uint32(v137-int32(16))+8)) = base.F64_mul(v134, float64(2.2250738585072014e-308))
							v141 = float64(0)
							v142 = float64(1)
							v143 = base.F64_add(v125, v142)
							v152 = base.F64_add(base.F64_add(v143, base.F64_add(base.F64_add(v124, base.F64_sub(v123, v125)), base.F64_add(v125, base.F64_sub(v142, v143)))), float64(-1))
							if base.F64_eq(v152, v141) != 0 {
								v155 = v141
							} else {
								v155 = v152
							}
							v157 = v155
						} else {
							v157 = v125
						}
						v164 = base.F64_mul(v157, float64(2.2250738585072014e-308))
					}
					return v164
				} else {
					v166 = base.F64_reinterpret_i64(v107)
					v176 = base.F64_add(base.F64_mul(v166, v103), v166)
					return v176
				}
			} else {
				v40 = base.I64_reinterpret_f64(l0)
				if v40 == int64(-4503599627370496) {
					v176 = float64(0)
					return v176
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v252 int32
	_ = v252
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	v9 = int32(0)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l5)+56))
	v29 = F_palloc0(m, int32(136))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(101)
	base.MemoryCopy(m, v29, l1, int32(136))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v27
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+119)))
	v38 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+128)) = v38
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+21)) = uint8(v37)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v38
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+20)) = uint8(base.B2i32(v37 == int32(112)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	v47 = F_lappend(m, v46, v29)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v47
	if v47 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v51 = v50
	goto L6
L5:
	;
	v51 = v9
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v29
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v51
	v55 = m.G0
	v57 = v55 - int32(48)
	m.G0 = v57
	v60 = F_palloc0(m, int32(36))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v332 = F_lappend(m, v331, v60)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L60
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(322)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+12)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l5)+56))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l5)+52))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+24)) = v76
	v80 = F_palloc0(m, v76<<(uint(int32(1))%32))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+28)) = v80
	if int32(0) < v74 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L56
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L52
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L49
	}
L13:
	;
	v92 = int32(0)
	v94 = int32(0)
	v99 = v9
	goto L16
L14:
	;
	v252 = v9
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+20)) = v252
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+32)) = v265
	m.G0 = v57 + int32(48)
	goto L7
L16:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v117 = v73 + v111<<(uint(int32(4))%32) + v92*int32(100)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+111)))
	if v118 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v252 = v231
	goto L15
L18:
	;
	v237 = v92 + int32(1)
	if v237 != v74 {
		v92 = v237
		v94 = v230
		v99 = v231
		goto L16
	} else {
		goto L48
	}
L19:
	;
	v122 = F_lappend(m, v99, int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v125 = v117 + int32(20)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+96))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v125)+76))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)+68))
	if l3 == l5 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v230 = v94
	v231 = v122
	goto L18
L23:
	;
	v131 = v92 + int32(1)
	v134 = F_makeVar(m, v51, base.I32_extend16_s(v131), v128, v127, v126, int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v143 = v117 + int32(24)
	if v76 <= v94 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v136 = F_lappend(m, v99, v134)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v80+v92<<(uint(int32(1))%32)))) = uint16(v131)
	v230 = v94
	v231 = v136
	goto L18
L28:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v207)+68))
	if v128 != v209 {
		goto L11
	} else {
		goto L43
	}
L29:
	;
	v186 = F_SearchSysCacheAttName(m, v72, v143)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L40
	}
L30:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v151 = v75 + v145<<(uint(int32(4))%32) + v94*int32(100)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+111)))
	if v152 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v156 = v151 + int32(24)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if base.B2i32(v159 == int32(0))|base.B2i32(v159 != v162) != 0 {
		v180 = v159
		v181 = v162
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v180-v181 == int32(0) {
		v207 = v151 + int32(20)
		v208 = v94
		goto L28
	} else {
		goto L39
	}
L33:
	;
	goto L32
L34:
	;
	v165 = v143
	v166 = v156
	goto L35
L35:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
	if v170 == int32(0) {
		v180 = v170
		v181 = v169
		goto L33
	} else {
		goto L37
	}
L36:
	;
	v180 = v170
	v181 = v169
	goto L33
L37:
	;
	v173 = int32(1)
	if v170 == v169 {
		v165 = v165 + v173
		v166 = v166 + v173
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	goto L29
L40:
	;
	if v186 == int32(0) {
		goto L12
	} else {
		goto L41
	}
L41:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v186)+16))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+22)))
	v193 = int32(*(*int16)(unsafe.Add(mBase, uint32(v190+v191)+74)))
	F_ReleaseCatCache(m, v186)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v207 = v75 + v196<<(uint(int32(4))%32) + v193*int32(100) - int32(80)
	v208 = v193 - int32(1)
	goto L28
L43:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v207)+76))
	if v127 != v211 {
		goto L11
	} else {
		goto L44
	}
L44:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v207)+96))
	if v126 != v213 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	v215 = int32(1)
	v218 = v208 + v215
	v221 = F_makeVar(m, v51, base.I32_extend16_s(v218), v128, v127, v126, int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v223 = F_lappend(m, v99, v221)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v227 = v92 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v208<<(uint(v215)%32)+v80))) = uint16(v227)
	v230 = v218
	v231 = v223
	goto L18
L48:
	;
	goto L17
L49:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v274 + int32(4)
	F_errmsg_internal(m, int32(_a_F_expand_single_inheritance_child_0), v57)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_expand_single_inheritance_child_1), int32(153), int32(_a_F_expand_single_inheritance_child_2))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
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
	F_errcode(m, int32(17064068))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+32)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v57)+36)) = v294 + int32(4)
	F_errmsg(m, int32(_a_F_expand_single_inheritance_child_3), v57+int32(32))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_expand_single_inheritance_child_1), int32(166), int32(_a_F_expand_single_inheritance_child_2))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
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
	F_errcode(m, int32(17064068))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+16)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v316 + int32(4)
	F_errmsg(m, int32(_a_F_expand_single_inheritance_child_4), v57+int32(16))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_expand_single_inheritance_child_1), int32(171), int32(_a_F_expand_single_inheritance_child_2))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v332
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v336 = F_copyObjectImpl(m, v335)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v336
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l5)+52))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	if int32(0) < v340 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+8))
	v352 = v340
	v354 = int32(0)
	v369 = v9
	goto L65
L63:
	;
	v436 = v9
	goto L64
L64:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	v440 = F_makeAlias(m, v439, v436)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L76
	}
L65:
	;
	v377 = v339 + v352<<(uint(int32(4))%32) + v354*int32(100)
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+111)))
	if v378 != 0 {
		v402 = int32(_a_F_expand_single_inheritance_child_5)
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v436 = v407
	goto L64
L67:
	;
	v403 = F_pstrdup(m, v402)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L72
	}
L68:
	;
	v379 = int32(0)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v60)+28))
	v385 = int32(*(*int16)(unsafe.Add(mBase, uint32(v381+v354<<(uint(int32(1))%32)))))
	if base.B2i32(v344 == v379)|base.B2i32(v385 <= v379) != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v402 = v377 + int32(24)
	goto L67
L70:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	if v389 < v385 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v344)+12))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v391+v385<<(uint(int32(2))%32)-int32(4))))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v397)+4))
	v402 = v398
	goto L67
L72:
	;
	v405 = F_makeString(m, v403)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v407 = F_lappend(m, v369, v405)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v410 = v354 + int32(1)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	if v410 < v411 {
		v352 = v411
		v354 = v410
		v369 = v407
		goto L65
	} else {
		goto L75
	}
L75:
	;
	goto L66
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v440
	v445 = v51 << (uint(int32(2)) % 32)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v445+v446))) = v29
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v449+v445))) = v60
	if l4 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v453 = F_palloc0(m, int32(36))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v527 = F_bms_is_member(m, l2, v526)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L97
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v453)+4)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v453))) = int32(374)
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v453)+8)) = v458
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v453)+12)) = v460
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
	v463 = m.G0
	v465 = v463 - int32(16)
	m.G0 = v465
	v467 = int32(5)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if v468 != 0 {
		v484 = v467
		goto L83
	} else {
		goto L84
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v453)+16)) = v484
	v504 = int32(1) << (uint(v484) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v453)+20)) = v504
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v453)+24)) = v506
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v453)+28)) = v508
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v453)+32)) = uint8(base.B2i32(v510 == int32(112)))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v514 | v504
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v518 = F_lappend(m, v517, v453)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L95
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L92
	}
L83:
	;
	m.G0 = v465 + int32(16)
	goto L81
L84:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+21)))
	if v469 == int32(102) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v473 = F_GetFdwRoutineByRelId(m, v472)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	if base.Ui32(int32(5)) <= base.Ui32(v462) {
		goto L82
	} else {
		goto L91
	}
L88:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v473)+104))
	if v475 == int32(0) {
		v484 = v467
		goto L83
	} else {
		goto L89
	}
L89:
	;
	v478 = m.T0[v475].(func(*base.Module, int32, int32) int32)(m, v29, v462)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v484 = v478
	goto L83
L91:
	;
	v484 = int32(4) - v462
	goto L83
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v465))) = v462
	F_errmsg_internal(m, int32(_a_F_expand_single_inheritance_child_6), v465)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_expand_single_inheritance_child_7), int32(2554), int32(_a_F_expand_single_inheritance_child_8))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v518
	goto L79
L96:
	;
	return
L97:
	;
	if v527 == int32(0) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v532 = F_bms_add_member(m, v531, v51)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v532
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+21)))
	if v535 == int32(112) {
		goto L96
	} else {
		goto L100
	}
L100:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v539 = F_bms_add_member(m, v538, v51)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v539
	v545 = int32(0)
	v547 = F_makeVar(m, v51, int32(-6), int32(26), int32(-1), v545, v545)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_add_row_identity_var(m, l0, v547, v51, int32(_a_F_expand_single_inheritance_child_9))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_add_row_identity_columns(m, l0, v51, v29, l5)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	goto L96
}
func F_extractNotNullColumn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v4 = F_SysCacheGetAttrNotNull(m, int32(19), l0, int32(21))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_pg_detoast_datum(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			if v10 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_extractNotNullColumn_0), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_extractNotNullColumn_1), int32(717), int32(_a_F_extractNotNullColumn_2))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
				if v13 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_extractNotNullColumn_0), int32(0))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_extractNotNullColumn_1), int32(717), int32(_a_F_extractNotNullColumn_2))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					if v14 != int32(21) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_extractNotNullColumn_0), int32(0))
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_extractNotNullColumn_1), int32(717), int32(_a_F_extractNotNullColumn_2))
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
						if v17 == int32(1) {
							v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8)+24)))
							return v33
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_extractNotNullColumn_0), int32(0))
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_extractNotNullColumn_1), int32(717), int32(_a_F_extractNotNullColumn_2))
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return int32(0)
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
	}
}
