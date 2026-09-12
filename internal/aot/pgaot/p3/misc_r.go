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
				v72 = int32(278979)
				v73 = int32(4104)
				F_errmsg(m, v72, int32(0))
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(517171), v73, int32(443517))
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
							v72 = int32(442941)
							v73 = int32(4114)
							F_errmsg(m, v72, int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(517171), v73, int32(443517))
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
								v72 = int32(442648)
								v73 = int32(4120)
								F_errmsg(m, v72, int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(517171), v73, int32(443517))
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
									v72 = int32(442577)
									v73 = int32(4128)
									F_errmsg(m, v72, int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(517171), v73, int32(443517))
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
										v72 = int32(442836)
										v73 = int32(4134)
										F_errmsg(m, v72, int32(0))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(517171), v73, int32(443517))
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
		v7 = *(*int32)(unsafe.Add(mBase, _consts[181]))
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
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
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
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
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
	F_ScanKeyInit(m, v14+int32(176), int32(4), int32(3), int32(184), int32(1259))
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
	v49 = F_systable_beginscan(m, v18, int32(2674), int32(1), int32(0), int32(3), v14+int32(176))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L151
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L148
	}
L8:
	;
	v51 = F_systable_getnext(m, v49)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v51 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v54 = base.B2i32(l1 != int32(24))
	v61 = v51
	goto L13
L11:
	;
	goto L12
L12:
	;
	F_systable_endscan(m, v49)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L146
	}
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+22)))
	v68 = v66 + v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+164)) = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+168)) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+172)) = v73
	if v69 <= int32(3255) {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	goto L12
L15:
	;
	v443 = F_systable_getnext(m, v49)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L144
	}
L16:
	;
	F_RememberConstraintForRebuilding(m, v71, l0)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L143
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L139
	}
L18:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v370 = int32(0)
	if v369 == v370 {
		goto L123
	} else {
		goto L124
	}
L19:
	;
	F_GetAttrDefaultColumnAddress(m, v14+int32(152), v71)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
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
	v115 = F_get_rel_relkind(m, v71)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L39
	}
L25:
	;
	switch v69 - int32(2604) {
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
	if v69 == int32(3256) {
		goto L20
	} else {
		goto L29
	}
L28:
	;
	switch v69 - int32(1255) {
	case 0:
		goto L23
	default:
		goto L17
	case 4:
		goto L24
	}
L29:
	;
	if v69 == int32(3381) {
		goto L18
	} else {
		goto L30
	}
L30:
	;
	if v69 != int32(6106) {
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
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errmsg(m, int32(378560), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v101 = F_getObjectDescription(m, v14+int32(164), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+148)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = v101
	F_errdetail(m, int32(747571), v14+int32(144))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(520068), int32(15245), int32(354853))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
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
	if v115&int32(-33) == int32(73) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+168))
	v123 = int32(0)
	if v121 == v123 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	goto L42
L42:
	;
	if v115 == int32(83) {
		goto L15
	} else {
		goto L75
	}
L43:
	;
	if v161 != 0 {
		goto L15
	} else {
		goto L56
	}
L44:
	;
	v161 = int32(0)
	goto L43
L45:
	;
	goto L46
L46:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	if v129 <= int32(0) {
		v154 = v123
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v161 = v154
	goto L43
L48:
	;
	v132 = int32(0)
	if v132 < v129 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v135 = v129
	goto L51
L50:
	;
	v135 = v132
	goto L51
L51:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	v138 = int32(0)
	goto L52
L52:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v136+v138<<(uint(int32(2))%32))))
	v147 = base.B2i32(v146 == v122)
	if v146 == v122 {
		v154 = v147
		goto L47
	} else {
		goto L54
	}
L53:
	;
	v154 = v147
	goto L47
L54:
	;
	v149 = v138 + int32(1)
	if v149 != v135 {
		v138 = v149
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v162 = F_get_index_constraint(m, v122)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	if v162 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	F_RememberConstraintForRebuilding(m, v162, l0)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v166 = int32(0)
	v170 = int32(1)
	v174 = F_pg_get_indexdef_worker(m, v122, v166, v166, v166, v166, v170, v170, v166, v166)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	goto L15
L62:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v177 = F_lappend_oid(m, v176, v122)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v177
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v181 = F_lappend(m, v180, v174)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v181
	v184 = F_get_index_isreplident(m, v122)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v184 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v186 != 0 {
		goto L7
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v190 = F_get_index_isclustered(m, v122)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	v187 = F_get_rel_name(m, v122)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v187
	goto L68
L71:
	;
	if v190 == int32(0) {
		goto L15
	} else {
		goto L72
	}
L72:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v194 != 0 {
		goto L6
	} else {
		goto L73
	}
L73:
	;
	v195 = F_get_rel_name(m, v122)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v195
	goto L15
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v207 = F_getObjectDescription(m, v14+int32(164), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v207
	F_errmsg_internal(m, int32(213311), v14+int32(48))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(520068), int32(15104), int32(354853))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
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
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errmsg(m, int32(383029), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v234 = F_getObjectDescription(m, v14+int32(164), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v234
	F_errdetail(m, int32(747571), v14-int32(-64))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(520068), int32(15132), int32(354853))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
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
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errmsg(m, int32(402364), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v262 = F_getObjectDescription(m, v14+int32(164), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v262
	F_errdetail(m, int32(747571), v14+int32(80))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(520068), int32(15147), int32(354853))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
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
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errmsg(m, int32(263781), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v290 = F_getObjectDescription(m, v14+int32(164), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+100)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v290
	F_errdetail(m, int32(747571), v14+int32(96))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(520068), int32(15167), int32(354853))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
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
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(263508), int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v318 = F_getObjectDescription(m, v14+int32(164), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+116)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v318
	F_errdetail(m, int32(747571), v14+int32(112))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(520068), int32(15186), int32(354853))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
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
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v14)+156))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	if v336 == v337 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
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
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v14)+160))
	if v339 != l3 {
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
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_errmsg(m, int32(288508), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v14)+156))
	v353 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+160)))
	v355 = F_get_attname(m, v352, v353, int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+132)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = l4
	F_errdetail(m, int32(699605), v14+int32(128))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(520068), int32(15219), int32(354853))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
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
	if v408 != 0 {
		goto L15
	} else {
		goto L135
	}
L123:
	;
	v408 = int32(0)
	goto L122
L124:
	;
	goto L125
L125:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	if v376 <= int32(0) {
		v401 = v370
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v408 = v401
	goto L122
L127:
	;
	v379 = int32(0)
	if v379 < v376 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v382 = v376
	goto L130
L129:
	;
	v382 = v379
	goto L130
L130:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	v385 = int32(0)
	goto L131
L131:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v383+v385<<(uint(int32(2))%32))))
	v394 = base.B2i32(v393 == v71)
	if v393 == v71 {
		v401 = v394
		goto L126
	} else {
		goto L133
	}
L132:
	;
	v401 = v394
	goto L126
L133:
	;
	v396 = v385 + int32(1)
	if v396 != v382 {
		v385 = v396
		goto L131
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	v409 = int32(0)
	v411 = F_pg_get_statisticsobj_worker(m, v71, v409, v409)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v414 = F_lappend_oid(m, v413, v71)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v414
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v418 = F_lappend(m, v417, v411)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v418
	goto L15
L139:
	;
	v428 = F_getObjectDescription(m, v14+int32(164), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v428
	F_errmsg_internal(m, int32(213311), v14)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(520068), int32(15255), int32(354853))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
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
	if v443 != 0 {
		v61 = v443
		goto L13
	} else {
		goto L145
	}
L145:
	;
	goto L14
L146:
	;
	F_sequence_close(m, v18, int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
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
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v468
	F_errmsg_internal(m, int32(10574), v14+int32(32))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(520068), int32(15275), int32(354816))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
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
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v484
	F_errmsg_internal(m, int32(167964), v14+int32(16))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(520068), int32(15290), int32(354887))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
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
								v91 = m.ExcPending
								if v91 != 0 {
									return
								} else {
									F_errcode(m, int32(290948))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
										F_errmsg(m, int32(125039), v12+int32(16))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											F_errfinish(m, int32(525105), int32(805), int32(328956))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
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
										v91 = m.ExcPending
										if v91 != 0 {
											return
										} else {
											F_errcode(m, int32(290948))
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
												F_errmsg(m, int32(125039), v12+int32(16))
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return
												} else {
													F_errfinish(m, int32(525105), int32(805), int32(328956))
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
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
											v50 = *(*int32)(unsafe.Add(mBase, _consts[444]))
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
														F_sequence_close(m, v16, int32(3))
														mBase = m.M
														v61 = m.ExcPending
														if v61 != 0 {
															return
														} else {
															if v25 == int32(0) {
																m.G0 = v12 + int32(32)
																return
															} else {
																if v25 == v29 {
																	m.G0 = v12 + int32(32)
																	return
																} else {
																	v65 = F_makeArrayTypeName(m, l1, l2)
																	mBase = m.M
																	v66 = m.ExcPending
																	if v66 != 0 {
																		return
																	} else {
																		F_RenameTypeInternal(m, v25, v65, l2)
																		mBase = m.M
																		v68 = m.ExcPending
																		if v68 != 0 {
																			return
																		} else {
																			F_pfree(m, v65)
																			mBase = m.M
																			v70 = m.ExcPending
																			if v70 != 0 {
																				return
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
											} else {
												F_pfree(m, v20)
												mBase = m.M
												v58 = m.ExcPending
												if v58 != 0 {
													return
												} else {
													F_sequence_close(m, v16, int32(3))
													mBase = m.M
													v61 = m.ExcPending
													if v61 != 0 {
														return
													} else {
														if v25 == int32(0) {
															m.G0 = v12 + int32(32)
															return
														} else {
															if v25 == v29 {
																m.G0 = v12 + int32(32)
																return
															} else {
																v65 = F_makeArrayTypeName(m, l1, l2)
																mBase = m.M
																v66 = m.ExcPending
																if v66 != 0 {
																	return
																} else {
																	F_RenameTypeInternal(m, v25, v65, l2)
																	mBase = m.M
																	v68 = m.ExcPending
																	if v68 != 0 {
																		return
																	} else {
																		F_pfree(m, v65)
																		mBase = m.M
																		v70 = m.ExcPending
																		if v70 != 0 {
																			return
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
							v50 = *(*int32)(unsafe.Add(mBase, _consts[444]))
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
										F_sequence_close(m, v16, int32(3))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return
										} else {
											if v25 == int32(0) {
												m.G0 = v12 + int32(32)
												return
											} else {
												if v25 == v29 {
													m.G0 = v12 + int32(32)
													return
												} else {
													v65 = F_makeArrayTypeName(m, l1, l2)
													mBase = m.M
													v66 = m.ExcPending
													if v66 != 0 {
														return
													} else {
														F_RenameTypeInternal(m, v25, v65, l2)
														mBase = m.M
														v68 = m.ExcPending
														if v68 != 0 {
															return
														} else {
															F_pfree(m, v65)
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return
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
							} else {
								F_pfree(m, v20)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									F_sequence_close(m, v16, int32(3))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										if v25 == int32(0) {
											m.G0 = v12 + int32(32)
											return
										} else {
											if v25 == v29 {
												m.G0 = v12 + int32(32)
												return
											} else {
												v65 = F_makeArrayTypeName(m, l1, l2)
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return
												} else {
													F_RenameTypeInternal(m, v25, v65, l2)
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return
													} else {
														F_pfree(m, v65)
														mBase = m.M
														v70 = m.ExcPending
														if v70 != 0 {
															return
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
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
					F_errmsg_internal(m, int32(55308), v12)
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return
					} else {
						F_errfinish(m, int32(525105), int32(777), int32(328956))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
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
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v171 int32
	_ = v171
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v216 int32
	_ = v216
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v249 int32
	_ = v249
	var v261 int32
	_ = v261
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v294 int32
	_ = v294
	var v306 int32
	_ = v306
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v339 int32
	_ = v339
	var v351 int32
	_ = v351
	var v362 int32
	_ = v362
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v384 int32
	_ = v384
	var v396 int32
	_ = v396
	var v407 int32
	_ = v407
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v429 int32
	_ = v429
	var v441 int32
	_ = v441
	var v452 int32
	_ = v452
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v518 int64
	_ = v518
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v615 int64
	_ = v615
	var v636 int32
	_ = v636
	var v647 int32
	_ = v647
	var v654 int32
	_ = v654
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v698 int32
	_ = v698
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v716 int32
	_ = v716
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v738 int32
	_ = v738
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v767 int32
	_ = v767
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v785 int32
	_ = v785
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v824 int32
	_ = v824
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v848 int32
	_ = v848
	var v856 int32
	_ = v856
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v958 int32
	_ = v958
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v983 int32
	_ = v983
	var v991 int32
	_ = v991
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1014 int32
	_ = v1014
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1032 int32
	_ = v1032
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1073 int32
	_ = v1073
	var v1082 int32
	_ = v1082
	var v1088 int32
	_ = v1088
	var v1089 int64
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	v19 = int32(-1)
	v20 = v3
	v21 = v3
	v22 = v3
	v23 = v3
	v24 = v15
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
	if v19 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v1088 = int32(m.ExcTag)
	v1089 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1088 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L7:
	;
	v32 = int32(16)
	v33 = v24 - v32
	m.G0 = v33
	v36 = v33 - int32(160)
	m.G0 = v36
	v39 = v36 - v32
	m.G0 = v39
	*(*int32)(unsafe.Add(mBase, _consts[200])) = int32(7)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v33
	goto L11
L8:
	;
	v69 = v20
	v70 = v21
	v71 = v22
	v72 = v24
	v73 = v23
	goto L9
L9:
	;
	if v73 != 0 {
		goto L20
	} else {
		goto L21
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v33
	F_InitProcess(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		v1082 = v39
		goto L6
	} else {
		goto L14
	}
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[200]))
	v52 = F_GetBackendTypeDesc(m, v51)
	mBase = m.M
	goto L13
L13:
	;
	goto L10
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v33
	F_BaseInit(m)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		v1082 = v39
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v15 - int32(-64)
	goto L19
L17:
	;
	v69 = v36
	v70 = v39
	v71 = v33
	v72 = v39
	v73 = int32(0)
	goto L9
L19:
	;
	goto L17
L20:
	;
	v74 = int32(4556748)
	v76 = *(*int32)(unsafe.Add(mBase, _consts[117]))
	*(*int32)(unsafe.Add(mBase, _consts[117])) = v76 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[52])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_EmitErrorReport(m)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, _consts[53])) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v100 = int32(914)
	v102 = m.G0
	v104 = v102 - int32(144)
	m.G0 = v104
	switch int32(916) {
	case 0, 2:
		v114 = v100
		goto L26
	default:
		goto L27
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_proc_exit(m, int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L24
	}
L24:
	;
	goto L3
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v145 = int32(915)
	v147 = m.G0
	v149 = v147 - int32(144)
	m.G0 = v149
	switch int32(917) {
	case 0, 2:
		v159 = v145
		goto L39
	default:
		goto L40
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v114
	F_sigemptyset(m, v104+int32(8))
	mBase = m.M
	goto L29
L27:
	;
	*(*int32)(unsafe.Add(mBase, _consts[630])) = v100
	v114 = int32(4733)
	goto L26
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+136)) = int32(268435456)
	v126 = v104 + int32(4)
	goto L33
L31:
	;
	m.G0 = v104 + int32(144)
	goto L25
L33:
	;
	goto L34
L34:
	;
	if v126 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v137 = F___memcpy(m, int32(4736156), v126, int32(140))
	mBase = m.M
	goto L37
L36:
	;
	goto L37
L37:
	;
	goto L31
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v190 = int32(295)
	v192 = m.G0
	v194 = v192 - int32(144)
	m.G0 = v194
	switch int32(297) {
	case 0, 2:
		v204 = v190
		goto L52
	default:
		goto L53
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v159
	F_sigemptyset(m, v149+int32(8))
	mBase = m.M
	goto L42
L40:
	;
	*(*int32)(unsafe.Add(mBase, _consts[624])) = v145
	v159 = int32(4733)
	goto L39
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149)+136)) = int32(268435456)
	v171 = v149 + int32(4)
	goto L46
L44:
	;
	m.G0 = v149 + int32(144)
	goto L38
L46:
	;
	goto L47
L47:
	;
	if v171 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v182 = F___memcpy(m, int32(4736296), v171, int32(140))
	mBase = m.M
	goto L50
L49:
	;
	goto L50
L50:
	;
	goto L44
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v235 = int32(919)
	v237 = m.G0
	v239 = v237 - int32(144)
	m.G0 = v239
	switch int32(921) {
	case 0, 2:
		v249 = v235
		goto L65
	default:
		goto L66
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v194)+4)) = v204
	F_sigemptyset(m, v194+int32(8))
	mBase = m.M
	goto L55
L53:
	;
	*(*int32)(unsafe.Add(mBase, _consts[627])) = v190
	v204 = int32(4733)
	goto L52
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v194)+136)) = int32(268435456)
	v216 = v194 + int32(4)
	goto L59
L57:
	;
	m.G0 = v194 + int32(144)
	goto L51
L59:
	;
	goto L60
L60:
	;
	if v216 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v227 = F___memcpy(m, int32(4738116), v216, int32(140))
	mBase = m.M
	goto L63
L62:
	;
	goto L63
L63:
	;
	goto L57
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v280 = int32(917)
	v282 = m.G0
	v284 = v282 - int32(144)
	m.G0 = v284
	switch int32(919) {
	case 0, 2:
		v294 = v280
		goto L78
	default:
		goto L79
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239)+4)) = v249
	F_sigemptyset(m, v239+int32(8))
	mBase = m.M
	goto L68
L66:
	;
	*(*int32)(unsafe.Add(mBase, _consts[626])) = v235
	v249 = int32(4733)
	goto L65
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239)+136)) = int32(268435456)
	v261 = v239 + int32(4)
	goto L72
L70:
	;
	m.G0 = v239 + int32(144)
	goto L64
L72:
	;
	goto L73
L73:
	;
	if v261 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v272 = F___memcpy(m, int32(4737136), v261, int32(140))
	mBase = m.M
	goto L76
L75:
	;
	goto L76
L76:
	;
	goto L70
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v325 = int32(-2)
	v327 = m.G0
	v329 = v327 - int32(144)
	m.G0 = v329
	switch int32(0) {
	case 0, 2:
		v339 = v325
		goto L91
	default:
		goto L92
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v284)+4)) = v294
	F_sigemptyset(m, v284+int32(8))
	mBase = m.M
	goto L81
L79:
	;
	*(*int32)(unsafe.Add(mBase, _consts[625])) = v280
	v294 = int32(4733)
	goto L78
L81:
	;
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v284)+136)) = int32(268435456)
	v306 = v284 + int32(4)
	goto L85
L83:
	;
	m.G0 = v284 + int32(144)
	goto L77
L85:
	;
	goto L86
L86:
	;
	if v306 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v317 = F___memcpy(m, int32(4737416), v306, int32(140))
	mBase = m.M
	goto L89
L88:
	;
	goto L89
L89:
	;
	goto L83
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v370 = int32(-2)
	v372 = m.G0
	v374 = v372 - int32(144)
	m.G0 = v374
	switch int32(0) {
	case 0, 2:
		v384 = v370
		goto L104
	default:
		goto L105
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+4)) = v339
	F_sigemptyset(m, v329+int32(8))
	mBase = m.M
	goto L94
L92:
	;
	*(*int32)(unsafe.Add(mBase, _consts[640])) = v325
	v339 = int32(4733)
	goto L91
L94:
	;
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+136)) = int32(268435456)
	v351 = v329 + int32(4)
	goto L98
L96:
	;
	m.G0 = v329 + int32(144)
	goto L90
L98:
	;
	goto L99
L99:
	;
	if v351 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v362 = F___memcpy(m, int32(4737696), v351, int32(140))
	mBase = m.M
	goto L102
L101:
	;
	goto L102
L102:
	;
	goto L96
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v415 = int32(0)
	v417 = m.G0
	v419 = v417 - int32(144)
	m.G0 = v419
	switch int32(2) {
	case 0, 2:
		v429 = v415
		goto L117
	default:
		goto L118
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v374)+4)) = v384
	F_sigemptyset(m, v374+int32(8))
	mBase = m.M
	goto L107
L105:
	;
	*(*int32)(unsafe.Add(mBase, _consts[639])) = v370
	v384 = int32(4733)
	goto L104
L107:
	;
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v374)+136)) = int32(268435456)
	v396 = v374 + int32(4)
	goto L111
L109:
	;
	m.G0 = v374 + int32(144)
	goto L103
L111:
	;
	goto L112
L112:
	;
	if v396 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v407 = F___memcpy(m, int32(4737836), v396, int32(140))
	mBase = m.M
	goto L115
L114:
	;
	goto L115
L115:
	;
	goto L109
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v460 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	F_check_and_set_sync_info(m, v460)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L129
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v419)+4)) = v429
	F_sigemptyset(m, v419+int32(8))
	mBase = m.M
	goto L119
L118:
	;
	*(*int32)(unsafe.Add(mBase, _consts[641])) = v415
	v429 = int32(4733)
	goto L117
L119:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v419)+136)) = int32(268435457)
	v441 = v419 + int32(4)
	goto L124
L122:
	;
	m.G0 = v419 + int32(144)
	goto L116
L124:
	;
	goto L125
L125:
	;
	if v441 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v452 = F___memcpy(m, int32(4738396), v441, int32(140))
	mBase = m.M
	goto L128
L127:
	;
	goto L128
L128:
	;
	goto L122
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v468 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L130
	}
L130:
	;
	if v468 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_errmsg(m, int32(467156), int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_before_shmem_exit(m, int32(1020), int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L136
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_errfinish(m, int32(526536), int32(1425), int32(292791))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v495 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[628])) = v495
	*(*int32)(unsafe.Add(mBase, _consts[629])) = v495
	v505 = v495
	goto L138
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_load_file(m, int32(226665), int32(0))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L143
	}
L138:
	;
	v507 = int32(40)
	v508 = v505 * v507
	v511 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v508)+uint32(_consts[631]))) = uint8(v511)
	*(*int32)(unsafe.Add(mBase, uint32(v508)+uint32(_consts[632]))) = v505
	v518 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v508)+uint32(_consts[633]))) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v508)+uint32(_consts[634]))) = v511
	*(*int64)(unsafe.Add(mBase, uint32(v508)+uint32(_consts[635]))) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v508)+uint32(_consts[636]))) = v511
	*(*uint8)(unsafe.Add(mBase, uint32(v508)+uint32(_consts[637]))) = uint8(v511)
	v537 = v505 | int32(1)
	v539 = v537 * v507
	*(*uint8)(unsafe.Add(mBase, uint32(v539)+uint32(_consts[631]))) = uint8(v511)
	*(*int32)(unsafe.Add(mBase, uint32(v539)+uint32(_consts[632]))) = v537
	*(*int64)(unsafe.Add(mBase, uint32(v539)+uint32(_consts[633]))) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v539)+uint32(_consts[634]))) = v511
	*(*int64)(unsafe.Add(mBase, uint32(v539)+uint32(_consts[635]))) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v539)+uint32(_consts[636]))) = v511
	*(*uint8)(unsafe.Add(mBase, uint32(v539)+uint32(_consts[637]))) = uint8(v511)
	v568 = v505 | int32(2)
	v570 = v568 * v507
	*(*uint8)(unsafe.Add(mBase, uint32(v570)+uint32(_consts[631]))) = uint8(v511)
	*(*int32)(unsafe.Add(mBase, uint32(v570)+uint32(_consts[632]))) = v568
	*(*int64)(unsafe.Add(mBase, uint32(v570)+uint32(_consts[633]))) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v570)+uint32(_consts[634]))) = v511
	*(*int64)(unsafe.Add(mBase, uint32(v570)+uint32(_consts[635]))) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v570)+uint32(_consts[636]))) = v511
	*(*uint8)(unsafe.Add(mBase, uint32(v570)+uint32(_consts[637]))) = uint8(v511)
	if base.B2i32(v505 == int32(20)) == v511 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v636 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[638])) = uint8(v636)
	F_pqsignal_be(m, int32(14), int32(1788))
	mBase = m.M
	goto L137
L140:
	;
	v603 = v505 | int32(3)
	v605 = v603 * int32(40)
	v608 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v605)+uint32(_consts[631]))) = uint8(v608)
	*(*int32)(unsafe.Add(mBase, uint32(v605)+uint32(_consts[632]))) = v603
	v615 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v605)+uint32(_consts[633]))) = v615
	*(*int32)(unsafe.Add(mBase, uint32(v605)+uint32(_consts[634]))) = v608
	*(*int64)(unsafe.Add(mBase, uint32(v605)+uint32(_consts[635]))) = v615
	*(*int32)(unsafe.Add(mBase, uint32(v605)+uint32(_consts[636]))) = v608
	*(*uint8)(unsafe.Add(mBase, uint32(v605)+uint32(_consts[637]))) = uint8(v608)
	v505 = v505 + int32(4)
	goto L138
L141:
	;
	goto L142
L142:
	;
	goto L139
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_sigprocmask(m, int32(4469288), int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_SetConfigOption(m, int32(338457), int32(794587), int32(5), int32(10))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L145
	}
L145:
	;
	v665 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v665)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v671 = *(*int32)(unsafe.Add(mBase, _consts[232]))
	v672 = m.T0[v666].(func(*base.Module, int32) int32)(m, v671)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L146
	}
L146:
	;
	if v672 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v710 = int32(0)
	F_InitPostgres(m, v672, v710, v710, v710, v710, v710)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L154
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_errcode(m, int32(50856066))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(254158)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(397841)
	F_errmsg(m, int32(749116), v15)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_errfinish(m, int32(526536), int32(1052), int32(254175))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L153
	}
L153:
	;
	goto L3
L154:
	;
	*(*int32)(unsafe.Add(mBase, _consts[87])) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_initStringInfo(m, v70)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L155
	}
L155:
	;
	v726 = *(*int32)(unsafe.Add(mBase, _consts[682]))
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726))))
	if v727 != 0 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v747 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v747)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v753 = *(*int32)(unsafe.Add(mBase, _consts[232]))
	v754 = int32(0)
	v757 = m.T0[v748].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v753, v754, v754, v754, v745, v71)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L162
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = int32(232949)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v726
	F_appendStringInfo(m, v70, int32(187122), v15+int32(48))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_appendStringInfoString(m, v70, int32(232949))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		v1082 = v72
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
	if v757 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_pfree(m, v794)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L170
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_errcode(m, int32(100663808))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L167
	}
L167:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v775
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v774
	F_errmsg(m, int32(212122), v15+int32(16))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_errfinish(m, int32(526536), int32(1488), int32(292791))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L169
	}
L169:
	;
	goto L3
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_before_shmem_exit(m, int32(1021), v757)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_validate_remote_info(m, v757)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L172
	}
L172:
	;
	goto L173
L173:
	;
	v824 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v824 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_ProcessInterrupts(m)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v831 = *(*int32)(unsafe.Add(mBase, _consts[312]))
	v832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v831)+4)))
	if v832 == int32(1) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	goto L177
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v840 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v864 = *(*int32)(unsafe.Add(mBase, _consts[348]))
	if v864 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L182:
	;
	if v840 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_errmsg(m, int32(473791), int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_proc_exit(m, int32(0))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L188
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_errfinish(m, int32(526536), int32(1184), int32(126296))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		v1082 = v72
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v1042 = F_synchronize_slots(m, v757)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L233
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v871 = *(*int32)(unsafe.Add(mBase, _consts[232]))
	v872 = F_pstrdup(m, v871)
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v878 = *(*int32)(unsafe.Add(mBase, _consts[398]))
	v879 = F_pstrdup(m, v878)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	*(*int32)(unsafe.Add(mBase, _consts[348])) = int32(0)
	v888 = int32(*(*uint8)(unsafe.Add(mBase, _consts[683])))
	v890 = int32(*(*uint8)(unsafe.Add(mBase, _consts[684])))
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v898 = *(*int32)(unsafe.Add(mBase, _consts[232]))
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898))))
	v902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v872))))
	if v902 == int32(0) {
		v921 = v901
		v922 = v902
		goto L195
	} else {
		goto L196
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v928 = *(*int32)(unsafe.Add(mBase, _consts[398]))
	v931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v928))))
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v879))))
	if v932 == int32(0) {
		v951 = v931
		v952 = v932
		goto L203
	} else {
		goto L204
	}
L195:
	;
	goto L194
L196:
	;
	if v901 != v902 {
		v921 = v901
		v922 = v902
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v906 = v872
	v907 = v898
	goto L198
L198:
	;
	v910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v907)+1)))
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v906)+1)))
	if v911 == int32(0) {
		v921 = v910
		v922 = v911
		goto L195
	} else {
		goto L200
	}
L199:
	;
	v921 = v910
	v922 = v911
	goto L195
L200:
	;
	v914 = int32(1)
	if v910 == v911 {
		v906 = v906 + v914
		v907 = v907 + v914
		goto L198
	} else {
		goto L201
	}
L201:
	;
	goto L199
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_pfree(m, v872)
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L210
	}
L203:
	;
	goto L202
L204:
	;
	if v931 != v932 {
		v951 = v931
		v952 = v932
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v936 = v879
	v937 = v928
	goto L206
L206:
	;
	v940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937)+1)))
	v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v936)+1)))
	if v941 == int32(0) {
		v951 = v940
		v952 = v941
		goto L203
	} else {
		goto L208
	}
L207:
	;
	v951 = v940
	v952 = v941
	goto L203
L208:
	;
	v944 = int32(1)
	if v940 == v941 {
		v936 = v936 + v944
		v937 = v937 + v944
		goto L206
	} else {
		goto L209
	}
L209:
	;
	goto L207
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_pfree(m, v879)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L211
	}
L211:
	;
	v965 = int32(*(*uint8)(unsafe.Add(mBase, _consts[684])))
	if v965 != v890 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v972 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	if v922-v921 != 0 {
		goto L222
	} else {
		goto L223
	}
L215:
	;
	if v972 != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = int32(126658)
	F_errmsg(m, int32(477939), v15+int32(32))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_proc_exit(m, int32(0))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L221
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_errfinish(m, int32(526536), int32(1151), int32(355505))
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L220
	}
L220:
	;
	goto L218
L221:
	;
	goto L3
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v1006 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L226
	}
L223:
	;
	if v952-v951 != 0 {
		goto L222
	} else {
		goto L224
	}
L224:
	;
	v999 = int32(*(*uint8)(unsafe.Add(mBase, _consts[683])))
	if v888 == v999 {
		goto L189
	} else {
		goto L225
	}
L225:
	;
	goto L222
L226:
	;
	if v1006 != 0 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_errmsg(m, int32(423317), int32(0))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, _consts[312]))
	*(*int64)(unsafe.Add(mBase, uint32(v1024)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_proc_exit(m, int32(0))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L232
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	F_errfinish(m, int32(526536), int32(1160), int32(355505))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	goto L3
L233:
	;
	v1044 = int32(4164308)
	v1046 = int32(30000)
	v1048 = *(*int32)(unsafe.Add(mBase, _consts[685]))
	v1050 = v1048 << (uint(int32(1)) % 32)
	if v1046 <= v1050 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v1053 = v1046
	goto L236
L235:
	;
	v1053 = v1050
	goto L236
L236:
	;
	if v1042 != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v1054 = int32(200)
	goto L239
L238:
	;
	v1054 = v1053
	goto L239
L239:
	;
	*(*int32)(unsafe.Add(mBase, _consts[685])) = v1054
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v1060 = *(*int32)(unsafe.Add(mBase, _consts[313]))
	v1063 = F_WaitLatch(m, v1060, int32(41), v1054, int32(83886091))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		v1082 = v72
		goto L6
	} else {
		goto L240
	}
L240:
	;
	if v1063&int32(1) == int32(0) {
		goto L173
	} else {
		goto L241
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v71
	v1073 = *(*int32)(unsafe.Add(mBase, _consts[313]))
	*(*int32)(unsafe.Add(mBase, uint32(v1073))) = int32(0)
	goto L242
L242:
	;
	goto L173
L243:
	;
	v1093 = int32(v1089)
	m.G0 = v1082
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+4))
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1093)))
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1096)))
	if v15-int32(-64) == v1100 {
		goto L246
	} else {
		goto L247
	}
L244:
	;
	m.ExcPending = 1
	goto L252
L245:
	;
	if v1103 != 0 {
		goto L249
	} else {
		goto L250
	}
L246:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1096)+4))
	v1103 = v1102
	goto L248
L247:
	;
	v1103 = int32(0)
	goto L248
L248:
	;
	goto L245
L249:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	v19 = v1103
	v20 = v1105
	v21 = v1106
	v22 = v1104
	v23 = v1095
	v24 = v1082
	goto L1
L250:
	;
	goto L251
L251:
	;
	F___wasm_longjmp(m, v1096, v1095)
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	return
L253:
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
			v7 = int32(4457908)
			v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[360])))
			v10 = v9 | int32(2)
			*(*uint8)(unsafe.Add(mBase, _consts[360])) = uint8(v10)
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int64
	_ = v161
	var v163 int64
	_ = v163
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
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
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
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
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
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
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
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
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v624 int32
	_ = v624
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v720 int32
	_ = v720
	var v727 int32
	_ = v727
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v742 int32
	_ = v742
	var v749 int32
	_ = v749
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v793 int32
	_ = v793
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v822 int32
	_ = v822
	var v828 int32
	_ = v828
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v886 int32
	_ = v886
	var v891 int32
	_ = v891
	var v897 int32
	_ = v897
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v953 int32
	_ = v953
	var v960 int32
	_ = v960
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1004 int32
	_ = v1004
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
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
	var v1091 int32
	_ = v1091
	var v1096 int32
	_ = v1096
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1116 int32
	_ = v1116
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1141 int32
	_ = v1141
	var v1146 int32
	_ = v1146
	v9 = m.G0
	v11 = v9 - int32(7392)
	m.G0 = v11
	v13 = F_AllocateDir(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L4
	} else {
		goto L242
	}
L2:
	;
	m.G0 = v11 + int32(7392)
	return
L3:
	;
	v36 = F_ReadDir(m, v13, l0)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L13
	}
L4:
	;
	return
L5:
	;
	if v13 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v16 != int32(44) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v21 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	if v21 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg(m, int32(311552), v11)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(518316), int32(127), int32(225248))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	goto L2
L13:
	;
	if v36 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v39 = l1 & int32(1)
	v41 = l1 & int32(2)
	v45 = v36
	goto L17
L15:
	;
	goto L16
L16:
	;
	F_FreeDir(m, v13)
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L4
	} else {
		goto L241
	}
L17:
	;
	v51 = v45 + int32(19)
	v52 = int32(580318)
	v56 = m.G0
	v58 = v56 - int32(32)
	v59 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v58)+24)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v58)+16)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v58)+8)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v58))) = v59
	v67 = int32(*(*uint8)(unsafe.Add(mBase, _consts[266])))
	if v67 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L16
L19:
	;
	v1105 = F_ReadDir(m, v13, l0)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L4
	} else {
		goto L239
	}
L20:
	;
	v136 = F_strlen(m, v51)
	mBase = m.M
	if v135 != v136 {
		goto L19
	} else {
		goto L41
	}
L21:
	;
	v135 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, _consts[267])))
	if v71 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v75 = v51
	goto L27
L25:
	;
	goto L26
L26:
	;
	v85 = v52
	v86 = v67
	goto L30
L27:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v81 == v67 {
		v75 = v75 + int32(1)
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v135 = v75 - v51
	goto L20
L29:
	;
	goto L28
L30:
	;
	v93 = v58 + int32(base.Ui32(v86)>>(uint(int32(3))%32))&int32(28)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v95 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v94 | v95<<(uint(v86)%32)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	if v99 != 0 {
		v85 = v85 + v95
		v86 = v99
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v102 == int32(0) {
		v127 = v51
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L31
L33:
	;
	v135 = v127 - v51
	goto L20
L34:
	;
	v106 = v51
	v107 = v102
	goto L35
L35:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v58+int32(base.Ui32(v107)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v115)>>(uint(v107)%32))&int32(1) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v127 = v123
	goto L33
L37:
	;
	v127 = v106
	goto L33
L38:
	;
	goto L39
L39:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	v123 = v106 + int32(1)
	if v121 != 0 {
		v106 = v123
		v107 = v121
		goto L35
	} else {
		goto L40
	}
L40:
	;
	goto L36
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = l0
	v146 = F_pg_snprintf(m, v11+int32(208), int32(2048), int32(188238), v11+int32(192))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	if v41 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v705 = F_AllocateDir(m, v11+int32(208))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L4
	} else {
		goto L153
	}
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+3296)) = int64(17179869188)
	v275 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+3320)) = v275
	v282 = F_hash_create(m, int32(186177), int32(32), v11+int32(3280), int32(1064))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L4
	} else {
		goto L69
	}
L45:
	;
	v155 = m.G0
	v157 = v155 - int32(16)
	m.G0 = v157
	v160 = *(*int32)(unsafe.Add(mBase, _consts[304]))
	if v160 != 0 {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	goto L47
L47:
	;
	if v39 == int32(0) {
		goto L19
	} else {
		goto L59
	}
L48:
	;
	if v39 != 0 {
		goto L44
	} else {
		goto L58
	}
L49:
	;
	if base.B2i32(v160 != int32(0)) == int32(0) {
		goto L48
	} else {
		goto L53
	}
L50:
	;
	v161 = F_GetCurrentTimestamp(m)
	mBase = m.M
	v163 = *(*int64)(unsafe.Add(mBase, _consts[305]))
	F_TimestampDifference(m, v163, v161, v157+int32(12), v157+int32(8))
	mBase = m.M
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[825]))) = v169
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(3280)))) = v171
	*(*int32)(unsafe.Add(mBase, _consts[304])) = int32(0)
	goto L52
L51:
	;
	goto L52
L52:
	;
	m.G0 = v157 + int32(16)
	goto L49
L53:
	;
	v186 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	if v186 == int32(0) {
		goto L48
	} else {
		goto L55
	}
L55:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[825])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+176)) = v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v11)+3280))
	v194 = base.I32_div_s(v192, int32(10000))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+180)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v11)+184)) = v11 + int32(208)
	F_errmsg(m, int32(213593), v11+int32(176))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(518316), int32(146), int32(225248))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	goto L48
L58:
	;
	goto L43
L59:
	;
	v218 = m.G0
	v220 = v218 - int32(16)
	m.G0 = v220
	v223 = *(*int32)(unsafe.Add(mBase, _consts[304]))
	if v223 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if base.B2i32(v223 != int32(0)) == int32(0) {
		goto L44
	} else {
		goto L64
	}
L61:
	;
	v224 = F_GetCurrentTimestamp(m)
	mBase = m.M
	v226 = *(*int64)(unsafe.Add(mBase, _consts[305]))
	F_TimestampDifference(m, v226, v224, v220+int32(12), v220+int32(8))
	mBase = m.M
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[825]))) = v232
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(3280)))) = v234
	*(*int32)(unsafe.Add(mBase, _consts[304])) = int32(0)
	goto L63
L62:
	;
	goto L63
L63:
	;
	m.G0 = v220 + int32(16)
	goto L60
L64:
	;
	v249 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	if v249 == int32(0) {
		goto L44
	} else {
		goto L66
	}
L66:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[825])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v253
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v11)+3280))
	v257 = base.I32_div_s(v255, int32(10000))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v11)+168)) = v11 + int32(208)
	F_errmsg(m, int32(213673), v11+int32(160))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(518316), int32(149), int32(225248))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	goto L44
L69:
	;
	v286 = F_AllocateDir(m, v11+int32(208))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v290 = F_ReadDir(m, v286, v11+int32(208))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	if v290 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v295 = v290
	goto L75
L73:
	;
	goto L74
L74:
	;
	F_FreeDir(m, v286)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L4
	} else {
		goto L102
	}
L75:
	;
	v301 = v295 + int32(19)
	v305 = v11 + int32(2256)
	v308 = int32(0)
	v313 = m.G0
	v315 = v313 - int32(16)
	m.G0 = v315
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[826]))) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v305))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[827]))) = v308
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	if base.Ui32((v323-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v403 = v308
		goto L79
	} else {
		goto L80
	}
L76:
	;
	goto L74
L77:
	;
	v420 = F_ReadDir(m, v286, v11+int32(208))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L4
	} else {
		goto L100
	}
L78:
	;
	if v403 == int32(0) {
		goto L77
	} else {
		goto L97
	}
L79:
	;
	m.G0 = v315 + int32(16)
	goto L78
L80:
	;
	v330 = int32(4735140)
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	v336 = F_strtoul(m, v301, v315+int32(8), int32(10))
	mBase = m.M
	v338 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v338 != 0 {
		v403 = v308
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v315)+8))
	if v301 == v339 {
		v403 = v308
		goto L79
	} else {
		goto L82
	}
L82:
	;
	if v336 == int32(0) {
		v403 = v308
		goto L79
	} else {
		goto L83
	}
L83:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339))))
	if v343 != int32(95) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	if v359&int32(255) == int32(46) {
		goto L89
	} else {
		goto L90
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v315)+12)) = int32(0)
	v359 = v343
	v360 = v339
	goto L84
L86:
	;
	goto L87
L87:
	;
	v352 = F_forkname_chars(m, v339+int32(1), v315+int32(12))
	mBase = m.M
	if v352 <= int32(0) {
		v403 = v308
		goto L79
	} else {
		goto L88
	}
L88:
	;
	v357 = v352 + v339 + int32(1)
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357))))
	v359 = v358
	v360 = v357
	goto L84
L89:
	;
	v366 = v360 + int32(1)
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366))))
	if base.Ui32((v367-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v403 = v308
		goto L79
	} else {
		goto L92
	}
L90:
	;
	v390 = v308
	v391 = v359
	goto L91
L91:
	;
	if v391&int32(255) != 0 {
		v403 = v308
		goto L79
	} else {
		goto L96
	}
L92:
	;
	v374 = int32(4735140)
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	v380 = F_strtoul(m, v366, v315+int32(8), int32(10))
	mBase = m.M
	v382 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v382 != 0 {
		v403 = v308
		goto L79
	} else {
		goto L93
	}
L93:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v315)+8))
	if v366 == v383 {
		v403 = v308
		goto L79
	} else {
		goto L94
	}
L94:
	;
	if v380 == int32(0) {
		v403 = v308
		goto L79
	} else {
		goto L95
	}
L95:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383))))
	v390 = v380
	v391 = v387
	goto L91
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[826]))) = v336
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v315)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v305))) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[827]))) = v390
	v403 = int32(1)
	goto L79
L97:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v11)+2256))
	if v409 != int32(3) {
		goto L77
	} else {
		goto L98
	}
L98:
	;
	v416 = F_hash_search(m, v282, v11+int32(5336), int32(1), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	goto L77
L100:
	;
	if v420 != 0 {
		v295 = v420
		goto L75
	} else {
		goto L101
	}
L101:
	;
	goto L76
L102:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v433)+4))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v433)+412))
	if v435 != 0 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v500 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L104:
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
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v433-int32(-64))))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v433)+52))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v433)+40))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v433)+28))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v433)+16))
	v500 = v436 + (v437 + (v438 + (v439 + (v440 + (v441 + (v442 + (v443 + (v444 + (v445 + (v446 + (v447 + (v448 + (v449 + (v450 + (v451 + (v452 + (v453 + (v454 + (v455 + (v456 + (v457 + (v458 + (v459 + (v460 + (v461 + (v464 + (v465 + (v466 + (v467 + (v468 + v434))))))))))))))))))))))))))))))
	goto L106
L105:
	;
	v500 = v434
	goto L106
L106:
	;
	goto L103
L107:
	;
	F_hash_destroy(m, v282)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L4
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v507 = F_AllocateDir(m, v11+int32(208))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L4
	} else {
		goto L111
	}
L110:
	;
	goto L19
L111:
	;
	v511 = F_ReadDir(m, v507, v11+int32(208))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	if v511 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v516 = v511
	goto L116
L114:
	;
	goto L115
L115:
	;
	F_FreeDir(m, v507)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L4
	} else {
		goto L150
	}
L116:
	;
	v522 = v516 + int32(19)
	v526 = v11 + int32(2256)
	v529 = int32(0)
	v534 = m.G0
	v536 = v534 - int32(16)
	m.G0 = v536
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[826]))) = v529
	*(*int32)(unsafe.Add(mBase, uint32(v526))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[827]))) = v529
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522))))
	if base.Ui32((v544-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v624 = v529
		goto L120
	} else {
		goto L121
	}
L117:
	;
	goto L115
L118:
	;
	v679 = F_ReadDir(m, v507, v11+int32(208))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L4
	} else {
		goto L148
	}
L119:
	;
	if v624 == int32(0) {
		goto L118
	} else {
		goto L138
	}
L120:
	;
	m.G0 = v536 + int32(16)
	goto L119
L121:
	;
	v551 = int32(4735140)
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	v557 = F_strtoul(m, v522, v536+int32(8), int32(10))
	mBase = m.M
	v559 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v559 != 0 {
		v624 = v529
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v536)+8))
	if v522 == v560 {
		v624 = v529
		goto L120
	} else {
		goto L123
	}
L123:
	;
	if v557 == int32(0) {
		v624 = v529
		goto L120
	} else {
		goto L124
	}
L124:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560))))
	if v564 != int32(95) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	if v580&int32(255) == int32(46) {
		goto L130
	} else {
		goto L131
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v536)+12)) = int32(0)
	v580 = v564
	v581 = v560
	goto L125
L127:
	;
	goto L128
L128:
	;
	v573 = F_forkname_chars(m, v560+int32(1), v536+int32(12))
	mBase = m.M
	if v573 <= int32(0) {
		v624 = v529
		goto L120
	} else {
		goto L129
	}
L129:
	;
	v578 = v573 + v560 + int32(1)
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578))))
	v580 = v579
	v581 = v578
	goto L125
L130:
	;
	v587 = v581 + int32(1)
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587))))
	if base.Ui32((v588-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v624 = v529
		goto L120
	} else {
		goto L133
	}
L131:
	;
	v611 = v529
	v612 = v580
	goto L132
L132:
	;
	if v612&int32(255) != 0 {
		v624 = v529
		goto L120
	} else {
		goto L137
	}
L133:
	;
	v595 = int32(4735140)
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	v601 = F_strtoul(m, v587, v536+int32(8), int32(10))
	mBase = m.M
	v603 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v603 != 0 {
		v624 = v529
		goto L120
	} else {
		goto L134
	}
L134:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v536)+8))
	if v587 == v604 {
		v624 = v529
		goto L120
	} else {
		goto L135
	}
L135:
	;
	if v601 == int32(0) {
		v624 = v529
		goto L120
	} else {
		goto L136
	}
L136:
	;
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604))))
	v611 = v601
	v612 = v608
	goto L132
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[826]))) = v557
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v536)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v526))) = v616
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[827]))) = v611
	v624 = int32(1)
	goto L120
L138:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v11)+2256))
	if v630 == int32(3) {
		goto L118
	} else {
		goto L139
	}
L139:
	;
	v635 = int32(0)
	v637 = F_hash_search(m, v282, v11+int32(5336), v635, v635)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	if v637 == int32(0) {
		goto L118
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+148)) = v522
	*(*int32)(unsafe.Add(mBase, uint32(v11)+144)) = v11 + int32(208)
	v651 = F_pg_snprintf(m, v11+int32(5344), int32(2048), int32(188238), v11+int32(144))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L4
	} else {
		goto L142
	}
L142:
	;
	v655 = F_unlink(m, v11+int32(5344))
	mBase = m.M
	if v655 < int32(0) {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v660 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L4
	} else {
		goto L144
	}
L144:
	;
	if v660 == int32(0) {
		goto L118
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v11 + int32(5344)
	F_errmsg_internal(m, int32(754169), v11+int32(128))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(518316), int32(264), int32(225286))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	goto L118
L148:
	;
	if v679 != 0 {
		v516 = v679
		goto L116
	} else {
		goto L149
	}
L149:
	;
	goto L117
L150:
	;
	F_hash_destroy(m, v282)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	if v41 == int32(0) {
		goto L19
	} else {
		goto L152
	}
L152:
	;
	goto L43
L153:
	;
	v709 = F_ReadDir(m, v705, v11+int32(208))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	if v709 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v714 = v709
	goto L158
L156:
	;
	goto L157
L157:
	;
	F_FreeDir(m, v705)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L4
	} else {
		goto L198
	}
L158:
	;
	v720 = v714 + int32(19)
	v727 = int32(0)
	v732 = m.G0
	v734 = v732 - int32(16)
	m.G0 = v734
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[826]))) = v727
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[827]))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[828]))) = v727
	v742 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v720))))
	if base.Ui32((v742-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v822 = v727
		goto L162
	} else {
		goto L163
	}
L159:
	;
	goto L157
L160:
	;
	v902 = F_ReadDir(m, v705, v11+int32(208))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L4
	} else {
		goto L196
	}
L161:
	;
	if v822 == int32(0) {
		goto L160
	} else {
		goto L180
	}
L162:
	;
	m.G0 = v734 + int32(16)
	goto L161
L163:
	;
	v749 = int32(4735140)
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	v755 = F_strtoul(m, v720, v734+int32(8), int32(10))
	mBase = m.M
	v757 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v757 != 0 {
		v822 = v727
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v734)+8))
	if v720 == v758 {
		v822 = v727
		goto L162
	} else {
		goto L165
	}
L165:
	;
	if v755 == int32(0) {
		v822 = v727
		goto L162
	} else {
		goto L166
	}
L166:
	;
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v758))))
	if v762 != int32(95) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	if v778&int32(255) == int32(46) {
		goto L172
	} else {
		goto L173
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v734)+12)) = int32(0)
	v778 = v762
	v779 = v758
	goto L167
L169:
	;
	goto L170
L170:
	;
	v771 = F_forkname_chars(m, v758+int32(1), v734+int32(12))
	mBase = m.M
	if v771 <= int32(0) {
		v822 = v727
		goto L162
	} else {
		goto L171
	}
L171:
	;
	v776 = v771 + v758 + int32(1)
	v777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v776))))
	v778 = v777
	v779 = v776
	goto L167
L172:
	;
	v785 = v779 + int32(1)
	v786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v785))))
	if base.Ui32((v786-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v822 = v727
		goto L162
	} else {
		goto L175
	}
L173:
	;
	v809 = v727
	v810 = v778
	goto L174
L174:
	;
	if v810&int32(255) != 0 {
		v822 = v727
		goto L162
	} else {
		goto L179
	}
L175:
	;
	v793 = int32(4735140)
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	v799 = F_strtoul(m, v785, v734+int32(8), int32(10))
	mBase = m.M
	v801 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v801 != 0 {
		v822 = v727
		goto L162
	} else {
		goto L176
	}
L176:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v734)+8))
	if v785 == v802 {
		v822 = v727
		goto L162
	} else {
		goto L177
	}
L177:
	;
	if v799 == int32(0) {
		v822 = v727
		goto L162
	} else {
		goto L178
	}
L178:
	;
	v806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v802))))
	v809 = v799
	v810 = v806
	goto L174
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[826]))) = v755
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v734)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[827]))) = v814
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[828]))) = v809
	v822 = int32(1)
	goto L162
L180:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[827])))
	if v828 != int32(3) {
		goto L160
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v11 + int32(208)
	v841 = F_pg_snprintf(m, v11+int32(3280), int32(2048), int32(188238), v11+int32(96))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[826])))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[828])))
	if v844 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v874 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L4
	} else {
		goto L189
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v11 + int32(208)
	v857 = F_pg_snprintf(m, v11+int32(2256), int32(1024), int32(42047), v11-int32(-64))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L4
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+88)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v11 + int32(208)
	v870 = F_pg_snprintf(m, v11+int32(2256), int32(1024), int32(42097), v11+int32(80))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L4
	} else {
		goto L188
	}
L187:
	;
	goto L183
L188:
	;
	goto L183
L189:
	;
	if v874 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v11 + int32(2256)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v11 + int32(3280)
	F_errmsg_internal(m, int32(194155), v11+int32(48))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L4
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	F_copy_file(m, v11+int32(3280), v11+int32(2256))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L4
	} else {
		goto L195
	}
L193:
	;
	F_errfinish(m, int32(518316), int32(314), int32(225286))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L4
	} else {
		goto L194
	}
L194:
	;
	goto L192
L195:
	;
	goto L160
L196:
	;
	if v902 != 0 {
		v714 = v902
		goto L158
	} else {
		goto L197
	}
L197:
	;
	goto L159
L198:
	;
	v916 = F_AllocateDir(m, v11+int32(208))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L4
	} else {
		goto L199
	}
L199:
	;
	v920 = F_ReadDir(m, v916, v11+int32(208))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	if v920 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v925 = v920
	goto L204
L202:
	;
	goto L203
L203:
	;
	F_FreeDir(m, v916)
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L4
	} else {
		goto L237
	}
L204:
	;
	v931 = v925 + int32(19)
	v933 = v11 + int32(2256)
	v938 = int32(0)
	v943 = m.G0
	v945 = v943 - int32(16)
	m.G0 = v945
	*(*int32)(unsafe.Add(mBase, uint32(v933))) = v938
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[827]))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[826]))) = v938
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v931))))
	if base.Ui32((v953-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v1033 = v938
		goto L208
	} else {
		goto L209
	}
L205:
	;
	goto L203
L206:
	;
	v1080 = F_ReadDir(m, v916, v11+int32(208))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L4
	} else {
		goto L235
	}
L207:
	;
	if v1033 == int32(0) {
		goto L206
	} else {
		goto L226
	}
L208:
	;
	m.G0 = v945 + int32(16)
	goto L207
L209:
	;
	v960 = int32(4735140)
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	v966 = F_strtoul(m, v931, v945+int32(8), int32(10))
	mBase = m.M
	v968 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v968 != 0 {
		v1033 = v938
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v945)+8))
	if v931 == v969 {
		v1033 = v938
		goto L208
	} else {
		goto L211
	}
L211:
	;
	if v966 == int32(0) {
		v1033 = v938
		goto L208
	} else {
		goto L212
	}
L212:
	;
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v969))))
	if v973 != int32(95) {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	if v989&int32(255) == int32(46) {
		goto L218
	} else {
		goto L219
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v945)+12)) = int32(0)
	v989 = v973
	v990 = v969
	goto L213
L215:
	;
	goto L216
L216:
	;
	v982 = F_forkname_chars(m, v969+int32(1), v945+int32(12))
	mBase = m.M
	if v982 <= int32(0) {
		v1033 = v938
		goto L208
	} else {
		goto L217
	}
L217:
	;
	v987 = v982 + v969 + int32(1)
	v988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v987))))
	v989 = v988
	v990 = v987
	goto L213
L218:
	;
	v996 = v990 + int32(1)
	v997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v996))))
	if base.Ui32((v997-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v1033 = v938
		goto L208
	} else {
		goto L221
	}
L219:
	;
	v1020 = v938
	v1021 = v989
	goto L220
L220:
	;
	if v1021&int32(255) != 0 {
		v1033 = v938
		goto L208
	} else {
		goto L225
	}
L221:
	;
	v1004 = int32(4735140)
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	v1010 = F_strtoul(m, v996, v945+int32(8), int32(10))
	mBase = m.M
	v1012 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v1012 != 0 {
		v1033 = v938
		goto L208
	} else {
		goto L222
	}
L222:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v945)+8))
	if v996 == v1013 {
		v1033 = v938
		goto L208
	} else {
		goto L223
	}
L223:
	;
	if v1010 == int32(0) {
		v1033 = v938
		goto L208
	} else {
		goto L224
	}
L224:
	;
	v1017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1013))))
	v1020 = v1010
	v1021 = v1017
	goto L220
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v933))) = v966
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v945)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[827]))) = v1025
	*(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[826]))) = v1020
	v1033 = int32(1)
	goto L208
L226:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[827])))
	if v1039 != int32(3) {
		goto L206
	} else {
		goto L227
	}
L227:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v11)+2256))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[826])))
	if v1043 == int32(0) {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	F_fsync_fname(m, v11+int32(3280), int32(0))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L4
	} else {
		goto L234
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v1042
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v11 + int32(208)
	v1056 = F_pg_snprintf(m, v11+int32(3280), int32(1024), int32(42047), v11+int32(16))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L4
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v1042
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v11 + int32(208)
	v1069 = F_pg_snprintf(m, v11+int32(3280), int32(1024), int32(42097), v11+int32(32))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L4
	} else {
		goto L233
	}
L232:
	;
	goto L228
L233:
	;
	goto L228
L234:
	;
	goto L206
L235:
	;
	if v1080 != 0 {
		v925 = v1080
		goto L204
	} else {
		goto L236
	}
L236:
	;
	goto L205
L237:
	;
	F_fsync_fname(m, v11+int32(208), int32(1))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L4
	} else {
		goto L238
	}
L238:
	;
	goto L19
L239:
	;
	if v1105 != 0 {
		v45 = v1105
		goto L17
	} else {
		goto L240
	}
L240:
	;
	goto L18
L241:
	;
	goto L2
L242:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L4
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v11 + int32(5344)
	F_errmsg(m, int32(314276), v11+int32(112))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L4
	} else {
		goto L244
	}
L244:
	;
	F_errfinish(m, int32(518316), int32(262), int32(225286))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L4
	} else {
		goto L245
	}
L245:
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
	v2 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	if v2 != 0 {
		v4 = int32(0)
		v9 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		v14 = F_set_config_with_handle(m, int32(338457), v4, int32(248145), int32(6), int32(13), v9, int32(2), int32(1), v4, v4)
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
	var v7 int32
	_ = v7
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v7 < v9 {
		v62 = v2
		return v62
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v7
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
		v16 = v7 - int32(1)
		if v16 <= v9 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v13
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v16))))
			switch v20 - int32(100) {
			case 0, 16:
				v28 = F_find_among_b(m, l0, int32(4240480), int32(4))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					if v28 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v13
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v13
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v38
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v42 = v40 + (v7 - v11)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v42
						if v42 <= v13 {
							v62 = v2
							return v62
						} else {
							v45 = int32(1)
							v46 = v42 - v45
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v46
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v46
							v50 = F_slice_del(m, l0)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								if int32(0) <= v50 {
									v57 = v45
								} else {
									v57 = v50 >> (uint(int32(31)) % 32) & v50
								}
								v62 = v57
								return v62
							}
						}
					}
				}
			default:
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v13
				return int32(0)
			}
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
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
	var v171 int32
	_ = v171
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	if v5 < v7 {
		v171 = v2
		return v171
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v26 <= v27 {
			v132 = int32(-1)
			v139 = v132
		} else {
			v44 = int32(1)
			v45 = v26 - v44
			v47 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23+v45))))
			v49 = v47 & int32(255)
			if v45 == v27 {
				v104 = v49
				v105 = v44
			} else {
				if int32(0) <= v47 {
					v104 = v49
					v105 = v44
				} else {
					v55 = v49 & int32(63)
					v57 = v26 - int32(2)
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v57))))
					v61 = v59 << (uint(int32(6)) % 32)
					if base.B2i32(v57 != v27)&base.B2i32(base.Ui32(v59) < base.Ui32(int32(192))) == int32(0) {
						v104 = v61&int32(1984) | v55
						v105 = int32(2)
					} else {
						v74 = v61&int32(4032) | v55
						v76 = v26 - int32(3)
						v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v76))))
						if base.B2i32(v76 != v27)&base.B2i32(base.Ui32(v78) < base.Ui32(int32(224))) == int32(0) {
							v104 = v78<<(uint(int32(12))%32)&int32(61440) | v74
							v105 = int32(3)
						} else {
							v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+(v23-int32(4))))))
							v104 = v78<<(uint(int32(12))%32)&int32(258048) | v96&int32(7)<<(uint(int32(18))%32) | v74
							v105 = int32(4)
						}
					}
				}
			}
			if int32(232) < v104 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v26 - v105
				v132 = int32(0)
				v139 = v132
			} else {
				v109 = v104 - int32(97)
				if v109 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v26 - v105
					v132 = int32(0)
					v139 = v132
				} else {
					v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v109)>>(uint(int32(3))%32)))+uint32(_consts[1295]))))
					if int32(base.Ui32(v115)>>(uint(v109&int32(7))%32))&int32(1) == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v26 - v105
						v132 = int32(0)
						v139 = v132
					} else {
						v139 = v105
					}
				}
			}
		}
		if v139 != 0 {
			v171 = v2
			return v171
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
				v155 = F_memcmp(m, v152+v142-v144, int32(2231761), v144)
				mBase = m.M
				if v155 != 0 {
					v159 = v146
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v142 - v144
					v159 = int32(1)
				}
			}
			if v159 != 0 {
				v171 = v2
				return v171
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
						v171 = v163
						return v171
					} else {
						v169 = F_r_undouble(m, l0)
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
							return int32(0)
						} else {
							v171 = v169
							return v171
						}
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v4
	v7 = v4 + int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v8 <= v7 {
		v66 = v2
		return v66
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v7))))
		if v12 != int32(101) {
			v66 = v2
			return v66
		} else {
			v17 = F_find_among(m, l0, int32(4343072), int32(6))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v17 == int32(0) {
					v66 = v2
					return v66
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v23
					v25 = int32(1)
					switch v17 - v25 {
					case 0:
						v28 = F_slice_del(m, l0)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							if v28 < int32(0) {
								v66 = v28
							} else {
								v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(2)
								v58 = v32
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								v61 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v60 - v61
								v66 = v61
							}
							return v66
						}
					case 1:
						v37 = F_slice_from_s(m, l0, int32(4), int32(2246069))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							if v37 < int32(0) {
								v66 = v37
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								v58 = v41
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								v61 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v60 - v61
								v66 = v61
							}
							return v66
						}
					case 2:
						v42 = F_slice_del(m, l0)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							if v42 < int32(0) {
								v66 = v42
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v46))) = int32(4)
								v58 = v46
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								v61 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v60 - v61
								v66 = v61
							}
							return v66
						}
					case 3:
						v51 = F_slice_from_s(m, l0, int32(4), int32(2246073))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							if v51 < int32(0) {
								v66 = v51
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v55))) = int32(4)
								v58 = v55
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								v61 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v60 - v61
								v66 = v61
							}
							return v66
						}
					default:
						v66 = v25
						return v66
					}
				}
			}
		}
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
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v32)>>(uint(int32(3))%32)))+uint32(_consts[1290]))))
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
					v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v189)>>(uint(int32(3))%32)))+uint32(_consts[1291]))))
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
						v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v241)>>(uint(int32(3))%32)))+uint32(_consts[1291]))))
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
					v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v84)>>(uint(int32(3))%32)))+uint32(_consts[1291]))))
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
						v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v189)>>(uint(int32(3))%32)))+uint32(_consts[1291]))))
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
							v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v241)>>(uint(int32(3))%32)))+uint32(_consts[1291]))))
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
						v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v134)>>(uint(int32(3))%32)))+uint32(_consts[1291]))))
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
							v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v189)>>(uint(int32(3))%32)))+uint32(_consts[1291]))))
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
								v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v241)>>(uint(int32(3))%32)))+uint32(_consts[1291]))))
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v299 int32
	_ = v299
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
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
	F_errstart_cold(m, int32(21), int32(584727))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L106
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(584727))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L101
	}
L5:
	;
	F_plpgsql_yyerror(m, l3, int32(0), l4, int32(223848))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L100
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(584727))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L95
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(584727))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L90
	}
L8:
	;
	m.G0 = v16 + int32(128)
	return v385
L9:
	;
	if v18 == int32(40) {
		goto L3
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if v18 != int32(40) {
		goto L4
	} else {
		goto L15
	}
L12:
	;
	if l1 == v18 {
		v385 = int32(0)
		goto L8
	} else {
		goto L13
	}
L13:
	;
	F_plpgsql_yyerror(m, l3, int32(0), l4, int32(223848))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[1272]))
	v37 = int32(2)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v22<<(uint(v37)%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+28))
	v44 = F_palloc0(m, v41<<(uint(v37)%32))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+28))
	if int32(0) < v46 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v58 = v6
	v59 = v6
	goto L20
L18:
	;
	v279 = v6
	goto L19
L19:
	;
	F_initStringInfo(m, v16+int32(112))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L68
	}
L20:
	;
	F_plpgsql_peek2(m, v16+int32(108), v16+int32(104), v16+int32(100), int32(0), l4)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	v279 = v213
	goto L19
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v16)+108))
	if v71 != int32(258) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v219 = v44 + v210<<(uint(int32(2))%32)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	if v220 != 0 {
		goto L7
	} else {
		goto L55
	}
L24:
	;
	v210 = v59
	v213 = v58
	goto L23
L25:
	;
	goto L26
L26:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v16)+104))
	if v74&int32(-2) != int32(270) {
		v210 = v59
		v213 = v58
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v79 = int32(4655372)
	v80 = *(*int32)(unsafe.Add(mBase, _consts[1280]))
	*(*int32)(unsafe.Add(mBase, _consts[1280])) = int32(1)
	v84 = F_plpgsql_yylex(m, l2, l3, l4)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, _consts[1280])) = v80
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v40)+28))
	if v89 <= int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(584727))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L50
	}
L30:
	;
	if v148 == v89 {
		goto L29
	} else {
		goto L46
	}
L31:
	;
	v148 = int32(0)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	v101 = int32(0)
	goto L34
L34:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v93+v101<<(uint(int32(2))%32))))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	if v115 == int32(0) {
		v134 = v114
		v135 = v115
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L29
L36:
	;
	if v135-v134 == int32(0) {
		v148 = v101
		goto L30
	} else {
		goto L44
	}
L37:
	;
	goto L36
L38:
	;
	if v114 != v115 {
		v134 = v114
		v135 = v115
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v119 = v111
	v120 = v86
	goto L40
L40:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+1)))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
	if v124 == int32(0) {
		v134 = v123
		v135 = v124
		goto L37
	} else {
		goto L42
	}
L41:
	;
	v134 = v123
	v135 = v124
	goto L37
L42:
	;
	v127 = int32(1)
	if v123 == v124 {
		v119 = v119 + v127
		v120 = v120 + v127
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v140 = v101 + int32(1)
	if v140 != v89 {
		v101 = v140
		goto L34
	} else {
		goto L45
	}
L45:
	;
	goto L35
L46:
	;
	v156 = F_plpgsql_yylex(m, l2, l3, l4)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+104)) = v156
	if base.Ui32(int32(-3)) < base.Ui32(v156-int32(272)) {
		v210 = v148
		v213 = int32(1)
		goto L23
	} else {
		goto L48
	}
L48:
	;
	F_plpgsql_yyerror(m, l3, int32(0), l4, int32(223848))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v188
	F_errmsg(m, int32(760108), v16+int32(16))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v197 = F_plpgsql_scanner_errposition(m, v196, l4)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(27578), int32(3986), int32(165606))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
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
	v223 = int32(0)
	v226 = int32(1)
	v231 = F_read_sql_construct(m, int32(44), int32(41), v223, int32(719152), int32(2), v226, v226, v223, v16+int32(112), l2, l3, l4)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	*(*int32)(unsafe.Add(mBase, uint32(v219))) = v233
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v40)+28))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v16)+112))
	switch v236 - int32(41) {
	case 0:
		goto L59
	default:
		goto L57
	case 3:
		goto L58
	}
L57:
	;
	v268 = v59 + int32(1)
	if v268 < v235 {
		v58 = v213
		v59 = v268
		goto L20
	} else {
		goto L67
	}
L58:
	;
	if v59 == v235-int32(1) {
		goto L6
	} else {
		goto L66
	}
L59:
	;
	if v59 == v235-int32(1) {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(584727))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v249
	F_errmsg(m, int32(736718), v16+int32(48))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v257 = F_plpgsql_scanner_errposition(m, v256, l4)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(27578), int32(4028), int32(165606))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	goto L57
L67:
	;
	goto L21
L68:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v40)+28))
	if int32(0) < v287 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v299 = int32(0)
	goto L72
L70:
	;
	goto L71
L71:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v16)+112))
	v356 = F_palloc0(m, int32(80))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L85
	}
L72:
	;
	v309 = v299 << (uint(int32(2)) % 32)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v44+v309)))
	F_appendStringInfoString(m, v16+int32(112), v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	goto L71
L74:
	;
	if v279&int32(1) != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v314+v309)))
	v317 = F_quote_identifier(m, v316)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v40)+28))
	if v299 < v327-int32(1) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v317
	F_appendStringInfo(m, v16+int32(112), int32(209214), v16+int32(32))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	F_appendStringInfoString(m, v16+int32(112), int32(783295))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L83
	}
L81:
	;
	v337 = v327
	goto L82
L82:
	;
	v339 = v299 + int32(1)
	if v339 < v337 {
		v299 = v339
		goto L72
	} else {
		goto L84
	}
L83:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v40)+28))
	v337 = v336
	goto L82
L84:
	;
	goto L73
L85:
	;
	v358 = F_pstrdup(m, v354)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v356)+4)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v356))) = v358
	v364 = *(*int32)(unsafe.Add(mBase, _consts[1261]))
	*(*int32)(unsafe.Add(mBase, uint32(v356)+8)) = v364
	v367 = *(*int32)(unsafe.Add(mBase, _consts[1267]))
	v368 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+20)) = uint8(v368)
	*(*int32)(unsafe.Add(mBase, uint32(v356)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v356)+12)) = v367
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v16)+112))
	F_pfree(m, v373)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v376 = F_plpgsql_yylex(m, l2, l3, l4)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	if v376 != l1 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	v385 = v356
	goto L8
L90:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v403+v210<<(uint(int32(2))%32))))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v407
	F_errmsg(m, int32(436578), v16+int32(80))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
	v417 = F_plpgsql_scanner_errposition(m, v416, l4)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(27578), int32(4006), int32(165606))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v431
	F_errmsg(m, int32(736683), v16-int32(-64))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v439 = F_plpgsql_scanner_errposition(m, v438, l4)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(27578), int32(4035), int32(165606))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v457
	F_errmsg(m, int32(129585), v16+int32(96))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v465 = F_plpgsql_scanner_errposition(m, v464, l4)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(27578), int32(3941), int32(165606))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v479
	F_errmsg(m, int32(129745), v16)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v485 = F_plpgsql_scanner_errposition(m, v484, l4)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(27578), int32(3927), int32(165606))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
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
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
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
						F_plpgsql_yyerror(m, l3, int32(0), l4, int32(223848))
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
									F_errstart_cold(m, int32(21), int32(584727))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										F_errcode(m, int32(16801924))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return
										} else {
											F_errmsg(m, int32(81657), int32(0))
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return
											} else {
												v79 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
												v80 = F_plpgsql_scanner_errposition(m, v79, l4)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return
												} else {
													F_errfinish(m, int32(27578), int32(3626), int32(115532))
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
							F_plpgsql_yyerror(m, l3, int32(0), l4, int32(223848))
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
										F_errstart_cold(m, int32(21), int32(584727))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											F_errcode(m, int32(16801924))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return
											} else {
												F_errmsg(m, int32(81657), int32(0))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return
												} else {
													v79 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
													v80 = F_plpgsql_scanner_errposition(m, v79, l4)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return
													} else {
														F_errfinish(m, int32(27578), int32(3626), int32(115532))
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
					F_plpgsql_yyerror(m, l3, int32(0), l4, int32(223848))
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
								F_errstart_cold(m, int32(21), int32(584727))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									F_errcode(m, int32(16801924))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return
									} else {
										F_errmsg(m, int32(81657), int32(0))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return
										} else {
											v79 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
											v80 = F_plpgsql_scanner_errposition(m, v79, l4)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return
											} else {
												F_errfinish(m, int32(27578), int32(3626), int32(115532))
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
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
	v15 = *(*int32)(unsafe.Add(mBase, _consts[47]))
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = int32(251189)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = l1
	v30 = F_pg_snprintf(m, v10+int32(112), int32(1024), int32(188238), v10+int32(96))
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
	v35 = F_OpenTransientFile(m, v10+int32(112), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L8
	}
L7:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(167772200)
	v63 = int32(524)
	v64 = F_read(m, v35, l0, v63)
	mBase = m.M
	if v64 == v63 {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	if int32(0) <= v35 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v40 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	if v40 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v10 + int32(112)
	F_errmsg(m, int32(313722), v10+int32(80))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(521070), int32(819), int32(406602))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L7
L15:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(0)
	v112 = F_CloseTransientFile(m, v35)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L30
	}
L16:
	;
	v68 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	if v64 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	F_errfinish(m, int32(521070), v103, int32(406602))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L28
	}
L19:
	;
	if v68 == int32(0) {
		goto L15
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v68 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L22:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v10 + int32(112)
	F_errmsg(m, int32(314641), v10+int32(48))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v103 = int32(829)
	goto L18
L25:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = int32(524)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v10 + int32(112)
	F_errmsg(m, int32(38923), v10-int32(-64))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v103 = int32(834)
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
	if v112 == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v117 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	if v117 == int32(0) {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v10 + int32(112)
	F_errmsg(m, int32(314465), v10+int32(32))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(521070), int32(842), int32(406602))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	goto L29
L37:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v139+int32(3200))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v144 == int32(5842711) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L39
L41:
	;
	v168 = int32(-1)
	v170 = m.Env.Pgmem_crc32c(m, v168, l0, int32(520))
	mBase = m.M
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
	if v170^v171 == v168 {
		goto L50
	} else {
		goto L51
	}
L42:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v147) < base.Ui32(int32(65)) {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v151 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	if v151 == int32(0) {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v10 + int32(112)
	F_errmsg(m, int32(532259), v10+int32(16))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(521070), int32(853), int32(406602))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
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
	v176 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	if v176 == int32(0) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v10 + int32(112)
	F_errmsg(m, int32(301034), v10)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(521070), int32(863), int32(406602))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
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
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = l3 - int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v13
	if v13 == v5 {
		v45 = v5
		v46 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v45
		v48 = v45
		v49 = v46
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v48
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v49)
		v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
		if v52&int32(1) != 0 {
			v58 = F_LogicalTapeRead(m, l2, v10+int32(12), int32(4))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				if v58 != int32(4) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(532168), int32(0))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							F_errfinish(m, int32(518771), int32(2088), int32(300850))
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
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
		if v18 == int32(0) {
			v23 = F_LogicalTapeRead(m, l2, l1+int32(4), v13)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				if v23 == v13 {
					v48 = v5
					v49 = v5
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v48
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v49)
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
					if v52&int32(1) != 0 {
						v58 = F_LogicalTapeRead(m, l2, v10+int32(12), int32(4))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							if v58 != int32(4) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									F_errmsg_internal(m, int32(532168), int32(0))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										F_errfinish(m, int32(518771), int32(2088), int32(300850))
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
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(532168), int32(0))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							F_errfinish(m, int32(518771), int32(2073), int32(300850))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
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
			v39 = F_tuplesort_readtup_alloc(m, l0, v13)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				v41 = F_LogicalTapeRead(m, l2, v39, v13)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					if v41 != v13 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(532168), int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								F_errfinish(m, int32(518771), int32(2081), int32(300850))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v45 = v39
						v46 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v45
						v48 = v45
						v49 = v46
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v48
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v49)
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
						if v52&int32(1) != 0 {
							v58 = F_LogicalTapeRead(m, l2, v10+int32(12), int32(4))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								if v58 != int32(4) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										F_errmsg_internal(m, int32(532168), int32(0))
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											F_errfinish(m, int32(518771), int32(2088), int32(300850))
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
								F_errmsg_internal(m, int32(532168), int32(0))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									F_errfinish(m, int32(518771), int32(1972), int32(290984))
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
					F_errmsg_internal(m, int32(532168), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						F_errfinish(m, int32(518771), int32(1970), int32(290984))
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
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	v4 = int32(0)
	v5 = F_strlen(m, l0)
	mBase = m.M
	if v5 != int32(16) {
		v124 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v125 = F_strlen(m, l1)
	mBase = m.M
	if v125 == int32(16) {
		goto L36
	} else {
		goto L37
	}
L2:
	;
	v8 = int32(564757)
	v12 = m.G0
	v14 = v12 - int32(32)
	v15 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v15
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _consts[273])))
	if v23 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v91 != int32(8) {
		v124 = v4
		goto L1
	} else {
		goto L24
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
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _consts[274])))
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
		v83 = l0
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v91 = v83 - l0
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
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v83 = v79
	goto L16
L20:
	;
	v83 = v62
	goto L16
L21:
	;
	goto L22
L22:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	v79 = v62 + int32(1)
	if v77 != 0 {
		v62 = v79
		v63 = v77
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	v95 = l0 + int32(8)
	v96 = int32(13035)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, _consts[364])))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v100 == int32(0) {
		v119 = v99
		v120 = v100
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v124 = base.B2i32(v120-v119 == int32(0))
	goto L1
L26:
	;
	goto L25
L27:
	;
	if v99 != v100 {
		v119 = v99
		v120 = v100
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v104 = v95
	v105 = v96
	goto L29
L29:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+1)))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
	if v109 == int32(0) {
		v119 = v108
		v120 = v109
		goto L26
	} else {
		goto L31
	}
L30:
	;
	v119 = v108
	v120 = v109
	goto L26
L31:
	;
	v112 = int32(1)
	if v108 == v109 {
		v104 = v104 + v112
		v105 = v105 + v112
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v252 == int32(0) {
		v271 = v251
		v272 = v252
		goto L75
	} else {
		goto L76
	}
L34:
	;
	if v124 != 0 {
		goto L71
	} else {
		goto L72
	}
L35:
	;
	v215 = l1 + int32(8)
	v216 = int32(13035)
	v219 = int32(*(*uint8)(unsafe.Add(mBase, _consts[364])))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if v220 == int32(0) {
		v239 = v219
		v240 = v220
		goto L63
	} else {
		goto L64
	}
L36:
	;
	v128 = int32(564757)
	v132 = m.G0
	v134 = v132 - int32(32)
	v135 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v134)+24)) = v135
	*(*int64)(unsafe.Add(mBase, uint32(v134)+16)) = v135
	*(*int64)(unsafe.Add(mBase, uint32(v134)+8)) = v135
	*(*int64)(unsafe.Add(mBase, uint32(v134))) = v135
	v143 = int32(*(*uint8)(unsafe.Add(mBase, _consts[273])))
	if v143 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	goto L38
L38:
	;
	if v124 != 0 {
		goto L34
	} else {
		goto L61
	}
L39:
	;
	if v211 == int32(8) {
		goto L35
	} else {
		goto L60
	}
L40:
	;
	v211 = int32(0)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, _consts[274])))
	if v147 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v151 = l1
	goto L46
L44:
	;
	goto L45
L45:
	;
	v161 = v128
	v162 = v143
	goto L49
L46:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v157 == v143 {
		v151 = v151 + int32(1)
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v211 = v151 - l1
	goto L39
L48:
	;
	goto L47
L49:
	;
	v169 = v134 + int32(base.Ui32(v162)>>(uint(int32(3))%32))&int32(28)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v171 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = v170 | v171<<(uint(v162)%32)
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
	if v175 != 0 {
		v161 = v161 + v171
		v162 = v175
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v178 == int32(0) {
		v203 = l1
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L50
L52:
	;
	v211 = v203 - l1
	goto L39
L53:
	;
	v182 = l1
	v183 = v178
	goto L54
L54:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v134+int32(base.Ui32(v183)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v191)>>(uint(v183)%32))&int32(1) == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v203 = v199
	goto L52
L56:
	;
	v203 = v182
	goto L52
L57:
	;
	goto L58
L58:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
	v199 = v182 + int32(1)
	if v197 != 0 {
		v182 = v199
		v183 = v197
		goto L54
	} else {
		goto L59
	}
L59:
	;
	goto L55
L60:
	;
	goto L38
L61:
	;
	goto L33
L62:
	;
	if v124 == base.B2i32(v240-v239 == int32(0)) {
		goto L33
	} else {
		goto L70
	}
L63:
	;
	goto L62
L64:
	;
	if v219 != v220 {
		v239 = v219
		v240 = v220
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v224 = v215
	v225 = v216
	goto L66
L66:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+1)))
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+1)))
	if v229 == int32(0) {
		v239 = v228
		v240 = v229
		goto L63
	} else {
		goto L68
	}
L67:
	;
	v239 = v228
	v240 = v229
	goto L63
L68:
	;
	v232 = int32(1)
	if v228 == v229 {
		v224 = v224 + v232
		v225 = v225 + v232
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	goto L34
L71:
	;
	v247 = int32(-1)
	goto L73
L72:
	;
	v247 = int32(1)
	goto L73
L73:
	;
	return v247
L74:
	;
	return v272 - v271
L75:
	;
	goto L74
L76:
	;
	if v251 != v252 {
		v271 = v251
		v272 = v252
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v256 = l0
	v257 = l1
	goto L78
L78:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+1)))
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+1)))
	if v261 == int32(0) {
		v271 = v260
		v272 = v261
		goto L75
	} else {
		goto L80
	}
L79:
	;
	v271 = v260
	v272 = v261
	goto L75
L80:
	;
	v264 = int32(1)
	if v260 == v261 {
		v256 = v256 + v264
		v257 = v257 + v264
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
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
	var v17 int32
	_ = v17
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
	var v39 int32
	_ = v39
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(int32(2)) <= base.Ui32(v9) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = v9
	v16 = int32(0)
	v17 = int32(1)
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
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v17<<(uint(int32(2))%32))))
	if v25 == int32(0) {
		v141 = v15
		v142 = v16
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v148 = v17 + int32(1)
	if base.Ui32(v148) < base.Ui32(v141) {
		v15 = v141
		v16 = v142
		v17 = v148
		goto L4
	} else {
		goto L48
	}
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v28 != 0 {
		v141 = v15
		v142 = v16
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+212))
	if v29 == int32(0) {
		v141 = v15
		v142 = v16
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
	v39 = v16
	goto L13
L11:
	;
	v133 = v16
	goto L12
L12:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v141 = v138
	v142 = v133
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
	v133 = v124
	goto L12
L15:
	;
	v127 = v38 + int32(1)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v127 < v128 {
		v38 = v127
		v39 = v124
		goto L13
	} else {
		goto L47
	}
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+56))
	v54 = F_bms_is_member(m, v53, v39)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v59 = v39
	goto L18
L18:
	;
	if v49 == int32(0) {
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
		v124 = v39
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)+56))
	v57 = F_bms_add_member(m, v39, v56)
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
	if v106 != int32(2) {
		v124 = v59
		goto L15
	} else {
		goto L39
	}
L24:
	;
	v106 = int32(0)
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
	v75 = int32(0)
	v77 = v75
	v78 = v75
	goto L30
L30:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v49+int32(8)+v77<<(uint(int32(2))%32))))
	if v86 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v106 = v99
	goto L23
L32:
	;
	goto L31
L33:
	;
	v87 = int32(2)
	if v78 != 0 {
		v99 = v87
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
	v95 = v77 + int32(1)
	if v95 != v72 {
		v77 = v95
		v78 = v92
		goto L30
	} else {
		goto L38
	}
L36:
	;
	v88 = int32(1)
	if base.Ui32(v88) < base.Ui32(base.I32_popcnt(v86)) {
		v99 = v87
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v92 = v88
	goto L35
L38:
	;
	v99 = v92
	goto L32
L39:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v111 = F_pull_var_clause(m, v109, int32(26))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L19
	} else {
		goto L40
	}
L40:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+12)))
	if v113 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v117 = F_bms_intersect(m, v49, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L19
	} else {
		goto L44
	}
L42:
	;
	v119 = v49
	goto L43
L43:
	;
	F_add_vars_to_attr_needed(m, l0, v111, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L19
	} else {
		goto L45
	}
L44:
	;
	v119 = v117
	goto L43
L45:
	;
	F_list_free(m, v111)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L19
	} else {
		goto L46
	}
L46:
	;
	v124 = v59
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 float64
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
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
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
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
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v390 int32
	_ = v390
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
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
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v526 int32
	_ = v526
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v560 int32
	_ = v560
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v620 int32
	_ = v620
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v636 int32
	_ = v636
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 float64
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
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
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v725 int32
	_ = v725
	var v731 int32
	_ = v731
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
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v900 float64
	_ = v900
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v911 float64
	_ = v911
	var v914 int32
	_ = v914
	var v917 float64
	_ = v917
	var v918 float64
	_ = v918
	var v923 int32
	_ = v923
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1003 float64
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1011 float64
	_ = v1011
	var v1016 float64
	_ = v1016
	var v1017 float64
	_ = v1017
	var v1019 float64
	_ = v1019
	var v1022 float64
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1041 int32
	_ = v1041
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1067 int32
	_ = v1067
	var v1071 int32
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1123 int32
	_ = v1123
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1157 int32
	_ = v1157
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1217 int32
	_ = v1217
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1233 int32
	_ = v1233
	var v1239 int32
	_ = v1239
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1256 int32
	_ = v1256
	var v1260 int32
	_ = v1260
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1282 int32
	_ = v1282
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1292 int32
	_ = v1292
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1312 int32
	_ = v1312
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1331 int32
	_ = v1331
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1346 int32
	_ = v1346
	var v1353 int32
	_ = v1353
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1406 int32
	_ = v1406
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1422 int32
	_ = v1422
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1475 float64
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 float64
	_ = v1487
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1498 int32
	_ = v1498
	var v1500 int32
	_ = v1500
	var v1503 int32
	_ = v1503
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1516 int32
	_ = v1516
	var v1522 int32
	_ = v1522
	var v1526 int32
	_ = v1526
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1564 int32
	_ = v1564
	var v1567 int32
	_ = v1567
	var v1571 int32
	_ = v1571
	var v1574 int32
	_ = v1574
	var v1580 int32
	_ = v1580
	var v1602 int32
	_ = v1602
	var v1606 int32
	_ = v1606
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1642 int32
	_ = v1642
	var v1646 int32
	_ = v1646
	var v1673 int32
	_ = v1673
	var v1700 int32
	_ = v1700
	var v1714 int32
	_ = v1714
	var v1741 int32
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1747 int32
	_ = v1747
	var v1764 int32
	_ = v1764
	var v1773 int32
	_ = v1773
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1786 int32
	_ = v1786
	var v1820 int32
	_ = v1820
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1829 int32
	_ = v1829
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1844 int32
	_ = v1844
	var v1848 int32
	_ = v1848
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1864 int32
	_ = v1864
	var v1867 int32
	_ = v1867
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1874 int32
	_ = v1874
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1884 int32
	_ = v1884
	var v1889 int32
	_ = v1889
	var v1894 int32
	_ = v1894
	var v1895 float64
	_ = v1895
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1903 int32
	_ = v1903
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1909 int32
	_ = v1909
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1917 int32
	_ = v1917
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1923 int32
	_ = v1923
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1939 int32
	_ = v1939
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1954 int32
	_ = v1954
	var v1963 int32
	_ = v1963
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1973 int32
	_ = v1973
	var v1977 int32
	_ = v1977
	var v1987 int32
	_ = v1987
	var v2006 int32
	_ = v2006
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2014 int32
	_ = v2014
	var v2017 int32
	_ = v2017
	var v2022 int32
	_ = v2022
	var v2033 int32
	_ = v2033
	var v2048 int32
	_ = v2048
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2064 int32
	_ = v2064
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2072 int32
	_ = v2072
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2091 int32
	_ = v2091
	var v2136 int32
	_ = v2136
	var v2165 int32
	_ = v2165
	var v2167 int32
	_ = v2167
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2173 int32
	_ = v2173
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2189 int32
	_ = v2189
	var v2215 int32
	_ = v2215
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2257 int32
	_ = v2257
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2264 int32
	_ = v2264
	var v2292 int32
	_ = v2292
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2334 int32
	_ = v2334
	var v2335 int32
	_ = v2335
	var v2339 int32
	_ = v2339
	var v2341 int32
	_ = v2341
	var v2351 int32
	_ = v2351
	v9 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(80)
	m.G0 = v31
	v33 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v33)
	F_check_stack_depth(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v39 != int32(142) {
		goto L11
	} else {
		goto L12
	}
L3:
	;
	m.G0 = v31 + int32(80)
	return v2351
L4:
	;
	v2006 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v2008 = F_tlist_same_datatypes(m, v2006, l3, int32(0))
	mBase = m.M
	v2009 = m.ExcPending
	if v2009 != 0 {
		goto L1
	} else {
		goto L537
	}
L5:
	;
	v1467 = F_fetch_upper_rel(m, l1, int32(0), v1447)
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L1
	} else {
		goto L411
	}
L6:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v983)+48))
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v984)+48))
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v984)+8))
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v983)+8))
	v993 = F_bms_union(m, v991, v992)
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L1
	} else {
		goto L244
	}
L7:
	;
	v983 = v699
	v984 = v690
	v985 = v706
	goto L6
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L1
	} else {
		goto L241
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L1
	} else {
		goto L233
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L1
	} else {
		goto L230
	}
L11:
	;
	if v39 != int32(63) {
		goto L8
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v76 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44+v45<<(uint(int32(2))%32))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+36))
	v52 = F_build_simple_rel(m, l1, v45, int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v56 = *(*float64)(unsafe.Add(mBase, uint32(l1)+296))
	v57 = F_subquery_planner(m, v54, v50, l1, int32(0), v56, l2)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+140)) = v57
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v60 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)+264))
	v66 = F_generate_setop_tlist(m, l3, l4, v61, int32(1), v63, l5, v31+int32(56))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v68 = F_make_pathtarget_from_tlist(m, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v70 = F_set_pathtarget_cost_width(m, l1, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+28)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v66
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+56)))
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v74)
	v2351 = v52
	goto L3
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v31)+56)) = l0
	v81 = int32(0)
	v85 = F_list_make1_impl(m, int32(1), v31+int32(12))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v680 = *(*float64)(unsafe.Add(mBase, uint32(l1)+296))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+296)) = int64(0)
	v683 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v690 = F_recurse_set_operations(m, v683, l1, l0, v684, v685, l5, v31+int32(76), v31+int32(71))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L171
	}
L24:
	;
	if v85 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v89 = v81
	v95 = v85
	v99 = v9
	v101 = v9
	goto L28
L26:
	;
	v176 = v81
	v186 = v9
	v188 = v9
	goto L27
L27:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v204 = F_generate_append_tlist(m, v202, v203, v186, l5)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L50
	}
L28:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v117 = F_list_delete_first(m, v95)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	v176 = v169
	v186 = v171
	v188 = v172
	goto L27
L30:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v119 != int32(142) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v170 != 0 {
		v89 = v169
		v95 = v170
		v99 = v171
		v101 = v172
		goto L28
	} else {
		goto L49
	}
L32:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v151 != 0 {
		goto L42
	} else {
		goto L43
	}
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v122 != v123 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+8)))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if base.B2i32(v125 != v126)&base.B2i32(v125 == int32(0)) != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v116)+20))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v133 = F_equal(m, v131, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v133 == int32(0) {
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v116)+28))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v139 = F_equal(m, v137, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v139 == int32(0) {
		goto L32
	} else {
		goto L39
	}
L39:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	v144 = F_lcons(m, v143, v117)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v147 = F_lcons(m, v146, v144)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v169 = v89
	v170 = v147
	v171 = v99
	v172 = v101
	goto L31
L42:
	;
	v152 = int32(0)
	goto L44
L43:
	;
	v152 = l0
	goto L44
L44:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v159 = F_recurse_set_operations(m, v116, l1, v152, v153, v154, l5, v31+int32(48), v31+int32(76))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v161 = F_lappend(m, v89, v159)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	v164 = F_lappend(m, v99, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+76)))
	v167 = F_lappend_int(m, v101, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v169 = v161
	v170 = v117
	v171 = v164
	v172 = v167
	goto L31
L49:
	;
	goto L29
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v204
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v208 != 0 {
		v365 = v9
		v372 = v9
		v380 = int32(0)
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v390 = int32(0)
	goto L81
L52:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v210 = F_copyObjectImpl(m, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v210 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	v214 = v212
	goto L56
L55:
	;
	v214 = int32(0)
	goto L56
L56:
	;
	if v204 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v299 = int32(0)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v300 == v299 {
		goto L67
	} else {
		goto L68
	}
L58:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v217 <= int32(0) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v229 = v214
	v231 = int32(0)
	goto L60
L60:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v253 = int32(2)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v252+v231<<(uint(v253)%32))))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v251)+4)) = v257
	v260 = v229 + int32(4)
	if base.Ui32(v260) < base.Ui32(v249+v250<<(uint(v253)%32)) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L57
L62:
	;
	v266 = v260
	goto L64
L63:
	;
	v266 = int32(0)
	goto L64
L64:
	;
	v268 = v231 + int32(1)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v268 < v269 {
		v229 = v266
		v231 = v268
		goto L60
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	if v345 == int32(0) {
		v365 = v210
		v372 = v9
		v380 = v299
		goto L51
	} else {
		goto L79
	}
L67:
	;
	v345 = int32(1)
	goto L66
L68:
	;
	goto L69
L69:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	if v309 <= int32(0) {
		v337 = int32(1)
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v345 = v337
	goto L66
L71:
	;
	v312 = int32(0)
	if v312 < v309 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v315 = v309
	goto L74
L73:
	;
	v315 = v312
	goto L74
L74:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v300)+12))
	v318 = int32(0)
	goto L75
L75:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v316+v318<<(uint(int32(2))%32))))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+12))
	v328 = int32(0)
	v329 = base.B2i32(v327 != v328)
	if v327 == v328 {
		v337 = v329
		goto L70
	} else {
		goto L77
	}
L76:
	;
	v337 = v329
	goto L70
L77:
	;
	v333 = v318 + int32(1)
	if v333 != v315 {
		v318 = v333
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v348 = F_make_pathkeys_for_sortclauses(m, l1, v210, v204)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+156)) = v348
	v365 = v210
	v372 = v348
	v380 = int32(1)
	goto L51
L81:
	;
	v410 = int32(0)
	if v176 == v410 {
		v420 = v410
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v471 = int32(0)
	v482 = v471
	v483 = v447
	v485 = v448
	v487 = v471
	v489 = v380
	v494 = int32(1)
	v495 = v9
	v496 = v9
	goto L106
L83:
	;
	v421 = int32(0)
	if v188 == v421 {
		v430 = v421
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	if v414 <= v390 {
		v420 = int32(0)
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
	v420 = v416 + v390<<(uint(int32(2))%32)
	goto L83
L86:
	;
	if v186 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v424 <= v390 {
		v430 = v421
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	v430 = v426 + v390<<(uint(int32(2))%32)
	goto L86
L89:
	;
	goto L82
L90:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v420)))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)+76))
	if v458 == int32(1) {
		goto L102
	} else {
		goto L103
	}
L91:
	;
	if v176 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L92:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	if v433 <= v390 {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	if v420 == int32(0) {
		goto L91
	} else {
		goto L94
	}
L94:
	;
	if v430 == int32(0) {
		goto L91
	} else {
		goto L95
	}
L95:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
	v442 = v439 + v390<<(uint(int32(2))%32)
	if v442 != 0 {
		goto L90
	} else {
		goto L96
	}
L96:
	;
	goto L91
L97:
	;
	v455 = int32(0)
	v1447 = v455
	v1450 = v453
	v1452 = v455
	v1454 = v380
	v1459 = int32(1)
	v1460 = v9
	v1461 = v9
	goto L5
L98:
	;
	v453 = int32(1)
	goto L97
L99:
	;
	goto L100
L100:
	;
	v447 = int32(0)
	v448 = int32(1)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	if v447 < v449 {
		goto L89
	} else {
		goto L101
	}
L101:
	;
	v453 = v448
	goto L97
L102:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v430)))
	v462 = int32(0)
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	F_build_setop_child_paths(m, l1, v457, base.B2i32(v461 != v462), v464, v372, v462)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v390 = v390 + int32(1)
	goto L81
L105:
	;
	goto L104
L106:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v501+v483<<(uint(int32(2))%32))))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v505)+48))
	v507 = F_lappend(m, v487, v506)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L108
	}
L107:
	;
	v1447 = v674
	v1450 = v669
	v1452 = v507
	v1454 = v647
	v1459 = v671
	v1460 = v672
	v1461 = v649
	goto L5
L108:
	;
	if v489&int32(1) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	if v485&int32(1) == int32(0) {
		goto L159
	} else {
		goto L160
	}
L110:
	;
	v647 = int32(0)
	v649 = v496
	goto L109
L111:
	;
	goto L112
L112:
	;
	v514 = int32(0)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v505)+32))
	if v515 != 0 {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	if v636 == int32(0) {
		v647 = v514
		v649 = v496
		goto L109
	} else {
		goto L156
	}
L114:
	;
	goto L113
L115:
	;
	v535 = v514
	v538 = v514
	goto L120
L116:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v515)+4))
	if int32(0) < v526 {
		goto L115
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v636 = v514
	goto L114
L119:
	;
	goto L118
L120:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v515)+12))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v541+v538<<(uint(int32(2))%32))))
	goto L124
L121:
	;
	v636 = v620
	goto L114
L122:
	;
	v627 = v538 + int32(1)
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v515)+4))
	if v627 < v628 {
		v535 = v620
		v538 = v627
		goto L120
	} else {
		goto L155
	}
L124:
	;
	goto L125
L125:
	;
	if v535 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v549 = F_compare_path_costs(m, v535, v545, int32(1))
	mBase = m.M
	if v549 <= int32(0) {
		v620 = v535
		goto L122
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v545)+64))
	if v372 == v552 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	goto L129
L131:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v545)+16))
	if v608 != 0 {
		goto L149
	} else {
		goto L150
	}
L132:
	;
	v560 = int32(0)
	goto L133
L133:
	;
	v567 = int32(0)
	if v372 == v567 {
		v577 = v567
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v577 != 0 {
		v620 = v535
		goto L122
	} else {
		goto L148
	}
L135:
	;
	if v552 != 0 {
		goto L139
	} else {
		goto L140
	}
L136:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	if v571 <= v560 {
		v577 = int32(0)
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v372)+12))
	v577 = v573 + v560<<(uint(int32(2))%32)
	goto L135
L138:
	;
	if v577 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L139:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v552)+4))
	if v560 < v578 {
		goto L138
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	if v577 == int32(0) {
		goto L131
	} else {
		goto L143
	}
L142:
	;
	goto L141
L143:
	;
	v620 = v535
	goto L122
L144:
	;
	goto L134
L145:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v552)+12))
	v587 = v584 + v560<<(uint(int32(2))%32)
	if v587 == int32(0) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v577)))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	if v592 == v593 {
		v560 = v560 + int32(1)
		goto L133
	} else {
		goto L147
	}
L147:
	;
	v620 = v535
	goto L122
L148:
	;
	goto L131
L149:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v608)+4))
	v611 = v609
	goto L151
L150:
	;
	v611 = int32(0)
	goto L151
L151:
	;
	v612 = F_bms_is_subset(m, v611, v514)
	mBase = m.M
	if v612 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v613 = v545
	goto L154
L153:
	;
	v613 = v535
	goto L154
L154:
	;
	v620 = v613
	goto L122
L155:
	;
	goto L121
L156:
	;
	v645 = F_lappend(m, v496, v636)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v647 = int32(1)
	v649 = v645
	goto L109
L158:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v505)+8))
	v674 = F_bms_union(m, v482, v673)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L1
	} else {
		goto L169
	}
L159:
	;
	v669 = int32(0)
	v671 = v494
	v672 = v495
	goto L158
L160:
	;
	goto L161
L161:
	;
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505)+26)))
	if v655 != int32(1) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v658 = int32(0)
	v669 = v658
	v671 = v658
	v672 = v495
	goto L158
L163:
	;
	goto L164
L164:
	;
	v660 = int32(1)
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v505)+40))
	if v661 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v669 = v660
	v671 = int32(0)
	v672 = v495
	goto L158
L166:
	;
	goto L167
L167:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v661)+12))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v665)))
	v667 = F_lappend(m, v495, v666)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	v669 = v660
	v671 = v494
	v672 = v667
	goto L158
L169:
	;
	v677 = v483 + int32(1)
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	if v677 < v678 {
		v482 = v674
		v483 = v677
		v485 = v669
		v487 = v507
		v489 = v647
		v494 = v671
		v495 = v672
		v496 = v649
		goto L106
	} else {
		goto L170
	}
L170:
	;
	goto L107
L171:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v693 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v699 = F_recurse_set_operations(m, v692, l1, l0, v693, v694, l5, v31+int32(72), v31+int32(70))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v701 = int32(0)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v31)+76))
	v709 = F_generate_setop_tlist(m, v702, v703, v701, v701, v706, l5, v31+int32(69))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v709
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v713 = F_copyObjectImpl(m, v712)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	if v713 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v713)+12))
	v716 = v715
	goto L177
L176:
	;
	v716 = v701
	goto L177
L177:
	;
	if v709 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	if v713 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L179:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v709)+4))
	if v719 <= int32(0) {
		goto L178
	} else {
		goto L180
	}
L180:
	;
	v725 = v716
	v731 = int32(0)
	goto L181
L181:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v713)+12))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v713)+4))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v725)))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v709)+12))
	v755 = int32(2)
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v754+v731<<(uint(v755)%32))))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v758)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v753)+4)) = v759
	v762 = v725 + int32(4)
	if base.Ui32(v762) < base.Ui32(v751+v752<<(uint(v755)%32)) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	goto L178
L183:
	;
	v768 = v762
	goto L185
L184:
	;
	v768 = int32(0)
	goto L185
L185:
	;
	v770 = v731 + int32(1)
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v709)+4))
	if v770 < v771 {
		v725 = v768
		v731 = v770
		goto L181
	} else {
		goto L186
	}
L186:
	;
	goto L182
L187:
	;
	v846 = int32(0)
	if v713 == v846 {
		goto L201
	} else {
		goto L202
	}
L188:
	;
	v845 = int32(1)
	goto L187
L189:
	;
	goto L190
L190:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v713)+4))
	if v809 <= int32(0) {
		v837 = int32(1)
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v845 = v837
	goto L187
L192:
	;
	v812 = int32(0)
	if v812 < v809 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v815 = v809
	goto L195
L194:
	;
	v815 = v812
	goto L195
L195:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v713)+12))
	v818 = int32(0)
	goto L196
L196:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v816+v818<<(uint(int32(2))%32))))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v826)+12))
	v828 = int32(0)
	v829 = base.B2i32(v827 != v828)
	if v827 == v828 {
		v837 = v829
		goto L191
	} else {
		goto L198
	}
L197:
	;
	v837 = v829
	goto L191
L198:
	;
	v833 = v818 + int32(1)
	if v833 != v815 {
		v818 = v833
		goto L196
	} else {
		goto L199
	}
L199:
	;
	goto L197
L200:
	;
	if v845|v883 == int32(0) {
		goto L9
	} else {
		goto L213
	}
L201:
	;
	v883 = int32(1)
	goto L200
L202:
	;
	goto L203
L203:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v713)+4))
	if v853 <= int32(0) {
		v878 = int32(1)
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v883 = v878
	goto L200
L205:
	;
	v856 = int32(0)
	if v856 < v853 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v859 = v853
	goto L208
L207:
	;
	v859 = v856
	goto L208
L208:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v713)+12))
	v863 = v846
	goto L209
L209:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v860+v863<<(uint(int32(2))%32))))
	v869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v868)+18)))
	if v869 != int32(1) {
		v878 = v869
		goto L204
	} else {
		goto L211
	}
L210:
	;
	v878 = v869
	goto L204
L211:
	;
	v873 = v863 + int32(1)
	if v873 != v859 {
		v863 = v873
		goto L209
	} else {
		goto L212
	}
L212:
	;
	goto L210
L213:
	;
	if v845 != 0 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v888 = F_make_pathkeys_for_sortclauses(m, l1, v713, v709)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L1
	} else {
		goto L217
	}
L215:
	;
	v891 = int32(0)
	goto L216
L216:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v690)+76))
	if v892 == int32(1) {
		goto L219
	} else {
		goto L220
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+156)) = v888
	v891 = v888
	goto L216
L218:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v699)+76))
	if v902 == int32(1) {
		goto L224
	} else {
		goto L225
	}
L219:
	;
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+71)))
	F_build_setop_child_paths(m, l1, v690, v895, v706, v891, v31+int32(56))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L1
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	v900 = *(*float64)(unsafe.Add(mBase, uint32(v690)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v31)+56)) = v900
	goto L218
L222:
	;
	goto L218
L223:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+296)) = v680
	v914 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v914 == int32(3) {
		goto L7
	} else {
		goto L228
	}
L224:
	;
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+70)))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
	F_build_setop_child_paths(m, l1, v699, v905, v906, v891, v31+int32(48))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L1
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v911 = *(*float64)(unsafe.Add(mBase, uint32(v699)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v31)+48)) = v911
	goto L223
L227:
	;
	goto L223
L228:
	;
	v917 = *(*float64)(unsafe.Add(mBase, uint32(v31)+56))
	v918 = *(*float64)(unsafe.Add(mBase, uint32(v31)+48))
	if base.F64_gt(v917, v918) == int32(0) {
		goto L7
	} else {
		goto L229
	}
L229:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v31)+56)) = v918
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+72)) = v706
	*(*float64)(unsafe.Add(mBase, uint32(v31)+48)) = v917
	v983 = v690
	v984 = v699
	v985 = v923
	goto L6
L230:
	;
	F_errmsg_internal(m, int32(15682), int32(0))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	F_errfinish(m, int32(522301), int32(254), int32(152243))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L233:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v948 == int32(2) {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v951 = int32(548894)
	goto L237
L236:
	;
	v951 = int32(545236)
	goto L237
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = v951
	F_errmsg(m, int32(189979), v31+int32(32))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	F_errdetail(m, int32(657505), int32(0))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	F_errfinish(m, int32(522301), int32(1073), int32(165308))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L241:
	;
	v971 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v971
	F_errmsg_internal(m, int32(510299), v31)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	F_errfinish(m, int32(522301), int32(345), int32(152243))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L244:
	;
	v995 = F_fetch_upper_rel(m, l1, int32(0), v993)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	v997 = F_make_pathtarget_from_tlist(m, v709)
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	v999 = F_set_pathtarget_cost_width(m, l1, v997)
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v995)+28)) = v999
	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v1003 = *(*float64)(unsafe.Add(mBase, uint32(v31)+56))
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1004 == int32(3) {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v995)+16)) = v1022
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v1024 - int32(2) {
	case 0:
		goto L258
	case 1:
		goto L260
	default:
		goto L259
	}
L249:
	;
	if v1002&int32(1) == int32(0) {
		v1022 = v1003
		goto L248
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	if v1002&int32(1) == int32(0) {
		v1022 = v1003
		goto L248
	} else {
		goto L253
	}
L252:
	;
	v1011 = *(*float64)(unsafe.Add(mBase, uint32(v989)+32))
	v1022 = v1011
	goto L248
L253:
	;
	v1016 = *(*float64)(unsafe.Add(mBase, uint32(v989)+32))
	v1017 = *(*float64)(unsafe.Add(mBase, uint32(v988)+32))
	if base.F64_lt(v1016, v1017) != 0 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1019 = v1016
	goto L256
L255:
	;
	v1019 = v1017
	goto L256
L256:
	;
	v1022 = v1019
	goto L248
L257:
	;
	if v883 != 0 {
		goto L267
	} else {
		goto L268
	}
L258:
	;
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v1048 = v1047
	goto L257
L259:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L1
	} else {
		goto L264
	}
L260:
	;
	v1029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v1029 != 0 {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v1030 = int32(3)
	goto L263
L262:
	;
	v1030 = int32(2)
	goto L263
L263:
	;
	v1048 = v1030
	goto L257
L264:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v1035
	F_errmsg_internal(m, int32(507058), v31+int32(16))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	F_errfinish(m, int32(522301), int32(1165), int32(165308))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L267:
	;
	v1050 = F_create_setop_path(m, v995, v989, v988, v1048, int32(1), v713, v1003, v1022)
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L1
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	if v845 == int32(0) {
		v1987 = v995
		goto L4
	} else {
		goto L272
	}
L270:
	;
	F_add_path(m, v995, v1050)
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	goto L269
L272:
	;
	v1056 = F_make_pathkeys_for_sortclauses(m, l1, v713, v985)
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L1
	} else {
		goto L274
	}
L273:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
	v1245 = F_make_pathkeys_for_sortclauses(m, l1, v713, v1244)
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L1
	} else {
		goto L342
	}
L274:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v989)+64))
	if v1056 == v1058 {
		goto L276
	} else {
		goto L277
	}
L275:
	;
	if v1111 != 0 {
		goto L293
	} else {
		goto L294
	}
L276:
	;
	v1111 = int32(1)
	goto L275
L277:
	;
	goto L278
L278:
	;
	v1067 = int32(0)
	goto L280
L279:
	;
	v1111 = v1103
	goto L275
L280:
	;
	v1071 = int32(0)
	if v1056 == v1071 {
		v1081 = v1071
		goto L282
	} else {
		goto L283
	}
L281:
	;
	v1103 = int32(0)
	goto L279
L282:
	;
	if v1058 != 0 {
		goto L286
	} else {
		goto L287
	}
L283:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+4))
	if v1075 <= v1067 {
		v1081 = int32(0)
		goto L282
	} else {
		goto L284
	}
L284:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+12))
	v1081 = v1077 + v1067<<(uint(int32(2))%32)
	goto L282
L285:
	;
	v1087 = base.B2i32(v1081 == int32(0))
	if v1081 == int32(0) {
		v1103 = v1087
		goto L279
	} else {
		goto L290
	}
L286:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1058)+4))
	if v1067 < v1082 {
		goto L285
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	v1111 = base.B2i32(v1081 == int32(0))
	goto L275
L289:
	;
	goto L288
L290:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1058)+12))
	v1093 = v1090 + v1067<<(uint(int32(2))%32)
	if v1093 == int32(0) {
		v1103 = v1087
		goto L279
	} else {
		goto L291
	}
L291:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1081)))
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1093)))
	if v1098 == v1099 {
		v1067 = v1067 + int32(1)
		goto L280
	} else {
		goto L292
	}
L292:
	;
	goto L281
L293:
	;
	v1243 = v989
	goto L273
L294:
	;
	goto L295
L295:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v984)+32))
	v1113 = int32(0)
	if v1112 != 0 {
		goto L299
	} else {
		goto L300
	}
L296:
	;
	if v1233 != 0 {
		v1243 = v1233
		goto L273
	} else {
		goto L339
	}
L297:
	;
	goto L296
L298:
	;
	v1132 = v1113
	v1135 = v1113
	goto L303
L299:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+4))
	if int32(0) < v1123 {
		goto L298
	} else {
		goto L302
	}
L300:
	;
	goto L301
L301:
	;
	v1233 = v1113
	goto L297
L302:
	;
	goto L301
L303:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+12))
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1138+v1135<<(uint(int32(2))%32))))
	goto L307
L304:
	;
	v1233 = v1217
	goto L297
L305:
	;
	v1224 = v1135 + int32(1)
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+4))
	if v1224 < v1225 {
		v1132 = v1217
		v1135 = v1224
		goto L303
	} else {
		goto L338
	}
L307:
	;
	goto L308
L308:
	;
	if v1132 != 0 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1146 = F_compare_path_costs(m, v1132, v1142, int32(1))
	mBase = m.M
	if v1146 <= int32(0) {
		v1217 = v1132
		goto L305
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+64))
	if v891 == v1149 {
		goto L314
	} else {
		goto L315
	}
L313:
	;
	goto L312
L314:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+16))
	if v1205 != 0 {
		goto L332
	} else {
		goto L333
	}
L315:
	;
	v1157 = int32(0)
	goto L316
L316:
	;
	v1164 = int32(0)
	if v891 == v1164 {
		v1174 = v1164
		goto L318
	} else {
		goto L319
	}
L317:
	;
	if v1174 != 0 {
		v1217 = v1132
		goto L305
	} else {
		goto L331
	}
L318:
	;
	if v1149 != 0 {
		goto L322
	} else {
		goto L323
	}
L319:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v891)+4))
	if v1168 <= v1157 {
		v1174 = int32(0)
		goto L318
	} else {
		goto L320
	}
L320:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v891)+12))
	v1174 = v1170 + v1157<<(uint(int32(2))%32)
	goto L318
L321:
	;
	if v1174 == int32(0) {
		goto L327
	} else {
		goto L328
	}
L322:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1149)+4))
	if v1157 < v1175 {
		goto L321
	} else {
		goto L325
	}
L323:
	;
	goto L324
L324:
	;
	if v1174 == int32(0) {
		goto L314
	} else {
		goto L326
	}
L325:
	;
	goto L324
L326:
	;
	v1217 = v1132
	goto L305
L327:
	;
	goto L317
L328:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1149)+12))
	v1184 = v1181 + v1157<<(uint(int32(2))%32)
	if v1184 == int32(0) {
		goto L327
	} else {
		goto L329
	}
L329:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1174)))
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1184)))
	if v1189 == v1190 {
		v1157 = v1157 + int32(1)
		goto L316
	} else {
		goto L330
	}
L330:
	;
	v1217 = v1132
	goto L305
L331:
	;
	goto L314
L332:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+4))
	v1208 = v1206
	goto L334
L333:
	;
	v1208 = int32(0)
	goto L334
L334:
	;
	v1209 = F_bms_is_subset(m, v1208, v1113)
	mBase = m.M
	if v1209 != 0 {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v1210 = v1142
	goto L337
L336:
	;
	v1210 = v1132
	goto L337
L337:
	;
	v1217 = v1210
	goto L305
L338:
	;
	goto L304
L339:
	;
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v989)+8))
	v1241 = F_create_sort_path(m, v1239, v989, v1056, float64(-1))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L1
	} else {
		goto L340
	}
L340:
	;
	v1243 = v1241
	goto L273
L341:
	;
	v1434 = F_create_setop_path(m, v995, v1243, v1432, v1048, int32(0), v713, v1003, v1022)
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L1
	} else {
		goto L409
	}
L342:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v988)+64))
	if v1245 == v1247 {
		goto L344
	} else {
		goto L345
	}
L343:
	;
	if v1300 != 0 {
		goto L361
	} else {
		goto L362
	}
L344:
	;
	v1300 = int32(1)
	goto L343
L345:
	;
	goto L346
L346:
	;
	v1256 = int32(0)
	goto L348
L347:
	;
	v1300 = v1292
	goto L343
L348:
	;
	v1260 = int32(0)
	if v1245 == v1260 {
		v1270 = v1260
		goto L350
	} else {
		goto L351
	}
L349:
	;
	v1292 = int32(0)
	goto L347
L350:
	;
	if v1247 != 0 {
		goto L354
	} else {
		goto L355
	}
L351:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+4))
	if v1264 <= v1256 {
		v1270 = int32(0)
		goto L350
	} else {
		goto L352
	}
L352:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+12))
	v1270 = v1266 + v1256<<(uint(int32(2))%32)
	goto L350
L353:
	;
	v1276 = base.B2i32(v1270 == int32(0))
	if v1270 == int32(0) {
		v1292 = v1276
		goto L347
	} else {
		goto L358
	}
L354:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1247)+4))
	if v1256 < v1271 {
		goto L353
	} else {
		goto L357
	}
L355:
	;
	goto L356
L356:
	;
	v1300 = base.B2i32(v1270 == int32(0))
	goto L343
L357:
	;
	goto L356
L358:
	;
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v1247)+12))
	v1282 = v1279 + v1256<<(uint(int32(2))%32)
	if v1282 == int32(0) {
		v1292 = v1276
		goto L347
	} else {
		goto L359
	}
L359:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1270)))
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1282)))
	if v1287 == v1288 {
		v1256 = v1256 + int32(1)
		goto L348
	} else {
		goto L360
	}
L360:
	;
	goto L349
L361:
	;
	v1432 = v988
	goto L341
L362:
	;
	goto L363
L363:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v983)+32))
	v1302 = int32(0)
	if v1301 != 0 {
		goto L367
	} else {
		goto L368
	}
L364:
	;
	if v1422 != 0 {
		v1432 = v1422
		goto L341
	} else {
		goto L407
	}
L365:
	;
	goto L364
L366:
	;
	v1321 = v1302
	v1324 = v1302
	goto L371
L367:
	;
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1301)+4))
	if int32(0) < v1312 {
		goto L366
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	v1422 = v1302
	goto L365
L370:
	;
	goto L369
L371:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v1301)+12))
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1327+v1324<<(uint(int32(2))%32))))
	goto L375
L372:
	;
	v1422 = v1406
	goto L365
L373:
	;
	v1413 = v1324 + int32(1)
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v1301)+4))
	if v1413 < v1414 {
		v1321 = v1406
		v1324 = v1413
		goto L371
	} else {
		goto L406
	}
L375:
	;
	goto L376
L376:
	;
	if v1321 != 0 {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	v1335 = F_compare_path_costs(m, v1321, v1331, int32(1))
	mBase = m.M
	if v1335 <= int32(0) {
		v1406 = v1321
		goto L373
	} else {
		goto L381
	}
L379:
	;
	goto L380
L380:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+64))
	if v891 == v1338 {
		goto L382
	} else {
		goto L383
	}
L381:
	;
	goto L380
L382:
	;
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+16))
	if v1394 != 0 {
		goto L400
	} else {
		goto L401
	}
L383:
	;
	v1346 = int32(0)
	goto L384
L384:
	;
	v1353 = int32(0)
	if v891 == v1353 {
		v1363 = v1353
		goto L386
	} else {
		goto L387
	}
L385:
	;
	if v1363 != 0 {
		v1406 = v1321
		goto L373
	} else {
		goto L399
	}
L386:
	;
	if v1338 != 0 {
		goto L390
	} else {
		goto L391
	}
L387:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v891)+4))
	if v1357 <= v1346 {
		v1363 = int32(0)
		goto L386
	} else {
		goto L388
	}
L388:
	;
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v891)+12))
	v1363 = v1359 + v1346<<(uint(int32(2))%32)
	goto L386
L389:
	;
	if v1363 == int32(0) {
		goto L395
	} else {
		goto L396
	}
L390:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+4))
	if v1346 < v1364 {
		goto L389
	} else {
		goto L393
	}
L391:
	;
	goto L392
L392:
	;
	if v1363 == int32(0) {
		goto L382
	} else {
		goto L394
	}
L393:
	;
	goto L392
L394:
	;
	v1406 = v1321
	goto L373
L395:
	;
	goto L385
L396:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+12))
	v1373 = v1370 + v1346<<(uint(int32(2))%32)
	if v1373 == int32(0) {
		goto L395
	} else {
		goto L397
	}
L397:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1363)))
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v1373)))
	if v1378 == v1379 {
		v1346 = v1346 + int32(1)
		goto L384
	} else {
		goto L398
	}
L398:
	;
	v1406 = v1321
	goto L373
L399:
	;
	goto L382
L400:
	;
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1394)+4))
	v1397 = v1395
	goto L402
L401:
	;
	v1397 = int32(0)
	goto L402
L402:
	;
	v1398 = F_bms_is_subset(m, v1397, v1302)
	mBase = m.M
	if v1398 != 0 {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	v1399 = v1331
	goto L405
L404:
	;
	v1399 = v1321
	goto L405
L405:
	;
	v1406 = v1399
	goto L373
L406:
	;
	goto L372
L407:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v988)+8))
	v1430 = F_create_sort_path(m, v1428, v988, v1245, float64(-1))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L1
	} else {
		goto L408
	}
L408:
	;
	v1432 = v1430
	goto L341
L409:
	;
	F_add_path(m, v995, v1434)
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L1
	} else {
		goto L410
	}
L410:
	;
	v1987 = v995
	goto L4
L411:
	;
	v1469 = F_make_pathtarget_from_tlist(m, v204)
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	v1471 = F_set_pathtarget_cost_width(m, l1, v1469)
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		goto L1
	} else {
		goto L413
	}
L413:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1467)+26)) = uint8(v1450)
	*(*int32)(unsafe.Add(mBase, uint32(v1467)+28)) = v1471
	v1475 = *(*float64)(unsafe.Add(mBase, uint32(l1)+296))
	v1477 = base.F64_gt(v1475, float64(0))
	*(*uint8)(unsafe.Add(mBase, uint32(v1467)+24)) = uint8(v1477)
	v1479 = int32(0)
	v1485 = F_create_append_path(m, l1, v1467, v1452, v1479, v1479, v1479, v1479, v1479, float64(-1))
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L1
	} else {
		goto L414
	}
L414:
	;
	v1487 = *(*float64)(unsafe.Add(mBase, uint32(v1485)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1467)+16)) = v1487
	if v1459&int32(1) != 0 {
		goto L422
	} else {
		goto L423
	}
L415:
	;
	F_add_path(m, v1467, v1485)
	mBase = m.M
	v1973 = m.ExcPending
	if v1973 != 0 {
		goto L1
	} else {
		goto L533
	}
L416:
	;
	if v365 == int32(0) {
		goto L470
	} else {
		goto L471
	}
L417:
	;
	v1773 = int32(0)
	v1777 = F_create_append_path(m, l1, v1467, v1773, v1460, v1773, v1773, v1747, v1764, float64(-1))
	mBase = m.M
	v1778 = m.ExcPending
	if v1778 != 0 {
		goto L1
	} else {
		goto L466
	}
L418:
	;
	v1741 = *(*int32)(unsafe.Add(mBase, _consts[612]))
	if v1714 < v1741 {
		goto L463
	} else {
		goto L464
	}
L419:
	;
	v1714 = int32(32) - base.I32_clz(v1700)
	goto L418
L420:
	;
	v1673 = int32(*(*uint8)(unsafe.Add(mBase, _consts[613])))
	if v1673 == int32(0) {
		goto L459
	} else {
		goto L460
	}
L421:
	;
	if v1500 == int32(0) {
		v1646 = v1574
		goto L420
	} else {
		goto L452
	}
L422:
	;
	if v1460 != 0 {
		goto L426
	} else {
		goto L427
	}
L423:
	;
	goto L424
L424:
	;
	v1567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v1567 != int32(1) {
		v1786 = int32(0)
		goto L416
	} else {
		goto L450
	}
L425:
	;
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v1460)+12))
	v1511 = int32(0)
	v1516 = v1511
	v1522 = v1511
	v1526 = v1511
	goto L435
L426:
	;
	v1491 = int32(0)
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v1460)+4))
	if v1492 <= v1491 {
		v1646 = v1491
		goto L420
	} else {
		goto L429
	}
L427:
	;
	goto L428
L428:
	;
	v1507 = int32(*(*uint8)(unsafe.Add(mBase, _consts[613])))
	if v1507 != 0 {
		v1700 = v9
		goto L419
	} else {
		goto L434
	}
L429:
	;
	v1495 = int32(0)
	if v1495 < v1492 {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v1498 = v1492
	goto L432
L431:
	;
	v1498 = v1495
	goto L432
L432:
	;
	v1500 = v1498 & int32(3)
	if int32(4) <= v1492 {
		goto L425
	} else {
		goto L433
	}
L433:
	;
	v1503 = int32(0)
	v1574 = v1503
	v1580 = v1503
	goto L421
L434:
	;
	v1747 = int32(0)
	v1764 = v9
	goto L417
L435:
	;
	v1544 = v1510 + v1522<<(uint(int32(2))%32)
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v1544)))
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1545)+24))
	if v1546 < v1516 {
		goto L437
	} else {
		goto L438
	}
L436:
	;
	v1574 = v1560
	v1580 = v1562
	goto L421
L437:
	;
	v1548 = v1516
	goto L439
L438:
	;
	v1548 = v1546
	goto L439
L439:
	;
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(v1544)+4))
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v1549)+24))
	if v1550 < v1548 {
		goto L440
	} else {
		goto L441
	}
L440:
	;
	v1552 = v1548
	goto L442
L441:
	;
	v1552 = v1550
	goto L442
L442:
	;
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v1544)+8))
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(v1553)+24))
	if v1554 < v1552 {
		goto L443
	} else {
		goto L444
	}
L443:
	;
	v1556 = v1552
	goto L445
L444:
	;
	v1556 = v1554
	goto L445
L445:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v1544)+12))
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(v1557)+24))
	if v1558 < v1556 {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	v1560 = v1556
	goto L448
L447:
	;
	v1560 = v1558
	goto L448
L448:
	;
	v1561 = int32(4)
	v1562 = v1522 + v1561
	v1564 = v1526 + v1561
	if v1564 != v1498&int32(2147483644) {
		v1516 = v1560
		v1522 = v1562
		v1526 = v1564
		goto L435
	} else {
		goto L449
	}
L449:
	;
	goto L436
L450:
	;
	F_add_path(m, v1467, v1485)
	mBase = m.M
	v1571 = m.ExcPending
	if v1571 != 0 {
		goto L1
	} else {
		goto L451
	}
L451:
	;
	v1987 = v1467
	goto L4
L452:
	;
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(v1460)+12))
	v1606 = v1574
	v1612 = v1580
	v1614 = int32(0)
	goto L453
L453:
	;
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1602+v1612<<(uint(int32(2))%32))))
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v1635)+24))
	if v1636 < v1606 {
		goto L455
	} else {
		goto L456
	}
L454:
	;
	v1646 = v1638
	goto L420
L455:
	;
	v1638 = v1606
	goto L457
L456:
	;
	v1638 = v1636
	goto L457
L457:
	;
	v1639 = int32(1)
	v1642 = v1614 + v1639
	if v1642 != v1500 {
		v1606 = v1638
		v1612 = v1612 + v1639
		v1614 = v1642
		goto L453
	} else {
		goto L458
	}
L458:
	;
	goto L454
L459:
	;
	v1747 = v1646
	v1764 = int32(0)
	goto L417
L460:
	;
	goto L461
L461:
	;
	if int32(32)-base.I32_clz(v1492) < v1646 {
		v1714 = v1646
		goto L418
	} else {
		goto L462
	}
L462:
	;
	v1700 = v1492
	goto L419
L463:
	;
	v1743 = v1714
	goto L465
L464:
	;
	v1743 = v1741
	goto L465
L465:
	;
	v1747 = v1743
	v1764 = int32(1)
	goto L417
L466:
	;
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(v1467)+28))
	v1781 = F_create_gather_path(m, l1, v1467, v1777, v1779, int32(0))
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L1
	} else {
		goto L467
	}
L467:
	;
	v1783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v1783 != 0 {
		goto L415
	} else {
		goto L468
	}
L468:
	;
	v1786 = v1781
	goto L416
L469:
	;
	v1857 = int32(0)
	if v365 == v1857 {
		goto L483
	} else {
		goto L484
	}
L470:
	;
	v1856 = int32(1)
	goto L469
L471:
	;
	goto L472
L472:
	;
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	if v1820 <= int32(0) {
		v1848 = int32(1)
		goto L473
	} else {
		goto L474
	}
L473:
	;
	v1856 = v1848
	goto L469
L474:
	;
	v1823 = int32(0)
	if v1823 < v1820 {
		goto L475
	} else {
		goto L476
	}
L475:
	;
	v1826 = v1820
	goto L477
L476:
	;
	v1826 = v1823
	goto L477
L477:
	;
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v365)+12))
	v1829 = int32(0)
	goto L478
L478:
	;
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v1827+v1829<<(uint(int32(2))%32))))
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v1837)+12))
	v1839 = int32(0)
	v1840 = base.B2i32(v1838 != v1839)
	if v1838 == v1839 {
		v1848 = v1840
		goto L473
	} else {
		goto L480
	}
L479:
	;
	v1848 = v1840
	goto L473
L480:
	;
	v1844 = v1829 + int32(1)
	if v1844 != v1826 {
		v1829 = v1844
		goto L478
	} else {
		goto L481
	}
L481:
	;
	goto L479
L482:
	;
	v1895 = *(*float64)(unsafe.Add(mBase, uint32(v1485)+32))
	if v1894 == int32(0) {
		goto L495
	} else {
		goto L496
	}
L483:
	;
	v1894 = int32(1)
	goto L482
L484:
	;
	goto L485
L485:
	;
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	if v1864 <= int32(0) {
		v1889 = int32(1)
		goto L486
	} else {
		goto L487
	}
L486:
	;
	v1894 = v1889
	goto L482
L487:
	;
	v1867 = int32(0)
	if v1867 < v1864 {
		goto L488
	} else {
		goto L489
	}
L488:
	;
	v1870 = v1864
	goto L490
L489:
	;
	v1870 = v1867
	goto L490
L490:
	;
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v365)+12))
	v1874 = v1857
	goto L491
L491:
	;
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v1871+v1874<<(uint(int32(2))%32))))
	v1880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1879)+18)))
	if v1880 != int32(1) {
		v1889 = v1880
		goto L486
	} else {
		goto L493
	}
L492:
	;
	v1889 = v1880
	goto L486
L493:
	;
	v1884 = v1874 + int32(1)
	if v1884 != v1870 {
		v1874 = v1884
		goto L491
	} else {
		goto L494
	}
L494:
	;
	goto L492
L495:
	;
	if v1856 == int32(0) {
		goto L506
	} else {
		goto L507
	}
L496:
	;
	v1898 = F_make_pathtarget_from_tlist(m, v204)
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	v1900 = F_set_pathtarget_cost_width(m, l1, v1898)
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
		goto L1
	} else {
		goto L498
	}
L498:
	;
	v1903 = int32(0)
	v1906 = F_create_agg_path(m, l1, v1467, v1485, v1900, int32(2), v1903, v365, v1903, v1903, v1895)
	mBase = m.M
	v1907 = m.ExcPending
	if v1907 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	F_add_path(m, v1467, v1906)
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L1
	} else {
		goto L500
	}
L500:
	;
	if v1786 == int32(0) {
		goto L495
	} else {
		goto L501
	}
L501:
	;
	v1912 = F_make_pathtarget_from_tlist(m, v204)
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		goto L1
	} else {
		goto L502
	}
L502:
	;
	v1914 = F_set_pathtarget_cost_width(m, l1, v1912)
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L1
	} else {
		goto L503
	}
L503:
	;
	v1917 = int32(0)
	v1920 = F_create_agg_path(m, l1, v1467, v1786, v1914, int32(2), v1917, v365, v1917, v1917, v1895)
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L1
	} else {
		goto L504
	}
L504:
	;
	F_add_path(m, v1467, v1920)
	mBase = m.M
	v1923 = m.ExcPending
	if v1923 != 0 {
		goto L1
	} else {
		goto L505
	}
L505:
	;
	goto L495
L506:
	;
	if base.B2i32(v365 == int32(0))|(v1454^int32(1)) != 0 {
		v1987 = v1467
		goto L4
	} else {
		goto L526
	}
L507:
	;
	if v365 != 0 {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v1926 = F_make_pathkeys_for_sortclauses(m, l1, v365, v204)
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L1
	} else {
		goto L511
	}
L509:
	;
	v1931 = v1485
	goto L510
L510:
	;
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v1931)+64))
	if v1932 != 0 {
		goto L513
	} else {
		goto L514
	}
L511:
	;
	v1929 = F_create_sort_path(m, v1467, v1485, v1926, float64(-1))
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	v1931 = v1929
	goto L510
L513:
	;
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(v1932)+4))
	v1935 = v1933
	goto L515
L514:
	;
	v1935 = int32(0)
	goto L515
L515:
	;
	v1936 = F_create_upper_unique_path(m, v1467, v1931, v1935, v1895)
	mBase = m.M
	v1937 = m.ExcPending
	if v1937 != 0 {
		goto L1
	} else {
		goto L516
	}
L516:
	;
	F_add_path(m, v1467, v1936)
	mBase = m.M
	v1939 = m.ExcPending
	if v1939 != 0 {
		goto L1
	} else {
		goto L517
	}
L517:
	;
	if v1786 == int32(0) {
		goto L506
	} else {
		goto L518
	}
L518:
	;
	v1942 = F_make_pathkeys_for_sortclauses(m, l1, v365, v204)
	mBase = m.M
	v1943 = m.ExcPending
	if v1943 != 0 {
		goto L1
	} else {
		goto L519
	}
L519:
	;
	v1945 = F_create_sort_path(m, v1467, v1786, v1942, float64(-1))
	mBase = m.M
	v1946 = m.ExcPending
	if v1946 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+64))
	if v1947 != 0 {
		goto L521
	} else {
		goto L522
	}
L521:
	;
	v1948 = *(*int32)(unsafe.Add(mBase, uint32(v1947)+4))
	v1950 = v1948
	goto L523
L522:
	;
	v1950 = int32(0)
	goto L523
L523:
	;
	v1951 = F_create_upper_unique_path(m, v1467, v1945, v1950, v1895)
	mBase = m.M
	v1952 = m.ExcPending
	if v1952 != 0 {
		goto L1
	} else {
		goto L524
	}
L524:
	;
	F_add_path(m, v1467, v1951)
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	goto L506
L526:
	;
	v1963 = F_create_merge_append_path(m, l1, v1467, v1461, v372)
	mBase = m.M
	v1964 = m.ExcPending
	if v1964 != 0 {
		goto L1
	} else {
		goto L527
	}
L527:
	;
	if v204 != 0 {
		goto L528
	} else {
		goto L529
	}
L528:
	;
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	v1967 = v1965
	goto L530
L529:
	;
	v1967 = int32(0)
	goto L530
L530:
	;
	v1968 = F_create_upper_unique_path(m, v1467, v1963, v1967, v1895)
	mBase = m.M
	v1969 = m.ExcPending
	if v1969 != 0 {
		goto L1
	} else {
		goto L531
	}
L531:
	;
	F_add_path(m, v1467, v1968)
	mBase = m.M
	v1971 = m.ExcPending
	if v1971 != 0 {
		goto L1
	} else {
		goto L532
	}
L532:
	;
	v1987 = v1467
	goto L4
L533:
	;
	if v1781 == int32(0) {
		v1987 = v1467
		goto L4
	} else {
		goto L534
	}
L534:
	;
	F_add_path(m, v1467, v1781)
	mBase = m.M
	v1977 = m.ExcPending
	if v1977 != 0 {
		goto L1
	} else {
		goto L535
	}
L535:
	;
	v1987 = v1467
	goto L4
L536:
	;
	v2334 = *(*int32)(unsafe.Add(mBase, _consts[614]))
	if v2334 != 0 {
		goto L582
	} else {
		goto L583
	}
L537:
	;
	if v2008 != 0 {
		goto L538
	} else {
		goto L539
	}
L538:
	;
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if l4 != 0 {
		goto L542
	} else {
		goto L543
	}
L539:
	;
	goto L540
L540:
	;
	v2165 = int32(0)
	v2167 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v2170 = F_generate_setop_tlist(m, l3, l4, v2165, v2165, v2167, l5, v31+int32(56))
	mBase = m.M
	v2171 = m.ExcPending
	if v2171 != 0 {
		goto L1
	} else {
		goto L563
	}
L541:
	;
	if v2136 != 0 {
		goto L536
	} else {
		goto L562
	}
L542:
	;
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v2014 = v2012
	goto L544
L543:
	;
	v2014 = int32(0)
	goto L544
L544:
	;
	if v2010 == int32(0) {
		v2091 = v2014
		goto L545
	} else {
		goto L546
	}
L545:
	;
	v2136 = base.B2i32(v2091 == int32(0))
	goto L541
L546:
	;
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(v2010)+4))
	if v2017 <= int32(0) {
		v2091 = v2014
		goto L545
	} else {
		goto L547
	}
L547:
	;
	v2022 = int32(0)
	v2033 = v2014
	goto L548
L548:
	;
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(v2010)+12))
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(v2048+v2022<<(uint(int32(2))%32))))
	v2053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2052)+26)))
	if v2053 == int32(0) {
		goto L551
	} else {
		goto L552
	}
L549:
	;
	v2091 = v2072
	goto L545
L550:
	;
	v2075 = v2022 + int32(1)
	v2076 = *(*int32)(unsafe.Add(mBase, uint32(v2010)+4))
	if v2075 < v2076 {
		v2022 = v2075
		v2033 = v2072
		goto L548
	} else {
		goto L561
	}
L551:
	;
	if v2033 == int32(0) {
		goto L554
	} else {
		goto L555
	}
L552:
	;
	goto L553
L553:
	;
	v2136 = int32(0)
	goto L541
L554:
	;
	goto L553
L555:
	;
	v2058 = *(*int32)(unsafe.Add(mBase, uint32(v2052)+4))
	v2059 = F_exprCollation(m, v2058)
	mBase = m.M
	v2060 = m.ExcPending
	if v2060 != 0 {
		goto L1
	} else {
		goto L556
	}
L556:
	;
	v2061 = *(*int32)(unsafe.Add(mBase, uint32(v2033)))
	if v2059 != v2061 {
		goto L554
	} else {
		goto L557
	}
L557:
	;
	v2064 = v2033 + int32(4)
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if base.Ui32(v2064) < base.Ui32(v2066+v2067<<(uint(int32(2))%32)) {
		goto L558
	} else {
		goto L559
	}
L558:
	;
	v2072 = v2064
	goto L560
L559:
	;
	v2072 = int32(0)
	goto L560
L560:
	;
	goto L550
L561:
	;
	goto L549
L562:
	;
	goto L540
L563:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v2170
	v2173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+56)))
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v2173)
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v2176 = F_make_pathtarget_from_tlist(m, v2175)
	mBase = m.M
	v2177 = m.ExcPending
	if v2177 != 0 {
		goto L1
	} else {
		goto L564
	}
L564:
	;
	v2178 = F_set_pathtarget_cost_width(m, l1, v2176)
	mBase = m.M
	v2179 = m.ExcPending
	if v2179 != 0 {
		goto L1
	} else {
		goto L565
	}
L565:
	;
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(v1987)+32))
	if v2180 == int32(0) {
		goto L566
	} else {
		goto L567
	}
L566:
	;
	v2257 = *(*int32)(unsafe.Add(mBase, uint32(v1987)+40))
	if v2257 == int32(0) {
		goto L536
	} else {
		goto L576
	}
L567:
	;
	v2183 = int32(0)
	v2184 = *(*int32)(unsafe.Add(mBase, uint32(v2180)+4))
	if v2184 <= v2183 {
		goto L566
	} else {
		goto L568
	}
L568:
	;
	v2189 = v2183
	goto L569
L569:
	;
	v2215 = *(*int32)(unsafe.Add(mBase, uint32(v2180)+12))
	v2218 = v2215 + v2189<<(uint(int32(2))%32)
	v2219 = *(*int32)(unsafe.Add(mBase, uint32(v2218)))
	v2220 = *(*int32)(unsafe.Add(mBase, uint32(v2219)+8))
	v2221 = F_apply_projection_to_path(m, l1, v2220, v2219, v2178)
	mBase = m.M
	v2222 = m.ExcPending
	if v2222 != 0 {
		goto L1
	} else {
		goto L571
	}
L570:
	;
	goto L566
L571:
	;
	if v2219 != v2221 {
		goto L572
	} else {
		goto L573
	}
L572:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2218))) = v2221
	goto L574
L573:
	;
	goto L574
L574:
	;
	v2226 = v2189 + int32(1)
	v2227 = *(*int32)(unsafe.Add(mBase, uint32(v2180)+4))
	if v2226 < v2227 {
		v2189 = v2226
		goto L569
	} else {
		goto L575
	}
L575:
	;
	goto L570
L576:
	;
	v2260 = int32(0)
	v2261 = *(*int32)(unsafe.Add(mBase, uint32(v2257)+4))
	if v2261 <= v2260 {
		goto L536
	} else {
		goto L577
	}
L577:
	;
	v2264 = v2260
	goto L578
L578:
	;
	v2292 = *(*int32)(unsafe.Add(mBase, uint32(v2257)+12))
	v2295 = v2292 + v2264<<(uint(int32(2))%32)
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v2295)))
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(v2296)+8))
	v2298 = F_create_projection_path(m, l1, v2297, v2296, v2178)
	mBase = m.M
	v2299 = m.ExcPending
	if v2299 != 0 {
		goto L1
	} else {
		goto L580
	}
L579:
	;
	goto L536
L580:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2295))) = v2298
	v2302 = v2264 + int32(1)
	v2303 = *(*int32)(unsafe.Add(mBase, uint32(v2257)+4))
	if v2302 < v2303 {
		v2264 = v2302
		goto L578
	} else {
		goto L581
	}
L581:
	;
	goto L579
L582:
	;
	v2335 = int32(0)
	m.T0[v2334].(func(*base.Module, int32, int32, int32, int32, int32))(m, l1, v2335, v2335, v1987, v2335)
	mBase = m.M
	v2339 = m.ExcPending
	if v2339 != 0 {
		goto L1
	} else {
		goto L585
	}
L583:
	;
	goto L584
L584:
	;
	F_set_cheapest(m, v1987)
	mBase = m.M
	v2341 = m.ExcPending
	if v2341 != 0 {
		goto L1
	} else {
		goto L586
	}
L585:
	;
	goto L584
L586:
	;
	v2351 = v1987
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
	var v98 int32
	_ = v98
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
		goto L31
	} else {
		goto L46
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v151
L3:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	if v119 == int32(0) {
		goto L1
	} else {
		goto L33
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
	v23 = int32(580318)
	v27 = m.G0
	v29 = v27 - int32(32)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v30
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _consts[266])))
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
		goto L30
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
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _consts[267])))
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
		v98 = v11
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v98 - v11
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
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v98 = v94
	goto L22
L26:
	;
	v98 = v77
	goto L22
L27:
	;
	goto L28
L28:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v94 = v77 + int32(1)
	if v92 != 0 {
		v77 = v94
		v78 = v92
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v113 = F_DirectInputFunctionCallSafe(m, int32(547), v11, int32(-1), v10, v8+int32(12))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(0)
L32:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v151 = v117
	goto L2
L33:
	;
	v122 = F_stringToQualifiedNameList(m, v11, v10)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	if v122 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v126)
	v151 = int32(0)
	goto L2
L36:
	;
	goto L37
L37:
	;
	v130 = F_get_ts_config_oid(m, v122, int32(1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L31
	} else {
		goto L38
	}
L38:
	;
	if v130 != 0 {
		v151 = v130
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v132 = int32(0)
	v133 = F_errsave_start(m, v10)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L31
	} else {
		goto L40
	}
L40:
	;
	if v133 == int32(0) {
		v151 = v132
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
		goto L31
	} else {
		goto L42
	}
L42:
	;
	v140 = F_NameListToString(m, v122)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L31
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v140
	F_errmsg(m, int32(77413), v8)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L31
	} else {
		goto L44
	}
L44:
	;
	F_errsave_finish(m, v10, int32(526468), int32(1350), int32(290897))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L31
	} else {
		goto L45
	}
L45:
	;
	v151 = v132
	goto L2
L46:
	;
	F_errmsg_internal(m, int32(433203), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L31
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(526468), int32(1334), int32(290897))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L31
	} else {
		goto L48
	}
L48:
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
	var v98 int32
	_ = v98
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
		goto L31
	} else {
		goto L46
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v151
L3:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	if v119 == int32(0) {
		goto L1
	} else {
		goto L33
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
	v23 = int32(580318)
	v27 = m.G0
	v29 = v27 - int32(32)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v30
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _consts[266])))
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
		goto L30
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
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _consts[267])))
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
		v98 = v11
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v98 - v11
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
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v98 = v94
	goto L22
L26:
	;
	v98 = v77
	goto L22
L27:
	;
	goto L28
L28:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v94 = v77 + int32(1)
	if v92 != 0 {
		v77 = v94
		v78 = v92
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v113 = F_DirectInputFunctionCallSafe(m, int32(547), v11, int32(-1), v10, v8+int32(12))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(0)
L32:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v151 = v117
	goto L2
L33:
	;
	v122 = F_stringToQualifiedNameList(m, v11, v10)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	if v122 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v126)
	v151 = int32(0)
	goto L2
L36:
	;
	goto L37
L37:
	;
	v130 = F_get_ts_dict_oid(m, v122, int32(1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L31
	} else {
		goto L38
	}
L38:
	;
	if v130 != 0 {
		v151 = v130
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v132 = int32(0)
	v133 = F_errsave_start(m, v10)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L31
	} else {
		goto L40
	}
L40:
	;
	if v133 == int32(0) {
		v151 = v132
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
		goto L31
	} else {
		goto L42
	}
L42:
	;
	v140 = F_NameListToString(m, v122)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L31
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v140
	F_errmsg(m, int32(76598), v8)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L31
	} else {
		goto L44
	}
L44:
	;
	F_errsave_finish(m, v10, int32(526468), int32(1460), int32(289151))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L31
	} else {
		goto L45
	}
L45:
	;
	v151 = v132
	goto L2
L46:
	;
	F_errmsg_internal(m, int32(432957), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L31
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(526468), int32(1444), int32(289151))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L31
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v167 int32
	_ = v167
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v430 int32
	_ = v430
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v475 int32
	_ = v475
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v631 int32
	_ = v631
	var v634 float64
	_ = v634
	var v635 int32
	_ = v635
	var v637 float64
	_ = v637
	var v638 int32
	_ = v638
	var v642 float64
	_ = v642
	var v650 float64
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v682 int32
	_ = v682
	var v685 float64
	_ = v685
	var v686 int32
	_ = v686
	var v688 float64
	_ = v688
	var v689 int32
	_ = v689
	var v693 float64
	_ = v693
	var v698 float64
	_ = v698
	var v702 float64
	_ = v702
	var v703 float64
	_ = v703
	var v705 float64
	_ = v705
	var v713 float64
	_ = v713
	var v716 int32
	_ = v716
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	v28 = m.G0
	v30 = v28 - int32(16)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v32 != int32(17) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v36 = F_pg_detoast_datum_packed(m, v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
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
	v743 = m.ExcPending
	if v743 != 0 {
		goto L4
	} else {
		goto L149
	}
L4:
	;
	return int32(0)
L5:
	;
	v40 = m.G0
	v42 = v40 - int32(128)
	m.G0 = v42
	v45 = v30 + int32(15)
	v46 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v46)
	if l1 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v50 = int32(27)
	goto L8
L7:
	;
	v50 = int32(19)
	goto L8
L8:
	;
	v51 = F_RE_compile_and_cache(m, v36, v50, l2)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v54 = v42 + int32(120)
	v55 = int32(16)
	v57 = v42 + int32(124)
	if v57 == int32(0) {
		v500 = v55
		goto L15
	} else {
		goto L16
	}
L10:
	;
	m.G0 = v42 + int32(128)
	if v602 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L11:
	;
	v581 = *(*int32)(unsafe.Add(mBase, _consts[462]))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)+4))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v582*int32(28))+uint32(_consts[994])))
	goto L94
L12:
	;
	v578 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v578)
	goto L11
L13:
	;
	v558 = F_pg_regerror(m, v553, v42+int32(16))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L4
	} else {
		goto L89
	}
L14:
	;
	switch v553 + int32(2) {
	case 0:
		goto L12
	case 1:
		goto L11
	default:
		goto L13
	case 3:
		v602 = int32(0)
		goto L10
	}
L15:
	;
	v553 = v500
	goto L14
L16:
	;
	if v54 == int32(0) {
		v500 = v55
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v62 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v62
	v67 = *(*int32)(unsafe.Add(mBase, _consts[995]))
	if v67 != int32(65239) {
		v500 = v55
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[996]))
	if v72 != int32(4) {
		v553 = int32(17)
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[997]))
	F_pg_set_regex_collation(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v79 = int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, _consts[998]))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+9)))
	if v82&int32(16) != 0 {
		v500 = v79
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+44)))
	if v86&int32(2) != 0 {
		v500 = v79
		goto L15
	} else {
		goto L22
	}
L22:
	;
	v90 = v85 + int32(36)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v92 = int32(2)
	v95 = F_palloc_extended(m, v91<<(uint(v92)%32), v92)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v95
	if v95 == int32(0) {
		v553 = int32(12)
		goto L14
	} else {
		goto L24
	}
L24:
	;
	v102 = v81 + int32(72)
	v103 = int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v90)+32))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v104+v105<<(uint(int32(2))%32))))
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109))))
	if v110 == int32(65535) {
		v475 = v103
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if base.Ui32(int32(-3)) < base.Ui32(v475) {
		v500 = v475
		goto L15
	} else {
		goto L87
	}
L26:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+20)))
	v118 = v109
	v119 = int32(-1)
	v122 = v110
	goto L27
L27:
	;
	v145 = v122 & int32(65535)
	if v113 != v145 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v153 == int32(-1) {
		v475 = v103
		goto L25
	} else {
		goto L39
	}
L29:
	;
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85+int32(58)))))
	if v145 != v147 {
		v475 = v103
		goto L25
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	if v119 == int32(-1) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L31
L33:
	;
	v155 = v118 + int32(8)
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155))))
	if v156 != int32(65535) {
		v118 = v155
		v119 = v153
		v122 = v156
		goto L27
	} else {
		goto L38
	}
L34:
	;
	v153 = v149
	goto L33
L35:
	;
	goto L36
L36:
	;
	if v119 != v149 {
		v475 = v103
		goto L25
	} else {
		goto L37
	}
L37:
	;
	v153 = v119
	goto L33
L38:
	;
	goto L28
L39:
	;
	v167 = v153
	goto L41
L40:
	;
	if v341 == int32(65535) {
		goto L69
	} else {
		goto L70
	}
L41:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v90)+32))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v192+v167<<(uint(int32(2))%32))))
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v196))))
	if v197 == int32(65535) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v90)+32))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v322+v167<<(uint(int32(2))%32))))
	v327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v326))))
	v338 = v326
	v341 = v327
	goto L40
L43:
	;
	goto L42
L44:
	;
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+20)))
	v204 = v197
	v205 = int32(-1)
	v211 = int32(65535)
	v212 = v196
	goto L45
L45:
	;
	v231 = v204 & int32(65535)
	if v231 == v202 {
		v251 = v205
		v252 = v211
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v259 = int32(65535)
	v260 = v252 & v259
	if v260 == v259 {
		goto L43
	} else {
		goto L59
	}
L47:
	;
	v255 = v212 + int32(8)
	v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v255))))
	if v256 != int32(65535) {
		v204 = v256
		v205 = v251
		v211 = v252
		v212 = v255
		goto L45
	} else {
		goto L58
	}
L48:
	;
	v233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85+int32(58)))))
	if v231 == v233 {
		v251 = v205
		v252 = v211
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+24)))
	if v231 == v235 {
		v338 = v196
		v341 = v197
		goto L40
	} else {
		goto L50
	}
L50:
	;
	if v231 == int32(65534) {
		v338 = v196
		v341 = v197
		goto L40
	} else {
		goto L51
	}
L51:
	;
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85+int32(62)))))
	if v231 == v239 {
		v338 = v196
		v341 = v197
		goto L40
	} else {
		goto L52
	}
L52:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v241 <= base.I32_extend16_s(v204) {
		v338 = v196
		v341 = v197
		goto L40
	} else {
		goto L53
	}
L53:
	;
	v244 = int32(65535)
	v245 = v211 & v244
	if v245 == v244 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	v251 = v248
	v252 = v204
	goto L47
L55:
	;
	goto L56
L56:
	;
	if v231 != v245 {
		v338 = v196
		v341 = v197
		goto L40
	} else {
		goto L57
	}
L57:
	;
	v251 = int32(-1)
	v252 = v204
	goto L47
L58:
	;
	goto L46
L59:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v102)+20))
	v267 = v263 + base.I32_extend16_s(v252)*int32(24)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if v268 != int32(1) {
		goto L43
	} else {
		goto L60
	}
L60:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v267)+4))
	if v271 != 0 {
		goto L43
	} else {
		goto L61
	}
L61:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v267)+16))
	if base.Ui32(v272) <= base.Ui32(int32(2047)) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if v281&int32(65535) != v260 {
		goto L43
	} else {
		goto L66
	}
L63:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v102)+24))
	v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v275+v272<<(uint(int32(1))%32)))))
	v281 = v279
	goto L62
L64:
	;
	goto L65
L65:
	;
	v280 = F_pg_reg_getcolor(m, v102, v272)
	mBase = m.M
	v281 = v280
	goto L62
L66:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v285 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v95+v285<<(uint(int32(2))%32)))) = v272
	if v251 != int32(-1) {
		v167 = v251
		goto L41
	} else {
		goto L67
	}
L67:
	;
	goto L43
L68:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	if v430 == v457 {
		v475 = int32(-2)
		goto L25
	} else {
		goto L83
	}
L69:
	;
	v430 = int32(-1)
	goto L68
L70:
	;
	v357 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+24)))
	v362 = int32(-1)
	v371 = v338
	v374 = v341
	goto L71
L71:
	;
	if v357 != v374 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v430 = v395
	goto L68
L73:
	;
	v389 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85+int32(62)))))
	if v374 != v389 {
		goto L69
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
	if v362 == int32(-1) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L75
L77:
	;
	v397 = v371 + int32(8)
	v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v397))))
	if v398 != int32(65535) {
		v362 = v395
		v371 = v397
		v374 = v398
		goto L71
	} else {
		goto L82
	}
L78:
	;
	v395 = v391
	goto L77
L79:
	;
	goto L80
L80:
	;
	if v362 != v391 {
		goto L69
	} else {
		goto L81
	}
L81:
	;
	v395 = v362
	goto L77
L82:
	;
	goto L72
L83:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v461 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v462 = int32(-1)
	goto L86
L85:
	;
	v462 = int32(1)
	goto L86
L86:
	;
	v475 = v462
	goto L25
L87:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	F_pfree(m, v492)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	v495 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v495
	v500 = v475
	goto L15
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v42 + int32(16)
	F_errmsg(m, int32(215208), v42)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(521411), int32(2068), int32(28014))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L4
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
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v42)+120))
	v592 = F_palloc(m, v587*v588+int32(1))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v42)+124))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v42)+120))
	v596 = F_pg_wchar2mb_with_len(m, v594, v592, v595)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+120)) = v596
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v42)+124))
	F_pfree(m, v599)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	v602 = v592
	goto L10
L98:
	;
	m.G0 = v30 + int32(16)
	return v731
L99:
	;
	v608 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v608
	if l4 == v608 {
		v731 = v608
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v654 = F_string_to_const(m, v602, v32)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L4
	} else {
		goto L118
	}
L102:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v615 = F_text_to_cstring(m, v614)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L4
	} else {
		goto L106
	}
L103:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = v650
	F_pfree(m, v615)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L4
	} else {
		goto L117
	}
L104:
	;
	if base.F64_lt(v642, float64(0)) != 0 {
		v650 = float64(0)
		goto L103
	} else {
		goto L115
	}
L105:
	;
	v637 = F_regex_selectivity_sub(m, v615, v617)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L4
	} else {
		goto L114
	}
L106:
	;
	v617 = F_strlen(m, v615)
	mBase = m.M
	if v617 <= int32(0) {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v621 = v617 - int32(1)
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615+v621))))
	if v623 != int32(36) {
		goto L105
	} else {
		goto L108
	}
L108:
	;
	if v617 != int32(1) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615+v617-int32(2)))))
	if v631 == int32(92) {
		goto L105
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v634 = F_regex_selectivity_sub(m, v615, v621)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L4
	} else {
		goto L113
	}
L112:
	;
	goto L111
L113:
	;
	v642 = v634
	goto L104
L114:
	;
	v642 = base.F64_mul(v637, float64(5))
	goto L104
L115:
	;
	if base.F64_gt(v642, float64(1)) == int32(0) {
		v650 = v642
		goto L103
	} else {
		goto L116
	}
L116:
	;
	v650 = float64(1)
	goto L103
L117:
	;
	v731 = v608
	goto L98
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v654
	if l4 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	F_pfree(m, v602)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L4
	} else {
		goto L145
	}
L120:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+15)))
	if v659 == int32(1) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(4607182418800017408)
	goto L119
L122:
	;
	goto L123
L123:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v665 = F_text_to_cstring(m, v664)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	v667 = F_strlen(m, v665)
	mBase = m.M
	v668 = F_strlen(m, v602)
	mBase = m.M
	if v667 <= int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	if int32(0) < v668 {
		goto L136
	} else {
		goto L137
	}
L126:
	;
	v688 = F_regex_selectivity_sub(m, v665, v667)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L4
	} else {
		goto L134
	}
L127:
	;
	v672 = v667 - int32(1)
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665+v672))))
	if v674 != int32(36) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	if v667 != int32(1) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667+v665-int32(2)))))
	if v682 == int32(92) {
		goto L126
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v685 = F_regex_selectivity_sub(m, v665, v672)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L4
	} else {
		goto L133
	}
L132:
	;
	goto L131
L133:
	;
	v693 = v685
	goto L125
L134:
	;
	v693 = base.F64_mul(v688, float64(5))
	goto L125
L135:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = v713
	F_pfree(m, v665)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L4
	} else {
		goto L144
	}
L136:
	;
	v698 = F_pow(m, float64(0.2), base.F64_convert_i32_u(v668))
	mBase = m.M
	if base.F64_gt(v698, float64(0)) != 0 {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	v703 = v693
	goto L138
L138:
	;
	v705 = float64(0)
	if base.F64_lt(v703, v705) != 0 {
		v713 = v705
		goto L135
	} else {
		goto L142
	}
L139:
	;
	v702 = base.F64_div(v693, v698)
	goto L141
L140:
	;
	v702 = v693
	goto L141
L141:
	;
	v703 = v702
	goto L138
L142:
	;
	if base.F64_gt(v703, float64(1)) == int32(0) {
		v713 = v703
		goto L135
	} else {
		goto L143
	}
L143:
	;
	v713 = float64(1)
	goto L135
L144:
	;
	goto L119
L145:
	;
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+15)))
	if v727 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v728 = int32(2)
	goto L148
L147:
	;
	v728 = int32(1)
	goto L148
L148:
	;
	v731 = v728
	goto L98
L149:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	F_errmsg(m, int32(534432), int32(0))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(517715), int32(1106), int32(27995))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
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
	var v195 int32
	_ = v195
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
	v195 = v4
	goto L36
L36:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v200+v195<<(uint(int32(2))%32))))
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
	v358 = v195 + int32(1)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v358 < v359 {
		v195 = v358
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
	*(*int32)(unsafe.Add(mBase, uint32(v305))) = l2 + v306
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
	var v98 int32
	_ = v98
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
		goto L31
	} else {
		goto L53
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v175
L3:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	if v119 == int32(0) {
		goto L1
	} else {
		goto L33
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
	v23 = int32(580318)
	v27 = m.G0
	v29 = v27 - int32(32)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v30
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _consts[266])))
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
		goto L30
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
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _consts[267])))
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
		v98 = v11
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v98 - v11
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
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v98 = v94
	goto L22
L26:
	;
	v98 = v77
	goto L22
L27:
	;
	goto L28
L28:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v94 = v77 + int32(1)
	if v92 != 0 {
		v77 = v94
		v78 = v92
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v113 = F_DirectInputFunctionCallSafe(m, int32(547), v11, int32(-1), v10, v8+int32(12))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(0)
L32:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v175 = v117
	goto L2
L33:
	;
	v122 = F_stringToQualifiedNameList(m, v11, v10)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	if v122 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v126)
	v175 = int32(0)
	goto L2
L36:
	;
	goto L37
L37:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v129 != int32(1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v132 = int32(0)
	v133 = F_errsave_start(m, v10)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L31
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v153 = F_get_role_oid(m, v151, int32(1))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L31
	} else {
		goto L46
	}
L41:
	;
	if v133 == int32(0) {
		v175 = v132
		goto L2
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L31
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(30646), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L31
	} else {
		goto L44
	}
L44:
	;
	F_errsave_finish(m, v10, int32(526468), int32(1564), int32(291124))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L31
	} else {
		goto L45
	}
L45:
	;
	v175 = v132
	goto L2
L46:
	;
	if v153 != 0 {
		v175 = v153
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v155 = int32(0)
	v156 = F_errsave_start(m, v10)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L31
	} else {
		goto L48
	}
L48:
	;
	if v156 == int32(0) {
		v175 = v155
		goto L2
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L31
	} else {
		goto L50
	}
L50:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v165
	F_errmsg(m, int32(78396), v8)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L31
	} else {
		goto L51
	}
L51:
	;
	F_errsave_finish(m, v10, int32(526468), int32(1572), int32(291124))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L31
	} else {
		goto L52
	}
L52:
	;
	v175 = v155
	goto L2
L53:
	;
	F_errmsg_internal(m, int32(433348), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L31
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(526468), int32(1554), int32(291124))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L31
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
				F_errmsg(m, int32(321140), v6)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(518464), int32(219), int32(321044))
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
	var v26 int32
	_ = v26
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
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
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
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v14)+288))
	F_pfree(m, v441)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L34
	} else {
		goto L135
	}
L3:
	;
	if int32(2) <= v255 {
		goto L122
	} else {
		goto L123
	}
L4:
	;
	v348 = F_getObjectDescription(m, l3, int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L34
	} else {
		goto L115
	}
L5:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v316 = F_getObjectDescription(m, v33+int32(4), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L34
	} else {
		goto L108
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = v5
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
	v33 = v19 + v26<<(uint(int32(4))%32)
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
	v40 = v26 + int32(1)
	if v40 != v16 {
		v26 = v40
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
		v93 = v62
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
	v100 = m.ExcPending
	if v100 != 0 {
		goto L34
	} else {
		goto L35
	}
L19:
	;
	if v93 == int32(0) {
		goto L1
	} else {
		goto L33
	}
L20:
	;
	goto L19
L21:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	if base.Ui32(v57-int32(15)) <= base.Ui32(int32(1)) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v78 = int32(0)
	if v57 == int32(16) {
		v93 = v78
		goto L20
	} else {
		goto L30
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
	if v57 == int32(20) {
		goto L22
	} else {
		goto L27
	}
L26:
	;
	v93 = v62
	goto L20
L27:
	;
	if v66 == int32(15) {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	if v66 <= v57 {
		v93 = v62
		goto L20
	} else {
		goto L29
	}
L29:
	;
	goto L22
L30:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	if v82 != int32(2) {
		v93 = v78
		goto L20
	} else {
		goto L31
	}
L31:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, _consts[174])))
	if v86 != 0 {
		v93 = v78
		goto L20
	} else {
		goto L32
	}
L32:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v93 = base.B2i32(v57 == int32(17)) | base.B2i32(v90 <= v57)
	goto L20
L33:
	;
	goto L18
L34:
	;
	return
L35:
	;
	F_initStringInfo(m, v14+int32(272))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v107 = v105 - int32(1)
	if v107 < int32(0) {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v111 = int32(0)
	v115 = v111
	v118 = v107
	v120 = v111
	v122 = int32(1)
	goto L38
L38:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v127 = v124 + v118<<(uint(int32(4))%32)
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127))))
	if v128&int32(257) != 0 {
		v253 = v115
		v255 = v120
		v257 = v122
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if int32(0) < v253 {
		goto L91
	} else {
		goto L92
	}
L40:
	;
	if int32(0) < v118 {
		v115 = v253
		v118 = v118 - int32(1)
		v120 = v255
		v122 = v257
		goto L38
	} else {
		goto L89
	}
L41:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v136 = F_getObjectDescription(m, v131+v118*int32(12), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L34
	} else {
		goto L42
	}
L42:
	;
	if v136 == int32(0) {
		v253 = v115
		v255 = v120
		v257 = v122
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v140&int32(60) != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	F_pfree(m, v136)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L34
	} else {
		goto L88
	}
L45:
	;
	v145 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L34
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if l1 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L48:
	;
	if v145 == int32(0) {
		v247 = v115
		v249 = v120
		v250 = v122
		goto L44
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+256)) = v136
	F_errmsg_internal(m, int32(193885), v14+int32(256))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L34
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(517236), int32(1083), int32(134009))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L34
	} else {
		goto L51
	}
L51:
	;
	v247 = v115
	v249 = v120
	v250 = v122
	goto L44
L52:
	;
	v165 = F_getObjectDescription(m, v127+int32(4), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L34
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if v120 <= int32(99) {
		goto L75
	} else {
		goto L76
	}
L55:
	;
	if v165 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if v120 <= int32(99) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	goto L58
L58:
	;
	v247 = v115 + int32(1)
	v249 = v120
	v250 = int32(0)
	goto L44
L59:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v14)+276))
	if v190 != 0 {
		goto L68
	} else {
		goto L69
	}
L60:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v14)+292))
	if v169 != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v188 = v115 + int32(1)
	v189 = v120
	goto L59
L63:
	;
	F_appendStringInfoChar(m, v14+int32(288), int32(10))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L34
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+212)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = v136
	F_appendStringInfo(m, v14+int32(288), int32(195563), v14+int32(208))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L34
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	v188 = v115
	v189 = v120 + int32(1)
	goto L59
L68:
	;
	F_appendStringInfoChar(m, v14+int32(272), int32(10))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L34
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+196)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v14)+192)) = v136
	F_appendStringInfo(m, v14+int32(272), int32(195563), v14+int32(192))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L34
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	F_pfree(m, v165)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L34
	} else {
		goto L73
	}
L73:
	;
	v247 = v188
	v249 = v189
	v250 = int32(0)
	goto L44
L74:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v14)+276))
	if v233 != 0 {
		goto L83
	} else {
		goto L84
	}
L75:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v14)+292))
	if v213 != 0 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	v231 = v115 + int32(1)
	v232 = v120
	goto L74
L78:
	;
	F_appendStringInfoChar(m, v14+int32(288), int32(10))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L34
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v136
	F_appendStringInfo(m, v14+int32(288), int32(193910), v14+int32(240))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L34
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	v231 = v115
	v232 = v120 + int32(1)
	goto L74
L83:
	;
	F_appendStringInfoChar(m, v14+int32(272), int32(10))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L34
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+224)) = v136
	F_appendStringInfo(m, v14+int32(272), int32(193910), v14+int32(224))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L34
	} else {
		goto L87
	}
L86:
	;
	goto L85
L87:
	;
	v247 = v231
	v249 = v232
	v250 = v122
	goto L44
L88:
	;
	v253 = v247
	v255 = v249
	v257 = v250
	goto L40
L89:
	;
	goto L39
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L34
	} else {
		goto L100
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+176)) = v253
	if v253 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	if v257 != 0 {
		goto L3
	} else {
		goto L99
	}
L94:
	;
	v271 = int32(705672)
	goto L96
L95:
	;
	v271 = int32(705719)
	goto L96
L96:
	;
	F_appendStringInfo(m, v14+int32(288), v271, v14+int32(176))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L34
	} else {
		goto L97
	}
L97:
	;
	if v257 == int32(0) {
		goto L90
	} else {
		goto L98
	}
L98:
	;
	goto L3
L99:
	;
	goto L90
L100:
	;
	F_errcode(m, int32(16909442))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L34
	} else {
		goto L101
	}
L101:
	;
	if l3 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(306663), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L34
	} else {
		goto L103
	}
L103:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v14)+288))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v289
	F_errdetail_internal(m, int32(217416), v14+int32(112))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L34
	} else {
		goto L104
	}
L104:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v14)+272))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v296
	F_errdetail_log(m, int32(217416), v14+int32(96))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L34
	} else {
		goto L105
	}
L105:
	;
	F_errhint(m, int32(643073), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L34
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(517236), int32(1161), int32(134009))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L34
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L34
	} else {
		goto L109
	}
L109:
	;
	F_errcode(m, int32(16909442))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L34
	} else {
		goto L110
	}
L110:
	;
	v329 = F_getObjectDescription(m, v312+v26*int32(12), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L34
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v316
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v329
	F_errmsg(m, int32(110839), v14+int32(16))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L34
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v316
	F_errhint(m, int32(683397), v14)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L34
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(517236), int32(1018), int32(134009))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L34
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = v348
	F_errmsg(m, int32(110963), v14+int32(160))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L34
	} else {
		goto L116
	}
L116:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v14)+288))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = v356
	F_errdetail_internal(m, int32(217416), v14+int32(144))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L34
	} else {
		goto L117
	}
L117:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v14)+272))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v363
	F_errdetail_log(m, int32(217416), v14+int32(128))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L34
	} else {
		goto L118
	}
L118:
	;
	F_errhint(m, int32(643073), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L34
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(517236), int32(1154), int32(134009))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L34
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
	F_errfinish(m, int32(517236), v426, int32(134009))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L34
	} else {
		goto L134
	}
L122:
	;
	v383 = F_errstart(m, v57, int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L34
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	if v255 != int32(1) {
		goto L2
	} else {
		goto L130
	}
L125:
	;
	if v383 == int32(0) {
		goto L2
	} else {
		goto L126
	}
L126:
	;
	v387 = v253 + v255
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v387
	F_errmsg_plural(m, int32(118390), int32(133725), v387, v14-int32(-64))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L34
	} else {
		goto L127
	}
L127:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v14)+288))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v395
	F_errdetail_internal(m, int32(217416), v14+int32(48))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L34
	} else {
		goto L128
	}
L128:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v14)+272))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v402
	F_errdetail_log(m, int32(217416), v14+int32(32))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L34
	} else {
		goto L129
	}
L129:
	;
	v426 = int32(1171)
	goto L121
L130:
	;
	v413 = F_errstart(m, v57, int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L34
	} else {
		goto L131
	}
L131:
	;
	if v413 == int32(0) {
		goto L2
	} else {
		goto L132
	}
L132:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v14)+288))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v417
	F_errmsg_internal(m, int32(217416), v14+int32(80))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L34
	} else {
		goto L133
	}
L133:
	;
	v426 = int32(1177)
	goto L121
L134:
	;
	goto L2
L135:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v14)+272))
	F_pfree(m, v444)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L34
	} else {
		goto L136
	}
L136:
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
	v2 = int32(*(*uint8)(unsafe.Add(mBase, _consts[638])))
	if v2 == int32(0) {
		return
	} else {
		v6 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[628])) = v6
		v9 = *(*int32)(unsafe.Add(mBase, _consts[629]))
		if v9 <= v6 {
			return
		} else {
			v15 = m.G0
			v16 = int32(16)
			v17 = v15 - v16
			m.G0 = v17
			F___gettimeofday(m, v17)
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v35 int64
	_ = v35
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v55 int64
	_ = v55
	var v60 int64
	_ = v60
	var v65 int64
	_ = v65
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v106 float64
	_ = v106
	var v109 float64
	_ = v109
	var v113 float64
	_ = v113
	v6 = l0 + int32(8)
	v9 = int32(4645608)
	v10 = int32(4645600)
	v11 = *(*int64)(unsafe.Add(mBase, _consts[824]))
	v13 = *(*int64)(unsafe.Add(mBase, _consts[823]))
	v14 = v11 ^ v13
	*(*int64)(unsafe.Add(mBase, _consts[823])) = base.I64_rotl(v14, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[824])) = v14<<(uint(int64(16))%64) ^ base.I64_rotl(v11, int64(24)) ^ v14
	v35 = base.I64_extend_i32_u(base.I32_wrap_i64(int64(base.Ui64(base.I64_rotl(v11*int64(5), int64(7))*int64(9)) >> (uint(int64(32)) % 64))))
	v39 = v35 + int64(4354685564936845354)
	v40 = int64(30)
	v43 = int64(-4658895280553007687)
	v44 = (int64(base.Ui64(v39)>>(uint(v40)%64)) ^ v39) * v43
	v45 = int64(27)
	v48 = int64(-7723592293110705685)
	v49 = (int64(base.Ui64(v44)>>(uint(v45)%64)) ^ v44) * v48
	v50 = int64(31)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(base.Ui64(v49)>>(uint(v50)%64)) ^ v49
	v55 = v35 - int64(7046029254386353131)
	v60 = (int64(base.Ui64(v55)>>(uint(v40)%64)) ^ v55) * v43
	v65 = (int64(base.Ui64(v60)>>(uint(v45)%64)) ^ v60) * v48
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(base.Ui64(v65)>>(uint(v50)%64)) ^ v65
	if v55|v39 == int64(0) {
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(1442695040888963407)
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(6364136223846793005)
	} else {
	}
	for {
		v83 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
		v84 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
		v85 = v83 ^ v84
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = base.I64_rotl(v85, int64(37))
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v85<<(uint(int64(16))%64) ^ base.I64_rotl(v83, int64(24)) ^ v85
		v106 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v83*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
		mBase = m.M
		if base.F64_eq(v106, float64(0)) != 0 {
			continue
		} else {
			break
		}
		break
	}
	v109 = F_log(m, v106)
	mBase = m.M
	v113 = F_exp(m, base.F64_div(base.F64_neg(v109), base.F64_convert_i32_s(l1)))
	mBase = m.M
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v113
	return
}
func F_resize_intArrayType(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	if l1 <= int32(0) {
		v18 = F_construct_empty_array(m, int32(23))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			return v18
		}
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v26 = F_ArrayGetNItems(m, v23, l0+int32(16))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			if v26 == l1 {
				v166 = l0
				return v166
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v29 != 0 {
					v37 = v29
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v37 = (v30<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				}
				v40 = v37 + l1<<(uint(int32(2))%32)
				v41 = F_repalloc(m, l0, v40)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = v40 << (uint(int32(2)) % 32)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
					if v46 <= int32(0) {
						v166 = v41
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = l1
						if v46 == int32(1) {
							v166 = v41
						} else {
							v53 = v41 + int32(16)
							v54 = int32(1)
							v55 = v46 - v54
							v56 = int32(7)
							v57 = v55 & v56
							if base.Ui32(v56) <= base.Ui32(v46-int32(2)) {
								v82 = v54
								v84 = int32(0)
								for {
									v95 = v82 << (uint(int32(2)) % 32)
									v97 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v53+v95))) = v97
									*(*int32)(unsafe.Add(mBase, uint32(v95+(v41+int32(20))))) = v97
									*(*int32)(unsafe.Add(mBase, uint32(v95+(v41+int32(24))))) = v97
									*(*int32)(unsafe.Add(mBase, uint32(v95+(v41+int32(28))))) = v97
									*(*int32)(unsafe.Add(mBase, uint32(v95+(v41+int32(32))))) = v97
									*(*int32)(unsafe.Add(mBase, uint32(v95+(v41+int32(36))))) = v97
									*(*int32)(unsafe.Add(mBase, uint32(v95+(v41+int32(40))))) = v97
									*(*int32)(unsafe.Add(mBase, uint32(v95+(v41+int32(44))))) = v97
									v120 = int32(8)
									v121 = v82 + v120
									v123 = v84 + v120
									if v123 != v55&int32(-8) {
										v82 = v121
										v84 = v123
										continue
									} else {
										break
									}
									break
								}
								v127 = v121
							} else {
								v127 = v54
							}
							if v57 == int32(0) {
								v166 = v41
							} else {
								v143 = int32(0)
								v144 = v127
								for {
									v159 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v53+v144<<(uint(int32(2))%32)))) = v159
									v164 = v143 + v159
									if v164 != v57 {
										v143 = v164
										v144 = v144 + v159
										continue
									} else {
										break
									}
									break
								}
								v166 = v41
							}
						}
					}
					return v166
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
					F_errmsg_internal(m, int32(390216), int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						F_errfinish(m, int32(524018), int32(674), int32(143218))
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
									F_errmsg(m, int32(204554), v6)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										F_errfinish(m, int32(524018), int32(670), int32(143218))
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
							F_errmsg(m, int32(204554), v6)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								F_errfinish(m, int32(524018), int32(670), int32(143218))
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v42 int32
	_ = v42
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int64
	_ = v110
	var v112 int32
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
	v21 = v14 + int32(76)
	v22 = m.G0
	v24 = v22 - int32(32)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+49)))
	v29 = v27 << (uint(int32(5)) % 32)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[302])))
	if v32 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_RmgrNotFound(m, v27)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v39 = v26
	v40 = v32
	goto L5
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[408])))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+48)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[409])))
	F_appendStringInfoString(m, v21, v40)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[302])))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v39 = v38
	v40 = v37
	goto L5
L7:
	;
	F_appendStringInfoChar(m, v21, int32(47))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v49 = m.T0[v43].(func(*base.Module, int32) int32)(m, v42)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	m.T0[v41].(func(*base.Module, int32, int32))(m, v21, l0)
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
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v42 & int32(240)
	F_appendStringInfo(m, v21, int32(782983), v24)
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
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v49
	F_appendStringInfo(m, v21, int32(782690), v24+int32(16))
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
	m.G0 = v24 + int32(32)
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
	v106 = v99 + v90*int32(52) + int32(76)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v107 != int32(1) {
		v124 = v97
		goto L24
	} else {
		goto L26
	}
L26:
	;
	if v92 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v106)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v92))) = v110
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+8)) = v112
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
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
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
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v106)+20))
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
	F_appendStringInfo(m, v14+int32(76), int32(51779), v14+int32(48))
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
	F_appendStringInfo(m, v14+int32(76), int32(51744), v14+int32(16))
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
	F_appendStringInfoString(m, v14+int32(76), int32(543478))
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
	F_errcontext_msg(m, int32(192890), v14)
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
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
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(1104)
	m.G0 = v11
	v13 = F_AllocateDir(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(1104)
	return v228 & int32(1)
L2:
	;
	return int32(0)
L3:
	;
	if v13 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v35 = F_palloc(m, int32(32))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L11
	}
L7:
	;
	if v21 == int32(0) {
		v228 = v2
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(311552), v11)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(525732), int32(63), int32(431164))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v228 = v2
	goto L1
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	v40 = int32(1)
	v41 = F_readdir(m, v13)
	mBase = m.M
	if v41 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v45 = v41
	v46 = int32(8)
	v47 = v40
	v49 = v35
	v50 = v2
	goto L15
L13:
	;
	v141 = v40
	v143 = v35
	v144 = v2
	goto L14
L14:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v146 == int32(0) {
		v167 = v141
		goto L39
	} else {
		goto L40
	}
L15:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+19)))
	if v51 != int32(46) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v141 = v130
	v143 = v131
	v144 = v132
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	v136 = F_readdir(m, v13)
	mBase = m.M
	if v136 != 0 {
		v45 = v136
		v46 = v129
		v47 = v130
		v49 = v131
		v50 = v132
		goto L15
	} else {
		goto L38
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v45 + int32(19)
	v73 = F_pg_snprintf(m, v11+int32(80), int32(1024), int32(188238), v11-int32(-64))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L2
	} else {
		goto L23
	}
L19:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+20)))
	if v54 == int32(0) {
		v129 = v46
		v130 = v47
		v131 = v49
		v132 = v50
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+20)))
	if v57 != int32(46) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+21)))
	if v60 == int32(0) {
		v129 = v46
		v130 = v47
		v131 = v49
		v132 = v50
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	v79 = F_get_dirent_type(m, v11+int32(80), v45, int32(0), int32(19))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L26
	}
L24:
	;
	v102 = F_unlink(m, v11+int32(80))
	mBase = m.M
	if v102 == int32(0) {
		v129 = v46
		v130 = v47
		v131 = v49
		v132 = v50
		goto L17
	} else {
		goto L32
	}
L25:
	;
	if v46 == v50 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	switch v79 {
	case 0:
		v129 = v46
		v130 = v47
		v131 = v49
		v132 = v50
		goto L17
	default:
		goto L24
	case 3:
		goto L25
	}
L27:
	;
	v84 = F_repalloc(m, v49, v46<<(uint(int32(3))%32))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L2
	} else {
		goto L30
	}
L28:
	;
	v88 = v46
	v89 = v49
	goto L29
L29:
	;
	v95 = F_pstrdup(m, v11+int32(80))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L31
	}
L30:
	;
	v88 = v46 << (uint(int32(1)) % 32)
	v89 = v84
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89+v50<<(uint(int32(2))%32)))) = v95
	v129 = v88
	v130 = v47
	v131 = v89
	v132 = v50 + int32(1)
	goto L17
L32:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v106 == int32(44) {
		v129 = v46
		v130 = v47
		v131 = v49
		v132 = v50
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
		v129 = v46
		v130 = v109
		v131 = v49
		v132 = v50
		goto L17
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v11 + int32(80)
	F_errmsg_internal(m, int32(314276), v11+int32(48))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(525732), int32(97), int32(431164))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v129 = v46
	v130 = v109
	v131 = v49
	v132 = v50
	goto L17
L38:
	;
	goto L16
L39:
	;
	F_FreeDir(m, v13)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L2
	} else {
		goto L45
	}
L40:
	;
	v149 = int32(0)
	v152 = F_errstart(m, int32(19), v149)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	if v152 == int32(0) {
		v167 = v149
		goto L39
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l0
	F_errmsg_internal(m, int32(311751), v11+int32(32))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(525732), int32(106), int32(431164))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v167 = v149
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
	v173 = int32(0)
	v175 = v167
	goto L49
L47:
	;
	v196 = v167
	goto L48
L48:
	;
	v200 = F_rmdir(m, l0)
	mBase = m.M
	if v200 == int32(0) {
		v221 = v196
		goto L54
	} else {
		goto L55
	}
L49:
	;
	v181 = v143 + v173<<(uint(int32(2))%32)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v183 = F_rmtree(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L51
	}
L50:
	;
	v196 = v188
	goto L48
L51:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	F_pfree(m, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v188 = v183 & v175
	v190 = v173 + int32(1)
	if v190 != v144 {
		v173 = v190
		v175 = v188
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
	v223 = m.ExcPending
	if v223 != 0 {
		goto L2
	} else {
		goto L60
	}
L55:
	;
	v203 = int32(0)
	v206 = F_errstart(m, int32(19), v203)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	if v206 == int32(0) {
		v221 = v203
		goto L54
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l0
	F_errmsg_internal(m, int32(311679), v11+int32(16))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(525732), int32(124), int32(431164))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v221 = v203
	goto L54
L60:
	;
	v228 = v221
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
	var v52 int32
	_ = v52
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
		v52 = v21
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v59 = v52
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
		v52 = v45
		goto L13
	} else {
		goto L20
	}
L19:
	;
	v52 = v45
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
	v69 = *(*int32)(unsafe.Add(mBase, _consts[538]))
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
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
	var v182 int32
	_ = v182
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
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
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
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v23 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if int32(0) < v17 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v26 = int32(4)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v28&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v41 = int32(1)
	if v23&v41 != 0 {
		v53 = int32(base.Ui32(v23)>>(uint(v41)%32)) - v41
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v37 = v26
	goto L10
L9:
	;
	v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
	goto L10
L10:
	;
	if v28 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v40 = v26
	goto L13
L12:
	;
	v40 = v37
	goto L13
L13:
	;
	v53 = v40
	goto L4
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v55 = v17
	goto L17
L16:
	;
	v55 = int32(0)
	goto L17
L17:
	;
	v56 = int32(0)
	if v56 < v53 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v59 = v53
	goto L20
L19:
	;
	v59 = v56
	goto L20
L20:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v60 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v91 = int32(1)
	v92 = v13 + v91
	v94 = v13 + int32(4)
	if v23&v91 != 0 {
		goto L32
	} else {
		goto L33
	}
L22:
	;
	v63 = int32(4)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v65&int32(254) == int32(2) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v78 = int32(1)
	if v60&v78 != 0 {
		v90 = int32(base.Ui32(v60)>>(uint(v78)%32)) - v78
		goto L21
	} else {
		goto L31
	}
L25:
	;
	v74 = v63
	goto L27
L26:
	;
	v74 = base.B2i32(v65 == int32(18)) << (uint(v63) % 32)
	goto L27
L27:
	;
	if v65 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v77 = v63
	goto L30
L29:
	;
	v77 = v74
	goto L30
L30:
	;
	v90 = v77
	goto L21
L31:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v90 = int32(base.Ui32(v84)>>(uint(int32(2))%32)) - int32(4)
	goto L21
L32:
	;
	v97 = v92
	goto L34
L33:
	;
	v97 = v94
	goto L34
L34:
	;
	v98 = F_pg_mbstrlen_with_len(m, v97, v59)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _consts[462]))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v102*int32(28))+uint32(_consts[994])))
	goto L37
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L82
	}
L37:
	;
	if v98 < v55 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v110 = v98
	goto L40
L39:
	;
	v110 = v55
	goto L40
L40:
	;
	if v90 <= int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v113 = v110
	goto L43
L42:
	;
	v113 = v55
	goto L43
L43:
	;
	v115 = base.I64_extend_i32_s(v107) * base.I64_extend_i32_s(v113)
	v119 = base.I32_wrap_i64(v115)
	if base.I32_wrap_i64(int64(base.Ui64(v115)>>(uint(int64(32))%64))) != v119>>(uint(int32(31))%32) {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v124 = v119 + int32(4)
	if v124 < v119 {
		goto L36
	} else {
		goto L45
	}
L45:
	;
	if base.Ui32(int32(1073741824)) <= base.Ui32(v124) {
		goto L36
	} else {
		goto L46
	}
L46:
	;
	v128 = v113 - v110
	v129 = F_palloc(m, v124)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v132 = v129 + int32(4)
	if v110 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v133&int32(1) != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v156 = v132
	goto L50
L50:
	;
	if v128 != 0 {
		goto L62
	} else {
		goto L63
	}
L51:
	;
	v136 = v92
	goto L53
L52:
	;
	v136 = v94
	goto L53
L53:
	;
	v137 = v132
	v138 = v136
	v141 = v110
	goto L54
L54:
	;
	v148 = F_pg_mblen_unbounded(m, v138)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L56
	}
L55:
	;
	v156 = v153
	goto L50
L56:
	;
	if v148 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v153 = v151 + v148
	v155 = v141 - int32(1)
	if v155 != 0 {
		v137 = v153
		v138 = v138 + v148
		v141 = v155
		goto L54
	} else {
		goto L61
	}
L58:
	;
	v150 = F__emscripten_memcpy_bulkmem(m, v137, v138, v148)
	mBase = m.M
	v151 = v150
	goto L60
L59:
	;
	v151 = v137
	goto L60
L60:
	;
	goto L57
L61:
	;
	goto L55
L62:
	;
	v167 = int32(1)
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v169&v167 != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v200 = v156
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = (v200 - v129) << (uint(int32(2)) % 32)
	return v129
L65:
	;
	v172 = v167
	goto L67
L66:
	;
	v172 = int32(4)
	goto L67
L67:
	;
	v173 = v21 + v172
	v174 = int32(0)
	if v174 < v90 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v177 = v90
	goto L70
L69:
	;
	v177 = v174
	goto L70
L70:
	;
	v178 = v173 + v177
	v179 = v156
	v180 = v173
	v182 = v128
	goto L71
L71:
	;
	v190 = F_pg_mblen_range(m, v180, v178)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	v200 = v197
	goto L64
L73:
	;
	if v190 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v194 = v180 + v190
	if v194 == v178 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v192 = F__emscripten_memcpy_bulkmem(m, v179, v180, v190)
	mBase = m.M
	v193 = v192
	goto L77
L76:
	;
	v193 = v179
	goto L77
L77:
	;
	goto L74
L78:
	;
	v196 = v173
	goto L80
L79:
	;
	v196 = v194
	goto L80
L80:
	;
	v197 = v193 + v190
	v199 = v182 - int32(1)
	if v199 != 0 {
		v179 = v197
		v180 = v196
		v182 = v199
		goto L71
	} else {
		goto L81
	}
L81:
	;
	goto L72
L82:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(421000), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(518624), int32(304), int32(486544))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
