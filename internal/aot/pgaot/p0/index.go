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
				F_errfinish(m, int32(_a_F_BuildIndexInfo_3), int32(2472), int32(_a_F_BuildIndexInfo_4))
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int64
	_ = v101
	var v102 int32
	_ = v102
	var v105 int64
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
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
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
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
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v367 int32
	_ = v367
	var v369 int64
	_ = v369
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v425 int32
	_ = v425
	var v427 int64
	_ = v427
	var v431 int64
	_ = v431
	var v434 int32
	_ = v434
	var v440 int64
	_ = v440
	var v441 int32
	_ = v441
	var v449 int64
	_ = v449
	var v453 int64
	_ = v453
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v634 int32
	_ = v634
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v688 int32
	_ = v688
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v744 int32
	_ = v744
	v6 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(48)
	m.G0 = v20
	v23 = F_IndexGetRelation(m, l0, v6)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
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
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v28 = v27
	goto L5
L4:
	;
	v28 = v6
	goto L5
L5:
	;
	v31 = F_SearchSysCache1(m, int32(1), base.I64_extend_i32_u(l1))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	m.G0 = v20 + int32(48)
	return v744
L7:
	;
	F_ReleaseCatCache(m, v76)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L180
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L177
	}
L9:
	;
	if v31 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
	v35 = v33 + v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+68))
	v38 = F_GetIndexAmRoutine(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
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
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L173
	}
L13:
	;
	F_ReleaseCatCache(m, v31)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+10)))
	v43 = int32(0)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+28)))
	v51 = F_makeIndexInfo(m, v28, v28, v36, v43, v43, v43, v43, v43, v43, v50, l4)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v54 = F_palloc_mul(m, int32(4), v28)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v57 = F_palloc_mul(m, int32(4), v28)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v60 = F_palloc_mul(m, int32(4), v28)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v63 = F_palloc_mul(m, int32(8), v28)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v66 = F_palloc_mul(m, int32(2), v28)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v68 = int32(0)
	F_ComputeIndexAttrs(m, v43, v51, v54, v57, v60, v63, v66, l2, l3, v23, l1, v36, v42, v68, l4, v68, v68, v68)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v76 = F_SearchSysCache1(m, int32(34), base.I64_extend_i32_u(l0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v76 == int32(0) {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+16))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+22)))
	v84 = F_heap_attisnull(m, v76, int32(21), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v84 == int32(0) {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v90 = F_heap_attisnull(m, v76, int32(20), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v90 == int32(0) {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	v94 = v80 + v81
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+18)))
	if v95 == int32(0) {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v98 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94)+10)))
	v101 = F_SysCacheGetAttrNotNull(m, int32(34), v76, int32(17))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v105 = F_SysCacheGetAttrNotNull(m, int32(34), v76, int32(18))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v109 = base.I32_wrap_i64(v105) + int32(24)
	v111 = v98 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v111) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	if v173 != 0 {
		goto L7
	} else {
		goto L49
	}
L32:
	;
	v173 = int32(0)
	goto L31
L33:
	;
	v147 = v142
	v148 = v143
	v149 = v144
	goto L43
L34:
	;
	if (v109|v60)&int32(3) != 0 {
		v142 = v109
		v143 = v60
		v144 = v111
		goto L33
	} else {
		goto L37
	}
L35:
	;
	v135 = v109
	v136 = v60
	v137 = v111
	goto L36
L36:
	;
	if v137 == int32(0) {
		goto L32
	} else {
		goto L42
	}
L37:
	;
	v119 = v109
	v120 = v60
	v121 = v111
	goto L38
L38:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v124 != v125 {
		v142 = v119
		v143 = v120
		v144 = v121
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v135 = v130
	v136 = v128
	v137 = v132
	goto L36
L40:
	;
	v127 = int32(4)
	v128 = v120 + v127
	v130 = v119 + v127
	v132 = v121 - v127
	if base.Ui32(int32(3)) < base.Ui32(v132) {
		v119 = v130
		v120 = v128
		v121 = v132
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v142 = v135
	v143 = v136
	v144 = v137
	goto L33
L43:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v152 == v153 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v173 = v152 - v153
	goto L31
L45:
	;
	v155 = int32(1)
	v160 = v149 - v155
	if v160 != 0 {
		v147 = v147 + v155
		v148 = v148 + v155
		v149 = v160
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
	v176 = base.I32_wrap_i64(v101) + int32(24)
	if base.Ui32(int32(4)) <= base.Ui32(v111) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	F_ReleaseCatCache(m, v76)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L68
	}
L51:
	;
	v238 = int32(0)
	goto L50
L52:
	;
	v212 = v207
	v213 = v208
	v214 = v209
	goto L62
L53:
	;
	if (v176|v57)&int32(3) != 0 {
		v207 = v176
		v208 = v57
		v209 = v111
		goto L52
	} else {
		goto L56
	}
L54:
	;
	v200 = v176
	v201 = v57
	v202 = v111
	goto L55
L55:
	;
	if v202 == int32(0) {
		goto L51
	} else {
		goto L61
	}
L56:
	;
	v184 = v176
	v185 = v57
	v186 = v111
	goto L57
L57:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	if v189 != v190 {
		v207 = v184
		v208 = v185
		v209 = v186
		goto L52
	} else {
		goto L59
	}
L58:
	;
	v200 = v195
	v201 = v193
	v202 = v197
	goto L55
L59:
	;
	v192 = int32(4)
	v193 = v185 + v192
	v195 = v184 + v192
	v197 = v186 - v192
	if base.Ui32(int32(3)) < base.Ui32(v197) {
		v184 = v195
		v185 = v193
		v186 = v197
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v207 = v200
	v208 = v201
	v209 = v202
	goto L52
L62:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
	if v217 == v218 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v238 = v217 - v218
	goto L50
L64:
	;
	v220 = int32(1)
	v225 = v214 - v220
	if v225 != 0 {
		v212 = v212 + v220
		v213 = v213 + v220
		v214 = v225
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
	v241 = int32(0)
	if v238 != 0 {
		v744 = v241
		goto L6
	} else {
		goto L69
	}
L69:
	;
	v243 = F_index_open(m, l0, int32(1))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	if v98 <= int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	F_relation_close(m, v243, int32(0))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L172
	}
L72:
	;
	v390 = int32(0)
	v391 = m.G0
	v393 = v391 - int32(32)
	m.G0 = v393
	if v377|v63 == v390 {
		goto L111
	} else {
		goto L112
	}
L73:
	;
	v248 = F_palloc_mul(m, int32(8), v98)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v255 = v241
	goto L77
L76:
	;
	v377 = v248
	goto L72
L77:
	;
	v268 = v255 << (uint(int32(2)) % 32)
	v269 = v60 + v268
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v271 = F_get_opclass_input_type(m, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L81
	}
L78:
	;
	v344 = F_palloc_mul(m, int32(8), v98)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L105
	}
L79:
	;
	v340 = v255 + int32(1)
	if v340 != v98 {
		v255 = v340
		goto L77
	} else {
		goto L104
	}
L80:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v243)+52))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v325+v326<<(uint(int32(3))%32)+v255*int32(100))+96))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v268+v54)))
	if v333 == v335 {
		goto L79
	} else {
		goto L103
	}
L81:
	;
	if v271 == int32(2283) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v276 = F_get_opclass_input_type(m, v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	if v276 == int32(2277) {
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v281 = F_get_opclass_input_type(m, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v281 == int32(2776) {
		goto L80
	} else {
		goto L86
	}
L86:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v286 = F_get_opclass_input_type(m, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	if v286 == int32(3500) {
		goto L80
	} else {
		goto L88
	}
L88:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v291 = F_get_opclass_input_type(m, v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	if v291 == int32(3831) {
		goto L80
	} else {
		goto L90
	}
L90:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v296 = F_get_opclass_input_type(m, v295)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	if v296 == int32(_a_F_CheckIndexCompatible_0) {
		goto L80
	} else {
		goto L92
	}
L92:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v301 = F_get_opclass_input_type(m, v300)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	if v301 == int32(_a_F_CheckIndexCompatible_1) {
		goto L80
	} else {
		goto L94
	}
L94:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v306 = F_get_opclass_input_type(m, v305)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	if v306 == int32(_a_F_CheckIndexCompatible_2) {
		goto L80
	} else {
		goto L96
	}
L96:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v311 = F_get_opclass_input_type(m, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	if v311 == int32(_a_F_CheckIndexCompatible_3) {
		goto L80
	} else {
		goto L98
	}
L98:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v316 = F_get_opclass_input_type(m, v315)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	if v316 == int32(_a_F_CheckIndexCompatible_4) {
		goto L80
	} else {
		goto L100
	}
L100:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v321 = F_get_opclass_input_type(m, v320)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	if v321 != int32(_a_F_CheckIndexCompatible_5) {
		goto L79
	} else {
		goto L102
	}
L102:
	;
	goto L80
L103:
	;
	v688 = int32(0)
	goto L71
L104:
	;
	goto L78
L105:
	;
	v351 = int32(0)
	goto L106
L106:
	;
	v367 = v351 + int32(1)
	v369 = F_get_attoptions(m, l0, base.I32_extend16_s(v367))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L108
	}
L107:
	;
	v377 = v344
	goto L72
L108:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v344+v351<<(uint(int32(3))%32)))) = v369
	if v98 != v367 {
		v351 = v367
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	m.G0 = v393 + int32(32)
	F_pfree(m, v377)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L134
	}
L111:
	;
	v472 = int32(1)
	goto L110
L112:
	;
	goto L113
L113:
	;
	F_fmgr_info(m, int32(744), v393+int32(4))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	if v98 <= int32(0) {
		v472 = int32(1)
		goto L110
	} else {
		goto L115
	}
L115:
	;
	v408 = v390
	goto L116
L116:
	;
	if v377 != 0 {
		goto L122
	} else {
		goto L123
	}
L117:
	;
	v472 = v465
	goto L110
L118:
	;
	v465 = int32(1)
	v467 = v408 + v465
	if v467 != v98 {
		v408 = v467
		goto L116
	} else {
		goto L133
	}
L119:
	;
	v472 = int32(0)
	goto L110
L120:
	;
	if v427 == int64(0) {
		goto L118
	} else {
		goto L132
	}
L121:
	;
	if v453 == int64(0) {
		goto L118
	} else {
		goto L131
	}
L122:
	;
	v425 = v408 << (uint(int32(3)) % 32)
	v427 = *(*int64)(unsafe.Add(mBase, uint32(v377+v425)))
	if v63 == int32(0) {
		goto L120
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	if v63 == int32(0) {
		goto L118
	} else {
		goto L130
	}
L125:
	;
	v431 = *(*int64)(unsafe.Add(mBase, uint32(v425+v63)))
	if v427 == int64(0) {
		v453 = v431
		goto L121
	} else {
		goto L126
	}
L126:
	;
	v434 = int32(0)
	if v431 == int64(0) {
		v472 = v434
		goto L110
	} else {
		goto L127
	}
L127:
	;
	v440 = F_FunctionCall2Coll(m, v393+int32(4), int32(950), v427, v431)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	if v440 == int64(0) {
		v472 = v434
		goto L110
	} else {
		goto L129
	}
L129:
	;
	goto L118
L130:
	;
	v449 = *(*int64)(unsafe.Add(mBase, uint32(v63+v408<<(uint(int32(3))%32))))
	v453 = v449
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
	if v472 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v688 = int32(0)
	goto L71
L136:
	;
	goto L137
L137:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v51)+92))
	if v494 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v688 = int32(1)
	goto L71
L139:
	;
	goto L140
L140:
	;
	F_RelationGetExclusionInfo(m, v243, v20+int32(44), v20+int32(40), v20+int32(36))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v51)+92))
	if base.Ui32(int32(4)) <= base.Ui32(v111) {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	v570 = int32(0)
	if v569|base.B2i32(v98 <= v570) != 0 {
		v688 = base.B2i32(v569 == v570)
		goto L71
	} else {
		goto L160
	}
L143:
	;
	v569 = int32(0)
	goto L142
L144:
	;
	v543 = v538
	v544 = v539
	v545 = v540
	goto L154
L145:
	;
	if (v506|v507)&int32(3) != 0 {
		v538 = v506
		v539 = v507
		v540 = v111
		goto L144
	} else {
		goto L148
	}
L146:
	;
	v531 = v506
	v532 = v507
	v533 = v111
	goto L147
L147:
	;
	if v533 == int32(0) {
		goto L143
	} else {
		goto L153
	}
L148:
	;
	v515 = v506
	v516 = v507
	v517 = v111
	goto L149
L149:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v516)))
	if v520 != v521 {
		v538 = v515
		v539 = v516
		v540 = v517
		goto L144
	} else {
		goto L151
	}
L150:
	;
	v531 = v526
	v532 = v524
	v533 = v528
	goto L147
L151:
	;
	v523 = int32(4)
	v524 = v516 + v523
	v526 = v515 + v523
	v528 = v517 - v523
	if base.Ui32(int32(3)) < base.Ui32(v528) {
		v515 = v526
		v516 = v524
		v517 = v528
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	v538 = v531
	v539 = v532
	v540 = v533
	goto L144
L154:
	;
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543))))
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544))))
	if v548 == v549 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v569 = v548 - v549
	goto L142
L156:
	;
	v551 = int32(1)
	v556 = v545 - v551
	if v556 != 0 {
		v543 = v543 + v551
		v544 = v544 + v551
		v545 = v556
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
	v576 = int32(0)
	goto L161
L161:
	;
	v594 = v576 << (uint(int32(2)) % 32)
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v51)+92))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v594+v595)))
	F_op_input_types(m, v597, v20+int32(32), v20+int32(28))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L163
	}
L162:
	;
	v688 = v679
	goto L71
L163:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	if base.B2i32(base.Ui32(v604-int32(_a_F_CheckIndexCompatible_1)) < base.Ui32(int32(2)))|base.B2i32(v604 == int32(_a_F_CheckIndexCompatible_0))|(base.B2i32(v604 == int32(3831))|base.B2i32(v604 == int32(3500)))|(base.B2i32(v604 == int32(2776))|base.B2i32(v604 == int32(2283))|(base.B2i32(v604 == int32(2277))|base.B2i32(v604 == int32(_a_F_CheckIndexCompatible_5)))) != 0 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v679 = int32(1)
	v681 = v576 + v679
	if v681 != v98 {
		v576 = v681
		goto L161
	} else {
		goto L171
	}
L165:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v243)+52))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v665)))
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v665+v666<<(uint(int32(3))%32)+v576*int32(100))+96))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v594+v54)))
	if v673 == v675 {
		goto L164
	} else {
		goto L170
	}
L166:
	;
	if base.Ui32(v604-int32(_a_F_CheckIndexCompatible_3)) < base.Ui32(int32(2)) {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	if base.B2i32(base.Ui32(v634-int32(_a_F_CheckIndexCompatible_1)) < base.Ui32(int32(2)))|base.B2i32(v634 == int32(_a_F_CheckIndexCompatible_0))|(base.B2i32(v634 == int32(3831))|base.B2i32(v634 == int32(3500)))|(base.B2i32(v634 == int32(2776))|base.B2i32(v634 == int32(2283))|(base.B2i32(v634 == int32(2277))|base.B2i32(v634 == int32(_a_F_CheckIndexCompatible_5)))) != 0 {
		goto L165
	} else {
		goto L168
	}
L168:
	;
	if base.Ui32(int32(1)) < base.Ui32(v634-int32(_a_F_CheckIndexCompatible_3)) {
		goto L164
	} else {
		goto L169
	}
L169:
	;
	goto L165
L170:
	;
	v688 = int32(0)
	goto L71
L171:
	;
	goto L162
L172:
	;
	v744 = v688
	goto L6
L173:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l1
	F_errmsg(m, int32(_a_F_CheckIndexCompatible_6), v20)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	F_errfinish(m, int32(_a_F_CheckIndexCompatible_7), int32(228), int32(_a_F_CheckIndexCompatible_8))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_CheckIndexCompatible_9), v20+int32(16))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(_a_F_CheckIndexCompatible_7), int32(264), int32(_a_F_CheckIndexCompatible_8))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
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
	v744 = int32(0)
	goto L6
}
func F_ExecIndexBuildScanKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v39 int32
	_ = v39
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
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
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
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int64
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int64
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
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
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int64
	_ = v321
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
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
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
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
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int64
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int64
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v445 int64
	_ = v445
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
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
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v520 int64
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int64
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v617 int32
	_ = v617
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v880 int32
	_ = v880
	var v885 int32
	_ = v885
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	v11 = int32(0)
	v39 = m.G0
	v41 = v39 - int32(48)
	m.G0 = v41
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v45 = v43
	goto L3
L2:
	;
	v45 = int32(0)
	goto L3
L3:
	;
	v48 = F_palloc(m, v45*int32(56))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v54 = F_palloc0(m, v45*int32(24))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
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
	v977 = m.ExcPending
	if v977 != 0 {
		goto L4
	} else {
		goto L199
	}
L8:
	;
	m.G0 = v41 + int32(48)
	return
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v918
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v917
	goto L8
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v48
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v45
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v821
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v820
	if l8 == int32(0) {
		goto L7
	} else {
		goto L198
	}
L11:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v56 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v855 = v50
	v859 = v51
	goto L13
L13:
	;
	F_pfree(m, v54)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L4
	} else {
		goto L196
	}
L14:
	;
	if v827 != 0 {
		goto L10
	} else {
		goto L195
	}
L15:
	;
	v820 = v50
	v821 = v51
	v827 = v11
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
	v61 = int32(256)
	goto L20
L19:
	;
	v61 = int32(0)
	goto L20
L20:
	;
	v77 = v48
	v78 = v50
	v81 = v50
	v82 = v51
	v88 = v11
	v90 = v11
	goto L30
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L4
	} else {
		goto L192
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L4
	} else {
		goto L189
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L4
	} else {
		goto L186
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L4
	} else {
		goto L183
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L4
	} else {
		goto L180
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L4
	} else {
		goto L177
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L4
	} else {
		goto L174
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L4
	} else {
		goto L171
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L4
	} else {
		goto L168
	}
L30:
	;
	v104 = v48 + v90*int32(56)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v106 = int32(*(*int16)(unsafe.Add(mBase, uint32(v105)+10)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+v90<<(uint(int32(2))%32))))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	switch v112 - int32(17) {
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
	v666 = m.ExcPending
	if v666 != 0 {
		goto L4
	} else {
		goto L165
	}
L32:
	;
	goto L31
L33:
	;
	v657 = v90 + int32(1)
	v661 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v657 < v661 {
		v77 = v48 + v657*int32(56)
		v78 = v632
		v81 = v635
		v82 = v636
		v88 = v642
		v90 = v657
		goto L30
	} else {
		goto L164
	}
L34:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v577)))
	if v578 == int32(27) {
		goto L152
	} else {
		goto L153
	}
L35:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)+12))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	if v468 == int32(27) {
		goto L120
	} else {
		goto L121
	}
L36:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	if v210 != 0 {
		goto L62
	} else {
		goto L63
	}
L37:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if v118 == int32(27) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v123 = v121
	v124 = v122
	goto L40
L39:
	;
	v123 = v117
	v124 = v118
	goto L40
L40:
	;
	if v124 != int32(6) {
		goto L32
	} else {
		goto L41
	}
L41:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v127 != int32(-3) {
		goto L32
	} else {
		goto L42
	}
L42:
	;
	v130 = int32(*(*int16)(unsafe.Add(mBase, uint32(v123)+8)))
	if base.B2i32(v130 <= int32(0))|base.B2i32(v106 < v130) != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v137+v130<<(uint(int32(2))%32)-int32(4))))
	F_get_op_opfamily_properties(m, v136, v143, l3, v41+int32(44), v41+int32(40), v41+int32(36))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v155 == int32(27) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+44)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v41)+36))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v111)+24))
	F_ScanKeyEntryInitialize(m, v104, v203, v130, v204, v205, v206, v135, v202)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L61
	}
L46:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v160 = v158
	v161 = v159
	goto L48
L47:
	;
	v160 = v154
	v161 = v155
	goto L48
L48:
	;
	if v161 == int32(7) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v164 = *(*int64)(unsafe.Add(mBase, uint32(v160)+24))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+32)))
	v199 = v78
	v200 = v81
	v201 = v82
	v202 = v164
	v203 = v61 | v165
	goto L45
L50:
	;
	goto L51
L51:
	;
	if v81 < v78 {
		v180 = v78
		v181 = v82
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v184 = v181 + v81*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = v104
	v186 = F_ExecInitExpr(m, v160, l0)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
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
	v172 = F_palloc(m, int32(96))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v176 = F_repalloc(m, v82, v78*int32(24))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L58
	}
L57:
	;
	v180 = int32(8)
	v181 = v172
	goto L52
L58:
	;
	v180 = v78 << (uint(int32(1)) % 32)
	v181 = v176
	goto L52
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v184)+4)) = v186
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v41)+36))
	v190 = F_get_typstorage(m, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v184)+8)) = uint8(base.B2i32(v190 != int32(112)))
	v199 = v180
	v200 = v81 + int32(1)
	v201 = v181
	v202 = int64(0)
	v203 = v61
	goto L45
L61:
	;
	v632 = v199
	v635 = v200
	v636 = v201
	v642 = v88
	goto L33
L62:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v215 = v211 * int32(56)
	goto L64
L63:
	;
	v215 = int32(0)
	goto L64
L64:
	;
	v216 = F_palloc(m, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v111)+24))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v232 = int32(0)
	v236 = v78
	v239 = v81
	v240 = v82
	goto L67
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = int32(4)
	v459 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v104)+4)) = uint16(v459)
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v104)+48)) = base.I64_extend_i32_u(v216)
	*(*uint16)(unsafe.Add(mBase, uint32(v104)+6)) = uint16(v461)
	v632 = v236
	v635 = v239
	v636 = v240
	v642 = v88
	goto L33
L67:
	;
	v260 = int32(0)
	if v221 == v260 {
		v271 = v260
		goto L69
	} else {
		goto L70
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+52)) = int32(0)
	v445 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v104)+44)) = v445
	*(*int64)(unsafe.Add(mBase, uint32(v104)+36)) = v445
	*(*int64)(unsafe.Add(mBase, uint32(v104)+28)) = v445
	*(*int64)(unsafe.Add(mBase, uint32(v104)+20)) = v445
	*(*int64)(unsafe.Add(mBase, uint32(v104)+12)) = v445
	*(*int64)(unsafe.Add(mBase, uint32(v104)+4)) = v445
	goto L66
L69:
	;
	if v220 == int32(0) {
		v280 = v260
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if v265 <= v232 {
		v271 = int32(0)
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v271 = v267 + v232<<(uint(int32(2))%32)
	goto L69
L72:
	;
	v281 = int32(0)
	if v219 == v281 {
		v292 = v281
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	if v274 <= v232 {
		v280 = v260
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	v280 = v276 + v232<<(uint(int32(2))%32)
	goto L72
L75:
	;
	if v218 == int32(0) {
		v301 = v281
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	if v286 <= v232 {
		v292 = int32(0)
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
	v292 = v288 + v232<<(uint(int32(2))%32)
	goto L75
L78:
	;
	v304 = v216 + v232*int32(56)
	v305 = int32(0)
	if v301 != 0 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	if v295 <= v232 {
		v301 = v281
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v218)+12))
	v301 = v297 + v232<<(uint(int32(2))%32)
	goto L78
L81:
	;
	goto L68
L82:
	;
	v314 = base.B2i32(v292 == v305) | (base.B2i32(v271 == v305) | base.B2i32(v280 == v305))
	goto L84
L83:
	;
	v314 = int32(1)
	goto L84
L84:
	;
	if v314 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v316 = v304 - int32(56)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	*(*int32)(unsafe.Add(mBase, uint32(v316))) = v317 | int32(16)
	if v48&int32(3) != 0 {
		goto L81
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)))
	if v336 == int32(27) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v321 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v77)+48)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v77)+40)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v77)+32)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v77)+24)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v77)+16)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v77)+8)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v321
	goto L66
L89:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v335)+4))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	v341 = v339
	v342 = v340
	goto L91
L90:
	;
	v341 = v335
	v342 = v336
	goto L91
L91:
	;
	if v342 != int32(6) {
		goto L28
	} else {
		goto L92
	}
L92:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	if v345 != int32(-3) {
		goto L28
	} else {
		goto L93
	}
L93:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+10)))
	if v349 != int32(1) {
		goto L27
	} else {
		goto L94
	}
L94:
	;
	v352 = int32(*(*int16)(unsafe.Add(mBase, uint32(v341)+8)))
	if base.B2i32(v352 <= int32(0))|base.B2i32(v106 < v352) != 0 {
		goto L27
	} else {
		goto L95
	}
L95:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v360+v352<<(uint(int32(2))%32)-int32(4))))
	F_get_op_opfamily_properties(m, v359, v366, l3, v41+int32(44), v41+int32(40), v41+int32(36))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v41)+44))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v375 != v376 {
		goto L26
	} else {
		goto L97
	}
L97:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v41)+40))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v41)+36))
	v381 = F_get_opfamily_proc(m, v366, v378, v379, int32(1))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	if v381 == int32(0) {
		goto L25
	} else {
		goto L99
	}
L99:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	if v385 == int32(27) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v437 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+44)))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v41)+36))
	F_ScanKeyEntryInitialize(m, v304, v436, v352, v437, v438, v357, v381, v435)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L4
	} else {
		goto L119
	}
L101:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	v390 = v388
	v391 = v389
	goto L103
L102:
	;
	v390 = v358
	v391 = v385
	goto L103
L103:
	;
	if v391 == int32(7) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v394 = *(*int64)(unsafe.Add(mBase, uint32(v390)+24))
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390)+32)))
	if v397 != 0 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	if v239 < v236 {
		v412 = v236
		v413 = v240
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v398 = int32(9)
	goto L109
L108:
	;
	v398 = int32(8)
	goto L109
L109:
	;
	v431 = v236
	v432 = v239
	v433 = v240
	v435 = v394
	v436 = v398
	goto L100
L110:
	;
	v416 = v413 + v239*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v416))) = v304
	v418 = F_ExecInitExpr(m, v390, l0)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L4
	} else {
		goto L117
	}
L111:
	;
	if v236 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v404 = F_palloc(m, int32(96))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L4
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v408 = F_repalloc(m, v240, v236*int32(24))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L4
	} else {
		goto L116
	}
L115:
	;
	v412 = int32(8)
	v413 = v404
	goto L110
L116:
	;
	v412 = v236 << (uint(int32(1)) % 32)
	v413 = v408
	goto L110
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416)+4)) = v418
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v41)+36))
	v422 = F_get_typstorage(m, v421)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v416)+8)) = uint8(base.B2i32(v422 != int32(112)))
	v431 = v412
	v432 = v239 + int32(1)
	v433 = v413
	v435 = int64(0)
	v436 = int32(8)
	goto L100
L119:
	;
	v232 = v232 + int32(1)
	v236 = v431
	v239 = v432
	v240 = v433
	goto L67
L120:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v467)+4))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)))
	v473 = v471
	v474 = v472
	goto L122
L121:
	;
	v473 = v467
	v474 = v468
	goto L122
L122:
	;
	if v474 != int32(6) {
		goto L24
	} else {
		goto L123
	}
L123:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v473)+4))
	if v477 != int32(-3) {
		goto L24
	} else {
		goto L124
	}
L124:
	;
	v480 = int32(*(*int16)(unsafe.Add(mBase, uint32(v473)+8)))
	if base.B2i32(v480 <= int32(0))|base.B2i32(v106 < v480) != 0 {
		goto L23
	} else {
		goto L125
	}
L125:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v487+v480<<(uint(int32(2))%32)-int32(4))))
	F_get_op_opfamily_properties(m, v486, v493, l3, v41+int32(44), v41+int32(40), v41+int32(36))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)+12))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v503)+4))
	if v504 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+19)))
	if v514 == int32(1) {
		goto L133
	} else {
		goto L134
	}
L128:
	;
	v512 = int32(0)
	goto L127
L129:
	;
	goto L130
L130:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v504)))
	if v508 != int32(27) {
		v512 = v504
		goto L127
	} else {
		goto L131
	}
L131:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v504)+4))
	v512 = v511
	goto L127
L132:
	;
	v572 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+44)))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v41)+36))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v111)+24))
	F_ScanKeyEntryInitialize(m, v104, v571, v480, v572, v573, v574, v485, v570)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L4
	} else {
		goto L151
	}
L133:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
	if v517 == int32(7) {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	goto L135
L135:
	;
	v555 = v54 + v88*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v555))) = v104
	v557 = F_ExecInitExpr(m, v512, l0)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L4
	} else {
		goto L150
	}
L136:
	;
	v520 = *(*int64)(unsafe.Add(mBase, uint32(v512)+24))
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+32)))
	if v523 != 0 {
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
		v538 = v78
		v539 = v82
		goto L142
	} else {
		goto L143
	}
L139:
	;
	v524 = int32(33)
	goto L141
L140:
	;
	v524 = int32(32)
	goto L141
L141:
	;
	v566 = v78
	v567 = v81
	v568 = v82
	v569 = v88
	v570 = v520
	v571 = v524
	goto L132
L142:
	;
	v542 = v539 + v81*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v542))) = v104
	v544 = F_ExecInitExpr(m, v512, l0)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
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
	v530 = F_palloc(m, int32(96))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L4
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v534 = F_repalloc(m, v82, v78*int32(24))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L4
	} else {
		goto L148
	}
L147:
	;
	v538 = int32(8)
	v539 = v530
	goto L142
L148:
	;
	v538 = v78 << (uint(int32(1)) % 32)
	v539 = v534
	goto L142
L149:
	;
	v546 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v542)+8)) = uint8(v546)
	*(*int32)(unsafe.Add(mBase, uint32(v542)+4)) = v544
	v566 = v538
	v567 = v81 + v546
	v568 = v539
	v569 = v88
	v570 = int64(0)
	v571 = int32(32)
	goto L132
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v555)+4)) = v557
	v566 = v78
	v567 = v81
	v568 = v82
	v569 = v88 + int32(1)
	v570 = int64(0)
	v571 = int32(0)
	goto L132
L151:
	;
	v632 = v566
	v635 = v567
	v636 = v568
	v642 = v569
	goto L33
L152:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v577)+4))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)))
	v583 = v581
	v584 = v582
	goto L154
L153:
	;
	v583 = v577
	v584 = v578
	goto L154
L154:
	;
	if v584 != int32(6) {
		goto L22
	} else {
		goto L155
	}
L155:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v583)+4))
	if v587 != int32(-3) {
		goto L22
	} else {
		goto L156
	}
L156:
	;
	v590 = int32(*(*int16)(unsafe.Add(mBase, uint32(v583)+8)))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	switch v592 {
	case 0:
		v610 = int32(65)
		goto L157
	case 1:
		goto L158
	default:
		goto L159
	}
L157:
	;
	v611 = int32(0)
	F_ScanKeyEntryInitialize(m, v104, v610, v590, v611, v611, v611, v611, int64(0))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L4
	} else {
		goto L163
	}
L158:
	;
	v610 = int32(129)
	goto L157
L159:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v597
	F_errmsg_internal(m, int32(_a_F_ExecIndexBuildScanKeys_0), v41+int32(32))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1616), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
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
	v632 = v78
	v635 = v81
	v636 = v82
	v642 = v88
	goto L33
L164:
	;
	v820 = v635
	v821 = v636
	v827 = v642
	goto L14
L165:
	;
	F_errmsg_internal(m, int32(_a_F_ExecIndexBuildScanKeys_3), int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1244), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
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
	v683 = m.ExcPending
	if v683 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1248), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
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
	v696 = m.ExcPending
	if v696 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1361), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
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
	v710 = m.ExcPending
	if v710 != 0 {
		goto L4
	} else {
		goto L175
	}
L175:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1371), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
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
	v723 = m.ExcPending
	if v723 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1380), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v366
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v41)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+20)) = v736
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v41)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = v738
	F_errmsg_internal(m, int32(_a_F_ExecIndexBuildScanKeys_7), v41+int32(16))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1388), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
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
	v757 = m.ExcPending
	if v757 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1485), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
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
	v770 = m.ExcPending
	if v770 != 0 {
		goto L4
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1489), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
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
	v783 = m.ExcPending
	if v783 != 0 {
		goto L4
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1599), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
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
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v793
	F_errmsg_internal(m, int32(_a_F_ExecIndexBuildScanKeys_9), v41)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L4
	} else {
		goto L193
	}
L193:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1632), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
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
	v855 = v820
	v859 = v821
	goto L13
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v48
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v45
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v859
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v855
	v885 = int32(0)
	if l8 != 0 {
		v917 = v885
		v918 = v885
		goto L9
	} else {
		goto L197
	}
L197:
	;
	goto L8
L198:
	;
	v917 = v827
	v918 = v54
	goto L9
L199:
	;
	F_errmsg_internal(m, int32(_a_F_ExecIndexBuildScanKeys_10), int32(0))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_ExecIndexBuildScanKeys_1), int32(1657), int32(_a_F_ExecIndexBuildScanKeys_2))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v42 int32
	_ = v42
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
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int64
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 float64
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int64
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v282 int32
	_ = v282
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
	var v294 int32
	_ = v294
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int64
	_ = v314
	var v316 int64
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
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
	var v353 int32
	_ = v353
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	var v403 int64
	_ = v403
	var v405 int64
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
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
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v483 int64
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v492 int64
	_ = v492
	var v493 int32
	_ = v493
	var v494 int64
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v539 int32
	_ = v539
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v20 != 0 {
		v48 = v20
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v59 = int32(0)
	goto L14
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v30 = F_ScanRelIsReadOnly(m, l0)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v30 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v34 = int32(1024)
	goto L7
L6:
	;
	v34 = int32(0)
	goto L7
L7:
	;
	v35 = F_index_beginscan(m, v21, v22, v24, v25, v26, v27, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v35
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v38 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)))
	if v39 != int32(1) {
		v48 = v35
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	F_index_rescan(m, v35, v42, v43, v44, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L3
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	v48 = v35
	goto L1
L14:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[0]))
	if v63 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	m.G0 = v16 + int32(16)
	return v18
L16:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	if v67 != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L18
L20:
	;
	goto L15
L21:
	;
	v160 = F_index_getnext_slot(m, v48, int32(1), v18)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L3
	} else {
		goto L44
	}
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+176)))
	if v69 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+176)))
	if v139 != int32(1) {
		v156 = v59
		goto L21
	} else {
		goto L41
	}
L25:
	;
	v134 = F_reorderqueue_pop(m, l0)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L3
	} else {
		goto L39
	}
L26:
	;
	v70 = int32(0)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v71 <= v70 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v48)+80))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v48)+76))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	v80 = v70
	goto L28
L28:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v74))))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v76))))
	if v94 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if int32(0) < v118 {
		v156 = v68
		goto L21
	} else {
		goto L38
	}
L30:
	;
	goto L29
L31:
	;
	v118 = v92 ^ int32(1)
	goto L30
L32:
	;
	goto L33
L33:
	;
	if v92&int32(1) != 0 {
		goto L25
	} else {
		goto L34
	}
L34:
	;
	v102 = v80 << (uint(int32(3)) % 32)
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v77+v102)))
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v102+v75)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v110 = v107 + v80*int32(36)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	v112 = m.T0[v111].(func(*base.Module, int64, int64, int32) int32)(m, v104, v106, v110)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	if v112 != 0 {
		v118 = v112
		goto L30
	} else {
		goto L36
	}
L36:
	;
	v115 = v80 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v115 < v116 {
		v80 = v115
		goto L28
	} else {
		goto L37
	}
L37:
	;
	goto L25
L38:
	;
	goto L25
L39:
	;
	F_ExecForceStoreHeapTuple(m, v134, v18, int32(1))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	goto L20
L41:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
	m.T0[v143].(func(*base.Module, int32))(m, v18)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	goto L20
L43:
	;
	v539 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+176)) = uint8(v539)
	v59 = v156
	goto L14
L44:
	;
	if v160 == int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	goto L46
L46:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+72)))
	if v177 != int32(1) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+84)))
	if v218 == int32(1) {
		goto L70
	} else {
		goto L71
	}
L48:
	;
	goto L47
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v18
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v181 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	F_MemoryContextReset(m, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L3
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v187 = int32(_a_F_IndexNextWithReorder_0)
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[1]))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[1])) = v190
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v181)+24))
	v195 = m.T0[v194].(func(*base.Module, int32, int32, int32) int64)(m, v181, v19, v16+int32(15))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L3
	} else {
		goto L54
	}
L53:
	;
	goto L48
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[1])) = v188
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	F_MemoryContextReset(m, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	if v195 != int64(0) {
		goto L48
	} else {
		goto L56
	}
L56:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v204 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v205 = *(*float64)(unsafe.Add(mBase, uint32(v204)+432))
	*(*float64)(unsafe.Add(mBase, uint32(v204)+432)) = base.F64_add(v205, float64(1))
	goto L59
L58:
	;
	goto L59
L59:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[0]))
	if v210 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L3
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v214 = F_index_getnext_slot(m, v48, int32(1), v18)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L3
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	if v214 != 0 {
		goto L46
	} else {
		goto L65
	}
L65:
	;
	goto L43
L66:
	;
	v434 = int32(_a_F_IndexNextWithReorder_0)
	v435 = *(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[1]))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[1])) = v439
	v442 = F_palloc(m, int32(24))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L3
	} else {
		goto L110
	}
L67:
	;
	if v156 == int32(0) {
		goto L20
	} else {
		goto L97
	}
L68:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if v306 == int32(0) {
		v428 = v353
		v429 = v352
		goto L66
	} else {
		goto L96
	}
L69:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v364 = v351
	v365 = v350
	goto L67
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v18
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	F_MemoryContextReset(m, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L3
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v48)+76))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v48)+80))
	v364 = v349
	v365 = v348
	goto L67
L73:
	;
	v225 = int32(_a_F_IndexNextWithReorder_0)
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[1]))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[1])) = v228
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v230 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v282 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[1])) = v226
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v287 <= v282 {
		v364 = v285
		v365 = v286
		goto L67
	} else {
		goto L81
	}
L75:
	;
	v233 = int32(0)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	if v234 <= v233 {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v239 = v233
	goto L77
L77:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v230)+12))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250+v239<<(uint(int32(2))%32))))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v254)+24))
	v258 = m.T0[v257].(func(*base.Module, int32, int32, int32) int64)(m, v254, v19, v255+v239)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L3
	} else {
		goto L79
	}
L78:
	;
	goto L74
L79:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	*(*int64)(unsafe.Add(mBase, uint32(v260+v239<<(uint(int32(3))%32)))) = v258
	v266 = v239 + int32(1)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	if v266 < v267 {
		v239 = v266
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v48)+80))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v48)+76))
	v294 = v282
	goto L83
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L3
	} else {
		goto L93
	}
L83:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294+v290))))
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294+v285))))
	if v308 != 0 {
		goto L68
	} else {
		goto L85
	}
L84:
	;
	if v322 < int32(0) {
		goto L82
	} else {
		goto L92
	}
L85:
	;
	if v306&int32(1) != 0 {
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v312 = v294 << (uint(int32(3)) % 32)
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v286+v312)))
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v312+v291)))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v320 = v317 + v294*int32(36)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+16))
	v322 = m.T0[v321].(func(*base.Module, int64, int64, int32) int32)(m, v314, v316, v320)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L3
	} else {
		goto L87
	}
L87:
	;
	if v322 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v327 = v294 + int32(1)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v328 <= v327 {
		goto L69
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	goto L84
L91:
	;
	v294 = v327
	goto L83
L92:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v428 = v333
	v429 = v332
	goto L66
L93:
	;
	F_errmsg_internal(m, int32(_a_F_IndexNextWithReorder_1), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L3
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_IndexNextWithReorder_2), int32(317), int32(_a_F_IndexNextWithReorder_3))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L3
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	v364 = v353
	v365 = v352
	goto L67
L97:
	;
	v371 = int32(0)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v372 <= v371 {
		goto L20
	} else {
		goto L98
	}
L98:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v156)+20))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v156)+16))
	v379 = v371
	goto L99
L99:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379+v375))))
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379+v364))))
	if v393 == int32(1) {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	if v417 <= int32(0) {
		goto L20
	} else {
		goto L109
	}
L101:
	;
	goto L100
L102:
	;
	v417 = v391 ^ int32(1)
	goto L101
L103:
	;
	goto L104
L104:
	;
	if v391&int32(1) != 0 {
		goto L20
	} else {
		goto L105
	}
L105:
	;
	v401 = v379 << (uint(int32(3)) % 32)
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v365+v401)))
	v405 = *(*int64)(unsafe.Add(mBase, uint32(v401+v376)))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v409 = v406 + v379*int32(36)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v409)+16))
	v411 = m.T0[v410].(func(*base.Module, int64, int64, int32) int32)(m, v403, v405, v409)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L3
	} else {
		goto L106
	}
L106:
	;
	if v411 != 0 {
		v417 = v411
		goto L101
	} else {
		goto L107
	}
L107:
	;
	v414 = v379 + int32(1)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v414 < v415 {
		v379 = v414
		goto L99
	} else {
		goto L108
	}
L108:
	;
	goto L20
L109:
	;
	v428 = v364
	v429 = v365
	goto L66
L110:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v444)+44))
	v446 = m.T0[v445].(func(*base.Module, int32) int32)(m, v18)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L3
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+12)) = v446
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v436)+16))
	v451 = F_palloc_mul(m, int32(8), v450)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L3
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+16)) = v451
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v436)+16))
	v456 = F_palloc_mul(m, int32(1), v455)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L3
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+20)) = v456
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if int32(0) < v459 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v464 = int32(0)
	goto L117
L115:
	;
	goto L116
L116:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	F_pairingheap_add(m, v521, v442)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L3
	} else {
		goto L124
	}
L117:
	;
	v476 = v464 + v428
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476))))
	if v477 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	goto L116
L119:
	;
	v483 = *(*int64)(unsafe.Add(mBase, uint32(v429+v464<<(uint(int32(3))%32))))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484+v464))))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v491 = int32(*(*int16)(unsafe.Add(mBase, uint32(v487+v464<<(uint(int32(1))%32)))))
	v492 = F_datumCopy(m, v483, v486, v491)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L3
	} else {
		goto L122
	}
L120:
	;
	v494 = int64(0)
	goto L121
L121:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v442)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v495+v464<<(uint(int32(3))%32)))) = v494
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v442)+20))
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476))))
	*(*uint8)(unsafe.Add(mBase, uint32(v500+v464))) = uint8(v502)
	v505 = v464 + int32(1)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v505 < v506 {
		v464 = v505
		goto L117
	} else {
		goto L123
	}
L122:
	;
	v494 = v492
	goto L121
L123:
	;
	goto L118
L124:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IndexNextWithReorder[1])) = v435
	v59 = v156
	goto L14
}
func F_IndexSetParentIndex(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
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
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
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
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(176)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v17 = F_relation_open(m, int32(2611), int32(3))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v20 = v12 - int32(-64)
		v24 = base.I64_extend_i32_u(v14)
		F_ScanKeyInit(m, v20, int32(1), int32(3), int32(184), v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			v29 = int32(3)
			F_ScanKeyInit(m, v12+int32(120), v29, v29, int32(65), int64(1))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				v39 = F_systable_beginscan(m, v17, int32(2680), int32(1), int32(0), int32(2), v20)
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
									F_relation_close(m, v17, int32(3))
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
														v81 = F_SearchSysCacheLockedCopy1(m, int32(57), v24)
														mBase = m.M
														v82 = m.ExcPending
														if v82 != 0 {
															return
														} else {
															if v81 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v170 = m.ExcPending
																if v170 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v14
																	F_errmsg_internal(m, int32(_a_F_IndexSetParentIndex_0), v12)
																	mBase = m.M
																	v174 = m.ExcPending
																	if v174 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_IndexSetParentIndex_1), int32(_a_F_IndexSetParentIndex_2), int32(_a_F_IndexSetParentIndex_3))
																		mBase = m.M
																		v179 = m.ExcPending
																		if v179 != 0 {
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
																*(*uint16)(unsafe.Add(mBase, uint32(v12)+56)) = uint16(v85)
																v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v87
																v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
																v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
																*(*uint8)(unsafe.Add(mBase, uint32(v89+v90)+131)) = uint8(base.B2i32(l1 != int32(0)))
																v96 = v12 + int32(52)
																F_CatalogTupleUpdate(m, v78, v96, v81)
																mBase = m.M
																v98 = m.ExcPending
																if v98 != 0 {
																	return
																} else {
																	F_UnlockTuple(m, v78, v96, int32(7))
																	mBase = m.M
																	v101 = m.ExcPending
																	if v101 != 0 {
																		return
																	} else {
																		F_pfree(m, v81)
																		mBase = m.M
																		v103 = m.ExcPending
																		if v103 != 0 {
																			return
																		} else {
																			F_relation_close(m, v78, int32(3))
																			mBase = m.M
																			v106 = m.ExcPending
																			if v106 != 0 {
																				return
																			} else {
																				if v64 != 0 {
																					if l1 != 0 {
																						v107 = int32(0)
																						*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v107
																						*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v14
																						v110 = int32(1259)
																						*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v110
																						*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v107
																						*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = l1
																						*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v110
																						*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v110
																						v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																						v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
																						*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v107
																						*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v120
																						F_recordDependencyOn(m, v96, v12+int32(40), int32(80))
																						mBase = m.M
																						v128 = m.ExcPending
																						if v128 != 0 {
																							return
																						} else {
																							F_recordDependencyOn(m, v96, v12+int32(28), int32(83))
																							mBase = m.M
																							v133 = m.ExcPending
																							if v133 != 0 {
																								return
																							} else {
																								F_CommandCounterIncrement(m)
																								mBase = m.M
																								v146 = m.ExcPending
																								if v146 != 0 {
																									return
																								} else {
																									m.G0 = v12 + int32(176)
																									return
																								}
																							}
																						}
																					} else {
																						v134 = int32(1259)
																						v137 = F_deleteDependencyRecordsForClass(m, v134, v14, v134, int32(80))
																						mBase = m.M
																						v138 = m.ExcPending
																						if v138 != 0 {
																							return
																						} else {
																							v139 = int32(1259)
																							v142 = F_deleteDependencyRecordsForClass(m, v139, v14, v139, int32(83))
																							mBase = m.M
																							v143 = m.ExcPending
																							if v143 != 0 {
																								return
																							} else {
																								F_CommandCounterIncrement(m)
																								mBase = m.M
																								v146 = m.ExcPending
																								if v146 != 0 {
																									return
																								} else {
																									m.G0 = v12 + int32(176)
																									return
																								}
																							}
																						}
																					}
																				} else {
																					m.G0 = v12 + int32(176)
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
												v81 = F_SearchSysCacheLockedCopy1(m, int32(57), v24)
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return
												} else {
													if v81 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v170 = m.ExcPending
														if v170 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v12))) = v14
															F_errmsg_internal(m, int32(_a_F_IndexSetParentIndex_0), v12)
															mBase = m.M
															v174 = m.ExcPending
															if v174 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_IndexSetParentIndex_1), int32(_a_F_IndexSetParentIndex_2), int32(_a_F_IndexSetParentIndex_3))
																mBase = m.M
																v179 = m.ExcPending
																if v179 != 0 {
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
														*(*uint16)(unsafe.Add(mBase, uint32(v12)+56)) = uint16(v85)
														v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v87
														v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
														v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
														*(*uint8)(unsafe.Add(mBase, uint32(v89+v90)+131)) = uint8(base.B2i32(l1 != int32(0)))
														v96 = v12 + int32(52)
														F_CatalogTupleUpdate(m, v78, v96, v81)
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return
														} else {
															F_UnlockTuple(m, v78, v96, int32(7))
															mBase = m.M
															v101 = m.ExcPending
															if v101 != 0 {
																return
															} else {
																F_pfree(m, v81)
																mBase = m.M
																v103 = m.ExcPending
																if v103 != 0 {
																	return
																} else {
																	F_relation_close(m, v78, int32(3))
																	mBase = m.M
																	v106 = m.ExcPending
																	if v106 != 0 {
																		return
																	} else {
																		if v64 != 0 {
																			if l1 != 0 {
																				v107 = int32(0)
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v107
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v14
																				v110 = int32(1259)
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v110
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v107
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = l1
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v110
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v110
																				v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																				v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v107
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v120
																				F_recordDependencyOn(m, v96, v12+int32(40), int32(80))
																				mBase = m.M
																				v128 = m.ExcPending
																				if v128 != 0 {
																					return
																				} else {
																					F_recordDependencyOn(m, v96, v12+int32(28), int32(83))
																					mBase = m.M
																					v133 = m.ExcPending
																					if v133 != 0 {
																						return
																					} else {
																						F_CommandCounterIncrement(m)
																						mBase = m.M
																						v146 = m.ExcPending
																						if v146 != 0 {
																							return
																						} else {
																							m.G0 = v12 + int32(176)
																							return
																						}
																					}
																				}
																			} else {
																				v134 = int32(1259)
																				v137 = F_deleteDependencyRecordsForClass(m, v134, v14, v134, int32(80))
																				mBase = m.M
																				v138 = m.ExcPending
																				if v138 != 0 {
																					return
																				} else {
																					v139 = int32(1259)
																					v142 = F_deleteDependencyRecordsForClass(m, v139, v14, v139, int32(83))
																					mBase = m.M
																					v143 = m.ExcPending
																					if v143 != 0 {
																						return
																					} else {
																						F_CommandCounterIncrement(m)
																						mBase = m.M
																						v146 = m.ExcPending
																						if v146 != 0 {
																							return
																						} else {
																							m.G0 = v12 + int32(176)
																							return
																						}
																					}
																				}
																			}
																		} else {
																			m.G0 = v12 + int32(176)
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
								F_StoreSingleInheritance(m, v14, l1, v47)
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
										F_relation_close(m, v17, int32(3))
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
															v81 = F_SearchSysCacheLockedCopy1(m, int32(57), v24)
															mBase = m.M
															v82 = m.ExcPending
															if v82 != 0 {
																return
															} else {
																if v81 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v170 = m.ExcPending
																	if v170 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v12))) = v14
																		F_errmsg_internal(m, int32(_a_F_IndexSetParentIndex_0), v12)
																		mBase = m.M
																		v174 = m.ExcPending
																		if v174 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_IndexSetParentIndex_1), int32(_a_F_IndexSetParentIndex_2), int32(_a_F_IndexSetParentIndex_3))
																			mBase = m.M
																			v179 = m.ExcPending
																			if v179 != 0 {
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
																	*(*uint16)(unsafe.Add(mBase, uint32(v12)+56)) = uint16(v85)
																	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
																	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v87
																	v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
																	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v89+v90)+131)) = uint8(base.B2i32(l1 != int32(0)))
																	v96 = v12 + int32(52)
																	F_CatalogTupleUpdate(m, v78, v96, v81)
																	mBase = m.M
																	v98 = m.ExcPending
																	if v98 != 0 {
																		return
																	} else {
																		F_UnlockTuple(m, v78, v96, int32(7))
																		mBase = m.M
																		v101 = m.ExcPending
																		if v101 != 0 {
																			return
																		} else {
																			F_pfree(m, v81)
																			mBase = m.M
																			v103 = m.ExcPending
																			if v103 != 0 {
																				return
																			} else {
																				F_relation_close(m, v78, int32(3))
																				mBase = m.M
																				v106 = m.ExcPending
																				if v106 != 0 {
																					return
																				} else {
																					if v64 != 0 {
																						if l1 != 0 {
																							v107 = int32(0)
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v107
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v14
																							v110 = int32(1259)
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v110
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v107
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = l1
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v110
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v110
																							v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																							v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v107
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v120
																							F_recordDependencyOn(m, v96, v12+int32(40), int32(80))
																							mBase = m.M
																							v128 = m.ExcPending
																							if v128 != 0 {
																								return
																							} else {
																								F_recordDependencyOn(m, v96, v12+int32(28), int32(83))
																								mBase = m.M
																								v133 = m.ExcPending
																								if v133 != 0 {
																									return
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v146 = m.ExcPending
																									if v146 != 0 {
																										return
																									} else {
																										m.G0 = v12 + int32(176)
																										return
																									}
																								}
																							}
																						} else {
																							v134 = int32(1259)
																							v137 = F_deleteDependencyRecordsForClass(m, v134, v14, v134, int32(80))
																							mBase = m.M
																							v138 = m.ExcPending
																							if v138 != 0 {
																								return
																							} else {
																								v139 = int32(1259)
																								v142 = F_deleteDependencyRecordsForClass(m, v139, v14, v139, int32(83))
																								mBase = m.M
																								v143 = m.ExcPending
																								if v143 != 0 {
																									return
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v146 = m.ExcPending
																									if v146 != 0 {
																										return
																									} else {
																										m.G0 = v12 + int32(176)
																										return
																									}
																								}
																							}
																						}
																					} else {
																						m.G0 = v12 + int32(176)
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
													v81 = F_SearchSysCacheLockedCopy1(m, int32(57), v24)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return
													} else {
														if v81 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v170 = m.ExcPending
															if v170 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v12))) = v14
																F_errmsg_internal(m, int32(_a_F_IndexSetParentIndex_0), v12)
																mBase = m.M
																v174 = m.ExcPending
																if v174 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_IndexSetParentIndex_1), int32(_a_F_IndexSetParentIndex_2), int32(_a_F_IndexSetParentIndex_3))
																	mBase = m.M
																	v179 = m.ExcPending
																	if v179 != 0 {
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
															*(*uint16)(unsafe.Add(mBase, uint32(v12)+56)) = uint16(v85)
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v87
															v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
															v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
															*(*uint8)(unsafe.Add(mBase, uint32(v89+v90)+131)) = uint8(base.B2i32(l1 != int32(0)))
															v96 = v12 + int32(52)
															F_CatalogTupleUpdate(m, v78, v96, v81)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return
															} else {
																F_UnlockTuple(m, v78, v96, int32(7))
																mBase = m.M
																v101 = m.ExcPending
																if v101 != 0 {
																	return
																} else {
																	F_pfree(m, v81)
																	mBase = m.M
																	v103 = m.ExcPending
																	if v103 != 0 {
																		return
																	} else {
																		F_relation_close(m, v78, int32(3))
																		mBase = m.M
																		v106 = m.ExcPending
																		if v106 != 0 {
																			return
																		} else {
																			if v64 != 0 {
																				if l1 != 0 {
																					v107 = int32(0)
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v107
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v14
																					v110 = int32(1259)
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v110
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v107
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = l1
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v110
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v110
																					v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																					v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v107
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v120
																					F_recordDependencyOn(m, v96, v12+int32(40), int32(80))
																					mBase = m.M
																					v128 = m.ExcPending
																					if v128 != 0 {
																						return
																					} else {
																						F_recordDependencyOn(m, v96, v12+int32(28), int32(83))
																						mBase = m.M
																						v133 = m.ExcPending
																						if v133 != 0 {
																							return
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v146 = m.ExcPending
																							if v146 != 0 {
																								return
																							} else {
																								m.G0 = v12 + int32(176)
																								return
																							}
																						}
																					}
																				} else {
																					v134 = int32(1259)
																					v137 = F_deleteDependencyRecordsForClass(m, v134, v14, v134, int32(80))
																					mBase = m.M
																					v138 = m.ExcPending
																					if v138 != 0 {
																						return
																					} else {
																						v139 = int32(1259)
																						v142 = F_deleteDependencyRecordsForClass(m, v139, v14, v139, int32(83))
																						mBase = m.M
																						v143 = m.ExcPending
																						if v143 != 0 {
																							return
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v146 = m.ExcPending
																							if v146 != 0 {
																								return
																							} else {
																								m.G0 = v12 + int32(176)
																								return
																							}
																						}
																					}
																				}
																			} else {
																				m.G0 = v12 + int32(176)
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
								F_simple_heap_delete(m, v17, v41+int32(4))
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
										F_relation_close(m, v17, int32(3))
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
															v81 = F_SearchSysCacheLockedCopy1(m, int32(57), v24)
															mBase = m.M
															v82 = m.ExcPending
															if v82 != 0 {
																return
															} else {
																if v81 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v170 = m.ExcPending
																	if v170 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v12))) = v14
																		F_errmsg_internal(m, int32(_a_F_IndexSetParentIndex_0), v12)
																		mBase = m.M
																		v174 = m.ExcPending
																		if v174 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_IndexSetParentIndex_1), int32(_a_F_IndexSetParentIndex_2), int32(_a_F_IndexSetParentIndex_3))
																			mBase = m.M
																			v179 = m.ExcPending
																			if v179 != 0 {
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
																	*(*uint16)(unsafe.Add(mBase, uint32(v12)+56)) = uint16(v85)
																	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
																	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v87
																	v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
																	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v89+v90)+131)) = uint8(base.B2i32(l1 != int32(0)))
																	v96 = v12 + int32(52)
																	F_CatalogTupleUpdate(m, v78, v96, v81)
																	mBase = m.M
																	v98 = m.ExcPending
																	if v98 != 0 {
																		return
																	} else {
																		F_UnlockTuple(m, v78, v96, int32(7))
																		mBase = m.M
																		v101 = m.ExcPending
																		if v101 != 0 {
																			return
																		} else {
																			F_pfree(m, v81)
																			mBase = m.M
																			v103 = m.ExcPending
																			if v103 != 0 {
																				return
																			} else {
																				F_relation_close(m, v78, int32(3))
																				mBase = m.M
																				v106 = m.ExcPending
																				if v106 != 0 {
																					return
																				} else {
																					if v64 != 0 {
																						if l1 != 0 {
																							v107 = int32(0)
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v107
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v14
																							v110 = int32(1259)
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v110
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v107
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = l1
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v110
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v110
																							v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																							v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v107
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v120
																							F_recordDependencyOn(m, v96, v12+int32(40), int32(80))
																							mBase = m.M
																							v128 = m.ExcPending
																							if v128 != 0 {
																								return
																							} else {
																								F_recordDependencyOn(m, v96, v12+int32(28), int32(83))
																								mBase = m.M
																								v133 = m.ExcPending
																								if v133 != 0 {
																									return
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v146 = m.ExcPending
																									if v146 != 0 {
																										return
																									} else {
																										m.G0 = v12 + int32(176)
																										return
																									}
																								}
																							}
																						} else {
																							v134 = int32(1259)
																							v137 = F_deleteDependencyRecordsForClass(m, v134, v14, v134, int32(80))
																							mBase = m.M
																							v138 = m.ExcPending
																							if v138 != 0 {
																								return
																							} else {
																								v139 = int32(1259)
																								v142 = F_deleteDependencyRecordsForClass(m, v139, v14, v139, int32(83))
																								mBase = m.M
																								v143 = m.ExcPending
																								if v143 != 0 {
																									return
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v146 = m.ExcPending
																									if v146 != 0 {
																										return
																									} else {
																										m.G0 = v12 + int32(176)
																										return
																									}
																								}
																							}
																						}
																					} else {
																						m.G0 = v12 + int32(176)
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
													v81 = F_SearchSysCacheLockedCopy1(m, int32(57), v24)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return
													} else {
														if v81 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v170 = m.ExcPending
															if v170 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v12))) = v14
																F_errmsg_internal(m, int32(_a_F_IndexSetParentIndex_0), v12)
																mBase = m.M
																v174 = m.ExcPending
																if v174 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_IndexSetParentIndex_1), int32(_a_F_IndexSetParentIndex_2), int32(_a_F_IndexSetParentIndex_3))
																	mBase = m.M
																	v179 = m.ExcPending
																	if v179 != 0 {
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
															*(*uint16)(unsafe.Add(mBase, uint32(v12)+56)) = uint16(v85)
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v87
															v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
															v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
															*(*uint8)(unsafe.Add(mBase, uint32(v89+v90)+131)) = uint8(base.B2i32(l1 != int32(0)))
															v96 = v12 + int32(52)
															F_CatalogTupleUpdate(m, v78, v96, v81)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return
															} else {
																F_UnlockTuple(m, v78, v96, int32(7))
																mBase = m.M
																v101 = m.ExcPending
																if v101 != 0 {
																	return
																} else {
																	F_pfree(m, v81)
																	mBase = m.M
																	v103 = m.ExcPending
																	if v103 != 0 {
																		return
																	} else {
																		F_relation_close(m, v78, int32(3))
																		mBase = m.M
																		v106 = m.ExcPending
																		if v106 != 0 {
																			return
																		} else {
																			if v64 != 0 {
																				if l1 != 0 {
																					v107 = int32(0)
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v107
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v14
																					v110 = int32(1259)
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v110
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v107
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = l1
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v110
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v110
																					v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																					v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v107
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v120
																					F_recordDependencyOn(m, v96, v12+int32(40), int32(80))
																					mBase = m.M
																					v128 = m.ExcPending
																					if v128 != 0 {
																						return
																					} else {
																						F_recordDependencyOn(m, v96, v12+int32(28), int32(83))
																						mBase = m.M
																						v133 = m.ExcPending
																						if v133 != 0 {
																							return
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v146 = m.ExcPending
																							if v146 != 0 {
																								return
																							} else {
																								m.G0 = v12 + int32(176)
																								return
																							}
																						}
																					}
																				} else {
																					v134 = int32(1259)
																					v137 = F_deleteDependencyRecordsForClass(m, v134, v14, v134, int32(80))
																					mBase = m.M
																					v138 = m.ExcPending
																					if v138 != 0 {
																						return
																					} else {
																						v139 = int32(1259)
																						v142 = F_deleteDependencyRecordsForClass(m, v139, v14, v139, int32(83))
																						mBase = m.M
																						v143 = m.ExcPending
																						if v143 != 0 {
																							return
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v146 = m.ExcPending
																							if v146 != 0 {
																								return
																							} else {
																								m.G0 = v12 + int32(176)
																								return
																							}
																						}
																					}
																				}
																			} else {
																				m.G0 = v12 + int32(176)
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
									v154 = m.ExcPending
									if v154 != 0 {
										return
									} else {
										v155 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
										*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v155
										F_errmsg_internal(m, int32(_a_F_IndexSetParentIndex_4), v12+int32(16))
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_IndexSetParentIndex_1), int32(_a_F_IndexSetParentIndex_5), int32(_a_F_IndexSetParentIndex_6))
											mBase = m.M
											v166 = m.ExcPending
											if v166 != 0 {
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
										F_relation_close(m, v17, int32(3))
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
															v81 = F_SearchSysCacheLockedCopy1(m, int32(57), v24)
															mBase = m.M
															v82 = m.ExcPending
															if v82 != 0 {
																return
															} else {
																if v81 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v170 = m.ExcPending
																	if v170 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v12))) = v14
																		F_errmsg_internal(m, int32(_a_F_IndexSetParentIndex_0), v12)
																		mBase = m.M
																		v174 = m.ExcPending
																		if v174 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_IndexSetParentIndex_1), int32(_a_F_IndexSetParentIndex_2), int32(_a_F_IndexSetParentIndex_3))
																			mBase = m.M
																			v179 = m.ExcPending
																			if v179 != 0 {
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
																	*(*uint16)(unsafe.Add(mBase, uint32(v12)+56)) = uint16(v85)
																	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
																	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v87
																	v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
																	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v89+v90)+131)) = uint8(base.B2i32(l1 != int32(0)))
																	v96 = v12 + int32(52)
																	F_CatalogTupleUpdate(m, v78, v96, v81)
																	mBase = m.M
																	v98 = m.ExcPending
																	if v98 != 0 {
																		return
																	} else {
																		F_UnlockTuple(m, v78, v96, int32(7))
																		mBase = m.M
																		v101 = m.ExcPending
																		if v101 != 0 {
																			return
																		} else {
																			F_pfree(m, v81)
																			mBase = m.M
																			v103 = m.ExcPending
																			if v103 != 0 {
																				return
																			} else {
																				F_relation_close(m, v78, int32(3))
																				mBase = m.M
																				v106 = m.ExcPending
																				if v106 != 0 {
																					return
																				} else {
																					if v64 != 0 {
																						if l1 != 0 {
																							v107 = int32(0)
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v107
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v14
																							v110 = int32(1259)
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v110
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v107
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = l1
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v110
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v110
																							v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																							v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v107
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v120
																							F_recordDependencyOn(m, v96, v12+int32(40), int32(80))
																							mBase = m.M
																							v128 = m.ExcPending
																							if v128 != 0 {
																								return
																							} else {
																								F_recordDependencyOn(m, v96, v12+int32(28), int32(83))
																								mBase = m.M
																								v133 = m.ExcPending
																								if v133 != 0 {
																									return
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v146 = m.ExcPending
																									if v146 != 0 {
																										return
																									} else {
																										m.G0 = v12 + int32(176)
																										return
																									}
																								}
																							}
																						} else {
																							v134 = int32(1259)
																							v137 = F_deleteDependencyRecordsForClass(m, v134, v14, v134, int32(80))
																							mBase = m.M
																							v138 = m.ExcPending
																							if v138 != 0 {
																								return
																							} else {
																								v139 = int32(1259)
																								v142 = F_deleteDependencyRecordsForClass(m, v139, v14, v139, int32(83))
																								mBase = m.M
																								v143 = m.ExcPending
																								if v143 != 0 {
																									return
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v146 = m.ExcPending
																									if v146 != 0 {
																										return
																									} else {
																										m.G0 = v12 + int32(176)
																										return
																									}
																								}
																							}
																						}
																					} else {
																						m.G0 = v12 + int32(176)
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
													v81 = F_SearchSysCacheLockedCopy1(m, int32(57), v24)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return
													} else {
														if v81 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v170 = m.ExcPending
															if v170 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v12))) = v14
																F_errmsg_internal(m, int32(_a_F_IndexSetParentIndex_0), v12)
																mBase = m.M
																v174 = m.ExcPending
																if v174 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_IndexSetParentIndex_1), int32(_a_F_IndexSetParentIndex_2), int32(_a_F_IndexSetParentIndex_3))
																	mBase = m.M
																	v179 = m.ExcPending
																	if v179 != 0 {
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
															*(*uint16)(unsafe.Add(mBase, uint32(v12)+56)) = uint16(v85)
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v87
															v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
															v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
															*(*uint8)(unsafe.Add(mBase, uint32(v89+v90)+131)) = uint8(base.B2i32(l1 != int32(0)))
															v96 = v12 + int32(52)
															F_CatalogTupleUpdate(m, v78, v96, v81)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return
															} else {
																F_UnlockTuple(m, v78, v96, int32(7))
																mBase = m.M
																v101 = m.ExcPending
																if v101 != 0 {
																	return
																} else {
																	F_pfree(m, v81)
																	mBase = m.M
																	v103 = m.ExcPending
																	if v103 != 0 {
																		return
																	} else {
																		F_relation_close(m, v78, int32(3))
																		mBase = m.M
																		v106 = m.ExcPending
																		if v106 != 0 {
																			return
																		} else {
																			if v64 != 0 {
																				if l1 != 0 {
																					v107 = int32(0)
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v107
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v14
																					v110 = int32(1259)
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v110
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v107
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = l1
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v110
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v110
																					v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
																					v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v107
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v120
																					F_recordDependencyOn(m, v96, v12+int32(40), int32(80))
																					mBase = m.M
																					v128 = m.ExcPending
																					if v128 != 0 {
																						return
																					} else {
																						F_recordDependencyOn(m, v96, v12+int32(28), int32(83))
																						mBase = m.M
																						v133 = m.ExcPending
																						if v133 != 0 {
																							return
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v146 = m.ExcPending
																							if v146 != 0 {
																								return
																							} else {
																								m.G0 = v12 + int32(176)
																								return
																							}
																						}
																					}
																				} else {
																					v134 = int32(1259)
																					v137 = F_deleteDependencyRecordsForClass(m, v134, v14, v134, int32(80))
																					mBase = m.M
																					v138 = m.ExcPending
																					if v138 != 0 {
																						return
																					} else {
																						v139 = int32(1259)
																						v142 = F_deleteDependencyRecordsForClass(m, v139, v14, v139, int32(83))
																						mBase = m.M
																						v143 = m.ExcPending
																						if v143 != 0 {
																							return
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v146 = m.ExcPending
																							if v146 != 0 {
																								return
																							} else {
																								m.G0 = v12 + int32(176)
																								return
																							}
																						}
																					}
																				}
																			} else {
																				m.G0 = v12 + int32(176)
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
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
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
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
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
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
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
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int64
	_ = v505
	var v507 int64
	_ = v507
	var v509 int64
	_ = v509
	var v511 int64
	_ = v511
	var v513 int32
	_ = v513
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v629 int32
	_ = v629
	var v642 int32
	_ = v642
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v681 int32
	_ = v681
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v735 int32
	_ = v735
	var v743 int32
	_ = v743
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
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
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
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
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+111)))
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
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
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
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+108)))
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
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+107)))
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
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
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
		goto L199
	} else {
		goto L200
	}
L32:
	;
	v197 = int32(1)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	if v198 != 0 {
		v203 = v197
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v773 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v773
	v776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+34)))
	if v776&int32(1) == v773 {
		v1000 = v207
		v1001 = v773
		v1005 = v760
		v1006 = v761
		v1013 = v768
		goto L31
	} else {
		goto L162
	}
L34:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l2)+60))
	v207 = v203 & base.B2i32(v204 != int32(0))
	if v207 != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	goto L34
L36:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+232)))
	if v199 != 0 {
		v203 = v197
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v203 = base.B2i32(v200 != int32(0))
	goto L35
L38:
	;
	v209 = F_build_index_pathkeys(m, l0, l2, int32(1))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L5
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+106)))
	if v213&v203&int32(1) == int32(0) {
		v735 = v8
		v743 = v8
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v211 = F_truncate_useless_pathkeys(m, l0, l1, v209)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v760 = v211
	v761 = v191
	v768 = v8
	goto L33
L43:
	;
	v760 = v735
	v761 = int32(0)
	v768 = v743
	goto L33
L44:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v219 == int32(0) {
		v760 = v8
		v761 = v191
		v768 = v8
		goto L33
	} else {
		goto L45
	}
L45:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	if int32(0) < v222 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v722 = F_list_copy_head(m, v709, v705)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L5
	} else {
		goto L161
	}
L47:
	;
	v694 = int32(0)
	if v673 == v694 {
		v760 = v681
		v761 = v694
		v768 = v689
		goto L33
	} else {
		goto L160
	}
L48:
	;
	v236 = v8
	v237 = v191
	v244 = v8
	goto L51
L49:
	;
	goto L50
L50:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v665 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L51:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v249+v236<<(uint(int32(2))%32))))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+12))
	if v254 != int32(1) {
		v642 = v237
		v649 = v244
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v654 != 0 {
		goto L151
	} else {
		goto L152
	}
L53:
	;
	goto L52
L54:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+16)))
	if v257 != 0 {
		v642 = v237
		v649 = v244
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+41)))
	if v259 != 0 {
		v642 = v237
		v649 = v244
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	v263 = v27 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v263))) = v258
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v258)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+4)) = int32(-1)
	if v265 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v269 = v261
	goto L59
L58:
	;
	v269 = int32(0)
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v263)+8)) = v269
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v258)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+16)) = v271
	if v271 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v271)+12))
	v275 = v273
	goto L62
L61:
	;
	v275 = int32(0)
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v263)+12)) = v275
	v278 = v27 + int32(20)
	v279 = int32(0)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v278)+16))
	if v283 == v279 {
		v332 = v279
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v332 == int32(0) {
		v642 = v237
		v649 = v244
		goto L53
	} else {
		goto L78
	}
L64:
	;
	goto L63
L65:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	if v286 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v313)))
	v319 = v313 + int32(4)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v314)+4))
	if base.Ui32(v319) < base.Ui32(v315+v321<<(uint(int32(2))%32)) {
		goto L75
	} else {
		goto L76
	}
L67:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v283)+12))
	v313 = v286
	v314 = v283
	v315 = v287
	goto L66
L68:
	;
	goto L69
L69:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	v290 = v288
	goto L70
L70:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v278)+8))
	v295 = F_bms_next_member(m, v294, v290)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v278)+4)) = v295
	if v295 <= int32(0) {
		v332 = v279
		goto L64
	} else {
		goto L72
	}
L71:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v306)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v278)+12)) = v310
	v313 = v310
	v314 = v306
	v315 = v310
	goto L66
L72:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+12))
	if v300 <= v295 {
		v332 = v279
		goto L64
	} else {
		goto L73
	}
L73:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v299)+20))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v302+v295<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v278)+16)) = v306
	if v306 == int32(0) {
		v290 = v295
		goto L70
	} else {
		goto L74
	}
L74:
	;
	goto L71
L75:
	;
	v326 = v319
	goto L77
L76:
	;
	v326 = int32(0)
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278)+12)) = v326
	v332 = v317
	goto L64
L78:
	;
	v357 = v332
	goto L79
L79:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v357)+8))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v360)+8))
	v362 = int32(0)
	if base.B2i32(v359 == v362)|base.B2i32(v361 == v362) != 0 {
		v408 = base.B2i32(v359|v361 == v362)
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v642 = v237
	v649 = v244
	goto L53
L81:
	;
	v575 = v27 + int32(20)
	v576 = int32(0)
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v575)+16))
	if v580 == v576 {
		v629 = v576
		goto L135
	} else {
		goto L136
	}
L82:
	;
	if v408 == int32(0) {
		goto L81
	} else {
		goto L93
	}
L83:
	;
	goto L82
L84:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v361)+4))
	if v376 != v377 {
		v408 = int32(0)
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v379 = int32(1)
	if v376 <= v379 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v382 = v379
	goto L88
L87:
	;
	v382 = v376
	goto L88
L88:
	;
	v383 = int32(8)
	v388 = int32(0)
	goto L89
L89:
	;
	v396 = v388 << (uint(int32(2)) % 32)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v359+v383+v396)))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v361+v383+v396)))
	v401 = base.B2i32(v398 == v400)
	if v398 != v400 {
		v408 = v401
		goto L83
	} else {
		goto L91
	}
L90:
	;
	v408 = v401
	goto L83
L91:
	;
	v404 = v388 + int32(1)
	if v404 != v382 {
		v388 = v404
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v415 <= int32(0) {
		goto L81
	} else {
		goto L94
	}
L94:
	;
	v422 = int32(0)
	goto L95
L95:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v357)+4))
	if v443 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L81
L97:
	;
	v547 = v422 + int32(1)
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v547 < v548 {
		v422 = v547
		goto L95
	} else {
		goto L133
	}
L98:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v443)))
	if v446 != int32(17) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v443)+28))
	if v449 == int32(0) {
		goto L97
	} else {
		goto L100
	}
L100:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v452 < int32(2) {
		goto L97
	} else {
		goto L101
	}
L101:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v449)+12))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	if v456 == int32(0) {
		goto L97
	} else {
		goto L102
	}
L102:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v455)+4))
	if v459 == int32(0) {
		goto L97
	} else {
		goto L103
	}
L103:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v253)+8))
	v464 = v422 << (uint(int32(2)) % 32)
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v464+v465)))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v469+v464)))
	if v471 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v443)+24))
	if v471 != v472 {
		goto L97
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v474 = F_match_index_to_operand(m, v456, v422, l2)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L5
	} else {
		goto L110
	}
L107:
	;
	goto L106
L108:
	;
	v532 = F_lappend(m, v237, v529)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L5
	} else {
		goto L130
	}
L109:
	;
	v485 = F_match_index_to_operand(m, v459, v422, l2)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L5
	} else {
		goto L118
	}
L110:
	;
	if v474 == int32(0) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v478 = F_contain_var_clause(m, v459)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	if v478 != 0 {
		goto L109
	} else {
		goto L113
	}
L113:
	;
	v480 = F_contain_volatile_functions(m, v459)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L5
	} else {
		goto L114
	}
L114:
	;
	if v480 != 0 {
		goto L109
	} else {
		goto L115
	}
L115:
	;
	v482 = F_get_op_opfamily_sortfamily(m, v468, v467)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L5
	} else {
		goto L116
	}
L116:
	;
	if v482 == v462 {
		v529 = v443
		goto L108
	} else {
		goto L117
	}
L117:
	;
	goto L97
L118:
	;
	if v485 == int32(0) {
		goto L97
	} else {
		goto L119
	}
L119:
	;
	v489 = F_contain_var_clause(m, v456)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L5
	} else {
		goto L120
	}
L120:
	;
	if v489 != 0 {
		goto L97
	} else {
		goto L121
	}
L121:
	;
	v491 = F_contain_volatile_functions(m, v456)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	if v491 != 0 {
		goto L97
	} else {
		goto L123
	}
L123:
	;
	v493 = F_get_commutator(m, v468)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	if v493 == int32(0) {
		goto L97
	} else {
		goto L125
	}
L125:
	;
	v497 = F_get_op_opfamily_sortfamily(m, v493, v467)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L5
	} else {
		goto L126
	}
L126:
	;
	if v497 != v462 {
		goto L97
	} else {
		goto L127
	}
L127:
	;
	v501 = F_palloc0(m, int32(36))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v501))) = int32(17)
	v505 = *(*int64)(unsafe.Add(mBase, uint32(v443)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v501)+8)) = v505
	v507 = *(*int64)(unsafe.Add(mBase, uint32(v443)))
	*(*int64)(unsafe.Add(mBase, uint32(v501))) = v507
	v509 = *(*int64)(unsafe.Add(mBase, uint32(v443)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v501)+16)) = v509
	v511 = *(*int64)(unsafe.Add(mBase, uint32(v443)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v501)+24)) = v511
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v443)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v501)+32)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v501)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v501)+4)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v456
	*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v456
	v526 = F_list_make2_impl(m, v27+int32(16), v27+int32(12))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L5
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v501)+28)) = v526
	v529 = v501
	goto L108
L130:
	;
	v534 = F_lappend_int(m, v244, v422)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L5
	} else {
		goto L131
	}
L131:
	;
	v537 = v236 + int32(1)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	if v537 < v538 {
		v236 = v537
		v237 = v532
		v244 = v534
		goto L51
	} else {
		goto L132
	}
L132:
	;
	v642 = v532
	v649 = v534
	goto L53
L133:
	;
	goto L96
L134:
	;
	if v629 != 0 {
		v357 = v629
		goto L79
	} else {
		goto L149
	}
L135:
	;
	goto L134
L136:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v575)+12))
	if v583 != 0 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v610)))
	v616 = v610 + int32(4)
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v611)+4))
	if base.Ui32(v616) < base.Ui32(v612+v618<<(uint(int32(2))%32)) {
		goto L146
	} else {
		goto L147
	}
L138:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v580)+12))
	v610 = v583
	v611 = v580
	v612 = v584
	goto L137
L139:
	;
	goto L140
L140:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v575)+4))
	v587 = v585
	goto L141
L141:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v575)+8))
	v592 = F_bms_next_member(m, v591, v587)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v575)+4)) = v592
	if v592 <= int32(0) {
		v629 = v576
		goto L135
	} else {
		goto L143
	}
L142:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v575)+12)) = v607
	v610 = v607
	v611 = v603
	v612 = v607
	goto L137
L143:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v596)+12))
	if v597 <= v592 {
		v629 = v576
		goto L135
	} else {
		goto L144
	}
L144:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v596)+20))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v599+v592<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v575)+16)) = v603
	if v603 == int32(0) {
		v587 = v592
		goto L141
	} else {
		goto L145
	}
L145:
	;
	goto L142
L146:
	;
	v623 = v616
	goto L148
L147:
	;
	v623 = int32(0)
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v575)+12)) = v623
	v629 = v614
	goto L135
L149:
	;
	goto L80
L150:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v642)+4))
	if v660 == v662 {
		v760 = v661
		v761 = v642
		v768 = v649
		goto L33
	} else {
		goto L156
	}
L151:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v654)+4))
	if v642 != 0 {
		v660 = v655
		v661 = v654
		goto L150
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v656 = int32(0)
	if v642 == v656 {
		v735 = v656
		v743 = v649
		goto L43
	} else {
		goto L155
	}
L154:
	;
	v673 = v655
	v681 = v654
	v689 = v649
	goto L47
L155:
	;
	v660 = v656
	v661 = v656
	goto L150
L156:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v642)+4))
	v705 = v664
	v709 = v661
	v710 = v642
	v717 = v649
	goto L46
L157:
	;
	v760 = int32(0)
	v761 = v191
	v768 = v8
	goto L33
L158:
	;
	goto L159
L159:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v665)+4))
	v673 = v669
	v681 = v665
	v689 = v8
	goto L47
L160:
	;
	v705 = v694
	v709 = v681
	v710 = v694
	v717 = v689
	goto L46
L161:
	;
	v760 = v722
	v761 = v710
	v768 = v717
	goto L33
L162:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v781)+4))
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	F_pull_varattnos(m, v782, v783, v27+int32(20))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L5
	} else {
		goto L163
	}
L163:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(l2)+96))
	if v788 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	if v858 <= int32(0) {
		goto L172
	} else {
		goto L173
	}
L165:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v788)+4))
	if v791 <= int32(0) {
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v798 = int32(0)
	goto L167
L167:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v788)+12))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v819+v798<<(uint(int32(2))%32))))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v823)+4))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	F_pull_varattnos(m, v824, v825, v27+int32(20))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L5
	} else {
		goto L169
	}
L168:
	;
	goto L164
L169:
	;
	v831 = v798 + int32(1)
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v788)+4))
	if v831 < v832 {
		v798 = v831
		goto L167
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v935 = int32(0)
	if v934 == v935 {
		goto L183
	} else {
		goto L184
	}
L172:
	;
	v919 = int32(0)
	goto L171
L173:
	;
	goto L174
L174:
	;
	v862 = int32(0)
	v867 = v862
	v871 = v858
	v873 = v862
	goto L175
L175:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v888+v867<<(uint(int32(2))%32))))
	if v892 == int32(0) {
		v905 = v871
		v906 = v873
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v919 = v906
	goto L171
L177:
	;
	v908 = v867 + int32(1)
	if v908 < v905 {
		v867 = v908
		v871 = v905
		v873 = v906
		goto L175
	} else {
		goto L181
	}
L178:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
	v897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v895+v867))))
	if v897 != int32(1) {
		v905 = v871
		v906 = v873
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v902 = F_bms_add_member(m, v873, v892+int32(7))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L5
	} else {
		goto L180
	}
L180:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	v905 = v904
	v906 = v902
	goto L177
L181:
	;
	goto L176
L182:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	F_bms_free(m, v989)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L5
	} else {
		goto L196
	}
L183:
	;
	v988 = int32(1)
	goto L182
L184:
	;
	goto L185
L185:
	;
	if v919 == int32(0) {
		v981 = v935
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v988 = v981
	goto L182
L187:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v934)+4))
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v919)+4))
	if v945 < v944 {
		v981 = v935
		goto L186
	} else {
		goto L188
	}
L188:
	;
	v947 = int32(1)
	if v944 <= v947 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v950 = v947
	goto L191
L190:
	;
	v950 = v944
	goto L191
L191:
	;
	v951 = int32(8)
	v956 = int32(0)
	goto L192
L192:
	;
	v963 = v956 << (uint(int32(2)) % 32)
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v934+v951+v963)))
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v919+v951+v963)))
	v970 = v965 & (v967 ^ int32(-1))
	v972 = base.B2i32(v970 == int32(0))
	if v970 != 0 {
		v981 = v972
		goto L186
	} else {
		goto L194
	}
L193:
	;
	v981 = v972
	goto L186
L194:
	;
	v974 = v956 + int32(1)
	if v974 != v950 {
		v956 = v974
		goto L192
	} else {
		goto L195
	}
L195:
	;
	goto L193
L196:
	;
	F_bms_free(m, v919)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L5
	} else {
		goto L197
	}
L197:
	;
	v1000 = v207
	v1001 = v988
	v1005 = v760
	v1006 = v761
	v1013 = v768
	goto L31
L198:
	;
	if v1000 == int32(0) {
		v1105 = v1057
		goto L1
	} else {
		goto L213
	}
L199:
	;
	v1022 = int32(0)
	if v1001 == v1022 {
		v1057 = v1022
		goto L198
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v1026 = int32(0)
	v1029 = F_create_index_path(m, l0, l2, v174, v1006, v1013, v1005, int32(1), v1001, v186, v189, v1026)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L5
	} else {
		goto L203
	}
L202:
	;
	goto L201
L203:
	;
	v1031 = F_lappend(m, v1026, v1029)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L5
	} else {
		goto L204
	}
L204:
	;
	v1033 = int32(1)
	v1035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+112)))
	if v186|(base.B2i32(l5 == v1033)|base.B2i32(v1035 != v1033)) != 0 {
		v1057 = v1031
		goto L198
	} else {
		goto L205
	}
L205:
	;
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v1040&int32(1) == int32(0) {
		v1057 = v1031
		goto L198
	} else {
		goto L206
	}
L206:
	;
	v1045 = int32(1)
	v1048 = F_create_index_path(m, l0, l2, v174, v1006, v1013, v1005, v1045, v1001, int32(0), v189, v1045)
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L5
	} else {
		goto L207
	}
L207:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+24))
	if int32(0) < v1050 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	F_add_partial_path(m, l1, v1048)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L5
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	F_pfree(m, v1048)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L5
	} else {
		goto L212
	}
L211:
	;
	v1057 = v1031
	goto L198
L212:
	;
	v1057 = v1031
	goto L198
L213:
	;
	v1062 = F_build_index_pathkeys(m, l0, l2, int32(-1))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L5
	} else {
		goto L214
	}
L214:
	;
	v1064 = F_truncate_useless_pathkeys(m, l0, l1, v1062)
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L5
	} else {
		goto L215
	}
L215:
	;
	if v1064 == int32(0) {
		v1105 = v1057
		goto L1
	} else {
		goto L216
	}
L216:
	;
	v1068 = int32(0)
	v1072 = F_create_index_path(m, l0, l2, v174, v1068, v1068, v1064, int32(-1), v1001, v186, v189, v1068)
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L5
	} else {
		goto L217
	}
L217:
	;
	v1074 = F_lappend(m, v1057, v1072)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L5
	} else {
		goto L218
	}
L218:
	;
	v1076 = int32(1)
	v1078 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+112)))
	if v186|(base.B2i32(l5 == v1076)|base.B2i32(v1078 != v1076)) != 0 {
		v1105 = v1074
		goto L1
	} else {
		goto L219
	}
L219:
	;
	v1083 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v1083&int32(1) == int32(0) {
		v1105 = v1074
		goto L1
	} else {
		goto L220
	}
L220:
	;
	v1088 = int32(0)
	v1093 = F_create_index_path(m, l0, l2, v174, v1088, v1088, v1064, int32(-1), v1001, v1088, v189, int32(1))
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L5
	} else {
		goto L221
	}
L221:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+24))
	if int32(0) < v1095 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	F_add_partial_path(m, l1, v1093)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L5
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	F_pfree(m, v1093)
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L5
	} else {
		goto L226
	}
L225:
	;
	v1105 = v1074
	goto L1
L226:
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
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v113 int32
	_ = v113
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
	return v113
L2:
	;
	return int32(0)
L3:
	;
	if v14 == int32(0) {
		v113 = v7
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(469)
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
	v45 = F_OidFunctionCall1Coll(m, v14, int32(0), base.I64_extend_i32_u(v12+int32(8)))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v47 = base.I32_wrap_i64(v45)
	if v47 == int32(0) {
		v113 = v7
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v51 = F_palloc0(m, int32(20))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(284)
	v55 = int32(0)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v55 < v56 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v62 = v55
	v65 = int32(0)
	goto L11
L9:
	;
	v92 = v55
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = l1
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+44)))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v51)+14)) = uint16(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v51)+12)) = uint8(v101)
	v113 = v51
	goto L1
L11:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+v65<<(uint(int32(2))%32))))
	v75 = int32(0)
	v82 = F_make_restrictinfo(m, l0, v73, int32(1), v75, v75, v75, v75, v75, v75, v75)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L13
	}
L12:
	;
	v92 = v84
	goto L10
L13:
	;
	v84 = F_lappend(m, v62, v82)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v87 = v65 + int32(1)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v87 < v88 {
		v62 = v84
		v65 = v87
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
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v7 = F_SearchSysCache1(m, int32(34), base.I64_extend_i32_u(l0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 == int32(0) {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15+v16)+10)))
			if l1 <= v18 {
				v22 = F_SysCacheGetAttrNotNull(m, int32(34), v7, int32(18))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v22)+l1<<(uint(int32(2))%32))+20))
					v29 = v28
					F_ReleaseCatCache(m, v7)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						return v29
					}
				}
			} else {
				v29 = int32(0)
				F_ReleaseCatCache(m, v7)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					return v29
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
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_SearchSysCache1(m, int32(34), base.I64_extend_i32_u(l0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_get_index_isvalid_0), v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_index_isvalid_1), int32(3920), int32(_a_F_get_index_isvalid_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v30)+18)))
			F_ReleaseCatCache(m, v10)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v32
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
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
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
					F_index_build(m, v12, v43, v45, int32(0), v47, v47)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						F_AtEOXact_GUC(m, int32(0), v38)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
							*(*int32)(unsafe.Add(mBase, _c_F_index_concurrently_build[1])) = v58
							*(*int32)(unsafe.Add(mBase, _c_F_index_concurrently_build[0])) = v57
							F_relation_close(m, v12, int32(0))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								F_relation_close(m, v43, int32(0))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									F_index_set_state_flags(m, l1, int32(0))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
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
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	v5 = int32(0)
	v14 = int64(0)
	v15 = m.G0
	v17 = v15 - int32(336)
	m.G0 = v17
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+334)) = uint16(v5)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	base.MemoryFill(m, v17-int32(-64), v5, int32(256))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+56)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v17)+48)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = v14
	if v21 <= int32(32) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v200 = v17 - int32(-64)
	v201 = F_heap_compute_data_size(m, l0, v200, l2)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L8
	} else {
		goto L38
	}
L2:
	;
	v188 = int32(8)
	v190 = v174
	v197 = int32(0)
	goto L1
L3:
	;
	v66 = v5
	goto L13
L4:
	;
	if int32(0) < v21 {
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
	v44 = m.ExcPending
	if v44 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v174 = v5
	goto L2
L8:
	;
	return int32(0)
L9:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v21
	F_errmsg(m, int32(_a_F_index_form_tuple_context_0), v17)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_index_form_tuple_context_1), int32(90), int32(_a_F_index_form_tuple_context_2))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
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
	v74 = v66 << (uint(int32(3)) % 32)
	v76 = *(*int64)(unsafe.Add(mBase, uint32(l1+v74)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v80 = v17 + int32(32) + v66
	v81 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v80))) = uint8(v81)
	v85 = v17 - int32(-64) + v74
	*(*int64)(unsafe.Add(mBase, uint32(v85))) = v76
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v66))))
	if v88 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v145 = int32(0)
	goto L32
L15:
	;
	v141 = v66 + int32(1)
	if v141 != v21 {
		v66 = v141
		goto L13
	} else {
		goto L31
	}
L16:
	;
	v96 = l0 + v77<<(uint(int32(3))%32) + v66*int32(100) + int32(28)
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96)+72)))
	if v97 != int32(_a_F_index_form_tuple_context_3) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v100 = base.I32_wrap_i64(v76)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	v103 = base.B2i32(v101 != int32(1))
	if v101 != int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v111 = v100
	v112 = v76
	v113 = v101
	goto L20
L19:
	;
	v104 = F_detoast_external_attr(m, v100)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L8
	} else {
		goto L21
	}
L20:
	;
	if v113&int32(3) != 0 {
		goto L15
	} else {
		goto L22
	}
L21:
	;
	v106 = base.I64_extend_i32_u(v104)
	*(*int64)(unsafe.Add(mBase, uint32(v85))) = v106
	v108 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v80))) = uint8(v108)
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	v111 = v104
	v112 = v106
	v113 = v110
	goto L20
L22:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if base.Ui32(v116) < base.Ui32(int32(2044)) {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+84)))
	switch v119 - int32(109) {
	case 0, 11:
		goto L24
	default:
		goto L15
	}
L24:
	;
	v122 = int32(*(*int8)(unsafe.Add(mBase, uint32(v96)+85)))
	v123 = F_toast_compress_datum(m, v112, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	if base.I32_wrap_i64(v123) == int32(0) {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	if v103 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_pfree(m, v111)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L8
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v85))) = v123
	v133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v80))) = uint8(v133)
	goto L15
L30:
	;
	goto L29
L31:
	;
	goto L14
L32:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+l2))))
	if v159 != int32(1) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v188 = int32(16)
	v190 = v159
	v197 = int32(_a_F_index_form_tuple_context_4)
	goto L1
L34:
	;
	v163 = v145 + int32(1)
	if v21 != v163 {
		v145 = v163
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
	v174 = v159
	goto L2
L38:
	;
	v205 = v201 + v188 + int32(7)
	v207 = v205 & int32(-8)
	v208 = F_MemoryContextAllocZero(m, l3, v207)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	if v190 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v216 = v208 + int32(8)
	goto L42
L41:
	;
	v216 = int32(0)
	goto L42
L42:
	;
	F_heap_fill_tuple(m, l0, v200, l2, v208+v188, v17+int32(334), v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	if int32(0) < v21 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v222 = int32(0)
	goto L47
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(int32(_a_F_index_form_tuple_context_5)) <= base.Ui32(v205) {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(32)+v222))))
	if v238 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L46
L49:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v17-int32(-64)+v222<<(uint(int32(3))%32))))
	F_pfree(m, v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L8
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v250 = v222 + int32(1)
	if v250 != v21 {
		v222 = v250
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
	v271 = m.ExcPending
	if v271 != 0 {
		goto L8
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+334)))
	v294 = v288<<(uint(int32(13))%32)&int32(_a_F_index_form_tuple_context_6) | (v207 | v197)
	*(*uint16)(unsafe.Add(mBase, uint32(v208)+6)) = uint16(v294)
	m.G0 = v17 + int32(336)
	return v208
L57:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L8
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(_a_F_index_form_tuple_context_7)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v207
	F_errmsg(m, int32(_a_F_index_form_tuple_context_8), v17+int32(16))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L8
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_index_form_tuple_context_1), int32(210), int32(_a_F_index_form_tuple_context_2))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
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
				F_errfinish(m, int32(_a_F_index_getnext_tid_2), int32(604), int32(_a_F_index_getnext_tid_3))
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
					F_errfinish(m, int32(_a_F_index_insert_2), int32(224), int32(_a_F_index_insert_3))
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
					F_errfinish(m, int32(_a_F_index_insert_2), int32(223), int32(_a_F_index_insert_3))
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
		*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(387)
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v92 int32
	_ = v92
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	if l1 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L7
	} else {
		goto L36
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L7
	} else {
		goto L33
	}
L3:
	;
	m.G0 = v13 + int32(32)
	return
L4:
	;
	v15 = F_get_index_isclustered(m, l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v19 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L7
	} else {
		goto L10
	}
L7:
	;
	return
L8:
	;
	if v15 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L6
L10:
	;
	v21 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L7
	} else {
		goto L12
	}
L11:
	;
	F_relation_close(m, v19, int32(3))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L7
	} else {
		goto L32
	}
L12:
	;
	if v21 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v25 <= int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v36 = int32(0)
	goto L15
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v36<<(uint(int32(2))%32))))
	v46 = F_SearchSysCacheCopy(m, int32(34), base.I64_extend_i32_u(v43), int64(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L7
	} else {
		goto L17
	}
L16:
	;
	goto L11
L17:
	;
	if v46 == int32(0) {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+22)))
	v52 = v50 + v51
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+17)))
	if v53 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_mark_index_clustered[0]))
	if v68 != 0 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	v60 = int32(0)
	goto L22
L21:
	;
	if l1 != v43 {
		goto L19
	} else {
		goto L23
	}
L22:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+17)) = uint8(v60)
	F_CatalogTupleUpdate(m, v19, v46+int32(4), v46)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L7
	} else {
		goto L25
	}
L23:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+18)))
	if v56 == int32(0) {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v60 = int32(1)
	goto L22
L25:
	;
	goto L19
L26:
	;
	v70 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2610), v43, v70, v70, l2)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L7
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_pfree(m, v46)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	v77 = v36 + int32(1)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v77 < v78 {
		v36 = v77
		goto L15
	} else {
		goto L31
	}
L31:
	;
	goto L16
L32:
	;
	goto L3
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v43
	F_errmsg_internal(m, int32(_a_F_mark_index_clustered_0), v13)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_mark_index_clustered_1), int32(857), int32(_a_F_mark_index_clustered_2))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_mark_index_clustered_3), v13+int32(16))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_mark_index_clustered_1), int32(873), int32(_a_F_mark_index_clustered_2))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
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
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int64
	_ = v108
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
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
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
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
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int64
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v176 int32
	_ = v176
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int64
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
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
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v304 int32
	_ = v304
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
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
	var v374 int32
	_ = v374
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
	var v397 int32
	_ = v397
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v642 int32
	_ = v642
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
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
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v692 int32
	_ = v692
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
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
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1016 int32
	_ = v1016
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1086 int32
	_ = v1086
	var v1091 int32
	_ = v1091
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1129 int32
	_ = v1129
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1153 int32
	_ = v1153
	var v1156 int32
	_ = v1156
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
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int64
	_ = v1201
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1279 int32
	_ = v1279
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1292 int32
	_ = v1292
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
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
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1344 int32
	_ = v1344
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1370 int32
	_ = v1370
	var v1374 int32
	_ = v1374
	var v1377 int32
	_ = v1377
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1387 int32
	_ = v1387
	var v1392 int32
	_ = v1392
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1413 int32
	_ = v1413
	var v1418 int32
	_ = v1418
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1435 int32
	_ = v1435
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1462 int32
	_ = v1462
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
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
	var v1486 int32
	_ = v1486
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1506 int32
	_ = v1506
	var v1529 int32
	_ = v1529
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1546 int32
	_ = v1546
	var v1551 int32
	_ = v1551
	var v1576 int32
	_ = v1576
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1590 int32
	_ = v1590
	var v1605 int32
	_ = v1605
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1614 int32
	_ = v1614
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1648 int32
	_ = v1648
	var v1651 int32
	_ = v1651
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1662 int32
	_ = v1662
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1675 int32
	_ = v1675
	var v1700 int32
	_ = v1700
	var v1705 int32
	_ = v1705
	var v1710 int32
	_ = v1710
	var v1715 int32
	_ = v1715
	var v1720 int32
	_ = v1720
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1733 int32
	_ = v1733
	var v1741 int32
	_ = v1741
	var v1759 int32
	_ = v1759
	var v1763 int32
	_ = v1763
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1791 int32
	_ = v1791
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1819 int32
	_ = v1819
	var v1822 int32
	_ = v1822
	var v1825 int32
	_ = v1825
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1836 int32
	_ = v1836
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1849 int32
	_ = v1849
	var v1875 int32
	_ = v1875
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1902 int32
	_ = v1902
	var v1906 int32
	_ = v1906
	var v1909 int32
	_ = v1909
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1919 int32
	_ = v1919
	var v1924 int32
	_ = v1924
	var v1927 int32
	_ = v1927
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1957 int64
	_ = v1957
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2007 int32
	_ = v2007
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2024 int32
	_ = v2024
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2035 int32
	_ = v2035
	var v2043 int32
	_ = v2043
	var v2053 int32
	_ = v2053
	var v2061 int32
	_ = v2061
	var v2065 int32
	_ = v2065
	var v2066 int32
	_ = v2066
	var v2070 int32
	_ = v2070
	var v2071 int32
	_ = v2071
	var v2075 int32
	_ = v2075
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2104 int32
	_ = v2104
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2125 int32
	_ = v2125
	var v2126 int32
	_ = v2126
	var v2129 int32
	_ = v2129
	var v2132 int32
	_ = v2132
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2143 int32
	_ = v2143
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2156 int32
	_ = v2156
	var v2157 int32
	_ = v2157
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2195 int32
	_ = v2195
	var v2196 int32
	_ = v2196
	var v2211 int32
	_ = v2211
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2237 int32
	_ = v2237
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2255 int32
	_ = v2255
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2265 int32
	_ = v2265
	var v2270 int32
	_ = v2270
	v2 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(320)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v27 == v2 {
		v2002 = l0
		v2007 = v25
		v2016 = v2
		v2017 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2255 = m.ExcPending
	if v2255 != 0 {
		goto L7
	} else {
		goto L478
	}
L2:
	;
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(v2002)+56))
	if v2024 != 0 {
		goto L437
	} else {
		goto L438
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v30 <= int32(0) {
		v2002 = l0
		v2007 = v25
		v2016 = v2
		v2017 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v33 = l0
	v38 = v25
	v47 = v2
	v48 = v2
	v49 = v27
	v52 = v2
	goto L5
L5:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v52<<(uint(int32(2))%32))))
	v61 = F_palloc0(m, int32(72))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v2002 = v33
	v2007 = v38
	v2016 = v47
	v2017 = v1996
	goto L2
L7:
	;
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = int32(204)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+60)) = uint8(base.B2i32(v65 != int32(8)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v71 = base.B2i32(v69 == int32(6))
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+62)) = uint8(v71)
	if v69 == int32(6) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v59)+40))
	if v1576 == int32(0) {
		goto L357
	} else {
		goto L358
	}
L10:
	;
	v1529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+36)))
	if v1529 == int32(0) {
		goto L9
	} else {
		goto L347
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L7
	} else {
		goto L342
	}
L12:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v33)+56))
	if v75 != 0 {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+30)))
	v78 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+63)) = uint8(v78)
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+61)) = uint8(v77)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+36)))
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+64)) = uint8(v81)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+65)) = uint8(v83)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+66)) = uint8(v85)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v87 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+56)) = v61
	goto L14
L16:
	;
	v88 = F_pstrdup(m, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L7
	} else {
		goto L19
	}
L17:
	;
	v91 = int32(0)
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v59)+64))
	if v95 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v91 = v88
	goto L18
L20:
	;
	v97 = v95
	goto L22
L21:
	;
	v97 = int32(_a_F_transformIndexConstraints_0)
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+12)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v59)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+28)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v59)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v59)+68))
	v104 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+69)) = uint8(v104)
	*(*uint16)(unsafe.Add(mBase, uint32(v61)+67)) = uint16(v104)
	v108 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v61)+20)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v61)+32)) = v103
	*(*int64)(unsafe.Add(mBase, uint32(v61)+36)) = v108
	*(*int64)(unsafe.Add(mBase, uint32(v61)+44)) = v108
	*(*int64)(unsafe.Add(mBase, uint32(v61)+52)) = v108
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+60)))
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+70)) = uint8(v117)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v59)+52))
	if v119 != 0 {
		goto L34
	} else {
		goto L35
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L7
	} else {
		goto L337
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L7
	} else {
		goto L332
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L7
	} else {
		goto L327
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L7
	} else {
		goto L322
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L7
	} else {
		goto L317
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L7
	} else {
		goto L311
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L7
	} else {
		goto L305
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L7
	} else {
		goto L299
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L7
	} else {
		goto L293
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L7
	} else {
		goto L288
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L7
	} else {
		goto L282
	}
L34:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+21)))
	if v120 == int32(0) {
		goto L23
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v328 == int32(8) {
		goto L89
	} else {
		goto L90
	}
L37:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+68))
	v126 = F_get_relname_relid(m, v119, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	if v126 == int32(0) {
		goto L24
	} else {
		goto L39
	}
L39:
	;
	v131 = F_index_open(m, v126, int32(1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)+192))
	v134 = F_get_index_constraint(m, v126)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	if v134 != 0 {
		goto L25
	} else {
		goto L42
	}
L42:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v123)+56))
	if v136 != v137 {
		goto L26
	} else {
		goto L43
	}
L43:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+18)))
	if v139 == int32(0) {
		goto L27
	} else {
		goto L44
	}
L44:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+12)))
	if v142 == int32(0) {
		goto L28
	} else {
		goto L45
	}
L45:
	;
	v145 = F_RelationGetIndexExpressions(m, v131)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	if v145 != 0 {
		goto L29
	} else {
		goto L47
	}
L47:
	;
	v147 = F_RelationGetIndexPredicate(m, v131)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	if v147 != 0 {
		goto L30
	} else {
		goto L49
	}
L49:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+16)))
	if v149 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+12)))
	if v152 == int32(0) {
		goto L31
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v131)+48))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+84))
	v158 = F_get_index_am_oid(m, int32(_a_F_transformIndexConstraints_0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L7
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	if v156 != v158 {
		goto L32
	} else {
		goto L55
	}
L55:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v131)+196))
	v164 = F_SysCacheGetAttrNotNull(m, int32(34), v162, int32(18))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	v166 = int32(*(*int16)(unsafe.Add(mBase, uint32(v133)+8)))
	if int32(0) < v166 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v176 = int32(0)
	goto L60
L58:
	;
	goto L59
L59:
	;
	F_relation_close(m, v131, int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L7
	} else {
		goto L88
	}
L60:
	;
	v198 = v176 << (uint(int32(1)) % 32)
	v200 = int32(*(*int16)(unsafe.Add(mBase, uint32(v133+int32(48)+v198))))
	if int32(0) < v200 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L59
L62:
	;
	v219 = F_pstrdup(m, v216+int32(4))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L7
	} else {
		goto L67
	}
L63:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v123)+52))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v216 = v203 + v204<<(uint(int32(3))%32) + v200*int32(100) - int32(72)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v213 = F_SystemAttributeDefinition(m, v200)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	v216 = v213
	goto L62
L67:
	;
	v221 = int32(*(*int16)(unsafe.Add(mBase, uint32(v133)+10)))
	if v176 < v221 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v278 = int32(*(*int16)(unsafe.Add(mBase, uint32(v133)+8)))
	if v275 < v278 {
		v176 = v275
		goto L60
	} else {
		goto L87
	}
L69:
	;
	v224 = v176 << (uint(int32(2)) % 32)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v131)+56))
	v227 = v176 + int32(1)
	v229 = F_get_attoptions(m, v225, base.I32_extend16_s(v227))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L7
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v59)+40))
	v268 = F_makeString(m, v219)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L7
	} else {
		goto L85
	}
L72:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v216)+68))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v131)+48))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+84))
	v234 = F_GetDefaultOpClass(m, v231, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L7
	} else {
		goto L73
	}
L73:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v224+(base.I32_wrap_i64(v164)+int32(24)))))
	if v234 != v237 {
		goto L33
	} else {
		goto L74
	}
L74:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v216)+96))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v131)+248))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v240+v224)))
	if base.B2i32(v239 != v242)|base.B2i32(v229 != int64(0)) != 0 {
		goto L33
	} else {
		goto L75
	}
L75:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v131)+224))
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247+v198))))
	if v249 != 0 {
		goto L33
	} else {
		goto L76
	}
L76:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v250 == int32(6) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
	v254 = F_makeString(m, v219)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L7
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v59)+32))
	v262 = F_makeString(m, v219)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L7
	} else {
		goto L83
	}
L80:
	;
	v256 = F_makeNotNullConstraint(m, v254)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L7
	} else {
		goto L81
	}
L81:
	;
	v258 = F_lappend(m, v253, v256)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L7
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v258
	goto L79
L83:
	;
	v264 = F_lappend(m, v261, v262)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L7
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+32)) = v264
	v275 = v227
	goto L68
L85:
	;
	v270 = F_lappend(m, v267, v268)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L7
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+40)) = v270
	v275 = v176 + int32(1)
	goto L68
L87:
	;
	goto L61
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+44)) = v126
	goto L36
L89:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v59)+44))
	if v331 == int32(0) {
		goto L9
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v59)+32))
	if v380 == int32(0) {
		goto L10
	} else {
		goto L99
	}
L92:
	;
	v334 = int32(0)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	if v335 <= v334 {
		goto L9
	} else {
		goto L93
	}
L93:
	;
	v339 = v334
	goto L94
L94:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v331)+12))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v360+v339<<(uint(int32(2))%32))))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)+12))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	v369 = F_lappend(m, v367, v368)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L7
	} else {
		goto L96
	}
L95:
	;
	goto L9
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+20)) = v369
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v61)+36))
	v373 = F_lappend(m, v372, v366)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L7
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+36)) = v373
	v377 = v339 + int32(1)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	if v377 < v378 {
		v339 = v377
		goto L94
	} else {
		goto L98
	}
L98:
	;
	goto L95
L99:
	;
	v383 = int32(0)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v380)+4))
	if v384 <= v383 {
		goto L10
	} else {
		goto L100
	}
L100:
	;
	v397 = v383
	goto L101
L101:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v380)+12))
	v412 = v409 + v397<<(uint(int32(2))%32)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)+4))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	if v415 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L102:
	;
	goto L10
L103:
	;
	v1195 = F_palloc0(m, int32(40))
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L7
	} else {
		goto L278
	}
L104:
	;
	v1129 = int32(0)
	if base.B2i32(v861 == v1129)|v1111 == v1129 {
		goto L261
	} else {
		goto L262
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v414
	F_errmsg(m, int32(_a_F_transformIndexConstraints_1), v38-int32(-64))
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L7
	} else {
		goto L258
	}
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L7
	} else {
		goto L254
	}
L107:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	if v876 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L108:
	;
	v844 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v448)+19)) = uint8(v844)
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
	v847 = F_makeString(m, v414)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L7
	} else {
		goto L211
	}
L109:
	;
	v598 = int32(0)
	v599 = int32(1)
	v601 = F_strcmp(m, int32(_a_F_transformIndexConstraints_2), v414)
	mBase = m.M
	if v601 == v598 {
		goto L149
	} else {
		goto L150
	}
L110:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v415)+4))
	if v418 <= int32(0) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v415)+12))
	v424 = int32(0)
	goto L112
L112:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v421+v424<<(uint(int32(2))%32))))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449))))
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414))))
	if base.B2i32(v452 == int32(0))|base.B2i32(v452 != v455) != 0 {
		v473 = v452
		v474 = v455
		goto L115
	} else {
		goto L116
	}
L113:
	;
	v479 = int32(0)
	v480 = int32(1)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v481 != int32(6) {
		v858 = v479
		v861 = v448
		v862 = v480
		goto L107
	} else {
		goto L125
	}
L114:
	;
	if v473-v474 != 0 {
		goto L121
	} else {
		goto L122
	}
L115:
	;
	goto L114
L116:
	;
	v458 = v449
	v459 = v414
	goto L117
L117:
	;
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459)+1)))
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458)+1)))
	if v463 == int32(0) {
		v473 = v463
		v474 = v462
		goto L115
	} else {
		goto L119
	}
L118:
	;
	v473 = v463
	v474 = v462
	goto L115
L119:
	;
	v466 = int32(1)
	if v463 == v462 {
		v458 = v458 + v466
		v459 = v459 + v466
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v477 = v424 + int32(1)
	if v477 != v418 {
		v424 = v477
		goto L112
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	goto L113
L124:
	;
	goto L109
L125:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+21)))
	if v484 != 0 {
		v858 = v479
		v861 = v448
		v862 = v480
		goto L107
	} else {
		goto L126
	}
L126:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+19)))
	if v485 != int32(1) {
		goto L108
	} else {
		goto L127
	}
L127:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
	if v488 == int32(0) {
		v858 = v479
		v861 = v448
		v862 = v480
		goto L107
	} else {
		goto L128
	}
L128:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	if v491 <= int32(0) {
		v858 = v479
		v861 = v448
		v862 = v480
		goto L107
	} else {
		goto L129
	}
L129:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v488)+12))
	v497 = int32(0)
	goto L130
L130:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v494+v497<<(uint(int32(2))%32))))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v521)+32))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v522)+12))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v524)+4))
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525))))
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414))))
	if base.B2i32(v528 == int32(0))|base.B2i32(v528 != v531) != 0 {
		v549 = v528
		v550 = v531
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521)+17)))
	if v555 != int32(1) {
		v858 = v479
		v861 = v448
		v862 = v480
		goto L107
	} else {
		goto L143
	}
L132:
	;
	if v549-v550 != 0 {
		goto L139
	} else {
		goto L140
	}
L133:
	;
	goto L132
L134:
	;
	v534 = v525
	v535 = v414
	goto L135
L135:
	;
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535)+1)))
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+1)))
	if v539 == int32(0) {
		v549 = v539
		v550 = v538
		goto L133
	} else {
		goto L137
	}
L136:
	;
	v549 = v539
	v550 = v538
	goto L133
L137:
	;
	v542 = int32(1)
	if v539 == v538 {
		v534 = v534 + v542
		v535 = v535 + v542
		goto L135
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	v553 = v497 + int32(1)
	if v553 != v491 {
		v497 = v553
		goto L130
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	goto L131
L142:
	;
	v858 = v479
	v861 = v448
	v862 = v480
	goto L107
L143:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L7
	} else {
		goto L144
	}
L144:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L7
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+128)) = v414
	F_errmsg(m, int32(_a_F_transformIndexConstraints_3), v38+int32(128))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L7
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2680), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L7
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	if v630 != 0 {
		goto L167
	} else {
		goto L168
	}
L149:
	;
	v630 = int32(_a_F_transformIndexConstraints_6)
	goto L148
L150:
	;
	goto L151
L151:
	;
	v606 = F_strcmp(m, int32(_a_F_transformIndexConstraints_7), v414)
	mBase = m.M
	if v606 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v630 = int32(_a_F_transformIndexConstraints_8)
	goto L148
L153:
	;
	goto L154
L154:
	;
	v611 = F_strcmp(m, int32(_a_F_transformIndexConstraints_9), v414)
	mBase = m.M
	if v611 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v630 = int32(_a_F_transformIndexConstraints_10)
	goto L148
L156:
	;
	goto L157
L157:
	;
	v616 = F_strcmp(m, int32(_a_F_transformIndexConstraints_11), v414)
	mBase = m.M
	if v616 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v630 = int32(_a_F_transformIndexConstraints_12)
	goto L148
L159:
	;
	goto L160
L160:
	;
	v621 = F_strcmp(m, int32(_a_F_transformIndexConstraints_13), v414)
	mBase = m.M
	if v621 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v630 = int32(_a_F_transformIndexConstraints_14)
	goto L148
L162:
	;
	goto L163
L163:
	;
	v628 = F_strcmp(m, int32(_a_F_transformIndexConstraints_15), v414)
	mBase = m.M
	if v628 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v629 = int32(0)
	goto L166
L165:
	;
	v629 = int32(_a_F_transformIndexConstraints_16)
	goto L166
L166:
	;
	v630 = v629
	goto L148
L167:
	;
	v858 = v598
	v861 = int32(0)
	v862 = v599
	goto L107
L168:
	;
	goto L169
L169:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v632 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+21)))
	if v818 != 0 {
		goto L203
	} else {
		goto L204
	}
L171:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v632)+4))
	if v635 <= int32(0) {
		goto L170
	} else {
		goto L172
	}
L172:
	;
	v642 = v598
	goto L173
L173:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v632)+12))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v660+v642<<(uint(int32(2))%32))))
	v666 = F_table_openrv(m, v664, int32(1))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L7
	} else {
		goto L175
	}
L174:
	;
	goto L170
L175:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v666)+48))
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668)+119)))
	v671 = v669 - int32(102)
	if base.B2i32(base.Ui32(int32(12)) < base.Ui32(v671))|base.B2i32(int32(1)<<(uint(v671)%32)&int32(_a_F_transformIndexConstraints_17) == int32(0)) != 0 {
		goto L106
	} else {
		goto L176
	}
L176:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v666)+52))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	if int32(0) < v682 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v692 = int32(0)
	goto L180
L178:
	;
	goto L179
L179:
	;
	F_relation_close(m, v666, int32(0))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L7
	} else {
		goto L201
	}
L180:
	;
	v715 = v681 + v682<<(uint(int32(3))%32) + int32(28) + v692*int32(100)
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v715)+91)))
	if v716 != 0 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	goto L179
L182:
	;
	v765 = v692 + int32(1)
	if v765 != v682 {
		v692 = v765
		goto L180
	} else {
		goto L200
	}
L183:
	;
	v718 = v715 + int32(4)
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414))))
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v718))))
	if base.B2i32(v721 == int32(0))|base.B2i32(v721 != v724) != 0 {
		v742 = v721
		v743 = v724
		goto L185
	} else {
		goto L186
	}
L184:
	;
	if v742-v743 != 0 {
		goto L182
	} else {
		goto L191
	}
L185:
	;
	goto L184
L186:
	;
	v727 = v414
	v728 = v718
	goto L187
L187:
	;
	v731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v728)+1)))
	v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v727)+1)))
	if v732 == int32(0) {
		v742 = v732
		v743 = v731
		goto L185
	} else {
		goto L189
	}
L188:
	;
	v742 = v732
	v743 = v731
	goto L185
L189:
	;
	v735 = int32(1)
	if v732 == v731 {
		v727 = v727 + v735
		v728 = v728 + v735
		goto L187
	} else {
		goto L190
	}
L190:
	;
	goto L188
L191:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v715)+68))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v746 == int32(6) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
	v750 = F_pstrdup(m, v718)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L7
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v759 = int32(0)
	F_relation_close(m, v666, v759)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L7
	} else {
		goto L199
	}
L195:
	;
	v752 = F_makeString(m, v750)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L7
	} else {
		goto L196
	}
L196:
	;
	v754 = F_makeNotNullConstraint(m, v752)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L7
	} else {
		goto L197
	}
L197:
	;
	v756 = F_lappend(m, v749, v754)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L7
	} else {
		goto L198
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v756
	goto L194
L199:
	;
	v858 = v745
	v861 = v759
	v862 = v599
	goto L107
L200:
	;
	goto L181
L201:
	;
	v793 = v642 + int32(1)
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v632)+4))
	if v793 < v794 {
		v642 = v793
		goto L173
	} else {
		goto L202
	}
L202:
	;
	goto L174
L203:
	;
	v819 = int32(0)
	v858 = v819
	v861 = v819
	v862 = v819
	goto L107
L204:
	;
	goto L205
L205:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L7
	} else {
		goto L206
	}
L206:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L7
	} else {
		goto L207
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v414
	F_errmsg(m, int32(_a_F_transformIndexConstraints_18), v38+int32(96))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L7
	} else {
		goto L208
	}
L208:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	F_parser_errposition(m, v835, v836)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L7
	} else {
		goto L209
	}
L209:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2760), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L7
	} else {
		goto L210
	}
L210:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L211:
	;
	v849 = F_makeNotNullConstraint(m, v847)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L7
	} else {
		goto L212
	}
L212:
	;
	v851 = F_lappend(m, v846, v849)
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L7
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v851
	v858 = v479
	v861 = v448
	v862 = v480
	goto L107
L214:
	;
	v989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+36)))
	if v989 == int32(0) {
		goto L103
	} else {
		goto L236
	}
L215:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v876)+4))
	if v879 <= int32(0) {
		goto L214
	} else {
		goto L216
	}
L216:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v876)+12))
	v885 = int32(0)
	goto L217
L217:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v882+v885<<(uint(int32(2))%32))))
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v909)+4))
	if v910 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	goto L214
L219:
	;
	v965 = v885 + int32(1)
	if v879 != v965 {
		v885 = v965
		goto L217
	} else {
		goto L235
	}
L220:
	;
	v915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414))))
	v918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v910))))
	if base.B2i32(v915 == int32(0))|base.B2i32(v915 != v918) != 0 {
		v936 = v915
		v937 = v918
		goto L222
	} else {
		goto L223
	}
L221:
	;
	if v936-v937 != 0 {
		goto L219
	} else {
		goto L228
	}
L222:
	;
	goto L221
L223:
	;
	v921 = v414
	v922 = v910
	goto L224
L224:
	;
	v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v922)+1)))
	v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v921)+1)))
	if v926 == int32(0) {
		v936 = v926
		v937 = v925
		goto L222
	} else {
		goto L226
	}
L225:
	;
	v936 = v926
	v937 = v925
	goto L222
L226:
	;
	v929 = int32(1)
	if v926 == v925 {
		v921 = v921 + v929
		v922 = v922 + v929
		goto L224
	} else {
		goto L227
	}
L227:
	;
	goto L225
L228:
	;
	v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+62)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L7
	} else {
		goto L229
	}
L229:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L7
	} else {
		goto L230
	}
L230:
	;
	if v939 == int32(1) {
		goto L105
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v414
	F_errmsg(m, int32(_a_F_transformIndexConstraints_19), v38+int32(80))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L7
	} else {
		goto L232
	}
L232:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	F_parser_errposition(m, v955, v956)
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L7
	} else {
		goto L233
	}
L233:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2779), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L7
	} else {
		goto L234
	}
L234:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L235:
	;
	goto L218
L236:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v59)+32))
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v992)+12))
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v992)+4))
	if v412 != v993+v994<<(uint(int32(2))%32)-int32(4) {
		goto L103
	} else {
		goto L237
	}
L237:
	;
	if v862 != 0 {
		v1111 = v858
		goto L104
	} else {
		goto L238
	}
L238:
	;
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+21)))
	if v1001 != int32(1) {
		goto L103
	} else {
		goto L239
	}
L239:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+52))
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v1005)))
	if v1006 <= int32(0) {
		goto L103
	} else {
		goto L240
	}
L240:
	;
	v1016 = int32(0)
	goto L241
L241:
	;
	v1039 = v1005 + v1006<<(uint(int32(3))%32) + int32(28) + v1016*int32(100)
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1039)+91)))
	if v1040 != 0 {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	goto L103
L243:
	;
	v1071 = v1016 + int32(1)
	if v1071 != v1006 {
		v1016 = v1071
		goto L241
	} else {
		goto L253
	}
L244:
	;
	v1042 = v1039 + int32(4)
	v1045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1042))))
	v1048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414))))
	if base.B2i32(v1045 == int32(0))|base.B2i32(v1045 != v1048) != 0 {
		v1066 = v1045
		v1067 = v1048
		goto L246
	} else {
		goto L247
	}
L245:
	;
	if v1066-v1067 != 0 {
		goto L243
	} else {
		goto L252
	}
L246:
	;
	goto L245
L247:
	;
	v1051 = v1042
	v1052 = v414
	goto L248
L248:
	;
	v1055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052)+1)))
	v1056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1051)+1)))
	if v1056 == int32(0) {
		v1066 = v1056
		v1067 = v1055
		goto L246
	} else {
		goto L250
	}
L249:
	;
	v1066 = v1056
	v1067 = v1055
	goto L246
L250:
	;
	v1059 = int32(1)
	if v1056 == v1055 {
		v1051 = v1051 + v1059
		v1052 = v1052 + v1059
		goto L248
	} else {
		goto L251
	}
L251:
	;
	goto L249
L252:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1039)+68))
	v1111 = v1069
	goto L104
L253:
	;
	goto L242
L254:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L7
	} else {
		goto L255
	}
L255:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v664)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+112)) = v1080
	F_errmsg(m, int32(_a_F_transformIndexConstraints_20), v38+int32(112))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L7
	} else {
		goto L256
	}
L256:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2724), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L7
	} else {
		goto L257
	}
L257:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L258:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	F_parser_errposition(m, v1098, v1099)
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L7
	} else {
		goto L259
	}
L259:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2773), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L7
	} else {
		goto L260
	}
L260:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L261:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v861)+8))
	v1136 = F_typenameTypeId(m, v1134, v1135)
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L7
	} else {
		goto L264
	}
L262:
	;
	v1138 = v1111
	goto L263
L263:
	;
	if v1138 == int32(0) {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	v1138 = v1136
	goto L263
L265:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L7
	} else {
		goto L273
	}
L266:
	;
	v1141 = F_getBaseType(m, v1138)
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L7
	} else {
		goto L267
	}
L267:
	;
	if v1141 == int32(0) {
		goto L265
	} else {
		goto L268
	}
L268:
	;
	v1145 = F_type_is_range(m, v1141)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L7
	} else {
		goto L269
	}
L269:
	;
	if v1145 != 0 {
		goto L103
	} else {
		goto L270
	}
L270:
	;
	v1147 = F_type_is_multirange(m, v1141)
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L7
	} else {
		goto L271
	}
L271:
	;
	if v1147 != 0 {
		goto L103
	} else {
		goto L272
	}
L272:
	;
	goto L265
L273:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L7
	} else {
		goto L274
	}
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v414
	F_errmsg(m, int32(_a_F_transformIndexConstraints_21), v38+int32(48))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L7
	} else {
		goto L275
	}
L275:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	F_parser_errposition(m, v1163, v1164)
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L7
	} else {
		goto L276
	}
L276:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2828), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L7
	} else {
		goto L277
	}
L277:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1195))) = int32(92)
	v1199 = F_pstrdup(m, v414)
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L7
	} else {
		goto L279
	}
L279:
	;
	v1201 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1195)+8)) = v1201
	*(*int32)(unsafe.Add(mBase, uint32(v1195)+4)) = v1199
	*(*int64)(unsafe.Add(mBase, uint32(v1195)+16)) = v1201
	*(*int64)(unsafe.Add(mBase, uint32(v1195)+24)) = v1201
	*(*int64)(unsafe.Add(mBase, uint32(v1195)+32)) = int64(-4294967296)
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	v1211 = F_lappend(m, v1210, v1195)
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L7
	} else {
		goto L280
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+20)) = v1211
	v1215 = v397 + int32(1)
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v380)+4))
	if v1215 < v1216 {
		v397 = v1215
		goto L101
	} else {
		goto L281
	}
L281:
	;
	goto L102
L282:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L7
	} else {
		goto L283
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+164)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v38)+160)) = v119
	F_errmsg(m, int32(_a_F_transformIndexConstraints_22), v38+int32(160))
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L7
	} else {
		goto L284
	}
L284:
	;
	v1234 = F_errdetail(m, int32(_a_F_transformIndexConstraints_23), int32(0))
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L7
	} else {
		goto L285
	}
L285:
	;
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	F_parser_errposition(m, v1236, v1237)
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L7
	} else {
		goto L286
	}
L286:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2581), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L7
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
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L7
	} else {
		goto L289
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+176)) = v119
	F_errmsg(m, int32(_a_F_transformIndexConstraints_24), v38+int32(176))
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L7
	} else {
		goto L290
	}
L290:
	;
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	F_parser_errposition(m, v1258, v1259)
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L7
	} else {
		goto L291
	}
L291:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2529), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L7
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
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L7
	} else {
		goto L294
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+192)) = v119
	F_errmsg(m, int32(_a_F_transformIndexConstraints_25), v38+int32(192))
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L7
	} else {
		goto L295
	}
L295:
	;
	v1282 = F_errdetail(m, int32(_a_F_transformIndexConstraints_26), int32(0))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L7
	} else {
		goto L296
	}
L296:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	F_parser_errposition(m, v1284, v1285)
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L7
	} else {
		goto L297
	}
L297:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2517), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L7
	} else {
		goto L298
	}
L298:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L299:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L7
	} else {
		goto L300
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+208)) = v119
	F_errmsg(m, int32(_a_F_transformIndexConstraints_27), v38+int32(208))
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L7
	} else {
		goto L301
	}
L301:
	;
	v1308 = F_errdetail(m, int32(_a_F_transformIndexConstraints_23), int32(0))
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L7
	} else {
		goto L302
	}
L302:
	;
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	F_parser_errposition(m, v1310, v1311)
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L7
	} else {
		goto L303
	}
L303:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2505), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1318 = m.ExcPending
	if v1318 != 0 {
		goto L7
	} else {
		goto L304
	}
L304:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L305:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L7
	} else {
		goto L306
	}
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+224)) = v119
	F_errmsg(m, int32(_a_F_transformIndexConstraints_28), v38+int32(224))
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L7
	} else {
		goto L307
	}
L307:
	;
	v1334 = F_errdetail(m, int32(_a_F_transformIndexConstraints_23), int32(0))
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L7
	} else {
		goto L308
	}
L308:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	F_parser_errposition(m, v1336, v1337)
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L7
	} else {
		goto L309
	}
L309:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2498), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L7
	} else {
		goto L310
	}
L310:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L311:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L7
	} else {
		goto L312
	}
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+240)) = v119
	F_errmsg(m, int32(_a_F_transformIndexConstraints_29), v38+int32(240))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L7
	} else {
		goto L313
	}
L313:
	;
	v1360 = F_errdetail(m, int32(_a_F_transformIndexConstraints_23), int32(0))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L7
	} else {
		goto L314
	}
L314:
	;
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	F_parser_errposition(m, v1362, v1363)
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L7
	} else {
		goto L315
	}
L315:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2491), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L7
	} else {
		goto L316
	}
L316:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L317:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L7
	} else {
		goto L318
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+256)) = v119
	F_errmsg(m, int32(_a_F_transformIndexConstraints_30), v38+int32(256))
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L7
	} else {
		goto L319
	}
L319:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	F_parser_errposition(m, v1384, v1385)
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L7
	} else {
		goto L320
	}
L320:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2479), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L7
	} else {
		goto L321
	}
L321:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L322:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L7
	} else {
		goto L323
	}
L323:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+272)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v38)+276)) = v1400 + int32(4)
	F_errmsg(m, int32(_a_F_transformIndexConstraints_31), v38+int32(272))
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L7
	} else {
		goto L324
	}
L324:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	F_parser_errposition(m, v1410, v1411)
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L7
	} else {
		goto L325
	}
L325:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2473), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
		goto L7
	} else {
		goto L326
	}
L326:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L327:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L7
	} else {
		goto L328
	}
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+288)) = v119
	F_errmsg(m, int32(_a_F_transformIndexConstraints_32), v38+int32(288))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L7
	} else {
		goto L329
	}
L329:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	F_parser_errposition(m, v1432, v1433)
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L7
	} else {
		goto L330
	}
L330:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2465), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L7
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
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L7
	} else {
		goto L333
	}
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+144)) = v119
	F_errmsg(m, int32(_a_F_transformIndexConstraints_33), v38+int32(144))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L7
	} else {
		goto L334
	}
L334:
	;
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	F_parser_errposition(m, v1454, v1455)
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L7
	} else {
		goto L335
	}
L335:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2453), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L7
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L7
	} else {
		goto L338
	}
L338:
	;
	F_errmsg(m, int32(_a_F_transformIndexConstraints_34), int32(0))
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L7
	} else {
		goto L339
	}
L339:
	;
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	F_parser_errposition(m, v1474, v1475)
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L7
	} else {
		goto L340
	}
L340:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2444), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
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
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1489 = m.ExcPending
	if v1489 != 0 {
		goto L7
	} else {
		goto L343
	}
L343:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1490)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+304)) = v1491
	F_errmsg(m, int32(_a_F_transformIndexConstraints_35), v38+int32(304))
	mBase = m.M
	v1497 = m.ExcPending
	if v1497 != 0 {
		goto L7
	} else {
		goto L344
	}
L344:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	F_parser_errposition(m, v1498, v1499)
	mBase = m.M
	v1501 = m.ExcPending
	if v1501 != 0 {
		goto L7
	} else {
		goto L345
	}
L345:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2378), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
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
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v59)+32))
	if v1532 != 0 {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+12)) = int32(_a_F_transformIndexConstraints_36)
	goto L9
L349:
	;
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v1532)+4))
	if int32(1) < v1533 {
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
	v1539 = m.ExcPending
	if v1539 != 0 {
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
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L7
	} else {
		goto L354
	}
L354:
	;
	F_errmsg(m, int32(_a_F_transformIndexConstraints_37), int32(0))
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L7
	} else {
		goto L355
	}
L355:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2856), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1551 = m.ExcPending
	if v1551 != 0 {
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
	v1996 = F_lappend(m, v48, v61)
	mBase = m.M
	v1997 = m.ExcPending
	if v1997 != 0 {
		goto L7
	} else {
		goto L435
	}
L358:
	;
	v1579 = int32(0)
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(v1576)+4))
	if v1580 <= v1579 {
		goto L357
	} else {
		goto L359
	}
L359:
	;
	v1590 = v1579
	goto L360
L360:
	;
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v1576)+12))
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1605+v1590<<(uint(int32(2))%32))))
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v1609)+4))
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	if v1611 == int32(0) {
		goto L363
	} else {
		goto L364
	}
L361:
	;
	goto L357
L362:
	;
	v1951 = F_palloc0(m, int32(40))
	mBase = m.M
	v1952 = m.ExcPending
	if v1952 != 0 {
		goto L7
	} else {
		goto L431
	}
L363:
	;
	v1700 = F_strcmp(m, int32(_a_F_transformIndexConstraints_2), v1610)
	mBase = m.M
	if v1700 == int32(0) {
		goto L378
	} else {
		goto L379
	}
L364:
	;
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+4))
	if v1614 <= int32(0) {
		goto L363
	} else {
		goto L365
	}
L365:
	;
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+12))
	v1620 = int32(0)
	goto L366
L366:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1617+v1620<<(uint(int32(2))%32))))
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v1644)+4))
	v1648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1645))))
	v1651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1610))))
	if base.B2i32(v1648 == int32(0))|base.B2i32(v1648 != v1651) != 0 {
		v1669 = v1648
		v1670 = v1651
		goto L369
	} else {
		goto L370
	}
L367:
	;
	goto L363
L368:
	;
	if v1669-v1670 == int32(0) {
		goto L362
	} else {
		goto L375
	}
L369:
	;
	goto L368
L370:
	;
	v1654 = v1645
	v1655 = v1610
	goto L371
L371:
	;
	v1658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1655)+1)))
	v1659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1654)+1)))
	if v1659 == int32(0) {
		v1669 = v1659
		v1670 = v1658
		goto L369
	} else {
		goto L373
	}
L372:
	;
	v1669 = v1659
	v1670 = v1658
	goto L369
L373:
	;
	v1662 = int32(1)
	if v1659 == v1658 {
		v1654 = v1654 + v1662
		v1655 = v1655 + v1662
		goto L371
	} else {
		goto L374
	}
L374:
	;
	goto L372
L375:
	;
	v1675 = v1620 + int32(1)
	if v1614 != v1675 {
		v1620 = v1675
		goto L366
	} else {
		goto L376
	}
L376:
	;
	goto L367
L377:
	;
	if v1729 != 0 {
		goto L362
	} else {
		goto L396
	}
L378:
	;
	v1729 = int32(_a_F_transformIndexConstraints_6)
	goto L377
L379:
	;
	goto L380
L380:
	;
	v1705 = F_strcmp(m, int32(_a_F_transformIndexConstraints_7), v1610)
	mBase = m.M
	if v1705 == int32(0) {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v1729 = int32(_a_F_transformIndexConstraints_8)
	goto L377
L382:
	;
	goto L383
L383:
	;
	v1710 = F_strcmp(m, int32(_a_F_transformIndexConstraints_9), v1610)
	mBase = m.M
	if v1710 == int32(0) {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v1729 = int32(_a_F_transformIndexConstraints_10)
	goto L377
L385:
	;
	goto L386
L386:
	;
	v1715 = F_strcmp(m, int32(_a_F_transformIndexConstraints_11), v1610)
	mBase = m.M
	if v1715 == int32(0) {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v1729 = int32(_a_F_transformIndexConstraints_12)
	goto L377
L388:
	;
	goto L389
L389:
	;
	v1720 = F_strcmp(m, int32(_a_F_transformIndexConstraints_13), v1610)
	mBase = m.M
	if v1720 == int32(0) {
		goto L390
	} else {
		goto L391
	}
L390:
	;
	v1729 = int32(_a_F_transformIndexConstraints_14)
	goto L377
L391:
	;
	goto L392
L392:
	;
	v1727 = F_strcmp(m, int32(_a_F_transformIndexConstraints_15), v1610)
	mBase = m.M
	if v1727 != 0 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v1728 = int32(0)
	goto L395
L394:
	;
	v1728 = int32(_a_F_transformIndexConstraints_16)
	goto L395
L395:
	;
	v1729 = v1728
	goto L377
L396:
	;
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v1730 == int32(0) {
		goto L398
	} else {
		goto L399
	}
L397:
	;
	F_relation_close(m, v1765, int32(0))
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L7
	} else {
		goto L430
	}
L398:
	;
	v1902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+21)))
	if v1902 != 0 {
		goto L362
	} else {
		goto L424
	}
L399:
	;
	v1733 = *(*int32)(unsafe.Add(mBase, uint32(v1730)+4))
	if v1733 <= int32(0) {
		goto L398
	} else {
		goto L400
	}
L400:
	;
	v1741 = int32(0)
	goto L401
L401:
	;
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v1730)+12))
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1759+v1741<<(uint(int32(2))%32))))
	v1765 = F_table_openrv(m, v1763, int32(1))
	mBase = m.M
	v1766 = m.ExcPending
	if v1766 != 0 {
		goto L7
	} else {
		goto L403
	}
L402:
	;
	goto L398
L403:
	;
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(v1765)+48))
	v1768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1767)+119)))
	v1770 = v1768 - int32(102)
	if base.B2i32(base.Ui32(int32(12)) < base.Ui32(v1770))|base.B2i32(int32(1)<<(uint(v1770)%32)&int32(_a_F_transformIndexConstraints_17) == int32(0)) != 0 {
		goto L1
	} else {
		goto L404
	}
L404:
	;
	v1780 = *(*int32)(unsafe.Add(mBase, uint32(v1765)+52))
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(v1780)))
	if int32(0) < v1781 {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v1791 = int32(0)
	goto L408
L406:
	;
	goto L407
L407:
	;
	F_relation_close(m, v1765, int32(0))
	mBase = m.M
	v1875 = m.ExcPending
	if v1875 != 0 {
		goto L7
	} else {
		goto L422
	}
L408:
	;
	v1814 = v1780 + v1781<<(uint(int32(3))%32) + int32(28) + v1791*int32(100)
	v1815 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1814)+91)))
	if v1815 == int32(0) {
		goto L410
	} else {
		goto L411
	}
L409:
	;
	goto L407
L410:
	;
	v1819 = v1814 + int32(4)
	v1822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1610))))
	v1825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1819))))
	if base.B2i32(v1822 == int32(0))|base.B2i32(v1822 != v1825) != 0 {
		v1843 = v1822
		v1844 = v1825
		goto L414
	} else {
		goto L415
	}
L411:
	;
	goto L412
L412:
	;
	v1849 = v1791 + int32(1)
	if v1849 != v1781 {
		v1791 = v1849
		goto L408
	} else {
		goto L421
	}
L413:
	;
	if v1843-v1844 == int32(0) {
		goto L397
	} else {
		goto L420
	}
L414:
	;
	goto L413
L415:
	;
	v1828 = v1610
	v1829 = v1819
	goto L416
L416:
	;
	v1832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1829)+1)))
	v1833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1828)+1)))
	if v1833 == int32(0) {
		v1843 = v1833
		v1844 = v1832
		goto L414
	} else {
		goto L418
	}
L417:
	;
	v1843 = v1833
	v1844 = v1832
	goto L414
L418:
	;
	v1836 = int32(1)
	if v1833 == v1832 {
		v1828 = v1828 + v1836
		v1829 = v1829 + v1836
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
	v1877 = v1741 + int32(1)
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1730)+4))
	if v1877 < v1878 {
		v1741 = v1877
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
	v1906 = m.ExcPending
	if v1906 != 0 {
		goto L7
	} else {
		goto L425
	}
L425:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L7
	} else {
		goto L426
	}
L426:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v1610
	F_errmsg(m, int32(_a_F_transformIndexConstraints_18), v38+int32(16))
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L7
	} else {
		goto L427
	}
L427:
	;
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	F_parser_errposition(m, v1916, v1917)
	mBase = m.M
	v1919 = m.ExcPending
	if v1919 != 0 {
		goto L7
	} else {
		goto L428
	}
L428:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2949), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v1924 = m.ExcPending
	if v1924 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v1951))) = int32(92)
	v1955 = F_pstrdup(m, v1610)
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L7
	} else {
		goto L432
	}
L432:
	;
	v1957 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1951)+8)) = v1957
	*(*int32)(unsafe.Add(mBase, uint32(v1951)+4)) = v1955
	*(*int64)(unsafe.Add(mBase, uint32(v1951)+16)) = v1957
	*(*int32)(unsafe.Add(mBase, uint32(v1951)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1951)+36)) = int32(-1)
	v1966 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
	v1967 = F_lappend(m, v1966, v1951)
	mBase = m.M
	v1968 = m.ExcPending
	if v1968 != 0 {
		goto L7
	} else {
		goto L433
	}
L433:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+24)) = v1967
	v1971 = v1590 + int32(1)
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v1576)+4))
	if v1971 < v1972 {
		v1590 = v1971
		goto L360
	} else {
		goto L434
	}
L434:
	;
	goto L361
L435:
	;
	v1999 = v52 + int32(1)
	v2000 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v1999 < v2000 {
		v48 = v1996
		v52 = v1999
		goto L5
	} else {
		goto L436
	}
L436:
	;
	goto L6
L437:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2007)+12)) = v2024
	*(*int32)(unsafe.Add(mBase, uint32(v2007)+316)) = v2024
	v2030 = F_list_make1_impl(m, int32(1), v2007+int32(12))
	mBase = m.M
	v2031 = m.ExcPending
	if v2031 != 0 {
		goto L7
	} else {
		goto L440
	}
L438:
	;
	v2032 = v2016
	goto L439
L439:
	;
	if v2017 == int32(0) {
		v2237 = v2032
		goto L441
	} else {
		goto L442
	}
L440:
	;
	v2032 = v2030
	goto L439
L441:
	;
	v2245 = *(*int32)(unsafe.Add(mBase, uint32(v2002)+52))
	v2246 = F_list_concat(m, v2245, v2237)
	mBase = m.M
	v2247 = m.ExcPending
	if v2247 != 0 {
		goto L7
	} else {
		goto L477
	}
L442:
	;
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v2017)+4))
	if v2035 <= int32(0) {
		v2237 = v2032
		goto L441
	} else {
		goto L443
	}
L443:
	;
	v2043 = int32(0)
	v2053 = v2032
	goto L444
L444:
	;
	v2061 = *(*int32)(unsafe.Add(mBase, uint32(v2017)+12))
	v2065 = *(*int32)(unsafe.Add(mBase, uint32(v2061+v2043<<(uint(int32(2))%32))))
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v2002)+56))
	if v2065 == v2066 {
		v2211 = v2053
		goto L446
	} else {
		goto L447
	}
L445:
	;
	v2237 = v2211
	goto L441
L446:
	;
	v2220 = v2043 + int32(1)
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(v2017)+4))
	if v2220 < v2221 {
		v2043 = v2220
		v2053 = v2211
		goto L444
	} else {
		goto L476
	}
L447:
	;
	if v2053 == int32(0) {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	v2195 = F_lappend(m, v2053, v2065)
	mBase = m.M
	v2196 = m.ExcPending
	if v2196 != 0 {
		goto L7
	} else {
		goto L475
	}
L449:
	;
	v2070 = int32(0)
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(v2053)+4))
	if v2071 <= v2070 {
		goto L448
	} else {
		goto L450
	}
L450:
	;
	v2075 = v2070
	goto L451
L451:
	;
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(v2065)+20))
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(v2053)+12))
	v2101 = *(*int32)(unsafe.Add(mBase, uint32(v2097+v2075<<(uint(int32(2))%32))))
	v2102 = *(*int32)(unsafe.Add(mBase, uint32(v2101)+20))
	v2103 = F_equal(m, v2096, v2102)
	mBase = m.M
	v2104 = m.ExcPending
	if v2104 != 0 {
		goto L7
	} else {
		goto L454
	}
L452:
	;
	goto L448
L453:
	;
	v2170 = v2075 + int32(1)
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v2053)+4))
	if v2170 < v2171 {
		v2075 = v2170
		goto L451
	} else {
		goto L474
	}
L454:
	;
	if v2103 == int32(0) {
		goto L453
	} else {
		goto L455
	}
L455:
	;
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(v2065)+24))
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(v2101)+24))
	v2109 = F_equal(m, v2107, v2108)
	mBase = m.M
	v2110 = m.ExcPending
	if v2110 != 0 {
		goto L7
	} else {
		goto L456
	}
L456:
	;
	if v2109 == int32(0) {
		goto L453
	} else {
		goto L457
	}
L457:
	;
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(v2065)+32))
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(v2101)+32))
	v2115 = F_equal(m, v2113, v2114)
	mBase = m.M
	v2116 = m.ExcPending
	if v2116 != 0 {
		goto L7
	} else {
		goto L458
	}
L458:
	;
	if v2115 == int32(0) {
		goto L453
	} else {
		goto L459
	}
L459:
	;
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v2065)+36))
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v2101)+36))
	v2121 = F_equal(m, v2119, v2120)
	mBase = m.M
	v2122 = m.ExcPending
	if v2122 != 0 {
		goto L7
	} else {
		goto L460
	}
L460:
	;
	if v2121 == int32(0) {
		goto L453
	} else {
		goto L461
	}
L461:
	;
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(v2065)+12))
	v2126 = *(*int32)(unsafe.Add(mBase, uint32(v2101)+12))
	v2129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2125))))
	v2132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2126))))
	if base.B2i32(v2129 == int32(0))|base.B2i32(v2129 != v2132) != 0 {
		v2150 = v2129
		v2151 = v2132
		goto L463
	} else {
		goto L464
	}
L462:
	;
	if v2150-v2151 != 0 {
		goto L453
	} else {
		goto L469
	}
L463:
	;
	goto L462
L464:
	;
	v2135 = v2125
	v2136 = v2126
	goto L465
L465:
	;
	v2139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2136)+1)))
	v2140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2135)+1)))
	if v2140 == int32(0) {
		v2150 = v2140
		v2151 = v2139
		goto L463
	} else {
		goto L467
	}
L466:
	;
	v2150 = v2140
	v2151 = v2139
	goto L463
L467:
	;
	v2143 = int32(1)
	if v2140 == v2139 {
		v2135 = v2135 + v2143
		v2136 = v2136 + v2143
		goto L465
	} else {
		goto L468
	}
L468:
	;
	goto L466
L469:
	;
	v2153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2065)+61)))
	v2154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2101)+61)))
	if v2153 != v2154 {
		goto L453
	} else {
		goto L470
	}
L470:
	;
	v2156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2065)+65)))
	v2157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2101)+65)))
	if v2156 != v2157 {
		goto L453
	} else {
		goto L471
	}
L471:
	;
	v2159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2065)+66)))
	v2160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2101)+66)))
	if v2159 != v2160 {
		goto L453
	} else {
		goto L472
	}
L472:
	;
	v2162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2101)+60)))
	v2163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2065)+60)))
	v2164 = v2162 | v2163
	*(*uint8)(unsafe.Add(mBase, uint32(v2101)+60)) = uint8(v2164)
	v2166 = *(*int32)(unsafe.Add(mBase, uint32(v2101)+4))
	if v2166 != 0 {
		v2211 = v2053
		goto L446
	} else {
		goto L473
	}
L473:
	;
	v2167 = *(*int32)(unsafe.Add(mBase, uint32(v2065)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2101)+4)) = v2167
	v2211 = v2053
	goto L446
L474:
	;
	goto L452
L475:
	;
	v2211 = v2195
	goto L446
L476:
	;
	goto L445
L477:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2002)+52)) = v2246
	m.G0 = v2007 + int32(320)
	return
L478:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v2258 = m.ExcPending
	if v2258 != 0 {
		goto L7
	} else {
		goto L479
	}
L479:
	;
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(v1763)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = v2259
	F_errmsg(m, int32(_a_F_transformIndexConstraints_20), v38+int32(32))
	mBase = m.M
	v2265 = m.ExcPending
	if v2265 != 0 {
		goto L7
	} else {
		goto L480
	}
L480:
	;
	F_errfinish(m, int32(_a_F_transformIndexConstraints_4), int32(2918), int32(_a_F_transformIndexConstraints_5))
	mBase = m.M
	v2270 = m.ExcPending
	if v2270 != 0 {
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
	F_errfinish(m, int32(_a_F_transformIndexStmt_3), int32(3154), int32(_a_F_transformIndexStmt_4))
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
