package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecAlterOwnerStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
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
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int64
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v609 int64
	_ = v609
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v640 int64
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v668 int64
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v19 = F_get_rolespec_oid(m, v17, int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		switch v21 - int32(1) {
		case 0, 6, 7, 18, 20, 21, 23, 24, 25, 28, 34, 39, 42, 45, 46:
			v551 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v552 = int32(0)
			F_get_object_address(m, l0, v21, v551, v552, int32(8), v552)
			mBase = m.M
			v556 = m.ExcPending
			if v556 != 0 {
				return
			} else {
				v557 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_AlterObjectOwner_internal(m, v557, v558, v19)
				mBase = m.M
				v560 = m.ExcPending
				if v560 != 0 {
					return
				} else {
					m.G0 = v15 + int32(16)
					return
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v564 = m.ExcPending
			if v564 != 0 {
				return
			} else {
				v565 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v565
				F_errmsg_internal(m, int32(_a_F_ExecAlterOwnerStmt_0), v15)
				mBase = m.M
				v569 = m.ExcPending
				if v569 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_1), int32(902), int32(_a_F_ExecAlterOwnerStmt_2))
					mBase = m.M
					v574 = m.ExcPending
					if v574 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 8:
			v575 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v576 = *(*int32)(unsafe.Add(mBase, uint32(v575)+4))
			v577 = m.G0
			v579 = v577 - int32(288)
			m.G0 = v579
			v583 = F_table_open(m, int32(1262), int32(3))
			mBase = m.M
			v584 = m.ExcPending
			if v584 != 0 {
				return
			} else {
				v586 = v579 + int32(232)
				F_ScanKeyInit(m, v586, int32(2), int32(3), int32(62), base.I64_extend_i32_u(v576))
				mBase = m.M
				v592 = m.ExcPending
				if v592 != 0 {
					return
				} else {
					v594 = int32(1)
					v597 = F_systable_beginscan(m, v583, int32(2671), v594, int32(0), v594, v586)
					mBase = m.M
					v598 = m.ExcPending
					if v598 != 0 {
						return
					} else {
						v599 = F_systable_getnext(m, v597)
						mBase = m.M
						v600 = m.ExcPending
						if v600 != 0 {
							return
						} else {
							if v599 != 0 {
								v601 = *(*int32)(unsafe.Add(mBase, uint32(v599)+16))
								v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+22)))
								v603 = v601 + v602
								v604 = *(*int32)(unsafe.Add(mBase, uint32(v603)))
								v605 = *(*int32)(unsafe.Add(mBase, uint32(v603)+68))
								if v19 != v605 {
									v607 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v579)+64)) = uint16(v607)
									v609 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v579)+56)) = v609
									*(*int64)(unsafe.Add(mBase, uint32(v579)+48)) = v609
									*(*uint16)(unsafe.Add(mBase, uint32(v579)+32)) = uint16(v607)
									*(*int64)(unsafe.Add(mBase, uint32(v579)+24)) = v609
									*(*int64)(unsafe.Add(mBase, uint32(v579)+16)) = v609
									v621 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
									v622 = F_object_ownercheck(m, int32(1262), v604, v621)
									mBase = m.M
									v623 = m.ExcPending
									if v623 != 0 {
										return
									} else {
										if v622 == int32(0) {
											F_aclcheck_error(m, int32(2), int32(9), v576)
											mBase = m.M
											v629 = m.ExcPending
											if v629 != 0 {
												return
											} else {
												v631 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
												F_check_can_set_role(m, v631, v19)
												mBase = m.M
												v633 = m.ExcPending
												if v633 != 0 {
													return
												} else {
													v634 = F_superuser(m)
													mBase = m.M
													v635 = m.ExcPending
													if v635 != 0 {
														return
													} else {
														if v634 == int32(0) {
															v640 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0])))
															v641 = F_SearchSysCache1(m, int32(11), v640)
															mBase = m.M
															v642 = m.ExcPending
															if v642 != 0 {
																return
															} else {
																if v641 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v751 = m.ExcPending
																	if v751 != 0 {
																		return
																	} else {
																		F_errcode(m, int32(16797828))
																		mBase = m.M
																		v754 = m.ExcPending
																		if v754 != 0 {
																			return
																		} else {
																			F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_3), int32(0))
																			mBase = m.M
																			v758 = m.ExcPending
																			if v758 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_4), int32(2748), int32(_a_F_ExecAlterOwnerStmt_5))
																				mBase = m.M
																				v763 = m.ExcPending
																				if v763 != 0 {
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
																	v645 = *(*int32)(unsafe.Add(mBase, uint32(v641)+16))
																	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645)+22)))
																	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645+v646)+71)))
																	F_ReleaseCatCache(m, v641)
																	mBase = m.M
																	v650 = m.ExcPending
																	if v650 != 0 {
																		return
																	} else {
																		if v648 == int32(0) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v751 = m.ExcPending
																			if v751 != 0 {
																				return
																			} else {
																				F_errcode(m, int32(16797828))
																				mBase = m.M
																				v754 = m.ExcPending
																				if v754 != 0 {
																					return
																				} else {
																					F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_3), int32(0))
																					mBase = m.M
																					v758 = m.ExcPending
																					if v758 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_4), int32(2748), int32(_a_F_ExecAlterOwnerStmt_5))
																						mBase = m.M
																						v763 = m.ExcPending
																						if v763 != 0 {
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
																			v656 = v599 + int32(4)
																			F_LockTuple(m, v583, v656, int32(7))
																			mBase = m.M
																			v659 = m.ExcPending
																			if v659 != 0 {
																				return
																			} else {
																				*(*int64)(unsafe.Add(mBase, uint32(v579)+96)) = base.I64_extend_i32_u(v19)
																				v662 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v579)+18)) = uint8(v662)
																				v665 = *(*int32)(unsafe.Add(mBase, uint32(v583)+52))
																				v668 = F_heap_getattr_6(m, v599, int32(18), v665, v579+int32(15))
																				mBase = m.M
																				v669 = m.ExcPending
																				if v669 != 0 {
																					return
																				} else {
																					v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+15)))
																					if v670 == int32(0) {
																						v674 = F_pg_detoast_datum(m, base.I32_wrap_i64(v668))
																						mBase = m.M
																						v675 = m.ExcPending
																						if v675 != 0 {
																							return
																						} else {
																							v676 = *(*int32)(unsafe.Add(mBase, uint32(v603)+68))
																							v677 = F_aclnewowner(m, v674, v676, v19)
																							mBase = m.M
																							v678 = m.ExcPending
																							if v678 != 0 {
																								return
																							} else {
																								v679 = int32(1)
																								*(*uint8)(unsafe.Add(mBase, uint32(v579)+33)) = uint8(v679)
																								*(*int64)(unsafe.Add(mBase, uint32(v579)+216)) = base.I64_extend_i32_u(v677)
																								v684 = *(*int32)(unsafe.Add(mBase, uint32(v583)+52))
																								v691 = F_heap_modify_tuple(m, v599, v684, v579+int32(80), v579+int32(48), v579+int32(16))
																								mBase = m.M
																								v692 = m.ExcPending
																								if v692 != 0 {
																									return
																								} else {
																									F_CatalogTupleUpdate(m, v583, v691+int32(4), v691)
																									mBase = m.M
																									v696 = m.ExcPending
																									if v696 != 0 {
																										return
																									} else {
																										F_UnlockTuple(m, v583, v656, int32(7))
																										mBase = m.M
																										v699 = m.ExcPending
																										if v699 != 0 {
																											return
																										} else {
																											F_pfree(m, v691)
																											mBase = m.M
																											v701 = m.ExcPending
																											if v701 != 0 {
																												return
																											} else {
																												F_changeDependencyOnOwner(m, int32(1262), v604, v19)
																												mBase = m.M
																												v704 = m.ExcPending
																												if v704 != 0 {
																													return
																												} else {
																													v711 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[1]))
																													if v711 != 0 {
																														v713 = int32(0)
																														F_RunObjectPostAlterHook(m, int32(1262), v604, v713, v713, v713)
																														mBase = m.M
																														v717 = m.ExcPending
																														if v717 != 0 {
																															return
																														} else {
																															*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v604
																															*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																															F_systable_endscan(m, v597)
																															mBase = m.M
																															v724 = m.ExcPending
																															if v724 != 0 {
																																return
																															} else {
																																F_relation_close(m, v583, int32(0))
																																mBase = m.M
																																v727 = m.ExcPending
																																if v727 != 0 {
																																	return
																																} else {
																																	m.G0 = v579 + int32(288)
																																	m.G0 = v15 + int32(16)
																																	return
																																}
																															}
																														}
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v604
																														*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																														F_systable_endscan(m, v597)
																														mBase = m.M
																														v724 = m.ExcPending
																														if v724 != 0 {
																															return
																														} else {
																															F_relation_close(m, v583, int32(0))
																															mBase = m.M
																															v727 = m.ExcPending
																															if v727 != 0 {
																																return
																															} else {
																																m.G0 = v579 + int32(288)
																																m.G0 = v15 + int32(16)
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
																						v684 = *(*int32)(unsafe.Add(mBase, uint32(v583)+52))
																						v691 = F_heap_modify_tuple(m, v599, v684, v579+int32(80), v579+int32(48), v579+int32(16))
																						mBase = m.M
																						v692 = m.ExcPending
																						if v692 != 0 {
																							return
																						} else {
																							F_CatalogTupleUpdate(m, v583, v691+int32(4), v691)
																							mBase = m.M
																							v696 = m.ExcPending
																							if v696 != 0 {
																								return
																							} else {
																								F_UnlockTuple(m, v583, v656, int32(7))
																								mBase = m.M
																								v699 = m.ExcPending
																								if v699 != 0 {
																									return
																								} else {
																									F_pfree(m, v691)
																									mBase = m.M
																									v701 = m.ExcPending
																									if v701 != 0 {
																										return
																									} else {
																										F_changeDependencyOnOwner(m, int32(1262), v604, v19)
																										mBase = m.M
																										v704 = m.ExcPending
																										if v704 != 0 {
																											return
																										} else {
																											v711 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[1]))
																											if v711 != 0 {
																												v713 = int32(0)
																												F_RunObjectPostAlterHook(m, int32(1262), v604, v713, v713, v713)
																												mBase = m.M
																												v717 = m.ExcPending
																												if v717 != 0 {
																													return
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v604
																													*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																													F_systable_endscan(m, v597)
																													mBase = m.M
																													v724 = m.ExcPending
																													if v724 != 0 {
																														return
																													} else {
																														F_relation_close(m, v583, int32(0))
																														mBase = m.M
																														v727 = m.ExcPending
																														if v727 != 0 {
																															return
																														} else {
																															m.G0 = v579 + int32(288)
																															m.G0 = v15 + int32(16)
																															return
																														}
																													}
																												}
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v604
																												*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																												F_systable_endscan(m, v597)
																												mBase = m.M
																												v724 = m.ExcPending
																												if v724 != 0 {
																													return
																												} else {
																													F_relation_close(m, v583, int32(0))
																													mBase = m.M
																													v727 = m.ExcPending
																													if v727 != 0 {
																														return
																													} else {
																														m.G0 = v579 + int32(288)
																														m.G0 = v15 + int32(16)
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
														} else {
															v656 = v599 + int32(4)
															F_LockTuple(m, v583, v656, int32(7))
															mBase = m.M
															v659 = m.ExcPending
															if v659 != 0 {
																return
															} else {
																*(*int64)(unsafe.Add(mBase, uint32(v579)+96)) = base.I64_extend_i32_u(v19)
																v662 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v579)+18)) = uint8(v662)
																v665 = *(*int32)(unsafe.Add(mBase, uint32(v583)+52))
																v668 = F_heap_getattr_6(m, v599, int32(18), v665, v579+int32(15))
																mBase = m.M
																v669 = m.ExcPending
																if v669 != 0 {
																	return
																} else {
																	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+15)))
																	if v670 == int32(0) {
																		v674 = F_pg_detoast_datum(m, base.I32_wrap_i64(v668))
																		mBase = m.M
																		v675 = m.ExcPending
																		if v675 != 0 {
																			return
																		} else {
																			v676 = *(*int32)(unsafe.Add(mBase, uint32(v603)+68))
																			v677 = F_aclnewowner(m, v674, v676, v19)
																			mBase = m.M
																			v678 = m.ExcPending
																			if v678 != 0 {
																				return
																			} else {
																				v679 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v579)+33)) = uint8(v679)
																				*(*int64)(unsafe.Add(mBase, uint32(v579)+216)) = base.I64_extend_i32_u(v677)
																				v684 = *(*int32)(unsafe.Add(mBase, uint32(v583)+52))
																				v691 = F_heap_modify_tuple(m, v599, v684, v579+int32(80), v579+int32(48), v579+int32(16))
																				mBase = m.M
																				v692 = m.ExcPending
																				if v692 != 0 {
																					return
																				} else {
																					F_CatalogTupleUpdate(m, v583, v691+int32(4), v691)
																					mBase = m.M
																					v696 = m.ExcPending
																					if v696 != 0 {
																						return
																					} else {
																						F_UnlockTuple(m, v583, v656, int32(7))
																						mBase = m.M
																						v699 = m.ExcPending
																						if v699 != 0 {
																							return
																						} else {
																							F_pfree(m, v691)
																							mBase = m.M
																							v701 = m.ExcPending
																							if v701 != 0 {
																								return
																							} else {
																								F_changeDependencyOnOwner(m, int32(1262), v604, v19)
																								mBase = m.M
																								v704 = m.ExcPending
																								if v704 != 0 {
																									return
																								} else {
																									v711 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[1]))
																									if v711 != 0 {
																										v713 = int32(0)
																										F_RunObjectPostAlterHook(m, int32(1262), v604, v713, v713, v713)
																										mBase = m.M
																										v717 = m.ExcPending
																										if v717 != 0 {
																											return
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v604
																											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																											F_systable_endscan(m, v597)
																											mBase = m.M
																											v724 = m.ExcPending
																											if v724 != 0 {
																												return
																											} else {
																												F_relation_close(m, v583, int32(0))
																												mBase = m.M
																												v727 = m.ExcPending
																												if v727 != 0 {
																													return
																												} else {
																													m.G0 = v579 + int32(288)
																													m.G0 = v15 + int32(16)
																													return
																												}
																											}
																										}
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v604
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																										F_systable_endscan(m, v597)
																										mBase = m.M
																										v724 = m.ExcPending
																										if v724 != 0 {
																											return
																										} else {
																											F_relation_close(m, v583, int32(0))
																											mBase = m.M
																											v727 = m.ExcPending
																											if v727 != 0 {
																												return
																											} else {
																												m.G0 = v579 + int32(288)
																												m.G0 = v15 + int32(16)
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
																		v684 = *(*int32)(unsafe.Add(mBase, uint32(v583)+52))
																		v691 = F_heap_modify_tuple(m, v599, v684, v579+int32(80), v579+int32(48), v579+int32(16))
																		mBase = m.M
																		v692 = m.ExcPending
																		if v692 != 0 {
																			return
																		} else {
																			F_CatalogTupleUpdate(m, v583, v691+int32(4), v691)
																			mBase = m.M
																			v696 = m.ExcPending
																			if v696 != 0 {
																				return
																			} else {
																				F_UnlockTuple(m, v583, v656, int32(7))
																				mBase = m.M
																				v699 = m.ExcPending
																				if v699 != 0 {
																					return
																				} else {
																					F_pfree(m, v691)
																					mBase = m.M
																					v701 = m.ExcPending
																					if v701 != 0 {
																						return
																					} else {
																						F_changeDependencyOnOwner(m, int32(1262), v604, v19)
																						mBase = m.M
																						v704 = m.ExcPending
																						if v704 != 0 {
																							return
																						} else {
																							v711 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[1]))
																							if v711 != 0 {
																								v713 = int32(0)
																								F_RunObjectPostAlterHook(m, int32(1262), v604, v713, v713, v713)
																								mBase = m.M
																								v717 = m.ExcPending
																								if v717 != 0 {
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v604
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																									F_systable_endscan(m, v597)
																									mBase = m.M
																									v724 = m.ExcPending
																									if v724 != 0 {
																										return
																									} else {
																										F_relation_close(m, v583, int32(0))
																										mBase = m.M
																										v727 = m.ExcPending
																										if v727 != 0 {
																											return
																										} else {
																											m.G0 = v579 + int32(288)
																											m.G0 = v15 + int32(16)
																											return
																										}
																									}
																								}
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v604
																								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																								F_systable_endscan(m, v597)
																								mBase = m.M
																								v724 = m.ExcPending
																								if v724 != 0 {
																									return
																								} else {
																									F_relation_close(m, v583, int32(0))
																									mBase = m.M
																									v727 = m.ExcPending
																									if v727 != 0 {
																										return
																									} else {
																										m.G0 = v579 + int32(288)
																										m.G0 = v15 + int32(16)
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
										} else {
											v631 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
											F_check_can_set_role(m, v631, v19)
											mBase = m.M
											v633 = m.ExcPending
											if v633 != 0 {
												return
											} else {
												v634 = F_superuser(m)
												mBase = m.M
												v635 = m.ExcPending
												if v635 != 0 {
													return
												} else {
													if v634 == int32(0) {
														v640 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0])))
														v641 = F_SearchSysCache1(m, int32(11), v640)
														mBase = m.M
														v642 = m.ExcPending
														if v642 != 0 {
															return
														} else {
															if v641 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v751 = m.ExcPending
																if v751 != 0 {
																	return
																} else {
																	F_errcode(m, int32(16797828))
																	mBase = m.M
																	v754 = m.ExcPending
																	if v754 != 0 {
																		return
																	} else {
																		F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_3), int32(0))
																		mBase = m.M
																		v758 = m.ExcPending
																		if v758 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_4), int32(2748), int32(_a_F_ExecAlterOwnerStmt_5))
																			mBase = m.M
																			v763 = m.ExcPending
																			if v763 != 0 {
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
																v645 = *(*int32)(unsafe.Add(mBase, uint32(v641)+16))
																v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645)+22)))
																v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645+v646)+71)))
																F_ReleaseCatCache(m, v641)
																mBase = m.M
																v650 = m.ExcPending
																if v650 != 0 {
																	return
																} else {
																	if v648 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v751 = m.ExcPending
																		if v751 != 0 {
																			return
																		} else {
																			F_errcode(m, int32(16797828))
																			mBase = m.M
																			v754 = m.ExcPending
																			if v754 != 0 {
																				return
																			} else {
																				F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_3), int32(0))
																				mBase = m.M
																				v758 = m.ExcPending
																				if v758 != 0 {
																					return
																				} else {
																					F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_4), int32(2748), int32(_a_F_ExecAlterOwnerStmt_5))
																					mBase = m.M
																					v763 = m.ExcPending
																					if v763 != 0 {
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
																		v656 = v599 + int32(4)
																		F_LockTuple(m, v583, v656, int32(7))
																		mBase = m.M
																		v659 = m.ExcPending
																		if v659 != 0 {
																			return
																		} else {
																			*(*int64)(unsafe.Add(mBase, uint32(v579)+96)) = base.I64_extend_i32_u(v19)
																			v662 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v579)+18)) = uint8(v662)
																			v665 = *(*int32)(unsafe.Add(mBase, uint32(v583)+52))
																			v668 = F_heap_getattr_6(m, v599, int32(18), v665, v579+int32(15))
																			mBase = m.M
																			v669 = m.ExcPending
																			if v669 != 0 {
																				return
																			} else {
																				v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+15)))
																				if v670 == int32(0) {
																					v674 = F_pg_detoast_datum(m, base.I32_wrap_i64(v668))
																					mBase = m.M
																					v675 = m.ExcPending
																					if v675 != 0 {
																						return
																					} else {
																						v676 = *(*int32)(unsafe.Add(mBase, uint32(v603)+68))
																						v677 = F_aclnewowner(m, v674, v676, v19)
																						mBase = m.M
																						v678 = m.ExcPending
																						if v678 != 0 {
																							return
																						} else {
																							v679 = int32(1)
																							*(*uint8)(unsafe.Add(mBase, uint32(v579)+33)) = uint8(v679)
																							*(*int64)(unsafe.Add(mBase, uint32(v579)+216)) = base.I64_extend_i32_u(v677)
																							v684 = *(*int32)(unsafe.Add(mBase, uint32(v583)+52))
																							v691 = F_heap_modify_tuple(m, v599, v684, v579+int32(80), v579+int32(48), v579+int32(16))
																							mBase = m.M
																							v692 = m.ExcPending
																							if v692 != 0 {
																								return
																							} else {
																								F_CatalogTupleUpdate(m, v583, v691+int32(4), v691)
																								mBase = m.M
																								v696 = m.ExcPending
																								if v696 != 0 {
																									return
																								} else {
																									F_UnlockTuple(m, v583, v656, int32(7))
																									mBase = m.M
																									v699 = m.ExcPending
																									if v699 != 0 {
																										return
																									} else {
																										F_pfree(m, v691)
																										mBase = m.M
																										v701 = m.ExcPending
																										if v701 != 0 {
																											return
																										} else {
																											F_changeDependencyOnOwner(m, int32(1262), v604, v19)
																											mBase = m.M
																											v704 = m.ExcPending
																											if v704 != 0 {
																												return
																											} else {
																												v711 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[1]))
																												if v711 != 0 {
																													v713 = int32(0)
																													F_RunObjectPostAlterHook(m, int32(1262), v604, v713, v713, v713)
																													mBase = m.M
																													v717 = m.ExcPending
																													if v717 != 0 {
																														return
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v604
																														*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																														F_systable_endscan(m, v597)
																														mBase = m.M
																														v724 = m.ExcPending
																														if v724 != 0 {
																															return
																														} else {
																															F_relation_close(m, v583, int32(0))
																															mBase = m.M
																															v727 = m.ExcPending
																															if v727 != 0 {
																																return
																															} else {
																																m.G0 = v579 + int32(288)
																																m.G0 = v15 + int32(16)
																																return
																															}
																														}
																													}
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v604
																													*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																													F_systable_endscan(m, v597)
																													mBase = m.M
																													v724 = m.ExcPending
																													if v724 != 0 {
																														return
																													} else {
																														F_relation_close(m, v583, int32(0))
																														mBase = m.M
																														v727 = m.ExcPending
																														if v727 != 0 {
																															return
																														} else {
																															m.G0 = v579 + int32(288)
																															m.G0 = v15 + int32(16)
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
																					v684 = *(*int32)(unsafe.Add(mBase, uint32(v583)+52))
																					v691 = F_heap_modify_tuple(m, v599, v684, v579+int32(80), v579+int32(48), v579+int32(16))
																					mBase = m.M
																					v692 = m.ExcPending
																					if v692 != 0 {
																						return
																					} else {
																						F_CatalogTupleUpdate(m, v583, v691+int32(4), v691)
																						mBase = m.M
																						v696 = m.ExcPending
																						if v696 != 0 {
																							return
																						} else {
																							F_UnlockTuple(m, v583, v656, int32(7))
																							mBase = m.M
																							v699 = m.ExcPending
																							if v699 != 0 {
																								return
																							} else {
																								F_pfree(m, v691)
																								mBase = m.M
																								v701 = m.ExcPending
																								if v701 != 0 {
																									return
																								} else {
																									F_changeDependencyOnOwner(m, int32(1262), v604, v19)
																									mBase = m.M
																									v704 = m.ExcPending
																									if v704 != 0 {
																										return
																									} else {
																										v711 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[1]))
																										if v711 != 0 {
																											v713 = int32(0)
																											F_RunObjectPostAlterHook(m, int32(1262), v604, v713, v713, v713)
																											mBase = m.M
																											v717 = m.ExcPending
																											if v717 != 0 {
																												return
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v604
																												*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																												F_systable_endscan(m, v597)
																												mBase = m.M
																												v724 = m.ExcPending
																												if v724 != 0 {
																													return
																												} else {
																													F_relation_close(m, v583, int32(0))
																													mBase = m.M
																													v727 = m.ExcPending
																													if v727 != 0 {
																														return
																													} else {
																														m.G0 = v579 + int32(288)
																														m.G0 = v15 + int32(16)
																														return
																													}
																												}
																											}
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v604
																											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																											F_systable_endscan(m, v597)
																											mBase = m.M
																											v724 = m.ExcPending
																											if v724 != 0 {
																												return
																											} else {
																												F_relation_close(m, v583, int32(0))
																												mBase = m.M
																												v727 = m.ExcPending
																												if v727 != 0 {
																													return
																												} else {
																													m.G0 = v579 + int32(288)
																													m.G0 = v15 + int32(16)
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
													} else {
														v656 = v599 + int32(4)
														F_LockTuple(m, v583, v656, int32(7))
														mBase = m.M
														v659 = m.ExcPending
														if v659 != 0 {
															return
														} else {
															*(*int64)(unsafe.Add(mBase, uint32(v579)+96)) = base.I64_extend_i32_u(v19)
															v662 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v579)+18)) = uint8(v662)
															v665 = *(*int32)(unsafe.Add(mBase, uint32(v583)+52))
															v668 = F_heap_getattr_6(m, v599, int32(18), v665, v579+int32(15))
															mBase = m.M
															v669 = m.ExcPending
															if v669 != 0 {
																return
															} else {
																v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+15)))
																if v670 == int32(0) {
																	v674 = F_pg_detoast_datum(m, base.I32_wrap_i64(v668))
																	mBase = m.M
																	v675 = m.ExcPending
																	if v675 != 0 {
																		return
																	} else {
																		v676 = *(*int32)(unsafe.Add(mBase, uint32(v603)+68))
																		v677 = F_aclnewowner(m, v674, v676, v19)
																		mBase = m.M
																		v678 = m.ExcPending
																		if v678 != 0 {
																			return
																		} else {
																			v679 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v579)+33)) = uint8(v679)
																			*(*int64)(unsafe.Add(mBase, uint32(v579)+216)) = base.I64_extend_i32_u(v677)
																			v684 = *(*int32)(unsafe.Add(mBase, uint32(v583)+52))
																			v691 = F_heap_modify_tuple(m, v599, v684, v579+int32(80), v579+int32(48), v579+int32(16))
																			mBase = m.M
																			v692 = m.ExcPending
																			if v692 != 0 {
																				return
																			} else {
																				F_CatalogTupleUpdate(m, v583, v691+int32(4), v691)
																				mBase = m.M
																				v696 = m.ExcPending
																				if v696 != 0 {
																					return
																				} else {
																					F_UnlockTuple(m, v583, v656, int32(7))
																					mBase = m.M
																					v699 = m.ExcPending
																					if v699 != 0 {
																						return
																					} else {
																						F_pfree(m, v691)
																						mBase = m.M
																						v701 = m.ExcPending
																						if v701 != 0 {
																							return
																						} else {
																							F_changeDependencyOnOwner(m, int32(1262), v604, v19)
																							mBase = m.M
																							v704 = m.ExcPending
																							if v704 != 0 {
																								return
																							} else {
																								v711 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[1]))
																								if v711 != 0 {
																									v713 = int32(0)
																									F_RunObjectPostAlterHook(m, int32(1262), v604, v713, v713, v713)
																									mBase = m.M
																									v717 = m.ExcPending
																									if v717 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v604
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																										F_systable_endscan(m, v597)
																										mBase = m.M
																										v724 = m.ExcPending
																										if v724 != 0 {
																											return
																										} else {
																											F_relation_close(m, v583, int32(0))
																											mBase = m.M
																											v727 = m.ExcPending
																											if v727 != 0 {
																												return
																											} else {
																												m.G0 = v579 + int32(288)
																												m.G0 = v15 + int32(16)
																												return
																											}
																										}
																									}
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v604
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																									F_systable_endscan(m, v597)
																									mBase = m.M
																									v724 = m.ExcPending
																									if v724 != 0 {
																										return
																									} else {
																										F_relation_close(m, v583, int32(0))
																										mBase = m.M
																										v727 = m.ExcPending
																										if v727 != 0 {
																											return
																										} else {
																											m.G0 = v579 + int32(288)
																											m.G0 = v15 + int32(16)
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
																	v684 = *(*int32)(unsafe.Add(mBase, uint32(v583)+52))
																	v691 = F_heap_modify_tuple(m, v599, v684, v579+int32(80), v579+int32(48), v579+int32(16))
																	mBase = m.M
																	v692 = m.ExcPending
																	if v692 != 0 {
																		return
																	} else {
																		F_CatalogTupleUpdate(m, v583, v691+int32(4), v691)
																		mBase = m.M
																		v696 = m.ExcPending
																		if v696 != 0 {
																			return
																		} else {
																			F_UnlockTuple(m, v583, v656, int32(7))
																			mBase = m.M
																			v699 = m.ExcPending
																			if v699 != 0 {
																				return
																			} else {
																				F_pfree(m, v691)
																				mBase = m.M
																				v701 = m.ExcPending
																				if v701 != 0 {
																					return
																				} else {
																					F_changeDependencyOnOwner(m, int32(1262), v604, v19)
																					mBase = m.M
																					v704 = m.ExcPending
																					if v704 != 0 {
																						return
																					} else {
																						v711 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[1]))
																						if v711 != 0 {
																							v713 = int32(0)
																							F_RunObjectPostAlterHook(m, int32(1262), v604, v713, v713, v713)
																							mBase = m.M
																							v717 = m.ExcPending
																							if v717 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v604
																								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																								F_systable_endscan(m, v597)
																								mBase = m.M
																								v724 = m.ExcPending
																								if v724 != 0 {
																									return
																								} else {
																									F_relation_close(m, v583, int32(0))
																									mBase = m.M
																									v727 = m.ExcPending
																									if v727 != 0 {
																										return
																									} else {
																										m.G0 = v579 + int32(288)
																										m.G0 = v15 + int32(16)
																										return
																									}
																								}
																							}
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v604
																							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																							F_systable_endscan(m, v597)
																							mBase = m.M
																							v724 = m.ExcPending
																							if v724 != 0 {
																								return
																							} else {
																								F_relation_close(m, v583, int32(0))
																								mBase = m.M
																								v727 = m.ExcPending
																								if v727 != 0 {
																									return
																								} else {
																									m.G0 = v579 + int32(288)
																									m.G0 = v15 + int32(16)
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
								} else {
									v711 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[1]))
									if v711 != 0 {
										v713 = int32(0)
										F_RunObjectPostAlterHook(m, int32(1262), v604, v713, v713, v713)
										mBase = m.M
										v717 = m.ExcPending
										if v717 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v604
											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
											F_systable_endscan(m, v597)
											mBase = m.M
											v724 = m.ExcPending
											if v724 != 0 {
												return
											} else {
												F_relation_close(m, v583, int32(0))
												mBase = m.M
												v727 = m.ExcPending
												if v727 != 0 {
													return
												} else {
													m.G0 = v579 + int32(288)
													m.G0 = v15 + int32(16)
													return
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v604
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
										F_systable_endscan(m, v597)
										mBase = m.M
										v724 = m.ExcPending
										if v724 != 0 {
											return
										} else {
											F_relation_close(m, v583, int32(0))
											mBase = m.M
											v727 = m.ExcPending
											if v727 != 0 {
												return
											} else {
												m.G0 = v579 + int32(288)
												m.G0 = v15 + int32(16)
												return
											}
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v734 = m.ExcPending
								if v734 != 0 {
									return
								} else {
									F_errcode(m, int32(1283))
									mBase = m.M
									v737 = m.ExcPending
									if v737 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v579))) = v576
										F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_6), v579)
										mBase = m.M
										v741 = m.ExcPending
										if v741 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_4), int32(2708), int32(_a_F_ExecAlterOwnerStmt_5))
											mBase = m.M
											v746 = m.ExcPending
											if v746 != 0 {
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
		case 11, 49:
			v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v76 = m.G0
			v78 = v76 - int32(128)
			m.G0 = v78
			v82 = F_table_open(m, int32(1247), int32(3))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return
			} else {
				v85 = F_makeTypeNameFromNameList(m, v75)
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					v88 = F_LookupTypeName(m, int32(0), v85, int32(0))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return
					} else {
						if v88 != 0 {
							v90 = F_typeTypeId(m, v88)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return
							} else {
								v92 = F_heap_copytuple(m, v88)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return
								} else {
									F_ReleaseCatCache(m, v88)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return
									} else {
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
										v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+22)))
										v98 = v96 + v97
										v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+79)))
										if v21 == int32(12) {
											if v99 == int32(100) {
												v131 = *(*int32)(unsafe.Add(mBase, uint32(v98)+92))
												if v131 != 0 {
													v132 = *(*int32)(unsafe.Add(mBase, uint32(v98)+88))
													if v132 == int32(_a_F_ExecAlterOwnerStmt_7) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v264 = m.ExcPending
														if v264 != 0 {
															return
														} else {
															F_errcode(m, int32(151027844))
															mBase = m.M
															v267 = m.ExcPending
															if v267 != 0 {
																return
															} else {
																v268 = F_format_type_be(m, v90)
																mBase = m.M
																v269 = m.ExcPending
																if v269 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v78)+64)) = v268
																	F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_8), v78-int32(-64))
																	mBase = m.M
																	v275 = m.ExcPending
																	if v275 != 0 {
																		return
																	} else {
																		v276 = *(*int32)(unsafe.Add(mBase, uint32(v98)+92))
																		v277 = F_format_type_be(m, v276)
																		mBase = m.M
																		v278 = m.ExcPending
																		if v278 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v78)+48)) = v277
																			F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_9), v78+int32(48))
																			mBase = m.M
																			v284 = m.ExcPending
																			if v284 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3949), int32(_a_F_ExecAlterOwnerStmt_11))
																				mBase = m.M
																				v289 = m.ExcPending
																				if v289 != 0 {
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
													} else {
														v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+79)))
														if v135 == int32(109) {
															v138 = F_get_multirange_range(m, v90)
															mBase = m.M
															v139 = m.ExcPending
															if v139 != 0 {
																return
															} else {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v143 = m.ExcPending
																if v143 != 0 {
																	return
																} else {
																	F_errcode(m, int32(151027844))
																	mBase = m.M
																	v146 = m.ExcPending
																	if v146 != 0 {
																		return
																	} else {
																		v147 = F_format_type_be(m, v90)
																		mBase = m.M
																		v148 = m.ExcPending
																		if v148 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = v147
																			F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_12), v78+int32(32))
																			mBase = m.M
																			v154 = m.ExcPending
																			if v154 != 0 {
																				return
																			} else {
																				if v138 != 0 {
																					v155 = F_format_type_be(m, v138)
																					mBase = m.M
																					v156 = m.ExcPending
																					if v156 != 0 {
																						return
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v155
																						F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_13), v78+int32(16))
																						mBase = m.M
																						v162 = m.ExcPending
																						if v162 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3963), int32(_a_F_ExecAlterOwnerStmt_11))
																							mBase = m.M
																							v167 = m.ExcPending
																							if v167 != 0 {
																								return
																							} else {
																								base.Wasm_trap_unreachable()
																								for {
																								}
																							}
																						}
																					}
																				} else {
																					F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3963), int32(_a_F_ExecAlterOwnerStmt_11))
																					mBase = m.M
																					v167 = m.ExcPending
																					if v167 != 0 {
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
														} else {
															v168 = *(*int32)(unsafe.Add(mBase, uint32(v98)+72))
															if v19 != v168 {
																v170 = F_superuser(m)
																mBase = m.M
																v171 = m.ExcPending
																if v171 != 0 {
																	return
																} else {
																	if v170 != 0 {
																		F_AlterTypeOwner_oid(m, v90, v19)
																		mBase = m.M
																		v203 = m.ExcPending
																		if v203 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																			F_relation_close(m, v82, int32(3))
																			mBase = m.M
																			v212 = m.ExcPending
																			if v212 != 0 {
																				return
																			} else {
																				m.G0 = v78 + int32(128)
																				m.G0 = v15 + int32(16)
																				return
																			}
																		}
																	} else {
																		v173 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
																		v175 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																		v176 = F_object_ownercheck(m, int32(1247), v173, v175)
																		mBase = m.M
																		v177 = m.ExcPending
																		if v177 != 0 {
																			return
																		} else {
																			if v176 == int32(0) {
																				v181 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
																				F_aclcheck_error_type(m, int32(2), v181)
																				mBase = m.M
																				v183 = m.ExcPending
																				if v183 != 0 {
																					return
																				} else {
																					v185 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																					F_check_can_set_role(m, v185, v19)
																					mBase = m.M
																					v187 = m.ExcPending
																					if v187 != 0 {
																						return
																					} else {
																						v189 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																						v191 = F_object_aclcheck(m, int32(2615), v189, v19, int64(512))
																						mBase = m.M
																						v192 = m.ExcPending
																						if v192 != 0 {
																							return
																						} else {
																							if v191 == int32(0) {
																								F_AlterTypeOwner_oid(m, v90, v19)
																								mBase = m.M
																								v203 = m.ExcPending
																								if v203 != 0 {
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																									F_relation_close(m, v82, int32(3))
																									mBase = m.M
																									v212 = m.ExcPending
																									if v212 != 0 {
																										return
																									} else {
																										m.G0 = v78 + int32(128)
																										m.G0 = v15 + int32(16)
																										return
																									}
																								}
																							} else {
																								v196 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																								v197 = F_get_namespace_name(m, v196)
																								mBase = m.M
																								v198 = m.ExcPending
																								if v198 != 0 {
																									return
																								} else {
																									F_aclcheck_error(m, v191, int32(37), v197)
																									mBase = m.M
																									v200 = m.ExcPending
																									if v200 != 0 {
																										return
																									} else {
																										F_AlterTypeOwner_oid(m, v90, v19)
																										mBase = m.M
																										v203 = m.ExcPending
																										if v203 != 0 {
																											return
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																											F_relation_close(m, v82, int32(3))
																											mBase = m.M
																											v212 = m.ExcPending
																											if v212 != 0 {
																												return
																											} else {
																												m.G0 = v78 + int32(128)
																												m.G0 = v15 + int32(16)
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
																				v185 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																				F_check_can_set_role(m, v185, v19)
																				mBase = m.M
																				v187 = m.ExcPending
																				if v187 != 0 {
																					return
																				} else {
																					v189 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																					v191 = F_object_aclcheck(m, int32(2615), v189, v19, int64(512))
																					mBase = m.M
																					v192 = m.ExcPending
																					if v192 != 0 {
																						return
																					} else {
																						if v191 == int32(0) {
																							F_AlterTypeOwner_oid(m, v90, v19)
																							mBase = m.M
																							v203 = m.ExcPending
																							if v203 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																								F_relation_close(m, v82, int32(3))
																								mBase = m.M
																								v212 = m.ExcPending
																								if v212 != 0 {
																									return
																								} else {
																									m.G0 = v78 + int32(128)
																									m.G0 = v15 + int32(16)
																									return
																								}
																							}
																						} else {
																							v196 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																							v197 = F_get_namespace_name(m, v196)
																							mBase = m.M
																							v198 = m.ExcPending
																							if v198 != 0 {
																								return
																							} else {
																								F_aclcheck_error(m, v191, int32(37), v197)
																								mBase = m.M
																								v200 = m.ExcPending
																								if v200 != 0 {
																									return
																								} else {
																									F_AlterTypeOwner_oid(m, v90, v19)
																									mBase = m.M
																									v203 = m.ExcPending
																									if v203 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																										F_relation_close(m, v82, int32(3))
																										mBase = m.M
																										v212 = m.ExcPending
																										if v212 != 0 {
																											return
																										} else {
																											m.G0 = v78 + int32(128)
																											m.G0 = v15 + int32(16)
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
																*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																F_relation_close(m, v82, int32(3))
																mBase = m.M
																v212 = m.ExcPending
																if v212 != 0 {
																	return
																} else {
																	m.G0 = v78 + int32(128)
																	m.G0 = v15 + int32(16)
																	return
																}
															}
														}
													}
												} else {
													v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+79)))
													if v135 == int32(109) {
														v138 = F_get_multirange_range(m, v90)
														mBase = m.M
														v139 = m.ExcPending
														if v139 != 0 {
															return
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v143 = m.ExcPending
															if v143 != 0 {
																return
															} else {
																F_errcode(m, int32(151027844))
																mBase = m.M
																v146 = m.ExcPending
																if v146 != 0 {
																	return
																} else {
																	v147 = F_format_type_be(m, v90)
																	mBase = m.M
																	v148 = m.ExcPending
																	if v148 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = v147
																		F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_12), v78+int32(32))
																		mBase = m.M
																		v154 = m.ExcPending
																		if v154 != 0 {
																			return
																		} else {
																			if v138 != 0 {
																				v155 = F_format_type_be(m, v138)
																				mBase = m.M
																				v156 = m.ExcPending
																				if v156 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v155
																					F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_13), v78+int32(16))
																					mBase = m.M
																					v162 = m.ExcPending
																					if v162 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3963), int32(_a_F_ExecAlterOwnerStmt_11))
																						mBase = m.M
																						v167 = m.ExcPending
																						if v167 != 0 {
																							return
																						} else {
																							base.Wasm_trap_unreachable()
																							for {
																							}
																						}
																					}
																				}
																			} else {
																				F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3963), int32(_a_F_ExecAlterOwnerStmt_11))
																				mBase = m.M
																				v167 = m.ExcPending
																				if v167 != 0 {
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
													} else {
														v168 = *(*int32)(unsafe.Add(mBase, uint32(v98)+72))
														if v19 != v168 {
															v170 = F_superuser(m)
															mBase = m.M
															v171 = m.ExcPending
															if v171 != 0 {
																return
															} else {
																if v170 != 0 {
																	F_AlterTypeOwner_oid(m, v90, v19)
																	mBase = m.M
																	v203 = m.ExcPending
																	if v203 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																		F_relation_close(m, v82, int32(3))
																		mBase = m.M
																		v212 = m.ExcPending
																		if v212 != 0 {
																			return
																		} else {
																			m.G0 = v78 + int32(128)
																			m.G0 = v15 + int32(16)
																			return
																		}
																	}
																} else {
																	v173 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
																	v175 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																	v176 = F_object_ownercheck(m, int32(1247), v173, v175)
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return
																	} else {
																		if v176 == int32(0) {
																			v181 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
																			F_aclcheck_error_type(m, int32(2), v181)
																			mBase = m.M
																			v183 = m.ExcPending
																			if v183 != 0 {
																				return
																			} else {
																				v185 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																				F_check_can_set_role(m, v185, v19)
																				mBase = m.M
																				v187 = m.ExcPending
																				if v187 != 0 {
																					return
																				} else {
																					v189 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																					v191 = F_object_aclcheck(m, int32(2615), v189, v19, int64(512))
																					mBase = m.M
																					v192 = m.ExcPending
																					if v192 != 0 {
																						return
																					} else {
																						if v191 == int32(0) {
																							F_AlterTypeOwner_oid(m, v90, v19)
																							mBase = m.M
																							v203 = m.ExcPending
																							if v203 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																								F_relation_close(m, v82, int32(3))
																								mBase = m.M
																								v212 = m.ExcPending
																								if v212 != 0 {
																									return
																								} else {
																									m.G0 = v78 + int32(128)
																									m.G0 = v15 + int32(16)
																									return
																								}
																							}
																						} else {
																							v196 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																							v197 = F_get_namespace_name(m, v196)
																							mBase = m.M
																							v198 = m.ExcPending
																							if v198 != 0 {
																								return
																							} else {
																								F_aclcheck_error(m, v191, int32(37), v197)
																								mBase = m.M
																								v200 = m.ExcPending
																								if v200 != 0 {
																									return
																								} else {
																									F_AlterTypeOwner_oid(m, v90, v19)
																									mBase = m.M
																									v203 = m.ExcPending
																									if v203 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																										F_relation_close(m, v82, int32(3))
																										mBase = m.M
																										v212 = m.ExcPending
																										if v212 != 0 {
																											return
																										} else {
																											m.G0 = v78 + int32(128)
																											m.G0 = v15 + int32(16)
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
																			v185 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																			F_check_can_set_role(m, v185, v19)
																			mBase = m.M
																			v187 = m.ExcPending
																			if v187 != 0 {
																				return
																			} else {
																				v189 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																				v191 = F_object_aclcheck(m, int32(2615), v189, v19, int64(512))
																				mBase = m.M
																				v192 = m.ExcPending
																				if v192 != 0 {
																					return
																				} else {
																					if v191 == int32(0) {
																						F_AlterTypeOwner_oid(m, v90, v19)
																						mBase = m.M
																						v203 = m.ExcPending
																						if v203 != 0 {
																							return
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																							F_relation_close(m, v82, int32(3))
																							mBase = m.M
																							v212 = m.ExcPending
																							if v212 != 0 {
																								return
																							} else {
																								m.G0 = v78 + int32(128)
																								m.G0 = v15 + int32(16)
																								return
																							}
																						}
																					} else {
																						v196 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																						v197 = F_get_namespace_name(m, v196)
																						mBase = m.M
																						v198 = m.ExcPending
																						if v198 != 0 {
																							return
																						} else {
																							F_aclcheck_error(m, v191, int32(37), v197)
																							mBase = m.M
																							v200 = m.ExcPending
																							if v200 != 0 {
																								return
																							} else {
																								F_AlterTypeOwner_oid(m, v90, v19)
																								mBase = m.M
																								v203 = m.ExcPending
																								if v203 != 0 {
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																									F_relation_close(m, v82, int32(3))
																									mBase = m.M
																									v212 = m.ExcPending
																									if v212 != 0 {
																										return
																									} else {
																										m.G0 = v78 + int32(128)
																										m.G0 = v15 + int32(16)
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
															*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
															*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
															F_relation_close(m, v82, int32(3))
															mBase = m.M
															v212 = m.ExcPending
															if v212 != 0 {
																return
															} else {
																m.G0 = v78 + int32(128)
																m.G0 = v15 + int32(16)
																return
															}
														}
													}
												}
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v107 = m.ExcPending
												if v107 != 0 {
													return
												} else {
													F_errcode(m, int32(151027844))
													mBase = m.M
													v110 = m.ExcPending
													if v110 != 0 {
														return
													} else {
														v111 = F_format_type_be(m, v90)
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v78)+80)) = v111
															F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_14), v78+int32(80))
															mBase = m.M
															v118 = m.ExcPending
															if v118 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3925), int32(_a_F_ExecAlterOwnerStmt_11))
																mBase = m.M
																v123 = m.ExcPending
																if v123 != 0 {
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
										} else {
											if v99 != int32(99) {
												v131 = *(*int32)(unsafe.Add(mBase, uint32(v98)+92))
												if v131 != 0 {
													v132 = *(*int32)(unsafe.Add(mBase, uint32(v98)+88))
													if v132 == int32(_a_F_ExecAlterOwnerStmt_7) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v264 = m.ExcPending
														if v264 != 0 {
															return
														} else {
															F_errcode(m, int32(151027844))
															mBase = m.M
															v267 = m.ExcPending
															if v267 != 0 {
																return
															} else {
																v268 = F_format_type_be(m, v90)
																mBase = m.M
																v269 = m.ExcPending
																if v269 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v78)+64)) = v268
																	F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_8), v78-int32(-64))
																	mBase = m.M
																	v275 = m.ExcPending
																	if v275 != 0 {
																		return
																	} else {
																		v276 = *(*int32)(unsafe.Add(mBase, uint32(v98)+92))
																		v277 = F_format_type_be(m, v276)
																		mBase = m.M
																		v278 = m.ExcPending
																		if v278 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v78)+48)) = v277
																			F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_9), v78+int32(48))
																			mBase = m.M
																			v284 = m.ExcPending
																			if v284 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3949), int32(_a_F_ExecAlterOwnerStmt_11))
																				mBase = m.M
																				v289 = m.ExcPending
																				if v289 != 0 {
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
													} else {
														v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+79)))
														if v135 == int32(109) {
															v138 = F_get_multirange_range(m, v90)
															mBase = m.M
															v139 = m.ExcPending
															if v139 != 0 {
																return
															} else {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v143 = m.ExcPending
																if v143 != 0 {
																	return
																} else {
																	F_errcode(m, int32(151027844))
																	mBase = m.M
																	v146 = m.ExcPending
																	if v146 != 0 {
																		return
																	} else {
																		v147 = F_format_type_be(m, v90)
																		mBase = m.M
																		v148 = m.ExcPending
																		if v148 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = v147
																			F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_12), v78+int32(32))
																			mBase = m.M
																			v154 = m.ExcPending
																			if v154 != 0 {
																				return
																			} else {
																				if v138 != 0 {
																					v155 = F_format_type_be(m, v138)
																					mBase = m.M
																					v156 = m.ExcPending
																					if v156 != 0 {
																						return
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v155
																						F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_13), v78+int32(16))
																						mBase = m.M
																						v162 = m.ExcPending
																						if v162 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3963), int32(_a_F_ExecAlterOwnerStmt_11))
																							mBase = m.M
																							v167 = m.ExcPending
																							if v167 != 0 {
																								return
																							} else {
																								base.Wasm_trap_unreachable()
																								for {
																								}
																							}
																						}
																					}
																				} else {
																					F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3963), int32(_a_F_ExecAlterOwnerStmt_11))
																					mBase = m.M
																					v167 = m.ExcPending
																					if v167 != 0 {
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
														} else {
															v168 = *(*int32)(unsafe.Add(mBase, uint32(v98)+72))
															if v19 != v168 {
																v170 = F_superuser(m)
																mBase = m.M
																v171 = m.ExcPending
																if v171 != 0 {
																	return
																} else {
																	if v170 != 0 {
																		F_AlterTypeOwner_oid(m, v90, v19)
																		mBase = m.M
																		v203 = m.ExcPending
																		if v203 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																			F_relation_close(m, v82, int32(3))
																			mBase = m.M
																			v212 = m.ExcPending
																			if v212 != 0 {
																				return
																			} else {
																				m.G0 = v78 + int32(128)
																				m.G0 = v15 + int32(16)
																				return
																			}
																		}
																	} else {
																		v173 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
																		v175 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																		v176 = F_object_ownercheck(m, int32(1247), v173, v175)
																		mBase = m.M
																		v177 = m.ExcPending
																		if v177 != 0 {
																			return
																		} else {
																			if v176 == int32(0) {
																				v181 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
																				F_aclcheck_error_type(m, int32(2), v181)
																				mBase = m.M
																				v183 = m.ExcPending
																				if v183 != 0 {
																					return
																				} else {
																					v185 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																					F_check_can_set_role(m, v185, v19)
																					mBase = m.M
																					v187 = m.ExcPending
																					if v187 != 0 {
																						return
																					} else {
																						v189 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																						v191 = F_object_aclcheck(m, int32(2615), v189, v19, int64(512))
																						mBase = m.M
																						v192 = m.ExcPending
																						if v192 != 0 {
																							return
																						} else {
																							if v191 == int32(0) {
																								F_AlterTypeOwner_oid(m, v90, v19)
																								mBase = m.M
																								v203 = m.ExcPending
																								if v203 != 0 {
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																									F_relation_close(m, v82, int32(3))
																									mBase = m.M
																									v212 = m.ExcPending
																									if v212 != 0 {
																										return
																									} else {
																										m.G0 = v78 + int32(128)
																										m.G0 = v15 + int32(16)
																										return
																									}
																								}
																							} else {
																								v196 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																								v197 = F_get_namespace_name(m, v196)
																								mBase = m.M
																								v198 = m.ExcPending
																								if v198 != 0 {
																									return
																								} else {
																									F_aclcheck_error(m, v191, int32(37), v197)
																									mBase = m.M
																									v200 = m.ExcPending
																									if v200 != 0 {
																										return
																									} else {
																										F_AlterTypeOwner_oid(m, v90, v19)
																										mBase = m.M
																										v203 = m.ExcPending
																										if v203 != 0 {
																											return
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																											F_relation_close(m, v82, int32(3))
																											mBase = m.M
																											v212 = m.ExcPending
																											if v212 != 0 {
																												return
																											} else {
																												m.G0 = v78 + int32(128)
																												m.G0 = v15 + int32(16)
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
																				v185 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																				F_check_can_set_role(m, v185, v19)
																				mBase = m.M
																				v187 = m.ExcPending
																				if v187 != 0 {
																					return
																				} else {
																					v189 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																					v191 = F_object_aclcheck(m, int32(2615), v189, v19, int64(512))
																					mBase = m.M
																					v192 = m.ExcPending
																					if v192 != 0 {
																						return
																					} else {
																						if v191 == int32(0) {
																							F_AlterTypeOwner_oid(m, v90, v19)
																							mBase = m.M
																							v203 = m.ExcPending
																							if v203 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																								F_relation_close(m, v82, int32(3))
																								mBase = m.M
																								v212 = m.ExcPending
																								if v212 != 0 {
																									return
																								} else {
																									m.G0 = v78 + int32(128)
																									m.G0 = v15 + int32(16)
																									return
																								}
																							}
																						} else {
																							v196 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																							v197 = F_get_namespace_name(m, v196)
																							mBase = m.M
																							v198 = m.ExcPending
																							if v198 != 0 {
																								return
																							} else {
																								F_aclcheck_error(m, v191, int32(37), v197)
																								mBase = m.M
																								v200 = m.ExcPending
																								if v200 != 0 {
																									return
																								} else {
																									F_AlterTypeOwner_oid(m, v90, v19)
																									mBase = m.M
																									v203 = m.ExcPending
																									if v203 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																										F_relation_close(m, v82, int32(3))
																										mBase = m.M
																										v212 = m.ExcPending
																										if v212 != 0 {
																											return
																										} else {
																											m.G0 = v78 + int32(128)
																											m.G0 = v15 + int32(16)
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
																*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																F_relation_close(m, v82, int32(3))
																mBase = m.M
																v212 = m.ExcPending
																if v212 != 0 {
																	return
																} else {
																	m.G0 = v78 + int32(128)
																	m.G0 = v15 + int32(16)
																	return
																}
															}
														}
													}
												} else {
													v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+79)))
													if v135 == int32(109) {
														v138 = F_get_multirange_range(m, v90)
														mBase = m.M
														v139 = m.ExcPending
														if v139 != 0 {
															return
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v143 = m.ExcPending
															if v143 != 0 {
																return
															} else {
																F_errcode(m, int32(151027844))
																mBase = m.M
																v146 = m.ExcPending
																if v146 != 0 {
																	return
																} else {
																	v147 = F_format_type_be(m, v90)
																	mBase = m.M
																	v148 = m.ExcPending
																	if v148 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = v147
																		F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_12), v78+int32(32))
																		mBase = m.M
																		v154 = m.ExcPending
																		if v154 != 0 {
																			return
																		} else {
																			if v138 != 0 {
																				v155 = F_format_type_be(m, v138)
																				mBase = m.M
																				v156 = m.ExcPending
																				if v156 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v155
																					F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_13), v78+int32(16))
																					mBase = m.M
																					v162 = m.ExcPending
																					if v162 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3963), int32(_a_F_ExecAlterOwnerStmt_11))
																						mBase = m.M
																						v167 = m.ExcPending
																						if v167 != 0 {
																							return
																						} else {
																							base.Wasm_trap_unreachable()
																							for {
																							}
																						}
																					}
																				}
																			} else {
																				F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3963), int32(_a_F_ExecAlterOwnerStmt_11))
																				mBase = m.M
																				v167 = m.ExcPending
																				if v167 != 0 {
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
													} else {
														v168 = *(*int32)(unsafe.Add(mBase, uint32(v98)+72))
														if v19 != v168 {
															v170 = F_superuser(m)
															mBase = m.M
															v171 = m.ExcPending
															if v171 != 0 {
																return
															} else {
																if v170 != 0 {
																	F_AlterTypeOwner_oid(m, v90, v19)
																	mBase = m.M
																	v203 = m.ExcPending
																	if v203 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																		F_relation_close(m, v82, int32(3))
																		mBase = m.M
																		v212 = m.ExcPending
																		if v212 != 0 {
																			return
																		} else {
																			m.G0 = v78 + int32(128)
																			m.G0 = v15 + int32(16)
																			return
																		}
																	}
																} else {
																	v173 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
																	v175 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																	v176 = F_object_ownercheck(m, int32(1247), v173, v175)
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return
																	} else {
																		if v176 == int32(0) {
																			v181 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
																			F_aclcheck_error_type(m, int32(2), v181)
																			mBase = m.M
																			v183 = m.ExcPending
																			if v183 != 0 {
																				return
																			} else {
																				v185 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																				F_check_can_set_role(m, v185, v19)
																				mBase = m.M
																				v187 = m.ExcPending
																				if v187 != 0 {
																					return
																				} else {
																					v189 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																					v191 = F_object_aclcheck(m, int32(2615), v189, v19, int64(512))
																					mBase = m.M
																					v192 = m.ExcPending
																					if v192 != 0 {
																						return
																					} else {
																						if v191 == int32(0) {
																							F_AlterTypeOwner_oid(m, v90, v19)
																							mBase = m.M
																							v203 = m.ExcPending
																							if v203 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																								F_relation_close(m, v82, int32(3))
																								mBase = m.M
																								v212 = m.ExcPending
																								if v212 != 0 {
																									return
																								} else {
																									m.G0 = v78 + int32(128)
																									m.G0 = v15 + int32(16)
																									return
																								}
																							}
																						} else {
																							v196 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																							v197 = F_get_namespace_name(m, v196)
																							mBase = m.M
																							v198 = m.ExcPending
																							if v198 != 0 {
																								return
																							} else {
																								F_aclcheck_error(m, v191, int32(37), v197)
																								mBase = m.M
																								v200 = m.ExcPending
																								if v200 != 0 {
																									return
																								} else {
																									F_AlterTypeOwner_oid(m, v90, v19)
																									mBase = m.M
																									v203 = m.ExcPending
																									if v203 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																										F_relation_close(m, v82, int32(3))
																										mBase = m.M
																										v212 = m.ExcPending
																										if v212 != 0 {
																											return
																										} else {
																											m.G0 = v78 + int32(128)
																											m.G0 = v15 + int32(16)
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
																			v185 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																			F_check_can_set_role(m, v185, v19)
																			mBase = m.M
																			v187 = m.ExcPending
																			if v187 != 0 {
																				return
																			} else {
																				v189 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																				v191 = F_object_aclcheck(m, int32(2615), v189, v19, int64(512))
																				mBase = m.M
																				v192 = m.ExcPending
																				if v192 != 0 {
																					return
																				} else {
																					if v191 == int32(0) {
																						F_AlterTypeOwner_oid(m, v90, v19)
																						mBase = m.M
																						v203 = m.ExcPending
																						if v203 != 0 {
																							return
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																							F_relation_close(m, v82, int32(3))
																							mBase = m.M
																							v212 = m.ExcPending
																							if v212 != 0 {
																								return
																							} else {
																								m.G0 = v78 + int32(128)
																								m.G0 = v15 + int32(16)
																								return
																							}
																						}
																					} else {
																						v196 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																						v197 = F_get_namespace_name(m, v196)
																						mBase = m.M
																						v198 = m.ExcPending
																						if v198 != 0 {
																							return
																						} else {
																							F_aclcheck_error(m, v191, int32(37), v197)
																							mBase = m.M
																							v200 = m.ExcPending
																							if v200 != 0 {
																								return
																							} else {
																								F_AlterTypeOwner_oid(m, v90, v19)
																								mBase = m.M
																								v203 = m.ExcPending
																								if v203 != 0 {
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																									F_relation_close(m, v82, int32(3))
																									mBase = m.M
																									v212 = m.ExcPending
																									if v212 != 0 {
																										return
																									} else {
																										m.G0 = v78 + int32(128)
																										m.G0 = v15 + int32(16)
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
															*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
															*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
															F_relation_close(m, v82, int32(3))
															mBase = m.M
															v212 = m.ExcPending
															if v212 != 0 {
																return
															} else {
																m.G0 = v78 + int32(128)
																m.G0 = v15 + int32(16)
																return
															}
														}
													}
												}
											} else {
												v126 = *(*int32)(unsafe.Add(mBase, uint32(v98)+84))
												v127 = F_get_rel_relkind(m, v126)
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return
												} else {
													if v127 != int32(99) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v237 = m.ExcPending
														if v237 != 0 {
															return
														} else {
															F_errcode(m, int32(151027844))
															mBase = m.M
															v240 = m.ExcPending
															if v240 != 0 {
																return
															} else {
																v241 = F_format_type_be(m, v90)
																mBase = m.M
																v242 = m.ExcPending
																if v242 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v78)+112)) = v241
																	F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_15), v78+int32(112))
																	mBase = m.M
																	v248 = m.ExcPending
																	if v248 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v78)+96)) = int32(_a_F_ExecAlterOwnerStmt_16)
																		F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_17), v78+int32(96))
																		mBase = m.M
																		v255 = m.ExcPending
																		if v255 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3940), int32(_a_F_ExecAlterOwnerStmt_11))
																			mBase = m.M
																			v260 = m.ExcPending
																			if v260 != 0 {
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
													} else {
														v131 = *(*int32)(unsafe.Add(mBase, uint32(v98)+92))
														if v131 != 0 {
															v132 = *(*int32)(unsafe.Add(mBase, uint32(v98)+88))
															if v132 == int32(_a_F_ExecAlterOwnerStmt_7) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v264 = m.ExcPending
																if v264 != 0 {
																	return
																} else {
																	F_errcode(m, int32(151027844))
																	mBase = m.M
																	v267 = m.ExcPending
																	if v267 != 0 {
																		return
																	} else {
																		v268 = F_format_type_be(m, v90)
																		mBase = m.M
																		v269 = m.ExcPending
																		if v269 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v78)+64)) = v268
																			F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_8), v78-int32(-64))
																			mBase = m.M
																			v275 = m.ExcPending
																			if v275 != 0 {
																				return
																			} else {
																				v276 = *(*int32)(unsafe.Add(mBase, uint32(v98)+92))
																				v277 = F_format_type_be(m, v276)
																				mBase = m.M
																				v278 = m.ExcPending
																				if v278 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v78)+48)) = v277
																					F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_9), v78+int32(48))
																					mBase = m.M
																					v284 = m.ExcPending
																					if v284 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3949), int32(_a_F_ExecAlterOwnerStmt_11))
																						mBase = m.M
																						v289 = m.ExcPending
																						if v289 != 0 {
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
															} else {
																v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+79)))
																if v135 == int32(109) {
																	v138 = F_get_multirange_range(m, v90)
																	mBase = m.M
																	v139 = m.ExcPending
																	if v139 != 0 {
																		return
																	} else {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v143 = m.ExcPending
																		if v143 != 0 {
																			return
																		} else {
																			F_errcode(m, int32(151027844))
																			mBase = m.M
																			v146 = m.ExcPending
																			if v146 != 0 {
																				return
																			} else {
																				v147 = F_format_type_be(m, v90)
																				mBase = m.M
																				v148 = m.ExcPending
																				if v148 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = v147
																					F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_12), v78+int32(32))
																					mBase = m.M
																					v154 = m.ExcPending
																					if v154 != 0 {
																						return
																					} else {
																						if v138 != 0 {
																							v155 = F_format_type_be(m, v138)
																							mBase = m.M
																							v156 = m.ExcPending
																							if v156 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v155
																								F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_13), v78+int32(16))
																								mBase = m.M
																								v162 = m.ExcPending
																								if v162 != 0 {
																									return
																								} else {
																									F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3963), int32(_a_F_ExecAlterOwnerStmt_11))
																									mBase = m.M
																									v167 = m.ExcPending
																									if v167 != 0 {
																										return
																									} else {
																										base.Wasm_trap_unreachable()
																										for {
																										}
																									}
																								}
																							}
																						} else {
																							F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3963), int32(_a_F_ExecAlterOwnerStmt_11))
																							mBase = m.M
																							v167 = m.ExcPending
																							if v167 != 0 {
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
																} else {
																	v168 = *(*int32)(unsafe.Add(mBase, uint32(v98)+72))
																	if v19 != v168 {
																		v170 = F_superuser(m)
																		mBase = m.M
																		v171 = m.ExcPending
																		if v171 != 0 {
																			return
																		} else {
																			if v170 != 0 {
																				F_AlterTypeOwner_oid(m, v90, v19)
																				mBase = m.M
																				v203 = m.ExcPending
																				if v203 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																					*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																					F_relation_close(m, v82, int32(3))
																					mBase = m.M
																					v212 = m.ExcPending
																					if v212 != 0 {
																						return
																					} else {
																						m.G0 = v78 + int32(128)
																						m.G0 = v15 + int32(16)
																						return
																					}
																				}
																			} else {
																				v173 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
																				v175 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																				v176 = F_object_ownercheck(m, int32(1247), v173, v175)
																				mBase = m.M
																				v177 = m.ExcPending
																				if v177 != 0 {
																					return
																				} else {
																					if v176 == int32(0) {
																						v181 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
																						F_aclcheck_error_type(m, int32(2), v181)
																						mBase = m.M
																						v183 = m.ExcPending
																						if v183 != 0 {
																							return
																						} else {
																							v185 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																							F_check_can_set_role(m, v185, v19)
																							mBase = m.M
																							v187 = m.ExcPending
																							if v187 != 0 {
																								return
																							} else {
																								v189 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																								v191 = F_object_aclcheck(m, int32(2615), v189, v19, int64(512))
																								mBase = m.M
																								v192 = m.ExcPending
																								if v192 != 0 {
																									return
																								} else {
																									if v191 == int32(0) {
																										F_AlterTypeOwner_oid(m, v90, v19)
																										mBase = m.M
																										v203 = m.ExcPending
																										if v203 != 0 {
																											return
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																											F_relation_close(m, v82, int32(3))
																											mBase = m.M
																											v212 = m.ExcPending
																											if v212 != 0 {
																												return
																											} else {
																												m.G0 = v78 + int32(128)
																												m.G0 = v15 + int32(16)
																												return
																											}
																										}
																									} else {
																										v196 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																										v197 = F_get_namespace_name(m, v196)
																										mBase = m.M
																										v198 = m.ExcPending
																										if v198 != 0 {
																											return
																										} else {
																											F_aclcheck_error(m, v191, int32(37), v197)
																											mBase = m.M
																											v200 = m.ExcPending
																											if v200 != 0 {
																												return
																											} else {
																												F_AlterTypeOwner_oid(m, v90, v19)
																												mBase = m.M
																												v203 = m.ExcPending
																												if v203 != 0 {
																													return
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																													*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																													F_relation_close(m, v82, int32(3))
																													mBase = m.M
																													v212 = m.ExcPending
																													if v212 != 0 {
																														return
																													} else {
																														m.G0 = v78 + int32(128)
																														m.G0 = v15 + int32(16)
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
																						v185 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																						F_check_can_set_role(m, v185, v19)
																						mBase = m.M
																						v187 = m.ExcPending
																						if v187 != 0 {
																							return
																						} else {
																							v189 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																							v191 = F_object_aclcheck(m, int32(2615), v189, v19, int64(512))
																							mBase = m.M
																							v192 = m.ExcPending
																							if v192 != 0 {
																								return
																							} else {
																								if v191 == int32(0) {
																									F_AlterTypeOwner_oid(m, v90, v19)
																									mBase = m.M
																									v203 = m.ExcPending
																									if v203 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																										F_relation_close(m, v82, int32(3))
																										mBase = m.M
																										v212 = m.ExcPending
																										if v212 != 0 {
																											return
																										} else {
																											m.G0 = v78 + int32(128)
																											m.G0 = v15 + int32(16)
																											return
																										}
																									}
																								} else {
																									v196 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																									v197 = F_get_namespace_name(m, v196)
																									mBase = m.M
																									v198 = m.ExcPending
																									if v198 != 0 {
																										return
																									} else {
																										F_aclcheck_error(m, v191, int32(37), v197)
																										mBase = m.M
																										v200 = m.ExcPending
																										if v200 != 0 {
																											return
																										} else {
																											F_AlterTypeOwner_oid(m, v90, v19)
																											mBase = m.M
																											v203 = m.ExcPending
																											if v203 != 0 {
																												return
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																												*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																												F_relation_close(m, v82, int32(3))
																												mBase = m.M
																												v212 = m.ExcPending
																												if v212 != 0 {
																													return
																												} else {
																													m.G0 = v78 + int32(128)
																													m.G0 = v15 + int32(16)
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
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																		F_relation_close(m, v82, int32(3))
																		mBase = m.M
																		v212 = m.ExcPending
																		if v212 != 0 {
																			return
																		} else {
																			m.G0 = v78 + int32(128)
																			m.G0 = v15 + int32(16)
																			return
																		}
																	}
																}
															}
														} else {
															v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+79)))
															if v135 == int32(109) {
																v138 = F_get_multirange_range(m, v90)
																mBase = m.M
																v139 = m.ExcPending
																if v139 != 0 {
																	return
																} else {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v143 = m.ExcPending
																	if v143 != 0 {
																		return
																	} else {
																		F_errcode(m, int32(151027844))
																		mBase = m.M
																		v146 = m.ExcPending
																		if v146 != 0 {
																			return
																		} else {
																			v147 = F_format_type_be(m, v90)
																			mBase = m.M
																			v148 = m.ExcPending
																			if v148 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = v147
																				F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_12), v78+int32(32))
																				mBase = m.M
																				v154 = m.ExcPending
																				if v154 != 0 {
																					return
																				} else {
																					if v138 != 0 {
																						v155 = F_format_type_be(m, v138)
																						mBase = m.M
																						v156 = m.ExcPending
																						if v156 != 0 {
																							return
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v155
																							F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_13), v78+int32(16))
																							mBase = m.M
																							v162 = m.ExcPending
																							if v162 != 0 {
																								return
																							} else {
																								F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3963), int32(_a_F_ExecAlterOwnerStmt_11))
																								mBase = m.M
																								v167 = m.ExcPending
																								if v167 != 0 {
																									return
																								} else {
																									base.Wasm_trap_unreachable()
																									for {
																									}
																								}
																							}
																						}
																					} else {
																						F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3963), int32(_a_F_ExecAlterOwnerStmt_11))
																						mBase = m.M
																						v167 = m.ExcPending
																						if v167 != 0 {
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
															} else {
																v168 = *(*int32)(unsafe.Add(mBase, uint32(v98)+72))
																if v19 != v168 {
																	v170 = F_superuser(m)
																	mBase = m.M
																	v171 = m.ExcPending
																	if v171 != 0 {
																		return
																	} else {
																		if v170 != 0 {
																			F_AlterTypeOwner_oid(m, v90, v19)
																			mBase = m.M
																			v203 = m.ExcPending
																			if v203 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																				F_relation_close(m, v82, int32(3))
																				mBase = m.M
																				v212 = m.ExcPending
																				if v212 != 0 {
																					return
																				} else {
																					m.G0 = v78 + int32(128)
																					m.G0 = v15 + int32(16)
																					return
																				}
																			}
																		} else {
																			v173 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
																			v175 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																			v176 = F_object_ownercheck(m, int32(1247), v173, v175)
																			mBase = m.M
																			v177 = m.ExcPending
																			if v177 != 0 {
																				return
																			} else {
																				if v176 == int32(0) {
																					v181 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
																					F_aclcheck_error_type(m, int32(2), v181)
																					mBase = m.M
																					v183 = m.ExcPending
																					if v183 != 0 {
																						return
																					} else {
																						v185 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																						F_check_can_set_role(m, v185, v19)
																						mBase = m.M
																						v187 = m.ExcPending
																						if v187 != 0 {
																							return
																						} else {
																							v189 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																							v191 = F_object_aclcheck(m, int32(2615), v189, v19, int64(512))
																							mBase = m.M
																							v192 = m.ExcPending
																							if v192 != 0 {
																								return
																							} else {
																								if v191 == int32(0) {
																									F_AlterTypeOwner_oid(m, v90, v19)
																									mBase = m.M
																									v203 = m.ExcPending
																									if v203 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																										F_relation_close(m, v82, int32(3))
																										mBase = m.M
																										v212 = m.ExcPending
																										if v212 != 0 {
																											return
																										} else {
																											m.G0 = v78 + int32(128)
																											m.G0 = v15 + int32(16)
																											return
																										}
																									}
																								} else {
																									v196 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																									v197 = F_get_namespace_name(m, v196)
																									mBase = m.M
																									v198 = m.ExcPending
																									if v198 != 0 {
																										return
																									} else {
																										F_aclcheck_error(m, v191, int32(37), v197)
																										mBase = m.M
																										v200 = m.ExcPending
																										if v200 != 0 {
																											return
																										} else {
																											F_AlterTypeOwner_oid(m, v90, v19)
																											mBase = m.M
																											v203 = m.ExcPending
																											if v203 != 0 {
																												return
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																												*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																												F_relation_close(m, v82, int32(3))
																												mBase = m.M
																												v212 = m.ExcPending
																												if v212 != 0 {
																													return
																												} else {
																													m.G0 = v78 + int32(128)
																													m.G0 = v15 + int32(16)
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
																					v185 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																					F_check_can_set_role(m, v185, v19)
																					mBase = m.M
																					v187 = m.ExcPending
																					if v187 != 0 {
																						return
																					} else {
																						v189 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																						v191 = F_object_aclcheck(m, int32(2615), v189, v19, int64(512))
																						mBase = m.M
																						v192 = m.ExcPending
																						if v192 != 0 {
																							return
																						} else {
																							if v191 == int32(0) {
																								F_AlterTypeOwner_oid(m, v90, v19)
																								mBase = m.M
																								v203 = m.ExcPending
																								if v203 != 0 {
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																									F_relation_close(m, v82, int32(3))
																									mBase = m.M
																									v212 = m.ExcPending
																									if v212 != 0 {
																										return
																									} else {
																										m.G0 = v78 + int32(128)
																										m.G0 = v15 + int32(16)
																										return
																									}
																								}
																							} else {
																								v196 = *(*int32)(unsafe.Add(mBase, uint32(v98)+68))
																								v197 = F_get_namespace_name(m, v196)
																								mBase = m.M
																								v198 = m.ExcPending
																								if v198 != 0 {
																									return
																								} else {
																									F_aclcheck_error(m, v191, int32(37), v197)
																									mBase = m.M
																									v200 = m.ExcPending
																									if v200 != 0 {
																										return
																									} else {
																										F_AlterTypeOwner_oid(m, v90, v19)
																										mBase = m.M
																										v203 = m.ExcPending
																										if v203 != 0 {
																											return
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																											F_relation_close(m, v82, int32(3))
																											mBase = m.M
																											v212 = m.ExcPending
																											if v212 != 0 {
																												return
																											} else {
																												m.G0 = v78 + int32(128)
																												m.G0 = v15 + int32(16)
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
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
																	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																	F_relation_close(m, v82, int32(3))
																	mBase = m.M
																	v212 = m.ExcPending
																	if v212 != 0 {
																		return
																	} else {
																		m.G0 = v78 + int32(128)
																		m.G0 = v15 + int32(16)
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
							v219 = m.ExcPending
							if v219 != 0 {
								return
							} else {
								F_errcode(m, int32(67137668))
								mBase = m.M
								v222 = m.ExcPending
								if v222 != 0 {
									return
								} else {
									v223 = F_TypeNameToString(m, v85)
									mBase = m.M
									v224 = m.ExcPending
									if v224 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v78))) = v223
										F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_18), v78)
										mBase = m.M
										v228 = m.ExcPending
										if v228 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3911), int32(_a_F_ExecAlterOwnerStmt_11))
											mBase = m.M
											v233 = m.ExcPending
											if v233 != 0 {
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
		case 13:
			v394 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
			v396 = m.G0
			v398 = v396 - int32(16)
			m.G0 = v398
			v402 = F_table_open(m, int32(3466), int32(3))
			mBase = m.M
			v403 = m.ExcPending
			if v403 != 0 {
				return
			} else {
				v407 = F_SearchSysCacheCopy(m, int32(25), base.I64_extend_i32_u(v395), int64(0))
				mBase = m.M
				v408 = m.ExcPending
				if v408 != 0 {
					return
				} else {
					if v407 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v414 = m.ExcPending
						if v414 != 0 {
							return
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v417 = m.ExcPending
							if v417 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v398))) = v395
								F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_19), v398)
								mBase = m.M
								v421 = m.ExcPending
								if v421 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_20), int32(496), int32(_a_F_ExecAlterOwnerStmt_21))
									mBase = m.M
									v426 = m.ExcPending
									if v426 != 0 {
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
						v427 = *(*int32)(unsafe.Add(mBase, uint32(v407)+16))
						v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427)+22)))
						v430 = *(*int32)(unsafe.Add(mBase, uint32(v427+v428)))
						F_AlterEventTriggerOwner_internal(m, v402, v407, v19)
						mBase = m.M
						v432 = m.ExcPending
						if v432 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v430
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(3466)
							F_pfree(m, v407)
							mBase = m.M
							v439 = m.ExcPending
							if v439 != 0 {
								return
							} else {
								F_relation_close(m, v402, int32(3))
								mBase = m.M
								v442 = m.ExcPending
								if v442 != 0 {
									return
								} else {
									m.G0 = v398 + int32(16)
									m.G0 = v15 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		case 15:
			v290 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)+4))
			v292 = m.G0
			v294 = v292 - int32(16)
			m.G0 = v294
			v298 = F_table_open(m, int32(2328), int32(3))
			mBase = m.M
			v299 = m.ExcPending
			if v299 != 0 {
				return
			} else {
				v303 = F_SearchSysCacheCopy(m, int32(29), base.I64_extend_i32_u(v291), int64(0))
				mBase = m.M
				v304 = m.ExcPending
				if v304 != 0 {
					return
				} else {
					if v303 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v310 = m.ExcPending
						if v310 != 0 {
							return
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v313 = m.ExcPending
							if v313 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v294))) = v291
								F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_22), v294)
								mBase = m.M
								v317 = m.ExcPending
								if v317 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_23), int32(302), int32(_a_F_ExecAlterOwnerStmt_24))
									mBase = m.M
									v322 = m.ExcPending
									if v322 != 0 {
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
						v323 = *(*int32)(unsafe.Add(mBase, uint32(v303)+16))
						v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323)+22)))
						v326 = *(*int32)(unsafe.Add(mBase, uint32(v323+v324)))
						F_AlterForeignDataWrapperOwner_internal(m, v298, v303, v19)
						mBase = m.M
						v328 = m.ExcPending
						if v328 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v326
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2328)
							F_pfree(m, v303)
							mBase = m.M
							v335 = m.ExcPending
							if v335 != 0 {
								return
							} else {
								F_relation_close(m, v298, int32(3))
								mBase = m.M
								v338 = m.ExcPending
								if v338 != 0 {
									return
								} else {
									m.G0 = v294 + int32(16)
									m.G0 = v15 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		case 16:
			v342 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)+4))
			v344 = m.G0
			v346 = v344 - int32(16)
			m.G0 = v346
			v350 = F_table_open(m, int32(1417), int32(3))
			mBase = m.M
			v351 = m.ExcPending
			if v351 != 0 {
				return
			} else {
				v355 = F_SearchSysCacheCopy(m, int32(31), base.I64_extend_i32_u(v343), int64(0))
				mBase = m.M
				v356 = m.ExcPending
				if v356 != 0 {
					return
				} else {
					if v355 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v362 = m.ExcPending
						if v362 != 0 {
							return
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v365 = m.ExcPending
							if v365 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v346))) = v343
								F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_25), v346)
								mBase = m.M
								v369 = m.ExcPending
								if v369 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_23), int32(441), int32(_a_F_ExecAlterOwnerStmt_26))
									mBase = m.M
									v374 = m.ExcPending
									if v374 != 0 {
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
						v375 = *(*int32)(unsafe.Add(mBase, uint32(v355)+16))
						v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375)+22)))
						v378 = *(*int32)(unsafe.Add(mBase, uint32(v375+v376)))
						F_AlterForeignServerOwner_internal(m, v350, v355, v19)
						mBase = m.M
						v380 = m.ExcPending
						if v380 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v378
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1417)
							F_pfree(m, v355)
							mBase = m.M
							v387 = m.ExcPending
							if v387 != 0 {
								return
							} else {
								F_relation_close(m, v350, int32(3))
								mBase = m.M
								v390 = m.ExcPending
								if v390 != 0 {
									return
								} else {
									m.G0 = v346 + int32(16)
									m.G0 = v15 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		case 29:
			v446 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v447 = *(*int32)(unsafe.Add(mBase, uint32(v446)+4))
			v448 = m.G0
			v450 = v448 - int32(16)
			m.G0 = v450
			v454 = F_table_open(m, int32(_a_F_ExecAlterOwnerStmt_27), int32(3))
			mBase = m.M
			v455 = m.ExcPending
			if v455 != 0 {
				return
			} else {
				v459 = F_SearchSysCacheCopy(m, int32(48), base.I64_extend_i32_u(v447), int64(0))
				mBase = m.M
				v460 = m.ExcPending
				if v460 != 0 {
					return
				} else {
					if v459 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v466 = m.ExcPending
						if v466 != 0 {
							return
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v469 = m.ExcPending
							if v469 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v450))) = v447
								F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_28), v450)
								mBase = m.M
								v473 = m.ExcPending
								if v473 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_29), int32(2250), int32(_a_F_ExecAlterOwnerStmt_30))
									mBase = m.M
									v478 = m.ExcPending
									if v478 != 0 {
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
						v479 = *(*int32)(unsafe.Add(mBase, uint32(v459)+16))
						v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479)+22)))
						v482 = *(*int32)(unsafe.Add(mBase, uint32(v479+v480)))
						F_AlterPublicationOwner_internal(m, v454, v459, v19)
						mBase = m.M
						v484 = m.ExcPending
						if v484 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v482
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a_F_ExecAlterOwnerStmt_27)
							F_pfree(m, v459)
							mBase = m.M
							v491 = m.ExcPending
							if v491 != 0 {
								return
							} else {
								F_relation_close(m, v454, int32(3))
								mBase = m.M
								v494 = m.ExcPending
								if v494 != 0 {
									return
								} else {
									m.G0 = v450 + int32(16)
									m.G0 = v15 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		case 36:
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
			v26 = m.G0
			v28 = v26 - int32(16)
			m.G0 = v28
			v32 = F_table_open(m, int32(2615), int32(3))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				v36 = F_SearchSysCache1(m, int32(37), base.I64_extend_i32_u(v25))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					if v36 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							F_errcode(m, int32(1411))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v28))) = v25
								F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_31), v28)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_32), int32(346), int32(_a_F_ExecAlterOwnerStmt_33))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
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
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
						v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+22)))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v56+v57)))
						F_AlterSchemaOwner_internal(m, v36, v32, v19)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v59
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2615)
							F_ReleaseCatCache(m, v36)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								F_relation_close(m, v32, int32(3))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
								} else {
									m.G0 = v28 + int32(16)
									m.G0 = v15 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		case 38:
			v498 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)+4))
			v500 = m.G0
			v502 = v500 - int32(16)
			m.G0 = v502
			v506 = F_table_open(m, int32(_a_F_ExecAlterOwnerStmt_34), int32(3))
			mBase = m.M
			v507 = m.ExcPending
			if v507 != 0 {
				return
			} else {
				v510 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[2])))
				v512 = F_SearchSysCacheCopy(m, int32(66), v510, base.I64_extend_i32_u(v499))
				mBase = m.M
				v513 = m.ExcPending
				if v513 != 0 {
					return
				} else {
					if v512 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v519 = m.ExcPending
						if v519 != 0 {
							return
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v522 = m.ExcPending
							if v522 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v502))) = v499
								F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_35), v502)
								mBase = m.M
								v526 = m.ExcPending
								if v526 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_36), int32(2850), int32(_a_F_ExecAlterOwnerStmt_37))
									mBase = m.M
									v531 = m.ExcPending
									if v531 != 0 {
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
						v532 = *(*int32)(unsafe.Add(mBase, uint32(v512)+16))
						v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532)+22)))
						v535 = *(*int32)(unsafe.Add(mBase, uint32(v532+v533)))
						F_AlterSubscriptionOwner_internal(m, v506, v512, v19)
						mBase = m.M
						v537 = m.ExcPending
						if v537 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v535
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a_F_ExecAlterOwnerStmt_34)
							F_pfree(m, v512)
							mBase = m.M
							v544 = m.ExcPending
							if v544 != 0 {
								return
							} else {
								F_relation_close(m, v506, int32(3))
								mBase = m.M
								v547 = m.ExcPending
								if v547 != 0 {
									return
								} else {
									m.G0 = v502 + int32(16)
									m.G0 = v15 + int32(16)
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
