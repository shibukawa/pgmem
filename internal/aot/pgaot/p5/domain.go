package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_domain_check(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	v8 = *(*int32)(unsafe.Add(mBase, _consts[9]))
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
func F_domain_check_safe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
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
	v9 = *(*int32)(unsafe.Add(mBase, _consts[9]))
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
						if v30 != int32(447) {
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
							if v30 != int32(447) {
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
						if v30 != int32(447) {
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
					if v30 != int32(447) {
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v23 = F_stringToNode(m, l1)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v25 = F_CreateExecutorState(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+152))
	if v27 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v30 = F_MakePerTupleExprContext(m, v25)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v32 = v27
	goto L6
L6:
	;
	v33 = F_ExecPrepareExpr(m, v23, v25)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v32 = v30
	goto L6
L8:
	;
	v35 = F_get_rels_with_domain(m, l0)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L54
	}
L10:
	;
	F_FreeExecutorState(m, v25)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L53
	}
L11:
	;
	if v35 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v39 <= int32(0) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v50 = int32(0)
	goto L14
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v50<<(uint(int32(2))%32))))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+52))
	v67 = F_GetLatestSnapshot(m)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L10
L16:
	;
	v69 = F_RegisterSnapshot(m, v67)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v71 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v65)+188))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	v77 = m.T0[v76].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v65, v69, v71, v71, v71, int32(449))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v80 = F_table_slot_create(m, v65, int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+36)) = v83
	v86 = *(*int32)(unsafe.Add(mBase, _consts[246]))
	if v86 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	F_ExecDropSingleTupleTableSlot(m, v80)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L48
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L45
	}
L22:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, _consts[247])))
	if v88&int32(1) == int32(0) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L26
L25:
	;
	goto L24
L26:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+188))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+20))
	v115 = m.T0[v114].(func(*base.Module, int32, int32, int32) int32)(m, v77, int32(1), v80)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	goto L21
L28:
	;
	if v115 == int32(0) {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v119 = int32(0)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v119 < v120 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v128 = v119
	goto L33
L31:
	;
	goto L32
L32:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	F_MemoryContextReset(m, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L42
	}
L33:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+v128<<(uint(int32(2))%32))))
	v147 = v145 - int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v149 = int32(*(*int16)(unsafe.Add(mBase, uint32(v80)+6)))
	if v149 < v145 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	F_slot_getsomeattrs_int(m, v80, v145)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153+v147))))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+15)) = uint8(v155)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157+v147<<(uint(int32(2))%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+52)) = uint8(v155)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v161
	v164 = int32(4455216)
	v165 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v167
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	v172 = m.T0[v171].(func(*base.Module, int32, int32, int32) int32)(m, v33, v32, v21+int32(15))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v165
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+15)))
	if v176|v172 == int32(0) {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	v181 = v128 + int32(1)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v181 < v182 {
		v128 = v181
		goto L33
	} else {
		goto L41
	}
L41:
	;
	goto L34
L42:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+36)) = v206
	v209 = *(*int32)(unsafe.Add(mBase, _consts[246]))
	if v209 == int32(0) {
		goto L26
	} else {
		goto L43
	}
L43:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, _consts[247])))
	if v213&int32(1) != 0 {
		goto L26
	} else {
		goto L44
	}
L44:
	;
	goto L27
L45:
	;
	F_errmsg_internal(m, int32(323972), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(314654), int32(1034), int32(81516))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+188))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+12))
	m.T0[v251].(func(*base.Module, int32))(m, v77)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_UnregisterSnapshot(m, v69)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_sequence_close(m, v65, int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v260 = v50 + int32(1)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v260 < v261 {
		v50 = v260
		goto L14
	} else {
		goto L52
	}
L52:
	;
	goto L15
L53:
	;
	m.G0 = v21 + int32(16)
	return
L54:
	;
	F_errcode(m, int32(67391682))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v65)+48))
	v294 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v293 + v294
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v66 + v148<<(uint(v294)%32) + v147*int32(100) + int32(24)
	F_errmsg(m, int32(86446), v21)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errtablecol(m, v65, v145)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(477716), int32(3267), int32(87515))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
