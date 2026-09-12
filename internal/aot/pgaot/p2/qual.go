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
							F_errmsg(m, int32(443313), int32(0))
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return
							} else {
								F_errfinish(m, int32(492736), int32(1175), int32(308155))
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
						F_errmsg(m, int32(443313), int32(0))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							F_errfinish(m, int32(492736), int32(1175), int32(308155))
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
						F_errmsg(m, int32(443364), int32(0))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							F_errfinish(m, int32(492736), int32(1187), int32(308155))
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
	var v47 int32
	_ = v47
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
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int64
	_ = v188
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v207
L2:
	;
	v207 = v3
	goto L1
L3:
	;
	goto L4
L4:
	;
	v17 = F_palloc0(m, int32(68))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
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
		goto L5
	} else {
		goto L7
	}
L7:
	;
	F_ExecPushExprSetupSteps(m, v17, v12)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v38 = v17 + int32(5)
	v40 = v17 + int32(8)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v41 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v47 = v3
	v49 = v3
	goto L12
L10:
	;
	v153 = v3
	goto L11
L11:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	if v159 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v49<<(uint(int32(2))%32))))
	F_ExecInitExprRec(m, v57, v17, v40, v38)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L14
	}
L13:
	;
	if v104 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	if v60 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
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
	v104 = F_lappend_int(m, v47, v101-v83)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L25
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v79
	v81 = v79
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = int32(16)
	v66 = F_palloc(m, int32(640))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if v68 != v60 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v79 = v66
	goto L16
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v81 = v70
	goto L15
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v60 << (uint(int32(1)) % 32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v77 = F_repalloc(m, v74, v60*int32(80))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v79 = v77
	goto L16
L25:
	;
	v107 = v49 + int32(1)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v107 < v108 {
		v47 = v104
		v49 = v107
		goto L12
	} else {
		goto L26
	}
L26:
	;
	goto L13
L27:
	;
	v153 = int32(-1)
	goto L11
L28:
	;
	v112 = int32(0)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v113 <= v112 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v117 = v112
	goto L30
L30:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126+v117<<(uint(int32(2))%32))))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v125+v130*int32(40))+16)) = v134
	v137 = v117 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v137 < v138 {
		v117 = v137
		goto L30
	} else {
		goto L32
	}
L31:
	;
	goto L27
L32:
	;
	goto L31
L33:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v181 + int32(1)
	v187 = v180 + v181*int32(40)
	v188 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v187)+20)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v187)+16)) = v153
	v191 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v187)+12)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v187)+8)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v187)+4)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = v191
	*(*int64)(unsafe.Add(mBase, uint32(v187)+28)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v187)+36)) = v191
	v201 = F_jit_compile_expr(m, v17)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L5
	} else {
		goto L43
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v178
	v180 = v178
	goto L33
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = int32(16)
	v165 = F_palloc(m, int32(640))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L5
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if v167 != v159 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v178 = v165
	goto L34
L39:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v180 = v169
	goto L33
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v159 << (uint(int32(1)) % 32)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v176 = F_repalloc(m, v173, v159*int32(80))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v178 = v176
	goto L34
L43:
	;
	if v201 != 0 {
		v207 = v17
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_ExecReadyInterpretedExpr(m, v17)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v207 = v17
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
	var v20 int32
	_ = v20
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
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v10
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l2
	v16 = v8 + int32(16)
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v46
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v48
	m.G0 = v8 + int32(32)
	return
L2:
	;
	v19 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v20 <= v19 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = v19
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v195 int32
	_ = v195
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
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
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
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
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
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
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
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
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
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v694 int32
	_ = v694
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
	return v694
L4:
	;
	v340 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+4)))
	v341 = v340 - v334
	v342 = int32(0)
	if v342 < v341 {
		goto L93
	} else {
		goto L94
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L90
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L87
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L84
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
		v694 = v4
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v53 = v4
	v54 = v4
	goto L13
L13:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v46+v54<<(uint(int32(2))%32))))
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
	v116 = v53
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
	v113 = F_lappend(m, v53, v112)
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
	v120 = v54 + int32(1)
	if v120 != v43 {
		v53 = v116
		v54 = v120
		goto L13
	} else {
		goto L32
	}
L32:
	;
	goto L14
L33:
	;
	v694 = int32(0)
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
	v694 = v159
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
	v179 = v165
	v180 = v4
	v195 = v168
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
	if v204 <= v180 {
		v210 = int32(0)
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
	v210 = v206 + v180<<(uint(int32(2))%32)
	goto L56
L59:
	;
	v213 = int32(0)
	v333 = v165
	v334 = v213
	v337 = v213
	v339 = v168
	goto L4
L60:
	;
	goto L61
L61:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	if v215 <= v180 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v333 = v179
	v334 = v180
	v337 = int32(0)
	v339 = v195
	goto L4
L63:
	;
	goto L64
L64:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
	v221 = v218 + v180<<(uint(int32(2))%32)
	if v210 == int32(0) {
		v333 = v179
		v334 = v180
		v337 = v221
		v339 = v195
		goto L4
	} else {
		goto L65
	}
L65:
	;
	if v221 == int32(0) {
		v333 = v179
		v334 = v180
		v337 = v221
		v339 = v195
		goto L4
	} else {
		goto L66
	}
L66:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	F_get_range_key_properties(m, v33, v180, v227, v228, v31+int32(44), v31+int32(40), v31+int32(36), v31+int32(32))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v31)+36))
	if v239 == int32(0) {
		v333 = v179
		v334 = v180
		v337 = v221
		v339 = v226
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	if v242 == int32(0) {
		v333 = v179
		v334 = v180
		v337 = v221
		v339 = v226
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v245 = F_CreateExecutorState(m)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v247 = int32(4486928)
	v248 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v245)+100))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v250
	v253 = F_make_partition_op_expr(m, v33, v180, int32(3), v239, v242)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_fix_opfuncids(m, v253)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v258 = F_ExecInitExpr(m, v253, int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v245)+152))
	if v260 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v263 = F_MakePerTupleExprContext(m, v245)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	v265 = v260
	goto L76
L76:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v265)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v267
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v258)+20))
	v272 = m.T0[v271].(func(*base.Module, int32, int32, int32) int32)(m, v258, v265, v31+int32(23))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	v265 = v263
	goto L76
L78:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v248
	F_FreeExecutorState(m, v245)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	if v272 == int32(0) {
		v333 = v179
		v334 = v180
		v337 = v221
		v339 = v226
		goto L4
	} else {
		goto L80
	}
L80:
	;
	v280 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+4)))
	if v180 == v280-int32(1) {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
	v286 = F_make_partition_op_expr(m, v33, v180, int32(3), v285, v239)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v288 = F_lappend(m, v179, v286)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v179 = v288
	v180 = v180 + int32(1)
	v195 = v226
	goto L54
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v79
	F_errmsg_internal(m, int32(46145), v31)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(491059), int32(4317), int32(399126))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errmsg_internal(m, int32(488611), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(491059), int32(4324), int32(399126))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	F_errmsg_internal(m, int32(266504), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(491059), int32(4448), int32(399126))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	v345 = v341
	goto L95
L94:
	;
	v345 = v342
	goto L95
L95:
	;
	v346 = int32(1)
	v360 = int32(0)
	v367 = v4
	v369 = v4
	v373 = v346
	v374 = v346
	goto L96
L96:
	;
	if v360 == v345 {
		v638 = v367
		v640 = v369
		goto L98
	} else {
		goto L99
	}
L97:
	;
	if v638 != 0 {
		goto L179
	} else {
		goto L180
	}
L98:
	;
	goto L97
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v339
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v210 != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v337 != 0 {
		goto L106
	} else {
		goto L107
	}
L101:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)+12))
	v388 = (v210 - v380) >> (uint(int32(2)) % 32)
	goto L100
L102:
	;
	goto L103
L103:
	;
	v384 = int32(0)
	if v379 == v384 {
		v388 = v384
		goto L100
	} else {
		goto L104
	}
L104:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v379)+4))
	v388 = v387
	goto L100
L105:
	;
	v399 = int32(0)
	v405 = v334
	v407 = v398
	v414 = v388
	v415 = v399
	v416 = v399
	goto L110
L106:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)+12))
	v398 = (v337 - v390) >> (uint(int32(2)) % 32)
	goto L105
L107:
	;
	goto L108
L108:
	;
	v394 = int32(0)
	if v389 == v394 {
		v398 = v394
		goto L105
	} else {
		goto L109
	}
L109:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	v398 = v397
	goto L105
L110:
	;
	v429 = int32(0)
	if v379 == v429 {
		v439 = v429
		goto L112
	} else {
		goto L113
	}
L112:
	;
	if v389 == int32(0) {
		v607 = v367
		v609 = v369
		v611 = v373
		v612 = v374
		goto L116
	} else {
		goto L117
	}
L113:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v379)+4))
	if v433 <= v414 {
		v439 = int32(0)
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v379)+12))
	v439 = v435 + v414<<(uint(int32(2))%32)
	goto L112
L115:
	;
	v616 = int32(1)
	v405 = v541
	v407 = v407 + v616
	v414 = v414 + v616
	v415 = v518
	v416 = v539
	goto L110
L116:
	;
	v615 = v360 + int32(1)
	if v612 != 0 {
		v360 = v615
		v367 = v607
		v369 = v609
		v373 = v611
		v374 = v612
		goto L96
	} else {
		goto L177
	}
L117:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	if v442 <= v407 {
		v566 = v415
		v567 = v416
		v571 = v373
		v572 = v374
		goto L118
	} else {
		goto L119
	}
L118:
	;
	if v566 != 0 {
		goto L161
	} else {
		goto L162
	}
L119:
	;
	if v439 == int32(0) {
		v566 = v415
		v567 = v416
		v571 = v373
		v572 = v374
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v389)+12))
	v449 = v446 + v407<<(uint(int32(2))%32)
	if v449 == int32(0) {
		v566 = v415
		v567 = v416
		v571 = v373
		v572 = v374
		goto L118
	} else {
		goto L121
	}
L121:
	;
	v452 = int32(0)
	v455 = v439 + int32(4)
	if v455 == v452 {
		v468 = v452
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v439)))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	v472 = v449 + int32(4)
	if v472 == int32(0) {
		v484 = v452
		goto L125
	} else {
		goto L126
	}
L123:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)+12))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v459)+4))
	if base.Ui32(v460+v461<<(uint(int32(2))%32)) <= base.Ui32(v455) {
		v468 = int32(0)
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	v468 = v466
	goto L122
L125:
	;
	F_get_range_key_properties(m, v33, v405, v469, v470, v31+int32(44), v31+int32(40), v31+int32(36), v31+int32(32))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L128
	}
L126:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v475)+12))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v475)+4))
	if base.Ui32(v476+v477<<(uint(int32(2))%32)) <= base.Ui32(v472) {
		v484 = v452
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	v484 = v482
	goto L125
L128:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v31)+36))
	if base.B2i32(v495 != int32(0))&v374 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	if v405-v334 < v360 {
		v511 = int32(3)
		goto L132
	} else {
		goto L133
	}
L130:
	;
	v518 = v415
	goto L131
L131:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	if base.B2i32(v519 != int32(0))&v373 != 0 {
		goto L141
	} else {
		goto L142
	}
L132:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
	v513 = F_make_partition_op_expr(m, v33, v405, v511, v512, v495)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L139
	}
L133:
	;
	v502 = int32(4)
	v503 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+4)))
	if v405 == v503-int32(1) {
		v511 = v502
		goto L132
	} else {
		goto L134
	}
L134:
	;
	if v468 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v468)+4))
	if v507 == int32(-1) {
		v511 = v502
		goto L132
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v511 = int32(5)
	goto L132
L138:
	;
	goto L137
L139:
	;
	v515 = F_lappend(m, v415, v513)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v518 = v515
	goto L131
L141:
	;
	if v405-v334 < v360 {
		v533 = int32(3)
		goto L144
	} else {
		goto L145
	}
L142:
	;
	v539 = v416
	goto L143
L143:
	;
	v541 = v405 + int32(1)
	if v541-v334 <= v360 {
		goto L115
	} else {
		goto L151
	}
L144:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
	v535 = F_make_partition_op_expr(m, v33, v405, v533, v534, v519)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L149
	}
L145:
	;
	if v484 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v533 = int32(1)
	goto L144
L147:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v484)+4))
	if v528 != int32(1) {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v533 = int32(2)
	goto L144
L149:
	;
	v537 = F_lappend(m, v416, v535)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v539 = v537
	goto L143
L151:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v31)+36))
	if v544 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	if v554 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L153:
	;
	v553 = int32(0)
	goto L152
L154:
	;
	if v468 == int32(0) {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v468)+4))
	if v549 == int32(0) {
		v553 = v374
		goto L152
	} else {
		goto L156
	}
L156:
	;
	goto L153
L157:
	;
	v566 = v518
	v567 = v539
	v571 = int32(0)
	v572 = v553
	goto L118
L158:
	;
	if v484 == int32(0) {
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v484)+4))
	if v559 == int32(0) {
		v566 = v518
		v567 = v539
		v571 = v373
		v572 = v553
		goto L118
	} else {
		goto L160
	}
L160:
	;
	goto L157
L161:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	if int32(2) <= v574 {
		goto L165
	} else {
		goto L166
	}
L162:
	;
	v586 = v367
	goto L163
L163:
	;
	if v567 == int32(0) {
		v607 = v586
		v609 = v369
		v611 = v571
		v612 = v572
		goto L116
	} else {
		goto L170
	}
L164:
	;
	v584 = F_lappend(m, v367, v583)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L169
	}
L165:
	;
	v579 = F_makeBoolExpr(m, int32(0), v566, int32(-1))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v566)+12))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)))
	v583 = v582
	goto L164
L168:
	;
	v583 = v579
	goto L164
L169:
	;
	v586 = v584
	goto L163
L170:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	if int32(2) <= v589 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v599 = F_lappend(m, v369, v598)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L176
	}
L172:
	;
	v594 = F_makeBoolExpr(m, int32(0), v567, int32(-1))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v567)+12))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v596)))
	v598 = v597
	goto L171
L175:
	;
	v598 = v594
	goto L171
L176:
	;
	v607 = v586
	v609 = v599
	v611 = v571
	v612 = v572
	goto L116
L177:
	;
	if v611 != 0 {
		v360 = v615
		v367 = v607
		v369 = v609
		v373 = v611
		v374 = v612
		goto L96
	} else {
		goto L178
	}
L178:
	;
	v638 = v607
	v640 = v609
	goto L98
L179:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v638)+4))
	if int32(2) <= v648 {
		goto L183
	} else {
		goto L184
	}
L180:
	;
	v660 = v333
	goto L181
L181:
	;
	if v640 != 0 {
		goto L188
	} else {
		goto L189
	}
L182:
	;
	v658 = F_lappend(m, v333, v657)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L1
	} else {
		goto L187
	}
L183:
	;
	v653 = F_makeBoolExpr(m, int32(1), v638, int32(-1))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v638)+12))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v655)))
	v657 = v656
	goto L182
L186:
	;
	v657 = v653
	goto L182
L187:
	;
	v660 = v658
	goto L181
L188:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v640)+4))
	if int32(2) <= v661 {
		goto L192
	} else {
		goto L193
	}
L189:
	;
	v673 = v660
	goto L190
L190:
	;
	if v673 != 0 {
		v694 = v673
		goto L3
	} else {
		goto L197
	}
L191:
	;
	v671 = F_lappend(m, v660, v670)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L196
	}
L192:
	;
	v666 = F_makeBoolExpr(m, int32(1), v640, int32(-1))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v640)+12))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v668)))
	v670 = v669
	goto L191
L195:
	;
	v670 = v666
	goto L191
L196:
	;
	v673 = v671
	goto L190
L197:
	;
	if l2 != 0 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v674 = F_get_range_nulltest(m, v33)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L1
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v678 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L202
	}
L201:
	;
	v694 = v674
	goto L3
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v678
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v678
	v685 = F_list_make1_impl(m, int32(1), v31+int32(12))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	v694 = v685
	goto L3
}
