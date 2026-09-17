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
							F_errmsg(m, int32(_a_F_ATExecSetOptions_0), v13+int32(16))
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ATExecSetOptions_1), int32(_a_F_ATExecSetOptions_2), int32(_a_F_ATExecSetOptions_3))
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
								*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v47
								*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v47
								*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v47
								*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v47
								*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v47
								*(*uint8)(unsafe.Add(mBase, uint32(v13)+56)) = uint8(v45)
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
										v80 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecSetOptions[0]))
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
														F_relation_close(m, v17, int32(3))
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
													F_relation_close(m, v17, int32(3))
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
						F_errmsg(m, int32(_a_F_ATExecSetOptions_4), v13)
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ATExecSetOptions_1), int32(_a_F_ATExecSetOptions_5), int32(_a_F_ATExecSetOptions_3))
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
	var v46 int32
	_ = v46
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
	var v105 int32
	_ = v105
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
	var v158 int32
	_ = v158
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
	var v225 int32
	_ = v225
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
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
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
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v441 int32
	_ = v441
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v636 int32
	_ = v636
	var v648 int32
	_ = v648
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
	F_relation_close(m, v636, int32(0))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
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
	v636 = v29
	goto L1
L8:
	;
	v40 = v7
	v46 = v7
	goto L11
L9:
	;
	v105 = v7
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
	v105 = v88
	goto L10
L13:
	;
	v90 = v40 + int32(1)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v90 < v91 {
		v40 = v90
		v46 = v88
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
	v63 = F_lappend(m, v46, v61)
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
	v67 = F_lappend(m, v46, v65)
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
	v71 = F_lappend(m, v46, v57)
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
	v80 = F_list_concat(m, v46, v79)
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
	if v105 == int32(0) {
		v636 = v112
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v116 <= int32(0) {
		v636 = v112
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
	v615 = m.ExcPending
	if v615 != 0 {
		goto L2
	} else {
		goto L133
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L2
	} else {
		goto L130
	}
L35:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
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
	v587 = m.ExcPending
	if v587 != 0 {
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
	v158 = int32(0)
	goto L41
L41:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v149+v158<<(uint(int32(2))%32))))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	if v173 == v142 {
		v225 = v172
		goto L37
	} else {
		goto L43
	}
L42:
	;
	goto L38
L43:
	;
	v176 = v158 + int32(1)
	if v146 != v176 {
		v158 = v176
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
	v213 = int32(_a_F_ATPostAlterTypeParse_0)
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
	v225 = v197
	goto L37
L48:
	;
	goto L36
L49:
	;
	v581 = v133 + int32(1)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v582 <= v581 {
		v636 = v112
		goto L1
	} else {
		goto L126
	}
L50:
	;
	v549 = F_GetComment(m, l0, int32(3381), int32(0))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L2
	} else {
		goto L123
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L2
	} else {
		goto L120
	}
L52:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+4)))
	if v497 == int32(67) {
		goto L111
	} else {
		goto L112
	}
L53:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	if v284 == int32(0) {
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
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v225)+32))
	v281 = F_lappend(m, v280, v275)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+32)) = v281
	goto L49
L68:
	;
	v287 = int32(0)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	if v288 <= v287 {
		goto L49
	} else {
		goto L69
	}
L69:
	;
	v292 = v287
	goto L70
L70:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v284)+12))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v309+v292<<(uint(int32(2))%32))))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)+4))
	switch v314 - int32(14) {
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
	v494 = v292 + int32(1)
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	if v494 < v495 {
		v292 = v494
		goto L70
	} else {
		goto L110
	}
L73:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v313)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v363)+100)) = l2
	if l5 != 0 {
		goto L88
	} else {
		goto L89
	}
L74:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v313)+20))
	v318 = F_get_constraint_index(m, l0)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
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
	v347 = F_GetComment(m, v318, int32(1259), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L2
	} else {
		goto L85
	}
L77:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v317)+12))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v317)+20))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v317)+36))
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317)+64)))
	v324 = F_CheckIndexCompatible(m, v318, v320, v321, v322, v323)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	if v324 == int32(0) {
		goto L76
	} else {
		goto L79
	}
L79:
	;
	v329 = F_index_open(m, v318, int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v329)+48))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+119)))
	if v332 != int32(73) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v329)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v317)+48)) = v335
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v329)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v317)+52)) = v337
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v329)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v317)+56)) = v339
	goto L83
L82:
	;
	goto L83
L83:
	;
	F_relation_close(m, v329, int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	goto L76
L85:
	;
	v349 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v317)+70)) = uint8(v349)
	*(*int32)(unsafe.Add(mBase, uint32(v317)+40)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v313)+4)) = int32(15)
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v225)+32))
	v355 = F_lappend(m, v354, v313)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L2
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+32)) = v355
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v317)+4))
	F_RebuildConstraintComment(m, v225, int32(4), l0, v112, int32(0), v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L2
	} else {
		goto L87
	}
L87:
	;
	goto L72
L88:
	;
	v460 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v363)+60)) = uint8(v460)
	*(*int32)(unsafe.Add(mBase, uint32(v313)+4)) = int32(17)
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v225)+36))
	v465 = F_lappend(m, v464, v313)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L2
	} else {
		goto L107
	}
L89:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v363)+4))
	if v365 != int32(9) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v225)+80))
	if v368 != 0 {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	v370 = F_SearchSysCache1(m, int32(19), l0)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L2
	} else {
		goto L92
	}
L92:
	;
	if v370 == int32(0) {
		goto L48
	} else {
		goto L93
	}
L93:
	;
	v376 = F_SysCacheGetAttrNotNull(m, int32(19), v370, int32(23))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v378 = F_pg_detoast_datum(m, v376)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v378)+4))
	if v380 != int32(1) {
		goto L34
	} else {
		goto L96
	}
L96:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v378)+8))
	if v383 != 0 {
		goto L34
	} else {
		goto L97
	}
L97:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v378)+12))
	if v384 != int32(26) {
		goto L34
	} else {
		goto L98
	}
L98:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v378)+16))
	if int32(0) < v387 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v363)+96))
	v397 = v392
	v401 = int32(0)
	goto L102
L100:
	;
	goto L101
L101:
	;
	F_ReleaseCatCache(m, v370)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L2
	} else {
		goto L106
	}
L102:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v378+int32(24)+v401<<(uint(int32(2))%32))))
	v416 = F_lappend_oid(m, v397, v415)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L2
	} else {
		goto L104
	}
L103:
	;
	goto L101
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363)+96)) = v416
	v420 = v401 + int32(1)
	if v420 != v387 {
		v397 = v416
		v401 = v420
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
	*(*int32)(unsafe.Add(mBase, uint32(v225)+36)) = v465
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v363)+8))
	if v468 == int32(0) {
		goto L72
	} else {
		goto L108
	}
L108:
	;
	F_RebuildConstraintComment(m, v225, int32(5), l0, v112, int32(0), v468)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
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
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v141)+16))
	v502 = F_palloc0(m, int32(32))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
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
	v520 = m.ExcPending
	if v520 != 0 {
		goto L2
	} else {
		goto L117
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v502)+20)) = v141
	*(*int64)(unsafe.Add(mBase, uint32(v502))) = int64(77309411475)
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v225)+36))
	v508 = F_lappend(m, v507, v502)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L2
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+36)) = v508
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v500)+8))
	F_RebuildConstraintComment(m, v225, int32(5), l0, int32(0), v513, v514)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L2
	} else {
		goto L116
	}
L116:
	;
	goto L49
L117:
	;
	v521 = int32(*(*int8)(unsafe.Add(mBase, uint32(v141)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v521
	F_errmsg_internal(m, int32(_a_F_ATPostAlterTypeParse_1), v19+int32(-16))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_ATPostAlterTypeParse_2), int32(_a_F_ATPostAlterTypeParse_3), int32(_a_F_ATPostAlterTypeParse_4))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
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
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v537
	F_errmsg_internal(m, int32(_a_F_ATPostAlterTypeParse_5), v21)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_ATPostAlterTypeParse_2), int32(_a_F_ATPostAlterTypeParse_6), int32(_a_F_ATPostAlterTypeParse_4))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v141)+20)) = v549
	v553 = F_palloc0(m, int32(32))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L2
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v553)+20)) = v141
	*(*int64)(unsafe.Add(mBase, uint32(v553))) = int64(279172874387)
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v225)+60))
	v559 = F_lappend(m, v558, v553)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L2
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+60)) = v559
	goto L49
L126:
	;
	v133 = v581
	goto L35
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = l0
	F_errmsg_internal(m, int32(_a_F_ATPostAlterTypeParse_7), v19+int32(-32))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L2
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_ATPostAlterTypeParse_2), int32(_a_F_ATPostAlterTypeParse_8), int32(_a_F_ATPostAlterTypeParse_9))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_ATPostAlterTypeParse_10), int32(0))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L2
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_ATPostAlterTypeParse_2), int32(_a_F_ATPostAlterTypeParse_11), int32(_a_F_ATPostAlterTypeParse_9))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
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
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v313)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v616
	F_errmsg_internal(m, int32(_a_F_ATPostAlterTypeParse_1), v19+int32(-48))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L2
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_ATPostAlterTypeParse_2), int32(_a_F_ATPostAlterTypeParse_12), int32(_a_F_ATPostAlterTypeParse_4))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
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
	var v117 int32
	_ = v117
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
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
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
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v653 int32
	_ = v653
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v714 int32
	_ = v714
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
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
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v920 int32
	_ = v920
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v965 int32
	_ = v965
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1008 int32
	_ = v1008
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1060 int32
	_ = v1060
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1153 int32
	_ = v1153
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1188 int32
	_ = v1188
	var v1192 int32
	_ = v1192
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1214 int32
	_ = v1214
	var v1218 int32
	_ = v1218
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1237 int32
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1299 int32
	_ = v1299
	var v1304 int32
	_ = v1304
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1329 int32
	_ = v1329
	var v1336 int32
	_ = v1336
	var v1340 int32
	_ = v1340
	var v1345 int32
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1361 int32
	_ = v1361
	var v1366 int32
	_ = v1366
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1383 int32
	_ = v1383
	var v1388 int32
	_ = v1388
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1405 int32
	_ = v1405
	var v1410 int32
	_ = v1410
	var v1414 int32
	_ = v1414
	var v1417 int32
	_ = v1417
	var v1421 int32
	_ = v1421
	var v1425 int32
	_ = v1425
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1443 int32
	_ = v1443
	var v1448 int32
	_ = v1448
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1490 int32
	_ = v1490
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1505 int32
	_ = v1505
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1548 int32
	_ = v1548
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1560 int32
	_ = v1560
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1569 int32
	_ = v1569
	var v1573 int32
	_ = v1573
	var v1578 int32
	_ = v1578
	var v1582 int32
	_ = v1582
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1591 int32
	_ = v1591
	var v1594 int32
	_ = v1594
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	var v1607 int32
	_ = v1607
	var v1612 int32
	_ = v1612
	var v1616 int32
	_ = v1616
	var v1619 int32
	_ = v1619
	var v1623 int32
	_ = v1623
	var v1628 int32
	_ = v1628
	var v1632 int32
	_ = v1632
	var v1635 int32
	_ = v1635
	var v1639 int32
	_ = v1639
	var v1644 int32
	_ = v1644
	var v1648 int32
	_ = v1648
	var v1652 int32
	_ = v1652
	var v1655 int32
	_ = v1655
	var v1660 int32
	_ = v1660
	var v1663 int32
	_ = v1663
	var v1668 int32
	_ = v1668
	var v1671 int32
	_ = v1671
	var v1674 int32
	_ = v1674
	var v1678 int32
	_ = v1678
	var v1682 int32
	_ = v1682
	var v1686 int32
	_ = v1686
	var v1690 int32
	_ = v1690
	var v1694 int32
	_ = v1694
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1703 int32
	_ = v1703
	var v1708 int32
	_ = v1708
	var v1712 int32
	_ = v1712
	var v1715 int32
	_ = v1715
	var v1720 int32
	_ = v1720
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1732 int32
	_ = v1732
	var v1736 int32
	_ = v1736
	var v1741 int32
	_ = v1741
	var v1745 int32
	_ = v1745
	var v1748 int32
	_ = v1748
	var v1752 int32
	_ = v1752
	var v1757 int32
	_ = v1757
	var v1761 int32
	_ = v1761
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1773 int32
	_ = v1773
	var v1777 int32
	_ = v1777
	var v1782 int32
	_ = v1782
	var v1786 int32
	_ = v1786
	var v1789 int32
	_ = v1789
	var v1793 int32
	_ = v1793
	var v1798 int32
	_ = v1798
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1808 int32
	_ = v1808
	var v1811 int32
	_ = v1811
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1825 int32
	_ = v1825
	var v1830 int32
	_ = v1830
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1868 int32
	_ = v1868
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1884 int32
	_ = v1884
	var v1888 int32
	_ = v1888
	var v1893 int32
	_ = v1893
	var v1898 int32
	_ = v1898
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1924 int32
	_ = v1924
	var v1927 int32
	_ = v1927
	var v1932 int32
	_ = v1932
	var v1934 int32
	_ = v1934
	var v1936 int32
	_ = v1936
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1962 int32
	_ = v1962
	var v1963 int32
	_ = v1963
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1973 int32
	_ = v1973
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1989 int32
	_ = v1989
	var v1993 int32
	_ = v1993
	var v1998 int32
	_ = v1998
	var v2022 int32
	_ = v2022
	var v2025 int32
	_ = v2025
	var v2029 int32
	_ = v2029
	var v2034 int32
	_ = v2034
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
		goto L22
	} else {
		goto L23
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
		v117 = v57
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
	v100 = int32(_a_F_ATPrepCmd_0)
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
	v117 = v84
	goto L1
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2022 = m.ExcPending
	if v2022 != 0 {
		goto L9
	} else {
		goto L569
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1973 = m.ExcPending
	if v1973 != 0 {
		goto L9
	} else {
		goto L564
	}
L15:
	;
	v1962 = v117 + v1949<<(uint(int32(2))%32)
	v1963 = *(*int32)(unsafe.Add(mBase, uint32(v1962)+16))
	v1964 = F_lappend(m, v1963, v1950)
	mBase = m.M
	v1965 = m.ExcPending
	if v1965 != 0 {
		goto L9
	} else {
		goto L563
	}
L16:
	;
	v1805 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v1806 = m.ExcPending
	if v1806 != 0 {
		goto L9
	} else {
		goto L528
	}
L17:
	;
	v1801 = int32(2665)
	v1802 = int32(9)
	goto L16
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1786 = m.ExcPending
	if v1786 != 0 {
		goto L9
	} else {
		goto L524
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1761 = m.ExcPending
	if v1761 != 0 {
		goto L9
	} else {
		goto L519
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L9
	} else {
		goto L515
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L9
	} else {
		goto L510
	}
L22:
	;
	v187 = int32(11)
	v188 = F_copyObjectImpl(m, l2)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L9
	} else {
		goto L79
	}
L23:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v130 == int32(61) {
		goto L22
	} else {
		goto L24
	}
L24:
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
		goto L26
	}
L25:
	;
	if v160 != 0 {
		goto L21
	} else {
		goto L38
	}
L26:
	;
	v143 = v134 + int32(-48)
	F_ScanKeyInit(m, v143, int32(1), int32(3), int32(184), v133)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v150 = int32(1)
	v153 = F_systable_beginscan(m, v140, int32(2680), v150, int32(0), v150, v143)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v155 = F_systable_getnext(m, v153)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	if v155 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v155)+16))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+22)))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157+v158)+12)))
	F_systable_endscan(m, v153)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L9
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L9
	} else {
		goto L35
	}
L33:
	;
	F_relation_close(m, v140, int32(3))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	m.G0 = v136 - int32(-64)
	goto L25
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v133
	F_errmsg_internal(m, int32(_a_F_ATPrepCmd_1), v136)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(654), int32(_a_F_ATPrepCmd_3))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	goto L22
L39:
	;
	F_ATSimplePermissions(m, int32(0), l1, int32(305))
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L9
	} else {
		goto L508
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1698 = m.ExcPending
	if v1698 != 0 {
		goto L9
	} else {
		goto L505
	}
L41:
	;
	F_ATSimplePermissions(m, int32(61), l1, int32(256))
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L9
	} else {
		goto L504
	}
L42:
	;
	F_ATSimplePermissions(m, int32(60), l1, int32(256))
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L9
	} else {
		goto L503
	}
L43:
	;
	F_ATSimplePermissions(m, int32(59), l1, int32(320))
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L9
	} else {
		goto L502
	}
L44:
	;
	F_ATSimplePermissions(m, int32(58), l1, int32(32))
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L9
	} else {
		goto L501
	}
L45:
	;
	F_ATSimplePermissions(m, v190, l1, int32(257))
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L9
	} else {
		goto L500
	}
L46:
	;
	F_ATSimplePermissions(m, v190, l1, int32(289))
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L9
	} else {
		goto L498
	}
L47:
	;
	F_ATSimplePermissions(m, int32(53), l1, int32(261))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L9
	} else {
		goto L497
	}
L48:
	;
	F_ATSimplePermissions(m, int32(20), l1, int32(289))
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L9
	} else {
		goto L495
	}
L49:
	;
	F_ATSimplePermissions(m, int32(19), l1, int32(257))
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		goto L9
	} else {
		goto L493
	}
L50:
	;
	F_ATSimplePermissions(m, int32(50), l1, int32(289))
	mBase = m.M
	v1648 = m.ExcPending
	if v1648 != 0 {
		goto L9
	} else {
		goto L492
	}
L51:
	;
	F_ATSimplePermissions(m, int32(49), l1, int32(289))
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L9
	} else {
		goto L471
	}
L52:
	;
	F_ATSimplePermissions(m, v190, l1, int32(271))
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L9
	} else {
		goto L470
	}
L53:
	;
	F_ATSimplePermissions(m, int32(33), l1, int32(333))
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L9
	} else {
		goto L455
	}
L54:
	;
	F_ATSimplePermissions(m, int32(32), l1, int32(261))
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L9
	} else {
		goto L444
	}
L55:
	;
	F_ATSimplePermissions(m, int32(31), l1, int32(289))
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		goto L9
	} else {
		goto L443
	}
L56:
	;
	F_ATSimplePermissions(m, v190, l1, int32(129))
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L9
	} else {
		goto L426
	}
L57:
	;
	F_ATSimplePermissions(m, v190, l1, int32(261))
	mBase = m.M
	v1456 = m.ExcPending
	if v1456 != 0 {
		goto L9
	} else {
		goto L425
	}
L58:
	;
	F_ATSimplePermissions(m, int32(25), l1, int32(32))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L9
	} else {
		goto L424
	}
L59:
	;
	F_ATSimplePermissions(m, int32(24), l1, int32(305))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L9
	} else {
		goto L218
	}
L60:
	;
	F_ATSimplePermissions(m, int32(22), l1, int32(289))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L9
	} else {
		goto L197
	}
L61:
	;
	F_ATSimplePermissions(m, int32(21), l1, int32(257))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L9
	} else {
		goto L196
	}
L62:
	;
	F_ATSimplePermissions(m, int32(16), l1, int32(289))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L9
	} else {
		goto L154
	}
L63:
	;
	F_ATSimplePermissions(m, int32(14), l1, int32(257))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L9
	} else {
		goto L153
	}
L64:
	;
	F_ATSimplePermissions(m, int32(13), l1, int32(305))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L9
	} else {
		goto L135
	}
L65:
	;
	F_ATSimplePermissions(m, int32(12), l1, int32(261))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L9
	} else {
		goto L134
	}
L66:
	;
	F_ATSimplePermissions(m, int32(11), l1, int32(293))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L9
	} else {
		goto L132
	}
L67:
	;
	F_ATSimplePermissions(m, v190, l1, int32(293))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L9
	} else {
		goto L131
	}
L68:
	;
	F_ATSimplePermissions(m, int32(8), l1, int32(365))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L9
	} else {
		goto L129
	}
L69:
	;
	F_ATSimplePermissions(m, int32(7), l1, int32(289))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L9
	} else {
		goto L100
	}
L70:
	;
	F_ATSimplePermissions(m, int32(6), l1, int32(289))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L9
	} else {
		goto L98
	}
L71:
	;
	F_ATSimplePermissions(m, int32(5), l1, int32(289))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L9
	} else {
		goto L96
	}
L72:
	;
	F_ATSimplePermissions(m, int32(4), l1, int32(289))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L9
	} else {
		goto L94
	}
L73:
	;
	F_ATSimplePermissions(m, int32(64), l1, int32(291))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L9
	} else {
		goto L92
	}
L74:
	;
	F_ATSimplePermissions(m, int32(63), l1, int32(291))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L9
	} else {
		goto L90
	}
L75:
	;
	F_ATSimplePermissions(m, int32(62), l1, int32(291))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L9
	} else {
		goto L88
	}
L76:
	;
	F_ATSimplePermissions(m, int32(3), l1, int32(289))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L87
	}
L77:
	;
	F_ATSimplePermissions(m, int32(2), l1, int32(291))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L9
	} else {
		goto L82
	}
L78:
	;
	v191 = int32(2)
	F_ATSimplePermissions(m, int32(1), l1, v191)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L9
	} else {
		goto L80
	}
L79:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	switch v190 {
	case 0:
		goto L39
	case 1:
		goto L78
	case 2:
		goto L77
	case 3:
		goto L76
	case 4:
		goto L72
	case 5:
		goto L71
	case 6:
		goto L70
	case 7:
		goto L69
	case 8:
		goto L68
	case 9, 10:
		goto L67
	case 11:
		goto L66
	case 12:
		goto L65
	case 13:
		goto L64
	case 14:
		goto L63
	default:
		goto L40
	case 16:
		goto L62
	case 19:
		goto L49
	case 20:
		goto L48
	case 21:
		goto L61
	case 22:
		goto L60
	case 24:
		goto L59
	case 25:
		goto L58
	case 26:
		v1949 = v187
		v1950 = v188
		goto L15
	case 27, 28:
		goto L57
	case 29, 30:
		goto L56
	case 31:
		goto L55
	case 32:
		goto L54
	case 33:
		goto L53
	case 34, 35, 36:
		goto L52
	case 37, 38, 39, 40, 41, 42, 43, 44:
		goto L46
	case 45, 46, 47, 48, 51, 52, 54, 55, 56, 57:
		goto L45
	case 49:
		goto L51
	case 50:
		goto L50
	case 53:
		goto L47
	case 58:
		goto L44
	case 59:
		goto L43
	case 60:
		goto L42
	case 61:
		goto L41
	case 62:
		goto L75
	case 63:
		goto L74
	case 64:
		goto L73
	}
L80:
	;
	F_ATPrepAddColumn(m, l0, l1, l3, l4, int32(1), v188, l5, l6)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L9
	} else {
		goto L81
	}
L81:
	;
	v1949 = v191
	v1950 = v188
	goto L15
L82:
	;
	F_ATSimpleRecursion(m, l0, l1, v188, l3, l5, l6)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L9
	} else {
		goto L83
	}
L83:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v188)+20))
	if v207 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v208 = int32(10)
	goto L86
L85:
	;
	v208 = int32(0)
	goto L86
L86:
	;
	v1949 = v208
	v1950 = v188
	goto L15
L87:
	;
	v1949 = int32(10)
	v1950 = v188
	goto L15
L88:
	;
	v218 = int32(10)
	if l3 == int32(0) {
		v1949 = v218
		v1950 = v188
		goto L15
	} else {
		goto L89
	}
L89:
	;
	v221 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v188)+29)) = uint8(v221)
	v1949 = v218
	v1950 = v188
	goto L15
L90:
	;
	if l3 == int32(0) {
		v1949 = v187
		v1950 = v188
		goto L15
	} else {
		goto L91
	}
L91:
	;
	v229 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v188)+29)) = uint8(v229)
	v1949 = v187
	v1950 = v188
	goto L15
L92:
	;
	v235 = int32(0)
	if l3 == v235 {
		v1949 = v235
		v1950 = v188
		goto L15
	} else {
		goto L93
	}
L93:
	;
	v238 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v188)+29)) = uint8(v238)
	v1949 = v235
	v1950 = v188
	goto L15
L94:
	;
	v244 = int32(0)
	if l3 == v244 {
		v1949 = v244
		v1950 = v188
		goto L15
	} else {
		goto L95
	}
L95:
	;
	v247 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v188)+29)) = uint8(v247)
	v1949 = v244
	v1950 = v188
	goto L15
L96:
	;
	v253 = int32(7)
	if l3 == int32(0) {
		v1949 = v253
		v1950 = v188
		goto L15
	} else {
		goto L97
	}
L97:
	;
	v256 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v188)+29)) = uint8(v256)
	v1949 = v253
	v1950 = v188
	goto L15
L98:
	;
	F_ATSimpleRecursion(m, l0, l1, v188, l3, l5, l6)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L9
	} else {
		goto L99
	}
L99:
	;
	v1949 = int32(3)
	v1950 = v188
	goto L15
L100:
	;
	F_ATSimpleRecursion(m, l0, l1, v188, l3, l5, l6)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L9
	} else {
		goto L101
	}
L101:
	;
	v271 = m.G0
	v273 = v271 - int32(16)
	m.G0 = v273
	if l3 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L102:
	;
	v1949 = int32(0)
	v1950 = v188
	goto L15
L103:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L9
	} else {
		goto L125
	}
L104:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L9
	} else {
		goto L121
	}
L105:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L9
	} else {
		goto L117
	}
L106:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v278 = F_find_inheritance_children(m, v277, l5)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L9
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	if l4 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	if v278 != 0 {
		goto L105
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
	v284 = F_SearchSysCacheCopyAttName(m, v282, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L9
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	m.G0 = v273 + int32(16)
	goto L102
L114:
	;
	if v284 == int32(0) {
		goto L104
	} else {
		goto L115
	}
L115:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v284)+16))
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288)+22)))
	v291 = int32(*(*int16)(unsafe.Add(mBase, uint32(v288+v289)+94)))
	if int32(0) < v291 {
		goto L103
	} else {
		goto L116
	}
L116:
	;
	goto L113
L117:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L9
	} else {
		goto L118
	}
L118:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_4), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L9
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_6), int32(_a_F_ATPrepCmd_7))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L9
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L9
	} else {
		goto L122
	}
L122:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v273))) = v323
	*(*int32)(unsafe.Add(mBase, uint32(v273)+4)) = v322 + int32(4)
	F_errmsg(m, int32(_a_F_ATPrepCmd_8), v273)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L9
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_9), int32(_a_F_ATPrepCmd_7))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L9
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L9
	} else {
		goto L126
	}
L126:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_10), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L9
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_11), int32(_a_F_ATPrepCmd_7))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L9
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	F_ATSimpleRecursion(m, l0, l1, v188, l3, l5, l6)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L9
	} else {
		goto L130
	}
L130:
	;
	v1949 = v187
	v1950 = v188
	goto L15
L131:
	;
	v1949 = v187
	v1950 = v188
	goto L15
L132:
	;
	F_ATSimpleRecursion(m, l0, l1, v188, l3, l5, l6)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L9
	} else {
		goto L133
	}
L133:
	;
	v1949 = v187
	v1950 = v188
	goto L15
L134:
	;
	v1949 = v187
	v1950 = v188
	goto L15
L135:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if l4 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v1949 = int32(0)
	v1950 = v188
	goto L15
L137:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L9
	} else {
		goto L149
	}
L138:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v376)+76))
	if v379 != 0 {
		goto L137
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+119)))
	if v380 == int32(99) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L140
L142:
	;
	F_ATTypedTableRecursion(m, l0, l1, v188, l5, l6)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L9
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	if l3 != 0 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	goto L144
L146:
	;
	v385 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v188)+29)) = uint8(v385)
	goto L148
L147:
	;
	goto L148
L148:
	;
	goto L136
L149:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L9
	} else {
		goto L150
	}
L150:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_12), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L9
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_13), int32(_a_F_ATPrepCmd_14))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L9
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
	v1949 = int32(9)
	v1950 = v188
	goto L15
L154:
	;
	v413 = int32(0)
	v415 = m.G0
	v417 = v415 - int32(16)
	m.G0 = v417
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v188)+20))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)+4))
	if v420 != int32(6) {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	v616 = int32(6)
	if l3 == int32(0) {
		v1949 = v616
		v1950 = v188
		goto L15
	} else {
		goto L194
	}
L156:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L9
	} else {
		goto L190
	}
L157:
	;
	m.G0 = v417 + int32(16)
	goto L155
L158:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v419)+32))
	if v423 == int32(0) {
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v423)+4))
	if v426 <= int32(0) {
		goto L157
	} else {
		goto L160
	}
L160:
	;
	v433 = v413
	v438 = v413
	v444 = v8
	goto L161
L161:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v423)+12))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v450+v444<<(uint(int32(2))%32))))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v454)+4))
	v456 = F_findNotNullConstraint(m, v449, v455)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L9
	} else {
		goto L164
	}
L162:
	;
	goto L157
L163:
	;
	v573 = v444 + int32(1)
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v423)+4))
	if v573 < v574 {
		v433 = v556
		v438 = v561
		v444 = v573
		goto L161
	} else {
		goto L189
	}
L164:
	;
	if v456 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v454)+4))
	F_verifyNotNullPKCompatible(m, v456, v458)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L9
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	if l3 != 0 {
		v522 = v433
		v527 = v438
		goto L170
	} else {
		goto L171
	}
L168:
	;
	F_pfree(m, v456)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L9
	} else {
		goto L169
	}
L169:
	;
	v556 = v433
	v561 = v438
	goto L163
L170:
	;
	v538 = F_makeNotNullConstraint(m, v454)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L9
	} else {
		goto L186
	}
L171:
	;
	if v433&int32(1) == int32(0) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v468 = F_find_inheritance_children(m, v467, l5)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L9
	} else {
		goto L175
	}
L173:
	;
	v470 = v438
	goto L174
L174:
	;
	if v470 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v470 = v468
	goto L174
L176:
	;
	v522 = int32(1)
	v527 = int32(0)
	goto L170
L177:
	;
	goto L178
L178:
	;
	v476 = int32(0)
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	if v477 <= v476 {
		v522 = int32(1)
		v527 = v470
		goto L170
	} else {
		goto L179
	}
L179:
	;
	v480 = v476
	goto L180
L180:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v470)+12))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v500+v480<<(uint(int32(2))%32))))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v454)+4))
	v506 = F_findNotNullConstraint(m, v504, v505)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L9
	} else {
		goto L182
	}
L181:
	;
	v522 = v513
	v527 = v470
	goto L170
L182:
	;
	if v506 == int32(0) {
		goto L156
	} else {
		goto L183
	}
L183:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v454)+4))
	F_verifyNotNullPKCompatible(m, v506, v510)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L9
	} else {
		goto L184
	}
L184:
	;
	v513 = int32(1)
	v515 = v480 + v513
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	if v515 < v516 {
		v480 = v515
		goto L180
	} else {
		goto L185
	}
L185:
	;
	goto L181
L186:
	;
	v541 = F_palloc0(m, int32(32))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L9
	} else {
		goto L187
	}
L187:
	;
	v543 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v541)+29)) = uint8(v543)
	*(*int64)(unsafe.Add(mBase, uint32(v541))) = int64(68719476883)
	*(*int32)(unsafe.Add(mBase, uint32(v541)+20)) = v538
	F_ATPrepCmd(m, l0, l1, v541, v543, int32(0), l5, l6)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L9
	} else {
		goto L188
	}
L188:
	;
	v556 = v522
	v561 = v527
	goto L163
L189:
	;
	goto L162
L190:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v454)+4))
	v604 = F_get_rel_name(m, v504)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L9
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v417)+4)) = v604
	*(*int32)(unsafe.Add(mBase, uint32(v417))) = v603
	F_errmsg(m, int32(_a_F_ATPrepCmd_15), v417)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L9
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_16), int32(_a_F_ATPrepCmd_17))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L9
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v621 = F_find_all_inheritors(m, v619, l5, int32(0))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L9
	} else {
		goto L195
	}
L195:
	;
	v623 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v188)+29)) = uint8(v623)
	v1949 = v616
	v1950 = v188
	goto L15
L196:
	;
	v1949 = int32(8)
	v1950 = v188
	goto L15
L197:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634)+119)))
	if v635 == int32(112) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v735 = int32(0)
	if l3 == v735 {
		v1949 = v735
		v1950 = v188
		goto L15
	} else {
		goto L217
	}
L199:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v640 = F_find_all_inheritors(m, v638, l5, int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L9
	} else {
		goto L203
	}
L200:
	;
	goto L201
L201:
	;
	goto L198
L202:
	;
	F_list_free(m, v640)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L9
	} else {
		goto L216
	}
L203:
	;
	if v640 == int32(0) {
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v640)+4))
	if v644 < int32(2) {
		goto L202
	} else {
		goto L205
	}
L205:
	;
	v653 = int32(1)
	goto L206
L206:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v640)+12))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v668+v653<<(uint(int32(2))%32))))
	v674 = F_table_open(m, v672, int32(0))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L9
	} else {
		goto L208
	}
L207:
	;
	goto L202
L208:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v674)+48))
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676)+118)))
	if v677 == int32(116) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674)+24)))
	if v680 == int32(0) {
		goto L13
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	F_CheckTableNotInUse(m, v674, int32(_a_F_ATPrepCmd_18))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L9
	} else {
		goto L213
	}
L212:
	;
	goto L211
L213:
	;
	F_relation_close(m, v674, int32(0))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L9
	} else {
		goto L214
	}
L214:
	;
	v690 = v653 + int32(1)
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v640)+4))
	if v690 < v691 {
		v653 = v690
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
	v738 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v188)+29)) = uint8(v738)
	v1949 = v735
	v1950 = v188
	goto L15
L218:
	;
	v745 = F_ATParseTransformCmd(m, v117, l1, v188, l3, int32(-1), l6)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L9
	} else {
		goto L219
	}
L219:
	;
	v747 = m.G0
	v749 = v747 - int32(208)
	m.G0 = v749
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v745)+8))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v745)+20))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v752)+32))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v752)+8))
	v756 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L9
	} else {
		goto L220
	}
L220:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v756)+4)) = v758
	if l4 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L221:
	;
	v1949 = int32(1)
	v1950 = v745
	goto L15
L222:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L9
	} else {
		goto L420
	}
L223:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L9
	} else {
		goto L415
	}
L224:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L9
	} else {
		goto L411
	}
L225:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L9
	} else {
		goto L407
	}
L226:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L9
	} else {
		goto L403
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v749)+148)) = v863
	*(*int32)(unsafe.Add(mBase, uint32(v749)+144)) = v751
	F_errmsg(m, int32(_a_F_ATPrepCmd_19), v749+int32(144))
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L9
	} else {
		goto L400
	}
L228:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L9
	} else {
		goto L395
	}
L229:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L9
	} else {
		goto L390
	}
L230:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L9
	} else {
		goto L384
	}
L231:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L9
	} else {
		goto L379
	}
L232:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L9
	} else {
		goto L374
	}
L233:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L9
	} else {
		goto L369
	}
L234:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v762)+76))
	if v763 != 0 {
		goto L233
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v765 = F_SearchSysCacheAttName(m, v764, v751)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L9
	} else {
		goto L238
	}
L237:
	;
	goto L236
L238:
	;
	if v765 == int32(0) {
		goto L232
	} else {
		goto L239
	}
L239:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v765)+16))
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769)+22)))
	v771 = v769 + v770
	v772 = int32(*(*int16)(unsafe.Add(mBase, uint32(v771)+74)))
	if v772 <= int32(0) {
		goto L231
	} else {
		goto L240
	}
L240:
	;
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v771)+90)))
	if v775 != 0 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v752)+32))
	if v776 != 0 {
		goto L230
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	if l4 == int32(0) {
		goto L245
	} else {
		goto L246
	}
L244:
	;
	goto L243
L245:
	;
	v779 = int32(*(*int16)(unsafe.Add(mBase, uint32(v771)+94)))
	if int32(0) < v779 {
		goto L229
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	v784 = F_bms_make_singleton(m, v772+int32(7))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L9
	} else {
		goto L249
	}
L248:
	;
	goto L247
L249:
	;
	v788 = F_has_partition_attrs(m, l1, v784, v749+int32(199))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L9
	} else {
		goto L250
	}
L250:
	;
	if v788 != 0 {
		goto L228
	} else {
		goto L251
	}
L251:
	;
	F_typenameTypeIdAndMod(m, v756, v754, v749+int32(204), v749+int32(200))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L9
	} else {
		goto L252
	}
L252:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v749)+204))
	v799 = *(*int32)(unsafe.Add(mBase, _c_F_ATPrepCmd[0]))
	v801 = F_object_aclcheck(m, int32(1247), v797, v799, int64(256))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L9
	} else {
		goto L253
	}
L253:
	;
	if v801 != 0 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v749)+204))
	F_aclcheck_error_type(m, v801, v803)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L9
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v749)+204))
	v807 = F_GetColumnDefCollation(m, v756, v752, v806)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L9
	} else {
		goto L258
	}
L257:
	;
	goto L256
L258:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v809)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v749)+192)) = v810
	*(*int32)(unsafe.Add(mBase, uint32(v749)+156)) = v810
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v749)+204))
	v817 = F_list_make1_impl(m, int32(472), v749+int32(156))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L9
	} else {
		goto L259
	}
L259:
	;
	v821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v771)+90)))
	if v821 == int32(118) {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v824 = int32(8)
	goto L262
L261:
	;
	v824 = int32(0)
	goto L262
L262:
	;
	F_CheckAttributeType(m, v751, v813, v807, v817, v824)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L9
	} else {
		goto L263
	}
L263:
	;
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v771)+90)))
	if v827 == int32(118) {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v1032 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+4)))
	switch v1032 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L319
	default:
		goto L318
	}
L265:
	;
	v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+4)))
	switch v830 - int32(112) {
	case 0, 2:
		goto L267
	default:
		goto L266
	}
L266:
	;
	if v753 != 0 {
		goto L226
	} else {
		goto L316
	}
L267:
	;
	if v753 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v771)+68))
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v771)+76))
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v771)+96))
	v840 = F_makeVar(m, int32(1), v772, v836, v837, v838, int32(0))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L9
	} else {
		goto L271
	}
L269:
	;
	v842 = v753
	goto L270
L270:
	;
	v843 = F_exprType(m, v842)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L9
	} else {
		goto L272
	}
L271:
	;
	v842 = v840
	goto L270
L272:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v749)+204))
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v749)+200))
	v850 = F_coerce_to_target_type(m, v756, v842, v843, v845, v846, int32(1), int32(2), int32(-1))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L9
	} else {
		goto L273
	}
L273:
	;
	if v850 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v752)+32))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L9
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	F_assign_expr_collations(m, v756, v850)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L9
	} else {
		goto L289
	}
L277:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L9
	} else {
		goto L278
	}
L278:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v749)+204))
	v863 = F_format_type_be(m, v862)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L9
	} else {
		goto L279
	}
L279:
	;
	if v854 != 0 {
		goto L227
	} else {
		goto L280
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v749)+132)) = v863
	*(*int32)(unsafe.Add(mBase, uint32(v749)+128)) = v751
	F_errmsg(m, int32(_a_F_ATPrepCmd_20), v749+int32(128))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L9
	} else {
		goto L281
	}
L281:
	;
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v771)+90)))
	if v872 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v875 = F_quote_identifier(m, v751)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L9
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_21), int32(_a_F_ATPrepCmd_22))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L9
	} else {
		goto L288
	}
L285:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v749)+204))
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v749)+200))
	v879 = F_format_type_with_typemod(m, v877, v878)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L9
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v749)+116)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v749)+112)) = v875
	F_errhint(m, int32(_a_F_ATPrepCmd_23), v749+int32(112))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L9
	} else {
		goto L287
	}
L287:
	;
	goto L284
L288:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L289:
	;
	v897 = F_expand_generated_columns_in_expr(m, v850, l1, int32(1))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L9
	} else {
		goto L290
	}
L290:
	;
	v899 = F_expression_planner(m, v897)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L9
	} else {
		goto L291
	}
L291:
	;
	v902 = F_palloc0(m, int32(16))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L9
	} else {
		goto L292
	}
L292:
	;
	v904 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v902)+12)) = uint8(v904)
	*(*int32)(unsafe.Add(mBase, uint32(v902)+4)) = v899
	*(*uint16)(unsafe.Add(mBase, uint32(v902))) = uint16(v772)
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v117)+68))
	v909 = F_lappend(m, v908, v902)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L9
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+68)) = v909
	v920 = v899
	goto L295
L294:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v117)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+80)) = v1008 | int32(4)
	goto L264
L295:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v920)))
	switch v932 - int32(6) {
	case 0:
		goto L297
	case 1, 2, 3, 4, 5, 6, 7, 8, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20:
		goto L294
	case 9:
		goto L298
	case 21:
		goto L300
	default:
		goto L299
	}
L296:
	;
	v1002 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920)+8)))
	if v1002 == v772&int32(_a_F_ATPrepCmd_24) {
		goto L264
	} else {
		goto L315
	}
L297:
	;
	goto L296
L298:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v920)+4))
	if base.Ui32(int32(1)) < base.Ui32(v942-int32(2027)) {
		goto L294
	} else {
		goto L304
	}
L299:
	;
	if v932 != int32(55) {
		goto L294
	} else {
		goto L301
	}
L300:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v920)+4))
	v920 = v935
	goto L295
L301:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v920)+8))
	v939 = F_DomainHasConstraints(m, v938)
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L9
	} else {
		goto L302
	}
L302:
	;
	if v939 != 0 {
		goto L294
	} else {
		goto L303
	}
L303:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v920)+4))
	v920 = v941
	goto L295
L304:
	;
	v947 = m.G0
	v949 = v947 - int32(16)
	m.G0 = v949
	v952 = *(*int32)(unsafe.Add(mBase, _c_F_ATPrepCmd[1]))
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v952)+uint32(_c_F_ATPrepCmd[2])))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v952)+264))
	if v959 < int32(2) {
		goto L306
	} else {
		goto L307
	}
L305:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v949)+12))
	m.G0 = v949 + int32(16)
	if v992|(v991^int32(1)) != 0 {
		goto L294
	} else {
		goto L314
	}
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v949+int32(12)))) = v958
	v991 = int32(1)
	goto L305
L307:
	;
	v965 = int32(1)
	goto L308
L308:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v952+int32(_a_F_ATPrepCmd_25)+v965<<(uint(int32(4))%32))))
	if v958 == v973 {
		goto L310
	} else {
		goto L311
	}
L309:
	;
	v991 = int32(0)
	goto L305
L310:
	;
	v976 = v965 + int32(1)
	if v959 != v976 {
		v965 = v976
		goto L308
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	goto L309
L313:
	;
	goto L306
L314:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v920)+28))
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v999)+12))
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v1000)))
	v920 = v1001
	goto L295
L315:
	;
	goto L294
L316:
	;
	goto L264
L317:
	;
	F_ReleaseCatCache(m, v765)
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L9
	} else {
		goto L322
	}
L318:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+72))
	F_find_composite_type_dependencies(m, v1039, l1, int32(0))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L9
	} else {
		goto L321
	}
L319:
	;
	v1035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v771)+90)))
	if v1035 != int32(118) {
		goto L317
	} else {
		goto L320
	}
L320:
	;
	goto L318
L321:
	;
	goto L317
L322:
	;
	if l3 != 0 {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	v1188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+4)))
	if v1188 == int32(99) {
		goto L365
	} else {
		goto L366
	}
L324:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v1048 = F_find_all_inheritors(m, v1045, l5, v749+int32(188))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L9
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	if l4 != 0 {
		v1170 = v745
		goto L323
	} else {
		goto L362
	}
L327:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v749)+188))
	v1052 = v745
	v1060 = int32(0)
	goto L328
L328:
	;
	v1072 = int32(0)
	if v1048 == v1072 {
		v1082 = v1072
		goto L330
	} else {
		goto L331
	}
L330:
	;
	if v1050 == int32(0) {
		v1170 = v745
		goto L323
	} else {
		goto L333
	}
L331:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+4))
	if v1076 <= v1060 {
		v1082 = int32(0)
		goto L330
	} else {
		goto L332
	}
L332:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+12))
	v1082 = v1078 + v1060<<(uint(int32(2))%32)
	goto L330
L333:
	;
	v1085 = int32(0)
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1050)+4))
	if base.B2i32(v1082 == v1085)|base.B2i32(v1087 <= v1060) == v1085 {
		goto L335
	} else {
		goto L336
	}
L334:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1082)))
	if v1045 != v1094 {
		goto L339
	} else {
		goto L340
	}
L335:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1050)+12))
	if v1092 != 0 {
		goto L334
	} else {
		goto L338
	}
L336:
	;
	goto L337
L337:
	;
	v1170 = v1052
	goto L323
L338:
	;
	goto L337
L339:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1092+v1060<<(uint(int32(2))%32))))
	v1101 = F_relation_open(m, v1094, int32(0))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L9
	} else {
		goto L342
	}
L340:
	;
	v1157 = v1052
	goto L341
L341:
	;
	v1052 = v1157
	v1060 = v1060 + int32(1)
	goto L328
L342:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+48))
	v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1103)+118)))
	if v1104 == int32(116) {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v1107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1101)+24)))
	if v1107 == int32(0) {
		goto L13
	} else {
		goto L346
	}
L344:
	;
	goto L345
L345:
	;
	F_CheckTableNotInUse(m, v1101, int32(_a_F_ATPrepCmd_18))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L9
	} else {
		goto L347
	}
L346:
	;
	goto L345
L347:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+56))
	v1114 = F_SearchSysCacheAttName(m, v1113, v751)
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L9
	} else {
		goto L348
	}
L348:
	;
	if v1114 == int32(0) {
		goto L225
	} else {
		goto L349
	}
L349:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v1114)+16))
	v1119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1118)+22)))
	v1121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1118+v1119)+94)))
	if v1099 < v1121 {
		goto L224
	} else {
		goto L350
	}
L350:
	;
	F_ReleaseCatCache(m, v1114)
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L9
	} else {
		goto L351
	}
L351:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v752)+32))
	if v1125 != 0 {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v1126 = F_copyObjectImpl(m, v1052)
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L9
	} else {
		goto L355
	}
L353:
	;
	v1147 = v1052
	goto L354
L354:
	;
	F_ATPrepCmd(m, l0, v1101, v1147, int32(0), int32(1), l5, l6)
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L9
	} else {
		goto L360
	}
L355:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+52))
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v1131 = F_build_attrmap_by_name(m, v1128, v1129, int32(0))
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L9
	} else {
		goto L356
	}
L356:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v752)+32))
	v1138 = F_map_variable_attnos(m, v1133, int32(1), v1131, int32(0), v749+int32(187))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L9
	} else {
		goto L357
	}
L357:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1140)+32)) = v1138
	v1142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749)+187)))
	if v1142 == int32(1) {
		goto L223
	} else {
		goto L358
	}
L358:
	;
	F_pfree(m, v1131)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L9
	} else {
		goto L359
	}
L359:
	;
	v1147 = v1126
	goto L354
L360:
	;
	F_relation_close(m, v1101, int32(0))
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L9
	} else {
		goto L361
	}
L361:
	;
	v1157 = v1147
	goto L341
L362:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v1166 = F_find_inheritance_children(m, v1164, int32(0))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L9
	} else {
		goto L363
	}
L363:
	;
	if v1166 != 0 {
		goto L222
	} else {
		goto L364
	}
L364:
	;
	v1170 = v745
	goto L323
L365:
	;
	F_ATTypedTableRecursion(m, l0, l1, v1170, l5, l6)
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L9
	} else {
		goto L368
	}
L366:
	;
	goto L367
L367:
	;
	m.G0 = v749 + int32(208)
	goto L221
L368:
	;
	goto L367
L369:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1202 = m.ExcPending
	if v1202 != 0 {
		goto L9
	} else {
		goto L370
	}
L370:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_26), int32(0))
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L9
	} else {
		goto L371
	}
L371:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v752)+64))
	F_parser_errposition(m, v756, v1207)
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L9
	} else {
		goto L372
	}
L372:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_27), int32(_a_F_ATPrepCmd_22))
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L9
	} else {
		goto L373
	}
L373:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L374:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L9
	} else {
		goto L375
	}
L375:
	;
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v749))) = v751
	*(*int32)(unsafe.Add(mBase, uint32(v749)+4)) = v1222 + int32(4)
	F_errmsg(m, int32(_a_F_ATPrepCmd_8), v749)
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L9
	} else {
		goto L376
	}
L376:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v752)+64))
	F_parser_errposition(m, v756, v1230)
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L9
	} else {
		goto L377
	}
L377:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_28), int32(_a_F_ATPrepCmd_22))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L9
	} else {
		goto L378
	}
L378:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L379:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L9
	} else {
		goto L380
	}
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v749)+16)) = v751
	F_errmsg(m, int32(_a_F_ATPrepCmd_29), v749+int32(16))
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L9
	} else {
		goto L381
	}
L381:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v752)+64))
	F_parser_errposition(m, v756, v1251)
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L9
	} else {
		goto L382
	}
L382:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_30), int32(_a_F_ATPrepCmd_22))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L9
	} else {
		goto L383
	}
L383:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L384:
	;
	F_errcode(m, int32(17064068))
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L9
	} else {
		goto L385
	}
L385:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_31), int32(0))
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L9
	} else {
		goto L386
	}
L386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v749)+176)) = v751
	F_errdetail(m, int32(_a_F_ATPrepCmd_32), v749+int32(176))
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L9
	} else {
		goto L387
	}
L387:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v752)+64))
	F_parser_errposition(m, v756, v1276)
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L9
	} else {
		goto L388
	}
L388:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_33), int32(_a_F_ATPrepCmd_22))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L9
	} else {
		goto L389
	}
L389:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L390:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L9
	} else {
		goto L391
	}
L391:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v749)+160)) = v751
	F_errmsg(m, int32(_a_F_ATPrepCmd_34), v749+int32(160))
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L9
	} else {
		goto L392
	}
L392:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v752)+64))
	F_parser_errposition(m, v756, v1297)
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L9
	} else {
		goto L393
	}
L393:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_35), int32(_a_F_ATPrepCmd_22))
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L9
	} else {
		goto L394
	}
L394:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L395:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L9
	} else {
		goto L396
	}
L396:
	;
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v749)+32)) = v751
	*(*int32)(unsafe.Add(mBase, uint32(v749)+36)) = v1312 + int32(4)
	F_errmsg(m, int32(_a_F_ATPrepCmd_36), v749+int32(32))
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L9
	} else {
		goto L397
	}
L397:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v752)+64))
	F_parser_errposition(m, v756, v1322)
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L9
	} else {
		goto L398
	}
L398:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_37), int32(_a_F_ATPrepCmd_22))
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L9
	} else {
		goto L399
	}
L399:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L400:
	;
	F_errhint(m, int32(_a_F_ATPrepCmd_38), int32(0))
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L9
	} else {
		goto L401
	}
L401:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_39), int32(_a_F_ATPrepCmd_22))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L9
	} else {
		goto L402
	}
L402:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L403:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L9
	} else {
		goto L404
	}
L404:
	;
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v749)+96)) = v1353 + int32(4)
	F_errmsg(m, int32(_a_F_ATPrepCmd_40), v749+int32(96))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L9
	} else {
		goto L405
	}
L405:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_41), int32(_a_F_ATPrepCmd_22))
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L9
	} else {
		goto L406
	}
L406:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L407:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L9
	} else {
		goto L408
	}
L408:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v749)+48)) = v751
	*(*int32)(unsafe.Add(mBase, uint32(v749)+52)) = v1374 + int32(4)
	F_errmsg(m, int32(_a_F_ATPrepCmd_8), v749+int32(48))
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L9
	} else {
		goto L409
	}
L409:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_42), int32(_a_F_ATPrepCmd_22))
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L9
	} else {
		goto L410
	}
L410:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L411:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L9
	} else {
		goto L412
	}
L412:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v749)+64)) = v751
	*(*int32)(unsafe.Add(mBase, uint32(v749)+68)) = v1396 + int32(4)
	F_errmsg(m, int32(_a_F_ATPrepCmd_43), v749-int32(-64))
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L9
	} else {
		goto L413
	}
L413:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_44), int32(_a_F_ATPrepCmd_22))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L9
	} else {
		goto L414
	}
L414:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L415:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L9
	} else {
		goto L416
	}
L416:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_45), int32(0))
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L9
	} else {
		goto L417
	}
L417:
	;
	F_errdetail(m, int32(_a_F_ATPrepCmd_46), int32(0))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L9
	} else {
		goto L418
	}
L418:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_47), int32(_a_F_ATPrepCmd_22))
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		goto L9
	} else {
		goto L419
	}
L419:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L420:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L9
	} else {
		goto L421
	}
L421:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v749)+80)) = v751
	F_errmsg(m, int32(_a_F_ATPrepCmd_48), v749+int32(80))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L9
	} else {
		goto L422
	}
L422:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_49), int32(_a_F_ATPrepCmd_22))
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L9
	} else {
		goto L423
	}
L423:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L424:
	;
	v1949 = v187
	v1950 = v188
	goto L15
L425:
	;
	v1949 = v187
	v1950 = v188
	goto L15
L426:
	;
	v1460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+96)))
	if v1460 == int32(1) {
		goto L20
	} else {
		goto L427
	}
L427:
	;
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1464)+118)))
	switch v1465 - int32(112) {
	case 0:
		goto L429
	default:
		goto L430
	case 4:
		goto L432
	case 5:
		goto L431
	}
L428:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v1498 = F_GetRelationPublications(m, v1497)
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L9
	} else {
		goto L441
	}
L429:
	;
	if v1463 == int32(29) {
		v1949 = v187
		v1950 = v188
		goto L15
	} else {
		goto L440
	}
L430:
	;
	if v1463 == int32(29) {
		goto L17
	} else {
		goto L439
	}
L431:
	;
	if v1463 != int32(29) {
		v1949 = v187
		v1950 = v188
		goto L15
	} else {
		goto L438
	}
L432:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L9
	} else {
		goto L433
	}
L433:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L9
	} else {
		goto L434
	}
L434:
	;
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v1475 + int32(4)
	F_errmsg(m, int32(_a_F_ATPrepCmd_50), v23-int32(-64))
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L9
	} else {
		goto L435
	}
L435:
	;
	F_errtable(m, l1)
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L9
	} else {
		goto L436
	}
L436:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_51), int32(_a_F_ATPrepCmd_52))
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		goto L9
	} else {
		goto L437
	}
L437:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L438:
	;
	goto L17
L439:
	;
	goto L428
L440:
	;
	goto L428
L441:
	;
	if v1498 != 0 {
		goto L19
	} else {
		goto L442
	}
L442:
	;
	v1801 = int32(0)
	v1802 = int32(13)
	goto L16
L443:
	;
	v1949 = int32(0)
	v1950 = v188
	goto L15
L444:
	;
	v1511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+84)))
	if v1511 == int32(1) {
		goto L18
	} else {
		goto L445
	}
L445:
	;
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
	if v1514 != 0 {
		goto L447
	} else {
		goto L448
	}
L446:
	;
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1527)+84))
	if v1526 != v1528 {
		goto L452
	} else {
		goto L453
	}
L447:
	;
	v1522 = v1514
	goto L449
L448:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1516)+119)))
	if v1517 == int32(112) {
		v1526 = int32(0)
		goto L446
	} else {
		goto L450
	}
L449:
	;
	v1524 = F_get_table_am_oid(m, v1522, int32(0))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L9
	} else {
		goto L451
	}
L450:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, _c_F_ATPrepCmd[3]))
	v1522 = v1521
	goto L449
L451:
	;
	v1526 = v1524
	goto L446
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+88)) = v1526
	v1531 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v117)+84)) = uint8(v1531)
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v117)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+80)) = v1533 | int32(8)
	goto L454
L453:
	;
	goto L454
L454:
	;
	v1949 = v187
	v1950 = v188
	goto L15
L455:
	;
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
	v1543 = F_get_tablespace_oid(m, v1541, int32(0))
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L9
	} else {
		goto L457
	}
L456:
	;
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v117)+92))
	if v1562 != 0 {
		goto L463
	} else {
		goto L464
	}
L457:
	;
	if v1543 == int32(0) {
		goto L456
	} else {
		goto L458
	}
L458:
	;
	v1548 = *(*int32)(unsafe.Add(mBase, _c_F_ATPrepCmd[4]))
	if v1543 == v1548 {
		goto L456
	} else {
		goto L459
	}
L459:
	;
	v1552 = *(*int32)(unsafe.Add(mBase, _c_F_ATPrepCmd[0]))
	v1554 = F_object_aclcheck(m, int32(1213), v1543, v1552, int64(512))
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L9
	} else {
		goto L460
	}
L460:
	;
	if v1554 == int32(0) {
		goto L456
	} else {
		goto L461
	}
L461:
	;
	F_aclcheck_error(m, v1554, int32(42), v1541)
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L9
	} else {
		goto L462
	}
L462:
	;
	goto L456
L463:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L9
	} else {
		goto L466
	}
L464:
	;
	goto L465
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+92)) = v1543
	v1949 = v187
	v1950 = v188
	goto L15
L466:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L9
	} else {
		goto L467
	}
L467:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_53), int32(0))
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L9
	} else {
		goto L468
	}
L468:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_54), int32(_a_F_ATPrepCmd_55))
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L9
	} else {
		goto L469
	}
L469:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L470:
	;
	v1949 = v187
	v1950 = v188
	goto L15
L471:
	;
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(v1587)+76))
	if v1588 == int32(0) {
		goto L475
	} else {
		goto L476
	}
L472:
	;
	v1949 = v187
	v1950 = v188
	goto L15
L473:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L9
	} else {
		goto L488
	}
L474:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		goto L9
	} else {
		goto L484
	}
L475:
	;
	v1591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1587)+131)))
	if v1591 == int32(1) {
		goto L474
	} else {
		goto L478
	}
L476:
	;
	goto L477
L477:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L9
	} else {
		goto L480
	}
L478:
	;
	v1594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1587)+119)))
	if v1594 == int32(112) {
		goto L473
	} else {
		goto L479
	}
L479:
	;
	goto L472
L480:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L9
	} else {
		goto L481
	}
L481:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_56), int32(0))
	mBase = m.M
	v1607 = m.ExcPending
	if v1607 != 0 {
		goto L9
	} else {
		goto L482
	}
L482:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_57), int32(_a_F_ATPrepCmd_58))
	mBase = m.M
	v1612 = m.ExcPending
	if v1612 != 0 {
		goto L9
	} else {
		goto L483
	}
L483:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L484:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L9
	} else {
		goto L485
	}
L485:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_59), int32(0))
	mBase = m.M
	v1623 = m.ExcPending
	if v1623 != 0 {
		goto L9
	} else {
		goto L486
	}
L486:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_60), int32(_a_F_ATPrepCmd_58))
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L9
	} else {
		goto L487
	}
L487:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L488:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1635 = m.ExcPending
	if v1635 != 0 {
		goto L9
	} else {
		goto L489
	}
L489:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_61), int32(0))
	mBase = m.M
	v1639 = m.ExcPending
	if v1639 != 0 {
		goto L9
	} else {
		goto L490
	}
L490:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_62), int32(_a_F_ATPrepCmd_58))
	mBase = m.M
	v1644 = m.ExcPending
	if v1644 != 0 {
		goto L9
	} else {
		goto L491
	}
L491:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L492:
	;
	v1949 = v187
	v1950 = v188
	goto L15
L493:
	;
	if l3 == int32(0) {
		v1949 = v187
		v1950 = v188
		goto L15
	} else {
		goto L494
	}
L494:
	;
	v1655 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v188)+29)) = uint8(v1655)
	v1949 = v187
	v1950 = v188
	goto L15
L495:
	;
	if l3 == int32(0) {
		v1949 = v187
		v1950 = v188
		goto L15
	} else {
		goto L496
	}
L496:
	;
	v1663 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v188)+29)) = uint8(v1663)
	v1949 = v187
	v1950 = v188
	goto L15
L497:
	;
	v1949 = v187
	v1950 = v188
	goto L15
L498:
	;
	if l3 == int32(0) {
		v1949 = v187
		v1950 = v188
		goto L15
	} else {
		goto L499
	}
L499:
	;
	v1674 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v188)+29)) = uint8(v1674)
	v1949 = v187
	v1950 = v188
	goto L15
L500:
	;
	v1949 = v187
	v1950 = v188
	goto L15
L501:
	;
	v1949 = v187
	v1950 = v188
	goto L15
L502:
	;
	v1949 = v187
	v1950 = v188
	goto L15
L503:
	;
	v1949 = v187
	v1950 = v188
	goto L15
L504:
	;
	v1949 = v187
	v1950 = v188
	goto L15
L505:
	;
	v1699 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v1699
	F_errmsg_internal(m, int32(_a_F_ATPrepCmd_63), v23)
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		goto L9
	} else {
		goto L506
	}
L506:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_64), int32(_a_F_ATPrepCmd_65))
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L9
	} else {
		goto L507
	}
L507:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L508:
	;
	F_ATPrepAddColumn(m, l0, l1, l3, l4, int32(0), v188, l5, l6)
	mBase = m.M
	v1715 = m.ExcPending
	if v1715 != 0 {
		goto L9
	} else {
		goto L509
	}
L509:
	;
	v1949 = int32(2)
	v1950 = v188
	goto L15
L510:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1723 = m.ExcPending
	if v1723 != 0 {
		goto L9
	} else {
		goto L511
	}
L511:
	;
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v1724 + int32(4)
	F_errmsg(m, int32(_a_F_ATPrepCmd_66), v23+int32(80))
	mBase = m.M
	v1732 = m.ExcPending
	if v1732 != 0 {
		goto L9
	} else {
		goto L512
	}
L512:
	;
	F_errhint(m, int32(_a_F_ATPrepCmd_67), int32(0))
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L9
	} else {
		goto L513
	}
L513:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_68), int32(_a_F_ATPrepCmd_65))
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L9
	} else {
		goto L514
	}
L514:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L515:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L9
	} else {
		goto L516
	}
L516:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_69), int32(0))
	mBase = m.M
	v1752 = m.ExcPending
	if v1752 != 0 {
		goto L9
	} else {
		goto L517
	}
L517:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_70), int32(_a_F_ATPrepCmd_65))
	mBase = m.M
	v1757 = m.ExcPending
	if v1757 != 0 {
		goto L9
	} else {
		goto L518
	}
L518:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L519:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1764 = m.ExcPending
	if v1764 != 0 {
		goto L9
	} else {
		goto L520
	}
L520:
	;
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v1765 + int32(4)
	F_errmsg(m, int32(_a_F_ATPrepCmd_71), v23+int32(48))
	mBase = m.M
	v1773 = m.ExcPending
	if v1773 != 0 {
		goto L9
	} else {
		goto L521
	}
L521:
	;
	F_errdetail(m, int32(_a_F_ATPrepCmd_72), int32(0))
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L9
	} else {
		goto L522
	}
L522:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_73), int32(_a_F_ATPrepCmd_52))
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L9
	} else {
		goto L523
	}
L523:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L524:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1789 = m.ExcPending
	if v1789 != 0 {
		goto L9
	} else {
		goto L525
	}
L525:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_74), int32(0))
	mBase = m.M
	v1793 = m.ExcPending
	if v1793 != 0 {
		goto L9
	} else {
		goto L526
	}
L526:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_75), int32(_a_F_ATPrepCmd_65))
	mBase = m.M
	v1798 = m.ExcPending
	if v1798 != 0 {
		goto L9
	} else {
		goto L527
	}
L527:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L528:
	;
	v1808 = v23 + int32(96)
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_ScanKeyInit(m, v1808, v1802, int32(3), int32(184), v1811)
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L9
	} else {
		goto L529
	}
L529:
	;
	v1814 = int32(1)
	v1817 = F_systable_beginscan(m, v1805, v1801, v1814, int32(0), v1814, v1808)
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L9
	} else {
		goto L530
	}
L530:
	;
	v1819 = F_systable_getnext(m, v1817)
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L9
	} else {
		goto L531
	}
L531:
	;
	if v1819 != 0 {
		goto L532
	} else {
		goto L533
	}
L532:
	;
	if v1463 == int32(29) {
		goto L535
	} else {
		goto L536
	}
L533:
	;
	goto L534
L534:
	;
	F_systable_endscan(m, v1817)
	mBase = m.M
	v1924 = m.ExcPending
	if v1924 != 0 {
		goto L9
	} else {
		goto L558
	}
L535:
	;
	v1825 = int32(96)
	goto L537
L536:
	;
	v1825 = int32(80)
	goto L537
L537:
	;
	v1830 = v1819
	goto L538
L538:
	;
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1830)+16))
	v1847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1846)+22)))
	v1848 = v1846 + v1847
	v1849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1848)+72)))
	if v1849 != int32(102) {
		goto L540
	} else {
		goto L541
	}
L539:
	;
	goto L534
L540:
	;
	v1901 = F_systable_getnext(m, v1817)
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L9
	} else {
		goto L556
	}
L541:
	;
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(v1848+v1825)))
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v1853 == v1854 {
		goto L540
	} else {
		goto L542
	}
L542:
	;
	v1857 = F_relation_open(m, v1853, int32(1))
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L9
	} else {
		goto L543
	}
L543:
	;
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1857)+48))
	v1860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1859)+118)))
	if v1463 == int32(29) {
		goto L545
	} else {
		goto L546
	}
L544:
	;
	F_relation_close(m, v1857, int32(1))
	mBase = m.M
	v1898 = m.ExcPending
	if v1898 != 0 {
		goto L9
	} else {
		goto L555
	}
L545:
	;
	if v1860 == int32(112) {
		goto L544
	} else {
		goto L548
	}
L546:
	;
	goto L547
L547:
	;
	if v1860 == int32(112) {
		goto L14
	} else {
		goto L554
	}
L548:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		goto L9
	} else {
		goto L549
	}
L549:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1871 = m.ExcPending
	if v1871 != 0 {
		goto L9
	} else {
		goto L550
	}
L550:
	;
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(v1857)+48))
	v1874 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v1873 + v1874
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v1872 + v1874
	F_errmsg(m, int32(_a_F_ATPrepCmd_76), v23+int32(16))
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		goto L9
	} else {
		goto L551
	}
L551:
	;
	F_errtableconstraint(m, l1, v1848+int32(4))
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L9
	} else {
		goto L552
	}
L552:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_77), int32(_a_F_ATPrepCmd_52))
	mBase = m.M
	v1893 = m.ExcPending
	if v1893 != 0 {
		goto L9
	} else {
		goto L553
	}
L553:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L554:
	;
	goto L544
L555:
	;
	goto L540
L556:
	;
	if v1901 != 0 {
		v1830 = v1901
		goto L538
	} else {
		goto L557
	}
L557:
	;
	goto L539
L558:
	;
	F_relation_close(m, v1805, int32(1))
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L9
	} else {
		goto L559
	}
L559:
	;
	if v1463 == int32(29) {
		goto L560
	} else {
		goto L561
	}
L560:
	;
	v1932 = int32(112)
	goto L562
L561:
	;
	v1932 = int32(117)
	goto L562
L562:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v117)+97)) = uint8(v1932)
	v1934 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v117)+96)) = uint8(v1934)
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v117)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+80)) = v1936 | v1934
	v1949 = v187
	v1950 = v188
	goto L15
L563:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1962)+16)) = v1964
	m.G0 = v23 + int32(144)
	return
L564:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L9
	} else {
		goto L565
	}
L565:
	;
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(v1857)+48))
	v1979 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v1978 + v1979
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v1977 + v1979
	F_errmsg(m, int32(_a_F_ATPrepCmd_78), v23+int32(32))
	mBase = m.M
	v1989 = m.ExcPending
	if v1989 != 0 {
		goto L9
	} else {
		goto L566
	}
L566:
	;
	F_errtableconstraint(m, l1, v1848+int32(4))
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L9
	} else {
		goto L567
	}
L567:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_79), int32(_a_F_ATPrepCmd_52))
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L9
	} else {
		goto L568
	}
L568:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L569:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2025 = m.ExcPending
	if v2025 != 0 {
		goto L9
	} else {
		goto L570
	}
L570:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_80), int32(0))
	mBase = m.M
	v2029 = m.ExcPending
	if v2029 != 0 {
		goto L9
	} else {
		goto L571
	}
L571:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_5), int32(_a_F_ATPrepCmd_81), int32(_a_F_ATPrepCmd_82))
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L9
	} else {
		goto L572
	}
L572:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
