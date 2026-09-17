package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AddQual(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
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
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	if l1 == int32(0) {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v6 == int32(6) {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v9 != 0 {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				if v10 == int32(222) {
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v19 = m.ExcPending
						if v19 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_AddQual_0), int32(0))
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_AddQual_1), int32(1175), int32(_a_F_AddQual_2))
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
									return
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
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_AddQual_0), int32(0))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_AddQual_1), int32(1175), int32(_a_F_AddQual_2))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return
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
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
			if v29 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_AddQual_3), int32(0))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_AddQual_1), int32(1187), int32(_a_F_AddQual_2))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v30 = F_copyObjectImpl(m, l1)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
					v34 = F_make_and_qual(m, v33, v30)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
						*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v34
						v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
						if v38 != 0 {
							return
						} else {
							v42 = F_query_or_expression_tree_walker_impl(m, v30, int32(1047), int32(0), int32(3))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)) = uint8(v42)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_ExecInitQual(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
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
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v179 int64
	_ = v179
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l0 == v3 {
		v198 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v198
L2:
	;
	v17 = F_palloc0(m, int32(68))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(380)
	v27 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)) = uint8(v27)
	v29 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v29
	v33 = F_expr_setup_walker(m, l0, v12)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	F_ExecPushExprSetupSteps(m, v17, v12)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v38 = v17 + int32(5)
	v40 = v17 + int32(8)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v41 <= int32(0) {
		v146 = v3
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	if v150 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L8:
	;
	v48 = v3
	v49 = v3
	goto L9
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v49<<(uint(int32(2))%32))))
	F_ExecInitExprRec(m, v57, v17, v40, v38)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L3
	} else {
		goto L11
	}
L10:
	;
	v110 = int32(-1)
	if v104 == int32(0) {
		v146 = v110
		goto L7
	} else {
		goto L24
	}
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	if v60 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v83 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v82 + v83
	v88 = v81 + v82*int32(40)
	v89 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+20)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v88)+12)) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = int32(39)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+28)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v88)+36)) = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v104 = F_lappend_int(m, v48, v101-v83)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L22
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v79
	v81 = v79
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = int32(16)
	v66 = F_palloc(m, int32(640))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L3
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if v68 != v60 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v79 = v66
	goto L13
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v81 = v70
	goto L12
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v60 << (uint(int32(1)) % 32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v77 = F_repalloc(m, v74, v60*int32(80))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v79 = v77
	goto L13
L22:
	;
	v107 = v49 + int32(1)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v107 < v108 {
		v48 = v104
		v49 = v107
		goto L9
	} else {
		goto L23
	}
L23:
	;
	goto L10
L24:
	;
	v113 = int32(0)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v114 <= v113 {
		v146 = v110
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v118 = v113
	goto L26
L26:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127+v118<<(uint(int32(2))%32))))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v126+v131*int32(40))+16)) = v135
	v138 = v118 + int32(1)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v138 < v139 {
		v118 = v138
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v146 = v110
	goto L7
L28:
	;
	goto L27
L29:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v172 + int32(1)
	v178 = v171 + v172*int32(40)
	v179 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v178)+20)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v178)+16)) = v146
	v182 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v178)+12)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v178)+8)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v178)+4)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v182
	*(*int64)(unsafe.Add(mBase, uint32(v178)+28)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v178)+36)) = v182
	v192 = F_jit_compile_expr(m, v17)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L3
	} else {
		goto L39
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v169
	v171 = v169
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = int32(16)
	v156 = F_palloc(m, int32(640))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L3
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if v158 != v150 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v169 = v156
	goto L30
L35:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v171 = v160
	goto L29
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v150 << (uint(int32(1)) % 32)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v167 = F_repalloc(m, v164, v150*int32(80))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v169 = v167
	goto L30
L39:
	;
	if v192 != 0 {
		v198 = v17
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_ExecReadyInterpretedExpr(m, v17)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	v198 = v17
	goto L1
}
func F_cost_qual_eval(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v10
	v16 = v8 + int32(16)
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v46
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v48
	m.G0 = v8 + int32(32)
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v19 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = int32(0)
	goto L4
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v25<<(uint(int32(2))%32))))
	v35 = F_cost_qual_eval_walker(m, v32, v8+int32(8))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	v38 = v25 + int32(1)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v38 < v39 {
		v25 = v38
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
func F_get_qual_for_range(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
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
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v618 int32
	_ = v618
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	v4 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(48)
	m.G0 = v31
	v33 = F_RelationGetPartitionKey(m, l0)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v37 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	m.G0 = v31 + int32(48)
	return v693
L4:
	;
	v341 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+4)))
	v342 = v341 - v334
	v343 = int32(0)
	if v343 < v342 {
		goto L92
	} else {
		goto L93
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L89
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L86
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L83
	}
L8:
	;
	v41 = F_RelationGetPartitionDesc(m, l0, int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	if l2 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v43 <= int32(0) {
		v693 = v4
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v51 = v4
	v52 = v4
	goto L13
L13:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v46+v52<<(uint(int32(2))%32))))
	v80 = F_SearchSysCache1(m, int32(57), v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	if v116 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L15:
	;
	if v80 == int32(0) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v86 = F_SysCacheGetAttrNotNull(m, int32(57), v80, int32(34))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v88 = F_text_to_cstring(m, v86)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v90 = F_stringToNode(m, v88)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v92 != int32(98) {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+5)))
	if v95 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v99 = F_get_qual_for_range(m, l0, v90, int32(1))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L26
	}
L22:
	;
	v116 = v51
	goto L23
L23:
	;
	F_ReleaseCatCache(m, v80)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L31
	}
L24:
	;
	v113 = F_lappend(m, v51, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L30
	}
L25:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v112 = v111
	goto L24
L26:
	;
	if v99 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v103 < int32(2) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v108 = F_makeBoolExpr(m, int32(0), v99, int32(-1))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v112 = v108
	goto L24
L30:
	;
	v116 = v113
	goto L23
L31:
	;
	v120 = v52 + int32(1)
	if v120 != v43 {
		v51 = v116
		v52 = v120
		goto L13
	} else {
		goto L32
	}
L32:
	;
	goto L14
L33:
	;
	v693 = int32(0)
	goto L3
L34:
	;
	goto L35
L35:
	;
	v126 = F_get_range_nulltest(m, v33)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	if int32(2) <= v128 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v138 = F_lappend(m, v126, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L42
	}
L38:
	;
	v133 = F_makeBoolExpr(m, int32(1), v116, int32(-1))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v137 = v136
	goto L37
L41:
	;
	v137 = v133
	goto L37
L42:
	;
	v141 = F_makeBoolExpr(m, int32(0), v138, int32(-1))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v141
	v149 = F_list_make1_impl(m, int32(1), v31+int32(8))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v152 = F_makeBoolExpr(m, int32(2), v149, int32(-1))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v31)+28)) = v152
	v159 = F_list_make1_impl(m, int32(1), v31+int32(4))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v693 = v159
	goto L3
L47:
	;
	v163 = F_get_range_nulltest(m, v33)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	v165 = v4
	goto L49
L49:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	if v166 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v165 = v163
	goto L49
L51:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+12))
	v168 = v167
	goto L53
L52:
	;
	v168 = v4
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v168
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v176 = v165
	v178 = v4
	v190 = v168
	goto L54
L54:
	;
	v200 = int32(0)
	if v171 == v200 {
		v210 = v200
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if v170 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	if v204 <= v178 {
		v210 = int32(0)
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
	v210 = v206 + v178<<(uint(int32(2))%32)
	goto L56
L59:
	;
	v213 = int32(0)
	v333 = v165
	v334 = v213
	v339 = v168
	v340 = v213
	goto L4
L60:
	;
	goto L61
L61:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	if v215 <= v178 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v333 = v176
	v334 = v178
	v339 = v190
	v340 = int32(0)
	goto L4
L63:
	;
	goto L64
L64:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
	v221 = v218 + v178<<(uint(int32(2))%32)
	v222 = int32(0)
	if base.B2i32(v210 == v222)|base.B2i32(v218 == v222) != 0 {
		v333 = v176
		v334 = v178
		v339 = v190
		v340 = v221
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	F_get_range_key_properties(m, v33, v178, v228, v229, v31+int32(44), v31+int32(40), v31+int32(36), v31+int32(32))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v31)+36))
	if v240 == int32(0) {
		v333 = v176
		v334 = v178
		v339 = v227
		v340 = v221
		goto L4
	} else {
		goto L67
	}
L67:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	if v243 == int32(0) {
		v333 = v176
		v334 = v178
		v339 = v227
		v340 = v221
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v246 = F_CreateExecutorState(m)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v248 = int32(_a_F_get_qual_for_range_0)
	v249 = *(*int32)(unsafe.Add(mBase, _c_F_get_qual_for_range[0]))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v246)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_get_qual_for_range[0])) = v251
	v254 = F_make_partition_op_expr(m, v33, v178, int32(3), v240, v243)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_fix_opfuncids(m, v254)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v259 = F_ExecInitExpr(m, v254, int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v246)+152))
	if v261 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v264 = F_MakePerTupleExprContext(m, v246)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	v266 = v261
	goto L75
L75:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v266)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_get_qual_for_range[0])) = v268
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v259)+20))
	v273 = m.T0[v272].(func(*base.Module, int32, int32, int32) int32)(m, v259, v266, v31+int32(23))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L77
	}
L76:
	;
	v266 = v264
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_qual_for_range[0])) = v249
	F_FreeExecutorState(m, v246)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	if v273 == int32(0) {
		v333 = v176
		v334 = v178
		v339 = v227
		v340 = v221
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v281 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+4)))
	if v178 == v281-int32(1) {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
	v287 = F_make_partition_op_expr(m, v33, v178, int32(3), v286, v240)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v289 = F_lappend(m, v176, v287)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v176 = v289
	v178 = v178 + int32(1)
	v190 = v227
	goto L54
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v79
	F_errmsg_internal(m, int32(_a_F_get_qual_for_range_1), v31)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_get_qual_for_range_2), int32(_a_F_get_qual_for_range_3), int32(_a_F_get_qual_for_range_4))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_errmsg_internal(m, int32(_a_F_get_qual_for_range_5), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_get_qual_for_range_2), int32(_a_F_get_qual_for_range_6), int32(_a_F_get_qual_for_range_4))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	F_errmsg_internal(m, int32(_a_F_get_qual_for_range_7), int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_get_qual_for_range_2), int32(_a_F_get_qual_for_range_8), int32(_a_F_get_qual_for_range_4))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	v346 = v342
	goto L94
L93:
	;
	v346 = v343
	goto L94
L94:
	;
	v347 = int32(1)
	v359 = int32(0)
	v367 = v4
	v370 = v4
	v373 = v347
	v374 = v347
	goto L95
L95:
	;
	if v359 == v346 {
		v639 = v367
		v642 = v370
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if v639 != 0 {
		goto L176
	} else {
		goto L177
	}
L97:
	;
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v339
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v210 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v340 != 0 {
		goto L105
	} else {
		goto L106
	}
L100:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+12))
	v389 = (v210 - v381) >> (uint(int32(2)) % 32)
	goto L99
L101:
	;
	goto L102
L102:
	;
	v385 = int32(0)
	if v380 == v385 {
		v389 = v385
		goto L99
	} else {
		goto L103
	}
L103:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v380)+4))
	v389 = v388
	goto L99
L104:
	;
	v400 = int32(0)
	v407 = v334
	v409 = v399
	v414 = v400
	v415 = v400
	v424 = v389
	goto L109
L105:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)+12))
	v399 = (v340 - v391) >> (uint(int32(2)) % 32)
	goto L104
L106:
	;
	goto L107
L107:
	;
	v395 = int32(0)
	if v390 == v395 {
		v399 = v395
		goto L104
	} else {
		goto L108
	}
L108:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	v399 = v398
	goto L104
L109:
	;
	v430 = int32(0)
	if v380 == v430 {
		v440 = v430
		goto L111
	} else {
		goto L112
	}
L111:
	;
	if v390 == int32(0) {
		v609 = v367
		v610 = v370
		v611 = v373
		v612 = v374
		goto L115
	} else {
		goto L116
	}
L112:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v380)+4))
	if v434 <= v424 {
		v440 = int32(0)
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v380)+12))
	v440 = v436 + v424<<(uint(int32(2))%32)
	goto L111
L114:
	;
	v618 = int32(1)
	v407 = v536
	v409 = v409 + v618
	v414 = v513
	v415 = v534
	v424 = v424 + v618
	goto L109
L115:
	;
	if v611|v612 != 0 {
		v359 = v359 + int32(1)
		v367 = v609
		v370 = v610
		v373 = v611
		v374 = v612
		goto L95
	} else {
		goto L175
	}
L116:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	if base.B2i32(v440 == int32(0))|base.B2i32(v445 <= v409) != 0 {
		v566 = v414
		v567 = v415
		v571 = v373
		v572 = v374
		goto L117
	} else {
		goto L118
	}
L117:
	;
	if v566 != 0 {
		goto L159
	} else {
		goto L160
	}
L118:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v390)+12))
	if v448 == int32(0) {
		v566 = v414
		v567 = v415
		v571 = v373
		v572 = v374
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v451 = int32(0)
	v454 = v440 + int32(4)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)+12))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v455)+4))
	if base.Ui32(v454) < base.Ui32(v456+v457<<(uint(int32(2))%32)) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	v463 = v462
	goto L122
L121:
	;
	v463 = v451
	goto L122
L122:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	v465 = int32(2)
	v467 = v448 + v409<<(uint(v465)%32)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	v470 = v467 + int32(4)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)+12))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v471)+4))
	if base.Ui32(v470) < base.Ui32(v472+v473<<(uint(v465)%32)) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v470)))
	v479 = v478
	goto L125
L124:
	;
	v479 = v451
	goto L125
L125:
	;
	F_get_range_key_properties(m, v33, v407, v464, v468, v31+int32(44), v31+int32(40), v31+int32(36), v31+int32(32))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v31)+36))
	if base.B2i32(v490 != int32(0))&v374 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	if v407-v334 < v359 {
		v506 = int32(3)
		goto L130
	} else {
		goto L131
	}
L128:
	;
	v513 = v414
	goto L129
L129:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	if base.B2i32(v514 != int32(0))&v373 != 0 {
		goto L139
	} else {
		goto L140
	}
L130:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
	v508 = F_make_partition_op_expr(m, v33, v407, v506, v507, v490)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L137
	}
L131:
	;
	v497 = int32(4)
	v498 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+4)))
	if v407 == v498-int32(1) {
		v506 = v497
		goto L130
	} else {
		goto L132
	}
L132:
	;
	if v463 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v463)+4))
	if v502 == int32(-1) {
		v506 = v497
		goto L130
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v506 = int32(5)
	goto L130
L136:
	;
	goto L135
L137:
	;
	v510 = F_lappend(m, v414, v508)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v513 = v510
	goto L129
L139:
	;
	if v407-v334 < v359 {
		v528 = int32(3)
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v534 = v415
	goto L141
L141:
	;
	v536 = v407 + int32(1)
	if v536-v334 <= v359 {
		goto L114
	} else {
		goto L149
	}
L142:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
	v530 = F_make_partition_op_expr(m, v33, v407, v528, v529, v514)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L147
	}
L143:
	;
	if v479 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v528 = int32(1)
	goto L142
L145:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v479)+4))
	if v523 != int32(1) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v528 = int32(2)
	goto L142
L147:
	;
	v532 = F_lappend(m, v415, v530)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v534 = v532
	goto L141
L149:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v31)+36))
	v540 = int32(0)
	if base.B2i32(v539 == v540)|base.B2i32(v463 == v540) == v540 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	v553 = int32(0)
	if base.B2i32(v552 == v553)|base.B2i32(v479 == v553) == v553 {
		goto L155
	} else {
		goto L156
	}
L151:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v463)+4))
	if v547 == int32(0) {
		v551 = v374
		goto L150
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v551 = int32(0)
	goto L150
L154:
	;
	goto L153
L155:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v479)+4))
	if v560 == int32(0) {
		v566 = v513
		v567 = v534
		v571 = v373
		v572 = v551
		goto L117
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v566 = v513
	v567 = v534
	v571 = int32(0)
	v572 = v551
	goto L117
L158:
	;
	goto L157
L159:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	if int32(2) <= v575 {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	v587 = v367
	goto L161
L161:
	;
	if v567 == int32(0) {
		v609 = v587
		v610 = v370
		v611 = v571
		v612 = v572
		goto L115
	} else {
		goto L168
	}
L162:
	;
	v585 = F_lappend(m, v367, v584)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L167
	}
L163:
	;
	v580 = F_makeBoolExpr(m, int32(0), v566, int32(-1))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v566)+12))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v582)))
	v584 = v583
	goto L162
L166:
	;
	v584 = v580
	goto L162
L167:
	;
	v587 = v585
	goto L161
L168:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	if int32(2) <= v590 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v600 = F_lappend(m, v370, v599)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L174
	}
L170:
	;
	v595 = F_makeBoolExpr(m, int32(0), v567, int32(-1))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v567)+12))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v597)))
	v599 = v598
	goto L169
L173:
	;
	v599 = v595
	goto L169
L174:
	;
	v609 = v587
	v610 = v600
	v611 = v571
	v612 = v572
	goto L115
L175:
	;
	v639 = v609
	v642 = v610
	goto L97
L176:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v639)+4))
	if int32(2) <= v650 {
		goto L180
	} else {
		goto L181
	}
L177:
	;
	v662 = v333
	goto L178
L178:
	;
	if v642 != 0 {
		goto L185
	} else {
		goto L186
	}
L179:
	;
	v660 = F_lappend(m, v333, v659)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L184
	}
L180:
	;
	v655 = F_makeBoolExpr(m, int32(1), v639, int32(-1))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v639)+12))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v657)))
	v659 = v658
	goto L179
L183:
	;
	v659 = v655
	goto L179
L184:
	;
	v662 = v660
	goto L178
L185:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v642)+4))
	if int32(2) <= v663 {
		goto L189
	} else {
		goto L190
	}
L186:
	;
	v675 = v662
	goto L187
L187:
	;
	if v675 != 0 {
		v693 = v675
		goto L3
	} else {
		goto L194
	}
L188:
	;
	v673 = F_lappend(m, v662, v672)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L193
	}
L189:
	;
	v668 = F_makeBoolExpr(m, int32(1), v642, int32(-1))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v642)+12))
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v670)))
	v672 = v671
	goto L188
L192:
	;
	v672 = v668
	goto L188
L193:
	;
	v675 = v673
	goto L187
L194:
	;
	if l2 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v676 = F_get_range_nulltest(m, v33)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v680 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L199
	}
L198:
	;
	v693 = v676
	goto L3
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v680
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v680
	v687 = F_list_make1_impl(m, int32(1), v31+int32(12))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v693 = v687
	goto L3
}
