package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inet_gist_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
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
	var v78 int32
	_ = v78
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
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
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
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
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
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
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v775 int32
	_ = v775
	var v782 int32
	_ = v782
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v21 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v21)
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
		if v23 == v21 {
			return int32(1)
		} else {
			v28 = int32(1)
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			v33 = v31 & v28
			if v33 != 0 {
				v34 = v28
			} else {
				v34 = int32(4)
			}
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v34))))
			if v36 != v23 {
				switch v18&int32(_a_F_inet_gist_consistent_0) - int32(19) {
				case 0:
					v768 = v28
					return v768
				case 1, 2:
					if base.Ui32(v36) <= base.Ui32(v23) {
						return int32(0)
					} else {
						v768 = v28
						return v768
					}
				case 3, 4:
					if base.Ui32(v36) < base.Ui32(v23) {
						v768 = v28
						return v768
					} else {
						return int32(0)
					}
				default:
					return int32(0)
				}
			} else {
				v44 = v14 + int32(1)
				v46 = v14 + int32(4)
				if v33 != 0 {
					v47 = v44
				} else {
					v47 = v46
				}
				switch v18&int32(_a_F_inet_gist_consistent_0) - int32(18) {
				case 0, 9:
					v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
					v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
					if base.Ui32(v74) <= base.Ui32(v75) {
						v82 = v74
						v83 = v75
						v85 = v19 + int32(4)
						v87 = v47 + int32(2)
						v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+3)))
						if base.Ui32(v88) < base.Ui32(v82) {
							v90 = v88
						} else {
							v90 = v82
						}
						if base.Ui32(v90) < base.Ui32(v83) {
							v92 = v90
						} else {
							v92 = v83
						}
						v97 = base.I32_div_s(v92, int32(8))
						v98 = F_memcmp(m, v85, v87, v97)
						mBase = m.M
						if v98 != 0 {
							v184 = v98
							v194 = v184
						} else {
							v99 = int32(0)
							v102 = v92 - v97<<(uint(int32(3))%32)
							if v102 <= v99 {
								v184 = v99
								v194 = v184
							} else {
								v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v97))))
								v107 = int32(128)
								v108 = v106 & v107
								v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+v97))))
								if v108 != v110&v107 {
									v185 = v108
									if v185 != 0 {
										v188 = int32(1)
									} else {
										v188 = int32(-1)
									}
									v194 = v188
								} else {
									if v102 == int32(1) {
										v184 = v99
										v194 = v184
									} else {
										v116 = int32(1)
										v118 = int32(128)
										v119 = v106 << (uint(v116) % 32) & v118
										if v119 != v110<<(uint(v116)%32)&v118 {
											v185 = v119
											if v185 != 0 {
												v188 = int32(1)
											} else {
												v188 = int32(-1)
											}
											v194 = v188
										} else {
											if v102 < int32(3) {
												v184 = v99
												v194 = v184
											} else {
												v127 = int32(2)
												v129 = int32(128)
												v130 = v106 << (uint(v127) % 32) & v129
												if v130 != v110<<(uint(v127)%32)&v129 {
													v185 = v130
													if v185 != 0 {
														v188 = int32(1)
													} else {
														v188 = int32(-1)
													}
													v194 = v188
												} else {
													if v102 == int32(3) {
														v184 = v99
														v194 = v184
													} else {
														v138 = int32(3)
														v140 = int32(128)
														v141 = v106 << (uint(v138) % 32) & v140
														if v141 != v110<<(uint(v138)%32)&v140 {
															v185 = v141
															if v185 != 0 {
																v188 = int32(1)
															} else {
																v188 = int32(-1)
															}
															v194 = v188
														} else {
															if v102 < int32(5) {
																v184 = v99
																v194 = v184
															} else {
																v149 = int32(4)
																v151 = int32(128)
																v152 = v106 << (uint(v149) % 32) & v151
																if v152 != v110<<(uint(v149)%32)&v151 {
																	v185 = v152
																	if v185 != 0 {
																		v188 = int32(1)
																	} else {
																		v188 = int32(-1)
																	}
																	v194 = v188
																} else {
																	if v102 == int32(5) {
																		v184 = v99
																		v194 = v184
																	} else {
																		v160 = int32(5)
																		v162 = int32(128)
																		v163 = v106 << (uint(v160) % 32) & v162
																		if v163 != v110<<(uint(v160)%32)&v162 {
																			v185 = v163
																			if v185 != 0 {
																				v188 = int32(1)
																			} else {
																				v188 = int32(-1)
																			}
																			v194 = v188
																		} else {
																			if v102 < int32(7) {
																				v184 = v99
																				v194 = v184
																			} else {
																				v171 = int32(6)
																				v173 = int32(128)
																				v174 = v106 << (uint(v171) % 32) & v173
																				if v174 != v110<<(uint(v171)%32)&v173 {
																					v185 = v174
																					if v185 != 0 {
																						v188 = int32(1)
																					} else {
																						v188 = int32(-1)
																					}
																					v194 = v188
																				} else {
																					v184 = v99
																					v194 = v184
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
						switch v18&int32(_a_F_inet_gist_consistent_0) - int32(3) {
						case 0, 21, 22, 23, 24:
							v313 = v194
							return base.B2i32(v313 == int32(0))
						default:
							v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
							v618 = v615 & int32(1)
							if v618 != 0 {
								v622 = v44
							} else {
								v622 = v46
							}
							v624 = v622 + int32(2)
							v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
							if v627 == int32(3) {
								v630 = int32(128)
							} else {
								v630 = int32(32)
							}
							v635 = base.I32_div_s(v630, int32(8))
							v636 = F_memcmp(m, v85, v624, v635)
							mBase = m.M
							if v636 != 0 {
								v722 = v636
								v732 = v722
							} else {
								v637 = int32(0)
								v640 = v630 - v635<<(uint(int32(3))%32)
								if v640 <= v637 {
									v722 = v637
									v732 = v722
								} else {
									v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v635))))
									v645 = int32(128)
									v646 = v644 & v645
									v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624+v635))))
									if v646 != v648&v645 {
										v723 = v646
										if v723 != 0 {
											v726 = int32(1)
										} else {
											v726 = int32(-1)
										}
										v732 = v726
									} else {
										if v640 == int32(1) {
											v722 = v637
											v732 = v722
										} else {
											v654 = int32(1)
											v656 = int32(128)
											v657 = v644 << (uint(v654) % 32) & v656
											if v657 != v648<<(uint(v654)%32)&v656 {
												v723 = v657
												if v723 != 0 {
													v726 = int32(1)
												} else {
													v726 = int32(-1)
												}
												v732 = v726
											} else {
												if v640 < int32(3) {
													v722 = v637
													v732 = v722
												} else {
													v665 = int32(2)
													v667 = int32(128)
													v668 = v644 << (uint(v665) % 32) & v667
													if v668 != v648<<(uint(v665)%32)&v667 {
														v723 = v668
														if v723 != 0 {
															v726 = int32(1)
														} else {
															v726 = int32(-1)
														}
														v732 = v726
													} else {
														if v640 == int32(3) {
															v722 = v637
															v732 = v722
														} else {
															v676 = int32(3)
															v678 = int32(128)
															v679 = v644 << (uint(v676) % 32) & v678
															if v679 != v648<<(uint(v676)%32)&v678 {
																v723 = v679
																if v723 != 0 {
																	v726 = int32(1)
																} else {
																	v726 = int32(-1)
																}
																v732 = v726
															} else {
																if v640 < int32(5) {
																	v722 = v637
																	v732 = v722
																} else {
																	v687 = int32(4)
																	v689 = int32(128)
																	v690 = v644 << (uint(v687) % 32) & v689
																	if v690 != v648<<(uint(v687)%32)&v689 {
																		v723 = v690
																		if v723 != 0 {
																			v726 = int32(1)
																		} else {
																			v726 = int32(-1)
																		}
																		v732 = v726
																	} else {
																		if v640 == int32(5) {
																			v722 = v637
																			v732 = v722
																		} else {
																			v698 = int32(5)
																			v700 = int32(128)
																			v701 = v644 << (uint(v698) % 32) & v700
																			if v701 != v648<<(uint(v698)%32)&v700 {
																				v723 = v701
																				if v723 != 0 {
																					v726 = int32(1)
																				} else {
																					v726 = int32(-1)
																				}
																				v732 = v726
																			} else {
																				if v640 < int32(7) {
																					v722 = v637
																					v732 = v722
																				} else {
																					v709 = int32(6)
																					v711 = int32(128)
																					v712 = v644 << (uint(v709) % 32) & v711
																					if v712 != v648<<(uint(v709)%32)&v711 {
																						v723 = v712
																						if v723 != 0 {
																							v726 = int32(1)
																						} else {
																							v726 = int32(-1)
																						}
																						v732 = v726
																					} else {
																						v722 = v637
																						v732 = v722
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
							switch v18&int32(_a_F_inet_gist_consistent_0) - int32(18) {
							case 0:
								v782 = v732
								return base.B2i32(v782 == int32(0))
							case 1:
								v775 = v732
								return base.B2i32(v775 != int32(0))
							case 2:
								return int32(base.Ui32(v732) >> (uint(int32(31)) % 32))
							case 3:
								return base.B2i32(v732 <= int32(0))
							case 4:
								return base.B2i32(int32(0) < v732)
							case 5:
								return int32(base.Ui32(v732^int32(-1)) >> (uint(int32(31)) % 32))
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v754 = m.ExcPending
								if v754 != 0 {
									return int32(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_1), int32(0))
									mBase = m.M
									v758 = m.ExcPending
									if v758 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_inet_gist_consistent_2), int32(327), int32(_a_F_inet_gist_consistent_3))
										mBase = m.M
										v763 = m.ExcPending
										if v763 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						case 15:
							if v194 != 0 {
								v768 = int32(0)
								return v768
							} else {
								v335 = int32(1)
								v336 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								v337 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v336)+16)))
								v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336+v337)+12)))
								if v339&v335 == int32(0) {
									v768 = v335
									return v768
								} else {
									v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
									v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
									if v346&int32(1) != 0 {
										v349 = v44
									} else {
										v349 = v46
									}
									v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+1)))
									if v345 != v350 {
										v768 = int32(0)
										return v768
									} else {
										v353 = v349 + int32(2)
										v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
										if v356 == int32(3) {
											v359 = int32(128)
										} else {
											v359 = int32(32)
										}
										v364 = base.I32_div_s(v359, int32(8))
										v365 = F_memcmp(m, v85, v353, v364)
										mBase = m.M
										if v365 != 0 {
											v451 = v365
											v461 = v451
										} else {
											v366 = int32(0)
											v369 = v359 - v364<<(uint(int32(3))%32)
											if v369 <= v366 {
												v451 = v366
												v461 = v451
											} else {
												v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v364))))
												v374 = int32(128)
												v375 = v373 & v374
												v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353+v364))))
												if v375 != v377&v374 {
													v452 = v375
													if v452 != 0 {
														v455 = int32(1)
													} else {
														v455 = int32(-1)
													}
													v461 = v455
												} else {
													if v369 == int32(1) {
														v451 = v366
														v461 = v451
													} else {
														v383 = int32(1)
														v385 = int32(128)
														v386 = v373 << (uint(v383) % 32) & v385
														if v386 != v377<<(uint(v383)%32)&v385 {
															v452 = v386
															if v452 != 0 {
																v455 = int32(1)
															} else {
																v455 = int32(-1)
															}
															v461 = v455
														} else {
															if v369 < int32(3) {
																v451 = v366
																v461 = v451
															} else {
																v394 = int32(2)
																v396 = int32(128)
																v397 = v373 << (uint(v394) % 32) & v396
																if v397 != v377<<(uint(v394)%32)&v396 {
																	v452 = v397
																	if v452 != 0 {
																		v455 = int32(1)
																	} else {
																		v455 = int32(-1)
																	}
																	v461 = v455
																} else {
																	if v369 == int32(3) {
																		v451 = v366
																		v461 = v451
																	} else {
																		v405 = int32(3)
																		v407 = int32(128)
																		v408 = v373 << (uint(v405) % 32) & v407
																		if v408 != v377<<(uint(v405)%32)&v407 {
																			v452 = v408
																			if v452 != 0 {
																				v455 = int32(1)
																			} else {
																				v455 = int32(-1)
																			}
																			v461 = v455
																		} else {
																			if v369 < int32(5) {
																				v451 = v366
																				v461 = v451
																			} else {
																				v416 = int32(4)
																				v418 = int32(128)
																				v419 = v373 << (uint(v416) % 32) & v418
																				if v419 != v377<<(uint(v416)%32)&v418 {
																					v452 = v419
																					if v452 != 0 {
																						v455 = int32(1)
																					} else {
																						v455 = int32(-1)
																					}
																					v461 = v455
																				} else {
																					if v369 == int32(5) {
																						v451 = v366
																						v461 = v451
																					} else {
																						v427 = int32(5)
																						v429 = int32(128)
																						v430 = v373 << (uint(v427) % 32) & v429
																						if v430 != v377<<(uint(v427)%32)&v429 {
																							v452 = v430
																							if v452 != 0 {
																								v455 = int32(1)
																							} else {
																								v455 = int32(-1)
																							}
																							v461 = v455
																						} else {
																							if v369 < int32(7) {
																								v451 = v366
																								v461 = v451
																							} else {
																								v438 = int32(6)
																								v440 = int32(128)
																								v441 = v373 << (uint(v438) % 32) & v440
																								if v441 != v377<<(uint(v438)%32)&v440 {
																									v452 = v441
																									if v452 != 0 {
																										v455 = int32(1)
																									} else {
																										v455 = int32(-1)
																									}
																									v461 = v455
																								} else {
																									v451 = v366
																									v461 = v451
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
										v782 = v461
										return base.B2i32(v782 == int32(0))
									}
								}
							}
						case 16:
							if v194 != 0 {
								v768 = v28
								return v768
							} else {
								v488 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								v489 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v488)+16)))
								v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488+v489)+12)))
								if v491&int32(1) == int32(0) {
									v768 = v28
									return v768
								} else {
									v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
									v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
									if v497&int32(1) != 0 {
										v500 = v44
									} else {
										v500 = v46
									}
									v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+1)))
									if v496 != v501 {
										v768 = v28
										return v768
									} else {
										v504 = v500 + int32(2)
										v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
										if v507 == int32(3) {
											v510 = int32(128)
										} else {
											v510 = int32(32)
										}
										v515 = base.I32_div_s(v510, int32(8))
										v516 = F_memcmp(m, v85, v504, v515)
										mBase = m.M
										if v516 != 0 {
											v602 = v516
											v612 = v602
										} else {
											v517 = int32(0)
											v520 = v510 - v515<<(uint(int32(3))%32)
											if v520 <= v517 {
												v602 = v517
												v612 = v602
											} else {
												v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v515))))
												v525 = int32(128)
												v526 = v524 & v525
												v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504+v515))))
												if v526 != v528&v525 {
													v603 = v526
													if v603 != 0 {
														v606 = int32(1)
													} else {
														v606 = int32(-1)
													}
													v612 = v606
												} else {
													if v520 == int32(1) {
														v602 = v517
														v612 = v602
													} else {
														v534 = int32(1)
														v536 = int32(128)
														v537 = v524 << (uint(v534) % 32) & v536
														if v537 != v528<<(uint(v534)%32)&v536 {
															v603 = v537
															if v603 != 0 {
																v606 = int32(1)
															} else {
																v606 = int32(-1)
															}
															v612 = v606
														} else {
															if v520 < int32(3) {
																v602 = v517
																v612 = v602
															} else {
																v545 = int32(2)
																v547 = int32(128)
																v548 = v524 << (uint(v545) % 32) & v547
																if v548 != v528<<(uint(v545)%32)&v547 {
																	v603 = v548
																	if v603 != 0 {
																		v606 = int32(1)
																	} else {
																		v606 = int32(-1)
																	}
																	v612 = v606
																} else {
																	if v520 == int32(3) {
																		v602 = v517
																		v612 = v602
																	} else {
																		v556 = int32(3)
																		v558 = int32(128)
																		v559 = v524 << (uint(v556) % 32) & v558
																		if v559 != v528<<(uint(v556)%32)&v558 {
																			v603 = v559
																			if v603 != 0 {
																				v606 = int32(1)
																			} else {
																				v606 = int32(-1)
																			}
																			v612 = v606
																		} else {
																			if v520 < int32(5) {
																				v602 = v517
																				v612 = v602
																			} else {
																				v567 = int32(4)
																				v569 = int32(128)
																				v570 = v524 << (uint(v567) % 32) & v569
																				if v570 != v528<<(uint(v567)%32)&v569 {
																					v603 = v570
																					if v603 != 0 {
																						v606 = int32(1)
																					} else {
																						v606 = int32(-1)
																					}
																					v612 = v606
																				} else {
																					if v520 == int32(5) {
																						v602 = v517
																						v612 = v602
																					} else {
																						v578 = int32(5)
																						v580 = int32(128)
																						v581 = v524 << (uint(v578) % 32) & v580
																						if v581 != v528<<(uint(v578)%32)&v580 {
																							v603 = v581
																							if v603 != 0 {
																								v606 = int32(1)
																							} else {
																								v606 = int32(-1)
																							}
																							v612 = v606
																						} else {
																							if v520 < int32(7) {
																								v602 = v517
																								v612 = v602
																							} else {
																								v589 = int32(6)
																								v591 = int32(128)
																								v592 = v524 << (uint(v589) % 32) & v591
																								if v592 != v528<<(uint(v589)%32)&v591 {
																									v603 = v592
																									if v603 != 0 {
																										v606 = int32(1)
																									} else {
																										v606 = int32(-1)
																									}
																									v612 = v606
																								} else {
																									v602 = v517
																									v612 = v602
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
										v775 = v612
										return base.B2i32(v775 != int32(0))
									}
								}
							}
						case 17, 18:
							v322 = int32(0)
							if v322 < v194 {
								v768 = v322
								return v768
							} else {
								v325 = int32(1)
								if v194 < int32(0) {
									v768 = v325
									return v768
								} else {
									v328 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									v329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v328)+16)))
									v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328+v329)+12)))
									if v331&int32(1) != 0 {
										v475 = v325
										v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
										v478 = v476 & int32(1)
										if v478 != 0 {
											v479 = v44
										} else {
											v479 = v46
										}
										v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479)+1)))
										v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
										if v18&int32(_a_F_inet_gist_consistent_4) != int32(20) {
											if base.Ui32(v480) < base.Ui32(v481) {
												v768 = v475
												return v768
											} else {
												if base.Ui32(v480) <= base.Ui32(v481) {
													v618 = v478
													if v618 != 0 {
														v622 = v44
													} else {
														v622 = v46
													}
													v624 = v622 + int32(2)
													v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
													if v627 == int32(3) {
														v630 = int32(128)
													} else {
														v630 = int32(32)
													}
													v635 = base.I32_div_s(v630, int32(8))
													v636 = F_memcmp(m, v85, v624, v635)
													mBase = m.M
													if v636 != 0 {
														v722 = v636
														v732 = v722
													} else {
														v637 = int32(0)
														v640 = v630 - v635<<(uint(int32(3))%32)
														if v640 <= v637 {
															v722 = v637
															v732 = v722
														} else {
															v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v635))))
															v645 = int32(128)
															v646 = v644 & v645
															v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624+v635))))
															if v646 != v648&v645 {
																v723 = v646
																if v723 != 0 {
																	v726 = int32(1)
																} else {
																	v726 = int32(-1)
																}
																v732 = v726
															} else {
																if v640 == int32(1) {
																	v722 = v637
																	v732 = v722
																} else {
																	v654 = int32(1)
																	v656 = int32(128)
																	v657 = v644 << (uint(v654) % 32) & v656
																	if v657 != v648<<(uint(v654)%32)&v656 {
																		v723 = v657
																		if v723 != 0 {
																			v726 = int32(1)
																		} else {
																			v726 = int32(-1)
																		}
																		v732 = v726
																	} else {
																		if v640 < int32(3) {
																			v722 = v637
																			v732 = v722
																		} else {
																			v665 = int32(2)
																			v667 = int32(128)
																			v668 = v644 << (uint(v665) % 32) & v667
																			if v668 != v648<<(uint(v665)%32)&v667 {
																				v723 = v668
																				if v723 != 0 {
																					v726 = int32(1)
																				} else {
																					v726 = int32(-1)
																				}
																				v732 = v726
																			} else {
																				if v640 == int32(3) {
																					v722 = v637
																					v732 = v722
																				} else {
																					v676 = int32(3)
																					v678 = int32(128)
																					v679 = v644 << (uint(v676) % 32) & v678
																					if v679 != v648<<(uint(v676)%32)&v678 {
																						v723 = v679
																						if v723 != 0 {
																							v726 = int32(1)
																						} else {
																							v726 = int32(-1)
																						}
																						v732 = v726
																					} else {
																						if v640 < int32(5) {
																							v722 = v637
																							v732 = v722
																						} else {
																							v687 = int32(4)
																							v689 = int32(128)
																							v690 = v644 << (uint(v687) % 32) & v689
																							if v690 != v648<<(uint(v687)%32)&v689 {
																								v723 = v690
																								if v723 != 0 {
																									v726 = int32(1)
																								} else {
																									v726 = int32(-1)
																								}
																								v732 = v726
																							} else {
																								if v640 == int32(5) {
																									v722 = v637
																									v732 = v722
																								} else {
																									v698 = int32(5)
																									v700 = int32(128)
																									v701 = v644 << (uint(v698) % 32) & v700
																									if v701 != v648<<(uint(v698)%32)&v700 {
																										v723 = v701
																										if v723 != 0 {
																											v726 = int32(1)
																										} else {
																											v726 = int32(-1)
																										}
																										v732 = v726
																									} else {
																										if v640 < int32(7) {
																											v722 = v637
																											v732 = v722
																										} else {
																											v709 = int32(6)
																											v711 = int32(128)
																											v712 = v644 << (uint(v709) % 32) & v711
																											if v712 != v648<<(uint(v709)%32)&v711 {
																												v723 = v712
																												if v723 != 0 {
																													v726 = int32(1)
																												} else {
																													v726 = int32(-1)
																												}
																												v732 = v726
																											} else {
																												v722 = v637
																												v732 = v722
																											}
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
													switch v18&int32(_a_F_inet_gist_consistent_0) - int32(18) {
													case 0:
														v782 = v732
														return base.B2i32(v782 == int32(0))
													case 1:
														v775 = v732
														return base.B2i32(v775 != int32(0))
													case 2:
														return int32(base.Ui32(v732) >> (uint(int32(31)) % 32))
													case 3:
														return base.B2i32(v732 <= int32(0))
													case 4:
														return base.B2i32(int32(0) < v732)
													case 5:
														return int32(base.Ui32(v732^int32(-1)) >> (uint(int32(31)) % 32))
													default:
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v754 = m.ExcPending
														if v754 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_1), int32(0))
															mBase = m.M
															v758 = m.ExcPending
															if v758 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_inet_gist_consistent_2), int32(327), int32(_a_F_inet_gist_consistent_3))
																mBase = m.M
																v763 = m.ExcPending
																if v763 != 0 {
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
													return int32(0)
												}
											}
										} else {
											if base.Ui32(v481) < base.Ui32(v480) {
												v768 = v475
												return v768
											} else {
												if base.Ui32(v481) <= base.Ui32(v480) {
													v618 = v478
													if v618 != 0 {
														v622 = v44
													} else {
														v622 = v46
													}
													v624 = v622 + int32(2)
													v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
													if v627 == int32(3) {
														v630 = int32(128)
													} else {
														v630 = int32(32)
													}
													v635 = base.I32_div_s(v630, int32(8))
													v636 = F_memcmp(m, v85, v624, v635)
													mBase = m.M
													if v636 != 0 {
														v722 = v636
														v732 = v722
													} else {
														v637 = int32(0)
														v640 = v630 - v635<<(uint(int32(3))%32)
														if v640 <= v637 {
															v722 = v637
															v732 = v722
														} else {
															v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v635))))
															v645 = int32(128)
															v646 = v644 & v645
															v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624+v635))))
															if v646 != v648&v645 {
																v723 = v646
																if v723 != 0 {
																	v726 = int32(1)
																} else {
																	v726 = int32(-1)
																}
																v732 = v726
															} else {
																if v640 == int32(1) {
																	v722 = v637
																	v732 = v722
																} else {
																	v654 = int32(1)
																	v656 = int32(128)
																	v657 = v644 << (uint(v654) % 32) & v656
																	if v657 != v648<<(uint(v654)%32)&v656 {
																		v723 = v657
																		if v723 != 0 {
																			v726 = int32(1)
																		} else {
																			v726 = int32(-1)
																		}
																		v732 = v726
																	} else {
																		if v640 < int32(3) {
																			v722 = v637
																			v732 = v722
																		} else {
																			v665 = int32(2)
																			v667 = int32(128)
																			v668 = v644 << (uint(v665) % 32) & v667
																			if v668 != v648<<(uint(v665)%32)&v667 {
																				v723 = v668
																				if v723 != 0 {
																					v726 = int32(1)
																				} else {
																					v726 = int32(-1)
																				}
																				v732 = v726
																			} else {
																				if v640 == int32(3) {
																					v722 = v637
																					v732 = v722
																				} else {
																					v676 = int32(3)
																					v678 = int32(128)
																					v679 = v644 << (uint(v676) % 32) & v678
																					if v679 != v648<<(uint(v676)%32)&v678 {
																						v723 = v679
																						if v723 != 0 {
																							v726 = int32(1)
																						} else {
																							v726 = int32(-1)
																						}
																						v732 = v726
																					} else {
																						if v640 < int32(5) {
																							v722 = v637
																							v732 = v722
																						} else {
																							v687 = int32(4)
																							v689 = int32(128)
																							v690 = v644 << (uint(v687) % 32) & v689
																							if v690 != v648<<(uint(v687)%32)&v689 {
																								v723 = v690
																								if v723 != 0 {
																									v726 = int32(1)
																								} else {
																									v726 = int32(-1)
																								}
																								v732 = v726
																							} else {
																								if v640 == int32(5) {
																									v722 = v637
																									v732 = v722
																								} else {
																									v698 = int32(5)
																									v700 = int32(128)
																									v701 = v644 << (uint(v698) % 32) & v700
																									if v701 != v648<<(uint(v698)%32)&v700 {
																										v723 = v701
																										if v723 != 0 {
																											v726 = int32(1)
																										} else {
																											v726 = int32(-1)
																										}
																										v732 = v726
																									} else {
																										if v640 < int32(7) {
																											v722 = v637
																											v732 = v722
																										} else {
																											v709 = int32(6)
																											v711 = int32(128)
																											v712 = v644 << (uint(v709) % 32) & v711
																											if v712 != v648<<(uint(v709)%32)&v711 {
																												v723 = v712
																												if v723 != 0 {
																													v726 = int32(1)
																												} else {
																													v726 = int32(-1)
																												}
																												v732 = v726
																											} else {
																												v722 = v637
																												v732 = v722
																											}
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
													switch v18&int32(_a_F_inet_gist_consistent_0) - int32(18) {
													case 0:
														v782 = v732
														return base.B2i32(v782 == int32(0))
													case 1:
														v775 = v732
														return base.B2i32(v775 != int32(0))
													case 2:
														return int32(base.Ui32(v732) >> (uint(int32(31)) % 32))
													case 3:
														return base.B2i32(v732 <= int32(0))
													case 4:
														return base.B2i32(int32(0) < v732)
													case 5:
														return int32(base.Ui32(v732^int32(-1)) >> (uint(int32(31)) % 32))
													default:
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v754 = m.ExcPending
														if v754 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_1), int32(0))
															mBase = m.M
															v758 = m.ExcPending
															if v758 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_inet_gist_consistent_2), int32(327), int32(_a_F_inet_gist_consistent_3))
																mBase = m.M
																v763 = m.ExcPending
																if v763 != 0 {
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
													return int32(0)
												}
											}
										}
									} else {
										v768 = v325
										return v768
									}
								}
							}
						case 19, 20:
							v462 = int32(0)
							if v194 < v462 {
								v768 = v462
								return v768
							} else {
								v465 = int32(1)
								if v194 != 0 {
									v768 = v465
									return v768
								} else {
									v466 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									v467 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v466)+16)))
									v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v466+v467)+12)))
									if v469&int32(1) == int32(0) {
										v768 = v465
										return v768
									} else {
										v475 = v465
										v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
										v478 = v476 & int32(1)
										if v478 != 0 {
											v479 = v44
										} else {
											v479 = v46
										}
										v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479)+1)))
										v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
										if v18&int32(_a_F_inet_gist_consistent_4) != int32(20) {
											if base.Ui32(v480) < base.Ui32(v481) {
												v768 = v475
												return v768
											} else {
												if base.Ui32(v480) <= base.Ui32(v481) {
													v618 = v478
													if v618 != 0 {
														v622 = v44
													} else {
														v622 = v46
													}
													v624 = v622 + int32(2)
													v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
													if v627 == int32(3) {
														v630 = int32(128)
													} else {
														v630 = int32(32)
													}
													v635 = base.I32_div_s(v630, int32(8))
													v636 = F_memcmp(m, v85, v624, v635)
													mBase = m.M
													if v636 != 0 {
														v722 = v636
														v732 = v722
													} else {
														v637 = int32(0)
														v640 = v630 - v635<<(uint(int32(3))%32)
														if v640 <= v637 {
															v722 = v637
															v732 = v722
														} else {
															v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v635))))
															v645 = int32(128)
															v646 = v644 & v645
															v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624+v635))))
															if v646 != v648&v645 {
																v723 = v646
																if v723 != 0 {
																	v726 = int32(1)
																} else {
																	v726 = int32(-1)
																}
																v732 = v726
															} else {
																if v640 == int32(1) {
																	v722 = v637
																	v732 = v722
																} else {
																	v654 = int32(1)
																	v656 = int32(128)
																	v657 = v644 << (uint(v654) % 32) & v656
																	if v657 != v648<<(uint(v654)%32)&v656 {
																		v723 = v657
																		if v723 != 0 {
																			v726 = int32(1)
																		} else {
																			v726 = int32(-1)
																		}
																		v732 = v726
																	} else {
																		if v640 < int32(3) {
																			v722 = v637
																			v732 = v722
																		} else {
																			v665 = int32(2)
																			v667 = int32(128)
																			v668 = v644 << (uint(v665) % 32) & v667
																			if v668 != v648<<(uint(v665)%32)&v667 {
																				v723 = v668
																				if v723 != 0 {
																					v726 = int32(1)
																				} else {
																					v726 = int32(-1)
																				}
																				v732 = v726
																			} else {
																				if v640 == int32(3) {
																					v722 = v637
																					v732 = v722
																				} else {
																					v676 = int32(3)
																					v678 = int32(128)
																					v679 = v644 << (uint(v676) % 32) & v678
																					if v679 != v648<<(uint(v676)%32)&v678 {
																						v723 = v679
																						if v723 != 0 {
																							v726 = int32(1)
																						} else {
																							v726 = int32(-1)
																						}
																						v732 = v726
																					} else {
																						if v640 < int32(5) {
																							v722 = v637
																							v732 = v722
																						} else {
																							v687 = int32(4)
																							v689 = int32(128)
																							v690 = v644 << (uint(v687) % 32) & v689
																							if v690 != v648<<(uint(v687)%32)&v689 {
																								v723 = v690
																								if v723 != 0 {
																									v726 = int32(1)
																								} else {
																									v726 = int32(-1)
																								}
																								v732 = v726
																							} else {
																								if v640 == int32(5) {
																									v722 = v637
																									v732 = v722
																								} else {
																									v698 = int32(5)
																									v700 = int32(128)
																									v701 = v644 << (uint(v698) % 32) & v700
																									if v701 != v648<<(uint(v698)%32)&v700 {
																										v723 = v701
																										if v723 != 0 {
																											v726 = int32(1)
																										} else {
																											v726 = int32(-1)
																										}
																										v732 = v726
																									} else {
																										if v640 < int32(7) {
																											v722 = v637
																											v732 = v722
																										} else {
																											v709 = int32(6)
																											v711 = int32(128)
																											v712 = v644 << (uint(v709) % 32) & v711
																											if v712 != v648<<(uint(v709)%32)&v711 {
																												v723 = v712
																												if v723 != 0 {
																													v726 = int32(1)
																												} else {
																													v726 = int32(-1)
																												}
																												v732 = v726
																											} else {
																												v722 = v637
																												v732 = v722
																											}
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
													switch v18&int32(_a_F_inet_gist_consistent_0) - int32(18) {
													case 0:
														v782 = v732
														return base.B2i32(v782 == int32(0))
													case 1:
														v775 = v732
														return base.B2i32(v775 != int32(0))
													case 2:
														return int32(base.Ui32(v732) >> (uint(int32(31)) % 32))
													case 3:
														return base.B2i32(v732 <= int32(0))
													case 4:
														return base.B2i32(int32(0) < v732)
													case 5:
														return int32(base.Ui32(v732^int32(-1)) >> (uint(int32(31)) % 32))
													default:
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v754 = m.ExcPending
														if v754 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_1), int32(0))
															mBase = m.M
															v758 = m.ExcPending
															if v758 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_inet_gist_consistent_2), int32(327), int32(_a_F_inet_gist_consistent_3))
																mBase = m.M
																v763 = m.ExcPending
																if v763 != 0 {
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
													return int32(0)
												}
											}
										} else {
											if base.Ui32(v481) < base.Ui32(v480) {
												v768 = v475
												return v768
											} else {
												if base.Ui32(v481) <= base.Ui32(v480) {
													v618 = v478
													if v618 != 0 {
														v622 = v44
													} else {
														v622 = v46
													}
													v624 = v622 + int32(2)
													v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
													if v627 == int32(3) {
														v630 = int32(128)
													} else {
														v630 = int32(32)
													}
													v635 = base.I32_div_s(v630, int32(8))
													v636 = F_memcmp(m, v85, v624, v635)
													mBase = m.M
													if v636 != 0 {
														v722 = v636
														v732 = v722
													} else {
														v637 = int32(0)
														v640 = v630 - v635<<(uint(int32(3))%32)
														if v640 <= v637 {
															v722 = v637
															v732 = v722
														} else {
															v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v635))))
															v645 = int32(128)
															v646 = v644 & v645
															v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624+v635))))
															if v646 != v648&v645 {
																v723 = v646
																if v723 != 0 {
																	v726 = int32(1)
																} else {
																	v726 = int32(-1)
																}
																v732 = v726
															} else {
																if v640 == int32(1) {
																	v722 = v637
																	v732 = v722
																} else {
																	v654 = int32(1)
																	v656 = int32(128)
																	v657 = v644 << (uint(v654) % 32) & v656
																	if v657 != v648<<(uint(v654)%32)&v656 {
																		v723 = v657
																		if v723 != 0 {
																			v726 = int32(1)
																		} else {
																			v726 = int32(-1)
																		}
																		v732 = v726
																	} else {
																		if v640 < int32(3) {
																			v722 = v637
																			v732 = v722
																		} else {
																			v665 = int32(2)
																			v667 = int32(128)
																			v668 = v644 << (uint(v665) % 32) & v667
																			if v668 != v648<<(uint(v665)%32)&v667 {
																				v723 = v668
																				if v723 != 0 {
																					v726 = int32(1)
																				} else {
																					v726 = int32(-1)
																				}
																				v732 = v726
																			} else {
																				if v640 == int32(3) {
																					v722 = v637
																					v732 = v722
																				} else {
																					v676 = int32(3)
																					v678 = int32(128)
																					v679 = v644 << (uint(v676) % 32) & v678
																					if v679 != v648<<(uint(v676)%32)&v678 {
																						v723 = v679
																						if v723 != 0 {
																							v726 = int32(1)
																						} else {
																							v726 = int32(-1)
																						}
																						v732 = v726
																					} else {
																						if v640 < int32(5) {
																							v722 = v637
																							v732 = v722
																						} else {
																							v687 = int32(4)
																							v689 = int32(128)
																							v690 = v644 << (uint(v687) % 32) & v689
																							if v690 != v648<<(uint(v687)%32)&v689 {
																								v723 = v690
																								if v723 != 0 {
																									v726 = int32(1)
																								} else {
																									v726 = int32(-1)
																								}
																								v732 = v726
																							} else {
																								if v640 == int32(5) {
																									v722 = v637
																									v732 = v722
																								} else {
																									v698 = int32(5)
																									v700 = int32(128)
																									v701 = v644 << (uint(v698) % 32) & v700
																									if v701 != v648<<(uint(v698)%32)&v700 {
																										v723 = v701
																										if v723 != 0 {
																											v726 = int32(1)
																										} else {
																											v726 = int32(-1)
																										}
																										v732 = v726
																									} else {
																										if v640 < int32(7) {
																											v722 = v637
																											v732 = v722
																										} else {
																											v709 = int32(6)
																											v711 = int32(128)
																											v712 = v644 << (uint(v709) % 32) & v711
																											if v712 != v648<<(uint(v709)%32)&v711 {
																												v723 = v712
																												if v723 != 0 {
																													v726 = int32(1)
																												} else {
																													v726 = int32(-1)
																												}
																												v732 = v726
																											} else {
																												v722 = v637
																												v732 = v722
																											}
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
													switch v18&int32(_a_F_inet_gist_consistent_0) - int32(18) {
													case 0:
														v782 = v732
														return base.B2i32(v782 == int32(0))
													case 1:
														v775 = v732
														return base.B2i32(v775 != int32(0))
													case 2:
														return int32(base.Ui32(v732) >> (uint(int32(31)) % 32))
													case 3:
														return base.B2i32(v732 <= int32(0))
													case 4:
														return base.B2i32(int32(0) < v732)
													case 5:
														return int32(base.Ui32(v732^int32(-1)) >> (uint(int32(31)) % 32))
													default:
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v754 = m.ExcPending
														if v754 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_1), int32(0))
															mBase = m.M
															v758 = m.ExcPending
															if v758 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_inet_gist_consistent_2), int32(327), int32(_a_F_inet_gist_consistent_3))
																mBase = m.M
																v763 = m.ExcPending
																if v763 != 0 {
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
													return int32(0)
												}
											}
										}
									}
								}
							}
						}
					} else {
						return int32(0)
					}
				default:
					v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
					v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
					v82 = v81
					v83 = v80
					v85 = v19 + int32(4)
					v87 = v47 + int32(2)
					v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+3)))
					if base.Ui32(v88) < base.Ui32(v82) {
						v90 = v88
					} else {
						v90 = v82
					}
					if base.Ui32(v90) < base.Ui32(v83) {
						v92 = v90
					} else {
						v92 = v83
					}
					v97 = base.I32_div_s(v92, int32(8))
					v98 = F_memcmp(m, v85, v87, v97)
					mBase = m.M
					if v98 != 0 {
						v184 = v98
						v194 = v184
					} else {
						v99 = int32(0)
						v102 = v92 - v97<<(uint(int32(3))%32)
						if v102 <= v99 {
							v184 = v99
							v194 = v184
						} else {
							v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v97))))
							v107 = int32(128)
							v108 = v106 & v107
							v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+v97))))
							if v108 != v110&v107 {
								v185 = v108
								if v185 != 0 {
									v188 = int32(1)
								} else {
									v188 = int32(-1)
								}
								v194 = v188
							} else {
								if v102 == int32(1) {
									v184 = v99
									v194 = v184
								} else {
									v116 = int32(1)
									v118 = int32(128)
									v119 = v106 << (uint(v116) % 32) & v118
									if v119 != v110<<(uint(v116)%32)&v118 {
										v185 = v119
										if v185 != 0 {
											v188 = int32(1)
										} else {
											v188 = int32(-1)
										}
										v194 = v188
									} else {
										if v102 < int32(3) {
											v184 = v99
											v194 = v184
										} else {
											v127 = int32(2)
											v129 = int32(128)
											v130 = v106 << (uint(v127) % 32) & v129
											if v130 != v110<<(uint(v127)%32)&v129 {
												v185 = v130
												if v185 != 0 {
													v188 = int32(1)
												} else {
													v188 = int32(-1)
												}
												v194 = v188
											} else {
												if v102 == int32(3) {
													v184 = v99
													v194 = v184
												} else {
													v138 = int32(3)
													v140 = int32(128)
													v141 = v106 << (uint(v138) % 32) & v140
													if v141 != v110<<(uint(v138)%32)&v140 {
														v185 = v141
														if v185 != 0 {
															v188 = int32(1)
														} else {
															v188 = int32(-1)
														}
														v194 = v188
													} else {
														if v102 < int32(5) {
															v184 = v99
															v194 = v184
														} else {
															v149 = int32(4)
															v151 = int32(128)
															v152 = v106 << (uint(v149) % 32) & v151
															if v152 != v110<<(uint(v149)%32)&v151 {
																v185 = v152
																if v185 != 0 {
																	v188 = int32(1)
																} else {
																	v188 = int32(-1)
																}
																v194 = v188
															} else {
																if v102 == int32(5) {
																	v184 = v99
																	v194 = v184
																} else {
																	v160 = int32(5)
																	v162 = int32(128)
																	v163 = v106 << (uint(v160) % 32) & v162
																	if v163 != v110<<(uint(v160)%32)&v162 {
																		v185 = v163
																		if v185 != 0 {
																			v188 = int32(1)
																		} else {
																			v188 = int32(-1)
																		}
																		v194 = v188
																	} else {
																		if v102 < int32(7) {
																			v184 = v99
																			v194 = v184
																		} else {
																			v171 = int32(6)
																			v173 = int32(128)
																			v174 = v106 << (uint(v171) % 32) & v173
																			if v174 != v110<<(uint(v171)%32)&v173 {
																				v185 = v174
																				if v185 != 0 {
																					v188 = int32(1)
																				} else {
																					v188 = int32(-1)
																				}
																				v194 = v188
																			} else {
																				v184 = v99
																				v194 = v184
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
					switch v18&int32(_a_F_inet_gist_consistent_0) - int32(3) {
					case 0, 21, 22, 23, 24:
						v313 = v194
						return base.B2i32(v313 == int32(0))
					default:
						v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
						v618 = v615 & int32(1)
						if v618 != 0 {
							v622 = v44
						} else {
							v622 = v46
						}
						v624 = v622 + int32(2)
						v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
						if v627 == int32(3) {
							v630 = int32(128)
						} else {
							v630 = int32(32)
						}
						v635 = base.I32_div_s(v630, int32(8))
						v636 = F_memcmp(m, v85, v624, v635)
						mBase = m.M
						if v636 != 0 {
							v722 = v636
							v732 = v722
						} else {
							v637 = int32(0)
							v640 = v630 - v635<<(uint(int32(3))%32)
							if v640 <= v637 {
								v722 = v637
								v732 = v722
							} else {
								v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v635))))
								v645 = int32(128)
								v646 = v644 & v645
								v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624+v635))))
								if v646 != v648&v645 {
									v723 = v646
									if v723 != 0 {
										v726 = int32(1)
									} else {
										v726 = int32(-1)
									}
									v732 = v726
								} else {
									if v640 == int32(1) {
										v722 = v637
										v732 = v722
									} else {
										v654 = int32(1)
										v656 = int32(128)
										v657 = v644 << (uint(v654) % 32) & v656
										if v657 != v648<<(uint(v654)%32)&v656 {
											v723 = v657
											if v723 != 0 {
												v726 = int32(1)
											} else {
												v726 = int32(-1)
											}
											v732 = v726
										} else {
											if v640 < int32(3) {
												v722 = v637
												v732 = v722
											} else {
												v665 = int32(2)
												v667 = int32(128)
												v668 = v644 << (uint(v665) % 32) & v667
												if v668 != v648<<(uint(v665)%32)&v667 {
													v723 = v668
													if v723 != 0 {
														v726 = int32(1)
													} else {
														v726 = int32(-1)
													}
													v732 = v726
												} else {
													if v640 == int32(3) {
														v722 = v637
														v732 = v722
													} else {
														v676 = int32(3)
														v678 = int32(128)
														v679 = v644 << (uint(v676) % 32) & v678
														if v679 != v648<<(uint(v676)%32)&v678 {
															v723 = v679
															if v723 != 0 {
																v726 = int32(1)
															} else {
																v726 = int32(-1)
															}
															v732 = v726
														} else {
															if v640 < int32(5) {
																v722 = v637
																v732 = v722
															} else {
																v687 = int32(4)
																v689 = int32(128)
																v690 = v644 << (uint(v687) % 32) & v689
																if v690 != v648<<(uint(v687)%32)&v689 {
																	v723 = v690
																	if v723 != 0 {
																		v726 = int32(1)
																	} else {
																		v726 = int32(-1)
																	}
																	v732 = v726
																} else {
																	if v640 == int32(5) {
																		v722 = v637
																		v732 = v722
																	} else {
																		v698 = int32(5)
																		v700 = int32(128)
																		v701 = v644 << (uint(v698) % 32) & v700
																		if v701 != v648<<(uint(v698)%32)&v700 {
																			v723 = v701
																			if v723 != 0 {
																				v726 = int32(1)
																			} else {
																				v726 = int32(-1)
																			}
																			v732 = v726
																		} else {
																			if v640 < int32(7) {
																				v722 = v637
																				v732 = v722
																			} else {
																				v709 = int32(6)
																				v711 = int32(128)
																				v712 = v644 << (uint(v709) % 32) & v711
																				if v712 != v648<<(uint(v709)%32)&v711 {
																					v723 = v712
																					if v723 != 0 {
																						v726 = int32(1)
																					} else {
																						v726 = int32(-1)
																					}
																					v732 = v726
																				} else {
																					v722 = v637
																					v732 = v722
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
						switch v18&int32(_a_F_inet_gist_consistent_0) - int32(18) {
						case 0:
							v782 = v732
							return base.B2i32(v782 == int32(0))
						case 1:
							v775 = v732
							return base.B2i32(v775 != int32(0))
						case 2:
							return int32(base.Ui32(v732) >> (uint(int32(31)) % 32))
						case 3:
							return base.B2i32(v732 <= int32(0))
						case 4:
							return base.B2i32(int32(0) < v732)
						case 5:
							return int32(base.Ui32(v732^int32(-1)) >> (uint(int32(31)) % 32))
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v754 = m.ExcPending
							if v754 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_1), int32(0))
								mBase = m.M
								v758 = m.ExcPending
								if v758 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_inet_gist_consistent_2), int32(327), int32(_a_F_inet_gist_consistent_3))
									mBase = m.M
									v763 = m.ExcPending
									if v763 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					case 15:
						if v194 != 0 {
							v768 = int32(0)
							return v768
						} else {
							v335 = int32(1)
							v336 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							v337 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v336)+16)))
							v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336+v337)+12)))
							if v339&v335 == int32(0) {
								v768 = v335
								return v768
							} else {
								v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
								v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
								if v346&int32(1) != 0 {
									v349 = v44
								} else {
									v349 = v46
								}
								v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+1)))
								if v345 != v350 {
									v768 = int32(0)
									return v768
								} else {
									v353 = v349 + int32(2)
									v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
									if v356 == int32(3) {
										v359 = int32(128)
									} else {
										v359 = int32(32)
									}
									v364 = base.I32_div_s(v359, int32(8))
									v365 = F_memcmp(m, v85, v353, v364)
									mBase = m.M
									if v365 != 0 {
										v451 = v365
										v461 = v451
									} else {
										v366 = int32(0)
										v369 = v359 - v364<<(uint(int32(3))%32)
										if v369 <= v366 {
											v451 = v366
											v461 = v451
										} else {
											v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v364))))
											v374 = int32(128)
											v375 = v373 & v374
											v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353+v364))))
											if v375 != v377&v374 {
												v452 = v375
												if v452 != 0 {
													v455 = int32(1)
												} else {
													v455 = int32(-1)
												}
												v461 = v455
											} else {
												if v369 == int32(1) {
													v451 = v366
													v461 = v451
												} else {
													v383 = int32(1)
													v385 = int32(128)
													v386 = v373 << (uint(v383) % 32) & v385
													if v386 != v377<<(uint(v383)%32)&v385 {
														v452 = v386
														if v452 != 0 {
															v455 = int32(1)
														} else {
															v455 = int32(-1)
														}
														v461 = v455
													} else {
														if v369 < int32(3) {
															v451 = v366
															v461 = v451
														} else {
															v394 = int32(2)
															v396 = int32(128)
															v397 = v373 << (uint(v394) % 32) & v396
															if v397 != v377<<(uint(v394)%32)&v396 {
																v452 = v397
																if v452 != 0 {
																	v455 = int32(1)
																} else {
																	v455 = int32(-1)
																}
																v461 = v455
															} else {
																if v369 == int32(3) {
																	v451 = v366
																	v461 = v451
																} else {
																	v405 = int32(3)
																	v407 = int32(128)
																	v408 = v373 << (uint(v405) % 32) & v407
																	if v408 != v377<<(uint(v405)%32)&v407 {
																		v452 = v408
																		if v452 != 0 {
																			v455 = int32(1)
																		} else {
																			v455 = int32(-1)
																		}
																		v461 = v455
																	} else {
																		if v369 < int32(5) {
																			v451 = v366
																			v461 = v451
																		} else {
																			v416 = int32(4)
																			v418 = int32(128)
																			v419 = v373 << (uint(v416) % 32) & v418
																			if v419 != v377<<(uint(v416)%32)&v418 {
																				v452 = v419
																				if v452 != 0 {
																					v455 = int32(1)
																				} else {
																					v455 = int32(-1)
																				}
																				v461 = v455
																			} else {
																				if v369 == int32(5) {
																					v451 = v366
																					v461 = v451
																				} else {
																					v427 = int32(5)
																					v429 = int32(128)
																					v430 = v373 << (uint(v427) % 32) & v429
																					if v430 != v377<<(uint(v427)%32)&v429 {
																						v452 = v430
																						if v452 != 0 {
																							v455 = int32(1)
																						} else {
																							v455 = int32(-1)
																						}
																						v461 = v455
																					} else {
																						if v369 < int32(7) {
																							v451 = v366
																							v461 = v451
																						} else {
																							v438 = int32(6)
																							v440 = int32(128)
																							v441 = v373 << (uint(v438) % 32) & v440
																							if v441 != v377<<(uint(v438)%32)&v440 {
																								v452 = v441
																								if v452 != 0 {
																									v455 = int32(1)
																								} else {
																									v455 = int32(-1)
																								}
																								v461 = v455
																							} else {
																								v451 = v366
																								v461 = v451
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
									v782 = v461
									return base.B2i32(v782 == int32(0))
								}
							}
						}
					case 16:
						if v194 != 0 {
							v768 = v28
							return v768
						} else {
							v488 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							v489 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v488)+16)))
							v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488+v489)+12)))
							if v491&int32(1) == int32(0) {
								v768 = v28
								return v768
							} else {
								v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
								v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
								if v497&int32(1) != 0 {
									v500 = v44
								} else {
									v500 = v46
								}
								v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+1)))
								if v496 != v501 {
									v768 = v28
									return v768
								} else {
									v504 = v500 + int32(2)
									v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
									if v507 == int32(3) {
										v510 = int32(128)
									} else {
										v510 = int32(32)
									}
									v515 = base.I32_div_s(v510, int32(8))
									v516 = F_memcmp(m, v85, v504, v515)
									mBase = m.M
									if v516 != 0 {
										v602 = v516
										v612 = v602
									} else {
										v517 = int32(0)
										v520 = v510 - v515<<(uint(int32(3))%32)
										if v520 <= v517 {
											v602 = v517
											v612 = v602
										} else {
											v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v515))))
											v525 = int32(128)
											v526 = v524 & v525
											v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504+v515))))
											if v526 != v528&v525 {
												v603 = v526
												if v603 != 0 {
													v606 = int32(1)
												} else {
													v606 = int32(-1)
												}
												v612 = v606
											} else {
												if v520 == int32(1) {
													v602 = v517
													v612 = v602
												} else {
													v534 = int32(1)
													v536 = int32(128)
													v537 = v524 << (uint(v534) % 32) & v536
													if v537 != v528<<(uint(v534)%32)&v536 {
														v603 = v537
														if v603 != 0 {
															v606 = int32(1)
														} else {
															v606 = int32(-1)
														}
														v612 = v606
													} else {
														if v520 < int32(3) {
															v602 = v517
															v612 = v602
														} else {
															v545 = int32(2)
															v547 = int32(128)
															v548 = v524 << (uint(v545) % 32) & v547
															if v548 != v528<<(uint(v545)%32)&v547 {
																v603 = v548
																if v603 != 0 {
																	v606 = int32(1)
																} else {
																	v606 = int32(-1)
																}
																v612 = v606
															} else {
																if v520 == int32(3) {
																	v602 = v517
																	v612 = v602
																} else {
																	v556 = int32(3)
																	v558 = int32(128)
																	v559 = v524 << (uint(v556) % 32) & v558
																	if v559 != v528<<(uint(v556)%32)&v558 {
																		v603 = v559
																		if v603 != 0 {
																			v606 = int32(1)
																		} else {
																			v606 = int32(-1)
																		}
																		v612 = v606
																	} else {
																		if v520 < int32(5) {
																			v602 = v517
																			v612 = v602
																		} else {
																			v567 = int32(4)
																			v569 = int32(128)
																			v570 = v524 << (uint(v567) % 32) & v569
																			if v570 != v528<<(uint(v567)%32)&v569 {
																				v603 = v570
																				if v603 != 0 {
																					v606 = int32(1)
																				} else {
																					v606 = int32(-1)
																				}
																				v612 = v606
																			} else {
																				if v520 == int32(5) {
																					v602 = v517
																					v612 = v602
																				} else {
																					v578 = int32(5)
																					v580 = int32(128)
																					v581 = v524 << (uint(v578) % 32) & v580
																					if v581 != v528<<(uint(v578)%32)&v580 {
																						v603 = v581
																						if v603 != 0 {
																							v606 = int32(1)
																						} else {
																							v606 = int32(-1)
																						}
																						v612 = v606
																					} else {
																						if v520 < int32(7) {
																							v602 = v517
																							v612 = v602
																						} else {
																							v589 = int32(6)
																							v591 = int32(128)
																							v592 = v524 << (uint(v589) % 32) & v591
																							if v592 != v528<<(uint(v589)%32)&v591 {
																								v603 = v592
																								if v603 != 0 {
																									v606 = int32(1)
																								} else {
																									v606 = int32(-1)
																								}
																								v612 = v606
																							} else {
																								v602 = v517
																								v612 = v602
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
									v775 = v612
									return base.B2i32(v775 != int32(0))
								}
							}
						}
					case 17, 18:
						v322 = int32(0)
						if v322 < v194 {
							v768 = v322
							return v768
						} else {
							v325 = int32(1)
							if v194 < int32(0) {
								v768 = v325
								return v768
							} else {
								v328 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								v329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v328)+16)))
								v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328+v329)+12)))
								if v331&int32(1) != 0 {
									v475 = v325
									v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
									v478 = v476 & int32(1)
									if v478 != 0 {
										v479 = v44
									} else {
										v479 = v46
									}
									v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479)+1)))
									v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
									if v18&int32(_a_F_inet_gist_consistent_4) != int32(20) {
										if base.Ui32(v480) < base.Ui32(v481) {
											v768 = v475
											return v768
										} else {
											if base.Ui32(v480) <= base.Ui32(v481) {
												v618 = v478
												if v618 != 0 {
													v622 = v44
												} else {
													v622 = v46
												}
												v624 = v622 + int32(2)
												v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
												if v627 == int32(3) {
													v630 = int32(128)
												} else {
													v630 = int32(32)
												}
												v635 = base.I32_div_s(v630, int32(8))
												v636 = F_memcmp(m, v85, v624, v635)
												mBase = m.M
												if v636 != 0 {
													v722 = v636
													v732 = v722
												} else {
													v637 = int32(0)
													v640 = v630 - v635<<(uint(int32(3))%32)
													if v640 <= v637 {
														v722 = v637
														v732 = v722
													} else {
														v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v635))))
														v645 = int32(128)
														v646 = v644 & v645
														v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624+v635))))
														if v646 != v648&v645 {
															v723 = v646
															if v723 != 0 {
																v726 = int32(1)
															} else {
																v726 = int32(-1)
															}
															v732 = v726
														} else {
															if v640 == int32(1) {
																v722 = v637
																v732 = v722
															} else {
																v654 = int32(1)
																v656 = int32(128)
																v657 = v644 << (uint(v654) % 32) & v656
																if v657 != v648<<(uint(v654)%32)&v656 {
																	v723 = v657
																	if v723 != 0 {
																		v726 = int32(1)
																	} else {
																		v726 = int32(-1)
																	}
																	v732 = v726
																} else {
																	if v640 < int32(3) {
																		v722 = v637
																		v732 = v722
																	} else {
																		v665 = int32(2)
																		v667 = int32(128)
																		v668 = v644 << (uint(v665) % 32) & v667
																		if v668 != v648<<(uint(v665)%32)&v667 {
																			v723 = v668
																			if v723 != 0 {
																				v726 = int32(1)
																			} else {
																				v726 = int32(-1)
																			}
																			v732 = v726
																		} else {
																			if v640 == int32(3) {
																				v722 = v637
																				v732 = v722
																			} else {
																				v676 = int32(3)
																				v678 = int32(128)
																				v679 = v644 << (uint(v676) % 32) & v678
																				if v679 != v648<<(uint(v676)%32)&v678 {
																					v723 = v679
																					if v723 != 0 {
																						v726 = int32(1)
																					} else {
																						v726 = int32(-1)
																					}
																					v732 = v726
																				} else {
																					if v640 < int32(5) {
																						v722 = v637
																						v732 = v722
																					} else {
																						v687 = int32(4)
																						v689 = int32(128)
																						v690 = v644 << (uint(v687) % 32) & v689
																						if v690 != v648<<(uint(v687)%32)&v689 {
																							v723 = v690
																							if v723 != 0 {
																								v726 = int32(1)
																							} else {
																								v726 = int32(-1)
																							}
																							v732 = v726
																						} else {
																							if v640 == int32(5) {
																								v722 = v637
																								v732 = v722
																							} else {
																								v698 = int32(5)
																								v700 = int32(128)
																								v701 = v644 << (uint(v698) % 32) & v700
																								if v701 != v648<<(uint(v698)%32)&v700 {
																									v723 = v701
																									if v723 != 0 {
																										v726 = int32(1)
																									} else {
																										v726 = int32(-1)
																									}
																									v732 = v726
																								} else {
																									if v640 < int32(7) {
																										v722 = v637
																										v732 = v722
																									} else {
																										v709 = int32(6)
																										v711 = int32(128)
																										v712 = v644 << (uint(v709) % 32) & v711
																										if v712 != v648<<(uint(v709)%32)&v711 {
																											v723 = v712
																											if v723 != 0 {
																												v726 = int32(1)
																											} else {
																												v726 = int32(-1)
																											}
																											v732 = v726
																										} else {
																											v722 = v637
																											v732 = v722
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
												switch v18&int32(_a_F_inet_gist_consistent_0) - int32(18) {
												case 0:
													v782 = v732
													return base.B2i32(v782 == int32(0))
												case 1:
													v775 = v732
													return base.B2i32(v775 != int32(0))
												case 2:
													return int32(base.Ui32(v732) >> (uint(int32(31)) % 32))
												case 3:
													return base.B2i32(v732 <= int32(0))
												case 4:
													return base.B2i32(int32(0) < v732)
												case 5:
													return int32(base.Ui32(v732^int32(-1)) >> (uint(int32(31)) % 32))
												default:
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v754 = m.ExcPending
													if v754 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_1), int32(0))
														mBase = m.M
														v758 = m.ExcPending
														if v758 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_inet_gist_consistent_2), int32(327), int32(_a_F_inet_gist_consistent_3))
															mBase = m.M
															v763 = m.ExcPending
															if v763 != 0 {
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
												return int32(0)
											}
										}
									} else {
										if base.Ui32(v481) < base.Ui32(v480) {
											v768 = v475
											return v768
										} else {
											if base.Ui32(v481) <= base.Ui32(v480) {
												v618 = v478
												if v618 != 0 {
													v622 = v44
												} else {
													v622 = v46
												}
												v624 = v622 + int32(2)
												v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
												if v627 == int32(3) {
													v630 = int32(128)
												} else {
													v630 = int32(32)
												}
												v635 = base.I32_div_s(v630, int32(8))
												v636 = F_memcmp(m, v85, v624, v635)
												mBase = m.M
												if v636 != 0 {
													v722 = v636
													v732 = v722
												} else {
													v637 = int32(0)
													v640 = v630 - v635<<(uint(int32(3))%32)
													if v640 <= v637 {
														v722 = v637
														v732 = v722
													} else {
														v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v635))))
														v645 = int32(128)
														v646 = v644 & v645
														v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624+v635))))
														if v646 != v648&v645 {
															v723 = v646
															if v723 != 0 {
																v726 = int32(1)
															} else {
																v726 = int32(-1)
															}
															v732 = v726
														} else {
															if v640 == int32(1) {
																v722 = v637
																v732 = v722
															} else {
																v654 = int32(1)
																v656 = int32(128)
																v657 = v644 << (uint(v654) % 32) & v656
																if v657 != v648<<(uint(v654)%32)&v656 {
																	v723 = v657
																	if v723 != 0 {
																		v726 = int32(1)
																	} else {
																		v726 = int32(-1)
																	}
																	v732 = v726
																} else {
																	if v640 < int32(3) {
																		v722 = v637
																		v732 = v722
																	} else {
																		v665 = int32(2)
																		v667 = int32(128)
																		v668 = v644 << (uint(v665) % 32) & v667
																		if v668 != v648<<(uint(v665)%32)&v667 {
																			v723 = v668
																			if v723 != 0 {
																				v726 = int32(1)
																			} else {
																				v726 = int32(-1)
																			}
																			v732 = v726
																		} else {
																			if v640 == int32(3) {
																				v722 = v637
																				v732 = v722
																			} else {
																				v676 = int32(3)
																				v678 = int32(128)
																				v679 = v644 << (uint(v676) % 32) & v678
																				if v679 != v648<<(uint(v676)%32)&v678 {
																					v723 = v679
																					if v723 != 0 {
																						v726 = int32(1)
																					} else {
																						v726 = int32(-1)
																					}
																					v732 = v726
																				} else {
																					if v640 < int32(5) {
																						v722 = v637
																						v732 = v722
																					} else {
																						v687 = int32(4)
																						v689 = int32(128)
																						v690 = v644 << (uint(v687) % 32) & v689
																						if v690 != v648<<(uint(v687)%32)&v689 {
																							v723 = v690
																							if v723 != 0 {
																								v726 = int32(1)
																							} else {
																								v726 = int32(-1)
																							}
																							v732 = v726
																						} else {
																							if v640 == int32(5) {
																								v722 = v637
																								v732 = v722
																							} else {
																								v698 = int32(5)
																								v700 = int32(128)
																								v701 = v644 << (uint(v698) % 32) & v700
																								if v701 != v648<<(uint(v698)%32)&v700 {
																									v723 = v701
																									if v723 != 0 {
																										v726 = int32(1)
																									} else {
																										v726 = int32(-1)
																									}
																									v732 = v726
																								} else {
																									if v640 < int32(7) {
																										v722 = v637
																										v732 = v722
																									} else {
																										v709 = int32(6)
																										v711 = int32(128)
																										v712 = v644 << (uint(v709) % 32) & v711
																										if v712 != v648<<(uint(v709)%32)&v711 {
																											v723 = v712
																											if v723 != 0 {
																												v726 = int32(1)
																											} else {
																												v726 = int32(-1)
																											}
																											v732 = v726
																										} else {
																											v722 = v637
																											v732 = v722
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
												switch v18&int32(_a_F_inet_gist_consistent_0) - int32(18) {
												case 0:
													v782 = v732
													return base.B2i32(v782 == int32(0))
												case 1:
													v775 = v732
													return base.B2i32(v775 != int32(0))
												case 2:
													return int32(base.Ui32(v732) >> (uint(int32(31)) % 32))
												case 3:
													return base.B2i32(v732 <= int32(0))
												case 4:
													return base.B2i32(int32(0) < v732)
												case 5:
													return int32(base.Ui32(v732^int32(-1)) >> (uint(int32(31)) % 32))
												default:
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v754 = m.ExcPending
													if v754 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_1), int32(0))
														mBase = m.M
														v758 = m.ExcPending
														if v758 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_inet_gist_consistent_2), int32(327), int32(_a_F_inet_gist_consistent_3))
															mBase = m.M
															v763 = m.ExcPending
															if v763 != 0 {
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
												return int32(0)
											}
										}
									}
								} else {
									v768 = v325
									return v768
								}
							}
						}
					case 19, 20:
						v462 = int32(0)
						if v194 < v462 {
							v768 = v462
							return v768
						} else {
							v465 = int32(1)
							if v194 != 0 {
								v768 = v465
								return v768
							} else {
								v466 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								v467 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v466)+16)))
								v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v466+v467)+12)))
								if v469&int32(1) == int32(0) {
									v768 = v465
									return v768
								} else {
									v475 = v465
									v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
									v478 = v476 & int32(1)
									if v478 != 0 {
										v479 = v44
									} else {
										v479 = v46
									}
									v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479)+1)))
									v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
									if v18&int32(_a_F_inet_gist_consistent_4) != int32(20) {
										if base.Ui32(v480) < base.Ui32(v481) {
											v768 = v475
											return v768
										} else {
											if base.Ui32(v480) <= base.Ui32(v481) {
												v618 = v478
												if v618 != 0 {
													v622 = v44
												} else {
													v622 = v46
												}
												v624 = v622 + int32(2)
												v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
												if v627 == int32(3) {
													v630 = int32(128)
												} else {
													v630 = int32(32)
												}
												v635 = base.I32_div_s(v630, int32(8))
												v636 = F_memcmp(m, v85, v624, v635)
												mBase = m.M
												if v636 != 0 {
													v722 = v636
													v732 = v722
												} else {
													v637 = int32(0)
													v640 = v630 - v635<<(uint(int32(3))%32)
													if v640 <= v637 {
														v722 = v637
														v732 = v722
													} else {
														v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v635))))
														v645 = int32(128)
														v646 = v644 & v645
														v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624+v635))))
														if v646 != v648&v645 {
															v723 = v646
															if v723 != 0 {
																v726 = int32(1)
															} else {
																v726 = int32(-1)
															}
															v732 = v726
														} else {
															if v640 == int32(1) {
																v722 = v637
																v732 = v722
															} else {
																v654 = int32(1)
																v656 = int32(128)
																v657 = v644 << (uint(v654) % 32) & v656
																if v657 != v648<<(uint(v654)%32)&v656 {
																	v723 = v657
																	if v723 != 0 {
																		v726 = int32(1)
																	} else {
																		v726 = int32(-1)
																	}
																	v732 = v726
																} else {
																	if v640 < int32(3) {
																		v722 = v637
																		v732 = v722
																	} else {
																		v665 = int32(2)
																		v667 = int32(128)
																		v668 = v644 << (uint(v665) % 32) & v667
																		if v668 != v648<<(uint(v665)%32)&v667 {
																			v723 = v668
																			if v723 != 0 {
																				v726 = int32(1)
																			} else {
																				v726 = int32(-1)
																			}
																			v732 = v726
																		} else {
																			if v640 == int32(3) {
																				v722 = v637
																				v732 = v722
																			} else {
																				v676 = int32(3)
																				v678 = int32(128)
																				v679 = v644 << (uint(v676) % 32) & v678
																				if v679 != v648<<(uint(v676)%32)&v678 {
																					v723 = v679
																					if v723 != 0 {
																						v726 = int32(1)
																					} else {
																						v726 = int32(-1)
																					}
																					v732 = v726
																				} else {
																					if v640 < int32(5) {
																						v722 = v637
																						v732 = v722
																					} else {
																						v687 = int32(4)
																						v689 = int32(128)
																						v690 = v644 << (uint(v687) % 32) & v689
																						if v690 != v648<<(uint(v687)%32)&v689 {
																							v723 = v690
																							if v723 != 0 {
																								v726 = int32(1)
																							} else {
																								v726 = int32(-1)
																							}
																							v732 = v726
																						} else {
																							if v640 == int32(5) {
																								v722 = v637
																								v732 = v722
																							} else {
																								v698 = int32(5)
																								v700 = int32(128)
																								v701 = v644 << (uint(v698) % 32) & v700
																								if v701 != v648<<(uint(v698)%32)&v700 {
																									v723 = v701
																									if v723 != 0 {
																										v726 = int32(1)
																									} else {
																										v726 = int32(-1)
																									}
																									v732 = v726
																								} else {
																									if v640 < int32(7) {
																										v722 = v637
																										v732 = v722
																									} else {
																										v709 = int32(6)
																										v711 = int32(128)
																										v712 = v644 << (uint(v709) % 32) & v711
																										if v712 != v648<<(uint(v709)%32)&v711 {
																											v723 = v712
																											if v723 != 0 {
																												v726 = int32(1)
																											} else {
																												v726 = int32(-1)
																											}
																											v732 = v726
																										} else {
																											v722 = v637
																											v732 = v722
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
												switch v18&int32(_a_F_inet_gist_consistent_0) - int32(18) {
												case 0:
													v782 = v732
													return base.B2i32(v782 == int32(0))
												case 1:
													v775 = v732
													return base.B2i32(v775 != int32(0))
												case 2:
													return int32(base.Ui32(v732) >> (uint(int32(31)) % 32))
												case 3:
													return base.B2i32(v732 <= int32(0))
												case 4:
													return base.B2i32(int32(0) < v732)
												case 5:
													return int32(base.Ui32(v732^int32(-1)) >> (uint(int32(31)) % 32))
												default:
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v754 = m.ExcPending
													if v754 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_1), int32(0))
														mBase = m.M
														v758 = m.ExcPending
														if v758 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_inet_gist_consistent_2), int32(327), int32(_a_F_inet_gist_consistent_3))
															mBase = m.M
															v763 = m.ExcPending
															if v763 != 0 {
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
												return int32(0)
											}
										}
									} else {
										if base.Ui32(v481) < base.Ui32(v480) {
											v768 = v475
											return v768
										} else {
											if base.Ui32(v481) <= base.Ui32(v480) {
												v618 = v478
												if v618 != 0 {
													v622 = v44
												} else {
													v622 = v46
												}
												v624 = v622 + int32(2)
												v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
												if v627 == int32(3) {
													v630 = int32(128)
												} else {
													v630 = int32(32)
												}
												v635 = base.I32_div_s(v630, int32(8))
												v636 = F_memcmp(m, v85, v624, v635)
												mBase = m.M
												if v636 != 0 {
													v722 = v636
													v732 = v722
												} else {
													v637 = int32(0)
													v640 = v630 - v635<<(uint(int32(3))%32)
													if v640 <= v637 {
														v722 = v637
														v732 = v722
													} else {
														v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v635))))
														v645 = int32(128)
														v646 = v644 & v645
														v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624+v635))))
														if v646 != v648&v645 {
															v723 = v646
															if v723 != 0 {
																v726 = int32(1)
															} else {
																v726 = int32(-1)
															}
															v732 = v726
														} else {
															if v640 == int32(1) {
																v722 = v637
																v732 = v722
															} else {
																v654 = int32(1)
																v656 = int32(128)
																v657 = v644 << (uint(v654) % 32) & v656
																if v657 != v648<<(uint(v654)%32)&v656 {
																	v723 = v657
																	if v723 != 0 {
																		v726 = int32(1)
																	} else {
																		v726 = int32(-1)
																	}
																	v732 = v726
																} else {
																	if v640 < int32(3) {
																		v722 = v637
																		v732 = v722
																	} else {
																		v665 = int32(2)
																		v667 = int32(128)
																		v668 = v644 << (uint(v665) % 32) & v667
																		if v668 != v648<<(uint(v665)%32)&v667 {
																			v723 = v668
																			if v723 != 0 {
																				v726 = int32(1)
																			} else {
																				v726 = int32(-1)
																			}
																			v732 = v726
																		} else {
																			if v640 == int32(3) {
																				v722 = v637
																				v732 = v722
																			} else {
																				v676 = int32(3)
																				v678 = int32(128)
																				v679 = v644 << (uint(v676) % 32) & v678
																				if v679 != v648<<(uint(v676)%32)&v678 {
																					v723 = v679
																					if v723 != 0 {
																						v726 = int32(1)
																					} else {
																						v726 = int32(-1)
																					}
																					v732 = v726
																				} else {
																					if v640 < int32(5) {
																						v722 = v637
																						v732 = v722
																					} else {
																						v687 = int32(4)
																						v689 = int32(128)
																						v690 = v644 << (uint(v687) % 32) & v689
																						if v690 != v648<<(uint(v687)%32)&v689 {
																							v723 = v690
																							if v723 != 0 {
																								v726 = int32(1)
																							} else {
																								v726 = int32(-1)
																							}
																							v732 = v726
																						} else {
																							if v640 == int32(5) {
																								v722 = v637
																								v732 = v722
																							} else {
																								v698 = int32(5)
																								v700 = int32(128)
																								v701 = v644 << (uint(v698) % 32) & v700
																								if v701 != v648<<(uint(v698)%32)&v700 {
																									v723 = v701
																									if v723 != 0 {
																										v726 = int32(1)
																									} else {
																										v726 = int32(-1)
																									}
																									v732 = v726
																								} else {
																									if v640 < int32(7) {
																										v722 = v637
																										v732 = v722
																									} else {
																										v709 = int32(6)
																										v711 = int32(128)
																										v712 = v644 << (uint(v709) % 32) & v711
																										if v712 != v648<<(uint(v709)%32)&v711 {
																											v723 = v712
																											if v723 != 0 {
																												v726 = int32(1)
																											} else {
																												v726 = int32(-1)
																											}
																											v732 = v726
																										} else {
																											v722 = v637
																											v732 = v722
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
												switch v18&int32(_a_F_inet_gist_consistent_0) - int32(18) {
												case 0:
													v782 = v732
													return base.B2i32(v782 == int32(0))
												case 1:
													v775 = v732
													return base.B2i32(v775 != int32(0))
												case 2:
													return int32(base.Ui32(v732) >> (uint(int32(31)) % 32))
												case 3:
													return base.B2i32(v732 <= int32(0))
												case 4:
													return base.B2i32(int32(0) < v732)
												case 5:
													return int32(base.Ui32(v732^int32(-1)) >> (uint(int32(31)) % 32))
												default:
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v754 = m.ExcPending
													if v754 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_inet_gist_consistent_1), int32(0))
														mBase = m.M
														v758 = m.ExcPending
														if v758 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_inet_gist_consistent_2), int32(327), int32(_a_F_inet_gist_consistent_3))
															mBase = m.M
															v763 = m.ExcPending
															if v763 != 0 {
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
												return int32(0)
											}
										}
									}
								}
							}
						}
					}
				case 6:
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+16)))
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v53)+12)))
					if v55&int32(1) == int32(0) {
						v201 = v19 + int32(4)
						v203 = v47 + int32(2)
						v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+3)))
						v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
						if base.Ui32(v204) < base.Ui32(v205) {
							v207 = v204
						} else {
							v207 = v205
						}
						v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
						if base.Ui32(v207) < base.Ui32(v208) {
							v210 = v207
						} else {
							v210 = v208
						}
						v215 = base.I32_div_s(v210, int32(8))
						v216 = F_memcmp(m, v201, v203, v215)
						mBase = m.M
						if v216 != 0 {
							v302 = v216
							v312 = v302
						} else {
							v217 = int32(0)
							v220 = v210 - v215<<(uint(int32(3))%32)
							if v220 <= v217 {
								v302 = v217
								v312 = v302
							} else {
								v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201+v215))))
								v225 = int32(128)
								v226 = v224 & v225
								v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203+v215))))
								if v226 != v228&v225 {
									v303 = v226
									if v303 != 0 {
										v306 = int32(1)
									} else {
										v306 = int32(-1)
									}
									v312 = v306
								} else {
									if v220 == int32(1) {
										v302 = v217
										v312 = v302
									} else {
										v234 = int32(1)
										v236 = int32(128)
										v237 = v224 << (uint(v234) % 32) & v236
										if v237 != v228<<(uint(v234)%32)&v236 {
											v303 = v237
											if v303 != 0 {
												v306 = int32(1)
											} else {
												v306 = int32(-1)
											}
											v312 = v306
										} else {
											if v220 < int32(3) {
												v302 = v217
												v312 = v302
											} else {
												v245 = int32(2)
												v247 = int32(128)
												v248 = v224 << (uint(v245) % 32) & v247
												if v248 != v228<<(uint(v245)%32)&v247 {
													v303 = v248
													if v303 != 0 {
														v306 = int32(1)
													} else {
														v306 = int32(-1)
													}
													v312 = v306
												} else {
													if v220 == int32(3) {
														v302 = v217
														v312 = v302
													} else {
														v256 = int32(3)
														v258 = int32(128)
														v259 = v224 << (uint(v256) % 32) & v258
														if v259 != v228<<(uint(v256)%32)&v258 {
															v303 = v259
															if v303 != 0 {
																v306 = int32(1)
															} else {
																v306 = int32(-1)
															}
															v312 = v306
														} else {
															if v220 < int32(5) {
																v302 = v217
																v312 = v302
															} else {
																v267 = int32(4)
																v269 = int32(128)
																v270 = v224 << (uint(v267) % 32) & v269
																if v270 != v228<<(uint(v267)%32)&v269 {
																	v303 = v270
																	if v303 != 0 {
																		v306 = int32(1)
																	} else {
																		v306 = int32(-1)
																	}
																	v312 = v306
																} else {
																	if v220 == int32(5) {
																		v302 = v217
																		v312 = v302
																	} else {
																		v278 = int32(5)
																		v280 = int32(128)
																		v281 = v224 << (uint(v278) % 32) & v280
																		if v281 != v228<<(uint(v278)%32)&v280 {
																			v303 = v281
																			if v303 != 0 {
																				v306 = int32(1)
																			} else {
																				v306 = int32(-1)
																			}
																			v312 = v306
																		} else {
																			if v220 < int32(7) {
																				v302 = v217
																				v312 = v302
																			} else {
																				v289 = int32(6)
																				v291 = int32(128)
																				v292 = v224 << (uint(v289) % 32) & v291
																				if v292 != v228<<(uint(v289)%32)&v291 {
																					v303 = v292
																					if v303 != 0 {
																						v306 = int32(1)
																					} else {
																						v306 = int32(-1)
																					}
																					v312 = v306
																				} else {
																					v302 = v217
																					v312 = v302
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
						v313 = v312
						return base.B2i32(v313 == int32(0))
					} else {
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
						v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
						if base.Ui32(v61) < base.Ui32(v60) {
							v201 = v19 + int32(4)
							v203 = v47 + int32(2)
							v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+3)))
							v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
							if base.Ui32(v204) < base.Ui32(v205) {
								v207 = v204
							} else {
								v207 = v205
							}
							v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
							if base.Ui32(v207) < base.Ui32(v208) {
								v210 = v207
							} else {
								v210 = v208
							}
							v215 = base.I32_div_s(v210, int32(8))
							v216 = F_memcmp(m, v201, v203, v215)
							mBase = m.M
							if v216 != 0 {
								v302 = v216
								v312 = v302
							} else {
								v217 = int32(0)
								v220 = v210 - v215<<(uint(int32(3))%32)
								if v220 <= v217 {
									v302 = v217
									v312 = v302
								} else {
									v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201+v215))))
									v225 = int32(128)
									v226 = v224 & v225
									v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203+v215))))
									if v226 != v228&v225 {
										v303 = v226
										if v303 != 0 {
											v306 = int32(1)
										} else {
											v306 = int32(-1)
										}
										v312 = v306
									} else {
										if v220 == int32(1) {
											v302 = v217
											v312 = v302
										} else {
											v234 = int32(1)
											v236 = int32(128)
											v237 = v224 << (uint(v234) % 32) & v236
											if v237 != v228<<(uint(v234)%32)&v236 {
												v303 = v237
												if v303 != 0 {
													v306 = int32(1)
												} else {
													v306 = int32(-1)
												}
												v312 = v306
											} else {
												if v220 < int32(3) {
													v302 = v217
													v312 = v302
												} else {
													v245 = int32(2)
													v247 = int32(128)
													v248 = v224 << (uint(v245) % 32) & v247
													if v248 != v228<<(uint(v245)%32)&v247 {
														v303 = v248
														if v303 != 0 {
															v306 = int32(1)
														} else {
															v306 = int32(-1)
														}
														v312 = v306
													} else {
														if v220 == int32(3) {
															v302 = v217
															v312 = v302
														} else {
															v256 = int32(3)
															v258 = int32(128)
															v259 = v224 << (uint(v256) % 32) & v258
															if v259 != v228<<(uint(v256)%32)&v258 {
																v303 = v259
																if v303 != 0 {
																	v306 = int32(1)
																} else {
																	v306 = int32(-1)
																}
																v312 = v306
															} else {
																if v220 < int32(5) {
																	v302 = v217
																	v312 = v302
																} else {
																	v267 = int32(4)
																	v269 = int32(128)
																	v270 = v224 << (uint(v267) % 32) & v269
																	if v270 != v228<<(uint(v267)%32)&v269 {
																		v303 = v270
																		if v303 != 0 {
																			v306 = int32(1)
																		} else {
																			v306 = int32(-1)
																		}
																		v312 = v306
																	} else {
																		if v220 == int32(5) {
																			v302 = v217
																			v312 = v302
																		} else {
																			v278 = int32(5)
																			v280 = int32(128)
																			v281 = v224 << (uint(v278) % 32) & v280
																			if v281 != v228<<(uint(v278)%32)&v280 {
																				v303 = v281
																				if v303 != 0 {
																					v306 = int32(1)
																				} else {
																					v306 = int32(-1)
																				}
																				v312 = v306
																			} else {
																				if v220 < int32(7) {
																					v302 = v217
																					v312 = v302
																				} else {
																					v289 = int32(6)
																					v291 = int32(128)
																					v292 = v224 << (uint(v289) % 32) & v291
																					if v292 != v228<<(uint(v289)%32)&v291 {
																						v303 = v292
																						if v303 != 0 {
																							v306 = int32(1)
																						} else {
																							v306 = int32(-1)
																						}
																						v312 = v306
																					} else {
																						v302 = v217
																						v312 = v302
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
							v313 = v312
							return base.B2i32(v313 == int32(0))
						} else {
							return int32(0)
						}
					}
				case 7:
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+16)))
					v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+v64)+12)))
					if v66&int32(1) == int32(0) {
						v201 = v19 + int32(4)
						v203 = v47 + int32(2)
						v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+3)))
						v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
						if base.Ui32(v204) < base.Ui32(v205) {
							v207 = v204
						} else {
							v207 = v205
						}
						v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
						if base.Ui32(v207) < base.Ui32(v208) {
							v210 = v207
						} else {
							v210 = v208
						}
						v215 = base.I32_div_s(v210, int32(8))
						v216 = F_memcmp(m, v201, v203, v215)
						mBase = m.M
						if v216 != 0 {
							v302 = v216
							v312 = v302
						} else {
							v217 = int32(0)
							v220 = v210 - v215<<(uint(int32(3))%32)
							if v220 <= v217 {
								v302 = v217
								v312 = v302
							} else {
								v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201+v215))))
								v225 = int32(128)
								v226 = v224 & v225
								v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203+v215))))
								if v226 != v228&v225 {
									v303 = v226
									if v303 != 0 {
										v306 = int32(1)
									} else {
										v306 = int32(-1)
									}
									v312 = v306
								} else {
									if v220 == int32(1) {
										v302 = v217
										v312 = v302
									} else {
										v234 = int32(1)
										v236 = int32(128)
										v237 = v224 << (uint(v234) % 32) & v236
										if v237 != v228<<(uint(v234)%32)&v236 {
											v303 = v237
											if v303 != 0 {
												v306 = int32(1)
											} else {
												v306 = int32(-1)
											}
											v312 = v306
										} else {
											if v220 < int32(3) {
												v302 = v217
												v312 = v302
											} else {
												v245 = int32(2)
												v247 = int32(128)
												v248 = v224 << (uint(v245) % 32) & v247
												if v248 != v228<<(uint(v245)%32)&v247 {
													v303 = v248
													if v303 != 0 {
														v306 = int32(1)
													} else {
														v306 = int32(-1)
													}
													v312 = v306
												} else {
													if v220 == int32(3) {
														v302 = v217
														v312 = v302
													} else {
														v256 = int32(3)
														v258 = int32(128)
														v259 = v224 << (uint(v256) % 32) & v258
														if v259 != v228<<(uint(v256)%32)&v258 {
															v303 = v259
															if v303 != 0 {
																v306 = int32(1)
															} else {
																v306 = int32(-1)
															}
															v312 = v306
														} else {
															if v220 < int32(5) {
																v302 = v217
																v312 = v302
															} else {
																v267 = int32(4)
																v269 = int32(128)
																v270 = v224 << (uint(v267) % 32) & v269
																if v270 != v228<<(uint(v267)%32)&v269 {
																	v303 = v270
																	if v303 != 0 {
																		v306 = int32(1)
																	} else {
																		v306 = int32(-1)
																	}
																	v312 = v306
																} else {
																	if v220 == int32(5) {
																		v302 = v217
																		v312 = v302
																	} else {
																		v278 = int32(5)
																		v280 = int32(128)
																		v281 = v224 << (uint(v278) % 32) & v280
																		if v281 != v228<<(uint(v278)%32)&v280 {
																			v303 = v281
																			if v303 != 0 {
																				v306 = int32(1)
																			} else {
																				v306 = int32(-1)
																			}
																			v312 = v306
																		} else {
																			if v220 < int32(7) {
																				v302 = v217
																				v312 = v302
																			} else {
																				v289 = int32(6)
																				v291 = int32(128)
																				v292 = v224 << (uint(v289) % 32) & v291
																				if v292 != v228<<(uint(v289)%32)&v291 {
																					v303 = v292
																					if v303 != 0 {
																						v306 = int32(1)
																					} else {
																						v306 = int32(-1)
																					}
																					v312 = v306
																				} else {
																					v302 = v217
																					v312 = v302
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
						v313 = v312
						return base.B2i32(v313 == int32(0))
					} else {
						v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
						if base.Ui32(v72) <= base.Ui32(v71) {
							v201 = v19 + int32(4)
							v203 = v47 + int32(2)
							v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+3)))
							v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
							if base.Ui32(v204) < base.Ui32(v205) {
								v207 = v204
							} else {
								v207 = v205
							}
							v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
							if base.Ui32(v207) < base.Ui32(v208) {
								v210 = v207
							} else {
								v210 = v208
							}
							v215 = base.I32_div_s(v210, int32(8))
							v216 = F_memcmp(m, v201, v203, v215)
							mBase = m.M
							if v216 != 0 {
								v302 = v216
								v312 = v302
							} else {
								v217 = int32(0)
								v220 = v210 - v215<<(uint(int32(3))%32)
								if v220 <= v217 {
									v302 = v217
									v312 = v302
								} else {
									v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201+v215))))
									v225 = int32(128)
									v226 = v224 & v225
									v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203+v215))))
									if v226 != v228&v225 {
										v303 = v226
										if v303 != 0 {
											v306 = int32(1)
										} else {
											v306 = int32(-1)
										}
										v312 = v306
									} else {
										if v220 == int32(1) {
											v302 = v217
											v312 = v302
										} else {
											v234 = int32(1)
											v236 = int32(128)
											v237 = v224 << (uint(v234) % 32) & v236
											if v237 != v228<<(uint(v234)%32)&v236 {
												v303 = v237
												if v303 != 0 {
													v306 = int32(1)
												} else {
													v306 = int32(-1)
												}
												v312 = v306
											} else {
												if v220 < int32(3) {
													v302 = v217
													v312 = v302
												} else {
													v245 = int32(2)
													v247 = int32(128)
													v248 = v224 << (uint(v245) % 32) & v247
													if v248 != v228<<(uint(v245)%32)&v247 {
														v303 = v248
														if v303 != 0 {
															v306 = int32(1)
														} else {
															v306 = int32(-1)
														}
														v312 = v306
													} else {
														if v220 == int32(3) {
															v302 = v217
															v312 = v302
														} else {
															v256 = int32(3)
															v258 = int32(128)
															v259 = v224 << (uint(v256) % 32) & v258
															if v259 != v228<<(uint(v256)%32)&v258 {
																v303 = v259
																if v303 != 0 {
																	v306 = int32(1)
																} else {
																	v306 = int32(-1)
																}
																v312 = v306
															} else {
																if v220 < int32(5) {
																	v302 = v217
																	v312 = v302
																} else {
																	v267 = int32(4)
																	v269 = int32(128)
																	v270 = v224 << (uint(v267) % 32) & v269
																	if v270 != v228<<(uint(v267)%32)&v269 {
																		v303 = v270
																		if v303 != 0 {
																			v306 = int32(1)
																		} else {
																			v306 = int32(-1)
																		}
																		v312 = v306
																	} else {
																		if v220 == int32(5) {
																			v302 = v217
																			v312 = v302
																		} else {
																			v278 = int32(5)
																			v280 = int32(128)
																			v281 = v224 << (uint(v278) % 32) & v280
																			if v281 != v228<<(uint(v278)%32)&v280 {
																				v303 = v281
																				if v303 != 0 {
																					v306 = int32(1)
																				} else {
																					v306 = int32(-1)
																				}
																				v312 = v306
																			} else {
																				if v220 < int32(7) {
																					v302 = v217
																					v312 = v302
																				} else {
																					v289 = int32(6)
																					v291 = int32(128)
																					v292 = v224 << (uint(v289) % 32) & v291
																					if v292 != v228<<(uint(v289)%32)&v291 {
																						v303 = v292
																						if v303 != 0 {
																							v306 = int32(1)
																						} else {
																							v306 = int32(-1)
																						}
																						v312 = v306
																					} else {
																						v302 = v217
																						v312 = v302
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
							v313 = v312
							return base.B2i32(v313 == int32(0))
						} else {
							return int32(0)
						}
					}
				case 8:
					v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
					v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
					if base.Ui32(v77) < base.Ui32(v78) {
						v201 = v19 + int32(4)
						v203 = v47 + int32(2)
						v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+3)))
						v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
						if base.Ui32(v204) < base.Ui32(v205) {
							v207 = v204
						} else {
							v207 = v205
						}
						v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
						if base.Ui32(v207) < base.Ui32(v208) {
							v210 = v207
						} else {
							v210 = v208
						}
						v215 = base.I32_div_s(v210, int32(8))
						v216 = F_memcmp(m, v201, v203, v215)
						mBase = m.M
						if v216 != 0 {
							v302 = v216
							v312 = v302
						} else {
							v217 = int32(0)
							v220 = v210 - v215<<(uint(int32(3))%32)
							if v220 <= v217 {
								v302 = v217
								v312 = v302
							} else {
								v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201+v215))))
								v225 = int32(128)
								v226 = v224 & v225
								v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203+v215))))
								if v226 != v228&v225 {
									v303 = v226
									if v303 != 0 {
										v306 = int32(1)
									} else {
										v306 = int32(-1)
									}
									v312 = v306
								} else {
									if v220 == int32(1) {
										v302 = v217
										v312 = v302
									} else {
										v234 = int32(1)
										v236 = int32(128)
										v237 = v224 << (uint(v234) % 32) & v236
										if v237 != v228<<(uint(v234)%32)&v236 {
											v303 = v237
											if v303 != 0 {
												v306 = int32(1)
											} else {
												v306 = int32(-1)
											}
											v312 = v306
										} else {
											if v220 < int32(3) {
												v302 = v217
												v312 = v302
											} else {
												v245 = int32(2)
												v247 = int32(128)
												v248 = v224 << (uint(v245) % 32) & v247
												if v248 != v228<<(uint(v245)%32)&v247 {
													v303 = v248
													if v303 != 0 {
														v306 = int32(1)
													} else {
														v306 = int32(-1)
													}
													v312 = v306
												} else {
													if v220 == int32(3) {
														v302 = v217
														v312 = v302
													} else {
														v256 = int32(3)
														v258 = int32(128)
														v259 = v224 << (uint(v256) % 32) & v258
														if v259 != v228<<(uint(v256)%32)&v258 {
															v303 = v259
															if v303 != 0 {
																v306 = int32(1)
															} else {
																v306 = int32(-1)
															}
															v312 = v306
														} else {
															if v220 < int32(5) {
																v302 = v217
																v312 = v302
															} else {
																v267 = int32(4)
																v269 = int32(128)
																v270 = v224 << (uint(v267) % 32) & v269
																if v270 != v228<<(uint(v267)%32)&v269 {
																	v303 = v270
																	if v303 != 0 {
																		v306 = int32(1)
																	} else {
																		v306 = int32(-1)
																	}
																	v312 = v306
																} else {
																	if v220 == int32(5) {
																		v302 = v217
																		v312 = v302
																	} else {
																		v278 = int32(5)
																		v280 = int32(128)
																		v281 = v224 << (uint(v278) % 32) & v280
																		if v281 != v228<<(uint(v278)%32)&v280 {
																			v303 = v281
																			if v303 != 0 {
																				v306 = int32(1)
																			} else {
																				v306 = int32(-1)
																			}
																			v312 = v306
																		} else {
																			if v220 < int32(7) {
																				v302 = v217
																				v312 = v302
																			} else {
																				v289 = int32(6)
																				v291 = int32(128)
																				v292 = v224 << (uint(v289) % 32) & v291
																				if v292 != v228<<(uint(v289)%32)&v291 {
																					v303 = v292
																					if v303 != 0 {
																						v306 = int32(1)
																					} else {
																						v306 = int32(-1)
																					}
																					v312 = v306
																				} else {
																					v302 = v217
																					v312 = v302
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
						v313 = v312
						return base.B2i32(v313 == int32(0))
					} else {
						return int32(0)
					}
				}
			}
		}
	}
}
