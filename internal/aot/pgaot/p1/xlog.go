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
	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = int32(395)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = int32(396)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[138]))
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
									F_errmsg(m, int32(194341), v7+int32(-48))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return
									} else {
										F_errfinish(m, int32(475602), int32(1430), int32(482102))
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
									F_errmsg(m, int32(490360), v9)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return
									} else {
										F_errfinish(m, int32(475602), int32(1435), int32(482102))
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
									F_errmsg(m, int32(490299), v7+int32(-32))
									mBase = m.M
									v128 = m.ExcPending
									if v128 != 0 {
										return
									} else {
										F_errfinish(m, int32(475602), int32(1443), int32(482102))
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
										F_errmsg(m, int32(490299), v7+int32(-32))
										mBase = m.M
										v128 = m.ExcPending
										if v128 != 0 {
											return
										} else {
											F_errfinish(m, int32(475602), int32(1443), int32(482102))
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
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+64))
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+68))
									if v72 != 0 {
										v73 = F__emscripten_memcpy_bulkmem(m, v67, v71, v72)
										mBase = m.M
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
				F_errcode(m, int32(8389))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					F_errmsg(m, int32(12790), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						F_errdetail(m, int32(566435), int32(0))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							F_errfinish(m, int32(475602), int32(1419), int32(482102))
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
				F_errmsg_internal(m, int32(56251), v10)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errfinish(m, int32(476369), int32(193), int32(394365))
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
			if v28&int32(28094) != 0 {
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
							F_errmsg_internal(m, int32(56251), v10)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								F_errfinish(m, int32(476369), int32(193), int32(394365))
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
									F_errmsg(m, int32(16796), int32(0))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										F_errfinish(m, int32(476369), int32(177), int32(394365))
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
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
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
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v112 int64
	_ = v112
	var v117 int32
	_ = v117
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
	var v134 int32
	_ = v134
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v198 int64
	_ = v198
	var v200 int32
	_ = v200
	var v206 int64
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int64
	_ = v230
	var v232 int32
	_ = v232
	var v233 int64
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v322 int64
	_ = v322
	var v324 int32
	_ = v324
	var v330 int64
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int64
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int64
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
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
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v492 int64
	_ = v492
	var v494 int64
	_ = v494
	var v496 int64
	_ = v496
	var v502 int64
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v593 int32
	_ = v593
	var v594 int64
	_ = v594
	var v596 int32
	_ = v596
	var v597 int64
	_ = v597
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
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
	v670 = m.ExcPending
	if v670 != 0 {
		goto L11
	} else {
		goto L159
	}
L2:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L11
	} else {
		goto L156
	}
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L11
	} else {
		goto L153
	}
L4:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L11
	} else {
		goto L150
	}
L5:
	;
	m.G0 = v15 + int32(176)
	return
L6:
	;
	v367 = v18 & int32(-16)
	if v367 == int32(32) {
		goto L5
	} else {
		goto L89
	}
L7:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+8))
	v345 = F_GetCurrentReplayRecPtr(m, v15+int32(88))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L11
	} else {
		goto L84
	}
L8:
	;
	v217 = int32(88)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	goto L49
L9:
	;
	v47 = int32(88)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	goto L15
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v28 = *(*int32)(unsafe.Add(mBase, _consts[24]))
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
	v34 = int32(4338040)
	v35 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v26
	v38 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, _consts[24]))
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
	v54 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v58 = F_LWLockAcquire(m, v54+int32(384), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L11
	} else {
		goto L18
	}
L15:
	;
	v51 = F__emscripten_memcpy_bulkmem(m, v15+v47, v49, v47)
	mBase = m.M
	goto L17
L17:
	;
	goto L14
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v15)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v61)+8)) = v62
	v65 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v65+int32(384))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v75 = F_LWLockAcquire(m, v71+int32(256), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	v77 = int32(4338040)
	v78 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v15)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v79
	v82 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = int32(0)
	v86 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v86+int32(256))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v15)+124))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v15)+128))
	F_MultiXactSetNextMXact(m, v91, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v15)+140))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v15)+144))
	F_MultiXactAdvanceOldest(m, v95, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L11
	} else {
		goto L23
	}
L23:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v15)+132))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+136))
	F_SetTransactionIdLimit(m, v99, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, _consts[247])))
	if v104 != int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _consts[248]))
	if v117 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v108)+152))
	if v109 == int64(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v108)+160))
	if v112 == int64(0) {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v122 = F_PrescanPreparedTransactions(m, v15+int32(84), v15+int32(80))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L11
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v173 = F_LWLockAcquire(m, v169+int32(1152), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L11
	} else {
		goto L38
	}
L32:
	;
	F_StandbyRecoverPreparedTransactions(m)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L11
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v122
	v127 = base.I32_wrap_i64(v62)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v127
	*(*int64)(unsafe.Add(mBase, uint32(v15)+52)) = int64(8589934592)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v131
	v134 = v127
	goto L34
L34:
	;
	v146 = v134 - int32(1)
	if base.Ui32(v146) < base.Ui32(int32(3)) {
		v134 = v146
		goto L34
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v146
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v15)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v150
	F_ProcArrayApplyRecoveryInfo(m, v15+int32(48))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L11
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	goto L31
L38:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	*(*int64)(unsafe.Add(mBase, uint32(v176)+64)) = v62
	v179 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v179+int32(1152))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+440)) = int32(1)
	if v186 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	F_s_lock(m, v190+int32(440), int32(475016), int32(8393), int32(230373))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L11
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v15)+112))
	v200 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	*(*int32)(unsafe.Add(mBase, uint32(v200)+440)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v200)+208)) = v198
	v206 = F_GetCurrentReplayRecPtr(m, v15+int32(48))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L11
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	if v208 != v209 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	F_RecoveryRestartPoint(m, v15+int32(88), l0)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L11
	} else {
		goto L46
	}
L46:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	goto L5
L48:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v228 = F_LWLockAcquire(m, v224+int32(384), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L11
	} else {
		goto L52
	}
L49:
	;
	v221 = F__emscripten_memcpy_bulkmem(m, v15+v217, v219, v217)
	mBase = m.M
	goto L51
L51:
	;
	goto L48
L52:
	;
	v230 = *(*int64)(unsafe.Add(mBase, uint32(v15)+112))
	v232 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	v233 = *(*int64)(unsafe.Add(mBase, uint32(v232)+8))
	if base.Ui64(v233) < base.Ui64(v230) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v232)+8)) = v230
	goto L55
L54:
	;
	goto L55
L55:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v237+int32(384))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v15)+124))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v15)+128))
	v245 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v249 = F_LWLockAcquire(m, v245+int32(1664), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _consts[129]))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	if v253-v242 < int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v252))) = v242
	goto L60
L59:
	;
	goto L60
L60:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	if v258-v243 < int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v252)+4)) = v243
	goto L63
L62:
	;
	goto L63
L63:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v264+int32(1664))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L11
	} else {
		goto L64
	}
L64:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v15)+140))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v15)+144))
	F_MultiXactAdvanceOldest(m, v269, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+16))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v15)+132))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v276))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v275)) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v288 != 0 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v288 = base.B2i32(base.Ui32(v275) < base.Ui32(v276))
	goto L66
L68:
	;
	goto L69
L69:
	;
	v288 = int32(base.Ui32(v275-v276) >> (uint(int32(31)) % 32))
	goto L66
L70:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v15)+136))
	F_SetTransactionIdLimit(m, v276, v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L11
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v297 = F_LWLockAcquire(m, v293+int32(1152), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L11
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	v300 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	*(*int64)(unsafe.Add(mBase, uint32(v300)+64)) = v230
	v303 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v303+int32(1152))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L11
	} else {
		goto L75
	}
L75:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v309)+440)) = int32(1)
	if v310 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	F_s_lock(m, v314+int32(440), int32(475016), int32(8462), int32(230373))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L11
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v322 = *(*int64)(unsafe.Add(mBase, uint32(v15)+112))
	v324 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	*(*int32)(unsafe.Add(mBase, uint32(v324)+440)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v324)+208)) = v322
	v330 = F_GetCurrentReplayRecPtr(m, v15+int32(48))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L11
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	if v332 != v333 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	F_RecoveryRestartPoint(m, v15+int32(88), l0)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L11
	} else {
		goto L82
	}
L82:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L11
	} else {
		goto L83
	}
L83:
	;
	goto L5
L84:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	if v342 == v347 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L11
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v342
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v354
	F_errmsg(m, int32(400518), v15+int32(32))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L11
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(475016), int32(8508), int32(230373))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L11
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	v371 = v367 & int32(255)
	if v371 == int32(64) {
		goto L5
	} else {
		goto L90
	}
L90:
	;
	if v371 == int32(112) {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	v377 = v367 & int32(255)
	switch v377 - int32(80) {
	case 0:
		goto L5
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L92
	case 16:
		goto L93
	default:
		goto L94
	}
L92:
	;
	if int32(-113) < v18 {
		goto L5
	} else {
		goto L139
	}
L93:
	;
	v439 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440)+25)))
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440)+24)))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v440)+20))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v440)+16))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v440)+12))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v440)+8))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	v450 = int32(*(*uint8)(unsafe.Add(mBase, _consts[113])))
	if v450 != int32(1) {
		goto L111
	} else {
		goto L112
	}
L94:
	;
	switch v377 - int32(160) {
	case 0, 16:
		goto L95
	default:
		goto L92
	}
L95:
	;
	v382 = int32(0)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	if v383 < v382 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	v390 = v17
	v392 = int32(0)
	v393 = v382
	goto L97
L97:
	;
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390+v393*int32(52))+105)))
	if v404 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L5
L99:
	;
	v434 = v392 + int32(1)
	v436 = v434 & int32(255)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v432)+72))
	if v436 <= v437 {
		v390 = v432
		v392 = v434
		v393 = v436
		goto L97
	} else {
		goto L110
	}
L100:
	;
	if v367 != int32(-80) {
		v432 = v390
		goto L99
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v424 = F_XLogReadBufferForRedo(m, l0, v392&int32(255), v15+int32(88))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L11
	} else {
		goto L107
	}
L103:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L11
	} else {
		goto L104
	}
L104:
	;
	F_errmsg_internal(m, int32(389425), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L11
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(475016), int32(8547), int32(230373))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L11
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	if v424 != int32(2) {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	F_UnlockReleaseBuffer(m, v428)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L11
	} else {
		goto L109
	}
L109:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v432 = v431
	goto L99
L110:
	;
	goto L98
L111:
	;
	v470 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v474 = F_LWLockAcquire(m, v470+int32(1152), int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L11
	} else {
		goto L117
	}
L112:
	;
	v454 = *(*int32)(unsafe.Add(mBase, _consts[248]))
	if base.Ui32(v454) < base.Ui32(int32(2)) {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	if int32(1) < v443 {
		goto L111
	} else {
		goto L114
	}
L114:
	;
	v460 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	if v460 < int32(2) {
		goto L111
	} else {
		goto L115
	}
L115:
	;
	v465 = int32(0)
	v467 = F_InvalidateObsoleteReplicationSlots(m, int32(4), int64(0), v465, v465)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L11
	} else {
		goto L116
	}
L116:
	;
	goto L111
L117:
	;
	v477 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	*(*int32)(unsafe.Add(mBase, uint32(v477)+196)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v477)+192)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v477)+188)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v477)+184)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v477)+180)) = v448
	v484 = v442 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v477)+176)) = uint8(v484)
	*(*int32)(unsafe.Add(mBase, uint32(v477)+172)) = v443
	v488 = int32(*(*uint8)(unsafe.Add(mBase, _consts[249])))
	if v488 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	if v496 == int64(0) {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v492 = *(*int64)(unsafe.Add(mBase, _consts[227]))
	v496 = v492
	goto L118
L120:
	;
	goto L121
L121:
	;
	v494 = *(*int64)(unsafe.Add(mBase, uint32(v477)+136))
	*(*int64)(unsafe.Add(mBase, _consts[227])) = v494
	v496 = v494
	goto L118
L122:
	;
	v512 = *(*int32)(unsafe.Add(mBase, _consts[250]))
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+24)))
	v515 = v441 & int32(1)
	if v515 != 0 {
		goto L127
	} else {
		goto L128
	}
L123:
	;
	if base.Ui64(v439) <= base.Ui64(v496) {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v502 = F_GetCurrentReplayRecPtr(m, v15+int32(88))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L11
	} else {
		goto L125
	}
L125:
	;
	v505 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	*(*int64)(unsafe.Add(mBase, uint32(v505)+136)) = v439
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v505)+144)) = v507
	goto L122
L126:
	;
	v558 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	*(*uint8)(unsafe.Add(mBase, uint32(v558)+200)) = uint8(v515)
	v561 = *(*int32)(unsafe.Add(mBase, _consts[251]))
	F_update_controlfile(m, v561, v558)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L11
	} else {
		goto L136
	}
L127:
	;
	if v513&int32(1) != 0 {
		goto L126
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	if v513&int32(1) == int32(0) {
		goto L126
	} else {
		goto L132
	}
L130:
	;
	F_ActivateCommitTs(m)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L11
	} else {
		goto L131
	}
L131:
	;
	goto L126
L132:
	;
	v525 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v529 = F_LWLockAcquire(m, v525+int32(4992), int32(0))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L11
	} else {
		goto L133
	}
L133:
	;
	v532 = *(*int32)(unsafe.Add(mBase, _consts[250]))
	v533 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v532)+16)) = uint16(v533)
	*(*int64)(unsafe.Add(mBase, uint32(v532)+8)) = int64(-9223372036854775807 - 1)
	*(*int32)(unsafe.Add(mBase, uint32(v532))) = v533
	*(*uint8)(unsafe.Add(mBase, uint32(v532)+24)) = uint8(v533)
	v542 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int64)(unsafe.Add(mBase, uint32(v542)+40)) = int64(0)
	v548 = F_SlruScanDirectory(m, int32(4337604), int32(290), v533)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L11
	} else {
		goto L134
	}
L134:
	;
	v551 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v551+int32(4992))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L11
	} else {
		goto L135
	}
L135:
	;
	goto L126
L136:
	;
	v565 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v565+int32(1152))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L11
	} else {
		goto L137
	}
L137:
	;
	F_CheckRequiredParameterValues(m)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L11
	} else {
		goto L138
	}
L138:
	;
	goto L5
L139:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574))))
	v577 = v575 & int32(1)
	if v577 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v581 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v581)+440)) = int32(1)
	if v582 != 0 {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	goto L142
L142:
	;
	*(*uint8)(unsafe.Add(mBase, _consts[252])) = uint8(v577)
	goto L5
L143:
	;
	v586 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	F_s_lock(m, v586+int32(440), int32(475016), int32(8636), int32(230373))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L11
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v594 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v596 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	v597 = *(*int64)(unsafe.Add(mBase, uint32(v596)+432))
	if base.Ui64(v597) < base.Ui64(v594) {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	goto L145
L147:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v596)+432)) = v594
	goto L149
L148:
	;
	goto L149
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v596)+440)) = int32(0)
	goto L142
L150:
	;
	F_errmsg(m, int32(327925), int32(0))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L11
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(475016), int32(8346), int32(230373))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L11
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v208
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v640
	F_errmsg(m, int32(400585), v15)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L11
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(475016), int32(8405), int32(230373))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L11
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v332
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v655
	F_errmsg(m, int32(400951), v15+int32(16))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L11
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(475016), int32(8471), int32(230373))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L11
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_errmsg_internal(m, int32(301134), int32(0))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L11
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(475016), int32(8552), int32(230373))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L11
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
