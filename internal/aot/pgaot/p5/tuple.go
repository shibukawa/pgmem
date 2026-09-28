package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecCheckTupleVisible(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCheckTupleVisible[0]))
	if v10 < int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L49
	}
L2:
	;
	m.G0 = v7 + int32(16)
	return
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+188))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v16 = m.T0[v15].(func(*base.Module, int32, int32, int32) int32)(m, l1, l2, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if v16 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v23 = m.T0[v22].(func(*base.Module, int32, int32, int32) int64)(m, l2, int32(-2), v7+int32(15))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v25 = base.I32_wrap_i64(v23)
	if base.Ui32(v25) < base.Ui32(int32(3)) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v157 == int32(0) {
		goto L1
	} else {
		goto L48
	}
L9:
	;
	v157 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCheckTupleVisible[1]))
	if v37 == v25 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v157 = int32(1)
	goto L8
L13:
	;
	goto L14
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCheckTupleVisible[2]))
	if v41 <= int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v157 = v147
	goto L8
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCheckTupleVisible[3]))
	if v45 == int32(0) {
		v147 = int32(0)
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCheckTupleVisible[4]))
	v117 = int32(0)
	v120 = v41 - int32(1)
	goto L38
L19:
	;
	v50 = v45
	goto L20
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	if v56 == int32(4) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v147 = int32(0)
	goto L15
L22:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v50)+80))
	if v110 != 0 {
		v50 = v110
		goto L20
	} else {
		goto L37
	}
L23:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v59 == int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v62 = int32(1)
	if v25 == v59 {
		v147 = v62
		goto L15
	} else {
		goto L25
	}
L25:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v50)+52))
	v66 = v64 - int32(1)
	if v66 < int32(0) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	v72 = int32(0)
	v75 = v66
	goto L27
L27:
	;
	v80 = int32(2)
	v81 = base.I32_div_s(v75-v72, v80)
	v82 = v81 + v72
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v69+v82<<(uint(v80)%32))))
	if v86 == v25 {
		v147 = v62
		goto L15
	} else {
		goto L29
	}
L28:
	;
	goto L22
L29:
	;
	v95 = base.B2i32(v86-v25 < int32(0)) | base.B2i32(base.Ui32(v86) < base.Ui32(int32(3)))
	if v95 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v96 = v82 + int32(1)
	goto L32
L31:
	;
	v96 = v72
	goto L32
L32:
	;
	if v95 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v99 = v75
	goto L35
L34:
	;
	v99 = v82 - int32(1)
	goto L35
L35:
	;
	if v96 <= v99 {
		v72 = v96
		v75 = v99
		goto L27
	} else {
		goto L36
	}
L36:
	;
	goto L28
L37:
	;
	goto L21
L38:
	;
	v125 = int32(2)
	v126 = base.I32_div_s(v120-v117, v125)
	v127 = v126 + v117
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v115+v127<<(uint(v125)%32))))
	v132 = base.B2i32(v131 == v25)
	if v131 == v25 {
		v147 = v132
		goto L15
	} else {
		goto L40
	}
L39:
	;
	v147 = v132
	goto L15
L40:
	;
	v135 = base.B2i32(base.Ui32(v131) < base.Ui32(v25))
	if base.Ui32(v131) < base.Ui32(v25) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v136 = v127 + int32(1)
	goto L43
L42:
	;
	v136 = v117
	goto L43
L43:
	;
	if base.Ui32(v131) < base.Ui32(v25) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v139 = v120
	goto L46
L45:
	;
	v139 = v127 - int32(1)
	goto L46
L46:
	;
	if v136 <= v139 {
		v117 = v136
		v120 = v139
		goto L38
	} else {
		goto L47
	}
L47:
	;
	goto L39
L48:
	;
	goto L2
L49:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(_a_F_ExecCheckTupleVisible_0), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_ExecCheckTupleVisible_1), int32(403), int32(_a_F_ExecCheckTupleVisible_2))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecCleanupTupleRouting(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if int32(2) <= v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = int32(1)
	goto L4
L2:
	;
	goto L3
L3:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if int32(0) < v36 {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15+v12<<(uint(int32(2))%32))))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	F_relation_close(m, v20, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v24 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_ExecDropSingleTupleTableSlot(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v28 = v12 + int32(1)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v28 < v29 {
		v12 = v28
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	goto L5
L13:
	;
	v42 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	return
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v42<<(uint(int32(2))%32))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+84))
	if v50 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v42))))
	if v62 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+80))
	if v53 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	m.T0[v53].(func(*base.Module, int32, int32))(m, v56, v49)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	F_ExecCloseIndices(m, v49)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v72 = v42 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v72 < v73 {
		v42 = v72
		goto L16
	} else {
		goto L27
	}
L25:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	F_relation_close(m, v67, int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	goto L17
}
func F_GetTupleForTrigger(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
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
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if l7 != 0 {
		v18 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l7))) = v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_GetTupleForTrigger[0]))
		if v26 <= int32(1) {
			v29 = int32(2)
		} else {
			v29 = v18
		}
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+188))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+104))
		v34 = m.T0[v33].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v17, l3, v20, l5, v21, l4, v18, v29, v15+int32(28))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			if l8 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l8))) = v34
			} else {
			}
			if l9 != 0 {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
				*(*int32)(unsafe.Add(mBase, uint32(l9)+16)) = v39
				v41 = *(*int64)(unsafe.Add(mBase, uint32(v15)+36))
				*(*int64)(unsafe.Add(mBase, uint32(l9)+8)) = v41
				v43 = *(*int64)(unsafe.Add(mBase, uint32(v15)+28))
				*(*int64)(unsafe.Add(mBase, uint32(l9))) = v43
			} else {
			}
			switch v34 {
			case 0:
				v69 = int32(1)
				v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+44)))
				if v70 != v69 {
					v211 = v69
					m.G0 = v15 + int32(48)
					return v211
				} else {
					if l6 != 0 {
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
						v74 = F_EvalPlanQual(m, l1, v17, v73, l5)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l7))) = v74
							if v74 != 0 {
								v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+4)))
								if v77&int32(2) == int32(0) {
									v211 = v69
								} else {
									v82 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l7))) = v82
									v211 = v82
								}
							} else {
								v82 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l7))) = v82
								v211 = v82
							}
							m.G0 = v15 + int32(48)
							return v211
						}
					} else {
						v85 = int32(0)
						if l8 == v85 {
							v211 = v85
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l8))) = int32(3)
							v211 = v85
						}
						m.G0 = v15 + int32(48)
						return v211
					}
				}
			case 1:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_GetTupleForTrigger_0), int32(0))
					mBase = m.M
					v138 = m.ExcPending
					if v138 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_GetTupleForTrigger_1), int32(3472), int32(_a_F_GetTupleForTrigger_2))
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			case 2:
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				if v46 == v47 {
					v211 = int32(0)
					m.G0 = v15 + int32(48)
					return v211
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(450))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_GetTupleForTrigger_3), int32(0))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								F_errhint(m, int32(_a_F_GetTupleForTrigger_4), int32(0))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_GetTupleForTrigger_1), int32(3417), int32(_a_F_GetTupleForTrigger_2))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
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
			case 3:
				v91 = *(*int32)(unsafe.Add(mBase, _c_F_GetTupleForTrigger[0]))
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return int32(0)
				} else {
					if int32(2) <= v91 {
						F_errcode(m, int32(16777220))
						mBase = m.M
						v186 = m.ExcPending
						if v186 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_GetTupleForTrigger_5), int32(0))
							mBase = m.M
							v190 = m.ExcPending
							if v190 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_GetTupleForTrigger_1), int32(3459), int32(_a_F_GetTupleForTrigger_2))
								mBase = m.M
								v195 = m.ExcPending
								if v195 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(3)
						F_errmsg_internal(m, int32(_a_F_GetTupleForTrigger_6), v15+int32(16))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_GetTupleForTrigger_1), int32(3460), int32(_a_F_GetTupleForTrigger_2))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 4:
				v112 = *(*int32)(unsafe.Add(mBase, _c_F_GetTupleForTrigger[0]))
				if v112 < int32(2) {
					v211 = int32(0)
					m.G0 = v15 + int32(48)
					return v211
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(16777220))
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_GetTupleForTrigger_7), int32(0))
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_GetTupleForTrigger_1), int32(3467), int32(_a_F_GetTupleForTrigger_2))
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v147 = m.ExcPending
				if v147 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v34
					F_errmsg_internal(m, int32(_a_F_GetTupleForTrigger_8), v15)
					mBase = m.M
					v151 = m.ExcPending
					if v151 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_GetTupleForTrigger_1), int32(3476), int32(_a_F_GetTupleForTrigger_2))
						mBase = m.M
						v156 = m.ExcPending
						if v156 != 0 {
							return int32(0)
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
		v158 = *(*int32)(unsafe.Add(mBase, _c_F_GetTupleForTrigger[1]))
		if v158 != 0 {
			v160 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetTupleForTrigger[2])))
			if v160&int32(1) == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v199 = m.ExcPending
				if v199 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_GetTupleForTrigger_9), int32(0))
					mBase = m.M
					v203 = m.ExcPending
					if v203 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_GetTupleForTrigger_10), int32(1355), int32(_a_F_GetTupleForTrigger_11))
						mBase = m.M
						v208 = m.ExcPending
						if v208 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v167 = *(*int32)(unsafe.Add(mBase, uint32(v17)+188))
				v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+60))
				v169 = m.T0[v168].(func(*base.Module, int32, int32, int32, int32) int32)(m, v17, l3, int32(_a_F_GetTupleForTrigger_12), l5)
				mBase = m.M
				v170 = m.ExcPending
				if v170 != 0 {
					return int32(0)
				} else {
					if v169 != 0 {
						v211 = int32(1)
						m.G0 = v15 + int32(48)
						return v211
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v174 = m.ExcPending
						if v174 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_GetTupleForTrigger_13), int32(0))
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_GetTupleForTrigger_1), int32(3488), int32(_a_F_GetTupleForTrigger_2))
								mBase = m.M
								v183 = m.ExcPending
								if v183 != 0 {
									return int32(0)
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
		} else {
			v167 = *(*int32)(unsafe.Add(mBase, uint32(v17)+188))
			v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+60))
			v169 = m.T0[v168].(func(*base.Module, int32, int32, int32, int32) int32)(m, v17, l3, int32(_a_F_GetTupleForTrigger_12), l5)
			mBase = m.M
			v170 = m.ExcPending
			if v170 != 0 {
				return int32(0)
			} else {
				if v169 != 0 {
					v211 = int32(1)
					m.G0 = v15 + int32(48)
					return v211
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v174 = m.ExcPending
					if v174 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_GetTupleForTrigger_13), int32(0))
						mBase = m.M
						v178 = m.ExcPending
						if v178 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_GetTupleForTrigger_1), int32(3488), int32(_a_F_GetTupleForTrigger_2))
							mBase = m.M
							v183 = m.ExcPending
							if v183 != 0 {
								return int32(0)
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
func F_TupleDescFinalize(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v2 < v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = l0 + int32(28)
	v22 = v2
	v23 = v11
	v25 = v2
	goto L5
L2:
	;
	v79 = v2
	v86 = v11
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v79
	return
L4:
	;
	v79 = v73
	v86 = v52
	goto L3
L5:
	;
	v31 = v15 + v11<<(uint(int32(3))%32) + v22*int32(100)
	v34 = v15 + v22<<(uint(int32(3))%32)
	if v11 != v23 {
		v52 = v23
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v73 = v11
	goto L4
L7:
	;
	v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34)+2)))
	if v53 <= int32(0) {
		v73 = v22
		goto L4
	} else {
		goto L15
	}
L8:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+7)))
	if v36 != int32(118) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v52 = v22
	goto L7
L10:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+4)))
	if v39 != int32(1) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+6)))
	if v42&int32(6) != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34)+2)))
	if v45 <= int32(0) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+90)))
	if v48 != int32(118) {
		v52 = v11
		goto L7
	} else {
		goto L14
	}
L14:
	;
	goto L9
L15:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+90)))
	if v56 == int32(118) {
		v73 = v22
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+5)))
	v65 = (v25 + v59 - int32(1)) & (int32(0) - v59)
	if int32(_a_F_TupleDescFinalize_0) < v65 {
		v73 = v22
		goto L4
	} else {
		goto L17
	}
L17:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v34))) = uint16(v65)
	v71 = v22 + int32(1)
	if v71 != v11 {
		v22 = v71
		v23 = v52
		v25 = v65 + v53
		goto L5
	} else {
		goto L18
	}
L18:
	;
	goto L6
}
func F_load_tuple_array(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v13 = v10 + l1<<(uint(int32(4))%32)
	v15 = v13 - int32(12)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v18 = v13 - int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v16 == v19 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v38 = v28
	goto L8
L4:
	;
	v21 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v21
	v28 = v21
	goto L3
L5:
	;
	goto L6
L6:
	;
	if int32(9) < v16 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v28 = v16
	goto L3
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_load_tuple_array[0]))
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L1
L10:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+l1<<(uint(int32(2))%32)-int32(4))))
	v52 = F_TupleQueueReaderNext(m, v50, int32(1), v13-int32(4))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L13
	} else {
		goto L15
	}
L13:
	;
	return
L14:
	;
	goto L12
L15:
	;
	if v52 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v57 = F_heap_copy_minimal_tuple(m, v52, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	if v57 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v13-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v61+v38<<(uint(int32(2))%32)))) = v57
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v67 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v66 + v67
	v71 = v38 + v67
	if v71 != int32(10) {
		v38 = v71
		goto L8
	} else {
		goto L19
	}
L19:
	;
	goto L9
}
func F_make_tuple_from_row(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v16 != v17 {
		v111 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v111
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v23 = F_MemoryContextAllocZero(m, v20, v16<<(uint(int32(3))%32))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v29 = F_MemoryContextAlloc(m, v28, v16)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if int32(0) < v16 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v37 = v4
	v42 = v33
	goto L9
L7:
	;
	goto L8
L8:
	;
	v99 = F_heap_form_tuple(m, l2, v23, v29)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L3
	} else {
		goto L18
	}
L9:
	;
	v45 = v37 * int32(100)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+(l2+v42<<(uint(int32(3))%32)))+119)))
	if v50 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L8
L11:
	;
	v87 = v37 + int32(1)
	if v87 != v16 {
		v37 = v87
		v42 = v85
		goto L9
	} else {
		goto L17
	}
L12:
	;
	v54 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v37+v29))) = uint8(v54)
	v85 = v42
	goto L11
L13:
	;
	goto L14
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v58 = int32(2)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v37<<(uint(v58)%32))))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v56+v61<<(uint(v58)%32))))
	F_exec_eval_datum(m, l0, v65, v13+int32(12), v13+int32(8), v23+v37<<(uint(int32(3))%32), v37+v29)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2+v77<<(uint(int32(3))%32)+v45)+96))
	if v76 == v82 {
		v85 = v77
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v111 = int32(0)
	goto L1
L17:
	;
	goto L10
L18:
	;
	v111 = v99
	goto L1
}
