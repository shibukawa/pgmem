package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_contain_agg_clause(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	if l0 == int32(0) {
		return int32(0)
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(v6-int32(9)) < base.Ui32(int32(2)) {
			return int32(1)
		} else {
			v15 = F_expression_tree_walker_impl(m, l0, int32(855), int32(0))
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
func F_contain_nonstrict_functions_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	v9 = l0
	goto L6
L4:
	;
	return v55
L5:
	;
	v50 = F_check_functions_in_node(m, v9, int32(866), l1)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L13
	} else {
		goto L25
	}
L6:
	;
	v13 = int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	switch v14 - int32(9) {
	case 0, 1, 2:
		v55 = v13
		goto L4
	default:
		v34 = v14
		goto L8
	case 5:
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v34-int32(28)) {
		goto L19
	} else {
		goto L20
	}
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	if v17 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(1)
L11:
	;
	goto L12
L12:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v22 = F_getSubscriptingRoutines(m, v20, int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	if v22 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(1)
L16:
	;
	goto L17
L17:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
	if v30 != int32(1) {
		v55 = v13
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v34 = v33
	goto L8
L19:
	;
	switch v34 - int32(18) {
	case 0, 1, 4, 5, 6, 8, 14, 17, 18, 19, 20, 21, 23, 27, 34, 35:
		v55 = v13
		goto L4
	default:
		goto L5
	case 3:
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v46 != 0 {
		v9 = v46
		goto L6
	} else {
		goto L24
	}
L22:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if base.Ui32(int32(2)) <= base.Ui32(v41) {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	return int32(1)
L24:
	;
	goto L7
L25:
	;
	if v50 != 0 {
		v55 = v13
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v53 = F_expression_tree_walker_impl(m, v9, int32(867), l1)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	v55 = v53
	goto L4
}
func F_contain_var_clause(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = int32(1)
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v9 != int32(319) {
			if v9 == int32(58) {
				v29 = v8
				return v29
			} else {
				if v9 != int32(6) {
					v25 = F_expression_tree_walker_impl(m, l0, int32(899), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = v25
						return v29
					}
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					return base.B2i32(v16 == int32(0))
				}
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v20 == int32(0) {
				v29 = v8
				return v29
			} else {
				v25 = F_expression_tree_walker_impl(m, l0, int32(899), int32(0))
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
func F_contain_volatile_functions_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	v3 = int32(0)
	if l0 == v3 {
		v55 = v3
		return v55
	} else {
		v7 = int32(1)
		v9 = F_check_functions_in_node(m, l0, int32(860), l1)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if v9 != 0 {
				v55 = v7
				return v55
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				switch v13 - int32(59) {
				case 0:
					v55 = v7
					return v55
				case 1, 2, 3, 4, 5, 6, 7:
					v50 = F_expression_tree_walker_impl(m, l0, int32(861), l1)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						return v50
					}
				case 8:
					v46 = F_query_tree_walker_impl(m, l0, int32(861), l1, int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						return v46
					}
				default:
					if v13 != int32(277) {
						if v13 != int32(318) {
							v50 = F_expression_tree_walker_impl(m, l0, int32(861), l1)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								return v50
							}
						} else {
							v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							switch v21 - int32(1) {
							case 0:
								v55 = int32(1)
								return v55
							case 1:
								v55 = int32(0)
								return v55
							default:
								v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v27 = F_contain_volatile_functions_walker(m, v26, l1)
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
									return int32(0)
								} else {
									if v27 != 0 {
										v29 = int32(1)
									} else {
										v29 = int32(2)
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v29
									return v27
								}
							}
						}
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						switch v33 - int32(1) {
						case 0:
							v55 = int32(1)
							return v55
						case 1:
							v55 = int32(0)
							return v55
						default:
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v39 = F_contain_volatile_functions_walker(m, v38, l1)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								if v39 != 0 {
									v41 = int32(1)
								} else {
									v41 = int32(2)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v41
								return v39
							}
						}
					}
				}
			}
		}
	}
}
