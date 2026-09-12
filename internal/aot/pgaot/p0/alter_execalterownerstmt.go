package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecAlterOwnerStmt(m *base.Module, l0 int32, l1 int32) {
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
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
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
	var v66 int32
	_ = v66
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
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
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
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
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
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
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
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
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
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
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
	var v640 int32
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
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v18 = F_get_rolespec_oid(m, v16, int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		switch v20 - int32(1) {
		case 0, 6, 7, 18, 20, 21, 23, 24, 25, 28, 33, 38, 41, 44, 45:
			v548 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v549 = int32(0)
			F_get_object_address(m, l0, v20, v548, v549, int32(8), v549)
			mBase = m.M
			v553 = m.ExcPending
			if v553 != 0 {
				return
			} else {
				v554 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_AlterObjectOwner_internal(m, v554, v555, v18)
				mBase = m.M
				v557 = m.ExcPending
				if v557 != 0 {
					return
				} else {
					m.G0 = v14 + int32(16)
					return
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v561 = m.ExcPending
			if v561 != 0 {
				return
			} else {
				v562 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v562
				F_errmsg_internal(m, int32(503310), v14)
				mBase = m.M
				v566 = m.ExcPending
				if v566 != 0 {
					return
				} else {
					F_errfinish(m, int32(515218), int32(908), int32(103384))
					mBase = m.M
					v571 = m.ExcPending
					if v571 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 8:
			v572 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
			v574 = m.G0
			v576 = v574 - int32(208)
			m.G0 = v576
			v580 = F_table_open(m, int32(1262), int32(3))
			mBase = m.M
			v581 = m.ExcPending
			if v581 != 0 {
				return
			} else {
				F_ScanKeyInit(m, v576+int32(160), int32(2), int32(3), int32(62), v573)
				mBase = m.M
				v588 = m.ExcPending
				if v588 != 0 {
					return
				} else {
					v590 = int32(1)
					v595 = F_systable_beginscan(m, v580, int32(2671), v590, int32(0), v590, v576+int32(160))
					mBase = m.M
					v596 = m.ExcPending
					if v596 != 0 {
						return
					} else {
						v597 = F_systable_getnext(m, v595)
						mBase = m.M
						v598 = m.ExcPending
						if v598 != 0 {
							return
						} else {
							if v597 != 0 {
								v599 = *(*int32)(unsafe.Add(mBase, uint32(v597)+16))
								v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599)+22)))
								v601 = v599 + v600
								v602 = *(*int32)(unsafe.Add(mBase, uint32(v601)))
								v603 = *(*int32)(unsafe.Add(mBase, uint32(v601)+68))
								if v18 != v603 {
									v607 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v576-int32(-64)))) = uint16(v607)
									v609 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v576)+56)) = v609
									*(*int64)(unsafe.Add(mBase, uint32(v576)+48)) = v609
									*(*uint16)(unsafe.Add(mBase, uint32(v576)+32)) = uint16(v607)
									*(*int64)(unsafe.Add(mBase, uint32(v576)+24)) = v609
									*(*int64)(unsafe.Add(mBase, uint32(v576)+16)) = v609
									v621 = *(*int32)(unsafe.Add(mBase, _consts[4]))
									v622 = F_object_ownercheck(m, int32(1262), v602, v621)
									mBase = m.M
									v623 = m.ExcPending
									if v623 != 0 {
										return
									} else {
										if v622 == int32(0) {
											F_aclcheck_error(m, int32(2), int32(9), v573)
											mBase = m.M
											v629 = m.ExcPending
											if v629 != 0 {
												return
											} else {
												v631 = *(*int32)(unsafe.Add(mBase, _consts[4]))
												F_check_can_set_role(m, v631, v18)
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
															v640 = *(*int32)(unsafe.Add(mBase, _consts[4]))
															v641 = F_SearchSysCache1(m, int32(11), v640)
															mBase = m.M
															v642 = m.ExcPending
															if v642 != 0 {
																return
															} else {
																if v641 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v745 = m.ExcPending
																	if v745 != 0 {
																		return
																	} else {
																		F_errcode(m, int32(16797828))
																		mBase = m.M
																		v748 = m.ExcPending
																		if v748 != 0 {
																			return
																		} else {
																			F_errmsg(m, int32(377210), int32(0))
																			mBase = m.M
																			v752 = m.ExcPending
																			if v752 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(514080), int32(2730), int32(228004))
																				mBase = m.M
																				v757 = m.ExcPending
																				if v757 != 0 {
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
																			v745 = m.ExcPending
																			if v745 != 0 {
																				return
																			} else {
																				F_errcode(m, int32(16797828))
																				mBase = m.M
																				v748 = m.ExcPending
																				if v748 != 0 {
																					return
																				} else {
																					F_errmsg(m, int32(377210), int32(0))
																					mBase = m.M
																					v752 = m.ExcPending
																					if v752 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(514080), int32(2730), int32(228004))
																						mBase = m.M
																						v757 = m.ExcPending
																						if v757 != 0 {
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
																			v656 = v597 + int32(4)
																			F_LockTuple(m, v580, v656, int32(7))
																			mBase = m.M
																			v659 = m.ExcPending
																			if v659 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v576)+88)) = v18
																				v661 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v576)+18)) = uint8(v661)
																				v664 = *(*int32)(unsafe.Add(mBase, uint32(v580)+52))
																				v667 = F_heap_getattr_6(m, v597, int32(18), v664, v576+int32(15))
																				mBase = m.M
																				v668 = m.ExcPending
																				if v668 != 0 {
																					return
																				} else {
																					v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+15)))
																					if v669 == int32(0) {
																						v672 = F_pg_detoast_datum(m, v667)
																						mBase = m.M
																						v673 = m.ExcPending
																						if v673 != 0 {
																							return
																						} else {
																							v674 = *(*int32)(unsafe.Add(mBase, uint32(v601)+68))
																							v675 = F_aclnewowner(m, v672, v674, v18)
																							mBase = m.M
																							v676 = m.ExcPending
																							if v676 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v576)+148)) = v675
																								v678 = int32(1)
																								*(*uint8)(unsafe.Add(mBase, uint32(v576)+33)) = uint8(v678)
																								v680 = *(*int32)(unsafe.Add(mBase, uint32(v580)+52))
																								v687 = F_heap_modify_tuple(m, v597, v680, v576+int32(80), v576+int32(48), v576+int32(16))
																								mBase = m.M
																								v688 = m.ExcPending
																								if v688 != 0 {
																									return
																								} else {
																									F_CatalogTupleUpdate(m, v580, v687+int32(4), v687)
																									mBase = m.M
																									v692 = m.ExcPending
																									if v692 != 0 {
																										return
																									} else {
																										F_UnlockTuple(m, v580, v656, int32(7))
																										mBase = m.M
																										v695 = m.ExcPending
																										if v695 != 0 {
																											return
																										} else {
																											F_pfree(m, v687)
																											mBase = m.M
																											v697 = m.ExcPending
																											if v697 != 0 {
																												return
																											} else {
																												F_changeDependencyOnOwner(m, int32(1262), v602, v18)
																												mBase = m.M
																												v700 = m.ExcPending
																												if v700 != 0 {
																													return
																												} else {
																													v705 = *(*int32)(unsafe.Add(mBase, _consts[203]))
																													if v705 != 0 {
																														v707 = int32(0)
																														F_RunObjectPostAlterHook(m, int32(1262), v602, v707, v707, v707)
																														mBase = m.M
																														v711 = m.ExcPending
																														if v711 != 0 {
																															return
																														} else {
																															*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
																															*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																															F_systable_endscan(m, v595)
																															mBase = m.M
																															v718 = m.ExcPending
																															if v718 != 0 {
																																return
																															} else {
																																F_sequence_close(m, v580, int32(0))
																																mBase = m.M
																																v721 = m.ExcPending
																																if v721 != 0 {
																																	return
																																} else {
																																	m.G0 = v576 + int32(208)
																																	m.G0 = v14 + int32(16)
																																	return
																																}
																															}
																														}
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
																														*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																														F_systable_endscan(m, v595)
																														mBase = m.M
																														v718 = m.ExcPending
																														if v718 != 0 {
																															return
																														} else {
																															F_sequence_close(m, v580, int32(0))
																															mBase = m.M
																															v721 = m.ExcPending
																															if v721 != 0 {
																																return
																															} else {
																																m.G0 = v576 + int32(208)
																																m.G0 = v14 + int32(16)
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
																						v680 = *(*int32)(unsafe.Add(mBase, uint32(v580)+52))
																						v687 = F_heap_modify_tuple(m, v597, v680, v576+int32(80), v576+int32(48), v576+int32(16))
																						mBase = m.M
																						v688 = m.ExcPending
																						if v688 != 0 {
																							return
																						} else {
																							F_CatalogTupleUpdate(m, v580, v687+int32(4), v687)
																							mBase = m.M
																							v692 = m.ExcPending
																							if v692 != 0 {
																								return
																							} else {
																								F_UnlockTuple(m, v580, v656, int32(7))
																								mBase = m.M
																								v695 = m.ExcPending
																								if v695 != 0 {
																									return
																								} else {
																									F_pfree(m, v687)
																									mBase = m.M
																									v697 = m.ExcPending
																									if v697 != 0 {
																										return
																									} else {
																										F_changeDependencyOnOwner(m, int32(1262), v602, v18)
																										mBase = m.M
																										v700 = m.ExcPending
																										if v700 != 0 {
																											return
																										} else {
																											v705 = *(*int32)(unsafe.Add(mBase, _consts[203]))
																											if v705 != 0 {
																												v707 = int32(0)
																												F_RunObjectPostAlterHook(m, int32(1262), v602, v707, v707, v707)
																												mBase = m.M
																												v711 = m.ExcPending
																												if v711 != 0 {
																													return
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
																													*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																													F_systable_endscan(m, v595)
																													mBase = m.M
																													v718 = m.ExcPending
																													if v718 != 0 {
																														return
																													} else {
																														F_sequence_close(m, v580, int32(0))
																														mBase = m.M
																														v721 = m.ExcPending
																														if v721 != 0 {
																															return
																														} else {
																															m.G0 = v576 + int32(208)
																															m.G0 = v14 + int32(16)
																															return
																														}
																													}
																												}
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
																												*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																												F_systable_endscan(m, v595)
																												mBase = m.M
																												v718 = m.ExcPending
																												if v718 != 0 {
																													return
																												} else {
																													F_sequence_close(m, v580, int32(0))
																													mBase = m.M
																													v721 = m.ExcPending
																													if v721 != 0 {
																														return
																													} else {
																														m.G0 = v576 + int32(208)
																														m.G0 = v14 + int32(16)
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
															v656 = v597 + int32(4)
															F_LockTuple(m, v580, v656, int32(7))
															mBase = m.M
															v659 = m.ExcPending
															if v659 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v576)+88)) = v18
																v661 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v576)+18)) = uint8(v661)
																v664 = *(*int32)(unsafe.Add(mBase, uint32(v580)+52))
																v667 = F_heap_getattr_6(m, v597, int32(18), v664, v576+int32(15))
																mBase = m.M
																v668 = m.ExcPending
																if v668 != 0 {
																	return
																} else {
																	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+15)))
																	if v669 == int32(0) {
																		v672 = F_pg_detoast_datum(m, v667)
																		mBase = m.M
																		v673 = m.ExcPending
																		if v673 != 0 {
																			return
																		} else {
																			v674 = *(*int32)(unsafe.Add(mBase, uint32(v601)+68))
																			v675 = F_aclnewowner(m, v672, v674, v18)
																			mBase = m.M
																			v676 = m.ExcPending
																			if v676 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v576)+148)) = v675
																				v678 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v576)+33)) = uint8(v678)
																				v680 = *(*int32)(unsafe.Add(mBase, uint32(v580)+52))
																				v687 = F_heap_modify_tuple(m, v597, v680, v576+int32(80), v576+int32(48), v576+int32(16))
																				mBase = m.M
																				v688 = m.ExcPending
																				if v688 != 0 {
																					return
																				} else {
																					F_CatalogTupleUpdate(m, v580, v687+int32(4), v687)
																					mBase = m.M
																					v692 = m.ExcPending
																					if v692 != 0 {
																						return
																					} else {
																						F_UnlockTuple(m, v580, v656, int32(7))
																						mBase = m.M
																						v695 = m.ExcPending
																						if v695 != 0 {
																							return
																						} else {
																							F_pfree(m, v687)
																							mBase = m.M
																							v697 = m.ExcPending
																							if v697 != 0 {
																								return
																							} else {
																								F_changeDependencyOnOwner(m, int32(1262), v602, v18)
																								mBase = m.M
																								v700 = m.ExcPending
																								if v700 != 0 {
																									return
																								} else {
																									v705 = *(*int32)(unsafe.Add(mBase, _consts[203]))
																									if v705 != 0 {
																										v707 = int32(0)
																										F_RunObjectPostAlterHook(m, int32(1262), v602, v707, v707, v707)
																										mBase = m.M
																										v711 = m.ExcPending
																										if v711 != 0 {
																											return
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
																											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																											F_systable_endscan(m, v595)
																											mBase = m.M
																											v718 = m.ExcPending
																											if v718 != 0 {
																												return
																											} else {
																												F_sequence_close(m, v580, int32(0))
																												mBase = m.M
																												v721 = m.ExcPending
																												if v721 != 0 {
																													return
																												} else {
																													m.G0 = v576 + int32(208)
																													m.G0 = v14 + int32(16)
																													return
																												}
																											}
																										}
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																										F_systable_endscan(m, v595)
																										mBase = m.M
																										v718 = m.ExcPending
																										if v718 != 0 {
																											return
																										} else {
																											F_sequence_close(m, v580, int32(0))
																											mBase = m.M
																											v721 = m.ExcPending
																											if v721 != 0 {
																												return
																											} else {
																												m.G0 = v576 + int32(208)
																												m.G0 = v14 + int32(16)
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
																		v680 = *(*int32)(unsafe.Add(mBase, uint32(v580)+52))
																		v687 = F_heap_modify_tuple(m, v597, v680, v576+int32(80), v576+int32(48), v576+int32(16))
																		mBase = m.M
																		v688 = m.ExcPending
																		if v688 != 0 {
																			return
																		} else {
																			F_CatalogTupleUpdate(m, v580, v687+int32(4), v687)
																			mBase = m.M
																			v692 = m.ExcPending
																			if v692 != 0 {
																				return
																			} else {
																				F_UnlockTuple(m, v580, v656, int32(7))
																				mBase = m.M
																				v695 = m.ExcPending
																				if v695 != 0 {
																					return
																				} else {
																					F_pfree(m, v687)
																					mBase = m.M
																					v697 = m.ExcPending
																					if v697 != 0 {
																						return
																					} else {
																						F_changeDependencyOnOwner(m, int32(1262), v602, v18)
																						mBase = m.M
																						v700 = m.ExcPending
																						if v700 != 0 {
																							return
																						} else {
																							v705 = *(*int32)(unsafe.Add(mBase, _consts[203]))
																							if v705 != 0 {
																								v707 = int32(0)
																								F_RunObjectPostAlterHook(m, int32(1262), v602, v707, v707, v707)
																								mBase = m.M
																								v711 = m.ExcPending
																								if v711 != 0 {
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																									F_systable_endscan(m, v595)
																									mBase = m.M
																									v718 = m.ExcPending
																									if v718 != 0 {
																										return
																									} else {
																										F_sequence_close(m, v580, int32(0))
																										mBase = m.M
																										v721 = m.ExcPending
																										if v721 != 0 {
																											return
																										} else {
																											m.G0 = v576 + int32(208)
																											m.G0 = v14 + int32(16)
																											return
																										}
																									}
																								}
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
																								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																								F_systable_endscan(m, v595)
																								mBase = m.M
																								v718 = m.ExcPending
																								if v718 != 0 {
																									return
																								} else {
																									F_sequence_close(m, v580, int32(0))
																									mBase = m.M
																									v721 = m.ExcPending
																									if v721 != 0 {
																										return
																									} else {
																										m.G0 = v576 + int32(208)
																										m.G0 = v14 + int32(16)
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
											v631 = *(*int32)(unsafe.Add(mBase, _consts[4]))
											F_check_can_set_role(m, v631, v18)
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
														v640 = *(*int32)(unsafe.Add(mBase, _consts[4]))
														v641 = F_SearchSysCache1(m, int32(11), v640)
														mBase = m.M
														v642 = m.ExcPending
														if v642 != 0 {
															return
														} else {
															if v641 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v745 = m.ExcPending
																if v745 != 0 {
																	return
																} else {
																	F_errcode(m, int32(16797828))
																	mBase = m.M
																	v748 = m.ExcPending
																	if v748 != 0 {
																		return
																	} else {
																		F_errmsg(m, int32(377210), int32(0))
																		mBase = m.M
																		v752 = m.ExcPending
																		if v752 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(514080), int32(2730), int32(228004))
																			mBase = m.M
																			v757 = m.ExcPending
																			if v757 != 0 {
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
																		v745 = m.ExcPending
																		if v745 != 0 {
																			return
																		} else {
																			F_errcode(m, int32(16797828))
																			mBase = m.M
																			v748 = m.ExcPending
																			if v748 != 0 {
																				return
																			} else {
																				F_errmsg(m, int32(377210), int32(0))
																				mBase = m.M
																				v752 = m.ExcPending
																				if v752 != 0 {
																					return
																				} else {
																					F_errfinish(m, int32(514080), int32(2730), int32(228004))
																					mBase = m.M
																					v757 = m.ExcPending
																					if v757 != 0 {
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
																		v656 = v597 + int32(4)
																		F_LockTuple(m, v580, v656, int32(7))
																		mBase = m.M
																		v659 = m.ExcPending
																		if v659 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v576)+88)) = v18
																			v661 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v576)+18)) = uint8(v661)
																			v664 = *(*int32)(unsafe.Add(mBase, uint32(v580)+52))
																			v667 = F_heap_getattr_6(m, v597, int32(18), v664, v576+int32(15))
																			mBase = m.M
																			v668 = m.ExcPending
																			if v668 != 0 {
																				return
																			} else {
																				v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+15)))
																				if v669 == int32(0) {
																					v672 = F_pg_detoast_datum(m, v667)
																					mBase = m.M
																					v673 = m.ExcPending
																					if v673 != 0 {
																						return
																					} else {
																						v674 = *(*int32)(unsafe.Add(mBase, uint32(v601)+68))
																						v675 = F_aclnewowner(m, v672, v674, v18)
																						mBase = m.M
																						v676 = m.ExcPending
																						if v676 != 0 {
																							return
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v576)+148)) = v675
																							v678 = int32(1)
																							*(*uint8)(unsafe.Add(mBase, uint32(v576)+33)) = uint8(v678)
																							v680 = *(*int32)(unsafe.Add(mBase, uint32(v580)+52))
																							v687 = F_heap_modify_tuple(m, v597, v680, v576+int32(80), v576+int32(48), v576+int32(16))
																							mBase = m.M
																							v688 = m.ExcPending
																							if v688 != 0 {
																								return
																							} else {
																								F_CatalogTupleUpdate(m, v580, v687+int32(4), v687)
																								mBase = m.M
																								v692 = m.ExcPending
																								if v692 != 0 {
																									return
																								} else {
																									F_UnlockTuple(m, v580, v656, int32(7))
																									mBase = m.M
																									v695 = m.ExcPending
																									if v695 != 0 {
																										return
																									} else {
																										F_pfree(m, v687)
																										mBase = m.M
																										v697 = m.ExcPending
																										if v697 != 0 {
																											return
																										} else {
																											F_changeDependencyOnOwner(m, int32(1262), v602, v18)
																											mBase = m.M
																											v700 = m.ExcPending
																											if v700 != 0 {
																												return
																											} else {
																												v705 = *(*int32)(unsafe.Add(mBase, _consts[203]))
																												if v705 != 0 {
																													v707 = int32(0)
																													F_RunObjectPostAlterHook(m, int32(1262), v602, v707, v707, v707)
																													mBase = m.M
																													v711 = m.ExcPending
																													if v711 != 0 {
																														return
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
																														*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																														F_systable_endscan(m, v595)
																														mBase = m.M
																														v718 = m.ExcPending
																														if v718 != 0 {
																															return
																														} else {
																															F_sequence_close(m, v580, int32(0))
																															mBase = m.M
																															v721 = m.ExcPending
																															if v721 != 0 {
																																return
																															} else {
																																m.G0 = v576 + int32(208)
																																m.G0 = v14 + int32(16)
																																return
																															}
																														}
																													}
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
																													*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																													F_systable_endscan(m, v595)
																													mBase = m.M
																													v718 = m.ExcPending
																													if v718 != 0 {
																														return
																													} else {
																														F_sequence_close(m, v580, int32(0))
																														mBase = m.M
																														v721 = m.ExcPending
																														if v721 != 0 {
																															return
																														} else {
																															m.G0 = v576 + int32(208)
																															m.G0 = v14 + int32(16)
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
																					v680 = *(*int32)(unsafe.Add(mBase, uint32(v580)+52))
																					v687 = F_heap_modify_tuple(m, v597, v680, v576+int32(80), v576+int32(48), v576+int32(16))
																					mBase = m.M
																					v688 = m.ExcPending
																					if v688 != 0 {
																						return
																					} else {
																						F_CatalogTupleUpdate(m, v580, v687+int32(4), v687)
																						mBase = m.M
																						v692 = m.ExcPending
																						if v692 != 0 {
																							return
																						} else {
																							F_UnlockTuple(m, v580, v656, int32(7))
																							mBase = m.M
																							v695 = m.ExcPending
																							if v695 != 0 {
																								return
																							} else {
																								F_pfree(m, v687)
																								mBase = m.M
																								v697 = m.ExcPending
																								if v697 != 0 {
																									return
																								} else {
																									F_changeDependencyOnOwner(m, int32(1262), v602, v18)
																									mBase = m.M
																									v700 = m.ExcPending
																									if v700 != 0 {
																										return
																									} else {
																										v705 = *(*int32)(unsafe.Add(mBase, _consts[203]))
																										if v705 != 0 {
																											v707 = int32(0)
																											F_RunObjectPostAlterHook(m, int32(1262), v602, v707, v707, v707)
																											mBase = m.M
																											v711 = m.ExcPending
																											if v711 != 0 {
																												return
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
																												*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																												F_systable_endscan(m, v595)
																												mBase = m.M
																												v718 = m.ExcPending
																												if v718 != 0 {
																													return
																												} else {
																													F_sequence_close(m, v580, int32(0))
																													mBase = m.M
																													v721 = m.ExcPending
																													if v721 != 0 {
																														return
																													} else {
																														m.G0 = v576 + int32(208)
																														m.G0 = v14 + int32(16)
																														return
																													}
																												}
																											}
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
																											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																											F_systable_endscan(m, v595)
																											mBase = m.M
																											v718 = m.ExcPending
																											if v718 != 0 {
																												return
																											} else {
																												F_sequence_close(m, v580, int32(0))
																												mBase = m.M
																												v721 = m.ExcPending
																												if v721 != 0 {
																													return
																												} else {
																													m.G0 = v576 + int32(208)
																													m.G0 = v14 + int32(16)
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
														v656 = v597 + int32(4)
														F_LockTuple(m, v580, v656, int32(7))
														mBase = m.M
														v659 = m.ExcPending
														if v659 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v576)+88)) = v18
															v661 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v576)+18)) = uint8(v661)
															v664 = *(*int32)(unsafe.Add(mBase, uint32(v580)+52))
															v667 = F_heap_getattr_6(m, v597, int32(18), v664, v576+int32(15))
															mBase = m.M
															v668 = m.ExcPending
															if v668 != 0 {
																return
															} else {
																v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+15)))
																if v669 == int32(0) {
																	v672 = F_pg_detoast_datum(m, v667)
																	mBase = m.M
																	v673 = m.ExcPending
																	if v673 != 0 {
																		return
																	} else {
																		v674 = *(*int32)(unsafe.Add(mBase, uint32(v601)+68))
																		v675 = F_aclnewowner(m, v672, v674, v18)
																		mBase = m.M
																		v676 = m.ExcPending
																		if v676 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v576)+148)) = v675
																			v678 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v576)+33)) = uint8(v678)
																			v680 = *(*int32)(unsafe.Add(mBase, uint32(v580)+52))
																			v687 = F_heap_modify_tuple(m, v597, v680, v576+int32(80), v576+int32(48), v576+int32(16))
																			mBase = m.M
																			v688 = m.ExcPending
																			if v688 != 0 {
																				return
																			} else {
																				F_CatalogTupleUpdate(m, v580, v687+int32(4), v687)
																				mBase = m.M
																				v692 = m.ExcPending
																				if v692 != 0 {
																					return
																				} else {
																					F_UnlockTuple(m, v580, v656, int32(7))
																					mBase = m.M
																					v695 = m.ExcPending
																					if v695 != 0 {
																						return
																					} else {
																						F_pfree(m, v687)
																						mBase = m.M
																						v697 = m.ExcPending
																						if v697 != 0 {
																							return
																						} else {
																							F_changeDependencyOnOwner(m, int32(1262), v602, v18)
																							mBase = m.M
																							v700 = m.ExcPending
																							if v700 != 0 {
																								return
																							} else {
																								v705 = *(*int32)(unsafe.Add(mBase, _consts[203]))
																								if v705 != 0 {
																									v707 = int32(0)
																									F_RunObjectPostAlterHook(m, int32(1262), v602, v707, v707, v707)
																									mBase = m.M
																									v711 = m.ExcPending
																									if v711 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																										F_systable_endscan(m, v595)
																										mBase = m.M
																										v718 = m.ExcPending
																										if v718 != 0 {
																											return
																										} else {
																											F_sequence_close(m, v580, int32(0))
																											mBase = m.M
																											v721 = m.ExcPending
																											if v721 != 0 {
																												return
																											} else {
																												m.G0 = v576 + int32(208)
																												m.G0 = v14 + int32(16)
																												return
																											}
																										}
																									}
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																									F_systable_endscan(m, v595)
																									mBase = m.M
																									v718 = m.ExcPending
																									if v718 != 0 {
																										return
																									} else {
																										F_sequence_close(m, v580, int32(0))
																										mBase = m.M
																										v721 = m.ExcPending
																										if v721 != 0 {
																											return
																										} else {
																											m.G0 = v576 + int32(208)
																											m.G0 = v14 + int32(16)
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
																	v680 = *(*int32)(unsafe.Add(mBase, uint32(v580)+52))
																	v687 = F_heap_modify_tuple(m, v597, v680, v576+int32(80), v576+int32(48), v576+int32(16))
																	mBase = m.M
																	v688 = m.ExcPending
																	if v688 != 0 {
																		return
																	} else {
																		F_CatalogTupleUpdate(m, v580, v687+int32(4), v687)
																		mBase = m.M
																		v692 = m.ExcPending
																		if v692 != 0 {
																			return
																		} else {
																			F_UnlockTuple(m, v580, v656, int32(7))
																			mBase = m.M
																			v695 = m.ExcPending
																			if v695 != 0 {
																				return
																			} else {
																				F_pfree(m, v687)
																				mBase = m.M
																				v697 = m.ExcPending
																				if v697 != 0 {
																					return
																				} else {
																					F_changeDependencyOnOwner(m, int32(1262), v602, v18)
																					mBase = m.M
																					v700 = m.ExcPending
																					if v700 != 0 {
																						return
																					} else {
																						v705 = *(*int32)(unsafe.Add(mBase, _consts[203]))
																						if v705 != 0 {
																							v707 = int32(0)
																							F_RunObjectPostAlterHook(m, int32(1262), v602, v707, v707, v707)
																							mBase = m.M
																							v711 = m.ExcPending
																							if v711 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
																								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																								F_systable_endscan(m, v595)
																								mBase = m.M
																								v718 = m.ExcPending
																								if v718 != 0 {
																									return
																								} else {
																									F_sequence_close(m, v580, int32(0))
																									mBase = m.M
																									v721 = m.ExcPending
																									if v721 != 0 {
																										return
																									} else {
																										m.G0 = v576 + int32(208)
																										m.G0 = v14 + int32(16)
																										return
																									}
																								}
																							}
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
																							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																							F_systable_endscan(m, v595)
																							mBase = m.M
																							v718 = m.ExcPending
																							if v718 != 0 {
																								return
																							} else {
																								F_sequence_close(m, v580, int32(0))
																								mBase = m.M
																								v721 = m.ExcPending
																								if v721 != 0 {
																									return
																								} else {
																									m.G0 = v576 + int32(208)
																									m.G0 = v14 + int32(16)
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
									v705 = *(*int32)(unsafe.Add(mBase, _consts[203]))
									if v705 != 0 {
										v707 = int32(0)
										F_RunObjectPostAlterHook(m, int32(1262), v602, v707, v707, v707)
										mBase = m.M
										v711 = m.ExcPending
										if v711 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
											F_systable_endscan(m, v595)
											mBase = m.M
											v718 = m.ExcPending
											if v718 != 0 {
												return
											} else {
												F_sequence_close(m, v580, int32(0))
												mBase = m.M
												v721 = m.ExcPending
												if v721 != 0 {
													return
												} else {
													m.G0 = v576 + int32(208)
													m.G0 = v14 + int32(16)
													return
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
										F_systable_endscan(m, v595)
										mBase = m.M
										v718 = m.ExcPending
										if v718 != 0 {
											return
										} else {
											F_sequence_close(m, v580, int32(0))
											mBase = m.M
											v721 = m.ExcPending
											if v721 != 0 {
												return
											} else {
												m.G0 = v576 + int32(208)
												m.G0 = v14 + int32(16)
												return
											}
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v728 = m.ExcPending
								if v728 != 0 {
									return
								} else {
									F_errcode(m, int32(1283))
									mBase = m.M
									v731 = m.ExcPending
									if v731 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v576))) = v573
										F_errmsg(m, int32(77637), v576)
										mBase = m.M
										v735 = m.ExcPending
										if v735 != 0 {
											return
										} else {
											F_errfinish(m, int32(514080), int32(2690), int32(228004))
											mBase = m.M
											v740 = m.ExcPending
											if v740 != 0 {
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
		case 11, 48:
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v74 = m.G0
			v76 = v74 - int32(128)
			m.G0 = v76
			v80 = F_table_open(m, int32(1247), int32(3))
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return
			} else {
				v83 = F_makeTypeNameFromNameList(m, v73)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return
				} else {
					v86 = F_LookupTypeName(m, int32(0), v83, int32(0))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return
					} else {
						if v86 != 0 {
							v88 = F_typeTypeId(m, v86)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								v90 = F_heap_copytuple(m, v86)
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return
								} else {
									F_ReleaseCatCache(m, v86)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
										v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+22)))
										v96 = v94 + v95
										v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+79)))
										if v20 == int32(12) {
											if v97&int32(255) == int32(100) {
												v133 = *(*int32)(unsafe.Add(mBase, uint32(v96)+92))
												if v133 != 0 {
													v134 = *(*int32)(unsafe.Add(mBase, uint32(v96)+88))
													if v134 == int32(6179) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v266 = m.ExcPending
														if v266 != 0 {
															return
														} else {
															F_errcode(m, int32(151027844))
															mBase = m.M
															v269 = m.ExcPending
															if v269 != 0 {
																return
															} else {
																v270 = F_format_type_be(m, v88)
																mBase = m.M
																v271 = m.ExcPending
																if v271 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v76)+64)) = v270
																	F_errmsg(m, int32(197326), v76-int32(-64))
																	mBase = m.M
																	v277 = m.ExcPending
																	if v277 != 0 {
																		return
																	} else {
																		v278 = *(*int32)(unsafe.Add(mBase, uint32(v96)+92))
																		v279 = F_format_type_be(m, v278)
																		mBase = m.M
																		v280 = m.ExcPending
																		if v280 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v76)+48)) = v279
																			F_errhint(m, int32(645120), v76+int32(48))
																			mBase = m.M
																			v286 = m.ExcPending
																			if v286 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(514319), int32(3881), int32(228023))
																				mBase = m.M
																				v291 = m.ExcPending
																				if v291 != 0 {
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
														v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+79)))
														if v137 == int32(109) {
															v140 = F_get_multirange_range(m, v88)
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return
															} else {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v145 = m.ExcPending
																if v145 != 0 {
																	return
																} else {
																	F_errcode(m, int32(151027844))
																	mBase = m.M
																	v148 = m.ExcPending
																	if v148 != 0 {
																		return
																	} else {
																		v149 = F_format_type_be(m, v88)
																		mBase = m.M
																		v150 = m.ExcPending
																		if v150 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v76)+32)) = v149
																			F_errmsg(m, int32(201809), v76+int32(32))
																			mBase = m.M
																			v156 = m.ExcPending
																			if v156 != 0 {
																				return
																			} else {
																				if v140 != 0 {
																					v157 = F_format_type_be(m, v140)
																					mBase = m.M
																					v158 = m.ExcPending
																					if v158 != 0 {
																						return
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v157
																						F_errhint(m, int32(645184), v76+int32(16))
																						mBase = m.M
																						v164 = m.ExcPending
																						if v164 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(514319), int32(3895), int32(228023))
																							mBase = m.M
																							v169 = m.ExcPending
																							if v169 != 0 {
																								return
																							} else {
																								base.Wasm_trap_unreachable()
																								for {
																								}
																							}
																						}
																					}
																				} else {
																					F_errfinish(m, int32(514319), int32(3895), int32(228023))
																					mBase = m.M
																					v169 = m.ExcPending
																					if v169 != 0 {
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
															v170 = *(*int32)(unsafe.Add(mBase, uint32(v96)+72))
															if v18 != v170 {
																v172 = F_superuser(m)
																mBase = m.M
																v173 = m.ExcPending
																if v173 != 0 {
																	return
																} else {
																	if v172 != 0 {
																		F_AlterTypeOwner_oid(m, v88, v18)
																		mBase = m.M
																		v205 = m.ExcPending
																		if v205 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																			F_sequence_close(m, v80, int32(3))
																			mBase = m.M
																			v214 = m.ExcPending
																			if v214 != 0 {
																				return
																			} else {
																				m.G0 = v76 + int32(128)
																				m.G0 = v14 + int32(16)
																				return
																			}
																		}
																	} else {
																		v175 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																		v177 = *(*int32)(unsafe.Add(mBase, _consts[4]))
																		v178 = F_object_ownercheck(m, int32(1247), v175, v177)
																		mBase = m.M
																		v179 = m.ExcPending
																		if v179 != 0 {
																			return
																		} else {
																			if v178 == int32(0) {
																				v183 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																				F_aclcheck_error_type(m, int32(2), v183)
																				mBase = m.M
																				v185 = m.ExcPending
																				if v185 != 0 {
																					return
																				} else {
																					v187 = *(*int32)(unsafe.Add(mBase, _consts[4]))
																					F_check_can_set_role(m, v187, v18)
																					mBase = m.M
																					v189 = m.ExcPending
																					if v189 != 0 {
																						return
																					} else {
																						v191 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																						v193 = F_object_aclcheck(m, int32(2615), v191, v18, int64(512))
																						mBase = m.M
																						v194 = m.ExcPending
																						if v194 != 0 {
																							return
																						} else {
																							if v193 == int32(0) {
																								F_AlterTypeOwner_oid(m, v88, v18)
																								mBase = m.M
																								v205 = m.ExcPending
																								if v205 != 0 {
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																									F_sequence_close(m, v80, int32(3))
																									mBase = m.M
																									v214 = m.ExcPending
																									if v214 != 0 {
																										return
																									} else {
																										m.G0 = v76 + int32(128)
																										m.G0 = v14 + int32(16)
																										return
																									}
																								}
																							} else {
																								v198 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																								v199 = F_get_namespace_name(m, v198)
																								mBase = m.M
																								v200 = m.ExcPending
																								if v200 != 0 {
																									return
																								} else {
																									F_aclcheck_error(m, v193, int32(36), v199)
																									mBase = m.M
																									v202 = m.ExcPending
																									if v202 != 0 {
																										return
																									} else {
																										F_AlterTypeOwner_oid(m, v88, v18)
																										mBase = m.M
																										v205 = m.ExcPending
																										if v205 != 0 {
																											return
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																											F_sequence_close(m, v80, int32(3))
																											mBase = m.M
																											v214 = m.ExcPending
																											if v214 != 0 {
																												return
																											} else {
																												m.G0 = v76 + int32(128)
																												m.G0 = v14 + int32(16)
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
																				v187 = *(*int32)(unsafe.Add(mBase, _consts[4]))
																				F_check_can_set_role(m, v187, v18)
																				mBase = m.M
																				v189 = m.ExcPending
																				if v189 != 0 {
																					return
																				} else {
																					v191 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																					v193 = F_object_aclcheck(m, int32(2615), v191, v18, int64(512))
																					mBase = m.M
																					v194 = m.ExcPending
																					if v194 != 0 {
																						return
																					} else {
																						if v193 == int32(0) {
																							F_AlterTypeOwner_oid(m, v88, v18)
																							mBase = m.M
																							v205 = m.ExcPending
																							if v205 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																								F_sequence_close(m, v80, int32(3))
																								mBase = m.M
																								v214 = m.ExcPending
																								if v214 != 0 {
																									return
																								} else {
																									m.G0 = v76 + int32(128)
																									m.G0 = v14 + int32(16)
																									return
																								}
																							}
																						} else {
																							v198 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																							v199 = F_get_namespace_name(m, v198)
																							mBase = m.M
																							v200 = m.ExcPending
																							if v200 != 0 {
																								return
																							} else {
																								F_aclcheck_error(m, v193, int32(36), v199)
																								mBase = m.M
																								v202 = m.ExcPending
																								if v202 != 0 {
																									return
																								} else {
																									F_AlterTypeOwner_oid(m, v88, v18)
																									mBase = m.M
																									v205 = m.ExcPending
																									if v205 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																										F_sequence_close(m, v80, int32(3))
																										mBase = m.M
																										v214 = m.ExcPending
																										if v214 != 0 {
																											return
																										} else {
																											m.G0 = v76 + int32(128)
																											m.G0 = v14 + int32(16)
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
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																F_sequence_close(m, v80, int32(3))
																mBase = m.M
																v214 = m.ExcPending
																if v214 != 0 {
																	return
																} else {
																	m.G0 = v76 + int32(128)
																	m.G0 = v14 + int32(16)
																	return
																}
															}
														}
													}
												} else {
													v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+79)))
													if v137 == int32(109) {
														v140 = F_get_multirange_range(m, v88)
														mBase = m.M
														v141 = m.ExcPending
														if v141 != 0 {
															return
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v145 = m.ExcPending
															if v145 != 0 {
																return
															} else {
																F_errcode(m, int32(151027844))
																mBase = m.M
																v148 = m.ExcPending
																if v148 != 0 {
																	return
																} else {
																	v149 = F_format_type_be(m, v88)
																	mBase = m.M
																	v150 = m.ExcPending
																	if v150 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v76)+32)) = v149
																		F_errmsg(m, int32(201809), v76+int32(32))
																		mBase = m.M
																		v156 = m.ExcPending
																		if v156 != 0 {
																			return
																		} else {
																			if v140 != 0 {
																				v157 = F_format_type_be(m, v140)
																				mBase = m.M
																				v158 = m.ExcPending
																				if v158 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v157
																					F_errhint(m, int32(645184), v76+int32(16))
																					mBase = m.M
																					v164 = m.ExcPending
																					if v164 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(514319), int32(3895), int32(228023))
																						mBase = m.M
																						v169 = m.ExcPending
																						if v169 != 0 {
																							return
																						} else {
																							base.Wasm_trap_unreachable()
																							for {
																							}
																						}
																					}
																				}
																			} else {
																				F_errfinish(m, int32(514319), int32(3895), int32(228023))
																				mBase = m.M
																				v169 = m.ExcPending
																				if v169 != 0 {
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
														v170 = *(*int32)(unsafe.Add(mBase, uint32(v96)+72))
														if v18 != v170 {
															v172 = F_superuser(m)
															mBase = m.M
															v173 = m.ExcPending
															if v173 != 0 {
																return
															} else {
																if v172 != 0 {
																	F_AlterTypeOwner_oid(m, v88, v18)
																	mBase = m.M
																	v205 = m.ExcPending
																	if v205 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																		F_sequence_close(m, v80, int32(3))
																		mBase = m.M
																		v214 = m.ExcPending
																		if v214 != 0 {
																			return
																		} else {
																			m.G0 = v76 + int32(128)
																			m.G0 = v14 + int32(16)
																			return
																		}
																	}
																} else {
																	v175 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																	v177 = *(*int32)(unsafe.Add(mBase, _consts[4]))
																	v178 = F_object_ownercheck(m, int32(1247), v175, v177)
																	mBase = m.M
																	v179 = m.ExcPending
																	if v179 != 0 {
																		return
																	} else {
																		if v178 == int32(0) {
																			v183 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																			F_aclcheck_error_type(m, int32(2), v183)
																			mBase = m.M
																			v185 = m.ExcPending
																			if v185 != 0 {
																				return
																			} else {
																				v187 = *(*int32)(unsafe.Add(mBase, _consts[4]))
																				F_check_can_set_role(m, v187, v18)
																				mBase = m.M
																				v189 = m.ExcPending
																				if v189 != 0 {
																					return
																				} else {
																					v191 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																					v193 = F_object_aclcheck(m, int32(2615), v191, v18, int64(512))
																					mBase = m.M
																					v194 = m.ExcPending
																					if v194 != 0 {
																						return
																					} else {
																						if v193 == int32(0) {
																							F_AlterTypeOwner_oid(m, v88, v18)
																							mBase = m.M
																							v205 = m.ExcPending
																							if v205 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																								F_sequence_close(m, v80, int32(3))
																								mBase = m.M
																								v214 = m.ExcPending
																								if v214 != 0 {
																									return
																								} else {
																									m.G0 = v76 + int32(128)
																									m.G0 = v14 + int32(16)
																									return
																								}
																							}
																						} else {
																							v198 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																							v199 = F_get_namespace_name(m, v198)
																							mBase = m.M
																							v200 = m.ExcPending
																							if v200 != 0 {
																								return
																							} else {
																								F_aclcheck_error(m, v193, int32(36), v199)
																								mBase = m.M
																								v202 = m.ExcPending
																								if v202 != 0 {
																									return
																								} else {
																									F_AlterTypeOwner_oid(m, v88, v18)
																									mBase = m.M
																									v205 = m.ExcPending
																									if v205 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																										F_sequence_close(m, v80, int32(3))
																										mBase = m.M
																										v214 = m.ExcPending
																										if v214 != 0 {
																											return
																										} else {
																											m.G0 = v76 + int32(128)
																											m.G0 = v14 + int32(16)
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
																			v187 = *(*int32)(unsafe.Add(mBase, _consts[4]))
																			F_check_can_set_role(m, v187, v18)
																			mBase = m.M
																			v189 = m.ExcPending
																			if v189 != 0 {
																				return
																			} else {
																				v191 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																				v193 = F_object_aclcheck(m, int32(2615), v191, v18, int64(512))
																				mBase = m.M
																				v194 = m.ExcPending
																				if v194 != 0 {
																					return
																				} else {
																					if v193 == int32(0) {
																						F_AlterTypeOwner_oid(m, v88, v18)
																						mBase = m.M
																						v205 = m.ExcPending
																						if v205 != 0 {
																							return
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																							F_sequence_close(m, v80, int32(3))
																							mBase = m.M
																							v214 = m.ExcPending
																							if v214 != 0 {
																								return
																							} else {
																								m.G0 = v76 + int32(128)
																								m.G0 = v14 + int32(16)
																								return
																							}
																						}
																					} else {
																						v198 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																						v199 = F_get_namespace_name(m, v198)
																						mBase = m.M
																						v200 = m.ExcPending
																						if v200 != 0 {
																							return
																						} else {
																							F_aclcheck_error(m, v193, int32(36), v199)
																							mBase = m.M
																							v202 = m.ExcPending
																							if v202 != 0 {
																								return
																							} else {
																								F_AlterTypeOwner_oid(m, v88, v18)
																								mBase = m.M
																								v205 = m.ExcPending
																								if v205 != 0 {
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																									F_sequence_close(m, v80, int32(3))
																									mBase = m.M
																									v214 = m.ExcPending
																									if v214 != 0 {
																										return
																									} else {
																										m.G0 = v76 + int32(128)
																										m.G0 = v14 + int32(16)
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
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
															*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
															F_sequence_close(m, v80, int32(3))
															mBase = m.M
															v214 = m.ExcPending
															if v214 != 0 {
																return
															} else {
																m.G0 = v76 + int32(128)
																m.G0 = v14 + int32(16)
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
														v111 = F_format_type_be(m, v88)
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v76)+80)) = v111
															F_errmsg(m, int32(288868), v76+int32(80))
															mBase = m.M
															v118 = m.ExcPending
															if v118 != 0 {
																return
															} else {
																F_errfinish(m, int32(514319), int32(3857), int32(228023))
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
											if v97&int32(255) != int32(99) {
												v133 = *(*int32)(unsafe.Add(mBase, uint32(v96)+92))
												if v133 != 0 {
													v134 = *(*int32)(unsafe.Add(mBase, uint32(v96)+88))
													if v134 == int32(6179) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v266 = m.ExcPending
														if v266 != 0 {
															return
														} else {
															F_errcode(m, int32(151027844))
															mBase = m.M
															v269 = m.ExcPending
															if v269 != 0 {
																return
															} else {
																v270 = F_format_type_be(m, v88)
																mBase = m.M
																v271 = m.ExcPending
																if v271 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v76)+64)) = v270
																	F_errmsg(m, int32(197326), v76-int32(-64))
																	mBase = m.M
																	v277 = m.ExcPending
																	if v277 != 0 {
																		return
																	} else {
																		v278 = *(*int32)(unsafe.Add(mBase, uint32(v96)+92))
																		v279 = F_format_type_be(m, v278)
																		mBase = m.M
																		v280 = m.ExcPending
																		if v280 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v76)+48)) = v279
																			F_errhint(m, int32(645120), v76+int32(48))
																			mBase = m.M
																			v286 = m.ExcPending
																			if v286 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(514319), int32(3881), int32(228023))
																				mBase = m.M
																				v291 = m.ExcPending
																				if v291 != 0 {
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
														v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+79)))
														if v137 == int32(109) {
															v140 = F_get_multirange_range(m, v88)
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return
															} else {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v145 = m.ExcPending
																if v145 != 0 {
																	return
																} else {
																	F_errcode(m, int32(151027844))
																	mBase = m.M
																	v148 = m.ExcPending
																	if v148 != 0 {
																		return
																	} else {
																		v149 = F_format_type_be(m, v88)
																		mBase = m.M
																		v150 = m.ExcPending
																		if v150 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v76)+32)) = v149
																			F_errmsg(m, int32(201809), v76+int32(32))
																			mBase = m.M
																			v156 = m.ExcPending
																			if v156 != 0 {
																				return
																			} else {
																				if v140 != 0 {
																					v157 = F_format_type_be(m, v140)
																					mBase = m.M
																					v158 = m.ExcPending
																					if v158 != 0 {
																						return
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v157
																						F_errhint(m, int32(645184), v76+int32(16))
																						mBase = m.M
																						v164 = m.ExcPending
																						if v164 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(514319), int32(3895), int32(228023))
																							mBase = m.M
																							v169 = m.ExcPending
																							if v169 != 0 {
																								return
																							} else {
																								base.Wasm_trap_unreachable()
																								for {
																								}
																							}
																						}
																					}
																				} else {
																					F_errfinish(m, int32(514319), int32(3895), int32(228023))
																					mBase = m.M
																					v169 = m.ExcPending
																					if v169 != 0 {
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
															v170 = *(*int32)(unsafe.Add(mBase, uint32(v96)+72))
															if v18 != v170 {
																v172 = F_superuser(m)
																mBase = m.M
																v173 = m.ExcPending
																if v173 != 0 {
																	return
																} else {
																	if v172 != 0 {
																		F_AlterTypeOwner_oid(m, v88, v18)
																		mBase = m.M
																		v205 = m.ExcPending
																		if v205 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																			F_sequence_close(m, v80, int32(3))
																			mBase = m.M
																			v214 = m.ExcPending
																			if v214 != 0 {
																				return
																			} else {
																				m.G0 = v76 + int32(128)
																				m.G0 = v14 + int32(16)
																				return
																			}
																		}
																	} else {
																		v175 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																		v177 = *(*int32)(unsafe.Add(mBase, _consts[4]))
																		v178 = F_object_ownercheck(m, int32(1247), v175, v177)
																		mBase = m.M
																		v179 = m.ExcPending
																		if v179 != 0 {
																			return
																		} else {
																			if v178 == int32(0) {
																				v183 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																				F_aclcheck_error_type(m, int32(2), v183)
																				mBase = m.M
																				v185 = m.ExcPending
																				if v185 != 0 {
																					return
																				} else {
																					v187 = *(*int32)(unsafe.Add(mBase, _consts[4]))
																					F_check_can_set_role(m, v187, v18)
																					mBase = m.M
																					v189 = m.ExcPending
																					if v189 != 0 {
																						return
																					} else {
																						v191 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																						v193 = F_object_aclcheck(m, int32(2615), v191, v18, int64(512))
																						mBase = m.M
																						v194 = m.ExcPending
																						if v194 != 0 {
																							return
																						} else {
																							if v193 == int32(0) {
																								F_AlterTypeOwner_oid(m, v88, v18)
																								mBase = m.M
																								v205 = m.ExcPending
																								if v205 != 0 {
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																									F_sequence_close(m, v80, int32(3))
																									mBase = m.M
																									v214 = m.ExcPending
																									if v214 != 0 {
																										return
																									} else {
																										m.G0 = v76 + int32(128)
																										m.G0 = v14 + int32(16)
																										return
																									}
																								}
																							} else {
																								v198 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																								v199 = F_get_namespace_name(m, v198)
																								mBase = m.M
																								v200 = m.ExcPending
																								if v200 != 0 {
																									return
																								} else {
																									F_aclcheck_error(m, v193, int32(36), v199)
																									mBase = m.M
																									v202 = m.ExcPending
																									if v202 != 0 {
																										return
																									} else {
																										F_AlterTypeOwner_oid(m, v88, v18)
																										mBase = m.M
																										v205 = m.ExcPending
																										if v205 != 0 {
																											return
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																											F_sequence_close(m, v80, int32(3))
																											mBase = m.M
																											v214 = m.ExcPending
																											if v214 != 0 {
																												return
																											} else {
																												m.G0 = v76 + int32(128)
																												m.G0 = v14 + int32(16)
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
																				v187 = *(*int32)(unsafe.Add(mBase, _consts[4]))
																				F_check_can_set_role(m, v187, v18)
																				mBase = m.M
																				v189 = m.ExcPending
																				if v189 != 0 {
																					return
																				} else {
																					v191 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																					v193 = F_object_aclcheck(m, int32(2615), v191, v18, int64(512))
																					mBase = m.M
																					v194 = m.ExcPending
																					if v194 != 0 {
																						return
																					} else {
																						if v193 == int32(0) {
																							F_AlterTypeOwner_oid(m, v88, v18)
																							mBase = m.M
																							v205 = m.ExcPending
																							if v205 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																								F_sequence_close(m, v80, int32(3))
																								mBase = m.M
																								v214 = m.ExcPending
																								if v214 != 0 {
																									return
																								} else {
																									m.G0 = v76 + int32(128)
																									m.G0 = v14 + int32(16)
																									return
																								}
																							}
																						} else {
																							v198 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																							v199 = F_get_namespace_name(m, v198)
																							mBase = m.M
																							v200 = m.ExcPending
																							if v200 != 0 {
																								return
																							} else {
																								F_aclcheck_error(m, v193, int32(36), v199)
																								mBase = m.M
																								v202 = m.ExcPending
																								if v202 != 0 {
																									return
																								} else {
																									F_AlterTypeOwner_oid(m, v88, v18)
																									mBase = m.M
																									v205 = m.ExcPending
																									if v205 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																										F_sequence_close(m, v80, int32(3))
																										mBase = m.M
																										v214 = m.ExcPending
																										if v214 != 0 {
																											return
																										} else {
																											m.G0 = v76 + int32(128)
																											m.G0 = v14 + int32(16)
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
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																F_sequence_close(m, v80, int32(3))
																mBase = m.M
																v214 = m.ExcPending
																if v214 != 0 {
																	return
																} else {
																	m.G0 = v76 + int32(128)
																	m.G0 = v14 + int32(16)
																	return
																}
															}
														}
													}
												} else {
													v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+79)))
													if v137 == int32(109) {
														v140 = F_get_multirange_range(m, v88)
														mBase = m.M
														v141 = m.ExcPending
														if v141 != 0 {
															return
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v145 = m.ExcPending
															if v145 != 0 {
																return
															} else {
																F_errcode(m, int32(151027844))
																mBase = m.M
																v148 = m.ExcPending
																if v148 != 0 {
																	return
																} else {
																	v149 = F_format_type_be(m, v88)
																	mBase = m.M
																	v150 = m.ExcPending
																	if v150 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v76)+32)) = v149
																		F_errmsg(m, int32(201809), v76+int32(32))
																		mBase = m.M
																		v156 = m.ExcPending
																		if v156 != 0 {
																			return
																		} else {
																			if v140 != 0 {
																				v157 = F_format_type_be(m, v140)
																				mBase = m.M
																				v158 = m.ExcPending
																				if v158 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v157
																					F_errhint(m, int32(645184), v76+int32(16))
																					mBase = m.M
																					v164 = m.ExcPending
																					if v164 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(514319), int32(3895), int32(228023))
																						mBase = m.M
																						v169 = m.ExcPending
																						if v169 != 0 {
																							return
																						} else {
																							base.Wasm_trap_unreachable()
																							for {
																							}
																						}
																					}
																				}
																			} else {
																				F_errfinish(m, int32(514319), int32(3895), int32(228023))
																				mBase = m.M
																				v169 = m.ExcPending
																				if v169 != 0 {
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
														v170 = *(*int32)(unsafe.Add(mBase, uint32(v96)+72))
														if v18 != v170 {
															v172 = F_superuser(m)
															mBase = m.M
															v173 = m.ExcPending
															if v173 != 0 {
																return
															} else {
																if v172 != 0 {
																	F_AlterTypeOwner_oid(m, v88, v18)
																	mBase = m.M
																	v205 = m.ExcPending
																	if v205 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																		F_sequence_close(m, v80, int32(3))
																		mBase = m.M
																		v214 = m.ExcPending
																		if v214 != 0 {
																			return
																		} else {
																			m.G0 = v76 + int32(128)
																			m.G0 = v14 + int32(16)
																			return
																		}
																	}
																} else {
																	v175 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																	v177 = *(*int32)(unsafe.Add(mBase, _consts[4]))
																	v178 = F_object_ownercheck(m, int32(1247), v175, v177)
																	mBase = m.M
																	v179 = m.ExcPending
																	if v179 != 0 {
																		return
																	} else {
																		if v178 == int32(0) {
																			v183 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																			F_aclcheck_error_type(m, int32(2), v183)
																			mBase = m.M
																			v185 = m.ExcPending
																			if v185 != 0 {
																				return
																			} else {
																				v187 = *(*int32)(unsafe.Add(mBase, _consts[4]))
																				F_check_can_set_role(m, v187, v18)
																				mBase = m.M
																				v189 = m.ExcPending
																				if v189 != 0 {
																					return
																				} else {
																					v191 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																					v193 = F_object_aclcheck(m, int32(2615), v191, v18, int64(512))
																					mBase = m.M
																					v194 = m.ExcPending
																					if v194 != 0 {
																						return
																					} else {
																						if v193 == int32(0) {
																							F_AlterTypeOwner_oid(m, v88, v18)
																							mBase = m.M
																							v205 = m.ExcPending
																							if v205 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																								F_sequence_close(m, v80, int32(3))
																								mBase = m.M
																								v214 = m.ExcPending
																								if v214 != 0 {
																									return
																								} else {
																									m.G0 = v76 + int32(128)
																									m.G0 = v14 + int32(16)
																									return
																								}
																							}
																						} else {
																							v198 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																							v199 = F_get_namespace_name(m, v198)
																							mBase = m.M
																							v200 = m.ExcPending
																							if v200 != 0 {
																								return
																							} else {
																								F_aclcheck_error(m, v193, int32(36), v199)
																								mBase = m.M
																								v202 = m.ExcPending
																								if v202 != 0 {
																									return
																								} else {
																									F_AlterTypeOwner_oid(m, v88, v18)
																									mBase = m.M
																									v205 = m.ExcPending
																									if v205 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																										F_sequence_close(m, v80, int32(3))
																										mBase = m.M
																										v214 = m.ExcPending
																										if v214 != 0 {
																											return
																										} else {
																											m.G0 = v76 + int32(128)
																											m.G0 = v14 + int32(16)
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
																			v187 = *(*int32)(unsafe.Add(mBase, _consts[4]))
																			F_check_can_set_role(m, v187, v18)
																			mBase = m.M
																			v189 = m.ExcPending
																			if v189 != 0 {
																				return
																			} else {
																				v191 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																				v193 = F_object_aclcheck(m, int32(2615), v191, v18, int64(512))
																				mBase = m.M
																				v194 = m.ExcPending
																				if v194 != 0 {
																					return
																				} else {
																					if v193 == int32(0) {
																						F_AlterTypeOwner_oid(m, v88, v18)
																						mBase = m.M
																						v205 = m.ExcPending
																						if v205 != 0 {
																							return
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																							F_sequence_close(m, v80, int32(3))
																							mBase = m.M
																							v214 = m.ExcPending
																							if v214 != 0 {
																								return
																							} else {
																								m.G0 = v76 + int32(128)
																								m.G0 = v14 + int32(16)
																								return
																							}
																						}
																					} else {
																						v198 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																						v199 = F_get_namespace_name(m, v198)
																						mBase = m.M
																						v200 = m.ExcPending
																						if v200 != 0 {
																							return
																						} else {
																							F_aclcheck_error(m, v193, int32(36), v199)
																							mBase = m.M
																							v202 = m.ExcPending
																							if v202 != 0 {
																								return
																							} else {
																								F_AlterTypeOwner_oid(m, v88, v18)
																								mBase = m.M
																								v205 = m.ExcPending
																								if v205 != 0 {
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																									F_sequence_close(m, v80, int32(3))
																									mBase = m.M
																									v214 = m.ExcPending
																									if v214 != 0 {
																										return
																									} else {
																										m.G0 = v76 + int32(128)
																										m.G0 = v14 + int32(16)
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
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
															*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
															F_sequence_close(m, v80, int32(3))
															mBase = m.M
															v214 = m.ExcPending
															if v214 != 0 {
																return
															} else {
																m.G0 = v76 + int32(128)
																m.G0 = v14 + int32(16)
																return
															}
														}
													}
												}
											} else {
												v128 = *(*int32)(unsafe.Add(mBase, uint32(v96)+84))
												v129 = F_get_rel_relkind(m, v128)
												mBase = m.M
												v130 = m.ExcPending
												if v130 != 0 {
													return
												} else {
													if v129 != int32(99) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v239 = m.ExcPending
														if v239 != 0 {
															return
														} else {
															F_errcode(m, int32(151027844))
															mBase = m.M
															v242 = m.ExcPending
															if v242 != 0 {
																return
															} else {
																v243 = F_format_type_be(m, v88)
																mBase = m.M
																v244 = m.ExcPending
																if v244 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = v243
																	F_errmsg(m, int32(382543), v76+int32(112))
																	mBase = m.M
																	v250 = m.ExcPending
																	if v250 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v76)+96)) = int32(562014)
																		F_errhint(m, int32(675236), v76+int32(96))
																		mBase = m.M
																		v257 = m.ExcPending
																		if v257 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(514319), int32(3872), int32(228023))
																			mBase = m.M
																			v262 = m.ExcPending
																			if v262 != 0 {
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
														v133 = *(*int32)(unsafe.Add(mBase, uint32(v96)+92))
														if v133 != 0 {
															v134 = *(*int32)(unsafe.Add(mBase, uint32(v96)+88))
															if v134 == int32(6179) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v266 = m.ExcPending
																if v266 != 0 {
																	return
																} else {
																	F_errcode(m, int32(151027844))
																	mBase = m.M
																	v269 = m.ExcPending
																	if v269 != 0 {
																		return
																	} else {
																		v270 = F_format_type_be(m, v88)
																		mBase = m.M
																		v271 = m.ExcPending
																		if v271 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v76)+64)) = v270
																			F_errmsg(m, int32(197326), v76-int32(-64))
																			mBase = m.M
																			v277 = m.ExcPending
																			if v277 != 0 {
																				return
																			} else {
																				v278 = *(*int32)(unsafe.Add(mBase, uint32(v96)+92))
																				v279 = F_format_type_be(m, v278)
																				mBase = m.M
																				v280 = m.ExcPending
																				if v280 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v76)+48)) = v279
																					F_errhint(m, int32(645120), v76+int32(48))
																					mBase = m.M
																					v286 = m.ExcPending
																					if v286 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(514319), int32(3881), int32(228023))
																						mBase = m.M
																						v291 = m.ExcPending
																						if v291 != 0 {
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
																v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+79)))
																if v137 == int32(109) {
																	v140 = F_get_multirange_range(m, v88)
																	mBase = m.M
																	v141 = m.ExcPending
																	if v141 != 0 {
																		return
																	} else {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v145 = m.ExcPending
																		if v145 != 0 {
																			return
																		} else {
																			F_errcode(m, int32(151027844))
																			mBase = m.M
																			v148 = m.ExcPending
																			if v148 != 0 {
																				return
																			} else {
																				v149 = F_format_type_be(m, v88)
																				mBase = m.M
																				v150 = m.ExcPending
																				if v150 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v76)+32)) = v149
																					F_errmsg(m, int32(201809), v76+int32(32))
																					mBase = m.M
																					v156 = m.ExcPending
																					if v156 != 0 {
																						return
																					} else {
																						if v140 != 0 {
																							v157 = F_format_type_be(m, v140)
																							mBase = m.M
																							v158 = m.ExcPending
																							if v158 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v157
																								F_errhint(m, int32(645184), v76+int32(16))
																								mBase = m.M
																								v164 = m.ExcPending
																								if v164 != 0 {
																									return
																								} else {
																									F_errfinish(m, int32(514319), int32(3895), int32(228023))
																									mBase = m.M
																									v169 = m.ExcPending
																									if v169 != 0 {
																										return
																									} else {
																										base.Wasm_trap_unreachable()
																										for {
																										}
																									}
																								}
																							}
																						} else {
																							F_errfinish(m, int32(514319), int32(3895), int32(228023))
																							mBase = m.M
																							v169 = m.ExcPending
																							if v169 != 0 {
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
																	v170 = *(*int32)(unsafe.Add(mBase, uint32(v96)+72))
																	if v18 != v170 {
																		v172 = F_superuser(m)
																		mBase = m.M
																		v173 = m.ExcPending
																		if v173 != 0 {
																			return
																		} else {
																			if v172 != 0 {
																				F_AlterTypeOwner_oid(m, v88, v18)
																				mBase = m.M
																				v205 = m.ExcPending
																				if v205 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																					*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																					F_sequence_close(m, v80, int32(3))
																					mBase = m.M
																					v214 = m.ExcPending
																					if v214 != 0 {
																						return
																					} else {
																						m.G0 = v76 + int32(128)
																						m.G0 = v14 + int32(16)
																						return
																					}
																				}
																			} else {
																				v175 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																				v177 = *(*int32)(unsafe.Add(mBase, _consts[4]))
																				v178 = F_object_ownercheck(m, int32(1247), v175, v177)
																				mBase = m.M
																				v179 = m.ExcPending
																				if v179 != 0 {
																					return
																				} else {
																					if v178 == int32(0) {
																						v183 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																						F_aclcheck_error_type(m, int32(2), v183)
																						mBase = m.M
																						v185 = m.ExcPending
																						if v185 != 0 {
																							return
																						} else {
																							v187 = *(*int32)(unsafe.Add(mBase, _consts[4]))
																							F_check_can_set_role(m, v187, v18)
																							mBase = m.M
																							v189 = m.ExcPending
																							if v189 != 0 {
																								return
																							} else {
																								v191 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																								v193 = F_object_aclcheck(m, int32(2615), v191, v18, int64(512))
																								mBase = m.M
																								v194 = m.ExcPending
																								if v194 != 0 {
																									return
																								} else {
																									if v193 == int32(0) {
																										F_AlterTypeOwner_oid(m, v88, v18)
																										mBase = m.M
																										v205 = m.ExcPending
																										if v205 != 0 {
																											return
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																											F_sequence_close(m, v80, int32(3))
																											mBase = m.M
																											v214 = m.ExcPending
																											if v214 != 0 {
																												return
																											} else {
																												m.G0 = v76 + int32(128)
																												m.G0 = v14 + int32(16)
																												return
																											}
																										}
																									} else {
																										v198 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																										v199 = F_get_namespace_name(m, v198)
																										mBase = m.M
																										v200 = m.ExcPending
																										if v200 != 0 {
																											return
																										} else {
																											F_aclcheck_error(m, v193, int32(36), v199)
																											mBase = m.M
																											v202 = m.ExcPending
																											if v202 != 0 {
																												return
																											} else {
																												F_AlterTypeOwner_oid(m, v88, v18)
																												mBase = m.M
																												v205 = m.ExcPending
																												if v205 != 0 {
																													return
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																													*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																													F_sequence_close(m, v80, int32(3))
																													mBase = m.M
																													v214 = m.ExcPending
																													if v214 != 0 {
																														return
																													} else {
																														m.G0 = v76 + int32(128)
																														m.G0 = v14 + int32(16)
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
																						v187 = *(*int32)(unsafe.Add(mBase, _consts[4]))
																						F_check_can_set_role(m, v187, v18)
																						mBase = m.M
																						v189 = m.ExcPending
																						if v189 != 0 {
																							return
																						} else {
																							v191 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																							v193 = F_object_aclcheck(m, int32(2615), v191, v18, int64(512))
																							mBase = m.M
																							v194 = m.ExcPending
																							if v194 != 0 {
																								return
																							} else {
																								if v193 == int32(0) {
																									F_AlterTypeOwner_oid(m, v88, v18)
																									mBase = m.M
																									v205 = m.ExcPending
																									if v205 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																										F_sequence_close(m, v80, int32(3))
																										mBase = m.M
																										v214 = m.ExcPending
																										if v214 != 0 {
																											return
																										} else {
																											m.G0 = v76 + int32(128)
																											m.G0 = v14 + int32(16)
																											return
																										}
																									}
																								} else {
																									v198 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																									v199 = F_get_namespace_name(m, v198)
																									mBase = m.M
																									v200 = m.ExcPending
																									if v200 != 0 {
																										return
																									} else {
																										F_aclcheck_error(m, v193, int32(36), v199)
																										mBase = m.M
																										v202 = m.ExcPending
																										if v202 != 0 {
																											return
																										} else {
																											F_AlterTypeOwner_oid(m, v88, v18)
																											mBase = m.M
																											v205 = m.ExcPending
																											if v205 != 0 {
																												return
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																												*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																												F_sequence_close(m, v80, int32(3))
																												mBase = m.M
																												v214 = m.ExcPending
																												if v214 != 0 {
																													return
																												} else {
																													m.G0 = v76 + int32(128)
																													m.G0 = v14 + int32(16)
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
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																		F_sequence_close(m, v80, int32(3))
																		mBase = m.M
																		v214 = m.ExcPending
																		if v214 != 0 {
																			return
																		} else {
																			m.G0 = v76 + int32(128)
																			m.G0 = v14 + int32(16)
																			return
																		}
																	}
																}
															}
														} else {
															v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+79)))
															if v137 == int32(109) {
																v140 = F_get_multirange_range(m, v88)
																mBase = m.M
																v141 = m.ExcPending
																if v141 != 0 {
																	return
																} else {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v145 = m.ExcPending
																	if v145 != 0 {
																		return
																	} else {
																		F_errcode(m, int32(151027844))
																		mBase = m.M
																		v148 = m.ExcPending
																		if v148 != 0 {
																			return
																		} else {
																			v149 = F_format_type_be(m, v88)
																			mBase = m.M
																			v150 = m.ExcPending
																			if v150 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v76)+32)) = v149
																				F_errmsg(m, int32(201809), v76+int32(32))
																				mBase = m.M
																				v156 = m.ExcPending
																				if v156 != 0 {
																					return
																				} else {
																					if v140 != 0 {
																						v157 = F_format_type_be(m, v140)
																						mBase = m.M
																						v158 = m.ExcPending
																						if v158 != 0 {
																							return
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v157
																							F_errhint(m, int32(645184), v76+int32(16))
																							mBase = m.M
																							v164 = m.ExcPending
																							if v164 != 0 {
																								return
																							} else {
																								F_errfinish(m, int32(514319), int32(3895), int32(228023))
																								mBase = m.M
																								v169 = m.ExcPending
																								if v169 != 0 {
																									return
																								} else {
																									base.Wasm_trap_unreachable()
																									for {
																									}
																								}
																							}
																						}
																					} else {
																						F_errfinish(m, int32(514319), int32(3895), int32(228023))
																						mBase = m.M
																						v169 = m.ExcPending
																						if v169 != 0 {
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
																v170 = *(*int32)(unsafe.Add(mBase, uint32(v96)+72))
																if v18 != v170 {
																	v172 = F_superuser(m)
																	mBase = m.M
																	v173 = m.ExcPending
																	if v173 != 0 {
																		return
																	} else {
																		if v172 != 0 {
																			F_AlterTypeOwner_oid(m, v88, v18)
																			mBase = m.M
																			v205 = m.ExcPending
																			if v205 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																				F_sequence_close(m, v80, int32(3))
																				mBase = m.M
																				v214 = m.ExcPending
																				if v214 != 0 {
																					return
																				} else {
																					m.G0 = v76 + int32(128)
																					m.G0 = v14 + int32(16)
																					return
																				}
																			}
																		} else {
																			v175 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																			v177 = *(*int32)(unsafe.Add(mBase, _consts[4]))
																			v178 = F_object_ownercheck(m, int32(1247), v175, v177)
																			mBase = m.M
																			v179 = m.ExcPending
																			if v179 != 0 {
																				return
																			} else {
																				if v178 == int32(0) {
																					v183 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																					F_aclcheck_error_type(m, int32(2), v183)
																					mBase = m.M
																					v185 = m.ExcPending
																					if v185 != 0 {
																						return
																					} else {
																						v187 = *(*int32)(unsafe.Add(mBase, _consts[4]))
																						F_check_can_set_role(m, v187, v18)
																						mBase = m.M
																						v189 = m.ExcPending
																						if v189 != 0 {
																							return
																						} else {
																							v191 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																							v193 = F_object_aclcheck(m, int32(2615), v191, v18, int64(512))
																							mBase = m.M
																							v194 = m.ExcPending
																							if v194 != 0 {
																								return
																							} else {
																								if v193 == int32(0) {
																									F_AlterTypeOwner_oid(m, v88, v18)
																									mBase = m.M
																									v205 = m.ExcPending
																									if v205 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																										F_sequence_close(m, v80, int32(3))
																										mBase = m.M
																										v214 = m.ExcPending
																										if v214 != 0 {
																											return
																										} else {
																											m.G0 = v76 + int32(128)
																											m.G0 = v14 + int32(16)
																											return
																										}
																									}
																								} else {
																									v198 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																									v199 = F_get_namespace_name(m, v198)
																									mBase = m.M
																									v200 = m.ExcPending
																									if v200 != 0 {
																										return
																									} else {
																										F_aclcheck_error(m, v193, int32(36), v199)
																										mBase = m.M
																										v202 = m.ExcPending
																										if v202 != 0 {
																											return
																										} else {
																											F_AlterTypeOwner_oid(m, v88, v18)
																											mBase = m.M
																											v205 = m.ExcPending
																											if v205 != 0 {
																												return
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																												*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																												F_sequence_close(m, v80, int32(3))
																												mBase = m.M
																												v214 = m.ExcPending
																												if v214 != 0 {
																													return
																												} else {
																													m.G0 = v76 + int32(128)
																													m.G0 = v14 + int32(16)
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
																					v187 = *(*int32)(unsafe.Add(mBase, _consts[4]))
																					F_check_can_set_role(m, v187, v18)
																					mBase = m.M
																					v189 = m.ExcPending
																					if v189 != 0 {
																						return
																					} else {
																						v191 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																						v193 = F_object_aclcheck(m, int32(2615), v191, v18, int64(512))
																						mBase = m.M
																						v194 = m.ExcPending
																						if v194 != 0 {
																							return
																						} else {
																							if v193 == int32(0) {
																								F_AlterTypeOwner_oid(m, v88, v18)
																								mBase = m.M
																								v205 = m.ExcPending
																								if v205 != 0 {
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																									F_sequence_close(m, v80, int32(3))
																									mBase = m.M
																									v214 = m.ExcPending
																									if v214 != 0 {
																										return
																									} else {
																										m.G0 = v76 + int32(128)
																										m.G0 = v14 + int32(16)
																										return
																									}
																								}
																							} else {
																								v198 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																								v199 = F_get_namespace_name(m, v198)
																								mBase = m.M
																								v200 = m.ExcPending
																								if v200 != 0 {
																									return
																								} else {
																									F_aclcheck_error(m, v193, int32(36), v199)
																									mBase = m.M
																									v202 = m.ExcPending
																									if v202 != 0 {
																										return
																									} else {
																										F_AlterTypeOwner_oid(m, v88, v18)
																										mBase = m.M
																										v205 = m.ExcPending
																										if v205 != 0 {
																											return
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																											F_sequence_close(m, v80, int32(3))
																											mBase = m.M
																											v214 = m.ExcPending
																											if v214 != 0 {
																												return
																											} else {
																												m.G0 = v76 + int32(128)
																												m.G0 = v14 + int32(16)
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
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																	F_sequence_close(m, v80, int32(3))
																	mBase = m.M
																	v214 = m.ExcPending
																	if v214 != 0 {
																		return
																	} else {
																		m.G0 = v76 + int32(128)
																		m.G0 = v14 + int32(16)
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
							v221 = m.ExcPending
							if v221 != 0 {
								return
							} else {
								F_errcode(m, int32(67137668))
								mBase = m.M
								v224 = m.ExcPending
								if v224 != 0 {
									return
								} else {
									v225 = F_TypeNameToString(m, v83)
									mBase = m.M
									v226 = m.ExcPending
									if v226 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v76))) = v225
										F_errmsg(m, int32(77684), v76)
										mBase = m.M
										v230 = m.ExcPending
										if v230 != 0 {
											return
										} else {
											F_errfinish(m, int32(514319), int32(3843), int32(228023))
											mBase = m.M
											v235 = m.ExcPending
											if v235 != 0 {
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
				v406 = F_SearchSysCacheCopy(m, int32(25), v395, int32(0))
				mBase = m.M
				v407 = m.ExcPending
				if v407 != 0 {
					return
				} else {
					if v406 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v413 = m.ExcPending
						if v413 != 0 {
							return
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v416 = m.ExcPending
							if v416 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v398))) = v395
								F_errmsg(m, int32(76535), v398)
								mBase = m.M
								v420 = m.ExcPending
								if v420 != 0 {
									return
								} else {
									F_errfinish(m, int32(515466), int32(494), int32(227904))
									mBase = m.M
									v425 = m.ExcPending
									if v425 != 0 {
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
						v426 = *(*int32)(unsafe.Add(mBase, uint32(v406)+16))
						v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426)+22)))
						v429 = *(*int32)(unsafe.Add(mBase, uint32(v426+v427)))
						F_AlterEventTriggerOwner_internal(m, v402, v406, v18)
						mBase = m.M
						v431 = m.ExcPending
						if v431 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v429
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(3466)
							F_pfree(m, v406)
							mBase = m.M
							v438 = m.ExcPending
							if v438 != 0 {
								return
							} else {
								F_sequence_close(m, v402, int32(3))
								mBase = m.M
								v441 = m.ExcPending
								if v441 != 0 {
									return
								} else {
									m.G0 = v398 + int32(16)
									m.G0 = v14 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		case 15:
			v292 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
			v294 = m.G0
			v296 = v294 - int32(16)
			m.G0 = v296
			v300 = F_table_open(m, int32(2328), int32(3))
			mBase = m.M
			v301 = m.ExcPending
			if v301 != 0 {
				return
			} else {
				v304 = F_SearchSysCacheCopy(m, int32(29), v293, int32(0))
				mBase = m.M
				v305 = m.ExcPending
				if v305 != 0 {
					return
				} else {
					if v304 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v311 = m.ExcPending
						if v311 != 0 {
							return
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v314 = m.ExcPending
							if v314 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v296))) = v293
								F_errmsg(m, int32(76437), v296)
								mBase = m.M
								v318 = m.ExcPending
								if v318 != 0 {
									return
								} else {
									F_errfinish(m, int32(514242), int32(302), int32(227875))
									mBase = m.M
									v323 = m.ExcPending
									if v323 != 0 {
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
						v324 = *(*int32)(unsafe.Add(mBase, uint32(v304)+16))
						v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+22)))
						v327 = *(*int32)(unsafe.Add(mBase, uint32(v324+v325)))
						F_AlterForeignDataWrapperOwner_internal(m, v300, v304, v18)
						mBase = m.M
						v329 = m.ExcPending
						if v329 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v327
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2328)
							F_pfree(m, v304)
							mBase = m.M
							v336 = m.ExcPending
							if v336 != 0 {
								return
							} else {
								F_sequence_close(m, v300, int32(3))
								mBase = m.M
								v339 = m.ExcPending
								if v339 != 0 {
									return
								} else {
									m.G0 = v296 + int32(16)
									m.G0 = v14 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		case 16:
			v343 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+4))
			v345 = m.G0
			v347 = v345 - int32(16)
			m.G0 = v347
			v351 = F_table_open(m, int32(1417), int32(3))
			mBase = m.M
			v352 = m.ExcPending
			if v352 != 0 {
				return
			} else {
				v355 = F_SearchSysCacheCopy(m, int32(31), v344, int32(0))
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
								*(*int32)(unsafe.Add(mBase, uint32(v347))) = v344
								F_errmsg(m, int32(76306), v347)
								mBase = m.M
								v369 = m.ExcPending
								if v369 != 0 {
									return
								} else {
									F_errfinish(m, int32(514242), int32(441), int32(227851))
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
						F_AlterForeignServerOwner_internal(m, v351, v355, v18)
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
								F_sequence_close(m, v351, int32(3))
								mBase = m.M
								v390 = m.ExcPending
								if v390 != 0 {
									return
								} else {
									m.G0 = v347 + int32(16)
									m.G0 = v14 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		case 29:
			v445 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
			v447 = m.G0
			v449 = v447 - int32(16)
			m.G0 = v449
			v453 = F_table_open(m, int32(6104), int32(3))
			mBase = m.M
			v454 = m.ExcPending
			if v454 != 0 {
				return
			} else {
				v457 = F_SearchSysCacheCopy(m, int32(48), v446, int32(0))
				mBase = m.M
				v458 = m.ExcPending
				if v458 != 0 {
					return
				} else {
					if v457 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v464 = m.ExcPending
						if v464 != 0 {
							return
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v467 = m.ExcPending
							if v467 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v449))) = v446
								F_errmsg(m, int32(77149), v449)
								mBase = m.M
								v471 = m.ExcPending
								if v471 != 0 {
									return
								} else {
									F_errfinish(m, int32(514207), int32(2073), int32(227950))
									mBase = m.M
									v476 = m.ExcPending
									if v476 != 0 {
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
						v477 = *(*int32)(unsafe.Add(mBase, uint32(v457)+16))
						v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477)+22)))
						v480 = *(*int32)(unsafe.Add(mBase, uint32(v477+v478)))
						F_AlterPublicationOwner_internal(m, v453, v457, v18)
						mBase = m.M
						v482 = m.ExcPending
						if v482 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v480
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(6104)
							F_pfree(m, v457)
							mBase = m.M
							v489 = m.ExcPending
							if v489 != 0 {
								return
							} else {
								F_sequence_close(m, v453, int32(3))
								mBase = m.M
								v492 = m.ExcPending
								if v492 != 0 {
									return
								} else {
									m.G0 = v449 + int32(16)
									m.G0 = v14 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		case 35:
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
			v25 = m.G0
			v27 = v25 - int32(16)
			m.G0 = v27
			v31 = F_table_open(m, int32(2615), int32(3))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v34 = F_SearchSysCache1(m, int32(37), v24)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					if v34 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							F_errcode(m, int32(1411))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v27))) = v24
								F_errmsg(m, int32(78120), v27)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									F_errfinish(m, int32(514342), int32(344), int32(228056))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
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
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+22)))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v54+v55)))
						F_AlterSchemaOwner_internal(m, v34, v31, v18)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v57
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2615)
							F_ReleaseCatCache(m, v34)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								F_sequence_close(m, v31, int32(3))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									m.G0 = v27 + int32(16)
									m.G0 = v14 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		case 37:
			v496 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)+4))
			v498 = m.G0
			v500 = v498 - int32(16)
			m.G0 = v500
			v504 = F_table_open(m, int32(6100), int32(3))
			mBase = m.M
			v505 = m.ExcPending
			if v505 != 0 {
				return
			} else {
				v508 = *(*int32)(unsafe.Add(mBase, _consts[223]))
				v509 = F_SearchSysCacheCopy(m, int32(66), v508, v497)
				mBase = m.M
				v510 = m.ExcPending
				if v510 != 0 {
					return
				} else {
					if v509 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v516 = m.ExcPending
						if v516 != 0 {
							return
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v519 = m.ExcPending
							if v519 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v500))) = v497
								F_errmsg(m, int32(76664), v500)
								mBase = m.M
								v523 = m.ExcPending
								if v523 != 0 {
									return
								} else {
									F_errfinish(m, int32(514157), int32(2046), int32(227927))
									mBase = m.M
									v528 = m.ExcPending
									if v528 != 0 {
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
						v529 = *(*int32)(unsafe.Add(mBase, uint32(v509)+16))
						v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529)+22)))
						v532 = *(*int32)(unsafe.Add(mBase, uint32(v529+v530)))
						F_AlterSubscriptionOwner_internal(m, v504, v509, v18)
						mBase = m.M
						v534 = m.ExcPending
						if v534 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v532
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(6100)
							F_pfree(m, v509)
							mBase = m.M
							v541 = m.ExcPending
							if v541 != 0 {
								return
							} else {
								F_sequence_close(m, v504, int32(3))
								mBase = m.M
								v544 = m.ExcPending
								if v544 != 0 {
									return
								} else {
									m.G0 = v500 + int32(16)
									m.G0 = v14 + int32(16)
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
