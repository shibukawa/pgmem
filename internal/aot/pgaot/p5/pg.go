package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_GetPGVariable(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
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
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v18 = l0
	v19 = int32(_a_F_GetPGVariable_0)
	goto L3
L1:
	;
	m.G0 = v13 + int32(32)
	return
L2:
	;
	if v60 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L3:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v60 = base.I32_extend8_s(v40) - base.I32_extend8_s(v49)
	goto L2
L5:
	;
	v28 = int32(1)
	if base.Ui32((v23-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	if v22 != 0 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v60 = int32(1)
	goto L2
L10:
	;
	v27 = int32(-1)
	goto L12
L11:
	;
	v27 = int32(0)
	goto L12
L12:
	;
	v60 = v27
	goto L2
L13:
	;
	v40 = v23 | int32(32)
	goto L15
L14:
	;
	v40 = v23
	goto L15
L15:
	;
	if base.Ui32((v22-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v49 = v22 | int32(32)
	goto L18
L17:
	;
	v49 = v22
	goto L18
L18:
	;
	if v40 == v49&int32(255) {
		v18 = v18 + v28
		v19 = v19 + v28
		goto L3
	} else {
		goto L19
	}
L19:
	;
	goto L4
L20:
	;
	v63 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v63)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)) = uint16(v63)
	v70 = F_get_guc_variables(m, v13+int32(28))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v192 = F_GetConfigOptionByName(m, l0, v13+int32(16), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L23
	} else {
		goto L63
	}
L23:
	;
	return
L24:
	;
	v73 = F_CreateTemplateTupleDesc(m, int32(3))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	F_TupleDescInitBuiltinEntry(m, v73, int32(1), int32(_a_F_GetPGVariable_1), int32(25))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	F_TupleDescInitBuiltinEntry(m, v73, int32(2), int32(_a_F_GetPGVariable_2), int32(25))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	F_TupleDescInitBuiltinEntry(m, v73, int32(3), int32(_a_F_GetPGVariable_3), int32(25))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v91 = F_begin_tup_output_tupdesc(m, l1, v73, int32(_a_F_GetPGVariable_4))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	if int32(0) < v93 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v96 = v63
	goto L33
L31:
	;
	goto L32
L32:
	;
	F_end_tup_output(m, v91)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L23
	} else {
		goto L62
	}
L33:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v70+v96<<(uint(int32(2))%32))))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+20))
	if v110&int32(4) != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v174 = v96 + int32(1)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	if v174 < v175 {
		v96 = v174
		goto L33
	} else {
		goto L61
	}
L36:
	;
	if v110&int32(1024) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_GetPGVariable[0]))
	v118 = F_has_privs_of_role(m, v116, int32(3374))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L23
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v123 = F_cstring_to_text(m, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L23
	} else {
		goto L42
	}
L40:
	;
	if v118 == int32(0) {
		goto L35
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v123
	v126 = int32(1)
	v127 = int32(0)
	v131 = F_ShowGUCOption(m, v109, v126)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L23
	} else {
		goto L43
	}
L43:
	;
	if v131 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v134 = F_cstring_to_text(m, v131)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L23
	} else {
		goto L47
	}
L45:
	;
	v136 = v127
	v137 = v126
	goto L46
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)) = uint8(v137)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v136
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	if v140 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v136 = v134
	v137 = int32(0)
	goto L46
L48:
	;
	v142 = F_cstring_to_text(m, v140)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L23
	} else {
		goto L51
	}
L49:
	;
	v144 = v127
	v145 = v126
	goto L50
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v145)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v144
	F_do_tup_output(m, v91, v13+int32(16), v13+int32(12))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L23
	} else {
		goto L52
	}
L51:
	;
	v144 = v142
	v145 = int32(0)
	goto L50
L52:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	F_pfree(m, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L23
	} else {
		goto L53
	}
L53:
	;
	if v131 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_pfree(m, v131)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L23
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	if v162 == int32(0) {
		goto L35
	} else {
		goto L59
	}
L57:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	F_pfree(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L23
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	F_pfree(m, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L23
	} else {
		goto L60
	}
L60:
	;
	goto L35
L61:
	;
	goto L34
L62:
	;
	goto L1
L63:
	;
	v195 = F_CreateTemplateTupleDesc(m, int32(1))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L23
	} else {
		goto L64
	}
L64:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	F_TupleDescInitBuiltinEntry(m, v195, int32(1), v198, int32(25))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L23
	} else {
		goto L65
	}
L65:
	;
	v203 = F_begin_tup_output_tupdesc(m, l1, v195, int32(_a_F_GetPGVariable_4))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L23
	} else {
		goto L66
	}
L66:
	;
	v205 = F_cstring_to_text(m, v192)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L23
	} else {
		goto L67
	}
L67:
	;
	v207 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+12)) = uint8(v207)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v205
	F_do_tup_output(m, v203, v13+int32(28), v13+int32(12))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L23
	} else {
		goto L68
	}
L68:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	F_pfree(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L23
	} else {
		goto L69
	}
L69:
	;
	F_end_tup_output(m, v203)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L23
	} else {
		goto L70
	}
L70:
	;
	goto L1
}
func F__PG_init_auto_explain(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v110 float64
	_ = v110
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	v6 = int32(-1)
	F_DefineCustomIntVariable(m, int32(_a_F__PG_init_auto_explain_0), int32(_a_F__PG_init_auto_explain_1), int32(_a_F__PG_init_auto_explain_2), int32(_a_F__PG_init_auto_explain_3), v6, v6, int32(2147483647), int32(5), int32(268435456))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v17 = int32(-1)
		F_DefineCustomIntVariable(m, int32(_a_F__PG_init_auto_explain_4), int32(_a_F__PG_init_auto_explain_5), int32(_a_F__PG_init_auto_explain_6), int32(_a_F__PG_init_auto_explain_7), v17, v17, int32(2147483647), int32(5), int32(83886080))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v26 = int32(0)
			F_DefineCustomBoolVariable(m, int32(_a_F__PG_init_auto_explain_8), int32(_a_F__PG_init_auto_explain_9), v26, int32(_a_F__PG_init_auto_explain_10), v26, int32(5))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v34 = int32(0)
				F_DefineCustomBoolVariable(m, int32(_a_F__PG_init_auto_explain_11), int32(_a_F__PG_init_auto_explain_12), v34, int32(_a_F__PG_init_auto_explain_13), v34, int32(5))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					v42 = int32(0)
					F_DefineCustomBoolVariable(m, int32(_a_F__PG_init_auto_explain_14), int32(_a_F__PG_init_auto_explain_15), v42, int32(_a_F__PG_init_auto_explain_16), v42, int32(5))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						v50 = int32(0)
						F_DefineCustomBoolVariable(m, int32(_a_F__PG_init_auto_explain_17), int32(_a_F__PG_init_auto_explain_18), v50, int32(_a_F__PG_init_auto_explain_19), v50, int32(5))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v58 = int32(0)
							F_DefineCustomBoolVariable(m, int32(_a_F__PG_init_auto_explain_20), int32(_a_F__PG_init_auto_explain_21), v58, int32(_a_F__PG_init_auto_explain_22), v58, int32(5))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								F_DefineCustomBoolVariable(m, int32(_a_F__PG_init_auto_explain_23), int32(_a_F__PG_init_auto_explain_24), int32(_a_F__PG_init_auto_explain_25), int32(_a_F__PG_init_auto_explain_26), int32(0), int32(5))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
								} else {
									v74 = int32(0)
									F_DefineCustomEnumVariable(m, int32(_a_F__PG_init_auto_explain_27), int32(_a_F__PG_init_auto_explain_28), v74, int32(_a_F__PG_init_auto_explain_29), v74, int32(_a_F__PG_init_auto_explain_30), int32(5))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										F_DefineCustomEnumVariable(m, int32(_a_F__PG_init_auto_explain_31), int32(_a_F__PG_init_auto_explain_32), int32(0), int32(_a_F__PG_init_auto_explain_33), int32(15), int32(_a_F__PG_init_auto_explain_34), int32(5))
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											v92 = int32(0)
											F_DefineCustomBoolVariable(m, int32(_a_F__PG_init_auto_explain_35), int32(_a_F__PG_init_auto_explain_36), v92, int32(_a_F__PG_init_auto_explain_37), v92, int32(5))
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
											} else {
												F_DefineCustomBoolVariable(m, int32(_a_F__PG_init_auto_explain_38), int32(_a_F__PG_init_auto_explain_39), int32(0), int32(_a_F__PG_init_auto_explain_40), int32(1), int32(5))
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return
												} else {
													v110 = float64(1)
													F_DefineCustomRealVariable(m, int32(_a_F__PG_init_auto_explain_41), int32(_a_F__PG_init_auto_explain_42), int32(0), int32(_a_F__PG_init_auto_explain_43), v110, float64(0), v110, int32(5))
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return
													} else {
														F_MarkGUCPrefixReserved(m, int32(_a_F__PG_init_auto_explain_44))
														mBase = m.M
														v118 = m.ExcPending
														if v118 != 0 {
															return
														} else {
															v120 = int32(_a_F__PG_init_auto_explain_45)
															v121 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_auto_explain[0]))
															*(*int32)(unsafe.Add(mBase, _c_F__PG_init_auto_explain[1])) = v121
															*(*int32)(unsafe.Add(mBase, _c_F__PG_init_auto_explain[0])) = int32(_a_F__PG_init_auto_explain_46)
															v126 = int32(_a_F__PG_init_auto_explain_47)
															v127 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_auto_explain[2]))
															*(*int32)(unsafe.Add(mBase, _c_F__PG_init_auto_explain[2])) = int32(_a_F__PG_init_auto_explain_48)
															*(*int32)(unsafe.Add(mBase, _c_F__PG_init_auto_explain[3])) = v127
															v133 = int32(_a_F__PG_init_auto_explain_49)
															v134 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_auto_explain[4]))
															*(*int32)(unsafe.Add(mBase, _c_F__PG_init_auto_explain[4])) = int32(_a_F__PG_init_auto_explain_50)
															*(*int32)(unsafe.Add(mBase, _c_F__PG_init_auto_explain[5])) = v134
															v140 = int32(_a_F__PG_init_auto_explain_51)
															v141 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_auto_explain[6]))
															*(*int32)(unsafe.Add(mBase, _c_F__PG_init_auto_explain[6])) = int32(_a_F__PG_init_auto_explain_52)
															*(*int32)(unsafe.Add(mBase, _c_F__PG_init_auto_explain[7])) = v141
															return
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
		}
	}
}
func F_do_pg_abort_backup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	if l1 == int32(0) {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_abort_backup[0])))
		if v6&int32(1) == int32(0) {
			return
		} else {
			F_WALInsertLockAcquireExclusive(m)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_abort_backup[1]))
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+164))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+164)) = v15 - int32(1)
				v20 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_abort_backup[0])) = uint8(v20)
				F_WALInsertLockRelease(m)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					if l1 != 0 {
						return
					} else {
						v26 = F_errstart(m, int32(19), int32(0))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							if v26 == int32(0) {
								return
							} else {
								F_errmsg(m, int32(_a_F_do_pg_abort_backup_0), int32(0))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_do_pg_abort_backup_1), int32(_a_F_do_pg_abort_backup_2), int32(_a_F_do_pg_abort_backup_3))
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_WALInsertLockAcquireExclusive(m)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_abort_backup[1]))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+164))
			*(*int32)(unsafe.Add(mBase, uint32(v14)+164)) = v15 - int32(1)
			v20 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_abort_backup[0])) = uint8(v20)
			F_WALInsertLockRelease(m)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				if l1 != 0 {
					return
				} else {
					v26 = F_errstart(m, int32(19), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						if v26 == int32(0) {
							return
						} else {
							F_errmsg(m, int32(_a_F_do_pg_abort_backup_0), int32(0))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_do_pg_abort_backup_1), int32(_a_F_do_pg_abort_backup_2), int32(_a_F_do_pg_abort_backup_3))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_do_pg_backup_start(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
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
	var v44 int32
	_ = v44
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v322 int32
	_ = v322
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v356 int32
	_ = v356
	var v366 int32
	_ = v366
	var v375 int64
	_ = v375
	var v376 int32
	_ = v376
	var v407 int32
	_ = v407
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int64
	_ = v423
	var v425 int64
	_ = v425
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int64
	_ = v466
	var v467 int32
	_ = v467
	var v472 int64
	_ = v472
	var v483 int32
	_ = v483
	var v492 int32
	_ = v492
	var v502 int32
	_ = v502
	var v512 int32
	_ = v512
	var v523 int32
	_ = v523
	var v533 int32
	_ = v533
	var v534 int64
	_ = v534
	var v536 int32
	_ = v536
	var v537 int64
	_ = v537
	var v547 int32
	_ = v547
	var v555 int32
	_ = v555
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v635 int32
	_ = v635
	var v640 int64
	_ = v640
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v691 int32
	_ = v691
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v738 int32
	_ = v738
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v806 int32
	_ = v806
	var v820 int32
	_ = v820
	var v832 int32
	_ = v832
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v858 int32
	_ = v858
	var v867 int32
	_ = v867
	var v874 int32
	_ = v874
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v926 int32
	_ = v926
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v958 int32
	_ = v958
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v980 int32
	_ = v980
	var v990 int32
	_ = v990
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1053 int32
	_ = v1053
	var v1060 int64
	_ = v1060
	var v1071 int32
	_ = v1071
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1109 int32
	_ = v1109
	var v1117 int32
	_ = v1117
	var v1141 int32
	_ = v1141
	var v1142 int64
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	v6 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(2352)
	m.G0 = v26
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = int32(44)
	goto L3
L2:
	;
	v30 = int32(40)
	goto L3
L3:
	;
	v33 = l1
	v38 = v6
	v39 = v6
	v40 = v6
	v41 = v6
	v42 = v6
	v43 = int32(-1)
	v44 = v6
	goto L4
L4:
	;
	goto L7
L5:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L6:
	;
	goto L5
L7:
	;
	if v43 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v1141 = int32(m.ExcTag)
	v1142 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1141 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L10:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_backup_start[0])))
	if v59 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v345 = v38
	v346 = v39
	v347 = v42
	v348 = v44
	goto L12
L12:
	;
	if v347 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L13:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[1]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+316))
	v67 = base.B2i32(v65 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_backup_start[0])) = uint8(v67)
	v69 = v67
	goto L15
L14:
	;
	v69 = int32(0)
	goto L15
L15:
	;
	if v69 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v69)
	v130 = F_strlen(m, l0)
	mBase = m.M
	if base.Ui32(int32(1025)) <= base.Ui32(v130) {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[2]))
	if int32(0) < v71 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v69)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v69)
	F_errcode(m, int32(325))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v69)
	F_errmsg(m, int32(_a_F_do_pg_backup_start_0), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v69)
	F_errhint(m, int32(_a_F_do_pg_backup_start_1), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v69)
	F_errfinish(m, int32(_a_F_do_pg_backup_start_2), int32(_a_F_do_pg_backup_start_3), int32(_a_F_do_pg_backup_start_4))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	goto L6
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v69)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L9
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v69)
	goto L34
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v69)
	F_errcode(m, int32(50856066))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v69)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = int32(1024)
	F_errmsg(m, int32(_a_F_do_pg_backup_start_5), v26-int32(-64))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v69)
	F_errfinish(m, int32(_a_F_do_pg_backup_start_2), int32(_a_F_do_pg_backup_start_6), int32(_a_F_do_pg_backup_start_4))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	goto L6
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v69)
	F_WALInsertLockAcquireExclusive(m)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L9
	} else {
		goto L62
	}
L32:
	;
	v298 = F_strlen(m, v287)
	mBase = m.M
	goto L31
L34:
	;
	goto L35
L35:
	;
	v188 = int32(1024)
	if (l3^l0)&int32(3) != 0 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v291 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v288))) = uint8(v291)
	goto L32
L37:
	;
	v272 = v267
	v273 = v268
	v274 = v269
	goto L58
L38:
	;
	if v262 == int32(0) {
		v287 = v260
		v288 = v261
		goto L36
	} else {
		goto L57
	}
L39:
	;
	v260 = l0
	v261 = l3
	v262 = v188
	goto L38
L40:
	;
	goto L41
L41:
	;
	v192 = int32(0)
	if base.B2i32(l0&int32(3) == v192)|int32(0) == v192 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v228 == int32(0) {
		v287 = v225
		v288 = v226
		goto L36
	} else {
		goto L51
	}
L43:
	;
	v204 = l0
	v205 = l3
	v206 = v188
	goto L46
L44:
	;
	goto L45
L45:
	;
	v225 = l0
	v226 = l3
	v227 = v188
	v228 = int32(1)
	goto L42
L46:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	*(*uint8)(unsafe.Add(mBase, uint32(v205))) = uint8(v208)
	if v208 == int32(0) {
		v267 = v204
		v268 = v205
		v269 = v206
		goto L37
	} else {
		goto L48
	}
L47:
	;
	v225 = v219
	v226 = v213
	v227 = v215
	v228 = v217
	goto L42
L48:
	;
	v212 = int32(1)
	v213 = v205 + v212
	v215 = v206 - v212
	v216 = int32(0)
	v217 = base.B2i32(v215 != v216)
	v219 = v204 + v212
	if v219&int32(3) == v216 {
		v225 = v219
		v226 = v213
		v227 = v215
		v228 = v217
		goto L42
	} else {
		goto L49
	}
L49:
	;
	if v215 != 0 {
		v204 = v219
		v205 = v213
		v206 = v215
		goto L46
	} else {
		goto L50
	}
L50:
	;
	goto L47
L51:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	if base.B2i32(v231 == int32(0))|base.B2i32(base.Ui32(v227) < base.Ui32(int32(4))) != 0 {
		v260 = v225
		v261 = v226
		v262 = v227
		goto L38
	} else {
		goto L52
	}
L52:
	;
	v238 = v225
	v239 = v226
	v240 = v227
	goto L53
L53:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v246 = int32(-2139062144)
	if (int32(16843008)-v243|v243)&v246 != v246 {
		v267 = v238
		v268 = v239
		v269 = v240
		goto L37
	} else {
		goto L55
	}
L54:
	;
	v260 = v254
	v261 = v252
	v262 = v256
	goto L38
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239))) = v243
	v251 = int32(4)
	v252 = v239 + v251
	v254 = v238 + v251
	v256 = v240 - v251
	if base.Ui32(int32(3)) < base.Ui32(v256) {
		v238 = v254
		v239 = v252
		v240 = v256
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v267 = v260
	v268 = v261
	v269 = v262
	goto L37
L58:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	*(*uint8)(unsafe.Add(mBase, uint32(v273))) = uint8(v276)
	if v276 == int32(0) {
		v287 = v272
		v288 = v273
		goto L36
	} else {
		goto L60
	}
L59:
	;
	v287 = v283
	v288 = v281
	goto L36
L60:
	;
	v280 = int32(1)
	v281 = v273 + v280
	v283 = v272 + v280
	v285 = v274 - v280
	if v285 != 0 {
		v272 = v283
		v273 = v281
		v274 = v285
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[1]))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v310)+164)) = v311 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v69)
	F_WALInsertLockRelease(m)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v39
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v69)
	F_before_shmem_exit(m, int32(407), int32(1))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[3]))
	v337 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[4]))
	goto L65
L65:
	;
	v339 = v26 + int32(2160)
	*(*int32)(unsafe.Add(mBase, uint32(v339)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v339))) = v26 + int32(72)
	goto L68
L66:
	;
	v345 = v335
	v346 = v337
	v347 = int32(0)
	v348 = v69
	goto L12
L68:
	;
	goto L66
L69:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[4])) = v26 + int32(2160)
	v356 = v348 & int32(1)
	if v356 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[3])) = v345
	*(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[4])) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	v1094 = int32(1)
	v1095 = v348 & v1094
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v1095)
	F_cancel_before_shmem_exit(m, int32(407), v1094)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L9
	} else {
		goto L176
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	F_XLogBeginInsert(m)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L9
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	goto L77
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v375 = F_XLogInsert(m, int32(0), int32(64))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	F_RequestCheckpoint(m, v30)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L9
	} else {
		goto L79
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v565 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[5]))
	v566 = F_strlen(m, v565)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v574 = F_AllocateDir(m, int32(_a_F_do_pg_backup_start_7))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L9
	} else {
		goto L105
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v415 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[6]))
	v419 = F_LWLockAcquire(m, v415+int32(1152), int32(1))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	v422 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[7]))
	v423 = *(*int64)(unsafe.Add(mBase, uint32(v422)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l3)+1048)) = v423
	v425 = *(*int64)(unsafe.Add(mBase, uint32(v422)+40))
	*(*int64)(unsafe.Add(mBase, uint32(l3)+1032)) = v425
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v422)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+1040)) = v427
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422)+56)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v437 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[6]))
	F_LWLockRelease(m, v437+int32(1152))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L9
	} else {
		goto L81
	}
L81:
	;
	if v356 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	F_WALInsertLockAcquireExclusive(m)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L9
	} else {
		goto L97
	}
L83:
	;
	v445 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[1]))
	v448 = base.AtomicRmwXchg32(m, v445, int32(440), int32(1))
	if v448 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v456 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[1]))
	F_s_lock(m, v456+int32(440), int32(_a_F_do_pg_backup_start_2), int32(_a_F_do_pg_backup_start_8), int32(_a_F_do_pg_backup_start_4))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L9
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v465 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[1]))
	v466 = *(*int64)(unsafe.Add(mBase, uint32(v465)+432))
	v467 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v465)+440)), uint32(v467))
	if v429&int32(1) != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L86
L88:
	;
	v472 = *(*int64)(unsafe.Add(mBase, uint32(l3)+1032))
	if base.Ui64(v466) < base.Ui64(v472) {
		goto L82
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L9
	} else {
		goto L92
	}
L91:
	;
	goto L90
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	F_errcode(m, int32(325))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L9
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	F_errmsg(m, int32(_a_F_do_pg_backup_start_9), int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L9
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	F_errhint(m, int32(_a_F_do_pg_backup_start_10), int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L9
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	F_errfinish(m, int32(_a_F_do_pg_backup_start_2), int32(_a_F_do_pg_backup_start_11), int32(_a_F_do_pg_backup_start_4))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L9
	} else {
		goto L96
	}
L96:
	;
	goto L6
L97:
	;
	v534 = *(*int64)(unsafe.Add(mBase, uint32(l3)+1032))
	v536 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[1]))
	v537 = *(*int64)(unsafe.Add(mBase, uint32(v536)+168))
	if base.Ui64(v537) < base.Ui64(v534) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	goto L78
L99:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v536)+168)) = v534
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	F_WALInsertLockRelease(m)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L9
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	F_WALInsertLockRelease(m)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L9
	} else {
		goto L103
	}
L102:
	;
	goto L98
L103:
	;
	if v356 == int32(0) {
		goto L77
	} else {
		goto L104
	}
L104:
	;
	goto L98
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v583 = F_ReadDir(m, v574, int32(_a_F_do_pg_backup_start_7))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L9
	} else {
		goto L106
	}
L106:
	;
	if v583 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v587 = v26 + int32(96) + v566
	v591 = v33
	v598 = v40
	v599 = v41
	v601 = v583
	goto L110
L108:
	;
	v1024 = v33
	v1031 = v40
	v1032 = v41
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v1031
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v1024
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v1032
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	F_FreeDir(m, v574)
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L9
	} else {
		goto L174
	}
L110:
	;
	v613 = int32(*(*int8)(unsafe.Add(mBase, uint32(v601)+19)))
	if v613 < int32(49) {
		v999 = v598
		v1000 = v599
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v1024 = v1021
	v1031 = v999
	v1032 = v1000
	goto L109
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v999
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v1021 = F_ReadDir(m, v574, int32(_a_F_do_pg_backup_start_7))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L9
	} else {
		goto L172
	}
L113:
	;
	v616 = int32(*(*int8)(unsafe.Add(mBase, uint32(v601)+20)))
	if int32(57) < v616 {
		v999 = v598
		v1000 = v599
		goto L112
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	*(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[8])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v635 = v601 + int32(19)
	v640 = F_strtox_2(m, v635, v26+int32(92), int32(10), int64(4294967295))
	mBase = m.M
	goto L115
L115:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v26)+92))
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642))))
	if v643 != 0 {
		v999 = v598
		v1000 = v599
		goto L112
	} else {
		goto L116
	}
L116:
	;
	v645 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[8]))
	if base.B2i32(v645 == int32(28))|base.B2i32(v645 == int32(68)) != 0 {
		v999 = v598
		v1000 = v599
		goto L112
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v635
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = int32(_a_F_do_pg_backup_start_7)
	v661 = v26 + int32(1120)
	v666 = F_pg_snprintf(m, v661, int32(1034), int32(_a_F_do_pg_backup_start_12), v26+int32(48))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L9
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v676 = F_get_dirent_type(m, v661, v601, int32(0), int32(21))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L9
	} else {
		goto L123
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v26 + int32(1120)
	F_errmsg(m, v968, v26)
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L9
	} else {
		goto L170
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v917
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v939 = F_palloc(m, int32(24))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L9
	} else {
		goto L166
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v635
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = int32(_a_F_do_pg_backup_start_7)
	v894 = v26 + int32(96)
	v899 = F_pg_snprintf(m, v894, int32(1024), int32(_a_F_do_pg_backup_start_12), v26+int32(32))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L9
	} else {
		goto L164
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v691 = F_readlink(m, v26+int32(1120), v26+int32(96), int32(1024))
	mBase = m.M
	if v691 < int32(0) {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	switch v676 - int32(3) {
	case 0:
		goto L121
	case 1:
		goto L122
	default:
		v999 = v598
		v1000 = v599
		goto L112
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v702 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L9
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	if base.Ui32(int32(1024)) <= base.Ui32(v691) {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	if v702 == int32(0) {
		v999 = v598
		v1000 = v599
		goto L112
	} else {
		goto L128
	}
L128:
	;
	v968 = int32(_a_F_do_pg_backup_start_13)
	v969 = int32(_a_F_do_pg_backup_start_14)
	goto L119
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v718 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L9
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v724 = int32(0)
	v726 = v26 + int32(96)
	*(*uint8)(unsafe.Add(mBase, uint32(v726+v691))) = uint8(v724)
	if v691 <= v566 {
		v795 = v599
		v796 = v724
		goto L134
	} else {
		goto L135
	}
L132:
	;
	if v718 == int32(0) {
		v999 = v598
		v1000 = v599
		goto L112
	} else {
		goto L133
	}
L133:
	;
	v968 = int32(_a_F_do_pg_backup_start_15)
	v969 = int32(_a_F_do_pg_backup_start_16)
	goto L119
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v795
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	F_initStringInfo(m, v26+int32(76))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L9
	} else {
		goto L152
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v738 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[5]))
	if v566 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	if v783 != 0 {
		v795 = v599
		v796 = v724
		goto L134
	} else {
		goto L149
	}
L137:
	;
	v783 = int32(0)
	goto L136
L138:
	;
	goto L139
L139:
	;
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726))))
	if v744 != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v745 = v726
	v746 = v738
	v747 = v566
	v748 = v744
	goto L144
L141:
	;
	v771 = v738
	v775 = int32(0)
	goto L142
L142:
	;
	v776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v771))))
	v783 = v775 - v776
	goto L136
L143:
	;
	v771 = v766
	v775 = v768
	goto L142
L144:
	;
	v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746))))
	if base.B2i32(v748 != v750)|base.B2i32(v750 == int32(0)) != 0 {
		v766 = v746
		v768 = v748
		goto L143
	} else {
		goto L146
	}
L145:
	;
	v766 = v760
	v768 = int32(0)
	goto L143
L146:
	;
	v756 = v747 - int32(1)
	if v756 == int32(0) {
		v766 = v746
		v768 = v748
		goto L143
	} else {
		goto L147
	}
L147:
	;
	v759 = int32(1)
	v760 = v746 + v759
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v745)+1)))
	if v761 != 0 {
		v745 = v745 + v759
		v746 = v760
		v747 = v756
		v748 = v761
		goto L144
	} else {
		goto L148
	}
L148:
	;
	goto L145
L149:
	;
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587))))
	if v784 != int32(47) {
		v795 = v599
		v796 = v724
		goto L134
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v793 = F_pstrdup(m, v587+int32(1))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L9
	} else {
		goto L151
	}
L151:
	;
	v795 = v793
	v796 = v793
	goto L134
L152:
	;
	v820 = v26 + int32(96)
	goto L153
L153:
	;
	v832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v820))))
	switch v832 {
	case 0:
		goto L155
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12:
		v847 = v832
		goto L156
	case 10, 13:
		goto L157
	default:
		goto L158
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v795
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v26)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v867
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v635
	F_appendStringInfo(m, l4, int32(_a_F_do_pg_backup_start_17), v26+int32(16))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L9
	} else {
		goto L162
	}
L155:
	;
	goto L154
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v795
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	F_appendStringInfoChar(m, v26+int32(76), base.I32_extend8_s(v847))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L9
	} else {
		goto L161
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v795
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	F_appendStringInfoChar(m, v26+int32(76), int32(92))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L9
	} else {
		goto L160
	}
L158:
	;
	if v832 != int32(92) {
		v847 = v832
		goto L156
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v820))))
	v847 = v846
	goto L156
L161:
	;
	v820 = v820 + int32(1)
	goto L153
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v795
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v26)+76))
	F_pfree(m, v881)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L9
	} else {
		goto L163
	}
L163:
	;
	v917 = v598
	v918 = v795
	v926 = v796
	goto L120
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v907 = F_pstrdup(m, v894)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L9
	} else {
		goto L165
	}
L165:
	;
	v917 = v907
	v918 = v599
	v926 = v907
	goto L120
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v939))) = base.I32_wrap_i64(v640)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v917
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v950 = F_pstrdup(m, v26+int32(96))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L9
	} else {
		goto L167
	}
L167:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v939)+16)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v939)+8)) = v926
	*(*int32)(unsafe.Add(mBase, uint32(v939)+4)) = v950
	if l2 == int32(0) {
		v999 = v917
		v1000 = v918
		goto L112
	} else {
		goto L168
	}
L168:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v917
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v918
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v965 = F_lappend(m, v958, v939)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L9
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v965
	v999 = v917
	v1000 = v918
	goto L112
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	F_errfinish(m, int32(_a_F_do_pg_backup_start_2), v969, int32(_a_F_do_pg_backup_start_4))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L9
	} else {
		goto L171
	}
L171:
	;
	v999 = v598
	v1000 = v599
	goto L112
L172:
	;
	if v1021 != 0 {
		v591 = v1021
		v598 = v999
		v599 = v1000
		v601 = v1021
		goto L110
	} else {
		goto L173
	}
L173:
	;
	goto L111
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v1031
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v1024
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v1032
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	v1060 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l3)+1056)) = v1060
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v1031
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v1024
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v1032
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v356)
	F_cancel_before_shmem_exit(m, int32(407), int32(1))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L9
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[3])) = v345
	*(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[4])) = v346
	v1076 = int32(1)
	v1077 = v348 & v1076
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+1064)) = uint8(v1077)
	*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_backup_start[9])) = uint8(v1076)
	m.G0 = v26 + int32(2352)
	return
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v1095)
	F_do_pg_abort_backup(m, v26, int32(1))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L9
	} else {
		goto L177
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2332)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2328)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2336)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2340)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v26)+2344)) = v346
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)) = uint8(v1095)
	F_pg_re_throw(m)
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L9
	} else {
		goto L178
	}
L178:
	;
	goto L8
L179:
	;
	v1146 = int32(v1142)
	m.G0 = v26
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1146)+4))
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1146)))
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1149)))
	if v26+int32(72) == v1152 {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	m.ExcPending = 1
	goto L188
L181:
	;
	if v1156 != 0 {
		goto L185
	} else {
		goto L186
	}
L182:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1149)+4))
	v1156 = v1154
	goto L184
L183:
	;
	v1156 = int32(0)
	goto L184
L184:
	;
	goto L181
L185:
	;
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+2351)))
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v26)+2344))
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v26)+2340))
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v26)+2336))
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v26)+2332))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v26)+2328))
	v33 = v1162
	v38 = v1159
	v39 = v1158
	v40 = v1161
	v41 = v1160
	v42 = v1148
	v43 = v1156
	v44 = v1157
	goto L4
L186:
	;
	goto L187
L187:
	;
	F___wasm_longjmp(m, v1149, v1148)
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	return
L189:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_armor(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	switch v49 - int32(1) {
	case 0:
		v184 = v2
		v187 = v2
		v188 = v2
		goto L24
	default:
		goto L14
	case 2:
		goto L25
	}
L2:
	;
	return int32(0)
L3:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v19 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v25 == int32(18) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v36 = int32(1)
	if v19&v36 != 0 {
		v48 = int32(base.Ui32(v19)>>(uint(v36)%32)) - v36
		goto L1
	} else {
		goto L13
	}
L7:
	;
	v28 = int32(16)
	goto L9
L8:
	;
	v28 = int32(0)
	goto L9
L9:
	;
	if base.Ui32((v25-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v35 = int32(4)
	goto L12
L11:
	;
	v35 = v28
	goto L12
L12:
	;
	v48 = v35
	goto L1
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L2
	} else {
		goto L117
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L2
	} else {
		goto L113
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L2
	} else {
		goto L109
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L2
	} else {
		goto L105
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L2
	} else {
		goto L101
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L2
	} else {
		goto L97
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L2
	} else {
		goto L93
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L2
	} else {
		goto L89
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L2
	} else {
		goto L85
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L2
	} else {
		goto L81
	}
L24:
	;
	v191 = v12 + int32(12)
	F_initStringInfo(m, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L2
	} else {
		goto L67
	}
L25:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v53 = F_pg_detoast_datum(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v56 = F_pg_detoast_datum(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if int32(1) < v58 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v58 != v61 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	if v58 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v184 = int32(0)
	v187 = v2
	v188 = v2
	goto L24
L31:
	;
	goto L32
L32:
	;
	F_deconstruct_array_builtin(m, v53, int32(25), v12+int32(12), v12+int32(40), v12+int32(32))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	F_deconstruct_array_builtin(m, v56, int32(25), v12+int32(44), v12+int32(36), v12+int32(28))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	if v84 != v85 {
		goto L22
	} else {
		goto L35
	}
L35:
	;
	v89 = F_palloc(m, v84<<(uint(int32(2))%32))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v94 = F_palloc(m, v91<<(uint(int32(2))%32))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v96 <= int32(0) {
		v184 = v96
		v187 = v89
		v188 = v94
		goto L24
	} else {
		goto L38
	}
L38:
	;
	v102 = int32(0)
	goto L39
L39:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+v102))))
	if v111 == int32(1) {
		goto L21
	} else {
		goto L41
	}
L40:
	;
	v184 = v179
	v187 = v89
	v188 = v94
	goto L24
L41:
	;
	v115 = v102 << (uint(int32(2)) % 32)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v115+v116)))
	v119 = F_text_to_cstring(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v122 = v119
	goto L44
L43:
	;
	if base.B2i32(v124 == int32(0)) == int32(0) {
		goto L20
	} else {
		goto L47
	}
L44:
	;
	v124 = int32(*(*int8)(unsafe.Add(mBase, uint32(v122))))
	if int32(0) < v124 {
		v122 = v122 + int32(1)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	goto L43
L46:
	;
	goto L45
L47:
	;
	v134 = F_strstr(m, v119, int32(_a_F_pg_armor_0))
	mBase = m.M
	if v134 != 0 {
		goto L19
	} else {
		goto L48
	}
L48:
	;
	v135 = int32(10)
	v136 = F___strchrnul(m, v119, v135)
	mBase = m.M
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	if v138 == v135 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v142 != 0 {
		goto L18
	} else {
		goto L53
	}
L50:
	;
	v142 = v136
	goto L52
L51:
	;
	v142 = int32(0)
	goto L52
L52:
	;
	goto L49
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115+v89))) = v119
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+v102))))
	if v147 == int32(1) {
		goto L17
	} else {
		goto L54
	}
L54:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v150+v115)))
	v153 = F_text_to_cstring(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v156 = v153
	goto L57
L56:
	;
	if base.B2i32(v158 == int32(0)) == int32(0) {
		goto L16
	} else {
		goto L60
	}
L57:
	;
	v158 = int32(*(*int8)(unsafe.Add(mBase, uint32(v156))))
	if int32(0) < v158 {
		v156 = v156 + int32(1)
		goto L57
	} else {
		goto L59
	}
L58:
	;
	goto L56
L59:
	;
	goto L58
L60:
	;
	v167 = int32(10)
	v168 = F___strchrnul(m, v153, v167)
	mBase = m.M
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if v170 == v167 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if v174 != 0 {
		goto L15
	} else {
		goto L65
	}
L62:
	;
	v174 = v168
	goto L64
L63:
	;
	v174 = int32(0)
	goto L64
L64:
	;
	goto L61
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115+v94))) = v153
	v178 = v102 + int32(1)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v178 < v179 {
		v102 = v178
		goto L39
	} else {
		goto L66
	}
L66:
	;
	goto L40
L67:
	;
	v194 = int32(1)
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v196&v194 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v199 = v194
	goto L70
L69:
	;
	v199 = int32(4)
	goto L70
L70:
	;
	F_pgp_armor_encode(m, v15+v199, v48, v191, v184, v187, v188)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v206 = F_palloc(m, v203+int32(4))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v208<<(uint(int32(2))%32) + int32(16)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v215 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	base.MemoryCopy(m, v206+int32(4), v214, v215)
	goto L75
L74:
	;
	goto L75
L75:
	;
	F_pfree(m, v214)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v221 != v15 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_pfree(m, v15)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L2
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	m.G0 = v12 + int32(48)
	return v206
L80:
	;
	goto L79
L81:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	F_errmsg(m, int32(_a_F_pg_armor_1), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(775), int32(_a_F_pg_armor_3))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L2
	} else {
		goto L86
	}
L86:
	;
	F_errmsg(m, int32(_a_F_pg_armor_4), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L2
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(785), int32(_a_F_pg_armor_3))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L2
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
	F_errcode(m, int32(67108994))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L2
	} else {
		goto L90
	}
L90:
	;
	F_errmsg(m, int32(_a_F_pg_armor_5), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L2
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(798), int32(_a_F_pg_armor_3))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L2
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(_a_F_pg_armor_6), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(805), int32(_a_F_pg_armor_3))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L2
	} else {
		goto L98
	}
L98:
	;
	F_errmsg(m, int32(_a_F_pg_armor_7), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L2
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(809), int32(_a_F_pg_armor_3))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L2
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L2
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(_a_F_pg_armor_8), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L2
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(813), int32(_a_F_pg_armor_3))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L2
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L2
	} else {
		goto L106
	}
L106:
	;
	F_errmsg(m, int32(_a_F_pg_armor_9), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L2
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(820), int32(_a_F_pg_armor_3))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L2
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L2
	} else {
		goto L110
	}
L110:
	;
	F_errmsg(m, int32(_a_F_pg_armor_10), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(827), int32(_a_F_pg_armor_3))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	F_errmsg(m, int32(_a_F_pg_armor_11), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L2
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(831), int32(_a_F_pg_armor_3))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L2
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	v377 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v377
	F_errmsg_internal(m, int32(_a_F_pg_armor_12), v12)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(863), int32(_a_F_pg_armor_13))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_available_wal_summaries(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+2)) = uint8(v16)
	*(*uint16)(unsafe.Add(mBase, uint32(v8))) = uint16(v16)
	v21 = int64(0)
	v23 = F_GetWalSummaries(m, v16, v21, v21)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v8 + int32(16)
	return int32(0)
L4:
	;
	if v23 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v27 <= int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v34 = v2
	goto L7
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v34<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_pg_available_wal_summaries[0]))
	if v41 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L3
L9:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v44 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v39)+16)))
	v45 = F_Int64GetDatum(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v45
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	v49 = F_Int64GetDatum(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v49
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
	v53 = F_Int64GetDatum(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v53
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v59 = F_heap_form_tuple(m, v56, v8+int32(4), v8)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	F_tuplestore_puttuple(m, v61, v59)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v65 = v34 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v65 < v66 {
		v34 = v65
		goto L7
	} else {
		goto L18
	}
L18:
	;
	goto L8
}
func F_pg_b64_enc_len(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = int32(2)
	v5 = base.I32_div_s(l0+v2, int32(3))
	return v5 << (uint(v2) % 32)
}
func F_pg_b64_encode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	if int32(0) < l1 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if l3 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	if l3 < v58-l2+int32(4) {
		goto L1
	} else {
		goto L14
	}
L3:
	;
	v13 = l0
	v14 = int32(0)
	v17 = l2
	v18 = int32(2)
	goto L6
L4:
	;
	v69 = l2
	goto L5
L5:
	;
	return v69 - l2
L6:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	v24 = v20<<(uint(v18<<(uint(int32(3))%32))%32) | v14
	if int32(0) < v18 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v59 != int32(2) {
		goto L2
	} else {
		goto L13
	}
L8:
	;
	v57 = v24
	v58 = v17
	v59 = v18 - int32(1)
	goto L10
L9:
	;
	if l3 < v17-l2+int32(4) {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v61 = v13 + int32(1)
	if base.Ui32(v61) < base.Ui32(l0+l1) {
		v13 = v61
		v14 = v57
		v17 = v58
		v18 = v59
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v33 = int32(63)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24&v33)+uint32(_c_F_pg_b64_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+3)) = uint8(v35)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v24)>>(uint(int32(18))%32)))+uint32(_c_F_pg_b64_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v39)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v24)>>(uint(int32(6))%32))&v33)+uint32(_c_F_pg_b64_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)) = uint8(v45)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v24)>>(uint(int32(12))%32))&v33)+uint32(_c_F_pg_b64_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)) = uint8(v51)
	v57 = int32(0)
	v58 = v17 + int32(4)
	v59 = int32(2)
	goto L10
L12:
	;
	goto L7
L13:
	;
	v69 = v58
	goto L5
L14:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v57)>>(uint(int32(18))%32)))+uint32(_c_F_pg_b64_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v80)
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v57)>>(uint(int32(12))%32))&int32(63))+uint32(_c_F_pg_b64_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)) = uint8(v86)
	if v59 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v57)>>(uint(int32(6))%32))&int32(63))+uint32(_c_F_pg_b64_encode[0]))))
	v96 = v95
	goto L17
L16:
	;
	v96 = int32(61)
	goto L17
L17:
	;
	v97 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+3)) = uint8(v97)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+2)) = uint8(v96)
	return v58 + int32(4) - l2
L18:
	;
	base.MemoryFill(m, l2, int32(0), l3)
	goto L20
L19:
	;
	goto L20
L20:
	;
	return int32(-1)
}
func F_pg_backup_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_backup_start[0])))
		v13 = F_text_to_cstring(m, v6)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			if v12 != int32(1) {
				v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[1]))
				if v18 == int32(0) {
					v23 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[2]))
					v28 = F_AllocSetContextCreateInternal(m, v23, int32(_a_F_pg_backup_start_0), int32(0), int32(1024), int32(_a_F_pg_backup_start_1))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[1])) = v28
						v41 = v28
						v42 = int32(_a_F_pg_backup_start_2)
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[3]))
						*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[3])) = v41
						v48 = F_palloc0(m, int32(1112))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4])) = v48
							v51 = F_makeStringInfo(m)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[3])) = v43
								*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[5])) = v51
								v58 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_backup_start[6])))
								if v58 == int32(0) {
									F_before_shmem_exit(m, int32(407), int32(0))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										v66 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_pg_backup_start[6])) = uint8(v66)
										v68 = int32(0)
										v72 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4]))
										v74 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[5]))
										F_do_pg_backup_start(m, v13, base.B2i32(v10 != v68), v68, v72, v74)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4]))
											v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+1032))
											v80 = F_Int64GetDatum(m, v79)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												return v80
											}
										}
									}
								} else {
									v68 = int32(0)
									v72 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4]))
									v74 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[5]))
									F_do_pg_backup_start(m, v13, base.B2i32(v10 != v68), v68, v72, v74)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4]))
										v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+1032))
										v80 = F_Int64GetDatum(m, v79)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											return v80
										}
									}
								}
							}
						}
					}
				} else {
					v32 = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[5])) = v32
					*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4])) = v32
					F_MemoryContextReset(m, v18)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[1]))
						v41 = v40
						v42 = int32(_a_F_pg_backup_start_2)
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[3]))
						*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[3])) = v41
						v48 = F_palloc0(m, int32(1112))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4])) = v48
							v51 = F_makeStringInfo(m)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[3])) = v43
								*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[5])) = v51
								v58 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_backup_start[6])))
								if v58 == int32(0) {
									F_before_shmem_exit(m, int32(407), int32(0))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										v66 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_pg_backup_start[6])) = uint8(v66)
										v68 = int32(0)
										v72 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4]))
										v74 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[5]))
										F_do_pg_backup_start(m, v13, base.B2i32(v10 != v68), v68, v72, v74)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4]))
											v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+1032))
											v80 = F_Int64GetDatum(m, v79)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												return v80
											}
										}
									}
								} else {
									v68 = int32(0)
									v72 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4]))
									v74 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[5]))
									F_do_pg_backup_start(m, v13, base.B2i32(v10 != v68), v68, v72, v74)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4]))
										v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+1032))
										v80 = F_Int64GetDatum(m, v79)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											return v80
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_pg_backup_start_3), int32(0))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_backup_start_4), int32(70), int32(_a_F_pg_backup_start_5))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
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
func F_pg_base64_decode_2(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v155 int32
	_ = v155
	v4 = int32(0)
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v155
L2:
	;
	v11 = l0 + l1
	v12 = l0
	v15 = l2
	v16 = v4
	v17 = v4
	v19 = v4
	goto L5
L3:
	;
	v139 = l2
	goto L4
L4:
	;
	v155 = v139 - l2
	goto L1
L5:
	;
	v23 = v12
	goto L13
L6:
	;
	if v133 != 0 {
		v155 = int32(-101)
		goto L1
	} else {
		goto L40
	}
L7:
	;
	goto L6
L8:
	;
	if base.Ui32(v33) < base.Ui32(v11) {
		v12 = v33
		v15 = v123
		v16 = v124
		v17 = v125
		v19 = v126
		goto L5
	} else {
		goto L39
	}
L9:
	;
	v123 = v116
	v124 = v117
	v125 = v118
	v126 = int32(0)
	goto L8
L10:
	;
	v116 = v113
	v117 = int32(0)
	v118 = v110
	goto L9
L11:
	;
	v85 = v81 + v16<<(uint(int32(6))%32)
	v87 = v19 + int32(1)
	if v87 != int32(4) {
		goto L30
	} else {
		goto L31
	}
L12:
	;
	v81 = int32(62)
	goto L11
L13:
	;
	v33 = v23 + int32(1)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v36 = v34 - int32(65)
	if base.Ui32(v36&int32(255)) <= base.Ui32(int32(25)) {
		v81 = v36
		goto L11
	} else {
		goto L15
	}
L14:
	;
	if v17 != 0 {
		goto L25
	} else {
		goto L26
	}
L15:
	;
	if base.Ui32((v34-int32(97))&int32(255)) <= base.Ui32(int32(25)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v81 = v34 - int32(71)
	goto L11
L17:
	;
	goto L18
L18:
	;
	if base.Ui32((v34-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v81 = (v34 + int32(4)) & int32(255)
	goto L11
L20:
	;
	goto L21
L21:
	;
	v60 = int32(-101)
	switch v34 - int32(9) {
	case 0, 1, 4, 23:
		goto L23
	default:
		v155 = v60
		goto L1
	case 34:
		goto L12
	case 38:
		v81 = int32(63)
		goto L11
	case 52:
		goto L22
	}
L22:
	;
	goto L14
L23:
	;
	if base.Ui32(v33) < base.Ui32(v11) {
		v23 = v33
		goto L13
	} else {
		goto L24
	}
L24:
	;
	v130 = v15
	v133 = v19
	goto L7
L25:
	;
	v81 = int32(0)
	goto L11
L26:
	;
	goto L27
L27:
	;
	switch v19 - int32(2) {
	case 0:
		goto L29
	case 1:
		goto L28
	default:
		v155 = v60
		goto L1
	}
L28:
	;
	v71 = int32(2)
	v73 = int32(base.Ui32(v16) >> (uint(v71) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v73)
	v76 = int32(base.Ui32(v16) >> (uint(int32(10)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v76)
	v110 = v71
	v113 = v15 + v71
	goto L10
L29:
	;
	v123 = v15
	v124 = v16 << (uint(int32(6)) % 32)
	v125 = int32(1)
	v126 = int32(3)
	goto L8
L30:
	;
	v123 = v15
	v124 = v85
	v125 = v17
	v126 = v87
	goto L8
L31:
	;
	goto L32
L32:
	;
	v91 = int32(base.Ui32(v85) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v91)
	v93 = int32(0)
	if v17 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = int32(1)
	v116 = v15 + v96
	v117 = v93
	v118 = v96
	goto L9
L34:
	;
	goto L35
L35:
	;
	v100 = int32(base.Ui32(v85) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v100)
	if v17 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v116 = v15 + int32(2)
	v117 = v93
	v118 = v17
	goto L9
L37:
	;
	goto L38
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+2)) = uint8(v85)
	v110 = int32(0)
	v113 = v15 + int32(3)
	goto L10
L39:
	;
	v130 = v123
	v133 = v126
	goto L7
L40:
	;
	v139 = v130
	goto L4
}
func F_pg_class_aclcheck_ext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_pg_class_aclmask_ext(m, l0, l1, l2, int32(1), l3)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 == int64(0))
	}
}
func F_pg_class_aclmask_ext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int64
	_ = v143
	var v145 int64
	_ = v145
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v161 int64
	_ = v161
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v161
L2:
	;
	return int64(0)
L3:
	;
	if v16 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l4 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+22)))
	v43 = v41 + v42
	if l2&int64(285) == int64(0) {
		v68 = l2
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v22)
	v161 = int64(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
	F_errmsg(m, int32(_a_F_pg_class_aclmask_ext_0), v13)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_pg_class_aclmask_ext_1), int32(3310), int32(_a_F_pg_class_aclmask_ext_2))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	v69 = F_superuser_arg(m, l1)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L26
	}
L15:
	;
	v49 = int32(1)
	if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_class_aclmask_ext_6)) {
		v57 = v49
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v57 == int32(0) {
		v68 = l2
		goto L14
	} else {
		goto L20
	}
L17:
	;
	goto L16
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)+68))
	if v52 == int32(99) {
		v57 = v49
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v55 = F_isTempToastNamespace(m, v52)
	mBase = m.M
	v57 = v55
	goto L17
L20:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+119)))
	if v60 == int32(118) {
		v68 = l2
		goto L14
	} else {
		goto L21
	}
L21:
	;
	v65 = F_superuser_arg(m, l1)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	if v65 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v67 = l2
	goto L25
L24:
	;
	v67 = l2 & int64(-286)
	goto L25
L25:
	;
	v68 = v67
	goto L14
L26:
	;
	if v69 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_ReleaseCatCache(m, v16)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	v78 = F_SysCacheGetAttr(m, int32(57), v16, int32(32), v13+int32(15))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L2
	} else {
		goto L31
	}
L30:
	;
	v161 = v68
	goto L1
L31:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	if v80 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v98 = F_aclmask(m, v97, l1, v73, v68, l3)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L42
	}
L33:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+119)))
	if v83 == int32(83) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v94 = F_pg_detoast_datum(m, v78)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L41
	}
L36:
	;
	v88 = F_acldefault(m, int32(37), v73)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v92 = F_acldefault(m, int32(41), v73)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L2
	} else {
		goto L40
	}
L39:
	;
	v96 = int32(0)
	v97 = v88
	goto L32
L40:
	;
	v96 = int32(0)
	v97 = v92
	goto L32
L41:
	;
	v96 = v78
	v97 = v94
	goto L32
L42:
	;
	v100 = int32(0)
	if base.B2i32(v97 == v100)|base.B2i32(v96 == v97) == v100 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_pfree(m, v97)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L2
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_ReleaseCatCache(m, v16)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	v110 = int64(2)
	v112 = int64(0)
	if base.B2i32(v68&v110 == v112)|base.B2i32(v98&v110 != v112) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v124 = F_has_privs_of_role(m, l1, int32(_a_F_pg_class_aclmask_ext_3))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L51
	}
L49:
	;
	v127 = v98
	goto L50
L50:
	;
	v128 = int64(13)
	v129 = v68 & v128
	v130 = int64(0)
	if base.B2i32(v129 == v130)|base.B2i32(v127&v128 != v130) == int32(0) {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	if v124 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v126 = v98 | int64(2)
	goto L54
L53:
	;
	v126 = v98
	goto L54
L54:
	;
	v127 = v126
	goto L50
L55:
	;
	v141 = F_has_privs_of_role(m, l1, int32(_a_F_pg_class_aclmask_ext_4))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L2
	} else {
		goto L58
	}
L56:
	;
	v145 = v127
	goto L57
L57:
	;
	if v68&int64(16384) == int64(0) {
		goto L62
	} else {
		goto L63
	}
L58:
	;
	if v141 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v143 = v129
	goto L61
L60:
	;
	v143 = int64(0)
	goto L61
L61:
	;
	v145 = v143 | v127
	goto L57
L62:
	;
	v161 = v145
	goto L1
L63:
	;
	goto L64
L64:
	;
	if v145&int64(16384) != int64(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v161 = v145
	goto L1
L66:
	;
	goto L67
L67:
	;
	v157 = F_has_privs_of_role(m, l1, int32(_a_F_pg_class_aclmask_ext_5))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	if v157 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v159 = v145 | int64(16384)
	goto L71
L70:
	;
	v159 = v145
	goto L71
L71:
	;
	v161 = v159
	goto L1
}
func F_pg_collation_actual_version(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
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
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v10 == int32(100) {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_pg_collation_actual_version[0]))
		v16 = F_SearchSysCache1(m, int32(21), v15)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v16 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, _c_F_pg_collation_actual_version[0]))
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v79
						F_errmsg(m, int32(_a_F_pg_collation_actual_version_0), v8)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_collation_actual_version_1), int32(524), int32(_a_F_pg_collation_actual_version_2))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v26)+76)))
				if v28 == int32(99) {
					v31 = int32(13)
				} else {
					v31 = int32(15)
				}
				v32 = F_SysCacheGetAttrNotNull(m, int32(21), v16, v31)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v51 = v28
					v52 = v16
					v53 = v32
					v54 = F_text_to_cstring(m, v53)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_ReleaseCatCache(m, v52)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v59 = F_get_collation_actual_version(m, base.I32_extend8_s(v51), v54)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								if v59 != 0 {
									v61 = F_cstring_to_text(m, v59)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										v66 = v61
										m.G0 = v8 + int32(32)
										return v66
									}
								} else {
									v63 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v63)
									v66 = int32(0)
									m.G0 = v8 + int32(32)
									return v66
								}
							}
						}
					}
				}
			}
		}
	} else {
		v35 = F_SearchSysCache1(m, int32(16), v10)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			if v35 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
						F_errmsg(m, int32(_a_F_pg_collation_actual_version_3), v8+int32(16))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_collation_actual_version_1), int32(550), int32(_a_F_pg_collation_actual_version_2))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+22)))
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v43)+76)))
				if v45 == int32(99) {
					v48 = int32(8)
				} else {
					v48 = int32(10)
				}
				v49 = F_SysCacheGetAttrNotNull(m, int32(16), v35, v48)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					v51 = v45
					v52 = v35
					v53 = v49
					v54 = F_text_to_cstring(m, v53)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_ReleaseCatCache(m, v52)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v59 = F_get_collation_actual_version(m, base.I32_extend8_s(v51), v54)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								if v59 != 0 {
									v61 = F_cstring_to_text(m, v59)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										v66 = v61
										m.G0 = v8 + int32(32)
										return v66
									}
								} else {
									v63 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v63)
									v66 = int32(0)
									m.G0 = v8 + int32(32)
									return v66
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_collation_is_visible(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_SearchSysCache1(m, int32(16), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v8 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v74 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v74)
	return int32(0)
L6:
	;
	v16 = v12 + v13
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
	if v17 != int32(11) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	F_ReleaseCatCache(m, v8)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L26
	}
L8:
	;
	v20 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_pg_collation_is_visible[0]))
	if v22 == v20 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	v67 = F_CollationGetCollid(m, v16+int32(4))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L25
	}
L11:
	;
	if v61 == int32(0) {
		v70 = v20
		goto L7
	} else {
		goto L24
	}
L12:
	;
	v61 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v29 <= int32(0) {
		v55 = v20
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v61 = v55
	goto L11
L16:
	;
	v32 = int32(0)
	if v32 < v29 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v35 = v29
	goto L19
L18:
	;
	v35 = v32
	goto L19
L19:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v38 = int32(0)
	goto L20
L20:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v36+v38<<(uint(int32(2))%32))))
	v47 = base.B2i32(v46 == v17)
	if v46 == v17 {
		v55 = v47
		goto L15
	} else {
		goto L22
	}
L21:
	;
	v55 = v47
	goto L15
L22:
	;
	v49 = v38 + int32(1)
	if v49 != v35 {
		v38 = v49
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L10
L25:
	;
	v70 = base.B2i32(v67 == v7)
	goto L7
L26:
	;
	return v70
}
func F_pg_column_compression(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v11 == int32(0) {
		v15 = F_get_fn_expr_argtype(m, v10, int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = F_get_typlen(m, v15)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v115 = m.ExcPending
					if v115 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
						F_errmsg_internal(m, int32(_a_F_pg_column_compression_0), v8)
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_column_compression_1), int32(_a_F_pg_column_compression_2), int32(_a_F_pg_column_compression_3))
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
					v26 = F_MemoryContextAlloc(m, v24, int32(4))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v26
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = v19
						v35 = v19
						if v35 != int32(-1) {
							v38 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v38)
							v104 = int32(0)
							m.G0 = v8 + int32(32)
							return v104
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
							if v43 == int32(1) {
								v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
								if v47 != int32(18) {
									v70 = int32(2)
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+6))
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+2))
									if base.Ui32(v54-int32(4)) <= base.Ui32(v51&int32(1073741823)) {
										v60 = int32(2)
									} else {
										v60 = int32(base.Ui32(v51) >> (uint(int32(30)) % 32))
									}
									v70 = v60
								}
							} else {
								v61 = int32(2)
								if v43&int32(3) != v61 {
									v70 = v61
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
									v70 = int32(base.Ui32(v66) >> (uint(int32(30)) % 32))
								}
							}
							switch v70 {
							case 0:
								v90 = int32(_a_F_pg_column_compression_4)
								v91 = F_strlen(m, v90)
								mBase = m.M
								v93 = v91 + int32(4)
								v94 = F_palloc(m, v93)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v94))) = v93 << (uint(int32(2)) % 32)
									if v91 == int32(0) {
										v104 = v94
									} else {
										base.MemoryCopy(m, v94+int32(4), v90, v91)
										v104 = v94
									}
									m.G0 = v8 + int32(32)
									return v104
								}
							case 1:
								v90 = int32(_a_F_pg_column_compression_5)
								v91 = F_strlen(m, v90)
								mBase = m.M
								v93 = v91 + int32(4)
								v94 = F_palloc(m, v93)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v94))) = v93 << (uint(int32(2)) % 32)
									if v91 == int32(0) {
										v104 = v94
									} else {
										base.MemoryCopy(m, v94+int32(4), v90, v91)
										v104 = v94
									}
									m.G0 = v8 + int32(32)
									return v104
								}
							case 2:
								v71 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v71)
								v104 = int32(0)
								m.G0 = v8 + int32(32)
								return v104
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
									F_errmsg_internal(m, int32(_a_F_pg_column_compression_6), v8+int32(16))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_pg_column_compression_1), int32(_a_F_pg_column_compression_7), int32(_a_F_pg_column_compression_3))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
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
	} else {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		v35 = v33
		if v35 != int32(-1) {
			v38 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v38)
			v104 = int32(0)
			m.G0 = v8 + int32(32)
			return v104
		} else {
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
			if v43 == int32(1) {
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
				if v47 != int32(18) {
					v70 = int32(2)
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+6))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+2))
					if base.Ui32(v54-int32(4)) <= base.Ui32(v51&int32(1073741823)) {
						v60 = int32(2)
					} else {
						v60 = int32(base.Ui32(v51) >> (uint(int32(30)) % 32))
					}
					v70 = v60
				}
			} else {
				v61 = int32(2)
				if v43&int32(3) != v61 {
					v70 = v61
				} else {
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
					v70 = int32(base.Ui32(v66) >> (uint(int32(30)) % 32))
				}
			}
			switch v70 {
			case 0:
				v90 = int32(_a_F_pg_column_compression_4)
				v91 = F_strlen(m, v90)
				mBase = m.M
				v93 = v91 + int32(4)
				v94 = F_palloc(m, v93)
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v94))) = v93 << (uint(int32(2)) % 32)
					if v91 == int32(0) {
						v104 = v94
					} else {
						base.MemoryCopy(m, v94+int32(4), v90, v91)
						v104 = v94
					}
					m.G0 = v8 + int32(32)
					return v104
				}
			case 1:
				v90 = int32(_a_F_pg_column_compression_5)
				v91 = F_strlen(m, v90)
				mBase = m.M
				v93 = v91 + int32(4)
				v94 = F_palloc(m, v93)
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v94))) = v93 << (uint(int32(2)) % 32)
					if v91 == int32(0) {
						v104 = v94
					} else {
						base.MemoryCopy(m, v94+int32(4), v90, v91)
						v104 = v94
					}
					m.G0 = v8 + int32(32)
					return v104
				}
			case 2:
				v71 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v71)
				v104 = int32(0)
				m.G0 = v8 + int32(32)
				return v104
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
					F_errmsg_internal(m, int32(_a_F_pg_column_compression_6), v8+int32(16))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_column_compression_1), int32(_a_F_pg_column_compression_7), int32(_a_F_pg_column_compression_3))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
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
func F_pg_copy_physical_replication_slot_a(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_copy_replication_slot(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_pg_cryptohash_update(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	if l0 == int32(0) {
		return int32(-1)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = m.Env.Pgmem_hash_update(m, v8, l1, l2)
		mBase = m.M
		return int32(0)
	}
}
func F_pg_ddl_command_recv(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_pg_ddl_command_recv_0), int32(359), int32(_a_F_pg_ddl_command_recv_1), int32(_a_F_pg_ddl_command_recv_2), int32(_a_F_pg_ddl_command_recv_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_pg_decrypt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
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
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
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
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
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
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L2
	} else {
		goto L94
	}
L2:
	;
	return int32(0)
L3:
	;
	v19 = int32(1)
	v20 = v15 + v19
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v25 = v23 & v19
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = v20
	goto L6
L5:
	;
	v26 = v15 + int32(4)
	goto L6
L6:
	;
	if v23 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v55 = F_downcase_truncate_identifier(m, v26, v53, int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L18
	}
L8:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v32 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v43 = int32(1)
	if v25 != 0 {
		v53 = int32(base.Ui32(v23)>>(uint(v43)%32)) - v43
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v35 = int32(16)
	goto L13
L12:
	;
	v35 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v42 = int32(4)
	goto L16
L15:
	;
	v42 = v35
	goto L16
L16:
	;
	v53 = v42
	goto L7
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v59 = F_px_find_combo(m, v55, v12+int32(28))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	if v59 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_pfree(m, v55)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L2
	} else {
		goto L76
	}
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v67 = F_pg_detoast_datum_packed(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v70 = F_pg_detoast_datum_packed(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v72 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v102 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L27:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	if v78 == int32(18) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v89 = int32(1)
	if v72&v89 != 0 {
		v101 = int32(base.Ui32(v72)>>(uint(v89)%32)) - v89
		goto L26
	} else {
		goto L36
	}
L30:
	;
	v81 = int32(16)
	goto L32
L31:
	;
	v81 = int32(0)
	goto L32
L32:
	;
	if base.Ui32((v78-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v88 = int32(4)
	goto L35
L34:
	;
	v88 = v81
	goto L35
L35:
	;
	v101 = v88
	goto L26
L36:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v101 = int32(base.Ui32(v95)>>(uint(int32(2))%32)) - int32(4)
	goto L26
L37:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	v133 = m.T0[v132].(func(*base.Module, int32, int32) int32)(m, v65, v101)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L2
	} else {
		goto L48
	}
L38:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v108 == int32(18) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v119 = int32(1)
	if v102&v119 != 0 {
		v131 = int32(base.Ui32(v102)>>(uint(v119)%32)) - v119
		goto L37
	} else {
		goto L47
	}
L41:
	;
	v111 = int32(16)
	goto L43
L42:
	;
	v111 = int32(0)
	goto L43
L43:
	;
	if base.Ui32((v108-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v118 = int32(4)
	goto L46
L45:
	;
	v118 = v111
	goto L46
L46:
	;
	v131 = v118
	goto L37
L47:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v131 = int32(base.Ui32(v125)>>(uint(int32(2))%32)) - int32(4)
	goto L37
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v133
	v138 = F_palloc(m, v133+int32(4))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	v140 = int32(1)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v142&v140 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v145 = v140
	goto L52
L51:
	;
	v145 = int32(4)
	goto L52
L52:
	;
	v147 = int32(0)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v150 = m.T0[v149].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v65, v70+v145, v131, v147, v147)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	if v150 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	m.T0[v152].(func(*base.Module, int32))(m, v65)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L2
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v155 = int32(1)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v157&v155 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v243 = v150
	goto L1
L58:
	;
	v160 = v155
	goto L60
L59:
	;
	v160 = int32(4)
	goto L60
L60:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v167 = m.T0[v166].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v65, v67+v160, v101, v138+int32(4), v12+int32(28))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	m.T0[v169].(func(*base.Module, int32))(m, v65)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	if v167 != 0 {
		v243 = v167
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v172<<(uint(int32(2))%32) + int32(16)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v178 != v67 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_pfree(m, v67)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v182 != v70 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L66
L68:
	;
	F_pfree(m, v70)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L2
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v186 != v15 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L70
L72:
	;
	F_pfree(m, v15)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L2
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	m.G0 = v12 + int32(32)
	return v138
L75:
	;
	goto L74
L76:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	if v59 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v55
	F_errmsg(m, int32(_a_F_pg_decrypt_0), v12+int32(16))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L2
	} else {
		goto L92
	}
L79:
	;
	v230 = int32(_a_F_pg_decrypt_1)
	goto L78
L80:
	;
	goto L81
L81:
	;
	v209 = int32(_a_F_pg_decrypt_2)
	goto L83
L82:
	;
	v230 = v224
	goto L78
L83:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)+8))
	if v59 != v212 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v209)+12))
	v224 = v222
	goto L82
L85:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
	if v214 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	goto L84
L88:
	;
	v230 = int32(_a_F_pg_decrypt_3)
	goto L78
L89:
	;
	goto L90
L90:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v209)+16))
	if v59 != v218 {
		v209 = v209 + int32(16)
		goto L83
	} else {
		goto L91
	}
L91:
	;
	v224 = v214
	goto L82
L92:
	;
	F_errfinish(m, int32(_a_F_pg_decrypt_4), int32(513), int32(_a_F_pg_decrypt_5))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	F_errcode(m, int32(579))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	if v243 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v280
	F_errmsg(m, int32(_a_F_pg_decrypt_6), v12)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L2
	} else {
		goto L110
	}
L97:
	;
	v280 = int32(_a_F_pg_decrypt_1)
	goto L96
L98:
	;
	goto L99
L99:
	;
	v259 = int32(_a_F_pg_decrypt_2)
	goto L101
L100:
	;
	v280 = v274
	goto L96
L101:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v259)+8))
	if v243 != v262 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v274 = v272
	goto L100
L103:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v259)+20))
	if v264 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	goto L102
L106:
	;
	v280 = int32(_a_F_pg_decrypt_3)
	goto L96
L107:
	;
	goto L108
L108:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v259)+16))
	if v243 != v268 {
		v259 = v259 + int32(16)
		goto L101
	} else {
		goto L109
	}
L109:
	;
	v274 = v264
	goto L100
L110:
	;
	F_errfinish(m, int32(_a_F_pg_decrypt_4), int32(334), int32(_a_F_pg_decrypt_7))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_decrypt_iv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
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
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L2
	} else {
		goto L113
	}
L2:
	;
	return int32(0)
L3:
	;
	v21 = int32(1)
	v22 = v17 + v21
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v27 = v25 & v21
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = v22
	goto L6
L5:
	;
	v28 = v17 + int32(4)
	goto L6
L6:
	;
	if v25 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v57 = F_downcase_truncate_identifier(m, v28, v55, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L18
	}
L8:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v34 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v45 = int32(1)
	if v27 != 0 {
		v55 = int32(base.Ui32(v25)>>(uint(v45)%32)) - v45
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v37 = int32(16)
	goto L13
L12:
	;
	v37 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v34-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v44 = int32(4)
	goto L16
L15:
	;
	v44 = v37
	goto L16
L16:
	;
	v55 = v44
	goto L7
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v61 = F_px_find_combo(m, v57, v14+int32(28))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	if v61 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_pfree(m, v57)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L2
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L2
	} else {
		goto L95
	}
L23:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v69 = F_pg_detoast_datum_packed(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v72 = F_pg_detoast_datum_packed(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v75 = F_pg_detoast_datum_packed(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v77 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v107 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+1)))
	if v83 == int32(18) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v94 = int32(1)
	if v77&v94 != 0 {
		v106 = int32(base.Ui32(v77)>>(uint(v94)%32)) - v94
		goto L27
	} else {
		goto L37
	}
L31:
	;
	v86 = int32(16)
	goto L33
L32:
	;
	v86 = int32(0)
	goto L33
L33:
	;
	if base.Ui32((v83-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v93 = int32(4)
	goto L36
L35:
	;
	v93 = v86
	goto L36
L36:
	;
	v106 = v93
	goto L27
L37:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v106 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
	goto L27
L38:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v137 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L39:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	if v113 == int32(18) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v124 = int32(1)
	if v107&v124 != 0 {
		v136 = int32(base.Ui32(v107)>>(uint(v124)%32)) - v124
		goto L38
	} else {
		goto L48
	}
L42:
	;
	v116 = int32(16)
	goto L44
L43:
	;
	v116 = int32(0)
	goto L44
L44:
	;
	if base.Ui32((v113-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v123 = int32(4)
	goto L47
L46:
	;
	v123 = v116
	goto L47
L47:
	;
	v136 = v123
	goto L38
L48:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v136 = int32(base.Ui32(v130)>>(uint(int32(2))%32)) - int32(4)
	goto L38
L49:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v168 = m.T0[v167].(func(*base.Module, int32, int32) int32)(m, v67, v106)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L2
	} else {
		goto L60
	}
L50:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
	if v143 == int32(18) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v154 = int32(1)
	if v137&v154 != 0 {
		v166 = int32(base.Ui32(v137)>>(uint(v154)%32)) - v154
		goto L49
	} else {
		goto L59
	}
L53:
	;
	v146 = int32(16)
	goto L55
L54:
	;
	v146 = int32(0)
	goto L55
L55:
	;
	if base.Ui32((v143-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v153 = int32(4)
	goto L58
L57:
	;
	v153 = v146
	goto L58
L58:
	;
	v166 = v153
	goto L49
L59:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v166 = int32(base.Ui32(v160)>>(uint(int32(2))%32)) - int32(4)
	goto L49
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v168
	v173 = F_palloc(m, v168+int32(4))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v175 = int32(1)
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v177&v175 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v180 = v175
	goto L64
L63:
	;
	v180 = int32(4)
	goto L64
L64:
	;
	v182 = int32(1)
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v184&v182 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v187 = v182
	goto L67
L66:
	;
	v187 = int32(4)
	goto L67
L67:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v190 = m.T0[v189].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v67, v72+v180, v136, v75+v187, v166)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	if v190 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	m.T0[v192].(func(*base.Module, int32))(m, v67)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L2
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v195 = int32(1)
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v197&v195 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v287 = v190
	goto L1
L73:
	;
	v200 = v195
	goto L75
L74:
	;
	v200 = int32(4)
	goto L75
L75:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	v207 = m.T0[v206].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v67, v69+v200, v106, v173+int32(4), v14+int32(28))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	m.T0[v209].(func(*base.Module, int32))(m, v67)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	if v207 != 0 {
		v287 = v207
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = v212<<(uint(int32(2))%32) + int32(16)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v218 != v69 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	F_pfree(m, v69)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L2
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v222 != v72 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L81
L83:
	;
	F_pfree(m, v72)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L2
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v226 != v75 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	F_pfree(m, v75)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L2
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v230 != v17 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	goto L89
L91:
	;
	F_pfree(m, v17)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L2
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	m.G0 = v14 + int32(32)
	return v173
L94:
	;
	goto L93
L95:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	if v61 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v274
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v57
	F_errmsg(m, int32(_a_F_pg_decrypt_iv_0), v14+int32(16))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L2
	} else {
		goto L111
	}
L98:
	;
	v274 = int32(_a_F_pg_decrypt_iv_1)
	goto L97
L99:
	;
	goto L100
L100:
	;
	v253 = int32(_a_F_pg_decrypt_iv_2)
	goto L102
L101:
	;
	v274 = v268
	goto L97
L102:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v253)+8))
	if v61 != v256 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v253)+12))
	v268 = v266
	goto L101
L104:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v253)+20))
	if v258 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	goto L103
L107:
	;
	v274 = int32(_a_F_pg_decrypt_iv_3)
	goto L97
L108:
	;
	goto L109
L109:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v253)+16))
	if v61 != v262 {
		v253 = v253 + int32(16)
		goto L102
	} else {
		goto L110
	}
L110:
	;
	v268 = v258
	goto L101
L111:
	;
	F_errfinish(m, int32(_a_F_pg_decrypt_iv_4), int32(513), int32(_a_F_pg_decrypt_iv_5))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errcode(m, int32(579))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	if v287 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v324
	F_errmsg(m, int32(_a_F_pg_decrypt_iv_6), v14)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L2
	} else {
		goto L129
	}
L116:
	;
	v324 = int32(_a_F_pg_decrypt_iv_1)
	goto L115
L117:
	;
	goto L118
L118:
	;
	v303 = int32(_a_F_pg_decrypt_iv_2)
	goto L120
L119:
	;
	v324 = v318
	goto L115
L120:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v303)+8))
	if v287 != v306 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v303)+12))
	v318 = v316
	goto L119
L122:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v303)+20))
	if v308 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	goto L121
L125:
	;
	v324 = int32(_a_F_pg_decrypt_iv_3)
	goto L115
L126:
	;
	goto L127
L127:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v303)+16))
	if v287 != v312 {
		v303 = v303 + int32(16)
		goto L120
	} else {
		goto L128
	}
L128:
	;
	v318 = v308
	goto L119
L129:
	;
	F_errfinish(m, int32(_a_F_pg_decrypt_iv_4), int32(441), int32(_a_F_pg_decrypt_iv_7))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L2
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_eucjp_increment(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	v3 = int32(0)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v6 - int32(142) {
	case 0:
		if l1 != int32(2) {
			return int32(0)
		} else {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			if base.Ui32(int32(223)) <= base.Ui32(v13) {
				v16 = int32(_a_F_pg_eucjp_increment_0)
				*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v16)
				return int32(1)
			} else {
				if base.Ui32(v13) <= base.Ui32(int32(160)) {
					v22 = int32(161)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v22)
					return int32(1)
				} else {
					v26 = int32(1)
					v27 = v13 + v26
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v27)
					return v26
				}
			}
		}
	case 1:
		if l1 != int32(3) {
			v80 = v3
			return v80
		} else {
			v34 = l0 + int32(2)
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
			if base.Ui32(v35) < base.Ui32(int32(161)) {
				v82 = v34
				v84 = int32(161)
				*(*uint8)(unsafe.Add(mBase, uint32(v82))) = uint8(v84)
				return int32(1)
			} else {
				if base.Ui32(int32(253)) < base.Ui32(v35) {
					v72 = l0 + int32(1)
					v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
					if base.Ui32(v73) < base.Ui32(int32(161)) {
						v82 = v72
						v84 = int32(161)
						*(*uint8)(unsafe.Add(mBase, uint32(v82))) = uint8(v84)
						return int32(1)
					} else {
						if base.Ui32(int32(254)) <= base.Ui32(v73) {
							v80 = v3
							return v80
						} else {
							v88 = v72
							v89 = v73
							v90 = int32(1)
							v91 = v89 + v90
							*(*uint8)(unsafe.Add(mBase, uint32(v88))) = uint8(v91)
							return v90
						}
					}
				} else {
					v88 = v34
					v89 = v35
					v90 = int32(1)
					v91 = v89 + v90
					*(*uint8)(unsafe.Add(mBase, uint32(v88))) = uint8(v91)
					return v90
				}
			}
		}
	default:
		if base.I32_extend8_s(v6) < int32(0) {
			if l1 != int32(2) {
				v80 = v3
				return v80
			} else {
				v46 = l0 + int32(1)
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				if base.Ui32(int32(161)) <= base.Ui32(v47) {
					if base.Ui32(v47) <= base.Ui32(int32(253)) {
						v69 = v46
						v70 = v47
						v88 = v69
						v89 = v70
						v90 = int32(1)
						v91 = v89 + v90
						*(*uint8)(unsafe.Add(mBase, uint32(v88))) = uint8(v91)
						return v90
					} else {
						if base.Ui32(int32(161)) <= base.Ui32(v6) {
							if base.Ui32(int32(254)) <= base.Ui32(v6) {
								v80 = v3
								return v80
							} else {
								v69 = l0
								v70 = v6
								v88 = v69
								v89 = v70
								v90 = int32(1)
								v91 = v89 + v90
								*(*uint8)(unsafe.Add(mBase, uint32(v88))) = uint8(v91)
								return v90
							}
						} else {
							v54 = l0
							v55 = int32(161)
							*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v55)
							return int32(1)
						}
					}
				} else {
					v54 = v46
					v55 = int32(161)
					*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v55)
					return int32(1)
				}
			}
		} else {
			if v6 == int32(127) {
				return int32(0)
			} else {
				v63 = int32(1)
				v65 = v6 + v63
				*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v65)
				v80 = v63
				return v80
			}
		}
	}
}
func F_pg_eucjp_verifychar(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v5 - int32(142) {
	case 0:
		v8 = int32(2)
		v9 = int32(-1)
		if l1 < v8 {
			v62 = v9
		} else {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			if base.Ui32(int32(193)) <= base.Ui32((v12+int32(32))&int32(255)) {
				v59 = v8
				v62 = v59
			} else {
				v62 = v9
			}
		}
	case 1:
		v19 = int32(-1)
		if l1 < int32(3) {
			v62 = v19
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			if base.Ui32(int32(93)) < base.Ui32((v22+int32(95))&int32(255)) {
				v62 = v19
			} else {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
				if base.Ui32(int32(94)) <= base.Ui32((v29+int32(95))&int32(255)) {
					v62 = v19
				} else {
					v59 = int32(3)
					v62 = v59
				}
			}
		}
	default:
		v37 = int32(-1)
		v38 = base.I32_extend8_s(v5)
		if int32(0) <= v38 {
			v59 = int32(1)
			v62 = v59
		} else {
			if base.B2i32(l1 < int32(2))|base.B2i32(base.Ui32(int32(93)) < base.Ui32((v38+int32(95))&int32(255))) != 0 {
				v62 = v37
			} else {
				v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				if base.Ui32(int32(93)) < base.Ui32((v51+int32(95))&int32(255)) {
					v62 = v37
				} else {
					v59 = int32(2)
					v62 = v59
				}
			}
		}
	}
	return v62
}
func F_pg_euckr_verifystr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	if l1 <= int32(0) {
		v41 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v41 - l0
L2:
	;
	v8 = l1
	v9 = l0
	goto L3
L3:
	;
	v11 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9))))
	if int32(0) <= v11 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v41 = v35
	goto L1
L5:
	;
	v35 = v34 + v9
	v36 = v8 - v34
	if int32(0) < v36 {
		v8 = v36
		v9 = v35
		goto L3
	} else {
		goto L12
	}
L6:
	;
	if v11 == int32(0) {
		v41 = v9
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if base.B2i32(v8 == int32(1))|base.B2i32(base.Ui32(int32(93)) < base.Ui32((v11+int32(95))&int32(255))) != 0 {
		v41 = v9
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v34 = int32(1)
	goto L5
L10:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if base.Ui32(int32(93)) < base.Ui32((v26+int32(95))&int32(255)) {
		v41 = v9
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v34 = int32(2)
	goto L5
L12:
	;
	goto L4
}
func F_pg_event_trigger_ddl_commands(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
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
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int64
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int64
	_ = v337
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int64
	_ = v355
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	v11 = m.G0
	v13 = v11 - int32(160)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_pg_event_trigger_ddl_commands[0]))
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L6
	} else {
		goto L116
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L6
	} else {
		goto L113
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L6
	} else {
		goto L109
	}
L6:
	;
	return int32(0)
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_pg_event_trigger_ddl_commands[0]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	if v25 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	m.G0 = v13 + int32(160)
	return int32(0)
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v28 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v37 = int32(0)
	goto L11
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v37<<(uint(int32(2))%32))))
	v46 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+104)) = uint8(v46)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = int64(0)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	switch v50 {
	case 0:
		goto L23
	case 1:
		goto L22
	case 2:
		goto L16
	case 3:
		goto L21
	case 4:
		goto L17
	case 5:
		goto L20
	case 6:
		goto L19
	default:
		goto L14
	}
L12:
	;
	goto L8
L13:
	;
	v298 = v37 + int32(1)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v298 < v299 {
		v37 = v298
		goto L11
	} else {
		goto L108
	}
L14:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	F_tuplestore_putvalues(m, v284, v285, v13+int32(112), v13+int32(96))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L6
	} else {
		goto L107
	}
L15:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+144)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v13)+140)) = v276
	goto L14
L16:
	;
	v219 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+98)) = uint8(v219)
	v221 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+96)) = uint16(v221)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	if v226 != 0 {
		goto L81
	} else {
		goto L82
	}
L17:
	;
	v168 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+98)) = uint8(v168)
	v170 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+96)) = uint16(v170)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v173 = F_CreateCommandTag(m, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L6
	} else {
		goto L57
	}
L18:
	;
	v78 = v13 + int32(80)
	v79 = int32(0)
	v82 = F_getObjectIdentityParts(m, v78, v79, v79, int32(1))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L25
	}
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v72
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v45)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+80)) = v74
	goto L18
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v68
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v45)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+80)) = v70
	goto L18
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v64
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v45)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+80)) = v66
	goto L18
L22:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v60
	goto L18
L23:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	if v51 == int32(0) {
		goto L13
	} else {
		goto L24
	}
L24:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v54
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v45)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+80)) = v56
	goto L18
L25:
	;
	if v82 == int32(0) {
		goto L13
	} else {
		goto L26
	}
L26:
	;
	v86 = int32(0)
	v88 = F_getObjectTypeDescription(m, v78, int32(1))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v13)+80))
	v94 = int32(0)
	goto L30
L28:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v13)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = v142
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v13)+84))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+116)) = v144
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v147 = F_CreateCommandTag(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L6
	} else {
		goto L47
	}
L29:
	;
	if v99 == int32(0) {
		v139 = v86
		goto L28
	} else {
		goto L36
	}
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94*int32(40))+uint32(_c_F_pg_event_trigger_ddl_commands[1])))
	v99 = base.B2i32(v98 == v90)
	if v99 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	v103 = v94 + int32(1)
	if v103 != int32(37) {
		v94 = v103
		goto L30
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	goto L31
L35:
	;
	goto L34
L36:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v13)+80))
	v110 = F_get_object_attnum_namespace(m, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	if v110 == int32(0) {
		v139 = v86
		goto L28
	} else {
		goto L38
	}
L38:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v13)+80))
	v116 = F_table_open(m, v114, int32(1))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+80))
	v119 = F_get_object_attnum_oid(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v13)+84))
	v122 = F_get_catalog_object_by_oid(m, v116, v119, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	if v122 == int32(0) {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v116)+52))
	v129 = F_heap_getattr_2(m, v122, v110, v126, v13+int32(112))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+112)))
	if v131 == int32(1) {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v134 = F_get_namespace_name_or_temp(m, v129)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	F_relation_close(m, v116, int32(1))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	v139 = v134
	goto L28
L47:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v147<<(uint(int32(3))%32))+uint32(_c_F_pg_event_trigger_ddl_commands[2])))
	goto L48
L48:
	;
	v152 = F_cstring_to_text(m, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+124)) = v152
	v155 = F_cstring_to_text(m, v88)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+128)) = v155
	if v139 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v165 = F_cstring_to_text(m, v82)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L6
	} else {
		goto L56
	}
L52:
	;
	v160 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+101)) = uint8(v160)
	goto L51
L53:
	;
	goto L54
L54:
	;
	v162 = F_cstring_to_text(m, v139)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+132)) = v162
	goto L51
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+136)) = v165
	goto L15
L57:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173<<(uint(int32(3))%32))+uint32(_c_F_pg_event_trigger_ddl_commands[2])))
	goto L58
L58:
	;
	v178 = F_cstring_to_text(m, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+124)) = v178
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	switch v182 {
	case 0, 1, 2, 3, 4, 5, 7, 8, 10, 11, 13, 14, 15, 18, 20, 23, 24, 25, 26, 27, 28, 30, 31, 32, 33, 35, 38, 39, 40, 43, 44, 45, 46, 47, 48, 50, 51:
		goto L62
	case 6:
		v213 = int32(_a_F_pg_event_trigger_ddl_commands_0)
		goto L60
	case 9:
		goto L74
	case 12:
		goto L73
	case 16:
		goto L72
	case 17:
		goto L71
	case 19:
		goto L70
	case 21:
		goto L69
	case 22:
		goto L68
	case 29:
		goto L66
	case 34:
		goto L65
	case 36:
		goto L67
	case 37:
		goto L75
	case 41:
		goto L76
	case 42:
		goto L64
	case 49:
		goto L63
	default:
		goto L61
	}
L60:
	;
	v214 = F_cstring_to_text(m, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L6
	} else {
		goto L80
	}
L61:
	;
	v213 = int32(_a_F_pg_event_trigger_ddl_commands_1)
	goto L60
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L6
	} else {
		goto L77
	}
L63:
	;
	v213 = int32(_a_F_pg_event_trigger_ddl_commands_2)
	goto L60
L64:
	;
	v213 = int32(_a_F_pg_event_trigger_ddl_commands_3)
	goto L60
L65:
	;
	v213 = int32(_a_F_pg_event_trigger_ddl_commands_4)
	goto L60
L66:
	;
	v213 = int32(_a_F_pg_event_trigger_ddl_commands_5)
	goto L60
L67:
	;
	v213 = int32(_a_F_pg_event_trigger_ddl_commands_6)
	goto L60
L68:
	;
	v213 = int32(_a_F_pg_event_trigger_ddl_commands_7)
	goto L60
L69:
	;
	v213 = int32(_a_F_pg_event_trigger_ddl_commands_8)
	goto L60
L70:
	;
	v213 = int32(_a_F_pg_event_trigger_ddl_commands_9)
	goto L60
L71:
	;
	v213 = int32(_a_F_pg_event_trigger_ddl_commands_10)
	goto L60
L72:
	;
	v213 = int32(_a_F_pg_event_trigger_ddl_commands_11)
	goto L60
L73:
	;
	v213 = int32(_a_F_pg_event_trigger_ddl_commands_12)
	goto L60
L74:
	;
	v213 = int32(_a_F_pg_event_trigger_ddl_commands_13)
	goto L60
L75:
	;
	v213 = int32(_a_F_pg_event_trigger_ddl_commands_14)
	goto L60
L76:
	;
	v213 = int32(_a_F_pg_event_trigger_ddl_commands_15)
	goto L60
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v182
	F_errmsg_internal(m, int32(_a_F_pg_event_trigger_ddl_commands_16), v13-int32(-64))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_pg_event_trigger_ddl_commands_17), int32(2405), int32(_a_F_pg_event_trigger_ddl_commands_18))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	v216 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+101)) = uint16(v216)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+128)) = v214
	goto L15
L81:
	;
	v227 = int32(_a_F_pg_event_trigger_ddl_commands_19)
	goto L83
L82:
	;
	v227 = int32(_a_F_pg_event_trigger_ddl_commands_20)
	goto L83
L83:
	;
	v228 = F_cstring_to_text(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+124)) = v228
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	switch v233 {
	case 0, 1, 2, 3, 4, 5, 7, 8, 10, 11, 13, 14, 15, 18, 20, 23, 24, 25, 26, 28, 30, 31, 32, 33, 35, 38, 39, 40, 43, 44, 45, 46, 47, 48, 50, 51:
		goto L87
	case 6:
		v265 = int32(_a_F_pg_event_trigger_ddl_commands_21)
		goto L85
	case 9:
		goto L100
	case 12:
		goto L99
	case 16:
		goto L98
	case 17:
		goto L97
	case 19:
		goto L96
	case 21:
		goto L95
	case 22:
		goto L94
	case 27:
		goto L92
	case 29:
		goto L91
	case 34:
		goto L90
	case 36:
		goto L93
	case 37:
		goto L101
	case 41:
		goto L102
	case 42:
		goto L89
	case 49:
		goto L88
	default:
		goto L86
	}
L85:
	;
	v266 = F_cstring_to_text(m, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L6
	} else {
		goto L106
	}
L86:
	;
	v265 = int32(_a_F_pg_event_trigger_ddl_commands_1)
	goto L85
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L6
	} else {
		goto L103
	}
L88:
	;
	v265 = int32(_a_F_pg_event_trigger_ddl_commands_22)
	goto L85
L89:
	;
	v265 = int32(_a_F_pg_event_trigger_ddl_commands_23)
	goto L85
L90:
	;
	v265 = int32(_a_F_pg_event_trigger_ddl_commands_24)
	goto L85
L91:
	;
	v265 = int32(_a_F_pg_event_trigger_ddl_commands_25)
	goto L85
L92:
	;
	v265 = int32(_a_F_pg_event_trigger_ddl_commands_26)
	goto L85
L93:
	;
	v265 = int32(_a_F_pg_event_trigger_ddl_commands_27)
	goto L85
L94:
	;
	v265 = int32(_a_F_pg_event_trigger_ddl_commands_28)
	goto L85
L95:
	;
	v265 = int32(_a_F_pg_event_trigger_ddl_commands_29)
	goto L85
L96:
	;
	v265 = int32(_a_F_pg_event_trigger_ddl_commands_30)
	goto L85
L97:
	;
	v265 = int32(_a_F_pg_event_trigger_ddl_commands_31)
	goto L85
L98:
	;
	v265 = int32(_a_F_pg_event_trigger_ddl_commands_32)
	goto L85
L99:
	;
	v265 = int32(_a_F_pg_event_trigger_ddl_commands_33)
	goto L85
L100:
	;
	v265 = int32(_a_F_pg_event_trigger_ddl_commands_34)
	goto L85
L101:
	;
	v265 = int32(_a_F_pg_event_trigger_ddl_commands_35)
	goto L85
L102:
	;
	v265 = int32(_a_F_pg_event_trigger_ddl_commands_36)
	goto L85
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v233
	F_errmsg_internal(m, int32(_a_F_pg_event_trigger_ddl_commands_16), v13+int32(48))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_pg_event_trigger_ddl_commands_17), int32(2321), int32(_a_F_pg_event_trigger_ddl_commands_37))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	v268 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+101)) = uint16(v268)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+128)) = v266
	goto L15
L107:
	;
	goto L13
L108:
	;
	goto L12
L109:
	;
	F_errcode(m, int32(50463299))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L6
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_pg_event_trigger_ddl_commands_38)
	F_errmsg(m, int32(_a_F_pg_event_trigger_ddl_commands_39), v13)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L6
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_pg_event_trigger_ddl_commands_17), int32(2064), int32(_a_F_pg_event_trigger_ddl_commands_40))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L6
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	v337 = *(*int64)(unsafe.Add(mBase, uint32(v13)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v337
	F_errmsg_internal(m, int32(_a_F_pg_event_trigger_ddl_commands_41), v13+int32(16))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_pg_event_trigger_ddl_commands_17), int32(2154), int32(_a_F_pg_event_trigger_ddl_commands_40))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v13)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v353
	v355 = *(*int64)(unsafe.Add(mBase, uint32(v13)+84))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+36)) = v355
	F_errmsg_internal(m, int32(_a_F_pg_event_trigger_ddl_commands_42), v13+int32(32))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_pg_event_trigger_ddl_commands_17), int32(2161), int32(_a_F_pg_event_trigger_ddl_commands_40))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L6
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_event_trigger_dropped_objects(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int64
	_ = v29
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_pg_event_trigger_dropped_objects[0]))
	if v10 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L39
	}
L2:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+8)))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_pg_event_trigger_dropped_objects[0]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v25 = v24
	goto L9
L7:
	;
	goto L8
L8:
	;
	m.G0 = v7 + int32(80)
	return int32(0)
L9:
	;
	v29 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+72)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v7)+64)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v7)+56)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v7)+48)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v29
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v25-int32(40))))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v47
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v25-int32(36))))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v51
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v25-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v55
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25-int32(4)))))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = v59
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25-int32(3)))))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v63
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25-int32(2)))))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v67
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v25-int32(16))))
	v72 = F_cstring_to_text(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v72
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v25-int32(28))))
	if v77 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v25-int32(24))))
	if v85 != 0 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v78 = F_cstring_to_text(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v81 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+23)) = uint8(v81)
	goto L12
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+60)) = v78
	goto L12
L17:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v25-int32(20))))
	if v93 != 0 {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v86 = F_cstring_to_text(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v89 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+24)) = uint8(v89)
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v86
	goto L17
L22:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v25-int32(12))))
	if v101 != 0 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v94 = F_cstring_to_text(m, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v97 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+25)) = uint8(v97)
	goto L22
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+68)) = v94
	goto L22
L27:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	F_tuplestore_putvalues(m, v118, v119, v7+int32(32), v7+int32(16))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L37
	}
L28:
	;
	v102 = F_strlist_to_textarray(m, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v115 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+26)) = uint16(v115)
	goto L27
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+72)) = v102
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v25-int32(8))))
	if v107 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v108 = F_strlist_to_textarray(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v112 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+76)) = v108
	goto L27
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+76)) = v112
	goto L27
L37:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v126 != 0 {
		v25 = v126
		goto L9
	} else {
		goto L38
	}
L38:
	;
	goto L10
L39:
	;
	F_errcode(m, int32(50463299))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_pg_event_trigger_dropped_objects_0)
	F_errmsg(m, int32(_a_F_pg_event_trigger_dropped_objects_1), v7)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_pg_event_trigger_dropped_objects_2), int32(1537), int32(_a_F_pg_event_trigger_dropped_objects_3))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_export_snapshot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pg_export_snapshot[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = F_ExportSnapshot(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_cstring_to_text(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
func F_pg_flush_data(m *base.Module, l0 int32) {
	var v2 int32
	_ = v2
	v2 = F_fsync(m, l0)
	return
}
func F_pg_fprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	v5 = m.G0
	v7 = v5 - int32(1072)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	if l0 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_pg_fprintf[0])) = int32(28)
		v58 = int32(-1)
		m.G0 = v7 + int32(1072)
		return v58
	} else {
		v16 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+1068)) = uint8(v16)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1064)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1060)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1056)) = v7 + int32(1040)
		v25 = v7 + int32(16)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1052)) = v25
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1048)) = v25
		F_dopr(m, v7+int32(1048), l1, l2)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1068)))
			if v34 == int32(0) {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1048))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1052))
				if v37 != v38 {
					v47 = v37 - v38
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1060))
					v49 = F_fwrite(m, v38, int32(1), v47, v48)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1064))
						if v47 != v49 {
							v54 = int32(-1)
						} else {
							v54 = v49 + v51
						}
						v58 = v54
						m.G0 = v7 + int32(1072)
						return v58
					}
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1064))
					if v34 != 0 {
						v44 = int32(-1)
					} else {
						v44 = v43
					}
					v58 = v44
					m.G0 = v7 + int32(1072)
					return v58
				}
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1064))
				if v34 != 0 {
					v44 = int32(-1)
				} else {
					v44 = v43
				}
				v58 = v44
				m.G0 = v7 + int32(1072)
				return v58
			}
		}
	}
}
func F_pg_freeaddrinfo_all(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	if l0 == int32(1) {
		if l1 == int32(0) {
		} else {
			v8 = l1
			for {
				v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
				v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
				F_emscripten_builtin_free(m, v10)
				mBase = m.M
				F_emscripten_builtin_free(m, v8)
				mBase = m.M
				if v9 != 0 {
					v8 = v9
					continue
				} else {
					break
				}
				break
			}
		}
	} else {
		if l1 == int32(0) {
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			F_emscripten_builtin_free(m, v15)
			mBase = m.M
			F_emscripten_builtin_free(m, l1)
			mBase = m.M
		}
	}
	return
}
func F_pg_gen_salt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v5 = m.G0
	v7 = v5 - int32(160)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = v7 + int32(16)
	F_text_to_cstring_buffer(m, v10, v15, int32(129))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = F_px_gen_salt(m, v15, v15, int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if int32(0) <= v20 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v24 != v10 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	F_pfree(m, v10)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v30 = F_cstring_to_text_with_len(m, v7+int32(16), v20)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	m.G0 = v7 + int32(160)
	return v30
L13:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v20 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v72
	F_errmsg(m, int32(_a_F_pg_gen_salt_0), v7)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L29
	}
L16:
	;
	v72 = int32(_a_F_pg_gen_salt_1)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v51 = int32(_a_F_pg_gen_salt_2)
	goto L20
L19:
	;
	v72 = v66
	goto L15
L20:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	if v20 != v54 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v66 = v64
	goto L19
L22:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
	if v56 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	v72 = int32(_a_F_pg_gen_salt_3)
	goto L15
L26:
	;
	goto L27
L27:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	if v20 != v60 {
		v51 = v51 + int32(16)
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v66 = v56
	goto L19
L29:
	;
	F_errfinish(m, int32(_a_F_pg_gen_salt_4), int32(179), int32(_a_F_pg_gen_salt_5))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_acl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v14|v15 == v2 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L32
	} else {
		goto L66
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L32
	} else {
		goto L63
	}
L3:
	;
	m.G0 = v12 + int32(32)
	return v194
L4:
	;
	v189 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v189)
	v194 = int32(0)
	goto L3
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v14 == int32(2613) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v23 = int32(2995)
	goto L8
L7:
	;
	v23 = v14
	goto L8
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_acl[0]))
	if v25 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(v66)+28)))
	if v74 == int32(0) {
		goto L4
	} else {
		goto L27
	}
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v26 == v23 {
		v66 = v25
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v35 = v2
	goto L18
L13:
	;
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_get_acl[0])) = v63
	v66 = v63
	goto L9
L15:
	;
	v63 = v39 + int32(_a_F_pg_get_acl_0)
	goto L14
L16:
	;
	v63 = v39 + int32(_a_F_pg_get_acl_1)
	goto L14
L17:
	;
	v63 = v39 + int32(_a_F_pg_get_acl_2)
	goto L14
L18:
	;
	v39 = v35 * int32(40)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_pg_get_acl[1])))
	if v23 != v40 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v63 = v39 + int32(_a_F_pg_get_acl_3)
	goto L14
L20:
	;
	if v35 == int32(36) {
		goto L2
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	goto L19
L23:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_pg_get_acl[2])))
	if v46 == v23 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_pg_get_acl[3])))
	if v48 == v23 {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_pg_get_acl[4])))
	if v50 == v23 {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	v35 = v35 + int32(4)
	goto L18
L27:
	;
	v77 = int32(0)
	if base.B2i32(v19 == v77)|base.B2i32(v14 != int32(1259)) == v77 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)))
	if v177 != int32(1) {
		v194 = v170
		goto L3
	} else {
		goto L62
	}
L29:
	;
	v85 = F_SearchSysCacheCopyAttNum(m, v15, base.I32_extend16_s(v19))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v99 = F_table_open(m, v23, int32(1))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L32
	} else {
		goto L36
	}
L32:
	;
	return int32(0)
L33:
	;
	if v85 == int32(0) {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v95 = F_SysCacheGetAttr(m, int32(7), v85, int32(22), v12+int32(31))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v170 = v95
	goto L28
L36:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_acl[0]))
	if v102 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v151 = int32(*(*int16)(unsafe.Add(mBase, uint32(v143)+20)))
	v153 = F_get_catalog_object_by_oid_extended(m, v99, v151, v15, int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L32
	} else {
		goto L55
	}
L38:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	if v103 == v23 {
		v143 = v102
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v112 = int32(0)
	goto L46
L41:
	;
	goto L40
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_get_acl[0])) = v140
	v143 = v140
	goto L37
L43:
	;
	v140 = v116 + int32(_a_F_pg_get_acl_0)
	goto L42
L44:
	;
	v140 = v116 + int32(_a_F_pg_get_acl_1)
	goto L42
L45:
	;
	v140 = v116 + int32(_a_F_pg_get_acl_2)
	goto L42
L46:
	;
	v116 = v112 * int32(40)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+uint32(_c_F_pg_get_acl[1])))
	if v23 != v117 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v140 = v116 + int32(_a_F_pg_get_acl_3)
	goto L42
L48:
	;
	if v112 == int32(36) {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v116)+uint32(_c_F_pg_get_acl[2])))
	if v123 == v23 {
		goto L43
	} else {
		goto L52
	}
L52:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v116)+uint32(_c_F_pg_get_acl[3])))
	if v125 == v23 {
		goto L44
	} else {
		goto L53
	}
L53:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v116)+uint32(_c_F_pg_get_acl[4])))
	if v127 == v23 {
		goto L45
	} else {
		goto L54
	}
L54:
	;
	v112 = v112 + int32(4)
	goto L46
L55:
	;
	if v153 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_relation_close(m, v99, int32(1))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L32
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v99)+52))
	v163 = F_heap_getattr_2(m, v153, v74, v160, v12+int32(31))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L32
	} else {
		goto L60
	}
L59:
	;
	goto L4
L60:
	;
	F_relation_close(m, v99, int32(1))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L32
	} else {
		goto L61
	}
L61:
	;
	v170 = v163
	goto L28
L62:
	;
	goto L4
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v23
	F_errmsg_internal(m, int32(_a_F_pg_get_acl_4), v12+int32(16))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L32
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_pg_get_acl_5), int32(2777), int32(_a_F_pg_get_acl_6))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L32
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v23
	F_errmsg_internal(m, int32(_a_F_pg_get_acl_4), v12)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L32
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_pg_get_acl_5), int32(2777), int32(_a_F_pg_get_acl_6))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L32
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_aios(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v40 int64
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int64
	_ = v55
	var v73 int64
	_ = v73
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int64
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v267 int64
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int64
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v381 int64
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v498 int64
	_ = v498
	var v500 int32
	_ = v500
	var v501 int64
	_ = v501
	v14 = m.G0
	v16 = v14 - int32(1232)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_aios[0]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	if v26 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v28 = v16 + int32(1156)
	v29 = v25
	v40 = int64(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	m.G0 = v16 + int32(1232)
	return int32(0)
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	v44 = int32(7)
	v46 = v42 + base.I32_wrap_i64(v40)<<(uint(v44)%32)
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_aios[0]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	goto L8
L7:
	;
	goto L5
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1224)) = int32(0)
	v55 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+1216)) = v55
	*(*int64)(unsafe.Add(mBase, uint32(v16)+1208)) = v55
	*(*int64)(unsafe.Add(mBase, uint32(v16)+1200)) = v55
	*(*int64)(unsafe.Add(mBase, uint32(v16)+1192)) = v55
	*(*int64)(unsafe.Add(mBase, uint32(v16)+1184)) = v55
	*(*int64)(unsafe.Add(mBase, uint32(v16)+1176)) = v55
	*(*int64)(unsafe.Add(mBase, uint32(v16)+1168)) = v55
	*(*int64)(unsafe.Add(mBase, uint32(v16)+1159)) = v55
	*(*int64)(unsafe.Add(mBase, uint32(v16)+1152)) = v55
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v46)+48))
	goto L10
L9:
	;
	v498 = v40 + int64(1)
	v500 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_aios[0]))
	v501 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v500)+20)))
	if base.Ui64(v498) < base.Ui64(v501) {
		v29 = v500
		v40 = v498
		goto L6
	} else {
		goto L84
	}
L10:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v87 == int32(0) {
		goto L9
	} else {
		goto L12
	}
L11:
	;
	if v110 != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v90 = int32(1024)
	base.MemoryCopy(m, v16+v90, v46, int32(128))
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_aios[0]))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1100))
	base.MemoryCopy(m, v16, v96+v97<<(uint(int32(3))%32), v90)
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_aios[1]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1040))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v105+v106*int32(640))+44))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v46)+48))
	if v111 != v73 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v113 != v87 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1172)) = (v46 - v49) >> (uint(v44) % 32)
	v119 = F_Int64GetDatum(m, v73)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L19
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1168)) = v110
	goto L15
L17:
	;
	goto L18
L18:
	;
	v116 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1152)) = uint8(v116)
	goto L15
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1176)) = v119
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(1024)))))
	if base.Ui32(v124) <= base.Ui32(int32(7)) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v132 = F_cstring_to_text(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L24
	}
L21:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v124<<(uint(int32(2))%32))+uint32(_c_F_pg_get_aios[2])))
	v131 = v129
	goto L23
L22:
	;
	v131 = int32(0)
	goto L23
L23:
	;
	goto L20
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1180)) = v132
	if v87 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	F_tuplestore_putvalues(m, v476, v477, v16+int32(1168), v16+int32(1152))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L83
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+7)) = int32(16843009)
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = int64(72340172838076673)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(1024))+2)))
	if base.Ui32(v143) <= base.Ui32(int32(2)) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v151 = F_cstring_to_text(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L33
	}
L30:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v143<<(uint(int32(2))%32))+uint32(_c_F_pg_get_aios[3])))
	v150 = v148
	goto L32
L31:
	;
	v150 = int32(0)
	goto L32
L32:
	;
	goto L29
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1184)) = v151
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1026)))
	switch v154 {
	case 0:
		goto L37
	case 1:
		goto L36
	case 2:
		goto L35
	default:
		goto L34
	}
L34:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(1024))+1)))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v400<<(uint(int32(2))%32))+uint32(_c_F_pg_get_aios[4])))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)+8))
	goto L70
L35:
	;
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v16)+1120))
	v272 = F_Int64GetDatum(m, v271)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L54
	}
L36:
	;
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v16)+1120))
	v158 = F_Int64GetDatum(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	v155 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+1157)) = uint16(v155)
	goto L34
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1188)) = v158
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+1116)))
	if v161 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v163 = v161 & int32(3)
	v164 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v161) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v267 = int64(0)
	goto L41
L41:
	;
	v268 = F_Int64GetDatum(m, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L53
	}
L42:
	;
	v267 = base.I64_extend_i32_u(v241)
	goto L41
L43:
	;
	v172 = v164
	v174 = v164
	v175 = int32(0)
	goto L46
L44:
	;
	v203 = v164
	v205 = v164
	goto L45
L45:
	;
	v216 = v203
	v218 = v205
	v220 = v164
	goto L50
L46:
	;
	v187 = v16 + v172<<(uint(int32(3))%32)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+28))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v187)+20))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v195 = v188 + (v189 + (v190 + (v191 + v174)))
	v196 = int32(4)
	v197 = v172 + v196
	v199 = v175 + v196
	if v199 != v161&int32(_a_F_pg_get_aios_0) {
		v172 = v197
		v174 = v195
		v175 = v199
		goto L46
	} else {
		goto L48
	}
L47:
	;
	if v163 == int32(0) {
		v241 = v195
		goto L42
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	v203 = v197
	v205 = v195
	goto L45
L50:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v16+v216<<(uint(int32(3))%32))+4))
	v233 = v232 + v218
	v234 = int32(1)
	v237 = v220 + v234
	if v237 != v163 {
		v216 = v216 + v234
		v218 = v233
		v220 = v237
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v241 = v233
	goto L42
L52:
	;
	goto L51
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1192)) = v268
	goto L34
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1188)) = v272
	v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+1116)))
	if v275 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v277 = v275 & int32(3)
	v278 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v275) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v381 = int64(0)
	goto L57
L57:
	;
	v382 = F_Int64GetDatum(m, v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L69
	}
L58:
	;
	v381 = base.I64_extend_i32_u(v355)
	goto L57
L59:
	;
	v286 = v278
	v288 = v278
	v289 = int32(0)
	goto L62
L60:
	;
	v317 = v278
	v319 = v278
	goto L61
L61:
	;
	v330 = v317
	v332 = v319
	v334 = v278
	goto L66
L62:
	;
	v301 = v16 + v286<<(uint(int32(3))%32)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)+28))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v301)+20))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v301)+12))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	v309 = v302 + (v303 + (v304 + (v305 + v288)))
	v310 = int32(4)
	v311 = v286 + v310
	v313 = v289 + v310
	if v313 != v275&int32(_a_F_pg_get_aios_0) {
		v286 = v311
		v288 = v309
		v289 = v313
		goto L62
	} else {
		goto L64
	}
L63:
	;
	if v277 == int32(0) {
		v355 = v309
		goto L58
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	v317 = v311
	v319 = v309
	goto L61
L66:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v16+v330<<(uint(int32(3))%32))+4))
	v347 = v346 + v332
	v348 = int32(1)
	v351 = v334 + v348
	if v351 != v277 {
		v330 = v330 + v348
		v332 = v347
		v334 = v351
		goto L66
	} else {
		goto L68
	}
L67:
	;
	v355 = v347
	goto L58
L68:
	;
	goto L67
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1192)) = v382
	goto L34
L70:
	;
	v405 = F_cstring_to_text(m, v404)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1196)) = v405
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1037)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1200)) = v408
	if base.Ui32((v87-int32(5))&int32(255)) <= base.Ui32(int32(2)) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1092))
	v424 = int32(base.Ui32(v420)>>(uint(int32(6))%32)) & int32(7)
	if base.Ui32(v424) <= base.Ui32(int32(4)) {
		goto L77
	} else {
		goto L78
	}
L73:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1044))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1204)) = v416
	goto L72
L74:
	;
	goto L75
L75:
	;
	v418 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1161)) = uint8(v418)
	goto L72
L76:
	;
	v432 = F_cstring_to_text(m, v431)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L80
	}
L77:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v424<<(uint(int32(2))%32))+uint32(_c_F_pg_get_aios[5])))
	v431 = v429
	goto L79
L78:
	;
	v431 = int32(0)
	goto L79
L79:
	;
	goto L76
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1208)) = v432
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(1024))+1)))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v439<<(uint(int32(2))%32))+uint32(_c_F_pg_get_aios[4])))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	v444 = m.T0[v443].(func(*base.Module, int32) int32)(m, v16+int32(1128))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v446 = F_cstring_to_text(m, v444)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1212)) = v446
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1027)))
	v450 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1216)) = v449 & v450
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1224)) = int32(base.Ui32(v449)>>(uint(int32(2))%32)) & v450
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1220)) = int32(base.Ui32(v449)>>(uint(v450)%32)) & v450
	goto L25
L83:
	;
	goto L9
L84:
	;
	goto L7
}
func F_pg_get_constraintdef_ext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v7 != 0 {
		v8 = int32(7)
	} else {
		v8 = int32(2)
	}
	v10 = F_pg_get_constraintdef_worker(m, v3, int32(0), v8, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			v16 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v16)
			return int32(0)
		} else {
			v20 = F_cstring_to_text(m, v10)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v10)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					return v20
				}
			}
		}
	}
}
func F_pg_get_expr_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int64
	_ = v225
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = F_text_to_cstring(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = F_stringToNode(m, v13)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_pfree(m, v13)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v17 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v62 = int32(0)
	v64 = F_pull_varnos(m, v62, v17)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L18
	}
L6:
	;
	v23 = v17
	goto L7
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v31 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L5
L9:
	;
	if v31 != int32(67) {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v53 != 0 {
		v23 = v53
		goto L7
	} else {
		goto L17
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errmsg(m, int32(_a_F_pg_get_expr_worker_0), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_pg_get_expr_worker_1), int32(2740), int32(_a_F_pg_get_expr_worker_2))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	goto L8
L18:
	;
	if l1 != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	m.G0 = v11 - int32(-64)
	return v251
L20:
	;
	v218 = v9 + int32(-16)
	F_initStringInfo(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L61
	}
L21:
	;
	v158 = F_try_relation_open(m, l1, int32(1))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L50
	}
L22:
	;
	v67 = F_bms_make_singleton(m, int32(1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v64 == int32(0) {
		v212 = v62
		v215 = int32(0)
		goto L20
	} else {
		goto L45
	}
L25:
	;
	v69 = int32(0)
	if v64 == v69 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v122 != 0 {
		goto L21
	} else {
		goto L40
	}
L27:
	;
	v122 = int32(1)
	goto L26
L28:
	;
	goto L29
L29:
	;
	if v67 == int32(0) {
		v115 = v69
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v122 = v115
	goto L26
L31:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v79 < v78 {
		v115 = v69
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v81 = int32(1)
	if v78 <= v81 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v84 = v81
	goto L35
L34:
	;
	v84 = v78
	goto L35
L35:
	;
	v85 = int32(8)
	v90 = int32(0)
	goto L36
L36:
	;
	v97 = v90 << (uint(int32(2)) % 32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v64+v85+v97)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v67+v85+v97)))
	v104 = v99 & (v101 ^ int32(-1))
	v106 = base.B2i32(v104 == int32(0))
	if v104 != 0 {
		v115 = v106
		goto L30
	} else {
		goto L38
	}
L37:
	;
	v115 = v106
	goto L30
L38:
	;
	v108 = v90 + int32(1)
	if v108 != v84 {
		v90 = v108
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(_a_F_pg_get_expr_worker_3), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_pg_get_expr_worker_1), int32(2752), int32(_a_F_pg_get_expr_worker_2))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(_a_F_pg_get_expr_worker_4), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_pg_get_expr_worker_1), int32(2759), int32(_a_F_pg_get_expr_worker_2))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	if v158 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v251 = v62
	goto L19
L52:
	;
	goto L53
L53:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158)+48))
	v164 = F_palloc0(m, int32(80))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v167 = F_palloc0(m, int32(136))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+24)) = int32(1)
	v171 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v167)+21)) = uint8(v171)
	*(*int32)(unsafe.Add(mBase, uint32(v167)+16)) = l1
	v174 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v167)+12)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = int32(101)
	v181 = F_makeAlias(m, v162+int32(4), v174)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+8)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v167)+4)) = v181
	v185 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v167)+124)) = uint16(v185)
	v187 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v167)+20)) = uint8(v187)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v167
	v194 = F_list_make1_impl(m, int32(1), v9+int32(-60))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v196 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v164)+20)) = v196
	*(*int64)(unsafe.Add(mBase, uint32(v164)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v164))) = v194
	F_set_rtable_names(m, v164, v196, v196)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_set_simple_column_names(m, v164)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v164
	v210 = F_list_make1_impl(m, int32(1), v11)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v212 = v210
	v215 = v158
	goto L20
L61:
	;
	v221 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v221)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v221
	v225 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v221
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+43)) = uint8(v221)
	v232 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+41)) = uint16(v232)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v218
	F_get_rule_expr(m, v17, v9+int32(-56), v221)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	if v215 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	F_relation_close(m, v215, int32(1))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v247 = F_cstring_to_text(m, v243)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	F_pfree(m, v243)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v251 = v247
	goto L19
}
func F_pg_get_function_identity_arguments(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13983(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_pg_get_function_result(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_SearchSysCache1(m, int32(47), v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
			v42 = int32(0)
			m.G0 = v7 + int32(16)
			return v42
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+22)))
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v21)+96)))
			if v23 == int32(112) {
				F_ReleaseCatCache(m, v11)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
					v42 = int32(0)
					m.G0 = v7 + int32(16)
					return v42
				}
			} else {
				F_initStringInfo(m, v7)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					F_print_function_rettype(m, v7, v11)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_ReleaseCatCache(m, v11)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
							v38 = F_cstring_to_text(m, v37)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v37)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v42 = v38
									m.G0 = v7 + int32(16)
									return v42
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_get_partkeydef_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int64
	_ = v213
	var v220 int32
	_ = v220
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	v5 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(160)
	m.G0 = v20
	v23 = F_SearchSysCache1(m, int32(45), l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L6
	} else {
		goto L99
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L6
	} else {
		goto L96
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L6
	} else {
		goto L93
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L6
	} else {
		goto L90
	}
L5:
	;
	m.G0 = v20 + int32(160)
	return v346
L6:
	;
	return int32(0)
L7:
	;
	if v23 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l3 != 0 {
		v346 = int32(0)
		goto L5
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
	v47 = F_SysCacheGetAttrNotNull(m, int32(45), v23, int32(6))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L15
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l0
	F_errmsg_internal(m, int32(_a_F_pg_get_partkeydef_worker_0), v20)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_pg_get_partkeydef_worker_1), int32(1958), int32(_a_F_pg_get_partkeydef_worker_2))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L15:
	;
	v51 = F_SysCacheGetAttrNotNull(m, int32(45), v23, int32(7))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v55 = F_heap_attisnull(m, v23, int32(8), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	if v55 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v61 = F_SysCacheGetAttrNotNull(m, int32(45), v23, int32(8))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	v74 = v5
	v75 = v5
	goto L20
L20:
	;
	v76 = F_get_rel_name(m, l0)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L6
	} else {
		goto L26
	}
L21:
	;
	v63 = F_text_to_cstring(m, v61)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v65 = F_stringToNode(m, v63)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v67 != int32(1) {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	F_pfree(m, v63)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v74 = v72
	v75 = v65
	goto L20
L26:
	;
	if v76 == int32(0) {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	v81 = F_palloc0(m, int32(80))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v84 = F_palloc0(m, int32(136))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+24)) = int32(1)
	v88 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+21)) = uint8(v88)
	*(*int32)(unsafe.Add(mBase, uint32(v84)+16)) = l0
	v91 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v84)+12)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = int32(101)
	v96 = F_makeAlias(m, v76, v91)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+8)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = v96
	v100 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v84)+124)) = uint16(v100)
	v102 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+20)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v84
	v109 = F_list_make1_impl(m, int32(1), v20+int32(76))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v111 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+20)) = v111
	*(*int64)(unsafe.Add(mBase, uint32(v81)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v109
	F_set_rtable_names(m, v81, v111, v111)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	F_set_simple_column_names(m, v81)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v20)+144)) = v81
	v127 = F_list_make1_impl(m, int32(1), v20+int32(72))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	v130 = v20 + int32(88)
	F_initStringInfo(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	v133 = v44 + v43
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+4)))
	switch v134 - int32(104) {
	case 0:
		goto L40
	default:
		goto L2
	case 4:
		goto L39
	case 10:
		goto L38
	}
L36:
	;
	v148 = int32(*(*int16)(unsafe.Add(mBase, uint32(v133)+6)))
	if int32(0) < v148 {
		goto L46
	} else {
		goto L47
	}
L37:
	;
	F_appendStringInfoString(m, v130, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L6
	} else {
		goto L44
	}
L38:
	;
	if l2 != 0 {
		goto L36
	} else {
		goto L43
	}
L39:
	;
	if l2 != 0 {
		goto L36
	} else {
		goto L42
	}
L40:
	;
	if l2 != 0 {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v140 = int32(_a_F_pg_get_partkeydef_worker_3)
	goto L37
L42:
	;
	v140 = int32(_a_F_pg_get_partkeydef_worker_4)
	goto L37
L43:
	;
	v140 = int32(_a_F_pg_get_partkeydef_worker_5)
	goto L37
L44:
	;
	F_appendStringInfoString(m, v20+int32(88), int32(_a_F_pg_get_partkeydef_worker_6))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	goto L36
L46:
	;
	v151 = int32(24)
	v165 = int32(_a_F_pg_get_partkeydef_worker_7)
	v167 = v74
	v169 = int32(0)
	goto L49
L47:
	;
	goto L48
L48:
	;
	if l2 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L49:
	;
	v179 = int32(*(*int16)(unsafe.Add(mBase, uint32(v133+int32(36)+v169<<(uint(int32(1))%32)))))
	v181 = v20 + int32(88)
	F_appendStringInfoString(m, v181, v165)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L6
	} else {
		goto L51
	}
L50:
	;
	goto L48
L51:
	;
	if v179 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v270 = v169 << (uint(int32(2)) % 32)
	if l2 != 0 {
		goto L75
	} else {
		goto L76
	}
L53:
	;
	v185 = F_get_attname(m, l0, v179, int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L6
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if v167 == int32(0) {
		goto L1
	} else {
		goto L60
	}
L56:
	;
	v187 = F_quote_identifier(m, v185)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	F_appendStringInfoString(m, v181, v187)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	F_get_atttypetypmodcoll(m, l0, v179, v20+int32(104), v20+int32(84), v20+int32(144))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v20)+144))
	v264 = v199
	v267 = v167
	goto L52
L60:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v206 = v20 + int32(144)
	F_initStringInfo(m, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	v209 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+136)) = uint8(v209)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+120)) = v209
	v213 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+112)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v20)+140)) = v209
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+139)) = uint8(v209)
	v220 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+137)) = uint16(v220)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+128)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v20)+124)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v206
	F_get_rule_expr(m, v204, v20+int32(104), v209)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	v232 = v167 + int32(4)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v20)+144))
	if v204 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if base.Ui32(v232) < base.Ui32(v202+v203<<(uint(int32(2))%32)) {
		goto L70
	} else {
		goto L71
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v237
	F_appendStringInfo(m, v20+int32(88), int32(_a_F_pg_get_partkeydef_worker_8), v20-int32(-64))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L6
	} else {
		goto L69
	}
L65:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	switch v240 - int32(15) {
	case 0:
		goto L67
	default:
		goto L64
	case 4, 23, 24, 25, 26, 33:
		goto L66
	}
L66:
	;
	F_appendStringInfoString(m, v20+int32(88), v237)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L6
	} else {
		goto L68
	}
L67:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v204)+16))
	switch v243 {
	case 0, 3:
		goto L66
	default:
		goto L64
	}
L68:
	;
	goto L63
L69:
	;
	goto L63
L70:
	;
	v257 = v232
	goto L72
L71:
	;
	v257 = int32(0)
	goto L72
L72:
	;
	v258 = F_exprType(m, v204)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L6
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v258
	v261 = F_exprCollation(m, v204)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+144)) = v261
	v264 = v261
	v267 = v257
	goto L52
L75:
	;
	if l2 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v270+(v51+v151))))
	if base.B2i32(v272 == int32(0))|base.B2i32(v264 == v272) != 0 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v277 = F_generate_collation_name(m, v272)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v277
	F_appendStringInfo(m, v20+int32(88), int32(_a_F_pg_get_partkeydef_worker_9), v20+int32(48))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	goto L75
L80:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v270+(v47+v151))))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v20)+104))
	F_get_opclass_name(m, v291, v292, v20+int32(88))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L6
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v299 = v169 + int32(1)
	v300 = int32(*(*int16)(unsafe.Add(mBase, uint32(v133)+6)))
	if v299 < v300 {
		v165 = int32(_a_F_pg_get_partkeydef_worker_10)
		v167 = v267
		v169 = v299
		goto L49
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	goto L50
L85:
	;
	F_appendStringInfoChar(m, v20+int32(88), int32(41))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L6
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	F_ReleaseCatCache(m, v23)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L6
	} else {
		goto L89
	}
L88:
	;
	goto L87
L89:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	v346 = v328
	goto L5
L90:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v355
	F_errmsg_internal(m, int32(_a_F_pg_get_partkeydef_worker_11), v20+int32(80))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L6
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_pg_get_partkeydef_worker_1), int32(1992), int32(_a_F_pg_get_partkeydef_worker_2))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L6
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
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_pg_get_partkeydef_worker_12), v20+int32(16))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_pg_get_partkeydef_worker_1), int32(_a_F_pg_get_partkeydef_worker_13), int32(_a_F_pg_get_partkeydef_worker_14))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	v386 = int32(*(*int8)(unsafe.Add(mBase, uint32(v133)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v386
	F_errmsg_internal(m, int32(_a_F_pg_get_partkeydef_worker_15), v20+int32(32))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L6
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_pg_get_partkeydef_worker_1), int32(2020), int32(_a_F_pg_get_partkeydef_worker_2))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L6
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	F_errmsg_internal(m, int32(_a_F_pg_get_partkeydef_worker_16), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L6
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_pg_get_partkeydef_worker_1), int32(2053), int32(_a_F_pg_get_partkeydef_worker_2))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L6
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_ruledef_worker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int64
	_ = v55
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
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
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
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
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
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v527 int32
	_ = v527
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v594 int32
	_ = v594
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	v11 = m.G0
	v13 = v11 - int32(224)
	m.G0 = v13
	F_initStringInfo(m, v13+int32(72))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_ruledef_worker[0]))
	if v25 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L200
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L197
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L194
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L191
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+100)) = int32(26)
	v34 = F_SPI_prepare(m, int32(_a_F_pg_get_ruledef_worker_0), int32(1), v13+int32(100))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v42 = v25
	goto L10
L10:
	;
	v43 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)) = uint8(v43)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+92)) = l0
	v50 = F_SPI_execute_plan(m, v42, v13+int32(92), v13+int32(91))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	if v34 == int32(0) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	F_SPI_keepplan(m, v34)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_get_ruledef_worker[0])) = v34
	v42 = v34
	goto L10
L14:
	;
	if v50 != int32(5) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v55 = *(*int64)(unsafe.Add(mBase, _c_F_pg_get_ruledef_worker[1]))
	if v55 == int64(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_ruledef_worker[2]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v63 = int32(_a_F_pg_get_ruledef_worker_1)
	v64 = int32(0)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v64 < v66 {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	goto L18
L18:
	;
	v605 = F_SPI_finish(m)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L186
	}
L19:
	;
	v108 = v13 + int32(223)
	v109 = F_SPI_getbinval(m, v62, v60, v106, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L34
	}
L20:
	;
	v106 = v72 + int32(1)
	goto L19
L21:
	;
	v71 = v66
	v72 = v64
	goto L24
L22:
	;
	goto L23
L23:
	;
	v95 = F_SystemAttributeByName(m, v63)
	mBase = m.M
	if v95 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L24:
	;
	v78 = v60 + v71<<(uint(int32(4))%32) + v72*int32(100)
	v81 = F_namestrcmp(m, v78+int32(24), v63)
	mBase = m.M
	if v81 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L23
L26:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+111)))
	if v84 != int32(1) {
		goto L20
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v88 = v72 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v88 < v89 {
		v71 = v89
		v72 = v88
		goto L24
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	goto L25
L31:
	;
	v106 = int32(-9)
	goto L19
L32:
	;
	goto L33
L33:
	;
	v99 = int32(*(*int16)(unsafe.Add(mBase, uint32(v95)+74)))
	v106 = v99
	goto L19
L34:
	;
	v111 = int32(_a_F_pg_get_ruledef_worker_2)
	v112 = int32(0)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v112 < v114 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v155 = F_SPI_getbinval(m, v62, v60, v154, v108)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L50
	}
L36:
	;
	v154 = v120 + int32(1)
	goto L35
L37:
	;
	v119 = v114
	v120 = v112
	goto L40
L38:
	;
	goto L39
L39:
	;
	v143 = F_SystemAttributeByName(m, v111)
	mBase = m.M
	if v143 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L40:
	;
	v126 = v60 + v119<<(uint(int32(4))%32) + v120*int32(100)
	v129 = F_namestrcmp(m, v126+int32(24), v111)
	mBase = m.M
	if v129 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L39
L42:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+111)))
	if v132 != int32(1) {
		goto L36
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v136 = v120 + int32(1)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v136 < v137 {
		v119 = v137
		v120 = v136
		goto L40
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	goto L41
L47:
	;
	v154 = int32(-9)
	goto L35
L48:
	;
	goto L49
L49:
	;
	v147 = int32(*(*int16)(unsafe.Add(mBase, uint32(v143)+74)))
	v154 = v147
	goto L35
L50:
	;
	v157 = int32(_a_F_pg_get_ruledef_worker_3)
	v158 = int32(0)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v158 < v160 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v201 = F_SPI_getbinval(m, v62, v60, v200, v108)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L66
	}
L52:
	;
	v200 = v166 + int32(1)
	goto L51
L53:
	;
	v165 = v160
	v166 = v158
	goto L56
L54:
	;
	goto L55
L55:
	;
	v189 = F_SystemAttributeByName(m, v157)
	mBase = m.M
	if v189 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L56:
	;
	v172 = v60 + v165<<(uint(int32(4))%32) + v166*int32(100)
	v175 = F_namestrcmp(m, v172+int32(24), v157)
	mBase = m.M
	if v175 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L55
L58:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+111)))
	if v178 != int32(1) {
		goto L52
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v182 = v166 + int32(1)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v182 < v183 {
		v165 = v183
		v166 = v182
		goto L56
	} else {
		goto L62
	}
L61:
	;
	goto L60
L62:
	;
	goto L57
L63:
	;
	v200 = int32(-9)
	goto L51
L64:
	;
	goto L65
L65:
	;
	v193 = int32(*(*int16)(unsafe.Add(mBase, uint32(v189)+74)))
	v200 = v193
	goto L51
L66:
	;
	v203 = int32(_a_F_pg_get_ruledef_worker_4)
	v204 = int32(0)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v204 < v206 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v247 = F_SPI_getbinval(m, v62, v60, v246, v108)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L82
	}
L68:
	;
	v246 = v212 + int32(1)
	goto L67
L69:
	;
	v211 = v206
	v212 = v204
	goto L72
L70:
	;
	goto L71
L71:
	;
	v235 = F_SystemAttributeByName(m, v203)
	mBase = m.M
	if v235 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L72:
	;
	v218 = v60 + v211<<(uint(int32(4))%32) + v212*int32(100)
	v221 = F_namestrcmp(m, v218+int32(24), v203)
	mBase = m.M
	if v221 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L71
L74:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+111)))
	if v224 != int32(1) {
		goto L68
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v228 = v212 + int32(1)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v228 < v229 {
		v211 = v229
		v212 = v228
		goto L72
	} else {
		goto L78
	}
L77:
	;
	goto L76
L78:
	;
	goto L73
L79:
	;
	v246 = int32(-9)
	goto L67
L80:
	;
	goto L81
L81:
	;
	v239 = int32(*(*int16)(unsafe.Add(mBase, uint32(v235)+74)))
	v246 = v239
	goto L67
L82:
	;
	v249 = int32(_a_F_pg_get_ruledef_worker_5)
	v250 = int32(0)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v250 < v252 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v293 = F_SPI_getvalue(m, v62, v60, v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L98
	}
L84:
	;
	v292 = v258 + int32(1)
	goto L83
L85:
	;
	v257 = v252
	v258 = v250
	goto L88
L86:
	;
	goto L87
L87:
	;
	v281 = F_SystemAttributeByName(m, v249)
	mBase = m.M
	if v281 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L88:
	;
	v264 = v60 + v257<<(uint(int32(4))%32) + v258*int32(100)
	v267 = F_namestrcmp(m, v264+int32(24), v249)
	mBase = m.M
	if v267 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L87
L90:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+111)))
	if v270 != int32(1) {
		goto L84
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v274 = v258 + int32(1)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v274 < v275 {
		v257 = v275
		v258 = v274
		goto L88
	} else {
		goto L94
	}
L93:
	;
	goto L92
L94:
	;
	goto L89
L95:
	;
	v292 = int32(-9)
	goto L83
L96:
	;
	goto L97
L97:
	;
	v285 = int32(*(*int16)(unsafe.Add(mBase, uint32(v281)+74)))
	v292 = v285
	goto L83
L98:
	;
	v295 = int32(_a_F_pg_get_ruledef_worker_6)
	v296 = int32(0)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v296 < v298 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v339 = F_SPI_getvalue(m, v62, v60, v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L114
	}
L100:
	;
	v338 = v304 + int32(1)
	goto L99
L101:
	;
	v303 = v298
	v304 = v296
	goto L104
L102:
	;
	goto L103
L103:
	;
	v327 = F_SystemAttributeByName(m, v295)
	mBase = m.M
	if v327 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L104:
	;
	v310 = v60 + v303<<(uint(int32(4))%32) + v304*int32(100)
	v313 = F_namestrcmp(m, v310+int32(24), v295)
	mBase = m.M
	if v313 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L103
L106:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+111)))
	if v316 != int32(1) {
		goto L100
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v320 = v304 + int32(1)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v320 < v321 {
		v303 = v321
		v304 = v320
		goto L104
	} else {
		goto L110
	}
L109:
	;
	goto L108
L110:
	;
	goto L105
L111:
	;
	v338 = int32(-9)
	goto L99
L112:
	;
	goto L113
L113:
	;
	v331 = int32(*(*int16)(unsafe.Add(mBase, uint32(v327)+74)))
	v338 = v331
	goto L99
L114:
	;
	v341 = F_stringToNode(m, v339)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	if v341 == int32(0) {
		goto L5
	} else {
		goto L116
	}
L116:
	;
	v346 = F_table_open(m, v201, int32(1))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v348 = F_quote_identifier(m, v109)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v348
	v352 = v13 + int32(72)
	F_appendStringInfo(m, v352, int32(_a_F_pg_get_ruledef_worker_7), v13+int32(48))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v361 = l1 & int32(2)
	if v361 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v362 = int32(_a_F_pg_get_ruledef_worker_8)
	goto L122
L121:
	;
	v362 = int32(_a_F_pg_get_ruledef_worker_9)
	goto L122
L122:
	;
	F_appendStringInfoString(m, v352, v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	switch v155&int32(255) - int32(49) {
	case 0:
		goto L125
	case 1:
		goto L129
	case 2:
		goto L128
	case 3:
		goto L127
	default:
		goto L126
	}
L124:
	;
	if base.Ui32(int32(4)) <= base.Ui32(l1) {
		goto L139
	} else {
		goto L140
	}
L125:
	;
	F_appendStringInfoString(m, v13+int32(72), int32(_a_F_pg_get_ruledef_worker_10))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L137
	}
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L133
	}
L127:
	;
	F_appendStringInfoString(m, v13+int32(72), int32(_a_F_pg_get_ruledef_worker_11))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L132
	}
L128:
	;
	F_appendStringInfoString(m, v13+int32(72), int32(_a_F_pg_get_ruledef_worker_12))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	F_appendStringInfoString(m, v13+int32(72), int32(_a_F_pg_get_ruledef_worker_13))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v413 = int32(0)
	goto L124
L131:
	;
	v413 = int32(0)
	goto L124
L132:
	;
	v413 = int32(0)
	goto L124
L133:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = base.I32_extend8_s(v155)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v109
	F_errmsg(m, int32(_a_F_pg_get_ruledef_worker_14), v13+int32(16))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_pg_get_ruledef_worker_15), int32(_a_F_pg_get_ruledef_worker_16), int32(_a_F_pg_get_ruledef_worker_17))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v346)+52))
	v413 = v412
	goto L124
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v421
	F_appendStringInfo(m, v13+int32(72), int32(_a_F_pg_get_ruledef_worker_18), v13+int32(32))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L144
	}
L139:
	;
	v417 = F_generate_relation_name(m, v201, int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v419 = F_generate_qualified_relation_name(m, v201)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L143
	}
L142:
	;
	v421 = v417
	goto L138
L143:
	;
	v421 = v419
	goto L138
L144:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	if v430 != int32(60) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v505 = v13 + int32(72)
	F_appendStringInfoString(m, v505, int32(_a_F_pg_get_ruledef_worker_19))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L164
	}
L146:
	;
	if v361 != 0 {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+1)))
	if v433 != int32(62) {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+2)))
	if v436 == int32(0) {
		goto L145
	} else {
		goto L149
	}
L149:
	;
	goto L146
L150:
	;
	F_appendStringInfoString(m, v13+int32(72), int32(_a_F_pg_get_ruledef_worker_20))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v445 = v13 + int32(72)
	F_appendStringInfoString(m, v445, int32(_a_F_pg_get_ruledef_worker_21))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L154
	}
L153:
	;
	goto L152
L154:
	;
	v449 = F_stringToNode(m, v293)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v341)+12))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	v454 = F_getInsertSelectQuery(m, v452, int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	v456 = int32(0)
	F_AcquireRewriteLocks(m, v454, v456, v456)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v461 = v13 + int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v461
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v461
	v468 = F_list_make1_impl(m, int32(1), v13+int32(28))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+196)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+188)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v468
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v454)+52))
	if v476 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	v480 = base.B2i32(v477 != int32(1))
	goto L161
L160:
	;
	v480 = int32(1)
	goto L161
L161:
	;
	v481 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+216)) = v481
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+215)) = uint8(v481)
	v485 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+213)) = uint16(v485)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+204)) = int64(34359738368)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+212)) = uint8(v480)
	F_set_deparse_for_query(m, v13+int32(100), v454, v481)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_get_rule_expr(m, v449, v13+int32(180), int32(0))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	goto L145
L164:
	;
	if v247 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	F_appendStringInfoString(m, v505, int32(_a_F_pg_get_ruledef_worker_22))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	if int32(2) <= v512 {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	goto L167
L169:
	;
	F_relation_close(m, v346, int32(1))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L185
	}
L170:
	;
	F_appendStringInfoChar(m, v13+int32(72), int32(40))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v341)+12))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	v572 = v13 + int32(72)
	v573 = int32(0)
	F_get_query_def(m, v570, v572, v573, v413, int32(1), l1, v573, v573)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L183
	}
L173:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	if int32(0) < v520 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v527 = int32(0)
	goto L177
L175:
	;
	goto L176
L176:
	;
	F_appendStringInfoString(m, v13+int32(72), int32(_a_F_pg_get_ruledef_worker_23))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L182
	}
L177:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v341)+12))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v534+v527<<(uint(int32(2))%32))))
	v540 = v13 + int32(72)
	v541 = int32(0)
	F_get_query_def(m, v538, v540, v541, v413, int32(1), l1, v541, v541)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L179
	}
L178:
	;
	goto L176
L179:
	;
	F_appendStringInfoString(m, v540, int32(_a_F_pg_get_ruledef_worker_24))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v551 = v527 + int32(1)
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	if v551 < v552 {
		v527 = v551
		goto L177
	} else {
		goto L181
	}
L181:
	;
	goto L178
L182:
	;
	goto L169
L183:
	;
	F_appendStringInfoChar(m, v572, int32(59))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	goto L169
L185:
	;
	goto L18
L186:
	;
	if v605 != int32(2) {
		goto L4
	} else {
		goto L187
	}
L187:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	m.G0 = v13 + int32(224)
	if v609 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v615 = v610
	goto L190
L189:
	;
	v615 = int32(0)
	goto L190
L190:
	;
	return v615
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_pg_get_ruledef_worker_0)
	F_errmsg_internal(m, int32(_a_F_pg_get_ruledef_worker_25), v13)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_pg_get_ruledef_worker_15), int32(629), int32(_a_F_pg_get_ruledef_worker_26))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = l0
	F_errmsg_internal(m, int32(_a_F_pg_get_ruledef_worker_27), v13-int32(-64))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(_a_F_pg_get_ruledef_worker_15), int32(641), int32(_a_F_pg_get_ruledef_worker_26))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	F_errmsg_internal(m, int32(_a_F_pg_get_ruledef_worker_28), int32(0))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_pg_get_ruledef_worker_15), int32(_a_F_pg_get_ruledef_worker_29), int32(_a_F_pg_get_ruledef_worker_17))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L200:
	;
	F_errmsg_internal(m, int32(_a_F_pg_get_ruledef_worker_30), int32(0))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(_a_F_pg_get_ruledef_worker_15), int32(663), int32(_a_F_pg_get_ruledef_worker_26))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_serial_sequence(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(160)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = F_pg_detoast_datum_packed(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = F_textToQualifiedNameList(m, v13)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = F_makeRangeVarFromNameList(m, v20)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v24 = int32(0)
	v28 = F_RangeVarGetRelidExtended(m, v22, v24, v24, v24, v24)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v30 = F_text_to_cstring(m, v18)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v32 = F_get_attnum(m, v28, v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v32 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v36 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L39
	}
L12:
	;
	v39 = v10 + int32(16)
	F_ScanKeyInit(m, v39, int32(4), int32(3), int32(184), int32(1259))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_ScanKeyInit(m, v10-int32(-64), int32(5), int32(3), int32(184), v28)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_ScanKeyInit(m, v10+int32(112), int32(6), int32(3), int32(65), v32)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v64 = F_systable_beginscan(m, v36, int32(2674), int32(1), int32(0), int32(3), v39)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	F_systable_endscan(m, v64)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L30
	}
L17:
	;
	v66 = F_systable_getnext(m, v64)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v66 == int32(0) {
		v100 = v2
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v71 = v66
	goto L20
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+22)))
	v79 = v77 + v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v80 != int32(1259) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v100 = v2
	goto L16
L22:
	;
	v93 = F_systable_getnext(m, v64)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L28
	}
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	if v83 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+24)))
	switch v84 - int32(97) {
	case 0, 8:
		goto L25
	default:
		goto L22
	}
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v88 = F_get_rel_relkind(m, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v88 != int32(83) {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v100 = v92
	goto L16
L28:
	;
	if v93 != 0 {
		v71 = v93
		goto L20
	} else {
		goto L29
	}
L29:
	;
	goto L21
L30:
	;
	F_relation_close(m, v36, int32(1))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v100 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	m.G0 = v10 + int32(160)
	return v117
L33:
	;
	v107 = F_generate_qualified_relation_name(m, v100)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v113 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v113)
	v117 = int32(0)
	goto L32
L36:
	;
	v109 = F_cstring_to_text(m, v107)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_pfree(m, v107)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v117 = v109
	goto L32
L39:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v30
	F_errmsg(m, int32(_a_F_pg_get_serial_sequence_0), v10)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_pg_get_serial_sequence_1), int32(2859), int32(_a_F_pg_get_serial_sequence_2))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_viewdef(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_get_viewdef_worker(m, v3, int32(2), int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v12 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
			return int32(0)
		} else {
			v16 = F_cstring_to_text(m, v6)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v6)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					return v16
				}
			}
		}
	}
}
func F_pg_get_viewdef_ext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
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
	var v22 int32
	_ = v22
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v6 != 0 {
		v7 = int32(7)
	} else {
		v7 = int32(2)
	}
	v9 = F_pg_get_viewdef_worker(m, v3, v7, int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			v15 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v15)
			return int32(0)
		} else {
			v19 = F_cstring_to_text(m, v9)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v9)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					return v19
				}
			}
		}
	}
}
func F_pg_get_viewdef_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_textToQualifiedNameList(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = F_makeRangeVarFromNameList(m, v8)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				v12 = int32(0)
				v16 = F_RangeVarGetRelidExtended(m, v10, v12, v12, v12, v12)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v20 = F_pg_get_viewdef_worker(m, v16, int32(2), int32(0))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						if v20 == int32(0) {
							v24 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
							return int32(0)
						} else {
							v28 = F_cstring_to_text(m, v20)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v20)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return int32(0)
								} else {
									return v28
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_getnameinfo_all(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	if l0 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v317
L2:
	;
	if l4 != 0 {
		goto L60
	} else {
		goto L61
	}
L3:
	;
	v76 = int32(_a_F_pg_getnameinfo_all_0)
	if l3 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L4:
	;
	if l2 == int32(0) {
		v197 = v70
		goto L2
	} else {
		goto L28
	}
L5:
	;
	v317 = int32(0)
	goto L1
L6:
	;
	v64 = m.Env.Getnameinfo(m, l0, l1, l2, l3, l4, l5, l6)
	mBase = m.M
	if v64 != 0 {
		v70 = v64
		goto L4
	} else {
		goto L27
	}
L7:
	;
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v15 != int32(1) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if l2|l4 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v317 = int32(-4)
	goto L1
L10:
	;
	goto L11
L11:
	;
	if l2 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(_a_F_pg_getnameinfo_all_1)
	v28 = F_pg_snprintf(m, l2, l3, int32(_a_F_pg_getnameinfo_all_2), v11+int32(32))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if l4 == int32(0) {
		goto L5
	} else {
		goto L18
	}
L15:
	;
	return int32(0)
L16:
	;
	if base.B2i32(v28 < int32(0))|base.B2i32(l3 <= v28) != 0 {
		v75 = int32(-10)
		goto L3
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if v40 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v59 = int32(0)
	v60 = int32(-10)
	if v58 < v59 {
		v70 = v60
		goto L4
	} else {
		goto L25
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0 + int32(2)
	v56 = F_pg_snprintf(m, l4, l5, int32(_a_F_pg_getnameinfo_all_2), v11)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L15
	} else {
		goto L24
	}
L21:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
	if v41 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l0 + int32(3)
	v50 = F_pg_snprintf(m, l4, l5, int32(_a_F_pg_getnameinfo_all_3), v11+int32(16))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v58 = v50
	goto L19
L24:
	;
	v58 = v56
	goto L19
L25:
	;
	if v58 < l5 {
		v317 = v59
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v70 = v60
	goto L4
L27:
	;
	goto L5
L28:
	;
	v75 = v70
	goto L3
L29:
	;
	v197 = v75
	goto L2
L30:
	;
	v192 = F_strlen(m, v188)
	mBase = m.M
	goto L29
L31:
	;
	v188 = v76
	goto L30
L32:
	;
	goto L33
L33:
	;
	v82 = l3 - int32(1)
	if (l2^v76)&int32(3) != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v185 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v182))) = uint8(v185)
	v188 = v181
	goto L30
L35:
	;
	v166 = v161
	v167 = v162
	v168 = v163
	goto L56
L36:
	;
	if v156 == int32(0) {
		v181 = v154
		v182 = v155
		goto L34
	} else {
		goto L55
	}
L37:
	;
	v154 = v76
	v155 = l2
	v156 = v82
	goto L36
L38:
	;
	goto L39
L39:
	;
	v86 = int32(0)
	if int32(0)|base.B2i32(v82 == v86) == v86 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v122 == int32(0) {
		v181 = v119
		v182 = v120
		goto L34
	} else {
		goto L49
	}
L41:
	;
	v98 = v76
	v99 = l2
	v100 = v82
	goto L44
L42:
	;
	goto L43
L43:
	;
	v119 = v76
	v120 = l2
	v121 = v82
	v122 = base.B2i32(v82 != v86)
	goto L40
L44:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v102)
	if v102 == int32(0) {
		v161 = v98
		v162 = v99
		v163 = v100
		goto L35
	} else {
		goto L46
	}
L45:
	;
	v119 = v113
	v120 = v107
	v121 = v109
	v122 = v111
	goto L40
L46:
	;
	v106 = int32(1)
	v107 = v99 + v106
	v109 = v100 - v106
	v110 = int32(0)
	v111 = base.B2i32(v109 != v110)
	v113 = v98 + v106
	if v113&int32(3) == v110 {
		v119 = v113
		v120 = v107
		v121 = v109
		v122 = v111
		goto L40
	} else {
		goto L47
	}
L47:
	;
	if v109 != 0 {
		v98 = v113
		v99 = v107
		v100 = v109
		goto L44
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if base.B2i32(v125 == int32(0))|base.B2i32(base.Ui32(v121) < base.Ui32(int32(4))) != 0 {
		v154 = v119
		v155 = v120
		v156 = v121
		goto L36
	} else {
		goto L50
	}
L50:
	;
	v132 = v119
	v133 = v120
	v134 = v121
	goto L51
L51:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v140 = int32(-2139062144)
	if (int32(16843008)-v137|v137)&v140 != v140 {
		v161 = v132
		v162 = v133
		v163 = v134
		goto L35
	} else {
		goto L53
	}
L52:
	;
	v154 = v148
	v155 = v146
	v156 = v150
	goto L36
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v137
	v145 = int32(4)
	v146 = v133 + v145
	v148 = v132 + v145
	v150 = v134 - v145
	if base.Ui32(int32(3)) < base.Ui32(v150) {
		v132 = v148
		v133 = v146
		v134 = v150
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v161 = v154
	v162 = v155
	v163 = v156
	goto L35
L56:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	*(*uint8)(unsafe.Add(mBase, uint32(v167))) = uint8(v170)
	if v170 == int32(0) {
		v181 = v166
		v182 = v167
		goto L34
	} else {
		goto L58
	}
L57:
	;
	v181 = v177
	v182 = v175
	goto L34
L58:
	;
	v174 = int32(1)
	v175 = v167 + v174
	v177 = v166 + v174
	v179 = v168 - v174
	if v179 != 0 {
		v166 = v177
		v167 = v175
		v168 = v179
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v198 = int32(_a_F_pg_getnameinfo_all_0)
	if l5 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	goto L62
L62:
	;
	v317 = v197
	goto L1
L63:
	;
	goto L62
L64:
	;
	v314 = F_strlen(m, v310)
	mBase = m.M
	goto L63
L65:
	;
	v310 = v198
	goto L64
L66:
	;
	goto L67
L67:
	;
	v204 = l5 - int32(1)
	if (l4^v198)&int32(3) != 0 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v307 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v304))) = uint8(v307)
	v310 = v303
	goto L64
L69:
	;
	v288 = v283
	v289 = v284
	v290 = v285
	goto L90
L70:
	;
	if v278 == int32(0) {
		v303 = v276
		v304 = v277
		goto L68
	} else {
		goto L89
	}
L71:
	;
	v276 = v198
	v277 = l4
	v278 = v204
	goto L70
L72:
	;
	goto L73
L73:
	;
	v208 = int32(0)
	if int32(0)|base.B2i32(v204 == v208) == v208 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	if v244 == int32(0) {
		v303 = v241
		v304 = v242
		goto L68
	} else {
		goto L83
	}
L75:
	;
	v220 = v198
	v221 = l4
	v222 = v204
	goto L78
L76:
	;
	goto L77
L77:
	;
	v241 = v198
	v242 = l4
	v243 = v204
	v244 = base.B2i32(v204 != v208)
	goto L74
L78:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	*(*uint8)(unsafe.Add(mBase, uint32(v221))) = uint8(v224)
	if v224 == int32(0) {
		v283 = v220
		v284 = v221
		v285 = v222
		goto L69
	} else {
		goto L80
	}
L79:
	;
	v241 = v235
	v242 = v229
	v243 = v231
	v244 = v233
	goto L74
L80:
	;
	v228 = int32(1)
	v229 = v221 + v228
	v231 = v222 - v228
	v232 = int32(0)
	v233 = base.B2i32(v231 != v232)
	v235 = v220 + v228
	if v235&int32(3) == v232 {
		v241 = v235
		v242 = v229
		v243 = v231
		v244 = v233
		goto L74
	} else {
		goto L81
	}
L81:
	;
	if v231 != 0 {
		v220 = v235
		v221 = v229
		v222 = v231
		goto L78
	} else {
		goto L82
	}
L82:
	;
	goto L79
L83:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
	if base.B2i32(v247 == int32(0))|base.B2i32(base.Ui32(v243) < base.Ui32(int32(4))) != 0 {
		v276 = v241
		v277 = v242
		v278 = v243
		goto L70
	} else {
		goto L84
	}
L84:
	;
	v254 = v241
	v255 = v242
	v256 = v243
	goto L85
L85:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	v262 = int32(-2139062144)
	if (int32(16843008)-v259|v259)&v262 != v262 {
		v283 = v254
		v284 = v255
		v285 = v256
		goto L69
	} else {
		goto L87
	}
L86:
	;
	v276 = v270
	v277 = v268
	v278 = v272
	goto L70
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = v259
	v267 = int32(4)
	v268 = v255 + v267
	v270 = v254 + v267
	v272 = v256 - v267
	if base.Ui32(int32(3)) < base.Ui32(v272) {
		v254 = v270
		v255 = v268
		v256 = v272
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v283 = v276
	v284 = v277
	v285 = v278
	goto L69
L90:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	*(*uint8)(unsafe.Add(mBase, uint32(v289))) = uint8(v292)
	if v292 == int32(0) {
		v303 = v288
		v304 = v289
		goto L68
	} else {
		goto L92
	}
L91:
	;
	v303 = v299
	v304 = v297
	goto L68
L92:
	;
	v296 = int32(1)
	v297 = v289 + v296
	v299 = v288 + v296
	v301 = v290 - v296
	if v301 != 0 {
		v288 = v299
		v289 = v297
		v290 = v301
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
}
func F_pg_has_role_id_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v21 = F_GetSysCacheOid(m, int32(10), v11, v18, v18, v18)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			if v21 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
						F_errmsg(m, int32(_a_F_pg_has_role_id_name_0), v8)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_has_role_id_name_1), int32(_a_F_pg_has_role_id_name_2), int32(_a_F_pg_has_role_id_name_3))
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
					}
				}
			} else {
				v42 = F_convert_any_priv_string(m, v13, int32(_a_F_pg_has_role_id_name_4))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					v44 = F_pg_role_aclcheck(m, v21, v10, v42)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return v44 ^ int32(1)
					}
				}
			}
		}
	}
}
func F_pg_has_role_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_pg_has_role_name[0]))
		v19 = int32(0)
		v22 = F_GetSysCacheOid(m, int32(10), v10, v19, v19, v19)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if v22 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
						F_errmsg(m, int32(_a_F_pg_has_role_name_0), v8)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_has_role_name_1), int32(_a_F_pg_has_role_name_2), int32(_a_F_pg_has_role_name_3))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v43 = F_convert_any_priv_string(m, v12, int32(_a_F_pg_has_role_name_4))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v45 = F_pg_role_aclcheck(m, v22, v17, v43)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return v45 ^ int32(1)
					}
				}
			}
		}
	}
}
func F_pg_has_role_name_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v21 = F_GetSysCacheOid(m, int32(10), v11, v18, v18, v18)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			if v21 != 0 {
				v24 = int32(0)
				v27 = F_GetSysCacheOid(m, int32(10), v10, v24, v24, v24)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
								F_errmsg(m, int32(_a_F_pg_has_role_name_name_0), v8+int32(16))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_has_role_name_name_1), int32(_a_F_pg_has_role_name_name_2), int32(_a_F_pg_has_role_name_name_3))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v32 = F_convert_any_priv_string(m, v13, int32(_a_F_pg_has_role_name_name_4))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = F_pg_role_aclcheck(m, v27, v21, v32)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(32)
								return v34 ^ int32(1)
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
						F_errmsg(m, int32(_a_F_pg_has_role_name_name_0), v8)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_has_role_name_name_1), int32(_a_F_pg_has_role_name_name_2), int32(_a_F_pg_has_role_name_name_3))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
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
func F_pg_hmac(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = int32(1)
	v21 = v16 + v20
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v26 = v24 & v20
	if v26 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v27 = v21
	goto L5
L4:
	;
	v27 = v16 + int32(4)
	goto L5
L5:
	;
	if v24 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v56 = F_downcase_truncate_identifier(m, v27, v54, int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v33 == int32(18) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v44 = int32(1)
	if v26 != 0 {
		v54 = int32(base.Ui32(v24)>>(uint(v44)%32)) - v44
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v36 = int32(16)
	goto L12
L11:
	;
	v36 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v33-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v43 = int32(4)
	goto L15
L14:
	;
	v43 = v36
	goto L15
L15:
	;
	v54 = v43
	goto L6
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v60 = F_px_find_hmac(m, v56, v13+int32(12))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v60 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_pfree(m, v56)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L71
	}
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v68 = m.T0[v67].(func(*base.Module, int32) int32)(m, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v71 = v68 + int32(4)
	v72 = F_palloc(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v71 << (uint(int32(2)) % 32)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v78 = F_pg_detoast_datum_packed(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v81 = F_pg_detoast_datum_packed(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v83 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v113 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	if v89 == int32(18) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v100 = int32(1)
	if v83&v100 != 0 {
		v112 = int32(base.Ui32(v83)>>(uint(v100)%32)) - v100
		goto L27
	} else {
		goto L37
	}
L31:
	;
	v92 = int32(16)
	goto L33
L32:
	;
	v92 = int32(0)
	goto L33
L33:
	;
	if base.Ui32((v89-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v99 = int32(4)
	goto L36
L35:
	;
	v99 = v92
	goto L36
L36:
	;
	v112 = v99
	goto L27
L37:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v112 = int32(base.Ui32(v106)>>(uint(int32(2))%32)) - int32(4)
	goto L27
L38:
	;
	v143 = int32(1)
	if v113&v143 != 0 {
		goto L49
	} else {
		goto L50
	}
L39:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	if v119 == int32(18) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v130 = int32(1)
	if v113&v130 != 0 {
		v142 = int32(base.Ui32(v113)>>(uint(v130)%32)) - v130
		goto L38
	} else {
		goto L48
	}
L42:
	;
	v122 = int32(16)
	goto L44
L43:
	;
	v122 = int32(0)
	goto L44
L44:
	;
	if base.Ui32((v119-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v129 = int32(4)
	goto L47
L46:
	;
	v129 = v122
	goto L47
L47:
	;
	v142 = v129
	goto L38
L48:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v142 = int32(base.Ui32(v136)>>(uint(int32(2))%32)) - int32(4)
	goto L38
L49:
	;
	v147 = v143
	goto L51
L50:
	;
	v147 = int32(4)
	goto L51
L51:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v66)+24))
	m.T0[v149].(func(*base.Module, int32, int32, int32))(m, v66, v81+v147, v142)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v152 = int32(1)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v154&v152 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v157 = v152
	goto L55
L54:
	;
	v157 = int32(4)
	goto L55
L55:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	m.T0[v159].(func(*base.Module, int32, int32, int32))(m, v66, v78+v157, v112)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	m.T0[v164].(func(*base.Module, int32, int32))(m, v66, v72+int32(4))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	m.T0[v167].(func(*base.Module, int32))(m, v66)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v170 != v78 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F_pfree(m, v78)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v174 != v81 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	F_pfree(m, v81)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v178 != v16 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L65
L67:
	;
	F_pfree(m, v16)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	m.G0 = v13 + int32(16)
	return v72
L70:
	;
	goto L69
L71:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	if v60 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v56
	F_errmsg(m, int32(_a_F_pg_hmac_0), v13)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L87
	}
L74:
	;
	v222 = int32(_a_F_pg_hmac_1)
	goto L73
L75:
	;
	goto L76
L76:
	;
	v201 = int32(_a_F_pg_hmac_2)
	goto L78
L77:
	;
	v222 = v216
	goto L73
L78:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	if v60 != v204 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	v216 = v214
	goto L77
L80:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v201)+20))
	if v206 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	goto L79
L83:
	;
	v222 = int32(_a_F_pg_hmac_3)
	goto L73
L84:
	;
	goto L85
L85:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v201)+16))
	if v60 != v210 {
		v201 = v201 + int32(16)
		goto L78
	} else {
		goto L86
	}
L86:
	;
	v216 = v206
	goto L77
L87:
	;
	F_errfinish(m, int32(_a_F_pg_hmac_4), int32(513), int32(_a_F_pg_hmac_5))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_ident_file_mappings(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
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
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	v29 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v29
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_pg_ident_file_mappings[0]))
	v36 = F_open_auth_file(m, v32, int32(21), v29, v29)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_pg_ident_file_mappings[0]))
	F_tokenize_auth_file(m, v39, v36, v19+int32(4), int32(12), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_pg_ident_file_mappings[1]))
	v52 = F_AllocSetContextCreateInternal(m, v47, int32(_a_F_pg_ident_file_mappings_0), int32(0), int32(1024), int32(_a_F_pg_ident_file_mappings_1))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v54 = int32(_a_F_pg_ident_file_mappings_2)
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_pg_ident_file_mappings[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_ident_file_mappings[1])) = v52
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v58 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_free_auth_file(m, v36)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L38
	}
L7:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v61 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v65 = v19 + int32(11)
	v76 = v2
	v77 = v2
	goto L9
L9:
	;
	v82 = int32(0)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83+v76<<(uint(int32(2))%32))))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	if v88 == v82 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L6
L11:
	;
	v92 = F_parse_ident_line(m, v87, int32(12))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v95 = v82
	v96 = v88
	goto L13
L13:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v99 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v99
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v99
	*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = v99
	v105 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v19)+11)) = v105
	v112 = v77 + int32(1)
	if v96 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	v95 = v92
	v96 = v94
	goto L13
L15:
	;
	v116 = F_cstring_to_text(m, v98)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L19
	}
L16:
	;
	v113 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+8)) = uint8(v113)
	goto L15
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v112
	goto L15
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v116
	if v95 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v96 != 0 {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	v121 = F_cstring_to_text(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v134 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v65)+2)) = uint8(v134)
	v136 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v65))) = uint16(v136)
	goto L20
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v121
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v126 = F_cstring_to_text(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v126
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v131 = F_cstring_to_text(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v131
	goto L20
L27:
	;
	if v96 != 0 {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	v138 = F_cstring_to_text(m, v96)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v141 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+14)) = uint8(v141)
	goto L27
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v138
	goto L27
L32:
	;
	v143 = v77
	goto L34
L33:
	;
	v143 = v112
	goto L34
L34:
	;
	v148 = F_heap_form_tuple(m, v27, v19+int32(16), v19+int32(8))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_tuplestore_puttuple(m, v28, v148)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v153 = v76 + int32(1)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v153 < v154 {
		v76 = v153
		v77 = v143
		goto L9
	} else {
		goto L37
	}
L37:
	;
	goto L10
L38:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_ident_file_mappings[1])) = v55
	F_MemoryContextDelete(m, v52)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v178 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v178)
	m.G0 = v19 + int32(48)
	return int32(0)
}
func F_pg_index_has_property(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_text_to_cstring(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = int32(0)
			v13 = F_indexam_property(m, l0, v9, v11, v3, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				return v13
			}
		}
	}
}
func F_pg_last_wal_replay_lsn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v4 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int64(0) {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
			return int32(0)
		} else {
			v14 = F_Int64GetDatum(m, v4)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_mb_radix_conv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	v7 = int32(0)
	switch l1 - int32(1) {
	case 0:
		v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
		if base.Ui32(l5) < base.Ui32(v215) {
			v245 = v7
			return v245
		} else {
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
			if base.Ui32(v217) < base.Ui32(l5) {
				v245 = v7
				return v245
			} else {
				v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v219 != 0 {
					v221 = int32(2)
					v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v228 = *(*int32)(unsafe.Add(mBase, uint32(v219+(l5-v215)<<(uint(v221)%32)+v224<<(uint(v221)%32))))
					return v228
				} else {
					v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v232 = int32(1)
					v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v230+(l5-v215)<<(uint(v232)%32)+v235<<(uint(v232)%32)))))
					v245 = v239
					return v245
				}
			}
		}
	case 1:
		v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
		if base.Ui32(l4) < base.Ui32(v168) {
			v245 = v7
			return v245
		} else {
			v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
			if base.Ui32(v170) < base.Ui32(l4) {
				v245 = v7
				return v245
			} else {
				v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+22)))
				if base.Ui32(l5) < base.Ui32(v172) {
					v245 = v7
					return v245
				} else {
					v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+23)))
					if base.Ui32(v174) < base.Ui32(l5) {
						v245 = v7
						return v245
					} else {
						v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v177 != 0 {
							v179 = int32(2)
							v189 = *(*int32)(unsafe.Add(mBase, uint32(v177+(l4-v168)<<(uint(v179)%32)+v176<<(uint(v179)%32))))
							v193 = *(*int32)(unsafe.Add(mBase, uint32(v177+(l5-v172)<<(uint(v179)%32)+v189<<(uint(v179)%32))))
							return v193
						} else {
							v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v197 = int32(1)
							v209 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195+(l4-v168)<<(uint(v197)%32)+v176&int32(_a_F_pg_mb_radix_conv_0)<<(uint(v197)%32)))))
							v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195+(l5-v172)<<(uint(v197)%32)+v209<<(uint(v197)%32)))))
							return v213
						}
					}
				}
			}
		}
	case 2:
		v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
		if base.Ui32(l3) < base.Ui32(v101) {
			v245 = v7
			return v245
		} else {
			v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
			if base.Ui32(v103) < base.Ui32(l3) {
				v245 = v7
				return v245
			} else {
				v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
				if base.Ui32(l4) < base.Ui32(v105) {
					v245 = v7
					return v245
				} else {
					v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
					if base.Ui32(v107) < base.Ui32(l4) {
						v245 = v7
						return v245
					} else {
						v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
						if base.Ui32(l5) < base.Ui32(v109) {
							v245 = v7
							return v245
						} else {
							v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
							if base.Ui32(v111) < base.Ui32(l5) {
								v245 = v7
								return v245
							} else {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v114 != 0 {
									v116 = int32(2)
									v130 = *(*int32)(unsafe.Add(mBase, uint32(v114+(l3-v101)<<(uint(v116)%32)+v113<<(uint(v116)%32))))
									v134 = *(*int32)(unsafe.Add(mBase, uint32(v114+(l4-v105)<<(uint(v116)%32)+v130<<(uint(v116)%32))))
									v138 = *(*int32)(unsafe.Add(mBase, uint32(v114+(l5-v109)<<(uint(v116)%32)+v134<<(uint(v116)%32))))
									return v138
								} else {
									v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v142 = int32(1)
									v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140+(l3-v101)<<(uint(v142)%32)+v113&int32(_a_F_pg_mb_radix_conv_0)<<(uint(v142)%32)))))
									v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140+(l4-v105)<<(uint(v142)%32)+v158<<(uint(v142)%32)))))
									v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140+(l5-v109)<<(uint(v142)%32)+v162<<(uint(v142)%32)))))
									return v166
								}
							}
						}
					}
				}
			}
		}
	case 3:
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
		if base.Ui32(l2) < base.Ui32(v14) {
			v245 = v7
			return v245
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+41)))
			if base.Ui32(v16) < base.Ui32(l2) {
				v245 = v7
				return v245
			} else {
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+42)))
				if base.Ui32(l3) < base.Ui32(v18) {
					v245 = v7
					return v245
				} else {
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)))
					if base.Ui32(v20) < base.Ui32(l3) {
						v245 = v7
						return v245
					} else {
						v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+44)))
						if base.Ui32(l4) < base.Ui32(v22) {
							v245 = v7
							return v245
						} else {
							v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+45)))
							if base.Ui32(v24) < base.Ui32(l4) {
								v245 = v7
								return v245
							} else {
								v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+46)))
								if base.Ui32(l5) < base.Ui32(v26) {
									v245 = v7
									return v245
								} else {
									v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+47)))
									if base.Ui32(v28) < base.Ui32(l5) {
										v245 = v7
										return v245
									} else {
										v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										if v31 != 0 {
											v33 = int32(2)
											v51 = *(*int32)(unsafe.Add(mBase, uint32(v31+(l2-v14)<<(uint(v33)%32)+v30<<(uint(v33)%32))))
											v55 = *(*int32)(unsafe.Add(mBase, uint32(v31+(l3-v18)<<(uint(v33)%32)+v51<<(uint(v33)%32))))
											v59 = *(*int32)(unsafe.Add(mBase, uint32(v31+(l4-v22)<<(uint(v33)%32)+v55<<(uint(v33)%32))))
											v63 = *(*int32)(unsafe.Add(mBase, uint32(v31+(l5-v26)<<(uint(v33)%32)+v59<<(uint(v33)%32))))
											return v63
										} else {
											v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v67 = int32(1)
											v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65+(l2-v14)<<(uint(v67)%32)+v30&int32(_a_F_pg_mb_radix_conv_0)<<(uint(v67)%32)))))
											v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65+(l3-v18)<<(uint(v67)%32)+v87<<(uint(v67)%32)))))
											v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65+(l4-v22)<<(uint(v67)%32)+v91<<(uint(v67)%32)))))
											v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65+(l5-v26)<<(uint(v67)%32)+v95<<(uint(v67)%32)))))
											return v99
										}
									}
								}
							}
						}
					}
				}
			}
		}
	default:
		v245 = v7
		return v245
	}
}
func F_pg_mbstrlen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pg_mbstrlen[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7*int32(28))+uint32(_c_F_pg_mbstrlen[1])))
	if v12 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_report_invalid_encoding_db(m, v16, v29, v36)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L20
	}
L2:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v59 = F_strlen(m, l0)
	mBase = m.M
	return v59
L5:
	;
	v16 = l0
	v19 = v2
	goto L8
L6:
	;
	v57 = v2
	goto L7
L7:
	;
	return v57
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_pg_mbstrlen[0]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23*int32(28))+uint32(_c_F_pg_mbstrlen[2])))
	v29 = m.T0[v28].(func(*base.Module, int32) int32)(m, v16)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v57 = v51
	goto L7
L10:
	;
	return int32(0)
L11:
	;
	if int32(2) <= v29 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v36 = int32(1)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v51 = v19 + int32(1)
	v52 = v16 + v29
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v53 != 0 {
		v16 = v52
		v19 = v51
		goto L8
	} else {
		goto L19
	}
L15:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v36))))
	if v40 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v44 = v36 + int32(1)
	if v44 != v29 {
		v36 = v44
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	goto L9
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_mcv_list_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_pg_mcv_list_in_0), int32(1480), int32(_a_F_pg_mcv_list_in_1), int32(_a_F_pg_mcv_list_in_2), int32(_a_F_pg_mcv_list_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_pg_mule_verifychar(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32((v6+int32(127))&int32(255)) < base.Ui32(int32(13)) {
		v27 = int32(2)
	} else {
		if base.Ui32((v6+int32(112))&int32(255)) < base.Ui32(int32(12)) {
			v27 = int32(3)
		} else {
			if v6&int32(254) == int32(156) {
				v26 = int32(4)
			} else {
				v26 = int32(1)
			}
			v27 = v26
		}
	}
	if l1 < v27 {
		v51 = int32(-1)
		return v51
	} else {
		if base.Ui32(v27) < base.Ui32(int32(2)) {
			return v27
		} else {
			v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
			if int32(0) <= v33 {
				v51 = int32(-1)
				return v51
			} else {
				if v27 == int32(2) {
					return v27
				} else {
					v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
					if int32(0) <= v40 {
						v51 = int32(-1)
						return v51
					} else {
						if base.Ui32(v27) < base.Ui32(int32(4)) {
							return v27
						} else {
							v47 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+3)))
							if int32(0) <= v47 {
								v50 = int32(-1)
							} else {
								v50 = v27
							}
							v51 = v50
							return v51
						}
					}
				}
			}
		}
	}
}
func F_pg_ndistinct_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 float64
	_ = v43
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = F_statext_ndistinct_deserialize(m, v14)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = v9 + int32(-16)
	F_initStringInfo(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_appendStringInfoChar(m, v21, int32(123))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v33 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	F_appendStringInfoChar(m, v9+int32(-16), int32(125))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L25
	}
L9:
	;
	v40 = v18 + int32(16) + v33<<(uint(int32(4))%32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v43 = *(*float64)(unsafe.Add(mBase, uint32(v40)))
	if int32(0) < v33 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	F_appendStringInfoString(m, v9+int32(-16), int32(_a_F_pg_ndistinct_out_0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v42 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = base.I32_trunc_sat_f64_s(v43)
	F_appendStringInfo(m, v9+int32(-16), int32(_a_F_pg_ndistinct_out_1), v11)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L23
	}
L16:
	;
	v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(_a_F_pg_ndistinct_out_2)
	F_appendStringInfo(m, v9+int32(-16), int32(_a_F_pg_ndistinct_out_3), v9+int32(-32))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v64 = int32(1)
	if v42 == v64 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v67 = v64
	goto L19
L19:
	;
	v78 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41+v67<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_pg_ndistinct_out_0)
	F_appendStringInfo(m, v9+int32(-16), int32(_a_F_pg_ndistinct_out_3), v9+int32(-48))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L15
L21:
	;
	v90 = v67 + int32(1)
	if v90 != v42 {
		v67 = v90
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v108 = v33 + int32(1)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if base.Ui32(v108) < base.Ui32(v109) {
		v33 = v108
		goto L9
	} else {
		goto L24
	}
L24:
	;
	goto L10
L25:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	m.G0 = v11 - int32(-64)
	return v124
}
func F_pg_nextoid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_superuser(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if v16 != 0 {
			v21 = F_table_open(m, v15, int32(3))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v24 = F_index_open(m, v13, int32(3))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
					if base.Ui32(v26) < base.Ui32(int32(_a_F_pg_nextoid_0)) {
						v43 = v26
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v24)+192))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
						if v46 != v43 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int32(0)
								} else {
									v123 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
									v124 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
									v125 = int32(4)
									*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v124 + v125
									*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v123 + v125
									F_errmsg(m, int32(_a_F_pg_nextoid_1), v11+int32(48))
									mBase = m.M
									v135 = m.ExcPending
									if v135 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(678), int32(_a_F_pg_nextoid_3))
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v48 = F_SearchSysCacheAttName(m, v15, v14)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								if v48 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50360452))
										mBase = m.M
										v147 = m.ExcPending
										if v147 != 0 {
											return int32(0)
										} else {
											v148 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
											*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v148 + int32(4)
											F_errmsg(m, int32(_a_F_pg_nextoid_4), v11)
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(685), int32(_a_F_pg_nextoid_3))
												mBase = m.M
												v160 = m.ExcPending
												if v160 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
									v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
									v54 = v52 + v53
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+68))
									if v55 != int32(26) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v14
												F_errmsg(m, int32(_a_F_pg_nextoid_5), v11+int32(32))
												mBase = m.M
												v173 = m.ExcPending
												if v173 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(694), int32(_a_F_pg_nextoid_3))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v24)+192))
										v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+10)))
										if v59 != int32(1) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v183 = m.ExcPending
											if v183 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v186 = m.ExcPending
												if v186 != 0 {
													return int32(0)
												} else {
													v187 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v14
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v187 + int32(4)
													F_errmsg(m, int32(_a_F_pg_nextoid_6), v11+int32(16))
													mBase = m.M
													v196 = m.ExcPending
													if v196 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(702), int32(_a_F_pg_nextoid_3))
														mBase = m.M
														v201 = m.ExcPending
														if v201 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+48)))
											v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+74)))
											if v62 != v63&int32(_a_F_pg_nextoid_7) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v183 = m.ExcPending
												if v183 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v186 = m.ExcPending
													if v186 != 0 {
														return int32(0)
													} else {
														v187 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
														*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v14
														*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v187 + int32(4)
														F_errmsg(m, int32(_a_F_pg_nextoid_6), v11+int32(16))
														mBase = m.M
														v196 = m.ExcPending
														if v196 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(702), int32(_a_F_pg_nextoid_3))
															mBase = m.M
															v201 = m.ExcPending
															if v201 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v67 = F_GetNewOidWithIndex(m, v21, v13, v63)
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v48)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return int32(0)
													} else {
														F_relation_close(m, v21, int32(3))
														mBase = m.M
														v73 = m.ExcPending
														if v73 != 0 {
															return int32(0)
														} else {
															F_relation_close(m, v24, int32(3))
															mBase = m.M
															v76 = m.ExcPending
															if v76 != 0 {
																return int32(0)
															} else {
																m.G0 = v11 + int32(80)
																return v67
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
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
						if v30 == int32(99) {
							v43 = v26
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v24)+192))
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
							if v46 != v43 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return int32(0)
									} else {
										v123 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
										v124 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
										v125 = int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v124 + v125
										*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v123 + v125
										F_errmsg(m, int32(_a_F_pg_nextoid_1), v11+int32(48))
										mBase = m.M
										v135 = m.ExcPending
										if v135 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(678), int32(_a_F_pg_nextoid_3))
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v48 = F_SearchSysCacheAttName(m, v15, v14)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									if v48 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50360452))
											mBase = m.M
											v147 = m.ExcPending
											if v147 != 0 {
												return int32(0)
											} else {
												v148 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
												*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v148 + int32(4)
												F_errmsg(m, int32(_a_F_pg_nextoid_4), v11)
												mBase = m.M
												v155 = m.ExcPending
												if v155 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(685), int32(_a_F_pg_nextoid_3))
													mBase = m.M
													v160 = m.ExcPending
													if v160 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
										v54 = v52 + v53
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+68))
										if v55 != int32(26) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v14
													F_errmsg(m, int32(_a_F_pg_nextoid_5), v11+int32(32))
													mBase = m.M
													v173 = m.ExcPending
													if v173 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(694), int32(_a_F_pg_nextoid_3))
														mBase = m.M
														v178 = m.ExcPending
														if v178 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v24)+192))
											v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+10)))
											if v59 != int32(1) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v183 = m.ExcPending
												if v183 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v186 = m.ExcPending
													if v186 != 0 {
														return int32(0)
													} else {
														v187 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
														*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v14
														*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v187 + int32(4)
														F_errmsg(m, int32(_a_F_pg_nextoid_6), v11+int32(16))
														mBase = m.M
														v196 = m.ExcPending
														if v196 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(702), int32(_a_F_pg_nextoid_3))
															mBase = m.M
															v201 = m.ExcPending
															if v201 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+48)))
												v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+74)))
												if v62 != v63&int32(_a_F_pg_nextoid_7) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v186 = m.ExcPending
														if v186 != 0 {
															return int32(0)
														} else {
															v187 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v14
															*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v187 + int32(4)
															F_errmsg(m, int32(_a_F_pg_nextoid_6), v11+int32(16))
															mBase = m.M
															v196 = m.ExcPending
															if v196 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(702), int32(_a_F_pg_nextoid_3))
																mBase = m.M
																v201 = m.ExcPending
																if v201 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v67 = F_GetNewOidWithIndex(m, v21, v13, v63)
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return int32(0)
													} else {
														F_ReleaseCatCache(m, v48)
														mBase = m.M
														v70 = m.ExcPending
														if v70 != 0 {
															return int32(0)
														} else {
															F_relation_close(m, v21, int32(3))
															mBase = m.M
															v73 = m.ExcPending
															if v73 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v24, int32(3))
																mBase = m.M
																v76 = m.ExcPending
																if v76 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v11 + int32(80)
																	return v67
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
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, _c_F_pg_nextoid[0]))
							if base.B2i32(v35 != int32(0))&base.B2i32(v30 == v35) == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_pg_nextoid_8), int32(0))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(671), int32(_a_F_pg_nextoid_3))
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
								v43 = v42
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v24)+192))
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
								if v46 != v43 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return int32(0)
										} else {
											v123 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
											v124 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
											v125 = int32(4)
											*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v124 + v125
											*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v123 + v125
											F_errmsg(m, int32(_a_F_pg_nextoid_1), v11+int32(48))
											mBase = m.M
											v135 = m.ExcPending
											if v135 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(678), int32(_a_F_pg_nextoid_3))
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v48 = F_SearchSysCacheAttName(m, v15, v14)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										if v48 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v144 = m.ExcPending
											if v144 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50360452))
												mBase = m.M
												v147 = m.ExcPending
												if v147 != 0 {
													return int32(0)
												} else {
													v148 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
													*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v148 + int32(4)
													F_errmsg(m, int32(_a_F_pg_nextoid_4), v11)
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(685), int32(_a_F_pg_nextoid_3))
														mBase = m.M
														v160 = m.ExcPending
														if v160 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
											v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
											v54 = v52 + v53
											v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+68))
											if v55 != int32(26) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v167 = m.ExcPending
													if v167 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v14
														F_errmsg(m, int32(_a_F_pg_nextoid_5), v11+int32(32))
														mBase = m.M
														v173 = m.ExcPending
														if v173 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(694), int32(_a_F_pg_nextoid_3))
															mBase = m.M
															v178 = m.ExcPending
															if v178 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(v24)+192))
												v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+10)))
												if v59 != int32(1) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v186 = m.ExcPending
														if v186 != 0 {
															return int32(0)
														} else {
															v187 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v14
															*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v187 + int32(4)
															F_errmsg(m, int32(_a_F_pg_nextoid_6), v11+int32(16))
															mBase = m.M
															v196 = m.ExcPending
															if v196 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(702), int32(_a_F_pg_nextoid_3))
																mBase = m.M
																v201 = m.ExcPending
																if v201 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+48)))
													v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+74)))
													if v62 != v63&int32(_a_F_pg_nextoid_7) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v183 = m.ExcPending
														if v183 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(50856066))
															mBase = m.M
															v186 = m.ExcPending
															if v186 != 0 {
																return int32(0)
															} else {
																v187 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
																*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v14
																*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v187 + int32(4)
																F_errmsg(m, int32(_a_F_pg_nextoid_6), v11+int32(16))
																mBase = m.M
																v196 = m.ExcPending
																if v196 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(702), int32(_a_F_pg_nextoid_3))
																	mBase = m.M
																	v201 = m.ExcPending
																	if v201 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														}
													} else {
														v67 = F_GetNewOidWithIndex(m, v21, v13, v63)
														mBase = m.M
														v68 = m.ExcPending
														if v68 != 0 {
															return int32(0)
														} else {
															F_ReleaseCatCache(m, v48)
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v21, int32(3))
																mBase = m.M
																v73 = m.ExcPending
																if v73 != 0 {
																	return int32(0)
																} else {
																	F_relation_close(m, v24, int32(3))
																	mBase = m.M
																	v76 = m.ExcPending
																	if v76 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v11 + int32(80)
																		return v67
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
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = int32(_a_F_pg_nextoid_3)
					F_errmsg(m, int32(_a_F_pg_nextoid_9), v11-int32(-64))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(663), int32(_a_F_pg_nextoid_3))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
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
func F_pg_opclass_is_visible(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_OpclassIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v23 = v2
		} else {
			v23 = v14
		}
		m.G0 = v7 + int32(16)
		return v23
	}
}
func F_pg_partition_root(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_SearchSysCacheExists(m, int32(57), v5, v2, v2, v2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 != 0 {
			v13 = F_get_rel_relkind(m, v5)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = F_get_rel_relispartition(m, v5)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					if v15|base.B2i32(v13 == int32(73))|base.B2i32(v13&int32(255) == int32(112)) != 0 {
						v30 = F_get_partition_ancestors(m, v5)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							if v30 != 0 {
								v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
								v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v32+v33<<(uint(int32(2))%32)-int32(4))))
								F_list_free(m, v30)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v42 = v39
									return v42
								}
							} else {
								v42 = v5
								return v42
							}
						}
					} else {
						v26 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
						return int32(0)
					}
				}
			}
		} else {
			v26 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
			return int32(0)
		}
	}
}
func F_pg_pwrite_zeros(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int64
	_ = v78
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
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v137 int32
	_ = v137
	var v153 int32
	_ = v153
	v13 = m.G0
	v15 = v13 - int32(1024)
	m.G0 = v15
	v18 = l1
	v19 = l2
	v27 = int32(0)
	goto L1
L1:
	;
	v29 = int32(0)
	if v18 == v29 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v15 + int32(1024)
	return v153
L3:
	;
	goto L2
L4:
	;
	v153 = v27
	goto L3
L5:
	;
	goto L6
L6:
	;
	v33 = v18
	v35 = v29
	goto L7
L7:
	;
	v46 = v15 + v35<<(uint(int32(3))%32)
	v47 = int32(_a_F_pg_pwrite_zeros_0)
	if base.Ui32(v47) <= base.Ui32(v33) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v61 = m.G0
	v63 = v61 - int32(1024)
	m.G0 = v63
	if v55 <= int32(128) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	goto L8
L10:
	;
	v50 = v47
	goto L12
L11:
	;
	v50 = v33
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = int32(_a_F_pg_pwrite_zeros_1)
	v55 = v35 + int32(1)
	v56 = v33 - v50
	if base.Ui32(int32(126)) < base.Ui32(v35) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	if v56 != 0 {
		v33 = v56
		v35 = v55
		goto L7
	} else {
		goto L14
	}
L14:
	;
	goto L9
L15:
	;
	m.G0 = v63 + int32(1024)
	if int32(0) <= v137 {
		v18 = v56
		v19 = v19 + base.I64_extend_i32_u(v137)
		v27 = v27 + v137
		goto L1
	} else {
		goto L38
	}
L16:
	;
	v70 = v15
	v71 = v55
	v74 = int32(0)
	v78 = v19
	goto L19
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_pwrite_zeros[0])) = int32(28)
	v137 = int32(-1)
	goto L15
L19:
	;
	if v71 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v137 = v89
	goto L15
L21:
	;
	if v85 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v83 = F_pwrite(m, l0, v81, v82, v78)
	mBase = m.M
	v85 = v83
	goto L21
L23:
	;
	goto L24
L24:
	;
	v84 = F_pwritev(m, l0, v70, v71, v78)
	mBase = m.M
	v85 = v84
	goto L21
L25:
	;
	v137 = int32(-1)
	goto L15
L26:
	;
	goto L27
L27:
	;
	v89 = v85 + v74
	v95 = v70
	v96 = v71
	v98 = v85
	goto L28
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if base.Ui32(v104) <= base.Ui32(v98) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v95 == v63 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v110 = v96 - int32(1)
	if v110 != 0 {
		v95 = v95 + int32(8)
		v96 = v110
		v98 = v98 - v104
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	v137 = v89
	goto L15
L34:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v118 + v98
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = v121 - v98
	if int32(0) < v96 {
		v70 = v63
		v71 = v96
		v74 = v89
		v78 = v78 + base.I64_extend_i32_u(v85)
		goto L19
	} else {
		goto L37
	}
L35:
	;
	v113 = v96 << (uint(int32(3)) % 32)
	if v113 == int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	base.MemoryCopy(m, v63, v95, v113)
	goto L34
L37:
	;
	goto L20
L38:
	;
	v153 = v137
	goto L3
}
func F_pg_qsort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
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
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
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
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v422 int32
	_ = v422
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v504 int32
	_ = v504
	var v517 int32
	_ = v517
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v642 int32
	_ = v642
	var v664 int32
	_ = v664
	var v671 int32
	_ = v671
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v723 int32
	_ = v723
	var v737 int32
	_ = v737
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v840 int32
	_ = v840
	var v862 int32
	_ = v862
	var v870 int32
	_ = v870
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v995 int32
	_ = v995
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1078 int32
	_ = v1078
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1148 int32
	_ = v1148
	var v1170 int32
	_ = v1170
	var v1177 int32
	_ = v1177
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1284 int32
	_ = v1284
	var v1293 int32
	_ = v1293
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1317 int32
	_ = v1317
	var v1329 int32
	_ = v1329
	var v1335 int32
	_ = v1335
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1386 int32
	_ = v1386
	var v1408 int32
	_ = v1408
	var v1414 int32
	_ = v1414
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1482 int32
	_ = v1482
	v24 = int32(0) - l2
	if base.Ui32(l1) < base.Ui32(int32(7)) {
		v1231 = l0
		v1232 = l1
		v1233 = l2
		v1234 = l3
		v1251 = v24
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v1253 = v1232 * v1233
	if base.Ui32(v1253) <= base.Ui32(v1233) {
		goto L1
	} else {
		goto L144
	}
L3:
	;
	v33 = l0
	v34 = l1
	v35 = l2
	v36 = l3
	v48 = l2 & int32(-4)
	v49 = l2 & int32(3)
	v50 = l2 - int32(1)
	v53 = v24
	goto L4
L4:
	;
	v55 = v33 + v35
	v57 = v34
	goto L6
L5:
	;
	v1231 = v33
	v1232 = v1228
	v1233 = v35
	v1234 = v36
	v1251 = v53
	goto L2
L6:
	;
	v78 = v57 * v35
	if base.Ui32(v78) <= base.Ui32(v35) {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v80 = v33 + v78
	v87 = v55
	goto L9
L9:
	;
	v104 = m.T0[v36].(func(*base.Module, int32, int32) int32)(m, v87+v53, v87)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v113 = v33 + int32(base.Ui32(v57)>>(uint(int32(1))%32))*v35
	if v57 != int32(7) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	return
L12:
	;
	if v104 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v108 = v35 + v87
	if base.Ui32(v108) < base.Ui32(v80) {
		v87 = v108
		goto L9
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L10
L16:
	;
	goto L1
L17:
	;
	v119 = v33 + (v57-int32(1))*v35
	if base.Ui32(v57) < base.Ui32(int32(41)) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v147 = v113
	goto L19
L19:
	;
	if v35 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	v144 = F_pg_qsort_med3(m, v142, v140, v141, v36)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L11
	} else {
		goto L27
	}
L21:
	;
	v140 = v113
	v141 = v119
	v142 = v33
	goto L20
L22:
	;
	goto L23
L23:
	;
	v124 = int32(base.Ui32(v57)>>(uint(int32(3))%32)) * v35
	v127 = v124 << (uint(int32(1)) % 32)
	v129 = F_pg_qsort_med3(m, v33, v33+v124, v33+v127, v36)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v133 = F_pg_qsort_med3(m, v113-v124, v113, v124+v113, v36)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	v137 = F_pg_qsort_med3(m, v119-v127, v119-v124, v119, v36)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	v140 = v133
	v141 = v137
	v142 = v129
	goto L20
L27:
	;
	v147 = v144
	goto L19
L28:
	;
	v298 = v33 + (v57-int32(1))*v35
	v305 = v298
	v307 = v298
	v308 = v55
	v309 = v55
	goto L40
L29:
	;
	v153 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v50) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v165 = v153
	v169 = v153
	goto L33
L31:
	;
	v222 = v153
	goto L32
L32:
	;
	v244 = v222
	v248 = v153
	goto L37
L33:
	;
	v181 = v33 + v165
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	v183 = v147 + v165
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	*(*uint8)(unsafe.Add(mBase, uint32(v181))) = uint8(v184)
	*(*uint8)(unsafe.Add(mBase, uint32(v183))) = uint8(v182)
	v188 = v165 | int32(1)
	v189 = v33 + v188
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	v191 = v188 + v147
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	*(*uint8)(unsafe.Add(mBase, uint32(v189))) = uint8(v192)
	*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v190)
	v196 = v165 | int32(2)
	v197 = v33 + v196
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	v199 = v196 + v147
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v200)
	*(*uint8)(unsafe.Add(mBase, uint32(v199))) = uint8(v198)
	v204 = v165 | int32(3)
	v205 = v33 + v204
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	v207 = v204 + v147
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	*(*uint8)(unsafe.Add(mBase, uint32(v205))) = uint8(v208)
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v206)
	v211 = int32(4)
	v212 = v165 + v211
	v214 = v169 + v211
	if v214 != v48 {
		v165 = v212
		v169 = v214
		goto L33
	} else {
		goto L35
	}
L34:
	;
	if v49 == int32(0) {
		goto L28
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	v222 = v212
	goto L32
L37:
	;
	v262 = v33 + v244
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	v264 = v244 + v147
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	*(*uint8)(unsafe.Add(mBase, uint32(v262))) = uint8(v265)
	*(*uint8)(unsafe.Add(mBase, uint32(v264))) = uint8(v263)
	v268 = int32(1)
	v271 = v248 + v268
	if v271 != v49 {
		v244 = v244 + v268
		v248 = v271
		goto L37
	} else {
		goto L39
	}
L38:
	;
	goto L28
L39:
	;
	goto L38
L40:
	;
	if base.Ui32(v305) < base.Ui32(v309) {
		v528 = v308
		v529 = v309
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if base.Ui32(v35) < base.Ui32(v913) {
		goto L138
	} else {
		goto L139
	}
L42:
	;
	if base.Ui32(v529) <= base.Ui32(v305) {
		goto L66
	} else {
		goto L67
	}
L43:
	;
	v331 = v308
	v332 = v309
	goto L44
L44:
	;
	v344 = m.T0[v36].(func(*base.Module, int32, int32) int32)(m, v332, v33)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L11
	} else {
		goto L46
	}
L45:
	;
	v528 = v504
	v529 = v517
	goto L42
L46:
	;
	if int32(0) < v344 {
		v528 = v331
		v529 = v332
		goto L42
	} else {
		goto L47
	}
L47:
	;
	if v344 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if v35 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v504 = v331
	goto L50
L50:
	;
	v517 = v35 + v332
	if base.Ui32(v517) <= base.Ui32(v305) {
		v331 = v504
		v332 = v517
		goto L44
	} else {
		goto L63
	}
L51:
	;
	v504 = v35 + v331
	goto L50
L52:
	;
	v352 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v50) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v359 = v352
	v362 = v352
	goto L56
L54:
	;
	v422 = v352
	goto L55
L55:
	;
	v444 = v422
	v451 = v352
	goto L60
L56:
	;
	v380 = v362 + v331
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380))))
	v382 = v362 + v332
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382))))
	*(*uint8)(unsafe.Add(mBase, uint32(v380))) = uint8(v383)
	*(*uint8)(unsafe.Add(mBase, uint32(v382))) = uint8(v381)
	v387 = v362 | int32(1)
	v388 = v331 + v387
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388))))
	v390 = v387 + v332
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390))))
	*(*uint8)(unsafe.Add(mBase, uint32(v388))) = uint8(v391)
	*(*uint8)(unsafe.Add(mBase, uint32(v390))) = uint8(v389)
	v395 = v362 | int32(2)
	v396 = v331 + v395
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396))))
	v398 = v395 + v332
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398))))
	*(*uint8)(unsafe.Add(mBase, uint32(v396))) = uint8(v399)
	*(*uint8)(unsafe.Add(mBase, uint32(v398))) = uint8(v397)
	v403 = v362 | int32(3)
	v404 = v331 + v403
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404))))
	v406 = v403 + v332
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406))))
	*(*uint8)(unsafe.Add(mBase, uint32(v404))) = uint8(v407)
	*(*uint8)(unsafe.Add(mBase, uint32(v406))) = uint8(v405)
	v410 = int32(4)
	v411 = v362 + v410
	v413 = v359 + v410
	if v413 != v48 {
		v359 = v413
		v362 = v411
		goto L56
	} else {
		goto L58
	}
L57:
	;
	if v49 == int32(0) {
		goto L51
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	v422 = v411
	goto L55
L60:
	;
	v461 = v444 + v331
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
	v463 = v444 + v332
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463))))
	*(*uint8)(unsafe.Add(mBase, uint32(v461))) = uint8(v464)
	*(*uint8)(unsafe.Add(mBase, uint32(v463))) = uint8(v462)
	v467 = int32(1)
	v470 = v451 + v467
	if v470 != v49 {
		v444 = v444 + v467
		v451 = v470
		goto L60
	} else {
		goto L62
	}
L61:
	;
	goto L51
L62:
	;
	goto L61
L63:
	;
	goto L45
L64:
	;
	goto L41
L65:
	;
	if v35 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L66:
	;
	v548 = v305
	v550 = v307
	goto L69
L67:
	;
	v745 = v305
	v747 = v307
	goto L68
L68:
	;
	v761 = v528 - v33
	v762 = v529 - v528
	if v761 < v762 {
		goto L90
	} else {
		goto L91
	}
L69:
	;
	v564 = m.T0[v36].(func(*base.Module, int32, int32) int32)(m, v548, v33)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L11
	} else {
		goto L71
	}
L70:
	;
	v745 = v737
	v747 = v723
	goto L68
L71:
	;
	if v564 < int32(0) {
		goto L65
	} else {
		goto L72
	}
L72:
	;
	if v564 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	if v35 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v723 = v550
	goto L75
L75:
	;
	v737 = v548 + v53
	if base.Ui32(v529) <= base.Ui32(v737) {
		v548 = v737
		v550 = v723
		goto L69
	} else {
		goto L88
	}
L76:
	;
	v723 = v550 + v53
	goto L75
L77:
	;
	v572 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v50) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v579 = v572
	v582 = v572
	goto L81
L79:
	;
	v642 = v572
	goto L80
L80:
	;
	v664 = v642
	v671 = v572
	goto L85
L81:
	;
	v600 = v582 + v548
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600))))
	v602 = v582 + v550
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602))))
	*(*uint8)(unsafe.Add(mBase, uint32(v600))) = uint8(v603)
	*(*uint8)(unsafe.Add(mBase, uint32(v602))) = uint8(v601)
	v607 = v582 | int32(1)
	v608 = v548 + v607
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608))))
	v610 = v607 + v550
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610))))
	*(*uint8)(unsafe.Add(mBase, uint32(v608))) = uint8(v611)
	*(*uint8)(unsafe.Add(mBase, uint32(v610))) = uint8(v609)
	v615 = v582 | int32(2)
	v616 = v548 + v615
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v616))))
	v618 = v615 + v550
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618))))
	*(*uint8)(unsafe.Add(mBase, uint32(v616))) = uint8(v619)
	*(*uint8)(unsafe.Add(mBase, uint32(v618))) = uint8(v617)
	v623 = v582 | int32(3)
	v624 = v548 + v623
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624))))
	v626 = v623 + v550
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626))))
	*(*uint8)(unsafe.Add(mBase, uint32(v624))) = uint8(v627)
	*(*uint8)(unsafe.Add(mBase, uint32(v626))) = uint8(v625)
	v630 = int32(4)
	v631 = v582 + v630
	v633 = v579 + v630
	if v633 != v48 {
		v579 = v633
		v582 = v631
		goto L81
	} else {
		goto L83
	}
L82:
	;
	if v49 == int32(0) {
		goto L76
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	v642 = v631
	goto L80
L85:
	;
	v681 = v664 + v548
	v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681))))
	v683 = v664 + v550
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683))))
	*(*uint8)(unsafe.Add(mBase, uint32(v681))) = uint8(v684)
	*(*uint8)(unsafe.Add(mBase, uint32(v683))) = uint8(v682)
	v687 = int32(1)
	v690 = v671 + v687
	if v690 != v49 {
		v664 = v664 + v687
		v671 = v690
		goto L85
	} else {
		goto L87
	}
L86:
	;
	goto L76
L87:
	;
	goto L86
L88:
	;
	goto L70
L89:
	;
	v913 = v747 - v745
	v915 = v80 - (v35 + v747)
	if base.Ui32(v913) < base.Ui32(v915) {
		goto L105
	} else {
		goto L106
	}
L90:
	;
	v764 = v761
	goto L92
L91:
	;
	v764 = v762
	goto L92
L92:
	;
	if v764 == int32(0) {
		goto L89
	} else {
		goto L93
	}
L93:
	;
	v767 = v529 - v764
	v769 = v764 & int32(3)
	v770 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v764) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v778 = int32(0)
	v781 = v770
	goto L97
L95:
	;
	v840 = v770
	goto L96
L96:
	;
	v862 = v840
	v870 = v770
	goto L101
L97:
	;
	v799 = v33 + v781
	v800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v799))))
	v801 = v781 + v767
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801))))
	*(*uint8)(unsafe.Add(mBase, uint32(v799))) = uint8(v802)
	*(*uint8)(unsafe.Add(mBase, uint32(v801))) = uint8(v800)
	v806 = v781 | int32(1)
	v807 = v33 + v806
	v808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807))))
	v809 = v767 + v806
	v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v809))))
	*(*uint8)(unsafe.Add(mBase, uint32(v807))) = uint8(v810)
	*(*uint8)(unsafe.Add(mBase, uint32(v809))) = uint8(v808)
	v814 = v781 | int32(2)
	v815 = v33 + v814
	v816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815))))
	v817 = v767 + v814
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v817))))
	*(*uint8)(unsafe.Add(mBase, uint32(v815))) = uint8(v818)
	*(*uint8)(unsafe.Add(mBase, uint32(v817))) = uint8(v816)
	v822 = v781 | int32(3)
	v823 = v33 + v822
	v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v823))))
	v825 = v767 + v822
	v826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v825))))
	*(*uint8)(unsafe.Add(mBase, uint32(v823))) = uint8(v826)
	*(*uint8)(unsafe.Add(mBase, uint32(v825))) = uint8(v824)
	v829 = int32(4)
	v830 = v781 + v829
	v832 = v778 + v829
	if v832 != v764&int32(-4) {
		v778 = v832
		v781 = v830
		goto L97
	} else {
		goto L99
	}
L98:
	;
	if v769 == int32(0) {
		goto L89
	} else {
		goto L100
	}
L99:
	;
	goto L98
L100:
	;
	v840 = v830
	goto L96
L101:
	;
	v880 = v33 + v862
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v880))))
	v882 = v862 + v767
	v883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v882))))
	*(*uint8)(unsafe.Add(mBase, uint32(v880))) = uint8(v883)
	*(*uint8)(unsafe.Add(mBase, uint32(v882))) = uint8(v881)
	v886 = int32(1)
	v889 = v870 + v886
	if v889 != v769 {
		v862 = v862 + v886
		v870 = v889
		goto L101
	} else {
		goto L103
	}
L102:
	;
	goto L89
L103:
	;
	goto L102
L104:
	;
	if base.Ui32(v913) < base.Ui32(v762) {
		goto L64
	} else {
		goto L119
	}
L105:
	;
	v917 = v913
	goto L107
L106:
	;
	v917 = v915
	goto L107
L107:
	;
	if v917 == int32(0) {
		goto L104
	} else {
		goto L108
	}
L108:
	;
	v920 = v80 - v917
	v922 = v917 & int32(3)
	v923 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v917) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v936 = v923
	v938 = int32(0)
	goto L112
L110:
	;
	v995 = v923
	goto L111
L111:
	;
	v1016 = v923
	v1017 = v995
	goto L116
L112:
	;
	v952 = v936 + v529
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v952))))
	v954 = v920 + v936
	v955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v954))))
	*(*uint8)(unsafe.Add(mBase, uint32(v952))) = uint8(v955)
	*(*uint8)(unsafe.Add(mBase, uint32(v954))) = uint8(v953)
	v959 = v936 | int32(1)
	v960 = v529 + v959
	v961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v960))))
	v962 = v920 + v959
	v963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v962))))
	*(*uint8)(unsafe.Add(mBase, uint32(v960))) = uint8(v963)
	*(*uint8)(unsafe.Add(mBase, uint32(v962))) = uint8(v961)
	v967 = v936 | int32(2)
	v968 = v529 + v967
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v968))))
	v970 = v920 + v967
	v971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v970))))
	*(*uint8)(unsafe.Add(mBase, uint32(v968))) = uint8(v971)
	*(*uint8)(unsafe.Add(mBase, uint32(v970))) = uint8(v969)
	v975 = v936 | int32(3)
	v976 = v529 + v975
	v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v976))))
	v978 = v920 + v975
	v979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v978))))
	*(*uint8)(unsafe.Add(mBase, uint32(v976))) = uint8(v979)
	*(*uint8)(unsafe.Add(mBase, uint32(v978))) = uint8(v977)
	v982 = int32(4)
	v983 = v936 + v982
	v985 = v938 + v982
	if v985 != v917&int32(-4) {
		v936 = v983
		v938 = v985
		goto L112
	} else {
		goto L114
	}
L113:
	;
	if v922 == int32(0) {
		goto L104
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	v995 = v983
	goto L111
L116:
	;
	v1033 = v1017 + v529
	v1034 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1033))))
	v1035 = v920 + v1017
	v1036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1035))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1033))) = uint8(v1036)
	*(*uint8)(unsafe.Add(mBase, uint32(v1035))) = uint8(v1034)
	v1039 = int32(1)
	v1042 = v1016 + v1039
	if v1042 != v922 {
		v1016 = v1042
		v1017 = v1017 + v1039
		goto L116
	} else {
		goto L118
	}
L117:
	;
	goto L104
L118:
	;
	goto L117
L119:
	;
	if base.Ui32(v35) < base.Ui32(v762) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v1068 = base.I32_div_u_s(v762, v35)
	F_pg_qsort(m, v33, v1068, v35, v36)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L11
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	if base.Ui32(v913) <= base.Ui32(v35) {
		goto L1
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	v1072 = v80 - v913
	v1073 = base.I32_div_u_s(v913, v35)
	if base.Ui32(int32(7)) <= base.Ui32(v1073) {
		v33 = v1072
		v34 = v1073
		goto L4
	} else {
		goto L125
	}
L125:
	;
	v1231 = v1072
	v1232 = v1073
	v1233 = v35
	v1234 = v36
	v1251 = v53
	goto L2
L126:
	;
	v305 = v548 + v53
	v307 = v550
	v308 = v528
	v309 = v35 + v529
	goto L40
L127:
	;
	v1078 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v50) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v1085 = v1078
	v1088 = v1078
	goto L131
L129:
	;
	v1148 = v1078
	goto L130
L130:
	;
	v1170 = v1148
	v1177 = v1078
	goto L135
L131:
	;
	v1106 = v1088 + v529
	v1107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1106))))
	v1108 = v1088 + v548
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1106))) = uint8(v1109)
	*(*uint8)(unsafe.Add(mBase, uint32(v1108))) = uint8(v1107)
	v1113 = v1088 | int32(1)
	v1114 = v529 + v1113
	v1115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1114))))
	v1116 = v1113 + v548
	v1117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1116))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1114))) = uint8(v1117)
	*(*uint8)(unsafe.Add(mBase, uint32(v1116))) = uint8(v1115)
	v1121 = v1088 | int32(2)
	v1122 = v529 + v1121
	v1123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1122))))
	v1124 = v1121 + v548
	v1125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1122))) = uint8(v1125)
	*(*uint8)(unsafe.Add(mBase, uint32(v1124))) = uint8(v1123)
	v1129 = v1088 | int32(3)
	v1130 = v529 + v1129
	v1131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1130))))
	v1132 = v1129 + v548
	v1133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1132))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1130))) = uint8(v1133)
	*(*uint8)(unsafe.Add(mBase, uint32(v1132))) = uint8(v1131)
	v1136 = int32(4)
	v1137 = v1088 + v1136
	v1139 = v1085 + v1136
	if v1139 != v48 {
		v1085 = v1139
		v1088 = v1137
		goto L131
	} else {
		goto L133
	}
L132:
	;
	if v49 == int32(0) {
		goto L126
	} else {
		goto L134
	}
L133:
	;
	goto L132
L134:
	;
	v1148 = v1137
	goto L130
L135:
	;
	v1187 = v1170 + v529
	v1188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1187))))
	v1189 = v1170 + v548
	v1190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1189))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1187))) = uint8(v1190)
	*(*uint8)(unsafe.Add(mBase, uint32(v1189))) = uint8(v1188)
	v1193 = int32(1)
	v1196 = v1177 + v1193
	if v1196 != v49 {
		v1170 = v1170 + v1193
		v1177 = v1196
		goto L135
	} else {
		goto L137
	}
L136:
	;
	goto L126
L137:
	;
	goto L136
L138:
	;
	v1224 = base.I32_div_u_s(v913, v35)
	F_pg_qsort(m, v80-v913, v1224, v35, v36)
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L11
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	if base.Ui32(v762) <= base.Ui32(v35) {
		goto L1
	} else {
		goto L142
	}
L141:
	;
	goto L140
L142:
	;
	v1228 = base.I32_div_u_s(v762, v35)
	if base.Ui32(int32(7)) <= base.Ui32(v1228) {
		v57 = v1228
		goto L6
	} else {
		goto L143
	}
L143:
	;
	goto L7
L144:
	;
	v1258 = int32(3)
	v1259 = v1233 & v1258
	v1284 = v1231 + v1233
	goto L145
L145:
	;
	if base.Ui32(v1284) <= base.Ui32(v1231) {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	goto L1
L147:
	;
	v1482 = v1233 + v1284
	if base.Ui32(v1482) < base.Ui32(v1231+v1253) {
		v1284 = v1482
		goto L145
	} else {
		goto L166
	}
L148:
	;
	v1293 = v1284
	goto L149
L149:
	;
	v1310 = v1293 + v1251
	v1311 = m.T0[v1234].(func(*base.Module, int32, int32) int32)(m, v1310, v1293)
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L11
	} else {
		goto L151
	}
L150:
	;
	goto L147
L151:
	;
	if v1311 <= int32(0) {
		goto L147
	} else {
		goto L152
	}
L152:
	;
	if v1233 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	if base.Ui32(v1231) < base.Ui32(v1310) {
		v1293 = v1310
		goto L149
	} else {
		goto L165
	}
L154:
	;
	v1317 = int32(0)
	if base.B2i32(base.Ui32(v1233-int32(1)) < base.Ui32(v1258)) == v1317 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v1329 = v1317
	v1335 = v1317
	goto L158
L156:
	;
	v1386 = v1317
	goto L157
L157:
	;
	v1408 = v1386
	v1414 = v1317
	goto L162
L158:
	;
	v1345 = v1293 + v1329
	v1346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1345))))
	v1347 = v1310 + v1329
	v1348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1347))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1345))) = uint8(v1348)
	*(*uint8)(unsafe.Add(mBase, uint32(v1347))) = uint8(v1346)
	v1352 = v1329 | int32(1)
	v1353 = v1293 + v1352
	v1354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1353))))
	v1355 = v1352 + v1310
	v1356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1355))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1353))) = uint8(v1356)
	*(*uint8)(unsafe.Add(mBase, uint32(v1355))) = uint8(v1354)
	v1360 = v1329 | int32(2)
	v1361 = v1293 + v1360
	v1362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1361))))
	v1363 = v1360 + v1310
	v1364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1363))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1361))) = uint8(v1364)
	*(*uint8)(unsafe.Add(mBase, uint32(v1363))) = uint8(v1362)
	v1368 = v1329 | int32(3)
	v1369 = v1293 + v1368
	v1370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1369))))
	v1371 = v1368 + v1310
	v1372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1371))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1369))) = uint8(v1372)
	*(*uint8)(unsafe.Add(mBase, uint32(v1371))) = uint8(v1370)
	v1375 = int32(4)
	v1376 = v1329 + v1375
	v1378 = v1335 + v1375
	if v1378 != v1233&int32(-4) {
		v1329 = v1376
		v1335 = v1378
		goto L158
	} else {
		goto L160
	}
L159:
	;
	if v1259 == int32(0) {
		goto L153
	} else {
		goto L161
	}
L160:
	;
	goto L159
L161:
	;
	v1386 = v1376
	goto L157
L162:
	;
	v1426 = v1293 + v1408
	v1427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1426))))
	v1428 = v1408 + v1310
	v1429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1428))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1426))) = uint8(v1429)
	*(*uint8)(unsafe.Add(mBase, uint32(v1428))) = uint8(v1427)
	v1432 = int32(1)
	v1435 = v1414 + v1432
	if v1435 != v1259 {
		v1408 = v1408 + v1432
		v1414 = v1435
		goto L162
	} else {
		goto L164
	}
L163:
	;
	goto L153
L164:
	;
	goto L163
L165:
	;
	goto L150
L166:
	;
	goto L146
}
func F_pg_read_file_common(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	if int64(0) <= l2 {
		v7 = F_convert_and_check_filename(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = F_read_binary_file(m, v7, l1, l2, l3)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				if v11 != 0 {
					v13 = int32(4)
					v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					F_pg_verifymbstr(m, v11+v13, int32(base.Ui32(v15)>>(uint(int32(2))%32))-v13)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						return v11
					}
				} else {
					return v11
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_read_file_common_0), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_read_file_common_1), int32(248), int32(_a_F_pg_read_file_common_2))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
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
func F_pg_regexec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v480 int32
	_ = v480
	var v488 int32
	_ = v488
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
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
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v706 int32
	_ = v706
	var v716 int32
	_ = v716
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v736 int32
	_ = v736
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v812 int32
	_ = v812
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v917 int32
	_ = v917
	var v923 int32
	_ = v923
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v975 int32
	_ = v975
	var v981 int32
	_ = v981
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v998 int32
	_ = v998
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1061 int32
	_ = v1061
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1130 int32
	_ = v1130
	var v1145 int32
	_ = v1145
	var v1153 int32
	_ = v1153
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1185 int32
	_ = v1185
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1238 int32
	_ = v1238
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1275 int32
	_ = v1275
	var v1279 int32
	_ = v1279
	var v1282 int32
	_ = v1282
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1304 int32
	_ = v1304
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1352 int32
	_ = v1352
	var v1361 int32
	_ = v1361
	v7 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(_a_F_pg_regexec_0)
	m.G0 = v19
	v21 = int32(16)
	if base.B2i32(l0 == v7)|base.B2i32(l1 == v7) != 0 {
		v1361 = v21
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v19 + int32(_a_F_pg_regexec_0)
	return v1361
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v27 != int32(_a_F_pg_regexec_1) {
		v1361 = v21
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v30 != int32(4) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v1361 = int32(17)
	goto L1
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(l2) < base.Ui32(l3) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v1361 = int32(1)
	goto L1
L8:
	;
	goto L9
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pg_set_regex_collation(m, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+320)) = l0
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+324)) = v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v44&int32(512) != 0 {
		v1361 = v21
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v48&int32(_a_F_pg_regexec_2) != 0 {
		v1361 = int32(1)
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+328)) = int32(0)
	v54 = v48 & int32(1)
	if v54 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v114 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+356)) = v114
	v116 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+340)) = v116
	*(*int64)(unsafe.Add(mBase, uint32(v19)+364)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v19)+372)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v19)+344)) = l1
	v123 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+352)) = l1 + l2<<(uint(v123)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+348)) = l1 + l3<<(uint(v123)%32)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v109)+68))
	if base.Ui32(int32(41)) <= base.Ui32(v131) {
		goto L42
	} else {
		goto L43
	}
L15:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v105 = v103 + int32(1)
	if base.Ui32(l4) < base.Ui32(v105) {
		goto L36
	} else {
		goto L37
	}
L16:
	;
	v94 = l4<<(uint(int32(3))%32) - int32(8)
	if v94 == int32(0) {
		goto L15
	} else {
		goto L35
	}
L17:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if base.Ui32(v55) < base.Ui32(l4) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+336)) = l5
	if base.Ui32(l4) < base.Ui32(int32(2)) {
		goto L15
	} else {
		goto L34
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+336)) = l5
	if l4 != int32(1) {
		goto L16
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v61 = v55 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+332)) = v61
	if base.Ui32(int32(21)) <= base.Ui32(v61) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L15
L24:
	;
	v81 = v61<<(uint(int32(3))%32) - int32(8)
	if v81 != 0 {
		goto L31
	} else {
		goto L32
	}
L25:
	;
	v68 = F_palloc_extended(m, v61<<(uint(int32(3))%32), int32(2))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v73 = v19 + int32(160)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+336)) = v73
	if v55 == int32(0) {
		v109 = v42
		v110 = l4
		goto L14
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+336)) = v68
	if v68 != 0 {
		v77 = v68
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v1361 = int32(12)
	goto L1
L30:
	;
	v77 = v73
	goto L24
L31:
	;
	base.MemoryFill(m, v77+int32(8), int32(255), v81)
	goto L33
L32:
	;
	goto L33
L33:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v19)+324))
	v109 = v86
	v110 = l4
	goto L14
L34:
	;
	goto L16
L35:
	;
	base.MemoryFill(m, l5+int32(8), int32(255), v94)
	goto L15
L36:
	;
	v107 = l4
	goto L38
L37:
	;
	v107 = v105
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+332)) = v107
	v109 = v42
	v110 = v107
	goto L14
L39:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v19)+336))
	if base.B2i32(v1153 == l5)|base.B2i32(v1153 == v19+int32(160)) == int32(0) {
		goto L340
	} else {
		goto L341
	}
L40:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v19)+324))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+428))
	if v257 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L41:
	;
	v145 = int32(3)
	v146 = v131 & v145
	v147 = int32(0)
	if base.Ui32(v145) <= base.Ui32(v131-int32(1)) {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	v134 = int32(2)
	v137 = F_palloc_extended(m, v131<<(uint(v134)%32), v134)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L10
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+360)) = v19
	if v131 == int32(0) {
		goto L40
	} else {
		goto L47
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+360)) = v137
	if v137 != 0 {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v1145 = int32(12)
	goto L39
L47:
	;
	goto L41
L48:
	;
	v155 = v147
	v156 = int32(0)
	goto L51
L49:
	;
	v196 = v147
	goto L50
L50:
	;
	v213 = v196
	v221 = int32(0)
	goto L55
L51:
	;
	v172 = v155 << (uint(int32(2)) % 32)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v19)+360))
	v175 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v172+v173))) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v19)+360))
	*(*int32)(unsafe.Add(mBase, uint32(v177+v172)+4)) = v175
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v19)+360))
	*(*int32)(unsafe.Add(mBase, uint32(v181+v172)+8)) = v175
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v19)+360))
	*(*int32)(unsafe.Add(mBase, uint32(v185+v172)+12)) = v175
	v189 = int32(4)
	v190 = v155 + v189
	v192 = v156 + v189
	if v192 != v131&int32(-4) {
		v155 = v190
		v156 = v192
		goto L51
	} else {
		goto L53
	}
L52:
	;
	if v146 == int32(0) {
		goto L40
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v196 = v190
	goto L50
L55:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v19)+360))
	*(*int32)(unsafe.Add(mBase, uint32(v229+v213<<(uint(int32(2))%32)))) = int32(0)
	v235 = int32(1)
	v238 = v221 + v235
	if v238 != v146 {
		v213 = v213 + v235
		v221 = v238
		goto L55
	} else {
		goto L57
	}
L56:
	;
	goto L40
L57:
	;
	goto L56
L58:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v19)+324))
	v529 = v527 + int32(72)
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v527)+16))
	v532 = v530 + int32(36)
	if v54 != 0 {
		goto L87
	} else {
		goto L88
	}
L59:
	;
	v260 = int32(2)
	v261 = v257 << (uint(v260) % 32)
	v263 = F_palloc_extended(m, v261, v260)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+364)) = v263
	v266 = int32(12)
	if v263 == int32(0) {
		v1145 = v266
		goto L39
	} else {
		goto L61
	}
L61:
	;
	v270 = v257 & int32(3)
	v271 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v257) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v379 = F_palloc_extended(m, v261, int32(2))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L10
	} else {
		goto L73
	}
L63:
	;
	v277 = v271
	v280 = int32(0)
	goto L66
L64:
	;
	v318 = v271
	goto L65
L65:
	;
	v335 = v318
	v336 = int32(0)
	goto L70
L66:
	;
	v294 = v277 << (uint(int32(2)) % 32)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v19)+364))
	v297 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v294+v295))) = v297
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v19)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v299+v294)+4)) = v297
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v19)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v303+v294)+8)) = v297
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v19)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v307+v294)+12)) = v297
	v311 = int32(4)
	v312 = v277 + v311
	v314 = v280 + v311
	if v314 != v257&int32(-4) {
		v277 = v312
		v280 = v314
		goto L66
	} else {
		goto L68
	}
L67:
	;
	if v270 == int32(0) {
		goto L62
	} else {
		goto L69
	}
L68:
	;
	goto L67
L69:
	;
	v318 = v312
	goto L65
L70:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v19)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v351+v335<<(uint(int32(2))%32)))) = int32(0)
	v357 = int32(1)
	v360 = v336 + v357
	if v360 != v270 {
		v335 = v335 + v357
		v336 = v360
		goto L70
	} else {
		goto L72
	}
L71:
	;
	goto L62
L72:
	;
	goto L71
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+368)) = v379
	v383 = F_palloc_extended(m, v261, int32(2))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+372)) = v383
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v19)+368))
	v387 = int32(0)
	if base.B2i32(v386 == v387)|base.B2i32(v383 == v387) != 0 {
		v1145 = v266
		goto L39
	} else {
		goto L75
	}
L75:
	;
	v393 = v257 & int32(3)
	v394 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v257) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v400 = v394
	v401 = int32(0)
	goto L79
L77:
	;
	v463 = v394
	goto L78
L78:
	;
	v480 = v463
	v488 = int32(0)
	goto L83
L79:
	;
	v417 = v400 << (uint(int32(2)) % 32)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v19)+368))
	v420 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v417+v418))) = v420
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v19)+372))
	*(*int32)(unsafe.Add(mBase, uint32(v422+v417))) = v420
	v426 = int32(4)
	v427 = v417 | v426
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v19)+368))
	*(*int32)(unsafe.Add(mBase, uint32(v427+v428))) = v420
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v19)+372))
	*(*int32)(unsafe.Add(mBase, uint32(v432+v427))) = v420
	v437 = v417 | int32(8)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v19)+368))
	*(*int32)(unsafe.Add(mBase, uint32(v437+v438))) = v420
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v19)+372))
	*(*int32)(unsafe.Add(mBase, uint32(v442+v437))) = v420
	v447 = v417 | int32(12)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v19)+368))
	*(*int32)(unsafe.Add(mBase, uint32(v447+v448))) = v420
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v19)+372))
	*(*int32)(unsafe.Add(mBase, uint32(v452+v447))) = v420
	v457 = v400 + v426
	v459 = v401 + v426
	if v459 != v257&int32(-4) {
		v400 = v457
		v401 = v459
		goto L79
	} else {
		goto L81
	}
L80:
	;
	if v393 == int32(0) {
		goto L58
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	v463 = v457
	goto L78
L83:
	;
	v497 = v480 << (uint(int32(2)) % 32)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v19)+368))
	v500 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v497+v498))) = v500
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v19)+372))
	*(*int32)(unsafe.Add(mBase, uint32(v502+v497))) = v500
	v506 = int32(1)
	v509 = v488 + v506
	if v509 != v393 {
		v480 = v480 + v506
		v488 = v509
		goto L83
	} else {
		goto L85
	}
L84:
	;
	goto L58
L85:
	;
	goto L84
L86:
	;
	if v1105|base.B2i32(v110 == int32(0)) != 0 {
		v1145 = v1105
		goto L39
	} else {
		goto L334
	}
L87:
	;
	v533 = m.G0
	v535 = v533 - int32(16)
	m.G0 = v535
	v538 = v19 + int32(320)
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v538)+4))
	v544 = F_newdfa(m, v538, v539+int32(20), v529, v19+int32(376))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L10
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v829 = m.G0
	v831 = v829 - int32(16)
	m.G0 = v831
	v834 = v19 + int32(320)
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v834)+4))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v835)+16))
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v836)+1)))
	v841 = v19 + int32(376)
	v842 = F_newdfa(m, v834, v835+int32(20), v529, v841)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L10
	} else {
		goto L219
	}
L90:
	;
	m.G0 = v535 + int32(16)
	v1105 = v812
	goto L86
L91:
	;
	if v544 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v538)+36))
	v812 = v548
	goto L90
L93:
	;
	goto L94
L94:
	;
	v551 = F_newdfa(m, v538, v532, v529, v19+int32(_a_F_pg_regexec_3))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L10
	} else {
		goto L95
	}
L95:
	;
	if v551 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544)+69)))
	if v555 != int32(1) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v538)+4))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v579)+16))
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v580)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v535)+12)) = int32(0)
	v585 = v581 & int32(2)
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v538)+32))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v538)+28))
	v588 = v586
	v590 = v587
	goto L119
L99:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544)+68)))
	if v573 == int32(1) {
		goto L115
	} else {
		goto L116
	}
L100:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v544)+20))
	if v558 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	F_pfree(m, v558)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L10
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v544)+24))
	if v561 != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	goto L103
L105:
	;
	F_pfree(m, v561)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L10
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v544)+32))
	if v564 != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	goto L107
L109:
	;
	F_pfree(m, v564)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L10
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v544)+36))
	if v567 == int32(0) {
		goto L99
	} else {
		goto L113
	}
L112:
	;
	goto L111
L113:
	;
	F_pfree(m, v567)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L10
	} else {
		goto L114
	}
L114:
	;
	goto L99
L115:
	;
	F_pfree(m, v544)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L10
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v538)+36))
	v812 = v578
	goto L90
L118:
	;
	goto L117
L119:
	;
	v607 = F_shortest(m, v538, v544, v590, v590, v588, v535+int32(12), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L10
	} else {
		goto L122
	}
L120:
	;
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+69)))
	if v744 != int32(1) {
		goto L171
	} else {
		goto L172
	}
L121:
	;
	goto L120
L122:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v538)+36))
	if v609 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v535)+12))
	v729 = v609
	v736 = v610
	goto L121
L124:
	;
	goto L125
L125:
	;
	v611 = int32(1)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v535)+12))
	if v607 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v729 = v611
	v736 = v612
	goto L121
L127:
	;
	goto L128
L128:
	;
	v615 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v535)+12)) = v615
	if base.Ui32(v612) <= base.Ui32(v607) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v619 = v612
	goto L132
L130:
	;
	v716 = v615
	goto L131
L131:
	;
	v725 = v607 + int32(4)
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v538)+32))
	if base.Ui32(v725) < base.Ui32(v726) {
		v588 = v726
		v590 = v725
		goto L119
	} else {
		goto L170
	}
L132:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v538)+32))
	v639 = v619
	v650 = v635
	goto L134
L133:
	;
	v716 = v667
	goto L131
L134:
	;
	if v585 != 0 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v706 = v619 + int32(4)
	if base.Ui32(v706) <= base.Ui32(v607) {
		v619 = v706
		goto L132
	} else {
		goto L169
	}
L136:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v538)+36))
	if v662 != 0 {
		goto L142
	} else {
		goto L143
	}
L137:
	;
	v655 = F_shortest(m, v538, v551, v619, v639, v650, int32(0), v535+int32(8))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L10
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v659 = F_longest(m, v538, v551, v619, v650, v535+int32(8))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L10
	} else {
		goto L141
	}
L140:
	;
	v661 = v655
	goto L136
L141:
	;
	v661 = v659
	goto L136
L142:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v535)+12))
	v729 = v662
	v736 = v663
	goto L121
L143:
	;
	goto L144
L144:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v535)+12))
	if v664 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v665 = v664
	goto L147
L146:
	;
	v665 = v619
	goto L147
L147:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v535)+8))
	if v666 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v667 = v665
	goto L150
L149:
	;
	v667 = v664
	goto L150
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v535)+12)) = v667
	if v661 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	goto L135
L152:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v538)+4))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v671)+16))
	v673 = F_cdissect(m, v538, v672, v619, v661)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L10
	} else {
		goto L153
	}
L153:
	;
	if v673 != int32(1) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	if v673 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	goto L156
L156:
	;
	if v585 != 0 {
		goto L164
	} else {
		goto L165
	}
L157:
	;
	v679 = int32(0)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v538)+12))
	if v680 == v679 {
		v729 = v679
		v736 = v667
		goto L121
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v538)+36))
	if v695 != 0 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v538)+16))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v538)+24))
	v686 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v683))) = (v619 - v684) >> (uint(v686) % 32)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v538)+16))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v538)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v689)+4)) = (v661 - v690) >> (uint(v686) % 32)
	v729 = v679
	v736 = v667
	goto L121
L161:
	;
	v696 = v695
	goto L163
L162:
	;
	v696 = v673
	goto L163
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538)+36)) = v696
	v729 = v673
	v736 = v667
	goto L121
L164:
	;
	if v661 == v650 {
		goto L151
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	if v619 == v661 {
		goto L151
	} else {
		goto L168
	}
L167:
	;
	v639 = v661 + int32(4)
	goto L134
L168:
	;
	v650 = v661 - int32(4)
	goto L134
L169:
	;
	goto L133
L170:
	;
	v729 = v611
	v736 = v716
	goto L121
L171:
	;
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+68)))
	if v762 == int32(1) {
		goto L187
	} else {
		goto L188
	}
L172:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v551)+20))
	if v747 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	F_pfree(m, v747)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L10
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v551)+24))
	if v750 != 0 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	goto L175
L177:
	;
	F_pfree(m, v750)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L10
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v551)+32))
	if v753 != 0 {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	goto L179
L181:
	;
	F_pfree(m, v753)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L10
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v551)+36))
	if v756 == int32(0) {
		goto L171
	} else {
		goto L185
	}
L184:
	;
	goto L183
L185:
	;
	F_pfree(m, v756)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L10
	} else {
		goto L186
	}
L186:
	;
	goto L171
L187:
	;
	F_pfree(m, v551)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L10
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544)+69)))
	if v767 != int32(1) {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	goto L189
L191:
	;
	v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544)+68)))
	if v785 == int32(1) {
		goto L207
	} else {
		goto L208
	}
L192:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v544)+20))
	if v770 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	F_pfree(m, v770)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L10
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v544)+24))
	if v773 != 0 {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	goto L195
L197:
	;
	F_pfree(m, v773)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L10
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v544)+32))
	if v776 != 0 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	goto L199
L201:
	;
	F_pfree(m, v776)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L10
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v544)+36))
	if v779 == int32(0) {
		goto L191
	} else {
		goto L205
	}
L204:
	;
	goto L203
L205:
	;
	F_pfree(m, v779)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L10
	} else {
		goto L206
	}
L206:
	;
	goto L191
L207:
	;
	F_pfree(m, v544)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L10
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v538)+36))
	if v790 != 0 {
		v812 = v790
		goto L90
	} else {
		goto L211
	}
L210:
	;
	goto L209
L211:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v538)+4))
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v791)+5)))
	if v792&int32(2) != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v538)+20))
	if v736 != 0 {
		goto L215
	} else {
		goto L216
	}
L213:
	;
	goto L214
L214:
	;
	v812 = v729
	goto L90
L215:
	;
	v797 = v736
	goto L217
L216:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v538)+32))
	v797 = v796
	goto L217
L217:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v538)+24))
	v800 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v795))) = (v797 - v798) >> (uint(v800) % 32)
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v538)+20))
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v538)+32))
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v538)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v803)+4)) = (v804 - v805) >> (uint(v800) % 32)
	goto L214
L218:
	;
	m.G0 = v831 + int32(16)
	v1105 = v1072
	goto L86
L219:
	;
	if v842 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v834)+36))
	v1072 = v846
	goto L218
L221:
	;
	goto L222
L222:
	;
	v847 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v831)+12)) = v847
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v834)+28))
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v834)+32))
	v854 = F_shortest(m, v834, v842, v849, v849, v850, v831+int32(12), v847)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L10
	} else {
		goto L223
	}
L223:
	;
	v856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842)+69)))
	if v856 != int32(1) {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842)+68)))
	if v874 == int32(1) {
		goto L240
	} else {
		goto L241
	}
L225:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v842)+20))
	if v859 != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	F_pfree(m, v859)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L10
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v842)+24))
	if v862 != 0 {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	goto L228
L230:
	;
	F_pfree(m, v862)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L10
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v842)+32))
	if v865 != 0 {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	goto L232
L234:
	;
	F_pfree(m, v865)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L10
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v842)+36))
	if v868 == int32(0) {
		goto L224
	} else {
		goto L238
	}
L237:
	;
	goto L236
L238:
	;
	F_pfree(m, v868)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L10
	} else {
		goto L239
	}
L239:
	;
	goto L224
L240:
	;
	F_pfree(m, v842)
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L10
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v834)+36))
	if v879 != 0 {
		v1072 = v879
		goto L218
	} else {
		goto L244
	}
L243:
	;
	goto L242
L244:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v834)+4))
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v880)+5)))
	if v881&int32(2) != 0 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v834)+20))
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v831)+12))
	if v885 != 0 {
		goto L248
	} else {
		goto L249
	}
L246:
	;
	goto L247
L247:
	;
	if v854 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L248:
	;
	v887 = v885
	goto L250
L249:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v834)+32))
	v887 = v886
	goto L250
L250:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v834)+24))
	v890 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v884))) = (v887 - v888) >> (uint(v890) % 32)
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v834)+20))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v834)+32))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v834)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v893)+4)) = (v894 - v895) >> (uint(v890) % 32)
	goto L247
L251:
	;
	v1072 = int32(1)
	goto L218
L252:
	;
	goto L253
L253:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v834)+12))
	if v904 == int32(0) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1072 = int32(0)
	goto L218
L255:
	;
	goto L256
L256:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v831)+12))
	v909 = F_newdfa(m, v834, v532, v529, v841)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L10
	} else {
		goto L257
	}
L257:
	;
	if v909 != 0 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v911 = int32(0)
	if base.Ui32(v908) <= base.Ui32(v854) {
		goto L262
	} else {
		goto L263
	}
L259:
	;
	goto L260
L260:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v834)+36))
	v1072 = v1069
	goto L218
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v831)+12)) = v998
	v1007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v909)+69)))
	if v1007 != int32(1) {
		goto L304
	} else {
		goto L305
	}
L262:
	;
	v917 = v908
	v923 = v911
	goto L265
L263:
	;
	v975 = v908
	v981 = v911
	goto L264
L264:
	;
	v990 = int32(0)
	v992 = v975
	v998 = v981
	goto L261
L265:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v834)+32))
	if v837&int32(2) != 0 {
		goto L268
	} else {
		goto L269
	}
L266:
	;
	v975 = v971
	v981 = v969
	goto L264
L267:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v834)+36))
	if v942 != 0 {
		goto L273
	} else {
		goto L274
	}
L268:
	;
	v935 = F_shortest(m, v834, v909, v917, v917, v931, int32(0), v831+int32(8))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L10
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v939 = F_longest(m, v834, v909, v917, v931, v831+int32(8))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L10
	} else {
		goto L272
	}
L271:
	;
	v941 = v935
	goto L267
L272:
	;
	v941 = v939
	goto L267
L273:
	;
	v943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v909)+69)))
	if v943 != int32(1) {
		goto L276
	} else {
		goto L277
	}
L274:
	;
	goto L275
L275:
	;
	if v923 != 0 {
		goto L296
	} else {
		goto L297
	}
L276:
	;
	v961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v909)+68)))
	if v961 == int32(1) {
		goto L292
	} else {
		goto L293
	}
L277:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v909)+20))
	if v946 != 0 {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	F_pfree(m, v946)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L10
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v909)+24))
	if v949 != 0 {
		goto L282
	} else {
		goto L283
	}
L281:
	;
	goto L280
L282:
	;
	F_pfree(m, v949)
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L10
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v909)+32))
	if v952 != 0 {
		goto L286
	} else {
		goto L287
	}
L285:
	;
	goto L284
L286:
	;
	F_pfree(m, v952)
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L10
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v909)+36))
	if v955 == int32(0) {
		goto L276
	} else {
		goto L290
	}
L289:
	;
	goto L288
L290:
	;
	F_pfree(m, v955)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L10
	} else {
		goto L291
	}
L291:
	;
	goto L276
L292:
	;
	F_pfree(m, v909)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L10
	} else {
		goto L295
	}
L293:
	;
	goto L294
L294:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v834)+36))
	v1072 = v966
	goto L218
L295:
	;
	goto L294
L296:
	;
	v967 = v923
	goto L298
L297:
	;
	v967 = v917
	goto L298
L298:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v831)+8))
	if v968 != 0 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v969 = v967
	goto L301
L300:
	;
	v969 = v923
	goto L301
L301:
	;
	if v941 != 0 {
		v990 = v941
		v992 = v917
		v998 = v969
		goto L261
	} else {
		goto L302
	}
L302:
	;
	v971 = v917 + int32(4)
	if base.Ui32(v971) <= base.Ui32(v854) {
		v917 = v971
		v923 = v969
		goto L265
	} else {
		goto L303
	}
L303:
	;
	goto L266
L304:
	;
	v1025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v909)+68)))
	if v1025 == int32(1) {
		goto L320
	} else {
		goto L321
	}
L305:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v909)+20))
	if v1010 != 0 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	F_pfree(m, v1010)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L10
	} else {
		goto L309
	}
L307:
	;
	goto L308
L308:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v909)+24))
	if v1013 != 0 {
		goto L310
	} else {
		goto L311
	}
L309:
	;
	goto L308
L310:
	;
	F_pfree(m, v1013)
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L10
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v909)+32))
	if v1016 != 0 {
		goto L314
	} else {
		goto L315
	}
L313:
	;
	goto L312
L314:
	;
	F_pfree(m, v1016)
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L10
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v909)+36))
	if v1019 == int32(0) {
		goto L304
	} else {
		goto L318
	}
L317:
	;
	goto L316
L318:
	;
	F_pfree(m, v1019)
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L10
	} else {
		goto L319
	}
L319:
	;
	goto L304
L320:
	;
	F_pfree(m, v909)
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L10
	} else {
		goto L323
	}
L321:
	;
	goto L322
L322:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v834)+16))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v834)+24))
	v1033 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1030))) = (v992 - v1031) >> (uint(v1033) % 32)
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v834)+16))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v834)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1036)+4)) = (v990 - v1037) >> (uint(v1033) % 32)
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v834)+4))
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1042)+5)))
	if v1043&v1033 != 0 {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	goto L322
L324:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v834)+20))
	if v998 != 0 {
		goto L327
	} else {
		goto L328
	}
L325:
	;
	goto L326
L326:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v834)+12))
	if v1061 == int32(1) {
		goto L330
	} else {
		goto L331
	}
L327:
	;
	v1048 = v998
	goto L329
L328:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v834)+32))
	v1048 = v1047
	goto L329
L329:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v834)+24))
	v1051 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1046))) = (v1048 - v1049) >> (uint(v1051) % 32)
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v834)+20))
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v834)+32))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v834)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1054)+4)) = (v1055 - v1056) >> (uint(v1051) % 32)
	goto L326
L330:
	;
	v1072 = int32(0)
	goto L218
L331:
	;
	goto L332
L332:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v834)+4))
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v1065)+16))
	v1067 = F_cdissect(m, v834, v1066, v992, v990)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L10
	} else {
		goto L333
	}
L333:
	;
	v1072 = v1067
	goto L218
L334:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v19)+336))
	if v1109 == l5 {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v1117 = int32(0)
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v19)+324))
	v1119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1118)+4)))
	if base.B2i32(v1119&int32(16) == v1117)|base.B2i32(v110 == int32(1)) != 0 {
		v1145 = v1117
		goto L39
	} else {
		goto L338
	}
L336:
	;
	v1112 = v110 << (uint(int32(3)) % 32)
	if v1112 == int32(0) {
		goto L335
	} else {
		goto L337
	}
L337:
	;
	base.MemoryCopy(m, l5, v1109, v1112)
	goto L335
L338:
	;
	v1130 = v110<<(uint(int32(3))%32) - int32(8)
	if v1130 == int32(0) {
		v1145 = v1117
		goto L39
	} else {
		goto L339
	}
L339:
	;
	base.MemoryFill(m, l5+int32(8), int32(255), v1130)
	v1145 = v1117
	goto L39
L340:
	;
	F_pfree(m, v1153)
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L10
	} else {
		goto L343
	}
L341:
	;
	goto L342
L342:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v19)+360))
	if v1163 == int32(0) {
		goto L344
	} else {
		goto L345
	}
L343:
	;
	goto L342
L344:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v19)+364))
	if v1255 != 0 {
		goto L374
	} else {
		goto L375
	}
L345:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v19)+324))
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+68))
	if v1167 != 0 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v1170 = int32(0)
	goto L349
L347:
	;
	v1220 = v1163
	goto L348
L348:
	;
	if v1220 == v19 {
		goto L344
	} else {
		goto L372
	}
L349:
	;
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v19)+360))
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1185+v1170<<(uint(int32(2))%32))))
	if v1189 == int32(0) {
		goto L351
	} else {
		goto L352
	}
L350:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v19)+360))
	v1220 = v1219
	goto L348
L351:
	;
	v1217 = v1170 + int32(1)
	if v1217 != v1167 {
		v1170 = v1217
		goto L349
	} else {
		goto L371
	}
L352:
	;
	v1192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1189)+69)))
	if v1192 != int32(1) {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v1210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1189)+68)))
	if v1210 != int32(1) {
		goto L351
	} else {
		goto L369
	}
L354:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v1189)+20))
	if v1195 != 0 {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	F_pfree(m, v1195)
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L10
	} else {
		goto L358
	}
L356:
	;
	goto L357
L357:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1189)+24))
	if v1198 != 0 {
		goto L359
	} else {
		goto L360
	}
L358:
	;
	goto L357
L359:
	;
	F_pfree(m, v1198)
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L10
	} else {
		goto L362
	}
L360:
	;
	goto L361
L361:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1189)+32))
	if v1201 != 0 {
		goto L363
	} else {
		goto L364
	}
L362:
	;
	goto L361
L363:
	;
	F_pfree(m, v1201)
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L10
	} else {
		goto L366
	}
L364:
	;
	goto L365
L365:
	;
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1189)+36))
	if v1204 == int32(0) {
		goto L353
	} else {
		goto L367
	}
L366:
	;
	goto L365
L367:
	;
	F_pfree(m, v1204)
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L10
	} else {
		goto L368
	}
L368:
	;
	goto L353
L369:
	;
	F_pfree(m, v1189)
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L10
	} else {
		goto L370
	}
L370:
	;
	goto L351
L371:
	;
	goto L350
L372:
	;
	F_pfree(m, v1220)
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L10
	} else {
		goto L373
	}
L373:
	;
	goto L344
L374:
	;
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v19)+324))
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1256)+428))
	if v1257 != 0 {
		goto L377
	} else {
		goto L378
	}
L375:
	;
	goto L376
L376:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v19)+368))
	if v1345 != 0 {
		goto L404
	} else {
		goto L405
	}
L377:
	;
	v1260 = int32(0)
	goto L380
L378:
	;
	v1326 = v1255
	goto L379
L379:
	;
	F_pfree(m, v1326)
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L10
	} else {
		goto L403
	}
L380:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v19)+364))
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v1275+v1260<<(uint(int32(2))%32))))
	if v1279 == int32(0) {
		goto L382
	} else {
		goto L383
	}
L381:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v19)+364))
	v1326 = v1309
	goto L379
L382:
	;
	v1307 = v1260 + int32(1)
	if v1307 != v1257 {
		v1260 = v1307
		goto L380
	} else {
		goto L402
	}
L383:
	;
	v1282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1279)+69)))
	if v1282 != int32(1) {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v1300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1279)+68)))
	if v1300 != int32(1) {
		goto L382
	} else {
		goto L400
	}
L385:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1279)+20))
	if v1285 != 0 {
		goto L386
	} else {
		goto L387
	}
L386:
	;
	F_pfree(m, v1285)
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L10
	} else {
		goto L389
	}
L387:
	;
	goto L388
L388:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1279)+24))
	if v1288 != 0 {
		goto L390
	} else {
		goto L391
	}
L389:
	;
	goto L388
L390:
	;
	F_pfree(m, v1288)
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L10
	} else {
		goto L393
	}
L391:
	;
	goto L392
L392:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1279)+32))
	if v1291 != 0 {
		goto L394
	} else {
		goto L395
	}
L393:
	;
	goto L392
L394:
	;
	F_pfree(m, v1291)
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L10
	} else {
		goto L397
	}
L395:
	;
	goto L396
L396:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1279)+36))
	if v1294 == int32(0) {
		goto L384
	} else {
		goto L398
	}
L397:
	;
	goto L396
L398:
	;
	F_pfree(m, v1294)
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L10
	} else {
		goto L399
	}
L399:
	;
	goto L384
L400:
	;
	F_pfree(m, v1279)
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L10
	} else {
		goto L401
	}
L401:
	;
	goto L382
L402:
	;
	goto L381
L403:
	;
	goto L376
L404:
	;
	F_pfree(m, v1345)
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L10
	} else {
		goto L407
	}
L405:
	;
	goto L406
L406:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v19)+372))
	if v1348 == int32(0) {
		v1361 = v1145
		goto L1
	} else {
		goto L408
	}
L407:
	;
	goto L406
L408:
	;
	F_pfree(m, v1348)
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L10
	} else {
		goto L409
	}
L409:
	;
	v1361 = v1145
	goto L1
}
func F_pg_safe_snapshot_blocking_pids(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_pg_safe_snapshot_blocking_pids[0]))
	v16 = F_palloc(m, v13<<(uint(int32(2))%32))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_pg_safe_snapshot_blocking_pids[0]))
		v22 = F_GetSafeSnapshotBlockingPids(m, v11, v16, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if v22 <= int32(0) {
				v111 = v2
				v120 = F_construct_array_builtin(m, v111, v22, int32(23))
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
					return int32(0)
				} else {
					return v120
				}
			} else {
				v27 = v22 & int32(3)
				v30 = F_palloc(m, v22<<(uint(int32(2))%32))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v32 = int32(0)
					if base.Ui32(int32(4)) <= base.Ui32(v22) {
						v37 = v32
						v45 = v2
						for {
							v48 = v37 << (uint(int32(2)) % 32)
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v48+v16)))
							*(*int32)(unsafe.Add(mBase, uint32(v30+v48))) = v51
							v53 = int32(4)
							v54 = v48 | v53
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v16+v54)))
							*(*int32)(unsafe.Add(mBase, uint32(v30+v54))) = v57
							v60 = v48 | int32(8)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v16+v60)))
							*(*int32)(unsafe.Add(mBase, uint32(v30+v60))) = v63
							v66 = v48 | int32(12)
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v66+v16)))
							*(*int32)(unsafe.Add(mBase, uint32(v30+v66))) = v69
							v72 = v37 + v53
							v74 = v45 + v53
							if v74 != v22&int32(2147483644) {
								v37 = v72
								v45 = v74
								continue
							} else {
								break
							}
							break
						}
						if v27 == int32(0) {
							v111 = v30
						} else {
							v78 = v72
							v88 = v78
							v97 = v2
							for {
								v99 = v88 << (uint(int32(2)) % 32)
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v99+v16)))
								*(*int32)(unsafe.Add(mBase, uint32(v30+v99))) = v102
								v104 = int32(1)
								v107 = v97 + v104
								if v107 != v27 {
									v88 = v88 + v104
									v97 = v107
									continue
								} else {
									break
								}
								break
							}
							v111 = v30
						}
					} else {
						v78 = v32
						v88 = v78
						v97 = v2
						for {
							v99 = v88 << (uint(int32(2)) % 32)
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v99+v16)))
							*(*int32)(unsafe.Add(mBase, uint32(v30+v99))) = v102
							v104 = int32(1)
							v107 = v97 + v104
							if v107 != v27 {
								v88 = v88 + v104
								v97 = v107
								continue
							} else {
								break
							}
							break
						}
						v111 = v30
					}
					v120 = F_construct_array_builtin(m, v111, v22, int32(23))
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int32(0)
					} else {
						return v120
					}
				}
			}
		}
	}
}
func F_pg_set_noblock(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v2 = m.G0
	v3 = int32(16)
	v4 = v2 - v3
	m.G0 = v4
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(2048)
	m.G0 = v4 + v3
	return int32(1)
}
func F_pg_size_bytes(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v149 int32
	_ = v149
	var v154 int64
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
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
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v359 int32
	_ = v359
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v404 int32
	_ = v404
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v449 int32
	_ = v449
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v472 int32
	_ = v472
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v494 int32
	_ = v494
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v517 int32
	_ = v517
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v540 int32
	_ = v540
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int64
	_ = v573
	var v574 int64
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
	var v587 int32
	_ = v587
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int64
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = F_text_to_cstring(m, v15)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = v19
	goto L5
L4:
	;
	v46 = v43 - int32(48)
	if base.Ui32(v46&int32(255)) <= base.Ui32(int32(9)) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if base.Ui32(v30-int32(9)) < base.Ui32(int32(5)) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	v43 = v42
	v44 = v23 + int32(1)
	goto L4
L7:
	;
	goto L6
L8:
	;
	v23 = v23 + int32(1)
	goto L5
L9:
	;
	switch v30 - int32(32) {
	case 0:
		goto L8
	default:
		v43 = v30
		v44 = v23
		goto L4
	case 11, 13:
		goto L7
	}
L10:
	;
	v54 = v44
	goto L13
L11:
	;
	v69 = v43
	v72 = v44
	goto L12
L12:
	;
	if v69&int32(255) != int32(46) {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v61 = v54 + int32(1)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	if base.Ui32((v62-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v54 = v61
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v69 = v62
	v72 = v61
	goto L12
L15:
	;
	goto L14
L16:
	;
	if (v133|int32(32))&int32(255) == int32(101) {
		goto L30
	} else {
		goto L31
	}
L17:
	;
	if base.Ui32(v46&int32(255)) < base.Ui32(int32(10)) {
		v133 = v109
		v134 = v110
		goto L16
	} else {
		goto L25
	}
L18:
	;
	v109 = v69
	v110 = v72
	goto L17
L19:
	;
	goto L20
L20:
	;
	v83 = v72 + int32(1)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	if base.Ui32(int32(9)) < base.Ui32((v84-int32(48))&int32(255)) {
		v109 = v84
		v110 = v83
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v92 = v83
	goto L22
L22:
	;
	v101 = v92 + int32(1)
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+1)))
	if base.Ui32((v102-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v92 = v101
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v133 = v102
	v134 = v101
	goto L16
L24:
	;
	goto L23
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v19
	F_errmsg(m, int32(_a_F_pg_size_bytes_0), v12+int32(32))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_pg_size_bytes_1), int32(782), int32(_a_F_pg_size_bytes_2))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	v149 = v134 + int32(1)
	v154 = F_strtox_2(m, v149, v12+int32(44), int32(10), int64(2147483648))
	mBase = m.M
	goto L33
L31:
	;
	v160 = v133
	v161 = v134
	goto L32
L32:
	;
	v163 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v161))) = uint8(v163)
	v169 = F_DirectFunctionCall3Coll(m, int32(408), v163, v23, v163, int32(-1))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L37
	}
L33:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	if base.Ui32(v149) < base.Ui32(v156) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v158 = v156
	goto L36
L35:
	;
	v158 = v134
	goto L36
L36:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	v160 = v159
	v161 = v158
	goto L32
L37:
	;
	v171 = F_pg_detoast_datum(m, v169)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v161))) = uint8(v160)
	v174 = v160
	v175 = v161
	goto L39
L39:
	;
	v184 = v174 & int32(255)
	if base.B2i32(base.Ui32(v184-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v184 == int32(32)) != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L172
	}
L41:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
	v174 = v192
	v175 = v175 + int32(1)
	goto L39
L42:
	;
	if v184 == int32(0) {
		v587 = v171
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L40
L44:
	;
	goto L43
L45:
	;
	v596 = F_DirectFunctionCall1Coll(m, int32(1261), int32(0), v587)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L170
	}
L46:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v197 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v228 = v226 + v19
	goto L58
L48:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v203 == int32(18) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v214 = int32(1)
	if v197&v214 != 0 {
		v226 = int32(base.Ui32(v197)>>(uint(v214)%32)) - v214
		goto L47
	} else {
		goto L57
	}
L51:
	;
	v206 = int32(16)
	goto L53
L52:
	;
	v206 = int32(0)
	goto L53
L53:
	;
	if base.Ui32((v203-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v213 = int32(4)
	goto L56
L55:
	;
	v213 = v206
	goto L56
L56:
	;
	v226 = v213
	goto L47
L57:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v226 = int32(base.Ui32(v220)>>(uint(int32(2))%32)) - int32(4)
	goto L47
L58:
	;
	v238 = v228 - int32(1)
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	if base.B2i32(base.Ui32(v239-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v239 == int32(32)) != 0 {
		v228 = v238
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v247 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v228))) = uint8(v247)
	v254 = v175
	v255 = int32(_a_F_pg_size_bytes_3)
	goto L63
L60:
	;
	goto L59
L61:
	;
	v573 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v572)+9)))
	v574 = int64(1) << (uint(v573) % 64)
	if v574 < int64(2) {
		v587 = v171
		goto L45
	} else {
		goto L166
	}
L62:
	;
	if v292 == int32(0) {
		v572 = int32(_a_F_pg_size_bytes_4)
		goto L61
	} else {
		goto L75
	}
L63:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	if v258 == v259 {
		v281 = v258
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v292 = int32(0)
	goto L62
L65:
	;
	v283 = int32(1)
	if v281 != 0 {
		v254 = v254 + v283
		v255 = v255 + v283
		goto L63
	} else {
		goto L74
	}
L66:
	;
	if base.Ui32((v258-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v269 = v258 | int32(32)
	goto L69
L68:
	;
	v269 = v258
	goto L69
L69:
	;
	if base.Ui32((v259-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v278 = v259 | int32(32)
	goto L72
L71:
	;
	v278 = v259
	goto L72
L72:
	;
	if v269 == v278 {
		v281 = v269
		goto L65
	} else {
		goto L73
	}
L73:
	;
	v292 = v269 - v278
	goto L62
L74:
	;
	goto L64
L75:
	;
	v299 = v175
	v300 = int32(_a_F_pg_size_bytes_5)
	goto L77
L76:
	;
	if v337 == int32(0) {
		v572 = int32(_a_F_pg_size_bytes_6)
		goto L61
	} else {
		goto L89
	}
L77:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	if v303 == v304 {
		v326 = v303
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v337 = int32(0)
	goto L76
L79:
	;
	v328 = int32(1)
	if v326 != 0 {
		v299 = v299 + v328
		v300 = v300 + v328
		goto L77
	} else {
		goto L88
	}
L80:
	;
	if base.Ui32((v303-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v314 = v303 | int32(32)
	goto L83
L82:
	;
	v314 = v303
	goto L83
L83:
	;
	if base.Ui32((v304-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v323 = v304 | int32(32)
	goto L86
L85:
	;
	v323 = v304
	goto L86
L86:
	;
	if v314 == v323 {
		v326 = v314
		goto L79
	} else {
		goto L87
	}
L87:
	;
	v337 = v314 - v323
	goto L76
L88:
	;
	goto L78
L89:
	;
	v344 = v175
	v345 = int32(_a_F_pg_size_bytes_7)
	goto L91
L90:
	;
	if v382 == int32(0) {
		v572 = int32(_a_F_pg_size_bytes_8)
		goto L61
	} else {
		goto L103
	}
L91:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344))))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345))))
	if v348 == v349 {
		v371 = v348
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v382 = int32(0)
	goto L90
L93:
	;
	v373 = int32(1)
	if v371 != 0 {
		v344 = v344 + v373
		v345 = v345 + v373
		goto L91
	} else {
		goto L102
	}
L94:
	;
	if base.Ui32((v348-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v359 = v348 | int32(32)
	goto L97
L96:
	;
	v359 = v348
	goto L97
L97:
	;
	if base.Ui32((v349-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v368 = v349 | int32(32)
	goto L100
L99:
	;
	v368 = v349
	goto L100
L100:
	;
	if v359 == v368 {
		v371 = v359
		goto L93
	} else {
		goto L101
	}
L101:
	;
	v382 = v359 - v368
	goto L90
L102:
	;
	goto L92
L103:
	;
	v389 = v175
	v390 = int32(_a_F_pg_size_bytes_9)
	goto L105
L104:
	;
	if v427 == int32(0) {
		v572 = int32(_a_F_pg_size_bytes_10)
		goto L61
	} else {
		goto L117
	}
L105:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389))))
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390))))
	if v393 == v394 {
		v416 = v393
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v427 = int32(0)
	goto L104
L107:
	;
	v418 = int32(1)
	if v416 != 0 {
		v389 = v389 + v418
		v390 = v390 + v418
		goto L105
	} else {
		goto L116
	}
L108:
	;
	if base.Ui32((v393-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v404 = v393 | int32(32)
	goto L111
L110:
	;
	v404 = v393
	goto L111
L111:
	;
	if base.Ui32((v394-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v413 = v394 | int32(32)
	goto L114
L113:
	;
	v413 = v394
	goto L114
L114:
	;
	if v404 == v413 {
		v416 = v404
		goto L107
	} else {
		goto L115
	}
L115:
	;
	v427 = v404 - v413
	goto L104
L116:
	;
	goto L106
L117:
	;
	v434 = v175
	v435 = int32(_a_F_pg_size_bytes_11)
	goto L119
L118:
	;
	if v472 == int32(0) {
		v572 = int32(_a_F_pg_size_bytes_12)
		goto L61
	} else {
		goto L131
	}
L119:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434))))
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435))))
	if v438 == v439 {
		v461 = v438
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v472 = int32(0)
	goto L118
L121:
	;
	v463 = int32(1)
	if v461 != 0 {
		v434 = v434 + v463
		v435 = v435 + v463
		goto L119
	} else {
		goto L130
	}
L122:
	;
	if base.Ui32((v438-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v449 = v438 | int32(32)
	goto L125
L124:
	;
	v449 = v438
	goto L125
L125:
	;
	if base.Ui32((v439-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v458 = v439 | int32(32)
	goto L128
L127:
	;
	v458 = v439
	goto L128
L128:
	;
	if v449 == v458 {
		v461 = v449
		goto L121
	} else {
		goto L129
	}
L129:
	;
	v472 = v449 - v458
	goto L118
L130:
	;
	goto L120
L131:
	;
	v479 = v175
	v480 = int32(_a_F_pg_size_bytes_13)
	goto L133
L132:
	;
	if v517 == int32(0) {
		v572 = int32(_a_F_pg_size_bytes_14)
		goto L61
	} else {
		goto L145
	}
L133:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	if v483 == v484 {
		v506 = v483
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v517 = int32(0)
	goto L132
L135:
	;
	v508 = int32(1)
	if v506 != 0 {
		v479 = v479 + v508
		v480 = v480 + v508
		goto L133
	} else {
		goto L144
	}
L136:
	;
	if base.Ui32((v483-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v494 = v483 | int32(32)
	goto L139
L138:
	;
	v494 = v483
	goto L139
L139:
	;
	if base.Ui32((v484-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v503 = v484 | int32(32)
	goto L142
L141:
	;
	v503 = v484
	goto L142
L142:
	;
	if v494 == v503 {
		v506 = v494
		goto L135
	} else {
		goto L143
	}
L143:
	;
	v517 = v494 - v503
	goto L132
L144:
	;
	goto L134
L145:
	;
	v525 = v175
	v526 = int32(_a_F_pg_size_bytes_15)
	goto L147
L146:
	;
	if v563 != 0 {
		goto L159
	} else {
		goto L160
	}
L147:
	;
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525))))
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526))))
	if v529 == v530 {
		v552 = v529
		goto L149
	} else {
		goto L150
	}
L148:
	;
	v563 = int32(0)
	goto L146
L149:
	;
	v554 = int32(1)
	if v552 != 0 {
		v525 = v525 + v554
		v526 = v526 + v554
		goto L147
	} else {
		goto L158
	}
L150:
	;
	if base.Ui32((v529-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v540 = v529 | int32(32)
	goto L153
L152:
	;
	v540 = v529
	goto L153
L153:
	;
	if base.Ui32((v530-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v549 = v530 | int32(32)
	goto L156
L155:
	;
	v549 = v530
	goto L156
L156:
	;
	if v540 == v549 {
		v552 = v540
		goto L149
	} else {
		goto L157
	}
L157:
	;
	v563 = v540 - v549
	goto L146
L158:
	;
	goto L148
L159:
	;
	v564 = int32(_a_F_pg_size_bytes_16)
	goto L161
L160:
	;
	v564 = int32(_a_F_pg_size_bytes_4)
	goto L161
L161:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)))
	if v565 == int32(0) {
		goto L44
	} else {
		goto L162
	}
L162:
	;
	if v563 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v570 = int32(_a_F_pg_size_bytes_16)
	goto L165
L164:
	;
	v570 = int32(_a_F_pg_size_bytes_4)
	goto L165
L165:
	;
	v572 = v570
	goto L61
L166:
	;
	v579 = F_int64_to_numeric(m, v574)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v581 = F_DirectFunctionCall2Coll(m, int32(1262), int32(0), v579, v171)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	v583 = F_pg_detoast_datum(m, v581)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	v587 = v583
	goto L45
L170:
	;
	v598 = *(*int64)(unsafe.Add(mBase, uint32(v596)))
	v599 = F_Int64GetDatum(m, v598)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	m.G0 = v12 + int32(48)
	return v599
L172:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v612 = F_text_to_cstring(m, v15)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v612
	F_errmsg(m, int32(_a_F_pg_size_bytes_0), v12+int32(16))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v175
	F_errdetail(m, int32(_a_F_pg_size_bytes_17), v12)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	F_errhint(m, int32(_a_F_pg_size_bytes_18), int32(0))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(_a_F_pg_size_bytes_1), int32(860), int32(_a_F_pg_size_bytes_2))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_snapshot_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v38 int64
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v81 int32
	_ = v81
	var v84 int64
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v143 int64
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int64
	_ = v150
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v161 int64
	_ = v161
	var v163 int64
	_ = v163
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_pq_begintypsend(m, v9)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	F_enlargeStringInfo(m, v9, int32(4))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v27 = int32(16711935)
	v31 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v22+v23))) = base.I32_rotr(v18, int32(24))&v27 | base.I32_rotr(v18&v27, v31)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v22 + int32(4)
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	F_enlargeStringInfo(m, v9, v31)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v45 = int64(56)
	v47 = int64(65280)
	v49 = int64(40)
	v52 = int64(16711680)
	v54 = int64(24)
	v56 = int64(4278190080)
	v58 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v42+v43))) = v38<<(uint(v45)%64) | v38&v47<<(uint(v49)%64) | (v38&v52<<(uint(v54)%64) | v38&v56<<(uint(v58)%64)) | (int64(base.Ui64(v38)>>(uint(v58)%64))&v56 | int64(base.Ui64(v38)>>(uint(v54)%64))&v52 | (int64(base.Ui64(v38)>>(uint(v49)%64))&v47 | int64(base.Ui64(v38)>>(uint(v45)%64))))
	v81 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v42 + v81
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	F_enlargeStringInfo(m, v9, v81)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v91 = int64(56)
	v93 = int64(65280)
	v95 = int64(40)
	v98 = int64(16711680)
	v100 = int64(24)
	v102 = int64(4278190080)
	v104 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v88+v89))) = v84<<(uint(v91)%64) | v84&v93<<(uint(v95)%64) | (v84&v98<<(uint(v100)%64) | v84&v102<<(uint(v104)%64)) | (int64(base.Ui64(v84)>>(uint(v104)%64))&v102 | int64(base.Ui64(v84)>>(uint(v100)%64))&v98 | (int64(base.Ui64(v84)>>(uint(v95)%64))&v93 | int64(base.Ui64(v84)>>(uint(v91)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v88 + int32(8)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v130 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v134 = int32(0)
	goto L10
L8:
	;
	goto L9
L9:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v201 << (uint(int32(2)) % 32)
	goto L14
L10:
	;
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(24)+v134<<(uint(int32(3))%32))))
	F_enlargeStringInfo(m, v9, int32(8))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v150 = int64(56)
	v152 = int64(65280)
	v154 = int64(40)
	v157 = int64(16711680)
	v159 = int64(24)
	v161 = int64(4278190080)
	v163 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v147+v148))) = v143<<(uint(v150)%64) | v143&v152<<(uint(v154)%64) | (v143&v157<<(uint(v159)%64) | v143&v161<<(uint(v163)%64)) | (int64(base.Ui64(v143)>>(uint(v163)%64))&v161 | int64(base.Ui64(v143)>>(uint(v159)%64))&v157 | (int64(base.Ui64(v143)>>(uint(v154)%64))&v152 | int64(base.Ui64(v143)>>(uint(v150)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v147 + int32(8)
	v190 = v134 + int32(1)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if base.Ui32(v190) < base.Ui32(v191) {
		v134 = v190
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	m.G0 = v9 + int32(16)
	return v200
}
func F_pg_snapshot_xmax(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int64)(unsafe.Add(mBase, uint32(v3)+16))
		v8 = F_Int64GetDatum(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_pg_snapshot_xmin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int64)(unsafe.Add(mBase, uint32(v3)+8))
		v8 = F_Int64GetDatum(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_pg_strcasecmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v5 = l0
	v6 = l1
	goto L1
L1:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v9 == v10 {
		v33 = v9
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return int32(0)
L3:
	;
	v35 = int32(1)
	if v33 != 0 {
		v5 = v5 + v35
		v6 = v6 + v35
		goto L1
	} else {
		goto L12
	}
L4:
	;
	if base.Ui32((v9-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v20 = v9 | int32(32)
	goto L7
L6:
	;
	v20 = v9
	goto L7
L7:
	;
	if base.Ui32((v10-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v29 = v10 | int32(32)
	goto L10
L9:
	;
	v29 = v10
	goto L10
L10:
	;
	if v20 == v29 {
		v33 = v20
		goto L3
	} else {
		goto L11
	}
L11:
	;
	return v20 - v29
L12:
	;
	goto L2
}
func F_pg_strncasecmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v6 = l0
	v7 = l1
	v8 = l2
	goto L1
L1:
	;
	if v8 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return int32(0)
L3:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v11 == v12 {
		v35 = v11
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	goto L2
L6:
	;
	v37 = int32(1)
	if v35 != 0 {
		v6 = v6 + v37
		v7 = v7 + v37
		v8 = v8 - v37
		goto L1
	} else {
		goto L15
	}
L7:
	;
	if base.Ui32((v11-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v22 = v11 | int32(32)
	goto L10
L9:
	;
	v22 = v11
	goto L10
L10:
	;
	if base.Ui32((v12-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = v12 | int32(32)
	goto L13
L12:
	;
	v31 = v12
	goto L13
L13:
	;
	if v22 == v31 {
		v35 = v22
		goto L6
	} else {
		goto L14
	}
L14:
	;
	return v22 - v31
L15:
	;
	goto L5
}
func F_pg_strsignal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	v4 = int32(_a_F_pg_strsignal_0)
	if base.Ui32(int32(-64)) <= base.Ui32(l0-int32(65)) {
		v9 = l0
		v10 = v4
		for {
			v13 = v10 + int32(1)
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			if v14 != 0 {
				v10 = v13
				continue
			} else {
			}
			v16 = v9 - int32(1)
			if v16 != 0 {
				v9 = v16
				v10 = v13
				continue
			} else {
				break
			}
			break
		}
		v18 = v13
	} else {
		v18 = v4
	}
	if v18 != 0 {
		v21 = v18
	} else {
		v21 = int32(_a_F_pg_strsignal_1)
	}
	return v21
}
func F_pg_strtoint64_safe(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int64
	_ = v46
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v72 int64
	_ = v72
	var v77 int64
	_ = v77
	var v83 int64
	_ = v83
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v159 int64
	_ = v159
	var v165 int32
	_ = v165
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v192 int64
	_ = v192
	var v193 int32
	_ = v193
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v232 int64
	_ = v232
	var v233 int32
	_ = v233
	var v252 int32
	_ = v252
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int64
	_ = v270
	var v272 int32
	_ = v272
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int64
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v333 int64
	_ = v333
	var v339 int64
	_ = v339
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v428 int64
	_ = v428
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v17 = base.B2i32(v15 == int32(45))
	v18 = l0 + v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	v21 = v19 - int32(48)
	if base.Ui32(int32(10)) <= base.Ui32(v21&int32(255)) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return v428
L2:
	;
	v428 = int64(0)
	goto L1
L3:
	;
	F_errsave_finish(m, l1, int32(_a_F_pg_strtoint64_safe_0), v406, int32(_a_F_pg_strtoint64_safe_1))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L82
	} else {
		goto L91
	}
L4:
	;
	v381 = F_errsave_start(m, l1)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L82
	} else {
		goto L87
	}
L5:
	;
	v356 = F_errsave_start(m, l1)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L82
	} else {
		goto L83
	}
L6:
	;
	v103 = l0
	v106 = v15
	goto L23
L7:
	;
	v28 = base.I64_extend_i32_u(v21) & int64(255)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	v31 = v29 - int32(48)
	if base.Ui32(v31&int32(255)) <= base.Ui32(int32(9)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v41 = v18 + int32(1)
	v42 = v31
	v46 = v28
	goto L11
L9:
	;
	v66 = v29
	v72 = v28
	goto L10
L10:
	;
	if v66&int32(255) != 0 {
		goto L6
	} else {
		goto L15
	}
L11:
	;
	if base.Ui64(int64(922337203685477580)) < base.Ui64(v46) {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v66 = v55
	v72 = v54
	goto L10
L13:
	;
	v54 = v46*int64(10) + base.I64_extend_i32_u(v42)&int64(255)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	v59 = v55 - int32(48)
	if base.Ui32(v59&int32(255)) < base.Ui32(int32(10)) {
		v41 = v41 + int32(1)
		v42 = v59
		v46 = v54
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	if v15 == int32(45) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v77 = int64(0)
	v83 = v77 - v72
	if v77-base.I64_extend_i32_u(base.B2i32(v72 != v77))^v83>>(uint(int64(63))%64) != v77 {
		goto L5
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v72 < int64(0) {
		goto L5
	} else {
		goto L20
	}
L19:
	;
	v428 = v83
	goto L1
L20:
	;
	v428 = v72
	goto L1
L21:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v125 != int32(48) {
		goto L29
	} else {
		goto L30
	}
L22:
	;
	v123 = v103 + int32(1)
	v124 = v17
	goto L21
L23:
	;
	if base.Ui32(v106-int32(9)) < base.Ui32(int32(5)) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v118 = int32(1)
	v123 = v103 + v118
	v124 = v118
	goto L21
L25:
	;
	goto L24
L26:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	v103 = v103 + int32(1)
	v106 = v115
	goto L23
L27:
	;
	switch v106 - int32(32) {
	case 0:
		goto L26
	default:
		v123 = v103
		v124 = v17
		goto L21
	case 11:
		goto L22
	case 13:
		goto L25
	}
L28:
	;
	if v304 == v305 {
		goto L4
	} else {
		goto L70
	}
L29:
	;
	v264 = v123
	v266 = v125
	v270 = int64(0)
	goto L61
L30:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
	switch v128 - int32(66) {
	case 0, 32:
		goto L31
	default:
		goto L29
	case 13, 45:
		goto L32
	case 22, 54:
		goto L33
	}
L31:
	;
	v223 = v123 + int32(2)
	v226 = v223
	v232 = int64(0)
	goto L53
L32:
	;
	v183 = v123 + int32(2)
	v186 = v183
	v192 = int64(0)
	goto L45
L33:
	;
	v133 = v123 + int32(2)
	v136 = v133
	v142 = int64(0)
	goto L34
L34:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	goto L36
L35:
	;
	goto L4
L36:
	;
	if base.B2i32(base.Ui32(v143-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v143|int32(32)-int32(97)) < base.Ui32(int32(6))) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if base.Ui64(int64(576460752303423488)) < base.Ui64(v142) {
		goto L5
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if v143 != int32(95) {
		v304 = v136
		v305 = v133
		v306 = v143
		v310 = v142
		goto L28
	} else {
		goto L41
	}
L40:
	;
	v159 = int64(*(*int8)(unsafe.Add(mBase, uint32(v143)+uint32(_c_F_pg_strtoint64_safe[0]))))
	v136 = v136 + int32(1)
	v142 = v159 + v142<<(uint(int64(4))%64)
	goto L34
L41:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	if v165 == int32(0) {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	goto L43
L43:
	;
	if base.B2i32(base.Ui32(v165-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v165|int32(32)-int32(97)) < base.Ui32(int32(6))) != 0 {
		v136 = v136 + int32(1)
		goto L34
	} else {
		goto L44
	}
L44:
	;
	goto L35
L45:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v193&int32(248) == int32(48) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L4
L47:
	;
	if base.Ui64(int64(1152921504606846976)) < base.Ui64(v192) {
		goto L5
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v193 != int32(95) {
		v304 = v186
		v305 = v183
		v306 = v193
		v310 = v192
		goto L28
	} else {
		goto L51
	}
L50:
	;
	v186 = v186 + int32(1)
	v192 = base.I64_extend_i32_u(v193-int32(48))&int64(255) | v192<<(uint(int64(3))%64)
	goto L45
L51:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+1)))
	if base.Ui32(int32(248)) <= base.Ui32((v212-int32(56))&int32(255)) {
		v186 = v186 + int32(1)
		goto L45
	} else {
		goto L52
	}
L52:
	;
	goto L46
L53:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	if v233&int32(254) == int32(48) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L4
L55:
	;
	if base.Ui64(int64(4611686018427387904)) < base.Ui64(v232) {
		goto L5
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	if v233 != int32(95) {
		v304 = v226
		v305 = v223
		v306 = v233
		v310 = v232
		goto L28
	} else {
		goto L59
	}
L58:
	;
	v226 = v226 + int32(1)
	v232 = base.I64_extend_i32_u(v233-int32(48))&int64(255) | v232<<(uint(int64(1))%64)
	goto L53
L59:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+1)))
	if base.Ui32(int32(254)) <= base.Ui32((v252-int32(50))&int32(255)) {
		v226 = v226 + int32(1)
		goto L53
	} else {
		goto L60
	}
L60:
	;
	goto L54
L61:
	;
	v272 = v266 - int32(48)
	if base.Ui32(v272&int32(255)) <= base.Ui32(int32(9)) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L4
L63:
	;
	if base.Ui64(int64(922337203685477580)) < base.Ui64(v270) {
		goto L5
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v266&int32(255) != int32(95) {
		v304 = v264
		v305 = v123
		v306 = v266
		v310 = v270
		goto L28
	} else {
		goto L67
	}
L66:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+1)))
	v264 = v264 + int32(1)
	v266 = v285
	v270 = v270*int64(10) + base.I64_extend_i32_u(v272)&int64(255)
	goto L61
L67:
	;
	if v264 == v123 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+1)))
	if base.Ui32((v293-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v264 = v264 + int32(1)
		v266 = v293
		goto L61
	} else {
		goto L69
	}
L69:
	;
	goto L62
L70:
	;
	v314 = v304
	v316 = v306
	goto L71
L71:
	;
	v322 = v316 & int32(255)
	if base.B2i32(base.Ui32(v322-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v322 == int32(32)) != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L5
L73:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+1)))
	v314 = v314 + int32(1)
	v316 = v330
	goto L71
L74:
	;
	if v322 != 0 {
		goto L4
	} else {
		goto L76
	}
L75:
	;
	goto L72
L76:
	;
	if v124 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v333 = int64(0)
	v339 = v333 - v310
	if v333-base.I64_extend_i32_u(base.B2i32(v310 != v333))^v339>>(uint(int64(63))%64) != v333 {
		goto L5
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if int64(0) <= v310 {
		v428 = v310
		goto L1
	} else {
		goto L81
	}
L80:
	;
	v428 = v339
	goto L1
L81:
	;
	goto L75
L82:
	;
	return int64(0)
L83:
	;
	if v356 == int32(0) {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(_a_F_pg_strtoint64_safe_2)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg(m, int32(_a_F_pg_strtoint64_safe_3), v12)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v406 = int32(873)
	goto L3
L87:
	;
	if v381 == int32(0) {
		goto L2
	} else {
		goto L88
	}
L88:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L82
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(_a_F_pg_strtoint64_safe_2)
	F_errmsg(m, int32(_a_F_pg_strtoint64_safe_4), v12+int32(16))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L82
	} else {
		goto L90
	}
L90:
	;
	v406 = int32(879)
	goto L3
L91:
	;
	goto L2
}
func F_pg_strtok(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_pg_strtok[0]))
	v9 = v7
	goto L5
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_strtok[0])) = v59
	return v58
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	v58 = v54
	v59 = v52
	goto L1
L3:
	;
	v40 = v37 - v9
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v40
	if v40 != int32(2) {
		v58 = v9
		v59 = v37
		goto L1
	} else {
		goto L19
	}
L4:
	;
	v37 = v9 + int32(1)
	goto L3
L5:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	switch v13 {
	case 0:
		v52 = v9
		v54 = int32(0)
		goto L2
	case 1, 2, 3, 4, 5, 6, 7, 8, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 33, 34, 35, 36, 37, 38, 39:
		goto L7
	case 9, 10, 32:
		goto L8
	case 40, 41:
		goto L4
	default:
		goto L9
	}
L6:
	;
	v20 = v9
	v21 = v13
	goto L10
L7:
	;
	goto L6
L8:
	;
	v9 = v9 + int32(1)
	goto L5
L9:
	;
	switch v13 - int32(123) {
	case 0, 2:
		goto L4
	default:
		goto L7
	}
L10:
	;
	if v21 != int32(92) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v31 = v30 + v20
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v20 = v31
	v21 = v32
	goto L10
L13:
	;
	v30 = int32(1)
	goto L12
L14:
	;
	switch v21 {
	case 0, 9, 10, 32, 40, 41:
		v37 = v20
		goto L3
	case 1, 2, 3, 4, 5, 6, 7, 8, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 33, 34, 35, 36, 37, 38, 39:
		goto L13
	default:
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v28 != 0 {
		v30 = int32(2)
		goto L12
	} else {
		goto L18
	}
L17:
	;
	switch v21 - int32(123) {
	case 0, 2:
		v37 = v20
		goto L3
	default:
		goto L13
	}
L18:
	;
	goto L13
L19:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v44 != int32(60) {
		v58 = v9
		v59 = v37
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v47 != int32(62) {
		v58 = v9
		v59 = v37
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v52 = v37
	v54 = v9
	goto L2
}
func F_pg_tablespace_size_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_SearchSysCacheExists(m, int32(69), v10, v2, v2, v2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			v18 = F_calculate_tablespace_size(m, v10)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 < int64(0) {
					v22 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
					v27 = int32(0)
					m.G0 = v7 + int32(16)
					return v27
				} else {
					v25 = F_Int64GetDatum(m, v18)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = v25
						m.G0 = v7 + int32(16)
						return v27
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10
					F_errmsg(m, int32(_a_F_pg_tablespace_size_oid_0), v7)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_tablespace_size_oid_1), int32(293), int32(_a_F_pg_tablespace_size_oid_2))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
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
func F_pg_timezone_abbrev_is_known(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	v6 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l4)+268))
	if v11 <= v6 {
		v112 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v112
L2:
	;
	v24 = v6
	goto L3
L3:
	;
	v26 = l4 + int32(_a_F_pg_timezone_abbrev_is_known_0) + v24
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if base.B2i32(v29 == int32(0))|base.B2i32(v29 != v32) != 0 {
		v50 = v29
		v51 = v32
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l4)+264))
	if v58 <= int32(0) {
		v112 = v6
		goto L1
	} else {
		goto L16
	}
L5:
	;
	if v50-v51 != 0 {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	goto L5
L7:
	;
	v35 = l0
	v36 = v26
	goto L8
L8:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	if v40 == int32(0) {
		v50 = v40
		v51 = v39
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v50 = v40
	v51 = v39
	goto L6
L10:
	;
	v43 = int32(1)
	if v40 == v39 {
		v35 = v35 + v43
		v36 = v36 + v43
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v53 = F_strlen(m, v26)
	mBase = m.M
	v56 = v53 + v24 + int32(1)
	if v56 < v11 {
		v24 = v56
		goto L3
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	goto L4
L15:
	;
	v112 = v6
	goto L1
L16:
	;
	v69 = int32(0)
	v70 = v58
	v71 = v6
	goto L17
L17:
	;
	v76 = l4 + int32(_a_F_pg_timezone_abbrev_is_known_1) + v69<<(uint(int32(4))%32)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	if v77 != v24 {
		v100 = v70
		v101 = v71
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v112 = v101
	goto L1
L19:
	;
	v103 = v69 + int32(1)
	if v103 < v100 {
		v69 = v103
		v70 = v100
		v71 = v101
		goto L17
	} else {
		goto L28
	}
L20:
	;
	if v71 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v81 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v81)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v84
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l4)+264))
	v100 = v88
	v101 = v81
	goto L19
L22:
	;
	goto L23
L23:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v89 == v90 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+4)))
	if v93 == v94 {
		v100 = v70
		v101 = int32(1)
		goto L19
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v97 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v97)
	v112 = int32(1)
	goto L1
L27:
	;
	goto L26
L28:
	;
	goto L18
}
func F_pg_ts_parser_is_visible(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_TSParserIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v23 = v2
		} else {
			v23 = v14
		}
		m.G0 = v7 + int32(16)
		return v23
	}
}
func F_pg_utf8_increment(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	switch l1 - int32(1) {
	case 0:
		v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		v43 = v42
		v46 = v43 & int32(255)
		v48 = v46 - int32(223)
		if int32(1)<<(uint(v48)%32)&int32(_a_F_pg_utf8_increment_0) != 0 {
			v56 = base.B2i32(base.Ui32(v48) <= base.Ui32(int32(21)))
		} else {
			v56 = int32(0)
		}
		if v56|base.B2i32(v46 == int32(127)) != 0 {
			return int32(0)
		} else {
			v60 = int32(1)
			v61 = v43 + v60
			*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v61)
			return v60
		}
	case 1:
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		if v29 == int32(244) {
			v32 = int32(143)
		} else {
			v32 = int32(191)
		}
		if v29 == int32(237) {
			v35 = int32(159)
		} else {
			v35 = v32
		}
		if base.Ui32(v35) <= base.Ui32(v25) {
			v43 = v29
			v46 = v43 & int32(255)
			v48 = v46 - int32(223)
			if int32(1)<<(uint(v48)%32)&int32(_a_F_pg_utf8_increment_0) != 0 {
				v56 = base.B2i32(base.Ui32(v48) <= base.Ui32(int32(21)))
			} else {
				v56 = int32(0)
			}
			if v56|base.B2i32(v46 == int32(127)) != 0 {
				return int32(0)
			} else {
				v60 = int32(1)
				v61 = v43 + v60
				*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v61)
				return v60
			}
		} else {
			v37 = int32(1)
			v38 = v25 + v37
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v38)
			return v37
		}
	case 2:
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
		if base.Ui32(int32(190)) < base.Ui32(v16) {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v29 == int32(244) {
				v32 = int32(143)
			} else {
				v32 = int32(191)
			}
			if v29 == int32(237) {
				v35 = int32(159)
			} else {
				v35 = v32
			}
			if base.Ui32(v35) <= base.Ui32(v25) {
				v43 = v29
				v46 = v43 & int32(255)
				v48 = v46 - int32(223)
				if int32(1)<<(uint(v48)%32)&int32(_a_F_pg_utf8_increment_0) != 0 {
					v56 = base.B2i32(base.Ui32(v48) <= base.Ui32(int32(21)))
				} else {
					v56 = int32(0)
				}
				if v56|base.B2i32(v46 == int32(127)) != 0 {
					return int32(0)
				} else {
					v60 = int32(1)
					v61 = v43 + v60
					*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v61)
					return v60
				}
			} else {
				v37 = int32(1)
				v38 = v25 + v37
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v38)
				return v37
			}
		} else {
			v19 = int32(1)
			v20 = v16 + v19
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v20)
			return v19
		}
	case 3:
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
		if base.Ui32(int32(190)) < base.Ui32(v7) {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
			if base.Ui32(int32(190)) < base.Ui32(v16) {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				if v29 == int32(244) {
					v32 = int32(143)
				} else {
					v32 = int32(191)
				}
				if v29 == int32(237) {
					v35 = int32(159)
				} else {
					v35 = v32
				}
				if base.Ui32(v35) <= base.Ui32(v25) {
					v43 = v29
					v46 = v43 & int32(255)
					v48 = v46 - int32(223)
					if int32(1)<<(uint(v48)%32)&int32(_a_F_pg_utf8_increment_0) != 0 {
						v56 = base.B2i32(base.Ui32(v48) <= base.Ui32(int32(21)))
					} else {
						v56 = int32(0)
					}
					if v56|base.B2i32(v46 == int32(127)) != 0 {
						return int32(0)
					} else {
						v60 = int32(1)
						v61 = v43 + v60
						*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v61)
						return v60
					}
				} else {
					v37 = int32(1)
					v38 = v25 + v37
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v38)
					return v37
				}
			} else {
				v19 = int32(1)
				v20 = v16 + v19
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v20)
				return v19
			}
		} else {
			v10 = int32(1)
			v11 = v7 + v10
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v11)
			return v10
		}
	default:
		return int32(0)
	}
}
func F_pg_utf8_islegal(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	v3 = int32(0)
	switch l1 - int32(1) {
	case 0:
		v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		v39 = v38
		if base.I32_extend8_s(v39) < int32(-62) {
			v52 = v3
		} else {
			v44 = v39
			v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
		}
	case 1:
		v12 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		switch v13 - int32(224) {
		case 0:
			v16 = int32(224)
			if base.Ui32(v16) <= base.Ui32((v12-int32(-64))&int32(255)) {
				v44 = v16
				v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
			} else {
				v52 = v3
			}
		default:
			if v12 <= int32(-65) {
				v39 = v13
				if base.I32_extend8_s(v39) < int32(-62) {
					v52 = v3
				} else {
					v44 = v39
					v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
				}
			} else {
				v52 = v3
			}
		case 13:
			if int32(-97) < v12 {
				v52 = v3
			} else {
				v44 = int32(237)
				v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
			}
		case 16:
			if base.Ui32((v12-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
				v52 = v3
			} else {
				v44 = int32(240)
				v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
			}
		case 20:
			if int32(-113) < v12 {
				v52 = v3
			} else {
				v44 = int32(244)
				v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
			}
		}
	case 2:
		v9 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
		if int32(-65) < v9 {
			v52 = v3
		} else {
			v12 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			switch v13 - int32(224) {
			case 0:
				v16 = int32(224)
				if base.Ui32(v16) <= base.Ui32((v12-int32(-64))&int32(255)) {
					v44 = v16
					v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
				} else {
					v52 = v3
				}
			default:
				if v12 <= int32(-65) {
					v39 = v13
					if base.I32_extend8_s(v39) < int32(-62) {
						v52 = v3
					} else {
						v44 = v39
						v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
					}
				} else {
					v52 = v3
				}
			case 13:
				if int32(-97) < v12 {
					v52 = v3
				} else {
					v44 = int32(237)
					v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
				}
			case 16:
				if base.Ui32((v12-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
					v52 = v3
				} else {
					v44 = int32(240)
					v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
				}
			case 20:
				if int32(-113) < v12 {
					v52 = v3
				} else {
					v44 = int32(244)
					v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
				}
			}
		}
	case 3:
		v6 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+3)))
		if int32(-65) < v6 {
			v52 = v3
		} else {
			v9 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
			if int32(-65) < v9 {
				v52 = v3
			} else {
				v12 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
				v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				switch v13 - int32(224) {
				case 0:
					v16 = int32(224)
					if base.Ui32(v16) <= base.Ui32((v12-int32(-64))&int32(255)) {
						v44 = v16
						v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
					} else {
						v52 = v3
					}
				default:
					if v12 <= int32(-65) {
						v39 = v13
						if base.I32_extend8_s(v39) < int32(-62) {
							v52 = v3
						} else {
							v44 = v39
							v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
						}
					} else {
						v52 = v3
					}
				case 13:
					if int32(-97) < v12 {
						v52 = v3
					} else {
						v44 = int32(237)
						v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
					}
				case 16:
					if base.Ui32((v12-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
						v52 = v3
					} else {
						v44 = int32(240)
						v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
					}
				case 20:
					if int32(-113) < v12 {
						v52 = v3
					} else {
						v44 = int32(244)
						v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
					}
				}
			}
		}
	default:
		v52 = v3
	}
	return v52
}
func F_pg_utf8_verifychar(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v6 = base.I32_extend8_s(v5)
	if int32(0) <= v6 {
		if v6 != 0 {
			v11 = int32(1)
		} else {
			v11 = int32(-1)
		}
		return v11
	} else {
		if v5&int32(224) == int32(192) {
			v30 = int32(2)
		} else {
			if v5&int32(240) == int32(224) {
				v30 = int32(3)
			} else {
				if v5&int32(248) == int32(240) {
					v29 = int32(4)
				} else {
					v29 = int32(1)
				}
				v30 = v29
			}
		}
		if v30 <= l1 {
			v33 = int32(0)
			switch v30 - int32(1) {
			case 0:
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				v69 = v68
				if base.I32_extend8_s(v69) < int32(-62) {
					v82 = v33
				} else {
					v74 = v69
					v82 = base.B2i32(base.Ui32(v74&int32(255)) < base.Ui32(int32(245)))
				}
			case 1:
				v42 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				switch v43 - int32(224) {
				case 0:
					v46 = int32(224)
					if base.Ui32(v46) <= base.Ui32((v42-int32(-64))&int32(255)) {
						v74 = v46
						v82 = base.B2i32(base.Ui32(v74&int32(255)) < base.Ui32(int32(245)))
					} else {
						v82 = v33
					}
				default:
					if v42 <= int32(-65) {
						v69 = v43
						if base.I32_extend8_s(v69) < int32(-62) {
							v82 = v33
						} else {
							v74 = v69
							v82 = base.B2i32(base.Ui32(v74&int32(255)) < base.Ui32(int32(245)))
						}
					} else {
						v82 = v33
					}
				case 13:
					if int32(-97) < v42 {
						v82 = v33
					} else {
						v74 = int32(237)
						v82 = base.B2i32(base.Ui32(v74&int32(255)) < base.Ui32(int32(245)))
					}
				case 16:
					if base.Ui32((v42-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
						v82 = v33
					} else {
						v74 = int32(240)
						v82 = base.B2i32(base.Ui32(v74&int32(255)) < base.Ui32(int32(245)))
					}
				case 20:
					if int32(-113) < v42 {
						v82 = v33
					} else {
						v74 = int32(244)
						v82 = base.B2i32(base.Ui32(v74&int32(255)) < base.Ui32(int32(245)))
					}
				}
			case 2:
				v39 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
				if int32(-65) < v39 {
					v82 = v33
				} else {
					v42 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					switch v43 - int32(224) {
					case 0:
						v46 = int32(224)
						if base.Ui32(v46) <= base.Ui32((v42-int32(-64))&int32(255)) {
							v74 = v46
							v82 = base.B2i32(base.Ui32(v74&int32(255)) < base.Ui32(int32(245)))
						} else {
							v82 = v33
						}
					default:
						if v42 <= int32(-65) {
							v69 = v43
							if base.I32_extend8_s(v69) < int32(-62) {
								v82 = v33
							} else {
								v74 = v69
								v82 = base.B2i32(base.Ui32(v74&int32(255)) < base.Ui32(int32(245)))
							}
						} else {
							v82 = v33
						}
					case 13:
						if int32(-97) < v42 {
							v82 = v33
						} else {
							v74 = int32(237)
							v82 = base.B2i32(base.Ui32(v74&int32(255)) < base.Ui32(int32(245)))
						}
					case 16:
						if base.Ui32((v42-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
							v82 = v33
						} else {
							v74 = int32(240)
							v82 = base.B2i32(base.Ui32(v74&int32(255)) < base.Ui32(int32(245)))
						}
					case 20:
						if int32(-113) < v42 {
							v82 = v33
						} else {
							v74 = int32(244)
							v82 = base.B2i32(base.Ui32(v74&int32(255)) < base.Ui32(int32(245)))
						}
					}
				}
			case 3:
				v36 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+3)))
				if int32(-65) < v36 {
					v82 = v33
				} else {
					v39 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
					if int32(-65) < v39 {
						v82 = v33
					} else {
						v42 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
						v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
						switch v43 - int32(224) {
						case 0:
							v46 = int32(224)
							if base.Ui32(v46) <= base.Ui32((v42-int32(-64))&int32(255)) {
								v74 = v46
								v82 = base.B2i32(base.Ui32(v74&int32(255)) < base.Ui32(int32(245)))
							} else {
								v82 = v33
							}
						default:
							if v42 <= int32(-65) {
								v69 = v43
								if base.I32_extend8_s(v69) < int32(-62) {
									v82 = v33
								} else {
									v74 = v69
									v82 = base.B2i32(base.Ui32(v74&int32(255)) < base.Ui32(int32(245)))
								}
							} else {
								v82 = v33
							}
						case 13:
							if int32(-97) < v42 {
								v82 = v33
							} else {
								v74 = int32(237)
								v82 = base.B2i32(base.Ui32(v74&int32(255)) < base.Ui32(int32(245)))
							}
						case 16:
							if base.Ui32((v42-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
								v82 = v33
							} else {
								v74 = int32(240)
								v82 = base.B2i32(base.Ui32(v74&int32(255)) < base.Ui32(int32(245)))
							}
						case 20:
							if int32(-113) < v42 {
								v82 = v33
							} else {
								v74 = int32(244)
								v82 = base.B2i32(base.Ui32(v74&int32(255)) < base.Ui32(int32(245)))
							}
						}
					}
				}
			default:
				v82 = v33
			}
			if v82 != 0 {
				v83 = v30
			} else {
				v83 = int32(-1)
			}
			v85 = v83
		} else {
			v85 = int32(-1)
		}
		return v85
	}
}
func F_pg_xact_status(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int64
	_ = v40
	var v47 int64
	_ = v47
	var v52 int64
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_pg_xact_status[0]))
	v20 = F_LWLockAcquire(m, v16+int32(_a_F_pg_xact_status_0), int32(1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		v24 = F_ReadNextFullTransactionId(m)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = base.I32_wrap_i64(v14)
			if v26 == int32(0) {
				v94 = *(*int32)(unsafe.Add(mBase, _c_F_pg_xact_status[0]))
				F_LWLockRelease(m, v94+int32(_a_F_pg_xact_status_0))
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int32(0)
				} else {
					v99 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v99)
					v104 = int32(0)
					m.G0 = v11 + int32(16)
					return v104
				}
			} else {
				if base.Ui32(int32(3)) <= base.Ui32(v26) {
					if base.Ui64(v24) <= base.Ui64(v14) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v11))) = v14
								F_errmsg(m, int32(_a_F_pg_xact_status_1), v11)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_xact_status_2), int32(121), int32(_a_F_pg_xact_status_3))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, _c_F_pg_xact_status[1]))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+64))
						if base.Ui32(v34) <= base.Ui32(int32(2)) {
							v52 = base.I64_extend_i32_u(v34)
						} else {
							v40 = int64(base.Ui64(v24) >> (uint(int64(32)) % 64))
							if base.Ui32(base.I32_wrap_i64(v24)) < base.Ui32(v34) {
								v47 = (v40 - int64(1)) & int64(4294967295)
							} else {
								v47 = v40
							}
							v52 = base.I64_extend_i32_u(v34) | v47<<(uint(int64(32))%64)
						}
						if base.Ui64(v14) < base.Ui64(v52) {
							v94 = *(*int32)(unsafe.Add(mBase, _c_F_pg_xact_status[0]))
							F_LWLockRelease(m, v94+int32(_a_F_pg_xact_status_0))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int32(0)
							} else {
								v99 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v99)
								v104 = int32(0)
								m.G0 = v11 + int32(16)
								return v104
							}
						} else {
							v57 = F_TransactionIdIsInProgress(m, v26)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								if v57 == int32(0) {
									v63 = F_TransactionIdDidCommit(m, v26)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										if v63 != 0 {
											v65 = int32(_a_F_pg_xact_status_4)
										} else {
											v65 = int32(_a_F_pg_xact_status_5)
										}
										v66 = v65
										v68 = *(*int32)(unsafe.Add(mBase, _c_F_pg_xact_status[0]))
										F_LWLockRelease(m, v68+int32(_a_F_pg_xact_status_0))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											v73 = F_cstring_to_text(m, v66)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												v104 = v73
												m.G0 = v11 + int32(16)
												return v104
											}
										}
									}
								} else {
									v66 = int32(_a_F_pg_xact_status_6)
									v68 = *(*int32)(unsafe.Add(mBase, _c_F_pg_xact_status[0]))
									F_LWLockRelease(m, v68+int32(_a_F_pg_xact_status_0))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int32(0)
									} else {
										v73 = F_cstring_to_text(m, v66)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											v104 = v73
											m.G0 = v11 + int32(16)
											return v104
										}
									}
								}
							}
						}
					}
				} else {
					v57 = F_TransactionIdIsInProgress(m, v26)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						if v57 == int32(0) {
							v63 = F_TransactionIdDidCommit(m, v26)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								if v63 != 0 {
									v65 = int32(_a_F_pg_xact_status_4)
								} else {
									v65 = int32(_a_F_pg_xact_status_5)
								}
								v66 = v65
								v68 = *(*int32)(unsafe.Add(mBase, _c_F_pg_xact_status[0]))
								F_LWLockRelease(m, v68+int32(_a_F_pg_xact_status_0))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									v73 = F_cstring_to_text(m, v66)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										v104 = v73
										m.G0 = v11 + int32(16)
										return v104
									}
								}
							}
						} else {
							v66 = int32(_a_F_pg_xact_status_6)
							v68 = *(*int32)(unsafe.Add(mBase, _c_F_pg_xact_status[0]))
							F_LWLockRelease(m, v68+int32(_a_F_pg_xact_status_0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								v73 = F_cstring_to_text(m, v66)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									v104 = v73
									m.G0 = v11 + int32(16)
									return v104
								}
							}
						}
					}
				}
			}
		}
	}
}
