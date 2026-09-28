package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_domain_check(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_domain_check[0]))
	if l4 != 0 {
		v9 = l4
	} else {
		v9 = v8
	}
	if l3 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		if v10 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if v11 == l2 {
				v20 = v10
				F_domain_check_input(m, l0, l1, v20, int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					return
				}
			} else {
				v14 = F_domain_state_setup(m, l2, int32(1), v9)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v14
					v20 = v14
					F_domain_check_input(m, l0, l1, v20, int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			v14 = F_domain_state_setup(m, l2, int32(1), v9)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v14
				v20 = v14
				F_domain_check_input(m, l0, l1, v20, int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v18 = F_domain_state_setup(m, l2, int32(1), v9)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v20 = v18
			F_domain_check_input(m, l0, l1, v20, int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_domain_check_safe(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_domain_check_safe[0]))
	if l4 != 0 {
		v10 = l4
	} else {
		v10 = v9
	}
	if l3 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		if v11 != 0 {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			if v12 == l2 {
				v23 = v11
				F_domain_check_input(m, l0, l1, v23, l5)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					if l5 == int32(0) {
						v36 = int32(1)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
						if v30 != int32(453) {
							v36 = int32(1)
						} else {
							v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+4)))
							v36 = v33 ^ int32(1)
						}
					}
					return v36 & int32(1)
				}
			} else {
				v15 = F_domain_state_setup(m, l2, int32(1), v10)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v15
					v23 = v15
					F_domain_check_input(m, l0, l1, v23, l5)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						if l5 == int32(0) {
							v36 = int32(1)
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
							if v30 != int32(453) {
								v36 = int32(1)
							} else {
								v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+4)))
								v36 = v33 ^ int32(1)
							}
						}
						return v36 & int32(1)
					}
				}
			}
		} else {
			v15 = F_domain_state_setup(m, l2, int32(1), v10)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v15
				v23 = v15
				F_domain_check_input(m, l0, l1, v23, l5)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					if l5 == int32(0) {
						v36 = int32(1)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
						if v30 != int32(453) {
							v36 = int32(1)
						} else {
							v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+4)))
							v36 = v33 ^ int32(1)
						}
					}
					return v36 & int32(1)
				}
			}
		}
	} else {
		v21 = F_domain_state_setup(m, l2, int32(1), v10)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = v21
			F_domain_check_input(m, l0, l1, v23, l5)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if l5 == int32(0) {
					v36 = int32(1)
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
					if v30 != int32(453) {
						v36 = int32(1)
					} else {
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+4)))
						v36 = v33 ^ int32(1)
					}
				}
				return v36 & int32(1)
			}
		}
	}
}
func F_validateDomainCheckConstraint(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int64
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int64
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = F_stringToNode(m, l1)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v24 = F_CreateExecutorState(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+152))
	if v26 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v29 = F_MakePerTupleExprContext(m, v24)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v31 = v26
	goto L6
L6:
	;
	v32 = F_ExecPrepareExpr(m, v22, v24)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v31 = v29
	goto L6
L8:
	;
	v34 = F_get_rels_with_domain(m, l0)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L12
	}
L9:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L55
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L52
	}
L11:
	;
	F_FreeExecutorState(m, v24)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L51
	}
L12:
	;
	if v34 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v38 <= int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v53 = int32(0)
	goto L15
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v53<<(uint(int32(2))%32))))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v65 = F_GetLatestSnapshot(m)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L11
L17:
	;
	v67 = F_RegisterSnapshot(m, v65)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_validateDomainCheckConstraint[0]))
	if v70 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_validateDomainCheckConstraint[1])))
	if v72&int32(1) == int32(0) {
		goto L10
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v77 = int32(0)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v63)+188))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v83 = m.T0[v82].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v63, v67, v77, v77, v77, int32(449))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v86 = F_table_slot_create(m, v63, int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+40)) = v89
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+188))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v95 = m.T0[v94].(func(*base.Module, int32, int32, int32) int32)(m, v83, int32(1), v86)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v95 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	goto L29
L27:
	;
	goto L28
L28:
	;
	F_ExecDropSingleTupleTableSlot(m, v86)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L46
	}
L29:
	;
	v114 = int32(0)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v114 < v115 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L28
L31:
	;
	v118 = v114
	goto L34
L32:
	;
	goto L33
L33:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	F_MemoryContextReset(m, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L43
	}
L34:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135+v118<<(uint(int32(2))%32))))
	v140 = int32(*(*int16)(unsafe.Add(mBase, uint32(v86)+6)))
	if v140 < v139 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L33
L36:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
	m.T0[v143].(func(*base.Module, int32, int32))(m, v86, v139)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v147 = v139 - int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+v148))))
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+15)) = uint8(v150)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v152+v147<<(uint(int32(3))%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+64)) = uint8(v150)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+56)) = v156
	v159 = int32(_a_F_validateDomainCheckConstraint_0)
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_validateDomainCheckConstraint[2]))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_validateDomainCheckConstraint[2])) = v162
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v167 = m.T0[v166].(func(*base.Module, int32, int32, int32) int64)(m, v32, v31, v20+int32(15))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_validateDomainCheckConstraint[2])) = v160
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+15)))
	if base.B2i32(v171 == int32(0))&base.B2i32(v167 == int64(0)) != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	v178 = v118 + int32(1)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v178 < v179 {
		v118 = v178
		goto L34
	} else {
		goto L42
	}
L42:
	;
	goto L35
L43:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+40)) = v202
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+188))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+20))
	v208 = m.T0[v207].(func(*base.Module, int32, int32, int32) int32)(m, v83, int32(1), v86)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v208 != 0 {
		goto L29
	} else {
		goto L45
	}
L45:
	;
	goto L30
L46:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+188))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+12))
	m.T0[v231].(func(*base.Module, int32))(m, v83)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_UnregisterSnapshot(m, v67)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_relation_close(m, v63, int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v240 = v53 + int32(1)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v240 < v241 {
		v53 = v240
		goto L15
	} else {
		goto L50
	}
L50:
	;
	goto L16
L51:
	;
	m.G0 = v20 + int32(16)
	return
L52:
	;
	F_errmsg_internal(m, int32(_a_F_validateDomainCheckConstraint_1), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_validateDomainCheckConstraint_2), int32(931), int32(_a_F_validateDomainCheckConstraint_3))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errcode(m, int32(67391682))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v63)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v286 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v64 + v278<<(uint(int32(3))%32) + v147*int32(100) + int32(32)
	F_errmsg(m, int32(_a_F_validateDomainCheckConstraint_4), v20)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errtablecol(m, v63, v139)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_validateDomainCheckConstraint_5), int32(3327), int32(_a_F_validateDomainCheckConstraint_6))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
