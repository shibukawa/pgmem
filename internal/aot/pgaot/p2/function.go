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
				F_errmsg_internal(m, int32(530590), v8)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493374), int32(1165), int32(303171))
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
				F_errmsg_internal(m, int32(530590), v14)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493374), int32(1348), int32(303039))
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 float64
	_ = v43
	var v54 int32
	_ = v54
	var v55 float64
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 float64
	_ = v65
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
			v65 = v16
			m.G0 = v14 - int32(-64)
			return v65
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
			v38 = F_OidFunctionCall1Coll(m, v17, int32(0), v12+int32(-48))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return float64(0)
			} else {
				if v38 != v12+int32(-48) {
					v65 = v16
					m.G0 = v14 - int32(-64)
					return v65
				} else {
					v43 = *(*float64)(unsafe.Add(mBase, uint32(v14)+56))
					if base.F64_lt(v43, float64(0))|base.F64_gt(v43, float64(1)) == int32(0) {
						v65 = v43
						m.G0 = v14 - int32(-64)
						return v65
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return float64(0)
						} else {
							v55 = *(*float64)(unsafe.Add(mBase, uint32(v14)+56))
							*(*float64)(unsafe.Add(mBase, uint32(v14))) = v55
							F_errmsg_internal(m, int32(339078), v14)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return float64(0)
							} else {
								F_errfinish(m, int32(491772), int32(2106), int32(9871))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
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
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
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
	v169 = F_quote_identifier(m, v26)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L3
	} else {
		goto L51
	}
L2:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	v147 = F_get_namespace_name_or_temp(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L3
	} else {
		goto L46
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
		v82 = v8
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
	v134 = m.ExcPending
	if v134 != 0 {
		goto L3
	} else {
		goto L43
	}
L8:
	;
	if l5 != 0 {
		goto L30
	} else {
		goto L31
	}
L9:
	;
	v29 = int32(418604)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1142])))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v33 == int32(0) {
		v52 = v32
		v53 = v33
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v53-v52 != 0 {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	goto L10
L12:
	;
	if v32 != v33 {
		v52 = v32
		v53 = v33
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v37 = v26
	v38 = v29
	goto L14
L14:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	if v42 == int32(0) {
		v52 = v41
		v53 = v42
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v52 = v41
	v53 = v42
	goto L11
L16:
	;
	v45 = int32(1)
	if v41 == v42 {
		v37 = v37 + v45
		v38 = v38 + v45
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v55 = int32(232486)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1143])))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v59 == int32(0) {
		v78 = v58
		v79 = v59
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	v82 = int32(1)
	goto L8
L21:
	;
	if v79-v78 != 0 {
		v82 = v8
		goto L8
	} else {
		goto L29
	}
L22:
	;
	goto L21
L23:
	;
	if v58 != v59 {
		v78 = v58
		v79 = v59
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v63 = v26
	v64 = v55
	goto L25
L25:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	if v68 == int32(0) {
		v78 = v67
		v79 = v68
		goto L22
	} else {
		goto L27
	}
L26:
	;
	v78 = v67
	v79 = v68
	goto L22
L27:
	;
	v71 = int32(1)
	if v67 == v68 {
		v63 = v63 + v71
		v64 = v64 + v71
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	goto L20
L30:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v5)
	v87 = v5 ^ int32(1)
	goto L32
L31:
	;
	v87 = int32(1)
	goto L32
L32:
	;
	if v82 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = int32(0)
	goto L2
L34:
	;
	goto L35
L35:
	;
	v90 = F_makeString(m, v26)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v90
	v97 = F_list_make1_impl(m, int32(1), v15+int32(32))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v99 = int32(0)
	v115 = F_func_get_detail(m, v97, v99, l2, l1, l3, v87, int32(1), v99, v15+int32(60), v15+int32(56), v15+int32(55), v15+int32(48), v15+int32(44), v15+int32(40), v99)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	if base.Ui32(int32(5)) < base.Ui32(v115) {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	if int32(1)<<(uint(v115)%32)&int32(52) == int32(0) {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	if v125 != l0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	F_initStringInfo(m, v15-int32(-64))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	goto L1
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	F_errmsg_internal(m, int32(44633), v15)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(492166), int32(13278), int32(377667))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_initStringInfo(m, v15-int32(-64))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	if v147 == int32(0) {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v155 = F_quote_identifier(m, v147)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v155
	F_appendStringInfo(m, v15-int32(-64), int32(589640), v15+int32(16))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	goto L1
L51:
	;
	F_appendStringInfoString(m, v15-int32(-64), v169)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	F_ReleaseCatCache(m, v18)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	m.G0 = v15 + int32(80)
	return v173
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
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
		if v15 == int32(1) {
			F_appendStringInfoString(m, v8, int32(668783))
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
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
							F_appendBinaryStringInfo(m, l0, v53, v54)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
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
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
							F_appendBinaryStringInfo(m, l0, v53, v54)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						} else {
							v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+100)))
							if v37&int32(1) != 0 {
								F_appendStringInfoString(m, v8, int32(727543))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
									v48 = F_format_type_be(m, v47)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return
									} else {
										F_appendStringInfoString(m, v8, v48)
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return
										} else {
											v53 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
											v54 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
											F_appendBinaryStringInfo(m, l0, v53, v54)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return
											} else {
												m.G0 = v8 + int32(16)
												return
											}
										}
									}
								}
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
								v48 = F_format_type_be(m, v47)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									F_appendStringInfoString(m, v8, v48)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
										F_appendBinaryStringInfo(m, l0, v53, v54)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
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
		} else {
			if v15 == int32(0) {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
				v48 = F_format_type_be(m, v47)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					F_appendStringInfoString(m, v8, v48)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
						F_appendBinaryStringInfo(m, l0, v53, v54)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			} else {
				F_appendStringInfoString(m, v8, int32(727543))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
					v48 = F_format_type_be(m, v47)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						F_appendStringInfoString(m, v8, v48)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
							F_appendBinaryStringInfo(m, l0, v53, v54)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
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
