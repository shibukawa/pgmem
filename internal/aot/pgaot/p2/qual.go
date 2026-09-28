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
								F_errfinish(m, int32(_a_F_AddQual_1), int32(1193), int32(_a_F_AddQual_2))
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
							F_errfinish(m, int32(_a_F_AddQual_1), int32(1193), int32(_a_F_AddQual_2))
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
							F_errfinish(m, int32(_a_F_AddQual_1), int32(1205), int32(_a_F_AddQual_2))
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
							v42 = F_query_or_expression_tree_walker_impl(m, v30, int32(1125), int32(0), int32(3))
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
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int64
	_ = v181
	var v184 int32
	_ = v184
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l0 == v3 {
		v200 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v200
L2:
	;
	v17 = F_palloc0(m, int32(72))
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
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(386)
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
		v147 = v3
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	if v151 == int32(0) {
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
	v111 = int32(-1)
	if v105 == int32(0) {
		v147 = v111
		goto L7
	} else {
		goto L24
	}
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	if v60 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	v84 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v83 + v84
	v89 = v82 + v83*int32(40)
	v90 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+20)) = v90
	*(*int64)(unsafe.Add(mBase, uint32(v89)+12)) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = int32(39)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+28)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v89)+36)) = int32(0)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	v105 = F_lappend_int(m, v48, v102-v84)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L3
	} else {
		goto L22
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v80
	v82 = v80
	goto L12
L14:
	;
	v63 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v63
	v67 = F_palloc_mul(m, int32(40), v63)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L3
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	if v69 != v60 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v80 = v67
	goto L13
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v82 = v71
	goto L12
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v60 << (uint(int32(1)) % 32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v78 = F_repalloc(m, v75, v60*int32(80))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v80 = v78
	goto L13
L22:
	;
	v108 = v49 + int32(1)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v108 < v109 {
		v48 = v105
		v49 = v108
		goto L9
	} else {
		goto L23
	}
L23:
	;
	goto L10
L24:
	;
	v114 = int32(0)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v115 <= v114 {
		v147 = v111
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v119 = v114
	goto L26
L26:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128+v119<<(uint(int32(2))%32))))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v127+v132*int32(40))+16)) = v136
	v139 = v119 + int32(1)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v139 < v140 {
		v119 = v139
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v147 = v111
	goto L7
L28:
	;
	goto L27
L29:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v174 + int32(1)
	v180 = v173 + v174*int32(40)
	v181 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v180)+20)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v180)+16)) = v147
	v184 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v180)+12)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v180)+8)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v180)+4)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v180))) = v184
	*(*int64)(unsafe.Add(mBase, uint32(v180)+28)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v180)+36)) = v184
	v194 = F_jit_compile_expr(m, v17)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L3
	} else {
		goto L39
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v171
	v173 = v171
	goto L29
L31:
	;
	v154 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v154
	v158 = F_palloc_mul(m, int32(40), v154)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L3
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	if v160 != v151 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v171 = v158
	goto L30
L35:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v173 = v162
	goto L29
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v151 << (uint(int32(1)) % 32)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v169 = F_repalloc(m, v166, v151*int32(80))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v171 = v169
	goto L30
L39:
	;
	if v194 != 0 {
		v200 = v17
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_ExecReadyInterpretedExpr(m, v17)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	v200 = v17
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
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
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v194 int32
	_ = v194
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int64
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v363 int32
	_ = v363
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v530 int32
	_ = v530
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
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
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
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	v4 = int32(0)
	v30 = m.G0
	v32 = v30 - int32(48)
	m.G0 = v32
	v34 = F_RelationGetPartitionKey(m, l0)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v38 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	m.G0 = v32 + int32(48)
	return v701
L4:
	;
	v346 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34)+4)))
	v347 = v346 - v339
	v348 = int32(0)
	if v348 < v347 {
		goto L92
	} else {
		goto L93
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L89
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L86
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L83
	}
L8:
	;
	v42 = F_RelationGetPartitionDesc(m, l0, int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
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
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v44 <= int32(0) {
		v701 = v4
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v52 = v4
	v53 = v4
	goto L13
L13:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v47+v53<<(uint(int32(2))%32))))
	v83 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(v81))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	if v120 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L15:
	;
	if v83 == int32(0) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v89 = F_SysCacheGetAttrNotNull(m, int32(57), v83, int32(34))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v92 = F_text_to_cstring(m, base.I32_wrap_i64(v89))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v94 = F_stringToNode(m, v92)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v96 != int32(98) {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+5)))
	if v99 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v103 = F_get_qual_for_range(m, l0, v94, int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L26
	}
L22:
	;
	v120 = v52
	goto L23
L23:
	;
	F_ReleaseCatCache(m, v83)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L31
	}
L24:
	;
	v117 = F_lappend(m, v52, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L30
	}
L25:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v116 = v115
	goto L24
L26:
	;
	if v103 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v107 < int32(2) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v112 = F_makeBoolExpr(m, int32(0), v103, int32(-1))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v116 = v112
	goto L24
L30:
	;
	v120 = v117
	goto L23
L31:
	;
	v124 = v53 + int32(1)
	if v124 != v44 {
		v52 = v120
		v53 = v124
		goto L13
	} else {
		goto L32
	}
L32:
	;
	goto L14
L33:
	;
	v701 = int32(0)
	goto L3
L34:
	;
	goto L35
L35:
	;
	v130 = F_get_range_nulltest(m, v34)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if int32(2) <= v132 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v142 = F_lappend(m, v130, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L42
	}
L38:
	;
	v137 = F_makeBoolExpr(m, int32(1), v120, int32(-1))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v141 = v140
	goto L37
L41:
	;
	v141 = v137
	goto L37
L42:
	;
	v145 = F_makeBoolExpr(m, int32(0), v142, int32(-1))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v145
	v153 = F_list_make1_impl(m, int32(1), v32+int32(8))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v156 = F_makeBoolExpr(m, int32(2), v153, int32(-1))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v156
	v163 = F_list_make1_impl(m, int32(1), v32+int32(4))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v701 = v163
	goto L3
L47:
	;
	v167 = F_get_range_nulltest(m, v34)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	v169 = v4
	goto L49
L49:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	if v170 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v169 = v167
	goto L49
L51:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
	v172 = v171
	goto L53
L52:
	;
	v172 = v4
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+44)) = v172
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v180 = v169
	v182 = v4
	v194 = v172
	goto L54
L54:
	;
	v205 = int32(0)
	if v175 == v205 {
		v215 = v205
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if v174 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	if v209 <= v182 {
		v215 = int32(0)
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	v215 = v211 + v182<<(uint(int32(2))%32)
	goto L56
L59:
	;
	v218 = int32(0)
	v338 = v169
	v339 = v218
	v344 = v172
	v345 = v218
	goto L4
L60:
	;
	goto L61
L61:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	if v220 <= v182 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v338 = v180
	v339 = v182
	v344 = v194
	v345 = int32(0)
	goto L4
L63:
	;
	goto L64
L64:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	v226 = v223 + v182<<(uint(int32(2))%32)
	v227 = int32(0)
	if base.B2i32(v215 == v227)|base.B2i32(v223 == v227) != 0 {
		v338 = v180
		v339 = v182
		v344 = v194
		v345 = v226
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v32)+44))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	F_get_range_key_properties(m, v34, v182, v233, v234, v32+int32(44), v32+int32(40), v32+int32(36), v32+int32(32))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v32)+36))
	if v245 == int32(0) {
		v338 = v180
		v339 = v182
		v344 = v232
		v345 = v226
		goto L4
	} else {
		goto L67
	}
L67:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
	if v248 == int32(0) {
		v338 = v180
		v339 = v182
		v344 = v232
		v345 = v226
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v251 = F_CreateExecutorState(m)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v253 = int32(_a_F_get_qual_for_range_0)
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_get_qual_for_range[0]))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v251)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_get_qual_for_range[0])) = v256
	v259 = F_make_partition_op_expr(m, v34, v182, int32(3), v245, v248)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_fix_opfuncids(m, v259)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v264 = F_ExecInitExpr(m, v259, int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v251)+152))
	if v266 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v269 = F_MakePerTupleExprContext(m, v251)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	v271 = v266
	goto L75
L75:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v271)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_get_qual_for_range[0])) = v273
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v264)+24))
	v278 = m.T0[v277].(func(*base.Module, int32, int32, int32) int64)(m, v264, v271, v32+int32(23))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L77
	}
L76:
	;
	v271 = v269
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_qual_for_range[0])) = v254
	F_FreeExecutorState(m, v251)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	if v278 == int64(0) {
		v338 = v180
		v339 = v182
		v344 = v232
		v345 = v226
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v286 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34)+4)))
	if v182 == v286-int32(1) {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	v292 = F_make_partition_op_expr(m, v34, v182, int32(3), v291, v245)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v294 = F_lappend(m, v180, v292)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v180 = v294
	v182 = v182 + int32(1)
	v194 = v232
	goto L54
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v81
	F_errmsg_internal(m, int32(_a_F_get_qual_for_range_1), v32)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_get_qual_for_range_2), int32(_a_F_get_qual_for_range_3), int32(_a_F_get_qual_for_range_4))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
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
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_get_qual_for_range_2), int32(_a_F_get_qual_for_range_6), int32(_a_F_get_qual_for_range_4))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
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
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_get_qual_for_range_2), int32(_a_F_get_qual_for_range_8), int32(_a_F_get_qual_for_range_4))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
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
	v351 = v347
	goto L94
L93:
	;
	v351 = v348
	goto L94
L94:
	;
	v352 = int32(1)
	v363 = int32(0)
	v372 = v4
	v375 = v4
	v378 = v352
	v379 = v352
	goto L95
L95:
	;
	if v363 == v351 {
		v646 = v372
		v649 = v375
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if v646 != 0 {
		goto L176
	} else {
		goto L177
	}
L97:
	;
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+44)) = v344
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v215 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v345 != 0 {
		goto L105
	} else {
		goto L106
	}
L100:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v386)+12))
	v395 = (v215 - v387) >> (uint(int32(2)) % 32)
	goto L99
L101:
	;
	goto L102
L102:
	;
	v391 = int32(0)
	if v386 == v391 {
		v395 = v391
		goto L99
	} else {
		goto L103
	}
L103:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v386)+4))
	v395 = v394
	goto L99
L104:
	;
	v406 = int32(0)
	v413 = v339
	v417 = v405
	v420 = v406
	v421 = v406
	v430 = v395
	goto L109
L105:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)+12))
	v405 = (v345 - v397) >> (uint(int32(2)) % 32)
	goto L104
L106:
	;
	goto L107
L107:
	;
	v401 = int32(0)
	if v396 == v401 {
		v405 = v401
		goto L104
	} else {
		goto L108
	}
L108:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	v405 = v404
	goto L104
L109:
	;
	v437 = int32(0)
	if v386 == v437 {
		v447 = v437
		goto L111
	} else {
		goto L112
	}
L111:
	;
	if v396 == int32(0) {
		v616 = v372
		v617 = v375
		v618 = v378
		v619 = v379
		goto L115
	} else {
		goto L116
	}
L112:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v386)+4))
	if v441 <= v430 {
		v447 = int32(0)
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v386)+12))
	v447 = v443 + v430<<(uint(int32(2))%32)
	goto L111
L114:
	;
	v625 = int32(1)
	v413 = v543
	v417 = v417 + v625
	v420 = v520
	v421 = v541
	v430 = v430 + v625
	goto L109
L115:
	;
	if v618|v619 != 0 {
		v363 = v363 + int32(1)
		v372 = v616
		v375 = v617
		v378 = v618
		v379 = v619
		goto L95
	} else {
		goto L175
	}
L116:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	if base.B2i32(v447 == int32(0))|base.B2i32(v452 <= v417) != 0 {
		v573 = v420
		v574 = v421
		v578 = v378
		v579 = v379
		goto L117
	} else {
		goto L118
	}
L117:
	;
	if v573 != 0 {
		goto L159
	} else {
		goto L160
	}
L118:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v396)+12))
	if v455 == int32(0) {
		v573 = v420
		v574 = v421
		v578 = v378
		v579 = v379
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v458 = int32(0)
	v461 = v447 + int32(4)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v462)+12))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v462)+4))
	if base.Ui32(v461) < base.Ui32(v463+v464<<(uint(int32(2))%32)) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v461)))
	v470 = v469
	goto L122
L121:
	;
	v470 = v458
	goto L122
L122:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	v472 = int32(2)
	v474 = v455 + v417<<(uint(v472)%32)
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v474)))
	v477 = v474 + int32(4)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)+12))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v478)+4))
	if base.Ui32(v477) < base.Ui32(v479+v480<<(uint(v472)%32)) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	v486 = v485
	goto L125
L124:
	;
	v486 = v458
	goto L125
L125:
	;
	F_get_range_key_properties(m, v34, v413, v471, v475, v32+int32(44), v32+int32(40), v32+int32(36), v32+int32(32))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v32)+36))
	if base.B2i32(v497 != int32(0))&v379 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	if v413-v339 < v363 {
		v513 = int32(3)
		goto L130
	} else {
		goto L131
	}
L128:
	;
	v520 = v420
	goto L129
L129:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
	if base.B2i32(v521 != int32(0))&v378 != 0 {
		goto L139
	} else {
		goto L140
	}
L130:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	v515 = F_make_partition_op_expr(m, v34, v413, v513, v514, v497)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L137
	}
L131:
	;
	v504 = int32(4)
	v505 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34)+4)))
	if v413 == v505-int32(1) {
		v513 = v504
		goto L130
	} else {
		goto L132
	}
L132:
	;
	if v470 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	if v509 == int32(-1) {
		v513 = v504
		goto L130
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v513 = int32(5)
	goto L130
L136:
	;
	goto L135
L137:
	;
	v517 = F_lappend(m, v420, v515)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v520 = v517
	goto L129
L139:
	;
	if v413-v339 < v363 {
		v535 = int32(3)
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v541 = v421
	goto L141
L141:
	;
	v543 = v413 + int32(1)
	if v543-v339 <= v363 {
		goto L114
	} else {
		goto L149
	}
L142:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	v537 = F_make_partition_op_expr(m, v34, v413, v535, v536, v521)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L147
	}
L143:
	;
	if v486 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v535 = int32(1)
	goto L142
L145:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v486)+4))
	if v530 != int32(1) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v535 = int32(2)
	goto L142
L147:
	;
	v539 = F_lappend(m, v421, v537)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v541 = v539
	goto L141
L149:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v32)+36))
	v547 = int32(0)
	if base.B2i32(v546 == v547)|base.B2i32(v470 == v547) == v547 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
	v560 = int32(0)
	if base.B2i32(v559 == v560)|base.B2i32(v486 == v560) == v560 {
		goto L155
	} else {
		goto L156
	}
L151:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	if v554 == int32(0) {
		v558 = v379
		goto L150
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v558 = int32(0)
	goto L150
L154:
	;
	goto L153
L155:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v486)+4))
	if v567 == int32(0) {
		v573 = v520
		v574 = v541
		v578 = v378
		v579 = v558
		goto L117
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v573 = v520
	v574 = v541
	v578 = int32(0)
	v579 = v558
	goto L117
L158:
	;
	goto L157
L159:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	if int32(2) <= v582 {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	v594 = v372
	goto L161
L161:
	;
	if v574 == int32(0) {
		v616 = v594
		v617 = v375
		v618 = v578
		v619 = v579
		goto L115
	} else {
		goto L168
	}
L162:
	;
	v592 = F_lappend(m, v372, v591)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L167
	}
L163:
	;
	v587 = F_makeBoolExpr(m, int32(0), v573, int32(-1))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v573)+12))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v589)))
	v591 = v590
	goto L162
L166:
	;
	v591 = v587
	goto L162
L167:
	;
	v594 = v592
	goto L161
L168:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	if int32(2) <= v597 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v607 = F_lappend(m, v375, v606)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L174
	}
L170:
	;
	v602 = F_makeBoolExpr(m, int32(0), v574, int32(-1))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v574)+12))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v604)))
	v606 = v605
	goto L169
L173:
	;
	v606 = v602
	goto L169
L174:
	;
	v616 = v594
	v617 = v607
	v618 = v578
	v619 = v579
	goto L115
L175:
	;
	v646 = v616
	v649 = v617
	goto L97
L176:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v646)+4))
	if int32(2) <= v658 {
		goto L180
	} else {
		goto L181
	}
L177:
	;
	v670 = v338
	goto L178
L178:
	;
	if v649 != 0 {
		goto L185
	} else {
		goto L186
	}
L179:
	;
	v668 = F_lappend(m, v338, v667)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L184
	}
L180:
	;
	v663 = F_makeBoolExpr(m, int32(1), v646, int32(-1))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v646)+12))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v665)))
	v667 = v666
	goto L179
L183:
	;
	v667 = v663
	goto L179
L184:
	;
	v670 = v668
	goto L178
L185:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v649)+4))
	if int32(2) <= v671 {
		goto L189
	} else {
		goto L190
	}
L186:
	;
	v683 = v670
	goto L187
L187:
	;
	if v683 != 0 {
		v701 = v683
		goto L3
	} else {
		goto L194
	}
L188:
	;
	v681 = F_lappend(m, v670, v680)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L193
	}
L189:
	;
	v676 = F_makeBoolExpr(m, int32(1), v649, int32(-1))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v649)+12))
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v678)))
	v680 = v679
	goto L188
L192:
	;
	v680 = v676
	goto L188
L193:
	;
	v683 = v681
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
	v684 = F_get_range_nulltest(m, v34)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v688 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L199
	}
L198:
	;
	v701 = v684
	goto L3
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v688
	v695 = F_list_make1_impl(m, int32(1), v32+int32(12))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v701 = v695
	goto L3
}
