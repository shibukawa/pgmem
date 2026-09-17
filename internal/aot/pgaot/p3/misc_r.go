package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReadCheckpointRecord(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	if base.Ui64(l1&int64(8184)) <= base.Ui64(int64(23)) {
		v9 = int32(0)
		v12 = F_errstart(m, int32(15), v9)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			if v12 == int32(0) {
				v81 = v9
				return v81
			} else {
				v71 = v9
				v72 = int32(_a_F_ReadCheckpointRecord_0)
				v73 = int32(_a_F_ReadCheckpointRecord_1)
				F_errmsg(m, v72, int32(0))
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_ReadCheckpointRecord_2), v73, int32(_a_F_ReadCheckpointRecord_3))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						v81 = v71
						return v81
					}
				}
			}
		}
	} else {
		F_XLogPrefetcherBeginRead(m, l0, l1)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v24 = F_ReadRecord(m, l0, int32(15), int32(1), l2)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v24 == int32(0) {
					v28 = int32(0)
					v31 = F_errstart(m, int32(15), v28)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if v31 == int32(0) {
							v81 = v28
							return v81
						} else {
							v71 = v28
							v72 = int32(_a_F_ReadCheckpointRecord_4)
							v73 = int32(_a_F_ReadCheckpointRecord_5)
							F_errmsg(m, v72, int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_ReadCheckpointRecord_2), v73, int32(_a_F_ReadCheckpointRecord_3))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									v81 = v71
									return v81
								}
							}
						}
					}
				} else {
					v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+17)))
					if v37 != 0 {
						v38 = int32(0)
						v41 = F_errstart(m, int32(15), v38)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							if v41 == int32(0) {
								v81 = v38
								return v81
							} else {
								v71 = v38
								v72 = int32(_a_F_ReadCheckpointRecord_6)
								v73 = int32(_a_F_ReadCheckpointRecord_7)
								F_errmsg(m, v72, int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ReadCheckpointRecord_2), v73, int32(_a_F_ReadCheckpointRecord_3))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										v81 = v71
										return v81
									}
								}
							}
						}
					} else {
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+16)))
						if base.Ui32(int32(32)) <= base.Ui32(v47) {
							v50 = int32(0)
							v53 = F_errstart(m, int32(15), v50)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								if v53 == int32(0) {
									v81 = v50
									return v81
								} else {
									v71 = v50
									v72 = int32(_a_F_ReadCheckpointRecord_8)
									v73 = int32(_a_F_ReadCheckpointRecord_9)
									F_errmsg(m, v72, int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_ReadCheckpointRecord_2), v73, int32(_a_F_ReadCheckpointRecord_3))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int32(0)
										} else {
											v81 = v71
											return v81
										}
									}
								}
							}
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
							if v59 == int32(114) {
								v81 = v24
								return v81
							} else {
								v62 = int32(0)
								v65 = F_errstart(m, int32(15), v62)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									if v65 == int32(0) {
										v81 = v62
										return v81
									} else {
										v71 = v62
										v72 = int32(_a_F_ReadCheckpointRecord_10)
										v73 = int32(_a_F_ReadCheckpointRecord_11)
										F_errmsg(m, v72, int32(0))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_ReadCheckpointRecord_2), v73, int32(_a_F_ReadCheckpointRecord_3))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int32(0)
											} else {
												v81 = v71
												return v81
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
func F_RegisterSnapshot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterSnapshot[0]))
		v8 = F_RegisterSnapshotOnOwner(m, l0, v7)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_RememberAllDependentForRebuilding(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
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
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	v12 = m.G0
	v14 = v12 - int32(320)
	m.G0 = v14
	v18 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = v14 + int32(176)
	F_ScanKeyInit(m, v21, int32(4), int32(3), int32(184), int32(1259))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	F_ScanKeyInit(m, v14+int32(224), int32(5), int32(3), int32(184), v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_ScanKeyInit(m, v14+int32(272), int32(6), int32(3), int32(65), l3)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v47 = F_systable_beginscan(m, v18, int32(2674), int32(1), int32(0), int32(3), v21)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L151
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L148
	}
L8:
	;
	v49 = F_systable_getnext(m, v47)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v49 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v52 = base.B2i32(l1 != int32(24))
	v59 = v49
	goto L13
L11:
	;
	goto L12
L12:
	;
	F_systable_endscan(m, v47)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L146
	}
L13:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+22)))
	v66 = v64 + v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+164)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+168)) = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+172)) = v71
	if v67 <= int32(3255) {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	goto L12
L15:
	;
	v441 = F_systable_getnext(m, v47)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L144
	}
L16:
	;
	F_RememberConstraintForRebuilding(m, v69, l0)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L143
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L139
	}
L18:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v368 = int32(0)
	if v367 == v368 {
		goto L123
	} else {
		goto L124
	}
L19:
	;
	F_GetAttrDefaultColumnAddress(m, v14+int32(152), v69)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L108
	}
L20:
	;
	if l1 != int32(24) {
		goto L15
	} else {
		goto L101
	}
L21:
	;
	if l1 != int32(24) {
		goto L15
	} else {
		goto L94
	}
L22:
	;
	if l1 != int32(24) {
		goto L15
	} else {
		goto L87
	}
L23:
	;
	if l1 != int32(24) {
		goto L15
	} else {
		goto L80
	}
L24:
	;
	v113 = F_get_rel_relkind(m, v69)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L39
	}
L25:
	;
	switch v67 - int32(2604) {
	case 0:
		goto L19
	case 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15:
		goto L17
	case 2:
		goto L16
	case 14:
		goto L22
	case 16:
		goto L21
	default:
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if v67 == int32(3256) {
		goto L20
	} else {
		goto L29
	}
L28:
	;
	switch v67 - int32(1255) {
	case 0:
		goto L23
	default:
		goto L17
	case 4:
		goto L24
	}
L29:
	;
	if v67 == int32(3381) {
		goto L18
	} else {
		goto L30
	}
L30:
	;
	if v67 != int32(_a_F_RememberAllDependentForRebuilding_0) {
		goto L17
	} else {
		goto L31
	}
L31:
	;
	if l1 != int32(24) {
		goto L15
	} else {
		goto L32
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errmsg(m, int32(_a_F_RememberAllDependentForRebuilding_1), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v99 = F_getObjectDescription(m, v14+int32(164), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+148)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = v99
	F_errdetail(m, int32(_a_F_RememberAllDependentForRebuilding_2), v14+int32(144))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_RememberAllDependentForRebuilding_3), int32(_a_F_RememberAllDependentForRebuilding_4), int32(_a_F_RememberAllDependentForRebuilding_5))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	if v113&int32(-33) == int32(73) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v14)+168))
	v121 = int32(0)
	if v119 == v121 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	goto L42
L42:
	;
	if v113 == int32(83) {
		goto L15
	} else {
		goto L75
	}
L43:
	;
	if v159 != 0 {
		goto L15
	} else {
		goto L56
	}
L44:
	;
	v159 = int32(0)
	goto L43
L45:
	;
	goto L46
L46:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	if v127 <= int32(0) {
		v153 = v121
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v159 = v153
	goto L43
L48:
	;
	v130 = int32(0)
	if v130 < v127 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v133 = v127
	goto L51
L50:
	;
	v133 = v130
	goto L51
L51:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	v136 = int32(0)
	goto L52
L52:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v134+v136<<(uint(int32(2))%32))))
	v145 = base.B2i32(v144 == v120)
	if v144 == v120 {
		v153 = v145
		goto L47
	} else {
		goto L54
	}
L53:
	;
	v153 = v145
	goto L47
L54:
	;
	v147 = v136 + int32(1)
	if v147 != v133 {
		v136 = v147
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v160 = F_get_index_constraint(m, v120)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	if v160 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	F_RememberConstraintForRebuilding(m, v160, l0)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v164 = int32(0)
	v168 = int32(1)
	v172 = F_pg_get_indexdef_worker(m, v120, v164, v164, v164, v164, v168, v168, v164, v164)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	goto L15
L62:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v175 = F_lappend_oid(m, v174, v120)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v175
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v179 = F_lappend(m, v178, v172)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v179
	v182 = F_get_index_isreplident(m, v120)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v182 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v184 != 0 {
		goto L7
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v188 = F_get_index_isclustered(m, v120)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	v185 = F_get_rel_name(m, v120)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v185
	goto L68
L71:
	;
	if v188 == int32(0) {
		goto L15
	} else {
		goto L72
	}
L72:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v192 != 0 {
		goto L6
	} else {
		goto L73
	}
L73:
	;
	v193 = F_get_rel_name(m, v120)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v193
	goto L15
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v205 = F_getObjectDescription(m, v14+int32(164), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v205
	F_errmsg_internal(m, int32(_a_F_RememberAllDependentForRebuilding_6), v14+int32(48))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_RememberAllDependentForRebuilding_3), int32(_a_F_RememberAllDependentForRebuilding_7), int32(_a_F_RememberAllDependentForRebuilding_5))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errmsg(m, int32(_a_F_RememberAllDependentForRebuilding_8), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v232 = F_getObjectDescription(m, v14+int32(164), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v232
	F_errdetail(m, int32(_a_F_RememberAllDependentForRebuilding_2), v14-int32(-64))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_RememberAllDependentForRebuilding_3), int32(_a_F_RememberAllDependentForRebuilding_9), int32(_a_F_RememberAllDependentForRebuilding_5))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errmsg(m, int32(_a_F_RememberAllDependentForRebuilding_10), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v260 = F_getObjectDescription(m, v14+int32(164), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v260
	F_errdetail(m, int32(_a_F_RememberAllDependentForRebuilding_2), v14+int32(80))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_RememberAllDependentForRebuilding_3), int32(_a_F_RememberAllDependentForRebuilding_11), int32(_a_F_RememberAllDependentForRebuilding_5))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errmsg(m, int32(_a_F_RememberAllDependentForRebuilding_12), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v288 = F_getObjectDescription(m, v14+int32(164), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+100)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v288
	F_errdetail(m, int32(_a_F_RememberAllDependentForRebuilding_2), v14+int32(96))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_RememberAllDependentForRebuilding_3), int32(_a_F_RememberAllDependentForRebuilding_13), int32(_a_F_RememberAllDependentForRebuilding_5))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(_a_F_RememberAllDependentForRebuilding_14), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v316 = F_getObjectDescription(m, v14+int32(164), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+116)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v316
	F_errdetail(m, int32(_a_F_RememberAllDependentForRebuilding_2), v14+int32(112))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_RememberAllDependentForRebuilding_3), int32(_a_F_RememberAllDependentForRebuilding_15), int32(_a_F_RememberAllDependentForRebuilding_5))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v14)+156))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	if v334 == v335 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L116
	}
L110:
	;
	if l1 != int32(24) {
		goto L15
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	if l1 != int32(24) {
		goto L15
	} else {
		goto L115
	}
L113:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v14)+160))
	if v337 != l3 {
		goto L109
	} else {
		goto L114
	}
L114:
	;
	goto L15
L115:
	;
	goto L109
L116:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_errmsg(m, int32(_a_F_RememberAllDependentForRebuilding_16), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v14)+156))
	v351 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+160)))
	v353 = F_get_attname(m, v350, v351, int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+132)) = v353
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = l4
	F_errdetail(m, int32(_a_F_RememberAllDependentForRebuilding_17), v14+int32(128))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_RememberAllDependentForRebuilding_3), int32(_a_F_RememberAllDependentForRebuilding_18), int32(_a_F_RememberAllDependentForRebuilding_5))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	if v406 != 0 {
		goto L15
	} else {
		goto L135
	}
L123:
	;
	v406 = int32(0)
	goto L122
L124:
	;
	goto L125
L125:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	if v374 <= int32(0) {
		v400 = v368
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v406 = v400
	goto L122
L127:
	;
	v377 = int32(0)
	if v377 < v374 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v380 = v374
	goto L130
L129:
	;
	v380 = v377
	goto L130
L130:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v367)+12))
	v383 = int32(0)
	goto L131
L131:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v381+v383<<(uint(int32(2))%32))))
	v392 = base.B2i32(v391 == v69)
	if v391 == v69 {
		v400 = v392
		goto L126
	} else {
		goto L133
	}
L132:
	;
	v400 = v392
	goto L126
L133:
	;
	v394 = v383 + int32(1)
	if v394 != v380 {
		v383 = v394
		goto L131
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	v407 = int32(0)
	v409 = F_pg_get_statisticsobj_worker(m, v69, v407, v407)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v412 = F_lappend_oid(m, v411, v69)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v412
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v416 = F_lappend(m, v415, v409)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v416
	goto L15
L139:
	;
	v426 = F_getObjectDescription(m, v14+int32(164), int32(0))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v426
	F_errmsg_internal(m, int32(_a_F_RememberAllDependentForRebuilding_6), v14)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_RememberAllDependentForRebuilding_3), int32(_a_F_RememberAllDependentForRebuilding_19), int32(_a_F_RememberAllDependentForRebuilding_5))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
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
	goto L15
L144:
	;
	if v441 != 0 {
		v59 = v441
		goto L13
	} else {
		goto L145
	}
L145:
	;
	goto L14
L146:
	;
	F_relation_close(m, v18, int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	m.G0 = v14 + int32(320)
	return
L148:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v466
	F_errmsg_internal(m, int32(_a_F_RememberAllDependentForRebuilding_20), v14+int32(32))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_RememberAllDependentForRebuilding_3), int32(_a_F_RememberAllDependentForRebuilding_21), int32(_a_F_RememberAllDependentForRebuilding_22))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v482
	F_errmsg_internal(m, int32(_a_F_RememberAllDependentForRebuilding_23), v14+int32(16))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(_a_F_RememberAllDependentForRebuilding_3), int32(_a_F_RememberAllDependentForRebuilding_24), int32(_a_F_RememberAllDependentForRebuilding_25))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RenameTypeInternal(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v27 int32
	_ = v27
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
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v16 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v20 = F_SearchSysCacheCopy(m, int32(82), l0, int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			if v20 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
				v24 = v22 + v23
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+96))
				v27 = int32(0)
				v29 = F_GetSysCacheOid(m, int32(81), l1, l2, v27, v27)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					if v29 != 0 {
						v31 = F_get_typisdefined(m, v29)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							if v31 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return
								} else {
									F_errcode(m, int32(_a_F_RenameTypeInternal_0))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
										F_errmsg(m, int32(_a_F_RenameTypeInternal_1), v12+int32(16))
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_RenameTypeInternal_2), int32(805), int32(_a_F_RenameTypeInternal_3))
											mBase = m.M
											v108 = m.ExcPending
											if v108 != 0 {
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
								v35 = F_moveArrayTypeName(m, v29, l1, l2)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									if v35 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return
										} else {
											F_errcode(m, int32(_a_F_RenameTypeInternal_0))
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
												F_errmsg(m, int32(_a_F_RenameTypeInternal_1), v12+int32(16))
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_RenameTypeInternal_2), int32(805), int32(_a_F_RenameTypeInternal_3))
													mBase = m.M
													v108 = m.ExcPending
													if v108 != 0 {
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
										v42 = F_strncpy(m, v24+int32(4), l1, int32(64))
										mBase = m.M
										v43 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v42)+63)) = uint8(v43)
										F_CatalogTupleUpdate(m, v16, v20+int32(4), v20)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											v50 = *(*int32)(unsafe.Add(mBase, _c_F_RenameTypeInternal[0]))
											if v50 != 0 {
												v52 = int32(0)
												F_RunObjectPostAlterHook(m, int32(1247), l0, v52, v52, v52)
												mBase = m.M
												v56 = m.ExcPending
												if v56 != 0 {
													return
												} else {
													F_pfree(m, v20)
													mBase = m.M
													v58 = m.ExcPending
													if v58 != 0 {
														return
													} else {
														F_relation_close(m, v16, int32(3))
														mBase = m.M
														v61 = m.ExcPending
														if v61 != 0 {
															return
														} else {
															v62 = int32(0)
															if base.B2i32(v25 == v62)|base.B2i32(v25 == v29) == v62 {
																v68 = F_makeArrayTypeName(m, l1, l2)
																mBase = m.M
																v69 = m.ExcPending
																if v69 != 0 {
																	return
																} else {
																	F_RenameTypeInternal(m, v25, v68, l2)
																	mBase = m.M
																	v71 = m.ExcPending
																	if v71 != 0 {
																		return
																	} else {
																		F_pfree(m, v68)
																		mBase = m.M
																		v73 = m.ExcPending
																		if v73 != 0 {
																			return
																		} else {
																			m.G0 = v12 + int32(32)
																			return
																		}
																	}
																}
															} else {
																m.G0 = v12 + int32(32)
																return
															}
														}
													}
												}
											} else {
												F_pfree(m, v20)
												mBase = m.M
												v58 = m.ExcPending
												if v58 != 0 {
													return
												} else {
													F_relation_close(m, v16, int32(3))
													mBase = m.M
													v61 = m.ExcPending
													if v61 != 0 {
														return
													} else {
														v62 = int32(0)
														if base.B2i32(v25 == v62)|base.B2i32(v25 == v29) == v62 {
															v68 = F_makeArrayTypeName(m, l1, l2)
															mBase = m.M
															v69 = m.ExcPending
															if v69 != 0 {
																return
															} else {
																F_RenameTypeInternal(m, v25, v68, l2)
																mBase = m.M
																v71 = m.ExcPending
																if v71 != 0 {
																	return
																} else {
																	F_pfree(m, v68)
																	mBase = m.M
																	v73 = m.ExcPending
																	if v73 != 0 {
																		return
																	} else {
																		m.G0 = v12 + int32(32)
																		return
																	}
																}
															}
														} else {
															m.G0 = v12 + int32(32)
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
					} else {
						v42 = F_strncpy(m, v24+int32(4), l1, int32(64))
						mBase = m.M
						v43 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v42)+63)) = uint8(v43)
						F_CatalogTupleUpdate(m, v16, v20+int32(4), v20)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, _c_F_RenameTypeInternal[0]))
							if v50 != 0 {
								v52 = int32(0)
								F_RunObjectPostAlterHook(m, int32(1247), l0, v52, v52, v52)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									F_pfree(m, v20)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return
									} else {
										F_relation_close(m, v16, int32(3))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return
										} else {
											v62 = int32(0)
											if base.B2i32(v25 == v62)|base.B2i32(v25 == v29) == v62 {
												v68 = F_makeArrayTypeName(m, l1, l2)
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return
												} else {
													F_RenameTypeInternal(m, v25, v68, l2)
													mBase = m.M
													v71 = m.ExcPending
													if v71 != 0 {
														return
													} else {
														F_pfree(m, v68)
														mBase = m.M
														v73 = m.ExcPending
														if v73 != 0 {
															return
														} else {
															m.G0 = v12 + int32(32)
															return
														}
													}
												}
											} else {
												m.G0 = v12 + int32(32)
												return
											}
										}
									}
								}
							} else {
								F_pfree(m, v20)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									F_relation_close(m, v16, int32(3))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										v62 = int32(0)
										if base.B2i32(v25 == v62)|base.B2i32(v25 == v29) == v62 {
											v68 = F_makeArrayTypeName(m, l1, l2)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return
											} else {
												F_RenameTypeInternal(m, v25, v68, l2)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return
												} else {
													F_pfree(m, v68)
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return
													} else {
														m.G0 = v12 + int32(32)
														return
													}
												}
											}
										} else {
											m.G0 = v12 + int32(32)
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
				v81 = m.ExcPending
				if v81 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
					F_errmsg_internal(m, int32(_a_F_RenameTypeInternal_4), v12)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_RenameTypeInternal_2), int32(777), int32(_a_F_RenameTypeInternal_3))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
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
func F_ReplSlotSyncWorkerMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int64
	_ = v147
	var v149 int64
	_ = v149
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v171 int32
	_ = v171
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int64
	_ = v193
	var v195 int64
	_ = v195
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v217 int32
	_ = v217
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int64
	_ = v239
	var v241 int64
	_ = v241
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int64
	_ = v285
	var v287 int64
	_ = v287
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v309 int32
	_ = v309
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int64
	_ = v331
	var v333 int64
	_ = v333
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v355 int32
	_ = v355
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int64
	_ = v377
	var v379 int64
	_ = v379
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v401 int32
	_ = v401
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int64
	_ = v423
	var v425 int64
	_ = v425
	var v433 int32
	_ = v433
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
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int64
	_ = v470
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int64
	_ = v523
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
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
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v849 int32
	_ = v849
	var v860 int32
	_ = v860
	var v861 int64
	_ = v861
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	v9 = m.G0
	v11 = v9 - int32(240)
	m.G0 = v11
	v15 = int32(-1)
	v16 = int32(0)
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v15 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v860 = int32(m.ExcTag)
	v861 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v860 == int32(0) {
		goto L241
	} else {
		goto L242
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[0])) = int32(7)
	goto L11
L8:
	;
	v45 = v16
	goto L9
L9:
	;
	if v45 != 0 {
		goto L20
	} else {
		goto L21
	}
L10:
	;
	F_InitProcess(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L14
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[0]))
	v32 = F_GetBackendTypeDesc(m, v31)
	mBase = m.M
	goto L13
L13:
	;
	goto L10
L14:
	;
	F_BaseInit(m)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L16
L16:
	;
	v38 = v11 + int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v11 + int32(60)
	goto L19
L17:
	;
	v45 = int32(0)
	goto L9
L19:
	;
	goto L17
L20:
	;
	v46 = int32(_a_F_ReplSlotSyncWorkerMain_0)
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[1])) = v48 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[2])) = int32(0)
	F_EmitErrorReport(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[3])) = v11 + int32(80)
	v65 = int32(914)
	v67 = m.G0
	v69 = v67 - int32(32)
	m.G0 = v69
	switch int32(916) {
	case 0, 2:
		v79 = v65
		goto L26
	default:
		goto L27
	}
L23:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	goto L3
L25:
	;
	v111 = int32(915)
	v113 = m.G0
	v115 = v113 - int32(32)
	m.G0 = v115
	switch int32(917) {
	case 0, 2:
		v125 = v111
		goto L39
	default:
		goto L40
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = v79
	F_sigemptyset(m, v69+int32(16))
	mBase = m.M
	goto L29
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[4])) = v65
	v79 = int32(_a_F_ReplSlotSyncWorkerMain_1)
	goto L26
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+24)) = int32(268435456)
	v91 = v69 + int32(12)
	goto L33
L31:
	;
	m.G0 = v69 + int32(32)
	goto L25
L33:
	;
	goto L34
L34:
	;
	if v91 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v97 = int32(20)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[5])) = v99
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v91)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[6])) = v101
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
	*(*int64)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[7])) = v103
	goto L37
L36:
	;
	goto L37
L37:
	;
	goto L31
L38:
	;
	v157 = int32(295)
	v159 = m.G0
	v161 = v159 - int32(32)
	m.G0 = v161
	switch int32(297) {
	case 0, 2:
		v171 = v157
		goto L52
	default:
		goto L53
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115)+12)) = v125
	F_sigemptyset(m, v115+int32(16))
	mBase = m.M
	goto L42
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[8])) = v111
	v125 = int32(_a_F_ReplSlotSyncWorkerMain_1)
	goto L39
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115)+24)) = int32(268435456)
	v137 = v115 + int32(12)
	goto L46
L44:
	;
	m.G0 = v115 + int32(32)
	goto L38
L46:
	;
	goto L47
L47:
	;
	if v137 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v144 = int32(40)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[9])) = v145
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v137)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[10])) = v147
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v137)))
	*(*int64)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[11])) = v149
	goto L50
L49:
	;
	goto L50
L50:
	;
	goto L44
L51:
	;
	v203 = int32(919)
	v205 = m.G0
	v207 = v205 - int32(32)
	m.G0 = v207
	switch int32(921) {
	case 0, 2:
		v217 = v203
		goto L65
	default:
		goto L66
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161)+12)) = v171
	F_sigemptyset(m, v161+int32(16))
	mBase = m.M
	goto L55
L53:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[12])) = v157
	v171 = int32(_a_F_ReplSlotSyncWorkerMain_1)
	goto L52
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161)+24)) = int32(268435456)
	v183 = v161 + int32(12)
	goto L59
L57:
	;
	m.G0 = v161 + int32(32)
	goto L51
L59:
	;
	goto L60
L60:
	;
	if v183 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v190 = int32(300)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v183)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[13])) = v191
	v193 = *(*int64)(unsafe.Add(mBase, uint32(v183)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[14])) = v193
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v183)))
	*(*int64)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[15])) = v195
	goto L63
L62:
	;
	goto L63
L63:
	;
	goto L57
L64:
	;
	v249 = int32(917)
	v251 = m.G0
	v253 = v251 - int32(32)
	m.G0 = v253
	switch int32(919) {
	case 0, 2:
		v263 = v249
		goto L78
	default:
		goto L79
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207)+12)) = v217
	F_sigemptyset(m, v207+int32(16))
	mBase = m.M
	goto L68
L66:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[16])) = v203
	v217 = int32(_a_F_ReplSlotSyncWorkerMain_1)
	goto L65
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207)+24)) = int32(268435456)
	v229 = v207 + int32(12)
	goto L72
L70:
	;
	m.G0 = v207 + int32(32)
	goto L64
L72:
	;
	goto L73
L73:
	;
	if v229 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v236 = int32(160)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v229)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[17])) = v237
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v229)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[18])) = v239
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v229)))
	*(*int64)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[19])) = v241
	goto L76
L75:
	;
	goto L76
L76:
	;
	goto L70
L77:
	;
	v295 = int32(-2)
	v297 = m.G0
	v299 = v297 - int32(32)
	m.G0 = v299
	switch int32(0) {
	case 0, 2:
		v309 = v295
		goto L91
	default:
		goto L92
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v253)+12)) = v263
	F_sigemptyset(m, v253+int32(16))
	mBase = m.M
	goto L81
L79:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[20])) = v249
	v263 = int32(_a_F_ReplSlotSyncWorkerMain_1)
	goto L78
L81:
	;
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v253)+24)) = int32(268435456)
	v275 = v253 + int32(12)
	goto L85
L83:
	;
	m.G0 = v253 + int32(32)
	goto L77
L85:
	;
	goto L86
L86:
	;
	if v275 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v282 = int32(200)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v275)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[21])) = v283
	v285 = *(*int64)(unsafe.Add(mBase, uint32(v275)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[22])) = v285
	v287 = *(*int64)(unsafe.Add(mBase, uint32(v275)))
	*(*int64)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[23])) = v287
	goto L89
L88:
	;
	goto L89
L89:
	;
	goto L83
L90:
	;
	v341 = int32(-2)
	v343 = m.G0
	v345 = v343 - int32(32)
	m.G0 = v345
	switch int32(0) {
	case 0, 2:
		v355 = v341
		goto L104
	default:
		goto L105
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v299)+12)) = v309
	F_sigemptyset(m, v299+int32(16))
	mBase = m.M
	goto L94
L92:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[24])) = v295
	v309 = int32(_a_F_ReplSlotSyncWorkerMain_1)
	goto L91
L94:
	;
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v299)+24)) = int32(268435456)
	v321 = v299 + int32(12)
	goto L98
L96:
	;
	m.G0 = v299 + int32(32)
	goto L90
L98:
	;
	goto L99
L99:
	;
	if v321 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v328 = int32(240)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v321)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[25])) = v329
	v331 = *(*int64)(unsafe.Add(mBase, uint32(v321)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[26])) = v331
	v333 = *(*int64)(unsafe.Add(mBase, uint32(v321)))
	*(*int64)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[27])) = v333
	goto L102
L101:
	;
	goto L102
L102:
	;
	goto L96
L103:
	;
	v387 = int32(0)
	v389 = m.G0
	v391 = v389 - int32(32)
	m.G0 = v391
	switch int32(2) {
	case 0, 2:
		v401 = v387
		goto L117
	default:
		goto L118
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v345)+12)) = v355
	F_sigemptyset(m, v345+int32(16))
	mBase = m.M
	goto L107
L105:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[28])) = v341
	v355 = int32(_a_F_ReplSlotSyncWorkerMain_1)
	goto L104
L107:
	;
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v345)+24)) = int32(268435456)
	v367 = v345 + int32(12)
	goto L111
L109:
	;
	m.G0 = v345 + int32(32)
	goto L103
L111:
	;
	goto L112
L112:
	;
	if v367 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v374 = int32(260)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v367)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[29])) = v375
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v367)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[30])) = v377
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v367)))
	*(*int64)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[31])) = v379
	goto L115
L114:
	;
	goto L115
L115:
	;
	goto L109
L116:
	;
	v433 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[32]))
	F_check_and_set_sync_info(m, v433)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L6
	} else {
		goto L129
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391)+12)) = v401
	F_sigemptyset(m, v391+int32(16))
	mBase = m.M
	goto L119
L118:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[33])) = v387
	v401 = int32(_a_F_ReplSlotSyncWorkerMain_1)
	goto L117
L119:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391)+24)) = int32(268435457)
	v413 = v391 + int32(12)
	goto L124
L122:
	;
	m.G0 = v391 + int32(32)
	goto L116
L124:
	;
	goto L125
L125:
	;
	if v413 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v420 = int32(340)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v413)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[34])) = v421
	v423 = *(*int64)(unsafe.Add(mBase, uint32(v413)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[35])) = v423
	v425 = *(*int64)(unsafe.Add(mBase, uint32(v413)))
	*(*int64)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[36])) = v425
	goto L128
L127:
	;
	goto L128
L128:
	;
	goto L122
L129:
	;
	v438 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L6
	} else {
		goto L130
	}
L130:
	;
	if v438 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	F_errmsg(m, int32(_a_F_ReplSlotSyncWorkerMain_2), int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L6
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	F_before_shmem_exit(m, int32(1020), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L6
	} else {
		goto L136
	}
L134:
	;
	F_errfinish(m, int32(_a_F_ReplSlotSyncWorkerMain_3), int32(1425), int32(_a_F_ReplSlotSyncWorkerMain_4))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L6
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	v453 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[37])) = v453
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[38])) = v453
	v463 = v453
	goto L138
L137:
	;
	F_load_file(m, int32(_a_F_ReplSlotSyncWorkerMain_5), int32(0))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L6
	} else {
		goto L143
	}
L138:
	;
	v465 = int32(40)
	v466 = v463 * v465
	v467 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v466)+uint32(_c_F_ReplSlotSyncWorkerMain[39]))) = uint8(v467)
	*(*int32)(unsafe.Add(mBase, uint32(v466)+uint32(_c_F_ReplSlotSyncWorkerMain[40]))) = v463
	v470 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v466)+uint32(_c_F_ReplSlotSyncWorkerMain[41]))) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v466)+uint32(_c_F_ReplSlotSyncWorkerMain[42]))) = v467
	*(*int64)(unsafe.Add(mBase, uint32(v466)+uint32(_c_F_ReplSlotSyncWorkerMain[43]))) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v466)+uint32(_c_F_ReplSlotSyncWorkerMain[44]))) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v466)+uint32(_c_F_ReplSlotSyncWorkerMain[45]))) = uint8(v467)
	v481 = v463 | int32(1)
	v483 = v481 * v465
	*(*uint8)(unsafe.Add(mBase, uint32(v483)+uint32(_c_F_ReplSlotSyncWorkerMain[39]))) = uint8(v467)
	*(*int32)(unsafe.Add(mBase, uint32(v483)+uint32(_c_F_ReplSlotSyncWorkerMain[40]))) = v481
	*(*int64)(unsafe.Add(mBase, uint32(v483)+uint32(_c_F_ReplSlotSyncWorkerMain[41]))) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v483)+uint32(_c_F_ReplSlotSyncWorkerMain[42]))) = v467
	*(*int64)(unsafe.Add(mBase, uint32(v483)+uint32(_c_F_ReplSlotSyncWorkerMain[43]))) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v483)+uint32(_c_F_ReplSlotSyncWorkerMain[44]))) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v483)+uint32(_c_F_ReplSlotSyncWorkerMain[45]))) = uint8(v467)
	v498 = v463 | int32(2)
	v500 = v498 * v465
	*(*uint8)(unsafe.Add(mBase, uint32(v500)+uint32(_c_F_ReplSlotSyncWorkerMain[39]))) = uint8(v467)
	*(*int32)(unsafe.Add(mBase, uint32(v500)+uint32(_c_F_ReplSlotSyncWorkerMain[40]))) = v498
	*(*int64)(unsafe.Add(mBase, uint32(v500)+uint32(_c_F_ReplSlotSyncWorkerMain[41]))) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v500)+uint32(_c_F_ReplSlotSyncWorkerMain[42]))) = v467
	*(*int64)(unsafe.Add(mBase, uint32(v500)+uint32(_c_F_ReplSlotSyncWorkerMain[43]))) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v500)+uint32(_c_F_ReplSlotSyncWorkerMain[44]))) = v467
	*(*uint8)(unsafe.Add(mBase, uint32(v500)+uint32(_c_F_ReplSlotSyncWorkerMain[45]))) = uint8(v467)
	if v463 != int32(20) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v536 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[46])) = uint8(v536)
	F_pqsignal_be(m, int32(14), int32(1769))
	mBase = m.M
	goto L137
L140:
	;
	v517 = v463 | int32(3)
	v519 = v517 * int32(40)
	v520 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v519)+uint32(_c_F_ReplSlotSyncWorkerMain[39]))) = uint8(v520)
	*(*int32)(unsafe.Add(mBase, uint32(v519)+uint32(_c_F_ReplSlotSyncWorkerMain[40]))) = v517
	v523 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v519)+uint32(_c_F_ReplSlotSyncWorkerMain[41]))) = v523
	*(*int32)(unsafe.Add(mBase, uint32(v519)+uint32(_c_F_ReplSlotSyncWorkerMain[42]))) = v520
	*(*int64)(unsafe.Add(mBase, uint32(v519)+uint32(_c_F_ReplSlotSyncWorkerMain[43]))) = v523
	*(*int32)(unsafe.Add(mBase, uint32(v519)+uint32(_c_F_ReplSlotSyncWorkerMain[44]))) = v520
	*(*uint8)(unsafe.Add(mBase, uint32(v519)+uint32(_c_F_ReplSlotSyncWorkerMain[45]))) = uint8(v520)
	v463 = v463 + int32(4)
	goto L138
L141:
	;
	goto L142
L142:
	;
	goto L139
L143:
	;
	F_sigprocmask(m, int32(_a_F_ReplSlotSyncWorkerMain_6), int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L6
	} else {
		goto L144
	}
L144:
	;
	F_SetConfigOption(m, int32(_a_F_ReplSlotSyncWorkerMain_7), int32(_a_F_ReplSlotSyncWorkerMain_8), int32(5), int32(10))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L6
	} else {
		goto L145
	}
L145:
	;
	v556 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[47]))
	v558 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[48]))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v558)+20))
	v560 = m.T0[v559].(func(*base.Module, int32) int32)(m, v556)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L6
	} else {
		goto L146
	}
L146:
	;
	if v560 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L6
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v583 = int32(0)
	F_InitPostgres(m, v560, v583, v583, v583, v583, v583)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L6
	} else {
		goto L154
	}
L150:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L6
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(_a_F_ReplSlotSyncWorkerMain_9)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_ReplSlotSyncWorkerMain_10)
	F_errmsg(m, int32(_a_F_ReplSlotSyncWorkerMain_11), v11)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L6
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(_a_F_ReplSlotSyncWorkerMain_3), int32(1052), int32(_a_F_ReplSlotSyncWorkerMain_12))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L6
	} else {
		goto L153
	}
L153:
	;
	goto L3
L154:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[49])) = int32(2)
	v594 = v11 - int32(-64)
	F_initStringInfo(m, v594)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L6
	} else {
		goto L155
	}
L155:
	;
	v598 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[50]))
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598))))
	if v599 != 0 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v614 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[47]))
	v615 = int32(0)
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v622 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[48]))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v622)))
	v624 = m.T0[v623].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v614, v615, v615, v615, v618, v11+int32(236))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L6
	} else {
		goto L162
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = int32(_a_F_ReplSlotSyncWorkerMain_13)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v598
	F_appendStringInfo(m, v594, int32(_a_F_ReplSlotSyncWorkerMain_14), v11+int32(48))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L6
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	F_appendStringInfoString(m, v11-int32(-64), int32(_a_F_ReplSlotSyncWorkerMain_13))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L6
	} else {
		goto L161
	}
L160:
	;
	goto L156
L161:
	;
	goto L156
L162:
	;
	if v624 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L6
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	F_pfree(m, v649)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L6
	} else {
		goto L170
	}
L166:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L6
	} else {
		goto L167
	}
L167:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v11)+236))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v636
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v635
	F_errmsg(m, int32(_a_F_ReplSlotSyncWorkerMain_15), v11+int32(16))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L6
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_ReplSlotSyncWorkerMain_3), int32(1488), int32(_a_F_ReplSlotSyncWorkerMain_4))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L6
	} else {
		goto L169
	}
L169:
	;
	goto L3
L170:
	;
	F_before_shmem_exit(m, int32(1021), v624)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L6
	} else {
		goto L171
	}
L171:
	;
	F_validate_remote_info(m, v624)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L6
	} else {
		goto L172
	}
L172:
	;
	goto L173
L173:
	;
	v666 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[51]))
	if v666 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L6
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v670 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[52]))
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670)+4)))
	if v671 == int32(1) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	goto L177
L179:
	;
	v676 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L6
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v691 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[53]))
	if v691 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L182:
	;
	if v676 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	F_errmsg(m, int32(_a_F_ReplSlotSyncWorkerMain_16), int32(0))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L6
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L6
	} else {
		goto L188
	}
L186:
	;
	F_errfinish(m, int32(_a_F_ReplSlotSyncWorkerMain_3), int32(1184), int32(_a_F_ReplSlotSyncWorkerMain_17))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L6
	} else {
		goto L187
	}
L187:
	;
	goto L185
L188:
	;
	goto L3
L189:
	;
	v824 = F_synchronize_slots(m, v624)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L6
	} else {
		goto L231
	}
L190:
	;
	v695 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[47]))
	v696 = F_pstrdup(m, v695)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L6
	} else {
		goto L191
	}
L191:
	;
	v699 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[54]))
	v700 = F_pstrdup(m, v699)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L6
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[53])) = int32(0)
	v706 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[55])))
	v708 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[56])))
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L6
	} else {
		goto L193
	}
L193:
	;
	v713 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[47]))
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696))))
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713))))
	if base.B2i32(v716 == int32(0))|base.B2i32(v716 != v719) != 0 {
		v737 = v716
		v738 = v719
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v741 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[54]))
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700))))
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v741))))
	if base.B2i32(v744 == int32(0))|base.B2i32(v744 != v747) != 0 {
		v765 = v744
		v766 = v747
		goto L202
	} else {
		goto L203
	}
L195:
	;
	goto L194
L196:
	;
	v722 = v696
	v723 = v713
	goto L197
L197:
	;
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v723)+1)))
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722)+1)))
	if v727 == int32(0) {
		v737 = v727
		v738 = v726
		goto L195
	} else {
		goto L199
	}
L198:
	;
	v737 = v727
	v738 = v726
	goto L195
L199:
	;
	v730 = int32(1)
	if v727 == v726 {
		v722 = v722 + v730
		v723 = v723 + v730
		goto L197
	} else {
		goto L200
	}
L200:
	;
	goto L198
L201:
	;
	F_pfree(m, v696)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L6
	} else {
		goto L208
	}
L202:
	;
	goto L201
L203:
	;
	v750 = v700
	v751 = v741
	goto L204
L204:
	;
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751)+1)))
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+1)))
	if v755 == int32(0) {
		v765 = v755
		v766 = v754
		goto L202
	} else {
		goto L206
	}
L205:
	;
	v765 = v755
	v766 = v754
	goto L202
L206:
	;
	v758 = int32(1)
	if v755 == v754 {
		v750 = v750 + v758
		v751 = v751 + v758
		goto L204
	} else {
		goto L207
	}
L207:
	;
	goto L205
L208:
	;
	F_pfree(m, v700)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L6
	} else {
		goto L209
	}
L209:
	;
	v773 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[56])))
	if v773 != v708 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v777 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L6
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	if v737-v738|(v765-v766) == int32(0) {
		goto L220
	} else {
		goto L221
	}
L213:
	;
	if v777 != 0 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(_a_F_ReplSlotSyncWorkerMain_18)
	F_errmsg(m, int32(_a_F_ReplSlotSyncWorkerMain_19), v11+int32(32))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L6
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L6
	} else {
		goto L219
	}
L217:
	;
	F_errfinish(m, int32(_a_F_ReplSlotSyncWorkerMain_3), int32(1151), int32(_a_F_ReplSlotSyncWorkerMain_20))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L6
	} else {
		goto L218
	}
L218:
	;
	goto L216
L219:
	;
	goto L3
L220:
	;
	v798 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[55])))
	if v706 == v798 {
		goto L189
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v802 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L6
	} else {
		goto L224
	}
L223:
	;
	goto L222
L224:
	;
	if v802 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	F_errmsg(m, int32(_a_F_ReplSlotSyncWorkerMain_21), int32(0))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L6
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	v814 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[52]))
	*(*int64)(unsafe.Add(mBase, uint32(v814)+8)) = int64(0)
	F_proc_exit(m, int32(0))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L6
	} else {
		goto L230
	}
L228:
	;
	F_errfinish(m, int32(_a_F_ReplSlotSyncWorkerMain_3), int32(1160), int32(_a_F_ReplSlotSyncWorkerMain_20))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L6
	} else {
		goto L229
	}
L229:
	;
	goto L227
L230:
	;
	goto L3
L231:
	;
	v826 = int32(_a_F_ReplSlotSyncWorkerMain_22)
	v828 = int32(_a_F_ReplSlotSyncWorkerMain_23)
	v830 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[57]))
	v832 = v830 << (uint(int32(1)) % 32)
	if v828 <= v832 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v835 = v828
	goto L234
L233:
	;
	v835 = v832
	goto L234
L234:
	;
	if v824 != 0 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v836 = int32(200)
	goto L237
L236:
	;
	v836 = v835
	goto L237
L237:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[57])) = v836
	v839 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[58]))
	v842 = F_WaitLatch(m, v839, int32(41), v836, int32(83886091))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L6
	} else {
		goto L238
	}
L238:
	;
	if v842&int32(1) == int32(0) {
		goto L173
	} else {
		goto L239
	}
L239:
	;
	v849 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[58]))
	*(*int32)(unsafe.Add(mBase, uint32(v849))) = int32(0)
	goto L240
L240:
	;
	goto L173
L241:
	;
	v865 = int32(v861)
	m.G0 = v11
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v865)+4))
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v865)))
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
	if v11+int32(60) == v871 {
		goto L244
	} else {
		goto L245
	}
L242:
	;
	m.ExcPending = 1
	goto L250
L243:
	;
	if v875 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L244:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v868)+4))
	v875 = v873
	goto L246
L245:
	;
	v875 = int32(0)
	goto L246
L246:
	;
	goto L243
L247:
	;
	F___wasm_longjmp(m, v868, v867)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L250
	} else {
		goto L251
	}
L248:
	;
	goto L249
L249:
	;
	v15 = v875
	v16 = v867
	goto L1
L250:
	;
	return
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RequestXLogSwitch(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	F_XLogBeginInsert(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		if l0 != 0 {
			v7 = int32(_a_F_RequestXLogSwitch_0)
			v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RequestXLogSwitch[0])))
			v10 = v9 | int32(2)
			*(*uint8)(unsafe.Add(mBase, _c_F_RequestXLogSwitch[0])) = uint8(v10)
		} else {
		}
		v14 = F_XLogInsert(m, int32(0), int32(64))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			return v14
		}
	}
}
func F_ResetUnloggedRelationsInTablespaceDir(m *base.Module, l0 int32, l1 int32) {
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
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v165 int64
	_ = v165
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int64
	_ = v224
	var v226 int64
	_ = v226
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
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
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v623 int32
	_ = v623
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v715 int32
	_ = v715
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v737 int32
	_ = v737
	var v744 int32
	_ = v744
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v819 int32
	_ = v819
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v892 int32
	_ = v892
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v934 int32
	_ = v934
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v949 int32
	_ = v949
	var v956 int32
	_ = v956
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1031 int32
	_ = v1031
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1073 int32
	_ = v1073
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1090 int32
	_ = v1090
	var v1095 int32
	_ = v1095
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1117 int32
	_ = v1117
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1143 int32
	_ = v1143
	var v1148 int32
	_ = v1148
	v10 = m.G0
	v12 = v10 - int32(_a_F_ResetUnloggedRelationsInTablespaceDir_0)
	m.G0 = v12
	v14 = F_AllocateDir(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L4
	} else {
		goto L232
	}
L2:
	;
	m.G0 = v12 + int32(_a_F_ResetUnloggedRelationsInTablespaceDir_0)
	return
L3:
	;
	v37 = F_ReadDir(m, v14, l0)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L13
	}
L4:
	;
	return
L5:
	;
	if v14 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0]))
	if v17 != int32(44) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v22 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	if v22 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg(m, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_1), v12)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_2), int32(127), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_3))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	goto L2
L13:
	;
	if v37 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v40 = l1 & int32(1)
	v42 = l1 & int32(2)
	v46 = v37
	goto L17
L15:
	;
	goto L16
L16:
	;
	F_FreeDir(m, v14)
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L4
	} else {
		goto L231
	}
L17:
	;
	v53 = v46 + int32(19)
	v54 = int32(_a_F_ResetUnloggedRelationsInTablespaceDir_4)
	v58 = m.G0
	v60 = v58 - int32(32)
	v61 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v60)+16)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v60)+8)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v60))) = v61
	v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[1])))
	if v69 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L16
L19:
	;
	v1105 = F_ReadDir(m, v14, l0)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L4
	} else {
		goto L229
	}
L20:
	;
	v138 = F_strlen(m, v53)
	mBase = m.M
	if v137 != v138 {
		goto L19
	} else {
		goto L39
	}
L21:
	;
	v137 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[2])))
	if v73 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v77 = v53
	goto L27
L25:
	;
	goto L26
L26:
	;
	v87 = v54
	v88 = v69
	goto L30
L27:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v83 == v69 {
		v77 = v77 + int32(1)
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v137 = v77 - v53
	goto L20
L29:
	;
	goto L28
L30:
	;
	v95 = v60 + int32(base.Ui32(v88)>>(uint(int32(3))%32))&int32(28)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v97 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = v96 | v97<<(uint(v88)%32)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	if v101 != 0 {
		v87 = v87 + v97
		v88 = v101
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v104 == int32(0) {
		v127 = v53
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L31
L33:
	;
	v137 = v127 - v53
	goto L20
L34:
	;
	v108 = v53
	v109 = v104
	goto L35
L35:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v60+int32(base.Ui32(v109)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v117)>>(uint(v109)%32))&int32(1) == int32(0) {
		v127 = v108
		goto L33
	} else {
		goto L37
	}
L36:
	;
	v127 = v125
	goto L33
L37:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
	v125 = v108 + int32(1)
	if v123 != 0 {
		v108 = v125
		v109 = v123
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = l0
	v143 = v12 + int32(208)
	v148 = F_pg_snprintf(m, v143, int32(2048), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_5), v12+int32(192))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	if v42 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v700 = v12 + int32(208)
	v701 = F_AllocateDir(m, v700)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L4
	} else {
		goto L147
	}
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+3296)) = int64(17179869188)
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+3320)) = v275
	v282 = F_hash_create(m, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_6), int32(32), v12+int32(3280), int32(1064))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L4
	} else {
		goto L67
	}
L43:
	;
	v157 = m.G0
	v159 = v157 - int32(16)
	m.G0 = v159
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[4]))
	if v162 != 0 {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	goto L45
L45:
	;
	if v40 == int32(0) {
		goto L19
	} else {
		goto L57
	}
L46:
	;
	if v40 != 0 {
		goto L42
	} else {
		goto L56
	}
L47:
	;
	if base.B2i32(v162 != int32(0)) == int32(0) {
		goto L46
	} else {
		goto L51
	}
L48:
	;
	v163 = F_GetCurrentTimestamp(m)
	mBase = m.M
	v165 = *(*int64)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[5]))
	F_TimestampDifference(m, v165, v163, v159+int32(12), v159+int32(8))
	mBase = m.M
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[6]))) = v171
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(3280)))) = v173
	*(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[4])) = int32(0)
	goto L50
L49:
	;
	goto L50
L50:
	;
	m.G0 = v159 + int32(16)
	goto L47
L51:
	;
	v188 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	if v188 == int32(0) {
		goto L46
	} else {
		goto L53
	}
L53:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[6])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+176)) = v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v12)+3280))
	v196 = base.I32_div_s(v194, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_7))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+180)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v12)+184)) = v143
	F_errmsg(m, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_8), v12+int32(176))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_2), int32(146), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_3))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	goto L46
L56:
	;
	goto L41
L57:
	;
	v218 = m.G0
	v220 = v218 - int32(16)
	m.G0 = v220
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[4]))
	if v223 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if base.B2i32(v223 != int32(0)) == int32(0) {
		goto L42
	} else {
		goto L62
	}
L59:
	;
	v224 = F_GetCurrentTimestamp(m)
	mBase = m.M
	v226 = *(*int64)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[5]))
	F_TimestampDifference(m, v226, v224, v220+int32(12), v220+int32(8))
	mBase = m.M
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[6]))) = v232
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(3280)))) = v234
	*(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[4])) = int32(0)
	goto L61
L60:
	;
	goto L61
L61:
	;
	m.G0 = v220 + int32(16)
	goto L58
L62:
	;
	v249 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	if v249 == int32(0) {
		goto L42
	} else {
		goto L64
	}
L64:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[6])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v253
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v12)+3280))
	v257 = base.I32_div_s(v255, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_7))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v12)+168)) = v12 + int32(208)
	F_errmsg(m, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_9), v12+int32(160))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_2), int32(149), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_3))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	goto L42
L67:
	;
	v285 = v12 + int32(208)
	v286 = F_AllocateDir(m, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v288 = F_ReadDir(m, v286, v285)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	if v288 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v293 = v288
	goto L73
L71:
	;
	goto L72
L72:
	;
	F_FreeDir(m, v286)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L4
	} else {
		goto L98
	}
L73:
	;
	v300 = v293 + int32(19)
	v304 = v12 + int32(2256)
	v307 = int32(0)
	v312 = m.G0
	v314 = v312 - int32(16)
	m.G0 = v314
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[7]))) = v307
	*(*int32)(unsafe.Add(mBase, uint32(v304))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[8]))) = v307
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	if base.Ui32((v322-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v404 = v307
		goto L77
	} else {
		goto L78
	}
L74:
	;
	goto L72
L75:
	;
	v419 = F_ReadDir(m, v286, v12+int32(208))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L4
	} else {
		goto L96
	}
L76:
	;
	if v404 == int32(0) {
		goto L75
	} else {
		goto L93
	}
L77:
	;
	m.G0 = v314 + int32(16)
	goto L76
L78:
	;
	v329 = int32(_a_F_ResetUnloggedRelationsInTablespaceDir_10)
	*(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0])) = int32(0)
	v335 = F_strtoul(m, v300, v314+int32(8), int32(10))
	mBase = m.M
	v337 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0]))
	if v337 != 0 {
		v404 = v307
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v314)+8))
	if base.B2i32(v335 == int32(0))|base.B2i32(v300 == v340) != 0 {
		v404 = v307
		goto L77
	} else {
		goto L80
	}
L80:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340))))
	if v343 != int32(95) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	if v359&int32(255) == int32(46) {
		goto L86
	} else {
		goto L87
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314)+12)) = int32(0)
	v359 = v343
	v360 = v340
	goto L81
L83:
	;
	goto L84
L84:
	;
	v352 = F_forkname_chars(m, v340+int32(1), v314+int32(12))
	mBase = m.M
	if v352 <= int32(0) {
		v404 = v307
		goto L77
	} else {
		goto L85
	}
L85:
	;
	v357 = v352 + v340 + int32(1)
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357))))
	v359 = v358
	v360 = v357
	goto L81
L86:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+1)))
	if base.Ui32((v365-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v404 = v307
		goto L77
	} else {
		goto L89
	}
L87:
	;
	v391 = v307
	v392 = v359
	goto L88
L88:
	;
	if v392&int32(255) != 0 {
		v404 = v307
		goto L77
	} else {
		goto L92
	}
L89:
	;
	v372 = int32(_a_F_ResetUnloggedRelationsInTablespaceDir_10)
	*(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0])) = int32(0)
	v376 = v360 + int32(1)
	v380 = F_strtoul(m, v376, v314+int32(8), int32(10))
	mBase = m.M
	v382 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0]))
	if v382 != 0 {
		v404 = v307
		goto L77
	} else {
		goto L90
	}
L90:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v314)+8))
	if base.B2i32(v380 == int32(0))|base.B2i32(v376 == v385) != 0 {
		v404 = v307
		goto L77
	} else {
		goto L91
	}
L91:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	v391 = v380
	v392 = v388
	goto L88
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[7]))) = v335
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v314)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v304))) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[8]))) = v391
	v404 = int32(1)
	goto L77
L93:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2256))
	if v410 != int32(3) {
		goto L75
	} else {
		goto L94
	}
L94:
	;
	v415 = F_hash_search(m, v282, v12+int32(_a_F_ResetUnloggedRelationsInTablespaceDir_11), int32(1), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	goto L75
L96:
	;
	if v419 != 0 {
		v293 = v419
		goto L73
	} else {
		goto L97
	}
L97:
	;
	goto L74
L98:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v433)+4))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v433)+412))
	if v435 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if v498 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v433)+376))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v433)+364))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v433)+352))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v433)+340))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v433)+328))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v433)+316))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v433)+304))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v433)+292))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v433)+280))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v433)+268))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v433)+256))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v433)+244))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v433)+232))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v433)+220))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v433)+208))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v433)+196))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v433)+184))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v433)+172))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v433)+160))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v433)+148))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v433)+136))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v433)+124))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v433)+112))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v433)+100))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v433)+88))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v433)+76))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v433)+64))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v433)+52))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v433)+40))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v433)+28))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v433)+16))
	v498 = v436 + (v437 + (v438 + (v439 + (v440 + (v441 + (v442 + (v443 + (v444 + (v445 + (v446 + (v447 + (v448 + (v449 + (v450 + (v451 + (v452 + (v453 + (v454 + (v455 + (v456 + (v457 + (v458 + (v459 + (v460 + (v461 + (v462 + (v463 + (v464 + (v465 + (v466 + v434))))))))))))))))))))))))))))))
	goto L102
L101:
	;
	v498 = v434
	goto L102
L102:
	;
	goto L99
L103:
	;
	F_hash_destroy(m, v282)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L4
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v504 = v12 + int32(208)
	v505 = F_AllocateDir(m, v504)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L4
	} else {
		goto L107
	}
L106:
	;
	goto L19
L107:
	;
	v507 = F_ReadDir(m, v505, v504)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	if v507 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v512 = v507
	goto L112
L110:
	;
	goto L111
L111:
	;
	F_FreeDir(m, v505)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L4
	} else {
		goto L144
	}
L112:
	;
	v519 = v512 + int32(19)
	v523 = v12 + int32(2256)
	v526 = int32(0)
	v531 = m.G0
	v533 = v531 - int32(16)
	m.G0 = v533
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[7]))) = v526
	*(*int32)(unsafe.Add(mBase, uint32(v523))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[8]))) = v526
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519))))
	if base.Ui32((v541-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v623 = v526
		goto L116
	} else {
		goto L117
	}
L113:
	;
	goto L111
L114:
	;
	v673 = F_ReadDir(m, v505, v12+int32(208))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L4
	} else {
		goto L142
	}
L115:
	;
	if v623 == int32(0) {
		goto L114
	} else {
		goto L132
	}
L116:
	;
	m.G0 = v533 + int32(16)
	goto L115
L117:
	;
	v548 = int32(_a_F_ResetUnloggedRelationsInTablespaceDir_10)
	*(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0])) = int32(0)
	v554 = F_strtoul(m, v519, v533+int32(8), int32(10))
	mBase = m.M
	v556 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0]))
	if v556 != 0 {
		v623 = v526
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v533)+8))
	if base.B2i32(v554 == int32(0))|base.B2i32(v519 == v559) != 0 {
		v623 = v526
		goto L116
	} else {
		goto L119
	}
L119:
	;
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559))))
	if v562 != int32(95) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	if v578&int32(255) == int32(46) {
		goto L125
	} else {
		goto L126
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v533)+12)) = int32(0)
	v578 = v562
	v579 = v559
	goto L120
L122:
	;
	goto L123
L123:
	;
	v571 = F_forkname_chars(m, v559+int32(1), v533+int32(12))
	mBase = m.M
	if v571 <= int32(0) {
		v623 = v526
		goto L116
	} else {
		goto L124
	}
L124:
	;
	v576 = v571 + v559 + int32(1)
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576))))
	v578 = v577
	v579 = v576
	goto L120
L125:
	;
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+1)))
	if base.Ui32((v584-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v623 = v526
		goto L116
	} else {
		goto L128
	}
L126:
	;
	v610 = v526
	v611 = v578
	goto L127
L127:
	;
	if v611&int32(255) != 0 {
		v623 = v526
		goto L116
	} else {
		goto L131
	}
L128:
	;
	v591 = int32(_a_F_ResetUnloggedRelationsInTablespaceDir_10)
	*(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0])) = int32(0)
	v595 = v579 + int32(1)
	v599 = F_strtoul(m, v595, v533+int32(8), int32(10))
	mBase = m.M
	v601 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0]))
	if v601 != 0 {
		v623 = v526
		goto L116
	} else {
		goto L129
	}
L129:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v533)+8))
	if base.B2i32(v599 == int32(0))|base.B2i32(v595 == v604) != 0 {
		v623 = v526
		goto L116
	} else {
		goto L130
	}
L130:
	;
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604))))
	v610 = v599
	v611 = v607
	goto L127
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[7]))) = v554
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v533)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v523))) = v615
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[8]))) = v610
	v623 = int32(1)
	goto L116
L132:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2256))
	if v629 == int32(3) {
		goto L114
	} else {
		goto L133
	}
L133:
	;
	v632 = int32(0)
	v634 = F_hash_search(m, v282, v12+int32(_a_F_ResetUnloggedRelationsInTablespaceDir_11), v632, v632)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	if v634 == int32(0) {
		goto L114
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+148)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v12 + int32(208)
	v643 = v12 + int32(_a_F_ResetUnloggedRelationsInTablespaceDir_12)
	v648 = F_pg_snprintf(m, v643, int32(2048), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_5), v12+int32(144))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	v650 = F_unlink(m, v643)
	mBase = m.M
	if v650 < int32(0) {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v655 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	if v655 == int32(0) {
		goto L114
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v643
	F_errmsg_internal(m, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_13), v12+int32(128))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_2), int32(264), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_14))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	goto L114
L142:
	;
	if v673 != 0 {
		v512 = v673
		goto L112
	} else {
		goto L143
	}
L143:
	;
	goto L113
L144:
	;
	F_hash_destroy(m, v282)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	if v42 == int32(0) {
		goto L19
	} else {
		goto L146
	}
L146:
	;
	goto L41
L147:
	;
	v703 = F_ReadDir(m, v701, v700)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	if v703 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v708 = v703
	goto L152
L150:
	;
	goto L151
L151:
	;
	F_FreeDir(m, v701)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L4
	} else {
		goto L190
	}
L152:
	;
	v715 = v708 + int32(19)
	v722 = int32(0)
	v727 = m.G0
	v729 = v727 - int32(16)
	m.G0 = v729
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[7]))) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[8]))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[9]))) = v722
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v715))))
	if base.Ui32((v737-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v819 = v722
		goto L156
	} else {
		goto L157
	}
L153:
	;
	goto L151
L154:
	;
	v898 = F_ReadDir(m, v701, v12+int32(208))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L4
	} else {
		goto L188
	}
L155:
	;
	if v819 == int32(0) {
		goto L154
	} else {
		goto L172
	}
L156:
	;
	m.G0 = v729 + int32(16)
	goto L155
L157:
	;
	v744 = int32(_a_F_ResetUnloggedRelationsInTablespaceDir_10)
	*(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0])) = int32(0)
	v750 = F_strtoul(m, v715, v729+int32(8), int32(10))
	mBase = m.M
	v752 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0]))
	if v752 != 0 {
		v819 = v722
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v729)+8))
	if base.B2i32(v750 == int32(0))|base.B2i32(v715 == v755) != 0 {
		v819 = v722
		goto L156
	} else {
		goto L159
	}
L159:
	;
	v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755))))
	if v758 != int32(95) {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	if v774&int32(255) == int32(46) {
		goto L165
	} else {
		goto L166
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v729)+12)) = int32(0)
	v774 = v758
	v775 = v755
	goto L160
L162:
	;
	goto L163
L163:
	;
	v767 = F_forkname_chars(m, v755+int32(1), v729+int32(12))
	mBase = m.M
	if v767 <= int32(0) {
		v819 = v722
		goto L156
	} else {
		goto L164
	}
L164:
	;
	v772 = v767 + v755 + int32(1)
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v772))))
	v774 = v773
	v775 = v772
	goto L160
L165:
	;
	v780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775)+1)))
	if base.Ui32((v780-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v819 = v722
		goto L156
	} else {
		goto L168
	}
L166:
	;
	v806 = v722
	v807 = v774
	goto L167
L167:
	;
	if v807&int32(255) != 0 {
		v819 = v722
		goto L156
	} else {
		goto L171
	}
L168:
	;
	v787 = int32(_a_F_ResetUnloggedRelationsInTablespaceDir_10)
	*(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0])) = int32(0)
	v791 = v775 + int32(1)
	v795 = F_strtoul(m, v791, v729+int32(8), int32(10))
	mBase = m.M
	v797 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0]))
	if v797 != 0 {
		v819 = v722
		goto L156
	} else {
		goto L169
	}
L169:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v729)+8))
	if base.B2i32(v795 == int32(0))|base.B2i32(v791 == v800) != 0 {
		v819 = v722
		goto L156
	} else {
		goto L170
	}
L170:
	;
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v800))))
	v806 = v795
	v807 = v803
	goto L167
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[7]))) = v750
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v729)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[8]))) = v811
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[9]))) = v806
	v819 = int32(1)
	goto L156
L172:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[8])))
	if v825 != int32(3) {
		goto L154
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+100)) = v715
	v830 = v12 + int32(208)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v830
	v838 = F_pg_snprintf(m, v12+int32(3280), int32(2048), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_5), v12+int32(96))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L4
	} else {
		goto L174
	}
L174:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[7])))
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[9])))
	if v841 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v869 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L4
	} else {
		goto L181
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v830
	v852 = F_pg_snprintf(m, v12+int32(2256), int32(1024), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_15), v12-int32(-64))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L4
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+88)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v12 + int32(208)
	v865 = F_pg_snprintf(m, v12+int32(2256), int32(1024), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_16), v12+int32(80))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L4
	} else {
		goto L180
	}
L179:
	;
	goto L175
L180:
	;
	goto L175
L181:
	;
	if v869 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v12 + int32(2256)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v12 + int32(3280)
	F_errmsg_internal(m, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_17), v12+int32(48))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L4
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	F_copy_file(m, v12+int32(3280), v12+int32(2256))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L4
	} else {
		goto L187
	}
L185:
	;
	F_errfinish(m, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_2), int32(314), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_14))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L4
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	goto L154
L188:
	;
	if v898 != 0 {
		v708 = v898
		goto L152
	} else {
		goto L189
	}
L189:
	;
	goto L153
L190:
	;
	v912 = v12 + int32(208)
	v913 = F_AllocateDir(m, v912)
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L4
	} else {
		goto L191
	}
L191:
	;
	v915 = F_ReadDir(m, v913, v912)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	if v915 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v920 = v915
	goto L196
L194:
	;
	goto L195
L195:
	;
	F_FreeDir(m, v913)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L4
	} else {
		goto L227
	}
L196:
	;
	v927 = v920 + int32(19)
	v929 = v12 + int32(2256)
	v934 = int32(0)
	v939 = m.G0
	v941 = v939 - int32(16)
	m.G0 = v941
	*(*int32)(unsafe.Add(mBase, uint32(v929))) = v934
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[8]))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[7]))) = v934
	v949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v927))))
	if base.Ui32((v949-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v1031 = v934
		goto L200
	} else {
		goto L201
	}
L197:
	;
	goto L195
L198:
	;
	v1078 = F_ReadDir(m, v913, v12+int32(208))
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L4
	} else {
		goto L225
	}
L199:
	;
	if v1031 == int32(0) {
		goto L198
	} else {
		goto L216
	}
L200:
	;
	m.G0 = v941 + int32(16)
	goto L199
L201:
	;
	v956 = int32(_a_F_ResetUnloggedRelationsInTablespaceDir_10)
	*(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0])) = int32(0)
	v962 = F_strtoul(m, v927, v941+int32(8), int32(10))
	mBase = m.M
	v964 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0]))
	if v964 != 0 {
		v1031 = v934
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v941)+8))
	if base.B2i32(v962 == int32(0))|base.B2i32(v927 == v967) != 0 {
		v1031 = v934
		goto L200
	} else {
		goto L203
	}
L203:
	;
	v970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v967))))
	if v970 != int32(95) {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	if v986&int32(255) == int32(46) {
		goto L209
	} else {
		goto L210
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v941)+12)) = int32(0)
	v986 = v970
	v987 = v967
	goto L204
L206:
	;
	goto L207
L207:
	;
	v979 = F_forkname_chars(m, v967+int32(1), v941+int32(12))
	mBase = m.M
	if v979 <= int32(0) {
		v1031 = v934
		goto L200
	} else {
		goto L208
	}
L208:
	;
	v984 = v979 + v967 + int32(1)
	v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v984))))
	v986 = v985
	v987 = v984
	goto L204
L209:
	;
	v992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v987)+1)))
	if base.Ui32((v992-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v1031 = v934
		goto L200
	} else {
		goto L212
	}
L210:
	;
	v1018 = v934
	v1019 = v986
	goto L211
L211:
	;
	if v1019&int32(255) != 0 {
		v1031 = v934
		goto L200
	} else {
		goto L215
	}
L212:
	;
	v999 = int32(_a_F_ResetUnloggedRelationsInTablespaceDir_10)
	*(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0])) = int32(0)
	v1003 = v987 + int32(1)
	v1007 = F_strtoul(m, v1003, v941+int32(8), int32(10))
	mBase = m.M
	v1009 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0]))
	if v1009 != 0 {
		v1031 = v934
		goto L200
	} else {
		goto L213
	}
L213:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v941)+8))
	if base.B2i32(v1007 == int32(0))|base.B2i32(v1003 == v1012) != 0 {
		v1031 = v934
		goto L200
	} else {
		goto L214
	}
L214:
	;
	v1015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1012))))
	v1018 = v1007
	v1019 = v1015
	goto L211
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v929))) = v962
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v941)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[8]))) = v1023
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[7]))) = v1018
	v1031 = int32(1)
	goto L200
L216:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[8])))
	if v1037 != int32(3) {
		goto L198
	} else {
		goto L217
	}
L217:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2256))
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[7])))
	if v1041 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	F_fsync_fname(m, v12+int32(3280), int32(0))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L4
	} else {
		goto L224
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v1040
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(208)
	v1054 = F_pg_snprintf(m, v12+int32(3280), int32(1024), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_15), v12+int32(16))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L4
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v1041
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v1040
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v12 + int32(208)
	v1067 = F_pg_snprintf(m, v12+int32(3280), int32(1024), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_16), v12+int32(32))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L4
	} else {
		goto L223
	}
L222:
	;
	goto L218
L223:
	;
	goto L218
L224:
	;
	goto L198
L225:
	;
	if v1078 != 0 {
		v920 = v1078
		goto L196
	} else {
		goto L226
	}
L226:
	;
	goto L197
L227:
	;
	F_fsync_fname(m, v12+int32(208), int32(1))
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L4
	} else {
		goto L228
	}
L228:
	;
	goto L19
L229:
	;
	if v1105 != 0 {
		v46 = v1105
		goto L17
	} else {
		goto L230
	}
L230:
	;
	goto L18
L231:
	;
	goto L2
L232:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L4
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v12 + int32(_a_F_ResetUnloggedRelationsInTablespaceDir_12)
	F_errmsg(m, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_18), v12+int32(112))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L4
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_2), int32(262), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_14))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L4
	} else {
		goto L235
	}
L235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RestrictSearchPath(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_RestrictSearchPath[0]))
	if v2 != 0 {
		v4 = int32(0)
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_RestrictSearchPath[1]))
		v14 = F_set_config_with_handle(m, int32(_a_F_RestrictSearchPath_0), v4, int32(_a_F_RestrictSearchPath_1), int32(6), int32(13), v9, int32(2), int32(1), v4, v4)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_r_consonant_pair_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v8 < v10 {
		v71 = v2
		return v71
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v8
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v10
		v17 = v8 - int32(1)
		if v10 < v17 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+v17))))
			v23 = v21 - int32(100)
			if base.B2i32(v23 == int32(0))|base.B2i32(v23 == int32(16)) != 0 {
				v35 = F_find_among_b(m, l0, int32(_a_F_r_consonant_pair_1_0), int32(4))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					if v35 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v14
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v14
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v45
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v49 = v47 + (v8 - v12)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49
						if v49 <= v14 {
							v71 = v2
							return v71
						} else {
							v52 = int32(1)
							v53 = v49 - v52
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v53
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v53
							v57 = F_slice_del(m, l0)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								if int32(0) <= v57 {
									v64 = v52
								} else {
									v64 = v57 >> (uint(int32(31)) % 32) & v57
								}
								v71 = v64
								return v71
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v14
				return int32(0)
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v14
			return int32(0)
		}
	}
}
func F_r_en_ending_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	if v5 < v7 {
		v172 = v2
		return v172
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v22 <= v23 {
			v132 = int32(-1)
			v139 = v132
		} else {
			v40 = int32(1)
			v41 = v22 - v40
			v43 = int32(*(*int8)(unsafe.Add(mBase, uint32(v24+v41))))
			v45 = v43 & int32(255)
			if base.B2i32(v41 == v23)|base.B2i32(int32(0) <= v43) != 0 {
				v103 = v45
				v107 = v40
			} else {
				v52 = v45 & int32(63)
				v54 = v22 - int32(2)
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v54))))
				v58 = v56 << (uint(int32(6)) % 32)
				if base.B2i32(v54 != v23)&base.B2i32(base.Ui32(v56) < base.Ui32(int32(192))) == int32(0) {
					v103 = v58&int32(1984) | v52
					v107 = int32(2)
				} else {
					v71 = v58&int32(4032) | v52
					v73 = v22 - int32(3)
					v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v73))))
					if base.B2i32(v73 != v23)&base.B2i32(base.Ui32(v75) < base.Ui32(int32(224))) == int32(0) {
						v103 = v75<<(uint(int32(12))%32)&int32(_a_F_r_en_ending_2_0) | v71
						v107 = int32(3)
					} else {
						v93 = int32(4)
						v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v24-v93))))
						v103 = v75<<(uint(int32(12))%32)&int32(_a_F_r_en_ending_2_1) | v95&int32(7)<<(uint(int32(18))%32) | v71
						v107 = v93
					}
				}
			}
			if int32(232) < v103 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v22 - v107
				v132 = int32(0)
				v139 = v132
			} else {
				v109 = v103 - int32(97)
				if v109 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v22 - v107
					v132 = int32(0)
					v139 = v132
				} else {
					v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v109)>>(uint(int32(3))%32)))+uint32(_c_F_r_en_ending_2[0]))))
					if int32(base.Ui32(v115)>>(uint(v109&int32(7))%32))&int32(1) == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v22 - v107
						v132 = int32(0)
						v139 = v132
					} else {
						v139 = v107
					}
				}
			}
		}
		if v139 != 0 {
			v172 = v2
			return v172
		} else {
			v140 = v5 - v9
			v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v142 = v140 + v141
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v142
			v144 = int32(3)
			v146 = int32(0)
			v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v142-v149 < v144 {
				v159 = v146
			} else {
				v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v155 = F_memcmp(m, v152+v142-v144, int32(_a_F_r_en_ending_2_2), v144)
				mBase = m.M
				if v155 != 0 {
					v159 = v146
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v142 - v144
					v159 = int32(1)
				}
			}
			if v159 != 0 {
				v172 = v2
				return v172
			} else {
				v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v160 + v140
				v163 = F_slice_del(m, l0)
				mBase = m.M
				v166 = m.ExcPending
				if v166 != 0 {
					return int32(0)
				} else {
					if v163 < int32(0) {
						v172 = v163
						return v172
					} else {
						v169 = F_r_undouble_3(m, l0)
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
							return int32(0)
						} else {
							v172 = v169
							return v172
						}
					}
				}
			}
		}
	}
}
func F_r_remove_second_order_prefix_2(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13983(m, l0, int32(_a_F_r_remove_second_order_prefix_2_0), int32(_a_F_r_remove_second_order_prefix_2_1), int32(_a_F_r_remove_second_order_prefix_2_2))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_r_shortv_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 <= v15 {
		v55 = int32(-1)
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v6-int32(1)))))
		if int32(121) < v30 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 - int32(1)
			v52 = int32(0)
		} else {
			v32 = v30 - int32(89)
			if v32 < int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 - int32(1)
				v52 = int32(0)
			} else {
				v35 = int32(1)
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v32)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_1[0]))))
				if int32(base.Ui32(v39)>>(uint(v32&int32(7))%32))&v35 != 0 {
					v52 = v35
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 - int32(1)
					v52 = int32(0)
				}
			}
		}
		v55 = v52
	}
	if v55 != 0 {
		v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v162 = v160 + (v6 - v5)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v162
		v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v162 <= v172 {
			v212 = int32(-1)
		} else {
			v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+v162-int32(1)))))
			if int32(121) < v187 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v162 - int32(1)
				v209 = int32(0)
			} else {
				v189 = v187 - int32(97)
				if v189 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v162 - int32(1)
					v209 = int32(0)
				} else {
					v192 = int32(1)
					v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v189)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_1[1]))))
					if int32(base.Ui32(v196)>>(uint(v189&int32(7))%32))&v192 != 0 {
						v209 = v192
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v162 - int32(1)
						v209 = int32(0)
					}
				}
			}
			v212 = v209
		}
		if v212 != 0 {
			v269 = v2
		} else {
			v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v221 <= v222 {
				v265 = int32(-1)
			} else {
				v234 = int32(1)
				v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235+v221-v234))))
				if int32(121) < v239 {
					v261 = v234
				} else {
					v241 = v239 - int32(97)
					if v241 < int32(0) {
						v261 = v234
					} else {
						v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v241)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_1[1]))))
						if int32(base.Ui32(v247)>>(uint(v241&int32(7))%32))&int32(1) == int32(0) {
							v261 = v234
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v221 - int32(1)
							v261 = int32(0)
						}
					}
				}
				v265 = v261
			}
			if v265 != 0 {
				v269 = v2
			} else {
				v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v269 = base.B2i32(v266 <= v267)
			}
		}
		return v269
	} else {
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v64 <= v65 {
			v108 = int32(-1)
		} else {
			v77 = int32(1)
			v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+v64-v77))))
			if int32(121) < v82 {
				v104 = v77
			} else {
				v84 = v82 - int32(97)
				if v84 < int32(0) {
					v104 = v77
				} else {
					v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v84)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_1[1]))))
					if int32(base.Ui32(v90)>>(uint(v84&int32(7))%32))&int32(1) == int32(0) {
						v104 = v77
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v64 - int32(1)
						v104 = int32(0)
					}
				}
			}
			v108 = v104
		}
		if v108 != 0 {
			v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v162 = v160 + (v6 - v5)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v162
			v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v162 <= v172 {
				v212 = int32(-1)
			} else {
				v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+v162-int32(1)))))
				if int32(121) < v187 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v162 - int32(1)
					v209 = int32(0)
				} else {
					v189 = v187 - int32(97)
					if v189 < int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v162 - int32(1)
						v209 = int32(0)
					} else {
						v192 = int32(1)
						v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v189)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_1[1]))))
						if int32(base.Ui32(v196)>>(uint(v189&int32(7))%32))&v192 != 0 {
							v209 = v192
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v162 - int32(1)
							v209 = int32(0)
						}
					}
				}
				v212 = v209
			}
			if v212 != 0 {
				v269 = v2
			} else {
				v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v221 <= v222 {
					v265 = int32(-1)
				} else {
					v234 = int32(1)
					v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235+v221-v234))))
					if int32(121) < v239 {
						v261 = v234
					} else {
						v241 = v239 - int32(97)
						if v241 < int32(0) {
							v261 = v234
						} else {
							v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v241)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_1[1]))))
							if int32(base.Ui32(v247)>>(uint(v241&int32(7))%32))&int32(1) == int32(0) {
								v261 = v234
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v221 - int32(1)
								v261 = int32(0)
							}
						}
					}
					v265 = v261
				}
				if v265 != 0 {
					v269 = v2
				} else {
					v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v269 = base.B2i32(v266 <= v267)
				}
			}
			return v269
		} else {
			v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v116 <= v117 {
				v157 = int32(-1)
			} else {
				v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128+v116-int32(1)))))
				if int32(121) < v132 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v116 - int32(1)
					v154 = int32(0)
				} else {
					v134 = v132 - int32(97)
					if v134 < int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v116 - int32(1)
						v154 = int32(0)
					} else {
						v137 = int32(1)
						v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v134)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_1[1]))))
						if int32(base.Ui32(v141)>>(uint(v134&int32(7))%32))&v137 != 0 {
							v154 = v137
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v116 - int32(1)
							v154 = int32(0)
						}
					}
				}
				v157 = v154
			}
			if v157 != 0 {
				v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v162 = v160 + (v6 - v5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v162
				v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v162 <= v172 {
					v212 = int32(-1)
				} else {
					v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+v162-int32(1)))))
					if int32(121) < v187 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v162 - int32(1)
						v209 = int32(0)
					} else {
						v189 = v187 - int32(97)
						if v189 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v162 - int32(1)
							v209 = int32(0)
						} else {
							v192 = int32(1)
							v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v189)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_1[1]))))
							if int32(base.Ui32(v196)>>(uint(v189&int32(7))%32))&v192 != 0 {
								v209 = v192
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v162 - int32(1)
								v209 = int32(0)
							}
						}
					}
					v212 = v209
				}
				if v212 != 0 {
					v269 = v2
				} else {
					v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v221 <= v222 {
						v265 = int32(-1)
					} else {
						v234 = int32(1)
						v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235+v221-v234))))
						if int32(121) < v239 {
							v261 = v234
						} else {
							v241 = v239 - int32(97)
							if v241 < int32(0) {
								v261 = v234
							} else {
								v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v241)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_1[1]))))
								if int32(base.Ui32(v247)>>(uint(v241&int32(7))%32))&int32(1) == int32(0) {
									v261 = v234
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v221 - int32(1)
									v261 = int32(0)
								}
							}
						}
						v265 = v261
					}
					if v265 != 0 {
						v269 = v2
					} else {
						v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v269 = base.B2i32(v266 <= v267)
					}
				}
				return v269
			} else {
				return int32(1)
			}
		}
	}
}
func F_r_undouble_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = v5 - int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 <= v8 {
		v56 = v2
		return v56
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v7))))
		if base.B2i32(v12&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v12)%32)&int32(_a_F_r_undouble_1_0) == int32(0)) != 0 {
			v56 = v2
			return v56
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v27 = F_find_among_b(m, l0, int32(_a_F_r_undouble_1_1), int32(3))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				if v27 == int32(0) {
					v56 = v2
					return v56
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v35 = v33 + (v5 - v24)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v35
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v35
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v35 <= v38 {
						v56 = v2
						return v56
					} else {
						v40 = int32(1)
						v41 = v35 - v40
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v41
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v41
						v45 = F_slice_del(m, l0)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							if int32(0) <= v45 {
								v52 = v40
							} else {
								v52 = v45 >> (uint(int32(31)) % 32) & v45
							}
							v56 = v52
							return v56
						}
					}
				}
			}
		}
	}
}
func F_read_cursor_args(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v493 int32
	_ = v493
	v6 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(128)
	m.G0 = v16
	v18 = F_plpgsql_yylex(m, l2, l3, l4)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v22 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	F_plpgsql_yyerror(m, l3, int32(0), l4, int32(_a_F_read_cursor_args_0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L107
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_read_cursor_args_1))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L102
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_read_cursor_args_1))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L97
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_read_cursor_args_1))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L92
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_read_cursor_args_1))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L87
	}
L8:
	;
	m.G0 = v16 + int32(128)
	return v374
L9:
	;
	if v18 == int32(40) {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if v18 != int32(40) {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	if l1 == v18 {
		v374 = int32(0)
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L3
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_read_cursor_args[0]))
	v33 = int32(2)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v22<<(uint(v33)%32))))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
	v40 = F_palloc0(m, v37<<(uint(v33)%32))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
	if int32(0) < v42 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v53 = v6
	v55 = v6
	goto L19
L17:
	;
	v273 = v6
	goto L18
L18:
	;
	F_initStringInfo(m, v16+int32(112))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L65
	}
L19:
	;
	F_plpgsql_peek2(m, v16+int32(108), v16+int32(104), v16+int32(100), int32(0), l4)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v273 = v207
	goto L18
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v16)+108))
	if v67 != int32(258) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v212 = v40 + v203<<(uint(int32(2))%32)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if v213 != 0 {
		goto L7
	} else {
		goto L52
	}
L23:
	;
	v203 = v53
	v207 = v55
	goto L22
L24:
	;
	goto L25
L25:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v16)+104))
	if v70&int32(-2) != int32(270) {
		v203 = v53
		v207 = v55
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v75 = int32(_a_F_read_cursor_args_2)
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_read_cursor_args[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_read_cursor_args[1])) = int32(1)
	v80 = F_plpgsql_yylex(m, l2, l3, l4)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, _c_F_read_cursor_args[1])) = v76
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
	if v85 <= int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_read_cursor_args_1))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L47
	}
L29:
	;
	if v145 == v85 {
		goto L28
	} else {
		goto L44
	}
L30:
	;
	v145 = int32(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v97 = int32(0)
	goto L33
L33:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v89+v97<<(uint(int32(2))%32))))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if base.B2i32(v110 == int32(0))|base.B2i32(v110 != v113) != 0 {
		v131 = v110
		v132 = v113
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L28
L35:
	;
	if v131-v132 == int32(0) {
		v145 = v97
		goto L29
	} else {
		goto L42
	}
L36:
	;
	goto L35
L37:
	;
	v116 = v107
	v117 = v82
	goto L38
L38:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	if v121 == int32(0) {
		v131 = v121
		v132 = v120
		goto L36
	} else {
		goto L40
	}
L39:
	;
	v131 = v121
	v132 = v120
	goto L36
L40:
	;
	v124 = int32(1)
	if v121 == v120 {
		v116 = v116 + v124
		v117 = v117 + v124
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v137 = v97 + int32(1)
	if v137 != v85 {
		v97 = v137
		goto L33
	} else {
		goto L43
	}
L43:
	;
	goto L34
L44:
	;
	v153 = F_plpgsql_yylex(m, l2, l3, l4)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+104)) = v153
	if base.Ui32(int32(-3)) < base.Ui32(v153-int32(272)) {
		v203 = v145
		v207 = int32(1)
		goto L22
	} else {
		goto L46
	}
L46:
	;
	goto L3
L47:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v181
	F_errmsg(m, int32(_a_F_read_cursor_args_3), v16+int32(16))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v190 = F_plpgsql_scanner_errposition(m, v189, l4)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_read_cursor_args_4), int32(3986), int32(_a_F_read_cursor_args_5))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	v216 = int32(0)
	v219 = int32(1)
	v224 = F_read_sql_construct(m, int32(44), int32(41), v216, int32(_a_F_read_cursor_args_6), int32(2), v219, v219, v216, v16+int32(112), l2, l3, l4)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	*(*int32)(unsafe.Add(mBase, uint32(v212))) = v226
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v16)+112))
	switch v229 - int32(41) {
	case 0:
		goto L56
	default:
		goto L54
	case 3:
		goto L55
	}
L54:
	;
	v261 = v53 + int32(1)
	if v261 < v228 {
		v53 = v261
		v55 = v207
		goto L19
	} else {
		goto L64
	}
L55:
	;
	if v53 == v228-int32(1) {
		goto L6
	} else {
		goto L63
	}
L56:
	;
	if v53 == v228-int32(1) {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_read_cursor_args_1))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v242
	F_errmsg(m, int32(_a_F_read_cursor_args_7), v16+int32(48))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v250 = F_plpgsql_scanner_errposition(m, v249, l4)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_read_cursor_args_4), int32(4028), int32(_a_F_read_cursor_args_5))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	goto L54
L64:
	;
	goto L20
L65:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
	if int32(0) < v280 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v290 = int32(0)
	goto L69
L67:
	;
	goto L68
L68:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v16)+112))
	v345 = F_palloc0(m, int32(80))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L82
	}
L69:
	;
	v298 = v16 + int32(112)
	v300 = v290 << (uint(int32(2)) % 32)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v40+v300)))
	F_appendStringInfoString(m, v298, v302)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L71
	}
L70:
	;
	goto L68
L71:
	;
	if v273 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v305+v300)))
	v308 = F_quote_identifier(m, v307)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
	if v290 < v316-int32(1) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v308
	F_appendStringInfo(m, v298, int32(_a_F_read_cursor_args_8), v16+int32(32))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	F_appendStringInfoString(m, v16+int32(112), int32(_a_F_read_cursor_args_9))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	v326 = v316
	goto L79
L79:
	;
	v328 = v290 + int32(1)
	if v328 < v326 {
		v290 = v328
		goto L69
	} else {
		goto L81
	}
L80:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
	v326 = v325
	goto L79
L81:
	;
	goto L70
L82:
	;
	v347 = F_pstrdup(m, v343)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v345)+4)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v345))) = v347
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_read_cursor_args[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v345)+8)) = v353
	v356 = *(*int32)(unsafe.Add(mBase, _c_F_read_cursor_args[3]))
	v357 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v345)+20)) = uint8(v357)
	*(*int32)(unsafe.Add(mBase, uint32(v345)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v345)+12)) = v356
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v16)+112))
	F_pfree(m, v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v365 = F_plpgsql_yylex(m, l2, l3, l4)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v365 != l1 {
		goto L3
	} else {
		goto L86
	}
L86:
	;
	v374 = v345
	goto L8
L87:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v392+v203<<(uint(int32(2))%32))))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v396
	F_errmsg(m, int32(_a_F_read_cursor_args_10), v16+int32(80))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
	v406 = F_plpgsql_scanner_errposition(m, v405, l4)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_read_cursor_args_4), int32(4006), int32(_a_F_read_cursor_args_5))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v420
	F_errmsg(m, int32(_a_F_read_cursor_args_11), v16-int32(-64))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v428 = F_plpgsql_scanner_errposition(m, v427, l4)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_read_cursor_args_4), int32(4035), int32(_a_F_read_cursor_args_5))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v442
	F_errmsg(m, int32(_a_F_read_cursor_args_12), v16+int32(96))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v450 = F_plpgsql_scanner_errposition(m, v449, l4)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_read_cursor_args_4), int32(3941), int32(_a_F_read_cursor_args_5))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v464
	F_errmsg(m, int32(_a_F_read_cursor_args_13), v16)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v470 = F_plpgsql_scanner_errposition(m, v469, l4)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_read_cursor_args_4), int32(3927), int32(_a_F_read_cursor_args_5))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
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
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_read_into_target(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
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
	var v83 int32
	_ = v83
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	if l1 != 0 {
		v9 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v9)
		v11 = F_plpgsql_yylex(m, l2, l3, l4)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			if v11 != int32(373) {
				v20 = v11
				if v20 != int32(277) {
					switch v20 - int32(275) {
					case 0:
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						F_word_is_not_variable(m, l2, v55, l4)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					case 1:
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						F_cword_is_not_variable(m, l2, v58, l4)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					default:
						F_plpgsql_yyerror(m, l3, int32(0), l4, int32(_a_F_read_into_target_0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					v27 = int32(1)
					if base.Ui32(v26-v27) <= base.Ui32(v27) {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						F_check_assignable(m, v25, v31, l4)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v34
							v36 = F_plpgsql_yylex(m, l2, l3, l4)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								if v36 == int32(44) {
									F_errstart_cold(m, int32(21), int32(_a_F_read_into_target_1))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										F_errcode(m, int32(16801924))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_read_into_target_2), int32(0))
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return
											} else {
												v76 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
												v77 = F_plpgsql_scanner_errposition(m, v76, l4)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_read_into_target_3), int32(3626), int32(_a_F_read_into_target_4))
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
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
									F_plpgsql_push_back_token(m, v36, l2, l3, l4)
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
						if v42 == int32(0) {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
							v46 = F_NameListToString(m, v45)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v49 = v48
								v50 = v46
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								v52 = F_read_into_scalar_list(m, v50, v49, v51, l2, l3, l4)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v52
									return
								}
							}
						} else {
							v49 = v25
							v50 = v42
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							v52 = F_read_into_scalar_list(m, v50, v49, v51, l2, l3, l4)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v52
								return
							}
						}
					}
				}
			} else {
				v15 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v15)
				v18 = F_plpgsql_yylex(m, l2, l3, l4)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = v18
					if v20 != int32(277) {
						switch v20 - int32(275) {
						case 0:
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							F_word_is_not_variable(m, l2, v55, l4)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						case 1:
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							F_cword_is_not_variable(m, l2, v58, l4)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						default:
							F_plpgsql_yyerror(m, l3, int32(0), l4, int32(_a_F_read_into_target_0))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
						v27 = int32(1)
						if base.Ui32(v26-v27) <= base.Ui32(v27) {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							F_check_assignable(m, v25, v31, l4)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v34
								v36 = F_plpgsql_yylex(m, l2, l3, l4)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									if v36 == int32(44) {
										F_errstart_cold(m, int32(21), int32(_a_F_read_into_target_1))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											F_errcode(m, int32(16801924))
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_read_into_target_2), int32(0))
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return
												} else {
													v76 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
													v77 = F_plpgsql_scanner_errposition(m, v76, l4)
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_read_into_target_3), int32(3626), int32(_a_F_read_into_target_4))
														mBase = m.M
														v83 = m.ExcPending
														if v83 != 0 {
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
										F_plpgsql_push_back_token(m, v36, l2, l3, l4)
										mBase = m.M
										v41 = m.ExcPending
										if v41 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
							if v42 == int32(0) {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
								v46 = F_NameListToString(m, v45)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return
								} else {
									v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									v49 = v48
									v50 = v46
									v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									v52 = F_read_into_scalar_list(m, v50, v49, v51, l2, l3, l4)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = v52
										return
									}
								}
							} else {
								v49 = v25
								v50 = v42
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								v52 = F_read_into_scalar_list(m, v50, v49, v51, l2, l3, l4)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v52
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v18 = F_plpgsql_yylex(m, l2, l3, l4)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v20 = v18
			if v20 != int32(277) {
				switch v20 - int32(275) {
				case 0:
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					F_word_is_not_variable(m, l2, v55, l4)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				case 1:
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					F_cword_is_not_variable(m, l2, v58, l4)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				default:
					F_plpgsql_yyerror(m, l3, int32(0), l4, int32(_a_F_read_into_target_0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				v27 = int32(1)
				if base.Ui32(v26-v27) <= base.Ui32(v27) {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					F_check_assignable(m, v25, v31, l4)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v34
						v36 = F_plpgsql_yylex(m, l2, l3, l4)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							if v36 == int32(44) {
								F_errstart_cold(m, int32(21), int32(_a_F_read_into_target_1))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									F_errcode(m, int32(16801924))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_read_into_target_2), int32(0))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
											v77 = F_plpgsql_scanner_errposition(m, v76, l4)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_read_into_target_3), int32(3626), int32(_a_F_read_into_target_4))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
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
								F_plpgsql_push_back_token(m, v36, l2, l3, l4)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
					if v42 == int32(0) {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
						v46 = F_NameListToString(m, v45)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v49 = v48
							v50 = v46
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							v52 = F_read_into_scalar_list(m, v50, v49, v51, l2, l3, l4)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v52
								return
							}
						}
					} else {
						v49 = v25
						v50 = v42
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						v52 = F_read_into_scalar_list(m, v50, v49, v51, l2, l3, l4)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v52
							return
						}
					}
				}
			}
		}
	}
}
func F_read_relmap_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	v8 = m.G0
	v10 = v8 - int32(1136)
	m.G0 = v10
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_read_relmap_file[0]))
	v19 = F_LWLockAcquire(m, v15+int32(3200), int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = int32(_a_F_read_relmap_file_0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = l1
	v25 = v10 + int32(112)
	v30 = F_pg_snprintf(m, v25, int32(1024), int32(_a_F_read_relmap_file_1), v10+int32(96))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v33 = F_OpenTransientFile(m, v25, int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L8
	}
L7:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_read_relmap_file[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = int32(167772200)
	v59 = int32(524)
	v60 = F_read(m, v33, l0, v59)
	mBase = m.M
	if v60 == v59 {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	if int32(0) <= v33 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v38 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	if v38 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v25
	F_errmsg(m, int32(_a_F_read_relmap_file_8), v10+int32(80))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_read_relmap_file_4), int32(819), int32(_a_F_read_relmap_file_5))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L7
L15:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_read_relmap_file[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(0)
	v108 = F_CloseTransientFile(m, v33)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L30
	}
L16:
	;
	v64 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	if v60 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	F_errfinish(m, int32(_a_F_read_relmap_file_4), v99, int32(_a_F_read_relmap_file_5))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L28
	}
L19:
	;
	if v64 == int32(0) {
		goto L15
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v64 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L22:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v10 + int32(112)
	F_errmsg(m, int32(_a_F_read_relmap_file_9), v10+int32(48))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v99 = int32(829)
	goto L18
L25:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = int32(524)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v10 + int32(112)
	F_errmsg(m, int32(_a_F_read_relmap_file_10), v10-int32(-64))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v99 = int32(834)
	goto L18
L28:
	;
	goto L15
L29:
	;
	if l2 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L30:
	;
	if v108 == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v113 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	if v113 == int32(0) {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v10 + int32(112)
	F_errmsg(m, int32(_a_F_read_relmap_file_7), v10+int32(32))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_read_relmap_file_4), int32(842), int32(_a_F_read_relmap_file_5))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	goto L29
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_read_relmap_file[0]))
	F_LWLockRelease(m, v135+int32(3200))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v140 == int32(_a_F_read_relmap_file_2) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L39
L41:
	;
	v164 = int32(-1)
	v166 = m.Env.Pgmem_crc32c(m, v164, l0, int32(520))
	mBase = m.M
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
	if v166^v167 == v164 {
		goto L50
	} else {
		goto L51
	}
L42:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v143) < base.Ui32(int32(65)) {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v147 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	if v147 == int32(0) {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v10 + int32(112)
	F_errmsg(m, int32(_a_F_read_relmap_file_6), v10+int32(16))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_read_relmap_file_4), int32(853), int32(_a_F_read_relmap_file_5))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	goto L41
L50:
	;
	m.G0 = v10 + int32(1136)
	return
L51:
	;
	v172 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	if v172 == int32(0) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v10 + int32(112)
	F_errmsg(m, int32(_a_F_read_relmap_file_3), v10)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_read_relmap_file_4), int32(863), int32(_a_F_read_relmap_file_5))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	goto L50
}
func F_readtup_datum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = l3 - int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v13
	if v13 != 0 {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
		if v15 == int32(0) {
			v20 = F_LogicalTapeRead(m, l2, l1+int32(4), v13)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				if v20 == v13 {
					v46 = v5
					v47 = v5
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v46
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v47)
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
					if v50&int32(1) != 0 {
						v56 = F_LogicalTapeRead(m, l2, v10+int32(12), int32(4))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							if v56 != int32(4) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return
								} else {
									F_errmsg_internal(m, int32(_a_F_readtup_datum_0), int32(0))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_readtup_datum_1), int32(2088), int32(_a_F_readtup_datum_2))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								m.G0 = v10 + int32(16)
								return
							}
						}
					} else {
						m.G0 = v10 + int32(16)
						return
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_readtup_datum_0), int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_readtup_datum_1), int32(2073), int32(_a_F_readtup_datum_2))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
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
			v36 = F_tuplesort_readtup_alloc(m, l0, v13)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				v38 = F_LogicalTapeRead(m, l2, v36, v13)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					if v38 != v13 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_readtup_datum_0), int32(0))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_readtup_datum_1), int32(2081), int32(_a_F_readtup_datum_2))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v43 = v36
						v44 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v43
						v46 = v43
						v47 = v44
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v46
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v47)
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
						if v50&int32(1) != 0 {
							v56 = F_LogicalTapeRead(m, l2, v10+int32(12), int32(4))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								if v56 != int32(4) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return
									} else {
										F_errmsg_internal(m, int32(_a_F_readtup_datum_0), int32(0))
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_readtup_datum_1), int32(2088), int32(_a_F_readtup_datum_2))
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							}
						} else {
							m.G0 = v10 + int32(16)
							return
						}
					}
				}
			}
		}
	} else {
		v43 = v5
		v44 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v43
		v46 = v43
		v47 = v44
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v46
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v47)
		v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
		if v50&int32(1) != 0 {
			v56 = F_LogicalTapeRead(m, l2, v10+int32(12), int32(4))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return
			} else {
				if v56 != int32(4) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_readtup_datum_0), int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_readtup_datum_1), int32(2088), int32(_a_F_readtup_datum_2))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					m.G0 = v10 + int32(16)
					return
				}
			}
		} else {
			m.G0 = v10 + int32(16)
			return
		}
	}
}
func F_readtup_index_gin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = l3 - int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v12
	v14 = F_tuplesort_readtup_alloc(m, l0, v12)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v14))) = v12
		v17 = F_LogicalTapeRead(m, l2, v14, v12)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			if v17 == v12 {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
				if v20&int32(1) != 0 {
					v26 = F_LogicalTapeRead(m, l2, v9+int32(12), int32(4))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						if v26 != int32(4) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F_readtup_index_gin_0), int32(0))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_readtup_index_gin_1), int32(1972), int32(_a_F_readtup_index_gin_2))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v14
							m.G0 = v9 + int32(16)
							return
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v14
					m.G0 = v9 + int32(16)
					return
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_readtup_index_gin_0), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_readtup_index_gin_1), int32(1970), int32(_a_F_readtup_index_gin_2))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
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
func F_ready_file_comparator(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	v4 = int32(0)
	v5 = F_strlen(m, l0)
	mBase = m.M
	if v5 != int32(16) {
		v125 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v126 = F_strlen(m, l1)
	mBase = m.M
	if v126 == int32(16) {
		goto L33
	} else {
		goto L34
	}
L2:
	;
	v8 = int32(_a_F_ready_file_comparator_0)
	v12 = m.G0
	v14 = v12 - int32(32)
	v15 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v15
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ready_file_comparator[0])))
	if v23 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v91 != int32(8) {
		v125 = v4
		goto L1
	} else {
		goto L22
	}
L4:
	;
	v91 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ready_file_comparator[1])))
	if v27 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v31 = l0
	goto L10
L8:
	;
	goto L9
L9:
	;
	v41 = v8
	v42 = v23
	goto L13
L10:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v37 == v23 {
		v31 = v31 + int32(1)
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v91 = v31 - l0
	goto L3
L12:
	;
	goto L11
L13:
	;
	v49 = v14 + int32(base.Ui32(v42)>>(uint(int32(3))%32))&int32(28)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v51 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v50 | v51<<(uint(v42)%32)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v55 != 0 {
		v41 = v41 + v51
		v42 = v55
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v58 == int32(0) {
		v81 = l0
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v91 = v81 - l0
	goto L3
L17:
	;
	v62 = l0
	v63 = v58
	goto L18
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(base.Ui32(v63)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v71)>>(uint(v63)%32))&int32(1) == int32(0) {
		v81 = v62
		goto L16
	} else {
		goto L20
	}
L19:
	;
	v81 = v79
	goto L16
L20:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	v79 = v62 + int32(1)
	if v77 != 0 {
		v62 = v79
		v63 = v77
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v95 = l0 + int32(8)
	v96 = int32(_a_F_ready_file_comparator_1)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ready_file_comparator[2])))
	if base.B2i32(v99 == int32(0))|base.B2i32(v99 != v102) != 0 {
		v120 = v99
		v121 = v102
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v125 = base.B2i32(v120-v121 == int32(0))
	goto L1
L24:
	;
	goto L23
L25:
	;
	v105 = v95
	v106 = v96
	goto L26
L26:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+1)))
	if v110 == int32(0) {
		v120 = v110
		v121 = v109
		goto L24
	} else {
		goto L28
	}
L27:
	;
	v120 = v110
	v121 = v109
	goto L24
L28:
	;
	v113 = int32(1)
	if v110 == v109 {
		v105 = v105 + v113
		v106 = v106 + v113
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v253 == int32(0))|base.B2i32(v253 != v256) != 0 {
		v274 = v253
		v275 = v256
		goto L69
	} else {
		goto L70
	}
L31:
	;
	if v125 != 0 {
		goto L65
	} else {
		goto L66
	}
L32:
	;
	v216 = l1 + int32(8)
	v217 = int32(_a_F_ready_file_comparator_1)
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ready_file_comparator[2])))
	if base.B2i32(v220 == int32(0))|base.B2i32(v220 != v223) != 0 {
		v241 = v220
		v242 = v223
		goto L58
	} else {
		goto L59
	}
L33:
	;
	v129 = int32(_a_F_ready_file_comparator_0)
	v133 = m.G0
	v135 = v133 - int32(32)
	v136 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v135)+24)) = v136
	*(*int64)(unsafe.Add(mBase, uint32(v135)+16)) = v136
	*(*int64)(unsafe.Add(mBase, uint32(v135)+8)) = v136
	*(*int64)(unsafe.Add(mBase, uint32(v135))) = v136
	v144 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ready_file_comparator[0])))
	if v144 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L35
L35:
	;
	if v125 != 0 {
		goto L31
	} else {
		goto L56
	}
L36:
	;
	if v212 == int32(8) {
		goto L32
	} else {
		goto L55
	}
L37:
	;
	v212 = int32(0)
	goto L36
L38:
	;
	goto L39
L39:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ready_file_comparator[1])))
	if v148 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v152 = l1
	goto L43
L41:
	;
	goto L42
L42:
	;
	v162 = v129
	v163 = v144
	goto L46
L43:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	if v158 == v144 {
		v152 = v152 + int32(1)
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v212 = v152 - l1
	goto L36
L45:
	;
	goto L44
L46:
	;
	v170 = v135 + int32(base.Ui32(v163)>>(uint(int32(3))%32))&int32(28)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v172 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v171 | v172<<(uint(v163)%32)
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
	if v176 != 0 {
		v162 = v162 + v172
		v163 = v176
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v179 == int32(0) {
		v202 = l1
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	v212 = v202 - l1
	goto L36
L50:
	;
	v183 = l1
	v184 = v179
	goto L51
L51:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v135+int32(base.Ui32(v184)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v192)>>(uint(v184)%32))&int32(1) == int32(0) {
		v202 = v183
		goto L49
	} else {
		goto L53
	}
L52:
	;
	v202 = v200
	goto L49
L53:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+1)))
	v200 = v183 + int32(1)
	if v198 != 0 {
		v183 = v200
		v184 = v198
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	goto L35
L56:
	;
	goto L30
L57:
	;
	if v125 == base.B2i32(v241-v242 == int32(0)) {
		goto L30
	} else {
		goto L64
	}
L58:
	;
	goto L57
L59:
	;
	v226 = v216
	v227 = v217
	goto L60
L60:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+1)))
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+1)))
	if v231 == int32(0) {
		v241 = v231
		v242 = v230
		goto L58
	} else {
		goto L62
	}
L61:
	;
	v241 = v231
	v242 = v230
	goto L58
L62:
	;
	v234 = int32(1)
	if v231 == v230 {
		v226 = v226 + v234
		v227 = v227 + v234
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	goto L31
L65:
	;
	v249 = int32(-1)
	goto L67
L66:
	;
	v249 = int32(1)
	goto L67
L67:
	;
	return v249
L68:
	;
	return v274 - v275
L69:
	;
	goto L68
L70:
	;
	v259 = l0
	v260 = l1
	goto L71
L71:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+1)))
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+1)))
	if v264 == int32(0) {
		v274 = v264
		v275 = v263
		goto L69
	} else {
		goto L73
	}
L72:
	;
	v274 = v264
	v275 = v263
	goto L69
L73:
	;
	v267 = int32(1)
	if v264 == v263 {
		v259 = v259 + v267
		v260 = v260 + v267
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
}
func F_rebuild_joinclause_attr_needed(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(int32(2)) <= base.Ui32(v9) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = v9
	v16 = int32(1)
	v19 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v16<<(uint(int32(2))%32))))
	if v25 == int32(0) {
		v140 = v15
		v144 = v19
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v147 = v16 + int32(1)
	if base.Ui32(v147) < base.Ui32(v140) {
		v15 = v140
		v16 = v147
		v19 = v144
		goto L4
	} else {
		goto L48
	}
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v28 != 0 {
		v140 = v15
		v144 = v19
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+212))
	if v29 == int32(0) {
		v140 = v15
		v144 = v19
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v32 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v32 < v33 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = v32
	v42 = v19
	goto L13
L11:
	;
	v135 = v19
	goto L12
L12:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v140 = v137
	v144 = v135
	goto L6
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v38<<(uint(int32(2))%32))))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+32))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+12)))
	if v50 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v135 = v123
	goto L12
L15:
	;
	v126 = v38 + int32(1)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v126 < v127 {
		v38 = v126
		v42 = v123
		goto L13
	} else {
		goto L47
	}
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+56))
	v54 = F_bms_is_member(m, v53, v42)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v59 = v42
	goto L18
L18:
	;
	v60 = int32(0)
	if v49 == v60 {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	return
L20:
	;
	if v54 != 0 {
		v123 = v42
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)+56))
	v57 = F_bms_add_member(m, v42, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v59 = v57
	goto L18
L23:
	;
	if v105 != int32(2) {
		v123 = v59
		goto L15
	} else {
		goto L39
	}
L24:
	;
	v105 = int32(0)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v68 = int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v69 <= v68 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v72 = v68
	goto L29
L28:
	;
	v72 = v69
	goto L29
L29:
	;
	v76 = int32(0)
	v78 = v60
	goto L30
L30:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v49+int32(8)+v76<<(uint(int32(2))%32))))
	if v85 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v105 = v97
	goto L23
L32:
	;
	goto L31
L33:
	;
	v86 = int32(2)
	if v78 != 0 {
		v97 = v86
		goto L32
	} else {
		goto L36
	}
L34:
	;
	v92 = v78
	goto L35
L35:
	;
	v94 = v76 + int32(1)
	if v94 != v72 {
		v76 = v94
		v78 = v92
		goto L30
	} else {
		goto L38
	}
L36:
	;
	v87 = int32(1)
	if base.Ui32(v87) < base.Ui32(base.I32_popcnt(v85)) {
		v97 = v86
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v92 = v87
	goto L35
L38:
	;
	v97 = v92
	goto L32
L39:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v110 = F_pull_var_clause(m, v108, int32(26))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L19
	} else {
		goto L40
	}
L40:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+12)))
	if v112 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v116 = F_bms_intersect(m, v49, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L19
	} else {
		goto L44
	}
L42:
	;
	v118 = v49
	goto L43
L43:
	;
	F_add_vars_to_attr_needed(m, l0, v110, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L19
	} else {
		goto L45
	}
L44:
	;
	v118 = v116
	goto L43
L45:
	;
	F_list_free(m, v110)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L19
	} else {
		goto L46
	}
L46:
	;
	v123 = v59
	goto L15
L47:
	;
	goto L14
L48:
	;
	goto L5
}
func F_rebuild_lateral_attr_needed(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+317)))
	if v5 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(v8) < base.Ui32(int32(2)) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = v8
	v14 = int32(1)
	goto L4
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16+v14<<(uint(int32(2))%32))))
	if v20 == int32(0) {
		v33 = v13
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	v35 = v14 + int32(1)
	if base.Ui32(v35) < base.Ui32(v33) {
		v13 = v33
		v14 = v35
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v23 != 0 {
		v33 = v13
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+100))
	if v24 == int32(0) {
		v33 = v13
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v27 = F_bms_make_singleton(m, v14)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+100))
	F_add_vars_to_attr_needed(m, l0, v29, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v33 = v32
	goto L6
L13:
	;
	goto L5
}
func F_rebuild_placeholder_attr_needed(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v6 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v9 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = v2
	goto L4
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17+v14<<(uint(int32(2))%32))))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v25 = F_pull_var_clause(m, v23, int32(26))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	F_add_vars_to_attr_needed(m, l0, v25, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F_list_free(m, v25)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v33 = v14 + int32(1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v33 < v34 {
		v14 = v33
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L5
}
func F_record_gt(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(int32(0) < v2)
	}
}
func F_recurse_set_operations(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 float64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
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
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v391 int32
	_ = v391
	var v401 int32
	_ = v401
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v546 int32
	_ = v546
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v581 int32
	_ = v581
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v606 int32
	_ = v606
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v644 int32
	_ = v644
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v661 int32
	_ = v661
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 float64
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
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
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
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
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v942 float64
	_ = v942
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v953 float64
	_ = v953
	var v956 int32
	_ = v956
	var v959 float64
	_ = v959
	var v960 float64
	_ = v960
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1042 float64
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1050 float64
	_ = v1050
	var v1055 float64
	_ = v1055
	var v1056 float64
	_ = v1056
	var v1058 float64
	_ = v1058
	var v1061 float64
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1080 int32
	_ = v1080
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1106 int32
	_ = v1106
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1142 int32
	_ = v1142
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1166 int32
	_ = v1166
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1201 int32
	_ = v1201
	var v1209 int32
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1226 int32
	_ = v1226
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1264 int32
	_ = v1264
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1281 int32
	_ = v1281
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1318 int32
	_ = v1318
	var v1322 int32
	_ = v1322
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1354 int32
	_ = v1354
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1378 int32
	_ = v1378
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1394 int32
	_ = v1394
	var v1398 int32
	_ = v1398
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1413 int32
	_ = v1413
	var v1421 int32
	_ = v1421
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1438 int32
	_ = v1438
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1476 int32
	_ = v1476
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1493 int32
	_ = v1493
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
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
	var v1559 int32
	_ = v1559
	var v1562 float64
	_ = v1562
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1574 float64
	_ = v1574
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1582 int32
	_ = v1582
	var v1585 int32
	_ = v1585
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1594 int32
	_ = v1594
	var v1598 int32
	_ = v1598
	var v1604 int32
	_ = v1604
	var v1611 int32
	_ = v1611
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1648 int32
	_ = v1648
	var v1654 int32
	_ = v1654
	var v1660 int32
	_ = v1660
	var v1682 int32
	_ = v1682
	var v1685 int32
	_ = v1685
	var v1691 int32
	_ = v1691
	var v1697 int32
	_ = v1697
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1723 int32
	_ = v1723
	var v1727 int32
	_ = v1727
	var v1756 int32
	_ = v1756
	var v1761 int32
	_ = v1761
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1766 int32
	_ = v1766
	var v1791 int32
	_ = v1791
	var v1800 int32
	_ = v1800
	var v1802 int32
	_ = v1802
	var v1806 int32
	_ = v1806
	var v1826 int32
	_ = v1826
	var v1834 int32
	_ = v1834
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1846 int32
	_ = v1846
	var v1850 int32
	_ = v1850
	var v1852 int32
	_ = v1852
	var v1856 int32
	_ = v1856
	var v1859 int32
	_ = v1859
	var v1895 int32
	_ = v1895
	var v1898 int32
	_ = v1898
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1904 int32
	_ = v1904
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1919 int32
	_ = v1919
	var v1923 int32
	_ = v1923
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1939 int32
	_ = v1939
	var v1942 int32
	_ = v1942
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1950 int32
	_ = v1950
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1959 int32
	_ = v1959
	var v1963 int32
	_ = v1963
	var v1969 int32
	_ = v1969
	var v1970 float64
	_ = v1970
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1990 int32
	_ = v1990
	var v1992 int32
	_ = v1992
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v2001 int32
	_ = v2001
	var v2002 int32
	_ = v2002
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2014 int32
	_ = v2014
	var v2017 int32
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2029 int32
	_ = v2029
	var v2038 int32
	_ = v2038
	var v2039 int32
	_ = v2039
	var v2040 int32
	_ = v2040
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2046 int32
	_ = v2046
	var v2056 int32
	_ = v2056
	var v2077 int32
	_ = v2077
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2083 int32
	_ = v2083
	var v2085 int32
	_ = v2085
	var v2088 int32
	_ = v2088
	var v2093 int32
	_ = v2093
	var v2104 int32
	_ = v2104
	var v2121 int32
	_ = v2121
	var v2125 int32
	_ = v2125
	var v2126 int32
	_ = v2126
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2137 int32
	_ = v2137
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2145 int32
	_ = v2145
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2164 int32
	_ = v2164
	var v2213 int32
	_ = v2213
	var v2244 int32
	_ = v2244
	var v2246 int32
	_ = v2246
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2252 int32
	_ = v2252
	var v2254 int32
	_ = v2254
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2262 int32
	_ = v2262
	var v2268 int32
	_ = v2268
	var v2296 int32
	_ = v2296
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2340 int32
	_ = v2340
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2347 int32
	_ = v2347
	var v2377 int32
	_ = v2377
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2426 int32
	_ = v2426
	var v2428 int32
	_ = v2428
	var v2438 int32
	_ = v2438
	v9 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(80)
	m.G0 = v33
	v35 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v35)
	F_check_stack_depth(m)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v41 != int32(142) {
		goto L11
	} else {
		goto L12
	}
L3:
	;
	m.G0 = v33 + int32(80)
	return v2438
L4:
	;
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v2079 = F_tlist_same_datatypes(m, v2077, l3, int32(0))
	mBase = m.M
	v2080 = m.ExcPending
	if v2080 != 0 {
		goto L1
	} else {
		goto L534
	}
L5:
	;
	v1554 = F_fetch_upper_rel(m, l1, int32(0), v1532)
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L1
	} else {
		goto L409
	}
L6:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+48))
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1024)+48))
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v1024)+8))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+8))
	v1032 = F_bms_union(m, v1030, v1031)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L1
	} else {
		goto L242
	}
L7:
	;
	v1023 = v738
	v1024 = v729
	v1025 = v745
	goto L6
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L1
	} else {
		goto L239
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L1
	} else {
		goto L231
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L1
	} else {
		goto L228
	}
L11:
	;
	if v41 != int32(63) {
		goto L8
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v78 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46+v47<<(uint(int32(2))%32))))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+36))
	v54 = F_build_simple_rel(m, l1, v47, int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v58 = *(*float64)(unsafe.Add(mBase, uint32(l1)+296))
	v59 = F_subquery_planner(m, v56, v52, l1, int32(0), v58, l2)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+140)) = v59
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v62 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)+264))
	v68 = F_generate_setop_tlist(m, l3, l4, v63, int32(1), v65, l5, v33+int32(56))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v70 = F_make_pathtarget_from_tlist(m, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v72 = F_set_pathtarget_cost_width(m, l1, v70)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+28)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v68
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+56)))
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v76)
	v2438 = v54
	goto L3
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+56)) = l0
	v83 = int32(0)
	v87 = F_list_make1_impl(m, int32(1), v33+int32(12))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v719 = *(*float64)(unsafe.Add(mBase, uint32(l1)+296))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+296)) = int64(0)
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v729 = F_recurse_set_operations(m, v722, l1, l0, v723, v724, l5, v33+int32(76), v33+int32(71))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L169
	}
L24:
	;
	if v87 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v91 = v83
	v97 = v87
	v103 = v9
	v104 = v9
	goto L28
L26:
	;
	v180 = v83
	v192 = v9
	v193 = v9
	goto L27
L27:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v210 = F_generate_append_tlist(m, v208, v209, v192, l5)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L50
	}
L28:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v121 = F_list_delete_first(m, v97)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	v180 = v173
	v192 = v176
	v193 = v177
	goto L27
L30:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v123 != int32(142) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v174 != 0 {
		v91 = v173
		v97 = v174
		v103 = v176
		v104 = v177
		goto L28
	} else {
		goto L49
	}
L32:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v155 != 0 {
		goto L42
	} else {
		goto L43
	}
L33:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v126 != v127 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+8)))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if base.B2i32(v129 != v130)&base.B2i32(v129 == int32(0)) != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v120)+20))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v137 = F_equal(m, v135, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v137 == int32(0) {
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v120)+28))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v143 = F_equal(m, v141, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v143 == int32(0) {
		goto L32
	} else {
		goto L39
	}
L39:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	v148 = F_lcons(m, v147, v121)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v151 = F_lcons(m, v150, v148)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v173 = v91
	v174 = v151
	v176 = v103
	v177 = v104
	goto L31
L42:
	;
	v156 = int32(0)
	goto L44
L43:
	;
	v156 = l0
	goto L44
L44:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v163 = F_recurse_set_operations(m, v120, l1, v156, v157, v158, l5, v33+int32(48), v33+int32(76))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v165 = F_lappend(m, v91, v163)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v168 = F_lappend(m, v103, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+76)))
	v171 = F_lappend_int(m, v104, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v173 = v165
	v174 = v121
	v176 = v168
	v177 = v171
	goto L31
L49:
	;
	goto L29
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v210
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v214 != 0 {
		v374 = v9
		v380 = v9
		v391 = int32(0)
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v401 = int32(0)
	goto L81
L52:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v216 = F_copyObjectImpl(m, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v216 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	v220 = v218
	goto L56
L55:
	;
	v220 = int32(0)
	goto L56
L56:
	;
	if v210 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v308 = int32(0)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v309 == v308 {
		goto L67
	} else {
		goto L68
	}
L58:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	if v223 <= int32(0) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v234 = v220
	v238 = v9
	goto L60
L60:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	v260 = int32(2)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v259+v238<<(uint(v260)%32))))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v258)+4)) = v264
	v267 = v234 + int32(4)
	if base.Ui32(v267) < base.Ui32(v256+v257<<(uint(v260)%32)) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L57
L62:
	;
	v273 = v267
	goto L64
L63:
	;
	v273 = int32(0)
	goto L64
L64:
	;
	v275 = v238 + int32(1)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	if v275 < v276 {
		v234 = v273
		v238 = v275
		goto L60
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	if v354 == int32(0) {
		v374 = v216
		v380 = v9
		v391 = v308
		goto L51
	} else {
		goto L79
	}
L67:
	;
	v354 = int32(1)
	goto L66
L68:
	;
	goto L69
L69:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	if v318 <= int32(0) {
		v346 = int32(1)
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v354 = v346
	goto L66
L71:
	;
	v321 = int32(0)
	if v321 < v318 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v324 = v318
	goto L74
L73:
	;
	v324 = v321
	goto L74
L74:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v309)+12))
	v327 = int32(0)
	goto L75
L75:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v325+v327<<(uint(int32(2))%32))))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+12))
	v337 = int32(0)
	v338 = base.B2i32(v336 != v337)
	if v336 == v337 {
		v346 = v338
		goto L70
	} else {
		goto L77
	}
L76:
	;
	v346 = v338
	goto L70
L77:
	;
	v342 = v327 + int32(1)
	if v342 != v324 {
		v327 = v342
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v357 = F_make_pathkeys_for_sortclauses(m, l1, v216, v210)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+156)) = v357
	v374 = v216
	v380 = v357
	v391 = int32(1)
	goto L51
L81:
	;
	v423 = int32(0)
	if v180 == v423 {
		v433 = v423
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v483 = int32(0)
	v495 = int32(1)
	v496 = v483
	v499 = v483
	v501 = v459
	v502 = v483
	v503 = v9
	v507 = v391
	v510 = v9
	goto L104
L83:
	;
	v434 = int32(0)
	if v193 == v434 {
		v443 = v434
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v427 <= v401 {
		v433 = int32(0)
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
	v433 = v429 + v401<<(uint(int32(2))%32)
	goto L83
L86:
	;
	if v192 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	if v437 <= v401 {
		v443 = v434
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v193)+12))
	v443 = v439 + v401<<(uint(int32(2))%32)
	goto L86
L89:
	;
	goto L82
L90:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+76))
	if v468 == int32(1) {
		goto L100
	} else {
		goto L101
	}
L91:
	;
	if v180 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L92:
	;
	v446 = int32(0)
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	if base.B2i32(v443 == v446)|(base.B2i32(v433 == v446)|base.B2i32(v450 <= v401)) != 0 {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v192)+12))
	if v454 != 0 {
		goto L90
	} else {
		goto L94
	}
L94:
	;
	goto L91
L95:
	;
	v465 = int32(0)
	v1531 = int32(1)
	v1532 = v465
	v1537 = v463
	v1538 = v465
	v1539 = v9
	v1543 = v391
	v1546 = v9
	goto L5
L96:
	;
	v463 = int32(1)
	goto L95
L97:
	;
	goto L98
L98:
	;
	v459 = int32(1)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if int32(0) < v460 {
		goto L89
	} else {
		goto L99
	}
L99:
	;
	v463 = v459
	goto L95
L100:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v443)))
	v472 = int32(0)
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v454+v401<<(uint(int32(2))%32))))
	F_build_setop_child_paths(m, l1, v467, base.B2i32(v471 != v472), v477, v380, v472)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v401 = v401 + int32(1)
	goto L81
L103:
	;
	goto L102
L104:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v517+v499<<(uint(int32(2))%32))))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v521)+48))
	v523 = F_lappend(m, v502, v522)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L106
	}
L105:
	;
	v1531 = v708
	v1532 = v713
	v1537 = v709
	v1538 = v523
	v1539 = v710
	v1543 = v687
	v1546 = v688
	goto L5
L106:
	;
	if v507&int32(1) == int32(0) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	if v501&int32(1) == int32(0) {
		goto L157
	} else {
		goto L158
	}
L108:
	;
	v687 = int32(0)
	v688 = v510
	goto L107
L109:
	;
	goto L110
L110:
	;
	v530 = int32(0)
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v521)+32))
	if v531 == v530 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v680 == int32(0) {
		v687 = v530
		v688 = v510
		goto L107
	} else {
		goto L154
	}
L112:
	;
	v680 = int32(0)
	goto L111
L113:
	;
	goto L114
L114:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v531)+4))
	if int32(0) < v546 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v556 = v530
	v559 = v530
	goto L118
L116:
	;
	v661 = v530
	goto L117
L117:
	;
	v680 = v661
	goto L111
L118:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v531)+12))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v562+v559<<(uint(int32(2))%32))))
	goto L122
L119:
	;
	v661 = v644
	goto L117
L120:
	;
	v651 = v559 + int32(1)
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v531)+4))
	if v651 < v652 {
		v556 = v644
		v559 = v651
		goto L118
	} else {
		goto L153
	}
L122:
	;
	goto L123
L123:
	;
	if v556 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v570 = F_compare_path_costs(m, v556, v566, int32(1))
	mBase = m.M
	if v570 <= int32(0) {
		v644 = v556
		goto L120
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v566)+64))
	if v380 == v573 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	goto L127
L129:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v566)+16))
	if v631 != 0 {
		goto L147
	} else {
		goto L148
	}
L130:
	;
	v581 = int32(0)
	goto L131
L131:
	;
	v589 = int32(0)
	if v380 == v589 {
		v599 = v589
		goto L133
	} else {
		goto L134
	}
L132:
	;
	if v599 != 0 {
		v644 = v556
		goto L120
	} else {
		goto L146
	}
L133:
	;
	if v573 != 0 {
		goto L137
	} else {
		goto L138
	}
L134:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v380)+4))
	if v593 <= v581 {
		v599 = int32(0)
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v380)+12))
	v599 = v595 + v581<<(uint(int32(2))%32)
	goto L133
L136:
	;
	if v599 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L137:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	if v581 < v600 {
		goto L136
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	if v599 == int32(0) {
		goto L129
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	v644 = v556
	goto L120
L142:
	;
	goto L132
L143:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v573)+12))
	if v606 == int32(0) {
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v599)))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v606+v581<<(uint(int32(2))%32))))
	if v613 == v615 {
		v581 = v581 + int32(1)
		goto L131
	} else {
		goto L145
	}
L145:
	;
	v644 = v556
	goto L120
L146:
	;
	goto L129
L147:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v631)+4))
	v634 = v632
	goto L149
L148:
	;
	v634 = int32(0)
	goto L149
L149:
	;
	v635 = F_bms_is_subset(m, v634, v530)
	mBase = m.M
	if v635 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v636 = v566
	goto L152
L151:
	;
	v636 = v556
	goto L152
L152:
	;
	v644 = v636
	goto L120
L153:
	;
	goto L119
L154:
	;
	v684 = F_lappend(m, v510, v680)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v687 = int32(1)
	v688 = v684
	goto L107
L156:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v521)+8))
	v713 = F_bms_union(m, v496, v712)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L167
	}
L157:
	;
	v708 = v495
	v709 = int32(0)
	v710 = v503
	goto L156
L158:
	;
	goto L159
L159:
	;
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521)+26)))
	if v694 != int32(1) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v697 = int32(0)
	v708 = v697
	v709 = v697
	v710 = v503
	goto L156
L161:
	;
	goto L162
L162:
	;
	v699 = int32(1)
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v521)+40))
	if v700 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v708 = int32(0)
	v709 = v699
	v710 = v503
	goto L156
L164:
	;
	goto L165
L165:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v700)+12))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	v706 = F_lappend(m, v503, v705)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	v708 = v495
	v709 = v699
	v710 = v706
	goto L156
L167:
	;
	v716 = v499 + int32(1)
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v716 < v717 {
		v495 = v708
		v496 = v713
		v499 = v716
		v501 = v709
		v502 = v523
		v503 = v710
		v507 = v687
		v510 = v688
		goto L104
	} else {
		goto L168
	}
L168:
	;
	goto L105
L169:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v738 = F_recurse_set_operations(m, v731, l1, l0, v732, v733, l5, v33+int32(72), v33+int32(70))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	v740 = int32(0)
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
	v748 = F_generate_setop_tlist(m, v741, v742, v740, v740, v745, l5, v33+int32(69))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v748
	v751 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v752 = F_copyObjectImpl(m, v751)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	if v752 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v752)+12))
	v755 = v754
	goto L175
L174:
	;
	v755 = v740
	goto L175
L175:
	;
	if v748 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	if v752 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L177:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v748)+4))
	if v758 <= int32(0) {
		goto L176
	} else {
		goto L178
	}
L178:
	;
	v764 = v755
	v770 = int32(0)
	goto L179
L179:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v752)+12))
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v752)+4))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v764)))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v748)+12))
	v796 = int32(2)
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v795+v770<<(uint(v796)%32))))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v799)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v794)+4)) = v800
	v803 = v764 + int32(4)
	if base.Ui32(v803) < base.Ui32(v792+v793<<(uint(v796)%32)) {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	goto L176
L181:
	;
	v809 = v803
	goto L183
L182:
	;
	v809 = int32(0)
	goto L183
L183:
	;
	v811 = v770 + int32(1)
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v748)+4))
	if v811 < v812 {
		v764 = v809
		v770 = v811
		goto L179
	} else {
		goto L184
	}
L184:
	;
	goto L180
L185:
	;
	v889 = int32(0)
	if v752 == v889 {
		goto L199
	} else {
		goto L200
	}
L186:
	;
	v888 = int32(1)
	goto L185
L187:
	;
	goto L188
L188:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v752)+4))
	if v852 <= int32(0) {
		v880 = int32(1)
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v888 = v880
	goto L185
L190:
	;
	v855 = int32(0)
	if v855 < v852 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v858 = v852
	goto L193
L192:
	;
	v858 = v855
	goto L193
L193:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v752)+12))
	v861 = int32(0)
	goto L194
L194:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v859+v861<<(uint(int32(2))%32))))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v869)+12))
	v871 = int32(0)
	v872 = base.B2i32(v870 != v871)
	if v870 == v871 {
		v880 = v872
		goto L189
	} else {
		goto L196
	}
L195:
	;
	v880 = v872
	goto L189
L196:
	;
	v876 = v861 + int32(1)
	if v876 != v858 {
		v861 = v876
		goto L194
	} else {
		goto L197
	}
L197:
	;
	goto L195
L198:
	;
	if v888|v926 == int32(0) {
		goto L9
	} else {
		goto L211
	}
L199:
	;
	v926 = int32(1)
	goto L198
L200:
	;
	goto L201
L201:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v752)+4))
	if v896 <= int32(0) {
		v920 = int32(1)
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v926 = v920
	goto L198
L203:
	;
	v899 = int32(0)
	if v899 < v896 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v902 = v896
	goto L206
L205:
	;
	v902 = v899
	goto L206
L206:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v752)+12))
	v907 = v889
	goto L207
L207:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v903+v907<<(uint(int32(2))%32))))
	v912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v911)+18)))
	if v912 != int32(1) {
		v920 = v912
		goto L202
	} else {
		goto L209
	}
L208:
	;
	v920 = v912
	goto L202
L209:
	;
	v916 = v907 + int32(1)
	if v916 != v902 {
		v907 = v916
		goto L207
	} else {
		goto L210
	}
L210:
	;
	goto L208
L211:
	;
	if v888 != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v930 = F_make_pathkeys_for_sortclauses(m, l1, v752, v748)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L1
	} else {
		goto L215
	}
L213:
	;
	v933 = v9
	goto L214
L214:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v729)+76))
	if v934 == int32(1) {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+156)) = v930
	v933 = v930
	goto L214
L216:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v738)+76))
	if v944 == int32(1) {
		goto L222
	} else {
		goto L223
	}
L217:
	;
	v937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+71)))
	F_build_setop_child_paths(m, l1, v729, v937, v745, v933, v33+int32(56))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L1
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v942 = *(*float64)(unsafe.Add(mBase, uint32(v729)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v33)+56)) = v942
	goto L216
L220:
	;
	goto L216
L221:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+296)) = v719
	v956 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v956 == int32(3) {
		goto L7
	} else {
		goto L226
	}
L222:
	;
	v947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+70)))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
	F_build_setop_child_paths(m, l1, v738, v947, v948, v933, v33+int32(48))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L1
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v953 = *(*float64)(unsafe.Add(mBase, uint32(v738)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v33)+48)) = v953
	goto L221
L225:
	;
	goto L221
L226:
	;
	v959 = *(*float64)(unsafe.Add(mBase, uint32(v33)+48))
	v960 = *(*float64)(unsafe.Add(mBase, uint32(v33)+56))
	if base.F64_lt(v959, v960) == int32(0) {
		goto L7
	} else {
		goto L227
	}
L227:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v33)+56)) = v959
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+72)) = v745
	v1023 = v729
	v1024 = v738
	v1025 = v965
	goto L6
L228:
	;
	F_errmsg_internal(m, int32(_a_F_recurse_set_operations_0), int32(0))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	F_errfinish(m, int32(_a_F_recurse_set_operations_1), int32(254), int32(_a_F_recurse_set_operations_2))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L231:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v989 == int32(2) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v992 = int32(_a_F_recurse_set_operations_3)
	goto L235
L234:
	;
	v992 = int32(_a_F_recurse_set_operations_4)
	goto L235
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v992
	F_errmsg(m, int32(_a_F_recurse_set_operations_5), v33+int32(32))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	F_errdetail(m, int32(_a_F_recurse_set_operations_6), int32(0))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(_a_F_recurse_set_operations_1), int32(1073), int32(_a_F_recurse_set_operations_7))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L239:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v1012
	F_errmsg_internal(m, int32(_a_F_recurse_set_operations_8), v33)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(_a_F_recurse_set_operations_1), int32(345), int32(_a_F_recurse_set_operations_2))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L242:
	;
	v1034 = F_fetch_upper_rel(m, l1, int32(0), v1032)
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	v1036 = F_make_pathtarget_from_tlist(m, v748)
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	v1038 = F_set_pathtarget_cost_width(m, l1, v1036)
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1034)+28)) = v1038
	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v1042 = *(*float64)(unsafe.Add(mBase, uint32(v33)+56))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1043 == int32(3) {
		goto L247
	} else {
		goto L248
	}
L246:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1034)+16)) = v1061
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v1063 - int32(2) {
	case 0:
		goto L256
	case 1:
		goto L258
	default:
		goto L257
	}
L247:
	;
	if v1041&int32(1) == int32(0) {
		v1061 = v1042
		goto L246
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	if v1041&int32(1) == int32(0) {
		v1061 = v1042
		goto L246
	} else {
		goto L251
	}
L250:
	;
	v1050 = *(*float64)(unsafe.Add(mBase, uint32(v1028)+32))
	v1061 = v1050
	goto L246
L251:
	;
	v1055 = *(*float64)(unsafe.Add(mBase, uint32(v1028)+32))
	v1056 = *(*float64)(unsafe.Add(mBase, uint32(v1027)+32))
	if base.F64_lt(v1055, v1056) != 0 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1058 = v1055
	goto L254
L253:
	;
	v1058 = v1056
	goto L254
L254:
	;
	v1061 = v1058
	goto L246
L255:
	;
	if v926 != 0 {
		goto L265
	} else {
		goto L266
	}
L256:
	;
	v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v1087 = v1086
	goto L255
L257:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L1
	} else {
		goto L262
	}
L258:
	;
	v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v1068 != 0 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v1069 = int32(3)
	goto L261
L260:
	;
	v1069 = int32(2)
	goto L261
L261:
	;
	v1087 = v1069
	goto L255
L262:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v1074
	F_errmsg_internal(m, int32(_a_F_recurse_set_operations_9), v33+int32(16))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	F_errfinish(m, int32(_a_F_recurse_set_operations_1), int32(1165), int32(_a_F_recurse_set_operations_7))
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L265:
	;
	v1089 = F_create_setop_path(m, v1034, v1028, v1027, v1087, int32(1), v752, v1042, v1061)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L1
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	if v888 == int32(0) {
		v2056 = v1034
		goto L4
	} else {
		goto L270
	}
L268:
	;
	F_add_path(m, v1034, v1089)
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	goto L267
L270:
	;
	v1095 = F_make_pathkeys_for_sortclauses(m, l1, v752, v1025)
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L1
	} else {
		goto L272
	}
L271:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
	v1307 = F_make_pathkeys_for_sortclauses(m, l1, v752, v1306)
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L1
	} else {
		goto L340
	}
L272:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1028)+64))
	if v1095 == v1097 {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	if v1150 != 0 {
		goto L291
	} else {
		goto L292
	}
L274:
	;
	v1150 = int32(1)
	goto L273
L275:
	;
	goto L276
L276:
	;
	v1106 = int32(0)
	goto L278
L277:
	;
	v1150 = v1142
	goto L273
L278:
	;
	v1110 = int32(0)
	if v1095 == v1110 {
		v1120 = v1110
		goto L280
	} else {
		goto L281
	}
L279:
	;
	v1142 = int32(0)
	goto L277
L280:
	;
	if v1097 != 0 {
		goto L284
	} else {
		goto L285
	}
L281:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1095)+4))
	if v1114 <= v1106 {
		v1120 = int32(0)
		goto L280
	} else {
		goto L282
	}
L282:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1095)+12))
	v1120 = v1116 + v1106<<(uint(int32(2))%32)
	goto L280
L283:
	;
	v1126 = base.B2i32(v1120 == int32(0))
	if v1120 == int32(0) {
		v1142 = v1126
		goto L277
	} else {
		goto L288
	}
L284:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+4))
	if v1106 < v1121 {
		goto L283
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	v1150 = base.B2i32(v1120 == int32(0))
	goto L273
L287:
	;
	goto L286
L288:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+12))
	if v1129 == int32(0) {
		v1142 = v1126
		goto L277
	} else {
		goto L289
	}
L289:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1120)))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1106<<(uint(int32(2))%32)+v1129)))
	if v1136 == v1138 {
		v1106 = v1106 + int32(1)
		goto L278
	} else {
		goto L290
	}
L290:
	;
	goto L279
L291:
	;
	v1305 = v1028
	goto L271
L292:
	;
	goto L293
L293:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1024)+32))
	v1152 = int32(0)
	if v1151 == v1152 {
		goto L295
	} else {
		goto L296
	}
L294:
	;
	if v1300 != 0 {
		v1305 = v1300
		goto L271
	} else {
		goto L337
	}
L295:
	;
	v1300 = int32(0)
	goto L294
L296:
	;
	goto L297
L297:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+4))
	if int32(0) < v1166 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v1176 = v1152
	v1179 = v1152
	goto L301
L299:
	;
	v1281 = v1152
	goto L300
L300:
	;
	v1300 = v1281
	goto L294
L301:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+12))
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v1182+v1179<<(uint(int32(2))%32))))
	goto L305
L302:
	;
	v1281 = v1264
	goto L300
L303:
	;
	v1271 = v1179 + int32(1)
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+4))
	if v1271 < v1272 {
		v1176 = v1264
		v1179 = v1271
		goto L301
	} else {
		goto L336
	}
L305:
	;
	goto L306
L306:
	;
	if v1176 != 0 {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1190 = F_compare_path_costs(m, v1176, v1186, int32(1))
	mBase = m.M
	if v1190 <= int32(0) {
		v1264 = v1176
		goto L303
	} else {
		goto L311
	}
L309:
	;
	goto L310
L310:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+64))
	if v933 == v1193 {
		goto L312
	} else {
		goto L313
	}
L311:
	;
	goto L310
L312:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+16))
	if v1251 != 0 {
		goto L330
	} else {
		goto L331
	}
L313:
	;
	v1201 = int32(0)
	goto L314
L314:
	;
	v1209 = int32(0)
	if v933 == v1209 {
		v1219 = v1209
		goto L316
	} else {
		goto L317
	}
L315:
	;
	if v1219 != 0 {
		v1264 = v1176
		goto L303
	} else {
		goto L329
	}
L316:
	;
	if v1193 != 0 {
		goto L320
	} else {
		goto L321
	}
L317:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v933)+4))
	if v1213 <= v1201 {
		v1219 = int32(0)
		goto L316
	} else {
		goto L318
	}
L318:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v933)+12))
	v1219 = v1215 + v1201<<(uint(int32(2))%32)
	goto L316
L319:
	;
	if v1219 == int32(0) {
		goto L325
	} else {
		goto L326
	}
L320:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+4))
	if v1201 < v1220 {
		goto L319
	} else {
		goto L323
	}
L321:
	;
	goto L322
L322:
	;
	if v1219 == int32(0) {
		goto L312
	} else {
		goto L324
	}
L323:
	;
	goto L322
L324:
	;
	v1264 = v1176
	goto L303
L325:
	;
	goto L315
L326:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+12))
	if v1226 == int32(0) {
		goto L325
	} else {
		goto L327
	}
L327:
	;
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1219)))
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1226+v1201<<(uint(int32(2))%32))))
	if v1233 == v1235 {
		v1201 = v1201 + int32(1)
		goto L314
	} else {
		goto L328
	}
L328:
	;
	v1264 = v1176
	goto L303
L329:
	;
	goto L312
L330:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1251)+4))
	v1254 = v1252
	goto L332
L331:
	;
	v1254 = int32(0)
	goto L332
L332:
	;
	v1255 = F_bms_is_subset(m, v1254, v1152)
	mBase = m.M
	if v1255 != 0 {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v1256 = v1186
	goto L335
L334:
	;
	v1256 = v1176
	goto L335
L335:
	;
	v1264 = v1256
	goto L303
L336:
	;
	goto L302
L337:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1028)+8))
	v1303 = F_create_sort_path(m, v1301, v1028, v1095, float64(-1))
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	v1305 = v1303
	goto L271
L339:
	;
	v1519 = F_create_setop_path(m, v1034, v1305, v1517, v1087, int32(0), v752, v1042, v1061)
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L1
	} else {
		goto L407
	}
L340:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+64))
	if v1307 == v1309 {
		goto L342
	} else {
		goto L343
	}
L341:
	;
	if v1362 != 0 {
		goto L359
	} else {
		goto L360
	}
L342:
	;
	v1362 = int32(1)
	goto L341
L343:
	;
	goto L344
L344:
	;
	v1318 = int32(0)
	goto L346
L345:
	;
	v1362 = v1354
	goto L341
L346:
	;
	v1322 = int32(0)
	if v1307 == v1322 {
		v1332 = v1322
		goto L348
	} else {
		goto L349
	}
L347:
	;
	v1354 = int32(0)
	goto L345
L348:
	;
	if v1309 != 0 {
		goto L352
	} else {
		goto L353
	}
L349:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1307)+4))
	if v1326 <= v1318 {
		v1332 = int32(0)
		goto L348
	} else {
		goto L350
	}
L350:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v1307)+12))
	v1332 = v1328 + v1318<<(uint(int32(2))%32)
	goto L348
L351:
	;
	v1338 = base.B2i32(v1332 == int32(0))
	if v1332 == int32(0) {
		v1354 = v1338
		goto L345
	} else {
		goto L356
	}
L352:
	;
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1309)+4))
	if v1318 < v1333 {
		goto L351
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	v1362 = base.B2i32(v1332 == int32(0))
	goto L341
L355:
	;
	goto L354
L356:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v1309)+12))
	if v1341 == int32(0) {
		v1354 = v1338
		goto L345
	} else {
		goto L357
	}
L357:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1332)))
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1318<<(uint(int32(2))%32)+v1341)))
	if v1348 == v1350 {
		v1318 = v1318 + int32(1)
		goto L346
	} else {
		goto L358
	}
L358:
	;
	goto L347
L359:
	;
	v1517 = v1027
	goto L339
L360:
	;
	goto L361
L361:
	;
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+32))
	v1364 = int32(0)
	if v1363 == v1364 {
		goto L363
	} else {
		goto L364
	}
L362:
	;
	if v1512 != 0 {
		v1517 = v1512
		goto L339
	} else {
		goto L405
	}
L363:
	;
	v1512 = int32(0)
	goto L362
L364:
	;
	goto L365
L365:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1363)+4))
	if int32(0) < v1378 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v1388 = v1364
	v1391 = v1364
	goto L369
L367:
	;
	v1493 = v1364
	goto L368
L368:
	;
	v1512 = v1493
	goto L362
L369:
	;
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1363)+12))
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v1394+v1391<<(uint(int32(2))%32))))
	goto L373
L370:
	;
	v1493 = v1476
	goto L368
L371:
	;
	v1483 = v1391 + int32(1)
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v1363)+4))
	if v1483 < v1484 {
		v1388 = v1476
		v1391 = v1483
		goto L369
	} else {
		goto L404
	}
L373:
	;
	goto L374
L374:
	;
	if v1388 != 0 {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v1402 = F_compare_path_costs(m, v1388, v1398, int32(1))
	mBase = m.M
	if v1402 <= int32(0) {
		v1476 = v1388
		goto L371
	} else {
		goto L379
	}
L377:
	;
	goto L378
L378:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1398)+64))
	if v933 == v1405 {
		goto L380
	} else {
		goto L381
	}
L379:
	;
	goto L378
L380:
	;
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v1398)+16))
	if v1463 != 0 {
		goto L398
	} else {
		goto L399
	}
L381:
	;
	v1413 = int32(0)
	goto L382
L382:
	;
	v1421 = int32(0)
	if v933 == v1421 {
		v1431 = v1421
		goto L384
	} else {
		goto L385
	}
L383:
	;
	if v1431 != 0 {
		v1476 = v1388
		goto L371
	} else {
		goto L397
	}
L384:
	;
	if v1405 != 0 {
		goto L388
	} else {
		goto L389
	}
L385:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v933)+4))
	if v1425 <= v1413 {
		v1431 = int32(0)
		goto L384
	} else {
		goto L386
	}
L386:
	;
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v933)+12))
	v1431 = v1427 + v1413<<(uint(int32(2))%32)
	goto L384
L387:
	;
	if v1431 == int32(0) {
		goto L393
	} else {
		goto L394
	}
L388:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v1405)+4))
	if v1413 < v1432 {
		goto L387
	} else {
		goto L391
	}
L389:
	;
	goto L390
L390:
	;
	if v1431 == int32(0) {
		goto L380
	} else {
		goto L392
	}
L391:
	;
	goto L390
L392:
	;
	v1476 = v1388
	goto L371
L393:
	;
	goto L383
L394:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1405)+12))
	if v1438 == int32(0) {
		goto L393
	} else {
		goto L395
	}
L395:
	;
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v1431)))
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1438+v1413<<(uint(int32(2))%32))))
	if v1445 == v1447 {
		v1413 = v1413 + int32(1)
		goto L382
	} else {
		goto L396
	}
L396:
	;
	v1476 = v1388
	goto L371
L397:
	;
	goto L380
L398:
	;
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1463)+4))
	v1466 = v1464
	goto L400
L399:
	;
	v1466 = int32(0)
	goto L400
L400:
	;
	v1467 = F_bms_is_subset(m, v1466, v1364)
	mBase = m.M
	if v1467 != 0 {
		goto L401
	} else {
		goto L402
	}
L401:
	;
	v1468 = v1398
	goto L403
L402:
	;
	v1468 = v1388
	goto L403
L403:
	;
	v1476 = v1468
	goto L371
L404:
	;
	goto L370
L405:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+8))
	v1515 = F_create_sort_path(m, v1513, v1027, v1307, float64(-1))
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L1
	} else {
		goto L406
	}
L406:
	;
	v1517 = v1515
	goto L339
L407:
	;
	F_add_path(m, v1034, v1519)
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L1
	} else {
		goto L408
	}
L408:
	;
	v2056 = v1034
	goto L4
L409:
	;
	v1556 = F_make_pathtarget_from_tlist(m, v210)
	mBase = m.M
	v1557 = m.ExcPending
	if v1557 != 0 {
		goto L1
	} else {
		goto L410
	}
L410:
	;
	v1558 = F_set_pathtarget_cost_width(m, l1, v1556)
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1554)+26)) = uint8(v1537)
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+28)) = v1558
	v1562 = *(*float64)(unsafe.Add(mBase, uint32(l1)+296))
	v1564 = base.F64_gt(v1562, float64(0))
	*(*uint8)(unsafe.Add(mBase, uint32(v1554)+24)) = uint8(v1564)
	v1566 = int32(0)
	v1572 = F_create_append_path(m, l1, v1554, v1538, v1566, v1566, v1566, v1566, v1566, float64(-1))
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	v1574 = *(*float64)(unsafe.Add(mBase, uint32(v1572)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1554)+16)) = v1574
	if v1531&int32(1) != 0 {
		goto L415
	} else {
		goto L416
	}
L413:
	;
	if v374 == int32(0) {
		goto L470
	} else {
		goto L471
	}
L414:
	;
	F_add_path(m, v1554, v1572)
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L1
	} else {
		goto L466
	}
L415:
	;
	if v1539 != 0 {
		goto L420
	} else {
		goto L421
	}
L416:
	;
	goto L417
L417:
	;
	v1846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v1846 != int32(1) {
		v1859 = int32(0)
		goto L413
	} else {
		goto L464
	}
L418:
	;
	v1834 = int32(0)
	v1838 = F_create_append_path(m, l1, v1554, v1834, v1539, v1834, v1834, v1806, v1826, float64(-1))
	mBase = m.M
	v1839 = m.ExcPending
	if v1839 != 0 {
		goto L1
	} else {
		goto L461
	}
L419:
	;
	v1800 = *(*int32)(unsafe.Add(mBase, _c_F_recurse_set_operations[0]))
	if v1791 < v1800 {
		goto L458
	} else {
		goto L459
	}
L420:
	;
	v1578 = int32(0)
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1539)+4))
	if v1579 <= v1578 {
		v1727 = v1578
		goto L423
	} else {
		goto L424
	}
L421:
	;
	goto L422
L422:
	;
	v1764 = int32(0)
	v1766 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_recurse_set_operations[1])))
	if v1766 == v1764 {
		v1806 = v1764
		v1826 = v9
		goto L418
	} else {
		goto L457
	}
L423:
	;
	v1756 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_recurse_set_operations[1])))
	if v1756 != int32(1) {
		v1806 = v1727
		v1826 = v9
		goto L418
	} else {
		goto L453
	}
L424:
	;
	v1582 = int32(0)
	if v1582 < v1579 {
		goto L425
	} else {
		goto L426
	}
L425:
	;
	v1585 = v1579
	goto L427
L426:
	;
	v1585 = v1582
	goto L427
L427:
	;
	v1587 = v1585 & int32(3)
	v1588 = int32(0)
	if int32(4) <= v1579 {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v1539)+12))
	v1598 = v1578
	v1604 = v1588
	v1611 = int32(0)
	goto L431
L429:
	;
	v1654 = v1578
	v1660 = v1588
	goto L430
L430:
	;
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v1539)+12))
	v1685 = v1654
	v1691 = v1660
	v1697 = v1588
	goto L447
L431:
	;
	v1628 = v1594 + v1604<<(uint(int32(2))%32)
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v1628)))
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1629)+24))
	if v1630 < v1598 {
		goto L433
	} else {
		goto L434
	}
L432:
	;
	if v1587 == int32(0) {
		v1727 = v1644
		goto L423
	} else {
		goto L446
	}
L433:
	;
	v1632 = v1598
	goto L435
L434:
	;
	v1632 = v1630
	goto L435
L435:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1628)+4))
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v1633)+24))
	if v1634 < v1632 {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	v1636 = v1632
	goto L438
L437:
	;
	v1636 = v1634
	goto L438
L438:
	;
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v1628)+8))
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1637)+24))
	if v1638 < v1636 {
		goto L439
	} else {
		goto L440
	}
L439:
	;
	v1640 = v1636
	goto L441
L440:
	;
	v1640 = v1638
	goto L441
L441:
	;
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1628)+12))
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(v1641)+24))
	if v1642 < v1640 {
		goto L442
	} else {
		goto L443
	}
L442:
	;
	v1644 = v1640
	goto L444
L443:
	;
	v1644 = v1642
	goto L444
L444:
	;
	v1645 = int32(4)
	v1646 = v1604 + v1645
	v1648 = v1611 + v1645
	if v1648 != v1585&int32(2147483644) {
		v1598 = v1644
		v1604 = v1646
		v1611 = v1648
		goto L431
	} else {
		goto L445
	}
L445:
	;
	goto L432
L446:
	;
	v1654 = v1644
	v1660 = v1646
	goto L430
L447:
	;
	v1716 = *(*int32)(unsafe.Add(mBase, uint32(v1682+v1691<<(uint(int32(2))%32))))
	v1717 = *(*int32)(unsafe.Add(mBase, uint32(v1716)+24))
	if v1717 < v1685 {
		goto L449
	} else {
		goto L450
	}
L448:
	;
	v1727 = v1719
	goto L423
L449:
	;
	v1719 = v1685
	goto L451
L450:
	;
	v1719 = v1717
	goto L451
L451:
	;
	v1720 = int32(1)
	v1723 = v1697 + v1720
	if v1723 != v1587 {
		v1685 = v1719
		v1691 = v1691 + v1720
		v1697 = v1723
		goto L447
	} else {
		goto L452
	}
L452:
	;
	goto L448
L453:
	;
	v1761 = int32(32) - base.I32_clz(v1579)
	if v1761 < v1727 {
		goto L454
	} else {
		goto L455
	}
L454:
	;
	v1763 = v1727
	goto L456
L455:
	;
	v1763 = v1761
	goto L456
L456:
	;
	v1791 = v1763
	goto L419
L457:
	;
	v1791 = v9
	goto L419
L458:
	;
	v1802 = v1791
	goto L460
L459:
	;
	v1802 = v1800
	goto L460
L460:
	;
	v1806 = v1802
	v1826 = int32(1)
	goto L418
L461:
	;
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+28))
	v1842 = F_create_gather_path(m, l1, v1554, v1838, v1840, int32(0))
	mBase = m.M
	v1843 = m.ExcPending
	if v1843 != 0 {
		goto L1
	} else {
		goto L462
	}
L462:
	;
	v1844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v1844 != 0 {
		goto L414
	} else {
		goto L463
	}
L463:
	;
	v1859 = v1842
	goto L413
L464:
	;
	F_add_path(m, v1554, v1572)
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	v2056 = v1554
	goto L4
L466:
	;
	if v1842 == int32(0) {
		v2056 = v1554
		goto L4
	} else {
		goto L467
	}
L467:
	;
	F_add_path(m, v1554, v1842)
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
		goto L1
	} else {
		goto L468
	}
L468:
	;
	v2056 = v1554
	goto L4
L469:
	;
	v1932 = int32(0)
	if v374 == v1932 {
		goto L483
	} else {
		goto L484
	}
L470:
	;
	v1931 = int32(1)
	goto L469
L471:
	;
	goto L472
L472:
	;
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	if v1895 <= int32(0) {
		v1923 = int32(1)
		goto L473
	} else {
		goto L474
	}
L473:
	;
	v1931 = v1923
	goto L469
L474:
	;
	v1898 = int32(0)
	if v1898 < v1895 {
		goto L475
	} else {
		goto L476
	}
L475:
	;
	v1901 = v1895
	goto L477
L476:
	;
	v1901 = v1898
	goto L477
L477:
	;
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(v374)+12))
	v1904 = int32(0)
	goto L478
L478:
	;
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v1902+v1904<<(uint(int32(2))%32))))
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v1912)+12))
	v1914 = int32(0)
	v1915 = base.B2i32(v1913 != v1914)
	if v1913 == v1914 {
		v1923 = v1915
		goto L473
	} else {
		goto L480
	}
L479:
	;
	v1923 = v1915
	goto L473
L480:
	;
	v1919 = v1904 + int32(1)
	if v1919 != v1901 {
		v1904 = v1919
		goto L478
	} else {
		goto L481
	}
L481:
	;
	goto L479
L482:
	;
	v1970 = *(*float64)(unsafe.Add(mBase, uint32(v1572)+32))
	if v1969 == int32(0) {
		goto L495
	} else {
		goto L496
	}
L483:
	;
	v1969 = int32(1)
	goto L482
L484:
	;
	goto L485
L485:
	;
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	if v1939 <= int32(0) {
		v1963 = int32(1)
		goto L486
	} else {
		goto L487
	}
L486:
	;
	v1969 = v1963
	goto L482
L487:
	;
	v1942 = int32(0)
	if v1942 < v1939 {
		goto L488
	} else {
		goto L489
	}
L488:
	;
	v1945 = v1939
	goto L490
L489:
	;
	v1945 = v1942
	goto L490
L490:
	;
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(v374)+12))
	v1950 = v1932
	goto L491
L491:
	;
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(v1946+v1950<<(uint(int32(2))%32))))
	v1955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1954)+18)))
	if v1955 != int32(1) {
		v1963 = v1955
		goto L486
	} else {
		goto L493
	}
L492:
	;
	v1963 = v1955
	goto L486
L493:
	;
	v1959 = v1950 + int32(1)
	if v1959 != v1945 {
		v1950 = v1959
		goto L491
	} else {
		goto L494
	}
L494:
	;
	goto L492
L495:
	;
	if v1931 == int32(0) {
		goto L506
	} else {
		goto L507
	}
L496:
	;
	v1973 = F_make_pathtarget_from_tlist(m, v210)
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	v1975 = F_set_pathtarget_cost_width(m, l1, v1973)
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L1
	} else {
		goto L498
	}
L498:
	;
	v1978 = int32(0)
	v1981 = F_create_agg_path(m, l1, v1554, v1572, v1975, int32(2), v1978, v374, v1978, v1978, v1970)
	mBase = m.M
	v1982 = m.ExcPending
	if v1982 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	F_add_path(m, v1554, v1981)
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		goto L1
	} else {
		goto L500
	}
L500:
	;
	if v1859 == int32(0) {
		goto L495
	} else {
		goto L501
	}
L501:
	;
	v1987 = F_make_pathtarget_from_tlist(m, v210)
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		goto L1
	} else {
		goto L502
	}
L502:
	;
	v1989 = F_set_pathtarget_cost_width(m, l1, v1987)
	mBase = m.M
	v1990 = m.ExcPending
	if v1990 != 0 {
		goto L1
	} else {
		goto L503
	}
L503:
	;
	v1992 = int32(0)
	v1995 = F_create_agg_path(m, l1, v1554, v1859, v1989, int32(2), v1992, v374, v1992, v1992, v1970)
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		goto L1
	} else {
		goto L504
	}
L504:
	;
	F_add_path(m, v1554, v1995)
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L1
	} else {
		goto L505
	}
L505:
	;
	goto L495
L506:
	;
	if base.B2i32(v374 == int32(0))|(v1543^int32(1)) != 0 {
		v2056 = v1554
		goto L4
	} else {
		goto L526
	}
L507:
	;
	if v374 != 0 {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v2001 = F_make_pathkeys_for_sortclauses(m, l1, v374, v210)
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		goto L1
	} else {
		goto L511
	}
L509:
	;
	v2006 = v1572
	goto L510
L510:
	;
	v2007 = *(*int32)(unsafe.Add(mBase, uint32(v2006)+64))
	if v2007 != 0 {
		goto L513
	} else {
		goto L514
	}
L511:
	;
	v2004 = F_create_sort_path(m, v1554, v1572, v2001, float64(-1))
	mBase = m.M
	v2005 = m.ExcPending
	if v2005 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	v2006 = v2004
	goto L510
L513:
	;
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(v2007)+4))
	v2010 = v2008
	goto L515
L514:
	;
	v2010 = int32(0)
	goto L515
L515:
	;
	v2011 = F_create_upper_unique_path(m, v1554, v2006, v2010, v1970)
	mBase = m.M
	v2012 = m.ExcPending
	if v2012 != 0 {
		goto L1
	} else {
		goto L516
	}
L516:
	;
	F_add_path(m, v1554, v2011)
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		goto L1
	} else {
		goto L517
	}
L517:
	;
	if v1859 == int32(0) {
		goto L506
	} else {
		goto L518
	}
L518:
	;
	v2017 = F_make_pathkeys_for_sortclauses(m, l1, v374, v210)
	mBase = m.M
	v2018 = m.ExcPending
	if v2018 != 0 {
		goto L1
	} else {
		goto L519
	}
L519:
	;
	v2020 = F_create_sort_path(m, v1554, v1859, v2017, float64(-1))
	mBase = m.M
	v2021 = m.ExcPending
	if v2021 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	v2022 = *(*int32)(unsafe.Add(mBase, uint32(v2020)+64))
	if v2022 != 0 {
		goto L521
	} else {
		goto L522
	}
L521:
	;
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(v2022)+4))
	v2025 = v2023
	goto L523
L522:
	;
	v2025 = int32(0)
	goto L523
L523:
	;
	v2026 = F_create_upper_unique_path(m, v1554, v2020, v2025, v1970)
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L1
	} else {
		goto L524
	}
L524:
	;
	F_add_path(m, v1554, v2026)
	mBase = m.M
	v2029 = m.ExcPending
	if v2029 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	goto L506
L526:
	;
	v2038 = F_create_merge_append_path(m, l1, v1554, v1546, v380)
	mBase = m.M
	v2039 = m.ExcPending
	if v2039 != 0 {
		goto L1
	} else {
		goto L527
	}
L527:
	;
	if v210 != 0 {
		goto L528
	} else {
		goto L529
	}
L528:
	;
	v2040 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v2042 = v2040
	goto L530
L529:
	;
	v2042 = int32(0)
	goto L530
L530:
	;
	v2043 = F_create_upper_unique_path(m, v1554, v2038, v2042, v1970)
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		goto L1
	} else {
		goto L531
	}
L531:
	;
	F_add_path(m, v1554, v2043)
	mBase = m.M
	v2046 = m.ExcPending
	if v2046 != 0 {
		goto L1
	} else {
		goto L532
	}
L532:
	;
	v2056 = v1554
	goto L4
L533:
	;
	v2421 = *(*int32)(unsafe.Add(mBase, _c_F_recurse_set_operations[2]))
	if v2421 != 0 {
		goto L579
	} else {
		goto L580
	}
L534:
	;
	if v2079 != 0 {
		goto L535
	} else {
		goto L536
	}
L535:
	;
	v2081 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if l4 != 0 {
		goto L539
	} else {
		goto L540
	}
L536:
	;
	goto L537
L537:
	;
	v2244 = int32(0)
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v2249 = F_generate_setop_tlist(m, l3, l4, v2244, v2244, v2246, l5, v33+int32(56))
	mBase = m.M
	v2250 = m.ExcPending
	if v2250 != 0 {
		goto L1
	} else {
		goto L560
	}
L538:
	;
	if v2213 != 0 {
		goto L533
	} else {
		goto L559
	}
L539:
	;
	v2083 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v2085 = v2083
	goto L541
L540:
	;
	v2085 = int32(0)
	goto L541
L541:
	;
	if v2081 == int32(0) {
		v2164 = v2085
		goto L542
	} else {
		goto L543
	}
L542:
	;
	v2213 = base.B2i32(v2164 == int32(0))
	goto L538
L543:
	;
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(v2081)+4))
	if v2088 <= int32(0) {
		v2164 = v2085
		goto L542
	} else {
		goto L544
	}
L544:
	;
	v2093 = int32(0)
	v2104 = v2085
	goto L545
L545:
	;
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(v2081)+12))
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(v2121+v2093<<(uint(int32(2))%32))))
	v2126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2125)+26)))
	if v2126 == int32(0) {
		goto L548
	} else {
		goto L549
	}
L546:
	;
	v2164 = v2145
	goto L542
L547:
	;
	v2148 = v2093 + int32(1)
	v2149 = *(*int32)(unsafe.Add(mBase, uint32(v2081)+4))
	if v2148 < v2149 {
		v2093 = v2148
		v2104 = v2145
		goto L545
	} else {
		goto L558
	}
L548:
	;
	if v2104 == int32(0) {
		goto L551
	} else {
		goto L552
	}
L549:
	;
	goto L550
L550:
	;
	v2213 = int32(0)
	goto L538
L551:
	;
	goto L550
L552:
	;
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v2125)+4))
	v2132 = F_exprCollation(m, v2131)
	mBase = m.M
	v2133 = m.ExcPending
	if v2133 != 0 {
		goto L1
	} else {
		goto L553
	}
L553:
	;
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v2104)))
	if v2132 != v2134 {
		goto L551
	} else {
		goto L554
	}
L554:
	;
	v2137 = v2104 + int32(4)
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if base.Ui32(v2137) < base.Ui32(v2139+v2140<<(uint(int32(2))%32)) {
		goto L555
	} else {
		goto L556
	}
L555:
	;
	v2145 = v2137
	goto L557
L556:
	;
	v2145 = int32(0)
	goto L557
L557:
	;
	goto L547
L558:
	;
	goto L546
L559:
	;
	goto L537
L560:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v2249
	v2252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+56)))
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v2252)
	v2254 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v2255 = F_make_pathtarget_from_tlist(m, v2254)
	mBase = m.M
	v2256 = m.ExcPending
	if v2256 != 0 {
		goto L1
	} else {
		goto L561
	}
L561:
	;
	v2257 = F_set_pathtarget_cost_width(m, l1, v2255)
	mBase = m.M
	v2258 = m.ExcPending
	if v2258 != 0 {
		goto L1
	} else {
		goto L562
	}
L562:
	;
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(v2056)+32))
	if v2259 == int32(0) {
		goto L563
	} else {
		goto L564
	}
L563:
	;
	v2340 = *(*int32)(unsafe.Add(mBase, uint32(v2056)+40))
	if v2340 == int32(0) {
		goto L533
	} else {
		goto L573
	}
L564:
	;
	v2262 = *(*int32)(unsafe.Add(mBase, uint32(v2259)+4))
	if v2262 <= int32(0) {
		goto L563
	} else {
		goto L565
	}
L565:
	;
	v2268 = int32(0)
	goto L566
L566:
	;
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v2259)+12))
	v2299 = v2296 + v2268<<(uint(int32(2))%32)
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(v2299)))
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(v2300)+8))
	v2302 = F_apply_projection_to_path(m, l1, v2301, v2300, v2257)
	mBase = m.M
	v2303 = m.ExcPending
	if v2303 != 0 {
		goto L1
	} else {
		goto L568
	}
L567:
	;
	goto L563
L568:
	;
	if v2300 != v2302 {
		goto L569
	} else {
		goto L570
	}
L569:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2299))) = v2302
	goto L571
L570:
	;
	goto L571
L571:
	;
	v2307 = v2268 + int32(1)
	v2308 = *(*int32)(unsafe.Add(mBase, uint32(v2259)+4))
	if v2307 < v2308 {
		v2268 = v2307
		goto L566
	} else {
		goto L572
	}
L572:
	;
	goto L567
L573:
	;
	v2343 = int32(0)
	v2344 = *(*int32)(unsafe.Add(mBase, uint32(v2340)+4))
	if v2344 <= v2343 {
		goto L533
	} else {
		goto L574
	}
L574:
	;
	v2347 = v2343
	goto L575
L575:
	;
	v2377 = *(*int32)(unsafe.Add(mBase, uint32(v2340)+12))
	v2380 = v2377 + v2347<<(uint(int32(2))%32)
	v2381 = *(*int32)(unsafe.Add(mBase, uint32(v2380)))
	v2382 = *(*int32)(unsafe.Add(mBase, uint32(v2381)+8))
	v2383 = F_create_projection_path(m, l1, v2382, v2381, v2257)
	mBase = m.M
	v2384 = m.ExcPending
	if v2384 != 0 {
		goto L1
	} else {
		goto L577
	}
L576:
	;
	goto L533
L577:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2380))) = v2383
	v2387 = v2347 + int32(1)
	v2388 = *(*int32)(unsafe.Add(mBase, uint32(v2340)+4))
	if v2387 < v2388 {
		v2347 = v2387
		goto L575
	} else {
		goto L578
	}
L578:
	;
	goto L576
L579:
	;
	v2422 = int32(0)
	m.T0[v2421].(func(*base.Module, int32, int32, int32, int32, int32))(m, l1, v2422, v2422, v2056, v2422)
	mBase = m.M
	v2426 = m.ExcPending
	if v2426 != 0 {
		goto L1
	} else {
		goto L582
	}
L580:
	;
	goto L581
L581:
	;
	F_set_cheapest(m, v2056)
	mBase = m.M
	v2428 = m.ExcPending
	if v2428 != 0 {
		goto L1
	} else {
		goto L583
	}
L582:
	;
	goto L581
L583:
	;
	v2438 = v2056
	goto L3
}
func F_regconfigin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v12 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L29
	} else {
		goto L44
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v151
L3:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_regconfigin[0]))
	if v119 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v15 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v12-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v151 = int32(0)
	goto L2
L8:
	;
	v23 = int32(_a_F_regconfigin_0)
	v27 = m.G0
	v29 = v27 - int32(32)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v30
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regconfigin[1])))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v107 = F_strlen(m, v11)
	mBase = m.M
	if v106 != v107 {
		goto L3
	} else {
		goto L28
	}
L10:
	;
	v106 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regconfigin[2])))
	if v42 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = v11
	goto L16
L14:
	;
	goto L15
L15:
	;
	v56 = v23
	v57 = v38
	goto L19
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v52 == v38 {
		v46 = v46 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v106 = v46 - v11
	goto L9
L18:
	;
	goto L17
L19:
	;
	v64 = v29 + int32(base.Ui32(v57)>>(uint(int32(3))%32))&int32(28)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v65 | v66<<(uint(v57)%32)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v70 != 0 {
		v56 = v56 + v66
		v57 = v70
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v73 == int32(0) {
		v96 = v11
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v96 - v11
	goto L9
L23:
	;
	v77 = v11
	v78 = v73
	goto L24
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(base.Ui32(v78)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v86)>>(uint(v78)%32))&int32(1) == int32(0) {
		v96 = v77
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v96 = v94
	goto L22
L26:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v94 = v77 + int32(1)
	if v92 != 0 {
		v77 = v94
		v78 = v92
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v113 = F_DirectInputFunctionCallSafe(m, int32(547), v11, int32(-1), v10, v8+int32(12))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int32(0)
L30:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v151 = v117
	goto L2
L31:
	;
	v122 = F_stringToQualifiedNameList(m, v11, v10)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if v122 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v126)
	v151 = int32(0)
	goto L2
L34:
	;
	goto L35
L35:
	;
	v130 = F_get_ts_config_oid(m, v122, int32(1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L29
	} else {
		goto L36
	}
L36:
	;
	if v130 != 0 {
		v151 = v130
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v132 = int32(0)
	v133 = F_errsave_start(m, v10)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L29
	} else {
		goto L38
	}
L38:
	;
	if v133 == int32(0) {
		v151 = v132
		goto L2
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L29
	} else {
		goto L40
	}
L40:
	;
	v140 = F_NameListToString(m, v122)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L29
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v140
	F_errmsg(m, int32(_a_F_regconfigin_1), v8)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L29
	} else {
		goto L42
	}
L42:
	;
	F_errsave_finish(m, v10, int32(_a_F_regconfigin_2), int32(1350), int32(_a_F_regconfigin_3))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	v151 = v132
	goto L2
L44:
	;
	F_errmsg_internal(m, int32(_a_F_regconfigin_4), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L29
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_regconfigin_2), int32(1334), int32(_a_F_regconfigin_3))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L29
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regdictionaryin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v12 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L29
	} else {
		goto L44
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v151
L3:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_regdictionaryin[0]))
	if v119 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v15 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v12-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v151 = int32(0)
	goto L2
L8:
	;
	v23 = int32(_a_F_regdictionaryin_0)
	v27 = m.G0
	v29 = v27 - int32(32)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v30
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regdictionaryin[1])))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v107 = F_strlen(m, v11)
	mBase = m.M
	if v106 != v107 {
		goto L3
	} else {
		goto L28
	}
L10:
	;
	v106 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regdictionaryin[2])))
	if v42 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = v11
	goto L16
L14:
	;
	goto L15
L15:
	;
	v56 = v23
	v57 = v38
	goto L19
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v52 == v38 {
		v46 = v46 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v106 = v46 - v11
	goto L9
L18:
	;
	goto L17
L19:
	;
	v64 = v29 + int32(base.Ui32(v57)>>(uint(int32(3))%32))&int32(28)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v65 | v66<<(uint(v57)%32)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v70 != 0 {
		v56 = v56 + v66
		v57 = v70
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v73 == int32(0) {
		v96 = v11
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v96 - v11
	goto L9
L23:
	;
	v77 = v11
	v78 = v73
	goto L24
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(base.Ui32(v78)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v86)>>(uint(v78)%32))&int32(1) == int32(0) {
		v96 = v77
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v96 = v94
	goto L22
L26:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v94 = v77 + int32(1)
	if v92 != 0 {
		v77 = v94
		v78 = v92
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v113 = F_DirectInputFunctionCallSafe(m, int32(547), v11, int32(-1), v10, v8+int32(12))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int32(0)
L30:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v151 = v117
	goto L2
L31:
	;
	v122 = F_stringToQualifiedNameList(m, v11, v10)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if v122 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v126)
	v151 = int32(0)
	goto L2
L34:
	;
	goto L35
L35:
	;
	v130 = F_get_ts_dict_oid(m, v122, int32(1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L29
	} else {
		goto L36
	}
L36:
	;
	if v130 != 0 {
		v151 = v130
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v132 = int32(0)
	v133 = F_errsave_start(m, v10)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L29
	} else {
		goto L38
	}
L38:
	;
	if v133 == int32(0) {
		v151 = v132
		goto L2
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L29
	} else {
		goto L40
	}
L40:
	;
	v140 = F_NameListToString(m, v122)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L29
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v140
	F_errmsg(m, int32(_a_F_regdictionaryin_1), v8)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L29
	} else {
		goto L42
	}
L42:
	;
	F_errsave_finish(m, v10, int32(_a_F_regdictionaryin_2), int32(1460), int32(_a_F_regdictionaryin_3))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	v151 = v132
	goto L2
L44:
	;
	F_errmsg_internal(m, int32(_a_F_regdictionaryin_4), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L29
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_regdictionaryin_2), int32(1444), int32(_a_F_regdictionaryin_3))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L29
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regex_fixed_prefix(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v164 int32
	_ = v164
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v411 int32
	_ = v411
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v478 int32
	_ = v478
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v598 int32
	_ = v598
	var v601 float64
	_ = v601
	var v602 int32
	_ = v602
	var v604 float64
	_ = v604
	var v605 int32
	_ = v605
	var v609 float64
	_ = v609
	var v617 float64
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v649 int32
	_ = v649
	var v652 float64
	_ = v652
	var v653 int32
	_ = v653
	var v655 float64
	_ = v655
	var v656 int32
	_ = v656
	var v660 float64
	_ = v660
	var v665 float64
	_ = v665
	var v669 float64
	_ = v669
	var v670 float64
	_ = v670
	var v672 float64
	_ = v672
	var v680 float64
	_ = v680
	var v683 int32
	_ = v683
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	v26 = m.G0
	v28 = v26 - int32(16)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v30 != int32(17) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v34 = F_pg_detoast_datum_packed(m, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L4
	} else {
		goto L147
	}
L4:
	;
	return int32(0)
L5:
	;
	v38 = m.G0
	v40 = v38 - int32(128)
	m.G0 = v40
	v43 = v28 + int32(15)
	v44 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43))) = uint8(v44)
	if l1 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v48 = int32(27)
	goto L8
L7:
	;
	v48 = int32(19)
	goto L8
L8:
	;
	v49 = F_RE_compile_and_cache(m, v34, v48, l2)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v51 = int32(16)
	v53 = v40 + int32(124)
	v54 = int32(0)
	v57 = v40 + int32(120)
	if base.B2i32(v53 == v54)|base.B2i32(v57 == v54) != 0 {
		v478 = v51
		goto L15
	} else {
		goto L16
	}
L10:
	;
	m.G0 = v40 + int32(128)
	if v569 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L11:
	;
	v548 = *(*int32)(unsafe.Add(mBase, _c_F_regex_fixed_prefix[0]))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)+4))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v549*int32(28))+uint32(_c_F_regex_fixed_prefix[1])))
	goto L92
L12:
	;
	v545 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43))) = uint8(v545)
	goto L11
L13:
	;
	v526 = v40 + int32(16)
	F_pg_regerror(m, v522, v526)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L4
	} else {
		goto L87
	}
L14:
	;
	switch v522 + int32(2) {
	case 0:
		goto L12
	case 1:
		goto L11
	default:
		goto L13
	case 3:
		v569 = int32(0)
		goto L10
	}
L15:
	;
	v522 = v478
	goto L14
L16:
	;
	v61 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v61
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_regex_fixed_prefix[2]))
	if v66 != int32(_a_F_regex_fixed_prefix_0) {
		v478 = v51
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_regex_fixed_prefix[3]))
	if v71 != int32(4) {
		v522 = int32(17)
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_regex_fixed_prefix[4]))
	F_pg_set_regex_collation(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v78 = int32(1)
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_regex_fixed_prefix[5]))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+9)))
	if v81&int32(16) != 0 {
		v478 = v78
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+44)))
	if v85&int32(2) != 0 {
		v478 = v78
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v89 = v84 + int32(36)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v91 = int32(2)
	v94 = F_palloc_extended(m, v90<<(uint(v91)%32), v91)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v94
	if v94 == int32(0) {
		v522 = int32(12)
		goto L14
	} else {
		goto L23
	}
L23:
	;
	v101 = v80 + int32(72)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v102+v103<<(uint(int32(2))%32))))
	v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107))))
	if v108 == int32(_a_F_regex_fixed_prefix_1) {
		v444 = v78
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if base.Ui32(int32(-3)) < base.Ui32(v444) {
		v478 = v444
		goto L15
	} else {
		goto L85
	}
L25:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+20)))
	v114 = v108
	v118 = v107
	v122 = int32(-1)
	goto L26
L26:
	;
	v139 = v114 & int32(_a_F_regex_fixed_prefix_1)
	if v111 != v139 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v147 == int32(-1) {
		v444 = v78
		goto L24
	} else {
		goto L38
	}
L28:
	;
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+22)))
	if v139 != v141 {
		v444 = v78
		goto L24
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	if v122 == int32(-1) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L30
L32:
	;
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v118)+8)))
	if v148 != int32(_a_F_regex_fixed_prefix_1) {
		v114 = v148
		v118 = v118 + int32(8)
		v122 = v147
		goto L26
	} else {
		goto L37
	}
L33:
	;
	v147 = v143
	goto L32
L34:
	;
	goto L35
L35:
	;
	if v143 != v122 {
		v444 = v78
		goto L24
	} else {
		goto L36
	}
L36:
	;
	v147 = v122
	goto L32
L37:
	;
	goto L27
L38:
	;
	v164 = v147
	goto L40
L39:
	;
	if v315 == int32(_a_F_regex_fixed_prefix_1) {
		goto L67
	} else {
		goto L68
	}
L40:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v180+v164<<(uint(int32(2))%32))))
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v184))))
	if v185 == int32(_a_F_regex_fixed_prefix_1) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v307+v164<<(uint(int32(2))%32))))
	v312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v311))))
	v315 = v312
	v319 = v311
	goto L39
L42:
	;
	goto L41
L43:
	;
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+20)))
	v192 = v185
	v196 = v184
	v200 = int32(-1)
	v201 = int32(_a_F_regex_fixed_prefix_1)
	goto L44
L44:
	;
	v217 = v192 & int32(_a_F_regex_fixed_prefix_1)
	if v217 == v190 {
		v239 = v200
		v240 = v201
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v246 = int32(_a_F_regex_fixed_prefix_1)
	v247 = v240 & v246
	if v247 == v246 {
		goto L42
	} else {
		goto L57
	}
L46:
	;
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v196)+8)))
	if v243 != int32(_a_F_regex_fixed_prefix_1) {
		v192 = v243
		v196 = v196 + int32(8)
		v200 = v239
		v201 = v240
		goto L44
	} else {
		goto L56
	}
L47:
	;
	v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+22)))
	if v217 == v219 {
		v239 = v200
		v240 = v201
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+24)))
	if base.B2i32(v217 == v221)|base.B2i32(v217 == int32(_a_F_regex_fixed_prefix_2)) != 0 {
		v315 = v185
		v319 = v184
		goto L39
	} else {
		goto L49
	}
L49:
	;
	v226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+26)))
	if v217 == v226 {
		v315 = v185
		v319 = v184
		goto L39
	} else {
		goto L50
	}
L50:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v228 <= base.I32_extend16_s(v192) {
		v315 = v185
		v319 = v184
		goto L39
	} else {
		goto L51
	}
L51:
	;
	v231 = int32(_a_F_regex_fixed_prefix_1)
	v232 = v201 & v231
	if v232 == v231 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	v239 = v235
	v240 = v192
	goto L46
L53:
	;
	goto L54
L54:
	;
	if v232 != v217 {
		v315 = v185
		v319 = v184
		goto L39
	} else {
		goto L55
	}
L55:
	;
	v239 = int32(-1)
	v240 = v192
	goto L46
L56:
	;
	goto L45
L57:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
	v254 = v250 + base.I32_extend16_s(v240)*int32(24)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	if v255 != int32(1) {
		goto L42
	} else {
		goto L58
	}
L58:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	if v258 != 0 {
		goto L42
	} else {
		goto L59
	}
L59:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v254)+16))
	if base.Ui32(v259) <= base.Ui32(int32(2047)) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v268&int32(_a_F_regex_fixed_prefix_1) != v247 {
		goto L42
	} else {
		goto L64
	}
L61:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v101)+24))
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262+v259<<(uint(int32(1))%32)))))
	v268 = v266
	goto L60
L62:
	;
	goto L63
L63:
	;
	v267 = F_pg_reg_getcolor(m, v101, v259)
	mBase = m.M
	v268 = v267
	goto L60
L64:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v272 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v94+v272<<(uint(int32(2))%32)))) = v259
	if v239 != int32(-1) {
		v164 = v239
		goto L40
	} else {
		goto L65
	}
L65:
	;
	goto L42
L66:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	if v411 == v432 {
		v444 = int32(-2)
		goto L24
	} else {
		goto L81
	}
L67:
	;
	v411 = int32(-1)
	goto L66
L68:
	;
	v340 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+24)))
	v344 = v315
	v347 = int32(-1)
	v348 = v319
	goto L69
L69:
	;
	if v344 != v340 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v411 = v374
	goto L66
L71:
	;
	v368 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+26)))
	if v344 != v368 {
		goto L67
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	if v347 == int32(-1) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L73
L75:
	;
	v375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v348)+8)))
	if v375 != int32(_a_F_regex_fixed_prefix_1) {
		v344 = v375
		v347 = v374
		v348 = v348 + int32(8)
		goto L69
	} else {
		goto L80
	}
L76:
	;
	v374 = v370
	goto L75
L77:
	;
	goto L78
L78:
	;
	if v370 != v347 {
		goto L67
	} else {
		goto L79
	}
L79:
	;
	v374 = v347
	goto L75
L80:
	;
	goto L70
L81:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v436 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v437 = int32(-1)
	goto L84
L83:
	;
	v437 = int32(1)
	goto L84
L84:
	;
	v444 = v437
	goto L24
L85:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	F_pfree(m, v465)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	v468 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v468
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v468
	v478 = v444
	goto L15
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v526
	F_errmsg(m, int32(_a_F_regex_fixed_prefix_3), v40)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_regex_fixed_prefix_4), int32(2068), int32(_a_F_regex_fixed_prefix_5))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v40)+120))
	v559 = F_palloc(m, v554*v555+int32(1))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v40)+124))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v40)+120))
	v563 = F_pg_wchar2mb_with_len(m, v561, v559, v562)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+120)) = v563
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v40)+124))
	F_pfree(m, v566)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	v569 = v559
	goto L10
L96:
	;
	m.G0 = v28 + int32(16)
	return v698
L97:
	;
	v575 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v575
	if l4 == v575 {
		v698 = v575
		goto L96
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v621 = F_string_to_const(m, v569, v30)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L4
	} else {
		goto L116
	}
L100:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v582 = F_text_to_cstring(m, v581)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L4
	} else {
		goto L104
	}
L101:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = v617
	F_pfree(m, v582)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L4
	} else {
		goto L115
	}
L102:
	;
	if base.F64_lt(v609, float64(0)) != 0 {
		v617 = float64(0)
		goto L101
	} else {
		goto L113
	}
L103:
	;
	v604 = F_regex_selectivity_sub(m, v582, v584)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L4
	} else {
		goto L112
	}
L104:
	;
	v584 = F_strlen(m, v582)
	mBase = m.M
	if v584 <= int32(0) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v588 = v584 - int32(1)
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582+v588))))
	if v590 != int32(36) {
		goto L103
	} else {
		goto L106
	}
L106:
	;
	if v584 != int32(1) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v584+v582-int32(2)))))
	if v598 == int32(92) {
		goto L103
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v601 = F_regex_selectivity_sub(m, v582, v588)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L4
	} else {
		goto L111
	}
L110:
	;
	goto L109
L111:
	;
	v609 = v601
	goto L102
L112:
	;
	v609 = base.F64_mul(v604, float64(5))
	goto L102
L113:
	;
	if base.F64_gt(v609, float64(1)) == int32(0) {
		v617 = v609
		goto L101
	} else {
		goto L114
	}
L114:
	;
	v617 = float64(1)
	goto L101
L115:
	;
	v698 = v575
	goto L96
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v621
	if l4 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	F_pfree(m, v569)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L4
	} else {
		goto L143
	}
L118:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+15)))
	if v626 == int32(1) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(4607182418800017408)
	goto L117
L120:
	;
	goto L121
L121:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v632 = F_text_to_cstring(m, v631)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	v634 = F_strlen(m, v632)
	mBase = m.M
	v635 = F_strlen(m, v569)
	mBase = m.M
	if v634 <= int32(0) {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	if int32(0) < v635 {
		goto L134
	} else {
		goto L135
	}
L124:
	;
	v655 = F_regex_selectivity_sub(m, v632, v634)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L4
	} else {
		goto L132
	}
L125:
	;
	v639 = v634 - int32(1)
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632+v639))))
	if v641 != int32(36) {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	if v634 != int32(1) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634+v632-int32(2)))))
	if v649 == int32(92) {
		goto L124
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v652 = F_regex_selectivity_sub(m, v632, v639)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L4
	} else {
		goto L131
	}
L130:
	;
	goto L129
L131:
	;
	v660 = v652
	goto L123
L132:
	;
	v660 = base.F64_mul(v655, float64(5))
	goto L123
L133:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = v680
	F_pfree(m, v632)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L4
	} else {
		goto L142
	}
L134:
	;
	v665 = F_pow(m, float64(0.2), base.F64_convert_i32_u(v635))
	mBase = m.M
	if base.F64_gt(v665, float64(0)) != 0 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v670 = v660
	goto L136
L136:
	;
	v672 = float64(0)
	if base.F64_lt(v670, v672) != 0 {
		v680 = v672
		goto L133
	} else {
		goto L140
	}
L137:
	;
	v669 = base.F64_div(v660, v665)
	goto L139
L138:
	;
	v669 = v660
	goto L139
L139:
	;
	v670 = v669
	goto L136
L140:
	;
	if base.F64_gt(v670, float64(1)) == int32(0) {
		v680 = v670
		goto L133
	} else {
		goto L141
	}
L141:
	;
	v680 = float64(1)
	goto L133
L142:
	;
	goto L117
L143:
	;
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+15)))
	if v694 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v695 = int32(2)
	goto L146
L145:
	;
	v695 = int32(1)
	goto L146
L146:
	;
	v698 = v695
	goto L96
L147:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	F_errmsg(m, int32(_a_F_regex_fixed_prefix_6), int32(0))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_regex_fixed_prefix_7), int32(1106), int32(_a_F_regex_fixed_prefix_8))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_register_partpruneinfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	v4 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+l1<<(uint(int32(2))%32))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l2 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v168
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v180 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L2:
	;
	v168 = v25
	goto L1
L3:
	;
	goto L4
L4:
	;
	if v25 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v85 < int32(0) {
		v168 = v4
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v85 = base.I32_ctz(v71) | v72<<(uint(int32(5))%32)
	goto L5
L7:
	;
	v85 = int32(-2)
	goto L5
L8:
	;
	v38 = base.I32_div_s(int32(0), int32(32))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v39 <= v38 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v42 = v25 + int32(8)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v38<<(uint(int32(2))%32))))
	v49 = v46 & int32(-1)
	if v49 != 0 {
		v71 = v49
		v72 = v38
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v51 = v38 + int32(1)
	if v51 == v39 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v54 = v51
	goto L12
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v42+v54<<(uint(int32(2))%32))))
	if v61 != 0 {
		v71 = v61
		v72 = v54
		goto L6
	} else {
		goto L14
	}
L13:
	;
	goto L7
L14:
	;
	v63 = v54 + int32(1)
	if v63 != v39 {
		v54 = v63
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v89 = v85
	v91 = v4
	goto L17
L17:
	;
	v103 = F_bms_add_member(m, v91, v89+l2)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v168 = v103
	goto L1
L19:
	;
	return int32(0)
L20:
	;
	if v25 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	if int32(0) <= v162 {
		v89 = v162
		v91 = v103
		goto L17
	} else {
		goto L32
	}
L22:
	;
	v162 = base.I32_ctz(v148) | v149<<(uint(int32(5))%32)
	goto L21
L23:
	;
	v162 = int32(-2)
	goto L21
L24:
	;
	v113 = v89 + int32(1)
	v115 = base.I32_div_s(v113, int32(32))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v116 <= v115 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v119 = v25 + int32(8)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119+v115<<(uint(int32(2))%32))))
	v126 = v123 & (int32(-1) << (uint(v113) % 32))
	if v126 != 0 {
		v148 = v126
		v149 = v115
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v128 = v115 + int32(1)
	if v128 == v116 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v131 = v128
	goto L28
L28:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v119+v131<<(uint(int32(2))%32))))
	if v138 != 0 {
		v148 = v138
		v149 = v131
		goto L22
	} else {
		goto L30
	}
L29:
	;
	goto L23
L30:
	;
	v140 = v131 + int32(1)
	if v140 != v116 {
		v131 = v140
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	goto L18
L33:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	v376 = F_lappend(m, v375, v24)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L19
	} else {
		goto L82
	}
L34:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v183 <= int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v198 = v4
	goto L36
L36:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v200+v198<<(uint(int32(2))%32))))
	if v204 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L33
L38:
	;
	v358 = v198 + int32(1)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v358 < v359 {
		v198 = v358
		goto L36
	} else {
		goto L81
	}
L39:
	;
	v207 = int32(0)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v208 <= v207 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v222 = v207
	goto L41
L41:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225+v222<<(uint(int32(2))%32))))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v229)+4)) = v230 + l2
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v229)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = int64(4607182418800017408)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
	if l2 != 0 {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	goto L38
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229)+36)) = v282
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v229)+12))
	if int32(0) < v284 {
		goto L70
	} else {
		goto L71
	}
L44:
	;
	if v263 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L45:
	;
	v272 = F_fix_scan_expr_mutator(m, v271, v17)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L19
	} else {
		goto L64
	}
L46:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v264 != 0 {
		v271 = v263
		goto L45
	} else {
		goto L60
	}
L47:
	;
	v255 = F_fix_scan_expr_mutator(m, v233, v17)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L19
	} else {
		goto L58
	}
L48:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v238 != 0 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+68))
	if v240 != 0 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	if v241 != 0 {
		goto L47
	} else {
		goto L51
	}
L51:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)))
	if v242 != 0 {
		goto L47
	} else {
		goto L52
	}
L52:
	;
	if v233 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_fix_expr_common(m, l0, v233)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L19
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229)+32)) = v233
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v229)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = int64(4607182418800017408)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
	v263 = v249
	goto L46
L56:
	;
	v246 = F_expression_tree_walker_impl(m, v233, int32(840), v17)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L19
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229)+32)) = v255
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v229)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = int64(4607182418800017408)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
	if l2 != 0 {
		v271 = v258
		goto L45
	} else {
		goto L59
	}
L59:
	;
	v263 = v258
	goto L46
L60:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+68))
	if v266 != 0 {
		v271 = v263
		goto L45
	} else {
		goto L61
	}
L61:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	if v267 != 0 {
		v271 = v263
		goto L45
	} else {
		goto L62
	}
L62:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)))
	if v268 != int32(1) {
		goto L44
	} else {
		goto L63
	}
L63:
	;
	v271 = v263
	goto L45
L64:
	;
	v282 = v272
	goto L43
L65:
	;
	v282 = int32(0)
	goto L43
L66:
	;
	goto L67
L67:
	;
	F_fix_expr_common(m, l0, v263)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L19
	} else {
		goto L68
	}
L68:
	;
	v280 = F_expression_tree_walker_impl(m, v263, int32(840), v17)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L19
	} else {
		goto L69
	}
L69:
	;
	v282 = v263
	goto L43
L70:
	;
	v291 = int32(0)
	goto L73
L71:
	;
	goto L72
L72:
	;
	v340 = v222 + int32(1)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v340 < v341 {
		v222 = v340
		goto L41
	} else {
		goto L80
	}
L73:
	;
	v303 = v291 << (uint(int32(2)) % 32)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v229)+24))
	v305 = v303 + v304
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	if v306 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L72
L75:
	;
	v322 = v291 + int32(1)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v229)+12))
	if v322 < v323 {
		v291 = v322
		goto L73
	} else {
		goto L79
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v305))) = v306 + l2
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v229)+32))
	if v311 == int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v229)+24))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v315+v303)))
	v318 = F_bms_add_member(m, v314, v317)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L19
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v318
	goto L75
L79:
	;
	goto L74
L80:
	;
	goto L42
L81:
	;
	goto L37
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v376
	if v376 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	v383 = v379 - int32(1)
	goto L85
L84:
	;
	v383 = int32(-1)
	goto L85
L85:
	;
	m.G0 = v17 + int32(16)
	return v383
}
func F_regrolein(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
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
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v12 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L29
	} else {
		goto L51
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v175
L3:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_regrolein[0]))
	if v119 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v15 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v12-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v175 = int32(0)
	goto L2
L8:
	;
	v23 = int32(_a_F_regrolein_0)
	v27 = m.G0
	v29 = v27 - int32(32)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v30
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regrolein[1])))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v107 = F_strlen(m, v11)
	mBase = m.M
	if v106 != v107 {
		goto L3
	} else {
		goto L28
	}
L10:
	;
	v106 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regrolein[2])))
	if v42 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = v11
	goto L16
L14:
	;
	goto L15
L15:
	;
	v56 = v23
	v57 = v38
	goto L19
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v52 == v38 {
		v46 = v46 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v106 = v46 - v11
	goto L9
L18:
	;
	goto L17
L19:
	;
	v64 = v29 + int32(base.Ui32(v57)>>(uint(int32(3))%32))&int32(28)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v65 | v66<<(uint(v57)%32)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v70 != 0 {
		v56 = v56 + v66
		v57 = v70
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v73 == int32(0) {
		v96 = v11
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v96 - v11
	goto L9
L23:
	;
	v77 = v11
	v78 = v73
	goto L24
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(base.Ui32(v78)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v86)>>(uint(v78)%32))&int32(1) == int32(0) {
		v96 = v77
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v96 = v94
	goto L22
L26:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v94 = v77 + int32(1)
	if v92 != 0 {
		v77 = v94
		v78 = v92
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v113 = F_DirectInputFunctionCallSafe(m, int32(547), v11, int32(-1), v10, v8+int32(12))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int32(0)
L30:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v175 = v117
	goto L2
L31:
	;
	v122 = F_stringToQualifiedNameList(m, v11, v10)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if v122 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v126)
	v175 = int32(0)
	goto L2
L34:
	;
	goto L35
L35:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v129 != int32(1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v132 = int32(0)
	v133 = F_errsave_start(m, v10)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L29
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v153 = F_get_role_oid(m, v151, int32(1))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L29
	} else {
		goto L44
	}
L39:
	;
	if v133 == int32(0) {
		v175 = v132
		goto L2
	} else {
		goto L40
	}
L40:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L29
	} else {
		goto L41
	}
L41:
	;
	F_errmsg(m, int32(_a_F_regrolein_1), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L29
	} else {
		goto L42
	}
L42:
	;
	F_errsave_finish(m, v10, int32(_a_F_regrolein_2), int32(1564), int32(_a_F_regrolein_3))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	v175 = v132
	goto L2
L44:
	;
	if v153 != 0 {
		v175 = v153
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v155 = int32(0)
	v156 = F_errsave_start(m, v10)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L29
	} else {
		goto L46
	}
L46:
	;
	if v156 == int32(0) {
		v175 = v155
		goto L2
	} else {
		goto L47
	}
L47:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L29
	} else {
		goto L48
	}
L48:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v165
	F_errmsg(m, int32(_a_F_regrolein_4), v8)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	F_errsave_finish(m, v10, int32(_a_F_regrolein_2), int32(1572), int32(_a_F_regrolein_3))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	v175 = v155
	goto L2
L51:
	;
	F_errmsg_internal(m, int32(_a_F_regrolein_5), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L29
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_regrolein_2), int32(1554), int32(_a_F_regrolein_3))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L29
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_reject_target_detail(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l1 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg(m, int32(_a_F_reject_target_detail_0), v6)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_reject_target_detail_1), int32(219), int32(_a_F_reject_target_detail_2))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.G0 = v6 + int32(16)
		return int32(0)
	}
}
func F_relabel_to_typmod(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = F_exprType(m, l0)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_exprCollation(m, l0)
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			v12 = F_applyRelabelType(m, l0, v3, l1, v7, int32(1), int32(-1), int32(0))
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v12
			}
		}
	}
}
func F_relatt_cache_syshash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = F_GetSysCacheHashValue(m, int32(7), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_reportDependentObjects(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	v5 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(304)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v5 < v16 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v14 + int32(304)
	return
L2:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v14)+288))
	F_pfree(m, v442)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L33
	} else {
		goto L134
	}
L3:
	;
	if int32(2) <= v257 {
		goto L121
	} else {
		goto L122
	}
L4:
	;
	v349 = F_getObjectDescription(m, l3, int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L33
	} else {
		goto L114
	}
L5:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v317 = F_getObjectDescription(m, v33+int32(4), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L33
	} else {
		goto L107
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = v5
	goto L9
L7:
	;
	goto L8
L8:
	;
	if l2&int32(4) != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v33 = v19 + v28<<(uint(int32(4))%32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v34&int32(144) == int32(128) {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v40 = v28 + int32(1)
	if v40 != v16 {
		v28 = v40
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v57 = int32(13)
	goto L15
L14:
	;
	v57 = int32(18)
	goto L15
L15:
	;
	if l1 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v62 = int32(1)
	if int32(20) < v57 {
		v96 = v62
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L18
L18:
	;
	F_initStringInfo(m, v14+int32(288))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L33
	} else {
		goto L34
	}
L19:
	;
	if v96 == int32(0) {
		goto L1
	} else {
		goto L32
	}
L20:
	;
	goto L19
L21:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_reportDependentObjects[0]))
	if base.Ui32(v57-int32(15)) <= base.Ui32(int32(1)) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v79 = int32(0)
	if v57 == int32(16) {
		v96 = v79
		goto L20
	} else {
		goto L29
	}
L23:
	;
	if int32(22) <= v66 {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if base.B2i32(v57 == int32(20))|base.B2i32(v66 == int32(15)) != 0 {
		goto L22
	} else {
		goto L27
	}
L26:
	;
	v96 = v62
	goto L20
L27:
	;
	if v66 <= v57 {
		v96 = v62
		goto L20
	} else {
		goto L28
	}
L28:
	;
	goto L22
L29:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_reportDependentObjects[1]))
	if v83 != int32(2) {
		v96 = v79
		goto L20
	} else {
		goto L30
	}
L30:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_reportDependentObjects[2])))
	if v87&int32(1) != 0 {
		v96 = v79
		goto L20
	} else {
		goto L31
	}
L31:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_reportDependentObjects[3]))
	v96 = base.B2i32(v57 == int32(17)) | base.B2i32(v93 <= v57)
	goto L20
L32:
	;
	goto L18
L33:
	;
	return
L34:
	;
	F_initStringInfo(m, v14+int32(272))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v110 = v108 - int32(1)
	if v110 < int32(0) {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v114 = int32(0)
	v118 = v114
	v121 = v110
	v122 = v114
	v123 = int32(1)
	goto L37
L37:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v130 = v127 + v121<<(uint(int32(4))%32)
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130))))
	if v131&int32(257) != 0 {
		v255 = v118
		v257 = v122
		v258 = v123
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if int32(0) < v255 {
		goto L90
	} else {
		goto L91
	}
L39:
	;
	if int32(0) < v121 {
		v118 = v255
		v121 = v121 - int32(1)
		v122 = v257
		v123 = v258
		goto L37
	} else {
		goto L88
	}
L40:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v139 = F_getObjectDescription(m, v134+v121*int32(12), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L33
	} else {
		goto L41
	}
L41:
	;
	if v139 == int32(0) {
		v255 = v118
		v257 = v122
		v258 = v123
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v143&int32(60) != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	F_pfree(m, v139)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L33
	} else {
		goto L87
	}
L44:
	;
	v148 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L33
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if l1 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L47:
	;
	if v148 == int32(0) {
		v250 = v118
		v251 = v122
		v252 = v123
		goto L43
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+256)) = v139
	F_errmsg_internal(m, int32(_a_F_reportDependentObjects_0), v14+int32(256))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L33
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_reportDependentObjects_1), int32(1083), int32(_a_F_reportDependentObjects_2))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L33
	} else {
		goto L50
	}
L50:
	;
	v250 = v118
	v251 = v122
	v252 = v123
	goto L43
L51:
	;
	v168 = F_getObjectDescription(m, v130+int32(4), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L33
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v122 <= int32(99) {
		goto L74
	} else {
		goto L75
	}
L54:
	;
	if v168 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	if v122 <= int32(99) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	goto L57
L57:
	;
	v250 = v118 + int32(1)
	v251 = v122
	v252 = int32(0)
	goto L43
L58:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v14)+276))
	if v193 != 0 {
		goto L67
	} else {
		goto L68
	}
L59:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v14)+292))
	if v172 != 0 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L61
L61:
	;
	v191 = v118 + int32(1)
	v192 = v122
	goto L58
L62:
	;
	F_appendStringInfoChar(m, v14+int32(288), int32(10))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L33
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+212)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = v139
	F_appendStringInfo(m, v14+int32(288), int32(_a_F_reportDependentObjects_3), v14+int32(208))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L33
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	v191 = v118
	v192 = v122 + int32(1)
	goto L58
L67:
	;
	F_appendStringInfoChar(m, v14+int32(272), int32(10))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L33
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+196)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v14)+192)) = v139
	F_appendStringInfo(m, v14+int32(272), int32(_a_F_reportDependentObjects_3), v14+int32(192))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L33
	} else {
		goto L71
	}
L70:
	;
	goto L69
L71:
	;
	F_pfree(m, v168)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L33
	} else {
		goto L72
	}
L72:
	;
	v250 = v191
	v251 = v192
	v252 = int32(0)
	goto L43
L73:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v14)+276))
	if v236 != 0 {
		goto L82
	} else {
		goto L83
	}
L74:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v14)+292))
	if v216 != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v234 = v118 + int32(1)
	v235 = v122
	goto L73
L77:
	;
	F_appendStringInfoChar(m, v14+int32(288), int32(10))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L33
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v139
	F_appendStringInfo(m, v14+int32(288), int32(_a_F_reportDependentObjects_4), v14+int32(240))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L33
	} else {
		goto L81
	}
L80:
	;
	goto L79
L81:
	;
	v234 = v118
	v235 = v122 + int32(1)
	goto L73
L82:
	;
	F_appendStringInfoChar(m, v14+int32(272), int32(10))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L33
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+224)) = v139
	F_appendStringInfo(m, v14+int32(272), int32(_a_F_reportDependentObjects_4), v14+int32(224))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L33
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	v250 = v234
	v251 = v235
	v252 = v123
	goto L43
L87:
	;
	v255 = v250
	v257 = v251
	v258 = v252
	goto L39
L88:
	;
	goto L38
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L33
	} else {
		goto L99
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+176)) = v255
	if v255 == int32(1) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	if v258 != 0 {
		goto L3
	} else {
		goto L98
	}
L93:
	;
	v272 = int32(_a_F_reportDependentObjects_5)
	goto L95
L94:
	;
	v272 = int32(_a_F_reportDependentObjects_6)
	goto L95
L95:
	;
	F_appendStringInfo(m, v14+int32(288), v272, v14+int32(176))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L33
	} else {
		goto L96
	}
L96:
	;
	if v258 == int32(0) {
		goto L89
	} else {
		goto L97
	}
L97:
	;
	goto L3
L98:
	;
	goto L89
L99:
	;
	F_errcode(m, int32(16909442))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L33
	} else {
		goto L100
	}
L100:
	;
	if l3 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	F_errmsg(m, int32(_a_F_reportDependentObjects_7), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L33
	} else {
		goto L102
	}
L102:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v14)+288))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v290
	F_errdetail_internal(m, int32(_a_F_reportDependentObjects_8), v14+int32(112))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L33
	} else {
		goto L103
	}
L103:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v14)+272))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v297
	F_errdetail_log(m, int32(_a_F_reportDependentObjects_8), v14+int32(96))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L33
	} else {
		goto L104
	}
L104:
	;
	F_errhint(m, int32(_a_F_reportDependentObjects_9), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L33
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_reportDependentObjects_1), int32(1161), int32(_a_F_reportDependentObjects_2))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L33
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L33
	} else {
		goto L108
	}
L108:
	;
	F_errcode(m, int32(16909442))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L33
	} else {
		goto L109
	}
L109:
	;
	v330 = F_getObjectDescription(m, v313+v28*int32(12), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L33
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v317
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v330
	F_errmsg(m, int32(_a_F_reportDependentObjects_10), v14+int32(16))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L33
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v317
	F_errhint(m, int32(_a_F_reportDependentObjects_11), v14)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L33
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_reportDependentObjects_1), int32(1018), int32(_a_F_reportDependentObjects_2))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L33
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = v349
	F_errmsg(m, int32(_a_F_reportDependentObjects_12), v14+int32(160))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L33
	} else {
		goto L115
	}
L115:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v14)+288))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = v357
	F_errdetail_internal(m, int32(_a_F_reportDependentObjects_8), v14+int32(144))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L33
	} else {
		goto L116
	}
L116:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v14)+272))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v364
	F_errdetail_log(m, int32(_a_F_reportDependentObjects_8), v14+int32(128))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L33
	} else {
		goto L117
	}
L117:
	;
	F_errhint(m, int32(_a_F_reportDependentObjects_9), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L33
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_reportDependentObjects_1), int32(1154), int32(_a_F_reportDependentObjects_2))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L33
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
	F_errfinish(m, int32(_a_F_reportDependentObjects_1), v427, int32(_a_F_reportDependentObjects_2))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L33
	} else {
		goto L133
	}
L121:
	;
	v384 = F_errstart(m, v57, int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L33
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	if v257 != int32(1) {
		goto L2
	} else {
		goto L129
	}
L124:
	;
	if v384 == int32(0) {
		goto L2
	} else {
		goto L125
	}
L125:
	;
	v388 = v255 + v257
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v388
	F_errmsg_plural(m, int32(_a_F_reportDependentObjects_13), int32(_a_F_reportDependentObjects_14), v388, v14-int32(-64))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L33
	} else {
		goto L126
	}
L126:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v14)+288))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v396
	F_errdetail_internal(m, int32(_a_F_reportDependentObjects_8), v14+int32(48))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L33
	} else {
		goto L127
	}
L127:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v14)+272))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v403
	F_errdetail_log(m, int32(_a_F_reportDependentObjects_8), v14+int32(32))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L33
	} else {
		goto L128
	}
L128:
	;
	v427 = int32(1171)
	goto L120
L129:
	;
	v414 = F_errstart(m, v57, int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L33
	} else {
		goto L130
	}
L130:
	;
	if v414 == int32(0) {
		goto L2
	} else {
		goto L131
	}
L131:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v14)+288))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v418
	F_errmsg_internal(m, int32(_a_F_reportDependentObjects_8), v14+int32(80))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L33
	} else {
		goto L132
	}
L132:
	;
	v427 = int32(1177)
	goto L120
L133:
	;
	goto L2
L134:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v14)+272))
	F_pfree(m, v445)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L33
	} else {
		goto L135
	}
L135:
	;
	goto L1
}
func F_reschedule_timeouts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v31 int32
	_ = v31
	v2 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_reschedule_timeouts[0])))
	if v2 == int32(0) {
		return
	} else {
		v6 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_reschedule_timeouts[1])) = v6
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_reschedule_timeouts[2]))
		if v9 <= v6 {
			return
		} else {
			v15 = m.G0
			v16 = int32(16)
			v17 = v15 - v16
			m.G0 = v17
			F_gettimeofday(m, v17)
			mBase = m.M
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
			v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(v17)+8)))
			m.G0 = v17 + v16
			F_schedule_alarm(m, v21+v20*int64(1000000)-int64(946684800000000))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_reservoir_init_selection_state(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v28 int64
	_ = v28
	var v33 int64
	_ = v33
	var v38 int64
	_ = v38
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v72 float64
	_ = v72
	var v75 float64
	_ = v75
	var v79 float64
	_ = v79
	v6 = l0 + int32(8)
	v8 = Fn13964(m, int64(32))
	mBase = m.M
	v9 = base.I64_extend_i32_u(v8)
	v12 = v9 + int64(4354685564936845354)
	v13 = int64(30)
	v16 = int64(-4658895280553007687)
	v17 = (int64(base.Ui64(v12)>>(uint(v13)%64)) ^ v12) * v16
	v18 = int64(27)
	v21 = int64(-7723592293110705685)
	v22 = (int64(base.Ui64(v17)>>(uint(v18)%64)) ^ v17) * v21
	v23 = int64(31)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(base.Ui64(v22)>>(uint(v23)%64)) ^ v22
	v28 = v9 - int64(7046029254386353131)
	v33 = (int64(base.Ui64(v28)>>(uint(v13)%64)) ^ v28) * v16
	v38 = (int64(base.Ui64(v33)>>(uint(v18)%64)) ^ v33) * v21
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(base.Ui64(v38)>>(uint(v23)%64)) ^ v38
	for {
		v49 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
		v50 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
		v51 = v49 ^ v50
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = base.I64_rotl(v51, int64(37))
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v51<<(uint(int64(16))%64) ^ base.I64_rotl(v49, int64(24)) ^ v51
		v72 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v49*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
		mBase = m.M
		if base.F64_eq(v72, float64(0)) != 0 {
			continue
		} else {
			break
		}
		break
	}
	v75 = F_log(m, v72)
	mBase = m.M
	v79 = F_exp(m, base.F64_div(base.F64_neg(v75), base.F64_convert_i32_s(l1)))
	mBase = m.M
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v79
	return
}
func F_resize_intArrayType(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	if l1 <= int32(0) {
		v11 = F_construct_empty_array(m, int32(23))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return v11
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v19 = F_ArrayGetNItemsSafe(m, v16, l0+int32(16))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v19 == l1 {
				v109 = l0
				return v109
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v22 != 0 {
					v30 = v22
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v30 = (v23<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				}
				v33 = v30 + l1<<(uint(int32(2))%32)
				v34 = F_repalloc(m, l0, v33)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v34))) = v33 << (uint(int32(2)) % 32)
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
					if v39 <= int32(0) {
						v109 = v34
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = l1
						if v39 == int32(1) {
							v109 = v34
						} else {
							v46 = v34 + int32(16)
							v47 = int32(1)
							v48 = v39 - v47
							v49 = int32(7)
							v50 = v48 & v49
							if base.Ui32(v49) <= base.Ui32(v39-int32(2)) {
								v60 = v47
								v62 = int32(0)
								for {
									v68 = v46 + v60<<(uint(int32(2))%32)
									v69 = int64(4294967297)
									*(*int64)(unsafe.Add(mBase, uint32(v68)+24)) = v69
									*(*int64)(unsafe.Add(mBase, uint32(v68)+16)) = v69
									*(*int64)(unsafe.Add(mBase, uint32(v68)+8)) = v69
									*(*int64)(unsafe.Add(mBase, uint32(v68))) = v69
									v77 = int32(8)
									v78 = v60 + v77
									v80 = v62 + v77
									if v80 != v48&int32(-8) {
										v60 = v78
										v62 = v80
										continue
									} else {
										break
									}
									break
								}
								if v50 == int32(0) {
									v109 = v34
								} else {
									v85 = v78
									v93 = v85
									v94 = int32(0)
									for {
										v102 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v46+v93<<(uint(int32(2))%32)))) = v102
										v107 = v94 + v102
										if v107 != v50 {
											v93 = v93 + v102
											v94 = v107
											continue
										} else {
											break
										}
										break
									}
									v109 = v34
								}
							} else {
								v85 = v47
								v93 = v85
								v94 = int32(0)
								for {
									v102 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v46+v93<<(uint(int32(2))%32)))) = v102
									v107 = v94 + v102
									if v107 != v50 {
										v93 = v93 + v102
										v94 = v107
										continue
									} else {
										break
									}
									break
								}
								v109 = v34
							}
						}
					}
					return v109
				}
			}
		}
	}
}
func F_resolve_anyarray_from_others(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 == int32(0) {
		F_resolve_anyelement_from_others(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v13 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_resolve_anyarray_from_others_0), int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_resolve_anyarray_from_others_1), int32(674), int32(_a_F_resolve_anyarray_from_others_2))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v16 = v13
				v17 = F_get_array_type(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					if v17 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v46 = F_format_type_be(m, v45)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v46
									F_errmsg(m, int32(_a_F_resolve_anyarray_from_others_3), v6)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_resolve_anyarray_from_others_1), int32(670), int32(_a_F_resolve_anyarray_from_others_2))
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
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v17
						m.G0 = v6 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v16 = v8
		v17 = F_get_array_type(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			if v17 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v46 = F_format_type_be(m, v45)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v46
							F_errmsg(m, int32(_a_F_resolve_anyarray_from_others_3), v6)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_resolve_anyarray_from_others_1), int32(670), int32(_a_F_resolve_anyarray_from_others_2))
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
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v17
				m.G0 = v6 + int32(16)
				return
			}
		}
	}
}
func F_rm_redo_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int64
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v146 int32
	_ = v146
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v194 int32
	_ = v194
	var v195 int64
	_ = v195
	var v197 int32
	_ = v197
	var v200 int64
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	v12 = m.G0
	v14 = v12 - int32(112)
	m.G0 = v14
	F_initStringInfo(m, v14+int32(76))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+49)))
	v27 = v25 << (uint(int32(5)) % 32)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_rm_redo_error_callback[0])))
	if v30 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_RmgrNotFound(m, v25)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v37 = v24
	v38 = v30
	goto L5
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_rm_redo_error_callback[1])))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+48)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_rm_redo_error_callback[2])))
	v43 = v14 + int32(76)
	F_appendStringInfoString(m, v43, v38)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_rm_redo_error_callback[0])))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v37 = v36
	v38 = v35
	goto L5
L7:
	;
	F_appendStringInfoChar(m, v43, int32(47))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v49 = m.T0[v41].(func(*base.Module, int32) int32)(m, v40)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	m.T0[v39].(func(*base.Module, int32, int32))(m, v43, l0)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L16
	}
L10:
	;
	if v49 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v40 & int32(240)
	F_appendStringInfo(m, v43, int32(_a_F_rm_redo_error_callback_0), v22)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v49
	F_appendStringInfo(m, v43, int32(_a_F_rm_redo_error_callback_1), v22+int32(16))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L9
L15:
	;
	goto L9
L16:
	;
	m.G0 = v22 + int32(32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+72))
	if int32(0) <= v71 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v85 = int32(0)
	goto L20
L18:
	;
	goto L19
L19:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L47
	}
L20:
	;
	v90 = v85 & int32(255)
	v92 = v14 + int32(100)
	v94 = v14 + int32(96)
	v96 = v14 + int32(92)
	v97 = int32(0)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+72))
	if v100 < v90 {
		v124 = v97
		goto L24
	} else {
		goto L25
	}
L21:
	;
	goto L19
L22:
	;
	v177 = v85 + int32(1)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+72))
	if v177 <= v179 {
		v85 = v177
		goto L20
	} else {
		goto L46
	}
L23:
	;
	if v124 == int32(0) {
		goto L22
	} else {
		goto L37
	}
L24:
	;
	goto L23
L25:
	;
	v104 = v99 + v90*int32(52)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+76)))
	if v105 != int32(1) {
		v124 = v97
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v109 = v104 + int32(76)
	if v92 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+8)) = v110
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v109)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v92))) = v112
	goto L29
L28:
	;
	goto L29
L29:
	;
	if v94 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v114
	goto L32
L31:
	;
	goto L32
L32:
	;
	if v96 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v109)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v116
	goto L35
L34:
	;
	goto L35
L35:
	;
	v124 = int32(1)
	goto L24
L37:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+100))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	if v131 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159+v85*int32(52))+105)))
	if v163 != int32(1) {
		goto L22
	} else {
		goto L44
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(68)))) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v14-int32(-64)))) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v85
	F_appendStringInfo(m, v14+int32(76), int32(_a_F_rm_redo_error_callback_2), v14+int32(48))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(32)))) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v85
	F_appendStringInfo(m, v14+int32(76), int32(_a_F_rm_redo_error_callback_3), v14+int32(16))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	goto L38
L43:
	;
	goto L38
L44:
	;
	F_appendStringInfoString(m, v14+int32(76), int32(_a_F_rm_redo_error_callback_4))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L22
L46:
	;
	goto L21
L47:
	;
	v195 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+4)) = uint32(v195)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v197
	v200 = int64(base.Ui64(v195) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v14))) = uint32(v200)
	F_errcontext_msg(m, int32(_a_F_rm_redo_error_callback_5), v14)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	F_pfree(m, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	m.G0 = v14 + int32(112)
	return
}
func F_rmtree(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(1104)
	m.G0 = v12
	v14 = F_AllocateDir(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(1104)
	return v231 & int32(1)
L2:
	;
	return int32(0)
L3:
	;
	if v14 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v36 = F_palloc(m, int32(32))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L11
	}
L7:
	;
	if v22 == int32(0) {
		v231 = v2
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg_internal(m, int32(_a_F_rmtree_0), v12)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(_a_F_rmtree_1), int32(63), int32(_a_F_rmtree_2))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v231 = v2
	goto L1
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_rmtree[0])) = int32(0)
	v41 = int32(1)
	v42 = F_readdir(m, v14)
	mBase = m.M
	if v42 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v46 = v42
	v47 = int32(8)
	v48 = v41
	v50 = v36
	v51 = v2
	goto L15
L13:
	;
	v141 = v41
	v143 = v36
	v144 = v2
	goto L14
L14:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_rmtree[0]))
	if v147 == int32(0) {
		v168 = v141
		goto L39
	} else {
		goto L40
	}
L15:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+19)))
	if v53 != int32(46) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v141 = v129
	v143 = v130
	v144 = v131
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_rmtree[0])) = int32(0)
	v136 = F_readdir(m, v14)
	mBase = m.M
	if v136 != 0 {
		v46 = v136
		v47 = v128
		v48 = v129
		v50 = v130
		v51 = v131
		goto L15
	} else {
		goto L38
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v46 + int32(19)
	v70 = v12 + int32(80)
	v75 = F_pg_snprintf(m, v70, int32(1024), int32(_a_F_rmtree_3), v12-int32(-64))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L2
	} else {
		goto L23
	}
L19:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+20)))
	if v56 == int32(0) {
		v128 = v47
		v129 = v48
		v130 = v50
		v131 = v51
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+20)))
	if v59 != int32(46) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+21)))
	if v62 == int32(0) {
		v128 = v47
		v129 = v48
		v130 = v50
		v131 = v51
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	v79 = F_get_dirent_type(m, v70, v46, int32(0), int32(19))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L26
	}
L24:
	;
	v101 = v12 + int32(80)
	v102 = F_unlink(m, v101)
	mBase = m.M
	if v102 == int32(0) {
		v128 = v47
		v129 = v48
		v130 = v50
		v131 = v51
		goto L17
	} else {
		goto L32
	}
L25:
	;
	if v47 == v51 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	switch v79 {
	case 0:
		v128 = v47
		v129 = v48
		v130 = v50
		v131 = v51
		goto L17
	default:
		goto L24
	case 3:
		goto L25
	}
L27:
	;
	v84 = F_repalloc(m, v50, v47<<(uint(int32(3))%32))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L2
	} else {
		goto L30
	}
L28:
	;
	v88 = v47
	v89 = v50
	goto L29
L29:
	;
	v95 = F_pstrdup(m, v12+int32(80))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L31
	}
L30:
	;
	v88 = v47 << (uint(int32(1)) % 32)
	v89 = v84
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89+v51<<(uint(int32(2))%32)))) = v95
	v128 = v88
	v129 = v48
	v130 = v89
	v131 = v51 + int32(1)
	goto L17
L32:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_rmtree[0]))
	if v106 == int32(44) {
		v128 = v47
		v129 = v48
		v130 = v50
		v131 = v51
		goto L17
	} else {
		goto L33
	}
L33:
	;
	v109 = int32(0)
	v112 = F_errstart(m, int32(19), v109)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	if v112 == int32(0) {
		v128 = v47
		v129 = v109
		v130 = v50
		v131 = v51
		goto L17
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v101
	F_errmsg_internal(m, int32(_a_F_rmtree_4), v12+int32(48))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_rmtree_1), int32(97), int32(_a_F_rmtree_2))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v128 = v47
	v129 = v109
	v130 = v50
	v131 = v51
	goto L17
L38:
	;
	goto L16
L39:
	;
	F_FreeDir(m, v14)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L2
	} else {
		goto L45
	}
L40:
	;
	v150 = int32(0)
	v153 = F_errstart(m, int32(19), v150)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	if v153 == int32(0) {
		v168 = v150
		goto L39
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l0
	F_errmsg_internal(m, int32(_a_F_rmtree_5), v12+int32(32))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_rmtree_1), int32(106), int32(_a_F_rmtree_2))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v168 = v150
	goto L39
L45:
	;
	if v144 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v174 = int32(0)
	v176 = v168
	goto L49
L47:
	;
	v198 = v168
	goto L48
L48:
	;
	v203 = F_rmdir(m, l0)
	mBase = m.M
	if v203 == int32(0) {
		v224 = v198
		goto L54
	} else {
		goto L55
	}
L49:
	;
	v183 = v143 + v174<<(uint(int32(2))%32)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v185 = F_rmtree(m, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L2
	} else {
		goto L51
	}
L50:
	;
	v198 = v190
	goto L48
L51:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	F_pfree(m, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v190 = v176 & v185
	v192 = v174 + int32(1)
	if v192 != v144 {
		v174 = v192
		v176 = v190
		goto L49
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	F_pfree(m, v143)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L2
	} else {
		goto L60
	}
L55:
	;
	v206 = int32(0)
	v209 = F_errstart(m, int32(19), v206)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	if v209 == int32(0) {
		v224 = v206
		goto L54
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_rmtree_6), v12+int32(16))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_rmtree_1), int32(124), int32(_a_F_rmtree_2))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v224 = v206
	goto L54
L60:
	;
	v231 = v224
	goto L1
}
func F_roles_list_append(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v116
L2:
	;
	if l0 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L3:
	;
	v15 = F_bloom_lacks_element(m, v11, v8+int32(12), int32(4))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v20 = l2
	goto L5
L5:
	;
	v21 = int32(0)
	if l0 == v21 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	return int32(0)
L7:
	;
	if v15 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v20 = v19
	goto L5
L9:
	;
	if v59 != 0 {
		v116 = l0
		goto L1
	} else {
		goto L22
	}
L10:
	;
	v59 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v27 <= int32(0) {
		v53 = v21
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v59 = v53
	goto L9
L14:
	;
	v30 = int32(0)
	if v30 < v27 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v33 = v27
	goto L17
L16:
	;
	v33 = v30
	goto L17
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v36 = int32(0)
	goto L18
L18:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v34+v36<<(uint(int32(2))%32))))
	v45 = base.B2i32(v44 == v20)
	if v44 == v20 {
		v53 = v45
		goto L13
	} else {
		goto L20
	}
L19:
	;
	v53 = v45
	goto L13
L20:
	;
	v47 = v36 + int32(1)
	if v47 != v33 {
		v36 = v47
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	goto L2
L23:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v106 = F_lappend_oid(m, l0, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L6
	} else {
		goto L33
	}
L24:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v62 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v63 < int32(1025) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_roles_list_append[0]))
	v71 = F_bloom_create(m, int64(10240), v69, int64(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v71
	v74 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v76 <= v74 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v81 = int32(0)
	goto L29
L29:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84+v81<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_bloom_add_element(m, v90, v8+int32(8), int32(4))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L31
	}
L30:
	;
	goto L23
L31:
	;
	v97 = v81 + int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v97 < v98 {
		v81 = v97
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v108 == int32(0) {
		v116 = v106
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_bloom_add_element(m, v108, v8+int32(12), int32(4))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	v116 = v106
	goto L1
}
func F_rpad(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v21 = F_pg_detoast_datum_packed(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if int32(0) < v17 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = v17
	goto L6
L5:
	;
	v24 = int32(0)
	goto L6
L6:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v25 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v55 = int32(0)
	if v55 < v54 {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v31 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v42 = int32(1)
	if v25&v42 != 0 {
		v54 = int32(base.Ui32(v25)>>(uint(v42)%32)) - v42
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v34 = int32(16)
	goto L13
L12:
	;
	v34 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v41 = int32(4)
	goto L16
L15:
	;
	v41 = v34
	goto L16
L16:
	;
	v54 = v41
	goto L7
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v58 = v54
	goto L20
L19:
	;
	v58 = v55
	goto L20
L20:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v59 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v89 = int32(1)
	v90 = v13 + v89
	v92 = v13 + int32(4)
	if v25&v89 != 0 {
		goto L32
	} else {
		goto L33
	}
L22:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v65 == int32(18) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v76 = int32(1)
	if v59&v76 != 0 {
		v88 = int32(base.Ui32(v59)>>(uint(v76)%32)) - v76
		goto L21
	} else {
		goto L31
	}
L25:
	;
	v68 = int32(16)
	goto L27
L26:
	;
	v68 = int32(0)
	goto L27
L27:
	;
	if base.Ui32((v65-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v75 = int32(4)
	goto L30
L29:
	;
	v75 = v68
	goto L30
L30:
	;
	v88 = v75
	goto L21
L31:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v88 = int32(base.Ui32(v82)>>(uint(int32(2))%32)) - int32(4)
	goto L21
L32:
	;
	v95 = v90
	goto L34
L33:
	;
	v95 = v92
	goto L34
L34:
	;
	v96 = F_pg_mbstrlen_with_len(m, v95, v58)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_rpad[0]))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100*int32(28))+uint32(_c_F_rpad[1])))
	goto L37
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L79
	}
L37:
	;
	if v96 < v24 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v108 = v96
	goto L40
L39:
	;
	v108 = v24
	goto L40
L40:
	;
	if v88 <= int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v111 = v108
	goto L43
L42:
	;
	v111 = v24
	goto L43
L43:
	;
	v113 = base.I64_extend_i32_s(v105) * base.I64_extend_i32_s(v111)
	v117 = base.I32_wrap_i64(v113)
	if base.I32_wrap_i64(int64(base.Ui64(v113)>>(uint(int64(32))%64))) != v117>>(uint(int32(31))%32) {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v122 = v117 + int32(4)
	if base.B2i32(v122 < v117)|base.B2i32(base.Ui32(int32(1073741824)) <= base.Ui32(v122)) != 0 {
		goto L36
	} else {
		goto L45
	}
L45:
	;
	v127 = v111 - v108
	v128 = F_palloc(m, v122)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v131 = v128 + int32(4)
	if v108 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v132&int32(1) != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v154 = v131
	goto L49
L49:
	;
	if v127 != 0 {
		goto L60
	} else {
		goto L61
	}
L50:
	;
	v135 = v90
	goto L52
L51:
	;
	v135 = v92
	goto L52
L52:
	;
	v136 = v131
	v137 = v135
	v145 = v108
	goto L53
L53:
	;
	v147 = F_pg_mblen_unbounded(m, v137)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	v154 = v151
	goto L49
L55:
	;
	if v147 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	base.MemoryCopy(m, v136, v137, v147)
	goto L58
L57:
	;
	goto L58
L58:
	;
	v151 = v136 + v147
	v153 = v145 - int32(1)
	if v153 != 0 {
		v136 = v151
		v137 = v137 + v147
		v145 = v153
		goto L53
	} else {
		goto L59
	}
L59:
	;
	goto L54
L60:
	;
	v165 = int32(1)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v167&v165 != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v197 = v154
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = (v197 - v128) << (uint(int32(2)) % 32)
	return v128
L63:
	;
	v170 = v165
	goto L65
L64:
	;
	v170 = int32(4)
	goto L65
L65:
	;
	v171 = v21 + v170
	v172 = int32(0)
	if v172 < v88 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v175 = v88
	goto L68
L67:
	;
	v175 = v172
	goto L68
L68:
	;
	v176 = v171 + v175
	v177 = v154
	v178 = v171
	v181 = v127
	goto L69
L69:
	;
	v188 = F_pg_mblen_range(m, v178, v176)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L71
	}
L70:
	;
	v197 = v194
	goto L62
L71:
	;
	if v188 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	base.MemoryCopy(m, v177, v178, v188)
	goto L74
L73:
	;
	goto L74
L74:
	;
	v191 = v178 + v188
	if v191 == v176 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v193 = v171
	goto L77
L76:
	;
	v193 = v191
	goto L77
L77:
	;
	v194 = v177 + v188
	v196 = v181 - int32(1)
	if v196 != 0 {
		v177 = v194
		v178 = v193
		v181 = v196
		goto L69
	} else {
		goto L78
	}
L78:
	;
	goto L70
L79:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errmsg(m, int32(_a_F_rpad_0), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_rpad_1), int32(304), int32(_a_F_rpad_2))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
