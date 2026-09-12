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
							F_errmsg(m, int32(375615), v9)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(486904), int32(7674), int32(268088))
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
								F_errmsg(m, int32(114788), v9+int32(32))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(486904), int32(7689), int32(268088))
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
									F_errmsg(m, int32(328147), v9+int32(16))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(486904), int32(7682), int32(268088))
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
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v102 int32
	_ = v102
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int64
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
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
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
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
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v251 int32
	_ = v251
	var v252 int64
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v278 int64
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v368 int64
	_ = v368
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int64
	_ = v377
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
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
	F_errmsg_internal(m, int32(445887), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_errfinish(m, int32(489522), int32(374), int32(305312))
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
	v53 = *(*int32)(unsafe.Add(mBase, _consts[501]))
	v55 = *(*int64)(unsafe.Add(mBase, _consts[500]))
	if v55 == int64(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	F_SPI_freetuptable(m, v383)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L4
	} else {
		goto L111
	}
L13:
	;
	v58 = int32(0)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	F_exec_move_row(m, l0, v24, v58, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
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
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v63 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_SPI_freetuptable(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v66 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v68 == v66 {
		v383 = v53
		v384 = v58
		goto L12
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	F_MemoryContextReset(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v383 = v53
	v384 = v58
	goto L12
L23:
	;
	v76 = int32(50)
	goto L25
L24:
	;
	v76 = int32(1)
	goto L25
L25:
	;
	v82 = v53
	v84 = int32(1)
	v94 = v55
	v95 = int64(1)
	goto L26
L26:
	;
	v102 = v84
	v110 = int64(0)
	v113 = v95
	goto L28
L27:
	;
	v383 = v375
	v384 = v364
	goto L12
L28:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v114 == int32(2) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	F_SPI_freetuptable(m, v82)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L4
	} else {
		goto L108
	}
L30:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v279 != 0 {
		goto L69
	} else {
		goto L70
	}
L31:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v117 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v253+base.I32_wrap_i64(v110)<<(uint(int32(2))%32))))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	F_exec_move_row(m, l0, v24, v258, v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L4
	} else {
		goto L68
	}
L34:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v140+base.I32_wrap_i64(v110)<<(uint(int32(2))%32))))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	F_exec_move_row(m, l0, v24, v145, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L38
	}
L35:
	;
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v117)+48))
	if base.B2i32(v120 == v113)&v102 == int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v125 = int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v126+base.I32_wrap_i64(v110)<<(uint(int32(2))%32))))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	F_expanded_record_set_tuple(m, v117, v131, v125, (v133^int32(-1))&v125)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v267 = v125
	v278 = v113
	goto L30
L38:
	;
	v151 = int32(0)
	if v102&int32(1) == v151 {
		v239 = v151
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v251)+48))
	v267 = v239
	v278 = v252
	goto L30
L40:
	;
	v154 = int32(1)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	if v155 == int32(2249) {
		v239 = v154
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v155 == v159 {
		v239 = v154
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+44))
	if v162 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v165 = F_expanded_record_fetch_tupdesc(m, v161)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	v167 = v162
	goto L45
L45:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v169 != v170 {
		v239 = int32(0)
		goto L39
	} else {
		goto L47
	}
L46:
	;
	v167 = v165
	goto L45
L47:
	;
	if v169 <= int32(0) {
		v239 = int32(1)
		goto L39
	} else {
		goto L48
	}
L48:
	;
	v176 = v169 << (uint(int32(4)) % 32)
	v178 = int32(20)
	v188 = int32(0)
	goto L49
L49:
	;
	v202 = v188 * int32(100)
	v203 = v167 + v176 + v178 + v202
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+91)))
	v205 = v202 + (v158 + v176 + v178)
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+91)))
	if v204 != v206 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v239 = v230
	goto L39
L51:
	;
	v239 = int32(0)
	goto L39
L52:
	;
	goto L53
L53:
	;
	if v204 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v230 = int32(1)
	v232 = v188 + v230
	if v232 != v169 {
		v188 = v232
		goto L49
	} else {
		goto L67
	}
L55:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v205)+68))
	if v211 != v212 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v203)+72)))
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v205)+72)))
	if v221 != v222 {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	v239 = int32(0)
	goto L39
L59:
	;
	goto L60
L60:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v203)+76))
	if v215 < int32(0) {
		goto L54
	} else {
		goto L61
	}
L61:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v205)+76))
	if v215 == v218 {
		goto L54
	} else {
		goto L62
	}
L62:
	;
	v239 = int32(0)
	goto L39
L63:
	;
	v239 = int32(0)
	goto L39
L64:
	;
	goto L65
L65:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+83)))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+83)))
	if v225 == v226 {
		goto L54
	} else {
		goto L66
	}
L66:
	;
	v239 = int32(0)
	goto L39
L67:
	;
	goto L50
L68:
	;
	v267 = v102
	v278 = v113
	goto L30
L69:
	;
	F_SPI_freetuptable(m, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
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
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v284 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L71
L73:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+20))
	F_MemoryContextReset(m, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v289 = F_exec_stmts(m, l0, v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L4
	} else {
		goto L80
	}
L76:
	;
	goto L75
L77:
	;
	v368 = v110 + int64(1)
	if v368 != v94 {
		v102 = v267
		v110 = v368
		v113 = v278
		goto L28
	} else {
		goto L107
	}
L78:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v329 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L79:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v293 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	switch v289 - int32(1) {
	case 0:
		goto L79
	case 1:
		v383 = v82
		v384 = v289
		goto L12
	case 2:
		goto L78
	default:
		v364 = v289
		goto L77
	}
L81:
	;
	v383 = v82
	v384 = int32(0)
	goto L12
L82:
	;
	goto L83
L83:
	;
	v297 = int32(1)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v298 == int32(0) {
		v383 = v82
		v384 = v297
		goto L12
	} else {
		goto L84
	}
L84:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	if v304 == int32(0) {
		v323 = v303
		v324 = v304
		goto L86
	} else {
		goto L87
	}
L85:
	;
	if v324-v323 != 0 {
		v383 = v82
		v384 = v297
		goto L12
	} else {
		goto L93
	}
L86:
	;
	goto L85
L87:
	;
	if v303 != v304 {
		v323 = v303
		v324 = v304
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v308 = v298
	v309 = v293
	goto L89
L89:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+1)))
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308)+1)))
	if v313 == int32(0) {
		v323 = v312
		v324 = v313
		goto L86
	} else {
		goto L91
	}
L90:
	;
	v323 = v312
	v324 = v313
	goto L86
L91:
	;
	v316 = int32(1)
	if v312 == v313 {
		v308 = v308 + v316
		v309 = v309 + v316
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v326 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v326
	v383 = v82
	v384 = v326
	goto L12
L94:
	;
	v364 = int32(0)
	goto L77
L95:
	;
	goto L96
L96:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v333 == int32(0) {
		v383 = v82
		v384 = v289
		goto L12
	} else {
		goto L97
	}
L97:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329))))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
	if v339 == int32(0) {
		v358 = v338
		v359 = v339
		goto L99
	} else {
		goto L100
	}
L98:
	;
	if v359-v358 != 0 {
		v383 = v82
		v384 = v289
		goto L12
	} else {
		goto L106
	}
L99:
	;
	goto L98
L100:
	;
	if v338 != v339 {
		v358 = v338
		v359 = v339
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v343 = v333
	v344 = v329
	goto L102
L102:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344)+1)))
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+1)))
	if v348 == int32(0) {
		v358 = v347
		v359 = v348
		goto L99
	} else {
		goto L104
	}
L103:
	;
	v358 = v347
	v359 = v348
	goto L99
L104:
	;
	v351 = int32(1)
	if v347 == v348 {
		v343 = v343 + v351
		v344 = v344 + v351
		goto L102
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	v361 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v361
	v364 = v361
	goto L77
L107:
	;
	goto L29
L108:
	;
	F_SPI_cursor_fetch(m, l2, v76)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v375 = *(*int32)(unsafe.Add(mBase, _consts[501]))
	v377 = *(*int64)(unsafe.Add(mBase, _consts[500]))
	if v377 != int64(0) {
		v82 = v375
		v84 = v267
		v94 = v377
		v95 = v278
		goto L26
	} else {
		goto L110
	}
L110:
	;
	goto L27
L111:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+84)))
	if v399 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L4
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v415 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+84)) = uint8(v415)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v417+v418<<(uint(int32(2))%32))))
	F_assign_simple_var(m, l0, v422, base.B2i32(v55 != int64(0)), v415, v415)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L4
	} else {
		goto L118
	}
L115:
	;
	F_errmsg_internal(m, int32(445975), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(489522), int32(383), int32(305300))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	return v384
}
