package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_coerce_fn_result_column(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	if l3 == int32(0) {
		v32 = int32(0)
		v35 = F_makeVarFromTargetEntry(m, int32(1), l0)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
			v41 = F_coerce_to_target_type(m, v32, v35, v37, l1, l2, int32(1), int32(2), int32(-1))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				if v41 == int32(0) {
					v73 = v32
					return v73
				} else {
					F_assign_expr_collations(m, int32(0), v41)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						if v41 == v35 {
							v51 = v41
						} else {
							v49 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v49)
							v51 = v41
						}
						v54 = int32(1)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						if v56 != 0 {
							v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)))
							v60 = v57 + int32(1)
						} else {
							v60 = v54
						}
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v64 = F_makeTargetEntry(m, v51, base.I32_extend16_s(v60), v62, int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							v67 = F_lappend(m, v66, v64)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v67
								v73 = v54
								return v73
							}
						}
					}
				}
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v10 != 0 {
			v32 = int32(0)
			v35 = F_makeVarFromTargetEntry(m, int32(1), l0)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
				v41 = F_coerce_to_target_type(m, v32, v35, v37, l1, l2, int32(1), int32(2), int32(-1))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					if v41 == int32(0) {
						v73 = v32
						return v73
					} else {
						F_assign_expr_collations(m, int32(0), v41)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							if v41 == v35 {
								v51 = v41
							} else {
								v49 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v49)
								v51 = v41
							}
							v54 = int32(1)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							if v56 != 0 {
								v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)))
								v60 = v57 + int32(1)
							} else {
								v60 = v54
							}
							v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v64 = F_makeTargetEntry(m, v51, base.I32_extend16_s(v60), v62, int32(0))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								v67 = F_lappend(m, v66, v64)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = v67
									v73 = v54
									return v73
								}
							}
						}
					}
				}
			}
		} else {
			v11 = int32(0)
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v14 = F_exprType(m, v13)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v21 = F_coerce_to_target_type(m, v11, v13, v14, l1, l2, int32(1), int32(2), int32(-1))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					if v21 == int32(0) {
						v73 = v11
						return v73
					} else {
						F_assign_expr_collations(m, int32(0), v21)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v21
							v30 = F_makeVarFromTargetEntry(m, int32(1), l0)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								v51 = v30
								v54 = int32(1)
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								if v56 != 0 {
									v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)))
									v60 = v57 + int32(1)
								} else {
									v60 = v54
								}
								v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v64 = F_makeTargetEntry(m, v51, base.I32_extend16_s(v60), v62, int32(0))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
									v67 = F_lappend(m, v66, v64)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l4))) = v67
										v73 = v54
										return v73
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
func F_get_fn_expr_variadic(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v2 = int32(0)
	if l0 == v2 {
		v13 = v2
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v5 == int32(0) {
			v13 = v2
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			if v8 != int32(15) {
				v13 = v2
			} else {
				v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+13)))
				v13 = v11
			}
		}
	}
	return v13 & int32(1)
}
