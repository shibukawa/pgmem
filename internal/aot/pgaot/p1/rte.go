package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetRTEByRangeTablePosn(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	v4 = int32(0)
	if l2 <= v4 {
		v53 = l0
	} else {
		v10 = l2 & int32(7)
		if v10 == int32(0) {
			v25 = l0
			v28 = l2
		} else {
			v13 = l0
			v16 = l2
			v17 = v4
			for {
				v19 = int32(1)
				v20 = v16 - v19
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v23 = v17 + v19
				if v23 != v10 {
					v13 = v21
					v16 = v20
					v17 = v23
					continue
				} else {
					break
				}
				break
			}
			v25 = v21
			v28 = v20
		}
		if base.Ui32(l2) < base.Ui32(int32(8)) {
			v53 = v25
		} else {
			v33 = v25
			v36 = v28
			for {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
				if base.Ui32(v36-int32(9)) < base.Ui32(int32(-2)) {
					v33 = v48
					v36 = v36 - int32(8)
					continue
				} else {
					break
				}
				break
			}
			v53 = v48
		}
	}
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v60+l1<<(uint(int32(2))%32)-int32(4))))
	return v66
}
func F_addRTEPermissionInfo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v24 int32
	_ = v24
	v5 = F_palloc0(m, int32(40))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(102)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v11
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+8)) = uint8(v13)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = F_lappend(m, v15, v5)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v16
			if v16 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(0)
				return v5
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v24
				return v5
			}
		}
	}
}
func F_replace_rte_variables_mutator(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v8 - int32(58) {
		case 0:
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v32 != v33 {
				v74 = F_expression_tree_mutator_impl(m, l0, int32(1053), l1)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					v76 = v74
					return v76
				}
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				if v35 != 0 {
					v74 = F_expression_tree_mutator_impl(m, l0, int32(1053), l1)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						v76 = v74
						return v76
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(424687), int32(0))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(472969), int32(1520), int32(198526))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
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
		case 1, 2, 3, 4, 5, 6, 7, 8:
			v74 = F_expression_tree_mutator_impl(m, l0, int32(1053), l1)
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				v76 = v74
				return v76
			}
		case 9:
			v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v52 + int32(1)
			v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
			v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v57)
			v61 = F_query_tree_mutator_impl(m, l0, int32(1053), l1, int32(0))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+39)))
				v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
				v65 = v63 | v64
				*(*uint8)(unsafe.Add(mBase, uint32(v61)+39)) = uint8(v65)
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v56)
				v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v68 - int32(1)
				return v61
			}
		default:
			if v8 != int32(6) {
				v74 = F_expression_tree_mutator_impl(m, l0, int32(1053), l1)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					v76 = v74
					return v76
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				if v13 != v14 {
					v74 = F_expression_tree_mutator_impl(m, l0, int32(1053), l1)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						v76 = v74
						return v76
					}
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					if v16 != v17 {
						v74 = F_expression_tree_mutator_impl(m, l0, int32(1053), l1)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							v76 = v74
							return v76
						}
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v20 = m.T0[v19].(func(*base.Module, int32, int32) int32)(m, l0, l1)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
							if v24 != 0 {
								v76 = v20
								return v76
							} else {
								v28 = F_query_or_expression_tree_walker_impl(m, v20, int32(1046), int32(0), int32(3))
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return int32(0)
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v28)
									return v20
								}
							}
						}
					}
				}
			}
		}
	}
}
