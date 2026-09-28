package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int64
	_ = v122
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	v1 = l0
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = int32(414)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = int32(415)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = int32(416)
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
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v102
									F_errmsg(m, int32(_a_F_XlogReadTwoPhaseData_0), v7+int32(-48))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_XlogReadTwoPhaseData_1), int32(1444), int32(_a_F_XlogReadTwoPhaseData_2))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
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
										F_errfinish(m, int32(_a_F_XlogReadTwoPhaseData_1), int32(1449), int32(_a_F_XlogReadTwoPhaseData_2))
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
							v117 = m.ExcPending
							if v117 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return
								} else {
									*(*uint32)(unsafe.Add(mBase, uint32(v9)+36)) = uint32(v1)
									v122 = int64(base.Ui64(v1) >> (uint(int64(32)) % 64))
									*(*uint32)(unsafe.Add(mBase, uint32(v9)+32)) = uint32(v122)
									F_errmsg(m, int32(_a_F_XlogReadTwoPhaseData_4), v7+int32(-32))
									mBase = m.M
									v128 = m.ExcPending
									if v128 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_XlogReadTwoPhaseData_1), int32(1457), int32(_a_F_XlogReadTwoPhaseData_2))
										mBase = m.M
										v133 = m.ExcPending
										if v133 != 0 {
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
								v117 = m.ExcPending
								if v117 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return
									} else {
										*(*uint32)(unsafe.Add(mBase, uint32(v9)+36)) = uint32(v1)
										v122 = int64(base.Ui64(v1) >> (uint(int64(32)) % 64))
										*(*uint32)(unsafe.Add(mBase, uint32(v9)+32)) = uint32(v122)
										F_errmsg(m, int32(_a_F_XlogReadTwoPhaseData_4), v7+int32(-32))
										mBase = m.M
										v128 = m.ExcPending
										if v128 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_XlogReadTwoPhaseData_1), int32(1457), int32(_a_F_XlogReadTwoPhaseData_2))
											mBase = m.M
											v133 = m.ExcPending
											if v133 != 0 {
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
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v53)+68))
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v63
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v22)+96))
									v66 = v65
								} else {
									v66 = v53
								}
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+68))
								v68 = F_palloc_mul(m, int32(1), v67)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = v68
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v22)+96))
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+68))
									if v72 != 0 {
										v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)+64))
										base.MemoryCopy(m, v68, v73, v72)
									} else {
									}
									F_XLogReaderFree(m, v22)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
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
			v83 = m.ExcPending
			if v83 != 0 {
				return
			} else {
				F_errcode(m, int32(_a_F_XlogReadTwoPhaseData_5))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_XlogReadTwoPhaseData_6), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						v93 = F_errdetail(m, int32(_a_F_XlogReadTwoPhaseData_7), int32(0))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_XlogReadTwoPhaseData_1), int32(1433), int32(_a_F_XlogReadTwoPhaseData_2))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	v8 = m.G0
	v10 = v8 - int32(32)
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
		switch int32(base.Ui32(v22)>>(uint(int32(4))%32)) - int32(1) {
		case 0, 1, 2, 3, 4, 5, 6, 7, 9, 10, 12, 13:
			m.G0 = v10 + int32(32)
			return
		default:
			v61 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			if v62 <= int32(1) {
				v65 = F_SnapBuildRestore(m, v12, v61)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					m.G0 = v10 + int32(32)
					return
				}
			} else {
				F_SnapBuildSerialize(m, v12, v61)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					m.G0 = v10 + int32(32)
					return
				}
			}
		case 11:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v22
				F_errmsg_internal(m, int32(_a_F_xlog_decode_0), v10+int32(16))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_xlog_decode_1), int32(189), int32(_a_F_xlog_decode_2))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 14:
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+96))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+64))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v30 & int32(1)
				F_errmsg_internal(m, int32(_a_F_xlog_decode_3), v10)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_xlog_decode_1), int32(172), int32(_a_F_xlog_decode_2))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
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
func F_xlog_redo(m *base.Module, l0 int32) {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v108 int64
	_ = v108
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int64
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int64
	_ = v199
	var v201 int32
	_ = v201
	var v202 int64
	_ = v202
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int64
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int64
	_ = v230
	var v234 int32
	_ = v234
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
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v279 int64
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int64
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int64
	_ = v315
	var v317 int32
	_ = v317
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int64
	_ = v402
	var v403 int64
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v426 int64
	_ = v426
	var v428 int64
	_ = v428
	var v430 int64
	_ = v430
	var v439 int64
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v531 int64
	_ = v531
	var v533 int32
	_ = v533
	var v534 int64
	_ = v534
	var v537 int32
	_ = v537
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	v10 = m.G0
	v12 = v10 - int32(192)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+48)))
	switch int32(base.Ui32((v15+int32(112))&int32(240)) >> (uint(int32(4)) % 32)) {
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
	v658 = m.ExcPending
	if v658 != 0 {
		goto L11
	} else {
		goto L163
	}
L2:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L11
	} else {
		goto L160
	}
L3:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L11
	} else {
		goto L157
	}
L4:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L11
	} else {
		goto L154
	}
L5:
	;
	m.G0 = v12 + int32(192)
	return
L6:
	;
	v315 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v317 = v15 & int32(-16)
	if base.Ui32(v317) <= base.Ui32(int32(111)) {
		goto L74
	} else {
		goto L75
	}
L7:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)+8))
	v294 = F_GetCurrentReplayRecPtr(m, v12+int32(96))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L11
	} else {
		goto L68
	}
L8:
	;
	v187 = int32(96)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	base.MemoryCopy(m, v12+v187, v189, v187)
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v197 = F_LWLockAcquire(m, v193+int32(384), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L11
	} else {
		goto L40
	}
L9:
	;
	v44 = int32(96)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	base.MemoryCopy(m, v12+v44, v46, v44)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v54 = F_LWLockAcquire(m, v50+int32(384), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L11
	} else {
		goto L14
	}
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v29 = F_LWLockAcquire(m, v25+int32(256), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	v31 = int32(_a_F_xlog_redo_0)
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v23
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = int32(0)
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v39+int32(256))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	goto L5
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[1]))
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v12)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v57)+8)) = v58
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v61+int32(384))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v71 = F_LWLockAcquire(m, v67+int32(256), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v73 = int32(_a_F_xlog_redo_0)
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[1]))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v12)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v75
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = int32(0)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v82+int32(256))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v12)+140))
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v12)+144))
	F_MultiXactSetNextMXact(m, v87, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v12)+164))
	F_MultiXactAdvanceOldest(m, v91, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v12)+156))
	F_SetTransactionIdLimit(m, v95, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_xlog_redo[2])))
	if v100 != int32(1) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[3]))
	if v113 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[4]))
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v104)+160))
	if v105 == int64(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v104)+168))
	if v108 == int64(0) {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v118 = F_PrescanPreparedTransactions(m, v12+int32(92), v12+int32(88))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L11
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v163 = F_LWLockAcquire(m, v159+int32(1152), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L11
	} else {
		goto L34
	}
L28:
	;
	F_StandbyRecoverPreparedTransactions(m)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L11
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v118
	v123 = base.I32_wrap_i64(v58)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v123
	*(*int64)(unsafe.Add(mBase, uint32(v12)+60)) = int64(8589934592)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v127
	v131 = v123
	goto L30
L30:
	;
	v139 = v131 - int32(1)
	if base.Ui32(v139) < base.Ui32(int32(3)) {
		v131 = v139
		goto L30
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v139
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v143
	F_ProcArrayApplyRecoveryInfo(m, v12+int32(56))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
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
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v166)+72)) = v58
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v169+int32(1152))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L11
	} else {
		goto L35
	}
L35:
	;
	v176 = F_GetCurrentReplayRecPtr(m, v12+int32(56))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	if v178 != v179 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	F_RecoveryRestartPoint(m, v12+int32(96), l0)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	goto L5
L40:
	;
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v12)+128))
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[1]))
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v201)+8))
	if base.Ui64(v202) < base.Ui64(v199) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v201)+8)) = v199
	goto L43
L42:
	;
	goto L43
L43:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v206+int32(384))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L11
	} else {
		goto L44
	}
L44:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v12)+140))
	v212 = *(*int64)(unsafe.Add(mBase, uint32(v12)+144))
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v218 = F_LWLockAcquire(m, v214+int32(1664), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L11
	} else {
		goto L45
	}
L45:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[5]))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	if v222-v211 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v221))) = v211
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[5]))
	v229 = v228
	goto L48
L47:
	;
	v229 = v221
	goto L48
L48:
	;
	v230 = *(*int64)(unsafe.Add(mBase, uint32(v229)+8))
	if base.Ui64(v230) < base.Ui64(v212) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v229)+8)) = v212
	goto L51
L50:
	;
	goto L51
L51:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v234+int32(1664))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L11
	} else {
		goto L52
	}
L52:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v12)+164))
	F_MultiXactAdvanceOldest(m, v239, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L11
	} else {
		goto L53
	}
L53:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[1]))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+16))
	v246 = int32(3)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
	if base.B2i32(base.Ui32(v245) < base.Ui32(v246))|base.B2i32(base.Ui32(v248) < base.Ui32(v246)) == int32(0) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v266 = F_LWLockAcquire(m, v262+int32(1152), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L11
	} else {
		goto L62
	}
L55:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v12)+156))
	F_SetTransactionIdLimit(m, v248, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L11
	} else {
		goto L61
	}
L56:
	;
	if v245-v248 < int32(0) {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if base.Ui32(v248) <= base.Ui32(v245) {
		goto L54
	} else {
		goto L60
	}
L59:
	;
	goto L54
L60:
	;
	goto L55
L61:
	;
	goto L54
L62:
	;
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v269)+72)) = v199
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v272+int32(1152))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	v279 = F_GetCurrentReplayRecPtr(m, v12+int32(56))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L11
	} else {
		goto L64
	}
L64:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	if v281 != v282 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	F_RecoveryRestartPoint(m, v12+int32(96), l0)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L11
	} else {
		goto L67
	}
L67:
	;
	goto L5
L68:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v12)+96))
	if v291 == v296 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L11
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v291
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v12)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v303
	F_errmsg(m, int32(_a_F_xlog_redo_1), v12+int32(48))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L11
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_xlog_redo_2), int32(_a_F_xlog_redo_3), int32(_a_F_xlog_redo_4))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L11
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	v330 = v317 & int32(255)
	switch v330 - int32(80) {
	case 0:
		goto L5
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L80
	case 16:
		goto L81
	default:
		goto L82
	}
L74:
	;
	if v317 == int32(32) {
		goto L5
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	if base.B2i32(v317 == int32(112))|base.B2i32(v317 == int32(192)) != 0 {
		goto L5
	} else {
		goto L79
	}
L77:
	;
	if v317 != int32(64) {
		goto L73
	} else {
		goto L78
	}
L78:
	;
	goto L5
L79:
	;
	goto L73
L80:
	;
	if base.I32_extend8_s(v15) <= int32(-113) {
		goto L121
	} else {
		goto L122
	}
L81:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398)+25)))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v398)+20))
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398)+24)))
	v402 = *(*int64)(unsafe.Add(mBase, uint32(v398)))
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v398)+8))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v398)+16))
	v406 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v410 = F_LWLockAcquire(m, v406+int32(1152), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L11
	} else {
		goto L99
	}
L82:
	;
	v334 = v330 - int32(160)
	v335 = int32(0)
	if base.B2i32(v334 == v335)|base.B2i32(v334 == int32(16)) == v335 {
		goto L80
	} else {
		goto L83
	}
L83:
	;
	v342 = int32(0)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	if v343 < v342 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	v349 = v14
	v350 = v342
	v351 = int32(0)
	goto L85
L85:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349+v351*int32(52))+105)))
	if v359 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L5
L87:
	;
	v393 = v350 + int32(1)
	v395 = v393 & int32(255)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v391)+72))
	if v395 <= v396 {
		v349 = v391
		v350 = v393
		v351 = v395
		goto L85
	} else {
		goto L98
	}
L88:
	;
	if v317&int32(255) != int32(176) {
		v391 = v349
		goto L87
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v383 = F_XLogReadBufferForRedo(m, l0, v350&int32(255), v12+int32(96))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L11
	} else {
		goto L95
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L11
	} else {
		goto L92
	}
L92:
	;
	F_errmsg_internal(m, int32(_a_F_xlog_redo_5), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L11
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_xlog_redo_2), int32(_a_F_xlog_redo_6), int32(_a_F_xlog_redo_4))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L11
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	if v383 != int32(2) {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v12)+96))
	F_UnlockReleaseBuffer(m, v387)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L11
	} else {
		goto L97
	}
L97:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v391 = v390
	goto L87
L98:
	;
	goto L86
L99:
	;
	v413 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v413)+204)) = v404
	*(*int64)(unsafe.Add(mBase, uint32(v413)+196)) = v403
	*(*int64)(unsafe.Add(mBase, uint32(v413)+188)) = v402
	v418 = v401 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v413)+184)) = uint8(v418)
	*(*int32)(unsafe.Add(mBase, uint32(v413)+180)) = v400
	v422 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_xlog_redo[6])))
	if v422 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	if base.B2i32(v430 == int64(0))|base.B2i32(base.Ui64(v315) <= base.Ui64(v430)) == int32(0) {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	v426 = *(*int64)(unsafe.Add(mBase, _c_F_xlog_redo[7]))
	v430 = v426
	goto L100
L102:
	;
	goto L103
L103:
	;
	v428 = *(*int64)(unsafe.Add(mBase, uint32(v413)+144))
	*(*int64)(unsafe.Add(mBase, _c_F_xlog_redo[7])) = v428
	v430 = v428
	goto L100
L104:
	;
	v439 = F_GetCurrentReplayRecPtr(m, v12+int32(96))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L11
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v449 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[8]))
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449)+24)))
	v452 = v399 & int32(1)
	if v452 != 0 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v442 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v442)+144)) = v315
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v12)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v442)+152)) = v444
	goto L106
L108:
	;
	v497 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[4]))
	*(*uint8)(unsafe.Add(mBase, uint32(v497)+208)) = uint8(v452)
	v500 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[9]))
	F_update_controlfile(m, v500, v497)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L11
	} else {
		goto L118
	}
L109:
	;
	if v450&int32(1) != 0 {
		goto L108
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	if v450&int32(1) == int32(0) {
		goto L108
	} else {
		goto L114
	}
L112:
	;
	F_ActivateCommitTs(m)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L11
	} else {
		goto L113
	}
L113:
	;
	goto L108
L114:
	;
	v462 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v466 = F_LWLockAcquire(m, v462+int32(_a_F_xlog_redo_7), int32(0))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L11
	} else {
		goto L115
	}
L115:
	;
	v468 = int32(_a_F_xlog_redo_8)
	v469 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[8]))
	v470 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v469))) = v470
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+24)) = uint8(v470)
	v475 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[8]))
	*(*uint16)(unsafe.Add(mBase, uint32(v475)+16)) = uint16(v470)
	*(*int64)(unsafe.Add(mBase, uint32(v475)+8)) = int64(-9223372036854775807 - 1)
	v481 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v481)+40)) = int64(0)
	v487 = F_SlruScanDirectory(m, int32(_a_F_xlog_redo_9), int32(302), v470)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L11
	} else {
		goto L116
	}
L116:
	;
	v490 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v490+int32(_a_F_xlog_redo_7))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L11
	} else {
		goto L117
	}
L117:
	;
	goto L108
L118:
	;
	v504 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v504+int32(1152))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L11
	} else {
		goto L119
	}
L119:
	;
	F_CheckRequiredParameterValues(m)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L11
	} else {
		goto L120
	}
L120:
	;
	goto L5
L121:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
	v518 = v516 & int32(1)
	if v518 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	if base.Ui32(v15) < base.Ui32(int32(240)) {
		goto L5
	} else {
		goto L134
	}
L124:
	;
	v522 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[10]))
	v525 = base.AtomicRmwXchg32(m, v522, int32(440), int32(1))
	if v525 != 0 {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	goto L126
L126:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_xlog_redo[11])) = uint8(v518)
	goto L5
L127:
	;
	F_s_lock(m, v522+int32(440), int32(_a_F_xlog_redo_10))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L11
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v531 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v533 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[10]))
	v534 = *(*int64)(unsafe.Add(mBase, uint32(v533)+432))
	if base.Ui64(v534) < base.Ui64(v531) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	goto L129
L131:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v533)+432)) = v531
	goto L133
L132:
	;
	goto L133
L133:
	;
	v537 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v533)+440)), uint32(v537))
	goto L126
L134:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547))))
	if v548&int32(1) != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v557 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L11
	} else {
		goto L141
	}
L136:
	;
	F_EnableLogicalDecoding(m)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L11
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	F_DisableLogicalDecoding(m)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L11
	} else {
		goto L140
	}
L139:
	;
	goto L135
L140:
	;
	goto L135
L141:
	;
	if v557 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v548 & int32(1)
	F_errmsg_internal(m, int32(_a_F_xlog_redo_11), v12)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L11
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v571 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_xlog_redo[12])))
	if v571 != int32(1) {
		goto L5
	} else {
		goto L147
	}
L145:
	;
	F_errfinish(m, int32(_a_F_xlog_redo_2), int32(_a_F_xlog_redo_12), int32(_a_F_xlog_redo_4))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L11
	} else {
		goto L146
	}
L146:
	;
	goto L144
L147:
	;
	v575 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[3]))
	if base.Ui32(v575) < base.Ui32(int32(2)) {
		goto L5
	} else {
		goto L148
	}
L148:
	;
	if v548&int32(1) == int32(0) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v584 = int32(0)
	v586 = F_InvalidateObsoleteReplicationSlots(m, int32(4), int64(0), v584, v584)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L11
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v589 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_xlog_redo[13])))
	if v589 != int32(1) {
		goto L5
	} else {
		goto L153
	}
L152:
	;
	goto L5
L153:
	;
	v593 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[14]))
	v595 = F_pgmem_kill(m, v593, int32(10))
	mBase = m.M
	goto L5
L154:
	;
	F_errmsg(m, int32(_a_F_xlog_redo_13), int32(0))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L11
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(_a_F_xlog_redo_2), int32(_a_F_xlog_redo_14), int32(_a_F_xlog_redo_4))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L11
	} else {
		goto L156
	}
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v178
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v626
	F_errmsg(m, int32(_a_F_xlog_redo_15), v12+int32(16))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L11
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_xlog_redo_2), int32(_a_F_xlog_redo_16), int32(_a_F_xlog_redo_4))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L11
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v281
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v643
	F_errmsg(m, int32(_a_F_xlog_redo_17), v12+int32(32))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L11
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_xlog_redo_2), int32(_a_F_xlog_redo_18), int32(_a_F_xlog_redo_4))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L11
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	F_errmsg_internal(m, int32(_a_F_xlog_redo_19), int32(0))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L11
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F_xlog_redo_2), int32(_a_F_xlog_redo_20), int32(_a_F_xlog_redo_4))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L11
	} else {
		goto L165
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
