package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FunctionCall2Coll(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64) int64 {
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
	var v27 int64
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
	v8 = v6 + int32(-64)
	m.G0 = v8
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+56)) = uint8(v5)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+40)) = uint8(v5)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = l2
	v16 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+26)) = uint16(v16)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+24)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l0
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = m.T0[v26].(func(*base.Module, int32) int64)(m, v6+int32(-56))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int64(0)
	} else {
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+24)))
		if v31 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int64(0)
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v38
				F_errmsg_internal(m, int32(_a_F_FunctionCall2Coll_0), v8)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_FunctionCall2Coll_1), int32(1167), int32(_a_F_FunctionCall2Coll_2))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v8 - int32(-64)
			return v27
		}
	}
}
func F_FunctionCall8Coll(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64) int64 {
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
	var v51 int64
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
	v14 = v12 - int32(160)
	m.G0 = v14
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+152)) = uint8(v11)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+144)) = l9
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+136)) = uint8(v11)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+128)) = l8
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+120)) = uint8(v11)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+112)) = l7
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+104)) = uint8(v11)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+96)) = l6
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+88)) = uint8(v11)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+72)) = uint8(v11)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+56)) = uint8(v11)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+48)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+40)) = uint8(v11)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = l2
	v40 = int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+26)) = uint16(v40)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)) = uint8(v11)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v14)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = l0
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = m.T0[v50].(func(*base.Module, int32) int64)(m, v14+v40)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		return int64(0)
	} else {
		v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)))
		if v55 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int64(0)
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v62
				F_errmsg_internal(m, int32(_a_F_FunctionCall8Coll_0), v14)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_FunctionCall8Coll_1), int32(1350), int32(_a_F_FunctionCall8Coll_2))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v14 + int32(160)
			return v51
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	v5 = l4
	v8 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	v19 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(l0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v174 = F_quote_identifier(m, v27)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L3
	} else {
		goto L48
	}
L2:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v25)+68))
	v153 = F_get_namespace_name_or_temp(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L3
	} else {
		goto L43
	}
L3:
	;
	return int32(0)
L4:
	;
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
	v25 = v23 + v24
	v27 = v25 + int32(4)
	if l6 == int32(0) {
		v85 = v8
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
	v140 = m.ExcPending
	if v140 != 0 {
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
	v30 = int32(_a_F_generate_function_name_0)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generate_function_name[0])))
	if base.B2i32(v33 == int32(0))|base.B2i32(v33 != v36) != 0 {
		v54 = v33
		v55 = v36
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v54-v55 != 0 {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	goto L10
L12:
	;
	v39 = v27
	v40 = v30
	goto L13
L13:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	if v44 == int32(0) {
		v54 = v44
		v55 = v43
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v54 = v44
	v55 = v43
	goto L11
L15:
	;
	v47 = int32(1)
	if v44 == v43 {
		v39 = v39 + v47
		v40 = v40 + v47
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v57 = int32(_a_F_generate_function_name_1)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generate_function_name[1])))
	if base.B2i32(v60 == int32(0))|base.B2i32(v60 != v63) != 0 {
		v81 = v60
		v82 = v63
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	v85 = int32(1)
	goto L8
L20:
	;
	if v81-v82 != 0 {
		v85 = v8
		goto L8
	} else {
		goto L27
	}
L21:
	;
	goto L20
L22:
	;
	v66 = v27
	v67 = v57
	goto L23
L23:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v71 == int32(0) {
		v81 = v71
		v82 = v70
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v81 = v71
	v82 = v70
	goto L21
L25:
	;
	v74 = int32(1)
	if v71 == v70 {
		v66 = v66 + v74
		v67 = v67 + v74
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
	v90 = v5 ^ int32(1)
	goto L30
L29:
	;
	v90 = int32(1)
	goto L30
L30:
	;
	if v85 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = int32(0)
	goto L2
L32:
	;
	goto L33
L33:
	;
	v93 = F_makeString(m, v27)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v93
	v100 = F_list_make1_impl(m, int32(1), v15+int32(28))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v102 = int32(0)
	v120 = F_func_get_detail(m, v100, v102, l2, l1, l3, v90, int32(1), v102, v15+int32(60), v15+int32(56), v15+int32(52), v15+int32(51), v15+int32(44), v15+int32(40), v15+int32(36), v102)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v120))|base.B2i32(int32(1)<<(uint(v120)%32)&int32(52) == int32(0)) != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	if v131 != l0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	F_initStringInfo(m, v15-int32(-64))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
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
	v144 = m.ExcPending
	if v144 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_generate_function_name_3), int32(_a_F_generate_function_name_4), int32(_a_F_generate_function_name_5))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
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
	v156 = v15 - int32(-64)
	F_initStringInfo(m, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	if v153 == int32(0) {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v161 = F_quote_identifier(m, v153)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v161
	F_appendStringInfo(m, v156, int32(_a_F_generate_function_name_6), v15+int32(16))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	goto L1
L48:
	;
	F_appendStringInfoString(m, v15-int32(-64), v174)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	F_ReleaseCatCache(m, v19)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	m.G0 = v15 + int32(80)
	return v178
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
