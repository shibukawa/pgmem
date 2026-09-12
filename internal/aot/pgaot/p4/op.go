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
	var v28 int32
	_ = v28
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
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
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
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
	return v194
L2:
	;
	v102 = int32(0)
	v104 = F_SearchSysCache1(m, int32(40), l0)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L32
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
	v28 = v2
	v30 = v2
	goto L9
L8:
	;
	goto L2
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(48)+v28<<(uint(int32(2))%32))))
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
	v92 = m.ExcPending
	if v92 != 0 {
		goto L3
	} else {
		goto L30
	}
L11:
	;
	v88 = v28 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	if v88 < v89 {
		v28 = v88
		v30 = v84
		goto L9
	} else {
		goto L29
	}
L12:
	;
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v67 = F_IndexAmTranslateStrategy(m, v65, v63, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L3
	} else {
		goto L25
	}
L13:
	;
	v55 = F_GetIndexAmRoutineByAmId(m, v41, int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L3
	} else {
		goto L22
	}
L14:
	;
	switch v41 - int32(403) {
	case 0:
		v63 = v41
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
	if v41 == int32(2742) {
		v84 = v30
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
	v84 = v30
	goto L11
L19:
	;
	if v41 == int32(3580) {
		v84 = v30
		goto L11
	} else {
		goto L20
	}
L20:
	;
	if v41 == int32(4000) {
		v84 = v30
		goto L11
	} else {
		goto L21
	}
L21:
	;
	goto L13
L22:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+10)))
	F_pfree(m, v55)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	if v57 != int32(1) {
		v84 = v30
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
	v63 = v62
	goto L12
L25:
	;
	if v67 == int32(0) {
		v84 = v30
		goto L11
	} else {
		goto L26
	}
L26:
	;
	v72 = F_palloc(m, int32(16))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v74
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = v79
	v81 = F_lappend(m, v30, v72)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	v84 = v81
	goto L11
L29:
	;
	goto L10
L30:
	;
	if v84 != 0 {
		v194 = v84
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L2
L32:
	;
	if v104 == int32(0) {
		v194 = v102
		goto L1
	} else {
		goto L33
	}
L33:
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
		goto L34
	}
L34:
	;
	if v111 == int32(0) {
		v194 = v102
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v118 = int32(0)
	v120 = F_SearchSysCacheList(m, int32(3), int32(1), v111, v118, v118)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v120)+40))
	if int32(0) < v122 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v131 = int32(0)
	v133 = v102
	goto L40
L38:
	;
	v183 = v102
	goto L39
L39:
	;
	F_ReleaseCatCacheList(m, v120)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L3
	} else {
		goto L50
	}
L40:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v120+int32(48)+v131<<(uint(int32(2))%32))))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+56))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+22)))
	v143 = v141 + v142
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+24))
	v146 = F_GetIndexAmRoutineByAmId(m, v144, int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L3
	} else {
		goto L43
	}
L41:
	;
	v183 = v172
	goto L39
L42:
	;
	v175 = v131 + int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v120)+40))
	if v175 < v176 {
		v131 = v175
		v133 = v172
		goto L40
	} else {
		goto L49
	}
L43:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+10)))
	if v148 != int32(1) {
		v172 = v133
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+16)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v143)+24))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v154 = F_IndexAmTranslateStrategy(m, v151, v152, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	if v154 != int32(3) {
		v172 = v133
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v159 = F_palloc(m, int32(16))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v161
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+8)) = v165
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+12)) = v167
	v169 = F_lappend(m, v133, v159)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	v172 = v169
	goto L42
L49:
	;
	goto L41
L50:
	;
	v194 = v183
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
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
						v136 = v27
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(52461700))
							mBase = m.M
							v146 = m.ExcPending
							if v146 != 0 {
								return int32(0)
							} else {
								v147 = *(*int32)(unsafe.Add(mBase, uint32(v136)+80))
								v148 = *(*int32)(unsafe.Add(mBase, uint32(v136)+84))
								v149 = F_op_signature_string(m, l1, v147, v148)
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14))) = v149
									F_errmsg(m, int32(201461), v14)
									mBase = m.M
									v154 = m.ExcPending
									if v154 != 0 {
										return int32(0)
									} else {
										F_parser_errposition(m, l0, l5)
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(492185), int32(706), int32(234114))
											mBase = m.M
											v161 = m.ExcPending
											if v161 != 0 {
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
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v74)+88))
							v84 = F_enforce_generic_type_consistency(m, v14+int32(40), v14+int32(32), v73, v82, int32(0))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								F_make_fn_arguments(m, l0, v72, v14+int32(40), v14+int32(32))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									v93 = F_palloc0(m, int32(36))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v93))) = int32(17)
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v76)+16))
										v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+22)))
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v97+v98)))
										*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v100
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v74)+100))
										*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = v84
										*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v102
										v105 = *(*int32)(unsafe.Add(mBase, uint32(v74)+100))
										v106 = F_get_func_retset(m, v105)
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v93)+32)) = l5
											*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v72
											*(*uint8)(unsafe.Add(mBase, uint32(v93)+16)) = uint8(v106)
											if v106 != 0 {
												F_check_srf_call_placement(m, l0, l4, l5)
												mBase = m.M
												v112 = m.ExcPending
												if v112 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v93
													F_ReleaseCatCache(m, v76)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int32(0)
													} else {
														m.G0 = v14 + int32(48)
														return v93
													}
												}
											} else {
												F_ReleaseCatCache(m, v76)
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int32(0)
												} else {
													m.G0 = v14 + int32(48)
													return v93
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
							v136 = v51
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(52461700))
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return int32(0)
								} else {
									v147 = *(*int32)(unsafe.Add(mBase, uint32(v136)+80))
									v148 = *(*int32)(unsafe.Add(mBase, uint32(v136)+84))
									v149 = F_op_signature_string(m, l1, v147, v148)
									mBase = m.M
									v150 = m.ExcPending
									if v150 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14))) = v149
										F_errmsg(m, int32(201461), v14)
										mBase = m.M
										v154 = m.ExcPending
										if v154 != 0 {
											return int32(0)
										} else {
											F_parser_errposition(m, l0, l5)
											mBase = m.M
											v156 = m.ExcPending
											if v156 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(492185), int32(706), int32(234114))
												mBase = m.M
												v161 = m.ExcPending
												if v161 != 0 {
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
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v74)+88))
								v84 = F_enforce_generic_type_consistency(m, v14+int32(40), v14+int32(32), v73, v82, int32(0))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									F_make_fn_arguments(m, l0, v72, v14+int32(40), v14+int32(32))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										v93 = F_palloc0(m, int32(36))
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v93))) = int32(17)
											v97 = *(*int32)(unsafe.Add(mBase, uint32(v76)+16))
											v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+22)))
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v97+v98)))
											*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v100
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v74)+100))
											*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = v84
											*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v102
											v105 = *(*int32)(unsafe.Add(mBase, uint32(v74)+100))
											v106 = F_get_func_retset(m, v105)
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v93)+32)) = l5
												*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v72
												*(*uint8)(unsafe.Add(mBase, uint32(v93)+16)) = uint8(v106)
												if v106 != 0 {
													F_check_srf_call_placement(m, l0, l4, l5)
													mBase = m.M
													v112 = m.ExcPending
													if v112 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v93
														F_ReleaseCatCache(m, v76)
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return int32(0)
														} else {
															m.G0 = v14 + int32(48)
															return v93
														}
													}
												} else {
													F_ReleaseCatCache(m, v76)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int32(0)
													} else {
														m.G0 = v14 + int32(48)
														return v93
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
		v123 = m.ExcPending
		if v123 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v126 = m.ExcPending
			if v126 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(440747), int32(0))
				mBase = m.M
				v130 = m.ExcPending
				if v130 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(492185), int32(678), int32(234114))
					mBase = m.M
					v135 = m.ExcPending
					if v135 != 0 {
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
				F_errmsg_internal(m, int32(68455), v7)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(496031), int32(1638), int32(383371))
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
						F_errmsg_internal(m, int32(68455), v7)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(496031), int32(1638), int32(383371))
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
								F_errmsg_internal(m, int32(44572), v7+int32(16))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(496031), int32(1927), int32(383383))
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
