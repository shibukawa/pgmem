package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_EA_get_flat_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L12
	} else {
		goto L48
	}
L2:
	;
	m.G0 = v15 + int32(32)
	return v150
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v150 = int32(base.Ui32(v18) >> (uint(int32(2)) % 32))
	goto L2
L4:
	;
	goto L5
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v21 != 0 {
		v150 = v21
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+47)))
	switch v27 - int32(99) {
	case 0:
		v49 = int32(1)
		goto L7
	case 1:
		goto L10
	default:
		goto L9
	case 6:
		goto L11
	case 16:
		goto L8
	}
L7:
	;
	if int32(0) < v25 {
		goto L16
	} else {
		goto L17
	}
L8:
	;
	v49 = int32(2)
	goto L7
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v49 = int32(8)
	goto L7
L11:
	;
	v49 = int32(4)
	goto L7
L12:
	;
	return int32(0)
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = base.I32_extend8_s(v27)
	F_errmsg_internal(m, int32(_a_F_EA_get_flat_size_0), v15)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_EA_get_flat_size_1), int32(322), int32(_a_F_EA_get_flat_size_2))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	v52 = int32(0)
	v59 = v52
	v64 = v2
	goto L19
L17:
	;
	v127 = v2
	goto L18
L18:
	;
	v133 = v24 << (uint(int32(3)) % 32)
	v137 = base.I32_div_s(v25+int32(7), int32(8))
	v139 = int32(23)
	if v22 != 0 {
		goto L45
	} else {
		goto L46
	}
L19:
	;
	if v22 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v127 = v116
	goto L18
L21:
	;
	v118 = v59 + int32(1)
	if v118 != v25 {
		v59 = v118
		v64 = v116
		goto L19
	} else {
		goto L44
	}
L22:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v22))))
	if v70 != 0 {
		v116 = v64
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v71 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+44)))
	if int32(0) < v71 {
		v107 = v71
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	v111 = (v64 + (v49 - int32(1)) + v107) & (v52 - v49)
	if base.Ui32(int32(1073741824)) <= base.Ui32(v111) {
		goto L1
	} else {
		goto L43
	}
L27:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v23+v59<<(uint(int32(3))%32))))
	if v71 == int32(-1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v80 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v104 = F_strlen(m, v77)
	mBase = m.M
	v107 = v104 + int32(1)
	goto L26
L31:
	;
	v84 = int32(18)
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	if v86 == v84 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	if v80&int32(1) != 0 {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	v89 = v84
	goto L36
L35:
	;
	v89 = int32(2)
	goto L36
L36:
	;
	if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v96 = int32(6)
	goto L39
L38:
	;
	v96 = v89
	goto L39
L39:
	;
	v107 = v96
	goto L26
L40:
	;
	v107 = int32(base.Ui32(v80) >> (uint(int32(1)) % 32))
	goto L26
L41:
	;
	goto L42
L42:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v107 = int32(base.Ui32(v101) >> (uint(int32(2)) % 32))
	goto L26
L43:
	;
	v116 = v111
	goto L21
L44:
	;
	goto L20
L45:
	;
	v143 = v133 + v137 + v139
	goto L47
L46:
	;
	v143 = v133 + v139
	goto L47
L47:
	;
	v146 = v143&int32(-8) + v127
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v146
	v150 = v146
	goto L2
L48:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(1073741823)
	F_errmsg(m, int32(_a_F_EA_get_flat_size_3), v15+int32(16))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_EA_get_flat_size_4), int32(277), int32(_a_F_EA_get_flat_size_5))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
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
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
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
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	if l2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L13
	} else {
		goto L42
	}
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v21 = v19
	goto L4
L3:
	;
	v21 = int32(0)
	goto L4
L4:
	;
	if v18 == v21 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if v18 == int32(0) {
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
	v167 = m.ExcPending
	if v167 != 0 {
		goto L13
	} else {
		goto L37
	}
L8:
	;
	m.G0 = v15 + int32(48)
	return v153
L9:
	;
	v153 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v27 = F_copyObjectImpl(m, l2)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v88 = F_ExecPrepareExprList(m, v27, l3)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L13
	} else {
		goto L25
	}
L13:
	;
	return int32(0)
L14:
	;
	if v27 == int32(0) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v33 <= int32(0) {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v39 = int32(0)
	goto L17
L17:
	;
	v50 = v39 + int32(1)
	v52 = v39 << (uint(int32(2)) % 32)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v26+v52)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v56 = v55 + v52
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v59 = F_transformExpr(m, l0, v57, int32(36))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L13
	} else {
		goto L19
	}
L18:
	;
	goto L12
L19:
	;
	v61 = F_exprType(m, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v63 = int32(-1)
	v67 = F_coerce_to_target_type(m, l0, v59, v61, v54, v63, int32(1), int32(2), v63)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	if v67 == int32(0) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_assign_expr_collations(m, l0, v67)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v67
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v50 < v74 {
		v39 = v50
		goto L17
	} else {
		goto L24
	}
L24:
	;
	goto L18
L25:
	;
	v90 = F_makeParamList(m, v18)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L13
	} else {
		goto L26
	}
L26:
	;
	if v88 == int32(0) {
		v153 = v90
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v94 = int32(0)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v95 <= v94 {
		v153 = v90
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v100 = v94
	goto L29
L29:
	;
	v113 = v100 << (uint(int32(2)) % 32)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v113+v114)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v113+v26)))
	v121 = v90 + int32(32) + v100<<(uint(int32(4))%32)
	v122 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v121)+10)) = uint16(v122)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+12)) = v118
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l3)+152))
	if v125 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v153 = v90
	goto L8
L31:
	;
	v128 = F_MakePerTupleExprContext(m, l3)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L13
	} else {
		goto L34
	}
L32:
	;
	v130 = v125
	goto L33
L33:
	;
	v131 = int32(_a_F_EvaluateParams_0)
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_EvaluateParams[0]))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v130)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_EvaluateParams[0])) = v134
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v116)+24))
	v139 = m.T0[v138].(func(*base.Module, int32, int32, int32) int64)(m, v116, v130, v121+int32(8))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L13
	} else {
		goto L35
	}
L34:
	;
	v130 = v128
	goto L33
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_EvaluateParams[0])) = v132
	*(*int64)(unsafe.Add(mBase, uint32(v121))) = v139
	v145 = v100 + int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v145 < v146 {
		v100 = v145
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
	v170 = m.ExcPending
	if v170 != 0 {
		goto L13
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l1
	F_errmsg(m, int32(_a_F_EvaluateParams_1), v15+int32(32))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L13
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v18
	v182 = F_errdetail(m, int32(_a_F_EvaluateParams_2), v15+int32(16))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L13
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_EvaluateParams_3), int32(300), int32(_a_F_EvaluateParams_4))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
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
	v195 = m.ExcPending
	if v195 != 0 {
		goto L13
	} else {
		goto L43
	}
L43:
	;
	v196 = F_format_type_be(m, v61)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	v198 = F_format_type_be(m, v54)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v50
	F_errmsg(m, int32(_a_F_EvaluateParams_5), v15)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	F_errhint(m, int32(_a_F_EvaluateParams_6), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v211 = F_exprLocation(m, v210)
	mBase = m.M
	F_parser_errposition(m, l0, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_EvaluateParams_3), int32(337), int32(_a_F_EvaluateParams_4))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
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
	*(*int64)(unsafe.Add(mBase, uint32(v15)+20)) = int64(55834575296)
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
	var v242 int32
	_ = v242
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
	*(*int64)(unsafe.Add(mBase, uint32(v16)+52)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v16)+44)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v16)+36)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v16)+28)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = int64(51539608000)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v29
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
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
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
	v71 = F_ExecFetchSlotHeapTuple(m, l2, int32(1), v16+int32(62))
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = l2
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
	v86 = F_ExecCallTriggerFunc(m, v16+int32(16), v45, v79, v80, v85)
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
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+62)))
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
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+24))
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
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113*int32(100)+(v98+v120<<(uint(int32(3))%32)))+118)))
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
	*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = int64(0)
	v138 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+63)) = uint8(v138)
	v147 = F_heap_modify_tuple_by_cols(m, v115, v98, v138, v16+int32(76), v16-int32(-64), v16+int32(63))
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
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+62)))
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
	v242 = F_errdetail(m, int32(_a_F_ExecBRInsertTriggers_1), v16)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_ExecBRInsertTriggers_2), int32(2545), int32(_a_F_ExecBRInsertTriggers_3))
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
	var v72 int64
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
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v335 int32
	_ = v335
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
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+46)) = uint8(v26)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+45)) = uint8(v26)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v26
	v32 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+28)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v19)+20)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v19)+12)) = v32
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
	return v335
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = int64(60129542592)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v110
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
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(0)
	v48 = F_GetTupleForTrigger(m, l0, l1, l2, l3, v38, v22, l8^int32(1), v17+int32(-16), l6, l7)
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
		v335 = v10
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	if v52 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v52
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+80))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
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
	v99 = F_ExecFetchSlotHeapTuple(m, v22, int32(1), v17+int32(-18))
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
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v53)+32))
	v72 = m.T0[v71].(func(*base.Module, int32, int32, int32) int64)(m, v53+int32(8), v57, int32(0))
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
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v112
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
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+46)))
	if v316 != int32(1) {
		v335 = v115
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
		v290 = v130
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L26
L29:
	;
	v297 = v127 + int32(1)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v297 < v298 {
		v127 = v297
		v130 = v290
		goto L27
	} else {
		goto L75
	}
L30:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
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
		v290 = v130
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
	v155 = F_ExecFetchSlotHeapTuple(m, l5, int32(1), v17+int32(-19))
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
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v139
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v165 != 0 {
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
	v168 = v165
	goto L39
L38:
	;
	v166 = F_MakePerTupleExprContext(m, l0)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+20))
	v170 = F_ExecCallTriggerFunc(m, v19, v127, v163, v164, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	v168 = v166
	goto L39
L41:
	;
	if v170 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+46)))
	if v174 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	if v170 == v157 {
		goto L51
	} else {
		goto L52
	}
L45:
	;
	F_pfree(m, v106)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v179 = int32(0)
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+45)))
	if v180 != int32(1) {
		v335 = v179
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
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v335 = v179
	goto L4
L51:
	;
	v290 = v170
	goto L29
L52:
	;
	goto L53
L53:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+52))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+24))
	if v188 == int32(0) {
		v245 = v170
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_ExecForceStoreHeapTuple(m, v245, l5, int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L66
	}
L55:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+18)))
	if v191 != int32(1) {
		v245 = v170
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v194 = int32(0)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	if v195 <= v194 {
		v245 = v170
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v199 = v170
	v201 = v195
	v206 = v194
	goto L58
L58:
	;
	v217 = v206 + int32(1)
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206*int32(100)+(v187+v201<<(uint(int32(3))%32)))+118)))
	if v222 != int32(118) {
		v241 = v199
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v245 = v241
	goto L54
L60:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	if v217 < v242 {
		v199 = v241
		v201 = v242
		v206 = v217
		goto L58
	} else {
		goto L65
	}
L61:
	;
	v225 = F_heap_attisnull(m, v199, v217, v187)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	if v225 != 0 {
		v241 = v199
		goto L60
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v217
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = int64(0)
	v230 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+47)) = uint8(v230)
	v239 = F_heap_modify_tuple_by_cols(m, v199, v187, v230, v17+int32(-4), v17+int32(-16), v17+int32(-17))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v241 = v239
	goto L60
L65:
	;
	goto L59
L66:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+46)))
	if base.B2i32(v263 != int32(1))|base.B2i32(v245 != v106) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+28))
	m.T0[v271].(func(*base.Module, int32))(m, l5)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+45)))
	if v274 == int32(1) {
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
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v290 = int32(0)
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
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v335 = v115
	goto L4
}
func F_ExecBuildHash32Expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v23 int32
	_ = v23
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
	var v37 int64
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
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
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
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
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v287 int64
	_ = v287
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	v6 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(16)
	m.G0 = v25
	v28 = F_palloc0(m, int32(72))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(386)
	if l2 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v35 = v34
	goto L5
L4:
	;
	v35 = v6
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+44)) = l4
	v37 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v37
	v41 = F_expr_setup_walker(m, l2, v25)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	F_ExecPushExprSetupSteps(m, v28, v25)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if int64(2) <= base.I64_extend_i32_s(v35) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v49 = F_palloc(m, int32(16))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v51 = v6
	goto L10
L10:
	;
	v54 = int32(8)
	v67 = int32(86)
	v69 = int32(0)
	v70 = v6
	v71 = v6
	v75 = v6
	v76 = v6
	v77 = int32(87)
	v78 = v6
	v84 = v6
	goto L14
L11:
	;
	v51 = v49
	goto L10
L12:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v280 + int32(1)
	v286 = v279 + v280*int32(40)
	v287 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v286)+8)) = v287
	*(*int64)(unsafe.Add(mBase, uint32(v286))) = v287
	*(*int32)(unsafe.Add(mBase, uint32(v286)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v286)+32)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v286)+28)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v286)+24)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v286)+20)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v286)+16)) = v76
	v298 = F_jit_compile_expr(m, v28)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L61
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v277
	v279 = v277
	goto L12
L14:
	;
	v85 = int32(0)
	if l2 == v85 {
		v95 = v85
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	if v266 != v170 {
		goto L57
	} else {
		goto L58
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
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v89 <= v71 {
		v95 = int32(0)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v95 = v91 + v71<<(uint(int32(2))%32)
	goto L16
L19:
	;
	goto L15
L20:
	;
	v178 = v71 << (uint(int32(2)) % 32)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v103+v178)))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0+v178)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v185 = F_palloc0(m, int32(28))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L33
	}
L21:
	;
	if v75 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(v95 == int32(0))|base.B2i32(v100 <= v71) != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v103 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
	if v170 != 0 {
		goto L19
	} else {
		goto L31
	}
L26:
	;
	v107 = int32(0)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v108 <= v107 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v115 = v107
	goto L28
L28:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v134+v115<<(uint(int32(2))%32))))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v133+v138*int32(40))+28)) = v142
	v145 = v115 + int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v145 < v146 {
		v115 = v145
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
	v171 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+40)) = v171
	v175 = F_palloc_mul(m, int32(40), v171)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v277 = v175
	goto L13
L33:
	;
	v188 = F_palloc0(m, int32(40))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_fmgr_info(m, v182, v185)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_ExecInitExprRec(m, v183, v28, v188+int32(24), v188+int32(32))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v198 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v188)+18)) = uint16(v198)
	v200 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v188)+16)) = uint8(v200)
	*(*int32)(unsafe.Add(mBase, uint32(v188)+12)) = v180
	*(*int64)(unsafe.Add(mBase, uint32(v188)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = v185
	v206 = base.B2i32(v71 == v35-int32(1))
	if v71 == v35-int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v207 = v28 + v54
	goto L39
L38:
	;
	v207 = v51
	goto L39
L39:
	;
	if v71 == v35-int32(1) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v208 = v28 + int32(5)
	goto L42
L41:
	;
	v208 = v51 + v54
	goto L42
L42:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v71))))
	if v210 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v211 = v77
	goto L45
L44:
	;
	v211 = v67
	goto L45
L45:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
	if v213 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	v237 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v236 + v237
	v242 = v235 + v236*int32(40)
	v243 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v242)+36)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v242)+32)) = v51
	v246 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v242)+28)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v242)+24)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v242)+20)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v242)+16)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v242)+12)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v242)+8)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v242)+4)) = v207
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v211
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	v264 = F_lappend_int(m, v75, v261-v237)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L56
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v233
	v235 = v233
	goto L46
L48:
	;
	v216 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+40)) = v216
	v220 = F_palloc_mul(m, int32(40), v216)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	if v222 != v213 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v233 = v220
	goto L47
L52:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v235 = v224
	goto L46
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+40)) = v213 << (uint(int32(1)) % 32)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v231 = F_repalloc(m, v228, v213*int32(80))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v233 = v231
	goto L47
L56:
	;
	v67 = int32(88)
	v69 = v188
	v70 = v51
	v71 = v71 + v237
	v75 = v264
	v76 = v185
	v77 = int32(89)
	v78 = v212
	v84 = v246
	goto L14
L57:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v279 = v268
	goto L12
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+40)) = v170 << (uint(int32(1)) % 32)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v275 = F_repalloc(m, v272, v170*int32(80))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v277 = v275
	goto L13
L61:
	;
	if v298 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	F_ExecReadyInterpretedExpr(m, v28)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	m.G0 = v25 + int32(16)
	return v28
L65:
	;
	goto L64
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
	v45 = int32(42)
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
	v47 = int32(52)
	goto L15
L21:
	;
	v47 = int32(38)
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
	var v63 int32
	_ = v63
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
	var v93 int32
	_ = v93
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
	var v134 int32
	_ = v134
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
	var v166 int32
	_ = v166
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
	var v186 int32
	_ = v186
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
	v41 = v20*int32(100) + (v12 + v27<<(uint(int32(3))%32))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+119)))
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
	v72 = v41 + int32(28)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+90)))
	if v73 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+32)))
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
	v63 = F_errdetail(m, int32(_a_F_ExecCheckPlanOutput_1), v10)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_ExecCheckPlanOutput_2), int32(249), int32(_a_F_ExecCheckPlanOutput_3))
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
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+32)))
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
	v93 = F_errdetail(m, int32(_a_F_ExecCheckPlanOutput_4), v10+int32(32))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L20
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_ExecCheckPlanOutput_2), int32(266), int32(_a_F_ExecCheckPlanOutput_3))
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
	v134 = F_errdetail(m, int32(_a_F_ExecCheckPlanOutput_5), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L20
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_ExecCheckPlanOutput_2), int32(229), int32(_a_F_ExecCheckPlanOutput_3))
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
	v166 = F_errdetail(m, int32(_a_F_ExecCheckPlanOutput_6), v10+int32(16))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L20
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_ExecCheckPlanOutput_2), int32(278), int32(_a_F_ExecCheckPlanOutput_3))
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
	v186 = F_errdetail(m, int32(_a_F_ExecCheckPlanOutput_7), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L20
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_ExecCheckPlanOutput_2), int32(285), int32(_a_F_ExecCheckPlanOutput_3))
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
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
	var v150 int32
	_ = v150
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
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
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
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
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
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
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+16)))
	if v20 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v321 = m.G0
	v323 = v321 - int32(32)
	m.G0 = v323
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)+52))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v334 != 0 {
		goto L89
	} else {
		goto L90
	}
L2:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	v93 = int32(*(*int16)(unsafe.Add(mBase, uint32(v92)+122)))
	if v93 <= int32(0) {
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
	v45 = v18 + v33<<(uint(int32(3))%32) + v31*int32(100)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+14)))
	if v46 != int32(1) {
		v70 = v32
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v70 == int32(0) {
		goto L2
	} else {
		goto L20
	}
L7:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v73 = v37 + int32(1)
	v74 = base.I32_extend16_s(v73)
	if v74 <= v71 {
		v31 = v74
		v32 = v70
		v33 = v71
		v37 = v73
		goto L5
	} else {
		goto L19
	}
L8:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45-int32(72))+90)))
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
	v70 = v54
	goto L7
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	m.T0[v60].(func(*base.Module, int32, int32))(m, l1, v31)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v65 = int32(1)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+v31-v65))))
	if v67 == v65 {
		v317 = v31
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v70 = v32
	goto L7
L19:
	;
	goto L6
L20:
	;
	v78 = F_ExecRelGenVirtualNotNull(m, l0, l1, l2, v70)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	if v78 != 0 {
		v317 = v78
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
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+52))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+14)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96)+48))
	v101 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100)+122)))
	if v99 == v101 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v104 == int32(0) {
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
	v284 = m.ExcPending
	if v284 != 0 {
		goto L12
	} else {
		goto L85
	}
L28:
	;
	v108 = int32(_a_F_ExecConstraints_0)
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_ExecConstraints[0]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l2)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecConstraints[0])) = v111
	v114 = F_palloc0_mul(m, int32(4), v99)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L12
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l2)+152))
	if v178 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v114
	if v99 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v120 = int32(0)
	goto L35
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecConstraints[0])) = v109
	goto L30
L35:
	;
	v131 = v103 + v120*int32(12)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+8)))
	if v132 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L34
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v136 = F_stringToNode(m, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L12
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v150 = v120 + int32(1)
	if v150 != v99 {
		v120 = v150
		goto L35
	} else {
		goto L43
	}
L40:
	;
	v139 = F_expand_generated_columns_in_expr(m, v136, v96, int32(1))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L12
	} else {
		goto L41
	}
L41:
	;
	v141 = F_ExecPrepareExpr(m, v139, l2)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v143+v120<<(uint(int32(2))%32)))) = v141
	goto L39
L43:
	;
	goto L36
L44:
	;
	v181 = F_MakePerTupleExprContext(m, l2)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L12
	} else {
		goto L47
	}
L45:
	;
	v183 = v178
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = l1
	if v99 == int32(0) {
		goto L23
	} else {
		goto L48
	}
L47:
	;
	v183 = v181
	goto L46
L48:
	;
	v191 = int32(0)
	goto L49
L49:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v200+v191<<(uint(int32(2))%32))))
	if v204 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v103+v191*int32(12))))
	if v215 == int32(0) {
		goto L23
	} else {
		goto L58
	}
L51:
	;
	goto L50
L52:
	;
	v205 = F_ExecCheck(m, v204, v183)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L12
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v210 = v191 + int32(1)
	if v210 != v99 {
		v191 = v210
		goto L49
	} else {
		goto L57
	}
L55:
	;
	if v205 == int32(0) {
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
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v218 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+56))
	v251 = F_ExecBuildSlotValueDescription(m, v250, v246, v248, v247)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L12
	} else {
		goto L75
	}
L60:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v218)+8))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+52))
	v223 = F_build_attrmap_by_name_if_req(m, v219, v221, int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L12
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v239 = F_ExecGetInsertedCols(m, l0, l2)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L12
	} else {
		goto L72
	}
L63:
	;
	if v223 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v227 = F_MakeTupleTableSlot(m, v221, int32(_a_F_ExecConstraints_1), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L12
	} else {
		goto L67
	}
L65:
	;
	v231 = l1
	goto L66
L66:
	;
	v232 = F_ExecGetInsertedCols(m, v218, l2)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L12
	} else {
		goto L69
	}
L67:
	;
	v229 = F_execute_attr_map_slot(m, v223, l1, v227)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L12
	} else {
		goto L68
	}
L68:
	;
	v231 = v229
	goto L66
L69:
	;
	v234 = F_ExecGetUpdatedCols(m, v218, l2)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L12
	} else {
		goto L70
	}
L70:
	;
	v236 = F_bms_union(m, v232, v234)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v218)+8))
	v246 = v231
	v247 = v236
	v248 = v221
	v249 = v238
	goto L59
L72:
	;
	v241 = F_ExecGetUpdatedCols(m, l0, l2)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	v243 = F_bms_union(m, v239, v241)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	v246 = l1
	v247 = v243
	v248 = v18
	v249 = v17
	goto L59
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L12
	} else {
		goto L76
	}
L76:
	;
	F_errcode(m, int32(67391682))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L12
	} else {
		goto L77
	}
L77:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v260 + int32(4)
	F_errmsg(m, int32(_a_F_ExecConstraints_2), v15+int32(16))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L12
	} else {
		goto L78
	}
L78:
	;
	if v251 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v251
	v272 = F_errdetail(m, int32(_a_F_ExecConstraints_3), v15)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L12
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	F_errtableconstraint(m, v17, v215)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L12
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	F_errfinish(m, int32(_a_F_ExecConstraints_4), int32(2108), int32(_a_F_ExecConstraints_5))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
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
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v96)+48))
	v286 = int32(*(*int16)(unsafe.Add(mBase, uint32(v285)+122)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v285 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v286 - v99
	F_errmsg_internal(m, int32(_a_F_ExecConstraints_6), v15+int32(32))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L12
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_ExecConstraints_4), int32(1825), int32(_a_F_ExecConstraints_7))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
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
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)+56))
	v366 = F_ExecBuildSlotValueDescription(m, v365, v361, v360, v362)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L12
	} else {
		goto L104
	}
L89:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)+8))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+52))
	v338 = F_build_attrmap_by_name_if_req(m, v326, v336, int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L12
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v354 = F_ExecGetInsertedCols(m, l0, l2)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L12
	} else {
		goto L101
	}
L92:
	;
	if v338 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v342 = F_MakeTupleTableSlot(m, v336, int32(_a_F_ExecConstraints_1), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L12
	} else {
		goto L96
	}
L94:
	;
	v346 = l1
	goto L95
L95:
	;
	v347 = F_ExecGetInsertedCols(m, v334, l2)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L12
	} else {
		goto L98
	}
L96:
	;
	v344 = F_execute_attr_map_slot(m, v338, l1, v342)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L12
	} else {
		goto L97
	}
L97:
	;
	v346 = v344
	goto L95
L98:
	;
	v349 = F_ExecGetUpdatedCols(m, v334, l2)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L12
	} else {
		goto L99
	}
L99:
	;
	v351 = F_bms_union(m, v347, v349)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L12
	} else {
		goto L100
	}
L100:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v334)+8))
	v360 = v336
	v361 = v346
	v362 = v351
	v364 = v353
	goto L88
L101:
	;
	v356 = F_ExecGetUpdatedCols(m, l0, l2)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L12
	} else {
		goto L102
	}
L102:
	;
	v358 = F_bms_union(m, v354, v356)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L12
	} else {
		goto L103
	}
L103:
	;
	v360 = v326
	v361 = l1
	v362 = v358
	v364 = v325
	goto L88
L104:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L12
	} else {
		goto L105
	}
L105:
	;
	F_errcode(m, int32(33575106))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L12
	} else {
		goto L106
	}
L106:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v325)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v323)+16)) = v326 + v327<<(uint(int32(3))%32) + v317*int32(100) - int32(68)
	*(*int32)(unsafe.Add(mBase, uint32(v323)+20)) = v375 + int32(4)
	F_errmsg(m, int32(_a_F_ExecConstraints_8), v323+int32(16))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L12
	} else {
		goto L107
	}
L107:
	;
	if v366 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = v366
	v389 = F_errdetail(m, int32(_a_F_ExecConstraints_3), v323)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L12
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	F_errtablecol(m, v325, v317)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L12
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	F_errfinish(m, int32(_a_F_ExecConstraints_4), int32(2246), int32(_a_F_ExecConstraints_9))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int64
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v12 < v11 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
		m.T0[v15].(func(*base.Module, int32, int32))(m, l1, v11)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
			m.T0[v27].(func(*base.Module, int32))(m, v25)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				if int32(0) < v24 {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
					v36 = int32(0)
					for {
						v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20+v36<<(uint(int32(1))%32)))))
						if v48 == int32(0) {
							*(*int64)(unsafe.Add(mBase, uint32(v33+v36<<(uint(int32(3))%32)))) = int64(0)
							v70 = int32(1)
						} else {
							v57 = int32(3)
							v61 = v48 - int32(1)
							v65 = *(*int64)(unsafe.Add(mBase, uint32(v22+v61<<(uint(v57)%32))))
							*(*int64)(unsafe.Add(mBase, uint32(v33+v36<<(uint(v57)%32)))) = v65
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v21))))
							v70 = v68
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v36+v32))) = uint8(v70)
						v73 = v36 + int32(1)
						if v73 != v24 {
							v36 = v73
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
				v86 = v84 & int32(_a_F_ExecFilterJunk_0)
				*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)) = uint16(v86)
				v88 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
				v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
				*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)) = uint16(v89)
				return v25
			}
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
		m.T0[v27].(func(*base.Module, int32))(m, v25)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			if int32(0) < v24 {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
				v36 = int32(0)
				for {
					v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20+v36<<(uint(int32(1))%32)))))
					if v48 == int32(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v33+v36<<(uint(int32(3))%32)))) = int64(0)
						v70 = int32(1)
					} else {
						v57 = int32(3)
						v61 = v48 - int32(1)
						v65 = *(*int64)(unsafe.Add(mBase, uint32(v22+v61<<(uint(v57)%32))))
						*(*int64)(unsafe.Add(mBase, uint32(v33+v36<<(uint(v57)%32)))) = v65
						v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v21))))
						v70 = v68
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v36+v32))) = uint8(v70)
					v73 = v36 + int32(1)
					if v73 != v24 {
						v36 = v73
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
			v86 = v84 & int32(_a_F_ExecFilterJunk_0)
			*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)) = uint16(v86)
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
			*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)) = uint16(v89)
			return v25
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int64
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
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v143 int64
	_ = v143
	var v144 int64
	_ = v144
	var v147 int64
	_ = v147
	var v148 int32
	_ = v148
	var v149 int64
	_ = v149
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v182 int64
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int64
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int64
	_ = v257
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v295 int32
	_ = v295
	var v300 int64
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int64
	_ = v336
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int64
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v375 int64
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v453 int32
	_ = v453
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int64
	_ = v484
	var v485 int32
	_ = v485
	var v490 int64
	_ = v490
	var v495 int32
	_ = v495
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v521 int64
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v575 int32
	_ = v575
	var v576 int64
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v766 int32
	_ = v766
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v807 int64
	_ = v807
	var v812 int32
	_ = v812
	var v817 int64
	_ = v817
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v848 int64
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int64
	_ = v957
	var v960 int32
	_ = v960
	var v964 int64
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1128 int32
	_ = v1128
	var v1133 int32
	_ = v1133
	var v1138 int32
	_ = v1138
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1156 int64
	_ = v1156
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1176 int32
	_ = v1176
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1211 int64
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1220 int64
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int64
	_ = v1242
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1251 int64
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1253 int64
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1258 int64
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1283 int32
	_ = v1283
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1325 int64
	_ = v1325
	var v1329 int32
	_ = v1329
	var v1335 int32
	_ = v1335
	var v1339 int32
	_ = v1339
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1358 int32
	_ = v1358
	var v1366 int32
	_ = v1366
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1423 int32
	_ = v1423
	var v1428 int32
	_ = v1428
	var v1433 int32
	_ = v1433
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1442 int64
	_ = v1442
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1454 int32
	_ = v1454
	var v1466 int32
	_ = v1466
	var v1484 int32
	_ = v1484
	var v1488 int32
	_ = v1488
	var v1489 int64
	_ = v1489
	var v1495 int32
	_ = v1495
	var v1504 int32
	_ = v1504
	var v1510 int32
	_ = v1510
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
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1529 int64
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1550 int64
	_ = v1550
	var v1556 int32
	_ = v1556
	var v1559 int32
	_ = v1559
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1567 int64
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1569 int64
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1572 int32
	_ = v1572
	var v1574 int64
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1609 int32
	_ = v1609
	var v1617 int32
	_ = v1617
	var v1619 int32
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1664 int32
	_ = v1664
	var v1668 int32
	_ = v1668
	var v1674 int32
	_ = v1674
	var v1679 int32
	_ = v1679
	var v1684 int32
	_ = v1684
	var v1689 int32
	_ = v1689
	var v1694 int32
	_ = v1694
	var v1699 int32
	_ = v1699
	var v1704 int32
	_ = v1704
	var v1732 int32
	_ = v1732
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1756 int64
	_ = v1756
	var v1758 int64
	_ = v1758
	var v1760 int64
	_ = v1760
	var v1762 int64
	_ = v1762
	var v1764 int64
	_ = v1764
	var v1766 int64
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1778 int32
	_ = v1778
	var v1781 int32
	_ = v1781
	var v1785 int32
	_ = v1785
	var v1787 int32
	_ = v1787
	var v1812 int32
	_ = v1812
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1859 int32
	_ = v1859
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1870 int32
	_ = v1870
	v2 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(672)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v32 - int32(9) {
	case 0:
		goto L12
	default:
		goto L11
	case 3, 41:
		goto L2
	case 7:
		goto L3
	case 8:
		goto L4
	case 10, 20, 26:
		goto L5
	case 12:
		goto L6
	case 13:
		goto L7
	case 18:
		goto L10
	case 28:
		goto L8
	case 29, 33:
		goto L13
	case 34:
		goto L9
	}
L1:
	;
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L377
L2:
	;
	F_ExecGrant_common(m, l0, int32(1247), int64(256), int32(496))
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L14
	} else {
		goto L376
	}
L3:
	;
	F_ExecGrant_common(m, l0, int32(2328), int64(256), int32(0))
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L14
	} else {
		goto L375
	}
L4:
	;
	F_ExecGrant_common(m, l0, int32(1417), int64(256), int32(0))
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L14
	} else {
		goto L374
	}
L5:
	;
	F_ExecGrant_common(m, l0, int32(1255), int64(128), int32(0))
	mBase = m.M
	v1689 = m.ExcPending
	if v1689 != 0 {
		goto L14
	} else {
		goto L373
	}
L6:
	;
	F_ExecGrant_common(m, l0, int32(2612), int64(256), int32(495))
	mBase = m.M
	v1684 = m.ExcPending
	if v1684 != 0 {
		goto L14
	} else {
		goto L372
	}
L7:
	;
	v1439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v1439 != int32(1) {
		goto L328
	} else {
		goto L329
	}
L8:
	;
	F_ExecGrant_common(m, l0, int32(2615), int64(768), int32(0))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L14
	} else {
		goto L327
	}
L9:
	;
	F_ExecGrant_common(m, l0, int32(1213), int64(512), int32(0))
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L14
	} else {
		goto L326
	}
L10:
	;
	v1153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v1153 != int32(1) {
		goto L259
	} else {
		goto L260
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L14
	} else {
		goto L256
	}
L12:
	;
	F_ExecGrant_common(m, l0, int32(1262), int64(3584), int32(0))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L14
	} else {
		goto L255
	}
L13:
	;
	v37 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	v41 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v43 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L14
	} else {
		goto L252
	}
L18:
	;
	F_relation_close(m, v41, int32(3))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L14
	} else {
		goto L250
	}
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v46 <= int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v69 = v2
	goto L21
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77+v69<<(uint(int32(2))%32))))
	v82 = base.I64_extend_i32_u(v81)
	v83 = F_SearchSysCacheLocked1(m, int32(57), v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L14
	} else {
		goto L30
	}
L22:
	;
	goto L18
L23:
	;
	v766 = int32(0)
	if base.B2i32(v742 == v766)|base.B2i32(v183 < int32(-7)) == v766 {
		goto L181
	} else {
		goto L182
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L14
	} else {
		goto L178
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L14
	} else {
		goto L173
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L14
	} else {
		goto L168
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L14
	} else {
		goto L165
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L14
	} else {
		goto L161
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L14
	} else {
		goto L157
	}
L30:
	;
	if v83 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+22)))
	v87 = v85 + v86
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+119)))
	switch v88 - int32(99) {
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
	v618 = m.ExcPending
	if v618 != 0 {
		goto L14
	} else {
		goto L154
	}
L34:
	;
	v134 = base.B2i32(v88 == int32(83))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.B2i32(v134 == int32(0))&base.B2i32(v137 == int32(38)) != 0 {
		goto L29
	} else {
		goto L47
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L14
	} else {
		goto L43
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L14
	} else {
		goto L39
	}
L37:
	;
	if v88 != int32(73) {
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
	v99 = m.ExcPending
	if v99 != 0 {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+160)) = v87 + int32(4)
	F_errmsg(m, int32(_a_F_ExecGrantStmt_oids_0), v30+int32(160))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1811), int32(_a_F_ExecGrantStmt_oids_2))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
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
	v119 = m.ExcPending
	if v119 != 0 {
		goto L14
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+176)) = v87 + int32(4)
	F_errmsg(m, int32(_a_F_ExecGrantStmt_oids_3), v30+int32(176))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L14
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1818), int32(_a_F_ExecGrantStmt_oids_2))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
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
	if v88 == int32(83) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v143 = int64(262)
	goto L50
L49:
	;
	v143 = int64(16511)
	goto L50
L50:
	;
	v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v144 == int64(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v147 = v143
	goto L53
L52:
	;
	v147 = v144
	goto L53
L53:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v148 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v149 = v147
	goto L56
L55:
	;
	v149 = v144
	goto L56
L56:
	;
	if v137 != int32(42) {
		v182 = v149
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v183 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87)+120)))
	v185 = v183 + int32(8)
	v188 = F_palloc0(m, v185<<(uint(int32(3))%32))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L14
	} else {
		goto L71
	}
L58:
	;
	if v88 == int32(83) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	if v149&int64(-263) == int64(0) {
		v182 = v149
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v149&int64(-16512) != int64(0) {
		goto L28
	} else {
		goto L70
	}
L62:
	;
	v158 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L14
	} else {
		goto L63
	}
L63:
	;
	if v158 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L14
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v182 = v149 & int64(262)
	goto L57
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+112)) = v87 + int32(4)
	F_errmsg(m, int32(_a_F_ExecGrantStmt_oids_4), v30+int32(112))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L14
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1864), int32(_a_F_ExecGrantStmt_oids_2))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L14
	} else {
		goto L69
	}
L69:
	;
	goto L66
L70:
	;
	v182 = v149
	goto L57
L71:
	;
	v190 = int32(0)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v191 != 0 {
		v271 = v190
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v87)+80))
	v300 = F_SysCacheGetAttr(m, int32(57), v83, int32(32), v30+int32(667))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L14
	} else {
		goto L89
	}
L73:
	;
	v193 = v182 & int64(39)
	if v193 == int64(0) {
		v271 = v190
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v196 = int32(-6)
	v199 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87)+120)))
	if v199 < v196 {
		v271 = int32(1)
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v204 = int32(_a_F_ExecGrantStmt_oids_5)
	v208 = v196
	goto L76
L76:
	;
	if v204&int32(_a_F_ExecGrantStmt_oids_6) == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v271 = v262
	goto L72
L78:
	;
	v262 = int32(1)
	v265 = base.I32_extend16_s(v204 + v262)
	v266 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87)+120)))
	if v265 <= v266 {
		v204 = v265
		v208 = v265
		goto L76
	} else {
		goto L88
	}
L79:
	;
	if base.I32_extend16_s(v204) < int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+119)))
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
	v242 = F_SearchSysCache2(m, int32(7), v82, base.I64_extend16_s(base.I64_extend_i32_u(v204)))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L14
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	if v242 == int32(0) {
		goto L27
	} else {
		goto L85
	}
L85:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242)+16))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+22)))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246+v247)+91)))
	F_ReleaseCatCache(m, v242)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L14
	} else {
		goto L86
	}
L86:
	;
	if v249 != 0 {
		goto L78
	} else {
		goto L87
	}
L87:
	;
	v256 = v188 + v208<<(uint(int32(3))%32) + int32(56)
	v257 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
	*(*int64)(unsafe.Add(mBase, uint32(v256))) = v257 | v193
	goto L78
L88:
	;
	goto L77
L89:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+667)))
	if v302 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v325 = F_aclcopy(m, v323)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L14
	} else {
		goto L100
	}
L91:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+119)))
	if v308 == int32(83) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	v317 = F_pg_detoast_datum_copy(m, base.I32_wrap_i64(v300))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L14
	} else {
		goto L98
	}
L94:
	;
	v311 = int32(38)
	goto L96
L95:
	;
	v311 = int32(42)
	goto L96
L96:
	;
	v312 = F_acldefault(m, v311, v295)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L14
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+572)) = int32(0)
	v323 = v312
	v324 = int32(0)
	goto L90
L98:
	;
	v321 = F_aclmembers(m, v317, v30+int32(572))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L14
	} else {
		goto L99
	}
L99:
	;
	v323 = v317
	v324 = v321
	goto L90
L100:
	;
	if v182 != int64(0) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v439 == int32(0) {
		v742 = v271
		goto L23
	} else {
		goto L124
	}
L102:
	;
	v330 = v30 + int32(288)
	v331 = int32(0)
	base.MemoryFill(m, v330, v331, int32(272))
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+608)) = uint16(v331)
	v336 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+600)) = v336
	*(*int64)(unsafe.Add(mBase, uint32(v30)+592)) = v336
	*(*int64)(unsafe.Add(mBase, uint32(v30)+584)) = v336
	*(*int64)(unsafe.Add(mBase, uint32(v30)+576)) = v336
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+272)) = uint16(v331)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+264)) = v336
	*(*int64)(unsafe.Add(mBase, uint32(v30)+256)) = v336
	*(*int64)(unsafe.Add(mBase, uint32(v30)+248)) = v336
	*(*int64)(unsafe.Add(mBase, uint32(v30)+240)) = v336
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_select_best_grantor(m, v354, v182, v323, v295, v30+int32(668), v30+int32(656))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L14
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	F_UnlockTuple(m, v37, v83+int32(4), int32(7))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L14
	} else {
		goto L123
	}
L105:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v362 = *(*int64)(unsafe.Add(mBase, uint32(v30)+656))
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v30)+668))
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+119)))
	if v367 == int32(83) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v370 = int32(38)
	goto L108
L107:
	;
	v370 = int32(42)
	goto L108
L108:
	;
	v373 = int32(0)
	v375 = F_restrict_and_check_grant(m, v361, v362, v363, v182, v81, v364, v370, v87+int32(4), v373, v373)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L14
	} else {
		goto L109
	}
L109:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v30)+668))
	v382 = F_merge_acl_with_grant(m, v323, v377, v378, v379, v380, v375, v381, v295)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L14
	} else {
		goto L110
	}
L110:
	;
	v386 = F_aclmembers(m, v382, v30+int32(648))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L14
	} else {
		goto L111
	}
L111:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v30)+536)) = base.I64_extend_i32_u(v382)
	v390 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+271)) = uint8(v390)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v37)+52))
	v397 = F_heap_modify_tuple(m, v83, v392, v330, v30+int32(576), v30+int32(240))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L14
	} else {
		goto L112
	}
L112:
	;
	F_CatalogTupleUpdate(m, v37, v397+int32(4), v397)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L14
	} else {
		goto L113
	}
L113:
	;
	F_UnlockTuple(m, v37, v83+int32(4), int32(7))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L14
	} else {
		goto L114
	}
L114:
	;
	v409 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[0])))
	if v409 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v30)+572))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v30)+648))
	F_updateAclDependencies(m, int32(1259), v81, int32(0), v295, v324, v424, v386, v425)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L14
	} else {
		goto L121
	}
L116:
	;
	v413 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[1])))
	if v413&int32(1) == int32(0) {
		goto L115
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	F_recordExtensionInitPrivWorker(m, v81, int32(1259), int32(0), v382)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
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
	F_pfree(m, v382)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
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
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	if v442 <= int32(0) {
		v742 = v271
		goto L23
	} else {
		goto L125
	}
L125:
	;
	v453 = int32(0)
	goto L126
L126:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v439)+12))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v475+v453<<(uint(int32(2))%32))))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v479)+4))
	if v480 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v742 = v610
	goto L23
L128:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+119)))
	if base.B2i32(v490&int64(37) == int64(0))|base.B2i32(v495 != int32(83)) == int32(0) {
		goto L134
	} else {
		goto L135
	}
L129:
	;
	v490 = int64(39)
	goto L128
L130:
	;
	goto L131
L131:
	;
	v484 = F_string_to_privilege(m, v480)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L14
	} else {
		goto L132
	}
L132:
	;
	if v484&int64(32728) != int64(0) {
		goto L26
	} else {
		goto L133
	}
L133:
	;
	v490 = v484
	goto L128
L134:
	;
	v503 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L14
	} else {
		goto L137
	}
L135:
	;
	v521 = v490
	goto L136
L136:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v479)+8))
	if v522 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L137:
	;
	if v503 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L14
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v521 = v490 & int64(2)
	goto L136
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+64)) = v87 + int32(4)
	F_errmsg(m, int32(_a_F_ExecGrantStmt_oids_7), v30-int32(-64))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L14
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(2059), int32(_a_F_ExecGrantStmt_oids_2))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L14
	} else {
		goto L143
	}
L143:
	;
	goto L140
L144:
	;
	v610 = int32(1)
	v612 = v453 + v610
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	if v612 < v613 {
		v453 = v612
		goto L126
	} else {
		goto L153
	}
L145:
	;
	v525 = int32(0)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v522)+4))
	if v526 <= v525 {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v531 = v525
	goto L147
L147:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v522)+12))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v556+v531<<(uint(int32(2))%32))))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v560)+4))
	v562 = F_get_attnum(m, v81, v561)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L14
	} else {
		goto L149
	}
L148:
	;
	goto L144
L149:
	;
	if v562 == int32(0) {
		goto L25
	} else {
		goto L150
	}
L150:
	;
	v568 = base.I32_extend16_s(v562 + int32(7))
	if base.B2i32(v568 <= int32(0))|base.B2i32(v185 <= v568) != 0 {
		goto L24
	} else {
		goto L151
	}
L151:
	;
	v575 = v188 + v568<<(uint(int32(3))%32)
	v576 = *(*int64)(unsafe.Add(mBase, uint32(v575)))
	*(*int64)(unsafe.Add(mBase, uint32(v575))) = v576 | v521
	v580 = v531 + int32(1)
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v522)+4))
	if v580 < v581 {
		v531 = v580
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
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v81
	F_errmsg_internal(m, int32(_a_F_ExecGrantStmt_oids_8), v30+int32(16))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L14
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1802), int32(_a_F_ExecGrantStmt_oids_2))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
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
	v636 = m.ExcPending
	if v636 != 0 {
		goto L14
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+144)) = v87 + int32(4)
	F_errmsg(m, int32(_a_F_ExecGrantStmt_oids_9), v30+int32(144))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L14
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1826), int32(_a_F_ExecGrantStmt_oids_2))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
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
	v656 = m.ExcPending
	if v656 != 0 {
		goto L14
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+128)) = int32(_a_F_ExecGrantStmt_oids_10)
	F_errmsg(m, int32(_a_F_ExecGrantStmt_oids_11), v30+int32(128))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L14
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1881), int32(_a_F_ExecGrantStmt_oids_2))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v30)+100)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v30)+96)) = v208
	F_errmsg_internal(m, int32(_a_F_ExecGrantStmt_oids_12), v30+int32(96))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L14
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1609), int32(_a_F_ExecGrantStmt_oids_13))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
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
	v691 = m.ExcPending
	if v691 != 0 {
		goto L14
	} else {
		goto L169
	}
L169:
	;
	v692 = F_privilege_to_string(m, v484)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L14
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+80)) = v692
	F_errmsg(m, int32(_a_F_ExecGrantStmt_oids_14), v30+int32(80))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L14
	} else {
		goto L171
	}
L171:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(2046), int32(_a_F_ExecGrantStmt_oids_2))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
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
	v711 = m.ExcPending
	if v711 != 0 {
		goto L14
	} else {
		goto L174
	}
L174:
	;
	v712 = F_get_rel_name(m, v81)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L14
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+52)) = v712
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = v561
	F_errmsg(m, int32(_a_F_ExecGrantStmt_oids_15), v30+int32(48))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L14
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1566), int32(_a_F_ExecGrantStmt_oids_16))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
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
	v733 = m.ExcPending
	if v733 != 0 {
		goto L14
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1569), int32(_a_F_ExecGrantStmt_oids_16))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
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
	v775 = int32(0)
	v779 = v775
	v783 = v775
	goto L184
L182:
	;
	goto L183
L183:
	;
	F_pfree(m, v325)
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L14
	} else {
		goto L245
	}
L184:
	;
	v807 = *(*int64)(unsafe.Add(mBase, uint32(v188+v779<<(uint(int32(3))%32))))
	if v807 != int64(0) {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	goto L183
L186:
	;
	v812 = int32(0)
	base.MemoryFill(m, v30+int32(288), v812, int32(200))
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+600)) = uint8(v812)
	v817 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+592)) = v817
	*(*int64)(unsafe.Add(mBase, uint32(v30)+584)) = v817
	*(*int64)(unsafe.Add(mBase, uint32(v30)+576)) = v817
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+264)) = uint8(v812)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+256)) = v817
	*(*int64)(unsafe.Add(mBase, uint32(v30)+248)) = v817
	*(*int64)(unsafe.Add(mBase, uint32(v30)+240)) = v817
	v831 = int32(7)
	v832 = v783 - v831
	v833 = base.I32_extend16_s(v832)
	v837 = F_SearchSysCache2(m, v831, v82, base.I64_extend16_s(base.I64_extend_i32_u(v832)))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L14
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v1043 = v783 + int32(1)
	v1044 = base.I32_extend16_s(v1043)
	if v1044 < v185 {
		v779 = v1044
		v783 = v1043
		goto L184
	} else {
		goto L244
	}
L189:
	;
	if v837 == int32(0) {
		goto L17
	} else {
		goto L190
	}
L190:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v837)+16))
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v841)+22)))
	v848 = F_SysCacheGetAttr(m, int32(7), v837, int32(22), v30+int32(652))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L14
	} else {
		goto L191
	}
L191:
	;
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+652)))
	if v850 == int32(1) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v868 = m.G0
	v870 = v868 - int32(16)
	m.G0 = v870
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v866)+16))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v325)+16))
	v874 = v872 + v873
	if int32(0) <= v874 {
		goto L200
	} else {
		goto L201
	}
L193:
	;
	v855 = F_acldefault(m, int32(6), v295)
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L14
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v860 = F_pg_detoast_datum_copy(m, base.I32_wrap_i64(v848))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L14
	} else {
		goto L197
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+648)) = int32(0)
	v866 = v855
	v867 = int32(0)
	goto L192
L197:
	;
	v864 = F_aclmembers(m, v860, v30+int32(648))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L14
	} else {
		goto L198
	}
L198:
	;
	v866 = v860
	v867 = v864
	goto L192
L199:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_select_best_grantor(m, v947, v807, v881, v295, v30+int32(668), v30+int32(656))
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L14
	} else {
		goto L219
	}
L200:
	;
	v880 = v874<<(uint(int32(4))%32) + int32(24)
	v881 = F_palloc0(m, v880)
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
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
	v937 = m.ExcPending
	if v937 != 0 {
		goto L14
	} else {
		goto L216
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v881)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v881)+12)) = int32(1033)
	*(*int64)(unsafe.Add(mBase, uint32(v881)+4)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v881))) = v880 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v881)+16)) = v874
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v325)+8))
	if v893 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
	v903 = (v896<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L206
L205:
	;
	v903 = v893
	goto L206
L206:
	;
	v905 = v881 + int32(24)
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v325)+16))
	v908 = v906 << (uint(int32(4)) % 32)
	if v908 != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	base.MemoryCopy(m, v905, v903+v325, v908)
	goto L209
L208:
	;
	goto L209
L209:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v866)+8))
	if v911 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v866)+4))
	v921 = (v914<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L212
L211:
	;
	v921 = v911
	goto L212
L212:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v866)+16))
	v924 = v922 << (uint(int32(4)) % 32)
	if v924 != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v325)+16))
	base.MemoryCopy(m, v905+v925<<(uint(int32(4))%32), v866+v921, v924)
	goto L215
L214:
	;
	goto L215
L215:
	;
	m.G0 = v870 + int32(16)
	goto L199
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v870))) = v874
	F_errmsg_internal(m, int32(_a_F_ExecGrantStmt_oids_18), v870)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L14
	} else {
		goto L217
	}
L217:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_19), int32(446), int32(_a_F_ExecGrantStmt_oids_20))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
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
	F_pfree(m, v881)
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L14
	} else {
		goto L220
	}
L220:
	;
	v956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v957 = *(*int64)(unsafe.Add(mBase, uint32(v30)+656))
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v30)+668))
	v964 = F_restrict_and_check_grant(m, v956, v957, base.B2i32(v807 == int64(39)), v807, v81, v960, int32(6), v87+int32(4), v833, v841+v842+int32(4))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L14
	} else {
		goto L221
	}
L221:
	;
	v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v969 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v30)+668))
	v971 = F_merge_acl_with_grant(m, v866, v966, v967, v968, v969, v964, v970, v295)
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L14
	} else {
		goto L222
	}
L222:
	;
	v975 = F_aclmembers(m, v971, v30+int32(644))
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L14
	} else {
		goto L223
	}
L223:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v971)+16))
	if int32(0) < v977 {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	F_pfree(m, v971)
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L14
	} else {
		goto L242
	}
L225:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v41)+52))
	v996 = F_heap_modify_tuple(m, v837, v989, v30+int32(288), v30+int32(576), v30+int32(240))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L14
	} else {
		goto L230
	}
L226:
	;
	v980 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+261)) = uint8(v980)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+456)) = base.I64_extend_i32_u(v971)
	goto L225
L227:
	;
	goto L228
L228:
	;
	v984 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+597)) = uint8(v984)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+261)) = uint8(v984)
	v988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+652)))
	if v988 != 0 {
		goto L224
	} else {
		goto L229
	}
L229:
	;
	goto L225
L230:
	;
	F_CatalogTupleUpdate(m, v41, v996+int32(4), v996)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L14
	} else {
		goto L231
	}
L231:
	;
	v1003 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[0])))
	if v1003 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v30)+648))
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v30)+644))
	F_updateAclDependencies(m, int32(1259), v81, v833, v295, v867, v1021, v975, v1022)
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L14
	} else {
		goto L241
	}
L233:
	;
	v1007 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[1])))
	if v1007&int32(1) == int32(0) {
		goto L232
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	v1013 = int32(0)
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v971)+16))
	if v1013 < v1014 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	goto L235
L237:
	;
	v1017 = v971
	goto L239
L238:
	;
	v1017 = v1013
	goto L239
L239:
	;
	F_recordExtensionInitPrivWorker(m, v81, int32(1259), v833, v1017)
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
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
	F_ReleaseCatCache(m, v837)
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
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
	F_pfree(m, v188)
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L14
	} else {
		goto L246
	}
L246:
	;
	F_ReleaseCatCache(m, v83)
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L14
	} else {
		goto L247
	}
L247:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L14
	} else {
		goto L248
	}
L248:
	;
	v1082 = v69 + int32(1)
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v1082 < v1083 {
		v69 = v1082
		goto L21
	} else {
		goto L249
	}
L249:
	;
	goto L22
L250:
	;
	F_relation_close(m, v37, int32(3))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L14
	} else {
		goto L251
	}
L251:
	;
	goto L1
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+36)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v833
	F_errmsg_internal(m, int32(_a_F_ExecGrantStmt_oids_12), v30+int32(32))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L14
	} else {
		goto L253
	}
L253:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(1656), int32(_a_F_ExecGrantStmt_oids_21))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
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
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v1143
	F_errmsg_internal(m, int32(_a_F_ExecGrantStmt_oids_22), v30)
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L14
	} else {
		goto L257
	}
L257:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(630), int32(_a_F_ExecGrantStmt_oids_23))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
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
	v1163 = F_table_open(m, int32(_a_F_ExecGrantStmt_oids_24), int32(3))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L14
	} else {
		goto L262
	}
L260:
	;
	v1156 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v1156 != int64(0) {
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
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1165 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L14
	} else {
		goto L323
	}
L264:
	;
	F_relation_close(m, v1163, int32(3))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L14
	} else {
		goto L322
	}
L265:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1165)+4))
	if v1168 <= int32(0) {
		goto L264
	} else {
		goto L266
	}
L266:
	;
	v1176 = v2
	goto L267
L267:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1165)+12))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1199+v1176<<(uint(int32(2))%32))))
	v1205 = F_SearchSysCache1(m, int32(44), base.I64_extend_i32_u(v1203))
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L14
	} else {
		goto L269
	}
L268:
	;
	goto L264
L269:
	;
	if v1205 == int32(0) {
		goto L263
	} else {
		goto L270
	}
L270:
	;
	v1211 = F_SysCacheGetAttrNotNull(m, int32(44), v1205, int32(2))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L14
	} else {
		goto L271
	}
L271:
	;
	v1214 = F_text_to_cstring(m, base.I32_wrap_i64(v1211))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L14
	} else {
		goto L272
	}
L272:
	;
	v1220 = F_SysCacheGetAttr(m, int32(44), v1205, int32(3), v30+int32(572))
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L14
	} else {
		goto L273
	}
L273:
	;
	v1222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+572)))
	if v1222 == int32(1) {
		goto L275
	} else {
		goto L276
	}
L274:
	;
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1242 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	F_select_best_grantor(m, v1241, v1242, v1239, int32(10), v30+int32(240), v30+int32(576))
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L14
	} else {
		goto L281
	}
L275:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1228 = F_acldefault(m, v1226, int32(10))
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L14
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	v1233 = F_pg_detoast_datum_copy(m, base.I32_wrap_i64(v1220))
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L14
	} else {
		goto L279
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+656)) = int32(0)
	v1239 = v1228
	v1240 = int32(0)
	goto L274
L279:
	;
	v1237 = F_aclmembers(m, v1233, v30+int32(656))
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L14
	} else {
		goto L280
	}
L280:
	;
	v1239 = v1233
	v1240 = v1237
	goto L274
L281:
	;
	v1250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v1251 = *(*int64)(unsafe.Add(mBase, uint32(v30)+576))
	v1252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v1253 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v30)+240))
	v1256 = int32(0)
	v1258 = F_restrict_and_check_grant(m, v1250, v1251, v1252, v1253, v1203, v1254, int32(27), v1214, v1256, v1256)
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L14
	} else {
		goto L282
	}
L282:
	;
	v1260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v1261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v30)+240))
	v1266 = F_merge_acl_with_grant(m, v1239, v1260, v1261, v1262, v1263, v1258, v1264, int32(10))
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L14
	} else {
		goto L283
	}
L283:
	;
	v1270 = F_aclmembers(m, v1266, v30+int32(668))
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L14
	} else {
		goto L284
	}
L284:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1274 = F_acldefault(m, v1272, int32(10))
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L14
	} else {
		goto L286
	}
L285:
	;
	v1354 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[0])))
	if v1354 == int32(0) {
		goto L312
	} else {
		goto L313
	}
L286:
	;
	v1276 = int32(0)
	if v1266 != 0 {
		goto L289
	} else {
		goto L290
	}
L287:
	;
	if v1320 != 0 {
		goto L305
	} else {
		goto L306
	}
L288:
	;
	if v1274 == int32(0) {
		v1316 = v1276
		goto L296
	} else {
		goto L297
	}
L289:
	;
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v1266)+16))
	if v1278 != 0 {
		goto L288
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	if v1274 == int32(0) {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	goto L291
L293:
	;
	v1320 = int32(1)
	goto L287
L294:
	;
	goto L295
L295:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+16))
	v1320 = base.B2i32(v1283 == int32(0))
	goto L287
L296:
	;
	v1320 = v1316
	goto L287
L297:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+16))
	if v1278 != v1288 {
		v1316 = v1276
		goto L296
	} else {
		goto L298
	}
L298:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v1266)+8))
	if v1290 != 0 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1298 = v1290
	goto L301
L300:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1266)+4))
	v1298 = (v1291<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L301
L301:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+8))
	if v1300 != 0 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1308 = v1300
	goto L304
L303:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+4))
	v1308 = (v1301<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L304
L304:
	;
	v1312 = F_memcmp(m, v1298+v1266, v1308+v1274, v1278<<(uint(int32(4))%32))
	mBase = m.M
	v1316 = base.B2i32(v1312 == int32(0))
	goto L296
L305:
	;
	F_simple_heap_delete(m, v1163, v1205+int32(4))
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L14
	} else {
		goto L308
	}
L306:
	;
	goto L307
L307:
	;
	v1325 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+296)) = v1325
	*(*int64)(unsafe.Add(mBase, uint32(v30)+288)) = v1325
	v1329 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+650)) = uint8(v1329)
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+648)) = uint16(v1329)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+304)) = base.I64_extend_i32_u(v1266)
	v1335 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+646)) = uint8(v1335)
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+644)) = uint16(v1329)
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1163)+52))
	v1346 = F_heap_modify_tuple(m, v1205, v1339, v30+int32(288), v30+int32(648), v30+int32(644))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L14
	} else {
		goto L309
	}
L308:
	;
	goto L285
L309:
	;
	F_CatalogTupleUpdate(m, v1163, v1346+int32(4), v1346)
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L14
	} else {
		goto L310
	}
L310:
	;
	goto L285
L311:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v30)+656))
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v30)+668))
	F_updateAclDependencies(m, int32(_a_F_ExecGrantStmt_oids_24), v1203, int32(0), int32(10), v1240, v1370, v1270, v1371)
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L14
	} else {
		goto L317
	}
L312:
	;
	v1358 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[1])))
	if v1358&int32(1) == int32(0) {
		goto L311
	} else {
		goto L315
	}
L313:
	;
	goto L314
L314:
	;
	F_recordExtensionInitPrivWorker(m, v1203, int32(_a_F_ExecGrantStmt_oids_24), int32(0), v1266)
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
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
	F_ReleaseCatCache(m, v1205)
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L14
	} else {
		goto L318
	}
L318:
	;
	F_pfree(m, v1266)
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L14
	} else {
		goto L319
	}
L319:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L14
	} else {
		goto L320
	}
L320:
	;
	v1381 = v1176 + int32(1)
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v1165)+4))
	if v1381 < v1382 {
		v1176 = v1381
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
	*(*int32)(unsafe.Add(mBase, uint32(v30)+224)) = v1203
	F_errmsg_internal(m, int32(_a_F_ExecGrantStmt_oids_25), v30+int32(224))
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L14
	} else {
		goto L324
	}
L324:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(2444), int32(_a_F_ExecGrantStmt_oids_26))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
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
	v1449 = F_table_open(m, int32(2995), int32(3))
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L14
	} else {
		goto L331
	}
L329:
	;
	v1442 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v1442 != int64(0) {
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
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1451 == int32(0) {
		goto L333
	} else {
		goto L334
	}
L332:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L14
	} else {
		goto L369
	}
L333:
	;
	F_relation_close(m, v1449, int32(3))
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		goto L14
	} else {
		goto L368
	}
L334:
	;
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1451)+4))
	if v1454 <= int32(0) {
		goto L333
	} else {
		goto L335
	}
L335:
	;
	v1466 = v2
	goto L336
L336:
	;
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v1451)+12))
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1484+v1466<<(uint(int32(2))%32))))
	v1489 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+256)) = v1489
	*(*int64)(unsafe.Add(mBase, uint32(v30)+248)) = v1489
	*(*int64)(unsafe.Add(mBase, uint32(v30)+240)) = v1489
	v1495 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+574)) = uint8(v1495)
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+572)) = uint16(v1495)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+654)) = uint8(v1495)
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+652)) = uint16(v1495)
	v1504 = v30 + int32(576)
	F_ScanKeyInit(m, v1504, int32(1), int32(3), int32(184), base.I64_extend_i32_u(v1488))
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L14
	} else {
		goto L338
	}
L337:
	;
	goto L333
L338:
	;
	v1512 = int32(1)
	v1515 = F_systable_beginscan(m, v1449, int32(2996), v1512, int32(0), v1512, v1504)
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L14
	} else {
		goto L339
	}
L339:
	;
	v1517 = F_systable_getnext(m, v1515)
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L14
	} else {
		goto L340
	}
L340:
	;
	if v1517 == int32(0) {
		goto L332
	} else {
		goto L341
	}
L341:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v1517)+16))
	v1522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1521)+22)))
	v1523 = v1521 + v1522
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(v1523)+4))
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1449)+52))
	v1529 = F_heap_getattr_2(m, v1517, int32(3), v1526, v30+int32(667))
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L14
	} else {
		goto L342
	}
L342:
	;
	v1531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+667)))
	if v1531 == int32(1) {
		goto L344
	} else {
		goto L345
	}
L343:
	;
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1550 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	F_select_best_grantor(m, v1549, v1550, v1547, v1524, v30+int32(668), v30+int32(656))
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L14
	} else {
		goto L350
	}
L344:
	;
	v1536 = F_acldefault(m, int32(22), v1524)
	mBase = m.M
	v1537 = m.ExcPending
	if v1537 != 0 {
		goto L14
	} else {
		goto L347
	}
L345:
	;
	goto L346
L346:
	;
	v1541 = F_pg_detoast_datum_copy(m, base.I32_wrap_i64(v1529))
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L14
	} else {
		goto L348
	}
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+648)) = int32(0)
	v1547 = v1536
	v1548 = int32(0)
	goto L343
L348:
	;
	v1545 = F_aclmembers(m, v1541, v30+int32(648))
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L14
	} else {
		goto L349
	}
L349:
	;
	v1547 = v1541
	v1548 = v1545
	goto L343
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+208)) = v1488
	v1559 = v30 + int32(288)
	v1564 = F_pg_snprintf(m, v1559, int32(64), int32(_a_F_ExecGrantStmt_oids_27), v30+int32(208))
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L14
	} else {
		goto L351
	}
L351:
	;
	v1566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v1567 = *(*int64)(unsafe.Add(mBase, uint32(v30)+656))
	v1568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v1569 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(v30)+668))
	v1572 = int32(0)
	v1574 = F_restrict_and_check_grant(m, v1566, v1567, v1568, v1569, v1488, v1570, int32(22), v1559, v1572, v1572)
	mBase = m.M
	v1575 = m.ExcPending
	if v1575 != 0 {
		goto L14
	} else {
		goto L352
	}
L352:
	;
	v1576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v1577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(v30)+668))
	v1581 = F_merge_acl_with_grant(m, v1547, v1576, v1577, v1578, v1579, v1574, v1580, v1524)
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L14
	} else {
		goto L353
	}
L353:
	;
	v1585 = F_aclmembers(m, v1581, v30+int32(644))
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L14
	} else {
		goto L354
	}
L354:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v30)+256)) = base.I64_extend_i32_u(v1581)
	v1589 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+654)) = uint8(v1589)
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v1449)+52))
	v1598 = F_heap_modify_tuple(m, v1517, v1591, v30+int32(240), v30+int32(572), v30+int32(652))
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L14
	} else {
		goto L355
	}
L355:
	;
	F_CatalogTupleUpdate(m, v1449, v1598+int32(4), v1598)
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L14
	} else {
		goto L356
	}
L356:
	;
	v1605 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[0])))
	if v1605 == int32(0) {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	v1619 = *(*int32)(unsafe.Add(mBase, uint32(v1523)))
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v30)+648))
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v30)+644))
	F_updateAclDependencies(m, int32(2613), v1619, int32(0), v1524, v1548, v1621, v1585, v1622)
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L14
	} else {
		goto L363
	}
L358:
	;
	v1609 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[1])))
	if v1609&int32(1) == int32(0) {
		goto L357
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	F_recordExtensionInitPrivWorker(m, v1488, int32(2613), int32(0), v1581)
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
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
	F_systable_endscan(m, v1515)
	mBase = m.M
	v1626 = m.ExcPending
	if v1626 != 0 {
		goto L14
	} else {
		goto L364
	}
L364:
	;
	F_pfree(m, v1581)
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L14
	} else {
		goto L365
	}
L365:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1630 = m.ExcPending
	if v1630 != 0 {
		goto L14
	} else {
		goto L366
	}
L366:
	;
	v1632 = v1466 + int32(1)
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1451)+4))
	if v1632 < v1633 {
		v1466 = v1632
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
	*(*int32)(unsafe.Add(mBase, uint32(v30)+192)) = v1488
	F_errmsg_internal(m, int32(_a_F_ExecGrantStmt_oids_28), v30+int32(192))
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L14
	} else {
		goto L370
	}
L370:
	;
	F_errfinish(m, int32(_a_F_ExecGrantStmt_oids_1), int32(2304), int32(_a_F_ExecGrantStmt_oids_29))
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
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
	if (base.I32_wrap_i64(int64(base.Ui64(int64(8778778918399))>>(uint(base.I64_extend_i32_u(v1732))%64)))|base.B2i32(base.Ui32(int32(43)) < base.Ui32(v1732)))&int32(1) != 0 {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	v1742 = int32(0)
	v1744 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[2]))
	if v1744 == v1742 {
		goto L381
	} else {
		goto L382
	}
L379:
	;
	goto L380
L380:
	;
	m.G0 = v30 + int32(672)
	return
L381:
	;
	goto L380
L382:
	;
	v1747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1744)+20)))
	if v1747 != 0 {
		goto L381
	} else {
		goto L383
	}
L383:
	;
	v1748 = int32(_a_F_ExecGrantStmt_oids_30)
	v1749 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[3]))
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1744)))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[3])) = v1751
	v1754 = F_palloc(m, int32(48))
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L14
	} else {
		goto L384
	}
L384:
	;
	v1756 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v1754)+40)) = v1756
	v1758 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1754)+32)) = v1758
	v1760 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1754)+24)) = v1760
	v1762 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1754)+16)) = v1762
	v1764 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1754)+8)) = v1764
	v1766 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v1754))) = v1766
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1769 = F_list_copy(m, v1768)
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L14
	} else {
		goto L385
	}
L385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1754)+8)) = v1769
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1773 = F_list_copy(m, v1772)
	mBase = m.M
	v1774 = m.ExcPending
	if v1774 != 0 {
		goto L14
	} else {
		goto L386
	}
L386:
	;
	v1775 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1754)+24)) = v1775
	*(*int32)(unsafe.Add(mBase, uint32(v1754)+28)) = v1773
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1778 == v1775 {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v1854 = F_palloc(m, int32(40))
	mBase = m.M
	v1855 = m.ExcPending
	if v1855 != 0 {
		goto L14
	} else {
		goto L395
	}
L388:
	;
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(v1778)+4))
	if v1781 <= int32(0) {
		goto L387
	} else {
		goto L389
	}
L389:
	;
	v1785 = int32(0)
	v1787 = v1742
	goto L390
L390:
	;
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v1778)+12))
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v1812+v1785<<(uint(int32(2))%32))))
	v1817 = F_copyObjectImpl(m, v1816)
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L14
	} else {
		goto L392
	}
L391:
	;
	goto L387
L392:
	;
	v1819 = F_lappend(m, v1787, v1817)
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L14
	} else {
		goto L393
	}
L393:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1754)+24)) = v1819
	v1823 = v1785 + int32(1)
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1778)+4))
	if v1823 < v1824 {
		v1785 = v1823
		v1787 = v1819
		goto L390
	} else {
		goto L394
	}
L394:
	;
	goto L391
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1854))) = int32(2)
	v1859 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[0])))
	*(*int32)(unsafe.Add(mBase, uint32(v1854)+12)) = v1754
	*(*uint8)(unsafe.Add(mBase, uint32(v1854)+4)) = uint8(v1859)
	*(*int32)(unsafe.Add(mBase, uint32(v1854)+8)) = int32(0)
	v1865 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[2]))
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(v1865)+28))
	v1867 = F_lappend(m, v1866, v1854)
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		goto L14
	} else {
		goto L396
	}
L396:
	;
	v1870 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v1870)+28)) = v1867
	*(*int32)(unsafe.Add(mBase, _c_F_ExecGrantStmt_oids[3])) = v1749
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
	v117 = F_palloc0(m, int32(56))
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
	*(*uint8)(unsafe.Add(mBase, uint32(v138)+32)) = uint8(v122)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+48)) = uint8(v122)
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
						*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(391)
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
							*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(391)
							return v67
						}
					}
				}
			}
		} else {
			v15 = F_MakeSingleTupleTableSlot(m, v8, int32(_a_F_ExecInitJunkFilter_0))
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
						*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(391)
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
							*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(391)
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
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
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
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v89 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v88+v54<<(uint(v89)%32)))) = l4
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v93+v54))) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(l3+l5<<(uint(v89)%32))+24)) = v54
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
	v62 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v62
	v66 = F_palloc_mul(m, int32(4), v62)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
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
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v80 = F_repalloc(m, v77, v58<<(uint(int32(3))%32))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v66
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v71 = F_palloc_mul(m, int32(1), v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v71
	goto L17
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v80
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v85 = F_repalloc(m, v83, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v85
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
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int64
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int64
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
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
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
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
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v287 int64
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v362 int64
	_ = v362
	var v365 int32
	_ = v365
	var v367 int64
	_ = v367
	var v369 int64
	_ = v369
	var v372 int64
	_ = v372
	var v373 int64
	_ = v373
	var v383 int64
	_ = v383
	var v388 int32
	_ = v388
	var v389 int64
	_ = v389
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int64
	_ = v398
	var v408 int64
	_ = v408
	var v416 int32
	_ = v416
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v489 int32
	_ = v489
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int64
	_ = v504
	var v506 int64
	_ = v506
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
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
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v619 int64
	_ = v619
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int64
	_ = v669
	var v671 int64
	_ = v671
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v695 int64
	_ = v695
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v737 int32
	_ = v737
	var v744 int32
	_ = v744
	var v747 int64
	_ = v747
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int64
	_ = v765
	var v769 int64
	_ = v769
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v783 int64
	_ = v783
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v796 int64
	_ = v796
	var v809 int64
	_ = v809
	var v812 int32
	_ = v812
	var v816 int64
	_ = v816
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
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
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
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
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
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
	var v1011 int32
	_ = v1011
	var v1024 int64
	_ = v1024
	var v1030 int32
	_ = v1030
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
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1060 int64
	_ = v1060
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1108 int32
	_ = v1108
	var v1113 int32
	_ = v1113
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1157 int32
	_ = v1157
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
		v1113 = v2
		goto L8
	default:
		goto L13
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L4
	} else {
		goto L270
	}
L8:
	;
	m.G0 = v18 + int32(16)
	return v1113
L9:
	;
	v347 = v344
	goto L100
L10:
	;
	v344 = int32(0)
	goto L9
L11:
	;
	v344 = int32(1)
	goto L9
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L4
	} else {
		goto L97
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L4
	} else {
		goto L94
	}
L14:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)+52))
	if v297 != 0 {
		goto L83
	} else {
		goto L84
	}
L15:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+52))
	if v263 != 0 {
		goto L67
	} else {
		goto L68
	}
L16:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v250
	if v250 == int32(0) {
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
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116)+4)))
	v238 = v236 & int32(_a_F_ExecMemoize_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v116)+4)) = uint16(v238)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)) = uint16(v241)
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
	v137 = v2
	v142 = v2
	goto L53
L51:
	;
	v191 = v2
	goto L52
L52:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v204+v191<<(uint(int32(2))%32))))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v116)+20))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v208)+24))
	v212 = m.T0[v211].(func(*base.Module, int32, int32, int32) int64)(m, v208, v124, v209+v191)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L59
	}
L53:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v150+v137<<(uint(int32(2))%32))))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v116)+20))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v154)+24))
	v158 = m.T0[v157].(func(*base.Module, int32, int32, int32) int64)(m, v154, v124, v155+v137)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
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
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v160+v137<<(uint(int32(3))%32)))) = v158
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v167 = v137 | int32(1)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v165+v167<<(uint(int32(2))%32))))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v116)+20))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v171)+24))
	v175 = m.T0[v174].(func(*base.Module, int32, int32, int32) int64)(m, v171, v124, v172+v167)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v177+v167<<(uint(int32(3))%32)))) = v175
	v182 = int32(2)
	v183 = v137 + v182
	v185 = v142 + v182
	if v185 != v115&int32(-2) {
		v137 = v183
		v142 = v185
		goto L53
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	v191 = v183
	goto L52
L59:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v214+v191<<(uint(int32(3))%32)))) = v212
	goto L48
L60:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v244 = F_MemoizeHash_hash(m, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v243)+16))
	if base.Ui32(v247) <= base.Ui32(v246) {
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
	v1113 = v2
	goto L8
L64:
	;
	goto L65
L65:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v259 = F_ExecStoreMinimalTuple(m, v256, v257, int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	v1113 = v257
	goto L8
L67:
	;
	F_ExecReScan(m, v262)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L4
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v262)+12))
	v267 = m.T0[v266].(func(*base.Module, int32) int32)(m, v262)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L4
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+13)))
	if v278 == int32(1) {
		goto L12
	} else {
		goto L77
	}
L72:
	;
	if v267 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267)+4)))
	if v269&int32(2) == int32(0) {
		goto L71
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v274 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v261)+13)) = uint8(v274)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(5)
	v1113 = v2
	goto L8
L76:
	;
	goto L75
L77:
	;
	v281 = F_cache_store_tuple(m, l0, v267)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	if v281 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(4)
	v287 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+224)) = v287 + int64(1)
	goto L81
L80:
	;
	goto L81
L81:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+8))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)+32))
	m.T0[v293].(func(*base.Module, int32, int32))(m, v291, v267)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v1113 = v291
	goto L8
L83:
	;
	F_ExecReScan(m, v296)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L4
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v296)+12))
	v301 = m.T0[v300].(func(*base.Module, int32) int32)(m, v296)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L4
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+8))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+32))
	m.T0[v312].(func(*base.Module, int32, int32))(m, v310, v301)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L4
	} else {
		goto L93
	}
L88:
	;
	if v301 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+4)))
	if v303&int32(2) == int32(0) {
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
	v1113 = v2
	goto L8
L92:
	;
	goto L91
L93:
	;
	v1113 = v310
	goto L8
L94:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v319
	F_errmsg_internal(m, int32(_a_F_ExecMemoize_2), v18)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_ExecMemoize_3), int32(947), int32(_a_F_ExecMemoize_4))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
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
	v336 = m.ExcPending
	if v336 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_ExecMemoize_3), int32(894), int32(_a_F_ExecMemoize_4))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
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
	if v347 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L102:
	;
	v1108 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v243)+16)) = v1108
	v347 = v1108
	goto L100
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(5)
	v1113 = int32(0)
	goto L8
L104:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+52))
	if v1044 != 0 {
		goto L249
	} else {
		goto L250
	}
L105:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	v848 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v243)+8)) = v847 + v848
	*(*uint8)(unsafe.Add(mBase, uint32(v834)+12)) = uint8(v848)
	*(*int32)(unsafe.Add(mBase, uint32(v834)+8)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v834))) = int32(0)
	v856 = int32(_a_F_ExecMemoize_0)
	v857 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMemoize[1]))
	v859 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMemoize[1])) = v859
	v862 = F_palloc(m, int32(12))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L4
	} else {
		goto L212
	}
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L4
	} else {
		goto L209
	}
L107:
	;
	v362 = *(*int64)(unsafe.Add(mBase, uint32(v243)))
	if v362 == int64(4294967296) {
		goto L106
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v243)+20))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v243)+12))
	v553 = v552 & v244
	v556 = v551 + v553<<(uint(int32(4))%32)
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+12)))
	if v557 != 0 {
		goto L152
	} else {
		goto L153
	}
L110:
	;
	v365 = int32(0)
	v367 = int64(2)
	v369 = v362 << (uint(int64(1)) % 64)
	if base.Ui64(v369) <= base.Ui64(v367) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v347 = int32(1)
	goto L100
L112:
	;
	v372 = v367
	goto L114
L113:
	;
	v372 = v369
	goto L114
L114:
	;
	v373 = int64(1)
	if v372&(v372-v373) == int64(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v383 = v372
	goto L117
L116:
	;
	v383 = v373 << (uint(int64(64)-base.I64_clz(v372)) % 64)
	goto L117
L117:
	;
	if base.Ui64(v383<<(uint(int64(4))%64)) < base.Ui64(int64(2147483647)) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v243)+20))
	v389 = *(*int64)(unsafe.Add(mBase, uint32(v243)))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v243)+24))
	v395 = F_MemoryContextAllocExtended(m, v390, base.I32_wrap_i64(v383)<<(uint(int32(4))%32), int32(5))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v243)+20)) = v395
	v398 = int64(1)
	if v383&(v383-v398) == int64(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v408 = v383
	goto L124
L123:
	;
	v408 = v398 << (uint(int64(64)-base.I64_clz(v383)) % 64)
	goto L124
L124:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v408<<(uint(int64(4))%64)) {
		goto L7
	} else {
		goto L125
	}
L125:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v243))) = v408
	v416 = base.I32_wrap_i64(v408) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v243)+12)) = v416
	if v408 == int64(4294967296) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v425 = int32(-85899346)
	goto L128
L127:
	;
	v425 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v408), float64(0.9)))
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243)+16)) = v425
	if v389 != int64(0) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v432 = v365
	goto L133
L130:
	;
	goto L131
L131:
	;
	F_pfree(m, v388)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L4
	} else {
		goto L150
	}
L132:
	;
	v461 = v365
	v462 = v458
	goto L138
L133:
	;
	v446 = v388 + v432<<(uint(int32(4))%32)
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446)+12)))
	if v447 != int32(1) {
		v458 = v432
		goto L132
	} else {
		goto L135
	}
L134:
	;
	v458 = int32(0)
	goto L132
L135:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v446)+8))
	if v450&v416 == v432 {
		v458 = v432
		goto L132
	} else {
		goto L136
	}
L136:
	;
	v454 = v432 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v454)) < base.Ui64(v389) {
		v432 = v454
		goto L133
	} else {
		goto L137
	}
L137:
	;
	goto L134
L138:
	;
	v476 = v388 + v462<<(uint(int32(4))%32)
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476)+12)))
	if v477 == int32(1) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	goto L131
L140:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v243)+12))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v476)+8))
	v489 = v481
	goto L143
L141:
	;
	goto L142
L142:
	;
	v524 = v462 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v524)) < base.Ui64(v389) {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	v497 = v489 & v480
	v502 = v395 + v497<<(uint(int32(4))%32)
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502)+12)))
	if v503 != 0 {
		v489 = v497 + int32(1)
		goto L143
	} else {
		goto L145
	}
L144:
	;
	v504 = *(*int64)(unsafe.Add(mBase, uint32(v476)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v502)+8)) = v504
	v506 = *(*int64)(unsafe.Add(mBase, uint32(v476)))
	*(*int64)(unsafe.Add(mBase, uint32(v502))) = v506
	goto L142
L145:
	;
	goto L144
L146:
	;
	v528 = v524
	goto L148
L147:
	;
	v528 = int32(0)
	goto L148
L148:
	;
	v530 = v461 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v530)) < base.Ui64(v389) {
		v461 = v530
		v462 = v528
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
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v561)))
	v721 = v719 + int32(4)
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v721 != v722 {
		goto L189
	} else {
		goto L190
	}
L152:
	;
	v561 = v556
	v562 = v553
	v563 = int32(0)
	v566 = v552
	goto L155
L153:
	;
	v706 = v556
	goto L154
L154:
	;
	v834 = v706
	goto L105
L155:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v561)+8))
	if v574 == v244 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v706 = v702
	goto L154
L157:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v561)))
	v577 = F_MemoizeHash_equal(m, v243, v576)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L4
	} else {
		goto L160
	}
L158:
	;
	v581 = v574
	v582 = v566
	goto L159
L159:
	;
	v583 = v581 & v582
	if base.Ui32(v562) < base.Ui32(v583) {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	if v577 != 0 {
		goto L151
	} else {
		goto L161
	}
L161:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v243)+12))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v561)+8))
	v581 = v580
	v582 = v579
	goto L159
L162:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	v587 = v562 + v585
	goto L164
L163:
	;
	v587 = v562
	goto L164
L164:
	;
	v590 = v582 & (v562 + int32(1))
	if base.Ui32(v587-v583) < base.Ui32(v563) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v595 = v551 + v590<<(uint(int32(4))%32)
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595)+12)))
	if v596 != 0 {
		goto L168
	} else {
		goto L169
	}
L166:
	;
	goto L167
L167:
	;
	v690 = v563 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v690) {
		goto L184
	} else {
		goto L185
	}
L168:
	;
	v599 = v590
	v603 = int32(0)
	goto L171
L169:
	;
	v632 = v590
	v635 = v595
	goto L170
L170:
	;
	if v632 != v562 {
		goto L178
	} else {
		goto L179
	}
L171:
	;
	v614 = v603 + int32(1)
	if int32(151) <= v614 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v632 = v626
	v635 = v629
	goto L170
L173:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	v619 = *(*int64)(unsafe.Add(mBase, uint32(v243)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v617), base.F64_convert_i64_u(v619)), float64(0.1)) != 0 {
		goto L102
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v626 = (v599 + int32(1)) & v582
	v629 = v551 + v626<<(uint(int32(4))%32)
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629)+12)))
	if v630 != 0 {
		v599 = v626
		v603 = v614
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
	v648 = v632
	v651 = v635
	goto L181
L179:
	;
	goto L180
L180:
	;
	v834 = v561
	goto L105
L181:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v243)+12))
	v665 = v662 & (v648 - int32(1))
	v668 = v551 + v665<<(uint(int32(4))%32)
	v669 = *(*int64)(unsafe.Add(mBase, uint32(v668)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v651)+8)) = v669
	v671 = *(*int64)(unsafe.Add(mBase, uint32(v668)))
	*(*int64)(unsafe.Add(mBase, uint32(v651))) = v671
	if v665 != v562 {
		v648 = v665
		v651 = v668
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
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	v695 = *(*int64)(unsafe.Add(mBase, uint32(v243)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v693), base.F64_convert_i64_u(v695)), float64(0.1)) != 0 {
		goto L102
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v702 = v551 + v590<<(uint(int32(4))%32)
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v702)+12)))
	if v703 != 0 {
		v561 = v702
		v562 = v590
		v563 = v690
		v566 = v582
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
	v725 = l0 + int32(180)
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v719)+4))
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v719)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v726)+4)) = v727
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v719)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v727))) = v729
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if v731 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L190:
	;
	goto L191
L191:
	;
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561)+13)))
	if v744 == int32(1) {
		goto L196
	} else {
		goto L197
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v725
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v725
	goto L194
L193:
	;
	goto L194
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v719)+8)) = v725
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v725)))
	*(*int32)(unsafe.Add(mBase, uint32(v719)+4)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v737)+4)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v725))) = v721
	goto L191
L195:
	;
	goto L103
L196:
	;
	v747 = *(*int64)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+200)) = v747 + int64(1)
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v561)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v561
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v751
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v561)+4))
	if v754 == int32(0) {
		goto L195
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	v765 = *(*int64)(unsafe.Add(mBase, uint32(l0)+208))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v765 + int64(1)
	v769 = int64(0)
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v561)+4))
	if v770 != 0 {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(2)
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v561)+4))
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v759)))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v763 = F_ExecStoreMinimalTuple(m, v760, v761, int32(0))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	v1113 = v761
	goto L8
L201:
	;
	v772 = v770
	v783 = v769
	goto L204
L202:
	;
	v809 = v769
	goto L203
L203:
	;
	v812 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v561)+4)) = v812
	*(*uint8)(unsafe.Add(mBase, uint32(v561)+13)) = uint8(v812)
	v816 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v816 - v809
	v1030 = v561
	goto L104
L204:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v772)+4))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v772)))
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v787)))
	F_pfree(m, v787)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L4
	} else {
		goto L206
	}
L205:
	;
	v809 = v796
	goto L203
L206:
	;
	F_pfree(m, v772)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L4
	} else {
		goto L207
	}
L207:
	;
	v796 = v783 + base.I64_extend_i32_u(v788+int32(8))
	if v786 != 0 {
		v772 = v786
		v783 = v796
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
	v826 = m.ExcPending
	if v826 != 0 {
		goto L4
	} else {
		goto L210
	}
L210:
	;
	F_errfinish(m, int32(_a_F_ExecMemoize_7), int32(635), int32(_a_F_ExecMemoize_8))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v834))) = v862
	v865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v865)+8))
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v867)+48))
	v869 = m.T0[v868].(func(*base.Module, int32, int32) int32)(m, v865, int32(0))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L4
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v862))) = v869
	v872 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v834)))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v873)))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v874)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v872 + base.I64_extend_i32_u(v875+int32(28))
	v881 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v834)+4)) = v881
	*(*uint8)(unsafe.Add(mBase, uint32(v834)+13)) = uint8(v881)
	v886 = l0 + int32(180)
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v834)))
	v889 = v887 + int32(4)
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if v890 == v881 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v886
	goto L216
L215:
	;
	goto L216
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v887)+8)) = v886
	v896 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	*(*int32)(unsafe.Add(mBase, uint32(v887)+4)) = v896
	*(*int32)(unsafe.Add(mBase, uint32(v896)+4)) = v889
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v889
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMemoize[1])) = v857
	v904 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v905 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	if base.Ui64(v904) <= base.Ui64(v905) {
		v1011 = v834
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v1024 = *(*int64)(unsafe.Add(mBase, uint32(l0)+208))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v1024 + int64(1)
	v1030 = v1011
	goto L104
L218:
	;
	v907 = F_cache_reduce_memory(m, l0, v862)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L4
	} else {
		goto L220
	}
L219:
	;
	v1011 = int32(0)
	goto L217
L220:
	;
	if v907 == int32(0) {
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v834)+12)))
	if v911 == int32(1) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v834)))
	if v914 == v862 {
		v1011 = v834
		goto L217
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
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
		goto L226
	}
L225:
	;
	goto L224
L226:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v862)))
	v925 = F_ExecStoreMinimalTuple(m, v923, v917, int32(0))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L4
	} else {
		goto L227
	}
L227:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v917)+12))
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v927)))
	v929 = int32(*(*int16)(unsafe.Add(mBase, uint32(v917)+6)))
	if v929 < v928 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v917)+8))
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v931)+16))
	m.T0[v932].(func(*base.Module, int32, int32))(m, v917, v928)
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L4
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v936 = v916 << (uint(int32(3)) % 32)
	if v936 != 0 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	goto L230
L232:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v918)+16))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v917)+16))
	base.MemoryCopy(m, v937, v938, v936)
	goto L234
L233:
	;
	goto L234
L234:
	;
	if v916 != 0 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v918)+20))
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v917)+20))
	base.MemoryCopy(m, v940, v941, v916)
	goto L237
L236:
	;
	goto L237
L237:
	;
	v943 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v918)+4)))
	v945 = v943 & int32(_a_F_ExecMemoize_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v918)+4)) = uint16(v945)
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v918)+12))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v947)))
	*(*uint16)(unsafe.Add(mBase, uint32(v918)+6)) = uint16(v948)
	goto L238
L238:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v951 = F_MemoizeHash_hash(m, v950)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L4
	} else {
		goto L239
	}
L239:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v950)+20))
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v950)+12))
	v955 = v951 & v954
	v958 = v953 + v955<<(uint(int32(4))%32)
	v959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958)+12)))
	if v959 == int32(0) {
		goto L219
	} else {
		goto L240
	}
L240:
	;
	v963 = v955
	v964 = v958
	v967 = v953
	v969 = v954
	goto L241
L241:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v964)+8))
	if v977 == v951 {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	goto L219
L243:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v964)))
	v980 = F_MemoizeHash_equal(m, v950, v979)
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L4
	} else {
		goto L246
	}
L244:
	;
	v984 = v967
	v985 = v969
	goto L245
L245:
	;
	v988 = v985 & (v963 + int32(1))
	v991 = v984 + v988<<(uint(int32(4))%32)
	v992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v991)+12)))
	if v992 != 0 {
		v963 = v988
		v964 = v991
		v967 = v984
		v969 = v985
		goto L241
	} else {
		goto L248
	}
L246:
	;
	if v980 != 0 {
		v1011 = v964
		goto L217
	} else {
		goto L247
	}
L247:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v950)+12))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v950)+20))
	v984 = v983
	v985 = v982
	goto L245
L248:
	;
	goto L242
L249:
	;
	F_ExecReScan(m, v1043)
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L4
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+12))
	v1048 = m.T0[v1047].(func(*base.Module, int32) int32)(m, v1043)
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L4
	} else {
		goto L254
	}
L252:
	;
	goto L251
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v1030
	if v1030 != 0 {
		goto L264
	} else {
		goto L265
	}
L254:
	;
	if v1048 != 0 {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1048)+4)))
	if v1050&int32(2) == int32(0) {
		goto L253
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	if v1030 != 0 {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	goto L257
L259:
	;
	v1055 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1030)+13)) = uint8(v1055)
	goto L261
L260:
	;
	goto L261
L261:
	;
	goto L103
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v1068
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+8))
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1071)+32))
	m.T0[v1072].(func(*base.Module, int32, int32))(m, v1070, v1048)
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L4
	} else {
		goto L269
	}
L263:
	;
	v1065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1030)+13)) = uint8(v1065)
	v1068 = int32(3)
	goto L262
L264:
	;
	v1058 = F_cache_store_tuple(m, l0, v1048)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L4
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v1060 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+224)) = v1060 + int64(1)
	v1068 = int32(4)
	goto L262
L267:
	;
	if v1058 != 0 {
		goto L263
	} else {
		goto L268
	}
L268:
	;
	goto L266
L269:
	;
	v1113 = v1070
	goto L8
L270:
	;
	F_errmsg_internal(m, int32(_a_F_ExecMemoize_9), int32(0))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L4
	} else {
		goto L271
	}
L271:
	;
	F_errfinish(m, int32(_a_F_ExecMemoize_7), int32(332), int32(_a_F_ExecMemoize_10))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v50 int32
	_ = v50
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
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
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
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
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
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v143 int32
	_ = v143
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v198 float64
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int64
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int64
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v248 float64
	_ = v248
	var v252 int32
	_ = v252
	var v255 float64
	_ = v255
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int64
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[0]))
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	F_MemoryContextReset(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
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
	m.G0 = v18 + int32(16)
	return v292
L8:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v50 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+80))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v264)+24))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)+8))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+12))
	m.T0[v268].(func(*base.Module, int32))(m, v266)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L4
	} else {
		goto L68
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
	if v53 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	if v159 != 0 {
		goto L32
	} else {
		goto L33
	}
L13:
	;
	F_ExecReScan(m, v27)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v56 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v58 = m.T0[v57].(func(*base.Module, int32) int32)(m, v27)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	if v58 == int32(0) {
		v292 = v56
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+4)))
	if v62&int32(2) != 0 {
		v292 = v56
		goto L7
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v58
	v66 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+116)) = uint16(v66)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v30)+88))
	if v68 == v66 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_ExecReScan(m, v26)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L31
	}
L21:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v71 <= int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v75 = v56
	goto L23
L23:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90+v75<<(uint(int32(2))%32))))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v98 = v89 + v95*int32(24)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v99)+8)))
	v101 = int32(*(*int16)(unsafe.Add(mBase, uint32(v58)+6)))
	if v101 < v100 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L20
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+16))
	m.T0[v104].(func(*base.Module, int32, int32))(m, v58, v100)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v108 = v100 - int32(1)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+v109))))
	*(*uint8)(unsafe.Add(mBase, uint32(v98)+16)) = uint8(v111)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v113+v108<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v98)+8)) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	v120 = F_bms_add_member(m, v119, v95)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v120
	v124 = v75 + int32(1)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v124 < v125 {
		v75 = v124
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
	F_ExecReScan(m, v26)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v163 = m.T0[v162].(func(*base.Module, int32) int32)(m, v26)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v163
	if v163 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L9
L38:
	;
	if v29 != 0 {
		goto L51
	} else {
		goto L52
	}
L39:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+4)))
	if v166&int32(2) == int32(0) {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v171 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)) = uint8(v171)
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	if v173 != 0 {
		goto L8
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	switch v174 - int32(1) {
	case 0, 4:
		goto L44
	default:
		goto L8
	}
L44:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v177
	if v28 == int32(0) {
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v181 = int32(_a_F_ExecNestLoop_0)
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1]))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1])) = v184
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v189 = m.T0[v188].(func(*base.Module, int32, int32, int32) int64)(m, v28, v31, v18+int32(13))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1])) = v182
	if v189 != int64(0) {
		goto L37
	} else {
		goto L47
	}
L47:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v195 == int32(0) {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	v198 = *(*float64)(unsafe.Add(mBase, uint32(v195)+432))
	*(*float64)(unsafe.Add(mBase, uint32(v195)+432)) = base.F64_add(v198, float64(1))
	goto L8
L49:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	F_MemoryContextReset(m, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L4
	} else {
		goto L67
	}
L50:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v252 == int32(0) {
		goto L49
	} else {
		goto L66
	}
L51:
	;
	v202 = int32(_a_F_ExecNestLoop_0)
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1]))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1])) = v205
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	v210 = m.T0[v209].(func(*base.Module, int32, int32, int32) int64)(m, v29, v31, v18+int32(14))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v217 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)) = uint8(v217)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v219 == int32(5) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1])) = v203
	if v210 == int64(0) {
		goto L50
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v222 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)) = uint8(v222)
	goto L8
L57:
	;
	goto L58
L58:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	if v224 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v227 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)) = uint8(v227)
	goto L61
L60:
	;
	goto L61
L61:
	;
	if v28 == int32(0) {
		goto L37
	} else {
		goto L62
	}
L62:
	;
	v231 = int32(_a_F_ExecNestLoop_0)
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1]))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1])) = v234
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v239 = m.T0[v238].(func(*base.Module, int32, int32, int32) int64)(m, v28, v31, v18+int32(15))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1])) = v232
	if v239 != int64(0) {
		goto L37
	} else {
		goto L64
	}
L64:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v245 == int32(0) {
		goto L49
	} else {
		goto L65
	}
L65:
	;
	v248 = *(*float64)(unsafe.Add(mBase, uint32(v245)+432))
	*(*float64)(unsafe.Add(mBase, uint32(v245)+432)) = base.F64_add(v248, float64(1))
	goto L49
L66:
	;
	v255 = *(*float64)(unsafe.Add(mBase, uint32(v252)+424))
	*(*float64)(unsafe.Add(mBase, uint32(v252)+424)) = base.F64_add(v255, float64(1))
	goto L49
L67:
	;
	goto L8
L68:
	;
	v271 = int32(_a_F_ExecNestLoop_0)
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1]))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v265)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1])) = v274
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v264)+32))
	v280 = m.T0[v279].(func(*base.Module, int32, int32, int32) int64)(m, v264+int32(8), v265, int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecNestLoop[1])) = v272
	v284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v266)+4)))
	v286 = v284 & int32(_a_F_ExecNestLoop_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v266)+4)) = uint16(v286)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v266)+12))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	*(*uint16)(unsafe.Add(mBase, uint32(v266)+6)) = uint16(v289)
	v292 = v266
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
	switch v14 - int32(400) {
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
				F_errfinish(m, int32(_a_F_ExecRestrPos_4), int32(406), int32(_a_F_ExecRestrPos_5))
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
								F_errfinish(m, int32(_a_F_ExecRestrPos_7), int32(893), int32(_a_F_ExecRestrPos_8))
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
							F_errfinish(m, int32(_a_F_ExecRestrPos_7), int32(893), int32(_a_F_ExecRestrPos_8))
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
								F_errfinish(m, int32(_a_F_ExecRestrPos_10), int32(516), int32(_a_F_ExecRestrPos_11))
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
							F_errfinish(m, int32(_a_F_ExecRestrPos_10), int32(516), int32(_a_F_ExecRestrPos_11))
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
						F_errfinish(m, int32(_a_F_ExecRestrPos_24), int32(2382), int32(_a_F_ExecRestrPos_25))
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
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7+v9*int32(24))+8))
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
				v26 = F_ExecScan(m, l0, int32(823), int32(824))
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
		v26 = F_ExecScan(m, l0, int32(823), int32(824))
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
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
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
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
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v469 int32
	_ = v469
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int64
	_ = v517
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v562 int64
	_ = v562
	var v563 int32
	_ = v563
	var v567 int64
	_ = v567
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v593 int32
	_ = v593
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(224)
	m.G0 = v16
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+56)) = uint8(v18)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	switch v22 {
	case 0:
		goto L7
	case 1:
		goto L6
	default:
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L16
	} else {
		goto L162
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L16
	} else {
		goto L159
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L16
	} else {
		goto L155
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L16
	} else {
		goto L152
	}
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v395
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+88)) = uint8(v407)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v409
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v411
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v413 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L6:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v240 == int32(0) {
		v395 = v2
		goto L5
	} else {
		goto L62
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v20 - int32(12) {
	case 0, 38:
		goto L8
	default:
		goto L10
	case 15:
		goto L11
	case 26, 30:
		goto L9
	}
L8:
	;
	if v23 == int32(0) {
		v395 = v2
		goto L5
	} else {
		goto L54
	}
L9:
	;
	if v23 == int32(0) {
		v395 = v2
		goto L5
	} else {
		goto L47
	}
L10:
	;
	if v23 == int32(0) {
		v395 = v2
		goto L5
	} else {
		goto L40
	}
L11:
	;
	if v23 == int32(0) {
		v395 = v2
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v28 <= int32(0) {
		v395 = v2
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v35 = v2
	v37 = v2
	goto L14
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v37<<(uint(int32(2))%32))))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v51 = F_ParameterAclLookup(m, v49, int32(1))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v395 = v123
	goto L5
L16:
	;
	return
L17:
	;
	v55 = int32(0)
	if v51|base.B2i32(v18&int32(1) == v55) == v55 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v60 = m.G0
	v62 = v60 - int32(48)
	m.G0 = v62
	*(*int64)(unsafe.Add(mBase, uint32(v62)+32)) = int64(0)
	v66 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v62)+12)) = uint16(v66)
	v71 = F_find_option(m, v49, v66, int32(1), int32(10))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	v116 = v51
	goto L20
L20:
	;
	if v116 != 0 {
		goto L35
	} else {
		goto L36
	}
L21:
	;
	if v71 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v77 = F_assignable_custom_variable_name(m, v49, int32(0), int32(21))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L16
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v79 = F_convert_GUC_name_for_parameter_acl(m, v49)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L16
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	v83 = F_table_open(m, int32(_a_F_ExecuteGrantStmt_0), int32(3))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L16
	} else {
		goto L27
	}
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v83)+52))
	v88 = F_GetNewOidWithIndex(m, v83, int32(_a_F_ExecuteGrantStmt_1), int32(1))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L16
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v62)+16)) = base.I64_extend_i32_u(v88)
	v92 = F_cstring_to_text(m, v79)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L16
	} else {
		goto L29
	}
L29:
	;
	v94 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+14)) = uint8(v94)
	*(*int64)(unsafe.Add(mBase, uint32(v62)+24)) = base.I64_extend_i32_u(v92)
	v102 = F_heap_form_tuple(m, v85, v62+int32(16), v62+int32(12))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L16
	} else {
		goto L30
	}
L30:
	;
	F_CatalogTupleInsert(m, v83, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L16
	} else {
		goto L31
	}
L31:
	;
	F_pfree(m, v102)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L16
	} else {
		goto L32
	}
L32:
	;
	F_relation_close(m, v83, int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L16
	} else {
		goto L33
	}
L33:
	;
	m.G0 = v62 + int32(48)
	F_CommandCounterIncrement(m)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L16
	} else {
		goto L34
	}
L34:
	;
	v116 = v88
	goto L20
L35:
	;
	v121 = F_lappend_oid(m, v35, v116)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L16
	} else {
		goto L38
	}
L36:
	;
	v123 = v35
	goto L37
L37:
	;
	v125 = v37 + int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v125 < v126 {
		v35 = v123
		v37 = v125
		goto L14
	} else {
		goto L39
	}
L38:
	;
	v123 = v121
	goto L37
L39:
	;
	goto L15
L40:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v130 <= int32(0) {
		v395 = v2
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v134 = v2
	v137 = v2
	goto L42
L42:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148+v134<<(uint(int32(2))%32))))
	v153 = int32(0)
	F_get_object_address(m, v16+int32(112), v20, v152, v153, int32(1), v153)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L16
	} else {
		goto L44
	}
L43:
	;
	v395 = v159
	goto L5
L44:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v16)+116))
	v159 = F_lappend_oid(m, v137, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L16
	} else {
		goto L45
	}
L45:
	;
	v162 = v134 + int32(1)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v162 < v163 {
		v134 = v162
		v137 = v159
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v167 <= int32(0) {
		v395 = v2
		goto L5
	} else {
		goto L48
	}
L48:
	;
	v171 = v2
	v174 = v2
	goto L49
L49:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v183+v171<<(uint(int32(2))%32))))
	v189 = int32(0)
	v192 = F_RangeVarGetRelidExtended(m, v187, int32(1), v189, v189, v189)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L16
	} else {
		goto L51
	}
L50:
	;
	v395 = v194
	goto L5
L51:
	;
	v194 = F_lappend_oid(m, v174, v192)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L16
	} else {
		goto L52
	}
L52:
	;
	v197 = v171 + int32(1)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v197 < v198 {
		v171 = v197
		v174 = v194
		goto L49
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v202 <= int32(0) {
		v395 = v2
		goto L5
	} else {
		goto L55
	}
L55:
	;
	v206 = v2
	v209 = v2
	goto L56
L56:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v220+v206<<(uint(int32(2))%32))))
	v225 = F_makeTypeNameFromNameList(m, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L16
	} else {
		goto L58
	}
L57:
	;
	v395 = v234
	goto L5
L58:
	;
	F_get_object_address(m, v16+int32(112), v20, v225, v16+int32(108), int32(1), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L16
	} else {
		goto L59
	}
L59:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v16)+116))
	v234 = F_lappend_oid(m, v209, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L16
	} else {
		goto L60
	}
L60:
	;
	v237 = v206 + int32(1)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v237 < v238 {
		v206 = v237
		v209 = v234
		goto L56
	} else {
		goto L61
	}
L61:
	;
	goto L57
L62:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v243 <= int32(0) {
		v395 = v2
		goto L5
	} else {
		goto L63
	}
L63:
	;
	v249 = v20 - int32(19)
	v254 = v2
	v256 = v2
	goto L64
L64:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v240)+12))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v263+v256<<(uint(int32(2))%32))))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+4))
	v270 = F_LookupExplicitNamespace(m, v268, int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L16
	} else {
		goto L66
	}
L65:
	;
	v395 = v378
	goto L5
L66:
	;
	switch v249 {
	case 0, 10, 16:
		goto L70
	default:
		goto L69
	case 19:
		goto L71
	case 23:
		goto L68
	}
L67:
	;
	v388 = v256 + int32(1)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v388 < v389 {
		v254 = v378
		v256 = v388
		goto L64
	} else {
		goto L103
	}
L68:
	;
	v350 = F_getRelationsInNamespace(m, v270, int32(114))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L16
	} else {
		goto L93
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L16
	} else {
		goto L90
	}
L70:
	;
	v279 = int32(3)
	F_ScanKeyInit(m, v16+int32(112), v279, v279, int32(184), base.I64_extend_i32_u(v270))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L16
	} else {
		goto L74
	}
L71:
	;
	v273 = F_getRelationsInNamespace(m, v270, int32(83))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L16
	} else {
		goto L72
	}
L72:
	;
	v275 = F_list_concat(m, v254, v273)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L16
	} else {
		goto L73
	}
L73:
	;
	v378 = v275
	goto L67
L74:
	;
	switch v249 {
	case 0:
		v288 = int32(70)
		goto L76
	default:
		v296 = int32(1)
		goto L75
	case 10:
		goto L77
	}
L75:
	;
	v299 = F_table_open(m, int32(1255), int32(1))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L16
	} else {
		goto L79
	}
L76:
	;
	F_ScanKeyInit(m, v16+int32(168), int32(10), int32(3), v288, int64(112))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L16
	} else {
		goto L78
	}
L77:
	;
	v288 = int32(61)
	goto L76
L78:
	;
	v296 = int32(2)
	goto L75
L79:
	;
	v303 = F_table_beginscan_catalog(m, v299, v296, v16+int32(112))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L16
	} else {
		goto L80
	}
L80:
	;
	v309 = v254
	goto L81
L81:
	;
	v318 = F_heap_getnext(m, v303)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L16
	} else {
		goto L83
	}
L82:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+188))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+12))
	m.T0[v328].(func(*base.Module, int32))(m, v303)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L16
	} else {
		goto L88
	}
L83:
	;
	if v318 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v318)+16))
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+22)))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v320+v321)))
	v324 = F_lappend_oid(m, v309, v323)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L16
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	goto L82
L87:
	;
	v309 = v324
	goto L81
L88:
	;
	F_relation_close(m, v299, int32(1))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L16
	} else {
		goto L89
	}
L89:
	;
	v378 = v309
	goto L67
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v20
	F_errmsg_internal(m, int32(_a_F_ExecuteGrantStmt_2), v16+int32(48))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L16
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_ExecuteGrantStmt_3), int32(851), int32(_a_F_ExecuteGrantStmt_4))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L16
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	v352 = F_list_concat(m, v254, v350)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L16
	} else {
		goto L94
	}
L94:
	;
	v355 = F_getRelationsInNamespace(m, v270, int32(118))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L16
	} else {
		goto L95
	}
L95:
	;
	v357 = F_list_concat(m, v352, v355)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L16
	} else {
		goto L96
	}
L96:
	;
	v360 = F_getRelationsInNamespace(m, v270, int32(109))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L16
	} else {
		goto L97
	}
L97:
	;
	v362 = F_list_concat(m, v357, v360)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L16
	} else {
		goto L98
	}
L98:
	;
	v365 = F_getRelationsInNamespace(m, v270, int32(102))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L16
	} else {
		goto L99
	}
L99:
	;
	v367 = F_list_concat(m, v362, v365)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L16
	} else {
		goto L100
	}
L100:
	;
	v370 = F_getRelationsInNamespace(m, v270, int32(112))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L16
	} else {
		goto L101
	}
L101:
	;
	v372 = F_list_concat(m, v367, v370)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L16
	} else {
		goto L102
	}
L102:
	;
	v378 = v372
	goto L67
L103:
	;
	goto L65
L104:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	switch v469 - int32(9) {
	case 0:
		goto L130
	default:
		goto L117
	case 3:
		goto L129
	case 7:
		goto L120
	case 8:
		goto L119
	case 10:
		goto L128
	case 12:
		goto L127
	case 13:
		goto L126
	case 18:
		goto L118
	case 20:
		goto L124
	case 26:
		goto L123
	case 28:
		goto L125
	case 29:
		goto L116
	case 33:
		v516 = int32(_a_F_ExecuteGrantStmt_5)
		v517 = int64(-16768)
		goto L115
	case 34:
		goto L122
	case 41:
		goto L121
	}
L105:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v413)+4))
	if v416 <= int32(0) {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v419 = int32(0)
	v425 = v419
	v427 = v419
	goto L107
L107:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v413)+12))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v434+v425<<(uint(int32(2))%32))))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	if v439 != int32(4) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	goto L104
L109:
	;
	v443 = F_get_rolespec_oid(m, v438, int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L16
	} else {
		goto L112
	}
L110:
	;
	v446 = int32(0)
	goto L111
L111:
	;
	v447 = F_lappend_oid(m, v427, v446)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L16
	} else {
		goto L113
	}
L112:
	;
	v446 = v443
	goto L111
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = v447
	v451 = v425 + int32(1)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v413)+4))
	if v451 < v452 {
		v425 = v451
		v427 = v447
		goto L107
	} else {
		goto L114
	}
L114:
	;
	goto L108
L115:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v518 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L116:
	;
	v516 = int32(_a_F_ExecuteGrantStmt_6)
	v517 = int64(-263)
	goto L115
L117:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L16
	} else {
		goto L131
	}
L118:
	;
	v516 = int32(_a_F_ExecuteGrantStmt_7)
	v517 = int64(-12289)
	goto L115
L119:
	;
	v516 = int32(_a_F_ExecuteGrantStmt_8)
	v517 = int64(-257)
	goto L115
L120:
	;
	v516 = int32(_a_F_ExecuteGrantStmt_9)
	v517 = int64(-257)
	goto L115
L121:
	;
	v516 = int32(_a_F_ExecuteGrantStmt_10)
	v517 = int64(-257)
	goto L115
L122:
	;
	v516 = int32(_a_F_ExecuteGrantStmt_11)
	v517 = int64(-513)
	goto L115
L123:
	;
	v516 = int32(_a_F_ExecuteGrantStmt_12)
	v517 = int64(-129)
	goto L115
L124:
	;
	v516 = int32(_a_F_ExecuteGrantStmt_13)
	v517 = int64(-129)
	goto L115
L125:
	;
	v516 = int32(_a_F_ExecuteGrantStmt_14)
	v517 = int64(-769)
	goto L115
L126:
	;
	v516 = int32(_a_F_ExecuteGrantStmt_15)
	v517 = int64(-7)
	goto L115
L127:
	;
	v516 = int32(_a_F_ExecuteGrantStmt_16)
	v517 = int64(-257)
	goto L115
L128:
	;
	v516 = int32(_a_F_ExecuteGrantStmt_17)
	v517 = int64(-129)
	goto L115
L129:
	;
	v516 = int32(_a_F_ExecuteGrantStmt_18)
	v517 = int64(-257)
	goto L115
L130:
	;
	v516 = int32(_a_F_ExecuteGrantStmt_19)
	v517 = int64(-3585)
	goto L115
L131:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v502
	F_errmsg_internal(m, int32(_a_F_ExecuteGrantStmt_2), v16+int32(16))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L16
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_ExecuteGrantStmt_3), int32(525), int32(_a_F_ExecuteGrantStmt_20))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L16
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
	F_ExecGrantStmt_oids(m, v16+int32(56))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L16
	} else {
		goto L151
	}
L135:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+72)) = int64(0)
	v523 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+68)) = uint8(v523)
	goto L134
L136:
	;
	goto L137
L137:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+72)) = int64(0)
	v527 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+68)) = uint8(v527)
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	if v530 <= v527 {
		goto L134
	} else {
		goto L138
	}
L138:
	;
	v538 = v527
	v540 = int32(0)
	goto L139
L139:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v518)+12))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v547+v538<<(uint(int32(2))%32))))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)+8))
	if v552 != 0 {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	goto L134
L141:
	;
	v574 = v538 + int32(1)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	if v574 < v575 {
		v538 = v574
		v540 = v571
		goto L139
	} else {
		goto L150
	}
L142:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v553 != int32(42) {
		goto L3
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v551)+4))
	if v559 == int32(0) {
		goto L2
	} else {
		goto L147
	}
L145:
	;
	v556 = F_lappend(m, v540, v551)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L16
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v556
	v571 = v556
	goto L141
L147:
	;
	v562 = F_string_to_privilege(m, v559)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L16
	} else {
		goto L148
	}
L148:
	;
	if v562&v517 != int64(0) {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v567 = *(*int64)(unsafe.Add(mBase, uint32(v16)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+72)) = v567 | v562
	v571 = v540
	goto L141
L150:
	;
	goto L140
L151:
	;
	m.G0 = v16 + int32(224)
	return
L152:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v601
	F_errmsg_internal(m, int32(_a_F_ExecuteGrantStmt_21), v16)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L16
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_ExecuteGrantStmt_3), int32(418), int32(_a_F_ExecuteGrantStmt_20))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L16
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L16
	} else {
		goto L156
	}
L156:
	;
	F_errmsg(m, int32(_a_F_ExecuteGrantStmt_22), int32(0))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L16
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_ExecuteGrantStmt_3), int32(560), int32(_a_F_ExecuteGrantStmt_20))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L16
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
	F_errmsg_internal(m, int32(_a_F_ExecuteGrantStmt_23), int32(0))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L16
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_ExecuteGrantStmt_3), int32(566), int32(_a_F_ExecuteGrantStmt_20))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L16
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L16
	} else {
		goto L163
	}
L163:
	;
	v647 = F_privilege_to_string(m, v562)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L16
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v647
	F_errmsg(m, v516, v16+int32(32))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L16
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_ExecuteGrantStmt_3), int32(572), int32(_a_F_ExecuteGrantStmt_20))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L16
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExtendBufferedRelCommon(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int64
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int64
	_ = v180
	var v188 int32
	_ = v188
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
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int64
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int64
	_ = v298
	var v301 int64
	_ = v301
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int64
	_ = v323
	var v326 int64
	_ = v326
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int64
	_ = v364
	var v367 int64
	_ = v367
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v380 int64
	_ = v380
	var v385 int32
	_ = v385
	var v386 int64
	_ = v386
	var v388 int32
	_ = v388
	var v389 int64
	_ = v389
	var v390 int32
	_ = v390
	var v392 int64
	_ = v392
	var v394 int64
	_ = v394
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v411 int64
	_ = v411
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v422 int64
	_ = v422
	var v425 int64
	_ = v425
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v464 int64
	_ = v464
	var v465 int64
	_ = v465
	var v469 int64
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int64
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
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
	var v509 int64
	_ = v509
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v521 int64
	_ = v521
	var v522 int64
	_ = v522
	var v526 int64
	_ = v526
	var v533 int32
	_ = v533
	var v535 int64
	_ = v535
	var v537 int64
	_ = v537
	var v540 int32
	_ = v540
	var v542 int64
	_ = v542
	var v545 int32
	_ = v545
	var v547 int64
	_ = v547
	var v576 int32
	_ = v576
	var v577 int64
	_ = v577
	var v581 int32
	_ = v581
	var v588 int32
	_ = v588
	var v593 int64
	_ = v593
	var v597 int32
	_ = v597
	var v613 int32
	_ = v613
	var v614 int64
	_ = v614
	var v618 int64
	_ = v618
	var v623 int32
	_ = v623
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int64
	_ = v675
	var v677 int32
	_ = v677
	var v678 int64
	_ = v678
	var v679 int64
	_ = v679
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v693 int64
	_ = v693
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v706 int32
	_ = v706
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v739 int64
	_ = v739
	var v742 int64
	_ = v742
	var v771 int32
	_ = v771
	var v773 int64
	_ = v773
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int64
	_ = v792
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v834 int32
	_ = v834
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v922 int32
	_ = v922
	var v930 int32
	_ = v930
	var v944 int32
	_ = v944
	var v957 int32
	_ = v957
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v971 int64
	_ = v971
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1005 int64
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1085 int32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1132 int32
	_ = v1132
	var v1145 int64
	_ = v1145
	var v1150 int64
	_ = v1150
	var v1155 int64
	_ = v1155
	var v1158 int64
	_ = v1158
	var v1172 int32
	_ = v1172
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1205 int64
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1331 int32
	_ = v1331
	var v1336 int32
	_ = v1336
	var v1363 int64
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1372 int32
	_ = v1372
	var v1373 int64
	_ = v1373
	var v1375 int64
	_ = v1375
	var v1400 int64
	_ = v1400
	var v1412 int64
	_ = v1412
	var v1445 int32
	_ = v1445
	var v1446 int64
	_ = v1446
	var v1449 int64
	_ = v1449
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1489 int32
	_ = v1489
	var v1492 int32
	_ = v1492
	var v1499 int32
	_ = v1499
	var v1501 int64
	_ = v1501
	var v1503 int64
	_ = v1503
	var v1528 int64
	_ = v1528
	var v1532 int32
	_ = v1532
	var v1534 int64
	_ = v1534
	var v1536 int64
	_ = v1536
	var v1539 int64
	_ = v1539
	var v1544 int64
	_ = v1544
	var v1566 int64
	_ = v1566
	var v1574 int64
	_ = v1574
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1632 int32
	_ = v1632
	var v1659 int32
	_ = v1659
	var v1662 int32
	_ = v1662
	var v1664 int32
	_ = v1664
	var v1668 int64
	_ = v1668
	var v1669 int64
	_ = v1669
	var v1673 int64
	_ = v1673
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1683 int64
	_ = v1683
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1695 int32
	_ = v1695
	var v1697 int32
	_ = v1697
	var v1699 int32
	_ = v1699
	var v1703 int32
	_ = v1703
	var v1705 int32
	_ = v1705
	var v1707 int32
	_ = v1707
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1718 int64
	_ = v1718
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1730 int64
	_ = v1730
	var v1731 int64
	_ = v1731
	var v1735 int64
	_ = v1735
	var v1742 int32
	_ = v1742
	var v1744 int64
	_ = v1744
	var v1746 int64
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1751 int64
	_ = v1751
	var v1754 int32
	_ = v1754
	var v1756 int64
	_ = v1756
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1786 int64
	_ = v1786
	var v1790 int32
	_ = v1790
	var v1802 int64
	_ = v1802
	var v1806 int32
	_ = v1806
	var v1818 int32
	_ = v1818
	var v1823 int64
	_ = v1823
	var v1827 int64
	_ = v1827
	var v1832 int32
	_ = v1832
	var v1843 int32
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1853 int32
	_ = v1853
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1875 int32
	_ = v1875
	var v1878 int32
	_ = v1878
	var v1903 int32
	_ = v1903
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1931 int32
	_ = v1931
	var v1933 int32
	_ = v1933
	var v1959 int32
	_ = v1959
	var v1961 int64
	_ = v1961
	var v1968 int32
	_ = v1968
	var v1975 int32
	_ = v1975
	var v1996 int32
	_ = v1996
	var v1999 int32
	_ = v1999
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2006 int64
	_ = v2006
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2022 int32
	_ = v2022
	var v2026 int32
	_ = v2026
	var v2028 int32
	_ = v2028
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2036 int32
	_ = v2036
	var v2044 int32
	_ = v2044
	var v2049 int32
	_ = v2049
	v25 = m.G0
	v27 = v25 - int32(224)
	m.G0 = v27
	*(*int32)(unsafe.Add(mBase, uint32(v27)+128)) = l4
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v31 == int32(116) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		goto L17
	} else {
		goto L346
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v1975
	m.G0 = v27 + int32(224)
	return v1968
L3:
	;
	v957 = l3 & int32(1)
	if v957 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L4:
	;
	v883 = int32(0)
	goto L157
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L17
	} else {
		goto L153
	}
L6:
	;
	if v30 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v842 = F_IOContextForStrategy(m, l2)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L17
	} else {
		goto L139
	}
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v43
	v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v27))) = v45
	v49 = m.G0
	v51 = v49 - int32(160)
	m.G0 = v51
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[0]))
	if v54 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+118)))
	if v37 != int32(116) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+24)))
	if v40 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v27)+128))
	v1968 = v177
	v1975 = v840
	goto L2
L14:
	;
	F_InitLocalBuffers(m)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l4) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	return int32(0)
L18:
	;
	goto L16
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[1]))
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[2]))
	v67 = v64 - v66
	if base.Ui32(l4) < base.Ui32(v67) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v71 = l4
	goto L21
L21:
	;
	if v71 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v69 = l4
	goto L24
L23:
	;
	v69 = v67
	goto L24
L24:
	;
	v71 = v69
	goto L21
L25:
	;
	v75 = int32(0)
	goto L28
L26:
	;
	goto L27
L27:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v149 != 0 {
		goto L33
	} else {
		goto L34
	}
L28:
	;
	v100 = F_GetLocalVictimBuffer(m)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L17
	} else {
		goto L30
	}
L29:
	;
	goto L27
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6+v75<<(uint(int32(2))%32)))) = v100
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[3]))
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[4]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v107+(v100^int32(-1))*int32(56))+20))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v104+(int32(-2)-v113)<<(uint(int32(2))%32))))
	base.MemoryFill(m, v118, int32(0), int32(_a_F_ExtendBufferedRelCommon_0))
	v123 = v75 + int32(1)
	if v123 != v71 {
		v75 = v123
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v177 = F_smgrnblocks(m, v176, l1)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L17
	} else {
		goto L42
	}
L33:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
	if v150 != 0 {
		v176 = v150
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v176 = v175
	goto L32
L36:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v149)+20))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+64)) = v152
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v149)))
	*(*int64)(unsafe.Add(mBase, uint32(v51)+56)) = v154
	v158 = F_smgropen(m, v51+int32(56), v151)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L17
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149)+12)) = v158
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158)+72))
	if v162 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
	v176 = v174
	goto L32
L39:
	;
	v170 = v162
	goto L41
L40:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158)+76))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v158)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v158)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v164))) = v166
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v158)+72))
	v170 = v168
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158)+72)) = v170 + int32(1)
	goto L38
L42:
	;
	v180 = base.I64_extend_i32_u(v71)
	if base.Ui64(base.I64_extend_i32_u(v177)+v180) <= base.Ui64(int64(4294967293)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if v71 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L17
	} else {
		goto L123
	}
L46:
	;
	v188 = int32(0)
	goto L49
L47:
	;
	v432 = v149
	goto L48
L48:
	;
	v455 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[5])))
	v458 = m.G0
	v460 = v458 - int32(16)
	m.G0 = v460
	if v455 != 0 {
		goto L83
	} else {
		goto L84
	}
L49:
	;
	v211 = l6 + v188<<(uint(int32(2))%32)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[4]))
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[6]))
	F_ResourceOwnerEnlarge(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L17
	} else {
		goto L51
	}
L50:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v432 = v429
	goto L48
L51:
	;
	v220 = v212 ^ int32(-1)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v223 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v251 = v214 + v220*int32(56)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+72)) = v252
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+76)) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v250)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+88)) = v188 + v177
	*(*int32)(unsafe.Add(mBase, uint32(v51)+84)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v51)+80)) = v256
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[0]))
	v268 = F_hash_search(m, v262, v51+int32(72), int32(1), v51+int32(71))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L17
	} else {
		goto L62
	}
L53:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	if v224 != 0 {
		v250 = v224
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v250 = v249
	goto L52
L56:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v223)+20))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+48)) = v226
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v223)))
	*(*int64)(unsafe.Add(mBase, uint32(v51)+40)) = v228
	v232 = F_smgropen(m, v51+int32(40), v225)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L17
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+12)) = v232
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v232)+72))
	if v236 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v250 = v248
	goto L52
L59:
	;
	v244 = v236
	goto L61
L60:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v232)+76))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v232)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v232)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = v240
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v232)+72))
	v244 = v242
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+72)) = v244 + int32(1)
	goto L58
L62:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+71)))
	if v270 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v422 = int64(0)
	v425 = base.AtomicRmwCmpxchg64(m, v418, int32(0), v422, v422)
	v427 = v188 + int32(1)
	if v427 != v71 {
		v188 = v427
		goto L49
	} else {
		goto L81
	}
L64:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[7]))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v251)+20))
	v277 = int32(-2) - v276
	v280 = v274 + v277<<(uint(int32(2))%32)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v283 = v281 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v280))) = v283
	if v283 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v386 = int64(0)
	v388 = int32(24)
	v389 = base.AtomicRmwCmpxchg64(m, v251, v388, v386, v386)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v51)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v251)+16)) = v390
	v392 = *(*int64)(unsafe.Add(mBase, uint32(v51)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v251)+8)) = v392
	v394 = *(*int64)(unsafe.Add(mBase, uint32(v51)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v251))) = v394
	*(*int64)(unsafe.Add(mBase, uint32(v251)+24)) = v389 | int64(33816576)
	*(*int32)(unsafe.Add(mBase, uint32(v268)+20)) = v220
	v401 = v251 + v388
	v403 = v251 + int32(36)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	goto L78
L67:
	;
	v287 = int32(_a_F_ExtendBufferedRelCommon_1)
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[2])) = v289 - int32(1)
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[4]))
	v297 = v294 + v277*int32(56)
	v298 = int64(0)
	v301 = base.AtomicRmwCmpxchg64(m, v297, int32(24), v298, v298)
	*(*int64)(unsafe.Add(mBase, uint32(v297)+24)) = v301 - int64(1)
	goto L69
L68:
	;
	goto L69
L69:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[6]))
	F_ResourceOwnerForget(m, v307, base.I64_extend_i32_s(v276+int32(1)), int32(_a_F_ExtendBufferedRelCommon_2))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L17
	} else {
		goto L70
	}
L70:
	;
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[4]))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v268)+20))
	v319 = v315 + v316*int32(56)
	v320 = int32(24)
	v321 = v319 + v320
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v319)+20))
	v323 = int64(0)
	v326 = base.AtomicRmwCmpxchg64(m, v319, v320, v323, v323)
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[7]))
	v333 = v328 + (int32(-2)-v322)<<(uint(int32(2))%32)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	if v334 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v337 = int32(_a_F_ExtendBufferedRelCommon_1)
	v339 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[2])) = v339 + int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v321))) = v326 + int64(1)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	v347 = v346
	goto L73
L72:
	;
	v347 = v334
	goto L73
L73:
	;
	v348 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v333))) = v347 + v348
	v352 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[6]))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v319)+20))
	F_ResourceOwnerRemember(m, v352, base.I64_extend_i32_s(v353+v348), int32(_a_F_ExtendBufferedRelCommon_2))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L17
	} else {
		goto L74
	}
L74:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v319)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v211))) = v360 + int32(1)
	v364 = int64(0)
	v367 = base.AtomicRmwCmpxchg64(m, v319, int32(24), v364, v364)
	*(*int64)(unsafe.Add(mBase, uint32(v319)+24)) = v367 & int64(-16777217)
	v372 = v319 + int32(36)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	goto L75
L75:
	;
	if base.B2i32(v373 != int32(-1)) == int32(0) {
		v418 = v321
		goto L63
	} else {
		goto L76
	}
L76:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+152)) = v378
	v380 = *(*int64)(unsafe.Add(mBase, uint32(v372)))
	*(*int64)(unsafe.Add(mBase, uint32(v51)+144)) = v380
	F_pgaio_wref_wait(m, v51+int32(144))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L17
	} else {
		goto L77
	}
L77:
	;
	v418 = v321
	goto L63
L78:
	;
	if base.B2i32(v404 != int32(-1)) == int32(0) {
		v418 = v401
		goto L63
	} else {
		goto L79
	}
L79:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v403)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+152)) = v409
	v411 = *(*int64)(unsafe.Add(mBase, uint32(v403)))
	*(*int64)(unsafe.Add(mBase, uint32(v51)+144)) = v411
	F_pgaio_wref_wait(m, v51+int32(144))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L17
	} else {
		goto L80
	}
L80:
	;
	v418 = v401
	goto L63
L81:
	;
	goto L50
L82:
	;
	if v432 != 0 {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	F___clock_gettime(m, int32(1), v460)
	mBase = m.M
	v464 = int64(*(*int32)(unsafe.Add(mBase, uint32(v460)+8)))
	v465 = *(*int64)(unsafe.Add(mBase, uint32(v460)))
	v469 = v464 + v465*int64(1000000000)
	goto L85
L84:
	;
	v469 = int64(0)
	goto L85
L85:
	;
	m.G0 = v460 + int32(16)
	goto L82
L86:
	;
	v500 = int32(0)
	F_smgrzeroextend(m, v499, l1, v177, v71)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L17
	} else {
		goto L96
	}
L87:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v432)+12))
	if v473 != 0 {
		v499 = v473
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v499 = v498
	goto L86
L90:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v432)+20))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v432)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+32)) = v475
	v477 = *(*int64)(unsafe.Add(mBase, uint32(v432)))
	*(*int64)(unsafe.Add(mBase, uint32(v51)+24)) = v477
	v481 = F_smgropen(m, v51+int32(24), v474)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L17
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432)+12)) = v481
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v481)+72))
	if v485 != 0 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v432)+12))
	v499 = v497
	goto L86
L93:
	;
	v493 = v485
	goto L95
L94:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v481)+76))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v481)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v486)+4)) = v487
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v481)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v487))) = v489
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v481)+72))
	v493 = v491
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v481)+72)) = v493 + int32(1)
	goto L92
L96:
	;
	v503 = int32(1)
	v509 = base.I64_extend_i32_u(v71 << (uint(int32(13)) % 32))
	v513 = m.G0
	v515 = v513 - int32(16)
	m.G0 = v515
	if v469 != int64(0) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	if v71 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L98:
	;
	F___clock_gettime(m, int32(1), v515)
	mBase = m.M
	v521 = int64(*(*int32)(unsafe.Add(mBase, uint32(v515)+8)))
	v522 = *(*int64)(unsafe.Add(mBase, uint32(v515)))
	v526 = v521 + (v522*int64(1000000000) - v469)
	goto L102
L99:
	;
	goto L100
L100:
	;
	v613 = int32(552)
	v614 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[8]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[8])) = v614 + base.I64_extend_i32_u(v503)
	v618 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[9]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[9])) = v618 + v509
	F_pgstat_count_backend_io_op(m, v503, int32(3), int32(5), v503, v509)
	mBase = m.M
	v623 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[10])) = uint8(v623)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[11])) = uint8(v623)
	m.G0 = v515 + int32(16)
	goto L97
L101:
	;
	v576 = int32(552)
	v577 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[12]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[12])) = v577 + v526
	v581 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[13]))
	v588 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v581))|base.B2i32(int32(1)<<(uint(v581)%32)&int32(_a_F_ExtendBufferedRelCommon_3) == v588) == v588 {
		goto L111
	} else {
		goto L112
	}
L102:
	;
	goto L103
L103:
	;
	v533 = int32(_a_F_ExtendBufferedRelCommon_4)
	v535 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[14]))
	v537 = base.I64_div_s(v526, int64(1000))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[14])) = v535 + v537
	switch v503 {
	case 0:
		goto L107
	case 1:
		goto L106
	default:
		goto L101
	}
L106:
	;
	v545 = int32(_a_F_ExtendBufferedRelCommon_5)
	v547 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[15]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[15])) = v547 + v526
	goto L101
L107:
	;
	v540 = int32(_a_F_ExtendBufferedRelCommon_6)
	v542 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[16]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[16])) = v542 + v526
	goto L101
L111:
	;
	v593 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[17]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[17])) = v593 + v526
	v597 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[10])) = uint8(v597)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[18])) = uint8(v597)
	goto L113
L112:
	;
	goto L113
L113:
	;
	goto L100
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27+int32(128)))) = v71
	v771 = int32(_a_F_ExtendBufferedRelCommon_7)
	v773 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[19]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[19])) = v773 + v180
	m.G0 = v51 + int32(160)
	goto L13
L115:
	;
	if v71 != int32(1) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v642 = v500
	v643 = int32(0)
	goto L119
L117:
	;
	v706 = v500
	goto L118
L118:
	;
	v729 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[4]))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l6+v706<<(uint(int32(2))%32))))
	v738 = v729 + (v733^int32(-1))*int32(56)
	v739 = int64(0)
	v742 = base.AtomicRmwCmpxchg64(m, v738, int32(24), v739, v739)
	*(*int64)(unsafe.Add(mBase, uint32(v738)+24)) = v742 | int64(16777216)
	goto L114
L119:
	;
	v664 = int32(_a_F_ExtendBufferedRelCommon_8)
	v665 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[4]))
	v666 = int32(2)
	v668 = l6 + v642<<(uint(v666)%32)
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v668)))
	v670 = int32(-1)
	v672 = int32(56)
	v674 = v665 + (v669^v670)*v672
	v675 = int64(0)
	v677 = int32(24)
	v678 = base.AtomicRmwCmpxchg64(m, v674, v677, v675, v675)
	v679 = int64(16777216)
	*(*int64)(unsafe.Add(mBase, uint32(v674)+24)) = v678 | v679
	v683 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[4]))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v668)+4))
	v689 = v683 + (v684^v670)*v672
	v693 = base.AtomicRmwCmpxchg64(m, v689, v677, v675, v675)
	*(*int64)(unsafe.Add(mBase, uint32(v689)+24)) = v693 | v679
	v698 = v642 + v666
	v700 = v643 + v666
	if v700 != v71&int32(-2) {
		v642 = v698
		v643 = v700
		goto L119
	} else {
		goto L121
	}
L120:
	;
	if v71&int32(1) == int32(0) {
		goto L114
	} else {
		goto L122
	}
L121:
	;
	goto L120
L122:
	;
	v706 = v698
	goto L118
L123:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L17
	} else {
		goto L124
	}
L124:
	;
	if v149 != 0 {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	v822 = v51 + int32(72)
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v819)+4))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v819)))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v819)+8))
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v819)+12))
	F_GetRelationPath(m, v822, v823, v824, v825, v826, l1)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L17
	} else {
		goto L136
	}
L126:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v815)))
	v819 = v817
	goto L125
L127:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
	if v786 != 0 {
		v819 = v786
		goto L125
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v815 = v27 + int32(4)
	goto L126
L130:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v149)+20))
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = v790
	v792 = *(*int64)(unsafe.Add(mBase, uint32(v149)))
	*(*int64)(unsafe.Add(mBase, uint32(v51)+8)) = v792
	v796 = F_smgropen(m, v51+int32(8), v789)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L17
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149)+12)) = v796
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v796)+72))
	if v800 != 0 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v815 = v149 + int32(12)
	goto L126
L133:
	;
	v808 = v800
	goto L135
L134:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v796)+76))
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v796)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v801)+4)) = v802
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v796)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v802))) = v804
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v796)+72))
	v808 = v806
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v796)+72)) = v808 + int32(1)
	goto L132
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v822
	F_errmsg(m, int32(_a_F_ExtendBufferedRelCommon_9), v51)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L17
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_ExtendBufferedRelCommon_10), int32(405), int32(_a_F_ExtendBufferedRelCommon_11))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L17
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l4) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v848 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[20]))
	v850 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[21]))
	v853 = v848 - v850 - int32(8)
	if base.Ui32(v853) <= base.Ui32(v848) {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	goto L142
L142:
	;
	if l4 != 0 {
		v881 = int32(1)
		goto L4
	} else {
		goto L152
	}
L143:
	;
	v856 = v853
	goto L145
L144:
	;
	v856 = int32(0)
	goto L145
L145:
	;
	if base.Ui32(v856) <= base.Ui32(int32(1)) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v859 = int32(1)
	goto L148
L147:
	;
	v859 = v856
	goto L148
L148:
	;
	if base.Ui32(v856) < base.Ui32(l4) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v861 = v859
	goto L151
L150:
	;
	v861 = l4
	goto L151
L151:
	;
	v881 = v861
	goto L4
L152:
	;
	v944 = int32(0)
	goto L3
L153:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L17
	} else {
		goto L154
	}
L154:
	;
	F_errmsg(m, int32(_a_F_ExtendBufferedRelCommon_12), int32(0))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L17
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(_a_F_ExtendBufferedRelCommon_13), int32(2781), int32(_a_F_ExtendBufferedRelCommon_14))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L17
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
	v910 = F_GetVictimBuffer(m, l2, v842)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L17
	} else {
		goto L159
	}
L158:
	;
	v944 = v881
	goto L3
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6+v883<<(uint(int32(2))%32)))) = v910
	v914 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[22]))
	v916 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[23]))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v916+v910*int32(56)-int32(36))))
	base.MemoryFill(m, v914+v922<<(uint(int32(13))%32), int32(0), int32(_a_F_ExtendBufferedRelCommon_0))
	v930 = v883 + int32(1)
	if v930 != v881 {
		v883 = v930
		goto L157
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	F_LockRelationForExtension(m, v30, int32(7))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L17
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	if l3&int32(16) != 0 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	goto L163
L165:
	;
	if v30 == int32(0) {
		v992 = v841
		goto L168
	} else {
		goto L169
	}
L166:
	;
	goto L167
L167:
	;
	if v30 == int32(0) {
		v1026 = v841
		goto L176
	} else {
		goto L177
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v992+l1<<(uint(int32(2))%32))+20)) = int32(-1)
	goto L167
L169:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	if v967 != 0 {
		v992 = v967
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+120)) = v969
	v971 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+112)) = v971
	v975 = F_smgropen(m, v27+int32(112), v968)
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L17
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v975
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v975)+72))
	if v979 != 0 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v992 = v991
	goto L168
L173:
	;
	v987 = v979
	goto L175
L174:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v975)+76))
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v975)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v980)+4)) = v981
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v975)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v981))) = v983
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v975)+72))
	v987 = v985
	goto L175
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v975)+72)) = v987 + int32(1)
	goto L172
L176:
	;
	v1027 = F_smgrnblocks(m, v1026, l1)
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L17
	} else {
		goto L184
	}
L177:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	if v1001 != 0 {
		v1026 = v1001
		goto L176
	} else {
		goto L178
	}
L178:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+104)) = v1003
	v1005 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+96)) = v1005
	v1009 = F_smgropen(m, v27+int32(96), v1002)
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L17
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v1009
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+72))
	if v1013 != 0 {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v1026 = v1025
	goto L176
L181:
	;
	v1021 = v1013
	goto L183
L182:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+76))
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1014)+4)) = v1015
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1015))) = v1017
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+72))
	v1021 = v1019
	goto L183
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1009)+72)) = v1021 + int32(1)
	goto L180
L184:
	;
	if l5 == int32(-1) {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v1145 = base.I64_extend_i32_u(v1132)
	if base.Ui64(int64(4294967293)) < base.Ui64(v1145+base.I64_extend_i32_u(v1027)) {
		goto L1
	} else {
		goto L206
	}
L186:
	;
	v1132 = v944
	goto L185
L187:
	;
	goto L188
L188:
	;
	if base.Ui64(base.I64_extend_i32_u(l5)) < base.Ui64(base.I64_extend_i32_u(v1027)+base.I64_extend_i32_u(v944)) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v1037 = l5 - v1027
	goto L191
L190:
	;
	v1037 = v944
	goto L191
L191:
	;
	if base.Ui32(v1027) <= base.Ui32(l5) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v1040 = v1037
	goto L194
L193:
	;
	v1040 = int32(0)
	goto L194
L194:
	;
	if base.Ui32(v1040) < base.Ui32(v944) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v1042 = v1040
	goto L198
L196:
	;
	goto L197
L197:
	;
	if v1040 != 0 {
		v1132 = v1040
		goto L185
	} else {
		goto L203
	}
L198:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[6]))
	v1069 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[23]))
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(l6+v1042<<(uint(int32(2))%32))))
	v1076 = v1069 + v1073*int32(56)
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1076-int32(36))))
	F_ResourceOwnerForget(m, v1067, base.I64_extend_i32_s(v1079+int32(1)), int32(_a_F_ExtendBufferedRelCommon_2))
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L17
	} else {
		goto L200
	}
L199:
	;
	goto L197
L200:
	;
	F_UnpinBufferNoOwner(m, v1076-int32(56))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L17
	} else {
		goto L201
	}
L201:
	;
	v1091 = v1042 + int32(1)
	if v1091 != v944 {
		v1042 = v1091
		goto L198
	} else {
		goto L202
	}
L202:
	;
	goto L199
L203:
	;
	v1117 = int32(0)
	if v957 != 0 {
		v1968 = v1027
		v1975 = v1117
		goto L2
	} else {
		goto L204
	}
L204:
	;
	F_UnlockRelationForExtension(m, v30, int32(7))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L17
	} else {
		goto L205
	}
L205:
	;
	v1968 = v1027
	v1975 = v1117
	goto L2
L206:
	;
	if v1132 != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v1150 = int64(2181300224)
	if v31 == int32(112) {
		goto L210
	} else {
		goto L211
	}
L208:
	;
	goto L209
L209:
	;
	v1659 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[5])))
	v1662 = m.G0
	v1664 = v1662 - int32(16)
	m.G0 = v1664
	if v1659 != 0 {
		goto L292
	} else {
		goto L293
	}
L210:
	;
	v1155 = v1150
	goto L212
L211:
	;
	v1155 = int64(33816576)
	goto L212
L212:
	;
	if l1 == int32(3) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v1158 = v1150
	goto L215
L214:
	;
	v1158 = v1155
	goto L215
L215:
	;
	v1172 = int32(0)
	goto L216
L216:
	;
	v1186 = l6 + v1172<<(uint(int32(2))%32)
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1186)))
	v1189 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[23]))
	v1191 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[6]))
	F_ResourceOwnerEnlarge(m, v1191)
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L17
	} else {
		goto L218
	}
L217:
	;
	goto L209
L218:
	;
	v1196 = v1189 + v1187*int32(56)
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L17
	} else {
		goto L219
	}
L219:
	;
	if v30 == int32(0) {
		v1226 = v841
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v1228 = v1196 - int32(56)
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v1226)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+132)) = v1229
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1226)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+136)) = v1231
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1226)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+148)) = v1027 + v1172
	*(*int32)(unsafe.Add(mBase, uint32(v27)+144)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v27)+140)) = v1233
	v1239 = v27 + int32(132)
	v1240 = F_BufTableHashCode(m, v1239)
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L17
	} else {
		goto L228
	}
L221:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	if v1201 != 0 {
		v1226 = v1201
		goto L220
	} else {
		goto L222
	}
L222:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+88)) = v1203
	v1205 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+80)) = v1205
	v1209 = F_smgropen(m, v27+int32(80), v1202)
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L17
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v1209
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+72))
	if v1213 != 0 {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v1226 = v1225
	goto L220
L225:
	;
	v1221 = v1213
	goto L227
L226:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+76))
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1214)+4)) = v1215
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1215))) = v1217
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+72))
	v1221 = v1219
	goto L227
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1209)+72)) = v1221 + int32(1)
	goto L224
L228:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[24]))
	v1250 = v1243 + v1240&int32(127)<<(uint(int32(7))%32) + int32(_a_F_ExtendBufferedRelCommon_15)
	v1252 = F_LWLockAcquire(m, v1250, int32(0))
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L17
	} else {
		goto L229
	}
L229:
	;
	v1255 = v1196 - int32(36)
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1255)))
	v1257 = F_BufTableInsert(m, v1239, v1240, v1256)
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L17
	} else {
		goto L231
	}
L230:
	;
	v1632 = v1172 + int32(1)
	if v1632 != v1132 {
		v1172 = v1632
		goto L216
	} else {
		goto L290
	}
L231:
	;
	if int32(0) <= v1257 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[23]))
	v1265 = v1262 + v1257*int32(56)
	v1267 = F_PinBuffer(m, v1265, l2, int32(0))
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L17
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	v1372 = v1196 - int32(32)
	v1373 = int64(4194304)
	v1375 = base.AtomicRmwOr64(m, v1372, int32(0), v1373)
	if v1375&v1373 != int64(0) {
		goto L258
	} else {
		goto L259
	}
L235:
	;
	F_LWLockRelease(m, v1250)
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L17
	} else {
		goto L236
	}
L236:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[6]))
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v1255)))
	F_ResourceOwnerForget(m, v1272, base.I64_extend_i32_s(v1273+int32(1)), int32(_a_F_ExtendBufferedRelCommon_2))
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L17
	} else {
		goto L237
	}
L237:
	;
	F_UnpinBufferNoOwner(m, v1228)
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		goto L17
	} else {
		goto L238
	}
L238:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1186))) = v1282 + int32(1)
	if v1267 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	goto L254
L240:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[22]))
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+20))
	v1294 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1289+v1290<<(uint(int32(13))%32))+14)))
	if v1294 == int32(0) {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L17
	} else {
		goto L242
	}
L242:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+16))
	if v30 != 0 {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	v1321 = v27 + int32(152)
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1318)+12))
	F_GetRelationPath(m, v1321, v1316, v1317, v1319, v1322, l1)
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L17
	} else {
		goto L251
	}
L244:
	;
	v1302 = F_RelationGetSmgr(m, v30)
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L17
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v841)+8))
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v841)))
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v841)+4))
	v1316 = v1315
	v1317 = v1314
	v1318 = v841
	v1319 = v1313
	goto L243
L247:
	;
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1302)+4))
	v1305 = F_RelationGetSmgr(m, v30)
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		goto L17
	} else {
		goto L248
	}
L248:
	;
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1305)))
	v1308 = F_RelationGetSmgr(m, v30)
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L17
	} else {
		goto L249
	}
L249:
	;
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v1308)+8))
	v1311 = F_RelationGetSmgr(m, v30)
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L17
	} else {
		goto L250
	}
L250:
	;
	v1316 = v1304
	v1317 = v1307
	v1318 = v1311
	v1319 = v1310
	goto L243
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v27)+68)) = v1321
	F_errmsg(m, int32(_a_F_ExtendBufferedRelCommon_16), v27-int32(-64))
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L17
	} else {
		goto L252
	}
L252:
	;
	F_errfinish(m, int32(_a_F_ExtendBufferedRelCommon_13), int32(2968), int32(_a_F_ExtendBufferedRelCommon_17))
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L17
	} else {
		goto L253
	}
L253:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L254:
	;
	v1363 = base.AtomicRmwAnd64(m, v1265, int32(24), int64(-16777217))
	v1364 = int32(1)
	v1367 = F_StartSharedBufferIO(m, v1265, v1364, v1364, int32(0))
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L17
	} else {
		goto L256
	}
L255:
	;
	goto L230
L256:
	;
	if v1367 == int32(0) {
		goto L254
	} else {
		goto L257
	}
L257:
	;
	goto L255
L258:
	;
	v1400 = v1375
	goto L261
L259:
	;
	v1528 = v1375
	goto L260
L260:
	;
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v27)+148))
	*(*int32)(unsafe.Add(mBase, uint32(v1228)+16)) = v1532
	v1534 = *(*int64)(unsafe.Add(mBase, uint32(v27)+140))
	*(*int64)(unsafe.Add(mBase, uint32(v1228)+8)) = v1534
	v1536 = *(*int64)(unsafe.Add(mBase, uint32(v27)+132))
	*(*int64)(unsafe.Add(mBase, uint32(v1228))) = v1536
	v1539 = v1528 | int64(4194304)
	v1544 = base.AtomicRmwCmpxchg64(m, v1372, int32(0), v1539, v1528&int64(-38010881)|v1158)
	if v1544 != v1539 {
		goto L282
	} else {
		goto L283
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+172)) = int32(_a_F_ExtendBufferedRelCommon_18)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+168)) = int32(_a_F_ExtendBufferedRelCommon_19)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+164)) = int32(_a_F_ExtendBufferedRelCommon_13)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+160)) = int32(0)
	v1412 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+152)) = v1412
	if v1400&int64(4194304) != v1412 {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	v1528 = v1503
	goto L260
L263:
	;
	goto L266
L264:
	;
	goto L265
L265:
	;
	v1481 = int32(_a_F_ExtendBufferedRelCommon_20)
	v1482 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[25]))
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(152))+8))
	if v1484 == int32(0) {
		goto L273
	} else {
		goto L274
	}
L266:
	;
	F_perform_spin_delay(m, v27+int32(152))
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L17
	} else {
		goto L268
	}
L267:
	;
	goto L265
L268:
	;
	v1446 = int64(0)
	v1449 = base.AtomicRmwCmpxchg64(m, v1372, int32(0), v1446, v1446)
	if v1449&int64(4194304) != v1446 {
		goto L266
	} else {
		goto L269
	}
L269:
	;
	goto L267
L270:
	;
	v1501 = int64(4194304)
	v1503 = base.AtomicRmwOr64(m, v1372, int32(0), v1501)
	if v1503&v1501 != int64(0) {
		v1400 = v1503
		goto L261
	} else {
		goto L281
	}
L271:
	;
	goto L270
L272:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[25])) = v1499
	goto L271
L273:
	;
	if int32(999) < v1482 {
		goto L271
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	if v1482 < int32(11) {
		goto L271
	} else {
		goto L280
	}
L276:
	;
	v1489 = int32(900)
	if v1489 <= v1482 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1492 = v1489
	goto L279
L278:
	;
	v1492 = v1482
	goto L279
L279:
	;
	v1499 = v1492 + int32(100)
	goto L272
L280:
	;
	v1499 = v1482 - int32(1)
	goto L272
L281:
	;
	goto L262
L282:
	;
	v1566 = v1544
	goto L285
L283:
	;
	goto L284
L284:
	;
	F_LWLockRelease(m, v1250)
	mBase = m.M
	v1601 = m.ExcPending
	if v1601 != 0 {
		goto L17
	} else {
		goto L288
	}
L285:
	;
	v1574 = base.AtomicRmwCmpxchg64(m, v1372, int32(0), v1566, v1566&int64(-38010881)|v1158)
	if v1566 != v1574 {
		v1566 = v1574
		goto L285
	} else {
		goto L287
	}
L286:
	;
	goto L284
L287:
	;
	goto L286
L288:
	;
	v1602 = int32(1)
	v1605 = F_StartSharedBufferIO(m, v1228, v1602, v1602, int32(0))
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
		goto L17
	} else {
		goto L289
	}
L289:
	;
	goto L230
L290:
	;
	goto L217
L291:
	;
	if v30 == int32(0) {
		v1705 = v841
		goto L295
	} else {
		goto L296
	}
L292:
	;
	F___clock_gettime(m, int32(1), v1664)
	mBase = m.M
	v1668 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1664)+8)))
	v1669 = *(*int64)(unsafe.Add(mBase, uint32(v1664)))
	v1673 = v1668 + v1669*int64(1000000000)
	goto L294
L293:
	;
	v1673 = int64(0)
	goto L294
L294:
	;
	m.G0 = v1664 + int32(16)
	goto L291
L295:
	;
	F_smgrzeroextend(m, v1705, l1, v1027, v1132)
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L17
	} else {
		goto L303
	}
L296:
	;
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	if v1679 != 0 {
		v1705 = v1679
		goto L295
	} else {
		goto L297
	}
L297:
	;
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = v1681
	v1683 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+48)) = v1683
	v1687 = F_smgropen(m, v27+int32(48), v1680)
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L17
	} else {
		goto L298
	}
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v1687
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1687)+72))
	if v1691 != 0 {
		goto L300
	} else {
		goto L301
	}
L299:
	;
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v1705 = v1703
	goto L295
L300:
	;
	v1699 = v1691
	goto L302
L301:
	;
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v1687)+76))
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v1687)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1692)+4)) = v1693
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v1687)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1693))) = v1695
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v1687)+72))
	v1699 = v1697
	goto L302
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1687)+72)) = v1699 + int32(1)
	goto L299
L303:
	;
	if v957 == int32(0) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	F_UnlockRelationForExtension(m, v30, int32(7))
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L17
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	v1713 = int32(0)
	v1715 = int32(1)
	v1718 = base.I64_extend_i32_u(v1132 << (uint(int32(13)) % 32))
	v1722 = m.G0
	v1724 = v1722 - int32(16)
	m.G0 = v1724
	if v1673 != int64(0) {
		goto L309
	} else {
		goto L310
	}
L307:
	;
	goto L306
L308:
	;
	if v1132 == int32(0) {
		goto L325
	} else {
		goto L326
	}
L309:
	;
	F___clock_gettime(m, int32(1), v1724)
	mBase = m.M
	v1730 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1724)+8)))
	v1731 = *(*int64)(unsafe.Add(mBase, uint32(v1724)))
	v1735 = v1730 + (v1731*int64(1000000000) - v1673)
	goto L313
L310:
	;
	goto L311
L311:
	;
	v1818 = v842 << (uint(int32(6)) % 32)
	v1823 = *(*int64)(unsafe.Add(mBase, uint32(v1818)+uint32(_c_F_ExtendBufferedRelCommon[26])))
	*(*int64)(unsafe.Add(mBase, uint32(v1818)+uint32(_c_F_ExtendBufferedRelCommon[26]))) = v1823 + base.I64_extend_i32_u(v1715)
	v1827 = *(*int64)(unsafe.Add(mBase, uint32(v1818)+uint32(_c_F_ExtendBufferedRelCommon[27])))
	*(*int64)(unsafe.Add(mBase, uint32(v1818)+uint32(_c_F_ExtendBufferedRelCommon[27]))) = v1827 + v1718
	F_pgstat_count_backend_io_op(m, v1713, v842, int32(5), v1715, v1718)
	mBase = m.M
	v1832 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[10])) = uint8(v1832)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[11])) = uint8(v1832)
	m.G0 = v1724 + int32(16)
	goto L308
L312:
	;
	v1779 = int32(0)
	v1781 = v842 << (uint(int32(6)) % 32)
	v1786 = *(*int64)(unsafe.Add(mBase, uint32(v1781)+uint32(_c_F_ExtendBufferedRelCommon[28])))
	*(*int64)(unsafe.Add(mBase, uint32(v1781)+uint32(_c_F_ExtendBufferedRelCommon[28]))) = v1786 + v1735
	v1790 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[13]))
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v1790))|base.B2i32(int32(1)<<(uint(v1790)%32)&int32(_a_F_ExtendBufferedRelCommon_3) == v1779) == v1779 {
		goto L322
	} else {
		goto L323
	}
L313:
	;
	goto L314
L314:
	;
	v1742 = int32(_a_F_ExtendBufferedRelCommon_4)
	v1744 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[14]))
	v1746 = base.I64_div_s(v1735, int64(1000))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[14])) = v1744 + v1746
	switch v1713 {
	case 0:
		goto L318
	case 1:
		goto L317
	default:
		goto L312
	}
L317:
	;
	v1754 = int32(_a_F_ExtendBufferedRelCommon_5)
	v1756 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[15]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[15])) = v1756 + v1735
	goto L312
L318:
	;
	v1749 = int32(_a_F_ExtendBufferedRelCommon_6)
	v1751 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[16]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[16])) = v1751 + v1735
	goto L312
L322:
	;
	v1802 = *(*int64)(unsafe.Add(mBase, uint32(v1781)+uint32(_c_F_ExtendBufferedRelCommon[29])))
	*(*int64)(unsafe.Add(mBase, uint32(v1781)+uint32(_c_F_ExtendBufferedRelCommon[29]))) = v1802 + v1735
	v1806 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[10])) = uint8(v1806)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[18])) = uint8(v1806)
	goto L324
L323:
	;
	goto L324
L324:
	;
	goto L311
L325:
	;
	v1959 = int32(_a_F_ExtendBufferedRelCommon_21)
	v1961 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[30]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[30])) = v1961 + v1145
	v1968 = v1027
	v1975 = v1132
	goto L2
L326:
	;
	v1843 = v1027 + int32(1)
	v1845 = l3 & int32(32)
	v1847 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[23]))
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v1849 = int32(56)
	v1853 = v1847 + v1848*v1849 - v1849
	if l3&int32(8) == int32(0) {
		goto L329
	} else {
		goto L330
	}
L327:
	;
	v1869 = int32(1)
	v1870 = int32(0)
	F_TerminateBufferIO(m, v1853, v1870, int64(16777216), v1869, v1870)
	mBase = m.M
	v1875 = m.ExcPending
	if v1875 != 0 {
		goto L17
	} else {
		goto L336
	}
L328:
	;
	F_BufferLockAcquire(m, v1848, v1853, int32(3))
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		goto L17
	} else {
		goto L335
	}
L329:
	;
	if base.B2i32(v1845 == int32(0))|base.B2i32(v1843 != l5) != 0 {
		goto L327
	} else {
		goto L332
	}
L330:
	;
	goto L331
L331:
	;
	if v1848 < int32(0) {
		goto L327
	} else {
		goto L334
	}
L332:
	;
	if int32(0) <= v1848 {
		goto L328
	} else {
		goto L333
	}
L333:
	;
	goto L327
L334:
	;
	goto L328
L335:
	;
	goto L327
L336:
	;
	if v1132 == int32(1) {
		goto L325
	} else {
		goto L337
	}
L337:
	;
	v1878 = v1869
	goto L338
L338:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelCommon[23]))
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(l6+v1878<<(uint(int32(2))%32))))
	v1908 = int32(56)
	v1912 = v1903 + v1907*v1908 - v1908
	v1913 = int32(0)
	if base.B2i32(v1845 == v1913)|base.B2i32(v1878+v1843 != l5)|base.B2i32(v1907 < v1913) == v1913 {
		goto L340
	} else {
		goto L341
	}
L339:
	;
	goto L325
L340:
	;
	F_BufferLockAcquire(m, v1907, v1912, int32(3))
	mBase = m.M
	v1925 = m.ExcPending
	if v1925 != 0 {
		goto L17
	} else {
		goto L343
	}
L341:
	;
	goto L342
L342:
	;
	v1926 = int32(0)
	F_TerminateBufferIO(m, v1912, v1926, int64(16777216), int32(1), v1926)
	mBase = m.M
	v1931 = m.ExcPending
	if v1931 != 0 {
		goto L17
	} else {
		goto L344
	}
L343:
	;
	goto L342
L344:
	;
	v1933 = v1878 + int32(1)
	if v1933 != v1132 {
		v1878 = v1933
		goto L338
	} else {
		goto L345
	}
L345:
	;
	goto L339
L346:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1999 = m.ExcPending
	if v1999 != 0 {
		goto L17
	} else {
		goto L347
	}
L347:
	;
	if v30 == int32(0) {
		v2028 = v841
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v2030 = v27 + int32(152)
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+4))
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(v2028)))
	v2033 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+8))
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+12))
	F_GetRelationPath(m, v2030, v2031, v2032, v2033, v2034, l1)
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L17
	} else {
		goto L356
	}
L349:
	;
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	if v2002 != 0 {
		v2028 = v2002
		goto L348
	} else {
		goto L350
	}
L350:
	;
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v2004 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v2004
	v2006 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+32)) = v2006
	v2010 = F_smgropen(m, v27+int32(32), v2003)
	mBase = m.M
	v2011 = m.ExcPending
	if v2011 != 0 {
		goto L17
	} else {
		goto L351
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v2010
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(v2010)+72))
	if v2014 != 0 {
		goto L353
	} else {
		goto L354
	}
L352:
	;
	v2026 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v2028 = v2026
	goto L348
L353:
	;
	v2022 = v2014
	goto L355
L354:
	;
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v2010)+76))
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(v2010)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v2015)+4)) = v2016
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(v2010)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v2016))) = v2018
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(v2010)+72))
	v2022 = v2020
	goto L355
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2010)+72)) = v2022 + int32(1)
	goto L352
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v2030
	F_errmsg(m, int32(_a_F_ExtendBufferedRelCommon_9), v27+int32(16))
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		goto L17
	} else {
		goto L357
	}
L357:
	;
	F_errfinish(m, int32(_a_F_ExtendBufferedRelCommon_13), int32(2904), int32(_a_F_ExtendBufferedRelCommon_17))
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L17
	} else {
		goto L358
	}
L358:
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
func F_ean13_out(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14316(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
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
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1309 int32
	_ = v1309
	var v1314 int32
	_ = v1314
	v11 = m.G0
	v13 = v11 - int32(16)
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
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
	v170 = int32(_a_F_ean2string_0) + v118<<(uint(int32(3))%32)
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
	v400 = int32(_a_F_ean2string_1)
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
	v616 = int32(_a_F_ean2string_2)
	v618 = v102
	v619 = int32(_a_F_ean2string_3)
	goto L84
L99:
	;
	goto L100
L100:
	;
	v450 = int32(_a_F_ean2string_4)
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
	v616 = int32(_a_F_ean2string_5)
	v618 = int32(5)
	v619 = int32(_a_F_ean2string_6)
	goto L84
L115:
	;
	goto L116
L116:
	;
	v501 = int32(_a_F_ean2string_7)
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
	v616 = int32(_a_F_ean2string_8)
	v618 = int32(4)
	v619 = int32(_a_F_ean2string_9)
	goto L84
L131:
	;
	goto L132
L132:
	;
	v554 = int32(_a_F_ean2string_10)
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
	v616 = int32(_a_F_ean2string_11)
	v618 = v102
	v619 = int32(_a_F_ean2string_12)
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
	v609 = int32(_a_F_ean2string_13)
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
	v615 = int32(_a_F_ean2string_14)
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
	m.G0 = v13 + int32(16)
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
	v939 = int32(_a_F_ean2string_1)
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
	F_errcode(m, int32(50331778))
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L296
	} else {
		goto L298
	}
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(_a_F_ean2string_15)
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v16
	F_errmsg(m, int32(_a_F_ean2string_16), v13)
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L296
	} else {
		goto L299
	}
L299:
	;
	F_errfinish(m, int32(_a_F_ean2string_17), int32(657), int32(_a_F_ean2string_18))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L296
	} else {
		goto L300
	}
L300:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_elem_contained_by_range_support(m *base.Module, l0 int32) int64 {
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
	var v19 int64
	_ = v19
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	if v3 == int32(463) {
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
			return int64(0)
		} else {
			v19 = base.I64_extend_i32_u(v12)
			return v19
		}
	} else {
		v19 = int64(0)
		return v19
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
	var v76 int32
	_ = v76
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
						v76 = F_errdetail(m, int32(_a_F_enlargeStringInfo_1), v7+int32(16))
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
					F_errfinish(m, int32(_a_F_err_generic_string_2), int32(1750), int32(_a_F_err_generic_string_3))
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
				F_errfinish(m, int32(_a_F_err_generic_string_2), int32(1730), int32(_a_F_err_generic_string_3))
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
	var v126 int32
	_ = v126
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
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
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v247 int32
	_ = v247
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	v3 = int32(0)
	if l0 < int32(21) {
		v112 = l0
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[0]))
		if v14 != 0 {
			v15 = int32(24)
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
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123<<(uint(int32(2))%32))+uint32(_c_F_errstart[10])))
	if base.Ui32(v112-int32(15)) <= base.Ui32(v121) {
		if v126 < int32(22) {
			v144 = v121
		} else {
			v144 = int32(0)
		}
	} else {
		switch v112 - int32(20) {
		case 0, 3:
			v144 = int32(0)
		default:
			if v126 == int32(15) {
				if v112 <= int32(21) {
					v144 = int32(0)
				} else {
					v144 = int32(1)
				}
			} else {
				if v112 < v126 {
					v144 = int32(0)
				} else {
					v144 = int32(1)
				}
			}
		}
	}
	v145 = int32(0)
	if v112 == int32(16) {
		v164 = v145
	} else {
		v149 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[11]))
		if v149 != int32(2) {
			v164 = v145
		} else {
			v153 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_errstart[12])))
			if v153 == int32(1) {
				v164 = base.B2i32(int32(20) < v112)
			} else {
				v161 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[13]))
				v164 = base.B2i32(v112 == int32(17)) | base.B2i32(v161 <= v112)
			}
		}
	}
	v170 = (base.B2i32(int32(20) < v112) | v144 | v164) & int32(1)
	if v170 != 0 {
		v172 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[14]))
		if v172 == int32(0) {
			F_write_stderr(m, int32(_a_F_errstart_0), int32(0))
			mBase = m.M
			v262 = m.ExcPending
			if v262 != 0 {
				return int32(0)
			} else {
				F_pgl_exit(m, int32(2))
				mBase = m.M
				v265 = m.ExcPending
				if v265 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		} else {
			v175 = int32(_a_F_errstart_1)
			v177 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[15]))
			v179 = v177 + int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_errstart[15])) = v179
			if base.B2i32(v112 < int32(21))|base.B2i32(v177 <= int32(0)) != 0 {
				v200 = v179
				v201 = int32(_a_F_errstart_2)
				v203 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[4]))
				v205 = v203 + int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_errstart[4])) = v205
				if int32(5) <= v205 {
					*(*int32)(unsafe.Add(mBase, _c_F_errstart[4])) = int32(-1)
					F_errstart_cold(m, int32(24), int32(0))
					mBase = m.M
					v272 = m.ExcPending
					if v272 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_errstart_3), int32(0))
						mBase = m.M
						v276 = m.ExcPending
						if v276 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_errstart_4), int32(783), int32(_a_F_errstart_5))
							mBase = m.M
							v281 = m.ExcPending
							if v281 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v209 = int32(100)
					v210 = v205 * v209
					base.MemoryFill(m, v210+int32(_a_F_errstart_6), int32(0), v209)
					v217 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[16]))
					if l1 != 0 {
						v221 = l1
					} else {
						v221 = int32(_a_F_errstart_7)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[17]))) = v221
					*(*int32)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[18]))) = v221
					*(*uint8)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[19]))) = uint8(v164)
					*(*uint8)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[20]))) = uint8(v144)
					*(*int32)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[5]))) = v112
					*(*int32)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[21]))) = v217
					if int32(21) <= v112 {
						*(*int32)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[22]))) = int32(2600)
					} else {
						if int32(19) <= v112 {
							*(*int32)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[22]))) = int32(64)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[22]))) = int32(0)
						}
					}
					v247 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[14]))
					*(*int32)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[23]))) = v247
					*(*int32)(unsafe.Add(mBase, _c_F_errstart[15])) = v200 - int32(1)
					return v170
				}
			} else {
				F_MemoryContextReset(m, v172)
				mBase = m.M
				v189 = m.ExcPending
				if v189 != 0 {
					return int32(0)
				} else {
					v191 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[15]))
					if v191 < int32(3) {
						v200 = v191
					} else {
						v195 = int32(0)
						*(*int32)(unsafe.Add(mBase, _c_F_errstart[24])) = v195
						*(*int32)(unsafe.Add(mBase, _c_F_errstart[25])) = v195
						v200 = v191
					}
					v201 = int32(_a_F_errstart_2)
					v203 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[4]))
					v205 = v203 + int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_errstart[4])) = v205
					if int32(5) <= v205 {
						*(*int32)(unsafe.Add(mBase, _c_F_errstart[4])) = int32(-1)
						F_errstart_cold(m, int32(24), int32(0))
						mBase = m.M
						v272 = m.ExcPending
						if v272 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_errstart_3), int32(0))
							mBase = m.M
							v276 = m.ExcPending
							if v276 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_errstart_4), int32(783), int32(_a_F_errstart_5))
								mBase = m.M
								v281 = m.ExcPending
								if v281 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v209 = int32(100)
						v210 = v205 * v209
						base.MemoryFill(m, v210+int32(_a_F_errstart_6), int32(0), v209)
						v217 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[16]))
						if l1 != 0 {
							v221 = l1
						} else {
							v221 = int32(_a_F_errstart_7)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[17]))) = v221
						*(*int32)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[18]))) = v221
						*(*uint8)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[19]))) = uint8(v164)
						*(*uint8)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[20]))) = uint8(v144)
						*(*int32)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[5]))) = v112
						*(*int32)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[21]))) = v217
						if int32(21) <= v112 {
							*(*int32)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[22]))) = int32(2600)
						} else {
							if int32(19) <= v112 {
								*(*int32)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[22]))) = int32(64)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[22]))) = int32(0)
							}
						}
						v247 = *(*int32)(unsafe.Add(mBase, _c_F_errstart[14]))
						*(*int32)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_errstart[23]))) = v247
						*(*int32)(unsafe.Add(mBase, _c_F_errstart[15])) = v200 - int32(1)
						return v170
					}
				}
			}
		}
	} else {
		return v170
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
			v27 = v7 + v8<<(uint(int32(3))%32) + l1*int32(100) - int32(68)
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
	var v19 int32
	_ = v19
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int64
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
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
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v403 int32
	_ = v403
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int64
	_ = v462
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int64
	_ = v481
	var v482 int32
	_ = v482
	var v483 int64
	_ = v483
	var v484 int32
	_ = v484
	var v485 int64
	_ = v485
	var v486 int32
	_ = v486
	var v487 int64
	_ = v487
	var v488 int32
	_ = v488
	var v489 int64
	_ = v489
	var v493 int64
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int64
	_ = v498
	var v501 int64
	_ = v501
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v629 int32
	_ = v629
	var v636 int32
	_ = v636
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v744 int32
	_ = v744
	var v759 float64
	_ = v759
	var v762 int32
	_ = v762
	v5 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v19 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v34 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+120))
	if v35 == v34 {
		v762 = v34
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v33 = v19 + v20<<(uint(int32(2))%32)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v33 = v26 + v27<<(uint(int32(2))%32) - int32(4)
	goto L1
L5:
	;
	return v762
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v38 <= int32(0) {
		v762 = v34
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v47 = v5
	v49 = v5
	v53 = v5
	v54 = v5
	v55 = v5
	goto L8
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v53<<(uint(int32(2))%32))))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+16)))
	if v65 != int32(100) {
		v214 = v47
		v216 = v49
		v221 = v54
		v222 = v55
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v231 = int32(0)
	if v221 == v231 {
		v762 = v231
		goto L5
	} else {
		goto L39
	}
L10:
	;
	v228 = v53 + int32(1)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v228 < v229 {
		v47 = v214
		v49 = v216
		v53 = v228
		v54 = v221
		v55 = v222
		goto L8
	} else {
		goto L38
	}
L11:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+8)))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+20)))
	if v68 != v69 {
		v214 = v47
		v216 = v49
		v221 = v54
		v222 = v55
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v71 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if (base.B2i32(v47 != v187)|base.B2i32(v185 <= v49))&base.B2i32(v187 <= v47)|base.B2i32(v185+v187 < int32(2)) != 0 {
		v214 = v47
		v216 = v49
		v221 = v54
		v222 = v55
		goto L10
	} else {
		goto L37
	}
L14:
	;
	v74 = int32(0)
	v185 = v74
	v187 = v74
	goto L13
L15:
	;
	goto L16
L16:
	;
	v76 = int32(0)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v79 <= v76 {
		v185 = v76
		v187 = v76
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v86 = v76
	v88 = v76
	v92 = v76
	goto L18
L18:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100+v92<<(uint(int32(2))%32))))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v106 == int32(6) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v185 = v163
	v187 = v165
	goto L13
L20:
	;
	v178 = v92 + int32(1)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v178 < v179 {
		v86 = v163
		v88 = v165
		v92 = v178
		goto L18
	} else {
		goto L36
	}
L21:
	;
	v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(v105)+8)))
	if v109 <= int32(0) {
		v163 = v86
		v165 = v88
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	if v118 == int32(0) {
		v163 = v86
		v165 = v88
		goto L20
	} else {
		goto L27
	}
L24:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	v113 = F_bms_is_member(m, v109, v112)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int32(0)
L26:
	;
	v163 = v113 + v86
	v165 = v88
	goto L20
L27:
	;
	v121 = int32(0)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	if v122 <= v121 {
		v163 = v86
		v165 = v88
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v126 = v121
	goto L29
L29:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v144+v126<<(uint(int32(2))%32))))
	v149 = F_equal(m, v143, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L25
	} else {
		goto L31
	}
L30:
	;
	v163 = v86
	v165 = v88 + int32(1)
	goto L20
L31:
	;
	if v149 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v154 = v126 + int32(1)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	if v154 < v155 {
		v126 = v154
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
	v163 = v86
	v165 = v88
	goto L20
L36:
	;
	goto L19
L37:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v214 = v187
	v216 = v185
	v221 = v208
	v222 = v64
	goto L10
L38:
	;
	goto L9
L39:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+20)))
	v235 = m.G0
	v237 = v235 - int32(32)
	m.G0 = v237
	v242 = F_SearchSysCache2(m, int32(62), base.I64_extend_i32_u(v221), base.I64_extend_i32_u(v234))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L25
	} else {
		goto L42
	}
L40:
	;
	if v256 == int32(0) {
		v762 = v231
		goto L5
	} else {
		goto L57
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L25
	} else {
		goto L54
	}
L42:
	;
	if v242 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v248 = F_SysCacheGetAttr(m, int32(62), v242, int32(3), v237+int32(31))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
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
	v266 = m.ExcPending
	if v266 != 0 {
		goto L25
	} else {
		goto L51
	}
L46:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+31)))
	if v250 == int32(1) {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v254 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v248))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L25
	} else {
		goto L48
	}
L48:
	;
	v256 = F_statext_ndistinct_deserialize(m, v254)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L25
	} else {
		goto L49
	}
L49:
	;
	F_ReleaseCatCache(m, v242)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L25
	} else {
		goto L50
	}
L50:
	;
	m.G0 = v237 + int32(32)
	goto L40
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = v221
	F_errmsg_internal(m, int32(_a_F_estimate_multivariate_ndistinct_0), v237)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L25
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_estimate_multivariate_ndistinct_1), int32(155), int32(_a_F_estimate_multivariate_ndistinct_2))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v237)+20)) = v221
	*(*int32)(unsafe.Add(mBase, uint32(v237)+16)) = int32(100)
	F_errmsg_internal(m, int32(_a_F_estimate_multivariate_ndistinct_3), v237+int32(16))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L25
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_estimate_multivariate_ndistinct_1), int32(162), int32(_a_F_estimate_multivariate_ndistinct_2))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
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
	v295 = int32(0)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v222)+24))
	if v297 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	v299 = int32(16)
	v305 = (v298<<(uint(v299)%32) + int32(_a_F_estimate_multivariate_ndistinct_4)) >> (uint(v299) % 32)
	goto L60
L59:
	;
	v305 = v295
	goto L60
L60:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v306 == int32(0) {
		v425 = v295
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
	if v435 != 0 {
		goto L86
	} else {
		goto L87
	}
L62:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v306)+4))
	if v309 <= int32(0) {
		v425 = v295
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v321 = v295
	v323 = int32(0)
	goto L64
L64:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v306)+12))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v331+v323<<(uint(int32(2))%32))))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	if v337 != int32(6) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v425 = v403
	goto L61
L66:
	;
	v414 = v323 + int32(1)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v306)+4))
	if v414 < v415 {
		v321 = v403
		v323 = v414
		goto L64
	} else {
		goto L84
	}
L67:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v222)+24))
	if v340 == int32(0) {
		v403 = v321
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v383 = int32(*(*int16)(unsafe.Add(mBase, uint32(v336)+8)))
	if v383 <= int32(0) {
		v403 = v321
		goto L66
	} else {
		goto L80
	}
L70:
	;
	v343 = int32(0)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v340)+4))
	if v344 <= v343 {
		v403 = v321
		goto L66
	} else {
		goto L71
	}
L71:
	;
	v348 = v343
	goto L72
L72:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v335)))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v340)+12))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v366+v348<<(uint(int32(2))%32))))
	v371 = F_equal(m, v365, v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L25
	} else {
		goto L74
	}
L73:
	;
	v403 = v321
	goto L66
L74:
	;
	if v371 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v377 = F_bms_add_member(m, v321, base.I32_extend16_s(v305+(v348^int32(-1))))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L25
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v380 = v348 + int32(1)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v340)+4))
	if v380 < v381 {
		v348 = v380
		goto L72
	} else {
		goto L79
	}
L78:
	;
	v403 = v377
	goto L66
L79:
	;
	goto L73
L80:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v222)+20))
	v387 = F_bms_is_member(m, v383, v386)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L25
	} else {
		goto L81
	}
L81:
	;
	if v387 == int32(0) {
		v403 = v321
		goto L66
	} else {
		goto L82
	}
L82:
	;
	v393 = F_bms_add_member(m, v321, base.I32_extend16_s(v383+v305))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L25
	} else {
		goto L83
	}
L83:
	;
	v403 = v393
	goto L66
L84:
	;
	goto L65
L85:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v616 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L86:
	;
	v439 = int32(0)
	goto L89
L87:
	;
	goto L88
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L25
	} else {
		goto L115
	}
L89:
	;
	v459 = v256 + int32(16) + v439<<(uint(int32(4))%32)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)+8))
	v462 = int64(0)
	if v425 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	goto L88
L91:
	;
	v564 = v439 + int32(1)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
	if base.Ui32(v564) < base.Ui32(v565) {
		v439 = v564
		goto L89
	} else {
		goto L114
	}
L92:
	;
	if v460 != v506 {
		goto L91
	} else {
		goto L107
	}
L93:
	;
	v506 = int32(0)
	goto L92
L94:
	;
	goto L95
L95:
	;
	v467 = v425 + int32(8)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v425)+4))
	if v468 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	v506 = base.I32_popcnt(v471)
	goto L92
L97:
	;
	goto L98
L98:
	;
	v474 = v468 << (uint(int32(2)) % 32)
	if v474 <= int32(7) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v506 = base.I32_wrap_i64(v501)
	goto L92
L100:
	;
	if v474 == int32(0) {
		v501 = v462
		goto L99
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v498 = F_pg_popcount_optimized(m, v467, v474)
	mBase = m.M
	v501 = v498
	goto L99
L103:
	;
	v479 = v474
	v480 = v467
	v481 = v462
	goto L104
L104:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480)+3)))
	v483 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v482)+uint32(_c_F_estimate_multivariate_ndistinct[0]))))
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480)+2)))
	v485 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v484)+uint32(_c_F_estimate_multivariate_ndistinct[0]))))
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480)+1)))
	v487 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v486)+uint32(_c_F_estimate_multivariate_ndistinct[0]))))
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	v489 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v488)+uint32(_c_F_estimate_multivariate_ndistinct[0]))))
	v493 = v483 + (v485 + (v487 + (v481 + v489)))
	v494 = int32(4)
	v497 = v479 - v494
	if v497 != 0 {
		v479 = v497
		v480 = v480 + v494
		v481 = v493
		goto L104
	} else {
		goto L106
	}
L105:
	;
	v501 = v493
	goto L99
L106:
	;
	goto L105
L107:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v459)+8))
	if v508 <= int32(0) {
		goto L85
	} else {
		goto L108
	}
L108:
	;
	v513 = int32(0)
	goto L109
L109:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v459)+12))
	v534 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v530+v513<<(uint(int32(1))%32)))))
	v537 = F_bms_is_member(m, base.I32_extend16_s(v534+v305), v425)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L25
	} else {
		goto L111
	}
L110:
	;
	goto L85
L111:
	;
	if v537 == int32(0) {
		goto L91
	} else {
		goto L112
	}
L112:
	;
	v542 = v513 + int32(1)
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v459)+8))
	if v542 < v543 {
		v513 = v542
		goto L109
	} else {
		goto L113
	}
L113:
	;
	goto L110
L114:
	;
	goto L90
L115:
	;
	F_errmsg_internal(m, int32(_a_F_estimate_multivariate_ndistinct_5), int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L25
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_estimate_multivariate_ndistinct_6), int32(_a_F_estimate_multivariate_ndistinct_7), int32(_a_F_estimate_multivariate_ndistinct_8))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L25
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
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v744
	v759 = *(*float64)(unsafe.Add(mBase, uint32(v459)))
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = v759
	v762 = int32(1)
	goto L5
L119:
	;
	v744 = int32(0)
	goto L118
L120:
	;
	goto L121
L121:
	;
	v620 = int32(0)
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v616)+4))
	if v621 <= v620 {
		v744 = v620
		goto L118
	} else {
		goto L122
	}
L122:
	;
	v629 = v620
	v636 = int32(0)
	goto L123
L123:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v616)+12))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v643+v636<<(uint(int32(2))%32))))
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v647)))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v648)))
	if v649 == int32(6) {
		goto L127
	} else {
		goto L128
	}
L124:
	;
	v744 = v722
	goto L118
L125:
	;
	v737 = v636 + int32(1)
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v616)+4))
	if v737 < v738 {
		v629 = v722
		v636 = v737
		goto L123
	} else {
		goto L141
	}
L126:
	;
	v716 = F_lappend(m, v629, v647)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L25
	} else {
		goto L140
	}
L127:
	;
	v652 = int32(*(*int16)(unsafe.Add(mBase, uint32(v648)+8)))
	if v652 <= int32(0) {
		goto L126
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v222)+24))
	if v661 == int32(0) {
		goto L126
	} else {
		goto L133
	}
L130:
	;
	v657 = F_bms_is_member(m, base.I32_extend16_s(v652+v305), v425)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L25
	} else {
		goto L131
	}
L131:
	;
	if v657 == int32(0) {
		goto L126
	} else {
		goto L132
	}
L132:
	;
	v722 = v629
	goto L125
L133:
	;
	v664 = int32(0)
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v661)+4))
	if v665 <= v664 {
		goto L126
	} else {
		goto L134
	}
L134:
	;
	v669 = v664
	goto L135
L135:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v647)))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v661)+12))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v687+v669<<(uint(int32(2))%32))))
	v692 = F_equal(m, v686, v691)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L25
	} else {
		goto L137
	}
L136:
	;
	goto L126
L137:
	;
	if v692 != 0 {
		v722 = v629
		goto L125
	} else {
		goto L138
	}
L138:
	;
	v695 = v669 + int32(1)
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v661)+4))
	if v695 < v696 {
		v669 = v695
		goto L135
	} else {
		goto L139
	}
L139:
	;
	goto L136
L140:
	;
	v722 = v716
	goto L125
L141:
	;
	goto L124
}
func F_exec_assign_value(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32) {
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
	var v20 int64
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
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
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
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int64
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int64
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
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
					v68 = v20
					v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
					v72 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
					if v72 != v68 {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)))
						F_assign_simple_var(m, l0, l1, v68, v71&int32(1), base.B2i32(v82|v71 == int32(0)))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							m.G0 = v11 + int32(80)
							return
						}
					} else {
						v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)))
						if v74 != 0 {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)))
							F_assign_simple_var(m, l0, l1, v68, v71&int32(1), base.B2i32(v82|v71 == int32(0)))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								m.G0 = v11 + int32(80)
								return
							}
						} else {
							if v71&int32(1) == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = int32(0)
								m.G0 = v11 + int32(80)
								return
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)))
								F_assign_simple_var(m, l0, l1, v68, v71&int32(1), base.B2i32(v82|v71 == int32(0)))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
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
					v68 = v20
					v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
					v72 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
					if v72 != v68 {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)))
						F_assign_simple_var(m, l0, l1, v68, v71&int32(1), base.B2i32(v82|v71 == int32(0)))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							m.G0 = v11 + int32(80)
							return
						}
					} else {
						v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)))
						if v74 != 0 {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)))
							F_assign_simple_var(m, l0, l1, v68, v71&int32(1), base.B2i32(v82|v71 == int32(0)))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								m.G0 = v11 + int32(80)
								return
							}
						} else {
							if v71&int32(1) == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = int32(0)
								m.G0 = v11 + int32(80)
								return
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)))
								F_assign_simple_var(m, l0, l1, v68, v71&int32(1), base.B2i32(v82|v71 == int32(0)))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
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
						v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+12)))
						v66 = F_datumTransfer(m, v20, int32(0), v65)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							v68 = v66
							v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
							v72 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
							if v72 != v68 {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)))
								F_assign_simple_var(m, l0, l1, v68, v71&int32(1), base.B2i32(v82|v71 == int32(0)))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									m.G0 = v11 + int32(80)
									return
								}
							} else {
								v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)))
								if v74 != 0 {
									v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)))
									F_assign_simple_var(m, l0, l1, v68, v71&int32(1), base.B2i32(v82|v71 == int32(0)))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										m.G0 = v11 + int32(80)
										return
									}
								} else {
									if v71&int32(1) == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = int32(0)
										m.G0 = v11 + int32(80)
										return
									} else {
										v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)))
										F_assign_simple_var(m, l0, l1, v68, v71&int32(1), base.B2i32(v82|v71 == int32(0)))
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
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
						v52 = base.I32_wrap_i64(v20)
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
						if v53 == int32(1) {
							v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
							if v56 == int32(3) {
								v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+12)))
								v66 = F_datumTransfer(m, v20, int32(0), v65)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									v68 = v66
									v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
									v72 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
									if v72 != v68 {
										v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)))
										F_assign_simple_var(m, l0, l1, v68, v71&int32(1), base.B2i32(v82|v71 == int32(0)))
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return
										} else {
											m.G0 = v11 + int32(80)
											return
										}
									} else {
										v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)))
										if v74 != 0 {
											v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
											v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)))
											F_assign_simple_var(m, l0, l1, v68, v71&int32(1), base.B2i32(v82|v71 == int32(0)))
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return
											} else {
												m.G0 = v11 + int32(80)
												return
											}
										} else {
											if v71&int32(1) == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = int32(0)
												m.G0 = v11 + int32(80)
												return
											} else {
												v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
												v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)))
												F_assign_simple_var(m, l0, l1, v68, v71&int32(1), base.B2i32(v82|v71 == int32(0)))
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
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
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
								v61 = F_expand_array(m, v20, v59, int32(0))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									v68 = v61
									v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
									v72 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
									if v72 != v68 {
										v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)))
										F_assign_simple_var(m, l0, l1, v68, v71&int32(1), base.B2i32(v82|v71 == int32(0)))
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return
										} else {
											m.G0 = v11 + int32(80)
											return
										}
									} else {
										v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)))
										if v74 != 0 {
											v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
											v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)))
											F_assign_simple_var(m, l0, l1, v68, v71&int32(1), base.B2i32(v82|v71 == int32(0)))
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return
											} else {
												m.G0 = v11 + int32(80)
												return
											}
										} else {
											if v71&int32(1) == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = int32(0)
												m.G0 = v11 + int32(80)
												return
											} else {
												v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
												v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)))
												F_assign_simple_var(m, l0, l1, v68, v71&int32(1), base.B2i32(v82|v71 == int32(0)))
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
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
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
							v61 = F_expand_array(m, v20, v59, int32(0))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								v68 = v61
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
								v72 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
								if v72 != v68 {
									v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)))
									F_assign_simple_var(m, l0, l1, v68, v71&int32(1), base.B2i32(v82|v71 == int32(0)))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										m.G0 = v11 + int32(80)
										return
									}
								} else {
									v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)))
									if v74 != 0 {
										v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)))
										F_assign_simple_var(m, l0, l1, v68, v71&int32(1), base.B2i32(v82|v71 == int32(0)))
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return
										} else {
											m.G0 = v11 + int32(80)
											return
										}
									} else {
										if v71&int32(1) == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = int32(0)
											m.G0 = v11 + int32(80)
											return
										} else {
											v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
											v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)))
											F_assign_simple_var(m, l0, l1, v68, v71&int32(1), base.B2i32(v82|v71 == int32(0)))
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
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
			v88 = int32(0)
			F_exec_move_row(m, l0, l1, v88, v88)
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return
			} else {
				m.G0 = v11 + int32(80)
				return
			}
		} else {
			v92 = F_type_is_rowtype(m, l4)
			mBase = m.M
			v93 = m.ExcPending
			if v93 != 0 {
				return
			} else {
				if v92 == int32(0) {
					F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
					mBase = m.M
					v181 = m.ExcPending
					if v181 != 0 {
						return
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v184 = m.ExcPending
						if v184 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_exec_assign_value_5), int32(0))
							mBase = m.M
							v188 = m.ExcPending
							if v188 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_6), int32(_a_F_exec_assign_value_4))
								mBase = m.M
								v193 = m.ExcPending
								if v193 != 0 {
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
					v97 = m.ExcPending
					if v97 != 0 {
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
			v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
			if v98 == int32(1) {
				F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
				mBase = m.M
				v197 = m.ExcPending
				if v197 != 0 {
					return
				} else {
					F_errcode(m, int32(67108994))
					mBase = m.M
					v200 = m.ExcPending
					if v200 != 0 {
						return
					} else {
						v201 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v201
						F_errmsg(m, int32(_a_F_exec_assign_value_1), v11+int32(32))
						mBase = m.M
						v207 = m.ExcPending
						if v207 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_7), int32(_a_F_exec_assign_value_4))
							mBase = m.M
							v212 = m.ExcPending
							if v212 != 0 {
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
				v101 = int32(0)
				F_exec_move_row(m, l0, l1, v101, v101)
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return
				} else {
					m.G0 = v11 + int32(80)
					return
				}
			}
		} else {
			v105 = F_type_is_rowtype(m, l4)
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return
			} else {
				if v105 == int32(0) {
					F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
					mBase = m.M
					v216 = m.ExcPending
					if v216 != 0 {
						return
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v219 = m.ExcPending
						if v219 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_exec_assign_value_8), int32(0))
							mBase = m.M
							v223 = m.ExcPending
							if v223 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_9), int32(_a_F_exec_assign_value_4))
								mBase = m.M
								v228 = m.ExcPending
								if v228 != 0 {
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
					v110 = m.ExcPending
					if v110 != 0 {
						return
					} else {
						m.G0 = v11 + int32(80)
						return
					}
				}
			}
		}
	case 3:
		v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v116 = *(*int32)(unsafe.Add(mBase, uint32(v111+v112<<(uint(int32(2))%32))))
		v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+36))
		if v117 == int32(0) {
			F_instantiate_empty_record_variable(m, l0, v116)
			mBase = m.M
			v121 = m.ExcPending
			if v121 != 0 {
				return
			} else {
				v122 = *(*int32)(unsafe.Add(mBase, uint32(v116)+36))
				v123 = v122
				v124 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
				v125 = *(*int64)(unsafe.Add(mBase, uint32(v123)+48))
				if v124 != v125 {
					v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v130 = F_expanded_record_lookup_field(m, v123, v127, l1+int32(32))
					mBase = m.M
					v131 = m.ExcPending
					if v131 != 0 {
						return
					} else {
						if v130 == int32(0) {
							F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
							mBase = m.M
							v232 = m.ExcPending
							if v232 != 0 {
								return
							} else {
								F_errcode(m, int32(50360452))
								mBase = m.M
								v235 = m.ExcPending
								if v235 != 0 {
									return
								} else {
									v236 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
									v237 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v237
									*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v236
									F_errmsg(m, int32(_a_F_exec_assign_value_10), v11-int32(-64))
									mBase = m.M
									v244 = m.ExcPending
									if v244 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_11), int32(_a_F_exec_assign_value_4))
										mBase = m.M
										v249 = m.ExcPending
										if v249 != 0 {
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
							v134 = *(*int64)(unsafe.Add(mBase, uint32(v123)+48))
							*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v134
							v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
							if v136 <= int32(0) {
								F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
								mBase = m.M
								v253 = m.ExcPending
								if v253 != 0 {
									return
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v256 = m.ExcPending
									if v256 != 0 {
										return
									} else {
										v257 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v257
										F_errmsg(m, int32(_a_F_exec_assign_value_12), v11+int32(48))
										mBase = m.M
										v263 = m.ExcPending
										if v263 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_13), int32(_a_F_exec_assign_value_4))
											mBase = m.M
											v268 = m.ExcPending
											if v268 != 0 {
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
								v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
								v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
								v143 = F_exec_cast_value(m, l0, l2, v11+int32(79), l4, l5, v141, v142)
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return
								} else {
									v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
									v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
									v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
									v150 = int32(1)
									F_expanded_record_set_field_internal(m, v123, v145, v143, v146, (v147^int32(-1))&v150, v150)
									mBase = m.M
									v154 = m.ExcPending
									if v154 != 0 {
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
					v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
					if v136 <= int32(0) {
						F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
						mBase = m.M
						v253 = m.ExcPending
						if v253 != 0 {
							return
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v256 = m.ExcPending
							if v256 != 0 {
								return
							} else {
								v257 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v257
								F_errmsg(m, int32(_a_F_exec_assign_value_12), v11+int32(48))
								mBase = m.M
								v263 = m.ExcPending
								if v263 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_13), int32(_a_F_exec_assign_value_4))
									mBase = m.M
									v268 = m.ExcPending
									if v268 != 0 {
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
						v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
						v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						v143 = F_exec_cast_value(m, l0, l2, v11+int32(79), l4, l5, v141, v142)
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return
						} else {
							v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
							v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
							v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
							v150 = int32(1)
							F_expanded_record_set_field_internal(m, v123, v145, v143, v146, (v147^int32(-1))&v150, v150)
							mBase = m.M
							v154 = m.ExcPending
							if v154 != 0 {
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
			v123 = v117
			v124 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
			v125 = *(*int64)(unsafe.Add(mBase, uint32(v123)+48))
			if v124 != v125 {
				v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v130 = F_expanded_record_lookup_field(m, v123, v127, l1+int32(32))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					if v130 == int32(0) {
						F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
						mBase = m.M
						v232 = m.ExcPending
						if v232 != 0 {
							return
						} else {
							F_errcode(m, int32(50360452))
							mBase = m.M
							v235 = m.ExcPending
							if v235 != 0 {
								return
							} else {
								v236 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
								v237 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v237
								*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v236
								F_errmsg(m, int32(_a_F_exec_assign_value_10), v11-int32(-64))
								mBase = m.M
								v244 = m.ExcPending
								if v244 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_11), int32(_a_F_exec_assign_value_4))
									mBase = m.M
									v249 = m.ExcPending
									if v249 != 0 {
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
						v134 = *(*int64)(unsafe.Add(mBase, uint32(v123)+48))
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v134
						v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						if v136 <= int32(0) {
							F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
							mBase = m.M
							v253 = m.ExcPending
							if v253 != 0 {
								return
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v256 = m.ExcPending
								if v256 != 0 {
									return
								} else {
									v257 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v257
									F_errmsg(m, int32(_a_F_exec_assign_value_12), v11+int32(48))
									mBase = m.M
									v263 = m.ExcPending
									if v263 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_13), int32(_a_F_exec_assign_value_4))
										mBase = m.M
										v268 = m.ExcPending
										if v268 != 0 {
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
							v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
							v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
							v143 = F_exec_cast_value(m, l0, l2, v11+int32(79), l4, l5, v141, v142)
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
								v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
								v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
								v150 = int32(1)
								F_expanded_record_set_field_internal(m, v123, v145, v143, v146, (v147^int32(-1))&v150, v150)
								mBase = m.M
								v154 = m.ExcPending
								if v154 != 0 {
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
				v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
				if v136 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(_a_F_exec_assign_value_0))
					mBase = m.M
					v253 = m.ExcPending
					if v253 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v256 = m.ExcPending
						if v256 != 0 {
							return
						} else {
							v257 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v257
							F_errmsg(m, int32(_a_F_exec_assign_value_12), v11+int32(48))
							mBase = m.M
							v263 = m.ExcPending
							if v263 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_13), int32(_a_F_exec_assign_value_4))
								mBase = m.M
								v268 = m.ExcPending
								if v268 != 0 {
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
					v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
					v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					v143 = F_exec_cast_value(m, l0, l2, v11+int32(79), l4, l5, v141, v142)
					mBase = m.M
					v144 = m.ExcPending
					if v144 != 0 {
						return
					} else {
						v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
						v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
						v150 = int32(1)
						F_expanded_record_set_field_internal(m, v123, v145, v143, v146, (v147^int32(-1))&v150, v150)
						mBase = m.M
						v154 = m.ExcPending
						if v154 != 0 {
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
		v158 = m.ExcPending
		if v158 != 0 {
			return
		} else {
			v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v159
			F_errmsg_internal(m, int32(_a_F_exec_assign_value_14), v11)
			mBase = m.M
			v163 = m.ExcPending
			if v163 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_exec_assign_value_2), int32(_a_F_exec_assign_value_15), int32(_a_F_exec_assign_value_4))
				mBase = m.M
				v168 = m.ExcPending
				if v168 != 0 {
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
func F_exec_cast_value(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int64 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
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
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
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
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int64
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v351 int64
	_ = v351
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	if (base.B2i32(l4 == l6)|base.B2i32(l6 == int32(-1)))&base.B2i32(l3 == l5) != 0 {
		v351 = l1
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v17 + int32(32)
	return v351
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l3
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v31 = v17 + int32(16)
	v34 = v17 + int32(15)
	v35 = F_hash_search(m, v29, v31, int32(1), v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int64(0)
L4:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)))
	if v39 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	if v61 != 0 {
		goto L17
	} else {
		goto L18
	}
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[0]))
	v45 = F_hash_search(m, v43, v31, int32(1), v34)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L3
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	v60 = v59
	goto L5
L9:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)))
	if v47 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = int32(0)
	goto L12
L11:
	;
	goto L12
L12:
	;
	v52 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v52
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)) = uint8(v52)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v45
	v60 = v45
	goto L5
L13:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2])) = v336
	*(*int64)(unsafe.Add(mBase, uint32(v335)+40)) = l1
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	*(*uint8)(unsafe.Add(mBase, uint32(v335)+48)) = uint8(v339)
	v341 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)) = uint8(v341)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v332)+24))
	v344 = m.T0[v343].(func(*base.Module, int32, int32, int32) int64)(m, v332, v335, l2)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L3
	} else {
		goto L83
	}
L14:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2])) = v323
	v326 = F_ExecInitExpr(m, v278, int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L3
	} else {
		goto L82
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_cast_value_4))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L3
	} else {
		goto L76
	}
L16:
	;
	if v278 == int32(0) {
		v351 = l1
		goto L1
	} else {
		goto L71
	}
L17:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+8)))
	if v62 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v77 = int32(_a_F_exec_cast_value_0)
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2]))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2])) = v81
	v84 = F_palloc0(m, int32(16))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L3
	} else {
		goto L24
	}
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	v278 = v65
	goto L16
L21:
	;
	goto L22
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	F_MemoryContextDelete(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+20)) = int32(0)
	goto L19
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+8)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = int32(34)
	v90 = F_get_typcollation(m, l3)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+12)) = v90
	if base.B2i32(l3 == int32(705))|base.B2i32(l3 == int32(2249)) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v132 = m.G0
	v134 = v132 - int32(16)
	m.G0 = v134
	v140 = m.G0
	v142 = v140 - int32(528)
	m.G0 = v142
	v144 = int32(400)
	v145 = v142 + v144
	v146 = int32(0)
	base.MemoryFill(m, v145, v146, int32(128))
	*(*int32)(unsafe.Add(mBase, uint32(v142)+468)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v142)+400)) = int32(268)
	base.MemoryFill(m, v142, v146, v144)
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = int32(269)
	*(*int32)(unsafe.Add(mBase, uint32(v142)+8)) = v145
	v159 = F_eval_const_expressions(m, v142, v131)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L3
	} else {
		goto L36
	}
L27:
	;
	v101 = int32(2)
	v104 = F_coerce_to_target_type(m, int32(0), v84, l3, l5, l6, v101, v101, int32(-1))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v108 = F_palloc0(m, int32(24))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	if v104 != 0 {
		v131 = v104
		goto L26
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = int32(28)
	v112 = int32(2281)
	if base.B2i32(l3 == v112)|base.B2i32(l5 == v112) != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	v117 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v108)+20)) = v117
	*(*int64)(unsafe.Add(mBase, uint32(v108)+12)) = int64(8589934592)
	*(*int32)(unsafe.Add(mBase, uint32(v108)+8)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v84
	if l6 == v117 {
		v131 = v108
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v129 = F_coerce_to_target_type(m, int32(0), v108, l5, l5, l6, int32(1), int32(2), int32(-1))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v131 = v129
	goto L26
L36:
	;
	F_fix_opfuncids(m, v159)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v163 = F_extract_query_dependencies_walker(m, v159, v142)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v142)+464))
	*(*int32)(unsafe.Add(mBase, uint32(v134+int32(12)))) = v165
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v142)+468))
	*(*int32)(unsafe.Add(mBase, uint32(v134+int32(8)))) = v167
	m.G0 = v142 + int32(528)
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2]))
	v178 = F_AllocSetContextCreateInternal(m, v173, int32(_a_F_exec_cast_value_1), int32(0), int32(1024), int32(_a_F_exec_cast_value_2))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	v180 = int32(_a_F_exec_cast_value_0)
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2])) = v178
	v185 = F_palloc(m, int32(32))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = int32(838275847)
	v189 = F_copyObjectImpl(m, v159)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	v191 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v185)+8)) = uint8(v191)
	*(*int32)(unsafe.Add(mBase, uint32(v185)+4)) = v189
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v195 = F_copyObjectImpl(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185)+12)) = v195
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	v199 = F_copyObjectImpl(m, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185)+20)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v185)+16)) = v199
	*(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2])) = v181
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[3]))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v178)+16))
	if v210 != v206 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[4]))
	if v240 != 0 {
		goto L62
	} else {
		goto L63
	}
L45:
	;
	if v210 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	if v206 != 0 {
		goto L55
	} else {
		goto L56
	}
L49:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v178)+28))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v178)+24))
	if v215 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v214 == int32(0) {
		goto L48
	} else {
		goto L54
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215)+28)) = v214
	goto L50
L52:
	;
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210)+20)) = v214
	goto L50
L54:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v178)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v214)+24)) = v220
	goto L48
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v178)+16)) = v206
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v206)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v178)+28)) = v227
	if v227 != 0 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v178)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v178)+16)) = int32(0)
	goto L47
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227)+24)) = v178
	goto L60
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206)+20)) = v178
	goto L44
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185)+24)) = v247
	v249 = int32(_a_F_exec_cast_value_3)
	*(*int32)(unsafe.Add(mBase, uint32(v185)+28)) = v249
	v252 = v185 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v247)+4)) = v252
	*(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[5])) = v252
	m.G0 = v134 + int32(16)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	if v260 == int32(27) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[5]))
	v247 = v242
	goto L61
L63:
	;
	goto L64
L64:
	;
	v244 = int32(_a_F_exec_cast_value_3)
	*(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[4])) = v244
	v247 = v244
	goto L61
L65:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	if v264 != v84 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v267 = v259
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v60)+20)) = v185
	v270 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v270
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)) = uint8(v270)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v270
	*(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2])) = v78
	v278 = v267
	goto L16
L68:
	;
	v266 = v259
	goto L70
L69:
	;
	v266 = int32(0)
	goto L70
L70:
	;
	v267 = v266
	goto L67
L71:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[1]))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+44))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
	if v290 != v291 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2]))
	v320 = v294
	goto L14
L73:
	;
	goto L74
L74:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2]))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)))
	if v297 != 0 {
		v320 = v296
		goto L14
	} else {
		goto L75
	}
L75:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	v332 = v298
	v333 = v296
	goto L13
L76:
	;
	F_errcode(m, int32(101744772))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L3
	} else {
		goto L77
	}
L77:
	;
	v306 = F_format_type_be(m, l3)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	v308 = F_format_type_be(m, l5)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v306
	F_errmsg(m, int32(_a_F_exec_cast_value_5), v17)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L3
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_exec_cast_value_6), int32(_a_F_exec_cast_value_7), int32(_a_F_exec_cast_value_8))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L3
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v290
	v329 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)) = uint8(v329)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v326
	v332 = v326
	v333 = v320
	goto L13
L83:
	;
	v346 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)) = uint8(v346)
	*(*int32)(unsafe.Add(mBase, _c_F_exec_cast_value[2])) = v333
	v351 = v344
	goto L1
}
func F_exec_move_row_from_datum(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
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
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = base.I32_wrap_i64(l2)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v14 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return
L2:
	;
	if v346&int32(5) == int32(0) {
		goto L144
	} else {
		goto L145
	}
L3:
	;
	v226 = int32(_a_F_exec_move_row_from_datum_0)
	v227 = *(*int32)(unsafe.Add(mBase, _c_F_exec_move_row_from_datum[0]))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_move_row_from_datum[0])) = v230
	v232 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L11
	} else {
		goto L104
	}
L4:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v17&int32(254) != int32(2) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l2))+2))
	goto L6
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v24 != int32(2) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v345 = int32(0)
	v346 = v27
	goto L2
L8:
	;
	goto L9
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v23 == v29 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_revalidate_rectypeid(m, l1)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v33 != int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v89 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L14:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v36 != int32(3) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
	if v39 != v40 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v39 != int32(2249) {
		goto L13
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	if v52 != v48 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+28)))
	if v44&int32(64) != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v81 != 0 {
		goto L38
	} else {
		goto L39
	}
L22:
	;
	if v52 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	if v48 != 0 {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	if v57 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v56 == int32(0) {
		goto L25
	} else {
		goto L31
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+28)) = v56
	goto L27
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v56
	goto L27
L31:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+24)) = v62
	goto L25
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v48
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v69
	if v69 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v47)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = int32(0)
	goto L24
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+24)) = v47
	goto L37
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v47
	goto L21
L38:
	;
	F_DeleteExpandedObject(m, base.I64_extend_i32_u(v81+int32(12)))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L11
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v23
	goto L1
L41:
	;
	goto L40
L42:
	;
	v118 = F_make_expanded_record_for_rec(m, l0, l1, v23)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L11
	} else {
		goto L52
	}
L43:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+28)))
	if v92&int32(1) == int32(0) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v23)+36))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+36))
	if v97 != v98 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	if v97 == int32(2249) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	if v102 < int32(0) {
		goto L42
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
	v109 = int32(1)
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
	F_expanded_record_set_tuple(m, v89, v108, v109, (v110^int32(-1))&v109)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L11
	} else {
		goto L51
	}
L49:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v89)+40))
	if v102 != v105 {
		goto L42
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	goto L1
L52:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	if v120&int32(1) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if v120&int32(5) != 0 {
		v345 = v118
		v346 = v120
		goto L2
	} else {
		goto L81
	}
L54:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v125 != int32(2249) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v23)+36))
	if v125 != v128 {
		goto L53
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
	v131 = int32(1)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
	F_expanded_record_set_tuple(m, v118, v130, v131, (v132^int32(-1))&v131)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L11
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v139)+16))
	if v144 != v140 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v173 != 0 {
		goto L77
	} else {
		goto L78
	}
L61:
	;
	if v144 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	if v140 != 0 {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v139)+28))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v139)+24))
	if v149 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v148 == int32(0) {
		goto L64
	} else {
		goto L70
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149)+28)) = v148
	goto L66
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144)+20)) = v148
	goto L66
L70:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v139)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v148)+24)) = v154
	goto L64
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v139)+16)) = v140
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v140)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v139)+28)) = v161
	if v161 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v139)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v139)+16)) = int32(0)
	goto L63
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161)+24)) = v139
	goto L76
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140)+20)) = v139
	goto L60
L77:
	;
	F_DeleteExpandedObject(m, base.I64_extend_i32_u(v173+int32(12)))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L11
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v118
	goto L1
L80:
	;
	goto L79
L81:
	;
	F_deconstruct_expanded_record(m, v118)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L11
	} else {
		goto L82
	}
L82:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v185)+16))
	if v190 != v186 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v219 != 0 {
		goto L100
	} else {
		goto L101
	}
L84:
	;
	if v190 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L86
L86:
	;
	goto L83
L87:
	;
	if v186 != 0 {
		goto L94
	} else {
		goto L95
	}
L88:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v185)+28))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v185)+24))
	if v195 != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	if v194 == int32(0) {
		goto L87
	} else {
		goto L93
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+28)) = v194
	goto L89
L91:
	;
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+20)) = v194
	goto L89
L93:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v185)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v194)+24)) = v200
	goto L87
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v185)+16)) = v186
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+28)) = v207
	if v207 != 0 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L96
L96:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v185)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v185)+16)) = int32(0)
	goto L86
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207)+24)) = v185
	goto L99
L98:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186)+20)) = v185
	goto L83
L100:
	;
	F_DeleteExpandedObject(m, base.I64_extend_i32_u(v219+int32(12)))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L11
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v118
	goto L1
L103:
	;
	goto L102
L104:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_move_row_from_datum[0])) = v227
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v232
	v238 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v238
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+20)) = uint16(v238)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(-1)
	v244 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(base.Ui32(v236) >> (uint(v244) % 32))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v232)+8))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v249 != v244 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v336 = F_lookup_rowtype_tupdesc(m, v248, v247)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L11
	} else {
		goto L140
	}
L106:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v252 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if base.B2i32(v273 != int32(2249))&base.B2i32(v248 != v273) != 0 {
		goto L105
	} else {
		goto L116
	}
L108:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v252)+36))
	if v248 != v255 {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	if v248 == int32(2249) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	if v247 < int32(0) {
		goto L107
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v265 = int32(1)
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
	F_expanded_record_set_tuple(m, v252, v11+int32(12), v265, (v266^int32(-1))&v265)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L11
	} else {
		goto L115
	}
L113:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v252)+40))
	if v247 != v261 {
		goto L107
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	goto L1
L116:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)+20))
	v280 = F_make_expanded_record_from_typeid(m, v248, v247, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L11
	} else {
		goto L117
	}
L117:
	;
	v284 = int32(1)
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
	F_expanded_record_set_tuple(m, v280, v11+int32(12), v284, (v285^int32(-1))&v284)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L11
	} else {
		goto L118
	}
L118:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v280)+8))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v292)+16))
	if v297 != v293 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v326 != 0 {
		goto L136
	} else {
		goto L137
	}
L120:
	;
	if v297 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	goto L122
L122:
	;
	goto L119
L123:
	;
	if v293 != 0 {
		goto L130
	} else {
		goto L131
	}
L124:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v292)+28))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v292)+24))
	if v302 != 0 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	if v301 == int32(0) {
		goto L123
	} else {
		goto L129
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v302)+28)) = v301
	goto L125
L127:
	;
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v297)+20)) = v301
	goto L125
L129:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v292)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v301)+24)) = v307
	goto L123
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v292)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v292)+16)) = v293
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v293)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v292)+28)) = v314
	if v314 != 0 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	goto L132
L132:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v292)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v292)+16)) = int32(0)
	goto L122
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314)+24)) = v292
	goto L135
L134:
	;
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293)+20)) = v292
	goto L119
L136:
	;
	F_DeleteExpandedObject(m, base.I64_extend_i32_u(v326+int32(12)))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L11
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v280
	goto L1
L139:
	;
	goto L138
L140:
	;
	F_exec_move_row(m, l0, l1, v11+int32(12), v336)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L11
	} else {
		goto L141
	}
L141:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v336)+12))
	if v340 < int32(0) {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_DecrTupleDescRefCount(m, v336)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L11
	} else {
		goto L143
	}
L143:
	;
	goto L1
L144:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
	if v353 != 0 {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	goto L146
L146:
	;
	F_deconstruct_expanded_record(m, v23)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L11
	} else {
		goto L152
	}
L147:
	;
	v356 = v353
	goto L149
L148:
	;
	v354 = F_expanded_record_fetch_tupdesc(m, v23)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L11
	} else {
		goto L150
	}
L149:
	;
	F_exec_move_row(m, l0, l1, int32(0), v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L11
	} else {
		goto L151
	}
L150:
	;
	v356 = v354
	goto L149
L151:
	;
	goto L1
L152:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
	if v363 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v366 = v363
	goto L155
L154:
	;
	v364 = F_expanded_record_fetch_tupdesc(m, v23)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L11
	} else {
		goto L156
	}
L155:
	;
	F_exec_move_row_from_fields(m, l0, l1, v345, v361, v362, v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L11
	} else {
		goto L157
	}
L156:
	;
	v366 = v364
	goto L155
L157:
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
					v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
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
						v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
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
								v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+44))
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
									v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
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
										v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
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
												v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+44))
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
						v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
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
							v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
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
									v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+44))
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
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
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
					v73 = F_RE_compile_and_cache(m, v65, v69|int32(16), int32(100))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						v78 = F_palloc_mul(m, int32(4), v68+int32(1))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							v80 = F_pg_mb2wchar_with_len(m, v67, v78, v68)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v82 = int32(0)
								v85 = F_RE_wchar_execute(m, v78, v80, v82, v82, v82)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v78)
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return int32(0)
									} else {
										v94 = v85
										return v94
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
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
						v73 = F_RE_compile_and_cache(m, v65, v69|int32(16), int32(100))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							v78 = F_palloc_mul(m, int32(4), v68+int32(1))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								v80 = F_pg_mb2wchar_with_len(m, v67, v78, v68)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									v82 = int32(0)
									v85 = F_RE_wchar_execute(m, v78, v80, v82, v82, v82)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v78)
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return int32(0)
										} else {
											v94 = v85
											return v94
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
										F_errsave_finish(m, int32(0), int32(_a_F_executeLikeRegex_1), int32(711), int32(_a_F_executeLikeRegex_2))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											v63 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
											v65 = v63
											v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
											v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											v69 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
											v73 = F_RE_compile_and_cache(m, v65, v69|int32(16), int32(100))
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												v78 = F_palloc_mul(m, int32(4), v68+int32(1))
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													v80 = F_pg_mb2wchar_with_len(m, v67, v78, v68)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														v82 = int32(0)
														v85 = F_RE_wchar_execute(m, v78, v80, v82, v82, v82)
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v78)
															mBase = m.M
															v88 = m.ExcPending
															if v88 != 0 {
																return int32(0)
															} else {
																v94 = v85
																return v94
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
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
								v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								v69 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
								v73 = F_RE_compile_and_cache(m, v65, v69|int32(16), int32(100))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									v78 = F_palloc_mul(m, int32(4), v68+int32(1))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										v80 = F_pg_mb2wchar_with_len(m, v67, v78, v68)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											v82 = int32(0)
											v85 = F_RE_wchar_execute(m, v78, v80, v82, v82, v82)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												F_pfree(m, v78)
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return int32(0)
												} else {
													v94 = v85
													return v94
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
			v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v69 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
			v73 = F_RE_compile_and_cache(m, v65, v69|int32(16), int32(100))
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				v78 = F_palloc_mul(m, int32(4), v68+int32(1))
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int32(0)
				} else {
					v80 = F_pg_mb2wchar_with_len(m, v67, v78, v68)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						v82 = int32(0)
						v85 = F_RE_wchar_execute(m, v78, v80, v82, v82, v82)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v78)
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int32(0)
							} else {
								v94 = v85
								return v94
							}
						}
					}
				}
			}
		}
	} else {
		v94 = int32(2)
		return v94
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
	var v25 int32
	_ = v25
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
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
	var v254 int32
	_ = v254
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v327 int32
	_ = v327
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
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
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
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v422 int32
	_ = v422
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
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
	v51 = int32(0)
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v29
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v51
	if v27 != v25 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_get_relation_notnullatts(m, l0, l5)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v57 = int32(0)
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
L10:
	;
	goto L9
L11:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v334 = F_lappend(m, v333, v64)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L64
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(324)
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
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L60
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L56
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L53
	}
L17:
	;
	v96 = int32(0)
	v101 = v57
	v102 = v57
	goto L20
L18:
	;
	v254 = v57
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = v254
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v267
	m.G0 = v61 + int32(48)
	goto L11
L20:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v120 = v77 + v114<<(uint(int32(3))%32) + v96*int32(100)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+119)))
	if v121 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v254 = v234
	goto L19
L22:
	;
	v240 = v96 + int32(1)
	if v240 != v78 {
		v96 = v240
		v101 = v233
		v102 = v234
		goto L20
	} else {
		goto L52
	}
L23:
	;
	v125 = F_lappend(m, v102, int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v128 = v120 + int32(28)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+96))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v128)+76))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v128)+68))
	if l3 == l5 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v233 = v101
	v234 = v125
	goto L22
L27:
	;
	v134 = v96 + int32(1)
	v137 = F_makeVar(m, v51, base.I32_extend16_s(v134), v131, v130, v129, int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v146 = v120 + int32(32)
	if v80 <= v101 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v139 = F_lappend(m, v102, v137)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v84+v96<<(uint(int32(1))%32)))) = uint16(v134)
	v233 = v101
	v234 = v139
	goto L22
L32:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v210)+68))
	if v131 != v212 {
		goto L15
	} else {
		goto L47
	}
L33:
	;
	v189 = F_SearchSysCacheAttName(m, v76, v146)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L44
	}
L34:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v154 = v79 + v148<<(uint(int32(3))%32) + v101*int32(100)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+119)))
	if v155 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v159 = v154 + int32(32)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if base.B2i32(v162 == int32(0))|base.B2i32(v162 != v165) != 0 {
		v183 = v162
		v184 = v165
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v183-v184 == int32(0) {
		v210 = v154 + int32(28)
		v211 = v101
		goto L32
	} else {
		goto L43
	}
L37:
	;
	goto L36
L38:
	;
	v168 = v146
	v169 = v159
	goto L39
L39:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+1)))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	if v173 == int32(0) {
		v183 = v173
		v184 = v172
		goto L37
	} else {
		goto L41
	}
L40:
	;
	v183 = v173
	v184 = v172
	goto L37
L41:
	;
	v176 = int32(1)
	if v173 == v172 {
		v168 = v168 + v176
		v169 = v169 + v176
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	goto L33
L44:
	;
	if v189 == int32(0) {
		goto L16
	} else {
		goto L45
	}
L45:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189)+16))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+22)))
	v196 = int32(*(*int16)(unsafe.Add(mBase, uint32(v193+v194)+74)))
	F_ReleaseCatCache(m, v189)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v210 = v79 + v199<<(uint(int32(3))%32) + v196*int32(100) - int32(72)
	v211 = v196 - int32(1)
	goto L32
L47:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210)+76))
	if v130 != v214 {
		goto L15
	} else {
		goto L48
	}
L48:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v210)+96))
	if v129 != v216 {
		goto L14
	} else {
		goto L49
	}
L49:
	;
	v218 = int32(1)
	v221 = v211 + v218
	v224 = F_makeVar(m, v51, base.I32_extend16_s(v221), v131, v130, v129, int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v226 = F_lappend(m, v102, v224)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v230 = v96 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v211<<(uint(v218)%32)+v84))) = uint16(v230)
	v233 = v221
	v234 = v226
	goto L22
L52:
	;
	goto L21
L53:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v276 + int32(4)
	F_errmsg_internal(m, int32(_a_F_expand_single_inheritance_child_0), v61)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_expand_single_inheritance_child_1), int32(154), int32(_a_F_expand_single_inheritance_child_2))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
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
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+32)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v61)+36)) = v296 + int32(4)
	F_errmsg(m, int32(_a_F_expand_single_inheritance_child_3), v61+int32(32))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_expand_single_inheritance_child_1), int32(167), int32(_a_F_expand_single_inheritance_child_2))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
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
	F_errcode(m, int32(17064068))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v61)+20)) = v318 + int32(4)
	F_errmsg(m, int32(_a_F_expand_single_inheritance_child_4), v61+int32(16))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_expand_single_inheritance_child_1), int32(172), int32(_a_F_expand_single_inheritance_child_2))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v334
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v338 = F_copyObjectImpl(m, v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v338
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l5)+52))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	if v342 <= int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	v442 = F_makeAlias(m, v441, v422)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L81
	}
L67:
	;
	v422 = int32(0)
	goto L66
L68:
	;
	goto L69
L69:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v346)+8))
	v348 = int32(0)
	v353 = v348
	v356 = v348
	v357 = v342
	goto L70
L70:
	;
	v380 = v341 + v357<<(uint(int32(3))%32) + v353*int32(100)
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+119)))
	if v381 != 0 {
		v405 = int32(_a_F_expand_single_inheritance_child_5)
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v422 = v410
	goto L66
L72:
	;
	v406 = F_pstrdup(m, v405)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L77
	}
L73:
	;
	v382 = int32(0)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	v388 = int32(*(*int16)(unsafe.Add(mBase, uint32(v384+v353<<(uint(int32(1))%32)))))
	if base.B2i32(v347 == v382)|base.B2i32(v388 <= v382) != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v405 = v380 + int32(32)
	goto L72
L75:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	if v392 < v388 {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v394+v388<<(uint(int32(2))%32)-int32(4))))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+4))
	v405 = v401
	goto L72
L77:
	;
	v408 = F_makeString(m, v406)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v410 = F_lappend(m, v356, v408)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v413 = v353 + int32(1)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	if v413 < v414 {
		v353 = v413
		v356 = v410
		v357 = v414
		goto L70
	} else {
		goto L80
	}
L80:
	;
	goto L71
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v442
	v447 = v51 << (uint(int32(2)) % 32)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v447+v448))) = v29
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v451+v447))) = v64
	if l4 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v455 = F_palloc0(m, int32(36))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v529 = F_bms_is_member(m, l2, v528)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L102
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v455)+4)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v455))) = int32(378)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v455)+8)) = v460
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v455)+12)) = v462
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
	v465 = m.G0
	v467 = v465 - int32(16)
	m.G0 = v467
	v469 = int32(5)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if v470 != 0 {
		v486 = v469
		goto L88
	} else {
		goto L89
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v455)+16)) = v486
	v506 = int32(1) << (uint(v486) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v455)+20)) = v506
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v455)+24)) = v508
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v455)+28)) = v510
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v455)+32)) = uint8(base.B2i32(v512 == int32(112)))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v516 | v506
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v520 = F_lappend(m, v519, v455)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L100
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L97
	}
L88:
	;
	m.G0 = v467 + int32(16)
	goto L86
L89:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+21)))
	if v471 == int32(102) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v475 = F_GetFdwRoutineByRelId(m, v474)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	if base.Ui32(int32(5)) <= base.Ui32(v464) {
		goto L87
	} else {
		goto L96
	}
L93:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v475)+104))
	if v477 == int32(0) {
		v486 = v469
		goto L88
	} else {
		goto L94
	}
L94:
	;
	v480 = m.T0[v477].(func(*base.Module, int32, int32) int32)(m, v29, v464)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v486 = v480
	goto L88
L96:
	;
	v486 = int32(4) - v464
	goto L88
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v467))) = v464
	F_errmsg_internal(m, int32(_a_F_expand_single_inheritance_child_6), v467)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_expand_single_inheritance_child_7), int32(2855), int32(_a_F_expand_single_inheritance_child_8))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v520
	goto L84
L101:
	;
	return
L102:
	;
	if v529 == int32(0) {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v534 = F_bms_add_member(m, v533, v51)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v534
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+21)))
	if v537 == int32(112) {
		goto L101
	} else {
		goto L105
	}
L105:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v541 = F_bms_add_member(m, v540, v51)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v541
	v547 = int32(0)
	v549 = F_makeVar(m, v51, int32(-6), int32(26), int32(-1), v547, v547)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_add_row_identity_var(m, l0, v549, v51, int32(_a_F_expand_single_inheritance_child_9))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_add_row_identity_columns(m, l0, v51, v29, l5)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	goto L101
}
func F_extractNotNullColumn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v4 = F_SysCacheGetAttrNotNull(m, int32(19), l0, int32(21))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = F_pg_detoast_datum(m, base.I32_wrap_i64(v4))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			if v11 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_extractNotNullColumn_0), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_extractNotNullColumn_1), int32(720), int32(_a_F_extractNotNullColumn_2))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				if v14 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_extractNotNullColumn_0), int32(0))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_extractNotNullColumn_1), int32(720), int32(_a_F_extractNotNullColumn_2))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					if v15 != int32(21) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_extractNotNullColumn_0), int32(0))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_extractNotNullColumn_1), int32(720), int32(_a_F_extractNotNullColumn_2))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
						if v18 == int32(1) {
							v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+24)))
							return v34
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_extractNotNullColumn_0), int32(0))
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_extractNotNullColumn_1), int32(720), int32(_a_F_extractNotNullColumn_2))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
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
