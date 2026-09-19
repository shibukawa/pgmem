package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_contain_aggs_of_level(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13867(m, l0, l1, int32(1043))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_contain_context_dependent_node_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v8 - int32(29) {
		case 0:
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v32 = F_contain_context_dependent_node_walker(m, v31, l1)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				if v32 != 0 {
					return int32(1)
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v36 | int32(1)
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v41 = F_contain_context_dependent_node_walker(m, v40, l1)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v36
						return v41
					}
				}
			}
		default:
			v46 = F_expression_tree_walker_impl(m, l0, int32(877), l1)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				return v46
			}
		case 3:
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v17 == int32(0) {
				v46 = F_expression_tree_walker_impl(m, l0, int32(877), l1)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					return v46
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v20 | int32(1)
				v25 = F_expression_tree_walker_impl(m, l0, int32(877), l1)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v20
					return v25
				}
			}
		case 5:
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			return base.B2i32(v11&int32(1) == int32(0))
		}
	}
}
func F_contain_exec_param(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 != int32(8) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v61
L5:
	;
	v57 = F_expression_tree_walker_impl(m, l0, int32(868), l1)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L22
	} else {
		goto L23
	}
L6:
	;
	v11 = int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 != v11 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = int32(0)
	if l1 == v16 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v54 != 0 {
		v61 = v11
		goto L4
	} else {
		goto L21
	}
L9:
	;
	v54 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v22 <= int32(0) {
		v48 = v16
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v54 = v48
	goto L8
L13:
	;
	v25 = int32(0)
	if v25 < v22 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v28 = v22
	goto L16
L15:
	;
	v28 = v25
	goto L16
L16:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v31 = int32(0)
	goto L17
L17:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v29+v31<<(uint(int32(2))%32))))
	v40 = base.B2i32(v39 == v15)
	if v39 == v15 {
		v48 = v40
		goto L12
	} else {
		goto L19
	}
L18:
	;
	v48 = v40
	goto L12
L19:
	;
	v42 = v31 + int32(1)
	if v42 != v28 {
		v31 = v42
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L5
L22:
	;
	return int32(0)
L23:
	;
	v61 = v57
	goto L4
}
func F_contain_leaked_vars(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_contain_leaked_vars_walker(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_contain_mutable_functions_after_planning(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = F_expression_planner(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		v7 = F_contain_mutable_functions_walker(m, v2, int32(0))
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_contain_mutable_functions_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
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
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	v3 = int32(0)
	if l0 == v3 {
		v229 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v229
L2:
	;
	v13 = F_check_functions_in_node(m, l0, int32(858), l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v13 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(1)
L6:
	;
	goto L7
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v19 == int32(45) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v220 = F_expression_tree_walker_impl(m, l0, int32(859), l1)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L3
	} else {
		goto L55
	}
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v22 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	v145 = v19
	goto L11
L11:
	;
	if v145 == int32(48) {
		goto L40
	} else {
		goto L41
	}
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if int32(0) < v25 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v36 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v145 = v142
	goto L11
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v36<<(uint(int32(2))%32))))
	v48 = F_exprType(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L18
	}
L17:
	;
	goto L15
L18:
	;
	if base.B2i32(v31 != int32(2)) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v130 = v36 + int32(1)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v130 < v131 {
		v36 = v130
		goto L16
	} else {
		goto L39
	}
L20:
	;
	v52 = m.G0
	v54 = v52 - int32(16)
	m.G0 = v54
	v56 = int32(1)
	F_json_categorize_type(m, v48, v56, v54+int32(12), v54+int32(8))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L3
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v89 = m.G0
	v91 = v89 - int32(16)
	m.G0 = v91
	F_json_categorize_type(m, v48, int32(0), v91+int32(12), v91+int32(8))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L3
	} else {
		goto L31
	}
L23:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	if base.Ui32(int32(11)) < base.Ui32(v64) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	m.G0 = v54 + int32(16)
	if v82 != 0 {
		goto L19
	} else {
		goto L30
	}
L25:
	;
	v82 = int32(0)
	goto L24
L26:
	;
	v68 = int32(1) << (uint(v64) % 32)
	if v68&int32(195) != 0 {
		v82 = v56
		goto L24
	} else {
		goto L27
	}
L27:
	;
	if v68&int32(3076) == int32(0) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v76 = F_func_volatile(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	v82 = base.B2i32(v76 == int32(105))
	goto L24
L30:
	;
	return int32(1)
L31:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if base.Ui32(int32(11)) < base.Ui32(v100) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	m.G0 = v91 + int32(16)
	if v120 != 0 {
		goto L19
	} else {
		goto L38
	}
L33:
	;
	v120 = int32(0)
	goto L32
L34:
	;
	v103 = int32(1)
	v105 = v103 << (uint(v100) % 32)
	if v105&int32(195) != 0 {
		v120 = v103
		goto L32
	} else {
		goto L35
	}
L35:
	;
	if v105&int32(3076) == int32(0) {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v113 = F_func_volatile(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v120 = base.B2i32(v113 == int32(105))
	goto L32
L38:
	;
	return int32(1)
L39:
	;
	goto L17
L40:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v155 != int32(7) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v198 = v145
	goto L42
L42:
	;
	switch v198 - int32(40) {
	case 0, 19:
		v229 = int32(1)
		goto L1
	default:
		goto L8
	case 27:
		goto L53
	}
L43:
	;
	return int32(1)
L44:
	;
	goto L45
L45:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+24)))
	if v160 != 0 {
		v229 = v3
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v154)+20))
	v162 = F_pg_detoast_datum(m, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v166 = m.G0
	v168 = v166 - int32(48)
	m.G0 = v168
	v170 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v168)+40)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v168)+36)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v168)+32)) = v164
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v168)+45)) = uint8(v170)
	v178 = int32(base.Ui32(v174) >> (uint(int32(31)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v168)+44)) = uint8(v178)
	v181 = v168 + int32(4)
	F_jspInitByBuffer(m, v181, v162+int32(8), v170)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	v189 = F_jspIsMutableWalker(m, v181, v168+int32(32))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+45)))
	m.G0 = v168 + int32(48)
	if v191 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	return int32(1)
L51:
	;
	goto L52
L52:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v198 = v197
	goto L42
L53:
	;
	v207 = F_query_tree_walker_impl(m, l0, int32(859), l1, int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	return v207
L55:
	;
	v229 = v220
	goto L1
}
