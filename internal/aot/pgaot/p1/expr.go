package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecPrepareExpr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v4 = int32(4515712)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v7
	v9 = F_expression_planner(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = F_ExecInitExpr(m, v9, int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[3])) = v5
			return v14
		}
	}
}
func F_ExecPrepareExprList(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	v3 = int32(0)
	v8 = int32(4515712)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v11
	if l0 == v3 {
		v52 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v9
	return v52
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v52 = v3
	goto L1
L4:
	;
	goto L5
L5:
	;
	v20 = v3
	v21 = v3
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v21<<(uint(int32(2))%32))))
	v30 = int32(4515712)
	v31 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v33
	v35 = F_expression_planner(m, v29)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v52 = v44
	goto L1
L8:
	;
	return int32(0)
L9:
	;
	v40 = F_ExecInitExpr(m, v35, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v31
	v44 = F_lappend(m, v20, v40)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v47 = v21 + int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v47 < v48 {
		v20 = v44
		v21 = v47
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L7
}
func F_checkExprHasSubLink(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_query_or_expression_tree_walker_impl(m, l0, int32(1047), int32(0), int32(3))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_evaluate_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = F_CreateExecutorState(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(4515712)
		v19 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+100))
		*(*int32)(unsafe.Add(mBase, _consts[3])) = v21
		F_fix_opfuncids(m, l0)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v26 = F_ExecInitExpr(m, l0, int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v14)+152))
				if v28 == int32(0) {
					v31 = F_MakePerTupleExprContext(m, v14)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = v31
						v34 = int32(4515712)
						v35 = *(*int32)(unsafe.Add(mBase, _consts[3]))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
						*(*int32)(unsafe.Add(mBase, _consts[3])) = v37
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
						v42 = m.T0[v41].(func(*base.Module, int32, int32, int32) int32)(m, v26, v33, v12+int32(15))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[3])) = v35
							F_get_typlenbyval(m, l1, v12+int32(12), v12+int32(11))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[3])) = v19
								v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
								if v54 != 0 {
									v63 = v42
									F_FreeExecutorState(m, v14)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+12)))
										v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
										v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
										v70 = F_makeConst(m, l1, l2, l3, v67, v63, v68, v69)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											m.G0 = v12 + int32(16)
											return v70
										}
									}
								} else {
									v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+12)))
									if v55 == int32(-1) {
										v58 = F_pg_detoast_datum_copy(m, v42)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											v63 = v58
											F_FreeExecutorState(m, v14)
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+12)))
												v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
												v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
												v70 = F_makeConst(m, l1, l2, l3, v67, v63, v68, v69)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													m.G0 = v12 + int32(16)
													return v70
												}
											}
										}
									} else {
										v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
										v61 = F_datumCopy(m, v42, v60, v55)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											v63 = v61
											F_FreeExecutorState(m, v14)
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+12)))
												v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
												v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
												v70 = F_makeConst(m, l1, l2, l3, v67, v63, v68, v69)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													m.G0 = v12 + int32(16)
													return v70
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v33 = v28
					v34 = int32(4515712)
					v35 = *(*int32)(unsafe.Add(mBase, _consts[3]))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
					*(*int32)(unsafe.Add(mBase, _consts[3])) = v37
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
					v42 = m.T0[v41].(func(*base.Module, int32, int32, int32) int32)(m, v26, v33, v12+int32(15))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[3])) = v35
						F_get_typlenbyval(m, l1, v12+int32(12), v12+int32(11))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[3])) = v19
							v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
							if v54 != 0 {
								v63 = v42
								F_FreeExecutorState(m, v14)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+12)))
									v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
									v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
									v70 = F_makeConst(m, l1, l2, l3, v67, v63, v68, v69)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										m.G0 = v12 + int32(16)
										return v70
									}
								}
							} else {
								v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+12)))
								if v55 == int32(-1) {
									v58 = F_pg_detoast_datum_copy(m, v42)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										v63 = v58
										F_FreeExecutorState(m, v14)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+12)))
											v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
											v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
											v70 = F_makeConst(m, l1, l2, l3, v67, v63, v68, v69)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												m.G0 = v12 + int32(16)
												return v70
											}
										}
									}
								} else {
									v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
									v61 = F_datumCopy(m, v42, v60, v55)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										v63 = v61
										F_FreeExecutorState(m, v14)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+12)))
											v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
											v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
											v70 = F_makeConst(m, l1, l2, l3, v67, v63, v68, v69)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												m.G0 = v12 + int32(16)
												return v70
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
func F_fix_expr_common(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v6 - int32(7) {
	case 0:
		goto L3
	default:
		goto L1
	case 2:
		goto L10
	case 3:
		goto L2
	case 4:
		goto L9
	case 8:
		goto L8
	case 10:
		goto L7
	case 11:
		goto L6
	case 12:
		goto L5
	case 13:
		goto L4
	}
L1:
	;
	return
L2:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v206 == int32(0) {
		goto L1
	} else {
		goto L59
	}
L3:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(v192 != int32(2205))&base.B2i32(v192 != int32(26)) != 0 {
		goto L1
	} else {
		goto L56
	}
L4:
	;
	F_set_opfuncid(m, l1)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L12
	} else {
		goto L39
	}
L5:
	;
	F_set_opfuncid(m, l1)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L12
	} else {
		goto L34
	}
L6:
	;
	F_set_opfuncid(m, l1)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L12
	} else {
		goto L29
	}
L7:
	;
	F_set_opfuncid(m, l1)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L12
	} else {
		goto L24
	}
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v47) < base.Ui32(int32(12000)) {
		goto L1
	} else {
		goto L20
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v28) < base.Ui32(int32(12000)) {
		goto L1
	} else {
		goto L16
	}
L10:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v9) < base.Ui32(int32(12000)) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v13 = F_palloc0(m, int32(12))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = int64(201863463291)
	v19 = F_GetSysCacheHashValue(m, int32(47), v9, int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v19
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+60))
	v24 = F_lappend(m, v23, v13)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v24
	return
L16:
	;
	v32 = F_palloc0(m, int32(12))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32))) = int64(201863463291)
	v38 = F_GetSysCacheHashValue(m, int32(47), v28, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v38
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+60))
	v43 = F_lappend(m, v42, v32)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+60)) = v43
	return
L20:
	;
	v51 = F_palloc0(m, int32(12))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = int64(201863463291)
	v57 = F_GetSysCacheHashValue(m, int32(47), v47, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v57
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+60))
	v62 = F_lappend(m, v61, v51)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+60)) = v62
	return
L24:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v68) < base.Ui32(int32(12000)) {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v72 = F_palloc0(m, int32(12))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v72))) = int64(201863463291)
	v78 = F_GetSysCacheHashValue(m, int32(47), v68, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v78
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+60))
	v83 = F_lappend(m, v82, v72)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+60)) = v83
	return
L29:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v89) < base.Ui32(int32(12000)) {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v93 = F_palloc0(m, int32(12))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v93))) = int64(201863463291)
	v99 = F_GetSysCacheHashValue(m, int32(47), v89, int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L12
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v99
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+60))
	v104 = F_lappend(m, v103, v93)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v106)+60)) = v104
	return
L34:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v110) < base.Ui32(int32(12000)) {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v114 = F_palloc0(m, int32(12))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v114))) = int64(201863463291)
	v120 = F_GetSysCacheHashValue(m, int32(47), v110, int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114)+8)) = v120
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+60))
	v125 = F_lappend(m, v124, v114)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L12
	} else {
		goto L38
	}
L38:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v127)+60)) = v125
	return
L39:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(int32(12000)) <= base.Ui32(v131) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v135 = F_palloc0(m, int32(12))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L12
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if base.Ui32(int32(12000)) <= base.Ui32(v152) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v135))) = int64(201863463291)
	v141 = F_GetSysCacheHashValue(m, int32(47), v131, int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L12
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+8)) = v141
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+60))
	v146 = F_lappend(m, v145, v135)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v148)+60)) = v146
	goto L42
L46:
	;
	v156 = F_palloc0(m, int32(12))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L12
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if base.Ui32(v173) < base.Ui32(int32(12000)) {
		goto L1
	} else {
		goto L52
	}
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v156))) = int64(201863463291)
	v162 = F_GetSysCacheHashValue(m, int32(47), v152, int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+8)) = v162
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+60))
	v167 = F_lappend(m, v166, v156)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+60)) = v167
	goto L48
L52:
	;
	v177 = F_palloc0(m, int32(12))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L12
	} else {
		goto L53
	}
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v177))) = int64(201863463291)
	v183 = F_GetSysCacheHashValue(m, int32(47), v173, int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177)+8)) = v183
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+60))
	v188 = F_lappend(m, v187, v177)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L12
	} else {
		goto L55
	}
L55:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v190)+60)) = v188
	return
L56:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v198 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+56))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v202 = F_lappend_oid(m, v200, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v204)+56)) = v202
	return
L59:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v209 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v241 != 0 {
		goto L1
	} else {
		goto L69
	}
L61:
	;
	v239 = v3
	goto L60
L62:
	;
	goto L63
L63:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	if v212 <= int32(0) {
		v239 = v3
		goto L60
	} else {
		goto L64
	}
L64:
	;
	v216 = int32(0)
	v219 = v3
	goto L65
L65:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v209)+12))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v221+v216<<(uint(int32(2))%32))))
	v229 = int32(*(*int16)(unsafe.Add(mBase, uint32(v206+v225<<(uint(int32(1))%32)))))
	v230 = F_lappend_int(m, v219, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L12
	} else {
		goto L67
	}
L66:
	;
	v239 = v230
	goto L60
L67:
	;
	v233 = v216 + int32(1)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	if v233 < v234 {
		v216 = v233
		v219 = v230
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v239
	goto L1
}
func F_fix_scan_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	if l2 != 0 {
		v20 = F_fix_scan_expr_mutator(m, l1, v8)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v26 = v20
			m.G0 = v8 + int32(16)
			return v26
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
		if v13 != 0 {
			v20 = F_fix_scan_expr_mutator(m, l1, v8)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v26 = v20
				m.G0 = v8 + int32(16)
				return v26
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
			if v15 != 0 {
				v20 = F_fix_scan_expr_mutator(m, l1, v8)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v26 = v20
					m.G0 = v8 + int32(16)
					return v26
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
				if v16 != 0 {
					v20 = F_fix_scan_expr_mutator(m, l1, v8)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v26 = v20
						m.G0 = v8 + int32(16)
						return v26
					}
				} else {
					v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)))
					if v17 != int32(1) {
						v24 = F_fix_scan_expr_walker(m, l1, v8)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							v26 = l1
							m.G0 = v8 + int32(16)
							return v26
						}
					} else {
						v20 = F_fix_scan_expr_mutator(m, l1, v8)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							v26 = v20
							m.G0 = v8 + int32(16)
							return v26
						}
					}
				}
			}
		}
	}
}
func F_get_expr_result_tupdesc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v12 = F_get_expr_result_type(m, l0, v3, v7+int32(12))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(1)
		if base.Ui32(v12-v16) <= base.Ui32(v16) {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			v52 = v20
			m.G0 = v7 + int32(16)
			return v52
		} else {
			if l1 != 0 {
				v52 = v3
				m.G0 = v7 + int32(16)
				return v52
			} else {
				v21 = F_exprType(m, l0)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							if v21 == int32(2249) {
								F_errmsg(m, int32(451351), int32(0))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(498152), int32(574), int32(488696))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							} else {
								v41 = F_format_type_be(m, v21)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v41
									F_errmsg(m, int32(349594), v7)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(498152), int32(570), int32(488696))
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
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
		}
	}
}
func F_get_expr_width(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v5 == int32(6) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if v8 < int32(0) {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			v38 = F_get_typavgwidth(m, v36, v37)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				return v38
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			if v11 <= v8 {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				v38 = F_get_typavgwidth(m, v36, v37)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					return v38
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v13+v8<<(uint(int32(2))%32))))
				if v17 == int32(0) {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v38 = F_get_typavgwidth(m, v36, v37)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						return v38
					}
				} else {
					v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+8)))
					v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+80)))
					if v20 < v21 {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v38 = F_get_typavgwidth(m, v36, v37)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							return v38
						}
					} else {
						v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+82)))
						if v23 < v20 {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v38 = F_get_typavgwidth(m, v36, v37)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								return v38
							}
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)+88))
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v25+(v20-v21)<<(uint(int32(2))%32))))
							if int32(0) < v30 {
								v49 = v30
								return v49
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
								v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
								v38 = F_get_typavgwidth(m, v36, v37)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									return v38
								}
							}
						}
					}
				}
			}
		}
	} else {
		v43 = F_exprType(m, l1)
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			v45 = F_exprTypmod(m, l1)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				v47 = F_get_typavgwidth(m, v43, v45)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v49 = v47
					return v49
				}
			}
		}
	}
}
