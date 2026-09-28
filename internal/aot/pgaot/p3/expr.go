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
	var v21 int64
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
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
					if v28 != 0 {
						F_SPI_freetuptable(m, v28)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
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
						*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
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
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
				if v28 != 0 {
					F_SPI_freetuptable(m, v28)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
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
					*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
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
	var v50 int32
	_ = v50
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
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
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
	return v249
L2:
	;
	v228 = F_exprType(m, l0)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L8
	} else {
		goto L68
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
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v200 != int32(2249) {
		goto L2
	} else {
		goto L56
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
	v249 = v24
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
	v249 = v32
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
	v50 = v4
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
	if v61 <= v50 {
		v67 = v58
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v67 = v63 + v50<<(uint(int32(2))%32)
	goto L19
L22:
	;
	v172 = base.I32_extend16_s(v55)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v75+v50<<(uint(int32(2))%32))))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v179 = F_exprType(m, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L8
	} else {
		goto L51
	}
L23:
	;
	v77 = int32(0)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v77 < v86 {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if base.B2i32(v67 == int32(0))|base.B2i32(v72 <= v50) != 0 {
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
	if l1 != 0 {
		goto L46
	} else {
		goto L47
	}
L28:
	;
	v90 = v41 + int32(28)
	v97 = v77
	v98 = v86
	v100 = v77
	goto L32
L29:
	;
	v154 = v77
	v161 = v86
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+20)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v154
	goto L27
L31:
	;
	v154 = v148
	v161 = v127
	goto L30
L32:
	;
	v106 = v90 + v86<<(uint(int32(3))%32) + v97*int32(100)
	v109 = v90 + v97<<(uint(int32(3))%32)
	if v86 != v98 {
		v127 = v98
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v148 = v86
	goto L31
L34:
	;
	v128 = int32(*(*int16)(unsafe.Add(mBase, uint32(v109)+2)))
	if v128 <= int32(0) {
		v148 = v97
		goto L31
	} else {
		goto L42
	}
L35:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+7)))
	if v111 != int32(118) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v127 = v97
	goto L34
L37:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+4)))
	if v114 != int32(1) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+6)))
	if v117&int32(6) != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v120 = int32(*(*int16)(unsafe.Add(mBase, uint32(v109)+2)))
	if v120 <= int32(0) {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+90)))
	if v123 != int32(118) {
		v127 = v86
		goto L34
	} else {
		goto L41
	}
L41:
	;
	goto L36
L42:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+90)))
	if v131 == int32(118) {
		v148 = v97
		goto L31
	} else {
		goto L43
	}
L43:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+5)))
	v140 = (v100 + v134 - int32(1)) & (int32(0) - v134)
	if int32(_a_F_get_expr_result_type_0) < v140 {
		v148 = v97
		goto L31
	} else {
		goto L44
	}
L44:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v109))) = uint16(v140)
	v146 = v97 + int32(1)
	if v146 != v86 {
		v97 = v146
		v98 = v127
		v100 = v140 + v128
		goto L32
	} else {
		goto L45
	}
L45:
	;
	goto L33
L46:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v164
	goto L48
L47:
	;
	goto L48
L48:
	;
	v166 = int32(1)
	if l2 == int32(0) {
		v249 = v166
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v169 = F_BlessTupleDesc(m, v41)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v169
	v249 = v166
	goto L1
L51:
	;
	v181 = F_exprTypmod(m, v178)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L8
	} else {
		goto L52
	}
L52:
	;
	F_TupleDescInitEntry(m, v41, v172, v177, v179, v181, int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	v186 = F_exprCollation(m, v178)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L8
	} else {
		goto L54
	}
L54:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	*(*int32)(unsafe.Add(mBase, uint32(v41+v188<<(uint(int32(3))%32)+v172*int32(100))+24)) = v186
	goto L55
L55:
	;
	v196 = int32(1)
	v50 = v50 + v196
	v55 = v55 + v196
	goto L17
L56:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v203 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v205 = F_pg_detoast_datum(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L8
	} else {
		goto L58
	}
L58:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v205)+8))
	if l1 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v208
	goto L61
L60:
	;
	goto L61
L61:
	;
	v212 = int32(0)
	if base.B2i32(v208 == int32(2249))&base.B2i32(v207 < v212) == v212 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v217 = int32(1)
	if l2 == int32(0) {
		v249 = v217
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v223 = int32(3)
	if l2 == int32(0) {
		v249 = v223
		goto L1
	} else {
		goto L67
	}
L65:
	;
	v220 = F_lookup_rowtype_tupdesc_copy(m, v208, v207)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L8
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v220
	v249 = v217
	goto L1
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v249 = v223
	goto L1
L68:
	;
	if l1 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v228
	goto L71
L70:
	;
	goto L71
L71:
	;
	if l2 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	goto L74
L73:
	;
	goto L74
L74:
	;
	v237 = F_get_type_func_class(m, v228, v15+int32(12))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L8
	} else {
		goto L75
	}
L75:
	;
	v239 = int32(1)
	if base.B2i32(l2 == int32(0))|base.B2i32(base.Ui32(v239) < base.Ui32(v237-v239)) != 0 {
		v249 = v237
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v246 = F_lookup_rowtype_tupdesc_copy(m, v244, int32(-1))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v246
	v249 = v237
	goto L1
}
