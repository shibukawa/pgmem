package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_multirange_adjacent_range(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int64
	_ = v52
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v7 = int64(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			v26 = int32(*(*int8)(unsafe.Add(mBase, uint32(v18+int32(base.Ui32(v20)>>(uint(int32(2))%32))-int32(1)))))
			if v26&int32(1) != 0 {
				v52 = v7
				m.G0 = v10 + int32(16)
				return v52
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
				if v29 == int32(0) {
					v52 = v7
					m.G0 = v10 + int32(16)
					return v52
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
					if v34 != 0 {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
						if v35 == v32 {
							v45 = v34
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
							v47 = F_range_adjacent_multirange_internal(m, v46, v18, v13)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int64(0)
							} else {
								v52 = base.I64_extend_i32_u(v47)
								m.G0 = v10 + int32(16)
								return v52
							}
						} else {
							v38 = F_lookup_type_cache(m, v32, int32(_a_F_multirange_adjacent_range_0))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int64(0)
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
								if v40 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v32
										F_errmsg_internal(m, int32(_a_F_multirange_adjacent_range_1), v10)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_multirange_adjacent_range_2), int32(561), int32(_a_F_multirange_adjacent_range_3))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v38
									v45 = v38
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
									v47 = F_range_adjacent_multirange_internal(m, v46, v18, v13)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int64(0)
									} else {
										v52 = base.I64_extend_i32_u(v47)
										m.G0 = v10 + int32(16)
										return v52
									}
								}
							}
						}
					} else {
						v38 = F_lookup_type_cache(m, v32, int32(_a_F_multirange_adjacent_range_0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
							if v40 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v32
									F_errmsg_internal(m, int32(_a_F_multirange_adjacent_range_1), v10)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_multirange_adjacent_range_2), int32(561), int32(_a_F_multirange_adjacent_range_3))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v38
								v45 = v38
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
								v47 = F_range_adjacent_multirange_internal(m, v46, v18, v13)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int64(0)
								} else {
									v52 = base.I64_extend_i32_u(v47)
									m.G0 = v10 + int32(16)
									return v52
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_multirange_agg_transfn(m *base.Module, l0 int32) int64 {
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v119 int64
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
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
	v178 = m.ExcPending
	if v178 != 0 {
		goto L20
	} else {
		goto L60
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
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
	switch v19 - int32(435) {
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
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+376))
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
	v152 = m.ExcPending
	if v152 != 0 {
		goto L20
	} else {
		goto L54
	}
L20:
	;
	return int64(0)
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
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
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
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
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
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v78 = v77
	goto L31
L35:
	;
	v78 = v75
	goto L31
L36:
	;
	m.G0 = v11 + int32(16)
	return base.I64_extend_i32_u(v78)
L37:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v88 = F_palloc_mul(m, int32(4), v83)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
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
	v91 = int32(0)
	goto L43
L43:
	;
	v102 = F_multirange_get_range(m, v68, v81, v91)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L20
	} else {
		goto L45
	}
L44:
	;
	v110 = int32(0)
	goto L47
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88+v91<<(uint(int32(2))%32)))) = v102
	v106 = v91 + int32(1)
	if v106 != v83 {
		v91 = v106
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v119 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v88+v110<<(uint(int32(2))%32)))))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v123 = F_accumArrayResult(m, v78, v119, int32(0), v121, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L20
	} else {
		goto L49
	}
L48:
	;
	goto L36
L49:
	;
	v126 = v110 + int32(1)
	if v126 != v83 {
		v110 = v126
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v128 = F_make_empty_range(m, v68)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L20
	} else {
		goto L52
	}
L52:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v134 = F_accumArrayResult(m, v78, base.I64_extend_i32_u(v128), int32(0), v132, v133)
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
	v156 = m.ExcPending
	if v156 != 0 {
		goto L20
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_multirange_agg_transfn_2), int32(1498), int32(_a_F_multirange_agg_transfn_3))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
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
	v169 = m.ExcPending
	if v169 != 0 {
		goto L20
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_multirange_agg_transfn_2), int32(1502), int32(_a_F_multirange_agg_transfn_3))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
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
	v182 = m.ExcPending
	if v182 != 0 {
		goto L20
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_multirange_agg_transfn_2), int32(561), int32(_a_F_multirange_agg_transfn_6))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
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
func F_multirange_constructor2(m *base.Module, l0 int32) int64 {
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
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
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
	return int64(0)
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
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L50
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L47
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L43
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L40
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
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
	return base.I64_extend_i32_u(v115)
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
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v37 == int32(1) {
		goto L6
	} else {
		goto L20
	}
L19:
	;
	v115 = v35
	goto L15
L20:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v107 = F_make_multirange(m, v12, v29, v103, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L36
	}
L25:
	;
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v51
	v103 = v51
	v106 = v51
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
		v103 = v71
		v106 = v69
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
	v103 = v99
	v106 = v69
	goto L24
L33:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v75<<(uint(int32(3))%32))))
	v94 = F_pg_detoast_datum(m, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69+v75<<(uint(int32(2))%32)))) = v94
	v98 = v75 + int32(1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if v98 < v99 {
		v75 = v98
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v115 = v107
	goto L15
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
	F_errmsg_internal(m, int32(_a_F_multirange_constructor2_1), v9)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_multirange_constructor2_2), int32(561), int32(_a_F_multirange_constructor2_3))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
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
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_multirange_constructor2_2), int32(979), int32(_a_F_multirange_constructor2_5))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
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
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(_a_F_multirange_constructor2_6), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_multirange_constructor2_2), int32(987), int32(_a_F_multirange_constructor2_5))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
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
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_multirange_constructor2_2), int32(991), int32(_a_F_multirange_constructor2_5))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
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
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errmsg(m, int32(_a_F_multirange_constructor2_4), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_multirange_constructor2_2), int32(1013), int32(_a_F_multirange_constructor2_5))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
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
func F_multirange_contains_multirange(m *base.Module, l0 int32) int64 {
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
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
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
						return int64(0)
					} else {
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v34)
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_contains_multirange_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_multirange_contains_multirange_1), v9)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_contains_multirange_2), int32(561), int32(_a_F_multirange_contains_multirange_3))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int64(0)
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
								return int64(0)
							} else {
								m.G0 = v9 + int32(16)
								return base.I64_extend_i32_u(v34)
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_contains_multirange_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_multirange_contains_multirange_1), v9)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_contains_multirange_2), int32(561), int32(_a_F_multirange_contains_multirange_3))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int64(0)
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
							return int64(0)
						} else {
							m.G0 = v9 + int32(16)
							return base.I64_extend_i32_u(v34)
						}
					}
				}
			}
		}
	}
}
func F_multirange_eq(m *base.Module, l0 int32) int64 {
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
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
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
						return int64(0)
					} else {
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v34)
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_eq_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_multirange_eq_1), v9)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_eq_2), int32(561), int32(_a_F_multirange_eq_3))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int64(0)
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
								return int64(0)
							} else {
								m.G0 = v9 + int32(16)
								return base.I64_extend_i32_u(v34)
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_eq_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_multirange_eq_1), v9)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_eq_2), int32(561), int32(_a_F_multirange_eq_3))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int64(0)
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
							return int64(0)
						} else {
							m.G0 = v9 + int32(16)
							return base.I64_extend_i32_u(v34)
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
	v13 = v11 + int32(-64)
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
	m.G0 = v13 - int32(-64)
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
	v35 = v11 + int32(-16)
	v37 = v11 + int32(-32)
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
	v43 = v11 + int32(-48)
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
	F_errfinish(m, int32(_a_F_multirange_eq_internal_1), int32(1955), int32(_a_F_multirange_eq_internal_2))
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
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v123 int64
	_ = v123
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int64
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v173 int64
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int64
	_ = v195
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int64
	_ = v229
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v259 int64
	_ = v259
	var v260 int64
	_ = v260
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	v6 = int32(0)
	v14 = int64(0)
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	v21 = l1 + int32(8)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v23 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22)+11)))
	v25 = v23 & int32(255)
	if l2 <= v6 {
		v61 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v71 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+8)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v75 = v73 << (uint(int32(2)) % 32)
	v76 = v73 + v75
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v75+l2))))
	switch v25 - int32(99) {
	case 0:
		v107 = v76 + int32(8)
		v108 = int32(-1)
		goto L7
	case 1:
		goto L8
	default:
		goto L11
	case 6:
		goto L9
	case 16:
		goto L10
	}
L2:
	;
	v28 = l2
	v33 = v6
	goto L3
L3:
	;
	v43 = int32(2)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v21+v28<<(uint(v43)%32))))
	v49 = v46&int32(2147483647) + v33
	if base.Ui32(v28) < base.Ui32(v43) {
		v61 = v49
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v61 = v49
	goto L1
L5:
	;
	if int32(0) <= v46 {
		v28 = v28 - int32(1)
		v33 = v49
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	v111 = l1 + v107&v108 + v61
	if v81&int32(41) != 0 {
		v173 = v14
		v174 = v111
		goto L22
	} else {
		goto L23
	}
L8:
	;
	v107 = v76 + int32(15)
	v108 = int32(-8)
	goto L7
L9:
	;
	v107 = v76 + int32(11)
	v108 = int32(-4)
	goto L7
L10:
	;
	v107 = v76 + int32(9)
	v108 = int32(-2)
	goto L7
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v23
	F_errmsg_internal(m, int32(_a_F_multirange_get_bounds_0), v18)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_multirange_get_bounds_1), int32(322), int32(_a_F_multirange_get_bounds_2))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	v261 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+10)) = uint8(v261)
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v260
	v267 = int32(base.Ui32(v81)>>(uint(v261)%32)) & v261
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v267)
	v272 = int32(base.Ui32(v81)>>(uint(int32(3))%32)) & v261
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v272)
	v274 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+10)) = uint8(v274)
	v279 = int32(base.Ui32(v81)>>(uint(int32(2))%32)) & v261
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+9)) = uint8(v279)
	v284 = int32(base.Ui32(v81)>>(uint(int32(4))%32)) & v261
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)) = uint8(v284)
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = v259
	m.G0 = v18 - int32(-64)
	return
L17:
	;
	if v72&int32(1) != 0 {
		goto L63
	} else {
		goto L64
	}
L18:
	;
	switch v25 - int32(99) {
	case 0:
		v223 = int32(-1)
		v224 = v193
		goto L55
	case 1:
		goto L56
	default:
		goto L59
	case 6:
		goto L57
	case 16:
		goto L58
	}
L19:
	;
	v190 = int32(-1)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v191 != 0 {
		v227 = v188
		v228 = v190
		v229 = v189
		goto L17
	} else {
		goto L54
	}
L20:
	;
	if v81&int32(80) != 0 {
		v259 = v14
		v260 = v140
		goto L16
	} else {
		goto L53
	}
L21:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v183 = int32(base.Ui32(v179) >> (uint(int32(2)) % 32))
	goto L20
L22:
	;
	if v81&int32(81) != 0 {
		v259 = v14
		v260 = v173
		goto L16
	} else {
		goto L51
	}
L23:
	;
	if v72&int32(1) != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if int32(0) < v71 {
		v173 = v140
		v174 = v111 + v71
		goto L22
	} else {
		goto L37
	}
L25:
	;
	if base.I32_popcnt(v71) != int32(1) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v140 = base.I64_extend_i32_u(v111)
	goto L24
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L12
	} else {
		goto L34
	}
L29:
	;
	switch base.I32_ctz(v71) {
	case 0:
		goto L33
	case 1:
		goto L32
	case 2:
		goto L31
	case 3:
		goto L30
	default:
		goto L28
	}
L30:
	;
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v111)))
	v140 = v123
	goto L24
L31:
	;
	v122 = int64(*(*int32)(unsafe.Add(mBase, uint32(v111))))
	v140 = v122
	goto L24
L32:
	;
	v121 = int64(*(*int16)(unsafe.Add(mBase, uint32(v111))))
	v140 = v121
	goto L24
L33:
	;
	v120 = int64(*(*int8)(unsafe.Add(mBase, uint32(v111))))
	v140 = v120
	goto L24
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v71
	F_errmsg_internal(m, int32(_a_F_multirange_get_bounds_3), v16+int32(-48))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_multirange_get_bounds_1), int32(123), int32(_a_F_multirange_get_bounds_4))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	if v71 == int32(-1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	if v146 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v169 = F_strlen(m, v111)
	mBase = m.M
	v173 = v140
	v174 = v169 + v111 + int32(1)
	goto L22
L41:
	;
	v150 = int32(18)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
	if v152 == v150 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	if v146&int32(1) == int32(0) {
		goto L21
	} else {
		goto L50
	}
L44:
	;
	v155 = v150
	goto L46
L45:
	;
	v155 = int32(2)
	goto L46
L46:
	;
	if base.Ui32((v152-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v162 = int32(6)
	goto L49
L48:
	;
	v162 = v155
	goto L49
L49:
	;
	v183 = v162
	goto L20
L50:
	;
	v183 = int32(base.Ui32(v146) >> (uint(int32(1)) % 32))
	goto L20
L51:
	;
	if v71 == int32(-1) {
		v188 = v174
		v189 = v173
		goto L19
	} else {
		goto L52
	}
L52:
	;
	v193 = v174
	v194 = v71
	v195 = v173
	goto L18
L53:
	;
	v188 = v183 + v111
	v189 = v140
	goto L19
L54:
	;
	v193 = v188
	v194 = v190
	v195 = v189
	goto L18
L55:
	;
	v227 = v223 & v224
	v228 = v194
	v229 = v195
	goto L17
L56:
	;
	v223 = int32(-8)
	v224 = v193 + int32(7)
	goto L55
L57:
	;
	v223 = int32(-4)
	v224 = v193 + int32(3)
	goto L55
L58:
	;
	v223 = int32(-2)
	v224 = v193 + int32(1)
	goto L55
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L12
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v25
	F_errmsg_internal(m, int32(_a_F_multirange_get_bounds_0), v16+int32(-32))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L12
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_multirange_get_bounds_1), int32(322), int32(_a_F_multirange_get_bounds_2))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	if base.I32_popcnt(v228) != int32(1) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v259 = base.I64_extend_i32_u(v227)
	v260 = v229
	goto L16
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L12
	} else {
		goto L72
	}
L67:
	;
	switch base.I32_ctz(v228) {
	case 0:
		goto L71
	case 1:
		goto L70
	case 2:
		goto L69
	case 3:
		goto L68
	default:
		goto L66
	}
L68:
	;
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v227)))
	v259 = v239
	v260 = v229
	goto L16
L69:
	;
	v238 = int64(*(*int32)(unsafe.Add(mBase, uint32(v227))))
	v259 = v238
	v260 = v229
	goto L16
L70:
	;
	v237 = int64(*(*int16)(unsafe.Add(mBase, uint32(v227))))
	v259 = v237
	v260 = v229
	goto L16
L71:
	;
	v236 = int64(*(*int8)(unsafe.Add(mBase, uint32(v227))))
	v259 = v236
	v260 = v229
	goto L16
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v228
	F_errmsg_internal(m, int32(_a_F_multirange_get_bounds_3), v16+int32(-16))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_multirange_get_bounds_1), int32(123), int32(_a_F_multirange_get_bounds_4))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
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
							F_errfinish(m, int32(_a_F_multirange_get_typcache_2), int32(561), int32(_a_F_multirange_get_typcache_3))
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
						F_errfinish(m, int32(_a_F_multirange_get_typcache_2), int32(561), int32(_a_F_multirange_get_typcache_3))
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
func F_multirange_lower_inc(m *base.Module, l0 int32) int64 {
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
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v43 int64
	_ = v43
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
		if v16 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
			if v19 != 0 {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				if v20 == v17 {
					v30 = v19
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+296))
					F_multirange_get_bounds(m, v31, v12, int32(0), v9+int32(32), v9+int32(16))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int64(0)
					} else {
						v39 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v9)+41)))
						v43 = v39
						m.G0 = v9 + int32(48)
						return v43
					}
				} else {
					v23 = F_lookup_type_cache(m, v17, int32(_a_F_multirange_lower_inc_0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int64(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+296))
						if v25 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
								F_errmsg_internal(m, int32(_a_F_multirange_lower_inc_1), v9)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_lower_inc_2), int32(561), int32(_a_F_multirange_lower_inc_3))
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int64(0)
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
							v30 = v23
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+296))
							F_multirange_get_bounds(m, v31, v12, int32(0), v9+int32(32), v9+int32(16))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int64(0)
							} else {
								v39 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v9)+41)))
								v43 = v39
								m.G0 = v9 + int32(48)
								return v43
							}
						}
					}
				}
			} else {
				v23 = F_lookup_type_cache(m, v17, int32(_a_F_multirange_lower_inc_0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+296))
					if v25 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
							F_errmsg_internal(m, int32(_a_F_multirange_lower_inc_1), v9)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_lower_inc_2), int32(561), int32(_a_F_multirange_lower_inc_3))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int64(0)
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
						v30 = v23
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+296))
						F_multirange_get_bounds(m, v31, v12, int32(0), v9+int32(32), v9+int32(16))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int64(0)
						} else {
							v39 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v9)+41)))
							v43 = v39
							m.G0 = v9 + int32(48)
							return v43
						}
					}
				}
			}
		} else {
			v43 = int64(0)
			m.G0 = v9 + int32(48)
			return v43
		}
	}
}
func F_multirange_lt(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_multirange_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return int64(base.Ui64(v2) >> (uint(int64(63)) % 64))
	}
}
func F_multirange_minus(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
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
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v23 = F_pg_detoast_datum(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L32
	}
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v39 == int32(0) {
		v129 = v18
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v28 == v25 {
		v38 = v27
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v31 = F_lookup_type_cache(m, v25, int32(_a_F_multirange_minus_0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+296))
	if v33 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v31
	v38 = v31
	goto L5
L12:
	;
	m.G0 = v15 + int32(16)
	return base.I64_extend_i32_u(v129)
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v42 == int32(0) {
		v129 = v18
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
	if int32(0) < v39 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v50 = F_palloc_mul(m, int32(4), v39)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v78 = v42
	v84 = v45
	v85 = v2
	goto L17
L17:
	;
	if int32(0) < v78 {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v52 = int32(0)
	goto L19
L19:
	;
	v67 = F_multirange_get_range(m, v45, v18, v52)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v78 = v74
	v84 = v73
	v85 = v50
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50+v52<<(uint(int32(2))%32)))) = v67
	v71 = v52 + int32(1)
	if v71 != v39 {
		v52 = v71
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v91 = F_palloc_mul(m, int32(4), v78)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v125 = v2
	goto L25
L25:
	;
	v126 = F_multirange_minus_internal(m, v25, v45, v39, v85, v78, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L31
	}
L26:
	;
	v93 = int32(0)
	goto L27
L27:
	;
	v108 = F_multirange_get_range(m, v84, v23, v93)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v125 = v91
	goto L25
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91+v93<<(uint(int32(2))%32)))) = v108
	v112 = v93 + int32(1)
	if v112 != v78 {
		v93 = v112
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v129 = v126
	goto L12
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v25
	F_errmsg_internal(m, int32(_a_F_multirange_minus_1), v15)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_multirange_minus_2), int32(561), int32(_a_F_multirange_minus_3))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_multirange_minus_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int64
	_ = v214
	var v215 int64
	_ = v215
	var v216 int64
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
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
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	v7 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v26 = F_palloc0(m, (l2+l4)<<(uint(int32(2))%32))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if int32(0) < l2 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v43 = v32
	v44 = v7
	v45 = v7
	v48 = v7
	goto L6
L4:
	;
	v389 = v7
	goto L5
L5:
	;
	v395 = F_make_multirange(m, l0, l1, v389, v26)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L101
	}
L6:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l3+v48<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v54
	if v43 == int32(0) {
		v341 = v44
		v342 = v45
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v389 = v373
	goto L5
L8:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v26+v361<<(uint(int32(2))%32)))) = v370
	v372 = int32(1)
	v373 = v361 + v372
	v375 = v48 + v372
	if v375 != l2 {
		v43 = v359
		v44 = v360
		v45 = v373
		v48 = v375
		goto L6
	} else {
		goto L100
	}
L9:
	;
	v359 = int32(0)
	v360 = v341
	v361 = v342
	goto L8
L10:
	;
	v68 = v43
	v69 = v44
	goto L11
L11:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v77 = F_range_before_internal(m, l1, v68, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v96 = v68
	v97 = v69
	v98 = v45
	goto L19
L13:
	;
	if v77 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v80 = v69 + int32(1)
	if l4 <= v80 {
		v341 = v80
		v342 = v45
		goto L9
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	goto L12
L17:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l5+v80<<(uint(int32(2))%32))))
	if v85 != 0 {
		v68 = v85
		v69 = v80
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v341 = v80
	v342 = v45
	goto L9
L19:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v110 = m.G0
	v112 = v110 - int32(80)
	m.G0 = v112
	F_range_deserialize(m, l1, v104, v112-int32(-64), v112+int32(32), v112+int32(15))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v341 = v324
	v342 = v325
	goto L9
L21:
	;
	F_range_deserialize(m, l1, v96, v112+int32(48), v112+int32(16), v112+int32(14))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+56)))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+72)))
	if v131 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	m.G0 = v112 + int32(80)
	if v288 != 0 {
		goto L87
	} else {
		goto L88
	}
L24:
	;
	v257 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+58)) = uint8(v257)
	v259 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+26)) = uint8(v259)
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+57)))
	v264 = v262 ^ v259
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+57)) = uint8(v264)
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+25)))
	v268 = v266 ^ v259
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+25)) = uint8(v268)
	v276 = F_make_range(m, l1, v112-int32(-64), v112+int32(48), v257, v257)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L84
	}
L25:
	;
	v288 = int32(0)
	goto L23
L26:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+24)))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+40)))
	if v191 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L27:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+74)))
	if v130&int32(1) != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	if v130&int32(1) != 0 {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	v137 = int32(0)
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+58)))
	if v138 == v134 {
		v288 = v137
		goto L23
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v134&int32(1) != 0 {
		goto L26
	} else {
		goto L35
	}
L33:
	;
	if v134&int32(1) != 0 {
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v288 = v137
	goto L23
L35:
	;
	goto L25
L36:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+58)))
	if v146 == int32(0) {
		goto L26
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v112)+64))
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v112)+48))
	v154 = F_FunctionCall2Coll(m, l1+int32(212), v151, v152, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	goto L25
L40:
	;
	v156 = base.I32_wrap_i64(v154)
	if v156 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+57)))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+73)))
	if v160 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	if int32(0) <= v156 {
		goto L25
	} else {
		goto L55
	}
L44:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+74)))
	if v159&int32(1) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v179 = int32(0)
	if v159&int32(1) != 0 {
		v288 = v179
		goto L23
	} else {
		goto L53
	}
L47:
	;
	v168 = int32(0)
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+58)))
	if v163 == v169 {
		v288 = v168
		goto L23
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v163&int32(1) == int32(0) {
		goto L26
	} else {
		goto L52
	}
L50:
	;
	if v163&int32(1) == int32(0) {
		goto L26
	} else {
		goto L51
	}
L51:
	;
	v288 = v168
	goto L23
L52:
	;
	goto L25
L53:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+58)))
	if v182&int32(1) != 0 {
		goto L26
	} else {
		goto L54
	}
L54:
	;
	v288 = v179
	goto L23
L55:
	;
	goto L26
L56:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+42)))
	if v190&int32(1) != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	if v190&int32(1) != 0 {
		goto L65
	} else {
		goto L66
	}
L59:
	;
	v197 = int32(0)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+26)))
	if v198 == v194 {
		v288 = v197
		goto L23
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v194&int32(1) == int32(0) {
		goto L24
	} else {
		goto L64
	}
L62:
	;
	if v194&int32(1) == int32(0) {
		goto L24
	} else {
		goto L63
	}
L63:
	;
	v288 = v197
	goto L23
L64:
	;
	goto L25
L65:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+26)))
	if v210 != 0 {
		goto L24
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
	v214 = *(*int64)(unsafe.Add(mBase, uint32(v112)+32))
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v112)+16))
	v216 = F_FunctionCall2Coll(m, l1+int32(212), v213, v214, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	goto L25
L69:
	;
	v218 = base.I32_wrap_i64(v216)
	if v218 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+25)))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+41)))
	if v222 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	if int32(0) < v218 {
		goto L24
	} else {
		goto L83
	}
L73:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+42)))
	if v221&int32(1) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v240 = int32(0)
	if v221&int32(1) != 0 {
		v288 = v240
		goto L23
	} else {
		goto L81
	}
L76:
	;
	v230 = int32(0)
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+26)))
	if base.B2i32(v225&int32(1) == v230)|base.B2i32(v225 == v235) != 0 {
		v288 = v230
		goto L23
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if v225&int32(1) != 0 {
		goto L24
	} else {
		goto L80
	}
L79:
	;
	goto L24
L80:
	;
	goto L25
L81:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+26)))
	if v243&int32(1) == int32(0) {
		goto L24
	} else {
		goto L82
	}
L82:
	;
	v288 = v240
	goto L23
L83:
	;
	goto L25
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+v98<<(uint(int32(2))%32)))) = v276
	v283 = int32(0)
	v285 = F_make_range(m, l1, v112+int32(16), v112+int32(32), v283, v283)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(12)))) = v285
	v288 = v259
	goto L23
L86:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l5+v324<<(uint(int32(2))%32))))
	if v329 != 0 {
		v96 = v329
		v97 = v324
		v98 = v325
		goto L19
	} else {
		goto L99
	}
L87:
	;
	v294 = int32(1)
	v295 = v98 + v294
	v297 = v97 + v294
	if v297 < l4 {
		v324 = v297
		v325 = v295
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v300 = F_range_overlaps_internal(m, l1, v299, v96)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L91
	}
L90:
	;
	v341 = v297
	v342 = v295
	goto L9
L91:
	;
	if v300 == int32(0) {
		v359 = v96
		v360 = v97
		v361 = v98
		goto L8
	} else {
		goto L92
	}
L92:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v305 = F_range_minus_internal(m, l1, v304, v96)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v305
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	v314 = int32(*(*int8)(unsafe.Add(mBase, uint32(v305+int32(base.Ui32(v308)>>(uint(int32(2))%32))-int32(1)))))
	goto L94
L94:
	;
	if v314&int32(1) != 0 {
		v359 = v96
		v360 = v97
		v361 = v98
		goto L8
	} else {
		goto L95
	}
L95:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v318 = F_range_before_internal(m, l1, v317, v96)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	if v318 != 0 {
		v359 = v96
		v360 = v97
		v361 = v98
		goto L8
	} else {
		goto L97
	}
L97:
	;
	v321 = v97 + int32(1)
	if l4 <= v321 {
		v341 = v321
		v342 = v98
		goto L9
	} else {
		goto L98
	}
L98:
	;
	v324 = v321
	v325 = v98
	goto L86
L99:
	;
	goto L20
L100:
	;
	goto L7
L101:
	;
	m.G0 = v21 + int32(16)
	return v395
}
func F_multirange_overleft_multirange(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v73 int64
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v8 = int64(0)
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v19 = F_pg_detoast_datum(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
			if v21 == int32(0) {
				v73 = v8
				m.G0 = v11 + int32(80)
				return v73
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
				if v24 == int32(0) {
					v73 = v8
					m.G0 = v11 + int32(80)
					return v73
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
					if v29 != 0 {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
						if v30 == v27 {
							v41 = v29
							v42 = v21
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+296))
							v49 = v11 + int32(48)
							F_multirange_get_bounds(m, v43, v14, v42-int32(1), v11-int32(-64), v49)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int64(0)
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v41)+296))
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
								v59 = v11 + int32(16)
								F_multirange_get_bounds(m, v52, v19, v53-int32(1), v11+int32(32), v59)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int64(0)
								} else {
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v41)+296))
									v63 = F_range_cmp_bounds(m, v62, v49, v59)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int64(0)
									} else {
										v73 = base.I64_extend_i32_u(base.B2i32(v63 <= int32(0)))
										m.G0 = v11 + int32(80)
										return v73
									}
								}
							}
						} else {
							v33 = F_lookup_type_cache(m, v27, int32(_a_F_multirange_overleft_multirange_0))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int64(0)
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+296))
								if v35 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = v27
										F_errmsg_internal(m, int32(_a_F_multirange_overleft_multirange_1), v11)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_multirange_overleft_multirange_2), int32(561), int32(_a_F_multirange_overleft_multirange_3))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v33
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
									v41 = v33
									v42 = v40
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+296))
									v49 = v11 + int32(48)
									F_multirange_get_bounds(m, v43, v14, v42-int32(1), v11-int32(-64), v49)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int64(0)
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v41)+296))
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
										v59 = v11 + int32(16)
										F_multirange_get_bounds(m, v52, v19, v53-int32(1), v11+int32(32), v59)
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int64(0)
										} else {
											v62 = *(*int32)(unsafe.Add(mBase, uint32(v41)+296))
											v63 = F_range_cmp_bounds(m, v62, v49, v59)
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int64(0)
											} else {
												v73 = base.I64_extend_i32_u(base.B2i32(v63 <= int32(0)))
												m.G0 = v11 + int32(80)
												return v73
											}
										}
									}
								}
							}
						}
					} else {
						v33 = F_lookup_type_cache(m, v27, int32(_a_F_multirange_overleft_multirange_0))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+296))
							if v35 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v27
									F_errmsg_internal(m, int32(_a_F_multirange_overleft_multirange_1), v11)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_multirange_overleft_multirange_2), int32(561), int32(_a_F_multirange_overleft_multirange_3))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v33
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
								v41 = v33
								v42 = v40
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+296))
								v49 = v11 + int32(48)
								F_multirange_get_bounds(m, v43, v14, v42-int32(1), v11-int32(-64), v49)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int64(0)
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v41)+296))
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
									v59 = v11 + int32(16)
									F_multirange_get_bounds(m, v52, v19, v53-int32(1), v11+int32(32), v59)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int64(0)
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(v41)+296))
										v63 = F_range_cmp_bounds(m, v62, v49, v59)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int64(0)
										} else {
											v73 = base.I64_extend_i32_u(base.B2i32(v63 <= int32(0)))
											m.G0 = v11 + int32(80)
											return v73
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
func F_multirange_recv(m *base.Module, l0 int32) int64 {
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
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
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return int64(0)
L2:
	;
	v24 = int32(4)
	v26 = F_pq_getmsgint(m, v17, v24)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = F_palloc_mul(m, v24, v26)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_initStringInfo(m, v14)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if int32(0) < v26 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v37 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	F_pfree(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L18
	}
L9:
	;
	v49 = F_pq_getmsgint(m, v17, int32(4))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v51 = F_pq_getmsgbytes(m, v17, v49)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v54
	goto L13
L13:
	;
	F_appendBinaryStringInfo(m, v14, v51, v49)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	v66 = F_ReceiveFunctionCall(m, v20+int32(4), v14, v65, v16)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v69 = F_pg_detoast_datum(m, base.I32_wrap_i64(v66))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28+v37<<(uint(int32(2))%32)))) = v69
	v73 = v37 + int32(1)
	if v73 != v26 {
		v37 = v73
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
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+296))
	v93 = F_make_multirange(m, v18, v92, v26, v28)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	m.G0 = v14 + int32(16)
	return base.I64_extend_i32_u(v93)
}
func F_multirange_upper_inf(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v48 int64
	_ = v48
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		if v17 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
			if v20 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				if v21 == v18 {
					v32 = v20
					v33 = v17
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
					F_multirange_get_bounds(m, v34, v13, v33-int32(1), v10+int32(32), v10+int32(16))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						v43 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
						v48 = v43
						m.G0 = v10 + int32(48)
						return v48
					}
				} else {
					v24 = F_lookup_type_cache(m, v18, int32(_a_F_multirange_upper_inf_0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int64(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+296))
						if v26 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
								F_errmsg_internal(m, int32(_a_F_multirange_upper_inf_1), v10)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_upper_inf_2), int32(561), int32(_a_F_multirange_upper_inf_3))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v24
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
							v32 = v24
							v33 = v31
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							F_multirange_get_bounds(m, v34, v13, v33-int32(1), v10+int32(32), v10+int32(16))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int64(0)
							} else {
								v43 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
								v48 = v43
								m.G0 = v10 + int32(48)
								return v48
							}
						}
					}
				}
			} else {
				v24 = F_lookup_type_cache(m, v18, int32(_a_F_multirange_upper_inf_0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+296))
					if v26 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
							F_errmsg_internal(m, int32(_a_F_multirange_upper_inf_1), v10)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_upper_inf_2), int32(561), int32(_a_F_multirange_upper_inf_3))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v24
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
						v32 = v24
						v33 = v31
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						F_multirange_get_bounds(m, v34, v13, v33-int32(1), v10+int32(32), v10+int32(16))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int64(0)
						} else {
							v43 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
							v48 = v43
							m.G0 = v10 + int32(48)
							return v48
						}
					}
				}
			}
		} else {
			v48 = int64(0)
			m.G0 = v10 + int32(48)
			return v48
		}
	}
}
