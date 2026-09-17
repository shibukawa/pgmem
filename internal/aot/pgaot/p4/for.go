package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_for_column_name_collision(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v13 = F_SearchSysCache2(m, int32(6), v12, l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			m.G0 = v9 + int32(48)
			return base.B2i32(v13 == int32(0))
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
			v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19+v20)+74)))
			F_ReleaseCatCache(m, v13)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v22 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(16806020))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
							F_errmsg(m, int32(_a_F_check_for_column_name_collision_0), v9)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_check_for_column_name_collision_1), int32(_a_F_check_for_column_name_collision_2), int32(_a_F_check_for_column_name_collision_3))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
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
					if l2 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(16806020))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v84 + int32(4)
								F_errmsg(m, int32(_a_F_check_for_column_name_collision_4), v9+int32(32))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_check_for_column_name_collision_1), int32(_a_F_check_for_column_name_collision_5), int32(_a_F_check_for_column_name_collision_3))
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
					} else {
						v31 = F_errstart(m, int32(18), int32(0))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							if v31 == int32(0) {
								m.G0 = v9 + int32(48)
								return base.B2i32(v13 == int32(0))
							} else {
								F_errcode(m, int32(16806020))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v38 + int32(4)
									F_errmsg(m, int32(_a_F_check_for_column_name_collision_6), v9+int32(16))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_check_for_column_name_collision_1), int32(_a_F_check_for_column_name_collision_7), int32(_a_F_check_for_column_name_collision_3))
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											m.G0 = v9 + int32(48)
											return base.B2i32(v13 == int32(0))
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
func F_exec_for_query(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v100 int32
	_ = v100
	var v109 int64
	_ = v109
	var v112 int64
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int64
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v186 int32
	_ = v186
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v248 int32
	_ = v248
	var v249 int64
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v275 int64
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v367 int64
	_ = v367
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int64
	_ = v376
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	v5 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18+v20<<(uint(int32(2))%32))))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+84)))
	if v25 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v43 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+84)) = uint8(v43)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	v48 = l3 & v47
	if v48 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	return int32(0)
L5:
	;
	F_errmsg_internal(m, int32(_a_F_exec_for_query_0), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_errfinish(m, int32(_a_F_exec_for_query_1), int32(374), int32(_a_F_exec_for_query_2))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L8:
	;
	v49 = int32(10)
	goto L10
L9:
	;
	v49 = v43
	goto L10
L10:
	;
	F_SPI_cursor_fetch(m, l2, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_exec_for_query[0]))
	v55 = *(*int64)(unsafe.Add(mBase, _c_F_exec_for_query[1]))
	if v55 == int64(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	F_SPI_freetuptable(m, v382)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L4
	} else {
		goto L109
	}
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	F_exec_move_row(m, l0, v24, int32(0), v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	if v48 != 0 {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v62 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_SPI_freetuptable(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v65 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v67 == v65 {
		v382 = v53
		v384 = v5
		goto L12
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	F_MemoryContextReset(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v382 = v53
	v384 = v5
	goto L12
L23:
	;
	v75 = int32(50)
	goto L25
L24:
	;
	v75 = int32(1)
	goto L25
L25:
	;
	v81 = v53
	v82 = int32(1)
	v93 = v55
	v94 = int64(1)
	goto L26
L26:
	;
	v100 = v82
	v109 = int64(0)
	v112 = v94
	goto L28
L27:
	;
	v382 = v374
	v384 = v363
	goto L12
L28:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v113 == int32(2) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	F_SPI_freetuptable(m, v81)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L4
	} else {
		goto L106
	}
L30:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v276 != 0 {
		goto L69
	} else {
		goto L70
	}
L31:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v116 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v250+base.I32_wrap_i64(v109)<<(uint(int32(2))%32))))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	F_exec_move_row(m, l0, v24, v255, v256)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L4
	} else {
		goto L68
	}
L34:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v139+base.I32_wrap_i64(v109)<<(uint(int32(2))%32))))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	F_exec_move_row(m, l0, v24, v144, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L38
	}
L35:
	;
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v116)+48))
	if base.B2i32(v119 == v112)&v100 == int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v124 = int32(1)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v125+base.I32_wrap_i64(v109)<<(uint(int32(2))%32))))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	F_expanded_record_set_tuple(m, v116, v130, v124, (v132^int32(-1))&v124)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v263 = v124
	v275 = v112
	goto L30
L38:
	;
	v148 = int32(0)
	if v100 == v148 {
		v235 = v148
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	v249 = *(*int64)(unsafe.Add(mBase, uint32(v248)+48))
	v263 = v235
	v275 = v249
	goto L30
L40:
	;
	v151 = int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	if v152 == int32(2249) {
		v235 = v151
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if v152 == v156 {
		v235 = v151
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+44))
	if v159 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v162 = F_expanded_record_fetch_tupdesc(m, v158)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	v164 = v159
	goto L45
L45:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	if v166 != v167 {
		v235 = int32(0)
		goto L39
	} else {
		goto L47
	}
L46:
	;
	v164 = v162
	goto L45
L47:
	;
	if v166 <= int32(0) {
		v235 = int32(1)
		goto L39
	} else {
		goto L48
	}
L48:
	;
	v173 = v166 << (uint(int32(4)) % 32)
	v175 = int32(20)
	v186 = int32(0)
	goto L49
L49:
	;
	v199 = v186 * int32(100)
	v200 = v173 + v164 + v175 + v199
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+91)))
	v202 = v199 + (v155 + v173 + v175)
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+91)))
	if v201 != v203 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v235 = v227
	goto L39
L51:
	;
	v235 = int32(0)
	goto L39
L52:
	;
	goto L53
L53:
	;
	if v201 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v227 = int32(1)
	v229 = v186 + v227
	if v229 != v166 {
		v186 = v229
		goto L49
	} else {
		goto L67
	}
L55:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v200)+68))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v202)+68))
	if v208 != v209 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200)+72)))
	v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202)+72)))
	if v218 != v219 {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	v235 = int32(0)
	goto L39
L59:
	;
	goto L60
L60:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v200)+76))
	if v212 < int32(0) {
		goto L54
	} else {
		goto L61
	}
L61:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v202)+76))
	if v212 == v215 {
		goto L54
	} else {
		goto L62
	}
L62:
	;
	v235 = int32(0)
	goto L39
L63:
	;
	v235 = int32(0)
	goto L39
L64:
	;
	goto L65
L65:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+83)))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+83)))
	if v222 == v223 {
		goto L54
	} else {
		goto L66
	}
L66:
	;
	v235 = int32(0)
	goto L39
L67:
	;
	goto L50
L68:
	;
	v263 = v100
	v275 = v112
	goto L30
L69:
	;
	F_SPI_freetuptable(m, v276)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L4
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v281 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L71
L73:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+20))
	F_MemoryContextReset(m, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L4
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v286 = F_exec_stmts(m, l0, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L80
	}
L76:
	;
	goto L75
L77:
	;
	v367 = v109 + int64(1)
	if v367 != v93 {
		v100 = v263
		v109 = v367
		v112 = v275
		goto L28
	} else {
		goto L105
	}
L78:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v327 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L79:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v290 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	switch v286 - int32(1) {
	case 0:
		goto L79
	case 1:
		v382 = v81
		v384 = v286
		goto L12
	case 2:
		goto L78
	default:
		v363 = v286
		goto L77
	}
L81:
	;
	v382 = v81
	v384 = int32(0)
	goto L12
L82:
	;
	goto L83
L83:
	;
	v294 = int32(1)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v295 == int32(0) {
		v382 = v81
		v384 = v294
		goto L12
	} else {
		goto L84
	}
L84:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290))))
	if base.B2i32(v300 == int32(0))|base.B2i32(v300 != v303) != 0 {
		v321 = v300
		v322 = v303
		goto L86
	} else {
		goto L87
	}
L85:
	;
	if v321-v322 != 0 {
		v382 = v81
		v384 = v294
		goto L12
	} else {
		goto L92
	}
L86:
	;
	goto L85
L87:
	;
	v306 = v295
	v307 = v290
	goto L88
L88:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+1)))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306)+1)))
	if v311 == int32(0) {
		v321 = v311
		v322 = v310
		goto L86
	} else {
		goto L90
	}
L89:
	;
	v321 = v311
	v322 = v310
	goto L86
L90:
	;
	v314 = int32(1)
	if v311 == v310 {
		v306 = v306 + v314
		v307 = v307 + v314
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v324 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v324
	v382 = v81
	v384 = v324
	goto L12
L93:
	;
	v363 = int32(0)
	goto L77
L94:
	;
	goto L95
L95:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v331 == int32(0) {
		v382 = v81
		v384 = v286
		goto L12
	} else {
		goto L96
	}
L96:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	if base.B2i32(v336 == int32(0))|base.B2i32(v336 != v339) != 0 {
		v357 = v336
		v358 = v339
		goto L98
	} else {
		goto L99
	}
L97:
	;
	if v357-v358 != 0 {
		v382 = v81
		v384 = v286
		goto L12
	} else {
		goto L104
	}
L98:
	;
	goto L97
L99:
	;
	v342 = v331
	v343 = v327
	goto L100
L100:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+1)))
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+1)))
	if v347 == int32(0) {
		v357 = v347
		v358 = v346
		goto L98
	} else {
		goto L102
	}
L101:
	;
	v357 = v347
	v358 = v346
	goto L98
L102:
	;
	v350 = int32(1)
	if v347 == v346 {
		v342 = v342 + v350
		v343 = v343 + v350
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v360 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v360
	v363 = v360
	goto L77
L105:
	;
	goto L29
L106:
	;
	F_SPI_cursor_fetch(m, l2, v75)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	v374 = *(*int32)(unsafe.Add(mBase, _c_F_exec_for_query[0]))
	v376 = *(*int64)(unsafe.Add(mBase, _c_F_exec_for_query[1]))
	if v376 != int64(0) {
		v81 = v374
		v82 = v263
		v93 = v376
		v94 = v275
		goto L26
	} else {
		goto L108
	}
L108:
	;
	goto L27
L109:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+84)))
	if v398 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L4
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v414 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+84)) = uint8(v414)
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v416+v417<<(uint(int32(2))%32))))
	F_assign_simple_var(m, l0, v421, base.B2i32(v55 != int64(0)), v414, v414)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L4
	} else {
		goto L116
	}
L113:
	;
	F_errmsg_internal(m, int32(_a_F_exec_for_query_3), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_exec_for_query_1), int32(383), int32(_a_F_exec_for_query_4))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L4
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
	return v384
}
