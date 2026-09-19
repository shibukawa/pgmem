package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_contain_agg_clause_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(v7-int32(9)) < base.Ui32(int32(2)) {
			return int32(1)
		} else {
			v15 = F_expression_tree_walker_impl(m, l0, int32(855), l1)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_contain_leaked_vars_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
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
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
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
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	v3 = int32(0)
	if l0 == v3 {
		v124 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v141 = F_expression_tree_walker_impl(m, l0, int32(870), l1)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L9
	} else {
		goto L52
	}
L2:
	;
	return v124
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v12 - int32(1) {
	case 0, 5, 6, 7, 15, 20, 24, 25, 26, 30, 31, 33, 34, 35, 39, 51, 52, 58, 60:
		goto L1
	default:
		goto L4
	case 13:
		goto L7
	case 14, 16, 17, 18, 19, 27, 28:
		goto L8
	case 36:
		goto L6
	case 38:
		goto L5
	case 57:
		v124 = v3
		goto L2
	}
L4:
	;
	v124 = int32(1)
	goto L2
L5:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v109 = F_lookup_type_cache(m, v107, int32(8))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L9
	} else {
		goto L44
	}
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v46 = v3
	goto L24
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = F_getSubscriptingRoutines(m, v26, int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L9
	} else {
		goto L15
	}
L8:
	;
	v16 = F_check_functions_in_node(m, l0, int32(869), l1)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	if v16 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v22 = F_contain_var_clause(m, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	if v22 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L4
L14:
	;
	v37 = F_contain_var_clause(m, l0)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L9
	} else {
		goto L22
	}
L15:
	;
	if v28 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v32 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+10)))
	if v33 == int32(0) {
		goto L14
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+9)))
	if v36 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L1
L21:
	;
	goto L14
L22:
	;
	if v37 == int32(0) {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	goto L4
L24:
	;
	v53 = int32(0)
	if v43 == v53 {
		v63 = v53
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v64 = int32(0)
	if v42 == v64 {
		v73 = v64
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v57 <= v46 {
		v63 = int32(0)
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v63 = v59 + v46<<(uint(int32(2))%32)
	goto L26
L29:
	;
	if v41 == int32(0) {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v67 <= v46 {
		v73 = v64
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v73 = v69 + v46<<(uint(int32(2))%32)
	goto L29
L32:
	;
	v76 = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if base.B2i32(v73 == v76)|(base.B2i32(v63 == v76)|base.B2i32(v80 <= v46)) != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	if v84 == int32(0) {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v88 = F_get_opcode(m, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	v90 = F_get_func_leakproof(m, v88)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	if v90 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v94 = int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v96 = F_contain_var_clause(m, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L9
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v46 = v46 + int32(1)
	goto L24
L40:
	;
	if v96 != 0 {
		v124 = v94
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v84+v46<<(uint(int32(2))%32))))
	v102 = F_contain_var_clause(m, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L9
	} else {
		goto L42
	}
L42:
	;
	if v102 != 0 {
		v124 = v94
		goto L2
	} else {
		goto L43
	}
L43:
	;
	goto L39
L44:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)+64))
	if v111 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v112 = F_get_func_leakproof(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L9
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v115 = F_contain_var_clause(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L9
	} else {
		goto L50
	}
L48:
	;
	if v112 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	if v115 == int32(0) {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L4
L52:
	;
	return v141
}
func F_contain_mutable_functions(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_contain_mutable_functions_walker(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_contain_mutable_functions_checker(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_func_volatile(m, l0)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v3 != int32(105))
	}
}
func F_contain_mutable_or_user_functions_checker(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_func_volatile(m, l0)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v3 != int32(105)) | base.B2i32(base.Ui32(int32(_a_F_contain_mutable_or_user_functions_checker_0)) < base.Ui32(l0))
	}
}
func F_contain_nonstrict_functions_checker(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_func_strict(m, l0)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3 ^ int32(1)
	}
}
func F_contain_outer_selfref_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v3 = int32(0)
	if l0 == v3 {
		v42 = v3
		return v42
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v6 != int32(67) {
			if v6 != int32(101) {
				v40 = F_expression_tree_walker_impl(m, l0, int32(843), l1)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = v40
					return v42
				}
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v11 != int32(6) {
					return int32(0)
				} else {
					v14 = int32(1)
					v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
					if v15 != v14 {
						return int32(0)
					} else {
						v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if base.Ui32(v19) <= base.Ui32(v18) {
							v42 = v14
							return v42
						} else {
							return int32(0)
						}
					}
				}
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v24 + int32(1)
			v30 = F_query_tree_walker_impl(m, l0, int32(843), l1, int32(16))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v34 - int32(1)
				return v30
			}
		}
	}
}
func F_contain_var_clause_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	if l0 == int32(0) {
		return int32(0)
	} else {
		v9 = int32(1)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v10 != int32(319) {
			if v10 == int32(58) {
				v29 = v9
				return v29
			} else {
				if v10 != int32(6) {
					v25 = F_expression_tree_walker_impl(m, l0, int32(899), l1)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = v25
						return v29
					}
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					return base.B2i32(v17 == int32(0))
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v21 == int32(0) {
				v29 = v9
				return v29
			} else {
				v25 = F_expression_tree_walker_impl(m, l0, int32(899), l1)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = v25
					return v29
				}
			}
		}
	}
}
func F_contain_vars_of_level_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v8 - int32(58) {
		case 0:
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			return base.B2i32(v19 == int32(0))
		case 1, 2, 3, 4, 5, 6, 7, 8:
			v44 = F_expression_tree_walker_impl(m, l0, int32(900), l1)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				return v44
			}
		case 9:
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v28 + int32(1)
			v34 = F_query_tree_walker_impl(m, l0, int32(900), l1, int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v38 - int32(1)
				return v34
			}
		default:
			if v8 == int32(319) {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if v23 != v24 {
					v44 = F_expression_tree_walker_impl(m, l0, int32(900), l1)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						return v44
					}
				} else {
					return int32(1)
				}
			} else {
				if v8 != int32(6) {
					v44 = F_expression_tree_walker_impl(m, l0, int32(900), l1)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						return v44
					}
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					return base.B2i32(v15 == v16)
				}
			}
		}
	}
}
func F_contain_volatile_functions(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_contain_volatile_functions_walker(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_contain_windowfuncs_walker(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13868(m, l0, l1, int32(1045), int32(11))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
