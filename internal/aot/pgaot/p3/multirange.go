package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_multirange_adjacent_range(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v25 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17+int32(base.Ui32(v19)>>(uint(int32(2))%32))-int32(1)))))
			if v25&int32(1) != 0 {
				v48 = v2
				m.G0 = v9 + int32(16)
				return v48
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
				if v28 == int32(0) {
					v48 = v2
					m.G0 = v9 + int32(16)
					return v48
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
					if v33 != 0 {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						if v34 == v31 {
							v44 = v33
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
							v46 = F_range_adjacent_multirange_internal(m, v45, v17, v12)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v48 = v46
								m.G0 = v9 + int32(16)
								return v48
							}
						} else {
							v37 = F_lookup_type_cache(m, v31, int32(65536))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+296))
								if v39 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v31
										F_errmsg_internal(m, int32(363791), v9)
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(485286), int32(558), int32(391974))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v37
									v44 = v37
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
									v46 = F_range_adjacent_multirange_internal(m, v45, v17, v12)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										v48 = v46
										m.G0 = v9 + int32(16)
										return v48
									}
								}
							}
						}
					} else {
						v37 = F_lookup_type_cache(m, v31, int32(65536))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+296))
							if v39 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v31
									F_errmsg_internal(m, int32(363791), v9)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(485286), int32(558), int32(391974))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v37
								v44 = v37
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
								v46 = F_range_adjacent_multirange_internal(m, v45, v17, v12)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v48 = v46
									m.G0 = v9 + int32(16)
									return v48
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_multirange_agg_transfn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
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
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = v11 + int32(12)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 == v2 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L20
	} else {
		goto L60
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L20
	} else {
		goto L57
	}
L3:
	;
	if v44 != 0 {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v44 = v41
	goto L3
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
	v41 = v37
	goto L4
L6:
	;
	v33 = int32(0)
	if v14 == v33 {
		v41 = v33
		goto L4
	} else {
		goto L16
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	switch v19 - int32(429) {
	case 0:
		goto L9
	case 1:
		goto L8
	default:
		goto L6
	}
L8:
	;
	if v14 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	if v14 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v44 = int32(1)
	goto L3
L11:
	;
	goto L12
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+168))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v36 = v26
	v37 = int32(1)
	goto L5
L13:
	;
	v44 = int32(2)
	goto L3
L14:
	;
	goto L15
L15:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+368))
	v36 = v31
	v37 = int32(2)
	goto L5
L16:
	;
	v36 = v33
	v37 = v2
	goto L5
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = F_get_fn_expr_argtype(m, v45, int32(1))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L20
	} else {
		goto L54
	}
L20:
	;
	return int32(0)
L21:
	;
	v51 = F_type_is_multirange(m, v47)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if v51 == int32(0) {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	if v56 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+296))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v69 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L25:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v57 == v47 {
		v67 = v56
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v60 = F_lookup_type_cache(m, v47, int32(65536))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L20
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v60)+296))
	if v62 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v60
	v67 = v60
	goto L24
L31:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v79 != 0 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v75 = F_initArrayResult(m, v72, v73, int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L20
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v78 = v77
	goto L31
L35:
	;
	v78 = v75
	goto L31
L36:
	;
	m.G0 = v11 + int32(16)
	return v78
L37:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v81 = F_pg_detoast_datum(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L20
	} else {
		goto L38
	}
L38:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if int32(0) < v83 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v89 = F_palloc(m, v83<<(uint(int32(2))%32))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L20
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if v83 != 0 {
		goto L36
	} else {
		goto L51
	}
L42:
	;
	v92 = int32(0)
	goto L43
L43:
	;
	v103 = F_multirange_get_range(m, v68, v81, v92)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L20
	} else {
		goto L45
	}
L44:
	;
	v111 = int32(0)
	goto L47
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89+v92<<(uint(int32(2))%32)))) = v103
	v107 = v92 + int32(1)
	if v107 != v83 {
		v92 = v107
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v89+v111<<(uint(int32(2))%32))))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v124 = F_accumArrayResult(m, v78, v120, int32(0), v122, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L20
	} else {
		goto L49
	}
L48:
	;
	goto L36
L49:
	;
	v127 = v111 + int32(1)
	if v127 != v83 {
		v111 = v127
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v129 = F_make_empty_range(m, v68)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L20
	} else {
		goto L52
	}
L52:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v134 = F_accumArrayResult(m, v78, v129, int32(0), v132, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L20
	} else {
		goto L53
	}
L53:
	;
	goto L36
L54:
	;
	F_errmsg_internal(m, int32(60776), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L20
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(485286), int32(1422), int32(275972))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L20
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errmsg_internal(m, int32(394255), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L20
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(485286), int32(1426), int32(275972))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L20
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v47
	F_errmsg_internal(m, int32(363791), v11)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L20
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(485286), int32(558), int32(391974))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L20
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_multirange_constructor2(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
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
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = F_get_fn_expr_rettype(m, v14)
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
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v20 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L50
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L47
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L43
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L40
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L37
	}
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+296))
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v33 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v21 == v15 {
		v31 = v20
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v24 = F_lookup_type_cache(m, v15, int32(65536))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+296))
	if v26 == int32(0) {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v24
	v31 = v24
	goto L8
L15:
	;
	m.G0 = v12 + int32(32)
	return v125
L16:
	;
	v36 = int32(0)
	v38 = F_make_multirange(m, v15, v32, v36, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v40 == int32(1) {
		goto L6
	} else {
		goto L20
	}
L19:
	;
	v125 = v38
	goto L15
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v44 = F_pg_detoast_datum(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if int32(2) <= v46 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v49 != v50 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	if v46 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v114 = F_make_multirange(m, v15, v32, v107, v109)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L36
	}
L25:
	;
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v54
	v107 = v54
	v109 = v54
	goto L24
L26:
	;
	goto L27
L27:
	;
	v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32)+8)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+10)))
	v60 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32)+11)))
	F_deconstruct_array(m, v44, v58, v59, v60, v12+int32(24), v12+int32(20), v12+int32(28))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v72 = F_palloc0(m, v69<<(uint(int32(2))%32))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	if v74 <= int32(0) {
		v107 = v74
		v109 = v72
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v78 = int32(0)
	goto L31
L31:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+v78))))
	if v89 == int32(1) {
		goto L3
	} else {
		goto L33
	}
L32:
	;
	v107 = v103
	v109 = v72
	goto L24
L33:
	;
	v93 = v78 << (uint(int32(2)) % 32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v95+v93)))
	v98 = F_pg_detoast_datum(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72+v93))) = v98
	v102 = v78 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	if v102 < v103 {
		v78 = v102
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v125 = v114
	goto L15
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v15
	F_errmsg_internal(m, int32(363791), v12)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(485286), int32(558), int32(391974))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	F_errmsg_internal(m, int32(133404), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(485286), int32(974), int32(543339))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	F_errcode(m, int32(66))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(111243), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(485286), int32(982), int32(543339))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v49
	F_errmsg_internal(m, int32(361715), v12+int32(16))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(485286), int32(986), int32(543339))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
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
	F_errcode(m, int32(67108994))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errmsg(m, int32(133404), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(485286), int32(1008), int32(543339))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_multirange_contains_multirange(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
					v34 = F_multirange_contains_multirange_internal(m, v33, v12, v17)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v34
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(65536))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(363791), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(485286), int32(558), int32(391974))
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
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v34 = F_multirange_contains_multirange_internal(m, v33, v12, v17)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v34
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(65536))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(363791), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(485286), int32(558), int32(391974))
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
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						v34 = F_multirange_contains_multirange_internal(m, v33, v12, v17)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v34
						}
					}
				}
			}
		}
	}
}
func F_multirange_eq(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
					v34 = F_multirange_eq_internal(m, v33, v12, v17)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v34
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(65536))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(363791), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(485286), int32(558), int32(391974))
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
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v34 = F_multirange_eq_internal(m, v33, v12, v17)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v34
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(65536))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(363791), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(485286), int32(558), int32(391974))
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
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						v34 = F_multirange_eq_internal(m, v33, v12, v17)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v34
						}
					}
				}
			}
		}
	}
}
func F_multirange_eq_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v12 == v13 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v15 != v16 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L13
	} else {
		goto L21
	}
L4:
	;
	m.G0 = v10 + int32(32)
	return v59
L5:
	;
	v59 = v4
	goto L4
L6:
	;
	goto L7
L7:
	;
	if v15 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v59 = int32(1)
	goto L4
L9:
	;
	goto L10
L10:
	;
	v26 = v4
	goto L11
L11:
	;
	F_multirange_get_bounds(m, l0, l1, v26, v10+int32(24), v10+int32(16))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v59 = v51
	goto L4
L13:
	;
	return int32(0)
L14:
	;
	F_multirange_get_bounds(m, l0, l2, v26, v10+int32(8), v10)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v40 = int32(0)
	v45 = F_range_cmp_bounds(m, l0, v10+int32(24), v10+int32(8))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v45 != 0 {
		v59 = v40
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v49 = F_range_cmp_bounds(m, l0, v10+int32(16), v10)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	if v49 != 0 {
		v59 = v40
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v51 = int32(1)
	v53 = v26 + v51
	if v53 != v15 {
		v26 = v53
		goto L11
	} else {
		goto L20
	}
L20:
	;
	goto L12
L21:
	;
	F_errmsg_internal(m, int32(319852), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(485286), int32(1879), int32(305888))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_multirange_get_bounds(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
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
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v18 = l1 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+11)))
	if l2 <= v6 {
		v54 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v62 = v60 << (uint(int32(2)) % 32)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+(v18+v62)))))
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+8)))
	v67 = base.I32_extend16_s(v66)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+10)))
	if v20 == int32(105) {
		v94 = (v60*int32(5) + int32(11)) & int32(-4)
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v23 = l2
	v29 = v6
	goto L3
L3:
	;
	v35 = int32(2)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v18+v23<<(uint(v35)%32))))
	v41 = v38&int32(2147483647) + v29
	if base.Ui32(v23) < base.Ui32(v35) {
		v54 = v41
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v54 = v41
	goto L1
L5:
	;
	if int32(0) <= v38 {
		v23 = v23 - int32(1)
		v29 = v41
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	v96 = v94 + l1 + v54
	v97 = int32(0)
	if v65&int32(41) != 0 {
		goto L21
	} else {
		goto L22
	}
L8:
	;
	switch v20 - int32(99) {
	case 0:
		goto L11
	case 1:
		goto L10
	default:
		goto L9
	}
L9:
	;
	v94 = (v60*int32(5) + int32(9)) & int32(-2)
	goto L7
L10:
	;
	v94 = (v60*int32(5) + int32(15)) & int32(-8)
	goto L7
L11:
	;
	v94 = v62 + v60 + int32(8)
	goto L7
L12:
	;
	v289 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v289)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v287
	v295 = int32(base.Ui32(v65)>>(uint(v289)%32)) & v289
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v295)
	v300 = int32(base.Ui32(v65)>>(uint(int32(3))%32)) & v289
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v300)
	v302 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+6)) = uint8(v302)
	v307 = int32(base.Ui32(v65)>>(uint(int32(2))%32)) & v289
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v307)
	v312 = int32(base.Ui32(v65)>>(uint(int32(4))%32)) & v289
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)) = uint8(v312)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v288
	m.G0 = v15 + int32(32)
	return
L13:
	;
	v287 = v281
	v288 = v283
	goto L12
L14:
	;
	if v68&int32(1) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L15:
	;
	v253 = v228
	v254 = v229
	v255 = int32(-1)
	goto L14
L16:
	;
	if v20 == int32(105) {
		goto L75
	} else {
		goto L76
	}
L17:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	if v231 != 0 {
		goto L15
	} else {
		goto L74
	}
L18:
	;
	if v65&int32(80) != 0 {
		v287 = v96
		v288 = v97
		goto L12
	} else {
		goto L73
	}
L19:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v224 = int32(base.Ui32(v220) >> (uint(int32(2)) % 32))
	goto L18
L20:
	;
	if v65&int32(81) != 0 {
		v281 = v215
		v283 = v97
		goto L13
	} else {
		goto L71
	}
L21:
	;
	v214 = v96
	v215 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	if v68&int32(1) != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	switch v66 - int32(1) {
	case 0:
		goto L30
	case 1:
		goto L29
	default:
		goto L27
	case 3:
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	if int32(0) < v67 {
		goto L35
	} else {
		goto L36
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v214 = v67 + v96
	v215 = v109
	goto L20
L29:
	;
	v107 = int32(*(*int16)(unsafe.Add(mBase, uint32(v96))))
	v214 = v67 + v96
	v215 = v107
	goto L20
L30:
	;
	v105 = int32(*(*int8)(unsafe.Add(mBase, uint32(v96))))
	v214 = v67 + v96
	v215 = v105
	goto L20
L31:
	;
	return
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v67
	F_errmsg_internal(m, int32(474656), v15)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(320933), int32(70), int32(66797))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	v214 = v67 + v96
	v215 = v96
	goto L20
L36:
	;
	goto L37
L37:
	;
	if v67 == int32(-1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v129 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	if v96&int32(3) == int32(0) {
		v177 = v96
		goto L56
	} else {
		goto L57
	}
L41:
	;
	v132 = int32(6)
	v134 = int32(18)
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)))
	if v136 == v134 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	if v129&int32(1) == int32(0) {
		goto L19
	} else {
		goto L53
	}
L44:
	;
	v139 = v134
	goto L46
L45:
	;
	v139 = int32(2)
	goto L46
L46:
	;
	if v136&int32(254) == int32(2) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v144 = v132
	goto L49
L48:
	;
	v144 = v139
	goto L49
L49:
	;
	if v136 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v147 = v132
	goto L52
L51:
	;
	v147 = v144
	goto L52
L52:
	;
	v224 = v147
	goto L18
L53:
	;
	v224 = int32(base.Ui32(v129) >> (uint(int32(1)) % 32))
	goto L18
L54:
	;
	v214 = v210 + v96 + int32(1)
	v215 = v96
	goto L20
L55:
	;
	v210 = v202 - v96
	goto L54
L56:
	;
	v181 = v177
	goto L65
L57:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v161 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v210 = int32(0)
	goto L54
L59:
	;
	goto L60
L60:
	;
	v166 = v96
	goto L61
L61:
	;
	v170 = v166 + int32(1)
	if v170&int32(3) == int32(0) {
		v177 = v170
		goto L56
	} else {
		goto L63
	}
L62:
	;
	v202 = v170
	goto L55
L63:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v175 != 0 {
		v166 = v170
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v190 = int32(-2139062144)
	if (int32(16843008)-v187|v187)&v190 == v190 {
		v181 = v181 + int32(4)
		goto L65
	} else {
		goto L67
	}
L66:
	;
	v196 = v181
	goto L68
L67:
	;
	goto L66
L68:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	if v200 != 0 {
		v196 = v196 + int32(1)
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v202 = v196
	goto L55
L70:
	;
	goto L69
L71:
	;
	if v67 == int32(-1) {
		v228 = v214
		v229 = v215
		goto L17
	} else {
		goto L72
	}
L72:
	;
	v233 = v214
	v234 = v215
	v235 = v67
	goto L16
L73:
	;
	v228 = v224 + v96
	v229 = v96
	goto L17
L74:
	;
	v233 = v228
	v234 = v229
	v235 = int32(-1)
	goto L16
L75:
	;
	v253 = (v233 + int32(3)) & int32(-4)
	v254 = v234
	v255 = v235
	goto L14
L76:
	;
	goto L77
L77:
	;
	switch v20 - int32(99) {
	case 0:
		v253 = v233
		v254 = v234
		v255 = v235
		goto L14
	case 1:
		goto L79
	default:
		goto L78
	}
L78:
	;
	v253 = (v233 + int32(1)) & int32(-2)
	v254 = v234
	v255 = v235
	goto L14
L79:
	;
	v253 = (v233 + int32(7)) & int32(-8)
	v254 = v234
	v255 = v235
	goto L14
L80:
	;
	v287 = v254
	v288 = v253
	goto L12
L81:
	;
	goto L82
L82:
	;
	switch v255 - int32(1) {
	case 0:
		goto L83
	case 1:
		goto L86
	default:
		goto L84
	case 3:
		goto L85
	}
L83:
	;
	v279 = int32(*(*int8)(unsafe.Add(mBase, uint32(v253))))
	v281 = v254
	v283 = v279
	goto L13
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L31
	} else {
		goto L87
	}
L85:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	v281 = v254
	v283 = v263
	goto L13
L86:
	;
	v262 = int32(*(*int16)(unsafe.Add(mBase, uint32(v253))))
	v281 = v254
	v283 = v262
	goto L13
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v255
	F_errmsg_internal(m, int32(474656), v15+int32(16))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L31
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(320933), int32(70), int32(66797))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L31
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_multirange_get_typcache(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v11 int32
	_ = v11
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
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if v10 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		if v11 == l1 {
			v23 = v10
			m.G0 = v7 + int32(16)
			return v23
		} else {
			v14 = F_lookup_type_cache(m, l1, int32(65536))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+296))
				if v18 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
						F_errmsg_internal(m, int32(363791), v7)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(485286), int32(558), int32(391974))
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
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v14
					v23 = v14
					m.G0 = v7 + int32(16)
					return v23
				}
			}
		}
	} else {
		v14 = F_lookup_type_cache(m, l1, int32(65536))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+296))
			if v18 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
					F_errmsg_internal(m, int32(363791), v7)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(485286), int32(558), int32(391974))
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
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v14
				v23 = v14
				m.G0 = v7 + int32(16)
				return v23
			}
		}
	}
}
func F_multirange_lower_inc(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
		if v15 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
			if v18 != 0 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				if v19 == v16 {
					v29 = v18
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+296))
					F_multirange_get_bounds(m, v30, v11, int32(0), v8+int32(24), v8+int32(16))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+29)))
						v42 = v38
						m.G0 = v8 + int32(32)
						return v42
					}
				} else {
					v22 = F_lookup_type_cache(m, v16, int32(65536))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+296))
						if v24 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
								F_errmsg_internal(m, int32(363791), v8)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(485286), int32(558), int32(391974))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v22
							v29 = v22
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+296))
							F_multirange_get_bounds(m, v30, v11, int32(0), v8+int32(24), v8+int32(16))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+29)))
								v42 = v38
								m.G0 = v8 + int32(32)
								return v42
							}
						}
					}
				}
			} else {
				v22 = F_lookup_type_cache(m, v16, int32(65536))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+296))
					if v24 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
							F_errmsg_internal(m, int32(363791), v8)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(485286), int32(558), int32(391974))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v22
						v29 = v22
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+296))
						F_multirange_get_bounds(m, v30, v11, int32(0), v8+int32(24), v8+int32(16))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+29)))
							v42 = v38
							m.G0 = v8 + int32(32)
							return v42
						}
					}
				}
			}
		} else {
			v42 = int32(0)
			m.G0 = v8 + int32(32)
			return v42
		}
	}
}
func F_multirange_lt(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_multirange_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v2) >> (uint(int32(31)) % 32))
	}
}
func F_multirange_minus(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v530 int32
	_ = v530
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v556 int32
	_ = v556
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	v2 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v30 = F_pg_detoast_datum(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v34 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L132
	}
L5:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v46 == int32(0) {
		v556 = v25
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v35 == v32 {
		v45 = v34
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v38 = F_lookup_type_cache(m, v32, int32(65536))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
	if v40 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v38
	v45 = v38
	goto L5
L12:
	;
	m.G0 = v22 + int32(16)
	return v556
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	if v49 == int32(0) {
		v556 = v25
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
	if int32(0) < v46 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v58 = F_palloc(m, v46<<(uint(int32(2))%32))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v98 = v49
	v103 = v52
	v105 = v2
	goto L17
L17:
	;
	if int32(0) < v98 {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v60 = int32(0)
	goto L19
L19:
	;
	v82 = F_multirange_get_range(m, v52, v25, v60)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v98 = v89
	v103 = v88
	v105 = v58
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58+v60<<(uint(int32(2))%32)))) = v82
	v86 = v60 + int32(1)
	if v86 != v46 {
		v60 = v86
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v114 = F_palloc(m, v98<<(uint(int32(2))%32))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v150 = v2
	goto L25
L25:
	;
	v163 = int32(0)
	v165 = m.G0
	v167 = v165 - int32(16)
	m.G0 = v167
	v172 = F_palloc0(m, (v98+v46)<<(uint(int32(2))%32))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L31
	}
L26:
	;
	v116 = int32(0)
	goto L27
L27:
	;
	v138 = F_multirange_get_range(m, v103, v30, v116)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v150 = v114
	goto L25
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114+v116<<(uint(int32(2))%32)))) = v138
	v142 = v116 + int32(1)
	if v142 != v98 {
		v116 = v142
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	if int32(0) < v46 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v182 = v163
	v183 = v176
	v184 = v163
	v193 = v2
	goto L35
L33:
	;
	v530 = v163
	goto L34
L34:
	;
	v542 = F_make_multirange(m, v32, v52, v530, v172)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L131
	}
L35:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v105+v193<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v167)+12)) = v199
	if v183 == int32(0) {
		v479 = v182
		v481 = v184
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v530 = v519
	goto L34
L37:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v172+v501<<(uint(int32(2))%32)))) = v516
	v518 = int32(1)
	v519 = v501 + v518
	v521 = v193 + v518
	if v521 != v46 {
		v182 = v499
		v183 = v500
		v184 = v519
		v193 = v521
		goto L35
	} else {
		goto L130
	}
L38:
	;
	v499 = v479
	v500 = int32(0)
	v501 = v481
	goto L37
L39:
	;
	v208 = v182
	v209 = v183
	goto L40
L40:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	v223 = F_range_before_internal(m, v52, v209, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	v237 = v208
	v238 = v209
	v239 = v184
	goto L48
L42:
	;
	if v223 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v226 = v208 + int32(1)
	if v98 <= v226 {
		v479 = v226
		v481 = v184
		goto L38
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	goto L41
L46:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v150+v226<<(uint(int32(2))%32))))
	if v231 != 0 {
		v208 = v226
		v209 = v231
		goto L40
	} else {
		goto L47
	}
L47:
	;
	v479 = v226
	v481 = v184
	goto L38
L48:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	v257 = m.G0
	v259 = v257 - int32(48)
	m.G0 = v259
	F_range_deserialize(m, v52, v251, v259+int32(40), v259+int32(24), v259+int32(15))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	v479 = v468
	v481 = v469
	goto L38
L50:
	;
	F_range_deserialize(m, v52, v238, v259+int32(32), v259+int32(16), v259+int32(14))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+36)))
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+44)))
	if v278 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	m.G0 = v259 + int32(48)
	if v432 != 0 {
		goto L117
	} else {
		goto L118
	}
L53:
	;
	v401 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v259)+38)) = uint8(v401)
	v403 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v259)+22)) = uint8(v403)
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+37)))
	v408 = v406 ^ v403
	*(*uint8)(unsafe.Add(mBase, uint32(v259)+37)) = uint8(v408)
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+21)))
	v412 = v410 ^ v403
	*(*uint8)(unsafe.Add(mBase, uint32(v259)+21)) = uint8(v412)
	v420 = F_make_range(m, v52, v259+int32(40), v259+int32(32), v401, v401)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L114
	}
L54:
	;
	v432 = int32(0)
	goto L52
L55:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+20)))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+28)))
	if v337 == int32(1) {
		goto L85
	} else {
		goto L86
	}
L56:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+46)))
	if v277&int32(1) != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	if v277&int32(1) != 0 {
		goto L65
	} else {
		goto L66
	}
L59:
	;
	v284 = int32(0)
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+38)))
	if v285 == v281&int32(255) {
		v432 = v284
		goto L52
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v281&int32(1) != 0 {
		goto L55
	} else {
		goto L64
	}
L62:
	;
	if v281&int32(1) != 0 {
		goto L55
	} else {
		goto L63
	}
L63:
	;
	v432 = v284
	goto L52
L64:
	;
	goto L54
L65:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+38)))
	if v295 == int32(0) {
		goto L55
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v52)+208))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v259)+40))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v259)+32))
	v303 = F_FunctionCall2Coll(m, v52+int32(212), v300, v301, v302)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	goto L54
L69:
	;
	if v303 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+37)))
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+45)))
	if v308 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	if int32(0) <= v303 {
		goto L54
	} else {
		goto L84
	}
L73:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+46)))
	if v307&int32(1) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v327 = int32(0)
	if v307&int32(1) != 0 {
		v432 = v327
		goto L52
	} else {
		goto L82
	}
L76:
	;
	v316 = int32(0)
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+38)))
	if v311 == v317 {
		v432 = v316
		goto L52
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if v311&int32(1) == int32(0) {
		goto L55
	} else {
		goto L81
	}
L79:
	;
	if v311&int32(1) == int32(0) {
		goto L55
	} else {
		goto L80
	}
L80:
	;
	v432 = v316
	goto L52
L81:
	;
	goto L54
L82:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+38)))
	if v330 != 0 {
		goto L55
	} else {
		goto L83
	}
L83:
	;
	v432 = v327
	goto L52
L84:
	;
	goto L55
L85:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+30)))
	if v336&int32(1) != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	if v336&int32(1) != 0 {
		goto L94
	} else {
		goto L95
	}
L88:
	;
	v343 = int32(0)
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+22)))
	if v344 == v340&int32(255) {
		v432 = v343
		goto L52
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	if v340&int32(1) == int32(0) {
		goto L53
	} else {
		goto L93
	}
L91:
	;
	if v340&int32(1) == int32(0) {
		goto L53
	} else {
		goto L92
	}
L92:
	;
	v432 = v343
	goto L52
L93:
	;
	goto L54
L94:
	;
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+22)))
	if v358 != 0 {
		goto L53
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v52)+208))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v259)+24))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v259)+16))
	v364 = F_FunctionCall2Coll(m, v52+int32(212), v361, v362, v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L98
	}
L97:
	;
	goto L54
L98:
	;
	if v364 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+21)))
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+29)))
	if v369 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	if int32(0) < v364 {
		goto L53
	} else {
		goto L113
	}
L102:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+30)))
	if v368&int32(1) == int32(0) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L104
L104:
	;
	v386 = int32(0)
	if v368&int32(1) != 0 {
		v432 = v386
		goto L52
	} else {
		goto L111
	}
L105:
	;
	v377 = int32(0)
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+22)))
	if v372 == v378 {
		v432 = v377
		goto L52
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if v372&int32(1) != 0 {
		goto L53
	} else {
		goto L110
	}
L108:
	;
	if v372&int32(1) == int32(0) {
		v432 = v377
		goto L52
	} else {
		goto L109
	}
L109:
	;
	goto L53
L110:
	;
	goto L54
L111:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+22)))
	if v389 == int32(0) {
		goto L53
	} else {
		goto L112
	}
L112:
	;
	v432 = v386
	goto L52
L113:
	;
	goto L54
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172+v239<<(uint(int32(2))%32)))) = v420
	v427 = int32(0)
	v429 = F_make_range(m, v52, v259+int32(16), v259+int32(24), v427, v427)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167+int32(12)))) = v429
	v432 = v403
	goto L52
L116:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v150+v468<<(uint(int32(2))%32))))
	if v473 != 0 {
		v237 = v468
		v238 = v473
		v239 = v469
		goto L48
	} else {
		goto L129
	}
L117:
	;
	v438 = int32(1)
	v439 = v239 + v438
	v441 = v237 + v438
	if v441 < v98 {
		v468 = v441
		v469 = v439
		goto L116
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	v444 = F_range_overlaps_internal(m, v52, v443, v238)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L121
	}
L120:
	;
	v479 = v441
	v481 = v439
	goto L38
L121:
	;
	if v444 == int32(0) {
		v499 = v237
		v500 = v238
		v501 = v239
		goto L37
	} else {
		goto L122
	}
L122:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	v449 = F_range_minus_internal(m, v52, v448, v238)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+12)) = v449
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	v458 = int32(*(*int8)(unsafe.Add(mBase, uint32(v449+int32(base.Ui32(v452)>>(uint(int32(2))%32))-int32(1)))))
	goto L124
L124:
	;
	if v458&int32(1) != 0 {
		v499 = v237
		v500 = v238
		v501 = v239
		goto L37
	} else {
		goto L125
	}
L125:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	v462 = F_range_before_internal(m, v52, v461, v238)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	if v462 != 0 {
		v499 = v237
		v500 = v238
		v501 = v239
		goto L37
	} else {
		goto L127
	}
L127:
	;
	v465 = v237 + int32(1)
	if v98 <= v465 {
		v479 = v465
		v481 = v239
		goto L38
	} else {
		goto L128
	}
L128:
	;
	v468 = v465
	v469 = v239
	goto L116
L129:
	;
	goto L49
L130:
	;
	goto L36
L131:
	;
	m.G0 = v167 + int32(16)
	v556 = v542
	goto L12
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v32
	F_errmsg_internal(m, int32(363791), v22)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(485286), int32(558), int32(391974))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_multirange_overleft_multirange(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			if v20 == int32(0) {
				v70 = v2
				m.G0 = v10 + int32(48)
				return v70
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
				if v23 == int32(0) {
					v70 = v2
					m.G0 = v10 + int32(48)
					return v70
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
					if v28 != 0 {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
						if v29 == v26 {
							v40 = v28
							v41 = v20
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+296))
							F_multirange_get_bounds(m, v42, v13, v41-int32(1), v10+int32(40), v10+int32(32))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v40)+296))
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
								F_multirange_get_bounds(m, v51, v18, v52-int32(1), v10+int32(24), v10+int32(16))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v40)+296))
									v66 = F_range_cmp_bounds(m, v61, v10+int32(32), v10+int32(16))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										v70 = base.B2i32(v66 <= int32(0))
										m.G0 = v10 + int32(48)
										return v70
									}
								}
							}
						} else {
							v32 = F_lookup_type_cache(m, v26, int32(65536))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
								if v34 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v26
										F_errmsg_internal(m, int32(363791), v10)
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(485286), int32(558), int32(391974))
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v32
									v39 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
									v40 = v32
									v41 = v39
									v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+296))
									F_multirange_get_bounds(m, v42, v13, v41-int32(1), v10+int32(40), v10+int32(32))
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int32(0)
									} else {
										v51 = *(*int32)(unsafe.Add(mBase, uint32(v40)+296))
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
										F_multirange_get_bounds(m, v51, v18, v52-int32(1), v10+int32(24), v10+int32(16))
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											v61 = *(*int32)(unsafe.Add(mBase, uint32(v40)+296))
											v66 = F_range_cmp_bounds(m, v61, v10+int32(32), v10+int32(16))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int32(0)
											} else {
												v70 = base.B2i32(v66 <= int32(0))
												m.G0 = v10 + int32(48)
												return v70
											}
										}
									}
								}
							}
						}
					} else {
						v32 = F_lookup_type_cache(m, v26, int32(65536))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							if v34 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v26
									F_errmsg_internal(m, int32(363791), v10)
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(485286), int32(558), int32(391974))
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v32
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
								v40 = v32
								v41 = v39
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+296))
								F_multirange_get_bounds(m, v42, v13, v41-int32(1), v10+int32(40), v10+int32(32))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v40)+296))
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
									F_multirange_get_bounds(m, v51, v18, v52-int32(1), v10+int32(24), v10+int32(16))
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v40)+296))
										v66 = F_range_cmp_bounds(m, v61, v10+int32(32), v10+int32(16))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											v70 = base.B2i32(v66 <= int32(0))
											m.G0 = v10 + int32(48)
											return v70
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
func F_multirange_recv(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v36 int32
	_ = v36
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
	var v71 int32
	_ = v71
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = F_get_multirange_io_data(m, l0, v18, int32(2))
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
	v25 = F_pq_getmsgint(m, v17, int32(4))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = F_palloc(m, v25<<(uint(int32(2))%32))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_initStringInfo(m, v14)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v25 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v36 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	F_pfree(m, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L18
	}
L9:
	;
	v48 = F_pq_getmsgint(m, v17, int32(4))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v50 = F_pq_getmsgbytes(m, v17, v48)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v53 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v52))) = uint8(v53)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v53
	goto L13
L13:
	;
	F_appendBinaryStringInfo(m, v14, v50, v48)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	v65 = F_ReceiveFunctionCall(m, v20+int32(4), v14, v64, v16)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v67 = F_pg_detoast_datum(m, v65)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29+v36<<(uint(int32(2))%32)))) = v67
	v71 = v36 + int32(1)
	if v71 != v25 {
		v36 = v71
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L10
L18:
	;
	F_pq_getmsgend(m, v17)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+296))
	v91 = F_make_multirange(m, v18, v90, v25, v29)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	m.G0 = v14 + int32(16)
	return v91
}
func F_multirange_upper_inf(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
		if v16 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
			if v19 != 0 {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				if v20 == v17 {
					v31 = v19
					v32 = v16
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+296))
					F_multirange_get_bounds(m, v33, v12, v32-int32(1), v9+int32(24), v9+int32(16))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
						v47 = v42
						m.G0 = v9 + int32(32)
						return v47
					}
				} else {
					v23 = F_lookup_type_cache(m, v17, int32(65536))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+296))
						if v25 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
								F_errmsg_internal(m, int32(363791), v9)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(485286), int32(558), int32(391974))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							v31 = v23
							v32 = v30
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+296))
							F_multirange_get_bounds(m, v33, v12, v32-int32(1), v9+int32(24), v9+int32(16))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
								v47 = v42
								m.G0 = v9 + int32(32)
								return v47
							}
						}
					}
				}
			} else {
				v23 = F_lookup_type_cache(m, v17, int32(65536))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+296))
					if v25 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
							F_errmsg_internal(m, int32(363791), v9)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(485286), int32(558), int32(391974))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
						v31 = v23
						v32 = v30
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+296))
						F_multirange_get_bounds(m, v33, v12, v32-int32(1), v9+int32(24), v9+int32(16))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
							v47 = v42
							m.G0 = v9 + int32(32)
							return v47
						}
					}
				}
			}
		} else {
			v47 = int32(0)
			m.G0 = v9 + int32(32)
			return v47
		}
	}
}
