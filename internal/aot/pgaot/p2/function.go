package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FunctionCall2Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v5 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+44)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l2
	v16 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+30)) = uint16(v16)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l0
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = m.T0[v26].(func(*base.Module, int32) int32)(m, v8+int32(12))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
		if v31 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v38
				F_errmsg_internal(m, int32(_a_F_FunctionCall2Coll_0), v8)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_FunctionCall2Coll_1), int32(1165), int32(_a_F_FunctionCall2Coll_2))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v8 + int32(48)
			return v27
		}
	}
}
func F_FunctionCall8Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v11 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+92)) = uint8(v11)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = l9
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+84)) = uint8(v11)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = l8
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)) = uint8(v11)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = l7
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+68)) = uint8(v11)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = l6
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+60)) = uint8(v11)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+52)) = uint8(v11)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+44)) = uint8(v11)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+36)) = uint8(v11)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l2
	v40 = int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+30)) = uint16(v40)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)) = uint8(v11)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l0
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, v14+int32(12))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		return int32(0)
	} else {
		v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)))
		if v55 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v62
				F_errmsg_internal(m, int32(_a_F_FunctionCall8Coll_0), v14)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_FunctionCall8Coll_1), int32(1348), int32(_a_F_FunctionCall8Coll_2))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v14 + int32(96)
			return v51
		}
	}
}
func F_function_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 float64
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 float64
	_ = v41
	var v52 int32
	_ = v52
	var v53 float64
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 float64
	_ = v64
	v5 = l4
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	v16 = float64(0.3333333)
	v17 = F_get_func_support(m, l1)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return float64(0)
	} else {
		if v17 == int32(0) {
			v64 = v16
			m.G0 = v14 - int32(-64)
			return v64
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v14)+56)) = int64(-4616189618054758400)
			*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l7
			*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l6
			*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = l5
			*(*uint8)(unsafe.Add(mBase, uint32(v14)+36)) = uint8(v5)
			*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(458)
			v37 = v12 + int32(-48)
			v38 = F_OidFunctionCall1Coll(m, v17, int32(0), v37)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return float64(0)
			} else {
				if v38 != v37 {
					v64 = v16
					m.G0 = v14 - int32(-64)
					return v64
				} else {
					v41 = *(*float64)(unsafe.Add(mBase, uint32(v14)+56))
					if base.F64_lt(v41, float64(0))|base.F64_gt(v41, float64(1)) == int32(0) {
						v64 = v41
						m.G0 = v14 - int32(-64)
						return v64
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return float64(0)
						} else {
							v53 = *(*float64)(unsafe.Add(mBase, uint32(v14)+56))
							*(*float64)(unsafe.Add(mBase, uint32(v14))) = v53
							F_errmsg_internal(m, int32(_a_F_function_selectivity_0), v14)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return float64(0)
							} else {
								F_errfinish(m, int32(_a_F_function_selectivity_1), int32(2106), int32(_a_F_function_selectivity_2))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return float64(0)
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
	}
}
func F_generate_function_name(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	v5 = l4
	v8 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	v18 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v171 = F_quote_identifier(m, v26)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L3
	} else {
		goto L48
	}
L2:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	v150 = F_get_namespace_name_or_temp(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L3
	} else {
		goto L43
	}
L3:
	;
	return int32(0)
L4:
	;
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
	v24 = v22 + v23
	v26 = v24 + int32(4)
	if l6 == int32(0) {
		v84 = v8
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L3
	} else {
		goto L40
	}
L8:
	;
	if l5 != 0 {
		goto L28
	} else {
		goto L29
	}
L9:
	;
	v29 = int32(_a_F_generate_function_name_0)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generate_function_name[0])))
	if base.B2i32(v32 == int32(0))|base.B2i32(v32 != v35) != 0 {
		v53 = v32
		v54 = v35
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v53-v54 != 0 {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	goto L10
L12:
	;
	v38 = v26
	v39 = v29
	goto L13
L13:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	if v43 == int32(0) {
		v53 = v43
		v54 = v42
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v53 = v43
	v54 = v42
	goto L11
L15:
	;
	v46 = int32(1)
	if v43 == v42 {
		v38 = v38 + v46
		v39 = v39 + v46
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v56 = int32(_a_F_generate_function_name_1)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generate_function_name[1])))
	if base.B2i32(v59 == int32(0))|base.B2i32(v59 != v62) != 0 {
		v80 = v59
		v81 = v62
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	v84 = int32(1)
	goto L8
L20:
	;
	if v80-v81 != 0 {
		v84 = v8
		goto L8
	} else {
		goto L27
	}
L21:
	;
	goto L20
L22:
	;
	v65 = v26
	v66 = v56
	goto L23
L23:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
	if v70 == int32(0) {
		v80 = v70
		v81 = v69
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v80 = v70
	v81 = v69
	goto L21
L25:
	;
	v73 = int32(1)
	if v70 == v69 {
		v65 = v65 + v73
		v66 = v66 + v73
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	goto L19
L28:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v5)
	v89 = v5 ^ int32(1)
	goto L30
L29:
	;
	v89 = int32(1)
	goto L30
L30:
	;
	if v84 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = int32(0)
	goto L2
L32:
	;
	goto L33
L33:
	;
	v92 = F_makeString(m, v26)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v92
	v99 = F_list_make1_impl(m, int32(1), v15+int32(32))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v101 = int32(0)
	v117 = F_func_get_detail(m, v99, v101, l2, l1, l3, v89, int32(1), v101, v15+int32(60), v15+int32(56), v15+int32(55), v15+int32(48), v15+int32(44), v15+int32(40), v101)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v117))|base.B2i32(int32(1)<<(uint(v117)%32)&int32(52) == int32(0)) != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	if v128 != l0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	F_initStringInfo(m, v15-int32(-64))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	goto L1
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	F_errmsg_internal(m, int32(_a_F_generate_function_name_2), v15)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_generate_function_name_3), int32(_a_F_generate_function_name_4), int32(_a_F_generate_function_name_5))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	v153 = v15 - int32(-64)
	F_initStringInfo(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	if v150 == int32(0) {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v158 = F_quote_identifier(m, v150)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v158
	F_appendStringInfo(m, v153, int32(_a_F_generate_function_name_6), v15+int32(16))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	goto L1
L48:
	;
	F_appendStringInfoString(m, v15-int32(-64), v171)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	F_ReleaseCatCache(m, v18)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	m.G0 = v15 + int32(80)
	return v175
}
func F_print_function_rettype(m *base.Module, l0 int32, l1 int32) {
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
	F_initStringInfo(m, v8)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = v10 + v11
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+100)))
		if v15 != int32(1) {
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
			v45 = F_format_type_be(m, v44)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				F_appendStringInfoString(m, v8, v45)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					F_appendBinaryStringInfo(m, l0, v50, v51)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		} else {
			F_appendStringInfoString(m, v8, int32(_a_F_print_function_rettype_0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v23 = F_print_function_arguments(m, v8, l1, int32(1), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					if int32(0) < v23 {
						F_appendStringInfoChar(m, v8, int32(41))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
							F_appendBinaryStringInfo(m, l0, v50, v51)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						v31 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v31)
						*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v31
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v31
						if v23 != 0 {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
							F_appendBinaryStringInfo(m, l0, v50, v51)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						} else {
							v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+100)))
							if v37 != int32(1) {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
								v45 = F_format_type_be(m, v44)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
								} else {
									F_appendStringInfoString(m, v8, v45)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
										v51 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
										F_appendBinaryStringInfo(m, l0, v50, v51)
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									}
								}
							} else {
								F_appendStringInfoString(m, v8, int32(_a_F_print_function_rettype_1))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
									v45 = F_format_type_be(m, v44)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										F_appendStringInfoString(m, v8, v45)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											v50 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
											v51 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
											F_appendBinaryStringInfo(m, l0, v50, v51)
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return
											} else {
												m.G0 = v8 + int32(16)
												return
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
	}
}
