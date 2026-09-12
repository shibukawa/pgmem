package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ATExecSetOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	v11 = m.G0
	v13 = v11 - int32(208)
	m.G0 = v13
	v17 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
		v20 = F_SearchSysCacheAttName(m, v19, l2)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			if v20 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
				v24 = v22 + v23
				v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+74)))
				if v25 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l2
							F_errmsg(m, int32(687318), v13+int32(16))
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
								return
							} else {
								F_errfinish(m, int32(486904), int32(9083), int32(136245))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
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
					v33 = F_SysCacheGetAttr(m, int32(6), v20, int32(23), v13+int32(207))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+207)))
						if v35 != 0 {
							v36 = int32(0)
						} else {
							v36 = v33
						}
						v37 = int32(0)
						v40 = F_transformRelOptions(m, v36, l3, v37, v37, v37, l4)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							v43 = F_attribute_reloptions(m, v40, int32(1))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								v45 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v13)+88)) = uint8(v45)
								v47 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v13)+80)) = v47
								*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v47
								*(*uint8)(unsafe.Add(mBase, uint32(v13)+56)) = uint8(v45)
								*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v47
								*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v47
								*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v47
								*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v47
								if v40 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v40
								} else {
									v62 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v13)+86)) = uint8(v62)
								}
								v64 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v13)+54)) = uint8(v64)
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
								v73 = F_heap_modify_tuple(m, v20, v66, v13+int32(96), v13-int32(-64), v13+int32(32))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return
								} else {
									F_CatalogTupleUpdate(m, v17, v73+int32(4), v73)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return
									} else {
										v80 = *(*int32)(unsafe.Add(mBase, _consts[380]))
										if v80 != 0 {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
											v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+74)))
											v84 = int32(0)
											F_RunObjectPostAlterHook(m, int32(1259), v82, v83, v84, v84)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
												v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v25
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
												F_pfree(m, v73)
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return
												} else {
													F_ReleaseCatCache(m, v20)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return
													} else {
														F_sequence_close(m, v17, int32(3))
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return
														} else {
															m.G0 = v13 + int32(208)
															return
														}
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
											v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v25
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
											F_pfree(m, v73)
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return
											} else {
												F_ReleaseCatCache(m, v20)
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return
												} else {
													F_sequence_close(m, v17, int32(3))
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return
													} else {
														m.G0 = v13 + int32(208)
														return
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
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return
				} else {
					F_errcode(m, int32(50360452))
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return
					} else {
						v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v110 + int32(4)
						F_errmsg(m, int32(70796), v13)
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return
						} else {
							F_errfinish(m, int32(486904), int32(9075), int32(136245))
							mBase = m.M
							v122 = m.ExcPending
							if v122 != 0 {
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
func F_ATPostAlterTypeParse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
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
	var v326 int32
	_ = v326
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
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v445 int32
	_ = v445
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v650 int32
	_ = v650
	var v658 int32
	_ = v658
	v7 = int32(0)
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	v24 = F_raw_parser(m, l3, v7)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_relation_close(m, v650, int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L2
	} else {
		goto L136
	}
L2:
	;
	return
L3:
	;
	if v24 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v29 = F_relation_open(m, l1, int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if int32(0) < v31 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v650 = v29
	goto L1
L8:
	;
	v40 = v7
	v45 = v7
	goto L11
L9:
	;
	v104 = v7
	goto L10
L10:
	;
	v112 = F_relation_open(m, l1, int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L30
	}
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v40<<(uint(int32(2))%32))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	switch v58 - int32(204) {
	case 0:
		goto L16
	case 1:
		goto L15
	default:
		goto L14
	}
L12:
	;
	v104 = v88
	goto L10
L13:
	;
	v90 = v40 + int32(1)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v90 < v91 {
		v40 = v90
		v45 = v88
		goto L11
	} else {
		goto L29
	}
L14:
	;
	if v58 != int32(146) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v65 = F_transformStatsStmt(m, l1, v57, l3)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L2
	} else {
		goto L19
	}
L16:
	;
	v61 = F_transformIndexStmt(m, l1, v57, l3)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v63 = F_lappend(m, v45, v61)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v88 = v63
	goto L13
L19:
	;
	v67 = F_lappend(m, v45, v65)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v88 = v67
	goto L13
L21:
	;
	v71 = F_lappend(m, v45, v57)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v77 = F_transformAlterTableStmt(m, l1, v57, l3, v19+int32(-4), v19+int32(-8))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L2
	} else {
		goto L25
	}
L24:
	;
	v88 = v71
	goto L13
L25:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	v80 = F_list_concat(m, v45, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v82 = F_lappend(m, v80, v77)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v85 = F_list_concat(m, v82, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v88 = v85
	goto L13
L29:
	;
	goto L12
L30:
	;
	if v104 == int32(0) {
		v650 = v112
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v116 <= int32(0) {
		v650 = v112
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v133 = v7
	goto L35
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L2
	} else {
		goto L133
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L2
	} else {
		goto L130
	}
L35:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137+v133<<(uint(int32(2))%32))))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v112)+56))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v143 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L2
	} else {
		goto L127
	}
L37:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	switch v237 - int32(146) {
	case 0:
		goto L53
	case 1, 2, 3, 4:
		goto L51
	case 5:
		goto L52
	default:
		goto L54
	}
L38:
	;
	v197 = F_palloc0(m, int32(140))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L2
	} else {
		goto L45
	}
L39:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	if v146 <= int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	v157 = int32(0)
	goto L41
L41:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v149+v157<<(uint(int32(2))%32))))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	if v173 == v142 {
		v226 = v172
		goto L37
	} else {
		goto L43
	}
L42:
	;
	goto L38
L43:
	;
	v176 = v157 + int32(1)
	if v146 != v176 {
		v157 = v176
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v142
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v112)+48))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+4)) = uint8(v203)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v112)+52))
	v206 = F_CreateTupleDescCopyConstr(m, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+8)) = v206
	*(*int64)(unsafe.Add(mBase, uint32(v197)+88)) = int64(0)
	v211 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+84)) = uint8(v211)
	v213 = int32(28672)
	*(*uint16)(unsafe.Add(mBase, uint32(v197)+96)) = uint16(v213)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v216 = F_lappend(m, v215, v197)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v216
	v226 = v197
	goto L37
L48:
	;
	goto L36
L49:
	;
	v591 = v133 + int32(1)
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v592 <= v591 {
		v650 = v112
		goto L1
	} else {
		goto L126
	}
L50:
	;
	v557 = F_GetComment(m, l0, int32(3381), int32(0))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L2
	} else {
		goto L123
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L2
	} else {
		goto L120
	}
L52:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+4)))
	if v503 == int32(67) {
		goto L111
	} else {
		goto L112
	}
L53:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	if v286 == int32(0) {
		goto L49
	} else {
		goto L68
	}
L54:
	;
	switch v237 - int32(204) {
	case 0:
		goto L55
	case 1:
		goto L50
	default:
		goto L51
	}
L55:
	;
	if l5 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v267 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+70)) = uint8(v267)
	v271 = F_GetComment(m, l0, int32(1259), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L2
	} else {
		goto L65
	}
L57:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v141)+20))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v141)+36))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+64)))
	v246 = F_CheckIndexCompatible(m, l0, v242, v243, v244, v245)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	if v246 == int32(0) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v251 = F_index_open(m, l0, int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v251)+48))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+119)))
	if v254 != int32(73) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v251)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+48)) = v257
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v251)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+52)) = v259
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v251)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+56)) = v261
	goto L63
L62:
	;
	goto L63
L63:
	;
	F_relation_close(m, v251, int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	goto L56
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141)+40)) = v271
	v275 = F_palloc0(m, int32(32))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275)+20)) = v141
	*(*int64)(unsafe.Add(mBase, uint32(v275))) = int64(64424509587)
	v281 = v226 + int32(32)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	v283 = F_lappend(m, v282, v275)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v281))) = v283
	goto L49
L68:
	;
	v289 = int32(0)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	if v290 <= v289 {
		goto L49
	} else {
		goto L69
	}
L69:
	;
	v294 = v289
	goto L70
L70:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v286)+12))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v311+v294<<(uint(int32(2))%32))))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)+4))
	switch v316 - int32(14) {
	case 0:
		goto L74
	default:
		goto L33
	case 2:
		goto L73
	}
L71:
	;
	goto L49
L72:
	;
	v500 = v294 + int32(1)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	if v500 < v501 {
		v294 = v500
		goto L70
	} else {
		goto L110
	}
L73:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v315)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v367)+100)) = l2
	if l5 != 0 {
		goto L88
	} else {
		goto L89
	}
L74:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v315)+20))
	v320 = F_get_constraint_index(m, l0)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	if l5 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v349 = F_GetComment(m, v320, int32(1259), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L2
	} else {
		goto L85
	}
L77:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v319)+12))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v319)+20))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v319)+36))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319)+64)))
	v326 = F_CheckIndexCompatible(m, v320, v322, v323, v324, v325)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	if v326 == int32(0) {
		goto L76
	} else {
		goto L79
	}
L79:
	;
	v331 = F_index_open(m, v320, int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v331)+48))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+119)))
	if v334 != int32(73) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v331)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v319)+48)) = v337
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v331)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v319)+52)) = v339
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v331)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v319)+56)) = v341
	goto L83
L82:
	;
	goto L83
L83:
	;
	F_relation_close(m, v331, int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	goto L76
L85:
	;
	v351 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v319)+70)) = uint8(v351)
	*(*int32)(unsafe.Add(mBase, uint32(v319)+40)) = v349
	*(*int32)(unsafe.Add(mBase, uint32(v315)+4)) = int32(15)
	v357 = v226 + int32(32)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	v359 = F_lappend(m, v358, v315)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L2
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357))) = v359
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	F_RebuildConstraintComment(m, v226, int32(4), l0, v112, int32(0), v364)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L2
	} else {
		goto L87
	}
L87:
	;
	goto L72
L88:
	;
	v464 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v367)+60)) = uint8(v464)
	*(*int32)(unsafe.Add(mBase, uint32(v315)+4)) = int32(17)
	v469 = v226 + int32(36)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v471 = F_lappend(m, v470, v315)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L2
	} else {
		goto L107
	}
L89:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	if v369 != int32(9) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v226)+80))
	if v372 != 0 {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	v374 = F_SearchSysCache1(m, int32(19), l0)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L2
	} else {
		goto L92
	}
L92:
	;
	if v374 == int32(0) {
		goto L48
	} else {
		goto L93
	}
L93:
	;
	v380 = F_SysCacheGetAttrNotNull(m, int32(19), v374, int32(23))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v382 = F_pg_detoast_datum(m, v380)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	if v384 != int32(1) {
		goto L34
	} else {
		goto L96
	}
L96:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	if v387 != 0 {
		goto L34
	} else {
		goto L97
	}
L97:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v382)+12))
	if v388 != int32(26) {
		goto L34
	} else {
		goto L98
	}
L98:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v382)+16))
	if int32(0) < v391 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v367)+96))
	v401 = v396
	v404 = int32(0)
	goto L102
L100:
	;
	goto L101
L101:
	;
	F_ReleaseCatCache(m, v374)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L2
	} else {
		goto L106
	}
L102:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v382+int32(24)+v404<<(uint(int32(2))%32))))
	v420 = F_lappend_oid(m, v401, v419)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L2
	} else {
		goto L104
	}
L103:
	;
	goto L101
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v367)+96)) = v420
	v424 = v404 + int32(1)
	if v424 != v391 {
		v401 = v420
		v404 = v424
		goto L102
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	goto L88
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v469))) = v471
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v367)+8))
	if v474 == int32(0) {
		goto L72
	} else {
		goto L108
	}
L108:
	;
	F_RebuildConstraintComment(m, v226, int32(5), l0, v112, int32(0), v474)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L2
	} else {
		goto L109
	}
L109:
	;
	goto L72
L110:
	;
	goto L71
L111:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v141)+16))
	v508 = F_palloc0(m, int32(32))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L2
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L2
	} else {
		goto L117
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v508)+20)) = v141
	*(*int64)(unsafe.Add(mBase, uint32(v508))) = int64(77309411475)
	v514 = v226 + int32(36)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v514)))
	v516 = F_lappend(m, v515, v508)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L2
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v514))) = v516
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v506)+8))
	F_RebuildConstraintComment(m, v226, int32(5), l0, int32(0), v521, v522)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L2
	} else {
		goto L116
	}
L116:
	;
	goto L49
L117:
	;
	v529 = int32(*(*int8)(unsafe.Add(mBase, uint32(v141)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v529
	F_errmsg_internal(m, int32(476907), v19+int32(-16))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(486904), int32(15807), int32(355400))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v545
	F_errmsg_internal(m, int32(477102), v21)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(486904), int32(15825), int32(355400))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141)+20)) = v557
	v561 = F_palloc0(m, int32(32))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L2
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v561)+20)) = v141
	*(*int64)(unsafe.Add(mBase, uint32(v561))) = int64(279172874387)
	v567 = v226 + int32(60)
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	v569 = F_lappend(m, v568, v561)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L2
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v567))) = v569
	goto L49
L126:
	;
	v133 = v591
	goto L35
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = l0
	F_errmsg_internal(m, int32(40615), v19+int32(-32))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L2
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(486904), int32(15929), int32(22495))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L2
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	F_errmsg_internal(m, int32(25685), int32(0))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L2
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(486904), int32(15939), int32(22495))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L2
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v315)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v626
	F_errmsg_internal(m, int32(476907), v19+int32(-48))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L2
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(486904), int32(15780), int32(355400))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L2
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	m.G0 = v21 - int32(-64)
	return
}
func F_ATPrepCmd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
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
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v552 int32
	_ = v552
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v654 int32
	_ = v654
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v715 int32
	_ = v715
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v903 int32
	_ = v903
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v936 int32
	_ = v936
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v981 int32
	_ = v981
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1025 int32
	_ = v1025
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1077 int32
	_ = v1077
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1202 int32
	_ = v1202
	var v1206 int32
	_ = v1206
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1228 int32
	_ = v1228
	var v1232 int32
	_ = v1232
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1251 int32
	_ = v1251
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1272 int32
	_ = v1272
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1283 int32
	_ = v1283
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1297 int32
	_ = v1297
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1318 int32
	_ = v1318
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1343 int32
	_ = v1343
	var v1350 int32
	_ = v1350
	var v1354 int32
	_ = v1354
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1375 int32
	_ = v1375
	var v1380 int32
	_ = v1380
	var v1384 int32
	_ = v1384
	var v1387 int32
	_ = v1387
	var v1391 int32
	_ = v1391
	var v1396 int32
	_ = v1396
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1413 int32
	_ = v1413
	var v1418 int32
	_ = v1418
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1435 int32
	_ = v1435
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1451 int32
	_ = v1451
	var v1455 int32
	_ = v1455
	var v1460 int32
	_ = v1460
	var v1464 int32
	_ = v1464
	var v1467 int32
	_ = v1467
	var v1473 int32
	_ = v1473
	var v1478 int32
	_ = v1478
	var v1483 int32
	_ = v1483
	var v1486 int32
	_ = v1486
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1501 int32
	_ = v1501
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1520 int32
	_ = v1520
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1535 int32
	_ = v1535
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1544 int32
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1563 int32
	_ = v1563
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1578 int32
	_ = v1578
	var v1582 int32
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1596 int32
	_ = v1596
	var v1599 int32
	_ = v1599
	var v1603 int32
	_ = v1603
	var v1608 int32
	_ = v1608
	var v1612 int32
	_ = v1612
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1621 int32
	_ = v1621
	var v1624 int32
	_ = v1624
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1637 int32
	_ = v1637
	var v1642 int32
	_ = v1642
	var v1646 int32
	_ = v1646
	var v1649 int32
	_ = v1649
	var v1653 int32
	_ = v1653
	var v1658 int32
	_ = v1658
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1669 int32
	_ = v1669
	var v1674 int32
	_ = v1674
	var v1678 int32
	_ = v1678
	var v1682 int32
	_ = v1682
	var v1685 int32
	_ = v1685
	var v1690 int32
	_ = v1690
	var v1693 int32
	_ = v1693
	var v1698 int32
	_ = v1698
	var v1701 int32
	_ = v1701
	var v1704 int32
	_ = v1704
	var v1708 int32
	_ = v1708
	var v1712 int32
	_ = v1712
	var v1716 int32
	_ = v1716
	var v1720 int32
	_ = v1720
	var v1724 int32
	_ = v1724
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1733 int32
	_ = v1733
	var v1738 int32
	_ = v1738
	var v1742 int32
	_ = v1742
	var v1745 int32
	_ = v1745
	var v1750 int32
	_ = v1750
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1762 int32
	_ = v1762
	var v1766 int32
	_ = v1766
	var v1771 int32
	_ = v1771
	var v1775 int32
	_ = v1775
	var v1778 int32
	_ = v1778
	var v1782 int32
	_ = v1782
	var v1787 int32
	_ = v1787
	var v1791 int32
	_ = v1791
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1803 int32
	_ = v1803
	var v1807 int32
	_ = v1807
	var v1812 int32
	_ = v1812
	var v1816 int32
	_ = v1816
	var v1819 int32
	_ = v1819
	var v1823 int32
	_ = v1823
	var v1828 int32
	_ = v1828
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1857 int32
	_ = v1857
	var v1865 int32
	_ = v1865
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1902 int32
	_ = v1902
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1918 int32
	_ = v1918
	var v1922 int32
	_ = v1922
	var v1927 int32
	_ = v1927
	var v1934 int32
	_ = v1934
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1960 int32
	_ = v1960
	var v1963 int32
	_ = v1963
	var v1968 int32
	_ = v1968
	var v1970 int32
	_ = v1970
	var v1972 int32
	_ = v1972
	var v1985 int32
	_ = v1985
	var v1987 int32
	_ = v1987
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2011 int32
	_ = v2011
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2027 int32
	_ = v2027
	var v2031 int32
	_ = v2031
	var v2036 int32
	_ = v2036
	v8 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(144)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v26 == v8 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+131)))
	if v127 != int32(1) {
		goto L21
	} else {
		goto L22
	}
L2:
	;
	v84 = F_palloc0(m, int32(140))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v29 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v43 = int32(0)
	goto L5
L5:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v32+v43<<(uint(int32(2))%32))))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v58 == v25 {
		v118 = v57
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L2
L7:
	;
	v61 = v43 + int32(1)
	if v29 != v61 {
		v43 = v61
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	return
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v25
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+4)) = uint8(v90)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v93 = F_CreateTupleDescCopyConstr(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+8)) = v93
	*(*int64)(unsafe.Add(mBase, uint32(v84)+88)) = int64(0)
	v98 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+84)) = uint8(v98)
	v100 = int32(28672)
	*(*uint16)(unsafe.Add(mBase, uint32(v84)+96)) = uint16(v100)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v103 = F_lappend(m, v102, v84)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v103
	v118 = v84
	goto L1
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2011 = m.ExcPending
	if v2011 != 0 {
		goto L9
	} else {
		goto L570
	}
L14:
	;
	v2000 = v118 + v1985<<(uint(int32(2))%32) + int32(16)
	v2001 = *(*int32)(unsafe.Add(mBase, uint32(v2000)))
	v2002 = F_lappend(m, v2001, v1987)
	mBase = m.M
	v2003 = m.ExcPending
	if v2003 != 0 {
		goto L9
	} else {
		goto L569
	}
L15:
	;
	v1835 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v1836 = m.ExcPending
	if v1836 != 0 {
		goto L9
	} else {
		goto L534
	}
L16:
	;
	v1831 = int32(2665)
	v1832 = int32(9)
	goto L15
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1816 = m.ExcPending
	if v1816 != 0 {
		goto L9
	} else {
		goto L530
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1791 = m.ExcPending
	if v1791 != 0 {
		goto L9
	} else {
		goto L525
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L9
	} else {
		goto L521
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1750 = m.ExcPending
	if v1750 != 0 {
		goto L9
	} else {
		goto L516
	}
L21:
	;
	v189 = int32(11)
	v190 = F_copyObjectImpl(m, l2)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L9
	} else {
		goto L78
	}
L22:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v130 == int32(61) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v134 = m.G0
	v136 = v134 + int32(-64)
	m.G0 = v136
	v140 = F_table_open(m, int32(2611), int32(3))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L9
	} else {
		goto L25
	}
L24:
	;
	if v162 != 0 {
		goto L20
	} else {
		goto L37
	}
L25:
	;
	F_ScanKeyInit(m, v134+int32(-48), int32(1), int32(3), int32(184), v133)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v150 = int32(1)
	v155 = F_systable_beginscan(m, v140, int32(2680), v150, int32(0), v150, v134+int32(-48))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v157 = F_systable_getnext(m, v155)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	if v157 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+22)))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159+v160)+12)))
	F_systable_endscan(m, v155)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L9
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L9
	} else {
		goto L34
	}
L32:
	;
	F_sequence_close(m, v140, int32(3))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	m.G0 = v136 - int32(-64)
	goto L24
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v133
	F_errmsg_internal(m, int32(245361), v136)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(486058), int32(654), int32(321558))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	goto L21
L38:
	;
	F_ATSimplePermissions(m, int32(0), l1, int32(305))
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L9
	} else {
		goto L514
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1728 = m.ExcPending
	if v1728 != 0 {
		goto L9
	} else {
		goto L511
	}
L40:
	;
	F_ATSimplePermissions(m, int32(61), l1, int32(256))
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L9
	} else {
		goto L510
	}
L41:
	;
	F_ATSimplePermissions(m, int32(60), l1, int32(256))
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L9
	} else {
		goto L509
	}
L42:
	;
	F_ATSimplePermissions(m, int32(59), l1, int32(320))
	mBase = m.M
	v1716 = m.ExcPending
	if v1716 != 0 {
		goto L9
	} else {
		goto L508
	}
L43:
	;
	F_ATSimplePermissions(m, int32(58), l1, int32(32))
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L9
	} else {
		goto L507
	}
L44:
	;
	F_ATSimplePermissions(m, v192, l1, int32(257))
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L9
	} else {
		goto L506
	}
L45:
	;
	F_ATSimplePermissions(m, v192, l1, int32(289))
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L9
	} else {
		goto L504
	}
L46:
	;
	F_ATSimplePermissions(m, int32(53), l1, int32(261))
	mBase = m.M
	v1698 = m.ExcPending
	if v1698 != 0 {
		goto L9
	} else {
		goto L503
	}
L47:
	;
	F_ATSimplePermissions(m, int32(20), l1, int32(289))
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L9
	} else {
		goto L501
	}
L48:
	;
	F_ATSimplePermissions(m, int32(19), l1, int32(257))
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L9
	} else {
		goto L499
	}
L49:
	;
	F_ATSimplePermissions(m, int32(50), l1, int32(289))
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L9
	} else {
		goto L498
	}
L50:
	;
	F_ATSimplePermissions(m, int32(49), l1, int32(289))
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		goto L9
	} else {
		goto L477
	}
L51:
	;
	F_ATSimplePermissions(m, v192, l1, int32(271))
	mBase = m.M
	v1612 = m.ExcPending
	if v1612 != 0 {
		goto L9
	} else {
		goto L476
	}
L52:
	;
	F_ATSimplePermissions(m, int32(33), l1, int32(333))
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L9
	} else {
		goto L461
	}
L53:
	;
	F_ATSimplePermissions(m, int32(32), l1, int32(261))
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L9
	} else {
		goto L450
	}
L54:
	;
	F_ATSimplePermissions(m, int32(31), l1, int32(289))
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L9
	} else {
		goto L449
	}
L55:
	;
	F_ATSimplePermissions(m, v192, l1, int32(129))
	mBase = m.M
	v1489 = m.ExcPending
	if v1489 != 0 {
		goto L9
	} else {
		goto L432
	}
L56:
	;
	F_ATSimplePermissions(m, v192, l1, int32(261))
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L9
	} else {
		goto L431
	}
L57:
	;
	F_ATSimplePermissions(m, int32(25), l1, int32(32))
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L9
	} else {
		goto L430
	}
L58:
	;
	F_ATSimplePermissions(m, int32(24), l1, int32(305))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L9
	} else {
		goto L222
	}
L59:
	;
	F_ATSimplePermissions(m, int32(22), l1, int32(289))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L9
	} else {
		goto L196
	}
L60:
	;
	F_ATSimplePermissions(m, int32(21), l1, int32(257))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L9
	} else {
		goto L195
	}
L61:
	;
	F_ATSimplePermissions(m, int32(16), l1, int32(289))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L9
	} else {
		goto L153
	}
L62:
	;
	F_ATSimplePermissions(m, int32(14), l1, int32(257))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L9
	} else {
		goto L152
	}
L63:
	;
	F_ATSimplePermissions(m, int32(13), l1, int32(305))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L9
	} else {
		goto L134
	}
L64:
	;
	F_ATSimplePermissions(m, int32(12), l1, int32(261))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L9
	} else {
		goto L133
	}
L65:
	;
	F_ATSimplePermissions(m, int32(11), l1, int32(293))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L9
	} else {
		goto L131
	}
L66:
	;
	F_ATSimplePermissions(m, v192, l1, int32(293))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L9
	} else {
		goto L130
	}
L67:
	;
	F_ATSimplePermissions(m, int32(8), l1, int32(365))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L9
	} else {
		goto L128
	}
L68:
	;
	F_ATSimplePermissions(m, int32(7), l1, int32(289))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L9
	} else {
		goto L99
	}
L69:
	;
	F_ATSimplePermissions(m, int32(6), l1, int32(289))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L9
	} else {
		goto L97
	}
L70:
	;
	F_ATSimplePermissions(m, int32(5), l1, int32(289))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L9
	} else {
		goto L95
	}
L71:
	;
	F_ATSimplePermissions(m, int32(4), l1, int32(289))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L9
	} else {
		goto L93
	}
L72:
	;
	F_ATSimplePermissions(m, int32(64), l1, int32(291))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L9
	} else {
		goto L91
	}
L73:
	;
	F_ATSimplePermissions(m, int32(63), l1, int32(291))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L9
	} else {
		goto L89
	}
L74:
	;
	F_ATSimplePermissions(m, int32(62), l1, int32(291))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L9
	} else {
		goto L87
	}
L75:
	;
	F_ATSimplePermissions(m, int32(3), l1, int32(289))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L9
	} else {
		goto L86
	}
L76:
	;
	F_ATSimplePermissions(m, int32(2), l1, int32(291))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L9
	} else {
		goto L81
	}
L77:
	;
	v193 = int32(2)
	F_ATSimplePermissions(m, int32(1), l1, v193)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L9
	} else {
		goto L79
	}
L78:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	switch v192 {
	case 0:
		goto L38
	case 1:
		goto L77
	case 2:
		goto L76
	case 3:
		goto L75
	case 4:
		goto L71
	case 5:
		goto L70
	case 6:
		goto L69
	case 7:
		goto L68
	case 8:
		goto L67
	case 9, 10:
		goto L66
	case 11:
		goto L65
	case 12:
		goto L64
	case 13:
		goto L63
	case 14:
		goto L62
	default:
		goto L39
	case 16:
		goto L61
	case 19:
		goto L48
	case 20:
		goto L47
	case 21:
		goto L60
	case 22:
		goto L59
	case 24:
		goto L58
	case 25:
		goto L57
	case 26:
		v1985 = v189
		v1987 = v190
		goto L14
	case 27, 28:
		goto L56
	case 29, 30:
		goto L55
	case 31:
		goto L54
	case 32:
		goto L53
	case 33:
		goto L52
	case 34, 35, 36:
		goto L51
	case 37, 38, 39, 40, 41, 42, 43, 44:
		goto L45
	case 45, 46, 47, 48, 51, 52, 54, 55, 56, 57:
		goto L44
	case 49:
		goto L50
	case 50:
		goto L49
	case 53:
		goto L46
	case 58:
		goto L43
	case 59:
		goto L42
	case 60:
		goto L41
	case 61:
		goto L40
	case 62:
		goto L74
	case 63:
		goto L73
	case 64:
		goto L72
	}
L79:
	;
	F_ATPrepAddColumn(m, l0, l1, l3, l4, int32(1), v190, l5, l6)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	v1985 = v193
	v1987 = v190
	goto L14
L81:
	;
	F_ATSimpleRecursion(m, l0, l1, v190, l3, l5, l6)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L9
	} else {
		goto L82
	}
L82:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v190)+20))
	if v209 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v210 = int32(10)
	goto L85
L84:
	;
	v210 = int32(0)
	goto L85
L85:
	;
	v1985 = v210
	v1987 = v190
	goto L14
L86:
	;
	v1985 = int32(10)
	v1987 = v190
	goto L14
L87:
	;
	v220 = int32(10)
	if l3 == int32(0) {
		v1985 = v220
		v1987 = v190
		goto L14
	} else {
		goto L88
	}
L88:
	;
	v223 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+29)) = uint8(v223)
	v1985 = v220
	v1987 = v190
	goto L14
L89:
	;
	if l3 == int32(0) {
		v1985 = v189
		v1987 = v190
		goto L14
	} else {
		goto L90
	}
L90:
	;
	v231 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+29)) = uint8(v231)
	v1985 = v189
	v1987 = v190
	goto L14
L91:
	;
	v237 = int32(0)
	if l3 == v237 {
		v1985 = v237
		v1987 = v190
		goto L14
	} else {
		goto L92
	}
L92:
	;
	v240 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+29)) = uint8(v240)
	v1985 = v237
	v1987 = v190
	goto L14
L93:
	;
	v246 = int32(0)
	if l3 == v246 {
		v1985 = v246
		v1987 = v190
		goto L14
	} else {
		goto L94
	}
L94:
	;
	v249 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+29)) = uint8(v249)
	v1985 = v246
	v1987 = v190
	goto L14
L95:
	;
	v255 = int32(7)
	if l3 == int32(0) {
		v1985 = v255
		v1987 = v190
		goto L14
	} else {
		goto L96
	}
L96:
	;
	v258 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+29)) = uint8(v258)
	v1985 = v255
	v1987 = v190
	goto L14
L97:
	;
	F_ATSimpleRecursion(m, l0, l1, v190, l3, l5, l6)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L9
	} else {
		goto L98
	}
L98:
	;
	v1985 = int32(3)
	v1987 = v190
	goto L14
L99:
	;
	F_ATSimpleRecursion(m, l0, l1, v190, l3, l5, l6)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L9
	} else {
		goto L100
	}
L100:
	;
	v273 = m.G0
	v275 = v273 - int32(16)
	m.G0 = v275
	if l3 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L101:
	;
	v1985 = int32(0)
	v1987 = v190
	goto L14
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L9
	} else {
		goto L124
	}
L103:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L9
	} else {
		goto L120
	}
L104:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L9
	} else {
		goto L116
	}
L105:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v280 = F_find_inheritance_children(m, v279, l5)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L9
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if l4 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	if v280 != 0 {
		goto L104
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v190)+8))
	v286 = F_SearchSysCacheCopyAttName(m, v284, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L9
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	m.G0 = v275 + int32(16)
	goto L101
L113:
	;
	if v286 == int32(0) {
		goto L103
	} else {
		goto L114
	}
L114:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v286)+16))
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290)+22)))
	v293 = int32(*(*int16)(unsafe.Add(mBase, uint32(v290+v291)+94)))
	if int32(0) < v293 {
		goto L102
	} else {
		goto L115
	}
L115:
	;
	goto L112
L116:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L9
	} else {
		goto L117
	}
L117:
	;
	F_errmsg(m, int32(236982), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L9
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(486904), int32(8772), int32(266830))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L9
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L9
	} else {
		goto L121
	}
L121:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v190)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v275))) = v324
	*(*int32)(unsafe.Add(mBase, uint32(v275)+4)) = v323 + int32(4)
	F_errmsg(m, int32(70796), v275)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L9
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(486904), int32(8787), int32(266830))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L9
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L9
	} else {
		goto L125
	}
L125:
	;
	F_errmsg(m, int32(270305), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L9
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(486904), int32(8794), int32(266830))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L9
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	F_ATSimpleRecursion(m, l0, l1, v190, l3, l5, l6)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L9
	} else {
		goto L129
	}
L129:
	;
	v1985 = v189
	v1987 = v190
	goto L14
L130:
	;
	v1985 = v189
	v1987 = v190
	goto L14
L131:
	;
	F_ATSimpleRecursion(m, l0, l1, v190, l3, l5, l6)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L9
	} else {
		goto L132
	}
L132:
	;
	v1985 = v189
	v1987 = v190
	goto L14
L133:
	;
	v1985 = v189
	v1987 = v190
	goto L14
L134:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if l4 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v1985 = int32(0)
	v1987 = v190
	goto L14
L136:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L9
	} else {
		goto L148
	}
L137:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v377)+76))
	if v380 != 0 {
		goto L136
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+119)))
	if v381 == int32(99) {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	goto L139
L141:
	;
	F_ATTypedTableRecursion(m, l0, l1, v190, l5, l6)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L9
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	if l3 != 0 {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	goto L143
L145:
	;
	v386 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+29)) = uint8(v386)
	goto L147
L146:
	;
	goto L147
L147:
	;
	goto L135
L148:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L9
	} else {
		goto L149
	}
L149:
	;
	F_errmsg(m, int32(388399), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L9
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(486904), int32(9263), int32(271033))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L9
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	v1985 = int32(9)
	v1987 = v190
	goto L14
L153:
	;
	v414 = int32(0)
	v416 = m.G0
	v418 = v416 - int32(16)
	m.G0 = v418
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v190)+20))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	if v421 != int32(6) {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	v617 = int32(6)
	if l3 == int32(0) {
		v1985 = v617
		v1987 = v190
		goto L14
	} else {
		goto L193
	}
L155:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L9
	} else {
		goto L189
	}
L156:
	;
	m.G0 = v418 + int32(16)
	goto L154
L157:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v420)+32))
	if v424 == int32(0) {
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v424)+4))
	if v427 <= int32(0) {
		goto L156
	} else {
		goto L159
	}
L159:
	;
	v437 = v414
	v439 = v414
	v443 = v8
	goto L160
L160:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v424)+12))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v451+v443<<(uint(int32(2))%32))))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)+4))
	v457 = F_findNotNullConstraint(m, v450, v456)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L9
	} else {
		goto L163
	}
L161:
	;
	goto L156
L162:
	;
	v574 = v443 + int32(1)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v424)+4))
	if v574 < v575 {
		v437 = v560
		v439 = v562
		v443 = v574
		goto L160
	} else {
		goto L188
	}
L163:
	;
	if v457 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v455)+4))
	F_verifyNotNullPKCompatible(m, v457, v459)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L9
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	if l3 != 0 {
		v526 = v437
		v528 = v439
		goto L169
	} else {
		goto L170
	}
L167:
	;
	F_pfree(m, v457)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L9
	} else {
		goto L168
	}
L168:
	;
	v560 = v437
	v562 = v439
	goto L162
L169:
	;
	v539 = F_makeNotNullConstraint(m, v455)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L9
	} else {
		goto L185
	}
L170:
	;
	if v439&int32(1) == int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v469 = F_find_inheritance_children(m, v468, l5)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L9
	} else {
		goto L174
	}
L172:
	;
	v471 = v437
	goto L173
L173:
	;
	if v471 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v471 = v469
	goto L173
L175:
	;
	v526 = int32(0)
	v528 = int32(1)
	goto L169
L176:
	;
	goto L177
L177:
	;
	v477 = int32(0)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v471)+4))
	if v478 <= v477 {
		v526 = v471
		v528 = int32(1)
		goto L169
	} else {
		goto L178
	}
L178:
	;
	v481 = v477
	goto L179
L179:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v471)+12))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v501+v481<<(uint(int32(2))%32))))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v455)+4))
	v507 = F_findNotNullConstraint(m, v505, v506)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L9
	} else {
		goto L181
	}
L180:
	;
	v526 = v471
	v528 = v514
	goto L169
L181:
	;
	if v507 == int32(0) {
		goto L155
	} else {
		goto L182
	}
L182:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v455)+4))
	F_verifyNotNullPKCompatible(m, v507, v511)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L9
	} else {
		goto L183
	}
L183:
	;
	v514 = int32(1)
	v516 = v481 + v514
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v471)+4))
	if v516 < v517 {
		v481 = v516
		goto L179
	} else {
		goto L184
	}
L184:
	;
	goto L180
L185:
	;
	v542 = F_palloc0(m, int32(32))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L9
	} else {
		goto L186
	}
L186:
	;
	v544 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v542)+29)) = uint8(v544)
	*(*int64)(unsafe.Add(mBase, uint32(v542))) = int64(68719476883)
	*(*int32)(unsafe.Add(mBase, uint32(v542)+20)) = v539
	F_ATPrepCmd(m, l0, l1, v542, v544, int32(0), l5, l6)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L9
	} else {
		goto L187
	}
L187:
	;
	v560 = v526
	v562 = v528
	goto L162
L188:
	;
	goto L161
L189:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v455)+4))
	v605 = F_get_rel_name(m, v505)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L9
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v418)+4)) = v605
	*(*int32)(unsafe.Add(mBase, uint32(v418))) = v604
	F_errmsg(m, int32(524646), v418)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L9
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(486904), int32(9554), int32(22385))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L9
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v622 = F_find_all_inheritors(m, v620, l5, int32(0))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L9
	} else {
		goto L194
	}
L194:
	;
	v624 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+29)) = uint8(v624)
	v1985 = v617
	v1987 = v190
	goto L14
L195:
	;
	v1985 = int32(8)
	v1987 = v190
	goto L14
L196:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v635)+119)))
	if v636 == int32(112) {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	v752 = int32(0)
	if l3 == v752 {
		v1985 = v752
		v1987 = v190
		goto L14
	} else {
		goto L221
	}
L198:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L9
	} else {
		goto L217
	}
L199:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v641 = F_find_all_inheritors(m, v639, l5, int32(0))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L9
	} else {
		goto L203
	}
L200:
	;
	goto L201
L201:
	;
	goto L197
L202:
	;
	F_list_free(m, v641)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L9
	} else {
		goto L216
	}
L203:
	;
	if v641 == int32(0) {
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v645 = int32(1)
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v641)+4))
	if v646 <= v645 {
		goto L202
	} else {
		goto L205
	}
L205:
	;
	v654 = v645
	goto L206
L206:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v641)+12))
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v669+v654<<(uint(int32(2))%32))))
	v675 = F_table_open(m, v673, int32(0))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L9
	} else {
		goto L208
	}
L207:
	;
	goto L202
L208:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v675)+48))
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677)+118)))
	if v678 == int32(116) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675)+24)))
	if v681 == int32(0) {
		goto L198
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	F_CheckTableNotInUse(m, v675, int32(532528))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L9
	} else {
		goto L213
	}
L212:
	;
	goto L211
L213:
	;
	F_sequence_close(m, v675, int32(0))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L9
	} else {
		goto L214
	}
L214:
	;
	v691 = v654 + int32(1)
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v641)+4))
	if v691 < v692 {
		v654 = v691
		goto L206
	} else {
		goto L215
	}
L215:
	;
	goto L207
L216:
	;
	goto L201
L217:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L9
	} else {
		goto L218
	}
L218:
	;
	F_errmsg(m, int32(141730), int32(0))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L9
	} else {
		goto L219
	}
L219:
	;
	F_errfinish(m, int32(486904), int32(4460), int32(403669))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L9
	} else {
		goto L220
	}
L220:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L221:
	;
	v755 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+29)) = uint8(v755)
	v1985 = v752
	v1987 = v190
	goto L14
L222:
	;
	v762 = F_ATParseTransformCmd(m, v118, l1, v190, l3, int32(-1), l6)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L9
	} else {
		goto L223
	}
L223:
	;
	v764 = m.G0
	v766 = v764 - int32(208)
	m.G0 = v766
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v762)+8))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v762)+20))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v769)+32))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v769)+8))
	v773 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L9
	} else {
		goto L224
	}
L224:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v773)+4)) = v775
	if l4 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L225:
	;
	v1985 = int32(1)
	v1987 = v762
	goto L14
L226:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L9
	} else {
		goto L426
	}
L227:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L9
	} else {
		goto L421
	}
L228:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L9
	} else {
		goto L417
	}
L229:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L9
	} else {
		goto L413
	}
L230:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L9
	} else {
		goto L409
	}
L231:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L9
	} else {
		goto L405
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v766)+148)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v766)+144)) = v768
	F_errmsg(m, int32(187830), v766+int32(144))
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L9
	} else {
		goto L402
	}
L233:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L9
	} else {
		goto L397
	}
L234:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L9
	} else {
		goto L392
	}
L235:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L9
	} else {
		goto L386
	}
L236:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L9
	} else {
		goto L381
	}
L237:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L9
	} else {
		goto L376
	}
L238:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L9
	} else {
		goto L371
	}
L239:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v779)+76))
	if v780 != 0 {
		goto L238
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v782 = F_SearchSysCacheAttName(m, v781, v768)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L9
	} else {
		goto L243
	}
L242:
	;
	goto L241
L243:
	;
	if v782 == int32(0) {
		goto L237
	} else {
		goto L244
	}
L244:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v782)+16))
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786)+22)))
	v788 = v786 + v787
	v789 = int32(*(*int16)(unsafe.Add(mBase, uint32(v788)+74)))
	if v789 <= int32(0) {
		goto L236
	} else {
		goto L245
	}
L245:
	;
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v788)+90)))
	if v792 != 0 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v769)+32))
	if v793 != 0 {
		goto L235
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	if l4 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L249:
	;
	goto L248
L250:
	;
	v796 = int32(*(*int16)(unsafe.Add(mBase, uint32(v788)+94)))
	if int32(0) < v796 {
		goto L234
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	v801 = F_bms_make_singleton(m, v789+int32(7))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L9
	} else {
		goto L254
	}
L253:
	;
	goto L252
L254:
	;
	v805 = F_has_partition_attrs(m, l1, v801, v766+int32(199))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L9
	} else {
		goto L255
	}
L255:
	;
	if v805 != 0 {
		goto L233
	} else {
		goto L256
	}
L256:
	;
	F_typenameTypeIdAndMod(m, v773, v771, v766+int32(204), v766+int32(200))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L9
	} else {
		goto L257
	}
L257:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v766)+204))
	v816 = *(*int32)(unsafe.Add(mBase, _consts[159]))
	v818 = F_object_aclcheck(m, int32(1247), v814, v816, int64(256))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L9
	} else {
		goto L258
	}
L258:
	;
	if v818 != 0 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v766)+204))
	F_aclcheck_error_type(m, v818, v820)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L9
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v766)+204))
	v824 = F_GetColumnDefCollation(m, v773, v769, v823)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L9
	} else {
		goto L263
	}
L262:
	;
	goto L261
L263:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v826)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v766)+192)) = v827
	*(*int32)(unsafe.Add(mBase, uint32(v766)+156)) = v827
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v766)+204))
	v834 = F_list_make1_impl(m, int32(472), v766+int32(156))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L9
	} else {
		goto L264
	}
L264:
	;
	v836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v788)+90)))
	F_CheckAttributeType(m, v768, v830, v824, v834, base.B2i32(v836 == int32(118))<<(uint(int32(3))%32))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L9
	} else {
		goto L265
	}
L265:
	;
	v843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v788)+90)))
	if v843 == int32(118) {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)))
	switch v1049 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L321
	default:
		goto L320
	}
L267:
	;
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)))
	switch v846 - int32(112) {
	case 0, 2:
		goto L269
	default:
		goto L268
	}
L268:
	;
	if v770 != 0 {
		goto L231
	} else {
		goto L318
	}
L269:
	;
	if v770 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v788)+68))
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v788)+76))
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v788)+96))
	v856 = F_makeVar(m, int32(1), v789, v852, v853, v854, int32(0))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L9
	} else {
		goto L273
	}
L271:
	;
	v858 = v770
	goto L272
L272:
	;
	v859 = F_exprType(m, v858)
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L9
	} else {
		goto L274
	}
L273:
	;
	v858 = v856
	goto L272
L274:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v766)+204))
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v766)+200))
	v866 = F_coerce_to_target_type(m, v773, v858, v859, v861, v862, int32(1), int32(2), int32(-1))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L9
	} else {
		goto L275
	}
L275:
	;
	if v866 == int32(0) {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v769)+32))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L9
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	F_assign_expr_collations(m, v773, v866)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L9
	} else {
		goto L291
	}
L279:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L9
	} else {
		goto L280
	}
L280:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v766)+204))
	v879 = F_format_type_be(m, v878)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L9
	} else {
		goto L281
	}
L281:
	;
	if v870 != 0 {
		goto L232
	} else {
		goto L282
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v766)+132)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v766)+128)) = v768
	F_errmsg(m, int32(187857), v766+int32(128))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L9
	} else {
		goto L283
	}
L283:
	;
	v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v788)+90)))
	if v888 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v891 = F_quote_identifier(m, v768)
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L9
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	F_errfinish(m, int32(486904), int32(14517), int32(365820))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L9
	} else {
		goto L290
	}
L287:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v766)+204))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v766)+200))
	v895 = F_format_type_with_typemod(m, v893, v894)
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L9
	} else {
		goto L288
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v766)+116)) = v895
	*(*int32)(unsafe.Add(mBase, uint32(v766)+112)) = v891
	F_errhint(m, int32(640549), v766+int32(112))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L9
	} else {
		goto L289
	}
L289:
	;
	goto L286
L290:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L291:
	;
	v913 = F_expand_generated_columns_in_expr(m, v866, l1, int32(1))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L9
	} else {
		goto L292
	}
L292:
	;
	v915 = F_expression_planner(m, v913)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L9
	} else {
		goto L293
	}
L293:
	;
	v918 = F_palloc0(m, int32(16))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L9
	} else {
		goto L294
	}
L294:
	;
	v920 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v918)+12)) = uint8(v920)
	*(*int32)(unsafe.Add(mBase, uint32(v918)+4)) = v915
	*(*uint16)(unsafe.Add(mBase, uint32(v918))) = uint16(v789)
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v118)+68))
	v925 = F_lappend(m, v924, v918)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L9
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+68)) = v925
	v936 = v915
	goto L297
L296:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v118)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+80)) = v1025 | int32(4)
	goto L266
L297:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v936)))
	switch v948 - int32(6) {
	case 0:
		goto L299
	case 1, 2, 3, 4, 5, 6, 7, 8, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20:
		goto L296
	case 9:
		goto L300
	case 21:
		goto L302
	default:
		goto L301
	}
L298:
	;
	v1018 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v936)+8)))
	if v1018 == v789&int32(65535) {
		goto L266
	} else {
		goto L317
	}
L299:
	;
	goto L298
L300:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v936)+4))
	if base.Ui32(int32(1)) < base.Ui32(v958-int32(2027)) {
		goto L296
	} else {
		goto L306
	}
L301:
	;
	if v948 != int32(55) {
		goto L296
	} else {
		goto L303
	}
L302:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v936)+4))
	v936 = v951
	goto L297
L303:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v936)+8))
	v955 = F_DomainHasConstraints(m, v954)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L9
	} else {
		goto L304
	}
L304:
	;
	if v955 != 0 {
		goto L296
	} else {
		goto L305
	}
L305:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v936)+4))
	v936 = v957
	goto L297
L306:
	;
	v963 = m.G0
	v965 = v963 - int32(16)
	m.G0 = v965
	v968 = *(*int32)(unsafe.Add(mBase, _consts[450]))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v968)+uint32(_consts[451])))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v968)+264))
	if v975 < int32(2) {
		goto L308
	} else {
		goto L309
	}
L307:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v965)+12))
	m.G0 = v965 + int32(16)
	if v1008|(v1007^int32(1)) != 0 {
		goto L296
	} else {
		goto L316
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v965+int32(12)))) = v974
	v1007 = int32(1)
	goto L307
L309:
	;
	v981 = int32(1)
	goto L310
L310:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v968+int32(18280)+v981<<(uint(int32(4))%32))))
	if v974 == v989 {
		goto L312
	} else {
		goto L313
	}
L311:
	;
	v1007 = int32(0)
	goto L307
L312:
	;
	v992 = v981 + int32(1)
	if v975 != v992 {
		v981 = v992
		goto L310
	} else {
		goto L315
	}
L313:
	;
	goto L314
L314:
	;
	goto L311
L315:
	;
	goto L308
L316:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v936)+28))
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v1015)+12))
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1016)))
	v936 = v1017
	goto L297
L317:
	;
	goto L296
L318:
	;
	goto L266
L319:
	;
	F_ReleaseCatCache(m, v782)
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L9
	} else {
		goto L324
	}
L320:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1055)+72))
	F_find_composite_type_dependencies(m, v1056, l1, int32(0))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L9
	} else {
		goto L323
	}
L321:
	;
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v788)+90)))
	if v1052 != int32(118) {
		goto L319
	} else {
		goto L322
	}
L322:
	;
	goto L320
L323:
	;
	goto L319
L324:
	;
	if l3 != 0 {
		goto L326
	} else {
		goto L327
	}
L325:
	;
	v1202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)))
	if v1202 == int32(99) {
		goto L367
	} else {
		goto L368
	}
L326:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v1065 = F_find_all_inheritors(m, v1062, l5, v766+int32(188))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L9
	} else {
		goto L329
	}
L327:
	;
	goto L328
L328:
	;
	if l4 != 0 {
		v1184 = v762
		goto L325
	} else {
		goto L364
	}
L329:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v766)+188))
	v1069 = v762
	v1077 = int32(0)
	goto L330
L330:
	;
	v1089 = int32(0)
	if v1065 == v1089 {
		v1099 = v1089
		goto L332
	} else {
		goto L333
	}
L332:
	;
	if v1067 == int32(0) {
		v1184 = v762
		goto L325
	} else {
		goto L335
	}
L333:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1065)+4))
	if v1093 <= v1077 {
		v1099 = int32(0)
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1065)+12))
	v1099 = v1095 + v1077<<(uint(int32(2))%32)
	goto L332
L335:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+4))
	if v1102 <= v1077 {
		goto L337
	} else {
		goto L338
	}
L336:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1099)))
	if v1062 != v1111 {
		goto L341
	} else {
		goto L342
	}
L337:
	;
	v1184 = v1069
	goto L325
L338:
	;
	if v1099 == int32(0) {
		goto L337
	} else {
		goto L339
	}
L339:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+12))
	v1109 = v1106 + v1077<<(uint(int32(2))%32)
	if v1109 != 0 {
		goto L336
	} else {
		goto L340
	}
L340:
	;
	goto L337
L341:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1109)))
	v1115 = F_relation_open(m, v1111, int32(0))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L9
	} else {
		goto L344
	}
L342:
	;
	v1171 = v1069
	goto L343
L343:
	;
	v1069 = v1171
	v1077 = v1077 + int32(1)
	goto L330
L344:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+48))
	v1118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1117)+118)))
	if v1118 == int32(116) {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1115)+24)))
	if v1121 == int32(0) {
		goto L230
	} else {
		goto L348
	}
L346:
	;
	goto L347
L347:
	;
	F_CheckTableNotInUse(m, v1115, int32(532528))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L9
	} else {
		goto L349
	}
L348:
	;
	goto L347
L349:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+56))
	v1128 = F_SearchSysCacheAttName(m, v1127, v768)
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L9
	} else {
		goto L350
	}
L350:
	;
	if v1128 == int32(0) {
		goto L229
	} else {
		goto L351
	}
L351:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+16))
	v1133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1132)+22)))
	v1135 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1132+v1133)+94)))
	if v1113 < v1135 {
		goto L228
	} else {
		goto L352
	}
L352:
	;
	F_ReleaseCatCache(m, v1128)
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L9
	} else {
		goto L353
	}
L353:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v769)+32))
	if v1139 != 0 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v1140 = F_copyObjectImpl(m, v1069)
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L9
	} else {
		goto L357
	}
L355:
	;
	v1161 = v1069
	goto L356
L356:
	;
	F_ATPrepCmd(m, l0, v1115, v1161, int32(0), int32(1), l5, l6)
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L9
	} else {
		goto L362
	}
L357:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+52))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v1145 = F_build_attrmap_by_name(m, v1142, v1143, int32(0))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L9
	} else {
		goto L358
	}
L358:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v769)+32))
	v1152 = F_map_variable_attnos(m, v1147, int32(1), v1145, int32(0), v766+int32(187))
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L9
	} else {
		goto L359
	}
L359:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1140)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1154)+32)) = v1152
	v1156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v766)+187)))
	if v1156 == int32(1) {
		goto L227
	} else {
		goto L360
	}
L360:
	;
	F_pfree(m, v1145)
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L9
	} else {
		goto L361
	}
L361:
	;
	v1161 = v1140
	goto L356
L362:
	;
	F_relation_close(m, v1115, int32(0))
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L9
	} else {
		goto L363
	}
L363:
	;
	v1171 = v1161
	goto L343
L364:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v1180 = F_find_inheritance_children(m, v1178, int32(0))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L9
	} else {
		goto L365
	}
L365:
	;
	if v1180 != 0 {
		goto L226
	} else {
		goto L366
	}
L366:
	;
	v1184 = v762
	goto L325
L367:
	;
	F_ATTypedTableRecursion(m, l0, l1, v1184, l5, l6)
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L9
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	m.G0 = v766 + int32(208)
	goto L225
L370:
	;
	goto L369
L371:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L9
	} else {
		goto L372
	}
L372:
	;
	F_errmsg(m, int32(388471), int32(0))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L9
	} else {
		goto L373
	}
L373:
	;
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v769)+64))
	F_parser_errposition(m, v773, v1221)
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L9
	} else {
		goto L374
	}
L374:
	;
	F_errfinish(m, int32(486904), int32(14400), int32(365820))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L9
	} else {
		goto L375
	}
L375:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L376:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L9
	} else {
		goto L377
	}
L377:
	;
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v766))) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v766)+4)) = v1236 + int32(4)
	F_errmsg(m, int32(70796), v766)
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L9
	} else {
		goto L378
	}
L378:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v769)+64))
	F_parser_errposition(m, v773, v1244)
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L9
	} else {
		goto L379
	}
L379:
	;
	F_errfinish(m, int32(486904), int32(14409), int32(365820))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L9
	} else {
		goto L380
	}
L380:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L381:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L9
	} else {
		goto L382
	}
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v766)+16)) = v768
	F_errmsg(m, int32(687318), v766+int32(16))
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L9
	} else {
		goto L383
	}
L383:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v769)+64))
	F_parser_errposition(m, v773, v1265)
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L9
	} else {
		goto L384
	}
L384:
	;
	F_errfinish(m, int32(486904), int32(14418), int32(365820))
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L9
	} else {
		goto L385
	}
L385:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L386:
	;
	F_errcode(m, int32(17064068))
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L9
	} else {
		goto L387
	}
L387:
	;
	F_errmsg(m, int32(270472), int32(0))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L9
	} else {
		goto L388
	}
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v766)+176)) = v768
	F_errdetail(m, int32(595654), v766+int32(176))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L9
	} else {
		goto L389
	}
L389:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v769)+64))
	F_parser_errposition(m, v773, v1290)
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L9
	} else {
		goto L390
	}
L390:
	;
	F_errfinish(m, int32(486904), int32(14429), int32(365820))
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L9
	} else {
		goto L391
	}
L391:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L392:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L9
	} else {
		goto L393
	}
L393:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v766)+160)) = v768
	F_errmsg(m, int32(687831), v766+int32(160))
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L9
	} else {
		goto L394
	}
L394:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v769)+64))
	F_parser_errposition(m, v773, v1311)
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L9
	} else {
		goto L395
	}
L395:
	;
	F_errfinish(m, int32(486904), int32(14440), int32(365820))
	mBase = m.M
	v1318 = m.ExcPending
	if v1318 != 0 {
		goto L9
	} else {
		goto L396
	}
L396:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L397:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L9
	} else {
		goto L398
	}
L398:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v766)+32)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v766)+36)) = v1326 + int32(4)
	F_errmsg(m, int32(683220), v766+int32(32))
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L9
	} else {
		goto L399
	}
L399:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v769)+64))
	F_parser_errposition(m, v773, v1336)
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L9
	} else {
		goto L400
	}
L400:
	;
	F_errfinish(m, int32(486904), int32(14450), int32(365820))
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L9
	} else {
		goto L401
	}
L401:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L402:
	;
	F_errhint(m, int32(555103), int32(0))
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L9
	} else {
		goto L403
	}
L403:
	;
	F_errfinish(m, int32(486904), int32(14506), int32(365820))
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L9
	} else {
		goto L404
	}
L404:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L405:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L9
	} else {
		goto L406
	}
L406:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v766)+96)) = v1367 + int32(4)
	F_errmsg(m, int32(389268), v766+int32(96))
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L9
	} else {
		goto L407
	}
L407:
	;
	F_errfinish(m, int32(486904), int32(14546), int32(365820))
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L9
	} else {
		goto L408
	}
L408:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L409:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L9
	} else {
		goto L410
	}
L410:
	;
	F_errmsg(m, int32(141730), int32(0))
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L9
	} else {
		goto L411
	}
L411:
	;
	F_errfinish(m, int32(486904), int32(4460), int32(403669))
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L9
	} else {
		goto L412
	}
L412:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L413:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L9
	} else {
		goto L414
	}
L414:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v766)+48)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v766)+52)) = v1404 + int32(4)
	F_errmsg(m, int32(70796), v766+int32(48))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L9
	} else {
		goto L415
	}
L415:
	;
	F_errfinish(m, int32(486904), int32(14611), int32(365820))
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
		goto L9
	} else {
		goto L416
	}
L416:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L417:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L9
	} else {
		goto L418
	}
L418:
	;
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v766)+64)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v766)+68)) = v1426 + int32(4)
	F_errmsg(m, int32(684761), v766-int32(-64))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L9
	} else {
		goto L419
	}
L419:
	;
	F_errfinish(m, int32(486904), int32(14618), int32(365820))
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L9
	} else {
		goto L420
	}
L420:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L421:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L9
	} else {
		goto L422
	}
L422:
	;
	F_errmsg(m, int32(410698), int32(0))
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L9
	} else {
		goto L423
	}
L423:
	;
	F_errdetail(m, int32(620294), int32(0))
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L9
	} else {
		goto L424
	}
L424:
	;
	F_errfinish(m, int32(486904), int32(14646), int32(365820))
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L9
	} else {
		goto L425
	}
L425:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L426:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L9
	} else {
		goto L427
	}
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v766)+80)) = v768
	F_errmsg(m, int32(237350), v766+int32(80))
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L9
	} else {
		goto L428
	}
L428:
	;
	F_errfinish(m, int32(486904), int32(14658), int32(365820))
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
		goto L9
	} else {
		goto L429
	}
L429:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L430:
	;
	v1985 = v189
	v1987 = v190
	goto L14
L431:
	;
	v1985 = v189
	v1987 = v190
	goto L14
L432:
	;
	v1490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+96)))
	if v1490 == int32(1) {
		goto L19
	} else {
		goto L433
	}
L433:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1494)+118)))
	switch v1495 - int32(112) {
	case 0:
		goto L435
	default:
		goto L436
	case 4:
		goto L438
	case 5:
		goto L437
	}
L434:
	;
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v1528 = F_GetRelationPublications(m, v1527)
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L9
	} else {
		goto L447
	}
L435:
	;
	if v1493 == int32(29) {
		v1985 = v189
		v1987 = v190
		goto L14
	} else {
		goto L446
	}
L436:
	;
	if v1493 == int32(29) {
		goto L16
	} else {
		goto L445
	}
L437:
	;
	if v1493 != int32(29) {
		v1985 = v189
		v1987 = v190
		goto L14
	} else {
		goto L444
	}
L438:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1501 = m.ExcPending
	if v1501 != 0 {
		goto L9
	} else {
		goto L439
	}
L439:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L9
	} else {
		goto L440
	}
L440:
	;
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v1505 + int32(4)
	F_errmsg(m, int32(17160), v23-int32(-64))
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L9
	} else {
		goto L441
	}
L441:
	;
	F_errtable(m, l1)
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L9
	} else {
		goto L442
	}
L442:
	;
	F_errfinish(m, int32(486904), int32(18840), int32(410323))
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L9
	} else {
		goto L443
	}
L443:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L444:
	;
	goto L16
L445:
	;
	goto L434
L446:
	;
	goto L434
L447:
	;
	if v1528 != 0 {
		goto L18
	} else {
		goto L448
	}
L448:
	;
	v1831 = int32(0)
	v1832 = int32(13)
	goto L15
L449:
	;
	v1985 = int32(0)
	v1987 = v190
	goto L14
L450:
	;
	v1541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+84)))
	if v1541 == int32(1) {
		goto L17
	} else {
		goto L451
	}
L451:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v190)+8))
	if v1544 != 0 {
		goto L453
	} else {
		goto L454
	}
L452:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(v1557)+84))
	if v1556 != v1558 {
		goto L458
	} else {
		goto L459
	}
L453:
	;
	v1552 = v1544
	goto L455
L454:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1546)+119)))
	if v1547 == int32(112) {
		v1556 = int32(0)
		goto L452
	} else {
		goto L456
	}
L455:
	;
	v1554 = F_get_table_am_oid(m, v1552, int32(0))
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L9
	} else {
		goto L457
	}
L456:
	;
	v1551 = *(*int32)(unsafe.Add(mBase, _consts[452]))
	v1552 = v1551
	goto L455
L457:
	;
	v1556 = v1554
	goto L452
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+88)) = v1556
	v1561 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v118)+84)) = uint8(v1561)
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(v118)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+80)) = v1563 | int32(8)
	goto L460
L459:
	;
	goto L460
L460:
	;
	v1985 = v189
	v1987 = v190
	goto L14
L461:
	;
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v190)+8))
	v1573 = F_get_tablespace_oid(m, v1571, int32(0))
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L9
	} else {
		goto L463
	}
L462:
	;
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(v118)+92))
	if v1592 != 0 {
		goto L469
	} else {
		goto L470
	}
L463:
	;
	if v1573 == int32(0) {
		goto L462
	} else {
		goto L464
	}
L464:
	;
	v1578 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	if v1573 == v1578 {
		goto L462
	} else {
		goto L465
	}
L465:
	;
	v1582 = *(*int32)(unsafe.Add(mBase, _consts[159]))
	v1584 = F_object_aclcheck(m, int32(1213), v1573, v1582, int64(512))
	mBase = m.M
	v1585 = m.ExcPending
	if v1585 != 0 {
		goto L9
	} else {
		goto L466
	}
L466:
	;
	if v1584 == int32(0) {
		goto L462
	} else {
		goto L467
	}
L467:
	;
	F_aclcheck_error(m, v1584, int32(42), v1571)
	mBase = m.M
	v1590 = m.ExcPending
	if v1590 != 0 {
		goto L9
	} else {
		goto L468
	}
L468:
	;
	goto L462
L469:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L9
	} else {
		goto L472
	}
L470:
	;
	goto L471
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+92)) = v1573
	v1985 = v189
	v1987 = v190
	goto L14
L472:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L9
	} else {
		goto L473
	}
L473:
	;
	F_errmsg(m, int32(170414), int32(0))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L9
	} else {
		goto L474
	}
L474:
	;
	F_errfinish(m, int32(486904), int32(16636), int32(413301))
	mBase = m.M
	v1608 = m.ExcPending
	if v1608 != 0 {
		goto L9
	} else {
		goto L475
	}
L475:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L476:
	;
	v1985 = v189
	v1987 = v190
	goto L14
L477:
	;
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v1617)+76))
	if v1618 == int32(0) {
		goto L481
	} else {
		goto L482
	}
L478:
	;
	v1985 = v189
	v1987 = v190
	goto L14
L479:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		goto L9
	} else {
		goto L494
	}
L480:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L9
	} else {
		goto L490
	}
L481:
	;
	v1621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1617)+131)))
	if v1621 == int32(1) {
		goto L480
	} else {
		goto L484
	}
L482:
	;
	goto L483
L483:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1630 = m.ExcPending
	if v1630 != 0 {
		goto L9
	} else {
		goto L486
	}
L484:
	;
	v1624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1617)+119)))
	if v1624 == int32(112) {
		goto L479
	} else {
		goto L485
	}
L485:
	;
	goto L478
L486:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L9
	} else {
		goto L487
	}
L487:
	;
	F_errmsg(m, int32(388511), int32(0))
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L9
	} else {
		goto L488
	}
L488:
	;
	F_errfinish(m, int32(486904), int32(17244), int32(98437))
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L9
	} else {
		goto L489
	}
L489:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L490:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L9
	} else {
		goto L491
	}
L491:
	;
	F_errmsg(m, int32(245658), int32(0))
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		goto L9
	} else {
		goto L492
	}
L492:
	;
	F_errfinish(m, int32(486904), int32(17249), int32(98437))
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L9
	} else {
		goto L493
	}
L493:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L494:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L9
	} else {
		goto L495
	}
L495:
	;
	F_errmsg(m, int32(388727), int32(0))
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
		goto L9
	} else {
		goto L496
	}
L496:
	;
	F_errfinish(m, int32(486904), int32(17254), int32(98437))
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L9
	} else {
		goto L497
	}
L497:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L498:
	;
	v1985 = v189
	v1987 = v190
	goto L14
L499:
	;
	if l3 == int32(0) {
		v1985 = v189
		v1987 = v190
		goto L14
	} else {
		goto L500
	}
L500:
	;
	v1685 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+29)) = uint8(v1685)
	v1985 = v189
	v1987 = v190
	goto L14
L501:
	;
	if l3 == int32(0) {
		v1985 = v189
		v1987 = v190
		goto L14
	} else {
		goto L502
	}
L502:
	;
	v1693 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+29)) = uint8(v1693)
	v1985 = v189
	v1987 = v190
	goto L14
L503:
	;
	v1985 = v189
	v1987 = v190
	goto L14
L504:
	;
	if l3 == int32(0) {
		v1985 = v189
		v1987 = v190
		goto L14
	} else {
		goto L505
	}
L505:
	;
	v1704 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+29)) = uint8(v1704)
	v1985 = v189
	v1987 = v190
	goto L14
L506:
	;
	v1985 = v189
	v1987 = v190
	goto L14
L507:
	;
	v1985 = v189
	v1987 = v190
	goto L14
L508:
	;
	v1985 = v189
	v1987 = v190
	goto L14
L509:
	;
	v1985 = v189
	v1987 = v190
	goto L14
L510:
	;
	v1985 = v189
	v1987 = v190
	goto L14
L511:
	;
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v1729
	F_errmsg_internal(m, int32(478227), v23)
	mBase = m.M
	v1733 = m.ExcPending
	if v1733 != 0 {
		goto L9
	} else {
		goto L512
	}
L512:
	;
	F_errfinish(m, int32(486904), int32(5284), int32(423023))
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L9
	} else {
		goto L513
	}
L513:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L514:
	;
	F_ATPrepAddColumn(m, l0, l1, l3, l4, int32(0), v190, l5, l6)
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L9
	} else {
		goto L515
	}
L515:
	;
	v1985 = int32(2)
	v1987 = v190
	goto L14
L516:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L9
	} else {
		goto L517
	}
L517:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v1754 + int32(4)
	F_errmsg(m, int32(321424), v23+int32(80))
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
		goto L9
	} else {
		goto L518
	}
L518:
	;
	F_errhint(m, int32(592402), int32(0))
	mBase = m.M
	v1766 = m.ExcPending
	if v1766 != 0 {
		goto L9
	} else {
		goto L519
	}
L519:
	;
	F_errfinish(m, int32(486904), int32(4926), int32(423023))
	mBase = m.M
	v1771 = m.ExcPending
	if v1771 != 0 {
		goto L9
	} else {
		goto L520
	}
L520:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L521:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1778 = m.ExcPending
	if v1778 != 0 {
		goto L9
	} else {
		goto L522
	}
L522:
	;
	F_errmsg(m, int32(411641), int32(0))
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L9
	} else {
		goto L523
	}
L523:
	;
	F_errfinish(m, int32(486904), int32(5153), int32(423023))
	mBase = m.M
	v1787 = m.ExcPending
	if v1787 != 0 {
		goto L9
	} else {
		goto L524
	}
L524:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L525:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1794 = m.ExcPending
	if v1794 != 0 {
		goto L9
	} else {
		goto L526
	}
L526:
	;
	v1795 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v1795 + int32(4)
	F_errmsg(m, int32(263202), v23+int32(48))
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L9
	} else {
		goto L527
	}
L527:
	;
	F_errdetail(m, int32(624632), int32(0))
	mBase = m.M
	v1807 = m.ExcPending
	if v1807 != 0 {
		goto L9
	} else {
		goto L528
	}
L528:
	;
	F_errfinish(m, int32(486904), int32(18864), int32(410323))
	mBase = m.M
	v1812 = m.ExcPending
	if v1812 != 0 {
		goto L9
	} else {
		goto L529
	}
L529:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L530:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L9
	} else {
		goto L531
	}
L531:
	;
	F_errmsg(m, int32(170462), int32(0))
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L9
	} else {
		goto L532
	}
L532:
	;
	F_errfinish(m, int32(486904), int32(5170), int32(423023))
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L9
	} else {
		goto L533
	}
L533:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L534:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_ScanKeyInit(m, v23+int32(96), v1832, int32(3), int32(184), v1841)
	mBase = m.M
	v1843 = m.ExcPending
	if v1843 != 0 {
		goto L9
	} else {
		goto L535
	}
L535:
	;
	v1844 = int32(1)
	v1849 = F_systable_beginscan(m, v1835, v1831, v1844, int32(0), v1844, v23+int32(96))
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L9
	} else {
		goto L536
	}
L536:
	;
	v1851 = F_systable_getnext(m, v1849)
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L9
	} else {
		goto L537
	}
L537:
	;
	if v1851 != 0 {
		goto L538
	} else {
		goto L539
	}
L538:
	;
	if v1493 == int32(29) {
		goto L541
	} else {
		goto L542
	}
L539:
	;
	goto L540
L540:
	;
	F_systable_endscan(m, v1849)
	mBase = m.M
	v1960 = m.ExcPending
	if v1960 != 0 {
		goto L9
	} else {
		goto L564
	}
L541:
	;
	v1857 = int32(96)
	goto L543
L542:
	;
	v1857 = int32(80)
	goto L543
L543:
	;
	v1865 = v1851
	goto L544
L544:
	;
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1865)+16))
	v1879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1878)+22)))
	v1880 = v1878 + v1879
	v1881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1880)+72)))
	if v1881 != int32(102) {
		goto L546
	} else {
		goto L547
	}
L545:
	;
	goto L540
L546:
	;
	v1937 = F_systable_getnext(m, v1849)
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L9
	} else {
		goto L562
	}
L547:
	;
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v1857+v1880)))
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v1885 == v1886 {
		goto L546
	} else {
		goto L548
	}
L548:
	;
	v1889 = F_relation_open(m, v1885, int32(1))
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L9
	} else {
		goto L549
	}
L549:
	;
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1889)+48))
	v1892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1891)+118)))
	if v1493 == int32(29) {
		goto L551
	} else {
		goto L552
	}
L550:
	;
	F_relation_close(m, v1889, int32(1))
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
		goto L9
	} else {
		goto L561
	}
L551:
	;
	if v1892&int32(255) == int32(112) {
		goto L550
	} else {
		goto L554
	}
L552:
	;
	goto L553
L553:
	;
	if v1892&int32(255) == int32(112) {
		goto L13
	} else {
		goto L560
	}
L554:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L9
	} else {
		goto L555
	}
L555:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L9
	} else {
		goto L556
	}
L556:
	;
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v1889)+48))
	v1908 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v1907 + v1908
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v1906 + v1908
	F_errmsg(m, int32(697054), v23+int32(16))
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		goto L9
	} else {
		goto L557
	}
L557:
	;
	F_errtableconstraint(m, l1, v1880+int32(4))
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L9
	} else {
		goto L558
	}
L558:
	;
	F_errfinish(m, int32(486904), int32(18912), int32(410323))
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L9
	} else {
		goto L559
	}
L559:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L560:
	;
	goto L550
L561:
	;
	goto L546
L562:
	;
	if v1937 != 0 {
		v1865 = v1937
		goto L544
	} else {
		goto L563
	}
L563:
	;
	goto L545
L564:
	;
	F_sequence_close(m, v1835, int32(1))
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		goto L9
	} else {
		goto L565
	}
L565:
	;
	if v1493 == int32(29) {
		goto L566
	} else {
		goto L567
	}
L566:
	;
	v1968 = int32(112)
	goto L568
L567:
	;
	v1968 = int32(117)
	goto L568
L568:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v118)+97)) = uint8(v1968)
	v1970 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v118)+96)) = uint8(v1970)
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v118)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+80)) = v1972 | v1970
	v1985 = v189
	v1987 = v190
	goto L14
L569:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2000))) = v2002
	m.G0 = v23 + int32(144)
	return
L570:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		goto L9
	} else {
		goto L571
	}
L571:
	;
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(v1889)+48))
	v2017 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v2016 + v2017
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v2015 + v2017
	F_errmsg(m, int32(697134), v23+int32(32))
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L9
	} else {
		goto L572
	}
L572:
	;
	F_errtableconstraint(m, l1, v1880+int32(4))
	mBase = m.M
	v2031 = m.ExcPending
	if v2031 != 0 {
		goto L9
	} else {
		goto L573
	}
L573:
	;
	F_errfinish(m, int32(486904), int32(18922), int32(410323))
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L9
	} else {
		goto L574
	}
L574:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
