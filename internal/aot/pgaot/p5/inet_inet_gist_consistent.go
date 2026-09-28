package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inet_gist_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
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
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
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
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
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
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v722 int32
	_ = v722
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v760 int64
	_ = v760
	var v762 int32
	_ = v762
	var v770 int32
	_ = v770
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
		v19 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v19)
		v21 = int64(1)
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
		if v22 == v19 {
			v760 = v21
			return v760
		} else {
			v25 = int32(1)
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v27&v25 != 0 {
				v30 = v25
			} else {
				v30 = int32(4)
			}
			v31 = v12 + v30
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
			if v32 != v22 {
				switch v16 - int32(19) {
				case 0:
					v760 = v21
					return v760
				case 1, 2:
					if base.Ui32(v32) <= base.Ui32(v22) {
						return int64(0)
					} else {
						v760 = v21
						return v760
					}
				case 3, 4:
					if base.Ui32(v32) < base.Ui32(v22) {
						v760 = v21
						return v760
					} else {
						return int64(0)
					}
				default:
					return int64(0)
				}
			} else {
				switch v16 - int32(18) {
				case 0, 9:
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
					if base.Ui32(v61) <= base.Ui32(v62) {
						v68 = v17 + int32(4)
						v70 = v31 + int32(2)
						v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+3)))
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
						if base.Ui32(v71) < base.Ui32(v72) {
							v74 = v71
						} else {
							v74 = v72
						}
						v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
						if base.Ui32(v74) < base.Ui32(v75) {
							v77 = v74
						} else {
							v77 = v75
						}
						v82 = base.I32_div_s(v77, int32(8))
						v83 = F_memcmp(m, v68, v70, v82)
						mBase = m.M
						if v83 != 0 {
							v169 = v83
							v179 = v169
						} else {
							v84 = int32(0)
							v87 = v77 - v82<<(uint(int32(3))%32)
							if v87 <= v84 {
								v169 = v84
								v179 = v169
							} else {
								v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v82))))
								v92 = int32(128)
								v93 = v91 & v92
								v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v82))))
								if v93 != v95&v92 {
									v170 = v93
									if v170 != 0 {
										v173 = int32(1)
									} else {
										v173 = int32(-1)
									}
									v179 = v173
								} else {
									if v87 == int32(1) {
										v169 = v84
										v179 = v169
									} else {
										v101 = int32(1)
										v103 = int32(128)
										v104 = v91 << (uint(v101) % 32) & v103
										if v104 != v95<<(uint(v101)%32)&v103 {
											v170 = v104
											if v170 != 0 {
												v173 = int32(1)
											} else {
												v173 = int32(-1)
											}
											v179 = v173
										} else {
											if v87 < int32(3) {
												v169 = v84
												v179 = v169
											} else {
												v112 = int32(2)
												v114 = int32(128)
												v115 = v91 << (uint(v112) % 32) & v114
												if v115 != v95<<(uint(v112)%32)&v114 {
													v170 = v115
													if v170 != 0 {
														v173 = int32(1)
													} else {
														v173 = int32(-1)
													}
													v179 = v173
												} else {
													if v87 == int32(3) {
														v169 = v84
														v179 = v169
													} else {
														v123 = int32(3)
														v125 = int32(128)
														v126 = v91 << (uint(v123) % 32) & v125
														if v126 != v95<<(uint(v123)%32)&v125 {
															v170 = v126
															if v170 != 0 {
																v173 = int32(1)
															} else {
																v173 = int32(-1)
															}
															v179 = v173
														} else {
															if v87 < int32(5) {
																v169 = v84
																v179 = v169
															} else {
																v134 = int32(4)
																v136 = int32(128)
																v137 = v91 << (uint(v134) % 32) & v136
																if v137 != v95<<(uint(v134)%32)&v136 {
																	v170 = v137
																	if v170 != 0 {
																		v173 = int32(1)
																	} else {
																		v173 = int32(-1)
																	}
																	v179 = v173
																} else {
																	if v87 == int32(5) {
																		v169 = v84
																		v179 = v169
																	} else {
																		v145 = int32(5)
																		v147 = int32(128)
																		v148 = v91 << (uint(v145) % 32) & v147
																		if v148 != v95<<(uint(v145)%32)&v147 {
																			v170 = v148
																			if v170 != 0 {
																				v173 = int32(1)
																			} else {
																				v173 = int32(-1)
																			}
																			v179 = v173
																		} else {
																			if v87 < int32(7) {
																				v169 = v84
																				v179 = v169
																			} else {
																				v156 = int32(6)
																				v158 = int32(128)
																				v159 = v91 << (uint(v156) % 32) & v158
																				if v159 != v95<<(uint(v156)%32)&v158 {
																					v170 = v159
																					if v170 != 0 {
																						v173 = int32(1)
																					} else {
																						v173 = int32(-1)
																					}
																					v179 = v173
																				} else {
																					v169 = v84
																					v179 = v169
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
						switch v16 - int32(3) {
						case 0, 21, 22, 23, 24:
							v296 = v179
							return base.I64_extend_i32_u(base.B2i32(v296 == int32(0)))
						default:
							v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
							v606 = v603 & int32(1)
							if v606 != 0 {
								v611 = int32(1)
							} else {
								v611 = int32(4)
							}
							v614 = v12 + v611 + int32(2)
							v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
							if v617 == int32(3) {
								v620 = int32(128)
							} else {
								v620 = int32(32)
							}
							v625 = base.I32_div_s(v620, int32(8))
							v626 = F_memcmp(m, v68, v614, v625)
							mBase = m.M
							if v626 != 0 {
								v712 = v626
								v722 = v712
							} else {
								v627 = int32(0)
								v630 = v620 - v625<<(uint(int32(3))%32)
								if v630 <= v627 {
									v712 = v627
									v722 = v712
								} else {
									v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v625))))
									v635 = int32(128)
									v636 = v634 & v635
									v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614+v625))))
									if v636 != v638&v635 {
										v713 = v636
										if v713 != 0 {
											v716 = int32(1)
										} else {
											v716 = int32(-1)
										}
										v722 = v716
									} else {
										if v630 == int32(1) {
											v712 = v627
											v722 = v712
										} else {
											v644 = int32(1)
											v646 = int32(128)
											v647 = v634 << (uint(v644) % 32) & v646
											if v647 != v638<<(uint(v644)%32)&v646 {
												v713 = v647
												if v713 != 0 {
													v716 = int32(1)
												} else {
													v716 = int32(-1)
												}
												v722 = v716
											} else {
												if v630 < int32(3) {
													v712 = v627
													v722 = v712
												} else {
													v655 = int32(2)
													v657 = int32(128)
													v658 = v634 << (uint(v655) % 32) & v657
													if v658 != v638<<(uint(v655)%32)&v657 {
														v713 = v658
														if v713 != 0 {
															v716 = int32(1)
														} else {
															v716 = int32(-1)
														}
														v722 = v716
													} else {
														if v630 == int32(3) {
															v712 = v627
															v722 = v712
														} else {
															v666 = int32(3)
															v668 = int32(128)
															v669 = v634 << (uint(v666) % 32) & v668
															if v669 != v638<<(uint(v666)%32)&v668 {
																v713 = v669
																if v713 != 0 {
																	v716 = int32(1)
																} else {
																	v716 = int32(-1)
																}
																v722 = v716
															} else {
																if v630 < int32(5) {
																	v712 = v627
																	v722 = v712
																} else {
																	v677 = int32(4)
																	v679 = int32(128)
																	v680 = v634 << (uint(v677) % 32) & v679
																	if v680 != v638<<(uint(v677)%32)&v679 {
																		v713 = v680
																		if v713 != 0 {
																			v716 = int32(1)
																		} else {
																			v716 = int32(-1)
																		}
																		v722 = v716
																	} else {
																		if v630 == int32(5) {
																			v712 = v627
																			v722 = v712
																		} else {
																			v688 = int32(5)
																			v690 = int32(128)
																			v691 = v634 << (uint(v688) % 32) & v690
																			if v691 != v638<<(uint(v688)%32)&v690 {
																				v713 = v691
																				if v713 != 0 {
																					v716 = int32(1)
																				} else {
																					v716 = int32(-1)
																				}
																				v722 = v716
																			} else {
																				if v630 < int32(7) {
																					v712 = v627
																					v722 = v712
																				} else {
																					v699 = int32(6)
																					v701 = int32(128)
																					v702 = v634 << (uint(v699) % 32) & v701
																					if v702 != v638<<(uint(v699)%32)&v701 {
																						v713 = v702
																						if v713 != 0 {
																							v716 = int32(1)
																						} else {
																							v716 = int32(-1)
																						}
																						v722 = v716
																					} else {
																						v712 = v627
																						v722 = v712
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
							switch v16 - int32(18) {
							case 0:
								v762 = v722
								return base.I64_extend_i32_u(base.B2i32(v762 == int32(0)))
							case 1:
								v770 = v722
								return base.I64_extend_i32_u(base.B2i32(v770 != int32(0)))
							case 2:
								return base.I64_extend_i32_u(int32(base.Ui32(v722) >> (uint(int32(31)) % 32)))
							case 3:
								return base.I64_extend_i32_u(base.B2i32(v722 <= int32(0)))
							case 4:
								return base.I64_extend_i32_u(base.B2i32(int32(0) < v722))
							case 5:
								return base.I64_extend_i32_u(base.B2i32(int32(0) <= v722))
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v744 = m.ExcPending
								if v744 != 0 {
									return int64(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_0), int32(0))
									mBase = m.M
									v748 = m.ExcPending
									if v748 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_inet_gist_consistent_1), int32(328), int32(_a_F_inet_gist_consistent_2))
										mBase = m.M
										v753 = m.ExcPending
										if v753 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						case 15:
							if v179 == int32(0) {
								v472 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
								v473 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v472)+16)))
								v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472+v473)+12)))
								if v475&int32(1) == int32(0) {
									v760 = v21
									return v760
								} else {
									v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
									v482 = int32(1)
									v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
									if v484&v482 != 0 {
										v487 = v482
									} else {
										v487 = int32(4)
									}
									v488 = v12 + v487
									v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+1)))
									if v481 != v489 {
										v760 = int64(0)
										return v760
									} else {
										v492 = v488 + int32(2)
										v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
										if v495 == int32(3) {
											v498 = int32(128)
										} else {
											v498 = int32(32)
										}
										v503 = base.I32_div_s(v498, int32(8))
										v504 = F_memcmp(m, v68, v492, v503)
										mBase = m.M
										if v504 != 0 {
											v590 = v504
											v600 = v590
										} else {
											v505 = int32(0)
											v508 = v498 - v503<<(uint(int32(3))%32)
											if v508 <= v505 {
												v590 = v505
												v600 = v590
											} else {
												v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v503))))
												v513 = int32(128)
												v514 = v512 & v513
												v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492+v503))))
												if v514 != v516&v513 {
													v591 = v514
													if v591 != 0 {
														v594 = int32(1)
													} else {
														v594 = int32(-1)
													}
													v600 = v594
												} else {
													if v508 == int32(1) {
														v590 = v505
														v600 = v590
													} else {
														v522 = int32(1)
														v524 = int32(128)
														v525 = v512 << (uint(v522) % 32) & v524
														if v525 != v516<<(uint(v522)%32)&v524 {
															v591 = v525
															if v591 != 0 {
																v594 = int32(1)
															} else {
																v594 = int32(-1)
															}
															v600 = v594
														} else {
															if v508 < int32(3) {
																v590 = v505
																v600 = v590
															} else {
																v533 = int32(2)
																v535 = int32(128)
																v536 = v512 << (uint(v533) % 32) & v535
																if v536 != v516<<(uint(v533)%32)&v535 {
																	v591 = v536
																	if v591 != 0 {
																		v594 = int32(1)
																	} else {
																		v594 = int32(-1)
																	}
																	v600 = v594
																} else {
																	if v508 == int32(3) {
																		v590 = v505
																		v600 = v590
																	} else {
																		v544 = int32(3)
																		v546 = int32(128)
																		v547 = v512 << (uint(v544) % 32) & v546
																		if v547 != v516<<(uint(v544)%32)&v546 {
																			v591 = v547
																			if v591 != 0 {
																				v594 = int32(1)
																			} else {
																				v594 = int32(-1)
																			}
																			v600 = v594
																		} else {
																			if v508 < int32(5) {
																				v590 = v505
																				v600 = v590
																			} else {
																				v555 = int32(4)
																				v557 = int32(128)
																				v558 = v512 << (uint(v555) % 32) & v557
																				if v558 != v516<<(uint(v555)%32)&v557 {
																					v591 = v558
																					if v591 != 0 {
																						v594 = int32(1)
																					} else {
																						v594 = int32(-1)
																					}
																					v600 = v594
																				} else {
																					if v508 == int32(5) {
																						v590 = v505
																						v600 = v590
																					} else {
																						v566 = int32(5)
																						v568 = int32(128)
																						v569 = v512 << (uint(v566) % 32) & v568
																						if v569 != v516<<(uint(v566)%32)&v568 {
																							v591 = v569
																							if v591 != 0 {
																								v594 = int32(1)
																							} else {
																								v594 = int32(-1)
																							}
																							v600 = v594
																						} else {
																							if v508 < int32(7) {
																								v590 = v505
																								v600 = v590
																							} else {
																								v577 = int32(6)
																								v579 = int32(128)
																								v580 = v512 << (uint(v577) % 32) & v579
																								if v580 != v516<<(uint(v577)%32)&v579 {
																									v591 = v580
																									if v591 != 0 {
																										v594 = int32(1)
																									} else {
																										v594 = int32(-1)
																									}
																									v600 = v594
																								} else {
																									v590 = v505
																									v600 = v590
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
										v762 = v600
										return base.I64_extend_i32_u(base.B2i32(v762 == int32(0)))
									}
								}
							} else {
								return int64(0)
							}
						case 16:
							if v179 != 0 {
								v760 = v21
								return v760
							} else {
								v344 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
								v345 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v344)+16)))
								v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344+v345)+12)))
								if v347&int32(1) == int32(0) {
									v760 = v21
									return v760
								} else {
									v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
									v353 = int32(1)
									v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
									if v355&v353 != 0 {
										v358 = v353
									} else {
										v358 = int32(4)
									}
									v359 = v12 + v358
									v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359)+1)))
									if v352 != v360 {
										v760 = v21
										return v760
									} else {
										v363 = v359 + int32(2)
										v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
										if v366 == int32(3) {
											v369 = int32(128)
										} else {
											v369 = int32(32)
										}
										v374 = base.I32_div_s(v369, int32(8))
										v375 = F_memcmp(m, v68, v363, v374)
										mBase = m.M
										if v375 != 0 {
											v461 = v375
											v471 = v461
										} else {
											v376 = int32(0)
											v379 = v369 - v374<<(uint(int32(3))%32)
											if v379 <= v376 {
												v461 = v376
												v471 = v461
											} else {
												v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v374))))
												v384 = int32(128)
												v385 = v383 & v384
												v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+v374))))
												if v385 != v387&v384 {
													v462 = v385
													if v462 != 0 {
														v465 = int32(1)
													} else {
														v465 = int32(-1)
													}
													v471 = v465
												} else {
													if v379 == int32(1) {
														v461 = v376
														v471 = v461
													} else {
														v393 = int32(1)
														v395 = int32(128)
														v396 = v383 << (uint(v393) % 32) & v395
														if v396 != v387<<(uint(v393)%32)&v395 {
															v462 = v396
															if v462 != 0 {
																v465 = int32(1)
															} else {
																v465 = int32(-1)
															}
															v471 = v465
														} else {
															if v379 < int32(3) {
																v461 = v376
																v471 = v461
															} else {
																v404 = int32(2)
																v406 = int32(128)
																v407 = v383 << (uint(v404) % 32) & v406
																if v407 != v387<<(uint(v404)%32)&v406 {
																	v462 = v407
																	if v462 != 0 {
																		v465 = int32(1)
																	} else {
																		v465 = int32(-1)
																	}
																	v471 = v465
																} else {
																	if v379 == int32(3) {
																		v461 = v376
																		v471 = v461
																	} else {
																		v415 = int32(3)
																		v417 = int32(128)
																		v418 = v383 << (uint(v415) % 32) & v417
																		if v418 != v387<<(uint(v415)%32)&v417 {
																			v462 = v418
																			if v462 != 0 {
																				v465 = int32(1)
																			} else {
																				v465 = int32(-1)
																			}
																			v471 = v465
																		} else {
																			if v379 < int32(5) {
																				v461 = v376
																				v471 = v461
																			} else {
																				v426 = int32(4)
																				v428 = int32(128)
																				v429 = v383 << (uint(v426) % 32) & v428
																				if v429 != v387<<(uint(v426)%32)&v428 {
																					v462 = v429
																					if v462 != 0 {
																						v465 = int32(1)
																					} else {
																						v465 = int32(-1)
																					}
																					v471 = v465
																				} else {
																					if v379 == int32(5) {
																						v461 = v376
																						v471 = v461
																					} else {
																						v437 = int32(5)
																						v439 = int32(128)
																						v440 = v383 << (uint(v437) % 32) & v439
																						if v440 != v387<<(uint(v437)%32)&v439 {
																							v462 = v440
																							if v462 != 0 {
																								v465 = int32(1)
																							} else {
																								v465 = int32(-1)
																							}
																							v471 = v465
																						} else {
																							if v379 < int32(7) {
																								v461 = v376
																								v471 = v461
																							} else {
																								v448 = int32(6)
																								v450 = int32(128)
																								v451 = v383 << (uint(v448) % 32) & v450
																								if v451 != v387<<(uint(v448)%32)&v450 {
																									v462 = v451
																									if v462 != 0 {
																										v465 = int32(1)
																									} else {
																										v465 = int32(-1)
																									}
																									v471 = v465
																								} else {
																									v461 = v376
																									v471 = v461
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
										v770 = v471
										return base.I64_extend_i32_u(base.B2i32(v770 != int32(0)))
									}
								}
							}
						case 17, 18:
							if int32(0) < v179 {
								return int64(0)
							} else {
								if v179 < int32(0) {
									v760 = v21
									return v760
								} else {
									v310 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
									v311 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v310)+16)))
									v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310+v311)+12)))
									if v313&int32(1) != 0 {
										v329 = int32(1)
										v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
										v333 = v331 & v329
										if v333 != 0 {
											v334 = v329
										} else {
											v334 = int32(4)
										}
										v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v334)+1)))
										v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
										if v16&int32(_a_F_inet_gist_consistent_3) != int32(20) {
											if base.Ui32(v336) < base.Ui32(v337) {
												v760 = v21
												return v760
											} else {
												if base.Ui32(v336) <= base.Ui32(v337) {
													v606 = v333
													if v606 != 0 {
														v611 = int32(1)
													} else {
														v611 = int32(4)
													}
													v614 = v12 + v611 + int32(2)
													v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
													if v617 == int32(3) {
														v620 = int32(128)
													} else {
														v620 = int32(32)
													}
													v625 = base.I32_div_s(v620, int32(8))
													v626 = F_memcmp(m, v68, v614, v625)
													mBase = m.M
													if v626 != 0 {
														v712 = v626
														v722 = v712
													} else {
														v627 = int32(0)
														v630 = v620 - v625<<(uint(int32(3))%32)
														if v630 <= v627 {
															v712 = v627
															v722 = v712
														} else {
															v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v625))))
															v635 = int32(128)
															v636 = v634 & v635
															v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614+v625))))
															if v636 != v638&v635 {
																v713 = v636
																if v713 != 0 {
																	v716 = int32(1)
																} else {
																	v716 = int32(-1)
																}
																v722 = v716
															} else {
																if v630 == int32(1) {
																	v712 = v627
																	v722 = v712
																} else {
																	v644 = int32(1)
																	v646 = int32(128)
																	v647 = v634 << (uint(v644) % 32) & v646
																	if v647 != v638<<(uint(v644)%32)&v646 {
																		v713 = v647
																		if v713 != 0 {
																			v716 = int32(1)
																		} else {
																			v716 = int32(-1)
																		}
																		v722 = v716
																	} else {
																		if v630 < int32(3) {
																			v712 = v627
																			v722 = v712
																		} else {
																			v655 = int32(2)
																			v657 = int32(128)
																			v658 = v634 << (uint(v655) % 32) & v657
																			if v658 != v638<<(uint(v655)%32)&v657 {
																				v713 = v658
																				if v713 != 0 {
																					v716 = int32(1)
																				} else {
																					v716 = int32(-1)
																				}
																				v722 = v716
																			} else {
																				if v630 == int32(3) {
																					v712 = v627
																					v722 = v712
																				} else {
																					v666 = int32(3)
																					v668 = int32(128)
																					v669 = v634 << (uint(v666) % 32) & v668
																					if v669 != v638<<(uint(v666)%32)&v668 {
																						v713 = v669
																						if v713 != 0 {
																							v716 = int32(1)
																						} else {
																							v716 = int32(-1)
																						}
																						v722 = v716
																					} else {
																						if v630 < int32(5) {
																							v712 = v627
																							v722 = v712
																						} else {
																							v677 = int32(4)
																							v679 = int32(128)
																							v680 = v634 << (uint(v677) % 32) & v679
																							if v680 != v638<<(uint(v677)%32)&v679 {
																								v713 = v680
																								if v713 != 0 {
																									v716 = int32(1)
																								} else {
																									v716 = int32(-1)
																								}
																								v722 = v716
																							} else {
																								if v630 == int32(5) {
																									v712 = v627
																									v722 = v712
																								} else {
																									v688 = int32(5)
																									v690 = int32(128)
																									v691 = v634 << (uint(v688) % 32) & v690
																									if v691 != v638<<(uint(v688)%32)&v690 {
																										v713 = v691
																										if v713 != 0 {
																											v716 = int32(1)
																										} else {
																											v716 = int32(-1)
																										}
																										v722 = v716
																									} else {
																										if v630 < int32(7) {
																											v712 = v627
																											v722 = v712
																										} else {
																											v699 = int32(6)
																											v701 = int32(128)
																											v702 = v634 << (uint(v699) % 32) & v701
																											if v702 != v638<<(uint(v699)%32)&v701 {
																												v713 = v702
																												if v713 != 0 {
																													v716 = int32(1)
																												} else {
																													v716 = int32(-1)
																												}
																												v722 = v716
																											} else {
																												v712 = v627
																												v722 = v712
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
													switch v16 - int32(18) {
													case 0:
														v762 = v722
														return base.I64_extend_i32_u(base.B2i32(v762 == int32(0)))
													case 1:
														v770 = v722
														return base.I64_extend_i32_u(base.B2i32(v770 != int32(0)))
													case 2:
														return base.I64_extend_i32_u(int32(base.Ui32(v722) >> (uint(int32(31)) % 32)))
													case 3:
														return base.I64_extend_i32_u(base.B2i32(v722 <= int32(0)))
													case 4:
														return base.I64_extend_i32_u(base.B2i32(int32(0) < v722))
													case 5:
														return base.I64_extend_i32_u(base.B2i32(int32(0) <= v722))
													default:
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v744 = m.ExcPending
														if v744 != 0 {
															return int64(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_0), int32(0))
															mBase = m.M
															v748 = m.ExcPending
															if v748 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_inet_gist_consistent_1), int32(328), int32(_a_F_inet_gist_consistent_2))
																mBase = m.M
																v753 = m.ExcPending
																if v753 != 0 {
																	return int64(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													return int64(0)
												}
											}
										} else {
											if base.Ui32(v337) < base.Ui32(v336) {
												v760 = v21
												return v760
											} else {
												if base.Ui32(v337) <= base.Ui32(v336) {
													v606 = v333
													if v606 != 0 {
														v611 = int32(1)
													} else {
														v611 = int32(4)
													}
													v614 = v12 + v611 + int32(2)
													v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
													if v617 == int32(3) {
														v620 = int32(128)
													} else {
														v620 = int32(32)
													}
													v625 = base.I32_div_s(v620, int32(8))
													v626 = F_memcmp(m, v68, v614, v625)
													mBase = m.M
													if v626 != 0 {
														v712 = v626
														v722 = v712
													} else {
														v627 = int32(0)
														v630 = v620 - v625<<(uint(int32(3))%32)
														if v630 <= v627 {
															v712 = v627
															v722 = v712
														} else {
															v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v625))))
															v635 = int32(128)
															v636 = v634 & v635
															v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614+v625))))
															if v636 != v638&v635 {
																v713 = v636
																if v713 != 0 {
																	v716 = int32(1)
																} else {
																	v716 = int32(-1)
																}
																v722 = v716
															} else {
																if v630 == int32(1) {
																	v712 = v627
																	v722 = v712
																} else {
																	v644 = int32(1)
																	v646 = int32(128)
																	v647 = v634 << (uint(v644) % 32) & v646
																	if v647 != v638<<(uint(v644)%32)&v646 {
																		v713 = v647
																		if v713 != 0 {
																			v716 = int32(1)
																		} else {
																			v716 = int32(-1)
																		}
																		v722 = v716
																	} else {
																		if v630 < int32(3) {
																			v712 = v627
																			v722 = v712
																		} else {
																			v655 = int32(2)
																			v657 = int32(128)
																			v658 = v634 << (uint(v655) % 32) & v657
																			if v658 != v638<<(uint(v655)%32)&v657 {
																				v713 = v658
																				if v713 != 0 {
																					v716 = int32(1)
																				} else {
																					v716 = int32(-1)
																				}
																				v722 = v716
																			} else {
																				if v630 == int32(3) {
																					v712 = v627
																					v722 = v712
																				} else {
																					v666 = int32(3)
																					v668 = int32(128)
																					v669 = v634 << (uint(v666) % 32) & v668
																					if v669 != v638<<(uint(v666)%32)&v668 {
																						v713 = v669
																						if v713 != 0 {
																							v716 = int32(1)
																						} else {
																							v716 = int32(-1)
																						}
																						v722 = v716
																					} else {
																						if v630 < int32(5) {
																							v712 = v627
																							v722 = v712
																						} else {
																							v677 = int32(4)
																							v679 = int32(128)
																							v680 = v634 << (uint(v677) % 32) & v679
																							if v680 != v638<<(uint(v677)%32)&v679 {
																								v713 = v680
																								if v713 != 0 {
																									v716 = int32(1)
																								} else {
																									v716 = int32(-1)
																								}
																								v722 = v716
																							} else {
																								if v630 == int32(5) {
																									v712 = v627
																									v722 = v712
																								} else {
																									v688 = int32(5)
																									v690 = int32(128)
																									v691 = v634 << (uint(v688) % 32) & v690
																									if v691 != v638<<(uint(v688)%32)&v690 {
																										v713 = v691
																										if v713 != 0 {
																											v716 = int32(1)
																										} else {
																											v716 = int32(-1)
																										}
																										v722 = v716
																									} else {
																										if v630 < int32(7) {
																											v712 = v627
																											v722 = v712
																										} else {
																											v699 = int32(6)
																											v701 = int32(128)
																											v702 = v634 << (uint(v699) % 32) & v701
																											if v702 != v638<<(uint(v699)%32)&v701 {
																												v713 = v702
																												if v713 != 0 {
																													v716 = int32(1)
																												} else {
																													v716 = int32(-1)
																												}
																												v722 = v716
																											} else {
																												v712 = v627
																												v722 = v712
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
													switch v16 - int32(18) {
													case 0:
														v762 = v722
														return base.I64_extend_i32_u(base.B2i32(v762 == int32(0)))
													case 1:
														v770 = v722
														return base.I64_extend_i32_u(base.B2i32(v770 != int32(0)))
													case 2:
														return base.I64_extend_i32_u(int32(base.Ui32(v722) >> (uint(int32(31)) % 32)))
													case 3:
														return base.I64_extend_i32_u(base.B2i32(v722 <= int32(0)))
													case 4:
														return base.I64_extend_i32_u(base.B2i32(int32(0) < v722))
													case 5:
														return base.I64_extend_i32_u(base.B2i32(int32(0) <= v722))
													default:
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v744 = m.ExcPending
														if v744 != 0 {
															return int64(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_0), int32(0))
															mBase = m.M
															v748 = m.ExcPending
															if v748 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_inet_gist_consistent_1), int32(328), int32(_a_F_inet_gist_consistent_2))
																mBase = m.M
																v753 = m.ExcPending
																if v753 != 0 {
																	return int64(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													return int64(0)
												}
											}
										}
									} else {
										v760 = v21
										return v760
									}
								}
							}
						case 19, 20:
							if v179 < int32(0) {
								return int64(0)
							} else {
								if v179 != 0 {
									v760 = v21
									return v760
								} else {
									v320 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
									v321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v320)+16)))
									v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320+v321)+12)))
									if v323&int32(1) == int32(0) {
										v760 = v21
										return v760
									} else {
										v329 = int32(1)
										v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
										v333 = v331 & v329
										if v333 != 0 {
											v334 = v329
										} else {
											v334 = int32(4)
										}
										v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v334)+1)))
										v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
										if v16&int32(_a_F_inet_gist_consistent_3) != int32(20) {
											if base.Ui32(v336) < base.Ui32(v337) {
												v760 = v21
												return v760
											} else {
												if base.Ui32(v336) <= base.Ui32(v337) {
													v606 = v333
													if v606 != 0 {
														v611 = int32(1)
													} else {
														v611 = int32(4)
													}
													v614 = v12 + v611 + int32(2)
													v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
													if v617 == int32(3) {
														v620 = int32(128)
													} else {
														v620 = int32(32)
													}
													v625 = base.I32_div_s(v620, int32(8))
													v626 = F_memcmp(m, v68, v614, v625)
													mBase = m.M
													if v626 != 0 {
														v712 = v626
														v722 = v712
													} else {
														v627 = int32(0)
														v630 = v620 - v625<<(uint(int32(3))%32)
														if v630 <= v627 {
															v712 = v627
															v722 = v712
														} else {
															v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v625))))
															v635 = int32(128)
															v636 = v634 & v635
															v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614+v625))))
															if v636 != v638&v635 {
																v713 = v636
																if v713 != 0 {
																	v716 = int32(1)
																} else {
																	v716 = int32(-1)
																}
																v722 = v716
															} else {
																if v630 == int32(1) {
																	v712 = v627
																	v722 = v712
																} else {
																	v644 = int32(1)
																	v646 = int32(128)
																	v647 = v634 << (uint(v644) % 32) & v646
																	if v647 != v638<<(uint(v644)%32)&v646 {
																		v713 = v647
																		if v713 != 0 {
																			v716 = int32(1)
																		} else {
																			v716 = int32(-1)
																		}
																		v722 = v716
																	} else {
																		if v630 < int32(3) {
																			v712 = v627
																			v722 = v712
																		} else {
																			v655 = int32(2)
																			v657 = int32(128)
																			v658 = v634 << (uint(v655) % 32) & v657
																			if v658 != v638<<(uint(v655)%32)&v657 {
																				v713 = v658
																				if v713 != 0 {
																					v716 = int32(1)
																				} else {
																					v716 = int32(-1)
																				}
																				v722 = v716
																			} else {
																				if v630 == int32(3) {
																					v712 = v627
																					v722 = v712
																				} else {
																					v666 = int32(3)
																					v668 = int32(128)
																					v669 = v634 << (uint(v666) % 32) & v668
																					if v669 != v638<<(uint(v666)%32)&v668 {
																						v713 = v669
																						if v713 != 0 {
																							v716 = int32(1)
																						} else {
																							v716 = int32(-1)
																						}
																						v722 = v716
																					} else {
																						if v630 < int32(5) {
																							v712 = v627
																							v722 = v712
																						} else {
																							v677 = int32(4)
																							v679 = int32(128)
																							v680 = v634 << (uint(v677) % 32) & v679
																							if v680 != v638<<(uint(v677)%32)&v679 {
																								v713 = v680
																								if v713 != 0 {
																									v716 = int32(1)
																								} else {
																									v716 = int32(-1)
																								}
																								v722 = v716
																							} else {
																								if v630 == int32(5) {
																									v712 = v627
																									v722 = v712
																								} else {
																									v688 = int32(5)
																									v690 = int32(128)
																									v691 = v634 << (uint(v688) % 32) & v690
																									if v691 != v638<<(uint(v688)%32)&v690 {
																										v713 = v691
																										if v713 != 0 {
																											v716 = int32(1)
																										} else {
																											v716 = int32(-1)
																										}
																										v722 = v716
																									} else {
																										if v630 < int32(7) {
																											v712 = v627
																											v722 = v712
																										} else {
																											v699 = int32(6)
																											v701 = int32(128)
																											v702 = v634 << (uint(v699) % 32) & v701
																											if v702 != v638<<(uint(v699)%32)&v701 {
																												v713 = v702
																												if v713 != 0 {
																													v716 = int32(1)
																												} else {
																													v716 = int32(-1)
																												}
																												v722 = v716
																											} else {
																												v712 = v627
																												v722 = v712
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
													switch v16 - int32(18) {
													case 0:
														v762 = v722
														return base.I64_extend_i32_u(base.B2i32(v762 == int32(0)))
													case 1:
														v770 = v722
														return base.I64_extend_i32_u(base.B2i32(v770 != int32(0)))
													case 2:
														return base.I64_extend_i32_u(int32(base.Ui32(v722) >> (uint(int32(31)) % 32)))
													case 3:
														return base.I64_extend_i32_u(base.B2i32(v722 <= int32(0)))
													case 4:
														return base.I64_extend_i32_u(base.B2i32(int32(0) < v722))
													case 5:
														return base.I64_extend_i32_u(base.B2i32(int32(0) <= v722))
													default:
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v744 = m.ExcPending
														if v744 != 0 {
															return int64(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_0), int32(0))
															mBase = m.M
															v748 = m.ExcPending
															if v748 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_inet_gist_consistent_1), int32(328), int32(_a_F_inet_gist_consistent_2))
																mBase = m.M
																v753 = m.ExcPending
																if v753 != 0 {
																	return int64(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													return int64(0)
												}
											}
										} else {
											if base.Ui32(v337) < base.Ui32(v336) {
												v760 = v21
												return v760
											} else {
												if base.Ui32(v337) <= base.Ui32(v336) {
													v606 = v333
													if v606 != 0 {
														v611 = int32(1)
													} else {
														v611 = int32(4)
													}
													v614 = v12 + v611 + int32(2)
													v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
													if v617 == int32(3) {
														v620 = int32(128)
													} else {
														v620 = int32(32)
													}
													v625 = base.I32_div_s(v620, int32(8))
													v626 = F_memcmp(m, v68, v614, v625)
													mBase = m.M
													if v626 != 0 {
														v712 = v626
														v722 = v712
													} else {
														v627 = int32(0)
														v630 = v620 - v625<<(uint(int32(3))%32)
														if v630 <= v627 {
															v712 = v627
															v722 = v712
														} else {
															v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v625))))
															v635 = int32(128)
															v636 = v634 & v635
															v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614+v625))))
															if v636 != v638&v635 {
																v713 = v636
																if v713 != 0 {
																	v716 = int32(1)
																} else {
																	v716 = int32(-1)
																}
																v722 = v716
															} else {
																if v630 == int32(1) {
																	v712 = v627
																	v722 = v712
																} else {
																	v644 = int32(1)
																	v646 = int32(128)
																	v647 = v634 << (uint(v644) % 32) & v646
																	if v647 != v638<<(uint(v644)%32)&v646 {
																		v713 = v647
																		if v713 != 0 {
																			v716 = int32(1)
																		} else {
																			v716 = int32(-1)
																		}
																		v722 = v716
																	} else {
																		if v630 < int32(3) {
																			v712 = v627
																			v722 = v712
																		} else {
																			v655 = int32(2)
																			v657 = int32(128)
																			v658 = v634 << (uint(v655) % 32) & v657
																			if v658 != v638<<(uint(v655)%32)&v657 {
																				v713 = v658
																				if v713 != 0 {
																					v716 = int32(1)
																				} else {
																					v716 = int32(-1)
																				}
																				v722 = v716
																			} else {
																				if v630 == int32(3) {
																					v712 = v627
																					v722 = v712
																				} else {
																					v666 = int32(3)
																					v668 = int32(128)
																					v669 = v634 << (uint(v666) % 32) & v668
																					if v669 != v638<<(uint(v666)%32)&v668 {
																						v713 = v669
																						if v713 != 0 {
																							v716 = int32(1)
																						} else {
																							v716 = int32(-1)
																						}
																						v722 = v716
																					} else {
																						if v630 < int32(5) {
																							v712 = v627
																							v722 = v712
																						} else {
																							v677 = int32(4)
																							v679 = int32(128)
																							v680 = v634 << (uint(v677) % 32) & v679
																							if v680 != v638<<(uint(v677)%32)&v679 {
																								v713 = v680
																								if v713 != 0 {
																									v716 = int32(1)
																								} else {
																									v716 = int32(-1)
																								}
																								v722 = v716
																							} else {
																								if v630 == int32(5) {
																									v712 = v627
																									v722 = v712
																								} else {
																									v688 = int32(5)
																									v690 = int32(128)
																									v691 = v634 << (uint(v688) % 32) & v690
																									if v691 != v638<<(uint(v688)%32)&v690 {
																										v713 = v691
																										if v713 != 0 {
																											v716 = int32(1)
																										} else {
																											v716 = int32(-1)
																										}
																										v722 = v716
																									} else {
																										if v630 < int32(7) {
																											v712 = v627
																											v722 = v712
																										} else {
																											v699 = int32(6)
																											v701 = int32(128)
																											v702 = v634 << (uint(v699) % 32) & v701
																											if v702 != v638<<(uint(v699)%32)&v701 {
																												v713 = v702
																												if v713 != 0 {
																													v716 = int32(1)
																												} else {
																													v716 = int32(-1)
																												}
																												v722 = v716
																											} else {
																												v712 = v627
																												v722 = v712
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
													switch v16 - int32(18) {
													case 0:
														v762 = v722
														return base.I64_extend_i32_u(base.B2i32(v762 == int32(0)))
													case 1:
														v770 = v722
														return base.I64_extend_i32_u(base.B2i32(v770 != int32(0)))
													case 2:
														return base.I64_extend_i32_u(int32(base.Ui32(v722) >> (uint(int32(31)) % 32)))
													case 3:
														return base.I64_extend_i32_u(base.B2i32(v722 <= int32(0)))
													case 4:
														return base.I64_extend_i32_u(base.B2i32(int32(0) < v722))
													case 5:
														return base.I64_extend_i32_u(base.B2i32(int32(0) <= v722))
													default:
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v744 = m.ExcPending
														if v744 != 0 {
															return int64(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_0), int32(0))
															mBase = m.M
															v748 = m.ExcPending
															if v748 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_inet_gist_consistent_1), int32(328), int32(_a_F_inet_gist_consistent_2))
																mBase = m.M
																v753 = m.ExcPending
																if v753 != 0 {
																	return int64(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													return int64(0)
												}
											}
										}
									}
								}
							}
						}
					} else {
						return int64(0)
					}
				default:
					v68 = v17 + int32(4)
					v70 = v31 + int32(2)
					v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+3)))
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
					if base.Ui32(v71) < base.Ui32(v72) {
						v74 = v71
					} else {
						v74 = v72
					}
					v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
					if base.Ui32(v74) < base.Ui32(v75) {
						v77 = v74
					} else {
						v77 = v75
					}
					v82 = base.I32_div_s(v77, int32(8))
					v83 = F_memcmp(m, v68, v70, v82)
					mBase = m.M
					if v83 != 0 {
						v169 = v83
						v179 = v169
					} else {
						v84 = int32(0)
						v87 = v77 - v82<<(uint(int32(3))%32)
						if v87 <= v84 {
							v169 = v84
							v179 = v169
						} else {
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v82))))
							v92 = int32(128)
							v93 = v91 & v92
							v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v82))))
							if v93 != v95&v92 {
								v170 = v93
								if v170 != 0 {
									v173 = int32(1)
								} else {
									v173 = int32(-1)
								}
								v179 = v173
							} else {
								if v87 == int32(1) {
									v169 = v84
									v179 = v169
								} else {
									v101 = int32(1)
									v103 = int32(128)
									v104 = v91 << (uint(v101) % 32) & v103
									if v104 != v95<<(uint(v101)%32)&v103 {
										v170 = v104
										if v170 != 0 {
											v173 = int32(1)
										} else {
											v173 = int32(-1)
										}
										v179 = v173
									} else {
										if v87 < int32(3) {
											v169 = v84
											v179 = v169
										} else {
											v112 = int32(2)
											v114 = int32(128)
											v115 = v91 << (uint(v112) % 32) & v114
											if v115 != v95<<(uint(v112)%32)&v114 {
												v170 = v115
												if v170 != 0 {
													v173 = int32(1)
												} else {
													v173 = int32(-1)
												}
												v179 = v173
											} else {
												if v87 == int32(3) {
													v169 = v84
													v179 = v169
												} else {
													v123 = int32(3)
													v125 = int32(128)
													v126 = v91 << (uint(v123) % 32) & v125
													if v126 != v95<<(uint(v123)%32)&v125 {
														v170 = v126
														if v170 != 0 {
															v173 = int32(1)
														} else {
															v173 = int32(-1)
														}
														v179 = v173
													} else {
														if v87 < int32(5) {
															v169 = v84
															v179 = v169
														} else {
															v134 = int32(4)
															v136 = int32(128)
															v137 = v91 << (uint(v134) % 32) & v136
															if v137 != v95<<(uint(v134)%32)&v136 {
																v170 = v137
																if v170 != 0 {
																	v173 = int32(1)
																} else {
																	v173 = int32(-1)
																}
																v179 = v173
															} else {
																if v87 == int32(5) {
																	v169 = v84
																	v179 = v169
																} else {
																	v145 = int32(5)
																	v147 = int32(128)
																	v148 = v91 << (uint(v145) % 32) & v147
																	if v148 != v95<<(uint(v145)%32)&v147 {
																		v170 = v148
																		if v170 != 0 {
																			v173 = int32(1)
																		} else {
																			v173 = int32(-1)
																		}
																		v179 = v173
																	} else {
																		if v87 < int32(7) {
																			v169 = v84
																			v179 = v169
																		} else {
																			v156 = int32(6)
																			v158 = int32(128)
																			v159 = v91 << (uint(v156) % 32) & v158
																			if v159 != v95<<(uint(v156)%32)&v158 {
																				v170 = v159
																				if v170 != 0 {
																					v173 = int32(1)
																				} else {
																					v173 = int32(-1)
																				}
																				v179 = v173
																			} else {
																				v169 = v84
																				v179 = v169
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
					switch v16 - int32(3) {
					case 0, 21, 22, 23, 24:
						v296 = v179
						return base.I64_extend_i32_u(base.B2i32(v296 == int32(0)))
					default:
						v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
						v606 = v603 & int32(1)
						if v606 != 0 {
							v611 = int32(1)
						} else {
							v611 = int32(4)
						}
						v614 = v12 + v611 + int32(2)
						v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
						if v617 == int32(3) {
							v620 = int32(128)
						} else {
							v620 = int32(32)
						}
						v625 = base.I32_div_s(v620, int32(8))
						v626 = F_memcmp(m, v68, v614, v625)
						mBase = m.M
						if v626 != 0 {
							v712 = v626
							v722 = v712
						} else {
							v627 = int32(0)
							v630 = v620 - v625<<(uint(int32(3))%32)
							if v630 <= v627 {
								v712 = v627
								v722 = v712
							} else {
								v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v625))))
								v635 = int32(128)
								v636 = v634 & v635
								v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614+v625))))
								if v636 != v638&v635 {
									v713 = v636
									if v713 != 0 {
										v716 = int32(1)
									} else {
										v716 = int32(-1)
									}
									v722 = v716
								} else {
									if v630 == int32(1) {
										v712 = v627
										v722 = v712
									} else {
										v644 = int32(1)
										v646 = int32(128)
										v647 = v634 << (uint(v644) % 32) & v646
										if v647 != v638<<(uint(v644)%32)&v646 {
											v713 = v647
											if v713 != 0 {
												v716 = int32(1)
											} else {
												v716 = int32(-1)
											}
											v722 = v716
										} else {
											if v630 < int32(3) {
												v712 = v627
												v722 = v712
											} else {
												v655 = int32(2)
												v657 = int32(128)
												v658 = v634 << (uint(v655) % 32) & v657
												if v658 != v638<<(uint(v655)%32)&v657 {
													v713 = v658
													if v713 != 0 {
														v716 = int32(1)
													} else {
														v716 = int32(-1)
													}
													v722 = v716
												} else {
													if v630 == int32(3) {
														v712 = v627
														v722 = v712
													} else {
														v666 = int32(3)
														v668 = int32(128)
														v669 = v634 << (uint(v666) % 32) & v668
														if v669 != v638<<(uint(v666)%32)&v668 {
															v713 = v669
															if v713 != 0 {
																v716 = int32(1)
															} else {
																v716 = int32(-1)
															}
															v722 = v716
														} else {
															if v630 < int32(5) {
																v712 = v627
																v722 = v712
															} else {
																v677 = int32(4)
																v679 = int32(128)
																v680 = v634 << (uint(v677) % 32) & v679
																if v680 != v638<<(uint(v677)%32)&v679 {
																	v713 = v680
																	if v713 != 0 {
																		v716 = int32(1)
																	} else {
																		v716 = int32(-1)
																	}
																	v722 = v716
																} else {
																	if v630 == int32(5) {
																		v712 = v627
																		v722 = v712
																	} else {
																		v688 = int32(5)
																		v690 = int32(128)
																		v691 = v634 << (uint(v688) % 32) & v690
																		if v691 != v638<<(uint(v688)%32)&v690 {
																			v713 = v691
																			if v713 != 0 {
																				v716 = int32(1)
																			} else {
																				v716 = int32(-1)
																			}
																			v722 = v716
																		} else {
																			if v630 < int32(7) {
																				v712 = v627
																				v722 = v712
																			} else {
																				v699 = int32(6)
																				v701 = int32(128)
																				v702 = v634 << (uint(v699) % 32) & v701
																				if v702 != v638<<(uint(v699)%32)&v701 {
																					v713 = v702
																					if v713 != 0 {
																						v716 = int32(1)
																					} else {
																						v716 = int32(-1)
																					}
																					v722 = v716
																				} else {
																					v712 = v627
																					v722 = v712
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
						switch v16 - int32(18) {
						case 0:
							v762 = v722
							return base.I64_extend_i32_u(base.B2i32(v762 == int32(0)))
						case 1:
							v770 = v722
							return base.I64_extend_i32_u(base.B2i32(v770 != int32(0)))
						case 2:
							return base.I64_extend_i32_u(int32(base.Ui32(v722) >> (uint(int32(31)) % 32)))
						case 3:
							return base.I64_extend_i32_u(base.B2i32(v722 <= int32(0)))
						case 4:
							return base.I64_extend_i32_u(base.B2i32(int32(0) < v722))
						case 5:
							return base.I64_extend_i32_u(base.B2i32(int32(0) <= v722))
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v744 = m.ExcPending
							if v744 != 0 {
								return int64(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_0), int32(0))
								mBase = m.M
								v748 = m.ExcPending
								if v748 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_inet_gist_consistent_1), int32(328), int32(_a_F_inet_gist_consistent_2))
									mBase = m.M
									v753 = m.ExcPending
									if v753 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					case 15:
						if v179 == int32(0) {
							v472 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
							v473 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v472)+16)))
							v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472+v473)+12)))
							if v475&int32(1) == int32(0) {
								v760 = v21
								return v760
							} else {
								v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
								v482 = int32(1)
								v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
								if v484&v482 != 0 {
									v487 = v482
								} else {
									v487 = int32(4)
								}
								v488 = v12 + v487
								v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+1)))
								if v481 != v489 {
									v760 = int64(0)
									return v760
								} else {
									v492 = v488 + int32(2)
									v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
									if v495 == int32(3) {
										v498 = int32(128)
									} else {
										v498 = int32(32)
									}
									v503 = base.I32_div_s(v498, int32(8))
									v504 = F_memcmp(m, v68, v492, v503)
									mBase = m.M
									if v504 != 0 {
										v590 = v504
										v600 = v590
									} else {
										v505 = int32(0)
										v508 = v498 - v503<<(uint(int32(3))%32)
										if v508 <= v505 {
											v590 = v505
											v600 = v590
										} else {
											v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v503))))
											v513 = int32(128)
											v514 = v512 & v513
											v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492+v503))))
											if v514 != v516&v513 {
												v591 = v514
												if v591 != 0 {
													v594 = int32(1)
												} else {
													v594 = int32(-1)
												}
												v600 = v594
											} else {
												if v508 == int32(1) {
													v590 = v505
													v600 = v590
												} else {
													v522 = int32(1)
													v524 = int32(128)
													v525 = v512 << (uint(v522) % 32) & v524
													if v525 != v516<<(uint(v522)%32)&v524 {
														v591 = v525
														if v591 != 0 {
															v594 = int32(1)
														} else {
															v594 = int32(-1)
														}
														v600 = v594
													} else {
														if v508 < int32(3) {
															v590 = v505
															v600 = v590
														} else {
															v533 = int32(2)
															v535 = int32(128)
															v536 = v512 << (uint(v533) % 32) & v535
															if v536 != v516<<(uint(v533)%32)&v535 {
																v591 = v536
																if v591 != 0 {
																	v594 = int32(1)
																} else {
																	v594 = int32(-1)
																}
																v600 = v594
															} else {
																if v508 == int32(3) {
																	v590 = v505
																	v600 = v590
																} else {
																	v544 = int32(3)
																	v546 = int32(128)
																	v547 = v512 << (uint(v544) % 32) & v546
																	if v547 != v516<<(uint(v544)%32)&v546 {
																		v591 = v547
																		if v591 != 0 {
																			v594 = int32(1)
																		} else {
																			v594 = int32(-1)
																		}
																		v600 = v594
																	} else {
																		if v508 < int32(5) {
																			v590 = v505
																			v600 = v590
																		} else {
																			v555 = int32(4)
																			v557 = int32(128)
																			v558 = v512 << (uint(v555) % 32) & v557
																			if v558 != v516<<(uint(v555)%32)&v557 {
																				v591 = v558
																				if v591 != 0 {
																					v594 = int32(1)
																				} else {
																					v594 = int32(-1)
																				}
																				v600 = v594
																			} else {
																				if v508 == int32(5) {
																					v590 = v505
																					v600 = v590
																				} else {
																					v566 = int32(5)
																					v568 = int32(128)
																					v569 = v512 << (uint(v566) % 32) & v568
																					if v569 != v516<<(uint(v566)%32)&v568 {
																						v591 = v569
																						if v591 != 0 {
																							v594 = int32(1)
																						} else {
																							v594 = int32(-1)
																						}
																						v600 = v594
																					} else {
																						if v508 < int32(7) {
																							v590 = v505
																							v600 = v590
																						} else {
																							v577 = int32(6)
																							v579 = int32(128)
																							v580 = v512 << (uint(v577) % 32) & v579
																							if v580 != v516<<(uint(v577)%32)&v579 {
																								v591 = v580
																								if v591 != 0 {
																									v594 = int32(1)
																								} else {
																									v594 = int32(-1)
																								}
																								v600 = v594
																							} else {
																								v590 = v505
																								v600 = v590
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
									v762 = v600
									return base.I64_extend_i32_u(base.B2i32(v762 == int32(0)))
								}
							}
						} else {
							return int64(0)
						}
					case 16:
						if v179 != 0 {
							v760 = v21
							return v760
						} else {
							v344 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
							v345 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v344)+16)))
							v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344+v345)+12)))
							if v347&int32(1) == int32(0) {
								v760 = v21
								return v760
							} else {
								v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
								v353 = int32(1)
								v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
								if v355&v353 != 0 {
									v358 = v353
								} else {
									v358 = int32(4)
								}
								v359 = v12 + v358
								v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359)+1)))
								if v352 != v360 {
									v760 = v21
									return v760
								} else {
									v363 = v359 + int32(2)
									v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
									if v366 == int32(3) {
										v369 = int32(128)
									} else {
										v369 = int32(32)
									}
									v374 = base.I32_div_s(v369, int32(8))
									v375 = F_memcmp(m, v68, v363, v374)
									mBase = m.M
									if v375 != 0 {
										v461 = v375
										v471 = v461
									} else {
										v376 = int32(0)
										v379 = v369 - v374<<(uint(int32(3))%32)
										if v379 <= v376 {
											v461 = v376
											v471 = v461
										} else {
											v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v374))))
											v384 = int32(128)
											v385 = v383 & v384
											v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+v374))))
											if v385 != v387&v384 {
												v462 = v385
												if v462 != 0 {
													v465 = int32(1)
												} else {
													v465 = int32(-1)
												}
												v471 = v465
											} else {
												if v379 == int32(1) {
													v461 = v376
													v471 = v461
												} else {
													v393 = int32(1)
													v395 = int32(128)
													v396 = v383 << (uint(v393) % 32) & v395
													if v396 != v387<<(uint(v393)%32)&v395 {
														v462 = v396
														if v462 != 0 {
															v465 = int32(1)
														} else {
															v465 = int32(-1)
														}
														v471 = v465
													} else {
														if v379 < int32(3) {
															v461 = v376
															v471 = v461
														} else {
															v404 = int32(2)
															v406 = int32(128)
															v407 = v383 << (uint(v404) % 32) & v406
															if v407 != v387<<(uint(v404)%32)&v406 {
																v462 = v407
																if v462 != 0 {
																	v465 = int32(1)
																} else {
																	v465 = int32(-1)
																}
																v471 = v465
															} else {
																if v379 == int32(3) {
																	v461 = v376
																	v471 = v461
																} else {
																	v415 = int32(3)
																	v417 = int32(128)
																	v418 = v383 << (uint(v415) % 32) & v417
																	if v418 != v387<<(uint(v415)%32)&v417 {
																		v462 = v418
																		if v462 != 0 {
																			v465 = int32(1)
																		} else {
																			v465 = int32(-1)
																		}
																		v471 = v465
																	} else {
																		if v379 < int32(5) {
																			v461 = v376
																			v471 = v461
																		} else {
																			v426 = int32(4)
																			v428 = int32(128)
																			v429 = v383 << (uint(v426) % 32) & v428
																			if v429 != v387<<(uint(v426)%32)&v428 {
																				v462 = v429
																				if v462 != 0 {
																					v465 = int32(1)
																				} else {
																					v465 = int32(-1)
																				}
																				v471 = v465
																			} else {
																				if v379 == int32(5) {
																					v461 = v376
																					v471 = v461
																				} else {
																					v437 = int32(5)
																					v439 = int32(128)
																					v440 = v383 << (uint(v437) % 32) & v439
																					if v440 != v387<<(uint(v437)%32)&v439 {
																						v462 = v440
																						if v462 != 0 {
																							v465 = int32(1)
																						} else {
																							v465 = int32(-1)
																						}
																						v471 = v465
																					} else {
																						if v379 < int32(7) {
																							v461 = v376
																							v471 = v461
																						} else {
																							v448 = int32(6)
																							v450 = int32(128)
																							v451 = v383 << (uint(v448) % 32) & v450
																							if v451 != v387<<(uint(v448)%32)&v450 {
																								v462 = v451
																								if v462 != 0 {
																									v465 = int32(1)
																								} else {
																									v465 = int32(-1)
																								}
																								v471 = v465
																							} else {
																								v461 = v376
																								v471 = v461
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
									v770 = v471
									return base.I64_extend_i32_u(base.B2i32(v770 != int32(0)))
								}
							}
						}
					case 17, 18:
						if int32(0) < v179 {
							return int64(0)
						} else {
							if v179 < int32(0) {
								v760 = v21
								return v760
							} else {
								v310 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
								v311 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v310)+16)))
								v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310+v311)+12)))
								if v313&int32(1) != 0 {
									v329 = int32(1)
									v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
									v333 = v331 & v329
									if v333 != 0 {
										v334 = v329
									} else {
										v334 = int32(4)
									}
									v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v334)+1)))
									v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
									if v16&int32(_a_F_inet_gist_consistent_3) != int32(20) {
										if base.Ui32(v336) < base.Ui32(v337) {
											v760 = v21
											return v760
										} else {
											if base.Ui32(v336) <= base.Ui32(v337) {
												v606 = v333
												if v606 != 0 {
													v611 = int32(1)
												} else {
													v611 = int32(4)
												}
												v614 = v12 + v611 + int32(2)
												v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
												if v617 == int32(3) {
													v620 = int32(128)
												} else {
													v620 = int32(32)
												}
												v625 = base.I32_div_s(v620, int32(8))
												v626 = F_memcmp(m, v68, v614, v625)
												mBase = m.M
												if v626 != 0 {
													v712 = v626
													v722 = v712
												} else {
													v627 = int32(0)
													v630 = v620 - v625<<(uint(int32(3))%32)
													if v630 <= v627 {
														v712 = v627
														v722 = v712
													} else {
														v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v625))))
														v635 = int32(128)
														v636 = v634 & v635
														v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614+v625))))
														if v636 != v638&v635 {
															v713 = v636
															if v713 != 0 {
																v716 = int32(1)
															} else {
																v716 = int32(-1)
															}
															v722 = v716
														} else {
															if v630 == int32(1) {
																v712 = v627
																v722 = v712
															} else {
																v644 = int32(1)
																v646 = int32(128)
																v647 = v634 << (uint(v644) % 32) & v646
																if v647 != v638<<(uint(v644)%32)&v646 {
																	v713 = v647
																	if v713 != 0 {
																		v716 = int32(1)
																	} else {
																		v716 = int32(-1)
																	}
																	v722 = v716
																} else {
																	if v630 < int32(3) {
																		v712 = v627
																		v722 = v712
																	} else {
																		v655 = int32(2)
																		v657 = int32(128)
																		v658 = v634 << (uint(v655) % 32) & v657
																		if v658 != v638<<(uint(v655)%32)&v657 {
																			v713 = v658
																			if v713 != 0 {
																				v716 = int32(1)
																			} else {
																				v716 = int32(-1)
																			}
																			v722 = v716
																		} else {
																			if v630 == int32(3) {
																				v712 = v627
																				v722 = v712
																			} else {
																				v666 = int32(3)
																				v668 = int32(128)
																				v669 = v634 << (uint(v666) % 32) & v668
																				if v669 != v638<<(uint(v666)%32)&v668 {
																					v713 = v669
																					if v713 != 0 {
																						v716 = int32(1)
																					} else {
																						v716 = int32(-1)
																					}
																					v722 = v716
																				} else {
																					if v630 < int32(5) {
																						v712 = v627
																						v722 = v712
																					} else {
																						v677 = int32(4)
																						v679 = int32(128)
																						v680 = v634 << (uint(v677) % 32) & v679
																						if v680 != v638<<(uint(v677)%32)&v679 {
																							v713 = v680
																							if v713 != 0 {
																								v716 = int32(1)
																							} else {
																								v716 = int32(-1)
																							}
																							v722 = v716
																						} else {
																							if v630 == int32(5) {
																								v712 = v627
																								v722 = v712
																							} else {
																								v688 = int32(5)
																								v690 = int32(128)
																								v691 = v634 << (uint(v688) % 32) & v690
																								if v691 != v638<<(uint(v688)%32)&v690 {
																									v713 = v691
																									if v713 != 0 {
																										v716 = int32(1)
																									} else {
																										v716 = int32(-1)
																									}
																									v722 = v716
																								} else {
																									if v630 < int32(7) {
																										v712 = v627
																										v722 = v712
																									} else {
																										v699 = int32(6)
																										v701 = int32(128)
																										v702 = v634 << (uint(v699) % 32) & v701
																										if v702 != v638<<(uint(v699)%32)&v701 {
																											v713 = v702
																											if v713 != 0 {
																												v716 = int32(1)
																											} else {
																												v716 = int32(-1)
																											}
																											v722 = v716
																										} else {
																											v712 = v627
																											v722 = v712
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
												switch v16 - int32(18) {
												case 0:
													v762 = v722
													return base.I64_extend_i32_u(base.B2i32(v762 == int32(0)))
												case 1:
													v770 = v722
													return base.I64_extend_i32_u(base.B2i32(v770 != int32(0)))
												case 2:
													return base.I64_extend_i32_u(int32(base.Ui32(v722) >> (uint(int32(31)) % 32)))
												case 3:
													return base.I64_extend_i32_u(base.B2i32(v722 <= int32(0)))
												case 4:
													return base.I64_extend_i32_u(base.B2i32(int32(0) < v722))
												case 5:
													return base.I64_extend_i32_u(base.B2i32(int32(0) <= v722))
												default:
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v744 = m.ExcPending
													if v744 != 0 {
														return int64(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_0), int32(0))
														mBase = m.M
														v748 = m.ExcPending
														if v748 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_inet_gist_consistent_1), int32(328), int32(_a_F_inet_gist_consistent_2))
															mBase = m.M
															v753 = m.ExcPending
															if v753 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												return int64(0)
											}
										}
									} else {
										if base.Ui32(v337) < base.Ui32(v336) {
											v760 = v21
											return v760
										} else {
											if base.Ui32(v337) <= base.Ui32(v336) {
												v606 = v333
												if v606 != 0 {
													v611 = int32(1)
												} else {
													v611 = int32(4)
												}
												v614 = v12 + v611 + int32(2)
												v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
												if v617 == int32(3) {
													v620 = int32(128)
												} else {
													v620 = int32(32)
												}
												v625 = base.I32_div_s(v620, int32(8))
												v626 = F_memcmp(m, v68, v614, v625)
												mBase = m.M
												if v626 != 0 {
													v712 = v626
													v722 = v712
												} else {
													v627 = int32(0)
													v630 = v620 - v625<<(uint(int32(3))%32)
													if v630 <= v627 {
														v712 = v627
														v722 = v712
													} else {
														v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v625))))
														v635 = int32(128)
														v636 = v634 & v635
														v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614+v625))))
														if v636 != v638&v635 {
															v713 = v636
															if v713 != 0 {
																v716 = int32(1)
															} else {
																v716 = int32(-1)
															}
															v722 = v716
														} else {
															if v630 == int32(1) {
																v712 = v627
																v722 = v712
															} else {
																v644 = int32(1)
																v646 = int32(128)
																v647 = v634 << (uint(v644) % 32) & v646
																if v647 != v638<<(uint(v644)%32)&v646 {
																	v713 = v647
																	if v713 != 0 {
																		v716 = int32(1)
																	} else {
																		v716 = int32(-1)
																	}
																	v722 = v716
																} else {
																	if v630 < int32(3) {
																		v712 = v627
																		v722 = v712
																	} else {
																		v655 = int32(2)
																		v657 = int32(128)
																		v658 = v634 << (uint(v655) % 32) & v657
																		if v658 != v638<<(uint(v655)%32)&v657 {
																			v713 = v658
																			if v713 != 0 {
																				v716 = int32(1)
																			} else {
																				v716 = int32(-1)
																			}
																			v722 = v716
																		} else {
																			if v630 == int32(3) {
																				v712 = v627
																				v722 = v712
																			} else {
																				v666 = int32(3)
																				v668 = int32(128)
																				v669 = v634 << (uint(v666) % 32) & v668
																				if v669 != v638<<(uint(v666)%32)&v668 {
																					v713 = v669
																					if v713 != 0 {
																						v716 = int32(1)
																					} else {
																						v716 = int32(-1)
																					}
																					v722 = v716
																				} else {
																					if v630 < int32(5) {
																						v712 = v627
																						v722 = v712
																					} else {
																						v677 = int32(4)
																						v679 = int32(128)
																						v680 = v634 << (uint(v677) % 32) & v679
																						if v680 != v638<<(uint(v677)%32)&v679 {
																							v713 = v680
																							if v713 != 0 {
																								v716 = int32(1)
																							} else {
																								v716 = int32(-1)
																							}
																							v722 = v716
																						} else {
																							if v630 == int32(5) {
																								v712 = v627
																								v722 = v712
																							} else {
																								v688 = int32(5)
																								v690 = int32(128)
																								v691 = v634 << (uint(v688) % 32) & v690
																								if v691 != v638<<(uint(v688)%32)&v690 {
																									v713 = v691
																									if v713 != 0 {
																										v716 = int32(1)
																									} else {
																										v716 = int32(-1)
																									}
																									v722 = v716
																								} else {
																									if v630 < int32(7) {
																										v712 = v627
																										v722 = v712
																									} else {
																										v699 = int32(6)
																										v701 = int32(128)
																										v702 = v634 << (uint(v699) % 32) & v701
																										if v702 != v638<<(uint(v699)%32)&v701 {
																											v713 = v702
																											if v713 != 0 {
																												v716 = int32(1)
																											} else {
																												v716 = int32(-1)
																											}
																											v722 = v716
																										} else {
																											v712 = v627
																											v722 = v712
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
												switch v16 - int32(18) {
												case 0:
													v762 = v722
													return base.I64_extend_i32_u(base.B2i32(v762 == int32(0)))
												case 1:
													v770 = v722
													return base.I64_extend_i32_u(base.B2i32(v770 != int32(0)))
												case 2:
													return base.I64_extend_i32_u(int32(base.Ui32(v722) >> (uint(int32(31)) % 32)))
												case 3:
													return base.I64_extend_i32_u(base.B2i32(v722 <= int32(0)))
												case 4:
													return base.I64_extend_i32_u(base.B2i32(int32(0) < v722))
												case 5:
													return base.I64_extend_i32_u(base.B2i32(int32(0) <= v722))
												default:
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v744 = m.ExcPending
													if v744 != 0 {
														return int64(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_0), int32(0))
														mBase = m.M
														v748 = m.ExcPending
														if v748 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_inet_gist_consistent_1), int32(328), int32(_a_F_inet_gist_consistent_2))
															mBase = m.M
															v753 = m.ExcPending
															if v753 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												return int64(0)
											}
										}
									}
								} else {
									v760 = v21
									return v760
								}
							}
						}
					case 19, 20:
						if v179 < int32(0) {
							return int64(0)
						} else {
							if v179 != 0 {
								v760 = v21
								return v760
							} else {
								v320 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
								v321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v320)+16)))
								v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320+v321)+12)))
								if v323&int32(1) == int32(0) {
									v760 = v21
									return v760
								} else {
									v329 = int32(1)
									v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
									v333 = v331 & v329
									if v333 != 0 {
										v334 = v329
									} else {
										v334 = int32(4)
									}
									v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v334)+1)))
									v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
									if v16&int32(_a_F_inet_gist_consistent_3) != int32(20) {
										if base.Ui32(v336) < base.Ui32(v337) {
											v760 = v21
											return v760
										} else {
											if base.Ui32(v336) <= base.Ui32(v337) {
												v606 = v333
												if v606 != 0 {
													v611 = int32(1)
												} else {
													v611 = int32(4)
												}
												v614 = v12 + v611 + int32(2)
												v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
												if v617 == int32(3) {
													v620 = int32(128)
												} else {
													v620 = int32(32)
												}
												v625 = base.I32_div_s(v620, int32(8))
												v626 = F_memcmp(m, v68, v614, v625)
												mBase = m.M
												if v626 != 0 {
													v712 = v626
													v722 = v712
												} else {
													v627 = int32(0)
													v630 = v620 - v625<<(uint(int32(3))%32)
													if v630 <= v627 {
														v712 = v627
														v722 = v712
													} else {
														v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v625))))
														v635 = int32(128)
														v636 = v634 & v635
														v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614+v625))))
														if v636 != v638&v635 {
															v713 = v636
															if v713 != 0 {
																v716 = int32(1)
															} else {
																v716 = int32(-1)
															}
															v722 = v716
														} else {
															if v630 == int32(1) {
																v712 = v627
																v722 = v712
															} else {
																v644 = int32(1)
																v646 = int32(128)
																v647 = v634 << (uint(v644) % 32) & v646
																if v647 != v638<<(uint(v644)%32)&v646 {
																	v713 = v647
																	if v713 != 0 {
																		v716 = int32(1)
																	} else {
																		v716 = int32(-1)
																	}
																	v722 = v716
																} else {
																	if v630 < int32(3) {
																		v712 = v627
																		v722 = v712
																	} else {
																		v655 = int32(2)
																		v657 = int32(128)
																		v658 = v634 << (uint(v655) % 32) & v657
																		if v658 != v638<<(uint(v655)%32)&v657 {
																			v713 = v658
																			if v713 != 0 {
																				v716 = int32(1)
																			} else {
																				v716 = int32(-1)
																			}
																			v722 = v716
																		} else {
																			if v630 == int32(3) {
																				v712 = v627
																				v722 = v712
																			} else {
																				v666 = int32(3)
																				v668 = int32(128)
																				v669 = v634 << (uint(v666) % 32) & v668
																				if v669 != v638<<(uint(v666)%32)&v668 {
																					v713 = v669
																					if v713 != 0 {
																						v716 = int32(1)
																					} else {
																						v716 = int32(-1)
																					}
																					v722 = v716
																				} else {
																					if v630 < int32(5) {
																						v712 = v627
																						v722 = v712
																					} else {
																						v677 = int32(4)
																						v679 = int32(128)
																						v680 = v634 << (uint(v677) % 32) & v679
																						if v680 != v638<<(uint(v677)%32)&v679 {
																							v713 = v680
																							if v713 != 0 {
																								v716 = int32(1)
																							} else {
																								v716 = int32(-1)
																							}
																							v722 = v716
																						} else {
																							if v630 == int32(5) {
																								v712 = v627
																								v722 = v712
																							} else {
																								v688 = int32(5)
																								v690 = int32(128)
																								v691 = v634 << (uint(v688) % 32) & v690
																								if v691 != v638<<(uint(v688)%32)&v690 {
																									v713 = v691
																									if v713 != 0 {
																										v716 = int32(1)
																									} else {
																										v716 = int32(-1)
																									}
																									v722 = v716
																								} else {
																									if v630 < int32(7) {
																										v712 = v627
																										v722 = v712
																									} else {
																										v699 = int32(6)
																										v701 = int32(128)
																										v702 = v634 << (uint(v699) % 32) & v701
																										if v702 != v638<<(uint(v699)%32)&v701 {
																											v713 = v702
																											if v713 != 0 {
																												v716 = int32(1)
																											} else {
																												v716 = int32(-1)
																											}
																											v722 = v716
																										} else {
																											v712 = v627
																											v722 = v712
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
												switch v16 - int32(18) {
												case 0:
													v762 = v722
													return base.I64_extend_i32_u(base.B2i32(v762 == int32(0)))
												case 1:
													v770 = v722
													return base.I64_extend_i32_u(base.B2i32(v770 != int32(0)))
												case 2:
													return base.I64_extend_i32_u(int32(base.Ui32(v722) >> (uint(int32(31)) % 32)))
												case 3:
													return base.I64_extend_i32_u(base.B2i32(v722 <= int32(0)))
												case 4:
													return base.I64_extend_i32_u(base.B2i32(int32(0) < v722))
												case 5:
													return base.I64_extend_i32_u(base.B2i32(int32(0) <= v722))
												default:
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v744 = m.ExcPending
													if v744 != 0 {
														return int64(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_0), int32(0))
														mBase = m.M
														v748 = m.ExcPending
														if v748 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_inet_gist_consistent_1), int32(328), int32(_a_F_inet_gist_consistent_2))
															mBase = m.M
															v753 = m.ExcPending
															if v753 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												return int64(0)
											}
										}
									} else {
										if base.Ui32(v337) < base.Ui32(v336) {
											v760 = v21
											return v760
										} else {
											if base.Ui32(v337) <= base.Ui32(v336) {
												v606 = v333
												if v606 != 0 {
													v611 = int32(1)
												} else {
													v611 = int32(4)
												}
												v614 = v12 + v611 + int32(2)
												v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
												if v617 == int32(3) {
													v620 = int32(128)
												} else {
													v620 = int32(32)
												}
												v625 = base.I32_div_s(v620, int32(8))
												v626 = F_memcmp(m, v68, v614, v625)
												mBase = m.M
												if v626 != 0 {
													v712 = v626
													v722 = v712
												} else {
													v627 = int32(0)
													v630 = v620 - v625<<(uint(int32(3))%32)
													if v630 <= v627 {
														v712 = v627
														v722 = v712
													} else {
														v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v625))))
														v635 = int32(128)
														v636 = v634 & v635
														v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614+v625))))
														if v636 != v638&v635 {
															v713 = v636
															if v713 != 0 {
																v716 = int32(1)
															} else {
																v716 = int32(-1)
															}
															v722 = v716
														} else {
															if v630 == int32(1) {
																v712 = v627
																v722 = v712
															} else {
																v644 = int32(1)
																v646 = int32(128)
																v647 = v634 << (uint(v644) % 32) & v646
																if v647 != v638<<(uint(v644)%32)&v646 {
																	v713 = v647
																	if v713 != 0 {
																		v716 = int32(1)
																	} else {
																		v716 = int32(-1)
																	}
																	v722 = v716
																} else {
																	if v630 < int32(3) {
																		v712 = v627
																		v722 = v712
																	} else {
																		v655 = int32(2)
																		v657 = int32(128)
																		v658 = v634 << (uint(v655) % 32) & v657
																		if v658 != v638<<(uint(v655)%32)&v657 {
																			v713 = v658
																			if v713 != 0 {
																				v716 = int32(1)
																			} else {
																				v716 = int32(-1)
																			}
																			v722 = v716
																		} else {
																			if v630 == int32(3) {
																				v712 = v627
																				v722 = v712
																			} else {
																				v666 = int32(3)
																				v668 = int32(128)
																				v669 = v634 << (uint(v666) % 32) & v668
																				if v669 != v638<<(uint(v666)%32)&v668 {
																					v713 = v669
																					if v713 != 0 {
																						v716 = int32(1)
																					} else {
																						v716 = int32(-1)
																					}
																					v722 = v716
																				} else {
																					if v630 < int32(5) {
																						v712 = v627
																						v722 = v712
																					} else {
																						v677 = int32(4)
																						v679 = int32(128)
																						v680 = v634 << (uint(v677) % 32) & v679
																						if v680 != v638<<(uint(v677)%32)&v679 {
																							v713 = v680
																							if v713 != 0 {
																								v716 = int32(1)
																							} else {
																								v716 = int32(-1)
																							}
																							v722 = v716
																						} else {
																							if v630 == int32(5) {
																								v712 = v627
																								v722 = v712
																							} else {
																								v688 = int32(5)
																								v690 = int32(128)
																								v691 = v634 << (uint(v688) % 32) & v690
																								if v691 != v638<<(uint(v688)%32)&v690 {
																									v713 = v691
																									if v713 != 0 {
																										v716 = int32(1)
																									} else {
																										v716 = int32(-1)
																									}
																									v722 = v716
																								} else {
																									if v630 < int32(7) {
																										v712 = v627
																										v722 = v712
																									} else {
																										v699 = int32(6)
																										v701 = int32(128)
																										v702 = v634 << (uint(v699) % 32) & v701
																										if v702 != v638<<(uint(v699)%32)&v701 {
																											v713 = v702
																											if v713 != 0 {
																												v716 = int32(1)
																											} else {
																												v716 = int32(-1)
																											}
																											v722 = v716
																										} else {
																											v712 = v627
																											v722 = v712
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
												switch v16 - int32(18) {
												case 0:
													v762 = v722
													return base.I64_extend_i32_u(base.B2i32(v762 == int32(0)))
												case 1:
													v770 = v722
													return base.I64_extend_i32_u(base.B2i32(v770 != int32(0)))
												case 2:
													return base.I64_extend_i32_u(int32(base.Ui32(v722) >> (uint(int32(31)) % 32)))
												case 3:
													return base.I64_extend_i32_u(base.B2i32(v722 <= int32(0)))
												case 4:
													return base.I64_extend_i32_u(base.B2i32(int32(0) < v722))
												case 5:
													return base.I64_extend_i32_u(base.B2i32(int32(0) <= v722))
												default:
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v744 = m.ExcPending
													if v744 != 0 {
														return int64(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_0), int32(0))
														mBase = m.M
														v748 = m.ExcPending
														if v748 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_inet_gist_consistent_1), int32(328), int32(_a_F_inet_gist_consistent_2))
															mBase = m.M
															v753 = m.ExcPending
															if v753 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												return int64(0)
											}
										}
									}
								}
							}
						}
					}
				case 6:
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+16)))
					v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v40)+12)))
					if v42&int32(1) == int32(0) {
						v184 = v17 + int32(4)
						v186 = v31 + int32(2)
						v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+3)))
						v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
						if base.Ui32(v187) < base.Ui32(v188) {
							v190 = v187
						} else {
							v190 = v188
						}
						v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
						if base.Ui32(v190) < base.Ui32(v191) {
							v193 = v190
						} else {
							v193 = v191
						}
						v198 = base.I32_div_s(v193, int32(8))
						v199 = F_memcmp(m, v184, v186, v198)
						mBase = m.M
						if v199 != 0 {
							v285 = v199
							v295 = v285
						} else {
							v200 = int32(0)
							v203 = v193 - v198<<(uint(int32(3))%32)
							if v203 <= v200 {
								v285 = v200
								v295 = v285
							} else {
								v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184+v198))))
								v208 = int32(128)
								v209 = v207 & v208
								v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186+v198))))
								if v209 != v211&v208 {
									v286 = v209
									if v286 != 0 {
										v289 = int32(1)
									} else {
										v289 = int32(-1)
									}
									v295 = v289
								} else {
									if v203 == int32(1) {
										v285 = v200
										v295 = v285
									} else {
										v217 = int32(1)
										v219 = int32(128)
										v220 = v207 << (uint(v217) % 32) & v219
										if v220 != v211<<(uint(v217)%32)&v219 {
											v286 = v220
											if v286 != 0 {
												v289 = int32(1)
											} else {
												v289 = int32(-1)
											}
											v295 = v289
										} else {
											if v203 < int32(3) {
												v285 = v200
												v295 = v285
											} else {
												v228 = int32(2)
												v230 = int32(128)
												v231 = v207 << (uint(v228) % 32) & v230
												if v231 != v211<<(uint(v228)%32)&v230 {
													v286 = v231
													if v286 != 0 {
														v289 = int32(1)
													} else {
														v289 = int32(-1)
													}
													v295 = v289
												} else {
													if v203 == int32(3) {
														v285 = v200
														v295 = v285
													} else {
														v239 = int32(3)
														v241 = int32(128)
														v242 = v207 << (uint(v239) % 32) & v241
														if v242 != v211<<(uint(v239)%32)&v241 {
															v286 = v242
															if v286 != 0 {
																v289 = int32(1)
															} else {
																v289 = int32(-1)
															}
															v295 = v289
														} else {
															if v203 < int32(5) {
																v285 = v200
																v295 = v285
															} else {
																v250 = int32(4)
																v252 = int32(128)
																v253 = v207 << (uint(v250) % 32) & v252
																if v253 != v211<<(uint(v250)%32)&v252 {
																	v286 = v253
																	if v286 != 0 {
																		v289 = int32(1)
																	} else {
																		v289 = int32(-1)
																	}
																	v295 = v289
																} else {
																	if v203 == int32(5) {
																		v285 = v200
																		v295 = v285
																	} else {
																		v261 = int32(5)
																		v263 = int32(128)
																		v264 = v207 << (uint(v261) % 32) & v263
																		if v264 != v211<<(uint(v261)%32)&v263 {
																			v286 = v264
																			if v286 != 0 {
																				v289 = int32(1)
																			} else {
																				v289 = int32(-1)
																			}
																			v295 = v289
																		} else {
																			if v203 < int32(7) {
																				v285 = v200
																				v295 = v285
																			} else {
																				v272 = int32(6)
																				v274 = int32(128)
																				v275 = v207 << (uint(v272) % 32) & v274
																				if v275 != v211<<(uint(v272)%32)&v274 {
																					v286 = v275
																					if v286 != 0 {
																						v289 = int32(1)
																					} else {
																						v289 = int32(-1)
																					}
																					v295 = v289
																				} else {
																					v285 = v200
																					v295 = v285
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
						v296 = v295
						return base.I64_extend_i32_u(base.B2i32(v296 == int32(0)))
					} else {
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
						v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
						if base.Ui32(v48) < base.Ui32(v47) {
							v184 = v17 + int32(4)
							v186 = v31 + int32(2)
							v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+3)))
							v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
							if base.Ui32(v187) < base.Ui32(v188) {
								v190 = v187
							} else {
								v190 = v188
							}
							v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
							if base.Ui32(v190) < base.Ui32(v191) {
								v193 = v190
							} else {
								v193 = v191
							}
							v198 = base.I32_div_s(v193, int32(8))
							v199 = F_memcmp(m, v184, v186, v198)
							mBase = m.M
							if v199 != 0 {
								v285 = v199
								v295 = v285
							} else {
								v200 = int32(0)
								v203 = v193 - v198<<(uint(int32(3))%32)
								if v203 <= v200 {
									v285 = v200
									v295 = v285
								} else {
									v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184+v198))))
									v208 = int32(128)
									v209 = v207 & v208
									v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186+v198))))
									if v209 != v211&v208 {
										v286 = v209
										if v286 != 0 {
											v289 = int32(1)
										} else {
											v289 = int32(-1)
										}
										v295 = v289
									} else {
										if v203 == int32(1) {
											v285 = v200
											v295 = v285
										} else {
											v217 = int32(1)
											v219 = int32(128)
											v220 = v207 << (uint(v217) % 32) & v219
											if v220 != v211<<(uint(v217)%32)&v219 {
												v286 = v220
												if v286 != 0 {
													v289 = int32(1)
												} else {
													v289 = int32(-1)
												}
												v295 = v289
											} else {
												if v203 < int32(3) {
													v285 = v200
													v295 = v285
												} else {
													v228 = int32(2)
													v230 = int32(128)
													v231 = v207 << (uint(v228) % 32) & v230
													if v231 != v211<<(uint(v228)%32)&v230 {
														v286 = v231
														if v286 != 0 {
															v289 = int32(1)
														} else {
															v289 = int32(-1)
														}
														v295 = v289
													} else {
														if v203 == int32(3) {
															v285 = v200
															v295 = v285
														} else {
															v239 = int32(3)
															v241 = int32(128)
															v242 = v207 << (uint(v239) % 32) & v241
															if v242 != v211<<(uint(v239)%32)&v241 {
																v286 = v242
																if v286 != 0 {
																	v289 = int32(1)
																} else {
																	v289 = int32(-1)
																}
																v295 = v289
															} else {
																if v203 < int32(5) {
																	v285 = v200
																	v295 = v285
																} else {
																	v250 = int32(4)
																	v252 = int32(128)
																	v253 = v207 << (uint(v250) % 32) & v252
																	if v253 != v211<<(uint(v250)%32)&v252 {
																		v286 = v253
																		if v286 != 0 {
																			v289 = int32(1)
																		} else {
																			v289 = int32(-1)
																		}
																		v295 = v289
																	} else {
																		if v203 == int32(5) {
																			v285 = v200
																			v295 = v285
																		} else {
																			v261 = int32(5)
																			v263 = int32(128)
																			v264 = v207 << (uint(v261) % 32) & v263
																			if v264 != v211<<(uint(v261)%32)&v263 {
																				v286 = v264
																				if v286 != 0 {
																					v289 = int32(1)
																				} else {
																					v289 = int32(-1)
																				}
																				v295 = v289
																			} else {
																				if v203 < int32(7) {
																					v285 = v200
																					v295 = v285
																				} else {
																					v272 = int32(6)
																					v274 = int32(128)
																					v275 = v207 << (uint(v272) % 32) & v274
																					if v275 != v211<<(uint(v272)%32)&v274 {
																						v286 = v275
																						if v286 != 0 {
																							v289 = int32(1)
																						} else {
																							v289 = int32(-1)
																						}
																						v295 = v289
																					} else {
																						v285 = v200
																						v295 = v285
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
							v296 = v295
							return base.I64_extend_i32_u(base.B2i32(v296 == int32(0)))
						} else {
							return int64(0)
						}
					}
				case 7:
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+16)))
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v51)+12)))
					if v53&int32(1) == int32(0) {
						v184 = v17 + int32(4)
						v186 = v31 + int32(2)
						v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+3)))
						v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
						if base.Ui32(v187) < base.Ui32(v188) {
							v190 = v187
						} else {
							v190 = v188
						}
						v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
						if base.Ui32(v190) < base.Ui32(v191) {
							v193 = v190
						} else {
							v193 = v191
						}
						v198 = base.I32_div_s(v193, int32(8))
						v199 = F_memcmp(m, v184, v186, v198)
						mBase = m.M
						if v199 != 0 {
							v285 = v199
							v295 = v285
						} else {
							v200 = int32(0)
							v203 = v193 - v198<<(uint(int32(3))%32)
							if v203 <= v200 {
								v285 = v200
								v295 = v285
							} else {
								v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184+v198))))
								v208 = int32(128)
								v209 = v207 & v208
								v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186+v198))))
								if v209 != v211&v208 {
									v286 = v209
									if v286 != 0 {
										v289 = int32(1)
									} else {
										v289 = int32(-1)
									}
									v295 = v289
								} else {
									if v203 == int32(1) {
										v285 = v200
										v295 = v285
									} else {
										v217 = int32(1)
										v219 = int32(128)
										v220 = v207 << (uint(v217) % 32) & v219
										if v220 != v211<<(uint(v217)%32)&v219 {
											v286 = v220
											if v286 != 0 {
												v289 = int32(1)
											} else {
												v289 = int32(-1)
											}
											v295 = v289
										} else {
											if v203 < int32(3) {
												v285 = v200
												v295 = v285
											} else {
												v228 = int32(2)
												v230 = int32(128)
												v231 = v207 << (uint(v228) % 32) & v230
												if v231 != v211<<(uint(v228)%32)&v230 {
													v286 = v231
													if v286 != 0 {
														v289 = int32(1)
													} else {
														v289 = int32(-1)
													}
													v295 = v289
												} else {
													if v203 == int32(3) {
														v285 = v200
														v295 = v285
													} else {
														v239 = int32(3)
														v241 = int32(128)
														v242 = v207 << (uint(v239) % 32) & v241
														if v242 != v211<<(uint(v239)%32)&v241 {
															v286 = v242
															if v286 != 0 {
																v289 = int32(1)
															} else {
																v289 = int32(-1)
															}
															v295 = v289
														} else {
															if v203 < int32(5) {
																v285 = v200
																v295 = v285
															} else {
																v250 = int32(4)
																v252 = int32(128)
																v253 = v207 << (uint(v250) % 32) & v252
																if v253 != v211<<(uint(v250)%32)&v252 {
																	v286 = v253
																	if v286 != 0 {
																		v289 = int32(1)
																	} else {
																		v289 = int32(-1)
																	}
																	v295 = v289
																} else {
																	if v203 == int32(5) {
																		v285 = v200
																		v295 = v285
																	} else {
																		v261 = int32(5)
																		v263 = int32(128)
																		v264 = v207 << (uint(v261) % 32) & v263
																		if v264 != v211<<(uint(v261)%32)&v263 {
																			v286 = v264
																			if v286 != 0 {
																				v289 = int32(1)
																			} else {
																				v289 = int32(-1)
																			}
																			v295 = v289
																		} else {
																			if v203 < int32(7) {
																				v285 = v200
																				v295 = v285
																			} else {
																				v272 = int32(6)
																				v274 = int32(128)
																				v275 = v207 << (uint(v272) % 32) & v274
																				if v275 != v211<<(uint(v272)%32)&v274 {
																					v286 = v275
																					if v286 != 0 {
																						v289 = int32(1)
																					} else {
																						v289 = int32(-1)
																					}
																					v295 = v289
																				} else {
																					v285 = v200
																					v295 = v285
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
						v296 = v295
						return base.I64_extend_i32_u(base.B2i32(v296 == int32(0)))
					} else {
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
						v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
						if base.Ui32(v59) <= base.Ui32(v58) {
							v184 = v17 + int32(4)
							v186 = v31 + int32(2)
							v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+3)))
							v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
							if base.Ui32(v187) < base.Ui32(v188) {
								v190 = v187
							} else {
								v190 = v188
							}
							v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
							if base.Ui32(v190) < base.Ui32(v191) {
								v193 = v190
							} else {
								v193 = v191
							}
							v198 = base.I32_div_s(v193, int32(8))
							v199 = F_memcmp(m, v184, v186, v198)
							mBase = m.M
							if v199 != 0 {
								v285 = v199
								v295 = v285
							} else {
								v200 = int32(0)
								v203 = v193 - v198<<(uint(int32(3))%32)
								if v203 <= v200 {
									v285 = v200
									v295 = v285
								} else {
									v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184+v198))))
									v208 = int32(128)
									v209 = v207 & v208
									v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186+v198))))
									if v209 != v211&v208 {
										v286 = v209
										if v286 != 0 {
											v289 = int32(1)
										} else {
											v289 = int32(-1)
										}
										v295 = v289
									} else {
										if v203 == int32(1) {
											v285 = v200
											v295 = v285
										} else {
											v217 = int32(1)
											v219 = int32(128)
											v220 = v207 << (uint(v217) % 32) & v219
											if v220 != v211<<(uint(v217)%32)&v219 {
												v286 = v220
												if v286 != 0 {
													v289 = int32(1)
												} else {
													v289 = int32(-1)
												}
												v295 = v289
											} else {
												if v203 < int32(3) {
													v285 = v200
													v295 = v285
												} else {
													v228 = int32(2)
													v230 = int32(128)
													v231 = v207 << (uint(v228) % 32) & v230
													if v231 != v211<<(uint(v228)%32)&v230 {
														v286 = v231
														if v286 != 0 {
															v289 = int32(1)
														} else {
															v289 = int32(-1)
														}
														v295 = v289
													} else {
														if v203 == int32(3) {
															v285 = v200
															v295 = v285
														} else {
															v239 = int32(3)
															v241 = int32(128)
															v242 = v207 << (uint(v239) % 32) & v241
															if v242 != v211<<(uint(v239)%32)&v241 {
																v286 = v242
																if v286 != 0 {
																	v289 = int32(1)
																} else {
																	v289 = int32(-1)
																}
																v295 = v289
															} else {
																if v203 < int32(5) {
																	v285 = v200
																	v295 = v285
																} else {
																	v250 = int32(4)
																	v252 = int32(128)
																	v253 = v207 << (uint(v250) % 32) & v252
																	if v253 != v211<<(uint(v250)%32)&v252 {
																		v286 = v253
																		if v286 != 0 {
																			v289 = int32(1)
																		} else {
																			v289 = int32(-1)
																		}
																		v295 = v289
																	} else {
																		if v203 == int32(5) {
																			v285 = v200
																			v295 = v285
																		} else {
																			v261 = int32(5)
																			v263 = int32(128)
																			v264 = v207 << (uint(v261) % 32) & v263
																			if v264 != v211<<(uint(v261)%32)&v263 {
																				v286 = v264
																				if v286 != 0 {
																					v289 = int32(1)
																				} else {
																					v289 = int32(-1)
																				}
																				v295 = v289
																			} else {
																				if v203 < int32(7) {
																					v285 = v200
																					v295 = v285
																				} else {
																					v272 = int32(6)
																					v274 = int32(128)
																					v275 = v207 << (uint(v272) % 32) & v274
																					if v275 != v211<<(uint(v272)%32)&v274 {
																						v286 = v275
																						if v286 != 0 {
																							v289 = int32(1)
																						} else {
																							v289 = int32(-1)
																						}
																						v295 = v289
																					} else {
																						v285 = v200
																						v295 = v285
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
							v296 = v295
							return base.I64_extend_i32_u(base.B2i32(v296 == int32(0)))
						} else {
							return int64(0)
						}
					}
				case 8:
					v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
					v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
					if base.Ui32(v64) < base.Ui32(v65) {
						v184 = v17 + int32(4)
						v186 = v31 + int32(2)
						v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+3)))
						v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
						if base.Ui32(v187) < base.Ui32(v188) {
							v190 = v187
						} else {
							v190 = v188
						}
						v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
						if base.Ui32(v190) < base.Ui32(v191) {
							v193 = v190
						} else {
							v193 = v191
						}
						v198 = base.I32_div_s(v193, int32(8))
						v199 = F_memcmp(m, v184, v186, v198)
						mBase = m.M
						if v199 != 0 {
							v285 = v199
							v295 = v285
						} else {
							v200 = int32(0)
							v203 = v193 - v198<<(uint(int32(3))%32)
							if v203 <= v200 {
								v285 = v200
								v295 = v285
							} else {
								v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184+v198))))
								v208 = int32(128)
								v209 = v207 & v208
								v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186+v198))))
								if v209 != v211&v208 {
									v286 = v209
									if v286 != 0 {
										v289 = int32(1)
									} else {
										v289 = int32(-1)
									}
									v295 = v289
								} else {
									if v203 == int32(1) {
										v285 = v200
										v295 = v285
									} else {
										v217 = int32(1)
										v219 = int32(128)
										v220 = v207 << (uint(v217) % 32) & v219
										if v220 != v211<<(uint(v217)%32)&v219 {
											v286 = v220
											if v286 != 0 {
												v289 = int32(1)
											} else {
												v289 = int32(-1)
											}
											v295 = v289
										} else {
											if v203 < int32(3) {
												v285 = v200
												v295 = v285
											} else {
												v228 = int32(2)
												v230 = int32(128)
												v231 = v207 << (uint(v228) % 32) & v230
												if v231 != v211<<(uint(v228)%32)&v230 {
													v286 = v231
													if v286 != 0 {
														v289 = int32(1)
													} else {
														v289 = int32(-1)
													}
													v295 = v289
												} else {
													if v203 == int32(3) {
														v285 = v200
														v295 = v285
													} else {
														v239 = int32(3)
														v241 = int32(128)
														v242 = v207 << (uint(v239) % 32) & v241
														if v242 != v211<<(uint(v239)%32)&v241 {
															v286 = v242
															if v286 != 0 {
																v289 = int32(1)
															} else {
																v289 = int32(-1)
															}
															v295 = v289
														} else {
															if v203 < int32(5) {
																v285 = v200
																v295 = v285
															} else {
																v250 = int32(4)
																v252 = int32(128)
																v253 = v207 << (uint(v250) % 32) & v252
																if v253 != v211<<(uint(v250)%32)&v252 {
																	v286 = v253
																	if v286 != 0 {
																		v289 = int32(1)
																	} else {
																		v289 = int32(-1)
																	}
																	v295 = v289
																} else {
																	if v203 == int32(5) {
																		v285 = v200
																		v295 = v285
																	} else {
																		v261 = int32(5)
																		v263 = int32(128)
																		v264 = v207 << (uint(v261) % 32) & v263
																		if v264 != v211<<(uint(v261)%32)&v263 {
																			v286 = v264
																			if v286 != 0 {
																				v289 = int32(1)
																			} else {
																				v289 = int32(-1)
																			}
																			v295 = v289
																		} else {
																			if v203 < int32(7) {
																				v285 = v200
																				v295 = v285
																			} else {
																				v272 = int32(6)
																				v274 = int32(128)
																				v275 = v207 << (uint(v272) % 32) & v274
																				if v275 != v211<<(uint(v272)%32)&v274 {
																					v286 = v275
																					if v286 != 0 {
																						v289 = int32(1)
																					} else {
																						v289 = int32(-1)
																					}
																					v295 = v289
																				} else {
																					v285 = v200
																					v295 = v285
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
						v296 = v295
						return base.I64_extend_i32_u(base.B2i32(v296 == int32(0)))
					} else {
						return int64(0)
					}
				}
			}
		}
	}
}
