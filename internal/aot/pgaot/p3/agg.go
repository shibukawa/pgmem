package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_agg_clause_costs(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v83 float64
	_ = v83
	var v84 float64
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v96 float64
	_ = v96
	var v97 float64
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 float64
	_ = v199
	var v200 float64
	_ = v200
	var v203 float64
	_ = v203
	var v204 float64
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	if v18 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	if v157 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v21 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = l1 & int32(1)
	if v33 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v34 = int32(24)
	goto L6
L5:
	;
	v34 = int32(12)
	goto L6
L6:
	;
	v41 = v4
	goto L7
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v41<<(uint(int32(2))%32))))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52+v34)))
	F_add_function_cost(m, l0, v54, int32(0), l2)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L1
L9:
	;
	return
L10:
	;
	if l1&int32(8) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if l1&int32(4) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	if v60 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	F_add_function_cost(m, l0, v60, int32(0), l2)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	if v33 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v69 == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_add_function_cost(m, l0, v69, int32(0), l2+int32(16))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+40)))
	if v101 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	F_cost_qual_eval_node(m, v16, v76, l0)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	v79 = *(*float64)(unsafe.Add(mBase, uint32(v16)))
	v80 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	*(*float64)(unsafe.Add(mBase, uint32(l2))) = base.F64_add(v79, v80)
	v83 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
	v84 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l2)+8)) = base.F64_add(v83, v84)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v87 == int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	F_cost_qual_eval_node(m, v16, v87, l0)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	v92 = *(*float64)(unsafe.Add(mBase, uint32(v16)))
	v93 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	*(*float64)(unsafe.Add(mBase, uint32(l2))) = base.F64_add(v92, v93)
	v96 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
	v97 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l2)+8)) = base.F64_add(v96, v97)
	goto L19
L24:
	;
	v141 = v41 + int32(1)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v141 < v142 {
		v41 = v141
		goto L7
	} else {
		goto L37
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v136
	goto L24
L26:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v52)+44))
	if int32(0) < v104 {
		v115 = v104
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v52)+28))
	if v124 != int32(2281) {
		goto L24
	} else {
		goto L33
	}
L29:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v136 = v116 + (v115+int32(7))&int32(-8) + int32(8)
	goto L25
L30:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	if v108 == int32(378) {
		v115 = int32(1024)
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v52)+28))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v52)+32))
	v113 = F_get_typavgwidth(m, v111, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	v115 = v113
	goto L29
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v52)+44))
	if int32(0) < v128 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v136 = v128 + v127
	goto L25
L35:
	;
	goto L36
L36:
	;
	v136 = v127 - int32(-8192)
	goto L25
L37:
	;
	goto L8
L38:
	;
	m.G0 = v16 + int32(16)
	return
L39:
	;
	v160 = int32(0)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v161 <= v160 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v175 = v160
	goto L41
L41:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v181+v175<<(uint(int32(2))%32))))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	if l1&int32(2) != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L38
L43:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v188)+28))
	if v196 != 0 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v185)+16))
	if v189 == int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	F_add_function_cost(m, l0, v189, int32(0), l2+int32(16))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L9
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	F_cost_qual_eval_node(m, v16, v196, l0)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L9
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v208 = v175 + int32(1)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v208 < v209 {
		v175 = v208
		goto L41
	} else {
		goto L51
	}
L50:
	;
	v199 = *(*float64)(unsafe.Add(mBase, uint32(v16)))
	v200 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l2)+16)) = base.F64_add(v199, v200)
	v203 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
	v204 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
	*(*float64)(unsafe.Add(mBase, uint32(l2)+24)) = base.F64_add(v203, v204)
	goto L49
L51:
	;
	goto L42
}
func F_get_agg_expr_helper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	v7 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(416)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)) = uint8(v7)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v21&int32(1) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(416)
	return
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	F_resolve_special_varno(m, v27, l1, int32(1700), l2)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+56)))
	if v31&int32(2) != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return
L6:
	;
	goto L1
L7:
	;
	F_appendStringInfoString(m, v18, int32(_a_F_get_agg_expr_helper_0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v38 = v16 + int32(16)
	v39 = F_get_aggregate_argtypes(m, l0, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	goto L9
L11:
	;
	if l3 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+49)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+34)))
	v49 = F_generate_function_name(m, v43, v39, int32(0), v38, v45, v16+int32(15), v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L15
	}
L13:
	;
	v51 = l3
	goto L14
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v51
	if v52 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v51 = v49
	goto L14
L16:
	;
	v56 = int32(_a_F_get_agg_expr_helper_1)
	goto L18
L17:
	;
	v56 = int32(_a_F_get_agg_expr_helper_2)
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v56
	F_appendStringInfo(m, v18, int32(_a_F_get_agg_expr_helper_3), v16)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v61 != int32(110) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if l4 != 0 {
		goto L56
	} else {
		goto L57
	}
L21:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_get_rule_orderby(m, v170, v171, int32(0), l1)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L5
	} else {
		goto L55
	}
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_get_rule_expr(m, v64, l1, int32(1))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v71 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	F_appendStringInfoString(m, v18, int32(_a_F_get_agg_expr_helper_4))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	goto L21
L27:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v151 == int32(0) {
		goto L20
	} else {
		goto L53
	}
L28:
	;
	F_appendStringInfoChar(m, v18, int32(42))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v77 == int32(0) {
		goto L27
	} else {
		goto L32
	}
L31:
	;
	goto L27
L32:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v80 <= int32(0) {
		goto L27
	} else {
		goto L33
	}
L33:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
	v90 = int32(0)
	v96 = v7
	goto L34
L34:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100+v90<<(uint(int32(2))%32))))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+26)))
	if v105 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L27
L36:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v110 = v96 + int32(1)
	if int32(0) < v96 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v133 = v96
	goto L38
L38:
	;
	v135 = v90 + int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v135 < v136 {
		v90 = v135
		v96 = v133
		goto L34
	} else {
		goto L52
	}
L39:
	;
	if l5 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v120 = int32(0)
	if base.B2i32(v83&int32(1) == v120)|base.B2i32(v110 != v39) == v120 {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	if int32(2) < v110 {
		goto L27
	} else {
		goto L45
	}
L43:
	;
	v117 = int32(_a_F_get_agg_expr_helper_5)
	goto L44
L44:
	;
	F_appendStringInfoString(m, v18, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L46
	}
L45:
	;
	v117 = int32(_a_F_get_agg_expr_helper_6)
	goto L44
L46:
	;
	goto L41
L47:
	;
	F_appendStringInfoString(m, v18, int32(_a_F_get_agg_expr_helper_7))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L5
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	F_get_rule_expr(m, v108, l1, int32(1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L5
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	v133 = v110
	goto L38
L52:
	;
	goto L35
L53:
	;
	F_appendStringInfoString(m, v18, int32(_a_F_get_agg_expr_helper_8))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	goto L21
L55:
	;
	goto L20
L56:
	;
	F_appendStringInfoString(m, v18, l4)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L5
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v190 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L58
L60:
	;
	F_appendStringInfoString(m, v18, int32(_a_F_get_agg_expr_helper_9))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L5
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_appendStringInfoChar(m, v18, int32(41))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L5
	} else {
		goto L65
	}
L63:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_get_rule_expr(m, v194, l1, int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	goto L1
}
