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
	v4 = int32(_a_F_ExecPrepareExpr_0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_ExecPrepareExpr[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecPrepareExpr[0])) = v7
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
			*(*int32)(unsafe.Add(mBase, _c_F_ExecPrepareExpr[0])) = v5
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v54 int32
	_ = v54
	v3 = int32(0)
	v8 = int32(_a_F_ExecPrepareExprList_0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_ExecPrepareExprList[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecPrepareExprList[0])) = v11
	if l0 == v3 {
		v54 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecPrepareExprList[0])) = v9
	return v54
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 <= int32(0) {
		v54 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = v3
	v22 = v3
	goto L4
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v21<<(uint(int32(2))%32))))
	v30 = int32(_a_F_ExecPrepareExprList_0)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_ExecPrepareExprList[0]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecPrepareExprList[0])) = v33
	v35 = F_expression_planner(m, v29)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v54 = v44
	goto L1
L6:
	;
	return int32(0)
L7:
	;
	v40 = F_ExecInitExpr(m, v35, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecPrepareExprList[0])) = v31
	v44 = F_lappend(m, v22, v40)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v47 = v21 + int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v47 < v48 {
		v21 = v47
		v22 = v44
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L5
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
		v18 = int32(_a_F_evaluate_expr_0)
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0]))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+100))
		*(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0])) = v21
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
						v34 = int32(_a_F_evaluate_expr_0)
						v35 = *(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0]))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
						*(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0])) = v37
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
						v42 = m.T0[v41].(func(*base.Module, int32, int32, int32) int32)(m, v26, v33, v12+int32(15))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0])) = v35
							F_get_typlenbyval(m, l1, v12+int32(12), v12+int32(11))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0])) = v19
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
					v34 = int32(_a_F_evaluate_expr_0)
					v35 = *(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0]))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
					*(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0])) = v37
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
					v42 = m.T0[v41].(func(*base.Module, int32, int32, int32) int32)(m, v26, v33, v12+int32(15))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0])) = v35
						F_get_typlenbyval(m, l1, v12+int32(12), v12+int32(11))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0])) = v19
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v6 - int32(7) {
	case 0:
		goto L4
	default:
		goto L2
	case 2:
		goto L11
	case 3:
		goto L3
	case 4:
		goto L10
	case 8:
		goto L9
	case 10:
		goto L8
	case 11:
		goto L7
	case 12:
		goto L6
	case 13:
		goto L5
	}
L1:
	;
	v137 = F_palloc0(m, int32(12))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L15
	} else {
		goto L48
	}
L2:
	;
	return
L3:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v92 == int32(0) {
		goto L2
	} else {
		goto L39
	}
L4:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(v78 != int32(2205))&base.B2i32(v78 != int32(26)) != 0 {
		goto L2
	} else {
		goto L36
	}
L5:
	;
	F_set_opfuncid(m, l1)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L15
	} else {
		goto L22
	}
L6:
	;
	F_set_opfuncid(m, l1)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L15
	} else {
		goto L20
	}
L7:
	;
	F_set_opfuncid(m, l1)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L15
	} else {
		goto L18
	}
L8:
	;
	F_set_opfuncid(m, l1)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v15) < base.Ui32(int32(_a_F_fix_expr_common_0)) {
		goto L2
	} else {
		goto L14
	}
L10:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v12) < base.Ui32(int32(_a_F_fix_expr_common_0)) {
		goto L2
	} else {
		goto L13
	}
L11:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v9) < base.Ui32(int32(_a_F_fix_expr_common_0)) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v134 = v9
	goto L1
L13:
	;
	v134 = v12
	goto L1
L14:
	;
	v134 = v15
	goto L1
L15:
	;
	return
L16:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v20) < base.Ui32(int32(_a_F_fix_expr_common_0)) {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v134 = v20
	goto L1
L18:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v25) < base.Ui32(int32(_a_F_fix_expr_common_0)) {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v134 = v25
	goto L1
L20:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v30) < base.Ui32(int32(_a_F_fix_expr_common_0)) {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v134 = v30
	goto L1
L22:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(int32(_a_F_fix_expr_common_0)) <= base.Ui32(v35) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v39 = F_palloc0(m, int32(12))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L15
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if base.Ui32(int32(_a_F_fix_expr_common_0)) <= base.Ui32(v55) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = int64(201863463291)
	v45 = F_GetSysCacheHashValue(m, int32(47), v35, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v45
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+60))
	v50 = F_lappend(m, v49, v39)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+60)) = v50
	goto L25
L29:
	;
	v59 = F_palloc0(m, int32(12))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L15
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if base.Ui32(v75) < base.Ui32(int32(_a_F_fix_expr_common_0)) {
		goto L2
	} else {
		goto L35
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v59))) = int64(201863463291)
	v65 = F_GetSysCacheHashValue(m, int32(47), v55, int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v65
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+60))
	v70 = F_lappend(m, v69, v59)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+60)) = v70
	goto L31
L35:
	;
	v134 = v75
	goto L1
L36:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v84 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+56))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v88 = F_lappend_oid(m, v86, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L15
	} else {
		goto L38
	}
L38:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+56)) = v88
	return
L39:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v95 == int32(0) {
		v125 = v3
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v127 != 0 {
		goto L2
	} else {
		goto L47
	}
L41:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v98 <= int32(0) {
		v125 = v3
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v102 = int32(0)
	v105 = v3
	goto L43
L43:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+v102<<(uint(int32(2))%32))))
	v115 = int32(*(*int16)(unsafe.Add(mBase, uint32(v92+v111<<(uint(int32(1))%32)))))
	v116 = F_lappend_int(m, v105, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L15
	} else {
		goto L45
	}
L44:
	;
	v125 = v116
	goto L40
L45:
	;
	v119 = v102 + int32(1)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v119 < v120 {
		v102 = v119
		v105 = v116
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v125
	goto L2
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v137))) = int64(201863463291)
	v143 = F_GetSysCacheHashValue(m, int32(47), v134, int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L15
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137)+8)) = v143
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+60))
	v148 = F_lappend(m, v147, v137)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L15
	} else {
		goto L50
	}
L50:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v150)+60)) = v148
	return
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
								F_errmsg(m, int32(_a_F_get_expr_result_tupdesc_0), int32(0))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_get_expr_result_tupdesc_1), int32(574), int32(_a_F_get_expr_result_tupdesc_2))
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
									F_errmsg(m, int32(_a_F_get_expr_result_tupdesc_3), v7)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_get_expr_result_tupdesc_1), int32(570), int32(_a_F_get_expr_result_tupdesc_2))
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
