package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_varbit_out(m *base.Module, l0 int32) int32 {
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v15 = v10 + int32(8)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v19 = F_palloc(m, v16+int32(1))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
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
			return v19
		}
	}
}
func F_varbit_recv(m *base.Module, l0 int32) int32 {
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
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pq_getmsgint(m, v12, int32(4))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if base.Ui32(v14) < base.Ui32(int32(2147483641)) {
			if base.B2i32(v11 < v14)&base.B2i32(int32(0) < v11) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16777346))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
						F_errmsg(m, int32(_a_F_varbit_recv_0), v9)
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_varbit_recv_1), int32(662), int32(_a_F_varbit_recv_2))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
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
				v27 = int32(base.Ui32(v14+int32(7)) >> (uint(int32(3)) % 32))
				v29 = v27 + int32(8)
				v30 = F_palloc(m, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v14
					*(*int32)(unsafe.Add(mBase, uint32(v30))) = v29 << (uint(int32(2)) % 32)
					F_pq_copymsgbytes(m, v12, v30+int32(8), v27)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
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
						return v30
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50462850))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_varbit_recv_3), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_varbit_recv_1), int32(652), int32(_a_F_varbit_recv_2))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
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
}
func F_varchar_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v39 int32
	_ = v39
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v6 != int32(457) {
		v39 = v2
		return v39
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v13 != int32(7) {
			v39 = v2
			return v39
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)))
			if v16 != 0 {
				v39 = v2
				return v39
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v18 = F_exprTypmod(m, v17)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
					v23 = int32(0)
					v27 = int32(4)
					if base.B2i32(v23 <= v22)&(base.B2i32(v18 < v23)|base.B2i32(v22-v27 < v18-v27)) != 0 {
						v39 = v2
						return v39
					} else {
						v34 = F_relabel_to_typmod(m, v17, v22)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v39 = v34
							return v39
						}
					}
				}
			}
		}
	}
}
func F_varstr_abbrev_convert(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
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
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
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
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
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
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v15 = F_pg_detoast_datum_packed(m, l0)
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
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
	v24 = int32(1)
	v25 = v15 + v24
	v27 = v19 & v24
	if v19 == v24 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v27 != 0 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v33 == int32(18) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v44 = int32(1)
	if v27 != 0 {
		v54 = int32(base.Ui32(v19)>>(uint(v44)%32)) - v44
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v36 = int32(16)
	goto L9
L8:
	;
	v36 = int32(0)
	goto L9
L9:
	;
	if base.Ui32((v33-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v43 = int32(4)
	goto L12
L11:
	;
	v43 = v36
	goto L12
L12:
	;
	v54 = v43
	goto L3
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
	goto L3
L14:
	;
	v55 = v25
	goto L16
L15:
	;
	v55 = v15 + int32(4)
	goto L16
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	if v56 == int32(1042) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v61 = int32(-1)
	v63 = v54 - int32(1)
	if v61 <= v63 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v82 = v54
	goto L19
L19:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+29)))
	if v83 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L20:
	;
	v82 = v81
	goto L19
L21:
	;
	v66 = v61
	goto L23
L22:
	;
	v66 = v63
	goto L23
L23:
	;
	v70 = v54
	goto L24
L24:
	;
	v74 = v70 - int32(1)
	if v74 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v81 = v70
	goto L20
L26:
	;
	v81 = v66 + int32(1)
	goto L20
L27:
	;
	goto L28
L28:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v74))))
	if v78 == int32(32) {
		v70 = v74
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	if l0 != v15 {
		goto L178
	} else {
		goto L179
	}
L31:
	;
	v684 = int32(4)
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if base.Ui32(v684) <= base.Ui32(v685) {
		goto L172
	} else {
		goto L173
	}
L32:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v86 <= v82 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v266 = v55
	v268 = v82
	goto L34
L34:
	;
	v274 = int32(4)
	if base.Ui32(v274) <= base.Ui32(v268) {
		goto L98
	} else {
		goto L99
	}
L35:
	;
	v88 = int32(1)
	v89 = v82 + v88
	v90 = int32(1073741823)
	v92 = v86 << (uint(v88) % 32)
	if base.Ui32(v90) <= base.Ui32(v92) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v82 != v105 {
		goto L46
	} else {
		goto L47
	}
L38:
	;
	v95 = v90
	goto L40
L39:
	;
	v95 = v92
	goto L40
L40:
	;
	if base.Ui32(v95) < base.Ui32(v89) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v97 = v89
	goto L43
L42:
	;
	v97 = v95
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v100 = F_repalloc(m, v99, v97)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v100
	goto L37
L45:
	;
	if v82 != 0 {
		goto L69
	} else {
		goto L70
	}
L46:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v176 = v107
	goto L45
L47:
	;
	goto L48
L48:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)))
	if v109 != int32(1) {
		v176 = v108
		goto L45
	} else {
		goto L49
	}
L49:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v82) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	if v173 == int32(0) {
		goto L31
	} else {
		goto L68
	}
L51:
	;
	v173 = int32(0)
	goto L50
L52:
	;
	v147 = v142
	v148 = v143
	v149 = v144
	goto L62
L53:
	;
	if (v108|v55)&int32(3) != 0 {
		v142 = v108
		v143 = v55
		v144 = v82
		goto L52
	} else {
		goto L56
	}
L54:
	;
	v135 = v108
	v136 = v55
	v137 = v82
	goto L55
L55:
	;
	if v137 == int32(0) {
		goto L51
	} else {
		goto L61
	}
L56:
	;
	v119 = v108
	v120 = v55
	v121 = v82
	goto L57
L57:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v124 != v125 {
		v142 = v119
		v143 = v120
		v144 = v121
		goto L52
	} else {
		goto L59
	}
L58:
	;
	v135 = v130
	v136 = v128
	v137 = v132
	goto L55
L59:
	;
	v127 = int32(4)
	v128 = v120 + v127
	v130 = v119 + v127
	v132 = v121 - v127
	if base.Ui32(int32(3)) < base.Ui32(v132) {
		v119 = v130
		v120 = v128
		v121 = v132
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v142 = v135
	v143 = v136
	v144 = v137
	goto L52
L62:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v152 == v153 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v173 = v152 - v153
	goto L50
L64:
	;
	v155 = int32(1)
	v160 = v149 - v155
	if v160 != 0 {
		v147 = v147 + v155
		v148 = v148 + v155
		v149 = v160
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
	v176 = v108
	goto L45
L69:
	;
	base.MemoryCopy(m, v176, v55, v82)
	goto L71
L70:
	;
	goto L71
L71:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v180 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v178+v82))) = uint8(v180)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v82
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+8))
	if v185 == v180 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v266 = v264
	v268 = v258
	goto L34
L73:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	v192 = F_pg_strxfrm(m, v188, v189, v190, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if base.Ui32(int32(4)) <= base.Ui32(v229) {
		goto L90
	} else {
		goto L91
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v192
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if base.Ui32(v192) < base.Ui32(v195) {
		v258 = v192
		goto L72
	} else {
		goto L77
	}
L77:
	;
	v198 = v195
	v200 = v192
	goto L78
L78:
	;
	v206 = int32(1)
	v207 = v200 + v206
	v208 = int32(1073741823)
	v210 = v198 << (uint(v206) % 32)
	if base.Ui32(v208) <= base.Ui32(v210) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v258 = v224
	goto L72
L80:
	;
	v213 = v208
	goto L82
L81:
	;
	v213 = v210
	goto L82
L82:
	;
	if base.Ui32(v213) < base.Ui32(v207) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v215 = v207
	goto L85
L84:
	;
	v215 = v213
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v215
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v218 = F_repalloc(m, v217, v215)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v218
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	v224 = F_pg_strxfrm(m, v218, v221, v222, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v224
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if base.Ui32(v227) <= base.Ui32(v224) {
		v198 = v227
		v200 = v224
		goto L78
	} else {
		goto L88
	}
L88:
	;
	goto L79
L89:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+8))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+8))
	v252 = m.T0[v251].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v245, int32(4), v247, int32(-1), v249)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L97
	}
L90:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v245 = v232
	goto L89
L91:
	;
	goto L92
L92:
	;
	v233 = int32(2)
	if base.Ui32(v229) <= base.Ui32(v233) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v236 = v233
	goto L95
L94:
	;
	v236 = v229
	goto L95
L95:
	;
	v238 = v236 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v241 = F_repalloc(m, v240, v238)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v241
	v245 = v241
	goto L89
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v252
	v258 = v252
	goto L72
L98:
	;
	v277 = v274
	goto L100
L99:
	;
	v277 = v268
	goto L100
L100:
	;
	if v277 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	base.MemoryCopy(m, v12+int32(12), v266, v277)
	goto L103
L102:
	;
	goto L103
L103:
	;
	v281 = int32(128)
	if v281 <= v82 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v284 = v281
	goto L106
L105:
	;
	v284 = v82
	goto L106
L106:
	;
	v290 = v284 - int32(1636608432)
	if v55&int32(3) != 0 {
		goto L111
	} else {
		goto L112
	}
L107:
	;
	v550 = v14 - int32(-64)
	if int32(129) <= v82 {
		goto L147
	} else {
		goto L148
	}
L108:
	;
	v522 = int32(14)
	v524 = v518 ^ v519 - base.I32_rotl(v518, v522)
	v528 = v524 ^ v517 - base.I32_rotl(v524, int32(11))
	v532 = v528 ^ v518 - base.I32_rotl(v528, int32(25))
	v536 = v532 ^ v524 - base.I32_rotl(v532, int32(16))
	v540 = v536 ^ v528 - base.I32_rotl(v536, int32(4))
	v544 = v540 ^ v532 - base.I32_rotl(v540, v522)
	v548 = v544 ^ v536 - base.I32_rotl(v544, int32(24))
	goto L107
L109:
	;
	switch v448 - int32(1) {
	case 0:
		v510 = v449
		v511 = v450
		v512 = v451
		goto L136
	case 1:
		v503 = v449
		v504 = v450
		v505 = v451
		goto L137
	case 2:
		v496 = v449
		v497 = v450
		v498 = v451
		goto L138
	case 3:
		v490 = v450
		v491 = v451
		goto L139
	case 4:
		v486 = v450
		v487 = v451
		goto L140
	case 5:
		v480 = v450
		v481 = v451
		goto L141
	case 6:
		v474 = v450
		v475 = v451
		goto L142
	case 7:
		v469 = v451
		goto L143
	case 8:
		v464 = v451
		goto L144
	case 9:
		v459 = v451
		goto L145
	case 10:
		goto L146
	default:
		v517 = v449
		v518 = v450
		v519 = v451
		goto L108
	}
L110:
	;
	v399 = v55
	v400 = v284
	v401 = v290
	v402 = v290
	v403 = v290
	goto L133
L111:
	;
	if base.Ui32(int32(11)) < base.Ui32(v284) {
		goto L110
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	if base.Ui32(v284) < base.Ui32(int32(12)) {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v447 = v55
	v448 = v284
	v449 = v290
	v450 = v290
	v451 = v290
	goto L109
L115:
	;
	switch v346 - int32(1) {
	case 0:
		v396 = v347
		goto L122
	case 1:
		v391 = v347
		goto L123
	case 2:
		goto L124
	case 3:
		v384 = v348
		goto L125
	case 4:
		v381 = v348
		goto L126
	case 5:
		v376 = v348
		goto L127
	case 6:
		goto L128
	case 7:
		v367 = v349
		goto L129
	case 8:
		v362 = v349
		goto L130
	case 9:
		v357 = v349
		goto L131
	case 10:
		goto L132
	default:
		v517 = v347
		v518 = v348
		v519 = v349
		goto L108
	}
L116:
	;
	v345 = v55
	v346 = v284
	v347 = v290
	v348 = v290
	v349 = v290
	goto L115
L117:
	;
	goto L118
L118:
	;
	v297 = v55
	v298 = v284
	v299 = v290
	v300 = v290
	v301 = v290
	goto L119
L119:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	v304 = v303 + v300
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v297)+8))
	v308 = v307 + v301
	v310 = int32(4)
	v312 = v305 + v299 - v308 ^ base.I32_rotl(v308, v310)
	v316 = v304 - v312 ^ base.I32_rotl(v312, int32(6))
	v317 = v308 + v304
	v318 = v312 + v317
	v319 = v316 + v318
	v323 = v317 - v316 ^ base.I32_rotl(v316, int32(8))
	v327 = v318 - v323 ^ base.I32_rotl(v323, int32(16))
	v331 = v319 - v327 ^ base.I32_rotl(v327, int32(19))
	v332 = v323 + v319
	v333 = v327 + v332
	v334 = v331 + v333
	v338 = v332 - v331 ^ base.I32_rotl(v331, v310)
	v339 = int32(12)
	v340 = v297 + v339
	v342 = v298 - v339
	if base.Ui32(int32(11)) < base.Ui32(v342) {
		v297 = v340
		v298 = v342
		v299 = v333
		v300 = v334
		v301 = v338
		goto L119
	} else {
		goto L121
	}
L120:
	;
	v345 = v340
	v346 = v342
	v347 = v333
	v348 = v334
	v349 = v338
	goto L115
L121:
	;
	goto L120
L122:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345))))
	v517 = v396 + v397
	v518 = v348
	v519 = v349
	goto L108
L123:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+1)))
	v396 = v392<<(uint(int32(8))%32) + v391
	goto L122
L124:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+2)))
	v391 = v387<<(uint(int32(16))%32) + v347
	goto L123
L125:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v517 = v385 + v347
	v518 = v384
	v519 = v349
	goto L108
L126:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+4)))
	v384 = v381 + v382
	goto L125
L127:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+5)))
	v381 = v377<<(uint(int32(8))%32) + v376
	goto L126
L128:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+6)))
	v376 = v372<<(uint(int32(16))%32) + v348
	goto L127
L129:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	v517 = v368 + v347
	v518 = v370 + v348
	v519 = v367
	goto L108
L130:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+8)))
	v367 = v363<<(uint(int32(8))%32) + v362
	goto L129
L131:
	;
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+9)))
	v362 = v358<<(uint(int32(16))%32) + v357
	goto L130
L132:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+10)))
	v357 = v353<<(uint(int32(24))%32) + v349
	goto L131
L133:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v399)+4))
	v406 = v405 + v402
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v399)+8))
	v410 = v409 + v403
	v412 = int32(4)
	v414 = v407 + v401 - v410 ^ base.I32_rotl(v410, v412)
	v418 = v406 - v414 ^ base.I32_rotl(v414, int32(6))
	v419 = v410 + v406
	v420 = v414 + v419
	v421 = v418 + v420
	v425 = v419 - v418 ^ base.I32_rotl(v418, int32(8))
	v429 = v420 - v425 ^ base.I32_rotl(v425, int32(16))
	v433 = v421 - v429 ^ base.I32_rotl(v429, int32(19))
	v434 = v425 + v421
	v435 = v429 + v434
	v436 = v433 + v435
	v440 = v434 - v433 ^ base.I32_rotl(v433, v412)
	v441 = int32(12)
	v442 = v399 + v441
	v444 = v400 - v441
	if base.Ui32(int32(11)) < base.Ui32(v444) {
		v399 = v442
		v400 = v444
		v401 = v435
		v402 = v436
		v403 = v440
		goto L133
	} else {
		goto L135
	}
L134:
	;
	v447 = v442
	v448 = v444
	v449 = v435
	v450 = v436
	v451 = v440
	goto L109
L135:
	;
	goto L134
L136:
	;
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447))))
	v517 = v510 + v513
	v518 = v511
	v519 = v512
	goto L108
L137:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+1)))
	v510 = v506<<(uint(int32(8))%32) + v503
	v511 = v504
	v512 = v505
	goto L136
L138:
	;
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+2)))
	v503 = v499<<(uint(int32(16))%32) + v496
	v504 = v497
	v505 = v498
	goto L137
L139:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+3)))
	v496 = v492<<(uint(int32(24))%32) + v449
	v497 = v490
	v498 = v491
	goto L138
L140:
	;
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+4)))
	v490 = v486 + v488
	v491 = v487
	goto L139
L141:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+5)))
	v486 = v482<<(uint(int32(8))%32) + v480
	v487 = v481
	goto L140
L142:
	;
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+6)))
	v480 = v476<<(uint(int32(16))%32) + v474
	v481 = v475
	goto L141
L143:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+7)))
	v474 = v470<<(uint(int32(24))%32) + v450
	v475 = v469
	goto L142
L144:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+8)))
	v469 = v465<<(uint(int32(8))%32) + v464
	goto L143
L145:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+9)))
	v464 = v460<<(uint(int32(16))%32) + v459
	goto L144
L146:
	;
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+10)))
	v459 = v455<<(uint(int32(24))%32) + v451
	goto L145
L147:
	;
	v557 = int32(711645284)
	v560 = v82 - int32(1636608428) ^ v557 - int32(1455628627)
	v565 = v560 ^ int32(-1636608428) - base.I32_rotl(v560, int32(25))
	v570 = v565 ^ v557 - base.I32_rotl(v565, int32(16))
	v574 = v570 ^ v560 - base.I32_rotl(v570, int32(4))
	v578 = v574 ^ v565 - base.I32_rotl(v574, int32(14))
	goto L150
L148:
	;
	v584 = v548
	goto L149
L149:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v550)+16))
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550))))
	v590 = int32(32) - v589
	v592 = v587 + int32(base.Ui32(v584)>>(uint(v590)%32))
	v593 = v584 << (uint(v589) % 32)
	if v593 != 0 {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	v584 = v578 ^ v570 - base.I32_rotl(v578, int32(24)) ^ v548
	goto L149
L151:
	;
	v618 = v14 + int32(40)
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v624 = int32(711645284)
	v627 = v619 - int32(1636608428) ^ v624 - int32(1455628627)
	v632 = v627 ^ int32(-1636608428) - base.I32_rotl(v627, int32(25))
	v637 = v632 ^ v624 - base.I32_rotl(v632, int32(16))
	v641 = v637 ^ v627 - base.I32_rotl(v637, int32(4))
	v645 = v641 ^ v632 - base.I32_rotl(v641, int32(14))
	v649 = v645 ^ v637 - base.I32_rotl(v645, int32(24))
	goto L161
L152:
	;
	v600 = int32(32) - (base.I32_clz(v593) ^ int32(31))
	v601 = int32(255)
	if base.Ui32(v590&v601) < base.Ui32(v600&v601) {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	v610 = v590 + int32(1)
	goto L154
L154:
	;
	v612 = v610 & int32(255)
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v592))))
	if base.Ui32(v613) < base.Ui32(v612) {
		goto L158
	} else {
		goto L159
	}
L155:
	;
	v606 = v590 + int32(1)
	goto L157
L156:
	;
	v606 = v600
	goto L157
L157:
	;
	v610 = v606
	goto L154
L158:
	;
	v615 = v612
	goto L160
L159:
	;
	v615 = v613
	goto L160
L160:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v592))) = uint8(v615)
	goto L151
L161:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v618)+16))
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618))))
	v655 = int32(32) - v654
	v657 = v652 + int32(base.Ui32(v649)>>(uint(v655)%32))
	v658 = v649 << (uint(v654) % 32)
	if v658 != 0 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v682 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)) = uint8(v682)
	v697 = v619
	goto L30
L163:
	;
	v665 = int32(32) - (base.I32_clz(v658) ^ int32(31))
	v666 = int32(255)
	if base.Ui32(v655&v666) < base.Ui32(v665&v666) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	v675 = v655 + int32(1)
	goto L165
L165:
	;
	v677 = v675 & int32(255)
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v657))))
	if base.Ui32(v678) < base.Ui32(v677) {
		goto L169
	} else {
		goto L170
	}
L166:
	;
	v671 = v655 + int32(1)
	goto L168
L167:
	;
	v671 = v665
	goto L168
L168:
	;
	v675 = v671
	goto L165
L169:
	;
	v680 = v677
	goto L171
L170:
	;
	v680 = v678
	goto L171
L171:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v657))) = uint8(v680)
	goto L162
L172:
	;
	v688 = v684
	goto L174
L173:
	;
	v688 = v685
	goto L174
L174:
	;
	if v688 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	base.MemoryCopy(m, v12+int32(12), v691, v688)
	goto L177
L176:
	;
	goto L177
L177:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v697 = v693
	goto L30
L178:
	;
	F_pfree(m, v15)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	m.G0 = v12 + int32(16)
	v711 = int32(16711935)
	return base.I32_rotr(v697, int32(24))&v711 | base.I32_rotr(v697&v711, int32(8))
L181:
	;
	goto L180
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
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+2)))
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
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
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
	F_errfinish(m, int32(_a_F_varstr_cmp_2), int32(1648), int32(_a_F_varstr_cmp_3))
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
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
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
	var v362 int32
	_ = v362
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
	v320 = int32(-1)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)+8))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	v325 = m.T0[v324].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v319, v320, v318, v320, v322)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L54
	} else {
		goto L121
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v360
	v362 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)) = uint8(v362)
	return v360
L121:
	;
	if v325 != 0 {
		v360 = v325
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+1)))
	if v329 != int32(1) {
		v360 = int32(0)
		goto L120
	} else {
		goto L123
	}
L123:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
	if base.B2i32(v336 == int32(0))|base.B2i32(v336 != v339) != 0 {
		v357 = v336
		v358 = v339
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v360 = v357 - v358
	goto L120
L125:
	;
	goto L124
L126:
	;
	v342 = v332
	v343 = v333
	goto L127
L127:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+1)))
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+1)))
	if v347 == int32(0) {
		v357 = v347
		v358 = v346
		goto L125
	} else {
		goto L129
	}
L128:
	;
	v357 = v347
	v358 = v346
	goto L125
L129:
	;
	v350 = int32(1)
	if v347 == v346 {
		v342 = v342 + v350
		v343 = v343 + v350
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
}
func F_void_out(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_pstrdup(m, int32(_a_F_void_out_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
