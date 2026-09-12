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
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
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
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
							F_errmsg_internal(m, int32(_a_F_ExplainIndexScanDetails_7), v8)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ExplainIndexScanDetails_8), int32(4035), int32(_a_F_ExplainIndexScanDetails_9))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
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
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(_a_F_ExplainIndexScanDetails_7), v8)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ExplainIndexScanDetails_8), int32(4035), int32(_a_F_ExplainIndexScanDetails_9))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v123 float64
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int64
	_ = v154
	var v158 int64
	_ = v158
	var v162 int64
	_ = v162
	var v165 int64
	_ = v165
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v172 int64
	_ = v172
	var v175 int64
	_ = v175
	var v178 int64
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int64
	_ = v184
	var v187 int64
	_ = v187
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v194 int64
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int64
	_ = v200
	var v203 int64
	_ = v203
	var v206 int32
	_ = v206
	var v207 int64
	_ = v207
	var v210 int64
	_ = v210
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v240 int32
	_ = v240
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
	var v329 int32
	_ = v329
	var v336 int64
	_ = v336
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int64
	_ = v353
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int64
	_ = v380
	var v381 int64
	_ = v381
	var v384 int64
	_ = v384
	var v385 int64
	_ = v385
	var v388 int64
	_ = v388
	var v389 int64
	_ = v389
	var v392 int64
	_ = v392
	var v393 int64
	_ = v393
	var v396 int64
	_ = v396
	var v397 int64
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int64
	_ = v409
	var v410 int64
	_ = v410
	var v413 int64
	_ = v413
	var v414 int64
	_ = v414
	var v417 int64
	_ = v417
	var v418 int64
	_ = v418
	var v421 int64
	_ = v421
	var v422 int64
	_ = v422
	var v425 int64
	_ = v425
	var v426 int64
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v459 int64
	_ = v459
	var v465 int64
	_ = v465
	var v474 int32
	_ = v474
	var v476 int64
	_ = v476
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int64
	_ = v493
	var v497 int64
	_ = v497
	var v501 int64
	_ = v501
	var v504 int64
	_ = v504
	var v507 int32
	_ = v507
	var v508 int64
	_ = v508
	var v511 int64
	_ = v511
	var v514 int64
	_ = v514
	var v517 int64
	_ = v517
	var v520 int32
	_ = v520
	var v521 int64
	_ = v521
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
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v567 int64
	_ = v567
	var v575 int32
	_ = v575
	var v578 int64
	_ = v578
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
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
	v21 = m.G0
	v23 = v21 - int32(256)
	m.G0 = v23
	v30 = F__emscripten_memset_bulkmem(m, v23-int32(-64), base.I32_extend8_s(v10), int32(144))
	mBase = m.M
	goto L1
L1:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v31 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v36 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v38 = v10
	goto L4
L4:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	F___clock_gettime(m, int32(1), v23+int32(208))
	mBase = m.M
	v45 = int64(*(*int32)(unsafe.Add(mBase, uint32(v23)+216)))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v23)+208))
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOnePlan[0]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	goto L8
L5:
	;
	v37 = int32(1)
	goto L7
L6:
	;
	v37 = int32(4)
	goto L7
L7:
	;
	v38 = v37
	goto L4
L8:
	;
	F_PushCopiedSnapshot(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	if v40 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v54 = v38 | int32(2)
	goto L13
L12:
	;
	v54 = v38
	goto L13
L13:
	;
	if v39 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v57 = v54 | int32(8)
	goto L16
L15:
	;
	v57 = v54
	goto L16
L16:
	;
	F_UpdateActiveSnapshotCommandId(m)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	if l1 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOnePlan[0]))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	goto L27
L19:
	;
	v60 = F_CreateIntoRelDestReceiver(m, l1)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L9
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v62 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v67 = v60
	goto L18
L23:
	;
	v63 = F_CreateExplainSerializeDestReceiver(m, l2)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L9
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOnePlan[2]))
	v67 = v66
	goto L18
L26:
	;
	v67 = v63
	goto L18
L27:
	;
	v72 = F_CreateQueryDesc(m, l0, l3, v70, int32(0), v67, l4, l5, v57)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	v76 = v74 ^ int32(1)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+13)))
	if v79 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v80 = v76 | int32(2)
	goto L31
L30:
	;
	v80 = v76
	goto L31
L31:
	;
	if l1 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v124 != 0 {
		goto L49
	} else {
		goto L50
	}
L33:
	;
	F_ExecutorRun(m, v72, v101, int64(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L9
	} else {
		goto L47
	}
L34:
	;
	v101 = int32(1)
	goto L33
L35:
	;
	F_ExecutorStart(m, v72, v80)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v89 != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v85 != 0 {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v123 = float64(0)
	goto L32
L40:
	;
	F_ExecutorStart(m, v72, v90|v80)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L9
	} else {
		goto L44
	}
L41:
	;
	v90 = int32(64)
	goto L43
L42:
	;
	v90 = int32(0)
	goto L43
L43:
	;
	goto L40
L44:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v95 != int32(1) {
		v123 = float64(0)
		goto L32
	} else {
		goto L45
	}
L45:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v99 != 0 {
		v101 = int32(0)
		goto L33
	} else {
		goto L46
	}
L46:
	;
	goto L34
L47:
	;
	F_ExecutorFinish(m, v72)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	F___clock_gettime(m, int32(1), v23+int32(208))
	mBase = m.M
	v111 = int64(*(*int32)(unsafe.Add(mBase, uint32(v23)+216)))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v23)+208))
	v123 = base.F64_add(base.F64_div(base.F64_convert_i64_s(v111-v45+(v113-v46)*int64(1000000000)), float64(1e+09)), float64(0))
	goto L32
L49:
	;
	v126 = v23 - int32(-64)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	if v127 == int32(12) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	goto L51
L51:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	m.T0[v140].(func(*base.Module, int32))(m, v67)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L9
	} else {
		goto L61
	}
L52:
	;
	goto L51
L53:
	;
	goto L57
L54:
	;
	goto L55
L55:
	;
	v138 = F__emscripten_memset_bulkmem(m, v126, base.I32_extend8_s(int32(0)), int32(144))
	mBase = m.M
	goto L60
L56:
	;
	goto L52
L57:
	;
	v133 = F__emscripten_memcpy_bulkmem(m, v126, v67-int32(-64), int32(144))
	mBase = m.M
	goto L59
L59:
	;
	goto L56
L60:
	;
	goto L52
L61:
	;
	F_ExplainOpenGroup(m, int32(_a_F_ExplainOnePlan_0), int32(0), int32(1), l2)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	F_ExplainPrintPlan(m, l2, v72)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	if l7 != 0 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	if l6 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L65:
	;
	v246 = int32(_a_F_ExplainOnePlan_1)
	F_ExplainOpenGroup(m, v246, v246, int32(1), l2)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L9
	} else {
		goto L94
	}
L66:
	;
	v240 = base.B2i32(l8 != int32(0))
	goto L65
L67:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v151 != 0 {
		goto L66
	} else {
		goto L70
	}
L68:
	;
	v221 = int32(0)
	goto L69
L69:
	;
	v227 = base.B2i32(l8 != int32(0))
	if l8 != 0 {
		v240 = v227
		goto L65
	} else {
		goto L92
	}
L70:
	;
	v152 = int32(1)
	v154 = *(*int64)(unsafe.Add(mBase, uint32(l7)))
	if int64(0) < v154 {
		v168 = v152
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v169 = *(*int64)(unsafe.Add(mBase, uint32(l7)+32))
	if int64(0) < v169 {
		v181 = v152
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v158 = *(*int64)(unsafe.Add(mBase, uint32(l7)+8))
	if int64(0) < v158 {
		v168 = int32(1)
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v162 = *(*int64)(unsafe.Add(mBase, uint32(l7)+16))
	if int64(0) < v162 {
		v168 = int32(1)
		goto L71
	} else {
		goto L74
	}
L74:
	;
	v165 = *(*int64)(unsafe.Add(mBase, uint32(l7)+24))
	v168 = base.B2i32(int64(0) < v165)
	goto L71
L75:
	;
	v182 = int32(1)
	v184 = *(*int64)(unsafe.Add(mBase, uint32(l7)+64))
	if v184 <= int64(0) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	v172 = *(*int64)(unsafe.Add(mBase, uint32(l7)+40))
	if int64(0) < v172 {
		v181 = v152
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v175 = *(*int64)(unsafe.Add(mBase, uint32(l7)+48))
	if int64(0) < v175 {
		v181 = v152
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v178 = *(*int64)(unsafe.Add(mBase, uint32(l7)+56))
	v181 = base.B2i32(int64(0) < v178)
	goto L75
L79:
	;
	v187 = *(*int64)(unsafe.Add(mBase, uint32(l7)+72))
	v190 = base.B2i32(int64(0) < v187)
	goto L81
L80:
	;
	v190 = v182
	goto L81
L81:
	;
	v191 = *(*int64)(unsafe.Add(mBase, uint32(l7)+80))
	if v191 == int64(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v194 = *(*int64)(unsafe.Add(mBase, uint32(l7)+88))
	v197 = base.B2i32(v194 != int64(0))
	goto L84
L83:
	;
	v197 = v182
	goto L84
L84:
	;
	v198 = int32(1)
	v200 = *(*int64)(unsafe.Add(mBase, uint32(l7)+96))
	if v200 == int64(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v203 = *(*int64)(unsafe.Add(mBase, uint32(l7)+104))
	v206 = base.B2i32(v203 != int64(0))
	goto L87
L86:
	;
	v206 = v198
	goto L87
L87:
	;
	v207 = *(*int64)(unsafe.Add(mBase, uint32(l7)+112))
	if v207 == int64(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v210 = *(*int64)(unsafe.Add(mBase, uint32(l7)+120))
	v213 = base.B2i32(v210 != int64(0))
	goto L90
L89:
	;
	v213 = v198
	goto L90
L90:
	;
	if (v168|v181|v190|v197)&int32(1) != 0 {
		goto L66
	} else {
		goto L91
	}
L91:
	;
	v221 = v213 | v206
	goto L69
L92:
	;
	if v221&int32(1) == int32(0) {
		goto L64
	} else {
		goto L93
	}
L93:
	;
	v240 = v227
	goto L65
L94:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v251 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	F_ExplainIndentText(m, l2)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L9
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	if l7 != 0 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v256, int32(_a_F_ExplainOnePlan_2))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L9
	} else {
		goto L99
	}
L99:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v260 + int32(1)
	goto L97
L100:
	;
	F_show_buffer_usage(m, l2, l7)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L9
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	if v240 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	goto L102
L104:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v307 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L105:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l8)+8))
	v270 = v268 + int32(1023)
	v271 = int32(10)
	v273 = base.I64_extend_i32_u(int32(base.Ui32(v270) >> (uint(v271) % 32)))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l8)+12))
	v278 = base.I64_extend_i32_u(int32(base.Ui32(v270-v274) >> (uint(v271) % 32)))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v279 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	F_ExplainIndentText(m, l2)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L9
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	F_ExplainPropertyInteger(m, int32(_a_F_ExplainOnePlan_10), int32(_a_F_ExplainOnePlan_11), v278, l2)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L9
	} else {
		goto L112
	}
L109:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v273
	*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v278
	F_appendStringInfo(m, v284, int32(_a_F_ExplainOnePlan_12), v23+int32(48))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L9
	} else {
		goto L110
	}
L110:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v292, int32(10))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L9
	} else {
		goto L111
	}
L111:
	;
	goto L104
L112:
	;
	F_ExplainPropertyInteger(m, int32(_a_F_ExplainOnePlan_13), int32(_a_F_ExplainOnePlan_11), v273, l2)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L9
	} else {
		goto L113
	}
L113:
	;
	goto L104
L114:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v310 - int32(1)
	goto L116
L115:
	;
	goto L116
L116:
	;
	F_ExplainCloseGroup(m, int32(_a_F_ExplainOnePlan_1), int32(1), l2)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L9
	} else {
		goto L117
	}
L117:
	;
	goto L64
L118:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v345 == int32(1) {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
	if v329&int32(1) == int32(0) {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v336 = *(*int64)(unsafe.Add(mBase, uint32(l6)))
	F_ExplainPropertyFloat(m, int32(_a_F_ExplainOnePlan_14), int32(_a_F_ExplainOnePlan_8), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v336), float64(1e+09)), float64(1000)), int32(3), l2)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L9
	} else {
		goto L121
	}
L121:
	;
	goto L118
L122:
	;
	F_ExplainPrintTriggers(m, l2, v72)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L9
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v350 != int32(1) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	goto L124
L126:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v438 != 0 {
		goto L138
	} else {
		goto L139
	}
L127:
	;
	v353 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+248)) = v353
	*(*int64)(unsafe.Add(mBase, uint32(v23)+240)) = v353
	*(*int64)(unsafe.Add(mBase, uint32(v23)+232)) = v353
	*(*int64)(unsafe.Add(mBase, uint32(v23)+224)) = v353
	*(*int64)(unsafe.Add(mBase, uint32(v23)+216)) = v353
	*(*int64)(unsafe.Add(mBase, uint32(v23)+208)) = v353
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v72)+40))
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+176)))
	if v366&int32(1) == int32(0) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v365)+180))
	if v371 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v373 = v23 + int32(208)
	v375 = v371 + int32(8)
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	*(*int32)(unsafe.Add(mBase, uint32(v373))) = v376 + v377
	v380 = *(*int64)(unsafe.Add(mBase, uint32(v373)+8))
	v381 = *(*int64)(unsafe.Add(mBase, uint32(v375)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v373)+8)) = v380 + v381
	v384 = *(*int64)(unsafe.Add(mBase, uint32(v373)+16))
	v385 = *(*int64)(unsafe.Add(mBase, uint32(v375)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v373)+16)) = v384 + v385
	v388 = *(*int64)(unsafe.Add(mBase, uint32(v373)+24))
	v389 = *(*int64)(unsafe.Add(mBase, uint32(v375)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v373)+24)) = v388 + v389
	v392 = *(*int64)(unsafe.Add(mBase, uint32(v373)+32))
	v393 = *(*int64)(unsafe.Add(mBase, uint32(v375)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v373)+32)) = v392 + v393
	v396 = *(*int64)(unsafe.Add(mBase, uint32(v373)+40))
	v397 = *(*int64)(unsafe.Add(mBase, uint32(v375)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v373)+40)) = v396 + v397
	goto L132
L130:
	;
	v401 = v365
	goto L131
L131:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+184))
	if v402 != 0 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v72)+40))
	v401 = v400
	goto L131
L133:
	;
	v404 = v23 + int32(208)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	*(*int32)(unsafe.Add(mBase, uint32(v404))) = v405 + v406
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v404)+8))
	v410 = *(*int64)(unsafe.Add(mBase, uint32(v402)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v404)+8)) = v409 + v410
	v413 = *(*int64)(unsafe.Add(mBase, uint32(v404)+16))
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v402)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v404)+16)) = v413 + v414
	v417 = *(*int64)(unsafe.Add(mBase, uint32(v404)+24))
	v418 = *(*int64)(unsafe.Add(mBase, uint32(v402)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v404)+24)) = v417 + v418
	v421 = *(*int64)(unsafe.Add(mBase, uint32(v404)+32))
	v422 = *(*int64)(unsafe.Add(mBase, uint32(v402)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v404)+32)) = v421 + v422
	v425 = *(*int64)(unsafe.Add(mBase, uint32(v404)+40))
	v426 = *(*int64)(unsafe.Add(mBase, uint32(v402)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v404)+40)) = v425 + v426
	goto L136
L134:
	;
	v430 = v401
	goto L135
L135:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)+176))
	F_ExplainPrintJIT(m, l2, v431, v23+int32(208))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L9
	} else {
		goto L137
	}
L136:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v72)+40))
	v430 = v429
	goto L135
L137:
	;
	goto L126
L138:
	;
	v439 = int32(_a_F_ExplainOnePlan_3)
	F_ExplainOpenGroup(m, v439, v439, int32(1), l2)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L9
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v604 = *(*int32)(unsafe.Add(mBase, _c_F_ExplainOnePlan[1]))
	if v604 != 0 {
		goto L186
	} else {
		goto L187
	}
L141:
	;
	if v438 == int32(1) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v448 = int32(_a_F_ExplainOnePlan_4)
	goto L144
L143:
	;
	v448 = int32(_a_F_ExplainOnePlan_5)
	goto L144
L144:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v449 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	F_ExplainCloseGroup(m, int32(_a_F_ExplainOnePlan_3), int32(1), l2)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L9
	} else {
		goto L185
	}
L146:
	;
	F_ExplainIndentText(m, l2)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L9
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v562 == int32(1) {
		goto L177
	} else {
		goto L178
	}
L149:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v455 == int32(1) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v487 != int32(1) {
		goto L145
	} else {
		goto L156
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v448
	v459 = *(*int64)(unsafe.Add(mBase, uint32(v23)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = int64(base.Ui64(v459+int64(1023)) >> (uint(int64(10)) % 64))
	v465 = *(*int64)(unsafe.Add(mBase, uint32(v23)+72))
	*(*float64)(unsafe.Add(mBase, uint32(v23))) = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v465), float64(1e+09)), float64(1000))
	F_appendStringInfo(m, v454, int32(_a_F_ExplainOnePlan_6), v23)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L9
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v448
	v476 = *(*int64)(unsafe.Add(mBase, uint32(v23)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = int64(base.Ui64(v476+int64(1023)) >> (uint(int64(10)) % 64))
	F_appendStringInfo(m, v454, int32(_a_F_ExplainOnePlan_9), v23+int32(32))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L9
	} else {
		goto L155
	}
L154:
	;
	goto L150
L155:
	;
	goto L150
L156:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v490 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v550 + int32(1)
	F_show_buffer_usage(m, l2, v23+int32(80))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L9
	} else {
		goto L176
	}
L158:
	;
	v491 = int32(1)
	v493 = *(*int64)(unsafe.Add(mBase, uint32(v23)+80))
	if int64(0) < v493 {
		v507 = v491
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v508 = *(*int64)(unsafe.Add(mBase, uint32(v23)+112))
	if int64(0) < v508 {
		v520 = v491
		goto L163
	} else {
		goto L164
	}
L160:
	;
	v497 = *(*int64)(unsafe.Add(mBase, uint32(v23)+88))
	if int64(0) < v497 {
		v507 = int32(1)
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v501 = *(*int64)(unsafe.Add(mBase, uint32(v23)+96))
	if int64(0) < v501 {
		v507 = int32(1)
		goto L159
	} else {
		goto L162
	}
L162:
	;
	v504 = *(*int64)(unsafe.Add(mBase, uint32(v23)+104))
	v507 = base.B2i32(int64(0) < v504)
	goto L159
L163:
	;
	v521 = *(*int64)(unsafe.Add(mBase, uint32(v23)+192))
	if v521 != int64(0) {
		goto L157
	} else {
		goto L167
	}
L164:
	;
	v511 = *(*int64)(unsafe.Add(mBase, uint32(v23)+120))
	if int64(0) < v511 {
		v520 = v491
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v514 = *(*int64)(unsafe.Add(mBase, uint32(v23)+128))
	if int64(0) < v514 {
		v520 = v491
		goto L163
	} else {
		goto L166
	}
L166:
	;
	v517 = *(*int64)(unsafe.Add(mBase, uint32(v23)+136))
	v520 = base.B2i32(int64(0) < v517)
	goto L163
L167:
	;
	if (v507|v520)&int32(1) != 0 {
		goto L157
	} else {
		goto L168
	}
L168:
	;
	v527 = *(*int64)(unsafe.Add(mBase, uint32(v23)+144))
	if int64(0) < v527 {
		goto L157
	} else {
		goto L169
	}
L169:
	;
	v530 = *(*int64)(unsafe.Add(mBase, uint32(v23)+152))
	if int64(0) < v530 {
		goto L157
	} else {
		goto L170
	}
L170:
	;
	v533 = *(*int64)(unsafe.Add(mBase, uint32(v23)+160))
	if v533 != int64(0) {
		goto L157
	} else {
		goto L171
	}
L171:
	;
	v536 = *(*int64)(unsafe.Add(mBase, uint32(v23)+168))
	if v536 != int64(0) {
		goto L157
	} else {
		goto L172
	}
L172:
	;
	v539 = *(*int64)(unsafe.Add(mBase, uint32(v23)+176))
	if v539 != int64(0) {
		goto L157
	} else {
		goto L173
	}
L173:
	;
	v542 = *(*int64)(unsafe.Add(mBase, uint32(v23)+184))
	if v542 != int64(0) {
		goto L157
	} else {
		goto L174
	}
L174:
	;
	v545 = *(*int64)(unsafe.Add(mBase, uint32(v23)+200))
	if v545 == int64(0) {
		goto L145
	} else {
		goto L175
	}
L175:
	;
	goto L157
L176:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v558 - int32(1)
	goto L145
L177:
	;
	v567 = *(*int64)(unsafe.Add(mBase, uint32(v23)+72))
	F_ExplainPropertyFloat(m, int32(_a_F_ExplainOnePlan_15), int32(_a_F_ExplainOnePlan_8), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v567), float64(1e+09)), float64(1000)), int32(3), l2)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L9
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v578 = *(*int64)(unsafe.Add(mBase, uint32(v23)+64))
	F_ExplainPropertyUInteger(m, int32(_a_F_ExplainOnePlan_16), int32(_a_F_ExplainOnePlan_11), int64(base.Ui64(v578+int64(1023))>>(uint(int64(10))%64)), l2)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L9
	} else {
		goto L181
	}
L180:
	;
	goto L179
L181:
	;
	F_ExplainPropertyText(m, int32(_a_F_ExplainOnePlan_17), v448, l2)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L9
	} else {
		goto L182
	}
L182:
	;
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v588 != int32(1) {
		goto L145
	} else {
		goto L183
	}
L183:
	;
	F_show_buffer_usage(m, l2, v23+int32(80))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L9
	} else {
		goto L184
	}
L184:
	;
	goto L145
L185:
	;
	goto L140
L186:
	;
	m.T0[v604].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, l4, l5)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L9
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	F___clock_gettime(m, int32(1), v23+int32(208))
	mBase = m.M
	v611 = int64(*(*int32)(unsafe.Add(mBase, uint32(v23)+216)))
	v612 = *(*int64)(unsafe.Add(mBase, uint32(v23)+208))
	F_ExecutorEnd(m, v72)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L9
	} else {
		goto L190
	}
L189:
	;
	goto L188
L190:
	;
	F_FreeQueryDesc(m, v72)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L9
	} else {
		goto L191
	}
L191:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L9
	} else {
		goto L192
	}
L192:
	;
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v619 == int32(1) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L9
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v624 = int32(1)
	F___clock_gettime(m, v624, v23+int32(208))
	mBase = m.M
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
	if v628 != v624 {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	goto L195
L197:
	;
	F_ExplainCloseGroup(m, int32(_a_F_ExplainOnePlan_0), int32(1), l2)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L9
	} else {
		goto L201
	}
L198:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v631 != int32(1) {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	v636 = int64(*(*int32)(unsafe.Add(mBase, uint32(v23)+216)))
	v638 = *(*int64)(unsafe.Add(mBase, uint32(v23)+208))
	F_ExplainPropertyFloat(m, int32(_a_F_ExplainOnePlan_7), int32(_a_F_ExplainOnePlan_8), base.F64_mul(base.F64_add(v123, base.F64_div(base.F64_convert_i64_s(v636-v611+(v638-v612)*int64(1000000000)), float64(1e+09))), float64(1000)), int32(3), l2)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L9
	} else {
		goto L200
	}
L200:
	;
	goto L197
L201:
	;
	m.G0 = v23 + int32(256)
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
