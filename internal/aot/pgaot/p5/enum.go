package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_call_enum_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v14 == int32(0) {
		v134 = v13
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L37
	}
L2:
	;
	m.G0 = v11 + int32(80)
	return v134
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_call_enum_check_hook[0])) = int32(50856066)
	v21 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_call_enum_check_hook[1])) = v21
	*(*int32)(unsafe.Add(mBase, _c_F_call_enum_check_hook[2])) = v21
	*(*int32)(unsafe.Add(mBase, _c_F_call_enum_check_hook[3])) = v21
	v29 = m.T0[v14].(func(*base.Module, int32, int32, int32) int32)(m, l1, l2, l3)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if v29 != 0 {
		v134 = v13
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v34 = F_errstart(m, l4, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if v34 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_call_enum_check_hook[0]))
	F_errcode(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_FlushErrorState(m)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L36
	}
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_call_enum_check_hook[1]))
	if v41 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_call_enum_check_hook[2]))
	if v97 != 0 {
		goto L27
	} else {
		goto L28
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v41
	F_errmsg_internal(m, int32(_a_F_call_enum_check_hook_0), v11-int32(-64))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v49 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L12
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v52 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v48 != v56 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v64 = v49
	goto L22
L20:
	;
	v75 = v52
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v55
	F_errmsg(m, int32(_a_F_call_enum_check_hook_1), v11+int32(48))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L26
	}
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	if v66 == int32(0) {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v75 = v66
	goto L21
L24:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	if v48 != v69 {
		v64 = v64 + int32(12)
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	goto L12
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v97
	F_errdetail_internal(m, int32(_a_F_call_enum_check_hook_0), v11+int32(32))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_call_enum_check_hook[3]))
	if v105 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v105
	F_errhint(m, int32(_a_F_call_enum_check_hook_0), v11+int32(16))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	F_errfinish(m, int32(_a_F_call_enum_check_hook_2), int32(_a_F_call_enum_check_hook_3), int32(_a_F_call_enum_check_hook_4))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	goto L10
L36:
	;
	v134 = int32(0)
	goto L2
L37:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v48
	F_errmsg_internal(m, int32(_a_F_call_enum_check_hook_5), v11)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_call_enum_check_hook_2), int32(2938), int32(_a_F_call_enum_check_hook_6))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_enum_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	var v83 int64
	_ = v83
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = base.I32_wrap_i64(v13)
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = base.I32_wrap_i64(v15)
	v17 = F_strlen(m, v16)
	mBase = m.M
	if base.Ui32(int32(64)) <= base.Ui32(v17) {
		v20 = int64(0)
		v21 = F_errsave_start(m, v12)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			if v21 == int32(0) {
				v83 = v20
				m.G0 = v10 + int32(32)
				return v83
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					v30 = F_format_type_be(m, v14)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v16
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v30
						F_errmsg(m, int32(_a_F_enum_in_0), v10)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int64(0)
						} else {
							F_errsave_finish(m, v12, int32(_a_F_enum_in_1), int32(123), int32(_a_F_enum_in_2))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int64(0)
							} else {
								v83 = v20
								m.G0 = v10 + int32(32)
								return v83
							}
						}
					}
				}
			}
		}
	} else {
		v43 = int64(4294967295)
		v47 = F_SearchSysCache2(m, int32(24), v13&v43, v15&v43)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int64(0)
		} else {
			if v47 == int32(0) {
				v51 = int64(0)
				v52 = F_errsave_start(m, v12)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int64(0)
				} else {
					if v52 == int32(0) {
						v83 = v51
						m.G0 = v10 + int32(32)
						return v83
					} else {
						F_errcode(m, int32(33685634))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int64(0)
						} else {
							v59 = F_format_type_be(m, v14)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v16
								*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v59
								F_errmsg(m, int32(_a_F_enum_in_0), v10+int32(16))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int64(0)
								} else {
									F_errsave_finish(m, v12, int32(_a_F_enum_in_1), int32(133), int32(_a_F_enum_in_2))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int64(0)
									} else {
										v83 = v51
										m.G0 = v10 + int32(32)
										return v83
									}
								}
							}
						}
					}
				}
			} else {
				F_check_safe_enum_use(m, v47)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int64(0)
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
					v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+22)))
					v78 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v75+v76))))
					F_ReleaseCatCache(m, v47)
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int64(0)
					} else {
						v83 = v78
						m.G0 = v10 + int32(32)
						return v83
					}
				}
			}
		}
	}
}
func F_enum_larger(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = F_enum_cmp_internal(m, base.I32_wrap_i64(v4), base.I32_wrap_i64(v5), l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		if int32(0) < v8 {
			v14 = v4
		} else {
			v14 = v5
		}
		return v14 & int64(4294967295)
	}
}
func F_enum_range_all(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = F_get_fn_expr_argtype(m, v2, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		if v4 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_enum_range_all_0), int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_enum_range_all_1), int32(540), int32(_a_F_enum_range_all_2))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v26 = int32(0)
			v28 = F_enum_range_internal(m, v4, v26, v26)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v28)
			}
		}
	}
}
func F_enum_range_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v19 = v14 + int32(-56)
	F_ScanKeyInit(m, v19, int32(2), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v33 = F_table_open(m, int32(3501), int32(1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v37 = F_index_open(m, int32(3534), int32(1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v41 = F_systable_beginscan_ordered(m, v33, v37, int32(0), int32(1), v19)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v44 = F_palloc(m, int32(512))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v49 = v44
	v50 = int32(64)
	v51 = int32(0)
	v54 = base.B2i32(l1 == int32(0))
	goto L7
L7:
	;
	v60 = F_systable_getnext_ordered(m, v41, int32(1))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	F_systable_endscan_ordered(m, v41)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L22
	}
L9:
	;
	if v60 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+22)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63)))
	v67 = v54 | base.B2i32(l1 == v65)
	if v67&int32(1) != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v95 = v49
	v97 = v51
	goto L12
L12:
	;
	goto L8
L13:
	;
	F_check_safe_enum_use(m, v60)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v88 = v49
	v89 = v50
	v90 = v51
	goto L15
L15:
	;
	if base.B2i32(l2 == int32(0))|base.B2i32(l2 != v65) != 0 {
		v49 = v88
		v50 = v89
		v51 = v90
		v54 = v67
		goto L7
	} else {
		goto L21
	}
L16:
	;
	if v50 <= v51 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v75 = F_repalloc(m, v49, v50<<(uint(int32(4))%32))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	v79 = v49
	v80 = v50
	goto L19
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v79+v51<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v65)
	v88 = v79
	v89 = v80
	v90 = v51 + int32(1)
	goto L15
L20:
	;
	v79 = v75
	v80 = v50 << (uint(int32(1)) % 32)
	goto L19
L21:
	;
	v95 = v88
	v97 = v90
	goto L12
L22:
	;
	F_relation_close(m, v37, int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_relation_close(m, v33, int32(1))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v111 = F_construct_array(m, v95, v97, l0, int32(4), int32(1), int32(105))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_pfree(m, v95)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	m.G0 = v16 - int32(-64)
	return v111
}
