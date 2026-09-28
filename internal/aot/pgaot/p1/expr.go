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
	v5 = F_query_or_expression_tree_walker_impl(m, l0, int32(1125), int32(0), int32(3))
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v69 int32
	_ = v69
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
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = F_CreateExecutorState(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(_a_F_evaluate_expr_0)
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0]))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
		*(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0])) = v22
		F_fix_opfuncids(m, l0)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v27 = F_ExecInitExpr(m, l0, int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v15)+152))
				if v29 == int32(0) {
					v32 = F_MakePerTupleExprContext(m, v15)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = v32
						v35 = int32(_a_F_evaluate_expr_0)
						v36 = *(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0]))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
						*(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0])) = v38
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
						v43 = m.T0[v42].(func(*base.Module, int32, int32, int32) int64)(m, v27, v34, v13+int32(15))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0])) = v36
							F_get_typlenbyval(m, l1, v13+int32(12), v13+int32(11))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0])) = v20
								v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
								if v55 != 0 {
									v67 = v43
									F_FreeExecutorState(m, v15)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										v70 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+12)))
										v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
										v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
										v73 = F_makeConst(m, l1, l2, l3, v70, v67, v71, v72)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											m.G0 = v13 + int32(16)
											return v73
										}
									}
								} else {
									v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+12)))
									if v56 == int32(-1) {
										v60 = F_pg_detoast_datum_copy(m, base.I32_wrap_i64(v43))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											v67 = base.I64_extend_i32_u(v60)
											F_FreeExecutorState(m, v15)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												v70 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+12)))
												v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
												v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
												v73 = F_makeConst(m, l1, l2, l3, v70, v67, v71, v72)
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return int32(0)
												} else {
													m.G0 = v13 + int32(16)
													return v73
												}
											}
										}
									} else {
										v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
										v64 = F_datumCopy(m, v43, v63, v56)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v67 = v64
											F_FreeExecutorState(m, v15)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												v70 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+12)))
												v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
												v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
												v73 = F_makeConst(m, l1, l2, l3, v70, v67, v71, v72)
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return int32(0)
												} else {
													m.G0 = v13 + int32(16)
													return v73
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v34 = v29
					v35 = int32(_a_F_evaluate_expr_0)
					v36 = *(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0]))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
					*(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0])) = v38
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
					v43 = m.T0[v42].(func(*base.Module, int32, int32, int32) int64)(m, v27, v34, v13+int32(15))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0])) = v36
						F_get_typlenbyval(m, l1, v13+int32(12), v13+int32(11))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_evaluate_expr[0])) = v20
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
							if v55 != 0 {
								v67 = v43
								F_FreeExecutorState(m, v15)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									v70 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+12)))
									v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
									v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
									v73 = F_makeConst(m, l1, l2, l3, v70, v67, v71, v72)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										m.G0 = v13 + int32(16)
										return v73
									}
								}
							} else {
								v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+12)))
								if v56 == int32(-1) {
									v60 = F_pg_detoast_datum_copy(m, base.I32_wrap_i64(v43))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v67 = base.I64_extend_i32_u(v60)
										F_FreeExecutorState(m, v15)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											v70 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+12)))
											v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
											v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
											v73 = F_makeConst(m, l1, l2, l3, v70, v67, v71, v72)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												m.G0 = v13 + int32(16)
												return v73
											}
										}
									}
								} else {
									v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
									v64 = F_datumCopy(m, v43, v63, v56)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										v67 = v64
										F_FreeExecutorState(m, v15)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											v70 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+12)))
											v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
											v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
											v73 = F_makeConst(m, l1, l2, l3, v70, v67, v71, v72)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												m.G0 = v13 + int32(16)
												return v73
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
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
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
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
	v139 = F_palloc0(m, int32(12))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L15
	} else {
		goto L48
	}
L2:
	;
	return
L3:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	if v94 == int32(0) {
		goto L2
	} else {
		goto L39
	}
L4:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(v80 != int32(2205))&base.B2i32(v80 != int32(26)) != 0 {
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
	v136 = v9
	goto L1
L13:
	;
	v136 = v12
	goto L1
L14:
	;
	v136 = v15
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
	v136 = v20
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
	v136 = v25
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
	v136 = v30
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
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if base.Ui32(int32(_a_F_fix_expr_common_0)) <= base.Ui32(v56) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = int64(201863463295)
	v46 = F_GetSysCacheHashValue(m, int32(47), base.I64_extend_i32_u(v35), int64(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v46
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+68))
	v51 = F_lappend(m, v50, v39)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+68)) = v51
	goto L25
L29:
	;
	v60 = F_palloc0(m, int32(12))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L15
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if base.Ui32(v77) < base.Ui32(int32(_a_F_fix_expr_common_0)) {
		goto L2
	} else {
		goto L35
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v60))) = int64(201863463295)
	v67 = F_GetSysCacheHashValue(m, int32(47), base.I64_extend_i32_u(v56), int64(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v67
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+68))
	v72 = F_lappend(m, v71, v60)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v74)+68)) = v72
	goto L31
L35:
	;
	v136 = v77
	goto L1
L36:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v86 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+64))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v90 = F_lappend_oid(m, v88, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L15
	} else {
		goto L38
	}
L38:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+64)) = v90
	return
L39:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v97 == int32(0) {
		v127 = v3
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v129 != 0 {
		goto L2
	} else {
		goto L47
	}
L41:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if v100 <= int32(0) {
		v127 = v3
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v104 = int32(0)
	v107 = v3
	goto L43
L43:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109+v104<<(uint(int32(2))%32))))
	v117 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94+v113<<(uint(int32(1))%32)))))
	v118 = F_lappend_int(m, v107, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L15
	} else {
		goto L45
	}
L44:
	;
	v127 = v118
	goto L40
L45:
	;
	v121 = v104 + int32(1)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if v121 < v122 {
		v104 = v121
		v107 = v118
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v127
	goto L2
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v139))) = int64(201863463295)
	v146 = F_GetSysCacheHashValue(m, int32(47), base.I64_extend_i32_u(v136), int64(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L15
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139)+8)) = v146
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+68))
	v151 = F_lappend(m, v150, v139)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L15
	} else {
		goto L50
	}
L50:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+68)) = v151
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
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
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
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
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
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
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
					v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
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
									F_errfinish(m, int32(_a_F_get_expr_result_tupdesc_1), int32(576), int32(_a_F_get_expr_result_tupdesc_2))
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
										F_errfinish(m, int32(_a_F_get_expr_result_tupdesc_1), int32(572), int32(_a_F_get_expr_result_tupdesc_2))
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
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
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
					v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+88)))
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
						v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+90)))
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
							v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)+96))
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
