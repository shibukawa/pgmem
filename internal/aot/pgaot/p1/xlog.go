package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_XlogReadTwoPhaseData(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int64
	_ = v121
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	v1 = l0
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = int32(395)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = int32(396)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_XlogReadTwoPhaseData[0]))
	v22 = F_XLogReaderAllocate(m, v18, v7+int32(-16), int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return
	} else {
		if v22 != 0 {
			F_XLogBeginRead(m, v22, v1)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v28 = F_XLogReadRecord(m, v22, v7+int32(-4))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					if v28 == int32(0) {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							v39 = base.I32_wrap_i64(int64(base.Ui64(v1) >> (uint(int64(32)) % 64)))
							F_errcode_for_file_access(m)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								v42 = base.I32_wrap_i64(v1)
								if v32 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v39
									*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v42
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v101
									F_errmsg(m, int32(_a_F_XlogReadTwoPhaseData_0), v7+int32(-48))
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_XlogReadTwoPhaseData_1), int32(1430), int32(_a_F_XlogReadTwoPhaseData_2))
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v42
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v39
									F_errmsg(m, int32(_a_F_XlogReadTwoPhaseData_3), v9)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_XlogReadTwoPhaseData_1), int32(1435), int32(_a_F_XlogReadTwoPhaseData_2))
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
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
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v22)+96))
						v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+49)))
						if v54 != int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return
								} else {
									*(*uint32)(unsafe.Add(mBase, uint32(v9)+36)) = uint32(v1)
									v121 = int64(base.Ui64(v1) >> (uint(int64(32)) % 64))
									*(*uint32)(unsafe.Add(mBase, uint32(v9)+32)) = uint32(v121)
									F_errmsg(m, int32(_a_F_XlogReadTwoPhaseData_4), v7+int32(-32))
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_XlogReadTwoPhaseData_1), int32(1443), int32(_a_F_XlogReadTwoPhaseData_2))
										mBase = m.M
										v132 = m.ExcPending
										if v132 != 0 {
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
							v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+48)))
							if v57&int32(112) != int32(16) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return
									} else {
										*(*uint32)(unsafe.Add(mBase, uint32(v9)+36)) = uint32(v1)
										v121 = int64(base.Ui64(v1) >> (uint(int64(32)) % 64))
										*(*uint32)(unsafe.Add(mBase, uint32(v9)+32)) = uint32(v121)
										F_errmsg(m, int32(_a_F_XlogReadTwoPhaseData_4), v7+int32(-32))
										mBase = m.M
										v127 = m.ExcPending
										if v127 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_XlogReadTwoPhaseData_1), int32(1443), int32(_a_F_XlogReadTwoPhaseData_2))
											mBase = m.M
											v132 = m.ExcPending
											if v132 != 0 {
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
								if l2 != 0 {
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v53)+68))
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v62
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v22)+96))
									v65 = v64
								} else {
									v65 = v53
								}
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+68))
								v67 = F_palloc(m, v66)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = v67
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v22)+96))
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+68))
									if v71 != 0 {
										v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+64))
										base.MemoryCopy(m, v67, v72, v71)
									} else {
									}
									F_XLogReaderFree(m, v22)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return
									} else {
										m.G0 = v9 - int32(-64)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return
			} else {
				F_errcode(m, int32(_a_F_XlogReadTwoPhaseData_5))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_XlogReadTwoPhaseData_6), int32(0))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return
					} else {
						F_errdetail(m, int32(_a_F_XlogReadTwoPhaseData_7), int32(0))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_XlogReadTwoPhaseData_1), int32(1419), int32(_a_F_XlogReadTwoPhaseData_2))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
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
		}
	}
}
func F_xlog_decode(m *base.Module, l0 int32, l1 int32) {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+96))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+48)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferProcessXid(m, v16, v17, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		v22 = v15 & int32(240)
		v24 = int32(base.Ui32(v22) >> (uint(int32(4)) % 32))
		if base.Ui32(int32(14)) < base.Ui32(v24) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v22
				F_errmsg_internal(m, int32(_a_F_xlog_decode_0), v10)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_xlog_decode_1), int32(193), int32(_a_F_xlog_decode_2))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = int32(1) << (uint(v24) % 32)
			if v28&int32(_a_F_xlog_decode_3) != 0 {
				m.G0 = v10 + int32(16)
				return
			} else {
				if v28&int32(513) != 0 {
					v71 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v72 <= int32(1) {
						v75 = F_SnapBuildRestore(m, v12, v71)
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return
						} else {
							m.G0 = v10 + int32(16)
							return
						}
					} else {
						F_SnapBuildSerialize(m, v12, v71)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return
						} else {
							m.G0 = v10 + int32(16)
							return
						}
					}
				} else {
					if v24 != int32(6) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v22
							F_errmsg_internal(m, int32(_a_F_xlog_decode_0), v10)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_xlog_decode_1), int32(193), int32(_a_F_xlog_decode_2))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+96))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+64))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
						if int32(1) < v38 {
							m.G0 = v10 + int32(16)
							return
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								F_errcode(m, int32(325))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_xlog_decode_4), int32(0))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_xlog_decode_1), int32(177), int32(_a_F_xlog_decode_2))
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
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
					}
				}
			}
		}
	}
}
func F_xlog_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int64
	_ = v108
	var v111 int64
	_ = v111
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v197 int64
	_ = v197
	var v199 int32
	_ = v199
	var v205 int64
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int64
	_ = v228
	var v230 int32
	_ = v230
	var v231 int64
	_ = v231
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
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
	var v274 int32
	_ = v274
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v320 int64
	_ = v320
	var v322 int32
	_ = v322
	var v328 int64
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int64
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int64
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
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
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v497 int64
	_ = v497
	var v499 int64
	_ = v499
	var v501 int64
	_ = v501
	var v510 int64
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v603 int64
	_ = v603
	var v605 int32
	_ = v605
	var v606 int64
	_ = v606
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	v13 = m.G0
	v15 = v13 - int32(176)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v18 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+48)))
	switch int32(base.Ui32((v18+int32(112))&int32(240)) >> (uint(int32(4)) % 32)) {
	case 0:
		goto L7
	default:
		goto L6
	case 4:
		goto L5
	case 7:
		goto L9
	case 8:
		goto L8
	case 10:
		goto L10
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L11
	} else {
		goto L149
	}
L2:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L11
	} else {
		goto L146
	}
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L11
	} else {
		goto L143
	}
L4:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L11
	} else {
		goto L140
	}
L5:
	;
	m.G0 = v15 + int32(176)
	return
L6:
	;
	v365 = v18 & int32(-16)
	if v365 == int32(32) {
		goto L5
	} else {
		goto L81
	}
L7:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)+8))
	v343 = F_GetCurrentReplayRecPtr(m, v15+int32(88))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L11
	} else {
		goto L76
	}
L8:
	;
	v216 = int32(88)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	base.MemoryCopy(m, v15+v216, v218, v216)
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v226 = F_LWLockAcquire(m, v222+int32(384), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L11
	} else {
		goto L44
	}
L9:
	;
	v47 = int32(88)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	base.MemoryCopy(m, v15+v47, v49, v47)
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v57 = F_LWLockAcquire(m, v53+int32(384), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L11
	} else {
		goto L14
	}
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v32 = F_LWLockAcquire(m, v28+int32(256), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	v34 = int32(_a_F_xlog_redo_0)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v26
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v42+int32(256))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	goto L5
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[1]))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v15)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v60)+8)) = v61
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v64+int32(384))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v74 = F_LWLockAcquire(m, v70+int32(256), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v76 = int32(_a_F_xlog_redo_0)
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[1]))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v15)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v78
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v85+int32(256))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+124))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v15)+128))
	F_MultiXactSetNextMXact(m, v90, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v15)+140))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v15)+144))
	F_MultiXactAdvanceOldest(m, v94, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v15)+132))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v15)+136))
	F_SetTransactionIdLimit(m, v98, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_xlog_redo[2])))
	if v103 != int32(1) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[3]))
	if v116 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[4]))
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v107)+152))
	if v108 == int64(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v107)+160))
	if v111 == int64(0) {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v121 = F_PrescanPreparedTransactions(m, v15+int32(84), v15+int32(80))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L11
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v172 = F_LWLockAcquire(m, v168+int32(1152), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L11
	} else {
		goto L34
	}
L28:
	;
	F_StandbyRecoverPreparedTransactions(m)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L11
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v121
	v126 = base.I32_wrap_i64(v61)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v126
	*(*int64)(unsafe.Add(mBase, uint32(v15)+52)) = int64(8589934592)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v130
	v134 = v126
	goto L30
L30:
	;
	v145 = v134 - int32(1)
	if base.Ui32(v145) < base.Ui32(int32(3)) {
		v134 = v145
		goto L30
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v145
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v15)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v149
	F_ProcArrayApplyRecoveryInfo(m, v15+int32(48))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L11
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	goto L27
L34:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v175)+64)) = v61
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v178+int32(1152))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L11
	} else {
		goto L35
	}
L35:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[5]))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v184)+440)) = int32(1)
	if v185 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[5]))
	F_s_lock(m, v189+int32(440), int32(_a_F_xlog_redo_1), int32(_a_F_xlog_redo_2), int32(_a_F_xlog_redo_3))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L11
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v197 = *(*int64)(unsafe.Add(mBase, uint32(v15)+112))
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+440)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v199)+208)) = v197
	v205 = F_GetCurrentReplayRecPtr(m, v15+int32(48))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L11
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	if v207 != v208 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	F_RecoveryRestartPoint(m, v15+int32(88), l0)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L11
	} else {
		goto L42
	}
L42:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	goto L5
L44:
	;
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v15)+112))
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[1]))
	v231 = *(*int64)(unsafe.Add(mBase, uint32(v230)+8))
	if base.Ui64(v231) < base.Ui64(v228) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v230)+8)) = v228
	goto L47
L46:
	;
	goto L47
L47:
	;
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v235+int32(384))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L11
	} else {
		goto L48
	}
L48:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v15)+124))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v15)+128))
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v247 = F_LWLockAcquire(m, v243+int32(1664), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L11
	} else {
		goto L49
	}
L49:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[6]))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	if v251-v240 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250))) = v240
	goto L52
L51:
	;
	goto L52
L52:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	if v256-v241 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+4)) = v241
	goto L55
L54:
	;
	goto L55
L55:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v262+int32(1664))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v15)+140))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v15)+144))
	F_MultiXactAdvanceOldest(m, v267, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[1]))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)+16))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v15)+132))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v274))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v273)) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v286 != 0 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v286 = base.B2i32(base.Ui32(v273) < base.Ui32(v274))
	goto L58
L60:
	;
	goto L61
L61:
	;
	v286 = int32(base.Ui32(v273-v274) >> (uint(int32(31)) % 32))
	goto L58
L62:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v15)+136))
	F_SetTransactionIdLimit(m, v274, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L11
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v295 = F_LWLockAcquire(m, v291+int32(1152), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L11
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v298)+64)) = v228
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v301+int32(1152))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L11
	} else {
		goto L67
	}
L67:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[5]))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v307)+440)) = int32(1)
	if v308 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[5]))
	F_s_lock(m, v312+int32(440), int32(_a_F_xlog_redo_1), int32(_a_F_xlog_redo_4), int32(_a_F_xlog_redo_3))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L11
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v320 = *(*int64)(unsafe.Add(mBase, uint32(v15)+112))
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v322)+440)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v322)+208)) = v320
	v328 = F_GetCurrentReplayRecPtr(m, v15+int32(48))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L11
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	if v330 != v331 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	F_RecoveryRestartPoint(m, v15+int32(88), l0)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L11
	} else {
		goto L74
	}
L74:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L11
	} else {
		goto L75
	}
L75:
	;
	goto L5
L76:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	if v340 == v345 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v340
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v352
	F_errmsg(m, int32(_a_F_xlog_redo_5), v15+int32(32))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L11
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_xlog_redo_1), int32(_a_F_xlog_redo_6), int32(_a_F_xlog_redo_3))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	v369 = v365 & int32(255)
	if base.B2i32(v369 == int32(64))|base.B2i32(v369 == int32(112)) != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	switch v369 - int32(80) {
	case 0:
		goto L5
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L83
	case 16:
		goto L84
	default:
		goto L85
	}
L83:
	;
	if int32(-113) < v18 {
		goto L5
	} else {
		goto L129
	}
L84:
	;
	v443 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+25)))
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+24)))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v444)+20))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v444)+16))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v444)+12))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v444)+8))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v444)+4))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	v454 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_xlog_redo[7])))
	if v454 != int32(1) {
		goto L102
	} else {
		goto L103
	}
L85:
	;
	v378 = v369 - int32(160)
	v379 = int32(0)
	if base.B2i32(v378 == v379)|base.B2i32(v378 == int32(16)) == v379 {
		goto L83
	} else {
		goto L86
	}
L86:
	;
	v386 = int32(0)
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	if v387 < v386 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	v395 = v17
	v396 = int32(0)
	v397 = v386
	goto L88
L88:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395+v396*int32(52))+105)))
	if v408 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L5
L90:
	;
	v438 = v397 + int32(1)
	v440 = v438 & int32(255)
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v436)+72))
	if v440 <= v441 {
		v395 = v436
		v396 = v440
		v397 = v438
		goto L88
	} else {
		goto L101
	}
L91:
	;
	if v365 != int32(-80) {
		v436 = v395
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v428 = F_XLogReadBufferForRedo(m, l0, v397&int32(255), v15+int32(88))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L11
	} else {
		goto L98
	}
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L11
	} else {
		goto L95
	}
L95:
	;
	F_errmsg_internal(m, int32(_a_F_xlog_redo_7), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L11
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_xlog_redo_1), int32(_a_F_xlog_redo_8), int32(_a_F_xlog_redo_3))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L11
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	if v428 != int32(2) {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	F_UnlockReleaseBuffer(m, v432)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L11
	} else {
		goto L100
	}
L100:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v436 = v435
	goto L90
L101:
	;
	goto L89
L102:
	;
	v475 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v479 = F_LWLockAcquire(m, v475+int32(1152), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L11
	} else {
		goto L107
	}
L103:
	;
	v458 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[3]))
	if base.B2i32(base.Ui32(v458) < base.Ui32(int32(2)))|base.B2i32(int32(1) < v447) != 0 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v465 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[8]))
	if v465 < int32(2) {
		goto L102
	} else {
		goto L105
	}
L105:
	;
	v470 = int32(0)
	v472 = F_InvalidateObsoleteReplicationSlots(m, int32(4), int64(0), v470, v470)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L11
	} else {
		goto L106
	}
L106:
	;
	goto L102
L107:
	;
	v482 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v482)+196)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v482)+192)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v482)+188)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v482)+184)) = v451
	*(*int32)(unsafe.Add(mBase, uint32(v482)+180)) = v452
	v489 = v446 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v482)+176)) = uint8(v489)
	*(*int32)(unsafe.Add(mBase, uint32(v482)+172)) = v447
	v493 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_xlog_redo[9])))
	if v493 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	if base.B2i32(v501 == int64(0))|base.B2i32(base.Ui64(v443) <= base.Ui64(v501)) == int32(0) {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	v497 = *(*int64)(unsafe.Add(mBase, _c_F_xlog_redo[10]))
	v501 = v497
	goto L108
L110:
	;
	goto L111
L111:
	;
	v499 = *(*int64)(unsafe.Add(mBase, uint32(v482)+136))
	*(*int64)(unsafe.Add(mBase, _c_F_xlog_redo[10])) = v499
	v501 = v499
	goto L108
L112:
	;
	v510 = F_GetCurrentReplayRecPtr(m, v15+int32(88))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L11
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v520 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[11]))
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520)+24)))
	v523 = v445 & int32(1)
	if v523 != 0 {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v513 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v513)+136)) = v443
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v513)+144)) = v515
	goto L114
L116:
	;
	v566 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[4]))
	*(*uint8)(unsafe.Add(mBase, uint32(v566)+200)) = uint8(v523)
	v569 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[12]))
	F_update_controlfile(m, v569, v566)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L11
	} else {
		goto L126
	}
L117:
	;
	if v521&int32(1) != 0 {
		goto L116
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	if v521&int32(1) == int32(0) {
		goto L116
	} else {
		goto L122
	}
L120:
	;
	F_ActivateCommitTs(m)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L11
	} else {
		goto L121
	}
L121:
	;
	goto L116
L122:
	;
	v533 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v537 = F_LWLockAcquire(m, v533+int32(_a_F_xlog_redo_9), int32(0))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L11
	} else {
		goto L123
	}
L123:
	;
	v540 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[11]))
	v541 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v540)+16)) = uint16(v541)
	*(*int64)(unsafe.Add(mBase, uint32(v540)+8)) = int64(-9223372036854775807 - 1)
	*(*int32)(unsafe.Add(mBase, uint32(v540))) = v541
	*(*uint8)(unsafe.Add(mBase, uint32(v540)+24)) = uint8(v541)
	v550 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v550)+40)) = int64(0)
	v556 = F_SlruScanDirectory(m, int32(_a_F_xlog_redo_10), int32(290), v541)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L11
	} else {
		goto L124
	}
L124:
	;
	v559 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v559+int32(_a_F_xlog_redo_9))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L11
	} else {
		goto L125
	}
L125:
	;
	goto L116
L126:
	;
	v573 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v573+int32(1152))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L11
	} else {
		goto L127
	}
L127:
	;
	F_CheckRequiredParameterValues(m)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L11
	} else {
		goto L128
	}
L128:
	;
	goto L5
L129:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583))))
	v586 = v584 & int32(1)
	if v586 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v590 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[5]))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v590)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v590)+440)) = int32(1)
	if v591 != 0 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	goto L132
L132:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_xlog_redo[13])) = uint8(v586)
	goto L5
L133:
	;
	v595 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[5]))
	F_s_lock(m, v595+int32(440), int32(_a_F_xlog_redo_1), int32(_a_F_xlog_redo_11), int32(_a_F_xlog_redo_3))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L11
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v603 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v605 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[5]))
	v606 = *(*int64)(unsafe.Add(mBase, uint32(v605)+432))
	if base.Ui64(v606) < base.Ui64(v603) {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	goto L135
L137:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v605)+432)) = v603
	goto L139
L138:
	;
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v605)+440)) = int32(0)
	goto L132
L140:
	;
	F_errmsg(m, int32(_a_F_xlog_redo_12), int32(0))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L11
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_xlog_redo_1), int32(_a_F_xlog_redo_13), int32(_a_F_xlog_redo_3))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L11
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v207
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v649
	F_errmsg(m, int32(_a_F_xlog_redo_14), v15)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L11
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_xlog_redo_1), int32(_a_F_xlog_redo_15), int32(_a_F_xlog_redo_3))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L11
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v330
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v664
	F_errmsg(m, int32(_a_F_xlog_redo_16), v15+int32(16))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L11
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_xlog_redo_1), int32(_a_F_xlog_redo_17), int32(_a_F_xlog_redo_3))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L11
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
	F_errmsg_internal(m, int32(_a_F_xlog_redo_18), int32(0))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L11
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_xlog_redo_1), int32(_a_F_xlog_redo_19), int32(_a_F_xlog_redo_3))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L11
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
