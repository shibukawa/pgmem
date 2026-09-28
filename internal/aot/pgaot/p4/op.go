package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_op_index_interpretation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int64
	_ = v9
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int64
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	v2 = int32(0)
	v9 = int64(0)
	v12 = base.I64_extend_i32_u(l0)
	v15 = F_SearchSysCacheList(m, int32(3), int32(1), v12, v9, v9)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v194
L2:
	;
	v102 = int32(0)
	v104 = F_SearchSysCache1(m, int32(40), v12)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L29
	}
L3:
	;
	return int32(0)
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	if v19 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_ReleaseCatCacheList(m, v15)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v28 = v2
	v30 = v2
	goto L9
L8:
	;
	goto L2
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v15-int32(-64)+v28<<(uint(int32(2))%32))))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+22)))
	v41 = v39 + v40
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	if v42 <= int32(2741) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	F_ReleaseCatCacheList(m, v15)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L3
	} else {
		goto L27
	}
L11:
	;
	v88 = v28 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	if v88 < v89 {
		v28 = v88
		v30 = v84
		goto L9
	} else {
		goto L26
	}
L12:
	;
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+16)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v67 = F_IndexAmTranslateStrategy(m, v65, v64, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L3
	} else {
		goto L22
	}
L13:
	;
	v58 = F_GetIndexAmRoutineByAmId(m, v42, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L3
	} else {
		goto L20
	}
L14:
	;
	switch v42 - int32(403) {
	case 0:
		v64 = v42
		goto L12
	case 1:
		goto L13
	case 2:
		v84 = v30
		goto L11
	default:
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if base.B2i32(v42 == int32(2742))|base.B2i32(v42 == int32(3580))|base.B2i32(v42 == int32(4000)) != 0 {
		v84 = v30
		goto L11
	} else {
		goto L19
	}
L17:
	;
	if v42 != int32(783) {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v84 = v30
	goto L11
L19:
	;
	goto L13
L20:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+10)))
	if v60 != int32(1) {
		v84 = v30
		goto L11
	} else {
		goto L21
	}
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	v64 = v63
	goto L12
L22:
	;
	if v67 == int32(0) {
		v84 = v30
		goto L11
	} else {
		goto L23
	}
L23:
	;
	v72 = F_palloc(m, int32(16))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v74
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = v79
	v81 = F_lappend(m, v30, v72)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v84 = v81
	goto L11
L26:
	;
	goto L10
L27:
	;
	if v84 != 0 {
		v194 = v84
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L2
L29:
	;
	if v104 == int32(0) {
		v194 = v102
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+22)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v108+v109)+96))
	F_ReleaseCatCache(m, v104)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	if v111 == int32(0) {
		v194 = v102
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v119 = int64(0)
	v121 = F_SearchSysCacheList(m, int32(3), int32(1), base.I64_extend_i32_u(v111), v119, v119)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v121)+56))
	if int32(0) < v123 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v131 = int32(0)
	v133 = v102
	goto L37
L35:
	;
	v183 = v102
	goto L36
L36:
	;
	F_ReleaseCatCacheList(m, v121)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L3
	} else {
		goto L47
	}
L37:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v121-int32(-64)+v131<<(uint(int32(2))%32))))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+72))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+22)))
	v144 = v142 + v143
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+24))
	v147 = F_GetIndexAmRoutineByAmId(m, v145, int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L3
	} else {
		goto L40
	}
L38:
	;
	v183 = v173
	goto L36
L39:
	;
	v176 = v131 + int32(1)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v121)+56))
	if v176 < v177 {
		v131 = v176
		v133 = v173
		goto L37
	} else {
		goto L46
	}
L40:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+10)))
	if v149 != int32(1) {
		v173 = v133
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144)+16)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v144)+24))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	v155 = F_IndexAmTranslateStrategy(m, v152, v153, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	if v155 != int32(3) {
		v173 = v133
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v160 = F_palloc(m, int32(16))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v162
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+8)) = v166
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+12)) = v168
	v170 = F_lappend(m, v133, v160)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	v173 = v170
	goto L39
L46:
	;
	goto L38
L47:
	;
	v194 = v183
	goto L1
}
func F_make_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	if l3 != 0 {
		if l2 == int32(0) {
			v20 = F_exprType(m, l3)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v25 = F_left_oper(m, l0, l1, v20, int32(0), l5)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
					v29 = v27 + v28
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+100))
					if v30 == int32(0) {
						v198 = v29
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v205 = m.ExcPending
						if v205 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(52461700))
							mBase = m.M
							v208 = m.ExcPending
							if v208 != 0 {
								return int32(0)
							} else {
								v209 = *(*int32)(unsafe.Add(mBase, uint32(v198)+80))
								v210 = *(*int32)(unsafe.Add(mBase, uint32(v198)+84))
								v211 = F_op_signature_string(m, l1, v209, v210)
								mBase = m.M
								v212 = m.ExcPending
								if v212 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v16))) = v211
									F_errmsg(m, int32(_a_F_make_op_0), v16)
									mBase = m.M
									v216 = m.ExcPending
									if v216 != 0 {
										return int32(0)
									} else {
										F_parser_errposition(m, l0, l5)
										mBase = m.M
										v218 = m.ExcPending
										if v218 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_make_op_1), int32(757), int32(_a_F_make_op_2))
											mBase = m.M
											v223 = m.ExcPending
											if v223 != 0 {
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
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = l3
						v39 = int32(1)
						v43 = F_list_make1_impl(m, v39, v16+int32(40))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v20
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v29)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v46
							v82 = v29 + int32(100)
							v83 = v39
							v84 = v29
							v85 = v25
							v87 = v43
							v88 = v25 + int32(16)
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v84)+88))
							v96 = F_enforce_generic_type_consistency(m, v16+int32(72), v16-int32(-64), v83, v94, int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
								if v98 == int32(2281) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v159 = m.ExcPending
									if v159 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(1088))
										mBase = m.M
										v162 = m.ExcPending
										if v162 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(_a_F_make_op_3)
											F_errmsg(m, int32(_a_F_make_op_4), v16+int32(16))
											mBase = m.M
											v169 = m.ExcPending
											if v169 != 0 {
												return int32(0)
											} else {
												F_parser_errposition(m, l0, l5)
												mBase = m.M
												v171 = m.ExcPending
												if v171 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_make_op_1), int32(803), int32(_a_F_make_op_2))
													mBase = m.M
													v176 = m.ExcPending
													if v176 != 0 {
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
								} else {
									if l2 != 0 {
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
										if v101 == int32(2281) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v159 = m.ExcPending
											if v159 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(1088))
												mBase = m.M
												v162 = m.ExcPending
												if v162 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(_a_F_make_op_3)
													F_errmsg(m, int32(_a_F_make_op_4), v16+int32(16))
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
														return int32(0)
													} else {
														F_parser_errposition(m, l0, l5)
														mBase = m.M
														v171 = m.ExcPending
														if v171 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_make_op_1), int32(803), int32(_a_F_make_op_2))
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
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
										} else {
											if v96 == int32(2281) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = int32(_a_F_make_op_3)
														F_errmsg(m, int32(_a_F_make_op_5), v16+int32(32))
														mBase = m.M
														v190 = m.ExcPending
														if v190 != 0 {
															return int32(0)
														} else {
															F_parser_errposition(m, l0, l5)
															mBase = m.M
															v192 = m.ExcPending
															if v192 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_make_op_1), int32(810), int32(_a_F_make_op_2))
																mBase = m.M
																v197 = m.ExcPending
																if v197 != 0 {
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
											} else {
												F_make_fn_arguments(m, l0, v87, v16+int32(72), v16-int32(-64))
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return int32(0)
												} else {
													v113 = F_palloc0(m, int32(36))
													mBase = m.M
													v114 = m.ExcPending
													if v114 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v113))) = int32(17)
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
														v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+22)))
														v120 = *(*int32)(unsafe.Add(mBase, uint32(v117+v118)))
														*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = v120
														v122 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
														*(*int32)(unsafe.Add(mBase, uint32(v113)+12)) = v96
														*(*int32)(unsafe.Add(mBase, uint32(v113)+8)) = v122
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
														v126 = F_get_func_retset(m, v125)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v113)+32)) = l5
															*(*int32)(unsafe.Add(mBase, uint32(v113)+28)) = v87
															*(*uint8)(unsafe.Add(mBase, uint32(v113)+16)) = uint8(v126)
															if v126 != 0 {
																F_check_srf_call_placement(m, l0, l4, l5)
																mBase = m.M
																v132 = m.ExcPending
																if v132 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v113
																	F_ReleaseCatCache(m, v85)
																	mBase = m.M
																	v135 = m.ExcPending
																	if v135 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v16 + int32(80)
																		return v113
																	}
																}
															} else {
																F_ReleaseCatCache(m, v85)
																mBase = m.M
																v135 = m.ExcPending
																if v135 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v16 + int32(80)
																	return v113
																}
															}
														}
													}
												}
											}
										}
									} else {
										if v96 == int32(2281) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v180 = m.ExcPending
											if v180 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(1088))
												mBase = m.M
												v183 = m.ExcPending
												if v183 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = int32(_a_F_make_op_3)
													F_errmsg(m, int32(_a_F_make_op_5), v16+int32(32))
													mBase = m.M
													v190 = m.ExcPending
													if v190 != 0 {
														return int32(0)
													} else {
														F_parser_errposition(m, l0, l5)
														mBase = m.M
														v192 = m.ExcPending
														if v192 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_make_op_1), int32(810), int32(_a_F_make_op_2))
															mBase = m.M
															v197 = m.ExcPending
															if v197 != 0 {
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
										} else {
											F_make_fn_arguments(m, l0, v87, v16+int32(72), v16-int32(-64))
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return int32(0)
											} else {
												v113 = F_palloc0(m, int32(36))
												mBase = m.M
												v114 = m.ExcPending
												if v114 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v113))) = int32(17)
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
													v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+22)))
													v120 = *(*int32)(unsafe.Add(mBase, uint32(v117+v118)))
													*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = v120
													v122 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
													*(*int32)(unsafe.Add(mBase, uint32(v113)+12)) = v96
													*(*int32)(unsafe.Add(mBase, uint32(v113)+8)) = v122
													v125 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
													v126 = F_get_func_retset(m, v125)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v113)+32)) = l5
														*(*int32)(unsafe.Add(mBase, uint32(v113)+28)) = v87
														*(*uint8)(unsafe.Add(mBase, uint32(v113)+16)) = uint8(v126)
														if v126 != 0 {
															F_check_srf_call_placement(m, l0, l4, l5)
															mBase = m.M
															v132 = m.ExcPending
															if v132 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v113
																F_ReleaseCatCache(m, v85)
																mBase = m.M
																v135 = m.ExcPending
																if v135 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v16 + int32(80)
																	return v113
																}
															}
														} else {
															F_ReleaseCatCache(m, v85)
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return int32(0)
															} else {
																m.G0 = v16 + int32(80)
																return v113
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
			v48 = F_exprType(m, l2)
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				v50 = F_exprType(m, l3)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					v53 = F_oper(m, l0, l1, v48, v50, int32(0), l5)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+22)))
						v57 = v55 + v56
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+100))
						if v58 == int32(0) {
							v198 = v57
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v205 = m.ExcPending
							if v205 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(52461700))
								mBase = m.M
								v208 = m.ExcPending
								if v208 != 0 {
									return int32(0)
								} else {
									v209 = *(*int32)(unsafe.Add(mBase, uint32(v198)+80))
									v210 = *(*int32)(unsafe.Add(mBase, uint32(v198)+84))
									v211 = F_op_signature_string(m, l1, v209, v210)
									mBase = m.M
									v212 = m.ExcPending
									if v212 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v16))) = v211
										F_errmsg(m, int32(_a_F_make_op_0), v16)
										mBase = m.M
										v216 = m.ExcPending
										if v216 != 0 {
											return int32(0)
										} else {
											F_parser_errposition(m, l0, l5)
											mBase = m.M
											v218 = m.ExcPending
											if v218 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_make_op_1), int32(757), int32(_a_F_make_op_2))
												mBase = m.M
												v223 = m.ExcPending
												if v223 != 0 {
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
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = l3
							*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = l3
							v73 = F_list_make2_impl(m, v16+int32(48), v16+int32(44))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v50
								*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v48
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v57)+80))
								*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v77
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+84))
								*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v79
								v82 = v57 + int32(100)
								v83 = int32(2)
								v84 = v57
								v85 = v53
								v87 = v73
								v88 = v53 + int32(16)
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v84)+88))
								v96 = F_enforce_generic_type_consistency(m, v16+int32(72), v16-int32(-64), v83, v94, int32(0))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
									if v98 == int32(2281) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v159 = m.ExcPending
										if v159 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(1088))
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(_a_F_make_op_3)
												F_errmsg(m, int32(_a_F_make_op_4), v16+int32(16))
												mBase = m.M
												v169 = m.ExcPending
												if v169 != 0 {
													return int32(0)
												} else {
													F_parser_errposition(m, l0, l5)
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_make_op_1), int32(803), int32(_a_F_make_op_2))
														mBase = m.M
														v176 = m.ExcPending
														if v176 != 0 {
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
									} else {
										if l2 != 0 {
											v101 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
											if v101 == int32(2281) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v159 = m.ExcPending
												if v159 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v162 = m.ExcPending
													if v162 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(_a_F_make_op_3)
														F_errmsg(m, int32(_a_F_make_op_4), v16+int32(16))
														mBase = m.M
														v169 = m.ExcPending
														if v169 != 0 {
															return int32(0)
														} else {
															F_parser_errposition(m, l0, l5)
															mBase = m.M
															v171 = m.ExcPending
															if v171 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_make_op_1), int32(803), int32(_a_F_make_op_2))
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
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
											} else {
												if v96 == int32(2281) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v180 = m.ExcPending
													if v180 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v183 = m.ExcPending
														if v183 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = int32(_a_F_make_op_3)
															F_errmsg(m, int32(_a_F_make_op_5), v16+int32(32))
															mBase = m.M
															v190 = m.ExcPending
															if v190 != 0 {
																return int32(0)
															} else {
																F_parser_errposition(m, l0, l5)
																mBase = m.M
																v192 = m.ExcPending
																if v192 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_make_op_1), int32(810), int32(_a_F_make_op_2))
																	mBase = m.M
																	v197 = m.ExcPending
																	if v197 != 0 {
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
												} else {
													F_make_fn_arguments(m, l0, v87, v16+int32(72), v16-int32(-64))
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
														return int32(0)
													} else {
														v113 = F_palloc0(m, int32(36))
														mBase = m.M
														v114 = m.ExcPending
														if v114 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v113))) = int32(17)
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
															v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+22)))
															v120 = *(*int32)(unsafe.Add(mBase, uint32(v117+v118)))
															*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = v120
															v122 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
															*(*int32)(unsafe.Add(mBase, uint32(v113)+12)) = v96
															*(*int32)(unsafe.Add(mBase, uint32(v113)+8)) = v122
															v125 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
															v126 = F_get_func_retset(m, v125)
															mBase = m.M
															v127 = m.ExcPending
															if v127 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v113)+32)) = l5
																*(*int32)(unsafe.Add(mBase, uint32(v113)+28)) = v87
																*(*uint8)(unsafe.Add(mBase, uint32(v113)+16)) = uint8(v126)
																if v126 != 0 {
																	F_check_srf_call_placement(m, l0, l4, l5)
																	mBase = m.M
																	v132 = m.ExcPending
																	if v132 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v113
																		F_ReleaseCatCache(m, v85)
																		mBase = m.M
																		v135 = m.ExcPending
																		if v135 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v16 + int32(80)
																			return v113
																		}
																	}
																} else {
																	F_ReleaseCatCache(m, v85)
																	mBase = m.M
																	v135 = m.ExcPending
																	if v135 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v16 + int32(80)
																		return v113
																	}
																}
															}
														}
													}
												}
											}
										} else {
											if v96 == int32(2281) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = int32(_a_F_make_op_3)
														F_errmsg(m, int32(_a_F_make_op_5), v16+int32(32))
														mBase = m.M
														v190 = m.ExcPending
														if v190 != 0 {
															return int32(0)
														} else {
															F_parser_errposition(m, l0, l5)
															mBase = m.M
															v192 = m.ExcPending
															if v192 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_make_op_1), int32(810), int32(_a_F_make_op_2))
																mBase = m.M
																v197 = m.ExcPending
																if v197 != 0 {
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
											} else {
												F_make_fn_arguments(m, l0, v87, v16+int32(72), v16-int32(-64))
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return int32(0)
												} else {
													v113 = F_palloc0(m, int32(36))
													mBase = m.M
													v114 = m.ExcPending
													if v114 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v113))) = int32(17)
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
														v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+22)))
														v120 = *(*int32)(unsafe.Add(mBase, uint32(v117+v118)))
														*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = v120
														v122 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
														*(*int32)(unsafe.Add(mBase, uint32(v113)+12)) = v96
														*(*int32)(unsafe.Add(mBase, uint32(v113)+8)) = v122
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
														v126 = F_get_func_retset(m, v125)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v113)+32)) = l5
															*(*int32)(unsafe.Add(mBase, uint32(v113)+28)) = v87
															*(*uint8)(unsafe.Add(mBase, uint32(v113)+16)) = uint8(v126)
															if v126 != 0 {
																F_check_srf_call_placement(m, l0, l4, l5)
																mBase = m.M
																v132 = m.ExcPending
																if v132 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v113
																	F_ReleaseCatCache(m, v85)
																	mBase = m.M
																	v135 = m.ExcPending
																	if v135 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v16 + int32(80)
																		return v113
																	}
																}
															} else {
																F_ReleaseCatCache(m, v85)
																mBase = m.M
																v135 = m.ExcPending
																if v135 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v16 + int32(80)
																	return v113
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
		v143 = m.ExcPending
		if v143 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v146 = m.ExcPending
			if v146 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_make_op_6), int32(0))
				mBase = m.M
				v150 = m.ExcPending
				if v150 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_make_op_1), int32(729), int32(_a_F_make_op_2))
					mBase = m.M
					v155 = m.ExcPending
					if v155 != 0 {
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
func F_op_mergejoinable(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	if l0 != int32(2988) {
		if l0 != int32(1070) {
			v23 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(l0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23 == int32(0) {
					v34 = int32(0)
					return v34 & int32(1)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v29)+77)))
					F_ReleaseCatCache(m, v23)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = v31
						return v34 & int32(1)
					}
				}
			}
		} else {
			v8 = F_lookup_type_cache(m, l1, int32(8))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+64))
				v34 = base.B2i32(v12 == int32(382))
				return v34 & int32(1)
			}
		}
	} else {
		v16 = F_lookup_type_cache(m, l1, int32(8))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
			v34 = base.B2i32(v18 == int32(2987))
			return v34 & int32(1)
		}
	}
}
func F_op_volatile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v11 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(l0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
				F_errmsg_internal(m, int32(_a_F_op_volatile_0), v7)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_op_volatile_1), int32(1810), int32(_a_F_op_volatile_2))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v17+v18)+100))
			F_ReleaseCatCache(m, v11)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				if v20 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg_internal(m, int32(_a_F_op_volatile_0), v7)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_op_volatile_1), int32(1810), int32(_a_F_op_volatile_2))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v27 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(v20))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v20
								F_errmsg_internal(m, int32(_a_F_op_volatile_3), v7+int32(16))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_op_volatile_1), int32(2099), int32(_a_F_op_volatile_4))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
							v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
							v34 = int32(*(*int8)(unsafe.Add(mBase, uint32(v31+v32)+101)))
							F_ReleaseCatCache(m, v27)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(32)
								return v34
							}
						}
					}
				}
			}
		}
	}
}
