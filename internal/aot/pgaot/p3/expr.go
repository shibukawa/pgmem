package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ParseExprKindName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	if base.Ui32(l0) <= base.Ui32(int32(44)) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_ParseExprKindName[0])))
		v8 = v6
	} else {
		v8 = int32(_a_F_ParseExprKindName_0)
	}
	return v8
}
func F_exec_assign_expr(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v9 == int32(0) {
		F_exec_prepare_plan(m, l0, l2, int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v21 = F_exec_eval_expr(m, l0, l2, v7+int32(15), v7+int32(8), v7+int32(4))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
				F_exec_assign_value(m, l0, l1, v21, v23, v24, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					if v28 != 0 {
						F_SPI_freetuptable(m, v28)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
							if v33 != 0 {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
								F_MemoryContextReset(m, v34)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						if v33 != 0 {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
							F_MemoryContextReset(m, v34)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			}
		}
	} else {
		v21 = F_exec_eval_expr(m, l0, l2, v7+int32(15), v7+int32(8), v7+int32(4))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			F_exec_assign_value(m, l0, l1, v21, v23, v24, v25)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				if v28 != 0 {
					F_SPI_freetuptable(m, v28)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						if v33 != 0 {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
							F_MemoryContextReset(m, v34)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
					if v33 != 0 {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
						F_MemoryContextReset(m, v34)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_get_expr_result_type(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if l0 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v162
L2:
	;
	v141 = F_exprType(m, l0)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L8
	} else {
		goto L49
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v19 - int32(7) {
	case 0:
		goto L4
	default:
		goto L2
	case 8:
		goto L7
	case 10:
		goto L6
	case 29:
		goto L5
	}
L4:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v113 != int32(2249) {
		goto L2
	} else {
		goto L37
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v34 != int32(2249) {
		goto L2
	} else {
		goto L12
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = F_get_opcode(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L10
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = F_internal_get_result_type(m, v22, l0, int32(0), l1, l2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v162 = v24
	goto L1
L10:
	;
	v32 = F_internal_get_result_type(m, v29, l0, int32(0), l1, l2)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v162 = v32
	goto L1
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v40 = v38
	goto L15
L14:
	;
	v40 = int32(0)
	goto L15
L15:
	;
	v41 = F_CreateTemplateTupleDesc(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v49 = v4
	v55 = int32(1)
	goto L17
L17:
	;
	v58 = int32(0)
	if v44 == v58 {
		v67 = v58
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v43 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v61 <= v49 {
		v67 = v58
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v67 = v63 + v49<<(uint(int32(2))%32)
	goto L19
L22:
	;
	v85 = base.I32_extend16_s(v55)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v75+v49<<(uint(int32(2))%32))))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v92 = F_exprType(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L32
	}
L23:
	;
	if l1 != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if base.B2i32(v67 == int32(0))|base.B2i32(v72 <= v49) != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	if v75 != 0 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v77
	goto L29
L28:
	;
	goto L29
L29:
	;
	v79 = int32(1)
	if l2 == int32(0) {
		v162 = v79
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v82 = F_BlessTupleDesc(m, v41)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v82
	v162 = v79
	goto L1
L32:
	;
	v94 = F_exprTypmod(m, v91)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	F_TupleDescInitEntry(m, v41, v85, v90, v92, v94, int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	v99 = F_exprCollation(m, v91)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	*(*int32)(unsafe.Add(mBase, uint32(v41+v101<<(uint(int32(4))%32)+v85*int32(100))+16)) = v99
	goto L36
L36:
	;
	v109 = int32(1)
	v49 = v49 + v109
	v55 = v55 + v109
	goto L17
L37:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v116 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v118 = F_pg_detoast_datum(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	if l1 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v121
	goto L42
L41:
	;
	goto L42
L42:
	;
	v125 = int32(0)
	if base.B2i32(v121 == int32(2249))&base.B2i32(v120 < v125) == v125 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v130 = int32(1)
	if l2 == int32(0) {
		v162 = v130
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v136 = int32(3)
	if l2 == int32(0) {
		v162 = v136
		goto L1
	} else {
		goto L48
	}
L46:
	;
	v133 = F_lookup_rowtype_tupdesc_copy(m, v121, v120)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v133
	v162 = v130
	goto L1
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v162 = v136
	goto L1
L49:
	;
	if l1 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v141
	goto L52
L51:
	;
	goto L52
L52:
	;
	if l2 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	goto L55
L54:
	;
	goto L55
L55:
	;
	v150 = F_get_type_func_class(m, v141, v15+int32(12))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	v152 = int32(1)
	if base.B2i32(l2 == int32(0))|base.B2i32(base.Ui32(v152) < base.Ui32(v150-v152)) != 0 {
		v162 = v150
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v159 = F_lookup_rowtype_tupdesc_copy(m, v157, int32(-1))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L8
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v159
	v162 = v150
	goto L1
}
