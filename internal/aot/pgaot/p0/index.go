package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BuildIndexInfo(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v33 int32
	_ = v33
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v150 int32
	_ = v150
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+8)))
	if base.Ui32(int32(65503)) < base.Ui32((v19-int32(33))&int32(65535)) {
		v27 = v19 & int32(3)
		v29 = v18 + int32(48)
		v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+10)))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+84))
		v33 = F_RelationGetIndexExpressions(m, l0)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			v37 = F_RelationGetIndexPredicate(m, l0)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+12)))
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+13)))
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+20)))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+28)))
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)))
				v47 = F_makeIndexInfo(m, v19, v30, v32, v33, v37, v39, v40, v41, int32(0), v44, v39&v45)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v50 = v47 + int32(12)
					v51 = int32(0)
					if base.Ui32(int32(3)) <= base.Ui32(v19-int32(1)) {
						v60 = v51
						v69 = v2
						for {
							v72 = v60 << (uint(int32(1)) % 32)
							v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72+v29))))
							*(*uint16)(unsafe.Add(mBase, uint32(v50+v72))) = uint16(v75)
							v78 = v72 | int32(2)
							v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+v78))))
							*(*uint16)(unsafe.Add(mBase, uint32(v50+v78))) = uint16(v81)
							v83 = int32(4)
							v84 = v72 | v83
							v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+v84))))
							*(*uint16)(unsafe.Add(mBase, uint32(v50+v84))) = uint16(v87)
							v90 = v72 | int32(6)
							v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90+v29))))
							*(*uint16)(unsafe.Add(mBase, uint32(v50+v90))) = uint16(v93)
							v96 = v60 + v83
							v98 = v69 + v83
							if v98 != v19&int32(60) {
								v60 = v96
								v69 = v98
								continue
							} else {
								break
							}
							break
						}
						v102 = v96
					} else {
						v102 = v51
					}
					if v27 != 0 {
						v115 = v102
						v123 = v2
						for {
							v126 = int32(1)
							v127 = v115 << (uint(v126) % 32)
							v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127+v29))))
							*(*uint16)(unsafe.Add(mBase, uint32(v50+v127))) = uint16(v130)
							v135 = v123 + v126
							if v135 != v27 {
								v115 = v115 + v126
								v123 = v135
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)))
					if v150 == int32(1) {
						F_RelationGetExclusionInfo(m, l0, v47+int32(92), v47+int32(96), v47+int32(100))
						mBase = m.M
						v160 = m.ExcPending
						if v160 != 0 {
							return int32(0)
						} else {
							m.G0 = v16 + int32(16)
							return v47
						}
					} else {
						m.G0 = v16 + int32(16)
						return v47
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v168 = m.ExcPending
		if v168 != 0 {
			return int32(0)
		} else {
			v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v169
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = v19
			F_errmsg_internal(m, int32(42701), v16)
			mBase = m.M
			v174 = m.ExcPending
			if v174 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(513935), int32(2439), int32(253525))
				mBase = m.M
				v179 = m.ExcPending
				if v179 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_CheckIndex(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v7 = l0 + l1<<(uint(int32(2))%32)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v8 < int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			F_errcode(m, int32(130))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_errmsg(m, int32(182121), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errfinish(m, int32(523450), int32(116), int32(30272))
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
	} else {
		if l2 <= v8 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				F_errcode(m, int32(130))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_errmsg(m, int32(182121), int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						F_errfinish(m, int32(523450), int32(116), int32(30272))
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
		} else {
			if int32(0) < l1 {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v7-int32(4))))
				if v8 < v16 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						F_errcode(m, int32(130))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							F_errmsg(m, int32(237775), int32(0))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								F_errfinish(m, int32(523450), int32(124), int32(30272))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
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
					if v8 == v16 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							F_errcode(m, int32(130))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								F_errmsg(m, int32(169987), int32(0))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									F_errfinish(m, int32(523450), int32(129), int32(30272))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
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
						return
					}
				}
			} else {
				return
			}
		}
	}
}
func F_CheckIndexCompatible(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
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
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v259 int32
	_ = v259
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
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v611 int32
	_ = v611
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v658 int32
	_ = v658
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	var v714 int32
	_ = v714
	v6 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v22 = F_IndexGetRelation(m, l0, v6)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l2 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v27 = v26
	goto L5
L4:
	;
	v27 = v6
	goto L5
L5:
	;
	v29 = F_SearchSysCache1(m, int32(1), l1)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	m.G0 = v19 + int32(48)
	return v714
L7:
	;
	F_ReleaseCatCache(m, v71)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L195
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L192
	}
L9:
	;
	if v29 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
	v33 = v31 + v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+68))
	v36 = F_GetIndexAmRoutine(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L1
	} else {
		goto L188
	}
L13:
	;
	F_ReleaseCatCache(m, v29)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+10)))
	v41 = int32(0)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+28)))
	v48 = F_makeIndexInfo(m, v27, v27, v34, v41, v41, v41, v41, v41, v41, v47, l4)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v51 = v27 << (uint(int32(2)) % 32)
	v52 = F_palloc(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v54 = F_palloc(m, v51)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v56 = F_palloc(m, v51)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v58 = F_palloc(m, v51)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v62 = F_palloc(m, v27<<(uint(int32(1))%32))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v64 = int32(0)
	F_ComputeIndexAttrs(m, v48, v52, v54, v56, v58, v62, l2, l3, v22, l1, v34, v40, v64, l4, v64, v64, v64)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v71 = F_SearchSysCache1(m, int32(34), l0)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v71 == int32(0) {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+22)))
	v79 = F_heap_attisnull(m, v71, int32(21), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v79 == int32(0) {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v85 = F_heap_attisnull(m, v71, int32(20), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v85 == int32(0) {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	v89 = v76 + v75
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+18)))
	if v90 == int32(0) {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v93 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+10)))
	v96 = F_SysCacheGetAttrNotNull(m, int32(34), v71, int32(17))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v100 = F_SysCacheGetAttrNotNull(m, int32(34), v71, int32(18))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v103 = v100 + int32(24)
	v105 = v93 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v105) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	if v167 != 0 {
		goto L7
	} else {
		goto L49
	}
L32:
	;
	v167 = int32(0)
	goto L31
L33:
	;
	v141 = v136
	v142 = v137
	v143 = v138
	goto L43
L34:
	;
	if (v103|v56)&int32(3) != 0 {
		v136 = v103
		v137 = v56
		v138 = v105
		goto L33
	} else {
		goto L37
	}
L35:
	;
	v129 = v103
	v130 = v56
	v131 = v105
	goto L36
L36:
	;
	if v131 == int32(0) {
		goto L32
	} else {
		goto L42
	}
L37:
	;
	v113 = v103
	v114 = v56
	v115 = v105
	goto L38
L38:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if v118 != v119 {
		v136 = v113
		v137 = v114
		v138 = v115
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v129 = v124
	v130 = v122
	v131 = v126
	goto L36
L40:
	;
	v121 = int32(4)
	v122 = v114 + v121
	v124 = v113 + v121
	v126 = v115 - v121
	if base.Ui32(int32(3)) < base.Ui32(v126) {
		v113 = v124
		v114 = v122
		v115 = v126
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v136 = v129
	v137 = v130
	v138 = v131
	goto L33
L43:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v146 == v147 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v167 = v146 - v147
	goto L31
L45:
	;
	v149 = int32(1)
	v154 = v143 - v149
	if v154 != 0 {
		v141 = v141 + v149
		v142 = v142 + v149
		v143 = v154
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	goto L32
L49:
	;
	v169 = v96 + int32(24)
	if base.Ui32(int32(4)) <= base.Ui32(v105) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	F_ReleaseCatCache(m, v71)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L68
	}
L51:
	;
	v231 = int32(0)
	goto L50
L52:
	;
	v205 = v200
	v206 = v201
	v207 = v202
	goto L62
L53:
	;
	if (v169|v54)&int32(3) != 0 {
		v200 = v169
		v201 = v54
		v202 = v105
		goto L52
	} else {
		goto L56
	}
L54:
	;
	v193 = v169
	v194 = v54
	v195 = v105
	goto L55
L55:
	;
	if v195 == int32(0) {
		goto L51
	} else {
		goto L61
	}
L56:
	;
	v177 = v169
	v178 = v54
	v179 = v105
	goto L57
L57:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	if v182 != v183 {
		v200 = v177
		v201 = v178
		v202 = v179
		goto L52
	} else {
		goto L59
	}
L58:
	;
	v193 = v188
	v194 = v186
	v195 = v190
	goto L55
L59:
	;
	v185 = int32(4)
	v186 = v178 + v185
	v188 = v177 + v185
	v190 = v179 - v185
	if base.Ui32(int32(3)) < base.Ui32(v190) {
		v177 = v188
		v178 = v186
		v179 = v190
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v200 = v193
	v201 = v194
	v202 = v195
	goto L52
L62:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	if v210 == v211 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v231 = v210 - v211
	goto L50
L64:
	;
	v213 = int32(1)
	v218 = v207 - v213
	if v218 != 0 {
		v205 = v205 + v213
		v206 = v206 + v213
		v207 = v218
		goto L62
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	goto L63
L67:
	;
	goto L51
L68:
	;
	v234 = int32(0)
	if v231 != 0 {
		v714 = v234
		goto L6
	} else {
		goto L69
	}
L69:
	;
	v236 = F_index_open(m, l0, int32(1))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	if v93 <= int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	F_relation_close(m, v236, int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L187
	}
L72:
	;
	v378 = int32(0)
	v379 = m.G0
	v381 = v379 - int32(32)
	m.G0 = v381
	if v58|v364 == v378 {
		goto L111
	} else {
		goto L112
	}
L73:
	;
	v240 = F_palloc(m, v105)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v247 = v234
	goto L77
L76:
	;
	v364 = v240
	goto L72
L77:
	;
	v259 = v247 << (uint(int32(2)) % 32)
	v260 = v56 + v259
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v262 = F_get_opclass_input_type(m, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L81
	}
L78:
	;
	v334 = F_palloc(m, v105)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L105
	}
L79:
	;
	v331 = v247 + int32(1)
	if v331 != v93 {
		v247 = v331
		goto L77
	} else {
		goto L104
	}
L80:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v236)+52))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v316+v317<<(uint(int32(4))%32)+v247*int32(100))+88))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v259+v52)))
	if v324 == v326 {
		goto L79
	} else {
		goto L103
	}
L81:
	;
	if v262 == int32(2283) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v267 = F_get_opclass_input_type(m, v266)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	if v267 == int32(2277) {
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v272 = F_get_opclass_input_type(m, v271)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v272 == int32(2776) {
		goto L80
	} else {
		goto L86
	}
L86:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v277 = F_get_opclass_input_type(m, v276)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	if v277 == int32(3500) {
		goto L80
	} else {
		goto L88
	}
L88:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v282 = F_get_opclass_input_type(m, v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	if v282 == int32(3831) {
		goto L80
	} else {
		goto L90
	}
L90:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v287 = F_get_opclass_input_type(m, v286)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	if v287 == int32(4537) {
		goto L80
	} else {
		goto L92
	}
L92:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v292 = F_get_opclass_input_type(m, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	if v292 == int32(5077) {
		goto L80
	} else {
		goto L94
	}
L94:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v297 = F_get_opclass_input_type(m, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	if v297 == int32(5078) {
		goto L80
	} else {
		goto L96
	}
L96:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v302 = F_get_opclass_input_type(m, v301)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	if v302 == int32(5079) {
		goto L80
	} else {
		goto L98
	}
L98:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v307 = F_get_opclass_input_type(m, v306)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	if v307 == int32(5080) {
		goto L80
	} else {
		goto L100
	}
L100:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v312 = F_get_opclass_input_type(m, v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	if v312 != int32(4538) {
		goto L79
	} else {
		goto L102
	}
L102:
	;
	goto L80
L103:
	;
	v658 = int32(0)
	goto L71
L104:
	;
	goto L78
L105:
	;
	v341 = int32(0)
	goto L106
L106:
	;
	v356 = v341 + int32(1)
	v358 = F_get_attoptions(m, l0, base.I32_extend16_s(v356))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L108
	}
L107:
	;
	v364 = v334
	goto L72
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334+v341<<(uint(int32(2))%32)))) = v358
	if v356 != v93 {
		v341 = v356
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	m.G0 = v381 + int32(32)
	F_pfree(m, v364)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L134
	}
L111:
	;
	v459 = int32(1)
	goto L110
L112:
	;
	goto L113
L113:
	;
	F_fmgr_info(m, int32(744), v381+int32(4))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	if v93 <= int32(0) {
		v459 = int32(1)
		goto L110
	} else {
		goto L115
	}
L115:
	;
	v399 = v378
	goto L116
L116:
	;
	if v364 != 0 {
		goto L122
	} else {
		goto L123
	}
L117:
	;
	v459 = v452
	goto L110
L118:
	;
	v452 = int32(1)
	v454 = v399 + v452
	if v454 != v93 {
		v399 = v454
		goto L116
	} else {
		goto L133
	}
L119:
	;
	v459 = int32(0)
	goto L110
L120:
	;
	if v414 == int32(0) {
		goto L118
	} else {
		goto L132
	}
L121:
	;
	if v440 == int32(0) {
		goto L118
	} else {
		goto L131
	}
L122:
	;
	v412 = v399 << (uint(int32(2)) % 32)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v364+v412)))
	if v58 == int32(0) {
		goto L120
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	if v58 == int32(0) {
		goto L118
	} else {
		goto L130
	}
L125:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v58+v412)))
	if v414 == int32(0) {
		v440 = v418
		goto L121
	} else {
		goto L126
	}
L126:
	;
	v421 = int32(0)
	if v418 == v421 {
		v459 = v421
		goto L110
	} else {
		goto L127
	}
L127:
	;
	v427 = F_FunctionCall2Coll(m, v381+int32(4), int32(950), v414, v418)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	if v427 == int32(0) {
		v459 = v421
		goto L110
	} else {
		goto L129
	}
L129:
	;
	goto L118
L130:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v58+v399<<(uint(int32(2))%32))))
	v440 = v436
	goto L121
L131:
	;
	goto L119
L132:
	;
	goto L119
L133:
	;
	goto L117
L134:
	;
	if v459 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v658 = int32(0)
	goto L71
L136:
	;
	goto L137
L137:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v48)+92))
	if v480 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v658 = int32(1)
	goto L71
L139:
	;
	goto L140
L140:
	;
	F_RelationGetExclusionInfo(m, v236, v19+int32(44), v19+int32(40), v19+int32(36))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v48)+92))
	if base.Ui32(int32(4)) <= base.Ui32(v105) {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	v557 = base.B2i32(v555 == int32(0))
	if v555 != 0 {
		v658 = v557
		goto L71
	} else {
		goto L160
	}
L143:
	;
	v555 = int32(0)
	goto L142
L144:
	;
	v529 = v524
	v530 = v525
	v531 = v526
	goto L154
L145:
	;
	if (v492|v493)&int32(3) != 0 {
		v524 = v492
		v525 = v493
		v526 = v105
		goto L144
	} else {
		goto L148
	}
L146:
	;
	v517 = v492
	v518 = v493
	v519 = v105
	goto L147
L147:
	;
	if v519 == int32(0) {
		goto L143
	} else {
		goto L153
	}
L148:
	;
	v501 = v492
	v502 = v493
	v503 = v105
	goto L149
L149:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v501)))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v502)))
	if v506 != v507 {
		v524 = v501
		v525 = v502
		v526 = v503
		goto L144
	} else {
		goto L151
	}
L150:
	;
	v517 = v512
	v518 = v510
	v519 = v514
	goto L147
L151:
	;
	v509 = int32(4)
	v510 = v502 + v509
	v512 = v501 + v509
	v514 = v503 - v509
	if base.Ui32(int32(3)) < base.Ui32(v514) {
		v501 = v512
		v502 = v510
		v503 = v514
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	v524 = v517
	v525 = v518
	v526 = v519
	goto L144
L154:
	;
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529))))
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530))))
	if v534 == v535 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v555 = v534 - v535
	goto L142
L156:
	;
	v537 = int32(1)
	v542 = v531 - v537
	if v542 != 0 {
		v529 = v529 + v537
		v530 = v530 + v537
		v531 = v542
		goto L154
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	goto L155
L159:
	;
	goto L143
L160:
	;
	if v93 <= int32(0) {
		v658 = v557
		goto L71
	} else {
		goto L161
	}
L161:
	;
	v563 = int32(0)
	goto L162
L162:
	;
	v578 = v563 << (uint(int32(2)) % 32)
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v48)+92))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v578+v579)))
	F_op_input_types(m, v581, v19+int32(32), v19+int32(28))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L164
	}
L163:
	;
	v658 = v649
	goto L71
L164:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	if base.Ui32(v588-int32(5077)) < base.Ui32(int32(2)) {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	v649 = int32(1)
	v651 = v563 + v649
	if v651 != v93 {
		v563 = v651
		goto L162
	} else {
		goto L186
	}
L166:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v236)+52))
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v635)))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v635+v636<<(uint(int32(4))%32)+v563*int32(100))+88))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v578+v52)))
	if v643 == v645 {
		goto L165
	} else {
		goto L185
	}
L167:
	;
	if v588 == int32(4537) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	if v588 == int32(3831) {
		goto L166
	} else {
		goto L169
	}
L169:
	;
	if v588 == int32(3500) {
		goto L166
	} else {
		goto L170
	}
L170:
	;
	if v588 == int32(2776) {
		goto L166
	} else {
		goto L171
	}
L171:
	;
	if v588 == int32(2283) {
		goto L166
	} else {
		goto L172
	}
L172:
	;
	if v588 == int32(2277) {
		goto L166
	} else {
		goto L173
	}
L173:
	;
	if v588 == int32(4538) {
		goto L166
	} else {
		goto L174
	}
L174:
	;
	if base.Ui32(v588-int32(5079)) < base.Ui32(int32(2)) {
		goto L166
	} else {
		goto L175
	}
L175:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	if base.Ui32(v611-int32(5077)) < base.Ui32(int32(2)) {
		goto L166
	} else {
		goto L176
	}
L176:
	;
	if v611 == int32(4537) {
		goto L166
	} else {
		goto L177
	}
L177:
	;
	if v611 == int32(3831) {
		goto L166
	} else {
		goto L178
	}
L178:
	;
	if v611 == int32(3500) {
		goto L166
	} else {
		goto L179
	}
L179:
	;
	if v611 == int32(2776) {
		goto L166
	} else {
		goto L180
	}
L180:
	;
	if v611 == int32(2283) {
		goto L166
	} else {
		goto L181
	}
L181:
	;
	if v611 == int32(2277) {
		goto L166
	} else {
		goto L182
	}
L182:
	;
	if v611 == int32(4538) {
		goto L166
	} else {
		goto L183
	}
L183:
	;
	if base.Ui32(int32(1)) < base.Ui32(v611-int32(5079)) {
		goto L165
	} else {
		goto L184
	}
L184:
	;
	goto L166
L185:
	;
	v658 = int32(0)
	goto L71
L186:
	;
	goto L163
L187:
	;
	v714 = v658
	goto L6
L188:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l1
	F_errmsg(m, int32(78403), v19)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(516405), int32(227), int32(408408))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = l0
	F_errmsg_internal(m, int32(42668), v19+int32(16))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	F_errfinish(m, int32(516405), int32(263), int32(408408))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L195:
	;
	v714 = int32(0)
	goto L6
}
func F_ExecIndexBuildScanKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v42 int32
	_ = v42
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
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
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
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
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
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
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
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
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
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
	var v446 int32
	_ = v446
	var v449 int64
	_ = v449
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v618 int32
	_ = v618
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v870 int32
	_ = v870
	var v875 int32
	_ = v875
	var v898 int32
	_ = v898
	var v903 int32
	_ = v903
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1010 int32
	_ = v1010
	v11 = int32(0)
	v42 = m.G0
	v44 = v42 - int32(48)
	m.G0 = v44
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v48 = v46
	goto L3
L2:
	;
	v48 = int32(0)
	goto L3
L3:
	;
	v51 = F_palloc(m, v48*int32(48))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v57 = F_palloc0(m, v48*int32(24))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if l2 != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L4
	} else {
		goto L210
	}
L8:
	;
	m.G0 = v44 + int32(48)
	return
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v936
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v932
	goto L8
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v51
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v48
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v834
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v833
	if l8 == int32(0) {
		goto L7
	} else {
		goto L209
	}
L11:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v59 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v870 = v53
	v875 = v54
	goto L13
L13:
	;
	F_pfree(m, v57)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L4
	} else {
		goto L207
	}
L14:
	;
	if v836 != 0 {
		goto L10
	} else {
		goto L206
	}
L15:
	;
	v833 = v53
	v834 = v54
	v836 = v11
	goto L14
L16:
	;
	goto L17
L17:
	;
	if l3 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v64 = int32(256)
	goto L20
L19:
	;
	v64 = int32(0)
	goto L20
L20:
	;
	v68 = v51 ^ int32(-1)
	v70 = v51 + int32(48)
	v72 = v51 + int32(4)
	if base.Ui32(v72) < base.Ui32(v70) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v74 = v70
	goto L23
L22:
	;
	v74 = v72
	goto L23
L23:
	;
	v87 = v68 + v74
	v90 = v53
	v94 = v53
	v95 = v54
	v97 = v11
	v98 = v11
	v113 = v51
	goto L33
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L4
	} else {
		goto L203
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L4
	} else {
		goto L200
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L4
	} else {
		goto L197
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L4
	} else {
		goto L194
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L4
	} else {
		goto L191
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L4
	} else {
		goto L188
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L4
	} else {
		goto L185
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L4
	} else {
		goto L182
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L4
	} else {
		goto L179
	}
L33:
	;
	v119 = v51 + v98*int32(48)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v120)+10)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122+v98<<(uint(int32(2))%32))))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	switch v127 - int32(17) {
	case 0:
		goto L40
	default:
		goto L24
	case 3:
		goto L38
	case 20:
		goto L39
	case 35:
		goto L37
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L4
	} else {
		goto L176
	}
L35:
	;
	goto L34
L36:
	;
	v661 = v98 + int32(1)
	v663 = v661 * int32(48)
	v664 = v70 + v663
	v665 = v663 + v72
	if base.Ui32(v665) < base.Ui32(v664) {
		goto L172
	} else {
		goto L173
	}
L37:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v578)))
	if v579 == int32(27) {
		goto L160
	} else {
		goto L161
	}
L38:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v126)+28))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v468)+12))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v470)))
	if v471 == int32(27) {
		goto L127
	} else {
		goto L128
	}
L39:
	;
	v222 = int32(0)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	if v228 != 0 {
		goto L66
	} else {
		goto L67
	}
L40:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+28))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v133 == int32(27) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v138 = v136
	v139 = v137
	goto L43
L42:
	;
	v138 = v132
	v139 = v133
	goto L43
L43:
	;
	if v139 != int32(6) {
		goto L35
	} else {
		goto L44
	}
L44:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v142 != int32(-3) {
		goto L35
	} else {
		goto L45
	}
L45:
	;
	v145 = int32(*(*int16)(unsafe.Add(mBase, uint32(v138)+8)))
	if v145 <= int32(0) {
		goto L32
	} else {
		goto L46
	}
L46:
	;
	if v121 < v145 {
		goto L32
	} else {
		goto L47
	}
L47:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v151+v145<<(uint(int32(2))%32)-int32(4))))
	F_get_op_opfamily_properties(m, v150, v157, l3, v44+int32(44), v44+int32(40), v44+int32(36))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v126)+28))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+12))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	if v169 == int32(27) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+44)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v44)+36))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v126)+24))
	F_ScanKeyEntryInitialize(m, v119, v212, v145, v217, v218, v219, v149, v216)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L4
	} else {
		goto L65
	}
L50:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v174 = v172
	v175 = v173
	goto L52
L51:
	;
	v174 = v168
	v175 = v169
	goto L52
L52:
	;
	if v175 == int32(7) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+24)))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v174)+20))
	v212 = v64 | v178
	v213 = v90
	v214 = v94
	v215 = v95
	v216 = v180
	goto L49
L54:
	;
	goto L55
L55:
	;
	if v94 < v90 {
		v194 = v90
		v195 = v95
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v198 = v195 + v94*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v198))) = v119
	v200 = F_ExecInitExpr(m, v174, l0)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L63
	}
L57:
	;
	if v90 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v186 = F_palloc(m, int32(96))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v190 = F_repalloc(m, v95, v90*int32(24))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L62
	}
L61:
	;
	v194 = int32(8)
	v195 = v186
	goto L56
L62:
	;
	v194 = v90 << (uint(int32(1)) % 32)
	v195 = v190
	goto L56
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+4)) = v200
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v44)+36))
	v204 = F_get_typstorage(m, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v198)+8)) = uint8(base.B2i32(v204 != int32(112)))
	v212 = v64
	v213 = v194
	v214 = v94 + int32(1)
	v215 = v195
	v216 = int32(0)
	goto L49
L65:
	;
	v633 = v213
	v637 = v214
	v638 = v215
	v640 = v97
	goto L36
L66:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	v232 = v229 * int32(48)
	goto L68
L67:
	;
	v232 = v222
	goto L68
L68:
	;
	v233 = F_palloc(m, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v126)+24))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
	v249 = v222
	v253 = v90
	v257 = v94
	v258 = v95
	goto L71
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = int32(4)
	v463 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v233)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v119)+4)) = uint16(v463)
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+44)) = v233
	*(*uint16)(unsafe.Add(mBase, uint32(v119)+6)) = uint16(v465)
	v633 = v253
	v637 = v257
	v638 = v258
	v640 = v97
	goto L36
L71:
	;
	v280 = int32(0)
	if v238 == v280 {
		v291 = v280
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v449 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v119)+4)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v119)+44)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v119)+36)) = v449
	*(*int64)(unsafe.Add(mBase, uint32(v119)+28)) = v449
	*(*int64)(unsafe.Add(mBase, uint32(v119)+20)) = v449
	*(*int64)(unsafe.Add(mBase, uint32(v119)+12)) = v449
	goto L70
L73:
	;
	if v237 == int32(0) {
		v300 = v280
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	if v285 <= v249 {
		v291 = int32(0)
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v291 = v287 + v249<<(uint(int32(2))%32)
	goto L73
L76:
	;
	v301 = int32(0)
	if v236 == v301 {
		v312 = v301
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v294 <= v249 {
		v300 = v280
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v237)+12))
	v300 = v296 + v249<<(uint(int32(2))%32)
	goto L76
L79:
	;
	if v235 == int32(0) {
		v321 = v301
		goto L82
	} else {
		goto L83
	}
L80:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	if v306 <= v249 {
		v312 = int32(0)
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v236)+12))
	v312 = v308 + v249<<(uint(int32(2))%32)
	goto L79
L82:
	;
	v324 = v233 + v249*int32(48)
	if v291 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	if v315 <= v249 {
		v321 = v301
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v235)+12))
	v321 = v317 + v249<<(uint(int32(2))%32)
	goto L82
L85:
	;
	goto L72
L86:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	if v344 == int32(27) {
		goto L95
	} else {
		goto L96
	}
L87:
	;
	v332 = v324 - int32(48)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	*(*int32)(unsafe.Add(mBase, uint32(v332))) = v333 | int32(16)
	if v51&int32(3) != 0 {
		goto L85
	} else {
		goto L92
	}
L88:
	;
	if v300 == int32(0) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	if v312 == int32(0) {
		goto L87
	} else {
		goto L90
	}
L90:
	;
	if v321 != 0 {
		goto L86
	} else {
		goto L91
	}
L91:
	;
	goto L87
L92:
	;
	if base.Ui32(v119+int32(48)) <= base.Ui32(v119) {
		goto L70
	} else {
		goto L93
	}
L93:
	;
	v342 = F__emscripten_memset_bulkmem(m, v113, base.I32_extend8_s(int32(0)), v87&int32(-4)+int32(4))
	mBase = m.M
	goto L94
L94:
	;
	goto L70
L95:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v343)+4))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	v349 = v347
	v350 = v348
	goto L97
L96:
	;
	v349 = v343
	v350 = v344
	goto L97
L97:
	;
	if v350 != int32(6) {
		goto L31
	} else {
		goto L98
	}
L98:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	if v353 != int32(-3) {
		goto L31
	} else {
		goto L99
	}
L99:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+10)))
	if v357 != int32(1) {
		goto L30
	} else {
		goto L100
	}
L100:
	;
	v360 = int32(*(*int16)(unsafe.Add(mBase, uint32(v349)+8)))
	if v360 <= int32(0) {
		goto L30
	} else {
		goto L101
	}
L101:
	;
	if v121 < v360 {
		goto L30
	} else {
		goto L102
	}
L102:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v367+v360<<(uint(int32(2))%32)-int32(4))))
	F_get_op_opfamily_properties(m, v366, v373, l3, v44+int32(44), v44+int32(40), v44+int32(36))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v44)+44))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v382 != v383 {
		goto L29
	} else {
		goto L104
	}
L104:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v44)+36))
	v388 = F_get_opfamily_proc(m, v373, v385, v386, int32(1))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	if v388 == int32(0) {
		goto L28
	} else {
		goto L106
	}
L106:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	if v392 == int32(27) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v443 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+44)))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v44)+36))
	F_ScanKeyEntryInitialize(m, v324, v438, v360, v443, v444, v364, v388, v442)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L4
	} else {
		goto L126
	}
L108:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	v397 = v395
	v398 = v396
	goto L110
L109:
	;
	v397 = v365
	v398 = v392
	goto L110
L110:
	;
	if v398 == int32(7) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+24)))
	if v403 != 0 {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	goto L113
L113:
	;
	if v257 < v253 {
		v419 = v253
		v420 = v258
		goto L117
	} else {
		goto L118
	}
L114:
	;
	v404 = int32(9)
	goto L116
L115:
	;
	v404 = int32(8)
	goto L116
L116:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v397)+20))
	v438 = v404
	v439 = v253
	v440 = v257
	v441 = v258
	v442 = v405
	goto L107
L117:
	;
	v423 = v420 + v257*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v423))) = v324
	v425 = F_ExecInitExpr(m, v397, l0)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L4
	} else {
		goto L124
	}
L118:
	;
	if v253 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v411 = F_palloc(m, int32(96))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L4
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v415 = F_repalloc(m, v258, v253*int32(24))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L4
	} else {
		goto L123
	}
L122:
	;
	v419 = int32(8)
	v420 = v411
	goto L117
L123:
	;
	v419 = v253 << (uint(int32(1)) % 32)
	v420 = v415
	goto L117
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v423)+4)) = v425
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v44)+36))
	v429 = F_get_typstorage(m, v428)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v423)+8)) = uint8(base.B2i32(v429 != int32(112)))
	v438 = int32(8)
	v439 = v419
	v440 = v257 + int32(1)
	v441 = v420
	v442 = int32(0)
	goto L107
L126:
	;
	v249 = v249 + int32(1)
	v253 = v439
	v257 = v440
	v258 = v441
	goto L71
L127:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v474)))
	v476 = v474
	v477 = v475
	goto L129
L128:
	;
	v476 = v470
	v477 = v471
	goto L129
L129:
	;
	if v477 != int32(6) {
		goto L27
	} else {
		goto L130
	}
L130:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	if v480 != int32(-3) {
		goto L27
	} else {
		goto L131
	}
L131:
	;
	v483 = int32(*(*int16)(unsafe.Add(mBase, uint32(v476)+8)))
	if v483 <= int32(0) {
		goto L26
	} else {
		goto L132
	}
L132:
	;
	if v121 < v483 {
		goto L26
	} else {
		goto L133
	}
L133:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v489+v483<<(uint(int32(2))%32)-int32(4))))
	F_get_op_opfamily_properties(m, v488, v495, l3, v44+int32(44), v44+int32(40), v44+int32(36))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v126)+28))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v504)+12))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v505)+4))
	if v506 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+19)))
	if v516 == int32(1) {
		goto L141
	} else {
		goto L142
	}
L136:
	;
	v514 = int32(0)
	goto L135
L137:
	;
	goto L138
L138:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v506)))
	if v510 != int32(27) {
		v514 = v506
		goto L135
	} else {
		goto L139
	}
L139:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v506)+4))
	v514 = v513
	goto L135
L140:
	;
	v573 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+44)))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v44)+36))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v126)+24))
	F_ScanKeyEntryInitialize(m, v119, v568, v483, v573, v574, v575, v487, v572)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L4
	} else {
		goto L159
	}
L141:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v514)))
	if v519 == int32(7) {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	goto L143
L143:
	;
	v557 = v57 + v97*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v557))) = v119
	v559 = F_ExecInitExpr(m, v514, l0)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L4
	} else {
		goto L158
	}
L144:
	;
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+24)))
	if v524 != 0 {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	goto L146
L146:
	;
	if v94 < v90 {
		v540 = v90
		v541 = v95
		goto L150
	} else {
		goto L151
	}
L147:
	;
	v525 = int32(33)
	goto L149
L148:
	;
	v525 = int32(32)
	goto L149
L149:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v514)+20))
	v567 = v90
	v568 = v525
	v569 = v94
	v570 = v95
	v571 = v97
	v572 = v526
	goto L140
L150:
	;
	v544 = v541 + v94*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v544))) = v119
	v546 = F_ExecInitExpr(m, v514, l0)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L4
	} else {
		goto L157
	}
L151:
	;
	if v90 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v532 = F_palloc(m, int32(96))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L4
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v536 = F_repalloc(m, v95, v90*int32(24))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L4
	} else {
		goto L156
	}
L155:
	;
	v540 = int32(8)
	v541 = v532
	goto L150
L156:
	;
	v540 = v90 << (uint(int32(1)) % 32)
	v541 = v536
	goto L150
L157:
	;
	v548 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v544)+8)) = uint8(v548)
	*(*int32)(unsafe.Add(mBase, uint32(v544)+4)) = v546
	v567 = v540
	v568 = int32(32)
	v569 = v94 + v548
	v570 = v541
	v571 = v97
	v572 = int32(0)
	goto L140
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v557)+4)) = v559
	v564 = int32(0)
	v567 = v90
	v568 = v564
	v569 = v94
	v570 = v95
	v571 = v97 + int32(1)
	v572 = v564
	goto L140
L159:
	;
	v633 = v567
	v637 = v569
	v638 = v570
	v640 = v571
	goto L36
L160:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v578)+4))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v582)))
	v584 = v582
	v585 = v583
	goto L162
L161:
	;
	v584 = v578
	v585 = v579
	goto L162
L162:
	;
	if v585 != int32(6) {
		goto L25
	} else {
		goto L163
	}
L163:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v584)+4))
	if v588 != int32(-3) {
		goto L25
	} else {
		goto L164
	}
L164:
	;
	v591 = int32(*(*int16)(unsafe.Add(mBase, uint32(v584)+8)))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	switch v593 {
	case 0:
		v611 = int32(65)
		goto L165
	case 1:
		goto L166
	default:
		goto L167
	}
L165:
	;
	v612 = int32(0)
	F_ScanKeyEntryInitialize(m, v119, v611, v591, v612, v612, v612, v612, v612)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L4
	} else {
		goto L171
	}
L166:
	;
	v611 = int32(129)
	goto L165
L167:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+32)) = v598
	F_errmsg_internal(m, int32(504938), v44+int32(32))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(519188), int32(1607), int32(120441))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L4
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	v633 = v90
	v637 = v94
	v638 = v95
	v640 = v97
	goto L36
L172:
	;
	v667 = v664
	goto L174
L173:
	;
	v667 = v665
	goto L174
L174:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v661 < v673 {
		v87 = v667 + (v661*int32(-48) + v68)
		v90 = v633
		v94 = v637
		v95 = v638
		v97 = v640
		v98 = v661
		v113 = v663 + v51
		goto L33
	} else {
		goto L175
	}
L175:
	;
	v833 = v637
	v834 = v638
	v836 = v640
	goto L14
L176:
	;
	F_errmsg_internal(m, int32(432827), int32(0))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(519188), int32(1235), int32(120441))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	F_errmsg_internal(m, int32(279318), int32(0))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	F_errfinish(m, int32(519188), int32(1239), int32(120441))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L4
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
	F_errmsg_internal(m, int32(432827), int32(0))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	F_errfinish(m, int32(519188), int32(1352), int32(120441))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L185:
	;
	F_errmsg_internal(m, int32(279344), int32(0))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L4
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(519188), int32(1362), int32(120441))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L4
	} else {
		goto L187
	}
L187:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L188:
	;
	F_errmsg_internal(m, int32(219490), int32(0))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L4
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(519188), int32(1371), int32(120441))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L4
	} else {
		goto L190
	}
L190:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+28)) = v373
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = v748
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v44)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v750
	F_errmsg_internal(m, int32(42270), v44+int32(16))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(519188), int32(1379), int32(120441))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L4
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
	F_errmsg_internal(m, int32(432827), int32(0))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L4
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(519188), int32(1476), int32(120441))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L4
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	F_errmsg_internal(m, int32(279318), int32(0))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L4
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(519188), int32(1480), int32(120441))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L4
	} else {
		goto L199
	}
L199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L200:
	;
	F_errmsg_internal(m, int32(22529), int32(0))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L4
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(519188), int32(1590), int32(120441))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L4
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L203:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v805
	F_errmsg_internal(m, int32(506337), v44)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L4
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(519188), int32(1623), int32(120441))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L4
	} else {
		goto L205
	}
L205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L206:
	;
	v870 = v833
	v875 = v834
	goto L13
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v51
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v48
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v875
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v870
	v903 = int32(0)
	if l8 != 0 {
		v932 = v903
		v936 = v903
		goto L9
	} else {
		goto L208
	}
L208:
	;
	goto L8
L209:
	;
	v932 = v836
	v936 = v57
	goto L9
L210:
	;
	F_errmsg_internal(m, int32(459232), int32(0))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L4
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(519188), int32(1648), int32(120441))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L4
	} else {
		goto L212
	}
L212:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_IndexNextWithReorder(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 float64
	_ = v201
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
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
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
	var v276 int32
	_ = v276
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
	var v287 int32
	_ = v287
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
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
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v537 int32
	_ = v537
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v20 != 0 {
		v44 = v20
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v55 = int32(0)
	goto L10
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v29 = F_index_beginscan(m, v21, v22, v24, l0+int32(168), v27, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v29
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v34 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)))
	if v35 != int32(1) {
		v44 = v29
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	F_index_rescan(m, v29, v38, v39, v40, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	v44 = v29
	goto L1
L10:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v59 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	m.G0 = v16 + int32(16)
	return v18
L12:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L3
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v63 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L14
L16:
	;
	goto L11
L17:
	;
	v158 = F_index_getnext_slot(m, v44, int32(1), v18)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L3
	} else {
		goto L40
	}
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+184)))
	if v65 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+184)))
	if v137 != int32(1) {
		v154 = v55
		goto L17
	} else {
		goto L37
	}
L21:
	;
	v132 = F_reorderqueue_pop(m, l0)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L3
	} else {
		goto L35
	}
L22:
	;
	v66 = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v67 <= v66 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	v75 = v66
	goto L24
L24:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+v70))))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+v72))))
	if v90 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	if int32(0) < v116 {
		v154 = v64
		goto L17
	} else {
		goto L34
	}
L26:
	;
	goto L25
L27:
	;
	v116 = (v88 ^ int32(1)) & int32(255)
	goto L26
L28:
	;
	goto L29
L29:
	;
	if v88&int32(1) != 0 {
		goto L21
	} else {
		goto L30
	}
L30:
	;
	v100 = v75 << (uint(int32(2)) % 32)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v73+v100)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100+v71)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v108 = v105 + v75*int32(36)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	v110 = m.T0[v109].(func(*base.Module, int32, int32, int32) int32)(m, v102, v104, v108)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	if v110 != 0 {
		v116 = v110
		goto L26
	} else {
		goto L32
	}
L32:
	;
	v113 = v75 + int32(1)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v113 < v114 {
		v75 = v113
		goto L24
	} else {
		goto L33
	}
L33:
	;
	goto L21
L34:
	;
	goto L21
L35:
	;
	F_ExecForceStoreHeapTuple(m, v132, v18, int32(1))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	goto L16
L37:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	m.T0[v141].(func(*base.Module, int32))(m, v18)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	goto L16
L39:
	;
	v537 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+184)) = uint8(v537)
	v55 = v154
	goto L10
L40:
	;
	if v158 == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	goto L42
L42:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+72)))
	if v175 != int32(1) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+84)))
	if v214 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L44:
	;
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v18
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v179 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	F_MemoryContextReset(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L3
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v185 = int32(4553888)
	v186 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v188
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
	v193 = m.T0[v192].(func(*base.Module, int32, int32, int32) int32)(m, v179, v19, v16+int32(15))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L3
	} else {
		goto L50
	}
L49:
	;
	goto L44
L50:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v186
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	F_MemoryContextReset(m, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	if v193 != 0 {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v200 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v201 = *(*float64)(unsafe.Add(mBase, uint32(v200)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v200)+248)) = base.F64_add(v201, float64(1))
	goto L55
L54:
	;
	goto L55
L55:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v206 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L3
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v210 = F_index_getnext_slot(m, v44, int32(1), v18)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L3
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	if v210 != 0 {
		goto L42
	} else {
		goto L61
	}
L61:
	;
	goto L39
L62:
	;
	v432 = int32(4553888)
	v433 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+100))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v437
	v440 = F_palloc(m, int32(24))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L3
	} else {
		goto L106
	}
L63:
	;
	if v154 == int32(0) {
		goto L16
	} else {
		goto L93
	}
L64:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v300&int32(255) == int32(0) {
		v422 = v347
		v427 = v346
		goto L62
	} else {
		goto L92
	}
L65:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v356 = v345
	v361 = v344
	goto L63
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v18
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	F_MemoryContextReset(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L3
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	v356 = v343
	v361 = v342
	goto L63
L69:
	;
	v221 = int32(4553888)
	v222 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v224
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v226 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v276 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v222
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v281 <= v276 {
		v356 = v279
		v361 = v280
		goto L63
	} else {
		goto L77
	}
L71:
	;
	v229 = int32(0)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v230 <= v229 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v234 = v229
	goto L73
L73:
	;
	v247 = v234 << (uint(int32(2)) % 32)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v247+v248)))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v250)+20))
	v254 = m.T0[v253].(func(*base.Module, int32, int32, int32) int32)(m, v250, v19, v251+v234)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L3
	} else {
		goto L75
	}
L74:
	;
	goto L70
L75:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v256+v247))) = v254
	v260 = v234 + int32(1)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v260 < v261 {
		v234 = v260
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	v287 = v276
	goto L79
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L3
	} else {
		goto L89
	}
L79:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287+v284))))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287+v279))))
	if v302 != 0 {
		goto L64
	} else {
		goto L81
	}
L80:
	;
	if v316 < int32(0) {
		goto L78
	} else {
		goto L88
	}
L81:
	;
	if v300&int32(1) != 0 {
		goto L78
	} else {
		goto L82
	}
L82:
	;
	v306 = v287 << (uint(int32(2)) % 32)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v280+v306)))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v306+v285)))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v314 = v311 + v287*int32(36)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+16))
	v316 = m.T0[v315].(func(*base.Module, int32, int32, int32) int32)(m, v308, v310, v314)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L3
	} else {
		goto L83
	}
L83:
	;
	if v316 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v321 = v287 + int32(1)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v322 <= v321 {
		goto L65
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	goto L80
L87:
	;
	v287 = v321
	goto L79
L88:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v422 = v327
	v427 = v326
	goto L62
L89:
	;
	F_errmsg_internal(m, int32(237738), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L3
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(519188), int32(311), int32(237586))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L3
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
	v356 = v347
	v361 = v346
	goto L63
L93:
	;
	v367 = int32(0)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v368 <= v367 {
		goto L16
	} else {
		goto L94
	}
L94:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v154)+20))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
	v374 = v367
	goto L95
L95:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374+v371))))
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374+v356))))
	if v389 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	if v415 <= int32(0) {
		goto L16
	} else {
		goto L105
	}
L97:
	;
	goto L96
L98:
	;
	v415 = (v387 ^ int32(1)) & int32(255)
	goto L97
L99:
	;
	goto L100
L100:
	;
	if v387&int32(1) != 0 {
		goto L16
	} else {
		goto L101
	}
L101:
	;
	v399 = v374 << (uint(int32(2)) % 32)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v361+v399)))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v399+v372)))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v407 = v404 + v374*int32(36)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)+16))
	v409 = m.T0[v408].(func(*base.Module, int32, int32, int32) int32)(m, v401, v403, v407)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L3
	} else {
		goto L102
	}
L102:
	;
	if v409 != 0 {
		v415 = v409
		goto L97
	} else {
		goto L103
	}
L103:
	;
	v412 = v374 + int32(1)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v412 < v413 {
		v374 = v412
		goto L95
	} else {
		goto L104
	}
L104:
	;
	goto L16
L105:
	;
	v422 = v356
	v427 = v361
	goto L62
L106:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)+44))
	v444 = m.T0[v443].(func(*base.Module, int32) int32)(m, v18)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L3
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+12)) = v444
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v434)+16))
	v450 = F_palloc(m, v447<<(uint(int32(2))%32))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L3
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+16)) = v450
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v434)+16))
	v454 = F_palloc(m, v453)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L3
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+20)) = v454
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if int32(0) < v457 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v461 = int32(0)
	goto L113
L111:
	;
	goto L112
L112:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	F_pairingheap_add(m, v519, v440)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L3
	} else {
		goto L120
	}
L113:
	;
	v473 = int32(0)
	v474 = v461 + v422
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
	if v475 == v473 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	goto L112
L115:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v427+v461<<(uint(int32(2))%32))))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482+v461))))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v489 = int32(*(*int16)(unsafe.Add(mBase, uint32(v485+v461<<(uint(int32(1))%32)))))
	v490 = F_datumCopy(m, v481, v484, v489)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L3
	} else {
		goto L118
	}
L116:
	;
	v492 = v473
	goto L117
L117:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v440)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v493+v461<<(uint(int32(2))%32)))) = v492
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v440)+20))
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
	*(*uint8)(unsafe.Add(mBase, uint32(v498+v461))) = uint8(v500)
	v503 = v461 + int32(1)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v503 < v504 {
		v461 = v503
		goto L113
	} else {
		goto L119
	}
L118:
	;
	v492 = v490
	goto L117
L119:
	;
	goto L114
L120:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v433
	v55 = v154
	goto L10
}
func F_IndexSetParentIndex(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
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
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
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
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v161 int64
	_ = v161
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(160)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v16 = F_relation_open(m, int32(2611), int32(3))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		F_ScanKeyInit(m, v11-int32(-64), int32(1), int32(3), int32(184), v13)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v27 = int32(3)
			F_ScanKeyInit(m, v11+int32(112), v27, v27, int32(65), int32(1))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v39 = F_systable_beginscan(m, v16, int32(2680), int32(1), int32(0), int32(2), v11-int32(-64))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					v41 = F_systable_getnext(m, v39)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						if v41 == int32(0) {
							if l1 == int32(0) {
								v64 = v3
								F_systable_endscan(m, v39)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									F_relation_close(m, v16, int32(3))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										if l1 != 0 {
											F_LockRelationOid(m, l1, int32(4))
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return
											} else {
												F_SetRelationHasSubclass(m, l1, int32(1))
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return
												} else {
													v78 = F_table_open(m, int32(1259), int32(3))
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return
													} else {
														v81 = F_SearchSysCacheLockedCopy1(m, int32(57), v13)
														mBase = m.M
														v82 = m.ExcPending
														if v82 != 0 {
															return
														} else {
															if v81 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
																	F_errmsg_internal(m, int32(49777), v11)
																	mBase = m.M
																	v180 = m.ExcPending
																	if v180 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(516405), int32(4584), int32(260026))
																		mBase = m.M
																		v185 = m.ExcPending
																		if v185 != 0 {
																			return
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	}
																}
															} else {
																v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+8)))
																*(*uint16)(unsafe.Add(mBase, uint32(v11)+56)) = uint16(v85)
																v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v87
																v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
																v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
																*(*uint8)(unsafe.Add(mBase, uint32(v89+v90)+131)) = uint8(base.B2i32(l1 != int32(0)))
																F_CatalogTupleUpdate(m, v78, v11+int32(52), v81)
																mBase = m.M
																v98 = m.ExcPending
																if v98 != 0 {
																	return
																} else {
																	F_UnlockTuple(m, v78, v11+int32(52), int32(7))
																	mBase = m.M
																	v103 = m.ExcPending
																	if v103 != 0 {
																		return
																	} else {
																		F_pfree(m, v81)
																		mBase = m.M
																		v105 = m.ExcPending
																		if v105 != 0 {
																			return
																		} else {
																			F_sequence_close(m, v78, int32(3))
																			mBase = m.M
																			v108 = m.ExcPending
																			if v108 != 0 {
																				return
																			} else {
																				if v64 != 0 {
																					if l1 != 0 {
																						v109 = int32(0)
																						*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v109
																						*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v13
																						v112 = int32(1259)
																						*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v112
																						*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v109
																						*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = l1
																						*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v112
																						*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v112
																						v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																						v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
																						*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v109
																						*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v122
																						F_recordDependencyOn(m, v11+int32(52), v11+int32(40), int32(80))
																						mBase = m.M
																						v132 = m.ExcPending
																						if v132 != 0 {
																							return
																						} else {
																							F_recordDependencyOn(m, v11+int32(52), v11+int32(28), int32(83))
																							mBase = m.M
																							v139 = m.ExcPending
																							if v139 != 0 {
																								return
																							} else {
																								F_CommandCounterIncrement(m)
																								mBase = m.M
																								v152 = m.ExcPending
																								if v152 != 0 {
																									return
																								} else {
																									m.G0 = v11 + int32(160)
																									return
																								}
																							}
																						}
																					} else {
																						v140 = int32(1259)
																						v143 = F_deleteDependencyRecordsForClass(m, v140, v13, v140, int32(80))
																						mBase = m.M
																						v144 = m.ExcPending
																						if v144 != 0 {
																							return
																						} else {
																							v145 = int32(1259)
																							v148 = F_deleteDependencyRecordsForClass(m, v145, v13, v145, int32(83))
																							mBase = m.M
																							v149 = m.ExcPending
																							if v149 != 0 {
																								return
																							} else {
																								F_CommandCounterIncrement(m)
																								mBase = m.M
																								v152 = m.ExcPending
																								if v152 != 0 {
																									return
																								} else {
																									m.G0 = v11 + int32(160)
																									return
																								}
																							}
																						}
																					}
																				} else {
																					m.G0 = v11 + int32(160)
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
											v78 = F_table_open(m, int32(1259), int32(3))
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return
											} else {
												v81 = F_SearchSysCacheLockedCopy1(m, int32(57), v13)
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return
												} else {
													if v81 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v176 = m.ExcPending
														if v176 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
															F_errmsg_internal(m, int32(49777), v11)
															mBase = m.M
															v180 = m.ExcPending
															if v180 != 0 {
																return
															} else {
																F_errfinish(m, int32(516405), int32(4584), int32(260026))
																mBase = m.M
																v185 = m.ExcPending
																if v185 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+8)))
														*(*uint16)(unsafe.Add(mBase, uint32(v11)+56)) = uint16(v85)
														v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v87
														v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
														v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
														*(*uint8)(unsafe.Add(mBase, uint32(v89+v90)+131)) = uint8(base.B2i32(l1 != int32(0)))
														F_CatalogTupleUpdate(m, v78, v11+int32(52), v81)
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return
														} else {
															F_UnlockTuple(m, v78, v11+int32(52), int32(7))
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return
															} else {
																F_pfree(m, v81)
																mBase = m.M
																v105 = m.ExcPending
																if v105 != 0 {
																	return
																} else {
																	F_sequence_close(m, v78, int32(3))
																	mBase = m.M
																	v108 = m.ExcPending
																	if v108 != 0 {
																		return
																	} else {
																		if v64 != 0 {
																			if l1 != 0 {
																				v109 = int32(0)
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v109
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v13
																				v112 = int32(1259)
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v112
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v109
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = l1
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v112
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v112
																				v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																				v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v109
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v122
																				F_recordDependencyOn(m, v11+int32(52), v11+int32(40), int32(80))
																				mBase = m.M
																				v132 = m.ExcPending
																				if v132 != 0 {
																					return
																				} else {
																					F_recordDependencyOn(m, v11+int32(52), v11+int32(28), int32(83))
																					mBase = m.M
																					v139 = m.ExcPending
																					if v139 != 0 {
																						return
																					} else {
																						F_CommandCounterIncrement(m)
																						mBase = m.M
																						v152 = m.ExcPending
																						if v152 != 0 {
																							return
																						} else {
																							m.G0 = v11 + int32(160)
																							return
																						}
																					}
																				}
																			} else {
																				v140 = int32(1259)
																				v143 = F_deleteDependencyRecordsForClass(m, v140, v13, v140, int32(80))
																				mBase = m.M
																				v144 = m.ExcPending
																				if v144 != 0 {
																					return
																				} else {
																					v145 = int32(1259)
																					v148 = F_deleteDependencyRecordsForClass(m, v145, v13, v145, int32(83))
																					mBase = m.M
																					v149 = m.ExcPending
																					if v149 != 0 {
																						return
																					} else {
																						F_CommandCounterIncrement(m)
																						mBase = m.M
																						v152 = m.ExcPending
																						if v152 != 0 {
																							return
																						} else {
																							m.G0 = v11 + int32(160)
																							return
																						}
																					}
																				}
																			}
																		} else {
																			m.G0 = v11 + int32(160)
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
								v47 = int32(1)
								F_StoreSingleInheritance(m, v13, l1, v47)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									v64 = v47
									F_systable_endscan(m, v39)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										F_relation_close(m, v16, int32(3))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											if l1 != 0 {
												F_LockRelationOid(m, l1, int32(4))
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return
												} else {
													F_SetRelationHasSubclass(m, l1, int32(1))
													mBase = m.M
													v75 = m.ExcPending
													if v75 != 0 {
														return
													} else {
														v78 = F_table_open(m, int32(1259), int32(3))
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return
														} else {
															v81 = F_SearchSysCacheLockedCopy1(m, int32(57), v13)
															mBase = m.M
															v82 = m.ExcPending
															if v82 != 0 {
																return
															} else {
																if v81 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v176 = m.ExcPending
																	if v176 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
																		F_errmsg_internal(m, int32(49777), v11)
																		mBase = m.M
																		v180 = m.ExcPending
																		if v180 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(516405), int32(4584), int32(260026))
																			mBase = m.M
																			v185 = m.ExcPending
																			if v185 != 0 {
																				return
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		}
																	}
																} else {
																	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+8)))
																	*(*uint16)(unsafe.Add(mBase, uint32(v11)+56)) = uint16(v85)
																	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v87
																	v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
																	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v89+v90)+131)) = uint8(base.B2i32(l1 != int32(0)))
																	F_CatalogTupleUpdate(m, v78, v11+int32(52), v81)
																	mBase = m.M
																	v98 = m.ExcPending
																	if v98 != 0 {
																		return
																	} else {
																		F_UnlockTuple(m, v78, v11+int32(52), int32(7))
																		mBase = m.M
																		v103 = m.ExcPending
																		if v103 != 0 {
																			return
																		} else {
																			F_pfree(m, v81)
																			mBase = m.M
																			v105 = m.ExcPending
																			if v105 != 0 {
																				return
																			} else {
																				F_sequence_close(m, v78, int32(3))
																				mBase = m.M
																				v108 = m.ExcPending
																				if v108 != 0 {
																					return
																				} else {
																					if v64 != 0 {
																						if l1 != 0 {
																							v109 = int32(0)
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v109
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v13
																							v112 = int32(1259)
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v112
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v109
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = l1
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v112
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v112
																							v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																							v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v109
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v122
																							F_recordDependencyOn(m, v11+int32(52), v11+int32(40), int32(80))
																							mBase = m.M
																							v132 = m.ExcPending
																							if v132 != 0 {
																								return
																							} else {
																								F_recordDependencyOn(m, v11+int32(52), v11+int32(28), int32(83))
																								mBase = m.M
																								v139 = m.ExcPending
																								if v139 != 0 {
																									return
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v152 = m.ExcPending
																									if v152 != 0 {
																										return
																									} else {
																										m.G0 = v11 + int32(160)
																										return
																									}
																								}
																							}
																						} else {
																							v140 = int32(1259)
																							v143 = F_deleteDependencyRecordsForClass(m, v140, v13, v140, int32(80))
																							mBase = m.M
																							v144 = m.ExcPending
																							if v144 != 0 {
																								return
																							} else {
																								v145 = int32(1259)
																								v148 = F_deleteDependencyRecordsForClass(m, v145, v13, v145, int32(83))
																								mBase = m.M
																								v149 = m.ExcPending
																								if v149 != 0 {
																									return
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v152 = m.ExcPending
																									if v152 != 0 {
																										return
																									} else {
																										m.G0 = v11 + int32(160)
																										return
																									}
																								}
																							}
																						}
																					} else {
																						m.G0 = v11 + int32(160)
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
												v78 = F_table_open(m, int32(1259), int32(3))
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return
												} else {
													v81 = F_SearchSysCacheLockedCopy1(m, int32(57), v13)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return
													} else {
														if v81 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
																F_errmsg_internal(m, int32(49777), v11)
																mBase = m.M
																v180 = m.ExcPending
																if v180 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(516405), int32(4584), int32(260026))
																	mBase = m.M
																	v185 = m.ExcPending
																	if v185 != 0 {
																		return
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+8)))
															*(*uint16)(unsafe.Add(mBase, uint32(v11)+56)) = uint16(v85)
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v87
															v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
															v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
															*(*uint8)(unsafe.Add(mBase, uint32(v89+v90)+131)) = uint8(base.B2i32(l1 != int32(0)))
															F_CatalogTupleUpdate(m, v78, v11+int32(52), v81)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return
															} else {
																F_UnlockTuple(m, v78, v11+int32(52), int32(7))
																mBase = m.M
																v103 = m.ExcPending
																if v103 != 0 {
																	return
																} else {
																	F_pfree(m, v81)
																	mBase = m.M
																	v105 = m.ExcPending
																	if v105 != 0 {
																		return
																	} else {
																		F_sequence_close(m, v78, int32(3))
																		mBase = m.M
																		v108 = m.ExcPending
																		if v108 != 0 {
																			return
																		} else {
																			if v64 != 0 {
																				if l1 != 0 {
																					v109 = int32(0)
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v109
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v13
																					v112 = int32(1259)
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v112
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v109
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = l1
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v112
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v112
																					v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																					v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v109
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v122
																					F_recordDependencyOn(m, v11+int32(52), v11+int32(40), int32(80))
																					mBase = m.M
																					v132 = m.ExcPending
																					if v132 != 0 {
																						return
																					} else {
																						F_recordDependencyOn(m, v11+int32(52), v11+int32(28), int32(83))
																						mBase = m.M
																						v139 = m.ExcPending
																						if v139 != 0 {
																							return
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v152 = m.ExcPending
																							if v152 != 0 {
																								return
																							} else {
																								m.G0 = v11 + int32(160)
																								return
																							}
																						}
																					}
																				} else {
																					v140 = int32(1259)
																					v143 = F_deleteDependencyRecordsForClass(m, v140, v13, v140, int32(80))
																					mBase = m.M
																					v144 = m.ExcPending
																					if v144 != 0 {
																						return
																					} else {
																						v145 = int32(1259)
																						v148 = F_deleteDependencyRecordsForClass(m, v145, v13, v145, int32(83))
																						mBase = m.M
																						v149 = m.ExcPending
																						if v149 != 0 {
																							return
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v152 = m.ExcPending
																							if v152 != 0 {
																								return
																							} else {
																								m.G0 = v11 + int32(160)
																								return
																							}
																						}
																					}
																				}
																			} else {
																				m.G0 = v11 + int32(160)
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
							if l1 == int32(0) {
								F_CatalogTupleDelete(m, v16, v41+int32(4))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									v64 = int32(1)
									F_systable_endscan(m, v39)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										F_relation_close(m, v16, int32(3))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											if l1 != 0 {
												F_LockRelationOid(m, l1, int32(4))
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return
												} else {
													F_SetRelationHasSubclass(m, l1, int32(1))
													mBase = m.M
													v75 = m.ExcPending
													if v75 != 0 {
														return
													} else {
														v78 = F_table_open(m, int32(1259), int32(3))
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return
														} else {
															v81 = F_SearchSysCacheLockedCopy1(m, int32(57), v13)
															mBase = m.M
															v82 = m.ExcPending
															if v82 != 0 {
																return
															} else {
																if v81 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v176 = m.ExcPending
																	if v176 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
																		F_errmsg_internal(m, int32(49777), v11)
																		mBase = m.M
																		v180 = m.ExcPending
																		if v180 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(516405), int32(4584), int32(260026))
																			mBase = m.M
																			v185 = m.ExcPending
																			if v185 != 0 {
																				return
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		}
																	}
																} else {
																	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+8)))
																	*(*uint16)(unsafe.Add(mBase, uint32(v11)+56)) = uint16(v85)
																	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v87
																	v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
																	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v89+v90)+131)) = uint8(base.B2i32(l1 != int32(0)))
																	F_CatalogTupleUpdate(m, v78, v11+int32(52), v81)
																	mBase = m.M
																	v98 = m.ExcPending
																	if v98 != 0 {
																		return
																	} else {
																		F_UnlockTuple(m, v78, v11+int32(52), int32(7))
																		mBase = m.M
																		v103 = m.ExcPending
																		if v103 != 0 {
																			return
																		} else {
																			F_pfree(m, v81)
																			mBase = m.M
																			v105 = m.ExcPending
																			if v105 != 0 {
																				return
																			} else {
																				F_sequence_close(m, v78, int32(3))
																				mBase = m.M
																				v108 = m.ExcPending
																				if v108 != 0 {
																					return
																				} else {
																					if v64 != 0 {
																						if l1 != 0 {
																							v109 = int32(0)
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v109
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v13
																							v112 = int32(1259)
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v112
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v109
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = l1
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v112
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v112
																							v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																							v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v109
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v122
																							F_recordDependencyOn(m, v11+int32(52), v11+int32(40), int32(80))
																							mBase = m.M
																							v132 = m.ExcPending
																							if v132 != 0 {
																								return
																							} else {
																								F_recordDependencyOn(m, v11+int32(52), v11+int32(28), int32(83))
																								mBase = m.M
																								v139 = m.ExcPending
																								if v139 != 0 {
																									return
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v152 = m.ExcPending
																									if v152 != 0 {
																										return
																									} else {
																										m.G0 = v11 + int32(160)
																										return
																									}
																								}
																							}
																						} else {
																							v140 = int32(1259)
																							v143 = F_deleteDependencyRecordsForClass(m, v140, v13, v140, int32(80))
																							mBase = m.M
																							v144 = m.ExcPending
																							if v144 != 0 {
																								return
																							} else {
																								v145 = int32(1259)
																								v148 = F_deleteDependencyRecordsForClass(m, v145, v13, v145, int32(83))
																								mBase = m.M
																								v149 = m.ExcPending
																								if v149 != 0 {
																									return
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v152 = m.ExcPending
																									if v152 != 0 {
																										return
																									} else {
																										m.G0 = v11 + int32(160)
																										return
																									}
																								}
																							}
																						}
																					} else {
																						m.G0 = v11 + int32(160)
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
												v78 = F_table_open(m, int32(1259), int32(3))
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return
												} else {
													v81 = F_SearchSysCacheLockedCopy1(m, int32(57), v13)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return
													} else {
														if v81 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
																F_errmsg_internal(m, int32(49777), v11)
																mBase = m.M
																v180 = m.ExcPending
																if v180 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(516405), int32(4584), int32(260026))
																	mBase = m.M
																	v185 = m.ExcPending
																	if v185 != 0 {
																		return
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+8)))
															*(*uint16)(unsafe.Add(mBase, uint32(v11)+56)) = uint16(v85)
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v87
															v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
															v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
															*(*uint8)(unsafe.Add(mBase, uint32(v89+v90)+131)) = uint8(base.B2i32(l1 != int32(0)))
															F_CatalogTupleUpdate(m, v78, v11+int32(52), v81)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return
															} else {
																F_UnlockTuple(m, v78, v11+int32(52), int32(7))
																mBase = m.M
																v103 = m.ExcPending
																if v103 != 0 {
																	return
																} else {
																	F_pfree(m, v81)
																	mBase = m.M
																	v105 = m.ExcPending
																	if v105 != 0 {
																		return
																	} else {
																		F_sequence_close(m, v78, int32(3))
																		mBase = m.M
																		v108 = m.ExcPending
																		if v108 != 0 {
																			return
																		} else {
																			if v64 != 0 {
																				if l1 != 0 {
																					v109 = int32(0)
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v109
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v13
																					v112 = int32(1259)
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v112
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v109
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = l1
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v112
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v112
																					v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																					v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v109
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v122
																					F_recordDependencyOn(m, v11+int32(52), v11+int32(40), int32(80))
																					mBase = m.M
																					v132 = m.ExcPending
																					if v132 != 0 {
																						return
																					} else {
																						F_recordDependencyOn(m, v11+int32(52), v11+int32(28), int32(83))
																						mBase = m.M
																						v139 = m.ExcPending
																						if v139 != 0 {
																							return
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v152 = m.ExcPending
																							if v152 != 0 {
																								return
																							} else {
																								m.G0 = v11 + int32(160)
																								return
																							}
																						}
																					}
																				} else {
																					v140 = int32(1259)
																					v143 = F_deleteDependencyRecordsForClass(m, v140, v13, v140, int32(80))
																					mBase = m.M
																					v144 = m.ExcPending
																					if v144 != 0 {
																						return
																					} else {
																						v145 = int32(1259)
																						v148 = F_deleteDependencyRecordsForClass(m, v145, v13, v145, int32(83))
																						mBase = m.M
																						v149 = m.ExcPending
																						if v149 != 0 {
																							return
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v152 = m.ExcPending
																							if v152 != 0 {
																								return
																							} else {
																								m.G0 = v11 + int32(160)
																								return
																							}
																						}
																					}
																				}
																			} else {
																				m.G0 = v11 + int32(160)
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
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
								v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+22)))
								v60 = v58 + v59
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
								if v61 != l1 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
										*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v161
										F_errmsg_internal(m, int32(43698), v11+int32(16))
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return
										} else {
											F_errfinish(m, int32(516405), int32(4511), int32(30241))
											mBase = m.M
											v172 = m.ExcPending
											if v172 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v64 = v3
									F_systable_endscan(m, v39)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										F_relation_close(m, v16, int32(3))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											if l1 != 0 {
												F_LockRelationOid(m, l1, int32(4))
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return
												} else {
													F_SetRelationHasSubclass(m, l1, int32(1))
													mBase = m.M
													v75 = m.ExcPending
													if v75 != 0 {
														return
													} else {
														v78 = F_table_open(m, int32(1259), int32(3))
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return
														} else {
															v81 = F_SearchSysCacheLockedCopy1(m, int32(57), v13)
															mBase = m.M
															v82 = m.ExcPending
															if v82 != 0 {
																return
															} else {
																if v81 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v176 = m.ExcPending
																	if v176 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
																		F_errmsg_internal(m, int32(49777), v11)
																		mBase = m.M
																		v180 = m.ExcPending
																		if v180 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(516405), int32(4584), int32(260026))
																			mBase = m.M
																			v185 = m.ExcPending
																			if v185 != 0 {
																				return
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		}
																	}
																} else {
																	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+8)))
																	*(*uint16)(unsafe.Add(mBase, uint32(v11)+56)) = uint16(v85)
																	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v87
																	v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
																	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v89+v90)+131)) = uint8(base.B2i32(l1 != int32(0)))
																	F_CatalogTupleUpdate(m, v78, v11+int32(52), v81)
																	mBase = m.M
																	v98 = m.ExcPending
																	if v98 != 0 {
																		return
																	} else {
																		F_UnlockTuple(m, v78, v11+int32(52), int32(7))
																		mBase = m.M
																		v103 = m.ExcPending
																		if v103 != 0 {
																			return
																		} else {
																			F_pfree(m, v81)
																			mBase = m.M
																			v105 = m.ExcPending
																			if v105 != 0 {
																				return
																			} else {
																				F_sequence_close(m, v78, int32(3))
																				mBase = m.M
																				v108 = m.ExcPending
																				if v108 != 0 {
																					return
																				} else {
																					if v64 != 0 {
																						if l1 != 0 {
																							v109 = int32(0)
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v109
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v13
																							v112 = int32(1259)
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v112
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v109
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = l1
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v112
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v112
																							v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																							v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v109
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v122
																							F_recordDependencyOn(m, v11+int32(52), v11+int32(40), int32(80))
																							mBase = m.M
																							v132 = m.ExcPending
																							if v132 != 0 {
																								return
																							} else {
																								F_recordDependencyOn(m, v11+int32(52), v11+int32(28), int32(83))
																								mBase = m.M
																								v139 = m.ExcPending
																								if v139 != 0 {
																									return
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v152 = m.ExcPending
																									if v152 != 0 {
																										return
																									} else {
																										m.G0 = v11 + int32(160)
																										return
																									}
																								}
																							}
																						} else {
																							v140 = int32(1259)
																							v143 = F_deleteDependencyRecordsForClass(m, v140, v13, v140, int32(80))
																							mBase = m.M
																							v144 = m.ExcPending
																							if v144 != 0 {
																								return
																							} else {
																								v145 = int32(1259)
																								v148 = F_deleteDependencyRecordsForClass(m, v145, v13, v145, int32(83))
																								mBase = m.M
																								v149 = m.ExcPending
																								if v149 != 0 {
																									return
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v152 = m.ExcPending
																									if v152 != 0 {
																										return
																									} else {
																										m.G0 = v11 + int32(160)
																										return
																									}
																								}
																							}
																						}
																					} else {
																						m.G0 = v11 + int32(160)
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
												v78 = F_table_open(m, int32(1259), int32(3))
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return
												} else {
													v81 = F_SearchSysCacheLockedCopy1(m, int32(57), v13)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return
													} else {
														if v81 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
																F_errmsg_internal(m, int32(49777), v11)
																mBase = m.M
																v180 = m.ExcPending
																if v180 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(516405), int32(4584), int32(260026))
																	mBase = m.M
																	v185 = m.ExcPending
																	if v185 != 0 {
																		return
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+8)))
															*(*uint16)(unsafe.Add(mBase, uint32(v11)+56)) = uint16(v85)
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v87
															v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
															v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
															*(*uint8)(unsafe.Add(mBase, uint32(v89+v90)+131)) = uint8(base.B2i32(l1 != int32(0)))
															F_CatalogTupleUpdate(m, v78, v11+int32(52), v81)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return
															} else {
																F_UnlockTuple(m, v78, v11+int32(52), int32(7))
																mBase = m.M
																v103 = m.ExcPending
																if v103 != 0 {
																	return
																} else {
																	F_pfree(m, v81)
																	mBase = m.M
																	v105 = m.ExcPending
																	if v105 != 0 {
																		return
																	} else {
																		F_sequence_close(m, v78, int32(3))
																		mBase = m.M
																		v108 = m.ExcPending
																		if v108 != 0 {
																			return
																		} else {
																			if v64 != 0 {
																				if l1 != 0 {
																					v109 = int32(0)
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v109
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v13
																					v112 = int32(1259)
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v112
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v109
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = l1
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v112
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v112
																					v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																					v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v109
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v122
																					F_recordDependencyOn(m, v11+int32(52), v11+int32(40), int32(80))
																					mBase = m.M
																					v132 = m.ExcPending
																					if v132 != 0 {
																						return
																					} else {
																						F_recordDependencyOn(m, v11+int32(52), v11+int32(28), int32(83))
																						mBase = m.M
																						v139 = m.ExcPending
																						if v139 != 0 {
																							return
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v152 = m.ExcPending
																							if v152 != 0 {
																								return
																							} else {
																								m.G0 = v11 + int32(160)
																								return
																							}
																						}
																					}
																				} else {
																					v140 = int32(1259)
																					v143 = F_deleteDependencyRecordsForClass(m, v140, v13, v140, int32(80))
																					mBase = m.M
																					v144 = m.ExcPending
																					if v144 != 0 {
																						return
																					} else {
																						v145 = int32(1259)
																						v148 = F_deleteDependencyRecordsForClass(m, v145, v13, v145, int32(83))
																						mBase = m.M
																						v149 = m.ExcPending
																						if v149 != 0 {
																							return
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v152 = m.ExcPending
																							if v152 != 0 {
																								return
																							} else {
																								m.G0 = v11 + int32(160)
																								return
																							}
																						}
																					}
																				}
																			} else {
																				m.G0 = v11 + int32(160)
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
						}
					}
				}
			}
		}
	}
}
func F_build_index_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
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
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 float64
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
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
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v345 int32
	_ = v345
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int64
	_ = v508
	var v510 int64
	_ = v510
	var v512 int64
	_ = v512
	var v514 int64
	_ = v514
	var v516 int32
	_ = v516
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v640 int32
	_ = v640
	var v648 int32
	_ = v648
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v777 int32
	_ = v777
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v897 int32
	_ = v897
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v934 int32
	_ = v934
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v957 int32
	_ = v957
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v979 int32
	_ = v979
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	v8 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(48)
	m.G0 = v27
	if l5 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v27 + int32(48)
	return v1099
L2:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v36 = F_bms_copy(m, v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+110)))
	if v31 == int32(1) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v1099 = int32(0)
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if int32(0) < v40 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v57 = v36
	v59 = v8
	v60 = v8
	goto L10
L8:
	;
	v173 = v36
	v176 = v8
	goto L9
L9:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v187 = F_bms_del_member(m, v173, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L5
	} else {
		goto L29
	}
L10:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l3+int32(4)+v59<<(uint(int32(2))%32))))
	if v72 == int32(0) {
		v141 = v57
		v144 = v60
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v173 = v141
	v176 = v144
	goto L9
L12:
	;
	if v144 != 0 {
		goto L25
	} else {
		goto L26
	}
L13:
	;
	v75 = int32(0)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v76 <= v75 {
		v141 = v57
		v144 = v60
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v82 = v75
	v91 = v57
	v94 = v60
	goto L15
L15:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103+v82<<(uint(int32(2))%32))))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if l6 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v141 = v123
	v144 = v124
	goto L12
L17:
	;
	v126 = v82 + int32(1)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v126 < v127 {
		v82 = v126
		v91 = v123
		v94 = v124
		goto L15
	} else {
		goto L24
	}
L18:
	;
	v118 = F_lappend(m, v94, v107)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L22
	}
L19:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+107)))
	if v111 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	if v113 != int32(20) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v116 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v116)
	v123 = v91
	v124 = v94
	goto L17
L22:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v108)+28))
	v121 = F_bms_add_members(m, v91, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v123 = v121
	v124 = v118
	goto L17
L24:
	;
	goto L16
L25:
	;
	v158 = v59 + int32(1)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v158 < v159 {
		v57 = v141
		v59 = v158
		v60 = v144
		goto L10
	} else {
		goto L28
	}
L26:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+106)))
	if v153 == int32(1) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v1099 = int32(0)
	goto L1
L28:
	;
	goto L11
L29:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v190 = F_get_loop_count(m, l0, v189, v187)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v192 = int32(0)
	if l5 != int32(1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v198 = int32(1)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	if v199 != 0 {
		v205 = v198
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v1004 = v192
	v1005 = v192
	v1006 = int32(1)
	v1008 = v8
	v1012 = v192
	goto L33
L33:
	;
	if l4 != 0 {
		goto L202
	} else {
		goto L203
	}
L34:
	;
	v753 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v753
	v757 = int32(*(*uint8)(unsafe.Add(mBase, _consts[386])))
	if v757 == int32(1) {
		goto L164
	} else {
		goto L165
	}
L35:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l2)+60))
	v209 = v205 & base.B2i32(v206 != int32(0))
	if v209 != 0 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	goto L35
L37:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+216)))
	if v200 != 0 {
		v205 = v198
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v201 != 0 {
		v205 = v198
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v205 = base.B2i32(v202 != int32(0))
	goto L36
L40:
	;
	v211 = F_build_index_pathkeys(m, l0, l2, int32(1))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L5
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+105)))
	if v215&v205&int32(1) == int32(0) {
		v735 = v192
		v739 = v8
		v743 = v192
		goto L34
	} else {
		goto L45
	}
L43:
	;
	v213 = F_truncate_useless_pathkeys(m, l0, l1, v211)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	v735 = v192
	v739 = v213
	v743 = v192
	goto L34
L45:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v221 == int32(0) {
		v735 = v192
		v739 = v8
		v743 = v192
		goto L34
	} else {
		goto L46
	}
L46:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if int32(0) < v224 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v727 = F_list_copy_head(m, v713, v710)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L5
	} else {
		goto L163
	}
L48:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v640)+4))
	if v699 == v701 {
		v735 = v640
		v739 = v700
		v743 = v648
		goto L34
	} else {
		goto L162
	}
L49:
	;
	v695 = int32(0)
	if v694 == v695 {
		v735 = v695
		v739 = v680
		v743 = v684
		goto L34
	} else {
		goto L161
	}
L50:
	;
	v233 = v192
	v241 = v192
	v246 = v8
	goto L53
L51:
	;
	goto L52
L52:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v665 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L53:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251+v246<<(uint(int32(2))%32))))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	if v256 != int32(1) {
		v640 = v233
		v648 = v241
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v658 != 0 {
		goto L153
	} else {
		goto L154
	}
L55:
	;
	goto L54
L56:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+16)))
	if v259 != 0 {
		v640 = v233
		v648 = v241
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+41)))
	if v261 != 0 {
		v640 = v233
		v648 = v241
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+8))
	v265 = v27 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v265))) = v260
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v265)+4)) = int32(-1)
	if v267 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v271 = v263
	goto L61
L60:
	;
	v271 = int32(0)
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v265)+8)) = v271
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v260)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v265)+16)) = v273
	if v273 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v273)+12))
	v277 = v275
	goto L64
L63:
	;
	v277 = int32(0)
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v265)+12)) = v277
	v280 = v27 + int32(20)
	v281 = int32(0)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v280)+16))
	if v285 == v281 {
		v334 = v281
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if v334 == int32(0) {
		v640 = v233
		v648 = v241
		goto L55
	} else {
		goto L80
	}
L66:
	;
	goto L65
L67:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v280)+12))
	if v288 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
	v321 = v315 + int32(4)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v316)+4))
	if base.Ui32(v321) < base.Ui32(v317+v323<<(uint(int32(2))%32)) {
		goto L77
	} else {
		goto L78
	}
L69:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v285)+12))
	v315 = v288
	v316 = v285
	v317 = v289
	goto L68
L70:
	;
	goto L71
L71:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	v292 = v290
	goto L72
L72:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v280)+8))
	v297 = F_bms_next_member(m, v296, v292)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v280)+4)) = v297
	if v297 <= int32(0) {
		v334 = v281
		goto L66
	} else {
		goto L74
	}
L73:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v308)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v280)+12)) = v312
	v315 = v312
	v316 = v308
	v317 = v312
	goto L68
L74:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)+12))
	if v302 <= v297 {
		v334 = v281
		goto L66
	} else {
		goto L75
	}
L75:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v301)+20))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v304+v297<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v280)+16)) = v308
	if v308 == int32(0) {
		v292 = v297
		goto L72
	} else {
		goto L76
	}
L76:
	;
	goto L73
L77:
	;
	v328 = v321
	goto L79
L78:
	;
	v328 = int32(0)
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280)+12)) = v328
	v334 = v319
	goto L66
L80:
	;
	v345 = v334
	goto L81
L81:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v345)+8))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v362)+8))
	v364 = int32(0)
	v371 = base.B2i32(v361|v363 == v364)
	if v361 == v364 {
		v410 = v371
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v640 = v233
	v648 = v241
	goto L55
L83:
	;
	v579 = v27 + int32(20)
	v580 = int32(0)
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v579)+16))
	if v584 == v580 {
		v633 = v580
		goto L138
	} else {
		goto L139
	}
L84:
	;
	if v410 == int32(0) {
		goto L83
	} else {
		goto L96
	}
L85:
	;
	goto L84
L86:
	;
	if v363 == int32(0) {
		v410 = v371
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v361)+4))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v363)+4))
	if v377 != v378 {
		v410 = int32(0)
		goto L85
	} else {
		goto L88
	}
L88:
	;
	v380 = int32(1)
	if v377 <= v380 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v383 = v380
	goto L91
L90:
	;
	v383 = v377
	goto L91
L91:
	;
	v384 = int32(8)
	v389 = int32(0)
	goto L92
L92:
	;
	v397 = v389 << (uint(int32(2)) % 32)
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v361+v384+v397)))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v397+(v363+v384))))
	v402 = base.B2i32(v399 == v401)
	if v401 != v399 {
		v410 = v402
		goto L85
	} else {
		goto L94
	}
L93:
	;
	v410 = v402
	goto L85
L94:
	;
	v405 = v389 + int32(1)
	if v405 != v383 {
		v389 = v405
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v416 <= int32(0) {
		goto L83
	} else {
		goto L97
	}
L97:
	;
	v423 = int32(0)
	goto L98
L98:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	if v444 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L83
L100:
	;
	v551 = v423 + int32(1)
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v551 < v552 {
		v423 = v551
		goto L98
	} else {
		goto L136
	}
L101:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	if v447 != int32(17) {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v444)+28))
	if v450 == int32(0) {
		goto L100
	} else {
		goto L103
	}
L103:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v450)+4))
	if v453 < int32(2) {
		goto L100
	} else {
		goto L104
	}
L104:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v450)+12))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v456)))
	if v457 == int32(0) {
		goto L100
	} else {
		goto L105
	}
L105:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v456)+4))
	if v460 == int32(0) {
		goto L100
	} else {
		goto L106
	}
L106:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v255)+8))
	v465 = v423 << (uint(int32(2)) % 32)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v465+v466)))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v444)+4))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v470+v465)))
	if v472 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v444)+24))
	if v472 != v473 {
		goto L100
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v475 = F_match_index_to_operand(m, v457, v423, l2)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L5
	} else {
		goto L113
	}
L110:
	;
	goto L109
L111:
	;
	v536 = F_lappend(m, v233, v532)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L5
	} else {
		goto L133
	}
L112:
	;
	v486 = F_match_index_to_operand(m, v460, v423, l2)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L5
	} else {
		goto L121
	}
L113:
	;
	if v475 == int32(0) {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v479 = F_contain_var_clause(m, v460)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	if v479 != 0 {
		goto L112
	} else {
		goto L116
	}
L116:
	;
	v481 = F_contain_volatile_functions(m, v460)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	if v481 != 0 {
		goto L112
	} else {
		goto L118
	}
L118:
	;
	v483 = F_get_op_opfamily_sortfamily(m, v469, v468)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	if v483 == v463 {
		v532 = v444
		goto L111
	} else {
		goto L120
	}
L120:
	;
	goto L100
L121:
	;
	if v486 == int32(0) {
		goto L100
	} else {
		goto L122
	}
L122:
	;
	v490 = F_contain_var_clause(m, v457)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	if v490 != 0 {
		goto L100
	} else {
		goto L124
	}
L124:
	;
	v492 = F_contain_volatile_functions(m, v457)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L5
	} else {
		goto L125
	}
L125:
	;
	if v492 != 0 {
		goto L100
	} else {
		goto L126
	}
L126:
	;
	v494 = F_get_commutator(m, v469)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L5
	} else {
		goto L127
	}
L127:
	;
	if v494 == int32(0) {
		goto L100
	} else {
		goto L128
	}
L128:
	;
	v498 = F_get_op_opfamily_sortfamily(m, v494, v468)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L5
	} else {
		goto L129
	}
L129:
	;
	if v498 != v463 {
		goto L100
	} else {
		goto L130
	}
L130:
	;
	v502 = F_palloc0(m, int32(36))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L5
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v502))) = int32(17)
	v507 = v502 + int32(8)
	v508 = *(*int64)(unsafe.Add(mBase, uint32(v444)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v507))) = v508
	v510 = *(*int64)(unsafe.Add(mBase, uint32(v444)))
	*(*int64)(unsafe.Add(mBase, uint32(v502))) = v510
	v512 = *(*int64)(unsafe.Add(mBase, uint32(v444)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v502)+16)) = v512
	v514 = *(*int64)(unsafe.Add(mBase, uint32(v444)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v502)+24)) = v514
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v444)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v502)+32)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v507))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v502)+4)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v460
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v460
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v457
	v529 = F_list_make2_impl(m, v27+int32(16), v27+int32(12))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L5
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v502)+28)) = v529
	v532 = v502
	goto L111
L133:
	;
	v538 = F_lappend_int(m, v241, v423)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L5
	} else {
		goto L134
	}
L134:
	;
	v541 = v246 + int32(1)
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if v541 < v542 {
		v233 = v536
		v241 = v538
		v246 = v541
		goto L53
	} else {
		goto L135
	}
L135:
	;
	v640 = v536
	v648 = v538
	goto L55
L136:
	;
	goto L99
L137:
	;
	if v633 != 0 {
		v345 = v633
		goto L81
	} else {
		goto L152
	}
L138:
	;
	goto L137
L139:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v579)+12))
	if v587 != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	v620 = v614 + int32(4)
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v615)+4))
	if base.Ui32(v620) < base.Ui32(v616+v622<<(uint(int32(2))%32)) {
		goto L149
	} else {
		goto L150
	}
L141:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v584)+12))
	v614 = v587
	v615 = v584
	v616 = v588
	goto L140
L142:
	;
	goto L143
L143:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v579)+4))
	v591 = v589
	goto L144
L144:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v579)+8))
	v596 = F_bms_next_member(m, v595, v591)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v579)+4)) = v596
	if v596 <= int32(0) {
		v633 = v580
		goto L138
	} else {
		goto L146
	}
L145:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v607)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v579)+12)) = v611
	v614 = v611
	v615 = v607
	v616 = v611
	goto L140
L146:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v579)))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v600)+12))
	if v601 <= v596 {
		v633 = v580
		goto L138
	} else {
		goto L147
	}
L147:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v600)+20))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v603+v596<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v579)+16)) = v607
	if v607 == int32(0) {
		v591 = v596
		goto L144
	} else {
		goto L148
	}
L148:
	;
	goto L145
L149:
	;
	v627 = v620
	goto L151
L150:
	;
	v627 = int32(0)
	goto L151
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v579)+12)) = v627
	v633 = v618
	goto L138
L152:
	;
	goto L82
L153:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v658)+4))
	if v640 == int32(0) {
		v680 = v658
		v684 = v648
		v694 = v659
		goto L49
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v662 = int32(0)
	if v640 != 0 {
		v699 = v662
		v700 = v662
		goto L48
	} else {
		goto L157
	}
L156:
	;
	v699 = v659
	v700 = v658
	goto L48
L157:
	;
	v735 = int32(0)
	v739 = v662
	v743 = v648
	goto L34
L158:
	;
	v735 = v192
	v739 = int32(0)
	v743 = v192
	goto L34
L159:
	;
	goto L160
L160:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v665)+4))
	v680 = v665
	v684 = v192
	v694 = v669
	goto L49
L161:
	;
	v709 = v695
	v710 = v695
	v713 = v680
	v717 = v684
	goto L47
L162:
	;
	v709 = v640
	v710 = v701
	v713 = v700
	v717 = v648
	goto L47
L163:
	;
	v735 = v709
	v739 = v727
	v743 = v717
	goto L34
L164:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v760)+4))
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	F_pull_varattnos(m, v761, v762, v27+int32(20))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L5
	} else {
		goto L167
	}
L165:
	;
	v979 = v753
	goto L166
L166:
	;
	v1004 = v735
	v1005 = v979
	v1006 = base.B2i32(v209 == int32(0))
	v1008 = v739
	v1012 = v743
	goto L33
L167:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(l2)+96))
	if v767 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v837 = int32(0)
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	if v837 < v838 {
		goto L175
	} else {
		goto L176
	}
L169:
	;
	v770 = int32(0)
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v767)+4))
	if v771 <= v770 {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v777 = v770
	goto L171
L171:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v767)+12))
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v798+v777<<(uint(int32(2))%32))))
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v802)+4))
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	F_pull_varattnos(m, v803, v804, v27+int32(20))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L5
	} else {
		goto L173
	}
L172:
	;
	goto L168
L173:
	;
	v810 = v777 + int32(1)
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v767)+4))
	if v810 < v811 {
		v777 = v810
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v845 = int32(0)
	v849 = v838
	v851 = v837
	goto L178
L176:
	;
	v897 = v837
	goto L177
L177:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v913 = int32(0)
	if v912 == v913 {
		goto L186
	} else {
		goto L187
	}
L178:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v866+v845<<(uint(int32(2))%32))))
	if v870 == int32(0) {
		v883 = v849
		v884 = v851
		goto L180
	} else {
		goto L181
	}
L179:
	;
	v897 = v884
	goto L177
L180:
	;
	v886 = v845 + int32(1)
	if v886 < v883 {
		v845 = v886
		v849 = v883
		v851 = v884
		goto L178
	} else {
		goto L184
	}
L181:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v873+v845))))
	if v875 != int32(1) {
		v883 = v849
		v884 = v851
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v880 = F_bms_add_member(m, v851, v870+int32(7))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L5
	} else {
		goto L183
	}
L183:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	v883 = v882
	v884 = v880
	goto L180
L184:
	;
	goto L179
L185:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	F_bms_free(m, v967)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L5
	} else {
		goto L199
	}
L186:
	;
	v966 = int32(1)
	goto L185
L187:
	;
	goto L188
L188:
	;
	if v897 == int32(0) {
		v957 = v913
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v966 = v957
	goto L185
L190:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v912)+4))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v897)+4))
	if v923 < v922 {
		v957 = v913
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v925 = int32(1)
	if v922 <= v925 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v928 = v925
	goto L194
L193:
	;
	v928 = v922
	goto L194
L194:
	;
	v929 = int32(8)
	v934 = int32(0)
	goto L195
L195:
	;
	v941 = v934 << (uint(int32(2)) % 32)
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v912+v929+v941)))
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v941+(v897+v929))))
	v948 = v943 & (v945 ^ int32(-1))
	v950 = base.B2i32(v948 == int32(0))
	if v948 != 0 {
		v957 = v950
		goto L189
	} else {
		goto L197
	}
L196:
	;
	v957 = v950
	goto L189
L197:
	;
	v952 = v934 + int32(1)
	if v952 != v928 {
		v934 = v952
		goto L195
	} else {
		goto L198
	}
L198:
	;
	goto L196
L199:
	;
	F_bms_free(m, v897)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L5
	} else {
		goto L200
	}
L200:
	;
	v979 = v966
	goto L166
L201:
	;
	if v1006 != 0 {
		v1099 = v1055
		goto L1
	} else {
		goto L219
	}
L202:
	;
	v1026 = int32(0)
	v1029 = F_create_index_path(m, l0, l2, v176, v1004, v1012, v1008, int32(1), v1005, v187, v190, v1026)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L5
	} else {
		goto L207
	}
L203:
	;
	if v176 != 0 {
		goto L202
	} else {
		goto L204
	}
L204:
	;
	if v1008 != 0 {
		goto L202
	} else {
		goto L205
	}
L205:
	;
	v1022 = int32(0)
	if v1005 == v1022 {
		v1055 = v1022
		goto L201
	} else {
		goto L206
	}
L206:
	;
	goto L202
L207:
	;
	v1031 = F_lappend(m, v1026, v1029)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L5
	} else {
		goto L208
	}
L208:
	;
	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+111)))
	if v1033 != int32(1) {
		v1055 = v1031
		goto L201
	} else {
		goto L209
	}
L209:
	;
	if l5 == int32(1) {
		v1055 = v1031
		goto L201
	} else {
		goto L210
	}
L210:
	;
	if v187 != 0 {
		v1055 = v1031
		goto L201
	} else {
		goto L211
	}
L211:
	;
	v1038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v1038&int32(1) == int32(0) {
		v1055 = v1031
		goto L201
	} else {
		goto L212
	}
L212:
	;
	v1043 = int32(1)
	v1046 = F_create_index_path(m, l0, l2, v176, v1004, v1012, v1008, v1043, v1005, int32(0), v190, v1043)
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L5
	} else {
		goto L213
	}
L213:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v1046)+24))
	if int32(0) < v1048 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	F_add_partial_path(m, l1, v1046)
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L5
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	F_pfree(m, v1046)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L5
	} else {
		goto L218
	}
L217:
	;
	v1055 = v1031
	goto L201
L218:
	;
	v1055 = v1031
	goto L201
L219:
	;
	v1058 = F_build_index_pathkeys(m, l0, l2, int32(-1))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L5
	} else {
		goto L220
	}
L220:
	;
	v1060 = F_truncate_useless_pathkeys(m, l0, l1, v1058)
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L5
	} else {
		goto L221
	}
L221:
	;
	if v1060 == int32(0) {
		v1099 = v1055
		goto L1
	} else {
		goto L222
	}
L222:
	;
	v1064 = int32(0)
	v1068 = F_create_index_path(m, l0, l2, v176, v1064, v1064, v1060, int32(-1), v1005, v187, v190, v1064)
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L5
	} else {
		goto L223
	}
L223:
	;
	v1070 = F_lappend(m, v1055, v1068)
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L5
	} else {
		goto L224
	}
L224:
	;
	v1072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+111)))
	if v1072 != int32(1) {
		v1099 = v1070
		goto L1
	} else {
		goto L225
	}
L225:
	;
	if l5 == int32(1) {
		v1099 = v1070
		goto L1
	} else {
		goto L226
	}
L226:
	;
	if v187 != 0 {
		v1099 = v1070
		goto L1
	} else {
		goto L227
	}
L227:
	;
	v1077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v1077&int32(1) == int32(0) {
		v1099 = v1070
		goto L1
	} else {
		goto L228
	}
L228:
	;
	v1082 = int32(0)
	v1087 = F_create_index_path(m, l0, l2, v176, v1082, v1082, v1060, int32(-1), v1005, v1082, v190, int32(1))
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L5
	} else {
		goto L229
	}
L229:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1087)+24))
	if int32(0) < v1089 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	F_add_partial_path(m, l1, v1087)
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L5
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	F_pfree(m, v1087)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L5
	} else {
		goto L234
	}
L233:
	;
	v1099 = v1070
	goto L1
L234:
	;
	v1099 = v1070
	goto L1
}
func F_get_index_clause_from_support(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	v5 = l4
	v7 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = F_get_func_support(m, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(48)
	return v111
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
	v111 = v7
	goto L1
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(461)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l5
	v30 = v5 << (uint(int32(2)) % 32)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l5)+52))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30+v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35+v30)))
	v38 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+44)) = uint8(v38)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v37
	v44 = F_OidFunctionCall1Coll(m, v14, int32(0), v12+int32(8))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if v44 == int32(0) {
		v111 = v7
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v49 = F_palloc0(m, int32(20))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = int32(281)
	v53 = int32(0)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v53 < v54 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v60 = v53
	v63 = int32(0)
	goto L13
L11:
	;
	v90 = v53
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = l1
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+44)))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v49)+14)) = uint16(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+12)) = uint8(v99)
	v111 = v49
	goto L1
L13:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67+v63<<(uint(int32(2))%32))))
	v73 = int32(0)
	v80 = F_make_restrictinfo(m, l0, v71, int32(1), v73, v73, v73, v73, v73, v73, v73)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L15
	}
L14:
	;
	v90 = v82
	goto L12
L15:
	;
	v82 = F_lappend(m, v60, v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v85 = v63 + int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v85 < v86 {
		v60 = v82
		v63 = v85
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
}
func F_get_index_column_opclass(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v6 = F_SearchSysCache1(m, int32(34), l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
			v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14+v15)+10)))
			if l1 <= v17 {
				v21 = F_SysCacheGetAttrNotNull(m, int32(34), v6, int32(18))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v21+l1<<(uint(int32(2))%32))+20))
					v27 = v26
					F_ReleaseCatCache(m, v6)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						return v27
					}
				}
			} else {
				v27 = int32(0)
				F_ReleaseCatCache(m, v6)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					return v27
				}
			}
		}
	}
}
func F_get_index_isvalid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_SearchSysCache1(m, int32(34), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(42668), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(522036), int32(3726), int32(454596))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v29)+18)))
			F_ReleaseCatCache(m, v9)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v31
			}
		}
	}
}
func F_index_concurrently_build(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_table_open(m, l0, int32(4))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, _consts[4]))
		*(*int32)(unsafe.Add(mBase, uint32(v9+int32(12)))) = v19
		v22 = *(*int32)(unsafe.Add(mBase, _consts[119]))
		*(*int32)(unsafe.Add(mBase, uint32(v9+int32(8)))) = v22
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
		*(*int32)(unsafe.Add(mBase, _consts[119])) = v26 | int32(2)
		*(*int32)(unsafe.Add(mBase, _consts[4])) = v25
		v34 = int32(4551928)
		v36 = *(*int32)(unsafe.Add(mBase, _consts[156]))
		v38 = v36 + int32(1)
		*(*int32)(unsafe.Add(mBase, _consts[156])) = v38
		F_RestrictSearchPath(m)
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return
		} else {
			v43 = F_index_open(m, l1, int32(3))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				v45 = F_BuildIndexInfo(m, v43)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					v47 = int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(v45)+121)) = uint16(v47)
					F_index_build(m, v12, v43, v45, int32(0), v47)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						F_AtEOXact_GUC(m, int32(0), v38)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
							*(*int32)(unsafe.Add(mBase, _consts[119])) = v57
							*(*int32)(unsafe.Add(mBase, _consts[4])) = v56
							F_sequence_close(m, v12, int32(0))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_relation_close(m, v43, int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									F_index_set_state_flags(m, l1, int32(0))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										m.G0 = v9 + int32(16)
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
func F_index_form_tuple_context(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v179 int32
	_ = v179
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	v5 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(208)
	m.G0 = v20
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+206)) = uint16(v5)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = F__emscripten_memset_bulkmem(m, v20-int32(-64), base.I32_extend8_s(v5), int32(128))
	mBase = m.M
	goto L1
L1:
	;
	v31 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+56)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v20)+40)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = v31
	if v24 <= int32(32) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v212 = F_heap_compute_data_size(m, l0, v20-int32(-64), l2)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L9
	} else {
		goto L39
	}
L3:
	;
	v198 = v179
	v201 = int32(0)
	v208 = int32(8)
	goto L2
L4:
	;
	v72 = v5
	goto L14
L5:
	;
	if int32(0) < v24 {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v179 = v5
	goto L3
L9:
	;
	return int32(0)
L10:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v24
	F_errmsg(m, int32(708003), v20)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(521746), int32(90), int32(64758))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	v83 = v72 << (uint(int32(2)) % 32)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1+v83)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v89 = v20 + int32(32) + v72
	v90 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v90)
	v94 = v20 - int32(-64) + v83
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v85
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v72))))
	if v97 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v151 = int32(0)
	goto L33
L16:
	;
	v143 = v72 + int32(1)
	if v143 != v24 {
		v72 = v143
		goto L14
	} else {
		goto L32
	}
L17:
	;
	v103 = l0 + int32(20) + v86<<(uint(int32(4))%32) + v72*int32(100)
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+72)))
	if v104 != int32(65535) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v109 = base.B2i32(v107 != int32(1))
	if v107 != int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v116 = v85
	v117 = v107
	goto L21
L20:
	;
	v110 = F_detoast_external_attr(m, v85)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L9
	} else {
		goto L22
	}
L21:
	;
	if v117&int32(3) != 0 {
		goto L16
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v110
	v113 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v113)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	v116 = v110
	v117 = v115
	goto L21
L23:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if base.Ui32(v120) < base.Ui32(int32(2044)) {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+84)))
	switch v123 - int32(109) {
	case 0, 11:
		goto L25
	default:
		goto L16
	}
L25:
	;
	v126 = int32(*(*int8)(unsafe.Add(mBase, uint32(v103)+85)))
	v127 = F_toast_compress_datum(m, v116, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	if v127 == int32(0) {
		goto L16
	} else {
		goto L27
	}
L27:
	;
	if v109 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_pfree(m, v116)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L9
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v127
	v136 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v136)
	goto L16
L31:
	;
	goto L30
L32:
	;
	goto L15
L33:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v151))))
	if v164 != int32(1) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v198 = v164
	v201 = int32(32768)
	v208 = int32(16)
	goto L2
L35:
	;
	v168 = v151 + int32(1)
	if v24 != v168 {
		v151 = v168
		goto L33
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	goto L34
L38:
	;
	v179 = v164
	goto L3
L39:
	;
	v216 = v212 + v208 + int32(7)
	v218 = v216 & int32(-8)
	v219 = F_MemoryContextAllocZero(m, l3, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	if v198 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v225 = v219 + int32(8)
	goto L43
L42:
	;
	v225 = int32(0)
	goto L43
L43:
	;
	F_heap_fill_tuple(m, l0, v20-int32(-64), l2, v219+v208, v20+int32(206), v225)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	if int32(0) < v24 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v239 = int32(0)
	goto L48
L46:
	;
	goto L47
L47:
	;
	if base.Ui32(int32(8192)) <= base.Ui32(v216) {
		goto L55
	} else {
		goto L56
	}
L48:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(32)+v239))))
	if v254 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L47
L50:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v20-int32(-64)+v239<<(uint(int32(2))%32))))
	F_pfree(m, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L9
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v266 = v239 + int32(1)
	if v266 != v24 {
		v239 = v266
		goto L48
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	goto L49
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L9
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+206)))
	v313 = v307<<(uint(int32(13))%32)&int32(16384) | (v201 | v218)
	*(*uint16)(unsafe.Add(mBase, uint32(v219)+6)) = uint16(v313)
	m.G0 = v20 + int32(208)
	return v219
L58:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L9
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(8191)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v218
	F_errmsg(m, int32(38225), v20+int32(16))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L9
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(521746), int32(210), int32(64758))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_index_getnext_tid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+204))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+100))
	if v12 != 0 {
		v13 = m.T0[v12].(func(*base.Module, int32, int32) int32)(m, l0, l1)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+66)) = uint8(v17)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v17)
			if v13 == v17 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
				if v23 == int32(0) {
					v52 = v3
					m.G0 = v8 + int32(16)
					return v52
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+188))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
					m.T0[v28].(func(*base.Module, int32))(m, v23)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v52 = v3
						m.G0 = v8 + int32(16)
						return v52
					}
				}
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+272))
				if v32 == int32(0) {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+268)))
					if v35 != int32(1) {
						v52 = l0 + int32(60)
						m.G0 = v8 + int32(16)
						return v52
					} else {
						F_pgstat_assoc_relation(m, v31)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+272))
							v42 = v41
							v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v42)+24)) = v43 + int64(1)
							v52 = l0 + int32(60)
							m.G0 = v8 + int32(16)
							return v52
						}
					}
				} else {
					v42 = v32
					v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v42)+24)) = v43 + int64(1)
					v52 = l0 + int32(60)
					m.G0 = v8 + int32(16)
					return v52
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return int32(0)
		} else {
			v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+48))
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(400360)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v62 + int32(4)
			F_errmsg_internal(m, int32(722793), v8)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(519997), int32(626), int32(452297))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_index_insert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v20 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v20 != v16 {
		v23 = *(*int32)(unsafe.Add(mBase, _consts[41]))
		v24 = F_list_member_ptr(m, v23, v16)
		mBase = m.M
		v25 = v24
	} else {
		v25 = int32(1)
	}
	if v25 == int32(0) {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+44))
		if v29 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(87529)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v71 + int32(4)
				F_errmsg_internal(m, int32(722793), v14+int32(16))
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(519997), int32(223), int32(87581))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+23)))
			if v32 != 0 {
				v41 = v29
				v42 = m.T0[v41].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3, l4, l5, l6, l7)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					m.G0 = v14 + int32(32)
					return v42
				}
			} else {
				F_CheckForSerializableConflictIn(m, l0, int32(0), int32(-1))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+44))
					v41 = v40
					v42 = m.T0[v41].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3, l4, l5, l6, l7)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						m.G0 = v14 + int32(32)
						return v42
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v55 + int32(4)
				F_errmsg(m, int32(458147), v14)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(519997), int32(222), int32(87581))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
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
func F_makeIndexInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v38 int64
	_ = v38
	var v50 int32
	_ = v50
	v6 = l5
	v7 = l6
	v8 = l7
	v9 = l8
	v10 = l9
	v11 = l10
	v14 = F_palloc0(m, int32(144))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+124)) = uint8(v11)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+123)) = uint8(v10)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+121)) = uint8(v9)
		v21 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v14)+119)) = uint16(v21)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+118)) = uint8(v8)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+117)) = uint8(v7)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+116)) = uint8(v6)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(381)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v21
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+122)) = uint8(v21)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = l3
		v38 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v38
		*(*int64)(unsafe.Add(mBase, uint32(v14)+96)) = v38
		*(*int64)(unsafe.Add(mBase, uint32(v14)+104)) = v38
		*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v14)+136)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v14)+132)) = l2
		v50 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v14)+140)) = v50
		return v14
	}
}
func F_mark_index_clustered(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
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
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v95 int32
	_ = v95
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+119)))
	if v16 != int32(112) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L10
	} else {
		goto L43
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L10
	} else {
		goto L40
	}
L3:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L10
	} else {
		goto L36
	}
L6:
	;
	m.G0 = v13 + int32(32)
	return
L7:
	;
	v19 = F_get_index_isclustered(m, l1)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v23 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L10
	} else {
		goto L13
	}
L10:
	;
	return
L11:
	;
	if v19 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	v25 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L10
	} else {
		goto L15
	}
L14:
	;
	F_sequence_close(m, v23, int32(3))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L10
	} else {
		goto L35
	}
L15:
	;
	if v25 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v29 <= int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v40 = int32(0)
	goto L18
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v40<<(uint(int32(2))%32))))
	v49 = F_SearchSysCacheCopy(m, int32(34), v47, int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L20
	}
L19:
	;
	goto L14
L20:
	;
	if v49 == int32(0) {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+22)))
	v55 = v53 + v54
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+17)))
	if v56 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[203]))
	if v71 != 0 {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	v63 = int32(0)
	goto L25
L24:
	;
	if l1 != v47 {
		goto L22
	} else {
		goto L26
	}
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v55)+17)) = uint8(v63)
	F_CatalogTupleUpdate(m, v23, v49+int32(4), v49)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L10
	} else {
		goto L28
	}
L26:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+18)))
	if v59 == int32(0) {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v63 = int32(1)
	goto L25
L28:
	;
	goto L22
L29:
	;
	v73 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2610), v47, v73, v73, l2)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L10
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	F_pfree(m, v49)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L10
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	v80 = v40 + int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v80 < v81 {
		v40 = v80
		goto L18
	} else {
		goto L34
	}
L34:
	;
	goto L19
L35:
	;
	goto L6
L36:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L10
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(412026), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(517563), int32(565), int32(470674))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v47
	F_errmsg_internal(m, int32(42668), v13)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L10
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(517563), int32(588), int32(470674))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L10
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
	F_errmsg_internal(m, int32(42825), v13+int32(16))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(517563), int32(604), int32(470674))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_transformIndexConstraints(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
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
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int64
	_ = v106
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
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
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
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
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
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
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v296 int32
	_ = v296
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v630 int32
	_ = v630
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v674 int32
	_ = v674
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
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
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
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
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v992 int32
	_ = v992
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1055 int32
	_ = v1055
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int64
	_ = v1111
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1215 int32
	_ = v1215
	var v1219 int32
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1241 int32
	_ = v1241
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1263 int32
	_ = v1263
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1289 int32
	_ = v1289
	var v1293 int32
	_ = v1293
	var v1296 int32
	_ = v1296
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1315 int32
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1328 int32
	_ = v1328
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1341 int32
	_ = v1341
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1354 int32
	_ = v1354
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1367 int32
	_ = v1367
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1389 int32
	_ = v1389
	var v1393 int32
	_ = v1393
	var v1396 int32
	_ = v1396
	var v1403 int32
	_ = v1403
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1411 int32
	_ = v1411
	var v1416 int32
	_ = v1416
	var v1420 int32
	_ = v1420
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1430 int32
	_ = v1430
	var v1435 int32
	_ = v1435
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1445 int32
	_ = v1445
	var v1450 int32
	_ = v1450
	var v1454 int32
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1467 int32
	_ = v1467
	var v1472 int32
	_ = v1472
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1504 int32
	_ = v1504
	var v1507 int32
	_ = v1507
	var v1511 int32
	_ = v1511
	var v1516 int32
	_ = v1516
	var v1540 int32
	_ = v1540
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1558 int32
	_ = v1558
	var v1568 int32
	_ = v1568
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1580 int32
	_ = v1580
	var v1583 int32
	_ = v1583
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1623 int32
	_ = v1623
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1636 int32
	_ = v1636
	var v1660 int32
	_ = v1660
	var v1665 int32
	_ = v1665
	var v1670 int32
	_ = v1670
	var v1675 int32
	_ = v1675
	var v1680 int32
	_ = v1680
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1693 int32
	_ = v1693
	var v1705 int32
	_ = v1705
	var v1718 int32
	_ = v1718
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1729 int32
	_ = v1729
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1749 int32
	_ = v1749
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1776 int32
	_ = v1776
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1792 int32
	_ = v1792
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1805 int32
	_ = v1805
	var v1809 int32
	_ = v1809
	var v1813 int32
	_ = v1813
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1823 int32
	_ = v1823
	var v1828 int32
	_ = v1828
	var v1852 int32
	_ = v1852
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1878 int32
	_ = v1878
	var v1882 int32
	_ = v1882
	var v1885 int32
	_ = v1885
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1900 int32
	_ = v1900
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1929 int64
	_ = v1929
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1975 int32
	_ = v1975
	var v1985 int32
	_ = v1985
	var v1987 int32
	_ = v1987
	var v1992 int32
	_ = v1992
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2014 int32
	_ = v2014
	var v2021 int32
	_ = v2021
	var v2028 int32
	_ = v2028
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2037 int32
	_ = v2037
	var v2038 int32
	_ = v2038
	var v2042 int32
	_ = v2042
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2070 int32
	_ = v2070
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2091 int32
	_ = v2091
	var v2092 int32
	_ = v2092
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2108 int32
	_ = v2108
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2129 int32
	_ = v2129
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2175 int32
	_ = v2175
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2200 int32
	_ = v2200
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	v2 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(320)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v26 == v2 {
		v1971 = l0
		v1975 = v24
		v1985 = v2
		v1987 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v1971)+56))
	if v1992 != 0 {
		goto L452
	} else {
		goto L453
	}
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v29 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v1971 = l0
	v1975 = v24
	v1985 = v2
	v1987 = v2
	goto L1
L4:
	;
	goto L5
L5:
	;
	v32 = l0
	v36 = v24
	v46 = v2
	v48 = v2
	v51 = v26
	v52 = v2
	goto L6
L6:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v52<<(uint(int32(2))%32))))
	v59 = F_palloc0(m, int32(72))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v1971 = v32
	v1975 = v36
	v1985 = v46
	v1987 = v1965
	goto L1
L8:
	;
	return
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = int32(204)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+60)) = uint8(base.B2i32(v63 != int32(8)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v69 = base.B2i32(v67 == int32(6))
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+62)) = uint8(v69)
	if v67 == int32(6) {
		goto L27
	} else {
		goto L28
	}
L10:
	;
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v57)+40))
	if v1540 == int32(0) {
		goto L365
	} else {
		goto L366
	}
L11:
	;
	v1494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+36)))
	if v1494 == int32(0) {
		goto L10
	} else {
		goto L355
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L8
	} else {
		goto L350
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+64)) = v403
	F_errmsg(m, int32(95518), v36-int32(-64))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L8
	} else {
		goto L347
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L8
	} else {
		goto L343
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L8
	} else {
		goto L337
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L8
	} else {
		goto L332
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L8
	} else {
		goto L326
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L8
	} else {
		goto L320
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L8
	} else {
		goto L314
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L8
	} else {
		goto L308
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L8
	} else {
		goto L303
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L8
	} else {
		goto L298
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L8
	} else {
		goto L293
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L8
	} else {
		goto L288
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L8
	} else {
		goto L283
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L8
	} else {
		goto L278
	}
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v32)+56))
	if v73 != 0 {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+30)))
	v76 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+63)) = uint8(v76)
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+61)) = uint8(v75)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+36)))
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+64)) = uint8(v79)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+65)) = uint8(v81)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+66)) = uint8(v83)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if v85 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = v59
	goto L29
L31:
	;
	v86 = F_pstrdup(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L34
	}
L32:
	;
	v89 = int32(0)
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v57)+64))
	if v93 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v89 = v86
	goto L33
L35:
	;
	v95 = v93
	goto L37
L36:
	;
	v95 = int32(428645)
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = v95
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v57)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v59)+28)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v57)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v59)+16)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v57)+68))
	v102 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+69)) = uint8(v102)
	*(*uint16)(unsafe.Add(mBase, uint32(v59)+67)) = uint16(v102)
	v106 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v59)+20)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v59)+32)) = v101
	*(*int64)(unsafe.Add(mBase, uint32(v59)+36)) = v106
	*(*int64)(unsafe.Add(mBase, uint32(v59)+44)) = v106
	*(*int64)(unsafe.Add(mBase, uint32(v59)+52)) = v106
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+60)))
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+70)) = uint8(v115)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v57)+52))
	if v117 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+21)))
	if v118 == int32(0) {
		goto L25
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v319 == int32(8) {
		goto L94
	} else {
		goto L95
	}
L41:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+48))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+68))
	v124 = F_get_relname_relid(m, v117, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	if v124 == int32(0) {
		goto L24
	} else {
		goto L43
	}
L43:
	;
	v129 = F_index_open(m, v124, int32(1))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+192))
	v132 = F_get_index_constraint(m, v124)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	if v132 != 0 {
		goto L23
	} else {
		goto L46
	}
L46:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v121)+56))
	if v134 != v135 {
		goto L22
	} else {
		goto L47
	}
L47:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+18)))
	if v137 == int32(0) {
		goto L21
	} else {
		goto L48
	}
L48:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+12)))
	if v140 == int32(0) {
		goto L20
	} else {
		goto L49
	}
L49:
	;
	v143 = F_RelationGetIndexExpressions(m, v129)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	if v143 != 0 {
		goto L19
	} else {
		goto L51
	}
L51:
	;
	v145 = F_RelationGetIndexPredicate(m, v129)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L8
	} else {
		goto L52
	}
L52:
	;
	if v145 != 0 {
		goto L18
	} else {
		goto L53
	}
L53:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+16)))
	if v147 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+12)))
	if v150 == int32(0) {
		goto L17
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v129)+48))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+84))
	v156 = F_get_index_am_oid(m, int32(428645))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L8
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	if v154 != v156 {
		goto L16
	} else {
		goto L59
	}
L59:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v129)+196))
	v162 = F_SysCacheGetAttrNotNull(m, int32(34), v160, int32(18))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L8
	} else {
		goto L60
	}
L60:
	;
	v164 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131)+8)))
	if int32(0) < v164 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v173 = int32(0)
	goto L64
L62:
	;
	goto L63
L63:
	;
	F_relation_close(m, v129, int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L8
	} else {
		goto L93
	}
L64:
	;
	v194 = v173 << (uint(int32(1)) % 32)
	v196 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131+int32(48)+v194))))
	if int32(0) < v196 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L63
L66:
	;
	v215 = F_pstrdup(m, v212+int32(4))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L8
	} else {
		goto L71
	}
L67:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v121)+52))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v212 = v199 + v200<<(uint(int32(4))%32) + v196*int32(100) - int32(80)
	goto L66
L68:
	;
	goto L69
L69:
	;
	v209 = F_SystemAttributeDefinition(m, v196)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L8
	} else {
		goto L70
	}
L70:
	;
	v212 = v209
	goto L66
L71:
	;
	v217 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131)+10)))
	if v173 < v217 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v271 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131)+8)))
	if v268 < v271 {
		v173 = v268
		goto L64
	} else {
		goto L92
	}
L73:
	;
	v220 = v173 << (uint(int32(2)) % 32)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v129)+56))
	v223 = v173 + int32(1)
	v225 = F_get_attoptions(m, v221, base.I32_extend16_s(v223))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L8
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v57)+40))
	v261 = F_makeString(m, v215)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L8
	} else {
		goto L90
	}
L76:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v212)+68))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v129)+48))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+84))
	v230 = F_GetDefaultOpClass(m, v227, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v220+(v162+int32(24)))))
	if v230 != v233 {
		goto L15
	} else {
		goto L78
	}
L78:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v212)+96))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v129)+248))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v236+v220)))
	if v235 != v238 {
		goto L15
	} else {
		goto L79
	}
L79:
	;
	if v225 != 0 {
		goto L15
	} else {
		goto L80
	}
L80:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v129)+224))
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v240+v194))))
	if v242 != 0 {
		goto L15
	} else {
		goto L81
	}
L81:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v243 == int32(6) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
	v247 = F_makeString(m, v215)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L8
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v57)+32))
	v255 = F_makeString(m, v215)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L8
	} else {
		goto L88
	}
L85:
	;
	v249 = F_makeNotNullConstraint(m, v247)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L8
	} else {
		goto L86
	}
L86:
	;
	v251 = F_lappend(m, v246, v249)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L8
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v251
	goto L84
L88:
	;
	v257 = F_lappend(m, v254, v255)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+32)) = v257
	v268 = v223
	goto L72
L90:
	;
	v263 = F_lappend(m, v260, v261)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L8
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+40)) = v263
	v268 = v173 + int32(1)
	goto L72
L92:
	;
	goto L65
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+44)) = v124
	goto L40
L94:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	if v322 == int32(0) {
		goto L10
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v57)+32))
	if v370 == int32(0) {
		goto L11
	} else {
		goto L104
	}
L97:
	;
	v325 = int32(0)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	if v326 <= v325 {
		goto L10
	} else {
		goto L98
	}
L98:
	;
	v330 = v325
	goto L99
L99:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v322)+12))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v350+v330<<(uint(int32(2))%32))))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)+12))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)+4))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v359 = F_lappend(m, v357, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L8
	} else {
		goto L101
	}
L100:
	;
	goto L10
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = v359
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v59)+36))
	v363 = F_lappend(m, v362, v356)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+36)) = v363
	v367 = v330 + int32(1)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	if v367 < v368 {
		v330 = v367
		goto L99
	} else {
		goto L103
	}
L103:
	;
	goto L100
L104:
	;
	v373 = int32(0)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	if v374 <= v373 {
		goto L11
	} else {
		goto L105
	}
L105:
	;
	v390 = v373
	goto L106
L106:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v370)+12))
	v401 = v398 + v390<<(uint(int32(2))%32)
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v404 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	goto L11
L108:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	if v853 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L109:
	;
	v822 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v436)+19)) = uint8(v822)
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
	v825 = F_makeString(m, v403)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L8
	} else {
		goto L216
	}
L110:
	;
	v582 = int32(0)
	v583 = int32(1)
	v585 = F_strcmp(m, int32(797468), v403)
	mBase = m.M
	if v585 == v582 {
		goto L152
	} else {
		goto L153
	}
L111:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v404)+4))
	if v407 <= int32(0) {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v404)+12))
	v413 = int32(0)
	goto L113
L113:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v410+v413<<(uint(int32(2))%32))))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+4))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403))))
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437))))
	if v441 == int32(0) {
		v460 = v440
		v461 = v441
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v466 = int32(0)
	v467 = int32(1)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v468 != int32(6) {
		v840 = v466
		v842 = v467
		v843 = v436
		goto L108
	} else {
		goto L127
	}
L115:
	;
	if v461-v460 != 0 {
		goto L123
	} else {
		goto L124
	}
L116:
	;
	goto L115
L117:
	;
	if v440 != v441 {
		v460 = v440
		v461 = v441
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v445 = v437
	v446 = v403
	goto L119
L119:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446)+1)))
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445)+1)))
	if v450 == int32(0) {
		v460 = v449
		v461 = v450
		goto L116
	} else {
		goto L121
	}
L120:
	;
	v460 = v449
	v461 = v450
	goto L116
L121:
	;
	v453 = int32(1)
	if v449 == v450 {
		v445 = v445 + v453
		v446 = v446 + v453
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	v464 = v413 + int32(1)
	if v464 != v407 {
		v413 = v464
		goto L113
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	goto L114
L126:
	;
	goto L110
L127:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+21)))
	if v471 != 0 {
		v840 = v466
		v842 = v467
		v843 = v436
		goto L108
	} else {
		goto L128
	}
L128:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436)+19)))
	if v472 != int32(1) {
		goto L109
	} else {
		goto L129
	}
L129:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
	if v475 == int32(0) {
		v840 = v466
		v842 = v467
		v843 = v436
		goto L108
	} else {
		goto L130
	}
L130:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v475)+4))
	if v478 <= int32(0) {
		v840 = v466
		v842 = v467
		v843 = v436
		goto L108
	} else {
		goto L131
	}
L131:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v475)+12))
	v484 = int32(0)
	goto L132
L132:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v481+v484<<(uint(int32(2))%32))))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v507)+32))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v508)+12))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v510)+4))
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403))))
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511))))
	if v515 == int32(0) {
		v534 = v514
		v535 = v515
		goto L135
	} else {
		goto L136
	}
L133:
	;
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+17)))
	if v540 != int32(1) {
		v840 = v466
		v842 = v467
		v843 = v436
		goto L108
	} else {
		goto L146
	}
L134:
	;
	if v535-v534 != 0 {
		goto L142
	} else {
		goto L143
	}
L135:
	;
	goto L134
L136:
	;
	if v514 != v515 {
		v534 = v514
		v535 = v515
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v519 = v511
	v520 = v403
	goto L138
L138:
	;
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520)+1)))
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+1)))
	if v524 == int32(0) {
		v534 = v523
		v535 = v524
		goto L135
	} else {
		goto L140
	}
L139:
	;
	v534 = v523
	v535 = v524
	goto L135
L140:
	;
	v527 = int32(1)
	if v523 == v524 {
		v519 = v519 + v527
		v520 = v520 + v527
		goto L138
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	v538 = v484 + int32(1)
	if v538 != v478 {
		v484 = v538
		goto L132
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	goto L133
L145:
	;
	v840 = v466
	v842 = v467
	v843 = v436
	goto L108
L146:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L8
	} else {
		goto L147
	}
L147:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L8
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v403
	F_errmsg(m, int32(742328), v36+int32(128))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L8
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(522528), int32(2656), int32(96590))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L8
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
	if v614 != 0 {
		goto L170
	} else {
		goto L171
	}
L152:
	;
	v614 = int32(797464)
	goto L151
L153:
	;
	goto L154
L154:
	;
	v590 = F_strcmp(m, int32(797568), v403)
	mBase = m.M
	if v590 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v614 = int32(797564)
	goto L151
L156:
	;
	goto L157
L157:
	;
	v595 = F_strcmp(m, int32(797668), v403)
	mBase = m.M
	if v595 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v614 = int32(797664)
	goto L151
L159:
	;
	goto L160
L160:
	;
	v600 = F_strcmp(m, int32(797768), v403)
	mBase = m.M
	if v600 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v614 = int32(797764)
	goto L151
L162:
	;
	goto L163
L163:
	;
	v605 = F_strcmp(m, int32(797868), v403)
	mBase = m.M
	if v605 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v614 = int32(797864)
	goto L151
L165:
	;
	goto L166
L166:
	;
	v612 = F_strcmp(m, int32(797968), v403)
	mBase = m.M
	if v612 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v613 = int32(0)
	goto L169
L168:
	;
	v613 = int32(797964)
	goto L169
L169:
	;
	v614 = v613
	goto L151
L170:
	;
	v840 = v582
	v842 = v583
	v843 = int32(0)
	goto L108
L171:
	;
	goto L172
L172:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v616 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+21)))
	if v796 != 0 {
		goto L208
	} else {
		goto L209
	}
L174:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v616)+4))
	if v619 <= int32(0) {
		goto L173
	} else {
		goto L175
	}
L175:
	;
	v630 = v582
	goto L176
L176:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v616)+12))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v643+v630<<(uint(int32(2))%32))))
	v649 = F_table_openrv(m, v647, int32(1))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L8
	} else {
		goto L178
	}
L177:
	;
	goto L173
L178:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v649)+48))
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651)+119)))
	v654 = v652 - int32(102)
	if base.Ui32(int32(12)) < base.Ui32(v654) {
		goto L14
	} else {
		goto L179
	}
L179:
	;
	if int32(1)<<(uint(v654)%32)&int32(5121) == int32(0) {
		goto L14
	} else {
		goto L180
	}
L180:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v649)+52))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v663)))
	if int32(0) < v664 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v674 = int32(0)
	goto L184
L182:
	;
	goto L183
L183:
	;
	F_sequence_close(m, v649, int32(0))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L8
	} else {
		goto L206
	}
L184:
	;
	v696 = v663 + v664<<(uint(int32(4))%32) + int32(20) + v674*int32(100)
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696)+91)))
	if v697 != 0 {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	goto L183
L186:
	;
	v745 = v674 + int32(1)
	if v745 != v664 {
		v674 = v745
		goto L184
	} else {
		goto L205
	}
L187:
	;
	v699 = v696 + int32(4)
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699))))
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403))))
	if v703 == int32(0) {
		v722 = v702
		v723 = v703
		goto L189
	} else {
		goto L190
	}
L188:
	;
	if v723-v722 != 0 {
		goto L186
	} else {
		goto L196
	}
L189:
	;
	goto L188
L190:
	;
	if v702 != v703 {
		v722 = v702
		v723 = v703
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v707 = v403
	v708 = v699
	goto L192
L192:
	;
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708)+1)))
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707)+1)))
	if v712 == int32(0) {
		v722 = v711
		v723 = v712
		goto L189
	} else {
		goto L194
	}
L193:
	;
	v722 = v711
	v723 = v712
	goto L189
L194:
	;
	v715 = int32(1)
	if v711 == v712 {
		v707 = v707 + v715
		v708 = v708 + v715
		goto L192
	} else {
		goto L195
	}
L195:
	;
	goto L193
L196:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v696)+68))
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v726 == int32(6) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
	v730 = F_pstrdup(m, v699)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L8
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v739 = int32(0)
	F_sequence_close(m, v649, v739)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L8
	} else {
		goto L204
	}
L200:
	;
	v732 = F_makeString(m, v730)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L8
	} else {
		goto L201
	}
L201:
	;
	v734 = F_makeNotNullConstraint(m, v732)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L8
	} else {
		goto L202
	}
L202:
	;
	v736 = F_lappend(m, v729, v734)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L8
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v736
	goto L199
L204:
	;
	v840 = v725
	v842 = v583
	v843 = v739
	goto L108
L205:
	;
	goto L185
L206:
	;
	v772 = v630 + int32(1)
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v616)+4))
	if v772 < v773 {
		v630 = v772
		goto L176
	} else {
		goto L207
	}
L207:
	;
	goto L177
L208:
	;
	v797 = int32(0)
	v840 = v797
	v842 = v797
	v843 = v797
	goto L108
L209:
	;
	goto L210
L210:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L8
	} else {
		goto L211
	}
L211:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L8
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v403
	F_errmsg(m, int32(74432), v36+int32(96))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L8
	} else {
		goto L213
	}
L213:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v813, v814)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L8
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(522528), int32(2736), int32(96590))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L8
	} else {
		goto L215
	}
L215:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L216:
	;
	v827 = F_makeNotNullConstraint(m, v825)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L8
	} else {
		goto L217
	}
L217:
	;
	v829 = F_lappend(m, v824, v827)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L8
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v829
	v840 = v466
	v842 = v467
	v843 = v436
	goto L108
L219:
	;
	v963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+36)))
	if v963 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L220:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v853)+4))
	if v856 <= int32(0) {
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v853)+12))
	v862 = int32(0)
	goto L222
L222:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v859+v862<<(uint(int32(2))%32))))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v885)+4))
	if v886 == int32(0) {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	goto L219
L224:
	;
	v940 = v862 + int32(1)
	if v856 != v940 {
		v862 = v940
		goto L222
	} else {
		goto L241
	}
L225:
	;
	v891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886))))
	v892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403))))
	if v892 == int32(0) {
		v911 = v891
		v912 = v892
		goto L227
	} else {
		goto L228
	}
L226:
	;
	if v912-v911 != 0 {
		goto L224
	} else {
		goto L234
	}
L227:
	;
	goto L226
L228:
	;
	if v891 != v892 {
		v911 = v891
		v912 = v892
		goto L227
	} else {
		goto L229
	}
L229:
	;
	v896 = v403
	v897 = v886
	goto L230
L230:
	;
	v900 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v897)+1)))
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v896)+1)))
	if v901 == int32(0) {
		v911 = v900
		v912 = v901
		goto L227
	} else {
		goto L232
	}
L231:
	;
	v911 = v900
	v912 = v901
	goto L227
L232:
	;
	v904 = int32(1)
	if v900 == v901 {
		v896 = v896 + v904
		v897 = v897 + v904
		goto L230
	} else {
		goto L233
	}
L233:
	;
	goto L231
L234:
	;
	v914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+62)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L8
	} else {
		goto L235
	}
L235:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L8
	} else {
		goto L236
	}
L236:
	;
	if v914 == int32(1) {
		goto L13
	} else {
		goto L237
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v403
	F_errmsg(m, int32(96380), v36+int32(80))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L8
	} else {
		goto L238
	}
L238:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v930, v931)
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L8
	} else {
		goto L239
	}
L239:
	;
	F_errfinish(m, int32(522528), int32(2755), int32(96590))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L8
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
	goto L223
L242:
	;
	v1105 = F_palloc0(m, int32(36))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L8
	} else {
		goto L274
	}
L243:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v57)+32))
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v966)+12))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v966)+4))
	if v401 != v967+v968<<(uint(int32(2))%32)-int32(4) {
		goto L242
	} else {
		goto L244
	}
L244:
	;
	if v842 == int32(0) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+21)))
	if v977 != int32(1) {
		goto L242
	} else {
		goto L248
	}
L246:
	;
	v1055 = v840
	goto L247
L247:
	;
	if v843 == int32(0) {
		v1074 = v1055
		goto L265
	} else {
		goto L266
	}
L248:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v980)+52))
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v981)))
	if v982 <= int32(0) {
		goto L242
	} else {
		goto L249
	}
L249:
	;
	v992 = int32(0)
	goto L250
L250:
	;
	v1014 = v981 + v982<<(uint(int32(4))%32) + int32(20) + v992*int32(100)
	v1015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1014)+91)))
	if v1015 != 0 {
		goto L242
	} else {
		goto L252
	}
L251:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+68))
	v1055 = v1046
	goto L247
L252:
	;
	v1017 = v1014 + int32(4)
	v1020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403))))
	v1021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1017))))
	if v1021 == int32(0) {
		v1040 = v1020
		v1041 = v1021
		goto L254
	} else {
		goto L255
	}
L253:
	;
	if v1041-v1040 != 0 {
		goto L261
	} else {
		goto L262
	}
L254:
	;
	goto L253
L255:
	;
	if v1020 != v1021 {
		v1040 = v1020
		v1041 = v1021
		goto L254
	} else {
		goto L256
	}
L256:
	;
	v1025 = v1017
	v1026 = v403
	goto L257
L257:
	;
	v1029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1026)+1)))
	v1030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1025)+1)))
	if v1030 == int32(0) {
		v1040 = v1029
		v1041 = v1030
		goto L254
	} else {
		goto L259
	}
L258:
	;
	v1040 = v1029
	v1041 = v1030
	goto L254
L259:
	;
	v1033 = int32(1)
	if v1029 == v1030 {
		v1025 = v1025 + v1033
		v1026 = v1026 + v1033
		goto L257
	} else {
		goto L260
	}
L260:
	;
	goto L258
L261:
	;
	v1044 = v992 + int32(1)
	if v1044 == v982 {
		goto L242
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	goto L251
L264:
	;
	v992 = v1044
	goto L250
L265:
	;
	if v1074 == int32(0) {
		goto L12
	} else {
		goto L269
	}
L266:
	;
	if v1055 != 0 {
		v1074 = v1055
		goto L265
	} else {
		goto L267
	}
L267:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v843)+8))
	v1072 = F_typenameTypeId(m, int32(0), v1071)
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L8
	} else {
		goto L268
	}
L268:
	;
	v1074 = v1072
	goto L265
L269:
	;
	v1077 = F_type_is_range(m, v1074)
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L8
	} else {
		goto L270
	}
L270:
	;
	if v1077 != 0 {
		goto L242
	} else {
		goto L271
	}
L271:
	;
	v1079 = F_type_is_multirange(m, v1074)
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L8
	} else {
		goto L272
	}
L272:
	;
	if v1079 == int32(0) {
		goto L12
	} else {
		goto L273
	}
L273:
	;
	goto L242
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1105))) = int32(92)
	v1109 = F_pstrdup(m, v403)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L8
	} else {
		goto L275
	}
L275:
	;
	v1111 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+8)) = v1111
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+4)) = v1109
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+16)) = v1111
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+24)) = v1111
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+32)) = int32(0)
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	v1121 = F_lappend(m, v1120, v1105)
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L8
	} else {
		goto L276
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = v1121
	v1125 = v390 + int32(1)
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	if v1125 < v1126 {
		v390 = v1125
		goto L106
	} else {
		goto L277
	}
L277:
	;
	goto L107
L278:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L8
	} else {
		goto L279
	}
L279:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1135)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+304)) = v1136
	F_errmsg(m, int32(459285), v36+int32(304))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L8
	} else {
		goto L280
	}
L280:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1143, v1144)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L8
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(522528), int32(2354), int32(96590))
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L8
	} else {
		goto L282
	}
L282:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L283:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L8
	} else {
		goto L284
	}
L284:
	;
	F_errmsg(m, int32(564988), int32(0))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L8
	} else {
		goto L285
	}
L285:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1163, v1164)
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L8
	} else {
		goto L286
	}
L286:
	;
	F_errfinish(m, int32(522528), int32(2420), int32(96590))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L8
	} else {
		goto L287
	}
L287:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L288:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L8
	} else {
		goto L289
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+144)) = v117
	F_errmsg(m, int32(76271), v36+int32(144))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L8
	} else {
		goto L290
	}
L290:
	;
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1185, v1186)
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L8
	} else {
		goto L291
	}
L291:
	;
	F_errfinish(m, int32(522528), int32(2429), int32(96590))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L8
	} else {
		goto L292
	}
L292:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L293:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L8
	} else {
		goto L294
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+288)) = v117
	F_errmsg(m, int32(96444), v36+int32(288))
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L8
	} else {
		goto L295
	}
L295:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1207, v1208)
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L8
	} else {
		goto L296
	}
L296:
	;
	F_errfinish(m, int32(522528), int32(2441), int32(96590))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L8
	} else {
		goto L297
	}
L297:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L298:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L8
	} else {
		goto L299
	}
L299:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v121)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+272)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v36)+276)) = v1223 + int32(4)
	F_errmsg(m, int32(750425), v36+int32(272))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L8
	} else {
		goto L300
	}
L300:
	;
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1233, v1234)
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L8
	} else {
		goto L301
	}
L301:
	;
	F_errfinish(m, int32(522528), int32(2449), int32(96590))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L8
	} else {
		goto L302
	}
L302:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L303:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L8
	} else {
		goto L304
	}
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+256)) = v117
	F_errmsg(m, int32(455447), v36+int32(256))
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L8
	} else {
		goto L305
	}
L305:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1255, v1256)
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L8
	} else {
		goto L306
	}
L306:
	;
	F_errfinish(m, int32(522528), int32(2455), int32(96590))
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L8
	} else {
		goto L307
	}
L307:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L308:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L8
	} else {
		goto L309
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v117
	F_errmsg(m, int32(29502), v36+int32(240))
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L8
	} else {
		goto L310
	}
L310:
	;
	F_errdetail(m, int32(602161), int32(0))
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L8
	} else {
		goto L311
	}
L311:
	;
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1281, v1282)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L8
	} else {
		goto L312
	}
L312:
	;
	F_errfinish(m, int32(522528), int32(2467), int32(96590))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L8
	} else {
		goto L313
	}
L313:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L314:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L8
	} else {
		goto L315
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+224)) = v117
	F_errmsg(m, int32(153585), v36+int32(224))
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L8
	} else {
		goto L316
	}
L316:
	;
	F_errdetail(m, int32(602161), int32(0))
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		goto L8
	} else {
		goto L317
	}
L317:
	;
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1307, v1308)
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L8
	} else {
		goto L318
	}
L318:
	;
	F_errfinish(m, int32(522528), int32(2474), int32(96590))
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		goto L8
	} else {
		goto L319
	}
L319:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L320:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L8
	} else {
		goto L321
	}
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v117
	F_errmsg(m, int32(29413), v36+int32(208))
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L8
	} else {
		goto L322
	}
L322:
	;
	F_errdetail(m, int32(602161), int32(0))
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L8
	} else {
		goto L323
	}
L323:
	;
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1333, v1334)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L8
	} else {
		goto L324
	}
L324:
	;
	F_errfinish(m, int32(522528), int32(2481), int32(96590))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L8
	} else {
		goto L325
	}
L325:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L326:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L8
	} else {
		goto L327
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+192)) = v117
	F_errmsg(m, int32(29529), v36+int32(192))
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L8
	} else {
		goto L328
	}
L328:
	;
	F_errdetail(m, int32(602266), int32(0))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L8
	} else {
		goto L329
	}
L329:
	;
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1359, v1360)
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L8
	} else {
		goto L330
	}
L330:
	;
	F_errfinish(m, int32(522528), int32(2493), int32(96590))
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L8
	} else {
		goto L331
	}
L331:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L332:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L8
	} else {
		goto L333
	}
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+176)) = v117
	F_errmsg(m, int32(428625), v36+int32(176))
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L8
	} else {
		goto L334
	}
L334:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1381, v1382)
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L8
	} else {
		goto L335
	}
L335:
	;
	F_errfinish(m, int32(522528), int32(2505), int32(96590))
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L8
	} else {
		goto L336
	}
L336:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L337:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L8
	} else {
		goto L338
	}
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+164)) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v36)+160)) = v117
	F_errmsg(m, int32(223126), v36+int32(160))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L8
	} else {
		goto L339
	}
L339:
	;
	F_errdetail(m, int32(602161), int32(0))
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L8
	} else {
		goto L340
	}
L340:
	;
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1408, v1409)
	mBase = m.M
	v1411 = m.ExcPending
	if v1411 != 0 {
		goto L8
	} else {
		goto L341
	}
L341:
	;
	F_errfinish(m, int32(522528), int32(2557), int32(96590))
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L8
	} else {
		goto L342
	}
L342:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L343:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L8
	} else {
		goto L344
	}
L344:
	;
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v647)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1424
	F_errmsg(m, int32(410624), v36+int32(112))
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		goto L8
	} else {
		goto L345
	}
L345:
	;
	F_errfinish(m, int32(522528), int32(2700), int32(96590))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L8
	} else {
		goto L346
	}
L346:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L347:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1442, v1443)
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L8
	} else {
		goto L348
	}
L348:
	;
	F_errfinish(m, int32(522528), int32(2749), int32(96590))
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L8
	} else {
		goto L349
	}
L349:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L350:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L8
	} else {
		goto L351
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+48)) = v403
	F_errmsg(m, int32(387112), v36+int32(48))
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L8
	} else {
		goto L352
	}
L352:
	;
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1464, v1465)
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L8
	} else {
		goto L353
	}
L353:
	;
	F_errfinish(m, int32(522528), int32(2799), int32(96590))
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		goto L8
	} else {
		goto L354
	}
L354:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L355:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v57)+32))
	if v1497 != 0 {
		goto L357
	} else {
		goto L358
	}
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = int32(81466)
	goto L10
L357:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v1497)+4))
	if int32(1) < v1498 {
		goto L356
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L8
	} else {
		goto L361
	}
L360:
	;
	goto L359
L361:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1507 = m.ExcPending
	if v1507 != 0 {
		goto L8
	} else {
		goto L362
	}
L362:
	;
	F_errmsg(m, int32(156102), int32(0))
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L8
	} else {
		goto L363
	}
L363:
	;
	F_errfinish(m, int32(522528), int32(2826), int32(96590))
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L8
	} else {
		goto L364
	}
L364:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L365:
	;
	v1965 = F_lappend(m, v48, v59)
	mBase = m.M
	v1966 = m.ExcPending
	if v1966 != 0 {
		goto L8
	} else {
		goto L450
	}
L366:
	;
	v1543 = int32(0)
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v1540)+4))
	if v1544 <= v1543 {
		goto L365
	} else {
		goto L367
	}
L367:
	;
	v1558 = v1543
	goto L368
L368:
	;
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v1540)+12))
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v1568+v1558<<(uint(int32(2))%32))))
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v1572)+4))
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v1574 == int32(0) {
		goto L371
	} else {
		goto L372
	}
L369:
	;
	goto L365
L370:
	;
	v1923 = F_palloc0(m, int32(36))
	mBase = m.M
	v1924 = m.ExcPending
	if v1924 != 0 {
		goto L8
	} else {
		goto L446
	}
L371:
	;
	v1660 = F_strcmp(m, int32(797468), v1573)
	mBase = m.M
	if v1660 == int32(0) {
		goto L387
	} else {
		goto L388
	}
L372:
	;
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v1574)+4))
	if v1577 <= int32(0) {
		goto L371
	} else {
		goto L373
	}
L373:
	;
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(v1574)+12))
	v1583 = int32(0)
	goto L374
L374:
	;
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(v1580+v1583<<(uint(int32(2))%32))))
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v1606)+4))
	v1610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1573))))
	v1611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1607))))
	if v1611 == int32(0) {
		v1630 = v1610
		v1631 = v1611
		goto L377
	} else {
		goto L378
	}
L375:
	;
	goto L371
L376:
	;
	if v1631-v1630 == int32(0) {
		goto L370
	} else {
		goto L384
	}
L377:
	;
	goto L376
L378:
	;
	if v1610 != v1611 {
		v1630 = v1610
		v1631 = v1611
		goto L377
	} else {
		goto L379
	}
L379:
	;
	v1615 = v1607
	v1616 = v1573
	goto L380
L380:
	;
	v1619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1616)+1)))
	v1620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1615)+1)))
	if v1620 == int32(0) {
		v1630 = v1619
		v1631 = v1620
		goto L377
	} else {
		goto L382
	}
L381:
	;
	v1630 = v1619
	v1631 = v1620
	goto L377
L382:
	;
	v1623 = int32(1)
	if v1619 == v1620 {
		v1615 = v1615 + v1623
		v1616 = v1616 + v1623
		goto L380
	} else {
		goto L383
	}
L383:
	;
	goto L381
L384:
	;
	v1636 = v1583 + int32(1)
	if v1577 != v1636 {
		v1583 = v1636
		goto L374
	} else {
		goto L385
	}
L385:
	;
	goto L375
L386:
	;
	if v1689 != 0 {
		goto L370
	} else {
		goto L405
	}
L387:
	;
	v1689 = int32(797464)
	goto L386
L388:
	;
	goto L389
L389:
	;
	v1665 = F_strcmp(m, int32(797568), v1573)
	mBase = m.M
	if v1665 == int32(0) {
		goto L390
	} else {
		goto L391
	}
L390:
	;
	v1689 = int32(797564)
	goto L386
L391:
	;
	goto L392
L392:
	;
	v1670 = F_strcmp(m, int32(797668), v1573)
	mBase = m.M
	if v1670 == int32(0) {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v1689 = int32(797664)
	goto L386
L394:
	;
	goto L395
L395:
	;
	v1675 = F_strcmp(m, int32(797768), v1573)
	mBase = m.M
	if v1675 == int32(0) {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v1689 = int32(797764)
	goto L386
L397:
	;
	goto L398
L398:
	;
	v1680 = F_strcmp(m, int32(797868), v1573)
	mBase = m.M
	if v1680 == int32(0) {
		goto L399
	} else {
		goto L400
	}
L399:
	;
	v1689 = int32(797864)
	goto L386
L400:
	;
	goto L401
L401:
	;
	v1687 = F_strcmp(m, int32(797968), v1573)
	mBase = m.M
	if v1687 != 0 {
		goto L402
	} else {
		goto L403
	}
L402:
	;
	v1688 = int32(0)
	goto L404
L403:
	;
	v1688 = int32(797964)
	goto L404
L404:
	;
	v1689 = v1688
	goto L386
L405:
	;
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v1690 == int32(0) {
		goto L406
	} else {
		goto L407
	}
L406:
	;
	v1878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+21)))
	if v1878 != 0 {
		goto L370
	} else {
		goto L440
	}
L407:
	;
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v1690)+4))
	if v1693 <= int32(0) {
		goto L406
	} else {
		goto L408
	}
L408:
	;
	v1705 = int32(0)
	goto L409
L409:
	;
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v1690)+12))
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v1718+v1705<<(uint(int32(2))%32))))
	v1724 = F_table_openrv(m, v1722, int32(1))
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L8
	} else {
		goto L413
	}
L410:
	;
	goto L406
L411:
	;
	F_sequence_close(m, v1724, int32(0))
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L8
	} else {
		goto L438
	}
L412:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L8
	} else {
		goto L434
	}
L413:
	;
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v1724)+48))
	v1727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1726)+119)))
	v1729 = v1727 - int32(102)
	if base.Ui32(int32(12)) < base.Ui32(v1729) {
		goto L412
	} else {
		goto L414
	}
L414:
	;
	if int32(1)<<(uint(v1729)%32)&int32(5121) == int32(0) {
		goto L412
	} else {
		goto L415
	}
L415:
	;
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v1724)+52))
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(v1738)))
	if v1739 <= int32(0) {
		goto L411
	} else {
		goto L416
	}
L416:
	;
	v1749 = int32(0)
	goto L417
L417:
	;
	v1771 = v1738 + v1739<<(uint(int32(4))%32) + int32(20) + v1749*int32(100)
	v1772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1771)+91)))
	if v1772 == int32(0) {
		goto L420
	} else {
		goto L421
	}
L418:
	;
	F_sequence_close(m, v1724, int32(0))
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L8
	} else {
		goto L433
	}
L419:
	;
	goto L418
L420:
	;
	v1776 = v1771 + int32(4)
	v1779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1776))))
	v1780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1573))))
	if v1780 == int32(0) {
		v1799 = v1779
		v1800 = v1780
		goto L424
	} else {
		goto L425
	}
L421:
	;
	goto L422
L422:
	;
	v1805 = v1749 + int32(1)
	if v1739 != v1805 {
		v1749 = v1805
		goto L417
	} else {
		goto L432
	}
L423:
	;
	if v1800-v1799 == int32(0) {
		goto L419
	} else {
		goto L431
	}
L424:
	;
	goto L423
L425:
	;
	if v1779 != v1780 {
		v1799 = v1779
		v1800 = v1780
		goto L424
	} else {
		goto L426
	}
L426:
	;
	v1784 = v1573
	v1785 = v1776
	goto L427
L427:
	;
	v1788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1785)+1)))
	v1789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1784)+1)))
	if v1789 == int32(0) {
		v1799 = v1788
		v1800 = v1789
		goto L424
	} else {
		goto L429
	}
L428:
	;
	v1799 = v1788
	v1800 = v1789
	goto L424
L429:
	;
	v1792 = int32(1)
	if v1788 == v1789 {
		v1784 = v1784 + v1792
		v1785 = v1785 + v1792
		goto L427
	} else {
		goto L430
	}
L430:
	;
	goto L428
L431:
	;
	goto L422
L432:
	;
	goto L411
L433:
	;
	goto L370
L434:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1816 = m.ExcPending
	if v1816 != 0 {
		goto L8
	} else {
		goto L435
	}
L435:
	;
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(v1722)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v1817
	F_errmsg(m, int32(410624), v36+int32(32))
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L8
	} else {
		goto L436
	}
L436:
	;
	F_errfinish(m, int32(522528), int32(2888), int32(96590))
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L8
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
	v1854 = v1705 + int32(1)
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v1690)+4))
	if v1854 < v1855 {
		v1705 = v1854
		goto L409
	} else {
		goto L439
	}
L439:
	;
	goto L410
L440:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L8
	} else {
		goto L441
	}
L441:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v1885 = m.ExcPending
	if v1885 != 0 {
		goto L8
	} else {
		goto L442
	}
L442:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v1573
	F_errmsg(m, int32(74432), v36+int32(16))
	mBase = m.M
	v1891 = m.ExcPending
	if v1891 != 0 {
		goto L8
	} else {
		goto L443
	}
L443:
	;
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1892, v1893)
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L8
	} else {
		goto L444
	}
L444:
	;
	F_errfinish(m, int32(522528), int32(2919), int32(96590))
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		goto L8
	} else {
		goto L445
	}
L445:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L446:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1923))) = int32(92)
	v1927 = F_pstrdup(m, v1573)
	mBase = m.M
	v1928 = m.ExcPending
	if v1928 != 0 {
		goto L8
	} else {
		goto L447
	}
L447:
	;
	v1929 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1923)+8)) = v1929
	*(*int32)(unsafe.Add(mBase, uint32(v1923)+4)) = v1927
	*(*int64)(unsafe.Add(mBase, uint32(v1923)+16)) = v1929
	*(*int32)(unsafe.Add(mBase, uint32(v1923)+24)) = int32(0)
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
	v1937 = F_lappend(m, v1936, v1923)
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L8
	} else {
		goto L448
	}
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+24)) = v1937
	v1941 = v1558 + int32(1)
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v1540)+4))
	if v1941 < v1942 {
		v1558 = v1941
		goto L368
	} else {
		goto L449
	}
L449:
	;
	goto L369
L450:
	;
	v1968 = v52 + int32(1)
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v1968 < v1969 {
		v48 = v1965
		v52 = v1968
		goto L6
	} else {
		goto L451
	}
L451:
	;
	goto L7
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1975)+12)) = v1992
	*(*int32)(unsafe.Add(mBase, uint32(v1975)+316)) = v1992
	v1998 = F_list_make1_impl(m, int32(1), v1975+int32(12))
	mBase = m.M
	v1999 = m.ExcPending
	if v1999 != 0 {
		goto L8
	} else {
		goto L455
	}
L453:
	;
	v2000 = v1985
	goto L454
L454:
	;
	if v1987 == int32(0) {
		v2200 = v2000
		goto L456
	} else {
		goto L457
	}
L455:
	;
	v2000 = v1998
	goto L454
L456:
	;
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(v1971)+52))
	v2208 = F_list_concat(m, v2207, v2200)
	mBase = m.M
	v2209 = m.ExcPending
	if v2209 != 0 {
		goto L8
	} else {
		goto L493
	}
L457:
	;
	v2003 = int32(0)
	v2004 = *(*int32)(unsafe.Add(mBase, uint32(v1987)+4))
	if v2004 <= v2003 {
		v2200 = v2000
		goto L456
	} else {
		goto L458
	}
L458:
	;
	v2014 = v2003
	v2021 = v2000
	goto L459
L459:
	;
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v1987)+12))
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(v2028+v2014<<(uint(int32(2))%32))))
	v2033 = *(*int32)(unsafe.Add(mBase, uint32(v1971)+56))
	if v2032 == v2033 {
		v2175 = v2021
		goto L461
	} else {
		goto L462
	}
L460:
	;
	v2200 = v2175
	goto L456
L461:
	;
	v2183 = v2014 + int32(1)
	v2184 = *(*int32)(unsafe.Add(mBase, uint32(v1987)+4))
	if v2183 < v2184 {
		v2014 = v2183
		v2021 = v2175
		goto L459
	} else {
		goto L492
	}
L462:
	;
	if v2021 == int32(0) {
		goto L463
	} else {
		goto L464
	}
L463:
	;
	v2159 = F_lappend(m, v2021, v2032)
	mBase = m.M
	v2160 = m.ExcPending
	if v2160 != 0 {
		goto L8
	} else {
		goto L491
	}
L464:
	;
	v2037 = int32(0)
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v2021)+4))
	if v2038 <= v2037 {
		goto L463
	} else {
		goto L465
	}
L465:
	;
	v2042 = v2037
	goto L466
L466:
	;
	v2062 = *(*int32)(unsafe.Add(mBase, uint32(v2032)+20))
	v2063 = *(*int32)(unsafe.Add(mBase, uint32(v2021)+12))
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(v2063+v2042<<(uint(int32(2))%32))))
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v2067)+20))
	v2069 = F_equal(m, v2062, v2068)
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L8
	} else {
		goto L469
	}
L467:
	;
	goto L463
L468:
	;
	v2135 = v2042 + int32(1)
	v2136 = *(*int32)(unsafe.Add(mBase, uint32(v2021)+4))
	if v2135 < v2136 {
		v2042 = v2135
		goto L466
	} else {
		goto L490
	}
L469:
	;
	if v2069 == int32(0) {
		goto L468
	} else {
		goto L470
	}
L470:
	;
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v2032)+24))
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(v2067)+24))
	v2075 = F_equal(m, v2073, v2074)
	mBase = m.M
	v2076 = m.ExcPending
	if v2076 != 0 {
		goto L8
	} else {
		goto L471
	}
L471:
	;
	if v2075 == int32(0) {
		goto L468
	} else {
		goto L472
	}
L472:
	;
	v2079 = *(*int32)(unsafe.Add(mBase, uint32(v2032)+32))
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v2067)+32))
	v2081 = F_equal(m, v2079, v2080)
	mBase = m.M
	v2082 = m.ExcPending
	if v2082 != 0 {
		goto L8
	} else {
		goto L473
	}
L473:
	;
	if v2081 == int32(0) {
		goto L468
	} else {
		goto L474
	}
L474:
	;
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(v2032)+36))
	v2086 = *(*int32)(unsafe.Add(mBase, uint32(v2067)+36))
	v2087 = F_equal(m, v2085, v2086)
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L8
	} else {
		goto L475
	}
L475:
	;
	if v2087 == int32(0) {
		goto L468
	} else {
		goto L476
	}
L476:
	;
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v2032)+12))
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(v2067)+12))
	v2095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2092))))
	v2096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2091))))
	if v2096 == int32(0) {
		v2115 = v2095
		v2116 = v2096
		goto L478
	} else {
		goto L479
	}
L477:
	;
	if v2116-v2115 != 0 {
		goto L468
	} else {
		goto L485
	}
L478:
	;
	goto L477
L479:
	;
	if v2095 != v2096 {
		v2115 = v2095
		v2116 = v2096
		goto L478
	} else {
		goto L480
	}
L480:
	;
	v2100 = v2091
	v2101 = v2092
	goto L481
L481:
	;
	v2104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2101)+1)))
	v2105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2100)+1)))
	if v2105 == int32(0) {
		v2115 = v2104
		v2116 = v2105
		goto L478
	} else {
		goto L483
	}
L482:
	;
	v2115 = v2104
	v2116 = v2105
	goto L478
L483:
	;
	v2108 = int32(1)
	if v2104 == v2105 {
		v2100 = v2100 + v2108
		v2101 = v2101 + v2108
		goto L481
	} else {
		goto L484
	}
L484:
	;
	goto L482
L485:
	;
	v2118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2032)+61)))
	v2119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2067)+61)))
	if v2118 != v2119 {
		goto L468
	} else {
		goto L486
	}
L486:
	;
	v2121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2032)+65)))
	v2122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2067)+65)))
	if v2121 != v2122 {
		goto L468
	} else {
		goto L487
	}
L487:
	;
	v2124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2032)+66)))
	v2125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2067)+66)))
	if v2124 != v2125 {
		goto L468
	} else {
		goto L488
	}
L488:
	;
	v2127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2067)+60)))
	v2128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2032)+60)))
	v2129 = v2127 | v2128
	*(*uint8)(unsafe.Add(mBase, uint32(v2067)+60)) = uint8(v2129)
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v2067)+4))
	if v2131 != 0 {
		v2175 = v2021
		goto L461
	} else {
		goto L489
	}
L489:
	;
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(v2032)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2067)+4)) = v2132
	v2175 = v2021
	goto L461
L490:
	;
	goto L467
L491:
	;
	v2175 = v2159
	goto L461
L492:
	;
	goto L460
L493:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1971)+52)) = v2208
	m.G0 = v1975 + int32(320)
	return
}
func F_transformIndexStmt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+67)))
	if v9 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L5
	} else {
		goto L34
	}
L2:
	;
	v13 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	return l1
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l2
	v19 = F_relation_open(m, l0, int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v21 = int32(1)
	v22 = int32(0)
	v25 = F_addRangeTableEntryForRelation(m, v13, v19, v21, v22, v22, v21)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v28 = int32(1)
	F_addNSItemToQuery(m, v13, v25, int32(0), v28, v28)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v32 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v35 = F_transformWhereClause(m, v13, v32, int32(33), int32(563654))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v41 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v35
	F_assign_expr_collations(m, v13, v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v102 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L16:
	;
	v44 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v45 <= v44 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v50 = v44
	goto L18
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v50<<(uint(int32(2))%32))))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v61 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L15
L20:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	if v62 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v91 = v50 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v91 < v92 {
		v50 = v91
		goto L18
	} else {
		goto L29
	}
L23:
	;
	v81 = v61
	goto L25
L24:
	;
	v63 = m.G0
	v65 = v63 - int32(16)
	m.G0 = v65
	*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = int32(0)
	v71 = F_FigureColnameInternal(m, v61, v65+int32(12))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L26
	}
L25:
	;
	v83 = F_transformExpr(m, v13, v81, int32(32))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L27
	}
L26:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	m.G0 = v65 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+12)) = v73
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v81 = v78
	goto L25
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v83
	F_assign_expr_collations(m, v13, v83)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	goto L22
L29:
	;
	goto L19
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	if v105 != int32(1) {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_free_parsestate(m, v13)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	F_sequence_close(m, v19, int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v113 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+67)) = uint8(v113)
	goto L4
L34:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	F_errmsg(m, int32(458200), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(522528), int32(3123), int32(103413))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
