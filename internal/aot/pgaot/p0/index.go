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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+8)))
	if base.Ui32(int32(_a_F_BuildIndexInfo_0)) < base.Ui32((v19-int32(33))&int32(_a_F_BuildIndexInfo_1)) {
		v27 = v18 + int32(48)
		v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+10)))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+84))
		v31 = F_RelationGetIndexExpressions(m, l0)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			v35 = F_RelationGetIndexPredicate(m, l0)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+12)))
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+13)))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+20)))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+28)))
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)))
				v45 = F_makeIndexInfo(m, v19, v28, v30, v31, v35, v37, v38, v39, int32(0), v42, v37&v43)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					v48 = v45 + int32(12)
					v49 = int32(0)
					if base.Ui32(int32(3)) <= base.Ui32(v19-int32(1)) {
						v58 = v49
						v67 = v2
						for {
							v70 = v58 << (uint(int32(1)) % 32)
							v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70+v27))))
							*(*uint16)(unsafe.Add(mBase, uint32(v48+v70))) = uint16(v73)
							v76 = v70 | int32(2)
							v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+v76))))
							*(*uint16)(unsafe.Add(mBase, uint32(v48+v76))) = uint16(v79)
							v81 = int32(4)
							v82 = v70 | v81
							v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+v82))))
							*(*uint16)(unsafe.Add(mBase, uint32(v48+v82))) = uint16(v85)
							v88 = v70 | int32(6)
							v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88+v27))))
							*(*uint16)(unsafe.Add(mBase, uint32(v48+v88))) = uint16(v91)
							v94 = v58 + v81
							v96 = v67 + v81
							if v96 != v19&int32(60) {
								v58 = v94
								v67 = v96
								continue
							} else {
								break
							}
							break
						}
						if v19&int32(3) == int32(0) {
						} else {
							v104 = v94
							v119 = v104
							v129 = v2
							for {
								v130 = int32(1)
								v131 = v119 << (uint(v130) % 32)
								v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131+v27))))
								*(*uint16)(unsafe.Add(mBase, uint32(v48+v131))) = uint16(v134)
								v139 = v129 + v130
								if v139 != v19&int32(3) {
									v119 = v119 + v130
									v129 = v139
									continue
								} else {
									break
								}
								break
							}
						}
					} else {
						v104 = v49
						v119 = v104
						v129 = v2
						for {
							v130 = int32(1)
							v131 = v119 << (uint(v130) % 32)
							v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131+v27))))
							*(*uint16)(unsafe.Add(mBase, uint32(v48+v131))) = uint16(v134)
							v139 = v129 + v130
							if v139 != v19&int32(3) {
								v119 = v119 + v130
								v129 = v139
								continue
							} else {
								break
							}
							break
						}
					}
					v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)))
					if v154 == int32(1) {
						F_RelationGetExclusionInfo(m, l0, v45+int32(92), v45+int32(96), v45+int32(100))
						mBase = m.M
						v164 = m.ExcPending
						if v164 != 0 {
							return int32(0)
						} else {
							m.G0 = v16 + int32(16)
							return v45
						}
					} else {
						m.G0 = v16 + int32(16)
						return v45
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v172 = m.ExcPending
		if v172 != 0 {
			return int32(0)
		} else {
			v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v173
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = v19
			F_errmsg_internal(m, int32(_a_F_BuildIndexInfo_2), v16)
			mBase = m.M
			v178 = m.ExcPending
			if v178 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_BuildIndexInfo_3), int32(2439), int32(_a_F_BuildIndexInfo_4))
				mBase = m.M
				v183 = m.ExcPending
				if v183 != 0 {
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v4 = int32(0)
	v7 = l0 + l1<<(uint(int32(2))%32)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if base.B2i32(v8 < v4)|base.B2i32(l2 <= v8) == v4 {
		if int32(0) < l1 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v7-int32(4))))
			if v8 < v19 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_CheckIndex_0), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_CheckIndex_1), int32(124), int32(_a_F_CheckIndex_2))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
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
				if v8 == v19 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						F_errcode(m, int32(130))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_CheckIndex_3), int32(0))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_CheckIndex_1), int32(129), int32(_a_F_CheckIndex_2))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
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
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			F_errcode(m, int32(130))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_CheckIndex_4), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_CheckIndex_1), int32(116), int32(_a_F_CheckIndex_2))
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
}
func F_CheckIndexCompatible(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
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
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
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
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
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
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v556 int32
	_ = v556
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v610 int32
	_ = v610
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v664 int32
	_ = v664
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v719 int32
	_ = v719
	v6 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v21 = F_IndexGetRelation(m, l0, v6)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
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
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v26 = v25
	goto L5
L4:
	;
	v26 = v6
	goto L5
L5:
	;
	v28 = F_SearchSysCache1(m, int32(1), l1)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	m.G0 = v18 + int32(48)
	return v719
L7:
	;
	F_ReleaseCatCache(m, v70)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L1
	} else {
		goto L180
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L1
	} else {
		goto L177
	}
L9:
	;
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
	v32 = v30 + v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+68))
	v35 = F_GetIndexAmRoutine(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
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
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L173
	}
L13:
	;
	F_ReleaseCatCache(m, v28)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+10)))
	v40 = int32(0)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+28)))
	v47 = F_makeIndexInfo(m, v26, v26, v33, v40, v40, v40, v40, v40, v40, v46, l4)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v50 = v26 << (uint(int32(2)) % 32)
	v51 = F_palloc(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v53 = F_palloc(m, v50)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v55 = F_palloc(m, v50)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v57 = F_palloc(m, v50)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v61 = F_palloc(m, v26<<(uint(int32(1))%32))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v63 = int32(0)
	F_ComputeIndexAttrs(m, v47, v51, v53, v55, v57, v61, l2, l3, v21, l1, v33, v39, v63, l4, v63, v63, v63)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v70 = F_SearchSysCache1(m, int32(34), l0)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v70 == int32(0) {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+22)))
	v78 = F_heap_attisnull(m, v70, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v78 == int32(0) {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v84 = F_heap_attisnull(m, v70, int32(20), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v84 == int32(0) {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	v88 = v75 + v74
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+18)))
	if v89 == int32(0) {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+10)))
	v95 = F_SysCacheGetAttrNotNull(m, int32(34), v70, int32(17))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v99 = F_SysCacheGetAttrNotNull(m, int32(34), v70, int32(18))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v102 = v99 + int32(24)
	v104 = v92 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v104) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	if v166 != 0 {
		goto L7
	} else {
		goto L49
	}
L32:
	;
	v166 = int32(0)
	goto L31
L33:
	;
	v140 = v135
	v141 = v136
	v142 = v137
	goto L43
L34:
	;
	if (v102|v55)&int32(3) != 0 {
		v135 = v102
		v136 = v55
		v137 = v104
		goto L33
	} else {
		goto L37
	}
L35:
	;
	v128 = v102
	v129 = v55
	v130 = v104
	goto L36
L36:
	;
	if v130 == int32(0) {
		goto L32
	} else {
		goto L42
	}
L37:
	;
	v112 = v102
	v113 = v55
	v114 = v104
	goto L38
L38:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	if v117 != v118 {
		v135 = v112
		v136 = v113
		v137 = v114
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v128 = v123
	v129 = v121
	v130 = v125
	goto L36
L40:
	;
	v120 = int32(4)
	v121 = v113 + v120
	v123 = v112 + v120
	v125 = v114 - v120
	if base.Ui32(int32(3)) < base.Ui32(v125) {
		v112 = v123
		v113 = v121
		v114 = v125
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v135 = v128
	v136 = v129
	v137 = v130
	goto L33
L43:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v145 == v146 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v166 = v145 - v146
	goto L31
L45:
	;
	v148 = int32(1)
	v153 = v142 - v148
	if v153 != 0 {
		v140 = v140 + v148
		v141 = v141 + v148
		v142 = v153
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
	v168 = v95 + int32(24)
	if base.Ui32(int32(4)) <= base.Ui32(v104) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	F_ReleaseCatCache(m, v70)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L68
	}
L51:
	;
	v230 = int32(0)
	goto L50
L52:
	;
	v204 = v199
	v205 = v200
	v206 = v201
	goto L62
L53:
	;
	if (v168|v53)&int32(3) != 0 {
		v199 = v168
		v200 = v53
		v201 = v104
		goto L52
	} else {
		goto L56
	}
L54:
	;
	v192 = v168
	v193 = v53
	v194 = v104
	goto L55
L55:
	;
	if v194 == int32(0) {
		goto L51
	} else {
		goto L61
	}
L56:
	;
	v176 = v168
	v177 = v53
	v178 = v104
	goto L57
L57:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	if v181 != v182 {
		v199 = v176
		v200 = v177
		v201 = v178
		goto L52
	} else {
		goto L59
	}
L58:
	;
	v192 = v187
	v193 = v185
	v194 = v189
	goto L55
L59:
	;
	v184 = int32(4)
	v185 = v177 + v184
	v187 = v176 + v184
	v189 = v178 - v184
	if base.Ui32(int32(3)) < base.Ui32(v189) {
		v176 = v187
		v177 = v185
		v178 = v189
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v199 = v192
	v200 = v193
	v201 = v194
	goto L52
L62:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	if v209 == v210 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v230 = v209 - v210
	goto L50
L64:
	;
	v212 = int32(1)
	v217 = v206 - v212
	if v217 != 0 {
		v204 = v204 + v212
		v205 = v205 + v212
		v206 = v217
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
	v233 = int32(0)
	if v230 != 0 {
		v719 = v233
		goto L6
	} else {
		goto L69
	}
L69:
	;
	v235 = F_index_open(m, l0, int32(1))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	if v92 <= int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	F_relation_close(m, v235, int32(0))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L172
	}
L72:
	;
	v374 = int32(0)
	v375 = m.G0
	v377 = v375 - int32(32)
	m.G0 = v377
	if v361|v57 == v374 {
		goto L111
	} else {
		goto L112
	}
L73:
	;
	v239 = F_palloc(m, v104)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v246 = v233
	goto L77
L76:
	;
	v361 = v239
	goto L72
L77:
	;
	v257 = v246 << (uint(int32(2)) % 32)
	v258 = v55 + v257
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v260 = F_get_opclass_input_type(m, v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L81
	}
L78:
	;
	v332 = F_palloc(m, v104)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L105
	}
L79:
	;
	v329 = v246 + int32(1)
	if v329 != v92 {
		v246 = v329
		goto L77
	} else {
		goto L104
	}
L80:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v235)+52))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v314+v315<<(uint(int32(4))%32)+v246*int32(100))+88))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v257+v51)))
	if v322 == v324 {
		goto L79
	} else {
		goto L103
	}
L81:
	;
	if v260 == int32(2283) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v265 = F_get_opclass_input_type(m, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	if v265 == int32(2277) {
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v270 = F_get_opclass_input_type(m, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v270 == int32(2776) {
		goto L80
	} else {
		goto L86
	}
L86:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v275 = F_get_opclass_input_type(m, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	if v275 == int32(3500) {
		goto L80
	} else {
		goto L88
	}
L88:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v280 = F_get_opclass_input_type(m, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	if v280 == int32(3831) {
		goto L80
	} else {
		goto L90
	}
L90:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v285 = F_get_opclass_input_type(m, v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	if v285 == int32(_a_F_CheckIndexCompatible_0) {
		goto L80
	} else {
		goto L92
	}
L92:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v290 = F_get_opclass_input_type(m, v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	if v290 == int32(_a_F_CheckIndexCompatible_1) {
		goto L80
	} else {
		goto L94
	}
L94:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v295 = F_get_opclass_input_type(m, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	if v295 == int32(_a_F_CheckIndexCompatible_2) {
		goto L80
	} else {
		goto L96
	}
L96:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v300 = F_get_opclass_input_type(m, v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	if v300 == int32(_a_F_CheckIndexCompatible_3) {
		goto L80
	} else {
		goto L98
	}
L98:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v305 = F_get_opclass_input_type(m, v304)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	if v305 == int32(_a_F_CheckIndexCompatible_4) {
		goto L80
	} else {
		goto L100
	}
L100:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v310 = F_get_opclass_input_type(m, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	if v310 != int32(_a_F_CheckIndexCompatible_5) {
		goto L79
	} else {
		goto L102
	}
L102:
	;
	goto L80
L103:
	;
	v664 = int32(0)
	goto L71
L104:
	;
	goto L78
L105:
	;
	v339 = int32(0)
	goto L106
L106:
	;
	v353 = v339 + int32(1)
	v355 = F_get_attoptions(m, l0, base.I32_extend16_s(v353))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L108
	}
L107:
	;
	v361 = v332
	goto L72
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v332+v339<<(uint(int32(2))%32)))) = v355
	if v353 != v92 {
		v339 = v353
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	m.G0 = v377 + int32(32)
	F_pfree(m, v361)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L134
	}
L111:
	;
	v452 = int32(1)
	goto L110
L112:
	;
	goto L113
L113:
	;
	F_fmgr_info(m, int32(744), v377+int32(4))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	if v92 <= int32(0) {
		v452 = int32(1)
		goto L110
	} else {
		goto L115
	}
L115:
	;
	v395 = v374
	goto L116
L116:
	;
	if v361 != 0 {
		goto L122
	} else {
		goto L123
	}
L117:
	;
	v452 = v445
	goto L110
L118:
	;
	v445 = int32(1)
	v447 = v395 + v445
	if v447 != v92 {
		v395 = v447
		goto L116
	} else {
		goto L133
	}
L119:
	;
	v452 = int32(0)
	goto L110
L120:
	;
	if v409 == int32(0) {
		goto L118
	} else {
		goto L132
	}
L121:
	;
	if v434 == int32(0) {
		goto L118
	} else {
		goto L131
	}
L122:
	;
	v407 = v395 << (uint(int32(2)) % 32)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v361+v407)))
	if v57 == int32(0) {
		goto L120
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	if v57 == int32(0) {
		goto L118
	} else {
		goto L130
	}
L125:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v407+v57)))
	if v409 == int32(0) {
		v434 = v413
		goto L121
	} else {
		goto L126
	}
L126:
	;
	v416 = int32(0)
	if v413 == v416 {
		v452 = v416
		goto L110
	} else {
		goto L127
	}
L127:
	;
	v422 = F_FunctionCall2Coll(m, v377+int32(4), int32(950), v409, v413)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	if v422 == int32(0) {
		v452 = v416
		goto L110
	} else {
		goto L129
	}
L129:
	;
	goto L118
L130:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v57+v395<<(uint(int32(2))%32))))
	v434 = v431
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
	if v452 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v664 = int32(0)
	goto L71
L136:
	;
	goto L137
L137:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	if v472 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v664 = int32(1)
	goto L71
L139:
	;
	goto L140
L140:
	;
	F_RelationGetExclusionInfo(m, v235, v18+int32(44), v18+int32(40), v18+int32(36))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	if base.Ui32(int32(4)) <= base.Ui32(v104) {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	v548 = int32(0)
	if v547|base.B2i32(v92 <= v548) != 0 {
		v664 = base.B2i32(v547 == v548)
		goto L71
	} else {
		goto L160
	}
L143:
	;
	v547 = int32(0)
	goto L142
L144:
	;
	v521 = v516
	v522 = v517
	v523 = v518
	goto L154
L145:
	;
	if (v484|v485)&int32(3) != 0 {
		v516 = v484
		v517 = v485
		v518 = v104
		goto L144
	} else {
		goto L148
	}
L146:
	;
	v509 = v484
	v510 = v485
	v511 = v104
	goto L147
L147:
	;
	if v511 == int32(0) {
		goto L143
	} else {
		goto L153
	}
L148:
	;
	v493 = v484
	v494 = v485
	v495 = v104
	goto L149
L149:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v493)))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	if v498 != v499 {
		v516 = v493
		v517 = v494
		v518 = v495
		goto L144
	} else {
		goto L151
	}
L150:
	;
	v509 = v504
	v510 = v502
	v511 = v506
	goto L147
L151:
	;
	v501 = int32(4)
	v502 = v494 + v501
	v504 = v493 + v501
	v506 = v495 - v501
	if base.Ui32(int32(3)) < base.Ui32(v506) {
		v493 = v504
		v494 = v502
		v495 = v506
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	v516 = v509
	v517 = v510
	v518 = v511
	goto L144
L154:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521))))
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522))))
	if v526 == v527 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v547 = v526 - v527
	goto L142
L156:
	;
	v529 = int32(1)
	v534 = v523 - v529
	if v534 != 0 {
		v521 = v521 + v529
		v522 = v522 + v529
		v523 = v534
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
	v556 = int32(0)
	goto L161
L161:
	;
	v570 = v556 << (uint(int32(2)) % 32)
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v570+v571)))
	F_op_input_types(m, v573, v18+int32(32), v18+int32(28))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L163
	}
L162:
	;
	v664 = v655
	goto L71
L163:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	if base.B2i32(base.Ui32(v580-int32(_a_F_CheckIndexCompatible_1)) < base.Ui32(int32(2)))|base.B2i32(v580 == int32(_a_F_CheckIndexCompatible_0))|(base.B2i32(v580 == int32(3831))|base.B2i32(v580 == int32(3500)))|(base.B2i32(v580 == int32(2776))|base.B2i32(v580 == int32(2283))|(base.B2i32(v580 == int32(2277))|base.B2i32(v580 == int32(_a_F_CheckIndexCompatible_5)))) != 0 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v655 = int32(1)
	v657 = v556 + v655
	if v657 != v92 {
		v556 = v657
		goto L161
	} else {
		goto L171
	}
L165:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v235)+52))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v641+v642<<(uint(int32(4))%32)+v556*int32(100))+88))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v570+v51)))
	if v649 == v651 {
		goto L164
	} else {
		goto L170
	}
L166:
	;
	if base.Ui32(v580-int32(_a_F_CheckIndexCompatible_3)) < base.Ui32(int32(2)) {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if base.B2i32(base.Ui32(v610-int32(_a_F_CheckIndexCompatible_1)) < base.Ui32(int32(2)))|base.B2i32(v610 == int32(_a_F_CheckIndexCompatible_0))|(base.B2i32(v610 == int32(3831))|base.B2i32(v610 == int32(3500)))|(base.B2i32(v610 == int32(2776))|base.B2i32(v610 == int32(2283))|(base.B2i32(v610 == int32(2277))|base.B2i32(v610 == int32(_a_F_CheckIndexCompatible_5)))) != 0 {
		goto L165
	} else {
		goto L168
	}
L168:
	;
	if base.Ui32(int32(1)) < base.Ui32(v610-int32(_a_F_CheckIndexCompatible_3)) {
		goto L164
	} else {
		goto L169
	}
L169:
	;
	goto L165
L170:
	;
	v664 = int32(0)
	goto L71
L171:
	;
	goto L162
L172:
	;
	v719 = v664
	goto L6
L173:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l1
	F_errmsg(m, int32(_a_F_CheckIndexCompatible_6), v18)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	F_errfinish(m, int32(_a_F_CheckIndexCompatible_7), int32(227), int32(_a_F_CheckIndexCompatible_8))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_CheckIndexCompatible_9), v18+int32(16))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(_a_F_CheckIndexCompatible_7), int32(263), int32(_a_F_CheckIndexCompatible_8))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L180:
	;
	v719 = int32(0)
	goto L6
}
func F_ExecIndexBuildScanKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
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
	var v60 int32
	_ = v60
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
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
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
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
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
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
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int64
	_ = v317
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
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
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v438 int64
	_ = v438
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
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
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
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
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v606 int32
	_ = v606
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	v11 = int32(0)
	v38 = m.G0
	v40 = v38 - int32(48)
	m.G0 = v40
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v44 = v42
	goto L3
L2:
	;
	v44 = int32(0)
	goto L3
L3:
	;
	v47 = F_palloc(m, v44*int32(48))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v53 = F_palloc0(m, v44*int32(24))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
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
	v961 = m.ExcPending
	if v961 != 0 {
		goto L4
	} else {
		goto L199
	}
L8:
	;
	m.G0 = v40 + int32(48)
	return
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v903
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v902
	goto L8
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v47
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v44
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v810
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v809
	if l8 == int32(0) {
		goto L7
	} else {
		goto L198
	}
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v55 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v843 = v49
	v847 = v50
	goto L13
L13:
	;
	F_pfree(m, v53)
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L4
	} else {
		goto L196
	}
L14:
	;
	if v814 != 0 {
		goto L10
	} else {
		goto L195
	}
L15:
	;
	v809 = v49
	v810 = v50
	v814 = v11
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
	v60 = int32(256)
	goto L20
L19:
	;
	v60 = int32(0)
	goto L20
L20:
	;
	v77 = v47
	v78 = v49
	v81 = v49
	v82 = v50
	v86 = v11
	v88 = v11
	goto L30
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L4
	} else {
		goto L192
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L4
	} else {
		goto L189
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L4
	} else {
		goto L186
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L4
	} else {
		goto L183
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L4
	} else {
		goto L180
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L4
	} else {
		goto L177
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L4
	} else {
		goto L174
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L4
	} else {
		goto L171
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L4
	} else {
		goto L168
	}
L30:
	;
	v102 = v47 + v88*int32(48)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v104 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+10)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105+v88<<(uint(int32(2))%32))))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	switch v110 - int32(17) {
	case 0:
		goto L37
	default:
		goto L21
	case 3:
		goto L35
	case 20:
		goto L36
	case 35:
		goto L34
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L4
	} else {
		goto L165
	}
L32:
	;
	goto L31
L33:
	;
	v645 = v88 + int32(1)
	v649 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v645 < v649 {
		v77 = v47 + v645*int32(48)
		v78 = v622
		v81 = v625
		v82 = v626
		v86 = v630
		v88 = v645
		goto L30
	} else {
		goto L164
	}
L34:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v566)))
	if v567 == int32(27) {
		goto L152
	} else {
		goto L153
	}
L35:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v109)+28))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)+12))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v456)))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	if v458 == int32(27) {
		goto L120
	} else {
		goto L121
	}
L36:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	if v207 != 0 {
		goto L62
	} else {
		goto L63
	}
L37:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)+28))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v116 == int32(27) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v121 = v119
	v122 = v120
	goto L40
L39:
	;
	v121 = v115
	v122 = v116
	goto L40
L40:
	;
	if v122 != int32(6) {
		goto L32
	} else {
		goto L41
	}
L41:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	if v125 != int32(-3) {
		goto L32
	} else {
		goto L42
	}
L42:
	;
	v128 = int32(*(*int16)(unsafe.Add(mBase, uint32(v121)+8)))
	if base.B2i32(v128 <= int32(0))|base.B2i32(v104 < v128) != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v135+v128<<(uint(int32(2))%32)-int32(4))))
	F_get_op_opfamily_properties(m, v134, v141, l3, v40+int32(44), v40+int32(40), v40+int32(36))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v109)+28))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+12))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	if v153 == int32(27) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+44)))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v40)+36))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v109)+24))
	F_ScanKeyEntryInitialize(m, v102, v196, v128, v201, v202, v203, v133, v200)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L61
	}
L46:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v158 = v156
	v159 = v157
	goto L48
L47:
	;
	v158 = v152
	v159 = v153
	goto L48
L48:
	;
	if v159 == int32(7) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+24)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v158)+20))
	v196 = v60 | v162
	v197 = v78
	v198 = v81
	v199 = v82
	v200 = v164
	goto L45
L50:
	;
	goto L51
L51:
	;
	if v81 < v78 {
		v178 = v78
		v179 = v82
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v182 = v179 + v81*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v102
	v184 = F_ExecInitExpr(m, v158, l0)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L59
	}
L53:
	;
	if v78 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v170 = F_palloc(m, int32(96))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v174 = F_repalloc(m, v82, v78*int32(24))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L58
	}
L57:
	;
	v178 = int32(8)
	v179 = v170
	goto L52
L58:
	;
	v178 = v78 << (uint(int32(1)) % 32)
	v179 = v174
	goto L52
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+4)) = v184
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v40)+36))
	v188 = F_get_typstorage(m, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+8)) = uint8(base.B2i32(v188 != int32(112)))
	v196 = v60
	v197 = v178
	v198 = v81 + int32(1)
	v199 = v179
	v200 = int32(0)
	goto L45
L61:
	;
	v622 = v197
	v625 = v198
	v626 = v199
	v630 = v86
	goto L33
L62:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	v212 = v208 * int32(48)
	goto L64
L63:
	;
	v212 = int32(0)
	goto L64
L64:
	;
	v213 = F_palloc(m, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v109)+24))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v109)+20))
	v229 = int32(0)
	v234 = v78
	v237 = v81
	v238 = v82
	goto L67
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = int32(4)
	v450 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v213)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v102)+4)) = uint16(v450)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v102)+44)) = v213
	*(*uint16)(unsafe.Add(mBase, uint32(v102)+6)) = uint16(v452)
	v622 = v234
	v625 = v237
	v626 = v238
	v630 = v86
	goto L33
L67:
	;
	v256 = int32(0)
	if v218 == v256 {
		v267 = v256
		goto L69
	} else {
		goto L70
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102)+44)) = int32(0)
	v438 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v102)+36)) = v438
	*(*int64)(unsafe.Add(mBase, uint32(v102)+28)) = v438
	*(*int64)(unsafe.Add(mBase, uint32(v102)+20)) = v438
	*(*int64)(unsafe.Add(mBase, uint32(v102)+12)) = v438
	*(*int64)(unsafe.Add(mBase, uint32(v102)+4)) = v438
	goto L66
L69:
	;
	if v217 == int32(0) {
		v276 = v256
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	if v261 <= v229 {
		v267 = int32(0)
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v218)+12))
	v267 = v263 + v229<<(uint(int32(2))%32)
	goto L69
L72:
	;
	v277 = int32(0)
	if v216 == v277 {
		v288 = v277
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	if v270 <= v229 {
		v276 = v256
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	v276 = v272 + v229<<(uint(int32(2))%32)
	goto L72
L75:
	;
	if v215 == int32(0) {
		v297 = v277
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	if v282 <= v229 {
		v288 = int32(0)
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	v288 = v284 + v229<<(uint(int32(2))%32)
	goto L75
L78:
	;
	v300 = v213 + v229*int32(48)
	v301 = int32(0)
	if v297 != 0 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v291 <= v229 {
		v297 = v277
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v297 = v293 + v229<<(uint(int32(2))%32)
	goto L78
L81:
	;
	goto L68
L82:
	;
	v310 = base.B2i32(v288 == v301) | (base.B2i32(v267 == v301) | base.B2i32(v276 == v301))
	goto L84
L83:
	;
	v310 = int32(1)
	goto L84
L84:
	;
	if v310 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v312 = v300 - int32(48)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	*(*int32)(unsafe.Add(mBase, uint32(v312))) = v313 | int32(16)
	if v47&int32(3) != 0 {
		goto L81
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	if v330 == int32(27) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v317 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v77)+40)) = v317
	*(*int64)(unsafe.Add(mBase, uint32(v77)+32)) = v317
	*(*int64)(unsafe.Add(mBase, uint32(v77)+24)) = v317
	*(*int64)(unsafe.Add(mBase, uint32(v77)+16)) = v317
	*(*int64)(unsafe.Add(mBase, uint32(v77)+8)) = v317
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v317
	goto L66
L89:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	v335 = v333
	v336 = v334
	goto L91
L90:
	;
	v335 = v329
	v336 = v330
	goto L91
L91:
	;
	if v336 != int32(6) {
		goto L28
	} else {
		goto L92
	}
L92:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v335)+4))
	if v339 != int32(-3) {
		goto L28
	} else {
		goto L93
	}
L93:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+10)))
	if v343 != int32(1) {
		goto L27
	} else {
		goto L94
	}
L94:
	;
	v346 = int32(*(*int16)(unsafe.Add(mBase, uint32(v335)+8)))
	if base.B2i32(v346 <= int32(0))|base.B2i32(v104 < v346) != 0 {
		goto L27
	} else {
		goto L95
	}
L95:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v354+v346<<(uint(int32(2))%32)-int32(4))))
	F_get_op_opfamily_properties(m, v353, v360, l3, v40+int32(44), v40+int32(40), v40+int32(36))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v40)+44))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v369 != v370 {
		goto L26
	} else {
		goto L97
	}
L97:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v40)+36))
	v375 = F_get_opfamily_proc(m, v360, v372, v373, int32(1))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	if v375 == int32(0) {
		goto L25
	} else {
		goto L99
	}
L99:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v352)))
	if v379 == int32(27) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v430 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+44)))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v40)+36))
	F_ScanKeyEntryInitialize(m, v300, v425, v346, v430, v431, v351, v375, v429)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L4
	} else {
		goto L119
	}
L101:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v384 = v382
	v385 = v383
	goto L103
L102:
	;
	v384 = v352
	v385 = v379
	goto L103
L103:
	;
	if v385 == int32(7) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+24)))
	if v390 != 0 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	if v237 < v234 {
		v406 = v234
		v407 = v238
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v391 = int32(9)
	goto L109
L108:
	;
	v391 = int32(8)
	goto L109
L109:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v384)+20))
	v425 = v391
	v426 = v234
	v427 = v237
	v428 = v238
	v429 = v392
	goto L100
L110:
	;
	v410 = v407 + v237*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v410))) = v300
	v412 = F_ExecInitExpr(m, v384, l0)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L4
	} else {
		goto L117
	}
L111:
	;
	if v234 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v398 = F_palloc(m, int32(96))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L4
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v402 = F_repalloc(m, v238, v234*int32(24))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L4
	} else {
		goto L116
	}
L115:
	;
	v406 = int32(8)
	v407 = v398
	goto L110
L116:
	;
	v406 = v234 << (uint(int32(1)) % 32)
	v407 = v402
	goto L110
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v410)+4)) = v412
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v40)+36))
	v416 = F_get_typstorage(m, v415)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v410)+8)) = uint8(base.B2i32(v416 != int32(112)))
	v425 = int32(8)
	v426 = v406
	v427 = v237 + int32(1)
	v428 = v407
	v429 = int32(0)
	goto L100
L119:
	;
	v229 = v229 + int32(1)
	v234 = v426
	v237 = v427
	v238 = v428
	goto L67
L120:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v457)+4))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)))
	v463 = v461
	v464 = v462
	goto L122
L121:
	;
	v463 = v457
	v464 = v458
	goto L122
L122:
	;
	if v464 != int32(6) {
		goto L24
	} else {
		goto L123
	}
L123:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v463)+4))
	if v467 != int32(-3) {
		goto L24
	} else {
		goto L124
	}
L124:
	;
	v470 = int32(*(*int16)(unsafe.Add(mBase, uint32(v463)+8)))
	if base.B2i32(v470 <= int32(0))|base.B2i32(v104 < v470) != 0 {
		goto L23
	} else {
		goto L125
	}
L125:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v477+v470<<(uint(int32(2))%32)-int32(4))))
	F_get_op_opfamily_properties(m, v476, v483, l3, v40+int32(44), v40+int32(40), v40+int32(36))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v109)+28))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+12))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)+4))
	if v494 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503)+19)))
	if v504 == int32(1) {
		goto L133
	} else {
		goto L134
	}
L128:
	;
	v502 = int32(0)
	goto L127
L129:
	;
	goto L130
L130:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	if v498 != int32(27) {
		v502 = v494
		goto L127
	} else {
		goto L131
	}
L131:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	v502 = v501
	goto L127
L132:
	;
	v561 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+44)))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v40)+36))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v109)+24))
	F_ScanKeyEntryInitialize(m, v102, v555, v470, v561, v562, v563, v475, v560)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L4
	} else {
		goto L151
	}
L133:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v502)))
	if v507 == int32(7) {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	goto L135
L135:
	;
	v545 = v53 + v86*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v545))) = v102
	v547 = F_ExecInitExpr(m, v502, l0)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L4
	} else {
		goto L150
	}
L136:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502)+24)))
	if v512 != 0 {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	goto L138
L138:
	;
	if v81 < v78 {
		v528 = v78
		v529 = v82
		goto L142
	} else {
		goto L143
	}
L139:
	;
	v513 = int32(33)
	goto L141
L140:
	;
	v513 = int32(32)
	goto L141
L141:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v502)+20))
	v555 = v513
	v556 = v78
	v557 = v81
	v558 = v82
	v559 = v86
	v560 = v514
	goto L132
L142:
	;
	v532 = v529 + v81*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v532))) = v102
	v534 = F_ExecInitExpr(m, v502, l0)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L4
	} else {
		goto L149
	}
L143:
	;
	if v78 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v520 = F_palloc(m, int32(96))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L4
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v524 = F_repalloc(m, v82, v78*int32(24))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L4
	} else {
		goto L148
	}
L147:
	;
	v528 = int32(8)
	v529 = v520
	goto L142
L148:
	;
	v528 = v78 << (uint(int32(1)) % 32)
	v529 = v524
	goto L142
L149:
	;
	v536 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v532)+8)) = uint8(v536)
	*(*int32)(unsafe.Add(mBase, uint32(v532)+4)) = v534
	v555 = int32(32)
	v556 = v528
	v557 = v81 + v536
	v558 = v529
	v559 = v86
	v560 = int32(0)
	goto L132
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v545)+4)) = v547
	v552 = int32(0)
	v555 = v552
	v556 = v78
	v557 = v81
	v558 = v82
	v559 = v86 + int32(1)
	v560 = v552
	goto L132
L151:
	;
	v622 = v556
	v625 = v557
	v626 = v558
	v630 = v559
	goto L33
L152:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	v572 = v570
	v573 = v571
	goto L154
L153:
	;
	v572 = v566
	v573 = v567
	goto L154
L154:
	;
	if v573 != int32(6) {
		goto L22
	} else {
		goto L155
	}
L155:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	if v576 != int32(-3) {
		goto L22
	} else {
		goto L156
	}
L156:
	;
	v579 = int32(*(*int16)(unsafe.Add(mBase, uint32(v572)+8)))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	switch v581 {
	case 0:
		v599 = int32(65)
		goto L157
	case 1:
		goto L158
	default:
		goto L159
	}
L157:
	;
	v600 = int32(0)
	F_ScanKeyEntryInitialize(m, v102, v599, v579, v600, v600, v600, v600, v600)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L4
	} else {
		goto L163
	}
L158:
	;
	v599 = int32(129)
	goto L157
L159:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v586
	F_errmsg_internal(m, int32(_a_F_ExecIndexBuildScanKeys_0), v40+int32(32))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1607), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	v622 = v78
	v625 = v81
	v626 = v82
	v630 = v86
	goto L33
L164:
	;
	v809 = v625
	v810 = v626
	v814 = v630
	goto L14
L165:
	;
	F_errmsg_internal(m, int32(_a_F_ExecIndexBuildScanKeys_3), int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1235), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L168:
	;
	F_errmsg_internal(m, int32(_a_F_ExecIndexBuildScanKeys_4), int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1239), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_ExecIndexBuildScanKeys_3), int32(0))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1352), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L174:
	;
	F_errmsg_internal(m, int32(_a_F_ExecIndexBuildScanKeys_5), int32(0))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L4
	} else {
		goto L175
	}
L175:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1362), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L4
	} else {
		goto L176
	}
L176:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L177:
	;
	F_errmsg_internal(m, int32(_a_F_ExecIndexBuildScanKeys_6), int32(0))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1371), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+28)) = v360
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v724
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v40)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = v726
	F_errmsg_internal(m, int32(_a_F_ExecIndexBuildScanKeys_7), v40+int32(16))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1379), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L183:
	;
	F_errmsg_internal(m, int32(_a_F_ExecIndexBuildScanKeys_3), int32(0))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1476), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L4
	} else {
		goto L185
	}
L185:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L186:
	;
	F_errmsg_internal(m, int32(_a_F_ExecIndexBuildScanKeys_4), int32(0))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L4
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1480), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L4
	} else {
		goto L188
	}
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L189:
	;
	F_errmsg_internal(m, int32(_a_F_ExecIndexBuildScanKeys_8), int32(0))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L4
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1590), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L4
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
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v781
	F_errmsg_internal(m, int32(_a_F_ExecIndexBuildScanKeys_9), v40)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L4
	} else {
		goto L193
	}
L193:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1623), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L4
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
	v843 = v809
	v847 = v810
	goto L13
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v47
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v44
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v847
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v843
	v871 = int32(0)
	if l8 != 0 {
		v902 = v871
		v903 = v871
		goto L9
	} else {
		goto L197
	}
L197:
	;
	goto L8
L198:
	;
	v902 = v814
	v903 = v53
	goto L9
L199:
	;
	F_errmsg_internal(m, int32(_a_F_ExecIndexBuildScanKeys_10), int32(0))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1648), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L4
	} else {
		goto L201
	}
L201:
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
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 float64
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v274 int32
	_ = v274
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
	var v285 int32
	_ = v285
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
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
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
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
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
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
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
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
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
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
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v531 int32
	_ = v531
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
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[0]))
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
	v156 = F_index_getnext_slot(m, v44, int32(1), v18)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
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
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+184)))
	if v135 != int32(1) {
		v152 = v55
		goto L17
	} else {
		goto L37
	}
L21:
	;
	v130 = F_reorderqueue_pop(m, l0)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
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
	if int32(0) < v114 {
		v152 = v64
		goto L17
	} else {
		goto L34
	}
L26:
	;
	goto L25
L27:
	;
	v114 = v88 ^ int32(1)
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
	v98 = v75 << (uint(int32(2)) % 32)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v73+v98)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+v71)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v106 = v103 + v75*int32(36)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	v108 = m.T0[v107].(func(*base.Module, int32, int32, int32) int32)(m, v100, v102, v106)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	if v108 != 0 {
		v114 = v108
		goto L26
	} else {
		goto L32
	}
L32:
	;
	v111 = v75 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v111 < v112 {
		v75 = v111
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
	F_ExecForceStoreHeapTuple(m, v130, v18, int32(1))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	goto L16
L37:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	m.T0[v139].(func(*base.Module, int32))(m, v18)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	goto L16
L39:
	;
	v531 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+184)) = uint8(v531)
	v55 = v152
	goto L10
L40:
	;
	if v156 == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	goto L42
L42:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+72)))
	if v173 != int32(1) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+84)))
	if v212 == int32(1) {
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
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v177 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	F_MemoryContextReset(m, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L3
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v183 = int32(_a_F_IndexNextWithReorder_0)
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[1]))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[1])) = v186
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
	v191 = m.T0[v190].(func(*base.Module, int32, int32, int32) int32)(m, v177, v19, v16+int32(15))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L3
	} else {
		goto L50
	}
L49:
	;
	goto L44
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[1])) = v184
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	F_MemoryContextReset(m, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	if v191 != 0 {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v198 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v199 = *(*float64)(unsafe.Add(mBase, uint32(v198)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v198)+248)) = base.F64_add(v199, float64(1))
	goto L55
L54:
	;
	goto L55
L55:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[0]))
	if v204 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L3
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v208 = F_index_getnext_slot(m, v44, int32(1), v18)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L3
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	if v208 != 0 {
		goto L42
	} else {
		goto L61
	}
L61:
	;
	goto L39
L62:
	;
	v426 = int32(_a_F_IndexNextWithReorder_0)
	v427 = *(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[1]))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[1])) = v431
	v434 = F_palloc(m, int32(24))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L3
	} else {
		goto L106
	}
L63:
	;
	if v152 == int32(0) {
		goto L16
	} else {
		goto L93
	}
L64:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v298 == int32(0) {
		v417 = v345
		v418 = v344
		goto L62
	} else {
		goto L92
	}
L65:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v353 = v343
	v354 = v342
	goto L63
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v18
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	F_MemoryContextReset(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L3
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	v353 = v341
	v354 = v340
	goto L63
L69:
	;
	v219 = int32(_a_F_IndexNextWithReorder_0)
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[1]))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[1])) = v222
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v224 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v274 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[1])) = v220
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v279 <= v274 {
		v353 = v277
		v354 = v278
		goto L63
	} else {
		goto L77
	}
L71:
	;
	v227 = int32(0)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	if v228 <= v227 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v232 = v227
	goto L73
L73:
	;
	v245 = v232 << (uint(int32(2)) % 32)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v224)+12))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v245+v246)))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v248)+20))
	v252 = m.T0[v251].(func(*base.Module, int32, int32, int32) int32)(m, v248, v19, v249+v232)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L3
	} else {
		goto L75
	}
L74:
	;
	goto L70
L75:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v254+v245))) = v252
	v258 = v232 + int32(1)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	if v258 < v259 {
		v232 = v258
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	v285 = v274
	goto L79
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L3
	} else {
		goto L89
	}
L79:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285+v282))))
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285+v277))))
	if v300 != 0 {
		goto L64
	} else {
		goto L81
	}
L80:
	;
	if v314 < int32(0) {
		goto L78
	} else {
		goto L88
	}
L81:
	;
	if v298&int32(1) != 0 {
		goto L78
	} else {
		goto L82
	}
L82:
	;
	v304 = v285 << (uint(int32(2)) % 32)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v278+v304)))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v304+v283)))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v312 = v309 + v285*int32(36)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+16))
	v314 = m.T0[v313].(func(*base.Module, int32, int32, int32) int32)(m, v306, v308, v312)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L3
	} else {
		goto L83
	}
L83:
	;
	if v314 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v319 = v285 + int32(1)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v320 <= v319 {
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
	v285 = v319
	goto L79
L88:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v417 = v325
	v418 = v324
	goto L62
L89:
	;
	F_errmsg_internal(m, int32(_a_F_IndexNextWithReorder_1), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L3
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_IndexNextWithReorder_2), int32(311), int32(_a_F_IndexNextWithReorder_3))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
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
	v353 = v345
	v354 = v344
	goto L63
L93:
	;
	v363 = int32(0)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v364 <= v363 {
		goto L16
	} else {
		goto L94
	}
L94:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v152)+20))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v152)+16))
	v370 = v363
	goto L95
L95:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370+v367))))
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370+v353))))
	if v385 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	if v409 <= int32(0) {
		goto L16
	} else {
		goto L105
	}
L97:
	;
	goto L96
L98:
	;
	v409 = v383 ^ int32(1)
	goto L97
L99:
	;
	goto L100
L100:
	;
	if v383&int32(1) != 0 {
		goto L16
	} else {
		goto L101
	}
L101:
	;
	v393 = v370 << (uint(int32(2)) % 32)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v354+v393)))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v393+v368)))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v401 = v398 + v370*int32(36)
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+16))
	v403 = m.T0[v402].(func(*base.Module, int32, int32, int32) int32)(m, v395, v397, v401)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L3
	} else {
		goto L102
	}
L102:
	;
	if v403 != 0 {
		v409 = v403
		goto L97
	} else {
		goto L103
	}
L103:
	;
	v406 = v370 + int32(1)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v406 < v407 {
		v370 = v406
		goto L95
	} else {
		goto L104
	}
L104:
	;
	goto L16
L105:
	;
	v417 = v353
	v418 = v354
	goto L62
L106:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+44))
	v438 = m.T0[v437].(func(*base.Module, int32) int32)(m, v18)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L3
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v434)+12)) = v438
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v428)+16))
	v444 = F_palloc(m, v441<<(uint(int32(2))%32))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L3
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v434)+16)) = v444
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v428)+16))
	v448 = F_palloc(m, v447)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L3
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v434)+20)) = v448
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if int32(0) < v451 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v455 = int32(0)
	goto L113
L111:
	;
	goto L112
L112:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	F_pairingheap_add(m, v513, v434)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L3
	} else {
		goto L120
	}
L113:
	;
	v467 = int32(0)
	v468 = v455 + v417
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	if v469 == v467 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	goto L112
L115:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v418+v455<<(uint(int32(2))%32))))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476+v455))))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v483 = int32(*(*int16)(unsafe.Add(mBase, uint32(v479+v455<<(uint(int32(1))%32)))))
	v484 = F_datumCopy(m, v475, v478, v483)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L3
	} else {
		goto L118
	}
L116:
	;
	v486 = v467
	goto L117
L117:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v434)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v487+v455<<(uint(int32(2))%32)))) = v486
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v434)+20))
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	*(*uint8)(unsafe.Add(mBase, uint32(v492+v455))) = uint8(v494)
	v497 = v455 + int32(1)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v497 < v498 {
		v455 = v497
		goto L113
	} else {
		goto L119
	}
L118:
	;
	v486 = v484
	goto L117
L119:
	;
	goto L114
L120:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[1])) = v427
	v55 = v152
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
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
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
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
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
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
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
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
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
		v19 = v11 - int32(-64)
		F_ScanKeyInit(m, v19, int32(1), int32(3), int32(184), v13)
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
				v37 = F_systable_beginscan(m, v16, int32(2680), int32(1), int32(0), int32(2), v19)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v39 = F_systable_getnext(m, v37)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						if v39 == int32(0) {
							if l1 == int32(0) {
								v62 = v3
								F_systable_endscan(m, v37)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									F_relation_close(m, v16, int32(3))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										if l1 != 0 {
											F_LockRelationOid(m, l1, int32(4))
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return
											} else {
												F_SetRelationHasSubclass(m, l1, int32(1))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return
												} else {
													v76 = F_table_open(m, int32(1259), int32(3))
													mBase = m.M
													v77 = m.ExcPending
													if v77 != 0 {
														return
													} else {
														v79 = F_SearchSysCacheLockedCopy1(m, int32(57), v13)
														mBase = m.M
														v80 = m.ExcPending
														if v80 != 0 {
															return
														} else {
															if v79 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v168 = m.ExcPending
																if v168 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
																	F_errmsg_internal(m, int32(_a_F_IndexSetParentIndex_0), v11)
																	mBase = m.M
																	v172 = m.ExcPending
																	if v172 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_IndexSetParentIndex_1), int32(_a_F_IndexSetParentIndex_2), int32(_a_F_IndexSetParentIndex_3))
																		mBase = m.M
																		v177 = m.ExcPending
																		if v177 != 0 {
																			return
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	}
																}
															} else {
																v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+8)))
																*(*uint16)(unsafe.Add(mBase, uint32(v11)+56)) = uint16(v83)
																v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v85
																v87 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
																v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+22)))
																*(*uint8)(unsafe.Add(mBase, uint32(v87+v88)+131)) = uint8(base.B2i32(l1 != int32(0)))
																v94 = v11 + int32(52)
																F_CatalogTupleUpdate(m, v76, v94, v79)
																mBase = m.M
																v96 = m.ExcPending
																if v96 != 0 {
																	return
																} else {
																	F_UnlockTuple(m, v76, v94, int32(7))
																	mBase = m.M
																	v99 = m.ExcPending
																	if v99 != 0 {
																		return
																	} else {
																		F_pfree(m, v79)
																		mBase = m.M
																		v101 = m.ExcPending
																		if v101 != 0 {
																			return
																		} else {
																			F_relation_close(m, v76, int32(3))
																			mBase = m.M
																			v104 = m.ExcPending
																			if v104 != 0 {
																				return
																			} else {
																				if v62 != 0 {
																					if l1 != 0 {
																						v105 = int32(0)
																						*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v105
																						*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v13
																						v108 = int32(1259)
																						*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v108
																						*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v105
																						*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = l1
																						*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v108
																						*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v108
																						v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																						v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
																						*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v105
																						*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v118
																						F_recordDependencyOn(m, v94, v11+int32(40), int32(80))
																						mBase = m.M
																						v126 = m.ExcPending
																						if v126 != 0 {
																							return
																						} else {
																							F_recordDependencyOn(m, v94, v11+int32(28), int32(83))
																							mBase = m.M
																							v131 = m.ExcPending
																							if v131 != 0 {
																								return
																							} else {
																								F_CommandCounterIncrement(m)
																								mBase = m.M
																								v144 = m.ExcPending
																								if v144 != 0 {
																									return
																								} else {
																									m.G0 = v11 + int32(160)
																									return
																								}
																							}
																						}
																					} else {
																						v132 = int32(1259)
																						v135 = F_deleteDependencyRecordsForClass(m, v132, v13, v132, int32(80))
																						mBase = m.M
																						v136 = m.ExcPending
																						if v136 != 0 {
																							return
																						} else {
																							v137 = int32(1259)
																							v140 = F_deleteDependencyRecordsForClass(m, v137, v13, v137, int32(83))
																							mBase = m.M
																							v141 = m.ExcPending
																							if v141 != 0 {
																								return
																							} else {
																								F_CommandCounterIncrement(m)
																								mBase = m.M
																								v144 = m.ExcPending
																								if v144 != 0 {
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
											v76 = F_table_open(m, int32(1259), int32(3))
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return
											} else {
												v79 = F_SearchSysCacheLockedCopy1(m, int32(57), v13)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return
												} else {
													if v79 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
															F_errmsg_internal(m, int32(_a_F_IndexSetParentIndex_0), v11)
															mBase = m.M
															v172 = m.ExcPending
															if v172 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_IndexSetParentIndex_1), int32(_a_F_IndexSetParentIndex_2), int32(_a_F_IndexSetParentIndex_3))
																mBase = m.M
																v177 = m.ExcPending
																if v177 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+8)))
														*(*uint16)(unsafe.Add(mBase, uint32(v11)+56)) = uint16(v83)
														v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v85
														v87 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
														v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+22)))
														*(*uint8)(unsafe.Add(mBase, uint32(v87+v88)+131)) = uint8(base.B2i32(l1 != int32(0)))
														v94 = v11 + int32(52)
														F_CatalogTupleUpdate(m, v76, v94, v79)
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return
														} else {
															F_UnlockTuple(m, v76, v94, int32(7))
															mBase = m.M
															v99 = m.ExcPending
															if v99 != 0 {
																return
															} else {
																F_pfree(m, v79)
																mBase = m.M
																v101 = m.ExcPending
																if v101 != 0 {
																	return
																} else {
																	F_relation_close(m, v76, int32(3))
																	mBase = m.M
																	v104 = m.ExcPending
																	if v104 != 0 {
																		return
																	} else {
																		if v62 != 0 {
																			if l1 != 0 {
																				v105 = int32(0)
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v105
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v13
																				v108 = int32(1259)
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v108
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v105
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = l1
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v108
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v108
																				v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																				v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v105
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v118
																				F_recordDependencyOn(m, v94, v11+int32(40), int32(80))
																				mBase = m.M
																				v126 = m.ExcPending
																				if v126 != 0 {
																					return
																				} else {
																					F_recordDependencyOn(m, v94, v11+int32(28), int32(83))
																					mBase = m.M
																					v131 = m.ExcPending
																					if v131 != 0 {
																						return
																					} else {
																						F_CommandCounterIncrement(m)
																						mBase = m.M
																						v144 = m.ExcPending
																						if v144 != 0 {
																							return
																						} else {
																							m.G0 = v11 + int32(160)
																							return
																						}
																					}
																				}
																			} else {
																				v132 = int32(1259)
																				v135 = F_deleteDependencyRecordsForClass(m, v132, v13, v132, int32(80))
																				mBase = m.M
																				v136 = m.ExcPending
																				if v136 != 0 {
																					return
																				} else {
																					v137 = int32(1259)
																					v140 = F_deleteDependencyRecordsForClass(m, v137, v13, v137, int32(83))
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return
																					} else {
																						F_CommandCounterIncrement(m)
																						mBase = m.M
																						v144 = m.ExcPending
																						if v144 != 0 {
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
								v45 = int32(1)
								F_StoreSingleInheritance(m, v13, l1, v45)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									v62 = v45
									F_systable_endscan(m, v37)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										F_relation_close(m, v16, int32(3))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return
										} else {
											if l1 != 0 {
												F_LockRelationOid(m, l1, int32(4))
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return
												} else {
													F_SetRelationHasSubclass(m, l1, int32(1))
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return
													} else {
														v76 = F_table_open(m, int32(1259), int32(3))
														mBase = m.M
														v77 = m.ExcPending
														if v77 != 0 {
															return
														} else {
															v79 = F_SearchSysCacheLockedCopy1(m, int32(57), v13)
															mBase = m.M
															v80 = m.ExcPending
															if v80 != 0 {
																return
															} else {
																if v79 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v168 = m.ExcPending
																	if v168 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
																		F_errmsg_internal(m, int32(_a_F_IndexSetParentIndex_0), v11)
																		mBase = m.M
																		v172 = m.ExcPending
																		if v172 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_IndexSetParentIndex_1), int32(_a_F_IndexSetParentIndex_2), int32(_a_F_IndexSetParentIndex_3))
																			mBase = m.M
																			v177 = m.ExcPending
																			if v177 != 0 {
																				return
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		}
																	}
																} else {
																	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+8)))
																	*(*uint16)(unsafe.Add(mBase, uint32(v11)+56)) = uint16(v83)
																	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v85
																	v87 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
																	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+22)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v87+v88)+131)) = uint8(base.B2i32(l1 != int32(0)))
																	v94 = v11 + int32(52)
																	F_CatalogTupleUpdate(m, v76, v94, v79)
																	mBase = m.M
																	v96 = m.ExcPending
																	if v96 != 0 {
																		return
																	} else {
																		F_UnlockTuple(m, v76, v94, int32(7))
																		mBase = m.M
																		v99 = m.ExcPending
																		if v99 != 0 {
																			return
																		} else {
																			F_pfree(m, v79)
																			mBase = m.M
																			v101 = m.ExcPending
																			if v101 != 0 {
																				return
																			} else {
																				F_relation_close(m, v76, int32(3))
																				mBase = m.M
																				v104 = m.ExcPending
																				if v104 != 0 {
																					return
																				} else {
																					if v62 != 0 {
																						if l1 != 0 {
																							v105 = int32(0)
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v105
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v13
																							v108 = int32(1259)
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v108
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v105
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = l1
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v108
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v108
																							v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																							v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v105
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v118
																							F_recordDependencyOn(m, v94, v11+int32(40), int32(80))
																							mBase = m.M
																							v126 = m.ExcPending
																							if v126 != 0 {
																								return
																							} else {
																								F_recordDependencyOn(m, v94, v11+int32(28), int32(83))
																								mBase = m.M
																								v131 = m.ExcPending
																								if v131 != 0 {
																									return
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v144 = m.ExcPending
																									if v144 != 0 {
																										return
																									} else {
																										m.G0 = v11 + int32(160)
																										return
																									}
																								}
																							}
																						} else {
																							v132 = int32(1259)
																							v135 = F_deleteDependencyRecordsForClass(m, v132, v13, v132, int32(80))
																							mBase = m.M
																							v136 = m.ExcPending
																							if v136 != 0 {
																								return
																							} else {
																								v137 = int32(1259)
																								v140 = F_deleteDependencyRecordsForClass(m, v137, v13, v137, int32(83))
																								mBase = m.M
																								v141 = m.ExcPending
																								if v141 != 0 {
																									return
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v144 = m.ExcPending
																									if v144 != 0 {
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
												v76 = F_table_open(m, int32(1259), int32(3))
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return
												} else {
													v79 = F_SearchSysCacheLockedCopy1(m, int32(57), v13)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return
													} else {
														if v79 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
																F_errmsg_internal(m, int32(_a_F_IndexSetParentIndex_0), v11)
																mBase = m.M
																v172 = m.ExcPending
																if v172 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_IndexSetParentIndex_1), int32(_a_F_IndexSetParentIndex_2), int32(_a_F_IndexSetParentIndex_3))
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+8)))
															*(*uint16)(unsafe.Add(mBase, uint32(v11)+56)) = uint16(v83)
															v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v85
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
															v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+22)))
															*(*uint8)(unsafe.Add(mBase, uint32(v87+v88)+131)) = uint8(base.B2i32(l1 != int32(0)))
															v94 = v11 + int32(52)
															F_CatalogTupleUpdate(m, v76, v94, v79)
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
																return
															} else {
																F_UnlockTuple(m, v76, v94, int32(7))
																mBase = m.M
																v99 = m.ExcPending
																if v99 != 0 {
																	return
																} else {
																	F_pfree(m, v79)
																	mBase = m.M
																	v101 = m.ExcPending
																	if v101 != 0 {
																		return
																	} else {
																		F_relation_close(m, v76, int32(3))
																		mBase = m.M
																		v104 = m.ExcPending
																		if v104 != 0 {
																			return
																		} else {
																			if v62 != 0 {
																				if l1 != 0 {
																					v105 = int32(0)
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v105
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v13
																					v108 = int32(1259)
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v108
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v105
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = l1
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v108
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v108
																					v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																					v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v105
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v118
																					F_recordDependencyOn(m, v94, v11+int32(40), int32(80))
																					mBase = m.M
																					v126 = m.ExcPending
																					if v126 != 0 {
																						return
																					} else {
																						F_recordDependencyOn(m, v94, v11+int32(28), int32(83))
																						mBase = m.M
																						v131 = m.ExcPending
																						if v131 != 0 {
																							return
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v144 = m.ExcPending
																							if v144 != 0 {
																								return
																							} else {
																								m.G0 = v11 + int32(160)
																								return
																							}
																						}
																					}
																				} else {
																					v132 = int32(1259)
																					v135 = F_deleteDependencyRecordsForClass(m, v132, v13, v132, int32(80))
																					mBase = m.M
																					v136 = m.ExcPending
																					if v136 != 0 {
																						return
																					} else {
																						v137 = int32(1259)
																						v140 = F_deleteDependencyRecordsForClass(m, v137, v13, v137, int32(83))
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v144 = m.ExcPending
																							if v144 != 0 {
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
								F_simple_heap_delete(m, v16, v39+int32(4))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									v62 = int32(1)
									F_systable_endscan(m, v37)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										F_relation_close(m, v16, int32(3))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return
										} else {
											if l1 != 0 {
												F_LockRelationOid(m, l1, int32(4))
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return
												} else {
													F_SetRelationHasSubclass(m, l1, int32(1))
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return
													} else {
														v76 = F_table_open(m, int32(1259), int32(3))
														mBase = m.M
														v77 = m.ExcPending
														if v77 != 0 {
															return
														} else {
															v79 = F_SearchSysCacheLockedCopy1(m, int32(57), v13)
															mBase = m.M
															v80 = m.ExcPending
															if v80 != 0 {
																return
															} else {
																if v79 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v168 = m.ExcPending
																	if v168 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
																		F_errmsg_internal(m, int32(_a_F_IndexSetParentIndex_0), v11)
																		mBase = m.M
																		v172 = m.ExcPending
																		if v172 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_IndexSetParentIndex_1), int32(_a_F_IndexSetParentIndex_2), int32(_a_F_IndexSetParentIndex_3))
																			mBase = m.M
																			v177 = m.ExcPending
																			if v177 != 0 {
																				return
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		}
																	}
																} else {
																	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+8)))
																	*(*uint16)(unsafe.Add(mBase, uint32(v11)+56)) = uint16(v83)
																	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v85
																	v87 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
																	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+22)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v87+v88)+131)) = uint8(base.B2i32(l1 != int32(0)))
																	v94 = v11 + int32(52)
																	F_CatalogTupleUpdate(m, v76, v94, v79)
																	mBase = m.M
																	v96 = m.ExcPending
																	if v96 != 0 {
																		return
																	} else {
																		F_UnlockTuple(m, v76, v94, int32(7))
																		mBase = m.M
																		v99 = m.ExcPending
																		if v99 != 0 {
																			return
																		} else {
																			F_pfree(m, v79)
																			mBase = m.M
																			v101 = m.ExcPending
																			if v101 != 0 {
																				return
																			} else {
																				F_relation_close(m, v76, int32(3))
																				mBase = m.M
																				v104 = m.ExcPending
																				if v104 != 0 {
																					return
																				} else {
																					if v62 != 0 {
																						if l1 != 0 {
																							v105 = int32(0)
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v105
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v13
																							v108 = int32(1259)
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v108
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v105
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = l1
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v108
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v108
																							v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																							v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v105
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v118
																							F_recordDependencyOn(m, v94, v11+int32(40), int32(80))
																							mBase = m.M
																							v126 = m.ExcPending
																							if v126 != 0 {
																								return
																							} else {
																								F_recordDependencyOn(m, v94, v11+int32(28), int32(83))
																								mBase = m.M
																								v131 = m.ExcPending
																								if v131 != 0 {
																									return
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v144 = m.ExcPending
																									if v144 != 0 {
																										return
																									} else {
																										m.G0 = v11 + int32(160)
																										return
																									}
																								}
																							}
																						} else {
																							v132 = int32(1259)
																							v135 = F_deleteDependencyRecordsForClass(m, v132, v13, v132, int32(80))
																							mBase = m.M
																							v136 = m.ExcPending
																							if v136 != 0 {
																								return
																							} else {
																								v137 = int32(1259)
																								v140 = F_deleteDependencyRecordsForClass(m, v137, v13, v137, int32(83))
																								mBase = m.M
																								v141 = m.ExcPending
																								if v141 != 0 {
																									return
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v144 = m.ExcPending
																									if v144 != 0 {
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
												v76 = F_table_open(m, int32(1259), int32(3))
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return
												} else {
													v79 = F_SearchSysCacheLockedCopy1(m, int32(57), v13)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return
													} else {
														if v79 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
																F_errmsg_internal(m, int32(_a_F_IndexSetParentIndex_0), v11)
																mBase = m.M
																v172 = m.ExcPending
																if v172 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_IndexSetParentIndex_1), int32(_a_F_IndexSetParentIndex_2), int32(_a_F_IndexSetParentIndex_3))
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+8)))
															*(*uint16)(unsafe.Add(mBase, uint32(v11)+56)) = uint16(v83)
															v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v85
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
															v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+22)))
															*(*uint8)(unsafe.Add(mBase, uint32(v87+v88)+131)) = uint8(base.B2i32(l1 != int32(0)))
															v94 = v11 + int32(52)
															F_CatalogTupleUpdate(m, v76, v94, v79)
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
																return
															} else {
																F_UnlockTuple(m, v76, v94, int32(7))
																mBase = m.M
																v99 = m.ExcPending
																if v99 != 0 {
																	return
																} else {
																	F_pfree(m, v79)
																	mBase = m.M
																	v101 = m.ExcPending
																	if v101 != 0 {
																		return
																	} else {
																		F_relation_close(m, v76, int32(3))
																		mBase = m.M
																		v104 = m.ExcPending
																		if v104 != 0 {
																			return
																		} else {
																			if v62 != 0 {
																				if l1 != 0 {
																					v105 = int32(0)
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v105
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v13
																					v108 = int32(1259)
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v108
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v105
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = l1
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v108
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v108
																					v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																					v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v105
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v118
																					F_recordDependencyOn(m, v94, v11+int32(40), int32(80))
																					mBase = m.M
																					v126 = m.ExcPending
																					if v126 != 0 {
																						return
																					} else {
																						F_recordDependencyOn(m, v94, v11+int32(28), int32(83))
																						mBase = m.M
																						v131 = m.ExcPending
																						if v131 != 0 {
																							return
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v144 = m.ExcPending
																							if v144 != 0 {
																								return
																							} else {
																								m.G0 = v11 + int32(160)
																								return
																							}
																						}
																					}
																				} else {
																					v132 = int32(1259)
																					v135 = F_deleteDependencyRecordsForClass(m, v132, v13, v132, int32(80))
																					mBase = m.M
																					v136 = m.ExcPending
																					if v136 != 0 {
																						return
																					} else {
																						v137 = int32(1259)
																						v140 = F_deleteDependencyRecordsForClass(m, v137, v13, v137, int32(83))
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v144 = m.ExcPending
																							if v144 != 0 {
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
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
								v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+22)))
								v58 = v56 + v57
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								if v59 != l1 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return
									} else {
										v153 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
										*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v153
										F_errmsg_internal(m, int32(_a_F_IndexSetParentIndex_4), v11+int32(16))
										mBase = m.M
										v159 = m.ExcPending
										if v159 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_IndexSetParentIndex_1), int32(_a_F_IndexSetParentIndex_5), int32(_a_F_IndexSetParentIndex_6))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v62 = v3
									F_systable_endscan(m, v37)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										F_relation_close(m, v16, int32(3))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return
										} else {
											if l1 != 0 {
												F_LockRelationOid(m, l1, int32(4))
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return
												} else {
													F_SetRelationHasSubclass(m, l1, int32(1))
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return
													} else {
														v76 = F_table_open(m, int32(1259), int32(3))
														mBase = m.M
														v77 = m.ExcPending
														if v77 != 0 {
															return
														} else {
															v79 = F_SearchSysCacheLockedCopy1(m, int32(57), v13)
															mBase = m.M
															v80 = m.ExcPending
															if v80 != 0 {
																return
															} else {
																if v79 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v168 = m.ExcPending
																	if v168 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
																		F_errmsg_internal(m, int32(_a_F_IndexSetParentIndex_0), v11)
																		mBase = m.M
																		v172 = m.ExcPending
																		if v172 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_IndexSetParentIndex_1), int32(_a_F_IndexSetParentIndex_2), int32(_a_F_IndexSetParentIndex_3))
																			mBase = m.M
																			v177 = m.ExcPending
																			if v177 != 0 {
																				return
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		}
																	}
																} else {
																	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+8)))
																	*(*uint16)(unsafe.Add(mBase, uint32(v11)+56)) = uint16(v83)
																	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v85
																	v87 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
																	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+22)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v87+v88)+131)) = uint8(base.B2i32(l1 != int32(0)))
																	v94 = v11 + int32(52)
																	F_CatalogTupleUpdate(m, v76, v94, v79)
																	mBase = m.M
																	v96 = m.ExcPending
																	if v96 != 0 {
																		return
																	} else {
																		F_UnlockTuple(m, v76, v94, int32(7))
																		mBase = m.M
																		v99 = m.ExcPending
																		if v99 != 0 {
																			return
																		} else {
																			F_pfree(m, v79)
																			mBase = m.M
																			v101 = m.ExcPending
																			if v101 != 0 {
																				return
																			} else {
																				F_relation_close(m, v76, int32(3))
																				mBase = m.M
																				v104 = m.ExcPending
																				if v104 != 0 {
																					return
																				} else {
																					if v62 != 0 {
																						if l1 != 0 {
																							v105 = int32(0)
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v105
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v13
																							v108 = int32(1259)
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v108
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v105
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = l1
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v108
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v108
																							v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																							v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v105
																							*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v118
																							F_recordDependencyOn(m, v94, v11+int32(40), int32(80))
																							mBase = m.M
																							v126 = m.ExcPending
																							if v126 != 0 {
																								return
																							} else {
																								F_recordDependencyOn(m, v94, v11+int32(28), int32(83))
																								mBase = m.M
																								v131 = m.ExcPending
																								if v131 != 0 {
																									return
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v144 = m.ExcPending
																									if v144 != 0 {
																										return
																									} else {
																										m.G0 = v11 + int32(160)
																										return
																									}
																								}
																							}
																						} else {
																							v132 = int32(1259)
																							v135 = F_deleteDependencyRecordsForClass(m, v132, v13, v132, int32(80))
																							mBase = m.M
																							v136 = m.ExcPending
																							if v136 != 0 {
																								return
																							} else {
																								v137 = int32(1259)
																								v140 = F_deleteDependencyRecordsForClass(m, v137, v13, v137, int32(83))
																								mBase = m.M
																								v141 = m.ExcPending
																								if v141 != 0 {
																									return
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v144 = m.ExcPending
																									if v144 != 0 {
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
												v76 = F_table_open(m, int32(1259), int32(3))
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return
												} else {
													v79 = F_SearchSysCacheLockedCopy1(m, int32(57), v13)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return
													} else {
														if v79 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
																F_errmsg_internal(m, int32(_a_F_IndexSetParentIndex_0), v11)
																mBase = m.M
																v172 = m.ExcPending
																if v172 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_IndexSetParentIndex_1), int32(_a_F_IndexSetParentIndex_2), int32(_a_F_IndexSetParentIndex_3))
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+8)))
															*(*uint16)(unsafe.Add(mBase, uint32(v11)+56)) = uint16(v83)
															v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v85
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
															v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+22)))
															*(*uint8)(unsafe.Add(mBase, uint32(v87+v88)+131)) = uint8(base.B2i32(l1 != int32(0)))
															v94 = v11 + int32(52)
															F_CatalogTupleUpdate(m, v76, v94, v79)
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
																return
															} else {
																F_UnlockTuple(m, v76, v94, int32(7))
																mBase = m.M
																v99 = m.ExcPending
																if v99 != 0 {
																	return
																} else {
																	F_pfree(m, v79)
																	mBase = m.M
																	v101 = m.ExcPending
																	if v101 != 0 {
																		return
																	} else {
																		F_relation_close(m, v76, int32(3))
																		mBase = m.M
																		v104 = m.ExcPending
																		if v104 != 0 {
																			return
																		} else {
																			if v62 != 0 {
																				if l1 != 0 {
																					v105 = int32(0)
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v105
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v13
																					v108 = int32(1259)
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v108
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v105
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = l1
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v108
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v108
																					v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																					v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v105
																					*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v118
																					F_recordDependencyOn(m, v94, v11+int32(40), int32(80))
																					mBase = m.M
																					v126 = m.ExcPending
																					if v126 != 0 {
																						return
																					} else {
																						F_recordDependencyOn(m, v94, v11+int32(28), int32(83))
																						mBase = m.M
																						v131 = m.ExcPending
																						if v131 != 0 {
																							return
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v144 = m.ExcPending
																							if v144 != 0 {
																								return
																							} else {
																								m.G0 = v11 + int32(160)
																								return
																							}
																						}
																					}
																				} else {
																					v132 = int32(1259)
																					v135 = F_deleteDependencyRecordsForClass(m, v132, v13, v132, int32(80))
																					mBase = m.M
																					v136 = m.ExcPending
																					if v136 != 0 {
																						return
																					} else {
																						v137 = int32(1259)
																						v140 = F_deleteDependencyRecordsForClass(m, v137, v13, v137, int32(83))
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v144 = m.ExcPending
																							if v144 != 0 {
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
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
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 float64
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
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
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
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
	var v409 int32
	_ = v409
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
	var v506 int64
	_ = v506
	var v508 int64
	_ = v508
	var v510 int64
	_ = v510
	var v512 int64
	_ = v512
	var v514 int32
	_ = v514
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v630 int32
	_ = v630
	var v643 int32
	_ = v643
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v682 int32
	_ = v682
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v736 int32
	_ = v736
	var v744 int32
	_ = v744
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v769 int32
	_ = v769
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v798 int32
	_ = v798
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v919 int32
	_ = v919
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v956 int32
	_ = v956
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v981 int32
	_ = v981
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1013 int32
	_ = v1013
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
	var v1035 int32
	_ = v1035
	var v1040 int32
	_ = v1040
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1083 int32
	_ = v1083
	var v1088 int32
	_ = v1088
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
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
	return v1105
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
	v1105 = int32(0)
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
	v57 = v8
	v58 = v8
	v59 = v36
	goto L10
L8:
	;
	v174 = v8
	v175 = v36
	goto L9
L9:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v186 = F_bms_del_member(m, v175, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L5
	} else {
		goto L29
	}
L10:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l3+int32(4)+v57<<(uint(int32(2))%32))))
	if v72 == int32(0) {
		v142 = v58
		v143 = v59
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v174 = v142
	v175 = v143
	goto L9
L12:
	;
	if v142 != 0 {
		goto L25
	} else {
		goto L26
	}
L13:
	;
	v75 = int32(0)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v76 <= v75 {
		v142 = v58
		v143 = v59
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v82 = v75
	v92 = v58
	v93 = v59
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
	v142 = v123
	v143 = v124
	goto L12
L17:
	;
	v126 = v82 + int32(1)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v126 < v127 {
		v82 = v126
		v92 = v123
		v93 = v124
		goto L15
	} else {
		goto L24
	}
L18:
	;
	v118 = F_lappend(m, v92, v107)
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
	v123 = v92
	v124 = v93
	goto L17
L22:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v108)+28))
	v121 = F_bms_add_members(m, v93, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v123 = v118
	v124 = v121
	goto L17
L24:
	;
	goto L16
L25:
	;
	v158 = v57 + int32(1)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v158 < v159 {
		v57 = v158
		v58 = v142
		v59 = v143
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
	v1105 = int32(0)
	goto L1
L28:
	;
	goto L11
L29:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v189 = F_get_loop_count(m, l0, v188, v186)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v191 = int32(0)
	if l5 == int32(1) {
		v1000 = v191
		v1001 = v191
		v1005 = v8
		v1006 = v191
		v1013 = v8
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v1005|(l4|v174) == int32(0) {
		goto L200
	} else {
		goto L201
	}
L32:
	;
	v197 = int32(1)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	if v198 != 0 {
		v204 = v197
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v774 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v774
	v778 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_build_index_paths[0])))
	if v778 != int32(1) {
		v1000 = v208
		v1001 = v774
		v1005 = v761
		v1006 = v762
		v1013 = v769
		goto L31
	} else {
		goto L163
	}
L34:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l2)+60))
	v208 = v204 & base.B2i32(v205 != int32(0))
	if v208 != 0 {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	goto L34
L36:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+216)))
	if v199 != 0 {
		v204 = v197
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v200 != 0 {
		v204 = v197
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v204 = base.B2i32(v201 != int32(0))
	goto L35
L39:
	;
	v210 = F_build_index_pathkeys(m, l0, l2, int32(1))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L5
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+105)))
	if v214&v204&int32(1) == int32(0) {
		v736 = v8
		v744 = v8
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v212 = F_truncate_useless_pathkeys(m, l0, l1, v210)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v761 = v212
	v762 = v191
	v769 = v8
	goto L33
L44:
	;
	v761 = v736
	v762 = int32(0)
	v769 = v744
	goto L33
L45:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v220 == int32(0) {
		v761 = v8
		v762 = v191
		v769 = v8
		goto L33
	} else {
		goto L46
	}
L46:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	if int32(0) < v223 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v723 = F_list_copy_head(m, v710, v706)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L5
	} else {
		goto L162
	}
L48:
	;
	v695 = int32(0)
	if v674 == v695 {
		v761 = v682
		v762 = v695
		v769 = v690
		goto L33
	} else {
		goto L161
	}
L49:
	;
	v237 = v8
	v238 = v191
	v245 = v8
	goto L52
L50:
	;
	goto L51
L51:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v666 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L52:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250+v237<<(uint(int32(2))%32))))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+12))
	if v255 != int32(1) {
		v643 = v238
		v650 = v245
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v655 != 0 {
		goto L152
	} else {
		goto L153
	}
L54:
	;
	goto L53
L55:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+16)))
	if v258 != 0 {
		v643 = v238
		v650 = v245
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+41)))
	if v260 != 0 {
		v643 = v238
		v650 = v245
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+8))
	v264 = v27 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v264))) = v259
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v259)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+4)) = int32(-1)
	if v266 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v270 = v262
	goto L60
L59:
	;
	v270 = int32(0)
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+8)) = v270
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v259)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+16)) = v272
	if v272 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
	v276 = v274
	goto L63
L62:
	;
	v276 = int32(0)
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+12)) = v276
	v279 = v27 + int32(20)
	v280 = int32(0)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v279)+16))
	if v284 == v280 {
		v333 = v280
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v333 == int32(0) {
		v643 = v238
		v650 = v245
		goto L54
	} else {
		goto L79
	}
L65:
	;
	goto L64
L66:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v279)+12))
	if v287 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	v320 = v314 + int32(4)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v315)+4))
	if base.Ui32(v320) < base.Ui32(v316+v322<<(uint(int32(2))%32)) {
		goto L76
	} else {
		goto L77
	}
L68:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v284)+12))
	v314 = v287
	v315 = v284
	v316 = v288
	goto L67
L69:
	;
	goto L70
L70:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	v291 = v289
	goto L71
L71:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v279)+8))
	v296 = F_bms_next_member(m, v295, v291)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v279)+4)) = v296
	if v296 <= int32(0) {
		v333 = v280
		goto L65
	} else {
		goto L73
	}
L72:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v307)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v279)+12)) = v311
	v314 = v311
	v315 = v307
	v316 = v311
	goto L67
L73:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)+12))
	if v301 <= v296 {
		v333 = v280
		goto L65
	} else {
		goto L74
	}
L74:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v300)+20))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v303+v296<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v279)+16)) = v307
	if v307 == int32(0) {
		v291 = v296
		goto L71
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	v327 = v320
	goto L78
L77:
	;
	v327 = int32(0)
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279)+12)) = v327
	v333 = v318
	goto L65
L79:
	;
	v358 = v333
	goto L80
L80:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v358)+8))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)+8))
	v363 = int32(0)
	if base.B2i32(v360 == v363)|base.B2i32(v362 == v363) != 0 {
		v409 = base.B2i32(v360|v362 == v363)
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v643 = v238
	v650 = v245
	goto L54
L82:
	;
	v576 = v27 + int32(20)
	v577 = int32(0)
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v576)+16))
	if v581 == v577 {
		v630 = v577
		goto L136
	} else {
		goto L137
	}
L83:
	;
	if v409 == int32(0) {
		goto L82
	} else {
		goto L94
	}
L84:
	;
	goto L83
L85:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v360)+4))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v362)+4))
	if v377 != v378 {
		v409 = int32(0)
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v380 = int32(1)
	if v377 <= v380 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v383 = v380
	goto L89
L88:
	;
	v383 = v377
	goto L89
L89:
	;
	v384 = int32(8)
	v389 = int32(0)
	goto L90
L90:
	;
	v397 = v389 << (uint(int32(2)) % 32)
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v360+v384+v397)))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v362+v384+v397)))
	v402 = base.B2i32(v399 == v401)
	if v399 != v401 {
		v409 = v402
		goto L84
	} else {
		goto L92
	}
L91:
	;
	v409 = v402
	goto L84
L92:
	;
	v405 = v389 + int32(1)
	if v405 != v383 {
		v389 = v405
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v416 <= int32(0) {
		goto L82
	} else {
		goto L95
	}
L95:
	;
	v423 = int32(0)
	goto L96
L96:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	if v444 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L82
L98:
	;
	v548 = v423 + int32(1)
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v548 < v549 {
		v423 = v548
		goto L96
	} else {
		goto L134
	}
L99:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	if v447 != int32(17) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v444)+28))
	if v450 == int32(0) {
		goto L98
	} else {
		goto L101
	}
L101:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v450)+4))
	if v453 < int32(2) {
		goto L98
	} else {
		goto L102
	}
L102:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v450)+12))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v456)))
	if v457 == int32(0) {
		goto L98
	} else {
		goto L103
	}
L103:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v456)+4))
	if v460 == int32(0) {
		goto L98
	} else {
		goto L104
	}
L104:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v254)+8))
	v465 = v423 << (uint(int32(2)) % 32)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v465+v466)))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v444)+4))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v470+v465)))
	if v472 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v444)+24))
	if v472 != v473 {
		goto L98
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v475 = F_match_index_to_operand(m, v457, v423, l2)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L5
	} else {
		goto L111
	}
L108:
	;
	goto L107
L109:
	;
	v533 = F_lappend(m, v238, v530)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L5
	} else {
		goto L131
	}
L110:
	;
	v486 = F_match_index_to_operand(m, v460, v423, l2)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L5
	} else {
		goto L119
	}
L111:
	;
	if v475 == int32(0) {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v479 = F_contain_var_clause(m, v460)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L5
	} else {
		goto L113
	}
L113:
	;
	if v479 != 0 {
		goto L110
	} else {
		goto L114
	}
L114:
	;
	v481 = F_contain_volatile_functions(m, v460)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	if v481 != 0 {
		goto L110
	} else {
		goto L116
	}
L116:
	;
	v483 = F_get_op_opfamily_sortfamily(m, v469, v468)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	if v483 == v463 {
		v530 = v444
		goto L109
	} else {
		goto L118
	}
L118:
	;
	goto L98
L119:
	;
	if v486 == int32(0) {
		goto L98
	} else {
		goto L120
	}
L120:
	;
	v490 = F_contain_var_clause(m, v457)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L5
	} else {
		goto L121
	}
L121:
	;
	if v490 != 0 {
		goto L98
	} else {
		goto L122
	}
L122:
	;
	v492 = F_contain_volatile_functions(m, v457)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	if v492 != 0 {
		goto L98
	} else {
		goto L124
	}
L124:
	;
	v494 = F_get_commutator(m, v469)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L5
	} else {
		goto L125
	}
L125:
	;
	if v494 == int32(0) {
		goto L98
	} else {
		goto L126
	}
L126:
	;
	v498 = F_get_op_opfamily_sortfamily(m, v494, v468)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L5
	} else {
		goto L127
	}
L127:
	;
	if v498 != v463 {
		goto L98
	} else {
		goto L128
	}
L128:
	;
	v502 = F_palloc0(m, int32(36))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L5
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v502))) = int32(17)
	v506 = *(*int64)(unsafe.Add(mBase, uint32(v444)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v502)+8)) = v506
	v508 = *(*int64)(unsafe.Add(mBase, uint32(v444)))
	*(*int64)(unsafe.Add(mBase, uint32(v502))) = v508
	v510 = *(*int64)(unsafe.Add(mBase, uint32(v444)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v502)+16)) = v510
	v512 = *(*int64)(unsafe.Add(mBase, uint32(v444)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v502)+24)) = v512
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v444)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v502)+32)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v502)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v502)+4)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v460
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v460
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v457
	v527 = F_list_make2_impl(m, v27+int32(16), v27+int32(12))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L5
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v502)+28)) = v527
	v530 = v502
	goto L109
L131:
	;
	v535 = F_lappend_int(m, v245, v423)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L5
	} else {
		goto L132
	}
L132:
	;
	v538 = v237 + int32(1)
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	if v538 < v539 {
		v237 = v538
		v238 = v533
		v245 = v535
		goto L52
	} else {
		goto L133
	}
L133:
	;
	v643 = v533
	v650 = v535
	goto L54
L134:
	;
	goto L97
L135:
	;
	if v630 != 0 {
		v358 = v630
		goto L80
	} else {
		goto L150
	}
L136:
	;
	goto L135
L137:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v576)+12))
	if v584 != 0 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v611)))
	v617 = v611 + int32(4)
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v612)+4))
	if base.Ui32(v617) < base.Ui32(v613+v619<<(uint(int32(2))%32)) {
		goto L147
	} else {
		goto L148
	}
L139:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v581)+12))
	v611 = v584
	v612 = v581
	v613 = v585
	goto L138
L140:
	;
	goto L141
L141:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
	v588 = v586
	goto L142
L142:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v576)+8))
	v593 = F_bms_next_member(m, v592, v588)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v576)+4)) = v593
	if v593 <= int32(0) {
		v630 = v577
		goto L136
	} else {
		goto L144
	}
L143:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v604)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v576)+12)) = v608
	v611 = v608
	v612 = v604
	v613 = v608
	goto L138
L144:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v597)+12))
	if v598 <= v593 {
		v630 = v577
		goto L136
	} else {
		goto L145
	}
L145:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v597)+20))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v600+v593<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v576)+16)) = v604
	if v604 == int32(0) {
		v588 = v593
		goto L142
	} else {
		goto L146
	}
L146:
	;
	goto L143
L147:
	;
	v624 = v617
	goto L149
L148:
	;
	v624 = int32(0)
	goto L149
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v576)+12)) = v624
	v630 = v615
	goto L136
L150:
	;
	goto L81
L151:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v643)+4))
	if v661 == v663 {
		v761 = v662
		v762 = v643
		v769 = v650
		goto L33
	} else {
		goto L157
	}
L152:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v655)+4))
	if v643 != 0 {
		v661 = v656
		v662 = v655
		goto L151
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v657 = int32(0)
	if v643 == v657 {
		v736 = v657
		v744 = v650
		goto L44
	} else {
		goto L156
	}
L155:
	;
	v674 = v656
	v682 = v655
	v690 = v650
	goto L48
L156:
	;
	v661 = v657
	v662 = v657
	goto L151
L157:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v643)+4))
	v706 = v665
	v710 = v662
	v711 = v643
	v718 = v650
	goto L47
L158:
	;
	v761 = int32(0)
	v762 = v191
	v769 = v8
	goto L33
L159:
	;
	goto L160
L160:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v666)+4))
	v674 = v670
	v682 = v666
	v690 = v8
	goto L48
L161:
	;
	v706 = v695
	v710 = v682
	v711 = v695
	v718 = v690
	goto L47
L162:
	;
	v761 = v723
	v762 = v711
	v769 = v718
	goto L33
L163:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v781)+4))
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	F_pull_varattnos(m, v782, v783, v27+int32(20))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L5
	} else {
		goto L164
	}
L164:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(l2)+96))
	if v788 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	if v858 <= int32(0) {
		goto L173
	} else {
		goto L174
	}
L166:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v788)+4))
	if v791 <= int32(0) {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v798 = int32(0)
	goto L168
L168:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v788)+12))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v819+v798<<(uint(int32(2))%32))))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v823)+4))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	F_pull_varattnos(m, v824, v825, v27+int32(20))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L5
	} else {
		goto L170
	}
L169:
	;
	goto L165
L170:
	;
	v831 = v798 + int32(1)
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v788)+4))
	if v831 < v832 {
		v798 = v831
		goto L168
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v935 = int32(0)
	if v934 == v935 {
		goto L184
	} else {
		goto L185
	}
L173:
	;
	v919 = int32(0)
	goto L172
L174:
	;
	goto L175
L175:
	;
	v862 = int32(0)
	v867 = v862
	v871 = v858
	v873 = v862
	goto L176
L176:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v888+v867<<(uint(int32(2))%32))))
	if v892 == int32(0) {
		v905 = v871
		v906 = v873
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v919 = v906
	goto L172
L178:
	;
	v908 = v867 + int32(1)
	if v908 < v905 {
		v867 = v908
		v871 = v905
		v873 = v906
		goto L176
	} else {
		goto L182
	}
L179:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
	v897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v895+v867))))
	if v897 != int32(1) {
		v905 = v871
		v906 = v873
		goto L178
	} else {
		goto L180
	}
L180:
	;
	v902 = F_bms_add_member(m, v873, v892+int32(7))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L5
	} else {
		goto L181
	}
L181:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	v905 = v904
	v906 = v902
	goto L178
L182:
	;
	goto L177
L183:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	F_bms_free(m, v989)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L5
	} else {
		goto L197
	}
L184:
	;
	v988 = int32(1)
	goto L183
L185:
	;
	goto L186
L186:
	;
	if v919 == int32(0) {
		v981 = v935
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v988 = v981
	goto L183
L188:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v934)+4))
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v919)+4))
	if v945 < v944 {
		v981 = v935
		goto L187
	} else {
		goto L189
	}
L189:
	;
	v947 = int32(1)
	if v944 <= v947 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v950 = v947
	goto L192
L191:
	;
	v950 = v944
	goto L192
L192:
	;
	v951 = int32(8)
	v956 = int32(0)
	goto L193
L193:
	;
	v963 = v956 << (uint(int32(2)) % 32)
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v934+v951+v963)))
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v919+v951+v963)))
	v970 = v965 & (v967 ^ int32(-1))
	v972 = base.B2i32(v970 == int32(0))
	if v970 != 0 {
		v981 = v972
		goto L187
	} else {
		goto L195
	}
L194:
	;
	v981 = v972
	goto L187
L195:
	;
	v974 = v956 + int32(1)
	if v974 != v950 {
		v956 = v974
		goto L193
	} else {
		goto L196
	}
L196:
	;
	goto L194
L197:
	;
	F_bms_free(m, v919)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L5
	} else {
		goto L198
	}
L198:
	;
	v1000 = v208
	v1001 = v988
	v1005 = v761
	v1006 = v762
	v1013 = v769
	goto L31
L199:
	;
	if v1000 == int32(0) {
		v1105 = v1057
		goto L1
	} else {
		goto L214
	}
L200:
	;
	v1022 = int32(0)
	if v1001 == v1022 {
		v1057 = v1022
		goto L199
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	v1026 = int32(0)
	v1029 = F_create_index_path(m, l0, l2, v174, v1006, v1013, v1005, int32(1), v1001, v186, v189, v1026)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L5
	} else {
		goto L204
	}
L203:
	;
	goto L202
L204:
	;
	v1031 = F_lappend(m, v1026, v1029)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L5
	} else {
		goto L205
	}
L205:
	;
	v1033 = int32(1)
	v1035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+111)))
	if v186|(base.B2i32(l5 == v1033)|base.B2i32(v1035 != v1033)) != 0 {
		v1057 = v1031
		goto L199
	} else {
		goto L206
	}
L206:
	;
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v1040&int32(1) == int32(0) {
		v1057 = v1031
		goto L199
	} else {
		goto L207
	}
L207:
	;
	v1045 = int32(1)
	v1048 = F_create_index_path(m, l0, l2, v174, v1006, v1013, v1005, v1045, v1001, int32(0), v189, v1045)
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L5
	} else {
		goto L208
	}
L208:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+24))
	if int32(0) < v1050 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	F_add_partial_path(m, l1, v1048)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L5
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	F_pfree(m, v1048)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L5
	} else {
		goto L213
	}
L212:
	;
	v1057 = v1031
	goto L199
L213:
	;
	v1057 = v1031
	goto L199
L214:
	;
	v1062 = F_build_index_pathkeys(m, l0, l2, int32(-1))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L5
	} else {
		goto L215
	}
L215:
	;
	v1064 = F_truncate_useless_pathkeys(m, l0, l1, v1062)
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L5
	} else {
		goto L216
	}
L216:
	;
	if v1064 == int32(0) {
		v1105 = v1057
		goto L1
	} else {
		goto L217
	}
L217:
	;
	v1068 = int32(0)
	v1072 = F_create_index_path(m, l0, l2, v174, v1068, v1068, v1064, int32(-1), v1001, v186, v189, v1068)
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L5
	} else {
		goto L218
	}
L218:
	;
	v1074 = F_lappend(m, v1057, v1072)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L5
	} else {
		goto L219
	}
L219:
	;
	v1076 = int32(1)
	v1078 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+111)))
	if v186|(base.B2i32(l5 == v1076)|base.B2i32(v1078 != v1076)) != 0 {
		v1105 = v1074
		goto L1
	} else {
		goto L220
	}
L220:
	;
	v1083 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v1083&int32(1) == int32(0) {
		v1105 = v1074
		goto L1
	} else {
		goto L221
	}
L221:
	;
	v1088 = int32(0)
	v1093 = F_create_index_path(m, l0, l2, v174, v1088, v1088, v1064, int32(-1), v1001, v1088, v189, int32(1))
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L5
	} else {
		goto L222
	}
L222:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+24))
	if int32(0) < v1095 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	F_add_partial_path(m, l1, v1093)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L5
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	F_pfree(m, v1093)
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L5
	} else {
		goto L227
	}
L226:
	;
	v1105 = v1074
	goto L1
L227:
	;
	v1105 = v1074
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
		v111 = v7
		goto L1
	} else {
		goto L4
	}
L4:
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
		goto L5
	}
L5:
	;
	if v44 == int32(0) {
		v111 = v7
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v49 = F_palloc0(m, int32(20))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = int32(281)
	v53 = int32(0)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v53 < v54 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v60 = v53
	v63 = int32(0)
	goto L11
L9:
	;
	v90 = v53
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = l1
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+44)))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v49)+14)) = uint16(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+12)) = uint8(v99)
	v111 = v49
	goto L1
L11:
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
		goto L13
	}
L12:
	;
	v90 = v82
	goto L10
L13:
	;
	v82 = F_lappend(m, v60, v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v85 = v63 + int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v85 < v86 {
		v60 = v82
		v63 = v85
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
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
				F_errmsg_internal(m, int32(_a_F_get_index_isvalid_0), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_index_isvalid_1), int32(3726), int32(_a_F_get_index_isvalid_2))
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
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_index_concurrently_build[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v9+int32(12)))) = v19
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_index_concurrently_build[1]))
		*(*int32)(unsafe.Add(mBase, uint32(v9+int32(8)))) = v22
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
		*(*int32)(unsafe.Add(mBase, _c_F_index_concurrently_build[1])) = v26 | int32(2)
		*(*int32)(unsafe.Add(mBase, _c_F_index_concurrently_build[0])) = v25
		v34 = int32(_a_F_index_concurrently_build_0)
		v36 = *(*int32)(unsafe.Add(mBase, _c_F_index_concurrently_build[2]))
		v38 = v36 + int32(1)
		*(*int32)(unsafe.Add(mBase, _c_F_index_concurrently_build[2])) = v38
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
							*(*int32)(unsafe.Add(mBase, _c_F_index_concurrently_build[1])) = v57
							*(*int32)(unsafe.Add(mBase, _c_F_index_concurrently_build[0])) = v56
							F_relation_close(m, v12, int32(0))
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int64
	_ = v26
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
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
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(208)
	m.G0 = v16
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+206)) = uint16(v5)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	base.MemoryFill(m, v16-int32(-64), v5, int32(128))
	v26 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+56)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v26
	if v20 <= int32(32) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v190 = v16 - int32(-64)
	v191 = F_heap_compute_data_size(m, l0, v190, l2)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L8
	} else {
		goto L38
	}
L2:
	;
	v175 = int32(0)
	v182 = v167
	v187 = int32(8)
	goto L1
L3:
	;
	v66 = v5
	goto L13
L4:
	;
	if int32(0) < v20 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v167 = v5
	goto L2
L8:
	;
	return int32(0)
L9:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v20
	F_errmsg(m, int32(_a_F_index_form_tuple_context_0), v16)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_index_form_tuple_context_1), int32(90), int32(_a_F_index_form_tuple_context_2))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v72 = v66 << (uint(int32(2)) % 32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1+v72)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v78 = v16 + int32(32) + v66
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v79)
	v83 = v16 - int32(-64) + v72
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v74
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v66))))
	if v86 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v142 = int32(0)
	goto L32
L15:
	;
	v134 = v66 + int32(1)
	if v134 != v20 {
		v66 = v134
		goto L13
	} else {
		goto L31
	}
L16:
	;
	v94 = l0 + v75<<(uint(int32(4))%32) + v66*int32(100) + int32(20)
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+72)))
	if v95 != int32(_a_F_index_form_tuple_context_3) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	v100 = base.B2i32(v98 != int32(1))
	if v98 != int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v107 = v74
	v108 = v98
	goto L20
L19:
	;
	v101 = F_detoast_external_attr(m, v74)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L8
	} else {
		goto L21
	}
L20:
	;
	if v108&int32(3) != 0 {
		goto L15
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v101
	v104 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v104)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	v107 = v101
	v108 = v106
	goto L20
L22:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if base.Ui32(v111) < base.Ui32(int32(2044)) {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+84)))
	switch v114 - int32(109) {
	case 0, 11:
		goto L24
	default:
		goto L15
	}
L24:
	;
	v117 = int32(*(*int8)(unsafe.Add(mBase, uint32(v94)+85)))
	v118 = F_toast_compress_datum(m, v107, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	if v118 == int32(0) {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	if v100 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_pfree(m, v107)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L8
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v118
	v127 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v127)
	goto L15
L30:
	;
	goto L29
L31:
	;
	goto L14
L32:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v142))))
	if v151 != int32(1) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v175 = int32(_a_F_index_form_tuple_context_4)
	v182 = v151
	v187 = int32(16)
	goto L1
L34:
	;
	v155 = v142 + int32(1)
	if v20 != v155 {
		v142 = v155
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	goto L33
L37:
	;
	v167 = v151
	goto L2
L38:
	;
	v195 = v191 + v187 + int32(7)
	v197 = v195 & int32(-8)
	v198 = F_MemoryContextAllocZero(m, l3, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	if v182 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v206 = v198 + int32(8)
	goto L42
L41:
	;
	v206 = int32(0)
	goto L42
L42:
	;
	F_heap_fill_tuple(m, l0, v190, l2, v198+v187, v16+int32(206), v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	if int32(0) < v20 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v216 = int32(0)
	goto L47
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(int32(_a_F_index_form_tuple_context_5)) <= base.Ui32(v195) {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(32)+v216))))
	if v227 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L46
L49:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v16-int32(-64)+v216<<(uint(int32(2))%32))))
	F_pfree(m, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L8
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v239 = v216 + int32(1)
	if v239 != v20 {
		v216 = v239
		goto L47
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	goto L48
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+206)))
	v282 = v276<<(uint(int32(13))%32)&int32(_a_F_index_form_tuple_context_6) | (v175 | v197)
	*(*uint16)(unsafe.Add(mBase, uint32(v198)+6)) = uint16(v282)
	m.G0 = v16 + int32(208)
	return v198
L57:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L8
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = int32(_a_F_index_form_tuple_context_7)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v197
	F_errmsg(m, int32(_a_F_index_form_tuple_context_8), v16+int32(16))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L8
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_index_form_tuple_context_1), int32(210), int32(_a_F_index_form_tuple_context_2))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L8
	} else {
		goto L60
	}
L60:
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
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_index_getnext_tid_0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v62 + int32(4)
			F_errmsg_internal(m, int32(_a_F_index_getnext_tid_1), v8)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_index_getnext_tid_2), int32(626), int32(_a_F_index_getnext_tid_3))
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_index_insert[0]))
	if v18 != v16 {
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_index_insert[1]))
		v22 = F_list_member_ptr(m, v21, v16)
		mBase = m.M
		v24 = v22
	} else {
		v24 = int32(1)
	}
	if v24 == int32(0) {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
		if v28 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_index_insert_0)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v70 + int32(4)
				F_errmsg_internal(m, int32(_a_F_index_insert_1), v14+int32(16))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_index_insert_2), int32(223), int32(_a_F_index_insert_3))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+23)))
			if v31 != 0 {
				v40 = v28
				v41 = m.T0[v40].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3, l4, l5, l6, l7)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					m.G0 = v14 + int32(32)
					return v41
				}
			} else {
				F_CheckForSerializableConflictIn(m, l0, int32(0), int32(-1))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+44))
					v40 = v39
					v41 = m.T0[v40].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3, l4, l5, l6, l7)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						m.G0 = v14 + int32(32)
						return v41
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v54 + int32(4)
				F_errmsg(m, int32(_a_F_index_insert_4), v14)
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_index_insert_2), int32(222), int32(_a_F_index_insert_3))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
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
		v50 = *(*int32)(unsafe.Add(mBase, _c_F_makeIndexInfo[0]))
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
	F_relation_close(m, v23, int32(3))
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
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_mark_index_clustered[0]))
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
	F_errmsg(m, int32(_a_F_mark_index_clustered_0), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_mark_index_clustered_1), int32(565), int32(_a_F_mark_index_clustered_2))
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
	F_errmsg_internal(m, int32(_a_F_mark_index_clustered_3), v13)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L10
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_mark_index_clustered_1), int32(588), int32(_a_F_mark_index_clustered_2))
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
	F_errmsg_internal(m, int32(_a_F_mark_index_clustered_4), v13+int32(16))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_mark_index_clustered_1), int32(604), int32(_a_F_mark_index_clustered_2))
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
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
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
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v297 int32
	_ = v297
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v351 int32
	_ = v351
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
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v388 int32
	_ = v388
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v628 int32
	_ = v628
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v678 int32
	_ = v678
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
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
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v998 int32
	_ = v998
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1075 int32
	_ = v1075
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int64
	_ = v1121
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1141 int32
	_ = v1141
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1161 int32
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1242 int32
	_ = v1242
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
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1277 int32
	_ = v1277
	var v1280 int32
	_ = v1280
	var v1286 int32
	_ = v1286
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1299 int32
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1312 int32
	_ = v1312
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1320 int32
	_ = v1320
	var v1325 int32
	_ = v1325
	var v1329 int32
	_ = v1329
	var v1332 int32
	_ = v1332
	var v1338 int32
	_ = v1338
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1364 int32
	_ = v1364
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1377 int32
	_ = v1377
	var v1381 int32
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1394 int32
	_ = v1394
	var v1399 int32
	_ = v1399
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1426 int32
	_ = v1426
	var v1430 int32
	_ = v1430
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1440 int32
	_ = v1440
	var v1445 int32
	_ = v1445
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
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
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1482 int32
	_ = v1482
	var v1504 int32
	_ = v1504
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1514 int32
	_ = v1514
	var v1517 int32
	_ = v1517
	var v1521 int32
	_ = v1521
	var v1526 int32
	_ = v1526
	var v1550 int32
	_ = v1550
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1564 int32
	_ = v1564
	var v1578 int32
	_ = v1578
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1590 int32
	_ = v1590
	var v1593 int32
	_ = v1593
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1623 int32
	_ = v1623
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1634 int32
	_ = v1634
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1647 int32
	_ = v1647
	var v1671 int32
	_ = v1671
	var v1676 int32
	_ = v1676
	var v1681 int32
	_ = v1681
	var v1686 int32
	_ = v1686
	var v1691 int32
	_ = v1691
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1704 int32
	_ = v1704
	var v1711 int32
	_ = v1711
	var v1729 int32
	_ = v1729
	var v1733 int32
	_ = v1733
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1740 int32
	_ = v1740
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1761 int32
	_ = v1761
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1788 int32
	_ = v1788
	var v1791 int32
	_ = v1791
	var v1794 int32
	_ = v1794
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1805 int32
	_ = v1805
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1818 int32
	_ = v1818
	var v1843 int32
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1869 int32
	_ = v1869
	var v1873 int32
	_ = v1873
	var v1876 int32
	_ = v1876
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1886 int32
	_ = v1886
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1923 int64
	_ = v1923
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1962 int32
	_ = v1962
	var v1963 int32
	_ = v1963
	var v1965 int32
	_ = v1965
	var v1970 int32
	_ = v1970
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1986 int32
	_ = v1986
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v2004 int32
	_ = v2004
	var v2015 int32
	_ = v2015
	var v2022 int32
	_ = v2022
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2036 int32
	_ = v2036
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
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
	var v2089 int32
	_ = v2089
	var v2092 int32
	_ = v2092
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2103 int32
	_ = v2103
	var v2110 int32
	_ = v2110
	var v2111 int32
	_ = v2111
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2170 int32
	_ = v2170
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2195 int32
	_ = v2195
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2212 int32
	_ = v2212
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2222 int32
	_ = v2222
	var v2227 int32
	_ = v2227
	v2 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(320)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v26 == v2 {
		v1965 = l0
		v1970 = v24
		v1979 = v2
		v1980 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2212 = m.ExcPending
	if v2212 != 0 {
		goto L7
	} else {
		goto L478
	}
L2:
	;
	v1986 = *(*int32)(unsafe.Add(mBase, uint32(v1965)+56))
	if v1986 != 0 {
		goto L437
	} else {
		goto L438
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v29 <= int32(0) {
		v1965 = l0
		v1970 = v24
		v1979 = v2
		v1980 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v32 = l0
	v37 = v24
	v46 = v2
	v47 = v2
	v48 = v26
	v52 = v2
	goto L5
L5:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v52<<(uint(int32(2))%32))))
	v59 = F_palloc0(m, int32(72))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v1965 = v32
	v1970 = v37
	v1979 = v46
	v1980 = v1959
	goto L2
L7:
	;
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = int32(204)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+60)) = uint8(base.B2i32(v63 != int32(8)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v69 = base.B2i32(v67 == int32(6))
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+62)) = uint8(v69)
	if v67 == int32(6) {
		goto L26
	} else {
		goto L27
	}
L9:
	;
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v57)+40))
	if v1550 == int32(0) {
		goto L357
	} else {
		goto L358
	}
L10:
	;
	v1504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+36)))
	if v1504 == int32(0) {
		goto L9
	} else {
		goto L347
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L7
	} else {
		goto L342
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+64)) = v404
	F_errmsg(m, int32(_a_F_transformIndexConstraints_0), v37-int32(-64))
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L7
	} else {
		goto L339
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		goto L7
	} else {
		goto L335
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L7
	} else {
		goto L329
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L7
	} else {
		goto L324
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L7
	} else {
		goto L318
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L7
	} else {
		goto L312
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L7
	} else {
		goto L306
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L7
	} else {
		goto L300
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L7
	} else {
		goto L295
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L7
	} else {
		goto L290
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L7
	} else {
		goto L285
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L7
	} else {
		goto L280
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L7
	} else {
		goto L275
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L7
	} else {
		goto L270
	}
L26:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v32)+56))
	if v73 != 0 {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
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
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = v59
	goto L28
L30:
	;
	v86 = F_pstrdup(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L7
	} else {
		goto L33
	}
L31:
	;
	v89 = int32(0)
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v57)+64))
	if v93 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v89 = v86
	goto L32
L34:
	;
	v95 = v93
	goto L36
L35:
	;
	v95 = int32(_a_F_transformIndexConstraints_1)
	goto L36
L36:
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
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+21)))
	if v118 == int32(0) {
		goto L24
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v320 == int32(8) {
		goto L92
	} else {
		goto L93
	}
L40:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+48))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+68))
	v124 = F_get_relname_relid(m, v117, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	if v124 == int32(0) {
		goto L23
	} else {
		goto L42
	}
L42:
	;
	v129 = F_index_open(m, v124, int32(1))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+192))
	v132 = F_get_index_constraint(m, v124)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	if v132 != 0 {
		goto L22
	} else {
		goto L45
	}
L45:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v121)+56))
	if v134 != v135 {
		goto L21
	} else {
		goto L46
	}
L46:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+18)))
	if v137 == int32(0) {
		goto L20
	} else {
		goto L47
	}
L47:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+12)))
	if v140 == int32(0) {
		goto L19
	} else {
		goto L48
	}
L48:
	;
	v143 = F_RelationGetIndexExpressions(m, v129)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	if v143 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	v145 = F_RelationGetIndexPredicate(m, v129)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	if v145 != 0 {
		goto L17
	} else {
		goto L52
	}
L52:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+16)))
	if v147 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+12)))
	if v150 == int32(0) {
		goto L16
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v129)+48))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+84))
	v156 = F_get_index_am_oid(m, int32(_a_F_transformIndexConstraints_1))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L7
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	if v154 != v156 {
		goto L15
	} else {
		goto L58
	}
L58:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v129)+196))
	v162 = F_SysCacheGetAttrNotNull(m, int32(34), v160, int32(18))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	v164 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131)+8)))
	if int32(0) < v164 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v173 = int32(0)
	goto L63
L61:
	;
	goto L62
L62:
	;
	F_relation_close(m, v129, int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L7
	} else {
		goto L91
	}
L63:
	;
	v194 = v173 << (uint(int32(1)) % 32)
	v196 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131+int32(48)+v194))))
	if int32(0) < v196 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L62
L65:
	;
	v215 = F_pstrdup(m, v212+int32(4))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L7
	} else {
		goto L70
	}
L66:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v121)+52))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v212 = v199 + v200<<(uint(int32(4))%32) + v196*int32(100) - int32(80)
	goto L65
L67:
	;
	goto L68
L68:
	;
	v209 = F_SystemAttributeDefinition(m, v196)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	v212 = v209
	goto L65
L70:
	;
	v217 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131)+10)))
	if v173 < v217 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v272 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131)+8)))
	if v269 < v272 {
		v173 = v269
		goto L63
	} else {
		goto L90
	}
L72:
	;
	v220 = v173 << (uint(int32(2)) % 32)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v129)+56))
	v223 = v173 + int32(1)
	v225 = F_get_attoptions(m, v221, base.I32_extend16_s(v223))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L7
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v57)+40))
	v262 = F_makeString(m, v215)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L7
	} else {
		goto L88
	}
L75:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v212)+68))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v129)+48))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+84))
	v230 = F_GetDefaultOpClass(m, v227, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v220+(v162+int32(24)))))
	if v230 != v233 {
		goto L14
	} else {
		goto L77
	}
L77:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v212)+96))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v129)+248))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v236+v220)))
	if base.B2i32(v235 != v238)|v225 != 0 {
		goto L14
	} else {
		goto L78
	}
L78:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v129)+224))
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v241+v194))))
	if v243 != 0 {
		goto L14
	} else {
		goto L79
	}
L79:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v244 == int32(6) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
	v248 = F_makeString(m, v215)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L7
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v57)+32))
	v256 = F_makeString(m, v215)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L7
	} else {
		goto L86
	}
L83:
	;
	v250 = F_makeNotNullConstraint(m, v248)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L7
	} else {
		goto L84
	}
L84:
	;
	v252 = F_lappend(m, v247, v250)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L7
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v252
	goto L82
L86:
	;
	v258 = F_lappend(m, v255, v256)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L7
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+32)) = v258
	v269 = v223
	goto L71
L88:
	;
	v264 = F_lappend(m, v261, v262)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L7
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+40)) = v264
	v269 = v173 + int32(1)
	goto L71
L90:
	;
	goto L64
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+44)) = v124
	goto L39
L92:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	if v323 == int32(0) {
		goto L9
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v57)+32))
	if v371 == int32(0) {
		goto L10
	} else {
		goto L102
	}
L95:
	;
	v326 = int32(0)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	if v327 <= v326 {
		goto L9
	} else {
		goto L96
	}
L96:
	;
	v331 = v326
	goto L97
L97:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v323)+12))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v351+v331<<(uint(int32(2))%32))))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)+12))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v360 = F_lappend(m, v358, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L7
	} else {
		goto L99
	}
L98:
	;
	goto L9
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = v360
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v59)+36))
	v364 = F_lappend(m, v363, v357)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L7
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+36)) = v364
	v368 = v331 + int32(1)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	if v368 < v369 {
		v331 = v368
		goto L97
	} else {
		goto L101
	}
L101:
	;
	goto L98
L102:
	;
	v374 = int32(0)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
	if v375 <= v374 {
		goto L10
	} else {
		goto L103
	}
L103:
	;
	v388 = v374
	goto L104
L104:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v371)+12))
	v402 = v399 + v388<<(uint(int32(2))%32)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v405 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	goto L10
L106:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	if v858 == int32(0) {
		goto L213
	} else {
		goto L214
	}
L107:
	;
	v827 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v437)+19)) = uint8(v827)
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
	v830 = F_makeString(m, v404)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L7
	} else {
		goto L210
	}
L108:
	;
	v585 = int32(0)
	v586 = int32(1)
	v588 = F_strcmp(m, int32(_a_F_transformIndexConstraints_2), v404)
	mBase = m.M
	if v588 == v585 {
		goto L148
	} else {
		goto L149
	}
L109:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	if v408 <= int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v405)+12))
	v414 = int32(0)
	goto L111
L111:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v411+v414<<(uint(int32(2))%32))))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404))))
	if base.B2i32(v441 == int32(0))|base.B2i32(v441 != v444) != 0 {
		v462 = v441
		v463 = v444
		goto L114
	} else {
		goto L115
	}
L112:
	;
	v468 = int32(0)
	v469 = int32(1)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v470 != int32(6) {
		v840 = v468
		v844 = v437
		v845 = v469
		goto L106
	} else {
		goto L124
	}
L113:
	;
	if v462-v463 != 0 {
		goto L120
	} else {
		goto L121
	}
L114:
	;
	goto L113
L115:
	;
	v447 = v438
	v448 = v404
	goto L116
L116:
	;
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+1)))
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+1)))
	if v452 == int32(0) {
		v462 = v452
		v463 = v451
		goto L114
	} else {
		goto L118
	}
L117:
	;
	v462 = v452
	v463 = v451
	goto L114
L118:
	;
	v455 = int32(1)
	if v452 == v451 {
		v447 = v447 + v455
		v448 = v448 + v455
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v466 = v414 + int32(1)
	if v466 != v408 {
		v414 = v466
		goto L111
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	goto L112
L123:
	;
	goto L108
L124:
	;
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+21)))
	if v473 != 0 {
		v840 = v468
		v844 = v437
		v845 = v469
		goto L106
	} else {
		goto L125
	}
L125:
	;
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437)+19)))
	if v474 != int32(1) {
		goto L107
	} else {
		goto L126
	}
L126:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
	if v477 == int32(0) {
		v840 = v468
		v844 = v437
		v845 = v469
		goto L106
	} else {
		goto L127
	}
L127:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v477)+4))
	if v480 <= int32(0) {
		v840 = v468
		v844 = v437
		v845 = v469
		goto L106
	} else {
		goto L128
	}
L128:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v477)+12))
	v486 = int32(0)
	goto L129
L129:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v483+v486<<(uint(int32(2))%32))))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)+32))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v510)+12))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v511)))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v512)+4))
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513))))
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404))))
	if base.B2i32(v516 == int32(0))|base.B2i32(v516 != v519) != 0 {
		v537 = v516
		v538 = v519
		goto L132
	} else {
		goto L133
	}
L130:
	;
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+17)))
	if v543 != int32(1) {
		v840 = v468
		v844 = v437
		v845 = v469
		goto L106
	} else {
		goto L142
	}
L131:
	;
	if v537-v538 != 0 {
		goto L138
	} else {
		goto L139
	}
L132:
	;
	goto L131
L133:
	;
	v522 = v513
	v523 = v404
	goto L134
L134:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523)+1)))
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522)+1)))
	if v527 == int32(0) {
		v537 = v527
		v538 = v526
		goto L132
	} else {
		goto L136
	}
L135:
	;
	v537 = v527
	v538 = v526
	goto L132
L136:
	;
	v530 = int32(1)
	if v527 == v526 {
		v522 = v522 + v530
		v523 = v523 + v530
		goto L134
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	v541 = v486 + int32(1)
	if v541 != v480 {
		v486 = v541
		goto L129
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	goto L130
L141:
	;
	v840 = v468
	v844 = v437
	v845 = v469
	goto L106
L142:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L7
	} else {
		goto L143
	}
L143:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L7
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+128)) = v404
	F_errmsg(m, int32(_a_F_transformIndexConstraints_3), v37+int32(128))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L7
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2656), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L7
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	if v617 != 0 {
		goto L166
	} else {
		goto L167
	}
L148:
	;
	v617 = int32(_a_F_transformIndexConstraints_6)
	goto L147
L149:
	;
	goto L150
L150:
	;
	v593 = F_strcmp(m, int32(_a_F_transformIndexConstraints_7), v404)
	mBase = m.M
	if v593 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v617 = int32(_a_F_transformIndexConstraints_8)
	goto L147
L152:
	;
	goto L153
L153:
	;
	v598 = F_strcmp(m, int32(_a_F_transformIndexConstraints_9), v404)
	mBase = m.M
	if v598 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v617 = int32(_a_F_transformIndexConstraints_10)
	goto L147
L155:
	;
	goto L156
L156:
	;
	v603 = F_strcmp(m, int32(_a_F_transformIndexConstraints_11), v404)
	mBase = m.M
	if v603 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v617 = int32(_a_F_transformIndexConstraints_12)
	goto L147
L158:
	;
	goto L159
L159:
	;
	v608 = F_strcmp(m, int32(_a_F_transformIndexConstraints_13), v404)
	mBase = m.M
	if v608 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v617 = int32(_a_F_transformIndexConstraints_14)
	goto L147
L161:
	;
	goto L162
L162:
	;
	v615 = F_strcmp(m, int32(_a_F_transformIndexConstraints_15), v404)
	mBase = m.M
	if v615 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v616 = int32(0)
	goto L165
L164:
	;
	v616 = int32(_a_F_transformIndexConstraints_16)
	goto L165
L165:
	;
	v617 = v616
	goto L147
L166:
	;
	v840 = v585
	v844 = int32(0)
	v845 = v586
	goto L106
L167:
	;
	goto L168
L168:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v619 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+21)))
	if v801 != 0 {
		goto L202
	} else {
		goto L203
	}
L170:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v619)+4))
	if v622 <= int32(0) {
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v628 = v585
	goto L172
L172:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v619)+12))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v646+v628<<(uint(int32(2))%32))))
	v652 = F_table_openrv(m, v650, int32(1))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L7
	} else {
		goto L174
	}
L173:
	;
	goto L169
L174:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v652)+48))
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654)+119)))
	v657 = v655 - int32(102)
	if base.B2i32(base.Ui32(int32(12)) < base.Ui32(v657))|base.B2i32(int32(1)<<(uint(v657)%32)&int32(_a_F_transformIndexConstraints_17) == int32(0)) != 0 {
		goto L13
	} else {
		goto L175
	}
L175:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v652)+52))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v667)))
	if int32(0) < v668 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v678 = int32(0)
	goto L179
L177:
	;
	goto L178
L178:
	;
	F_relation_close(m, v652, int32(0))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L7
	} else {
		goto L200
	}
L179:
	;
	v700 = v667 + v668<<(uint(int32(4))%32) + int32(20) + v678*int32(100)
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700)+91)))
	if v701 != 0 {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	goto L178
L181:
	;
	v750 = v678 + int32(1)
	if v750 != v668 {
		v678 = v750
		goto L179
	} else {
		goto L199
	}
L182:
	;
	v703 = v700 + int32(4)
	v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404))))
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703))))
	if base.B2i32(v706 == int32(0))|base.B2i32(v706 != v709) != 0 {
		v727 = v706
		v728 = v709
		goto L184
	} else {
		goto L185
	}
L183:
	;
	if v727-v728 != 0 {
		goto L181
	} else {
		goto L190
	}
L184:
	;
	goto L183
L185:
	;
	v712 = v404
	v713 = v703
	goto L186
L186:
	;
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713)+1)))
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v712)+1)))
	if v717 == int32(0) {
		v727 = v717
		v728 = v716
		goto L184
	} else {
		goto L188
	}
L187:
	;
	v727 = v717
	v728 = v716
	goto L184
L188:
	;
	v720 = int32(1)
	if v717 == v716 {
		v712 = v712 + v720
		v713 = v713 + v720
		goto L186
	} else {
		goto L189
	}
L189:
	;
	goto L187
L190:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v700)+68))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v731 == int32(6) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
	v735 = F_pstrdup(m, v703)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L7
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v744 = int32(0)
	F_relation_close(m, v652, v744)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L7
	} else {
		goto L198
	}
L194:
	;
	v737 = F_makeString(m, v735)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L7
	} else {
		goto L195
	}
L195:
	;
	v739 = F_makeNotNullConstraint(m, v737)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L7
	} else {
		goto L196
	}
L196:
	;
	v741 = F_lappend(m, v734, v739)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L7
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v741
	goto L193
L198:
	;
	v840 = v730
	v844 = v744
	v845 = v586
	goto L106
L199:
	;
	goto L180
L200:
	;
	v777 = v628 + int32(1)
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v619)+4))
	if v777 < v778 {
		v628 = v777
		goto L172
	} else {
		goto L201
	}
L201:
	;
	goto L173
L202:
	;
	v802 = int32(0)
	v840 = v802
	v844 = v802
	v845 = v802
	goto L106
L203:
	;
	goto L204
L204:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L7
	} else {
		goto L205
	}
L205:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L7
	} else {
		goto L206
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+96)) = v404
	F_errmsg(m, int32(_a_F_transformIndexConstraints_18), v37+int32(96))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L7
	} else {
		goto L207
	}
L207:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v818, v819)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L7
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2736), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L7
	} else {
		goto L209
	}
L209:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L210:
	;
	v832 = F_makeNotNullConstraint(m, v830)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L7
	} else {
		goto L211
	}
L211:
	;
	v834 = F_lappend(m, v829, v832)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L7
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v834
	v840 = v468
	v844 = v437
	v845 = v469
	goto L106
L213:
	;
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+36)))
	if v969 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L214:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v858)+4))
	if v861 <= int32(0) {
		goto L213
	} else {
		goto L215
	}
L215:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v858)+12))
	v867 = int32(0)
	goto L216
L216:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v864+v867<<(uint(int32(2))%32))))
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v890)+4))
	if v891 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	goto L213
L218:
	;
	v946 = v867 + int32(1)
	if v861 != v946 {
		v867 = v946
		goto L216
	} else {
		goto L234
	}
L219:
	;
	v896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404))))
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v891))))
	if base.B2i32(v896 == int32(0))|base.B2i32(v896 != v899) != 0 {
		v917 = v896
		v918 = v899
		goto L221
	} else {
		goto L222
	}
L220:
	;
	if v917-v918 != 0 {
		goto L218
	} else {
		goto L227
	}
L221:
	;
	goto L220
L222:
	;
	v902 = v404
	v903 = v891
	goto L223
L223:
	;
	v906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v903)+1)))
	v907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902)+1)))
	if v907 == int32(0) {
		v917 = v907
		v918 = v906
		goto L221
	} else {
		goto L225
	}
L224:
	;
	v917 = v907
	v918 = v906
	goto L221
L225:
	;
	v910 = int32(1)
	if v907 == v906 {
		v902 = v902 + v910
		v903 = v903 + v910
		goto L223
	} else {
		goto L226
	}
L226:
	;
	goto L224
L227:
	;
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+62)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L7
	} else {
		goto L228
	}
L228:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L7
	} else {
		goto L229
	}
L229:
	;
	if v920 == int32(1) {
		goto L12
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+80)) = v404
	F_errmsg(m, int32(_a_F_transformIndexConstraints_19), v37+int32(80))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L7
	} else {
		goto L231
	}
L231:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v936, v937)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L7
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2755), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L7
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	goto L217
L235:
	;
	v1115 = F_palloc0(m, int32(36))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L7
	} else {
		goto L266
	}
L236:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v57)+32))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v972)+12))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v972)+4))
	if v402 != v973+v974<<(uint(int32(2))%32)-int32(4) {
		goto L235
	} else {
		goto L237
	}
L237:
	;
	if v845 == int32(0) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+21)))
	if v983 != int32(1) {
		goto L235
	} else {
		goto L241
	}
L239:
	;
	v1057 = v840
	goto L240
L240:
	;
	v1075 = int32(0)
	if base.B2i32(v844 == v1075)|v1057 == v1075 {
		goto L257
	} else {
		goto L258
	}
L241:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v986)+52))
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v987)))
	if v988 <= int32(0) {
		goto L235
	} else {
		goto L242
	}
L242:
	;
	v998 = int32(0)
	goto L243
L243:
	;
	v1020 = v987 + v988<<(uint(int32(4))%32) + int32(20) + v998*int32(100)
	v1021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1020)+91)))
	if v1021 != 0 {
		goto L235
	} else {
		goto L245
	}
L244:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1020)+68))
	v1057 = v1053
	goto L240
L245:
	;
	v1023 = v1020 + int32(4)
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1023))))
	v1029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404))))
	if base.B2i32(v1026 == int32(0))|base.B2i32(v1026 != v1029) != 0 {
		v1047 = v1026
		v1048 = v1029
		goto L247
	} else {
		goto L248
	}
L246:
	;
	if v1047-v1048 != 0 {
		goto L253
	} else {
		goto L254
	}
L247:
	;
	goto L246
L248:
	;
	v1032 = v1023
	v1033 = v404
	goto L249
L249:
	;
	v1036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1033)+1)))
	v1037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1032)+1)))
	if v1037 == int32(0) {
		v1047 = v1037
		v1048 = v1036
		goto L247
	} else {
		goto L251
	}
L250:
	;
	v1047 = v1037
	v1048 = v1036
	goto L247
L251:
	;
	v1040 = int32(1)
	if v1037 == v1036 {
		v1032 = v1032 + v1040
		v1033 = v1033 + v1040
		goto L249
	} else {
		goto L252
	}
L252:
	;
	goto L250
L253:
	;
	v1051 = v998 + int32(1)
	if v1051 == v988 {
		goto L235
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	goto L244
L256:
	;
	v998 = v1051
	goto L243
L257:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v844)+8))
	v1082 = F_typenameTypeId(m, int32(0), v1081)
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L7
	} else {
		goto L260
	}
L258:
	;
	v1084 = v1057
	goto L259
L259:
	;
	if v1084 == int32(0) {
		goto L11
	} else {
		goto L261
	}
L260:
	;
	v1084 = v1082
	goto L259
L261:
	;
	v1087 = F_type_is_range(m, v1084)
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L7
	} else {
		goto L262
	}
L262:
	;
	if v1087 != 0 {
		goto L235
	} else {
		goto L263
	}
L263:
	;
	v1089 = F_type_is_multirange(m, v1084)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L7
	} else {
		goto L264
	}
L264:
	;
	if v1089 == int32(0) {
		goto L11
	} else {
		goto L265
	}
L265:
	;
	goto L235
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1115))) = int32(92)
	v1119 = F_pstrdup(m, v404)
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L7
	} else {
		goto L267
	}
L267:
	;
	v1121 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1115)+8)) = v1121
	*(*int32)(unsafe.Add(mBase, uint32(v1115)+4)) = v1119
	*(*int64)(unsafe.Add(mBase, uint32(v1115)+16)) = v1121
	*(*int64)(unsafe.Add(mBase, uint32(v1115)+24)) = v1121
	*(*int32)(unsafe.Add(mBase, uint32(v1115)+32)) = int32(0)
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	v1131 = F_lappend(m, v1130, v1115)
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L7
	} else {
		goto L268
	}
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = v1131
	v1135 = v388 + int32(1)
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
	if v1135 < v1136 {
		v388 = v1135
		goto L104
	} else {
		goto L269
	}
L269:
	;
	goto L105
L270:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L7
	} else {
		goto L271
	}
L271:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+304)) = v1146
	F_errmsg(m, int32(_a_F_transformIndexConstraints_20), v37+int32(304))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L7
	} else {
		goto L272
	}
L272:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1153, v1154)
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L7
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2354), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L7
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L275:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L7
	} else {
		goto L276
	}
L276:
	;
	F_errmsg(m, int32(_a_F_transformIndexConstraints_21), int32(0))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L7
	} else {
		goto L277
	}
L277:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1173, v1174)
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L7
	} else {
		goto L278
	}
L278:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2420), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L7
	} else {
		goto L279
	}
L279:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L280:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L7
	} else {
		goto L281
	}
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+144)) = v117
	F_errmsg(m, int32(_a_F_transformIndexConstraints_22), v37+int32(144))
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L7
	} else {
		goto L282
	}
L282:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1195, v1196)
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L7
	} else {
		goto L283
	}
L283:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2429), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L7
	} else {
		goto L284
	}
L284:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L285:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L7
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+288)) = v117
	F_errmsg(m, int32(_a_F_transformIndexConstraints_23), v37+int32(288))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L7
	} else {
		goto L287
	}
L287:
	;
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1217, v1218)
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L7
	} else {
		goto L288
	}
L288:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2441), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L7
	} else {
		goto L289
	}
L289:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L290:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L7
	} else {
		goto L291
	}
L291:
	;
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v121)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+272)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v37)+276)) = v1233 + int32(4)
	F_errmsg(m, int32(_a_F_transformIndexConstraints_24), v37+int32(272))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L7
	} else {
		goto L292
	}
L292:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1243, v1244)
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L7
	} else {
		goto L293
	}
L293:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2449), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L7
	} else {
		goto L294
	}
L294:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L295:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L7
	} else {
		goto L296
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+256)) = v117
	F_errmsg(m, int32(_a_F_transformIndexConstraints_25), v37+int32(256))
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L7
	} else {
		goto L297
	}
L297:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1265, v1266)
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L7
	} else {
		goto L298
	}
L298:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2455), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L7
	} else {
		goto L299
	}
L299:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L300:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L7
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v117
	F_errmsg(m, int32(_a_F_transformIndexConstraints_26), v37+int32(240))
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L7
	} else {
		goto L302
	}
L302:
	;
	F_errdetail(m, int32(_a_F_transformIndexConstraints_27), int32(0))
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L7
	} else {
		goto L303
	}
L303:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1291, v1292)
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L7
	} else {
		goto L304
	}
L304:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2467), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L7
	} else {
		goto L305
	}
L305:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L306:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		goto L7
	} else {
		goto L307
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+224)) = v117
	F_errmsg(m, int32(_a_F_transformIndexConstraints_28), v37+int32(224))
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L7
	} else {
		goto L308
	}
L308:
	;
	F_errdetail(m, int32(_a_F_transformIndexConstraints_27), int32(0))
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L7
	} else {
		goto L309
	}
L309:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1317, v1318)
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L7
	} else {
		goto L310
	}
L310:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2474), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L7
	} else {
		goto L311
	}
L311:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L312:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L7
	} else {
		goto L313
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+208)) = v117
	F_errmsg(m, int32(_a_F_transformIndexConstraints_29), v37+int32(208))
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L7
	} else {
		goto L314
	}
L314:
	;
	F_errdetail(m, int32(_a_F_transformIndexConstraints_27), int32(0))
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L7
	} else {
		goto L315
	}
L315:
	;
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1343, v1344)
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L7
	} else {
		goto L316
	}
L316:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2481), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L7
	} else {
		goto L317
	}
L317:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L318:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L7
	} else {
		goto L319
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+192)) = v117
	F_errmsg(m, int32(_a_F_transformIndexConstraints_30), v37+int32(192))
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L7
	} else {
		goto L320
	}
L320:
	;
	F_errdetail(m, int32(_a_F_transformIndexConstraints_31), int32(0))
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L7
	} else {
		goto L321
	}
L321:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1369, v1370)
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L7
	} else {
		goto L322
	}
L322:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2493), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L7
	} else {
		goto L323
	}
L323:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L324:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L7
	} else {
		goto L325
	}
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+176)) = v117
	F_errmsg(m, int32(_a_F_transformIndexConstraints_32), v37+int32(176))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L7
	} else {
		goto L326
	}
L326:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1391, v1392)
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L7
	} else {
		goto L327
	}
L327:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2505), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L7
	} else {
		goto L328
	}
L328:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L329:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L7
	} else {
		goto L330
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+164)) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v37)+160)) = v117
	F_errmsg(m, int32(_a_F_transformIndexConstraints_33), v37+int32(160))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L7
	} else {
		goto L331
	}
L331:
	;
	F_errdetail(m, int32(_a_F_transformIndexConstraints_27), int32(0))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L7
	} else {
		goto L332
	}
L332:
	;
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1418, v1419)
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L7
	} else {
		goto L333
	}
L333:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2557), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L7
	} else {
		goto L334
	}
L334:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L335:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L7
	} else {
		goto L336
	}
L336:
	;
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v650)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+112)) = v1434
	F_errmsg(m, int32(_a_F_transformIndexConstraints_34), v37+int32(112))
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L7
	} else {
		goto L337
	}
L337:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2700), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L7
	} else {
		goto L338
	}
L338:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L339:
	;
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1452, v1453)
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L7
	} else {
		goto L340
	}
L340:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2749), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L7
	} else {
		goto L341
	}
L341:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L342:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L7
	} else {
		goto L343
	}
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+48)) = v404
	F_errmsg(m, int32(_a_F_transformIndexConstraints_35), v37+int32(48))
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L7
	} else {
		goto L344
	}
L344:
	;
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1474, v1475)
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L7
	} else {
		goto L345
	}
L345:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2799), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L7
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
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v57)+32))
	if v1507 != 0 {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = int32(_a_F_transformIndexConstraints_36)
	goto L9
L349:
	;
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+4))
	if int32(1) < v1508 {
		goto L348
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1514 = m.ExcPending
	if v1514 != 0 {
		goto L7
	} else {
		goto L353
	}
L352:
	;
	goto L351
L353:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		goto L7
	} else {
		goto L354
	}
L354:
	;
	F_errmsg(m, int32(_a_F_transformIndexConstraints_37), int32(0))
	mBase = m.M
	v1521 = m.ExcPending
	if v1521 != 0 {
		goto L7
	} else {
		goto L355
	}
L355:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2826), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L7
	} else {
		goto L356
	}
L356:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L357:
	;
	v1959 = F_lappend(m, v47, v59)
	mBase = m.M
	v1960 = m.ExcPending
	if v1960 != 0 {
		goto L7
	} else {
		goto L435
	}
L358:
	;
	v1553 = int32(0)
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(v1550)+4))
	if v1554 <= v1553 {
		goto L357
	} else {
		goto L359
	}
L359:
	;
	v1564 = v1553
	goto L360
L360:
	;
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(v1550)+12))
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v1578+v1564<<(uint(int32(2))%32))))
	v1583 = *(*int32)(unsafe.Add(mBase, uint32(v1582)+4))
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v1584 == int32(0) {
		goto L363
	} else {
		goto L364
	}
L361:
	;
	goto L357
L362:
	;
	v1917 = F_palloc0(m, int32(36))
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		goto L7
	} else {
		goto L431
	}
L363:
	;
	v1671 = F_strcmp(m, int32(_a_F_transformIndexConstraints_2), v1583)
	mBase = m.M
	if v1671 == int32(0) {
		goto L378
	} else {
		goto L379
	}
L364:
	;
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v1584)+4))
	if v1587 <= int32(0) {
		goto L363
	} else {
		goto L365
	}
L365:
	;
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v1584)+12))
	v1593 = int32(0)
	goto L366
L366:
	;
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v1590+v1593<<(uint(int32(2))%32))))
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1616)+4))
	v1620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1617))))
	v1623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1583))))
	if base.B2i32(v1620 == int32(0))|base.B2i32(v1620 != v1623) != 0 {
		v1641 = v1620
		v1642 = v1623
		goto L369
	} else {
		goto L370
	}
L367:
	;
	goto L363
L368:
	;
	if v1641-v1642 == int32(0) {
		goto L362
	} else {
		goto L375
	}
L369:
	;
	goto L368
L370:
	;
	v1626 = v1617
	v1627 = v1583
	goto L371
L371:
	;
	v1630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1627)+1)))
	v1631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1626)+1)))
	if v1631 == int32(0) {
		v1641 = v1631
		v1642 = v1630
		goto L369
	} else {
		goto L373
	}
L372:
	;
	v1641 = v1631
	v1642 = v1630
	goto L369
L373:
	;
	v1634 = int32(1)
	if v1631 == v1630 {
		v1626 = v1626 + v1634
		v1627 = v1627 + v1634
		goto L371
	} else {
		goto L374
	}
L374:
	;
	goto L372
L375:
	;
	v1647 = v1593 + int32(1)
	if v1587 != v1647 {
		v1593 = v1647
		goto L366
	} else {
		goto L376
	}
L376:
	;
	goto L367
L377:
	;
	if v1700 != 0 {
		goto L362
	} else {
		goto L396
	}
L378:
	;
	v1700 = int32(_a_F_transformIndexConstraints_6)
	goto L377
L379:
	;
	goto L380
L380:
	;
	v1676 = F_strcmp(m, int32(_a_F_transformIndexConstraints_7), v1583)
	mBase = m.M
	if v1676 == int32(0) {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v1700 = int32(_a_F_transformIndexConstraints_8)
	goto L377
L382:
	;
	goto L383
L383:
	;
	v1681 = F_strcmp(m, int32(_a_F_transformIndexConstraints_9), v1583)
	mBase = m.M
	if v1681 == int32(0) {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v1700 = int32(_a_F_transformIndexConstraints_10)
	goto L377
L385:
	;
	goto L386
L386:
	;
	v1686 = F_strcmp(m, int32(_a_F_transformIndexConstraints_11), v1583)
	mBase = m.M
	if v1686 == int32(0) {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v1700 = int32(_a_F_transformIndexConstraints_12)
	goto L377
L388:
	;
	goto L389
L389:
	;
	v1691 = F_strcmp(m, int32(_a_F_transformIndexConstraints_13), v1583)
	mBase = m.M
	if v1691 == int32(0) {
		goto L390
	} else {
		goto L391
	}
L390:
	;
	v1700 = int32(_a_F_transformIndexConstraints_14)
	goto L377
L391:
	;
	goto L392
L392:
	;
	v1698 = F_strcmp(m, int32(_a_F_transformIndexConstraints_15), v1583)
	mBase = m.M
	if v1698 != 0 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v1699 = int32(0)
	goto L395
L394:
	;
	v1699 = int32(_a_F_transformIndexConstraints_16)
	goto L395
L395:
	;
	v1700 = v1699
	goto L377
L396:
	;
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v1701 == int32(0) {
		goto L398
	} else {
		goto L399
	}
L397:
	;
	F_relation_close(m, v1735, int32(0))
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L7
	} else {
		goto L430
	}
L398:
	;
	v1869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+21)))
	if v1869 != 0 {
		goto L362
	} else {
		goto L424
	}
L399:
	;
	v1704 = *(*int32)(unsafe.Add(mBase, uint32(v1701)+4))
	if v1704 <= int32(0) {
		goto L398
	} else {
		goto L400
	}
L400:
	;
	v1711 = int32(0)
	goto L401
L401:
	;
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v1701)+12))
	v1733 = *(*int32)(unsafe.Add(mBase, uint32(v1729+v1711<<(uint(int32(2))%32))))
	v1735 = F_table_openrv(m, v1733, int32(1))
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L7
	} else {
		goto L403
	}
L402:
	;
	goto L398
L403:
	;
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(v1735)+48))
	v1738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1737)+119)))
	v1740 = v1738 - int32(102)
	if base.B2i32(base.Ui32(int32(12)) < base.Ui32(v1740))|base.B2i32(int32(1)<<(uint(v1740)%32)&int32(_a_F_transformIndexConstraints_17) == int32(0)) != 0 {
		goto L1
	} else {
		goto L404
	}
L404:
	;
	v1750 = *(*int32)(unsafe.Add(mBase, uint32(v1735)+52))
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1750)))
	if int32(0) < v1751 {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v1761 = int32(0)
	goto L408
L406:
	;
	goto L407
L407:
	;
	F_relation_close(m, v1735, int32(0))
	mBase = m.M
	v1843 = m.ExcPending
	if v1843 != 0 {
		goto L7
	} else {
		goto L422
	}
L408:
	;
	v1783 = v1750 + v1751<<(uint(int32(4))%32) + int32(20) + v1761*int32(100)
	v1784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1783)+91)))
	if v1784 == int32(0) {
		goto L410
	} else {
		goto L411
	}
L409:
	;
	goto L407
L410:
	;
	v1788 = v1783 + int32(4)
	v1791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1583))))
	v1794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1788))))
	if base.B2i32(v1791 == int32(0))|base.B2i32(v1791 != v1794) != 0 {
		v1812 = v1791
		v1813 = v1794
		goto L414
	} else {
		goto L415
	}
L411:
	;
	goto L412
L412:
	;
	v1818 = v1761 + int32(1)
	if v1818 != v1751 {
		v1761 = v1818
		goto L408
	} else {
		goto L421
	}
L413:
	;
	if v1812-v1813 == int32(0) {
		goto L397
	} else {
		goto L420
	}
L414:
	;
	goto L413
L415:
	;
	v1797 = v1583
	v1798 = v1788
	goto L416
L416:
	;
	v1801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1798)+1)))
	v1802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1797)+1)))
	if v1802 == int32(0) {
		v1812 = v1802
		v1813 = v1801
		goto L414
	} else {
		goto L418
	}
L417:
	;
	v1812 = v1802
	v1813 = v1801
	goto L414
L418:
	;
	v1805 = int32(1)
	if v1802 == v1801 {
		v1797 = v1797 + v1805
		v1798 = v1798 + v1805
		goto L416
	} else {
		goto L419
	}
L419:
	;
	goto L417
L420:
	;
	goto L412
L421:
	;
	goto L409
L422:
	;
	v1845 = v1711 + int32(1)
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1701)+4))
	if v1845 < v1846 {
		v1711 = v1845
		goto L401
	} else {
		goto L423
	}
L423:
	;
	goto L402
L424:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L7
	} else {
		goto L425
	}
L425:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v1876 = m.ExcPending
	if v1876 != 0 {
		goto L7
	} else {
		goto L426
	}
L426:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v1583
	F_errmsg(m, int32(_a_F_transformIndexConstraints_18), v37+int32(16))
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L7
	} else {
		goto L427
	}
L427:
	;
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	F_parser_errposition(m, v1883, v1884)
	mBase = m.M
	v1886 = m.ExcPending
	if v1886 != 0 {
		goto L7
	} else {
		goto L428
	}
L428:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2919), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1891 = m.ExcPending
	if v1891 != 0 {
		goto L7
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
	goto L362
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1917))) = int32(92)
	v1921 = F_pstrdup(m, v1583)
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L7
	} else {
		goto L432
	}
L432:
	;
	v1923 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1917)+8)) = v1923
	*(*int32)(unsafe.Add(mBase, uint32(v1917)+4)) = v1921
	*(*int64)(unsafe.Add(mBase, uint32(v1917)+16)) = v1923
	*(*int32)(unsafe.Add(mBase, uint32(v1917)+24)) = int32(0)
	v1930 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
	v1931 = F_lappend(m, v1930, v1917)
	mBase = m.M
	v1932 = m.ExcPending
	if v1932 != 0 {
		goto L7
	} else {
		goto L433
	}
L433:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+24)) = v1931
	v1935 = v1564 + int32(1)
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v1550)+4))
	if v1935 < v1936 {
		v1564 = v1935
		goto L360
	} else {
		goto L434
	}
L434:
	;
	goto L361
L435:
	;
	v1962 = v52 + int32(1)
	v1963 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v1962 < v1963 {
		v47 = v1959
		v52 = v1962
		goto L5
	} else {
		goto L436
	}
L436:
	;
	goto L6
L437:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1970)+12)) = v1986
	*(*int32)(unsafe.Add(mBase, uint32(v1970)+316)) = v1986
	v1992 = F_list_make1_impl(m, int32(1), v1970+int32(12))
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L7
	} else {
		goto L440
	}
L438:
	;
	v1994 = v1979
	goto L439
L439:
	;
	if v1980 == int32(0) {
		v2195 = v1994
		goto L441
	} else {
		goto L442
	}
L440:
	;
	v1994 = v1992
	goto L439
L441:
	;
	v2202 = *(*int32)(unsafe.Add(mBase, uint32(v1965)+52))
	v2203 = F_list_concat(m, v2202, v2195)
	mBase = m.M
	v2204 = m.ExcPending
	if v2204 != 0 {
		goto L7
	} else {
		goto L477
	}
L442:
	;
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(v1980)+4))
	if v1997 <= int32(0) {
		v2195 = v1994
		goto L441
	} else {
		goto L443
	}
L443:
	;
	v2004 = int32(0)
	v2015 = v1994
	goto L444
L444:
	;
	v2022 = *(*int32)(unsafe.Add(mBase, uint32(v1980)+12))
	v2026 = *(*int32)(unsafe.Add(mBase, uint32(v2022+v2004<<(uint(int32(2))%32))))
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(v1965)+56))
	if v2026 == v2027 {
		v2170 = v2015
		goto L446
	} else {
		goto L447
	}
L445:
	;
	v2195 = v2170
	goto L441
L446:
	;
	v2178 = v2004 + int32(1)
	v2179 = *(*int32)(unsafe.Add(mBase, uint32(v1980)+4))
	if v2178 < v2179 {
		v2004 = v2178
		v2015 = v2170
		goto L444
	} else {
		goto L476
	}
L447:
	;
	if v2015 == int32(0) {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	v2154 = F_lappend(m, v2015, v2026)
	mBase = m.M
	v2155 = m.ExcPending
	if v2155 != 0 {
		goto L7
	} else {
		goto L475
	}
L449:
	;
	v2031 = int32(0)
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(v2015)+4))
	if v2032 <= v2031 {
		goto L448
	} else {
		goto L450
	}
L450:
	;
	v2036 = v2031
	goto L451
L451:
	;
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(v2026)+20))
	v2057 = *(*int32)(unsafe.Add(mBase, uint32(v2015)+12))
	v2061 = *(*int32)(unsafe.Add(mBase, uint32(v2057+v2036<<(uint(int32(2))%32))))
	v2062 = *(*int32)(unsafe.Add(mBase, uint32(v2061)+20))
	v2063 = F_equal(m, v2056, v2062)
	mBase = m.M
	v2064 = m.ExcPending
	if v2064 != 0 {
		goto L7
	} else {
		goto L454
	}
L452:
	;
	goto L448
L453:
	;
	v2130 = v2036 + int32(1)
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v2015)+4))
	if v2130 < v2131 {
		v2036 = v2130
		goto L451
	} else {
		goto L474
	}
L454:
	;
	if v2063 == int32(0) {
		goto L453
	} else {
		goto L455
	}
L455:
	;
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(v2026)+24))
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v2061)+24))
	v2069 = F_equal(m, v2067, v2068)
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L7
	} else {
		goto L456
	}
L456:
	;
	if v2069 == int32(0) {
		goto L453
	} else {
		goto L457
	}
L457:
	;
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v2026)+32))
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(v2061)+32))
	v2075 = F_equal(m, v2073, v2074)
	mBase = m.M
	v2076 = m.ExcPending
	if v2076 != 0 {
		goto L7
	} else {
		goto L458
	}
L458:
	;
	if v2075 == int32(0) {
		goto L453
	} else {
		goto L459
	}
L459:
	;
	v2079 = *(*int32)(unsafe.Add(mBase, uint32(v2026)+36))
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v2061)+36))
	v2081 = F_equal(m, v2079, v2080)
	mBase = m.M
	v2082 = m.ExcPending
	if v2082 != 0 {
		goto L7
	} else {
		goto L460
	}
L460:
	;
	if v2081 == int32(0) {
		goto L453
	} else {
		goto L461
	}
L461:
	;
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(v2026)+12))
	v2086 = *(*int32)(unsafe.Add(mBase, uint32(v2061)+12))
	v2089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2085))))
	v2092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2086))))
	if base.B2i32(v2089 == int32(0))|base.B2i32(v2089 != v2092) != 0 {
		v2110 = v2089
		v2111 = v2092
		goto L463
	} else {
		goto L464
	}
L462:
	;
	if v2110-v2111 != 0 {
		goto L453
	} else {
		goto L469
	}
L463:
	;
	goto L462
L464:
	;
	v2095 = v2085
	v2096 = v2086
	goto L465
L465:
	;
	v2099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2096)+1)))
	v2100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2095)+1)))
	if v2100 == int32(0) {
		v2110 = v2100
		v2111 = v2099
		goto L463
	} else {
		goto L467
	}
L466:
	;
	v2110 = v2100
	v2111 = v2099
	goto L463
L467:
	;
	v2103 = int32(1)
	if v2100 == v2099 {
		v2095 = v2095 + v2103
		v2096 = v2096 + v2103
		goto L465
	} else {
		goto L468
	}
L468:
	;
	goto L466
L469:
	;
	v2113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2026)+61)))
	v2114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2061)+61)))
	if v2113 != v2114 {
		goto L453
	} else {
		goto L470
	}
L470:
	;
	v2116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2026)+65)))
	v2117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2061)+65)))
	if v2116 != v2117 {
		goto L453
	} else {
		goto L471
	}
L471:
	;
	v2119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2026)+66)))
	v2120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2061)+66)))
	if v2119 != v2120 {
		goto L453
	} else {
		goto L472
	}
L472:
	;
	v2122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2061)+60)))
	v2123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2026)+60)))
	v2124 = v2122 | v2123
	*(*uint8)(unsafe.Add(mBase, uint32(v2061)+60)) = uint8(v2124)
	v2126 = *(*int32)(unsafe.Add(mBase, uint32(v2061)+4))
	if v2126 != 0 {
		v2170 = v2015
		goto L446
	} else {
		goto L473
	}
L473:
	;
	v2127 = *(*int32)(unsafe.Add(mBase, uint32(v2026)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2061)+4)) = v2127
	v2170 = v2015
	goto L446
L474:
	;
	goto L452
L475:
	;
	v2170 = v2154
	goto L446
L476:
	;
	goto L445
L477:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1965)+52)) = v2203
	m.G0 = v1970 + int32(320)
	return
L478:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v2215 = m.ExcPending
	if v2215 != 0 {
		goto L7
	} else {
		goto L479
	}
L479:
	;
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(v1733)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = v2216
	F_errmsg(m, int32(_a_F_transformIndexConstraints_34), v37+int32(32))
	mBase = m.M
	v2222 = m.ExcPending
	if v2222 != 0 {
		goto L7
	} else {
		goto L480
	}
L480:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2888), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v2227 = m.ExcPending
	if v2227 != 0 {
		goto L7
	} else {
		goto L481
	}
L481:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
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
	v128 = m.ExcPending
	if v128 != 0 {
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
	v35 = F_transformWhereClause(m, v13, v32, int32(33), int32(_a_F_transformIndexStmt_0))
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
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v103 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v44 <= int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v50 = int32(0)
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
	v92 = v50 + int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v92 < v93 {
		v50 = v92
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
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v106 != int32(1) {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_free_parsestate(m, v13)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	F_relation_close(m, v19, int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v114 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+67)) = uint8(v114)
	goto L4
L34:
	;
	F_errcode(m, int32(_a_F_transformIndexStmt_1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	F_errmsg(m, int32(_a_F_transformIndexStmt_2), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_transformIndexStmt_3), int32(3123), int32(_a_F_transformIndexStmt_4))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
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
