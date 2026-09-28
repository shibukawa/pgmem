package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_varbit_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	v2 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v15 = v10 + int32(8)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v19 = F_palloc(m, v16+int32(1))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v22 = v16 - int32(8)
			if v22 < int32(0) {
				v89 = v19
				v91 = v15
				v92 = v2
			} else {
				v25 = v19
				v27 = v15
				v28 = v2
				for {
					v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
					v36 = int32(48)
					v37 = v33&int32(1) | v36
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+7)) = uint8(v37)
					if v33&int32(2) != 0 {
						v43 = int32(49)
					} else {
						v43 = v36
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+6)) = uint8(v43)
					if v33&int32(4) != 0 {
						v49 = int32(49)
					} else {
						v49 = int32(48)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+5)) = uint8(v49)
					if v33&int32(8) != 0 {
						v55 = int32(49)
					} else {
						v55 = int32(48)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+4)) = uint8(v55)
					if v33&int32(16) != 0 {
						v61 = int32(49)
					} else {
						v61 = int32(48)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+3)) = uint8(v61)
					if v33&int32(32) != 0 {
						v67 = int32(49)
					} else {
						v67 = int32(48)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+2)) = uint8(v67)
					if v33&int32(64) != 0 {
						v73 = int32(49)
					} else {
						v73 = int32(48)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)) = uint8(v73)
					if int32(0) <= base.I32_extend8_s(v33) {
						v80 = int32(48)
					} else {
						v80 = int32(49)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v80)
					v83 = v27 + int32(1)
					v84 = int32(8)
					v85 = v25 + v84
					v87 = v28 + v84
					if v87 <= v22 {
						v25 = v85
						v27 = v83
						v28 = v87
						continue
					} else {
						break
					}
					break
				}
				v89 = v85
				v91 = v83
				v92 = v87
			}
			if v16 <= v92 {
				v180 = v89
			} else {
				v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
				v101 = (v16 - v92) & int32(3)
				if v101 == int32(0) {
					v129 = v89
					v130 = v98
					v131 = v92
				} else {
					v105 = v89
					v106 = v98
					v107 = v92
					v110 = int32(0)
					for {
						if int32(0) <= base.I32_extend8_s(v106) {
							v118 = int32(48)
						} else {
							v118 = int32(49)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v105))) = uint8(v118)
						v120 = int32(1)
						v121 = v107 + v120
						v123 = v106 << (uint(v120) % 32)
						v125 = v105 + v120
						v127 = v110 + v120
						if v127 != v101 {
							v105 = v125
							v106 = v123
							v107 = v121
							v110 = v127
							continue
						} else {
							break
						}
						break
					}
					v129 = v125
					v130 = v123
					v131 = v121
				}
				if base.Ui32(int32(-4)) < base.Ui32(v92-v16) {
					v180 = v129
				} else {
					v140 = v129
					v141 = v130
					v142 = v131
					for {
						if v141&int32(16) != 0 {
							v152 = int32(49)
						} else {
							v152 = int32(48)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v140)+3)) = uint8(v152)
						if v141&int32(32) != 0 {
							v158 = int32(49)
						} else {
							v158 = int32(48)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v140)+2)) = uint8(v158)
						if v141&int32(64) != 0 {
							v164 = int32(49)
						} else {
							v164 = int32(48)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v140)+1)) = uint8(v164)
						if int32(0) <= base.I32_extend8_s(v141) {
							v171 = int32(48)
						} else {
							v171 = int32(49)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v171)
						v173 = int32(4)
						v176 = v140 + v173
						v178 = v142 + v173
						if v178 != v16 {
							v140 = v176
							v141 = v141 << (uint(v173) % 32)
							v142 = v178
							continue
						} else {
							break
						}
						break
					}
					v180 = v176
				}
			}
			v188 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v180))) = uint8(v188)
			return base.I64_extend_i32_u(v19)
		}
	}
}
func F_varbit_recv(m *base.Module, l0 int32) int64 {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pq_getmsgint(m, v12, int32(4))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		if base.Ui32(v14) < base.Ui32(int32(2147483641)) {
			if base.B2i32(v11 < v14)&base.B2i32(int32(0) < v11) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(16777346))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
						F_errmsg(m, int32(_a_F_varbit_recv_0), v9)
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_varbit_recv_1), int32(662), int32(_a_F_varbit_recv_2))
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v27 = int32(base.Ui32(v14+int32(7)) >> (uint(int32(3)) % 32))
				v29 = v27 + int32(8)
				v30 = F_palloc(m, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v14
					*(*int32)(unsafe.Add(mBase, uint32(v30))) = v29 << (uint(int32(2)) % 32)
					F_pq_copymsgbytes(m, v12, v30+int32(8), v27)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int64(0)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
						v42 = int32(base.Ui32(v40) >> (uint(int32(2)) % 32))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
						v48 = v42<<(uint(int32(3))%32) - v45 + int32(-64)
						if int32(0) < v48 {
							v53 = v30 + v42 - int32(1)
							v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
							v57 = v54 & (int32(255) << (uint(v48) % 32))
							*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v57)
						} else {
						}
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v30)
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50462850))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_varbit_recv_3), int32(0))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_varbit_recv_1), int32(652), int32(_a_F_varbit_recv_2))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int64(0)
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
}
func F_varchar_support(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int64
	_ = v40
	v4 = int64(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v6 != int32(463) {
		v40 = v4
		return v40
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v13 != int32(7) {
			v40 = v4
			return v40
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+32)))
			if v16 != 0 {
				v40 = v4
				return v40
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v18 = F_exprTypmod(m, v17)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
					v23 = int32(0)
					v27 = int32(4)
					if base.B2i32(v23 <= v22)&(base.B2i32(v18 < v23)|base.B2i32(v22-v27 < v18-v27)) != 0 {
						v40 = v4
						return v40
					} else {
						v34 = F_relabel_to_typmod(m, v17, v22)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							v40 = base.I64_extend_i32_u(v34)
							return v40
						}
					}
				}
			}
		}
	}
}
func F_varstr_abbrev_convert(m *base.Module, l0 int64, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int64
	_ = v643
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v721 int64
	_ = v721
	var v732 int64
	_ = v732
	var v736 int32
	_ = v736
	var v740 int64
	_ = v740
	var v742 int64
	_ = v742
	var v744 int64
	_ = v744
	var v747 int64
	_ = v747
	var v749 int64
	_ = v749
	var v751 int64
	_ = v751
	var v753 int64
	_ = v753
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v18 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(l0))
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
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(0)
	v25 = int32(1)
	v28 = v22 & v25
	if v28 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v29 = v25
	goto L5
L4:
	;
	v29 = int32(4)
	goto L5
L5:
	;
	if v22 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v57 = v29 + v18
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	if v58 == int32(1042) {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v35 == int32(18) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v46 = int32(1)
	if v28 != 0 {
		v56 = int32(base.Ui32(v22)>>(uint(v46)%32)) - v46
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v38 = int32(16)
	goto L12
L11:
	;
	v38 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v35-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v45 = int32(4)
	goto L15
L14:
	;
	v45 = v38
	goto L15
L15:
	;
	v56 = v45
	goto L6
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v56 = int32(base.Ui32(v50)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v63 = int32(-1)
	v65 = v56 - int32(1)
	if v63 <= v65 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v84 = v56
	goto L19
L19:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+29)))
	if v85 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L20:
	;
	v84 = v83
	goto L19
L21:
	;
	v68 = v63
	goto L23
L22:
	;
	v68 = v65
	goto L23
L23:
	;
	v72 = v56
	goto L24
L24:
	;
	v76 = v72 - int32(1)
	if v76 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v83 = v72
	goto L20
L26:
	;
	v83 = v68 + int32(1)
	goto L20
L27:
	;
	goto L28
L28:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+v76))))
	if v80 == int32(32) {
		v72 = v76
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	if base.I64_extend_i32_u(v18) != l0 {
		goto L190
	} else {
		goto L191
	}
L31:
	;
	v712 = int32(8)
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	if base.Ui32(v712) <= base.Ui32(v713) {
		goto L184
	} else {
		goto L185
	}
L32:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v88 <= v84 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v290 = v57
	v291 = v84
	goto L34
L34:
	;
	v298 = int32(8)
	if base.Ui32(v298) <= base.Ui32(v291) {
		goto L110
	} else {
		goto L111
	}
L35:
	;
	v90 = int32(1)
	v91 = v84 + v90
	v92 = int32(1073741823)
	v94 = v88 << (uint(v90) % 32)
	if base.Ui32(v92) <= base.Ui32(v94) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v84 != v107 {
		goto L46
	} else {
		goto L47
	}
L38:
	;
	v97 = v92
	goto L40
L39:
	;
	v97 = v94
	goto L40
L40:
	;
	if base.Ui32(v97) < base.Ui32(v91) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v99 = v91
	goto L43
L42:
	;
	v99 = v97
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v102 = F_repalloc(m, v101, v99)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v102
	goto L37
L45:
	;
	if v84 != 0 {
		goto L69
	} else {
		goto L70
	}
L46:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v178 = v109
	goto L45
L47:
	;
	goto L48
L48:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+28)))
	if v111 != int32(1) {
		v178 = v110
		goto L45
	} else {
		goto L49
	}
L49:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v84) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	if v175 == int32(0) {
		goto L31
	} else {
		goto L68
	}
L51:
	;
	v175 = int32(0)
	goto L50
L52:
	;
	v149 = v144
	v150 = v145
	v151 = v146
	goto L62
L53:
	;
	if (v110|v57)&int32(3) != 0 {
		v144 = v110
		v145 = v57
		v146 = v84
		goto L52
	} else {
		goto L56
	}
L54:
	;
	v137 = v110
	v138 = v57
	v139 = v84
	goto L55
L55:
	;
	if v139 == int32(0) {
		goto L51
	} else {
		goto L61
	}
L56:
	;
	v121 = v110
	v122 = v57
	v123 = v84
	goto L57
L57:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v126 != v127 {
		v144 = v121
		v145 = v122
		v146 = v123
		goto L52
	} else {
		goto L59
	}
L58:
	;
	v137 = v132
	v138 = v130
	v139 = v134
	goto L55
L59:
	;
	v129 = int32(4)
	v130 = v122 + v129
	v132 = v121 + v129
	v134 = v123 - v129
	if base.Ui32(int32(3)) < base.Ui32(v134) {
		v121 = v132
		v122 = v130
		v123 = v134
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v144 = v137
	v145 = v138
	v146 = v139
	goto L52
L62:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	if v154 == v155 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v175 = v154 - v155
	goto L50
L64:
	;
	v157 = int32(1)
	v162 = v151 - v157
	if v162 != 0 {
		v149 = v149 + v157
		v150 = v150 + v157
		v151 = v162
		goto L62
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	goto L63
L67:
	;
	goto L51
L68:
	;
	v178 = v110
	goto L45
L69:
	;
	base.MemoryCopy(m, v178, v57, v84)
	goto L71
L70:
	;
	goto L71
L71:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v182 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v180+v84))) = uint8(v182)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v84
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v16)+96))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v186 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v290 = v286
	v291 = v279
	goto L34
L73:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+16))
	v191 = base.B2i32(v187 != int32(0))
	goto L75
L74:
	;
	v191 = int32(1)
	goto L75
L75:
	;
	if v191 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v16)+96))
	v198 = F_pg_strxfrm(m, v194, v195, v196, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if base.Ui32(int32(8)) <= base.Ui32(v237) {
		goto L93
	} else {
		goto L94
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v198
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if base.Ui32(v198) < base.Ui32(v201) {
		v279 = v198
		goto L72
	} else {
		goto L80
	}
L80:
	;
	v206 = v201
	v207 = v198
	goto L81
L81:
	;
	v214 = int32(1)
	v215 = v207 + v214
	v216 = int32(1073741823)
	v218 = v206 << (uint(v214) % 32)
	if base.Ui32(v216) <= base.Ui32(v218) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v279 = v232
	goto L72
L83:
	;
	v221 = v216
	goto L85
L84:
	;
	v221 = v218
	goto L85
L85:
	;
	if base.Ui32(v221) < base.Ui32(v215) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v223 = v215
	goto L88
L87:
	;
	v223 = v221
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v223
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v226 = F_repalloc(m, v225, v223)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v226
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v16)+96))
	v232 = F_pg_strxfrm(m, v226, v229, v230, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v232
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if base.Ui32(v235) <= base.Ui32(v232) {
		v206 = v235
		v207 = v232
		goto L81
	} else {
		goto L91
	}
L91:
	;
	goto L82
L92:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v16)+96))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	if v256 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L93:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v253 = v240
	goto L92
L94:
	;
	goto L95
L95:
	;
	v241 = int32(4)
	if base.Ui32(v237) <= base.Ui32(v241) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v244 = v241
	goto L98
L97:
	;
	v244 = v237
	goto L98
L98:
	;
	v246 = v244 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v246
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v249 = F_repalloc(m, v248, v246)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v249
	v253 = v249
	goto L92
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v273
	v279 = v273
	goto L72
L101:
	;
	v273 = v271
	goto L100
L102:
	;
	v259 = int32(8)
	v260 = F_strlen(m, v254)
	mBase = m.M
	if base.Ui32(v259) <= base.Ui32(v260) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L104
L104:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v256)+20))
	v269 = m.T0[v268].(func(*base.Module, int32, int32, int32, int32) int32)(m, v253, int32(8), v254, v255)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L109
	}
L105:
	;
	v263 = v259
	goto L107
L106:
	;
	v263 = v260
	goto L107
L107:
	;
	if v263 == int32(0) {
		v271 = v263
		goto L101
	} else {
		goto L108
	}
L108:
	;
	base.MemoryCopy(m, v253, v254, v263)
	v273 = v263
	goto L100
L109:
	;
	v271 = v269
	goto L101
L110:
	;
	v301 = v298
	goto L112
L111:
	;
	v301 = v291
	goto L112
L112:
	;
	if v301 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	base.MemoryCopy(m, v14+int32(8), v290, v301)
	goto L115
L114:
	;
	goto L115
L115:
	;
	v305 = int32(128)
	if v305 <= v84 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v308 = v305
	goto L118
L117:
	;
	v308 = v84
	goto L118
L118:
	;
	v314 = v308 - int32(1636608432)
	if v57&int32(3) != 0 {
		goto L123
	} else {
		goto L124
	}
L119:
	;
	v574 = v16 - int32(-64)
	if int32(129) <= v84 {
		goto L159
	} else {
		goto L160
	}
L120:
	;
	v546 = int32(14)
	v548 = v542 ^ v543 - base.I32_rotl(v542, v546)
	v552 = v548 ^ v541 - base.I32_rotl(v548, int32(11))
	v556 = v552 ^ v542 - base.I32_rotl(v552, int32(25))
	v560 = v556 ^ v548 - base.I32_rotl(v556, int32(16))
	v564 = v560 ^ v552 - base.I32_rotl(v560, int32(4))
	v568 = v564 ^ v556 - base.I32_rotl(v564, v546)
	v572 = v568 ^ v560 - base.I32_rotl(v568, int32(24))
	goto L119
L121:
	;
	switch v472 - int32(1) {
	case 0:
		v534 = v473
		v535 = v474
		v536 = v475
		goto L148
	case 1:
		v527 = v473
		v528 = v474
		v529 = v475
		goto L149
	case 2:
		v520 = v473
		v521 = v474
		v522 = v475
		goto L150
	case 3:
		v514 = v474
		v515 = v475
		goto L151
	case 4:
		v510 = v474
		v511 = v475
		goto L152
	case 5:
		v504 = v474
		v505 = v475
		goto L153
	case 6:
		v498 = v474
		v499 = v475
		goto L154
	case 7:
		v493 = v475
		goto L155
	case 8:
		v488 = v475
		goto L156
	case 9:
		v483 = v475
		goto L157
	case 10:
		goto L158
	default:
		v541 = v473
		v542 = v474
		v543 = v475
		goto L120
	}
L122:
	;
	v423 = v57
	v424 = v308
	v425 = v314
	v426 = v314
	v427 = v314
	goto L145
L123:
	;
	if base.Ui32(int32(11)) < base.Ui32(v308) {
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	if base.Ui32(v308) < base.Ui32(int32(12)) {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	v471 = v57
	v472 = v308
	v473 = v314
	v474 = v314
	v475 = v314
	goto L121
L127:
	;
	switch v370 - int32(1) {
	case 0:
		v420 = v371
		goto L134
	case 1:
		v415 = v371
		goto L135
	case 2:
		goto L136
	case 3:
		v408 = v372
		goto L137
	case 4:
		v405 = v372
		goto L138
	case 5:
		v400 = v372
		goto L139
	case 6:
		goto L140
	case 7:
		v391 = v373
		goto L141
	case 8:
		v386 = v373
		goto L142
	case 9:
		v381 = v373
		goto L143
	case 10:
		goto L144
	default:
		v541 = v371
		v542 = v372
		v543 = v373
		goto L120
	}
L128:
	;
	v369 = v57
	v370 = v308
	v371 = v314
	v372 = v314
	v373 = v314
	goto L127
L129:
	;
	goto L130
L130:
	;
	v321 = v57
	v322 = v308
	v323 = v314
	v324 = v314
	v325 = v314
	goto L131
L131:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	v328 = v327 + v324
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v321)+8))
	v332 = v331 + v325
	v334 = int32(4)
	v336 = v329 + v323 - v332 ^ base.I32_rotl(v332, v334)
	v340 = v328 - v336 ^ base.I32_rotl(v336, int32(6))
	v341 = v332 + v328
	v342 = v336 + v341
	v343 = v340 + v342
	v347 = v341 - v340 ^ base.I32_rotl(v340, int32(8))
	v351 = v342 - v347 ^ base.I32_rotl(v347, int32(16))
	v355 = v343 - v351 ^ base.I32_rotl(v351, int32(19))
	v356 = v347 + v343
	v357 = v351 + v356
	v358 = v355 + v357
	v362 = v356 - v355 ^ base.I32_rotl(v355, v334)
	v363 = int32(12)
	v364 = v321 + v363
	v366 = v322 - v363
	if base.Ui32(int32(11)) < base.Ui32(v366) {
		v321 = v364
		v322 = v366
		v323 = v357
		v324 = v358
		v325 = v362
		goto L131
	} else {
		goto L133
	}
L132:
	;
	v369 = v364
	v370 = v366
	v371 = v357
	v372 = v358
	v373 = v362
	goto L127
L133:
	;
	goto L132
L134:
	;
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369))))
	v541 = v420 + v421
	v542 = v372
	v543 = v373
	goto L120
L135:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369)+1)))
	v420 = v416<<(uint(int32(8))%32) + v415
	goto L134
L136:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369)+2)))
	v415 = v411<<(uint(int32(16))%32) + v371
	goto L135
L137:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	v541 = v409 + v371
	v542 = v408
	v543 = v373
	goto L120
L138:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369)+4)))
	v408 = v405 + v406
	goto L137
L139:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369)+5)))
	v405 = v401<<(uint(int32(8))%32) + v400
	goto L138
L140:
	;
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369)+6)))
	v400 = v396<<(uint(int32(16))%32) + v372
	goto L139
L141:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	v541 = v392 + v371
	v542 = v394 + v372
	v543 = v391
	goto L120
L142:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369)+8)))
	v391 = v387<<(uint(int32(8))%32) + v386
	goto L141
L143:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369)+9)))
	v386 = v382<<(uint(int32(16))%32) + v381
	goto L142
L144:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369)+10)))
	v381 = v377<<(uint(int32(24))%32) + v373
	goto L143
L145:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v423)+4))
	v430 = v429 + v426
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v423)))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v423)+8))
	v434 = v433 + v427
	v436 = int32(4)
	v438 = v431 + v425 - v434 ^ base.I32_rotl(v434, v436)
	v442 = v430 - v438 ^ base.I32_rotl(v438, int32(6))
	v443 = v434 + v430
	v444 = v438 + v443
	v445 = v442 + v444
	v449 = v443 - v442 ^ base.I32_rotl(v442, int32(8))
	v453 = v444 - v449 ^ base.I32_rotl(v449, int32(16))
	v457 = v445 - v453 ^ base.I32_rotl(v453, int32(19))
	v458 = v449 + v445
	v459 = v453 + v458
	v460 = v457 + v459
	v464 = v458 - v457 ^ base.I32_rotl(v457, v436)
	v465 = int32(12)
	v466 = v423 + v465
	v468 = v424 - v465
	if base.Ui32(int32(11)) < base.Ui32(v468) {
		v423 = v466
		v424 = v468
		v425 = v459
		v426 = v460
		v427 = v464
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v471 = v466
	v472 = v468
	v473 = v459
	v474 = v460
	v475 = v464
	goto L121
L147:
	;
	goto L146
L148:
	;
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471))))
	v541 = v534 + v537
	v542 = v535
	v543 = v536
	goto L120
L149:
	;
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471)+1)))
	v534 = v530<<(uint(int32(8))%32) + v527
	v535 = v528
	v536 = v529
	goto L148
L150:
	;
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471)+2)))
	v527 = v523<<(uint(int32(16))%32) + v520
	v528 = v521
	v529 = v522
	goto L149
L151:
	;
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471)+3)))
	v520 = v516<<(uint(int32(24))%32) + v473
	v521 = v514
	v522 = v515
	goto L150
L152:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471)+4)))
	v514 = v510 + v512
	v515 = v511
	goto L151
L153:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471)+5)))
	v510 = v506<<(uint(int32(8))%32) + v504
	v511 = v505
	goto L152
L154:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471)+6)))
	v504 = v500<<(uint(int32(16))%32) + v498
	v505 = v499
	goto L153
L155:
	;
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471)+7)))
	v498 = v494<<(uint(int32(24))%32) + v474
	v499 = v493
	goto L154
L156:
	;
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471)+8)))
	v493 = v489<<(uint(int32(8))%32) + v488
	goto L155
L157:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471)+9)))
	v488 = v484<<(uint(int32(16))%32) + v483
	goto L156
L158:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471)+10)))
	v483 = v479<<(uint(int32(24))%32) + v475
	goto L157
L159:
	;
	v581 = int32(711645284)
	v584 = v84 - int32(1636608428) ^ v581 - int32(1455628627)
	v589 = v584 ^ int32(-1636608428) - base.I32_rotl(v584, int32(25))
	v594 = v589 ^ v581 - base.I32_rotl(v589, int32(16))
	v598 = v594 ^ v584 - base.I32_rotl(v594, int32(4))
	v602 = v598 ^ v589 - base.I32_rotl(v598, int32(14))
	goto L162
L160:
	;
	v608 = v572
	goto L161
L161:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v574)+16))
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574))))
	v614 = int32(32) - v613
	v616 = v611 + int32(base.Ui32(v608)>>(uint(v614)%32))
	v617 = v608 << (uint(v613) % 32)
	if v617 != 0 {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	v608 = v602 ^ v594 - base.I32_rotl(v602, int32(24)) ^ v572
	goto L161
L163:
	;
	v642 = v16 + int32(40)
	v643 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v652 = int32(711645284)
	v655 = base.I32_wrap_i64(int64(base.Ui64(v643)>>(uint(int64(32))%64))^v643) - int32(1636608428) ^ v652 - int32(1455628627)
	v660 = v655 ^ int32(-1636608428) - base.I32_rotl(v655, int32(25))
	v665 = v660 ^ v652 - base.I32_rotl(v660, int32(16))
	v669 = v665 ^ v655 - base.I32_rotl(v665, int32(4))
	v673 = v669 ^ v660 - base.I32_rotl(v669, int32(14))
	v677 = v673 ^ v665 - base.I32_rotl(v673, int32(24))
	goto L173
L164:
	;
	v624 = int32(32) - (base.I32_clz(v617) ^ int32(31))
	v625 = int32(255)
	if base.Ui32(v614&v625) < base.Ui32(v624&v625) {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	v634 = v614 + int32(1)
	goto L166
L166:
	;
	v636 = v634 & int32(255)
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v616))))
	if base.Ui32(v637) < base.Ui32(v636) {
		goto L170
	} else {
		goto L171
	}
L167:
	;
	v630 = v614 + int32(1)
	goto L169
L168:
	;
	v630 = v624
	goto L169
L169:
	;
	v634 = v630
	goto L166
L170:
	;
	v639 = v636
	goto L172
L171:
	;
	v639 = v637
	goto L172
L172:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v616))) = uint8(v639)
	goto L163
L173:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v642)+16))
	v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642))))
	v683 = int32(32) - v682
	v685 = v680 + int32(base.Ui32(v677)>>(uint(v683)%32))
	v686 = v677 << (uint(v682) % 32)
	if v686 != 0 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v710 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+28)) = uint8(v710)
	v732 = v643
	goto L30
L175:
	;
	v693 = int32(32) - (base.I32_clz(v686) ^ int32(31))
	v694 = int32(255)
	if base.Ui32(v683&v694) < base.Ui32(v693&v694) {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	v703 = v683 + int32(1)
	goto L177
L177:
	;
	v705 = v703 & int32(255)
	v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v685))))
	if base.Ui32(v706) < base.Ui32(v705) {
		goto L181
	} else {
		goto L182
	}
L178:
	;
	v699 = v683 + int32(1)
	goto L180
L179:
	;
	v699 = v693
	goto L180
L180:
	;
	v703 = v699
	goto L177
L181:
	;
	v708 = v705
	goto L183
L182:
	;
	v708 = v706
	goto L183
L183:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v685))) = uint8(v708)
	goto L174
L184:
	;
	v716 = v712
	goto L186
L185:
	;
	v716 = v713
	goto L186
L186:
	;
	if v716 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	base.MemoryCopy(m, v14+int32(8), v719, v716)
	goto L189
L188:
	;
	goto L189
L189:
	;
	v721 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v732 = v721
	goto L30
L190:
	;
	F_pfree(m, v18)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	m.G0 = v14 + int32(16)
	v740 = int64(56)
	v742 = int64(65280)
	v744 = int64(40)
	v747 = int64(16711680)
	v749 = int64(24)
	v751 = int64(4278190080)
	v753 = int64(8)
	return v732<<(uint(v740)%64) | v732&v742<<(uint(v744)%64) | (v732&v747<<(uint(v749)%64) | v732&v751<<(uint(v753)%64)) | (int64(base.Ui64(v732)>>(uint(v753)%64))&v751 | int64(base.Ui64(v732)>>(uint(v749)%64))&v747 | (int64(base.Ui64(v732)>>(uint(v744)%64))&v742 | int64(base.Ui64(v732)>>(uint(v740)%64))))
L193:
	;
	goto L192
}
func F_varstr_cmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
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
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	if l4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = F_pg_newlocale_from_collation(m, l4)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
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
	v225 = m.ExcPending
	if v225 != 0 {
		goto L5
	} else {
		goto L80
	}
L4:
	;
	return v220
L5:
	;
	return int32(0)
L6:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
	if v11 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v14 = base.B2i32(l1 < l3)
	if l1 < l3 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	if l1 != l3 {
		goto L32
	} else {
		goto L33
	}
L10:
	;
	v15 = l1
	goto L12
L11:
	;
	v15 = l3
	goto L12
L12:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v15) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	if v77 != 0 {
		v220 = v77
		goto L4
	} else {
		goto L31
	}
L14:
	;
	v77 = int32(0)
	goto L13
L15:
	;
	v51 = v46
	v52 = v47
	v53 = v48
	goto L25
L16:
	;
	if (l0|l2)&int32(3) != 0 {
		v46 = l0
		v47 = l2
		v48 = v15
		goto L15
	} else {
		goto L19
	}
L17:
	;
	v39 = l0
	v40 = l2
	v41 = v15
	goto L18
L18:
	;
	if v41 == int32(0) {
		goto L14
	} else {
		goto L24
	}
L19:
	;
	v23 = l0
	v24 = l2
	v25 = v15
	goto L20
L20:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v28 != v29 {
		v46 = v23
		v47 = v24
		v48 = v25
		goto L15
	} else {
		goto L22
	}
L21:
	;
	v39 = v34
	v40 = v32
	v41 = v36
	goto L18
L22:
	;
	v31 = int32(4)
	v32 = v24 + v31
	v34 = v23 + v31
	v36 = v25 - v31
	if base.Ui32(int32(3)) < base.Ui32(v36) {
		v23 = v34
		v24 = v32
		v25 = v36
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v46 = v39
	v47 = v40
	v48 = v41
	goto L15
L25:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v56 == v57 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v77 = v56 - v57
	goto L13
L27:
	;
	v59 = int32(1)
	v64 = v53 - v59
	if v64 != 0 {
		v51 = v51 + v59
		v52 = v52 + v59
		v53 = v64
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	goto L14
L31:
	;
	return base.B2i32(l3 < l1) - v14
L32:
	;
	v146 = F_pg_strncoll(m, l0, l1, l2, l3, v7)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L53
	}
L33:
	;
	if base.Ui32(int32(4)) <= base.Ui32(l1) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	if v143 != 0 {
		goto L32
	} else {
		goto L52
	}
L35:
	;
	v143 = int32(0)
	goto L34
L36:
	;
	v117 = v112
	v118 = v113
	v119 = v114
	goto L46
L37:
	;
	if (l0|l2)&int32(3) != 0 {
		v112 = l0
		v113 = l2
		v114 = l1
		goto L36
	} else {
		goto L40
	}
L38:
	;
	v105 = l0
	v106 = l2
	v107 = l1
	goto L39
L39:
	;
	if v107 == int32(0) {
		goto L35
	} else {
		goto L45
	}
L40:
	;
	v89 = l0
	v90 = l2
	v91 = l1
	goto L41
L41:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v94 != v95 {
		v112 = v89
		v113 = v90
		v114 = v91
		goto L36
	} else {
		goto L43
	}
L42:
	;
	v105 = v100
	v106 = v98
	v107 = v102
	goto L39
L43:
	;
	v97 = int32(4)
	v98 = v90 + v97
	v100 = v89 + v97
	v102 = v91 - v97
	if base.Ui32(int32(3)) < base.Ui32(v102) {
		v89 = v100
		v90 = v98
		v91 = v102
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v112 = v105
	v113 = v106
	v114 = v107
	goto L36
L46:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v122 == v123 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v143 = v122 - v123
	goto L34
L48:
	;
	v125 = int32(1)
	v130 = v119 - v125
	if v130 != 0 {
		v117 = v117 + v125
		v118 = v118 + v125
		v119 = v130
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	goto L35
L52:
	;
	return int32(0)
L53:
	;
	if v146 != 0 {
		v220 = v146
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v148 != int32(1) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	return int32(0)
L56:
	;
	goto L57
L57:
	;
	v153 = base.B2i32(l1 < l3)
	if l1 < l3 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v154 = l1
	goto L60
L59:
	;
	v154 = l3
	goto L60
L60:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v154) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	if v216 != 0 {
		v220 = v216
		goto L4
	} else {
		goto L79
	}
L62:
	;
	v216 = int32(0)
	goto L61
L63:
	;
	v190 = v185
	v191 = v186
	v192 = v187
	goto L73
L64:
	;
	if (l0|l2)&int32(3) != 0 {
		v185 = l0
		v186 = l2
		v187 = v154
		goto L63
	} else {
		goto L67
	}
L65:
	;
	v178 = l0
	v179 = l2
	v180 = v154
	goto L66
L66:
	;
	if v180 == int32(0) {
		goto L62
	} else {
		goto L72
	}
L67:
	;
	v162 = l0
	v163 = l2
	v164 = v154
	goto L68
L68:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	if v167 != v168 {
		v185 = v162
		v186 = v163
		v187 = v164
		goto L63
	} else {
		goto L70
	}
L69:
	;
	v178 = v173
	v179 = v171
	v180 = v175
	goto L66
L70:
	;
	v170 = int32(4)
	v171 = v163 + v170
	v173 = v162 + v170
	v175 = v164 - v170
	if base.Ui32(int32(3)) < base.Ui32(v175) {
		v162 = v173
		v163 = v171
		v164 = v175
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v185 = v178
	v186 = v179
	v187 = v180
	goto L63
L73:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	if v195 == v196 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v216 = v195 - v196
	goto L61
L75:
	;
	v198 = int32(1)
	v203 = v192 - v198
	if v203 != 0 {
		v190 = v190 + v198
		v191 = v191 + v198
		v192 = v203
		goto L73
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	goto L74
L78:
	;
	goto L62
L79:
	;
	v220 = base.B2i32(l3 < l1) - v153
	goto L4
L80:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	F_errmsg(m, int32(_a_F_varstr_cmp_0), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	F_errhint(m, int32(_a_F_varstr_cmp_1), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_varstr_cmp_2), int32(1337), int32(_a_F_varstr_cmp_3))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_varstrfastcmp_locale(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
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
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
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
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if l1 != l3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	if v74 == int32(1042) {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	if base.Ui32(int32(4)) <= base.Ui32(l1) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	if v71 != 0 {
		goto L1
	} else {
		goto L21
	}
L4:
	;
	v71 = int32(0)
	goto L3
L5:
	;
	v45 = v40
	v46 = v41
	v47 = v42
	goto L15
L6:
	;
	if (l0|l2)&int32(3) != 0 {
		v40 = l0
		v41 = l2
		v42 = l1
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v33 = l0
	v34 = l2
	v35 = l1
	goto L8
L8:
	;
	if v35 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L9:
	;
	v17 = l0
	v18 = l2
	v19 = l1
	goto L10
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v22 != v23 {
		v40 = v17
		v41 = v18
		v42 = v19
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v33 = v28
	v34 = v26
	v35 = v30
	goto L8
L12:
	;
	v25 = int32(4)
	v26 = v18 + v25
	v28 = v17 + v25
	v30 = v19 - v25
	if base.Ui32(int32(3)) < base.Ui32(v30) {
		v17 = v28
		v18 = v26
		v19 = v30
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v40 = v33
	v41 = v34
	v42 = v35
	goto L5
L15:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 == v51 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v71 = v50 - v51
	goto L3
L17:
	;
	v53 = int32(1)
	v58 = v47 - v53
	if v58 != 0 {
		v45 = v45 + v53
		v46 = v46 + v53
		v47 = v58
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L4
L21:
	;
	return int32(0)
L22:
	;
	v79 = int32(-1)
	v81 = l1 - int32(1)
	if v79 <= v81 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v123 = l1
	v124 = l3
	goto L24
L24:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v125 <= v123 {
		goto L45
	} else {
		goto L46
	}
L25:
	;
	v102 = int32(-1)
	v104 = l3 - int32(1)
	if v102 <= v104 {
		goto L36
	} else {
		goto L37
	}
L26:
	;
	v84 = v79
	goto L28
L27:
	;
	v84 = v81
	goto L28
L28:
	;
	v88 = l1
	goto L29
L29:
	;
	v92 = v88 - int32(1)
	if v92 < int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v99 = v88
	goto L25
L31:
	;
	v99 = v84 + int32(1)
	goto L25
L32:
	;
	goto L33
L33:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v92))))
	if v96 == int32(32) {
		v88 = v92
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	v123 = v99
	v124 = v122
	goto L24
L36:
	;
	v107 = v102
	goto L38
L37:
	;
	v107 = v104
	goto L38
L38:
	;
	v111 = l3
	goto L39
L39:
	;
	v115 = v111 - int32(1)
	if v115 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v122 = v111
	goto L35
L41:
	;
	v122 = v107 + int32(1)
	goto L35
L42:
	;
	goto L43
L43:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v115))))
	if v119 == int32(32) {
		v111 = v115
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	v127 = int32(1)
	v128 = v123 + v127
	v129 = int32(1073741823)
	v131 = v125 << (uint(v127) % 32)
	if base.Ui32(v129) <= base.Ui32(v131) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v146 <= v124 {
		goto L56
	} else {
		goto L57
	}
L48:
	;
	v134 = v129
	goto L50
L49:
	;
	v134 = v131
	goto L50
L50:
	;
	if base.Ui32(v134) < base.Ui32(v128) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v136 = v128
	goto L53
L52:
	;
	v136 = v134
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v139 = F_repalloc(m, v138, v136)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	return int32(0)
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v139
	goto L47
L56:
	;
	v148 = int32(1)
	v149 = v124 + v148
	v150 = int32(1073741823)
	v152 = v146 << (uint(v148) % 32)
	if base.Ui32(v150) <= base.Ui32(v152) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v123 != v166 {
		goto L67
	} else {
		goto L68
	}
L59:
	;
	v155 = v150
	goto L61
L60:
	;
	v155 = v152
	goto L61
L61:
	;
	if base.Ui32(v155) < base.Ui32(v149) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v157 = v149
	goto L64
L63:
	;
	v157 = v155
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v160 = F_repalloc(m, v159, v157)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L54
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v160
	goto L58
L66:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v240 == v124 {
		goto L93
	} else {
		goto L94
	}
L67:
	;
	if v123 != 0 {
		goto L88
	} else {
		goto L89
	}
L68:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v123) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	if v229 != 0 {
		goto L67
	} else {
		goto L87
	}
L70:
	;
	v229 = int32(0)
	goto L69
L71:
	;
	v203 = v198
	v204 = v199
	v205 = v200
	goto L81
L72:
	;
	if (v165|l0)&int32(3) != 0 {
		v198 = v165
		v199 = l0
		v200 = v123
		goto L71
	} else {
		goto L75
	}
L73:
	;
	v191 = v165
	v192 = l0
	v193 = v123
	goto L74
L74:
	;
	if v193 == int32(0) {
		goto L70
	} else {
		goto L80
	}
L75:
	;
	v175 = v165
	v176 = l0
	v177 = v123
	goto L76
L76:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	if v180 != v181 {
		v198 = v175
		v199 = v176
		v200 = v177
		goto L71
	} else {
		goto L78
	}
L77:
	;
	v191 = v186
	v192 = v184
	v193 = v188
	goto L74
L78:
	;
	v183 = int32(4)
	v184 = v176 + v183
	v186 = v175 + v183
	v188 = v177 - v183
	if base.Ui32(int32(3)) < base.Ui32(v188) {
		v175 = v186
		v176 = v184
		v177 = v188
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v198 = v191
	v199 = v192
	v200 = v193
	goto L71
L81:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v208 == v209 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v229 = v208 - v209
	goto L69
L83:
	;
	v211 = int32(1)
	v216 = v205 - v211
	if v216 != 0 {
		v203 = v203 + v211
		v204 = v204 + v211
		v205 = v216
		goto L81
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	goto L82
L86:
	;
	goto L70
L87:
	;
	v238 = int32(1)
	goto L66
L88:
	;
	base.MemoryCopy(m, v165, l0, v123)
	goto L90
L89:
	;
	goto L90
L90:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v234 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v123))) = uint8(v234)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v123
	v238 = v234
	goto L66
L91:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+4))
	if v321 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L92:
	;
	if v238 == int32(0) {
		v318 = v239
		goto L91
	} else {
		goto L118
	}
L93:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v124) {
		goto L99
	} else {
		goto L100
	}
L94:
	;
	goto L95
L95:
	;
	if v124 != 0 {
		goto L115
	} else {
		goto L116
	}
L96:
	;
	if v303 == int32(0) {
		goto L92
	} else {
		goto L114
	}
L97:
	;
	v303 = int32(0)
	goto L96
L98:
	;
	v277 = v272
	v278 = v273
	v279 = v274
	goto L108
L99:
	;
	if (v239|l2)&int32(3) != 0 {
		v272 = v239
		v273 = l2
		v274 = v124
		goto L98
	} else {
		goto L102
	}
L100:
	;
	v265 = v239
	v266 = l2
	v267 = v124
	goto L101
L101:
	;
	if v267 == int32(0) {
		goto L97
	} else {
		goto L107
	}
L102:
	;
	v249 = v239
	v250 = l2
	v251 = v124
	goto L103
L103:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	if v254 != v255 {
		v272 = v249
		v273 = v250
		v274 = v251
		goto L98
	} else {
		goto L105
	}
L104:
	;
	v265 = v260
	v266 = v258
	v267 = v262
	goto L101
L105:
	;
	v257 = int32(4)
	v258 = v250 + v257
	v260 = v249 + v257
	v262 = v251 - v257
	if base.Ui32(int32(3)) < base.Ui32(v262) {
		v249 = v260
		v250 = v258
		v251 = v262
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	v272 = v265
	v273 = v266
	v274 = v267
	goto L98
L108:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278))))
	if v282 == v283 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v303 = v282 - v283
	goto L96
L110:
	;
	v285 = int32(1)
	v290 = v279 - v285
	if v290 != 0 {
		v277 = v277 + v285
		v278 = v278 + v285
		v279 = v290
		goto L108
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	goto L109
L113:
	;
	goto L97
L114:
	;
	goto L95
L115:
	;
	base.MemoryCopy(m, v239, l2, v124)
	goto L117
L116:
	;
	goto L117
L117:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v309 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v307+v124))) = uint8(v309)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v124
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v318 = v312
	goto L91
L118:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
	if v315 != 0 {
		v318 = v239
		goto L91
	} else {
		goto L119
	}
L119:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	return v316
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v387
	v389 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)) = uint8(v389)
	return v387
L121:
	;
	if v353 != 0 {
		v387 = v353
		goto L120
	} else {
		goto L133
	}
L122:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	if base.B2i32(v326 == int32(0))|base.B2i32(v326 != v329) != 0 {
		v347 = v326
		v348 = v329
		goto L126
	} else {
		goto L127
	}
L123:
	;
	goto L124
L124:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	v351 = m.T0[v350].(func(*base.Module, int32, int32, int32) int32)(m, v319, v318, v320)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L54
	} else {
		goto L132
	}
L125:
	;
	v353 = v347 - v348
	goto L121
L126:
	;
	goto L125
L127:
	;
	v332 = v319
	v333 = v318
	goto L128
L128:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+1)))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+1)))
	if v337 == int32(0) {
		v347 = v337
		v348 = v336
		goto L126
	} else {
		goto L130
	}
L129:
	;
	v347 = v337
	v348 = v336
	goto L126
L130:
	;
	v340 = int32(1)
	if v337 == v336 {
		v332 = v332 + v340
		v333 = v333 + v340
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v353 = v351
	goto L121
L133:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355))))
	if v356 != int32(1) {
		v387 = int32(0)
		goto L120
	} else {
		goto L134
	}
L134:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359))))
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360))))
	if base.B2i32(v363 == int32(0))|base.B2i32(v363 != v366) != 0 {
		v384 = v363
		v385 = v366
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v387 = v384 - v385
	goto L120
L136:
	;
	goto L135
L137:
	;
	v369 = v359
	v370 = v360
	goto L138
L138:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370)+1)))
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369)+1)))
	if v374 == int32(0) {
		v384 = v374
		v385 = v373
		goto L136
	} else {
		goto L140
	}
L139:
	;
	v384 = v374
	v385 = v373
	goto L136
L140:
	;
	v377 = int32(1)
	if v374 == v373 {
		v369 = v369 + v377
		v370 = v370 + v377
		goto L138
	} else {
		goto L141
	}
L141:
	;
	goto L139
}
func F_void_out(m *base.Module, l0 int32) int64 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_pstrdup(m, int32(_a_F_void_out_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v3)
	}
}
