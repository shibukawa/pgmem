package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_datumGetSize(m *base.Module, l0 int64, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1|base.B2i32(int32(0) < l2) != 0 {
		v65 = l2
		m.G0 = v7 + int32(16)
		return v65
	} else {
		switch l2 + int32(2) {
		case 0:
			v40 = base.I32_wrap_i64(l0)
			if v40 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_datumGetSize_0), int32(0))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_datumGetSize_1), int32(102), int32(_a_F_datumGetSize_2))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
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
				v43 = F_strlen(m, v40)
				mBase = m.M
				v65 = v43 + int32(1)
				m.G0 = v7 + int32(16)
				return v65
			}
		case 1:
			v14 = base.I32_wrap_i64(l0)
			if v14 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_datumGetSize_0), int32(0))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_datumGetSize_1), int32(90), int32(_a_F_datumGetSize_2))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
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
				v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
				if v17 == int32(1) {
					v21 = int32(18)
					v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
					if v23 == v21 {
						v26 = v21
					} else {
						v26 = int32(2)
					}
					if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v33 = int32(6)
					} else {
						v33 = v26
					}
					v65 = v33
				} else {
					if v17&int32(1) == int32(0) {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
						v65 = int32(base.Ui32(v61) >> (uint(int32(2)) % 32))
					} else {
						v65 = int32(base.Ui32(v17) >> (uint(int32(1)) % 32))
					}
				}
				m.G0 = v7 + int32(16)
				return v65
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l2
				F_errmsg_internal(m, int32(_a_F_datumGetSize_3), v7)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_datumGetSize_1), int32(108), int32(_a_F_datumGetSize_2))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
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
func F_datumTransfer(m *base.Module, l0 int64, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	if l1|base.B2i32(l2 != int32(-1)) != 0 {
		v27 = F_datumCopy(m, l0, l1, l2)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int64(0)
		} else {
			return v27
		}
	} else {
		v8 = base.I32_wrap_i64(l0)
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		if v9 != int32(1) {
			v27 = F_datumCopy(m, l0, l1, l2)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				return v27
			}
		} else {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
			if v12 != int32(3) {
				v27 = F_datumCopy(m, l0, l1, l2)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					return v27
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, _c_F_datumTransfer[0]))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0))+2))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
				F_MemoryContextSetParent(m, v20, v16)
				mBase = m.M
				return base.I64_extend_i32_u(v19 + int32(12))
			}
		}
	}
}
func F_datum_to_json_internal(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
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
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int64
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
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
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
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
						if l0 != int64(0) {
							F_appendBinaryStringInfo(m, l2, int32(_a_F_datum_to_json_internal_1), int32(4))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								if l5 == int32(0) {
									m.G0 = v10 + int32(176)
									return
								} else {
									F_appendStringInfoChar(m, l2, int32(34))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
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
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								if l5 == int32(0) {
									m.G0 = v10 + int32(176)
									return
								} else {
									F_appendStringInfoChar(m, l2, int32(34))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
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
						v86 = v84
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v86)
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
									F_pfree(m, v86)
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
					v98 = base.I32_wrap_i64(l0)
					if base.Ui32(v98-int32(2147483647)) <= base.Ui32(int32(1)) {
						F_EncodeSpecialDate(m, v98, v10)
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return
						} else {
							F_appendStringInfoChar(m, l2, int32(34))
							mBase = m.M
							v315 = m.ExcPending
							if v315 != 0 {
								return
							} else {
								F_appendStringInfoString(m, l2, v10)
								mBase = m.M
								v317 = m.ExcPending
								if v317 != 0 {
									return
								} else {
									F_appendStringInfoChar(m, l2, int32(34))
									mBase = m.M
									v320 = m.ExcPending
									if v320 != 0 {
										return
									} else {
										m.G0 = v10 + int32(176)
										return
									}
								}
							}
						}
					} else {
						v116 = v98 + int32(_a_F_datum_to_json_internal_3)
						v117 = int32(_a_F_datum_to_json_internal_4)
						v118 = base.I32_div_u_s(v116, v117)
						v119 = int32(3)
						v125 = int32(2)
						v130 = base.I32_div_u_s((v118*int32(1073595727)+v116)<<(uint(v125)%32)|v119, v117)
						v133 = v98 + int32(_a_F_datum_to_json_internal_5) + v118*v119 + v130 + int32(_a_F_datum_to_json_internal_6)
						v134 = int32(1461)
						v135 = base.I32_div_u_s(v133, v134)
						v138 = v135*int32(-1461) + v133
						v140 = v138 << (uint(v125) % 32)
						if base.Ui32(v134) <= base.Ui32(v140) {
							v146 = base.I32_rem_u_s(v138+int32(305), int32(365))
							v151 = v146
						} else {
							v150 = base.I32_rem_u_s(v138+int32(306), int32(366))
							v151 = v150
						}
						v153 = base.I32_div_u_s(v140, int32(1461))
						*(*int32)(unsafe.Add(mBase, uint32(v10+int32(152)))) = v153 + v135<<(uint(int32(2))%32) - int32(_a_F_datum_to_json_internal_7)
						v161 = v151 + int32(123)
						v165 = int32(base.Ui32(v161*int32(2141)) >> (uint(int32(16)) % 32))
						*(*int32)(unsafe.Add(mBase, uint32(v10+int32(144)))) = v161 - int32(base.Ui32(v165*int32(_a_F_datum_to_json_internal_8))>>(uint(int32(8))%32))
						v175 = base.I32_rem_u_s(v165+int32(10), int32(12))
						*(*int32)(unsafe.Add(mBase, uint32(v10+int32(148)))) = v175 + int32(1)
						v180 = v10 + int32(132)
						switch int32(3) {
						case 0, 3:
							v184 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
							if int32(0) < v184 {
								v189 = v184
							} else {
								v189 = int32(1) - v184
							}
							v191 = F_pg_ultostr_zeropad(m, v10, v189, int32(4))
							mBase = m.M
							v192 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v192)
							v194 = int32(1)
							v196 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
							v197 = int32(2)
							v198 = F_pg_ultostr_zeropad(m, v191+v194, v196, v197)
							mBase = m.M
							*(*uint8)(unsafe.Add(mBase, uint32(v198))) = uint8(v192)
							v203 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
							v205 = F_pg_ultostr_zeropad(m, v198+v194, v203, v197)
							mBase = m.M
							v298 = v205
						case 1:
							v209 = *(*int32)(unsafe.Add(mBase, _c_F_datum_to_json_internal[0]))
							v211 = base.B2i32(v209 == int32(1))
							if v209 == int32(1) {
								v212 = int32(12)
							} else {
								v212 = int32(16)
							}
							v214 = *(*int32)(unsafe.Add(mBase, uint32(v180+v212)))
							v216 = F_pg_ultostr_zeropad(m, v10, v214, int32(2))
							mBase = m.M
							v217 = int32(47)
							*(*uint8)(unsafe.Add(mBase, uint32(v216))) = uint8(v217)
							if v209 == int32(1) {
								v223 = int32(16)
							} else {
								v223 = int32(12)
							}
							v225 = *(*int32)(unsafe.Add(mBase, uint32(v180+v223)))
							v227 = F_pg_ultostr_zeropad(m, v216+int32(1), v225, int32(2))
							mBase = m.M
							v228 = int32(47)
							*(*uint8)(unsafe.Add(mBase, uint32(v227))) = uint8(v228)
							v230 = int32(1)
							v232 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
							if int32(0) < v232 {
								v237 = v232
							} else {
								v237 = v230 - v232
							}
							v239 = F_pg_ultostr_zeropad(m, v227+v230, v237, int32(4))
							mBase = m.M
							v298 = v239
						case 2:
							v240 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
							v241 = int32(2)
							v242 = F_pg_ultostr_zeropad(m, v10, v240, v241)
							mBase = m.M
							v243 = int32(46)
							*(*uint8)(unsafe.Add(mBase, uint32(v242))) = uint8(v243)
							v245 = int32(1)
							v247 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
							v249 = F_pg_ultostr_zeropad(m, v242+v245, v247, v241)
							mBase = m.M
							*(*uint8)(unsafe.Add(mBase, uint32(v249))) = uint8(v243)
							v254 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
							if int32(0) < v254 {
								v259 = v254
							} else {
								v259 = v245 - v254
							}
							v261 = F_pg_ultostr_zeropad(m, v249+v245, v259, int32(4))
							mBase = m.M
							v298 = v261
						default:
							v265 = *(*int32)(unsafe.Add(mBase, _c_F_datum_to_json_internal[0]))
							v267 = base.B2i32(v265 == int32(1))
							if v265 == int32(1) {
								v268 = int32(12)
							} else {
								v268 = int32(16)
							}
							v270 = *(*int32)(unsafe.Add(mBase, uint32(v180+v268)))
							v272 = F_pg_ultostr_zeropad(m, v10, v270, int32(2))
							mBase = m.M
							v273 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v272))) = uint8(v273)
							if v265 == int32(1) {
								v279 = int32(16)
							} else {
								v279 = int32(12)
							}
							v281 = *(*int32)(unsafe.Add(mBase, uint32(v180+v279)))
							v283 = F_pg_ultostr_zeropad(m, v272+int32(1), v281, int32(2))
							mBase = m.M
							v284 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v283))) = uint8(v284)
							v286 = int32(1)
							v288 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
							if int32(0) < v288 {
								v293 = v288
							} else {
								v293 = v286 - v288
							}
							v295 = F_pg_ultostr_zeropad(m, v283+v286, v293, int32(4))
							mBase = m.M
							v298 = v295
						}
						v299 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
						if v299 <= int32(0) {
							v303 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_datum_to_json_internal[1])))
							*(*uint8)(unsafe.Add(mBase, uint32(v298)+2)) = uint8(v303)
							v306 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_datum_to_json_internal[2])))
							*(*uint16)(unsafe.Add(mBase, uint32(v298))) = uint16(v306)
							v310 = v298 + int32(3)
						} else {
							v310 = v298
						}
						v311 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v310))) = uint8(v311)
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v315 = m.ExcPending
						if v315 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v10)
							mBase = m.M
							v317 = m.ExcPending
							if v317 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v320 = m.ExcPending
								if v320 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					}
				case 3:
					v323 = F_JsonEncodeDateTime(m, v10, l0, int32(1114), int32(0))
					mBase = m.M
					v324 = m.ExcPending
					if v324 != 0 {
						return
					} else {
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v327 = m.ExcPending
						if v327 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v10)
							mBase = m.M
							v329 = m.ExcPending
							if v329 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v332 = m.ExcPending
								if v332 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					}
				case 4:
					v335 = F_JsonEncodeDateTime(m, v10, l0, int32(1184), int32(0))
					mBase = m.M
					v336 = m.ExcPending
					if v336 != 0 {
						return
					} else {
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v339 = m.ExcPending
						if v339 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v10)
							mBase = m.M
							v341 = m.ExcPending
							if v341 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v344 = m.ExcPending
								if v344 != 0 {
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
								F_errfinish(m, int32(_a_F_datum_to_json_internal_10), int32(203), int32(_a_F_datum_to_json_internal_11))
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
						v399 = base.I32_wrap_i64(l0)
						v400 = F_pg_detoast_datum_packed(m, v399)
						mBase = m.M
						v401 = m.ExcPending
						if v401 != 0 {
							return
						} else {
							v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400))))
							if v402 == int32(1) {
								v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400)+1)))
								if v408 == int32(18) {
									v411 = int32(16)
								} else {
									v411 = int32(0)
								}
								if base.Ui32((v408-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v418 = int32(4)
								} else {
									v418 = v411
								}
								v431 = v418
							} else {
								v419 = int32(1)
								if v402&v419 != 0 {
									v431 = int32(base.Ui32(v402)>>(uint(v419)%32)) - v419
								} else {
									v425 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
									v431 = int32(base.Ui32(v425)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							v432 = int32(1)
							if v402&v432 != 0 {
								v436 = v432
							} else {
								v436 = int32(4)
							}
							F_escape_json_with_len(m, l2, v400+v436, v431)
							mBase = m.M
							v439 = m.ExcPending
							if v439 != 0 {
								return
							} else {
								if v400 == v399 {
									m.G0 = v10 + int32(176)
									return
								} else {
									F_pfree(m, v400)
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
							v399 = base.I32_wrap_i64(l0)
							v400 = F_pg_detoast_datum_packed(m, v399)
							mBase = m.M
							v401 = m.ExcPending
							if v401 != 0 {
								return
							} else {
								v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400))))
								if v402 == int32(1) {
									v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400)+1)))
									if v408 == int32(18) {
										v411 = int32(16)
									} else {
										v411 = int32(0)
									}
									if base.Ui32((v408-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v418 = int32(4)
									} else {
										v418 = v411
									}
									v431 = v418
								} else {
									v419 = int32(1)
									if v402&v419 != 0 {
										v431 = int32(base.Ui32(v402)>>(uint(v419)%32)) - v419
									} else {
										v425 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
										v431 = int32(base.Ui32(v425)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								v432 = int32(1)
								if v402&v432 != 0 {
									v436 = v432
								} else {
									v436 = int32(4)
								}
								F_escape_json_with_len(m, l2, v400+v436, v431)
								mBase = m.M
								v439 = m.ExcPending
								if v439 != 0 {
									return
								} else {
									if v400 == v399 {
										m.G0 = v10 + int32(176)
										return
									} else {
										F_pfree(m, v400)
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
					if l0 != int64(0) {
						F_appendBinaryStringInfo(m, l2, int32(_a_F_datum_to_json_internal_1), int32(4))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							if l5 == int32(0) {
								m.G0 = v10 + int32(176)
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
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
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							if l5 == int32(0) {
								m.G0 = v10 + int32(176)
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					}
				case 1:
					v62 = F_OidOutputFunctionCall(m, l4, l0)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
						if base.Ui32(int32(10)) <= base.Ui32((v64-int32(48))&int32(255)) {
							if v64 != int32(45) {
								v86 = v62
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return
								} else {
									F_appendStringInfoString(m, l2, v86)
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
											F_pfree(m, v86)
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
								v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
								if base.Ui32(int32(9)) < base.Ui32((v73-int32(48))&int32(255)) {
									v86 = v62
									F_appendStringInfoChar(m, l2, int32(34))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										F_appendStringInfoString(m, l2, v86)
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
												F_pfree(m, v86)
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
									F_appendStringInfoString(m, l2, v62)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										F_pfree(m, v62)
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
							F_appendStringInfoString(m, l2, v62)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								F_pfree(m, v62)
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
					v98 = base.I32_wrap_i64(l0)
					if base.Ui32(v98-int32(2147483647)) <= base.Ui32(int32(1)) {
						F_EncodeSpecialDate(m, v98, v10)
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return
						} else {
							F_appendStringInfoChar(m, l2, int32(34))
							mBase = m.M
							v315 = m.ExcPending
							if v315 != 0 {
								return
							} else {
								F_appendStringInfoString(m, l2, v10)
								mBase = m.M
								v317 = m.ExcPending
								if v317 != 0 {
									return
								} else {
									F_appendStringInfoChar(m, l2, int32(34))
									mBase = m.M
									v320 = m.ExcPending
									if v320 != 0 {
										return
									} else {
										m.G0 = v10 + int32(176)
										return
									}
								}
							}
						}
					} else {
						v116 = v98 + int32(_a_F_datum_to_json_internal_3)
						v117 = int32(_a_F_datum_to_json_internal_4)
						v118 = base.I32_div_u_s(v116, v117)
						v119 = int32(3)
						v125 = int32(2)
						v130 = base.I32_div_u_s((v118*int32(1073595727)+v116)<<(uint(v125)%32)|v119, v117)
						v133 = v98 + int32(_a_F_datum_to_json_internal_5) + v118*v119 + v130 + int32(_a_F_datum_to_json_internal_6)
						v134 = int32(1461)
						v135 = base.I32_div_u_s(v133, v134)
						v138 = v135*int32(-1461) + v133
						v140 = v138 << (uint(v125) % 32)
						if base.Ui32(v134) <= base.Ui32(v140) {
							v146 = base.I32_rem_u_s(v138+int32(305), int32(365))
							v151 = v146
						} else {
							v150 = base.I32_rem_u_s(v138+int32(306), int32(366))
							v151 = v150
						}
						v153 = base.I32_div_u_s(v140, int32(1461))
						*(*int32)(unsafe.Add(mBase, uint32(v10+int32(152)))) = v153 + v135<<(uint(int32(2))%32) - int32(_a_F_datum_to_json_internal_7)
						v161 = v151 + int32(123)
						v165 = int32(base.Ui32(v161*int32(2141)) >> (uint(int32(16)) % 32))
						*(*int32)(unsafe.Add(mBase, uint32(v10+int32(144)))) = v161 - int32(base.Ui32(v165*int32(_a_F_datum_to_json_internal_8))>>(uint(int32(8))%32))
						v175 = base.I32_rem_u_s(v165+int32(10), int32(12))
						*(*int32)(unsafe.Add(mBase, uint32(v10+int32(148)))) = v175 + int32(1)
						v180 = v10 + int32(132)
						switch int32(3) {
						case 0, 3:
							v184 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
							if int32(0) < v184 {
								v189 = v184
							} else {
								v189 = int32(1) - v184
							}
							v191 = F_pg_ultostr_zeropad(m, v10, v189, int32(4))
							mBase = m.M
							v192 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v192)
							v194 = int32(1)
							v196 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
							v197 = int32(2)
							v198 = F_pg_ultostr_zeropad(m, v191+v194, v196, v197)
							mBase = m.M
							*(*uint8)(unsafe.Add(mBase, uint32(v198))) = uint8(v192)
							v203 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
							v205 = F_pg_ultostr_zeropad(m, v198+v194, v203, v197)
							mBase = m.M
							v298 = v205
						case 1:
							v209 = *(*int32)(unsafe.Add(mBase, _c_F_datum_to_json_internal[0]))
							v211 = base.B2i32(v209 == int32(1))
							if v209 == int32(1) {
								v212 = int32(12)
							} else {
								v212 = int32(16)
							}
							v214 = *(*int32)(unsafe.Add(mBase, uint32(v180+v212)))
							v216 = F_pg_ultostr_zeropad(m, v10, v214, int32(2))
							mBase = m.M
							v217 = int32(47)
							*(*uint8)(unsafe.Add(mBase, uint32(v216))) = uint8(v217)
							if v209 == int32(1) {
								v223 = int32(16)
							} else {
								v223 = int32(12)
							}
							v225 = *(*int32)(unsafe.Add(mBase, uint32(v180+v223)))
							v227 = F_pg_ultostr_zeropad(m, v216+int32(1), v225, int32(2))
							mBase = m.M
							v228 = int32(47)
							*(*uint8)(unsafe.Add(mBase, uint32(v227))) = uint8(v228)
							v230 = int32(1)
							v232 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
							if int32(0) < v232 {
								v237 = v232
							} else {
								v237 = v230 - v232
							}
							v239 = F_pg_ultostr_zeropad(m, v227+v230, v237, int32(4))
							mBase = m.M
							v298 = v239
						case 2:
							v240 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
							v241 = int32(2)
							v242 = F_pg_ultostr_zeropad(m, v10, v240, v241)
							mBase = m.M
							v243 = int32(46)
							*(*uint8)(unsafe.Add(mBase, uint32(v242))) = uint8(v243)
							v245 = int32(1)
							v247 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
							v249 = F_pg_ultostr_zeropad(m, v242+v245, v247, v241)
							mBase = m.M
							*(*uint8)(unsafe.Add(mBase, uint32(v249))) = uint8(v243)
							v254 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
							if int32(0) < v254 {
								v259 = v254
							} else {
								v259 = v245 - v254
							}
							v261 = F_pg_ultostr_zeropad(m, v249+v245, v259, int32(4))
							mBase = m.M
							v298 = v261
						default:
							v265 = *(*int32)(unsafe.Add(mBase, _c_F_datum_to_json_internal[0]))
							v267 = base.B2i32(v265 == int32(1))
							if v265 == int32(1) {
								v268 = int32(12)
							} else {
								v268 = int32(16)
							}
							v270 = *(*int32)(unsafe.Add(mBase, uint32(v180+v268)))
							v272 = F_pg_ultostr_zeropad(m, v10, v270, int32(2))
							mBase = m.M
							v273 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v272))) = uint8(v273)
							if v265 == int32(1) {
								v279 = int32(16)
							} else {
								v279 = int32(12)
							}
							v281 = *(*int32)(unsafe.Add(mBase, uint32(v180+v279)))
							v283 = F_pg_ultostr_zeropad(m, v272+int32(1), v281, int32(2))
							mBase = m.M
							v284 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v283))) = uint8(v284)
							v286 = int32(1)
							v288 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
							if int32(0) < v288 {
								v293 = v288
							} else {
								v293 = v286 - v288
							}
							v295 = F_pg_ultostr_zeropad(m, v283+v286, v293, int32(4))
							mBase = m.M
							v298 = v295
						}
						v299 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
						if v299 <= int32(0) {
							v303 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_datum_to_json_internal[1])))
							*(*uint8)(unsafe.Add(mBase, uint32(v298)+2)) = uint8(v303)
							v306 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_datum_to_json_internal[2])))
							*(*uint16)(unsafe.Add(mBase, uint32(v298))) = uint16(v306)
							v310 = v298 + int32(3)
						} else {
							v310 = v298
						}
						v311 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v310))) = uint8(v311)
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v315 = m.ExcPending
						if v315 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v10)
							mBase = m.M
							v317 = m.ExcPending
							if v317 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v320 = m.ExcPending
								if v320 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					}
				case 3:
					v323 = F_JsonEncodeDateTime(m, v10, l0, int32(1114), int32(0))
					mBase = m.M
					v324 = m.ExcPending
					if v324 != 0 {
						return
					} else {
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v327 = m.ExcPending
						if v327 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v10)
							mBase = m.M
							v329 = m.ExcPending
							if v329 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v332 = m.ExcPending
								if v332 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					}
				case 4:
					v335 = F_JsonEncodeDateTime(m, v10, l0, int32(1184), int32(0))
					mBase = m.M
					v336 = m.ExcPending
					if v336 != 0 {
						return
					} else {
						F_appendStringInfoChar(m, l2, int32(34))
						mBase = m.M
						v339 = m.ExcPending
						if v339 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l2, v10)
							mBase = m.M
							v341 = m.ExcPending
							if v341 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l2, int32(34))
								mBase = m.M
								v344 = m.ExcPending
								if v344 != 0 {
									return
								} else {
									m.G0 = v10 + int32(176)
									return
								}
							}
						}
					}
				case 5:
					v345 = F_OidOutputFunctionCall(m, l4, l0)
					mBase = m.M
					v346 = m.ExcPending
					if v346 != 0 {
						return
					} else {
						F_appendStringInfoString(m, l2, v345)
						mBase = m.M
						v348 = m.ExcPending
						if v348 != 0 {
							return
						} else {
							F_pfree(m, v345)
							mBase = m.M
							v350 = m.ExcPending
							if v350 != 0 {
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
						v399 = base.I32_wrap_i64(l0)
						v400 = F_pg_detoast_datum_packed(m, v399)
						mBase = m.M
						v401 = m.ExcPending
						if v401 != 0 {
							return
						} else {
							v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400))))
							if v402 == int32(1) {
								v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400)+1)))
								if v408 == int32(18) {
									v411 = int32(16)
								} else {
									v411 = int32(0)
								}
								if base.Ui32((v408-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v418 = int32(4)
								} else {
									v418 = v411
								}
								v431 = v418
							} else {
								v419 = int32(1)
								if v402&v419 != 0 {
									v431 = int32(base.Ui32(v402)>>(uint(v419)%32)) - v419
								} else {
									v425 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
									v431 = int32(base.Ui32(v425)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							v432 = int32(1)
							if v402&v432 != 0 {
								v436 = v432
							} else {
								v436 = int32(4)
							}
							F_escape_json_with_len(m, l2, v400+v436, v431)
							mBase = m.M
							v439 = m.ExcPending
							if v439 != 0 {
								return
							} else {
								if v400 == v399 {
									m.G0 = v10 + int32(176)
									return
								} else {
									F_pfree(m, v400)
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
							v399 = base.I32_wrap_i64(l0)
							v400 = F_pg_detoast_datum_packed(m, v399)
							mBase = m.M
							v401 = m.ExcPending
							if v401 != 0 {
								return
							} else {
								v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400))))
								if v402 == int32(1) {
									v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400)+1)))
									if v408 == int32(18) {
										v411 = int32(16)
									} else {
										v411 = int32(0)
									}
									if base.Ui32((v408-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v418 = int32(4)
									} else {
										v418 = v411
									}
									v431 = v418
								} else {
									v419 = int32(1)
									if v402&v419 != 0 {
										v431 = int32(base.Ui32(v402)>>(uint(v419)%32)) - v419
									} else {
										v425 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
										v431 = int32(base.Ui32(v425)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								v432 = int32(1)
								if v402&v432 != 0 {
									v436 = v432
								} else {
									v436 = int32(4)
								}
								F_escape_json_with_len(m, l2, v400+v436, v431)
								mBase = m.M
								v439 = m.ExcPending
								if v439 != 0 {
									return
								} else {
									if v400 == v399 {
										m.G0 = v10 + int32(176)
										return
									} else {
										F_pfree(m, v400)
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
					v352 = F_OidFunctionCall1Coll(m, l4, int32(0), l0)
					mBase = m.M
					v353 = m.ExcPending
					if v353 != 0 {
						return
					} else {
						v355 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v352))
						mBase = m.M
						v356 = m.ExcPending
						if v356 != 0 {
							return
						} else {
							v357 = int32(1)
							v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355))))
							v361 = v359 & v357
							if v361 != 0 {
								v362 = v357
							} else {
								v362 = int32(4)
							}
							if v359 == int32(1) {
								v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355)+1)))
								if v369 == int32(18) {
									v372 = int32(16)
								} else {
									v372 = int32(0)
								}
								if base.Ui32((v369-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v379 = int32(4)
								} else {
									v379 = v372
								}
								v390 = v379
							} else {
								v380 = int32(1)
								if v361 != 0 {
									v390 = int32(base.Ui32(v359)>>(uint(v380)%32)) - v380
								} else {
									v384 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
									v390 = int32(base.Ui32(v384)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							F_appendBinaryStringInfo(m, l2, v355+v362, v390)
							mBase = m.M
							v392 = m.ExcPending
							if v392 != 0 {
								return
							} else {
								F_pfree(m, v355)
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
	var v22 int64
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
	var v48 int64
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
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
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
		v22 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
		*(*int64)(unsafe.Add(mBase, uint32(l4))) = v22
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)))
		*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v24)
		m.G0 = v12 + int32(32)
		return
	case 1:
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		if v26 == int32(0) {
			F_errstart_cold(m, int32(21), int32(_a_F_exec_eval_datum_0))
			mBase = m.M
			v167 = m.ExcPending
			if v167 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_exec_eval_datum_1), int32(0))
				mBase = m.M
				v171 = m.ExcPending
				if v171 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_exec_eval_datum_2), int32(_a_F_exec_eval_datum_3), int32(_a_F_exec_eval_datum_4))
					mBase = m.M
					v176 = m.ExcPending
					if v176 != 0 {
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
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
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
						v180 = m.ExcPending
						if v180 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_exec_eval_datum_6), int32(0))
							mBase = m.M
							v184 = m.ExcPending
							if v184 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_exec_eval_datum_2), int32(_a_F_exec_eval_datum_7), int32(_a_F_exec_eval_datum_4))
								mBase = m.M
								v189 = m.ExcPending
								if v189 != 0 {
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
							*(*int64)(unsafe.Add(mBase, uint32(l4))) = v48
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
			*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(0)
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
			*(*int64)(unsafe.Add(mBase, uint32(l4))) = base.I64_extend_i32_u(v72)
			*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(base.B2i32(v71 == int32(0)))
			v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			if v78 == int32(2249) {
				v152 = *(*int32)(unsafe.Add(mBase, uint32(v55)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v152
				v154 = *(*int32)(unsafe.Add(mBase, uint32(v55)+40))
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v154
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v78
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(-1)
			}
		}
		m.G0 = v12 + int32(32)
		return
	case 3:
		v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v89 = *(*int32)(unsafe.Add(mBase, uint32(v84+v85<<(uint(int32(2))%32))))
		v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+36))
		if v90 == int32(0) {
			F_instantiate_empty_record_variable(m, l0, v89)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return
			} else {
				v95 = *(*int32)(unsafe.Add(mBase, uint32(v89)+36))
				v96 = v95
				v97 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
				v98 = *(*int64)(unsafe.Add(mBase, uint32(v96)+48))
				if v97 != v98 {
					v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v103 = F_expanded_record_lookup_field(m, v96, v100, l1+int32(32))
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return
					} else {
						if v103 == int32(0) {
							F_errstart_cold(m, int32(21), int32(_a_F_exec_eval_datum_0))
							mBase = m.M
							v193 = m.ExcPending
							if v193 != 0 {
								return
							} else {
								F_errcode(m, int32(50360452))
								mBase = m.M
								v196 = m.ExcPending
								if v196 != 0 {
									return
								} else {
									v197 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
									v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v198
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v197
									F_errmsg(m, int32(_a_F_exec_eval_datum_8), v12+int32(16))
									mBase = m.M
									v205 = m.ExcPending
									if v205 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_exec_eval_datum_2), int32(_a_F_exec_eval_datum_9), int32(_a_F_exec_eval_datum_4))
										mBase = m.M
										v210 = m.ExcPending
										if v210 != 0 {
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
							v107 = *(*int64)(unsafe.Add(mBase, uint32(v96)+48))
							*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v107
							v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v109
							v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v111
							v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
							if v113 <= int32(0) {
								v135 = F_expanded_record_fetch_field(m, v96, v113, l5)
								mBase = m.M
								v136 = m.ExcPending
								if v136 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l4))) = v135
									m.G0 = v12 + int32(32)
									return
								}
							} else {
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
								if v116&int32(4) == int32(0) {
									v135 = F_expanded_record_fetch_field(m, v96, v113, l5)
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(l4))) = v135
										m.G0 = v12 + int32(32)
										return
									}
								} else {
									v121 = *(*int32)(unsafe.Add(mBase, uint32(v96)+64))
									if v121 < v113 {
										v135 = F_expanded_record_fetch_field(m, v96, v113, l5)
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(l4))) = v135
											m.G0 = v12 + int32(32)
											return
										}
									} else {
										v124 = v113 - int32(1)
										v125 = *(*int32)(unsafe.Add(mBase, uint32(v96)+60))
										v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124+v125))))
										*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v127)
										v129 = *(*int32)(unsafe.Add(mBase, uint32(v96)+56))
										v133 = *(*int64)(unsafe.Add(mBase, uint32(v129+v124<<(uint(int32(3))%32))))
										*(*int64)(unsafe.Add(mBase, uint32(l4))) = v133
										m.G0 = v12 + int32(32)
										return
									}
								}
							}
						}
					}
				} else {
					v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v109
					v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v111
					v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
					if v113 <= int32(0) {
						v135 = F_expanded_record_fetch_field(m, v96, v113, l5)
						mBase = m.M
						v136 = m.ExcPending
						if v136 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l4))) = v135
							m.G0 = v12 + int32(32)
							return
						}
					} else {
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
						if v116&int32(4) == int32(0) {
							v135 = F_expanded_record_fetch_field(m, v96, v113, l5)
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l4))) = v135
								m.G0 = v12 + int32(32)
								return
							}
						} else {
							v121 = *(*int32)(unsafe.Add(mBase, uint32(v96)+64))
							if v121 < v113 {
								v135 = F_expanded_record_fetch_field(m, v96, v113, l5)
								mBase = m.M
								v136 = m.ExcPending
								if v136 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l4))) = v135
									m.G0 = v12 + int32(32)
									return
								}
							} else {
								v124 = v113 - int32(1)
								v125 = *(*int32)(unsafe.Add(mBase, uint32(v96)+60))
								v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124+v125))))
								*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v127)
								v129 = *(*int32)(unsafe.Add(mBase, uint32(v96)+56))
								v133 = *(*int64)(unsafe.Add(mBase, uint32(v129+v124<<(uint(int32(3))%32))))
								*(*int64)(unsafe.Add(mBase, uint32(l4))) = v133
								m.G0 = v12 + int32(32)
								return
							}
						}
					}
				}
			}
		} else {
			v96 = v90
			v97 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
			v98 = *(*int64)(unsafe.Add(mBase, uint32(v96)+48))
			if v97 != v98 {
				v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v103 = F_expanded_record_lookup_field(m, v96, v100, l1+int32(32))
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return
				} else {
					if v103 == int32(0) {
						F_errstart_cold(m, int32(21), int32(_a_F_exec_eval_datum_0))
						mBase = m.M
						v193 = m.ExcPending
						if v193 != 0 {
							return
						} else {
							F_errcode(m, int32(50360452))
							mBase = m.M
							v196 = m.ExcPending
							if v196 != 0 {
								return
							} else {
								v197 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
								v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v198
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v197
								F_errmsg(m, int32(_a_F_exec_eval_datum_8), v12+int32(16))
								mBase = m.M
								v205 = m.ExcPending
								if v205 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_exec_eval_datum_2), int32(_a_F_exec_eval_datum_9), int32(_a_F_exec_eval_datum_4))
									mBase = m.M
									v210 = m.ExcPending
									if v210 != 0 {
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
						v107 = *(*int64)(unsafe.Add(mBase, uint32(v96)+48))
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v107
						v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v109
						v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v111
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						if v113 <= int32(0) {
							v135 = F_expanded_record_fetch_field(m, v96, v113, l5)
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l4))) = v135
								m.G0 = v12 + int32(32)
								return
							}
						} else {
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
							if v116&int32(4) == int32(0) {
								v135 = F_expanded_record_fetch_field(m, v96, v113, l5)
								mBase = m.M
								v136 = m.ExcPending
								if v136 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l4))) = v135
									m.G0 = v12 + int32(32)
									return
								}
							} else {
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v96)+64))
								if v121 < v113 {
									v135 = F_expanded_record_fetch_field(m, v96, v113, l5)
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(l4))) = v135
										m.G0 = v12 + int32(32)
										return
									}
								} else {
									v124 = v113 - int32(1)
									v125 = *(*int32)(unsafe.Add(mBase, uint32(v96)+60))
									v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124+v125))))
									*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v127)
									v129 = *(*int32)(unsafe.Add(mBase, uint32(v96)+56))
									v133 = *(*int64)(unsafe.Add(mBase, uint32(v129+v124<<(uint(int32(3))%32))))
									*(*int64)(unsafe.Add(mBase, uint32(l4))) = v133
									m.G0 = v12 + int32(32)
									return
								}
							}
						}
					}
				}
			} else {
				v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v109
				v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v111
				v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
				if v113 <= int32(0) {
					v135 = F_expanded_record_fetch_field(m, v96, v113, l5)
					mBase = m.M
					v136 = m.ExcPending
					if v136 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l4))) = v135
						m.G0 = v12 + int32(32)
						return
					}
				} else {
					v116 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
					if v116&int32(4) == int32(0) {
						v135 = F_expanded_record_fetch_field(m, v96, v113, l5)
						mBase = m.M
						v136 = m.ExcPending
						if v136 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l4))) = v135
							m.G0 = v12 + int32(32)
							return
						}
					} else {
						v121 = *(*int32)(unsafe.Add(mBase, uint32(v96)+64))
						if v121 < v113 {
							v135 = F_expanded_record_fetch_field(m, v96, v113, l5)
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l4))) = v135
								m.G0 = v12 + int32(32)
								return
							}
						} else {
							v124 = v113 - int32(1)
							v125 = *(*int32)(unsafe.Add(mBase, uint32(v96)+60))
							v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124+v125))))
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v127)
							v129 = *(*int32)(unsafe.Add(mBase, uint32(v96)+56))
							v133 = *(*int64)(unsafe.Add(mBase, uint32(v129+v124<<(uint(int32(3))%32))))
							*(*int64)(unsafe.Add(mBase, uint32(l4))) = v133
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
			v22 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
			*(*int64)(unsafe.Add(mBase, uint32(l4))) = v22
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)))
			*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v24)
			m.G0 = v12 + int32(32)
			return
		}
	default:
		F_errstart_cold(m, int32(21), int32(_a_F_exec_eval_datum_0))
		mBase = m.M
		v141 = m.ExcPending
		if v141 != 0 {
			return
		} else {
			v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v142
			F_errmsg_internal(m, int32(_a_F_exec_eval_datum_10), v12)
			mBase = m.M
			v146 = m.ExcPending
			if v146 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_exec_eval_datum_2), int32(_a_F_exec_eval_datum_11), int32(_a_F_exec_eval_datum_4))
				mBase = m.M
				v151 = m.ExcPending
				if v151 != 0 {
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
