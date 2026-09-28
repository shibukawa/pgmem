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
	v4 = Fn14250(m, l0, l1, int32(1121))
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
	var v29 int32
	_ = v29
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
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
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v8 - int32(29) {
		case 0:
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v30 = F_contain_context_dependent_node_walker(m, v29, l1)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				if v30 != 0 {
					return int32(1)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v32 | int32(1)
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v37 = F_contain_context_dependent_node_walker(m, v36, l1)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v58 = v37
						v59 = v32
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v59
						return v58
					}
				}
			}
		default:
			v53 = F_expression_tree_walker_impl(m, l0, int32(926), l1)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				return v53
			}
		case 3:
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v17 == int32(0) {
				v53 = F_expression_tree_walker_impl(m, l0, int32(926), l1)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					return v53
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v20 | int32(1)
				v25 = F_expression_tree_walker_impl(m, l0, int32(926), l1)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v58 = v25
					v59 = v20
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v59
					return v58
				}
			}
		case 5:
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			return base.B2i32(v11&int32(1) == int32(0))
		case 16:
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v40 = F_contain_context_dependent_node_walker(m, v39, l1)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				if v40 != 0 {
					return int32(1)
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v43 = F_contain_context_dependent_node_walker(m, v42, l1)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						if v43 != 0 {
							return int32(1)
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v45 | int32(1)
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v50 = F_contain_context_dependent_node_walker(m, v49, l1)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v58 = v50
								v59 = v45
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v59
								return v58
							}
						}
					}
				}
			}
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
	v57 = F_expression_tree_walker_impl(m, l0, int32(916), l1)
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
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
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	v3 = int32(0)
	if l0 == v3 {
		v189 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v194 = F_query_tree_walker_impl(m, l0, int32(907), l1, int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L46
	}
L2:
	;
	return v189
L3:
	;
	v12 = F_check_functions_in_node(m, l0, int32(906), l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if v12 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(1)
L7:
	;
	goto L8
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v18 == int32(45) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v181 = F_expression_tree_walker_impl(m, l0, int32(907), l1)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L45
	}
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v21 == int32(0) {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	v108 = v18
	goto L12
L12:
	;
	if v108 == int32(48) {
		goto L29
	} else {
		goto L30
	}
L13:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if int32(0) < v24 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v35 = int32(0)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v108 = v105
	goto L12
L17:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v35<<(uint(int32(2))%32))))
	v46 = F_exprType(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L16
L19:
	;
	if base.B2i32(v30 != int32(2)) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v94 = v35 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v94 < v95 {
		v35 = v94
		goto L17
	} else {
		goto L28
	}
L21:
	;
	v50 = m.G0
	v52 = v50 - int32(16)
	m.G0 = v52
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+15)) = uint8(v54)
	F_json_check_mutability(m, v46, int32(1), v52+int32(15))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v71 = m.G0
	v73 = v71 - int32(16)
	m.G0 = v73
	v75 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+15)) = uint8(v75)
	F_json_check_mutability(m, v46, v75, v73+int32(15))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+15)))
	m.G0 = v52 + int32(16)
	if (v61^int32(-1))&int32(1) != 0 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	return int32(1)
L26:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+15)))
	m.G0 = v73 + int32(16)
	if (v82^int32(-1))&int32(1) != 0 {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	return int32(1)
L28:
	;
	goto L18
L29:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v117 != int32(7) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v160 = v108
	goto L31
L31:
	;
	v164 = int32(1)
	switch v160 - int32(40) {
	case 0, 19:
		v189 = v164
		goto L2
	case 1:
		goto L42
	default:
		goto L9
	case 27:
		goto L1
	}
L32:
	;
	return int32(1)
L33:
	;
	goto L34
L34:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+32)))
	if v122 != 0 {
		v189 = v3
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v116)+24))
	v124 = F_pg_detoast_datum(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v128 = m.G0
	v130 = v128 - int32(48)
	m.G0 = v130
	v132 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v130)+40)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v130)+36)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v130)+32)) = v126
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v130)+45)) = uint8(v132)
	v140 = int32(base.Ui32(v136) >> (uint(int32(31)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v130)+44)) = uint8(v140)
	v143 = v130 + int32(4)
	F_jspInitByBuffer(m, v143, v124+int32(8), v132)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v151 = F_jspIsMutableWalker(m, v143, v130+int32(32))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+45)))
	m.G0 = v130 + int32(48)
	if v153 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	return int32(1)
L40:
	;
	goto L41
L41:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v160 = v159
	goto L31
L42:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v167-int32(3)) < base.Ui32(int32(5)) {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	if v167 != 0 {
		v189 = v164
		goto L2
	} else {
		goto L44
	}
L44:
	;
	goto L9
L45:
	;
	v189 = v181
	goto L2
L46:
	;
	return v194
}
