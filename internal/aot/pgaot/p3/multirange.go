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
							v37 = F_lookup_type_cache(m, v31, int32(_a_F_multirange_adjacent_range_0))
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
										F_errmsg_internal(m, int32(_a_F_multirange_adjacent_range_1), v9)
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_multirange_adjacent_range_2), int32(558), int32(_a_F_multirange_adjacent_range_3))
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
						v37 = F_lookup_type_cache(m, v31, int32(_a_F_multirange_adjacent_range_0))
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
									F_errmsg_internal(m, int32(_a_F_multirange_adjacent_range_1), v9)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_multirange_adjacent_range_2), int32(558), int32(_a_F_multirange_adjacent_range_3))
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
	v60 = F_lookup_type_cache(m, v47, int32(_a_F_multirange_agg_transfn_0))
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
	F_errmsg_internal(m, int32(_a_F_multirange_agg_transfn_1), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L20
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_multirange_agg_transfn_2), int32(1422), int32(_a_F_multirange_agg_transfn_3))
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
	F_errmsg_internal(m, int32(_a_F_multirange_agg_transfn_4), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L20
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_multirange_agg_transfn_2), int32(1426), int32(_a_F_multirange_agg_transfn_3))
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
	F_errmsg_internal(m, int32(_a_F_multirange_agg_transfn_5), v11)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L20
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_multirange_agg_transfn_2), int32(558), int32(_a_F_multirange_agg_transfn_6))
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
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
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = F_get_fn_expr_rettype(m, v11)
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
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v17 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L50
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L47
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L43
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L40
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L37
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v30 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v18 == v12 {
		v28 = v17
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v21 = F_lookup_type_cache(m, v12, int32(_a_F_multirange_constructor2_0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+296))
	if v23 == int32(0) {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v21
	v28 = v21
	goto L8
L15:
	;
	m.G0 = v9 + int32(32)
	return v113
L16:
	;
	v33 = int32(0)
	v35 = F_make_multirange(m, v12, v29, v33, v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v37 == int32(1) {
		goto L6
	} else {
		goto L20
	}
L19:
	;
	v113 = v35
	goto L15
L20:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v41 = F_pg_detoast_datum(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if int32(2) <= v43 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v46 != v47 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	if v43 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v105 = F_make_multirange(m, v12, v29, v101, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L36
	}
L25:
	;
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v51
	v101 = v51
	v104 = v51
	goto L24
L26:
	;
	goto L27
L27:
	;
	v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+8)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+10)))
	v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(v29)+11)))
	F_deconstruct_array(m, v41, v55, v56, v57, v9+int32(24), v9+int32(20), v9+int32(28))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v69 = F_palloc0(m, v66<<(uint(int32(2))%32))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if v71 <= int32(0) {
		v101 = v71
		v104 = v69
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v75 = int32(0)
	goto L31
L31:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v75))))
	if v83 == int32(1) {
		goto L3
	} else {
		goto L33
	}
L32:
	;
	v101 = v97
	v104 = v69
	goto L24
L33:
	;
	v87 = v75 << (uint(int32(2)) % 32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v89+v87)))
	v92 = F_pg_detoast_datum(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69+v87))) = v92
	v96 = v75 + int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if v96 < v97 {
		v75 = v96
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v113 = v105
	goto L15
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
	F_errmsg_internal(m, int32(_a_F_multirange_constructor2_1), v9)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_multirange_constructor2_2), int32(558), int32(_a_F_multirange_constructor2_3))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_multirange_constructor2_4), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_multirange_constructor2_2), int32(974), int32(_a_F_multirange_constructor2_5))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
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
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(_a_F_multirange_constructor2_6), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_multirange_constructor2_2), int32(982), int32(_a_F_multirange_constructor2_5))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v46
	F_errmsg_internal(m, int32(_a_F_multirange_constructor2_7), v9+int32(16))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_multirange_constructor2_2), int32(986), int32(_a_F_multirange_constructor2_5))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
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
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errmsg(m, int32(_a_F_multirange_constructor2_4), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_multirange_constructor2_2), int32(1008), int32(_a_F_multirange_constructor2_5))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
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
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_contains_multirange_0))
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
								F_errmsg_internal(m, int32(_a_F_multirange_contains_multirange_1), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_contains_multirange_2), int32(558), int32(_a_F_multirange_contains_multirange_3))
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
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_contains_multirange_0))
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
							F_errmsg_internal(m, int32(_a_F_multirange_contains_multirange_1), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_contains_multirange_2), int32(558), int32(_a_F_multirange_contains_multirange_3))
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
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_eq_0))
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
								F_errmsg_internal(m, int32(_a_F_multirange_eq_1), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_eq_2), int32(558), int32(_a_F_multirange_eq_3))
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
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_eq_0))
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
							F_errmsg_internal(m, int32(_a_F_multirange_eq_1), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_eq_2), int32(558), int32(_a_F_multirange_eq_3))
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v15 == v16 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v18 != v19 {
		v59 = v4
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L11
	} else {
		goto L19
	}
L4:
	;
	m.G0 = v13 + int32(32)
	return v59
L5:
	;
	if v18 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v59 = int32(1)
	goto L4
L7:
	;
	goto L8
L8:
	;
	v29 = v4
	goto L9
L9:
	;
	v35 = v13 + int32(24)
	v37 = v13 + int32(16)
	F_multirange_get_bounds(m, l0, l1, v29, v35, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v59 = v51
	goto L4
L11:
	;
	return int32(0)
L12:
	;
	v43 = v13 + int32(8)
	F_multirange_get_bounds(m, l0, l2, v29, v43, v13)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v46 = int32(0)
	v47 = F_range_cmp_bounds(m, l0, v35, v43)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	if v47 != 0 {
		v59 = v46
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v49 = F_range_cmp_bounds(m, l0, v37, v13)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	if v49 != 0 {
		v59 = v46
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v51 = int32(1)
	v53 = v29 + v51
	if v53 != v18 {
		v29 = v53
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	F_errmsg_internal(m, int32(_a_F_multirange_eq_internal_0), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_multirange_eq_internal_1), int32(1879), int32(_a_F_multirange_eq_internal_2))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L11
	} else {
		goto L21
	}
L21:
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
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v53 int32
	_ = v53
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
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
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
	var v176 int32
	_ = v176
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v18 = l1 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+11)))
	if l2 <= v6 {
		v53 = v6
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
	v28 = v6
	goto L3
L3:
	;
	v35 = int32(2)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v18+v23<<(uint(v35)%32))))
	v41 = v38&int32(2147483647) + v28
	if base.Ui32(v23) < base.Ui32(v35) {
		v53 = v41
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v53 = v41
	goto L1
L5:
	;
	if int32(0) <= v38 {
		v23 = v23 - int32(1)
		v28 = v41
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	v96 = v94 + l1 + v53
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
	v94 = v60 + v62 + int32(8)
	goto L7
L12:
	;
	v230 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v230)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v228
	v236 = int32(base.Ui32(v65)>>(uint(v230)%32)) & v230
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v236)
	v241 = int32(base.Ui32(v65)>>(uint(int32(3))%32)) & v230
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v241)
	v243 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+6)) = uint8(v243)
	v248 = int32(base.Ui32(v65)>>(uint(int32(2))%32)) & v230
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v248)
	v253 = int32(base.Ui32(v65)>>(uint(int32(4))%32)) & v230
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)) = uint8(v253)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v225
	m.G0 = v15 + int32(32)
	return
L13:
	;
	v225 = v221
	v228 = v223
	goto L12
L14:
	;
	if v68&int32(1) == int32(0) {
		goto L58
	} else {
		goto L59
	}
L15:
	;
	v194 = v170
	v195 = v171
	v196 = int32(-1)
	goto L14
L16:
	;
	if v20 == int32(105) {
		goto L53
	} else {
		goto L54
	}
L17:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v172 != 0 {
		goto L15
	} else {
		goto L52
	}
L18:
	;
	if v65&int32(80) != 0 {
		v225 = v97
		v228 = v96
		goto L12
	} else {
		goto L51
	}
L19:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v166 = int32(base.Ui32(v162) >> (uint(int32(2)) % 32))
	goto L18
L20:
	;
	if v65&int32(81) != 0 {
		v221 = v97
		v223 = v156
		goto L13
	} else {
		goto L49
	}
L21:
	;
	v156 = int32(0)
	v157 = v96
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
		v156 = v96
		v157 = v96 + v67
		goto L20
	} else {
		goto L35
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
	v156 = v109
	v157 = v96 + v67
	goto L20
L29:
	;
	v107 = int32(*(*int16)(unsafe.Add(mBase, uint32(v96))))
	v156 = v107
	v157 = v96 + v67
	goto L20
L30:
	;
	v105 = int32(*(*int8)(unsafe.Add(mBase, uint32(v96))))
	v156 = v105
	v157 = v96 + v67
	goto L20
L31:
	;
	return
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v67
	F_errmsg_internal(m, int32(_a_F_multirange_get_bounds_0), v15)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_multirange_get_bounds_1), int32(70), int32(_a_F_multirange_get_bounds_2))
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
	if v67 == int32(-1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v129 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v152 = F_strlen(m, v96)
	mBase = m.M
	v156 = v96
	v157 = v152 + v96 + int32(1)
	goto L20
L39:
	;
	v133 = int32(18)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)))
	if v135 == v133 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	if v129&int32(1) == int32(0) {
		goto L19
	} else {
		goto L48
	}
L42:
	;
	v138 = v133
	goto L44
L43:
	;
	v138 = int32(2)
	goto L44
L44:
	;
	if base.Ui32((v135-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v145 = int32(6)
	goto L47
L46:
	;
	v145 = v138
	goto L47
L47:
	;
	v166 = v145
	goto L18
L48:
	;
	v166 = int32(base.Ui32(v129) >> (uint(int32(1)) % 32))
	goto L18
L49:
	;
	if v67 == int32(-1) {
		v170 = v157
		v171 = v156
		goto L17
	} else {
		goto L50
	}
L50:
	;
	v174 = v157
	v175 = v156
	v176 = v67
	goto L16
L51:
	;
	v170 = v166 + v96
	v171 = v96
	goto L17
L52:
	;
	v174 = v170
	v175 = v171
	v176 = int32(-1)
	goto L16
L53:
	;
	v194 = (v174 + int32(3)) & int32(-4)
	v195 = v175
	v196 = v176
	goto L14
L54:
	;
	goto L55
L55:
	;
	switch v20 - int32(99) {
	case 0:
		v194 = v174
		v195 = v175
		v196 = v176
		goto L14
	case 1:
		goto L57
	default:
		goto L56
	}
L56:
	;
	v194 = (v174 + int32(1)) & int32(-2)
	v195 = v175
	v196 = v176
	goto L14
L57:
	;
	v194 = (v174 + int32(7)) & int32(-8)
	v195 = v175
	v196 = v176
	goto L14
L58:
	;
	v225 = v194
	v228 = v195
	goto L12
L59:
	;
	goto L60
L60:
	;
	switch v196 - int32(1) {
	case 0:
		goto L61
	case 1:
		goto L64
	default:
		goto L62
	case 3:
		goto L63
	}
L61:
	;
	v220 = int32(*(*int8)(unsafe.Add(mBase, uint32(v194))))
	v221 = v220
	v223 = v195
	goto L13
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L31
	} else {
		goto L65
	}
L63:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v221 = v204
	v223 = v195
	goto L13
L64:
	;
	v203 = int32(*(*int16)(unsafe.Add(mBase, uint32(v194))))
	v221 = v203
	v223 = v195
	goto L13
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v196
	F_errmsg_internal(m, int32(_a_F_multirange_get_bounds_0), v15+int32(16))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L31
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_multirange_get_bounds_1), int32(70), int32(_a_F_multirange_get_bounds_2))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L31
	} else {
		goto L67
	}
L67:
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
			v14 = F_lookup_type_cache(m, l1, int32(_a_F_multirange_get_typcache_0))
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
						F_errmsg_internal(m, int32(_a_F_multirange_get_typcache_1), v7)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_multirange_get_typcache_2), int32(558), int32(_a_F_multirange_get_typcache_3))
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
		v14 = F_lookup_type_cache(m, l1, int32(_a_F_multirange_get_typcache_0))
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
					F_errmsg_internal(m, int32(_a_F_multirange_get_typcache_1), v7)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_multirange_get_typcache_2), int32(558), int32(_a_F_multirange_get_typcache_3))
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
					v22 = F_lookup_type_cache(m, v16, int32(_a_F_multirange_lower_inc_0))
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
								F_errmsg_internal(m, int32(_a_F_multirange_lower_inc_1), v8)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_lower_inc_2), int32(558), int32(_a_F_multirange_lower_inc_3))
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
				v22 = F_lookup_type_cache(m, v16, int32(_a_F_multirange_lower_inc_0))
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
							F_errmsg_internal(m, int32(_a_F_multirange_lower_inc_1), v8)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_lower_inc_2), int32(558), int32(_a_F_multirange_lower_inc_3))
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
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
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
	var v149 int32
	_ = v149
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
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
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v532 int32
	_ = v532
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v554 int32
	_ = v554
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
		goto L131
	}
L5:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v46 == int32(0) {
		v554 = v25
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
	v38 = F_lookup_type_cache(m, v32, int32(_a_F_multirange_minus_0))
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
	return v554
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	if v49 == int32(0) {
		v554 = v25
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
	v92 = v52
	v98 = v49
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
	v92 = v88
	v98 = v89
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
	v149 = v2
	goto L25
L25:
	;
	v164 = m.G0
	v166 = v164 - int32(16)
	m.G0 = v166
	v171 = F_palloc0(m, (v98+v46)<<(uint(int32(2))%32))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
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
	v138 = F_multirange_get_range(m, v92, v30, v116)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v149 = v114
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
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v181 = v175
	v182 = int32(0)
	v185 = v2
	v192 = v2
	goto L35
L33:
	;
	v532 = v2
	goto L34
L34:
	;
	v542 = F_make_multirange(m, v32, v52, v532, v171)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L130
	}
L35:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v105+v192<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v166)+12)) = v198
	if v181 == int32(0) {
		v480 = v182
		v483 = v185
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v532 = v519
	goto L34
L37:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v166)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v171+v503<<(uint(int32(2))%32)))) = v516
	v518 = int32(1)
	v519 = v503 + v518
	v521 = v192 + v518
	if v521 != v46 {
		v181 = v499
		v182 = v500
		v185 = v519
		v192 = v521
		goto L35
	} else {
		goto L129
	}
L38:
	;
	v499 = int32(0)
	v500 = v480
	v503 = v483
	goto L37
L39:
	;
	v207 = v181
	v208 = v182
	goto L40
L40:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v166)+12))
	v222 = F_range_before_internal(m, v52, v207, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	v236 = v207
	v237 = v208
	v240 = v185
	goto L48
L42:
	;
	if v222 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v225 = v208 + int32(1)
	if v98 <= v225 {
		v480 = v225
		v483 = v185
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
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v149+v225<<(uint(int32(2))%32))))
	if v230 != 0 {
		v207 = v230
		v208 = v225
		goto L40
	} else {
		goto L47
	}
L47:
	;
	v480 = v225
	v483 = v185
	goto L38
L48:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v166)+12))
	v256 = m.G0
	v258 = v256 - int32(48)
	m.G0 = v258
	F_range_deserialize(m, v52, v250, v258+int32(40), v258+int32(24), v258+int32(15))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	v480 = v468
	v483 = v469
	goto L38
L50:
	;
	F_range_deserialize(m, v52, v236, v258+int32(32), v258+int32(16), v258+int32(14))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+36)))
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+44)))
	if v277 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	m.G0 = v258 + int32(48)
	if v432 != 0 {
		goto L116
	} else {
		goto L117
	}
L53:
	;
	v401 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v258)+38)) = uint8(v401)
	v403 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v258)+22)) = uint8(v403)
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+37)))
	v408 = v406 ^ v403
	*(*uint8)(unsafe.Add(mBase, uint32(v258)+37)) = uint8(v408)
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+21)))
	v412 = v410 ^ v403
	*(*uint8)(unsafe.Add(mBase, uint32(v258)+21)) = uint8(v412)
	v420 = F_make_range(m, v52, v258+int32(40), v258+int32(32), v401, v401)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L113
	}
L54:
	;
	v432 = int32(0)
	goto L52
L55:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+20)))
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+28)))
	if v336 == int32(1) {
		goto L85
	} else {
		goto L86
	}
L56:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+46)))
	if v276&int32(1) != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	if v276&int32(1) != 0 {
		goto L65
	} else {
		goto L66
	}
L59:
	;
	v283 = int32(0)
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+38)))
	if v284 == v280 {
		v432 = v283
		goto L52
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v280&int32(1) != 0 {
		goto L55
	} else {
		goto L64
	}
L62:
	;
	if v280&int32(1) != 0 {
		goto L55
	} else {
		goto L63
	}
L63:
	;
	v432 = v283
	goto L52
L64:
	;
	goto L54
L65:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+38)))
	if v292 == int32(0) {
		goto L55
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v52)+208))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v258)+40))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v258)+32))
	v300 = F_FunctionCall2Coll(m, v52+int32(212), v297, v298, v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	goto L54
L69:
	;
	if v300 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+37)))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+45)))
	if v305 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	if int32(0) <= v300 {
		goto L54
	} else {
		goto L84
	}
L73:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+46)))
	if v304&int32(1) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v324 = int32(0)
	if v304&int32(1) != 0 {
		v432 = v324
		goto L52
	} else {
		goto L82
	}
L76:
	;
	v313 = int32(0)
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+38)))
	if v308 == v314 {
		v432 = v313
		goto L52
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if v308&int32(1) == int32(0) {
		goto L55
	} else {
		goto L81
	}
L79:
	;
	if v308&int32(1) == int32(0) {
		goto L55
	} else {
		goto L80
	}
L80:
	;
	v432 = v313
	goto L52
L81:
	;
	goto L54
L82:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+38)))
	if v327&int32(1) != 0 {
		goto L55
	} else {
		goto L83
	}
L83:
	;
	v432 = v324
	goto L52
L84:
	;
	goto L55
L85:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+30)))
	if v335&int32(1) != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	if v335&int32(1) != 0 {
		goto L94
	} else {
		goto L95
	}
L88:
	;
	v342 = int32(0)
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+22)))
	if v343 == v339 {
		v432 = v342
		goto L52
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	if v339&int32(1) == int32(0) {
		goto L53
	} else {
		goto L93
	}
L91:
	;
	if v339&int32(1) == int32(0) {
		goto L53
	} else {
		goto L92
	}
L92:
	;
	v432 = v342
	goto L52
L93:
	;
	goto L54
L94:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+22)))
	if v355 != 0 {
		goto L53
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v52)+208))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v258)+24))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v258)+16))
	v361 = F_FunctionCall2Coll(m, v52+int32(212), v358, v359, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L98
	}
L97:
	;
	goto L54
L98:
	;
	if v361 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+21)))
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+29)))
	if v366 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	if int32(0) < v361 {
		goto L53
	} else {
		goto L112
	}
L102:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+30)))
	if v365&int32(1) == int32(0) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L104
L104:
	;
	v384 = int32(0)
	if v365&int32(1) != 0 {
		v432 = v384
		goto L52
	} else {
		goto L110
	}
L105:
	;
	v374 = int32(0)
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+22)))
	if base.B2i32(v369&int32(1) == v374)|base.B2i32(v369 == v379) != 0 {
		v432 = v374
		goto L52
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if v369&int32(1) != 0 {
		goto L53
	} else {
		goto L109
	}
L108:
	;
	goto L53
L109:
	;
	goto L54
L110:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+22)))
	if v387&int32(1) == int32(0) {
		goto L53
	} else {
		goto L111
	}
L111:
	;
	v432 = v384
	goto L52
L112:
	;
	goto L54
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171+v240<<(uint(int32(2))%32)))) = v420
	v427 = int32(0)
	v429 = F_make_range(m, v52, v258+int32(16), v258+int32(24), v427, v427)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166+int32(12)))) = v429
	v432 = v403
	goto L52
L115:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v149+v468<<(uint(int32(2))%32))))
	if v473 != 0 {
		v236 = v473
		v237 = v468
		v240 = v469
		goto L48
	} else {
		goto L128
	}
L116:
	;
	v438 = int32(1)
	v439 = v240 + v438
	v441 = v237 + v438
	if v441 < v98 {
		v468 = v441
		v469 = v439
		goto L115
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v166)+12))
	v444 = F_range_overlaps_internal(m, v52, v443, v236)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L120
	}
L119:
	;
	v480 = v441
	v483 = v439
	goto L38
L120:
	;
	if v444 == int32(0) {
		v499 = v236
		v500 = v237
		v503 = v240
		goto L37
	} else {
		goto L121
	}
L121:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v166)+12))
	v449 = F_range_minus_internal(m, v52, v448, v236)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166)+12)) = v449
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	v458 = int32(*(*int8)(unsafe.Add(mBase, uint32(v449+int32(base.Ui32(v452)>>(uint(int32(2))%32))-int32(1)))))
	goto L123
L123:
	;
	if v458&int32(1) != 0 {
		v499 = v236
		v500 = v237
		v503 = v240
		goto L37
	} else {
		goto L124
	}
L124:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v166)+12))
	v462 = F_range_before_internal(m, v52, v461, v236)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	if v462 != 0 {
		v499 = v236
		v500 = v237
		v503 = v240
		goto L37
	} else {
		goto L126
	}
L126:
	;
	v465 = v237 + int32(1)
	if v98 <= v465 {
		v480 = v465
		v483 = v240
		goto L38
	} else {
		goto L127
	}
L127:
	;
	v468 = v465
	v469 = v240
	goto L115
L128:
	;
	goto L49
L129:
	;
	goto L36
L130:
	;
	m.G0 = v166 + int32(16)
	v554 = v542
	goto L12
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v32
	F_errmsg_internal(m, int32(_a_F_multirange_minus_1), v22)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_multirange_minus_2), int32(558), int32(_a_F_multirange_minus_3))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
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
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
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
				v67 = v2
				m.G0 = v10 + int32(48)
				return v67
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
				if v23 == int32(0) {
					v67 = v2
					m.G0 = v10 + int32(48)
					return v67
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
							v48 = v10 + int32(32)
							F_multirange_get_bounds(m, v42, v13, v41-int32(1), v10+int32(40), v48)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v40)+296))
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
								v58 = v10 + int32(16)
								F_multirange_get_bounds(m, v51, v18, v52-int32(1), v10+int32(24), v58)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v40)+296))
									v62 = F_range_cmp_bounds(m, v61, v48, v58)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										v67 = base.B2i32(v62 <= int32(0))
										m.G0 = v10 + int32(48)
										return v67
									}
								}
							}
						} else {
							v32 = F_lookup_type_cache(m, v26, int32(_a_F_multirange_overleft_multirange_0))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
								if v34 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v26
										F_errmsg_internal(m, int32(_a_F_multirange_overleft_multirange_1), v10)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_multirange_overleft_multirange_2), int32(558), int32(_a_F_multirange_overleft_multirange_3))
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
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
									v48 = v10 + int32(32)
									F_multirange_get_bounds(m, v42, v13, v41-int32(1), v10+int32(40), v48)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int32(0)
									} else {
										v51 = *(*int32)(unsafe.Add(mBase, uint32(v40)+296))
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
										v58 = v10 + int32(16)
										F_multirange_get_bounds(m, v51, v18, v52-int32(1), v10+int32(24), v58)
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											v61 = *(*int32)(unsafe.Add(mBase, uint32(v40)+296))
											v62 = F_range_cmp_bounds(m, v61, v48, v58)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												v67 = base.B2i32(v62 <= int32(0))
												m.G0 = v10 + int32(48)
												return v67
											}
										}
									}
								}
							}
						}
					} else {
						v32 = F_lookup_type_cache(m, v26, int32(_a_F_multirange_overleft_multirange_0))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							if v34 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v26
									F_errmsg_internal(m, int32(_a_F_multirange_overleft_multirange_1), v10)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_multirange_overleft_multirange_2), int32(558), int32(_a_F_multirange_overleft_multirange_3))
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
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
								v48 = v10 + int32(32)
								F_multirange_get_bounds(m, v42, v13, v41-int32(1), v10+int32(40), v48)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v40)+296))
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
									v58 = v10 + int32(16)
									F_multirange_get_bounds(m, v51, v18, v52-int32(1), v10+int32(24), v58)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v40)+296))
										v62 = F_range_cmp_bounds(m, v61, v48, v58)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											v67 = base.B2i32(v62 <= int32(0))
											m.G0 = v10 + int32(48)
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
					v23 = F_lookup_type_cache(m, v17, int32(_a_F_multirange_upper_inf_0))
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
								F_errmsg_internal(m, int32(_a_F_multirange_upper_inf_1), v9)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_upper_inf_2), int32(558), int32(_a_F_multirange_upper_inf_3))
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
				v23 = F_lookup_type_cache(m, v17, int32(_a_F_multirange_upper_inf_0))
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
							F_errmsg_internal(m, int32(_a_F_multirange_upper_inf_1), v9)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_upper_inf_2), int32(558), int32(_a_F_multirange_upper_inf_3))
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
