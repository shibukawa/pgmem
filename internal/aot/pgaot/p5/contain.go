package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_contain_aggs_of_level_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v8 - int32(9) {
		case 0:
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v11 != v12 {
				v24 = F_expression_tree_walker_impl(m, l0, int32(1042), l1)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					return v24
				}
			} else {
				return int32(1)
			}
		case 1:
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v16 != v17 {
				v24 = F_expression_tree_walker_impl(m, l0, int32(1042), l1)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					return v24
				}
			} else {
				return int32(1)
			}
		default:
			if v8 == int32(67) {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29 + int32(1)
				v35 = F_query_tree_walker_impl(m, l0, int32(1042), l1, int32(0))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v37 - int32(1)
					return v35
				}
			} else {
				v24 = F_expression_tree_walker_impl(m, l0, int32(1042), l1)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					return v24
				}
			}
		}
	}
}
func F_contain_non_const_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v3 = int32(0)
	if l0 == v3 {
		v16 = v3
		return v16
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v6 - int32(1) {
		case 0:
			v10 = F_expression_tree_walker_impl(m, l0, int32(871), l1)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v10
			}
		default:
			v16 = int32(1)
			return v16
		case 6:
			v16 = v3
			return v16
		}
	}
}
func F_contain_strippable_phv_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 != int32(319) {
			v14 = F_expression_tree_walker_impl(m, l0, int32(824), l1)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v14
			}
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v10 != 0 {
				v14 = F_expression_tree_walker_impl(m, l0, int32(824), l1)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					return v14
				}
			} else {
				return int32(1)
			}
		}
	}
}
func F_contain_subplans(m *base.Module, l0 int32) int32 {
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
		if base.Ui32(v6-int32(22)) < base.Ui32(int32(3)) {
			return int32(1)
		} else {
			v15 = F_expression_tree_walker_impl(m, l0, int32(856), int32(0))
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
func F_contain_volatile_functions_checker(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_func_volatile(m, l0)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v3 == int32(118))
	}
}
func F_contain_volatile_functions_not_nextval_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = F_check_functions_in_node(m, l0, int32(861), l1)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			if v8 != 0 {
				return int32(1)
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v14 == int32(67) {
					v19 = F_query_tree_walker_impl(m, l0, int32(862), l1, int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						return v19
					}
				} else {
					v23 = F_expression_tree_walker_impl(m, l0, int32(862), l1)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						return v23
					}
				}
			}
		}
	}
}
