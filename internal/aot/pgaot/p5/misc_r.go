package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ReadControlFile(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 float64
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
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
	var v155 int32
	_ = v155
	var v158 float64
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	v4 = m.G0
	v6 = v4 - int32(336)
	m.G0 = v6
	v10 = F_BasicOpenFile(m, int32(_a_F_ReadControlFile_0), int32(2))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		if int32(0) <= v10 {
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(167772169)
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[1]))
			v20 = int32(312)
			v21 = F_read(m, v10, v19, v20)
			mBase = m.M
			if v21 != v20 {
				F_errstart_cold(m, int32(24), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					if v21 < int32(0) {
						F_errcode_for_file_access(m)
						mBase = m.M
						v198 = m.ExcPending
						if v198 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+272)) = int32(_a_F_ReadControlFile_0)
							F_errmsg(m, int32(_a_F_ReadControlFile_1), v6+int32(272))
							mBase = m.M
							v205 = m.ExcPending
							if v205 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_3), int32(_a_F_ReadControlFile_4))
								mBase = m.M
								v210 = m.ExcPending
								if v210 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						F_errcode(m, int32(16779816))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+296)) = int32(312)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+292)) = v21
							*(*int32)(unsafe.Add(mBase, uint32(v6)+288)) = int32(_a_F_ReadControlFile_0)
							F_errmsg(m, int32(_a_F_ReadControlFile_5), v6+int32(288))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_6), int32(_a_F_ReadControlFile_4))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
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
				v49 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[0]))
				v50 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v49))) = v50
				v52 = F_close(m, v10)
				mBase = m.M
				v55 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[1]))
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
				if v56&int32(_a_F_ReadControlFile_7) != 0 {
					v59 = v50
				} else {
					v59 = v56
				}
				if v59 != 0 {
					F_errstart_cold(m, int32(22), int32(0))
					mBase = m.M
					v214 = m.ExcPending
					if v214 != 0 {
						return
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v217 = m.ExcPending
						if v217 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_ReadControlFile_8), int32(0))
							mBase = m.M
							v221 = m.ExcPending
							if v221 != 0 {
								return
							} else {
								v223 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[1]))
								v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v6)+256)) = v224
								*(*int32)(unsafe.Add(mBase, uint32(v6)+260)) = v224
								*(*int64)(unsafe.Add(mBase, uint32(v6)+264)) = int64(8181912700785)
								v232 = F_errdetail(m, int32(_a_F_ReadControlFile_9), v6+int32(256))
								mBase = m.M
								v233 = m.ExcPending
								if v233 != 0 {
									return
								} else {
									F_errhint(m, int32(_a_F_ReadControlFile_10), int32(0))
									mBase = m.M
									v237 = m.ExcPending
									if v237 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_11), int32(_a_F_ReadControlFile_4))
										mBase = m.M
										v242 = m.ExcPending
										if v242 != 0 {
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
					if v56 != int32(1905) {
						F_errstart_cold(m, int32(22), int32(0))
						mBase = m.M
						v246 = m.ExcPending
						if v246 != 0 {
							return
						} else {
							F_errcode(m, int32(325))
							mBase = m.M
							v249 = m.ExcPending
							if v249 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_ReadControlFile_8), int32(0))
								mBase = m.M
								v253 = m.ExcPending
								if v253 != 0 {
									return
								} else {
									v255 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[1]))
									v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v6)+240)) = v256
									*(*int32)(unsafe.Add(mBase, uint32(v6)+244)) = int32(1905)
									v263 = F_errdetail(m, int32(_a_F_ReadControlFile_12), v6+int32(240))
									mBase = m.M
									v264 = m.ExcPending
									if v264 != 0 {
										return
									} else {
										F_errhint(m, int32(_a_F_ReadControlFile_13), int32(0))
										mBase = m.M
										v268 = m.ExcPending
										if v268 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_14), int32(_a_F_ReadControlFile_4))
											mBase = m.M
											v273 = m.ExcPending
											if v273 != 0 {
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
						v62 = int32(-1)
						v64 = m.Env.Pgmem_crc32c(m, v62, v55, int32(308))
						mBase = m.M
						v66 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[1]))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+308))
						if v64^v67 != v62 {
							F_errstart_cold(m, int32(22), int32(0))
							mBase = m.M
							v277 = m.ExcPending
							if v277 != 0 {
								return
							} else {
								F_errcode(m, int32(325))
								mBase = m.M
								v280 = m.ExcPending
								if v280 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_ReadControlFile_15), int32(0))
									mBase = m.M
									v284 = m.ExcPending
									if v284 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_16), int32(_a_F_ReadControlFile_4))
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
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
							if v71 != int32(202609165) {
								F_errstart_cold(m, int32(22), int32(0))
								mBase = m.M
								v293 = m.ExcPending
								if v293 != 0 {
									return
								} else {
									F_errcode(m, int32(325))
									mBase = m.M
									v296 = m.ExcPending
									if v296 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_ReadControlFile_8), int32(0))
										mBase = m.M
										v300 = m.ExcPending
										if v300 != 0 {
											return
										} else {
											v302 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[1]))
											v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)+12))
											v304 = int32(_a_F_ReadControlFile_17)
											*(*int32)(unsafe.Add(mBase, uint32(v6)+224)) = v304
											*(*int32)(unsafe.Add(mBase, uint32(v6)+228)) = v303
											*(*int32)(unsafe.Add(mBase, uint32(v6)+232)) = v304
											*(*int32)(unsafe.Add(mBase, uint32(v6)+236)) = int32(202609165)
											v314 = F_errdetail(m, int32(_a_F_ReadControlFile_18), v6+int32(224))
											mBase = m.M
											v315 = m.ExcPending
											if v315 != 0 {
												return
											} else {
												F_errhint(m, int32(_a_F_ReadControlFile_13), int32(0))
												mBase = m.M
												v319 = m.ExcPending
												if v319 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_19), int32(_a_F_ReadControlFile_4))
													mBase = m.M
													v324 = m.ExcPending
													if v324 != 0 {
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
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v66)+212))
								if v74 != int32(8) {
									F_errstart_cold(m, int32(22), int32(0))
									mBase = m.M
									v328 = m.ExcPending
									if v328 != 0 {
										return
									} else {
										F_errcode(m, int32(325))
										mBase = m.M
										v331 = m.ExcPending
										if v331 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_ReadControlFile_8), int32(0))
											mBase = m.M
											v335 = m.ExcPending
											if v335 != 0 {
												return
											} else {
												v337 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[1]))
												v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+212))
												v339 = int32(_a_F_ReadControlFile_20)
												*(*int32)(unsafe.Add(mBase, uint32(v6)+208)) = v339
												*(*int32)(unsafe.Add(mBase, uint32(v6)+212)) = v338
												*(*int32)(unsafe.Add(mBase, uint32(v6)+216)) = v339
												*(*int32)(unsafe.Add(mBase, uint32(v6)+220)) = int32(8)
												v349 = F_errdetail(m, int32(_a_F_ReadControlFile_18), v6+int32(208))
												mBase = m.M
												v350 = m.ExcPending
												if v350 != 0 {
													return
												} else {
													F_errhint(m, int32(_a_F_ReadControlFile_13), int32(0))
													mBase = m.M
													v354 = m.ExcPending
													if v354 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_21), int32(_a_F_ReadControlFile_4))
														mBase = m.M
														v359 = m.ExcPending
														if v359 != 0 {
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
									v77 = *(*float64)(unsafe.Add(mBase, uint32(v66)+216))
									if base.F64_ne(v77, float64(1.234567e+06)) != 0 {
										F_errstart_cold(m, int32(22), int32(0))
										mBase = m.M
										v363 = m.ExcPending
										if v363 != 0 {
											return
										} else {
											F_errcode(m, int32(325))
											mBase = m.M
											v366 = m.ExcPending
											if v366 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_ReadControlFile_8), int32(0))
												mBase = m.M
												v370 = m.ExcPending
												if v370 != 0 {
													return
												} else {
													v373 = F_errdetail(m, int32(_a_F_ReadControlFile_22), int32(0))
													mBase = m.M
													v374 = m.ExcPending
													if v374 != 0 {
														return
													} else {
														F_errhint(m, int32(_a_F_ReadControlFile_13), int32(0))
														mBase = m.M
														v378 = m.ExcPending
														if v378 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_23), int32(_a_F_ReadControlFile_4))
															mBase = m.M
															v383 = m.ExcPending
															if v383 != 0 {
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
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v66)+224))
										if v80 != int32(_a_F_ReadControlFile_24) {
											F_errstart_cold(m, int32(22), int32(0))
											mBase = m.M
											v387 = m.ExcPending
											if v387 != 0 {
												return
											} else {
												F_errcode(m, int32(325))
												mBase = m.M
												v390 = m.ExcPending
												if v390 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_ReadControlFile_8), int32(0))
													mBase = m.M
													v394 = m.ExcPending
													if v394 != 0 {
														return
													} else {
														v396 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[1]))
														v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)+224))
														v398 = int32(_a_F_ReadControlFile_25)
														*(*int32)(unsafe.Add(mBase, uint32(v6)+192)) = v398
														*(*int32)(unsafe.Add(mBase, uint32(v6)+196)) = v397
														*(*int32)(unsafe.Add(mBase, uint32(v6)+200)) = v398
														*(*int32)(unsafe.Add(mBase, uint32(v6)+204)) = int32(_a_F_ReadControlFile_24)
														v408 = F_errdetail(m, int32(_a_F_ReadControlFile_18), v6+int32(192))
														mBase = m.M
														v409 = m.ExcPending
														if v409 != 0 {
															return
														} else {
															F_errhint(m, int32(_a_F_ReadControlFile_26), int32(0))
															mBase = m.M
															v413 = m.ExcPending
															if v413 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_27), int32(_a_F_ReadControlFile_4))
																mBase = m.M
																v418 = m.ExcPending
																if v418 != 0 {
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
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v66)+228))
											if v83 != int32(_a_F_ReadControlFile_28) {
												F_errstart_cold(m, int32(22), int32(0))
												mBase = m.M
												v422 = m.ExcPending
												if v422 != 0 {
													return
												} else {
													F_errcode(m, int32(325))
													mBase = m.M
													v425 = m.ExcPending
													if v425 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_ReadControlFile_8), int32(0))
														mBase = m.M
														v429 = m.ExcPending
														if v429 != 0 {
															return
														} else {
															v431 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[1]))
															v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)+228))
															v433 = int32(_a_F_ReadControlFile_29)
															*(*int32)(unsafe.Add(mBase, uint32(v6)+176)) = v433
															*(*int32)(unsafe.Add(mBase, uint32(v6)+180)) = v432
															*(*int32)(unsafe.Add(mBase, uint32(v6)+184)) = v433
															*(*int32)(unsafe.Add(mBase, uint32(v6)+188)) = int32(_a_F_ReadControlFile_28)
															v443 = F_errdetail(m, int32(_a_F_ReadControlFile_18), v6+int32(176))
															mBase = m.M
															v444 = m.ExcPending
															if v444 != 0 {
																return
															} else {
																F_errhint(m, int32(_a_F_ReadControlFile_26), int32(0))
																mBase = m.M
																v448 = m.ExcPending
																if v448 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_30), int32(_a_F_ReadControlFile_4))
																	mBase = m.M
																	v453 = m.ExcPending
																	if v453 != 0 {
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
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v66)+232))
												if v86 != int32(32) {
													F_errstart_cold(m, int32(22), int32(0))
													mBase = m.M
													v457 = m.ExcPending
													if v457 != 0 {
														return
													} else {
														F_errcode(m, int32(325))
														mBase = m.M
														v460 = m.ExcPending
														if v460 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_ReadControlFile_8), int32(0))
															mBase = m.M
															v464 = m.ExcPending
															if v464 != 0 {
																return
															} else {
																v466 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[1]))
																v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)+232))
																v468 = int32(_a_F_ReadControlFile_31)
																*(*int32)(unsafe.Add(mBase, uint32(v6)+160)) = v468
																*(*int32)(unsafe.Add(mBase, uint32(v6)+164)) = v467
																*(*int32)(unsafe.Add(mBase, uint32(v6)+168)) = v468
																*(*int32)(unsafe.Add(mBase, uint32(v6)+172)) = int32(32)
																v478 = F_errdetail(m, int32(_a_F_ReadControlFile_18), v6+int32(160))
																mBase = m.M
																v479 = m.ExcPending
																if v479 != 0 {
																	return
																} else {
																	F_errhint(m, int32(_a_F_ReadControlFile_26), int32(0))
																	mBase = m.M
																	v483 = m.ExcPending
																	if v483 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_32), int32(_a_F_ReadControlFile_4))
																		mBase = m.M
																		v488 = m.ExcPending
																		if v488 != 0 {
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
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v66)+236))
													if v89 != int32(_a_F_ReadControlFile_24) {
														F_errstart_cold(m, int32(22), int32(0))
														mBase = m.M
														v492 = m.ExcPending
														if v492 != 0 {
															return
														} else {
															F_errcode(m, int32(325))
															mBase = m.M
															v495 = m.ExcPending
															if v495 != 0 {
																return
															} else {
																F_errmsg(m, int32(_a_F_ReadControlFile_8), int32(0))
																mBase = m.M
																v499 = m.ExcPending
																if v499 != 0 {
																	return
																} else {
																	v501 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[1]))
																	v502 = *(*int32)(unsafe.Add(mBase, uint32(v501)+236))
																	v503 = int32(_a_F_ReadControlFile_33)
																	*(*int32)(unsafe.Add(mBase, uint32(v6)+144)) = v503
																	*(*int32)(unsafe.Add(mBase, uint32(v6)+148)) = v502
																	*(*int32)(unsafe.Add(mBase, uint32(v6)+152)) = v503
																	*(*int32)(unsafe.Add(mBase, uint32(v6)+156)) = int32(_a_F_ReadControlFile_24)
																	v513 = F_errdetail(m, int32(_a_F_ReadControlFile_18), v6+int32(144))
																	mBase = m.M
																	v514 = m.ExcPending
																	if v514 != 0 {
																		return
																	} else {
																		F_errhint(m, int32(_a_F_ReadControlFile_26), int32(0))
																		mBase = m.M
																		v518 = m.ExcPending
																		if v518 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_34), int32(_a_F_ReadControlFile_4))
																			mBase = m.M
																			v523 = m.ExcPending
																			if v523 != 0 {
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
														v92 = *(*int32)(unsafe.Add(mBase, uint32(v66)+244))
														if v92 != int32(64) {
															F_errstart_cold(m, int32(22), int32(0))
															mBase = m.M
															v527 = m.ExcPending
															if v527 != 0 {
																return
															} else {
																F_errcode(m, int32(325))
																mBase = m.M
																v530 = m.ExcPending
																if v530 != 0 {
																	return
																} else {
																	F_errmsg(m, int32(_a_F_ReadControlFile_8), int32(0))
																	mBase = m.M
																	v534 = m.ExcPending
																	if v534 != 0 {
																		return
																	} else {
																		v536 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[1]))
																		v537 = *(*int32)(unsafe.Add(mBase, uint32(v536)+244))
																		v538 = int32(_a_F_ReadControlFile_35)
																		*(*int32)(unsafe.Add(mBase, uint32(v6)+128)) = v538
																		*(*int32)(unsafe.Add(mBase, uint32(v6)+132)) = v537
																		*(*int32)(unsafe.Add(mBase, uint32(v6)+136)) = v538
																		*(*int32)(unsafe.Add(mBase, uint32(v6)+140)) = int32(64)
																		v548 = F_errdetail(m, int32(_a_F_ReadControlFile_18), v6+int32(128))
																		mBase = m.M
																		v549 = m.ExcPending
																		if v549 != 0 {
																			return
																		} else {
																			F_errhint(m, int32(_a_F_ReadControlFile_26), int32(0))
																			mBase = m.M
																			v553 = m.ExcPending
																			if v553 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_36), int32(_a_F_ReadControlFile_4))
																				mBase = m.M
																				v558 = m.ExcPending
																				if v558 != 0 {
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
															v95 = *(*int32)(unsafe.Add(mBase, uint32(v66)+248))
															if v95 != int32(32) {
																F_errstart_cold(m, int32(22), int32(0))
																mBase = m.M
																v562 = m.ExcPending
																if v562 != 0 {
																	return
																} else {
																	F_errcode(m, int32(325))
																	mBase = m.M
																	v565 = m.ExcPending
																	if v565 != 0 {
																		return
																	} else {
																		F_errmsg(m, int32(_a_F_ReadControlFile_8), int32(0))
																		mBase = m.M
																		v569 = m.ExcPending
																		if v569 != 0 {
																			return
																		} else {
																			v571 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[1]))
																			v572 = *(*int32)(unsafe.Add(mBase, uint32(v571)+248))
																			v573 = int32(_a_F_ReadControlFile_37)
																			*(*int32)(unsafe.Add(mBase, uint32(v6)+112)) = v573
																			*(*int32)(unsafe.Add(mBase, uint32(v6)+116)) = v572
																			*(*int32)(unsafe.Add(mBase, uint32(v6)+120)) = v573
																			*(*int32)(unsafe.Add(mBase, uint32(v6)+124)) = int32(32)
																			v583 = F_errdetail(m, int32(_a_F_ReadControlFile_18), v6+int32(112))
																			mBase = m.M
																			v584 = m.ExcPending
																			if v584 != 0 {
																				return
																			} else {
																				F_errhint(m, int32(_a_F_ReadControlFile_26), int32(0))
																				mBase = m.M
																				v588 = m.ExcPending
																				if v588 != 0 {
																					return
																				} else {
																					F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_38), int32(_a_F_ReadControlFile_4))
																					mBase = m.M
																					v593 = m.ExcPending
																					if v593 != 0 {
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
																v98 = *(*int32)(unsafe.Add(mBase, uint32(v66)+252))
																if v98 != int32(1996) {
																	F_errstart_cold(m, int32(22), int32(0))
																	mBase = m.M
																	v597 = m.ExcPending
																	if v597 != 0 {
																		return
																	} else {
																		F_errcode(m, int32(325))
																		mBase = m.M
																		v600 = m.ExcPending
																		if v600 != 0 {
																			return
																		} else {
																			F_errmsg(m, int32(_a_F_ReadControlFile_8), int32(0))
																			mBase = m.M
																			v604 = m.ExcPending
																			if v604 != 0 {
																				return
																			} else {
																				v606 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[1]))
																				v607 = *(*int32)(unsafe.Add(mBase, uint32(v606)+252))
																				v608 = int32(_a_F_ReadControlFile_39)
																				*(*int32)(unsafe.Add(mBase, uint32(v6)+96)) = v608
																				*(*int32)(unsafe.Add(mBase, uint32(v6)+100)) = v607
																				*(*int32)(unsafe.Add(mBase, uint32(v6)+104)) = v608
																				*(*int32)(unsafe.Add(mBase, uint32(v6)+108)) = int32(1996)
																				v618 = F_errdetail(m, int32(_a_F_ReadControlFile_18), v6+int32(96))
																				mBase = m.M
																				v619 = m.ExcPending
																				if v619 != 0 {
																					return
																				} else {
																					F_errhint(m, int32(_a_F_ReadControlFile_26), int32(0))
																					mBase = m.M
																					v623 = m.ExcPending
																					if v623 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_40), int32(_a_F_ReadControlFile_4))
																						mBase = m.M
																						v628 = m.ExcPending
																						if v628 != 0 {
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
																	v101 = *(*int32)(unsafe.Add(mBase, uint32(v66)+256))
																	if v101 != int32(2048) {
																		F_errstart_cold(m, int32(22), int32(0))
																		mBase = m.M
																		v632 = m.ExcPending
																		if v632 != 0 {
																			return
																		} else {
																			F_errcode(m, int32(325))
																			mBase = m.M
																			v635 = m.ExcPending
																			if v635 != 0 {
																				return
																			} else {
																				F_errmsg(m, int32(_a_F_ReadControlFile_8), int32(0))
																				mBase = m.M
																				v639 = m.ExcPending
																				if v639 != 0 {
																					return
																				} else {
																					v641 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[1]))
																					v642 = *(*int32)(unsafe.Add(mBase, uint32(v641)+256))
																					v643 = int32(_a_F_ReadControlFile_41)
																					*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = v643
																					*(*int32)(unsafe.Add(mBase, uint32(v6)+84)) = v642
																					*(*int32)(unsafe.Add(mBase, uint32(v6)+88)) = v643
																					*(*int32)(unsafe.Add(mBase, uint32(v6)+92)) = int32(2048)
																					v653 = F_errdetail(m, int32(_a_F_ReadControlFile_18), v6+int32(80))
																					mBase = m.M
																					v654 = m.ExcPending
																					if v654 != 0 {
																						return
																					} else {
																						F_errhint(m, int32(_a_F_ReadControlFile_26), int32(0))
																						mBase = m.M
																						v658 = m.ExcPending
																						if v658 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_42), int32(_a_F_ReadControlFile_4))
																							mBase = m.M
																							v663 = m.ExcPending
																							if v663 != 0 {
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
																		v105 = *(*int32)(unsafe.Add(mBase, uint32(v66)+240))
																		*(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[2])) = v105
																		if v105&(v105-int32(1))|base.B2i32(v105 <= int32(0))|base.B2i32(base.Ui32(int32(1072693249)) <= base.Ui32(v105+int32(-1048576))) != 0 {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v667 = m.ExcPending
																			if v667 != 0 {
																				return
																			} else {
																				F_errcode(m, int32(50856066))
																				mBase = m.M
																				v670 = m.ExcPending
																				if v670 != 0 {
																					return
																				} else {
																					v672 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[2]))
																					*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v672
																					F_errmsg_plural(m, int32(_a_F_ReadControlFile_43), int32(_a_F_ReadControlFile_44), v672, v6-int32(-64))
																					mBase = m.M
																					v679 = m.ExcPending
																					if v679 != 0 {
																						return
																					} else {
																						v682 = F_errdetail(m, int32(_a_F_ReadControlFile_45), int32(0))
																						mBase = m.M
																						v683 = m.ExcPending
																						if v683 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_46), int32(_a_F_ReadControlFile_4))
																							mBase = m.M
																							v688 = m.ExcPending
																							if v688 != 0 {
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
																			*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v105
																			v120 = v6 + int32(304)
																			v125 = F_pg_snprintf(m, v120, int32(20), int32(_a_F_ReadControlFile_47), v6+int32(48))
																			mBase = m.M
																			v126 = m.ExcPending
																			if v126 != 0 {
																				return
																			} else {
																				F_SetConfigOption(m, int32(_a_F_ReadControlFile_48), v120, int32(0), int32(1))
																				mBase = m.M
																				v131 = m.ExcPending
																				if v131 != 0 {
																					return
																				} else {
																					v133 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[3]))
																					v135 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[2]))
																					v137 = base.I32_div_s(v135, int32(_a_F_ReadControlFile_49))
																					v138 = base.I32_div_s(v133, v137)
																					if v138 <= int32(1) {
																						F_errstart_cold(m, int32(21), int32(0))
																						mBase = m.M
																						v692 = m.ExcPending
																						if v692 != 0 {
																							return
																						} else {
																							F_errcode(m, int32(50856066))
																							mBase = m.M
																							v695 = m.ExcPending
																							if v695 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(_a_F_ReadControlFile_48)
																								*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a_F_ReadControlFile_50)
																								F_errmsg(m, int32(_a_F_ReadControlFile_51), v6+int32(16))
																								mBase = m.M
																								v704 = m.ExcPending
																								if v704 != 0 {
																									return
																								} else {
																									F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_52), int32(_a_F_ReadControlFile_4))
																									mBase = m.M
																									v709 = m.ExcPending
																									if v709 != 0 {
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
																						v142 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[4]))
																						v143 = base.I32_div_s(v142, v137)
																						if v143 <= int32(1) {
																							F_errstart_cold(m, int32(21), int32(0))
																							mBase = m.M
																							v713 = m.ExcPending
																							if v713 != 0 {
																								return
																							} else {
																								F_errcode(m, int32(50856066))
																								mBase = m.M
																								v716 = m.ExcPending
																								if v716 != 0 {
																									return
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = int32(_a_F_ReadControlFile_48)
																									*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = int32(_a_F_ReadControlFile_53)
																									F_errmsg(m, int32(_a_F_ReadControlFile_51), v6+int32(32))
																									mBase = m.M
																									v725 = m.ExcPending
																									if v725 != 0 {
																										return
																									} else {
																										F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_54), int32(_a_F_ReadControlFile_4))
																										mBase = m.M
																										v730 = m.ExcPending
																										if v730 != 0 {
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
																							v148 = base.I32_div_s(v135, int32(_a_F_ReadControlFile_24))
																							*(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[5])) = v148*int32(_a_F_ReadControlFile_55) - int32(16)
																							v155 = int32(1)
																							v158 = *(*float64)(unsafe.Add(mBase, _c_F_ReadControlFile[6]))
																							v162 = base.I32_trunc_sat_f64_s(base.F64_div(base.F64_convert_i32_u(v143), base.F64_add(v158, float64(1))))
																							if v162 <= v155 {
																								v165 = v155
																							} else {
																								v165 = v162
																							}
																							*(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[7])) = v165
																							v171 = *(*int32)(unsafe.Add(mBase, _c_F_ReadControlFile[1]))
																							v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+268))
																							if v172 != 0 {
																								v173 = int32(_a_F_ReadControlFile_56)
																							} else {
																								v173 = int32(_a_F_ReadControlFile_57)
																							}
																							F_SetConfigOption(m, int32(_a_F_ReadControlFile_58), v173, int32(0), int32(1))
																							mBase = m.M
																							v177 = m.ExcPending
																							if v177 != 0 {
																								return
																							} else {
																								m.G0 = v6 + int32(336)
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
			}
		} else {
			F_errstart_cold(m, int32(24), int32(0))
			mBase = m.M
			v184 = m.ExcPending
			if v184 != 0 {
				return
			} else {
				F_errcode_for_file_access(m)
				mBase = m.M
				v186 = m.ExcPending
				if v186 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_ReadControlFile_0)
					F_errmsg(m, int32(_a_F_ReadControlFile_59), v6)
					mBase = m.M
					v191 = m.ExcPending
					if v191 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ReadControlFile_2), int32(_a_F_ReadControlFile_60), int32(_a_F_ReadControlFile_4))
						mBase = m.M
						v196 = m.ExcPending
						if v196 != 0 {
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
func F_ReadDimensionInt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if base.Ui32((v7-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v18 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_ReadDimensionInt[0])) = v18
		v24 = F_strtox_2(m, v6, l0, int32(10), int64(2147483648))
		mBase = m.M
		v27 = *(*int32)(unsafe.Add(mBase, _c_F_ReadDimensionInt[0]))
		if v27 == int32(68) {
			v30 = F_errsave_start(m, l2)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				if v30 == int32(0) {
					v55 = v18
					return v55
				} else {
					F_errcode(m, int32(261))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_ReadDimensionInt_0), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, l2, int32(_a_F_ReadDimensionInt_1), int32(542), int32(_a_F_ReadDimensionInt_2))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								return int32(0)
							}
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = base.I32_wrap_i64(v24)
			v55 = int32(1)
			return v55
		}
	} else {
		switch v7 - int32(43) {
		case 0, 2:
			v18 = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_ReadDimensionInt[0])) = v18
			v24 = F_strtox_2(m, v6, l0, int32(10), int64(2147483648))
			mBase = m.M
			v27 = *(*int32)(unsafe.Add(mBase, _c_F_ReadDimensionInt[0]))
			if v27 == int32(68) {
				v30 = F_errsave_start(m, l2)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					if v30 == int32(0) {
						v55 = v18
						return v55
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_ReadDimensionInt_0), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, l2, int32(_a_F_ReadDimensionInt_1), int32(542), int32(_a_F_ReadDimensionInt_2))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									return int32(0)
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = base.I32_wrap_i64(v24)
				v55 = int32(1)
				return v55
			}
		default:
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
			v55 = int32(1)
			return v55
		}
	}
}
func F_ReadRecord(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v36 int32
	_ = v36
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v92 int64
	_ = v92
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int64
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v190 int64
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v215 int32
	_ = v215
	var v218 int64
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v280 int32
	_ = v280
	var v281 int64
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
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
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v606 int64
	_ = v606
	var v608 int64
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v622 int64
	_ = v622
	var v623 int64
	_ = v623
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int64
	_ = v631
	var v632 int32
	_ = v632
	var v633 int64
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v647 int64
	_ = v647
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int64
	_ = v658
	var v664 int64
	_ = v664
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v683 int64
	_ = v683
	var v685 int64
	_ = v685
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v743 int32
	_ = v743
	var v749 int32
	_ = v749
	var v750 int64
	_ = v750
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v755 int64
	_ = v755
	var v756 int64
	_ = v756
	var v758 int64
	_ = v758
	var v759 int64
	_ = v759
	var v762 int64
	_ = v762
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v777 int64
	_ = v777
	var v779 int64
	_ = v779
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v792 int64
	_ = v792
	var v801 int64
	_ = v801
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v858 int64
	_ = v858
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v871 int64
	_ = v871
	var v875 int64
	_ = v875
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v907 int64
	_ = v907
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	v3 = l2
	v5 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(144)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)) = uint8(v3)
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+5)) = uint8(base.B2i32(v27 == int64(0)))
	*(*uint8)(unsafe.Add(mBase, _c_F_ReadRecord[0])) = uint8(v5)
	v36 = base.B2i32(l1 != int32(15))
	goto L1
L1:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[1]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v58 != v59 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	m.G0 = v21 + int32(144)
	return v929
L3:
	;
	goto L2
L4:
	;
	v823 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ReadRecord[0])) = uint8(v823)
	if v3 != 0 {
		goto L171
	} else {
		goto L172
	}
L5:
	;
	if v653 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L6:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v61 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v107 = v21 + int32(140)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+96))
	if v110 == int32(0) {
		v190 = int64(0)
		goto L21
	} else {
		goto L22
	}
L9:
	;
	F_pfree(m, v61)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[2]))
	v68 = int32(2)
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[3]))
	v75 = int32(0)
	v79 = base.B2i32(v74 != v75) & base.B2i32(v75 < v67)
	if v79 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	return int32(0)
L13:
	;
	goto L11
L14:
	;
	v80 = v67<<(uint(v68)%32) | int32(1)
	goto L16
L15:
	;
	v80 = v68
	goto L16
L16:
	;
	v85 = F_palloc(m, v80<<(uint(int32(4))%32)+int32(32))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = v80
	if v79 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v89 = v67
	goto L20
L19:
	;
	v89 = int32(1)
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+8)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = l0
	v92 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v85)+12)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(436)
	*(*int64)(unsafe.Add(mBase, uint32(v85)+20)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v85
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[1]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v100
	goto L8
L21:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v191 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L22:
	;
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v110)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+96)) = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+120)) = v116
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v109)+124))
	if v110 == v118 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+124)) = int32(0)
	goto L25
L24:
	;
	goto L25
L25:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+4)))
	if v122 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+112)) = v130
	v190 = v113
	goto L21
L27:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	if v125 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	F_pfree(m, v110)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L12
	} else {
		goto L37
	}
L30:
	;
	v130 = v125
	goto L33
L31:
	;
	goto L32
L32:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v109)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+116)) = v166
	*(*int32)(unsafe.Add(mBase, uint32(v109)+112)) = v166
	v190 = v113
	goto L21
L33:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+4)))
	if v144 != int32(1) {
		goto L26
	} else {
		goto L35
	}
L34:
	;
	goto L32
L35:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	if v147 != 0 {
		v130 = v147
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v190 = v113
	goto L21
L38:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v256 = v254 + int32(32)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v254)+24))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v254)+20))
	if v257 == v258 {
		goto L47
	} else {
		goto L48
	}
L39:
	;
	v195 = l0 + int32(28)
	if v191 == v195 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	goto L41
L41:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	v218 = *(*int64)(unsafe.Add(mBase, uint32(v215-int32(12))))
	if base.Ui64(v190) <= base.Ui64(v218) {
		goto L38
	} else {
		goto L43
	}
L42:
	;
	goto L38
L43:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+4)) = v221
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	*(*int32)(unsafe.Add(mBase, uint32(v221))) = v223
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v230 = F_hash_search(m, v225, v215-int32(28), int32(2), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L12
	} else {
		goto L44
	}
L44:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v232 == int32(0) {
		goto L38
	} else {
		goto L45
	}
L45:
	;
	if v232 != v195 {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	goto L42
L47:
	;
	v321 = *(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[3]))
	if v321 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L48:
	;
	v266 = v257
	goto L49
L49:
	;
	v280 = v256 + v266<<(uint(int32(4))%32)
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v280)+8))
	if base.Ui64(v190) <= base.Ui64(v281) {
		goto L47
	} else {
		goto L51
	}
L50:
	;
	goto L47
L51:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	if v283 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v295 = v266 + int32(1)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v254)+28))
	if v295 != v297 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v254)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v254)+12)) = v286 - int32(1)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v254)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v254)+16)) = v290 - int32(1)
	goto L52
L56:
	;
	v299 = v295
	goto L58
L57:
	;
	v299 = int32(0)
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254)+24)) = v299
	if v299 != v258 {
		v266 = v299
		goto L49
	} else {
		goto L59
	}
L59:
	;
	goto L50
L60:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v414)+120))
	if v415 != 0 {
		v497 = v414
		goto L76
	} else {
		goto L77
	}
L61:
	;
	v325 = *(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[2]))
	if v325 <= int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v254)+12))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v254)+8))
	if base.Ui32(v329) <= base.Ui32(v328) {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v254)+28))
	v337 = v331
	v338 = v328
	v343 = v258
	goto L64
L64:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v254)+16))
	if base.Ui32(v337-int32(1)) <= base.Ui32(v350+v338) {
		goto L60
	} else {
		goto L66
	}
L65:
	;
	goto L60
L66:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	v362 = m.T0[v361].(func(*base.Module, int32, int32) int32)(m, v355, v256+v343<<(uint(int32(4))%32)+int32(8))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L12
	} else {
		goto L71
	}
L67:
	;
	v387 = v385 + int32(1)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v254)+28))
	if v387 != v389 {
		goto L72
	} else {
		goto L73
	}
L68:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v254)+20))
	v379 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v256+v375<<(uint(int32(4))%32)))) = uint8(v379)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v254)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v254)+16)) = v381 + int32(1)
	v385 = v375
	goto L67
L69:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v254)+20))
	v369 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v256+v365<<(uint(int32(4))%32)))) = uint8(v369)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v254)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v254)+12)) = v371 + v369
	v385 = v365
	goto L67
L70:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v254)+20))
	v385 = v364
	goto L67
L71:
	;
	switch v362 {
	case 0:
		goto L68
	case 1:
		goto L69
	case 2:
		goto L60
	default:
		goto L70
	}
L72:
	;
	v391 = v387
	goto L74
L73:
	;
	v391 = int32(0)
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254)+20)) = v391
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v254)+12))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v254)+8))
	if base.Ui32(v393) < base.Ui32(v394) {
		v337 = v389
		v338 = v393
		v343 = v391
		goto L64
	} else {
		goto L75
	}
L75:
	;
	goto L65
L76:
	;
	v510 = int32(0)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v497)+96))
	if v511 == v510 {
		goto L94
	} else {
		goto L95
	}
L77:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414)+1256)))
	if v416 != 0 {
		v497 = v414
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)+12))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v417)+8))
	if base.Ui32(v419) <= base.Ui32(v418) {
		v497 = v414
		goto L76
	} else {
		goto L79
	}
L79:
	;
	v422 = v417 + int32(32)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v417)+28))
	v429 = v423
	v430 = v418
	goto L80
L80:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v417)+16))
	if base.Ui32(v429-int32(1)) <= base.Ui32(v442+v430) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v497 = v491
	goto L76
L82:
	;
	goto L81
L83:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v417)+4))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v417)+20))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
	v455 = m.T0[v454].(func(*base.Module, int32, int32) int32)(m, v447, v422+v448<<(uint(int32(4))%32)+int32(8))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L12
	} else {
		goto L88
	}
L84:
	;
	v480 = v478 + int32(1)
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v417)+28))
	if v480 != v482 {
		goto L89
	} else {
		goto L90
	}
L85:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v417)+20))
	v472 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v422+v468<<(uint(int32(4))%32)))) = uint8(v472)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v417)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v417)+16)) = v474 + int32(1)
	v478 = v468
	goto L84
L86:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v417)+20))
	v462 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v422+v458<<(uint(int32(4))%32)))) = uint8(v462)
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v417)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v417)+12)) = v464 + v462
	v478 = v458
	goto L84
L87:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v417)+20))
	v478 = v457
	goto L84
L88:
	;
	switch v455 {
	case 0:
		goto L85
	case 1:
		goto L86
	case 2:
		goto L82
	default:
		goto L87
	}
L89:
	;
	v484 = v480
	goto L91
L90:
	;
	v484 = int32(0)
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v417)+20)) = v484
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v417)+12))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v417)+8))
	if base.Ui32(v486) < base.Ui32(v487) {
		v429 = v482
		v430 = v486
		goto L80
	} else {
		goto L92
	}
L92:
	;
	goto L82
L93:
	;
	if v615 == int32(0) {
		v653 = v510
		goto L5
	} else {
		goto L119
	}
L94:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v497)+120))
	if v590 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v497)+96)) = int32(0)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v511)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v497)+120)) = v516
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v497)+124))
	if v511 == v518 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v497)+124)) = int32(0)
	goto L98
L97:
	;
	goto L98
L98:
	;
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511)+4)))
	if v522 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v497)+112)) = v530
	goto L94
L100:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v511)+8))
	if v525 != 0 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	F_pfree(m, v511)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L12
	} else {
		goto L110
	}
L103:
	;
	v530 = v525
	goto L106
L104:
	;
	goto L105
L105:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v497)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v497)+116)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v497)+112)) = v566
	goto L94
L106:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530)+4)))
	if v544 != int32(1) {
		goto L99
	} else {
		goto L108
	}
L107:
	;
	goto L105
L108:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v530)+8))
	if v547 != 0 {
		v530 = v547
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	goto L94
L111:
	;
	v615 = v613
	goto L93
L112:
	;
	v593 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v593
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+1256)))
	if v596 != int32(1) {
		v613 = v593
		goto L111
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v497)+96)) = v590
	v606 = *(*int64)(unsafe.Add(mBase, uint32(v590)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v497)+32)) = v606
	v608 = *(*int64)(unsafe.Add(mBase, uint32(v590)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v497)+40)) = v608
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = int32(0)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v497)+96))
	v613 = v612
	goto L111
L115:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v497)+1252))
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599))))
	if v600 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v599
	goto L118
L117:
	;
	goto L118
L118:
	;
	v602 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v497)+1256)) = uint8(v602)
	v615 = v602
	goto L93
L119:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v618 == v615 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	goto L122
L121:
	;
	goto L122
L122:
	;
	v622 = *(*int64)(unsafe.Add(mBase, uint32(v615)+16))
	v623 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui64(v623) <= base.Ui64(v622) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v629)+124))
	if v630 != 0 {
		goto L127
	} else {
		goto L128
	}
L124:
	;
	goto L125
L125:
	;
	v653 = v615 + int32(32)
	goto L5
L126:
	;
	goto L125
L127:
	;
	v631 = *(*int64)(unsafe.Add(mBase, uint32(v630)+16))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v629)+120))
	v633 = *(*int64)(unsafe.Add(mBase, uint32(v632)+16))
	v636 = base.I32_wrap_i64(v631 - v633)
	goto L129
L128:
	;
	v636 = int32(0)
	goto L129
L129:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v637)+16))
	v640 = *(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[4]))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v637)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v640)+64)) = v641
	*(*int32)(unsafe.Add(mBase, uint32(v640)+56)) = v636
	*(*int32)(unsafe.Add(mBase, uint32(v640)+60)) = v641 + v638
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v647 = *(*int64)(unsafe.Add(mBase, uint32(v646)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v647 - int64(-8192)
	goto L126
L130:
	;
	v657 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReadRecord[5])))
	if v657 != 0 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	goto L132
L132:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v23)+1208))
	v709 = *(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[6]))
	v710 = int32(0)
	if v709 == v710 {
		goto L149
	} else {
		goto L150
	}
L133:
	;
	v668 = *(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[7]))
	if int32(0) <= v668 {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	v658 = *(*int64)(unsafe.Add(mBase, uint32(v23)+48))
	if v658 == int64(0) {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_ReadRecord[8])) = v658
	v664 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int64)(unsafe.Add(mBase, _c_F_ReadRecord[9])) = v664
	goto L133
L136:
	;
	v671 = F_close(m, v668)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[7])) = int32(-1)
	goto L138
L137:
	;
	goto L138
L138:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v21)+140))
	if v675 == int32(0) {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	if l1 != int32(15) {
		v690 = l1
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v693 = F_errstart(m, v690, int32(0))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L12
	} else {
		goto L144
	}
L141:
	;
	v679 = *(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[10]))
	if v679 != int32(2) {
		v690 = l1
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v683 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	v685 = *(*int64)(unsafe.Add(mBase, _c_F_ReadRecord[11]))
	if v683 == v685 {
		v690 = int32(14)
		goto L140
	} else {
		goto L143
	}
L143:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_ReadRecord[11])) = v683
	v690 = int32(15)
	goto L140
L144:
	;
	if v693 == int32(0) {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v21)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v697
	F_errmsg_internal(m, int32(_a_F_ReadRecord_0), v21)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L12
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_ReadRecord_1), int32(3175), int32(_a_F_ReadRecord_2))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L12
	} else {
		goto L147
	}
L147:
	;
	goto L4
L148:
	;
	if v749 != 0 {
		v929 = v653
		goto L3
	} else {
		goto L161
	}
L149:
	;
	v749 = int32(0)
	goto L148
L150:
	;
	goto L151
L151:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v709)+4))
	if v716 <= int32(0) {
		v743 = v710
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v749 = v743
	goto L148
L153:
	;
	v719 = int32(0)
	if v719 < v716 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v722 = v716
	goto L156
L155:
	;
	v722 = v719
	goto L156
L156:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v709)+12))
	v726 = int32(0)
	goto L157
L157:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v723+v726<<(uint(int32(2))%32))))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v733)))
	v735 = base.B2i32(v734 == v707)
	if v734 == v707 {
		v743 = v735
		goto L152
	} else {
		goto L159
	}
L158:
	;
	v743 = v735
	goto L152
L159:
	;
	v737 = v726 + int32(1)
	if v737 != v722 {
		v726 = v737
		goto L157
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	v750 = *(*int64)(unsafe.Add(mBase, uint32(v23)+1200))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v23)+1184))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v751
	v754 = *(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[12]))
	v755 = base.I64_extend_i32_s(v754)
	v756 = base.I64_div_u_s(v750, v755)
	v758 = base.I64_div_u_s(int64(4294967296), v755)
	v759 = base.I64_div_u_s(v756, v758)
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+52)) = uint32(v759)
	v762 = v756 - v758*v759
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+56)) = uint32(v762)
	v770 = F_pg_snprintf(m, v21-int32(-64), int32(64), int32(_a_F_ReadRecord_3), v21+int32(48))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L12
	} else {
		goto L162
	}
L162:
	;
	if l1 != int32(15) {
		v784 = l1
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v787 = F_errstart(m, v784, int32(0))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L12
	} else {
		goto L167
	}
L164:
	;
	v773 = *(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[10]))
	if v773 != int32(2) {
		v784 = l1
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v777 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	v779 = *(*int64)(unsafe.Add(mBase, _c_F_ReadRecord[11]))
	if v777 == v779 {
		v784 = int32(14)
		goto L163
	} else {
		goto L166
	}
L166:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_ReadRecord[11])) = v777
	v784 = int32(15)
	goto L163
L167:
	;
	if v787 == int32(0) {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v23)+1208))
	v792 = *(*int64)(unsafe.Add(mBase, uint32(v23)+1200))
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(32)))) = base.I32_wrap_i64(v750) & (v754 - int32(1))
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+28)) = uint32(v792)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v791
	v801 = int64(base.Ui64(v792) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+24)) = uint32(v801)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v21 - int32(-64)
	F_errmsg(m, int32(_a_F_ReadRecord_4), v21+int32(16))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L12
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_ReadRecord_1), int32(3197), int32(_a_F_ReadRecord_2))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L12
	} else {
		goto L170
	}
L170:
	;
	goto L4
L171:
	;
	v920 = int32(0)
	v922 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReadRecord[13])))
	if v922 != int32(1) {
		v929 = v920
		goto L3
	} else {
		goto L196
	}
L172:
	;
	v826 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReadRecord[14])))
	if v826&int32(1) != 0 {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v830 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReadRecord[5])))
	if v830&int32(1) == int32(0) {
		goto L171
	} else {
		goto L174
	}
L174:
	;
	v837 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L12
	} else {
		goto L175
	}
L175:
	;
	if v837 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	F_errmsg_internal(m, int32(_a_F_ReadRecord_5), int32(0))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L12
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v849 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ReadRecord[14])) = uint8(v849)
	v852 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReadRecord[15])))
	if v852 != 0 {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	F_errfinish(m, int32(_a_F_ReadRecord_1), int32(3227), int32(_a_F_ReadRecord_2))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L12
	} else {
		goto L180
	}
L180:
	;
	goto L178
L181:
	;
	v854 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ReadRecord[13])) = uint8(v854)
	F_disable_startup_progress_timeout(m)
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L12
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v858 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	v860 = *(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[16]))
	v864 = F_LWLockAcquire(m, v860+int32(1152), int32(0))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L12
	} else {
		goto L185
	}
L184:
	;
	goto L183
L185:
	;
	v867 = *(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[17]))
	*(*int32)(unsafe.Add(mBase, uint32(v867)+16)) = int32(5)
	v871 = *(*int64)(unsafe.Add(mBase, uint32(v867)+144))
	if base.Ui64(v871) < base.Ui64(v858) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v867)+152)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v867)+144)) = v858
	v875 = v858
	goto L188
L187:
	;
	v875 = v871
	goto L188
L188:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_ReadRecord[18])) = v875
	v878 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_ReadRecord[19])) = uint8(v878)
	v881 = *(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[20]))
	F_update_controlfile(m, v881, v867)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L12
	} else {
		goto L189
	}
L189:
	;
	v885 = *(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[21]))
	v888 = base.AtomicRmwXchg32(m, v885, int32(440), int32(1))
	if v888 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	F_s_lock(m, v885+int32(440), int32(_a_F_ReadRecord_6))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L12
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v895 = *(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v895)+308)) = int32(1)
	v898 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v895)+440)), uint32(v898))
	v902 = *(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[16]))
	F_LWLockRelease(m, v902+int32(1152))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L12
	} else {
		goto L194
	}
L193:
	;
	goto L192
L194:
	;
	v907 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[22])) = l3
	*(*int64)(unsafe.Add(mBase, _c_F_ReadRecord[23])) = v907
	F_CheckRecoveryConsistency(m)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L12
	} else {
		goto L195
	}
L195:
	;
	v915 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ReadRecord[24])) = v915
	*(*uint8)(unsafe.Add(mBase, _c_F_ReadRecord[0])) = uint8(v915)
	goto L1
L196:
	;
	v925 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L12
	} else {
		goto L197
	}
L197:
	;
	if v925 == int32(0) {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v929 = v920
	goto L3
}
func F_RegisterSnapshotOnOwner(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	if l0 == int32(0) {
		return int32(0)
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
		if v11 != 0 {
			v91 = l0
			F_ResourceOwnerEnlarge(m, l1)
			mBase = m.M
			v96 = m.ExcPending
			if v96 != 0 {
				return int32(0)
			} else {
				v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v91)+48)) = v97 + int32(1)
				F_ResourceOwnerRemember(m, l1, base.I64_extend_i32_u(v91), int32(_a_F_RegisterSnapshotOnOwner_0))
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return int32(0)
				} else {
					v105 = *(*int32)(unsafe.Add(mBase, uint32(v91)+48))
					if v105 == int32(1) {
						F_pairingheap_add(m, int32(_a_F_RegisterSnapshotOnOwner_1), v91+int32(52))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							return v91
						}
					} else {
						return v91
					}
				}
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterSnapshotOnOwner[0]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v17 = int32(2)
			v19 = int32(72)
			v24 = v15<<(uint(v17)%32) + v19
			if int32(0) < v14 {
				v27 = (v14+v15)<<(uint(v17)%32) + v19
			} else {
				v27 = v24
			}
			v28 = F_MemoryContextAlloc(m, v13, v27)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int64)(unsafe.Add(mBase, uint32(v28)+48)) = v32
				v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
				*(*int64)(unsafe.Add(mBase, uint32(v28)+40)) = v34
				v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v28)+24)) = v36
				v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
				*(*int64)(unsafe.Add(mBase, uint32(v28)+56)) = v38
				v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v28)+32)) = v40
				v42 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v28)+16)) = v42
				v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v44
				v46 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(v28))) = v46
				*(*int64)(unsafe.Add(mBase, uint32(v28)+64)) = int64(0)
				v50 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v50
				*(*int32)(unsafe.Add(mBase, uint32(v28)+44)) = v50
				v54 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v28)+30)) = uint8(v54)
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v56 != 0 {
					v58 = v28 + int32(72)
					*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v58
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v62 = v60 << (uint(int32(2)) % 32)
					if v62 == int32(0) {
					} else {
						v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						base.MemoryCopy(m, v58, v65, v62)
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = int32(0)
				}
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v71 <= int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = int32(0)
					v91 = v28
				} else {
					v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
					if v74 == int32(1) {
						v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
						if v77 != int32(1) {
							*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = int32(0)
							v91 = v28
						} else {
							v80 = v28 + v24
							*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v80
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v84 = v82 << (uint(int32(2)) % 32)
							if v84 == int32(0) {
								v91 = v28
							} else {
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								base.MemoryCopy(m, v80, v87, v84)
								v91 = v28
							}
						}
					} else {
						v80 = v28 + v24
						*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v80
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v84 = v82 << (uint(int32(2)) % 32)
						if v84 == int32(0) {
							v91 = v28
						} else {
							v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							base.MemoryCopy(m, v80, v87, v84)
							v91 = v28
						}
					}
				}
				F_ResourceOwnerEnlarge(m, l1)
				mBase = m.M
				v96 = m.ExcPending
				if v96 != 0 {
					return int32(0)
				} else {
					v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v91)+48)) = v97 + int32(1)
					F_ResourceOwnerRemember(m, l1, base.I64_extend_i32_u(v91), int32(_a_F_RegisterSnapshotOnOwner_0))
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return int32(0)
					} else {
						v105 = *(*int32)(unsafe.Add(mBase, uint32(v91)+48))
						if v105 == int32(1) {
							F_pairingheap_add(m, int32(_a_F_RegisterSnapshotOnOwner_1), v91+int32(52))
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								return v91
							}
						} else {
							return v91
						}
					}
				}
			}
		}
	}
}
func F_RelidByRelfilenumber(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
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
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v386 int32
	_ = v386
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(144)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_RelidByRelfilenumber[0]))
	if v13 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_RelidByRelfilenumber[1]))
	if v17 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v75 = v13
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+140)) = l1
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_RelidByRelfilenumber[2]))
	if l0 != v79 {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v26 = v17
	goto L6
L6:
	;
	base.MemoryFill(m, int32(_a_F_RelidByRelfilenumber_0), int32(0), int32(112))
	F_fmgr_info_cxt(m, int32(184), int32(_a_F_RelidByRelfilenumber_1), v26)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_RelidByRelfilenumber[1]))
	v26 = v25
	goto L6
L9:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_RelidByRelfilenumber[3])) = int64(0)
	v39 = int32(3)
	*(*uint16)(unsafe.Add(mBase, _c_F_RelidByRelfilenumber[4])) = uint16(v39)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_RelidByRelfilenumber[1]))
	F_fmgr_info_cxt(m, int32(184), int32(_a_F_RelidByRelfilenumber_2), v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_RelidByRelfilenumber[5])) = int64(196616)
	v51 = int32(9)
	*(*uint16)(unsafe.Add(mBase, _c_F_RelidByRelfilenumber[6])) = uint16(v51)
	*(*int32)(unsafe.Add(mBase, _c_F_RelidByRelfilenumber[7])) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(51539607560)
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_RelidByRelfilenumber[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v59
	v67 = F_hash_create(m, int32(_a_F_RelidByRelfilenumber_3), int64(64), v10+int32(16), int32(1064))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelidByRelfilenumber[0])) = v67
	F_CacheRegisterRelcacheCallback(m, int32(1808))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_RelidByRelfilenumber[0]))
	v75 = v74
	goto L3
L13:
	;
	v81 = l0
	goto L15
L14:
	;
	v81 = int32(0)
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+136)) = v81
	v88 = F_hash_search(m, v75, v10+int32(136), int32(0), v10+int32(135))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+135)))
	if v90 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L7
	} else {
		goto L125
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L7
	} else {
		goto L122
	}
L19:
	;
	m.G0 = v10 + int32(144)
	return v386
L20:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v386 = v93
	goto L19
L21:
	;
	goto L22
L22:
	;
	if v81 == int32(1664) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_RelidByRelfilenumber[0]))
	v376 = F_hash_search(m, v370, v10+int32(136), int32(1), v10+int32(135))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L7
	} else {
		goto L120
	}
L24:
	;
	v97 = int32(0)
	goto L31
L25:
	;
	goto L26
L26:
	;
	v202 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L7
	} else {
		goto L65
	}
L27:
	;
	v366 = v199
	goto L23
L28:
	;
	goto L27
L29:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v199 = v194
	goto L28
L31:
	;
	goto L32
L32:
	;
	v145 = int32(0)
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_RelidByRelfilenumber[8]))
	if v145 < v147 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v175 = int32(0)
	goto L59
L50:
	;
	v193 = v156 + int32(_a_F_RelidByRelfilenumber_4)
	goto L29
L51:
	;
	v151 = v145
	goto L54
L52:
	;
	goto L53
L53:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_RelidByRelfilenumber[9]))
	if v168 <= int32(0) {
		v199 = v97
		goto L28
	} else {
		goto L58
	}
L54:
	;
	v156 = v151 << (uint(int32(3)) % 32)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+uint32(_c_F_RelidByRelfilenumber[10])))
	if l1 == v157 {
		goto L50
	} else {
		goto L56
	}
L55:
	;
	goto L53
L56:
	;
	v160 = v151 + int32(1)
	if v160 != v147 {
		v151 = v160
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	goto L49
L59:
	;
	v180 = v175 << (uint(int32(3)) % 32)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+uint32(_c_F_RelidByRelfilenumber[11])))
	if v181 != l1 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v193 = v180 + int32(_a_F_RelidByRelfilenumber_5)
	goto L29
L61:
	;
	v184 = v175 + int32(1)
	if v168 != v184 {
		v175 = v184
		goto L59
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	v199 = v97
	goto L28
L65:
	;
	v205 = v10 + int32(16)
	base.MemoryCopy(m, v205, int32(_a_F_RelidByRelfilenumber_0), int32(112))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+120)) = base.I64_extend_i32_u(l1)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = base.I64_extend_i32_u(v81)
	v217 = F_systable_beginscan(m, v202, int32(3455), int32(1), int32(0), int32(2), v205)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	v219 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+135)) = uint8(v219)
	v221 = F_systable_getnext(m, v217)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	if v221 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v223 = v221
	v227 = v3
	goto L71
L69:
	;
	v249 = v3
	goto L70
L70:
	;
	F_systable_endscan(m, v217)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L7
	} else {
		goto L79
	}
L71:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+22)))
	v232 = v230 + v231
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+118)))
	if v233 != int32(116) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v249 = v242
	goto L70
L73:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+135)))
	if v236 == int32(1) {
		goto L18
	} else {
		goto L76
	}
L74:
	;
	v242 = v227
	goto L75
L75:
	;
	v243 = F_systable_getnext(m, v217)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L7
	} else {
		goto L77
	}
L76:
	;
	v239 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+135)) = uint8(v239)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	v242 = v241
	goto L75
L77:
	;
	if v243 != 0 {
		v223 = v243
		v227 = v242
		goto L71
	} else {
		goto L78
	}
L78:
	;
	goto L72
L79:
	;
	F_relation_close(m, v202, int32(1))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L7
	} else {
		goto L80
	}
L80:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+135)))
	if v257 != 0 {
		v366 = v249
		goto L23
	} else {
		goto L81
	}
L81:
	;
	v258 = int32(0)
	goto L85
L82:
	;
	v366 = v361
	goto L23
L83:
	;
	goto L82
L84:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v361 = v356
	goto L83
L85:
	;
	v264 = int32(0)
	v266 = *(*int32)(unsafe.Add(mBase, _c_F_RelidByRelfilenumber[12]))
	if v264 < v266 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v270 = v264
	goto L91
L89:
	;
	goto L90
L90:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_RelidByRelfilenumber[13]))
	if v289 <= int32(0) {
		v361 = v258
		goto L83
	} else {
		goto L97
	}
L91:
	;
	v275 = v270 << (uint(int32(3)) % 32)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+uint32(_c_F_RelidByRelfilenumber[14])))
	if v276 == l1 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	goto L90
L93:
	;
	v355 = v275 + int32(_a_F_RelidByRelfilenumber_6)
	goto L84
L94:
	;
	goto L95
L95:
	;
	v281 = v270 + int32(1)
	if v281 != v266 {
		v270 = v281
		goto L91
	} else {
		goto L96
	}
L96:
	;
	goto L92
L97:
	;
	v294 = int32(0)
	goto L98
L98:
	;
	v299 = v294 << (uint(int32(3)) % 32)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+uint32(_c_F_RelidByRelfilenumber[15])))
	if v300 != l1 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v355 = v299 + int32(_a_F_RelidByRelfilenumber_7)
	goto L84
L100:
	;
	v303 = v294 + int32(1)
	if v289 != v303 {
		v294 = v303
		goto L98
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	goto L99
L103:
	;
	v361 = v258
	goto L83
L120:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+135)))
	if v378 == int32(1) {
		goto L17
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v376)+8)) = v366
	v386 = v366
	goto L19
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v81
	F_errmsg_internal(m, int32(_a_F_RelidByRelfilenumber_8), v10)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L7
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_RelidByRelfilenumber_9), int32(222), int32(_a_F_RelidByRelfilenumber_10))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L7
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_errmsg_internal(m, int32(_a_F_RelidByRelfilenumber_11), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L7
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_RelidByRelfilenumber_9), int32(245), int32(_a_F_RelidByRelfilenumber_10))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L7
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RememberIndexForRebuilding(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	if v9 == v3 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L18
	} else {
		goto L40
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L18
	} else {
		goto L37
	}
L3:
	;
	m.G0 = v7 + int32(32)
	return
L4:
	;
	if v48 != 0 {
		goto L3
	} else {
		goto L17
	}
L5:
	;
	v48 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v16 <= int32(0) {
		v42 = v3
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v48 = v42
	goto L4
L9:
	;
	v19 = int32(0)
	if v19 < v16 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v22 = v16
	goto L12
L11:
	;
	v22 = v19
	goto L12
L12:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v25 = int32(0)
	goto L13
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v23+v25<<(uint(int32(2))%32))))
	v34 = base.B2i32(v33 == l0)
	if v33 == l0 {
		v42 = v34
		goto L8
	} else {
		goto L15
	}
L14:
	;
	v42 = v34
	goto L8
L15:
	;
	v36 = v25 + int32(1)
	if v36 != v22 {
		v25 = v36
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v49 = F_get_index_constraint(m, l0)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	if v49 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_RememberConstraintForRebuilding(m, v49, l1)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v53 = int32(0)
	v57 = int32(1)
	v61 = F_pg_get_indexdef_worker(m, l0, v53, v53, v53, v53, v57, v57, v53, v53)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L18
	} else {
		goto L24
	}
L23:
	;
	goto L3
L24:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	v64 = F_lappend_oid(m, v63, l0)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+116)) = v64
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+120))
	v68 = F_lappend(m, v67, v61)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+120)) = v68
	v71 = F_get_index_isreplident(m, l0)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	if v71 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	if v73 != 0 {
		goto L2
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v77 = F_get_index_isclustered(m, l0)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L18
	} else {
		goto L33
	}
L31:
	;
	v74 = F_get_rel_name(m, l0)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L18
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+124)) = v74
	goto L30
L33:
	;
	if v77 == int32(0) {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	if v81 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v82 = F_get_rel_name(m, l0)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L18
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+128)) = v82
	goto L3
L37:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v93
	F_errmsg_internal(m, int32(_a_F_RememberIndexForRebuilding_0), v7+int32(16))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L18
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_RememberIndexForRebuilding_1), int32(_a_F_RememberIndexForRebuilding_2), int32(_a_F_RememberIndexForRebuilding_3))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L18
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
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v109
	F_errmsg_internal(m, int32(_a_F_RememberIndexForRebuilding_4), v7)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L18
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_RememberIndexForRebuilding_1), int32(_a_F_RememberIndexForRebuilding_5), int32(_a_F_RememberIndexForRebuilding_6))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L18
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RememberSyncRequest(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
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
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	switch l1 - int32(1) {
	case 0:
		goto L3
	case 1:
		goto L5
	case 2:
		goto L4
	default:
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(32)
	return
L2:
	;
	v128 = int32(_a_F_RememberSyncRequest_0)
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_RememberSyncRequest[0]))
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_RememberSyncRequest[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RememberSyncRequest[0])) = v132
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_RememberSyncRequest[2]))
	v139 = F_hash_search(m, v135, l0, int32(1), v9+int32(12))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L6
	} else {
		goto L33
	}
L3:
	;
	v100 = int32(_a_F_RememberSyncRequest_0)
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_RememberSyncRequest[0]))
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_RememberSyncRequest[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RememberSyncRequest[0])) = v104
	v107 = F_palloc(m, int32(32))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L6
	} else {
		goto L31
	}
L4:
	;
	v24 = v9 + int32(12)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_RememberSyncRequest[2]))
	F_hash_seq_init(m, v24, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L9
	}
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_RememberSyncRequest[2]))
	v15 = int32(0)
	v17 = F_hash_search(m, v14, l0, v15, v15)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	if v17 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v21 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+26)) = uint8(v21)
	goto L1
L9:
	;
	v29 = F_hash_seq_search(m, v24)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	if v29 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = v29
	goto L14
L12:
	;
	goto L13
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_RememberSyncRequest[3]))
	if v63 == int32(0) {
		goto L1
	} else {
		goto L22
	}
L14:
	;
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32))))
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v37 != v38 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	v54 = F_hash_seq_search(m, v9+int32(12))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L20
	}
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_extend16_s(v37)*int32(12))+uint32(_c_F_RememberSyncRequest[4])))
	v46 = m.T0[v45].(func(*base.Module, int32, int32) int32)(m, l0, v32)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	if v46 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v50 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+26)) = uint8(v50)
	goto L16
L20:
	;
	if v54 != 0 {
		v32 = v54
		goto L14
	} else {
		goto L21
	}
L21:
	;
	goto L15
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v66 <= int32(0) {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v71 = int32(0)
	goto L24
L24:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76+v71<<(uint(int32(2))%32))))
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80))))
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v81 != v82 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L1
L26:
	;
	v97 = v71 + int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v97 < v98 {
		v71 = v97
		goto L24
	} else {
		goto L30
	}
L27:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_extend16_s(v81)*int32(12))+uint32(_c_F_RememberSyncRequest[4])))
	v90 = m.T0[v89].(func(*base.Module, int32, int32) int32)(m, l0, v80)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	if v90 == int32(0) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v94 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v80)+26)) = uint8(v94)
	goto L26
L30:
	;
	goto L25
L31:
	;
	v109 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v107)+16)) = v109
	v111 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v107)+8)) = v111
	v113 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v107))) = v113
	v116 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_RememberSyncRequest[5])))
	v117 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v107)+26)) = uint8(v117)
	*(*uint16)(unsafe.Add(mBase, uint32(v107)+24)) = uint16(v116)
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_RememberSyncRequest[3]))
	v122 = F_lappend(m, v121, v107)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RememberSyncRequest[0])) = v101
	*(*int32)(unsafe.Add(mBase, _c_F_RememberSyncRequest[3])) = v122
	goto L1
L33:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+12)))
	if v141 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RememberSyncRequest[0])) = v129
	goto L1
L35:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+26)))
	if v144 != int32(1) {
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v148 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_RememberSyncRequest[6])))
	v149 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v139)+26)) = uint8(v149)
	*(*uint16)(unsafe.Add(mBase, uint32(v139)+24)) = uint16(v148)
	goto L34
L38:
	;
	goto L37
}
func F_RemoveStatistics(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	v5 = m.G0
	v7 = v5 - int32(112)
	m.G0 = v7
	v11 = F_table_open(m, int32(2619), int32(3))
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_ScanKeyInit(m, v7, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l0))
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l1 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = int32(2)
	F_ScanKeyInit(m, v7+int32(56), v20, int32(3), int32(63), base.I64_extend_i32_s(l1))
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v29 = int32(1)
	goto L6
L6:
	;
	v33 = F_systable_beginscan(m, v11, int32(2696), int32(1), int32(0), v29, v7)
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v29 = v20
	goto L6
L8:
	;
	v35 = F_systable_getnext(m, v33)
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v35 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = v35
	goto L13
L11:
	;
	goto L12
L12:
	;
	F_systable_endscan(m, v33)
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	F_simple_heap_delete(m, v11, v38+int32(4))
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v45 = F_systable_getnext(m, v33)
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v45 != 0 {
		v38 = v45
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	F_relation_close(m, v11, int32(3))
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	m.G0 = v7 + int32(112)
	return
}
func F_RenameConstraintById(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v14 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v19 = F_SearchSysCacheCopy(m, int32(19), base.I64_extend_i32_u(l0), int64(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			if v19 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
				v23 = v21 + v22
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
				if v24 != 0 {
					v26 = F_ConstraintNameIsUsed(m, int32(0), v24, l1)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						if v26 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								F_errcode(m, int32(_a_F_RenameConstraintById_0))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
									v79 = F_get_rel_name(m, v78)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v79
										*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l1
										F_errmsg(m, int32(_a_F_RenameConstraintById_1), v10+int32(32))
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_RenameConstraintById_2), int32(1031), int32(_a_F_RenameConstraintById_3))
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
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
							if v28 != 0 {
								v30 = F_ConstraintNameIsUsed(m, int32(1), v28, l1)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									if v30 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return
										} else {
											F_errcode(m, int32(_a_F_RenameConstraintById_0))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
												v101 = F_format_type_be(m, v100)
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v101
													*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
													F_errmsg(m, int32(_a_F_RenameConstraintById_4), v10+int32(16))
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_RenameConstraintById_2), int32(1039), int32(_a_F_RenameConstraintById_3))
														mBase = m.M
														v114 = m.ExcPending
														if v114 != 0 {
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
										v35 = F_strncpy(m, v23+int32(4), l1, int32(64))
										mBase = m.M
										v36 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v35)+63)) = uint8(v36)
										F_CatalogTupleUpdate(m, v14, v19+int32(4), v19)
										mBase = m.M
										v41 = m.ExcPending
										if v41 != 0 {
											return
										} else {
											v43 = *(*int32)(unsafe.Add(mBase, _c_F_RenameConstraintById[0]))
											if v43 != 0 {
												v45 = int32(0)
												F_RunObjectPostAlterHook(m, int32(2606), l0, v45, v45, v45)
												mBase = m.M
												v49 = m.ExcPending
												if v49 != 0 {
													return
												} else {
													F_pfree(m, v19)
													mBase = m.M
													v51 = m.ExcPending
													if v51 != 0 {
														return
													} else {
														F_relation_close(m, v14, int32(3))
														mBase = m.M
														v54 = m.ExcPending
														if v54 != 0 {
															return
														} else {
															m.G0 = v10 + int32(48)
															return
														}
													}
												}
											} else {
												F_pfree(m, v19)
												mBase = m.M
												v51 = m.ExcPending
												if v51 != 0 {
													return
												} else {
													F_relation_close(m, v14, int32(3))
													mBase = m.M
													v54 = m.ExcPending
													if v54 != 0 {
														return
													} else {
														m.G0 = v10 + int32(48)
														return
													}
												}
											}
										}
									}
								}
							} else {
								v35 = F_strncpy(m, v23+int32(4), l1, int32(64))
								mBase = m.M
								v36 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v35)+63)) = uint8(v36)
								F_CatalogTupleUpdate(m, v14, v19+int32(4), v19)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, _c_F_RenameConstraintById[0]))
									if v43 != 0 {
										v45 = int32(0)
										F_RunObjectPostAlterHook(m, int32(2606), l0, v45, v45, v45)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return
										} else {
											F_pfree(m, v19)
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
												return
											} else {
												F_relation_close(m, v14, int32(3))
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return
												} else {
													m.G0 = v10 + int32(48)
													return
												}
											}
										}
									} else {
										F_pfree(m, v19)
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return
										} else {
											F_relation_close(m, v14, int32(3))
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return
											} else {
												m.G0 = v10 + int32(48)
												return
											}
										}
									}
								}
							}
						}
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
					if v28 != 0 {
						v30 = F_ConstraintNameIsUsed(m, int32(1), v28, l1)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							if v30 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return
								} else {
									F_errcode(m, int32(_a_F_RenameConstraintById_0))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
										v101 = F_format_type_be(m, v100)
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v101
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
											F_errmsg(m, int32(_a_F_RenameConstraintById_4), v10+int32(16))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_RenameConstraintById_2), int32(1039), int32(_a_F_RenameConstraintById_3))
												mBase = m.M
												v114 = m.ExcPending
												if v114 != 0 {
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
								v35 = F_strncpy(m, v23+int32(4), l1, int32(64))
								mBase = m.M
								v36 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v35)+63)) = uint8(v36)
								F_CatalogTupleUpdate(m, v14, v19+int32(4), v19)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, _c_F_RenameConstraintById[0]))
									if v43 != 0 {
										v45 = int32(0)
										F_RunObjectPostAlterHook(m, int32(2606), l0, v45, v45, v45)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return
										} else {
											F_pfree(m, v19)
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
												return
											} else {
												F_relation_close(m, v14, int32(3))
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return
												} else {
													m.G0 = v10 + int32(48)
													return
												}
											}
										}
									} else {
										F_pfree(m, v19)
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return
										} else {
											F_relation_close(m, v14, int32(3))
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return
											} else {
												m.G0 = v10 + int32(48)
												return
											}
										}
									}
								}
							}
						}
					} else {
						v35 = F_strncpy(m, v23+int32(4), l1, int32(64))
						mBase = m.M
						v36 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v35)+63)) = uint8(v36)
						F_CatalogTupleUpdate(m, v14, v19+int32(4), v19)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, _c_F_RenameConstraintById[0]))
							if v43 != 0 {
								v45 = int32(0)
								F_RunObjectPostAlterHook(m, int32(2606), l0, v45, v45, v45)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									F_pfree(m, v19)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										F_relation_close(m, v14, int32(3))
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return
										} else {
											m.G0 = v10 + int32(48)
											return
										}
									}
								}
							} else {
								F_pfree(m, v19)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									F_relation_close(m, v14, int32(3))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										m.G0 = v10 + int32(48)
										return
									}
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
					F_errmsg_internal(m, int32(_a_F_RenameConstraintById_5), v10)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_RenameConstraintById_2), int32(1018), int32(_a_F_RenameConstraintById_3))
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
		}
	}
}
func F_RenameRelationInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
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
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	if l3 != 0 {
		v17 = int32(4)
	} else {
		v17 = int32(8)
	}
	v18 = F_relation_open(m, l0, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
		v24 = F_table_open(m, int32(1259), int32(3))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			v28 = F_SearchSysCacheLockedCopy1(m, int32(57), base.I64_extend_i32_u(l0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				if v28 != 0 {
					v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+8)))
					*(*uint16)(unsafe.Add(mBase, uint32(v13)+28)) = uint16(v30)
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
					v36 = F_get_relname_relid(m, l1, v21)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						if v36 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return
							} else {
								F_errcode(m, int32(117571716))
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
									F_errmsg(m, int32(_a_F_RenameRelationInternal_0), v13+int32(16))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_RenameRelationInternal_1), int32(_a_F_RenameRelationInternal_2), int32(_a_F_RenameRelationInternal_3))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
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
							v42 = F_strncpy(m, v34+v35+int32(4), l1, int32(64))
							mBase = m.M
							v43 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v42)+63)) = uint8(v43)
							v46 = v13 + int32(24)
							F_CatalogTupleUpdate(m, v24, v46, v28)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								F_UnlockTuple(m, v24, v46, int32(7))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, _c_F_RenameRelationInternal[0]))
									if v53 != 0 {
										v55 = int32(0)
										F_RunObjectPostAlterHook(m, int32(1259), l0, v55, v55, l2)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return
										} else {
											F_pfree(m, v28)
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return
											} else {
												F_relation_close(m, v24, int32(3))
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return
												} else {
													v64 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
													v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
													if v65 != 0 {
														F_RenameTypeInternal(m, v65, l1, v21)
														mBase = m.M
														v67 = m.ExcPending
														if v67 != 0 {
															return
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
															v69 = v68
															v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+119)))
															if v70|int32(32) != int32(105) {
																F_relation_close(m, v18, int32(0))
																mBase = m.M
																v84 = m.ExcPending
																if v84 != 0 {
																	return
																} else {
																	m.G0 = v13 + int32(32)
																	return
																}
															} else {
																v75 = F_get_index_constraint(m, l0)
																mBase = m.M
																v76 = m.ExcPending
																if v76 != 0 {
																	return
																} else {
																	if v75 == int32(0) {
																		F_relation_close(m, v18, int32(0))
																		mBase = m.M
																		v84 = m.ExcPending
																		if v84 != 0 {
																			return
																		} else {
																			m.G0 = v13 + int32(32)
																			return
																		}
																	} else {
																		F_RenameConstraintById(m, v75, l1)
																		mBase = m.M
																		v80 = m.ExcPending
																		if v80 != 0 {
																			return
																		} else {
																			F_relation_close(m, v18, int32(0))
																			mBase = m.M
																			v84 = m.ExcPending
																			if v84 != 0 {
																				return
																			} else {
																				m.G0 = v13 + int32(32)
																				return
																			}
																		}
																	}
																}
															}
														}
													} else {
														v69 = v64
														v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+119)))
														if v70|int32(32) != int32(105) {
															F_relation_close(m, v18, int32(0))
															mBase = m.M
															v84 = m.ExcPending
															if v84 != 0 {
																return
															} else {
																m.G0 = v13 + int32(32)
																return
															}
														} else {
															v75 = F_get_index_constraint(m, l0)
															mBase = m.M
															v76 = m.ExcPending
															if v76 != 0 {
																return
															} else {
																if v75 == int32(0) {
																	F_relation_close(m, v18, int32(0))
																	mBase = m.M
																	v84 = m.ExcPending
																	if v84 != 0 {
																		return
																	} else {
																		m.G0 = v13 + int32(32)
																		return
																	}
																} else {
																	F_RenameConstraintById(m, v75, l1)
																	mBase = m.M
																	v80 = m.ExcPending
																	if v80 != 0 {
																		return
																	} else {
																		F_relation_close(m, v18, int32(0))
																		mBase = m.M
																		v84 = m.ExcPending
																		if v84 != 0 {
																			return
																		} else {
																			m.G0 = v13 + int32(32)
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
									} else {
										F_pfree(m, v28)
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return
										} else {
											F_relation_close(m, v24, int32(3))
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return
											} else {
												v64 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
												v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
												if v65 != 0 {
													F_RenameTypeInternal(m, v65, l1, v21)
													mBase = m.M
													v67 = m.ExcPending
													if v67 != 0 {
														return
													} else {
														v68 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
														v69 = v68
														v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+119)))
														if v70|int32(32) != int32(105) {
															F_relation_close(m, v18, int32(0))
															mBase = m.M
															v84 = m.ExcPending
															if v84 != 0 {
																return
															} else {
																m.G0 = v13 + int32(32)
																return
															}
														} else {
															v75 = F_get_index_constraint(m, l0)
															mBase = m.M
															v76 = m.ExcPending
															if v76 != 0 {
																return
															} else {
																if v75 == int32(0) {
																	F_relation_close(m, v18, int32(0))
																	mBase = m.M
																	v84 = m.ExcPending
																	if v84 != 0 {
																		return
																	} else {
																		m.G0 = v13 + int32(32)
																		return
																	}
																} else {
																	F_RenameConstraintById(m, v75, l1)
																	mBase = m.M
																	v80 = m.ExcPending
																	if v80 != 0 {
																		return
																	} else {
																		F_relation_close(m, v18, int32(0))
																		mBase = m.M
																		v84 = m.ExcPending
																		if v84 != 0 {
																			return
																		} else {
																			m.G0 = v13 + int32(32)
																			return
																		}
																	}
																}
															}
														}
													}
												} else {
													v69 = v64
													v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+119)))
													if v70|int32(32) != int32(105) {
														F_relation_close(m, v18, int32(0))
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return
														} else {
															m.G0 = v13 + int32(32)
															return
														}
													} else {
														v75 = F_get_index_constraint(m, l0)
														mBase = m.M
														v76 = m.ExcPending
														if v76 != 0 {
															return
														} else {
															if v75 == int32(0) {
																F_relation_close(m, v18, int32(0))
																mBase = m.M
																v84 = m.ExcPending
																if v84 != 0 {
																	return
																} else {
																	m.G0 = v13 + int32(32)
																	return
																}
															} else {
																F_RenameConstraintById(m, v75, l1)
																mBase = m.M
																v80 = m.ExcPending
																if v80 != 0 {
																	return
																} else {
																	F_relation_close(m, v18, int32(0))
																	mBase = m.M
																	v84 = m.ExcPending
																	if v84 != 0 {
																		return
																	} else {
																		m.G0 = v13 + int32(32)
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
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
						F_errmsg_internal(m, int32(_a_F_RenameRelationInternal_4), v13)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_RenameRelationInternal_1), int32(_a_F_RenameRelationInternal_5), int32(_a_F_RenameRelationInternal_3))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
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
func F_RepackWorkerMain(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
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
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int64
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int64
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int64
	_ = v194
	var v196 int64
	_ = v196
	var v197 int64
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
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
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int64
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int64
	_ = v364
	var v365 int32
	_ = v365
	var v368 int64
	_ = v368
	var v375 int32
	_ = v375
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int64
	_ = v396
	var v398 int64
	_ = v398
	var v399 int64
	_ = v399
	var v401 int64
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int64
	_ = v413
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int64
	_ = v447
	var v448 int32
	_ = v448
	var v455 int64
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int64
	_ = v458
	var v461 int64
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int64
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
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
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int64
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v582 int64
	_ = v582
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	v14 = m.G0
	v16 = v14 - int32(1152)
	m.G0 = v16
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_RepackWorkerMain[0])) = uint8(v19)
	F_BackgroundWorkerUnblockSignals(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v24 = F_dsm_attach(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L147
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L143
	}
L5:
	;
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	F_before_shmem_exit(m, int32(613), base.I64_extend_i32_u(v24))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L139
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_RepackWorkerMain[1]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v26)+116))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)+112))
	v39 = F_BecomeLockGroupMember(m, v33+v34*int32(768), v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v39 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v44 = (v26 + int32(151)) & int32(-32)
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_RepackWorkerMain[2]))
	F_shm_mq_set_sender(m, v44, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	m.G0 = v16 + int32(1152)
	return
L14:
	;
	v49 = F_shm_mq_attach(m, v44, v24)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_pq_redirect_to_shm_mq(m, v24, v49)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v26)+112))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v26)+116))
	*(*int32)(unsafe.Add(mBase, _c_F_RepackWorkerMain[3])) = v54
	*(*int32)(unsafe.Add(mBase, _c_F_RepackWorkerMain[4])) = v53
	goto L17
L17:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v26)+80))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v26)+84))
	F_BackgroundWorkerInitializeConnectionByOid(m, v59, v60, int32(3))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v26)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v64
	v67 = v16 + int32(80)
	v72 = F_pg_snprintf(m, v67, int32(32), int32(_a_F_RepackWorkerMain_0), v16-int32(-64))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_SetConfigOption(m, int32(_a_F_RepackWorkerMain_1), v67, int32(5), int32(10))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v26)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v79
	v85 = F_pg_snprintf(m, v67, int32(32), int32(_a_F_RepackWorkerMain_0), v16+int32(48))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_SetConfigOption(m, int32(_a_F_RepackWorkerMain_2), v67, int32(5), int32(10))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v95 = int32(76)
	v96 = v94 + v95
	v99 = base.AtomicRmwXchg32(m, v94, v95, int32(1))
	if v99 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_s_lock(m, v96, int32(_a_F_RepackWorkerMain_3))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v103 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v94)+76)), uint32(v103))
	v107 = v94 + int32(20)
	F_SharedFileSetAttach(m, v107, v24)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v94)+96))
	F_CheckLogicalDecodingRequirements(m, int32(1))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_RepackWorkerMain[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v115
	v118 = v16 + int32(128)
	v123 = F_pg_snprintf(m, v118, int32(64), int32(_a_F_RepackWorkerMain_4), v16+int32(32))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v125 = int32(1)
	v127 = int32(0)
	F_ReplicationSlotCreate(m, v118, v125, int32(2), v127, v125, v127, v127)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_EnsureLogicalDecodingEnabled(m)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v137 = F_table_open(m, v110, int32(1))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_RepackWorkerMain[6])) = v139
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v137)))
	*(*int64)(unsafe.Add(mBase, _c_F_RepackWorkerMain[7])) = v142
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v137)+48))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+112))
	if v145 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v148 = F_table_open(m, v145, int32(1))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	F_relation_close(m, v137, int32(1))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_RepackWorkerMain[8])) = v150
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v148)))
	*(*int64)(unsafe.Add(mBase, _c_F_RepackWorkerMain[9])) = v153
	F_relation_close(m, v148, int32(1))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+124)) = int32(414)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = int32(415)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+116)) = int32(416)
	v169 = int32(1)
	v174 = int32(0)
	v177 = F_CreateInitDecodingContext(m, int32(_a_F_RepackWorkerMain_5), v169, v169, int64(0), v16+int32(116), v174, v174, v174)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v177)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v179))) = v110
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_RepackWorkerMain[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v179)+12)) = v182
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_RepackWorkerMain[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v179)+16)) = v185
	F_DecodingContextFindStartpoint(m, v177)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = int32(614)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v193)+40))
	v196 = int64(*(*int32)(unsafe.Add(mBase, _c_F_RepackWorkerMain[12])))
	v197 = base.I64_div_u_s(v194, v196)
	*(*int64)(unsafe.Add(mBase, _c_F_RepackWorkerMain[13])) = v197
	v199 = int32(_a_F_RepackWorkerMain_6)
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_RepackWorkerMain[10]))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	*(*int32)(unsafe.Add(mBase, _c_F_RepackWorkerMain[10])) = v202
	v205 = F_palloc0(m, int32(1))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+24)) = v205
	*(*int32)(unsafe.Add(mBase, _c_F_RepackWorkerMain[10])) = v200
	v213 = base.AtomicRmwXchg32(m, v96, int32(0), int32(1))
	if v213 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_s_lock(m, v96, int32(_a_F_RepackWorkerMain_3))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v217 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v94))) = uint8(v217)
	v219 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v94)+76)), uint32(v219))
	v223 = v94 + int32(100)
	F_ConditionVariableSignal(m, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	v227 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_RepackWorkerMain[14])) = uint8(v227)
	*(*int32)(unsafe.Add(mBase, _c_F_RepackWorkerMain[15])) = int32(2)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
	v233 = F_SnapBuildInitialSnapshot(m, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v235 = F_EstimateSnapshotSpace(m, v233)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+116)) = v235
	v238 = F_palloc(m, v235)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v240 = int32(0)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v233)+24))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+29)))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v233)+16))
	v249 = *(*int64)(unsafe.Add(mBase, uint32(v233)+4))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v233)+32))
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v238)+16)) = uint8(v251)
	*(*uint16)(unsafe.Add(mBase, uint32(v238)+18)) = uint16(v240)
	*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v250
	*(*int64)(unsafe.Add(mBase, uint32(v238))) = v249
	*(*int32)(unsafe.Add(mBase, uint32(v238)+8)) = v248
	*(*uint8)(unsafe.Add(mBase, uint32(v238)+17)) = uint8(v247)
	if v247&int32(1) != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v294 = v16 + int32(128)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v94)+96))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v94)+72))
	F_DecodingWorkerFileName(m, v294, v295, v296+int32(1))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L64
	}
L52:
	;
	v262 = v246
	goto L54
L53:
	;
	v262 = v240
	goto L54
L54:
	;
	if v251 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v263 = v262
	goto L57
L56:
	;
	v263 = v246
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = v263
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v233)+16))
	if v265 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v263 <= int32(0) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v269 = v265 << (uint(int32(2)) % 32)
	if v269 == int32(0) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	base.MemoryCopy(m, v238+int32(24), v274, v269)
	goto L58
L61:
	;
	goto L51
L62:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v233)+24))
	v281 = v279 << (uint(int32(2)) % 32)
	if v281 == int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v233)+16))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v233)+20))
	base.MemoryCopy(m, v238+v284<<(uint(int32(2))%32)+int32(24), v290, v281)
	goto L61
L64:
	;
	v301 = F_BufFileCreateFileSet(m, v107, v294)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_BufFileWrite(m, v301, v16+int32(116), int32(4))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v16)+116))
	F_BufFileWrite(m, v301, v238, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_BufFileClose(m, v301)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_pfree(m, v238)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v317 = base.AtomicRmwXchg32(m, v94, int32(76), int32(1))
	if v317 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	F_s_lock(m, v96, int32(_a_F_RepackWorkerMain_3))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v94)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+72)) = v321 + int32(1)
	v325 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v94)+76)), uint32(v325))
	F_ConditionVariableSignal(m, v223)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	F_InvalidateCatalogSnapshot(m)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	goto L76
L76:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v177)+140))
	v347 = v16 + int32(128)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v94)+96))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v94)+72))
	F_DecodingWorkerFileName(m, v347, v348, v349+int32(1))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v177)+140))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v517)+8))
	if v518 != 0 {
		goto L131
	} else {
		goto L132
	}
L78:
	;
	v354 = F_BufFileCreateFileSet(m, v107, v347)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v345)+20)) = v354
	v359 = base.AtomicRmwXchg32(m, v94, int32(76), int32(1))
	if v359 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	F_s_lock(m, v96, int32(_a_F_RepackWorkerMain_3))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+16)))
	v364 = *(*int64)(unsafe.Add(mBase, uint32(v94)+8))
	v365 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v94)+76)), uint32(v365))
	v368 = v364
	v375 = v363
	goto L85
L83:
	;
	goto L82
L84:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v345)+20))
	F_BufFileClose(m, v491)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L124
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+116)) = int32(0)
	v384 = *(*int32)(unsafe.Add(mBase, _c_F_RepackWorkerMain[16]))
	if v384 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L120
	}
L87:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	v390 = F_XLogReadRecord(m, v387, v16+int32(116))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	goto L89
L91:
	;
	if v368 != int64(0) {
		v455 = v368
		v456 = v375
		goto L109
	} else {
		goto L110
	}
L92:
	;
	if v390 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	F_LogicalDecodingProcessRecord(m, v177, v392)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v16)+116))
	if v427 != 0 {
		goto L4
	} else {
		goto L106
	}
L96:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	v396 = *(*int64)(unsafe.Add(mBase, uint32(v395)+40))
	v398 = int64(*(*int32)(unsafe.Add(mBase, _c_F_RepackWorkerMain[12])))
	v399 = base.I64_div_u_s(v396, v398)
	v401 = *(*int64)(unsafe.Add(mBase, _c_F_RepackWorkerMain[13]))
	if v399 == v401 {
		goto L91
	} else {
		goto L97
	}
L97:
	;
	F_LogicalIncreaseRestartDecodingForSlot(m, v396, v396)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_LogicalConfirmReceivedLocation(m, v396)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v409 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	if v409 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v16)+20)) = uint32(v396)
	v413 = int64(base.Ui64(v396) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v16)+16)) = uint32(v413)
	F_errmsg_internal(m, int32(_a_F_RepackWorkerMain_7), v16+int32(16))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_RepackWorkerMain[13])) = v399
	goto L91
L104:
	;
	F_errfinish(m, int32(_a_F_RepackWorkerMain_8), int32(429), int32(_a_F_RepackWorkerMain_9))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v428)+24))
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429))))
	if v430 != int32(1) {
		goto L3
	} else {
		goto L107
	}
L107:
	;
	v433 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v429))) = uint8(v433)
	goto L91
L108:
	;
	if v390 != 0 {
		v368 = v461
		v375 = v462
		goto L85
	} else {
		goto L117
	}
L109:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	v458 = *(*int64)(unsafe.Add(mBase, uint32(v457)+40))
	if base.Ui64(v455) <= base.Ui64(v458) {
		goto L84
	} else {
		goto L116
	}
L110:
	;
	v442 = base.AtomicRmwXchg32(m, v96, int32(0), int32(1))
	if v442 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	F_s_lock(m, v96, int32(_a_F_RepackWorkerMain_3))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+16)))
	v447 = *(*int64)(unsafe.Add(mBase, uint32(v94)+8))
	v448 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v94)+76)), uint32(v448))
	if v447 != int64(0) {
		v455 = v447
		v456 = v446
		goto L109
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	v461 = int64(0)
	v462 = v446
	v463 = int32(100)
	goto L108
L116:
	;
	v461 = v455
	v462 = v456
	v463 = int32(0)
	goto L108
L117:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	v466 = *(*int64)(unsafe.Add(mBase, uint32(v465)+40))
	v469 = F_WaitForLSN(m, int32(3), v466+int64(1), v463)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	if v469&int32(-3) == int32(0) {
		v368 = v461
		v375 = v462
		goto L85
	} else {
		goto L119
	}
L119:
	;
	goto L86
L120:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errmsg(m, int32(_a_F_RepackWorkerMain_10), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_RepackWorkerMain_8), int32(505), int32(_a_F_RepackWorkerMain_9))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	v494 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v345)+20)) = v494
	v498 = base.AtomicRmwXchg32(m, v96, v494, int32(1))
	if v498 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	F_s_lock(m, v96, int32(_a_F_RepackWorkerMain_3))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v94)+8)) = int64(0)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v94)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+72)) = v504 + int32(1)
	v508 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v94)+76)), uint32(v508))
	F_ConditionVariableSignal(m, v223)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L129
	}
L128:
	;
	goto L127
L129:
	;
	if v456&int32(1) == int32(0) {
		goto L76
	} else {
		goto L130
	}
L130:
	;
	goto L77
L131:
	;
	F_ExecDropSingleTupleTableSlot(m, v518)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	F_FreeDecodingContext(m, v177)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L135
	}
L134:
	;
	goto L133
L135:
	;
	F_ReplicationSlotDropAcquired(m, int32(1))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v529 = int32(0)
	v532 = *(*int32)(unsafe.Add(mBase, _c_F_RepackWorkerMain[17]))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v532)+16))
	v534 = m.T0[v533].(func(*base.Module, int32, int32, int32) int32)(m, int32(88), v529, v529)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	goto L13
L139:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errmsg(m, int32(_a_F_RepackWorkerMain_11), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_RepackWorkerMain_8), int32(78), int32(_a_F_RepackWorkerMain_12))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	v575 = *(*int64)(unsafe.Add(mBase, uint32(v574)+40))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v574)+1224))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v576
	*(*uint32)(unsafe.Add(mBase, uint32(v16)+8)) = uint32(v575)
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v16)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v579
	v582 = int64(base.Ui64(v575) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v16)+4)) = uint32(v582)
	F_errmsg(m, int32(_a_F_RepackWorkerMain_13), v16)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_RepackWorkerMain_8), int32(443), int32(_a_F_RepackWorkerMain_9))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
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
	F_errcode(m, int32(16779816))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	F_errmsg(m, int32(_a_F_RepackWorkerMain_14), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_RepackWorkerMain_8), int32(457), int32(_a_F_RepackWorkerMain_9))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RequestDisableLogicalDecoding(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_RequestDisableLogicalDecoding[0]))
	if v2 != int32(1) {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_RequestDisableLogicalDecoding[1]))
		v10 = F_LWLockAcquire(m, v6+int32(_a_F_RequestDisableLogicalDecoding_0), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_RequestDisableLogicalDecoding[2]))
			v14 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)) = uint8(v14)
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_RequestDisableLogicalDecoding[1]))
			F_LWLockRelease(m, v17+int32(_a_F_RequestDisableLogicalDecoding_0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_RequestDisableLogicalDecoding[3]))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+68))
				if v26 != int32(-1) {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					F_SetLatch(m, v29+v26*int32(768)+int32(316))
					mBase = m.M
				} else {
				}
				v38 = F_errstart(m, int32(14), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					if v38 == int32(0) {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_RequestDisableLogicalDecoding_1), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_RequestDisableLogicalDecoding_2), int32(474), int32(_a_F_RequestDisableLogicalDecoding_3))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
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
func F_ReservePrivateRefCountEntry(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v147 int64
	_ = v147
	var v150 int32
	_ = v150
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v157 int64
	_ = v157
	var v158 int64
	_ = v158
	var v168 int64
	_ = v168
	var v173 int32
	_ = v173
	var v174 int64
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int64
	_ = v183
	var v193 int64
	_ = v193
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int64
	_ = v316
	var v318 int64
	_ = v318
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
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
	var v376 int32
	_ = v376
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v437 int64
	_ = v437
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int64
	_ = v487
	var v489 int64
	_ = v489
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v513 int64
	_ = v513
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v528 int32
	_ = v528
	var v543 int32
	_ = v543
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v565 int32
	_ = v565
	var v575 int64
	_ = v575
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v617 int32
	_ = v617
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	v16 = int32(-1)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[0]))
	if v18 != v16 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L47
	} else {
		goto L116
	}
L2:
	;
	return
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[1]))
	if v22 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[0])) = v26
	v29 = v26
	goto L6
L5:
	;
	v29 = v16
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[2]))
	if v31 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v35 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[0])) = v35
	v38 = v35
	goto L9
L8:
	;
	v38 = v29
	goto L9
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[3]))
	if v40 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v44 = int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[0])) = v44
	v47 = v44
	goto L12
L11:
	;
	v47 = v38
	goto L12
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[4]))
	if v49 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v53 = int32(3)
	*(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[0])) = v53
	v56 = v53
	goto L15
L14:
	;
	v56 = v47
	goto L15
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[5]))
	if v58 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v62 = int32(4)
	*(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[0])) = v62
	v65 = v62
	goto L18
L17:
	;
	v65 = v56
	goto L18
L18:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[6]))
	if v67 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v71 = int32(5)
	*(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[0])) = v71
	v74 = v71
	goto L21
L20:
	;
	v74 = v65
	goto L21
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[7]))
	if v76 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v80 = int32(6)
	*(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[0])) = v80
	v83 = v80
	goto L24
L23:
	;
	v83 = v74
	goto L24
L24:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[8]))
	if v85 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[0])) = int32(7)
	return
L26:
	;
	goto L27
L27:
	;
	if v83 != int32(-1) {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v94 = int32(_a_F_ReservePrivateRefCountEntry_0)
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[9]))
	v97 = v95 & int32(7)
	*(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[0])) = v97
	*(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[9])) = v95 + int32(1)
	v104 = v97 << (uint(int32(2)) % 32)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+uint32(_c_F_ReservePrivateRefCountEntry[1])))
	v106 = int32(16)
	v110 = (int32(base.Ui32(v105)>>(uint(v106)%32)) ^ v105) * int32(-2048144789)
	v115 = (int32(base.Ui32(v110)>>(uint(int32(13))%32)) ^ v110) * int32(-1028477387)
	v122 = v97 << (uint(int32(4)) % 32)
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[10]))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	v130 = base.B2i32(base.Ui32(v127) < base.Ui32(v128))
	goto L29
L29:
	;
	if v130 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v617 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v126)+16)) = v617
	v130 = v617
	goto L29
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L47
	} else {
		goto L113
	}
L33:
	;
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v126)))
	if v147 == int64(4294967296) {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v365 = v364 & (int32(base.Ui32(v115)>>(uint(v106)%32)) ^ v115)
	v368 = v363 + v365<<(uint(int32(4))%32)
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368)+4)))
	if v369 != 0 {
		goto L79
	} else {
		goto L80
	}
L36:
	;
	v150 = int32(0)
	v152 = int64(2)
	v154 = v147 << (uint(int64(1)) % 64)
	if base.Ui64(v154) <= base.Ui64(v152) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v130 = int32(1)
	goto L29
L38:
	;
	v157 = v152
	goto L40
L39:
	;
	v157 = v154
	goto L40
L40:
	;
	v158 = int64(1)
	if v157&(v157-v158) == int64(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v168 = v157
	goto L43
L42:
	;
	v168 = v158 << (uint(int64(64)-base.I64_clz(v157)) % 64)
	goto L43
L43:
	;
	if base.Ui64(v168<<(uint(int64(4))%64)) < base.Ui64(int64(2147483647)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v126)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v126)+24))
	v180 = F_MemoryContextAllocExtended(m, v175, base.I32_wrap_i64(v168)<<(uint(int32(4))%32), int32(5))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	goto L1
L47:
	;
	return
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+20)) = v180
	v183 = int64(1)
	if v168&(v168-v183) == int64(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v193 = v168
	goto L51
L50:
	;
	v193 = v183 << (uint(int64(64)-base.I64_clz(v168)) % 64)
	goto L51
L51:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v193<<(uint(int64(4))%64)) {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v126))) = v193
	v201 = base.I32_wrap_i64(v193) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v126)+12)) = v201
	if v193 == int64(4294967296) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v210 = int32(-85899346)
	goto L55
L54:
	;
	v210 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v193), float64(0.9)))
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+16)) = v210
	if v174 != int64(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v217 = v150
	goto L60
L57:
	;
	goto L58
L58:
	;
	F_pfree(m, v173)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L47
	} else {
		goto L77
	}
L59:
	;
	v261 = v257
	v263 = v150
	goto L65
L60:
	;
	v231 = v173 + v217<<(uint(int32(4))%32)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+4)))
	if v232 != int32(1) {
		v257 = v217
		goto L59
	} else {
		goto L62
	}
L61:
	;
	v257 = int32(0)
	goto L59
L62:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v236 = int32(16)
	v240 = (int32(base.Ui32(v235)>>(uint(v236)%32)) ^ v235) * int32(-2048144789)
	v245 = (int32(base.Ui32(v240)>>(uint(int32(13))%32)) ^ v240) * int32(-1028477387)
	if (int32(base.Ui32(v245)>>(uint(v236)%32))^v245)&v201 == v217 {
		v257 = v217
		goto L59
	} else {
		goto L63
	}
L63:
	;
	v252 = v217 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v252)) < base.Ui64(v174) {
		v217 = v252
		goto L60
	} else {
		goto L64
	}
L64:
	;
	goto L61
L65:
	;
	v275 = v173 + v261<<(uint(int32(4))%32)
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+4)))
	if v276 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L58
L67:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	v280 = int32(16)
	v284 = (int32(base.Ui32(v279)>>(uint(v280)%32)) ^ v279) * int32(-2048144789)
	v289 = (int32(base.Ui32(v284)>>(uint(int32(13))%32)) ^ v284) * int32(-1028477387)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v294 = int32(base.Ui32(v289)>>(uint(v280)%32)) ^ v289
	goto L70
L68:
	;
	goto L69
L69:
	;
	v336 = v261 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v336)) < base.Ui64(v174) {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v309 = v294 & v293
	v314 = v180 + v309<<(uint(int32(4))%32)
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+4)))
	if v315 != 0 {
		v294 = v309 + int32(1)
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v275)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v314)+8)) = v316
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v275)))
	*(*int64)(unsafe.Add(mBase, uint32(v314))) = v318
	goto L69
L72:
	;
	goto L71
L73:
	;
	v340 = v336
	goto L75
L74:
	;
	v340 = int32(0)
	goto L75
L75:
	;
	v342 = v263 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v342)) < base.Ui64(v174) {
		v261 = v340
		v263 = v342
		goto L65
	} else {
		goto L76
	}
L76:
	;
	goto L66
L77:
	;
	goto L37
L78:
	;
	v575 = *(*int64)(unsafe.Add(mBase, uint32(v122)+uint32(_c_F_ReservePrivateRefCountEntry[11])))
	*(*int64)(unsafe.Add(mBase, uint32(v565)+8)) = v575
	v577 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v104)+uint32(_c_F_ReservePrivateRefCountEntry[1]))) = v577
	*(*int64)(unsafe.Add(mBase, uint32(v122)+uint32(_c_F_ReservePrivateRefCountEntry[11]))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v122)+uint32(_c_F_ReservePrivateRefCountEntry[12]))) = v577
	v583 = int32(_a_F_ReservePrivateRefCountEntry_1)
	v585 = *(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[13]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReservePrivateRefCountEntry[13])) = v585 + int32(1)
	goto L2
L79:
	;
	v371 = v365
	v374 = int32(0)
	v376 = v368
	goto L83
L80:
	;
	v543 = v368
	goto L81
L81:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	v554 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v126)+8)) = v553 + v554
	*(*uint8)(unsafe.Add(mBase, uint32(v543)+4)) = uint8(v554)
	*(*int32)(unsafe.Add(mBase, uint32(v543))) = v105
	v565 = v543
	goto L78
L82:
	;
	v543 = v528
	goto L81
L83:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v376)))
	if v386 == v105 {
		v565 = v376
		goto L78
	} else {
		goto L85
	}
L84:
	;
	v528 = v521
	goto L82
L85:
	;
	v389 = v371 + int32(1)
	v390 = int32(16)
	v394 = (int32(base.Ui32(v386)>>(uint(v390)%32)) ^ v386) * int32(-2048144789)
	v399 = (int32(base.Ui32(v394)>>(uint(int32(13))%32)) ^ v394) * int32(-1028477387)
	v403 = (int32(base.Ui32(v399)>>(uint(v390)%32)) ^ v399) & v364
	if base.Ui32(v371) < base.Ui32(v403) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v407 = v371 + v405
	goto L88
L87:
	;
	v407 = v371
	goto L88
L88:
	;
	if base.Ui32(v407-v403) < base.Ui32(v374) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v410 = v389 & v364
	v413 = v363 + v410<<(uint(int32(4))%32)
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413)+4)))
	if v414 != 0 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	v508 = v374 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v508) {
		goto L108
	} else {
		goto L109
	}
L92:
	;
	v417 = int32(0)
	v420 = v410
	goto L95
L93:
	;
	v452 = v413
	v453 = v410
	goto L94
L94:
	;
	if v371 != v453 {
		goto L102
	} else {
		goto L103
	}
L95:
	;
	v432 = v417 + int32(1)
	if int32(151) <= v432 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v452 = v447
	v453 = v444
	goto L94
L97:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	v437 = *(*int64)(unsafe.Add(mBase, uint32(v126)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v435), base.F64_convert_i64_u(v437)), float64(0.1)) != 0 {
		goto L31
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v444 = (v420 + int32(1)) & v364
	v447 = v363 + v444<<(uint(int32(4))%32)
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+4)))
	if v448 != 0 {
		v417 = v432
		v420 = v444
		goto L95
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	goto L96
L102:
	;
	v468 = v452
	v469 = v453
	goto L105
L103:
	;
	goto L104
L104:
	;
	v528 = v376
	goto L82
L105:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v483 = v480 & (v469 - int32(1))
	v486 = v363 + v483<<(uint(int32(4))%32)
	v487 = *(*int64)(unsafe.Add(mBase, uint32(v486)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v468)+8)) = v487
	v489 = *(*int64)(unsafe.Add(mBase, uint32(v486)))
	*(*int64)(unsafe.Add(mBase, uint32(v468))) = v489
	if v371 != v483 {
		v468 = v486
		v469 = v483
		goto L105
	} else {
		goto L107
	}
L106:
	;
	goto L104
L107:
	;
	goto L106
L108:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	v513 = *(*int64)(unsafe.Add(mBase, uint32(v126)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v511), base.F64_convert_i64_u(v513)), float64(0.1)) != 0 {
		goto L31
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v518 = v389 & v364
	v521 = v363 + v518<<(uint(int32(4))%32)
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521)+4)))
	if v522 != 0 {
		v371 = v518
		v374 = v508
		v376 = v521
		goto L83
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	goto L84
L113:
	;
	F_errmsg_internal(m, int32(_a_F_ReservePrivateRefCountEntry_2), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L47
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_ReservePrivateRefCountEntry_3), int32(635), int32(_a_F_ReservePrivateRefCountEntry_4))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L47
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	F_errmsg_internal(m, int32(_a_F_ReservePrivateRefCountEntry_5), int32(0))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L47
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_ReservePrivateRefCountEntry_3), int32(332), int32(_a_F_ReservePrivateRefCountEntry_6))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L47
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ResetLatch(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	v2 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v2
	v7 = base.AtomicRmwOr32(m, v2, int32(_a_F_ResetLatch_0), v2)
	return
}
func F_ResetUnloggedRelations(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
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
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	v7 = m.G0
	v9 = v7 - int32(1088)
	m.G0 = v9
	v13 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v13 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v15 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0 & v15
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(base.Ui32(l0)>>(uint(v15)%32)) & v15
	F_errmsg_internal(m, int32(_a_F_ResetUnloggedRelations_0), v9+int32(16))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelations[0]))
	v39 = F_AllocSetContextCreateInternal(m, v34, int32(_a_F_ResetUnloggedRelations_1), int32(0), int32(_a_F_ResetUnloggedRelations_2), int32(_a_F_ResetUnloggedRelations_3))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	F_errfinish(m, int32(_a_F_ResetUnloggedRelations_4), int32(58), int32(_a_F_ResetUnloggedRelations_1))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v41 = int32(_a_F_ResetUnloggedRelations_5)
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelations[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelations[0])) = v39
	F_begin_startup_progress_phase(m)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_ResetUnloggedRelationsInTablespaceDir(m, int32(_a_F_ResetUnloggedRelations_6), l0)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v51 = F_AllocateDir(m, int32(_a_F_ResetUnloggedRelations_7))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v54 = F_ReadDir(m, v51, int32(_a_F_ResetUnloggedRelations_7))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v54 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v58 = v54
	goto L16
L14:
	;
	goto L15
L15:
	;
	F_FreeDir(m, v51)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L28
	}
L16:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+19)))
	if v62 != int32(46) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L15
L18:
	;
	v91 = F_ReadDir(m, v51, int32(_a_F_ResetUnloggedRelations_7))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L26
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(_a_F_ResetUnloggedRelations_8)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v58 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(_a_F_ResetUnloggedRelations_7)
	v82 = v9 + int32(32)
	v85 = F_pg_snprintf(m, v82, int32(1050), int32(_a_F_ResetUnloggedRelations_9), v9)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L24
	}
L20:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+20)))
	if v65 == int32(0) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+20)))
	if v68 != int32(46) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+21)))
	if v71 == int32(0) {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	F_ResetUnloggedRelationsInTablespaceDir(m, v82, l0)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	goto L18
L26:
	;
	if v91 != 0 {
		v58 = v91
		goto L16
	} else {
		goto L27
	}
L27:
	;
	goto L17
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ResetUnloggedRelations[0])) = v42
	F_MemoryContextDelete(m, v39)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	m.G0 = v9 + int32(1088)
	return
}
func F_RmgrNotFound(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		F_errmsg(m, int32(_a_F_RmgrNotFound_0), v5)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_errhint(m, int32(_a_F_RmgrNotFound_1), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_RmgrNotFound_2), int32(94), int32(_a_F_RmgrNotFound_3))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
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
func F___rem_pio2(m *base.Module, l0 float64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v48 float64
	_ = v48
	var v49 float64
	_ = v49
	var v50 float64
	_ = v50
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v60 float64
	_ = v60
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	var v82 float64
	_ = v82
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v108 float64
	_ = v108
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v122 float64
	_ = v122
	var v123 float64
	_ = v123
	var v124 float64
	_ = v124
	var v132 float64
	_ = v132
	var v133 float64
	_ = v133
	var v134 float64
	_ = v134
	var v148 float64
	_ = v148
	var v149 int32
	_ = v149
	var v152 float64
	_ = v152
	var v154 float64
	_ = v154
	var v155 float64
	_ = v155
	var v161 float64
	_ = v161
	var v174 float64
	_ = v174
	var v180 int32
	_ = v180
	var v181 float64
	_ = v181
	var v182 float64
	_ = v182
	var v183 float64
	_ = v183
	var v184 float64
	_ = v184
	var v187 int32
	_ = v187
	var v198 float64
	_ = v198
	var v199 float64
	_ = v199
	var v204 float64
	_ = v204
	var v205 float64
	_ = v205
	var v217 float64
	_ = v217
	var v218 float64
	_ = v218
	var v223 float64
	_ = v223
	var v224 float64
	_ = v224
	var v226 float64
	_ = v226
	var v227 float64
	_ = v227
	var v228 float64
	_ = v228
	var v235 float64
	_ = v235
	var v240 int32
	_ = v240
	var v249 float64
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v276 float64
	_ = v276
	var v280 float64
	_ = v280
	var v289 int32
	_ = v289
	var v319 float64
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v389 int32
	_ = v389
	var v391 float64
	_ = v391
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v465 float64
	_ = v465
	var v467 int32
	_ = v467
	var v491 int32
	_ = v491
	var v494 float64
	_ = v494
	var v501 float64
	_ = v501
	var v503 float64
	_ = v503
	var v505 int32
	_ = v505
	var v507 float64
	_ = v507
	var v545 int32
	_ = v545
	var v555 int32
	_ = v555
	var v581 float64
	_ = v581
	var v582 int32
	_ = v582
	var v585 float64
	_ = v585
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v619 float64
	_ = v619
	var v630 float64
	_ = v630
	var v631 float64
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 float64
	_ = v637
	var v666 float64
	_ = v666
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v682 float64
	_ = v682
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v695 float64
	_ = v695
	var v696 int32
	_ = v696
	var v703 float64
	_ = v703
	var v709 float64
	_ = v709
	var v710 int32
	_ = v710
	var v712 float64
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v806 int32
	_ = v806
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v846 float64
	_ = v846
	var v847 int32
	_ = v847
	var v848 float64
	_ = v848
	var v852 float64
	_ = v852
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v868 float64
	_ = v868
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v881 float64
	_ = v881
	var v882 int32
	_ = v882
	var v891 float64
	_ = v891
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v1026 int32
	_ = v1026
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1065 int32
	_ = v1065
	var v1090 int32
	_ = v1090
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1103 float64
	_ = v1103
	var v1106 float64
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1132 int32
	_ = v1132
	var v1135 float64
	_ = v1135
	var v1142 float64
	_ = v1142
	var v1144 float64
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1148 float64
	_ = v1148
	var v1180 int32
	_ = v1180
	var v1184 float64
	_ = v1184
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1200 float64
	_ = v1200
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1213 float64
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1221 float64
	_ = v1221
	var v1231 int32
	_ = v1231
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1253 int32
	_ = v1253
	var v1258 int32
	_ = v1258
	var v1276 float64
	_ = v1276
	var v1280 float64
	_ = v1280
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1296 float64
	_ = v1296
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1309 float64
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1320 float64
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1354 int32
	_ = v1354
	var v1369 int32
	_ = v1369
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1399 float64
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1426 int32
	_ = v1426
	var v1427 float64
	_ = v1427
	var v1429 float64
	_ = v1429
	var v1431 float64
	_ = v1431
	var v1435 float64
	_ = v1435
	var v1497 float64
	_ = v1497
	var v1500 float64
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1533 float64
	_ = v1533
	var v1534 float64
	_ = v1534
	var v1535 float64
	_ = v1535
	var v1562 float64
	_ = v1562
	var v1564 float64
	_ = v1564
	var v1565 float64
	_ = v1565
	var v1569 float64
	_ = v1569
	var v1571 int32
	_ = v1571
	var v1600 float64
	_ = v1600
	var v1601 float64
	_ = v1601
	var v1605 float64
	_ = v1605
	var v1632 float64
	_ = v1632
	var v1638 int32
	_ = v1638
	var v1639 float64
	_ = v1639
	var v1644 float64
	_ = v1644
	var v1650 float64
	_ = v1650
	var v1655 int32
	_ = v1655
	v27 = m.G0
	v29 = v27 - int32(48)
	m.G0 = v29
	v31 = base.I64_reinterpret_f64(l0)
	v34 = base.I32_wrap_i64(int64(base.Ui64(v31) >> (uint(int64(32)) % 64)))
	v36 = v34 & int32(2147483647)
	if base.Ui32(v36) <= base.Ui32(int32(1074752122)) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v29 + int32(48)
	return v1655
L2:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v36) {
		goto L42
	} else {
		goto L43
	}
L3:
	;
	v148 = base.F64_add(base.F64_add(base.F64_mul(l0, float64(0.6366197723675814)), float64(6.755399441055744e+15)), float64(-6.755399441055744e+15))
	v149 = base.I32_trunc_sat_f64_s(v148)
	v152 = base.F64_add(l0, base.F64_mul(v148, float64(-1.5707963267341256)))
	v154 = base.F64_mul(v148, float64(6.077100506506192e-11))
	v155 = base.F64_sub(v152, v154)
	if base.F64_lt(v155, float64(-0.7853981633974483)) != 0 {
		goto L33
	} else {
		goto L34
	}
L4:
	;
	if v34&int32(_a_F___rem_pio2_0) == int32(_a_F___rem_pio2_1) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(v36) <= base.Ui32(int32(1075594811)) {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	if base.Ui32(v36) <= base.Ui32(int32(1073928572)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if int64(0) <= v31 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	if int64(0) <= v31 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v48 = base.F64_add(l0, float64(-1.5707963267341256))
	v49 = float64(-6.077100506506192e-11)
	v50 = base.F64_add(v48, v49)
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v50
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(base.F64_sub(v48, v50), v49)
	v1655 = int32(1)
	goto L1
L12:
	;
	goto L13
L13:
	;
	v58 = base.F64_add(l0, float64(1.5707963267341256))
	v59 = float64(6.077100506506192e-11)
	v60 = base.F64_add(v58, v59)
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v60
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(base.F64_sub(v58, v60), v59)
	v1655 = int32(-1)
	goto L1
L14:
	;
	v70 = base.F64_add(l0, float64(-3.1415926534682512))
	v71 = float64(-1.2154201013012384e-10)
	v72 = base.F64_add(v70, v71)
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v72
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(base.F64_sub(v70, v72), v71)
	v1655 = int32(2)
	goto L1
L15:
	;
	goto L16
L16:
	;
	v80 = base.F64_add(l0, float64(3.1415926534682512))
	v81 = float64(1.2154201013012384e-10)
	v82 = base.F64_add(v80, v81)
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v82
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(base.F64_sub(v80, v82), v81)
	v1655 = int32(-2)
	goto L1
L17:
	;
	if base.Ui32(v36) <= base.Ui32(int32(1075183036)) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	if base.Ui32(int32(1094263290)) < base.Ui32(v36) {
		goto L2
	} else {
		goto L31
	}
L20:
	;
	if v36 == int32(1074977148) {
		goto L3
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v36 == int32(1075388923) {
		goto L3
	} else {
		goto L27
	}
L23:
	;
	if int64(0) <= v31 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v98 = base.F64_add(l0, float64(-4.712388980202377))
	v99 = float64(-1.8231301519518578e-10)
	v100 = base.F64_add(v98, v99)
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v100
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(base.F64_sub(v98, v100), v99)
	v1655 = int32(3)
	goto L1
L25:
	;
	goto L26
L26:
	;
	v108 = base.F64_add(l0, float64(4.712388980202377))
	v109 = float64(1.8231301519518578e-10)
	v110 = base.F64_add(v108, v109)
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v110
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(base.F64_sub(v108, v110), v109)
	v1655 = int32(-3)
	goto L1
L27:
	;
	if int64(0) <= v31 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v122 = base.F64_add(l0, float64(-6.2831853069365025))
	v123 = float64(-2.430840202602477e-10)
	v124 = base.F64_add(v122, v123)
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v124
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(base.F64_sub(v122, v124), v123)
	v1655 = int32(4)
	goto L1
L29:
	;
	goto L30
L30:
	;
	v132 = base.F64_add(l0, float64(6.2831853069365025))
	v133 = float64(2.430840202602477e-10)
	v134 = base.F64_add(v132, v133)
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v134
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(base.F64_sub(v132, v134), v133)
	v1655 = int32(-4)
	goto L1
L31:
	;
	goto L3
L32:
	;
	v184 = base.F64_sub(v181, v183)
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v184
	v187 = int32(base.Ui32(v36) >> (uint(int32(20)) % 32))
	if v187-base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v184))>>(uint(int64(52))%64)))&int32(2047) < int32(17) {
		v226 = v184
		v227 = v181
		v228 = v183
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v161 = base.F64_add(v148, float64(-1))
	v180 = v149 - int32(1)
	v181 = base.F64_add(l0, base.F64_mul(v161, float64(-1.5707963267341256)))
	v182 = v161
	v183 = base.F64_mul(v161, float64(6.077100506506192e-11))
	goto L32
L34:
	;
	goto L35
L35:
	;
	if base.F64_gt(v155, float64(0.7853981633974483)) == int32(0) {
		v180 = v149
		v181 = v152
		v182 = v148
		v183 = v154
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v174 = base.F64_add(v148, float64(1))
	v180 = v149 + int32(1)
	v181 = base.F64_add(l0, base.F64_mul(v174, float64(-1.5707963267341256)))
	v182 = v174
	v183 = base.F64_mul(v174, float64(6.077100506506192e-11))
	goto L32
L37:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_sub(base.F64_sub(v227, v226), v228)
	v1655 = v180
	goto L1
L38:
	;
	v198 = base.F64_mul(v182, float64(6.077100506303966e-11))
	v199 = base.F64_sub(v181, v198)
	v204 = base.F64_sub(base.F64_mul(v182, float64(2.0222662487959506e-21)), base.F64_sub(base.F64_sub(v181, v199), v198))
	v205 = base.F64_sub(v199, v204)
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v205
	if v187-base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v205))>>(uint(int64(52))%64)))&int32(2047) < int32(50) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v226 = v205
	v227 = v199
	v228 = v204
	goto L37
L40:
	;
	goto L41
L41:
	;
	v217 = base.F64_mul(v182, float64(2.0222662487111665e-21))
	v218 = base.F64_sub(v199, v217)
	v223 = base.F64_sub(base.F64_mul(v182, float64(8.4784276603689e-32)), base.F64_sub(base.F64_sub(v199, v218), v217))
	v224 = base.F64_sub(v218, v223)
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v224
	v226 = v224
	v227 = v218
	v228 = v223
	goto L37
L42:
	;
	v235 = base.F64_sub(l0, l0)
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v235
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v235
	v1655 = int32(0)
	goto L1
L43:
	;
	goto L44
L44:
	;
	v240 = v29 + int32(16)
	v249 = base.F64_reinterpret_i64(v31&int64(4503599627370495) | int64(4710765210229538816))
	v251 = int32(1)
	v252 = v240
	goto L45
L45:
	;
	v276 = base.F64_convert_i32_s(base.I32_trunc_sat_f64_s(v249))
	*(*float64)(unsafe.Add(mBase, uint32(v252))) = v276
	v280 = base.F64_mul(base.F64_sub(v249, v276), float64(1.6777216e+07))
	if v251&int32(1) != 0 {
		v249 = v280
		v251 = int32(0)
		v252 = v240 | int32(8)
		goto L45
	} else {
		goto L47
	}
L46:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29)+32)) = v280
	v289 = int32(2)
	goto L48
L47:
	;
	goto L46
L48:
	;
	v319 = *(*float64)(unsafe.Add(mBase, uint32(v29+int32(16)+v289<<(uint(int32(3))%32))))
	if base.F64_eq(v319, float64(0)) != 0 {
		v289 = v289 - int32(1)
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v322 = m.G0
	v324 = v322 - int32(560)
	m.G0 = v324
	v329 = int32(base.Ui32(v36)>>(uint(int32(20))%32)) - int32(1046)
	v333 = base.I32_div_s(v329-int32(3), int32(24))
	v334 = int32(0)
	if v334 < v333 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L49
L51:
	;
	v337 = v333
	goto L53
L52:
	;
	v337 = v334
	goto L53
L53:
	;
	v340 = v329 + v337*int32(-24)
	v342 = *(*int32)(unsafe.Add(mBase, _c_F___rem_pio2[0]))
	v343 = int32(1)
	v344 = v289 + v343
	v346 = v344 - v343
	if int32(0) <= v342+v346 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v355 = v337 - v346
	v357 = int32(0)
	goto L57
L55:
	;
	goto L56
L56:
	;
	v425 = v29 + int32(16)
	v427 = v340 - int32(24)
	v428 = int32(0)
	if v428 < v342 {
		goto L63
	} else {
		goto L64
	}
L57:
	;
	if v355 < int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L56
L59:
	;
	v391 = float64(0)
	goto L61
L60:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v355<<(uint(int32(2))%32))+uint32(_c_F___rem_pio2[1])))
	v391 = base.F64_convert_i32_s(v389)
	goto L61
L61:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v324+int32(320)+v357<<(uint(int32(3))%32)))) = v391
	v393 = int32(1)
	v396 = v357 + v393
	if v396 != v342+v344 {
		v355 = v355 + v393
		v357 = v396
		goto L57
	} else {
		goto L62
	}
L62:
	;
	goto L58
L63:
	;
	v432 = v342
	goto L65
L64:
	;
	v432 = v428
	goto L65
L65:
	;
	v438 = v428
	goto L66
L66:
	;
	if v344 <= int32(0) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v545 = int32(48) - v340
	v555 = v342
	goto L77
L68:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v324+v438<<(uint(int32(3))%32)))) = v507
	if base.B2i32(v438 == v432) == int32(0) {
		v438 = v438 + int32(1)
		goto L66
	} else {
		goto L75
	}
L69:
	;
	v507 = float64(0)
	goto L68
L70:
	;
	goto L71
L71:
	;
	v465 = float64(0)
	v467 = int32(0)
	goto L72
L72:
	;
	v491 = int32(3)
	v494 = *(*float64)(unsafe.Add(mBase, uint32(v425+v467<<(uint(v491)%32))))
	v501 = *(*float64)(unsafe.Add(mBase, uint32(v324+int32(320)+(v438+v346-v467)<<(uint(v491)%32))))
	v503 = base.F64_add(base.F64_mul(v494, v501), v465)
	v505 = v467 + int32(1)
	if v505 != v344 {
		v465 = v503
		v467 = v505
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v507 = v503
	goto L68
L74:
	;
	goto L73
L75:
	;
	goto L67
L76:
	;
	v1276 = float64(1)
	if int32(1024) <= v1258 {
		goto L197
	} else {
		goto L198
	}
L77:
	;
	v581 = *(*float64)(unsafe.Add(mBase, uint32(v324+v555<<(uint(int32(3))%32))))
	v582 = int32(0)
	if v582 < v555 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v1180 = int32(24) - v340
	if int32(1024) <= v1180 {
		goto L176
	} else {
		goto L177
	}
L79:
	;
	v585 = v581
	v587 = v582
	v589 = v555
	goto L82
L80:
	;
	v637 = v581
	goto L81
L81:
	;
	if int32(1024) <= v427 {
		goto L87
	} else {
		goto L88
	}
L82:
	;
	v619 = base.F64_convert_i32_s(base.I32_trunc_sat_f64_s(base.F64_mul(v585, float64(5.960464477539063e-08))))
	*(*int32)(unsafe.Add(mBase, uint32(v324+int32(480)+v587<<(uint(int32(2))%32)))) = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v619, float64(-1.6777216e+07)), v585))
	v630 = *(*float64)(unsafe.Add(mBase, uint32(v324+v589<<(uint(int32(3))%32)-int32(8))))
	v631 = base.F64_add(v630, v619)
	v632 = int32(1)
	v635 = v587 + v632
	if v635 != v555 {
		v585 = v631
		v587 = v635
		v589 = v589 - v632
		goto L82
	} else {
		goto L84
	}
L83:
	;
	v637 = v631
	goto L81
L84:
	;
	goto L83
L85:
	;
	v709 = base.F64_add(v703, base.F64_mul(base.F64_floor(base.F64_mul(v703, float64(0.125))), float64(-8)))
	v710 = base.I32_trunc_sat_f64_s(v709)
	v712 = base.F64_sub(v709, base.F64_convert_i32_s(v710))
	v713 = int32(0)
	v714 = base.B2i32(v427 <= v713)
	if v714 == v713 {
		goto L107
	} else {
		goto L108
	}
L86:
	;
	v703 = base.F64_mul(v695, base.F64_reinterpret_i64(base.I64_extend_i32_u(v696+int32(1023))<<(uint(int64(52))%64)))
	goto L85
L87:
	;
	v666 = base.F64_mul(v637, float64(8.98846567431158e+307))
	if base.Ui32(v427) < base.Ui32(int32(2047)) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	if int32(-1023) < v427 {
		v695 = v637
		v696 = v427
		goto L86
	} else {
		goto L96
	}
L90:
	;
	v695 = v666
	v696 = v427 - int32(1023)
	goto L86
L91:
	;
	goto L92
L92:
	;
	v673 = int32(3069)
	if base.Ui32(v673) <= base.Ui32(v427) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v676 = v673
	goto L95
L94:
	;
	v676 = v427
	goto L95
L95:
	;
	v695 = base.F64_mul(v666, float64(8.98846567431158e+307))
	v696 = v676 - int32(2046)
	goto L86
L96:
	;
	v682 = base.F64_mul(v637, float64(2.004168360008973e-292))
	if base.Ui32(int32(-1992)) < base.Ui32(v427) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v695 = v682
	v696 = v427 + int32(969)
	goto L86
L98:
	;
	goto L99
L99:
	;
	v689 = int32(-2960)
	if base.Ui32(v427) <= base.Ui32(v689) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v692 = v689
	goto L102
L101:
	;
	v692 = v427
	goto L102
L102:
	;
	v695 = base.F64_mul(v682, float64(2.004168360008973e-292))
	v696 = v692 + int32(1938)
	goto L86
L103:
	;
	if base.F64_eq(v891, float64(0)) != 0 {
		goto L149
	} else {
		goto L150
	}
L104:
	;
	v749 = int32(0)
	if v749 < v555 {
		goto L113
	} else {
		goto L114
	}
L105:
	;
	if base.F64_ge(v712, float64(0.5)) != 0 {
		v747 = int32(2)
		v748 = v710
		goto L104
	} else {
		goto L112
	}
L106:
	;
	if v738 <= int32(0) {
		v891 = v712
		v902 = v738
		v904 = v737
		goto L103
	} else {
		goto L111
	}
L107:
	;
	v721 = v555<<(uint(int32(2))%32) + v324 + int32(476)
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v721)))
	v723 = v722 >> (uint(v545) % 32)
	v725 = v722 - v723<<(uint(v545)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v721))) = v725
	v737 = v723 + v710
	v738 = v725 >> (uint(int32(47)-v340) % 32)
	goto L106
L108:
	;
	goto L109
L109:
	;
	if v427 != 0 {
		goto L105
	} else {
		goto L110
	}
L110:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v555<<(uint(int32(2))%32)+v324)+476))
	v737 = v710
	v738 = v732 >> (uint(int32(23)) % 32)
	goto L106
L111:
	;
	v747 = v738
	v748 = v737
	goto L104
L112:
	;
	v891 = v712
	v902 = int32(0)
	v904 = v710
	goto L103
L113:
	;
	v756 = v749
	v761 = v749
	goto L116
L114:
	;
	v806 = int32(1)
	goto L115
L115:
	;
	if v427 <= v713 {
		goto L125
	} else {
		goto L126
	}
L116:
	;
	v784 = v324 + int32(480) + v756<<(uint(int32(2))%32)
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v784)))
	if v761 != 0 {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	v806 = v797
	goto L115
L118:
	;
	v800 = v756 + int32(1)
	if v800 != v555 {
		v756 = v800
		v761 = v798
		goto L116
	} else {
		goto L124
	}
L119:
	;
	v797 = int32(1)
	v798 = int32(0)
	goto L118
L120:
	;
	v790 = int32(16777215)
	goto L122
L121:
	;
	if v785 == int32(0) {
		goto L119
	} else {
		goto L123
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v784))) = v790 - v785
	v797 = int32(0)
	v798 = int32(1)
	goto L118
L123:
	;
	v790 = int32(16777216)
	goto L122
L124:
	;
	goto L117
L125:
	;
	v842 = v748 + int32(1)
	if v747 != int32(2) {
		v891 = v712
		v902 = v747
		v904 = v842
		goto L103
	} else {
		goto L129
	}
L126:
	;
	switch v340 - int32(25) {
	case 0:
		v830 = int32(_a_F___rem_pio2_2)
		goto L127
	case 1:
		goto L128
	default:
		goto L125
	}
L127:
	;
	v835 = v555<<(uint(int32(2))%32) + v324 + int32(476)
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v835)))
	*(*int32)(unsafe.Add(mBase, uint32(v835))) = v836 & v830
	goto L125
L128:
	;
	v830 = int32(_a_F___rem_pio2_3)
	goto L127
L129:
	;
	v846 = base.F64_sub(float64(1), v712)
	v847 = int32(2)
	if v806 != 0 {
		v891 = v846
		v902 = v847
		v904 = v842
		goto L103
	} else {
		goto L130
	}
L130:
	;
	v848 = float64(1)
	if int32(1024) <= v427 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v891 = base.F64_sub(v846, base.F64_mul(v881, base.F64_reinterpret_i64(base.I64_extend_i32_u(v882+int32(1023))<<(uint(int64(52))%64))))
	v902 = v847
	v904 = v842
	goto L103
L132:
	;
	goto L131
L133:
	;
	v852 = base.F64_mul(v848, float64(8.98846567431158e+307))
	if base.Ui32(v427) < base.Ui32(int32(2047)) {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	goto L135
L135:
	;
	if int32(-1023) < v427 {
		v881 = v848
		v882 = v427
		goto L132
	} else {
		goto L142
	}
L136:
	;
	v881 = v852
	v882 = v427 - int32(1023)
	goto L132
L137:
	;
	goto L138
L138:
	;
	v859 = int32(3069)
	if base.Ui32(v859) <= base.Ui32(v427) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v862 = v859
	goto L141
L140:
	;
	v862 = v427
	goto L141
L141:
	;
	v881 = base.F64_mul(v852, float64(8.98846567431158e+307))
	v882 = v862 - int32(2046)
	goto L132
L142:
	;
	v868 = base.F64_mul(v848, float64(2.004168360008973e-292))
	if base.Ui32(int32(-1992)) < base.Ui32(v427) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v881 = v868
	v882 = v427 + int32(969)
	goto L132
L144:
	;
	goto L145
L145:
	;
	v875 = int32(-2960)
	if base.Ui32(v427) <= base.Ui32(v875) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v878 = v875
	goto L148
L147:
	;
	v878 = v427
	goto L148
L148:
	;
	v881 = base.F64_mul(v868, float64(2.004168360008973e-292))
	v882 = v878 + int32(1938)
	goto L132
L149:
	;
	if v555 <= v342 {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	goto L151
L151:
	;
	goto L78
L152:
	;
	v1026 = int32(1)
	goto L161
L153:
	;
	v923 = v555
	v925 = int32(0)
	goto L154
L154:
	;
	v950 = v923 - int32(1)
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v324+int32(480)+v950<<(uint(int32(2))%32))))
	v955 = v954 | v925
	if v342 < v950 {
		v923 = v950
		v925 = v955
		goto L154
	} else {
		goto L156
	}
L155:
	;
	if v955 == int32(0) {
		goto L152
	} else {
		goto L157
	}
L156:
	;
	goto L155
L157:
	;
	v962 = v555
	v967 = v427
	goto L158
L158:
	;
	v986 = v967 - int32(24)
	v990 = v962 - int32(1)
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v324+int32(480)+v990<<(uint(int32(2))%32))))
	if v994 == int32(0) {
		v962 = v990
		v967 = v986
		goto L158
	} else {
		goto L160
	}
L159:
	;
	v1253 = v990
	v1258 = v986
	goto L76
L160:
	;
	goto L159
L161:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v324+int32(480)+(v342-v1026)<<(uint(int32(2))%32))))
	if v1058 == int32(0) {
		v1026 = v1026 + int32(1)
		goto L161
	} else {
		goto L163
	}
L162:
	;
	v1061 = v555 + v1026
	v1065 = v555
	goto L164
L163:
	;
	goto L162
L164:
	;
	v1090 = v1065 + v344
	v1095 = v1065 + int32(1)
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v337<<(uint(int32(2))%32)+int32(_a_F___rem_pio2_4)+v1095<<(uint(int32(2))%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v324+int32(320)+v1090<<(uint(int32(3))%32)))) = base.F64_convert_i32_s(v1099)
	v1102 = int32(0)
	v1103 = float64(0)
	if v1102 < v344 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	v555 = v1061
	goto L77
L166:
	;
	v1106 = v1103
	v1108 = v1102
	goto L169
L167:
	;
	v1148 = v1103
	goto L168
L168:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v324+v1095<<(uint(int32(3))%32)))) = v1148
	if v1095 < v1061 {
		v1065 = v1095
		goto L164
	} else {
		goto L172
	}
L169:
	;
	v1132 = int32(3)
	v1135 = *(*float64)(unsafe.Add(mBase, uint32(v425+v1108<<(uint(v1132)%32))))
	v1142 = *(*float64)(unsafe.Add(mBase, uint32(v324+int32(320)+(v1090-v1108)<<(uint(v1132)%32))))
	v1144 = base.F64_add(base.F64_mul(v1135, v1142), v1106)
	v1146 = v1108 + int32(1)
	if v1146 != v344 {
		v1106 = v1144
		v1108 = v1146
		goto L169
	} else {
		goto L171
	}
L170:
	;
	v1148 = v1144
	goto L168
L171:
	;
	goto L170
L172:
	;
	goto L165
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v324+int32(480)+v1242<<(uint(int32(2))%32)))) = v1241
	v1253 = v1242
	v1258 = v1243
	goto L76
L174:
	;
	if base.F64_ge(v1221, float64(1.6777216e+07)) != 0 {
		goto L192
	} else {
		goto L193
	}
L175:
	;
	v1221 = base.F64_mul(v1213, base.F64_reinterpret_i64(base.I64_extend_i32_u(v1214+int32(1023))<<(uint(int64(52))%64)))
	goto L174
L176:
	;
	v1184 = base.F64_mul(v891, float64(8.98846567431158e+307))
	if base.Ui32(v1180) < base.Ui32(int32(2047)) {
		goto L179
	} else {
		goto L180
	}
L177:
	;
	goto L178
L178:
	;
	if int32(-1023) < v1180 {
		v1213 = v891
		v1214 = v1180
		goto L175
	} else {
		goto L185
	}
L179:
	;
	v1213 = v1184
	v1214 = v1180 - int32(1023)
	goto L175
L180:
	;
	goto L181
L181:
	;
	v1191 = int32(3069)
	if base.Ui32(v1191) <= base.Ui32(v1180) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v1194 = v1191
	goto L184
L183:
	;
	v1194 = v1180
	goto L184
L184:
	;
	v1213 = base.F64_mul(v1184, float64(8.98846567431158e+307))
	v1214 = v1194 - int32(2046)
	goto L175
L185:
	;
	v1200 = base.F64_mul(v891, float64(2.004168360008973e-292))
	if base.Ui32(int32(-1992)) < base.Ui32(v1180) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v1213 = v1200
	v1214 = v1180 + int32(969)
	goto L175
L187:
	;
	goto L188
L188:
	;
	v1207 = int32(-2960)
	if base.Ui32(v1180) <= base.Ui32(v1207) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v1210 = v1207
	goto L191
L190:
	;
	v1210 = v1180
	goto L191
L191:
	;
	v1213 = base.F64_mul(v1200, float64(2.004168360008973e-292))
	v1214 = v1210 + int32(1938)
	goto L175
L192:
	;
	v1231 = base.I32_trunc_sat_f64_s(base.F64_mul(v1221, float64(5.960464477539063e-08)))
	*(*int32)(unsafe.Add(mBase, uint32(v324+int32(480)+v555<<(uint(int32(2))%32)))) = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1231), float64(-1.6777216e+07)), v1221))
	v1241 = v1231
	v1242 = v555 + int32(1)
	v1243 = v340
	goto L173
L193:
	;
	goto L194
L194:
	;
	v1241 = base.I32_trunc_sat_f64_s(v1221)
	v1242 = v555
	v1243 = v427
	goto L173
L195:
	;
	if int32(0) <= v1253 {
		goto L213
	} else {
		goto L214
	}
L196:
	;
	goto L195
L197:
	;
	v1280 = base.F64_mul(v1276, float64(8.98846567431158e+307))
	if base.Ui32(v1258) < base.Ui32(int32(2047)) {
		goto L200
	} else {
		goto L201
	}
L198:
	;
	goto L199
L199:
	;
	if int32(-1023) < v1258 {
		v1309 = v1276
		v1310 = v1258
		goto L196
	} else {
		goto L206
	}
L200:
	;
	v1309 = v1280
	v1310 = v1258 - int32(1023)
	goto L196
L201:
	;
	goto L202
L202:
	;
	v1287 = int32(3069)
	if base.Ui32(v1287) <= base.Ui32(v1258) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v1290 = v1287
	goto L205
L204:
	;
	v1290 = v1258
	goto L205
L205:
	;
	v1309 = base.F64_mul(v1280, float64(8.98846567431158e+307))
	v1310 = v1290 - int32(2046)
	goto L196
L206:
	;
	v1296 = base.F64_mul(v1276, float64(2.004168360008973e-292))
	if base.Ui32(int32(-1992)) < base.Ui32(v1258) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v1309 = v1296
	v1310 = v1258 + int32(969)
	goto L196
L208:
	;
	goto L209
L209:
	;
	v1303 = int32(-2960)
	if base.Ui32(v1258) <= base.Ui32(v1303) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v1306 = v1303
	goto L212
L211:
	;
	v1306 = v1258
	goto L212
L212:
	;
	v1309 = base.F64_mul(v1296, float64(2.004168360008973e-292))
	v1310 = v1306 + int32(1938)
	goto L196
L213:
	;
	v1320 = base.F64_mul(v1309, base.F64_reinterpret_i64(base.I64_extend_i32_u(v1310+int32(1023))<<(uint(int64(52))%64)))
	v1322 = v1253
	goto L216
L214:
	;
	goto L215
L215:
	;
	v1497 = float64(0)
	if int32(0) <= v1253 {
		goto L232
	} else {
		goto L233
	}
L216:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v324+int32(480)+v1322<<(uint(int32(2))%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v324+v1322<<(uint(int32(3))%32)))) = base.F64_mul(v1320, base.F64_convert_i32_s(v1354))
	if v1322 != 0 {
		v1320 = base.F64_mul(v1320, float64(5.960464477539063e-08))
		v1322 = v1322 - int32(1)
		goto L216
	} else {
		goto L218
	}
L217:
	;
	v1369 = v1253
	goto L219
L218:
	;
	goto L217
L219:
	;
	v1388 = v1253 - v1369
	if v342 < v1388 {
		goto L222
	} else {
		goto L223
	}
L220:
	;
	goto L215
L221:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v324+int32(160)+v1388<<(uint(int32(3))%32)))) = v1435
	if int32(0) < v1369 {
		v1369 = v1369 - int32(1)
		goto L219
	} else {
		goto L231
	}
L222:
	;
	v1390 = v342
	goto L224
L223:
	;
	v1390 = v1388
	goto L224
L224:
	;
	if v1390 < int32(0) {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1435 = float64(0)
	goto L221
L226:
	;
	goto L227
L227:
	;
	v1399 = float64(0)
	v1401 = int32(0)
	goto L228
L228:
	;
	v1426 = v1401 << (uint(int32(3)) % 32)
	v1427 = *(*float64)(unsafe.Add(mBase, uint32(v1426)+uint32(_c_F___rem_pio2[2])))
	v1429 = *(*float64)(unsafe.Add(mBase, uint32(v1426+(v324+v1369<<(uint(int32(3))%32)))))
	v1431 = base.F64_add(base.F64_mul(v1427, v1429), v1399)
	if v1401 != v1390 {
		v1399 = v1431
		v1401 = v1401 + int32(1)
		goto L228
	} else {
		goto L230
	}
L229:
	;
	v1435 = v1431
	goto L221
L230:
	;
	goto L229
L231:
	;
	goto L220
L232:
	;
	v1500 = v1497
	v1502 = v1253
	goto L235
L233:
	;
	v1535 = v1497
	goto L234
L234:
	;
	if v902 != 0 {
		goto L238
	} else {
		goto L239
	}
L235:
	;
	v1533 = *(*float64)(unsafe.Add(mBase, uint32(v324+int32(160)+v1502<<(uint(int32(3))%32))))
	v1534 = base.F64_add(v1500, v1533)
	if v1502 != 0 {
		v1500 = v1534
		v1502 = v1502 - int32(1)
		goto L235
	} else {
		goto L237
	}
L236:
	;
	v1535 = v1534
	goto L234
L237:
	;
	goto L236
L238:
	;
	v1562 = base.F64_neg(v1535)
	goto L240
L239:
	;
	v1562 = v1535
	goto L240
L240:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29))) = v1562
	v1564 = *(*float64)(unsafe.Add(mBase, uint32(v324)+160))
	v1565 = base.F64_sub(v1564, v1535)
	if int32(0) < v1253 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v1569 = v1565
	v1571 = int32(1)
	goto L244
L242:
	;
	v1605 = v1565
	goto L243
L243:
	;
	if v902 != 0 {
		goto L247
	} else {
		goto L248
	}
L244:
	;
	v1600 = *(*float64)(unsafe.Add(mBase, uint32(v324+int32(160)+v1571<<(uint(int32(3))%32))))
	v1601 = base.F64_add(v1569, v1600)
	if v1571 != v1253 {
		v1569 = v1601
		v1571 = v1571 + int32(1)
		goto L244
	} else {
		goto L246
	}
L245:
	;
	v1605 = v1601
	goto L243
L246:
	;
	goto L245
L247:
	;
	v1632 = base.F64_neg(v1605)
	goto L249
L248:
	;
	v1632 = v1605
	goto L249
L249:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29)+8)) = v1632
	m.G0 = v324 + int32(560)
	v1638 = v904 & int32(7)
	v1639 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
	if v31 < int64(0) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_neg(v1639)
	v1644 = *(*float64)(unsafe.Add(mBase, uint32(v29)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_neg(v1644)
	v1655 = int32(0) - v1638
	goto L1
L251:
	;
	goto L252
L252:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v1639
	v1650 = *(*float64)(unsafe.Add(mBase, uint32(v29)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v1650
	v1655 = v1638
	goto L1
}
func F_r_Suffix_Noun_Step2a(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	v2 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3
	v8 = F_find_among_b(m, l0, int32(_a_F_r_Suffix_Noun_Step2a_0), int32(3), v2)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 == int32(0) {
			v107 = v2
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v14
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v17 = int32(0)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v16-int32(4))))
			if v24 == v17 {
				v98 = int32(0)
			} else {
				v29 = v24 & int32(3)
				if base.Ui32(v24) < base.Ui32(int32(4)) {
					v65 = v16
					v66 = int32(0)
					v71 = v65
					v72 = v66
					v76 = v17
					for {
						v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71))))
						v80 = v72 + base.B2i32(int32(-65) < v77)
						v81 = int32(1)
						v84 = v76 + v81
						if v84 != v29 {
							v71 = v71 + v81
							v72 = v80
							v76 = v84
							continue
						} else {
							break
						}
						break
					}
					v87 = v80
				} else {
					v36 = v16
					v37 = int32(0)
					v40 = v17
					for {
						v42 = int32(*(*int8)(unsafe.Add(mBase, uint32(v36))))
						v43 = int32(-65)
						v46 = int32(*(*int8)(unsafe.Add(mBase, uint32(v36)+1)))
						v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(v36)+2)))
						v54 = int32(*(*int8)(unsafe.Add(mBase, uint32(v36)+3)))
						v57 = v37 + base.B2i32(v43 < v42) + base.B2i32(v43 < v46) + base.B2i32(v43 < v50) + base.B2i32(v43 < v54)
						v58 = int32(4)
						v59 = v36 + v58
						v61 = v40 + v58
						if v61 != v24&int32(-4) {
							v36 = v59
							v37 = v57
							v40 = v61
							continue
						} else {
							break
						}
						break
					}
					if v29 == int32(0) {
						v87 = v57
					} else {
						v65 = v59
						v66 = v57
						v71 = v65
						v72 = v66
						v76 = v17
						for {
							v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71))))
							v80 = v72 + base.B2i32(int32(-65) < v77)
							v81 = int32(1)
							v84 = v76 + v81
							if v84 != v29 {
								v71 = v71 + v81
								v72 = v80
								v76 = v84
								continue
							} else {
								break
							}
							break
						}
						v87 = v80
					}
				}
				v98 = v87
			}
			if v98 < int32(5) {
				v107 = v2
			} else {
				v102 = F_slice_del(m, l0)
				mBase = m.M
				if int32(0) <= v102 {
					v105 = int32(1)
				} else {
					v105 = v102
				}
				v107 = v105
			}
		}
		return v107
	}
}
func F_r_et_condition_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7 <= v21 {
		v130 = int32(-1)
		v137 = v130
	} else {
		v38 = int32(1)
		v39 = v7 - v38
		v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22+v39))))
		v43 = v41 & int32(255)
		if base.B2i32(v39 == v21)|base.B2i32(int32(0) <= v41) != 0 {
			v101 = v43
			v105 = v38
		} else {
			v50 = v43 & int32(63)
			v52 = v7 - int32(2)
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v52))))
			v56 = v54 << (uint(int32(6)) % 32)
			if base.B2i32(v52 != v21)&base.B2i32(base.Ui32(v54) < base.Ui32(int32(192))) == int32(0) {
				v101 = v56&int32(1984) | v50
				v105 = int32(2)
			} else {
				v69 = v56&int32(4032) | v50
				v71 = v7 - int32(3)
				v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v71))))
				if base.B2i32(v71 != v21)&base.B2i32(base.Ui32(v73) < base.Ui32(int32(224))) == int32(0) {
					v101 = v73<<(uint(int32(12))%32)&int32(_a_F_r_et_condition_2_0) | v69
					v105 = int32(3)
				} else {
					v91 = int32(4)
					v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+v22-v91))))
					v101 = v73<<(uint(int32(12))%32)&int32(_a_F_r_et_condition_2_1) | v93&int32(7)<<(uint(int32(18))%32) | v69
					v105 = v91
				}
			}
		}
		if int32(246) < v101 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7 - v105
			v130 = int32(0)
			v137 = v130
		} else {
			v107 = v101 - int32(97)
			if v107 < int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7 - v105
				v130 = int32(0)
				v137 = v130
			} else {
				v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v107)>>(uint(int32(3))%32)))+uint32(_c_F_r_et_condition_2[0]))))
				if int32(base.Ui32(v113)>>(uint(v107&int32(7))%32))&int32(1) == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7 - v105
					v130 = int32(0)
					v137 = v130
				} else {
					v137 = v105
				}
			}
		}
	}
	if v137 != 0 {
		v305 = v2
		return v305
	} else {
		v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v150 <= v151 {
			v259 = int32(-1)
			v266 = v259
		} else {
			v168 = int32(1)
			v169 = v150 - v168
			v171 = int32(*(*int8)(unsafe.Add(mBase, uint32(v152+v169))))
			v173 = v171 & int32(255)
			if base.B2i32(v169 == v151)|base.B2i32(int32(0) <= v171) != 0 {
				v231 = v173
				v235 = v168
			} else {
				v180 = v173 & int32(63)
				v182 = v150 - int32(2)
				v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152+v182))))
				v186 = v184 << (uint(int32(6)) % 32)
				if base.B2i32(v182 != v151)&base.B2i32(base.Ui32(v184) < base.Ui32(int32(192))) == int32(0) {
					v231 = v186&int32(1984) | v180
					v235 = int32(2)
				} else {
					v199 = v186&int32(4032) | v180
					v201 = v150 - int32(3)
					v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152+v201))))
					if base.B2i32(v201 != v151)&base.B2i32(base.Ui32(v203) < base.Ui32(int32(224))) == int32(0) {
						v231 = v203<<(uint(int32(12))%32)&int32(_a_F_r_et_condition_2_0) | v199
						v235 = int32(3)
					} else {
						v221 = int32(4)
						v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+v152-v221))))
						v231 = v203<<(uint(int32(12))%32)&int32(_a_F_r_et_condition_2_1) | v223&int32(7)<<(uint(int32(18))%32) | v199
						v235 = v221
					}
				}
			}
			if int32(246) < v231 {
				v266 = v235
			} else {
				v237 = v231 - int32(97)
				if v237 < int32(0) {
					v266 = v235
				} else {
					v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v237)>>(uint(int32(3))%32)))+uint32(_c_F_r_et_condition_2[0]))))
					if int32(base.Ui32(v243)>>(uint(v237&int32(7))%32))&int32(1) == int32(0) {
						v266 = v235
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v150 - v235
						v259 = int32(0)
						v266 = v259
					}
				}
			}
		}
		if v266 != 0 {
			v305 = v2
			return v305
		} else {
			v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v268 <= v267 {
				v305 = v2
				return v305
			} else {
				v270 = v7 - v6
				v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v272 = v270 + v271
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v272
				if v272 <= v267 {
					v301 = v272
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v301
					v305 = int32(1)
					return v305
				} else {
					v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v277 = int32(1)
					v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275+v272-v277))))
					if base.B2i32(v279&int32(224) != int32(96))|base.B2i32(v277<<(uint(v279)%32)&int32(_a_F_r_et_condition_2_2) == int32(0)) != 0 {
						v301 = v272
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v301
						v305 = int32(1)
						return v305
					} else {
						v291 = int32(0)
						v295 = F_find_among_b(m, l0, int32(_a_F_r_et_condition_2_3), int32(21), v291)
						mBase = m.M
						v298 = m.ExcPending
						if v298 != 0 {
							return int32(0)
						} else {
							if v295 != 0 {
								v305 = v291
							} else {
								v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v301 = v299 + v270
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v301
								v305 = int32(1)
							}
							return v305
						}
					}
				}
			}
		}
	}
}
func F_r_fix_va_start(m *base.Module, l0 int32) int32 {
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v4
	v7 = v4 + int32(5)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v8 <= v7 {
		v64 = v2
		return v64
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v7))))
		if base.B2i32(v12&int32(224) != int32(128))|base.B2i32(int32(1)<<(uint(v12)%32)&int32(3078) == int32(0)) != 0 {
			v64 = v2
			return v64
		} else {
			v27 = F_find_among(m, l0, int32(_a_F_r_fix_va_start_0), int32(4), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				if v27 == int32(0) {
					v64 = v2
					return v64
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v33
					switch v27 - int32(1) {
					case 0:
						v39 = F_slice_from_s(m, l0, int32(3), int32(_a_F_r_fix_va_start_1))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if int32(0) <= v39 {
								v64 = int32(1)
							} else {
								v64 = v39
							}
							return v64
						}
					case 1:
						v45 = F_slice_from_s(m, l0, int32(3), int32(_a_F_r_fix_va_start_2))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							if int32(0) <= v45 {
								v64 = int32(1)
							} else {
								v64 = v45
							}
							return v64
						}
					case 2:
						v51 = F_slice_from_s(m, l0, int32(3), int32(_a_F_r_fix_va_start_3))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							if int32(0) <= v51 {
								v64 = int32(1)
							} else {
								v64 = v51
							}
							return v64
						}
					case 3:
						v57 = F_slice_from_s(m, l0, int32(3), int32(_a_F_r_fix_va_start_4))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							if v57 < int32(0) {
								v64 = v57
							} else {
								v64 = int32(1)
							}
							return v64
						}
					default:
						v64 = int32(1)
						return v64
					}
				}
			}
		}
	}
}
func F_r_lengthen_V_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
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
	var v135 int32
	_ = v135
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
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v270 int32
	_ = v270
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v334 int32
	_ = v334
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 <= v17 {
		v57 = int32(-1)
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v8-int32(1)))))
		if int32(252) < v32 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8 - int32(1)
			v54 = int32(0)
		} else {
			v34 = v32 - int32(97)
			if v34 < int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8 - int32(1)
				v54 = int32(0)
			} else {
				v37 = int32(1)
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v34)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_1[0]))))
				if int32(base.Ui32(v41)>>(uint(v34&int32(7))%32))&v37 != 0 {
					v54 = v37
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8 - int32(1)
					v54 = int32(0)
				}
			}
		}
		v57 = v54
	}
	if v57 != 0 {
		v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
		v465 = int32(1)
		return v465
	} else {
		v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v58
		v63 = F_find_among_b(m, l0, int32(_a_F_r_lengthen_V_1_0), int32(21), int32(0))
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return int32(0)
		} else {
			if v63 == int32(0) {
				v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
				v465 = int32(1)
				return v465
			} else {
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v69
				switch v63 - int32(1) {
				case 0:
					v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v81 <= v82 {
						v122 = int32(-1)
					} else {
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v81-int32(1)))))
						if int32(252) < v97 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v81 - int32(1)
							v119 = int32(0)
						} else {
							v99 = v97 - int32(97)
							if v99 < int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v81 - int32(1)
								v119 = int32(0)
							} else {
								v102 = int32(1)
								v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v99)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_1[1]))))
								if int32(base.Ui32(v106)>>(uint(v99&int32(7))%32))&v102 != 0 {
									v119 = v102
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v81 - int32(1)
									v119 = int32(0)
								}
							}
						}
						v122 = v119
					}
					v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v125 = v123 + (v69 - v73)
					if v122 != 0 {
						v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v126 < v125 {
							v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
							v465 = int32(1)
							return v465
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v125
							v131 = F_slice_to(m, l0, l0+int32(40))
							mBase = m.M
							v132 = m.ExcPending
							if v132 != 0 {
								return int32(0)
							} else {
								if v131 < int32(0) {
									v465 = v131
									return v465
								} else {
									v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									v137 = F_insert_v(m, l0, v135, v135, v136)
									mBase = m.M
									v138 = m.ExcPending
									if v138 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v135
										if int32(0) <= v137 {
											v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
											v465 = int32(1)
										} else {
											v465 = v137
										}
										return v465
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v125
						v131 = F_slice_to(m, l0, l0+int32(40))
						mBase = m.M
						v132 = m.ExcPending
						if v132 != 0 {
							return int32(0)
						} else {
							if v131 < int32(0) {
								v465 = v131
								return v465
							} else {
								v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								v137 = F_insert_v(m, l0, v135, v135, v136)
								mBase = m.M
								v138 = m.ExcPending
								if v138 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v135
									if int32(0) <= v137 {
										v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
										v465 = int32(1)
									} else {
										v465 = v137
									}
									return v465
								}
							}
						}
					}
				case 1:
					v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v143 = v142 - v69
					v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v151 <= v152 {
						v192 = int32(-1)
					} else {
						v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v151-int32(1)))))
						if int32(252) < v167 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v151 - int32(1)
							v189 = int32(0)
						} else {
							v169 = v167 - int32(97)
							if v169 < int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v151 - int32(1)
								v189 = int32(0)
							} else {
								v172 = int32(1)
								v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v169)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_1[1]))))
								if int32(base.Ui32(v176)>>(uint(v169&int32(7))%32))&v172 != 0 {
									v189 = v172
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v151 - int32(1)
									v189 = int32(0)
								}
							}
						}
						v192 = v189
					}
					v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v192 == int32(0) {
						v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v201 = v196
						v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v210 <= v211 {
							v254 = int32(-1)
						} else {
							v223 = int32(1)
							v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224+v210-v223))))
							if int32(252) < v228 {
								v250 = v223
							} else {
								v230 = v228 - int32(97)
								if v230 < int32(0) {
									v250 = v223
								} else {
									v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v230)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_1[2]))))
									if int32(base.Ui32(v236)>>(uint(v230&int32(7))%32))&int32(1) == int32(0) {
										v250 = v223
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v210 - int32(1)
										v250 = int32(0)
									}
								}
							}
							v254 = v250
						}
						if v254 == int32(0) {
							v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
							v465 = int32(1)
							return v465
						} else {
							v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v258 = v193 - v201
							v259 = v257 - v258
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v259
							v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v259 <= v270 {
								v313 = int32(-1)
							} else {
								v282 = int32(1)
								v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283+v259-v282))))
								if int32(235) < v287 {
									v309 = v282
								} else {
									v289 = v287 - int32(101)
									if v289 < int32(0) {
										v309 = v282
									} else {
										v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v289)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_1[3]))))
										if int32(base.Ui32(v295)>>(uint(v289&int32(7))%32))&int32(1) == int32(0) {
											v309 = v282
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v259 - int32(1)
											v309 = int32(0)
										}
									}
								}
								v313 = v309
							}
							if v313 != 0 {
								v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v318 = v314
								v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v320 = v319 - v258
								if v320 <= v318 {
									v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v429 - v143
									v434 = F_slice_to(m, l0, l0+int32(40))
									mBase = m.M
									v435 = m.ExcPending
									if v435 != 0 {
										return int32(0)
									} else {
										if v434 < int32(0) {
											v465 = v434
											return v465
										} else {
											v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											v440 = F_insert_v(m, l0, v438, v438, v439)
											mBase = m.M
											v441 = m.ExcPending
											if v441 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v438
												if int32(0) <= v440 {
													v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
													v465 = int32(1)
												} else {
													v465 = v440
												}
												return v465
											}
										}
									}
								} else {
									v323 = v320 - int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323
									v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v323 <= v334 {
										v377 = int32(-1)
									} else {
										v346 = int32(1)
										v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347+v323-v346))))
										if int32(252) < v351 {
											v373 = v346
										} else {
											v353 = v351 - int32(97)
											if v353 < int32(0) {
												v373 = v346
											} else {
												v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v353)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_1[2]))))
												if int32(base.Ui32(v359)>>(uint(v353&int32(7))%32))&int32(1) == int32(0) {
													v373 = v346
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - int32(1)
													v373 = int32(0)
												}
											}
										}
										v377 = v373
									}
									if v377 != 0 {
										v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v429 - v143
										v434 = F_slice_to(m, l0, l0+int32(40))
										mBase = m.M
										v435 = m.ExcPending
										if v435 != 0 {
											return int32(0)
										} else {
											if v434 < int32(0) {
												v465 = v434
												return v465
											} else {
												v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
												v440 = F_insert_v(m, l0, v438, v438, v439)
												mBase = m.M
												v441 = m.ExcPending
												if v441 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v438
													if int32(0) <= v440 {
														v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
														v465 = int32(1)
													} else {
														v465 = v440
													}
													return v465
												}
											}
										}
									} else {
										v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v385 <= v386 {
											v426 = int32(-1)
										} else {
											v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397+v385-int32(1)))))
											if int32(252) < v401 {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385 - int32(1)
												v423 = int32(0)
											} else {
												v403 = v401 - int32(97)
												if v403 < int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385 - int32(1)
													v423 = int32(0)
												} else {
													v406 = int32(1)
													v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v403)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_1[1]))))
													if int32(base.Ui32(v410)>>(uint(v403&int32(7))%32))&v406 != 0 {
														v423 = v406
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385 - int32(1)
														v423 = int32(0)
													}
												}
											}
											v426 = v423
										}
										if v426 == int32(0) {
											v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
											v465 = int32(1)
											return v465
										} else {
											v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v429 - v143
											v434 = F_slice_to(m, l0, l0+int32(40))
											mBase = m.M
											v435 = m.ExcPending
											if v435 != 0 {
												return int32(0)
											} else {
												if v434 < int32(0) {
													v465 = v434
													return v465
												} else {
													v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
													v440 = F_insert_v(m, l0, v438, v438, v439)
													mBase = m.M
													v441 = m.ExcPending
													if v441 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v438
														if int32(0) <= v440 {
															v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
															v465 = int32(1)
														} else {
															v465 = v440
														}
														return v465
													}
												}
											}
										}
									}
								}
							} else {
								v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v316 <= v315 {
									v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
									v465 = int32(1)
									return v465
								} else {
									v318 = v315
									v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v320 = v319 - v258
									if v320 <= v318 {
										v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v429 - v143
										v434 = F_slice_to(m, l0, l0+int32(40))
										mBase = m.M
										v435 = m.ExcPending
										if v435 != 0 {
											return int32(0)
										} else {
											if v434 < int32(0) {
												v465 = v434
												return v465
											} else {
												v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
												v440 = F_insert_v(m, l0, v438, v438, v439)
												mBase = m.M
												v441 = m.ExcPending
												if v441 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v438
													if int32(0) <= v440 {
														v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
														v465 = int32(1)
													} else {
														v465 = v440
													}
													return v465
												}
											}
										}
									} else {
										v323 = v320 - int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323
										v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v323 <= v334 {
											v377 = int32(-1)
										} else {
											v346 = int32(1)
											v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347+v323-v346))))
											if int32(252) < v351 {
												v373 = v346
											} else {
												v353 = v351 - int32(97)
												if v353 < int32(0) {
													v373 = v346
												} else {
													v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v353)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_1[2]))))
													if int32(base.Ui32(v359)>>(uint(v353&int32(7))%32))&int32(1) == int32(0) {
														v373 = v346
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - int32(1)
														v373 = int32(0)
													}
												}
											}
											v377 = v373
										}
										if v377 != 0 {
											v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v429 - v143
											v434 = F_slice_to(m, l0, l0+int32(40))
											mBase = m.M
											v435 = m.ExcPending
											if v435 != 0 {
												return int32(0)
											} else {
												if v434 < int32(0) {
													v465 = v434
													return v465
												} else {
													v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
													v440 = F_insert_v(m, l0, v438, v438, v439)
													mBase = m.M
													v441 = m.ExcPending
													if v441 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v438
														if int32(0) <= v440 {
															v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
															v465 = int32(1)
														} else {
															v465 = v440
														}
														return v465
													}
												}
											}
										} else {
											v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v385 <= v386 {
												v426 = int32(-1)
											} else {
												v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397+v385-int32(1)))))
												if int32(252) < v401 {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385 - int32(1)
													v423 = int32(0)
												} else {
													v403 = v401 - int32(97)
													if v403 < int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385 - int32(1)
														v423 = int32(0)
													} else {
														v406 = int32(1)
														v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v403)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_1[1]))))
														if int32(base.Ui32(v410)>>(uint(v403&int32(7))%32))&v406 != 0 {
															v423 = v406
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385 - int32(1)
															v423 = int32(0)
														}
													}
												}
												v426 = v423
											}
											if v426 == int32(0) {
												v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
												v465 = int32(1)
												return v465
											} else {
												v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v429 - v143
												v434 = F_slice_to(m, l0, l0+int32(40))
												mBase = m.M
												v435 = m.ExcPending
												if v435 != 0 {
													return int32(0)
												} else {
													if v434 < int32(0) {
														v465 = v434
														return v465
													} else {
														v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
														v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
														v440 = F_insert_v(m, l0, v438, v438, v439)
														mBase = m.M
														v441 = m.ExcPending
														if v441 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v438
															if int32(0) <= v440 {
																v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
																v465 = int32(1)
															} else {
																v465 = v440
															}
															return v465
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
						v197 = v193 - v143
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v197
						v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v199 < v197 {
							v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
							v465 = int32(1)
							return v465
						} else {
							v201 = v197
							v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v210 <= v211 {
								v254 = int32(-1)
							} else {
								v223 = int32(1)
								v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224+v210-v223))))
								if int32(252) < v228 {
									v250 = v223
								} else {
									v230 = v228 - int32(97)
									if v230 < int32(0) {
										v250 = v223
									} else {
										v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v230)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_1[2]))))
										if int32(base.Ui32(v236)>>(uint(v230&int32(7))%32))&int32(1) == int32(0) {
											v250 = v223
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v210 - int32(1)
											v250 = int32(0)
										}
									}
								}
								v254 = v250
							}
							if v254 == int32(0) {
								v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
								v465 = int32(1)
								return v465
							} else {
								v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v258 = v193 - v201
								v259 = v257 - v258
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v259
								v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v259 <= v270 {
									v313 = int32(-1)
								} else {
									v282 = int32(1)
									v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283+v259-v282))))
									if int32(235) < v287 {
										v309 = v282
									} else {
										v289 = v287 - int32(101)
										if v289 < int32(0) {
											v309 = v282
										} else {
											v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v289)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_1[3]))))
											if int32(base.Ui32(v295)>>(uint(v289&int32(7))%32))&int32(1) == int32(0) {
												v309 = v282
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v259 - int32(1)
												v309 = int32(0)
											}
										}
									}
									v313 = v309
								}
								if v313 != 0 {
									v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v318 = v314
									v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v320 = v319 - v258
									if v320 <= v318 {
										v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v429 - v143
										v434 = F_slice_to(m, l0, l0+int32(40))
										mBase = m.M
										v435 = m.ExcPending
										if v435 != 0 {
											return int32(0)
										} else {
											if v434 < int32(0) {
												v465 = v434
												return v465
											} else {
												v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
												v440 = F_insert_v(m, l0, v438, v438, v439)
												mBase = m.M
												v441 = m.ExcPending
												if v441 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v438
													if int32(0) <= v440 {
														v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
														v465 = int32(1)
													} else {
														v465 = v440
													}
													return v465
												}
											}
										}
									} else {
										v323 = v320 - int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323
										v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v323 <= v334 {
											v377 = int32(-1)
										} else {
											v346 = int32(1)
											v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347+v323-v346))))
											if int32(252) < v351 {
												v373 = v346
											} else {
												v353 = v351 - int32(97)
												if v353 < int32(0) {
													v373 = v346
												} else {
													v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v353)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_1[2]))))
													if int32(base.Ui32(v359)>>(uint(v353&int32(7))%32))&int32(1) == int32(0) {
														v373 = v346
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - int32(1)
														v373 = int32(0)
													}
												}
											}
											v377 = v373
										}
										if v377 != 0 {
											v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v429 - v143
											v434 = F_slice_to(m, l0, l0+int32(40))
											mBase = m.M
											v435 = m.ExcPending
											if v435 != 0 {
												return int32(0)
											} else {
												if v434 < int32(0) {
													v465 = v434
													return v465
												} else {
													v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
													v440 = F_insert_v(m, l0, v438, v438, v439)
													mBase = m.M
													v441 = m.ExcPending
													if v441 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v438
														if int32(0) <= v440 {
															v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
															v465 = int32(1)
														} else {
															v465 = v440
														}
														return v465
													}
												}
											}
										} else {
											v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v385 <= v386 {
												v426 = int32(-1)
											} else {
												v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397+v385-int32(1)))))
												if int32(252) < v401 {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385 - int32(1)
													v423 = int32(0)
												} else {
													v403 = v401 - int32(97)
													if v403 < int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385 - int32(1)
														v423 = int32(0)
													} else {
														v406 = int32(1)
														v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v403)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_1[1]))))
														if int32(base.Ui32(v410)>>(uint(v403&int32(7))%32))&v406 != 0 {
															v423 = v406
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385 - int32(1)
															v423 = int32(0)
														}
													}
												}
												v426 = v423
											}
											if v426 == int32(0) {
												v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
												v465 = int32(1)
												return v465
											} else {
												v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v429 - v143
												v434 = F_slice_to(m, l0, l0+int32(40))
												mBase = m.M
												v435 = m.ExcPending
												if v435 != 0 {
													return int32(0)
												} else {
													if v434 < int32(0) {
														v465 = v434
														return v465
													} else {
														v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
														v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
														v440 = F_insert_v(m, l0, v438, v438, v439)
														mBase = m.M
														v441 = m.ExcPending
														if v441 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v438
															if int32(0) <= v440 {
																v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
																v465 = int32(1)
															} else {
																v465 = v440
															}
															return v465
														}
													}
												}
											}
										}
									}
								} else {
									v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v316 <= v315 {
										v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
										v465 = int32(1)
										return v465
									} else {
										v318 = v315
										v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v320 = v319 - v258
										if v320 <= v318 {
											v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v429 - v143
											v434 = F_slice_to(m, l0, l0+int32(40))
											mBase = m.M
											v435 = m.ExcPending
											if v435 != 0 {
												return int32(0)
											} else {
												if v434 < int32(0) {
													v465 = v434
													return v465
												} else {
													v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
													v440 = F_insert_v(m, l0, v438, v438, v439)
													mBase = m.M
													v441 = m.ExcPending
													if v441 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v438
														if int32(0) <= v440 {
															v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
															v465 = int32(1)
														} else {
															v465 = v440
														}
														return v465
													}
												}
											}
										} else {
											v323 = v320 - int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323
											v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v323 <= v334 {
												v377 = int32(-1)
											} else {
												v346 = int32(1)
												v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347+v323-v346))))
												if int32(252) < v351 {
													v373 = v346
												} else {
													v353 = v351 - int32(97)
													if v353 < int32(0) {
														v373 = v346
													} else {
														v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v353)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_1[2]))))
														if int32(base.Ui32(v359)>>(uint(v353&int32(7))%32))&int32(1) == int32(0) {
															v373 = v346
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v323 - int32(1)
															v373 = int32(0)
														}
													}
												}
												v377 = v373
											}
											if v377 != 0 {
												v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v429 - v143
												v434 = F_slice_to(m, l0, l0+int32(40))
												mBase = m.M
												v435 = m.ExcPending
												if v435 != 0 {
													return int32(0)
												} else {
													if v434 < int32(0) {
														v465 = v434
														return v465
													} else {
														v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
														v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
														v440 = F_insert_v(m, l0, v438, v438, v439)
														mBase = m.M
														v441 = m.ExcPending
														if v441 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v438
															if int32(0) <= v440 {
																v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
																v465 = int32(1)
															} else {
																v465 = v440
															}
															return v465
														}
													}
												}
											} else {
												v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v385 <= v386 {
													v426 = int32(-1)
												} else {
													v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397+v385-int32(1)))))
													if int32(252) < v401 {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385 - int32(1)
														v423 = int32(0)
													} else {
														v403 = v401 - int32(97)
														if v403 < int32(0) {
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385 - int32(1)
															v423 = int32(0)
														} else {
															v406 = int32(1)
															v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v403)>>(uint(int32(3))%32)))+uint32(_c_F_r_lengthen_V_1[1]))))
															if int32(base.Ui32(v410)>>(uint(v403&int32(7))%32))&v406 != 0 {
																v423 = v406
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385 - int32(1)
																v423 = int32(0)
															}
														}
													}
													v426 = v423
												}
												if v426 == int32(0) {
													v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
													v465 = int32(1)
													return v465
												} else {
													v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v429 - v143
													v434 = F_slice_to(m, l0, l0+int32(40))
													mBase = m.M
													v435 = m.ExcPending
													if v435 != 0 {
														return int32(0)
													} else {
														if v434 < int32(0) {
															v465 = v434
															return v465
														} else {
															v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
															v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
															v440 = F_insert_v(m, l0, v438, v438, v439)
															mBase = m.M
															v441 = m.ExcPending
															if v441 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v438
																if int32(0) <= v440 {
																	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
																	v465 = int32(1)
																} else {
																	v465 = v440
																}
																return v465
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				case 2:
					v447 = F_slice_from_s(m, l0, int32(3), int32(_a_F_r_lengthen_V_1_1))
					mBase = m.M
					v448 = m.ExcPending
					if v448 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v447 {
							v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
							v465 = int32(1)
						} else {
							v465 = v447
						}
						return v465
					}
				case 3:
					v453 = F_slice_from_s(m, l0, int32(3), int32(_a_F_r_lengthen_V_1_2))
					mBase = m.M
					v454 = m.ExcPending
					if v454 != 0 {
						return int32(0)
					} else {
						if v453 < int32(0) {
							v465 = v453
						} else {
							v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
							v465 = int32(1)
						}
						return v465
					}
				default:
					v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + (v8 - v7)
					v465 = int32(1)
					return v465
				}
			}
		}
	}
}
func F_r_measure_1(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v5
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	for {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v21 < v20 {
			v23 = v20
		} else {
			v23 = v21
		}
		if v20 == v23 {
			v61 = int32(-1)
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+v20))))
			if int32(252) < v36 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20 + int32(1)
				v58 = int32(0)
			} else {
				v38 = v36 - int32(97)
				if v38 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20 + int32(1)
					v58 = int32(0)
				} else {
					v41 = int32(1)
					v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v38)>>(uint(int32(3))%32)))+uint32(_c_F_r_measure_1[0]))))
					if int32(base.Ui32(v45)>>(uint(v38&int32(7))%32))&v41 != 0 {
						v58 = v41
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20 + int32(1)
						v58 = int32(0)
					}
				}
			}
			v61 = v58
		}
		if v61 == int32(0) {
			continue
		} else {
			break
		}
		break
	}
	v66 = int32(1)
	for {
		v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v70 = int32(2)
		v72 = int32(0)
		v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v74-v69 < v70 {
			v84 = v72
		} else {
			v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v80 = F_memcmp(m, v78+v69, int32(_a_F_r_measure_1_0), v70)
			mBase = m.M
			if v80 != 0 {
				v84 = v72
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v70 + v69
				v84 = int32(1)
			}
		}
		if v84 == int32(0) {
		} else {
			v66 = v66 - int32(1)
			continue
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
		v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v97 < v69 {
			v99 = v69
		} else {
			v99 = v97
		}
		if v69 == v99 {
			v140 = int32(-1)
		} else {
			v111 = int32(1)
			v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+v69))))
			if int32(252) < v114 {
				v135 = v111
			} else {
				v116 = v114 - int32(97)
				if v116 < int32(0) {
					v135 = v111
				} else {
					v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v116)>>(uint(int32(3))%32)))+uint32(_c_F_r_measure_1[0]))))
					if int32(base.Ui32(v122)>>(uint(v116&int32(7))%32))&int32(1) == int32(0) {
						v135 = v111
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69 + int32(1)
						v135 = int32(0)
					}
				}
			}
			v140 = v135
		}
		if v140 != 0 {
			break
		} else {
			v66 = v66 - int32(1)
			continue
		}
		break
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
	if int32(0) < v66 {
	} else {
		v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v154 < v153 {
			v156 = v153
		} else {
			v156 = v154
		}
		if v153 == v156 {
			v194 = int32(-1)
		} else {
			v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167+v153))))
			if int32(252) < v169 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v153 + int32(1)
				v191 = int32(0)
			} else {
				v171 = v169 - int32(97)
				if v171 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v153 + int32(1)
					v191 = int32(0)
				} else {
					v174 = int32(1)
					v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v171)>>(uint(int32(3))%32)))+uint32(_c_F_r_measure_1[0]))))
					if int32(base.Ui32(v178)>>(uint(v171&int32(7))%32))&v174 != 0 {
						v191 = v174
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v153 + int32(1)
						v191 = int32(0)
					}
				}
			}
			v194 = v191
		}
		if v194 != 0 {
		} else {
			v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v195
			for {
				v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v209 < v208 {
					v211 = v208
				} else {
					v211 = v209
				}
				if v208 == v211 {
					v249 = int32(-1)
				} else {
					v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222+v208))))
					if int32(252) < v224 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v208 + int32(1)
						v246 = int32(0)
					} else {
						v226 = v224 - int32(97)
						if v226 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v208 + int32(1)
							v246 = int32(0)
						} else {
							v229 = int32(1)
							v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v226)>>(uint(int32(3))%32)))+uint32(_c_F_r_measure_1[0]))))
							if int32(base.Ui32(v233)>>(uint(v226&int32(7))%32))&v229 != 0 {
								v246 = v229
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v208 + int32(1)
								v246 = int32(0)
							}
						}
					}
					v249 = v246
				}
				if v249 == int32(0) {
					continue
				} else {
					break
				}
				break
			}
			v254 = int32(1)
			for {
				v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v258 = int32(2)
				v260 = int32(0)
				v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v262-v257 < v258 {
					v272 = v260
				} else {
					v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v268 = F_memcmp(m, v266+v257, int32(_a_F_r_measure_1_1), v258)
					mBase = m.M
					if v268 != 0 {
						v272 = v260
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v258 + v257
						v272 = int32(1)
					}
				}
				if v272 == int32(0) {
				} else {
					v254 = v254 - int32(1)
					continue
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v257
				v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v285 < v257 {
					v287 = v257
				} else {
					v287 = v285
				}
				if v257 == v287 {
					v328 = int32(-1)
				} else {
					v299 = int32(1)
					v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300+v257))))
					if int32(252) < v302 {
						v323 = v299
					} else {
						v304 = v302 - int32(97)
						if v304 < int32(0) {
							v323 = v299
						} else {
							v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v304)>>(uint(int32(3))%32)))+uint32(_c_F_r_measure_1[0]))))
							if int32(base.Ui32(v310)>>(uint(v304&int32(7))%32))&int32(1) == int32(0) {
								v323 = v299
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v257 + int32(1)
								v323 = int32(0)
							}
						}
					}
					v328 = v323
				}
				if v328 != 0 {
					break
				} else {
					v254 = v254 - int32(1)
					continue
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v257
			if int32(0) < v254 {
			} else {
				v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v342 < v341 {
					v344 = v341
				} else {
					v344 = v342
				}
				if v341 == v344 {
					v382 = int32(-1)
				} else {
					v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355+v341))))
					if int32(252) < v357 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v341 + int32(1)
						v379 = int32(0)
					} else {
						v359 = v357 - int32(97)
						if v359 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v341 + int32(1)
							v379 = int32(0)
						} else {
							v362 = int32(1)
							v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v359)>>(uint(int32(3))%32)))+uint32(_c_F_r_measure_1[0]))))
							if int32(base.Ui32(v366)>>(uint(v359&int32(7))%32))&v362 != 0 {
								v379 = v362
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v341 + int32(1)
								v379 = int32(0)
							}
						}
					}
					v382 = v379
				}
				if v382 != 0 {
				} else {
					v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v383
				}
			}
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8
	return
}
func F_rangesel(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 float64
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int64
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
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
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
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
	var v126 float64
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 float32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 float32
	_ = v144
	var v146 int32
	_ = v146
	var v148 float64
	_ = v148
	var v151 float64
	_ = v151
	var v152 float64
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
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
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int64
	_ = v285
	var v302 int32
	_ = v302
	var v308 float64
	_ = v308
	var v309 int32
	_ = v309
	var v314 float64
	_ = v314
	var v315 int32
	_ = v315
	var v321 float64
	_ = v321
	var v322 int32
	_ = v322
	var v327 float64
	_ = v327
	var v328 int32
	_ = v328
	var v333 float64
	_ = v333
	var v334 int32
	_ = v334
	var v340 float64
	_ = v340
	var v341 int32
	_ = v341
	var v346 float64
	_ = v346
	var v347 int32
	_ = v347
	var v352 float64
	_ = v352
	var v353 int32
	_ = v353
	var v358 float64
	_ = v358
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 float64
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v377 float64
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v386 float64
	_ = v386
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 float64
	_ = v395
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v413 float64
	_ = v413
	var v414 int32
	_ = v414
	var v416 float64
	_ = v416
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v446 int32
	_ = v446
	var v465 int32
	_ = v465
	var v470 float64
	_ = v470
	var v480 float64
	_ = v480
	var v488 float64
	_ = v488
	var v503 float64
	_ = v503
	var v510 float64
	_ = v510
	var v513 float64
	_ = v513
	var v527 int32
	_ = v527
	var v532 float64
	_ = v532
	var v543 float64
	_ = v543
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 float64
	_ = v553
	var v556 float64
	_ = v556
	var v560 int64
	_ = v560
	var v563 int64
	_ = v563
	var v577 int64
	_ = v577
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	v14 = int64(0)
	v17 = m.G0
	v19 = v17 - int32(176)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v31 = F_get_restriction_variable(m, v22, v23, v24, v19+int32(28), v19+int32(24), v19+int32(23))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L5
	} else {
		goto L153
	}
L2:
	;
	m.G0 = v19 + int32(176)
	return v577
L3:
	;
	v563 = *(*int64)(unsafe.Add(mBase, uint32(v51<<(uint(int32(3))%32))+uint32(_c_F_rangesel[0])))
	v577 = v563
	goto L2
L4:
	;
	v560 = *(*int64)(unsafe.Add(mBase, uint32(v38<<(uint(int32(3))%32))+uint32(_c_F_rangesel[0])))
	v577 = v560
	goto L2
L5:
	;
	return int64(0)
L6:
	;
	if v31 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v38 = v21 - int32(3884)
	if base.Ui32(v38) < base.Ui32(int32(13)) {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v43 != int32(7) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v577 = int64(4576918229304087675)
	goto L2
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	if v46 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+32)))
	if v55 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	m.T0[v47].(func(*base.Module, int32))(m, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v51 = v21 - int32(3884)
	if base.Ui32(v51) < base.Ui32(int32(13)) {
		goto L3
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v577 = int64(4576918229304087675)
	goto L2
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	if v58 == int32(0) {
		v577 = v14
		goto L2
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+23)))
	if v64 != 0 {
		v74 = v21
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	m.T0[v61].(func(*base.Module, int32))(m, v58)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v577 = v14
	goto L2
L24:
	;
	v75 = float64(0.005)
	switch v74 - int32(3889) {
	case 0:
		goto L34
	default:
		goto L33
	case 2:
		v543 = v75
		goto L30
	}
L25:
	;
	v65 = F_get_commutator(m, v21)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	if v65 != 0 {
		v74 = v65
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v67 = int64(4576918229304087675)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	if v68 == int32(0) {
		v577 = v67
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	m.T0[v71].(func(*base.Module, int32))(m, v68)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v577 = v67
	goto L2
L30:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	if v549 != 0 {
		goto L146
	} else {
		goto L147
	}
L31:
	;
	v527 = v74 - int32(3884)
	if base.Ui32(int32(12)) < base.Ui32(v527) {
		v543 = float64(0.01)
		goto L30
	} else {
		goto L145
	}
L32:
	;
	if v120 == int32(0) {
		goto L31
	} else {
		goto L41
	}
L33:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if v108 != v109 {
		goto L31
	} else {
		goto L38
	}
L34:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v79 = F_range_get_typcache(m, l0, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+200))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	if v82 != v84 {
		v543 = v75
		goto L30
	} else {
		goto L36
	}
L36:
	;
	v86 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+145)) = uint8(v86)
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v81)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+146)) = uint8(v86)
	v91 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+144)) = uint8(v91)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+136)) = v88
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+106)) = uint8(v91)
	v96 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+104)) = uint16(v96)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = v88
	v105 = F_range_serialize(m, v79, v19+int32(136), v19+int32(96), v91, v91)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v117 = v79
	v120 = v105
	goto L32
L38:
	;
	v111 = F_range_get_typcache(m, l0, v108)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+24))
	v115 = F_pg_detoast_datum(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v117 = v111
	v120 = v115
	goto L32
L41:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	if v123 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v159 = int32(*(*int8)(unsafe.Add(mBase, uint32(v120+int32(base.Ui32(v153)>>(uint(int32(2))%32))-int32(1)))))
	goto L53
L43:
	;
	v126 = float64(0)
	v151 = v126
	v152 = v126
	goto L42
L44:
	;
	goto L45
L45:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+22)))
	v131 = *(*float32)(unsafe.Add(mBase, uint32(v128+v129)+8))
	v134 = v19 + int32(136)
	v138 = F_get_attstatsslot(m, v134, v123, int32(6), int32(0), int32(2))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	if v138 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v19)+160))
	if v140 != int32(1) {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	v148 = float64(0)
	goto L49
L49:
	;
	v151 = v148
	v152 = base.F64_promote_f32(v131)
	goto L42
L50:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v19)+156))
	v144 = *(*float32)(unsafe.Add(mBase, uint32(v143)))
	F_free_attstatsslot(m, v134)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	v148 = base.F64_promote_f32(v144)
	goto L49
L52:
	;
	v510 = float64(0)
	v513 = base.F64_mul(base.F64_sub(float64(1), v152), v503)
	if base.F64_lt(v513, v510) != 0 {
		v543 = v510
		goto L30
	} else {
		goto L143
	}
L53:
	;
	if v159&int32(1) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	switch v74 - int32(3884) {
	case 0, 4, 9, 10, 11, 12:
		v503 = float64(0)
		goto L52
	case 1, 8:
		goto L60
	case 2, 6:
		goto L59
	case 3:
		goto L58
	default:
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v183 = v19 + int32(28)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v117)+216))
	v185 = F_statistic_proc_security_check(m, v183, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L5
	} else {
		goto L66
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L5
	} else {
		goto L61
	}
L58:
	;
	v503 = base.F64_sub(float64(1), v151)
	goto L52
L59:
	;
	v503 = float64(1)
	goto L52
L60:
	;
	v503 = v151
	goto L52
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v74
	F_errmsg_internal(m, int32(_a_F_rangesel_0), v19+int32(16))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_rangesel_1), int32(320), int32(_a_F_rangesel_2))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L5
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
	v488 = base.F64_sub(float64(1), v151)
	if v74 == int32(3892) {
		goto L140
	} else {
		goto L141
	}
L65:
	;
	v465 = v74 - int32(3884)
	if base.Ui32(int32(12)) < base.Ui32(v465) {
		v480 = float64(0.01)
		goto L64
	} else {
		goto L139
	}
L66:
	;
	if v185 == int32(0) {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v117)+272))
	if v189 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v190 = F_statistic_proc_security_check(m, v183, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	if v194 == int32(0) {
		goto L65
	} else {
		goto L73
	}
L71:
	;
	if v190 == int32(0) {
		goto L65
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v202 = F_get_attstatsslot(m, v19+int32(136), v194, int32(7), int32(0), int32(1))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	if v202 == int32(0) {
		goto L65
	} else {
		goto L75
	}
L75:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v19)+152))
	if v206 < int32(2) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_free_attstatsslot(m, v19+int32(136))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L5
	} else {
		goto L138
	}
L77:
	;
	v210 = F_palloc_mul(m, int32(16), v206)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	v213 = F_palloc_mul(m, int32(16), v206)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	v220 = int32(0)
	goto L81
L80:
	;
	switch v74 - int32(3890) {
	case 0, 2:
		goto L94
	default:
		goto L93
	}
L81:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v19)+148))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231+v220<<(uint(int32(3))%32))))
	v236 = F_pg_detoast_datum(m, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L5
	} else {
		goto L84
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L5
	} else {
		goto L88
	}
L83:
	;
	goto L82
L84:
	;
	v239 = v220 << (uint(int32(4)) % 32)
	F_range_deserialize(m, v117, v236, v210+v239, v213+v239, v19+int32(63))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+63)))
	if v246 == int32(1) {
		goto L83
	} else {
		goto L86
	}
L86:
	;
	v250 = v220 + int32(1)
	if v206 != v250 {
		v220 = v250
		goto L81
	} else {
		goto L87
	}
L87:
	;
	goto L80
L88:
	;
	F_errmsg_internal(m, int32(_a_F_rangesel_3), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_rangesel_1), int32(423), int32(_a_F_rangesel_4))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L5
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	F_free_attstatsslot(m, v19+int32(96))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L5
	} else {
		goto L135
	}
L92:
	;
	F_range_deserialize(m, v117, v120, v19+int32(80), v19-int32(-64), v19+int32(63))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L5
	} else {
		goto L99
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = int32(0)
	v285 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+120)) = v285
	*(*int64)(unsafe.Add(mBase, uint32(v19)+112)) = v285
	*(*int64)(unsafe.Add(mBase, uint32(v19)+104)) = v285
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = v285
	goto L92
L94:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	if v267 == int32(0) {
		goto L76
	} else {
		goto L95
	}
L95:
	;
	v275 = F_get_attstatsslot(m, v19+int32(96), v267, int32(6), int32(0), int32(1))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	if v275 == int32(0) {
		goto L76
	} else {
		goto L97
	}
L97:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v19)+112))
	if int32(2) <= v280 {
		goto L92
	} else {
		goto L98
	}
L98:
	;
	v416 = float64(-1)
	goto L91
L99:
	;
	switch v74 - int32(3884) {
	case 0:
		goto L100
	case 1:
		goto L111
	case 2:
		goto L109
	case 3:
		goto L110
	case 4, 5:
		goto L104
	case 6:
		goto L103
	default:
		goto L101
	case 8:
		goto L102
	case 9:
		goto L108
	case 10:
		goto L107
	case 11:
		goto L105
	case 12:
		goto L106
	}
L100:
	;
	v413 = F_calc_hist_selectivity_scalar(m, v117, v19+int32(80), v210, v206, int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L5
	} else {
		goto L134
	}
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L5
	} else {
		goto L131
	}
L102:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+88)))
	if v371 == int32(1) {
		goto L122
	} else {
		goto L123
	}
L103:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v19)+112))
	v369 = F_calc_hist_selectivity_contains(m, v117, v19+int32(80), v19-int32(-64), v210, v206, v367, v368)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L5
	} else {
		goto L121
	}
L104:
	;
	v352 = F_calc_hist_selectivity_scalar(m, v117, v19+int32(80), v213, v206, int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L5
	} else {
		goto L119
	}
L105:
	;
	v346 = F_calc_hist_selectivity_scalar(m, v117, v19-int32(-64), v213, v206, int32(1))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L5
	} else {
		goto L118
	}
L106:
	;
	v340 = F_calc_hist_selectivity_scalar(m, v117, v19+int32(80), v210, v206, int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L5
	} else {
		goto L117
	}
L107:
	;
	v333 = F_calc_hist_selectivity_scalar(m, v117, v19-int32(-64), v210, v206, int32(1))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L5
	} else {
		goto L116
	}
L108:
	;
	v327 = F_calc_hist_selectivity_scalar(m, v117, v19+int32(80), v213, v206, int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L5
	} else {
		goto L115
	}
L109:
	;
	v321 = F_calc_hist_selectivity_scalar(m, v117, v19+int32(80), v210, v206, int32(1))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L5
	} else {
		goto L114
	}
L110:
	;
	v314 = F_calc_hist_selectivity_scalar(m, v117, v19+int32(80), v210, v206, int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L5
	} else {
		goto L113
	}
L111:
	;
	v308 = F_calc_hist_selectivity_scalar(m, v117, v19+int32(80), v210, v206, int32(1))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	v416 = v308
	goto L91
L113:
	;
	v416 = base.F64_sub(float64(1), v314)
	goto L91
L114:
	;
	v416 = base.F64_sub(float64(1), v321)
	goto L91
L115:
	;
	v416 = v327
	goto L91
L116:
	;
	v416 = base.F64_sub(float64(1), v333)
	goto L91
L117:
	;
	v416 = base.F64_sub(float64(1), v340)
	goto L91
L118:
	;
	v416 = v346
	goto L91
L119:
	;
	v358 = F_calc_hist_selectivity_scalar(m, v117, v19-int32(-64), v210, v206, int32(1))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L5
	} else {
		goto L120
	}
L120:
	;
	v416 = base.F64_sub(float64(1), base.F64_add(v352, base.F64_sub(float64(1), v358)))
	goto L91
L121:
	;
	v416 = v369
	goto L91
L122:
	;
	v377 = F_calc_hist_selectivity_scalar(m, v117, v19-int32(-64), v213, v206, int32(1))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L5
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+72)))
	if v379 == int32(1) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v416 = v377
	goto L91
L126:
	;
	v386 = F_calc_hist_selectivity_scalar(m, v117, v19+int32(80), v210, v206, int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L5
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v19)+112))
	v395 = F_calc_hist_selectivity_contained(m, v117, v19+int32(80), v19-int32(-64), v210, v206, v393, v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L5
	} else {
		goto L130
	}
L129:
	;
	v416 = base.F64_sub(float64(1), v386)
	goto L91
L130:
	;
	v416 = v395
	goto L91
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v74
	F_errmsg_internal(m, int32(_a_F_rangesel_5), v19)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L5
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_rangesel_1), int32(579), int32(_a_F_rangesel_4))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L5
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	v416 = v413
	goto L91
L135:
	;
	F_free_attstatsslot(m, v19+int32(136))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L5
	} else {
		goto L136
	}
L136:
	;
	if base.F64_lt(v416, float64(0)) != 0 {
		goto L65
	} else {
		goto L137
	}
L137:
	;
	v480 = v416
	goto L64
L138:
	;
	goto L65
L139:
	;
	v470 = *(*float64)(unsafe.Add(mBase, uint32(v465<<(uint(int32(3))%32))+uint32(_c_F_rangesel[1])))
	v480 = v470
	goto L64
L140:
	;
	v503 = base.F64_add(base.F64_mul(v488, v480), v151)
	goto L52
L141:
	;
	goto L142
L142:
	;
	v503 = base.F64_mul(v488, v480)
	goto L52
L143:
	;
	if base.F64_gt(v513, float64(1)) == int32(0) {
		v543 = v513
		goto L30
	} else {
		goto L144
	}
L144:
	;
	v543 = float64(1)
	goto L30
L145:
	;
	v532 = *(*float64)(unsafe.Add(mBase, uint32(v527<<(uint(int32(3))%32))+uint32(_c_F_rangesel[2])))
	v543 = v532
	goto L30
L146:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	m.T0[v550].(func(*base.Module, int32))(m, v549)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L5
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v553 = float64(1)
	if base.F64_gt(v543, v553) != 0 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	goto L148
L150:
	;
	v556 = v553
	goto L152
L151:
	;
	v556 = v543
	goto L152
L152:
	;
	v577 = base.I64_reinterpret_f64(v556)
	goto L2
L153:
	;
	F_errmsg_internal(m, int32(_a_F_rangesel_6), int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L5
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_rangesel_1), int32(257), int32(_a_F_rangesel_2))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L5
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_read_into_scalar_list(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	v10 = m.G0
	v12 = v10 - int32(_a_F_read_into_scalar_list_0)
	m.G0 = v12
	F_check_assignable(m, l1, l2, l5)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_read_into_scalar_list[0]))) = l0
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v19
	v21 = int32(1)
	v22 = F_plpgsql_yylex(m, l3, l4, l5)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L8
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_read_into_scalar_list_1))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L62
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_read_into_scalar_list_1))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L57
	}
L5:
	;
	F_plpgsql_yyerror(m, l4, int32(0), l5, int32(_a_F_read_into_scalar_list_2))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L56
	}
L6:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	F_cword_is_not_variable(m, l3, v236, l5)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L55
	}
L7:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	F_word_is_not_variable(m, l3, v233, l5)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L54
	}
L8:
	;
	if v22 == int32(44) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v27 = v21
	goto L12
L10:
	;
	v79 = v22
	v80 = v21
	goto L11
L11:
	;
	F_plpgsql_push_back_token(m, v79, l3, l4, l5)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L27
	}
L12:
	;
	if v27 == int32(1024) {
		goto L4
	} else {
		goto L14
	}
L13:
	;
	v79 = v75
	v80 = v74
	goto L11
L14:
	;
	v37 = F_plpgsql_yylex(m, l3, l4, l5)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v37 != int32(277) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	switch v37 - int32(275) {
	case 0:
		goto L7
	case 1:
		goto L6
	default:
		goto L5
	}
L17:
	;
	goto L18
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	F_check_assignable(m, v43, v44, l5)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v49 = int32(1)
	if base.Ui32(v48-v49) <= base.Ui32(v49) {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v53 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v57 = F_NameListToString(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	v60 = v47
	v61 = v53
	goto L23
L23:
	;
	v63 = v27 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v63+(v12+int32(_a_F_read_into_scalar_list_3))))) = v61
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(16)+v63))) = v71
	v74 = v27 + int32(1)
	v75 = F_plpgsql_yylex(m, l3, l4, l5)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v60 = v59
	v61 = v57
	goto L23
L25:
	;
	if v75 == int32(44) {
		v27 = v74
		goto L12
	} else {
		goto L26
	}
L26:
	;
	goto L13
L27:
	;
	v91 = F_palloc0(m, int32(40))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = int32(_a_F_read_into_scalar_list_4)
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(1)
	v97 = int32(0)
	if l2 < v97 {
		v142 = v97
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v91)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = v142
	v149 = F_palloc_mul(m, int32(4), v80)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L42
	}
L30:
	;
	goto L29
L31:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+56))
	if v103 == int32(0) {
		v142 = v97
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v106 = l2 + v103
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v102)+184))
	if base.Ui32(v107) <= base.Ui32(v106) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v102)+192))
	if base.B2i32(v116 == int32(0))|base.B2i32(base.Ui32(v106) <= base.Ui32(v116)) != 0 {
		v142 = v117
		goto L30
	} else {
		goto L37
	}
L34:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v102)+188))
	v116 = v109
	goto L33
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v102)+184)) = v103
	v114 = F_strchr(m, v103, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v102)+188)) = v114
	v116 = v114
	goto L33
L37:
	;
	v122 = v116
	v125 = v117
	goto L38
L38:
	;
	v127 = int32(1)
	v128 = v125 + v127
	*(*int32)(unsafe.Add(mBase, uint32(v102)+192)) = v128
	v131 = v122 + v127
	*(*int32)(unsafe.Add(mBase, uint32(v102)+184)) = v131
	v134 = F_strchr(m, v131, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v102)+188)) = v134
	if v134 == int32(0) {
		v142 = v128
		goto L30
	} else {
		goto L40
	}
L39:
	;
	v142 = v128
	goto L30
L40:
	;
	if base.Ui32(v134) < base.Ui32(v106) {
		v122 = v134
		v125 = v128
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+32)) = v149
	v153 = F_palloc_mul(m, int32(4), v80)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+36)) = v153
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	if v80&int32(1) != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v160 = v80 - int32(1)
	v162 = v160 << (uint(int32(2)) % 32)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(_a_F_read_into_scalar_list_3)+v162)))
	*(*int32)(unsafe.Add(mBase, uint32(v156+v162))) = v167
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(16)+v162)))
	*(*int32)(unsafe.Add(mBase, uint32(v162+v153))) = v173
	v176 = v160
	goto L46
L45:
	;
	v176 = v80
	goto L46
L46:
	;
	if v80 != int32(1) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v182 = v176
	goto L50
L48:
	;
	goto L49
L49:
	;
	F_plpgsql_adddatum(m, v91)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L53
	}
L50:
	;
	v188 = int32(2)
	v191 = v182<<(uint(v188)%32) - int32(4)
	v194 = v12 + int32(_a_F_read_into_scalar_list_3)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v194+v191)))
	*(*int32)(unsafe.Add(mBase, uint32(v156+v191))) = v196
	v200 = v12 + int32(16)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v200+v191)))
	*(*int32)(unsafe.Add(mBase, uint32(v191+v153))) = v202
	v205 = v182 - v188
	v207 = v205 << (uint(v188) % 32)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v207+v194)))
	*(*int32)(unsafe.Add(mBase, uint32(v156+v207))) = v210
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v207+v200)))
	*(*int32)(unsafe.Add(mBase, uint32(v207+v153))) = v214
	if v188 < v182 {
		v182 = v205
		goto L50
	} else {
		goto L52
	}
L51:
	;
	goto L49
L52:
	;
	goto L51
L53:
	;
	m.G0 = v12 + int32(_a_F_read_into_scalar_list_0)
	return v91
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(_a_F_read_into_scalar_list_5), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v255 = F_plpgsql_scanner_errposition(m, v254, l5)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_read_into_scalar_list_6), int32(3679), int32(_a_F_read_into_scalar_list_7))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v269 = F_NameOfDatum(m, l3)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v269
	F_errmsg(m, int32(_a_F_read_into_scalar_list_8), v12)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v276 = F_plpgsql_scanner_errposition(m, v275, l5)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_read_into_scalar_list_6), int32(3692), int32(_a_F_read_into_scalar_list_7))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_readtup_heap_2(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = l1 + int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v11
	v13 = F_palloc(m, v11)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v13))) = v11
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		F_BufFileReadExact(m, v18, v13+int32(10), l1-int32(4))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
			if v25 == int32(1) {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				F_BufFileReadExact(m, v28, v8+int32(12), int32(4))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(16)
					return v13
				}
			} else {
				m.G0 = v8 + int32(16)
				return v13
			}
		}
	}
}
func F_recompute_limits(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int64
	_ = v92
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int64
	_ = v114
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v54 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v14 = int32(_a_F_recompute_limits_0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_recompute_limits[0]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_recompute_limits[0])) = v17
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v22 = m.T0[v21].(func(*base.Module, int32, int32, int32) int64)(m, v13, v12, v10+int32(15))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(0)
	goto L1
L5:
	;
	return
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_recompute_limits[0])) = v15
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
	if v26 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v22
	if int64(0) <= v22 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	F_errcode(m, int32(671350914))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	F_errmsg(m, int32(_a_F_recompute_limits_1), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_recompute_limits_2), int32(373), int32(_a_F_recompute_limits_3))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
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
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v122 = v119
	goto L33
L16:
	;
	v103 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v103
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+136)) = uint8(v103)
	v109 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v111 == v109 {
		v117 = int64(-1)
		goto L15
	} else {
		goto L30
	}
L17:
	;
	v55 = int32(_a_F_recompute_limits_0)
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_recompute_limits[0]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_recompute_limits[0])) = v58
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v54)+24))
	v63 = m.T0[v62].(func(*base.Module, int32, int32, int32) int64)(m, v54, v12, v10+int32(15))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v92 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = v92
	v98 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+136)) = uint8(v98)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v98
	v117 = int64(-1)
	goto L15
L20:
	;
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_recompute_limits[0])) = v56
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
	if v67 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	goto L20
L23:
	;
	goto L24
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v63
	if int64(0) <= v63 {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(654573698))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	F_errmsg(m, int32(_a_F_recompute_limits_4), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_recompute_limits_2), int32(399), int32(_a_F_recompute_limits_3))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	v114 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v117 = v114 + v63
	goto L15
L31:
	;
	m.G0 = v10 + int32(16)
	return
L32:
	;
	goto L31
L33:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	switch v124 - int32(400) {
	case 0:
		goto L39
	default:
		goto L32
	case 3:
		goto L42
	case 4:
		goto L43
	case 17:
		goto L38
	case 32:
		goto L41
	case 33:
		goto L40
	case 38, 39:
		goto L37
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v122)+120)) = v117
	v183 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v122)+117)) = uint8(v183)
	goto L32
L35:
	;
	goto L34
L36:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v122+v179)))
	v122 = v181
	goto L33
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v122)+112)) = v117
	v179 = int32(36)
	goto L36
L38:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v122)+32))
	if v174 == int32(0) {
		v179 = int32(116)
		goto L36
	} else {
		goto L57
	}
L39:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v122)+36))
	if v172 != 0 {
		v122 = v172
		goto L33
	} else {
		goto L56
	}
L40:
	;
	if v117 < int64(0) {
		goto L53
	} else {
		goto L54
	}
L41:
	;
	if int64(0) <= v117 {
		goto L35
	} else {
		goto L52
	}
L42:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v122)+108))
	if v144 <= int32(0) {
		goto L32
	} else {
		goto L48
	}
L43:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122)+108))
	if v127 <= int32(0) {
		goto L32
	} else {
		goto L44
	}
L44:
	;
	v133 = int32(0)
	goto L45
L45:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v122)+104))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v134+v133<<(uint(int32(2))%32))))
	F_ExecSetTupleBound(m, v117, v138)
	mBase = m.M
	v141 = v133 + int32(1)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v122)+108))
	if v141 < v142 {
		v133 = v141
		goto L45
	} else {
		goto L47
	}
L46:
	;
	goto L32
L47:
	;
	goto L46
L48:
	;
	v150 = int32(0)
	goto L49
L49:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v122)+104))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151+v150<<(uint(int32(2))%32))))
	F_ExecSetTupleBound(m, v117, v155)
	mBase = m.M
	v158 = v150 + int32(1)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v122)+108))
	if v158 < v159 {
		v150 = v158
		goto L49
	} else {
		goto L51
	}
L50:
	;
	goto L32
L51:
	;
	goto L50
L52:
	;
	v163 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v122)+117)) = uint8(v163)
	goto L31
L53:
	;
	v167 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v122)+116)) = uint8(v167)
	goto L31
L54:
	;
	goto L55
L55:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v122)+120)) = v117
	v170 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v122)+116)) = uint8(v170)
	goto L31
L56:
	;
	goto L32
L57:
	;
	goto L32
}
func F_reconsider_outer_join_clause(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
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
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
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
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
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
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v390 int32
	_ = v390
	v4 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v28 = v18 + int32(8)
	F_op_input_types(m, v24, v18+int32(12), v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+28))
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v61 == int32(0) {
		v390 = v4
		goto L13
	} else {
		goto L14
	}
L4:
	;
	v35 = int32(48)
	if v34 == int32(0) {
		v56 = v28
		v58 = v35
		v59 = v4
		v60 = v4
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v45 = v18 + int32(12)
	v46 = int32(44)
	if v34 == int32(0) {
		v56 = v45
		v58 = v46
		v59 = v4
		v60 = v4
		goto L3
	} else {
		goto L9
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v40 < int32(2) {
		v56 = v28
		v58 = v35
		v59 = v39
		v60 = v4
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v56 = v28
	v58 = v35
	v59 = v39
	v60 = v43
	goto L3
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if int32(2) <= v50 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v54 = v53
	goto L12
L11:
	;
	v54 = v4
	goto L12
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v56 = v45
	v58 = v46
	v59 = v54
	v60 = v55
	goto L3
L13:
	;
	m.G0 = v18 + int32(16)
	return v390 & int32(1)
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v64 <= int32(0) {
		v390 = v4
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v21+v58)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v72 = int32(0)
	goto L16
L16:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86+v72<<(uint(int32(2))%32))))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+40)))
	if v91 != int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v390 = v4
	goto L13
L18:
	;
	v376 = v72 + int32(1)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v376 < v377 {
		v72 = v376
		goto L16
	} else {
		goto L83
	}
L19:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+41)))
	if v94 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	if v23 != v95 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v21)+96))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v99 = F_equal(m, v97, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v99 == int32(0) {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	if v103 == int32(0) {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v106 = int32(0)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v107 <= v106 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v115 = v106
	goto L26
L26:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125+v115<<(uint(int32(2))%32))))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v131 = F_equal(m, v59, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	if v139 == int32(0) {
		v390 = v4
		goto L13
	} else {
		goto L33
	}
L28:
	;
	if v131 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v136 = v115 + int32(1)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v136 < v137 {
		v115 = v136
		goto L26
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	goto L27
L32:
	;
	goto L18
L33:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v142 <= int32(0) {
		v390 = v4
		goto L13
	} else {
		goto L34
	}
L34:
	;
	v153 = int32(0)
	v157 = v4
	goto L35
L35:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v139)+12))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161+v153<<(uint(int32(2))%32))))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+12)))
	if v166 != int32(1) {
		v324 = v157
		goto L38
	} else {
		goto L39
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L80
	}
L37:
	;
	goto L36
L38:
	;
	v329 = v153 + int32(1)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v329 < v330 {
		v153 = v329
		v157 = v324
		goto L35
	} else {
		goto L79
	}
L39:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v169 == int32(0) {
		v324 = v157
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	if v172 <= int32(0) {
		v324 = v157
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v178 = int32(0)
	goto L42
L42:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v169)+12))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v192+v178<<(uint(int32(2))%32))))
	v198 = F_get_opfamily_member_for_cmptype(m, v196, v69, v175, int32(3))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v213 = F_bms_copy(m, v68)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L54
	}
L44:
	;
	goto L43
L45:
	;
	if v198 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v90)+52))
	if v200 == int32(0) {
		goto L44
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v208 = v178 + int32(1)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	if v208 < v209 {
		v178 = v208
		goto L42
	} else {
		goto L53
	}
L49:
	;
	v203 = F_get_opcode(m, v198)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v205 = F_get_func_leakproof(m, v203)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v205 != 0 {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	goto L48
L53:
	;
	v324 = v157
	goto L38
L54:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v90)+48))
	v216 = F_build_implied_join_equality(m, l0, v198, v211, v60, v212, v213, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v216
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if v219 == int32(0) {
		goto L37
	} else {
		goto L56
	}
L56:
	;
	v222 = int32(0)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	if v223 <= v222 {
		goto L37
	} else {
		goto L57
	}
L57:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v228 = v222
	goto L58
L58:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242+v228<<(uint(int32(2))%32))))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	v248 = int32(0)
	if v247 == v248 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v310 = F_process_equivalence(m, l0, v18+int32(4), v246)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L78
	}
L60:
	;
	if v301 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L61:
	;
	v301 = int32(1)
	goto L60
L62:
	;
	goto L63
L63:
	;
	if v226 == int32(0) {
		v294 = v248
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v301 = v294
	goto L60
L65:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v247)+4))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v258 < v257 {
		v294 = v248
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v260 = int32(1)
	if v257 <= v260 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v263 = v260
	goto L69
L68:
	;
	v263 = v257
	goto L69
L69:
	;
	v264 = int32(8)
	v269 = int32(0)
	goto L70
L70:
	;
	v276 = v269 << (uint(int32(2)) % 32)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v247+v264+v276)))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v226+v264+v276)))
	v283 = v278 & (v280 ^ int32(-1))
	v285 = base.B2i32(v283 == int32(0))
	if v283 != 0 {
		v294 = v285
		goto L64
	} else {
		goto L72
	}
L71:
	;
	v294 = v285
	goto L64
L72:
	;
	v287 = v269 + int32(1)
	if v287 != v263 {
		v269 = v287
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v305 = v228 + int32(1)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	if v305 < v306 {
		v228 = v305
		goto L58
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	goto L59
L77:
	;
	goto L37
L78:
	;
	v324 = v310 | v157
	goto L38
L79:
	;
	v390 = v324
	goto L13
L80:
	;
	F_errmsg_internal(m, int32(_a_F_reconsider_outer_join_clause_0), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_reconsider_outer_join_clause_1), int32(2597), int32(_a_F_reconsider_outer_join_clause_2))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	goto L17
}
func F_record_image_gt(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_image_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(int32(0) < v2))
	}
}
func F_record_image_le(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_image_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(v2 <= int32(0)))
	}
}
func F_record_image_ne(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_image_eq(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2 ^ int64(1)
	}
}
func F_recv_password_packet(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
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
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	F_pq_startmsgread(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = F_pq_getbyte(m)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			if v13 == int32(-1) {
				v51 = v1
				m.G0 = v7 + int32(32)
				return v51
			} else {
				if v13 != int32(112) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(16908800))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v13
							F_errmsg(m, int32(_a_F_recv_password_packet_0), v7)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_recv_password_packet_1), int32(734), int32(_a_F_recv_password_packet_2))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
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
					v20 = v7 + int32(16)
					F_initStringInfo(m, v20)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v24 = F_pq_getmessage(m, v20, int32(_a_F_recv_password_packet_3))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
							if v24 != 0 {
								F_pfree(m, v26)
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
									return int32(0)
								} else {
									v51 = v1
									m.G0 = v7 + int32(32)
									return v51
								}
							} else {
								v29 = F_strlen(m, v26)
								mBase = m.M
								v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
								if v29+int32(1) != v32 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(16908800))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_recv_password_packet_4), int32(0))
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_recv_password_packet_1), int32(754), int32(_a_F_recv_password_packet_2))
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
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
									if v29 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(16908802))
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_recv_password_packet_5), int32(0))
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_recv_password_packet_1), int32(772), int32(_a_F_recv_password_packet_2))
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
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
										v38 = F_errstart(m, int32(10), int32(0))
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											if v38 != 0 {
												F_errmsg_internal(m, int32(_a_F_recv_password_packet_6), int32(0))
												mBase = m.M
												v43 = m.ExcPending
												if v43 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_recv_password_packet_1), int32(775), int32(_a_F_recv_password_packet_2))
													mBase = m.M
													v48 = m.ExcPending
													if v48 != 0 {
														return int32(0)
													} else {
														v49 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
														v51 = v49
														m.G0 = v7 + int32(32)
														return v51
													}
												}
											} else {
												v49 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
												v51 = v49
												m.G0 = v7 + int32(32)
												return v51
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_regdatabaseout(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v8 == int32(0) {
		v12 = F_pstrdup(m, int32(_a_F_regdatabaseout_0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v30 = v12
			m.G0 = v6 + int32(16)
			return base.I64_extend_i32_u(v30)
		}
	} else {
		v16 = F_get_database_name(m, v8)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			if v16 != 0 {
				v18 = F_quote_identifier(m, v16)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					v20 = F_pstrdup(m, v18)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int64(0)
					} else {
						v30 = v20
						m.G0 = v6 + int32(16)
						return base.I64_extend_i32_u(v30)
					}
				}
			} else {
				v23 = F_palloc(m, int32(64))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
					v28 = F_pg_snprintf(m, v23, int32(64), int32(_a_F_regdatabaseout_1), v6)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int64(0)
					} else {
						v30 = v23
						m.G0 = v6 + int32(16)
						return base.I64_extend_i32_u(v30)
					}
				}
			}
		}
	}
}
func F_regexeqsel(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14370(m, l0, int32(2))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_regnamespaceout(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v8 == int32(0) {
		v12 = F_pstrdup(m, int32(_a_F_regnamespaceout_0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v30 = v12
			m.G0 = v6 + int32(16)
			return base.I64_extend_i32_u(v30)
		}
	} else {
		v16 = F_get_namespace_name(m, v8)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			if v16 != 0 {
				v18 = F_quote_identifier(m, v16)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					v20 = F_pstrdup(m, v18)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int64(0)
					} else {
						v30 = v20
						m.G0 = v6 + int32(16)
						return base.I64_extend_i32_u(v30)
					}
				}
			} else {
				v23 = F_palloc(m, int32(64))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
					v28 = F_pg_snprintf(m, v23, int32(64), int32(_a_F_regnamespaceout_1), v6)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int64(0)
					} else {
						v30 = v23
						m.G0 = v6 + int32(16)
						return base.I64_extend_i32_u(v30)
					}
				}
			}
		}
	}
}
func F_regtypein(m *base.Module, l0 int32) int64 {
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v11 == int32(45) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L30
	} else {
		goto L34
	}
L2:
	;
	m.G0 = v7 + int32(16)
	return base.I64_extend_i32_u(v129)
L3:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v129 = v128
	goto L2
L4:
	;
	v116 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v116
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_regtypein[0]))
	if v119 == v116 {
		goto L1
	} else {
		goto L32
	}
L5:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v14 != 0 {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if base.Ui32(int32(9)) < base.Ui32((v11-int32(48))&int32(255)) {
		goto L4
	} else {
		goto L9
	}
L8:
	;
	v129 = int32(0)
	goto L2
L9:
	;
	v22 = int32(_a_F_regtypein_0)
	v26 = m.G0
	v28 = v26 - int32(32)
	v29 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+24)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v28)+16)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = v29
	v37 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regtypein[1])))
	if v37 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v106 = F_strlen(m, v10)
	mBase = m.M
	if v105 != v106 {
		goto L4
	} else {
		goto L29
	}
L11:
	;
	v105 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regtypein[2])))
	if v41 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v45 = v10
	goto L17
L15:
	;
	goto L16
L16:
	;
	v55 = v22
	v56 = v37
	goto L20
L17:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v51 == v37 {
		v45 = v45 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v105 = v45 - v10
	goto L10
L19:
	;
	goto L18
L20:
	;
	v63 = v28 + int32(base.Ui32(v56)>>(uint(int32(3))%32))&int32(28)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v65 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v64 | v65<<(uint(v56)%32)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v69 != 0 {
		v55 = v55 + v65
		v56 = v69
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v72 == int32(0) {
		v95 = v10
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v105 = v95 - v10
	goto L10
L24:
	;
	v76 = v10
	v77 = v72
	goto L25
L25:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(base.Ui32(v77)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v85)>>(uint(v77)%32))&int32(1) == int32(0) {
		v95 = v76
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v95 = v93
	goto L23
L27:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	v93 = v76 + int32(1)
	if v91 != 0 {
		v76 = v93
		v77 = v91
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v112 = F_DirectInputFunctionCallSafe(m, int32(588), v10, int32(-1), v9, v7+int32(8))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return int64(0)
L31:
	;
	goto L3
L32:
	;
	v126 = F_parseTypeString(m, v10, v7+int32(8), v7+int32(4), v9)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L3
L34:
	;
	F_errmsg_internal(m, int32(_a_F_regtypein_1), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_regtypein_2), int32(1199), int32(_a_F_regtypein_3))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_relmap_desc(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	Fn14372(m, l0, l1, int32(_a_F_relmap_desc_0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_remove_nulling_relids_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
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
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
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
	var v132 int32
	_ = v132
	var v134 int32
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
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 != int32(321) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v161 = F_expression_tree_mutator_impl(m, l0, int32(1131), l1)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L12
	} else {
		goto L50
	}
L5:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v147 + int32(1)
	v153 = F_query_tree_mutator_impl(m, l0, int32(1131), l1, int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L12
	} else {
		goto L49
	}
L6:
	;
	if v8 == int32(67) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v82 != v83 {
		goto L4
	} else {
		goto L31
	}
L9:
	;
	if v8 != int32(6) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v15 != v16 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v20 = F_bms_is_member(m, v18, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if v20 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v26 = int32(0)
	if base.B2i32(v24 == v26)|base.B2i32(v25 == v26) != 0 {
		v71 = v26
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v71 == int32(0) {
		goto L4
	} else {
		goto L28
	}
L16:
	;
	goto L15
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v36 < v37 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v39 = v36
	goto L20
L19:
	;
	v39 = v37
	goto L20
L20:
	;
	if v39 <= int32(1) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v42 = int32(1)
	goto L23
L22:
	;
	v42 = v39
	goto L23
L23:
	;
	v43 = int32(8)
	v48 = int32(0)
	goto L24
L24:
	;
	v55 = v48 << (uint(int32(2)) % 32)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v25+v43+v55)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v24+v43+v55)))
	v60 = v57 & v59
	v62 = base.B2i32(v60 != int32(0))
	if v60 != 0 {
		v71 = v62
		goto L16
	} else {
		goto L26
	}
L25:
	;
	v71 = v62
	goto L16
L26:
	;
	v64 = v48 + int32(1)
	if v64 != v42 {
		v48 = v64
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v74 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v78 = F_bms_difference(m, v76, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L12
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+24)) = v78
	return v74
L31:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v87 = int32(0)
	if base.B2i32(v85 == v87)|base.B2i32(v86 == v87) != 0 {
		v132 = v87
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v132 != 0 {
		goto L4
	} else {
		goto L45
	}
L33:
	;
	goto L32
L34:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v97 < v98 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v100 = v97
	goto L37
L36:
	;
	v100 = v98
	goto L37
L37:
	;
	if v100 <= int32(1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v103 = int32(1)
	goto L40
L39:
	;
	v103 = v100
	goto L40
L40:
	;
	v104 = int32(8)
	v109 = int32(0)
	goto L41
L41:
	;
	v116 = v109 << (uint(int32(2)) % 32)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v86+v104+v116)))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v85+v104+v116)))
	v121 = v118 & v120
	v123 = base.B2i32(v121 != int32(0))
	if v121 != 0 {
		v132 = v123
		goto L33
	} else {
		goto L43
	}
L42:
	;
	v132 = v123
	goto L33
L43:
	;
	v125 = v109 + int32(1)
	if v125 != v103 {
		v109 = v125
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v134 = F_expression_tree_mutator_impl(m, l0, int32(1131), l1)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L12
	} else {
		goto L46
	}
L46:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v138 = F_bms_difference(m, v136, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = v138
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v143 = F_bms_difference(m, v141, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+8)) = v143
	return v134
L49:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v155 - int32(1)
	return v153
L50:
	;
	return v161
}
func F_removeabbrev_heap(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	if v4 < l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = v4
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v11 + int32(32)
	return
L4:
	;
	v25 = l1 + v19*int32(24)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v28 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v26 - v28
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v27 + v28
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v37 = int32(*(*int16)(unsafe.Add(mBase, uint32(v36)+10)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v41 = F_heap_getattr_1(m, v11+int32(12), v37, v38, v25+int32(16))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v41
	v45 = v19 + int32(1)
	if v45 != l2 {
		v19 = v45
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
func F_removeabbrev_index(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v4 = int32(0)
	if v4 < l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v12 = v4
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v16 = l1 + v12*int32(24)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v23 = F_index_getattr_2(m, v17, int32(1), v20, v16+int32(16))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v23
	v27 = v12 + int32(1)
	if v27 != l2 {
		v12 = v27
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
func F_removeabbrev_index_brin(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	v4 = int32(0)
	if l2 <= v4 {
	} else {
		v11 = l2 & int32(3)
		v12 = int32(0)
		if base.Ui32(int32(4)) <= base.Ui32(l2) {
			v17 = v12
			v22 = v4
			for {
				v26 = l1 + v17*int32(24)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				v28 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v27)+4)))
				*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = v28
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
				v31 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v30)+4)))
				*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v31
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
				v34 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v33)+4)))
				*(*int64)(unsafe.Add(mBase, uint32(v26)+56)) = v34
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v26)+72))
				v37 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v36)+4)))
				*(*int64)(unsafe.Add(mBase, uint32(v26)+80)) = v37
				v39 = int32(4)
				v40 = v17 + v39
				v42 = v22 + v39
				if v42 != l2&int32(2147483644) {
					v17 = v40
					v22 = v42
					continue
				} else {
					break
				}
				break
			}
			if v11 == int32(0) {
			} else {
				v46 = v40
				v53 = v46
				v59 = v4
				for {
					v62 = l1 + v53*int32(24)
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
					v64 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v63)+4)))
					*(*int64)(unsafe.Add(mBase, uint32(v62)+8)) = v64
					v66 = int32(1)
					v69 = v59 + v66
					if v69 != v11 {
						v53 = v53 + v66
						v59 = v69
						continue
					} else {
						break
					}
					break
				}
			}
		} else {
			v46 = v12
			v53 = v46
			v59 = v4
			for {
				v62 = l1 + v53*int32(24)
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
				v64 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v63)+4)))
				*(*int64)(unsafe.Add(mBase, uint32(v62)+8)) = v64
				v66 = int32(1)
				v69 = v59 + v66
				if v69 != v11 {
					v53 = v53 + v66
					v59 = v69
					continue
				} else {
					break
				}
				break
			}
		}
	}
	return
}
func F_reorderqueue_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+136))
	if v8 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v18 = int32(0)
	goto L5
L4:
	;
	return v49
L5:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v13))))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v15))))
	if v28 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	return int32(0)
L7:
	;
	return v26 ^ int32(1)
L8:
	;
	goto L9
L9:
	;
	if v26&int32(1) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(-1)
L11:
	;
	goto L12
L12:
	;
	v39 = v18 << (uint(int32(3)) % 32)
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v16+v39)))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v39+v14)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+188))
	v47 = v44 + v18*int32(36)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v49 = m.T0[v48].(func(*base.Module, int64, int64, int32) int32)(m, v41, v43, v47)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	if v49 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v54 = v18 + int32(1)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+136))
	if v54 < v55 {
		v18 = v54
		goto L5
	} else {
		goto L16
	}
L16:
	;
	goto L6
}
func F_reparameterize_path_by_child(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
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
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
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
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
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
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
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
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
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
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
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
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
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
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 float64
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v494 int32
	_ = v494
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v10 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v494
L2:
	;
	v494 = l1
	goto L1
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+252))
	v15 = int32(0)
	if base.B2i32(v13 == v15)|base.B2i32(v14 == v15) != 0 {
		v60 = v15
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v60 == int32(0) {
		goto L2
	} else {
		goto L17
	}
L5:
	;
	goto L4
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v25 < v26 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v28 = v25
	goto L9
L8:
	;
	v28 = v26
	goto L9
L9:
	;
	if v28 <= int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v31 = int32(1)
	goto L12
L11:
	;
	v31 = v28
	goto L12
L12:
	;
	v32 = int32(8)
	v37 = int32(0)
	goto L13
L13:
	;
	v44 = v37 << (uint(int32(2)) % 32)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v14+v32+v44)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v13+v32+v44)))
	v49 = v46 & v48
	v51 = base.B2i32(v49 != int32(0))
	if v49 != 0 {
		v60 = v51
		goto L5
	} else {
		goto L15
	}
L14:
	;
	v60 = v51
	goto L5
L15:
	;
	v53 = v37 + int32(1)
	if v53 != v31 {
		v37 = v53
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v63 = int32(0)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v64 - int32(282) {
	case 0:
		goto L32
	case 1:
		goto L31
	default:
		v494 = v63
		goto L1
	case 3:
		goto L30
	case 4:
		goto L29
	case 5:
		goto L28
	case 9:
		goto L27
	case 10:
		goto L26
	case 11:
		goto L22
	case 14:
		goto L21
	case 15:
		goto L20
	case 16:
		goto L19
	case 18:
		goto L25
	case 19:
		goto L24
	case 20:
		goto L23
	}
L18:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	v285 = F_adjust_child_relids_multilevel(m, l0, v283, l2, v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L33
	} else {
		goto L105
	}
L19:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v275 = F_reparameterize_path_by_child(m, l0, v274, l2)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L33
	} else {
		goto L102
	}
L20:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v264 = F_reparameterize_path_by_child(m, l0, v263, l2)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L33
	} else {
		goto L99
	}
L21:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v260 = F_reparameterize_path_by_child(m, l0, v259, l2)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L33
	} else {
		goto L97
	}
L22:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v253 == int32(0) {
		goto L18
	} else {
		goto L94
	}
L23:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v232 = F_reparameterize_path_by_child(m, l0, v231, l2)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L33
	} else {
		goto L88
	}
L24:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v210 = F_reparameterize_path_by_child(m, l0, v209, l2)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L33
	} else {
		goto L82
	}
L25:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v193 = F_reparameterize_path_by_child(m, l0, v192, l2)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L33
	} else {
		goto L77
	}
L26:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+204))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	v166 = F_adjust_appendrel_attrs_multilevel(m, l0, v164, l2, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L33
	} else {
		goto L64
	}
L27:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+204))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	v138 = F_adjust_appendrel_attrs_multilevel(m, l0, v136, l2, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L33
	} else {
		goto L52
	}
L28:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v129 == int32(0) {
		goto L18
	} else {
		goto L49
	}
L29:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v123 == int32(0) {
		goto L18
	} else {
		goto L46
	}
L30:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+204))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	v115 = F_adjust_appendrel_attrs_multilevel(m, l0, v113, l2, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L33
	} else {
		goto L43
	}
L31:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+96))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	v103 = F_adjust_appendrel_attrs_multilevel(m, l0, v101, l2, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L33
	} else {
		goto L41
	}
L32:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+204))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	v70 = F_adjust_appendrel_attrs_multilevel(m, l0, v68, l2, v69)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	return int32(0)
L34:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v74)+204)) = v70
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v76 != int32(344) {
		goto L18
	} else {
		goto L35
	}
L35:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+76))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v81 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+32))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	v97 = F_adjust_appendrel_attrs_multilevel(m, l0, v95, l2, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L33
	} else {
		goto L40
	}
L37:
	;
	v93 = v81 + v80<<(uint(int32(2))%32)
	goto L36
L38:
	;
	goto L39
L39:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+52))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v93 = v87 + v80<<(uint(int32(2))%32) - int32(4)
	goto L36
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+32)) = v97
	goto L18
L41:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v105)+96)) = v103
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	v109 = F_adjust_appendrel_attrs_multilevel(m, l0, v107, l2, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L33
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v109
	goto L18
L43:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+204)) = v115
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v120 = F_reparameterize_path_by_child(m, l0, v119, l2)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L33
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v120
	if v120 != 0 {
		goto L18
	} else {
		goto L45
	}
L45:
	;
	v494 = v63
	goto L1
L46:
	;
	v126 = F_reparameterize_pathlist_by_child(m, l0, v123, l2)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L33
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v126
	if v126 != 0 {
		goto L18
	} else {
		goto L48
	}
L48:
	;
	v494 = v63
	goto L1
L49:
	;
	v132 = F_reparameterize_pathlist_by_child(m, l0, v129, l2)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L33
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v132
	if v132 != 0 {
		goto L18
	} else {
		goto L51
	}
L51:
	;
	v494 = v63
	goto L1
L52:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v140)+204)) = v138
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v142 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v143 = F_reparameterize_path_by_child(m, l0, v142, l2)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L33
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v149 != 0 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v143
	if v143 == int32(0) {
		v494 = v63
		goto L1
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	v151 = F_adjust_appendrel_attrs_multilevel(m, l0, v149, l2, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L33
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+176))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+168))
	if v156 == int32(0) {
		goto L18
	} else {
		goto L62
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v151
	goto L60
L62:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v160 = m.T0[v156].(func(*base.Module, int32, int32, int32) int32)(m, l0, v159, l2)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L33
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v160
	goto L18
L64:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v168)+204)) = v166
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v170 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v171 = F_reparameterize_pathlist_by_child(m, l0, v170, l2)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L33
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v177 != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v171
	if v171 == int32(0) {
		v494 = v63
		goto L1
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	v179 = F_adjust_appendrel_attrs_multilevel(m, l0, v177, l2, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L33
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	if v182 == int32(0) {
		goto L18
	} else {
		goto L74
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v179
	goto L72
L74:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v182)+8))
	if v185 == int32(0) {
		goto L18
	} else {
		goto L75
	}
L75:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v189 = m.T0[v185].(func(*base.Module, int32, int32, int32) int32)(m, l0, v188, l2)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L33
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+84)) = v189
	goto L18
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v193
	if v193 == int32(0) {
		v494 = v63
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v199 = F_reparameterize_path_by_child(m, l0, v198, l2)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L33
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+84)) = v199
	if v199 == int32(0) {
		v494 = v63
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	v206 = F_adjust_appendrel_attrs_multilevel(m, l0, v204, l2, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L33
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v206
	goto L18
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v210
	if v210 == int32(0) {
		v494 = v63
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v216 = F_reparameterize_path_by_child(m, l0, v215, l2)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L33
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+84)) = v216
	if v216 == int32(0) {
		v494 = v63
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	v223 = F_adjust_appendrel_attrs_multilevel(m, l0, v221, l2, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L33
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v223
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	v228 = F_adjust_appendrel_attrs_multilevel(m, l0, v226, l2, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L33
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v228
	goto L18
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v232
	if v232 == int32(0) {
		v494 = v63
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v238 = F_reparameterize_path_by_child(m, l0, v237, l2)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L33
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+84)) = v238
	if v238 == int32(0) {
		v494 = v63
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	v245 = F_adjust_appendrel_attrs_multilevel(m, l0, v243, l2, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L33
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v245
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	v250 = F_adjust_appendrel_attrs_multilevel(m, l0, v248, l2, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L33
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v250
	goto L18
L94:
	;
	v256 = F_reparameterize_pathlist_by_child(m, l0, v253, l2)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L33
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v256
	if v256 != 0 {
		goto L18
	} else {
		goto L96
	}
L96:
	;
	v494 = v63
	goto L1
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v260
	if v260 != 0 {
		goto L18
	} else {
		goto L98
	}
L98:
	;
	v494 = v63
	goto L1
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v264
	if v264 == int32(0) {
		v494 = v63
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	v271 = F_adjust_appendrel_attrs_multilevel(m, l0, v269, l2, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L33
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v271
	goto L18
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v275
	if v275 == int32(0) {
		v494 = v63
		goto L1
	} else {
		goto L103
	}
L103:
	;
	goto L18
L104:
	;
	if v377 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L105:
	;
	v287 = int32(0)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+48))
	if v289 == v287 {
		v377 = v287
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v292 = int32(0)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	if v293 <= v292 {
		v377 = v292
		goto L104
	} else {
		goto L107
	}
L107:
	;
	v301 = int32(0)
	goto L109
L108:
	;
	v377 = v310
	goto L104
L109:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v289)+12))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v306+v301<<(uint(int32(2))%32))))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+4))
	v312 = int32(0)
	if base.B2i32(v311 == v312)|base.B2i32(v285 == v312) != 0 {
		v358 = base.B2i32(v311|v285 == v312)
		goto L112
	} else {
		goto L113
	}
L110:
	;
	v377 = int32(0)
	goto L104
L111:
	;
	if v358 != 0 {
		goto L108
	} else {
		goto L122
	}
L112:
	;
	goto L111
L113:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v311)+4))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v285)+4))
	if v326 != v327 {
		v358 = int32(0)
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v329 = int32(1)
	if v326 <= v329 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v332 = v329
	goto L117
L116:
	;
	v332 = v326
	goto L117
L117:
	;
	v333 = int32(8)
	v338 = int32(0)
	goto L118
L118:
	;
	v346 = v338 << (uint(int32(2)) % 32)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v311+v333+v346)))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v285+v333+v346)))
	v351 = base.B2i32(v348 == v350)
	if v348 != v350 {
		v358 = v351
		goto L112
	} else {
		goto L120
	}
L119:
	;
	v358 = v351
	goto L112
L120:
	;
	v354 = v338 + int32(1)
	if v354 != v332 {
		v338 = v354
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v364 = v301 + int32(1)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	if v364 < v365 {
		v301 = v364
		goto L109
	} else {
		goto L123
	}
L123:
	;
	goto L110
L124:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v381 = F_GetMemoryChunkContext(m, v380)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L33
	} else {
		goto L127
	}
L125:
	;
	v413 = v377
	goto L126
L126:
	;
	F_bms_free(m, v285)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L33
	} else {
		goto L133
	}
L127:
	;
	v383 = int32(_a_F_reparameterize_path_by_child_0)
	v384 = *(*int32)(unsafe.Add(mBase, _c_F_reparameterize_path_by_child[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_reparameterize_path_by_child[0])) = v381
	v388 = F_palloc0(m, int32(24))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L33
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v388))) = int32(281)
	v392 = F_bms_copy(m, v285)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L33
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v388)+4)) = v392
	v395 = *(*float64)(unsafe.Add(mBase, uint32(v282)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v388)+8)) = v395
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v282)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v388)+16)) = v397
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	v400 = F_adjust_appendrel_attrs_multilevel(m, l0, v397, l2, v399)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L33
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v388)+16)) = v400
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v282)+20))
	v404 = F_bms_copy(m, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L33
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v388)+20)) = v404
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v380)+48))
	v408 = F_lappend(m, v407, v388)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L33
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v380)+48)) = v408
	*(*int32)(unsafe.Add(mBase, _c_F_reparameterize_path_by_child[0])) = v384
	v413 = v388
	goto L126
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v413
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v420)+72))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l2)+252))
	v423 = int32(0)
	if base.B2i32(v421 == v423)|base.B2i32(v422 == v423) != 0 {
		v468 = v423
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v468 == int32(0) {
		goto L2
	} else {
		goto L147
	}
L135:
	;
	goto L134
L136:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v421)+4))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v422)+4))
	if v433 < v434 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v436 = v433
	goto L139
L138:
	;
	v436 = v434
	goto L139
L139:
	;
	if v436 <= int32(1) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v439 = int32(1)
	goto L142
L141:
	;
	v439 = v436
	goto L142
L142:
	;
	v440 = int32(8)
	v445 = int32(0)
	goto L143
L143:
	;
	v452 = v445 << (uint(int32(2)) % 32)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v422+v440+v452)))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v421+v440+v452)))
	v457 = v454 & v456
	v459 = base.B2i32(v457 != int32(0))
	if v457 != 0 {
		v468 = v459
		goto L135
	} else {
		goto L145
	}
L144:
	;
	v468 = v459
	goto L135
L145:
	;
	v461 = v445 + int32(1)
	if v461 != v439 {
		v445 = v461
		goto L143
	} else {
		goto L146
	}
L146:
	;
	goto L144
L147:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v472 = F_copy_pathtarget(m, v471)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L33
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v472
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v472)+4))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	v477 = F_adjust_appendrel_attrs_multilevel(m, l0, v475, l2, v476)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L33
	} else {
		goto L149
	}
L149:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v479)+4)) = v477
	goto L2
}
func F_reparameterize_pathlist_by_child(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v4 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 < v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = v4
	v14 = v4
	goto L4
L2:
	;
	v40 = v4
	goto L3
L3:
	;
	return v40
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16+v14<<(uint(int32(2))%32))))
	v21 = F_reparameterize_path_by_child(m, l0, v20, l2)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v40 = v31
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	if v21 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_list_free(m, v13)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v31 = F_lappend(m, v13, v21)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v34 = v14 + int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v34 < v35 {
		v13 = v31
		v14 = v34
		goto L4
	} else {
		goto L13
	}
L13:
	;
	goto L5
}
func F_repeat_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
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
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
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
	var v230 int32
	_ = v230
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
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
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
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
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
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
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
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v414 int32
	_ = v414
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v474 int32
	_ = v474
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
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
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v576 int32
	_ = v576
	var v586 int32
	_ = v586
	v10 = l2
	v11 = l3
	v12 = l4
	goto L3
L1:
	;
	F_createarc(m, v576, int32(110), int32(0), l1, v25)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L30
	} else {
		goto L267
	}
L2:
	;
	return
L3:
	;
	v16 = int32(2)
	if v16 <= v12 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v565 != 0 {
		goto L264
	} else {
		goto L265
	}
L5:
	;
	v19 = v16
	goto L7
L6:
	;
	v19 = v12
	goto L7
L7:
	;
	if v12 == int32(256) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v22 = int32(3)
	goto L10
L9:
	;
	v22 = v19
	goto L10
L10:
	;
	v25 = v10
	v26 = v11
	goto L11
L11:
	;
	v31 = int32(2)
	if v31 <= v26 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L4
L13:
	;
	goto L12
L14:
	;
	v34 = v31
	goto L16
L15:
	;
	v34 = v26
	goto L16
L16:
	;
	if v26 == int32(256) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v39 = int32(12)
	goto L19
L18:
	;
	v39 = v34 << (uint(int32(2)) % 32)
	goto L19
L19:
	;
	v40 = v39 + v22
	if v40 != int32(11) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	switch v40 {
	case 0:
		goto L29
	case 1:
		goto L28
	case 2:
		goto L27
	case 3:
		goto L26
	default:
		goto L13
	case 5:
		goto L2
	case 6:
		goto L25
	case 7:
		goto L24
	case 10:
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v551 = F_newstate(m, v550)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L30
	} else {
		goto L259
	}
L23:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v536 = F_newstate(m, v535)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L30
	} else {
		goto L254
	}
L24:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v350 = F_newstate(m, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L30
	} else {
		goto L172
	}
L25:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v330 = F_newstate(m, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L30
	} else {
		goto L164
	}
L26:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v207 = F_newstate(m, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L30
	} else {
		goto L109
	}
L27:
	;
	F_repeat_1(m, l0, l1, v25, int32(1), v12)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L30
	} else {
		goto L83
	}
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_repeat_1[0]))
	if v106 != 0 {
		goto L59
	} else {
		goto L60
	}
L29:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v25
	F_deltraverse(m, v43, l1)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return
L31:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+76))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if v48 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v51
	goto L34
L33:
	;
	goto L34
L34:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_repeat_1[0]))
	if v57 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L30
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v60 <= v61 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L37
L39:
	;
	v576 = v55
	goto L1
L40:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v63 == int32(0) {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v80 == int32(0) {
		goto L39
	} else {
		goto L51
	}
L43:
	;
	v69 = v63
	goto L44
L44:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	if v73 != v25 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L39
L46:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	if v79 != 0 {
		v69 = v79
		goto L44
	} else {
		goto L50
	}
L47:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
	if v75 != 0 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v76 == int32(110) {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	goto L45
L51:
	;
	v86 = v80
	goto L52
L52:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	if v90 != l1 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L39
L54:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v86)+24))
	if v96 != 0 {
		v86 = v96
		goto L52
	} else {
		goto L58
	}
L55:
	;
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+4)))
	if v92 != 0 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if v93 == int32(110) {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	goto L53
L59:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L30
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v109 <= v110 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L61
L63:
	;
	v576 = v104
	goto L1
L64:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v112 == int32(0) {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v129 == int32(0) {
		goto L63
	} else {
		goto L75
	}
L67:
	;
	v118 = v112
	goto L68
L68:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	if v122 != v25 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L63
L70:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
	if v128 != 0 {
		v118 = v128
		goto L68
	} else {
		goto L74
	}
L71:
	;
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v118)+4)))
	if v124 != 0 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v125 == int32(110) {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	goto L69
L75:
	;
	v135 = v129
	goto L76
L76:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	if v139 != l1 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L63
L78:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v135)+24))
	if v145 != 0 {
		v135 = v145
		goto L76
	} else {
		goto L82
	}
L79:
	;
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+4)))
	if v141 != 0 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	if v142 == int32(110) {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	goto L78
L82:
	;
	goto L77
L83:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v156 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_repeat_1[0]))
	if v159 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L30
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v162 <= v163 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L87
L89:
	;
	v576 = v157
	goto L1
L90:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v165 == int32(0) {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v182 == int32(0) {
		goto L89
	} else {
		goto L101
	}
L93:
	;
	v171 = v165
	goto L94
L94:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
	if v175 != v25 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	goto L89
L96:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v171)+16))
	if v181 != 0 {
		v171 = v181
		goto L94
	} else {
		goto L100
	}
L97:
	;
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+4)))
	if v177 != 0 {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	if v178 == int32(110) {
		goto L2
	} else {
		goto L99
	}
L99:
	;
	goto L96
L100:
	;
	goto L95
L101:
	;
	v188 = v182
	goto L102
L102:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
	if v192 != l1 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	goto L89
L104:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v188)+24))
	if v198 != 0 {
		v188 = v198
		goto L102
	} else {
		goto L108
	}
L105:
	;
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+4)))
	if v194 != 0 {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	if v195 == int32(110) {
		goto L2
	} else {
		goto L107
	}
L107:
	;
	goto L104
L108:
	;
	goto L103
L109:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v209 != 0 {
		goto L2
	} else {
		goto L110
	}
L110:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_moveouts(m, v210, l1, v207)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L30
	} else {
		goto L111
	}
L111:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_moveins(m, v213, v25, v207)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L30
	} else {
		goto L112
	}
L112:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_repeat_1[0]))
	if v218 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L30
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v207)+8))
	if v221 <= v222 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	goto L115
L117:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_repeat_1[0]))
	if v278 != 0 {
		goto L139
	} else {
		goto L140
	}
L118:
	;
	F_createarc(m, v216, int32(110), int32(0), l1, v207)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L30
	} else {
		goto L138
	}
L119:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v224 == int32(0) {
		goto L118
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v207)+16))
	if v241 == int32(0) {
		goto L118
	} else {
		goto L130
	}
L122:
	;
	v230 = v224
	goto L123
L123:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v230)+12))
	if v234 != v207 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	goto L118
L125:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v230)+16))
	if v240 != 0 {
		v230 = v240
		goto L123
	} else {
		goto L129
	}
L126:
	;
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v230)+4)))
	if v236 != 0 {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	if v237 == int32(110) {
		goto L117
	} else {
		goto L128
	}
L128:
	;
	goto L125
L129:
	;
	goto L124
L130:
	;
	v247 = v241
	goto L131
L131:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v247)+8))
	if v251 != l1 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	goto L118
L133:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v247)+24))
	if v257 != 0 {
		v247 = v257
		goto L131
	} else {
		goto L137
	}
L134:
	;
	v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+4)))
	if v253 != 0 {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	if v254 == int32(110) {
		goto L117
	} else {
		goto L136
	}
L136:
	;
	goto L133
L137:
	;
	goto L132
L138:
	;
	goto L117
L139:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L30
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v281 <= v282 {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	goto L141
L143:
	;
	F_createarc(m, v276, int32(110), int32(0), v207, v25)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L30
	} else {
		goto L163
	}
L144:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v207)+20))
	if v284 == int32(0) {
		goto L143
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v301 == int32(0) {
		goto L143
	} else {
		goto L155
	}
L147:
	;
	v290 = v284
	goto L148
L148:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v290)+12))
	if v294 != v25 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	goto L143
L150:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v290)+16))
	if v300 != 0 {
		v290 = v300
		goto L148
	} else {
		goto L154
	}
L151:
	;
	v296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v290)+4)))
	if v296 != 0 {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	if v297 == int32(110) {
		goto L2
	} else {
		goto L153
	}
L153:
	;
	goto L150
L154:
	;
	goto L149
L155:
	;
	v307 = v301
	goto L156
L156:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v307)+8))
	if v311 != v207 {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	goto L143
L158:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v307)+24))
	if v317 != 0 {
		v307 = v317
		goto L156
	} else {
		goto L162
	}
L159:
	;
	v313 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v307)+4)))
	if v313 != 0 {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	if v314 == int32(110) {
		goto L2
	} else {
		goto L161
	}
L161:
	;
	goto L158
L162:
	;
	goto L157
L163:
	;
	return
L164:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v332 != 0 {
		goto L2
	} else {
		goto L165
	}
L165:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_moveouts(m, v333, l1, v330)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L30
	} else {
		goto L166
	}
L166:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_dupnfa(m, v336, v330, v25, l1, v330)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L30
	} else {
		goto L167
	}
L167:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v339 != 0 {
		goto L2
	} else {
		goto L168
	}
L168:
	;
	v340 = int32(1)
	F_repeat_1(m, l0, l1, v330, v340, v12-v340)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L30
	} else {
		goto L169
	}
L169:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v345 != 0 {
		goto L2
	} else {
		goto L170
	}
L170:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_newarc(m, v346, l1, v330)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L30
	} else {
		goto L171
	}
L171:
	;
	return
L172:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v353 = F_newstate(m, v352)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L30
	} else {
		goto L173
	}
L173:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v355 != 0 {
		goto L2
	} else {
		goto L174
	}
L174:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_moveouts(m, v356, l1, v350)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L30
	} else {
		goto L175
	}
L175:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_moveins(m, v359, v25, v353)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L30
	} else {
		goto L176
	}
L176:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_repeat_1[0]))
	if v364 != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L30
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v350)+8))
	if v367 <= v368 {
		goto L183
	} else {
		goto L184
	}
L180:
	;
	goto L179
L181:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_repeat_1[0]))
	if v424 != 0 {
		goto L203
	} else {
		goto L204
	}
L182:
	;
	F_createarc(m, v362, int32(110), int32(0), l1, v350)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L30
	} else {
		goto L202
	}
L183:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v370 == int32(0) {
		goto L182
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v350)+16))
	if v387 == int32(0) {
		goto L182
	} else {
		goto L194
	}
L186:
	;
	v376 = v370
	goto L187
L187:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v376)+12))
	if v380 != v350 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	goto L182
L189:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v376)+16))
	if v386 != 0 {
		v376 = v386
		goto L187
	} else {
		goto L193
	}
L190:
	;
	v382 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v376)+4)))
	if v382 != 0 {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v376)))
	if v383 == int32(110) {
		goto L181
	} else {
		goto L192
	}
L192:
	;
	goto L189
L193:
	;
	goto L188
L194:
	;
	v393 = v387
	goto L195
L195:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v393)+8))
	if v397 != l1 {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	goto L182
L197:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v393)+24))
	if v403 != 0 {
		v393 = v403
		goto L195
	} else {
		goto L201
	}
L198:
	;
	v399 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v393)+4)))
	if v399 != 0 {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	if v400 == int32(110) {
		goto L181
	} else {
		goto L200
	}
L200:
	;
	goto L197
L201:
	;
	goto L196
L202:
	;
	goto L181
L203:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L30
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v353)+12))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v427 <= v428 {
		goto L209
	} else {
		goto L210
	}
L206:
	;
	goto L205
L207:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v484 = *(*int32)(unsafe.Add(mBase, _c_F_repeat_1[0]))
	if v484 != 0 {
		goto L229
	} else {
		goto L230
	}
L208:
	;
	F_createarc(m, v422, int32(110), int32(0), v353, v25)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L30
	} else {
		goto L228
	}
L209:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v353)+20))
	if v430 == int32(0) {
		goto L208
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v447 == int32(0) {
		goto L208
	} else {
		goto L220
	}
L212:
	;
	v436 = v430
	goto L213
L213:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v436)+12))
	if v440 != v25 {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	goto L208
L215:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v436)+16))
	if v446 != 0 {
		v436 = v446
		goto L213
	} else {
		goto L219
	}
L216:
	;
	v442 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v436)+4)))
	if v442 != 0 {
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v436)))
	if v443 == int32(110) {
		goto L207
	} else {
		goto L218
	}
L218:
	;
	goto L215
L219:
	;
	goto L214
L220:
	;
	v453 = v447
	goto L221
L221:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v453)+8))
	if v457 != v353 {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	goto L208
L223:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v453)+24))
	if v463 != 0 {
		v453 = v463
		goto L221
	} else {
		goto L227
	}
L224:
	;
	v459 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v453)+4)))
	if v459 != 0 {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	if v460 == int32(110) {
		goto L207
	} else {
		goto L226
	}
L226:
	;
	goto L223
L227:
	;
	goto L222
L228:
	;
	goto L207
L229:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L30
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v353)+12))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v350)+8))
	if v487 <= v488 {
		goto L234
	} else {
		goto L235
	}
L232:
	;
	goto L231
L233:
	;
	F_createarc(m, v482, int32(110), int32(0), v353, v350)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L30
	} else {
		goto L253
	}
L234:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v353)+20))
	if v490 == int32(0) {
		goto L233
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v350)+16))
	if v507 == int32(0) {
		goto L233
	} else {
		goto L245
	}
L237:
	;
	v496 = v490
	goto L238
L238:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v496)+12))
	if v500 != v350 {
		goto L240
	} else {
		goto L241
	}
L239:
	;
	goto L233
L240:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v496)+16))
	if v506 != 0 {
		v496 = v506
		goto L238
	} else {
		goto L244
	}
L241:
	;
	v502 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v496)+4)))
	if v502 != 0 {
		goto L240
	} else {
		goto L242
	}
L242:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v496)))
	if v503 == int32(110) {
		goto L2
	} else {
		goto L243
	}
L243:
	;
	goto L240
L244:
	;
	goto L239
L245:
	;
	v513 = v507
	goto L246
L246:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v513)+8))
	if v517 != v353 {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	goto L233
L248:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v513)+24))
	if v523 != 0 {
		v513 = v523
		goto L246
	} else {
		goto L252
	}
L249:
	;
	v519 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v513)+4)))
	if v519 != 0 {
		goto L248
	} else {
		goto L250
	}
L250:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	if v520 == int32(110) {
		goto L2
	} else {
		goto L251
	}
L251:
	;
	goto L248
L252:
	;
	goto L247
L253:
	;
	return
L254:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v538 != 0 {
		goto L2
	} else {
		goto L255
	}
L255:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_moveouts(m, v539, l1, v536)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L30
	} else {
		goto L256
	}
L256:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_dupnfa(m, v542, v536, v25, l1, v536)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L30
	} else {
		goto L257
	}
L257:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v545 != 0 {
		goto L2
	} else {
		goto L258
	}
L258:
	;
	v546 = int32(1)
	v10 = v536
	v11 = v26 - v546
	v12 = v12 - v546
	goto L3
L259:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v553 != 0 {
		goto L2
	} else {
		goto L260
	}
L260:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_moveouts(m, v554, l1, v551)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L30
	} else {
		goto L261
	}
L261:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_dupnfa(m, v557, v551, v25, l1, v551)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L30
	} else {
		goto L262
	}
L262:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v560 != 0 {
		goto L2
	} else {
		goto L263
	}
L263:
	;
	v25 = v551
	v26 = v26 - int32(1)
	goto L11
L264:
	;
	v567 = v565
	goto L266
L265:
	;
	v567 = int32(15)
	goto L266
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v567
	goto L2
L267:
	;
	return
}
func F_replace_empty_jointree(m *base.Module, l0 int32) {
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v10 != 0 {
		m.G0 = v7 + int32(16)
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
		if v11 != 0 {
			m.G0 = v7 + int32(16)
			return
		} else {
			v13 = F_palloc0(m, int32(136))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(101)
				v21 = F_makeAlias(m, int32(_a_F_replace_empty_jointree_0), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v21
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					v25 = F_lappend(m, v24, v13)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v25
						if v25 != 0 {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
							v29 = v28
						} else {
							v29 = int32(0)
						}
						v31 = F_palloc0(m, int32(8))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v29
							*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(63)
							*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v31
							*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v31
							v41 = F_list_make1_impl(m, int32(1), v7+int32(8))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
								*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v41
								m.G0 = v7 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_replace_percent_placeholders(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	F_initStringInfo(m, v11+int32(-16))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = l0
	goto L3
L3:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v31 != int32(37) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L36
	}
L5:
	;
	goto L4
L6:
	;
	F_appendStringInfoChar(m, v11+int32(-16), int32(37))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L35
	}
L7:
	;
	v26 = v127 + int32(1)
	goto L3
L8:
	;
	F_appendStringInfoChar(m, v11+int32(-16), base.I32_extend8_s(v31))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L34
	}
L9:
	;
	if v31 != 0 {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v40 = v26 + int32(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	if v41 == int32(37) {
		goto L6
	} else {
		goto L13
	}
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	m.G0 = v13 - int32(-64)
	return v34
L13:
	;
	if v41 == int32(0) {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = l3
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v48 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v54 = v48
	v56 = l2
	v58 = int32(0)
	goto L18
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L29
	}
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v59 + int32(4)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v66 = base.B2i32(v63 != v54&int32(255))
	if v63 != v54&int32(255) {
		v76 = v58
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v76 != 0 {
		v127 = v40
		goto L7
	} else {
		goto L28
	}
L20:
	;
	if v63 != v54&int32(255) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v67 == int32(0) {
		v76 = v58
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_appendStringInfoString(m, v11+int32(-16), v67)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v76 = int32(1)
	goto L20
L24:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v77 != 0 {
		v54 = v77
		v56 = v56 + int32(1)
		v58 = v76
		goto L18
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L19
L27:
	;
	goto L26
L28:
	;
	goto L17
L29:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
	F_errmsg(m, int32(_a_F_replace_percent_placeholders_0), v11+int32(-48))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v106 = int32(*(*int8)(unsafe.Add(mBase, uint32(v40))))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v106
	v109 = F_errdetail(m, int32(_a_F_replace_percent_placeholders_1), v13)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_replace_percent_placeholders_2), int32(125), int32(_a_F_replace_percent_placeholders_3))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	v127 = v26
	goto L7
L35:
	;
	v26 = v26 + int32(2)
	goto L3
L36:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l1
	F_errmsg(m, int32(_a_F_replace_percent_placeholders_0), v11+int32(-32))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v156 = F_errdetail(m, int32(_a_F_replace_percent_placeholders_4), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_replace_percent_placeholders_2), int32(86), int32(_a_F_replace_percent_placeholders_3))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_resolve_aggregate_transtype(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 <= int32(3830) {
		switch l1 - int32(2277) {
		case 0, 6:
			F_get_func_signature(m, l0, v7+int32(12), v7+int32(8))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
				v40 = F_enforce_generic_type_consistency(m, l2, v37, v38, l1, int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					F_pfree(m, v42)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = v40
						m.G0 = v7 + int32(16)
						return v45
					}
				}
			}
		case 1, 2, 3, 4, 5:
			v45 = l1
			m.G0 = v7 + int32(16)
			return v45
		default:
			if base.B2i32(l1 == int32(2776))|base.B2i32(l1 == int32(3500)) != 0 {
				F_get_func_signature(m, l0, v7+int32(12), v7+int32(8))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
					v40 = F_enforce_generic_type_consistency(m, l2, v37, v38, l1, int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						F_pfree(m, v42)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							v45 = v40
							m.G0 = v7 + int32(16)
							return v45
						}
					}
				}
			} else {
				v45 = l1
				m.G0 = v7 + int32(16)
				return v45
			}
		}
	} else {
		if base.B2i32(base.Ui32(l1-int32(_a_F_resolve_aggregate_transtype_0)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l1-int32(_a_F_resolve_aggregate_transtype_1)) < base.Ui32(int32(2))) != 0 {
			F_get_func_signature(m, l0, v7+int32(12), v7+int32(8))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
				v40 = F_enforce_generic_type_consistency(m, l2, v37, v38, l1, int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					F_pfree(m, v42)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = v40
						m.G0 = v7 + int32(16)
						return v45
					}
				}
			}
		} else {
			if l1 != int32(3831) {
				v45 = l1
				m.G0 = v7 + int32(16)
				return v45
			} else {
				F_get_func_signature(m, l0, v7+int32(12), v7+int32(8))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
					v40 = F_enforce_generic_type_consistency(m, l2, v37, v38, l1, int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						F_pfree(m, v42)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							v45 = v40
							m.G0 = v7 + int32(16)
							return v45
						}
					}
				}
			}
		}
	}
}
func F_restriction_is_always_false(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v3 = int32(0)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)))
	if v6 != 0 {
		v61 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v61
L2:
	;
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v7 != 0 {
		v61 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v9 == int32(52) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v12 != 0 {
		v61 = v3
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	goto L11
L7:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)))
	if v13 != 0 {
		v61 = v3
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v16 = F_expr_is_nonnullable(m, l0, v14, int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	return v16
L11:
	;
	if base.B2i32(v21 != int32(0)) == int32(0) {
		v61 = v3
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v26 = int32(1)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v28 == int32(0) {
		v61 = v26
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v31 = int32(0)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v32 <= v31 {
		v61 = v26
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v36 = v31
	goto L15
L15:
	;
	v40 = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v36<<(uint(int32(2))%32))))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v46 != int32(320) {
		v61 = v40
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v61 = v53
	goto L1
L17:
	;
	v49 = F_restriction_is_always_false(m, l0, v45)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	if v49 == int32(0) {
		v61 = v40
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v53 = int32(1)
	v55 = v36 + v53
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v55 < v56 {
		v36 = v55
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
}
func F_restriction_is_or_clause(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	return base.B2i32(v2 != int32(0))
}
func F_revalidate_rectypeid(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
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
	var v48 int64
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v10 != int32(2249) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
		if v14 == int32(0) {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
			if v22 != 0 {
				F_typenameTypeIdAndMod(m, int32(0), v22, v13+int32(4), v13+int32(24))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v32 = F_lookup_type_cache(m, v30, int32(_a_F_revalidate_rectypeid_0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+13)))
						if v34 == int32(100) {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+300))
							v39 = F_lookup_type_cache(m, v37, int32(256))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								v41 = v39
								v43 = v13 + int32(4)
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+188))
								if v44 == int32(0) {
									F_errstart_cold(m, int32(21), int32(_a_F_revalidate_rectypeid_1))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										F_errcode(m, int32(151027844))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
											v68 = F_format_type_be(m, v67)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v68
												F_errmsg(m, int32(_a_F_revalidate_rectypeid_2), v8)
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_revalidate_rectypeid_3), int32(_a_F_revalidate_rectypeid_4), int32(_a_F_revalidate_rectypeid_5))
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
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
									*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v41
									v48 = *(*int64)(unsafe.Add(mBase, uint32(v41)+192))
									*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v48
									v51 = v43
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v52
									m.G0 = v8 + int32(16)
									return
								}
							}
						} else {
							v41 = v32
							v43 = v13 + int32(4)
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+188))
							if v44 == int32(0) {
								F_errstart_cold(m, int32(21), int32(_a_F_revalidate_rectypeid_1))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									F_errcode(m, int32(151027844))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
										v68 = F_format_type_be(m, v67)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v68
											F_errmsg(m, int32(_a_F_revalidate_rectypeid_2), v8)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_revalidate_rectypeid_3), int32(_a_F_revalidate_rectypeid_4), int32(_a_F_revalidate_rectypeid_5))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
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
								*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v41
								v48 = *(*int64)(unsafe.Add(mBase, uint32(v41)+192))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v48
								v51 = v43
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v52
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
				v32 = F_lookup_type_cache(m, v30, int32(_a_F_revalidate_rectypeid_0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+13)))
					if v34 == int32(100) {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+300))
						v39 = F_lookup_type_cache(m, v37, int32(256))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							v41 = v39
							v43 = v13 + int32(4)
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+188))
							if v44 == int32(0) {
								F_errstart_cold(m, int32(21), int32(_a_F_revalidate_rectypeid_1))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									F_errcode(m, int32(151027844))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
										v68 = F_format_type_be(m, v67)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v68
											F_errmsg(m, int32(_a_F_revalidate_rectypeid_2), v8)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_revalidate_rectypeid_3), int32(_a_F_revalidate_rectypeid_4), int32(_a_F_revalidate_rectypeid_5))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
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
								*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v41
								v48 = *(*int64)(unsafe.Add(mBase, uint32(v41)+192))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v48
								v51 = v43
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v52
								m.G0 = v8 + int32(16)
								return
							}
						}
					} else {
						v41 = v32
						v43 = v13 + int32(4)
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+188))
						if v44 == int32(0) {
							F_errstart_cold(m, int32(21), int32(_a_F_revalidate_rectypeid_1))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								F_errcode(m, int32(151027844))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
									v68 = F_format_type_be(m, v67)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v68
										F_errmsg(m, int32(_a_F_revalidate_rectypeid_2), v8)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_revalidate_rectypeid_3), int32(_a_F_revalidate_rectypeid_4), int32(_a_F_revalidate_rectypeid_5))
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v41
							v48 = *(*int64)(unsafe.Add(mBase, uint32(v41)+192))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v48
							v51 = v43
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v52
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		} else {
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v14)+192))
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
			if v17 != v18 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
				if v22 != 0 {
					F_typenameTypeIdAndMod(m, int32(0), v22, v13+int32(4), v13+int32(24))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
						v32 = F_lookup_type_cache(m, v30, int32(_a_F_revalidate_rectypeid_0))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+13)))
							if v34 == int32(100) {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+300))
								v39 = F_lookup_type_cache(m, v37, int32(256))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									v41 = v39
									v43 = v13 + int32(4)
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+188))
									if v44 == int32(0) {
										F_errstart_cold(m, int32(21), int32(_a_F_revalidate_rectypeid_1))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return
										} else {
											F_errcode(m, int32(151027844))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return
											} else {
												v67 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
												v68 = F_format_type_be(m, v67)
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8))) = v68
													F_errmsg(m, int32(_a_F_revalidate_rectypeid_2), v8)
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_revalidate_rectypeid_3), int32(_a_F_revalidate_rectypeid_4), int32(_a_F_revalidate_rectypeid_5))
														mBase = m.M
														v78 = m.ExcPending
														if v78 != 0 {
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
										*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v41
										v48 = *(*int64)(unsafe.Add(mBase, uint32(v41)+192))
										*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v48
										v51 = v43
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v52
										m.G0 = v8 + int32(16)
										return
									}
								}
							} else {
								v41 = v32
								v43 = v13 + int32(4)
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+188))
								if v44 == int32(0) {
									F_errstart_cold(m, int32(21), int32(_a_F_revalidate_rectypeid_1))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										F_errcode(m, int32(151027844))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
											v68 = F_format_type_be(m, v67)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v68
												F_errmsg(m, int32(_a_F_revalidate_rectypeid_2), v8)
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_revalidate_rectypeid_3), int32(_a_F_revalidate_rectypeid_4), int32(_a_F_revalidate_rectypeid_5))
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
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
									*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v41
									v48 = *(*int64)(unsafe.Add(mBase, uint32(v41)+192))
									*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v48
									v51 = v43
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v52
									m.G0 = v8 + int32(16)
									return
								}
							}
						}
					}
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v32 = F_lookup_type_cache(m, v30, int32(_a_F_revalidate_rectypeid_0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+13)))
						if v34 == int32(100) {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+300))
							v39 = F_lookup_type_cache(m, v37, int32(256))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								v41 = v39
								v43 = v13 + int32(4)
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+188))
								if v44 == int32(0) {
									F_errstart_cold(m, int32(21), int32(_a_F_revalidate_rectypeid_1))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										F_errcode(m, int32(151027844))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
											v68 = F_format_type_be(m, v67)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v68
												F_errmsg(m, int32(_a_F_revalidate_rectypeid_2), v8)
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_revalidate_rectypeid_3), int32(_a_F_revalidate_rectypeid_4), int32(_a_F_revalidate_rectypeid_5))
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
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
									*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v41
									v48 = *(*int64)(unsafe.Add(mBase, uint32(v41)+192))
									*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v48
									v51 = v43
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v52
									m.G0 = v8 + int32(16)
									return
								}
							}
						} else {
							v41 = v32
							v43 = v13 + int32(4)
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+188))
							if v44 == int32(0) {
								F_errstart_cold(m, int32(21), int32(_a_F_revalidate_rectypeid_1))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									F_errcode(m, int32(151027844))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
										v68 = F_format_type_be(m, v67)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v68
											F_errmsg(m, int32(_a_F_revalidate_rectypeid_2), v8)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_revalidate_rectypeid_3), int32(_a_F_revalidate_rectypeid_4), int32(_a_F_revalidate_rectypeid_5))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
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
								*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v41
								v48 = *(*int64)(unsafe.Add(mBase, uint32(v41)+192))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v48
								v51 = v43
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v52
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			} else {
				v51 = v13 + int32(4)
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v52
				m.G0 = v8 + int32(16)
				return
			}
		}
	} else {
		m.G0 = v8 + int32(16)
		return
	}
}
func F_rlocator_comparator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v6) < base.Ui32(v7) {
		return int32(-1)
	} else {
		v11 = int32(1)
		if base.Ui32(v7) < base.Ui32(v6) {
			v28 = v11
			return v28
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if base.Ui32(v13) < base.Ui32(v14) {
				return int32(-1)
			} else {
				if base.Ui32(v14) < base.Ui32(v13) {
					v28 = v11
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if base.Ui32(v20) < base.Ui32(v21) {
						v28 = int32(-1)
					} else {
						v28 = base.B2i32(base.Ui32(v21) < base.Ui32(v20))
					}
				}
				return v28
			}
		}
	}
}
func F_roles_is_member_of(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
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
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	v5 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v5
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v27 = l1 << (uint(int32(2)) % 32)
	if l2 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v20 + int32(32)
	return v270
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_roles_is_member_of[0]))
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_roles_is_member_of[1])))
	if base.B2i32(v28 != l0)|base.B2i32(v28 == int32(0)) != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_roles_is_member_of[2])))
	v270 = v33
	goto L4
L8:
	;
	v245 = int32(_a_F_roles_is_member_of_0)
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_roles_is_member_of[3]))
	v249 = *(*int32)(unsafe.Add(mBase, _c_F_roles_is_member_of[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_roles_is_member_of[3])) = v249
	v251 = F_list_copy(m, v235)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L14
	} else {
		goto L54
	}
L9:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	if v223 == int32(0) {
		v235 = v213
		goto L8
	} else {
		goto L52
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L14
	} else {
		goto L49
	}
L11:
	;
	v39 = F_SearchSysCache1(m, int32(21), base.I64_extend_i32_u(v36))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v53 = v5
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = l0
	v61 = F_list_make1_impl(m, int32(480), v20+int32(12))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L14
	} else {
		goto L18
	}
L14:
	;
	return int32(0)
L15:
	;
	if v39 == int32(0) {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+22)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45+v46)+68))
	F_ReleaseCatCache(m, v39)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v53 = v48
	goto L13
L18:
	;
	if v61 == int32(0) {
		v235 = v5
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v65 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v213 = v61
	goto L9
L21:
	;
	goto L22
L22:
	;
	v77 = v61
	v82 = v5
	goto L23
L23:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v82<<(uint(int32(2))%32))))
	v95 = int64(0)
	v97 = F_SearchSysCacheList(m, int32(8), int32(1), base.I64_extend_i32_u(v93), v95, v95)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L14
	} else {
		goto L25
	}
L24:
	;
	v213 = v184
	goto L9
L25:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)+56))
	if int32(0) < v99 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v109 = int32(0)
	v112 = v77
	goto L29
L27:
	;
	v161 = v77
	goto L28
L28:
	;
	F_ReleaseCatCacheList(m, v97)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L14
	} else {
		goto L43
	}
L29:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v97-int32(-64)+v109<<(uint(int32(2))%32))))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+72))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+22)))
	v130 = v128 + v129
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	if base.B2i32(l2 == int32(0))|base.B2i32(v131 != l2) != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v161 = v149
	goto L28
L31:
	;
	switch l1 - int32(1) {
	case 0:
		goto L38
	case 1:
		goto L37
	default:
		goto L36
	}
L32:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+16)))
	if v134&int32(1) == int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v139 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v93
	goto L31
L35:
	;
	v151 = v109 + int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v97)+56))
	if v151 < v152 {
		v109 = v151
		v112 = v149
		goto L29
	} else {
		goto L42
	}
L36:
	;
	v147 = F_roles_list_append(m, v112, v20+int32(28), v131)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L14
	} else {
		goto L41
	}
L37:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+18)))
	if v142 != int32(1) {
		v149 = v112
		goto L35
	} else {
		goto L40
	}
L38:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+17)))
	if v141 != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v149 = v112
	goto L35
L40:
	;
	goto L36
L41:
	;
	v149 = v147
	goto L35
L42:
	;
	goto L30
L43:
	;
	v173 = int32(0)
	if base.B2i32(v53 == v173)|base.B2i32(v53 != v93) == v173 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v182 = F_roles_list_append(m, v161, v20+int32(28), int32(_a_F_roles_is_member_of_1))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L14
	} else {
		goto L47
	}
L45:
	;
	v184 = v161
	goto L46
L46:
	;
	v186 = v82 + int32(1)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v186 < v187 {
		v77 = v184
		v82 = v186
		goto L23
	} else {
		goto L48
	}
L47:
	;
	v184 = v182
	goto L46
L48:
	;
	goto L24
L49:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_roles_is_member_of[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v194
	F_errmsg_internal(m, int32(_a_F_roles_is_member_of_2), v20+int32(16))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L14
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_roles_is_member_of_3), int32(_a_F_roles_is_member_of_4), int32(_a_F_roles_is_member_of_5))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L14
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	F_pfree(m, v223)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L14
	} else {
		goto L53
	}
L53:
	;
	v235 = v213
	goto L8
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_roles_is_member_of[3])) = v246
	F_list_free(m, v235)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L14
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_roles_is_member_of[1]))) = int32(0)
	v260 = l1 << (uint(int32(2)) % 32)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+uint32(_c_F_roles_is_member_of[2])))
	F_list_free(m, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L14
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260)+uint32(_c_F_roles_is_member_of[2]))) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_roles_is_member_of[1]))) = l0
	v270 = v251
	goto L4
}
func F_rtree_internal_consistent(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v110 int64
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	switch l2 - int32(1) {
	case 0:
		v110 = F_DirectFunctionCall2Coll(m, int32(106), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
		mBase = m.M
		v111 = m.ExcPending
		if v111 != 0 {
			return int32(0)
		} else {
			v114 = base.B2i32(v110 == int64(0))
			m.G0 = v7 + int32(16)
			return v114
		}
	case 1:
		v15 = F_DirectFunctionCall2Coll(m, int32(97), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v114 = base.B2i32(v15 == int64(0))
			m.G0 = v7 + int32(16)
			return v114
		}
	case 2:
		v25 = F_DirectFunctionCall2Coll(m, int32(98), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v114 = base.B2i32(v25 != int64(0))
			m.G0 = v7 + int32(16)
			return v114
		}
	case 3:
		v33 = F_DirectFunctionCall2Coll(m, int32(99), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			v114 = base.B2i32(v33 == int64(0))
			m.G0 = v7 + int32(16)
			return v114
		}
	case 4:
		v41 = F_DirectFunctionCall2Coll(m, int32(100), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			v114 = base.B2i32(v41 == int64(0))
			m.G0 = v7 + int32(16)
			return v114
		}
	case 5, 6:
		v49 = F_DirectFunctionCall2Coll(m, int32(101), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			v114 = base.B2i32(v49 != int64(0))
			m.G0 = v7 + int32(16)
			return v114
		}
	case 7:
		v57 = F_DirectFunctionCall2Coll(m, int32(98), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return int32(0)
		} else {
			v114 = base.B2i32(v57 != int64(0))
			m.G0 = v7 + int32(16)
			return v114
		}
	case 8:
		v65 = F_DirectFunctionCall2Coll(m, int32(102), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return int32(0)
		} else {
			v114 = base.B2i32(v65 == int64(0))
			m.G0 = v7 + int32(16)
			return v114
		}
	case 9:
		v73 = F_DirectFunctionCall2Coll(m, int32(103), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return int32(0)
		} else {
			v114 = base.B2i32(v73 == int64(0))
			m.G0 = v7 + int32(16)
			return v114
		}
	case 10:
		v81 = F_DirectFunctionCall2Coll(m, int32(104), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
		mBase = m.M
		v82 = m.ExcPending
		if v82 != 0 {
			return int32(0)
		} else {
			v114 = base.B2i32(v81 == int64(0))
			m.G0 = v7 + int32(16)
			return v114
		}
	case 11:
		v89 = F_DirectFunctionCall2Coll(m, int32(105), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
		mBase = m.M
		v90 = m.ExcPending
		if v90 != 0 {
			return int32(0)
		} else {
			v114 = base.B2i32(v89 == int64(0))
			m.G0 = v7 + int32(16)
			return v114
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v96 = m.ExcPending
		if v96 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l2
			F_errmsg_internal(m, int32(_a_F_rtree_internal_consistent_0), v7)
			mBase = m.M
			v100 = m.ExcPending
			if v100 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_rtree_internal_consistent_1), int32(1021), int32(_a_F_rtree_internal_consistent_2))
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
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
