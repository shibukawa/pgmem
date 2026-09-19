package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_fn_expr_arg_stable(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	v3 = int32(0)
	if l0 == v3 {
		v49 = v3
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v7 == int32(0) {
			v49 = v3
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v12 = v10 - int32(11)
			v19 = int32(0)
			if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v12))|base.B2i32(int32(base.Ui32(int32(977))>>(uint(v12)%32))&int32(1) == v19)|base.B2i32(l1 < v19) != 0 {
				v49 = v3
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(2))%32))+uint32(_c_F_get_fn_expr_arg_stable[0])))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v7+v27)))
				if v29 == int32(0) {
					v49 = v3
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
					if v32 <= l1 {
						v49 = v3
					} else {
						v34 = int32(1)
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+l1<<(uint(int32(2))%32))))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
						switch v40 - int32(7) {
						case 0:
							v49 = v34
						case 1:
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							if v43 == int32(0) {
								v49 = v34
							} else {
								v49 = int32(0)
							}
						default:
							v49 = int32(0)
						}
					}
				}
			}
		}
	}
	return v49
}
func F_has_fn_opclass_options(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	v2 = int32(0)
	if l0 == v2 {
		v18 = v2
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v5 == int32(0) {
			v18 = v2
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			if v8 != int32(7) {
				v18 = v2
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
				if v11 != int32(17) {
					v18 = v2
				} else {
					v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+24)))
					v18 = v14 ^ int32(1)
				}
			}
		}
	}
	return v18 & int32(1)
}
func F_make_fn_arguments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v5 = int32(0)
	if l1 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v12 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = v5
	goto L4
L4:
	;
	v25 = v20 << (uint(int32(2)) % 32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l2+v25)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l3+v25)))
	if v27 == v29 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	v55 = v20 + int32(1)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v55 < v56 {
		v20 = v55
		goto L4
	} else {
		goto L14
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v32 = v31 + v25
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v34 == int32(16) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v38 = int32(-1)
	v42 = F_coerce_type(m, l0, v37, v27, v29, v38, int32(0), int32(2), v38)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v45 = int32(-1)
	v49 = F_coerce_type(m, l0, v33, v27, v29, v45, int32(0), int32(2), v45)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L11
	} else {
		goto L13
	}
L11:
	;
	return
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v42
	goto L6
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v49
	goto L6
L14:
	;
	goto L5
}
func Fn13823(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int32 {
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v9 = int32(4)
	v10 = base.I32_wrap_i64(l0)<<(uint(l3)%32) | v9
	v12 = base.I32_wrap_i64(l1) << (uint(l3) % 32)
	v14 = v12 | v9
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v14))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v10)) == int32(0) {
		v26 = base.B2i32(base.Ui32(v10) < base.Ui32(v14))
	} else {
		v26 = int32(base.Ui32(v10-v14) >> (uint(int32(31)) % 32))
	}
	if v26 != 0 {
		v27 = v12 + l2
		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v27))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v10)) == int32(0) {
			v39 = base.B2i32(base.Ui32(v10) < base.Ui32(v27))
		} else {
			v39 = int32(base.Ui32(v10-v27) >> (uint(int32(31)) % 32))
		}
		v41 = v39
	} else {
		v41 = int32(0)
	}
	return v41
}
func Fn13830(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	v3 = int32(0)
	v9 = F_SearchSysCacheList(m, l1, int32(1), l0, v3, v3)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	if int32(0) < v13 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = int32(0)
	v21 = v3
	goto L6
L4:
	;
	v40 = v3
	goto L5
L5:
	;
	F_ReleaseCatCacheList(m, v9)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L10
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(48)+v19<<(uint(int32(2))%32))))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28+v29)+4))
	v32 = F_lappend_oid(m, v21, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v40 = v32
	goto L5
L8:
	;
	v35 = v19 + int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	if v35 < v36 {
		v19 = v35
		v21 = v32
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	return v40
}
func Fn13843(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	F_errstart_cold(m, int32(21), int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		F_errcode(m, int32(1088))
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			F_errmsg(m, int32(_a_Fn13843_0), int32(0))
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_errdetail(m, int32(_a_Fn13843_1), int32(0))
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_Fn13843_2), l3, l2)
					v22 = m.ExcPending
					if v22 != 0 {
						return
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
func Fn13849(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v11 < int32(20) {
		v15 = v11 << (uint(int32(3)) % 32)
		*(*int32)(unsafe.Add(mBase, uint32(v15+l7))) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v15+l6))) = l0
		*(*int32)(unsafe.Add(mBase, uint32(l5))) = v11 + int32(1)
		v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_Fn13849[0])))
		if v24 == int32(0) {
			v29 = *(*int32)(unsafe.Add(mBase, _c_Fn13849[1]))
			if v29 <= int32(31) {
				*(*int32)(unsafe.Add(mBase, _c_Fn13849[1])) = v29 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v29<<(uint(int32(2))%32))+uint32(_c_Fn13849[2]))) = int32(1103)
			} else {
			}
			v46 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_Fn13849[0])) = uint8(v46)
		} else {
		}
		return
	} else {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return
			} else {
				F_errmsg_internal(m, l4, int32(0))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_Fn13849_0), l3, l2)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
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
func Fn13850(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v13, v14, v15, l1, int32(6))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v23 = F_LocalToUtf(m, v11, v15, v10, l5, l4, l3, l2, l1, base.B2i32(v12 != int32(0)))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			return v23
		}
	}
}
func Fn13856(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v5 = int32(_a_Fn13856_0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_Fn13856[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	*(*int32)(unsafe.Add(mBase, _c_Fn13856[0])) = v10
	F_varstr_sortsupport(m, v7, l1, v8)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_Fn13856[0])) = v6
		return int32(0)
	}
}
func Fn13867(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
	v13 = F_query_or_expression_tree_walker_impl(m, l0, l2, v7+int32(12), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v13
	}
}
func Fn13872(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	if v11 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v23
				F_errmsg(m, l3, v8)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn13872_0), l2, l1)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.G0 = v8 + int32(16)
		return int32(0)
	}
}
func Fn13887(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v22 = v11 + int32(8)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v25 = int32(4)
		v26 = v23 + v25
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = v26
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		v29 = int32(2)
		v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
		if base.Ui32(v30+v25) < base.Ui32(int32(base.Ui32(v38)>>(uint(v29)%32))) {
			v42 = v26 + (v30+int32(3))&int32(2147483644)
		} else {
			v42 = v26
		}
		*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v42
		v44 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v44)
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+16)))
		v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47+v48)+12)))
		v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v54 = F_gbt_var_consistent(m, v22, v15, v19, v46, v50&int32(1), l1, v53)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int32(0)
		} else {
			m.G0 = v11 + int32(16)
			return v54
		}
	}
}
func Fn13892(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+14)) = uint16(v14)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v16 + l2
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v16
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+16)))
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+v28)+12)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = F_gbt_num_consistent(m, v10+int32(4), v12, v10+int32(14), v30&int32(1), l1, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return int32(0)
	} else {
		m.G0 = v10 + int32(16)
		return v34
	}
}
func Fn13894(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_palloc0(m, int32(16))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(16)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = F_gbt_num_union(m, v7, v5, l1, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	}
}
func Fn13900(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, l1, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+68))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func Fn13911(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	if l1 < int32(0) {
		v16 = *(*int32)(unsafe.Add(mBase, _c_Fn13911[0]))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v16+(l1^int32(-1))<<(uint(int32(2))%32))))
		v30 = v22
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, _c_Fn13911[1]))
		v30 = v24 + l1<<(uint(int32(13))%32) + int32(-8192)
	}
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+14)))
	if v31 != 0 {
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+19)))
		v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+16)))
		if (v32<<(uint(int32(8))%32)-v35)&int32(_a_Fn13911_0) != int32(16) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return
			} else {
				F_errcode(m, int32(33557032))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return
				} else {
					v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					if l1 < int32(0) {
						v95 = *(*int32)(unsafe.Add(mBase, _c_Fn13911[2]))
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v95+(l1^int32(-1))<<(uint(int32(6))%32))+16))
						v110 = v101
					} else {
						v103 = *(*int32)(unsafe.Add(mBase, _c_Fn13911[3]))
						v109 = *(*int32)(unsafe.Add(mBase, uint32(v103+l1<<(uint(int32(6))%32)+int32(-64))+16))
						v110 = v109
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v110
					*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v91 + int32(4)
					F_errmsg(m, int32(_a_Fn13911_1), v11+int32(16))
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return
					} else {
						F_errhint(m, int32(_a_Fn13911_2), int32(0))
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return
						} else {
							F_errfinish(m, l4, l3, l2)
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		} else {
			m.G0 = v11 + int32(32)
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return
		} else {
			F_errcode(m, int32(33557032))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if l1 < int32(0) {
					v55 = *(*int32)(unsafe.Add(mBase, _c_Fn13911[2]))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v55+(l1^int32(-1))<<(uint(int32(6))%32))+16))
					v70 = v61
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, _c_Fn13911[3]))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v63+l1<<(uint(int32(6))%32)+int32(-64))+16))
					v70 = v69
				}
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v70
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v51 + int32(4)
				F_errmsg(m, int32(_a_Fn13911_3), v11)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return
				} else {
					F_errhint(m, int32(_a_Fn13911_2), int32(0))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						F_errfinish(m, l4, l5, l2)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
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
}
func Fn13917(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v20)
		v22 = F_get_role_oid_or_public(m, v14)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = F_convert_any_priv_string(m, v16, l1)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v28 = F_object_aclcheck_ext(m, l2, v13, v22, v24, v11+int32(15))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
					if v30 == int32(1) {
						v33 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v33)
						v37 = int32(0)
					} else {
						v37 = base.B2i32(v28 == int32(0))
					}
					m.G0 = v11 + int32(16)
					return v37
				}
			}
		}
	}
}
func Fn13922(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = F_table_open(m, l3, int32(1))
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		F_ScanKeyInit(m, v10, l2, int32(3), int32(184), l0)
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(1)
			v24 = F_systable_beginscan(m, v13, l1, v21, int32(0), v21, v10)
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = F_systable_getnext(m, v24)
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_systable_endscan(m, v24)
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						F_relation_close(m, v13, int32(1))
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							m.G0 = v10 + int32(48)
							return base.B2i32(v26 != int32(0))
						}
					}
				}
			}
		}
	}
}
func Fn13924(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if int32(0) < l1 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+18)))
		if base.Ui32(v16&int32(2047)) < base.Ui32(l1) {
			v20 = F_getmissingattr(m, l2, l1, l3)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v93 = v20
				m.G0 = v11 + int32(16)
				return v93
			}
		} else {
			v24 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v24)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+20)))
			if v27&int32(1) == v24 {
				v32 = int32(4)
				v36 = l2 + l1<<(uint(v32)%32) + v32
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
				if int32(0) <= v37 {
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+22)))
					v42 = v26 + v40 + v37
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+6)))
					if v43 != int32(1) {
						v93 = v42
						m.G0 = v11 + int32(16)
						return v93
					} else {
						v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v36)+4)))
						switch v46&int32(_a_Fn13924_0) - int32(1) {
						case 0:
							v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42))))
							v93 = v51
							m.G0 = v11 + int32(16)
							return v93
						case 1:
							v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42))))
							v93 = v52
							m.G0 = v11 + int32(16)
							return v93
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v46
								F_errmsg_internal(m, int32(_a_Fn13924_1), v11)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, l4, int32(70), int32(_a_Fn13924_2))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 3:
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
							v93 = v53
							m.G0 = v11 + int32(16)
							return v93
						}
					}
				} else {
					v66 = F_nocachegetattr(m, l0, l1, l2)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						v93 = v66
						m.G0 = v11 + int32(16)
						return v93
					}
				}
			} else {
				v68 = int32(1)
				v69 = l1 - v68
				v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+int32(base.Ui32(v69)>>(uint(int32(3))%32)))+23)))
				if int32(base.Ui32(v73)>>(uint(v69&int32(7))%32))&v68 == int32(0) {
					v81 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v81)
					v93 = int32(0)
					m.G0 = v11 + int32(16)
					return v93
				} else {
					v84 = F_nocachegetattr(m, l0, l1, l2)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						v93 = v84
						m.G0 = v11 + int32(16)
						return v93
					}
				}
			}
		}
	} else {
		v86 = F_heap_getsysattr(m, l0, l1, l3)
		mBase = m.M
		v87 = m.ExcPending
		if v87 != 0 {
			return int32(0)
		} else {
			v93 = v86
			m.G0 = v11 + int32(16)
			return v93
		}
	}
}
func Fn13937(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v10, v11, v12, l1, int32(7))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v20 = F_latin2mic(m, v8, v7, v12, l2, l1, base.B2i32(v9 != int32(0)))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			return v20
		}
	}
}
func Fn13939(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	if l0 == int32(0) {
		v11 = F_palloc(m, int32(32))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(4)
			*(*int64)(unsafe.Add(mBase, uint32(v11))) = l2
			v19 = v11 + int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v19
			v75 = v11
			v76 = v19
			v77 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v76+v77<<(uint(int32(2))%32)-int32(4)))) = l1
			return v75
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v23 <= v22 {
			v25 = int32(1)
			v27 = int32(16)
			v29 = v22 + v25
			if v29 <= v27 {
				v32 = v27
			} else {
				v32 = v29
			}
			if v32&(v32-int32(1)) != 0 {
				v39 = v25 << (uint(int32(32)-base.I32_clz(v32)) % 32)
			} else {
				v39 = v32
			}
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v42 = l0 + int32(16)
			if v40 == v42 {
				v44 = F_GetMemoryChunkContext(m, l0)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v48 = F_MemoryContextAlloc(m, v44, v39<<(uint(int32(2))%32))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v48
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v53 = v51 << (uint(int32(2)) % 32)
						if v53 == int32(0) {
						} else {
							base.MemoryCopy(m, v48, v42, v53)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v39
						v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v70 = v65
						v72 = v70 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
						v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v75 = l0
						v76 = v74
						v77 = v72
						*(*int32)(unsafe.Add(mBase, uint32(v76+v77<<(uint(int32(2))%32)-int32(4)))) = l1
						return v75
					}
				}
			} else {
				v59 = F_repalloc(m, v40, v39<<(uint(int32(2))%32))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v59
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v39
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v70 = v65
					v72 = v70 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v75 = l0
					v76 = v74
					v77 = v72
					*(*int32)(unsafe.Add(mBase, uint32(v76+v77<<(uint(int32(2))%32)-int32(4)))) = l1
					return v75
				}
			}
		} else {
			v70 = v22
			v72 = v70 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
			v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v75 = l0
			v76 = v74
			v77 = v72
			*(*int32)(unsafe.Add(mBase, uint32(v76+v77<<(uint(int32(2))%32)-int32(4)))) = l1
			return v75
		}
	}
}
func Fn13944(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l3
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
	if v11 == int32(1) {
		v14 = int32(_a_Fn13944_0)
		v15 = *(*int32)(unsafe.Add(mBase, _c_Fn13944[0]))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		*(*int32)(unsafe.Add(mBase, _c_Fn13944[0])) = v17
		v20 = F_palloc(m, int32(40))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)) = uint8(v24)
			*(*int64)(unsafe.Add(mBase, uint32(v20))) = int64(0)
			F_initHyperLogLog(m, v20+int32(16), int32(10))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(116)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v20
				*(*int32)(unsafe.Add(mBase, _c_Fn13944[0])) = v15
				return int32(0)
			}
		}
	} else {
		return int32(0)
	}
}
func Fn13953(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v10 != 0 {
		v13 = int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v13 = v12
	}
	v16 = F_numeric_stddev_internal(m, v13, l2, l1, v8+int32(15))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
		if v20 == int32(1) {
			v23 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
			v26 = int32(0)
		} else {
			v26 = v16
		}
		m.G0 = v8 + int32(16)
		return v26
	}
}
func Fn13955(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	v13 = int32(1)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v20 == v13 {
		if v16&int32(1) == int32(0) {
			v40 = v19
			v41 = v19
			v42 = v13
			if v15&int32(1) != 0 {
				if v14&int32(1) == int32(0) {
					v53 = v17
					v55 = F_DirectFunctionCall2Coll(m, l2, int32(0), v41, v53)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							v105 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
							v115 = int32(0)
							return v115
						} else {
							v79 = v53
							v80 = int32(1)
							v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								if v83 != 0 {
									if v42 != 0 {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
										return v115
									} else {
										v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											if v80&base.B2i32(v86 == int32(0)) != 0 {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v115 = int32(0)
												return v115
											} else {
												return base.B2i32(v86 != int32(0))
											}
										}
									}
								} else {
									if v80|v42 == int32(0) {
										v115 = int32(1)
									} else {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
									}
									return v115
								}
							}
						}
					}
				} else {
					v105 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
					v115 = int32(0)
					return v115
				}
			} else {
				if v14&int32(1) == int32(0) {
					v58 = int32(0)
					v61 = F_DirectFunctionCall2Coll(m, l2, v58, v18, v17)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						if v61 != 0 {
							v63 = v17
						} else {
							v63 = v18
						}
						v64 = F_DirectFunctionCall2Coll(m, l2, v58, v41, v63)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							if v64 == int32(0) {
								v79 = v63
								v80 = v58
								v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									if v83 != 0 {
										if v42 != 0 {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
											return v115
										} else {
											v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												if v80&base.B2i32(v86 == int32(0)) != 0 {
													v105 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
													v115 = int32(0)
													return v115
												} else {
													return base.B2i32(v86 != int32(0))
												}
											}
										}
									} else {
										if v80|v42 == int32(0) {
											v115 = int32(1)
										} else {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
										}
										return v115
									}
								}
							} else {
								if v61 != 0 {
									v69 = v18
								} else {
									v69 = v17
								}
								v70 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v69)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									if v42&base.B2i32(v70 == int32(0)) != 0 {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
										return v115
									} else {
										return base.B2i32(v70 != int32(0))
									}
								}
							}
						}
					}
				} else {
					v53 = v18
					v55 = F_DirectFunctionCall2Coll(m, l2, int32(0), v41, v53)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							v105 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
							v115 = int32(0)
							return v115
						} else {
							v79 = v53
							v80 = int32(1)
							v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								if v83 != 0 {
									if v42 != 0 {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
										return v115
									} else {
										v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											if v80&base.B2i32(v86 == int32(0)) != 0 {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v115 = int32(0)
												return v115
											} else {
												return base.B2i32(v86 != int32(0))
											}
										}
									}
								} else {
									if v80|v42 == int32(0) {
										v115 = int32(1)
									} else {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
									}
									return v115
								}
							}
						}
					}
				}
			}
		} else {
			v105 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
			v115 = int32(0)
			return v115
		}
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v16&int32(1) != 0 {
			v40 = v19
			v41 = v27
			v42 = v13
			if v15&int32(1) != 0 {
				if v14&int32(1) == int32(0) {
					v53 = v17
					v55 = F_DirectFunctionCall2Coll(m, l2, int32(0), v41, v53)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							v105 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
							v115 = int32(0)
							return v115
						} else {
							v79 = v53
							v80 = int32(1)
							v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								if v83 != 0 {
									if v42 != 0 {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
										return v115
									} else {
										v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											if v80&base.B2i32(v86 == int32(0)) != 0 {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v115 = int32(0)
												return v115
											} else {
												return base.B2i32(v86 != int32(0))
											}
										}
									}
								} else {
									if v80|v42 == int32(0) {
										v115 = int32(1)
									} else {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
									}
									return v115
								}
							}
						}
					}
				} else {
					v105 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
					v115 = int32(0)
					return v115
				}
			} else {
				if v14&int32(1) == int32(0) {
					v58 = int32(0)
					v61 = F_DirectFunctionCall2Coll(m, l2, v58, v18, v17)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						if v61 != 0 {
							v63 = v17
						} else {
							v63 = v18
						}
						v64 = F_DirectFunctionCall2Coll(m, l2, v58, v41, v63)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							if v64 == int32(0) {
								v79 = v63
								v80 = v58
								v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									if v83 != 0 {
										if v42 != 0 {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
											return v115
										} else {
											v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												if v80&base.B2i32(v86 == int32(0)) != 0 {
													v105 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
													v115 = int32(0)
													return v115
												} else {
													return base.B2i32(v86 != int32(0))
												}
											}
										}
									} else {
										if v80|v42 == int32(0) {
											v115 = int32(1)
										} else {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
										}
										return v115
									}
								}
							} else {
								if v61 != 0 {
									v69 = v18
								} else {
									v69 = v17
								}
								v70 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v69)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									if v42&base.B2i32(v70 == int32(0)) != 0 {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
										return v115
									} else {
										return base.B2i32(v70 != int32(0))
									}
								}
							}
						}
					}
				} else {
					v53 = v18
					v55 = F_DirectFunctionCall2Coll(m, l2, int32(0), v41, v53)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							v105 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
							v115 = int32(0)
							return v115
						} else {
							v79 = v53
							v80 = int32(1)
							v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								if v83 != 0 {
									if v42 != 0 {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
										return v115
									} else {
										v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											if v80&base.B2i32(v86 == int32(0)) != 0 {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v115 = int32(0)
												return v115
											} else {
												return base.B2i32(v86 != int32(0))
											}
										}
									}
								} else {
									if v80|v42 == int32(0) {
										v115 = int32(1)
									} else {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
									}
									return v115
								}
							}
						}
					}
				}
			}
		} else {
			v30 = int32(0)
			v32 = F_DirectFunctionCall2Coll(m, l2, v30, v27, v19)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				if v32 != 0 {
					v36 = v19
				} else {
					v36 = v27
				}
				if v32 != 0 {
					v37 = v27
				} else {
					v37 = v19
				}
				v40 = v37
				v41 = v36
				v42 = v30
				if v15&int32(1) != 0 {
					if v14&int32(1) == int32(0) {
						v53 = v17
						v55 = F_DirectFunctionCall2Coll(m, l2, int32(0), v41, v53)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							if v55 != 0 {
								v105 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
								v115 = int32(0)
								return v115
							} else {
								v79 = v53
								v80 = int32(1)
								v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									if v83 != 0 {
										if v42 != 0 {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
											return v115
										} else {
											v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												if v80&base.B2i32(v86 == int32(0)) != 0 {
													v105 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
													v115 = int32(0)
													return v115
												} else {
													return base.B2i32(v86 != int32(0))
												}
											}
										}
									} else {
										if v80|v42 == int32(0) {
											v115 = int32(1)
										} else {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
										}
										return v115
									}
								}
							}
						}
					} else {
						v105 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
						v115 = int32(0)
						return v115
					}
				} else {
					if v14&int32(1) == int32(0) {
						v58 = int32(0)
						v61 = F_DirectFunctionCall2Coll(m, l2, v58, v18, v17)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							if v61 != 0 {
								v63 = v17
							} else {
								v63 = v18
							}
							v64 = F_DirectFunctionCall2Coll(m, l2, v58, v41, v63)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								if v64 == int32(0) {
									v79 = v63
									v80 = v58
									v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										if v83 != 0 {
											if v42 != 0 {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v115 = int32(0)
												return v115
											} else {
												v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return int32(0)
												} else {
													if v80&base.B2i32(v86 == int32(0)) != 0 {
														v105 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
														v115 = int32(0)
														return v115
													} else {
														return base.B2i32(v86 != int32(0))
													}
												}
											}
										} else {
											if v80|v42 == int32(0) {
												v115 = int32(1)
											} else {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v115 = int32(0)
											}
											return v115
										}
									}
								} else {
									if v61 != 0 {
										v69 = v18
									} else {
										v69 = v17
									}
									v70 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v69)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										if v42&base.B2i32(v70 == int32(0)) != 0 {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
											return v115
										} else {
											return base.B2i32(v70 != int32(0))
										}
									}
								}
							}
						}
					} else {
						v53 = v18
						v55 = F_DirectFunctionCall2Coll(m, l2, int32(0), v41, v53)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							if v55 != 0 {
								v105 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
								v115 = int32(0)
								return v115
							} else {
								v79 = v53
								v80 = int32(1)
								v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									if v83 != 0 {
										if v42 != 0 {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
											return v115
										} else {
											v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												if v80&base.B2i32(v86 == int32(0)) != 0 {
													v105 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
													v115 = int32(0)
													return v115
												} else {
													return base.B2i32(v86 != int32(0))
												}
											}
										}
									} else {
										if v80|v42 == int32(0) {
											v115 = int32(1)
										} else {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
										}
										return v115
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
func Fn13964(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_get_statisticsobj_worker(m, v4, l1, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v12 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
			return int32(0)
		} else {
			v16 = F_cstring_to_text(m, v6)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v6)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					return v16
				}
			}
		}
	}
}
func Fn13973(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(176)
	m.G0 = v12
	v19 = v4
	v20 = int32(-1)
	v21 = v4
	v22 = v4
	goto L1
L1:
	;
	if v20 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn13973[0])) = v43
	*(*int32)(unsafe.Add(mBase, _c_Fn13973[1])) = v44
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v94 - int32(1)
	m.G0 = v12 + int32(176)
	return
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v27 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v26 + v27
	v31 = *(*int32)(unsafe.Add(mBase, _c_Fn13973[0]))
	v33 = *(*int32)(unsafe.Add(mBase, _c_Fn13973[1]))
	v35 = v12 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v12 + int32(12)
	goto L6
L4:
	;
	v42 = v19
	v43 = v21
	v44 = v22
	goto L5
L5:
	;
	goto L8
L6:
	;
	v42 = int32(0)
	v43 = v31
	v44 = v33
	goto L5
L7:
	;
	goto L2
L8:
	;
	if v42 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L7
L10:
	;
	v69 = int32(m.ExcTag)
	v70 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v69 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	F_standard_ExecutorFinish(m, l0)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L10
	} else {
		goto L18
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn13973[0])) = v12 + int32(16)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v51 == int32(0) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn13973[1])) = v44
	*(*int32)(unsafe.Add(mBase, _c_Fn13973[0])) = v43
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v60 - int32(1)
	F_pg_re_throw(m)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L10
	} else {
		goto L17
	}
L15:
	;
	m.T0[v51].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L7
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	goto L9
L19:
	;
	v74 = int32(v70)
	m.G0 = v12
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v12+int32(12) == v80 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	m.ExcPending = 1
	goto L28
L21:
	;
	if v84 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v84 = v82
	goto L24
L23:
	;
	v84 = int32(0)
	goto L24
L24:
	;
	goto L21
L25:
	;
	F___wasm_longjmp(m, v77, v76)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v19 = v76
	v20 = v84
	v21 = v43
	v22 = v44
	goto L1
L28:
	;
	return
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func Fn13979(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v4
	v17 = F_query_or_expression_tree_walker_impl(m, l1, int32(896), v7+int32(4), v4)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		m.G0 = v7 + int32(16)
		return v21
	}
}
func Fn13980(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l3
	v15 = F_query_or_expression_tree_walker_impl(m, l0, l2, v8+int32(8), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		m.G0 = v8 + int32(16)
		return v19
	}
}
func Fn13991(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 float64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = F_get_negator(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_Fn13991_0), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn13991_1), int32(773), int32(_a_Fn13991_2))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v33 = F_patternsel_common(m, v10, v12, int32(0), v9, v8, v7, l1, int32(1))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = F_Float8GetDatum(m, v33)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					return v35
				}
			}
		}
	}
}
func Fn13997(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = v9 + int32(1)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			v19 = int32(1)
			v20 = v18 & v19
			if v18 == v19 {
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
				if v26 == int32(18) {
					v29 = int32(16)
				} else {
					v29 = int32(0)
				}
				if base.Ui32((v26-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v36 = int32(4)
				} else {
					v36 = v29
				}
				v47 = v36
			} else {
				v37 = int32(1)
				if v20 != 0 {
					v47 = int32(base.Ui32(v18)>>(uint(v37)%32)) - v37
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v49 = F_RE_compile_and_cache(m, v16, l1, v48)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				v55 = F_palloc(m, v47<<(uint(int32(2))%32)+int32(4))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					if v20 != 0 {
						v59 = v14
					} else {
						v59 = v9 + int32(4)
					}
					v60 = F_pg_mb2wchar_with_len(m, v59, v55, v47)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v62 = int32(0)
						v65 = F_RE_wchar_execute(m, v55, v60, v62, v62, v62)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v55)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								return v65
							}
						}
					}
				}
			}
		}
	}
}
func Fn14004(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) <= v8 {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
		v14 = F_psprintf(m, int32(_a_Fn14004_0), v6)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v20 = v14
			m.G0 = v6 + int32(16)
			return v20
		}
	} else {
		v18 = F_pstrdup(m, l1)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = v18
			m.G0 = v6 + int32(16)
			return v20
		}
	}
}
func Fn14008(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v25 int32
	_ = v25
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v15
		v18 = F_text_to_cstring(m, v11)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v24 = F_parse_tsquery(m, v18, int32(1158), v8+int32(8), l1, int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(16)
				return v24
			}
		}
	}
}
func Fn14013(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v14)+79)))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v16 == l1)
			}
		}
	}
}
func Fn14015(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v13, v14, v15, int32(6), l1)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v23 = F_UtfToLocal(m, v11, v15, v10, l5, l4, l3, l2, l1, base.B2i32(v12 != int32(0)))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			return v23
		}
	}
}
func Fn14019(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = F_WinGetFuncArgInFrame(m, v9, int32(0), l1, int32(1), v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v24 = int32(0)
		} else {
			v24 = v14
		}
		m.G0 = v7 + int32(16)
		return v24
	}
}
func Fn14024(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 float32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v17 = v10 + int32(1)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = F_pg_detoast_datum_packed(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v23 = v21 & int32(1)
			if v23 != 0 {
				v24 = v17
			} else {
				v24 = v10 + int32(4)
			}
			if v21 == int32(1) {
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				if v30 == int32(18) {
					v33 = int32(16)
				} else {
					v33 = int32(0)
				}
				if base.Ui32((v30-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v40 = int32(4)
				} else {
					v40 = v33
				}
				v51 = v40
			} else {
				v41 = int32(1)
				if v23 != 0 {
					v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v52 = int32(1)
			v53 = v19 + v52
			v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
			v58 = v56 & v52
			if v58 != 0 {
				v59 = v53
			} else {
				v59 = v19 + int32(4)
			}
			if v56 == int32(1) {
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
				if v65 == int32(18) {
					v68 = int32(16)
				} else {
					v68 = int32(0)
				}
				if base.Ui32((v65-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v75 = int32(4)
				} else {
					v75 = v68
				}
				v86 = v75
			} else {
				v76 = int32(1)
				if v58 != 0 {
					v86 = int32(base.Ui32(v56)>>(uint(v76)%32)) - v76
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					v86 = int32(base.Ui32(v80)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v87 = F_calc_word_similarity(m, v24, v51, v59, v86, l1)
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return int32(0)
			} else {
				v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v89 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v93 != v19 {
							F_pfree(m, v19)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int32(0)
							} else {
								return base.I32_reinterpret_f32(base.F32_sub(float32(1), v87))
							}
						} else {
							return base.I32_reinterpret_f32(base.F32_sub(float32(1), v87))
						}
					}
				} else {
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v93 != v19 {
						F_pfree(m, v19)
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							return base.I32_reinterpret_f32(base.F32_sub(float32(1), v87))
						}
					} else {
						return base.I32_reinterpret_f32(base.F32_sub(float32(1), v87))
					}
				}
			}
		}
	}
}
