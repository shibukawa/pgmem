package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgp_armor_headers(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
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
	var v29 int32
	_ = v29
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
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
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v16 == int32(0) {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v20 = F_pg_detoast_datum_packed(m, v19)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = int32(_a_F_pgp_armor_headers_0)
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_pgp_armor_headers[0]))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
				*(*int32)(unsafe.Add(mBase, _c_F_pgp_armor_headers[0])) = v29
				v34 = F_get_call_result_type(m, l0, int32(0), v13+int32(12))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					if v34 != int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v157 = m.ExcPending
						if v157 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_pgp_armor_headers_1), int32(0))
							mBase = m.M
							v161 = m.ExcPending
							if v161 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_pgp_armor_headers_2), int32(937), int32(_a_F_pgp_armor_headers_3))
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
						v39 = F_TupleDescGetAttInMetadata(m, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v39
							v43 = F_palloc(m, int32(12))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								v45 = int32(1)
								v46 = v20 + v45
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
								v51 = v49 & v45
								if v51 != 0 {
									v52 = v46
								} else {
									v52 = v20 + int32(4)
								}
								if v49 == int32(1) {
									v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
									if v58 == int32(18) {
										v61 = int32(16)
									} else {
										v61 = int32(0)
									}
									if base.Ui32((v58-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v68 = int32(4)
									} else {
										v68 = v61
									}
									v79 = v68
								} else {
									v69 = int32(1)
									if v51 != 0 {
										v79 = int32(base.Ui32(v49)>>(uint(v69)%32)) - v69
									} else {
										v73 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
										v79 = int32(base.Ui32(v73)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								v84 = F_pgp_extract_armor_headers(m, v52, v79, v43, v43+int32(4), v43+int32(8))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									if v84 < int32(0) {
										F_px_THROW_ERROR(m, v84)
										mBase = m.M
										v168 = m.ExcPending
										if v168 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_pgp_armor_headers[0])) = v27
										*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v43
										v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+16))
										v100 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v99)+16))
										v102 = int64(*(*int32)(unsafe.Add(mBase, uint32(v101))))
										if base.Ui64(v102) <= base.Ui64(v100) {
											F_end_MultiFuncCall(m, l0)
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return int32(0)
											} else {
												v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v106)+20)) = int32(2)
												v109 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v109)
												v149 = int32(0)
												m.G0 = v13 + int32(16)
												return v149
											}
										} else {
											v114 = base.I32_wrap_i64(v100) << (uint(int32(2)) % 32)
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
											v117 = *(*int32)(unsafe.Add(mBase, uint32(v114+v115)))
											v118 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
											v120 = *(*int32)(unsafe.Add(mBase, uint32(v118+v114)))
											v121 = F_strlen(m, v120)
											mBase = m.M
											v123 = F_pg_any_to_server(m, v120, v121, int32(6))
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v123
												v126 = F_strlen(m, v117)
												mBase = m.M
												v128 = F_pg_any_to_server(m, v117, v126, int32(6))
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v128
													v131 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
													v134 = F_BuildTupleFromCStrings(m, v131, v13+int32(4))
													mBase = m.M
													v135 = m.ExcPending
													if v135 != 0 {
														return int32(0)
													} else {
														v136 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
														*(*int64)(unsafe.Add(mBase, uint32(v99))) = v136 + int64(1)
														v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v140)+20)) = int32(1)
														v143 = *(*int32)(unsafe.Add(mBase, uint32(v134)+16))
														v144 = F_HeapTupleHeaderGetDatum(m, v143)
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
															return int32(0)
														} else {
															v149 = v144
															m.G0 = v13 + int32(16)
															return v149
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
			}
		}
	} else {
		v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+16))
		v100 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
		v101 = *(*int32)(unsafe.Add(mBase, uint32(v99)+16))
		v102 = int64(*(*int32)(unsafe.Add(mBase, uint32(v101))))
		if base.Ui64(v102) <= base.Ui64(v100) {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v105 = m.ExcPending
			if v105 != 0 {
				return int32(0)
			} else {
				v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v106)+20)) = int32(2)
				v109 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v109)
				v149 = int32(0)
				m.G0 = v13 + int32(16)
				return v149
			}
		} else {
			v114 = base.I32_wrap_i64(v100) << (uint(int32(2)) % 32)
			v115 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
			v117 = *(*int32)(unsafe.Add(mBase, uint32(v114+v115)))
			v118 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
			v120 = *(*int32)(unsafe.Add(mBase, uint32(v118+v114)))
			v121 = F_strlen(m, v120)
			mBase = m.M
			v123 = F_pg_any_to_server(m, v120, v121, int32(6))
			mBase = m.M
			v124 = m.ExcPending
			if v124 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v123
				v126 = F_strlen(m, v117)
				mBase = m.M
				v128 = F_pg_any_to_server(m, v117, v126, int32(6))
				mBase = m.M
				v129 = m.ExcPending
				if v129 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v128
					v131 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
					v134 = F_BuildTupleFromCStrings(m, v131, v13+int32(4))
					mBase = m.M
					v135 = m.ExcPending
					if v135 != 0 {
						return int32(0)
					} else {
						v136 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
						*(*int64)(unsafe.Add(mBase, uint32(v99))) = v136 + int64(1)
						v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v140)+20)) = int32(1)
						v143 = *(*int32)(unsafe.Add(mBase, uint32(v134)+16))
						v144 = F_HeapTupleHeaderGetDatum(m, v143)
						mBase = m.M
						v145 = m.ExcPending
						if v145 != 0 {
							return int32(0)
						} else {
							v149 = v144
							m.G0 = v13 + int32(16)
							return v149
						}
					}
				}
			}
		}
	}
}
func F_pgp_expect_packet_end(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v10 = F_pullf_read(m, l0, int32(_a_F_pgp_expect_packet_end_0), v5+int32(12))
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if int32(0) < v10 {
			F_px_debug(m, int32(_a_F_pgp_expect_packet_end_1), int32(0))
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v21 = int32(-100)
				m.G0 = v5 + int32(16)
				return v21
			}
		} else {
			v21 = v10
			m.G0 = v5 + int32(16)
			return v21
		}
	}
}
func F_pgp_extract_armor_headers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
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
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v493 int32
	_ = v493
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(-101)
	v17 = l0 + l1
	if base.Ui32(v17) <= base.Ui32(l0) {
		v129 = v16
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L116
	} else {
		goto L149
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L116
	} else {
		goto L146
	}
L3:
	;
	m.G0 = v14 + int32(16)
	return v493
L4:
	;
	if v129 <= int32(0) {
		v493 = v16
		goto L3
	} else {
		goto L43
	}
L5:
	;
	goto L4
L6:
	;
	v29 = int32(10)
	goto L8
L8:
	;
	goto L9
L9:
	;
	if v17-l0 < v29 {
		v129 = v16
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v34 = int32(_a_F_pgp_extract_armor_headers_0)
	goto L12
L12:
	;
	goto L13
L13:
	;
	v36 = int32(*(*int8)(unsafe.Add(mBase, _c_F_pgp_extract_armor_headers[0])))
	v40 = l0
	goto L14
L14:
	;
	v47 = F_memchr(m, v40, v36, v17-v40)
	mBase = m.M
	if v47 == int32(0) {
		v129 = v16
		goto L5
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(12)))) = v47
	if base.Ui32(v17) <= base.Ui32(v50) {
		v87 = v50
		goto L27
	} else {
		goto L28
	}
L16:
	;
	v50 = v47 + v29
	if base.Ui32(v17) < base.Ui32(v50) {
		v129 = v16
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v52 = F_memcmp(m, v47, v34, v29)
	mBase = m.M
	if v52 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v54 = v47 + int32(1)
	if base.Ui32(v54) < base.Ui32(v17) {
		v40 = v54
		goto L14
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if l0 == v47 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v129 = v16
	goto L5
L22:
	;
	goto L15
L23:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47-int32(1)))))
	if v59 == int32(10) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if base.Ui32(v17) <= base.Ui32(v50) {
		v129 = v16
		goto L5
	} else {
		goto L25
	}
L25:
	;
	if v29 <= v17-v50 {
		v40 = v50
		goto L14
	} else {
		goto L26
	}
L26:
	;
	v129 = v16
	goto L5
L27:
	;
	if v17-v87 < int32(5) {
		v129 = v16
		goto L5
	} else {
		goto L34
	}
L28:
	;
	v70 = v50
	goto L29
L29:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v76 == int32(45) {
		v87 = v70
		goto L27
	} else {
		goto L31
	}
L30:
	;
	v87 = v17
	goto L27
L31:
	;
	if base.Ui32(v76) < base.Ui32(int32(32)) {
		v129 = v16
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v82 = v70 + int32(1)
	if base.Ui32(v82) < base.Ui32(v17) {
		v70 = v82
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_pgp_extract_armor_headers[0]))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+4)))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgp_extract_armor_headers[1])))
	if v96^v97|(v99^v100) != 0 {
		v129 = v16
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v104 = v87 + int32(5)
	if base.Ui32(v17) <= base.Ui32(v104) {
		v119 = v104
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v129 = v119 - v47
	goto L5
L37:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	switch v106 - int32(10) {
	case 0, 3:
		goto L38
	default:
		v129 = v16
		goto L5
	}
L38:
	;
	if v106 == int32(13) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v113 = v87 + int32(6)
	goto L41
L40:
	;
	v113 = v104
	goto L41
L41:
	;
	if base.Ui32(v17) <= base.Ui32(v113) {
		v119 = v113
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	v119 = v113 + base.B2i32(v115 == int32(10))
	goto L36
L43:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v134 = v133 + v129
	v143 = int32(-101)
	if base.Ui32(v17) <= base.Ui32(v134) {
		v246 = v143
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v246 <= int32(0) {
		v493 = v16
		goto L3
	} else {
		goto L83
	}
L45:
	;
	goto L44
L46:
	;
	v145 = int32(8)
	goto L47
L47:
	;
	goto L49
L49:
	;
	if v17-v134 < v145 {
		v246 = v143
		goto L45
	} else {
		goto L50
	}
L50:
	;
	v150 = int32(_a_F_pgp_extract_armor_headers_1)
	goto L51
L51:
	;
	goto L53
L53:
	;
	v153 = int32(*(*int8)(unsafe.Add(mBase, _c_F_pgp_extract_armor_headers[2])))
	v157 = v134
	goto L54
L54:
	;
	v164 = F_memchr(m, v157, v153, v17-v157)
	mBase = m.M
	if v164 == int32(0) {
		v246 = v143
		goto L45
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(8)))) = v164
	if base.Ui32(v17) <= base.Ui32(v167) {
		v204 = v167
		goto L67
	} else {
		goto L68
	}
L56:
	;
	v167 = v164 + v145
	if base.Ui32(v17) < base.Ui32(v167) {
		v246 = v143
		goto L45
	} else {
		goto L57
	}
L57:
	;
	v169 = F_memcmp(m, v164, v150, v145)
	mBase = m.M
	if v169 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v171 = v164 + int32(1)
	if base.Ui32(v171) < base.Ui32(v17) {
		v157 = v171
		goto L54
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v134 == v164 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v246 = v143
	goto L45
L62:
	;
	goto L55
L63:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164-int32(1)))))
	if v176 == int32(10) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	if base.Ui32(v17) <= base.Ui32(v167) {
		v246 = v143
		goto L45
	} else {
		goto L65
	}
L65:
	;
	if v145 <= v17-v167 {
		v157 = v167
		goto L54
	} else {
		goto L66
	}
L66:
	;
	v246 = v143
	goto L45
L67:
	;
	if v17-v204 < int32(5) {
		v246 = v143
		goto L45
	} else {
		goto L74
	}
L68:
	;
	v187 = v167
	goto L69
L69:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	if v193 == int32(45) {
		v204 = v187
		goto L67
	} else {
		goto L71
	}
L70:
	;
	v204 = v17
	goto L67
L71:
	;
	if base.Ui32(v193) < base.Ui32(int32(32)) {
		v246 = v143
		goto L45
	} else {
		goto L72
	}
L72:
	;
	v199 = v187 + int32(1)
	if base.Ui32(v199) < base.Ui32(v17) {
		v187 = v199
		goto L69
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_pgp_extract_armor_headers[2]))
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+4)))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgp_extract_armor_headers[3])))
	if v213^v214|(v216^v217) != 0 {
		v246 = v143
		goto L45
	} else {
		goto L75
	}
L75:
	;
	v221 = v204 + int32(5)
	if base.Ui32(v17) <= base.Ui32(v221) {
		v236 = v221
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v246 = v236 - v164
	goto L45
L77:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	switch v223 - int32(10) {
	case 0, 3:
		goto L78
	default:
		v246 = v143
		goto L45
	}
L78:
	;
	if v223 == int32(13) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v230 = v204 + int32(6)
	goto L81
L80:
	;
	v230 = v221
	goto L81
L81:
	;
	if base.Ui32(v17) <= base.Ui32(v230) {
		v236 = v230
		goto L76
	} else {
		goto L82
	}
L82:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	v236 = v230 + base.B2i32(v232 == int32(10))
	goto L76
L83:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if base.Ui32(v250) <= base.Ui32(v134) {
		v380 = v134
		v387 = v6
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v391 = v380 - v134
	v394 = F_palloc(m, v391+int32(1))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L116
	} else {
		goto L117
	}
L85:
	;
	v252 = v134
	v259 = v6
	goto L86
L86:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	switch v263 - int32(10) {
	case 0, 3:
		v380 = v252
		v387 = v259
		goto L84
	default:
		goto L88
	}
L87:
	;
	v380 = v378
	v387 = v376
	goto L84
L88:
	;
	v267 = v250 - v252
	v268 = int32(0)
	if base.B2i32(v252&int32(3) == v268)|base.B2i32(v267 == v268) != 0 {
		v298 = v252
		v300 = v267
		v301 = base.B2i32(v267 != v268)
		goto L92
	} else {
		goto L93
	}
L89:
	;
	if v372 == int32(0) {
		v493 = v16
		goto L3
	} else {
		goto L114
	}
L90:
	;
	v372 = int32(0)
	goto L89
L91:
	;
	v350 = v343
	v352 = v345
	goto L108
L92:
	;
	if v301 == int32(0) {
		goto L90
	} else {
		goto L99
	}
L93:
	;
	v281 = v252
	v283 = v267
	goto L94
L94:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	if v286 == int32(10) {
		v343 = v281
		v345 = v283
		goto L91
	} else {
		goto L96
	}
L95:
	;
	v298 = v293
	v300 = v289
	v301 = v291
	goto L92
L96:
	;
	v288 = int32(1)
	v289 = v283 - v288
	v290 = int32(0)
	v291 = base.B2i32(v289 != v290)
	v293 = v281 + v288
	if v293&int32(3) == v290 {
		v298 = v293
		v300 = v289
		v301 = v291
		goto L92
	} else {
		goto L97
	}
L97:
	;
	if v289 != 0 {
		v281 = v293
		v283 = v289
		goto L94
	} else {
		goto L98
	}
L98:
	;
	goto L95
L99:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	if base.B2i32(int32(10) == v307)|base.B2i32(base.Ui32(v300) < base.Ui32(int32(4))) == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v316 = v298
	v318 = v300
	goto L103
L101:
	;
	v336 = v298
	v338 = v300
	goto L102
L102:
	;
	if v338 == int32(0) {
		goto L90
	} else {
		goto L107
	}
L103:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	v323 = v322 ^ int32(168430090)
	v326 = int32(-2139062144)
	if (int32(16843008)-v323|v323)&v326 != v326 {
		v343 = v316
		v345 = v318
		goto L91
	} else {
		goto L105
	}
L104:
	;
	v336 = v331
	v338 = v333
	goto L102
L105:
	;
	v330 = int32(4)
	v331 = v316 + v330
	v333 = v318 - v330
	if base.Ui32(int32(3)) < base.Ui32(v333) {
		v316 = v331
		v318 = v333
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	v343 = v336
	v345 = v338
	goto L91
L108:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350))))
	if int32(10) == v355 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L90
L110:
	;
	v372 = v350
	goto L89
L111:
	;
	goto L112
L112:
	;
	v357 = int32(1)
	v360 = v352 - v357
	if v360 != 0 {
		v350 = v350 + v357
		v352 = v360
		goto L108
	} else {
		goto L113
	}
L113:
	;
	goto L109
L114:
	;
	v375 = int32(1)
	v376 = v259 + v375
	v378 = v372 + v375
	if base.Ui32(v378) < base.Ui32(v250) {
		v252 = v378
		v259 = v376
		goto L86
	} else {
		goto L115
	}
L115:
	;
	goto L87
L116:
	;
	return int32(0)
L117:
	;
	if v391 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	base.MemoryCopy(m, v394, v134, v391)
	goto L120
L119:
	;
	goto L120
L120:
	;
	v399 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v391+v394))) = uint8(v399)
	v404 = v387 << (uint(int32(2)) % 32)
	v405 = F_palloc(m, v404)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L116
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v405
	v408 = F_palloc(m, v404)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L116
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v408
	v411 = int32(10)
	v412 = F___strchrnul(m, v394, v411)
	mBase = m.M
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412))))
	if v414 == v411 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	if v418 != 0 {
		goto L127
	} else {
		goto L128
	}
L124:
	;
	v418 = v412
	goto L126
L125:
	;
	v418 = int32(0)
	goto L126
L126:
	;
	goto L123
L127:
	;
	v419 = v418
	v420 = v394
	v424 = v399
	goto L130
L128:
	;
	v475 = v399
	goto L129
L129:
	;
	if v475 != v387 {
		goto L1
	} else {
		goto L145
	}
L130:
	;
	if base.Ui32(v420) < base.Ui32(v419) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v475 = v459
	goto L129
L132:
	;
	v432 = v419 - int32(1)
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432))))
	if v433 == int32(13) {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	v438 = v419
	goto L134
L134:
	;
	v439 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v438))) = uint8(v439)
	v442 = F_strstr(m, v420, int32(_a_F_pgp_extract_armor_headers_2))
	mBase = m.M
	if v442 == v439 {
		v493 = v16
		goto L3
	} else {
		goto L138
	}
L135:
	;
	v436 = v432
	goto L137
L136:
	;
	v436 = v419
	goto L137
L137:
	;
	v438 = v436
	goto L134
L138:
	;
	v445 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v442))) = uint8(v445)
	if v424 == v387 {
		goto L2
	} else {
		goto L139
	}
L139:
	;
	v448 = int32(2)
	v449 = v424 << (uint(v448) % 32)
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v449+v450))) = v420
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v453+v449))) = v442 + v448
	v458 = int32(1)
	v459 = v424 + v458
	v461 = v419 + v458
	v462 = int32(10)
	v463 = F___strchrnul(m, v461, v462)
	mBase = m.M
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463))))
	if v465 == v462 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	if v469 != 0 {
		v419 = v469
		v420 = v461
		v424 = v459
		goto L130
	} else {
		goto L144
	}
L141:
	;
	v469 = v463
	goto L143
L142:
	;
	v469 = int32(0)
	goto L143
L143:
	;
	goto L140
L144:
	;
	goto L131
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v387
	v493 = int32(0)
	goto L3
L146:
	;
	F_errmsg_internal(m, int32(_a_F_pgp_extract_armor_headers_3), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L116
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_pgp_extract_armor_headers_4), int32(473), int32(_a_F_pgp_extract_armor_headers_5))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L116
	} else {
		goto L148
	}
L148:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	F_errmsg_internal(m, int32(_a_F_pgp_extract_armor_headers_3), int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L116
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_pgp_extract_armor_headers_4), int32(484), int32(_a_F_pgp_extract_armor_headers_5))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L116
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgp_get_cipher_code(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v337 int32
	_ = v337
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v382 int32
	_ = v382
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	v7 = int32(_a_F_pgp_get_cipher_code_0)
	v8 = l0
	goto L4
L1:
	;
	return v410
L2:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v408)+4))
	v410 = v409
	goto L1
L3:
	;
	if v45 == int32(0) {
		v408 = int32(_a_F_pgp_get_cipher_code_1)
		goto L2
	} else {
		goto L16
	}
L4:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v11 == v12 {
		v34 = v11
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v45 = int32(0)
	goto L3
L6:
	;
	v36 = int32(1)
	if v34 != 0 {
		v7 = v7 + v36
		v8 = v8 + v36
		goto L4
	} else {
		goto L15
	}
L7:
	;
	if base.Ui32((v11-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v22 = v11 | int32(32)
	goto L10
L9:
	;
	v22 = v11
	goto L10
L10:
	;
	if base.Ui32((v12-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = v12 | int32(32)
	goto L13
L12:
	;
	v31 = v12
	goto L13
L13:
	;
	if v22 == v31 {
		v34 = v22
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v45 = v22 - v31
	goto L3
L15:
	;
	goto L5
L16:
	;
	v52 = int32(_a_F_pgp_get_cipher_code_2)
	v53 = l0
	goto L18
L17:
	;
	if v90 == int32(0) {
		v408 = int32(_a_F_pgp_get_cipher_code_3)
		goto L2
	} else {
		goto L30
	}
L18:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v56 == v57 {
		v79 = v56
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v90 = int32(0)
	goto L17
L20:
	;
	v81 = int32(1)
	if v79 != 0 {
		v52 = v52 + v81
		v53 = v53 + v81
		goto L18
	} else {
		goto L29
	}
L21:
	;
	if base.Ui32((v56-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v67 = v56 | int32(32)
	goto L24
L23:
	;
	v67 = v56
	goto L24
L24:
	;
	if base.Ui32((v57-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v76 = v57 | int32(32)
	goto L27
L26:
	;
	v76 = v57
	goto L27
L27:
	;
	if v67 == v76 {
		v79 = v67
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v90 = v67 - v76
	goto L17
L29:
	;
	goto L19
L30:
	;
	v97 = int32(_a_F_pgp_get_cipher_code_4)
	v98 = l0
	goto L32
L31:
	;
	if v135 == int32(0) {
		v408 = int32(_a_F_pgp_get_cipher_code_5)
		goto L2
	} else {
		goto L44
	}
L32:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v101 == v102 {
		v124 = v101
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v135 = int32(0)
	goto L31
L34:
	;
	v126 = int32(1)
	if v124 != 0 {
		v97 = v97 + v126
		v98 = v98 + v126
		goto L32
	} else {
		goto L43
	}
L35:
	;
	if base.Ui32((v101-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v112 = v101 | int32(32)
	goto L38
L37:
	;
	v112 = v101
	goto L38
L38:
	;
	if base.Ui32((v102-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v121 = v102 | int32(32)
	goto L41
L40:
	;
	v121 = v102
	goto L41
L41:
	;
	if v112 == v121 {
		v124 = v112
		goto L34
	} else {
		goto L42
	}
L42:
	;
	v135 = v112 - v121
	goto L31
L43:
	;
	goto L33
L44:
	;
	v142 = int32(_a_F_pgp_get_cipher_code_6)
	v143 = l0
	goto L46
L45:
	;
	if v180 == int32(0) {
		v408 = int32(_a_F_pgp_get_cipher_code_7)
		goto L2
	} else {
		goto L58
	}
L46:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	if v146 == v147 {
		v169 = v146
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v180 = int32(0)
	goto L45
L48:
	;
	v171 = int32(1)
	if v169 != 0 {
		v142 = v142 + v171
		v143 = v143 + v171
		goto L46
	} else {
		goto L57
	}
L49:
	;
	if base.Ui32((v146-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v157 = v146 | int32(32)
	goto L52
L51:
	;
	v157 = v146
	goto L52
L52:
	;
	if base.Ui32((v147-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v166 = v147 | int32(32)
	goto L55
L54:
	;
	v166 = v147
	goto L55
L55:
	;
	if v157 == v166 {
		v169 = v157
		goto L48
	} else {
		goto L56
	}
L56:
	;
	v180 = v157 - v166
	goto L45
L57:
	;
	goto L47
L58:
	;
	v187 = int32(_a_F_pgp_get_cipher_code_8)
	v188 = l0
	goto L60
L59:
	;
	if v225 == int32(0) {
		v408 = int32(_a_F_pgp_get_cipher_code_9)
		goto L2
	} else {
		goto L72
	}
L60:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v191 == v192 {
		v214 = v191
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v225 = int32(0)
	goto L59
L62:
	;
	v216 = int32(1)
	if v214 != 0 {
		v187 = v187 + v216
		v188 = v188 + v216
		goto L60
	} else {
		goto L71
	}
L63:
	;
	if base.Ui32((v191-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v202 = v191 | int32(32)
	goto L66
L65:
	;
	v202 = v191
	goto L66
L66:
	;
	if base.Ui32((v192-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v211 = v192 | int32(32)
	goto L69
L68:
	;
	v211 = v192
	goto L69
L69:
	;
	if v202 == v211 {
		v214 = v202
		goto L62
	} else {
		goto L70
	}
L70:
	;
	v225 = v202 - v211
	goto L59
L71:
	;
	goto L61
L72:
	;
	v232 = int32(_a_F_pgp_get_cipher_code_10)
	v233 = l0
	goto L74
L73:
	;
	if v270 == int32(0) {
		v408 = int32(_a_F_pgp_get_cipher_code_11)
		goto L2
	} else {
		goto L86
	}
L74:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	if v236 == v237 {
		v259 = v236
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v270 = int32(0)
	goto L73
L76:
	;
	v261 = int32(1)
	if v259 != 0 {
		v232 = v232 + v261
		v233 = v233 + v261
		goto L74
	} else {
		goto L85
	}
L77:
	;
	if base.Ui32((v236-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v247 = v236 | int32(32)
	goto L80
L79:
	;
	v247 = v236
	goto L80
L80:
	;
	if base.Ui32((v237-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v256 = v237 | int32(32)
	goto L83
L82:
	;
	v256 = v237
	goto L83
L83:
	;
	if v247 == v256 {
		v259 = v247
		goto L76
	} else {
		goto L84
	}
L84:
	;
	v270 = v247 - v256
	goto L73
L85:
	;
	goto L75
L86:
	;
	v277 = int32(_a_F_pgp_get_cipher_code_12)
	v278 = l0
	goto L88
L87:
	;
	if v315 == int32(0) {
		v408 = int32(_a_F_pgp_get_cipher_code_13)
		goto L2
	} else {
		goto L100
	}
L88:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278))))
	if v281 == v282 {
		v304 = v281
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v315 = int32(0)
	goto L87
L90:
	;
	v306 = int32(1)
	if v304 != 0 {
		v277 = v277 + v306
		v278 = v278 + v306
		goto L88
	} else {
		goto L99
	}
L91:
	;
	if base.Ui32((v281-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v292 = v281 | int32(32)
	goto L94
L93:
	;
	v292 = v281
	goto L94
L94:
	;
	if base.Ui32((v282-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v301 = v282 | int32(32)
	goto L97
L96:
	;
	v301 = v282
	goto L97
L97:
	;
	if v292 == v301 {
		v304 = v292
		goto L90
	} else {
		goto L98
	}
L98:
	;
	v315 = v292 - v301
	goto L87
L99:
	;
	goto L89
L100:
	;
	v322 = int32(_a_F_pgp_get_cipher_code_14)
	v323 = l0
	goto L102
L101:
	;
	if v360 == int32(0) {
		v408 = int32(_a_F_pgp_get_cipher_code_15)
		goto L2
	} else {
		goto L114
	}
L102:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
	if v326 == v327 {
		v349 = v326
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v360 = int32(0)
	goto L101
L104:
	;
	v351 = int32(1)
	if v349 != 0 {
		v322 = v322 + v351
		v323 = v323 + v351
		goto L102
	} else {
		goto L113
	}
L105:
	;
	if base.Ui32((v326-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v337 = v326 | int32(32)
	goto L108
L107:
	;
	v337 = v326
	goto L108
L108:
	;
	if base.Ui32((v327-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v346 = v327 | int32(32)
	goto L111
L110:
	;
	v346 = v327
	goto L111
L111:
	;
	if v337 == v346 {
		v349 = v337
		goto L104
	} else {
		goto L112
	}
L112:
	;
	v360 = v337 - v346
	goto L101
L113:
	;
	goto L103
L114:
	;
	v367 = int32(_a_F_pgp_get_cipher_code_16)
	v368 = l0
	goto L116
L115:
	;
	if v405 != 0 {
		v410 = int32(-103)
		goto L1
	} else {
		goto L128
	}
L116:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367))))
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368))))
	if v371 == v372 {
		v394 = v371
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v405 = int32(0)
	goto L115
L118:
	;
	v396 = int32(1)
	if v394 != 0 {
		v367 = v367 + v396
		v368 = v368 + v396
		goto L116
	} else {
		goto L127
	}
L119:
	;
	if base.Ui32((v371-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v382 = v371 | int32(32)
	goto L122
L121:
	;
	v382 = v371
	goto L122
L122:
	;
	if base.Ui32((v372-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v391 = v372 | int32(32)
	goto L125
L124:
	;
	v391 = v372
	goto L125
L125:
	;
	if v382 == v391 {
		v394 = v382
		goto L118
	} else {
		goto L126
	}
L126:
	;
	v405 = v382 - v391
	goto L115
L127:
	;
	goto L117
L128:
	;
	v408 = int32(_a_F_pgp_get_cipher_code_17)
	goto L2
}
func F_pgp_get_cipher_key_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v4 = l0 - int32(2)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v4))|base.B2i32(int32(base.Ui32(int32(487))>>(uint(v4)%32))&int32(1) == int32(0)) != 0 {
		v21 = int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v4<<(uint(int32(2))%32))+uint32(_c_F_pgp_get_cipher_key_size[0])))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
		v21 = v20
	}
	return v21
}
func F_pgp_init(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v20 int32
	_ = v20
	v4 = F_palloc0(m, int32(168))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v4)+60)) = int64(7)
		v10 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v4)+72)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(v4)+68)) = int32(6)
		*(*int64)(unsafe.Add(mBase, uint32(v4)+52)) = int64(-4294967294)
		*(*int64)(unsafe.Add(mBase, uint32(v4)+44)) = int64(-4294967293)
		*(*int64)(unsafe.Add(mBase, uint32(v4)+80)) = v10
		v20 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+88)) = v20
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v4
		return v20
	}
}
func F_pgp_rsa_decrypt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v12 = int32(-109)
	v13 = F_mpi_check(m, l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			v58 = v12
			return v58
		} else {
			v19 = F_mpi_check(m, v9)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					v58 = v12
					return v58
				} else {
					v23 = F_mpi_check(m, v10)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						if v23 == int32(0) {
							v58 = v12
							return v58
						} else {
							v27 = int32(1)
							if v11 <= v27 {
								v30 = v27
							} else {
								v30 = v11
							}
							v31 = F_palloc(m, v30)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
								v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								v40 = m.Env.Pgmem_bn_op(m, int32(1), v34, v35, v36, v37, v38, v39, v31, v11)
								mBase = m.M
								if v40 < int32(0) {
									v51 = int32(-109)
									if v31 == int32(0) {
										v58 = v51
										return v58
									} else {
										if v30 != 0 {
											base.MemoryFill(m, v31, int32(0), v30)
										} else {
										}
										F_pfree(m, v31)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v58 = v51
											return v58
										}
									}
								} else {
									v44 = F_bytes_to_mpi(m, v31, v40)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v44
										if v44 != 0 {
											v49 = int32(0)
										} else {
											v49 = int32(-109)
										}
										v51 = v49
										if v31 == int32(0) {
											v58 = v51
											return v58
										} else {
											if v30 != 0 {
												base.MemoryFill(m, v31, int32(0), v30)
											} else {
											}
											F_pfree(m, v31)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												v58 = v51
												return v58
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
}
func F_pgp_set_unicode_mode(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = base.B2i32(l1 != v3)
	return v3
}
func F_pgp_skip_packet(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	goto L1
L1:
	;
	v14 = F_pullf_read(m, l0, int32(_a_F_pgp_skip_packet_0), v6+int32(12))
	v17 = m.ExcPending
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	m.G0 = v6 + int32(16)
	return v14
L3:
	;
	return int32(0)
L4:
	;
	if int32(0) < v14 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L2
}
