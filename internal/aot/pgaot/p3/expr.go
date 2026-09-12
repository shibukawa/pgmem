package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ParseExprKindName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	if base.Ui32(l0) <= base.Ui32(int32(44)) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[443])))
		v11 = v10
	} else {
		v11 = int32(416162)
	}
	return v11
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
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
	return v159
L2:
	;
	v139 = F_exprType(m, l0)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L8
	} else {
		goto L50
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
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v111 != int32(2249) {
		goto L2
	} else {
		goto L38
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
	v159 = v24
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
	v159 = v32
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
	v39 = v38
	goto L15
L14:
	;
	v39 = v4
	goto L15
L15:
	;
	v40 = F_CreateTemplateTupleDesc(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v48 = v4
	v54 = int32(1)
	goto L17
L17:
	;
	v57 = int32(0)
	if v43 == v57 {
		v66 = v57
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v42 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v60 <= v48 {
		v66 = v57
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v66 = v62 + v48<<(uint(int32(2))%32)
	goto L19
L22:
	;
	v86 = base.I32_extend16_s(v54)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v90 = F_exprType(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L8
	} else {
		goto L33
	}
L23:
	;
	if l1 != 0 {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v69 <= v48 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	if v66 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v76 = v73 + v48<<(uint(int32(2))%32)
	if v76 != 0 {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v78
	goto L30
L29:
	;
	goto L30
L30:
	;
	v80 = int32(1)
	if l2 == int32(0) {
		v159 = v80
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v83 = F_BlessTupleDesc(m, v40)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v83
	v159 = v80
	goto L1
L33:
	;
	v92 = F_exprTypmod(m, v89)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	F_TupleDescInitEntry(m, v40, v86, v88, v90, v92, int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	v97 = F_exprCollation(m, v89)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	*(*int32)(unsafe.Add(mBase, uint32(v40+v99<<(uint(int32(4))%32)+v86*int32(100))+16)) = v97
	goto L37
L37:
	;
	v107 = int32(1)
	v48 = v48 + v107
	v54 = v54 + v107
	goto L17
L38:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v114 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v116 = F_pg_detoast_datum(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	if l1 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v119
	goto L43
L42:
	;
	goto L43
L43:
	;
	v123 = int32(0)
	if base.B2i32(v119 == int32(2249))&base.B2i32(v118 < v123) == v123 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v128 = int32(1)
	if l2 == int32(0) {
		v159 = v128
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v134 = int32(3)
	if l2 == int32(0) {
		v159 = v134
		goto L1
	} else {
		goto L49
	}
L47:
	;
	v131 = F_lookup_rowtype_tupdesc_copy(m, v119, v118)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v131
	v159 = v128
	goto L1
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v159 = v134
	goto L1
L50:
	;
	if l1 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v139
	goto L53
L52:
	;
	goto L53
L53:
	;
	if l2 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	goto L56
L55:
	;
	goto L56
L56:
	;
	v146 = F_get_type_func_class(m, v139, v15+int32(12))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L8
	} else {
		goto L57
	}
L57:
	;
	if l2 == int32(0) {
		v159 = v146
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v150 = int32(1)
	if base.Ui32(v150) < base.Ui32(v146-v150) {
		v159 = v146
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v156 = F_lookup_rowtype_tupdesc_copy(m, v154, int32(-1))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L8
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v156
	v159 = v146
	goto L1
}
