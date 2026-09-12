package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_datumGetSize(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 != 0 {
		v119 = l2
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L42
	} else {
		goto L50
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L42
	} else {
		goto L46
	}
L3:
	;
	m.G0 = v7 + int32(16)
	return v119
L4:
	;
	if int32(0) < l2 {
		v119 = l2
		goto L3
	} else {
		goto L5
	}
L5:
	;
	switch l2 + int32(2) {
	case 0:
		goto L8
	case 1:
		goto L9
	default:
		goto L7
	}
L6:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v119 = int32(base.Ui32(v116) >> (uint(int32(2)) % 32))
	goto L3
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L42
	} else {
		goto L43
	}
L8:
	;
	if l0 == int32(0) {
		goto L1
	} else {
		goto L24
	}
L9:
	;
	if l0 == int32(0) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v15 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v18 = int32(6)
	v20 = int32(18)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v22 == v20 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	if v15&int32(1) == int32(0) {
		goto L6
	} else {
		goto L23
	}
L14:
	;
	v25 = v20
	goto L16
L15:
	;
	v25 = int32(2)
	goto L16
L16:
	;
	if v22&int32(254) == int32(2) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v30 = v18
	goto L19
L18:
	;
	v30 = v25
	goto L19
L19:
	;
	if v22 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v33 = v18
	goto L22
L21:
	;
	v33 = v30
	goto L22
L22:
	;
	v119 = v33
	goto L3
L23:
	;
	v119 = int32(base.Ui32(v15) >> (uint(int32(1)) % 32))
	goto L3
L24:
	;
	if l0&int32(3) == int32(0) {
		v65 = l0
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v119 = v98 + int32(1)
	goto L3
L26:
	;
	v98 = v90 - l0
	goto L25
L27:
	;
	v69 = v65
	goto L36
L28:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v49 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v98 = int32(0)
	goto L25
L30:
	;
	goto L31
L31:
	;
	v54 = l0
	goto L32
L32:
	;
	v58 = v54 + int32(1)
	if v58&int32(3) == int32(0) {
		v65 = v58
		goto L27
	} else {
		goto L34
	}
L33:
	;
	v90 = v58
	goto L26
L34:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v63 != 0 {
		v54 = v58
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v78 = int32(-2139062144)
	if (int32(16843008)-v75|v75)&v78 == v78 {
		v69 = v69 + int32(4)
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v84 = v69
	goto L39
L38:
	;
	goto L37
L39:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v88 != 0 {
		v84 = v84 + int32(1)
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v90 = v84
	goto L26
L41:
	;
	goto L40
L42:
	;
	return int32(0)
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l2
	F_errmsg_internal(m, int32(460871), v7)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(474028), int32(108), int32(325570))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L42
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(204922), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L42
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(474028), int32(90), int32(325570))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L42
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
	F_errcode(m, int32(130))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L42
	} else {
		goto L51
	}
L51:
	;
	F_errmsg(m, int32(204922), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L42
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(474028), int32(102), int32(325570))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L42
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_datumTransfer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	if l1 != 0 {
		v20 = F_datumCopy(m, l0, l1, l2)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v20
		}
	} else {
		if l2 != int32(-1) {
			v20 = F_datumCopy(m, l0, l1, l2)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				return v20
			}
		} else {
			v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v6 != int32(1) {
				v20 = F_datumCopy(m, l0, l1, l2)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					return v20
				}
			} else {
				v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				if v9 != int32(3) {
					v20 = F_datumCopy(m, l0, l1, l2)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						return v20
					}
				} else {
					v13 = *(*int32)(unsafe.Add(mBase, _consts[28]))
					v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
					v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
					F_MemoryContextSetParent(m, v15, v13)
					mBase = m.M
					return v14 + int32(12)
				}
			}
		}
	}
}
func F_datum_to_json_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	v8 = m.G0
	v10 = v8 - int32(176)
	m.G0 = v10
	F_check_stack_depth(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if l1 != 0 {
			F_appendBinaryStringInfo(m, l2, int32(288610), int32(4))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				m.G0 = v10 + int32(176)
				return
			}
		} else {
			if l5 != 0 {
				switch l3 - int32(1) {
				case 0:
					F_appendStringInfoChar(m, l2, int32(34))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						if l0 != 0 {
							F_appendBinaryStringInfo(m, l2, int32(327532), int32(4))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								if l5 == int32(0) {
									m.G0 = v10 + int32(176)
									return
								} else {
									F_appendStringInfoChar(m, l2, int32(34))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										m.G0 = v10 + int32(176)
										return
									}
								}
							}
						} else {
							F_appendBinaryStringInfo(m, l2, int32(344107), int32(5))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								if l5 == int32(0) {
									m.G0 = v10 + int32(176)
									return
								} else {
									F_appendStringInfoChar(m, l2, int32(34))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										m.G0 = v10 + int32(176)
										return
									}
								}
							}
						}
					}
				case 1:
					v84 = F_OidOutputFunctionCall(m, l4, l0)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						v87 = v84
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v87)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return
								} else {
									F_pfree(m, v87)
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return
									} else {
										m.G0 = v10 + int32(176)
										return
									}
								}
							}
						}
					}
				case 2:
					if base.Ui32(l0-int32(2147483647)) <= base.Ui32(int32(1)) {
						F_EncodeSpecialDate(m, l0, v10)
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return
						} else {
							F_appendStringInfoChar(m, l2, int32(34))
							mBase = m.M
							v314 = m.ExcPending
							if v314 != 0 {
								return
							} else {
								F_appendStringInfoString(m, l2, v10)
								mBase = m.M
								v316 = m.ExcPending
								if v316 != 0 {
									return
								} else {
									F_appendStringInfoChar(m, l2, int32(34))
									mBase = m.M
									v319 = m.ExcPending
									if v319 != 0 {
										return
									} else {
										m.G0 = v10 + int32(176)
										return
									}
								}
							}
						}
					} else {
						v115 = l0 + int32(2483589)
						v116 = int32(146097)
						v117 = base.I32_div_u_s(v115, v116)
						v118 = int32(3)
						v124 = int32(2)
						v129 = base.I32_div_u_s((v117*int32(1073595727)+v115)<<(uint(v124)%32)|v118, v116)
						v132 = l0 + int32(2451545) + v117*v118 + v129 + int32(32104)
						v133 = int32(1461)
						v134 = base.I32_div_u_s(v132, v133)
						v137 = v134*int32(-1461) + v132
						v139 = v137 << (uint(v124) % 32)
						if base.Ui32(v133) <= base.Ui32(v139) {
							v145 = base.I32_rem_u_s(v137+int32(305), int32(365))
							v150 = v145
						} else {
							v149 = base.I32_rem_u_s(v137+int32(306), int32(366))
							v150 = v149
						}
						v152 = base.I32_div_u_s(v139, int32(1461))
						*(*int32)(unsafe.Add(mBase, uint32(v10+int32(152)))) = v152 + v134<<(uint(int32(2))%32) - int32(4800)
						v160 = v150 + int32(123)
						v164 = int32(base.Ui32(v160*int32(2141)) >> (uint(int32(16)) % 32))
						*(*int32)(unsafe.Add(mBase, uint32(v10+int32(144)))) = v160 - int32(base.Ui32(v164*int32(7834))>>(uint(int32(8))%32))
						v174 = base.I32_rem_u_s(v164+int32(10), int32(12))
						*(*int32)(unsafe.Add(mBase, uint32(v10+int32(148)))) = v174 + int32(1)
						v179 = v10 + int32(132)
						switch int32(3) {
						case 0, 3:
							v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
							if int32(0) < v183 {
								v188 = v183
							} else {
								v188 = int32(1) - v183
							}
							v190 = F_pg_ultostr_zeropad(m, v10, v188, int32(4))
							mBase = m.M
							v191 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v190))) = uint8(v191)
							v193 = int32(1)
							v195 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
							v196 = int32(2)
							v197 = F_pg_ultostr_zeropad(m, v190+v193, v195, v196)
							mBase = m.M
							*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v191)
							v202 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
							v204 = F_pg_ultostr_zeropad(m, v197+v193, v202, v196)
							mBase = m.M
							v297 = v204
						case 1:
							v208 = *(*int32)(unsafe.Add(mBase, _consts[467]))
							v210 = base.B2i32(v208 == int32(1))
							if v208 == int32(1) {
								v211 = int32(12)
							} else {
								v211 = int32(16)
							}
							v213 = *(*int32)(unsafe.Add(mBase, uint32(v179+v211)))
							v215 = F_pg_ultostr_zeropad(m, v10, v213, int32(2))
							mBase = m.M
							v216 = int32(47)
							*(*uint8)(unsafe.Add(mBase, uint32(v215))) = uint8(v216)
							if v208 == int32(1) {
								v222 = int32(16)
							} else {
								v222 = int32(12)
							}
							v224 = *(*int32)(unsafe.Add(mBase, uint32(v179+v222)))
							v226 = F_pg_ultostr_zeropad(m, v215+int32(1), v224, int32(2))
							mBase = m.M
							v227 = int32(47)
							*(*uint8)(unsafe.Add(mBase, uint32(v226))) = uint8(v227)
							v229 = int32(1)
							v231 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
							if int32(0) < v231 {
								v236 = v231
							} else {
								v236 = v229 - v231
							}
							v238 = F_pg_ultostr_zeropad(m, v226+v229, v236, int32(4))
							mBase = m.M
							v297 = v238
						case 2:
							v239 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
							v240 = int32(2)
							v241 = F_pg_ultostr_zeropad(m, v10, v239, v240)
							mBase = m.M
							v242 = int32(46)
							*(*uint8)(unsafe.Add(mBase, uint32(v241))) = uint8(v242)
							v244 = int32(1)
							v246 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
							v248 = F_pg_ultostr_zeropad(m, v241+v244, v246, v240)
							mBase = m.M
							*(*uint8)(unsafe.Add(mBase, uint32(v248))) = uint8(v242)
							v253 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
							if int32(0) < v253 {
								v258 = v253
							} else {
								v258 = v244 - v253
							}
							v260 = F_pg_ultostr_zeropad(m, v248+v244, v258, int32(4))
							mBase = m.M
							v297 = v260
						default:
							v264 = *(*int32)(unsafe.Add(mBase, _consts[467]))
							v266 = base.B2i32(v264 == int32(1))
							if v264 == int32(1) {
								v267 = int32(12)
							} else {
								v267 = int32(16)
							}
							v269 = *(*int32)(unsafe.Add(mBase, uint32(v179+v267)))
							v271 = F_pg_ultostr_zeropad(m, v10, v269, int32(2))
							mBase = m.M
							v272 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v271))) = uint8(v272)
							if v264 == int32(1) {
								v278 = int32(16)
							} else {
								v278 = int32(12)
							}
							v280 = *(*int32)(unsafe.Add(mBase, uint32(v179+v278)))
							v282 = F_pg_ultostr_zeropad(m, v271+int32(1), v280, int32(2))
							mBase = m.M
							v283 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v282))) = uint8(v283)
							v285 = int32(1)
							v287 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
							if int32(0) < v287 {
								v292 = v287
							} else {
								v292 = v285 - v287
							}
							v294 = F_pg_ultostr_zeropad(m, v282+v285, v292, int32(4))
							mBase = m.M
							v297 = v294
						}
						v298 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
						if v298 <= int32(0) {
							v302 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1291])))
							*(*uint8)(unsafe.Add(mBase, uint32(v297)+2)) = uint8(v302)
							v305 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1292])))
							*(*uint16)(unsafe.Add(mBase, uint32(v297))) = uint16(v305)
							v309 = v297 + int32(3)
						} else {
							v309 = v297
						}
						v310 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v309))) = uint8(v310)
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v314 = m.ExcPending
						if v314 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v10)
							mBase = m.M
							v316 = m.ExcPending
							if v316 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v319 = m.ExcPending
								if v319 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					}
				case 3:
					v322 = F_JsonEncodeDateTime(m, v10, l0, int32(1114), int32(0))
					mBase = m.M
					v323 = m.ExcPending
					if v323 != 0 {
						return
					} else {
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v326 = m.ExcPending
						if v326 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v10)
							mBase = m.M
							v328 = m.ExcPending
							if v328 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v331 = m.ExcPending
								if v331 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					}
				case 4:
					v334 = F_JsonEncodeDateTime(m, v10, l0, int32(1184), int32(0))
					mBase = m.M
					v335 = m.ExcPending
					if v335 != 0 {
						return
					} else {
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v338 = m.ExcPending
						if v338 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v10)
							mBase = m.M
							v340 = m.ExcPending
							if v340 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v343 = m.ExcPending
								if v343 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					}
				case 5, 7, 8, 9:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							F_errmsg(m, int32(233004), int32(0))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								F_errfinish(m, int32(473239), int32(204), int32(295993))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				default:
					switch l4 - int32(1045) {
					case 0, 2:
						v399 = F_pg_detoast_datum_packed(m, l0)
						mBase = m.M
						v400 = m.ExcPending
						if v400 != 0 {
							return
						} else {
							v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399))))
							if v401 == int32(1) {
								v404 = int32(4)
								v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399)+1)))
								if v406&int32(254) == int32(2) {
									v415 = v404
								} else {
									v415 = base.B2i32(v406 == int32(18)) << (uint(v404) % 32)
								}
								if v406 == int32(1) {
									v418 = v404
								} else {
									v418 = v415
								}
								v431 = v418
							} else {
								v419 = int32(1)
								if v401&v419 != 0 {
									v431 = int32(base.Ui32(v401)>>(uint(v419)%32)) - v419
								} else {
									v425 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
									v431 = int32(base.Ui32(v425)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							v432 = int32(1)
							if v401&v432 != 0 {
								v436 = v432
							} else {
								v436 = int32(4)
							}
							F_escape_json_with_len(m, l2, v399+v436, v431)
							mBase = m.M
							v439 = m.ExcPending
							if v439 != 0 {
								return
							} else {
								if l0 == v399 {
									m.G0 = v10 + int32(176)
									return
								} else {
									F_pfree(m, v399)
									mBase = m.M
									v442 = m.ExcPending
									if v442 != 0 {
										return
									} else {
										m.G0 = v10 + int32(176)
										return
									}
								}
							}
						}
					case 1:
						v443 = F_OidOutputFunctionCall(m, l4, l0)
						mBase = m.M
						v444 = m.ExcPending
						if v444 != 0 {
							return
						} else {
							F_escape_json(m, l2, v443)
							mBase = m.M
							v446 = m.ExcPending
							if v446 != 0 {
								return
							} else {
								F_pfree(m, v443)
								mBase = m.M
								v448 = m.ExcPending
								if v448 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					default:
						if l4 != int32(47) {
							v443 = F_OidOutputFunctionCall(m, l4, l0)
							mBase = m.M
							v444 = m.ExcPending
							if v444 != 0 {
								return
							} else {
								F_escape_json(m, l2, v443)
								mBase = m.M
								v446 = m.ExcPending
								if v446 != 0 {
									return
								} else {
									F_pfree(m, v443)
									mBase = m.M
									v448 = m.ExcPending
									if v448 != 0 {
										return
									} else {
										m.G0 = v10 + int32(176)
										return
									}
								}
							}
						} else {
							v399 = F_pg_detoast_datum_packed(m, l0)
							mBase = m.M
							v400 = m.ExcPending
							if v400 != 0 {
								return
							} else {
								v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399))))
								if v401 == int32(1) {
									v404 = int32(4)
									v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399)+1)))
									if v406&int32(254) == int32(2) {
										v415 = v404
									} else {
										v415 = base.B2i32(v406 == int32(18)) << (uint(v404) % 32)
									}
									if v406 == int32(1) {
										v418 = v404
									} else {
										v418 = v415
									}
									v431 = v418
								} else {
									v419 = int32(1)
									if v401&v419 != 0 {
										v431 = int32(base.Ui32(v401)>>(uint(v419)%32)) - v419
									} else {
										v425 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
										v431 = int32(base.Ui32(v425)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								v432 = int32(1)
								if v401&v432 != 0 {
									v436 = v432
								} else {
									v436 = int32(4)
								}
								F_escape_json_with_len(m, l2, v399+v436, v431)
								mBase = m.M
								v439 = m.ExcPending
								if v439 != 0 {
									return
								} else {
									if l0 == v399 {
										m.G0 = v10 + int32(176)
										return
									} else {
										F_pfree(m, v399)
										mBase = m.M
										v442 = m.ExcPending
										if v442 != 0 {
											return
										} else {
											m.G0 = v10 + int32(176)
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				switch l3 - int32(1) {
				case 0:
					if l0 != 0 {
						F_appendBinaryStringInfo(m, l2, int32(327532), int32(4))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							if l5 == int32(0) {
								m.G0 = v10 + int32(176)
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					} else {
						F_appendBinaryStringInfo(m, l2, int32(344107), int32(5))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							if l5 == int32(0) {
								m.G0 = v10 + int32(176)
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					}
				case 1:
					v60 = F_OidOutputFunctionCall(m, l4, l0)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
						if base.Ui32(int32(10)) <= base.Ui32((v62-int32(48))&int32(255)) {
							if v62&int32(255) != int32(45) {
								v87 = v60
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return
								} else {
									F_appendStringInfoString(m, l2, v87)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return
									} else {
										F_appendStringInfoChar(m, l2, int32(34))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											F_pfree(m, v87)
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
											} else {
												m.G0 = v10 + int32(176)
												return
											}
										}
									}
								}
							} else {
								v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
								if base.Ui32(int32(9)) < base.Ui32((v73-int32(48))&int32(255)) {
									v87 = v60
									F_appendStringInfoChar(m, l2, int32(34))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										F_appendStringInfoString(m, l2, v87)
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return
										} else {
											F_appendStringInfoChar(m, l2, int32(34))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												F_pfree(m, v87)
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return
												} else {
													m.G0 = v10 + int32(176)
													return
												}
											}
										}
									}
								} else {
									F_appendStringInfoString(m, l2, v60)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										F_pfree(m, v60)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											m.G0 = v10 + int32(176)
											return
										}
									}
								}
							}
						} else {
							F_appendStringInfoString(m, l2, v60)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								F_pfree(m, v60)
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					}
				case 2:
					if base.Ui32(l0-int32(2147483647)) <= base.Ui32(int32(1)) {
						F_EncodeSpecialDate(m, l0, v10)
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return
						} else {
							F_appendStringInfoChar(m, l2, int32(34))
							mBase = m.M
							v314 = m.ExcPending
							if v314 != 0 {
								return
							} else {
								F_appendStringInfoString(m, l2, v10)
								mBase = m.M
								v316 = m.ExcPending
								if v316 != 0 {
									return
								} else {
									F_appendStringInfoChar(m, l2, int32(34))
									mBase = m.M
									v319 = m.ExcPending
									if v319 != 0 {
										return
									} else {
										m.G0 = v10 + int32(176)
										return
									}
								}
							}
						}
					} else {
						v115 = l0 + int32(2483589)
						v116 = int32(146097)
						v117 = base.I32_div_u_s(v115, v116)
						v118 = int32(3)
						v124 = int32(2)
						v129 = base.I32_div_u_s((v117*int32(1073595727)+v115)<<(uint(v124)%32)|v118, v116)
						v132 = l0 + int32(2451545) + v117*v118 + v129 + int32(32104)
						v133 = int32(1461)
						v134 = base.I32_div_u_s(v132, v133)
						v137 = v134*int32(-1461) + v132
						v139 = v137 << (uint(v124) % 32)
						if base.Ui32(v133) <= base.Ui32(v139) {
							v145 = base.I32_rem_u_s(v137+int32(305), int32(365))
							v150 = v145
						} else {
							v149 = base.I32_rem_u_s(v137+int32(306), int32(366))
							v150 = v149
						}
						v152 = base.I32_div_u_s(v139, int32(1461))
						*(*int32)(unsafe.Add(mBase, uint32(v10+int32(152)))) = v152 + v134<<(uint(int32(2))%32) - int32(4800)
						v160 = v150 + int32(123)
						v164 = int32(base.Ui32(v160*int32(2141)) >> (uint(int32(16)) % 32))
						*(*int32)(unsafe.Add(mBase, uint32(v10+int32(144)))) = v160 - int32(base.Ui32(v164*int32(7834))>>(uint(int32(8))%32))
						v174 = base.I32_rem_u_s(v164+int32(10), int32(12))
						*(*int32)(unsafe.Add(mBase, uint32(v10+int32(148)))) = v174 + int32(1)
						v179 = v10 + int32(132)
						switch int32(3) {
						case 0, 3:
							v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
							if int32(0) < v183 {
								v188 = v183
							} else {
								v188 = int32(1) - v183
							}
							v190 = F_pg_ultostr_zeropad(m, v10, v188, int32(4))
							mBase = m.M
							v191 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v190))) = uint8(v191)
							v193 = int32(1)
							v195 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
							v196 = int32(2)
							v197 = F_pg_ultostr_zeropad(m, v190+v193, v195, v196)
							mBase = m.M
							*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v191)
							v202 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
							v204 = F_pg_ultostr_zeropad(m, v197+v193, v202, v196)
							mBase = m.M
							v297 = v204
						case 1:
							v208 = *(*int32)(unsafe.Add(mBase, _consts[467]))
							v210 = base.B2i32(v208 == int32(1))
							if v208 == int32(1) {
								v211 = int32(12)
							} else {
								v211 = int32(16)
							}
							v213 = *(*int32)(unsafe.Add(mBase, uint32(v179+v211)))
							v215 = F_pg_ultostr_zeropad(m, v10, v213, int32(2))
							mBase = m.M
							v216 = int32(47)
							*(*uint8)(unsafe.Add(mBase, uint32(v215))) = uint8(v216)
							if v208 == int32(1) {
								v222 = int32(16)
							} else {
								v222 = int32(12)
							}
							v224 = *(*int32)(unsafe.Add(mBase, uint32(v179+v222)))
							v226 = F_pg_ultostr_zeropad(m, v215+int32(1), v224, int32(2))
							mBase = m.M
							v227 = int32(47)
							*(*uint8)(unsafe.Add(mBase, uint32(v226))) = uint8(v227)
							v229 = int32(1)
							v231 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
							if int32(0) < v231 {
								v236 = v231
							} else {
								v236 = v229 - v231
							}
							v238 = F_pg_ultostr_zeropad(m, v226+v229, v236, int32(4))
							mBase = m.M
							v297 = v238
						case 2:
							v239 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
							v240 = int32(2)
							v241 = F_pg_ultostr_zeropad(m, v10, v239, v240)
							mBase = m.M
							v242 = int32(46)
							*(*uint8)(unsafe.Add(mBase, uint32(v241))) = uint8(v242)
							v244 = int32(1)
							v246 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
							v248 = F_pg_ultostr_zeropad(m, v241+v244, v246, v240)
							mBase = m.M
							*(*uint8)(unsafe.Add(mBase, uint32(v248))) = uint8(v242)
							v253 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
							if int32(0) < v253 {
								v258 = v253
							} else {
								v258 = v244 - v253
							}
							v260 = F_pg_ultostr_zeropad(m, v248+v244, v258, int32(4))
							mBase = m.M
							v297 = v260
						default:
							v264 = *(*int32)(unsafe.Add(mBase, _consts[467]))
							v266 = base.B2i32(v264 == int32(1))
							if v264 == int32(1) {
								v267 = int32(12)
							} else {
								v267 = int32(16)
							}
							v269 = *(*int32)(unsafe.Add(mBase, uint32(v179+v267)))
							v271 = F_pg_ultostr_zeropad(m, v10, v269, int32(2))
							mBase = m.M
							v272 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v271))) = uint8(v272)
							if v264 == int32(1) {
								v278 = int32(16)
							} else {
								v278 = int32(12)
							}
							v280 = *(*int32)(unsafe.Add(mBase, uint32(v179+v278)))
							v282 = F_pg_ultostr_zeropad(m, v271+int32(1), v280, int32(2))
							mBase = m.M
							v283 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v282))) = uint8(v283)
							v285 = int32(1)
							v287 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
							if int32(0) < v287 {
								v292 = v287
							} else {
								v292 = v285 - v287
							}
							v294 = F_pg_ultostr_zeropad(m, v282+v285, v292, int32(4))
							mBase = m.M
							v297 = v294
						}
						v298 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
						if v298 <= int32(0) {
							v302 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1291])))
							*(*uint8)(unsafe.Add(mBase, uint32(v297)+2)) = uint8(v302)
							v305 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1292])))
							*(*uint16)(unsafe.Add(mBase, uint32(v297))) = uint16(v305)
							v309 = v297 + int32(3)
						} else {
							v309 = v297
						}
						v310 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v309))) = uint8(v310)
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v314 = m.ExcPending
						if v314 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v10)
							mBase = m.M
							v316 = m.ExcPending
							if v316 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v319 = m.ExcPending
								if v319 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					}
				case 3:
					v322 = F_JsonEncodeDateTime(m, v10, l0, int32(1114), int32(0))
					mBase = m.M
					v323 = m.ExcPending
					if v323 != 0 {
						return
					} else {
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v326 = m.ExcPending
						if v326 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v10)
							mBase = m.M
							v328 = m.ExcPending
							if v328 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v331 = m.ExcPending
								if v331 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					}
				case 4:
					v334 = F_JsonEncodeDateTime(m, v10, l0, int32(1184), int32(0))
					mBase = m.M
					v335 = m.ExcPending
					if v335 != 0 {
						return
					} else {
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v338 = m.ExcPending
						if v338 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v10)
							mBase = m.M
							v340 = m.ExcPending
							if v340 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v343 = m.ExcPending
								if v343 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					}
				case 5:
					v344 = F_OidOutputFunctionCall(m, l4, l0)
					mBase = m.M
					v345 = m.ExcPending
					if v345 != 0 {
						return
					} else {
						F_appendStringInfoString(m, l2, v344)
						mBase = m.M
						v347 = m.ExcPending
						if v347 != 0 {
							return
						} else {
							F_pfree(m, v344)
							mBase = m.M
							v349 = m.ExcPending
							if v349 != 0 {
								return
							} else {
								m.G0 = v10 + int32(176)
								return
							}
						}
					}
				default:
					switch l4 - int32(1045) {
					case 0, 2:
						v399 = F_pg_detoast_datum_packed(m, l0)
						mBase = m.M
						v400 = m.ExcPending
						if v400 != 0 {
							return
						} else {
							v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399))))
							if v401 == int32(1) {
								v404 = int32(4)
								v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399)+1)))
								if v406&int32(254) == int32(2) {
									v415 = v404
								} else {
									v415 = base.B2i32(v406 == int32(18)) << (uint(v404) % 32)
								}
								if v406 == int32(1) {
									v418 = v404
								} else {
									v418 = v415
								}
								v431 = v418
							} else {
								v419 = int32(1)
								if v401&v419 != 0 {
									v431 = int32(base.Ui32(v401)>>(uint(v419)%32)) - v419
								} else {
									v425 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
									v431 = int32(base.Ui32(v425)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							v432 = int32(1)
							if v401&v432 != 0 {
								v436 = v432
							} else {
								v436 = int32(4)
							}
							F_escape_json_with_len(m, l2, v399+v436, v431)
							mBase = m.M
							v439 = m.ExcPending
							if v439 != 0 {
								return
							} else {
								if l0 == v399 {
									m.G0 = v10 + int32(176)
									return
								} else {
									F_pfree(m, v399)
									mBase = m.M
									v442 = m.ExcPending
									if v442 != 0 {
										return
									} else {
										m.G0 = v10 + int32(176)
										return
									}
								}
							}
						}
					case 1:
						v443 = F_OidOutputFunctionCall(m, l4, l0)
						mBase = m.M
						v444 = m.ExcPending
						if v444 != 0 {
							return
						} else {
							F_escape_json(m, l2, v443)
							mBase = m.M
							v446 = m.ExcPending
							if v446 != 0 {
								return
							} else {
								F_pfree(m, v443)
								mBase = m.M
								v448 = m.ExcPending
								if v448 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					default:
						if l4 != int32(47) {
							v443 = F_OidOutputFunctionCall(m, l4, l0)
							mBase = m.M
							v444 = m.ExcPending
							if v444 != 0 {
								return
							} else {
								F_escape_json(m, l2, v443)
								mBase = m.M
								v446 = m.ExcPending
								if v446 != 0 {
									return
								} else {
									F_pfree(m, v443)
									mBase = m.M
									v448 = m.ExcPending
									if v448 != 0 {
										return
									} else {
										m.G0 = v10 + int32(176)
										return
									}
								}
							}
						} else {
							v399 = F_pg_detoast_datum_packed(m, l0)
							mBase = m.M
							v400 = m.ExcPending
							if v400 != 0 {
								return
							} else {
								v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399))))
								if v401 == int32(1) {
									v404 = int32(4)
									v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399)+1)))
									if v406&int32(254) == int32(2) {
										v415 = v404
									} else {
										v415 = base.B2i32(v406 == int32(18)) << (uint(v404) % 32)
									}
									if v406 == int32(1) {
										v418 = v404
									} else {
										v418 = v415
									}
									v431 = v418
								} else {
									v419 = int32(1)
									if v401&v419 != 0 {
										v431 = int32(base.Ui32(v401)>>(uint(v419)%32)) - v419
									} else {
										v425 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
										v431 = int32(base.Ui32(v425)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								v432 = int32(1)
								if v401&v432 != 0 {
									v436 = v432
								} else {
									v436 = int32(4)
								}
								F_escape_json_with_len(m, l2, v399+v436, v431)
								mBase = m.M
								v439 = m.ExcPending
								if v439 != 0 {
									return
								} else {
									if l0 == v399 {
										m.G0 = v10 + int32(176)
										return
									} else {
										F_pfree(m, v399)
										mBase = m.M
										v442 = m.ExcPending
										if v442 != 0 {
											return
										} else {
											m.G0 = v10 + int32(176)
											return
										}
									}
								}
							}
						}
					}
				case 7:
					F_array_to_json_internal(m, l0, l2, int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						m.G0 = v10 + int32(176)
						return
					}
				case 8:
					F_composite_to_json(m, l0, l2, int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						m.G0 = v10 + int32(176)
						return
					}
				case 9:
					v351 = F_OidFunctionCall1Coll(m, l4, int32(0), l0)
					mBase = m.M
					v352 = m.ExcPending
					if v352 != 0 {
						return
					} else {
						v353 = F_pg_detoast_datum_packed(m, v351)
						mBase = m.M
						v354 = m.ExcPending
						if v354 != 0 {
							return
						} else {
							v355 = int32(1)
							v356 = v353 + v355
							v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353))))
							v361 = v359 & v355
							if v361 != 0 {
								v362 = v356
							} else {
								v362 = v353 + int32(4)
							}
							if v359 == int32(1) {
								v365 = int32(4)
								v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356))))
								if v367&int32(254) == int32(2) {
									v376 = v365
								} else {
									v376 = base.B2i32(v367 == int32(18)) << (uint(v365) % 32)
								}
								if v367 == int32(1) {
									v379 = v365
								} else {
									v379 = v376
								}
								v390 = v379
							} else {
								v380 = int32(1)
								if v361 != 0 {
									v390 = int32(base.Ui32(v359)>>(uint(v380)%32)) - v380
								} else {
									v384 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
									v390 = int32(base.Ui32(v384)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							F_appendBinaryStringInfo(m, l2, v362, v390)
							mBase = m.M
							v392 = m.ExcPending
							if v392 != 0 {
								return
							} else {
								F_pfree(m, v353)
								mBase = m.M
								v394 = m.ExcPending
								if v394 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
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
func F_exec_eval_datum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
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
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v14 {
	case 0:
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v20
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		*(*int32)(unsafe.Add(mBase, uint32(l4))) = v22
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
		*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v24)
		m.G0 = v12 + int32(32)
		return
	case 1:
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		if v26 == int32(0) {
			F_errstart_cold(m, int32(21), int32(525467))
			mBase = m.M
			v171 = m.ExcPending
			if v171 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(466067), int32(0))
				mBase = m.M
				v176 = m.ExcPending
				if v176 != 0 {
					return
				} else {
					F_errfinish(m, int32(476889), int32(5323), int32(272333))
					mBase = m.M
					v183 = m.ExcPending
					if v183 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = F_BlessTupleDesc(m, v26)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				v31 = int32(4442992)
				v32 = *(*int32)(unsafe.Add(mBase, _consts[28]))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
				*(*int32)(unsafe.Add(mBase, _consts[28])) = v34
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
				v37 = F_make_tuple_from_row(m, l0, l1, v36)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					if v37 == int32(0) {
						F_errstart_cold(m, int32(21), int32(525467))
						mBase = m.M
						v189 = m.ExcPending
						if v189 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(466095), int32(0))
							mBase = m.M
							v194 = m.ExcPending
							if v194 != 0 {
								return
							} else {
								F_errfinish(m, int32(476889), int32(5329), int32(272333))
								mBase = m.M
								v201 = m.ExcPending
								if v201 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v42
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v44
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
						v47 = F_HeapTupleHeaderGetDatum(m, v46)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v47
							v50 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v50)
							*(*int32)(unsafe.Add(mBase, _consts[28])) = v32
							m.G0 = v12 + int32(32)
							return
						}
					}
				}
			}
		}
	case 2:
		v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
		if v54 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
			v59 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v59)
			v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v61
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(-1)
		} else {
			v68 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
			v70 = v68 & int32(5)
			if v70 != 0 {
				v71 = v54 + int32(12)
			} else {
				v71 = int32(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v71
			*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(base.B2i32(v70 == int32(0)))
			v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			if v76 == int32(2249) {
				v155 = *(*int32)(unsafe.Add(mBase, uint32(v54)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v155
				v157 = *(*int32)(unsafe.Add(mBase, uint32(v54)+40))
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v157
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v76
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(-1)
			}
		}
		m.G0 = v12 + int32(32)
		return
	case 3:
		v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v87 = *(*int32)(unsafe.Add(mBase, uint32(v82+v83<<(uint(int32(2))%32))))
		v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+36))
		if v88 == int32(0) {
			F_instantiate_empty_record_variable(m, l0, v87)
			mBase = m.M
			v92 = m.ExcPending
			if v92 != 0 {
				return
			} else {
				v93 = *(*int32)(unsafe.Add(mBase, uint32(v87)+36))
				v94 = v93
				v95 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
				v96 = *(*int64)(unsafe.Add(mBase, uint32(v94)+48))
				if v95 != v96 {
					v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v101 = F_expanded_record_lookup_field(m, v94, v98, l1+int32(32))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return
					} else {
						if v101 == int32(0) {
							F_errstart_cold(m, int32(21), int32(525467))
							mBase = m.M
							v207 = m.ExcPending
							if v207 != 0 {
								return
							} else {
								F_errcode(m, int32(50360452))
								mBase = m.M
								v210 = m.ExcPending
								if v210 != 0 {
									return
								} else {
									v211 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
									v212 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v212
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v211
									F_errmsg(m, int32(681354), v12+int32(16))
									mBase = m.M
									v220 = m.ExcPending
									if v220 != 0 {
										return
									} else {
										F_errfinish(m, int32(476889), int32(5412), int32(272333))
										mBase = m.M
										v227 = m.ExcPending
										if v227 != 0 {
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
							v105 = *(*int64)(unsafe.Add(mBase, uint32(v94)+48))
							*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v105
							v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v107
							v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v109
							v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
							if v111 <= int32(0) {
								v133 = F_expanded_record_fetch_field(m, v94, v111, l5)
								mBase = m.M
								v134 = m.ExcPending
								if v134 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = v133
									m.G0 = v12 + int32(32)
									return
								}
							} else {
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v94)+28))
								if v114&int32(4) == int32(0) {
									v133 = F_expanded_record_fetch_field(m, v94, v111, l5)
									mBase = m.M
									v134 = m.ExcPending
									if v134 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l4))) = v133
										m.G0 = v12 + int32(32)
										return
									}
								} else {
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v94)+64))
									if v119 < v111 {
										v133 = F_expanded_record_fetch_field(m, v94, v111, l5)
										mBase = m.M
										v134 = m.ExcPending
										if v134 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l4))) = v133
											m.G0 = v12 + int32(32)
											return
										}
									} else {
										v122 = v111 - int32(1)
										v123 = *(*int32)(unsafe.Add(mBase, uint32(v94)+60))
										v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v123))))
										*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v125)
										v127 = *(*int32)(unsafe.Add(mBase, uint32(v94)+56))
										v131 = *(*int32)(unsafe.Add(mBase, uint32(v127+v122<<(uint(int32(2))%32))))
										*(*int32)(unsafe.Add(mBase, uint32(l4))) = v131
										m.G0 = v12 + int32(32)
										return
									}
								}
							}
						}
					}
				} else {
					v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v107
					v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v109
					v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
					if v111 <= int32(0) {
						v133 = F_expanded_record_fetch_field(m, v94, v111, l5)
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v133
							m.G0 = v12 + int32(32)
							return
						}
					} else {
						v114 = *(*int32)(unsafe.Add(mBase, uint32(v94)+28))
						if v114&int32(4) == int32(0) {
							v133 = F_expanded_record_fetch_field(m, v94, v111, l5)
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v133
								m.G0 = v12 + int32(32)
								return
							}
						} else {
							v119 = *(*int32)(unsafe.Add(mBase, uint32(v94)+64))
							if v119 < v111 {
								v133 = F_expanded_record_fetch_field(m, v94, v111, l5)
								mBase = m.M
								v134 = m.ExcPending
								if v134 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = v133
									m.G0 = v12 + int32(32)
									return
								}
							} else {
								v122 = v111 - int32(1)
								v123 = *(*int32)(unsafe.Add(mBase, uint32(v94)+60))
								v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v123))))
								*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v125)
								v127 = *(*int32)(unsafe.Add(mBase, uint32(v94)+56))
								v131 = *(*int32)(unsafe.Add(mBase, uint32(v127+v122<<(uint(int32(2))%32))))
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v131
								m.G0 = v12 + int32(32)
								return
							}
						}
					}
				}
			}
		} else {
			v94 = v88
			v95 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
			v96 = *(*int64)(unsafe.Add(mBase, uint32(v94)+48))
			if v95 != v96 {
				v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v101 = F_expanded_record_lookup_field(m, v94, v98, l1+int32(32))
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return
				} else {
					if v101 == int32(0) {
						F_errstart_cold(m, int32(21), int32(525467))
						mBase = m.M
						v207 = m.ExcPending
						if v207 != 0 {
							return
						} else {
							F_errcode(m, int32(50360452))
							mBase = m.M
							v210 = m.ExcPending
							if v210 != 0 {
								return
							} else {
								v211 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
								v212 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v212
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v211
								F_errmsg(m, int32(681354), v12+int32(16))
								mBase = m.M
								v220 = m.ExcPending
								if v220 != 0 {
									return
								} else {
									F_errfinish(m, int32(476889), int32(5412), int32(272333))
									mBase = m.M
									v227 = m.ExcPending
									if v227 != 0 {
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
						v105 = *(*int64)(unsafe.Add(mBase, uint32(v94)+48))
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v105
						v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v107
						v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v109
						v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						if v111 <= int32(0) {
							v133 = F_expanded_record_fetch_field(m, v94, v111, l5)
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v133
								m.G0 = v12 + int32(32)
								return
							}
						} else {
							v114 = *(*int32)(unsafe.Add(mBase, uint32(v94)+28))
							if v114&int32(4) == int32(0) {
								v133 = F_expanded_record_fetch_field(m, v94, v111, l5)
								mBase = m.M
								v134 = m.ExcPending
								if v134 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = v133
									m.G0 = v12 + int32(32)
									return
								}
							} else {
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v94)+64))
								if v119 < v111 {
									v133 = F_expanded_record_fetch_field(m, v94, v111, l5)
									mBase = m.M
									v134 = m.ExcPending
									if v134 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l4))) = v133
										m.G0 = v12 + int32(32)
										return
									}
								} else {
									v122 = v111 - int32(1)
									v123 = *(*int32)(unsafe.Add(mBase, uint32(v94)+60))
									v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v123))))
									*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v125)
									v127 = *(*int32)(unsafe.Add(mBase, uint32(v94)+56))
									v131 = *(*int32)(unsafe.Add(mBase, uint32(v127+v122<<(uint(int32(2))%32))))
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = v131
									m.G0 = v12 + int32(32)
									return
								}
							}
						}
					}
				}
			} else {
				v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v107
				v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v109
				v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
				if v111 <= int32(0) {
					v133 = F_expanded_record_fetch_field(m, v94, v111, l5)
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v133
						m.G0 = v12 + int32(32)
						return
					}
				} else {
					v114 = *(*int32)(unsafe.Add(mBase, uint32(v94)+28))
					if v114&int32(4) == int32(0) {
						v133 = F_expanded_record_fetch_field(m, v94, v111, l5)
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v133
							m.G0 = v12 + int32(32)
							return
						}
					} else {
						v119 = *(*int32)(unsafe.Add(mBase, uint32(v94)+64))
						if v119 < v111 {
							v133 = F_expanded_record_fetch_field(m, v94, v111, l5)
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v133
								m.G0 = v12 + int32(32)
								return
							}
						} else {
							v122 = v111 - int32(1)
							v123 = *(*int32)(unsafe.Add(mBase, uint32(v94)+60))
							v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v123))))
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v125)
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v94)+56))
							v131 = *(*int32)(unsafe.Add(mBase, uint32(v127+v122<<(uint(int32(2))%32))))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v131
							m.G0 = v12 + int32(32)
							return
						}
					}
				}
			}
		}
	case 4:
		F_plpgsql_fulfill_promise(m, l0, l1)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v18
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v20
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v22
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
			*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v24)
			m.G0 = v12 + int32(32)
			return
		}
	default:
		F_errstart_cold(m, int32(21), int32(525467))
		mBase = m.M
		v141 = m.ExcPending
		if v141 != 0 {
			return
		} else {
			v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v142
			F_errmsg_internal(m, int32(461960), v12)
			mBase = m.M
			v147 = m.ExcPending
			if v147 != 0 {
				return
			} else {
				F_errfinish(m, int32(476889), int32(5428), int32(272333))
				mBase = m.M
				v154 = m.ExcPending
				if v154 != 0 {
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
