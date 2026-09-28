package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ATExecSetOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v26 int32
	_ = v26
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	v12 = m.G0
	v14 = v12 - int32(304)
	m.G0 = v14
	v18 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
		v21 = F_SearchSysCacheAttName(m, v20, l2)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			if v21 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
				v25 = v23 + v24
				v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+74)))
				if v26 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v132 = m.ExcPending
						if v132 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l2
							F_errmsg(m, int32(_a_F_ATExecSetOptions_0), v14+int32(16))
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ATExecSetOptions_1), int32(_a_F_ATExecSetOptions_2), int32(_a_F_ATExecSetOptions_3))
								mBase = m.M
								v143 = m.ExcPending
								if v143 != 0 {
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
					v34 = F_SysCacheGetAttr(m, int32(6), v21, int32(23), v14+int32(303))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+303)))
						if v36 != 0 {
							v37 = int64(0)
						} else {
							v37 = v34
						}
						v38 = int32(0)
						v41 = F_transformRelOptions(m, v37, l3, v38, v38, v38, l4)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							v44 = F_attribute_reloptions(m, v41, int32(1))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								v46 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v14)+88)) = uint8(v46)
								v48 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v48
								*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v48
								*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v48
								*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v48
								*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v48
								*(*int64)(unsafe.Add(mBase, uint32(v14)+48)) = v48
								*(*uint8)(unsafe.Add(mBase, uint32(v14)+56)) = uint8(v46)
								if v41 != v48 {
									*(*int64)(unsafe.Add(mBase, uint32(v14)+272)) = v41
								} else {
									v65 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v14)+86)) = uint8(v65)
								}
								v67 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v14)+54)) = uint8(v67)
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
								v76 = F_heap_modify_tuple(m, v21, v69, v14+int32(96), v14-int32(-64), v14+int32(32))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return
								} else {
									F_CatalogTupleUpdate(m, v18, v76+int32(4), v76)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										v83 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecSetOptions[0]))
										if v83 != 0 {
											v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
											v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+74)))
											v87 = int32(0)
											F_RunObjectPostAlterHook(m, int32(1259), v85, v86, v87, v87)
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
												v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v26
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
												F_pfree(m, v76)
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return
												} else {
													F_ReleaseCatCache(m, v21)
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return
													} else {
														F_relation_close(m, v18, int32(3))
														mBase = m.M
														v102 = m.ExcPending
														if v102 != 0 {
															return
														} else {
															m.G0 = v14 + int32(304)
															return
														}
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
											v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v26
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
											F_pfree(m, v76)
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
											} else {
												F_ReleaseCatCache(m, v21)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return
												} else {
													F_relation_close(m, v18, int32(3))
													mBase = m.M
													v102 = m.ExcPending
													if v102 != 0 {
														return
													} else {
														m.G0 = v14 + int32(304)
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
				v109 = m.ExcPending
				if v109 != 0 {
					return
				} else {
					F_errcode(m, int32(50360452))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return
					} else {
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v113 + int32(4)
						F_errmsg(m, int32(_a_F_ATExecSetOptions_4), v14)
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ATExecSetOptions_1), int32(_a_F_ATExecSetOptions_5), int32(_a_F_ATExecSetOptions_3))
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
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
func F_ATPostAlterTypeParse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v234 int32
	_ = v234
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v386 int64
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v454 int32
	_ = v454
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v652 int32
	_ = v652
	var v665 int32
	_ = v665
	v8 = int32(0)
	v20 = m.G0
	v22 = v20 + int32(-64)
	m.G0 = v22
	v25 = F_raw_parser(m, l4, v8)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_relation_close(m, v652, int32(0))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L2
	} else {
		goto L136
	}
L2:
	;
	return
L3:
	;
	if v25 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v30 = F_relation_open(m, l1, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if int32(0) < v32 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v652 = v30
	goto L1
L8:
	;
	v42 = v8
	v47 = v8
	goto L11
L9:
	;
	v108 = v8
	goto L10
L10:
	;
	v116 = F_relation_open(m, l1, int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L2
	} else {
		goto L30
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54+v42<<(uint(int32(2))%32))))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	switch v60 - int32(204) {
	case 0:
		goto L16
	case 1:
		goto L15
	default:
		goto L14
	}
L12:
	;
	v108 = v91
	goto L10
L13:
	;
	v93 = v42 + int32(1)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v93 < v94 {
		v42 = v93
		v47 = v91
		goto L11
	} else {
		goto L29
	}
L14:
	;
	if v60 != int32(146) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v67 = F_transformStatsStmt(m, l1, v59, l4)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L19
	}
L16:
	;
	v63 = F_transformIndexStmt(m, l1, v59, l4)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v65 = F_lappend(m, v47, v63)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v91 = v65
	goto L13
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+28)) = l3
	v70 = F_lappend(m, v47, v67)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v91 = v70
	goto L13
L21:
	;
	v74 = F_lappend(m, v47, v59)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v80 = F_transformAlterTableStmt(m, l1, v59, l4, v20+int32(-4), v20+int32(-8))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L25
	}
L24:
	;
	v91 = v74
	goto L13
L25:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v22)+60))
	v83 = F_list_concat(m, v47, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v85 = F_lappend(m, v83, v80)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v88 = F_list_concat(m, v85, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v91 = v88
	goto L13
L29:
	;
	goto L12
L30:
	;
	if v108 == int32(0) {
		v652 = v116
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v120 <= int32(0) {
		v652 = v116
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v138 = v8
	goto L35
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L2
	} else {
		goto L133
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L2
	} else {
		goto L130
	}
L35:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v143+v138<<(uint(int32(2))%32))))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v116)+56))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v149 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L2
	} else {
		goto L127
	}
L37:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	switch v246 - int32(146) {
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
	v205 = F_palloc0(m, int32(144))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L2
	} else {
		goto L45
	}
L39:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v152 <= int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
	v160 = int32(0)
	goto L41
L41:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v155+v160<<(uint(int32(2))%32))))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	if v180 == v148 {
		v234 = v179
		goto L37
	} else {
		goto L43
	}
L42:
	;
	goto L38
L43:
	;
	v183 = v160 + int32(1)
	if v152 != v183 {
		v160 = v183
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v205)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = v148
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v116)+48))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v205)+4)) = uint8(v211)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v116)+52))
	v214 = F_CreateTupleDescCopyConstr(m, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v205)+8)) = v214
	*(*int64)(unsafe.Add(mBase, uint32(v205)+88)) = int64(0)
	v219 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v205)+84)) = uint8(v219)
	v221 = int32(_a_F_ATPostAlterTypeParse_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v205)+96)) = uint16(v221)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v224 = F_lappend(m, v223, v205)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v224
	v234 = v205
	goto L37
L48:
	;
	goto L36
L49:
	;
	v597 = v138 + int32(1)
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v598 <= v597 {
		v652 = v116
		goto L1
	} else {
		goto L126
	}
L50:
	;
	v564 = F_GetComment(m, l0, int32(3381), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L2
	} else {
		goto L123
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L2
	} else {
		goto L120
	}
L52:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	if v512 == int32(67) {
		goto L111
	} else {
		goto L112
	}
L53:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v147)+8))
	if v293 == int32(0) {
		goto L49
	} else {
		goto L68
	}
L54:
	;
	switch v246 - int32(204) {
	case 0:
		goto L55
	case 1:
		goto L50
	default:
		goto L51
	}
L55:
	;
	if l6 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v276 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v147)+70)) = uint8(v276)
	v280 = F_GetComment(m, l0, int32(1259), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L2
	} else {
		goto L65
	}
L57:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v147)+12))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v147)+20))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v147)+36))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+64)))
	v255 = F_CheckIndexCompatible(m, l0, v251, v252, v253, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	if v255 == int32(0) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v260 = F_index_open(m, l0, int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v260)+48))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+119)))
	if v263 != int32(73) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+48)) = v266
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v260)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+52)) = v268
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v260)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+56)) = v270
	goto L63
L62:
	;
	goto L63
L63:
	;
	F_relation_close(m, v260, int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	goto L56
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+40)) = v280
	v284 = F_palloc0(m, int32(32))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v284)+20)) = v147
	*(*int64)(unsafe.Add(mBase, uint32(v284))) = int64(64424509587)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v234)+32))
	v290 = F_lappend(m, v289, v284)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+32)) = v290
	goto L49
L68:
	;
	v296 = int32(0)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	if v297 <= v296 {
		goto L49
	} else {
		goto L69
	}
L69:
	;
	v301 = v296
	goto L70
L70:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v293)+12))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v319+v301<<(uint(int32(2))%32))))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	switch v324 - int32(14) {
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
	v509 = v301 + int32(1)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	if v509 < v510 {
		v301 = v509
		goto L70
	} else {
		goto L110
	}
L73:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v323)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v373)+100)) = l2
	if l6 != 0 {
		goto L88
	} else {
		goto L89
	}
L74:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v323)+20))
	v328 = F_get_constraint_index(m, l0)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	if l6 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v357 = F_GetComment(m, v328, int32(1259), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L2
	} else {
		goto L85
	}
L77:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v327)+12))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v327)+20))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v327)+36))
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327)+64)))
	v334 = F_CheckIndexCompatible(m, v328, v330, v331, v332, v333)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	if v334 == int32(0) {
		goto L76
	} else {
		goto L79
	}
L79:
	;
	v339 = F_index_open(m, v328, int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v339)+48))
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341)+119)))
	if v342 != int32(73) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v339)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v327)+48)) = v345
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v339)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v327)+52)) = v347
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v339)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v327)+56)) = v349
	goto L83
L82:
	;
	goto L83
L83:
	;
	F_relation_close(m, v339, int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	goto L76
L85:
	;
	v359 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v327)+70)) = uint8(v359)
	*(*int32)(unsafe.Add(mBase, uint32(v327)+40)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v323)+4)) = int32(15)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v234)+32))
	v365 = F_lappend(m, v364, v323)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L2
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+32)) = v365
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	F_RebuildConstraintComment(m, v234, int32(4), l0, v116, int32(0), v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L2
	} else {
		goto L87
	}
L87:
	;
	goto L72
L88:
	;
	v474 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v373)+60)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v323)+4)) = int32(17)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v234)+36))
	v479 = F_lappend(m, v478, v323)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L2
	} else {
		goto L107
	}
L89:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	if v375 != int32(9) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v234)+80))
	if v378 != 0 {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	v380 = F_SearchSysCache1(m, int32(19), base.I64_extend_i32_u(l0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L2
	} else {
		goto L92
	}
L92:
	;
	if v380 == int32(0) {
		goto L48
	} else {
		goto L93
	}
L93:
	;
	v386 = F_SysCacheGetAttrNotNull(m, int32(19), v380, int32(23))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v389 = F_pg_detoast_datum(m, base.I32_wrap_i64(v386))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	if v391 != int32(1) {
		goto L34
	} else {
		goto L96
	}
L96:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v389)+8))
	if v394 != 0 {
		goto L34
	} else {
		goto L97
	}
L97:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v389)+12))
	if v395 != int32(26) {
		goto L34
	} else {
		goto L98
	}
L98:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v389)+16))
	if int32(0) < v398 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v373)+96))
	v408 = int32(0)
	v409 = v403
	goto L102
L100:
	;
	goto L101
L101:
	;
	F_ReleaseCatCache(m, v380)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L2
	} else {
		goto L106
	}
L102:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v389+int32(24)+v408<<(uint(int32(2))%32))))
	v428 = F_lappend_oid(m, v409, v427)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L2
	} else {
		goto L104
	}
L103:
	;
	goto L101
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v373)+96)) = v428
	v432 = v408 + int32(1)
	if v432 != v398 {
		v408 = v432
		v409 = v428
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
	*(*int32)(unsafe.Add(mBase, uint32(v234)+36)) = v479
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v373)+8))
	if v482 == int32(0) {
		goto L72
	} else {
		goto L108
	}
L108:
	;
	F_RebuildConstraintComment(m, v234, int32(5), l0, v116, int32(0), v482)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
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
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v147)+16))
	v517 = F_palloc0(m, int32(32))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
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
	v535 = m.ExcPending
	if v535 != 0 {
		goto L2
	} else {
		goto L117
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v517)+20)) = v147
	*(*int64)(unsafe.Add(mBase, uint32(v517))) = int64(77309411475)
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v234)+36))
	v523 = F_lappend(m, v522, v517)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L2
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+36)) = v523
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v147)+8))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v515)+8))
	F_RebuildConstraintComment(m, v234, int32(5), l0, int32(0), v528, v529)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L2
	} else {
		goto L116
	}
L116:
	;
	goto L49
L117:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v536
	F_errmsg_internal(m, int32(_a_F_ATPostAlterTypeParse_1), v20+int32(-16))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_ATPostAlterTypeParse_2), int32(_a_F_ATPostAlterTypeParse_3), int32(_a_F_ATPostAlterTypeParse_4))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
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
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v552
	F_errmsg_internal(m, int32(_a_F_ATPostAlterTypeParse_5), v22)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_ATPostAlterTypeParse_2), int32(_a_F_ATPostAlterTypeParse_6), int32(_a_F_ATPostAlterTypeParse_4))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v147)+20)) = v564
	v568 = F_palloc0(m, int32(32))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L2
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v568)+20)) = v147
	*(*int64)(unsafe.Add(mBase, uint32(v568))) = int64(279172874387)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v234)+60))
	v574 = F_lappend(m, v573, v568)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L2
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+60)) = v574
	goto L49
L126:
	;
	v138 = v597
	goto L35
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = l0
	F_errmsg_internal(m, int32(_a_F_ATPostAlterTypeParse_7), v20+int32(-32))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L2
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_ATPostAlterTypeParse_2), int32(_a_F_ATPostAlterTypeParse_8), int32(_a_F_ATPostAlterTypeParse_9))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
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
	v622 = m.ExcPending
	if v622 != 0 {
		goto L2
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_ATPostAlterTypeParse_2), int32(_a_F_ATPostAlterTypeParse_11), int32(_a_F_ATPostAlterTypeParse_9))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
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
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v632
	F_errmsg_internal(m, int32(_a_F_ATPostAlterTypeParse_1), v20+int32(-48))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L2
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_ATPostAlterTypeParse_2), int32(_a_F_ATPostAlterTypeParse_12), int32(_a_F_ATPostAlterTypeParse_4))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
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
	m.G0 = v22 - int32(-64)
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
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
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
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v604 int32
	_ = v604
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v665 int32
	_ = v665
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
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
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v838 int32
	_ = v838
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v871 int32
	_ = v871
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v916 int32
	_ = v916
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v959 int32
	_ = v959
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1011 int32
	_ = v1011
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1188 int32
	_ = v1188
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1209 int32
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1234 int32
	_ = v1234
	var v1238 int32
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1255 int32
	_ = v1255
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1280 int32
	_ = v1280
	var v1287 int32
	_ = v1287
	var v1291 int32
	_ = v1291
	var v1296 int32
	_ = v1296
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1312 int32
	_ = v1312
	var v1317 int32
	_ = v1317
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1334 int32
	_ = v1334
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1356 int32
	_ = v1356
	var v1361 int32
	_ = v1361
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1372 int32
	_ = v1372
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1388 int32
	_ = v1388
	var v1394 int32
	_ = v1394
	var v1399 int32
	_ = v1399
	var v1404 int32
	_ = v1404
	var v1407 int32
	_ = v1407
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1441 int32
	_ = v1441
	var v1450 int64
	_ = v1450
	var v1451 int64
	_ = v1451
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1465 int32
	_ = v1465
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1508 int32
	_ = v1508
	var v1512 int32
	_ = v1512
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1533 int32
	_ = v1533
	var v1538 int32
	_ = v1538
	var v1542 int32
	_ = v1542
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1566 int32
	_ = v1566
	var v1569 int32
	_ = v1569
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1580 int32
	_ = v1580
	var v1584 int32
	_ = v1584
	var v1588 int32
	_ = v1588
	var v1592 int32
	_ = v1592
	var v1596 int32
	_ = v1596
	var v1600 int32
	_ = v1600
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1609 int32
	_ = v1609
	var v1614 int32
	_ = v1614
	var v1618 int32
	_ = v1618
	var v1621 int32
	_ = v1621
	var v1626 int32
	_ = v1626
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1638 int32
	_ = v1638
	var v1642 int32
	_ = v1642
	var v1647 int32
	_ = v1647
	var v1651 int32
	_ = v1651
	var v1654 int32
	_ = v1654
	var v1658 int32
	_ = v1658
	var v1663 int32
	_ = v1663
	var v1667 int32
	_ = v1667
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1679 int32
	_ = v1679
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1687 int32
	_ = v1687
	var v1692 int32
	_ = v1692
	var v1696 int32
	_ = v1696
	var v1699 int32
	_ = v1699
	var v1703 int32
	_ = v1703
	var v1708 int32
	_ = v1708
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1722 int64
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1736 int32
	_ = v1736
	var v1741 int32
	_ = v1741
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1779 int32
	_ = v1779
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1795 int32
	_ = v1795
	var v1799 int32
	_ = v1799
	var v1804 int32
	_ = v1804
	var v1809 int32
	_ = v1809
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1843 int32
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1847 int32
	_ = v1847
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1884 int32
	_ = v1884
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1900 int32
	_ = v1900
	var v1904 int32
	_ = v1904
	var v1909 int32
	_ = v1909
	var v1933 int32
	_ = v1933
	var v1936 int32
	_ = v1936
	var v1940 int32
	_ = v1940
	var v1945 int32
	_ = v1945
	v8 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(160)
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
	v84 = F_palloc0(m, int32(144))
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
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L9
	} else {
		goto L541
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		goto L9
	} else {
		goto L536
	}
L15:
	;
	v1873 = v117 + v1860<<(uint(int32(2))%32)
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v1873)+16))
	v1875 = F_lappend(m, v1874, v1861)
	mBase = m.M
	v1876 = m.ExcPending
	if v1876 != 0 {
		goto L9
	} else {
		goto L535
	}
L16:
	;
	v1716 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L9
	} else {
		goto L500
	}
L17:
	;
	v1712 = int32(2665)
	v1713 = int32(9)
	goto L16
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L9
	} else {
		goto L496
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1667 = m.ExcPending
	if v1667 != 0 {
		goto L9
	} else {
		goto L490
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L9
	} else {
		goto L486
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1626 = m.ExcPending
	if v1626 != 0 {
		goto L9
	} else {
		goto L481
	}
L22:
	;
	v136 = int32(11)
	v137 = F_copyObjectImpl(m, l2)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L9
	} else {
		goto L67
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
	v134 = F_PartitionHasPendingDetach(m, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	if v134 != 0 {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	F_ATSimplePermissions(m, int32(0), l1, int32(305))
	mBase = m.M
	v1618 = m.ExcPending
	if v1618 != 0 {
		goto L9
	} else {
		goto L479
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L9
	} else {
		goto L476
	}
L29:
	;
	F_ATSimplePermissions(m, int32(61), l1, int32(256))
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L9
	} else {
		goto L475
	}
L30:
	;
	F_ATSimplePermissions(m, int32(60), l1, int32(256))
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L9
	} else {
		goto L474
	}
L31:
	;
	F_ATSimplePermissions(m, int32(59), l1, int32(320))
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		goto L9
	} else {
		goto L473
	}
L32:
	;
	F_ATSimplePermissions(m, int32(58), l1, int32(32))
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L9
	} else {
		goto L472
	}
L33:
	;
	F_ATSimplePermissions(m, v139, l1, int32(257))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L9
	} else {
		goto L471
	}
L34:
	;
	F_ATSimplePermissions(m, v139, l1, int32(289))
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		goto L9
	} else {
		goto L469
	}
L35:
	;
	F_ATSimplePermissions(m, int32(53), l1, int32(261))
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L9
	} else {
		goto L468
	}
L36:
	;
	F_ATSimplePermissions(m, int32(20), l1, int32(289))
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L9
	} else {
		goto L466
	}
L37:
	;
	F_ATSimplePermissions(m, int32(19), l1, int32(257))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L9
	} else {
		goto L464
	}
L38:
	;
	F_ATSimplePermissions(m, int32(50), l1, int32(33))
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L9
	} else {
		goto L462
	}
L39:
	;
	F_ATSimplePermissions(m, int32(49), l1, int32(33))
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L9
	} else {
		goto L460
	}
L40:
	;
	F_ATSimplePermissions(m, v139, l1, int32(271))
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L9
	} else {
		goto L459
	}
L41:
	;
	F_ATSimplePermissions(m, int32(33), l1, int32(333))
	mBase = m.M
	v1500 = m.ExcPending
	if v1500 != 0 {
		goto L9
	} else {
		goto L444
	}
L42:
	;
	F_ATSimplePermissions(m, int32(32), l1, int32(261))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L9
	} else {
		goto L433
	}
L43:
	;
	F_ATSimplePermissions(m, int32(31), l1, int32(289))
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L9
	} else {
		goto L432
	}
L44:
	;
	F_ATSimplePermissions(m, v139, l1, int32(129))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L9
	} else {
		goto L414
	}
L45:
	;
	F_ATSimplePermissions(m, v139, l1, int32(5))
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L9
	} else {
		goto L413
	}
L46:
	;
	F_ATSimplePermissions(m, int32(25), l1, int32(32))
	mBase = m.M
	v1404 = m.ExcPending
	if v1404 != 0 {
		goto L9
	} else {
		goto L412
	}
L47:
	;
	F_ATSimplePermissions(m, int32(24), l1, int32(305))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L9
	} else {
		goto L206
	}
L48:
	;
	F_ATSimplePermissions(m, int32(22), l1, int32(289))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L9
	} else {
		goto L185
	}
L49:
	;
	F_ATSimplePermissions(m, int32(21), l1, int32(257))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L9
	} else {
		goto L184
	}
L50:
	;
	F_ATSimplePermissions(m, int32(16), l1, int32(289))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L9
	} else {
		goto L142
	}
L51:
	;
	F_ATSimplePermissions(m, int32(14), l1, int32(257))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L9
	} else {
		goto L141
	}
L52:
	;
	F_ATSimplePermissions(m, int32(13), l1, int32(305))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L9
	} else {
		goto L123
	}
L53:
	;
	F_ATSimplePermissions(m, int32(12), l1, int32(261))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L9
	} else {
		goto L122
	}
L54:
	;
	F_ATSimplePermissions(m, int32(11), l1, int32(293))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L9
	} else {
		goto L120
	}
L55:
	;
	F_ATSimplePermissions(m, v139, l1, int32(293))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L9
	} else {
		goto L119
	}
L56:
	;
	F_ATSimplePermissions(m, int32(8), l1, int32(365))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L9
	} else {
		goto L117
	}
L57:
	;
	F_ATSimplePermissions(m, int32(7), l1, int32(289))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L9
	} else {
		goto L88
	}
L58:
	;
	F_ATSimplePermissions(m, int32(6), l1, int32(289))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L9
	} else {
		goto L86
	}
L59:
	;
	F_ATSimplePermissions(m, int32(5), l1, int32(289))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L9
	} else {
		goto L84
	}
L60:
	;
	F_ATSimplePermissions(m, int32(4), l1, int32(289))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L9
	} else {
		goto L82
	}
L61:
	;
	F_ATSimplePermissions(m, int32(64), l1, int32(291))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L9
	} else {
		goto L80
	}
L62:
	;
	F_ATSimplePermissions(m, int32(63), l1, int32(291))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L9
	} else {
		goto L78
	}
L63:
	;
	F_ATSimplePermissions(m, int32(62), l1, int32(291))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L9
	} else {
		goto L76
	}
L64:
	;
	F_ATSimplePermissions(m, int32(3), l1, int32(289))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L9
	} else {
		goto L75
	}
L65:
	;
	F_ATSimplePermissions(m, int32(2), l1, int32(291))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L9
	} else {
		goto L70
	}
L66:
	;
	v140 = int32(2)
	F_ATSimplePermissions(m, int32(1), l1, v140)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L9
	} else {
		goto L68
	}
L67:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	switch v139 {
	case 0:
		goto L27
	case 1:
		goto L66
	case 2:
		goto L65
	case 3:
		goto L64
	case 4:
		goto L60
	case 5:
		goto L59
	case 6:
		goto L58
	case 7:
		goto L57
	case 8:
		goto L56
	case 9, 10:
		goto L55
	case 11:
		goto L54
	case 12:
		goto L53
	case 13:
		goto L52
	case 14:
		goto L51
	default:
		goto L28
	case 16:
		goto L50
	case 19:
		goto L37
	case 20:
		goto L36
	case 21:
		goto L49
	case 22:
		goto L48
	case 24:
		goto L47
	case 25:
		goto L46
	case 26:
		v1860 = v136
		v1861 = v137
		goto L15
	case 27, 28:
		goto L45
	case 29, 30:
		goto L44
	case 31:
		goto L43
	case 32:
		goto L42
	case 33:
		goto L41
	case 34, 35, 36:
		goto L40
	case 37, 38, 39, 40, 41, 42, 43, 44:
		goto L34
	case 45, 46, 47, 48, 51, 52, 54, 55, 56, 57:
		goto L33
	case 49:
		goto L39
	case 50:
		goto L38
	case 53:
		goto L35
	case 58:
		goto L32
	case 59:
		goto L31
	case 60:
		goto L30
	case 61:
		goto L29
	case 62:
		goto L63
	case 63:
		goto L62
	case 64:
		goto L61
	}
L68:
	;
	F_ATPrepAddColumn(m, l0, l1, l3, l4, int32(1), v137, l5, l6)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	v1860 = v140
	v1861 = v137
	goto L15
L70:
	;
	F_ATSimpleRecursion(m, l0, l1, v137, l3, l5, l6)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
	if v156 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v157 = int32(10)
	goto L74
L73:
	;
	v157 = int32(0)
	goto L74
L74:
	;
	v1860 = v157
	v1861 = v137
	goto L15
L75:
	;
	v1860 = int32(10)
	v1861 = v137
	goto L15
L76:
	;
	v167 = int32(10)
	if l3 == int32(0) {
		v1860 = v167
		v1861 = v137
		goto L15
	} else {
		goto L77
	}
L77:
	;
	v170 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+29)) = uint8(v170)
	v1860 = v167
	v1861 = v137
	goto L15
L78:
	;
	if l3 == int32(0) {
		v1860 = v136
		v1861 = v137
		goto L15
	} else {
		goto L79
	}
L79:
	;
	v178 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+29)) = uint8(v178)
	v1860 = v136
	v1861 = v137
	goto L15
L80:
	;
	v184 = int32(0)
	if l3 == v184 {
		v1860 = v184
		v1861 = v137
		goto L15
	} else {
		goto L81
	}
L81:
	;
	v187 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+29)) = uint8(v187)
	v1860 = v184
	v1861 = v137
	goto L15
L82:
	;
	v193 = int32(0)
	if l3 == v193 {
		v1860 = v193
		v1861 = v137
		goto L15
	} else {
		goto L83
	}
L83:
	;
	v196 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+29)) = uint8(v196)
	v1860 = v193
	v1861 = v137
	goto L15
L84:
	;
	v202 = int32(7)
	if l3 == int32(0) {
		v1860 = v202
		v1861 = v137
		goto L15
	} else {
		goto L85
	}
L85:
	;
	v205 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+29)) = uint8(v205)
	v1860 = v202
	v1861 = v137
	goto L15
L86:
	;
	F_ATSimpleRecursion(m, l0, l1, v137, l3, l5, l6)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L87
	}
L87:
	;
	v1860 = int32(3)
	v1861 = v137
	goto L15
L88:
	;
	F_ATSimpleRecursion(m, l0, l1, v137, l3, l5, l6)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L9
	} else {
		goto L89
	}
L89:
	;
	v220 = m.G0
	v222 = v220 - int32(16)
	m.G0 = v222
	if l3|l4 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L90:
	;
	v1860 = int32(0)
	v1861 = v137
	goto L15
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L9
	} else {
		goto L113
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L9
	} else {
		goto L109
	}
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L9
	} else {
		goto L105
	}
L94:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v229 = F_find_inheritance_children(m, v227, int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L9
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	if l4 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	if v229 != 0 {
		goto L93
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	v235 = F_SearchSysCacheCopyAttName(m, v233, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L9
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	m.G0 = v222 + int32(16)
	goto L90
L102:
	;
	if v235 == int32(0) {
		goto L92
	} else {
		goto L103
	}
L103:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v235)+16))
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+22)))
	v242 = int32(*(*int16)(unsafe.Add(mBase, uint32(v239+v240)+94)))
	if int32(0) < v242 {
		goto L91
	} else {
		goto L104
	}
L104:
	;
	goto L101
L105:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L9
	} else {
		goto L106
	}
L106:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_1), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L9
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_3), int32(_a_F_ATPrepCmd_4))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L9
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L9
	} else {
		goto L110
	}
L110:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v222))) = v274
	*(*int32)(unsafe.Add(mBase, uint32(v222)+4)) = v273 + int32(4)
	F_errmsg(m, int32(_a_F_ATPrepCmd_5), v222)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L9
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_6), int32(_a_F_ATPrepCmd_4))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L9
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L9
	} else {
		goto L114
	}
L114:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_7), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L9
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_8), int32(_a_F_ATPrepCmd_4))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L9
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	F_ATSimpleRecursion(m, l0, l1, v137, l3, l5, l6)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L9
	} else {
		goto L118
	}
L118:
	;
	v1860 = v136
	v1861 = v137
	goto L15
L119:
	;
	v1860 = v136
	v1861 = v137
	goto L15
L120:
	;
	F_ATSimpleRecursion(m, l0, l1, v137, l3, l5, l6)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L9
	} else {
		goto L121
	}
L121:
	;
	v1860 = v136
	v1861 = v137
	goto L15
L122:
	;
	v1860 = v136
	v1861 = v137
	goto L15
L123:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if l4 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	v1860 = int32(0)
	v1861 = v137
	goto L15
L125:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L9
	} else {
		goto L137
	}
L126:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v327)+76))
	if v330 != 0 {
		goto L125
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327)+119)))
	if v331 == int32(99) {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	goto L128
L130:
	;
	F_ATTypedTableRecursion(m, l0, l1, v137, l5, l6)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L9
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	if l3 != 0 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	goto L132
L134:
	;
	v336 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+29)) = uint8(v336)
	goto L136
L135:
	;
	goto L136
L136:
	;
	goto L124
L137:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L9
	} else {
		goto L138
	}
L138:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_9), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L9
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_10), int32(_a_F_ATPrepCmd_11))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L9
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	v1860 = int32(9)
	v1861 = v137
	goto L15
L142:
	;
	v364 = int32(0)
	v366 = m.G0
	v368 = v366 - int32(16)
	m.G0 = v368
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	if v371 != int32(6) {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	v567 = int32(6)
	if l3 == int32(0) {
		v1860 = v567
		v1861 = v137
		goto L15
	} else {
		goto L182
	}
L144:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L9
	} else {
		goto L178
	}
L145:
	;
	m.G0 = v368 + int32(16)
	goto L143
L146:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v370)+32))
	if v374 == int32(0) {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	if v377 <= int32(0) {
		goto L145
	} else {
		goto L148
	}
L148:
	;
	v384 = v364
	v389 = v364
	v395 = v8
	goto L149
L149:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v374)+12))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v401+v395<<(uint(int32(2))%32))))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	v407 = F_findNotNullConstraint(m, v400, v406)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L9
	} else {
		goto L152
	}
L150:
	;
	goto L145
L151:
	;
	v524 = v395 + int32(1)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	if v524 < v525 {
		v384 = v507
		v389 = v512
		v395 = v524
		goto L149
	} else {
		goto L177
	}
L152:
	;
	if v407 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	F_verifyNotNullPKCompatible(m, v407, v409)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L9
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	if l3 != 0 {
		v473 = v384
		v478 = v389
		goto L158
	} else {
		goto L159
	}
L156:
	;
	F_pfree(m, v407)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L9
	} else {
		goto L157
	}
L157:
	;
	v507 = v384
	v512 = v389
	goto L151
L158:
	;
	v489 = F_makeNotNullConstraint(m, v405)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L9
	} else {
		goto L174
	}
L159:
	;
	if v384&int32(1) == int32(0) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v419 = F_find_inheritance_children(m, v418, l5)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L9
	} else {
		goto L163
	}
L161:
	;
	v421 = v389
	goto L162
L162:
	;
	if v421 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v421 = v419
	goto L162
L164:
	;
	v473 = int32(1)
	v478 = int32(0)
	goto L158
L165:
	;
	goto L166
L166:
	;
	v427 = int32(0)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v421)+4))
	if v428 <= v427 {
		v473 = int32(1)
		v478 = v421
		goto L158
	} else {
		goto L167
	}
L167:
	;
	v431 = v427
	goto L168
L168:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v421)+12))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v451+v431<<(uint(int32(2))%32))))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	v457 = F_findNotNullConstraint(m, v455, v456)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L9
	} else {
		goto L170
	}
L169:
	;
	v473 = v464
	v478 = v421
	goto L158
L170:
	;
	if v457 == int32(0) {
		goto L144
	} else {
		goto L171
	}
L171:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	F_verifyNotNullPKCompatible(m, v457, v461)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L9
	} else {
		goto L172
	}
L172:
	;
	v464 = int32(1)
	v466 = v431 + v464
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v421)+4))
	if v466 < v467 {
		v431 = v466
		goto L168
	} else {
		goto L173
	}
L173:
	;
	goto L169
L174:
	;
	v492 = F_palloc0(m, int32(32))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L9
	} else {
		goto L175
	}
L175:
	;
	v494 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v492)+29)) = uint8(v494)
	*(*int64)(unsafe.Add(mBase, uint32(v492))) = int64(68719476883)
	*(*int32)(unsafe.Add(mBase, uint32(v492)+20)) = v489
	F_ATPrepCmd(m, l0, l1, v492, v494, int32(0), l5, l6)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L9
	} else {
		goto L176
	}
L176:
	;
	v507 = v473
	v512 = v478
	goto L151
L177:
	;
	goto L150
L178:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	v555 = F_get_rel_name(m, v455)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L9
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v368)+4)) = v555
	*(*int32)(unsafe.Add(mBase, uint32(v368))) = v554
	F_errmsg(m, int32(_a_F_ATPrepCmd_12), v368)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L9
	} else {
		goto L180
	}
L180:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_13), int32(_a_F_ATPrepCmd_14))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L9
	} else {
		goto L181
	}
L181:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L182:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v572 = F_find_all_inheritors(m, v570, l5, int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L9
	} else {
		goto L183
	}
L183:
	;
	v574 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+29)) = uint8(v574)
	v1860 = v567
	v1861 = v137
	goto L15
L184:
	;
	v1860 = int32(8)
	v1861 = v137
	goto L15
L185:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585)+119)))
	if v586 == int32(112) {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v686 = int32(0)
	if l3 == v686 {
		v1860 = v686
		v1861 = v137
		goto L15
	} else {
		goto L205
	}
L187:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v591 = F_find_all_inheritors(m, v589, l5, int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L9
	} else {
		goto L191
	}
L188:
	;
	goto L189
L189:
	;
	goto L186
L190:
	;
	F_list_free(m, v591)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L9
	} else {
		goto L204
	}
L191:
	;
	if v591 == int32(0) {
		goto L190
	} else {
		goto L192
	}
L192:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v591)+4))
	if v595 < int32(2) {
		goto L190
	} else {
		goto L193
	}
L193:
	;
	v604 = int32(1)
	goto L194
L194:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v591)+12))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v619+v604<<(uint(int32(2))%32))))
	v625 = F_table_open(m, v623, int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L9
	} else {
		goto L196
	}
L195:
	;
	goto L190
L196:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v625)+48))
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627)+118)))
	if v628 == int32(116) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625)+24)))
	if v631 == int32(0) {
		goto L13
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	F_CheckTableNotInUse(m, v625, int32(_a_F_ATPrepCmd_15))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L9
	} else {
		goto L201
	}
L200:
	;
	goto L199
L201:
	;
	F_relation_close(m, v625, int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L9
	} else {
		goto L202
	}
L202:
	;
	v641 = v604 + int32(1)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v591)+4))
	if v641 < v642 {
		v604 = v641
		goto L194
	} else {
		goto L203
	}
L203:
	;
	goto L195
L204:
	;
	goto L189
L205:
	;
	v689 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+29)) = uint8(v689)
	v1860 = v686
	v1861 = v137
	goto L15
L206:
	;
	v696 = F_ATParseTransformCmd(m, v117, l1, v137, l3, int32(-1), l6)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L9
	} else {
		goto L207
	}
L207:
	;
	v698 = m.G0
	v700 = v698 - int32(208)
	m.G0 = v700
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v696)+8))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v696)+20))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v703)+32))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v703)+8))
	v707 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L9
	} else {
		goto L208
	}
L208:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v707)+4)) = v709
	if l4 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L209:
	;
	v1860 = int32(1)
	v1861 = v696
	goto L15
L210:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L9
	} else {
		goto L408
	}
L211:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L9
	} else {
		goto L403
	}
L212:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L9
	} else {
		goto L399
	}
L213:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L9
	} else {
		goto L395
	}
L214:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L9
	} else {
		goto L391
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700)+148)) = v814
	*(*int32)(unsafe.Add(mBase, uint32(v700)+144)) = v702
	F_errmsg(m, int32(_a_F_ATPrepCmd_16), v700+int32(144))
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L9
	} else {
		goto L388
	}
L216:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L9
	} else {
		goto L383
	}
L217:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L9
	} else {
		goto L378
	}
L218:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L9
	} else {
		goto L372
	}
L219:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L9
	} else {
		goto L367
	}
L220:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L9
	} else {
		goto L362
	}
L221:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L9
	} else {
		goto L357
	}
L222:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v713)+76))
	if v714 != 0 {
		goto L221
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v716 = F_SearchSysCacheAttName(m, v715, v702)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L9
	} else {
		goto L226
	}
L225:
	;
	goto L224
L226:
	;
	if v716 == int32(0) {
		goto L220
	} else {
		goto L227
	}
L227:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v716)+16))
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v720)+22)))
	v722 = v720 + v721
	v723 = int32(*(*int16)(unsafe.Add(mBase, uint32(v722)+74)))
	if v723 <= int32(0) {
		goto L219
	} else {
		goto L228
	}
L228:
	;
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722)+90)))
	if v726 != 0 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v703)+32))
	if v727 != 0 {
		goto L218
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	if l4 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	goto L231
L233:
	;
	v730 = int32(*(*int16)(unsafe.Add(mBase, uint32(v722)+94)))
	if int32(0) < v730 {
		goto L217
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	v735 = F_bms_make_singleton(m, v723+int32(7))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L9
	} else {
		goto L237
	}
L236:
	;
	goto L235
L237:
	;
	v739 = F_has_partition_attrs(m, l1, v735, v700+int32(199))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L9
	} else {
		goto L238
	}
L238:
	;
	if v739 != 0 {
		goto L216
	} else {
		goto L239
	}
L239:
	;
	F_typenameTypeIdAndMod(m, v707, v705, v700+int32(204), v700+int32(200))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L9
	} else {
		goto L240
	}
L240:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v700)+204))
	v750 = *(*int32)(unsafe.Add(mBase, _c_F_ATPrepCmd[0]))
	v752 = F_object_aclcheck(m, int32(1247), v748, v750, int64(256))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L9
	} else {
		goto L241
	}
L241:
	;
	if v752 != 0 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v700)+204))
	F_aclcheck_error_type(m, v752, v754)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L9
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v700)+204))
	v758 = F_GetColumnDefCollation(m, v707, v703, v757)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L9
	} else {
		goto L246
	}
L245:
	;
	goto L244
L246:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v760)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+192)) = v761
	*(*int32)(unsafe.Add(mBase, uint32(v700)+156)) = v761
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v700)+204))
	v768 = F_list_make1_impl(m, int32(480), v700+int32(156))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L9
	} else {
		goto L247
	}
L247:
	;
	v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722)+90)))
	if v772 == int32(118) {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v775 = int32(8)
	goto L250
L249:
	;
	v775 = int32(0)
	goto L250
L250:
	;
	F_CheckAttributeType(m, v702, v764, v758, v768, v775)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L9
	} else {
		goto L251
	}
L251:
	;
	v778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722)+90)))
	if v778 == int32(118) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+4)))
	switch v983 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L307
	default:
		goto L306
	}
L253:
	;
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+4)))
	switch v781 - int32(112) {
	case 0, 2:
		goto L255
	default:
		goto L254
	}
L254:
	;
	if v704 != 0 {
		goto L214
	} else {
		goto L304
	}
L255:
	;
	if v704 == int32(0) {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v722)+68))
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v722)+76))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v722)+96))
	v791 = F_makeVar(m, int32(1), v723, v787, v788, v789, int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L9
	} else {
		goto L259
	}
L257:
	;
	v793 = v704
	goto L258
L258:
	;
	v794 = F_exprType(m, v793)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L9
	} else {
		goto L260
	}
L259:
	;
	v793 = v791
	goto L258
L260:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v700)+204))
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v700)+200))
	v801 = F_coerce_to_target_type(m, v707, v793, v794, v796, v797, int32(1), int32(2), int32(-1))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L9
	} else {
		goto L261
	}
L261:
	;
	if v801 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v703)+32))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L9
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	F_assign_expr_collations(m, v707, v801)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L9
	} else {
		goto L277
	}
L265:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L9
	} else {
		goto L266
	}
L266:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v700)+204))
	v814 = F_format_type_be(m, v813)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L9
	} else {
		goto L267
	}
L267:
	;
	if v805 != 0 {
		goto L215
	} else {
		goto L268
	}
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700)+132)) = v814
	*(*int32)(unsafe.Add(mBase, uint32(v700)+128)) = v702
	F_errmsg(m, int32(_a_F_ATPrepCmd_17), v700+int32(128))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L9
	} else {
		goto L269
	}
L269:
	;
	v823 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722)+90)))
	if v823 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v826 = F_quote_identifier(m, v702)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L9
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_18), int32(_a_F_ATPrepCmd_19))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L9
	} else {
		goto L276
	}
L273:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v700)+204))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v700)+200))
	v830 = F_format_type_with_typemod(m, v828, v829)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L9
	} else {
		goto L274
	}
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700)+116)) = v830
	*(*int32)(unsafe.Add(mBase, uint32(v700)+112)) = v826
	F_errhint(m, int32(_a_F_ATPrepCmd_20), v700+int32(112))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L9
	} else {
		goto L275
	}
L275:
	;
	goto L272
L276:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L277:
	;
	v848 = F_expand_generated_columns_in_expr(m, v801, l1, int32(1))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L9
	} else {
		goto L278
	}
L278:
	;
	v850 = F_expression_planner(m, v848)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L9
	} else {
		goto L279
	}
L279:
	;
	v853 = F_palloc0(m, int32(16))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L9
	} else {
		goto L280
	}
L280:
	;
	v855 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v853)+12)) = uint8(v855)
	*(*int32)(unsafe.Add(mBase, uint32(v853)+4)) = v850
	*(*uint16)(unsafe.Add(mBase, uint32(v853))) = uint16(v723)
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v117)+68))
	v860 = F_lappend(m, v859, v853)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L9
	} else {
		goto L281
	}
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+68)) = v860
	v871 = v850
	goto L283
L282:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v117)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+80)) = v959 | int32(4)
	goto L252
L283:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v871)))
	switch v883 - int32(6) {
	case 0:
		goto L285
	case 1, 2, 3, 4, 5, 6, 7, 8, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20:
		goto L282
	case 9:
		goto L286
	case 21:
		goto L288
	default:
		goto L287
	}
L284:
	;
	v953 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v871)+8)))
	if v953 == v723&int32(_a_F_ATPrepCmd_21) {
		goto L252
	} else {
		goto L303
	}
L285:
	;
	goto L284
L286:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v871)+4))
	if base.Ui32(int32(1)) < base.Ui32(v893-int32(2027)) {
		goto L282
	} else {
		goto L292
	}
L287:
	;
	if v883 != int32(55) {
		goto L282
	} else {
		goto L289
	}
L288:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v871)+4))
	v871 = v886
	goto L283
L289:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v871)+8))
	v890 = F_DomainHasConstraints(m, v889)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L9
	} else {
		goto L290
	}
L290:
	;
	if v890 != 0 {
		goto L282
	} else {
		goto L291
	}
L291:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v871)+4))
	v871 = v892
	goto L283
L292:
	;
	v898 = m.G0
	v900 = v898 - int32(16)
	m.G0 = v900
	v903 = *(*int32)(unsafe.Add(mBase, _c_F_ATPrepCmd[1]))
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v903)+uint32(_c_F_ATPrepCmd[2])))
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v903)+264))
	if v910 < int32(2) {
		goto L294
	} else {
		goto L295
	}
L293:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v900)+12))
	m.G0 = v900 + int32(16)
	if v943|(v942^int32(1)) != 0 {
		goto L282
	} else {
		goto L302
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v900+int32(12)))) = v909
	v942 = int32(1)
	goto L293
L295:
	;
	v916 = int32(1)
	goto L296
L296:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v903+int32(_a_F_ATPrepCmd_22)+v916<<(uint(int32(4))%32))))
	if v909 == v924 {
		goto L298
	} else {
		goto L299
	}
L297:
	;
	v942 = int32(0)
	goto L293
L298:
	;
	v927 = v916 + int32(1)
	if v910 != v927 {
		v916 = v927
		goto L296
	} else {
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	goto L297
L301:
	;
	goto L294
L302:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v871)+28))
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v950)+12))
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v951)))
	v871 = v952
	goto L283
L303:
	;
	goto L282
L304:
	;
	goto L252
L305:
	;
	F_ReleaseCatCache(m, v716)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L9
	} else {
		goto L310
	}
L306:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v989)+72))
	F_find_composite_type_dependencies(m, v990, l1, int32(0))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L9
	} else {
		goto L309
	}
L307:
	;
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722)+90)))
	if v986 != int32(118) {
		goto L305
	} else {
		goto L308
	}
L308:
	;
	goto L306
L309:
	;
	goto L305
L310:
	;
	if l3 != 0 {
		goto L312
	} else {
		goto L313
	}
L311:
	;
	v1139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+4)))
	if v1139 == int32(99) {
		goto L353
	} else {
		goto L354
	}
L312:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v999 = F_find_all_inheritors(m, v996, l5, v700+int32(188))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L9
	} else {
		goto L315
	}
L313:
	;
	goto L314
L314:
	;
	if l4 != 0 {
		v1121 = v696
		goto L311
	} else {
		goto L350
	}
L315:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v700)+188))
	v1003 = v696
	v1011 = int32(0)
	goto L316
L316:
	;
	v1023 = int32(0)
	if v999 == v1023 {
		v1033 = v1023
		goto L318
	} else {
		goto L319
	}
L318:
	;
	if v1001 == int32(0) {
		v1121 = v696
		goto L311
	} else {
		goto L321
	}
L319:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v999)+4))
	if v1027 <= v1011 {
		v1033 = int32(0)
		goto L318
	} else {
		goto L320
	}
L320:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v999)+12))
	v1033 = v1029 + v1011<<(uint(int32(2))%32)
	goto L318
L321:
	;
	v1036 = int32(0)
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+4))
	if base.B2i32(v1033 == v1036)|base.B2i32(v1038 <= v1011) == v1036 {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v1033)))
	if v996 != v1045 {
		goto L327
	} else {
		goto L328
	}
L323:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+12))
	if v1043 != 0 {
		goto L322
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	v1121 = v1003
	goto L311
L326:
	;
	goto L325
L327:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1043+v1011<<(uint(int32(2))%32))))
	v1052 = F_relation_open(m, v1045, int32(0))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L9
	} else {
		goto L330
	}
L328:
	;
	v1108 = v1003
	goto L329
L329:
	;
	v1003 = v1108
	v1011 = v1011 + int32(1)
	goto L316
L330:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+48))
	v1055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1054)+118)))
	if v1055 == int32(116) {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v1058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052)+24)))
	if v1058 == int32(0) {
		goto L13
	} else {
		goto L334
	}
L332:
	;
	goto L333
L333:
	;
	F_CheckTableNotInUse(m, v1052, int32(_a_F_ATPrepCmd_15))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L9
	} else {
		goto L335
	}
L334:
	;
	goto L333
L335:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+56))
	v1065 = F_SearchSysCacheAttName(m, v1064, v702)
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L9
	} else {
		goto L336
	}
L336:
	;
	if v1065 == int32(0) {
		goto L213
	} else {
		goto L337
	}
L337:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1065)+16))
	v1070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1069)+22)))
	v1072 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1069+v1070)+94)))
	if v1050 < v1072 {
		goto L212
	} else {
		goto L338
	}
L338:
	;
	F_ReleaseCatCache(m, v1065)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L9
	} else {
		goto L339
	}
L339:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v703)+32))
	if v1076 != 0 {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v1077 = F_copyObjectImpl(m, v1003)
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L9
	} else {
		goto L343
	}
L341:
	;
	v1098 = v1003
	goto L342
L342:
	;
	F_ATPrepCmd(m, l0, v1052, v1098, int32(0), int32(1), l5, l6)
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L9
	} else {
		goto L348
	}
L343:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+52))
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v1082 = F_build_attrmap_by_name(m, v1079, v1080, int32(0))
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L9
	} else {
		goto L344
	}
L344:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v703)+32))
	v1089 = F_map_variable_attnos(m, v1084, int32(1), v1082, int32(0), v700+int32(187))
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L9
	} else {
		goto L345
	}
L345:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1091)+32)) = v1089
	v1093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700)+187)))
	if v1093 == int32(1) {
		goto L211
	} else {
		goto L346
	}
L346:
	;
	F_pfree(m, v1082)
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L9
	} else {
		goto L347
	}
L347:
	;
	v1098 = v1077
	goto L342
L348:
	;
	F_relation_close(m, v1052, int32(0))
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L9
	} else {
		goto L349
	}
L349:
	;
	v1108 = v1098
	goto L329
L350:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v1117 = F_find_inheritance_children(m, v1115, int32(0))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L9
	} else {
		goto L351
	}
L351:
	;
	if v1117 != 0 {
		goto L210
	} else {
		goto L352
	}
L352:
	;
	v1121 = v696
	goto L311
L353:
	;
	F_ATTypedTableRecursion(m, l0, l1, v1121, l5, l6)
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L9
	} else {
		goto L356
	}
L354:
	;
	goto L355
L355:
	;
	m.G0 = v700 + int32(208)
	goto L209
L356:
	;
	goto L355
L357:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L9
	} else {
		goto L358
	}
L358:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_23), int32(0))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L9
	} else {
		goto L359
	}
L359:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v703)+64))
	F_parser_errposition(m, v707, v1158)
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L9
	} else {
		goto L360
	}
L360:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_24), int32(_a_F_ATPrepCmd_19))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L9
	} else {
		goto L361
	}
L361:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L362:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L9
	} else {
		goto L363
	}
L363:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v700))) = v702
	*(*int32)(unsafe.Add(mBase, uint32(v700)+4)) = v1173 + int32(4)
	F_errmsg(m, int32(_a_F_ATPrepCmd_5), v700)
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L9
	} else {
		goto L364
	}
L364:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v703)+64))
	F_parser_errposition(m, v707, v1181)
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L9
	} else {
		goto L365
	}
L365:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_25), int32(_a_F_ATPrepCmd_19))
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L9
	} else {
		goto L366
	}
L366:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L367:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L9
	} else {
		goto L368
	}
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700)+16)) = v702
	F_errmsg(m, int32(_a_F_ATPrepCmd_26), v700+int32(16))
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L9
	} else {
		goto L369
	}
L369:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v703)+64))
	F_parser_errposition(m, v707, v1202)
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L9
	} else {
		goto L370
	}
L370:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_27), int32(_a_F_ATPrepCmd_19))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L9
	} else {
		goto L371
	}
L371:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L372:
	;
	F_errcode(m, int32(17064068))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L9
	} else {
		goto L373
	}
L373:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_28), int32(0))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L9
	} else {
		goto L374
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700)+176)) = v702
	v1225 = F_errdetail(m, int32(_a_F_ATPrepCmd_29), v700+int32(176))
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L9
	} else {
		goto L375
	}
L375:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v703)+64))
	F_parser_errposition(m, v707, v1227)
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L9
	} else {
		goto L376
	}
L376:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_30), int32(_a_F_ATPrepCmd_19))
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L9
	} else {
		goto L377
	}
L377:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L378:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L9
	} else {
		goto L379
	}
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700)+160)) = v702
	F_errmsg(m, int32(_a_F_ATPrepCmd_31), v700+int32(160))
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L9
	} else {
		goto L380
	}
L380:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v703)+64))
	F_parser_errposition(m, v707, v1248)
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L9
	} else {
		goto L381
	}
L381:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_32), int32(_a_F_ATPrepCmd_19))
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L9
	} else {
		goto L382
	}
L382:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L383:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L9
	} else {
		goto L384
	}
L384:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+32)) = v702
	*(*int32)(unsafe.Add(mBase, uint32(v700)+36)) = v1263 + int32(4)
	F_errmsg(m, int32(_a_F_ATPrepCmd_33), v700+int32(32))
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L9
	} else {
		goto L385
	}
L385:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v703)+64))
	F_parser_errposition(m, v707, v1273)
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L9
	} else {
		goto L386
	}
L386:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_34), int32(_a_F_ATPrepCmd_19))
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L9
	} else {
		goto L387
	}
L387:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L388:
	;
	F_errhint(m, int32(_a_F_ATPrepCmd_35), int32(0))
	mBase = m.M
	v1291 = m.ExcPending
	if v1291 != 0 {
		goto L9
	} else {
		goto L389
	}
L389:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_36), int32(_a_F_ATPrepCmd_19))
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L9
	} else {
		goto L390
	}
L390:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L391:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L9
	} else {
		goto L392
	}
L392:
	;
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+96)) = v1304 + int32(4)
	F_errmsg(m, int32(_a_F_ATPrepCmd_37), v700+int32(96))
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L9
	} else {
		goto L393
	}
L393:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_38), int32(_a_F_ATPrepCmd_19))
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
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
	F_errcode(m, int32(50360452))
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L9
	} else {
		goto L396
	}
L396:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+48)) = v702
	*(*int32)(unsafe.Add(mBase, uint32(v700)+52)) = v1325 + int32(4)
	F_errmsg(m, int32(_a_F_ATPrepCmd_5), v700+int32(48))
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L9
	} else {
		goto L397
	}
L397:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_39), int32(_a_F_ATPrepCmd_19))
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L9
	} else {
		goto L398
	}
L398:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L399:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L9
	} else {
		goto L400
	}
L400:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+64)) = v702
	*(*int32)(unsafe.Add(mBase, uint32(v700)+68)) = v1347 + int32(4)
	F_errmsg(m, int32(_a_F_ATPrepCmd_40), v700-int32(-64))
	mBase = m.M
	v1356 = m.ExcPending
	if v1356 != 0 {
		goto L9
	} else {
		goto L401
	}
L401:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_41), int32(_a_F_ATPrepCmd_19))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L9
	} else {
		goto L404
	}
L404:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_42), int32(0))
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L9
	} else {
		goto L405
	}
L405:
	;
	v1375 = F_errdetail(m, int32(_a_F_ATPrepCmd_43), int32(0))
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L9
	} else {
		goto L406
	}
L406:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_44), int32(_a_F_ATPrepCmd_19))
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L9
	} else {
		goto L407
	}
L407:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L408:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L9
	} else {
		goto L409
	}
L409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700)+80)) = v702
	F_errmsg(m, int32(_a_F_ATPrepCmd_45), v700+int32(80))
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L9
	} else {
		goto L410
	}
L410:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_46), int32(_a_F_ATPrepCmd_19))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L9
	} else {
		goto L411
	}
L411:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L412:
	;
	v1860 = v136
	v1861 = v137
	goto L15
L413:
	;
	v1860 = v136
	v1861 = v137
	goto L15
L414:
	;
	v1411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+96)))
	if v1411 == int32(1) {
		goto L20
	} else {
		goto L415
	}
L415:
	;
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1415)+118)))
	switch v1416 - int32(112) {
	case 0:
		goto L417
	default:
		goto L418
	case 4:
		goto L420
	case 5:
		goto L419
	}
L416:
	;
	v1450 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+56)))
	v1451 = int64(0)
	v1453 = F_SearchSysCacheList(m, int32(53), int32(1), v1450, v1451, v1451)
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L9
	} else {
		goto L429
	}
L417:
	;
	if v1414 == int32(29) {
		v1860 = v136
		v1861 = v137
		goto L15
	} else {
		goto L428
	}
L418:
	;
	if v1414 == int32(29) {
		goto L17
	} else {
		goto L427
	}
L419:
	;
	if v1414 != int32(29) {
		v1860 = v136
		v1861 = v137
		goto L15
	} else {
		goto L426
	}
L420:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L9
	} else {
		goto L421
	}
L421:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L9
	} else {
		goto L422
	}
L422:
	;
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v1426 + int32(4)
	F_errmsg(m, int32(_a_F_ATPrepCmd_47), v23-int32(-64))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L9
	} else {
		goto L423
	}
L423:
	;
	F_errtable(m, l1)
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L9
	} else {
		goto L424
	}
L424:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_48), int32(_a_F_ATPrepCmd_49))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
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
	goto L17
L427:
	;
	goto L416
L428:
	;
	goto L416
L429:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v1453)+56))
	F_ReleaseCatCacheList(m, v1453)
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L9
	} else {
		goto L430
	}
L430:
	;
	if int32(0) < v1455 {
		goto L19
	} else {
		goto L431
	}
L431:
	;
	v1712 = int32(0)
	v1713 = int32(13)
	goto L16
L432:
	;
	v1860 = int32(0)
	v1861 = v137
	goto L15
L433:
	;
	v1471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+84)))
	if v1471 == int32(1) {
		goto L18
	} else {
		goto L434
	}
L434:
	;
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	if v1474 != 0 {
		goto L436
	} else {
		goto L437
	}
L435:
	;
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1487)+84))
	if v1486 != v1488 {
		goto L441
	} else {
		goto L442
	}
L436:
	;
	v1482 = v1474
	goto L438
L437:
	;
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1476)+119)))
	if v1477 == int32(112) {
		v1486 = int32(0)
		goto L435
	} else {
		goto L439
	}
L438:
	;
	v1484 = F_get_table_am_oid(m, v1482, int32(0))
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L9
	} else {
		goto L440
	}
L439:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, _c_F_ATPrepCmd[3]))
	v1482 = v1481
	goto L438
L440:
	;
	v1486 = v1484
	goto L435
L441:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+88)) = v1486
	v1491 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v117)+84)) = uint8(v1491)
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v117)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+80)) = v1493 | int32(8)
	goto L443
L442:
	;
	goto L443
L443:
	;
	v1860 = v136
	v1861 = v137
	goto L15
L444:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	v1503 = F_get_tablespace_oid(m, v1501, int32(0))
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L9
	} else {
		goto L446
	}
L445:
	;
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v117)+92))
	if v1522 != 0 {
		goto L452
	} else {
		goto L453
	}
L446:
	;
	if v1503 == int32(0) {
		goto L445
	} else {
		goto L447
	}
L447:
	;
	v1508 = *(*int32)(unsafe.Add(mBase, _c_F_ATPrepCmd[4]))
	if v1503 == v1508 {
		goto L445
	} else {
		goto L448
	}
L448:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, _c_F_ATPrepCmd[0]))
	v1514 = F_object_aclcheck(m, int32(1213), v1503, v1512, int64(512))
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L9
	} else {
		goto L449
	}
L449:
	;
	if v1514 == int32(0) {
		goto L445
	} else {
		goto L450
	}
L450:
	;
	F_aclcheck_error(m, v1514, int32(43), v1501)
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L9
	} else {
		goto L451
	}
L451:
	;
	goto L445
L452:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L9
	} else {
		goto L455
	}
L453:
	;
	goto L454
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+92)) = v1503
	v1860 = v136
	v1861 = v137
	goto L15
L455:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L9
	} else {
		goto L456
	}
L456:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_50), int32(0))
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		goto L9
	} else {
		goto L457
	}
L457:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_51), int32(_a_F_ATPrepCmd_52))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L9
	} else {
		goto L458
	}
L458:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L459:
	;
	v1860 = v136
	v1861 = v137
	goto L15
L460:
	;
	F_ATPrepChangeInherit(m, l1)
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L9
	} else {
		goto L461
	}
L461:
	;
	v1860 = v136
	v1861 = v137
	goto L15
L462:
	;
	F_ATPrepChangeInherit(m, l1)
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L9
	} else {
		goto L463
	}
L463:
	;
	v1860 = v136
	v1861 = v137
	goto L15
L464:
	;
	if l3 == int32(0) {
		v1860 = v136
		v1861 = v137
		goto L15
	} else {
		goto L465
	}
L465:
	;
	v1561 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+29)) = uint8(v1561)
	v1860 = v136
	v1861 = v137
	goto L15
L466:
	;
	if l3 == int32(0) {
		v1860 = v136
		v1861 = v137
		goto L15
	} else {
		goto L467
	}
L467:
	;
	v1569 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+29)) = uint8(v1569)
	v1860 = v136
	v1861 = v137
	goto L15
L468:
	;
	v1860 = v136
	v1861 = v137
	goto L15
L469:
	;
	if l3 == int32(0) {
		v1860 = v136
		v1861 = v137
		goto L15
	} else {
		goto L470
	}
L470:
	;
	v1580 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+29)) = uint8(v1580)
	v1860 = v136
	v1861 = v137
	goto L15
L471:
	;
	v1860 = v136
	v1861 = v137
	goto L15
L472:
	;
	v1860 = v136
	v1861 = v137
	goto L15
L473:
	;
	v1860 = v136
	v1861 = v137
	goto L15
L474:
	;
	v1860 = v136
	v1861 = v137
	goto L15
L475:
	;
	v1860 = v136
	v1861 = v137
	goto L15
L476:
	;
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v1605
	F_errmsg_internal(m, int32(_a_F_ATPrepCmd_53), v23)
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		goto L9
	} else {
		goto L477
	}
L477:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_54), int32(_a_F_ATPrepCmd_55))
	mBase = m.M
	v1614 = m.ExcPending
	if v1614 != 0 {
		goto L9
	} else {
		goto L478
	}
L478:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L479:
	;
	F_ATPrepAddColumn(m, l0, l1, l3, l4, int32(0), v137, l5, l6)
	mBase = m.M
	v1621 = m.ExcPending
	if v1621 != 0 {
		goto L9
	} else {
		goto L480
	}
L480:
	;
	v1860 = int32(2)
	v1861 = v137
	goto L15
L481:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1629 = m.ExcPending
	if v1629 != 0 {
		goto L9
	} else {
		goto L482
	}
L482:
	;
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v1630 + int32(4)
	F_errmsg(m, int32(_a_F_ATPrepCmd_56), v23+int32(80))
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L9
	} else {
		goto L483
	}
L483:
	;
	F_errhint(m, int32(_a_F_ATPrepCmd_57), int32(0))
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L9
	} else {
		goto L484
	}
L484:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_58), int32(_a_F_ATPrepCmd_55))
	mBase = m.M
	v1647 = m.ExcPending
	if v1647 != 0 {
		goto L9
	} else {
		goto L485
	}
L485:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L486:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1654 = m.ExcPending
	if v1654 != 0 {
		goto L9
	} else {
		goto L487
	}
L487:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_59), int32(0))
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L9
	} else {
		goto L488
	}
L488:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_60), int32(_a_F_ATPrepCmd_55))
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
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
	F_errcode(m, int32(325))
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L9
	} else {
		goto L491
	}
L491:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v1671 + int32(4)
	F_errmsg(m, int32(_a_F_ATPrepCmd_61), v23+int32(48))
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L9
	} else {
		goto L492
	}
L492:
	;
	v1682 = F_errdetail(m, int32(_a_F_ATPrepCmd_62), int32(0))
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L9
	} else {
		goto L493
	}
L493:
	;
	F_errhint(m, int32(_a_F_ATPrepCmd_63), int32(0))
	mBase = m.M
	v1687 = m.ExcPending
	if v1687 != 0 {
		goto L9
	} else {
		goto L494
	}
L494:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_64), int32(_a_F_ATPrepCmd_49))
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L9
	} else {
		goto L495
	}
L495:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L496:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L9
	} else {
		goto L497
	}
L497:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_65), int32(0))
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		goto L9
	} else {
		goto L498
	}
L498:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_66), int32(_a_F_ATPrepCmd_55))
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L9
	} else {
		goto L499
	}
L499:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L500:
	;
	v1719 = v23 + int32(96)
	v1722 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+56)))
	F_ScanKeyInit(m, v1719, v1713, int32(3), int32(184), v1722)
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L9
	} else {
		goto L501
	}
L501:
	;
	v1725 = int32(1)
	v1728 = F_systable_beginscan(m, v1716, v1712, v1725, int32(0), v1725, v1719)
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		goto L9
	} else {
		goto L502
	}
L502:
	;
	v1730 = F_systable_getnext(m, v1728)
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L9
	} else {
		goto L503
	}
L503:
	;
	if v1730 != 0 {
		goto L504
	} else {
		goto L505
	}
L504:
	;
	if v1414 == int32(29) {
		goto L507
	} else {
		goto L508
	}
L505:
	;
	goto L506
L506:
	;
	F_systable_endscan(m, v1728)
	mBase = m.M
	v1835 = m.ExcPending
	if v1835 != 0 {
		goto L9
	} else {
		goto L530
	}
L507:
	;
	v1736 = int32(96)
	goto L509
L508:
	;
	v1736 = int32(80)
	goto L509
L509:
	;
	v1741 = v1730
	goto L510
L510:
	;
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v1741)+16))
	v1758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1757)+22)))
	v1759 = v1757 + v1758
	v1760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1759)+72)))
	if v1760 != int32(102) {
		goto L512
	} else {
		goto L513
	}
L511:
	;
	goto L506
L512:
	;
	v1812 = F_systable_getnext(m, v1728)
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L9
	} else {
		goto L528
	}
L513:
	;
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v1759+v1736)))
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v1764 == v1765 {
		goto L512
	} else {
		goto L514
	}
L514:
	;
	v1768 = F_relation_open(m, v1764, int32(1))
	mBase = m.M
	v1769 = m.ExcPending
	if v1769 != 0 {
		goto L9
	} else {
		goto L515
	}
L515:
	;
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(v1768)+48))
	v1771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1770)+118)))
	if v1414 == int32(29) {
		goto L517
	} else {
		goto L518
	}
L516:
	;
	F_relation_close(m, v1768, int32(1))
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L9
	} else {
		goto L527
	}
L517:
	;
	if v1771 == int32(112) {
		goto L516
	} else {
		goto L520
	}
L518:
	;
	goto L519
L519:
	;
	if v1771 == int32(112) {
		goto L14
	} else {
		goto L526
	}
L520:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1779 = m.ExcPending
	if v1779 != 0 {
		goto L9
	} else {
		goto L521
	}
L521:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L9
	} else {
		goto L522
	}
L522:
	;
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(v1768)+48))
	v1785 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v1784 + v1785
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v1783 + v1785
	F_errmsg(m, int32(_a_F_ATPrepCmd_67), v23+int32(16))
	mBase = m.M
	v1795 = m.ExcPending
	if v1795 != 0 {
		goto L9
	} else {
		goto L523
	}
L523:
	;
	F_errtableconstraint(m, l1, v1759+int32(4))
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L9
	} else {
		goto L524
	}
L524:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_68), int32(_a_F_ATPrepCmd_49))
	mBase = m.M
	v1804 = m.ExcPending
	if v1804 != 0 {
		goto L9
	} else {
		goto L525
	}
L525:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L526:
	;
	goto L516
L527:
	;
	goto L512
L528:
	;
	if v1812 != 0 {
		v1741 = v1812
		goto L510
	} else {
		goto L529
	}
L529:
	;
	goto L511
L530:
	;
	F_relation_close(m, v1716, int32(1))
	mBase = m.M
	v1838 = m.ExcPending
	if v1838 != 0 {
		goto L9
	} else {
		goto L531
	}
L531:
	;
	if v1414 == int32(29) {
		goto L532
	} else {
		goto L533
	}
L532:
	;
	v1843 = int32(112)
	goto L534
L533:
	;
	v1843 = int32(117)
	goto L534
L534:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v117)+97)) = uint8(v1843)
	v1845 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v117)+96)) = uint8(v1845)
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(v117)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+80)) = v1847 | v1845
	v1860 = v136
	v1861 = v137
	goto L15
L535:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1873)+16)) = v1875
	m.G0 = v23 + int32(160)
	return
L536:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1887 = m.ExcPending
	if v1887 != 0 {
		goto L9
	} else {
		goto L537
	}
L537:
	;
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1768)+48))
	v1890 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v1889 + v1890
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v1888 + v1890
	F_errmsg(m, int32(_a_F_ATPrepCmd_69), v23+int32(32))
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		goto L9
	} else {
		goto L538
	}
L538:
	;
	F_errtableconstraint(m, l1, v1759+int32(4))
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
		goto L9
	} else {
		goto L539
	}
L539:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_70), int32(_a_F_ATPrepCmd_49))
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L9
	} else {
		goto L540
	}
L540:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L541:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		goto L9
	} else {
		goto L542
	}
L542:
	;
	F_errmsg(m, int32(_a_F_ATPrepCmd_71), int32(0))
	mBase = m.M
	v1940 = m.ExcPending
	if v1940 != 0 {
		goto L9
	} else {
		goto L543
	}
L543:
	;
	F_errfinish(m, int32(_a_F_ATPrepCmd_2), int32(_a_F_ATPrepCmd_72), int32(_a_F_ATPrepCmd_73))
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L9
	} else {
		goto L544
	}
L544:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
