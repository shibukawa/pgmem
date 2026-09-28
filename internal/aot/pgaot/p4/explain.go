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
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(590)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(591)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = int32(592)
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(593)
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
								F_errfinish(m, int32(_a_F_ExplainIndexScanDetails_8), int32(_a_F_ExplainIndexScanDetails_9), int32(_a_F_ExplainIndexScanDetails_10))
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
						F_errfinish(m, int32(_a_F_ExplainIndexScanDetails_8), int32(_a_F_ExplainIndexScanDetails_9), int32(_a_F_ExplainIndexScanDetails_10))
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
	var v39 int32
	_ = v39
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v134 float64
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int64
	_ = v165
	var v169 int64
	_ = v169
	var v173 int64
	_ = v173
	var v176 int64
	_ = v176
	var v179 int32
	_ = v179
	var v180 int64
	_ = v180
	var v183 int64
	_ = v183
	var v186 int64
	_ = v186
	var v189 int64
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int64
	_ = v195
	var v198 int64
	_ = v198
	var v201 int32
	_ = v201
	var v202 int64
	_ = v202
	var v205 int64
	_ = v205
	var v208 int32
	_ = v208
	var v210 int64
	_ = v210
	var v213 int64
	_ = v213
	var v216 int32
	_ = v216
	var v217 int64
	_ = v217
	var v220 int64
	_ = v220
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int64
	_ = v273
	var v274 int32
	_ = v274
	var v278 int64
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v328 int32
	_ = v328
	var v335 int64
	_ = v335
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int64
	_ = v352
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int64
	_ = v379
	var v380 int64
	_ = v380
	var v383 int64
	_ = v383
	var v384 int64
	_ = v384
	var v387 int64
	_ = v387
	var v388 int64
	_ = v388
	var v391 int64
	_ = v391
	var v392 int64
	_ = v392
	var v395 int64
	_ = v395
	var v396 int64
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int64
	_ = v408
	var v409 int64
	_ = v409
	var v412 int64
	_ = v412
	var v413 int64
	_ = v413
	var v416 int64
	_ = v416
	var v417 int64
	_ = v417
	var v420 int64
	_ = v420
	var v421 int64
	_ = v421
	var v424 int64
	_ = v424
	var v425 int64
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int64
	_ = v458
	var v464 int64
	_ = v464
	var v473 int32
	_ = v473
	var v475 int64
	_ = v475
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int64
	_ = v492
	var v496 int64
	_ = v496
	var v500 int64
	_ = v500
	var v503 int64
	_ = v503
	var v506 int32
	_ = v506
	var v507 int64
	_ = v507
	var v510 int64
	_ = v510
	var v513 int64
	_ = v513
	var v516 int64
	_ = v516
	var v519 int32
	_ = v519
	var v523 int64
	_ = v523
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
	var v542 int64
	_ = v542
	var v545 int64
	_ = v545
	var v549 int32
	_ = v549
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v566 int64
	_ = v566
	var v574 int32
	_ = v574
	var v577 int64
	_ = v577
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v611 int64
	_ = v611
	var v612 int64
	_ = v612
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v636 int64
	_ = v636
	var v638 int64
	_ = v638
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
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
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+13)))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	F___clock_gettime(m, int32(1), v22+int32(208))
	mBase = m.M
	v44 = int64(*(*int32)(unsafe.Add(mBase, uint32(v22)+216)))
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v22)+208))
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOnePlan[0]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
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
	F_PushCopiedSnapshot(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
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
	v67 = m.ExcPending
	if v67 != 0 {
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
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOnePlan[0]))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	goto L20
L12:
	;
	v68 = F_CreateIntoRelDestReceiver(m, l1)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v70 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v75 = v68
	goto L11
L16:
	;
	v71 = F_CreateExplainSerializeDestReceiver(m, l2)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOnePlan[2]))
	v75 = v74
	goto L11
L19:
	;
	v75 = v71
	goto L11
L20:
	;
	v80 = F_CreateQueryDesc(m, l0, l3, v78, int32(0), v75, l4, l5, v37<<(uint(int32(4))%32)&int32(16)|(v39<<(uint(int32(3))%32)&int32(8)|(v36|v38<<(uint(int32(1))%32)&int32(2))))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+14)))
	v84 = int32(1)
	v90 = v82 | v83<<(uint(v84)%32)&int32(2) ^ v84
	if l1 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v135 != 0 {
		goto L36
	} else {
		goto L37
	}
L23:
	;
	F_ExecutorRun(m, v80, v112, int64(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L8
	} else {
		goto L34
	}
L24:
	;
	v112 = int32(1)
	goto L23
L25:
	;
	F_ExecutorStart(m, v80, v90)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L8
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	goto L30
L28:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v95 != 0 {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v134 = float64(0)
	goto L22
L30:
	;
	F_ExecutorStart(m, v80, v97<<(uint(int32(6))%32)&int32(64)|v90)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v106 != int32(1) {
		v134 = float64(0)
		goto L22
	} else {
		goto L32
	}
L32:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v110 != 0 {
		v112 = int32(0)
		goto L23
	} else {
		goto L33
	}
L33:
	;
	goto L24
L34:
	;
	F_ExecutorFinish(m, v80)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	F___clock_gettime(m, int32(1), v22+int32(208))
	mBase = m.M
	v122 = int64(*(*int32)(unsafe.Add(mBase, uint32(v22)+216)))
	v124 = *(*int64)(unsafe.Add(mBase, uint32(v22)+208))
	v134 = base.F64_add(base.F64_div(base.F64_convert_i64_s(v122-v44+(v124-v45)*int64(1000000000)), float64(1e+09)), float64(0))
	goto L22
L36:
	;
	v137 = v22 - int32(-64)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	if v138 == int32(12) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	goto L38
L38:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	m.T0[v149].(func(*base.Module, int32))(m, v75)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L8
	} else {
		goto L43
	}
L39:
	;
	goto L38
L40:
	;
	base.MemoryCopy(m, v137, v75-int32(-64), int32(144))
	goto L39
L41:
	;
	goto L42
L42:
	;
	base.MemoryFill(m, v137, int32(0), int32(144))
	goto L39
L43:
	;
	F_ExplainOpenGroup(m, int32(_a_F_ExplainOnePlan_0), int32(0), int32(1), l2)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	F_ExplainPrintPlan(m, l2, v80)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
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
	v246 = int32(_a_F_ExplainOnePlan_1)
	F_ExplainOpenGroup(m, v246, v246, int32(1), l2)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L8
	} else {
		goto L76
	}
L48:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v160 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v231 = int32(0)
	goto L50
L50:
	;
	v236 = base.B2i32(l8 != int32(0))
	if l8 != 0 {
		v241 = v236
		goto L47
	} else {
		goto L74
	}
L51:
	;
	v241 = base.B2i32(l8 != int32(0))
	goto L47
L52:
	;
	goto L53
L53:
	;
	v163 = int32(1)
	v165 = *(*int64)(unsafe.Add(mBase, uint32(l7)))
	if int64(0) < v165 {
		v179 = v163
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v180 = *(*int64)(unsafe.Add(mBase, uint32(l7)+32))
	if int64(0) < v180 {
		v192 = v163
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v169 = *(*int64)(unsafe.Add(mBase, uint32(l7)+8))
	if int64(0) < v169 {
		v179 = int32(1)
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v173 = *(*int64)(unsafe.Add(mBase, uint32(l7)+16))
	if int64(0) < v173 {
		v179 = int32(1)
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v176 = *(*int64)(unsafe.Add(mBase, uint32(l7)+24))
	v179 = base.B2i32(int64(0) < v176)
	goto L54
L58:
	;
	v193 = int32(1)
	v195 = *(*int64)(unsafe.Add(mBase, uint32(l7)+64))
	if v195 <= int64(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v183 = *(*int64)(unsafe.Add(mBase, uint32(l7)+40))
	if int64(0) < v183 {
		v192 = v163
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v186 = *(*int64)(unsafe.Add(mBase, uint32(l7)+48))
	if int64(0) < v186 {
		v192 = v163
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v189 = *(*int64)(unsafe.Add(mBase, uint32(l7)+56))
	v192 = base.B2i32(int64(0) < v189)
	goto L58
L62:
	;
	v198 = *(*int64)(unsafe.Add(mBase, uint32(l7)+72))
	v201 = base.B2i32(int64(0) < v198)
	goto L64
L63:
	;
	v201 = v193
	goto L64
L64:
	;
	v202 = *(*int64)(unsafe.Add(mBase, uint32(l7)+80))
	if v202 == int64(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v205 = *(*int64)(unsafe.Add(mBase, uint32(l7)+88))
	v208 = base.B2i32(v205 != int64(0))
	goto L67
L66:
	;
	v208 = v193
	goto L67
L67:
	;
	v210 = *(*int64)(unsafe.Add(mBase, uint32(l7)+96))
	if v210 == int64(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v213 = *(*int64)(unsafe.Add(mBase, uint32(l7)+104))
	v216 = base.B2i32(v213 != int64(0))
	goto L70
L69:
	;
	v216 = int32(1)
	goto L70
L70:
	;
	v217 = *(*int64)(unsafe.Add(mBase, uint32(l7)+112))
	if v217 == int64(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v220 = *(*int64)(unsafe.Add(mBase, uint32(l7)+120))
	v224 = base.B2i32(v220 != int64(0))
	goto L73
L72:
	;
	v224 = int32(1)
	goto L73
L73:
	;
	v231 = v224 | (v192 | v179 | v201 | v208 | v216)
	goto L50
L74:
	;
	if v231&int32(1) == int32(0) {
		goto L46
	} else {
		goto L75
	}
L75:
	;
	v241 = v236
	goto L47
L76:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v251 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_ExplainIndentText(m, l2)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
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
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v256, int32(_a_F_ExplainOnePlan_2))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L81
	}
L81:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v260 + int32(1)
	goto L79
L82:
	;
	F_show_buffer_usage(m, l2, l7)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L8
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v241 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	goto L84
L86:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v307 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L87:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l8)+8))
	v270 = v268 + int32(1023)
	v271 = int32(10)
	v273 = base.I64_extend_i32_u(int32(base.Ui32(v270) >> (uint(v271) % 32)))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l8)+12))
	v278 = base.I64_extend_i32_u(int32(base.Ui32(v270-v274) >> (uint(v271) % 32)))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v279 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	F_ExplainIndentText(m, l2)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L8
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	F_ExplainPropertyInteger(m, int32(_a_F_ExplainOnePlan_10), int32(_a_F_ExplainOnePlan_11), v278, l2)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L8
	} else {
		goto L94
	}
L91:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v273
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v278
	F_appendStringInfo(m, v284, int32(_a_F_ExplainOnePlan_12), v22+int32(48))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L8
	} else {
		goto L92
	}
L92:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v292, int32(10))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	goto L86
L94:
	;
	F_ExplainPropertyInteger(m, int32(_a_F_ExplainOnePlan_13), int32(_a_F_ExplainOnePlan_11), v273, l2)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L8
	} else {
		goto L95
	}
L95:
	;
	goto L86
L96:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v310 - int32(1)
	goto L98
L97:
	;
	goto L98
L98:
	;
	F_ExplainCloseGroup(m, int32(_a_F_ExplainOnePlan_1), int32(1), l2)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	goto L46
L100:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v344 == int32(1) {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
	if v328&int32(1) == int32(0) {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v335 = *(*int64)(unsafe.Add(mBase, uint32(l6)))
	F_ExplainPropertyFloat(m, int32(_a_F_ExplainOnePlan_14), int32(_a_F_ExplainOnePlan_8), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v335), float64(1e+09)), float64(1000)), int32(3), l2)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	goto L100
L104:
	;
	F_ExplainPrintTriggers(m, l2, v80)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L8
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v349 != int32(1) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	goto L106
L108:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v437 != 0 {
		goto L120
	} else {
		goto L121
	}
L109:
	;
	v352 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+248)) = v352
	*(*int64)(unsafe.Add(mBase, uint32(v22)+240)) = v352
	*(*int64)(unsafe.Add(mBase, uint32(v22)+232)) = v352
	*(*int64)(unsafe.Add(mBase, uint32(v22)+224)) = v352
	*(*int64)(unsafe.Add(mBase, uint32(v22)+216)) = v352
	*(*int64)(unsafe.Add(mBase, uint32(v22)+208)) = v352
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v80)+44))
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364)+176)))
	if v365&int32(1) == int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v364)+180))
	if v370 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v372 = v22 + int32(208)
	v374 = v370 + int32(8)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	*(*int32)(unsafe.Add(mBase, uint32(v372))) = v375 + v376
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v372)+8))
	v380 = *(*int64)(unsafe.Add(mBase, uint32(v374)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v372)+8)) = v379 + v380
	v383 = *(*int64)(unsafe.Add(mBase, uint32(v372)+16))
	v384 = *(*int64)(unsafe.Add(mBase, uint32(v374)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v372)+16)) = v383 + v384
	v387 = *(*int64)(unsafe.Add(mBase, uint32(v372)+24))
	v388 = *(*int64)(unsafe.Add(mBase, uint32(v374)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v372)+24)) = v387 + v388
	v391 = *(*int64)(unsafe.Add(mBase, uint32(v372)+32))
	v392 = *(*int64)(unsafe.Add(mBase, uint32(v374)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v372)+32)) = v391 + v392
	v395 = *(*int64)(unsafe.Add(mBase, uint32(v372)+40))
	v396 = *(*int64)(unsafe.Add(mBase, uint32(v374)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v372)+40)) = v395 + v396
	goto L114
L112:
	;
	v400 = v364
	goto L113
L113:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+184))
	if v401 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v80)+44))
	v400 = v399
	goto L113
L115:
	;
	v403 = v22 + int32(208)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	*(*int32)(unsafe.Add(mBase, uint32(v403))) = v404 + v405
	v408 = *(*int64)(unsafe.Add(mBase, uint32(v403)+8))
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v401)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v403)+8)) = v408 + v409
	v412 = *(*int64)(unsafe.Add(mBase, uint32(v403)+16))
	v413 = *(*int64)(unsafe.Add(mBase, uint32(v401)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v403)+16)) = v412 + v413
	v416 = *(*int64)(unsafe.Add(mBase, uint32(v403)+24))
	v417 = *(*int64)(unsafe.Add(mBase, uint32(v401)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v403)+24)) = v416 + v417
	v420 = *(*int64)(unsafe.Add(mBase, uint32(v403)+32))
	v421 = *(*int64)(unsafe.Add(mBase, uint32(v401)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v403)+32)) = v420 + v421
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v403)+40))
	v425 = *(*int64)(unsafe.Add(mBase, uint32(v401)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v403)+40)) = v424 + v425
	goto L118
L116:
	;
	v429 = v400
	goto L117
L117:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)+176))
	F_ExplainPrintJIT(m, l2, v430, v22+int32(208))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L8
	} else {
		goto L119
	}
L118:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v80)+44))
	v429 = v428
	goto L117
L119:
	;
	goto L108
L120:
	;
	v438 = int32(_a_F_ExplainOnePlan_3)
	F_ExplainOpenGroup(m, v438, v438, int32(1), l2)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L8
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v604 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOnePlan[1]))
	if v604 != 0 {
		goto L167
	} else {
		goto L168
	}
L123:
	;
	if v437 == int32(1) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v447 = int32(_a_F_ExplainOnePlan_4)
	goto L126
L125:
	;
	v447 = int32(_a_F_ExplainOnePlan_5)
	goto L126
L126:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v448 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	F_ExplainCloseGroup(m, int32(_a_F_ExplainOnePlan_3), int32(1), l2)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L8
	} else {
		goto L166
	}
L128:
	;
	F_ExplainIndentText(m, l2)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L8
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v561 == int32(1) {
		goto L158
	} else {
		goto L159
	}
L131:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v454 == int32(1) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v486 != int32(1) {
		goto L127
	} else {
		goto L138
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v447
	v458 = *(*int64)(unsafe.Add(mBase, uint32(v22)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = int64(base.Ui64(v458+int64(1023)) >> (uint(int64(10)) % 64))
	v464 = *(*int64)(unsafe.Add(mBase, uint32(v22)+72))
	*(*float64)(unsafe.Add(mBase, uint32(v22))) = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v464), float64(1e+09)), float64(1000))
	F_appendStringInfo(m, v453, int32(_a_F_ExplainOnePlan_6), v22)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L8
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v447
	v475 = *(*int64)(unsafe.Add(mBase, uint32(v22)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = int64(base.Ui64(v475+int64(1023)) >> (uint(int64(10)) % 64))
	F_appendStringInfo(m, v453, int32(_a_F_ExplainOnePlan_9), v22+int32(32))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
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
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v489 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v549 + int32(1)
	F_show_buffer_usage(m, l2, v22+int32(80))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L8
	} else {
		goto L157
	}
L140:
	;
	v490 = int32(1)
	v492 = *(*int64)(unsafe.Add(mBase, uint32(v22)+80))
	if int64(0) < v492 {
		v506 = v490
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v507 = *(*int64)(unsafe.Add(mBase, uint32(v22)+112))
	if int64(0) < v507 {
		v519 = v490
		goto L145
	} else {
		goto L146
	}
L142:
	;
	v496 = *(*int64)(unsafe.Add(mBase, uint32(v22)+88))
	if int64(0) < v496 {
		v506 = int32(1)
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v500 = *(*int64)(unsafe.Add(mBase, uint32(v22)+96))
	if int64(0) < v500 {
		v506 = int32(1)
		goto L141
	} else {
		goto L144
	}
L144:
	;
	v503 = *(*int64)(unsafe.Add(mBase, uint32(v22)+104))
	v506 = base.B2i32(int64(0) < v503)
	goto L141
L145:
	;
	v523 = *(*int64)(unsafe.Add(mBase, uint32(v22)+192))
	if (v506|v519)&int32(1)|base.B2i32(v523 != int64(0)) != 0 {
		goto L139
	} else {
		goto L149
	}
L146:
	;
	v510 = *(*int64)(unsafe.Add(mBase, uint32(v22)+120))
	if int64(0) < v510 {
		v519 = v490
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v513 = *(*int64)(unsafe.Add(mBase, uint32(v22)+128))
	if int64(0) < v513 {
		v519 = v490
		goto L145
	} else {
		goto L148
	}
L148:
	;
	v516 = *(*int64)(unsafe.Add(mBase, uint32(v22)+136))
	v519 = base.B2i32(int64(0) < v516)
	goto L145
L149:
	;
	v527 = *(*int64)(unsafe.Add(mBase, uint32(v22)+144))
	if int64(0) < v527 {
		goto L139
	} else {
		goto L150
	}
L150:
	;
	v530 = *(*int64)(unsafe.Add(mBase, uint32(v22)+152))
	if int64(0) < v530 {
		goto L139
	} else {
		goto L151
	}
L151:
	;
	v533 = *(*int64)(unsafe.Add(mBase, uint32(v22)+160))
	if v533 != int64(0) {
		goto L139
	} else {
		goto L152
	}
L152:
	;
	v536 = *(*int64)(unsafe.Add(mBase, uint32(v22)+168))
	if v536 != int64(0) {
		goto L139
	} else {
		goto L153
	}
L153:
	;
	v539 = *(*int64)(unsafe.Add(mBase, uint32(v22)+176))
	if v539 != int64(0) {
		goto L139
	} else {
		goto L154
	}
L154:
	;
	v542 = *(*int64)(unsafe.Add(mBase, uint32(v22)+184))
	if v542 != int64(0) {
		goto L139
	} else {
		goto L155
	}
L155:
	;
	v545 = *(*int64)(unsafe.Add(mBase, uint32(v22)+200))
	if v545 == int64(0) {
		goto L127
	} else {
		goto L156
	}
L156:
	;
	goto L139
L157:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v557 - int32(1)
	goto L127
L158:
	;
	v566 = *(*int64)(unsafe.Add(mBase, uint32(v22)+72))
	F_ExplainPropertyFloat(m, int32(_a_F_ExplainOnePlan_15), int32(_a_F_ExplainOnePlan_8), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v566), float64(1e+09)), float64(1000)), int32(3), l2)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L8
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v577 = *(*int64)(unsafe.Add(mBase, uint32(v22)+64))
	F_ExplainPropertyUInteger(m, int32(_a_F_ExplainOnePlan_16), int32(_a_F_ExplainOnePlan_11), int64(base.Ui64(v577+int64(1023))>>(uint(int64(10))%64)), l2)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L8
	} else {
		goto L162
	}
L161:
	;
	goto L160
L162:
	;
	F_ExplainPropertyText(m, int32(_a_F_ExplainOnePlan_17), v447, l2)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L8
	} else {
		goto L163
	}
L163:
	;
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v587 != int32(1) {
		goto L127
	} else {
		goto L164
	}
L164:
	;
	F_show_buffer_usage(m, l2, v22+int32(80))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
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
	m.T0[v604].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, l4, l5)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
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
	v611 = int64(*(*int32)(unsafe.Add(mBase, uint32(v22)+216)))
	v612 = *(*int64)(unsafe.Add(mBase, uint32(v22)+208))
	F_ExecutorEnd(m, v80)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L8
	} else {
		goto L171
	}
L170:
	;
	goto L169
L171:
	;
	F_FreeQueryDesc(m, v80)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L8
	} else {
		goto L172
	}
L172:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L8
	} else {
		goto L173
	}
L173:
	;
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v619 == int32(1) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L8
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v624 = int32(1)
	F___clock_gettime(m, v624, v22+int32(208))
	mBase = m.M
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
	if v628 != v624 {
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
	v655 = m.ExcPending
	if v655 != 0 {
		goto L8
	} else {
		goto L182
	}
L179:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v631 != int32(1) {
		goto L178
	} else {
		goto L180
	}
L180:
	;
	v636 = int64(*(*int32)(unsafe.Add(mBase, uint32(v22)+216)))
	v638 = *(*int64)(unsafe.Add(mBase, uint32(v22)+208))
	F_ExplainPropertyFloat(m, int32(_a_F_ExplainOnePlan_7), int32(_a_F_ExplainOnePlan_8), base.F64_mul(base.F64_add(v134, base.F64_div(base.F64_convert_i64_s(v636-v611+(v638-v612)*int64(1000000000)), float64(1e+09))), float64(1000)), int32(3), l2)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
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
