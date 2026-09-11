package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_contain_aggs_of_level(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l1
	v13 = F_query_or_expression_tree_walker_impl(m, l0, int32(1042), v6+int32(12), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v13
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
			v46 = F_expression_tree_walker_impl(m, l0, int32(876), l1)
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
				v46 = F_expression_tree_walker_impl(m, l0, int32(876), l1)
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
				v25 = F_expression_tree_walker_impl(m, l0, int32(876), l1)
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
	var v47 int32
	_ = v47
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
	v57 = F_expression_tree_walker_impl(m, l0, int32(867), l1)
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
		v47 = v16
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v54 = v47
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
		v47 = v40
		goto L12
	} else {
		goto L19
	}
L18:
	;
	v47 = v40
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
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
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v248 int32
	_ = v248
	v3 = int32(0)
	if l0 == v3 {
		v248 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v248
L2:
	;
	v14 = F_check_functions_in_node(m, l0, int32(857), l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v14 != 0 {
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
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v20 == int32(45) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v238 = F_expression_tree_walker_impl(m, l0, int32(858), l1)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L3
	} else {
		goto L60
	}
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v23 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	v160 = v20
	goto L11
L11:
	;
	if v160 == int32(48) {
		goto L45
	} else {
		goto L46
	}
L12:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if int32(0) < v26 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v38 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v160 = v156
	goto L11
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v38<<(uint(int32(2))%32))))
	v50 = F_exprType(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L3
	} else {
		goto L18
	}
L17:
	;
	goto L15
L18:
	;
	if base.B2i32(v32 != int32(2)) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v143 = v38 + int32(1)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v143 < v144 {
		v38 = v143
		goto L16
	} else {
		goto L44
	}
L20:
	;
	v54 = m.G0
	v56 = v54 - int32(16)
	m.G0 = v56
	v58 = int32(1)
	F_json_categorize_type(m, v50, v58, v56+int32(12), v56+int32(8))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v95 = int32(0)
	v96 = m.G0
	v98 = v96 - int32(16)
	m.G0 = v98
	F_json_categorize_type(m, v50, v95, v98+int32(12), v98+int32(8))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L3
	} else {
		goto L32
	}
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if base.Ui32(int32(11)) < base.Ui32(v66) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	m.G0 = v56 + int32(16)
	if v88 != 0 {
		goto L19
	} else {
		goto L31
	}
L25:
	;
	v88 = int32(0)
	goto L24
L26:
	;
	v70 = int32(1) << (uint(v66) % 32)
	if v70&int32(195) != 0 {
		v88 = v58
		goto L24
	} else {
		goto L27
	}
L27:
	;
	if v70&int32(56) != 0 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	if int32(1)<<(uint(v66)%32)&int32(3076) == int32(0) {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v82 = F_func_volatile(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	v88 = base.B2i32(v82 == int32(105))
	goto L24
L31:
	;
	return int32(1)
L32:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	if base.Ui32(int32(11)) < base.Ui32(v107) {
		v131 = v95
		goto L33
	} else {
		goto L34
	}
L33:
	;
	m.G0 = v98 + int32(16)
	if v131 != 0 {
		goto L19
	} else {
		goto L43
	}
L34:
	;
	v111 = int32(1) << (uint(v107) % 32)
	if v111&int32(195) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v131 = int32(1)
	goto L33
L36:
	;
	goto L37
L37:
	;
	if v111&int32(56) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if int32(1)<<(uint(v107)%32)&int32(3076) == int32(0) {
		v131 = v95
		goto L33
	} else {
		goto L41
	}
L39:
	;
	v130 = v95
	goto L40
L40:
	;
	v131 = v130
	goto L33
L41:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	v126 = F_func_volatile(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	v130 = base.B2i32(v126 == int32(105))
	goto L40
L43:
	;
	return int32(1)
L44:
	;
	goto L17
L45:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	if v170 != int32(7) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v216 = v160
	goto L47
L47:
	;
	switch v216 - int32(40) {
	case 0, 19:
		v248 = int32(1)
		goto L1
	default:
		goto L8
	case 27:
		goto L58
	}
L48:
	;
	return int32(1)
L49:
	;
	goto L50
L50:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+24)))
	if v175 != 0 {
		v248 = v3
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v169)+20))
	v177 = F_pg_detoast_datum(m, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v181 = m.G0
	v183 = v181 - int32(48)
	m.G0 = v183
	v185 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v183)+40)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v183)+36)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v183)+32)) = v179
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v183)+45)) = uint8(v185)
	v193 = int32(base.Ui32(v189) >> (uint(int32(31)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v183)+44)) = uint8(v193)
	F_jspInitByBuffer(m, v183+int32(4), v177+int32(8), v185)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	v206 = F_jspIsMutableWalker(m, v183+int32(4), v183+int32(32))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+45)))
	m.G0 = v183 + int32(48)
	if v208 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	return int32(1)
L56:
	;
	goto L57
L57:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v216 = v214
	goto L47
L58:
	;
	v224 = F_query_tree_walker_impl(m, l0, int32(858), l1, int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	return v224
L60:
	;
	v248 = v238
	goto L1
}
