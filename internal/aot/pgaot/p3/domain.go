package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_domain_recv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v50 int64
	_ = v50
	v2 = int32(0)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v7 == v2 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v11 = v10
	} else {
		v11 = v2
	}
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v12 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
		if v17 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			if v18 == v15 {
				v28 = v17
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
				v33 = F_ReceiveFunctionCall(m, v28+int32(16), v11, v31, v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int64(0)
				} else {
					v35 = int32(0)
					F_domain_check_input(m, v33, base.B2i32(v11 == v35), v28, v35)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int64(0)
					} else {
						if v11 != 0 {
							v50 = v33
						} else {
							v44 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v44)
							v50 = int64(0)
						}
						return v50
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
				v22 = F_domain_state_setup(m, v15, int32(1), v21)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v22
					v28 = v22
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
					v33 = F_ReceiveFunctionCall(m, v28+int32(16), v11, v31, v32)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int64(0)
					} else {
						v35 = int32(0)
						F_domain_check_input(m, v33, base.B2i32(v11 == v35), v28, v35)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							if v11 != 0 {
								v50 = v33
							} else {
								v44 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v44)
								v50 = int64(0)
							}
							return v50
						}
					}
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
			v22 = F_domain_state_setup(m, v15, int32(1), v21)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int64(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v22
				v28 = v22
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
				v33 = F_ReceiveFunctionCall(m, v28+int32(16), v11, v31, v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int64(0)
				} else {
					v35 = int32(0)
					F_domain_check_input(m, v33, base.B2i32(v11 == v35), v28, v35)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int64(0)
					} else {
						if v11 != 0 {
							v50 = v33
						} else {
							v44 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v44)
							v50 = int64(0)
						}
						return v50
					}
				}
			}
		}
	} else {
		v44 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v44)
		v50 = int64(0)
		return v50
	}
}
func F_replace_domain_constraint_value(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v5 == v3 {
		v48 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v48
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v8 != int32(1) {
		v48 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v14 = int32(_a_F_replace_domain_constraint_value_0)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_replace_domain_constraint_value[0])))
	if base.B2i32(v17 == int32(0))|base.B2i32(v17 != v20) != 0 {
		v38 = v17
		v39 = v20
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v38-v39 != 0 {
		v48 = v3
		goto L1
	} else {
		goto L11
	}
L5:
	;
	goto L4
L6:
	;
	v23 = v13
	v24 = v14
	goto L7
L7:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v28 == int32(0) {
		v38 = v28
		v39 = v27
		goto L5
	} else {
		goto L9
	}
L8:
	;
	v38 = v28
	v39 = v27
	goto L5
L9:
	;
	v31 = int32(1)
	if v28 == v27 {
		v23 = v23 + v31
		v24 = v24 + v31
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v42 = F_copyObjectImpl(m, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v46
	v48 = v42
	goto L1
}
func F_validateDomainNotNullConstraint(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
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
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = F_get_rels_with_domain(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return
L2:
	;
	return
L3:
	;
	if v17 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v21 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v32 = int32(0)
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v32<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+52))
	v43 = F_GetLatestSnapshot(m)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	goto L1
L8:
	;
	v45 = F_RegisterSnapshot(m, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_validateDomainNotNullConstraint[0]))
	if v48 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_ExecDropSingleTupleTableSlot(m, v64)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L2
	} else {
		goto L44
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L2
	} else {
		goto L41
	}
L12:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_validateDomainNotNullConstraint[1])))
	if v50&int32(1) == int32(0) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v55 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v41)+188))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v61 = m.T0[v60].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v41, v45, v55, v55, v55, int32(449))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v64 = F_table_slot_create(m, v41, int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v67
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+188))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	v73 = m.T0[v72].(func(*base.Module, int32, int32, int32) int32)(m, v61, int32(1), v64)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	if v73 == int32(0) {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	goto L20
L20:
	;
	v89 = int32(0)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v90 <= v89 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L10
L22:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v171
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+188))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
	v177 = m.T0[v176].(func(*base.Module, int32, int32, int32) int32)(m, v61, int32(1), v64)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L39
	}
L23:
	;
	v100 = v89
	goto L24
L24:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106+v100<<(uint(int32(2))%32))))
	v111 = int32(*(*int16)(unsafe.Add(mBase, uint32(v64)+6)))
	if v111 < v110 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L2
	} else {
		goto L34
	}
L26:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+16))
	m.T0[v114].(func(*base.Module, int32, int32))(m, v64, v110)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L2
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	v119 = int32(1)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+v110-v119))))
	if v121 != v119 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	v125 = v100 + int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v126 <= v125 {
		goto L22
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L25
L33:
	;
	v100 = v125
	goto L24
L34:
	;
	F_errcode(m, int32(33575106))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v41)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v135 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v42 + v105<<(uint(int32(3))%32) + v110*int32(100) - int32(68)
	F_errmsg(m, int32(_a_F_validateDomainNotNullConstraint_0), v15)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	F_errtablecol(m, v41, v110)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_validateDomainNotNullConstraint_1), int32(3236), int32(_a_F_validateDomainNotNullConstraint_2))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	if v177 != 0 {
		goto L20
	} else {
		goto L40
	}
L40:
	;
	goto L21
L41:
	;
	F_errmsg_internal(m, int32(_a_F_validateDomainNotNullConstraint_3), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_validateDomainNotNullConstraint_4), int32(931), int32(_a_F_validateDomainNotNullConstraint_5))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+188))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
	m.T0[v208].(func(*base.Module, int32))(m, v61)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	F_UnregisterSnapshot(m, v45)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	F_relation_close(m, v41, int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v217 = v32 + int32(1)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v217 < v218 {
		v32 = v217
		goto L6
	} else {
		goto L48
	}
L48:
	;
	goto L7
}
