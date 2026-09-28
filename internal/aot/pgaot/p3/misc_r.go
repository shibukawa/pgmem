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
				v73 = int32(4082)
				F_errmsg(m, v72, int32(0))
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_ReadCheckpointRecord_1), v73, int32(_a_F_ReadCheckpointRecord_2))
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
							v72 = int32(_a_F_ReadCheckpointRecord_3)
							v73 = int32(4092)
							F_errmsg(m, v72, int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_ReadCheckpointRecord_1), v73, int32(_a_F_ReadCheckpointRecord_2))
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
								v72 = int32(_a_F_ReadCheckpointRecord_4)
								v73 = int32(_a_F_ReadCheckpointRecord_5)
								F_errmsg(m, v72, int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ReadCheckpointRecord_1), v73, int32(_a_F_ReadCheckpointRecord_2))
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
									v72 = int32(_a_F_ReadCheckpointRecord_6)
									v73 = int32(_a_F_ReadCheckpointRecord_7)
									F_errmsg(m, v72, int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_ReadCheckpointRecord_1), v73, int32(_a_F_ReadCheckpointRecord_2))
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
							if v59 == int32(122) {
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
										v72 = int32(_a_F_ReadCheckpointRecord_8)
										v73 = int32(_a_F_ReadCheckpointRecord_9)
										F_errmsg(m, v72, int32(0))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_ReadCheckpointRecord_1), v73, int32(_a_F_ReadCheckpointRecord_2))
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
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
	var v62 int32
	_ = v62
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
	var v76 int32
	_ = v76
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
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
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	v14 = m.G0
	v16 = v14 - int32(336)
	m.G0 = v16
	v20 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v23 = v16 + int32(160)
	F_ScanKeyInit(m, v23, int32(4), int32(3), int32(184), int64(1259))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v35 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+56)))
	F_ScanKeyInit(m, v16+int32(216), int32(5), int32(3), int32(184), v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_ScanKeyInit(m, v16+int32(272), int32(6), int32(3), int32(65), base.I64_extend_i32_s(l3))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v50 = F_systable_beginscan(m, v20, int32(2674), int32(1), int32(0), int32(3), v23)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L120
	}
L7:
	;
	v52 = F_systable_getnext(m, v50)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v52 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v55 = base.B2i32(l1 != int32(24))
	v62 = v52
	goto L12
L10:
	;
	goto L11
L11:
	;
	F_systable_endscan(m, v50)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L118
	}
L12:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+22)))
	v71 = v69 + v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+148)) = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+152)) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+156)) = v76
	if v72 <= int32(3255) {
		goto L24
	} else {
		goto L25
	}
L13:
	;
	goto L11
L14:
	;
	v391 = F_systable_getnext(m, v50)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L116
	}
L15:
	;
	F_RememberConstraintForRebuilding(m, v74, l0)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L115
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L111
	}
L17:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v299 = int32(0)
	if v298 == v299 {
		goto L91
	} else {
		goto L92
	}
L18:
	;
	F_GetAttrDefaultColumnAddress(m, v16+int32(136), v74)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L76
	}
L19:
	;
	if l1 != int32(24) {
		goto L14
	} else {
		goto L69
	}
L20:
	;
	if l1 != int32(24) {
		goto L14
	} else {
		goto L62
	}
L21:
	;
	if l1 != int32(24) {
		goto L14
	} else {
		goto L55
	}
L22:
	;
	if l1 != int32(24) {
		goto L14
	} else {
		goto L48
	}
L23:
	;
	v118 = F_get_rel_relkind(m, v74)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L38
	}
L24:
	;
	switch v72 - int32(2604) {
	case 0:
		goto L18
	case 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15:
		goto L16
	case 2:
		goto L15
	case 14:
		goto L21
	case 16:
		goto L20
	default:
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v72 == int32(3256) {
		goto L19
	} else {
		goto L28
	}
L27:
	;
	switch v72 - int32(1255) {
	case 0:
		goto L22
	default:
		goto L16
	case 4:
		goto L23
	}
L28:
	;
	if v72 == int32(3381) {
		goto L17
	} else {
		goto L29
	}
L29:
	;
	if v72 != int32(_a_F_RememberAllDependentForRebuilding_0) {
		goto L16
	} else {
		goto L30
	}
L30:
	;
	if l1 != int32(24) {
		goto L14
	} else {
		goto L31
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errmsg(m, int32(_a_F_RememberAllDependentForRebuilding_1), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v104 = F_getObjectDescription(m, v16+int32(148), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+132)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v104
	v111 = F_errdetail(m, int32(_a_F_RememberAllDependentForRebuilding_2), v16+int32(128))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_RememberAllDependentForRebuilding_3), int32(_a_F_RememberAllDependentForRebuilding_4), int32(_a_F_RememberAllDependentForRebuilding_5))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
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
	if v118&int32(-33) == int32(73) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v16)+152))
	F_RememberIndexForRebuilding(m, v124, l0)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if v118 == int32(83) {
		goto L14
	} else {
		goto L43
	}
L42:
	;
	goto L14
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v136 = F_getObjectDescription(m, v16+int32(148), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v136
	F_errmsg_internal(m, int32(_a_F_RememberAllDependentForRebuilding_6), v16+int32(16))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_RememberAllDependentForRebuilding_3), int32(_a_F_RememberAllDependentForRebuilding_7), int32(_a_F_RememberAllDependentForRebuilding_5))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(_a_F_RememberAllDependentForRebuilding_8), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v163 = F_getObjectDescription(m, v16+int32(148), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v163
	v170 = F_errdetail(m, int32(_a_F_RememberAllDependentForRebuilding_2), v16+int32(32))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_RememberAllDependentForRebuilding_3), int32(_a_F_RememberAllDependentForRebuilding_9), int32(_a_F_RememberAllDependentForRebuilding_5))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errmsg(m, int32(_a_F_RememberAllDependentForRebuilding_10), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v191 = F_getObjectDescription(m, v16+int32(148), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v191
	v198 = F_errdetail(m, int32(_a_F_RememberAllDependentForRebuilding_2), v16+int32(48))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_RememberAllDependentForRebuilding_3), int32(_a_F_RememberAllDependentForRebuilding_11), int32(_a_F_RememberAllDependentForRebuilding_5))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(_a_F_RememberAllDependentForRebuilding_12), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v219 = F_getObjectDescription(m, v16+int32(148), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v219
	v226 = F_errdetail(m, int32(_a_F_RememberAllDependentForRebuilding_2), v16-int32(-64))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_RememberAllDependentForRebuilding_3), int32(_a_F_RememberAllDependentForRebuilding_13), int32(_a_F_RememberAllDependentForRebuilding_5))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errmsg(m, int32(_a_F_RememberAllDependentForRebuilding_14), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v247 = F_getObjectDescription(m, v16+int32(148), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v247
	v254 = F_errdetail(m, int32(_a_F_RememberAllDependentForRebuilding_2), v16+int32(80))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_RememberAllDependentForRebuilding_3), int32(_a_F_RememberAllDependentForRebuilding_15), int32(_a_F_RememberAllDependentForRebuilding_5))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v16)+140))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	if v265 == v266 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L84
	}
L78:
	;
	if l1 != int32(24) {
		goto L14
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	if l1 != int32(24) {
		goto L14
	} else {
		goto L83
	}
L81:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v16)+144))
	if v268 != l3 {
		goto L77
	} else {
		goto L82
	}
L82:
	;
	goto L14
L83:
	;
	goto L77
L84:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errmsg(m, int32(_a_F_RememberAllDependentForRebuilding_16), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v16)+140))
	v282 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+144)))
	v284 = F_get_attname(m, v281, v282, int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+100)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = l4
	v291 = F_errdetail(m, int32(_a_F_RememberAllDependentForRebuilding_17), v16+int32(96))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_RememberAllDependentForRebuilding_3), int32(_a_F_RememberAllDependentForRebuilding_18), int32(_a_F_RememberAllDependentForRebuilding_5))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	if v337 != 0 {
		goto L14
	} else {
		goto L103
	}
L91:
	;
	v337 = int32(0)
	goto L90
L92:
	;
	goto L93
L93:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	if v305 <= int32(0) {
		v331 = v299
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v337 = v331
	goto L90
L95:
	;
	v308 = int32(0)
	if v308 < v305 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v311 = v305
	goto L98
L97:
	;
	v311 = v308
	goto L98
L98:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v298)+12))
	v314 = int32(0)
	goto L99
L99:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v312+v314<<(uint(int32(2))%32))))
	v323 = base.B2i32(v322 == v74)
	if v322 == v74 {
		v331 = v323
		goto L94
	} else {
		goto L101
	}
L100:
	;
	v331 = v323
	goto L94
L101:
	;
	v325 = v314 + int32(1)
	if v325 != v311 {
		v314 = v325
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v338 = int32(0)
	v340 = F_pg_get_statisticsobj_worker(m, v74, v338, v338)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v344 = F_SearchSysCache1(m, int32(64), base.I64_extend_i32_u(v74))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	if v344 == int32(0) {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v344)+16))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+22)))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v351 = F_lappend_oid(m, v350, v74)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v351
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v355 = F_lappend(m, v354, v340)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v355
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v348+v349)+76))
	v361 = F_lappend_oid(m, v358, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v361
	F_ReleaseCatCache(m, v344)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	goto L14
L111:
	;
	v373 = F_getObjectDescription(m, v16+int32(148), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v373
	F_errmsg_internal(m, int32(_a_F_RememberAllDependentForRebuilding_6), v16)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_RememberAllDependentForRebuilding_3), int32(_a_F_RememberAllDependentForRebuilding_19), int32(_a_F_RememberAllDependentForRebuilding_5))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	goto L14
L116:
	;
	if v391 != 0 {
		v62 = v391
		goto L12
	} else {
		goto L117
	}
L117:
	;
	goto L13
L118:
	;
	F_relation_close(m, v20, int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	m.G0 = v16 + int32(336)
	return
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = v74
	F_errmsg_internal(m, int32(_a_F_RememberAllDependentForRebuilding_20), v16+int32(112))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_RememberAllDependentForRebuilding_3), int32(_a_F_RememberAllDependentForRebuilding_21), int32(_a_F_RememberAllDependentForRebuilding_22))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
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
	var v30 int64
	_ = v30
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
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v16 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v21 = F_SearchSysCacheCopy(m, int32(82), base.I64_extend_i32_u(l0), int64(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			if v21 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
				v25 = v23 + v24
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+96))
				v30 = int64(0)
				v32 = F_GetSysCacheOid(m, int32(81), base.I64_extend_i32_u(l1), base.I64_extend_i32_u(l2), v30, v30)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					if v32 != 0 {
						v34 = F_get_typisdefined(m, v32)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							if v34 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return
								} else {
									F_errcode(m, int32(_a_F_RenameTypeInternal_0))
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
										F_errmsg(m, int32(_a_F_RenameTypeInternal_1), v12+int32(16))
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_RenameTypeInternal_2), int32(809), int32(_a_F_RenameTypeInternal_3))
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
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
								v38 = F_moveArrayTypeName(m, v32, l1, l2)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									if v38 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return
										} else {
											F_errcode(m, int32(_a_F_RenameTypeInternal_0))
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
												F_errmsg(m, int32(_a_F_RenameTypeInternal_1), v12+int32(16))
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_RenameTypeInternal_2), int32(809), int32(_a_F_RenameTypeInternal_3))
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
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
										v45 = F_strncpy(m, v25+int32(4), l1, int32(64))
										mBase = m.M
										v46 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v45)+63)) = uint8(v46)
										F_CatalogTupleUpdate(m, v16, v21+int32(4), v21)
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return
										} else {
											v53 = *(*int32)(unsafe.Add(mBase, _c_F_RenameTypeInternal[0]))
											if v53 != 0 {
												v55 = int32(0)
												F_RunObjectPostAlterHook(m, int32(1247), l0, v55, v55, v55)
												mBase = m.M
												v59 = m.ExcPending
												if v59 != 0 {
													return
												} else {
													F_pfree(m, v21)
													mBase = m.M
													v61 = m.ExcPending
													if v61 != 0 {
														return
													} else {
														F_relation_close(m, v16, int32(3))
														mBase = m.M
														v64 = m.ExcPending
														if v64 != 0 {
															return
														} else {
															v65 = int32(0)
															if base.B2i32(v26 == v65)|base.B2i32(v26 == v32) == v65 {
																v71 = F_makeArrayTypeName(m, l1, l2)
																mBase = m.M
																v72 = m.ExcPending
																if v72 != 0 {
																	return
																} else {
																	F_RenameTypeInternal(m, v26, v71, l2)
																	mBase = m.M
																	v74 = m.ExcPending
																	if v74 != 0 {
																		return
																	} else {
																		F_pfree(m, v71)
																		mBase = m.M
																		v76 = m.ExcPending
																		if v76 != 0 {
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
												F_pfree(m, v21)
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return
												} else {
													F_relation_close(m, v16, int32(3))
													mBase = m.M
													v64 = m.ExcPending
													if v64 != 0 {
														return
													} else {
														v65 = int32(0)
														if base.B2i32(v26 == v65)|base.B2i32(v26 == v32) == v65 {
															v71 = F_makeArrayTypeName(m, l1, l2)
															mBase = m.M
															v72 = m.ExcPending
															if v72 != 0 {
																return
															} else {
																F_RenameTypeInternal(m, v26, v71, l2)
																mBase = m.M
																v74 = m.ExcPending
																if v74 != 0 {
																	return
																} else {
																	F_pfree(m, v71)
																	mBase = m.M
																	v76 = m.ExcPending
																	if v76 != 0 {
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
						v45 = F_strncpy(m, v25+int32(4), l1, int32(64))
						mBase = m.M
						v46 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v45)+63)) = uint8(v46)
						F_CatalogTupleUpdate(m, v16, v21+int32(4), v21)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, _c_F_RenameTypeInternal[0]))
							if v53 != 0 {
								v55 = int32(0)
								F_RunObjectPostAlterHook(m, int32(1247), l0, v55, v55, v55)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									F_pfree(m, v21)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										F_relation_close(m, v16, int32(3))
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return
										} else {
											v65 = int32(0)
											if base.B2i32(v26 == v65)|base.B2i32(v26 == v32) == v65 {
												v71 = F_makeArrayTypeName(m, l1, l2)
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return
												} else {
													F_RenameTypeInternal(m, v26, v71, l2)
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return
													} else {
														F_pfree(m, v71)
														mBase = m.M
														v76 = m.ExcPending
														if v76 != 0 {
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
								F_pfree(m, v21)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									F_relation_close(m, v16, int32(3))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										v65 = int32(0)
										if base.B2i32(v26 == v65)|base.B2i32(v26 == v32) == v65 {
											v71 = F_makeArrayTypeName(m, l1, l2)
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return
											} else {
												F_RenameTypeInternal(m, v26, v71, l2)
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return
												} else {
													F_pfree(m, v71)
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
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
				v84 = m.ExcPending
				if v84 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
					F_errmsg_internal(m, int32(_a_F_RenameTypeInternal_4), v12)
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_RenameTypeInternal_2), int32(781), int32(_a_F_RenameTypeInternal_3))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
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
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v353 int32
	_ = v353
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int64
	_ = v437
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int64
	_ = v490
	var v503 int32
	_ = v503
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v695 int64
	_ = v695
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	v6 = m.G0
	v8 = v6 - int32(224)
	m.G0 = v8
	v12 = int32(-1)
	v13 = int32(0)
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
	if v12 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v694 = int32(m.ExcTag)
	v695 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v694 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[0]))
	if v19 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v44 = v13
	goto L9
L9:
	;
	if v44 != 0 {
		goto L24
	} else {
		goto L25
	}
L10:
	;
	F_MemoryContextDelete(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	goto L15
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[0])) = int32(0)
	goto L12
L14:
	;
	F_InitProcess(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L18
	}
L15:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[1]))
	v30 = F_GetBackendTypeDesc(m, v29)
	mBase = m.M
	goto L17
L17:
	;
	goto L14
L18:
	;
	F_BaseInit(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	goto L20
L20:
	;
	v36 = v8 - int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v8 + int32(44)
	goto L23
L21:
	;
	v44 = int32(0)
	goto L9
L23:
	;
	goto L21
L24:
	;
	v45 = int32(_a_F_ReplSlotSyncWorkerMain_0)
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[2])) = v47 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[3])) = int32(0)
	F_EmitErrorReport(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[4])) = v8 - int32(-64)
	v67 = m.G0
	v69 = v67 - int32(32)
	m.G0 = v69
	v72 = int32(967)
	switch v72 {
	case 0, 2:
		goto L30
	default:
		goto L31
	}
L27:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	goto L3
L29:
	;
	v109 = m.G0
	v111 = v109 - int32(32)
	m.G0 = v111
	v114 = int32(968)
	switch v114 {
	case 0, 2:
		goto L40
	default:
		goto L41
	}
L30:
	;
	F_sigemptyset(m, v69+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v69)+24)) = int32(268435456)
	switch v72 {
	case 0:
		goto L35
	default:
		goto L33
	case 2:
		goto L34
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[5])) = int32(965)
	goto L30
L32:
	;
	goto L37
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = int32(_a_F_ReplSlotSyncWorkerMain_1)
	goto L32
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = int32(0)
	goto L32
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = int32(-2)
	goto L32
L37:
	;
	goto L38
L38:
	;
	v101 = F___sigaction(m, int32(1), v69+int32(12), int32(0))
	mBase = m.M
	m.G0 = v69 + int32(32)
	goto L29
L39:
	;
	v151 = m.G0
	v153 = v151 - int32(32)
	m.G0 = v153
	v156 = int32(974)
	switch v156 {
	case 0, 2:
		goto L50
	default:
		goto L51
	}
L40:
	;
	F_sigemptyset(m, v111+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v111)+24)) = int32(268435456)
	switch v114 {
	case 0:
		goto L45
	default:
		goto L43
	case 2:
		goto L44
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[6])) = int32(966)
	goto L40
L42:
	;
	goto L47
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = int32(_a_F_ReplSlotSyncWorkerMain_1)
	goto L42
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = int32(0)
	goto L42
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = int32(-2)
	goto L42
L47:
	;
	goto L48
L48:
	;
	v143 = F___sigaction(m, int32(2), v111+int32(12), int32(0))
	mBase = m.M
	m.G0 = v111 + int32(32)
	goto L39
L49:
	;
	v193 = m.G0
	v195 = v193 - int32(32)
	m.G0 = v195
	v198 = int32(972)
	switch v198 {
	case 0, 2:
		goto L60
	default:
		goto L61
	}
L50:
	;
	F_sigemptyset(m, v153+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v153)+24)) = int32(268435456)
	switch v156 {
	case 0:
		goto L55
	default:
		goto L53
	case 2:
		goto L54
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[7])) = int32(972)
	goto L50
L52:
	;
	goto L57
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v153)+12)) = int32(_a_F_ReplSlotSyncWorkerMain_1)
	goto L52
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+12)) = int32(0)
	goto L52
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+12)) = int32(-2)
	goto L52
L57:
	;
	goto L58
L58:
	;
	v185 = F___sigaction(m, int32(15), v153+int32(12), int32(0))
	mBase = m.M
	m.G0 = v153 + int32(32)
	goto L49
L59:
	;
	v235 = m.G0
	v237 = v235 - int32(32)
	m.G0 = v237
	v240 = int32(970)
	switch v240 {
	case 0, 2:
		goto L70
	default:
		goto L71
	}
L60:
	;
	F_sigemptyset(m, v195+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v195)+24)) = int32(268435456)
	switch v198 {
	case 0:
		goto L65
	default:
		goto L63
	case 2:
		goto L64
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[8])) = int32(970)
	goto L60
L62:
	;
	goto L67
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+12)) = int32(_a_F_ReplSlotSyncWorkerMain_1)
	goto L62
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+12)) = int32(0)
	goto L62
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+12)) = int32(-2)
	goto L62
L67:
	;
	goto L68
L68:
	;
	v227 = F___sigaction(m, int32(8), v195+int32(12), int32(0))
	mBase = m.M
	m.G0 = v195 + int32(32)
	goto L59
L69:
	;
	v275 = int32(0)
	v277 = m.G0
	v279 = v277 - int32(32)
	m.G0 = v279
	switch v275 {
	case 0, 2:
		goto L80
	default:
		goto L81
	}
L70:
	;
	F_sigemptyset(m, v237+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v237)+24)) = int32(268435456)
	switch v240 {
	case 0:
		goto L75
	default:
		goto L73
	case 2:
		goto L74
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[9])) = int32(968)
	goto L70
L72:
	;
	goto L77
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v237)+12)) = int32(_a_F_ReplSlotSyncWorkerMain_1)
	goto L72
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+12)) = int32(0)
	goto L72
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+12)) = int32(-2)
	goto L72
L77:
	;
	goto L78
L78:
	;
	v269 = F___sigaction(m, int32(10), v237+int32(12), int32(0))
	mBase = m.M
	m.G0 = v237 + int32(32)
	goto L69
L79:
	;
	v317 = int32(0)
	v319 = m.G0
	v321 = v319 - int32(32)
	m.G0 = v321
	switch v317 {
	case 0, 2:
		goto L90
	default:
		goto L91
	}
L80:
	;
	F_sigemptyset(m, v279+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v279)+24)) = int32(268435456)
	switch v275 {
	case 0:
		goto L85
	default:
		goto L83
	case 2:
		goto L84
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[10])) = int32(-2)
	goto L80
L82:
	;
	goto L87
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v279)+12)) = int32(_a_F_ReplSlotSyncWorkerMain_1)
	goto L82
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279)+12)) = int32(0)
	goto L82
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279)+12)) = int32(-2)
	goto L82
L87:
	;
	goto L88
L88:
	;
	v311 = F___sigaction(m, int32(12), v279+int32(12), int32(0))
	mBase = m.M
	m.G0 = v279 + int32(32)
	goto L79
L89:
	;
	v361 = m.G0
	v363 = v361 - int32(32)
	m.G0 = v363
	v365 = int32(2)
	switch v365 {
	case 0, 2:
		goto L100
	default:
		goto L101
	}
L90:
	;
	F_sigemptyset(m, v321+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v321)+24)) = int32(268435456)
	switch v317 {
	case 0:
		goto L95
	default:
		goto L93
	case 2:
		goto L94
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[11])) = int32(-2)
	goto L90
L92:
	;
	goto L97
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v321)+12)) = int32(_a_F_ReplSlotSyncWorkerMain_1)
	goto L92
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+12)) = int32(0)
	goto L92
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+12)) = int32(-2)
	goto L92
L97:
	;
	goto L98
L98:
	;
	v353 = F___sigaction(m, int32(13), v321+int32(12), int32(0))
	mBase = m.M
	m.G0 = v321 + int32(32)
	goto L89
L99:
	;
	v400 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[12]))
	F_check_and_set_sync_info(m, v400)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L6
	} else {
		goto L109
	}
L100:
	;
	F_sigemptyset(m, v363+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v363)+24)) = int32(268435456)
	switch v365 {
	case 0:
		goto L105
	default:
		goto L103
	case 2:
		goto L104
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[13])) = int32(0)
	goto L100
L102:
	;
	goto L106
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v363)+12)) = int32(_a_F_ReplSlotSyncWorkerMain_1)
	v388 = int32(268435461)
	goto L102
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363)+12)) = int32(0)
	v388 = int32(268435457)
	goto L102
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363)+12)) = int32(-2)
	v388 = int32(268435457)
	goto L102
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363)+24)) = v388
	goto L108
L108:
	;
	v395 = F___sigaction(m, int32(17), v363+int32(12), int32(0))
	mBase = m.M
	m.G0 = v363 + int32(32)
	goto L99
L109:
	;
	v405 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L6
	} else {
		goto L110
	}
L110:
	;
	if v405 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	F_errmsg(m, int32(_a_F_ReplSlotSyncWorkerMain_2), int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L6
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	F_before_shmem_exit(m, int32(1090), int64(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L6
	} else {
		goto L116
	}
L114:
	;
	F_errfinish(m, int32(_a_F_ReplSlotSyncWorkerMain_3), int32(1675), int32(_a_F_ReplSlotSyncWorkerMain_4))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v420 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[14])) = v420
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[15])) = v420
	v430 = v420
	goto L118
L117:
	;
	F_load_file(m, int32(_a_F_ReplSlotSyncWorkerMain_5), int32(0))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L6
	} else {
		goto L123
	}
L118:
	;
	v432 = int32(40)
	v433 = v430 * v432
	v434 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v433)+uint32(_c_F_ReplSlotSyncWorkerMain[16]))) = uint8(v434)
	*(*int32)(unsafe.Add(mBase, uint32(v433)+uint32(_c_F_ReplSlotSyncWorkerMain[17]))) = v430
	v437 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v433)+uint32(_c_F_ReplSlotSyncWorkerMain[18]))) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v433)+uint32(_c_F_ReplSlotSyncWorkerMain[19]))) = v434
	*(*int64)(unsafe.Add(mBase, uint32(v433)+uint32(_c_F_ReplSlotSyncWorkerMain[20]))) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v433)+uint32(_c_F_ReplSlotSyncWorkerMain[21]))) = v434
	*(*uint8)(unsafe.Add(mBase, uint32(v433)+uint32(_c_F_ReplSlotSyncWorkerMain[22]))) = uint8(v434)
	v448 = v430 | int32(1)
	v450 = v448 * v432
	*(*uint8)(unsafe.Add(mBase, uint32(v450)+uint32(_c_F_ReplSlotSyncWorkerMain[16]))) = uint8(v434)
	*(*int32)(unsafe.Add(mBase, uint32(v450)+uint32(_c_F_ReplSlotSyncWorkerMain[17]))) = v448
	*(*int64)(unsafe.Add(mBase, uint32(v450)+uint32(_c_F_ReplSlotSyncWorkerMain[18]))) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v450)+uint32(_c_F_ReplSlotSyncWorkerMain[19]))) = v434
	*(*int64)(unsafe.Add(mBase, uint32(v450)+uint32(_c_F_ReplSlotSyncWorkerMain[20]))) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v450)+uint32(_c_F_ReplSlotSyncWorkerMain[21]))) = v434
	*(*uint8)(unsafe.Add(mBase, uint32(v450)+uint32(_c_F_ReplSlotSyncWorkerMain[22]))) = uint8(v434)
	v465 = v430 | int32(2)
	v467 = v465 * v432
	*(*uint8)(unsafe.Add(mBase, uint32(v467)+uint32(_c_F_ReplSlotSyncWorkerMain[16]))) = uint8(v434)
	*(*int32)(unsafe.Add(mBase, uint32(v467)+uint32(_c_F_ReplSlotSyncWorkerMain[17]))) = v465
	*(*int64)(unsafe.Add(mBase, uint32(v467)+uint32(_c_F_ReplSlotSyncWorkerMain[18]))) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v467)+uint32(_c_F_ReplSlotSyncWorkerMain[19]))) = v434
	*(*int64)(unsafe.Add(mBase, uint32(v467)+uint32(_c_F_ReplSlotSyncWorkerMain[20]))) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v467)+uint32(_c_F_ReplSlotSyncWorkerMain[21]))) = v434
	*(*uint8)(unsafe.Add(mBase, uint32(v467)+uint32(_c_F_ReplSlotSyncWorkerMain[22]))) = uint8(v434)
	if v430 != int32(20) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v503 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[23])) = uint8(v503)
	F_pqsignal_be(m, int32(14), int32(1992))
	mBase = m.M
	goto L117
L120:
	;
	v484 = v430 | int32(3)
	v486 = v484 * int32(40)
	v487 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v486)+uint32(_c_F_ReplSlotSyncWorkerMain[16]))) = uint8(v487)
	*(*int32)(unsafe.Add(mBase, uint32(v486)+uint32(_c_F_ReplSlotSyncWorkerMain[17]))) = v484
	v490 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v486)+uint32(_c_F_ReplSlotSyncWorkerMain[18]))) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v486)+uint32(_c_F_ReplSlotSyncWorkerMain[19]))) = v487
	*(*int64)(unsafe.Add(mBase, uint32(v486)+uint32(_c_F_ReplSlotSyncWorkerMain[20]))) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v486)+uint32(_c_F_ReplSlotSyncWorkerMain[21]))) = v487
	*(*uint8)(unsafe.Add(mBase, uint32(v486)+uint32(_c_F_ReplSlotSyncWorkerMain[22]))) = uint8(v487)
	v430 = v430 + int32(4)
	goto L118
L121:
	;
	goto L122
L122:
	;
	goto L119
L123:
	;
	F_pgmem_sigprocmask(m, int32(_a_F_ReplSlotSyncWorkerMain_6), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	F_SetConfigOption(m, int32(_a_F_ReplSlotSyncWorkerMain_7), int32(_a_F_ReplSlotSyncWorkerMain_8), int32(5), int32(10))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L6
	} else {
		goto L125
	}
L125:
	;
	v523 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[24]))
	v525 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[25]))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)+20))
	v527 = m.T0[v526].(func(*base.Module, int32) int32)(m, v523)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L6
	} else {
		goto L126
	}
L126:
	;
	if v527 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L6
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v550 = int32(0)
	F_InitPostgres(m, v527, v550, v550, v550, v550, v550)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L6
	} else {
		goto L134
	}
L130:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L6
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(_a_F_ReplSlotSyncWorkerMain_9)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_ReplSlotSyncWorkerMain_10)
	F_errmsg(m, int32(_a_F_ReplSlotSyncWorkerMain_11), v8)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L6
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_ReplSlotSyncWorkerMain_3), int32(1215), int32(_a_F_ReplSlotSyncWorkerMain_12))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L6
	} else {
		goto L133
	}
L133:
	;
	goto L3
L134:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[26])) = int32(2)
	v561 = v8 + int32(48)
	F_initStringInfo(m, v561)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L6
	} else {
		goto L135
	}
L135:
	;
	v565 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[27]))
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565))))
	if v566 != 0 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v581 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[24]))
	v582 = int32(0)
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
	v589 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[25]))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v589)))
	v591 = m.T0[v590].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v581, v582, v582, v582, v585, v8+int32(220))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L6
	} else {
		goto L142
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(_a_F_ReplSlotSyncWorkerMain_13)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v565
	F_appendStringInfo(m, v561, int32(_a_F_ReplSlotSyncWorkerMain_14), v8+int32(32))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L6
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	F_appendStringInfoString(m, v8+int32(48), int32(_a_F_ReplSlotSyncWorkerMain_13))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L6
	} else {
		goto L141
	}
L140:
	;
	goto L136
L141:
	;
	goto L136
L142:
	;
	if v591 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L6
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
	F_pfree(m, v616)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L6
	} else {
		goto L150
	}
L146:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L6
	} else {
		goto L147
	}
L147:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v8)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v603
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v602
	F_errmsg(m, int32(_a_F_ReplSlotSyncWorkerMain_15), v8+int32(16))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L6
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_ReplSlotSyncWorkerMain_3), int32(1738), int32(_a_F_ReplSlotSyncWorkerMain_4))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L6
	} else {
		goto L149
	}
L149:
	;
	goto L3
L150:
	;
	F_before_shmem_exit(m, int32(1091), base.I64_extend_i32_u(v591))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L6
	} else {
		goto L151
	}
L151:
	;
	F_validate_remote_info(m, v591)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L6
	} else {
		goto L152
	}
L152:
	;
	goto L153
L153:
	;
	v631 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[28]))
	if v631 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L6
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v635 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[29]))
	if v635 != 0 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	goto L157
L159:
	;
	F_slotsync_reread_config(m)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L6
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v639 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[30]))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v639)+20))
	v642 = base.B2i32(v640 == int32(2))
	goto L163
L162:
	;
	goto L161
L163:
	;
	if v642 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L6
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v648 = F_fetch_remote_slots(m, v591, int32(0))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L6
	} else {
		goto L168
	}
L167:
	;
	goto L166
L168:
	;
	v651 = F_synchronize_slots(m, v648, int32(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L6
	} else {
		goto L169
	}
L169:
	;
	F_list_free_deep(m, v648)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L6
	} else {
		goto L170
	}
L170:
	;
	if v642 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L6
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v659 = int32(_a_F_ReplSlotSyncWorkerMain_16)
	v661 = int32(_a_F_ReplSlotSyncWorkerMain_17)
	v663 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[31]))
	v665 = v663 << (uint(int32(1)) % 32)
	if v661 <= v665 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	goto L173
L175:
	;
	v668 = v661
	goto L177
L176:
	;
	v668 = v665
	goto L177
L177:
	;
	if v651 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v669 = int32(200)
	goto L180
L179:
	;
	v669 = v668
	goto L180
L180:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[31])) = v669
	v672 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[32]))
	v675 = F_WaitLatch(m, v672, int32(41), v669, int32(83886091))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L6
	} else {
		goto L181
	}
L181:
	;
	if v675&int32(1) == int32(0) {
		goto L153
	} else {
		goto L182
	}
L182:
	;
	v682 = *(*int32)(unsafe.Add(mBase, _c_F_ReplSlotSyncWorkerMain[32]))
	v683 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v682))) = v683
	v688 = base.AtomicRmwOr32(m, v683, int32(_a_F_ReplSlotSyncWorkerMain_18), v683)
	goto L183
L183:
	;
	goto L153
L184:
	;
	v699 = int32(v695)
	m.G0 = v8
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v699)+4))
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v699)))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v702)))
	if v8+int32(44) == v705 {
		goto L187
	} else {
		goto L188
	}
L185:
	;
	m.ExcPending = 1
	goto L193
L186:
	;
	if v709 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L187:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v702)+4))
	v709 = v707
	goto L189
L188:
	;
	v709 = int32(0)
	goto L189
L189:
	;
	goto L186
L190:
	;
	F___wasm_longjmp(m, v702, v701)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	goto L192
L192:
	;
	v12 = v709
	v13 = v701
	goto L1
L193:
	;
	return
L194:
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
	var v434 int64
	_ = v434
	var v435 int64
	_ = v435
	var v438 int64
	_ = v438
	var v439 int64
	_ = v439
	var v440 int64
	_ = v440
	var v441 int64
	_ = v441
	var v442 int64
	_ = v442
	var v443 int64
	_ = v443
	var v444 int64
	_ = v444
	var v445 int64
	_ = v445
	var v446 int64
	_ = v446
	var v447 int64
	_ = v447
	var v448 int64
	_ = v448
	var v449 int64
	_ = v449
	var v450 int64
	_ = v450
	var v451 int64
	_ = v451
	var v452 int64
	_ = v452
	var v453 int64
	_ = v453
	var v454 int64
	_ = v454
	var v455 int64
	_ = v455
	var v456 int64
	_ = v456
	var v457 int64
	_ = v457
	var v458 int64
	_ = v458
	var v459 int64
	_ = v459
	var v460 int64
	_ = v460
	var v461 int64
	_ = v461
	var v462 int64
	_ = v462
	var v463 int64
	_ = v463
	var v464 int64
	_ = v464
	var v465 int64
	_ = v465
	var v466 int64
	_ = v466
	var v467 int64
	_ = v467
	var v468 int64
	_ = v468
	var v500 int64
	_ = v500
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v543 int32
	_ = v543
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v625 int32
	_ = v625
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v717 int32
	_ = v717
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v739 int32
	_ = v739
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v769 int32
	_ = v769
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v883 int32
	_ = v883
	var v888 int32
	_ = v888
	var v894 int32
	_ = v894
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v951 int32
	_ = v951
	var v958 int32
	_ = v958
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v994 int32
	_ = v994
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1033 int32
	_ = v1033
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1075 int32
	_ = v1075
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1092 int32
	_ = v1092
	var v1097 int32
	_ = v1097
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1119 int32
	_ = v1119
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1145 int32
	_ = v1145
	var v1150 int32
	_ = v1150
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
	v1135 = m.ExcPending
	if v1135 != 0 {
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
	v1119 = m.ExcPending
	if v1119 != 0 {
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
	v1107 = F_ReadDir(m, v14, l0)
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
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
	v702 = v12 + int32(208)
	v703 = F_AllocateDir(m, v702)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L4
	} else {
		goto L147
	}
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+3288)) = int64(17179869188)
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+3316)) = v275
	v282 = F_hash_create(m, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_6), int64(32), v12+int32(3280), int32(1064))
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
	v434 = *(*int64)(unsafe.Add(mBase, uint32(v433)+8))
	v435 = *(*int64)(unsafe.Add(mBase, uint32(v433)+808))
	if v435 != int64(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if v500 == int64(0) {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v438 = *(*int64)(unsafe.Add(mBase, uint32(v433)+752))
	v439 = *(*int64)(unsafe.Add(mBase, uint32(v433)+728))
	v440 = *(*int64)(unsafe.Add(mBase, uint32(v433)+704))
	v441 = *(*int64)(unsafe.Add(mBase, uint32(v433)+680))
	v442 = *(*int64)(unsafe.Add(mBase, uint32(v433)+656))
	v443 = *(*int64)(unsafe.Add(mBase, uint32(v433)+632))
	v444 = *(*int64)(unsafe.Add(mBase, uint32(v433)+608))
	v445 = *(*int64)(unsafe.Add(mBase, uint32(v433)+584))
	v446 = *(*int64)(unsafe.Add(mBase, uint32(v433)+560))
	v447 = *(*int64)(unsafe.Add(mBase, uint32(v433)+536))
	v448 = *(*int64)(unsafe.Add(mBase, uint32(v433)+512))
	v449 = *(*int64)(unsafe.Add(mBase, uint32(v433)+488))
	v450 = *(*int64)(unsafe.Add(mBase, uint32(v433)+464))
	v451 = *(*int64)(unsafe.Add(mBase, uint32(v433)+440))
	v452 = *(*int64)(unsafe.Add(mBase, uint32(v433)+416))
	v453 = *(*int64)(unsafe.Add(mBase, uint32(v433)+392))
	v454 = *(*int64)(unsafe.Add(mBase, uint32(v433)+368))
	v455 = *(*int64)(unsafe.Add(mBase, uint32(v433)+344))
	v456 = *(*int64)(unsafe.Add(mBase, uint32(v433)+320))
	v457 = *(*int64)(unsafe.Add(mBase, uint32(v433)+296))
	v458 = *(*int64)(unsafe.Add(mBase, uint32(v433)+272))
	v459 = *(*int64)(unsafe.Add(mBase, uint32(v433)+248))
	v460 = *(*int64)(unsafe.Add(mBase, uint32(v433)+224))
	v461 = *(*int64)(unsafe.Add(mBase, uint32(v433)+200))
	v462 = *(*int64)(unsafe.Add(mBase, uint32(v433)+176))
	v463 = *(*int64)(unsafe.Add(mBase, uint32(v433)+152))
	v464 = *(*int64)(unsafe.Add(mBase, uint32(v433)+128))
	v465 = *(*int64)(unsafe.Add(mBase, uint32(v433)+104))
	v466 = *(*int64)(unsafe.Add(mBase, uint32(v433)+80))
	v467 = *(*int64)(unsafe.Add(mBase, uint32(v433)+56))
	v468 = *(*int64)(unsafe.Add(mBase, uint32(v433)+32))
	v500 = v438 + (v439 + (v440 + (v441 + (v442 + (v443 + (v444 + (v445 + (v446 + (v447 + (v448 + (v449 + (v450 + (v451 + (v452 + (v453 + (v454 + (v455 + (v456 + (v457 + (v458 + (v459 + (v460 + (v461 + (v462 + (v463 + (v464 + (v465 + (v466 + (v467 + (v468 + v434))))))))))))))))))))))))))))))
	goto L102
L101:
	;
	v500 = v434
	goto L102
L102:
	;
	goto L99
L103:
	;
	F_hash_destroy(m, v282)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L4
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v506 = v12 + int32(208)
	v507 = F_AllocateDir(m, v506)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L4
	} else {
		goto L107
	}
L106:
	;
	goto L19
L107:
	;
	v509 = F_ReadDir(m, v507, v506)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	if v509 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v514 = v509
	goto L112
L110:
	;
	goto L111
L111:
	;
	F_FreeDir(m, v507)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L4
	} else {
		goto L144
	}
L112:
	;
	v521 = v514 + int32(19)
	v525 = v12 + int32(2256)
	v528 = int32(0)
	v533 = m.G0
	v535 = v533 - int32(16)
	m.G0 = v535
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[7]))) = v528
	*(*int32)(unsafe.Add(mBase, uint32(v525))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[8]))) = v528
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521))))
	if base.Ui32((v543-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v625 = v528
		goto L116
	} else {
		goto L117
	}
L113:
	;
	goto L111
L114:
	;
	v675 = F_ReadDir(m, v507, v12+int32(208))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L4
	} else {
		goto L142
	}
L115:
	;
	if v625 == int32(0) {
		goto L114
	} else {
		goto L132
	}
L116:
	;
	m.G0 = v535 + int32(16)
	goto L115
L117:
	;
	v550 = int32(_a_F_ResetUnloggedRelationsInTablespaceDir_10)
	*(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0])) = int32(0)
	v556 = F_strtoul(m, v521, v535+int32(8), int32(10))
	mBase = m.M
	v558 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0]))
	if v558 != 0 {
		v625 = v528
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v535)+8))
	if base.B2i32(v556 == int32(0))|base.B2i32(v521 == v561) != 0 {
		v625 = v528
		goto L116
	} else {
		goto L119
	}
L119:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561))))
	if v564 != int32(95) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	if v580&int32(255) == int32(46) {
		goto L125
	} else {
		goto L126
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v535)+12)) = int32(0)
	v580 = v564
	v581 = v561
	goto L120
L122:
	;
	goto L123
L123:
	;
	v573 = F_forkname_chars(m, v561+int32(1), v535+int32(12))
	mBase = m.M
	if v573 <= int32(0) {
		v625 = v528
		goto L116
	} else {
		goto L124
	}
L124:
	;
	v578 = v573 + v561 + int32(1)
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578))))
	v580 = v579
	v581 = v578
	goto L120
L125:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v581)+1)))
	if base.Ui32((v586-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v625 = v528
		goto L116
	} else {
		goto L128
	}
L126:
	;
	v612 = v528
	v613 = v580
	goto L127
L127:
	;
	if v613&int32(255) != 0 {
		v625 = v528
		goto L116
	} else {
		goto L131
	}
L128:
	;
	v593 = int32(_a_F_ResetUnloggedRelationsInTablespaceDir_10)
	*(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0])) = int32(0)
	v597 = v581 + int32(1)
	v601 = F_strtoul(m, v597, v535+int32(8), int32(10))
	mBase = m.M
	v603 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0]))
	if v603 != 0 {
		v625 = v528
		goto L116
	} else {
		goto L129
	}
L129:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v535)+8))
	if base.B2i32(v601 == int32(0))|base.B2i32(v597 == v606) != 0 {
		v625 = v528
		goto L116
	} else {
		goto L130
	}
L130:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v606))))
	v612 = v601
	v613 = v609
	goto L127
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[7]))) = v556
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v535)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v525))) = v617
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[8]))) = v612
	v625 = int32(1)
	goto L116
L132:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2256))
	if v631 == int32(3) {
		goto L114
	} else {
		goto L133
	}
L133:
	;
	v634 = int32(0)
	v636 = F_hash_search(m, v282, v12+int32(_a_F_ResetUnloggedRelationsInTablespaceDir_11), v634, v634)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	if v636 == int32(0) {
		goto L114
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+148)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v12 + int32(208)
	v645 = v12 + int32(_a_F_ResetUnloggedRelationsInTablespaceDir_12)
	v650 = F_pg_snprintf(m, v645, int32(2048), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_5), v12+int32(144))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	v652 = F_unlink(m, v645)
	mBase = m.M
	if v652 < int32(0) {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v657 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	if v657 == int32(0) {
		goto L114
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v645
	F_errmsg_internal(m, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_13), v12+int32(128))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_2), int32(264), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_14))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	goto L114
L142:
	;
	if v675 != 0 {
		v514 = v675
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
	v689 = m.ExcPending
	if v689 != 0 {
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
	v705 = F_ReadDir(m, v703, v702)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	if v705 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v710 = v705
	goto L152
L150:
	;
	goto L151
L151:
	;
	F_FreeDir(m, v703)
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L4
	} else {
		goto L190
	}
L152:
	;
	v717 = v710 + int32(19)
	v724 = int32(0)
	v729 = m.G0
	v731 = v729 - int32(16)
	m.G0 = v731
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[7]))) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[8]))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[9]))) = v724
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v717))))
	if base.Ui32((v739-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v821 = v724
		goto L156
	} else {
		goto L157
	}
L153:
	;
	goto L151
L154:
	;
	v900 = F_ReadDir(m, v703, v12+int32(208))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L4
	} else {
		goto L188
	}
L155:
	;
	if v821 == int32(0) {
		goto L154
	} else {
		goto L172
	}
L156:
	;
	m.G0 = v731 + int32(16)
	goto L155
L157:
	;
	v746 = int32(_a_F_ResetUnloggedRelationsInTablespaceDir_10)
	*(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0])) = int32(0)
	v752 = F_strtoul(m, v717, v731+int32(8), int32(10))
	mBase = m.M
	v754 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0]))
	if v754 != 0 {
		v821 = v724
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v731)+8))
	if base.B2i32(v752 == int32(0))|base.B2i32(v717 == v757) != 0 {
		v821 = v724
		goto L156
	} else {
		goto L159
	}
L159:
	;
	v760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v757))))
	if v760 != int32(95) {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	if v776&int32(255) == int32(46) {
		goto L165
	} else {
		goto L166
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v731)+12)) = int32(0)
	v776 = v760
	v777 = v757
	goto L160
L162:
	;
	goto L163
L163:
	;
	v769 = F_forkname_chars(m, v757+int32(1), v731+int32(12))
	mBase = m.M
	if v769 <= int32(0) {
		v821 = v724
		goto L156
	} else {
		goto L164
	}
L164:
	;
	v774 = v769 + v757 + int32(1)
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v774))))
	v776 = v775
	v777 = v774
	goto L160
L165:
	;
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v777)+1)))
	if base.Ui32((v782-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v821 = v724
		goto L156
	} else {
		goto L168
	}
L166:
	;
	v808 = v724
	v809 = v776
	goto L167
L167:
	;
	if v809&int32(255) != 0 {
		v821 = v724
		goto L156
	} else {
		goto L171
	}
L168:
	;
	v789 = int32(_a_F_ResetUnloggedRelationsInTablespaceDir_10)
	*(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0])) = int32(0)
	v793 = v777 + int32(1)
	v797 = F_strtoul(m, v793, v731+int32(8), int32(10))
	mBase = m.M
	v799 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0]))
	if v799 != 0 {
		v821 = v724
		goto L156
	} else {
		goto L169
	}
L169:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v731)+8))
	if base.B2i32(v797 == int32(0))|base.B2i32(v793 == v802) != 0 {
		v821 = v724
		goto L156
	} else {
		goto L170
	}
L170:
	;
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v802))))
	v808 = v797
	v809 = v805
	goto L167
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[7]))) = v752
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v731)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[8]))) = v813
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[9]))) = v808
	v821 = int32(1)
	goto L156
L172:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[8])))
	if v827 != int32(3) {
		goto L154
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+100)) = v717
	v832 = v12 + int32(208)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v832
	v840 = F_pg_snprintf(m, v12+int32(3280), int32(2048), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_5), v12+int32(96))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L4
	} else {
		goto L174
	}
L174:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[7])))
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[9])))
	if v843 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v871 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L4
	} else {
		goto L181
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v832
	v854 = F_pg_snprintf(m, v12+int32(2256), int32(1024), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_15), v12-int32(-64))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L4
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+88)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v12 + int32(208)
	v867 = F_pg_snprintf(m, v12+int32(2256), int32(1024), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_16), v12+int32(80))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
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
	if v871 != 0 {
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
	v883 = m.ExcPending
	if v883 != 0 {
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
	v894 = m.ExcPending
	if v894 != 0 {
		goto L4
	} else {
		goto L187
	}
L185:
	;
	F_errfinish(m, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_2), int32(314), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_14))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
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
	if v900 != 0 {
		v710 = v900
		goto L152
	} else {
		goto L189
	}
L189:
	;
	goto L153
L190:
	;
	v914 = v12 + int32(208)
	v915 = F_AllocateDir(m, v914)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L4
	} else {
		goto L191
	}
L191:
	;
	v917 = F_ReadDir(m, v915, v914)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	if v917 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v922 = v917
	goto L196
L194:
	;
	goto L195
L195:
	;
	F_FreeDir(m, v915)
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L4
	} else {
		goto L227
	}
L196:
	;
	v929 = v922 + int32(19)
	v931 = v12 + int32(2256)
	v936 = int32(0)
	v941 = m.G0
	v943 = v941 - int32(16)
	m.G0 = v943
	*(*int32)(unsafe.Add(mBase, uint32(v931))) = v936
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[8]))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[7]))) = v936
	v951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v929))))
	if base.Ui32((v951-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v1033 = v936
		goto L200
	} else {
		goto L201
	}
L197:
	;
	goto L195
L198:
	;
	v1080 = F_ReadDir(m, v915, v12+int32(208))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L4
	} else {
		goto L225
	}
L199:
	;
	if v1033 == int32(0) {
		goto L198
	} else {
		goto L216
	}
L200:
	;
	m.G0 = v943 + int32(16)
	goto L199
L201:
	;
	v958 = int32(_a_F_ResetUnloggedRelationsInTablespaceDir_10)
	*(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0])) = int32(0)
	v964 = F_strtoul(m, v929, v943+int32(8), int32(10))
	mBase = m.M
	v966 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0]))
	if v966 != 0 {
		v1033 = v936
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v943)+8))
	if base.B2i32(v964 == int32(0))|base.B2i32(v929 == v969) != 0 {
		v1033 = v936
		goto L200
	} else {
		goto L203
	}
L203:
	;
	v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v969))))
	if v972 != int32(95) {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	if v988&int32(255) == int32(46) {
		goto L209
	} else {
		goto L210
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v943)+12)) = int32(0)
	v988 = v972
	v989 = v969
	goto L204
L206:
	;
	goto L207
L207:
	;
	v981 = F_forkname_chars(m, v969+int32(1), v943+int32(12))
	mBase = m.M
	if v981 <= int32(0) {
		v1033 = v936
		goto L200
	} else {
		goto L208
	}
L208:
	;
	v986 = v981 + v969 + int32(1)
	v987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v986))))
	v988 = v987
	v989 = v986
	goto L204
L209:
	;
	v994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v989)+1)))
	if base.Ui32((v994-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v1033 = v936
		goto L200
	} else {
		goto L212
	}
L210:
	;
	v1020 = v936
	v1021 = v988
	goto L211
L211:
	;
	if v1021&int32(255) != 0 {
		v1033 = v936
		goto L200
	} else {
		goto L215
	}
L212:
	;
	v1001 = int32(_a_F_ResetUnloggedRelationsInTablespaceDir_10)
	*(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0])) = int32(0)
	v1005 = v989 + int32(1)
	v1009 = F_strtoul(m, v1005, v943+int32(8), int32(10))
	mBase = m.M
	v1011 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelationsInTablespaceDir[0]))
	if v1011 != 0 {
		v1033 = v936
		goto L200
	} else {
		goto L213
	}
L213:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v943)+8))
	if base.B2i32(v1009 == int32(0))|base.B2i32(v1005 == v1014) != 0 {
		v1033 = v936
		goto L200
	} else {
		goto L214
	}
L214:
	;
	v1017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1014))))
	v1020 = v1009
	v1021 = v1017
	goto L211
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v931))) = v964
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v943)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[8]))) = v1025
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[7]))) = v1020
	v1033 = int32(1)
	goto L200
L216:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[8])))
	if v1039 != int32(3) {
		goto L198
	} else {
		goto L217
	}
L217:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2256))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_ResetUnloggedRelationsInTablespaceDir[7])))
	if v1043 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	F_fsync_fname(m, v12+int32(3280), int32(0))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L4
	} else {
		goto L224
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v1042
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(208)
	v1056 = F_pg_snprintf(m, v12+int32(3280), int32(1024), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_15), v12+int32(16))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L4
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v1042
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v12 + int32(208)
	v1069 = F_pg_snprintf(m, v12+int32(3280), int32(1024), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_16), v12+int32(32))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
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
	if v1080 != 0 {
		v922 = v1080
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
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L4
	} else {
		goto L228
	}
L228:
	;
	goto L19
L229:
	;
	if v1107 != 0 {
		v46 = v1107
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
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L4
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v12 + int32(_a_F_ResetUnloggedRelationsInTablespaceDir_12)
	F_errmsg(m, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_18), v12+int32(112))
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L4
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(_a_F_ResetUnloggedRelationsInTablespaceDir_2), int32(262), int32(_a_F_ResetUnloggedRelationsInTablespaceDir_14))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
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
func F_r_Step_1c_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 <= v8 {
		v285 = v2
		return v285
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v6-int32(1)))))
		v16 = v14 - int32(100)
		v17 = int32(0)
		if base.B2i32(v16 == v17)|base.B2i32(v16 == int32(16)) == v17 {
			v285 = v2
			return v285
		} else {
			v27 = F_find_among_b(m, l0, int32(_a_F_r_Step_1c_2_0), int32(2), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				if v27 == int32(0) {
					v285 = v2
					return v285
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v33
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					if v33 < v35 {
						v285 = v2
						return v285
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v38 = int32(2)
						v40 = int32(0)
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v42-v43 < v38 {
							v53 = v40
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v49 = F_memcmp(m, v46+v42-v38, int32(_a_F_r_Step_1c_2_1), v38)
							mBase = m.M
							if v49 != 0 {
								v53 = v40
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v42 - v38
								v53 = int32(1)
							}
						}
						if v53 != 0 {
							v285 = v2
							return v285
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v55 = v37 - v33
							v56 = v54 - v55
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v56
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							if v56 <= v71 {
								v180 = int32(-1)
								v187 = v180
							} else {
								v88 = int32(1)
								v89 = v56 - v88
								v91 = int32(*(*int8)(unsafe.Add(mBase, uint32(v72+v89))))
								v93 = v91 & int32(255)
								if base.B2i32(v89 == v71)|base.B2i32(int32(0) <= v91) != 0 {
									v151 = v93
									v155 = v88
								} else {
									v100 = v93 & int32(63)
									v102 = v56 - int32(2)
									v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v102))))
									v106 = v104 << (uint(int32(6)) % 32)
									if base.B2i32(v102 != v71)&base.B2i32(base.Ui32(v104) < base.Ui32(int32(192))) == int32(0) {
										v151 = v106&int32(1984) | v100
										v155 = int32(2)
									} else {
										v119 = v106&int32(4032) | v100
										v121 = v56 - int32(3)
										v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v121))))
										if base.B2i32(v121 != v71)&base.B2i32(base.Ui32(v123) < base.Ui32(int32(224))) == int32(0) {
											v151 = v123<<(uint(int32(12))%32)&int32(_a_F_r_Step_1c_2_2) | v119
											v155 = int32(3)
										} else {
											v141 = int32(4)
											v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v72-v141))))
											v151 = v123<<(uint(int32(12))%32)&int32(_a_F_r_Step_1c_2_3) | v143&int32(7)<<(uint(int32(18))%32) | v119
											v155 = v141
										}
									}
								}
								if int32(252) < v151 {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v56 - v155
									v180 = int32(0)
									v187 = v180
								} else {
									v157 = v151 - int32(97)
									if v157 < int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v56 - v155
										v180 = int32(0)
										v187 = v180
									} else {
										v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v157)>>(uint(int32(3))%32)))+uint32(_c_F_r_Step_1c_2[0]))))
										if int32(base.Ui32(v163)>>(uint(v157&int32(7))%32))&int32(1) == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v56 - v155
											v180 = int32(0)
											v187 = v180
										} else {
											v187 = v155
										}
									}
								}
							}
							if v187 != 0 {
								v285 = v2
								return v285
							} else {
								v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v189 = v188 - v55
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v189
								switch v27 - int32(1) {
								case 0:
									v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v189 <= v193 {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v189
										v208 = int32(2)
										v210 = int32(0)
										v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v189-v213 < v208 {
											v223 = v210
										} else {
											v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v219 = F_memcmp(m, v216+v189-v208, int32(_a_F_r_Step_1c_2_4), v208)
											mBase = m.M
											if v219 != 0 {
												v223 = v210
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v189 - v208
												v223 = int32(1)
											}
										}
										if v223 == int32(0) {
											v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v235 - v55
											v238 = F_slice_del(m, l0)
											mBase = m.M
											if int32(0) <= v238 {
												v285 = int32(1)
											} else {
												v285 = v238
											}
											return v285
										} else {
											v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v227 < v226 {
												v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v235 - v55
												v238 = F_slice_del(m, l0)
												mBase = m.M
												if int32(0) <= v238 {
													v285 = int32(1)
												} else {
													v285 = v238
												}
												return v285
											} else {
												v231 = F_slice_from_s(m, l0, int32(1), int32(_a_F_r_Step_1c_2_5))
												mBase = m.M
												v232 = m.ExcPending
												if v232 != 0 {
													return int32(0)
												} else {
													if int32(0) <= v231 {
														v285 = int32(1)
													} else {
														v285 = v231
													}
													return v285
												}
											}
										}
									} else {
										v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195+v189-int32(1)))))
										if v199 != int32(110) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v189
											v208 = int32(2)
											v210 = int32(0)
											v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v189-v213 < v208 {
												v223 = v210
											} else {
												v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v219 = F_memcmp(m, v216+v189-v208, int32(_a_F_r_Step_1c_2_4), v208)
												mBase = m.M
												if v219 != 0 {
													v223 = v210
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v189 - v208
													v223 = int32(1)
												}
											}
											if v223 == int32(0) {
												v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v235 - v55
												v238 = F_slice_del(m, l0)
												mBase = m.M
												if int32(0) <= v238 {
													v285 = int32(1)
												} else {
													v285 = v238
												}
												return v285
											} else {
												v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v227 < v226 {
													v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v235 - v55
													v238 = F_slice_del(m, l0)
													mBase = m.M
													if int32(0) <= v238 {
														v285 = int32(1)
													} else {
														v285 = v238
													}
													return v285
												} else {
													v231 = F_slice_from_s(m, l0, int32(1), int32(_a_F_r_Step_1c_2_5))
													mBase = m.M
													v232 = m.ExcPending
													if v232 != 0 {
														return int32(0)
													} else {
														if int32(0) <= v231 {
															v285 = int32(1)
														} else {
															v285 = v231
														}
														return v285
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v189 - int32(1)
											v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v205 < v189 {
												v285 = v2
												return v285
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v189
												v208 = int32(2)
												v210 = int32(0)
												v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v189-v213 < v208 {
													v223 = v210
												} else {
													v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v219 = F_memcmp(m, v216+v189-v208, int32(_a_F_r_Step_1c_2_4), v208)
													mBase = m.M
													if v219 != 0 {
														v223 = v210
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v189 - v208
														v223 = int32(1)
													}
												}
												if v223 == int32(0) {
													v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v235 - v55
													v238 = F_slice_del(m, l0)
													mBase = m.M
													if int32(0) <= v238 {
														v285 = int32(1)
													} else {
														v285 = v238
													}
													return v285
												} else {
													v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													if v227 < v226 {
														v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v235 - v55
														v238 = F_slice_del(m, l0)
														mBase = m.M
														if int32(0) <= v238 {
															v285 = int32(1)
														} else {
															v285 = v238
														}
														return v285
													} else {
														v231 = F_slice_from_s(m, l0, int32(1), int32(_a_F_r_Step_1c_2_5))
														mBase = m.M
														v232 = m.ExcPending
														if v232 != 0 {
															return int32(0)
														} else {
															if int32(0) <= v231 {
																v285 = int32(1)
															} else {
																v285 = v231
															}
															return v285
														}
													}
												}
											}
										}
									}
								case 1:
									v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v189 <= v241 {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v189
										v256 = int32(2)
										v258 = int32(0)
										v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v189-v261 < v256 {
											v271 = v258
										} else {
											v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v267 = F_memcmp(m, v264+v189-v256, int32(_a_F_r_Step_1c_2_6), v256)
											mBase = m.M
											if v267 != 0 {
												v271 = v258
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v189 - v256
												v271 = int32(1)
											}
										}
										if v271 != 0 {
											v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v272 <= v273 {
												v285 = v2
											} else {
												v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v275 - v55
												v278 = F_slice_del(m, l0)
												mBase = m.M
												if v278 < int32(0) {
													v285 = v278
												} else {
													v285 = int32(1)
												}
											}
										} else {
											v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v275 - v55
											v278 = F_slice_del(m, l0)
											mBase = m.M
											if v278 < int32(0) {
												v285 = v278
											} else {
												v285 = int32(1)
											}
										}
									} else {
										v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243+v189-int32(1)))))
										if v247 != int32(104) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v189
											v256 = int32(2)
											v258 = int32(0)
											v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v189-v261 < v256 {
												v271 = v258
											} else {
												v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v267 = F_memcmp(m, v264+v189-v256, int32(_a_F_r_Step_1c_2_6), v256)
												mBase = m.M
												if v267 != 0 {
													v271 = v258
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v189 - v256
													v271 = int32(1)
												}
											}
											if v271 != 0 {
												v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v272 <= v273 {
													v285 = v2
												} else {
													v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v275 - v55
													v278 = F_slice_del(m, l0)
													mBase = m.M
													if v278 < int32(0) {
														v285 = v278
													} else {
														v285 = int32(1)
													}
												}
											} else {
												v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v275 - v55
												v278 = F_slice_del(m, l0)
												mBase = m.M
												if v278 < int32(0) {
													v285 = v278
												} else {
													v285 = int32(1)
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v189 - int32(1)
											v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v253 < v189 {
												v285 = v2
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v189
												v256 = int32(2)
												v258 = int32(0)
												v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v189-v261 < v256 {
													v271 = v258
												} else {
													v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v267 = F_memcmp(m, v264+v189-v256, int32(_a_F_r_Step_1c_2_6), v256)
													mBase = m.M
													if v267 != 0 {
														v271 = v258
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v189 - v256
														v271 = int32(1)
													}
												}
												if v271 != 0 {
													v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													if v272 <= v273 {
														v285 = v2
													} else {
														v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v275 - v55
														v278 = F_slice_del(m, l0)
														mBase = m.M
														if v278 < int32(0) {
															v285 = v278
														} else {
															v285 = int32(1)
														}
													}
												} else {
													v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v275 - v55
													v278 = F_slice_del(m, l0)
													mBase = m.M
													if v278 < int32(0) {
														v285 = v278
													} else {
														v285 = int32(1)
													}
												}
											}
										}
									}
									return v285
								default:
									v285 = int32(1)
									return v285
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_r_consonant_pair_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v7 < v8 {
		v67 = v2
		return v67
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v7
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
		v15 = v7 - int32(1)
		if v8 < v15 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v15))))
			v21 = v19 - int32(100)
			if base.B2i32(v21 == int32(0))|base.B2i32(v21 == int32(16)) != 0 {
				v34 = F_find_among_b(m, l0, int32(_a_F_r_consonant_pair_1_0), int32(4), int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					if v34 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v44
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v48 = v46 + (v7 - v10)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48
						if v48 <= v12 {
							v67 = v2
						} else {
							v51 = int32(1)
							v52 = v48 - v51
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v52
							v56 = F_slice_del(m, l0)
							mBase = m.M
							if int32(0) <= v56 {
								v62 = v51
							} else {
								v62 = v56 >> (uint(int32(31)) % 32) & v56
							}
							v67 = v62
						}
						return v67
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
				return int32(0)
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
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
	var v8 int32
	_ = v8
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v5 < v6 {
		v170 = v2
		return v170
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v21 <= v22 {
			v131 = int32(-1)
			v138 = v131
		} else {
			v39 = int32(1)
			v40 = v21 - v39
			v42 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23+v40))))
			v44 = v42 & int32(255)
			if base.B2i32(v40 == v22)|base.B2i32(int32(0) <= v42) != 0 {
				v102 = v44
				v106 = v39
			} else {
				v51 = v44 & int32(63)
				v53 = v21 - int32(2)
				v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v53))))
				v57 = v55 << (uint(int32(6)) % 32)
				if base.B2i32(v53 != v22)&base.B2i32(base.Ui32(v55) < base.Ui32(int32(192))) == int32(0) {
					v102 = v57&int32(1984) | v51
					v106 = int32(2)
				} else {
					v70 = v57&int32(4032) | v51
					v72 = v21 - int32(3)
					v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v72))))
					if base.B2i32(v72 != v22)&base.B2i32(base.Ui32(v74) < base.Ui32(int32(224))) == int32(0) {
						v102 = v74<<(uint(int32(12))%32)&int32(_a_F_r_en_ending_2_0) | v70
						v106 = int32(3)
					} else {
						v92 = int32(4)
						v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v23-v92))))
						v102 = v74<<(uint(int32(12))%32)&int32(_a_F_r_en_ending_2_1) | v94&int32(7)<<(uint(int32(18))%32) | v70
						v106 = v92
					}
				}
			}
			if int32(232) < v102 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v21 - v106
				v131 = int32(0)
				v138 = v131
			} else {
				v108 = v102 - int32(97)
				if v108 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v21 - v106
					v131 = int32(0)
					v138 = v131
				} else {
					v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v108)>>(uint(int32(3))%32)))+uint32(_c_F_r_en_ending_2[0]))))
					if int32(base.Ui32(v114)>>(uint(v108&int32(7))%32))&int32(1) == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v21 - v106
						v131 = int32(0)
						v138 = v131
					} else {
						v138 = v106
					}
				}
			}
		}
		if v138 != 0 {
			v170 = v2
			return v170
		} else {
			v139 = v5 - v8
			v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v141 = v139 + v140
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v141
			v143 = int32(3)
			v145 = int32(0)
			v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v141-v148 < v143 {
				v158 = v145
			} else {
				v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v154 = F_memcmp(m, v151+v141-v143, int32(_a_F_r_en_ending_2_2), v143)
				mBase = m.M
				if v154 != 0 {
					v158 = v145
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v141 - v143
					v158 = int32(1)
				}
			}
			if v158 != 0 {
				v170 = v2
				return v170
			} else {
				v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v159 + v139
				v162 = F_slice_del(m, l0)
				mBase = m.M
				if v162 < int32(0) {
					v170 = v162
					return v170
				} else {
					v165 = F_r_undouble_3(m, l0)
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return int32(0)
					} else {
						v170 = v165
						return v170
					}
				}
			}
		}
	}
}
func F_r_remove_second_order_prefix_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v160 int32
	_ = v160
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v5
	v8 = v5 + int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v9 <= v8 {
		v245 = v2
		return v245
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v8))))
		if v13 != int32(101) {
			v245 = v2
			return v245
		} else {
			v16 = int32(2)
			v20 = F_find_among(m, l0, int32(_a_F_r_remove_second_order_prefix_2_0), v16, int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				switch v20 {
				case 0:
					v245 = v20
				case 1:
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v24 == v25 {
						v223 = v24
						v224 = v16
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v223
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v223
						v229 = v224
						*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v229
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v24))))
						switch v29 - int32(108) {
						case 0:
							v33 = v24 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v33
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v33
							v36 = int32(4)
							v38 = int32(0)
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v40-v33 < v36 {
								v50 = v38
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v46 = F_memcmp(m, v44+v33, int32(_a_F_r_remove_second_order_prefix_2_1), v36)
								mBase = m.M
								if v46 != 0 {
									v50 = v38
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v24 + int32(5)
									v50 = int32(1)
								}
							}
							if v50 == int32(0) {
								v223 = v24
								v224 = v16
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v223
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v223
								v229 = v224
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v229
							} else {
							}
						default:
							v223 = v24
							v224 = v16
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v223
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v223
							v229 = v224
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v229
						case 6:
							v223 = v24 + int32(1)
							v224 = v16
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v223
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v223
							v229 = v224
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v229
						}
					}
					v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v236 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v235 - v236
					v240 = F_slice_del(m, l0)
					mBase = m.M
					if int32(0) <= v240 {
						v243 = v236
					} else {
						v243 = v240
					}
					v245 = v243
				case 2:
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v53 == v54 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v53
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v53
						v87 = int32(0)
						v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						if v100 <= v53 {
							v204 = int32(-1)
						} else {
							v116 = int32(1)
							v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v101))))
							if base.Ui32(v118) < base.Ui32(int32(192)) {
								v175 = v118
								v176 = v116
							} else {
								v122 = v53 + int32(1)
								if v122 == v100 {
									v175 = v118
									v176 = v116
								} else {
									v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v101))))
									v127 = v125 & int32(63)
									if base.Ui32(int32(224)) <= base.Ui32(v118) {
										v131 = v53 + int32(2)
										if v131 != v100 {
											v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v101))))
											v143 = v141 & int32(63)
											if base.Ui32(int32(240)) <= base.Ui32(v118) {
												v147 = v53 + int32(3)
												if v147 != v100 {
													v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+v147))))
													v175 = v160&int32(63) | (v118<<(uint(int32(18))%32)&int32(_a_F_r_remove_second_order_prefix_2_2) | v127<<(uint(int32(12))%32) | v143<<(uint(int32(6))%32))
													v176 = int32(4)
												} else {
													v175 = v118<<(uint(int32(12))%32)&int32(_a_F_r_remove_second_order_prefix_2_3) | v127<<(uint(int32(6))%32) | v143
													v176 = int32(3)
												}
											} else {
												v175 = v118<<(uint(int32(12))%32)&int32(_a_F_r_remove_second_order_prefix_2_3) | v127<<(uint(int32(6))%32) | v143
												v176 = int32(3)
											}
										} else {
											v175 = v118<<(uint(int32(6))%32)&int32(1984) | v127
											v176 = int32(2)
										}
									} else {
										v175 = v118<<(uint(int32(6))%32)&int32(1984) | v127
										v176 = int32(2)
									}
								}
							}
							if int32(117) < v175 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v176 + v53
								v197 = int32(0)
							} else {
								v180 = v175 - int32(97)
								if v180 < int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v176 + v53
									v197 = int32(0)
								} else {
									v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v180)>>(uint(int32(3))%32)))+uint32(_c_F_r_remove_second_order_prefix_2[0]))))
									if int32(base.Ui32(v186)>>(uint(v180&int32(7))%32))&int32(1) != 0 {
										v197 = v176
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v176 + v53
										v197 = int32(0)
									}
								}
							}
							v204 = v197
						}
						if v204 != 0 {
							v245 = v87
						} else {
							v206 = int32(2)
							v208 = int32(0)
							v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v210-v211 < v206 {
								v220 = v208
							} else {
								v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v216 = F_memcmp(m, v214+v211, int32(_a_F_r_remove_second_order_prefix_2_4), v206)
								mBase = m.M
								if v216 != 0 {
									v220 = v208
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v206 + v211
									v220 = int32(1)
								}
							}
							if v220 != 0 {
								v229 = int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v229
								v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								v236 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v235 - v236
								v240 = F_slice_del(m, l0)
								mBase = m.M
								if int32(0) <= v240 {
									v243 = v236
								} else {
									v243 = v240
								}
								v245 = v243
							} else {
								v245 = v87
							}
						}
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v53))))
						switch v58 - int32(108) {
						case 0:
							v65 = v53 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v65
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v65
							v68 = int32(4)
							v71 = int32(0)
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v73-v65 < v68 {
								v83 = v71
							} else {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v79 = F_memcmp(m, v77+v65, int32(_a_F_r_remove_second_order_prefix_2_5), v68)
								mBase = m.M
								if v79 != 0 {
									v83 = v71
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v53 + int32(5)
									v83 = int32(1)
								}
							}
							if v83 != 0 {
								v229 = v68
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v229
								v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								v236 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v235 - v236
								v240 = F_slice_del(m, l0)
								mBase = m.M
								if int32(0) <= v240 {
									v243 = v236
								} else {
									v243 = v240
								}
								v245 = v243
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v53
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v53
								v87 = int32(0)
								v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if v100 <= v53 {
									v204 = int32(-1)
								} else {
									v116 = int32(1)
									v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v101))))
									if base.Ui32(v118) < base.Ui32(int32(192)) {
										v175 = v118
										v176 = v116
									} else {
										v122 = v53 + int32(1)
										if v122 == v100 {
											v175 = v118
											v176 = v116
										} else {
											v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v101))))
											v127 = v125 & int32(63)
											if base.Ui32(int32(224)) <= base.Ui32(v118) {
												v131 = v53 + int32(2)
												if v131 != v100 {
													v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v101))))
													v143 = v141 & int32(63)
													if base.Ui32(int32(240)) <= base.Ui32(v118) {
														v147 = v53 + int32(3)
														if v147 != v100 {
															v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+v147))))
															v175 = v160&int32(63) | (v118<<(uint(int32(18))%32)&int32(_a_F_r_remove_second_order_prefix_2_2) | v127<<(uint(int32(12))%32) | v143<<(uint(int32(6))%32))
															v176 = int32(4)
														} else {
															v175 = v118<<(uint(int32(12))%32)&int32(_a_F_r_remove_second_order_prefix_2_3) | v127<<(uint(int32(6))%32) | v143
															v176 = int32(3)
														}
													} else {
														v175 = v118<<(uint(int32(12))%32)&int32(_a_F_r_remove_second_order_prefix_2_3) | v127<<(uint(int32(6))%32) | v143
														v176 = int32(3)
													}
												} else {
													v175 = v118<<(uint(int32(6))%32)&int32(1984) | v127
													v176 = int32(2)
												}
											} else {
												v175 = v118<<(uint(int32(6))%32)&int32(1984) | v127
												v176 = int32(2)
											}
										}
									}
									if int32(117) < v175 {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v176 + v53
										v197 = int32(0)
									} else {
										v180 = v175 - int32(97)
										if v180 < int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v176 + v53
											v197 = int32(0)
										} else {
											v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v180)>>(uint(int32(3))%32)))+uint32(_c_F_r_remove_second_order_prefix_2[0]))))
											if int32(base.Ui32(v186)>>(uint(v180&int32(7))%32))&int32(1) != 0 {
												v197 = v176
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v176 + v53
												v197 = int32(0)
											}
										}
									}
									v204 = v197
								}
								if v204 != 0 {
									v245 = v87
								} else {
									v206 = int32(2)
									v208 = int32(0)
									v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v210-v211 < v206 {
										v220 = v208
									} else {
										v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v216 = F_memcmp(m, v214+v211, int32(_a_F_r_remove_second_order_prefix_2_4), v206)
										mBase = m.M
										if v216 != 0 {
											v220 = v208
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v206 + v211
											v220 = int32(1)
										}
									}
									if v220 != 0 {
										v229 = int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v229
										v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										v236 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v235 - v236
										v240 = F_slice_del(m, l0)
										mBase = m.M
										if int32(0) <= v240 {
											v243 = v236
										} else {
											v243 = v240
										}
										v245 = v243
									} else {
										v245 = v87
									}
								}
							}
						default:
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v53
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v53
							v87 = int32(0)
							v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							if v100 <= v53 {
								v204 = int32(-1)
							} else {
								v116 = int32(1)
								v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v101))))
								if base.Ui32(v118) < base.Ui32(int32(192)) {
									v175 = v118
									v176 = v116
								} else {
									v122 = v53 + int32(1)
									if v122 == v100 {
										v175 = v118
										v176 = v116
									} else {
										v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v101))))
										v127 = v125 & int32(63)
										if base.Ui32(int32(224)) <= base.Ui32(v118) {
											v131 = v53 + int32(2)
											if v131 != v100 {
												v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v101))))
												v143 = v141 & int32(63)
												if base.Ui32(int32(240)) <= base.Ui32(v118) {
													v147 = v53 + int32(3)
													if v147 != v100 {
														v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+v147))))
														v175 = v160&int32(63) | (v118<<(uint(int32(18))%32)&int32(_a_F_r_remove_second_order_prefix_2_2) | v127<<(uint(int32(12))%32) | v143<<(uint(int32(6))%32))
														v176 = int32(4)
													} else {
														v175 = v118<<(uint(int32(12))%32)&int32(_a_F_r_remove_second_order_prefix_2_3) | v127<<(uint(int32(6))%32) | v143
														v176 = int32(3)
													}
												} else {
													v175 = v118<<(uint(int32(12))%32)&int32(_a_F_r_remove_second_order_prefix_2_3) | v127<<(uint(int32(6))%32) | v143
													v176 = int32(3)
												}
											} else {
												v175 = v118<<(uint(int32(6))%32)&int32(1984) | v127
												v176 = int32(2)
											}
										} else {
											v175 = v118<<(uint(int32(6))%32)&int32(1984) | v127
											v176 = int32(2)
										}
									}
								}
								if int32(117) < v175 {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v176 + v53
									v197 = int32(0)
								} else {
									v180 = v175 - int32(97)
									if v180 < int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v176 + v53
										v197 = int32(0)
									} else {
										v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v180)>>(uint(int32(3))%32)))+uint32(_c_F_r_remove_second_order_prefix_2[0]))))
										if int32(base.Ui32(v186)>>(uint(v180&int32(7))%32))&int32(1) != 0 {
											v197 = v176
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v176 + v53
											v197 = int32(0)
										}
									}
								}
								v204 = v197
							}
							if v204 != 0 {
								v245 = v87
							} else {
								v206 = int32(2)
								v208 = int32(0)
								v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v210-v211 < v206 {
									v220 = v208
								} else {
									v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v216 = F_memcmp(m, v214+v211, int32(_a_F_r_remove_second_order_prefix_2_4), v206)
									mBase = m.M
									if v216 != 0 {
										v220 = v208
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v206 + v211
										v220 = int32(1)
									}
								}
								if v220 != 0 {
									v229 = int32(4)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v229
									v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									v236 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v235 - v236
									v240 = F_slice_del(m, l0)
									mBase = m.M
									if int32(0) <= v240 {
										v243 = v236
									} else {
										v243 = v240
									}
									v245 = v243
								} else {
									v245 = v87
								}
							}
						case 6:
							v223 = v53 + int32(1)
							v224 = int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v223
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v223
							v229 = v224
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v229
							v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v236 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v235 - v236
							v240 = F_slice_del(m, l0)
							mBase = m.M
							if int32(0) <= v240 {
								v243 = v236
							} else {
								v243 = v240
							}
							v245 = v243
						}
					}
				default:
					v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v236 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v235 - v236
					v240 = F_slice_del(m, l0)
					mBase = m.M
					if int32(0) <= v240 {
						v243 = v236
					} else {
						v243 = v240
					}
					v245 = v243
				}
				return v245
			}
		}
	}
}
func F_r_shortv_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v171 int32
	_ = v171
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v4 <= v14 {
		v54 = int32(-1)
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v4-int32(1)))))
		if int32(121) < v29 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4 - int32(1)
			v51 = int32(0)
		} else {
			v31 = v29 - int32(89)
			if v31 < int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4 - int32(1)
				v51 = int32(0)
			} else {
				v34 = int32(1)
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v31)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_1[0]))))
				if int32(base.Ui32(v38)>>(uint(v31&int32(7))%32))&v34 != 0 {
					v51 = v34
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4 - int32(1)
					v51 = int32(0)
				}
			}
		}
		v54 = v51
	}
	if v54 != 0 {
		v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v160 = v5 - v4
		v161 = v159 - v160
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v161
		v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v161 <= v171 {
			v211 = int32(-1)
		} else {
			v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182+v161-int32(1)))))
			if int32(121) < v186 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v161 - int32(1)
				v208 = int32(0)
			} else {
				v188 = v186 - int32(97)
				if v188 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v161 - int32(1)
					v208 = int32(0)
				} else {
					v191 = int32(1)
					v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v188)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_1[1]))))
					if int32(base.Ui32(v195)>>(uint(v188&int32(7))%32))&v191 != 0 {
						v208 = v191
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v161 - int32(1)
						v208 = int32(0)
					}
				}
			}
			v211 = v208
		}
		if v211 != 0 {
			v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v269 = v268 - v160
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269
			v271 = int32(4)
			v273 = int32(0)
			v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v269-v276 < v271 {
				v286 = v273
			} else {
				v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v282 = F_memcmp(m, v279+v269-v271, int32(_a_F_r_shortv_1_0), v271)
				mBase = m.M
				if v282 != 0 {
					v286 = v273
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269 - v271
					v286 = int32(1)
				}
			}
			if v286 != 0 {
				return int32(1)
			} else {
				return int32(0)
			}
		} else {
			v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v220 <= v221 {
				v264 = int32(-1)
			} else {
				v233 = int32(1)
				v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234+v220-v233))))
				if int32(121) < v238 {
					v260 = v233
				} else {
					v240 = v238 - int32(97)
					if v240 < int32(0) {
						v260 = v233
					} else {
						v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v240)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_1[1]))))
						if int32(base.Ui32(v246)>>(uint(v240&int32(7))%32))&int32(1) == int32(0) {
							v260 = v233
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v220 - int32(1)
							v260 = int32(0)
						}
					}
				}
				v264 = v260
			}
			if v264 != 0 {
				v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v269 = v268 - v160
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269
				v271 = int32(4)
				v273 = int32(0)
				v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v269-v276 < v271 {
					v286 = v273
				} else {
					v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v282 = F_memcmp(m, v279+v269-v271, int32(_a_F_r_shortv_1_0), v271)
					mBase = m.M
					if v282 != 0 {
						v286 = v273
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269 - v271
						v286 = int32(1)
					}
				}
				if v286 != 0 {
					return int32(1)
				} else {
					return int32(0)
				}
			} else {
				v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v265 <= v266 {
					return int32(1)
				} else {
					v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v269 = v268 - v160
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269
					v271 = int32(4)
					v273 = int32(0)
					v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v269-v276 < v271 {
						v286 = v273
					} else {
						v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v282 = F_memcmp(m, v279+v269-v271, int32(_a_F_r_shortv_1_0), v271)
						mBase = m.M
						if v282 != 0 {
							v286 = v273
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269 - v271
							v286 = int32(1)
						}
					}
					if v286 != 0 {
						return int32(1)
					} else {
						return int32(0)
					}
				}
			}
		}
	} else {
		v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v63 <= v64 {
			v107 = int32(-1)
		} else {
			v76 = int32(1)
			v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v63-v76))))
			if int32(121) < v81 {
				v103 = v76
			} else {
				v83 = v81 - int32(97)
				if v83 < int32(0) {
					v103 = v76
				} else {
					v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v83)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_1[1]))))
					if int32(base.Ui32(v89)>>(uint(v83&int32(7))%32))&int32(1) == int32(0) {
						v103 = v76
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v63 - int32(1)
						v103 = int32(0)
					}
				}
			}
			v107 = v103
		}
		if v107 != 0 {
			v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v160 = v5 - v4
			v161 = v159 - v160
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v161
			v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v161 <= v171 {
				v211 = int32(-1)
			} else {
				v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182+v161-int32(1)))))
				if int32(121) < v186 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v161 - int32(1)
					v208 = int32(0)
				} else {
					v188 = v186 - int32(97)
					if v188 < int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v161 - int32(1)
						v208 = int32(0)
					} else {
						v191 = int32(1)
						v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v188)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_1[1]))))
						if int32(base.Ui32(v195)>>(uint(v188&int32(7))%32))&v191 != 0 {
							v208 = v191
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v161 - int32(1)
							v208 = int32(0)
						}
					}
				}
				v211 = v208
			}
			if v211 != 0 {
				v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v269 = v268 - v160
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269
				v271 = int32(4)
				v273 = int32(0)
				v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v269-v276 < v271 {
					v286 = v273
				} else {
					v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v282 = F_memcmp(m, v279+v269-v271, int32(_a_F_r_shortv_1_0), v271)
					mBase = m.M
					if v282 != 0 {
						v286 = v273
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269 - v271
						v286 = int32(1)
					}
				}
				if v286 != 0 {
					return int32(1)
				} else {
					return int32(0)
				}
			} else {
				v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v220 <= v221 {
					v264 = int32(-1)
				} else {
					v233 = int32(1)
					v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234+v220-v233))))
					if int32(121) < v238 {
						v260 = v233
					} else {
						v240 = v238 - int32(97)
						if v240 < int32(0) {
							v260 = v233
						} else {
							v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v240)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_1[1]))))
							if int32(base.Ui32(v246)>>(uint(v240&int32(7))%32))&int32(1) == int32(0) {
								v260 = v233
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v220 - int32(1)
								v260 = int32(0)
							}
						}
					}
					v264 = v260
				}
				if v264 != 0 {
					v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v269 = v268 - v160
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269
					v271 = int32(4)
					v273 = int32(0)
					v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v269-v276 < v271 {
						v286 = v273
					} else {
						v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v282 = F_memcmp(m, v279+v269-v271, int32(_a_F_r_shortv_1_0), v271)
						mBase = m.M
						if v282 != 0 {
							v286 = v273
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269 - v271
							v286 = int32(1)
						}
					}
					if v286 != 0 {
						return int32(1)
					} else {
						return int32(0)
					}
				} else {
					v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v265 <= v266 {
						return int32(1)
					} else {
						v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v269 = v268 - v160
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269
						v271 = int32(4)
						v273 = int32(0)
						v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v269-v276 < v271 {
							v286 = v273
						} else {
							v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v282 = F_memcmp(m, v279+v269-v271, int32(_a_F_r_shortv_1_0), v271)
							mBase = m.M
							if v282 != 0 {
								v286 = v273
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269 - v271
								v286 = int32(1)
							}
						}
						if v286 != 0 {
							return int32(1)
						} else {
							return int32(0)
						}
					}
				}
			}
		} else {
			v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v115 <= v116 {
				v156 = int32(-1)
			} else {
				v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127+v115-int32(1)))))
				if int32(121) < v131 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v115 - int32(1)
					v153 = int32(0)
				} else {
					v133 = v131 - int32(97)
					if v133 < int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v115 - int32(1)
						v153 = int32(0)
					} else {
						v136 = int32(1)
						v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v133)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_1[1]))))
						if int32(base.Ui32(v140)>>(uint(v133&int32(7))%32))&v136 != 0 {
							v153 = v136
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v115 - int32(1)
							v153 = int32(0)
						}
					}
				}
				v156 = v153
			}
			if v156 == int32(0) {
				return int32(1)
			} else {
				v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v160 = v5 - v4
				v161 = v159 - v160
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v161
				v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v161 <= v171 {
					v211 = int32(-1)
				} else {
					v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182+v161-int32(1)))))
					if int32(121) < v186 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v161 - int32(1)
						v208 = int32(0)
					} else {
						v188 = v186 - int32(97)
						if v188 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v161 - int32(1)
							v208 = int32(0)
						} else {
							v191 = int32(1)
							v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v188)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_1[1]))))
							if int32(base.Ui32(v195)>>(uint(v188&int32(7))%32))&v191 != 0 {
								v208 = v191
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v161 - int32(1)
								v208 = int32(0)
							}
						}
					}
					v211 = v208
				}
				if v211 != 0 {
					v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v269 = v268 - v160
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269
					v271 = int32(4)
					v273 = int32(0)
					v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v269-v276 < v271 {
						v286 = v273
					} else {
						v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v282 = F_memcmp(m, v279+v269-v271, int32(_a_F_r_shortv_1_0), v271)
						mBase = m.M
						if v282 != 0 {
							v286 = v273
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269 - v271
							v286 = int32(1)
						}
					}
					if v286 != 0 {
						return int32(1)
					} else {
						return int32(0)
					}
				} else {
					v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v220 <= v221 {
						v264 = int32(-1)
					} else {
						v233 = int32(1)
						v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234+v220-v233))))
						if int32(121) < v238 {
							v260 = v233
						} else {
							v240 = v238 - int32(97)
							if v240 < int32(0) {
								v260 = v233
							} else {
								v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v240)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_1[1]))))
								if int32(base.Ui32(v246)>>(uint(v240&int32(7))%32))&int32(1) == int32(0) {
									v260 = v233
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v220 - int32(1)
									v260 = int32(0)
								}
							}
						}
						v264 = v260
					}
					if v264 != 0 {
						v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v269 = v268 - v160
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269
						v271 = int32(4)
						v273 = int32(0)
						v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v269-v276 < v271 {
							v286 = v273
						} else {
							v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v282 = F_memcmp(m, v279+v269-v271, int32(_a_F_r_shortv_1_0), v271)
							mBase = m.M
							if v282 != 0 {
								v286 = v273
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269 - v271
								v286 = int32(1)
							}
						}
						if v286 != 0 {
							return int32(1)
						} else {
							return int32(0)
						}
					} else {
						v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v265 <= v266 {
							return int32(1)
						} else {
							v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v269 = v268 - v160
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269
							v271 = int32(4)
							v273 = int32(0)
							v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v269-v276 < v271 {
								v286 = v273
							} else {
								v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v282 = F_memcmp(m, v279+v269-v271, int32(_a_F_r_shortv_1_0), v271)
								mBase = m.M
								if v282 != 0 {
									v286 = v273
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v269 - v271
									v286 = int32(1)
								}
							}
							if v286 != 0 {
								return int32(1)
							} else {
								return int32(0)
							}
						}
					}
				}
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
			v28 = F_find_among_b(m, l0, int32(_a_F_r_undouble_1_1), int32(3), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				if v28 == int32(0) {
					v56 = v2
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v36 = v34 + (v5 - v24)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v36
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v36
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v36 <= v39 {
						v56 = v2
					} else {
						v41 = int32(1)
						v42 = v36 - v41
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v42
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v42
						v46 = F_slice_del(m, l0)
						mBase = m.M
						if int32(0) <= v46 {
							v52 = v41
						} else {
							v52 = v46 >> (uint(int32(31)) % 32) & v46
						}
						v56 = v52
					}
				}
				return v56
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
	var v33 int32
	_ = v33
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
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v373 int32
	_ = v373
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
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
	var v492 int32
	_ = v492
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
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L107
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_read_cursor_args_1))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L102
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_read_cursor_args_1))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L97
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_read_cursor_args_1))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L92
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_read_cursor_args_1))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L87
	}
L8:
	;
	m.G0 = v16 + int32(128)
	return v373
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
		v373 = int32(0)
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L3
L14:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_read_cursor_args[0]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v22<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	v39 = F_palloc0_mul(m, int32(4), v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	if int32(0) < v41 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v52 = v6
	v54 = v6
	goto L19
L17:
	;
	v272 = v6
	goto L18
L18:
	;
	F_initStringInfo(m, v16+int32(112))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L65
	}
L19:
	;
	F_plpgsql_peek2(m, v16+int32(108), v16+int32(104), v16+int32(100), int32(0), l4)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v272 = v206
	goto L18
L21:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v16)+108))
	if v66 != int32(258) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v211 = v39 + v202<<(uint(int32(2))%32)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	if v212 != 0 {
		goto L7
	} else {
		goto L52
	}
L23:
	;
	v202 = v52
	v206 = v54
	goto L22
L24:
	;
	goto L25
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v16)+104))
	if v69&int32(-2) != int32(270) {
		v202 = v52
		v206 = v54
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v74 = int32(_a_F_read_cursor_args_2)
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_read_cursor_args[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_read_cursor_args[1])) = int32(1)
	v79 = F_plpgsql_yylex(m, l2, l3, l4)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, _c_F_read_cursor_args[1])) = v75
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	if v84 <= int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_read_cursor_args_1))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L47
	}
L29:
	;
	if v144 == v84 {
		goto L28
	} else {
		goto L44
	}
L30:
	;
	v144 = int32(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	v96 = int32(0)
	goto L33
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v88+v96<<(uint(int32(2))%32))))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if base.B2i32(v109 == int32(0))|base.B2i32(v109 != v112) != 0 {
		v130 = v109
		v131 = v112
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L28
L35:
	;
	if v130-v131 == int32(0) {
		v144 = v96
		goto L29
	} else {
		goto L42
	}
L36:
	;
	goto L35
L37:
	;
	v115 = v106
	v116 = v81
	goto L38
L38:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
	if v120 == int32(0) {
		v130 = v120
		v131 = v119
		goto L36
	} else {
		goto L40
	}
L39:
	;
	v130 = v120
	v131 = v119
	goto L36
L40:
	;
	v123 = int32(1)
	if v120 == v119 {
		v115 = v115 + v123
		v116 = v116 + v123
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v136 = v96 + int32(1)
	if v136 != v84 {
		v96 = v136
		goto L33
	} else {
		goto L43
	}
L43:
	;
	goto L34
L44:
	;
	v152 = F_plpgsql_yylex(m, l2, l3, l4)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+104)) = v152
	if base.Ui32(int32(-3)) < base.Ui32(v152-int32(272)) {
		v202 = v144
		v206 = int32(1)
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
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v180
	F_errmsg(m, int32(_a_F_read_cursor_args_3), v16+int32(16))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v189 = F_plpgsql_scanner_errposition(m, v188, l4)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_read_cursor_args_4), int32(3997), int32(_a_F_read_cursor_args_5))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
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
	v215 = int32(0)
	v218 = int32(1)
	v223 = F_read_sql_construct(m, int32(44), int32(41), v215, int32(_a_F_read_cursor_args_6), int32(2), v218, v218, v215, v16+int32(112), l2, l3, l4)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	*(*int32)(unsafe.Add(mBase, uint32(v211))) = v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v16)+112))
	switch v228 - int32(41) {
	case 0:
		goto L56
	default:
		goto L54
	case 3:
		goto L55
	}
L54:
	;
	v260 = v52 + int32(1)
	if v260 < v227 {
		v52 = v260
		v54 = v206
		goto L19
	} else {
		goto L64
	}
L55:
	;
	if v52 == v227-int32(1) {
		goto L6
	} else {
		goto L63
	}
L56:
	;
	if v52 == v227-int32(1) {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_read_cursor_args_1))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v241
	F_errmsg(m, int32(_a_F_read_cursor_args_7), v16+int32(48))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v249 = F_plpgsql_scanner_errposition(m, v248, l4)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_read_cursor_args_4), int32(4039), int32(_a_F_read_cursor_args_5))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
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
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	if int32(0) < v279 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v289 = int32(0)
	goto L69
L67:
	;
	goto L68
L68:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v16)+112))
	v344 = F_palloc0(m, int32(80))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L82
	}
L69:
	;
	v297 = v16 + int32(112)
	v299 = v289 << (uint(int32(2)) % 32)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v39+v299)))
	F_appendStringInfoString(m, v297, v301)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L71
	}
L70:
	;
	goto L68
L71:
	;
	if v272 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v304+v299)))
	v307 = F_quote_identifier(m, v306)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	if v289 < v315-int32(1) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v307
	F_appendStringInfo(m, v297, int32(_a_F_read_cursor_args_8), v16+int32(32))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
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
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	v325 = v315
	goto L79
L79:
	;
	v327 = v289 + int32(1)
	if v327 < v325 {
		v289 = v327
		goto L69
	} else {
		goto L81
	}
L80:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	v325 = v324
	goto L79
L81:
	;
	goto L70
L82:
	;
	v346 = F_pstrdup(m, v342)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v344)+4)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v344))) = v346
	v352 = *(*int32)(unsafe.Add(mBase, _c_F_read_cursor_args[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v344)+8)) = v352
	v355 = *(*int32)(unsafe.Add(mBase, _c_F_read_cursor_args[3]))
	v356 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v344)+20)) = uint8(v356)
	*(*int32)(unsafe.Add(mBase, uint32(v344)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v344)+12)) = v355
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v16)+112))
	F_pfree(m, v361)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v364 = F_plpgsql_yylex(m, l2, l3, l4)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v364 != l1 {
		goto L3
	} else {
		goto L86
	}
L86:
	;
	v373 = v344
	goto L8
L87:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v391+v202<<(uint(int32(2))%32))))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v395
	F_errmsg(m, int32(_a_F_read_cursor_args_10), v16+int32(80))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
	v405 = F_plpgsql_scanner_errposition(m, v404, l4)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_read_cursor_args_4), int32(4017), int32(_a_F_read_cursor_args_5))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
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
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v419
	F_errmsg(m, int32(_a_F_read_cursor_args_11), v16-int32(-64))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v427 = F_plpgsql_scanner_errposition(m, v426, l4)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_read_cursor_args_4), int32(4046), int32(_a_F_read_cursor_args_5))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
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
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v441
	F_errmsg(m, int32(_a_F_read_cursor_args_12), v16+int32(96))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v449 = F_plpgsql_scanner_errposition(m, v448, l4)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_read_cursor_args_4), int32(3952), int32(_a_F_read_cursor_args_5))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
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
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v463
	F_errmsg(m, int32(_a_F_read_cursor_args_13), v16)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v469 = F_plpgsql_scanner_errposition(m, v468, l4)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_read_cursor_args_4), int32(3938), int32(_a_F_read_cursor_args_5))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
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
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_read_into_target_0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L10
	} else {
		goto L43
	}
L2:
	;
	F_plpgsql_yyerror(m, l3, int32(0), l4, int32(_a_F_read_into_target_1))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L10
	} else {
		goto L42
	}
L3:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_cword_is_not_variable(m, l2, v93, l4)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L10
	} else {
		goto L41
	}
L4:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_word_is_not_variable(m, l2, v90, l4)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L10
	} else {
		goto L40
	}
L5:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v62 = int32(1)
	if base.Ui32(v61-v62) <= base.Ui32(v62) {
		goto L28
	} else {
		goto L29
	}
L6:
	;
	if v54 == int32(277) {
		goto L5
	} else {
		goto L27
	}
L7:
	;
	v9 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v9)
	v11 = F_plpgsql_yylex(m, l2, l3, l4)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v52 = F_plpgsql_yylex(m, l2, l3, l4)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L10
	} else {
		goto L26
	}
L10:
	;
	return
L11:
	;
	if v11 != int32(373) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v11 != int32(277) {
		v54 = v11
		goto L6
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v49 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v49)
	goto L9
L15:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v17 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v18 == int32(0) {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v21 = int32(_a_F_read_into_target_2)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_read_into_target[0])))
	if base.B2i32(v24 == int32(0))|base.B2i32(v24 != v27) != 0 {
		v45 = v24
		v46 = v27
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v45-v46 != 0 {
		goto L5
	} else {
		goto L25
	}
L19:
	;
	goto L18
L20:
	;
	v30 = v18
	v31 = v21
	goto L21
L21:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	if v35 == int32(0) {
		v45 = v35
		v46 = v34
		goto L19
	} else {
		goto L23
	}
L22:
	;
	v45 = v35
	v46 = v34
	goto L19
L23:
	;
	v38 = int32(1)
	if v35 == v34 {
		v30 = v30 + v38
		v31 = v31 + v38
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	goto L14
L26:
	;
	v54 = v52
	goto L6
L27:
	;
	switch v54 - int32(275) {
	case 0:
		goto L4
	case 1:
		goto L3
	default:
		goto L2
	}
L28:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_check_assignable(m, v60, v66, l4)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L10
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v77 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v69
	v71 = F_plpgsql_yylex(m, l2, l3, l4)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	if v71 == int32(44) {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_plpgsql_push_back_token(m, v71, l2, l3, l4)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	return
L35:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v81 = F_NameListToString(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L10
	} else {
		goto L38
	}
L36:
	;
	v84 = v60
	v85 = v77
	goto L37
L37:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v87 = F_read_into_scalar_list(m, v85, v84, v86, l2, l3, l4)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L10
	} else {
		goto L39
	}
L38:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v84 = v83
	v85 = v81
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v87
	return
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(_a_F_read_into_target_3), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v112 = F_plpgsql_scanner_errposition(m, v111, l4)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_read_into_target_4), int32(3631), int32(_a_F_read_into_target_5))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
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
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = int32(167772202)
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
	F_errfinish(m, int32(_a_F_read_relmap_file_4), int32(820), int32(_a_F_read_relmap_file_5))
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
	v99 = int32(830)
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
	v99 = int32(835)
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
	F_errfinish(m, int32(_a_F_read_relmap_file_4), int32(843), int32(_a_F_read_relmap_file_5))
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
	F_errfinish(m, int32(_a_F_read_relmap_file_4), int32(854), int32(_a_F_read_relmap_file_5))
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
	F_errfinish(m, int32(_a_F_read_relmap_file_4), int32(864), int32(_a_F_read_relmap_file_5))
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = l3 - int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v14
	if v14 != 0 {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
		if v16 == int32(0) {
			v21 = F_LogicalTapeRead(m, l2, l1+int32(8), v14)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				if v21 == v14 {
					v49 = v5
					v50 = v5
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v49
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v50)
					v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
					if v54&int32(1) != 0 {
						v60 = F_LogicalTapeRead(m, l2, v11+int32(12), int32(4))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							if v60 != int32(4) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									F_errmsg_internal(m, int32(_a_F_readtup_datum_0), int32(0))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_readtup_datum_1), int32(2088), int32(_a_F_readtup_datum_2))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								m.G0 = v11 + int32(16)
								return
							}
						}
					} else {
						m.G0 = v11 + int32(16)
						return
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_readtup_datum_0), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_readtup_datum_1), int32(2073), int32(_a_F_readtup_datum_2))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
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
			v37 = F_tuplesort_readtup_alloc(m, l0, v14)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				v39 = F_LogicalTapeRead(m, l2, v37, v14)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					if v39 != v14 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_readtup_datum_0), int32(0))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_readtup_datum_1), int32(2081), int32(_a_F_readtup_datum_2))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v45 = v37
						v46 = base.I64_extend_i32_u(v37)
						v47 = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v46
						v49 = v45
						v50 = v47
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v49
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v50)
						v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
						if v54&int32(1) != 0 {
							v60 = F_LogicalTapeRead(m, l2, v11+int32(12), int32(4))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								if v60 != int32(4) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										F_errmsg_internal(m, int32(_a_F_readtup_datum_0), int32(0))
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_readtup_datum_1), int32(2088), int32(_a_F_readtup_datum_2))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									m.G0 = v11 + int32(16)
									return
								}
							}
						} else {
							m.G0 = v11 + int32(16)
							return
						}
					}
				}
			}
		}
	} else {
		v45 = v5
		v46 = int64(0)
		v47 = int32(1)
		*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v46
		v49 = v45
		v50 = v47
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v49
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v50)
		v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
		if v54&int32(1) != 0 {
			v60 = F_LogicalTapeRead(m, l2, v11+int32(12), int32(4))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return
			} else {
				if v60 != int32(4) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_readtup_datum_0), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_readtup_datum_1), int32(2088), int32(_a_F_readtup_datum_2))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					m.G0 = v11 + int32(16)
					return
				}
			}
		} else {
			m.G0 = v11 + int32(16)
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
							*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v14
							m.G0 = v9 + int32(16)
							return
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = int64(0)
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
func F_ready_file_comparator(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
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
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	v4 = int32(0)
	v6 = base.I32_wrap_i64(l1)
	v7 = base.I32_wrap_i64(l0)
	v8 = F_strlen(m, v7)
	mBase = m.M
	if v8 != int32(16) {
		v128 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v129 = F_strlen(m, v6)
	mBase = m.M
	if v129 == int32(16) {
		goto L33
	} else {
		goto L34
	}
L2:
	;
	v11 = int32(_a_F_ready_file_comparator_0)
	v15 = m.G0
	v17 = v15 - int32(32)
	v18 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v18
	v26 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ready_file_comparator[0])))
	if v26 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v94 != int32(8) {
		v128 = v4
		goto L1
	} else {
		goto L22
	}
L4:
	;
	v94 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ready_file_comparator[1])))
	if v30 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v34 = v7
	goto L10
L8:
	;
	goto L9
L9:
	;
	v44 = v11
	v45 = v26
	goto L13
L10:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v40 == v26 {
		v34 = v34 + int32(1)
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v94 = v34 - v7
	goto L3
L12:
	;
	goto L11
L13:
	;
	v52 = v17 + int32(base.Ui32(v45)>>(uint(int32(3))%32))&int32(28)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v54 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v53 | v54<<(uint(v45)%32)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	if v58 != 0 {
		v44 = v44 + v54
		v45 = v58
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v61 == int32(0) {
		v84 = v7
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v94 = v84 - v7
	goto L3
L17:
	;
	v65 = v7
	v66 = v61
	goto L18
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(base.Ui32(v66)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v74)>>(uint(v66)%32))&int32(1) == int32(0) {
		v84 = v65
		goto L16
	} else {
		goto L20
	}
L19:
	;
	v84 = v82
	goto L16
L20:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
	v82 = v65 + int32(1)
	if v80 != 0 {
		v65 = v82
		v66 = v80
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v98 = v7 + int32(8)
	v99 = int32(_a_F_ready_file_comparator_1)
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ready_file_comparator[2])))
	if base.B2i32(v102 == int32(0))|base.B2i32(v102 != v105) != 0 {
		v123 = v102
		v124 = v105
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v128 = base.B2i32(v123-v124 == int32(0))
	goto L1
L24:
	;
	goto L23
L25:
	;
	v108 = v98
	v109 = v99
	goto L26
L26:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+1)))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
	if v113 == int32(0) {
		v123 = v113
		v124 = v112
		goto L24
	} else {
		goto L28
	}
L27:
	;
	v123 = v113
	v124 = v112
	goto L24
L28:
	;
	v116 = int32(1)
	if v113 == v112 {
		v108 = v108 + v116
		v109 = v109 + v116
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if base.B2i32(v256 == int32(0))|base.B2i32(v256 != v259) != 0 {
		v277 = v256
		v278 = v259
		goto L69
	} else {
		goto L70
	}
L31:
	;
	if v128 != 0 {
		goto L65
	} else {
		goto L66
	}
L32:
	;
	v219 = v6 + int32(8)
	v220 = int32(_a_F_ready_file_comparator_1)
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ready_file_comparator[2])))
	if base.B2i32(v223 == int32(0))|base.B2i32(v223 != v226) != 0 {
		v244 = v223
		v245 = v226
		goto L58
	} else {
		goto L59
	}
L33:
	;
	v132 = int32(_a_F_ready_file_comparator_0)
	v136 = m.G0
	v138 = v136 - int32(32)
	v139 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v138)+24)) = v139
	*(*int64)(unsafe.Add(mBase, uint32(v138)+16)) = v139
	*(*int64)(unsafe.Add(mBase, uint32(v138)+8)) = v139
	*(*int64)(unsafe.Add(mBase, uint32(v138))) = v139
	v147 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ready_file_comparator[0])))
	if v147 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L35
L35:
	;
	if v128 != 0 {
		goto L31
	} else {
		goto L56
	}
L36:
	;
	if v215 == int32(8) {
		goto L32
	} else {
		goto L55
	}
L37:
	;
	v215 = int32(0)
	goto L36
L38:
	;
	goto L39
L39:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ready_file_comparator[1])))
	if v151 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v155 = v6
	goto L43
L41:
	;
	goto L42
L42:
	;
	v165 = v132
	v166 = v147
	goto L46
L43:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	if v161 == v147 {
		v155 = v155 + int32(1)
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v215 = v155 - v6
	goto L36
L45:
	;
	goto L44
L46:
	;
	v173 = v138 + int32(base.Ui32(v166)>>(uint(int32(3))%32))&int32(28)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v175 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = v174 | v175<<(uint(v166)%32)
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
	if v179 != 0 {
		v165 = v165 + v175
		v166 = v179
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v182 == int32(0) {
		v205 = v6
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	v215 = v205 - v6
	goto L36
L50:
	;
	v186 = v6
	v187 = v182
	goto L51
L51:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v138+int32(base.Ui32(v187)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v195)>>(uint(v187)%32))&int32(1) == int32(0) {
		v205 = v186
		goto L49
	} else {
		goto L53
	}
L52:
	;
	v205 = v203
	goto L49
L53:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+1)))
	v203 = v186 + int32(1)
	if v201 != 0 {
		v186 = v203
		v187 = v201
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
	if v128 == base.B2i32(v244-v245 == int32(0)) {
		goto L30
	} else {
		goto L64
	}
L58:
	;
	goto L57
L59:
	;
	v229 = v219
	v230 = v220
	goto L60
L60:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+1)))
	if v234 == int32(0) {
		v244 = v234
		v245 = v233
		goto L58
	} else {
		goto L62
	}
L61:
	;
	v244 = v234
	v245 = v233
	goto L58
L62:
	;
	v237 = int32(1)
	if v234 == v233 {
		v229 = v229 + v237
		v230 = v230 + v237
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
	v252 = int32(-1)
	goto L67
L66:
	;
	v252 = int32(1)
	goto L67
L67:
	;
	return v252
L68:
	;
	return v277 - v278
L69:
	;
	goto L68
L70:
	;
	v262 = v7
	v263 = v6
	goto L71
L71:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+1)))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+1)))
	if v267 == int32(0) {
		v277 = v267
		v278 = v266
		goto L69
	} else {
		goto L73
	}
L72:
	;
	v277 = v267
	v278 = v266
	goto L69
L73:
	;
	v270 = int32(1)
	if v267 == v266 {
		v262 = v262 + v270
		v263 = v263 + v270
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
}
func F_record_gt(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(int32(0) < v2))
	}
}
func F_recurse_set_operations(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v29 float64
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v66 float64
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
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
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int64
	_ = v91
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
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
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
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
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v399 int32
	_ = v399
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v428 int32
	_ = v428
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
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
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v607 int32
	_ = v607
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v642 int32
	_ = v642
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v667 int32
	_ = v667
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v705 int32
	_ = v705
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v722 int32
	_ = v722
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 float64
	_ = v782
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
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
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v833 int32
	_ = v833
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
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1008 float64
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1019 float64
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1025 float64
	_ = v1025
	var v1026 float64
	_ = v1026
	var v1031 int32
	_ = v1031
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1045 int32
	_ = v1045
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1141 int32
	_ = v1141
	var v1146 int32
	_ = v1146
	var v1153 int32
	_ = v1153
	var v1172 float64
	_ = v1172
	var v1173 float64
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1184 float64
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1191 float64
	_ = v1191
	var v1194 float64
	_ = v1194
	var v1196 float64
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1206 int32
	_ = v1206
	var v1232 float64
	_ = v1232
	var v1233 float64
	_ = v1233
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1244 float64
	_ = v1244
	var v1276 float64
	_ = v1276
	var v1277 float64
	_ = v1277
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1339 int32
	_ = v1339
	var v1345 int32
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1367 int32
	_ = v1367
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1389 int32
	_ = v1389
	var v1391 int64
	_ = v1391
	var v1395 int32
	_ = v1395
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1423 int32
	_ = v1423
	var v1429 int32
	_ = v1429
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1435 float64
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1443 float64
	_ = v1443
	var v1448 float64
	_ = v1448
	var v1449 float64
	_ = v1449
	var v1451 float64
	_ = v1451
	var v1454 float64
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1473 int32
	_ = v1473
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1499 int32
	_ = v1499
	var v1503 int32
	_ = v1503
	var v1507 int32
	_ = v1507
	var v1509 int32
	_ = v1509
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1519 int32
	_ = v1519
	var v1522 int32
	_ = v1522
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1535 int32
	_ = v1535
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1559 int32
	_ = v1559
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1575 int32
	_ = v1575
	var v1579 int32
	_ = v1579
	var v1583 int32
	_ = v1583
	var v1586 int32
	_ = v1586
	var v1594 int32
	_ = v1594
	var v1602 int32
	_ = v1602
	var v1606 int32
	_ = v1606
	var v1608 int32
	_ = v1608
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1619 int32
	_ = v1619
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1657 int32
	_ = v1657
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1674 int32
	_ = v1674
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1711 int32
	_ = v1711
	var v1715 int32
	_ = v1715
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1731 int32
	_ = v1731
	var v1734 int32
	_ = v1734
	var v1741 int32
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1747 int32
	_ = v1747
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1771 int32
	_ = v1771
	var v1781 int32
	_ = v1781
	var v1784 int32
	_ = v1784
	var v1787 int32
	_ = v1787
	var v1791 int32
	_ = v1791
	var v1795 int32
	_ = v1795
	var v1798 int32
	_ = v1798
	var v1806 int32
	_ = v1806
	var v1814 int32
	_ = v1814
	var v1818 int32
	_ = v1818
	var v1820 int32
	_ = v1820
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1831 int32
	_ = v1831
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1869 int32
	_ = v1869
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1886 int32
	_ = v1886
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1915 int32
	_ = v1915
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1936 int32
	_ = v1936
	var v1940 int32
	_ = v1940
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1958 int32
	_ = v1958
	var v1961 int32
	_ = v1961
	var v1969 int32
	_ = v1969
	var v1973 int32
	_ = v1973
	var v1980 int32
	_ = v1980
	var v1999 float64
	_ = v1999
	var v2000 float64
	_ = v2000
	var v2003 int32
	_ = v2003
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2011 float64
	_ = v2011
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2018 float64
	_ = v2018
	var v2021 float64
	_ = v2021
	var v2023 float64
	_ = v2023
	var v2025 int32
	_ = v2025
	var v2027 int32
	_ = v2027
	var v2033 int32
	_ = v2033
	var v2059 float64
	_ = v2059
	var v2060 float64
	_ = v2060
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2071 float64
	_ = v2071
	var v2103 float64
	_ = v2103
	var v2104 float64
	_ = v2104
	var v2149 int32
	_ = v2149
	var v2151 float64
	_ = v2151
	var v2153 int32
	_ = v2153
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2161 int64
	_ = v2161
	var v2163 int32
	_ = v2163
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2173 float64
	_ = v2173
	var v2179 int32
	_ = v2179
	var v2182 int32
	_ = v2182
	var v2185 int32
	_ = v2185
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2194 int32
	_ = v2194
	var v2198 int32
	_ = v2198
	var v2205 int32
	_ = v2205
	var v2209 int32
	_ = v2209
	var v2230 int32
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2232 int32
	_ = v2232
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2250 int32
	_ = v2250
	var v2256 int32
	_ = v2256
	var v2263 int32
	_ = v2263
	var v2286 int32
	_ = v2286
	var v2289 int32
	_ = v2289
	var v2296 int32
	_ = v2296
	var v2302 int32
	_ = v2302
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2325 int32
	_ = v2325
	var v2326 int32
	_ = v2326
	var v2329 int32
	_ = v2329
	var v2333 int32
	_ = v2333
	var v2364 int32
	_ = v2364
	var v2369 int32
	_ = v2369
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2376 int32
	_ = v2376
	var v2378 int32
	_ = v2378
	var v2403 int32
	_ = v2403
	var v2414 int32
	_ = v2414
	var v2416 int32
	_ = v2416
	var v2420 int32
	_ = v2420
	var v2440 int32
	_ = v2440
	var v2450 int32
	_ = v2450
	var v2452 int64
	_ = v2452
	var v2456 int32
	_ = v2456
	var v2459 int32
	_ = v2459
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2463 int32
	_ = v2463
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2468 int32
	_ = v2468
	var v2506 int32
	_ = v2506
	var v2509 int32
	_ = v2509
	var v2512 int32
	_ = v2512
	var v2513 int32
	_ = v2513
	var v2515 int32
	_ = v2515
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2525 int32
	_ = v2525
	var v2526 int32
	_ = v2526
	var v2530 int32
	_ = v2530
	var v2534 int32
	_ = v2534
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2550 int32
	_ = v2550
	var v2553 int32
	_ = v2553
	var v2556 int32
	_ = v2556
	var v2557 int32
	_ = v2557
	var v2561 int32
	_ = v2561
	var v2565 int32
	_ = v2565
	var v2566 int32
	_ = v2566
	var v2570 int32
	_ = v2570
	var v2574 int32
	_ = v2574
	var v2580 int32
	_ = v2580
	var v2581 int32
	_ = v2581
	var v2584 int32
	_ = v2584
	var v2585 int32
	_ = v2585
	var v2586 int32
	_ = v2586
	var v2587 int32
	_ = v2587
	var v2590 int32
	_ = v2590
	var v2591 int32
	_ = v2591
	var v2592 float64
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2595 float64
	_ = v2595
	var v2596 int32
	_ = v2596
	var v2598 float64
	_ = v2598
	var v2600 float64
	_ = v2600
	var v2603 int32
	_ = v2603
	var v2605 int32
	_ = v2605
	var v2608 int32
	_ = v2608
	var v2609 int32
	_ = v2609
	var v2611 int32
	_ = v2611
	var v2614 int32
	_ = v2614
	var v2616 int32
	_ = v2616
	var v2619 int32
	_ = v2619
	var v2620 int32
	_ = v2620
	var v2622 int32
	_ = v2622
	var v2625 int32
	_ = v2625
	var v2626 int32
	_ = v2626
	var v2628 int32
	_ = v2628
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2631 int32
	_ = v2631
	var v2632 int32
	_ = v2632
	var v2634 int32
	_ = v2634
	var v2635 int32
	_ = v2635
	var v2636 int32
	_ = v2636
	var v2638 int32
	_ = v2638
	var v2641 int32
	_ = v2641
	var v2642 int32
	_ = v2642
	var v2644 int32
	_ = v2644
	var v2645 int32
	_ = v2645
	var v2646 int32
	_ = v2646
	var v2647 int32
	_ = v2647
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2651 int32
	_ = v2651
	var v2653 int32
	_ = v2653
	var v2665 int32
	_ = v2665
	var v2666 int32
	_ = v2666
	var v2667 int32
	_ = v2667
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2671 int32
	_ = v2671
	var v2673 int32
	_ = v2673
	var v2675 int32
	_ = v2675
	var v2679 int32
	_ = v2679
	var v2690 int32
	_ = v2690
	var v2712 int32
	_ = v2712
	var v2714 int32
	_ = v2714
	var v2715 int32
	_ = v2715
	var v2716 int32
	_ = v2716
	var v2718 int32
	_ = v2718
	var v2720 int32
	_ = v2720
	var v2723 int32
	_ = v2723
	var v2726 int32
	_ = v2726
	var v2740 int32
	_ = v2740
	var v2758 int32
	_ = v2758
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2768 int32
	_ = v2768
	var v2769 int32
	_ = v2769
	var v2770 int32
	_ = v2770
	var v2771 int32
	_ = v2771
	var v2774 int32
	_ = v2774
	var v2776 int32
	_ = v2776
	var v2777 int32
	_ = v2777
	var v2782 int32
	_ = v2782
	var v2785 int32
	_ = v2785
	var v2786 int32
	_ = v2786
	var v2788 int32
	_ = v2788
	var v2854 int32
	_ = v2854
	var v2887 int32
	_ = v2887
	var v2889 int32
	_ = v2889
	var v2892 int32
	_ = v2892
	var v2893 int32
	_ = v2893
	var v2895 int32
	_ = v2895
	var v2897 int32
	_ = v2897
	var v2898 int32
	_ = v2898
	var v2899 int32
	_ = v2899
	var v2900 int32
	_ = v2900
	var v2901 int32
	_ = v2901
	var v2902 int32
	_ = v2902
	var v2905 int32
	_ = v2905
	var v2911 int32
	_ = v2911
	var v2941 int32
	_ = v2941
	var v2944 int32
	_ = v2944
	var v2945 int32
	_ = v2945
	var v2946 int32
	_ = v2946
	var v2947 int32
	_ = v2947
	var v2948 int32
	_ = v2948
	var v2952 int32
	_ = v2952
	var v2953 int32
	_ = v2953
	var v2987 int32
	_ = v2987
	var v2990 int32
	_ = v2990
	var v2991 int32
	_ = v2991
	var v2994 int32
	_ = v2994
	var v3026 int32
	_ = v3026
	var v3029 int32
	_ = v3029
	var v3030 int32
	_ = v3030
	var v3031 int32
	_ = v3031
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3036 int32
	_ = v3036
	var v3037 int32
	_ = v3037
	var v3072 int32
	_ = v3072
	var v3073 int32
	_ = v3073
	var v3077 int32
	_ = v3077
	var v3079 int32
	_ = v3079
	var v3090 int32
	_ = v3090
	v9 = int32(0)
	v29 = float64(0)
	v33 = m.G0
	v35 = v33 - int32(160)
	m.G0 = v35
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v37)
	F_check_stack_depth(m)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v43 != int32(142) {
		goto L11
	} else {
		goto L12
	}
L3:
	;
	m.G0 = v35 + int32(160)
	return v3090
L4:
	;
	v2712 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v2714 = F_tlist_same_datatypes(m, v2712, l3, int32(0))
	mBase = m.M
	v2715 = m.ExcPending
	if v2715 != 0 {
		goto L1
	} else {
		goto L630
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+112)) = v1928
	v1950 = F_fetch_upper_rel(m, l1, int32(0), v1931)
	mBase = m.M
	v1951 = m.ExcPending
	if v1951 != 0 {
		goto L1
	} else {
		goto L488
	}
L6:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+60))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1089)+60))
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1089)+8))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+8))
	v1098 = F_bms_union(m, v1096, v1097)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L1
	} else {
		goto L254
	}
L7:
	;
	v1089 = v792
	v1090 = v801
	v1091 = v808
	goto L6
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L1
	} else {
		goto L251
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L1
	} else {
		goto L243
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L1
	} else {
		goto L240
	}
L11:
	;
	if v43 != int32(63) {
		goto L8
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v86 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48+v49<<(uint(int32(2))%32))))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+36))
	v56 = F_build_simple_rel(m, l1, v49, int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v61 = F_choose_plan_name(m, v58, int32(_a_F_recurse_set_operations_0), int32(1))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v64 = int32(0)
	v66 = *(*float64)(unsafe.Add(mBase, uint32(l1)+312))
	v67 = F_subquery_planner(m, v63, v54, v61, l1, v64, v64, v66, l2)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+148)) = v67
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v70 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v67)+284))
	v76 = F_generate_setop_tlist(m, l3, l4, v71, int32(1), v73, l5, v35+int32(112))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v78 = F_make_pathtarget_from_tlist(m, v76)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v80 = F_set_pathtarget_cost_width(m, l1, v78)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+40)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v76
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+112)))
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v84)
	v3090 = v56
	goto L3
L22:
	;
	v89 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+120)) = v89
	v91 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+112)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v35)+104)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v35)+96)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v35)+44)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v35)+136)) = l0
	v103 = F_list_make1_impl(m, int32(1), v35+int32(44))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v782 = *(*float64)(unsafe.Add(mBase, uint32(l1)+312))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+312)) = int64(0)
	v785 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v792 = F_recurse_set_operations(m, v785, l1, l0, v786, v787, l5, v35+int32(156), v35+int32(151))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L1
	} else {
		goto L181
	}
L25:
	;
	if v103 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v107 = v89
	v114 = v103
	v117 = v9
	v120 = v9
	goto L29
L27:
	;
	v198 = v89
	v208 = v9
	v211 = v9
	goto L28
L28:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v230 = F_generate_append_tlist(m, v228, v229, v211, l5)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L51
	}
L29:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	v139 = F_list_delete_first(m, v114)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	v198 = v191
	v208 = v193
	v211 = v194
	goto L28
L31:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	if v141 != int32(142) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v192 != 0 {
		v107 = v191
		v114 = v192
		v117 = v193
		v120 = v194
		goto L29
	} else {
		goto L50
	}
L33:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v173 != 0 {
		goto L43
	} else {
		goto L44
	}
L34:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v144 != v145 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+8)))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if base.B2i32(v147 != v148)&base.B2i32(v147 == int32(0)) != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v155 = F_equal(m, v153, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v155 == int32(0) {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v138)+28))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v161 = F_equal(m, v159, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v161 == int32(0) {
		goto L33
	} else {
		goto L40
	}
L40:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	v166 = F_lcons(m, v165, v139)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	v169 = F_lcons(m, v168, v166)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v191 = v107
	v192 = v169
	v193 = v117
	v194 = v120
	goto L32
L43:
	;
	v174 = int32(0)
	goto L45
L44:
	;
	v174 = l0
	goto L45
L45:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v181 = F_recurse_set_operations(m, v138, l1, v174, v175, v176, l5, v35+int32(156), v35+int32(152))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v183 = F_lappend(m, v107, v181)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v35)+156))
	v186 = F_lappend(m, v120, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+152)))
	v189 = F_lappend_int(m, v117, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v191 = v183
	v192 = v139
	v193 = v189
	v194 = v186
	goto L32
L50:
	;
	goto L30
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v230
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v234 != 0 {
		v399 = v9
		v410 = v9
		v417 = int32(0)
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v428 = int32(0)
	goto L82
L53:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v236 = F_copyObjectImpl(m, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	if v236 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v236)+12))
	v240 = v238
	goto L57
L56:
	;
	v240 = int32(0)
	goto L57
L57:
	;
	if v230 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v332 = int32(0)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v333 == v332 {
		goto L68
	} else {
		goto L69
	}
L59:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	if v243 <= int32(0) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v255 = v240
	v256 = v9
	goto L61
L61:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v236)+12))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v230)+12))
	v282 = int32(2)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v281+v256<<(uint(v282)%32))))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v280)+4)) = v286
	v289 = v255 + int32(4)
	if base.Ui32(v289) < base.Ui32(v278+v279<<(uint(v282)%32)) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L58
L63:
	;
	v295 = v289
	goto L65
L64:
	;
	v295 = int32(0)
	goto L65
L65:
	;
	v297 = v256 + int32(1)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	if v297 < v298 {
		v255 = v295
		v256 = v297
		goto L61
	} else {
		goto L66
	}
L66:
	;
	goto L62
L67:
	;
	if v378 == int32(0) {
		v399 = v236
		v410 = v9
		v417 = v332
		goto L52
	} else {
		goto L80
	}
L68:
	;
	v378 = int32(1)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	if v342 <= int32(0) {
		v370 = int32(1)
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v378 = v370
	goto L67
L72:
	;
	v345 = int32(0)
	if v345 < v342 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v348 = v342
	goto L75
L74:
	;
	v348 = v345
	goto L75
L75:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v333)+12))
	v351 = int32(0)
	goto L76
L76:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v349+v351<<(uint(int32(2))%32))))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+12))
	v361 = int32(0)
	v362 = base.B2i32(v360 != v361)
	if v360 == v361 {
		v370 = v362
		goto L71
	} else {
		goto L78
	}
L77:
	;
	v370 = v362
	goto L71
L78:
	;
	v366 = v351 + int32(1)
	if v366 != v348 {
		v351 = v366
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v381 = F_make_pathkeys_for_sortclauses(m, l1, v236, v230)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+176)) = v381
	v399 = v236
	v410 = v381
	v417 = int32(1)
	goto L52
L82:
	;
	v451 = int32(0)
	if v198 == v451 {
		v461 = v451
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v516 = int32(0)
	v529 = v516
	v530 = v417
	v531 = v489
	v532 = v490
	v534 = v516
	v535 = int32(1)
	v539 = v9
	v543 = v9
	goto L105
L84:
	;
	v462 = int32(0)
	if v208 == v462 {
		v471 = v462
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v455 <= v428 {
		v461 = int32(0)
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v198)+12))
	v461 = v457 + v428<<(uint(int32(2))%32)
	goto L84
L87:
	;
	if v211 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L88:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	if v465 <= v428 {
		v471 = v462
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v208)+12))
	v471 = v467 + v428<<(uint(int32(2))%32)
	goto L87
L90:
	;
	goto L83
L91:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v461)))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v500)+84))
	if v501 == int32(1) {
		goto L101
	} else {
		goto L102
	}
L92:
	;
	if v198 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L93:
	;
	v474 = int32(0)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	if base.B2i32(v471 == v474)|(base.B2i32(v461 == v474)|base.B2i32(v478 <= v428)) != 0 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	if v482 != 0 {
		goto L91
	} else {
		goto L95
	}
L95:
	;
	goto L92
L96:
	;
	v1927 = v417
	v1928 = v495
	v1929 = v496
	v1931 = int32(0)
	v1932 = int32(1)
	v1936 = v497
	v1940 = v9
	goto L5
L97:
	;
	v487 = int32(0)
	v495 = v487
	v496 = int32(1)
	v497 = v487
	goto L96
L98:
	;
	goto L99
L99:
	;
	v489 = int32(0)
	v490 = int32(1)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v489 < v491 {
		goto L90
	} else {
		goto L100
	}
L100:
	;
	v495 = v489
	v496 = v490
	v497 = int32(0)
	goto L96
L101:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v471)))
	v505 = int32(0)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v482+v428<<(uint(int32(2))%32))))
	F_build_setop_child_paths(m, l1, v500, base.B2i32(v504 != v505), v510, v410, v505)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v428 = v428 + int32(1)
	goto L82
L104:
	;
	goto L103
L105:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v198)+12))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v551+v529<<(uint(int32(2))%32))))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v555)+8))
	v557 = F_bms_add_members(m, v534, v556)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L107
	}
L106:
	;
	v1927 = v771
	v1928 = v772
	v1929 = v773
	v1931 = v557
	v1932 = v774
	v1936 = v776
	v1940 = v777
	goto L5
L107:
	;
	v559 = int32(0)
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v555)+44))
	if v561 == v559 {
		v582 = v559
		goto L110
	} else {
		goto L111
	}
L108:
	;
	v779 = v529 + int32(1)
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v779 < v780 {
		v529 = v779
		v530 = v771
		v531 = v772
		v532 = v773
		v534 = v557
		v535 = v774
		v539 = v776
		v543 = v777
		goto L105
	} else {
		goto L180
	}
L109:
	;
	if v582 != 0 {
		v771 = v530
		v772 = v531
		v773 = v532
		v774 = v535
		v776 = v539
		v777 = v543
		goto L108
	} else {
		goto L119
	}
L110:
	;
	goto L109
L111:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v561)+12))
	v565 = v564
	goto L112
L112:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v565)))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	if base.Ui32(int32(2)) <= base.Ui32(v569-int32(303)) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v582 = int32(1)
	goto L110
L114:
	;
	if v569 != int32(293) {
		v582 = v559
		goto L110
	} else {
		goto L117
	}
L115:
	;
	v565 = v568 + int32(72)
	goto L112
L116:
	;
	goto L113
L117:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v568)+72))
	if v576 != 0 {
		v582 = v559
		goto L110
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v555)+60))
	v584 = F_lappend(m, v531, v583)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	if v530&int32(1) == int32(0) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	if v532&int32(1) == int32(0) {
		goto L170
	} else {
		goto L171
	}
L122:
	;
	v747 = int32(0)
	v749 = v543
	goto L121
L123:
	;
	goto L124
L124:
	;
	v591 = int32(0)
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v555)+44))
	if v592 == v591 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	if v741 == int32(0) {
		v747 = v591
		v749 = v543
		goto L121
	} else {
		goto L168
	}
L126:
	;
	v741 = int32(0)
	goto L125
L127:
	;
	goto L128
L128:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v592)+4))
	if int32(0) < v607 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v617 = v591
	v620 = v591
	goto L132
L130:
	;
	v722 = v591
	goto L131
L131:
	;
	v741 = v722
	goto L125
L132:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v592)+12))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v623+v620<<(uint(int32(2))%32))))
	goto L136
L133:
	;
	v722 = v705
	goto L131
L134:
	;
	v712 = v620 + int32(1)
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v592)+4))
	if v712 < v713 {
		v617 = v705
		v620 = v712
		goto L132
	} else {
		goto L167
	}
L136:
	;
	goto L137
L137:
	;
	if v617 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v631 = F_compare_path_costs(m, v617, v627, int32(1))
	mBase = m.M
	if v631 <= int32(0) {
		v705 = v617
		goto L134
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v627)+64))
	if v410 == v634 {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	goto L141
L143:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v627)+16))
	if v692 != 0 {
		goto L161
	} else {
		goto L162
	}
L144:
	;
	v642 = int32(0)
	goto L145
L145:
	;
	v650 = int32(0)
	if v410 == v650 {
		v660 = v650
		goto L147
	} else {
		goto L148
	}
L146:
	;
	if v660 != 0 {
		v705 = v617
		goto L134
	} else {
		goto L160
	}
L147:
	;
	if v634 != 0 {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	if v654 <= v642 {
		v660 = int32(0)
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v410)+12))
	v660 = v656 + v642<<(uint(int32(2))%32)
	goto L147
L150:
	;
	if v660 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L151:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v634)+4))
	if v642 < v661 {
		goto L150
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	if v660 == int32(0) {
		goto L143
	} else {
		goto L155
	}
L154:
	;
	goto L153
L155:
	;
	v705 = v617
	goto L134
L156:
	;
	goto L146
L157:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v634)+12))
	if v667 == int32(0) {
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v660)))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v667+v642<<(uint(int32(2))%32))))
	if v674 == v676 {
		v642 = v642 + int32(1)
		goto L145
	} else {
		goto L159
	}
L159:
	;
	v705 = v617
	goto L134
L160:
	;
	goto L143
L161:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v692)+4))
	v695 = v693
	goto L163
L162:
	;
	v695 = int32(0)
	goto L163
L163:
	;
	v696 = F_bms_is_subset(m, v695, v591)
	mBase = m.M
	if v696 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v697 = v627
	goto L166
L165:
	;
	v697 = v617
	goto L166
L166:
	;
	v705 = v697
	goto L134
L167:
	;
	goto L133
L168:
	;
	v745 = F_lappend(m, v543, v741)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	v747 = int32(1)
	v749 = v745
	goto L121
L170:
	;
	v771 = v747
	v772 = v584
	v773 = int32(0)
	v774 = v535
	v776 = v539
	v777 = v749
	goto L108
L171:
	;
	goto L172
L172:
	;
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+26)))
	if v755 != int32(1) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v758 = int32(0)
	v771 = v747
	v772 = v584
	v773 = v758
	v774 = v758
	v776 = v539
	v777 = v749
	goto L108
L174:
	;
	goto L175
L175:
	;
	v760 = int32(1)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v555)+52))
	if v761 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v771 = v747
	v772 = v584
	v773 = v760
	v774 = int32(0)
	v776 = v539
	v777 = v749
	goto L108
L177:
	;
	goto L178
L178:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v761)+12))
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v765)))
	v767 = F_lappend(m, v539, v766)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+100)) = v767
	v771 = v747
	v772 = v584
	v773 = v760
	v774 = v535
	v776 = v767
	v777 = v749
	goto L108
L180:
	;
	goto L106
L181:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v801 = F_recurse_set_operations(m, v794, l1, l0, v795, v796, l5, v35+int32(152), v35+int32(150))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	v803 = int32(0)
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v805 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v35)+156))
	v811 = F_generate_setop_tlist(m, v804, v805, v803, v803, v808, l5, v35+int32(149))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v811
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v815 = F_copyObjectImpl(m, v814)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	if v815 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v815)+12))
	v818 = v817
	goto L187
L186:
	;
	v818 = v803
	goto L187
L187:
	;
	if v811 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	if v815 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L189:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v811)+4))
	if v821 <= int32(0) {
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v826 = v818
	v833 = v9
	goto L191
L191:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v815)+12))
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v815)+4))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v826)))
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v811)+12))
	v860 = int32(2)
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v859+v833<<(uint(v860)%32))))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v863)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v858)+4)) = v864
	v867 = v826 + int32(4)
	if base.Ui32(v867) < base.Ui32(v856+v857<<(uint(v860)%32)) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	goto L188
L193:
	;
	v873 = v867
	goto L195
L194:
	;
	v873 = int32(0)
	goto L195
L195:
	;
	v875 = v833 + int32(1)
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v811)+4))
	if v875 < v876 {
		v826 = v873
		v833 = v875
		goto L191
	} else {
		goto L196
	}
L196:
	;
	goto L192
L197:
	;
	v955 = int32(0)
	if v815 == v955 {
		goto L211
	} else {
		goto L212
	}
L198:
	;
	v954 = int32(1)
	goto L197
L199:
	;
	goto L200
L200:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v815)+4))
	if v918 <= int32(0) {
		v946 = int32(1)
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v954 = v946
	goto L197
L202:
	;
	v921 = int32(0)
	if v921 < v918 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v924 = v918
	goto L205
L204:
	;
	v924 = v921
	goto L205
L205:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v815)+12))
	v927 = int32(0)
	goto L206
L206:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v925+v927<<(uint(int32(2))%32))))
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v935)+12))
	v937 = int32(0)
	v938 = base.B2i32(v936 != v937)
	if v936 == v937 {
		v946 = v938
		goto L201
	} else {
		goto L208
	}
L207:
	;
	v946 = v938
	goto L201
L208:
	;
	v942 = v927 + int32(1)
	if v942 != v924 {
		v927 = v942
		goto L206
	} else {
		goto L209
	}
L209:
	;
	goto L207
L210:
	;
	if v954|v992 == int32(0) {
		goto L9
	} else {
		goto L223
	}
L211:
	;
	v992 = int32(1)
	goto L210
L212:
	;
	goto L213
L213:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v815)+4))
	if v962 <= int32(0) {
		v986 = int32(1)
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v992 = v986
	goto L210
L215:
	;
	v965 = int32(0)
	if v965 < v962 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v968 = v962
	goto L218
L217:
	;
	v968 = v965
	goto L218
L218:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v815)+12))
	v973 = v955
	goto L219
L219:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v969+v973<<(uint(int32(2))%32))))
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v977)+18)))
	if v978 != int32(1) {
		v986 = v978
		goto L214
	} else {
		goto L221
	}
L220:
	;
	v986 = v978
	goto L214
L221:
	;
	v982 = v973 + int32(1)
	if v982 != v968 {
		v973 = v982
		goto L219
	} else {
		goto L222
	}
L222:
	;
	goto L220
L223:
	;
	if v954 != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v996 = F_make_pathkeys_for_sortclauses(m, l1, v815, v811)
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L1
	} else {
		goto L227
	}
L225:
	;
	v999 = v9
	goto L226
L226:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v792)+84))
	if v1000 == int32(1) {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+176)) = v996
	v999 = v996
	goto L226
L228:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v801)+84))
	if v1010 == int32(1) {
		goto L234
	} else {
		goto L235
	}
L229:
	;
	v1003 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+151)))
	F_build_setop_child_paths(m, l1, v792, v1003, v808, v999, v35+int32(96))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L1
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	v1008 = *(*float64)(unsafe.Add(mBase, uint32(v792)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v35)+96)) = v1008
	goto L228
L232:
	;
	goto L228
L233:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+312)) = v782
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1022 == int32(3) {
		goto L7
	} else {
		goto L238
	}
L234:
	;
	v1013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+150)))
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v35)+152))
	F_build_setop_child_paths(m, l1, v801, v1013, v1014, v999, v35+int32(136))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L1
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v1019 = *(*float64)(unsafe.Add(mBase, uint32(v801)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v35)+136)) = v1019
	goto L233
L237:
	;
	goto L233
L238:
	;
	v1025 = *(*float64)(unsafe.Add(mBase, uint32(v35)+136))
	v1026 = *(*float64)(unsafe.Add(mBase, uint32(v35)+96))
	if base.F64_lt(v1025, v1026) == int32(0) {
		goto L7
	} else {
		goto L239
	}
L239:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v35)+96)) = v1025
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v35)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+152)) = v808
	v1089 = v801
	v1090 = v792
	v1091 = v1031
	goto L6
L240:
	;
	F_errmsg_internal(m, int32(_a_F_recurse_set_operations_1), int32(0))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(_a_F_recurse_set_operations_2), int32(262), int32(_a_F_recurse_set_operations_3))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L243:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1055 == int32(2) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1058 = int32(_a_F_recurse_set_operations_4)
	goto L247
L246:
	;
	v1058 = int32(_a_F_recurse_set_operations_5)
	goto L247
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+80)) = v1058
	F_errmsg(m, int32(_a_F_recurse_set_operations_6), v35+int32(80))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	v1067 = F_errdetail(m, int32(_a_F_recurse_set_operations_7), int32(0))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(_a_F_recurse_set_operations_2), int32(1127), int32(_a_F_recurse_set_operations_8))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L251:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v1078
	F_errmsg_internal(m, int32(_a_F_recurse_set_operations_9), v35)
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	F_errfinish(m, int32(_a_F_recurse_set_operations_2), int32(353), int32(_a_F_recurse_set_operations_3))
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L254:
	;
	v1100 = F_fetch_upper_rel(m, l1, int32(0), v1098)
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1102 == int32(3) {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	v1123 = F_make_pathtarget_from_tlist(m, v811)
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L1
	} else {
		goto L262
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+68)) = v1094
	*(*int32)(unsafe.Add(mBase, uint32(v35)+132)) = v1094
	v1110 = F_list_make1_impl(m, int32(1), v35+int32(68))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L1
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+128)) = v1093
	*(*int32)(unsafe.Add(mBase, uint32(v35)+132)) = v1094
	*(*int32)(unsafe.Add(mBase, uint32(v35)+76)) = v1094
	*(*int32)(unsafe.Add(mBase, uint32(v35)+72)) = v1093
	v1120 = F_list_make2_impl(m, v35+int32(76), v35+int32(72))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L1
	} else {
		goto L261
	}
L260:
	;
	v1122 = v1110
	goto L256
L261:
	;
	v1122 = v1120
	goto L256
L262:
	;
	v1125 = F_set_pathtarget_cost_width(m, l1, v1123)
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	if v1122 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1100)+40)) = v1125
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1322 = int32(0)
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v1089)+44))
	if v1324 == v1322 {
		v1345 = v1322
		goto L278
	} else {
		goto L279
	}
L265:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1122)+4))
	if v1129 <= int32(0) {
		goto L264
	} else {
		goto L266
	}
L266:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1122)+12))
	if v1129 == int32(1) {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	if base.F64_gt(v1276, float64(0)) == int32(0) {
		goto L264
	} else {
		goto L276
	}
L268:
	;
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1132+v1206<<(uint(int32(2))%32))))
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1239)+8))
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v1240)+40))
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+32))
	v1244 = *(*float64)(unsafe.Add(mBase, uint32(v1239)+32))
	v1276 = base.F64_add(v1232, v1244)
	v1277 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1242), v1244), v1233)
	goto L267
L269:
	;
	v1206 = int32(0)
	v1232 = float64(0)
	v1233 = v29
	goto L268
L270:
	;
	goto L271
L271:
	;
	v1141 = int32(0)
	v1146 = v1141
	v1153 = v1141
	v1172 = float64(0)
	v1173 = v29
	goto L272
L272:
	;
	v1176 = int32(2)
	v1178 = v1132 + v1146<<(uint(v1176)%32)
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1178)+4))
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1179)+8))
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1180)+40))
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v1181)+32))
	v1184 = *(*float64)(unsafe.Add(mBase, uint32(v1179)+32))
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v1178)))
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+8))
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1187)+40))
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1188)+32))
	v1191 = *(*float64)(unsafe.Add(mBase, uint32(v1186)+32))
	v1194 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1182), v1184), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1189), v1191), v1173))
	v1196 = base.F64_add(base.F64_add(v1172, v1191), v1184)
	v1198 = v1146 + v1176
	v1200 = v1153 + v1176
	if v1200 != v1129&int32(2147483646) {
		v1146 = v1198
		v1153 = v1200
		v1172 = v1196
		v1173 = v1194
		goto L272
	} else {
		goto L274
	}
L273:
	;
	if v1129&int32(1) == int32(0) {
		v1276 = v1196
		v1277 = v1194
		goto L267
	} else {
		goto L275
	}
L274:
	;
	goto L273
L275:
	;
	v1206 = v1198
	v1232 = v1196
	v1233 = v1194
	goto L268
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1125)+32)) = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_div(v1277, v1276)))
	goto L264
L277:
	;
	if v1321 == int32(3) {
		goto L288
	} else {
		goto L289
	}
L278:
	;
	goto L277
L279:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v1324)+12))
	v1328 = v1327
	goto L280
L280:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1328)))
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1331)))
	if base.Ui32(int32(2)) <= base.Ui32(v1332-int32(303)) {
		goto L282
	} else {
		goto L283
	}
L281:
	;
	v1345 = int32(1)
	goto L278
L282:
	;
	if v1332 != int32(293) {
		v1345 = v1322
		goto L278
	} else {
		goto L285
	}
L283:
	;
	v1328 = v1331 + int32(72)
	goto L280
L284:
	;
	goto L281
L285:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+72))
	if v1339 != 0 {
		v1345 = v1322
		goto L278
	} else {
		goto L286
	}
L286:
	;
	goto L284
L287:
	;
	v1434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v1435 = *(*float64)(unsafe.Add(mBase, uint32(v35)+96))
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1436 == int32(3) {
		goto L326
	} else {
		goto L327
	}
L288:
	;
	if v1345 != 0 {
		goto L291
	} else {
		goto L292
	}
L289:
	;
	goto L290
L290:
	;
	if v1345 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L291:
	;
	F_mark_dummy_rel(m, v1100)
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L1
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	v1350 = int32(0)
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+44))
	if v1352 == v1350 {
		v1373 = v1350
		goto L296
	} else {
		goto L297
	}
L294:
	;
	v2690 = v1100
	goto L4
L295:
	;
	if v1373 == int32(0) {
		goto L287
	} else {
		goto L305
	}
L296:
	;
	goto L295
L297:
	;
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1352)+12))
	v1356 = v1355
	goto L298
L298:
	;
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1356)))
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v1359)))
	if base.Ui32(int32(2)) <= base.Ui32(v1360-int32(303)) {
		goto L300
	} else {
		goto L301
	}
L299:
	;
	v1373 = int32(1)
	goto L296
L300:
	;
	if v1360 != int32(293) {
		v1373 = v1350
		goto L296
	} else {
		goto L303
	}
L301:
	;
	v1356 = v1359 + int32(72)
	goto L298
L302:
	;
	goto L299
L303:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v1359)+72))
	if v1367 != 0 {
		v1373 = v1350
		goto L296
	} else {
		goto L304
	}
L304:
	;
	goto L302
L305:
	;
	v1376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v1376 != int32(1) {
		goto L287
	} else {
		goto L306
	}
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+108)) = v1094
	*(*int64)(unsafe.Add(mBase, uint32(v35)+116)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+60)) = v1094
	v1386 = F_list_make1_impl(m, int32(1), v35+int32(60))
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+112)) = v1386
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v35)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v1389
	v1391 = *(*int64)(unsafe.Add(mBase, uint32(v35)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v35)+48)) = v1391
	v1395 = int32(0)
	v1400 = F_create_append_path(m, l1, v1100, v35+int32(48), v1395, v1395, v1395, v1395, float64(-1))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	F_add_path(m, v1100, v1400)
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	v2690 = v1100
	goto L4
L310:
	;
	v1406 = int32(0)
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+44))
	if v1408 == v1406 {
		v1429 = v1406
		goto L314
	} else {
		goto L315
	}
L311:
	;
	goto L312
L312:
	;
	F_mark_dummy_rel(m, v1100)
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L1
	} else {
		goto L324
	}
L313:
	;
	if v1429 == int32(0) {
		goto L287
	} else {
		goto L323
	}
L314:
	;
	goto L313
L315:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v1408)+12))
	v1412 = v1411
	goto L316
L316:
	;
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1412)))
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(v1415)))
	if base.Ui32(int32(2)) <= base.Ui32(v1416-int32(303)) {
		goto L318
	} else {
		goto L319
	}
L317:
	;
	v1429 = int32(1)
	goto L314
L318:
	;
	if v1416 != int32(293) {
		v1429 = v1406
		goto L314
	} else {
		goto L321
	}
L319:
	;
	v1412 = v1415 + int32(72)
	goto L316
L320:
	;
	goto L317
L321:
	;
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(v1415)+72))
	if v1423 != 0 {
		v1429 = v1406
		goto L314
	} else {
		goto L322
	}
L322:
	;
	goto L320
L323:
	;
	goto L312
L324:
	;
	v2690 = v1100
	goto L4
L325:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1100)+16)) = v1454
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v1456 - int32(2) {
	case 0:
		goto L335
	case 1:
		goto L337
	default:
		goto L336
	}
L326:
	;
	if v1434&int32(1) == int32(0) {
		v1454 = v1435
		goto L325
	} else {
		goto L329
	}
L327:
	;
	goto L328
L328:
	;
	if v1434&int32(1) == int32(0) {
		v1454 = v1435
		goto L325
	} else {
		goto L330
	}
L329:
	;
	v1443 = *(*float64)(unsafe.Add(mBase, uint32(v1094)+32))
	v1454 = v1443
	goto L325
L330:
	;
	v1448 = *(*float64)(unsafe.Add(mBase, uint32(v1094)+32))
	v1449 = *(*float64)(unsafe.Add(mBase, uint32(v1093)+32))
	if base.F64_lt(v1448, v1449) != 0 {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v1451 = v1448
	goto L333
L332:
	;
	v1451 = v1449
	goto L333
L333:
	;
	v1454 = v1451
	goto L325
L334:
	;
	if v992 != 0 {
		goto L344
	} else {
		goto L345
	}
L335:
	;
	v1479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v1480 = v1479
	goto L334
L336:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L1
	} else {
		goto L341
	}
L337:
	;
	v1461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v1461 != 0 {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v1462 = int32(3)
	goto L340
L339:
	;
	v1462 = int32(2)
	goto L340
L340:
	;
	v1480 = v1462
	goto L334
L341:
	;
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+64)) = v1467
	F_errmsg_internal(m, int32(_a_F_recurse_set_operations_10), v35-int32(-64))
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L1
	} else {
		goto L342
	}
L342:
	;
	F_errfinish(m, int32(_a_F_recurse_set_operations_2), int32(1297), int32(_a_F_recurse_set_operations_8))
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
		goto L1
	} else {
		goto L343
	}
L343:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L344:
	;
	v1482 = F_create_setop_path(m, v1100, v1094, v1093, v1480, int32(1), v815, v1435, v1454)
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L1
	} else {
		goto L347
	}
L345:
	;
	goto L346
L346:
	;
	if v954 == int32(0) {
		v2690 = v1100
		goto L4
	} else {
		goto L349
	}
L347:
	;
	F_add_path(m, v1100, v1482)
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L1
	} else {
		goto L348
	}
L348:
	;
	goto L346
L349:
	;
	v1488 = F_make_pathkeys_for_sortclauses(m, l1, v815, v1091)
	mBase = m.M
	v1489 = m.ExcPending
	if v1489 != 0 {
		goto L1
	} else {
		goto L351
	}
L350:
	;
	v1699 = *(*int32)(unsafe.Add(mBase, uint32(v35)+152))
	v1700 = F_make_pathkeys_for_sortclauses(m, l1, v815, v1699)
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L1
	} else {
		goto L419
	}
L351:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+64))
	if v1488 == v1490 {
		goto L353
	} else {
		goto L354
	}
L352:
	;
	if v1543 != 0 {
		goto L370
	} else {
		goto L371
	}
L353:
	;
	v1543 = int32(1)
	goto L352
L354:
	;
	goto L355
L355:
	;
	v1499 = int32(0)
	goto L357
L356:
	;
	v1543 = v1535
	goto L352
L357:
	;
	v1503 = int32(0)
	if v1488 == v1503 {
		v1513 = v1503
		goto L359
	} else {
		goto L360
	}
L358:
	;
	v1535 = int32(0)
	goto L356
L359:
	;
	if v1490 != 0 {
		goto L363
	} else {
		goto L364
	}
L360:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1488)+4))
	if v1507 <= v1499 {
		v1513 = int32(0)
		goto L359
	} else {
		goto L361
	}
L361:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v1488)+12))
	v1513 = v1509 + v1499<<(uint(int32(2))%32)
	goto L359
L362:
	;
	v1519 = base.B2i32(v1513 == int32(0))
	if v1513 == int32(0) {
		v1535 = v1519
		goto L356
	} else {
		goto L367
	}
L363:
	;
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v1490)+4))
	if v1499 < v1514 {
		goto L362
	} else {
		goto L366
	}
L364:
	;
	goto L365
L365:
	;
	v1543 = base.B2i32(v1513 == int32(0))
	goto L352
L366:
	;
	goto L365
L367:
	;
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v1490)+12))
	if v1522 == int32(0) {
		v1535 = v1519
		goto L356
	} else {
		goto L368
	}
L368:
	;
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v1513)))
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v1499<<(uint(int32(2))%32)+v1522)))
	if v1529 == v1531 {
		v1499 = v1499 + int32(1)
		goto L357
	} else {
		goto L369
	}
L369:
	;
	goto L358
L370:
	;
	v1698 = v1094
	goto L350
L371:
	;
	goto L372
L372:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v1089)+44))
	v1545 = int32(0)
	if v1544 == v1545 {
		goto L374
	} else {
		goto L375
	}
L373:
	;
	if v1693 != 0 {
		v1698 = v1693
		goto L350
	} else {
		goto L416
	}
L374:
	;
	v1693 = int32(0)
	goto L373
L375:
	;
	goto L376
L376:
	;
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v1544)+4))
	if int32(0) < v1559 {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v1569 = v1545
	v1572 = v1545
	goto L380
L378:
	;
	v1674 = v1545
	goto L379
L379:
	;
	v1693 = v1674
	goto L373
L380:
	;
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v1544)+12))
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1575+v1572<<(uint(int32(2))%32))))
	goto L384
L381:
	;
	v1674 = v1657
	goto L379
L382:
	;
	v1664 = v1572 + int32(1)
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v1544)+4))
	if v1664 < v1665 {
		v1569 = v1657
		v1572 = v1664
		goto L380
	} else {
		goto L415
	}
L384:
	;
	goto L385
L385:
	;
	if v1569 != 0 {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v1583 = F_compare_path_costs(m, v1569, v1579, int32(1))
	mBase = m.M
	if v1583 <= int32(0) {
		v1657 = v1569
		goto L382
	} else {
		goto L390
	}
L388:
	;
	goto L389
L389:
	;
	v1586 = *(*int32)(unsafe.Add(mBase, uint32(v1579)+64))
	if v999 == v1586 {
		goto L391
	} else {
		goto L392
	}
L390:
	;
	goto L389
L391:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1579)+16))
	if v1644 != 0 {
		goto L409
	} else {
		goto L410
	}
L392:
	;
	v1594 = int32(0)
	goto L393
L393:
	;
	v1602 = int32(0)
	if v999 == v1602 {
		v1612 = v1602
		goto L395
	} else {
		goto L396
	}
L394:
	;
	if v1612 != 0 {
		v1657 = v1569
		goto L382
	} else {
		goto L408
	}
L395:
	;
	if v1586 != 0 {
		goto L399
	} else {
		goto L400
	}
L396:
	;
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(v999)+4))
	if v1606 <= v1594 {
		v1612 = int32(0)
		goto L395
	} else {
		goto L397
	}
L397:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v999)+12))
	v1612 = v1608 + v1594<<(uint(int32(2))%32)
	goto L395
L398:
	;
	if v1612 == int32(0) {
		goto L404
	} else {
		goto L405
	}
L399:
	;
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1586)+4))
	if v1594 < v1613 {
		goto L398
	} else {
		goto L402
	}
L400:
	;
	goto L401
L401:
	;
	if v1612 == int32(0) {
		goto L391
	} else {
		goto L403
	}
L402:
	;
	goto L401
L403:
	;
	v1657 = v1569
	goto L382
L404:
	;
	goto L394
L405:
	;
	v1619 = *(*int32)(unsafe.Add(mBase, uint32(v1586)+12))
	if v1619 == int32(0) {
		goto L404
	} else {
		goto L406
	}
L406:
	;
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v1612)))
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v1619+v1594<<(uint(int32(2))%32))))
	if v1626 == v1628 {
		v1594 = v1594 + int32(1)
		goto L393
	} else {
		goto L407
	}
L407:
	;
	v1657 = v1569
	goto L382
L408:
	;
	goto L391
L409:
	;
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v1644)+4))
	v1647 = v1645
	goto L411
L410:
	;
	v1647 = int32(0)
	goto L411
L411:
	;
	v1648 = F_bms_is_subset(m, v1647, v1545)
	mBase = m.M
	if v1648 != 0 {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v1649 = v1579
	goto L414
L413:
	;
	v1649 = v1569
	goto L414
L414:
	;
	v1657 = v1649
	goto L382
L415:
	;
	goto L381
L416:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+8))
	v1696 = F_create_sort_path(m, v1694, v1094, v1488, float64(-1))
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		goto L1
	} else {
		goto L417
	}
L417:
	;
	v1698 = v1696
	goto L350
L418:
	;
	v1912 = F_create_setop_path(m, v1100, v1698, v1910, v1480, int32(0), v815, v1435, v1454)
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		goto L1
	} else {
		goto L486
	}
L419:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+64))
	if v1700 == v1702 {
		goto L421
	} else {
		goto L422
	}
L420:
	;
	if v1755 != 0 {
		goto L438
	} else {
		goto L439
	}
L421:
	;
	v1755 = int32(1)
	goto L420
L422:
	;
	goto L423
L423:
	;
	v1711 = int32(0)
	goto L425
L424:
	;
	v1755 = v1747
	goto L420
L425:
	;
	v1715 = int32(0)
	if v1700 == v1715 {
		v1725 = v1715
		goto L427
	} else {
		goto L428
	}
L426:
	;
	v1747 = int32(0)
	goto L424
L427:
	;
	if v1702 != 0 {
		goto L431
	} else {
		goto L432
	}
L428:
	;
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1700)+4))
	if v1719 <= v1711 {
		v1725 = int32(0)
		goto L427
	} else {
		goto L429
	}
L429:
	;
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(v1700)+12))
	v1725 = v1721 + v1711<<(uint(int32(2))%32)
	goto L427
L430:
	;
	v1731 = base.B2i32(v1725 == int32(0))
	if v1725 == int32(0) {
		v1747 = v1731
		goto L424
	} else {
		goto L435
	}
L431:
	;
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v1702)+4))
	if v1711 < v1726 {
		goto L430
	} else {
		goto L434
	}
L432:
	;
	goto L433
L433:
	;
	v1755 = base.B2i32(v1725 == int32(0))
	goto L420
L434:
	;
	goto L433
L435:
	;
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v1702)+12))
	if v1734 == int32(0) {
		v1747 = v1731
		goto L424
	} else {
		goto L436
	}
L436:
	;
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v1725)))
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(v1711<<(uint(int32(2))%32)+v1734)))
	if v1741 == v1743 {
		v1711 = v1711 + int32(1)
		goto L425
	} else {
		goto L437
	}
L437:
	;
	goto L426
L438:
	;
	v1910 = v1093
	goto L418
L439:
	;
	goto L440
L440:
	;
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+44))
	v1757 = int32(0)
	if v1756 == v1757 {
		goto L442
	} else {
		goto L443
	}
L441:
	;
	if v1905 != 0 {
		v1910 = v1905
		goto L418
	} else {
		goto L484
	}
L442:
	;
	v1905 = int32(0)
	goto L441
L443:
	;
	goto L444
L444:
	;
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1756)+4))
	if int32(0) < v1771 {
		goto L445
	} else {
		goto L446
	}
L445:
	;
	v1781 = v1757
	v1784 = v1757
	goto L448
L446:
	;
	v1886 = v1757
	goto L447
L447:
	;
	v1905 = v1886
	goto L441
L448:
	;
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v1756)+12))
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v1787+v1784<<(uint(int32(2))%32))))
	goto L452
L449:
	;
	v1886 = v1869
	goto L447
L450:
	;
	v1876 = v1784 + int32(1)
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1756)+4))
	if v1876 < v1877 {
		v1781 = v1869
		v1784 = v1876
		goto L448
	} else {
		goto L483
	}
L452:
	;
	goto L453
L453:
	;
	if v1781 != 0 {
		goto L455
	} else {
		goto L456
	}
L455:
	;
	v1795 = F_compare_path_costs(m, v1781, v1791, int32(1))
	mBase = m.M
	if v1795 <= int32(0) {
		v1869 = v1781
		goto L450
	} else {
		goto L458
	}
L456:
	;
	goto L457
L457:
	;
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(v1791)+64))
	if v999 == v1798 {
		goto L459
	} else {
		goto L460
	}
L458:
	;
	goto L457
L459:
	;
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(v1791)+16))
	if v1856 != 0 {
		goto L477
	} else {
		goto L478
	}
L460:
	;
	v1806 = int32(0)
	goto L461
L461:
	;
	v1814 = int32(0)
	if v999 == v1814 {
		v1824 = v1814
		goto L463
	} else {
		goto L464
	}
L462:
	;
	if v1824 != 0 {
		v1869 = v1781
		goto L450
	} else {
		goto L476
	}
L463:
	;
	if v1798 != 0 {
		goto L467
	} else {
		goto L468
	}
L464:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v999)+4))
	if v1818 <= v1806 {
		v1824 = int32(0)
		goto L463
	} else {
		goto L465
	}
L465:
	;
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v999)+12))
	v1824 = v1820 + v1806<<(uint(int32(2))%32)
	goto L463
L466:
	;
	if v1824 == int32(0) {
		goto L472
	} else {
		goto L473
	}
L467:
	;
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v1798)+4))
	if v1806 < v1825 {
		goto L466
	} else {
		goto L470
	}
L468:
	;
	goto L469
L469:
	;
	if v1824 == int32(0) {
		goto L459
	} else {
		goto L471
	}
L470:
	;
	goto L469
L471:
	;
	v1869 = v1781
	goto L450
L472:
	;
	goto L462
L473:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v1798)+12))
	if v1831 == int32(0) {
		goto L472
	} else {
		goto L474
	}
L474:
	;
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v1824)))
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v1831+v1806<<(uint(int32(2))%32))))
	if v1838 == v1840 {
		v1806 = v1806 + int32(1)
		goto L461
	} else {
		goto L475
	}
L475:
	;
	v1869 = v1781
	goto L450
L476:
	;
	goto L459
L477:
	;
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(v1856)+4))
	v1859 = v1857
	goto L479
L478:
	;
	v1859 = int32(0)
	goto L479
L479:
	;
	v1860 = F_bms_is_subset(m, v1859, v1757)
	mBase = m.M
	if v1860 != 0 {
		goto L480
	} else {
		goto L481
	}
L480:
	;
	v1861 = v1791
	goto L482
L481:
	;
	v1861 = v1781
	goto L482
L482:
	;
	v1869 = v1861
	goto L450
L483:
	;
	goto L449
L484:
	;
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+8))
	v1908 = F_create_sort_path(m, v1906, v1093, v1700, float64(-1))
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L1
	} else {
		goto L485
	}
L485:
	;
	v1910 = v1908
	goto L418
L486:
	;
	F_add_path(m, v1100, v1912)
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L1
	} else {
		goto L487
	}
L487:
	;
	v2690 = v1100
	goto L4
L488:
	;
	v1952 = F_make_pathtarget_from_tlist(m, v230)
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L1
	} else {
		goto L489
	}
L489:
	;
	v1954 = F_set_pathtarget_cost_width(m, l1, v1952)
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		goto L1
	} else {
		goto L490
	}
L490:
	;
	if v1928 == int32(0) {
		goto L491
	} else {
		goto L492
	}
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1950)+40)) = v1954
	v2149 = v1929 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1950)+26)) = uint8(v2149)
	v2151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+312))
	v2153 = base.F64_gt(v2151, float64(0))
	*(*uint8)(unsafe.Add(mBase, uint32(v1950)+24)) = uint8(v2153)
	if v1928 == int32(0) {
		goto L504
	} else {
		goto L505
	}
L492:
	;
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v1928)+4))
	if v1958 <= int32(0) {
		goto L491
	} else {
		goto L493
	}
L493:
	;
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(v1928)+12))
	if v1958 == int32(1) {
		goto L496
	} else {
		goto L497
	}
L494:
	;
	if base.F64_gt(v2103, float64(0)) == int32(0) {
		goto L491
	} else {
		goto L503
	}
L495:
	;
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v1961+v2033<<(uint(int32(2))%32))))
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(v2066)+8))
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v2067)+40))
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v2068)+32))
	v2071 = *(*float64)(unsafe.Add(mBase, uint32(v2066)+32))
	v2103 = base.F64_add(v2059, v2071)
	v2104 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v2069), v2071), v2060)
	goto L494
L496:
	;
	v2033 = int32(0)
	v2059 = v29
	v2060 = v29
	goto L495
L497:
	;
	goto L498
L498:
	;
	v1969 = int32(0)
	v1973 = v1969
	v1980 = v1969
	v1999 = v29
	v2000 = v29
	goto L499
L499:
	;
	v2003 = int32(2)
	v2005 = v1961 + v1973<<(uint(v2003)%32)
	v2006 = *(*int32)(unsafe.Add(mBase, uint32(v2005)+4))
	v2007 = *(*int32)(unsafe.Add(mBase, uint32(v2006)+8))
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(v2007)+40))
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(v2008)+32))
	v2011 = *(*float64)(unsafe.Add(mBase, uint32(v2006)+32))
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(v2005)))
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(v2013)+8))
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v2014)+40))
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(v2015)+32))
	v2018 = *(*float64)(unsafe.Add(mBase, uint32(v2013)+32))
	v2021 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v2009), v2011), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v2016), v2018), v2000))
	v2023 = base.F64_add(base.F64_add(v1999, v2018), v2011)
	v2025 = v1973 + v2003
	v2027 = v1980 + v2003
	if v2027 != v1958&int32(2147483646) {
		v1973 = v2025
		v1980 = v2027
		v1999 = v2023
		v2000 = v2021
		goto L499
	} else {
		goto L501
	}
L500:
	;
	if v1958&int32(1) == int32(0) {
		v2103 = v2023
		v2104 = v2021
		goto L494
	} else {
		goto L502
	}
L501:
	;
	goto L500
L502:
	;
	v2033 = v2025
	v2059 = v2023
	v2060 = v2021
	goto L495
L503:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1954)+32)) = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_div(v2104, v2103)))
	goto L491
L504:
	;
	F_mark_dummy_rel(m, v1950)
	mBase = m.M
	v2158 = m.ExcPending
	if v2158 != 0 {
		goto L1
	} else {
		goto L507
	}
L505:
	;
	goto L506
L506:
	;
	v2159 = *(*int32)(unsafe.Add(mBase, uint32(v35)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+40)) = v2159
	v2161 = *(*int64)(unsafe.Add(mBase, uint32(v35)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v35)+32)) = v2161
	v2163 = int32(0)
	v2171 = F_create_append_path(m, l1, v1950, v35+int32(32), v2163, v2163, v2163, v2163, float64(-1))
	mBase = m.M
	v2172 = m.ExcPending
	if v2172 != 0 {
		goto L1
	} else {
		goto L508
	}
L507:
	;
	v2690 = v1950
	goto L4
L508:
	;
	v2173 = *(*float64)(unsafe.Add(mBase, uint32(v2171)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1950)+16)) = v2173
	if v1932&int32(1) != 0 {
		goto L514
	} else {
		goto L515
	}
L509:
	;
	F_add_path(m, v1950, v2171)
	mBase = m.M
	v2675 = m.ExcPending
	if v2675 != 0 {
		goto L1
	} else {
		goto L626
	}
L510:
	;
	if v399 == int32(0) {
		goto L562
	} else {
		goto L563
	}
L511:
	;
	v2450 = *(*int32)(unsafe.Add(mBase, uint32(v35)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v2450
	v2452 = *(*int64)(unsafe.Add(mBase, uint32(v35)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v2452
	v2456 = int32(0)
	v2459 = F_create_append_path(m, l1, v1950, v35+int32(16), v2456, v2456, v2420, v2440, float64(-1))
	mBase = m.M
	v2460 = m.ExcPending
	if v2460 != 0 {
		goto L1
	} else {
		goto L558
	}
L512:
	;
	v2414 = *(*int32)(unsafe.Add(mBase, _c_F_recurse_set_operations[0]))
	if v2403 < v2414 {
		goto L555
	} else {
		goto L556
	}
L513:
	;
	v2378 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_recurse_set_operations[1])))
	if v2378 == int32(0) {
		v2420 = v2163
		v2440 = v9
		goto L511
	} else {
		goto L554
	}
L514:
	;
	if v1936 == int32(0) {
		goto L513
	} else {
		goto L517
	}
L515:
	;
	goto L516
L516:
	;
	v2372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v2372 != int32(1) {
		v2468 = v2163
		goto L510
	} else {
		goto L552
	}
L517:
	;
	v2179 = *(*int32)(unsafe.Add(mBase, uint32(v1936)+4))
	if v2179 <= int32(0) {
		v2333 = v2163
		goto L518
	} else {
		goto L519
	}
L518:
	;
	v2364 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_recurse_set_operations[1])))
	if v2364 != int32(1) {
		v2420 = v2333
		v2440 = v9
		goto L511
	} else {
		goto L548
	}
L519:
	;
	v2182 = int32(0)
	if v2182 < v2179 {
		goto L520
	} else {
		goto L521
	}
L520:
	;
	v2185 = v2179
	goto L522
L521:
	;
	v2185 = v2182
	goto L522
L522:
	;
	v2187 = v2185 & int32(3)
	v2188 = int32(0)
	if int32(4) <= v2179 {
		goto L523
	} else {
		goto L524
	}
L523:
	;
	v2194 = *(*int32)(unsafe.Add(mBase, uint32(v1936)+12))
	v2198 = v2163
	v2205 = v2188
	v2209 = int32(0)
	goto L526
L524:
	;
	v2256 = v2163
	v2263 = v2188
	goto L525
L525:
	;
	v2286 = *(*int32)(unsafe.Add(mBase, uint32(v1936)+12))
	v2289 = v2256
	v2296 = v2263
	v2302 = v2188
	goto L542
L526:
	;
	v2230 = v2194 + v2205<<(uint(int32(2))%32)
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(v2230)))
	v2232 = *(*int32)(unsafe.Add(mBase, uint32(v2231)+24))
	if v2232 < v2198 {
		goto L528
	} else {
		goto L529
	}
L527:
	;
	if v2187 == int32(0) {
		v2333 = v2246
		goto L518
	} else {
		goto L541
	}
L528:
	;
	v2234 = v2198
	goto L530
L529:
	;
	v2234 = v2232
	goto L530
L530:
	;
	v2235 = *(*int32)(unsafe.Add(mBase, uint32(v2230)+4))
	v2236 = *(*int32)(unsafe.Add(mBase, uint32(v2235)+24))
	if v2236 < v2234 {
		goto L531
	} else {
		goto L532
	}
L531:
	;
	v2238 = v2234
	goto L533
L532:
	;
	v2238 = v2236
	goto L533
L533:
	;
	v2239 = *(*int32)(unsafe.Add(mBase, uint32(v2230)+8))
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v2239)+24))
	if v2240 < v2238 {
		goto L534
	} else {
		goto L535
	}
L534:
	;
	v2242 = v2238
	goto L536
L535:
	;
	v2242 = v2240
	goto L536
L536:
	;
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(v2230)+12))
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(v2243)+24))
	if v2244 < v2242 {
		goto L537
	} else {
		goto L538
	}
L537:
	;
	v2246 = v2242
	goto L539
L538:
	;
	v2246 = v2244
	goto L539
L539:
	;
	v2247 = int32(4)
	v2248 = v2205 + v2247
	v2250 = v2209 + v2247
	if v2250 != v2185&int32(2147483644) {
		v2198 = v2246
		v2205 = v2248
		v2209 = v2250
		goto L526
	} else {
		goto L540
	}
L540:
	;
	goto L527
L541:
	;
	v2256 = v2246
	v2263 = v2248
	goto L525
L542:
	;
	v2322 = *(*int32)(unsafe.Add(mBase, uint32(v2286+v2296<<(uint(int32(2))%32))))
	v2323 = *(*int32)(unsafe.Add(mBase, uint32(v2322)+24))
	if v2323 < v2289 {
		goto L544
	} else {
		goto L545
	}
L543:
	;
	v2333 = v2325
	goto L518
L544:
	;
	v2325 = v2289
	goto L546
L545:
	;
	v2325 = v2323
	goto L546
L546:
	;
	v2326 = int32(1)
	v2329 = v2302 + v2326
	if v2329 != v2187 {
		v2289 = v2325
		v2296 = v2296 + v2326
		v2302 = v2329
		goto L542
	} else {
		goto L547
	}
L547:
	;
	goto L543
L548:
	;
	v2369 = int32(32) - base.I32_clz(v2179)
	if v2369 < v2333 {
		goto L549
	} else {
		goto L550
	}
L549:
	;
	v2371 = v2333
	goto L551
L550:
	;
	v2371 = v2369
	goto L551
L551:
	;
	v2403 = v2371
	goto L512
L552:
	;
	F_add_path(m, v1950, v2171)
	mBase = m.M
	v2376 = m.ExcPending
	if v2376 != 0 {
		goto L1
	} else {
		goto L553
	}
L553:
	;
	v2690 = v1950
	goto L4
L554:
	;
	v2403 = v9
	goto L512
L555:
	;
	v2416 = v2403
	goto L557
L556:
	;
	v2416 = v2414
	goto L557
L557:
	;
	v2420 = v2416
	v2440 = int32(1)
	goto L511
L558:
	;
	v2461 = *(*int32)(unsafe.Add(mBase, uint32(v1950)+40))
	v2463 = F_create_gather_path(m, l1, v1950, v2459, v2461, int32(0))
	mBase = m.M
	v2464 = m.ExcPending
	if v2464 != 0 {
		goto L1
	} else {
		goto L559
	}
L559:
	;
	v2465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v2465 != 0 {
		goto L509
	} else {
		goto L560
	}
L560:
	;
	v2468 = v2463
	goto L510
L561:
	;
	v2543 = int32(0)
	if v399 == v2543 {
		goto L575
	} else {
		goto L576
	}
L562:
	;
	v2542 = int32(1)
	goto L561
L563:
	;
	goto L564
L564:
	;
	v2506 = *(*int32)(unsafe.Add(mBase, uint32(v399)+4))
	if v2506 <= int32(0) {
		v2534 = int32(1)
		goto L565
	} else {
		goto L566
	}
L565:
	;
	v2542 = v2534
	goto L561
L566:
	;
	v2509 = int32(0)
	if v2509 < v2506 {
		goto L567
	} else {
		goto L568
	}
L567:
	;
	v2512 = v2506
	goto L569
L568:
	;
	v2512 = v2509
	goto L569
L569:
	;
	v2513 = *(*int32)(unsafe.Add(mBase, uint32(v399)+12))
	v2515 = int32(0)
	goto L570
L570:
	;
	v2523 = *(*int32)(unsafe.Add(mBase, uint32(v2513+v2515<<(uint(int32(2))%32))))
	v2524 = *(*int32)(unsafe.Add(mBase, uint32(v2523)+12))
	v2525 = int32(0)
	v2526 = base.B2i32(v2524 != v2525)
	if v2524 == v2525 {
		v2534 = v2526
		goto L565
	} else {
		goto L572
	}
L571:
	;
	v2534 = v2526
	goto L565
L572:
	;
	v2530 = v2515 + int32(1)
	if v2530 != v2512 {
		v2515 = v2530
		goto L570
	} else {
		goto L573
	}
L573:
	;
	goto L571
L574:
	;
	v2581 = *(*int32)(unsafe.Add(mBase, uint32(v1928)+4))
	if v2581 != int32(1) {
		goto L588
	} else {
		goto L589
	}
L575:
	;
	v2580 = int32(1)
	goto L574
L576:
	;
	goto L577
L577:
	;
	v2550 = *(*int32)(unsafe.Add(mBase, uint32(v399)+4))
	if v2550 <= int32(0) {
		v2574 = int32(1)
		goto L578
	} else {
		goto L579
	}
L578:
	;
	v2580 = v2574
	goto L574
L579:
	;
	v2553 = int32(0)
	if v2553 < v2550 {
		goto L580
	} else {
		goto L581
	}
L580:
	;
	v2556 = v2550
	goto L582
L581:
	;
	v2556 = v2553
	goto L582
L582:
	;
	v2557 = *(*int32)(unsafe.Add(mBase, uint32(v399)+12))
	v2561 = v2543
	goto L583
L583:
	;
	v2565 = *(*int32)(unsafe.Add(mBase, uint32(v2557+v2561<<(uint(int32(2))%32))))
	v2566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2565)+18)))
	if v2566 != int32(1) {
		v2574 = v2566
		goto L578
	} else {
		goto L585
	}
L584:
	;
	v2574 = v2566
	goto L578
L585:
	;
	v2570 = v2561 + int32(1)
	if v2570 != v2556 {
		v2561 = v2570
		goto L583
	} else {
		goto L586
	}
L586:
	;
	goto L584
L587:
	;
	if v2580 == int32(0) {
		goto L592
	} else {
		goto L593
	}
L588:
	;
	v2598 = *(*float64)(unsafe.Add(mBase, uint32(v2171)+32))
	v2600 = v2598
	goto L587
L589:
	;
	v2584 = *(*int32)(unsafe.Add(mBase, uint32(v1928)+12))
	v2585 = *(*int32)(unsafe.Add(mBase, uint32(v2584)))
	v2586 = *(*int32)(unsafe.Add(mBase, uint32(v2585)+8))
	v2587 = *(*int32)(unsafe.Add(mBase, uint32(v2586)+4))
	if v2587 == int32(4) {
		goto L588
	} else {
		goto L590
	}
L590:
	;
	v2590 = *(*int32)(unsafe.Add(mBase, uint32(v2585)+12))
	v2591 = *(*int32)(unsafe.Add(mBase, uint32(v2590)+4))
	v2592 = *(*float64)(unsafe.Add(mBase, uint32(v2585)+32))
	v2593 = int32(0)
	v2595 = F_estimate_num_groups(m, l1, v2591, v2592, v2593, v2593)
	mBase = m.M
	v2596 = m.ExcPending
	if v2596 != 0 {
		goto L1
	} else {
		goto L591
	}
L591:
	;
	v2600 = v2595
	goto L587
L592:
	;
	if v2542 == int32(0) {
		goto L599
	} else {
		goto L600
	}
L593:
	;
	v2603 = *(*int32)(unsafe.Add(mBase, uint32(v1950)+40))
	v2605 = int32(0)
	v2608 = F_create_agg_path(m, l1, v1950, v2171, v2603, int32(2), v2605, v399, v2605, v2605, v2600)
	mBase = m.M
	v2609 = m.ExcPending
	if v2609 != 0 {
		goto L1
	} else {
		goto L594
	}
L594:
	;
	F_add_path(m, v1950, v2608)
	mBase = m.M
	v2611 = m.ExcPending
	if v2611 != 0 {
		goto L1
	} else {
		goto L595
	}
L595:
	;
	if v2468 == int32(0) {
		goto L592
	} else {
		goto L596
	}
L596:
	;
	v2614 = *(*int32)(unsafe.Add(mBase, uint32(v1950)+40))
	v2616 = int32(0)
	v2619 = F_create_agg_path(m, l1, v1950, v2468, v2614, int32(2), v2616, v399, v2616, v2616, v2600)
	mBase = m.M
	v2620 = m.ExcPending
	if v2620 != 0 {
		goto L1
	} else {
		goto L597
	}
L597:
	;
	F_add_path(m, v1950, v2619)
	mBase = m.M
	v2622 = m.ExcPending
	if v2622 != 0 {
		goto L1
	} else {
		goto L598
	}
L598:
	;
	goto L592
L599:
	;
	if (base.B2i32(v399 == int32(0))|(v1927^int32(-1)))&int32(1) != 0 {
		v2690 = v1950
		goto L4
	} else {
		goto L619
	}
L600:
	;
	if v399 != 0 {
		goto L601
	} else {
		goto L602
	}
L601:
	;
	v2625 = F_make_pathkeys_for_sortclauses(m, l1, v399, v230)
	mBase = m.M
	v2626 = m.ExcPending
	if v2626 != 0 {
		goto L1
	} else {
		goto L604
	}
L602:
	;
	v2630 = v2171
	goto L603
L603:
	;
	v2631 = *(*int32)(unsafe.Add(mBase, uint32(v2630)+64))
	if v2631 != 0 {
		goto L606
	} else {
		goto L607
	}
L604:
	;
	v2628 = F_create_sort_path(m, v1950, v2171, v2625, float64(-1))
	mBase = m.M
	v2629 = m.ExcPending
	if v2629 != 0 {
		goto L1
	} else {
		goto L605
	}
L605:
	;
	v2630 = v2628
	goto L603
L606:
	;
	v2632 = *(*int32)(unsafe.Add(mBase, uint32(v2631)+4))
	v2634 = v2632
	goto L608
L607:
	;
	v2634 = int32(0)
	goto L608
L608:
	;
	v2635 = F_create_unique_path(m, v1950, v2630, v2634, v2600)
	mBase = m.M
	v2636 = m.ExcPending
	if v2636 != 0 {
		goto L1
	} else {
		goto L609
	}
L609:
	;
	F_add_path(m, v1950, v2635)
	mBase = m.M
	v2638 = m.ExcPending
	if v2638 != 0 {
		goto L1
	} else {
		goto L610
	}
L610:
	;
	if v2468 == int32(0) {
		goto L599
	} else {
		goto L611
	}
L611:
	;
	v2641 = F_make_pathkeys_for_sortclauses(m, l1, v399, v230)
	mBase = m.M
	v2642 = m.ExcPending
	if v2642 != 0 {
		goto L1
	} else {
		goto L612
	}
L612:
	;
	v2644 = F_create_sort_path(m, v1950, v2468, v2641, float64(-1))
	mBase = m.M
	v2645 = m.ExcPending
	if v2645 != 0 {
		goto L1
	} else {
		goto L613
	}
L613:
	;
	v2646 = *(*int32)(unsafe.Add(mBase, uint32(v2644)+64))
	if v2646 != 0 {
		goto L614
	} else {
		goto L615
	}
L614:
	;
	v2647 = *(*int32)(unsafe.Add(mBase, uint32(v2646)+4))
	v2649 = v2647
	goto L616
L615:
	;
	v2649 = int32(0)
	goto L616
L616:
	;
	v2650 = F_create_unique_path(m, v1950, v2644, v2649, v2600)
	mBase = m.M
	v2651 = m.ExcPending
	if v2651 != 0 {
		goto L1
	} else {
		goto L617
	}
L617:
	;
	F_add_path(m, v1950, v2650)
	mBase = m.M
	v2653 = m.ExcPending
	if v2653 != 0 {
		goto L1
	} else {
		goto L618
	}
L618:
	;
	goto L599
L619:
	;
	v2665 = F_create_merge_append_path(m, l1, v1950, v1940, int32(0), v410)
	mBase = m.M
	v2666 = m.ExcPending
	if v2666 != 0 {
		goto L1
	} else {
		goto L620
	}
L620:
	;
	if v230 != 0 {
		goto L621
	} else {
		goto L622
	}
L621:
	;
	v2667 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	v2669 = v2667
	goto L623
L622:
	;
	v2669 = int32(0)
	goto L623
L623:
	;
	v2670 = F_create_unique_path(m, v1950, v2665, v2669, v2600)
	mBase = m.M
	v2671 = m.ExcPending
	if v2671 != 0 {
		goto L1
	} else {
		goto L624
	}
L624:
	;
	F_add_path(m, v1950, v2670)
	mBase = m.M
	v2673 = m.ExcPending
	if v2673 != 0 {
		goto L1
	} else {
		goto L625
	}
L625:
	;
	v2690 = v1950
	goto L4
L626:
	;
	if v2463 == int32(0) {
		v2690 = v1950
		goto L4
	} else {
		goto L627
	}
L627:
	;
	F_add_path(m, v1950, v2463)
	mBase = m.M
	v2679 = m.ExcPending
	if v2679 != 0 {
		goto L1
	} else {
		goto L628
	}
L628:
	;
	v2690 = v1950
	goto L4
L629:
	;
	v3072 = *(*int32)(unsafe.Add(mBase, _c_F_recurse_set_operations[2]))
	if v3072 != 0 {
		goto L675
	} else {
		goto L676
	}
L630:
	;
	if v2714 != 0 {
		goto L631
	} else {
		goto L632
	}
L631:
	;
	v2716 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if l4 != 0 {
		goto L635
	} else {
		goto L636
	}
L632:
	;
	goto L633
L633:
	;
	v2887 = int32(0)
	v2889 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v2892 = F_generate_setop_tlist(m, l3, l4, v2887, v2887, v2889, l5, v35+int32(112))
	mBase = m.M
	v2893 = m.ExcPending
	if v2893 != 0 {
		goto L1
	} else {
		goto L656
	}
L634:
	;
	if v2854 != 0 {
		goto L629
	} else {
		goto L655
	}
L635:
	;
	v2718 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v2720 = v2718
	goto L637
L636:
	;
	v2720 = int32(0)
	goto L637
L637:
	;
	if v2716 == int32(0) {
		v2788 = v2720
		goto L638
	} else {
		goto L639
	}
L638:
	;
	v2854 = base.B2i32(v2788 == int32(0))
	goto L634
L639:
	;
	v2723 = *(*int32)(unsafe.Add(mBase, uint32(v2716)+4))
	if v2723 <= int32(0) {
		v2788 = v2720
		goto L638
	} else {
		goto L640
	}
L640:
	;
	v2726 = v2720
	v2740 = int32(0)
	goto L641
L641:
	;
	v2758 = *(*int32)(unsafe.Add(mBase, uint32(v2716)+12))
	v2762 = *(*int32)(unsafe.Add(mBase, uint32(v2758+v2740<<(uint(int32(2))%32))))
	v2763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2762)+26)))
	if v2763 == int32(0) {
		goto L644
	} else {
		goto L645
	}
L642:
	;
	v2788 = v2782
	goto L638
L643:
	;
	v2785 = v2740 + int32(1)
	v2786 = *(*int32)(unsafe.Add(mBase, uint32(v2716)+4))
	if v2785 < v2786 {
		v2726 = v2782
		v2740 = v2785
		goto L641
	} else {
		goto L654
	}
L644:
	;
	if v2726 == int32(0) {
		goto L647
	} else {
		goto L648
	}
L645:
	;
	goto L646
L646:
	;
	v2854 = int32(0)
	goto L634
L647:
	;
	goto L646
L648:
	;
	v2768 = *(*int32)(unsafe.Add(mBase, uint32(v2762)+4))
	v2769 = F_exprCollation(m, v2768)
	mBase = m.M
	v2770 = m.ExcPending
	if v2770 != 0 {
		goto L1
	} else {
		goto L649
	}
L649:
	;
	v2771 = *(*int32)(unsafe.Add(mBase, uint32(v2726)))
	if v2769 != v2771 {
		goto L647
	} else {
		goto L650
	}
L650:
	;
	v2774 = v2726 + int32(4)
	v2776 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v2777 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if base.Ui32(v2774) < base.Ui32(v2776+v2777<<(uint(int32(2))%32)) {
		goto L651
	} else {
		goto L652
	}
L651:
	;
	v2782 = v2774
	goto L653
L652:
	;
	v2782 = int32(0)
	goto L653
L653:
	;
	goto L643
L654:
	;
	goto L642
L655:
	;
	goto L633
L656:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v2892
	v2895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+112)))
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v2895)
	v2897 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v2898 = F_make_pathtarget_from_tlist(m, v2897)
	mBase = m.M
	v2899 = m.ExcPending
	if v2899 != 0 {
		goto L1
	} else {
		goto L657
	}
L657:
	;
	v2900 = F_set_pathtarget_cost_width(m, l1, v2898)
	mBase = m.M
	v2901 = m.ExcPending
	if v2901 != 0 {
		goto L1
	} else {
		goto L658
	}
L658:
	;
	v2902 = *(*int32)(unsafe.Add(mBase, uint32(v2690)+44))
	if v2902 == int32(0) {
		goto L659
	} else {
		goto L660
	}
L659:
	;
	v2987 = *(*int32)(unsafe.Add(mBase, uint32(v2690)+52))
	if v2987 == int32(0) {
		goto L629
	} else {
		goto L669
	}
L660:
	;
	v2905 = *(*int32)(unsafe.Add(mBase, uint32(v2902)+4))
	if v2905 <= int32(0) {
		goto L659
	} else {
		goto L661
	}
L661:
	;
	v2911 = int32(0)
	goto L662
L662:
	;
	v2941 = *(*int32)(unsafe.Add(mBase, uint32(v2902)+12))
	v2944 = v2941 + v2911<<(uint(int32(2))%32)
	v2945 = *(*int32)(unsafe.Add(mBase, uint32(v2944)))
	v2946 = *(*int32)(unsafe.Add(mBase, uint32(v2945)+8))
	v2947 = F_apply_projection_to_path(m, l1, v2946, v2945, v2900)
	mBase = m.M
	v2948 = m.ExcPending
	if v2948 != 0 {
		goto L1
	} else {
		goto L664
	}
L663:
	;
	goto L659
L664:
	;
	if v2945 != v2947 {
		goto L665
	} else {
		goto L666
	}
L665:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2944))) = v2947
	goto L667
L666:
	;
	goto L667
L667:
	;
	v2952 = v2911 + int32(1)
	v2953 = *(*int32)(unsafe.Add(mBase, uint32(v2902)+4))
	if v2952 < v2953 {
		v2911 = v2952
		goto L662
	} else {
		goto L668
	}
L668:
	;
	goto L663
L669:
	;
	v2990 = int32(0)
	v2991 = *(*int32)(unsafe.Add(mBase, uint32(v2987)+4))
	if v2991 <= v2990 {
		goto L629
	} else {
		goto L670
	}
L670:
	;
	v2994 = v2990
	goto L671
L671:
	;
	v3026 = *(*int32)(unsafe.Add(mBase, uint32(v2987)+12))
	v3029 = v3026 + v2994<<(uint(int32(2))%32)
	v3030 = *(*int32)(unsafe.Add(mBase, uint32(v3029)))
	v3031 = *(*int32)(unsafe.Add(mBase, uint32(v3030)+8))
	v3032 = F_create_projection_path(m, l1, v3031, v3030, v2900)
	mBase = m.M
	v3033 = m.ExcPending
	if v3033 != 0 {
		goto L1
	} else {
		goto L673
	}
L672:
	;
	goto L629
L673:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3029))) = v3032
	v3036 = v2994 + int32(1)
	v3037 = *(*int32)(unsafe.Add(mBase, uint32(v2987)+4))
	if v3036 < v3037 {
		v2994 = v3036
		goto L671
	} else {
		goto L674
	}
L674:
	;
	goto L672
L675:
	;
	v3073 = int32(0)
	m.T0[v3072].(func(*base.Module, int32, int32, int32, int32, int32))(m, l1, v3073, v3073, v2690, v3073)
	mBase = m.M
	v3077 = m.ExcPending
	if v3077 != 0 {
		goto L1
	} else {
		goto L678
	}
L676:
	;
	goto L677
L677:
	;
	F_set_cheapest(m, v2690)
	mBase = m.M
	v3079 = m.ExcPending
	if v3079 != 0 {
		goto L1
	} else {
		goto L679
	}
L678:
	;
	goto L677
L679:
	;
	v3090 = v2690
	goto L3
}
func F_regconfigin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v117 int64
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
	var v130 int32
	_ = v130
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
	var v154 int64
	_ = v154
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	v6 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v13 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L29
	} else {
		goto L46
	}
L2:
	;
	m.G0 = v9 + int32(16)
	return v154
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
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v16 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v13-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v154 = v6
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
	v107 = F_strlen(m, v12)
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
	v46 = v12
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
	v106 = v46 - v12
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
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v73 == int32(0) {
		v96 = v12
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v96 - v12
	goto L9
L23:
	;
	v77 = v12
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
	v113 = F_DirectInputFunctionCallSafe(m, int32(588), v12, int32(-1), v11, v9+int32(8))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int64(0)
L30:
	;
	v117 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)))
	v154 = v117
	goto L2
L31:
	;
	v122 = F_stringToQualifiedNameList(m, v12, v11)
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
	v154 = v6
	goto L2
L34:
	;
	goto L35
L35:
	;
	v129 = F_get_ts_config_oid(m, v122, int32(1))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L29
	} else {
		goto L36
	}
L36:
	;
	if v129 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v133 = F_errsave_start(m, v11)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L29
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v154 = base.I64_extend_i32_u(v129)
	goto L2
L40:
	;
	if v133 == int32(0) {
		v154 = v6
		goto L2
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L29
	} else {
		goto L42
	}
L42:
	;
	v140 = F_NameListToString(m, v122)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v140
	F_errmsg(m, int32(_a_F_regconfigin_1), v9)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L29
	} else {
		goto L44
	}
L44:
	;
	F_errsave_finish(m, v11, int32(_a_F_regconfigin_2), int32(1358), int32(_a_F_regconfigin_3))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L29
	} else {
		goto L45
	}
L45:
	;
	v154 = v6
	goto L2
L46:
	;
	F_errmsg_internal(m, int32(_a_F_regconfigin_4), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L29
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_regconfigin_2), int32(1342), int32(_a_F_regconfigin_3))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L29
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regdictionaryin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v117 int64
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
	var v130 int32
	_ = v130
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
	var v154 int64
	_ = v154
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	v6 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v13 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L29
	} else {
		goto L46
	}
L2:
	;
	m.G0 = v9 + int32(16)
	return v154
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
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v16 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v13-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v154 = v6
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
	v107 = F_strlen(m, v12)
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
	v46 = v12
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
	v106 = v46 - v12
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
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v73 == int32(0) {
		v96 = v12
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v96 - v12
	goto L9
L23:
	;
	v77 = v12
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
	v113 = F_DirectInputFunctionCallSafe(m, int32(588), v12, int32(-1), v11, v9+int32(8))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int64(0)
L30:
	;
	v117 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)))
	v154 = v117
	goto L2
L31:
	;
	v122 = F_stringToQualifiedNameList(m, v12, v11)
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
	v154 = v6
	goto L2
L34:
	;
	goto L35
L35:
	;
	v129 = F_get_ts_dict_oid(m, v122, int32(1))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L29
	} else {
		goto L36
	}
L36:
	;
	if v129 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v133 = F_errsave_start(m, v11)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L29
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v154 = base.I64_extend_i32_u(v129)
	goto L2
L40:
	;
	if v133 == int32(0) {
		v154 = v6
		goto L2
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L29
	} else {
		goto L42
	}
L42:
	;
	v140 = F_NameListToString(m, v122)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v140
	F_errmsg(m, int32(_a_F_regdictionaryin_1), v9)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L29
	} else {
		goto L44
	}
L44:
	;
	F_errsave_finish(m, v11, int32(_a_F_regdictionaryin_2), int32(1468), int32(_a_F_regdictionaryin_3))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L29
	} else {
		goto L45
	}
L45:
	;
	v154 = v6
	goto L2
L46:
	;
	F_errmsg_internal(m, int32(_a_F_regdictionaryin_4), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L29
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_regdictionaryin_2), int32(1452), int32(_a_F_regdictionaryin_3))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L29
	} else {
		goto L48
	}
L48:
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
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v412 int32
	_ = v412
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v479 int32
	_ = v479
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
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
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v599 int32
	_ = v599
	var v602 float64
	_ = v602
	var v603 int32
	_ = v603
	var v605 float64
	_ = v605
	var v606 int32
	_ = v606
	var v610 float64
	_ = v610
	var v618 float64
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v650 int32
	_ = v650
	var v653 float64
	_ = v653
	var v654 int32
	_ = v654
	var v656 float64
	_ = v656
	var v657 int32
	_ = v657
	var v661 float64
	_ = v661
	var v666 float64
	_ = v666
	var v670 float64
	_ = v670
	var v671 float64
	_ = v671
	var v673 float64
	_ = v673
	var v681 float64
	_ = v681
	var v684 int32
	_ = v684
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
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
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v710 = m.ExcPending
	if v710 != 0 {
		goto L4
	} else {
		goto L148
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
		v479 = v51
		goto L15
	} else {
		goto L16
	}
L10:
	;
	m.G0 = v40 + int32(128)
	if v570 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L11:
	;
	v549 = *(*int32)(unsafe.Add(mBase, _c_F_regex_fixed_prefix[0]))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)+4))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v550*int32(28))+uint32(_c_F_regex_fixed_prefix[1])))
	goto L93
L12:
	;
	v546 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43))) = uint8(v546)
	goto L11
L13:
	;
	v527 = v40 + int32(16)
	F_pg_regerror(m, v523, v527)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L4
	} else {
		goto L88
	}
L14:
	;
	switch v523 + int32(2) {
	case 0:
		goto L12
	case 1:
		goto L11
	default:
		goto L13
	case 3:
		v570 = int32(0)
		goto L10
	}
L15:
	;
	v523 = v479
	goto L14
L16:
	;
	v61 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v61
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_regex_fixed_prefix[2]))
	if v66 != int32(_a_F_regex_fixed_prefix_0) {
		v479 = v51
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_regex_fixed_prefix[3]))
	if v71 != int32(4) {
		v523 = int32(17)
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
		v479 = v78
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+44)))
	if v85&int32(2) != 0 {
		v479 = v78
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
		v523 = int32(12)
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
		v445 = v78
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if base.Ui32(int32(-3)) < base.Ui32(v445) {
		v479 = v445
		goto L15
	} else {
		goto L86
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
		v445 = v78
		goto L24
	} else {
		goto L38
	}
L28:
	;
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+22)))
	if v139 != v141 {
		v445 = v78
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
		v445 = v78
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
	if v316 == int32(_a_F_regex_fixed_prefix_1) {
		goto L68
	} else {
		goto L69
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
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v308+v164<<(uint(int32(2))%32))))
	v313 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v312))))
	v316 = v313
	v320 = v312
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
		v316 = v185
		v320 = v184
		goto L39
	} else {
		goto L49
	}
L49:
	;
	v226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+26)))
	if v217 == v226 {
		v316 = v185
		v320 = v184
		goto L39
	} else {
		goto L50
	}
L50:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v228 <= base.I32_extend16_s(v192) {
		v316 = v185
		v320 = v184
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
		v316 = v185
		v320 = v184
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
	if v269&int32(_a_F_regex_fixed_prefix_1) != v247 {
		goto L42
	} else {
		goto L65
	}
L61:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v101)+24))
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262+v259<<(uint(int32(1))%32)))))
	v269 = v266
	goto L60
L62:
	;
	goto L63
L63:
	;
	v267 = F_pg_reg_getcolor(m, v101, v259)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v269 = v267
	goto L60
L65:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v273 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v94+v273<<(uint(int32(2))%32)))) = v259
	if v239 != int32(-1) {
		v164 = v239
		goto L40
	} else {
		goto L66
	}
L66:
	;
	goto L42
L67:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	if v412 == v433 {
		v445 = int32(-2)
		goto L24
	} else {
		goto L82
	}
L68:
	;
	v412 = int32(-1)
	goto L67
L69:
	;
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+24)))
	v345 = v316
	v348 = int32(-1)
	v349 = v320
	goto L70
L70:
	;
	if v345 != v341 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v412 = v375
	goto L67
L72:
	;
	v369 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+26)))
	if v345 != v369 {
		goto L68
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	if v348 == int32(-1) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L74
L76:
	;
	v376 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v349)+8)))
	if v376 != int32(_a_F_regex_fixed_prefix_1) {
		v345 = v376
		v348 = v375
		v349 = v349 + int32(8)
		goto L70
	} else {
		goto L81
	}
L77:
	;
	v375 = v371
	goto L76
L78:
	;
	goto L79
L79:
	;
	if v371 != v348 {
		goto L68
	} else {
		goto L80
	}
L80:
	;
	v375 = v348
	goto L76
L81:
	;
	goto L71
L82:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v437 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v438 = int32(-1)
	goto L85
L84:
	;
	v438 = int32(1)
	goto L85
L85:
	;
	v445 = v438
	goto L24
L86:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	F_pfree(m, v466)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	v469 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v469
	v479 = v445
	goto L15
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v527
	F_errmsg(m, int32(_a_F_regex_fixed_prefix_3), v40)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_regex_fixed_prefix_4), int32(2069), int32(_a_F_regex_fixed_prefix_5))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v40)+120))
	v560 = F_palloc(m, v555*v556+int32(1))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v40)+124))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v40)+120))
	v564 = F_pg_wchar2mb_with_len(m, v562, v560, v563)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+120)) = v564
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v40)+124))
	F_pfree(m, v567)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	v570 = v560
	goto L10
L97:
	;
	m.G0 = v28 + int32(16)
	return v699
L98:
	;
	v576 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v576
	if l4 == v576 {
		v699 = v576
		goto L97
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v622 = F_string_to_const(m, v570, v30)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L4
	} else {
		goto L117
	}
L101:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v583 = F_text_to_cstring(m, v582)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L4
	} else {
		goto L105
	}
L102:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = v618
	F_pfree(m, v583)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L4
	} else {
		goto L116
	}
L103:
	;
	if base.F64_lt(v610, float64(0)) != 0 {
		v618 = float64(0)
		goto L102
	} else {
		goto L114
	}
L104:
	;
	v605 = F_regex_selectivity_sub(m, v583, v585)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L4
	} else {
		goto L113
	}
L105:
	;
	v585 = F_strlen(m, v583)
	mBase = m.M
	if v585 <= int32(0) {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v589 = v585 - int32(1)
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583+v589))))
	if v591 != int32(36) {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	if v585 != int32(1) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585+v583-int32(2)))))
	if v599 == int32(92) {
		goto L104
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v602 = F_regex_selectivity_sub(m, v583, v589)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L4
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	v610 = v602
	goto L103
L113:
	;
	v610 = base.F64_mul(v605, float64(5))
	goto L103
L114:
	;
	if base.F64_gt(v610, float64(1)) == int32(0) {
		v618 = v610
		goto L102
	} else {
		goto L115
	}
L115:
	;
	v618 = float64(1)
	goto L102
L116:
	;
	v699 = v576
	goto L97
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v622
	if l4 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	F_pfree(m, v570)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L4
	} else {
		goto L144
	}
L119:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+15)))
	if v627 == int32(1) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(4607182418800017408)
	goto L118
L121:
	;
	goto L122
L122:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v633 = F_text_to_cstring(m, v632)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	v635 = F_strlen(m, v633)
	mBase = m.M
	v636 = F_strlen(m, v570)
	mBase = m.M
	if v635 <= int32(0) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	if int32(0) < v636 {
		goto L135
	} else {
		goto L136
	}
L125:
	;
	v656 = F_regex_selectivity_sub(m, v633, v635)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L4
	} else {
		goto L133
	}
L126:
	;
	v640 = v635 - int32(1)
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633+v640))))
	if v642 != int32(36) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	if v635 != int32(1) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v635+v633-int32(2)))))
	if v650 == int32(92) {
		goto L125
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v653 = F_regex_selectivity_sub(m, v633, v640)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L4
	} else {
		goto L132
	}
L131:
	;
	goto L130
L132:
	;
	v661 = v653
	goto L124
L133:
	;
	v661 = base.F64_mul(v656, float64(5))
	goto L124
L134:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = v681
	F_pfree(m, v633)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L4
	} else {
		goto L143
	}
L135:
	;
	v666 = F_pow(m, float64(0.2), base.F64_convert_i32_u(v636))
	mBase = m.M
	if base.F64_gt(v666, float64(0)) != 0 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v671 = v661
	goto L137
L137:
	;
	v673 = float64(0)
	if base.F64_lt(v671, v673) != 0 {
		v681 = v673
		goto L134
	} else {
		goto L141
	}
L138:
	;
	v670 = base.F64_div(v661, v666)
	goto L140
L139:
	;
	v670 = v661
	goto L140
L140:
	;
	v671 = v670
	goto L137
L141:
	;
	if base.F64_gt(v671, float64(1)) == int32(0) {
		v681 = v671
		goto L134
	} else {
		goto L142
	}
L142:
	;
	v681 = float64(1)
	goto L134
L143:
	;
	goto L118
L144:
	;
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+15)))
	if v695 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v696 = int32(2)
	goto L147
L146:
	;
	v696 = int32(1)
	goto L147
L147:
	;
	v699 = v696
	goto L97
L148:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	F_errmsg(m, int32(_a_F_regex_fixed_prefix_6), int32(0))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_regex_fixed_prefix_7), int32(1192), int32(_a_F_regex_fixed_prefix_8))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+384))
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
	v36 = int32(0)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v39 <= v36 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v42 = v25 + int32(8)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v49 = v46 & int32(-1)
	if v49 != 0 {
		v71 = v49
		v72 = v36
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v50 = int32(1)
	if v50 == v39 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v54 = v50
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
	v115 = int32(base.Ui32(v113) >> (uint(int32(5)) % 32))
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
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v26)+60))
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
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
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
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v238 != 0 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+80))
	if v240 != 0 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
	if v241 != 0 {
		goto L47
	} else {
		goto L51
	}
L51:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
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
	v246 = F_expression_tree_walker_impl(m, v233, int32(888), v17)
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
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+80))
	if v266 != 0 {
		v271 = v263
		goto L45
	} else {
		goto L61
	}
L61:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
	if v267 != 0 {
		v271 = v263
		goto L45
	} else {
		goto L62
	}
L62:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
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
	v280 = F_expression_tree_walker_impl(m, v263, int32(888), v17)
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
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
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
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v318
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
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v376
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
func F_regrolein(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int64
	_ = v177
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	v6 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v13 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L29
	} else {
		goto L53
	}
L2:
	;
	m.G0 = v9 + int32(16)
	return v177
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
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v16 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v13-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v177 = v6
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
	v107 = F_strlen(m, v12)
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
	v46 = v12
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
	v106 = v46 - v12
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
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v73 == int32(0) {
		v96 = v12
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v96 - v12
	goto L9
L23:
	;
	v77 = v12
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
	v113 = F_DirectInputFunctionCallSafe(m, int32(588), v12, int32(-1), v11, v9+int32(8))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int64(0)
L30:
	;
	v117 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)))
	v177 = v117
	goto L2
L31:
	;
	v122 = F_stringToQualifiedNameList(m, v12, v11)
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
	v177 = v6
	goto L2
L34:
	;
	goto L35
L35:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v128 != int32(1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v131 = F_errsave_start(m, v11)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L29
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v151 = F_get_role_oid(m, v149, int32(1))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L29
	} else {
		goto L44
	}
L39:
	;
	if v131 == int32(0) {
		v177 = v6
		goto L2
	} else {
		goto L40
	}
L40:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L29
	} else {
		goto L41
	}
L41:
	;
	F_errmsg(m, int32(_a_F_regrolein_1), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L29
	} else {
		goto L42
	}
L42:
	;
	F_errsave_finish(m, v11, int32(_a_F_regrolein_2), int32(1572), int32(_a_F_regrolein_3))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	v177 = v6
	goto L2
L44:
	;
	if v151 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v155 = F_errsave_start(m, v11)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L29
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v177 = base.I64_extend_i32_u(v151)
	goto L2
L48:
	;
	if v155 == int32(0) {
		v177 = v6
		goto L2
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v164
	F_errmsg(m, int32(_a_F_regrolein_4), v9)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L29
	} else {
		goto L51
	}
L51:
	;
	F_errsave_finish(m, v11, int32(_a_F_regrolein_2), int32(1580), int32(_a_F_regrolein_3))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L29
	} else {
		goto L52
	}
L52:
	;
	v177 = v6
	goto L2
L53:
	;
	F_errmsg_internal(m, int32(_a_F_regrolein_5), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L29
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_regrolein_2), int32(1562), int32(_a_F_regrolein_3))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L29
	} else {
		goto L55
	}
L55:
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
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
	v5 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+4)))
	v6 = F_GetSysCacheHashValue(m, int32(7), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_repalloc_mul_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14373(m, l0, l1, l2, int32(2))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
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
	var v69 int32
	_ = v69
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
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
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v14)+288))
	F_pfree(m, v445)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L33
	} else {
		goto L134
	}
L3:
	;
	if int32(2) <= v260 {
		goto L121
	} else {
		goto L122
	}
L4:
	;
	v352 = F_getObjectDescription(m, l3, int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L33
	} else {
		goto L114
	}
L5:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v320 = F_getObjectDescription(m, v33+int32(4), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
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
		v99 = v62
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
	v106 = m.ExcPending
	if v106 != 0 {
		goto L33
	} else {
		goto L34
	}
L19:
	;
	if v99 == int32(0) {
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
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66<<(uint(int32(2))%32))+uint32(_c_F_reportDependentObjects[1])))
	if base.Ui32(v57-int32(15)) <= base.Ui32(int32(1)) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v82 = int32(0)
	if v57 == int32(16) {
		v99 = v82
		goto L20
	} else {
		goto L29
	}
L23:
	;
	if int32(22) <= v69 {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if base.B2i32(v57 == int32(20))|base.B2i32(v69 == int32(15)) != 0 {
		goto L22
	} else {
		goto L27
	}
L26:
	;
	v99 = v62
	goto L20
L27:
	;
	if v69 <= v57 {
		v99 = v62
		goto L20
	} else {
		goto L28
	}
L28:
	;
	goto L22
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_reportDependentObjects[2]))
	if v86 != int32(2) {
		v99 = v82
		goto L20
	} else {
		goto L30
	}
L30:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_reportDependentObjects[3])))
	if v90&int32(1) != 0 {
		v99 = v82
		goto L20
	} else {
		goto L31
	}
L31:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_reportDependentObjects[4]))
	v99 = base.B2i32(v57 == int32(17)) | base.B2i32(v96 <= v57)
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
	v110 = m.ExcPending
	if v110 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v113 = v111 - int32(1)
	if v113 < int32(0) {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v117 = int32(0)
	v121 = v117
	v124 = v113
	v125 = v117
	v126 = int32(1)
	goto L37
L37:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v133 = v130 + v124<<(uint(int32(4))%32)
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133))))
	if v134&int32(257) != 0 {
		v258 = v121
		v260 = v125
		v261 = v126
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if int32(0) < v258 {
		goto L90
	} else {
		goto L91
	}
L39:
	;
	if int32(0) < v124 {
		v121 = v258
		v124 = v124 - int32(1)
		v125 = v260
		v126 = v261
		goto L37
	} else {
		goto L88
	}
L40:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v142 = F_getObjectDescription(m, v137+v124*int32(12), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L33
	} else {
		goto L41
	}
L41:
	;
	if v142 == int32(0) {
		v258 = v121
		v260 = v125
		v261 = v126
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v146&int32(60) != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	F_pfree(m, v142)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L33
	} else {
		goto L87
	}
L44:
	;
	v151 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
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
	if v151 == int32(0) {
		v253 = v121
		v254 = v125
		v255 = v126
		goto L43
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+256)) = v142
	F_errmsg_internal(m, int32(_a_F_reportDependentObjects_0), v14+int32(256))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L33
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_reportDependentObjects_1), int32(1099), int32(_a_F_reportDependentObjects_2))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L33
	} else {
		goto L50
	}
L50:
	;
	v253 = v121
	v254 = v125
	v255 = v126
	goto L43
L51:
	;
	v171 = F_getObjectDescription(m, v133+int32(4), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L33
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v125 <= int32(99) {
		goto L74
	} else {
		goto L75
	}
L54:
	;
	if v171 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	if v125 <= int32(99) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	goto L57
L57:
	;
	v253 = v121 + int32(1)
	v254 = v125
	v255 = int32(0)
	goto L43
L58:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v14)+276))
	if v196 != 0 {
		goto L67
	} else {
		goto L68
	}
L59:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v14)+292))
	if v175 != 0 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L61
L61:
	;
	v194 = v121 + int32(1)
	v195 = v125
	goto L58
L62:
	;
	F_appendStringInfoChar(m, v14+int32(288), int32(10))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L33
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+212)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = v142
	F_appendStringInfo(m, v14+int32(288), int32(_a_F_reportDependentObjects_3), v14+int32(208))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L33
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	v194 = v121
	v195 = v125 + int32(1)
	goto L58
L67:
	;
	F_appendStringInfoChar(m, v14+int32(272), int32(10))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L33
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+196)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v14)+192)) = v142
	F_appendStringInfo(m, v14+int32(272), int32(_a_F_reportDependentObjects_3), v14+int32(192))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L33
	} else {
		goto L71
	}
L70:
	;
	goto L69
L71:
	;
	F_pfree(m, v171)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L33
	} else {
		goto L72
	}
L72:
	;
	v253 = v194
	v254 = v195
	v255 = int32(0)
	goto L43
L73:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v14)+276))
	if v239 != 0 {
		goto L82
	} else {
		goto L83
	}
L74:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v14)+292))
	if v219 != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v237 = v121 + int32(1)
	v238 = v125
	goto L73
L77:
	;
	F_appendStringInfoChar(m, v14+int32(288), int32(10))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L33
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v142
	F_appendStringInfo(m, v14+int32(288), int32(_a_F_reportDependentObjects_4), v14+int32(240))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L33
	} else {
		goto L81
	}
L80:
	;
	goto L79
L81:
	;
	v237 = v121
	v238 = v125 + int32(1)
	goto L73
L82:
	;
	F_appendStringInfoChar(m, v14+int32(272), int32(10))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L33
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+224)) = v142
	F_appendStringInfo(m, v14+int32(272), int32(_a_F_reportDependentObjects_4), v14+int32(224))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L33
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	v253 = v237
	v254 = v238
	v255 = v126
	goto L43
L87:
	;
	v258 = v253
	v260 = v254
	v261 = v255
	goto L39
L88:
	;
	goto L38
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L33
	} else {
		goto L99
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+176)) = v258
	if v258 == int32(1) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	if v261 != 0 {
		goto L3
	} else {
		goto L98
	}
L93:
	;
	v275 = int32(_a_F_reportDependentObjects_5)
	goto L95
L94:
	;
	v275 = int32(_a_F_reportDependentObjects_6)
	goto L95
L95:
	;
	F_appendStringInfo(m, v14+int32(288), v275, v14+int32(176))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L33
	} else {
		goto L96
	}
L96:
	;
	if v261 == int32(0) {
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
	v288 = m.ExcPending
	if v288 != 0 {
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
	v292 = m.ExcPending
	if v292 != 0 {
		goto L33
	} else {
		goto L102
	}
L102:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v14)+288))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v293
	F_errdetail_internal(m, int32(_a_F_reportDependentObjects_8), v14+int32(112))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L33
	} else {
		goto L103
	}
L103:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v14)+272))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v300
	F_errdetail_log(m, int32(_a_F_reportDependentObjects_8), v14+int32(96))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L33
	} else {
		goto L104
	}
L104:
	;
	F_errhint(m, int32(_a_F_reportDependentObjects_9), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L33
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_reportDependentObjects_1), int32(1177), int32(_a_F_reportDependentObjects_2))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
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
	v325 = m.ExcPending
	if v325 != 0 {
		goto L33
	} else {
		goto L108
	}
L108:
	;
	F_errcode(m, int32(16909442))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L33
	} else {
		goto L109
	}
L109:
	;
	v333 = F_getObjectDescription(m, v316+v28*int32(12), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L33
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v320
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v333
	F_errmsg(m, int32(_a_F_reportDependentObjects_10), v14+int32(16))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L33
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v320
	F_errhint(m, int32(_a_F_reportDependentObjects_11), v14)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L33
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_reportDependentObjects_1), int32(1034), int32(_a_F_reportDependentObjects_2))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = v352
	F_errmsg(m, int32(_a_F_reportDependentObjects_12), v14+int32(160))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L33
	} else {
		goto L115
	}
L115:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v14)+288))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = v360
	F_errdetail_internal(m, int32(_a_F_reportDependentObjects_8), v14+int32(144))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L33
	} else {
		goto L116
	}
L116:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v14)+272))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v367
	F_errdetail_log(m, int32(_a_F_reportDependentObjects_8), v14+int32(128))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L33
	} else {
		goto L117
	}
L117:
	;
	F_errhint(m, int32(_a_F_reportDependentObjects_9), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L33
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_reportDependentObjects_1), int32(1170), int32(_a_F_reportDependentObjects_2))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
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
	F_errfinish(m, int32(_a_F_reportDependentObjects_1), v430, int32(_a_F_reportDependentObjects_2))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L33
	} else {
		goto L133
	}
L121:
	;
	v387 = F_errstart(m, v57, int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L33
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	if v260 != int32(1) {
		goto L2
	} else {
		goto L129
	}
L124:
	;
	if v387 == int32(0) {
		goto L2
	} else {
		goto L125
	}
L125:
	;
	v391 = v258 + v260
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v391
	F_errmsg_plural(m, int32(_a_F_reportDependentObjects_13), int32(_a_F_reportDependentObjects_14), v391, v14-int32(-64))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L33
	} else {
		goto L126
	}
L126:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v14)+288))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v399
	F_errdetail_internal(m, int32(_a_F_reportDependentObjects_8), v14+int32(48))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L33
	} else {
		goto L127
	}
L127:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v14)+272))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v406
	F_errdetail_log(m, int32(_a_F_reportDependentObjects_8), v14+int32(32))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L33
	} else {
		goto L128
	}
L128:
	;
	v430 = int32(1187)
	goto L120
L129:
	;
	v417 = F_errstart(m, v57, int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L33
	} else {
		goto L130
	}
L130:
	;
	if v417 == int32(0) {
		goto L2
	} else {
		goto L131
	}
L131:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v14)+288))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v421
	F_errmsg_internal(m, int32(_a_F_reportDependentObjects_8), v14+int32(80))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L33
	} else {
		goto L132
	}
L132:
	;
	v430 = int32(1193)
	goto L120
L133:
	;
	goto L2
L134:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v14)+272))
	F_pfree(m, v448)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
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
	v8 = Fn14349(m, int64(32))
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
						F_errfinish(m, int32(_a_F_resolve_anyarray_from_others_1), int32(676), int32(_a_F_resolve_anyarray_from_others_2))
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
										F_errfinish(m, int32(_a_F_resolve_anyarray_from_others_1), int32(672), int32(_a_F_resolve_anyarray_from_others_2))
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
								F_errfinish(m, int32(_a_F_resolve_anyarray_from_others_1), int32(672), int32(_a_F_resolve_anyarray_from_others_2))
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
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
	return v232 & int32(1)
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
	v35 = int32(8)
	v38 = F_palloc_mul(m, int32(4), v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L11
	}
L7:
	;
	if v22 == int32(0) {
		v232 = v2
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
	v232 = v2
	goto L1
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_rmtree[0])) = int32(0)
	v43 = int32(1)
	v44 = F_readdir(m, v14)
	mBase = m.M
	if v44 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v47 = v44
	v48 = v35
	v49 = v43
	v51 = v38
	v52 = v2
	goto L15
L13:
	;
	v142 = v43
	v144 = v38
	v145 = v2
	goto L14
L14:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_rmtree[0]))
	if v148 == int32(0) {
		v169 = v142
		goto L39
	} else {
		goto L40
	}
L15:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+19)))
	if v54 != int32(46) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v142 = v130
	v144 = v131
	v145 = v132
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_rmtree[0])) = int32(0)
	v137 = F_readdir(m, v14)
	mBase = m.M
	if v137 != 0 {
		v47 = v137
		v48 = v129
		v49 = v130
		v51 = v131
		v52 = v132
		goto L15
	} else {
		goto L38
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v47 + int32(19)
	v71 = v12 + int32(80)
	v76 = F_pg_snprintf(m, v71, int32(1024), int32(_a_F_rmtree_3), v12-int32(-64))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L23
	}
L19:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+20)))
	if v57 == int32(0) {
		v129 = v48
		v130 = v49
		v131 = v51
		v132 = v52
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+20)))
	if v60 != int32(46) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+21)))
	if v63 == int32(0) {
		v129 = v48
		v130 = v49
		v131 = v51
		v132 = v52
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	v80 = F_get_dirent_type(m, v71, v47, int32(0), int32(19))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L26
	}
L24:
	;
	v102 = v12 + int32(80)
	v103 = F_unlink(m, v102)
	mBase = m.M
	if v103 == int32(0) {
		v129 = v48
		v130 = v49
		v131 = v51
		v132 = v52
		goto L17
	} else {
		goto L32
	}
L25:
	;
	if v48 == v52 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	switch v80 {
	case 0:
		v129 = v48
		v130 = v49
		v131 = v51
		v132 = v52
		goto L17
	default:
		goto L24
	case 3:
		goto L25
	}
L27:
	;
	v85 = F_repalloc(m, v51, v48<<(uint(int32(3))%32))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L30
	}
L28:
	;
	v89 = v48
	v90 = v51
	goto L29
L29:
	;
	v96 = F_pstrdup(m, v12+int32(80))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L2
	} else {
		goto L31
	}
L30:
	;
	v89 = v48 << (uint(int32(1)) % 32)
	v90 = v85
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90+v52<<(uint(int32(2))%32)))) = v96
	v129 = v89
	v130 = v49
	v131 = v90
	v132 = v52 + int32(1)
	goto L17
L32:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_rmtree[0]))
	if v107 == int32(44) {
		v129 = v48
		v130 = v49
		v131 = v51
		v132 = v52
		goto L17
	} else {
		goto L33
	}
L33:
	;
	v110 = int32(0)
	v113 = F_errstart(m, int32(19), v110)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	if v113 == int32(0) {
		v129 = v48
		v130 = v110
		v131 = v51
		v132 = v52
		goto L17
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v102
	F_errmsg_internal(m, int32(_a_F_rmtree_4), v12+int32(48))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_rmtree_1), int32(97), int32(_a_F_rmtree_2))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v129 = v48
	v130 = v110
	v131 = v51
	v132 = v52
	goto L17
L38:
	;
	goto L16
L39:
	;
	F_FreeDir(m, v14)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L2
	} else {
		goto L45
	}
L40:
	;
	v151 = int32(0)
	v154 = F_errstart(m, int32(19), v151)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	if v154 == int32(0) {
		v169 = v151
		goto L39
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l0
	F_errmsg_internal(m, int32(_a_F_rmtree_5), v12+int32(32))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_rmtree_1), int32(106), int32(_a_F_rmtree_2))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v169 = v151
	goto L39
L45:
	;
	if v145 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v175 = int32(0)
	v177 = v169
	goto L49
L47:
	;
	v199 = v169
	goto L48
L48:
	;
	v204 = F_rmdir(m, l0)
	mBase = m.M
	if v204 == int32(0) {
		v225 = v199
		goto L54
	} else {
		goto L55
	}
L49:
	;
	v184 = v144 + v175<<(uint(int32(2))%32)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v186 = F_rmtree(m, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L2
	} else {
		goto L51
	}
L50:
	;
	v199 = v191
	goto L48
L51:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	F_pfree(m, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v191 = v177 & v186
	v193 = v175 + int32(1)
	if v193 != v145 {
		v175 = v193
		v177 = v191
		goto L49
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	F_pfree(m, v144)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L2
	} else {
		goto L60
	}
L55:
	;
	v207 = int32(0)
	v210 = F_errstart(m, int32(19), v207)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	if v210 == int32(0) {
		v225 = v207
		goto L54
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_rmtree_6), v12+int32(16))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_rmtree_1), int32(124), int32(_a_F_rmtree_2))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v225 = v207
	goto L54
L60:
	;
	v232 = v225
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var __phi129 int32
	_ = __phi129
	var v130 int32
	_ = v130
	var __phi130 int32
	_ = __phi130
	var v131 int32
	_ = v131
	var __phi131 int32
	_ = __phi131
	var v134 int32
	_ = v134
	var __phi134 int32
	_ = __phi134
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v174 int32
	_ = v174
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
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
	return v318
L2:
	;
	if l0 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L3:
	;
	v25 = m.G0
	v27 = v25 - int32(48)
	m.G0 = v27
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v30 = F_hash_bytes_extended(m, v8+int32(12), int32(4), v29)
	mBase = m.M
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v33 = v31 - int32(1)
	v35 = v33 & base.I32_wrap_i64(v30)
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if int32(2) <= v37 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v222 = l2
	goto L5
L5:
	;
	v223 = int32(0)
	if l0 == v223 {
		goto L30
	} else {
		goto L31
	}
L6:
	;
	if v209 != 0 {
		goto L2
	} else {
		goto L28
	}
L7:
	;
	m.G0 = v27 + int32(48)
	goto L6
L8:
	;
	v174 = int32(0)
	goto L24
L9:
	;
	v41 = v37 - int32(1)
	v42 = int32(3)
	v43 = v41 & v42
	v46 = base.I32_wrap_i64(int64(base.Ui64(v30) >> (uint(int64(32)) % 64)))
	if base.Ui32(v37-int32(2)) < base.Ui32(v42) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	if v37 != int32(1) {
		v209 = int32(0)
		goto L7
	} else {
		goto L23
	}
L12:
	;
	__phi129 = v115
	__phi130 = v116
	__phi131 = v117
	__phi134 = int32(0)
	v129 = __phi129
	v130 = __phi130
	v131 = __phi131
	v134 = __phi134
	goto L20
L13:
	;
	v115 = int32(1)
	v116 = v46
	v117 = v35
	goto L12
L14:
	;
	goto L15
L15:
	;
	v57 = int32(1)
	v58 = v46
	v59 = v35
	v62 = int32(0)
	goto L16
L16:
	;
	v68 = int32(2)
	v71 = v33 & v58
	v73 = (v59 + v71) & v33
	*(*int32)(unsafe.Add(mBase, uint32(v27+v57<<(uint(v68)%32)))) = v73
	v76 = v57 + int32(1)
	v81 = (v57 + v71) & v33
	v83 = (v81 + v73) & v33
	*(*int32)(unsafe.Add(mBase, uint32(v27+v76<<(uint(v68)%32)))) = v83
	v86 = v57 + v68
	v91 = (v81 + v76) & v33
	v93 = (v91 + v83) & v33
	*(*int32)(unsafe.Add(mBase, uint32(v27+v86<<(uint(v68)%32)))) = v93
	v96 = v57 + int32(3)
	v101 = (v91 + v86) & v33
	v103 = (v93 + v101) & v33
	*(*int32)(unsafe.Add(mBase, uint32(v27+v96<<(uint(v68)%32)))) = v103
	v105 = v101 + v96
	v106 = int32(4)
	v107 = v57 + v106
	v109 = v62 + v106
	if v109 != v41&int32(-4) {
		v57 = v107
		v58 = v105
		v59 = v103
		v62 = v109
		goto L16
	} else {
		goto L18
	}
L17:
	;
	if v43 == int32(0) {
		goto L8
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v115 = v107
	v116 = v105
	v117 = v103
	goto L12
L20:
	;
	v143 = v33 & v130
	v145 = (v143 + v131) & v33
	*(*int32)(unsafe.Add(mBase, uint32(v27+v129<<(uint(int32(2))%32)))) = v145
	v148 = int32(1)
	v151 = v134 + v148
	if v151 != v43 {
		__phi129 = v129 + v148
		__phi130 = v129 + v143
		__phi131 = v145
		__phi134 = v151
		v129 = __phi129
		v130 = __phi130
		v131 = __phi131
		v134 = __phi134
		goto L20
	} else {
		goto L22
	}
L21:
	;
	goto L8
L22:
	;
	goto L21
L23:
	;
	goto L8
L24:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v27+v174<<(uint(int32(2))%32))))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(24)+int32(base.Ui32(v188)>>(uint(int32(3))%32))))))
	v199 = base.B2i32(int32(base.Ui32(v192)>>(uint(v188&int32(7))%32))&int32(1) == int32(0))
	if int32(base.Ui32(v192)>>(uint(v188&int32(7))%32))&int32(1) == int32(0) {
		v209 = v199
		goto L7
	} else {
		goto L26
	}
L25:
	;
	v209 = v199
	goto L7
L26:
	;
	v203 = v174 + int32(1)
	if v203 != v37 {
		v174 = v203
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v222 = v221
	goto L5
L29:
	;
	if v261 != 0 {
		v318 = l0
		goto L1
	} else {
		goto L42
	}
L30:
	;
	v261 = int32(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v229 <= int32(0) {
		v255 = v223
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v261 = v255
	goto L29
L34:
	;
	v232 = int32(0)
	if v232 < v229 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v235 = v229
	goto L37
L36:
	;
	v235 = v232
	goto L37
L37:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v238 = int32(0)
	goto L38
L38:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v236+v238<<(uint(int32(2))%32))))
	v247 = base.B2i32(v246 == v222)
	if v246 == v222 {
		v255 = v247
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v255 = v247
	goto L33
L40:
	;
	v249 = v238 + int32(1)
	if v249 != v235 {
		v238 = v249
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L2
L43:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v309 = F_lappend_oid(m, l0, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L47
	} else {
		goto L53
	}
L44:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v264 != 0 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v265 < int32(1025) {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v271 = *(*int32)(unsafe.Add(mBase, _c_F_roles_list_append[0]))
	v273 = F_bloom_create(m, int64(10240), v271, int64(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	return int32(0)
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v273
	v278 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v278
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v280 <= v278 {
		goto L43
	} else {
		goto L49
	}
L49:
	;
	v285 = int32(0)
	goto L50
L50:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v288+v285<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v292
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_bloom_add_element(m, v294, v8+int32(8), int32(4))
	mBase = m.M
	v300 = v285 + int32(1)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v300 < v301 {
		v285 = v300
		goto L50
	} else {
		goto L52
	}
L51:
	;
	goto L43
L52:
	;
	goto L51
L53:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v311 == int32(0) {
		v318 = v309
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_bloom_add_element(m, v311, v8+int32(12), int32(4))
	mBase = m.M
	v318 = v309
	goto L1
}
func F_rpad(m *base.Module, l0 int32) int64 {
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
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int64
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	v93 = v89
	goto L34
L33:
	;
	v93 = int32(4)
	goto L34
L34:
	;
	v95 = F_pg_mbstrlen_with_len(m, v13+v93, v58)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_rpad[0]))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99*int32(28))+uint32(_c_F_rpad[1])))
	goto L37
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L79
	}
L37:
	;
	if v95 < v24 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v107 = v95
	goto L40
L39:
	;
	v107 = v24
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
	v110 = v107
	goto L43
L42:
	;
	v110 = v24
	goto L43
L43:
	;
	v112 = base.I64_extend_i32_s(v104) * base.I64_extend_i32_s(v110)
	v116 = base.I32_wrap_i64(v112)
	if base.I32_wrap_i64(int64(base.Ui64(v112)>>(uint(int64(32))%64))) != v116>>(uint(int32(31))%32) {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v121 = v116 + int32(4)
	if base.B2i32(v121 < v116)|base.B2i32(base.Ui32(int32(1073741824)) <= base.Ui32(v121)) != 0 {
		goto L36
	} else {
		goto L45
	}
L45:
	;
	v126 = v110 - v107
	v127 = F_palloc(m, v121)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v130 = v127 + int32(4)
	if v107 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v131 = int32(1)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v133&v131 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v156 = v130
	goto L49
L49:
	;
	if v126 != 0 {
		goto L60
	} else {
		goto L61
	}
L50:
	;
	v136 = v131
	goto L52
L51:
	;
	v136 = int32(4)
	goto L52
L52:
	;
	v138 = v130
	v139 = v13 + v136
	v145 = v107
	goto L53
L53:
	;
	v149 = F_pg_mblen_unbounded(m, v139)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	v156 = v153
	goto L49
L55:
	;
	if v149 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	base.MemoryCopy(m, v138, v139, v149)
	goto L58
L57:
	;
	goto L58
L58:
	;
	v153 = v138 + v149
	v155 = v145 - int32(1)
	if v155 != 0 {
		v138 = v153
		v139 = v139 + v149
		v145 = v155
		goto L53
	} else {
		goto L59
	}
L59:
	;
	goto L54
L60:
	;
	v167 = int32(1)
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v169&v167 != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v199 = v156
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127))) = (v199 - v127) << (uint(int32(2)) % 32)
	return base.I64_extend_i32_u(v127)
L63:
	;
	v172 = v167
	goto L65
L64:
	;
	v172 = int32(4)
	goto L65
L65:
	;
	v173 = v21 + v172
	v174 = int32(0)
	if v174 < v88 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v177 = v88
	goto L68
L67:
	;
	v177 = v174
	goto L68
L68:
	;
	v178 = v173 + v177
	v179 = v156
	v180 = v173
	v187 = v126
	goto L69
L69:
	;
	v190 = F_pg_mblen_range(m, v180, v178)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L71
	}
L70:
	;
	v199 = v196
	goto L62
L71:
	;
	if v190 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	base.MemoryCopy(m, v179, v180, v190)
	goto L74
L73:
	;
	goto L74
L74:
	;
	v193 = v180 + v190
	if v193 == v178 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v195 = v173
	goto L77
L76:
	;
	v195 = v193
	goto L77
L77:
	;
	v196 = v179 + v190
	v198 = v187 - int32(1)
	if v198 != 0 {
		v179 = v196
		v180 = v195
		v187 = v198
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
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errmsg(m, int32(_a_F_rpad_0), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_rpad_1), int32(304), int32(_a_F_rpad_2))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
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
