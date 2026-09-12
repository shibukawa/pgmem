package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ER_flatten_into(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v10&int32(17) == int32(1) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
		if l2 != 0 {
			v17 = F__emscripten_memcpy_bulkmem(m, l1, v16, l2)
			mBase = m.M
			v18 = v17
		} else {
			v18 = l1
		}
		*(*int32)(unsafe.Add(mBase, uint32(v18))) = l2 << (uint(int32(2)) % 32)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v22
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v24
		return
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v26 == int32(0) {
			v29 = F_expanded_record_fetch_tupdesc(m, l0)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				v31 = v29
				v34 = F__emscripten_memset_bulkmem(m, l1, base.I32_extend8_s(int32(0)), l2)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v34))) = l2 << (uint(int32(2)) % 32)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v38
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v41 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v34)+16)) = uint16(v41)
				*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v40
				v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31))))
				v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)))
				v50 = v46 | v47&int32(63488)
				*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)) = uint16(v50)
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)) = uint8(v52)
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)))
				if v62 != 0 {
					v63 = v34 + int32(23)
				} else {
					v63 = v41
				}
				F_heap_fill_tuple(m, v31, v54, v55, v34+v56, v34+int32(20), v63)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v31 = v26
			v34 = F__emscripten_memset_bulkmem(m, l1, base.I32_extend8_s(int32(0)), l2)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v34))) = l2 << (uint(int32(2)) % 32)
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v38
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v41 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v34)+16)) = uint16(v41)
			*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v40
			v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31))))
			v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)))
			v50 = v46 | v47&int32(63488)
			*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)) = uint16(v50)
			v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)) = uint8(v52)
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)))
			if v62 != 0 {
				v63 = v34 + int32(23)
			} else {
				v63 = v41
			}
			F_heap_fill_tuple(m, v31, v54, v55, v34+v56, v34+int32(20), v63)
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_EvictUnpinnedBufferInternal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	v3 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v3)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v8&int32(16777216) == v3 {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v8 & int32(-20971521)
		return int32(0)
	} else {
		if v8&int32(262143) != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v8 & int32(-4194305)
			return int32(0)
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v26 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = (v25 + v26) & int32(-4194305)
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v32 = int32(4477976)
			v33 = *(*int32)(unsafe.Add(mBase, _consts[278]))
			*(*int32)(unsafe.Add(mBase, _consts[278])) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v26
			v40 = v31 + v26
			*(*int32)(unsafe.Add(mBase, uint32(v33))) = v40
			v43 = *(*int32)(unsafe.Add(mBase, _consts[175]))
			F_ResourceOwnerRemember(m, v43, v40, int32(1664464))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				if v8&int32(8388608) != 0 {
					v52 = l0 + int32(48)
					v54 = F_LWLockAcquire(m, v52, int32(1))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_FlushBuffer(m, l0, int32(0), int32(3))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							v60 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v60)
							F_LWLockRelease(m, v52)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v65 = F_InvalidateVictimBuffer(m, l0)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, _consts[175]))
									v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									F_ResourceOwnerForget(m, v68, v69+int32(1), int32(1664464))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										F_UnpinBufferNoOwner(m, l0)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											return v65
										}
									}
								}
							}
						}
					}
				} else {
					v65 = F_InvalidateVictimBuffer(m, l0)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, _consts[175]))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						F_ResourceOwnerForget(m, v68, v69+int32(1), int32(1664464))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							F_UnpinBufferNoOwner(m, l0)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								return v65
							}
						}
					}
				}
			}
		}
	}
}
func F_ExecASInsertTriggers(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v5 == int32(0) {
		return
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+12)))
		if v8 != int32(1) {
			return
		} else {
			v11 = int32(0)
			F_AfterTriggerSaveEvent(m, l0, l1, v11, v11, v11, v11, v11, v11, v11, v11, l2, v11)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_ExecAppend(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
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
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v598 int32
	_ = v598
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v7 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	goto L83
L4:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	m.T0[v12].(func(*base.Module, int32))(m, v10)
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
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v18 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	return int32(0)
L8:
	;
	return v10
L9:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v272 = m.T0[v271].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L7
	} else {
		goto L78
	}
L10:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)))
	if v21 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v96 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)) = uint8(base.B2i32(v95 == v96))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v99 == v96 {
		goto L37
	} else {
		goto L38
	}
L12:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v23 = int32(0)
	v25 = F_ExecFindMatchingSubPlans(m, v22, v23, v23)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v27 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)) = uint8(v27)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v25
	if v25 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(0)
	goto L11
L15:
	;
	goto L16
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v35 = int32(0)
	if v25 == v35 {
		v76 = v35
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v76 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L18:
	;
	goto L17
L19:
	;
	if v34 == int32(0) {
		v76 = v35
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v44 < v45 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v47 = v44
	goto L23
L22:
	;
	v47 = v45
	goto L23
L23:
	;
	if v47 <= int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v50 = int32(1)
	goto L26
L25:
	;
	v50 = v47
	goto L26
L26:
	;
	v51 = int32(8)
	v56 = int32(0)
	goto L27
L27:
	;
	v63 = v56 << (uint(int32(2)) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v34+v51+v63)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+(v25+v51))))
	v68 = v65 & v67
	v70 = base.B2i32(v68 != int32(0))
	if v68 != 0 {
		v76 = v70
		goto L18
	} else {
		goto L29
	}
L28:
	;
	v76 = v70
	goto L18
L29:
	;
	v72 = v56 + int32(1)
	if v72 != v50 {
		v56 = v72
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(0)
	goto L11
L32:
	;
	goto L33
L33:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v86 = F_bms_intersect(m, v84, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v89 = F_bms_del_members(m, v88, v86)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v89
	goto L11
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v135
	if v135 == int32(0) {
		goto L9
	} else {
		goto L49
	}
L37:
	;
	v135 = int32(0)
	goto L36
L38:
	;
	goto L39
L39:
	;
	v107 = int32(1)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v108 <= v107 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v111 = v107
	goto L42
L41:
	;
	v111 = v108
	goto L42
L42:
	;
	v115 = int32(0)
	v117 = v96
	goto L43
L43:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v99+int32(8)+v115<<(uint(int32(2))%32))))
	if v123 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v135 = v126
	goto L36
L45:
	;
	v126 = v117 + base.I32_popcnt(v123)
	goto L47
L46:
	;
	v126 = v117
	goto L47
L47:
	;
	v128 = v115 + int32(1)
	if v128 != v111 {
		v115 = v128
		v117 = v126
		goto L43
	} else {
		goto L48
	}
L48:
	;
	goto L44
L49:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v139 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	if v196 < int32(0) {
		goto L9
	} else {
		goto L61
	}
L51:
	;
	v196 = base.I32_ctz(v182) | v183<<(uint(int32(5))%32)
	goto L50
L52:
	;
	v196 = int32(-2)
	goto L50
L53:
	;
	v149 = base.I32_div_s(int32(0), int32(32))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v150 <= v149 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v153 = v139 + int32(8)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153+v149<<(uint(int32(2))%32))))
	v160 = v157 & int32(-1)
	if v160 != 0 {
		v182 = v160
		v183 = v149
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v162 = v149 + int32(1)
	if v162 == v150 {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v165 = v162
	goto L57
L57:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v153+v165<<(uint(int32(2))%32))))
	if v172 != 0 {
		v182 = v172
		v183 = v165
		goto L51
	} else {
		goto L59
	}
L58:
	;
	goto L52
L59:
	;
	v174 = v165 + int32(1)
	if v174 != v150 {
		v165 = v174
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v200 = v196
	goto L62
L62:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202+v200<<(uint(int32(2))%32))))
	F_ExecAsyncRequest(m, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L7
	} else {
		goto L64
	}
L63:
	;
	goto L9
L64:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v209 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	if int32(0) <= v265 {
		v200 = v265
		goto L62
	} else {
		goto L76
	}
L66:
	;
	v265 = base.I32_ctz(v251) | v252<<(uint(int32(5))%32)
	goto L65
L67:
	;
	v265 = int32(-2)
	goto L65
L68:
	;
	v216 = v200 + int32(1)
	v218 = base.I32_div_s(v216, int32(32))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	if v219 <= v218 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v222 = v209 + int32(8)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v222+v218<<(uint(int32(2))%32))))
	v229 = v226 & (int32(-1) << (uint(v216) % 32))
	if v229 != 0 {
		v251 = v229
		v252 = v218
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v231 = v218 + int32(1)
	if v231 == v219 {
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v234 = v231
	goto L72
L72:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v222+v234<<(uint(int32(2))%32))))
	if v241 != 0 {
		v251 = v241
		v252 = v234
		goto L66
	} else {
		goto L74
	}
L73:
	;
	goto L67
L74:
	;
	v243 = v234 + int32(1)
	if v243 != v219 {
		v234 = v243
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	goto L63
L77:
	;
	v281 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)) = uint8(v281)
	goto L3
L78:
	;
	if v272 != 0 {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v274 != 0 {
		goto L77
	} else {
		goto L80
	}
L80:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+8))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+12))
	m.T0[v277].(func(*base.Module, int32))(m, v275)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L7
	} else {
		goto L81
	}
L81:
	;
	return v275
L82:
	;
	return v670
L83:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v290 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v665)+8))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v666)+12))
	m.T0[v667].(func(*base.Module, int32))(m, v665)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L7
	} else {
		goto L196
	}
L85:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L7
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
	if v294 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L88:
	;
	goto L87
L89:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v639+v640<<(uint(int32(2))%32))))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v644)+52))
	if v645 != 0 {
		goto L180
	} else {
		goto L181
	}
L90:
	;
	goto L131
L91:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v299 <= int32(0) {
		goto L97
	} else {
		goto L98
	}
L92:
	;
	if v293 != 0 {
		goto L91
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	if v293 == int32(0) {
		goto L90
	} else {
		goto L96
	}
L95:
	;
	goto L89
L96:
	;
	goto L91
L97:
	;
	v302 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v302
	if v293 == v302 {
		goto L102
	} else {
		goto L103
	}
L98:
	;
	v440 = v299
	goto L99
L99:
	;
	v443 = v440 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v443
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v445+v443<<(uint(int32(2))%32))))
	return v449
L100:
	;
	if int32(0) <= v360 {
		goto L111
	} else {
		goto L112
	}
L101:
	;
	v360 = base.I32_ctz(v346) | v347<<(uint(int32(5))%32)
	goto L100
L102:
	;
	v360 = int32(-2)
	goto L100
L103:
	;
	v313 = base.I32_div_s(int32(0), int32(32))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	if v314 <= v313 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v317 = v293 + int32(8)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v317+v313<<(uint(int32(2))%32))))
	v324 = v321 & int32(-1)
	if v324 != 0 {
		v346 = v324
		v347 = v313
		goto L101
	} else {
		goto L105
	}
L105:
	;
	v326 = v313 + int32(1)
	if v326 == v314 {
		goto L102
	} else {
		goto L106
	}
L106:
	;
	v329 = v326
	goto L107
L107:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v317+v329<<(uint(int32(2))%32))))
	if v336 != 0 {
		v346 = v336
		v347 = v329
		goto L101
	} else {
		goto L109
	}
L108:
	;
	goto L102
L109:
	;
	v338 = v329 + int32(1)
	if v338 != v314 {
		v329 = v338
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v364 = v360
	goto L114
L112:
	;
	goto L113
L113:
	;
	F_bms_free(m, v293)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L7
	} else {
		goto L129
	}
L114:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v366+v364<<(uint(int32(2))%32))))
	F_ExecAsyncRequest(m, v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L7
	} else {
		goto L116
	}
L115:
	;
	goto L113
L116:
	;
	if v293 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	if int32(0) <= v428 {
		v364 = v428
		goto L114
	} else {
		goto L128
	}
L118:
	;
	v428 = base.I32_ctz(v414) | v415<<(uint(int32(5))%32)
	goto L117
L119:
	;
	v428 = int32(-2)
	goto L117
L120:
	;
	v379 = v364 + int32(1)
	v381 = base.I32_div_s(v379, int32(32))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	if v382 <= v381 {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v385 = v293 + int32(8)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v385+v381<<(uint(int32(2))%32))))
	v392 = v389 & (int32(-1) << (uint(v379) % 32))
	if v392 != 0 {
		v414 = v392
		v415 = v381
		goto L118
	} else {
		goto L122
	}
L122:
	;
	v394 = v381 + int32(1)
	if v394 == v382 {
		goto L119
	} else {
		goto L123
	}
L123:
	;
	v397 = v394
	goto L124
L124:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v385+v397<<(uint(int32(2))%32))))
	if v404 != 0 {
		v414 = v404
		v415 = v397
		goto L118
	} else {
		goto L126
	}
L125:
	;
	goto L119
L126:
	;
	v406 = v397 + int32(1)
	if v406 != v382 {
		v397 = v406
		goto L124
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	goto L115
L129:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v436 <= int32(0) {
		goto L90
	} else {
		goto L130
	}
L130:
	;
	v440 = v436
	goto L99
L131:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if int32(0) < v457 {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	goto L89
L133:
	;
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
	if v635 != 0 {
		goto L131
	} else {
		goto L179
	}
L134:
	;
	v461 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v461 != 0 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	goto L136
L136:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
	if v621&int32(1) == int32(0) {
		goto L89
	} else {
		goto L177
	}
L137:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L7
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	F_ExecAppendAsyncEventWait(m, l0)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L7
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v466 == int32(0) {
		goto L133
	} else {
		goto L142
	}
L142:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v469 <= int32(0) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v472 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v472
	if v466 == v472 {
		goto L148
	} else {
		goto L149
	}
L144:
	;
	v610 = v469
	goto L145
L145:
	;
	v613 = v610 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v613
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v615+v613<<(uint(int32(2))%32))))
	return v619
L146:
	;
	if int32(0) <= v530 {
		goto L157
	} else {
		goto L158
	}
L147:
	;
	v530 = base.I32_ctz(v516) | v517<<(uint(int32(5))%32)
	goto L146
L148:
	;
	v530 = int32(-2)
	goto L146
L149:
	;
	v483 = base.I32_div_s(int32(0), int32(32))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v466)+4))
	if v484 <= v483 {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v487 = v466 + int32(8)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v487+v483<<(uint(int32(2))%32))))
	v494 = v491 & int32(-1)
	if v494 != 0 {
		v516 = v494
		v517 = v483
		goto L147
	} else {
		goto L151
	}
L151:
	;
	v496 = v483 + int32(1)
	if v496 == v484 {
		goto L148
	} else {
		goto L152
	}
L152:
	;
	v499 = v496
	goto L153
L153:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v487+v499<<(uint(int32(2))%32))))
	if v506 != 0 {
		v516 = v506
		v517 = v499
		goto L147
	} else {
		goto L155
	}
L154:
	;
	goto L148
L155:
	;
	v508 = v499 + int32(1)
	if v508 != v484 {
		v499 = v508
		goto L153
	} else {
		goto L156
	}
L156:
	;
	goto L154
L157:
	;
	v534 = v530
	goto L160
L158:
	;
	goto L159
L159:
	;
	F_bms_free(m, v466)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L7
	} else {
		goto L175
	}
L160:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v536+v534<<(uint(int32(2))%32))))
	F_ExecAsyncRequest(m, v540)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L7
	} else {
		goto L162
	}
L161:
	;
	goto L159
L162:
	;
	if v466 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	if int32(0) <= v598 {
		v534 = v598
		goto L160
	} else {
		goto L174
	}
L164:
	;
	v598 = base.I32_ctz(v584) | v585<<(uint(int32(5))%32)
	goto L163
L165:
	;
	v598 = int32(-2)
	goto L163
L166:
	;
	v549 = v534 + int32(1)
	v551 = base.I32_div_s(v549, int32(32))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v466)+4))
	if v552 <= v551 {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v555 = v466 + int32(8)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v555+v551<<(uint(int32(2))%32))))
	v562 = v559 & (int32(-1) << (uint(v549) % 32))
	if v562 != 0 {
		v584 = v562
		v585 = v551
		goto L164
	} else {
		goto L168
	}
L168:
	;
	v564 = v551 + int32(1)
	if v564 == v552 {
		goto L165
	} else {
		goto L169
	}
L169:
	;
	v567 = v564
	goto L170
L170:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v555+v567<<(uint(int32(2))%32))))
	if v574 != 0 {
		v584 = v574
		v585 = v567
		goto L164
	} else {
		goto L172
	}
L171:
	;
	goto L165
L172:
	;
	v576 = v567 + int32(1)
	if v576 != v552 {
		v567 = v576
		goto L170
	} else {
		goto L173
	}
L173:
	;
	goto L171
L174:
	;
	goto L161
L175:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v606 <= int32(0) {
		goto L133
	} else {
		goto L176
	}
L176:
	;
	v610 = v606
	goto L145
L177:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v626)+8))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v627)+12))
	m.T0[v628].(func(*base.Module, int32))(m, v626)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L7
	} else {
		goto L178
	}
L178:
	;
	return v626
L179:
	;
	goto L132
L180:
	;
	F_ExecReScan(m, v644)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L7
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v644)+12))
	v649 = m.T0[v648].(func(*base.Module, int32) int32)(m, v644)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L7
	} else {
		goto L184
	}
L183:
	;
	goto L182
L184:
	;
	if v649 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649)+4)))
	if v651&int32(2) == int32(0) {
		v670 = v649
		goto L82
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if int32(0) < v656 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	goto L187
L189:
	;
	F_ExecAppendAsyncEventWait(m, l0)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L7
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v662 = m.T0[v661].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L7
	} else {
		goto L193
	}
L192:
	;
	goto L191
L193:
	;
	if v662 != 0 {
		goto L83
	} else {
		goto L194
	}
L194:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v664 != 0 {
		goto L83
	} else {
		goto L195
	}
L195:
	;
	goto L84
L196:
	;
	v670 = v665
	goto L82
}
func F_ExecCheck(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 != 0 {
		v10 = int32(4562096)
		v11 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		*(*int32)(unsafe.Add(mBase, _consts[3])) = v13
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v18 = m.T0[v17].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, v7+int32(15))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[3])) = v11
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
			v29 = v24 | base.B2i32(v18 != int32(0))
			m.G0 = v7 + int32(16)
			return v29 & int32(1)
		}
	} else {
		v29 = int32(1)
		m.G0 = v7 + int32(16)
		return v29 & int32(1)
	}
}
func F_ExecCheckOneRelPerms(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v185 int32
	_ = v185
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v9 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v14 = v13
	goto L3
L2:
	;
	v14 = v9
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = int32(0)
	v18 = F_pg_class_aclmask_ext(m, v15, v14, v7, v16, v16)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v218
L5:
	;
	return int32(0)
L6:
	;
	v24 = v7 & (v18 ^ int64(-1))
	if v24 == int64(0) {
		v218 = int32(1)
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if base.Ui64(int64(7)) < base.Ui64(v24) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	goto L10
L10:
	;
	if v24&int64(2) == int64(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v24&int64(1) == int64(0) {
		goto L54
	} else {
		goto L55
	}
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v35 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v43 = v35
	goto L15
L14:
	;
	v38 = F_pg_attribute_aclcheck_all(m, v15, v14, int64(2), int32(1))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L16
	}
L15:
	;
	if v43 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	if v38 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	goto L19
L19:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v43 = v42
	goto L15
L20:
	;
	if v100 < int32(0) {
		goto L11
	} else {
		goto L31
	}
L21:
	;
	v100 = base.I32_ctz(v86) | v87<<(uint(int32(5))%32)
	goto L20
L22:
	;
	v100 = int32(-2)
	goto L20
L23:
	;
	v53 = base.I32_div_s(int32(0), int32(32))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v54 <= v53 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v57 = v43 + int32(8)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v53<<(uint(int32(2))%32))))
	v64 = v61 & int32(-1)
	if v64 != 0 {
		v86 = v64
		v87 = v53
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v66 = v53 + int32(1)
	if v66 == v54 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v69 = v66
	goto L27
L27:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v57+v69<<(uint(int32(2))%32))))
	if v76 != 0 {
		v86 = v76
		v87 = v69
		goto L21
	} else {
		goto L29
	}
L28:
	;
	goto L22
L29:
	;
	v78 = v69 + int32(1)
	if v78 != v54 {
		v69 = v78
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v108 = v100
	goto L32
L32:
	;
	v110 = v108 - int32(7)
	if v110&int32(65535) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L11
L34:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v129 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L35:
	;
	v115 = int32(0)
	v118 = F_pg_attribute_aclcheck_all(m, v15, v14, int64(2), v115)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v122 = F_pg_attribute_aclcheck(m, v15, base.I32_extend16_s(v110), v14, int64(2))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L5
	} else {
		goto L40
	}
L38:
	;
	if v118 != 0 {
		v218 = v115
		goto L4
	} else {
		goto L39
	}
L39:
	;
	goto L34
L40:
	;
	if v122 == int32(0) {
		goto L34
	} else {
		goto L41
	}
L41:
	;
	return int32(0)
L42:
	;
	if int32(0) <= v185 {
		v108 = v185
		goto L32
	} else {
		goto L53
	}
L43:
	;
	v185 = base.I32_ctz(v171) | v172<<(uint(int32(5))%32)
	goto L42
L44:
	;
	v185 = int32(-2)
	goto L42
L45:
	;
	v136 = v108 + int32(1)
	v138 = base.I32_div_s(v136, int32(32))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v139 <= v138 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v142 = v129 + int32(8)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142+v138<<(uint(int32(2))%32))))
	v149 = v146 & (int32(-1) << (uint(v136) % 32))
	if v149 != 0 {
		v171 = v149
		v172 = v138
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v151 = v138 + int32(1)
	if v151 == v139 {
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v154 = v151
	goto L49
L49:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v142+v154<<(uint(int32(2))%32))))
	if v161 != 0 {
		v171 = v161
		v172 = v154
		goto L43
	} else {
		goto L51
	}
L50:
	;
	goto L44
L51:
	;
	v163 = v154 + int32(1)
	if v163 != v139 {
		v154 = v163
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	goto L33
L54:
	;
	if base.Ui64(int64(4)) <= base.Ui64(v24) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v200 = F_ExecCheckPermissionsModified(m, v15, v14, v198, int64(1))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	if v200 != 0 {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	return int32(0)
L58:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v209 = F_ExecCheckPermissionsModified(m, v15, v14, v207, int64(4))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L5
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v218 = int32(1)
	goto L4
L61:
	;
	if v209 == int32(0) {
		v218 = int32(0)
		goto L4
	} else {
		goto L62
	}
L62:
	;
	goto L60
}
func F_ExecCleanTargetListLength(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v2 = int32(0)
	if l0 == v2 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v13 <= int32(0) {
			v77 = v2
		} else {
			v16 = int32(0)
			if v16 < v13 {
				v19 = v13
			} else {
				v19 = v16
			}
			v20 = int32(1)
			if v13 == v20 {
				v24 = int32(0)
				v58 = v24
				v59 = v24
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v29 = int32(0)
				v32 = v29
				v33 = v29
				v34 = v2
				for {
					v39 = int32(2)
					v41 = v28 + v33<<(uint(v39)%32)
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+26)))
					v44 = int32(1)
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+26)))
					v51 = v32 + (v43 ^ v44) + (v48 ^ v44)
					v53 = v33 + v39
					v55 = v34 + v39
					if v55 != v19&int32(2147483646) {
						v32 = v51
						v33 = v53
						v34 = v55
						continue
					} else {
						break
					}
					break
				}
				v58 = v51
				v59 = v53
			}
			if v19&v20 == int32(0) {
				v77 = v58
			} else {
				v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v67+v59<<(uint(int32(2))%32))))
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+26)))
				v77 = v58 + (v72 ^ int32(1))
			}
		}
		return v77
	}
}
func F_ExecConditionalAssignProjectionInfo(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
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
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v15 = v13
	goto L3
L2:
	;
	v15 = int32(0)
	goto L3
L3:
	;
	if int32(0) < v10 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+v136))) = v142
	return
L5:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v96 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L6:
	;
	v27 = int32(1)
	v29 = v15
	goto L9
L7:
	;
	v75 = v15
	goto L8
L8:
	;
	if v75 != 0 {
		goto L5
	} else {
		goto L26
	}
L9:
	;
	if v29 == int32(0) {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	v75 = v66
	goto L8
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v36 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v39 != int32(6) {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v42 = int32(*(*int16)(unsafe.Add(mBase, uint32(v36)+8)))
	if v27 != v42 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v46 = l1 + v10<<(uint(int32(4))%32) - int32(80) + v27*int32(100)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+91)))
	if v47 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+88)))
	if v48 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+68))
	if v49 != v50 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	if v52 != int32(-1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v46)+76))
	if v52 != v55 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v58 = v29 + int32(4)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if base.Ui32(v58) < base.Ui32(v60+v61<<(uint(int32(2))%32)) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v66 = v58
	goto L24
L23:
	;
	v66 = int32(0)
	goto L24
L24:
	;
	if v27 != v10 {
		v27 = v27 + int32(1)
		v29 = v66
		goto L9
	} else {
		goto L25
	}
L25:
	;
	goto L10
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = int32(0)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+103)) = uint8(v81)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+99)) = uint8(v83)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v136 = int32(92)
	v142 = v86
	goto L4
L27:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v102 = F_MakeTupleTableSlot(m, v100, int32(1654020))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v126 = v96
	v128 = v12
	goto L29
L29:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v131 = F_ExecBuildProjectionInfo(m, v128, v130, v126, l0, l1)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L30
	} else {
		goto L33
	}
L30:
	;
	return
L31:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)+104))
	v105 = F_lappend(m, v104, v102)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+104)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v102
	v109 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+103)) = uint8(v109)
	v111 = int32(1654020)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v111
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+99)) = uint8(base.B2i32(v113 != int32(0)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+103)) = uint8(v109)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+99)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v111
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+44))
	v126 = v102
	v128 = v124
	goto L29
L33:
	;
	v136 = int32(68)
	v142 = v131
	goto L4
}
func F_ExecForceStoreHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
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
	var v59 int32
	_ = v59
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v6 == int32(1654072) {
		v9 = F_ExecStoreHeapTuple(m, l0, l1, l2)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	} else {
		if v6 == int32(1654176) {
			v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			if v13&int32(4) != 0 {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
				F_pfree(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
					v21 = v19 & int32(-5)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v21)
					v23 = v21
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
					if v24 != 0 {
						F_ReleaseBuffer(m, v24)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
							v28 = v27
							v29 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v29
							*(*int64)(unsafe.Add(mBase, uint32(l1)+40)) = int64(0)
							*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v29)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
							*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v29)
							v40 = v28 & int32(65533)
							*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v40)
							v42 = int32(4562096)
							v43 = *(*int32)(unsafe.Add(mBase, _consts[3]))
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							*(*int32)(unsafe.Add(mBase, _consts[3])) = v45
							v47 = F_heap_copytuple(m, l0)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v47
								v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
								v52 = v50 | int32(4)
								*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v52)
								*(*int32)(unsafe.Add(mBase, _consts[3])) = v43
								if l2 != 0 {
									F_pfree(m, l0)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										return
									}
								} else {
									return
								}
							}
						}
					} else {
						v28 = v23
						v29 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v29
						*(*int64)(unsafe.Add(mBase, uint32(l1)+40)) = int64(0)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v29)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v29)
						v40 = v28 & int32(65533)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v40)
						v42 = int32(4562096)
						v43 = *(*int32)(unsafe.Add(mBase, _consts[3]))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						*(*int32)(unsafe.Add(mBase, _consts[3])) = v45
						v47 = F_heap_copytuple(m, l0)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v47
							v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
							v52 = v50 | int32(4)
							*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v52)
							*(*int32)(unsafe.Add(mBase, _consts[3])) = v43
							if l2 != 0 {
								F_pfree(m, l0)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									return
								}
							} else {
								return
							}
						}
					}
				}
			} else {
				v23 = v13
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
				if v24 != 0 {
					F_ReleaseBuffer(m, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
						v28 = v27
						v29 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v29
						*(*int64)(unsafe.Add(mBase, uint32(l1)+40)) = int64(0)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v29)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v29)
						v40 = v28 & int32(65533)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v40)
						v42 = int32(4562096)
						v43 = *(*int32)(unsafe.Add(mBase, _consts[3]))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						*(*int32)(unsafe.Add(mBase, _consts[3])) = v45
						v47 = F_heap_copytuple(m, l0)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v47
							v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
							v52 = v50 | int32(4)
							*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v52)
							*(*int32)(unsafe.Add(mBase, _consts[3])) = v43
							if l2 != 0 {
								F_pfree(m, l0)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									return
								}
							} else {
								return
							}
						}
					}
				} else {
					v28 = v23
					v29 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v29
					*(*int64)(unsafe.Add(mBase, uint32(l1)+40)) = int64(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v29)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v29)
					v40 = v28 & int32(65533)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v40)
					v42 = int32(4562096)
					v43 = *(*int32)(unsafe.Add(mBase, _consts[3]))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
					*(*int32)(unsafe.Add(mBase, _consts[3])) = v45
					v47 = F_heap_copytuple(m, l0)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v47
						v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
						v52 = v50 | int32(4)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v52)
						*(*int32)(unsafe.Add(mBase, _consts[3])) = v43
						if l2 != 0 {
							F_pfree(m, l0)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								return
							}
						} else {
							return
						}
					}
				}
			}
		} else {
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			m.T0[v56].(func(*base.Module, int32))(m, l1)
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				F_heap_deform_tuple(m, l0, v59, v60, v61)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
					v66 = v64 & int32(65533)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v66)
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v69)
					if l2 == int32(0) {
						return
					} else {
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+28))
						m.T0[v74].(func(*base.Module, int32))(m, l1)
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return
						} else {
							F_pfree(m, l0)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_ExecGrant_Language_check(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
	v9 = v7 + v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+73)))
	if v10 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			F_errcode(m, int32(151027844))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v9 + int32(4)
				F_errmsg(m, int32(462961), v5)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					F_errdetail(m, int32(631447), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						F_errfinish(m, int32(523480), int32(2264), int32(334589))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
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
		m.G0 = v5 + int32(16)
		return
	}
}
func F_ExecGrant_Type_check(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+22)))
	v5 = v3 + v4
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+92))
	if v6 != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+88))
		if v7 == int32(6179) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errcode(m, int32(16910080))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_errmsg(m, int32(172244), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						F_errhint(m, int32(682672), int32(0))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							F_errfinish(m, int32(523480), int32(2415), int32(334568))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
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
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+79)))
			if v10 == int32(109) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					F_errcode(m, int32(16910080))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						F_errmsg(m, int32(172860), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							F_errhint(m, int32(682720), int32(0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								F_errfinish(m, int32(523480), int32(2420), int32(334568))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
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
				return
			}
		}
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+79)))
		if v10 == int32(109) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				F_errcode(m, int32(16910080))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					F_errmsg(m, int32(172860), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						F_errhint(m, int32(682720), int32(0))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							F_errfinish(m, int32(523480), int32(2420), int32(334568))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
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
			return
		}
	}
}
func F_ExecGrant_common(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v55 int32
	_ = v55
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
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
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
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
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
	var v114 int64
	_ = v114
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
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int64
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
	var v141 int32
	_ = v141
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
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
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
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v23 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = F_get_object_catcache_oid(m, l1)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v26 != int64(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = l2
	goto L1
L4:
	;
	return
L5:
	;
	v33 = F_table_open(m, l1, int32(3))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v35 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L58
	}
L8:
	;
	F_sequence_close(m, v33, int32(3))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L57
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v38 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v55 = int32(0)
	goto L11
L11:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v60 = int32(2)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59+v55<<(uint(v60)%32))))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v68 = F_palloc0(m, v65<<(uint(v60)%32))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	goto L8
L13:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v72 = F_palloc0(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v76 = F_palloc0(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v78 = F_SearchSysCacheLocked1(m, v30, v63)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if v78 == int32(0) {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	if l3 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	m.T0[l3].(func(*base.Module, int32, int32))(m, l0, v78)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v84 = F_get_object_attnum_owner(m, l1)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v86 = F_SysCacheGetAttrNotNull(m, v30, v78, v84)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v88 = F_get_object_attnum_acl(m, l1)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v92 = F_SysCacheGetAttr(m, v30, v78, v88, v21+int32(47))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)))
	if v94 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	F_select_best_grantor(m, v113, v114, v110, v86, v21+int32(28), v21+int32(32))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L34
	}
L27:
	;
	v98 = F_get_object_type(m, l1, v63)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v104 = F_pg_detoast_datum_copy(m, v92)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	v100 = F_acldefault(m, v98, v86)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = int32(0)
	v110 = v100
	v111 = int32(0)
	goto L26
L32:
	;
	v108 = F_aclmembers(m, v104, v21+int32(24))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v110 = v104
	v111 = v108
	goto L26
L34:
	;
	v121 = F_get_object_attnum_name(m, l1)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v123 = F_SysCacheGetAttrNotNull(m, v30, v78, v121)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v21)+32))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v128 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v130 = F_get_object_type(m, l1, v63)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v132 = int32(0)
	v134 = F_restrict_and_check_grant(m, v125, v126, v127, v128, v63, v129, v130, v123, v132, v132)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v141 = F_merge_acl_with_grant(m, v110, v136, v137, v138, v139, v134, v140, v86)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v145 = F_aclmembers(m, v141, v21+int32(20))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v147 = F_get_object_attnum_acl(m, l1)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v150 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v147+v76-v150))) = uint8(v150)
	v154 = F_get_object_attnum_acl(m, l1)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68+v154<<(uint(int32(2))%32)-int32(4)))) = v141
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v163 = F_heap_modify_tuple(m, v78, v162, v68, v72, v76)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	F_CatalogTupleUpdate(m, v33, v163+int32(4), v163)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	F_UnlockTuple(m, v33, v78+int32(4), int32(7))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, _consts[295])))
	if v175 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	F_updateAclDependencies(m, l1, v63, int32(0), v86, v111, v186, v145, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L52
	}
L47:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, _consts[296])))
	if v179 != int32(1) {
		goto L46
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	F_recordExtensionInitPrivWorker(m, v63, l1, int32(0), v141)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	goto L46
L52:
	;
	F_ReleaseCatCache(m, v78)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	F_pfree(m, v141)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v197 = v55 + int32(1)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v197 < v198 {
		v55 = v197
		goto L11
	} else {
		goto L56
	}
L56:
	;
	goto L12
L57:
	;
	m.G0 = v21 + int32(48)
	return
L58:
	;
	v228 = F_get_object_class_descr(m, l1)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v228
	F_errmsg_internal(m, int32(46411), v21)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(523480), int32(2154), int32(258838))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecMergeAppend(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v177 int32
	_ = v177
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
	var v192 int32
	_ = v192
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
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	v6 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)))
	if v11 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	if v224 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L7:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v14 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+20))
	v192 = v190 << (uint(int32(2)) % 32)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v192+v193)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	if v196 != 0 {
		goto L56
	} else {
		goto L57
	}
L10:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	m.T0[v19].(func(*base.Module, int32))(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v23 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	return v17
L14:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v27 = int32(0)
	v29 = F_ExecFindMatchingSubPlans(m, v26, v27, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	v32 = v23
	goto L16
L16:
	;
	if v32 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v29
	v32 = v29
	goto L16
L18:
	;
	if int32(0) <= v89 {
		goto L29
	} else {
		goto L30
	}
L19:
	;
	v89 = base.I32_ctz(v75) | v76<<(uint(int32(5))%32)
	goto L18
L20:
	;
	v89 = int32(-2)
	goto L18
L21:
	;
	v42 = base.I32_div_s(int32(0), int32(32))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v43 <= v42 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v46 = v32 + int32(8)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v42<<(uint(int32(2))%32))))
	v53 = v50 & int32(-1)
	if v53 != 0 {
		v75 = v53
		v76 = v42
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v55 = v42 + int32(1)
	if v55 == v43 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v58 = v55
	goto L25
L25:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v46+v58<<(uint(int32(2))%32))))
	if v65 != 0 {
		v75 = v65
		v76 = v58
		goto L19
	} else {
		goto L27
	}
L26:
	;
	goto L20
L27:
	;
	v67 = v58 + int32(1)
	if v67 != v43 {
		v58 = v67
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v94 = v89
	goto L32
L30:
	;
	goto L31
L31:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	F_binaryheap_build(m, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L55
	}
L32:
	;
	v97 = v94 << (uint(int32(2)) % 32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97+v98)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+52))
	if v101 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L31
L34:
	;
	F_ExecReScan(m, v100)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	v105 = m.T0[v104].(func(*base.Module, int32) int32)(m, v100)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v107+v97))) = v105
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v110+v97)))
	if v112 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v121 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+4)))
	if v115&int32(2) != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	F_binaryheap_add_unordered(m, v118, v94)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	goto L39
L43:
	;
	if int32(0) <= v177 {
		v94 = v177
		goto L32
	} else {
		goto L54
	}
L44:
	;
	v177 = base.I32_ctz(v163) | v164<<(uint(int32(5))%32)
	goto L43
L45:
	;
	v177 = int32(-2)
	goto L43
L46:
	;
	v128 = v94 + int32(1)
	v130 = base.I32_div_s(v128, int32(32))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	if v131 <= v130 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v134 = v121 + int32(8)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v134+v130<<(uint(int32(2))%32))))
	v141 = v138 & (int32(-1) << (uint(v128) % 32))
	if v141 != 0 {
		v163 = v141
		v164 = v130
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v143 = v130 + int32(1)
	if v143 == v131 {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v146 = v143
	goto L50
L50:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v134+v146<<(uint(int32(2))%32))))
	if v153 != 0 {
		v163 = v153
		v164 = v146
		goto L44
	} else {
		goto L52
	}
L51:
	;
	goto L45
L52:
	;
	v155 = v146 + int32(1)
	if v155 != v131 {
		v146 = v155
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	goto L33
L55:
	;
	v187 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)) = uint8(v187)
	goto L6
L56:
	;
	F_ExecReScan(m, v195)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v195)+12))
	v200 = m.T0[v199].(func(*base.Module, int32) int32)(m, v195)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v202+v192))) = v200
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v205+v192)))
	if v207 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v217 = F_binaryheap_remove_first(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L65
	}
L62:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+4)))
	if v210&int32(2) != 0 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	F_binaryheap_replace_first(m, v213, v190)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	goto L6
L65:
	;
	goto L6
L66:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+8))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+12))
	m.T0[v229].(func(*base.Module, int32))(m, v227)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v223)+20))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v233+v234<<(uint(int32(2))%32))))
	return v238
L69:
	;
	return v227
}
func F_ExecReadyInterpretedExpr(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
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
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _consts[430]))
	if v8 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = int32(0)
	v15 = F_ExecInterpExpr(m, v12, v12, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v94&int32(32) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, _consts[430])) = v15
	v19 = v2
	goto L6
L6:
	;
	v24 = int32(2)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15+v19<<(uint(v24)%32))))
	v28 = int32(3)
	v29 = v19 << (uint(v28) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[431]))) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[432]))) = v27
	v37 = v19 | int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v15+v37<<(uint(v24)%32))))
	v43 = v37 << (uint(v28) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_consts[431]))) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_consts[432]))) = v41
	v51 = v19 | v24
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v15+v51<<(uint(v24)%32))))
	v57 = v51 << (uint(v28) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+uint32(_consts[431]))) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v57)+uint32(_consts[432]))) = v55
	v65 = v19 | v28
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v15+v65<<(uint(v24)%32))))
	v71 = v65 << (uint(v28) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+uint32(_consts[431]))) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v71)+uint32(_consts[432]))) = v69
	v79 = v19 + int32(4)
	if v79 != int32(120) {
		v19 = v79
		goto L6
	} else {
		goto L8
	}
L7:
	;
	F_pg_qsort(m, int32(4459344), int32(120), int32(8), int32(588))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	goto L3
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(589)
	v102 = v94 | int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v102)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	switch v104 - int32(2) {
	case 0:
		goto L18
	case 1:
		goto L19
	case 2:
		goto L20
	case 3:
		goto L21
	default:
		goto L17
	}
L11:
	;
	goto L12
L12:
	;
	return
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v250
	goto L12
L14:
	;
	v250 = int32(610)
	goto L13
L15:
	;
	v245 = v239 | int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v245)
	v250 = int32(609)
	goto L13
L16:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _consts[430]))
	v218 = int32(0)
	goto L66
L17:
	;
	if v104 <= int32(0) {
		v239 = v102
		goto L15
	} else {
		goto L65
	}
L18:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	switch v200 - int32(7) {
	case 0:
		goto L14
	case 1:
		goto L64
	case 2:
		goto L63
	default:
		goto L16
	case 11:
		goto L62
	case 12:
		goto L61
	case 13:
		goto L60
	case 18:
		v250 = int32(603)
		goto L13
	}
L19:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+40))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v150 != int32(2) {
		goto L34
	} else {
		goto L35
	}
L20:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+80))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v121)+40))
	v129 = base.B2i32(v123 == int32(3)) & base.B2i32(v126 == int32(8))
	if v129 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v108 != int32(2) {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107)+40))
	if v111 != int32(85) {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v107)+80))
	if v114 != int32(7) {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v107)+120))
	if v117 != int32(88) {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	v250 = int32(590)
	goto L13
L26:
	;
	if v123 != int32(2) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	if v122 != int32(86) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v250 = int32(591)
	goto L13
L29:
	;
	if v129&base.B2i32(v122 == int32(87)) == int32(0) {
		goto L16
	} else {
		goto L33
	}
L30:
	;
	if v126 != int32(7) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	if v122 != int32(86) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v250 = int32(592)
	goto L13
L33:
	;
	v250 = int32(593)
	goto L13
L34:
	;
	if v150 != int32(3) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	if v149 != int32(7) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v250 = int32(594)
	goto L13
L37:
	;
	if v150 != int32(4) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	if v149 != int32(8) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v250 = int32(595)
	goto L13
L40:
	;
	if v150 != int32(2) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	if v149 != int32(9) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v250 = int32(596)
	goto L13
L43:
	;
	if v150 != int32(3) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	if v149 != int32(18) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v250 = int32(597)
	goto L13
L46:
	;
	if v150 != int32(4) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	if v149 != int32(19) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v250 = int32(598)
	goto L13
L49:
	;
	if v150 != int32(56) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	if v149 != int32(20) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v250 = int32(599)
	goto L13
L52:
	;
	if v150 != int32(7) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v149-int32(27)) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v250 = int32(600)
	goto L13
L55:
	;
	if v150 != int32(8) {
		goto L16
	} else {
		goto L58
	}
L56:
	;
	if v149 != int32(86) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v250 = int32(601)
	goto L13
L58:
	;
	if v149 != int32(86) {
		goto L16
	} else {
		goto L59
	}
L59:
	;
	v250 = int32(602)
	goto L13
L60:
	;
	v250 = int32(608)
	goto L13
L61:
	;
	v250 = int32(607)
	goto L13
L62:
	;
	v250 = int32(606)
	goto L13
L63:
	;
	v250 = int32(605)
	goto L13
L64:
	;
	v250 = int32(604)
	goto L13
L65:
	;
	goto L16
L66:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v226 = v223 + v218*int32(40)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v216+v227<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v226))) = v231
	v234 = v218 + int32(1)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v234 < v235 {
		v218 = v234
		goto L66
	} else {
		goto L68
	}
L67:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	v239 = v237
	goto L15
L68:
	;
	goto L67
}
func F_ExecRenameStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
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
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
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
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
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
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
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
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
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
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v871 int32
	_ = v871
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v925 int32
	_ = v925
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v979 int32
	_ = v979
	var v981 int64
	_ = v981
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
	var v1035 int64
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1123 int32
	_ = v1123
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1141 int32
	_ = v1141
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1163 int32
	_ = v1163
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1217 int32
	_ = v1217
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1259 int32
	_ = v1259
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1325 int32
	_ = v1325
	var v1330 int32
	_ = v1330
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1359 int32
	_ = v1359
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1394 int32
	_ = v1394
	var v1396 int32
	_ = v1396
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1447 int32
	_ = v1447
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1463 int32
	_ = v1463
	var v1470 int32
	_ = v1470
	var v1472 int32
	_ = v1472
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
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
	var v1504 int32
	_ = v1504
	var v1508 int32
	_ = v1508
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1520 int32
	_ = v1520
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1567 int32
	_ = v1567
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1578 int32
	_ = v1578
	var v1583 int32
	_ = v1583
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1595 int32
	_ = v1595
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1613 int32
	_ = v1613
	var v1620 int32
	_ = v1620
	var v1624 int32
	_ = v1624
	var v1629 int32
	_ = v1629
	var v1633 int32
	_ = v1633
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1644 int32
	_ = v1644
	var v1651 int32
	_ = v1651
	var v1656 int32
	_ = v1656
	var v1660 int32
	_ = v1660
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1680 int32
	_ = v1680
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
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
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1713 int32
	_ = v1713
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1720 int32
	_ = v1720
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1732 int32
	_ = v1732
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1752 int32
	_ = v1752
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1778 int32
	_ = v1778
	var v1780 int32
	_ = v1780
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1800 int32
	_ = v1800
	var v1803 int32
	_ = v1803
	var v1807 int32
	_ = v1807
	var v1811 int32
	_ = v1811
	var v1816 int32
	_ = v1816
	var v1834 int32
	_ = v1834
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1846 int32
	_ = v1846
	var v1848 int32
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1856 int32
	_ = v1856
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1893 int32
	_ = v1893
	var v1897 int32
	_ = v1897
	var v1900 int32
	_ = v1900
	var v1905 int32
	_ = v1905
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1918 int32
	_ = v1918
	var v1920 int32
	_ = v1920
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1929 int32
	_ = v1929
	var v1934 int32
	_ = v1934
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1970 int32
	_ = v1970
	var v1971 int32
	_ = v1971
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1983 int32
	_ = v1983
	var v1988 int32
	_ = v1988
	var v1990 int64
	_ = v1990
	var v1993 int32
	_ = v1993
	var v1998 int32
	_ = v1998
	var v2002 int32
	_ = v2002
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2017 int32
	_ = v2017
	var v2022 int32
	_ = v2022
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2044 int32
	_ = v2044
	var v2049 int32
	_ = v2049
	var v2053 int32
	_ = v2053
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2064 int32
	_ = v2064
	var v2069 int32
	_ = v2069
	var v2073 int32
	_ = v2073
	var v2076 int32
	_ = v2076
	var v2082 int32
	_ = v2082
	var v2087 int32
	_ = v2087
	var v2092 int32
	_ = v2092
	var v2098 int32
	_ = v2098
	var v2103 int32
	_ = v2103
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2109 int32
	_ = v2109
	var v2111 int32
	_ = v2111
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2158 int32
	_ = v2158
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2168 int32
	_ = v2168
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2178 int32
	_ = v2178
	var v2186 int32
	_ = v2186
	var v2194 int32
	_ = v2194
	var v2198 int32
	_ = v2198
	var v2200 int32
	_ = v2200
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2242 int32
	_ = v2242
	var v2244 int32
	_ = v2244
	var v2246 int32
	_ = v2246
	var v2248 int32
	_ = v2248
	var v2250 int32
	_ = v2250
	var v2253 int32
	_ = v2253
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(160)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v17 - int32(1) {
	case 0, 6, 7, 13, 15, 16, 18, 20, 23, 25, 28, 29, 33, 37, 38, 44, 45, 46, 47:
		goto L10
	default:
		goto L9
	case 3, 5:
		goto L15
	case 8:
		goto L20
	case 11, 48:
		goto L11
	case 12, 39:
		goto L8
	case 17, 19, 22, 36, 40, 50:
		goto L16
	case 27:
		goto L12
	case 32:
		goto L19
	case 34:
		goto L14
	case 35:
		goto L18
	case 41:
		goto L17
	case 43:
		goto L13
	}
L1:
	;
	m.G0 = v15 + int32(160)
	return
L2:
	;
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(v1694)+48))
	v2115 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2114)+120)))
	v2118 = F_palloc0(m, v2115<<(uint(int32(2))%32))
	mBase = m.M
	v2119 = m.ExcPending
	if v2119 != 0 {
		goto L21
	} else {
		goto L648
	}
L3:
	;
	v2105 = *(*int32)(unsafe.Add(mBase, uint32(v1709)+16))
	v2106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2105)+22)))
	v2107 = v2105 + v2106
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(v2107)+4))
	v2109 = *(*int32)(unsafe.Add(mBase, uint32(v2107)+72))
	F_IsThereOpClassInNamespace(m, v1696, v2108, v2109)
	mBase = m.M
	v2111 = m.ExcPending
	if v2111 != 0 {
		goto L21
	} else {
		goto L647
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2092 = m.ExcPending
	if v2092 != 0 {
		goto L21
	} else {
		goto L644
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2073 = m.ExcPending
	if v2073 != 0 {
		goto L21
	} else {
		goto L640
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2053 = m.ExcPending
	if v2053 != 0 {
		goto L21
	} else {
		goto L635
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L21
	} else {
		goto L632
	}
L8:
	;
	v1936 = m.G0
	v1938 = v1936 - int32(32)
	m.G0 = v1938
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1940 == int32(13) {
		goto L606
	} else {
		goto L607
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1924 = m.ExcPending
	if v1924 != 0 {
		goto L21
	} else {
		goto L600
	}
L10:
	;
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1687 = int32(0)
	F_get_object_address(m, l0, v17, v1686, v1687, int32(8), v1687)
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L21
	} else {
		goto L502
	}
L11:
	;
	v1526 = m.G0
	v1528 = v1526 - int32(96)
	m.G0 = v1528
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1533 = F_makeTypeNameFromNameList(m, v1532)
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L21
	} else {
		goto L447
	}
L12:
	;
	v1365 = m.G0
	v1367 = v1365 - int32(128)
	m.G0 = v1367
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1373 = F_RangeVarGetRelidExtended(m, v1369, int32(8), int32(0), int32(566), l1)
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L21
	} else {
		goto L410
	}
L13:
	;
	v1185 = m.G0
	v1187 = v1185 - int32(144)
	m.G0 = v1187
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1191 = int32(0)
	v1194 = F_RangeVarGetRelidExtended(m, v1189, int32(8), v1191, int32(580), v1191)
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L21
	} else {
		goto L370
	}
L14:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v1059 = m.G0
	v1061 = v1059 - int32(32)
	m.G0 = v1061
	v1064 = int32(0)
	v1067 = F_RangeVarGetRelidExtended(m, v1056, int32(8), v1064, int32(1040), v1064)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L21
	} else {
		goto L335
	}
L15:
	;
	v1006 = m.G0
	v1008 = v1006 - int32(16)
	m.G0 = v1008
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	v1015 = F_RangeVarGetRelidExtended(m, v1010, int32(8), v1012, int32(575), int32(0))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L21
	} else {
		goto L324
	}
L16:
	;
	v931 = m.G0
	v933 = v931 - int32(16)
	m.G0 = v933
	v935 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v940 = base.B2i32(v938 == int32(20))
	if v938 == int32(20) {
		goto L302
	} else {
		goto L303
	}
L17:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v768 = m.G0
	v770 = v768 - int32(96)
	m.G0 = v770
	v774 = F_table_open(m, int32(1213), int32(3))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L21
	} else {
		goto L249
	}
L18:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v619 = m.G0
	v621 = v619 - int32(48)
	m.G0 = v621
	v625 = F_table_open(m, int32(2615), int32(3))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L21
	} else {
		goto L199
	}
L19:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v246 = m.G0
	v248 = v246 - int32(208)
	m.G0 = v248
	v252 = F_table_open(m, int32(1260), int32(3))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L21
	} else {
		goto L85
	}
L20:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v22 = m.G0
	v24 = v22 - int32(80)
	m.G0 = v24
	v28 = F_table_open(m, int32(1262), int32(3))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return
L22:
	;
	v33 = int32(0)
	v47 = F_get_db_info(m, v20, int32(8), v24+int32(76), v33, v33, v33, v33, v33, v33, v33, v33, v33, v33, v33, v33, v33, v33)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L21
	} else {
		goto L29
	}
L23:
	;
	goto L1
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L21
	} else {
		goto L82
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L21
	} else {
		goto L77
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L21
	} else {
		goto L73
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L21
	} else {
		goto L69
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L21
	} else {
		goto L65
	}
L29:
	;
	if v47 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v24)+76))
	v52 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v53 = F_object_ownercheck(m, int32(1262), v50, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L21
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
	v143 = m.ExcPending
	if v143 != 0 {
		goto L21
	} else {
		goto L61
	}
L33:
	;
	if v53 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_aclcheck_error(m, int32(2), int32(9), v20)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L21
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v61 = F_superuser(m)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L21
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	if v61 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v68 = F_SearchSysCache1(m, int32(11), v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L21
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v83 = F_get_database_oid(m, v21, int32(1))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L21
	} else {
		goto L46
	}
L42:
	;
	if v68 == int32(0) {
		goto L28
	} else {
		goto L43
	}
L43:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+22)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v73)+71)))
	F_ReleaseCatCache(m, v68)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L21
	} else {
		goto L44
	}
L44:
	;
	if v75 == int32(0) {
		goto L28
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	if v83 != 0 {
		goto L27
	} else {
		goto L47
	}
L47:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	if v50 == v86 {
		goto L26
	} else {
		goto L48
	}
L48:
	;
	v92 = F_CountOtherDBBackends(m, v50, v24-int32(-64), v24+int32(60))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L21
	} else {
		goto L49
	}
L49:
	;
	if v92 != 0 {
		goto L25
	} else {
		goto L50
	}
L50:
	;
	v95 = F_SearchSysCacheLockedCopy1(m, int32(21), v50)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L21
	} else {
		goto L51
	}
L51:
	;
	if v95 == int32(0) {
		goto L24
	} else {
		goto L52
	}
L52:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+72)) = uint16(v99)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+22)))
	v109 = F_strncpy(m, v103+v104+int32(4), v21, int32(64))
	mBase = m.M
	v110 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+63)) = uint8(v110)
	goto L53
L53:
	;
	F_CatalogTupleUpdate(m, v28, v24+int32(68), v95)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L21
	} else {
		goto L54
	}
L54:
	;
	F_UnlockTuple(m, v28, v24+int32(68), int32(7))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L21
	} else {
		goto L55
	}
L55:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _consts[297]))
	if v122 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v124 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1262), v50, v124, v124, v124)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L21
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v129 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
	F_sequence_close(m, v28, v129)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L21
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	m.G0 = v24 + int32(80)
	goto L23
L61:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L21
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v20
	F_errmsg(m, int32(78313), v24+int32(48))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L21
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(519377), int32(1922), int32(381343))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L21
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L21
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(381058), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L21
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(519377), int32(1933), int32(381343))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L21
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
	F_errcode(m, int32(67240068))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L21
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v21
	F_errmsg(m, int32(124975), v24+int32(32))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L21
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(519377), int32(1951), int32(381343))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L21
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L21
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(476145), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L21
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(519377), int32(1962), int32(381343))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L21
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_errcode(m, int32(100663621))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L21
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v20
	F_errmsg(m, int32(142501), v24)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L21
	} else {
		goto L79
	}
L79:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
	F_errdetail_busy_db(m, v220, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L21
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(519377), int32(1975), int32(381343))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L21
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v50
	F_errmsg_internal(m, int32(54369), v24+int32(16))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L21
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(519377), int32(1980), int32(381343))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L21
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v252)+52))
	v256 = F_SearchSysCache1(m, int32(10), v244)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L21
	} else {
		goto L93
	}
L86:
	;
	goto L1
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L21
	} else {
		goto L194
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L21
	} else {
		goto L190
	}
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L21
	} else {
		goto L185
	}
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L21
	} else {
		goto L180
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L21
	} else {
		goto L176
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L21
	} else {
		goto L172
	}
L93:
	;
	if v256 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v256)+16))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+22)))
	v260 = v258 + v259
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v263 = *(*int32)(unsafe.Add(mBase, _consts[347]))
	if v261 == v263 {
		goto L92
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L21
	} else {
		goto L168
	}
L97:
	;
	v266 = *(*int32)(unsafe.Add(mBase, _consts[348]))
	if v261 == v266 {
		goto L91
	} else {
		goto L98
	}
L98:
	;
	v269 = v260 + int32(4)
	v270 = int32(0)
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	if v271 != int32(112) {
		v280 = v270
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if v280 != 0 {
		goto L90
	} else {
		goto L103
	}
L100:
	;
	goto L99
L101:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+1)))
	if v274 != int32(103) {
		v280 = v270
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+2)))
	v280 = base.B2i32(v277 == int32(95))
	goto L100
L103:
	;
	v281 = int32(0)
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	if v282 != int32(112) {
		v291 = v281
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v291 != 0 {
		goto L89
	} else {
		goto L108
	}
L105:
	;
	goto L104
L106:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+1)))
	if v285 != int32(103) {
		v291 = v281
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+2)))
	v291 = base.B2i32(v288 == int32(95))
	goto L105
L108:
	;
	v293 = int32(0)
	v296 = F_SearchSysCacheExists(m, int32(10), v245, v293, v293, v293)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L21
	} else {
		goto L109
	}
L109:
	;
	if v296 != 0 {
		goto L88
	} else {
		goto L110
	}
L110:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+68)))
	if v298 == int32(1) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v340 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v248)+128)) = v340
	*(*int64)(unsafe.Add(mBase, uint32(v248)+120)) = int64(0)
	v344 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v248)+121)) = uint8(v344)
	v348 = F_DirectFunctionCall1Coll(m, int32(500), v340, v245)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L21
	} else {
		goto L126
	}
L112:
	;
	v301 = F_superuser(m)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L21
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v330 = F_has_createrole_privilege(m, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L21
	} else {
		goto L122
	}
L115:
	;
	if v301 != 0 {
		goto L111
	} else {
		goto L116
	}
L116:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L21
	} else {
		goto L117
	}
L117:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L21
	} else {
		goto L118
	}
L118:
	;
	F_errmsg(m, int32(405206), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L21
	} else {
		goto L119
	}
L119:
	;
	v314 = int32(552379)
	*(*int32)(unsafe.Add(mBase, uint32(v248)+84)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v248)+80)) = v314
	F_errdetail(m, int32(660296), v248+int32(80))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L21
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(520594), int32(1423), int32(405372))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L21
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
	if v330 == int32(0) {
		goto L87
	} else {
		goto L123
	}
L123:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v336 = F_is_admin_of_role(m, v335, v261)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L21
	} else {
		goto L124
	}
L124:
	;
	if v336 == int32(0) {
		goto L87
	} else {
		goto L125
	}
L125:
	;
	goto L111
L126:
	;
	v350 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v248)+133)) = uint8(v350)
	*(*int32)(unsafe.Add(mBase, uint32(v248)+148)) = v348
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v256)+16))
	v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v353)+18)))
	if base.Ui32(int32(11)) <= base.Ui32(v354&int32(2047)) {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v456 = F_heap_modify_tuple(m, v256, v254, v248+int32(144), v248+int32(132), v248+int32(120))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L21
	} else {
		goto L160
	}
L128:
	;
	v421 = F_text_to_cstring(m, v419)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L21
	} else {
		goto L153
	}
L129:
	;
	v359 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v248)+207)) = uint8(v359)
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353)+20)))
	if v361&int32(1) == v359 {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	goto L131
L131:
	;
	v414 = F_getmissingattr(m, v254, int32(11), v248+int32(207))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L21
	} else {
		goto L151
	}
L132:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v254)+180))
	if int32(0) <= v366 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	goto L134
L134:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353)+24)))
	if v401&int32(4) == int32(0) {
		goto L147
	} else {
		goto L148
	}
L135:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353)+22)))
	v371 = v353 + v369 + v366
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+186)))
	if v372 != int32(1) {
		v419 = v371
		goto L128
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v399 = F_nocachegetattr(m, v256, int32(11), v254)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L21
	} else {
		goto L146
	}
L138:
	;
	v375 = int32(*(*int16)(unsafe.Add(mBase, uint32(v254)+184)))
	switch v375&int32(65535) - int32(1) {
	case 0:
		goto L142
	case 1:
		goto L141
	default:
		goto L139
	case 3:
		goto L140
	}
L139:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L21
	} else {
		goto L143
	}
L140:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v371)))
	v419 = v382
	goto L128
L141:
	;
	v381 = int32(*(*int16)(unsafe.Add(mBase, uint32(v371))))
	v419 = v381
	goto L128
L142:
	;
	v380 = int32(*(*int8)(unsafe.Add(mBase, uint32(v371))))
	v419 = v380
	goto L128
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248)+64)) = v375
	F_errmsg_internal(m, int32(507173), v248-int32(-64))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L21
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(343882), int32(70), int32(73857))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L21
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	v419 = v399
	goto L128
L147:
	;
	v406 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v248)+207)) = uint8(v406)
	goto L127
L148:
	;
	goto L149
L149:
	;
	v409 = F_nocachegetattr(m, v256, int32(11), v254)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L21
	} else {
		goto L150
	}
L150:
	;
	v419 = v409
	goto L128
L151:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+207)))
	if v416&int32(1) != 0 {
		goto L127
	} else {
		goto L152
	}
L152:
	;
	v419 = v414
	goto L128
L153:
	;
	v423 = F_get_password_type(m, v421)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L21
	} else {
		goto L154
	}
L154:
	;
	if v423 != int32(1) {
		goto L127
	} else {
		goto L155
	}
L155:
	;
	v427 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v248)+142)) = uint8(v427)
	*(*uint8)(unsafe.Add(mBase, uint32(v248)+130)) = uint8(v427)
	v433 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L21
	} else {
		goto L156
	}
L156:
	;
	if v433 == int32(0) {
		goto L127
	} else {
		goto L157
	}
L157:
	;
	F_errmsg(m, int32(397202), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L21
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(520594), int32(1454), int32(405372))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L21
	} else {
		goto L159
	}
L159:
	;
	goto L127
L160:
	;
	F_CatalogTupleUpdate(m, v252, v256+int32(4), v456)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L21
	} else {
		goto L161
	}
L161:
	;
	v461 = *(*int32)(unsafe.Add(mBase, _consts[297]))
	if v461 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v463 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1260), v261, v463, v463, v463)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L21
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v261
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1260)
	F_ReleaseCatCache(m, v256)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L21
	} else {
		goto L166
	}
L165:
	;
	goto L164
L166:
	;
	F_sequence_close(m, v252, int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L21
	} else {
		goto L167
	}
L167:
	;
	m.G0 = v248 + int32(208)
	goto L86
L168:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L21
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248))) = v244
	F_errmsg(m, int32(78385), v248)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L21
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(520594), int32(1357), int32(405372))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L21
	} else {
		goto L171
	}
L171:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L172:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L21
	} else {
		goto L173
	}
L173:
	;
	F_errmsg(m, int32(476114), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L21
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(520594), int32(1373), int32(405372))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L21
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L21
	} else {
		goto L177
	}
L177:
	;
	F_errmsg(m, int32(476083), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L21
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(520594), int32(1377), int32(405372))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L21
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
	F_errcode(m, int32(151818372))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L21
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248)+16)) = v269
	F_errmsg(m, int32(462194), v248+int32(16))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L21
	} else {
		goto L182
	}
L182:
	;
	F_errdetail(m, int32(677674), int32(0))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L21
	} else {
		goto L183
	}
L183:
	;
	F_errfinish(m, int32(520594), int32(1388), int32(405372))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L21
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
	F_errcode(m, int32(151818372))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L21
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248)+32)) = v245
	F_errmsg(m, int32(462194), v248+int32(32))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L21
	} else {
		goto L187
	}
L187:
	;
	F_errdetail(m, int32(677674), int32(0))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L21
	} else {
		goto L188
	}
L188:
	;
	F_errfinish(m, int32(520594), int32(1395), int32(405372))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L21
	} else {
		goto L189
	}
L189:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L190:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L21
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248)+48)) = v245
	F_errmsg(m, int32(125029), v248+int32(48))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L21
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(520594), int32(1410), int32(405372))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L21
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
	F_errcode(m, int32(16797828))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L21
	} else {
		goto L195
	}
L195:
	;
	F_errmsg(m, int32(405206), int32(0))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L21
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248)+104)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v248)+100)) = int32(557325)
	*(*int32)(unsafe.Add(mBase, uint32(v248)+96)) = int32(567348)
	F_errdetail(m, int32(667619), v248+int32(96))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L21
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(520594), int32(1433), int32(405372))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L21
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	v629 = F_SearchSysCacheCopy(m, int32(37), v617, int32(0))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L21
	} else {
		goto L203
	}
L200:
	;
	goto L1
L201:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L21
	} else {
		goto L244
	}
L202:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L21
	} else {
		goto L240
	}
L203:
	;
	if v629 != 0 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v629)+16))
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v631)+22)))
	v633 = v631 + v632
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v633)))
	v636 = F_get_namespace_oid(m, v618, int32(1))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L21
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L21
	} else {
		goto L236
	}
L207:
	;
	if v636 != 0 {
		goto L202
	} else {
		goto L208
	}
L208:
	;
	v640 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v641 = F_object_ownercheck(m, int32(2615), v634, v640)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L21
	} else {
		goto L209
	}
L209:
	;
	if v641 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	F_aclcheck_error(m, int32(2), int32(36), v617)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L21
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v651 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v653 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v655 = F_object_aclcheck(m, int32(1262), v651, v653, int64(512))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L21
	} else {
		goto L214
	}
L213:
	;
	goto L212
L214:
	;
	if v655 != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v659 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v660 = F_get_database_name(m, v659)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L21
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v665 = int32(*(*uint8)(unsafe.Add(mBase, _consts[349])))
	if v665 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L218:
	;
	F_aclcheck_error(m, v655, int32(9), v660)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L21
	} else {
		goto L219
	}
L219:
	;
	goto L217
L220:
	;
	v668 = int32(0)
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618))))
	if v669 != int32(112) {
		v678 = v668
		goto L224
	} else {
		goto L225
	}
L221:
	;
	goto L222
L222:
	;
	v682 = F_strncpy(m, v633+int32(4), v618, int32(64))
	mBase = m.M
	v683 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v682)+63)) = uint8(v683)
	goto L228
L223:
	;
	if v678 != 0 {
		goto L201
	} else {
		goto L227
	}
L224:
	;
	goto L223
L225:
	;
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618)+1)))
	if v672 != int32(103) {
		v678 = v668
		goto L224
	} else {
		goto L226
	}
L226:
	;
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618)+2)))
	v678 = base.B2i32(v675 == int32(95))
	goto L224
L227:
	;
	goto L222
L228:
	;
	F_CatalogTupleUpdate(m, v625, v629+int32(4), v629)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L21
	} else {
		goto L229
	}
L229:
	;
	v690 = *(*int32)(unsafe.Add(mBase, _consts[297]))
	if v690 != 0 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v692 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2615), v634, v692, v692, v692)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L21
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v697 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v697
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v634
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2615)
	F_sequence_close(m, v625, v697)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L21
	} else {
		goto L234
	}
L233:
	;
	goto L232
L234:
	;
	F_pfree(m, v629)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L21
	} else {
		goto L235
	}
L235:
	;
	m.G0 = v621 + int32(48)
	goto L200
L236:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L21
	} else {
		goto L237
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v621))) = v617
	F_errmsg(m, int32(78796), v621)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L21
	} else {
		goto L238
	}
L238:
	;
	F_errfinish(m, int32(519639), int32(264), int32(533469))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L21
	} else {
		goto L239
	}
L239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L240:
	;
	F_errcode(m, int32(100794500))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L21
	} else {
		goto L241
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v621)+32)) = v618
	F_errmsg(m, int32(125325), v621+int32(32))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L21
	} else {
		goto L242
	}
L242:
	;
	F_errfinish(m, int32(519639), int32(273), int32(533469))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L21
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
	F_errcode(m, int32(151818372))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L21
	} else {
		goto L245
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v621)+16)) = v618
	F_errmsg(m, int32(750892), v621+int32(16))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L21
	} else {
		goto L246
	}
L246:
	;
	F_errdetail(m, int32(632478), int32(0))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L21
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(519639), int32(290), int32(533469))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L21
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L249:
	;
	F_ScanKeyInit(m, v770+int32(48), int32(2), int32(3), int32(62), v766)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L21
	} else {
		goto L250
	}
L250:
	;
	v786 = F_table_beginscan_catalog(m, v774, int32(1), v770+int32(48))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L21
	} else {
		goto L254
	}
L251:
	;
	goto L1
L252:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L21
	} else {
		goto L295
	}
L253:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L21
	} else {
		goto L290
	}
L254:
	;
	v788 = F_heap_getnext(m, v786)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L21
	} else {
		goto L255
	}
L255:
	;
	if v788 != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v790 = F_heap_copytuple(m, v788)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L21
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L21
	} else {
		goto L286
	}
L259:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v790)+16))
	v793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v792)+22)))
	v794 = v792 + v793
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v794)))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v786)))
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v796)+188))
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v797)+12))
	m.T0[v798].(func(*base.Module, int32))(m, v786)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L21
	} else {
		goto L260
	}
L260:
	;
	v803 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v804 = F_object_ownercheck(m, int32(1213), v795, v803)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L21
	} else {
		goto L261
	}
L261:
	;
	if v804 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	F_aclcheck_error(m, int32(1), int32(42), v766)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L21
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	v813 = int32(*(*uint8)(unsafe.Add(mBase, _consts[349])))
	if v813 == int32(0) {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	goto L264
L266:
	;
	v816 = int32(0)
	v817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767))))
	if v817 != int32(112) {
		v826 = v816
		goto L270
	} else {
		goto L271
	}
L267:
	;
	goto L268
L268:
	;
	F_ScanKeyInit(m, v770+int32(48), int32(2), int32(3), int32(62), v767)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L21
	} else {
		goto L274
	}
L269:
	;
	if v826 != 0 {
		goto L253
	} else {
		goto L273
	}
L270:
	;
	goto L269
L271:
	;
	v820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767)+1)))
	if v820 != int32(103) {
		v826 = v816
		goto L270
	} else {
		goto L272
	}
L272:
	;
	v823 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767)+2)))
	v826 = base.B2i32(v823 == int32(95))
	goto L270
L273:
	;
	goto L268
L274:
	;
	v837 = F_table_beginscan_catalog(m, v774, int32(1), v770+int32(48))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L21
	} else {
		goto L275
	}
L275:
	;
	v839 = F_heap_getnext(m, v837)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L21
	} else {
		goto L276
	}
L276:
	;
	if v839 != 0 {
		goto L252
	} else {
		goto L277
	}
L277:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v837)))
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v841)+188))
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v842)+12))
	m.T0[v843].(func(*base.Module, int32))(m, v837)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L21
	} else {
		goto L278
	}
L278:
	;
	v849 = F_strncpy(m, v794+int32(4), v767, int32(64))
	mBase = m.M
	v850 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v849)+63)) = uint8(v850)
	goto L279
L279:
	;
	F_CatalogTupleUpdate(m, v774, v790+int32(4), v790)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L21
	} else {
		goto L280
	}
L280:
	;
	v857 = *(*int32)(unsafe.Add(mBase, _consts[297]))
	if v857 != 0 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v859 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1213), v795, v859, v859, v859)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L21
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	v864 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v864
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v795
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1213)
	F_sequence_close(m, v774, v864)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L21
	} else {
		goto L285
	}
L284:
	;
	goto L283
L285:
	;
	m.G0 = v770 + int32(96)
	goto L251
L286:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L21
	} else {
		goto L287
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v770))) = v766
	F_errmsg(m, int32(78725), v770)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L21
	} else {
		goto L288
	}
L288:
	;
	F_errfinish(m, int32(525439), int32(954), int32(441001))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L21
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
	F_errcode(m, int32(151818372))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L21
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v770)+32)) = v767
	F_errmsg(m, int32(750858), v770+int32(32))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L21
	} else {
		goto L292
	}
L292:
	;
	F_errdetail(m, int32(631903), int32(0))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L21
	} else {
		goto L293
	}
L293:
	;
	F_errfinish(m, int32(525439), int32(971), int32(441001))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L21
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
	F_errcode(m, int32(290948))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L21
	} else {
		goto L296
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v770)+16)) = v767
	F_errmsg(m, int32(125177), v770+int32(16))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L21
	} else {
		goto L297
	}
L297:
	;
	F_errfinish(m, int32(525439), int32(993), int32(441001))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L21
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
	m.G0 = v933 + int32(16)
	goto L1
L300:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	F_RenameRelationInternal(m, v989, v991, int32(0), v990)
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L21
	} else {
		goto L322
	}
L301:
	;
	v986 = F_get_rel_relkind(m, v961)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L21
	} else {
		goto L321
	}
L302:
	;
	v941 = int32(4)
	goto L304
L303:
	;
	v941 = int32(8)
	goto L304
L304:
	;
	v942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	v944 = F_RangeVarGetRelidExtended(m, v935, v941, v942, int32(576), l1)
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L21
	} else {
		goto L305
	}
L305:
	;
	if v944 != 0 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v946 = F_get_rel_relkind(m, v944)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L21
	} else {
		goto L309
	}
L307:
	;
	goto L308
L308:
	;
	v967 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L21
	} else {
		goto L315
	}
L309:
	;
	if v938 != int32(20) {
		v989 = v944
		v990 = v940
		goto L300
	} else {
		goto L310
	}
L310:
	;
	if v946&int32(223) == int32(73) {
		v989 = v944
		v990 = v940
		goto L300
	} else {
		goto L311
	}
L311:
	;
	F_UnlockRelationOid(m, v944, int32(4))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L21
	} else {
		goto L312
	}
L312:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	v961 = F_RangeVarGetRelidExtended(m, v957, int32(8), v959, int32(576), l1)
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L21
	} else {
		goto L313
	}
L313:
	;
	if v961 != 0 {
		goto L301
	} else {
		goto L314
	}
L314:
	;
	goto L308
L315:
	;
	if v967 != 0 {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v969)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v933))) = v970
	F_errmsg(m, int32(350183), v933)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L21
	} else {
		goto L319
	}
L317:
	;
	goto L318
L318:
	;
	v981 = *(*int64)(unsafe.Add(mBase, _consts[350]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v981
	v984 = *(*int32)(unsafe.Add(mBase, _consts[351]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v984
	goto L299
L319:
	;
	F_errfinish(m, int32(519627), int32(4238), int32(277835))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L21
	} else {
		goto L320
	}
L320:
	;
	goto L318
L321:
	;
	v989 = v961
	v990 = int32(0)
	goto L300
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
	goto L299
L323:
	;
	m.G0 = v1008 + int32(16)
	goto L1
L324:
	;
	if v1015 == int32(0) {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v1021 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L21
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1042)+16)))
	v1044 = int32(0)
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v1047 = F_renameatt_internal(m, v1015, v1040, v1041, v1043, v1044, v1044, v1046)
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L21
	} else {
		goto L334
	}
L328:
	;
	if v1021 != 0 {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1008))) = v1024
	F_errmsg(m, int32(350183), v1008)
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L21
	} else {
		goto L332
	}
L330:
	;
	goto L331
L331:
	;
	v1035 = *(*int64)(unsafe.Add(mBase, _consts[350]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v1035
	v1038 = *(*int32)(unsafe.Add(mBase, _consts[351]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1038
	goto L323
L332:
	;
	F_errfinish(m, int32(519627), int32(4025), int32(73847))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L21
	} else {
		goto L333
	}
L333:
	;
	goto L331
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1047
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1015
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
	goto L323
L335:
	;
	v1070 = F_relation_open(m, v1067, int32(0))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L21
	} else {
		goto L336
	}
L336:
	;
	v1074 = F_table_open(m, int32(2618), int32(3))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L21
	} else {
		goto L337
	}
L337:
	;
	v1077 = F_SearchSysCacheCopy(m, int32(60), v1067, v1057)
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L21
	} else {
		goto L341
	}
L338:
	;
	goto L1
L339:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L21
	} else {
		goto L366
	}
L340:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L21
	} else {
		goto L362
	}
L341:
	;
	if v1077 != 0 {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+16))
	v1080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1079)+22)))
	v1081 = v1079 + v1080
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1081)))
	v1084 = int32(0)
	v1086 = F_SearchSysCacheExists(m, int32(60), v1067, v1058, v1084, v1084)
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L21
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L21
	} else {
		goto L358
	}
L345:
	;
	if v1086 != 0 {
		goto L340
	} else {
		goto L346
	}
L346:
	;
	v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1081)+72)))
	if v1088 == int32(49) {
		goto L339
	} else {
		goto L347
	}
L347:
	;
	v1094 = F_strncpy(m, v1081+int32(4), v1058, int32(64))
	mBase = m.M
	v1095 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1094)+63)) = uint8(v1095)
	goto L348
L348:
	;
	F_CatalogTupleUpdate(m, v1074, v1077+int32(4), v1077)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L21
	} else {
		goto L349
	}
L349:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, _consts[297]))
	if v1102 != 0 {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v1104 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2618), v1082, v1104, v1104, v1104)
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L21
	} else {
		goto L353
	}
L351:
	;
	goto L352
L352:
	;
	F_pfree(m, v1077)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L21
	} else {
		goto L354
	}
L353:
	;
	goto L352
L354:
	;
	F_sequence_close(m, v1074, int32(3))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L21
	} else {
		goto L355
	}
L355:
	;
	F_CacheInvalidateRelcache(m, v1070)
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L21
	} else {
		goto L356
	}
L356:
	;
	v1116 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1116
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1082
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2618)
	F_relation_close(m, v1070, v1116)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L21
	} else {
		goto L357
	}
L357:
	;
	m.G0 = v1061 + int32(32)
	goto L338
L358:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L21
	} else {
		goto L359
	}
L359:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1061))) = v1057
	*(*int32)(unsafe.Add(mBase, uint32(v1061)+4)) = v1134 + int32(4)
	F_errmsg(m, int32(77448), v1061)
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L21
	} else {
		goto L360
	}
L360:
	;
	F_errfinish(m, int32(524724), int32(827), int32(402321))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L21
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
	F_errcode(m, int32(290948))
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L21
	} else {
		goto L363
	}
L363:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1061)+16)) = v1058
	*(*int32)(unsafe.Add(mBase, uint32(v1061)+20)) = v1154 + int32(4)
	F_errmsg(m, int32(124669), v1061+int32(16))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L21
	} else {
		goto L364
	}
L364:
	;
	F_errfinish(m, int32(524724), int32(836), int32(402321))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L21
	} else {
		goto L365
	}
L365:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L366:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L21
	} else {
		goto L367
	}
L367:
	;
	F_errmsg(m, int32(461134), int32(0))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L21
	} else {
		goto L368
	}
L368:
	;
	F_errfinish(m, int32(524724), int32(845), int32(402321))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L21
	} else {
		goto L369
	}
L369:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L370:
	;
	v1197 = F_relation_open(m, v1194, int32(0))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L21
	} else {
		goto L371
	}
L371:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1197)+48))
	v1200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1199)+119)))
	if v1200 == int32(112) {
		goto L372
	} else {
		goto L373
	}
L372:
	;
	v1205 = F_find_all_inheritors(m, v1194, int32(8), int32(0))
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L21
	} else {
		goto L375
	}
L373:
	;
	goto L374
L374:
	;
	v1209 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L21
	} else {
		goto L376
	}
L375:
	;
	goto L374
L376:
	;
	F_ScanKeyInit(m, v1187+int32(48), int32(2), int32(3), int32(184), v1194)
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L21
	} else {
		goto L377
	}
L377:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_ScanKeyInit(m, v1187+int32(96), int32(4), int32(3), int32(62), v1223)
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L21
	} else {
		goto L378
	}
L378:
	;
	v1232 = F_systable_beginscan(m, v1209, int32(2701), int32(1), int32(0), int32(2), v1187+int32(48))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L21
	} else {
		goto L381
	}
L379:
	;
	goto L1
L380:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L21
	} else {
		goto L403
	}
L381:
	;
	v1234 = F_systable_getnext(m, v1232)
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L21
	} else {
		goto L382
	}
L382:
	;
	if v1234 != 0 {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+16))
	v1237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1236)+22)))
	v1238 = v1236 + v1237
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1238)+8))
	if v1239 != 0 {
		goto L380
	} else {
		goto L386
	}
L384:
	;
	goto L385
L385:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L21
	} else {
		goto L399
	}
L386:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1238)))
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_renametrig_internal(m, v1209, v1197, v1234, v1241, v1242)
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L21
	} else {
		goto L387
	}
L387:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1197)+48))
	v1246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1245)+119)))
	if v1246 != int32(112) {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1240
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2620)
	F_systable_endscan(m, v1232)
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L21
	} else {
		goto L396
	}
L389:
	;
	v1250 = F_RelationGetPartitionDesc(m, v1197, int32(1))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L21
	} else {
		goto L390
	}
L390:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1250)))
	if v1252 <= int32(0) {
		goto L388
	} else {
		goto L391
	}
L391:
	;
	v1259 = int32(0)
	goto L392
L392:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+8))
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1268+v1259<<(uint(int32(2))%32))))
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v1238)))
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_renametrig_partition(m, v1209, v1272, v1273, v1274, v1275)
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L21
	} else {
		goto L394
	}
L393:
	;
	goto L388
L394:
	;
	v1279 = v1259 + int32(1)
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1250)))
	if v1279 < v1280 {
		v1259 = v1279
		goto L392
	} else {
		goto L395
	}
L395:
	;
	goto L393
L396:
	;
	F_sequence_close(m, v1209, int32(3))
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L21
	} else {
		goto L397
	}
L397:
	;
	F_relation_close(m, v1197, int32(0))
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		goto L21
	} else {
		goto L398
	}
L398:
	;
	m.G0 = v1187 + int32(144)
	goto L379
L399:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L21
	} else {
		goto L400
	}
L400:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v1197)+48))
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1187))) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v1187)+4)) = v1317 + int32(4)
	F_errmsg(m, int32(78539), v1187)
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L21
	} else {
		goto L401
	}
L401:
	;
	F_errfinish(m, int32(520840), int32(1558), int32(354938))
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L21
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
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L21
	} else {
		goto L404
	}
L404:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1197)+48))
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1187)+32)) = v1339
	*(*int32)(unsafe.Add(mBase, uint32(v1187)+36)) = v1338 + int32(4)
	F_errmsg(m, int32(754885), v1187+int32(32))
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L21
	} else {
		goto L405
	}
L405:
	;
	v1350 = F_get_partition_parent(m, v1194, int32(0))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L21
	} else {
		goto L406
	}
L406:
	;
	v1352 = F_get_rel_name(m, v1350)
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L21
	} else {
		goto L407
	}
L407:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1187)+16)) = v1352
	F_errhint(m, int32(683125), v1187+int32(16))
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L21
	} else {
		goto L408
	}
L408:
	;
	F_errfinish(m, int32(520840), int32(1532), int32(354938))
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L21
	} else {
		goto L409
	}
L409:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L410:
	;
	v1376 = F_relation_open(m, v1373, int32(0))
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L21
	} else {
		goto L411
	}
L411:
	;
	v1380 = F_table_open(m, int32(3256), int32(3))
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L21
	} else {
		goto L412
	}
L412:
	;
	v1384 = int32(3)
	F_ScanKeyInit(m, v1367+int32(32), v1384, v1384, int32(184), v1373)
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L21
	} else {
		goto L413
	}
L413:
	;
	v1390 = v1367 + int32(80)
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	F_ScanKeyInit(m, v1390, int32(2), int32(3), int32(62), v1394)
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L21
	} else {
		goto L414
	}
L414:
	;
	v1403 = F_systable_beginscan(m, v1380, int32(3258), int32(1), int32(0), int32(2), v1367+int32(32))
	mBase = m.M
	v1404 = m.ExcPending
	if v1404 != 0 {
		goto L21
	} else {
		goto L417
	}
L415:
	;
	goto L1
L416:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L21
	} else {
		goto L443
	}
L417:
	;
	v1405 = F_systable_getnext(m, v1403)
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L21
	} else {
		goto L418
	}
L418:
	;
	if v1405 == int32(0) {
		goto L419
	} else {
		goto L420
	}
L419:
	;
	F_systable_endscan(m, v1403)
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L21
	} else {
		goto L422
	}
L420:
	;
	goto L421
L421:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L21
	} else {
		goto L439
	}
L422:
	;
	v1413 = int32(3)
	F_ScanKeyInit(m, v1367+int32(32), v1413, v1413, int32(184), v1373)
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L21
	} else {
		goto L423
	}
L423:
	;
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_ScanKeyInit(m, v1390, int32(2), int32(3), int32(62), v1421)
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L21
	} else {
		goto L424
	}
L424:
	;
	v1430 = F_systable_beginscan(m, v1380, int32(3258), int32(1), int32(0), int32(2), v1367+int32(32))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L21
	} else {
		goto L425
	}
L425:
	;
	v1432 = F_systable_getnext(m, v1430)
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L21
	} else {
		goto L426
	}
L426:
	;
	if v1432 == int32(0) {
		goto L416
	} else {
		goto L427
	}
L427:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+16))
	v1437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1436)+22)))
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1436+v1437)))
	v1440 = F_heap_copytuple(m, v1432)
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L21
	} else {
		goto L428
	}
L428:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1440)+16))
	v1443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1442)+22)))
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v1449 = F_strncpy(m, v1442+v1443+int32(4), v1447, int32(64))
	mBase = m.M
	v1450 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1449)+63)) = uint8(v1450)
	goto L429
L429:
	;
	F_CatalogTupleUpdate(m, v1380, v1440+int32(4), v1440)
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L21
	} else {
		goto L430
	}
L430:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, _consts[297]))
	if v1457 != 0 {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v1459 = int32(0)
	F_RunObjectPostAlterHook(m, int32(3256), v1439, v1459, v1459, v1459)
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L21
	} else {
		goto L434
	}
L432:
	;
	goto L433
L433:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1439
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(3256)
	F_CacheInvalidateRelcache(m, v1376)
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L21
	} else {
		goto L435
	}
L434:
	;
	goto L433
L435:
	;
	F_systable_endscan(m, v1430)
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		goto L21
	} else {
		goto L436
	}
L436:
	;
	F_sequence_close(m, v1380, int32(3))
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L21
	} else {
		goto L437
	}
L437:
	;
	F_relation_close(m, v1376, int32(0))
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
		goto L21
	} else {
		goto L438
	}
L438:
	;
	m.G0 = v1367 + int32(128)
	goto L415
L439:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L21
	} else {
		goto L440
	}
L440:
	;
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+48))
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1367)+16)) = v1490
	*(*int32)(unsafe.Add(mBase, uint32(v1367)+20)) = v1489 + int32(4)
	F_errmsg(m, int32(125084), v1367+int32(16))
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L21
	} else {
		goto L441
	}
L441:
	;
	F_errfinish(m, int32(516808), int32(1139), int32(23856))
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L21
	} else {
		goto L442
	}
L442:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L443:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L21
	} else {
		goto L444
	}
L444:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+48))
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1367))) = v1513
	*(*int32)(unsafe.Add(mBase, uint32(v1367)+4)) = v1512 + int32(4)
	F_errmsg(m, int32(78410), v1367)
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L21
	} else {
		goto L445
	}
L445:
	;
	F_errfinish(m, int32(516808), int32(1167), int32(23856))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L21
	} else {
		goto L446
	}
L446:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L447:
	;
	v1535 = F_typenameTypeId(m, int32(0), v1533)
	mBase = m.M
	v1536 = m.ExcPending
	if v1536 != 0 {
		goto L21
	} else {
		goto L448
	}
L448:
	;
	v1539 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L21
	} else {
		goto L449
	}
L449:
	;
	v1543 = F_SearchSysCacheCopy(m, int32(82), v1535, int32(0))
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L21
	} else {
		goto L453
	}
L450:
	;
	goto L1
L451:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L21
	} else {
		goto L495
	}
L452:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L21
	} else {
		goto L489
	}
L453:
	;
	if v1543 != 0 {
		goto L454
	} else {
		goto L455
	}
L454:
	;
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v1543)+16))
	v1546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1545)+22)))
	v1547 = v1545 + v1546
	v1550 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v1551 = F_object_ownercheck(m, int32(1247), v1535, v1550)
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L21
	} else {
		goto L457
	}
L455:
	;
	goto L456
L456:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		goto L21
	} else {
		goto L486
	}
L457:
	;
	if v1551 == int32(0) {
		goto L458
	} else {
		goto L459
	}
L458:
	;
	F_aclcheck_error_type(m, int32(2), v1535)
	mBase = m.M
	v1557 = m.ExcPending
	if v1557 != 0 {
		goto L21
	} else {
		goto L461
	}
L459:
	;
	goto L460
L460:
	;
	v1558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1547)+79)))
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1559 == int32(12) {
		goto L463
	} else {
		goto L464
	}
L461:
	;
	goto L460
L462:
	;
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v1547)+92))
	if v1591 != 0 {
		goto L475
	} else {
		goto L476
	}
L463:
	;
	if v1558 == int32(100) {
		goto L462
	} else {
		goto L466
	}
L464:
	;
	goto L465
L465:
	;
	if v1558 != int32(99) {
		goto L462
	} else {
		goto L472
	}
L466:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L21
	} else {
		goto L467
	}
L467:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L21
	} else {
		goto L468
	}
L468:
	;
	v1571 = F_format_type_be(m, v1535)
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L21
	} else {
		goto L469
	}
L469:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1528)+48)) = v1571
	F_errmsg(m, int32(291651), v1528+int32(48))
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L21
	} else {
		goto L470
	}
L470:
	;
	F_errfinish(m, int32(519616), int32(3772), int32(390801))
	mBase = m.M
	v1583 = m.ExcPending
	if v1583 != 0 {
		goto L21
	} else {
		goto L471
	}
L471:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L472:
	;
	v1586 = *(*int32)(unsafe.Add(mBase, uint32(v1547)+84))
	v1587 = F_get_rel_relkind(m, v1586)
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L21
	} else {
		goto L473
	}
L473:
	;
	if v1587 != int32(99) {
		goto L452
	} else {
		goto L474
	}
L474:
	;
	goto L462
L475:
	;
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(v1547)+88))
	if v1592 == int32(6179) {
		goto L451
	} else {
		goto L478
	}
L476:
	;
	goto L477
L477:
	;
	v1595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1547)+79)))
	if v1595 == int32(99) {
		goto L480
	} else {
		goto L481
	}
L478:
	;
	goto L477
L479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1535
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
	F_sequence_close(m, v1539, int32(3))
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L21
	} else {
		goto L485
	}
L480:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v1547)+84))
	v1599 = int32(0)
	F_RenameRelationInternal(m, v1598, v1530, v1599, v1599)
	mBase = m.M
	v1602 = m.ExcPending
	if v1602 != 0 {
		goto L21
	} else {
		goto L483
	}
L481:
	;
	goto L482
L482:
	;
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(v1547)+68))
	F_RenameTypeInternal(m, v1535, v1530, v1603)
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
		goto L21
	} else {
		goto L484
	}
L483:
	;
	goto L479
L484:
	;
	goto L479
L485:
	;
	m.G0 = v1528 + int32(96)
	goto L450
L486:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1528))) = v1535
	F_errmsg_internal(m, int32(55308), v1528)
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L21
	} else {
		goto L487
	}
L487:
	;
	F_errfinish(m, int32(519616), int32(3760), int32(390801))
	mBase = m.M
	v1629 = m.ExcPending
	if v1629 != 0 {
		goto L21
	} else {
		goto L488
	}
L488:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L489:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1636 = m.ExcPending
	if v1636 != 0 {
		goto L21
	} else {
		goto L490
	}
L490:
	;
	v1637 = F_format_type_be(m, v1535)
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L21
	} else {
		goto L491
	}
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1528)+80)) = v1637
	F_errmsg(m, int32(386347), v1528+int32(80))
	mBase = m.M
	v1644 = m.ExcPending
	if v1644 != 0 {
		goto L21
	} else {
		goto L492
	}
L492:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1528)+64)) = int32(567924)
	F_errhint(m, int32(682375), v1528-int32(-64))
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L21
	} else {
		goto L493
	}
L493:
	;
	F_errfinish(m, int32(519616), int32(3787), int32(390801))
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L21
	} else {
		goto L494
	}
L494:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L495:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L21
	} else {
		goto L496
	}
L496:
	;
	v1664 = F_format_type_be(m, v1535)
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L21
	} else {
		goto L497
	}
L497:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1528)+32)) = v1664
	F_errmsg(m, int32(199387), v1528+int32(32))
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L21
	} else {
		goto L498
	}
L498:
	;
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1547)+92))
	v1673 = F_format_type_be(m, v1672)
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L21
	} else {
		goto L499
	}
L499:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1528)+16)) = v1673
	F_errhint(m, int32(652211), v1528+int32(16))
	mBase = m.M
	v1680 = m.ExcPending
	if v1680 != 0 {
		goto L21
	} else {
		goto L500
	}
L500:
	;
	F_errfinish(m, int32(519616), int32(3796), int32(390801))
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		goto L21
	} else {
		goto L501
	}
L501:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L502:
	;
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1694 = F_table_open(m, v1692, int32(3))
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L21
	} else {
		goto L503
	}
L503:
	;
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(v1694)+56))
	v1699 = F_get_object_catcache_oid(m, v1698)
	mBase = m.M
	v1700 = m.ExcPending
	if v1700 != 0 {
		goto L21
	} else {
		goto L504
	}
L504:
	;
	v1701 = F_get_object_catcache_name(m, v1698)
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L21
	} else {
		goto L505
	}
L505:
	;
	v1703 = F_get_object_attnum_name(m, v1698)
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L21
	} else {
		goto L506
	}
L506:
	;
	v1705 = F_get_object_attnum_namespace(m, v1698)
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L21
	} else {
		goto L507
	}
L507:
	;
	v1707 = F_get_object_attnum_owner(m, v1698)
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L21
	} else {
		goto L508
	}
L508:
	;
	v1709 = F_SearchSysCache1(m, v1699, v1697)
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		goto L21
	} else {
		goto L509
	}
L509:
	;
	if v1709 == int32(0) {
		goto L7
	} else {
		goto L510
	}
L510:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1694)+52))
	v1716 = F_heap_getattr_2(m, v1709, v1703, v1713, v15+int32(159))
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L21
	} else {
		goto L511
	}
L511:
	;
	if int32(0) < v1705 {
		goto L512
	} else {
		goto L513
	}
L512:
	;
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1694)+52))
	v1723 = F_heap_getattr_2(m, v1709, v1705, v1720, v15+int32(159))
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L21
	} else {
		goto L515
	}
L513:
	;
	v1725 = v3
	goto L514
L514:
	;
	v1726 = F_superuser(m)
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L21
	} else {
		goto L522
	}
L515:
	;
	v1725 = v1723
	goto L514
L516:
	;
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v1709)+16))
	v1913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1912)+22)))
	v1914 = v1912 + v1913
	v1915 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1914)+104)))
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(v1914)+68))
	F_IsThereFunctionInNamespace(m, v1696, v1915, v1914+int32(112), v1918)
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L21
	} else {
		goto L599
	}
L517:
	;
	v1868 = int32(0)
	v1871 = F_SearchSysCacheExists(m, v1701, v1696, v1868, v1868, v1868)
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		goto L21
	} else {
		goto L580
	}
L518:
	;
	v1861 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v1862 = int32(0)
	v1864 = F_SearchSysCacheExists(m, int32(66), v1861, v1696, v1862, v1862)
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L21
	} else {
		goto L577
	}
L519:
	;
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(v1709)+16))
	v1851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1850)+22)))
	v1852 = v1850 + v1851
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(v1852)+4))
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(v1852)+72))
	F_IsThereOpFamilyInNamespace(m, v1696, v1853, v1854)
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
		goto L21
	} else {
		goto L576
	}
L520:
	;
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v1709)+16))
	v1844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1843)+22)))
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1843+v1844)+68))
	F_IsThereCollationInNamespace(m, v1696, v1846)
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L21
	} else {
		goto L575
	}
L521:
	;
	if v1701 < int32(0) {
		goto L2
	} else {
		goto L570
	}
L522:
	;
	if v1726 == int32(0) {
		goto L523
	} else {
		goto L524
	}
L523:
	;
	if v1707 <= int32(0) {
		goto L6
	} else {
		goto L526
	}
L524:
	;
	goto L525
L525:
	;
	if v1698 <= int32(2752) {
		goto L562
	} else {
		goto L563
	}
L526:
	;
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v1694)+52))
	v1735 = F_heap_getattr_2(m, v1709, v1707, v1732, v15+int32(159))
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L21
	} else {
		goto L527
	}
L527:
	;
	v1738 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v1739 = F_has_privs_of_role(m, v1738, v1735)
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L21
	} else {
		goto L528
	}
L528:
	;
	if v1739 == int32(0) {
		goto L529
	} else {
		goto L530
	}
L529:
	;
	v1744 = F_get_object_type(m, v1698, v1697)
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L21
	} else {
		goto L532
	}
L530:
	;
	goto L531
L531:
	;
	if v1725 == int32(0) {
		goto L534
	} else {
		goto L535
	}
L532:
	;
	F_aclcheck_error(m, int32(2), v1744, v1716)
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L21
	} else {
		goto L533
	}
L533:
	;
	goto L531
L534:
	;
	if v1698 <= int32(2752) {
		goto L540
	} else {
		goto L541
	}
L535:
	;
	v1752 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v1754 = F_object_aclcheck(m, int32(2615), v1725, v1752, int64(512))
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L21
	} else {
		goto L536
	}
L536:
	;
	if v1754 == int32(0) {
		goto L534
	} else {
		goto L537
	}
L537:
	;
	v1759 = F_get_namespace_name(m, v1725)
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		goto L21
	} else {
		goto L538
	}
L538:
	;
	F_aclcheck_error(m, v1754, int32(36), v1759)
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
		goto L21
	} else {
		goto L539
	}
L539:
	;
	goto L534
L540:
	;
	if v1698 == int32(1255) {
		goto L516
	} else {
		goto L543
	}
L541:
	;
	goto L542
L542:
	;
	if v1698 == int32(2753) {
		goto L519
	} else {
		goto L545
	}
L543:
	;
	if v1698 != int32(2616) {
		goto L521
	} else {
		goto L544
	}
L544:
	;
	goto L3
L545:
	;
	if v1698 == int32(3456) {
		goto L520
	} else {
		goto L546
	}
L546:
	;
	if v1698 != int32(6100) {
		goto L521
	} else {
		goto L547
	}
L547:
	;
	v1778 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v1780 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v1782 = F_object_aclcheck(m, int32(1262), v1778, v1780, int64(512))
	mBase = m.M
	v1783 = m.ExcPending
	if v1783 != 0 {
		goto L21
	} else {
		goto L548
	}
L548:
	;
	if v1782 != 0 {
		goto L549
	} else {
		goto L550
	}
L549:
	;
	v1786 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v1787 = F_get_database_name(m, v1786)
	mBase = m.M
	v1788 = m.ExcPending
	if v1788 != 0 {
		goto L21
	} else {
		goto L552
	}
L550:
	;
	goto L551
L551:
	;
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v1709)+16))
	v1792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1791)+22)))
	v1794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1791+v1792)+89)))
	if v1794 != 0 {
		goto L518
	} else {
		goto L554
	}
L552:
	;
	F_aclcheck_error(m, v1782, int32(9), v1787)
	mBase = m.M
	v1790 = m.ExcPending
	if v1790 != 0 {
		goto L21
	} else {
		goto L553
	}
L553:
	;
	goto L551
L554:
	;
	v1795 = F_superuser(m)
	mBase = m.M
	v1796 = m.ExcPending
	if v1796 != 0 {
		goto L21
	} else {
		goto L555
	}
L555:
	;
	if v1795 != 0 {
		goto L518
	} else {
		goto L556
	}
L556:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L21
	} else {
		goto L557
	}
L557:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L21
	} else {
		goto L558
	}
L558:
	;
	F_errmsg(m, int32(19931), int32(0))
	mBase = m.M
	v1807 = m.ExcPending
	if v1807 != 0 {
		goto L21
	} else {
		goto L559
	}
L559:
	;
	F_errhint(m, int32(639340), int32(0))
	mBase = m.M
	v1811 = m.ExcPending
	if v1811 != 0 {
		goto L21
	} else {
		goto L560
	}
L560:
	;
	F_errfinish(m, int32(520586), int32(257), int32(328079))
	mBase = m.M
	v1816 = m.ExcPending
	if v1816 != 0 {
		goto L21
	} else {
		goto L561
	}
L561:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L562:
	;
	if v1698 == int32(1255) {
		goto L516
	} else {
		goto L565
	}
L563:
	;
	goto L564
L564:
	;
	if v1698 == int32(2753) {
		goto L519
	} else {
		goto L567
	}
L565:
	;
	if v1698 == int32(2616) {
		goto L3
	} else {
		goto L566
	}
L566:
	;
	goto L521
L567:
	;
	if v1698 == int32(3456) {
		goto L520
	} else {
		goto L568
	}
L568:
	;
	if v1698 == int32(6100) {
		goto L518
	} else {
		goto L569
	}
L569:
	;
	goto L521
L570:
	;
	if v1725 == int32(0) {
		goto L517
	} else {
		goto L571
	}
L571:
	;
	v1834 = int32(0)
	v1836 = F_SearchSysCacheExists(m, v1701, v1696, v1725, v1834, v1834)
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L21
	} else {
		goto L572
	}
L572:
	;
	if v1836 == int32(0) {
		goto L2
	} else {
		goto L573
	}
L573:
	;
	F_report_namespace_conflict(m, v1698, v1696, v1725)
	mBase = m.M
	v1841 = m.ExcPending
	if v1841 != 0 {
		goto L21
	} else {
		goto L574
	}
L574:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L575:
	;
	goto L2
L576:
	;
	goto L2
L577:
	;
	if v1864 != 0 {
		goto L5
	} else {
		goto L578
	}
L578:
	;
	F_LogicalRepWorkersWakeupAtCommit(m, v1697)
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		goto L21
	} else {
		goto L579
	}
L579:
	;
	goto L2
L580:
	;
	if v1871 == int32(0) {
		goto L2
	} else {
		goto L581
	}
L581:
	;
	if v1698 <= int32(3465) {
		goto L588
	} else {
		goto L589
	}
L582:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L21
	} else {
		goto L595
	}
L583:
	;
	if v1698 != int32(3466) {
		goto L4
	} else {
		goto L594
	}
L584:
	;
	v1893 = int32(124511)
	goto L582
L585:
	;
	v1893 = int32(124756)
	goto L582
L586:
	;
	v1893 = int32(124409)
	goto L582
L587:
	;
	v1893 = int32(124436)
	goto L582
L588:
	;
	if v1698 == int32(1417) {
		goto L586
	} else {
		goto L591
	}
L589:
	;
	goto L590
L590:
	;
	switch v1698 - int32(6100) {
	case 0:
		goto L584
	case 1, 2, 3:
		goto L4
	case 4:
		goto L585
	default:
		goto L583
	}
L591:
	;
	if v1698 == int32(2328) {
		goto L587
	} else {
		goto L592
	}
L592:
	;
	if v1698 != int32(2612) {
		goto L4
	} else {
		goto L593
	}
L593:
	;
	v1893 = int32(125148)
	goto L582
L594:
	;
	v1893 = int32(124477)
	goto L582
L595:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		goto L21
	} else {
		goto L596
	}
L596:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v1696
	F_errmsg(m, v1893, v15+int32(48))
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L21
	} else {
		goto L597
	}
L597:
	;
	F_errfinish(m, int32(520586), int32(107), int32(116892))
	mBase = m.M
	v1910 = m.ExcPending
	if v1910 != 0 {
		goto L21
	} else {
		goto L598
	}
L598:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L599:
	;
	goto L2
L600:
	;
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v1925
	F_errmsg_internal(m, int32(508500), v15)
	mBase = m.M
	v1929 = m.ExcPending
	if v1929 != 0 {
		goto L21
	} else {
		goto L601
	}
L601:
	;
	F_errfinish(m, int32(520586), int32(458), int32(104411))
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
		goto L21
	} else {
		goto L602
	}
L602:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L603:
	;
	m.G0 = v1938 + int32(32)
	goto L1
L604:
	;
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v2014 != 0 {
		goto L628
	} else {
		goto L629
	}
L605:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L21
	} else {
		goto L625
	}
L606:
	;
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1945 = F_makeTypeNameFromNameList(m, v1944)
	mBase = m.M
	v1946 = m.ExcPending
	if v1946 != 0 {
		goto L21
	} else {
		goto L609
	}
L607:
	;
	goto L608
L608:
	;
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	v1970 = F_RangeVarGetRelidExtended(m, v1965, int32(8), v1967, int32(575), int32(0))
	mBase = m.M
	v1971 = m.ExcPending
	if v1971 != 0 {
		goto L21
	} else {
		goto L617
	}
L609:
	;
	v1947 = F_typenameTypeId(m, int32(0), v1945)
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L21
	} else {
		goto L610
	}
L610:
	;
	v1951 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v1952 = m.ExcPending
	if v1952 != 0 {
		goto L21
	} else {
		goto L611
	}
L611:
	;
	v1954 = F_SearchSysCache1(m, int32(82), v1947)
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		goto L21
	} else {
		goto L612
	}
L612:
	;
	if v1954 == int32(0) {
		goto L605
	} else {
		goto L613
	}
L613:
	;
	F_checkDomainOwner(m, v1954)
	mBase = m.M
	v1959 = m.ExcPending
	if v1959 != 0 {
		goto L21
	} else {
		goto L614
	}
L614:
	;
	F_ReleaseCatCache(m, v1954)
	mBase = m.M
	v1961 = m.ExcPending
	if v1961 != 0 {
		goto L21
	} else {
		goto L615
	}
L615:
	;
	F_sequence_close(m, v1951, int32(0))
	mBase = m.M
	v1964 = m.ExcPending
	if v1964 != 0 {
		goto L21
	} else {
		goto L616
	}
L616:
	;
	v2008 = v1947
	v2010 = v3
	goto L604
L617:
	;
	if v1970 != 0 {
		v2008 = int32(0)
		v2010 = v1970
		goto L604
	} else {
		goto L618
	}
L618:
	;
	v1974 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L21
	} else {
		goto L619
	}
L619:
	;
	if v1974 != 0 {
		goto L620
	} else {
		goto L621
	}
L620:
	;
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(v1976)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1938)+16)) = v1977
	F_errmsg(m, int32(350183), v1938+int32(16))
	mBase = m.M
	v1983 = m.ExcPending
	if v1983 != 0 {
		goto L21
	} else {
		goto L623
	}
L621:
	;
	goto L622
L622:
	;
	v1990 = *(*int64)(unsafe.Add(mBase, _consts[350]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v1990
	v1993 = *(*int32)(unsafe.Add(mBase, _consts[351]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1993
	goto L603
L623:
	;
	F_errfinish(m, int32(519627), int32(4186), int32(97359))
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		goto L21
	} else {
		goto L624
	}
L624:
	;
	goto L622
L625:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1938))) = v1947
	F_errmsg_internal(m, int32(55308), v1938)
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		goto L21
	} else {
		goto L626
	}
L626:
	;
	F_errfinish(m, int32(519627), int32(4170), int32(97359))
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L21
	} else {
		goto L627
	}
L627:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L628:
	;
	v2015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2014)+16)))
	v2017 = v2015
	goto L630
L629:
	;
	v2017 = int32(0)
	goto L630
L630:
	;
	F_rename_constraint_internal(m, l0, v2010, v2008, v2012, v2013, v2017&int32(1), int32(0))
	mBase = m.M
	v2022 = m.ExcPending
	if v2022 != 0 {
		goto L21
	} else {
		goto L631
	}
L631:
	;
	goto L603
L632:
	;
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v1694)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v1697
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v2035 + int32(4)
	F_errmsg_internal(m, int32(748442), v15+int32(16))
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		goto L21
	} else {
		goto L633
	}
L633:
	;
	F_errfinish(m, int32(520586), int32(189), int32(328079))
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L21
	} else {
		goto L634
	}
L634:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L635:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v2056 = m.ExcPending
	if v2056 != 0 {
		goto L21
	} else {
		goto L636
	}
L636:
	;
	v2057 = F_getObjectDescriptionOids(m, v1698, v1697)
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L21
	} else {
		goto L637
	}
L637:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v2057
	F_errmsg(m, int32(205423), v15+int32(80))
	mBase = m.M
	v2064 = m.ExcPending
	if v2064 != 0 {
		goto L21
	} else {
		goto L638
	}
L638:
	;
	F_errfinish(m, int32(520586), int32(215), int32(328079))
	mBase = m.M
	v2069 = m.ExcPending
	if v2069 != 0 {
		goto L21
	} else {
		goto L639
	}
L639:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L640:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v2076 = m.ExcPending
	if v2076 != 0 {
		goto L21
	} else {
		goto L641
	}
L641:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v1696
	F_errmsg(m, int32(124511), v15-int32(-64))
	mBase = m.M
	v2082 = m.ExcPending
	if v2082 != 0 {
		goto L21
	} else {
		goto L642
	}
L642:
	;
	F_errfinish(m, int32(520586), int32(107), int32(116892))
	mBase = m.M
	v2087 = m.ExcPending
	if v2087 != 0 {
		goto L21
	} else {
		goto L643
	}
L643:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L644:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v1698
	F_errmsg_internal(m, int32(63269), v15+int32(32))
	mBase = m.M
	v2098 = m.ExcPending
	if v2098 != 0 {
		goto L21
	} else {
		goto L645
	}
L645:
	;
	F_errfinish(m, int32(520586), int32(101), int32(116892))
	mBase = m.M
	v2103 = m.ExcPending
	if v2103 != 0 {
		goto L21
	} else {
		goto L646
	}
L646:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L647:
	;
	goto L2
L648:
	;
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v1694)+48))
	v2121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2120)+120)))
	v2122 = F_palloc0(m, v2121)
	mBase = m.M
	v2123 = m.ExcPending
	if v2123 != 0 {
		goto L21
	} else {
		goto L649
	}
L649:
	;
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v1694)+48))
	v2125 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2124)+120)))
	v2126 = F_palloc0(m, v2125)
	mBase = m.M
	v2127 = m.ExcPending
	if v2127 != 0 {
		goto L21
	} else {
		goto L650
	}
L650:
	;
	v2131 = F_strncpy(m, v15+int32(95), v1696, int32(64))
	mBase = m.M
	v2132 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2131)+63)) = uint8(v2132)
	goto L651
L651:
	;
	v2134 = int32(1)
	v2135 = v1703 - v2134
	*(*int32)(unsafe.Add(mBase, uint32(v2118+v2135<<(uint(int32(2))%32)))) = v15 + int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v2126+v2135))) = uint8(v2134)
	v2147 = *(*int32)(unsafe.Add(mBase, uint32(v1694)+52))
	v2148 = F_heap_modify_tuple(m, v1709, v2147, v2118, v2122, v2126)
	mBase = m.M
	v2149 = m.ExcPending
	if v2149 != 0 {
		goto L21
	} else {
		goto L652
	}
L652:
	;
	F_CatalogTupleUpdate(m, v1694, v1709+int32(4), v2148)
	mBase = m.M
	v2151 = m.ExcPending
	if v2151 != 0 {
		goto L21
	} else {
		goto L653
	}
L653:
	;
	v2153 = *(*int32)(unsafe.Add(mBase, _consts[297]))
	if v2153 != 0 {
		goto L654
	} else {
		goto L655
	}
L654:
	;
	v2154 = int32(0)
	F_RunObjectPostAlterHook(m, v1698, v1697, v2154, v2154, v2154)
	mBase = m.M
	v2158 = m.ExcPending
	if v2158 != 0 {
		goto L21
	} else {
		goto L657
	}
L655:
	;
	goto L656
L656:
	;
	if v1698 == int32(6104) {
		goto L658
	} else {
		goto L659
	}
L657:
	;
	goto L656
L658:
	;
	v2161 = *(*int32)(unsafe.Add(mBase, uint32(v1709)+16))
	v2162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2161)+22)))
	v2163 = v2161 + v2162
	v2164 = *(*int32)(unsafe.Add(mBase, uint32(v2163)))
	v2165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2163)+72)))
	if v2165 != 0 {
		goto L662
	} else {
		goto L663
	}
L659:
	;
	goto L660
L660:
	;
	F_pfree(m, v2118)
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L21
	} else {
		goto L676
	}
L661:
	;
	goto L660
L662:
	;
	F_CacheInvalidateRelSync(m, int32(0))
	mBase = m.M
	v2168 = m.ExcPending
	if v2168 != 0 {
		goto L21
	} else {
		goto L665
	}
L663:
	;
	goto L664
L664:
	;
	v2170 = F_GetPublicationRelations(m, v2164, int32(2))
	mBase = m.M
	v2171 = m.ExcPending
	if v2171 != 0 {
		goto L21
	} else {
		goto L667
	}
L665:
	;
	goto L661
L666:
	;
	goto L661
L667:
	;
	v2172 = F_GetAllSchemaPublicationRelations(m, v2164)
	mBase = m.M
	v2173 = m.ExcPending
	if v2173 != 0 {
		goto L21
	} else {
		goto L668
	}
L668:
	;
	v2174 = F_list_concat_unique_oid(m, v2170, v2172)
	mBase = m.M
	v2175 = m.ExcPending
	if v2175 != 0 {
		goto L21
	} else {
		goto L669
	}
L669:
	;
	if v2174 == int32(0) {
		goto L666
	} else {
		goto L670
	}
L670:
	;
	v2178 = *(*int32)(unsafe.Add(mBase, uint32(v2174)+4))
	if v2178 <= int32(0) {
		goto L666
	} else {
		goto L671
	}
L671:
	;
	v2186 = int32(0)
	goto L672
L672:
	;
	v2194 = *(*int32)(unsafe.Add(mBase, uint32(v2174)+12))
	v2198 = *(*int32)(unsafe.Add(mBase, uint32(v2194+v2186<<(uint(int32(2))%32))))
	F_CacheInvalidateRelSync(m, v2198)
	mBase = m.M
	v2200 = m.ExcPending
	if v2200 != 0 {
		goto L21
	} else {
		goto L674
	}
L673:
	;
	goto L666
L674:
	;
	v2202 = v2186 + int32(1)
	v2203 = *(*int32)(unsafe.Add(mBase, uint32(v2174)+4))
	if v2202 < v2203 {
		v2186 = v2202
		goto L672
	} else {
		goto L675
	}
L675:
	;
	goto L673
L676:
	;
	F_pfree(m, v2122)
	mBase = m.M
	v2244 = m.ExcPending
	if v2244 != 0 {
		goto L21
	} else {
		goto L677
	}
L677:
	;
	F_pfree(m, v2126)
	mBase = m.M
	v2246 = m.ExcPending
	if v2246 != 0 {
		goto L21
	} else {
		goto L678
	}
L678:
	;
	F_pfree(m, v2148)
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L21
	} else {
		goto L679
	}
L679:
	;
	F_ReleaseCatCache(m, v1709)
	mBase = m.M
	v2250 = m.ExcPending
	if v2250 != 0 {
		goto L21
	} else {
		goto L680
	}
L680:
	;
	F_sequence_close(m, v1694, int32(3))
	mBase = m.M
	v2253 = m.ExcPending
	if v2253 != 0 {
		goto L21
	} else {
		goto L681
	}
L681:
	;
	goto L1
}
func F_ExecStoreMinimalTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v6 == int32(1654124) {
		v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		if v9&int32(4) != 0 {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
			F_pfree(m, v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v20 = v17 & int32(-5)
				v21 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v21)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v21
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v21)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = l0
				v31 = v20 & int32(65533)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v31)
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v34 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = l0 - v34
				*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v33 + v34
				if l2 != 0 {
					v41 = v31 | int32(4)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v41)
				} else {
				}
				return l1
			}
		} else {
			v20 = v9
			v21 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v21)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v21
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v21)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = l0
			v31 = v20 & int32(65533)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v31)
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v34 = int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = l0 - v34
			*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v33 + v34
			if l2 != 0 {
				v41 = v31 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v41)
			} else {
			}
			return l1
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(92280), int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(519289), int32(1647), int32(404679))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
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
func F_ExecStoreVirtualTuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v2 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v4 = v2 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v4)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v7)
	return l0
}
func F_ExecSupportsBackwardScan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
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
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	v2 = int32(0)
	if l0 == v2 {
		v60 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v60
L2:
	;
	v6 = l0
	goto L3
L3:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+36)))
	if v9 != 0 {
		v60 = v2
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v60 = int32(1)
	goto L1
L5:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	switch v10 - int32(331) {
	case 0:
		goto L14
	default:
		v60 = v2
		goto L1
	case 3:
		goto L13
	case 8, 14, 15, 17, 18, 20, 29, 31:
		goto L6
	case 10:
		goto L12
	case 11:
		goto L11
	case 16:
		goto L10
	case 24:
		goto L8
	case 41, 42:
		goto L7
	}
L6:
	;
	goto L4
L7:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
	if v57 != 0 {
		v6 = v57
		goto L3
	} else {
		goto L28
	}
L8:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+80)))
	return v53 & int32(1)
L9:
	;
	if v52 != 0 {
		v6 = v52
		goto L3
	} else {
		goto L27
	}
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
	v52 = v51
	goto L9
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
	v48 = F_IndexSupportsBackwardScan(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L21
	} else {
		goto L26
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
	v44 = F_IndexSupportsBackwardScan(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L21
	} else {
		goto L25
	}
L13:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
	if int32(0) < v14 {
		v60 = v2
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
	if v13 != 0 {
		v52 = v13
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v60 = v2
	goto L1
L16:
	;
	v17 = int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+76))
	if v18 == int32(0) {
		v60 = v17
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v21 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v22 <= v21 {
		v60 = v17
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v25 = v21
	goto L19
L19:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v25<<(uint(int32(2))%32))))
	v33 = F_ExecSupportsBackwardScan(m, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v60 = v33
	goto L1
L21:
	;
	return int32(0)
L22:
	;
	if v33 == int32(0) {
		v60 = v33
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v40 = v25 + int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v40 < v41 {
		v25 = v40
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	return v44
L26:
	;
	return v48
L27:
	;
	v60 = v2
	goto L1
L28:
	;
	v60 = v2
	goto L1
}
func F_ExtendBufferedRelTo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
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
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int64
	_ = v137
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
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v198 int32
	_ = v198
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v226 int64
	_ = v226
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
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
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int64
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int64
	_ = v346
	var v348 int64
	_ = v348
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v397 int64
	_ = v397
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int64
	_ = v418
	var v424 int64
	_ = v424
	var v432 int32
	_ = v432
	var v433 int64
	_ = v433
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int64
	_ = v444
	var v462 int32
	_ = v462
	var v468 int64
	_ = v468
	var v474 int64
	_ = v474
	var v476 int32
	_ = v476
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v538 int32
	_ = v538
	v6 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(400)
	m.G0 = v18
	*(*int32)(unsafe.Add(mBase, uint32(v18)+312)) = v6
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v22 == v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v60 = v22
	goto L3
L3:
	;
	if l2&int32(4) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+48))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+118)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v58)
	v60 = v54
	goto L3
L5:
	;
	v54 = v26
	v55 = v25
	goto L4
L6:
	;
	goto L7
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v28
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v30
	v34 = F_smgropen(m, v18+int32(32), v27)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v34
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	if v40 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v54 = v52
	v55 = v53
	goto L4
L11:
	;
	v48 = v40
	goto L13
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v34)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	v48 = v46
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v48 + int32(1)
	goto L10
L14:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v99 = l4 - int32(1)
	if l2&int32(16) != 0 {
		goto L27
	} else {
		goto L28
	}
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v60+l1<<(uint(int32(2))%32))+20))
	v71 = int32(1)
	if base.Ui32(v71) < base.Ui32(v70+v71) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v75 = F_smgrexists(m, v60, l1)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	if v75 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LockRelationForExtension(m, v77, int32(7))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v82 = F_smgrexists(m, v81, l1)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	if v82 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_smgrcreate(m, v86, l1, int32(base.Ui32(l2&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L8
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_UnlockRelationForExtension(m, v93, int32(7))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L8
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	goto L14
L26:
	;
	m.G0 = v18 + int32(400)
	return v538
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97+l1<<(uint(int32(2))%32))+20)) = int32(-1)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v108 = v107
	goto L29
L28:
	;
	v108 = v97
	goto L29
L29:
	;
	v109 = F_smgrnblocks(m, v108, l1)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	if base.Ui32(v109) < base.Ui32(l3) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if base.Ui32(v99) < base.Ui32(int32(2)) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l3 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L34:
	;
	v116 = l2 | int32(32)
	goto L36
L35:
	;
	v116 = l2
	goto L36
L36:
	;
	v129 = v109
	v131 = v6
	goto L37
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v135
	v137 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v137
	if base.Ui64(base.I64_extend_i32_u(v129)-int64(-64)) <= base.Ui64(base.I64_extend_i32_u(l3)) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v198 != 0 {
		v538 = v198
		goto L26
	} else {
		goto L55
	}
L39:
	;
	v148 = int32(64)
	goto L41
L40:
	;
	v148 = l3 - v129
	goto L41
L41:
	;
	v153 = F_ExtendBufferedRelCommon(m, v18+int32(16), l1, int32(0), v116, v148, l3, v18+int32(48), v18+int32(312))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v18)+312))
	v156 = v153 + v155
	if v155 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v160 = int32(0)
	v169 = v131
	goto L46
L44:
	;
	v198 = v131
	goto L45
L45:
	;
	if base.Ui32(v156) < base.Ui32(l3) {
		v129 = v156
		v131 = v198
		goto L37
	} else {
		goto L54
	}
L46:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(48)+v160<<(uint(int32(2))%32))))
	if l3-int32(1) == v160+v153 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v198 = v183
	goto L45
L48:
	;
	v185 = v160 + int32(1)
	if v185 != v155 {
		v160 = v185
		v169 = v183
		goto L46
	} else {
		goto L53
	}
L49:
	;
	v183 = v178
	goto L48
L50:
	;
	goto L51
L51:
	;
	F_ReleaseBuffer(m, v178)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L8
	} else {
		goto L52
	}
L52:
	;
	v183 = v169
	goto L48
L53:
	;
	goto L47
L54:
	;
	goto L38
L55:
	;
	goto L33
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+324)) = int64(0)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v18)+328))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v18)+320)) = v218
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v18)+320))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v226
	if base.Ui32(v99) < base.Ui32(int32(2)) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	if v218 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	v233 = int32(9)
	goto L61
L60:
	;
	v233 = int32(1)
	goto L61
L61:
	;
	v234 = F_ExtendBufferedRel(m, v18, l1, int32(0), v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L8
	} else {
		goto L62
	}
L62:
	;
	v538 = v234
	goto L26
L63:
	;
	v242 = int32(1)
	v243 = l3 - v242
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v99) <= base.Ui32(v242) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v241 = v238
	goto L63
L65:
	;
	goto L66
L66:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v218)+48))
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+118)))
	v241 = v240
	goto L63
L67:
	;
	if v241&int32(255) == int32(116) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+336)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+332)) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+328)) = uint8(v241)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+320)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v18)+324)) = v244
	if l4 == int32(3) {
		goto L133
	} else {
		goto L134
	}
L70:
	;
	if v218 != 0 {
		goto L111
	} else {
		goto L112
	}
L71:
	;
	v251 = int32(1)
	v254 = F_LocalBufferAlloc(m, v244, l1, v243, v18+int32(316))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L8
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v267 = F_IOContextForStrategy(m, int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L8
	} else {
		goto L76
	}
L74:
	;
	v256 = int32(3)
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+316)))
	if v257 != int32(1) {
		v402 = v254
		v405 = v251
		v406 = v257
		v407 = v256
		goto L70
	} else {
		goto L75
	}
L75:
	;
	v260 = int32(4460360)
	v262 = *(*int64)(unsafe.Add(mBase, _consts[266]))
	*(*int64)(unsafe.Add(mBase, _consts[266])) = v262 + int64(1)
	v402 = v254
	v405 = v251
	v406 = v257
	v407 = v256
	goto L70
L76:
	;
	v271 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_ResourceOwnerEnlarge(m, v271)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+320)) = v276
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+324)) = v278
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v244)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+336)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v18)+332)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v18)+328)) = v280
	v286 = F_BufTableHashCode(m, v18+int32(320))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v296 = v289 + v286&int32(127)<<(uint(int32(7))%32) + int32(6912)
	v298 = F_LWLockAcquire(m, v296, int32(1))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L8
	} else {
		goto L80
	}
L80:
	;
	v302 = F_BufTableLookup(m, v18+int32(320), v286)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L8
	} else {
		goto L82
	}
L81:
	;
	v395 = int32(4460328)
	v397 = *(*int64)(unsafe.Add(mBase, _consts[277]))
	*(*int64)(unsafe.Add(mBase, _consts[277])) = v397 + int64(1)
	v402 = v390
	v405 = int32(0)
	v406 = int32(1)
	v407 = v267
	goto L70
L82:
	;
	if int32(0) <= v302 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v310 = v307 + v302<<(uint(int32(6))%32)
	v312 = F_PinBuffer(m, v310, int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L8
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	F_LWLockRelease(m, v296)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L8
	} else {
		goto L89
	}
L86:
	;
	F_LWLockRelease(m, v296)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L8
	} else {
		goto L87
	}
L87:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+316)) = uint8(v312)
	if v312 != 0 {
		v390 = v310
		goto L81
	} else {
		goto L88
	}
L88:
	;
	v402 = v310
	v405 = int32(0)
	v406 = int32(0)
	v407 = v267
	goto L70
L89:
	;
	v321 = F_GetVictimBuffer(m, int32(0), v267)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v326 = F_LWLockAcquire(m, v296, int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L8
	} else {
		goto L91
	}
L91:
	;
	v330 = v324 + v321<<(uint(int32(6))%32)
	v332 = v330 + int32(-64)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v330-int32(44))))
	v338 = F_BufTableInsert(m, v18+int32(320), v286, v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L8
	} else {
		goto L92
	}
L92:
	;
	if v338 < int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v342 = F_LockBufHdr(m, v332)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L8
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	F_UnpinBuffer(m, v332)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L8
	} else {
		goto L104
	}
L96:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v18)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v332)+16)) = v344
	v346 = *(*int64)(unsafe.Add(mBase, uint32(v18)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v332)+8)) = v346
	v348 = *(*int64)(unsafe.Add(mBase, uint32(v18)+320))
	*(*int64)(unsafe.Add(mBase, uint32(v332))) = v348
	v354 = int32(-2113667072)
	if v241&int32(255) == int32(112) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v361 = v354
	goto L99
L98:
	;
	v361 = int32(33816576)
	goto L99
L99:
	;
	if l1 == int32(3) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v364 = v354
	goto L102
L101:
	;
	v364 = v361
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330-int32(40)))) = v342&int32(-38010881) | v364
	F_LWLockRelease(m, v296)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	v369 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+316)) = uint8(v369)
	v402 = v332
	v405 = v369
	v406 = v369
	v407 = v267
	goto L70
L104:
	;
	F_StrategyFreeBuffer(m, v332)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	v377 = int32(0)
	v379 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v382 = v379 + v338<<(uint(int32(6))%32)
	v384 = F_PinBuffer(m, v382, v377)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L8
	} else {
		goto L106
	}
L106:
	;
	F_LWLockRelease(m, v296)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L8
	} else {
		goto L107
	}
L107:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+316)) = uint8(v384)
	if v384 != 0 {
		v390 = v382
		goto L81
	} else {
		goto L108
	}
L108:
	;
	v402 = v382
	v405 = int32(0)
	v406 = v377
	v407 = v267
	goto L70
L109:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v402)+20))
	v501 = v499 + int32(1)
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+316)))
	F_ZeroAndLockBuffer(m, v501, l4, v502)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L8
	} else {
		goto L132
	}
L110:
	;
	v462 = v405*int32(320) + v407<<(uint(int32(6))%32)
	v468 = *(*int64)(unsafe.Add(mBase, uint32(v462)+uint32(_consts[727])))
	*(*int64)(unsafe.Add(mBase, uint32(v462)+uint32(_consts[727]))) = v468 + int64(1)
	v474 = *(*int64)(unsafe.Add(mBase, uint32(v462)+uint32(_consts[728])))
	*(*int64)(unsafe.Add(mBase, uint32(v462)+uint32(_consts[728]))) = v474
	v476 = int32(1)
	F_pgstat_count_backend_io_op(m, v405, v407, int32(2), v476, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _consts[729])) = uint8(v476)
	*(*uint8)(unsafe.Add(mBase, _consts[730])) = uint8(v476)
	goto L130
L111:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v218)+272))
	if v408 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L112:
	;
	goto L113
L113:
	;
	if v406 == int32(0) {
		goto L109
	} else {
		goto L129
	}
L114:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v218)+272))
	if v432 != 0 {
		goto L123
	} else {
		goto L124
	}
L115:
	;
	if v406 == int32(0) {
		goto L109
	} else {
		goto L122
	}
L116:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+268)))
	if v411 != int32(1) {
		goto L115
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v408)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v408)+112)) = v424 + int64(1)
	goto L115
L119:
	;
	F_pgstat_assoc_relation(m, v218)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L8
	} else {
		goto L120
	}
L120:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+316)))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v218)+272))
	v418 = *(*int64)(unsafe.Add(mBase, uint32(v417)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v417)+112)) = v418 + int64(1)
	if v416&int32(1) != 0 {
		goto L114
	} else {
		goto L121
	}
L121:
	;
	goto L109
L122:
	;
	goto L114
L123:
	;
	v433 = *(*int64)(unsafe.Add(mBase, uint32(v432)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v432)+120)) = v433 + int64(1)
	goto L110
L124:
	;
	goto L125
L125:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+268)))
	if v437 != int32(1) {
		goto L110
	} else {
		goto L126
	}
L126:
	;
	F_pgstat_assoc_relation(m, v218)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L8
	} else {
		goto L127
	}
L127:
	;
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+316)))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v218)+272))
	v444 = *(*int64)(unsafe.Add(mBase, uint32(v443)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v443)+120)) = v444 + int64(1)
	if v442&int32(1) != 0 {
		goto L110
	} else {
		goto L128
	}
L128:
	;
	goto L109
L129:
	;
	goto L110
L130:
	;
	v486 = int32(*(*uint8)(unsafe.Add(mBase, _consts[409])))
	if v486 != int32(1) {
		goto L109
	} else {
		goto L131
	}
L131:
	;
	v489 = int32(4556816)
	v491 = *(*int32)(unsafe.Add(mBase, _consts[410]))
	v493 = *(*int32)(unsafe.Add(mBase, _consts[731]))
	*(*int32)(unsafe.Add(mBase, _consts[410])) = v491 + v493
	goto L109
L132:
	;
	v538 = v501
	goto L26
L133:
	;
	v519 = int32(9)
	goto L135
L134:
	;
	v519 = int32(8)
	goto L135
L135:
	;
	v520 = F_StartReadBuffer(m, v18+int32(320), v18+int32(316), v243, v519)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L8
	} else {
		goto L136
	}
L136:
	;
	if v520 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	F_WaitReadBuffers(m, v18+int32(320))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L8
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v18)+316))
	v538 = v526
	goto L26
L140:
	;
	goto L139
}
func F_ExtractReplicaIdentity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(8000)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+130)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v5)
	v20 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	if v20 < int32(2) {
		v120 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(8000)
	return v120
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+118)))
	if v24 != int32(112) {
		v120 = v5
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+119)))
	if v27 == int32(102) {
		v120 = v5
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L5
L5:
	;
	if base.Ui32(v30) < base.Ui32(int32(12000)) {
		v120 = v5
		goto L1
	} else {
		goto L6
	}
L6:
	;
	switch v16 - int32(102) {
	case 0:
		goto L8
	default:
		goto L7
	case 8:
		v120 = v5
		goto L1
	}
L7:
	;
	if l2 == int32(0) {
		v120 = v5
		goto L1
	} else {
		goto L14
	}
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+20)))
	if v36&int32(4) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v120 = l1
	goto L1
L10:
	;
	goto L11
L11:
	;
	v41 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v41)
	v43 = F_toast_flatten_tuple(m, l1, v14)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v120 = v43
	goto L1
L14:
	;
	v50 = F_RelationGetIndexAttrBitmap(m, l0, int32(2))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	if v50 == int32(0) {
		v120 = v5
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_heap_deform_tuple(m, l1, v14, v12, v12+int32(6400))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v58 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v58 < v59 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v62 = v58
	goto L21
L19:
	;
	goto L20
L20:
	;
	v97 = F_heap_form_tuple(m, v14, v12, v12+int32(6400))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L12
	} else {
		goto L28
	}
L21:
	;
	v73 = F_bms_is_member(m, v62+int32(8), v50)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L12
	} else {
		goto L23
	}
L22:
	;
	goto L20
L23:
	;
	if v73 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v80 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(6400)+v62))) = uint8(v80)
	goto L26
L25:
	;
	goto L26
L26:
	;
	v83 = v62 + int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v83 < v84 {
		v62 = v83
		goto L21
	} else {
		goto L27
	}
L27:
	;
	goto L22
L28:
	;
	v99 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v99)
	F_bms_free(m, v50)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)))
	if v104&int32(4) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v120 = v97
	goto L1
L31:
	;
	goto L32
L32:
	;
	v109 = F_toast_flatten_tuple(m, v97, v14)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	F_pfree(m, v97)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	v120 = v109
	goto L1
}
func F___env_rm_add(m *base.Module, l0 int32, l1 int32) {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v9 = *(*int32)(unsafe.Add(mBase, _consts[478]))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _consts[479]))
	v13 = l1
	v14 = int32(0)
	goto L4
L2:
	;
	v35 = l1
	goto L3
L3:
	;
	if v35 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	v21 = v11 + v14<<(uint(int32(2))%32)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if l0 == v22 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v35 = v30
	goto L3
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v13
	F_emscripten_builtin_free(m, l0)
	mBase = m.M
	return
L7:
	;
	goto L8
L8:
	;
	if v22 != 0 {
		v30 = v13
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v32 = v14 + int32(1)
	if v32 != v9 {
		v13 = v30
		v14 = v32
		goto L4
	} else {
		goto L12
	}
L10:
	;
	if v13 == int32(0) {
		v30 = v13
		goto L9
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v13
	v30 = int32(0)
	goto L9
L12:
	;
	goto L5
L13:
	;
	return
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[479]))
	v49 = F_emscripten_builtin_realloc(m, v44, v9<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	if v49 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[479])) = v49
	v54 = int32(4735764)
	v56 = *(*int32)(unsafe.Add(mBase, _consts[478]))
	*(*int32)(unsafe.Add(mBase, _consts[478])) = v56 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v49+v56<<(uint(int32(2))%32)))) = v35
	goto L13
}
func F__exit(m *base.Module, l0 int32) {
	F__Exit(m, l0)
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_each(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_each(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_each_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+32))
	if v5 == int32(1) {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
		if v8 != int32(1) {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v18
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
			if v11 != int32(1) {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v18
				return int32(0)
			} else {
				v14 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)) = uint8(v14)
				return int32(0)
			}
		}
	} else {
		return int32(0)
	}
}
func F_eclass_member_iterator_next(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v6 == v2 {
		v110 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v110
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v97 = v91 + int32(4)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if base.Ui32(v97) < base.Ui32(v93+v99<<(uint(int32(2))%32)) {
		goto L23
	} else {
		goto L24
	}
L4:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v91 = v9
	v92 = v6
	v93 = v10
	goto L3
L5:
	;
	goto L6
L6:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = v11
	goto L7
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v17 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v88
	v91 = v88
	v92 = v84
	v93 = v88
	goto L3
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v73
	if v73 <= int32(0) {
		v110 = v2
		goto L1
	} else {
		goto L20
	}
L10:
	;
	v73 = base.I32_ctz(v59) | v60<<(uint(int32(5))%32)
	goto L9
L11:
	;
	v73 = int32(-2)
	goto L9
L12:
	;
	v24 = v13 + int32(1)
	v26 = base.I32_div_s(v24, int32(32))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v27 <= v26 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v30 = v17 + int32(8)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v26<<(uint(int32(2))%32))))
	v37 = v34 & (int32(-1) << (uint(v24) % 32))
	if v37 != 0 {
		v59 = v37
		v60 = v26
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v39 = v26 + int32(1)
	if v39 == v27 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v42 = v39
	goto L16
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v30+v42<<(uint(int32(2))%32))))
	if v49 != 0 {
		v59 = v49
		v60 = v42
		goto L10
	} else {
		goto L18
	}
L17:
	;
	goto L11
L18:
	;
	v51 = v42 + int32(1)
	if v51 != v27 {
		v42 = v51
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	if v78 <= v73 {
		v110 = v2
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80+v73<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v84
	if v84 == int32(0) {
		v13 = v73
		goto L7
	} else {
		goto L22
	}
L22:
	;
	goto L8
L23:
	;
	v104 = v97
	goto L25
L24:
	;
	v104 = int32(0)
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v104
	v110 = v95
	goto L1
}
func F_elem_contained_by_range(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
		if v19 != 0 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v20 == v17 {
				v30 = v19
				v31 = F_range_contains_elem_internal(m, v30, v13, v11)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					m.G0 = v9 + int32(16)
					return v31
				}
			} else {
				v23 = F_lookup_type_cache(m, v17, int32(2048))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
					if v25 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
							F_errmsg_internal(m, int32(389532), v9)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(519276), int32(1776), int32(418852))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
						v30 = v23
						v31 = F_range_contains_elem_internal(m, v30, v13, v11)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v31
						}
					}
				}
			}
		} else {
			v23 = F_lookup_type_cache(m, v17, int32(2048))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
				if v25 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
						F_errmsg_internal(m, int32(389532), v9)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(519276), int32(1776), int32(418852))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
					v30 = v23
					v31 = F_range_contains_elem_internal(m, v30, v13, v11)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v31
					}
				}
			}
		}
	}
}
func F_english_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v291 int32
	_ = v291
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v525 int32
	_ = v525
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v596 int32
	_ = v596
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v647 int32
	_ = v647
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v673 int32
	_ = v673
	var v681 int32
	_ = v681
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v722 int32
	_ = v722
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v773 int32
	_ = v773
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v799 int32
	_ = v799
	var v806 int32
	_ = v806
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v844 int32
	_ = v844
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v895 int32
	_ = v895
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v921 int32
	_ = v921
	var v929 int32
	_ = v929
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1029 int32
	_ = v1029
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1045 int32
	_ = v1045
	var v1059 int32
	_ = v1059
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1101 int32
	_ = v1101
	var v1106 int32
	_ = v1106
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1131 int32
	_ = v1131
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1161 int32
	_ = v1161
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1221 int32
	_ = v1221
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1234 int32
	_ = v1234
	var v1240 int32
	_ = v1240
	var v1249 int32
	_ = v1249
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1277 int32
	_ = v1277
	var v1280 int32
	_ = v1280
	var v1284 int32
	_ = v1284
	var v1288 int32
	_ = v1288
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1368 int32
	_ = v1368
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1387 int32
	_ = v1387
	var v1389 int32
	_ = v1389
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1410 int32
	_ = v1410
	var v1428 int32
	_ = v1428
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1441 int32
	_ = v1441
	var v1447 int32
	_ = v1447
	var v1456 int32
	_ = v1456
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1502 int32
	_ = v1502
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1559 int32
	_ = v1559
	var v1561 int32
	_ = v1561
	var v1574 int32
	_ = v1574
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1596 int32
	_ = v1596
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1609 int32
	_ = v1609
	var v1615 int32
	_ = v1615
	var v1631 int32
	_ = v1631
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1654 int32
	_ = v1654
	var v1659 int32
	_ = v1659
	var v1664 int32
	_ = v1664
	var v1669 int32
	_ = v1669
	var v1671 int32
	_ = v1671
	var v1679 int32
	_ = v1679
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1693 int32
	_ = v1693
	var v1699 int32
	_ = v1699
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1718 int32
	_ = v1718
	var v1725 int32
	_ = v1725
	var v1727 int32
	_ = v1727
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1739 int32
	_ = v1739
	var v1744 int32
	_ = v1744
	var v1748 int32
	_ = v1748
	var v1751 int32
	_ = v1751
	var v1755 int32
	_ = v1755
	var v1769 int32
	_ = v1769
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1782 int32
	_ = v1782
	var v1784 int32
	_ = v1784
	var v1787 int32
	_ = v1787
	var v1789 int32
	_ = v1789
	var v1793 int32
	_ = v1793
	var v1799 int32
	_ = v1799
	var v1815 int32
	_ = v1815
	var v1819 int32
	_ = v1819
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1839 int32
	_ = v1839
	var v1841 int32
	_ = v1841
	var v1847 int32
	_ = v1847
	var v1849 int32
	_ = v1849
	var v1851 int32
	_ = v1851
	var v1853 int32
	_ = v1853
	var v1866 int32
	_ = v1866
	var v1868 int32
	_ = v1868
	var v1870 int32
	_ = v1870
	var v1888 int32
	_ = v1888
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1901 int32
	_ = v1901
	var v1907 int32
	_ = v1907
	var v1924 int32
	_ = v1924
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1935 int32
	_ = v1935
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1945 int32
	_ = v1945
	var v1947 int32
	_ = v1947
	var v1950 int32
	_ = v1950
	var v1952 int32
	_ = v1952
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1978 int32
	_ = v1978
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2017 int32
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2023 int32
	_ = v2023
	var v2024 int32
	_ = v2024
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2057 int32
	_ = v2057
	var v2059 int32
	_ = v2059
	var v2063 int32
	_ = v2063
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2094 int32
	_ = v2094
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2118 int32
	_ = v2118
	var v2120 int32
	_ = v2120
	var v2126 int32
	_ = v2126
	var v2128 int32
	_ = v2128
	var v2130 int32
	_ = v2130
	var v2132 int32
	_ = v2132
	var v2145 int32
	_ = v2145
	var v2147 int32
	_ = v2147
	var v2149 int32
	_ = v2149
	var v2167 int32
	_ = v2167
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2180 int32
	_ = v2180
	var v2186 int32
	_ = v2186
	var v2202 int32
	_ = v2202
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2216 int32
	_ = v2216
	var v2221 int32
	_ = v2221
	var v2223 int32
	_ = v2223
	var v2226 int32
	_ = v2226
	var v2230 int32
	_ = v2230
	var v2234 int32
	_ = v2234
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2251 int32
	_ = v2251
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2266 int32
	_ = v2266
	var v2267 int32
	_ = v2267
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2286 int32
	_ = v2286
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2294 int32
	_ = v2294
	var v2300 int32
	_ = v2300
	var v2302 int32
	_ = v2302
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2309 int32
	_ = v2309
	var v2311 int32
	_ = v2311
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2328 int32
	_ = v2328
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2335 int32
	_ = v2335
	var v2336 int32
	_ = v2336
	var v2339 int32
	_ = v2339
	var v2341 int32
	_ = v2341
	var v2343 int32
	_ = v2343
	var v2345 int32
	_ = v2345
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2362 int32
	_ = v2362
	var v2366 int32
	_ = v2366
	var v2368 int32
	_ = v2368
	var v2371 int32
	_ = v2371
	var v2373 int32
	_ = v2373
	var v2377 int32
	_ = v2377
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2386 int32
	_ = v2386
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2393 int32
	_ = v2393
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2400 int32
	_ = v2400
	var v2405 int32
	_ = v2405
	var v2410 int32
	_ = v2410
	var v2415 int32
	_ = v2415
	var v2417 int32
	_ = v2417
	var v2425 int32
	_ = v2425
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2434 int32
	_ = v2434
	var v2436 int32
	_ = v2436
	var v2437 int32
	_ = v2437
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2449 int32
	_ = v2449
	var v2451 int32
	_ = v2451
	var v2455 int32
	_ = v2455
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2470 int32
	_ = v2470
	var v2473 int32
	_ = v2473
	var v2476 int32
	_ = v2476
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2490 int32
	_ = v2490
	var v2492 int32
	_ = v2492
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2502 int32
	_ = v2502
	var v2511 int32
	_ = v2511
	var v2513 int32
	_ = v2513
	var v2518 int32
	_ = v2518
	var v2520 int32
	_ = v2520
	var v2527 int32
	_ = v2527
	var v2530 int32
	_ = v2530
	var v2534 int32
	_ = v2534
	var v2541 int32
	_ = v2541
	var v2542 int32
	_ = v2542
	var v2556 int32
	_ = v2556
	var v2561 int32
	_ = v2561
	var v2566 int32
	_ = v2566
	var v2567 int32
	_ = v2567
	var v2585 int32
	_ = v2585
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v7
	v10 = v7 + int32(2)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v11 <= v10 {
		v106 = v11
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v2585
L2:
	;
	v2585 = int32(1)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L46
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v10))))
	if v15&int32(224) != int32(96) {
		v106 = v11
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if int32(1)<<(uint(v15)%32)&int32(42750482) == int32(0) {
		v106 = v11
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v28 = F_find_among(m, l0, int32(4302496), int32(18))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v28 == int32(0) {
		v106 = v32
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v35
	if v35 < v32 {
		v106 = v32
		goto L3
	} else {
		goto L10
	}
L10:
	;
	switch v28 - int32(1) {
	case 0:
		goto L21
	case 1:
		goto L20
	case 2:
		goto L19
	case 3:
		goto L18
	case 4:
		goto L17
	case 5:
		goto L16
	case 6:
		goto L15
	case 7:
		goto L14
	case 8:
		goto L13
	case 9:
		goto L12
	case 10:
		goto L11
	default:
		goto L2
	}
L11:
	;
	v102 = F_slice_from_s(m, l0, int32(5), int32(2231805))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L7
	} else {
		goto L42
	}
L12:
	;
	v96 = F_slice_from_s(m, l0, int32(4), int32(2231801))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L40
	}
L13:
	;
	v90 = F_slice_from_s(m, l0, int32(5), int32(2231796))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L7
	} else {
		goto L38
	}
L14:
	;
	v84 = F_slice_from_s(m, l0, int32(4), int32(2231792))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L7
	} else {
		goto L36
	}
L15:
	;
	v78 = F_slice_from_s(m, l0, int32(5), int32(2231787))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L7
	} else {
		goto L34
	}
L16:
	;
	v72 = F_slice_from_s(m, l0, int32(3), int32(2231784))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L7
	} else {
		goto L32
	}
L17:
	;
	v66 = F_slice_from_s(m, l0, int32(3), int32(2231781))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L30
	}
L18:
	;
	v60 = F_slice_from_s(m, l0, int32(3), int32(2231778))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L7
	} else {
		goto L28
	}
L19:
	;
	v54 = F_slice_from_s(m, l0, int32(3), int32(2231775))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L7
	} else {
		goto L26
	}
L20:
	;
	v48 = F_slice_from_s(m, l0, int32(3), int32(2231772))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L7
	} else {
		goto L24
	}
L21:
	;
	v42 = F_slice_from_s(m, l0, int32(3), int32(2231769))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	if int32(0) <= v42 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v2585 = v42
	goto L1
L24:
	;
	if int32(0) <= v48 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v2585 = v48
	goto L1
L26:
	;
	if int32(0) <= v54 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v2585 = v54
	goto L1
L28:
	;
	if int32(0) <= v60 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	v2585 = v60
	goto L1
L30:
	;
	if int32(0) <= v66 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v2585 = v66
	goto L1
L32:
	;
	if int32(0) <= v72 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v2585 = v72
	goto L1
L34:
	;
	if int32(0) <= v78 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v2585 = v78
	goto L1
L36:
	;
	if int32(0) <= v84 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v2585 = v84
	goto L1
L38:
	;
	if int32(0) <= v90 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v2585 = v90
	goto L1
L40:
	;
	if int32(0) <= v96 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v2585 = v96
	goto L1
L42:
	;
	if int32(0) <= v102 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v2585 = v102
	goto L1
L44:
	;
	v164 = base.B2i32(v162 < int32(0))
	if v162 < int32(0) {
		goto L64
	} else {
		goto L65
	}
L46:
	;
	goto L47
L47:
	;
	goto L48
L48:
	;
	v117 = v7
	v119 = int32(3)
	goto L51
L50:
	;
	v162 = v147
	goto L44
L51:
	;
	if v106 <= v117 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L50
L53:
	;
	v162 = int32(-1)
	goto L44
L54:
	;
	goto L55
L55:
	;
	v124 = v117 + int32(1)
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+v117))))
	if base.Ui32(v126) < base.Ui32(int32(192)) {
		v147 = v124
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v148 = int32(1)
	if v148 < v119 {
		v117 = v147
		v119 = v119 - v148
		goto L51
	} else {
		goto L63
	}
L57:
	;
	if v106 <= v124 {
		v147 = v124
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v133 = v124
	goto L59
L59:
	;
	v136 = int32(*(*int8)(unsafe.Add(mBase, uint32(v110+v133))))
	if int32(-65) < v136 {
		v147 = v133
		goto L56
	} else {
		goto L61
	}
L60:
	;
	v147 = v106
	goto L56
L61:
	;
	v140 = v133 + int32(1)
	if v140 != v106 {
		v133 = v140
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	goto L52
L64:
	;
	v165 = v7
	goto L66
L65:
	;
	v165 = v162
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v165
	if v162 < int32(0) {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v168)+8)) = int32(0)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v171
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v171 == v173 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v171
	v214 = v171
	goto L79
L69:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+v171))))
	if v177 == int32(39) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v181 = v171 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v181
	v184 = F_slice_del(m, l0)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L7
	} else {
		goto L73
	}
L71:
	;
	v189 = v173
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v171
	if v189 == v171 {
		goto L68
	} else {
		goto L75
	}
L73:
	;
	if v184 < int32(0) {
		v2585 = v184
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v189 = v188
	goto L72
L75:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+v171))))
	if v195 != int32(121) {
		goto L68
	} else {
		goto L76
	}
L76:
	;
	v198 = int32(1)
	v199 = v171 + v198
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v199
	v204 = F_slice_from_s(m, l0, v198, int32(2231894))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L7
	} else {
		goto L77
	}
L77:
	;
	if v204 < int32(0) {
		v2585 = v204
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v208)+8)) = int32(1)
	goto L68
L79:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L86
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v171
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v422))) = v423
	*(*int32)(unsafe.Add(mBase, uint32(v422)+4)) = v423
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v428 = v426 + int32(4)
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v429 <= v428 {
		goto L140
	} else {
		goto L141
	}
L81:
	;
	goto L80
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v214
	v408 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v339 + v408
	v413 = F_slice_from_s(m, l0, v408, int32(2231899))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L7
	} else {
		goto L136
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v214
	goto L117
L84:
	;
	if v336 != 0 {
		goto L108
	} else {
		goto L109
	}
L85:
	;
	v336 = v329
	goto L84
L86:
	;
	if v231 <= v230 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v329 = int32(0)
	goto L85
L88:
	;
	v336 = int32(-1)
	goto L84
L89:
	;
	goto L90
L90:
	;
	v247 = int32(1)
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v232))))
	if base.Ui32(v249) < base.Ui32(int32(192)) {
		v306 = v249
		v307 = v247
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if int32(121) < v306 {
		v329 = v307
		goto L85
	} else {
		goto L104
	}
L92:
	;
	v253 = v230 + int32(1)
	if v253 == v231 {
		v306 = v249
		v307 = v247
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253+v232))))
	v258 = v256 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v249) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262+v232))))
	v274 = v272 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v249) {
		goto L100
	} else {
		goto L101
	}
L95:
	;
	v262 = v230 + int32(2)
	if v262 != v231 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v306 = v249<<(uint(int32(6))%32)&int32(1984) | v258
	v307 = int32(2)
	goto L91
L98:
	;
	goto L97
L99:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232+v278))))
	v306 = v291&int32(63) | (v249<<(uint(int32(18))%32)&int32(1835008) | v258<<(uint(int32(12))%32) | v274<<(uint(int32(6))%32))
	v307 = int32(4)
	goto L91
L100:
	;
	v278 = v230 + int32(3)
	if v278 != v231 {
		goto L99
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v306 = v249<<(uint(int32(12))%32)&int32(61440) | v258<<(uint(int32(6))%32) | v274
	v307 = int32(3)
	goto L91
L103:
	;
	goto L102
L104:
	;
	v311 = v306 - int32(97)
	if v311 < int32(0) {
		v329 = v307
		goto L85
	} else {
		goto L105
	}
L105:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v311)>>(uint(int32(3))%32)))+uint32(_consts[1303]))))
	if int32(base.Ui32(v317)>>(uint(v311&int32(7))%32))&int32(1) == int32(0) {
		v329 = v307
		goto L85
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v307 + v230
	goto L107
L107:
	;
	goto L87
L108:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v348 = v337
	v350 = v338
	goto L83
L109:
	;
	goto L110
L110:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v339
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v342 == v339 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v348 = v339
	v350 = v341
	goto L83
L112:
	;
	goto L113
L113:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339+v341))))
	if v345 == int32(121) {
		goto L82
	} else {
		goto L114
	}
L114:
	;
	v348 = v342
	v350 = v341
	goto L83
L115:
	;
	if v403 < int32(0) {
		goto L81
	} else {
		goto L135
	}
L117:
	;
	goto L118
L118:
	;
	goto L119
L119:
	;
	v358 = v214
	v360 = int32(1)
	goto L122
L121:
	;
	v403 = v388
	goto L115
L122:
	;
	if v348 <= v358 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	goto L121
L124:
	;
	v403 = int32(-1)
	goto L115
L125:
	;
	goto L126
L126:
	;
	v365 = v358 + int32(1)
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350+v358))))
	if base.Ui32(v367) < base.Ui32(int32(192)) {
		v388 = v365
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v389 = int32(1)
	if v389 < v360 {
		v358 = v388
		v360 = v360 - v389
		goto L122
	} else {
		goto L134
	}
L128:
	;
	if v348 <= v365 {
		v388 = v365
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v374 = v365
	goto L130
L130:
	;
	v377 = int32(*(*int8)(unsafe.Add(mBase, uint32(v350+v374))))
	if int32(-65) < v377 {
		v388 = v374
		goto L127
	} else {
		goto L132
	}
L131:
	;
	v388 = v348
	goto L127
L132:
	;
	v381 = v374 + int32(1)
	if v381 != v348 {
		v374 = v381
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	goto L123
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v403
	v214 = v403
	goto L79
L136:
	;
	if v413 < int32(0) {
		v2585 = v413
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v417)+8)) = int32(1)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v214 = v420
	goto L79
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v426
	v949 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v949
	if v949 <= v426 {
		v978 = v949
		goto L248
	} else {
		goto L249
	}
L139:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v699)+4)) = v698
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v722 = v712
	goto L199
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v426
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v474 = v426
	goto L148
L141:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431+v428))))
	if v433&int32(224) != int32(96) {
		goto L140
	} else {
		goto L142
	}
L142:
	;
	if int32(1)<<(uint(v433)%32)&int32(2375680) == int32(0) {
		goto L140
	} else {
		goto L143
	}
L143:
	;
	v446 = F_find_among(m, l0, int32(4302864), int32(3))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L7
	} else {
		goto L144
	}
L144:
	;
	if v446 == int32(0) {
		goto L140
	} else {
		goto L145
	}
L145:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v698 = v450
	goto L139
L146:
	;
	if v569 < int32(0) {
		goto L138
	} else {
		goto L171
	}
L147:
	;
	v569 = v541
	goto L146
L148:
	;
	if v465 <= v474 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v569 = int32(-1)
	goto L146
L151:
	;
	goto L152
L152:
	;
	v481 = int32(1)
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474+v466))))
	if base.Ui32(v483) < base.Ui32(int32(192)) {
		v540 = v483
		v541 = v481
		goto L153
	} else {
		goto L154
	}
L153:
	;
	if int32(121) < v540 {
		goto L166
	} else {
		goto L167
	}
L154:
	;
	v487 = v474 + int32(1)
	if v487 == v465 {
		v540 = v483
		v541 = v481
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487+v466))))
	v492 = v490 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v483) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496+v466))))
	v508 = v506 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v483) {
		goto L162
	} else {
		goto L163
	}
L157:
	;
	v496 = v474 + int32(2)
	if v496 != v465 {
		goto L156
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v540 = v483<<(uint(int32(6))%32)&int32(1984) | v492
	v541 = int32(2)
	goto L153
L160:
	;
	goto L159
L161:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v466+v512))))
	v540 = v525&int32(63) | (v483<<(uint(int32(18))%32)&int32(1835008) | v492<<(uint(int32(12))%32) | v508<<(uint(int32(6))%32))
	v541 = int32(4)
	goto L153
L162:
	;
	v512 = v474 + int32(3)
	if v512 != v465 {
		goto L161
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v540 = v483<<(uint(int32(12))%32)&int32(61440) | v492<<(uint(int32(6))%32) | v508
	v541 = int32(3)
	goto L153
L165:
	;
	goto L164
L166:
	;
	v558 = v541 + v474
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v558
	v474 = v558
	goto L148
L167:
	;
	v545 = v540 - int32(97)
	if v545 < int32(0) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v545)>>(uint(int32(3))%32)))+uint32(_consts[1303]))))
	if int32(base.Ui32(v551)>>(uint(v545&int32(7))%32))&int32(1) != 0 {
		goto L147
	} else {
		goto L169
	}
L169:
	;
	goto L166
L171:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v573 = v572 + v569
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v573
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v596 = v573
	goto L174
L172:
	;
	if v692 < int32(0) {
		goto L138
	} else {
		goto L196
	}
L173:
	;
	v692 = v663
	goto L172
L174:
	;
	if v587 <= v596 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v692 = int32(-1)
	goto L172
L177:
	;
	goto L178
L178:
	;
	v603 = int32(1)
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596+v588))))
	if base.Ui32(v605) < base.Ui32(int32(192)) {
		v662 = v605
		v663 = v603
		goto L179
	} else {
		goto L180
	}
L179:
	;
	if int32(121) < v662 {
		goto L173
	} else {
		goto L192
	}
L180:
	;
	v609 = v596 + int32(1)
	if v609 == v587 {
		v662 = v605
		v663 = v603
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609+v588))))
	v614 = v612 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v605) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618+v588))))
	v630 = v628 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v605) {
		goto L188
	} else {
		goto L189
	}
L183:
	;
	v618 = v596 + int32(2)
	if v618 != v587 {
		goto L182
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v662 = v605<<(uint(int32(6))%32)&int32(1984) | v614
	v663 = int32(2)
	goto L179
L186:
	;
	goto L185
L187:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v588+v634))))
	v662 = v647&int32(63) | (v605<<(uint(int32(18))%32)&int32(1835008) | v614<<(uint(int32(12))%32) | v630<<(uint(int32(6))%32))
	v663 = int32(4)
	goto L179
L188:
	;
	v634 = v596 + int32(3)
	if v634 != v587 {
		goto L187
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v662 = v605<<(uint(int32(12))%32)&int32(61440) | v614<<(uint(int32(6))%32) | v630
	v663 = int32(3)
	goto L179
L191:
	;
	goto L190
L192:
	;
	v667 = v662 - int32(97)
	if v667 < int32(0) {
		goto L173
	} else {
		goto L193
	}
L193:
	;
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v667)>>(uint(int32(3))%32)))+uint32(_consts[1303]))))
	if int32(base.Ui32(v673)>>(uint(v667&int32(7))%32))&int32(1) == int32(0) {
		goto L173
	} else {
		goto L194
	}
L194:
	;
	v681 = v663 + v596
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v681
	v596 = v681
	goto L174
L196:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v696 = v695 + v692
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v696
	v698 = v696
	goto L139
L197:
	;
	if v817 < int32(0) {
		goto L138
	} else {
		goto L222
	}
L198:
	;
	v817 = v789
	goto L197
L199:
	;
	if v713 <= v722 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v817 = int32(-1)
	goto L197
L202:
	;
	goto L203
L203:
	;
	v729 = int32(1)
	v731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722+v714))))
	if base.Ui32(v731) < base.Ui32(int32(192)) {
		v788 = v731
		v789 = v729
		goto L204
	} else {
		goto L205
	}
L204:
	;
	if int32(121) < v788 {
		goto L217
	} else {
		goto L218
	}
L205:
	;
	v735 = v722 + int32(1)
	if v735 == v713 {
		v788 = v731
		v789 = v729
		goto L204
	} else {
		goto L206
	}
L206:
	;
	v738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v735+v714))))
	v740 = v738 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v731) {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744+v714))))
	v756 = v754 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v731) {
		goto L213
	} else {
		goto L214
	}
L208:
	;
	v744 = v722 + int32(2)
	if v744 != v713 {
		goto L207
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v788 = v731<<(uint(int32(6))%32)&int32(1984) | v740
	v789 = int32(2)
	goto L204
L211:
	;
	goto L210
L212:
	;
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v714+v760))))
	v788 = v773&int32(63) | (v731<<(uint(int32(18))%32)&int32(1835008) | v740<<(uint(int32(12))%32) | v756<<(uint(int32(6))%32))
	v789 = int32(4)
	goto L204
L213:
	;
	v760 = v722 + int32(3)
	if v760 != v713 {
		goto L212
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v788 = v731<<(uint(int32(12))%32)&int32(61440) | v740<<(uint(int32(6))%32) | v756
	v789 = int32(3)
	goto L204
L216:
	;
	goto L215
L217:
	;
	v806 = v789 + v722
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v806
	v722 = v806
	goto L199
L218:
	;
	v793 = v788 - int32(97)
	if v793 < int32(0) {
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v793)>>(uint(int32(3))%32)))+uint32(_consts[1303]))))
	if int32(base.Ui32(v799)>>(uint(v793&int32(7))%32))&int32(1) != 0 {
		goto L198
	} else {
		goto L220
	}
L220:
	;
	goto L217
L222:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v821 = v820 + v817
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v821
	v835 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v844 = v821
	goto L225
L223:
	;
	if v940 < int32(0) {
		goto L138
	} else {
		goto L247
	}
L224:
	;
	v940 = v911
	goto L223
L225:
	;
	if v835 <= v844 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v940 = int32(-1)
	goto L223
L228:
	;
	goto L229
L229:
	;
	v851 = int32(1)
	v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v844+v836))))
	if base.Ui32(v853) < base.Ui32(int32(192)) {
		v910 = v853
		v911 = v851
		goto L230
	} else {
		goto L231
	}
L230:
	;
	if int32(121) < v910 {
		goto L224
	} else {
		goto L243
	}
L231:
	;
	v857 = v844 + int32(1)
	if v857 == v835 {
		v910 = v853
		v911 = v851
		goto L230
	} else {
		goto L232
	}
L232:
	;
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v857+v836))))
	v862 = v860 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v853) {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	v876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v866+v836))))
	v878 = v876 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v853) {
		goto L239
	} else {
		goto L240
	}
L234:
	;
	v866 = v844 + int32(2)
	if v866 != v835 {
		goto L233
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v910 = v853<<(uint(int32(6))%32)&int32(1984) | v862
	v911 = int32(2)
	goto L230
L237:
	;
	goto L236
L238:
	;
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v836+v882))))
	v910 = v895&int32(63) | (v853<<(uint(int32(18))%32)&int32(1835008) | v862<<(uint(int32(12))%32) | v878<<(uint(int32(6))%32))
	v911 = int32(4)
	goto L230
L239:
	;
	v882 = v844 + int32(3)
	if v882 != v835 {
		goto L238
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v910 = v853<<(uint(int32(12))%32)&int32(61440) | v862<<(uint(int32(6))%32) | v878
	v911 = int32(3)
	goto L230
L242:
	;
	goto L241
L243:
	;
	v915 = v910 - int32(97)
	if v915 < int32(0) {
		goto L224
	} else {
		goto L244
	}
L244:
	;
	v921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v915)>>(uint(int32(3))%32)))+uint32(_consts[1303]))))
	if int32(base.Ui32(v921)>>(uint(v915&int32(7))%32))&int32(1) == int32(0) {
		goto L224
	} else {
		goto L245
	}
L245:
	;
	v929 = v911 + v844
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v929
	v844 = v929
	goto L225
L247:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v943))) = v944 + v940
	goto L138
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v978
	v981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v978 <= v981 {
		goto L257
	} else {
		goto L258
	}
L249:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v953+v949-int32(1)))))
	if base.B2i32(v957 != int32(115))&base.B2i32(v957 != int32(39)) != 0 {
		v978 = v949
		goto L248
	} else {
		goto L250
	}
L250:
	;
	v965 = F_find_among_b(m, l0, int32(4302928), int32(3))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L7
	} else {
		goto L251
	}
L251:
	;
	if v965 == int32(0) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v969
	v978 = v969
	goto L248
L253:
	;
	goto L254
L254:
	;
	v971 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v971
	v973 = F_slice_del(m, l0)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L7
	} else {
		goto L255
	}
L255:
	;
	if v973 < int32(0) {
		v2585 = v973
		goto L1
	} else {
		goto L256
	}
L256:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v978 = v977
	goto L248
L257:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1277
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1277
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1277-int32(5) <= v1280 {
		v1302 = v1280
		goto L338
	} else {
		goto L339
	}
L258:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v983+v978-int32(1)))))
	switch v987 - int32(100) {
	case 0, 15:
		goto L259
	default:
		goto L257
	}
L259:
	;
	v992 = F_find_among_b(m, l0, int32(4302992), int32(6))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L7
	} else {
		goto L260
	}
L260:
	;
	if v992 == int32(0) {
		goto L257
	} else {
		goto L261
	}
L261:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v996
	switch v992 - int32(1) {
	case 0:
		goto L264
	case 1:
		goto L263
	case 2:
		goto L262
	default:
		goto L257
	}
L262:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L296
L263:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L269
L264:
	;
	v1002 = F_slice_from_s(m, l0, int32(2), int32(2231916))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L7
	} else {
		goto L265
	}
L265:
	;
	if int32(0) <= v1002 {
		goto L257
	} else {
		goto L266
	}
L266:
	;
	v2585 = v1002
	goto L1
L267:
	;
	if int32(0) <= v1059 {
		goto L287
	} else {
		goto L288
	}
L269:
	;
	goto L270
L270:
	;
	goto L271
L271:
	;
	v1015 = v996
	v1017 = int32(2)
	goto L274
L273:
	;
	v1059 = v1041
	goto L267
L274:
	;
	if v1015 <= v1008 {
		goto L276
	} else {
		goto L277
	}
L275:
	;
	goto L273
L276:
	;
	v1059 = int32(-1)
	goto L267
L277:
	;
	goto L278
L278:
	;
	v1022 = v1015 - int32(1)
	v1024 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1007+v1022))))
	if int32(0) <= v1024 {
		v1041 = v1022
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v1045 = int32(1)
	if v1045 < v1017 {
		v1015 = v1041
		v1017 = v1017 - v1045
		goto L274
	} else {
		goto L286
	}
L280:
	;
	if v1022 <= v1008 {
		v1041 = v1022
		goto L279
	} else {
		goto L281
	}
L281:
	;
	v1029 = v1022
	goto L282
L282:
	;
	v1034 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1007+v1029))))
	if base.Ui32(int32(191)) < base.Ui32(v1034) {
		v1041 = v1029
		goto L279
	} else {
		goto L284
	}
L283:
	;
	v1041 = v1008
	goto L279
L284:
	;
	v1038 = v1029 - int32(1)
	if v1008 < v1038 {
		v1029 = v1038
		goto L282
	} else {
		goto L285
	}
L285:
	;
	goto L283
L286:
	;
	goto L275
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1059
	v1065 = F_slice_from_s(m, l0, int32(1), int32(2231918))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L7
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1069 + (v996 - v1006)
	v1075 = F_slice_from_s(m, l0, int32(2), int32(2231919))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L7
	} else {
		goto L292
	}
L290:
	;
	if int32(0) <= v1065 {
		goto L257
	} else {
		goto L291
	}
L291:
	;
	v2585 = v1065
	goto L1
L292:
	;
	if int32(0) <= v1075 {
		goto L257
	} else {
		goto L293
	}
L293:
	;
	v2585 = v1075
	goto L1
L294:
	;
	if v1131 < int32(0) {
		goto L257
	} else {
		goto L314
	}
L296:
	;
	goto L297
L297:
	;
	goto L298
L298:
	;
	v1087 = v996
	v1089 = int32(1)
	goto L301
L300:
	;
	v1131 = v1113
	goto L294
L301:
	;
	if v1087 <= v1080 {
		goto L303
	} else {
		goto L304
	}
L302:
	;
	goto L300
L303:
	;
	v1131 = int32(-1)
	goto L294
L304:
	;
	goto L305
L305:
	;
	v1094 = v1087 - int32(1)
	v1096 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1079+v1094))))
	if int32(0) <= v1096 {
		v1113 = v1094
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1117 = int32(1)
	if v1117 < v1089 {
		v1087 = v1113
		v1089 = v1089 - v1117
		goto L301
	} else {
		goto L313
	}
L307:
	;
	if v1094 <= v1080 {
		v1113 = v1094
		goto L306
	} else {
		goto L308
	}
L308:
	;
	v1101 = v1094
	goto L309
L309:
	;
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1079+v1101))))
	if base.Ui32(int32(191)) < base.Ui32(v1106) {
		v1113 = v1101
		goto L306
	} else {
		goto L311
	}
L310:
	;
	v1113 = v1080
	goto L306
L311:
	;
	v1110 = v1101 - int32(1)
	if v1080 < v1110 {
		v1101 = v1110
		goto L309
	} else {
		goto L312
	}
L312:
	;
	goto L310
L313:
	;
	goto L302
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1131
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1161 = v1131
	goto L317
L315:
	;
	if v1264 < int32(0) {
		goto L257
	} else {
		goto L334
	}
L316:
	;
	v1264 = int32(-1)
	goto L315
L317:
	;
	if v1161 <= v1152 {
		goto L316
	} else {
		goto L319
	}
L319:
	;
	v1169 = int32(1)
	v1170 = v1161 - v1169
	v1172 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1148+v1170))))
	v1174 = v1172 & int32(255)
	if v1170 == v1152 {
		v1229 = v1174
		v1230 = v1169
		goto L320
	} else {
		goto L321
	}
L320:
	;
	if int32(121) < v1229 {
		goto L329
	} else {
		goto L330
	}
L321:
	;
	if int32(0) <= v1172 {
		v1229 = v1174
		v1230 = v1169
		goto L320
	} else {
		goto L322
	}
L322:
	;
	v1180 = v1174 & int32(63)
	v1182 = v1161 - int32(2)
	v1184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1148+v1182))))
	v1186 = v1184 << (uint(int32(6)) % 32)
	if base.B2i32(v1182 != v1152)&base.B2i32(base.Ui32(v1184) < base.Ui32(int32(192))) == int32(0) {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v1229 = v1186&int32(1984) | v1180
	v1230 = int32(2)
	goto L320
L324:
	;
	goto L325
L325:
	;
	v1199 = v1186&int32(4032) | v1180
	v1201 = v1161 - int32(3)
	v1203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1148+v1201))))
	if base.B2i32(v1201 != v1152)&base.B2i32(base.Ui32(v1203) < base.Ui32(int32(224))) == int32(0) {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v1229 = v1203<<(uint(int32(12))%32)&int32(61440) | v1199
	v1230 = int32(3)
	goto L320
L327:
	;
	goto L328
L328:
	;
	v1221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1161+(v1148-int32(4))))))
	v1229 = v1203<<(uint(int32(12))%32)&int32(258048) | v1221&int32(7)<<(uint(int32(18))%32) | v1199
	v1230 = int32(4)
	goto L320
L329:
	;
	v1249 = v1161 - v1230
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1249
	v1161 = v1249
	goto L317
L330:
	;
	v1234 = v1229 - int32(97)
	if v1234 < int32(0) {
		goto L329
	} else {
		goto L331
	}
L331:
	;
	v1240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1234)>>(uint(int32(3))%32)))+uint32(_consts[1303]))))
	if int32(base.Ui32(v1240)>>(uint(v1234&int32(7))%32))&int32(1) == int32(0) {
		goto L329
	} else {
		goto L332
	}
L332:
	;
	v1264 = v1230
	goto L315
L334:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1267 - v1264
	v1270 = F_slice_del(m, l0)
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L7
	} else {
		goto L335
	}
L335:
	;
	if v1270 < int32(0) {
		v2585 = v1270
		goto L1
	} else {
		goto L336
	}
L336:
	;
	goto L257
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2476
	v2480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2481 = *(*int32)(unsafe.Add(mBase, uint32(v2480)+8))
	if v2481 == int32(0) {
		goto L632
	} else {
		goto L633
	}
L338:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1303
	v1307 = v1303 - int32(1)
	if v1307 <= v1302 {
		goto L344
	} else {
		goto L345
	}
L339:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1284+v1277-int32(1)))))
	switch v1288 - int32(100) {
	case 0, 3:
		goto L340
	default:
		v1302 = v1280
		goto L338
	}
L340:
	;
	v1293 = F_find_among_b(m, l0, int32(4303120), int32(8))
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L7
	} else {
		goto L341
	}
L341:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1293 == int32(0) {
		v1302 = v1295
		goto L338
	} else {
		goto L342
	}
L342:
	;
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1298
	if v1298 <= v1295 {
		v2476 = v1295
		goto L337
	} else {
		goto L343
	}
L343:
	;
	v1302 = v1295
	goto L338
L344:
	;
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1782
	v1784 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1782
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1782 <= v1787 {
		v1947 = v1784
		goto L448
	} else {
		goto L449
	}
L345:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1309+v1307))))
	if v1311&int32(224) != int32(96) {
		goto L344
	} else {
		goto L346
	}
L346:
	;
	if int32(1)<<(uint(v1311)%32)&int32(33554576) == int32(0) {
		goto L344
	} else {
		goto L347
	}
L347:
	;
	v1324 = F_find_among_b(m, l0, int32(4303280), int32(6))
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L7
	} else {
		goto L348
	}
L348:
	;
	if v1324 == int32(0) {
		goto L344
	} else {
		goto L349
	}
L349:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1328
	switch v1324 - int32(1) {
	case 0:
		goto L351
	case 1:
		goto L350
	default:
		goto L344
	}
L350:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1368 = v1358
	goto L357
L351:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+4))
	if v1328 < v1333 {
		goto L344
	} else {
		goto L352
	}
L352:
	;
	v1337 = F_slice_from_s(m, l0, int32(2), int32(2231995))
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L7
	} else {
		goto L353
	}
L353:
	;
	if int32(0) <= v1337 {
		goto L344
	} else {
		goto L354
	}
L354:
	;
	v2585 = v1337
	goto L1
L355:
	;
	if v1471 < int32(0) {
		goto L344
	} else {
		goto L374
	}
L356:
	;
	v1471 = int32(-1)
	goto L355
L357:
	;
	if v1368 <= v1359 {
		goto L356
	} else {
		goto L359
	}
L359:
	;
	v1376 = int32(1)
	v1377 = v1368 - v1376
	v1379 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1355+v1377))))
	v1381 = v1379 & int32(255)
	if v1377 == v1359 {
		v1436 = v1381
		v1437 = v1376
		goto L360
	} else {
		goto L361
	}
L360:
	;
	if int32(121) < v1436 {
		goto L369
	} else {
		goto L370
	}
L361:
	;
	if int32(0) <= v1379 {
		v1436 = v1381
		v1437 = v1376
		goto L360
	} else {
		goto L362
	}
L362:
	;
	v1387 = v1381 & int32(63)
	v1389 = v1368 - int32(2)
	v1391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1355+v1389))))
	v1393 = v1391 << (uint(int32(6)) % 32)
	if base.B2i32(v1389 != v1359)&base.B2i32(base.Ui32(v1391) < base.Ui32(int32(192))) == int32(0) {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1436 = v1393&int32(1984) | v1387
	v1437 = int32(2)
	goto L360
L364:
	;
	goto L365
L365:
	;
	v1406 = v1393&int32(4032) | v1387
	v1408 = v1368 - int32(3)
	v1410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1355+v1408))))
	if base.B2i32(v1408 != v1359)&base.B2i32(base.Ui32(v1410) < base.Ui32(int32(224))) == int32(0) {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v1436 = v1410<<(uint(int32(12))%32)&int32(61440) | v1406
	v1437 = int32(3)
	goto L360
L367:
	;
	goto L368
L368:
	;
	v1428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1368+(v1355-int32(4))))))
	v1436 = v1410<<(uint(int32(12))%32)&int32(258048) | v1428&int32(7)<<(uint(int32(18))%32) | v1406
	v1437 = int32(4)
	goto L360
L369:
	;
	v1456 = v1368 - v1437
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1456
	v1368 = v1456
	goto L357
L370:
	;
	v1441 = v1436 - int32(97)
	if v1441 < int32(0) {
		goto L369
	} else {
		goto L371
	}
L371:
	;
	v1447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1441)>>(uint(int32(3))%32)))+uint32(_consts[1303]))))
	if int32(base.Ui32(v1447)>>(uint(v1441&int32(7))%32))&int32(1) == int32(0) {
		goto L369
	} else {
		goto L372
	}
L372:
	;
	v1471 = v1437
	goto L355
L374:
	;
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1474 + (v1328 - v1341)
	v1478 = F_slice_del(m, l0)
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L7
	} else {
		goto L375
	}
L375:
	;
	if v1478 < int32(0) {
		v2585 = v1478
		goto L1
	} else {
		goto L376
	}
L376:
	;
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1482
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1482
	v1486 = v1482 - int32(1)
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1486 <= v1487 {
		v1643 = v1482
		goto L380
	} else {
		goto L381
	}
L377:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1715 = v1713 + (v1482 - v1502)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1715
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1715
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L427
L378:
	;
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1711 = v1710
	goto L377
L379:
	;
	v1702 = F_slice_from_s(m, l0, int32(1), v1699)
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		goto L7
	} else {
		goto L423
	}
L380:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v1646)+4))
	if v1643 != v1647 {
		goto L344
	} else {
		goto L413
	}
L381:
	;
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1489+v1486))))
	if v1491&int32(224) != int32(96) {
		v1643 = v1482
		goto L380
	} else {
		goto L382
	}
L382:
	;
	if int32(1)<<(uint(v1491)%32)&int32(68514004) == int32(0) {
		v1643 = v1482
		goto L380
	} else {
		goto L383
	}
L383:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1506 = F_find_among_b(m, l0, int32(4303408), int32(13))
	mBase = m.M
	v1507 = m.ExcPending
	if v1507 != 0 {
		goto L7
	} else {
		goto L386
	}
L384:
	;
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1643 = v1642
	goto L380
L385:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L389
L386:
	;
	switch v1506 - int32(1) {
	case 0:
		v1699 = int32(2231997)
		goto L379
	case 1:
		goto L385
	case 2:
		goto L384
	default:
		goto L378
	}
L387:
	;
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1638 != 0 {
		v1711 = v1639
		goto L377
	} else {
		goto L411
	}
L388:
	;
	v1638 = v1631
	goto L387
L389:
	;
	if v1526 <= v1527 {
		v1631 = int32(-1)
		goto L388
	} else {
		goto L391
	}
L390:
	;
	v1631 = int32(0)
	goto L388
L391:
	;
	v1544 = int32(1)
	v1545 = v1526 - v1544
	v1547 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1523+v1545))))
	v1549 = v1547 & int32(255)
	if v1545 == v1527 {
		v1604 = v1549
		v1605 = v1544
		goto L392
	} else {
		goto L393
	}
L392:
	;
	if int32(111) < v1604 {
		goto L401
	} else {
		goto L402
	}
L393:
	;
	if int32(0) <= v1547 {
		v1604 = v1549
		v1605 = v1544
		goto L392
	} else {
		goto L394
	}
L394:
	;
	v1555 = v1549 & int32(63)
	v1557 = v1526 - int32(2)
	v1559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1523+v1557))))
	v1561 = v1559 << (uint(int32(6)) % 32)
	if base.B2i32(v1557 != v1527)&base.B2i32(base.Ui32(v1559) < base.Ui32(int32(192))) == int32(0) {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	v1604 = v1561&int32(1984) | v1555
	v1605 = int32(2)
	goto L392
L396:
	;
	goto L397
L397:
	;
	v1574 = v1561&int32(4032) | v1555
	v1576 = v1526 - int32(3)
	v1578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1523+v1576))))
	if base.B2i32(v1576 != v1527)&base.B2i32(base.Ui32(v1578) < base.Ui32(int32(224))) == int32(0) {
		goto L398
	} else {
		goto L399
	}
L398:
	;
	v1604 = v1578<<(uint(int32(12))%32)&int32(61440) | v1574
	v1605 = int32(3)
	goto L392
L399:
	;
	goto L400
L400:
	;
	v1596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1526+(v1523-int32(4))))))
	v1604 = v1578<<(uint(int32(12))%32)&int32(258048) | v1596&int32(7)<<(uint(int32(18))%32) | v1574
	v1605 = int32(4)
	goto L392
L401:
	;
	v1638 = v1605
	goto L387
L402:
	;
	goto L403
L403:
	;
	v1609 = v1604 - int32(97)
	if v1609 < int32(0) {
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v1638 = v1605
	goto L387
L405:
	;
	goto L406
L406:
	;
	v1615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1609)>>(uint(int32(3))%32)))+uint32(_consts[1304]))))
	if int32(base.Ui32(v1615)>>(uint(v1609&int32(7))%32))&int32(1) == int32(0) {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	v1638 = v1605
	goto L387
L408:
	;
	goto L409
L409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1526 - v1605
	goto L410
L410:
	;
	goto L390
L411:
	;
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1639 < v1640 {
		v1711 = v1639
		goto L377
	} else {
		goto L412
	}
L412:
	;
	goto L344
L413:
	;
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1650 = int32(0)
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1659 = F_out_grouping_b_U(m, l0, int32(2232047), int32(89), int32(121), v1650)
	mBase = m.M
	if v1659 != 0 {
		goto L415
	} else {
		goto L416
	}
L414:
	;
	if v1690 == int32(0) {
		goto L344
	} else {
		goto L422
	}
L415:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1671 + (v1654 - v1649)
	v1679 = F_out_grouping_b_U(m, l0, int32(2231895), int32(97), int32(121), int32(0))
	mBase = m.M
	if v1679 != 0 {
		v1688 = v1650
		goto L419
	} else {
		goto L420
	}
L416:
	;
	v1664 = F_in_grouping_b_U(m, l0, int32(2231895), int32(97), int32(121), int32(0))
	mBase = m.M
	if v1664 != 0 {
		goto L415
	} else {
		goto L417
	}
L417:
	;
	v1669 = F_out_grouping_b_U(m, l0, int32(2231895), int32(97), int32(121), int32(0))
	mBase = m.M
	if v1669 != 0 {
		goto L415
	} else {
		goto L418
	}
L418:
	;
	v1690 = int32(1)
	goto L414
L419:
	;
	v1690 = v1688
	goto L414
L420:
	;
	v1684 = F_in_grouping_b_U(m, l0, int32(2231895), int32(97), int32(121), int32(0))
	mBase = m.M
	if v1684 != 0 {
		v1688 = v1650
		goto L419
	} else {
		goto L421
	}
L421:
	;
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1688 = base.B2i32(v1685 <= v1686)
	goto L419
L422:
	;
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1693 + (v1643 - v1649)
	v1699 = int32(2232000)
	goto L379
L423:
	;
	if int32(0) <= v1702 {
		goto L344
	} else {
		goto L424
	}
L424:
	;
	return v1702 >> (uint(int32(31)) % 32) & v1702
L425:
	;
	if v1769 < int32(0) {
		goto L344
	} else {
		goto L445
	}
L427:
	;
	goto L428
L428:
	;
	goto L429
L429:
	;
	v1725 = v1715
	v1727 = int32(1)
	goto L432
L431:
	;
	v1769 = v1751
	goto L425
L432:
	;
	if v1725 <= v1711 {
		goto L434
	} else {
		goto L435
	}
L433:
	;
	goto L431
L434:
	;
	v1769 = int32(-1)
	goto L425
L435:
	;
	goto L436
L436:
	;
	v1732 = v1725 - int32(1)
	v1734 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1718+v1732))))
	if int32(0) <= v1734 {
		v1751 = v1732
		goto L437
	} else {
		goto L438
	}
L437:
	;
	v1755 = int32(1)
	if v1755 < v1727 {
		v1725 = v1751
		v1727 = v1727 - v1755
		goto L432
	} else {
		goto L444
	}
L438:
	;
	if v1732 <= v1711 {
		v1751 = v1732
		goto L437
	} else {
		goto L439
	}
L439:
	;
	v1739 = v1732
	goto L440
L440:
	;
	v1744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1718+v1739))))
	if base.Ui32(int32(191)) < base.Ui32(v1744) {
		v1751 = v1739
		goto L437
	} else {
		goto L442
	}
L441:
	;
	v1751 = v1711
	goto L437
L442:
	;
	v1748 = v1739 - int32(1)
	if v1711 < v1748 {
		v1739 = v1748
		goto L440
	} else {
		goto L443
	}
L443:
	;
	goto L441
L444:
	;
	goto L433
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1769
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1769
	v1774 = F_slice_del(m, l0)
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L7
	} else {
		goto L446
	}
L446:
	;
	if v1774 < int32(0) {
		v2585 = v1774
		goto L1
	} else {
		goto L447
	}
L447:
	;
	goto L344
L448:
	;
	if v1947 < int32(0) {
		v2585 = v1947
		goto L1
	} else {
		goto L476
	}
L449:
	;
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1789+v1782-int32(1)))))
	if v1793|int32(32) != int32(121) {
		v1947 = v1784
		goto L448
	} else {
		goto L450
	}
L450:
	;
	v1799 = v1782 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1799
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1799
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L453
L451:
	;
	if v1931 != 0 {
		v1947 = v1784
		goto L448
	} else {
		goto L470
	}
L452:
	;
	v1931 = v1924
	goto L451
L453:
	;
	if v1799 <= v1819 {
		v1924 = int32(-1)
		goto L452
	} else {
		goto L455
	}
L454:
	;
	v1924 = int32(0)
	goto L452
L455:
	;
	v1836 = int32(1)
	v1837 = v1799 - v1836
	v1839 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1815+v1837))))
	v1841 = v1839 & int32(255)
	if v1837 == v1819 {
		v1896 = v1841
		v1897 = v1836
		goto L456
	} else {
		goto L457
	}
L456:
	;
	if int32(121) < v1896 {
		goto L465
	} else {
		goto L466
	}
L457:
	;
	if int32(0) <= v1839 {
		v1896 = v1841
		v1897 = v1836
		goto L456
	} else {
		goto L458
	}
L458:
	;
	v1847 = v1841 & int32(63)
	v1849 = v1799 - int32(2)
	v1851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1815+v1849))))
	v1853 = v1851 << (uint(int32(6)) % 32)
	if base.B2i32(v1849 != v1819)&base.B2i32(base.Ui32(v1851) < base.Ui32(int32(192))) == int32(0) {
		goto L459
	} else {
		goto L460
	}
L459:
	;
	v1896 = v1853&int32(1984) | v1847
	v1897 = int32(2)
	goto L456
L460:
	;
	goto L461
L461:
	;
	v1866 = v1853&int32(4032) | v1847
	v1868 = v1799 - int32(3)
	v1870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1815+v1868))))
	if base.B2i32(v1868 != v1819)&base.B2i32(base.Ui32(v1870) < base.Ui32(int32(224))) == int32(0) {
		goto L462
	} else {
		goto L463
	}
L462:
	;
	v1896 = v1870<<(uint(int32(12))%32)&int32(61440) | v1866
	v1897 = int32(3)
	goto L456
L463:
	;
	goto L464
L464:
	;
	v1888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1799+(v1815-int32(4))))))
	v1896 = v1870<<(uint(int32(12))%32)&int32(258048) | v1888&int32(7)<<(uint(int32(18))%32) | v1866
	v1897 = int32(4)
	goto L456
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1799 - v1897
	goto L469
L466:
	;
	v1901 = v1896 - int32(97)
	if v1901 < int32(0) {
		goto L465
	} else {
		goto L467
	}
L467:
	;
	v1907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1901)>>(uint(int32(3))%32)))+uint32(_consts[1303]))))
	if int32(base.Ui32(v1907)>>(uint(v1901&int32(7))%32))&int32(1) == int32(0) {
		goto L465
	} else {
		goto L468
	}
L468:
	;
	v1931 = v1897
	goto L451
L469:
	;
	goto L454
L470:
	;
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1932 <= v1933 {
		v1947 = v1784
		goto L448
	} else {
		goto L471
	}
L471:
	;
	v1935 = int32(1)
	v1938 = F_slice_from_s(m, l0, v1935, int32(2232052))
	mBase = m.M
	v1939 = m.ExcPending
	if v1939 != 0 {
		goto L7
	} else {
		goto L472
	}
L472:
	;
	if int32(0) <= v1938 {
		goto L473
	} else {
		goto L474
	}
L473:
	;
	v1945 = v1935
	goto L475
L474:
	;
	v1945 = v1938 >> (uint(int32(31)) % 32) & v1938
	goto L475
L475:
	;
	v1947 = v1945
	goto L448
L476:
	;
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1950
	v1952 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1950
	v1956 = v1950 - int32(1)
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1956 <= v1957 {
		v2216 = v1952
		goto L477
	} else {
		goto L478
	}
L477:
	;
	if v2216 < int32(0) {
		v2585 = v2216
		goto L1
	} else {
		goto L557
	}
L478:
	;
	v1959 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1959+v1956))))
	if v1961&int32(224) != int32(96) {
		v2216 = v1952
		goto L477
	} else {
		goto L479
	}
L479:
	;
	if int32(1)<<(uint(v1961)%32)&int32(815616) == int32(0) {
		v2216 = v1952
		goto L477
	} else {
		goto L480
	}
L480:
	;
	v1974 = F_find_among_b(m, l0, int32(4303680), int32(24))
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L7
	} else {
		goto L481
	}
L481:
	;
	if v1974 == int32(0) {
		v2216 = v1952
		goto L477
	} else {
		goto L482
	}
L482:
	;
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1978
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(v1980)+4))
	if v1978 < v1981 {
		v2216 = v1952
		goto L477
	} else {
		goto L483
	}
L483:
	;
	switch v1974 - int32(1) {
	case 0:
		goto L499
	case 1:
		goto L498
	case 2:
		goto L497
	case 3:
		goto L496
	case 4:
		goto L495
	case 5:
		goto L494
	case 6:
		goto L493
	case 7:
		goto L492
	case 8:
		goto L491
	case 9:
		goto L490
	case 10:
		goto L489
	case 11:
		goto L488
	case 12:
		goto L487
	case 13:
		goto L486
	case 14:
		goto L485
	default:
		goto L484
	}
L484:
	;
	v2216 = int32(1)
	goto L477
L485:
	;
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L532
L486:
	;
	v2077 = F_slice_from_s(m, l0, int32(4), int32(2232094))
	mBase = m.M
	v2078 = m.ExcPending
	if v2078 != 0 {
		goto L7
	} else {
		goto L528
	}
L487:
	;
	v2057 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1978 <= v2057 {
		v2216 = v1952
		goto L477
	} else {
		goto L524
	}
L488:
	;
	v2053 = F_slice_from_s(m, l0, int32(3), int32(2232089))
	mBase = m.M
	v2054 = m.ExcPending
	if v2054 != 0 {
		goto L7
	} else {
		goto L522
	}
L489:
	;
	v2047 = F_slice_from_s(m, l0, int32(3), int32(2232086))
	mBase = m.M
	v2048 = m.ExcPending
	if v2048 != 0 {
		goto L7
	} else {
		goto L520
	}
L490:
	;
	v2041 = F_slice_from_s(m, l0, int32(3), int32(2232083))
	mBase = m.M
	v2042 = m.ExcPending
	if v2042 != 0 {
		goto L7
	} else {
		goto L518
	}
L491:
	;
	v2035 = F_slice_from_s(m, l0, int32(3), int32(2232080))
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L7
	} else {
		goto L516
	}
L492:
	;
	v2029 = F_slice_from_s(m, l0, int32(2), int32(2232078))
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		goto L7
	} else {
		goto L514
	}
L493:
	;
	v2023 = F_slice_from_s(m, l0, int32(3), int32(2232075))
	mBase = m.M
	v2024 = m.ExcPending
	if v2024 != 0 {
		goto L7
	} else {
		goto L512
	}
L494:
	;
	v2017 = F_slice_from_s(m, l0, int32(3), int32(2232072))
	mBase = m.M
	v2018 = m.ExcPending
	if v2018 != 0 {
		goto L7
	} else {
		goto L510
	}
L495:
	;
	v2011 = F_slice_from_s(m, l0, int32(3), int32(2232069))
	mBase = m.M
	v2012 = m.ExcPending
	if v2012 != 0 {
		goto L7
	} else {
		goto L508
	}
L496:
	;
	v2005 = F_slice_from_s(m, l0, int32(4), int32(2232065))
	mBase = m.M
	v2006 = m.ExcPending
	if v2006 != 0 {
		goto L7
	} else {
		goto L506
	}
L497:
	;
	v1999 = F_slice_from_s(m, l0, int32(4), int32(2232061))
	mBase = m.M
	v2000 = m.ExcPending
	if v2000 != 0 {
		goto L7
	} else {
		goto L504
	}
L498:
	;
	v1993 = F_slice_from_s(m, l0, int32(4), int32(2232057))
	mBase = m.M
	v1994 = m.ExcPending
	if v1994 != 0 {
		goto L7
	} else {
		goto L502
	}
L499:
	;
	v1987 = F_slice_from_s(m, l0, int32(4), int32(2232053))
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		goto L7
	} else {
		goto L500
	}
L500:
	;
	if int32(0) <= v1987 {
		goto L484
	} else {
		goto L501
	}
L501:
	;
	v2216 = v1987
	goto L477
L502:
	;
	if int32(0) <= v1993 {
		goto L484
	} else {
		goto L503
	}
L503:
	;
	v2216 = v1993
	goto L477
L504:
	;
	if int32(0) <= v1999 {
		goto L484
	} else {
		goto L505
	}
L505:
	;
	v2216 = v1999
	goto L477
L506:
	;
	if int32(0) <= v2005 {
		goto L484
	} else {
		goto L507
	}
L507:
	;
	v2216 = v2005
	goto L477
L508:
	;
	if int32(0) <= v2011 {
		goto L484
	} else {
		goto L509
	}
L509:
	;
	v2216 = v2011
	goto L477
L510:
	;
	if int32(0) <= v2017 {
		goto L484
	} else {
		goto L511
	}
L511:
	;
	v2216 = v2017
	goto L477
L512:
	;
	if int32(0) <= v2023 {
		goto L484
	} else {
		goto L513
	}
L513:
	;
	v2216 = v2023
	goto L477
L514:
	;
	if int32(0) <= v2029 {
		goto L484
	} else {
		goto L515
	}
L515:
	;
	v2216 = v2029
	goto L477
L516:
	;
	if int32(0) <= v2035 {
		goto L484
	} else {
		goto L517
	}
L517:
	;
	v2216 = v2035
	goto L477
L518:
	;
	if int32(0) <= v2041 {
		goto L484
	} else {
		goto L519
	}
L519:
	;
	v2216 = v2041
	goto L477
L520:
	;
	if int32(0) <= v2047 {
		goto L484
	} else {
		goto L521
	}
L521:
	;
	v2216 = v2047
	goto L477
L522:
	;
	if int32(0) <= v2053 {
		goto L484
	} else {
		goto L523
	}
L523:
	;
	v2216 = v2053
	goto L477
L524:
	;
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2059+v1978-int32(1)))))
	if v2063 != int32(108) {
		v2216 = v1952
		goto L477
	} else {
		goto L525
	}
L525:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1978 - int32(1)
	v2071 = F_slice_from_s(m, l0, int32(2), int32(2232092))
	mBase = m.M
	v2072 = m.ExcPending
	if v2072 != 0 {
		goto L7
	} else {
		goto L526
	}
L526:
	;
	if int32(0) <= v2071 {
		goto L484
	} else {
		goto L527
	}
L527:
	;
	v2216 = v2071
	goto L477
L528:
	;
	if int32(0) <= v2077 {
		goto L484
	} else {
		goto L529
	}
L529:
	;
	v2216 = v2077
	goto L477
L530:
	;
	if v2209 != 0 {
		v2216 = v1952
		goto L477
	} else {
		goto L554
	}
L531:
	;
	v2209 = v2202
	goto L530
L532:
	;
	if v2097 <= v2098 {
		v2202 = int32(-1)
		goto L531
	} else {
		goto L534
	}
L533:
	;
	v2202 = int32(0)
	goto L531
L534:
	;
	v2115 = int32(1)
	v2116 = v2097 - v2115
	v2118 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2094+v2116))))
	v2120 = v2118 & int32(255)
	if v2116 == v2098 {
		v2175 = v2120
		v2176 = v2115
		goto L535
	} else {
		goto L536
	}
L535:
	;
	if int32(116) < v2175 {
		goto L544
	} else {
		goto L545
	}
L536:
	;
	if int32(0) <= v2118 {
		v2175 = v2120
		v2176 = v2115
		goto L535
	} else {
		goto L537
	}
L537:
	;
	v2126 = v2120 & int32(63)
	v2128 = v2097 - int32(2)
	v2130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2094+v2128))))
	v2132 = v2130 << (uint(int32(6)) % 32)
	if base.B2i32(v2128 != v2098)&base.B2i32(base.Ui32(v2130) < base.Ui32(int32(192))) == int32(0) {
		goto L538
	} else {
		goto L539
	}
L538:
	;
	v2175 = v2132&int32(1984) | v2126
	v2176 = int32(2)
	goto L535
L539:
	;
	goto L540
L540:
	;
	v2145 = v2132&int32(4032) | v2126
	v2147 = v2097 - int32(3)
	v2149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2094+v2147))))
	if base.B2i32(v2147 != v2098)&base.B2i32(base.Ui32(v2149) < base.Ui32(int32(224))) == int32(0) {
		goto L541
	} else {
		goto L542
	}
L541:
	;
	v2175 = v2149<<(uint(int32(12))%32)&int32(61440) | v2145
	v2176 = int32(3)
	goto L535
L542:
	;
	goto L543
L543:
	;
	v2167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2097+(v2094-int32(4))))))
	v2175 = v2149<<(uint(int32(12))%32)&int32(258048) | v2167&int32(7)<<(uint(int32(18))%32) | v2145
	v2176 = int32(4)
	goto L535
L544:
	;
	v2209 = v2176
	goto L530
L545:
	;
	goto L546
L546:
	;
	v2180 = v2175 - int32(99)
	if v2180 < int32(0) {
		goto L547
	} else {
		goto L548
	}
L547:
	;
	v2209 = v2176
	goto L530
L548:
	;
	goto L549
L549:
	;
	v2186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2180)>>(uint(int32(3))%32)))+uint32(_consts[1305]))))
	if int32(base.Ui32(v2186)>>(uint(v2180&int32(7))%32))&int32(1) == int32(0) {
		goto L550
	} else {
		goto L551
	}
L550:
	;
	v2209 = v2176
	goto L530
L551:
	;
	goto L552
L552:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2097 - v2176
	goto L553
L553:
	;
	goto L533
L554:
	;
	v2210 = F_slice_del(m, l0)
	mBase = m.M
	v2211 = m.ExcPending
	if v2211 != 0 {
		goto L7
	} else {
		goto L555
	}
L555:
	;
	if v2210 < int32(0) {
		v2216 = v2210
		goto L477
	} else {
		goto L556
	}
L556:
	;
	goto L484
L557:
	;
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2221
	v2223 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2221
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2221-int32(2) <= v2226 {
		v2294 = v2223
		goto L558
	} else {
		goto L559
	}
L558:
	;
	if v2294 < int32(0) {
		v2585 = v2294
		goto L1
	} else {
		goto L585
	}
L559:
	;
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2230+v2221-int32(1)))))
	if v2234&int32(224) != int32(96) {
		v2294 = v2223
		goto L558
	} else {
		goto L560
	}
L560:
	;
	if int32(1)<<(uint(v2234)%32)&int32(528928) == int32(0) {
		v2294 = v2223
		goto L558
	} else {
		goto L561
	}
L561:
	;
	v2247 = F_find_among_b(m, l0, int32(4304160), int32(9))
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L7
	} else {
		goto L562
	}
L562:
	;
	if v2247 == int32(0) {
		v2294 = v2223
		goto L558
	} else {
		goto L563
	}
L563:
	;
	v2251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2251
	v2253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2254 = *(*int32)(unsafe.Add(mBase, uint32(v2253)+4))
	if v2251 < v2254 {
		v2294 = v2223
		goto L558
	} else {
		goto L564
	}
L564:
	;
	switch v2247 - int32(1) {
	case 0:
		goto L571
	case 1:
		goto L570
	case 2:
		goto L569
	case 3:
		goto L568
	case 4:
		goto L567
	case 5:
		goto L566
	default:
		goto L565
	}
L565:
	;
	v2294 = int32(1)
	goto L558
L566:
	;
	v2286 = *(*int32)(unsafe.Add(mBase, uint32(v2253)))
	if v2251 < v2286 {
		v2294 = v2223
		goto L558
	} else {
		goto L582
	}
L567:
	;
	v2282 = F_slice_del(m, l0)
	mBase = m.M
	v2283 = m.ExcPending
	if v2283 != 0 {
		goto L7
	} else {
		goto L580
	}
L568:
	;
	v2278 = F_slice_from_s(m, l0, int32(2), int32(2232230))
	mBase = m.M
	v2279 = m.ExcPending
	if v2279 != 0 {
		goto L7
	} else {
		goto L578
	}
L569:
	;
	v2272 = F_slice_from_s(m, l0, int32(2), int32(2232228))
	mBase = m.M
	v2273 = m.ExcPending
	if v2273 != 0 {
		goto L7
	} else {
		goto L576
	}
L570:
	;
	v2266 = F_slice_from_s(m, l0, int32(3), int32(2232225))
	mBase = m.M
	v2267 = m.ExcPending
	if v2267 != 0 {
		goto L7
	} else {
		goto L574
	}
L571:
	;
	v2260 = F_slice_from_s(m, l0, int32(4), int32(2232221))
	mBase = m.M
	v2261 = m.ExcPending
	if v2261 != 0 {
		goto L7
	} else {
		goto L572
	}
L572:
	;
	if int32(0) <= v2260 {
		goto L565
	} else {
		goto L573
	}
L573:
	;
	v2294 = v2260
	goto L558
L574:
	;
	if int32(0) <= v2266 {
		goto L565
	} else {
		goto L575
	}
L575:
	;
	v2294 = v2266
	goto L558
L576:
	;
	if int32(0) <= v2272 {
		goto L565
	} else {
		goto L577
	}
L577:
	;
	v2294 = v2272
	goto L558
L578:
	;
	if int32(0) <= v2278 {
		goto L565
	} else {
		goto L579
	}
L579:
	;
	v2294 = v2278
	goto L558
L580:
	;
	if int32(0) <= v2282 {
		goto L565
	} else {
		goto L581
	}
L581:
	;
	v2294 = v2282
	goto L558
L582:
	;
	v2288 = F_slice_del(m, l0)
	mBase = m.M
	v2289 = m.ExcPending
	if v2289 != 0 {
		goto L7
	} else {
		goto L583
	}
L583:
	;
	if v2288 < int32(0) {
		v2294 = v2288
		goto L558
	} else {
		goto L584
	}
L584:
	;
	goto L565
L585:
	;
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2300
	v2302 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2300
	v2306 = v2300 - int32(1)
	v2307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2306 <= v2307 {
		v2362 = v2302
		goto L586
	} else {
		goto L587
	}
L586:
	;
	if v2362 < int32(0) {
		v2585 = v2362
		goto L1
	} else {
		goto L602
	}
L587:
	;
	v2309 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2309+v2306))))
	if v2311&int32(224) != int32(96) {
		v2362 = v2302
		goto L586
	} else {
		goto L588
	}
L588:
	;
	if int32(1)<<(uint(v2311)%32)&int32(1864232) == int32(0) {
		v2362 = v2302
		goto L586
	} else {
		goto L589
	}
L589:
	;
	v2324 = F_find_among_b(m, l0, int32(4304352), int32(18))
	mBase = m.M
	v2325 = m.ExcPending
	if v2325 != 0 {
		goto L7
	} else {
		goto L590
	}
L590:
	;
	if v2324 == int32(0) {
		v2362 = v2302
		goto L586
	} else {
		goto L591
	}
L591:
	;
	v2328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2328
	v2330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2331 = *(*int32)(unsafe.Add(mBase, uint32(v2330)))
	if v2328 < v2331 {
		v2362 = v2302
		goto L586
	} else {
		goto L592
	}
L592:
	;
	switch v2324 - int32(1) {
	case 0:
		goto L595
	case 1:
		goto L594
	default:
		goto L593
	}
L593:
	;
	v2362 = int32(1)
	goto L586
L594:
	;
	v2339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2328 <= v2339 {
		v2362 = v2302
		goto L586
	} else {
		goto L598
	}
L595:
	;
	v2335 = F_slice_del(m, l0)
	mBase = m.M
	v2336 = m.ExcPending
	if v2336 != 0 {
		goto L7
	} else {
		goto L596
	}
L596:
	;
	if int32(0) <= v2335 {
		goto L593
	} else {
		goto L597
	}
L597:
	;
	v2362 = v2335
	goto L586
L598:
	;
	v2341 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2343 = int32(1)
	v2345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2341+v2328-v2343))))
	if base.Ui32(v2343) < base.Ui32((v2345-int32(115))&int32(255)) {
		v2362 = v2302
		goto L586
	} else {
		goto L599
	}
L599:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2328 - int32(1)
	v2355 = F_slice_del(m, l0)
	mBase = m.M
	v2356 = m.ExcPending
	if v2356 != 0 {
		goto L7
	} else {
		goto L600
	}
L600:
	;
	if v2355 < int32(0) {
		v2362 = v2355
		goto L586
	} else {
		goto L601
	}
L601:
	;
	goto L593
L602:
	;
	v2366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2366
	v2368 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2366
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2366 <= v2371 {
		v2470 = v2368
		goto L603
	} else {
		goto L604
	}
L603:
	;
	if v2470 < int32(0) {
		v2585 = v2470
		goto L1
	} else {
		goto L631
	}
L604:
	;
	v2373 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2373+v2366-int32(1)))))
	switch v2377 - int32(101) {
	case 0, 7:
		goto L605
	default:
		v2470 = v2368
		goto L603
	}
L605:
	;
	v2382 = F_find_among_b(m, l0, int32(4304720), int32(2))
	mBase = m.M
	v2383 = m.ExcPending
	if v2383 != 0 {
		goto L7
	} else {
		goto L606
	}
L606:
	;
	if v2382 == int32(0) {
		v2470 = v2368
		goto L603
	} else {
		goto L607
	}
L607:
	;
	v2386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2386
	switch v2382 - int32(1) {
	case 0:
		goto L610
	case 1:
		goto L609
	default:
		goto L608
	}
L608:
	;
	v2470 = int32(1)
	goto L603
L609:
	;
	v2446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2447 = *(*int32)(unsafe.Add(mBase, uint32(v2446)))
	if v2386 < v2447 {
		v2470 = v2368
		goto L603
	} else {
		goto L626
	}
L610:
	;
	v2390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2391 = *(*int32)(unsafe.Add(mBase, uint32(v2390)))
	if v2386 < v2391 {
		goto L611
	} else {
		goto L612
	}
L611:
	;
	v2393 = *(*int32)(unsafe.Add(mBase, uint32(v2390)+4))
	if v2386 < v2393 {
		v2470 = v2368
		goto L603
	} else {
		goto L614
	}
L612:
	;
	goto L613
L613:
	;
	v2442 = F_slice_del(m, l0)
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		goto L7
	} else {
		goto L624
	}
L614:
	;
	v2395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2396 = int32(0)
	v2400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2405 = F_out_grouping_b_U(m, l0, int32(2232047), int32(89), int32(121), v2396)
	mBase = m.M
	if v2405 != 0 {
		goto L616
	} else {
		goto L617
	}
L615:
	;
	if v2436 != 0 {
		v2470 = v2368
		goto L603
	} else {
		goto L623
	}
L616:
	;
	v2417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2417 + (v2400 - v2395)
	v2425 = F_out_grouping_b_U(m, l0, int32(2231895), int32(97), int32(121), int32(0))
	mBase = m.M
	if v2425 != 0 {
		v2434 = v2396
		goto L620
	} else {
		goto L621
	}
L617:
	;
	v2410 = F_in_grouping_b_U(m, l0, int32(2231895), int32(97), int32(121), int32(0))
	mBase = m.M
	if v2410 != 0 {
		goto L616
	} else {
		goto L618
	}
L618:
	;
	v2415 = F_out_grouping_b_U(m, l0, int32(2231895), int32(97), int32(121), int32(0))
	mBase = m.M
	if v2415 != 0 {
		goto L616
	} else {
		goto L619
	}
L619:
	;
	v2436 = int32(1)
	goto L615
L620:
	;
	v2436 = v2434
	goto L615
L621:
	;
	v2430 = F_in_grouping_b_U(m, l0, int32(2231895), int32(97), int32(121), int32(0))
	mBase = m.M
	if v2430 != 0 {
		v2434 = v2396
		goto L620
	} else {
		goto L622
	}
L622:
	;
	v2431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2434 = base.B2i32(v2431 <= v2432)
	goto L620
L623:
	;
	v2437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2437 + (v2386 - v2395)
	goto L613
L624:
	;
	if int32(0) <= v2442 {
		goto L608
	} else {
		goto L625
	}
L625:
	;
	v2470 = v2442
	goto L603
L626:
	;
	v2449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2386 <= v2449 {
		v2470 = v2368
		goto L603
	} else {
		goto L627
	}
L627:
	;
	v2451 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2451+v2386-int32(1)))))
	if v2455 != int32(108) {
		v2470 = v2368
		goto L603
	} else {
		goto L628
	}
L628:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2386 - int32(1)
	v2461 = F_slice_del(m, l0)
	mBase = m.M
	v2462 = m.ExcPending
	if v2462 != 0 {
		goto L7
	} else {
		goto L629
	}
L629:
	;
	if v2461 < int32(0) {
		v2470 = v2461
		goto L603
	} else {
		goto L630
	}
L630:
	;
	goto L608
L631:
	;
	v2473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2476 = v2473
	goto L337
L632:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2476
	goto L2
L633:
	;
	goto L634
L634:
	;
	v2490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2492 = v2490
	goto L636
L635:
	;
	v2585 = v2566
	goto L1
L636:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2492
	v2498 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2499 != v2492 {
		goto L639
	} else {
		goto L640
	}
L637:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2492
	v2561 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2492 + v2561
	v2566 = F_slice_from_s(m, l0, v2561, int32(2232336))
	mBase = m.M
	v2567 = m.ExcPending
	if v2567 != 0 {
		goto L7
	} else {
		goto L664
	}
L638:
	;
	goto L637
L639:
	;
	v2502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2492+v2498))))
	if v2502 == int32(89) {
		goto L638
	} else {
		goto L642
	}
L640:
	;
	goto L641
L641:
	;
	goto L645
L642:
	;
	goto L641
L643:
	;
	if v2556 < int32(0) {
		goto L632
	} else {
		goto L663
	}
L645:
	;
	goto L646
L646:
	;
	goto L647
L647:
	;
	v2511 = v2492
	v2513 = int32(1)
	goto L650
L649:
	;
	v2556 = v2541
	goto L643
L650:
	;
	if v2499 <= v2511 {
		goto L652
	} else {
		goto L653
	}
L651:
	;
	goto L649
L652:
	;
	v2556 = int32(-1)
	goto L643
L653:
	;
	goto L654
L654:
	;
	v2518 = v2511 + int32(1)
	v2520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2498+v2511))))
	if base.Ui32(v2520) < base.Ui32(int32(192)) {
		v2541 = v2518
		goto L655
	} else {
		goto L656
	}
L655:
	;
	v2542 = int32(1)
	if v2542 < v2513 {
		v2511 = v2541
		v2513 = v2513 - v2542
		goto L650
	} else {
		goto L662
	}
L656:
	;
	if v2499 <= v2518 {
		v2541 = v2518
		goto L655
	} else {
		goto L657
	}
L657:
	;
	v2527 = v2518
	goto L658
L658:
	;
	v2530 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2498+v2527))))
	if int32(-65) < v2530 {
		v2541 = v2527
		goto L655
	} else {
		goto L660
	}
L659:
	;
	v2541 = v2499
	goto L655
L660:
	;
	v2534 = v2527 + int32(1)
	if v2534 != v2499 {
		v2527 = v2534
		goto L658
	} else {
		goto L661
	}
L661:
	;
	goto L659
L662:
	;
	goto L651
L663:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2556
	v2492 = v2556
	goto L636
L664:
	;
	if int32(0) <= v2566 {
		goto L634
	} else {
		goto L665
	}
L665:
	;
	goto L635
}
func F_eq_s_b(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6-v7 < l1 {
		v78 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v78
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = v10 + v6 - l1
	if base.Ui32(int32(4)) <= base.Ui32(l1) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	if v74 != 0 {
		v78 = v4
		goto L1
	} else {
		goto L21
	}
L4:
	;
	v74 = int32(0)
	goto L3
L5:
	;
	v48 = v43
	v49 = v44
	v50 = v45
	goto L15
L6:
	;
	if (v12|l2)&int32(3) != 0 {
		v43 = v12
		v44 = l2
		v45 = l1
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v36 = v12
	v37 = l2
	v38 = l1
	goto L8
L8:
	;
	if v38 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L9:
	;
	v20 = v12
	v21 = l2
	v22 = l1
	goto L10
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v25 != v26 {
		v43 = v20
		v44 = v21
		v45 = v22
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v36 = v31
	v37 = v29
	v38 = v33
	goto L8
L12:
	;
	v28 = int32(4)
	v29 = v21 + v28
	v31 = v20 + v28
	v33 = v22 - v28
	if base.Ui32(int32(3)) < base.Ui32(v33) {
		v20 = v31
		v21 = v29
		v22 = v33
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v43 = v36
	v44 = v37
	v45 = v38
	goto L5
L15:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v53 == v54 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v74 = v53 - v54
	goto L3
L17:
	;
	v56 = int32(1)
	v61 = v50 - v56
	if v61 != 0 {
		v48 = v48 + v56
		v49 = v49 + v56
		v50 = v61
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L4
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 - l1
	v78 = int32(1)
	goto L1
}
func F_eq_v(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1-int32(4))))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9-v10 < v8 {
		v80 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v80
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = v13 + v10
	if base.Ui32(int32(4)) <= base.Ui32(v8) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	if v76 != 0 {
		v80 = v3
		goto L1
	} else {
		goto L21
	}
L4:
	;
	v76 = int32(0)
	goto L3
L5:
	;
	v50 = v45
	v51 = v46
	v52 = v47
	goto L15
L6:
	;
	if (v14|l1)&int32(3) != 0 {
		v45 = v14
		v46 = l1
		v47 = v8
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v38 = v14
	v39 = l1
	v40 = v8
	goto L8
L8:
	;
	if v40 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L9:
	;
	v22 = v14
	v23 = l1
	v24 = v8
	goto L10
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v27 != v28 {
		v45 = v22
		v46 = v23
		v47 = v24
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v38 = v33
	v39 = v31
	v40 = v35
	goto L8
L12:
	;
	v30 = int32(4)
	v31 = v23 + v30
	v33 = v22 + v30
	v35 = v24 - v30
	if base.Ui32(int32(3)) < base.Ui32(v35) {
		v22 = v33
		v23 = v31
		v24 = v35
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v45 = v38
	v46 = v39
	v47 = v40
	goto L5
L15:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 == v56 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v76 = v55 - v56
	goto L3
L17:
	;
	v58 = int32(1)
	v63 = v52 - v58
	if v63 != 0 {
		v50 = v50 + v58
		v51 = v51 + v58
		v52 = v63
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L4
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8 + v10
	v80 = int32(1)
	goto L1
}
func F_eqjoinsel_semi(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64, l4 float64, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) float64 {
	mBase := m.M
	_ = mBase
	var v14 float64
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 float64
	_ = v29
	var v34 float64
	_ = v34
	var v35 int32
	_ = v35
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
	var v39 float64
	_ = v39
	var v42 int32
	_ = v42
	var v49 float32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v66 float64
	_ = v66
	var v68 float64
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
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
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v124 int32
	_ = v124
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
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
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 float32
	_ = v191
	var v194 float64
	_ = v194
	var v214 float64
	_ = v214
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v232 float64
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v257 float32
	_ = v257
	var v260 float64
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 float32
	_ = v271
	var v274 float64
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v284 float64
	_ = v284
	var v289 int32
	_ = v289
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v311 float32
	_ = v311
	var v314 float64
	_ = v314
	var v317 float64
	_ = v317
	var v334 float64
	_ = v334
	var v335 float64
	_ = v335
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v350 float64
	_ = v350
	var v351 float64
	_ = v351
	var v352 float64
	_ = v352
	var v358 float64
	_ = v358
	var v360 float64
	_ = v360
	var v363 float64
	_ = v363
	var v371 float64
	_ = v371
	var v395 float64
	_ = v395
	v14 = float64(0)
	v18 = int32(0)
	v22 = m.G0
	v24 = v22 + int32(-64)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v26 == v18 {
		v34 = l4
		v35 = l6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v37 = *(*float64)(unsafe.Add(mBase, uint32(l12)+16))
	v38 = base.F64_ge(v34, v37)
	if v38 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v29 = *(*float64)(unsafe.Add(mBase, uint32(v26)+16))
	if base.F64_ge(l4, v29) == int32(0) {
		v34 = l4
		v35 = l6
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v34 = v29
	v35 = int32(0)
	goto L1
L4:
	;
	v39 = v37
	goto L6
L5:
	;
	v39 = v34
	goto L6
L6:
	;
	v42 = v35 & (v38 ^ int32(1))
	if l10 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	m.G0 = v24 - int32(-64)
	return v395
L8:
	;
	F_pfree(m, v77)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L14
	} else {
		goto L67
	}
L9:
	;
	v214 = base.F64_convert_i32_s(v184)
	if v189 <= int32(0) {
		v334 = v14
		v335 = v214
		goto L8
	} else {
		goto L48
	}
L10:
	;
	if l9 != 0 {
		goto L41
	} else {
		goto L42
	}
L11:
	;
	if l0 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	if l11 == int32(0) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v49 = *(*float32)(unsafe.Add(mBase, uint32(l9)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l8)+16))
	F_fmgr_info(m, l0, v24)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return float64(0)
L15:
	;
	v55 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+60)) = uint8(v55)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+52)) = uint8(v55)
	v59 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+46)) = uint16(v59)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+44)) = uint8(v55)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v24)+32)) = int64(0)
	v66 = base.F64_convert_i32_s(v50)
	if base.F64_lt(v66, v39) != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v24
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l7)+16))
	v77 = F_palloc0(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L14
	} else {
		goto L23
	}
L17:
	;
	v68 = v66
	goto L19
L18:
	;
	v68 = v39
	goto L19
L19:
	;
	if base.F64_lt(base.F64_abs(v68), float64(2.147483648e+09)) != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v72 = base.I32_trunc_f64_s(v68)
	v74 = v72
	goto L16
L21:
	;
	goto L22
L22:
	;
	v74 = int32(-2147483648)
	goto L16
L23:
	;
	v79 = F_palloc0(m, v74)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L14
	} else {
		goto L24
	}
L24:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l7)+16))
	if v81 <= int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v334 = v14
	v335 = v14
	goto L8
L26:
	;
	goto L27
L27:
	;
	v84 = int32(0)
	v97 = v84
	v105 = v18
	goto L28
L28:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v97<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v112
	if v74 <= v84 {
		v184 = v105
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L9
L30:
	;
	v188 = v97 + int32(1)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l7)+16))
	if v188 < v189 {
		v97 = v188
		v105 = v184
		goto L28
	} else {
		goto L40
	}
L31:
	;
	v124 = int32(0)
	goto L32
L32:
	;
	v136 = v79 + v124
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	if v137 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v184 = v105
	goto L30
L34:
	;
	v164 = v124 + int32(1)
	if v164 != v74 {
		v124 = v164
		goto L32
	} else {
		goto L39
	}
L35:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l8)+12))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v138+v124<<(uint(int32(2))%32))))
	v143 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+44)) = uint8(v143)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+56)) = v142
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v150 = m.T0[v149].(func(*base.Module, int32) int32)(m, v22+int32(-36))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L14
	} else {
		goto L36
	}
L36:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+44)))
	if v152 != 0 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	if v150 == int32(0) {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v155 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v136))) = uint8(v155)
	*(*uint8)(unsafe.Add(mBase, uint32(v77+v97))) = uint8(v155)
	v184 = v105 + v155
	goto L30
L39:
	;
	goto L33
L40:
	;
	goto L29
L41:
	;
	v191 = *(*float32)(unsafe.Add(mBase, uint32(l9)+8))
	v194 = base.F64_promote_f32(v191)
	goto L43
L42:
	;
	v194 = float64(0)
	goto L43
L43:
	;
	if (l5|v42)&int32(1) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if base.F64_lt(v39, float64(0))|base.F64_le(l3, v39) != 0 {
		v395 = base.F64_sub(float64(1), v194)
		goto L7
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v395 = base.F64_mul(base.F64_sub(float64(1), v194), float64(0.5))
	goto L7
L47:
	;
	v395 = base.F64_mul(base.F64_div(v39, l3), base.F64_sub(float64(1), v194))
	goto L7
L48:
	;
	v217 = int32(1)
	if v189 == v217 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v189&v217 == int32(0) {
		v314 = v284
		goto L62
	} else {
		goto L63
	}
L50:
	;
	v284 = float64(0)
	v289 = int32(0)
	goto L49
L51:
	;
	goto L52
L52:
	;
	v226 = int32(0)
	v232 = float64(0)
	v234 = v226
	v237 = v226
	goto L53
L53:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v237))))
	if v250 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v284 = v274
	v289 = v276
	goto L49
L55:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l7)+20))
	v257 = *(*float32)(unsafe.Add(mBase, uint32(v253+v237<<(uint(int32(2))%32))))
	v260 = base.F64_add(v232, base.F64_promote_f32(v257))
	goto L57
L56:
	;
	v260 = v232
	goto L57
L57:
	;
	v261 = int32(1)
	v262 = v237 | v261
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v262))))
	if v264 == v261 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l7)+20))
	v271 = *(*float32)(unsafe.Add(mBase, uint32(v267+v262<<(uint(int32(2))%32))))
	v274 = base.F64_add(v260, base.F64_promote_f32(v271))
	goto L60
L59:
	;
	v274 = v260
	goto L60
L60:
	;
	v275 = int32(2)
	v276 = v237 + v275
	v278 = v234 + v275
	if v278 != v189&int32(2147483646) {
		v232 = v274
		v234 = v278
		v237 = v276
		goto L53
	} else {
		goto L61
	}
L61:
	;
	goto L54
L62:
	;
	if base.F64_lt(v314, float64(0)) != 0 {
		v334 = v14
		v335 = v214
		goto L8
	} else {
		goto L65
	}
L63:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v289))))
	if v304 != int32(1) {
		v314 = v284
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l7)+20))
	v311 = *(*float32)(unsafe.Add(mBase, uint32(v307+v289<<(uint(int32(2))%32))))
	v314 = base.F64_add(v284, base.F64_promote_f32(v311))
	goto L62
L65:
	;
	v317 = float64(1)
	if base.F64_gt(v314, v317) != 0 {
		v334 = v317
		v335 = v214
		goto L8
	} else {
		goto L66
	}
L66:
	;
	v334 = v314
	v335 = v214
	goto L8
L67:
	;
	F_pfree(m, v79)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L14
	} else {
		goto L68
	}
L68:
	;
	if (l5|v42)&int32(1) != 0 {
		v358 = float64(0.5)
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v360 = float64(0)
	v363 = base.F64_sub(base.F64_sub(float64(1), v334), base.F64_promote_f32(v49))
	if base.F64_lt(v363, v360) != 0 {
		v371 = v360
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v350 = float64(1)
	v351 = base.F64_sub(l3, v335)
	v352 = base.F64_sub(v39, v335)
	if base.F64_le(v351, v352) != 0 {
		v358 = v350
		goto L69
	} else {
		goto L71
	}
L71:
	;
	if base.F64_lt(v352, float64(0)) != 0 {
		v358 = v350
		goto L69
	} else {
		goto L72
	}
L72:
	;
	v358 = base.F64_div(v352, v351)
	goto L69
L73:
	;
	v395 = base.F64_add(base.F64_mul(v358, v371), v334)
	goto L7
L74:
	;
	if base.F64_gt(v363, float64(1)) == int32(0) {
		v371 = v363
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v371 = float64(1)
	goto L73
}
func F_errhidestmt(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, _consts[1177]))
	if v3 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[1177])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(476389), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errfinish(m, int32(523980), int32(1438), int32(103786))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v26 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v3*int32(100))+uint32(_consts[1187]))) = uint8(v26)
		return
	}
}
func F_errhint_plural(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(4555020)
	v14 = *(*int32)(unsafe.Add(mBase, _consts[1176]))
	*(*int32)(unsafe.Add(mBase, _consts[1176])) = v14 + int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, _consts[1177]))
	if int32(0) <= v19 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = int32(4562096)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v26 = v19 * int32(100)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1182])))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v29
	F_initStringInfo(m, v10+int32(16))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1177])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L24
	}
L4:
	;
	return
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1181])))
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l3
	if l2 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v43 = l0
	goto L8
L7:
	;
	v43 = l1
	goto L8
L8:
	;
	v44 = F_appendStringInfoVA(m, v10+int32(16), v43, l3)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v44 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v48 = v44
	goto L13
L11:
	;
	goto L12
L12:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1186])))
	if v72 != 0 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	F_enlargeStringInfo(m, v10+int32(16), v48)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1181])))
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l3
	v63 = F_appendStringInfoVA(m, v10+int32(16), v43, l3)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if v63 != 0 {
		v48 = v63
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	F_pfree(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v76 = F_pstrdup(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1186]))) = v76
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	F_pfree(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v23
	v84 = int32(4555020)
	v86 = *(*int32)(unsafe.Add(mBase, _consts[1176]))
	*(*int32)(unsafe.Add(mBase, _consts[1176])) = v86 - int32(1)
	m.G0 = v10 + int32(32)
	return
L24:
	;
	F_errmsg_internal(m, int32(476389), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(523980), int32(1368), int32(325723))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_errmsg_plural(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(4555020)
	v14 = *(*int32)(unsafe.Add(mBase, _consts[1176]))
	*(*int32)(unsafe.Add(mBase, _consts[1176])) = v14 + int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, _consts[1177]))
	if int32(0) <= v19 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = int32(4562096)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v26 = v19 * int32(100)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1182])))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1183]))) = l0
	F_initStringInfo(m, v10+int32(16))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1177])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L24
	}
L4:
	;
	return
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1181])))
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l3
	if l2 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v44 = l0
	goto L8
L7:
	;
	v44 = l1
	goto L8
L8:
	;
	v45 = F_appendStringInfoVA(m, v10+int32(16), v44, l3)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v47 = v45
	goto L13
L11:
	;
	goto L12
L12:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1184])))
	if v73 != 0 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	F_enlargeStringInfo(m, v10+int32(16), v47)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1181])))
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l3
	v64 = F_appendStringInfoVA(m, v10+int32(16), v44, l3)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if v64 != 0 {
		v47 = v64
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	F_pfree(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v77 = F_pstrdup(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1184]))) = v77
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	F_pfree(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v23
	v85 = int32(4555020)
	v87 = *(*int32)(unsafe.Add(mBase, _consts[1176]))
	*(*int32)(unsafe.Add(mBase, _consts[1176])) = v87 - int32(1)
	m.G0 = v10 + int32(32)
	return
L24:
	;
	F_errmsg_internal(m, int32(476389), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(523980), int32(1188), int32(325755))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_error_view_not_updatable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	v9 = m.G0
	v11 = v9 - int32(224)
	m.G0 = v11
	if l1 != int32(5) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L25
	} else {
		goto L86
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L25
	} else {
		goto L77
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L25
	} else {
		goto L68
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L25
	} else {
		goto L65
	}
L5:
	;
	switch l1 - int32(2) {
	case 0:
		goto L2
	case 1:
		goto L1
	case 2:
		goto L3
	default:
		goto L4
	}
L6:
	;
	goto L7
L7:
	;
	if l2 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	m.G0 = v11 + int32(224)
	return
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v19 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v23 = int32(0)
	if v23 < v19 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v27 = v19
	goto L13
L12:
	;
	v27 = v23
	goto L13
L13:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v30 = v23
	goto L14
L14:
	;
	v37 = int32(2)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v28+v30<<(uint(v37)%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	switch v41 - v37 {
	case 0:
		goto L19
	case 1:
		goto L20
	case 2:
		goto L18
	default:
		goto L17
	case 5:
		goto L16
	}
L15:
	;
	goto L8
L16:
	;
	v157 = v30 + int32(1)
	if v157 != v27 {
		v30 = v157
		goto L14
	} else {
		goto L64
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L25
	} else {
		goto L61
	}
L18:
	;
	if v22 != 0 {
		goto L48
	} else {
		goto L49
	}
L19:
	;
	if v22 != 0 {
		goto L35
	} else {
		goto L36
	}
L20:
	;
	if v22 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
	if v44 != 0 {
		goto L16
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	return
L26:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+144)) = v52 + int32(4)
	F_errmsg(m, int32(732614), v11+int32(144))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	if l3 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = l3
	F_errdetail_internal(m, int32(217224), v11+int32(128))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L25
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	F_errhint(m, int32(640361), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L25
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	F_errfinish(m, int32(520718), int32(3171), int32(411675))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L25
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+15)))
	if v76 != 0 {
		goto L16
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L25
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+176)) = v84 + int32(4)
	F_errmsg(m, int32(732908), v11+int32(176))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L25
	} else {
		goto L41
	}
L41:
	;
	if l3 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = l3
	F_errdetail_internal(m, int32(217224), v11+int32(160))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L25
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_errhint(m, int32(640530), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L25
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	F_errfinish(m, int32(520718), int32(3180), int32(411675))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L25
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
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+20)))
	if v108 != 0 {
		goto L16
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L25
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L25
	} else {
		goto L53
	}
L53:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+208)) = v116 + int32(4)
	F_errmsg(m, int32(732753), v11+int32(208))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L25
	} else {
		goto L54
	}
L54:
	;
	if l3 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = l3
	F_errdetail_internal(m, int32(217224), v11+int32(192))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L25
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	F_errhint(m, int32(640446), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L25
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	F_errfinish(m, int32(520718), int32(3189), int32(411675))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L25
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v144
	F_errmsg_internal(m, int32(510175), v11+int32(112))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L25
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(520718), int32(3194), int32(411675))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L25
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	goto L15
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
	F_errmsg_internal(m, int32(510204), v11)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L25
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(520718), int32(3200), int32(411675))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L25
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L25
	} else {
		goto L69
	}
L69:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v190 + int32(4)
	F_errmsg(m, int32(732753), v11+int32(96))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L25
	} else {
		goto L70
	}
L70:
	;
	if l3 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = l3
	F_errdetail_internal(m, int32(217224), v11+int32(80))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L25
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	F_errhint(m, int32(666823), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L25
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	F_errfinish(m, int32(520718), int32(3152), int32(411675))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L25
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L25
	} else {
		goto L78
	}
L78:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v221 + int32(4)
	F_errmsg(m, int32(732908), v11-int32(-64))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L25
	} else {
		goto L79
	}
L79:
	;
	if l3 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = l3
	F_errdetail_internal(m, int32(217224), v11+int32(48))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L25
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	F_errhint(m, int32(666941), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L25
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	F_errfinish(m, int32(520718), int32(3144), int32(411675))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L25
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L25
	} else {
		goto L87
	}
L87:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v252 + int32(4)
	F_errmsg(m, int32(732614), v11+int32(32))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L25
	} else {
		goto L88
	}
L88:
	;
	if l3 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l3
	F_errdetail_internal(m, int32(217224), v11+int32(16))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L25
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	F_errhint(m, int32(666704), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L25
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	F_errfinish(m, int32(520718), int32(3136), int32(411675))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L25
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_errsave_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	if l0 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v5 == int32(447) {
			v15 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v15)
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
			if v17 == v15 {
				v20 = int32(4555020)
				v22 = *(*int32)(unsafe.Add(mBase, _consts[1176]))
				v23 = int32(1)
				*(*int32)(unsafe.Add(mBase, _consts[1176])) = v22 + v23
				v26 = int32(4164900)
				v28 = *(*int32)(unsafe.Add(mBase, _consts[1177]))
				v30 = v28 + v23
				*(*int32)(unsafe.Add(mBase, _consts[1177])) = v30
				if int32(5) <= v30 {
					*(*int32)(unsafe.Add(mBase, _consts[1177])) = int32(-1)
					F_errstart_cold(m, int32(23), int32(0))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(485084), int32(0))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(523980), int32(762), int32(12326))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v34 = int32(100)
					v35 = v30 * v34
					v41 = F__emscripten_memset_bulkmem(m, v35+int32(4555024), base.I32_extend8_s(int32(0)), v34)
					mBase = m.M
					v43 = *(*int32)(unsafe.Add(mBase, _consts[137]))
					*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[1178]))) = int32(2600)
					v50 = int32(584269)
					*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[1179]))) = v50
					*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[1180]))) = v50
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(15)
					*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[1181]))) = v43
					v64 = *(*int32)(unsafe.Add(mBase, _consts[3]))
					*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_consts[1182]))) = v64
					*(*int32)(unsafe.Add(mBase, _consts[1176])) = v22
					v74 = int32(1)
					return v74
				}
			} else {
				v74 = int32(0)
				return v74
			}
		} else {
			v10 = F_errstart(m, int32(21), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	} else {
		v10 = F_errstart(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_escape_xml(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_initStringInfo(m, v7)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = l0
	goto L3
L3:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	switch v17 - int32(38) {
	case 0:
		goto L10
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 23:
		goto L6
	case 22:
		goto L9
	case 24:
		goto L8
	default:
		goto L11
	}
L4:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	m.G0 = v7 + int32(16)
	return v63
L5:
	;
	goto L4
L6:
	;
	v40 = base.I32_extend8_s(v17)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v41 <= v42+int32(1) {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	F_appendStringInfoString(m, v7, int32(574399))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L15
	}
L8:
	;
	F_appendStringInfoString(m, v7, int32(574058))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L14
	}
L9:
	;
	F_appendStringInfoString(m, v7, int32(574053))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L13
	}
L10:
	;
	F_appendStringInfoString(m, v7, int32(574063))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	switch v17 {
	case 0:
		goto L5
	default:
		goto L6
	case 13:
		goto L7
	}
L12:
	;
	v13 = v13 + int32(1)
	goto L3
L13:
	;
	v13 = v13 + int32(1)
	goto L3
L14:
	;
	v13 = v13 + int32(1)
	goto L3
L15:
	;
	v13 = v13 + int32(1)
	goto L3
L16:
	;
	F_appendStringInfoChar(m, v7, v40)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v50+v42))) = uint8(v40)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v54 = int32(1)
	v55 = v53 + v54
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v59 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v57+v55))) = uint8(v59)
	v13 = v13 + v54
	goto L3
L19:
	;
	v13 = v13 + int32(1)
	goto L3
}
func F_examine_indexcol_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v7+l2<<(uint(int32(2))%32))))
	if v11 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if v12 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
			v28 = v13
			v29 = v12 + v14<<(uint(int32(2))%32)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
			v28 = v21
			v29 = v20 + v22<<(uint(int32(2))%32) - int32(4)
		}
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
		*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v28
		v34 = *(*int32)(unsafe.Add(mBase, _consts[1141]))
		if v34 == int32(0) {
			v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+20)))
			v62 = F_SearchSysCache3(m, int32(65), v31, base.I32_extend16_s(v11), v61)
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return
			} else {
				v101 = v62
				*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(1509)
				*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v101
				return
			}
		} else {
			v38 = m.T0[v34].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v30, base.I32_extend16_s(v11), l3)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				if v38 == int32(0) {
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+20)))
					v62 = F_SearchSysCache3(m, int32(65), v31, base.I32_extend16_s(v11), v61)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						v101 = v62
						*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(1509)
						*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v101
						return
					}
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
					if v42 == int32(0) {
						return
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						if v45 != 0 {
							return
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(337074), int32(0))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									F_errfinish(m, int32(519996), int32(6181), int32(415924))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
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
	} else {
		v66 = base.I32_extend16_s(l2 + int32(1))
		v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v69 = *(*int32)(unsafe.Add(mBase, _consts[1142]))
		if v69 == int32(0) {
			v95 = F_SearchSysCache3(m, int32(65), v67, v66, int32(0))
			mBase = m.M
			v96 = m.ExcPending
			if v96 != 0 {
				return
			} else {
				v101 = v95
				*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(1509)
				*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v101
				return
			}
		} else {
			v72 = m.T0[v69].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v67, v66, l3)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return
			} else {
				if v72 == int32(0) {
					v95 = F_SearchSysCache3(m, int32(65), v67, v66, int32(0))
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return
					} else {
						v101 = v95
						*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(1509)
						*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v101
						return
					}
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
					if v76 == int32(0) {
						return
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						if v79 != 0 {
							return
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(337074), int32(0))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									F_errfinish(m, int32(519996), int32(6207), int32(415924))
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
						}
					}
				}
			}
		}
	}
}
func F_execTuplesHashPrepare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v16 = F_palloc(m, l0<<(uint(int32(2))%32))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v16
	v21 = F_palloc(m, l0*int32(28))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v21
	if int32(0) < l0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L15
	}
L5:
	;
	v31 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	m.G0 = v12 + int32(16)
	return
L8:
	;
	v36 = v31 << (uint(int32(2)) % 32)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1+v36)))
	v39 = F_get_opcode(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	v45 = F_get_op_hash_functions(m, v38, v12+int32(12), v12+int32(8))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v45 == int32(0) {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v49+v36))) = v39
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_fmgr_info(m, v52, v53+v31*int32(28))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v60 = v31 + int32(1)
	if v60 != l0 {
		v31 = v60
		goto L8
	} else {
		goto L14
	}
L14:
	;
	goto L9
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v38
	F_errmsg_internal(m, int32(46942), v12)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(524043), int32(118), int32(384466))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_exec_move_row_from_fields(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
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
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v323 int32
	_ = v323
	var v335 int32
	_ = v335
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v615 int32
	_ = v615
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v657 int32
	_ = v657
	var v664 int32
	_ = v664
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v695 int32
	_ = v695
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v727 int32
	_ = v727
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v824 int32
	_ = v824
	var v843 int32
	_ = v843
	var v855 int32
	_ = v855
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v907 int32
	_ = v907
	v7 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(400)
	m.G0 = v26
	if l5 == v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v58 - int32(1) {
	case 0:
		goto L13
	case 1:
		goto L12
	default:
		goto L11
	}
L2:
	;
	v54 = v7
	v55 = int32(166581)
	v56 = v7
	v57 = int32(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1287])))
	if v34&int32(8) != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v54 = v32
	v55 = int32(140433)
	v56 = int32(21)
	v57 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[1288]))
	v54 = v32
	v55 = int32(166581)
	v56 = v42 << (uint(int32(28)) % 32) >> (uint(int32(31)) % 32) & int32(19)
	v57 = base.B2i32(v42&int32(8) == int32(0))
	goto L1
L8:
	;
	m.G0 = v26 + int32(400)
	return
L9:
	;
	if v57|base.B2i32(v54 <= v824) != 0 {
		goto L8
	} else {
		goto L161
	}
L10:
	;
	v657 = int32(0)
	v664 = v7
	goto L140
L11:
	;
	F_errstart_cold(m, int32(21), int32(584281))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L18
	} else {
		goto L137
	}
L12:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	if v65 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if int32(0) < v61 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v824 = int32(0)
	goto L9
L15:
	;
	v68 = F_expanded_record_fetch_tupdesc(m, l2)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v70 = v65
	goto L17
L17:
	;
	if l5 == v70 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	return
L19:
	;
	v70 = v68
	goto L17
L20:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	if v416&int32(4) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L21:
	;
	v405 = l3
	v408 = l4
	goto L20
L22:
	;
	goto L23
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if base.Ui32(v72) < base.Ui32(int32(65)) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v72 <= int32(0) {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v88 = v26 + int32(144)
	v89 = v26 + int32(80)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	v83 = F_MemoryContextAlloc(m, v80, v72*int32(5))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	v88 = v83
	v89 = v83 + v72<<(uint(int32(2))%32)
	goto L24
L29:
	;
	if v57|base.B2i32(v54 <= v304) != 0 {
		v405 = v88
		v408 = v89
		goto L20
	} else {
		goto L57
	}
L30:
	;
	v304 = int32(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v93 = int32(20)
	v104 = int32(0)
	v111 = v7
	goto L33
L33:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v127 = v70 + v93 + v121<<(uint(int32(4))%32) + v111*int32(100)
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+91)))
	if v128 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v304 = v278
	goto L29
L35:
	;
	if v104 < v54 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v278 = v104
	goto L37
L37:
	;
	v296 = v111 + int32(1)
	if v296 != v72 {
		v104 = v278
		v111 = v296
		goto L33
	} else {
		goto L56
	}
L38:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v127)+68))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v127)+76))
	v266 = F_exec_cast_value(m, l0, v246, v26+int32(79), v244, v255, v264, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L18
	} else {
		goto L55
	}
L39:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l3+v142<<(uint(int32(2))%32))))
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v142))))
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+79)) = uint8(v230)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v161)+76))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v161)+68))
	v242 = v142 + int32(1)
	v244 = v235
	v246 = v228
	v255 = v234
	goto L38
L40:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v142 = v104
	goto L43
L41:
	;
	v174 = v104
	goto L42
L42:
	;
	v191 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+79)) = uint8(v191)
	v193 = int32(-1)
	v194 = int32(705)
	v195 = int32(0)
	if v57 != 0 {
		v242 = v174
		v244 = v194
		v246 = v195
		v255 = v193
		goto L38
	} else {
		goto L47
	}
L43:
	;
	v161 = l5 + v93 + v132<<(uint(int32(4))%32) + v142*int32(100)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+91)))
	if v162 != int32(1) {
		goto L39
	} else {
		goto L45
	}
L44:
	;
	v174 = v54
	goto L42
L45:
	;
	v166 = v142 + int32(1)
	if v166 != v54 {
		v142 = v166
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v197 = F_errstart(m, v56, int32(584281))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L18
	} else {
		goto L48
	}
L48:
	;
	if v197 == int32(0) {
		v242 = v174
		v244 = v194
		v246 = v195
		v255 = v193
		goto L38
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(341991), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L18
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = int32(101422)
	F_errdetail(m, int32(658588), v26+int32(32))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L18
	} else {
		goto L52
	}
L52:
	;
	F_errhint(m, int32(622264), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L18
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(526193), int32(7297), int32(184085))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L18
	} else {
		goto L54
	}
L54:
	;
	v242 = v174
	v244 = v194
	v246 = v195
	v255 = v193
	goto L38
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88+v111<<(uint(int32(2))%32)))) = v266
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+79)))
	*(*uint8)(unsafe.Add(mBase, uint32(v111+v89))) = uint8(v270)
	v278 = v242
	goto L37
L56:
	;
	goto L34
L57:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v335 = v304
	goto L58
L58:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v323<<(uint(int32(4))%32)+int32(111)+v335*int32(100)))))
	if v355 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v360 = F_errstart(m, v56, int32(584281))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L18
	} else {
		goto L64
	}
L60:
	;
	v357 = v335 + int32(1)
	if v54 != v357 {
		v335 = v357
		goto L58
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	goto L59
L63:
	;
	v405 = v88
	v408 = v89
	goto L20
L64:
	;
	if v360 == int32(0) {
		v405 = v88
		v408 = v89
		goto L20
	} else {
		goto L65
	}
L65:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L18
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(341991), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L18
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = int32(101422)
	F_errdetail(m, int32(658588), v26+int32(16))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L18
	} else {
		goto L68
	}
L68:
	;
	F_errhint(m, int32(622264), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L18
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(526193), int32(7331), int32(184085))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L18
	} else {
		goto L70
	}
L70:
	;
	v405 = v88
	v408 = v89
	goto L20
L71:
	;
	F_deconstruct_expanded_record(m, l2)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L18
	} else {
		goto L74
	}
L72:
	;
	v424 = v416
	goto L73
L73:
	;
	v425 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+68)) = v425
	v428 = v424 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v428
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v431 = int32(4562096)
	v432 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v434
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	if v425 < v436 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v424 = v423
	goto L73
L75:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l2)+60))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v447 = int32(0)
	v448 = v436
	goto L78
L76:
	;
	v561 = v428
	goto L77
L77:
	;
	if v561&int32(64) != 0 {
		goto L106
	} else {
		goto L107
	}
L78:
	;
	v469 = v430 + int32(20) + v447<<(uint(int32(4))%32)
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+9)))
	if v470 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v561 = v537
	goto L77
L80:
	;
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447+v408))))
	v476 = v447 << (uint(int32(2)) % 32)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v405+v476)))
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+6)))
	if v479 != 0 {
		v523 = v478
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v530 = v448
	goto L82
L82:
	;
	v535 = v447 + int32(1)
	if v535 < v530 {
		v447 = v535
		v448 = v530
		goto L78
	} else {
		goto L105
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440+v476))) = v523
	*(*uint8)(unsafe.Add(mBase, uint32(v447+v439))) = uint8(v474)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	v530 = v529
	goto L82
L84:
	;
	if v474&int32(1) == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v484 = int32(*(*int16)(unsafe.Add(mBase, uint32(v469)+4)))
	if v484 != int32(-1) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v511 = v478
	goto L87
L87:
	;
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447+v439))))
	if v514 != 0 {
		v523 = v511
		goto L83
	} else {
		goto L99
	}
L88:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v507 | int32(8)
	v511 = v506
	goto L87
L89:
	;
	v504 = F_datumCopy(m, v478, int32(0), v484)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L18
	} else {
		goto L98
	}
L90:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478))))
	if v487 != int32(1) {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	if (v411^int32(-1))&int32(1) != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v490 = F_detoast_external_attr(m, v478)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L18
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v494 = F_datumCopy(m, v478, int32(0), int32(-1))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L18
	} else {
		goto L96
	}
L95:
	;
	v506 = v490
	goto L88
L96:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494))))
	if v496 != int32(1) {
		v506 = v494
		goto L88
	} else {
		goto L97
	}
L97:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v499 | int32(16)
	v506 = v494
	goto L88
L98:
	;
	v506 = v504
	goto L88
L99:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v440+v476)))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l2)+88))
	if base.Ui32(v517) <= base.Ui32(v516) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l2)+92))
	if base.Ui32(v516) < base.Ui32(v519) {
		v523 = v511
		goto L83
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	F_pfree(m, v516)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L18
	} else {
		goto L104
	}
L103:
	;
	goto L102
L104:
	;
	v523 = v511
	goto L83
L105:
	;
	goto L79
L106:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l2)+96))
	if v564 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v432
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v593)+16))
	if v598 != v594 {
		goto L117
	} else {
		goto L118
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v578
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	F_domain_check(m, l2+int32(18), int32(0), v584, l2+int32(104), v587)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L18
	} else {
		goto L115
	}
L110:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v572 = F_AllocSetContextCreateInternal(m, v567, int32(66675), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L18
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	F_MemoryContextReset(m, v564)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L18
	} else {
		goto L114
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+96)) = v572
	v578 = v572
	goto L109
L114:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l2)+96))
	v578 = v577
	goto L109
L115:
	;
	goto L108
L116:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v628 != 0 {
		goto L133
	} else {
		goto L134
	}
L117:
	;
	if v598 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	goto L119
L119:
	;
	goto L116
L120:
	;
	if v594 != 0 {
		goto L127
	} else {
		goto L128
	}
L121:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v593)+28))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v593)+24))
	if v603 != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	if v602 == int32(0) {
		goto L120
	} else {
		goto L126
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+28)) = v602
	goto L122
L124:
	;
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v598)+20)) = v602
	goto L122
L126:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v593)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v602)+24)) = v608
	goto L120
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v593)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v593)+16)) = v594
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v594)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v593)+28)) = v615
	if v615 != 0 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v593)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v593)+16)) = int32(0)
	goto L119
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v615)+24)) = v593
	goto L132
L131:
	;
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v594)+20)) = v593
	goto L116
L133:
	;
	F_DeleteExpandedObject(m, v628+int32(12))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L18
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = l2
	goto L8
L136:
	;
	goto L135
L137:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v638
	F_errmsg_internal(m, int32(508615), v26)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L18
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(526193), int32(7437), int32(184085))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L18
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v676 = int32(2)
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v675+v664<<(uint(v676)%32))))
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v674+v679<<(uint(v676)%32))))
	if v657 < v54 {
		goto L144
	} else {
		goto L145
	}
L141:
	;
	v824 = v793
	goto L9
L142:
	;
	F_exec_assign_value(m, l0, v683, v797, v799&int32(1), v795, v798)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L18
	} else {
		goto L159
	}
L143:
	;
	v778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v695))))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(l3+v695<<(uint(int32(2))%32))))
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v714)+76))
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v714)+68))
	v793 = v695 + int32(1)
	v795 = v784
	v797 = v782
	v798 = v783
	v799 = v778
	goto L142
L144:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v695 = v657
	goto L147
L145:
	;
	v727 = v657
	goto L146
L146:
	;
	v744 = int32(-1)
	v745 = int32(705)
	v746 = int32(1)
	v747 = int32(0)
	if v57 != 0 {
		v793 = v727
		v795 = v745
		v797 = v747
		v798 = v744
		v799 = v746
		goto L142
	} else {
		goto L151
	}
L147:
	;
	v714 = l5 + int32(20) + v685<<(uint(int32(4))%32) + v695*int32(100)
	v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v714)+91)))
	if v715 != int32(1) {
		goto L143
	} else {
		goto L149
	}
L148:
	;
	v727 = v54
	goto L146
L149:
	;
	v719 = v695 + int32(1)
	if v719 != v54 {
		v695 = v719
		goto L147
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	v749 = F_errstart(m, v56, int32(584281))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L18
	} else {
		goto L152
	}
L152:
	;
	if v749 == int32(0) {
		v793 = v727
		v795 = v745
		v797 = v747
		v798 = v744
		v799 = v746
		goto L142
	} else {
		goto L153
	}
L153:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L18
	} else {
		goto L154
	}
L154:
	;
	F_errmsg(m, int32(341991), int32(0))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L18
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = int32(101422)
	F_errdetail(m, int32(658588), v26-int32(-64))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L18
	} else {
		goto L156
	}
L156:
	;
	F_errhint(m, int32(622264), int32(0))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L18
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(526193), int32(7405), int32(184085))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L18
	} else {
		goto L158
	}
L158:
	;
	v793 = v727
	v795 = v745
	v797 = v747
	v798 = v744
	v799 = v746
	goto L142
L159:
	;
	v815 = v664 + int32(1)
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v815 < v816 {
		v657 = v793
		v664 = v815
		goto L140
	} else {
		goto L160
	}
L160:
	;
	goto L141
L161:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v855 = v824
	goto L162
L162:
	;
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v843<<(uint(int32(4))%32)+int32(111)+v855*int32(100)))))
	if v875 != 0 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v880 = F_errstart(m, v56, int32(584281))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L18
	} else {
		goto L168
	}
L164:
	;
	v877 = v855 + int32(1)
	if v54 != v877 {
		v855 = v877
		goto L162
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	goto L163
L167:
	;
	goto L8
L168:
	;
	if v880 == int32(0) {
		goto L8
	} else {
		goto L169
	}
L169:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L18
	} else {
		goto L170
	}
L170:
	;
	F_errmsg(m, int32(341991), int32(0))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L18
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = int32(101422)
	F_errdetail(m, int32(658588), v26+int32(48))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L18
	} else {
		goto L172
	}
L172:
	;
	F_errhint(m, int32(622264), int32(0))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L18
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(526193), int32(7431), int32(184085))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L18
	} else {
		goto L174
	}
L174:
	;
	goto L8
}
func F_exp_var(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v49 float64
	_ = v49
	var v50 int32
	_ = v50
	var v51 float64
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int64
	_ = v62
	var v69 float64
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v85 float64
	_ = v85
	var v90 float64
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 float64
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v14
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = F_palloc(m, v20<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v25))) = uint16(v27)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v27 < v29 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v36 = v29 << (uint(int32(1)) % 32)
	if v36 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v39
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v25 + int32(2)
	v49 = F_numericvar_to_double_no_overflow(m, v12+int32(24))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L6:
	;
	goto L5
L7:
	;
	v37 = F__emscripten_memcpy_bulkmem(m, v25+int32(2), v34, v36)
	mBase = m.M
	goto L9
L8:
	;
	goto L9
L9:
	;
	goto L6
L10:
	;
	m.G0 = v12 + int32(48)
	return
L11:
	;
	if int32(0) < v106 {
		goto L55
	} else {
		goto L56
	}
L12:
	;
	v173 = v145
	goto L49
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L45
	}
L14:
	;
	v51 = base.F64_abs(v49)
	if base.F64_ge(v51, float64(6000)) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if base.F64_gt(v49, float64(0)) != 0 {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v69 = base.F64_mul(v49, float64(0.434294481903252))
	if base.F64_lt(base.F64_abs(v69), float64(2.147483648e+09)) != 0 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v56 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_pfree(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	v62 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v62
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v62
	goto L10
L22:
	;
	goto L21
L23:
	;
	if base.F64_gt(v51, float64(0.01)) != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v73 = base.I32_trunc_f64_s(v69)
	v75 = v73
	goto L23
L25:
	;
	goto L26
L26:
	;
	v75 = int32(-2147483648)
	goto L23
L27:
	;
	v82 = int32(1)
	v85 = v49
	goto L30
L28:
	;
	v106 = int32(0)
	goto L29
L29:
	;
	F_add_var(m, int32(1774696), v12+int32(24), l1)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L34
	}
L30:
	;
	v90 = base.F64_mul(v85, float64(0.5))
	if base.F64_gt(base.F64_abs(v90), float64(0.01)) != 0 {
		v82 = v82 + int32(1)
		v85 = v90
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v95 = v12 + int32(24)
	v96 = int32(1)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	F_div_var_int(m, v95, v96<<(uint(v82)%32), int32(0), v95, v101+v82, v96)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	v106 = v82
	goto L29
L34:
	;
	v121 = v12 + int32(24)
	v126 = base.F64_mul(base.F64_convert_i32_s(v106), float64(0.301029995663981))
	if base.F64_lt(base.F64_abs(v126), float64(2.147483648e+09)) != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v136 = v132 + (l2 + v75) + int32(1)
	v137 = int32(0)
	if v137 < v136 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v130 = base.I32_trunc_f64_s(v126)
	v132 = v130
	goto L35
L37:
	;
	goto L38
L38:
	;
	v132 = int32(-2147483648)
	goto L35
L39:
	;
	v140 = v136
	goto L41
L40:
	;
	v140 = v137
	goto L41
L41:
	;
	v142 = v140 + int32(7)
	F_mul_var(m, v121, v121, v12, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v145 = int32(2)
	F_div_var_int(m, v12, v145, int32(0), v12, v142, int32(1))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v151 != 0 {
		goto L12
	} else {
		goto L44
	}
L44:
	;
	goto L11
L45:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(120003), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(526183), int32(10929), int32(241552))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_add_var(m, l1, v12, l1)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L51
	}
L50:
	;
	goto L11
L51:
	;
	F_mul_var(m, v12, v12+int32(24), v12, v142)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v183 = int32(1)
	v184 = v173 + v183
	F_div_var_int(m, v12, v184, int32(0), v12, v142, v183)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v189 != 0 {
		v173 = v184
		goto L49
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	v203 = v106
	goto L58
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = l2
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v244 = l2 + v241<<(uint(int32(2))%32)
	if v244+int32(4) < int32(0) {
		goto L66
	} else {
		goto L67
	}
L58:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v215 = v140 + int32(8) - v212<<(uint(int32(3))%32)
	v216 = int32(0)
	if v216 < v215 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L57
L60:
	;
	v219 = v215
	goto L62
L61:
	;
	v219 = v216
	goto L62
L62:
	;
	F_mul_var(m, l1, l1, l1, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v222 = int32(1)
	if base.Ui32(v222) < base.Ui32(v203) {
		v203 = v203 - v222
		goto L58
	} else {
		goto L64
	}
L64:
	;
	goto L59
L65:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	if v360 != 0 {
		goto L92
	} else {
		goto L93
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
	goto L65
L67:
	;
	goto L68
L68:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v255 = l2 & int32(3)
	v259 = base.I32_div_s(v244+int32(7), int32(4))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v260 <= v259 {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	goto L65
L70:
	;
	if int32(0) <= v326 {
		goto L69
	} else {
		goto L91
	}
L71:
	;
	v306 = v300
	goto L85
L72:
	;
	v273 = int32(1)
	v274 = v259 - v273
	v277 = v253 + v274<<(uint(v273)%32)
	v278 = int32(*(*int16)(unsafe.Add(mBase, uint32(v277))))
	v279 = int32(2)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v255<<(uint(v279)%32))+uint32(_consts[1101])))
	v284 = base.I32_rem_s(v278, v283)
	v285 = v278 - v284
	*(*uint16)(unsafe.Add(mBase, uint32(v277))) = uint16(v285)
	v288 = base.I32_div_s(v283, v279)
	if v284 < v288 {
		v326 = v274
		goto L70
	} else {
		goto L80
	}
L73:
	;
	if v255 == int32(0) {
		goto L69
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v259
	if v255 != 0 {
		goto L72
	} else {
		goto L78
	}
L76:
	;
	if v259 != v260 {
		goto L69
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v259
	goto L72
L78:
	;
	v270 = int32(*(*int16)(unsafe.Add(mBase, uint32(v253+v259<<(uint(int32(1))%32)))))
	if v270 <= int32(4999) {
		v326 = v259
		goto L70
	} else {
		goto L79
	}
L79:
	;
	v300 = v259
	goto L71
L80:
	;
	v291 = v283 + base.I32_extend16_s(v285)
	if int32(9999) < v291 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v296 = v291 + int32(55536)
	goto L83
L82:
	;
	v296 = v291
	goto L83
L83:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v277))) = uint16(v296)
	if v291 < int32(10000) {
		v326 = v274
		goto L70
	} else {
		goto L84
	}
L84:
	;
	v300 = v274
	goto L71
L85:
	;
	v312 = int32(1)
	v313 = v306 - v312
	v316 = v253 + v313<<(uint(v312)%32)
	v319 = int32(*(*int16)(unsafe.Add(mBase, uint32(v316))))
	v321 = base.B2i32(int32(9998) < v319)
	if int32(9998) < v319 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v326 = v313
	goto L70
L87:
	;
	v322 = int32(-9999)
	goto L89
L88:
	;
	v322 = v312
	goto L89
L89:
	;
	v323 = v322 + v319
	*(*uint16)(unsafe.Add(mBase, uint32(v316))) = uint16(v323)
	if int32(9998) < v319 {
		v306 = v313
		goto L85
	} else {
		goto L90
	}
L90:
	;
	goto L86
L91:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v334 - int32(2)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v339 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v338 + v339
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v342 + v339
	goto L69
L92:
	;
	F_pfree(m, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v363 == int32(0) {
		goto L10
	} else {
		goto L96
	}
L95:
	;
	goto L94
L96:
	;
	F_pfree(m, v363)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	goto L10
}
func F_expand_groupingset_node(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
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
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
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
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
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
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v13 {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v177 = v2
		goto L1
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v177
L2:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v148 == int32(0) {
		v177 = v2
		goto L1
	} else {
		goto L44
	}
L3:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v89 != 0 {
		goto L25
	} else {
		goto L26
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v31 == int32(0) {
		v79 = v2
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v23
	v29 = F_list_make1_impl(m, int32(1), v11+int32(4))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L9
	}
L6:
	;
	v14 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v14
	v19 = F_list_make1_impl(m, int32(1), v11)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v177 = v19
	goto L1
L9:
	;
	v177 = v29
	goto L1
L10:
	;
	v87 = F_lappend(m, v79, int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L24
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v34 <= int32(0) {
		v79 = v2
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v37 = v34
	v38 = v2
	goto L13
L13:
	;
	v45 = int32(0)
	v47 = v45
	v49 = v45
	v50 = v37
	goto L15
L14:
	;
	v79 = v74
	goto L10
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v47 < v55 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v74 = F_lappend(m, v38, v70)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L22
	}
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v47<<(uint(int32(2))%32))))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v63 = F_list_concat(m, v49, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L20
	}
L18:
	;
	v70 = v49
	goto L19
L19:
	;
	goto L16
L20:
	;
	v65 = int32(1)
	v68 = v50 - v65
	if v68 != 0 {
		v47 = v47 + v65
		v49 = v63
		v50 = v68
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v70 = v63
	goto L19
L22:
	;
	if int32(1) < v37 {
		v37 = v37 - int32(1)
		v38 = v74
		goto L13
	} else {
		goto L23
	}
L23:
	;
	goto L14
L24:
	;
	v177 = v87
	goto L1
L25:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v91 = v90
	goto L27
L26:
	;
	v91 = v2
	goto L27
L27:
	;
	v93 = v2
	v98 = v2
	goto L28
L28:
	;
	if v89 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v177 = v141
	goto L1
L30:
	;
	v141 = F_lappend(m, v93, v137)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L7
	} else {
		goto L42
	}
L31:
	;
	v137 = int32(0)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v103 = int32(0)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v106 <= v103 {
		v137 = v103
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v109 = int32(1)
	v111 = v103
	v113 = v103
	goto L35
L35:
	;
	if v109&v98 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v137 = v126
	goto L30
L37:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118+v111<<(uint(int32(2))%32))))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	v124 = F_list_concat(m, v113, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L7
	} else {
		goto L40
	}
L38:
	;
	v126 = v113
	goto L39
L39:
	;
	v127 = int32(1)
	v130 = v111 + v127
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v130 < v131 {
		v109 = v109 << (uint(v127) % 32)
		v111 = v130
		v113 = v126
		goto L35
	} else {
		goto L41
	}
L40:
	;
	v126 = v124
	goto L39
L41:
	;
	goto L36
L42:
	;
	v144 = v98 + int32(1)
	if int32(base.Ui32(v144)>>(uint(v91)%32)) == int32(0) {
		v93 = v141
		v98 = v144
		goto L28
	} else {
		goto L43
	}
L43:
	;
	goto L29
L44:
	;
	v151 = int32(0)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v152 <= v151 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v177 = v2
	goto L1
L46:
	;
	goto L47
L47:
	;
	v155 = v151
	v156 = v2
	goto L48
L48:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v163+v155<<(uint(int32(2))%32))))
	v168 = F_expand_groupingset_node(m, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L7
	} else {
		goto L50
	}
L49:
	;
	v177 = v170
	goto L1
L50:
	;
	v170 = F_list_concat(m, v156, v168)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	v173 = v155 + int32(1)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v173 < v174 {
		v155 = v173
		v156 = v170
		goto L48
	} else {
		goto L52
	}
L52:
	;
	goto L49
}
func F_explicit_bzero(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	v5 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(0)), l1)
	return
}
func F_exprs_known_equal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
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
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	v5 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v11 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v14 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = v5
	goto L4
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v22<<(uint(int32(2))%32))))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+41)))
	if v32 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	v133 = v22 + int32(1)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v133 < v134 {
		v22 = v133
		goto L4
	} else {
		goto L40
	}
L7:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v34 = int32(0)
	if v33 == v34 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	if v75 == int32(0) {
		goto L6
	} else {
		goto L25
	}
L11:
	;
	if v72 == int32(0) {
		goto L6
	} else {
		goto L24
	}
L12:
	;
	v72 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v40 <= int32(0) {
		v65 = v34
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v72 = v65
	goto L11
L16:
	;
	v43 = int32(0)
	if v43 < v40 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v46 = v40
	goto L19
L18:
	;
	v46 = v43
	goto L19
L19:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v49 = int32(0)
	goto L20
L20:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v47+v49<<(uint(int32(2))%32))))
	v58 = base.B2i32(v57 == l3)
	if v57 == l3 {
		v65 = v58
		goto L15
	} else {
		goto L22
	}
L21:
	;
	v65 = v58
	goto L15
L22:
	;
	v60 = v49 + int32(1)
	if v60 != v46 {
		v49 = v60
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L10
L25:
	;
	v78 = int32(0)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v81 <= v78 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v84 = v78
	v91 = v78
	v92 = v78
	goto L27
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v84<<(uint(int32(2))%32))))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v100 = F_equal(m, l1, v99)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	return int32(1)
L29:
	;
	if v109&int32(1)&v110 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	return int32(0)
L31:
	;
	if v100 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v109 = v91
	v110 = int32(1)
	goto L29
L33:
	;
	goto L34
L34:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v106 = F_equal(m, l2, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	v109 = v106 | v91
	v110 = v92
	goto L29
L36:
	;
	v117 = v84 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v118 <= v117 {
		goto L6
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L28
L39:
	;
	v84 = v117
	v91 = v109
	v92 = v110
	goto L27
L40:
	;
	goto L5
}
