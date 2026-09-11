package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_contain_invalid_rfcolumn_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v19 int32
	_ = v19
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
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	if l0 == int32(0) {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v9 == int32(6) {
			v12 = int32(1)
			v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
			if v14 == v12 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v19 = F_get_attname(m, v17, v13, int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v24 = F_get_attnum(m, v23, v19)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v27 = v24
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v31 = F_bms_is_member(m, v27+int32(7), v30)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							if v31 == int32(0) {
								v41 = v12
								return v41
							} else {
								v38 = F_expression_tree_walker_impl(m, l0, int32(566), l1)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									v41 = v38
									return v41
								}
							}
						}
					}
				}
			} else {
				v27 = v13
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v31 = F_bms_is_member(m, v27+int32(7), v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					if v31 == int32(0) {
						v41 = v12
						return v41
					} else {
						v38 = F_expression_tree_walker_impl(m, l0, int32(566), l1)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v41 = v38
							return v41
						}
					}
				}
			}
		} else {
			v38 = F_expression_tree_walker_impl(m, l0, int32(566), l1)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v41 = v38
				return v41
			}
		}
	}
}
func F_contain_leaked_vars_checker(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_get_func_leakproof(m, l0)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3 ^ int32(1)
	}
}
func F_contain_placeholder_references_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v8 != int32(67) {
			if v8 != int32(319) {
				v37 = F_expression_tree_walker_impl(m, l0, int32(880), l1)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					return v37
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v13 != v14 {
					v37 = F_expression_tree_walker_impl(m, l0, int32(880), l1)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						return v37
					}
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v18 = F_bms_is_member(m, v16, v17)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						return v18
					}
				}
			}
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v23 + int32(1)
			v29 = F_query_tree_walker_impl(m, l0, int32(880), l1, int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v31 - int32(1)
				return v29
			}
		}
	}
}
func F_contain_subplans_walker(m *base.Module, l0 int32, l1 int32) int32 {
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
		if base.Ui32(v7-int32(22)) < base.Ui32(int32(3)) {
			return int32(1)
		} else {
			v15 = F_expression_tree_walker_impl(m, l0, int32(856), l1)
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
func F_contain_vars_of_level(m *base.Module, l0 int32, l1 int32) int32 {
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
	v13 = F_query_or_expression_tree_walker_impl(m, l0, int32(899), v6+int32(12), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v13
	}
}
func F_contain_vars_returning_old_or_new_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v3 = int32(0)
	if l0 == v3 {
		v29 = v3
		return v29
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 != int32(61) {
			if v7 != int32(6) {
				v24 = F_expression_tree_walker_impl(m, l0, int32(900), l1)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v29 = v24
					return v29
				}
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v12 == int32(0) {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					if v16 != 0 {
						v29 = int32(1)
						return v29
					} else {
						return int32(0)
					}
				} else {
					return int32(0)
				}
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			return base.B2i32(v19 == int32(0))
		}
	}
}
func F_contain_volatile_functions_not_nextval_checker(m *base.Module, l0 int32, l1 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	if l0 != int32(1574) {
		v6 = F_func_volatile(m, l0)
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v12 = base.B2i32(v6 == int32(118))
			return v12
		}
	} else {
		v12 = int32(0)
		return v12
	}
}
func F_contain_windowfuncs(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v3 = int32(0)
	v5 = F_query_or_expression_tree_walker_impl(m, l0, int32(1044), v3, v3)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
