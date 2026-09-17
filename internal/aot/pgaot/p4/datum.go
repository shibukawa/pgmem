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
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1|base.B2i32(int32(0) < l2) != 0 {
		v64 = l2
		m.G0 = v7 + int32(16)
		return v64
	} else {
		switch l2 + int32(2) {
		case 0:
			if l0 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_datumGetSize_0), int32(0))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_datumGetSize_1), int32(102), int32(_a_F_datumGetSize_2))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
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
				v41 = F_strlen(m, l0)
				mBase = m.M
				v64 = v41 + int32(1)
				m.G0 = v7 + int32(16)
				return v64
			}
		case 1:
			if l0 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_datumGetSize_0), int32(0))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_datumGetSize_1), int32(90), int32(_a_F_datumGetSize_2))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
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
				v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				if v16 == int32(1) {
					v20 = int32(18)
					v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
					if v22 == v20 {
						v25 = v20
					} else {
						v25 = int32(2)
					}
					if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v32 = int32(6)
					} else {
						v32 = v25
					}
					v64 = v32
				} else {
					if v16&int32(1) == int32(0) {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v64 = int32(base.Ui32(v59) >> (uint(int32(2)) % 32))
					} else {
						v64 = int32(base.Ui32(v16) >> (uint(int32(1)) % 32))
					}
				}
				m.G0 = v7 + int32(16)
				return v64
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l2
				F_errmsg_internal(m, int32(_a_F_datumGetSize_3), v7)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_datumGetSize_1), int32(108), int32(_a_F_datumGetSize_2))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
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
func F_datumTransfer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	if l1|base.B2i32(l2 != int32(-1)) != 0 {
		v21 = F_datumCopy(m, l0, l1, l2)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			return v21
		}
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		if v7 != int32(1) {
			v21 = F_datumCopy(m, l0, l1, l2)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				return v21
			}
		} else {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			if v10 != int32(3) {
				v21 = F_datumCopy(m, l0, l1, l2)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return v21
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, _c_F_datumTransfer[0]))
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				F_MemoryContextSetParent(m, v16, v14)
				mBase = m.M
				return v15 + int32(12)
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
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
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
			F_appendBinaryStringInfo(m, l2, int32(_a_F_datum_to_json_internal_0), int32(4))
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
							F_appendBinaryStringInfo(m, l2, int32(_a_F_datum_to_json_internal_1), int32(4))
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
							F_appendBinaryStringInfo(m, l2, int32(_a_F_datum_to_json_internal_2), int32(5))
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
					v82 = F_OidOutputFunctionCall(m, l4, l0)
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						v85 = v82
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v85)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return
								} else {
									F_pfree(m, v85)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
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
						v101 = m.ExcPending
						if v101 != 0 {
							return
						} else {
							F_appendStringInfoChar(m, l2, int32(34))
							mBase = m.M
							v312 = m.ExcPending
							if v312 != 0 {
								return
							} else {
								F_appendStringInfoString(m, l2, v10)
								mBase = m.M
								v314 = m.ExcPending
								if v314 != 0 {
									return
								} else {
									F_appendStringInfoChar(m, l2, int32(34))
									mBase = m.M
									v317 = m.ExcPending
									if v317 != 0 {
										return
									} else {
										m.G0 = v10 + int32(176)
										return
									}
								}
							}
						}
					} else {
						v113 = l0 + int32(_a_F_datum_to_json_internal_3)
						v114 = int32(_a_F_datum_to_json_internal_4)
						v115 = base.I32_div_u_s(v113, v114)
						v116 = int32(3)
						v122 = int32(2)
						v127 = base.I32_div_u_s((v115*int32(1073595727)+v113)<<(uint(v122)%32)|v116, v114)
						v130 = l0 + int32(_a_F_datum_to_json_internal_5) + v115*v116 + v127 + int32(_a_F_datum_to_json_internal_6)
						v131 = int32(1461)
						v132 = base.I32_div_u_s(v130, v131)
						v135 = v132*int32(-1461) + v130
						v137 = v135 << (uint(v122) % 32)
						if base.Ui32(v131) <= base.Ui32(v137) {
							v143 = base.I32_rem_u_s(v135+int32(305), int32(365))
							v148 = v143
						} else {
							v147 = base.I32_rem_u_s(v135+int32(306), int32(366))
							v148 = v147
						}
						v150 = base.I32_div_u_s(v137, int32(1461))
						*(*int32)(unsafe.Add(mBase, uint32(v10+int32(152)))) = v150 + v132<<(uint(int32(2))%32) - int32(_a_F_datum_to_json_internal_7)
						v158 = v148 + int32(123)
						v162 = int32(base.Ui32(v158*int32(2141)) >> (uint(int32(16)) % 32))
						*(*int32)(unsafe.Add(mBase, uint32(v10+int32(144)))) = v158 - int32(base.Ui32(v162*int32(_a_F_datum_to_json_internal_8))>>(uint(int32(8))%32))
						v172 = base.I32_rem_u_s(v162+int32(10), int32(12))
						*(*int32)(unsafe.Add(mBase, uint32(v10+int32(148)))) = v172 + int32(1)
						v177 = v10 + int32(132)
						switch int32(3) {
						case 0, 3:
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
							if int32(0) < v181 {
								v186 = v181
							} else {
								v186 = int32(1) - v181
							}
							v188 = F_pg_ultostr_zeropad(m, v10, v186, int32(4))
							mBase = m.M
							v189 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v189)
							v191 = int32(1)
							v193 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
							v194 = int32(2)
							v195 = F_pg_ultostr_zeropad(m, v188+v191, v193, v194)
							mBase = m.M
							*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v189)
							v200 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
							v202 = F_pg_ultostr_zeropad(m, v195+v191, v200, v194)
							mBase = m.M
							v295 = v202
						case 1:
							v206 = *(*int32)(unsafe.Add(mBase, _c_F_datum_to_json_internal[0]))
							v208 = base.B2i32(v206 == int32(1))
							if v206 == int32(1) {
								v209 = int32(12)
							} else {
								v209 = int32(16)
							}
							v211 = *(*int32)(unsafe.Add(mBase, uint32(v177+v209)))
							v213 = F_pg_ultostr_zeropad(m, v10, v211, int32(2))
							mBase = m.M
							v214 = int32(47)
							*(*uint8)(unsafe.Add(mBase, uint32(v213))) = uint8(v214)
							if v206 == int32(1) {
								v220 = int32(16)
							} else {
								v220 = int32(12)
							}
							v222 = *(*int32)(unsafe.Add(mBase, uint32(v177+v220)))
							v224 = F_pg_ultostr_zeropad(m, v213+int32(1), v222, int32(2))
							mBase = m.M
							v225 = int32(47)
							*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v225)
							v227 = int32(1)
							v229 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
							if int32(0) < v229 {
								v234 = v229
							} else {
								v234 = v227 - v229
							}
							v236 = F_pg_ultostr_zeropad(m, v224+v227, v234, int32(4))
							mBase = m.M
							v295 = v236
						case 2:
							v237 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
							v238 = int32(2)
							v239 = F_pg_ultostr_zeropad(m, v10, v237, v238)
							mBase = m.M
							v240 = int32(46)
							*(*uint8)(unsafe.Add(mBase, uint32(v239))) = uint8(v240)
							v242 = int32(1)
							v244 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
							v246 = F_pg_ultostr_zeropad(m, v239+v242, v244, v238)
							mBase = m.M
							*(*uint8)(unsafe.Add(mBase, uint32(v246))) = uint8(v240)
							v251 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
							if int32(0) < v251 {
								v256 = v251
							} else {
								v256 = v242 - v251
							}
							v258 = F_pg_ultostr_zeropad(m, v246+v242, v256, int32(4))
							mBase = m.M
							v295 = v258
						default:
							v262 = *(*int32)(unsafe.Add(mBase, _c_F_datum_to_json_internal[0]))
							v264 = base.B2i32(v262 == int32(1))
							if v262 == int32(1) {
								v265 = int32(12)
							} else {
								v265 = int32(16)
							}
							v267 = *(*int32)(unsafe.Add(mBase, uint32(v177+v265)))
							v269 = F_pg_ultostr_zeropad(m, v10, v267, int32(2))
							mBase = m.M
							v270 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v269))) = uint8(v270)
							if v262 == int32(1) {
								v276 = int32(16)
							} else {
								v276 = int32(12)
							}
							v278 = *(*int32)(unsafe.Add(mBase, uint32(v177+v276)))
							v280 = F_pg_ultostr_zeropad(m, v269+int32(1), v278, int32(2))
							mBase = m.M
							v281 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v280))) = uint8(v281)
							v283 = int32(1)
							v285 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
							if int32(0) < v285 {
								v290 = v285
							} else {
								v290 = v283 - v285
							}
							v292 = F_pg_ultostr_zeropad(m, v280+v283, v290, int32(4))
							mBase = m.M
							v295 = v292
						}
						v296 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
						if v296 <= int32(0) {
							v300 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_datum_to_json_internal[1])))
							*(*uint8)(unsafe.Add(mBase, uint32(v295)+2)) = uint8(v300)
							v303 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_datum_to_json_internal[2])))
							*(*uint16)(unsafe.Add(mBase, uint32(v295))) = uint16(v303)
							v307 = v295 + int32(3)
						} else {
							v307 = v295
						}
						v308 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v307))) = uint8(v308)
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v312 = m.ExcPending
						if v312 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v10)
							mBase = m.M
							v314 = m.ExcPending
							if v314 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v317 = m.ExcPending
								if v317 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					}
				case 3:
					v320 = F_JsonEncodeDateTime(m, v10, l0, int32(1114), int32(0))
					mBase = m.M
					v321 = m.ExcPending
					if v321 != 0 {
						return
					} else {
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v324 = m.ExcPending
						if v324 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v10)
							mBase = m.M
							v326 = m.ExcPending
							if v326 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v329 = m.ExcPending
								if v329 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					}
				case 4:
					v332 = F_JsonEncodeDateTime(m, v10, l0, int32(1184), int32(0))
					mBase = m.M
					v333 = m.ExcPending
					if v333 != 0 {
						return
					} else {
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v336 = m.ExcPending
						if v336 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v10)
							mBase = m.M
							v338 = m.ExcPending
							if v338 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v341 = m.ExcPending
								if v341 != 0 {
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
							F_errmsg(m, int32(_a_F_datum_to_json_internal_9), int32(0))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_datum_to_json_internal_10), int32(204), int32(_a_F_datum_to_json_internal_11))
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
						v396 = F_pg_detoast_datum_packed(m, l0)
						mBase = m.M
						v397 = m.ExcPending
						if v397 != 0 {
							return
						} else {
							v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396))))
							if v398 == int32(1) {
								v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+1)))
								if v404 == int32(18) {
									v407 = int32(16)
								} else {
									v407 = int32(0)
								}
								if base.Ui32((v404-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v414 = int32(4)
								} else {
									v414 = v407
								}
								v427 = v414
							} else {
								v415 = int32(1)
								if v398&v415 != 0 {
									v427 = int32(base.Ui32(v398)>>(uint(v415)%32)) - v415
								} else {
									v421 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
									v427 = int32(base.Ui32(v421)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							v428 = int32(1)
							if v398&v428 != 0 {
								v432 = v428
							} else {
								v432 = int32(4)
							}
							F_escape_json_with_len(m, l2, v396+v432, v427)
							mBase = m.M
							v435 = m.ExcPending
							if v435 != 0 {
								return
							} else {
								if l0 == v396 {
									m.G0 = v10 + int32(176)
									return
								} else {
									F_pfree(m, v396)
									mBase = m.M
									v438 = m.ExcPending
									if v438 != 0 {
										return
									} else {
										m.G0 = v10 + int32(176)
										return
									}
								}
							}
						}
					case 1:
						v439 = F_OidOutputFunctionCall(m, l4, l0)
						mBase = m.M
						v440 = m.ExcPending
						if v440 != 0 {
							return
						} else {
							F_escape_json(m, l2, v439)
							mBase = m.M
							v442 = m.ExcPending
							if v442 != 0 {
								return
							} else {
								F_pfree(m, v439)
								mBase = m.M
								v444 = m.ExcPending
								if v444 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					default:
						if l4 != int32(47) {
							v439 = F_OidOutputFunctionCall(m, l4, l0)
							mBase = m.M
							v440 = m.ExcPending
							if v440 != 0 {
								return
							} else {
								F_escape_json(m, l2, v439)
								mBase = m.M
								v442 = m.ExcPending
								if v442 != 0 {
									return
								} else {
									F_pfree(m, v439)
									mBase = m.M
									v444 = m.ExcPending
									if v444 != 0 {
										return
									} else {
										m.G0 = v10 + int32(176)
										return
									}
								}
							}
						} else {
							v396 = F_pg_detoast_datum_packed(m, l0)
							mBase = m.M
							v397 = m.ExcPending
							if v397 != 0 {
								return
							} else {
								v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396))))
								if v398 == int32(1) {
									v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+1)))
									if v404 == int32(18) {
										v407 = int32(16)
									} else {
										v407 = int32(0)
									}
									if base.Ui32((v404-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v414 = int32(4)
									} else {
										v414 = v407
									}
									v427 = v414
								} else {
									v415 = int32(1)
									if v398&v415 != 0 {
										v427 = int32(base.Ui32(v398)>>(uint(v415)%32)) - v415
									} else {
										v421 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
										v427 = int32(base.Ui32(v421)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								v428 = int32(1)
								if v398&v428 != 0 {
									v432 = v428
								} else {
									v432 = int32(4)
								}
								F_escape_json_with_len(m, l2, v396+v432, v427)
								mBase = m.M
								v435 = m.ExcPending
								if v435 != 0 {
									return
								} else {
									if l0 == v396 {
										m.G0 = v10 + int32(176)
										return
									} else {
										F_pfree(m, v396)
										mBase = m.M
										v438 = m.ExcPending
										if v438 != 0 {
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
						F_appendBinaryStringInfo(m, l2, int32(_a_F_datum_to_json_internal_1), int32(4))
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
						F_appendBinaryStringInfo(m, l2, int32(_a_F_datum_to_json_internal_2), int32(5))
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
							if v62 != int32(45) {
								v85 = v60
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									F_appendStringInfoString(m, l2, v85)
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										F_appendStringInfoChar(m, l2, int32(34))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											F_pfree(m, v85)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												m.G0 = v10 + int32(176)
												return
											}
										}
									}
								}
							} else {
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
								if base.Ui32(int32(9)) < base.Ui32((v71-int32(48))&int32(255)) {
									v85 = v60
									F_appendStringInfoChar(m, l2, int32(34))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										F_appendStringInfoString(m, l2, v85)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											F_appendStringInfoChar(m, l2, int32(34))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												F_pfree(m, v85)
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
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
									v79 = m.ExcPending
									if v79 != 0 {
										return
									} else {
										F_pfree(m, v60)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
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
							v79 = m.ExcPending
							if v79 != 0 {
								return
							} else {
								F_pfree(m, v60)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
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
						v101 = m.ExcPending
						if v101 != 0 {
							return
						} else {
							F_appendStringInfoChar(m, l2, int32(34))
							mBase = m.M
							v312 = m.ExcPending
							if v312 != 0 {
								return
							} else {
								F_appendStringInfoString(m, l2, v10)
								mBase = m.M
								v314 = m.ExcPending
								if v314 != 0 {
									return
								} else {
									F_appendStringInfoChar(m, l2, int32(34))
									mBase = m.M
									v317 = m.ExcPending
									if v317 != 0 {
										return
									} else {
										m.G0 = v10 + int32(176)
										return
									}
								}
							}
						}
					} else {
						v113 = l0 + int32(_a_F_datum_to_json_internal_3)
						v114 = int32(_a_F_datum_to_json_internal_4)
						v115 = base.I32_div_u_s(v113, v114)
						v116 = int32(3)
						v122 = int32(2)
						v127 = base.I32_div_u_s((v115*int32(1073595727)+v113)<<(uint(v122)%32)|v116, v114)
						v130 = l0 + int32(_a_F_datum_to_json_internal_5) + v115*v116 + v127 + int32(_a_F_datum_to_json_internal_6)
						v131 = int32(1461)
						v132 = base.I32_div_u_s(v130, v131)
						v135 = v132*int32(-1461) + v130
						v137 = v135 << (uint(v122) % 32)
						if base.Ui32(v131) <= base.Ui32(v137) {
							v143 = base.I32_rem_u_s(v135+int32(305), int32(365))
							v148 = v143
						} else {
							v147 = base.I32_rem_u_s(v135+int32(306), int32(366))
							v148 = v147
						}
						v150 = base.I32_div_u_s(v137, int32(1461))
						*(*int32)(unsafe.Add(mBase, uint32(v10+int32(152)))) = v150 + v132<<(uint(int32(2))%32) - int32(_a_F_datum_to_json_internal_7)
						v158 = v148 + int32(123)
						v162 = int32(base.Ui32(v158*int32(2141)) >> (uint(int32(16)) % 32))
						*(*int32)(unsafe.Add(mBase, uint32(v10+int32(144)))) = v158 - int32(base.Ui32(v162*int32(_a_F_datum_to_json_internal_8))>>(uint(int32(8))%32))
						v172 = base.I32_rem_u_s(v162+int32(10), int32(12))
						*(*int32)(unsafe.Add(mBase, uint32(v10+int32(148)))) = v172 + int32(1)
						v177 = v10 + int32(132)
						switch int32(3) {
						case 0, 3:
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
							if int32(0) < v181 {
								v186 = v181
							} else {
								v186 = int32(1) - v181
							}
							v188 = F_pg_ultostr_zeropad(m, v10, v186, int32(4))
							mBase = m.M
							v189 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v189)
							v191 = int32(1)
							v193 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
							v194 = int32(2)
							v195 = F_pg_ultostr_zeropad(m, v188+v191, v193, v194)
							mBase = m.M
							*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v189)
							v200 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
							v202 = F_pg_ultostr_zeropad(m, v195+v191, v200, v194)
							mBase = m.M
							v295 = v202
						case 1:
							v206 = *(*int32)(unsafe.Add(mBase, _c_F_datum_to_json_internal[0]))
							v208 = base.B2i32(v206 == int32(1))
							if v206 == int32(1) {
								v209 = int32(12)
							} else {
								v209 = int32(16)
							}
							v211 = *(*int32)(unsafe.Add(mBase, uint32(v177+v209)))
							v213 = F_pg_ultostr_zeropad(m, v10, v211, int32(2))
							mBase = m.M
							v214 = int32(47)
							*(*uint8)(unsafe.Add(mBase, uint32(v213))) = uint8(v214)
							if v206 == int32(1) {
								v220 = int32(16)
							} else {
								v220 = int32(12)
							}
							v222 = *(*int32)(unsafe.Add(mBase, uint32(v177+v220)))
							v224 = F_pg_ultostr_zeropad(m, v213+int32(1), v222, int32(2))
							mBase = m.M
							v225 = int32(47)
							*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v225)
							v227 = int32(1)
							v229 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
							if int32(0) < v229 {
								v234 = v229
							} else {
								v234 = v227 - v229
							}
							v236 = F_pg_ultostr_zeropad(m, v224+v227, v234, int32(4))
							mBase = m.M
							v295 = v236
						case 2:
							v237 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
							v238 = int32(2)
							v239 = F_pg_ultostr_zeropad(m, v10, v237, v238)
							mBase = m.M
							v240 = int32(46)
							*(*uint8)(unsafe.Add(mBase, uint32(v239))) = uint8(v240)
							v242 = int32(1)
							v244 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
							v246 = F_pg_ultostr_zeropad(m, v239+v242, v244, v238)
							mBase = m.M
							*(*uint8)(unsafe.Add(mBase, uint32(v246))) = uint8(v240)
							v251 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
							if int32(0) < v251 {
								v256 = v251
							} else {
								v256 = v242 - v251
							}
							v258 = F_pg_ultostr_zeropad(m, v246+v242, v256, int32(4))
							mBase = m.M
							v295 = v258
						default:
							v262 = *(*int32)(unsafe.Add(mBase, _c_F_datum_to_json_internal[0]))
							v264 = base.B2i32(v262 == int32(1))
							if v262 == int32(1) {
								v265 = int32(12)
							} else {
								v265 = int32(16)
							}
							v267 = *(*int32)(unsafe.Add(mBase, uint32(v177+v265)))
							v269 = F_pg_ultostr_zeropad(m, v10, v267, int32(2))
							mBase = m.M
							v270 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v269))) = uint8(v270)
							if v262 == int32(1) {
								v276 = int32(16)
							} else {
								v276 = int32(12)
							}
							v278 = *(*int32)(unsafe.Add(mBase, uint32(v177+v276)))
							v280 = F_pg_ultostr_zeropad(m, v269+int32(1), v278, int32(2))
							mBase = m.M
							v281 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v280))) = uint8(v281)
							v283 = int32(1)
							v285 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
							if int32(0) < v285 {
								v290 = v285
							} else {
								v290 = v283 - v285
							}
							v292 = F_pg_ultostr_zeropad(m, v280+v283, v290, int32(4))
							mBase = m.M
							v295 = v292
						}
						v296 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
						if v296 <= int32(0) {
							v300 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_datum_to_json_internal[1])))
							*(*uint8)(unsafe.Add(mBase, uint32(v295)+2)) = uint8(v300)
							v303 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_datum_to_json_internal[2])))
							*(*uint16)(unsafe.Add(mBase, uint32(v295))) = uint16(v303)
							v307 = v295 + int32(3)
						} else {
							v307 = v295
						}
						v308 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v307))) = uint8(v308)
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v312 = m.ExcPending
						if v312 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v10)
							mBase = m.M
							v314 = m.ExcPending
							if v314 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v317 = m.ExcPending
								if v317 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					}
				case 3:
					v320 = F_JsonEncodeDateTime(m, v10, l0, int32(1114), int32(0))
					mBase = m.M
					v321 = m.ExcPending
					if v321 != 0 {
						return
					} else {
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v324 = m.ExcPending
						if v324 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v10)
							mBase = m.M
							v326 = m.ExcPending
							if v326 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v329 = m.ExcPending
								if v329 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					}
				case 4:
					v332 = F_JsonEncodeDateTime(m, v10, l0, int32(1184), int32(0))
					mBase = m.M
					v333 = m.ExcPending
					if v333 != 0 {
						return
					} else {
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v336 = m.ExcPending
						if v336 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v10)
							mBase = m.M
							v338 = m.ExcPending
							if v338 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v341 = m.ExcPending
								if v341 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					}
				case 5:
					v342 = F_OidOutputFunctionCall(m, l4, l0)
					mBase = m.M
					v343 = m.ExcPending
					if v343 != 0 {
						return
					} else {
						F_appendStringInfoString(m, l2, v342)
						mBase = m.M
						v345 = m.ExcPending
						if v345 != 0 {
							return
						} else {
							F_pfree(m, v342)
							mBase = m.M
							v347 = m.ExcPending
							if v347 != 0 {
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
						v396 = F_pg_detoast_datum_packed(m, l0)
						mBase = m.M
						v397 = m.ExcPending
						if v397 != 0 {
							return
						} else {
							v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396))))
							if v398 == int32(1) {
								v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+1)))
								if v404 == int32(18) {
									v407 = int32(16)
								} else {
									v407 = int32(0)
								}
								if base.Ui32((v404-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v414 = int32(4)
								} else {
									v414 = v407
								}
								v427 = v414
							} else {
								v415 = int32(1)
								if v398&v415 != 0 {
									v427 = int32(base.Ui32(v398)>>(uint(v415)%32)) - v415
								} else {
									v421 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
									v427 = int32(base.Ui32(v421)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							v428 = int32(1)
							if v398&v428 != 0 {
								v432 = v428
							} else {
								v432 = int32(4)
							}
							F_escape_json_with_len(m, l2, v396+v432, v427)
							mBase = m.M
							v435 = m.ExcPending
							if v435 != 0 {
								return
							} else {
								if l0 == v396 {
									m.G0 = v10 + int32(176)
									return
								} else {
									F_pfree(m, v396)
									mBase = m.M
									v438 = m.ExcPending
									if v438 != 0 {
										return
									} else {
										m.G0 = v10 + int32(176)
										return
									}
								}
							}
						}
					case 1:
						v439 = F_OidOutputFunctionCall(m, l4, l0)
						mBase = m.M
						v440 = m.ExcPending
						if v440 != 0 {
							return
						} else {
							F_escape_json(m, l2, v439)
							mBase = m.M
							v442 = m.ExcPending
							if v442 != 0 {
								return
							} else {
								F_pfree(m, v439)
								mBase = m.M
								v444 = m.ExcPending
								if v444 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					default:
						if l4 != int32(47) {
							v439 = F_OidOutputFunctionCall(m, l4, l0)
							mBase = m.M
							v440 = m.ExcPending
							if v440 != 0 {
								return
							} else {
								F_escape_json(m, l2, v439)
								mBase = m.M
								v442 = m.ExcPending
								if v442 != 0 {
									return
								} else {
									F_pfree(m, v439)
									mBase = m.M
									v444 = m.ExcPending
									if v444 != 0 {
										return
									} else {
										m.G0 = v10 + int32(176)
										return
									}
								}
							}
						} else {
							v396 = F_pg_detoast_datum_packed(m, l0)
							mBase = m.M
							v397 = m.ExcPending
							if v397 != 0 {
								return
							} else {
								v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396))))
								if v398 == int32(1) {
									v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+1)))
									if v404 == int32(18) {
										v407 = int32(16)
									} else {
										v407 = int32(0)
									}
									if base.Ui32((v404-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v414 = int32(4)
									} else {
										v414 = v407
									}
									v427 = v414
								} else {
									v415 = int32(1)
									if v398&v415 != 0 {
										v427 = int32(base.Ui32(v398)>>(uint(v415)%32)) - v415
									} else {
										v421 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
										v427 = int32(base.Ui32(v421)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								v428 = int32(1)
								if v398&v428 != 0 {
									v432 = v428
								} else {
									v432 = int32(4)
								}
								F_escape_json_with_len(m, l2, v396+v432, v427)
								mBase = m.M
								v435 = m.ExcPending
								if v435 != 0 {
									return
								} else {
									if l0 == v396 {
										m.G0 = v10 + int32(176)
										return
									} else {
										F_pfree(m, v396)
										mBase = m.M
										v438 = m.ExcPending
										if v438 != 0 {
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
					v349 = F_OidFunctionCall1Coll(m, l4, int32(0), l0)
					mBase = m.M
					v350 = m.ExcPending
					if v350 != 0 {
						return
					} else {
						v351 = F_pg_detoast_datum_packed(m, v349)
						mBase = m.M
						v352 = m.ExcPending
						if v352 != 0 {
							return
						} else {
							v353 = int32(1)
							v354 = v351 + v353
							v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
							v359 = v357 & v353
							if v359 != 0 {
								v360 = v354
							} else {
								v360 = v351 + int32(4)
							}
							if v357 == int32(1) {
								v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354))))
								if v366 == int32(18) {
									v369 = int32(16)
								} else {
									v369 = int32(0)
								}
								if base.Ui32((v366-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v376 = int32(4)
								} else {
									v376 = v369
								}
								v387 = v376
							} else {
								v377 = int32(1)
								if v359 != 0 {
									v387 = int32(base.Ui32(v357)>>(uint(v377)%32)) - v377
								} else {
									v381 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
									v387 = int32(base.Ui32(v381)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							F_appendBinaryStringInfo(m, l2, v360, v387)
							mBase = m.M
							v389 = m.ExcPending
							if v389 != 0 {
								return
							} else {
								F_pfree(m, v351)
								mBase = m.M
								v391 = m.ExcPending
								if v391 != 0 {
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int64
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
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
			F_errstart_cold(m, int32(21), int32(_a_F_exec_eval_datum_0))
			mBase = m.M
			v166 = m.ExcPending
			if v166 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_exec_eval_datum_1), int32(0))
				mBase = m.M
				v170 = m.ExcPending
				if v170 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_exec_eval_datum_2), int32(_a_F_exec_eval_datum_3), int32(_a_F_exec_eval_datum_4))
					mBase = m.M
					v175 = m.ExcPending
					if v175 != 0 {
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
				v31 = int32(_a_F_exec_eval_datum_5)
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_exec_eval_datum[0]))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
				*(*int32)(unsafe.Add(mBase, _c_F_exec_eval_datum[0])) = v35
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
				v38 = F_make_tuple_from_row(m, l0, l1, v37)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					if v38 == int32(0) {
						F_errstart_cold(m, int32(21), int32(_a_F_exec_eval_datum_0))
						mBase = m.M
						v179 = m.ExcPending
						if v179 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_exec_eval_datum_6), int32(0))
							mBase = m.M
							v183 = m.ExcPending
							if v183 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_exec_eval_datum_2), int32(_a_F_exec_eval_datum_7), int32(_a_F_exec_eval_datum_4))
								mBase = m.M
								v188 = m.ExcPending
								if v188 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v43
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v45
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
						v48 = F_HeapTupleHeaderGetDatum(m, v47)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v48
							v51 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v51)
							*(*int32)(unsafe.Add(mBase, _c_F_exec_eval_datum[0])) = v32
							m.G0 = v12 + int32(32)
							return
						}
					}
				}
			}
		}
	case 2:
		v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
		if v55 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
			v60 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v60)
			v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v62
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(-1)
		} else {
			v69 = *(*int32)(unsafe.Add(mBase, uint32(v55)+28))
			v71 = v69 & int32(5)
			if v71 != 0 {
				v72 = v55 + int32(12)
			} else {
				v72 = int32(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v72
			*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(base.B2i32(v71 == int32(0)))
			v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			if v77 == int32(2249) {
				v151 = *(*int32)(unsafe.Add(mBase, uint32(v55)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v151
				v153 = *(*int32)(unsafe.Add(mBase, uint32(v55)+40))
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v153
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(-1)
			}
		}
		m.G0 = v12 + int32(32)
		return
	case 3:
		v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v88 = *(*int32)(unsafe.Add(mBase, uint32(v83+v84<<(uint(int32(2))%32))))
		v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+36))
		if v89 == int32(0) {
			F_instantiate_empty_record_variable(m, l0, v88)
			mBase = m.M
			v93 = m.ExcPending
			if v93 != 0 {
				return
			} else {
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v88)+36))
				v95 = v94
				v96 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
				v97 = *(*int64)(unsafe.Add(mBase, uint32(v95)+48))
				if v96 != v97 {
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v102 = F_expanded_record_lookup_field(m, v95, v99, l1+int32(32))
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return
					} else {
						if v102 == int32(0) {
							F_errstart_cold(m, int32(21), int32(_a_F_exec_eval_datum_0))
							mBase = m.M
							v192 = m.ExcPending
							if v192 != 0 {
								return
							} else {
								F_errcode(m, int32(50360452))
								mBase = m.M
								v195 = m.ExcPending
								if v195 != 0 {
									return
								} else {
									v196 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
									v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v197
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v196
									F_errmsg(m, int32(_a_F_exec_eval_datum_8), v12+int32(16))
									mBase = m.M
									v204 = m.ExcPending
									if v204 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_exec_eval_datum_2), int32(_a_F_exec_eval_datum_9), int32(_a_F_exec_eval_datum_4))
										mBase = m.M
										v209 = m.ExcPending
										if v209 != 0 {
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
							v106 = *(*int64)(unsafe.Add(mBase, uint32(v95)+48))
							*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v106
							v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v108
							v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v110
							v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
							if v112 <= int32(0) {
								v134 = F_expanded_record_fetch_field(m, v95, v112, l5)
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = v134
									m.G0 = v12 + int32(32)
									return
								}
							} else {
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
								if v115&int32(4) == int32(0) {
									v134 = F_expanded_record_fetch_field(m, v95, v112, l5)
									mBase = m.M
									v135 = m.ExcPending
									if v135 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l4))) = v134
										m.G0 = v12 + int32(32)
										return
									}
								} else {
									v120 = *(*int32)(unsafe.Add(mBase, uint32(v95)+64))
									if v120 < v112 {
										v134 = F_expanded_record_fetch_field(m, v95, v112, l5)
										mBase = m.M
										v135 = m.ExcPending
										if v135 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l4))) = v134
											m.G0 = v12 + int32(32)
											return
										}
									} else {
										v123 = v112 - int32(1)
										v124 = *(*int32)(unsafe.Add(mBase, uint32(v95)+60))
										v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123+v124))))
										*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v126)
										v128 = *(*int32)(unsafe.Add(mBase, uint32(v95)+56))
										v132 = *(*int32)(unsafe.Add(mBase, uint32(v128+v123<<(uint(int32(2))%32))))
										*(*int32)(unsafe.Add(mBase, uint32(l4))) = v132
										m.G0 = v12 + int32(32)
										return
									}
								}
							}
						}
					}
				} else {
					v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v108
					v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v110
					v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
					if v112 <= int32(0) {
						v134 = F_expanded_record_fetch_field(m, v95, v112, l5)
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v134
							m.G0 = v12 + int32(32)
							return
						}
					} else {
						v115 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
						if v115&int32(4) == int32(0) {
							v134 = F_expanded_record_fetch_field(m, v95, v112, l5)
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v134
								m.G0 = v12 + int32(32)
								return
							}
						} else {
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v95)+64))
							if v120 < v112 {
								v134 = F_expanded_record_fetch_field(m, v95, v112, l5)
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = v134
									m.G0 = v12 + int32(32)
									return
								}
							} else {
								v123 = v112 - int32(1)
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v95)+60))
								v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123+v124))))
								*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v126)
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v95)+56))
								v132 = *(*int32)(unsafe.Add(mBase, uint32(v128+v123<<(uint(int32(2))%32))))
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v132
								m.G0 = v12 + int32(32)
								return
							}
						}
					}
				}
			}
		} else {
			v95 = v89
			v96 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
			v97 = *(*int64)(unsafe.Add(mBase, uint32(v95)+48))
			if v96 != v97 {
				v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v102 = F_expanded_record_lookup_field(m, v95, v99, l1+int32(32))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return
				} else {
					if v102 == int32(0) {
						F_errstart_cold(m, int32(21), int32(_a_F_exec_eval_datum_0))
						mBase = m.M
						v192 = m.ExcPending
						if v192 != 0 {
							return
						} else {
							F_errcode(m, int32(50360452))
							mBase = m.M
							v195 = m.ExcPending
							if v195 != 0 {
								return
							} else {
								v196 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
								v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v197
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v196
								F_errmsg(m, int32(_a_F_exec_eval_datum_8), v12+int32(16))
								mBase = m.M
								v204 = m.ExcPending
								if v204 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_exec_eval_datum_2), int32(_a_F_exec_eval_datum_9), int32(_a_F_exec_eval_datum_4))
									mBase = m.M
									v209 = m.ExcPending
									if v209 != 0 {
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
						v106 = *(*int64)(unsafe.Add(mBase, uint32(v95)+48))
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v106
						v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v108
						v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v110
						v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						if v112 <= int32(0) {
							v134 = F_expanded_record_fetch_field(m, v95, v112, l5)
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v134
								m.G0 = v12 + int32(32)
								return
							}
						} else {
							v115 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
							if v115&int32(4) == int32(0) {
								v134 = F_expanded_record_fetch_field(m, v95, v112, l5)
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = v134
									m.G0 = v12 + int32(32)
									return
								}
							} else {
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v95)+64))
								if v120 < v112 {
									v134 = F_expanded_record_fetch_field(m, v95, v112, l5)
									mBase = m.M
									v135 = m.ExcPending
									if v135 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l4))) = v134
										m.G0 = v12 + int32(32)
										return
									}
								} else {
									v123 = v112 - int32(1)
									v124 = *(*int32)(unsafe.Add(mBase, uint32(v95)+60))
									v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123+v124))))
									*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v126)
									v128 = *(*int32)(unsafe.Add(mBase, uint32(v95)+56))
									v132 = *(*int32)(unsafe.Add(mBase, uint32(v128+v123<<(uint(int32(2))%32))))
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = v132
									m.G0 = v12 + int32(32)
									return
								}
							}
						}
					}
				}
			} else {
				v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v108
				v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v110
				v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
				if v112 <= int32(0) {
					v134 = F_expanded_record_fetch_field(m, v95, v112, l5)
					mBase = m.M
					v135 = m.ExcPending
					if v135 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v134
						m.G0 = v12 + int32(32)
						return
					}
				} else {
					v115 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
					if v115&int32(4) == int32(0) {
						v134 = F_expanded_record_fetch_field(m, v95, v112, l5)
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v134
							m.G0 = v12 + int32(32)
							return
						}
					} else {
						v120 = *(*int32)(unsafe.Add(mBase, uint32(v95)+64))
						if v120 < v112 {
							v134 = F_expanded_record_fetch_field(m, v95, v112, l5)
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v134
								m.G0 = v12 + int32(32)
								return
							}
						} else {
							v123 = v112 - int32(1)
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v95)+60))
							v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123+v124))))
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v126)
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v95)+56))
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v128+v123<<(uint(int32(2))%32))))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v132
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
		F_errstart_cold(m, int32(21), int32(_a_F_exec_eval_datum_0))
		mBase = m.M
		v140 = m.ExcPending
		if v140 != 0 {
			return
		} else {
			v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v141
			F_errmsg_internal(m, int32(_a_F_exec_eval_datum_10), v12)
			mBase = m.M
			v145 = m.ExcPending
			if v145 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_exec_eval_datum_2), int32(_a_F_exec_eval_datum_11), int32(_a_F_exec_eval_datum_4))
				mBase = m.M
				v150 = m.ExcPending
				if v150 != 0 {
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
