package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateExplainSerializeDestReceiver(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc0(m, int32(208))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(12)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(549)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(550)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = int32(551)
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(552)
		return v4
	}
}
func F_ExplainIndentText(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if v5 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+v5-int32(1)))))
		if v10 != int32(10) {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			F_appendStringInfoSpaces(m, v4, v13<<(uint(int32(1))%32))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		F_appendStringInfoSpaces(m, v4, v13<<(uint(int32(1))%32))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			return
		}
	}
}
func F_ExplainIndexScanDetails(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainIndexScanDetails[0]))
	if v11 != 0 {
		v12 = m.T0[v11].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			if v12 != 0 {
				v19 = v12
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
				if v20 == int32(0) {
					if l1 == int32(-1) {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						F_appendStringInfoString(m, v25, int32(_a_F_ExplainIndexScanDetails_0))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v30 = F_quote_identifier(m, v19)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v30
								F_appendStringInfo(m, v29, int32(_a_F_ExplainIndexScanDetails_1), v8+int32(16))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									m.G0 = v8 + int32(32)
									return
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v30 = F_quote_identifier(m, v19)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v30
							F_appendStringInfo(m, v29, int32(_a_F_ExplainIndexScanDetails_1), v8+int32(16))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								m.G0 = v8 + int32(32)
								return
							}
						}
					}
				} else {
					if l1 == int32(1) {
						v44 = int32(_a_F_ExplainIndexScanDetails_2)
					} else {
						v44 = int32(_a_F_ExplainIndexScanDetails_3)
					}
					if l1 == int32(-1) {
						v47 = int32(_a_F_ExplainIndexScanDetails_4)
					} else {
						v47 = v44
					}
					F_ExplainPropertyText(m, int32(_a_F_ExplainIndexScanDetails_5), v47, l2)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						F_ExplainPropertyText(m, int32(_a_F_ExplainIndexScanDetails_6), v19, l2)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					}
				}
			} else {
				v15 = F_get_rel_name(m, l0)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					if v15 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
							F_errmsg_internal(m, int32(_a_F_ExplainIndexScanDetails_7), v8)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ExplainIndexScanDetails_8), int32(4035), int32(_a_F_ExplainIndexScanDetails_9))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v19 = v15
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
						if v20 == int32(0) {
							if l1 == int32(-1) {
								v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								F_appendStringInfoString(m, v25, int32(_a_F_ExplainIndexScanDetails_0))
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
									return
								} else {
									v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									v30 = F_quote_identifier(m, v19)
									mBase = m.M
									v31 = m.ExcPending
									if v31 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v30
										F_appendStringInfo(m, v29, int32(_a_F_ExplainIndexScanDetails_1), v8+int32(16))
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return
										} else {
											m.G0 = v8 + int32(32)
											return
										}
									}
								}
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v30 = F_quote_identifier(m, v19)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v30
									F_appendStringInfo(m, v29, int32(_a_F_ExplainIndexScanDetails_1), v8+int32(16))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										m.G0 = v8 + int32(32)
										return
									}
								}
							}
						} else {
							if l1 == int32(1) {
								v44 = int32(_a_F_ExplainIndexScanDetails_2)
							} else {
								v44 = int32(_a_F_ExplainIndexScanDetails_3)
							}
							if l1 == int32(-1) {
								v47 = int32(_a_F_ExplainIndexScanDetails_4)
							} else {
								v47 = v44
							}
							F_ExplainPropertyText(m, int32(_a_F_ExplainIndexScanDetails_5), v47, l2)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								F_ExplainPropertyText(m, int32(_a_F_ExplainIndexScanDetails_6), v19, l2)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									m.G0 = v8 + int32(32)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v15 = F_get_rel_name(m, l0)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if v15 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(_a_F_ExplainIndexScanDetails_7), v8)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ExplainIndexScanDetails_8), int32(4035), int32(_a_F_ExplainIndexScanDetails_9))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v19 = v15
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
				if v20 == int32(0) {
					if l1 == int32(-1) {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						F_appendStringInfoString(m, v25, int32(_a_F_ExplainIndexScanDetails_0))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v30 = F_quote_identifier(m, v19)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v30
								F_appendStringInfo(m, v29, int32(_a_F_ExplainIndexScanDetails_1), v8+int32(16))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									m.G0 = v8 + int32(32)
									return
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v30 = F_quote_identifier(m, v19)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v30
							F_appendStringInfo(m, v29, int32(_a_F_ExplainIndexScanDetails_1), v8+int32(16))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								m.G0 = v8 + int32(32)
								return
							}
						}
					}
				} else {
					if l1 == int32(1) {
						v44 = int32(_a_F_ExplainIndexScanDetails_2)
					} else {
						v44 = int32(_a_F_ExplainIndexScanDetails_3)
					}
					if l1 == int32(-1) {
						v47 = int32(_a_F_ExplainIndexScanDetails_4)
					} else {
						v47 = v44
					}
					F_ExplainPropertyText(m, int32(_a_F_ExplainIndexScanDetails_5), v47, l2)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						F_ExplainPropertyText(m, int32(_a_F_ExplainIndexScanDetails_6), v19, l2)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					}
				}
			}
		}
	}
}
func F_ExplainOnePlan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
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
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v128 float64
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int64
	_ = v159
	var v163 int64
	_ = v163
	var v167 int64
	_ = v167
	var v170 int64
	_ = v170
	var v173 int32
	_ = v173
	var v174 int64
	_ = v174
	var v177 int64
	_ = v177
	var v180 int64
	_ = v180
	var v183 int64
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int64
	_ = v189
	var v192 int64
	_ = v192
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v199 int64
	_ = v199
	var v202 int32
	_ = v202
	var v204 int64
	_ = v204
	var v207 int64
	_ = v207
	var v210 int32
	_ = v210
	var v211 int64
	_ = v211
	var v214 int64
	_ = v214
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int64
	_ = v267
	var v268 int32
	_ = v268
	var v272 int64
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v322 int32
	_ = v322
	var v329 int64
	_ = v329
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int64
	_ = v346
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int64
	_ = v373
	var v374 int64
	_ = v374
	var v377 int64
	_ = v377
	var v378 int64
	_ = v378
	var v381 int64
	_ = v381
	var v382 int64
	_ = v382
	var v385 int64
	_ = v385
	var v386 int64
	_ = v386
	var v389 int64
	_ = v389
	var v390 int64
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int64
	_ = v402
	var v403 int64
	_ = v403
	var v406 int64
	_ = v406
	var v407 int64
	_ = v407
	var v410 int64
	_ = v410
	var v411 int64
	_ = v411
	var v414 int64
	_ = v414
	var v415 int64
	_ = v415
	var v418 int64
	_ = v418
	var v419 int64
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int64
	_ = v452
	var v458 int64
	_ = v458
	var v467 int32
	_ = v467
	var v469 int64
	_ = v469
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int64
	_ = v486
	var v490 int64
	_ = v490
	var v494 int64
	_ = v494
	var v497 int64
	_ = v497
	var v500 int32
	_ = v500
	var v501 int64
	_ = v501
	var v504 int64
	_ = v504
	var v507 int64
	_ = v507
	var v510 int64
	_ = v510
	var v513 int32
	_ = v513
	var v517 int64
	_ = v517
	var v521 int64
	_ = v521
	var v524 int64
	_ = v524
	var v527 int64
	_ = v527
	var v530 int64
	_ = v530
	var v533 int64
	_ = v533
	var v536 int64
	_ = v536
	var v539 int64
	_ = v539
	var v543 int32
	_ = v543
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v560 int64
	_ = v560
	var v568 int32
	_ = v568
	var v571 int64
	_ = v571
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v605 int64
	_ = v605
	var v606 int64
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v630 int64
	_ = v630
	var v632 int64
	_ = v632
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	v10 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(256)
	m.G0 = v22
	base.MemoryFill(m, v22-int32(-64), v10, int32(144))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v29 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v34 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v36 = v10
	goto L3
L3:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	F___clock_gettime(m, int32(1), v22+int32(208))
	mBase = m.M
	v43 = int64(*(*int32)(unsafe.Add(mBase, uint32(v22)+216)))
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v22)+208))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOnePlan[0]))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	goto L7
L4:
	;
	v35 = int32(1)
	goto L6
L5:
	;
	v35 = int32(4)
	goto L6
L6:
	;
	v36 = v35
	goto L3
L7:
	;
	F_PushCopiedSnapshot(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	F_UpdateActiveSnapshotCommandId(m)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if l1 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOnePlan[0]))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	goto L20
L12:
	;
	v62 = F_CreateIntoRelDestReceiver(m, l1)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v64 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v69 = v62
	goto L11
L16:
	;
	v65 = F_CreateExplainSerializeDestReceiver(m, l2)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L8
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOnePlan[2]))
	v69 = v68
	goto L11
L19:
	;
	v69 = v65
	goto L11
L20:
	;
	v74 = F_CreateQueryDesc(m, l0, l3, v72, int32(0), v69, l4, l5, v38<<(uint(int32(3))%32)&int32(8)|(v36|v37<<(uint(int32(1))%32)&int32(2)))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+13)))
	v78 = int32(1)
	v84 = v76 | v77<<(uint(v78)%32)&int32(2) ^ v78
	if l1 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v129 != 0 {
		goto L36
	} else {
		goto L37
	}
L23:
	;
	F_ExecutorRun(m, v74, v106, int64(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L34
	}
L24:
	;
	v106 = int32(1)
	goto L23
L25:
	;
	F_ExecutorStart(m, v74, v84)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	goto L30
L28:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v89 != 0 {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v128 = float64(0)
	goto L22
L30:
	;
	F_ExecutorStart(m, v74, v91<<(uint(int32(6))%32)&int32(64)|v84)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v100 != int32(1) {
		v128 = float64(0)
		goto L22
	} else {
		goto L32
	}
L32:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v104 != 0 {
		v106 = int32(0)
		goto L23
	} else {
		goto L33
	}
L33:
	;
	goto L24
L34:
	;
	F_ExecutorFinish(m, v74)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	F___clock_gettime(m, int32(1), v22+int32(208))
	mBase = m.M
	v116 = int64(*(*int32)(unsafe.Add(mBase, uint32(v22)+216)))
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v22)+208))
	v128 = base.F64_add(base.F64_div(base.F64_convert_i64_s(v116-v43+(v118-v44)*int64(1000000000)), float64(1e+09)), float64(0))
	goto L22
L36:
	;
	v131 = v22 - int32(-64)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	if v132 == int32(12) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	goto L38
L38:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	m.T0[v143].(func(*base.Module, int32))(m, v69)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L8
	} else {
		goto L43
	}
L39:
	;
	goto L38
L40:
	;
	base.MemoryCopy(m, v131, v69-int32(-64), int32(144))
	goto L39
L41:
	;
	goto L42
L42:
	;
	base.MemoryFill(m, v131, int32(0), int32(144))
	goto L39
L43:
	;
	F_ExplainOpenGroup(m, int32(_a_F_ExplainOnePlan_0), int32(0), int32(1), l2)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	F_ExplainPrintPlan(m, l2, v74)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	if l7 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	if l6 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L47:
	;
	v240 = int32(_a_F_ExplainOnePlan_1)
	F_ExplainOpenGroup(m, v240, v240, int32(1), l2)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L8
	} else {
		goto L76
	}
L48:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v154 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v225 = int32(0)
	goto L50
L50:
	;
	v230 = base.B2i32(l8 != int32(0))
	if l8 != 0 {
		v235 = v230
		goto L47
	} else {
		goto L74
	}
L51:
	;
	v235 = base.B2i32(l8 != int32(0))
	goto L47
L52:
	;
	goto L53
L53:
	;
	v157 = int32(1)
	v159 = *(*int64)(unsafe.Add(mBase, uint32(l7)))
	if int64(0) < v159 {
		v173 = v157
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v174 = *(*int64)(unsafe.Add(mBase, uint32(l7)+32))
	if int64(0) < v174 {
		v186 = v157
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v163 = *(*int64)(unsafe.Add(mBase, uint32(l7)+8))
	if int64(0) < v163 {
		v173 = int32(1)
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v167 = *(*int64)(unsafe.Add(mBase, uint32(l7)+16))
	if int64(0) < v167 {
		v173 = int32(1)
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v170 = *(*int64)(unsafe.Add(mBase, uint32(l7)+24))
	v173 = base.B2i32(int64(0) < v170)
	goto L54
L58:
	;
	v187 = int32(1)
	v189 = *(*int64)(unsafe.Add(mBase, uint32(l7)+64))
	if v189 <= int64(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v177 = *(*int64)(unsafe.Add(mBase, uint32(l7)+40))
	if int64(0) < v177 {
		v186 = v157
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v180 = *(*int64)(unsafe.Add(mBase, uint32(l7)+48))
	if int64(0) < v180 {
		v186 = v157
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v183 = *(*int64)(unsafe.Add(mBase, uint32(l7)+56))
	v186 = base.B2i32(int64(0) < v183)
	goto L58
L62:
	;
	v192 = *(*int64)(unsafe.Add(mBase, uint32(l7)+72))
	v195 = base.B2i32(int64(0) < v192)
	goto L64
L63:
	;
	v195 = v187
	goto L64
L64:
	;
	v196 = *(*int64)(unsafe.Add(mBase, uint32(l7)+80))
	if v196 == int64(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v199 = *(*int64)(unsafe.Add(mBase, uint32(l7)+88))
	v202 = base.B2i32(v199 != int64(0))
	goto L67
L66:
	;
	v202 = v187
	goto L67
L67:
	;
	v204 = *(*int64)(unsafe.Add(mBase, uint32(l7)+96))
	if v204 == int64(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v207 = *(*int64)(unsafe.Add(mBase, uint32(l7)+104))
	v210 = base.B2i32(v207 != int64(0))
	goto L70
L69:
	;
	v210 = int32(1)
	goto L70
L70:
	;
	v211 = *(*int64)(unsafe.Add(mBase, uint32(l7)+112))
	if v211 == int64(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v214 = *(*int64)(unsafe.Add(mBase, uint32(l7)+120))
	v218 = base.B2i32(v214 != int64(0))
	goto L73
L72:
	;
	v218 = int32(1)
	goto L73
L73:
	;
	v225 = v218 | (v186 | v173 | v195 | v202 | v210)
	goto L50
L74:
	;
	if v225&int32(1) == int32(0) {
		goto L46
	} else {
		goto L75
	}
L75:
	;
	v235 = v230
	goto L47
L76:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v245 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_ExplainIndentText(m, l2)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L8
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if l7 != 0 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v250, int32(_a_F_ExplainOnePlan_2))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L8
	} else {
		goto L81
	}
L81:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v254 + int32(1)
	goto L79
L82:
	;
	F_show_buffer_usage(m, l2, l7)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v235 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	goto L84
L86:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v301 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L87:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l8)+8))
	v264 = v262 + int32(1023)
	v265 = int32(10)
	v267 = base.I64_extend_i32_u(int32(base.Ui32(v264) >> (uint(v265) % 32)))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l8)+12))
	v272 = base.I64_extend_i32_u(int32(base.Ui32(v264-v268) >> (uint(v265) % 32)))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v273 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	F_ExplainIndentText(m, l2)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L8
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	F_ExplainPropertyInteger(m, int32(_a_F_ExplainOnePlan_10), int32(_a_F_ExplainOnePlan_11), v272, l2)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L8
	} else {
		goto L94
	}
L91:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v267
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v272
	F_appendStringInfo(m, v278, int32(_a_F_ExplainOnePlan_12), v22+int32(48))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L8
	} else {
		goto L92
	}
L92:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v286, int32(10))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	goto L86
L94:
	;
	F_ExplainPropertyInteger(m, int32(_a_F_ExplainOnePlan_13), int32(_a_F_ExplainOnePlan_11), v267, l2)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L8
	} else {
		goto L95
	}
L95:
	;
	goto L86
L96:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v304 - int32(1)
	goto L98
L97:
	;
	goto L98
L98:
	;
	F_ExplainCloseGroup(m, int32(_a_F_ExplainOnePlan_1), int32(1), l2)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	goto L46
L100:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v338 == int32(1) {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
	if v322&int32(1) == int32(0) {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v329 = *(*int64)(unsafe.Add(mBase, uint32(l6)))
	F_ExplainPropertyFloat(m, int32(_a_F_ExplainOnePlan_14), int32(_a_F_ExplainOnePlan_8), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v329), float64(1e+09)), float64(1000)), int32(3), l2)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	goto L100
L104:
	;
	F_ExplainPrintTriggers(m, l2, v74)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L8
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v343 != int32(1) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	goto L106
L108:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v431 != 0 {
		goto L120
	} else {
		goto L121
	}
L109:
	;
	v346 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+248)) = v346
	*(*int64)(unsafe.Add(mBase, uint32(v22)+240)) = v346
	*(*int64)(unsafe.Add(mBase, uint32(v22)+232)) = v346
	*(*int64)(unsafe.Add(mBase, uint32(v22)+224)) = v346
	*(*int64)(unsafe.Add(mBase, uint32(v22)+216)) = v346
	*(*int64)(unsafe.Add(mBase, uint32(v22)+208)) = v346
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v74)+40))
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+176)))
	if v359&int32(1) == int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v358)+180))
	if v364 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v366 = v22 + int32(208)
	v368 = v364 + int32(8)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v366)))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v368)))
	*(*int32)(unsafe.Add(mBase, uint32(v366))) = v369 + v370
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v366)+8))
	v374 = *(*int64)(unsafe.Add(mBase, uint32(v368)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v366)+8)) = v373 + v374
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v366)+16))
	v378 = *(*int64)(unsafe.Add(mBase, uint32(v368)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v366)+16)) = v377 + v378
	v381 = *(*int64)(unsafe.Add(mBase, uint32(v366)+24))
	v382 = *(*int64)(unsafe.Add(mBase, uint32(v368)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v366)+24)) = v381 + v382
	v385 = *(*int64)(unsafe.Add(mBase, uint32(v366)+32))
	v386 = *(*int64)(unsafe.Add(mBase, uint32(v368)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v366)+32)) = v385 + v386
	v389 = *(*int64)(unsafe.Add(mBase, uint32(v366)+40))
	v390 = *(*int64)(unsafe.Add(mBase, uint32(v368)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v366)+40)) = v389 + v390
	goto L114
L112:
	;
	v394 = v358
	goto L113
L113:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+184))
	if v395 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v74)+40))
	v394 = v393
	goto L113
L115:
	;
	v397 = v22 + int32(208)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	*(*int32)(unsafe.Add(mBase, uint32(v397))) = v398 + v399
	v402 = *(*int64)(unsafe.Add(mBase, uint32(v397)+8))
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v395)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v397)+8)) = v402 + v403
	v406 = *(*int64)(unsafe.Add(mBase, uint32(v397)+16))
	v407 = *(*int64)(unsafe.Add(mBase, uint32(v395)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v397)+16)) = v406 + v407
	v410 = *(*int64)(unsafe.Add(mBase, uint32(v397)+24))
	v411 = *(*int64)(unsafe.Add(mBase, uint32(v395)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v397)+24)) = v410 + v411
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v397)+32))
	v415 = *(*int64)(unsafe.Add(mBase, uint32(v395)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v397)+32)) = v414 + v415
	v418 = *(*int64)(unsafe.Add(mBase, uint32(v397)+40))
	v419 = *(*int64)(unsafe.Add(mBase, uint32(v395)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v397)+40)) = v418 + v419
	goto L118
L116:
	;
	v423 = v394
	goto L117
L117:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)+176))
	F_ExplainPrintJIT(m, l2, v424, v22+int32(208))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L8
	} else {
		goto L119
	}
L118:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v74)+40))
	v423 = v422
	goto L117
L119:
	;
	goto L108
L120:
	;
	v432 = int32(_a_F_ExplainOnePlan_3)
	F_ExplainOpenGroup(m, v432, v432, int32(1), l2)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L8
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v598 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOnePlan[1]))
	if v598 != 0 {
		goto L167
	} else {
		goto L168
	}
L123:
	;
	if v431 == int32(1) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v441 = int32(_a_F_ExplainOnePlan_4)
	goto L126
L125:
	;
	v441 = int32(_a_F_ExplainOnePlan_5)
	goto L126
L126:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v442 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	F_ExplainCloseGroup(m, int32(_a_F_ExplainOnePlan_3), int32(1), l2)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L8
	} else {
		goto L166
	}
L128:
	;
	F_ExplainIndentText(m, l2)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L8
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v555 == int32(1) {
		goto L158
	} else {
		goto L159
	}
L131:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v448 == int32(1) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v480 != int32(1) {
		goto L127
	} else {
		goto L138
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v441
	v452 = *(*int64)(unsafe.Add(mBase, uint32(v22)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = int64(base.Ui64(v452+int64(1023)) >> (uint(int64(10)) % 64))
	v458 = *(*int64)(unsafe.Add(mBase, uint32(v22)+72))
	*(*float64)(unsafe.Add(mBase, uint32(v22))) = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v458), float64(1e+09)), float64(1000))
	F_appendStringInfo(m, v447, int32(_a_F_ExplainOnePlan_6), v22)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L8
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v441
	v469 = *(*int64)(unsafe.Add(mBase, uint32(v22)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = int64(base.Ui64(v469+int64(1023)) >> (uint(int64(10)) % 64))
	F_appendStringInfo(m, v447, int32(_a_F_ExplainOnePlan_9), v22+int32(32))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L8
	} else {
		goto L137
	}
L136:
	;
	goto L132
L137:
	;
	goto L132
L138:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v483 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v543 + int32(1)
	F_show_buffer_usage(m, l2, v22+int32(80))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L8
	} else {
		goto L157
	}
L140:
	;
	v484 = int32(1)
	v486 = *(*int64)(unsafe.Add(mBase, uint32(v22)+80))
	if int64(0) < v486 {
		v500 = v484
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v501 = *(*int64)(unsafe.Add(mBase, uint32(v22)+112))
	if int64(0) < v501 {
		v513 = v484
		goto L145
	} else {
		goto L146
	}
L142:
	;
	v490 = *(*int64)(unsafe.Add(mBase, uint32(v22)+88))
	if int64(0) < v490 {
		v500 = int32(1)
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v494 = *(*int64)(unsafe.Add(mBase, uint32(v22)+96))
	if int64(0) < v494 {
		v500 = int32(1)
		goto L141
	} else {
		goto L144
	}
L144:
	;
	v497 = *(*int64)(unsafe.Add(mBase, uint32(v22)+104))
	v500 = base.B2i32(int64(0) < v497)
	goto L141
L145:
	;
	v517 = *(*int64)(unsafe.Add(mBase, uint32(v22)+192))
	if (v500|v513)&int32(1)|base.B2i32(v517 != int64(0)) != 0 {
		goto L139
	} else {
		goto L149
	}
L146:
	;
	v504 = *(*int64)(unsafe.Add(mBase, uint32(v22)+120))
	if int64(0) < v504 {
		v513 = v484
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v507 = *(*int64)(unsafe.Add(mBase, uint32(v22)+128))
	if int64(0) < v507 {
		v513 = v484
		goto L145
	} else {
		goto L148
	}
L148:
	;
	v510 = *(*int64)(unsafe.Add(mBase, uint32(v22)+136))
	v513 = base.B2i32(int64(0) < v510)
	goto L145
L149:
	;
	v521 = *(*int64)(unsafe.Add(mBase, uint32(v22)+144))
	if int64(0) < v521 {
		goto L139
	} else {
		goto L150
	}
L150:
	;
	v524 = *(*int64)(unsafe.Add(mBase, uint32(v22)+152))
	if int64(0) < v524 {
		goto L139
	} else {
		goto L151
	}
L151:
	;
	v527 = *(*int64)(unsafe.Add(mBase, uint32(v22)+160))
	if v527 != int64(0) {
		goto L139
	} else {
		goto L152
	}
L152:
	;
	v530 = *(*int64)(unsafe.Add(mBase, uint32(v22)+168))
	if v530 != int64(0) {
		goto L139
	} else {
		goto L153
	}
L153:
	;
	v533 = *(*int64)(unsafe.Add(mBase, uint32(v22)+176))
	if v533 != int64(0) {
		goto L139
	} else {
		goto L154
	}
L154:
	;
	v536 = *(*int64)(unsafe.Add(mBase, uint32(v22)+184))
	if v536 != int64(0) {
		goto L139
	} else {
		goto L155
	}
L155:
	;
	v539 = *(*int64)(unsafe.Add(mBase, uint32(v22)+200))
	if v539 == int64(0) {
		goto L127
	} else {
		goto L156
	}
L156:
	;
	goto L139
L157:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v551 - int32(1)
	goto L127
L158:
	;
	v560 = *(*int64)(unsafe.Add(mBase, uint32(v22)+72))
	F_ExplainPropertyFloat(m, int32(_a_F_ExplainOnePlan_15), int32(_a_F_ExplainOnePlan_8), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v560), float64(1e+09)), float64(1000)), int32(3), l2)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L8
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v571 = *(*int64)(unsafe.Add(mBase, uint32(v22)+64))
	F_ExplainPropertyUInteger(m, int32(_a_F_ExplainOnePlan_16), int32(_a_F_ExplainOnePlan_11), int64(base.Ui64(v571+int64(1023))>>(uint(int64(10))%64)), l2)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L8
	} else {
		goto L162
	}
L161:
	;
	goto L160
L162:
	;
	F_ExplainPropertyText(m, int32(_a_F_ExplainOnePlan_17), v441, l2)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L8
	} else {
		goto L163
	}
L163:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v581 != int32(1) {
		goto L127
	} else {
		goto L164
	}
L164:
	;
	F_show_buffer_usage(m, l2, v22+int32(80))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L8
	} else {
		goto L165
	}
L165:
	;
	goto L127
L166:
	;
	goto L122
L167:
	;
	m.T0[v598].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, l4, l5)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L8
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	F___clock_gettime(m, int32(1), v22+int32(208))
	mBase = m.M
	v605 = int64(*(*int32)(unsafe.Add(mBase, uint32(v22)+216)))
	v606 = *(*int64)(unsafe.Add(mBase, uint32(v22)+208))
	F_ExecutorEnd(m, v74)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L8
	} else {
		goto L171
	}
L170:
	;
	goto L169
L171:
	;
	F_FreeQueryDesc(m, v74)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L8
	} else {
		goto L172
	}
L172:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L8
	} else {
		goto L173
	}
L173:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v613 == int32(1) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L8
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v618 = int32(1)
	F___clock_gettime(m, v618, v22+int32(208))
	mBase = m.M
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
	if v622 != v618 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	goto L176
L178:
	;
	F_ExplainCloseGroup(m, int32(_a_F_ExplainOnePlan_0), int32(1), l2)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L8
	} else {
		goto L182
	}
L179:
	;
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v625 != int32(1) {
		goto L178
	} else {
		goto L180
	}
L180:
	;
	v630 = int64(*(*int32)(unsafe.Add(mBase, uint32(v22)+216)))
	v632 = *(*int64)(unsafe.Add(mBase, uint32(v22)+208))
	F_ExplainPropertyFloat(m, int32(_a_F_ExplainOnePlan_7), int32(_a_F_ExplainOnePlan_8), base.F64_mul(base.F64_add(v128, base.F64_div(base.F64_convert_i64_s(v630-v605+(v632-v606)*int64(1000000000)), float64(1e+09))), float64(1000)), int32(3), l2)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L8
	} else {
		goto L181
	}
L181:
	;
	goto L178
L182:
	;
	m.G0 = v22 + int32(256)
	return
}
func F_NewExplainState(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = F_palloc0(m, int32(72))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+6)) = uint8(v7)
		v9 = F_makeStringInfo(m)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v3))) = v9
			return v3
		}
	}
}
