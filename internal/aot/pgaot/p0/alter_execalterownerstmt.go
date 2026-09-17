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
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
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
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
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
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
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
	var v601 int64
	_ = v601
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
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
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
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
			v544 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v545 = int32(0)
			F_get_object_address(m, l0, v20, v544, v545, int32(8), v545)
			mBase = m.M
			v549 = m.ExcPending
			if v549 != 0 {
				return
			} else {
				v550 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_AlterObjectOwner_internal(m, v550, v551, v18)
				mBase = m.M
				v553 = m.ExcPending
				if v553 != 0 {
					return
				} else {
					m.G0 = v14 + int32(16)
					return
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v557 = m.ExcPending
			if v557 != 0 {
				return
			} else {
				v558 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v558
				F_errmsg_internal(m, int32(_a_F_ExecAlterOwnerStmt_0), v14)
				mBase = m.M
				v562 = m.ExcPending
				if v562 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_1), int32(908), int32(_a_F_ExecAlterOwnerStmt_2))
					mBase = m.M
					v567 = m.ExcPending
					if v567 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 8:
			v568 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v569 = *(*int32)(unsafe.Add(mBase, uint32(v568)+4))
			v570 = m.G0
			v572 = v570 - int32(208)
			m.G0 = v572
			v576 = F_table_open(m, int32(1262), int32(3))
			mBase = m.M
			v577 = m.ExcPending
			if v577 != 0 {
				return
			} else {
				v579 = v572 + int32(160)
				F_ScanKeyInit(m, v579, int32(2), int32(3), int32(62), v569)
				mBase = m.M
				v584 = m.ExcPending
				if v584 != 0 {
					return
				} else {
					v586 = int32(1)
					v589 = F_systable_beginscan(m, v576, int32(2671), v586, int32(0), v586, v579)
					mBase = m.M
					v590 = m.ExcPending
					if v590 != 0 {
						return
					} else {
						v591 = F_systable_getnext(m, v589)
						mBase = m.M
						v592 = m.ExcPending
						if v592 != 0 {
							return
						} else {
							if v591 != 0 {
								v593 = *(*int32)(unsafe.Add(mBase, uint32(v591)+16))
								v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593)+22)))
								v595 = v593 + v594
								v596 = *(*int32)(unsafe.Add(mBase, uint32(v595)))
								v597 = *(*int32)(unsafe.Add(mBase, uint32(v595)+68))
								if v18 != v597 {
									v599 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v572)+64)) = uint16(v599)
									v601 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v572)+56)) = v601
									*(*int64)(unsafe.Add(mBase, uint32(v572)+48)) = v601
									*(*uint16)(unsafe.Add(mBase, uint32(v572)+32)) = uint16(v599)
									*(*int64)(unsafe.Add(mBase, uint32(v572)+24)) = v601
									*(*int64)(unsafe.Add(mBase, uint32(v572)+16)) = v601
									v613 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
									v614 = F_object_ownercheck(m, int32(1262), v596, v613)
									mBase = m.M
									v615 = m.ExcPending
									if v615 != 0 {
										return
									} else {
										if v614 == int32(0) {
											F_aclcheck_error(m, int32(2), int32(9), v569)
											mBase = m.M
											v621 = m.ExcPending
											if v621 != 0 {
												return
											} else {
												v623 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
												F_check_can_set_role(m, v623, v18)
												mBase = m.M
												v625 = m.ExcPending
												if v625 != 0 {
													return
												} else {
													v626 = F_superuser(m)
													mBase = m.M
													v627 = m.ExcPending
													if v627 != 0 {
														return
													} else {
														if v626 == int32(0) {
															v632 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
															v633 = F_SearchSysCache1(m, int32(11), v632)
															mBase = m.M
															v634 = m.ExcPending
															if v634 != 0 {
																return
															} else {
																if v633 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v737 = m.ExcPending
																	if v737 != 0 {
																		return
																	} else {
																		F_errcode(m, int32(16797828))
																		mBase = m.M
																		v740 = m.ExcPending
																		if v740 != 0 {
																			return
																		} else {
																			F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_3), int32(0))
																			mBase = m.M
																			v744 = m.ExcPending
																			if v744 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_4), int32(2730), int32(_a_F_ExecAlterOwnerStmt_5))
																				mBase = m.M
																				v749 = m.ExcPending
																				if v749 != 0 {
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
																	v637 = *(*int32)(unsafe.Add(mBase, uint32(v633)+16))
																	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637)+22)))
																	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637+v638)+71)))
																	F_ReleaseCatCache(m, v633)
																	mBase = m.M
																	v642 = m.ExcPending
																	if v642 != 0 {
																		return
																	} else {
																		if v640 == int32(0) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v737 = m.ExcPending
																			if v737 != 0 {
																				return
																			} else {
																				F_errcode(m, int32(16797828))
																				mBase = m.M
																				v740 = m.ExcPending
																				if v740 != 0 {
																					return
																				} else {
																					F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_3), int32(0))
																					mBase = m.M
																					v744 = m.ExcPending
																					if v744 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_4), int32(2730), int32(_a_F_ExecAlterOwnerStmt_5))
																						mBase = m.M
																						v749 = m.ExcPending
																						if v749 != 0 {
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
																			v648 = v591 + int32(4)
																			F_LockTuple(m, v576, v648, int32(7))
																			mBase = m.M
																			v651 = m.ExcPending
																			if v651 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v572)+88)) = v18
																				v653 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v572)+18)) = uint8(v653)
																				v656 = *(*int32)(unsafe.Add(mBase, uint32(v576)+52))
																				v659 = F_heap_getattr_6(m, v591, int32(18), v656, v572+int32(15))
																				mBase = m.M
																				v660 = m.ExcPending
																				if v660 != 0 {
																					return
																				} else {
																					v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+15)))
																					if v661 == int32(0) {
																						v664 = F_pg_detoast_datum(m, v659)
																						mBase = m.M
																						v665 = m.ExcPending
																						if v665 != 0 {
																							return
																						} else {
																							v666 = *(*int32)(unsafe.Add(mBase, uint32(v595)+68))
																							v667 = F_aclnewowner(m, v664, v666, v18)
																							mBase = m.M
																							v668 = m.ExcPending
																							if v668 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v572)+148)) = v667
																								v670 = int32(1)
																								*(*uint8)(unsafe.Add(mBase, uint32(v572)+33)) = uint8(v670)
																								v672 = *(*int32)(unsafe.Add(mBase, uint32(v576)+52))
																								v679 = F_heap_modify_tuple(m, v591, v672, v572+int32(80), v572+int32(48), v572+int32(16))
																								mBase = m.M
																								v680 = m.ExcPending
																								if v680 != 0 {
																									return
																								} else {
																									F_CatalogTupleUpdate(m, v576, v679+int32(4), v679)
																									mBase = m.M
																									v684 = m.ExcPending
																									if v684 != 0 {
																										return
																									} else {
																										F_UnlockTuple(m, v576, v648, int32(7))
																										mBase = m.M
																										v687 = m.ExcPending
																										if v687 != 0 {
																											return
																										} else {
																											F_pfree(m, v679)
																											mBase = m.M
																											v689 = m.ExcPending
																											if v689 != 0 {
																												return
																											} else {
																												F_changeDependencyOnOwner(m, int32(1262), v596, v18)
																												mBase = m.M
																												v692 = m.ExcPending
																												if v692 != 0 {
																													return
																												} else {
																													v697 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[1]))
																													if v697 != 0 {
																														v699 = int32(0)
																														F_RunObjectPostAlterHook(m, int32(1262), v596, v699, v699, v699)
																														mBase = m.M
																														v703 = m.ExcPending
																														if v703 != 0 {
																															return
																														} else {
																															*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
																															*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																															F_systable_endscan(m, v589)
																															mBase = m.M
																															v710 = m.ExcPending
																															if v710 != 0 {
																																return
																															} else {
																																F_relation_close(m, v576, int32(0))
																																mBase = m.M
																																v713 = m.ExcPending
																																if v713 != 0 {
																																	return
																																} else {
																																	m.G0 = v572 + int32(208)
																																	m.G0 = v14 + int32(16)
																																	return
																																}
																															}
																														}
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
																														*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																														F_systable_endscan(m, v589)
																														mBase = m.M
																														v710 = m.ExcPending
																														if v710 != 0 {
																															return
																														} else {
																															F_relation_close(m, v576, int32(0))
																															mBase = m.M
																															v713 = m.ExcPending
																															if v713 != 0 {
																																return
																															} else {
																																m.G0 = v572 + int32(208)
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
																						v672 = *(*int32)(unsafe.Add(mBase, uint32(v576)+52))
																						v679 = F_heap_modify_tuple(m, v591, v672, v572+int32(80), v572+int32(48), v572+int32(16))
																						mBase = m.M
																						v680 = m.ExcPending
																						if v680 != 0 {
																							return
																						} else {
																							F_CatalogTupleUpdate(m, v576, v679+int32(4), v679)
																							mBase = m.M
																							v684 = m.ExcPending
																							if v684 != 0 {
																								return
																							} else {
																								F_UnlockTuple(m, v576, v648, int32(7))
																								mBase = m.M
																								v687 = m.ExcPending
																								if v687 != 0 {
																									return
																								} else {
																									F_pfree(m, v679)
																									mBase = m.M
																									v689 = m.ExcPending
																									if v689 != 0 {
																										return
																									} else {
																										F_changeDependencyOnOwner(m, int32(1262), v596, v18)
																										mBase = m.M
																										v692 = m.ExcPending
																										if v692 != 0 {
																											return
																										} else {
																											v697 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[1]))
																											if v697 != 0 {
																												v699 = int32(0)
																												F_RunObjectPostAlterHook(m, int32(1262), v596, v699, v699, v699)
																												mBase = m.M
																												v703 = m.ExcPending
																												if v703 != 0 {
																													return
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
																													*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																													F_systable_endscan(m, v589)
																													mBase = m.M
																													v710 = m.ExcPending
																													if v710 != 0 {
																														return
																													} else {
																														F_relation_close(m, v576, int32(0))
																														mBase = m.M
																														v713 = m.ExcPending
																														if v713 != 0 {
																															return
																														} else {
																															m.G0 = v572 + int32(208)
																															m.G0 = v14 + int32(16)
																															return
																														}
																													}
																												}
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
																												*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																												F_systable_endscan(m, v589)
																												mBase = m.M
																												v710 = m.ExcPending
																												if v710 != 0 {
																													return
																												} else {
																													F_relation_close(m, v576, int32(0))
																													mBase = m.M
																													v713 = m.ExcPending
																													if v713 != 0 {
																														return
																													} else {
																														m.G0 = v572 + int32(208)
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
															v648 = v591 + int32(4)
															F_LockTuple(m, v576, v648, int32(7))
															mBase = m.M
															v651 = m.ExcPending
															if v651 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v572)+88)) = v18
																v653 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v572)+18)) = uint8(v653)
																v656 = *(*int32)(unsafe.Add(mBase, uint32(v576)+52))
																v659 = F_heap_getattr_6(m, v591, int32(18), v656, v572+int32(15))
																mBase = m.M
																v660 = m.ExcPending
																if v660 != 0 {
																	return
																} else {
																	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+15)))
																	if v661 == int32(0) {
																		v664 = F_pg_detoast_datum(m, v659)
																		mBase = m.M
																		v665 = m.ExcPending
																		if v665 != 0 {
																			return
																		} else {
																			v666 = *(*int32)(unsafe.Add(mBase, uint32(v595)+68))
																			v667 = F_aclnewowner(m, v664, v666, v18)
																			mBase = m.M
																			v668 = m.ExcPending
																			if v668 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v572)+148)) = v667
																				v670 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v572)+33)) = uint8(v670)
																				v672 = *(*int32)(unsafe.Add(mBase, uint32(v576)+52))
																				v679 = F_heap_modify_tuple(m, v591, v672, v572+int32(80), v572+int32(48), v572+int32(16))
																				mBase = m.M
																				v680 = m.ExcPending
																				if v680 != 0 {
																					return
																				} else {
																					F_CatalogTupleUpdate(m, v576, v679+int32(4), v679)
																					mBase = m.M
																					v684 = m.ExcPending
																					if v684 != 0 {
																						return
																					} else {
																						F_UnlockTuple(m, v576, v648, int32(7))
																						mBase = m.M
																						v687 = m.ExcPending
																						if v687 != 0 {
																							return
																						} else {
																							F_pfree(m, v679)
																							mBase = m.M
																							v689 = m.ExcPending
																							if v689 != 0 {
																								return
																							} else {
																								F_changeDependencyOnOwner(m, int32(1262), v596, v18)
																								mBase = m.M
																								v692 = m.ExcPending
																								if v692 != 0 {
																									return
																								} else {
																									v697 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[1]))
																									if v697 != 0 {
																										v699 = int32(0)
																										F_RunObjectPostAlterHook(m, int32(1262), v596, v699, v699, v699)
																										mBase = m.M
																										v703 = m.ExcPending
																										if v703 != 0 {
																											return
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
																											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																											F_systable_endscan(m, v589)
																											mBase = m.M
																											v710 = m.ExcPending
																											if v710 != 0 {
																												return
																											} else {
																												F_relation_close(m, v576, int32(0))
																												mBase = m.M
																												v713 = m.ExcPending
																												if v713 != 0 {
																													return
																												} else {
																													m.G0 = v572 + int32(208)
																													m.G0 = v14 + int32(16)
																													return
																												}
																											}
																										}
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																										F_systable_endscan(m, v589)
																										mBase = m.M
																										v710 = m.ExcPending
																										if v710 != 0 {
																											return
																										} else {
																											F_relation_close(m, v576, int32(0))
																											mBase = m.M
																											v713 = m.ExcPending
																											if v713 != 0 {
																												return
																											} else {
																												m.G0 = v572 + int32(208)
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
																		v672 = *(*int32)(unsafe.Add(mBase, uint32(v576)+52))
																		v679 = F_heap_modify_tuple(m, v591, v672, v572+int32(80), v572+int32(48), v572+int32(16))
																		mBase = m.M
																		v680 = m.ExcPending
																		if v680 != 0 {
																			return
																		} else {
																			F_CatalogTupleUpdate(m, v576, v679+int32(4), v679)
																			mBase = m.M
																			v684 = m.ExcPending
																			if v684 != 0 {
																				return
																			} else {
																				F_UnlockTuple(m, v576, v648, int32(7))
																				mBase = m.M
																				v687 = m.ExcPending
																				if v687 != 0 {
																					return
																				} else {
																					F_pfree(m, v679)
																					mBase = m.M
																					v689 = m.ExcPending
																					if v689 != 0 {
																						return
																					} else {
																						F_changeDependencyOnOwner(m, int32(1262), v596, v18)
																						mBase = m.M
																						v692 = m.ExcPending
																						if v692 != 0 {
																							return
																						} else {
																							v697 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[1]))
																							if v697 != 0 {
																								v699 = int32(0)
																								F_RunObjectPostAlterHook(m, int32(1262), v596, v699, v699, v699)
																								mBase = m.M
																								v703 = m.ExcPending
																								if v703 != 0 {
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																									F_systable_endscan(m, v589)
																									mBase = m.M
																									v710 = m.ExcPending
																									if v710 != 0 {
																										return
																									} else {
																										F_relation_close(m, v576, int32(0))
																										mBase = m.M
																										v713 = m.ExcPending
																										if v713 != 0 {
																											return
																										} else {
																											m.G0 = v572 + int32(208)
																											m.G0 = v14 + int32(16)
																											return
																										}
																									}
																								}
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
																								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																								F_systable_endscan(m, v589)
																								mBase = m.M
																								v710 = m.ExcPending
																								if v710 != 0 {
																									return
																								} else {
																									F_relation_close(m, v576, int32(0))
																									mBase = m.M
																									v713 = m.ExcPending
																									if v713 != 0 {
																										return
																									} else {
																										m.G0 = v572 + int32(208)
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
											v623 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
											F_check_can_set_role(m, v623, v18)
											mBase = m.M
											v625 = m.ExcPending
											if v625 != 0 {
												return
											} else {
												v626 = F_superuser(m)
												mBase = m.M
												v627 = m.ExcPending
												if v627 != 0 {
													return
												} else {
													if v626 == int32(0) {
														v632 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
														v633 = F_SearchSysCache1(m, int32(11), v632)
														mBase = m.M
														v634 = m.ExcPending
														if v634 != 0 {
															return
														} else {
															if v633 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v737 = m.ExcPending
																if v737 != 0 {
																	return
																} else {
																	F_errcode(m, int32(16797828))
																	mBase = m.M
																	v740 = m.ExcPending
																	if v740 != 0 {
																		return
																	} else {
																		F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_3), int32(0))
																		mBase = m.M
																		v744 = m.ExcPending
																		if v744 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_4), int32(2730), int32(_a_F_ExecAlterOwnerStmt_5))
																			mBase = m.M
																			v749 = m.ExcPending
																			if v749 != 0 {
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
																v637 = *(*int32)(unsafe.Add(mBase, uint32(v633)+16))
																v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637)+22)))
																v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637+v638)+71)))
																F_ReleaseCatCache(m, v633)
																mBase = m.M
																v642 = m.ExcPending
																if v642 != 0 {
																	return
																} else {
																	if v640 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v737 = m.ExcPending
																		if v737 != 0 {
																			return
																		} else {
																			F_errcode(m, int32(16797828))
																			mBase = m.M
																			v740 = m.ExcPending
																			if v740 != 0 {
																				return
																			} else {
																				F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_3), int32(0))
																				mBase = m.M
																				v744 = m.ExcPending
																				if v744 != 0 {
																					return
																				} else {
																					F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_4), int32(2730), int32(_a_F_ExecAlterOwnerStmt_5))
																					mBase = m.M
																					v749 = m.ExcPending
																					if v749 != 0 {
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
																		v648 = v591 + int32(4)
																		F_LockTuple(m, v576, v648, int32(7))
																		mBase = m.M
																		v651 = m.ExcPending
																		if v651 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v572)+88)) = v18
																			v653 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v572)+18)) = uint8(v653)
																			v656 = *(*int32)(unsafe.Add(mBase, uint32(v576)+52))
																			v659 = F_heap_getattr_6(m, v591, int32(18), v656, v572+int32(15))
																			mBase = m.M
																			v660 = m.ExcPending
																			if v660 != 0 {
																				return
																			} else {
																				v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+15)))
																				if v661 == int32(0) {
																					v664 = F_pg_detoast_datum(m, v659)
																					mBase = m.M
																					v665 = m.ExcPending
																					if v665 != 0 {
																						return
																					} else {
																						v666 = *(*int32)(unsafe.Add(mBase, uint32(v595)+68))
																						v667 = F_aclnewowner(m, v664, v666, v18)
																						mBase = m.M
																						v668 = m.ExcPending
																						if v668 != 0 {
																							return
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v572)+148)) = v667
																							v670 = int32(1)
																							*(*uint8)(unsafe.Add(mBase, uint32(v572)+33)) = uint8(v670)
																							v672 = *(*int32)(unsafe.Add(mBase, uint32(v576)+52))
																							v679 = F_heap_modify_tuple(m, v591, v672, v572+int32(80), v572+int32(48), v572+int32(16))
																							mBase = m.M
																							v680 = m.ExcPending
																							if v680 != 0 {
																								return
																							} else {
																								F_CatalogTupleUpdate(m, v576, v679+int32(4), v679)
																								mBase = m.M
																								v684 = m.ExcPending
																								if v684 != 0 {
																									return
																								} else {
																									F_UnlockTuple(m, v576, v648, int32(7))
																									mBase = m.M
																									v687 = m.ExcPending
																									if v687 != 0 {
																										return
																									} else {
																										F_pfree(m, v679)
																										mBase = m.M
																										v689 = m.ExcPending
																										if v689 != 0 {
																											return
																										} else {
																											F_changeDependencyOnOwner(m, int32(1262), v596, v18)
																											mBase = m.M
																											v692 = m.ExcPending
																											if v692 != 0 {
																												return
																											} else {
																												v697 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[1]))
																												if v697 != 0 {
																													v699 = int32(0)
																													F_RunObjectPostAlterHook(m, int32(1262), v596, v699, v699, v699)
																													mBase = m.M
																													v703 = m.ExcPending
																													if v703 != 0 {
																														return
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
																														*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																														F_systable_endscan(m, v589)
																														mBase = m.M
																														v710 = m.ExcPending
																														if v710 != 0 {
																															return
																														} else {
																															F_relation_close(m, v576, int32(0))
																															mBase = m.M
																															v713 = m.ExcPending
																															if v713 != 0 {
																																return
																															} else {
																																m.G0 = v572 + int32(208)
																																m.G0 = v14 + int32(16)
																																return
																															}
																														}
																													}
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
																													*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																													F_systable_endscan(m, v589)
																													mBase = m.M
																													v710 = m.ExcPending
																													if v710 != 0 {
																														return
																													} else {
																														F_relation_close(m, v576, int32(0))
																														mBase = m.M
																														v713 = m.ExcPending
																														if v713 != 0 {
																															return
																														} else {
																															m.G0 = v572 + int32(208)
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
																					v672 = *(*int32)(unsafe.Add(mBase, uint32(v576)+52))
																					v679 = F_heap_modify_tuple(m, v591, v672, v572+int32(80), v572+int32(48), v572+int32(16))
																					mBase = m.M
																					v680 = m.ExcPending
																					if v680 != 0 {
																						return
																					} else {
																						F_CatalogTupleUpdate(m, v576, v679+int32(4), v679)
																						mBase = m.M
																						v684 = m.ExcPending
																						if v684 != 0 {
																							return
																						} else {
																							F_UnlockTuple(m, v576, v648, int32(7))
																							mBase = m.M
																							v687 = m.ExcPending
																							if v687 != 0 {
																								return
																							} else {
																								F_pfree(m, v679)
																								mBase = m.M
																								v689 = m.ExcPending
																								if v689 != 0 {
																									return
																								} else {
																									F_changeDependencyOnOwner(m, int32(1262), v596, v18)
																									mBase = m.M
																									v692 = m.ExcPending
																									if v692 != 0 {
																										return
																									} else {
																										v697 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[1]))
																										if v697 != 0 {
																											v699 = int32(0)
																											F_RunObjectPostAlterHook(m, int32(1262), v596, v699, v699, v699)
																											mBase = m.M
																											v703 = m.ExcPending
																											if v703 != 0 {
																												return
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
																												*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																												F_systable_endscan(m, v589)
																												mBase = m.M
																												v710 = m.ExcPending
																												if v710 != 0 {
																													return
																												} else {
																													F_relation_close(m, v576, int32(0))
																													mBase = m.M
																													v713 = m.ExcPending
																													if v713 != 0 {
																														return
																													} else {
																														m.G0 = v572 + int32(208)
																														m.G0 = v14 + int32(16)
																														return
																													}
																												}
																											}
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
																											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																											F_systable_endscan(m, v589)
																											mBase = m.M
																											v710 = m.ExcPending
																											if v710 != 0 {
																												return
																											} else {
																												F_relation_close(m, v576, int32(0))
																												mBase = m.M
																												v713 = m.ExcPending
																												if v713 != 0 {
																													return
																												} else {
																													m.G0 = v572 + int32(208)
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
														v648 = v591 + int32(4)
														F_LockTuple(m, v576, v648, int32(7))
														mBase = m.M
														v651 = m.ExcPending
														if v651 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v572)+88)) = v18
															v653 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v572)+18)) = uint8(v653)
															v656 = *(*int32)(unsafe.Add(mBase, uint32(v576)+52))
															v659 = F_heap_getattr_6(m, v591, int32(18), v656, v572+int32(15))
															mBase = m.M
															v660 = m.ExcPending
															if v660 != 0 {
																return
															} else {
																v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+15)))
																if v661 == int32(0) {
																	v664 = F_pg_detoast_datum(m, v659)
																	mBase = m.M
																	v665 = m.ExcPending
																	if v665 != 0 {
																		return
																	} else {
																		v666 = *(*int32)(unsafe.Add(mBase, uint32(v595)+68))
																		v667 = F_aclnewowner(m, v664, v666, v18)
																		mBase = m.M
																		v668 = m.ExcPending
																		if v668 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v572)+148)) = v667
																			v670 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v572)+33)) = uint8(v670)
																			v672 = *(*int32)(unsafe.Add(mBase, uint32(v576)+52))
																			v679 = F_heap_modify_tuple(m, v591, v672, v572+int32(80), v572+int32(48), v572+int32(16))
																			mBase = m.M
																			v680 = m.ExcPending
																			if v680 != 0 {
																				return
																			} else {
																				F_CatalogTupleUpdate(m, v576, v679+int32(4), v679)
																				mBase = m.M
																				v684 = m.ExcPending
																				if v684 != 0 {
																					return
																				} else {
																					F_UnlockTuple(m, v576, v648, int32(7))
																					mBase = m.M
																					v687 = m.ExcPending
																					if v687 != 0 {
																						return
																					} else {
																						F_pfree(m, v679)
																						mBase = m.M
																						v689 = m.ExcPending
																						if v689 != 0 {
																							return
																						} else {
																							F_changeDependencyOnOwner(m, int32(1262), v596, v18)
																							mBase = m.M
																							v692 = m.ExcPending
																							if v692 != 0 {
																								return
																							} else {
																								v697 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[1]))
																								if v697 != 0 {
																									v699 = int32(0)
																									F_RunObjectPostAlterHook(m, int32(1262), v596, v699, v699, v699)
																									mBase = m.M
																									v703 = m.ExcPending
																									if v703 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																										F_systable_endscan(m, v589)
																										mBase = m.M
																										v710 = m.ExcPending
																										if v710 != 0 {
																											return
																										} else {
																											F_relation_close(m, v576, int32(0))
																											mBase = m.M
																											v713 = m.ExcPending
																											if v713 != 0 {
																												return
																											} else {
																												m.G0 = v572 + int32(208)
																												m.G0 = v14 + int32(16)
																												return
																											}
																										}
																									}
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																									F_systable_endscan(m, v589)
																									mBase = m.M
																									v710 = m.ExcPending
																									if v710 != 0 {
																										return
																									} else {
																										F_relation_close(m, v576, int32(0))
																										mBase = m.M
																										v713 = m.ExcPending
																										if v713 != 0 {
																											return
																										} else {
																											m.G0 = v572 + int32(208)
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
																	v672 = *(*int32)(unsafe.Add(mBase, uint32(v576)+52))
																	v679 = F_heap_modify_tuple(m, v591, v672, v572+int32(80), v572+int32(48), v572+int32(16))
																	mBase = m.M
																	v680 = m.ExcPending
																	if v680 != 0 {
																		return
																	} else {
																		F_CatalogTupleUpdate(m, v576, v679+int32(4), v679)
																		mBase = m.M
																		v684 = m.ExcPending
																		if v684 != 0 {
																			return
																		} else {
																			F_UnlockTuple(m, v576, v648, int32(7))
																			mBase = m.M
																			v687 = m.ExcPending
																			if v687 != 0 {
																				return
																			} else {
																				F_pfree(m, v679)
																				mBase = m.M
																				v689 = m.ExcPending
																				if v689 != 0 {
																					return
																				} else {
																					F_changeDependencyOnOwner(m, int32(1262), v596, v18)
																					mBase = m.M
																					v692 = m.ExcPending
																					if v692 != 0 {
																						return
																					} else {
																						v697 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[1]))
																						if v697 != 0 {
																							v699 = int32(0)
																							F_RunObjectPostAlterHook(m, int32(1262), v596, v699, v699, v699)
																							mBase = m.M
																							v703 = m.ExcPending
																							if v703 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
																								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																								F_systable_endscan(m, v589)
																								mBase = m.M
																								v710 = m.ExcPending
																								if v710 != 0 {
																									return
																								} else {
																									F_relation_close(m, v576, int32(0))
																									mBase = m.M
																									v713 = m.ExcPending
																									if v713 != 0 {
																										return
																									} else {
																										m.G0 = v572 + int32(208)
																										m.G0 = v14 + int32(16)
																										return
																									}
																								}
																							}
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
																							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
																							F_systable_endscan(m, v589)
																							mBase = m.M
																							v710 = m.ExcPending
																							if v710 != 0 {
																								return
																							} else {
																								F_relation_close(m, v576, int32(0))
																								mBase = m.M
																								v713 = m.ExcPending
																								if v713 != 0 {
																									return
																								} else {
																									m.G0 = v572 + int32(208)
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
									v697 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[1]))
									if v697 != 0 {
										v699 = int32(0)
										F_RunObjectPostAlterHook(m, int32(1262), v596, v699, v699, v699)
										mBase = m.M
										v703 = m.ExcPending
										if v703 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
											F_systable_endscan(m, v589)
											mBase = m.M
											v710 = m.ExcPending
											if v710 != 0 {
												return
											} else {
												F_relation_close(m, v576, int32(0))
												mBase = m.M
												v713 = m.ExcPending
												if v713 != 0 {
													return
												} else {
													m.G0 = v572 + int32(208)
													m.G0 = v14 + int32(16)
													return
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
										F_systable_endscan(m, v589)
										mBase = m.M
										v710 = m.ExcPending
										if v710 != 0 {
											return
										} else {
											F_relation_close(m, v576, int32(0))
											mBase = m.M
											v713 = m.ExcPending
											if v713 != 0 {
												return
											} else {
												m.G0 = v572 + int32(208)
												m.G0 = v14 + int32(16)
												return
											}
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v720 = m.ExcPending
								if v720 != 0 {
									return
								} else {
									F_errcode(m, int32(1283))
									mBase = m.M
									v723 = m.ExcPending
									if v723 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v572))) = v569
										F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_6), v572)
										mBase = m.M
										v727 = m.ExcPending
										if v727 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_4), int32(2690), int32(_a_F_ExecAlterOwnerStmt_5))
											mBase = m.M
											v732 = m.ExcPending
											if v732 != 0 {
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
											if v97 == int32(100) {
												v129 = *(*int32)(unsafe.Add(mBase, uint32(v96)+92))
												if v129 != 0 {
													v130 = *(*int32)(unsafe.Add(mBase, uint32(v96)+88))
													if v130 == int32(_a_F_ExecAlterOwnerStmt_7) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v262 = m.ExcPending
														if v262 != 0 {
															return
														} else {
															F_errcode(m, int32(151027844))
															mBase = m.M
															v265 = m.ExcPending
															if v265 != 0 {
																return
															} else {
																v266 = F_format_type_be(m, v88)
																mBase = m.M
																v267 = m.ExcPending
																if v267 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v76)+64)) = v266
																	F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_8), v76-int32(-64))
																	mBase = m.M
																	v273 = m.ExcPending
																	if v273 != 0 {
																		return
																	} else {
																		v274 = *(*int32)(unsafe.Add(mBase, uint32(v96)+92))
																		v275 = F_format_type_be(m, v274)
																		mBase = m.M
																		v276 = m.ExcPending
																		if v276 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v76)+48)) = v275
																			F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_9), v76+int32(48))
																			mBase = m.M
																			v282 = m.ExcPending
																			if v282 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3881), int32(_a_F_ExecAlterOwnerStmt_11))
																				mBase = m.M
																				v287 = m.ExcPending
																				if v287 != 0 {
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
														v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+79)))
														if v133 == int32(109) {
															v136 = F_get_multirange_range(m, v88)
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
																return
															} else {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v141 = m.ExcPending
																if v141 != 0 {
																	return
																} else {
																	F_errcode(m, int32(151027844))
																	mBase = m.M
																	v144 = m.ExcPending
																	if v144 != 0 {
																		return
																	} else {
																		v145 = F_format_type_be(m, v88)
																		mBase = m.M
																		v146 = m.ExcPending
																		if v146 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v76)+32)) = v145
																			F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_12), v76+int32(32))
																			mBase = m.M
																			v152 = m.ExcPending
																			if v152 != 0 {
																				return
																			} else {
																				if v136 != 0 {
																					v153 = F_format_type_be(m, v136)
																					mBase = m.M
																					v154 = m.ExcPending
																					if v154 != 0 {
																						return
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v153
																						F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_13), v76+int32(16))
																						mBase = m.M
																						v160 = m.ExcPending
																						if v160 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3895), int32(_a_F_ExecAlterOwnerStmt_11))
																							mBase = m.M
																							v165 = m.ExcPending
																							if v165 != 0 {
																								return
																							} else {
																								base.Wasm_trap_unreachable()
																								for {
																								}
																							}
																						}
																					}
																				} else {
																					F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3895), int32(_a_F_ExecAlterOwnerStmt_11))
																					mBase = m.M
																					v165 = m.ExcPending
																					if v165 != 0 {
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
															v166 = *(*int32)(unsafe.Add(mBase, uint32(v96)+72))
															if v18 != v166 {
																v168 = F_superuser(m)
																mBase = m.M
																v169 = m.ExcPending
																if v169 != 0 {
																	return
																} else {
																	if v168 != 0 {
																		F_AlterTypeOwner_oid(m, v88, v18)
																		mBase = m.M
																		v201 = m.ExcPending
																		if v201 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																			F_relation_close(m, v80, int32(3))
																			mBase = m.M
																			v210 = m.ExcPending
																			if v210 != 0 {
																				return
																			} else {
																				m.G0 = v76 + int32(128)
																				m.G0 = v14 + int32(16)
																				return
																			}
																		}
																	} else {
																		v171 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																		v173 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																		v174 = F_object_ownercheck(m, int32(1247), v171, v173)
																		mBase = m.M
																		v175 = m.ExcPending
																		if v175 != 0 {
																			return
																		} else {
																			if v174 == int32(0) {
																				v179 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																				F_aclcheck_error_type(m, int32(2), v179)
																				mBase = m.M
																				v181 = m.ExcPending
																				if v181 != 0 {
																					return
																				} else {
																					v183 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																					F_check_can_set_role(m, v183, v18)
																					mBase = m.M
																					v185 = m.ExcPending
																					if v185 != 0 {
																						return
																					} else {
																						v187 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																						v189 = F_object_aclcheck(m, int32(2615), v187, v18, int64(512))
																						mBase = m.M
																						v190 = m.ExcPending
																						if v190 != 0 {
																							return
																						} else {
																							if v189 == int32(0) {
																								F_AlterTypeOwner_oid(m, v88, v18)
																								mBase = m.M
																								v201 = m.ExcPending
																								if v201 != 0 {
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																									F_relation_close(m, v80, int32(3))
																									mBase = m.M
																									v210 = m.ExcPending
																									if v210 != 0 {
																										return
																									} else {
																										m.G0 = v76 + int32(128)
																										m.G0 = v14 + int32(16)
																										return
																									}
																								}
																							} else {
																								v194 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																								v195 = F_get_namespace_name(m, v194)
																								mBase = m.M
																								v196 = m.ExcPending
																								if v196 != 0 {
																									return
																								} else {
																									F_aclcheck_error(m, v189, int32(36), v195)
																									mBase = m.M
																									v198 = m.ExcPending
																									if v198 != 0 {
																										return
																									} else {
																										F_AlterTypeOwner_oid(m, v88, v18)
																										mBase = m.M
																										v201 = m.ExcPending
																										if v201 != 0 {
																											return
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																											F_relation_close(m, v80, int32(3))
																											mBase = m.M
																											v210 = m.ExcPending
																											if v210 != 0 {
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
																				v183 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																				F_check_can_set_role(m, v183, v18)
																				mBase = m.M
																				v185 = m.ExcPending
																				if v185 != 0 {
																					return
																				} else {
																					v187 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																					v189 = F_object_aclcheck(m, int32(2615), v187, v18, int64(512))
																					mBase = m.M
																					v190 = m.ExcPending
																					if v190 != 0 {
																						return
																					} else {
																						if v189 == int32(0) {
																							F_AlterTypeOwner_oid(m, v88, v18)
																							mBase = m.M
																							v201 = m.ExcPending
																							if v201 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																								F_relation_close(m, v80, int32(3))
																								mBase = m.M
																								v210 = m.ExcPending
																								if v210 != 0 {
																									return
																								} else {
																									m.G0 = v76 + int32(128)
																									m.G0 = v14 + int32(16)
																									return
																								}
																							}
																						} else {
																							v194 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																							v195 = F_get_namespace_name(m, v194)
																							mBase = m.M
																							v196 = m.ExcPending
																							if v196 != 0 {
																								return
																							} else {
																								F_aclcheck_error(m, v189, int32(36), v195)
																								mBase = m.M
																								v198 = m.ExcPending
																								if v198 != 0 {
																									return
																								} else {
																									F_AlterTypeOwner_oid(m, v88, v18)
																									mBase = m.M
																									v201 = m.ExcPending
																									if v201 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																										F_relation_close(m, v80, int32(3))
																										mBase = m.M
																										v210 = m.ExcPending
																										if v210 != 0 {
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
																F_relation_close(m, v80, int32(3))
																mBase = m.M
																v210 = m.ExcPending
																if v210 != 0 {
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
													v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+79)))
													if v133 == int32(109) {
														v136 = F_get_multirange_range(m, v88)
														mBase = m.M
														v137 = m.ExcPending
														if v137 != 0 {
															return
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return
															} else {
																F_errcode(m, int32(151027844))
																mBase = m.M
																v144 = m.ExcPending
																if v144 != 0 {
																	return
																} else {
																	v145 = F_format_type_be(m, v88)
																	mBase = m.M
																	v146 = m.ExcPending
																	if v146 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v76)+32)) = v145
																		F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_12), v76+int32(32))
																		mBase = m.M
																		v152 = m.ExcPending
																		if v152 != 0 {
																			return
																		} else {
																			if v136 != 0 {
																				v153 = F_format_type_be(m, v136)
																				mBase = m.M
																				v154 = m.ExcPending
																				if v154 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v153
																					F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_13), v76+int32(16))
																					mBase = m.M
																					v160 = m.ExcPending
																					if v160 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3895), int32(_a_F_ExecAlterOwnerStmt_11))
																						mBase = m.M
																						v165 = m.ExcPending
																						if v165 != 0 {
																							return
																						} else {
																							base.Wasm_trap_unreachable()
																							for {
																							}
																						}
																					}
																				}
																			} else {
																				F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3895), int32(_a_F_ExecAlterOwnerStmt_11))
																				mBase = m.M
																				v165 = m.ExcPending
																				if v165 != 0 {
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
														v166 = *(*int32)(unsafe.Add(mBase, uint32(v96)+72))
														if v18 != v166 {
															v168 = F_superuser(m)
															mBase = m.M
															v169 = m.ExcPending
															if v169 != 0 {
																return
															} else {
																if v168 != 0 {
																	F_AlterTypeOwner_oid(m, v88, v18)
																	mBase = m.M
																	v201 = m.ExcPending
																	if v201 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																		F_relation_close(m, v80, int32(3))
																		mBase = m.M
																		v210 = m.ExcPending
																		if v210 != 0 {
																			return
																		} else {
																			m.G0 = v76 + int32(128)
																			m.G0 = v14 + int32(16)
																			return
																		}
																	}
																} else {
																	v171 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																	v173 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																	v174 = F_object_ownercheck(m, int32(1247), v171, v173)
																	mBase = m.M
																	v175 = m.ExcPending
																	if v175 != 0 {
																		return
																	} else {
																		if v174 == int32(0) {
																			v179 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																			F_aclcheck_error_type(m, int32(2), v179)
																			mBase = m.M
																			v181 = m.ExcPending
																			if v181 != 0 {
																				return
																			} else {
																				v183 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																				F_check_can_set_role(m, v183, v18)
																				mBase = m.M
																				v185 = m.ExcPending
																				if v185 != 0 {
																					return
																				} else {
																					v187 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																					v189 = F_object_aclcheck(m, int32(2615), v187, v18, int64(512))
																					mBase = m.M
																					v190 = m.ExcPending
																					if v190 != 0 {
																						return
																					} else {
																						if v189 == int32(0) {
																							F_AlterTypeOwner_oid(m, v88, v18)
																							mBase = m.M
																							v201 = m.ExcPending
																							if v201 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																								F_relation_close(m, v80, int32(3))
																								mBase = m.M
																								v210 = m.ExcPending
																								if v210 != 0 {
																									return
																								} else {
																									m.G0 = v76 + int32(128)
																									m.G0 = v14 + int32(16)
																									return
																								}
																							}
																						} else {
																							v194 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																							v195 = F_get_namespace_name(m, v194)
																							mBase = m.M
																							v196 = m.ExcPending
																							if v196 != 0 {
																								return
																							} else {
																								F_aclcheck_error(m, v189, int32(36), v195)
																								mBase = m.M
																								v198 = m.ExcPending
																								if v198 != 0 {
																									return
																								} else {
																									F_AlterTypeOwner_oid(m, v88, v18)
																									mBase = m.M
																									v201 = m.ExcPending
																									if v201 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																										F_relation_close(m, v80, int32(3))
																										mBase = m.M
																										v210 = m.ExcPending
																										if v210 != 0 {
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
																			v183 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																			F_check_can_set_role(m, v183, v18)
																			mBase = m.M
																			v185 = m.ExcPending
																			if v185 != 0 {
																				return
																			} else {
																				v187 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																				v189 = F_object_aclcheck(m, int32(2615), v187, v18, int64(512))
																				mBase = m.M
																				v190 = m.ExcPending
																				if v190 != 0 {
																					return
																				} else {
																					if v189 == int32(0) {
																						F_AlterTypeOwner_oid(m, v88, v18)
																						mBase = m.M
																						v201 = m.ExcPending
																						if v201 != 0 {
																							return
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																							F_relation_close(m, v80, int32(3))
																							mBase = m.M
																							v210 = m.ExcPending
																							if v210 != 0 {
																								return
																							} else {
																								m.G0 = v76 + int32(128)
																								m.G0 = v14 + int32(16)
																								return
																							}
																						}
																					} else {
																						v194 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																						v195 = F_get_namespace_name(m, v194)
																						mBase = m.M
																						v196 = m.ExcPending
																						if v196 != 0 {
																							return
																						} else {
																							F_aclcheck_error(m, v189, int32(36), v195)
																							mBase = m.M
																							v198 = m.ExcPending
																							if v198 != 0 {
																								return
																							} else {
																								F_AlterTypeOwner_oid(m, v88, v18)
																								mBase = m.M
																								v201 = m.ExcPending
																								if v201 != 0 {
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																									F_relation_close(m, v80, int32(3))
																									mBase = m.M
																									v210 = m.ExcPending
																									if v210 != 0 {
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
															F_relation_close(m, v80, int32(3))
															mBase = m.M
															v210 = m.ExcPending
															if v210 != 0 {
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
												v105 = m.ExcPending
												if v105 != 0 {
													return
												} else {
													F_errcode(m, int32(151027844))
													mBase = m.M
													v108 = m.ExcPending
													if v108 != 0 {
														return
													} else {
														v109 = F_format_type_be(m, v88)
														mBase = m.M
														v110 = m.ExcPending
														if v110 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v76)+80)) = v109
															F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_14), v76+int32(80))
															mBase = m.M
															v116 = m.ExcPending
															if v116 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3857), int32(_a_F_ExecAlterOwnerStmt_11))
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
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
											if v97 != int32(99) {
												v129 = *(*int32)(unsafe.Add(mBase, uint32(v96)+92))
												if v129 != 0 {
													v130 = *(*int32)(unsafe.Add(mBase, uint32(v96)+88))
													if v130 == int32(_a_F_ExecAlterOwnerStmt_7) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v262 = m.ExcPending
														if v262 != 0 {
															return
														} else {
															F_errcode(m, int32(151027844))
															mBase = m.M
															v265 = m.ExcPending
															if v265 != 0 {
																return
															} else {
																v266 = F_format_type_be(m, v88)
																mBase = m.M
																v267 = m.ExcPending
																if v267 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v76)+64)) = v266
																	F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_8), v76-int32(-64))
																	mBase = m.M
																	v273 = m.ExcPending
																	if v273 != 0 {
																		return
																	} else {
																		v274 = *(*int32)(unsafe.Add(mBase, uint32(v96)+92))
																		v275 = F_format_type_be(m, v274)
																		mBase = m.M
																		v276 = m.ExcPending
																		if v276 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v76)+48)) = v275
																			F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_9), v76+int32(48))
																			mBase = m.M
																			v282 = m.ExcPending
																			if v282 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3881), int32(_a_F_ExecAlterOwnerStmt_11))
																				mBase = m.M
																				v287 = m.ExcPending
																				if v287 != 0 {
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
														v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+79)))
														if v133 == int32(109) {
															v136 = F_get_multirange_range(m, v88)
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
																return
															} else {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v141 = m.ExcPending
																if v141 != 0 {
																	return
																} else {
																	F_errcode(m, int32(151027844))
																	mBase = m.M
																	v144 = m.ExcPending
																	if v144 != 0 {
																		return
																	} else {
																		v145 = F_format_type_be(m, v88)
																		mBase = m.M
																		v146 = m.ExcPending
																		if v146 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v76)+32)) = v145
																			F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_12), v76+int32(32))
																			mBase = m.M
																			v152 = m.ExcPending
																			if v152 != 0 {
																				return
																			} else {
																				if v136 != 0 {
																					v153 = F_format_type_be(m, v136)
																					mBase = m.M
																					v154 = m.ExcPending
																					if v154 != 0 {
																						return
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v153
																						F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_13), v76+int32(16))
																						mBase = m.M
																						v160 = m.ExcPending
																						if v160 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3895), int32(_a_F_ExecAlterOwnerStmt_11))
																							mBase = m.M
																							v165 = m.ExcPending
																							if v165 != 0 {
																								return
																							} else {
																								base.Wasm_trap_unreachable()
																								for {
																								}
																							}
																						}
																					}
																				} else {
																					F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3895), int32(_a_F_ExecAlterOwnerStmt_11))
																					mBase = m.M
																					v165 = m.ExcPending
																					if v165 != 0 {
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
															v166 = *(*int32)(unsafe.Add(mBase, uint32(v96)+72))
															if v18 != v166 {
																v168 = F_superuser(m)
																mBase = m.M
																v169 = m.ExcPending
																if v169 != 0 {
																	return
																} else {
																	if v168 != 0 {
																		F_AlterTypeOwner_oid(m, v88, v18)
																		mBase = m.M
																		v201 = m.ExcPending
																		if v201 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																			F_relation_close(m, v80, int32(3))
																			mBase = m.M
																			v210 = m.ExcPending
																			if v210 != 0 {
																				return
																			} else {
																				m.G0 = v76 + int32(128)
																				m.G0 = v14 + int32(16)
																				return
																			}
																		}
																	} else {
																		v171 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																		v173 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																		v174 = F_object_ownercheck(m, int32(1247), v171, v173)
																		mBase = m.M
																		v175 = m.ExcPending
																		if v175 != 0 {
																			return
																		} else {
																			if v174 == int32(0) {
																				v179 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																				F_aclcheck_error_type(m, int32(2), v179)
																				mBase = m.M
																				v181 = m.ExcPending
																				if v181 != 0 {
																					return
																				} else {
																					v183 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																					F_check_can_set_role(m, v183, v18)
																					mBase = m.M
																					v185 = m.ExcPending
																					if v185 != 0 {
																						return
																					} else {
																						v187 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																						v189 = F_object_aclcheck(m, int32(2615), v187, v18, int64(512))
																						mBase = m.M
																						v190 = m.ExcPending
																						if v190 != 0 {
																							return
																						} else {
																							if v189 == int32(0) {
																								F_AlterTypeOwner_oid(m, v88, v18)
																								mBase = m.M
																								v201 = m.ExcPending
																								if v201 != 0 {
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																									F_relation_close(m, v80, int32(3))
																									mBase = m.M
																									v210 = m.ExcPending
																									if v210 != 0 {
																										return
																									} else {
																										m.G0 = v76 + int32(128)
																										m.G0 = v14 + int32(16)
																										return
																									}
																								}
																							} else {
																								v194 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																								v195 = F_get_namespace_name(m, v194)
																								mBase = m.M
																								v196 = m.ExcPending
																								if v196 != 0 {
																									return
																								} else {
																									F_aclcheck_error(m, v189, int32(36), v195)
																									mBase = m.M
																									v198 = m.ExcPending
																									if v198 != 0 {
																										return
																									} else {
																										F_AlterTypeOwner_oid(m, v88, v18)
																										mBase = m.M
																										v201 = m.ExcPending
																										if v201 != 0 {
																											return
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																											F_relation_close(m, v80, int32(3))
																											mBase = m.M
																											v210 = m.ExcPending
																											if v210 != 0 {
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
																				v183 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																				F_check_can_set_role(m, v183, v18)
																				mBase = m.M
																				v185 = m.ExcPending
																				if v185 != 0 {
																					return
																				} else {
																					v187 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																					v189 = F_object_aclcheck(m, int32(2615), v187, v18, int64(512))
																					mBase = m.M
																					v190 = m.ExcPending
																					if v190 != 0 {
																						return
																					} else {
																						if v189 == int32(0) {
																							F_AlterTypeOwner_oid(m, v88, v18)
																							mBase = m.M
																							v201 = m.ExcPending
																							if v201 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																								F_relation_close(m, v80, int32(3))
																								mBase = m.M
																								v210 = m.ExcPending
																								if v210 != 0 {
																									return
																								} else {
																									m.G0 = v76 + int32(128)
																									m.G0 = v14 + int32(16)
																									return
																								}
																							}
																						} else {
																							v194 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																							v195 = F_get_namespace_name(m, v194)
																							mBase = m.M
																							v196 = m.ExcPending
																							if v196 != 0 {
																								return
																							} else {
																								F_aclcheck_error(m, v189, int32(36), v195)
																								mBase = m.M
																								v198 = m.ExcPending
																								if v198 != 0 {
																									return
																								} else {
																									F_AlterTypeOwner_oid(m, v88, v18)
																									mBase = m.M
																									v201 = m.ExcPending
																									if v201 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																										F_relation_close(m, v80, int32(3))
																										mBase = m.M
																										v210 = m.ExcPending
																										if v210 != 0 {
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
																F_relation_close(m, v80, int32(3))
																mBase = m.M
																v210 = m.ExcPending
																if v210 != 0 {
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
													v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+79)))
													if v133 == int32(109) {
														v136 = F_get_multirange_range(m, v88)
														mBase = m.M
														v137 = m.ExcPending
														if v137 != 0 {
															return
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return
															} else {
																F_errcode(m, int32(151027844))
																mBase = m.M
																v144 = m.ExcPending
																if v144 != 0 {
																	return
																} else {
																	v145 = F_format_type_be(m, v88)
																	mBase = m.M
																	v146 = m.ExcPending
																	if v146 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v76)+32)) = v145
																		F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_12), v76+int32(32))
																		mBase = m.M
																		v152 = m.ExcPending
																		if v152 != 0 {
																			return
																		} else {
																			if v136 != 0 {
																				v153 = F_format_type_be(m, v136)
																				mBase = m.M
																				v154 = m.ExcPending
																				if v154 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v153
																					F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_13), v76+int32(16))
																					mBase = m.M
																					v160 = m.ExcPending
																					if v160 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3895), int32(_a_F_ExecAlterOwnerStmt_11))
																						mBase = m.M
																						v165 = m.ExcPending
																						if v165 != 0 {
																							return
																						} else {
																							base.Wasm_trap_unreachable()
																							for {
																							}
																						}
																					}
																				}
																			} else {
																				F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3895), int32(_a_F_ExecAlterOwnerStmt_11))
																				mBase = m.M
																				v165 = m.ExcPending
																				if v165 != 0 {
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
														v166 = *(*int32)(unsafe.Add(mBase, uint32(v96)+72))
														if v18 != v166 {
															v168 = F_superuser(m)
															mBase = m.M
															v169 = m.ExcPending
															if v169 != 0 {
																return
															} else {
																if v168 != 0 {
																	F_AlterTypeOwner_oid(m, v88, v18)
																	mBase = m.M
																	v201 = m.ExcPending
																	if v201 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																		F_relation_close(m, v80, int32(3))
																		mBase = m.M
																		v210 = m.ExcPending
																		if v210 != 0 {
																			return
																		} else {
																			m.G0 = v76 + int32(128)
																			m.G0 = v14 + int32(16)
																			return
																		}
																	}
																} else {
																	v171 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																	v173 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																	v174 = F_object_ownercheck(m, int32(1247), v171, v173)
																	mBase = m.M
																	v175 = m.ExcPending
																	if v175 != 0 {
																		return
																	} else {
																		if v174 == int32(0) {
																			v179 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																			F_aclcheck_error_type(m, int32(2), v179)
																			mBase = m.M
																			v181 = m.ExcPending
																			if v181 != 0 {
																				return
																			} else {
																				v183 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																				F_check_can_set_role(m, v183, v18)
																				mBase = m.M
																				v185 = m.ExcPending
																				if v185 != 0 {
																					return
																				} else {
																					v187 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																					v189 = F_object_aclcheck(m, int32(2615), v187, v18, int64(512))
																					mBase = m.M
																					v190 = m.ExcPending
																					if v190 != 0 {
																						return
																					} else {
																						if v189 == int32(0) {
																							F_AlterTypeOwner_oid(m, v88, v18)
																							mBase = m.M
																							v201 = m.ExcPending
																							if v201 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																								F_relation_close(m, v80, int32(3))
																								mBase = m.M
																								v210 = m.ExcPending
																								if v210 != 0 {
																									return
																								} else {
																									m.G0 = v76 + int32(128)
																									m.G0 = v14 + int32(16)
																									return
																								}
																							}
																						} else {
																							v194 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																							v195 = F_get_namespace_name(m, v194)
																							mBase = m.M
																							v196 = m.ExcPending
																							if v196 != 0 {
																								return
																							} else {
																								F_aclcheck_error(m, v189, int32(36), v195)
																								mBase = m.M
																								v198 = m.ExcPending
																								if v198 != 0 {
																									return
																								} else {
																									F_AlterTypeOwner_oid(m, v88, v18)
																									mBase = m.M
																									v201 = m.ExcPending
																									if v201 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																										F_relation_close(m, v80, int32(3))
																										mBase = m.M
																										v210 = m.ExcPending
																										if v210 != 0 {
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
																			v183 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																			F_check_can_set_role(m, v183, v18)
																			mBase = m.M
																			v185 = m.ExcPending
																			if v185 != 0 {
																				return
																			} else {
																				v187 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																				v189 = F_object_aclcheck(m, int32(2615), v187, v18, int64(512))
																				mBase = m.M
																				v190 = m.ExcPending
																				if v190 != 0 {
																					return
																				} else {
																					if v189 == int32(0) {
																						F_AlterTypeOwner_oid(m, v88, v18)
																						mBase = m.M
																						v201 = m.ExcPending
																						if v201 != 0 {
																							return
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																							F_relation_close(m, v80, int32(3))
																							mBase = m.M
																							v210 = m.ExcPending
																							if v210 != 0 {
																								return
																							} else {
																								m.G0 = v76 + int32(128)
																								m.G0 = v14 + int32(16)
																								return
																							}
																						}
																					} else {
																						v194 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																						v195 = F_get_namespace_name(m, v194)
																						mBase = m.M
																						v196 = m.ExcPending
																						if v196 != 0 {
																							return
																						} else {
																							F_aclcheck_error(m, v189, int32(36), v195)
																							mBase = m.M
																							v198 = m.ExcPending
																							if v198 != 0 {
																								return
																							} else {
																								F_AlterTypeOwner_oid(m, v88, v18)
																								mBase = m.M
																								v201 = m.ExcPending
																								if v201 != 0 {
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																									F_relation_close(m, v80, int32(3))
																									mBase = m.M
																									v210 = m.ExcPending
																									if v210 != 0 {
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
															F_relation_close(m, v80, int32(3))
															mBase = m.M
															v210 = m.ExcPending
															if v210 != 0 {
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
												v124 = *(*int32)(unsafe.Add(mBase, uint32(v96)+84))
												v125 = F_get_rel_relkind(m, v124)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return
												} else {
													if v125 != int32(99) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v235 = m.ExcPending
														if v235 != 0 {
															return
														} else {
															F_errcode(m, int32(151027844))
															mBase = m.M
															v238 = m.ExcPending
															if v238 != 0 {
																return
															} else {
																v239 = F_format_type_be(m, v88)
																mBase = m.M
																v240 = m.ExcPending
																if v240 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = v239
																	F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_15), v76+int32(112))
																	mBase = m.M
																	v246 = m.ExcPending
																	if v246 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v76)+96)) = int32(_a_F_ExecAlterOwnerStmt_16)
																		F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_17), v76+int32(96))
																		mBase = m.M
																		v253 = m.ExcPending
																		if v253 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3872), int32(_a_F_ExecAlterOwnerStmt_11))
																			mBase = m.M
																			v258 = m.ExcPending
																			if v258 != 0 {
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
														v129 = *(*int32)(unsafe.Add(mBase, uint32(v96)+92))
														if v129 != 0 {
															v130 = *(*int32)(unsafe.Add(mBase, uint32(v96)+88))
															if v130 == int32(_a_F_ExecAlterOwnerStmt_7) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v262 = m.ExcPending
																if v262 != 0 {
																	return
																} else {
																	F_errcode(m, int32(151027844))
																	mBase = m.M
																	v265 = m.ExcPending
																	if v265 != 0 {
																		return
																	} else {
																		v266 = F_format_type_be(m, v88)
																		mBase = m.M
																		v267 = m.ExcPending
																		if v267 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v76)+64)) = v266
																			F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_8), v76-int32(-64))
																			mBase = m.M
																			v273 = m.ExcPending
																			if v273 != 0 {
																				return
																			} else {
																				v274 = *(*int32)(unsafe.Add(mBase, uint32(v96)+92))
																				v275 = F_format_type_be(m, v274)
																				mBase = m.M
																				v276 = m.ExcPending
																				if v276 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v76)+48)) = v275
																					F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_9), v76+int32(48))
																					mBase = m.M
																					v282 = m.ExcPending
																					if v282 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3881), int32(_a_F_ExecAlterOwnerStmt_11))
																						mBase = m.M
																						v287 = m.ExcPending
																						if v287 != 0 {
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
																v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+79)))
																if v133 == int32(109) {
																	v136 = F_get_multirange_range(m, v88)
																	mBase = m.M
																	v137 = m.ExcPending
																	if v137 != 0 {
																		return
																	} else {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v141 = m.ExcPending
																		if v141 != 0 {
																			return
																		} else {
																			F_errcode(m, int32(151027844))
																			mBase = m.M
																			v144 = m.ExcPending
																			if v144 != 0 {
																				return
																			} else {
																				v145 = F_format_type_be(m, v88)
																				mBase = m.M
																				v146 = m.ExcPending
																				if v146 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v76)+32)) = v145
																					F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_12), v76+int32(32))
																					mBase = m.M
																					v152 = m.ExcPending
																					if v152 != 0 {
																						return
																					} else {
																						if v136 != 0 {
																							v153 = F_format_type_be(m, v136)
																							mBase = m.M
																							v154 = m.ExcPending
																							if v154 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v153
																								F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_13), v76+int32(16))
																								mBase = m.M
																								v160 = m.ExcPending
																								if v160 != 0 {
																									return
																								} else {
																									F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3895), int32(_a_F_ExecAlterOwnerStmt_11))
																									mBase = m.M
																									v165 = m.ExcPending
																									if v165 != 0 {
																										return
																									} else {
																										base.Wasm_trap_unreachable()
																										for {
																										}
																									}
																								}
																							}
																						} else {
																							F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3895), int32(_a_F_ExecAlterOwnerStmt_11))
																							mBase = m.M
																							v165 = m.ExcPending
																							if v165 != 0 {
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
																	v166 = *(*int32)(unsafe.Add(mBase, uint32(v96)+72))
																	if v18 != v166 {
																		v168 = F_superuser(m)
																		mBase = m.M
																		v169 = m.ExcPending
																		if v169 != 0 {
																			return
																		} else {
																			if v168 != 0 {
																				F_AlterTypeOwner_oid(m, v88, v18)
																				mBase = m.M
																				v201 = m.ExcPending
																				if v201 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																					*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																					F_relation_close(m, v80, int32(3))
																					mBase = m.M
																					v210 = m.ExcPending
																					if v210 != 0 {
																						return
																					} else {
																						m.G0 = v76 + int32(128)
																						m.G0 = v14 + int32(16)
																						return
																					}
																				}
																			} else {
																				v171 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																				v173 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																				v174 = F_object_ownercheck(m, int32(1247), v171, v173)
																				mBase = m.M
																				v175 = m.ExcPending
																				if v175 != 0 {
																					return
																				} else {
																					if v174 == int32(0) {
																						v179 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																						F_aclcheck_error_type(m, int32(2), v179)
																						mBase = m.M
																						v181 = m.ExcPending
																						if v181 != 0 {
																							return
																						} else {
																							v183 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																							F_check_can_set_role(m, v183, v18)
																							mBase = m.M
																							v185 = m.ExcPending
																							if v185 != 0 {
																								return
																							} else {
																								v187 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																								v189 = F_object_aclcheck(m, int32(2615), v187, v18, int64(512))
																								mBase = m.M
																								v190 = m.ExcPending
																								if v190 != 0 {
																									return
																								} else {
																									if v189 == int32(0) {
																										F_AlterTypeOwner_oid(m, v88, v18)
																										mBase = m.M
																										v201 = m.ExcPending
																										if v201 != 0 {
																											return
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																											F_relation_close(m, v80, int32(3))
																											mBase = m.M
																											v210 = m.ExcPending
																											if v210 != 0 {
																												return
																											} else {
																												m.G0 = v76 + int32(128)
																												m.G0 = v14 + int32(16)
																												return
																											}
																										}
																									} else {
																										v194 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																										v195 = F_get_namespace_name(m, v194)
																										mBase = m.M
																										v196 = m.ExcPending
																										if v196 != 0 {
																											return
																										} else {
																											F_aclcheck_error(m, v189, int32(36), v195)
																											mBase = m.M
																											v198 = m.ExcPending
																											if v198 != 0 {
																												return
																											} else {
																												F_AlterTypeOwner_oid(m, v88, v18)
																												mBase = m.M
																												v201 = m.ExcPending
																												if v201 != 0 {
																													return
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																													*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																													F_relation_close(m, v80, int32(3))
																													mBase = m.M
																													v210 = m.ExcPending
																													if v210 != 0 {
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
																						v183 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																						F_check_can_set_role(m, v183, v18)
																						mBase = m.M
																						v185 = m.ExcPending
																						if v185 != 0 {
																							return
																						} else {
																							v187 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																							v189 = F_object_aclcheck(m, int32(2615), v187, v18, int64(512))
																							mBase = m.M
																							v190 = m.ExcPending
																							if v190 != 0 {
																								return
																							} else {
																								if v189 == int32(0) {
																									F_AlterTypeOwner_oid(m, v88, v18)
																									mBase = m.M
																									v201 = m.ExcPending
																									if v201 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																										F_relation_close(m, v80, int32(3))
																										mBase = m.M
																										v210 = m.ExcPending
																										if v210 != 0 {
																											return
																										} else {
																											m.G0 = v76 + int32(128)
																											m.G0 = v14 + int32(16)
																											return
																										}
																									}
																								} else {
																									v194 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																									v195 = F_get_namespace_name(m, v194)
																									mBase = m.M
																									v196 = m.ExcPending
																									if v196 != 0 {
																										return
																									} else {
																										F_aclcheck_error(m, v189, int32(36), v195)
																										mBase = m.M
																										v198 = m.ExcPending
																										if v198 != 0 {
																											return
																										} else {
																											F_AlterTypeOwner_oid(m, v88, v18)
																											mBase = m.M
																											v201 = m.ExcPending
																											if v201 != 0 {
																												return
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																												*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																												F_relation_close(m, v80, int32(3))
																												mBase = m.M
																												v210 = m.ExcPending
																												if v210 != 0 {
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
																		F_relation_close(m, v80, int32(3))
																		mBase = m.M
																		v210 = m.ExcPending
																		if v210 != 0 {
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
															v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+79)))
															if v133 == int32(109) {
																v136 = F_get_multirange_range(m, v88)
																mBase = m.M
																v137 = m.ExcPending
																if v137 != 0 {
																	return
																} else {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v141 = m.ExcPending
																	if v141 != 0 {
																		return
																	} else {
																		F_errcode(m, int32(151027844))
																		mBase = m.M
																		v144 = m.ExcPending
																		if v144 != 0 {
																			return
																		} else {
																			v145 = F_format_type_be(m, v88)
																			mBase = m.M
																			v146 = m.ExcPending
																			if v146 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v76)+32)) = v145
																				F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_12), v76+int32(32))
																				mBase = m.M
																				v152 = m.ExcPending
																				if v152 != 0 {
																					return
																				} else {
																					if v136 != 0 {
																						v153 = F_format_type_be(m, v136)
																						mBase = m.M
																						v154 = m.ExcPending
																						if v154 != 0 {
																							return
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v153
																							F_errhint(m, int32(_a_F_ExecAlterOwnerStmt_13), v76+int32(16))
																							mBase = m.M
																							v160 = m.ExcPending
																							if v160 != 0 {
																								return
																							} else {
																								F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3895), int32(_a_F_ExecAlterOwnerStmt_11))
																								mBase = m.M
																								v165 = m.ExcPending
																								if v165 != 0 {
																									return
																								} else {
																									base.Wasm_trap_unreachable()
																									for {
																									}
																								}
																							}
																						}
																					} else {
																						F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3895), int32(_a_F_ExecAlterOwnerStmt_11))
																						mBase = m.M
																						v165 = m.ExcPending
																						if v165 != 0 {
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
																v166 = *(*int32)(unsafe.Add(mBase, uint32(v96)+72))
																if v18 != v166 {
																	v168 = F_superuser(m)
																	mBase = m.M
																	v169 = m.ExcPending
																	if v169 != 0 {
																		return
																	} else {
																		if v168 != 0 {
																			F_AlterTypeOwner_oid(m, v88, v18)
																			mBase = m.M
																			v201 = m.ExcPending
																			if v201 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																				F_relation_close(m, v80, int32(3))
																				mBase = m.M
																				v210 = m.ExcPending
																				if v210 != 0 {
																					return
																				} else {
																					m.G0 = v76 + int32(128)
																					m.G0 = v14 + int32(16)
																					return
																				}
																			}
																		} else {
																			v171 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																			v173 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																			v174 = F_object_ownercheck(m, int32(1247), v171, v173)
																			mBase = m.M
																			v175 = m.ExcPending
																			if v175 != 0 {
																				return
																			} else {
																				if v174 == int32(0) {
																					v179 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
																					F_aclcheck_error_type(m, int32(2), v179)
																					mBase = m.M
																					v181 = m.ExcPending
																					if v181 != 0 {
																						return
																					} else {
																						v183 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																						F_check_can_set_role(m, v183, v18)
																						mBase = m.M
																						v185 = m.ExcPending
																						if v185 != 0 {
																							return
																						} else {
																							v187 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																							v189 = F_object_aclcheck(m, int32(2615), v187, v18, int64(512))
																							mBase = m.M
																							v190 = m.ExcPending
																							if v190 != 0 {
																								return
																							} else {
																								if v189 == int32(0) {
																									F_AlterTypeOwner_oid(m, v88, v18)
																									mBase = m.M
																									v201 = m.ExcPending
																									if v201 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																										F_relation_close(m, v80, int32(3))
																										mBase = m.M
																										v210 = m.ExcPending
																										if v210 != 0 {
																											return
																										} else {
																											m.G0 = v76 + int32(128)
																											m.G0 = v14 + int32(16)
																											return
																										}
																									}
																								} else {
																									v194 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																									v195 = F_get_namespace_name(m, v194)
																									mBase = m.M
																									v196 = m.ExcPending
																									if v196 != 0 {
																										return
																									} else {
																										F_aclcheck_error(m, v189, int32(36), v195)
																										mBase = m.M
																										v198 = m.ExcPending
																										if v198 != 0 {
																											return
																										} else {
																											F_AlterTypeOwner_oid(m, v88, v18)
																											mBase = m.M
																											v201 = m.ExcPending
																											if v201 != 0 {
																												return
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																												*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																												F_relation_close(m, v80, int32(3))
																												mBase = m.M
																												v210 = m.ExcPending
																												if v210 != 0 {
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
																					v183 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[0]))
																					F_check_can_set_role(m, v183, v18)
																					mBase = m.M
																					v185 = m.ExcPending
																					if v185 != 0 {
																						return
																					} else {
																						v187 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																						v189 = F_object_aclcheck(m, int32(2615), v187, v18, int64(512))
																						mBase = m.M
																						v190 = m.ExcPending
																						if v190 != 0 {
																							return
																						} else {
																							if v189 == int32(0) {
																								F_AlterTypeOwner_oid(m, v88, v18)
																								mBase = m.M
																								v201 = m.ExcPending
																								if v201 != 0 {
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																									F_relation_close(m, v80, int32(3))
																									mBase = m.M
																									v210 = m.ExcPending
																									if v210 != 0 {
																										return
																									} else {
																										m.G0 = v76 + int32(128)
																										m.G0 = v14 + int32(16)
																										return
																									}
																								}
																							} else {
																								v194 = *(*int32)(unsafe.Add(mBase, uint32(v96)+68))
																								v195 = F_get_namespace_name(m, v194)
																								mBase = m.M
																								v196 = m.ExcPending
																								if v196 != 0 {
																									return
																								} else {
																									F_aclcheck_error(m, v189, int32(36), v195)
																									mBase = m.M
																									v198 = m.ExcPending
																									if v198 != 0 {
																										return
																									} else {
																										F_AlterTypeOwner_oid(m, v88, v18)
																										mBase = m.M
																										v201 = m.ExcPending
																										if v201 != 0 {
																											return
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
																											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
																											*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
																											F_relation_close(m, v80, int32(3))
																											mBase = m.M
																											v210 = m.ExcPending
																											if v210 != 0 {
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
																	F_relation_close(m, v80, int32(3))
																	mBase = m.M
																	v210 = m.ExcPending
																	if v210 != 0 {
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
							v217 = m.ExcPending
							if v217 != 0 {
								return
							} else {
								F_errcode(m, int32(67137668))
								mBase = m.M
								v220 = m.ExcPending
								if v220 != 0 {
									return
								} else {
									v221 = F_TypeNameToString(m, v83)
									mBase = m.M
									v222 = m.ExcPending
									if v222 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v76))) = v221
										F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_18), v76)
										mBase = m.M
										v226 = m.ExcPending
										if v226 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_10), int32(3843), int32(_a_F_ExecAlterOwnerStmt_11))
											mBase = m.M
											v231 = m.ExcPending
											if v231 != 0 {
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
			v390 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
			v392 = m.G0
			v394 = v392 - int32(16)
			m.G0 = v394
			v398 = F_table_open(m, int32(3466), int32(3))
			mBase = m.M
			v399 = m.ExcPending
			if v399 != 0 {
				return
			} else {
				v402 = F_SearchSysCacheCopy(m, int32(25), v391, int32(0))
				mBase = m.M
				v403 = m.ExcPending
				if v403 != 0 {
					return
				} else {
					if v402 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v409 = m.ExcPending
						if v409 != 0 {
							return
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v412 = m.ExcPending
							if v412 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v394))) = v391
								F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_19), v394)
								mBase = m.M
								v416 = m.ExcPending
								if v416 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_20), int32(494), int32(_a_F_ExecAlterOwnerStmt_21))
									mBase = m.M
									v421 = m.ExcPending
									if v421 != 0 {
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
						v422 = *(*int32)(unsafe.Add(mBase, uint32(v402)+16))
						v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422)+22)))
						v425 = *(*int32)(unsafe.Add(mBase, uint32(v422+v423)))
						F_AlterEventTriggerOwner_internal(m, v398, v402, v18)
						mBase = m.M
						v427 = m.ExcPending
						if v427 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v425
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(3466)
							F_pfree(m, v402)
							mBase = m.M
							v434 = m.ExcPending
							if v434 != 0 {
								return
							} else {
								F_relation_close(m, v398, int32(3))
								mBase = m.M
								v437 = m.ExcPending
								if v437 != 0 {
									return
								} else {
									m.G0 = v394 + int32(16)
									m.G0 = v14 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		case 15:
			v288 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
			v290 = m.G0
			v292 = v290 - int32(16)
			m.G0 = v292
			v296 = F_table_open(m, int32(2328), int32(3))
			mBase = m.M
			v297 = m.ExcPending
			if v297 != 0 {
				return
			} else {
				v300 = F_SearchSysCacheCopy(m, int32(29), v289, int32(0))
				mBase = m.M
				v301 = m.ExcPending
				if v301 != 0 {
					return
				} else {
					if v300 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v307 = m.ExcPending
						if v307 != 0 {
							return
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v310 = m.ExcPending
							if v310 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v292))) = v289
								F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_22), v292)
								mBase = m.M
								v314 = m.ExcPending
								if v314 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_23), int32(302), int32(_a_F_ExecAlterOwnerStmt_24))
									mBase = m.M
									v319 = m.ExcPending
									if v319 != 0 {
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
						v320 = *(*int32)(unsafe.Add(mBase, uint32(v300)+16))
						v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+22)))
						v323 = *(*int32)(unsafe.Add(mBase, uint32(v320+v321)))
						F_AlterForeignDataWrapperOwner_internal(m, v296, v300, v18)
						mBase = m.M
						v325 = m.ExcPending
						if v325 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2328)
							F_pfree(m, v300)
							mBase = m.M
							v332 = m.ExcPending
							if v332 != 0 {
								return
							} else {
								F_relation_close(m, v296, int32(3))
								mBase = m.M
								v335 = m.ExcPending
								if v335 != 0 {
									return
								} else {
									m.G0 = v292 + int32(16)
									m.G0 = v14 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		case 16:
			v339 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
			v341 = m.G0
			v343 = v341 - int32(16)
			m.G0 = v343
			v347 = F_table_open(m, int32(1417), int32(3))
			mBase = m.M
			v348 = m.ExcPending
			if v348 != 0 {
				return
			} else {
				v351 = F_SearchSysCacheCopy(m, int32(31), v340, int32(0))
				mBase = m.M
				v352 = m.ExcPending
				if v352 != 0 {
					return
				} else {
					if v351 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v358 = m.ExcPending
						if v358 != 0 {
							return
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v361 = m.ExcPending
							if v361 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v343))) = v340
								F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_25), v343)
								mBase = m.M
								v365 = m.ExcPending
								if v365 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_23), int32(441), int32(_a_F_ExecAlterOwnerStmt_26))
									mBase = m.M
									v370 = m.ExcPending
									if v370 != 0 {
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
						v371 = *(*int32)(unsafe.Add(mBase, uint32(v351)+16))
						v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371)+22)))
						v374 = *(*int32)(unsafe.Add(mBase, uint32(v371+v372)))
						F_AlterForeignServerOwner_internal(m, v347, v351, v18)
						mBase = m.M
						v376 = m.ExcPending
						if v376 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v374
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1417)
							F_pfree(m, v351)
							mBase = m.M
							v383 = m.ExcPending
							if v383 != 0 {
								return
							} else {
								F_relation_close(m, v347, int32(3))
								mBase = m.M
								v386 = m.ExcPending
								if v386 != 0 {
									return
								} else {
									m.G0 = v343 + int32(16)
									m.G0 = v14 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		case 29:
			v441 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)+4))
			v443 = m.G0
			v445 = v443 - int32(16)
			m.G0 = v445
			v449 = F_table_open(m, int32(_a_F_ExecAlterOwnerStmt_27), int32(3))
			mBase = m.M
			v450 = m.ExcPending
			if v450 != 0 {
				return
			} else {
				v453 = F_SearchSysCacheCopy(m, int32(48), v442, int32(0))
				mBase = m.M
				v454 = m.ExcPending
				if v454 != 0 {
					return
				} else {
					if v453 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v460 = m.ExcPending
						if v460 != 0 {
							return
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v463 = m.ExcPending
							if v463 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v445))) = v442
								F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_28), v445)
								mBase = m.M
								v467 = m.ExcPending
								if v467 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_29), int32(2073), int32(_a_F_ExecAlterOwnerStmt_30))
									mBase = m.M
									v472 = m.ExcPending
									if v472 != 0 {
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
						v473 = *(*int32)(unsafe.Add(mBase, uint32(v453)+16))
						v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473)+22)))
						v476 = *(*int32)(unsafe.Add(mBase, uint32(v473+v474)))
						F_AlterPublicationOwner_internal(m, v449, v453, v18)
						mBase = m.M
						v478 = m.ExcPending
						if v478 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v476
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a_F_ExecAlterOwnerStmt_27)
							F_pfree(m, v453)
							mBase = m.M
							v485 = m.ExcPending
							if v485 != 0 {
								return
							} else {
								F_relation_close(m, v449, int32(3))
								mBase = m.M
								v488 = m.ExcPending
								if v488 != 0 {
									return
								} else {
									m.G0 = v445 + int32(16)
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
								F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_31), v27)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_32), int32(344), int32(_a_F_ExecAlterOwnerStmt_33))
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
								F_relation_close(m, v31, int32(3))
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
			v492 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+4))
			v494 = m.G0
			v496 = v494 - int32(16)
			m.G0 = v496
			v500 = F_table_open(m, int32(_a_F_ExecAlterOwnerStmt_34), int32(3))
			mBase = m.M
			v501 = m.ExcPending
			if v501 != 0 {
				return
			} else {
				v504 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterOwnerStmt[2]))
				v505 = F_SearchSysCacheCopy(m, int32(66), v504, v493)
				mBase = m.M
				v506 = m.ExcPending
				if v506 != 0 {
					return
				} else {
					if v505 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v512 = m.ExcPending
						if v512 != 0 {
							return
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v515 = m.ExcPending
							if v515 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v496))) = v493
								F_errmsg(m, int32(_a_F_ExecAlterOwnerStmt_35), v496)
								mBase = m.M
								v519 = m.ExcPending
								if v519 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ExecAlterOwnerStmt_36), int32(2046), int32(_a_F_ExecAlterOwnerStmt_37))
									mBase = m.M
									v524 = m.ExcPending
									if v524 != 0 {
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
						v525 = *(*int32)(unsafe.Add(mBase, uint32(v505)+16))
						v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525)+22)))
						v528 = *(*int32)(unsafe.Add(mBase, uint32(v525+v526)))
						F_AlterSubscriptionOwner_internal(m, v500, v505, v18)
						mBase = m.M
						v530 = m.ExcPending
						if v530 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v528
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a_F_ExecAlterOwnerStmt_34)
							F_pfree(m, v505)
							mBase = m.M
							v537 = m.ExcPending
							if v537 != 0 {
								return
							} else {
								F_relation_close(m, v500, int32(3))
								mBase = m.M
								v540 = m.ExcPending
								if v540 != 0 {
									return
								} else {
									m.G0 = v496 + int32(16)
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
