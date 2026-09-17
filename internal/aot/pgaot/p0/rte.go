package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_markRTEForSelectPriv(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v15 = l1 - int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13+v15<<(uint(int32(2))%32))))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v20 == int32(0) {
		v128 = l2
		v132 = v19
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L14
	} else {
		goto L40
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L14
	} else {
		goto L37
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L14
	} else {
		goto L34
	}
L4:
	;
	m.G0 = v10 + int32(32)
	return
L5:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v134 = F_getRTEPermissionInfo(m, v133, v132)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L14
	} else {
		goto L32
	}
L6:
	;
	if l2|base.B2i32(v20 != int32(2)) != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if l1 <= int32(0) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v28 == int32(0) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v31 < l1 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v15<<(uint(int32(2))%32))))
	if v37 == int32(0) {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v40 = int32(36)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	switch v43 - int32(63) {
	case 0:
		goto L13
	case 1:
		v47 = v40
		goto L12
	default:
		v201 = v37
		goto L1
	}
L12:
	;
	v48 = int32(0)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47+v42)))
	F_markRTEForSelectPriv(m, l0, v50, v48)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v47 = int32(4)
	goto L12
L14:
	;
	return
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	switch v55 - int32(63) {
	case 0:
		goto L17
	case 1:
		v59 = v40
		goto L16
	default:
		v177 = v37
		goto L2
	}
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59+v54)))
	v65 = v63 - int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v61+v65<<(uint(int32(2))%32))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	if v70 == int32(0) {
		v128 = v48
		v132 = v69
		goto L5
	} else {
		goto L18
	}
L17:
	;
	v59 = int32(4)
	goto L16
L18:
	;
	v74 = v63
	v76 = v70
	v77 = v65
	goto L19
L19:
	;
	if v76 != int32(2) {
		goto L4
	} else {
		goto L21
	}
L20:
	;
	v128 = v48
	v132 = v124
	goto L5
L21:
	;
	if v74 <= int32(0) {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v84 == int32(0) {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v87 < v74 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v77<<(uint(int32(2))%32))))
	if v93 == int32(0) {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	switch v98 - int32(63) {
	case 0:
		v102 = int32(4)
		goto L26
	case 1:
		goto L27
	default:
		v201 = v93
		goto L1
	}
L26:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v102+v97)))
	F_markRTEForSelectPriv(m, l0, v104, int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L14
	} else {
		goto L28
	}
L27:
	;
	v102 = int32(36)
	goto L26
L28:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	switch v110 - int32(63) {
	case 0:
		goto L30
	case 1:
		v114 = int32(36)
		goto L29
	default:
		v177 = v93
		goto L2
	}
L29:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114+v109)))
	v120 = v118 - int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v116+v120<<(uint(int32(2))%32))))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	if v125 != 0 {
		v74 = v118
		v76 = v125
		v77 = v120
		goto L19
	} else {
		goto L31
	}
L30:
	;
	v114 = int32(4)
	goto L29
L31:
	;
	goto L20
L32:
	;
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v134)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v134)+16)) = v136 | int64(2)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v134)+28))
	v143 = F_bms_add_member(m, v140, v128+int32(7))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L14
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+28)) = v143
	goto L4
L34:
	;
	F_errmsg_internal(m, int32(_a_F_markRTEForSelectPriv_0), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L14
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_markRTEForSelectPriv_1), int32(1112), int32(_a_F_markRTEForSelectPriv_2))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L14
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v188
	F_errmsg_internal(m, int32(_a_F_markRTEForSelectPriv_3), v10+int32(16))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L14
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_markRTEForSelectPriv_1), int32(1144), int32(_a_F_markRTEForSelectPriv_2))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v212
	F_errmsg_internal(m, int32(_a_F_markRTEForSelectPriv_3), v10)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_markRTEForSelectPriv_1), int32(1129), int32(_a_F_markRTEForSelectPriv_2))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L14
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_replace_rte_variables(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l3
	if l0 == int32(0) {
		v22 = int32(0)
		if l5 == v22 {
			v26 = v22
		} else {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5))))
			v26 = v25
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v18 != int32(67) {
			v22 = int32(0)
			if l5 == v22 {
				v26 = v22
			} else {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5))))
				v26 = v25
			}
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
			v26 = v21
		}
	}
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)) = uint8(v26)
	v31 = F_query_or_expression_tree_mutator_impl(m, l0, int32(1054), v10+int32(12))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return int32(0)
	} else {
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
		if v35 != int32(1) {
			m.G0 = v10 + int32(32)
			return v31
		} else {
			if v31 == int32(0) {
				if l5 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_replace_rte_variables_0), int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_replace_rte_variables_1), int32(1475), int32(_a_F_replace_rte_variables_2))
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
				} else {
					v47 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v47)
					m.G0 = v10 + int32(32)
					return v31
				}
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
				if v40 != int32(67) {
					if l5 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_replace_rte_variables_0), int32(0))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_replace_rte_variables_1), int32(1475), int32(_a_F_replace_rte_variables_2))
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
					} else {
						v47 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v47)
						m.G0 = v10 + int32(32)
						return v31
					}
				} else {
					v43 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v31)+39)) = uint8(v43)
					m.G0 = v10 + int32(32)
					return v31
				}
			}
		}
	}
}
