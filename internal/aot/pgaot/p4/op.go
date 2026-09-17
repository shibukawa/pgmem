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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
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
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
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
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	v2 = int32(0)
	v14 = F_SearchSysCacheList(m, int32(3), int32(1), l0, v2, v2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v195
L2:
	;
	v103 = int32(0)
	v105 = F_SearchSysCache1(m, int32(40), l0)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L3
	} else {
		goto L30
	}
L3:
	;
	return int32(0)
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	if v18 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_ReleaseCatCacheList(m, v14)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v29 = v2
	v30 = v2
	goto L9
L8:
	;
	goto L2
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(48)+v29<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+56))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+22)))
	v40 = v38 + v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
	if v41 <= int32(2741) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	F_ReleaseCatCacheList(m, v14)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L3
	} else {
		goto L28
	}
L11:
	;
	v89 = v29 + int32(1)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	if v89 < v90 {
		v29 = v89
		v30 = v85
		goto L9
	} else {
		goto L27
	}
L12:
	;
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v68 = F_IndexAmTranslateStrategy(m, v66, v65, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L3
	} else {
		goto L23
	}
L13:
	;
	v57 = F_GetIndexAmRoutineByAmId(m, v41, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L3
	} else {
		goto L20
	}
L14:
	;
	switch v41 - int32(403) {
	case 0:
		v65 = v41
		goto L12
	case 1:
		goto L13
	case 2:
		v85 = v30
		goto L11
	default:
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if base.B2i32(v41 == int32(2742))|base.B2i32(v41 == int32(3580))|base.B2i32(v41 == int32(4000)) != 0 {
		v85 = v30
		goto L11
	} else {
		goto L19
	}
L17:
	;
	if v41 != int32(783) {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v85 = v30
	goto L11
L19:
	;
	goto L13
L20:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+10)))
	F_pfree(m, v57)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	if v59 != int32(1) {
		v85 = v30
		goto L11
	} else {
		goto L22
	}
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
	v65 = v64
	goto L12
L23:
	;
	if v68 == int32(0) {
		v85 = v30
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v73 = F_palloc(m, int32(16))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v75
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v80
	v82 = F_lappend(m, v30, v73)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v85 = v82
	goto L11
L27:
	;
	goto L10
L28:
	;
	if v85 != 0 {
		v195 = v85
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L2
L30:
	;
	if v105 == int32(0) {
		v195 = v103
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+22)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109+v110)+96))
	F_ReleaseCatCache(m, v105)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	if v112 == int32(0) {
		v195 = v103
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v119 = int32(0)
	v121 = F_SearchSysCacheList(m, int32(3), int32(1), v112, v119, v119)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v121)+40))
	if int32(0) < v123 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v133 = int32(0)
	v134 = v103
	goto L38
L36:
	;
	v184 = v103
	goto L37
L37:
	;
	F_ReleaseCatCacheList(m, v121)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L3
	} else {
		goto L48
	}
L38:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v121+int32(48)+v133<<(uint(int32(2))%32))))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+56))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+22)))
	v144 = v142 + v143
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+24))
	v147 = F_GetIndexAmRoutineByAmId(m, v145, int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L3
	} else {
		goto L41
	}
L39:
	;
	v184 = v173
	goto L37
L40:
	;
	v176 = v133 + int32(1)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v121)+40))
	if v176 < v177 {
		v133 = v176
		v134 = v173
		goto L38
	} else {
		goto L47
	}
L41:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+10)))
	if v149 != int32(1) {
		v173 = v134
		goto L40
	} else {
		goto L42
	}
L42:
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
		goto L43
	}
L43:
	;
	if v155 != int32(3) {
		v173 = v134
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v160 = F_palloc(m, int32(16))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v162
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+8)) = v166
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+12)) = v168
	v170 = F_lappend(m, v134, v160)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	v173 = v170
	goto L40
L47:
	;
	goto L39
L48:
	;
	v195 = v184
	goto L1
}
func F_make_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v52 int32
	_ = v52
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	if l3 != 0 {
		if l2 == int32(0) {
			v18 = F_exprType(m, l3)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v23 = F_left_oper(m, l0, l1, v18, int32(0), l5)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
					v27 = v25 + v26
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+100))
					if v28 == int32(0) {
						v133 = v27
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(52461700))
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return int32(0)
							} else {
								v144 = *(*int32)(unsafe.Add(mBase, uint32(v133)+80))
								v145 = *(*int32)(unsafe.Add(mBase, uint32(v133)+84))
								v146 = F_op_signature_string(m, l1, v144, v145)
								mBase = m.M
								v147 = m.ExcPending
								if v147 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14))) = v146
									F_errmsg(m, int32(_a_F_make_op_0), v14)
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
										return int32(0)
									} else {
										F_parser_errposition(m, l0, l5)
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_make_op_1), int32(706), int32(_a_F_make_op_2))
											mBase = m.M
											v158 = m.ExcPending
											if v158 != 0 {
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
						*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = l3
						v33 = int32(1)
						v37 = F_list_make1_impl(m, v33, v14+int32(8))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v18
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v27)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v40
							v72 = v37
							v73 = v33
							v74 = v27
							v76 = v23
							v79 = v14 + int32(40)
							v81 = v14 + int32(32)
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v74)+88))
							v84 = F_enforce_generic_type_consistency(m, v79, v81, v73, v82, int32(0))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								F_make_fn_arguments(m, l0, v72, v79, v81)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									v89 = F_palloc0(m, int32(36))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v89))) = int32(17)
										v93 = *(*int32)(unsafe.Add(mBase, uint32(v76)+16))
										v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+22)))
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v93+v94)))
										*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v96
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v74)+100))
										*(*int32)(unsafe.Add(mBase, uint32(v89)+12)) = v84
										*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = v98
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v74)+100))
										v102 = F_get_func_retset(m, v101)
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v89)+32)) = l5
											*(*int32)(unsafe.Add(mBase, uint32(v89)+28)) = v72
											*(*uint8)(unsafe.Add(mBase, uint32(v89)+16)) = uint8(v102)
											if v102 != 0 {
												F_check_srf_call_placement(m, l0, l4, l5)
												mBase = m.M
												v108 = m.ExcPending
												if v108 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v89
													F_ReleaseCatCache(m, v76)
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
														return int32(0)
													} else {
														m.G0 = v14 + int32(48)
														return v89
													}
												}
											} else {
												F_ReleaseCatCache(m, v76)
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return int32(0)
												} else {
													m.G0 = v14 + int32(48)
													return v89
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
			v42 = F_exprType(m, l2)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v44 = F_exprType(m, l3)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v47 = F_oper(m, l0, l1, v42, v44, int32(0), l5)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+22)))
						v51 = v49 + v50
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+100))
						if v52 == int32(0) {
							v133 = v51
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(52461700))
								mBase = m.M
								v143 = m.ExcPending
								if v143 != 0 {
									return int32(0)
								} else {
									v144 = *(*int32)(unsafe.Add(mBase, uint32(v133)+80))
									v145 = *(*int32)(unsafe.Add(mBase, uint32(v133)+84))
									v146 = F_op_signature_string(m, l1, v144, v145)
									mBase = m.M
									v147 = m.ExcPending
									if v147 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14))) = v146
										F_errmsg(m, int32(_a_F_make_op_0), v14)
										mBase = m.M
										v151 = m.ExcPending
										if v151 != 0 {
											return int32(0)
										} else {
											F_parser_errposition(m, l0, l5)
											mBase = m.M
											v153 = m.ExcPending
											if v153 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_make_op_1), int32(706), int32(_a_F_make_op_2))
												mBase = m.M
												v158 = m.ExcPending
												if v158 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l3
							*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l3
							v63 = F_list_make2_impl(m, v14+int32(16), v14+int32(12))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v44
								*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v42
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v51)+80))
								*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v67
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v51)+84))
								*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v69
								v72 = v63
								v73 = int32(2)
								v74 = v51
								v76 = v47
								v79 = v14 + int32(40)
								v81 = v14 + int32(32)
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v74)+88))
								v84 = F_enforce_generic_type_consistency(m, v79, v81, v73, v82, int32(0))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									F_make_fn_arguments(m, l0, v72, v79, v81)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										v89 = F_palloc0(m, int32(36))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v89))) = int32(17)
											v93 = *(*int32)(unsafe.Add(mBase, uint32(v76)+16))
											v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+22)))
											v96 = *(*int32)(unsafe.Add(mBase, uint32(v93+v94)))
											*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v96
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v74)+100))
											*(*int32)(unsafe.Add(mBase, uint32(v89)+12)) = v84
											*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = v98
											v101 = *(*int32)(unsafe.Add(mBase, uint32(v74)+100))
											v102 = F_get_func_retset(m, v101)
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v89)+32)) = l5
												*(*int32)(unsafe.Add(mBase, uint32(v89)+28)) = v72
												*(*uint8)(unsafe.Add(mBase, uint32(v89)+16)) = uint8(v102)
												if v102 != 0 {
													F_check_srf_call_placement(m, l0, l4, l5)
													mBase = m.M
													v108 = m.ExcPending
													if v108 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v89
														F_ReleaseCatCache(m, v76)
														mBase = m.M
														v111 = m.ExcPending
														if v111 != 0 {
															return int32(0)
														} else {
															m.G0 = v14 + int32(48)
															return v89
														}
													}
												} else {
													F_ReleaseCatCache(m, v76)
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
														return int32(0)
													} else {
														m.G0 = v14 + int32(48)
														return v89
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
		v119 = m.ExcPending
		if v119 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v122 = m.ExcPending
			if v122 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_make_op_3), int32(0))
				mBase = m.M
				v126 = m.ExcPending
				if v126 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_make_op_1), int32(678), int32(_a_F_make_op_2))
					mBase = m.M
					v131 = m.ExcPending
					if v131 != 0 {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	if l0 != int32(2988) {
		if l0 != int32(1070) {
			v22 = F_SearchSysCache1(m, int32(40), l0)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					v33 = int32(0)
					return v33 & int32(1)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v28)+77)))
					F_ReleaseCatCache(m, v22)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = v30
						return v33 & int32(1)
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
				v33 = base.B2i32(v12 == int32(382))
				return v33 & int32(1)
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
			v33 = base.B2i32(v18 == int32(2987))
			return v33 & int32(1)
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
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = F_SearchSysCache1(m, int32(40), l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
				F_errmsg_internal(m, int32(_a_F_op_volatile_0), v7)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_op_volatile_1), int32(1638), int32(_a_F_op_volatile_2))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v16+v17)+100))
			F_ReleaseCatCache(m, v10)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg_internal(m, int32(_a_F_op_volatile_0), v7)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_op_volatile_1), int32(1638), int32(_a_F_op_volatile_2))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v25 = F_SearchSysCache1(m, int32(47), v19)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if v25 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v19
								F_errmsg_internal(m, int32(_a_F_op_volatile_3), v7+int32(16))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_op_volatile_1), int32(1927), int32(_a_F_op_volatile_4))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
							v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
							v32 = int32(*(*int8)(unsafe.Add(mBase, uint32(v29+v30)+101)))
							F_ReleaseCatCache(m, v25)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(32)
								return v32
							}
						}
					}
				}
			}
		}
	}
}
