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
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int64
	_ = v199
	var v201 int32
	_ = v201
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
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int64
	_ = v229
	var v231 int32
	_ = v231
	var v232 int64
	_ = v232
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int64
	_ = v323
	var v325 int32
	_ = v325
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
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int64
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
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v499 int64
	_ = v499
	var v501 int64
	_ = v501
	var v503 int64
	_ = v503
	var v512 int64
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v552 int32
	_ = v552
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v604 int32
	_ = v604
	var v605 int64
	_ = v605
	var v607 int32
	_ = v607
	var v608 int64
	_ = v608
	var v611 int32
	_ = v611
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v672 int32
	_ = v672
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
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
	v681 = m.ExcPending
	if v681 != 0 {
		goto L11
	} else {
		goto L149
	}
L2:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L11
	} else {
		goto L146
	}
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L11
	} else {
		goto L143
	}
L4:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
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
	v367 = v18 & int32(-16)
	if v367 == int32(32) {
		goto L5
	} else {
		goto L81
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
		goto L76
	}
L8:
	;
	v217 = int32(88)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	base.MemoryCopy(m, v15+v217, v219, v217)
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v227 = F_LWLockAcquire(m, v223+int32(384), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
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
	v187 = base.AtomicRmwXchg32(m, v184, int32(440), int32(1))
	if v187 != 0 {
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
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[5]))
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v15)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v198)+208)) = v199
	v201 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v198)+440)), uint32(v201))
	v206 = F_GetCurrentReplayRecPtr(m, v15+int32(48))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L11
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	if v208 != v209 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	F_RecoveryRestartPoint(m, v15+int32(88), l0)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L11
	} else {
		goto L42
	}
L42:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	goto L5
L44:
	;
	v229 = *(*int64)(unsafe.Add(mBase, uint32(v15)+112))
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[1]))
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v231)+8))
	if base.Ui64(v232) < base.Ui64(v229) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v231)+8)) = v229
	goto L47
L46:
	;
	goto L47
L47:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v236+int32(384))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L11
	} else {
		goto L48
	}
L48:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v15)+124))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v15)+128))
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v248 = F_LWLockAcquire(m, v244+int32(1664), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L11
	} else {
		goto L49
	}
L49:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[6]))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	if v252-v241 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v241
	goto L52
L51:
	;
	goto L52
L52:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	if v257-v242 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251)+4)) = v242
	goto L55
L54:
	;
	goto L55
L55:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v263+int32(1664))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v15)+140))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v15)+144))
	F_MultiXactAdvanceOldest(m, v268, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[1]))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+16))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v15)+132))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v275))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v274)) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v287 != 0 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v287 = base.B2i32(base.Ui32(v274) < base.Ui32(v275))
	goto L58
L60:
	;
	goto L61
L61:
	;
	v287 = int32(base.Ui32(v274-v275) >> (uint(int32(31)) % 32))
	goto L58
L62:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v15)+136))
	F_SetTransactionIdLimit(m, v275, v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L11
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v296 = F_LWLockAcquire(m, v292+int32(1152), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L11
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v299)+64)) = v229
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v302+int32(1152))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L11
	} else {
		goto L67
	}
L67:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[5]))
	v311 = base.AtomicRmwXchg32(m, v308, int32(440), int32(1))
	if v311 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[5]))
	F_s_lock(m, v313+int32(440), int32(_a_F_xlog_redo_1), int32(_a_F_xlog_redo_4), int32(_a_F_xlog_redo_3))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L11
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[5]))
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v15)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v322)+208)) = v323
	v325 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v322)+440)), uint32(v325))
	v330 = F_GetCurrentReplayRecPtr(m, v15+int32(48))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L11
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	if v332 != v333 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	F_RecoveryRestartPoint(m, v15+int32(88), l0)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L11
	} else {
		goto L74
	}
L74:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L11
	} else {
		goto L75
	}
L75:
	;
	goto L5
L76:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	if v342 == v347 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v342
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v354
	F_errmsg(m, int32(_a_F_xlog_redo_5), v15+int32(32))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L11
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_xlog_redo_1), int32(_a_F_xlog_redo_6), int32(_a_F_xlog_redo_3))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
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
	v371 = v367 & int32(255)
	if base.B2i32(v371 == int32(64))|base.B2i32(v371 == int32(112)) != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	switch v371 - int32(80) {
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
	v445 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446)+25)))
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446)+24)))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v446)+20))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v446)+16))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v446)+12))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v446)+8))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v446)+4))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v446)))
	v456 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_xlog_redo[7])))
	if v456 != int32(1) {
		goto L102
	} else {
		goto L103
	}
L85:
	;
	v380 = v371 - int32(160)
	v381 = int32(0)
	if base.B2i32(v380 == v381)|base.B2i32(v380 == int32(16)) == v381 {
		goto L83
	} else {
		goto L86
	}
L86:
	;
	v388 = int32(0)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	if v389 < v388 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	v397 = v17
	v398 = v388
	v399 = int32(0)
	goto L88
L88:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397+v399*int32(52))+105)))
	if v410 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L5
L90:
	;
	v440 = v398 + int32(1)
	v442 = v440 & int32(255)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v438)+72))
	if v442 <= v443 {
		v397 = v438
		v398 = v440
		v399 = v442
		goto L88
	} else {
		goto L101
	}
L91:
	;
	if v367 != int32(-80) {
		v438 = v397
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v430 = F_XLogReadBufferForRedo(m, l0, v398&int32(255), v15+int32(88))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L11
	} else {
		goto L98
	}
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L11
	} else {
		goto L95
	}
L95:
	;
	F_errmsg_internal(m, int32(_a_F_xlog_redo_7), int32(0))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L11
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_xlog_redo_1), int32(_a_F_xlog_redo_8), int32(_a_F_xlog_redo_3))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
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
	if v430 != int32(2) {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	F_UnlockReleaseBuffer(m, v434)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L11
	} else {
		goto L100
	}
L100:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v438 = v437
	goto L90
L101:
	;
	goto L89
L102:
	;
	v477 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v481 = F_LWLockAcquire(m, v477+int32(1152), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L11
	} else {
		goto L107
	}
L103:
	;
	v460 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[3]))
	if base.B2i32(base.Ui32(v460) < base.Ui32(int32(2)))|base.B2i32(int32(1) < v449) != 0 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v467 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[8]))
	if v467 < int32(2) {
		goto L102
	} else {
		goto L105
	}
L105:
	;
	v472 = int32(0)
	v474 = F_InvalidateObsoleteReplicationSlots(m, int32(4), int64(0), v472, v472)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L11
	} else {
		goto L106
	}
L106:
	;
	goto L102
L107:
	;
	v484 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v484)+196)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v484)+192)) = v451
	*(*int32)(unsafe.Add(mBase, uint32(v484)+188)) = v452
	*(*int32)(unsafe.Add(mBase, uint32(v484)+184)) = v453
	*(*int32)(unsafe.Add(mBase, uint32(v484)+180)) = v454
	v491 = v448 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v484)+176)) = uint8(v491)
	*(*int32)(unsafe.Add(mBase, uint32(v484)+172)) = v449
	v495 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_xlog_redo[9])))
	if v495 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	if base.B2i32(v503 == int64(0))|base.B2i32(base.Ui64(v445) <= base.Ui64(v503)) == int32(0) {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	v499 = *(*int64)(unsafe.Add(mBase, _c_F_xlog_redo[10]))
	v503 = v499
	goto L108
L110:
	;
	goto L111
L111:
	;
	v501 = *(*int64)(unsafe.Add(mBase, uint32(v484)+136))
	*(*int64)(unsafe.Add(mBase, _c_F_xlog_redo[10])) = v501
	v503 = v501
	goto L108
L112:
	;
	v512 = F_GetCurrentReplayRecPtr(m, v15+int32(88))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L11
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v522 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[11]))
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522)+24)))
	v525 = v447 & int32(1)
	if v525 != 0 {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v515 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v515)+136)) = v445
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v515)+144)) = v517
	goto L114
L116:
	;
	v568 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[4]))
	*(*uint8)(unsafe.Add(mBase, uint32(v568)+200)) = uint8(v525)
	v571 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[12]))
	F_update_controlfile(m, v571, v568)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L11
	} else {
		goto L126
	}
L117:
	;
	if v523&int32(1) != 0 {
		goto L116
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	if v523&int32(1) == int32(0) {
		goto L116
	} else {
		goto L122
	}
L120:
	;
	F_ActivateCommitTs(m)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L11
	} else {
		goto L121
	}
L121:
	;
	goto L116
L122:
	;
	v535 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	v539 = F_LWLockAcquire(m, v535+int32(_a_F_xlog_redo_9), int32(0))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L11
	} else {
		goto L123
	}
L123:
	;
	v542 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[11]))
	v543 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v542)+16)) = uint16(v543)
	*(*int64)(unsafe.Add(mBase, uint32(v542)+8)) = int64(-9223372036854775807 - 1)
	*(*int32)(unsafe.Add(mBase, uint32(v542))) = v543
	*(*uint8)(unsafe.Add(mBase, uint32(v542)+24)) = uint8(v543)
	v552 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v552)+40)) = int64(0)
	v558 = F_SlruScanDirectory(m, int32(_a_F_xlog_redo_10), int32(290), v543)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L11
	} else {
		goto L124
	}
L124:
	;
	v561 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v561+int32(_a_F_xlog_redo_9))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L11
	} else {
		goto L125
	}
L125:
	;
	goto L116
L126:
	;
	v575 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[0]))
	F_LWLockRelease(m, v575+int32(1152))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L11
	} else {
		goto L127
	}
L127:
	;
	F_CheckRequiredParameterValues(m)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L11
	} else {
		goto L128
	}
L128:
	;
	goto L5
L129:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585))))
	v588 = v586 & int32(1)
	if v588 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v592 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[5]))
	v595 = base.AtomicRmwXchg32(m, v592, int32(440), int32(1))
	if v595 != 0 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	goto L132
L132:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_xlog_redo[13])) = uint8(v588)
	goto L5
L133:
	;
	v597 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[5]))
	F_s_lock(m, v597+int32(440), int32(_a_F_xlog_redo_1), int32(_a_F_xlog_redo_11), int32(_a_F_xlog_redo_3))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L11
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v605 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v607 = *(*int32)(unsafe.Add(mBase, _c_F_xlog_redo[5]))
	v608 = *(*int64)(unsafe.Add(mBase, uint32(v607)+432))
	if base.Ui64(v608) < base.Ui64(v605) {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	goto L135
L137:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v607)+432)) = v605
	goto L139
L138:
	;
	goto L139
L139:
	;
	v611 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v607)+440)), uint32(v611))
	goto L132
L140:
	;
	F_errmsg(m, int32(_a_F_xlog_redo_12), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L11
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_xlog_redo_1), int32(_a_F_xlog_redo_13), int32(_a_F_xlog_redo_3))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v208
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v651
	F_errmsg(m, int32(_a_F_xlog_redo_14), v15)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L11
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_xlog_redo_1), int32(_a_F_xlog_redo_15), int32(_a_F_xlog_redo_3))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v332
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v666
	F_errmsg(m, int32(_a_F_xlog_redo_16), v15+int32(16))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L11
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_xlog_redo_1), int32(_a_F_xlog_redo_17), int32(_a_F_xlog_redo_3))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
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
	v685 = m.ExcPending
	if v685 != 0 {
		goto L11
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_xlog_redo_1), int32(_a_F_xlog_redo_19), int32(_a_F_xlog_redo_3))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
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
