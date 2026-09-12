package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReadControlFile(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 float64
	_ = v78
	var v81 int32
	_ = v81
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
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v159 float64
	_ = v159
	var v162 float64
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	v5 = m.G0
	v7 = v5 - int32(320)
	m.G0 = v7
	v11 = F_BasicOpenFile(m, int32(286739), int32(2))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		if int32(0) <= v11 {
			v16 = *(*int32)(unsafe.Add(mBase, _consts[96]))
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(167772169)
			v20 = *(*int32)(unsafe.Add(mBase, _consts[120]))
			v21 = int32(296)
			v22 = F_read(m, v11, v20, v21)
			mBase = m.M
			if v22 != v21 {
				F_errstart_cold(m, int32(23), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					if v22 < int32(0) {
						F_errcode_for_file_access(m)
						mBase = m.M
						v204 = m.ExcPending
						if v204 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+256)) = int32(286739)
							F_errmsg(m, int32(284935), v7+int32(256))
							mBase = m.M
							v211 = m.ExcPending
							if v211 != 0 {
								return
							} else {
								F_errfinish(m, int32(475016), int32(4372), int32(370959))
								mBase = m.M
								v216 = m.ExcPending
								if v216 != 0 {
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
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+280)) = int32(296)
							*(*int32)(unsafe.Add(mBase, uint32(v7)+276)) = v22
							*(*int32)(unsafe.Add(mBase, uint32(v7)+272)) = int32(286739)
							F_errmsg(m, int32(34779), v7+int32(272))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								F_errfinish(m, int32(475016), int32(4377), int32(370959))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
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
				v50 = *(*int32)(unsafe.Add(mBase, _consts[96]))
				v51 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v50))) = v51
				v53 = F_close(m, v11)
				mBase = m.M
				v56 = *(*int32)(unsafe.Add(mBase, _consts[120]))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
				if v57&int32(65535) != 0 {
					v60 = v51
				} else {
					v60 = v57
				}
				if v60 != 0 {
					F_errstart_cold(m, int32(22), int32(0))
					mBase = m.M
					v220 = m.ExcPending
					if v220 != 0 {
						return
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v223 = m.ExcPending
						if v223 != 0 {
							return
						} else {
							F_errmsg(m, int32(204068), int32(0))
							mBase = m.M
							v227 = m.ExcPending
							if v227 != 0 {
								return
							} else {
								v229 = *(*int32)(unsafe.Add(mBase, _consts[120]))
								v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v7)+240)) = v230
								*(*int32)(unsafe.Add(mBase, uint32(v7)+244)) = v230
								*(*int64)(unsafe.Add(mBase, uint32(v7)+248)) = int64(7730941134600)
								F_errdetail(m, int32(619310), v7+int32(240))
								mBase = m.M
								v239 = m.ExcPending
								if v239 != 0 {
									return
								} else {
									F_errhint(m, int32(613465), int32(0))
									mBase = m.M
									v243 = m.ExcPending
									if v243 != 0 {
										return
									} else {
										F_errfinish(m, int32(475016), int32(4398), int32(370959))
										mBase = m.M
										v248 = m.ExcPending
										if v248 != 0 {
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
					if v57 != int32(1800) {
						F_errstart_cold(m, int32(22), int32(0))
						mBase = m.M
						v252 = m.ExcPending
						if v252 != 0 {
							return
						} else {
							F_errcode(m, int32(325))
							mBase = m.M
							v255 = m.ExcPending
							if v255 != 0 {
								return
							} else {
								F_errmsg(m, int32(204068), int32(0))
								mBase = m.M
								v259 = m.ExcPending
								if v259 != 0 {
									return
								} else {
									v261 = *(*int32)(unsafe.Add(mBase, _consts[120]))
									v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v7)+224)) = v262
									*(*int32)(unsafe.Add(mBase, uint32(v7)+228)) = int32(1800)
									F_errdetail(m, int32(613263), v7+int32(224))
									mBase = m.M
									v270 = m.ExcPending
									if v270 != 0 {
										return
									} else {
										F_errhint(m, int32(613519), int32(0))
										mBase = m.M
										v274 = m.ExcPending
										if v274 != 0 {
											return
										} else {
											F_errfinish(m, int32(475016), int32(4407), int32(370959))
											mBase = m.M
											v279 = m.ExcPending
											if v279 != 0 {
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
						v63 = int32(-1)
						v65 = m.Env.Pgmem_crc32c(m, v63, v56, int32(292))
						mBase = m.M
						v67 = *(*int32)(unsafe.Add(mBase, _consts[120]))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+292))
						if v65^v68 != v63 {
							F_errstart_cold(m, int32(22), int32(0))
							mBase = m.M
							v283 = m.ExcPending
							if v283 != 0 {
								return
							} else {
								F_errcode(m, int32(325))
								mBase = m.M
								v286 = m.ExcPending
								if v286 != 0 {
									return
								} else {
									F_errmsg(m, int32(369163), int32(0))
									mBase = m.M
									v290 = m.ExcPending
									if v290 != 0 {
										return
									} else {
										F_errfinish(m, int32(475016), int32(4419), int32(370959))
										mBase = m.M
										v295 = m.ExcPending
										if v295 != 0 {
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
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
							if v72 != int32(202506291) {
								F_errstart_cold(m, int32(22), int32(0))
								mBase = m.M
								v299 = m.ExcPending
								if v299 != 0 {
									return
								} else {
									F_errcode(m, int32(325))
									mBase = m.M
									v302 = m.ExcPending
									if v302 != 0 {
										return
									} else {
										F_errmsg(m, int32(204068), int32(0))
										mBase = m.M
										v306 = m.ExcPending
										if v306 != 0 {
											return
										} else {
											v308 = *(*int32)(unsafe.Add(mBase, _consts[120]))
											v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+12))
											v310 = int32(503349)
											*(*int32)(unsafe.Add(mBase, uint32(v7)+208)) = v310
											*(*int32)(unsafe.Add(mBase, uint32(v7)+212)) = v309
											*(*int32)(unsafe.Add(mBase, uint32(v7)+216)) = v310
											*(*int32)(unsafe.Add(mBase, uint32(v7)+220)) = int32(202506291)
											F_errdetail(m, int32(612512), v7+int32(208))
											mBase = m.M
											v321 = m.ExcPending
											if v321 != 0 {
												return
											} else {
												F_errhint(m, int32(613519), int32(0))
												mBase = m.M
												v325 = m.ExcPending
												if v325 != 0 {
													return
												} else {
													F_errfinish(m, int32(475016), int32(4435), int32(370959))
													mBase = m.M
													v330 = m.ExcPending
													if v330 != 0 {
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
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v67)+204))
								if v75 != int32(8) {
									F_errstart_cold(m, int32(22), int32(0))
									mBase = m.M
									v334 = m.ExcPending
									if v334 != 0 {
										return
									} else {
										F_errcode(m, int32(325))
										mBase = m.M
										v337 = m.ExcPending
										if v337 != 0 {
											return
										} else {
											F_errmsg(m, int32(204068), int32(0))
											mBase = m.M
											v341 = m.ExcPending
											if v341 != 0 {
												return
											} else {
												v343 = *(*int32)(unsafe.Add(mBase, _consts[120]))
												v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+204))
												v345 = int32(506949)
												*(*int32)(unsafe.Add(mBase, uint32(v7)+192)) = v345
												*(*int32)(unsafe.Add(mBase, uint32(v7)+196)) = v344
												*(*int32)(unsafe.Add(mBase, uint32(v7)+200)) = v345
												*(*int32)(unsafe.Add(mBase, uint32(v7)+204)) = int32(8)
												F_errdetail(m, int32(612512), v7+int32(192))
												mBase = m.M
												v356 = m.ExcPending
												if v356 != 0 {
													return
												} else {
													F_errhint(m, int32(613519), int32(0))
													mBase = m.M
													v360 = m.ExcPending
													if v360 != 0 {
														return
													} else {
														F_errfinish(m, int32(475016), int32(4445), int32(370959))
														mBase = m.M
														v365 = m.ExcPending
														if v365 != 0 {
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
									v78 = *(*float64)(unsafe.Add(mBase, uint32(v67)+208))
									if base.F64_ne(v78, float64(1.234567e+06)) != 0 {
										F_errstart_cold(m, int32(22), int32(0))
										mBase = m.M
										v369 = m.ExcPending
										if v369 != 0 {
											return
										} else {
											F_errcode(m, int32(325))
											mBase = m.M
											v372 = m.ExcPending
											if v372 != 0 {
												return
											} else {
												F_errmsg(m, int32(204068), int32(0))
												mBase = m.M
												v376 = m.ExcPending
												if v376 != 0 {
													return
												} else {
													F_errdetail(m, int32(598002), int32(0))
													mBase = m.M
													v380 = m.ExcPending
													if v380 != 0 {
														return
													} else {
														F_errhint(m, int32(613519), int32(0))
														mBase = m.M
														v384 = m.ExcPending
														if v384 != 0 {
															return
														} else {
															F_errfinish(m, int32(475016), int32(4451), int32(370959))
															mBase = m.M
															v389 = m.ExcPending
															if v389 != 0 {
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
										v81 = *(*int32)(unsafe.Add(mBase, uint32(v67)+216))
										if v81 != int32(8192) {
											F_errstart_cold(m, int32(22), int32(0))
											mBase = m.M
											v393 = m.ExcPending
											if v393 != 0 {
												return
											} else {
												F_errcode(m, int32(325))
												mBase = m.M
												v396 = m.ExcPending
												if v396 != 0 {
													return
												} else {
													F_errmsg(m, int32(204068), int32(0))
													mBase = m.M
													v400 = m.ExcPending
													if v400 != 0 {
														return
													} else {
														v402 = *(*int32)(unsafe.Add(mBase, _consts[120]))
														v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)+216))
														v404 = int32(484655)
														*(*int32)(unsafe.Add(mBase, uint32(v7)+176)) = v404
														*(*int32)(unsafe.Add(mBase, uint32(v7)+180)) = v403
														*(*int32)(unsafe.Add(mBase, uint32(v7)+184)) = v404
														*(*int32)(unsafe.Add(mBase, uint32(v7)+188)) = int32(8192)
														F_errdetail(m, int32(612512), v7+int32(176))
														mBase = m.M
														v415 = m.ExcPending
														if v415 != 0 {
															return
														} else {
															F_errhint(m, int32(613418), int32(0))
															mBase = m.M
															v419 = m.ExcPending
															if v419 != 0 {
																return
															} else {
																F_errfinish(m, int32(475016), int32(4461), int32(370959))
																mBase = m.M
																v424 = m.ExcPending
																if v424 != 0 {
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
											v84 = *(*int32)(unsafe.Add(mBase, uint32(v67)+220))
											if v84 != int32(131072) {
												F_errstart_cold(m, int32(22), int32(0))
												mBase = m.M
												v428 = m.ExcPending
												if v428 != 0 {
													return
												} else {
													F_errcode(m, int32(325))
													mBase = m.M
													v431 = m.ExcPending
													if v431 != 0 {
														return
													} else {
														F_errmsg(m, int32(204068), int32(0))
														mBase = m.M
														v435 = m.ExcPending
														if v435 != 0 {
															return
														} else {
															v437 = *(*int32)(unsafe.Add(mBase, _consts[120]))
															v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)+220))
															v439 = int32(513498)
															*(*int32)(unsafe.Add(mBase, uint32(v7)+160)) = v439
															*(*int32)(unsafe.Add(mBase, uint32(v7)+164)) = v438
															*(*int32)(unsafe.Add(mBase, uint32(v7)+168)) = v439
															*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = int32(131072)
															F_errdetail(m, int32(612512), v7+int32(160))
															mBase = m.M
															v450 = m.ExcPending
															if v450 != 0 {
																return
															} else {
																F_errhint(m, int32(613418), int32(0))
																mBase = m.M
																v454 = m.ExcPending
																if v454 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(475016), int32(4471), int32(370959))
																	mBase = m.M
																	v459 = m.ExcPending
																	if v459 != 0 {
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
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v67)+224))
												if v87 != int32(8192) {
													F_errstart_cold(m, int32(22), int32(0))
													mBase = m.M
													v463 = m.ExcPending
													if v463 != 0 {
														return
													} else {
														F_errcode(m, int32(325))
														mBase = m.M
														v466 = m.ExcPending
														if v466 != 0 {
															return
														} else {
															F_errmsg(m, int32(204068), int32(0))
															mBase = m.M
															v470 = m.ExcPending
															if v470 != 0 {
																return
															} else {
																v472 = *(*int32)(unsafe.Add(mBase, _consts[120]))
																v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)+224))
																v474 = int32(484650)
																*(*int32)(unsafe.Add(mBase, uint32(v7)+144)) = v474
																*(*int32)(unsafe.Add(mBase, uint32(v7)+148)) = v473
																*(*int32)(unsafe.Add(mBase, uint32(v7)+152)) = v474
																*(*int32)(unsafe.Add(mBase, uint32(v7)+156)) = int32(8192)
																F_errdetail(m, int32(612512), v7+int32(144))
																mBase = m.M
																v485 = m.ExcPending
																if v485 != 0 {
																	return
																} else {
																	F_errhint(m, int32(613418), int32(0))
																	mBase = m.M
																	v489 = m.ExcPending
																	if v489 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(475016), int32(4481), int32(370959))
																		mBase = m.M
																		v494 = m.ExcPending
																		if v494 != 0 {
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
													v90 = *(*int32)(unsafe.Add(mBase, uint32(v67)+232))
													if v90 != int32(64) {
														F_errstart_cold(m, int32(22), int32(0))
														mBase = m.M
														v498 = m.ExcPending
														if v498 != 0 {
															return
														} else {
															F_errcode(m, int32(325))
															mBase = m.M
															v501 = m.ExcPending
															if v501 != 0 {
																return
															} else {
																F_errmsg(m, int32(204068), int32(0))
																mBase = m.M
																v505 = m.ExcPending
																if v505 != 0 {
																	return
																} else {
																	v507 = *(*int32)(unsafe.Add(mBase, _consts[120]))
																	v508 = *(*int32)(unsafe.Add(mBase, uint32(v507)+232))
																	v509 = int32(507015)
																	*(*int32)(unsafe.Add(mBase, uint32(v7)+128)) = v509
																	*(*int32)(unsafe.Add(mBase, uint32(v7)+132)) = v508
																	*(*int32)(unsafe.Add(mBase, uint32(v7)+136)) = v509
																	*(*int32)(unsafe.Add(mBase, uint32(v7)+140)) = int32(64)
																	F_errdetail(m, int32(612512), v7+int32(128))
																	mBase = m.M
																	v520 = m.ExcPending
																	if v520 != 0 {
																		return
																	} else {
																		F_errhint(m, int32(613418), int32(0))
																		mBase = m.M
																		v524 = m.ExcPending
																		if v524 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(475016), int32(4491), int32(370959))
																			mBase = m.M
																			v529 = m.ExcPending
																			if v529 != 0 {
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
														v93 = *(*int32)(unsafe.Add(mBase, uint32(v67)+236))
														if v93 != int32(32) {
															F_errstart_cold(m, int32(22), int32(0))
															mBase = m.M
															v533 = m.ExcPending
															if v533 != 0 {
																return
															} else {
																F_errcode(m, int32(325))
																mBase = m.M
																v536 = m.ExcPending
																if v536 != 0 {
																	return
																} else {
																	F_errmsg(m, int32(204068), int32(0))
																	mBase = m.M
																	v540 = m.ExcPending
																	if v540 != 0 {
																		return
																	} else {
																		v542 = *(*int32)(unsafe.Add(mBase, _consts[120]))
																		v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)+236))
																		v544 = int32(498690)
																		*(*int32)(unsafe.Add(mBase, uint32(v7)+112)) = v544
																		*(*int32)(unsafe.Add(mBase, uint32(v7)+116)) = v543
																		*(*int32)(unsafe.Add(mBase, uint32(v7)+120)) = v544
																		*(*int32)(unsafe.Add(mBase, uint32(v7)+124)) = int32(32)
																		F_errdetail(m, int32(612512), v7+int32(112))
																		mBase = m.M
																		v555 = m.ExcPending
																		if v555 != 0 {
																			return
																		} else {
																			F_errhint(m, int32(613418), int32(0))
																			mBase = m.M
																			v559 = m.ExcPending
																			if v559 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(475016), int32(4501), int32(370959))
																				mBase = m.M
																				v564 = m.ExcPending
																				if v564 != 0 {
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
															v96 = *(*int32)(unsafe.Add(mBase, uint32(v67)+240))
															if v96 != int32(1996) {
																F_errstart_cold(m, int32(22), int32(0))
																mBase = m.M
																v568 = m.ExcPending
																if v568 != 0 {
																	return
																} else {
																	F_errcode(m, int32(325))
																	mBase = m.M
																	v571 = m.ExcPending
																	if v571 != 0 {
																		return
																	} else {
																		F_errmsg(m, int32(204068), int32(0))
																		mBase = m.M
																		v575 = m.ExcPending
																		if v575 != 0 {
																			return
																		} else {
																			v577 = *(*int32)(unsafe.Add(mBase, _consts[120]))
																			v578 = *(*int32)(unsafe.Add(mBase, uint32(v577)+240))
																			v579 = int32(513477)
																			*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = v579
																			*(*int32)(unsafe.Add(mBase, uint32(v7)+100)) = v578
																			*(*int32)(unsafe.Add(mBase, uint32(v7)+104)) = v579
																			*(*int32)(unsafe.Add(mBase, uint32(v7)+108)) = int32(1996)
																			F_errdetail(m, int32(612512), v7+int32(96))
																			mBase = m.M
																			v590 = m.ExcPending
																			if v590 != 0 {
																				return
																			} else {
																				F_errhint(m, int32(613418), int32(0))
																				mBase = m.M
																				v594 = m.ExcPending
																				if v594 != 0 {
																					return
																				} else {
																					F_errfinish(m, int32(475016), int32(4511), int32(370959))
																					mBase = m.M
																					v599 = m.ExcPending
																					if v599 != 0 {
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
																v99 = *(*int32)(unsafe.Add(mBase, uint32(v67)+244))
																if v99 != int32(2048) {
																	F_errstart_cold(m, int32(22), int32(0))
																	mBase = m.M
																	v603 = m.ExcPending
																	if v603 != 0 {
																		return
																	} else {
																		F_errcode(m, int32(325))
																		mBase = m.M
																		v606 = m.ExcPending
																		if v606 != 0 {
																			return
																		} else {
																			F_errmsg(m, int32(204068), int32(0))
																			mBase = m.M
																			v610 = m.ExcPending
																			if v610 != 0 {
																				return
																			} else {
																				v612 = *(*int32)(unsafe.Add(mBase, _consts[120]))
																				v613 = *(*int32)(unsafe.Add(mBase, uint32(v612)+244))
																				v614 = int32(513510)
																				*(*int32)(unsafe.Add(mBase, uint32(v7)+80)) = v614
																				*(*int32)(unsafe.Add(mBase, uint32(v7)+84)) = v613
																				*(*int32)(unsafe.Add(mBase, uint32(v7)+88)) = v614
																				*(*int32)(unsafe.Add(mBase, uint32(v7)+92)) = int32(2048)
																				F_errdetail(m, int32(612512), v7+int32(80))
																				mBase = m.M
																				v625 = m.ExcPending
																				if v625 != 0 {
																					return
																				} else {
																					F_errhint(m, int32(613418), int32(0))
																					mBase = m.M
																					v629 = m.ExcPending
																					if v629 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(475016), int32(4521), int32(370959))
																						mBase = m.M
																						v634 = m.ExcPending
																						if v634 != 0 {
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
																	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+248)))
																	if v102 == int32(1) {
																		F_errstart_cold(m, int32(22), int32(0))
																		mBase = m.M
																		v638 = m.ExcPending
																		if v638 != 0 {
																			return
																		} else {
																			F_errcode(m, int32(325))
																			mBase = m.M
																			v641 = m.ExcPending
																			if v641 != 0 {
																				return
																			} else {
																				F_errmsg(m, int32(204068), int32(0))
																				mBase = m.M
																				v645 = m.ExcPending
																				if v645 != 0 {
																					return
																				} else {
																					F_errdetail(m, int32(617211), int32(0))
																					mBase = m.M
																					v649 = m.ExcPending
																					if v649 != 0 {
																						return
																					} else {
																						F_errhint(m, int32(613418), int32(0))
																						mBase = m.M
																						v653 = m.ExcPending
																						if v653 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(475016), int32(4538), int32(370959))
																							mBase = m.M
																							v658 = m.ExcPending
																							if v658 != 0 {
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
																		v106 = *(*int32)(unsafe.Add(mBase, uint32(v67)+228))
																		*(*int32)(unsafe.Add(mBase, _consts[116])) = v106
																		if v106 <= int32(0) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v662 = m.ExcPending
																			if v662 != 0 {
																				return
																			} else {
																				F_errcode(m, int32(50856066))
																				mBase = m.M
																				v665 = m.ExcPending
																				if v665 != 0 {
																					return
																				} else {
																					v667 = *(*int32)(unsafe.Add(mBase, _consts[116]))
																					*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v667
																					F_errmsg_plural(m, int32(635405), int32(632656), v667, v7-int32(-64))
																					mBase = m.M
																					v674 = m.ExcPending
																					if v674 != 0 {
																						return
																					} else {
																						F_errdetail(m, int32(618533), int32(0))
																						mBase = m.M
																						v678 = m.ExcPending
																						if v678 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(475016), int32(4549), int32(370959))
																							mBase = m.M
																							v683 = m.ExcPending
																							if v683 != 0 {
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
																			if v106&(v106-int32(1)) != 0 {
																				F_errstart_cold(m, int32(21), int32(0))
																				mBase = m.M
																				v662 = m.ExcPending
																				if v662 != 0 {
																					return
																				} else {
																					F_errcode(m, int32(50856066))
																					mBase = m.M
																					v665 = m.ExcPending
																					if v665 != 0 {
																						return
																					} else {
																						v667 = *(*int32)(unsafe.Add(mBase, _consts[116]))
																						*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v667
																						F_errmsg_plural(m, int32(635405), int32(632656), v667, v7-int32(-64))
																						mBase = m.M
																						v674 = m.ExcPending
																						if v674 != 0 {
																							return
																						} else {
																							F_errdetail(m, int32(618533), int32(0))
																							mBase = m.M
																							v678 = m.ExcPending
																							if v678 != 0 {
																								return
																							} else {
																								F_errfinish(m, int32(475016), int32(4549), int32(370959))
																								mBase = m.M
																								v683 = m.ExcPending
																								if v683 != 0 {
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
																				if base.Ui32(int32(1072693249)) <= base.Ui32(v106+int32(-1048576)) {
																					F_errstart_cold(m, int32(21), int32(0))
																					mBase = m.M
																					v662 = m.ExcPending
																					if v662 != 0 {
																						return
																					} else {
																						F_errcode(m, int32(50856066))
																						mBase = m.M
																						v665 = m.ExcPending
																						if v665 != 0 {
																							return
																						} else {
																							v667 = *(*int32)(unsafe.Add(mBase, _consts[116]))
																							*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v667
																							F_errmsg_plural(m, int32(635405), int32(632656), v667, v7-int32(-64))
																							mBase = m.M
																							v674 = m.ExcPending
																							if v674 != 0 {
																								return
																							} else {
																								F_errdetail(m, int32(618533), int32(0))
																								mBase = m.M
																								v678 = m.ExcPending
																								if v678 != 0 {
																									return
																								} else {
																									F_errfinish(m, int32(475016), int32(4549), int32(370959))
																									mBase = m.M
																									v683 = m.ExcPending
																									if v683 != 0 {
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
																					*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v106
																					v124 = F_pg_snprintf(m, v7+int32(288), int32(20), int32(465932), v7+int32(48))
																					mBase = m.M
																					v125 = m.ExcPending
																					if v125 != 0 {
																						return
																					} else {
																						F_SetConfigOption(m, int32(323890), v7+int32(288), int32(0), int32(1))
																						mBase = m.M
																						v132 = m.ExcPending
																						if v132 != 0 {
																							return
																						} else {
																							v134 = *(*int32)(unsafe.Add(mBase, _consts[125]))
																							v136 = *(*int32)(unsafe.Add(mBase, _consts[116]))
																							v138 = base.I32_div_s(v136, int32(1048576))
																							v139 = base.I32_div_s(v134, v138)
																							if v139 <= int32(1) {
																								F_errstart_cold(m, int32(21), int32(0))
																								mBase = m.M
																								v687 = m.ExcPending
																								if v687 != 0 {
																									return
																								} else {
																									F_errcode(m, int32(50856066))
																									mBase = m.M
																									v690 = m.ExcPending
																									if v690 != 0 {
																										return
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(323890)
																										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(324203)
																										F_errmsg(m, int32(680203), v7+int32(16))
																										mBase = m.M
																										v699 = m.ExcPending
																										if v699 != 0 {
																											return
																										} else {
																											F_errfinish(m, int32(475016), int32(4560), int32(370959))
																											mBase = m.M
																											v704 = m.ExcPending
																											if v704 != 0 {
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
																								v143 = *(*int32)(unsafe.Add(mBase, _consts[126]))
																								v144 = base.I32_div_s(v143, v138)
																								if v144 <= int32(1) {
																									F_errstart_cold(m, int32(21), int32(0))
																									mBase = m.M
																									v708 = m.ExcPending
																									if v708 != 0 {
																										return
																									} else {
																										F_errcode(m, int32(50856066))
																										mBase = m.M
																										v711 = m.ExcPending
																										if v711 != 0 {
																											return
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = int32(323890)
																											*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(324190)
																											F_errmsg(m, int32(680203), v7+int32(32))
																											mBase = m.M
																											v720 = m.ExcPending
																											if v720 != 0 {
																												return
																											} else {
																												F_errfinish(m, int32(475016), int32(4566), int32(370959))
																												mBase = m.M
																												v725 = m.ExcPending
																												if v725 != 0 {
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
																									v149 = base.I32_div_s(v136, int32(8192))
																									*(*int32)(unsafe.Add(mBase, _consts[127])) = v149*int32(8168) - int32(16)
																									v159 = *(*float64)(unsafe.Add(mBase, _consts[128]))
																									v162 = base.F64_div(base.F64_convert_i32_u(v144), base.F64_add(v159, float64(1)))
																									if base.F64_lt(base.F64_abs(v162), float64(2.147483648e+09)) != 0 {
																										v166 = base.I32_trunc_f64_s(v162)
																										v168 = v166
																									} else {
																										v168 = int32(-2147483648)
																									}
																									if v168 <= int32(1) {
																										v171 = int32(1)
																									} else {
																										v171 = v168
																									}
																									*(*int32)(unsafe.Add(mBase, _consts[114])) = v171
																									v177 = *(*int32)(unsafe.Add(mBase, _consts[120]))
																									v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)+252))
																									if v178 != 0 {
																										v179 = int32(147589)
																									} else {
																										v179 = int32(228702)
																									}
																									F_SetConfigOption(m, int32(141051), v179, int32(0), int32(1))
																									mBase = m.M
																									v183 = m.ExcPending
																									if v183 != 0 {
																										return
																									} else {
																										m.G0 = v7 + int32(320)
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
				}
			}
		} else {
			F_errstart_cold(m, int32(23), int32(0))
			mBase = m.M
			v190 = m.ExcPending
			if v190 != 0 {
				return
			} else {
				F_errcode_for_file_access(m)
				mBase = m.M
				v192 = m.ExcPending
				if v192 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(286739)
					F_errmsg(m, int32(284016), v7)
					mBase = m.M
					v197 = m.ExcPending
					if v197 != 0 {
						return
					} else {
						F_errfinish(m, int32(475016), int32(4362), int32(370959))
						mBase = m.M
						v202 = m.ExcPending
						if v202 != 0 {
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
		*(*int32)(unsafe.Add(mBase, _consts[86])) = v18
		v24 = F_strtox_2(m, v6, l0, int32(10), int64(2147483648))
		mBase = m.M
		v27 = *(*int32)(unsafe.Add(mBase, _consts[86]))
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
						F_errmsg(m, int32(382860), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, l2, int32(471701), int32(538), int32(91045))
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
			*(*int32)(unsafe.Add(mBase, _consts[86])) = v18
			v24 = F_strtox_2(m, v6, l0, int32(10), int64(2147483648))
			mBase = m.M
			v27 = *(*int32)(unsafe.Add(mBase, _consts[86]))
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
							F_errmsg(m, int32(382860), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, l2, int32(471701), int32(538), int32(91045))
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int64
	_ = v94
	var v102 int32
	_ = v102
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
	var v132 int32
	_ = v132
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
	var v267 int32
	_ = v267
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
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
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
	var v431 int32
	_ = v431
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
	var v532 int32
	_ = v532
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
	var v742 int32
	_ = v742
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
	var v828 int32
	_ = v828
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v856 int64
	_ = v856
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v869 int64
	_ = v869
	var v873 int64
	_ = v873
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v908 int64
	_ = v908
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
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
	*(*uint8)(unsafe.Add(mBase, _consts[157])) = uint8(v5)
	v36 = base.B2i32(l1 != int32(15))
	goto L1
L1:
	;
	v58 = v21 + int32(140)
	v60 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v60 != v61 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	m.G0 = v21 + int32(144)
	return v930
L3:
	;
	goto L2
L4:
	;
	v823 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[157])) = uint8(v823)
	v826 = int32(*(*uint8)(unsafe.Add(mBase, _consts[158])))
	if v826 != 0 {
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
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v63 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
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
	F_pfree(m, v63)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _consts[159]))
	v70 = int32(2)
	v76 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	v77 = int32(0)
	v81 = base.B2i32(v76 != v77) & base.B2i32(v77 < v69)
	if v81 != 0 {
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
	v82 = v69<<(uint(v70)%32) | int32(1)
	goto L16
L15:
	;
	v82 = v70
	goto L16
L16:
	;
	v87 = F_palloc(m, v82<<(uint(int32(4))%32)+int32(32))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = v82
	if v81 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v91 = v69
	goto L20
L19:
	;
	v91 = int32(1)
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = l0
	v94 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v87)+12)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = int32(410)
	*(*int64)(unsafe.Add(mBase, uint32(v87)+20)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v87
	v102 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v102
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
	*(*int32)(unsafe.Add(mBase, uint32(v109)+112)) = v132
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
	v132 = v125
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
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+4)))
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
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	if v147 != 0 {
		v132 = v147
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
	v321 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	if v321 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L48:
	;
	v267 = v257
	goto L49
L49:
	;
	v280 = v256 + v267<<(uint(int32(4))%32)
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
	v295 = v267 + int32(1)
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
		v267 = v299
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
	v325 = *(*int32)(unsafe.Add(mBase, _consts[159]))
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
	v339 = v328
	v342 = v258
	goto L64
L64:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v254)+16))
	if base.Ui32(v337-int32(1)) <= base.Ui32(v350+v339) {
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
	v362 = m.T0[v361].(func(*base.Module, int32, int32) int32)(m, v355, v256+v342<<(uint(int32(4))%32)+int32(8))
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
		v339 = v393
		v342 = v391
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
	v431 = v418
	goto L80
L80:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v417)+16))
	if base.Ui32(v429-int32(1)) <= base.Ui32(v442+v431) {
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
		v431 = v486
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
	*(*int32)(unsafe.Add(mBase, uint32(v497)+112)) = v532
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
	v532 = v525
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
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532)+4)))
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
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v532)+8))
	if v547 != 0 {
		v532 = v547
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
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v593
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
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(0)
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
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v599
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
	v640 = *(*int32)(unsafe.Add(mBase, _consts[160]))
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
	v657 = int32(*(*uint8)(unsafe.Add(mBase, _consts[95])))
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
	v709 = *(*int32)(unsafe.Add(mBase, _consts[161]))
	v710 = int32(0)
	if v709 == v710 {
		goto L149
	} else {
		goto L150
	}
L133:
	;
	v668 = *(*int32)(unsafe.Add(mBase, _consts[162]))
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
	*(*int64)(unsafe.Add(mBase, _consts[163])) = v658
	v664 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int64)(unsafe.Add(mBase, _consts[164])) = v664
	goto L133
L136:
	;
	v671 = F_close(m, v668)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[162])) = int32(-1)
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
	v679 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	if v679 != int32(2) {
		v690 = l1
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v683 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	v685 = *(*int64)(unsafe.Add(mBase, _consts[166]))
	if v683 == v685 {
		v690 = int32(14)
		goto L140
	} else {
		goto L143
	}
L143:
	;
	*(*int64)(unsafe.Add(mBase, _consts[166])) = v683
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
	F_errmsg_internal(m, int32(195849), v21)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L12
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(469636), int32(3209), int32(401689))
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
		v930 = v653
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
		v742 = v710
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v749 = v742
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
		v742 = v735
		goto L152
	} else {
		goto L159
	}
L158:
	;
	v742 = v735
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
	v754 = *(*int32)(unsafe.Add(mBase, _consts[116]))
	v755 = base.I64_extend_i32_s(v754)
	v756 = base.I64_div_u_s(v750, v755)
	v758 = base.I64_div_u_s(int64(4294967296), v755)
	v759 = base.I64_div_u_s(v756, v758)
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+52)) = uint32(v759)
	v762 = v756 - v758*v759
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+56)) = uint32(v762)
	v770 = F_pg_snprintf(m, v21-int32(-64), int32(64), int32(486532), v21+int32(48))
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
	v773 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	if v773 != int32(2) {
		v784 = l1
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v777 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	v779 = *(*int64)(unsafe.Add(mBase, _consts[166]))
	if v777 == v779 {
		v784 = int32(14)
		goto L163
	} else {
		goto L166
	}
L166:
	;
	*(*int64)(unsafe.Add(mBase, _consts[166])) = v777
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
	F_errmsg(m, int32(38936), v21+int32(16))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L12
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(469636), int32(3231), int32(401689))
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
	v921 = int32(0)
	v923 = int32(*(*uint8)(unsafe.Add(mBase, _consts[167])))
	if v923 != int32(1) {
		v930 = v921
		goto L3
	} else {
		goto L196
	}
L172:
	;
	if v3 != 0 {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v828 = int32(*(*uint8)(unsafe.Add(mBase, _consts[95])))
	if v828&int32(1) == int32(0) {
		goto L171
	} else {
		goto L174
	}
L174:
	;
	v835 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L12
	} else {
		goto L175
	}
L175:
	;
	if v835 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	F_errmsg_internal(m, int32(13842), int32(0))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L12
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v847 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[158])) = uint8(v847)
	v850 = int32(*(*uint8)(unsafe.Add(mBase, _consts[168])))
	if v850 != 0 {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	F_errfinish(m, int32(469636), int32(3261), int32(401689))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L12
	} else {
		goto L180
	}
L180:
	;
	goto L178
L181:
	;
	v852 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[167])) = uint8(v852)
	F_disable_startup_progress_timeout(m)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L12
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v856 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	v858 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v862 = F_LWLockAcquire(m, v858+int32(1152), int32(0))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L12
	} else {
		goto L185
	}
L184:
	;
	goto L183
L185:
	;
	v865 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	*(*int32)(unsafe.Add(mBase, uint32(v865)+16)) = int32(5)
	v869 = *(*int64)(unsafe.Add(mBase, uint32(v865)+136))
	if base.Ui64(v869) < base.Ui64(v856) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v865)+144)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v865)+136)) = v856
	v873 = v856
	goto L188
L187:
	;
	v873 = v869
	goto L188
L188:
	;
	*(*int64)(unsafe.Add(mBase, _consts[135])) = v873
	v876 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[169])) = uint8(v876)
	v879 = *(*int32)(unsafe.Add(mBase, _consts[132]))
	F_update_controlfile(m, v879, v865)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L12
	} else {
		goto L189
	}
L189:
	;
	v883 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v883)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v883)+440)) = int32(1)
	if v884 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v888 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	F_s_lock(m, v888+int32(440), int32(475016), int32(6275), int32(13995))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L12
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v897 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	*(*int32)(unsafe.Add(mBase, uint32(v897)+440)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v897)+316)) = int32(1)
	v903 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v903+int32(1152))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L12
	} else {
		goto L194
	}
L193:
	;
	goto L192
L194:
	;
	v908 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int32)(unsafe.Add(mBase, _consts[170])) = l3
	*(*int64)(unsafe.Add(mBase, _consts[171])) = v908
	F_CheckRecoveryConsistency(m)
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L12
	} else {
		goto L195
	}
L195:
	;
	v916 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[172])) = v916
	*(*uint8)(unsafe.Add(mBase, _consts[157])) = uint8(v916)
	goto L1
L196:
	;
	v926 = F_CheckForStandbyTrigger(m)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L12
	} else {
		goto L197
	}
L197:
	;
	if v926 == int32(0) {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v930 = v921
	goto L3
}
func F_RegisterSnapshotOnOwner(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
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
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	if l0 == int32(0) {
		return int32(0)
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
		if v12 != 0 {
			v95 = l0
			F_ResourceOwnerEnlarge(m, l1)
			mBase = m.M
			v101 = m.ExcPending
			if v101 != 0 {
				return int32(0)
			} else {
				v102 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v95)+48)) = v102 + int32(1)
				F_ResourceOwnerRemember(m, l1, v95, int32(1728324))
				mBase = m.M
				v108 = m.ExcPending
				if v108 != 0 {
					return int32(0)
				} else {
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
					if v109 == int32(1) {
						F_pairingheap_add(m, int32(4114840), v95+int32(52))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							return v95
						}
					} else {
						return v95
					}
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _consts[319]))
			v16 = l0 + int32(24)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v19 = l0 + int32(16)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			v22 = int32(2)
			v24 = int32(72)
			v29 = v20<<(uint(v22)%32) + v24
			if int32(0) < v17 {
				v32 = (v17+v20)<<(uint(v22)%32) + v24
			} else {
				v32 = v29
			}
			v33 = F_MemoryContextAlloc(m, v14, v32)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v38 = v33 + int32(48)
				v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int64)(unsafe.Add(mBase, uint32(v38))) = v39
				v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
				*(*int64)(unsafe.Add(mBase, uint32(v33)+40)) = v41
				v43 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
				*(*int64)(unsafe.Add(mBase, uint32(v33)+24)) = v43
				v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
				*(*int64)(unsafe.Add(mBase, uint32(v33)+56)) = v45
				v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v33)+32)) = v47
				v49 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
				*(*int64)(unsafe.Add(mBase, uint32(v33)+16)) = v49
				v51 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = v51
				v53 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(v33))) = v53
				*(*int64)(unsafe.Add(mBase, uint32(v33)+64)) = int64(0)
				v57 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v38))) = v57
				*(*int32)(unsafe.Add(mBase, uint32(v33)+44)) = v57
				v61 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v33)+30)) = uint8(v61)
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				if v63 != 0 {
					v65 = v33 + int32(72)
					*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v65
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v70 = v68 << (uint(int32(2)) % 32)
					if v70 != 0 {
						v71 = F__emscripten_memcpy_bulkmem(m, v65, v67, v70)
						mBase = m.M
					} else {
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = int32(0)
				}
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v76 <= int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = int32(0)
					v95 = v33
				} else {
					v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
					if v79 == int32(1) {
						v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
						if v82 != int32(1) {
							*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = int32(0)
							v95 = v33
						} else {
							v85 = v33 + v29
							*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v85
							v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v90 = v88 << (uint(int32(2)) % 32)
							if v90 != 0 {
								v91 = F__emscripten_memcpy_bulkmem(m, v85, v87, v90)
								mBase = m.M
							} else {
							}
							v95 = v33
						}
					} else {
						v85 = v33 + v29
						*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v85
						v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v90 = v88 << (uint(int32(2)) % 32)
						if v90 != 0 {
							v91 = F__emscripten_memcpy_bulkmem(m, v85, v87, v90)
							mBase = m.M
						} else {
						}
						v95 = v33
					}
				}
				F_ResourceOwnerEnlarge(m, l1)
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return int32(0)
				} else {
					v102 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v95)+48)) = v102 + int32(1)
					F_ResourceOwnerRemember(m, l1, v95, int32(1728324))
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int32(0)
					} else {
						v109 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
						if v109 == int32(1) {
							F_pairingheap_add(m, int32(4114840), v95+int32(52))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int32(0)
							} else {
								return v95
							}
						} else {
							return v95
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
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
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
	var v98 int32
	_ = v98
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
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
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
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
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
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
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(128)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[1134]))
	if v13 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[367]))
	if v17 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v76 = v13
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+124)) = l1
	v80 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	if l0 != v80 {
		goto L14
	} else {
		goto L15
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
	v31 = F__emscripten_memset_bulkmem(m, int32(4431168), base.I32_extend8_s(int32(0)), int32(96))
	mBase = m.M
	goto L9
L7:
	;
	return int32(0)
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[367]))
	v26 = v25
	goto L6
L9:
	;
	F_fmgr_info_cxt(m, int32(184), int32(4431184), v26)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	*(*int64)(unsafe.Add(mBase, _consts[1135])) = int64(0)
	v40 = int32(3)
	*(*uint16)(unsafe.Add(mBase, _consts[1136])) = uint16(v40)
	v45 = *(*int32)(unsafe.Add(mBase, _consts[367]))
	F_fmgr_info_cxt(m, int32(184), int32(4431232), v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	*(*int64)(unsafe.Add(mBase, _consts[1137])) = int64(196616)
	v52 = int32(9)
	*(*uint16)(unsafe.Add(mBase, _consts[1138])) = uint16(v52)
	*(*int32)(unsafe.Add(mBase, _consts[1139])) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = int64(51539607560)
	v60 = *(*int32)(unsafe.Add(mBase, _consts[367]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v60
	v68 = F_hash_create(m, int32(380375), int32(64), v10+int32(16), int32(1064))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1134])) = v68
	F_CacheRegisterRelcacheCallback(m, int32(1609))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[1134]))
	v76 = v75
	goto L3
L14:
	;
	v82 = l0
	goto L16
L15:
	;
	v82 = int32(0)
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+120)) = v82
	v89 = F_hash_search(m, v76, v10+int32(120), int32(0), v10+int32(119))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+119)))
	if v91 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L7
	} else {
		goto L130
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L7
	} else {
		goto L127
	}
L20:
	;
	m.G0 = v10 + int32(128)
	return v404
L21:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v404 = v94
	goto L20
L22:
	;
	goto L23
L23:
	;
	if v82 == int32(1664) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _consts[1134]))
	v394 = F_hash_search(m, v388, v10+int32(120), int32(1), v10+int32(119))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L7
	} else {
		goto L125
	}
L25:
	;
	v98 = int32(0)
	goto L32
L26:
	;
	goto L27
L27:
	;
	v211 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L7
	} else {
		goto L66
	}
L28:
	;
	v384 = v208
	goto L24
L29:
	;
	goto L28
L30:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	v208 = v203
	goto L29
L32:
	;
	goto L33
L33:
	;
	v150 = int32(0)
	v152 = *(*int32)(unsafe.Add(mBase, _consts[1047]))
	if v150 < v152 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v182 = int32(0)
	goto L60
L51:
	;
	v202 = v161 + int32(4431272)
	goto L30
L52:
	;
	v156 = v150
	goto L55
L53:
	;
	goto L54
L54:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _consts[1050]))
	if v175 <= int32(0) {
		v208 = v98
		goto L29
	} else {
		goto L59
	}
L55:
	;
	v161 = v156 << (uint(int32(3)) % 32)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v161)+uint32(_consts[1140])))
	if l1 == v164 {
		goto L51
	} else {
		goto L57
	}
L56:
	;
	goto L54
L57:
	;
	v167 = v156 + int32(1)
	if v167 != v152 {
		v156 = v167
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	goto L50
L60:
	;
	v187 = v182 << (uint(int32(3)) % 32)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_consts[1141])))
	if v190 != l1 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v202 = v187 + int32(4431796)
	goto L30
L62:
	;
	v193 = v182 + int32(1)
	if v175 != v193 {
		v182 = v193
		goto L60
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	goto L61
L65:
	;
	v208 = v98
	goto L29
L66:
	;
	goto L68
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+108)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v82
	v227 = F_systable_beginscan(m, v211, int32(3455), int32(1), int32(0), int32(2), v10+int32(16))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L7
	} else {
		goto L71
	}
L68:
	;
	v217 = F__emscripten_memcpy_bulkmem(m, v10+int32(16), int32(4431168), int32(96))
	mBase = m.M
	goto L70
L70:
	;
	goto L67
L71:
	;
	v229 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+119)) = uint8(v229)
	v231 = F_systable_getnext(m, v227)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	if v231 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v233 = v231
	v237 = v3
	goto L76
L74:
	;
	v259 = v3
	goto L75
L75:
	;
	F_systable_endscan(m, v227)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L7
	} else {
		goto L84
	}
L76:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v233)+16))
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+22)))
	v242 = v240 + v241
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+118)))
	if v243 != int32(116) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v259 = v252
	goto L75
L78:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+119)))
	if v246 == int32(1) {
		goto L19
	} else {
		goto L81
	}
L79:
	;
	v252 = v237
	goto L80
L80:
	;
	v253 = F_systable_getnext(m, v227)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L7
	} else {
		goto L82
	}
L81:
	;
	v249 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+119)) = uint8(v249)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	v252 = v251
	goto L80
L82:
	;
	if v253 != 0 {
		v233 = v253
		v237 = v252
		goto L76
	} else {
		goto L83
	}
L83:
	;
	goto L77
L84:
	;
	F_sequence_close(m, v211, int32(1))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L7
	} else {
		goto L85
	}
L85:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+119)))
	if v267 != 0 {
		v384 = v259
		goto L24
	} else {
		goto L86
	}
L86:
	;
	v268 = int32(0)
	goto L90
L87:
	;
	v384 = v379
	goto L24
L88:
	;
	goto L87
L89:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	v379 = v374
	goto L88
L90:
	;
	v274 = int32(0)
	v276 = *(*int32)(unsafe.Add(mBase, _consts[1046]))
	if v274 < v276 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v280 = v274
	goto L96
L94:
	;
	goto L95
L95:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _consts[1048]))
	if v301 <= int32(0) {
		v379 = v268
		goto L88
	} else {
		goto L102
	}
L96:
	;
	v285 = v280 << (uint(int32(3)) % 32)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v285)+uint32(_consts[1142])))
	if v288 == l1 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L95
L98:
	;
	v373 = v285 + int32(4432320)
	goto L89
L99:
	;
	goto L100
L100:
	;
	v293 = v280 + int32(1)
	if v293 != v276 {
		v280 = v293
		goto L96
	} else {
		goto L101
	}
L101:
	;
	goto L97
L102:
	;
	v306 = int32(0)
	goto L103
L103:
	;
	v311 = v306 << (uint(int32(3)) % 32)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v311)+uint32(_consts[1143])))
	if v314 != l1 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v373 = v311 + int32(4432844)
	goto L89
L105:
	;
	v317 = v306 + int32(1)
	if v301 != v317 {
		v306 = v317
		goto L103
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	goto L104
L108:
	;
	v379 = v268
	goto L88
L125:
	;
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+119)))
	if v396 == int32(1) {
		goto L18
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v394)+8)) = v384
	v404 = v384
	goto L20
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v82
	F_errmsg_internal(m, int32(41474), v10)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L7
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(473057), int32(222), int32(217032))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L7
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	F_errmsg_internal(m, int32(373035), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L7
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(473057), int32(245), int32(217032))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L7
	} else {
		goto L132
	}
L132:
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
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
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
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
	v130 = int32(4442992)
	v131 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v134 = *(*int32)(unsafe.Add(mBase, _consts[806]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v134
	v137 = *(*int32)(unsafe.Add(mBase, _consts[807]))
	v141 = F_hash_search(m, v137, l0, int32(1), v9+int32(12))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L6
	} else {
		goto L33
	}
L3:
	;
	v102 = int32(4442992)
	v103 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v106 = *(*int32)(unsafe.Add(mBase, _consts[806]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v106
	v109 = F_palloc(m, int32(32))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L6
	} else {
		goto L31
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[807]))
	F_hash_seq_init(m, v9+int32(12), v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L9
	}
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[807]))
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
	v31 = F_hash_seq_search(m, v9+int32(12))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	if v31 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v34 = v31
	goto L14
L12:
	;
	goto L13
L13:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[808]))
	if v65 == int32(0) {
		goto L1
	} else {
		goto L22
	}
L14:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34))))
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v39 != v40 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	v56 = F_hash_seq_search(m, v9+int32(12))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L6
	} else {
		goto L20
	}
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_extend16_s(v39)*int32(12))+uint32(_consts[809])))
	v48 = m.T0[v47].(func(*base.Module, int32, int32) int32)(m, l0, v34)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	if v48 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v52 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+26)) = uint8(v52)
	goto L16
L20:
	;
	if v56 != 0 {
		v34 = v56
		goto L14
	} else {
		goto L21
	}
L21:
	;
	goto L15
L22:
	;
	v68 = int32(0)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v69 <= v68 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v73 = v68
	goto L24
L24:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+v73<<(uint(int32(2))%32))))
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82))))
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v83 != v84 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L1
L26:
	;
	v99 = v73 + int32(1)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v99 < v100 {
		v73 = v99
		goto L24
	} else {
		goto L30
	}
L27:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_extend16_s(v83)*int32(12))+uint32(_consts[809])))
	v92 = m.T0[v91].(func(*base.Module, int32, int32) int32)(m, l0, v82)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	if v92 == int32(0) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v96 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v82)+26)) = uint8(v96)
	goto L26
L30:
	;
	goto L25
L31:
	;
	v111 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+16)) = v111
	v113 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+8)) = v113
	v115 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v109))) = v115
	v118 = int32(*(*uint16)(unsafe.Add(mBase, _consts[810])))
	v119 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+26)) = uint8(v119)
	*(*uint16)(unsafe.Add(mBase, uint32(v109)+24)) = uint16(v118)
	v123 = *(*int32)(unsafe.Add(mBase, _consts[808]))
	v124 = F_lappend(m, v123, v109)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v103
	*(*int32)(unsafe.Add(mBase, _consts[808])) = v124
	goto L1
L33:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+12)))
	if v143 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v131
	goto L1
L35:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+26)))
	if v146 != int32(1) {
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v150 = int32(*(*uint16)(unsafe.Add(mBase, _consts[811])))
	v151 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+26)) = uint8(v151)
	*(*uint16)(unsafe.Add(mBase, uint32(v141)+24)) = uint16(v150)
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
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
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	v5 = m.G0
	v7 = v5 - int32(96)
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
	F_ScanKeyInit(m, v7, int32(1), int32(3), int32(184), l0)
	v17 = m.ExcPending
	if v17 != 0 {
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
	v19 = int32(2)
	F_ScanKeyInit(m, v7+int32(48), v19, int32(3), int32(63), l1)
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v27 = int32(1)
	goto L6
L6:
	;
	v31 = F_systable_beginscan(m, v11, int32(2696), int32(1), int32(0), v27, v7)
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v27 = v19
	goto L6
L8:
	;
	v33 = F_systable_getnext(m, v31)
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v33 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v36 = v33
	goto L13
L11:
	;
	goto L12
L12:
	;
	F_systable_endscan(m, v31)
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	F_CatalogTupleDelete(m, v11, v36+int32(4))
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v43 = F_systable_getnext(m, v31)
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v43 != 0 {
		v36 = v43
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	F_sequence_close(m, v11, int32(3))
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	m.G0 = v7 + int32(96)
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v14 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v18 = F_SearchSysCacheCopy(m, int32(19), l0, int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			if v18 != 0 {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+22)))
				v22 = v20 + v21
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+80))
				if v23 != 0 {
					v25 = F_ConstraintNameIsUsed(m, int32(0), v23, l1)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						if v25 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return
							} else {
								F_errcode(m, int32(290948))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v22)+80))
									v78 = F_get_rel_name(m, v77)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v78
										*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l1
										F_errmsg(m, int32(108557), v10+int32(32))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											F_errfinish(m, int32(470264), int32(1026), int32(443415))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
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
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
							if v27 != 0 {
								v29 = F_ConstraintNameIsUsed(m, int32(1), v27, l1)
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return
								} else {
									if v29 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											F_errcode(m, int32(290948))
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return
											} else {
												v99 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
												v100 = F_format_type_be(m, v99)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v100
													*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
													F_errmsg(m, int32(108073), v10+int32(16))
													mBase = m.M
													v108 = m.ExcPending
													if v108 != 0 {
														return
													} else {
														F_errfinish(m, int32(470264), int32(1034), int32(443415))
														mBase = m.M
														v113 = m.ExcPending
														if v113 != 0 {
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
										v34 = F_strncpy(m, v22+int32(4), l1, int32(64))
										mBase = m.M
										v35 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v34)+63)) = uint8(v35)
										F_CatalogTupleUpdate(m, v14, v18+int32(4), v18)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return
										} else {
											v42 = *(*int32)(unsafe.Add(mBase, _consts[230]))
											if v42 != 0 {
												v44 = int32(0)
												F_RunObjectPostAlterHook(m, int32(2606), l0, v44, v44, v44)
												mBase = m.M
												v48 = m.ExcPending
												if v48 != 0 {
													return
												} else {
													F_pfree(m, v18)
													mBase = m.M
													v50 = m.ExcPending
													if v50 != 0 {
														return
													} else {
														F_sequence_close(m, v14, int32(3))
														mBase = m.M
														v53 = m.ExcPending
														if v53 != 0 {
															return
														} else {
															m.G0 = v10 + int32(48)
															return
														}
													}
												}
											} else {
												F_pfree(m, v18)
												mBase = m.M
												v50 = m.ExcPending
												if v50 != 0 {
													return
												} else {
													F_sequence_close(m, v14, int32(3))
													mBase = m.M
													v53 = m.ExcPending
													if v53 != 0 {
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
								v34 = F_strncpy(m, v22+int32(4), l1, int32(64))
								mBase = m.M
								v35 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v34)+63)) = uint8(v35)
								F_CatalogTupleUpdate(m, v14, v18+int32(4), v18)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									v42 = *(*int32)(unsafe.Add(mBase, _consts[230]))
									if v42 != 0 {
										v44 = int32(0)
										F_RunObjectPostAlterHook(m, int32(2606), l0, v44, v44, v44)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											F_pfree(m, v18)
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return
											} else {
												F_sequence_close(m, v14, int32(3))
												mBase = m.M
												v53 = m.ExcPending
												if v53 != 0 {
													return
												} else {
													m.G0 = v10 + int32(48)
													return
												}
											}
										}
									} else {
										F_pfree(m, v18)
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return
										} else {
											F_sequence_close(m, v14, int32(3))
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
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
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
					if v27 != 0 {
						v29 = F_ConstraintNameIsUsed(m, int32(1), v27, l1)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							if v29 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return
								} else {
									F_errcode(m, int32(290948))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return
									} else {
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
										v100 = F_format_type_be(m, v99)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v100
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
											F_errmsg(m, int32(108073), v10+int32(16))
											mBase = m.M
											v108 = m.ExcPending
											if v108 != 0 {
												return
											} else {
												F_errfinish(m, int32(470264), int32(1034), int32(443415))
												mBase = m.M
												v113 = m.ExcPending
												if v113 != 0 {
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
								v34 = F_strncpy(m, v22+int32(4), l1, int32(64))
								mBase = m.M
								v35 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v34)+63)) = uint8(v35)
								F_CatalogTupleUpdate(m, v14, v18+int32(4), v18)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									v42 = *(*int32)(unsafe.Add(mBase, _consts[230]))
									if v42 != 0 {
										v44 = int32(0)
										F_RunObjectPostAlterHook(m, int32(2606), l0, v44, v44, v44)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											F_pfree(m, v18)
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return
											} else {
												F_sequence_close(m, v14, int32(3))
												mBase = m.M
												v53 = m.ExcPending
												if v53 != 0 {
													return
												} else {
													m.G0 = v10 + int32(48)
													return
												}
											}
										}
									} else {
										F_pfree(m, v18)
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return
										} else {
											F_sequence_close(m, v14, int32(3))
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
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
						v34 = F_strncpy(m, v22+int32(4), l1, int32(64))
						mBase = m.M
						v35 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v34)+63)) = uint8(v35)
						F_CatalogTupleUpdate(m, v14, v18+int32(4), v18)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, _consts[230]))
							if v42 != 0 {
								v44 = int32(0)
								F_RunObjectPostAlterHook(m, int32(2606), l0, v44, v44, v44)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									F_pfree(m, v18)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										F_sequence_close(m, v14, int32(3))
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return
										} else {
											m.G0 = v10 + int32(48)
											return
										}
									}
								}
							} else {
								F_pfree(m, v18)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									F_sequence_close(m, v14, int32(3))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
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
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
					F_errmsg_internal(m, int32(38585), v10)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						F_errfinish(m, int32(470264), int32(1013), int32(443415))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
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
			v27 = F_SearchSysCacheLockedCopy1(m, int32(57), l0)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				if v27 != 0 {
					v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+8)))
					*(*uint16)(unsafe.Add(mBase, uint32(v13)+28)) = uint16(v29)
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v31
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
					v35 = F_get_relname_relid(m, l1, v21)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						if v35 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return
							} else {
								F_errcode(m, int32(117571716))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
									F_errmsg(m, int32(108710), v13+int32(16))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return
									} else {
										F_errfinish(m, int32(471646), int32(4307), int32(297300))
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
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
							v41 = F_strncpy(m, v33+v34+int32(4), l1, int32(64))
							mBase = m.M
							v42 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v41)+63)) = uint8(v42)
							F_CatalogTupleUpdate(m, v24, v13+int32(24), v27)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								F_UnlockTuple(m, v24, v13+int32(24), int32(7))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, _consts[230]))
									if v54 != 0 {
										v56 = int32(0)
										F_RunObjectPostAlterHook(m, int32(1259), l0, v56, v56, l2)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											F_pfree(m, v27)
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return
											} else {
												F_sequence_close(m, v24, int32(3))
												mBase = m.M
												v64 = m.ExcPending
												if v64 != 0 {
													return
												} else {
													v65 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
													v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+72))
													if v66 != 0 {
														F_RenameTypeInternal(m, v66, l1, v21)
														mBase = m.M
														v68 = m.ExcPending
														if v68 != 0 {
															return
														} else {
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
															v70 = v69
															v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+119)))
															if v71|int32(32) != int32(105) {
																F_relation_close(m, v18, int32(0))
																mBase = m.M
																v85 = m.ExcPending
																if v85 != 0 {
																	return
																} else {
																	m.G0 = v13 + int32(32)
																	return
																}
															} else {
																v76 = F_get_index_constraint(m, l0)
																mBase = m.M
																v77 = m.ExcPending
																if v77 != 0 {
																	return
																} else {
																	if v76 == int32(0) {
																		F_relation_close(m, v18, int32(0))
																		mBase = m.M
																		v85 = m.ExcPending
																		if v85 != 0 {
																			return
																		} else {
																			m.G0 = v13 + int32(32)
																			return
																		}
																	} else {
																		F_RenameConstraintById(m, v76, l1)
																		mBase = m.M
																		v81 = m.ExcPending
																		if v81 != 0 {
																			return
																		} else {
																			F_relation_close(m, v18, int32(0))
																			mBase = m.M
																			v85 = m.ExcPending
																			if v85 != 0 {
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
														v70 = v65
														v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+119)))
														if v71|int32(32) != int32(105) {
															F_relation_close(m, v18, int32(0))
															mBase = m.M
															v85 = m.ExcPending
															if v85 != 0 {
																return
															} else {
																m.G0 = v13 + int32(32)
																return
															}
														} else {
															v76 = F_get_index_constraint(m, l0)
															mBase = m.M
															v77 = m.ExcPending
															if v77 != 0 {
																return
															} else {
																if v76 == int32(0) {
																	F_relation_close(m, v18, int32(0))
																	mBase = m.M
																	v85 = m.ExcPending
																	if v85 != 0 {
																		return
																	} else {
																		m.G0 = v13 + int32(32)
																		return
																	}
																} else {
																	F_RenameConstraintById(m, v76, l1)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return
																	} else {
																		F_relation_close(m, v18, int32(0))
																		mBase = m.M
																		v85 = m.ExcPending
																		if v85 != 0 {
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
										F_pfree(m, v27)
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return
										} else {
											F_sequence_close(m, v24, int32(3))
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return
											} else {
												v65 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
												v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+72))
												if v66 != 0 {
													F_RenameTypeInternal(m, v66, l1, v21)
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return
													} else {
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
														v70 = v69
														v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+119)))
														if v71|int32(32) != int32(105) {
															F_relation_close(m, v18, int32(0))
															mBase = m.M
															v85 = m.ExcPending
															if v85 != 0 {
																return
															} else {
																m.G0 = v13 + int32(32)
																return
															}
														} else {
															v76 = F_get_index_constraint(m, l0)
															mBase = m.M
															v77 = m.ExcPending
															if v77 != 0 {
																return
															} else {
																if v76 == int32(0) {
																	F_relation_close(m, v18, int32(0))
																	mBase = m.M
																	v85 = m.ExcPending
																	if v85 != 0 {
																		return
																	} else {
																		m.G0 = v13 + int32(32)
																		return
																	}
																} else {
																	F_RenameConstraintById(m, v76, l1)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return
																	} else {
																		F_relation_close(m, v18, int32(0))
																		mBase = m.M
																		v85 = m.ExcPending
																		if v85 != 0 {
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
													v70 = v65
													v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+119)))
													if v71|int32(32) != int32(105) {
														F_relation_close(m, v18, int32(0))
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return
														} else {
															m.G0 = v13 + int32(32)
															return
														}
													} else {
														v76 = F_get_index_constraint(m, l0)
														mBase = m.M
														v77 = m.ExcPending
														if v77 != 0 {
															return
														} else {
															if v76 == int32(0) {
																F_relation_close(m, v18, int32(0))
																mBase = m.M
																v85 = m.ExcPending
																if v85 != 0 {
																	return
																} else {
																	m.G0 = v13 + int32(32)
																	return
																}
															} else {
																F_RenameConstraintById(m, v76, l1)
																mBase = m.M
																v81 = m.ExcPending
																if v81 != 0 {
																	return
																} else {
																	F_relation_close(m, v18, int32(0))
																	mBase = m.M
																	v85 = m.ExcPending
																	if v85 != 0 {
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
					v92 = m.ExcPending
					if v92 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
						F_errmsg_internal(m, int32(43821), v13)
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return
						} else {
							F_errfinish(m, int32(471646), int32(4299), int32(297300))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
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
func F_ReservePrivateRefCountEntry(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[623]))
	if v8 != 0 {
		m.G0 = v5 + int32(16)
		return
	} else {
		v10 = int32(4358832)
		v12 = *(*int32)(unsafe.Add(mBase, _consts[614]))
		if v12 == int32(0) {
			v48 = v10
			*(*int32)(unsafe.Add(mBase, _consts[623])) = v48
			m.G0 = v5 + int32(16)
			return
		} else {
			v15 = int32(4358840)
			v17 = *(*int32)(unsafe.Add(mBase, _consts[615]))
			if v17 == int32(0) {
				v48 = v15
				*(*int32)(unsafe.Add(mBase, _consts[623])) = v48
				m.G0 = v5 + int32(16)
				return
			} else {
				v20 = int32(4358848)
				v22 = *(*int32)(unsafe.Add(mBase, _consts[616]))
				if v22 == int32(0) {
					v48 = v20
					*(*int32)(unsafe.Add(mBase, _consts[623])) = v48
					m.G0 = v5 + int32(16)
					return
				} else {
					v25 = int32(4358856)
					v27 = *(*int32)(unsafe.Add(mBase, _consts[617]))
					if v27 == int32(0) {
						v48 = v25
						*(*int32)(unsafe.Add(mBase, _consts[623])) = v48
						m.G0 = v5 + int32(16)
						return
					} else {
						v30 = int32(4358864)
						v32 = *(*int32)(unsafe.Add(mBase, _consts[618]))
						if v32 == int32(0) {
							v48 = v30
							*(*int32)(unsafe.Add(mBase, _consts[623])) = v48
							m.G0 = v5 + int32(16)
							return
						} else {
							v35 = int32(4358872)
							v37 = *(*int32)(unsafe.Add(mBase, _consts[619]))
							if v37 == int32(0) {
								v48 = v35
								*(*int32)(unsafe.Add(mBase, _consts[623])) = v48
								m.G0 = v5 + int32(16)
								return
							} else {
								v40 = int32(4358880)
								v42 = *(*int32)(unsafe.Add(mBase, _consts[620]))
								if v42 == int32(0) {
									v48 = v40
									*(*int32)(unsafe.Add(mBase, _consts[623])) = v48
									m.G0 = v5 + int32(16)
									return
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, _consts[621]))
									if v46 != 0 {
										v50 = int32(4358908)
										v52 = *(*int32)(unsafe.Add(mBase, _consts[624]))
										v53 = int32(1)
										*(*int32)(unsafe.Add(mBase, _consts[624])) = v52 + v53
										v62 = v52&int32(7)<<(uint(int32(3))%32) + int32(4358832)
										*(*int32)(unsafe.Add(mBase, _consts[623])) = v62
										v65 = *(*int32)(unsafe.Add(mBase, _consts[622]))
										v69 = F_hash_search(m, v65, v62, v53, v5+int32(15))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											v72 = *(*int32)(unsafe.Add(mBase, _consts[623]))
											v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = v73
											*(*int64)(unsafe.Add(mBase, uint32(v72))) = int64(0)
											v77 = int32(4358792)
											v79 = *(*int32)(unsafe.Add(mBase, _consts[605]))
											*(*int32)(unsafe.Add(mBase, _consts[605])) = v79 + int32(1)
											m.G0 = v5 + int32(16)
											return
										}
									} else {
										v48 = int32(4358888)
										*(*int32)(unsafe.Add(mBase, _consts[623])) = v48
										m.G0 = v5 + int32(16)
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
func F_ResetLatch(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
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
	F_errmsg_internal(m, int32(446912), v9+int32(16))
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
	v34 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v39 = F_AllocSetContextCreateInternal(m, v34, int32(133996), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	F_errfinish(m, int32(470380), int32(58), int32(133996))
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
	v41 = int32(4442992)
	v42 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v39
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
	F_ResetUnloggedRelationsInTablespaceDir(m, int32(345315), l0)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v51 = F_AllocateDir(m, int32(466376))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v54 = F_ReadDir(m, v51, int32(466376))
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
	v101 = m.ExcPending
	if v101 != 0 {
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
	v92 = F_ReadDir(m, v51, int32(466376))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L26
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(529293)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v58 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(466376)
	v85 = F_pg_snprintf(m, v9+int32(32), int32(1050), int32(166941), v9)
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
	F_ResetUnloggedRelationsInTablespaceDir(m, v9+int32(32), l0)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	goto L18
L26:
	;
	if v92 != 0 {
		v58 = v92
		goto L16
	} else {
		goto L27
	}
L27:
	;
	goto L17
L28:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v42
	F_MemoryContextDelete(m, v39)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
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
		F_errmsg(m, int32(430133), v5)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_errhint(m, int32(622617), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				F_errfinish(m, int32(472252), int32(94), int32(404829))
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
	var v151 float64
	_ = v151
	var v153 float64
	_ = v153
	var v154 float64
	_ = v154
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 float64
	_ = v166
	var v179 float64
	_ = v179
	var v185 int32
	_ = v185
	var v186 float64
	_ = v186
	var v187 float64
	_ = v187
	var v188 float64
	_ = v188
	var v189 float64
	_ = v189
	var v192 int32
	_ = v192
	var v203 float64
	_ = v203
	var v204 float64
	_ = v204
	var v209 float64
	_ = v209
	var v210 float64
	_ = v210
	var v222 float64
	_ = v222
	var v223 float64
	_ = v223
	var v228 float64
	_ = v228
	var v229 float64
	_ = v229
	var v231 float64
	_ = v231
	var v232 float64
	_ = v232
	var v233 float64
	_ = v233
	var v240 float64
	_ = v240
	var v245 int32
	_ = v245
	var v256 float64
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 float64
	_ = v288
	var v292 float64
	_ = v292
	var v301 int32
	_ = v301
	var v331 float64
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v405 int32
	_ = v405
	var v407 float64
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v479 float64
	_ = v479
	var v481 int32
	_ = v481
	var v505 int32
	_ = v505
	var v508 float64
	_ = v508
	var v515 float64
	_ = v515
	var v517 float64
	_ = v517
	var v519 int32
	_ = v519
	var v521 float64
	_ = v521
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v591 float64
	_ = v591
	var v592 int32
	_ = v592
	var v595 float64
	_ = v595
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v627 float64
	_ = v627
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 float64
	_ = v634
	var v637 float64
	_ = v637
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v650 float64
	_ = v650
	var v651 float64
	_ = v651
	var v653 int32
	_ = v653
	var v655 float64
	_ = v655
	var v684 float64
	_ = v684
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v700 float64
	_ = v700
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v713 float64
	_ = v713
	var v714 int32
	_ = v714
	var v721 float64
	_ = v721
	var v727 float64
	_ = v727
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v735 float64
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v829 int32
	_ = v829
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v865 int32
	_ = v865
	var v869 float64
	_ = v869
	var v870 int32
	_ = v870
	var v871 float64
	_ = v871
	var v875 float64
	_ = v875
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v891 float64
	_ = v891
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v904 float64
	_ = v904
	var v905 int32
	_ = v905
	var v914 float64
	_ = v914
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1049 int32
	_ = v1049
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1088 int32
	_ = v1088
	var v1113 int32
	_ = v1113
	var v1118 int32
	_ = v1118
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1128 float64
	_ = v1128
	var v1131 float64
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1157 int32
	_ = v1157
	var v1160 float64
	_ = v1160
	var v1167 float64
	_ = v1167
	var v1169 float64
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1173 float64
	_ = v1173
	var v1205 int32
	_ = v1205
	var v1209 float64
	_ = v1209
	var v1216 int32
	_ = v1216
	var v1219 int32
	_ = v1219
	var v1225 float64
	_ = v1225
	var v1232 int32
	_ = v1232
	var v1235 int32
	_ = v1235
	var v1238 float64
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1246 float64
	_ = v1246
	var v1255 float64
	_ = v1255
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1265 float64
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1295 int32
	_ = v1295
	var v1300 int32
	_ = v1300
	var v1318 float64
	_ = v1318
	var v1322 float64
	_ = v1322
	var v1329 int32
	_ = v1329
	var v1332 int32
	_ = v1332
	var v1338 float64
	_ = v1338
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1351 float64
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1362 float64
	_ = v1362
	var v1367 int32
	_ = v1367
	var v1396 int32
	_ = v1396
	var v1408 int32
	_ = v1408
	var v1430 float64
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1437 float64
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1463 int32
	_ = v1463
	var v1467 float64
	_ = v1467
	var v1472 float64
	_ = v1472
	var v1474 float64
	_ = v1474
	var v1478 float64
	_ = v1478
	var v1540 float64
	_ = v1540
	var v1543 float64
	_ = v1543
	var v1548 int32
	_ = v1548
	var v1576 float64
	_ = v1576
	var v1577 float64
	_ = v1577
	var v1578 float64
	_ = v1578
	var v1605 float64
	_ = v1605
	var v1607 float64
	_ = v1607
	var v1608 float64
	_ = v1608
	var v1612 float64
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1643 float64
	_ = v1643
	var v1644 float64
	_ = v1644
	var v1648 float64
	_ = v1648
	var v1675 float64
	_ = v1675
	var v1681 int32
	_ = v1681
	var v1682 float64
	_ = v1682
	var v1687 float64
	_ = v1687
	var v1693 float64
	_ = v1693
	var v1698 int32
	_ = v1698
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
	return v1698
L2:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v36) {
		goto L46
	} else {
		goto L47
	}
L3:
	;
	v148 = base.F64_add(base.F64_add(base.F64_mul(l0, float64(0.6366197723675814)), float64(6.755399441055744e+15)), float64(-6.755399441055744e+15))
	v151 = base.F64_add(l0, base.F64_mul(v148, float64(-1.5707963267341256)))
	v153 = base.F64_mul(v148, float64(6.077100506506192e-11))
	v154 = base.F64_sub(v151, v153)
	if base.F64_lt(base.F64_abs(v148), float64(2.147483648e+09)) != 0 {
		goto L33
	} else {
		goto L34
	}
L4:
	;
	if v34&int32(1048575) == int32(598523) {
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
	v1698 = int32(1)
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
	v1698 = int32(-1)
	goto L1
L14:
	;
	v70 = base.F64_add(l0, float64(-3.1415926534682512))
	v71 = float64(-1.2154201013012384e-10)
	v72 = base.F64_add(v70, v71)
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v72
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(base.F64_sub(v70, v72), v71)
	v1698 = int32(2)
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
	v1698 = int32(-2)
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
	v1698 = int32(3)
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
	v1698 = int32(-3)
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
	v1698 = int32(4)
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
	v1698 = int32(-4)
	goto L1
L31:
	;
	goto L3
L32:
	;
	if base.F64_lt(v154, float64(-0.7853981633974483)) != 0 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v160 = base.I32_trunc_f64_s(v148)
	v162 = v160
	goto L32
L34:
	;
	goto L35
L35:
	;
	v162 = int32(-2147483648)
	goto L32
L36:
	;
	v189 = base.F64_sub(v187, v188)
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v189
	v192 = int32(base.Ui32(v36) >> (uint(int32(20)) % 32))
	if v192-base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v189))>>(uint(int64(52))%64)))&int32(2047) < int32(17) {
		v231 = v189
		v232 = v187
		v233 = v188
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v166 = base.F64_add(v148, float64(-1))
	v185 = v162 - int32(1)
	v186 = v166
	v187 = base.F64_add(l0, base.F64_mul(v166, float64(-1.5707963267341256)))
	v188 = base.F64_mul(v166, float64(6.077100506506192e-11))
	goto L36
L38:
	;
	goto L39
L39:
	;
	if base.F64_gt(v154, float64(0.7853981633974483)) == int32(0) {
		v185 = v162
		v186 = v148
		v187 = v151
		v188 = v153
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v179 = base.F64_add(v148, float64(1))
	v185 = v162 + int32(1)
	v186 = v179
	v187 = base.F64_add(l0, base.F64_mul(v179, float64(-1.5707963267341256)))
	v188 = base.F64_mul(v179, float64(6.077100506506192e-11))
	goto L36
L41:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_sub(base.F64_sub(v232, v231), v233)
	v1698 = v185
	goto L1
L42:
	;
	v203 = base.F64_mul(v186, float64(6.077100506303966e-11))
	v204 = base.F64_sub(v187, v203)
	v209 = base.F64_sub(base.F64_mul(v186, float64(2.0222662487959506e-21)), base.F64_sub(base.F64_sub(v187, v204), v203))
	v210 = base.F64_sub(v204, v209)
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v210
	if v192-base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v210))>>(uint(int64(52))%64)))&int32(2047) < int32(50) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v231 = v210
	v232 = v204
	v233 = v209
	goto L41
L44:
	;
	goto L45
L45:
	;
	v222 = base.F64_mul(v186, float64(2.0222662487111665e-21))
	v223 = base.F64_sub(v204, v222)
	v228 = base.F64_sub(base.F64_mul(v186, float64(8.4784276603689e-32)), base.F64_sub(base.F64_sub(v204, v223), v222))
	v229 = base.F64_sub(v223, v228)
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v229
	v231 = v229
	v232 = v223
	v233 = v228
	goto L41
L46:
	;
	v240 = base.F64_sub(l0, l0)
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v240
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v240
	v1698 = int32(0)
	goto L1
L47:
	;
	goto L48
L48:
	;
	v245 = v29 + int32(16)
	v256 = base.F64_reinterpret_i64(v31&int64(4503599627370495) | int64(4710765210229538816))
	v259 = v245
	v261 = int32(1)
	goto L49
L49:
	;
	if base.F64_lt(base.F64_abs(v256), float64(2.147483648e+09)) != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29)+32)) = v292
	v301 = int32(2)
	goto L56
L51:
	;
	v288 = base.F64_convert_i32_s(v287)
	*(*float64)(unsafe.Add(mBase, uint32(v259))) = v288
	v292 = base.F64_mul(base.F64_sub(v256, v288), float64(1.6777216e+07))
	if v261&int32(1) != 0 {
		v256 = v292
		v259 = v245 | int32(8)
		v261 = int32(0)
		goto L49
	} else {
		goto L55
	}
L52:
	;
	v285 = base.I32_trunc_f64_s(v256)
	v287 = v285
	goto L51
L53:
	;
	goto L54
L54:
	;
	v287 = int32(-2147483648)
	goto L51
L55:
	;
	goto L50
L56:
	;
	v331 = *(*float64)(unsafe.Add(mBase, uint32(v29+int32(16)+v301<<(uint(int32(3))%32))))
	if base.F64_eq(v331, float64(0)) != 0 {
		v301 = v301 - int32(1)
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v335 = v29 + int32(16)
	v336 = m.G0
	v338 = v336 - int32(560)
	m.G0 = v338
	v343 = int32(base.Ui32(v36)>>(uint(int32(20))%32)) - int32(1046)
	v347 = base.I32_div_s(v343-int32(3), int32(24))
	v348 = int32(0)
	if v348 < v347 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L57
L59:
	;
	v351 = v347
	goto L61
L60:
	;
	v351 = v348
	goto L61
L61:
	;
	v354 = v343 + v351*int32(-24)
	v356 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	v357 = int32(1)
	v358 = v301 + v357
	v360 = v358 - v357
	if int32(0) <= v356+v360 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v369 = v351 - v360
	v371 = int32(0)
	goto L65
L63:
	;
	goto L64
L64:
	;
	v441 = v354 - int32(24)
	v442 = int32(0)
	if v442 < v356 {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	if v369 < int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L64
L67:
	;
	v407 = float64(0)
	goto L69
L68:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v369<<(uint(int32(2))%32))+uint32(_consts[1459])))
	v407 = base.F64_convert_i32_s(v405)
	goto L69
L69:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v338+int32(320)+v371<<(uint(int32(3))%32)))) = v407
	v409 = int32(1)
	v412 = v371 + v409
	if v412 != v358+v356 {
		v369 = v369 + v409
		v371 = v412
		goto L65
	} else {
		goto L70
	}
L70:
	;
	goto L66
L71:
	;
	v446 = v356
	goto L73
L72:
	;
	v446 = v442
	goto L73
L73:
	;
	v452 = v442
	goto L74
L74:
	;
	if v358 <= int32(0) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v559 = int32(48) - v354
	v565 = v356
	goto L85
L76:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v338+v452<<(uint(int32(3))%32)))) = v521
	if base.B2i32(v452 == v446) == int32(0) {
		v452 = v452 + int32(1)
		goto L74
	} else {
		goto L83
	}
L77:
	;
	v521 = float64(0)
	goto L76
L78:
	;
	goto L79
L79:
	;
	v479 = float64(0)
	v481 = int32(0)
	goto L80
L80:
	;
	v505 = int32(3)
	v508 = *(*float64)(unsafe.Add(mBase, uint32(v335+v481<<(uint(v505)%32))))
	v515 = *(*float64)(unsafe.Add(mBase, uint32(v338+int32(320)+(v452+v360-v481)<<(uint(v505)%32))))
	v517 = base.F64_add(base.F64_mul(v508, v515), v479)
	v519 = v481 + int32(1)
	if v519 != v358 {
		v479 = v517
		v481 = v519
		goto L80
	} else {
		goto L82
	}
L81:
	;
	v521 = v517
	goto L76
L82:
	;
	goto L81
L83:
	;
	goto L75
L84:
	;
	v1318 = float64(1)
	if int32(1024) <= v1300 {
		goto L229
	} else {
		goto L230
	}
L85:
	;
	v591 = *(*float64)(unsafe.Add(mBase, uint32(v338+v565<<(uint(int32(3))%32))))
	v592 = int32(0)
	if v592 < v565 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v1205 = int32(24) - v354
	if int32(1024) <= v1205 {
		goto L196
	} else {
		goto L197
	}
L87:
	;
	v595 = v591
	v597 = v592
	v599 = v565
	goto L90
L88:
	;
	v655 = v591
	goto L89
L89:
	;
	if int32(1024) <= v441 {
		goto L104
	} else {
		goto L105
	}
L90:
	;
	v627 = base.F64_mul(v595, float64(5.960464477539063e-08))
	if base.F64_lt(base.F64_abs(v627), float64(2.147483648e+09)) != 0 {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	v655 = v651
	goto L89
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v338+int32(480)+v597<<(uint(int32(2))%32)))) = v643
	v645 = int32(1)
	v646 = v599 - v645
	v650 = *(*float64)(unsafe.Add(mBase, uint32(v338+v646<<(uint(int32(3))%32))))
	v651 = base.F64_add(v650, v634)
	v653 = v597 + v645
	if v653 != v565 {
		v595 = v651
		v597 = v653
		v599 = v646
		goto L90
	} else {
		goto L100
	}
L93:
	;
	v634 = base.F64_convert_i32_s(v633)
	v637 = base.F64_add(base.F64_mul(v634, float64(-1.6777216e+07)), v595)
	if base.F64_lt(base.F64_abs(v637), float64(2.147483648e+09)) != 0 {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	v631 = base.I32_trunc_f64_s(v627)
	v633 = v631
	goto L93
L95:
	;
	goto L96
L96:
	;
	v633 = int32(-2147483648)
	goto L93
L97:
	;
	v641 = base.I32_trunc_f64_s(v637)
	v643 = v641
	goto L92
L98:
	;
	goto L99
L99:
	;
	v643 = int32(-2147483648)
	goto L92
L100:
	;
	goto L91
L101:
	;
	v735 = base.F64_sub(v727, base.F64_convert_i32_s(v733))
	v736 = int32(0)
	v737 = base.B2i32(v441 <= v736)
	if v737 == v736 {
		goto L127
	} else {
		goto L128
	}
L102:
	;
	v727 = base.F64_add(v721, base.F64_mul(base.F64_floor(base.F64_mul(v721, float64(0.125))), float64(-8)))
	if base.F64_lt(base.F64_abs(v727), float64(2.147483648e+09)) != 0 {
		goto L120
	} else {
		goto L121
	}
L103:
	;
	v721 = base.F64_mul(v713, base.F64_reinterpret_i64(base.I64_extend_i32_u(v714+int32(1023))<<(uint(int64(52))%64)))
	goto L102
L104:
	;
	v684 = base.F64_mul(v655, float64(8.98846567431158e+307))
	if base.Ui32(v441) < base.Ui32(int32(2047)) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	if int32(-1023) < v441 {
		v713 = v655
		v714 = v441
		goto L103
	} else {
		goto L113
	}
L107:
	;
	v713 = v684
	v714 = v441 - int32(1023)
	goto L103
L108:
	;
	goto L109
L109:
	;
	v691 = int32(3069)
	if base.Ui32(v691) <= base.Ui32(v441) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v694 = v691
	goto L112
L111:
	;
	v694 = v441
	goto L112
L112:
	;
	v713 = base.F64_mul(v684, float64(8.98846567431158e+307))
	v714 = v694 - int32(2046)
	goto L103
L113:
	;
	v700 = base.F64_mul(v655, float64(2.004168360008973e-292))
	if base.Ui32(int32(-1992)) < base.Ui32(v441) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v713 = v700
	v714 = v441 + int32(969)
	goto L103
L115:
	;
	goto L116
L116:
	;
	v707 = int32(-2960)
	if base.Ui32(v441) <= base.Ui32(v707) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v710 = v707
	goto L119
L118:
	;
	v710 = v441
	goto L119
L119:
	;
	v713 = base.F64_mul(v700, float64(2.004168360008973e-292))
	v714 = v710 + int32(1938)
	goto L103
L120:
	;
	v731 = base.I32_trunc_f64_s(v727)
	v733 = v731
	goto L101
L121:
	;
	goto L122
L122:
	;
	v733 = int32(-2147483648)
	goto L101
L123:
	;
	if base.F64_eq(v914, float64(0)) != 0 {
		goto L169
	} else {
		goto L170
	}
L124:
	;
	v772 = int32(0)
	if v772 < v565 {
		goto L133
	} else {
		goto L134
	}
L125:
	;
	if base.F64_ge(v735, float64(0.5)) != 0 {
		v770 = v733
		v771 = int32(2)
		goto L124
	} else {
		goto L132
	}
L126:
	;
	if v761 <= int32(0) {
		v914 = v735
		v925 = v760
		v927 = v761
		goto L123
	} else {
		goto L131
	}
L127:
	;
	v744 = v565<<(uint(int32(2))%32) + v338 + int32(476)
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v744)))
	v746 = v745 >> (uint(v559) % 32)
	v748 = v745 - v746<<(uint(v559)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v744))) = v748
	v760 = v746 + v733
	v761 = v748 >> (uint(int32(47)-v354) % 32)
	goto L126
L128:
	;
	goto L129
L129:
	;
	if v441 != 0 {
		goto L125
	} else {
		goto L130
	}
L130:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v565<<(uint(int32(2))%32)+v338)+476))
	v760 = v733
	v761 = v755 >> (uint(int32(23)) % 32)
	goto L126
L131:
	;
	v770 = v760
	v771 = v761
	goto L124
L132:
	;
	v914 = v735
	v925 = v733
	v927 = int32(0)
	goto L123
L133:
	;
	v779 = v772
	v784 = v772
	goto L136
L134:
	;
	v829 = int32(1)
	goto L135
L135:
	;
	if v441 <= v736 {
		goto L145
	} else {
		goto L146
	}
L136:
	;
	v807 = v338 + int32(480) + v779<<(uint(int32(2))%32)
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v807)))
	if v784 != 0 {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	v829 = v821
	goto L135
L138:
	;
	v823 = v779 + int32(1)
	if v823 != v565 {
		v779 = v823
		v784 = v820
		goto L136
	} else {
		goto L144
	}
L139:
	;
	v820 = int32(0)
	v821 = int32(1)
	goto L138
L140:
	;
	v813 = int32(16777215)
	goto L142
L141:
	;
	if v808 == int32(0) {
		goto L139
	} else {
		goto L143
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v807))) = v813 - v808
	v820 = int32(1)
	v821 = int32(0)
	goto L138
L143:
	;
	v813 = int32(16777216)
	goto L142
L144:
	;
	goto L137
L145:
	;
	v865 = v770 + int32(1)
	if v771 != int32(2) {
		v914 = v735
		v925 = v865
		v927 = v771
		goto L123
	} else {
		goto L149
	}
L146:
	;
	switch v354 - int32(25) {
	case 0:
		v853 = int32(8388607)
		goto L147
	case 1:
		goto L148
	default:
		goto L145
	}
L147:
	;
	v858 = v565<<(uint(int32(2))%32) + v338 + int32(476)
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v858)))
	*(*int32)(unsafe.Add(mBase, uint32(v858))) = v859 & v853
	goto L145
L148:
	;
	v853 = int32(4194303)
	goto L147
L149:
	;
	v869 = base.F64_sub(float64(1), v735)
	v870 = int32(2)
	if v829 != 0 {
		v914 = v869
		v925 = v865
		v927 = v870
		goto L123
	} else {
		goto L150
	}
L150:
	;
	v871 = float64(1)
	if int32(1024) <= v441 {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	v914 = base.F64_sub(v869, base.F64_mul(v904, base.F64_reinterpret_i64(base.I64_extend_i32_u(v905+int32(1023))<<(uint(int64(52))%64))))
	v925 = v865
	v927 = v870
	goto L123
L152:
	;
	goto L151
L153:
	;
	v875 = base.F64_mul(v871, float64(8.98846567431158e+307))
	if base.Ui32(v441) < base.Ui32(int32(2047)) {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	goto L155
L155:
	;
	if int32(-1023) < v441 {
		v904 = v871
		v905 = v441
		goto L152
	} else {
		goto L162
	}
L156:
	;
	v904 = v875
	v905 = v441 - int32(1023)
	goto L152
L157:
	;
	goto L158
L158:
	;
	v882 = int32(3069)
	if base.Ui32(v882) <= base.Ui32(v441) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v885 = v882
	goto L161
L160:
	;
	v885 = v441
	goto L161
L161:
	;
	v904 = base.F64_mul(v875, float64(8.98846567431158e+307))
	v905 = v885 - int32(2046)
	goto L152
L162:
	;
	v891 = base.F64_mul(v871, float64(2.004168360008973e-292))
	if base.Ui32(int32(-1992)) < base.Ui32(v441) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v904 = v891
	v905 = v441 + int32(969)
	goto L152
L164:
	;
	goto L165
L165:
	;
	v898 = int32(-2960)
	if base.Ui32(v441) <= base.Ui32(v898) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v901 = v898
	goto L168
L167:
	;
	v901 = v441
	goto L168
L168:
	;
	v904 = base.F64_mul(v891, float64(2.004168360008973e-292))
	v905 = v901 + int32(1938)
	goto L152
L169:
	;
	if v565 <= v356 {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	goto L171
L171:
	;
	goto L86
L172:
	;
	v1049 = int32(1)
	goto L181
L173:
	;
	v946 = v565
	v948 = int32(0)
	goto L174
L174:
	;
	v973 = v946 - int32(1)
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v338+int32(480)+v973<<(uint(int32(2))%32))))
	v978 = v977 | v948
	if v356 < v973 {
		v946 = v973
		v948 = v978
		goto L174
	} else {
		goto L176
	}
L175:
	;
	if v978 == int32(0) {
		goto L172
	} else {
		goto L177
	}
L176:
	;
	goto L175
L177:
	;
	v985 = v565
	v990 = v441
	goto L178
L178:
	;
	v1009 = v990 - int32(24)
	v1013 = v985 - int32(1)
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v338+int32(480)+v1013<<(uint(int32(2))%32))))
	if v1017 == int32(0) {
		v985 = v1013
		v990 = v1009
		goto L178
	} else {
		goto L180
	}
L179:
	;
	v1295 = v1013
	v1300 = v1009
	goto L84
L180:
	;
	goto L179
L181:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v338+int32(480)+(v356-v1049)<<(uint(int32(2))%32))))
	if v1081 == int32(0) {
		v1049 = v1049 + int32(1)
		goto L181
	} else {
		goto L183
	}
L182:
	;
	v1084 = v565 + v1049
	v1088 = v565
	goto L184
L183:
	;
	goto L182
L184:
	;
	v1113 = v1088 + v358
	v1118 = v1088 + int32(1)
	v1124 = *(*int32)(unsafe.Add(mBase, uint32((v1118+v351)<<(uint(int32(2))%32))+uint32(_consts[1459])))
	*(*float64)(unsafe.Add(mBase, uint32(v338+int32(320)+v1113<<(uint(int32(3))%32)))) = base.F64_convert_i32_s(v1124)
	v1127 = int32(0)
	v1128 = float64(0)
	if v1127 < v358 {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v565 = v1084
	goto L85
L186:
	;
	v1131 = v1128
	v1133 = v1127
	goto L189
L187:
	;
	v1173 = v1128
	goto L188
L188:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v338+v1118<<(uint(int32(3))%32)))) = v1173
	if v1118 < v1084 {
		v1088 = v1118
		goto L184
	} else {
		goto L192
	}
L189:
	;
	v1157 = int32(3)
	v1160 = *(*float64)(unsafe.Add(mBase, uint32(v335+v1133<<(uint(v1157)%32))))
	v1167 = *(*float64)(unsafe.Add(mBase, uint32(v338+int32(320)+(v1113-v1133)<<(uint(v1157)%32))))
	v1169 = base.F64_add(base.F64_mul(v1160, v1167), v1131)
	v1171 = v1133 + int32(1)
	if v1171 != v358 {
		v1131 = v1169
		v1133 = v1171
		goto L189
	} else {
		goto L191
	}
L190:
	;
	v1173 = v1169
	goto L188
L191:
	;
	goto L190
L192:
	;
	goto L185
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v338+int32(480)+v1283<<(uint(int32(2))%32)))) = v1282
	v1295 = v1283
	v1300 = v1284
	goto L84
L194:
	;
	if base.F64_ge(v1246, float64(1.6777216e+07)) != 0 {
		goto L212
	} else {
		goto L213
	}
L195:
	;
	v1246 = base.F64_mul(v1238, base.F64_reinterpret_i64(base.I64_extend_i32_u(v1239+int32(1023))<<(uint(int64(52))%64)))
	goto L194
L196:
	;
	v1209 = base.F64_mul(v914, float64(8.98846567431158e+307))
	if base.Ui32(v1205) < base.Ui32(int32(2047)) {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	goto L198
L198:
	;
	if int32(-1023) < v1205 {
		v1238 = v914
		v1239 = v1205
		goto L195
	} else {
		goto L205
	}
L199:
	;
	v1238 = v1209
	v1239 = v1205 - int32(1023)
	goto L195
L200:
	;
	goto L201
L201:
	;
	v1216 = int32(3069)
	if base.Ui32(v1216) <= base.Ui32(v1205) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v1219 = v1216
	goto L204
L203:
	;
	v1219 = v1205
	goto L204
L204:
	;
	v1238 = base.F64_mul(v1209, float64(8.98846567431158e+307))
	v1239 = v1219 - int32(2046)
	goto L195
L205:
	;
	v1225 = base.F64_mul(v914, float64(2.004168360008973e-292))
	if base.Ui32(int32(-1992)) < base.Ui32(v1205) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v1238 = v1225
	v1239 = v1205 + int32(969)
	goto L195
L207:
	;
	goto L208
L208:
	;
	v1232 = int32(-2960)
	if base.Ui32(v1205) <= base.Ui32(v1232) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v1235 = v1232
	goto L211
L210:
	;
	v1235 = v1205
	goto L211
L211:
	;
	v1238 = base.F64_mul(v1225, float64(2.004168360008973e-292))
	v1239 = v1235 + int32(1938)
	goto L195
L212:
	;
	v1255 = base.F64_mul(v1246, float64(5.960464477539063e-08))
	if base.F64_lt(base.F64_abs(v1255), float64(2.147483648e+09)) != 0 {
		goto L217
	} else {
		goto L218
	}
L213:
	;
	goto L214
L214:
	;
	if base.F64_lt(base.F64_abs(v1246), float64(2.147483648e+09)) != 0 {
		goto L224
	} else {
		goto L225
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v338+int32(480)+v565<<(uint(int32(2))%32)))) = v1271
	v1282 = v1261
	v1283 = v565 + int32(1)
	v1284 = v354
	goto L193
L216:
	;
	v1265 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1261), float64(-1.6777216e+07)), v1246)
	if base.F64_lt(base.F64_abs(v1265), float64(2.147483648e+09)) != 0 {
		goto L220
	} else {
		goto L221
	}
L217:
	;
	v1259 = base.I32_trunc_f64_s(v1255)
	v1261 = v1259
	goto L216
L218:
	;
	goto L219
L219:
	;
	v1261 = int32(-2147483648)
	goto L216
L220:
	;
	v1269 = base.I32_trunc_f64_s(v1265)
	v1271 = v1269
	goto L215
L221:
	;
	goto L222
L222:
	;
	v1271 = int32(-2147483648)
	goto L215
L223:
	;
	v1282 = v1280
	v1283 = v565
	v1284 = v441
	goto L193
L224:
	;
	v1278 = base.I32_trunc_f64_s(v1246)
	v1280 = v1278
	goto L223
L225:
	;
	goto L226
L226:
	;
	v1280 = int32(-2147483648)
	goto L223
L227:
	;
	if int32(0) <= v1295 {
		goto L245
	} else {
		goto L246
	}
L228:
	;
	goto L227
L229:
	;
	v1322 = base.F64_mul(v1318, float64(8.98846567431158e+307))
	if base.Ui32(v1300) < base.Ui32(int32(2047)) {
		goto L232
	} else {
		goto L233
	}
L230:
	;
	goto L231
L231:
	;
	if int32(-1023) < v1300 {
		v1351 = v1318
		v1352 = v1300
		goto L228
	} else {
		goto L238
	}
L232:
	;
	v1351 = v1322
	v1352 = v1300 - int32(1023)
	goto L228
L233:
	;
	goto L234
L234:
	;
	v1329 = int32(3069)
	if base.Ui32(v1329) <= base.Ui32(v1300) {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v1332 = v1329
	goto L237
L236:
	;
	v1332 = v1300
	goto L237
L237:
	;
	v1351 = base.F64_mul(v1322, float64(8.98846567431158e+307))
	v1352 = v1332 - int32(2046)
	goto L228
L238:
	;
	v1338 = base.F64_mul(v1318, float64(2.004168360008973e-292))
	if base.Ui32(int32(-1992)) < base.Ui32(v1300) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1351 = v1338
	v1352 = v1300 + int32(969)
	goto L228
L240:
	;
	goto L241
L241:
	;
	v1345 = int32(-2960)
	if base.Ui32(v1300) <= base.Ui32(v1345) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1348 = v1345
	goto L244
L243:
	;
	v1348 = v1300
	goto L244
L244:
	;
	v1351 = base.F64_mul(v1338, float64(2.004168360008973e-292))
	v1352 = v1348 + int32(1938)
	goto L228
L245:
	;
	v1362 = base.F64_mul(v1351, base.F64_reinterpret_i64(base.I64_extend_i32_u(v1352+int32(1023))<<(uint(int64(52))%64)))
	v1367 = v1295
	goto L248
L246:
	;
	goto L247
L247:
	;
	v1540 = float64(0)
	if int32(0) <= v1295 {
		goto L263
	} else {
		goto L264
	}
L248:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v338+int32(480)+v1367<<(uint(int32(2))%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v338+v1367<<(uint(int32(3))%32)))) = base.F64_mul(v1362, base.F64_convert_i32_s(v1396))
	if v1367 != 0 {
		v1362 = base.F64_mul(v1362, float64(5.960464477539063e-08))
		v1367 = v1367 - int32(1)
		goto L248
	} else {
		goto L250
	}
L249:
	;
	v1408 = v1295
	goto L251
L250:
	;
	goto L249
L251:
	;
	v1430 = float64(0)
	v1432 = v1295 - v1408
	if v356 < v1432 {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	goto L247
L253:
	;
	v1434 = v356
	goto L255
L254:
	;
	v1434 = v1432
	goto L255
L255:
	;
	if int32(0) <= v1434 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v1437 = v1430
	v1439 = int32(0)
	goto L259
L257:
	;
	v1478 = v1430
	goto L258
L258:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v338+int32(160)+v1432<<(uint(int32(3))%32)))) = v1478
	if int32(0) < v1408 {
		v1408 = v1408 - int32(1)
		goto L251
	} else {
		goto L262
	}
L259:
	;
	v1463 = int32(3)
	v1467 = *(*float64)(unsafe.Add(mBase, uint32(v1439<<(uint(v1463)%32))+uint32(_consts[1460])))
	v1472 = *(*float64)(unsafe.Add(mBase, uint32(v338+(v1439+v1408)<<(uint(v1463)%32))))
	v1474 = base.F64_add(base.F64_mul(v1467, v1472), v1437)
	if v1439 != v1434 {
		v1437 = v1474
		v1439 = v1439 + int32(1)
		goto L259
	} else {
		goto L261
	}
L260:
	;
	v1478 = v1474
	goto L258
L261:
	;
	goto L260
L262:
	;
	goto L252
L263:
	;
	v1543 = v1540
	v1548 = v1295
	goto L266
L264:
	;
	v1578 = v1540
	goto L265
L265:
	;
	if v927 != 0 {
		goto L269
	} else {
		goto L270
	}
L266:
	;
	v1576 = *(*float64)(unsafe.Add(mBase, uint32(v338+int32(160)+v1548<<(uint(int32(3))%32))))
	v1577 = base.F64_add(v1543, v1576)
	if v1548 != 0 {
		v1543 = v1577
		v1548 = v1548 - int32(1)
		goto L266
	} else {
		goto L268
	}
L267:
	;
	v1578 = v1577
	goto L265
L268:
	;
	goto L267
L269:
	;
	v1605 = base.F64_neg(v1578)
	goto L271
L270:
	;
	v1605 = v1578
	goto L271
L271:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29))) = v1605
	v1607 = *(*float64)(unsafe.Add(mBase, uint32(v338)+160))
	v1608 = base.F64_sub(v1607, v1578)
	if int32(0) < v1295 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v1612 = v1608
	v1614 = int32(1)
	goto L275
L273:
	;
	v1648 = v1608
	goto L274
L274:
	;
	if v927 != 0 {
		goto L278
	} else {
		goto L279
	}
L275:
	;
	v1643 = *(*float64)(unsafe.Add(mBase, uint32(v338+int32(160)+v1614<<(uint(int32(3))%32))))
	v1644 = base.F64_add(v1612, v1643)
	if v1614 != v1295 {
		v1612 = v1644
		v1614 = v1614 + int32(1)
		goto L275
	} else {
		goto L277
	}
L276:
	;
	v1648 = v1644
	goto L274
L277:
	;
	goto L276
L278:
	;
	v1675 = base.F64_neg(v1648)
	goto L280
L279:
	;
	v1675 = v1648
	goto L280
L280:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29)+8)) = v1675
	m.G0 = v338 + int32(560)
	v1681 = v925 & int32(7)
	v1682 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
	if v31 < int64(0) {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_neg(v1682)
	v1687 = *(*float64)(unsafe.Add(mBase, uint32(v29)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_neg(v1687)
	v1698 = int32(0) - v1681
	goto L1
L282:
	;
	goto L283
L283:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v1682
	v1693 = *(*float64)(unsafe.Add(mBase, uint32(v29)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v1693
	v1698 = v1681
	goto L1
}
func F_r_KER_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v11 < v10 {
		v13 = v10
	} else {
		v13 = v11
	}
	if v10 == v13 {
		v51 = int32(-1)
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v10))))
		if int32(117) < v26 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10 + int32(1)
			v48 = int32(0)
		} else {
			v28 = v26 - int32(97)
			if v28 < int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10 + int32(1)
				v48 = int32(0)
			} else {
				v31 = int32(1)
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v28)>>(uint(int32(3))%32)))+uint32(_consts[1442]))))
				if int32(base.Ui32(v35)>>(uint(v28&int32(7))%32))&v31 != 0 {
					v48 = v31
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10 + int32(1)
					v48 = int32(0)
				}
			}
		}
		v51 = v48
	}
	if v51 != 0 {
		v69 = int32(0)
	} else {
		v52 = int32(2)
		v54 = int32(0)
		v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v56-v57 < v52 {
			v66 = v54
		} else {
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v62 = F_memcmp(m, v60+v57, int32(2126361), v52)
			mBase = m.M
			if v62 != 0 {
				v66 = v54
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v52 + v57
				v66 = int32(1)
			}
		}
		v69 = base.B2i32(v66 != int32(0))
	}
	return v69
}
func F_r_SUFFIX_I_OK(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if int32(2) < v5 {
		v23 = int32(0)
	} else {
		v8 = int32(1)
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v9 <= v10 {
			v23 = v8
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v9-int32(1)))))
			if v16 != int32(115) {
				v23 = v8
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9 - int32(1)
				v23 = int32(0)
			}
		}
	}
	return v23
}
func F_r_Suffix_Noun_Step2a(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
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
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	v2 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3
	v7 = F_find_among_b(m, l0, int32(4174784), int32(3))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 == int32(0) {
			v105 = v2
			return v105
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v13
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v16 = int32(0)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v15-int32(4))))
			if v23 == v16 {
				v95 = int32(0)
			} else {
				v28 = v23 & int32(3)
				if base.Ui32(v23) < base.Ui32(int32(4)) {
					v62 = v15
					v63 = int32(0)
				} else {
					v35 = v15
					v36 = int32(0)
					v39 = v16
					for {
						v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35))))
						v42 = int32(-65)
						v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35)+1)))
						v49 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35)+2)))
						v53 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35)+3)))
						v56 = v36 + base.B2i32(v42 < v41) + base.B2i32(v42 < v45) + base.B2i32(v42 < v49) + base.B2i32(v42 < v53)
						v57 = int32(4)
						v58 = v35 + v57
						v60 = v39 + v57
						if v60 != v23&int32(-4) {
							v35 = v58
							v36 = v56
							v39 = v60
							continue
						} else {
							break
						}
						break
					}
					v62 = v58
					v63 = v56
				}
				if v28 != 0 {
					v68 = v62
					v69 = v63
					v71 = v16
					for {
						v74 = int32(*(*int8)(unsafe.Add(mBase, uint32(v68))))
						v77 = v69 + base.B2i32(int32(-65) < v74)
						v78 = int32(1)
						v81 = v71 + v78
						if v81 != v28 {
							v68 = v68 + v78
							v69 = v77
							v71 = v81
							continue
						} else {
							break
						}
						break
					}
					v84 = v77
				} else {
					v84 = v63
				}
				v95 = v84
			}
			if v95 < int32(5) {
				v105 = v2
				return v105
			} else {
				v99 = F_slice_del(m, l0)
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return int32(0)
				} else {
					if int32(0) <= v99 {
						v103 = int32(1)
					} else {
						v103 = v99
					}
					v105 = v103
					return v105
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v4
	v7 = v4 + int32(5)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v8 <= v7 {
		v62 = v2
		return v62
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v7))))
		if v12&int32(224) != int32(128) {
			v62 = v2
			return v62
		} else {
			if int32(1)<<(uint(v12)%32)&int32(3078) == int32(0) {
				v62 = v2
				return v62
			} else {
				v25 = F_find_among(m, l0, int32(4319280), int32(4))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					if v25 == int32(0) {
						v62 = v2
						return v62
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v31
						switch v25 - int32(1) {
						case 0:
							v37 = F_slice_from_s(m, l0, int32(3), int32(2172514))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								if int32(0) <= v37 {
									v62 = int32(1)
								} else {
									v62 = v37
								}
								return v62
							}
						case 1:
							v43 = F_slice_from_s(m, l0, int32(3), int32(2172517))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								if int32(0) <= v43 {
									v62 = int32(1)
								} else {
									v62 = v43
								}
								return v62
							}
						case 2:
							v49 = F_slice_from_s(m, l0, int32(3), int32(2172520))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								if int32(0) <= v49 {
									v62 = int32(1)
								} else {
									v62 = v49
								}
								return v62
							}
						case 3:
							v55 = F_slice_from_s(m, l0, int32(3), int32(2172523))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								if v55 < int32(0) {
									v62 = v55
								} else {
									v62 = int32(1)
								}
								return v62
							}
						default:
							v62 = int32(1)
							return v62
						}
					}
				}
			}
		}
	}
}
func F_rangesel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 float64
	_ = v10
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
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
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 float64
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 float64
	_ = v74
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
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
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 float64
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 float32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 float32
	_ = v142
	var v146 int32
	_ = v146
	var v148 float64
	_ = v148
	var v152 float64
	_ = v152
	var v155 float64
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
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
	var v225 int32
	_ = v225
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int64
	_ = v290
	var v307 int32
	_ = v307
	var v313 float64
	_ = v313
	var v314 int32
	_ = v314
	var v319 float64
	_ = v319
	var v320 int32
	_ = v320
	var v326 float64
	_ = v326
	var v327 int32
	_ = v327
	var v332 float64
	_ = v332
	var v333 int32
	_ = v333
	var v338 float64
	_ = v338
	var v339 int32
	_ = v339
	var v345 float64
	_ = v345
	var v346 int32
	_ = v346
	var v351 float64
	_ = v351
	var v352 int32
	_ = v352
	var v357 float64
	_ = v357
	var v358 int32
	_ = v358
	var v363 float64
	_ = v363
	var v364 int32
	_ = v364
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 float64
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v382 float64
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v391 float64
	_ = v391
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 float64
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v418 float64
	_ = v418
	var v419 int32
	_ = v419
	var v421 float64
	_ = v421
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v450 int32
	_ = v450
	var v468 int32
	_ = v468
	var v475 float64
	_ = v475
	var v485 float64
	_ = v485
	var v492 float64
	_ = v492
	var v507 float64
	_ = v507
	var v513 float64
	_ = v513
	var v516 float64
	_ = v516
	var v528 int32
	_ = v528
	var v535 float64
	_ = v535
	var v546 float64
	_ = v546
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 float64
	_ = v555
	var v558 float64
	_ = v558
	var v563 float64
	_ = v563
	var v568 float64
	_ = v568
	var v579 float64
	_ = v579
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	v10 = float64(0)
	v16 = m.G0
	v18 = v16 - int32(160)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v30 = F_get_restriction_variable(m, v21, v22, v23, v18+int32(36), v18+int32(32), v18+int32(31))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L5
	} else {
		goto L154
	}
L2:
	;
	v584 = F_Float8GetDatum(m, v579)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L5
	} else {
		goto L153
	}
L3:
	;
	v568 = *(*float64)(unsafe.Add(mBase, uint32(v50<<(uint(int32(3))%32))+uint32(_consts[1110])))
	v579 = v568
	goto L2
L4:
	;
	v563 = *(*float64)(unsafe.Add(mBase, uint32(v37<<(uint(int32(3))%32))+uint32(_consts[1110])))
	v579 = v563
	goto L2
L5:
	;
	return int32(0)
L6:
	;
	if v30 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v37 = v20 - int32(3884)
	if base.Ui32(v37) < base.Ui32(int32(13)) {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v42 != int32(7) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v579 = float64(0.01)
	goto L2
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	if v45 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+24)))
	if v54 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	m.T0[v46].(func(*base.Module, int32))(m, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v50 = v20 - int32(3884)
	if base.Ui32(v50) < base.Ui32(int32(13)) {
		goto L3
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v579 = float64(0.01)
	goto L2
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	if v57 == int32(0) {
		v579 = v10
		goto L2
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+31)))
	if v63 != 0 {
		v73 = v20
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	m.T0[v60].(func(*base.Module, int32))(m, v57)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v579 = v10
	goto L2
L24:
	;
	v74 = float64(0.005)
	switch v73 - int32(3889) {
	case 0:
		goto L34
	default:
		goto L33
	case 2:
		v546 = v74
		goto L30
	}
L25:
	;
	v64 = F_get_commutator(m, v20)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	if v64 != 0 {
		v73 = v64
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v66 = float64(0.01)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	if v67 == int32(0) {
		v579 = v66
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	m.T0[v70].(func(*base.Module, int32))(m, v67)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v579 = v66
	goto L2
L30:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	if v551 != 0 {
		goto L146
	} else {
		goto L147
	}
L31:
	;
	v528 = v73 - int32(3884)
	if base.Ui32(int32(12)) < base.Ui32(v528) {
		v546 = float64(0.01)
		goto L30
	} else {
		goto L145
	}
L32:
	;
	if v118 == int32(0) {
		goto L31
	} else {
		goto L41
	}
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	if v107 != v108 {
		goto L31
	} else {
		goto L38
	}
L34:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v78 = F_range_get_typcache(m, l0, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+200))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v81 != v83 {
		v546 = v74
		goto L30
	} else {
		goto L36
	}
L36:
	;
	v85 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+129)) = uint8(v85)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+130)) = uint8(v85)
	v90 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+128)) = uint8(v90)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = v87
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+94)) = uint8(v90)
	v95 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+92)) = uint16(v95)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v87
	v104 = F_range_serialize(m, v78, v18+int32(124), v18+int32(88), v90, v90)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v116 = v78
	v118 = v104
	goto L32
L38:
	;
	v110 = F_range_get_typcache(m, l0, v107)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	v114 = F_pg_detoast_datum(m, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v116 = v110
	v118 = v114
	goto L32
L41:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	if v121 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v162 = int32(*(*int8)(unsafe.Add(mBase, uint32(v118+int32(base.Ui32(v156)>>(uint(int32(2))%32))-int32(1)))))
	goto L53
L43:
	;
	v124 = float64(0)
	v152 = v124
	v155 = v124
	goto L42
L44:
	;
	goto L45
L45:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v121)+16))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+22)))
	v129 = *(*float32)(unsafe.Add(mBase, uint32(v126+v127)+8))
	v136 = F_get_attstatsslot(m, v18+int32(124), v121, int32(6), int32(0), int32(2))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	if v136 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v18)+148))
	if v138 != int32(1) {
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
	v152 = v148
	v155 = base.F64_promote_f32(v129)
	goto L42
L50:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v18)+144))
	v142 = *(*float32)(unsafe.Add(mBase, uint32(v141)))
	F_free_attstatsslot(m, v18+int32(124))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	v148 = base.F64_promote_f32(v142)
	goto L49
L52:
	;
	v513 = float64(0)
	v516 = base.F64_mul(base.F64_sub(float64(1), v155), v507)
	if base.F64_lt(v516, v513) != 0 {
		v546 = v513
		goto L30
	} else {
		goto L143
	}
L53:
	;
	if v162&int32(1) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	switch v73 - int32(3884) {
	case 0, 4, 9, 10, 11, 12:
		v507 = v10
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
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v116)+216))
	v188 = F_statistic_proc_security_check(m, v18+int32(36), v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L5
	} else {
		goto L66
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L5
	} else {
		goto L61
	}
L58:
	;
	v507 = base.F64_sub(float64(1), v152)
	goto L52
L59:
	;
	v507 = float64(1)
	goto L52
L60:
	;
	v507 = v152
	goto L52
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v73
	F_errmsg_internal(m, int32(41068), v18+int32(16))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(471911), int32(320), int32(292265))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
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
	v492 = base.F64_sub(float64(1), v152)
	if v73 == int32(3892) {
		goto L140
	} else {
		goto L141
	}
L65:
	;
	v468 = v73 - int32(3884)
	if base.Ui32(int32(12)) < base.Ui32(v468) {
		v485 = float64(0.01)
		goto L64
	} else {
		goto L139
	}
L66:
	;
	if v188 == int32(0) {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v116)+272))
	if v192 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v195 = F_statistic_proc_security_check(m, v18+int32(36), v192)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L5
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	if v199 == int32(0) {
		goto L65
	} else {
		goto L73
	}
L71:
	;
	if v195 == int32(0) {
		goto L65
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v207 = F_get_attstatsslot(m, v18+int32(124), v199, int32(7), int32(0), int32(1))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	if v207 == int32(0) {
		goto L65
	} else {
		goto L75
	}
L75:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
	if v211 < int32(2) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_free_attstatsslot(m, v18+int32(124))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L5
	} else {
		goto L138
	}
L77:
	;
	v216 = v211 << (uint(int32(3)) % 32)
	v217 = F_palloc(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	v219 = F_palloc(m, v216)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	v225 = int32(0)
	goto L81
L80:
	;
	switch v73 - int32(3890) {
	case 0, 2:
		goto L94
	default:
		goto L93
	}
L81:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v236+v225<<(uint(int32(2))%32))))
	v241 = F_pg_detoast_datum(m, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L5
	} else {
		goto L84
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L5
	} else {
		goto L88
	}
L83:
	;
	goto L82
L84:
	;
	v244 = v225 << (uint(int32(3)) % 32)
	F_range_deserialize(m, v116, v241, v217+v244, v244+v219, v18+int32(71))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+71)))
	if v251 == int32(1) {
		goto L83
	} else {
		goto L86
	}
L86:
	;
	v255 = v225 + int32(1)
	if v211 != v255 {
		v225 = v255
		goto L81
	} else {
		goto L87
	}
L87:
	;
	goto L80
L88:
	;
	F_errmsg_internal(m, int32(382746), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(471911), int32(423), int32(8825))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
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
	F_free_attstatsslot(m, v18+int32(88))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L5
	} else {
		goto L135
	}
L92:
	;
	F_range_deserialize(m, v116, v118, v18+int32(80), v18+int32(72), v18+int32(71))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L5
	} else {
		goto L99
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = int32(0)
	v290 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = v290
	*(*int64)(unsafe.Add(mBase, uint32(v18)+104)) = v290
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = v290
	*(*int64)(unsafe.Add(mBase, uint32(v18)+88)) = v290
	goto L92
L94:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	if v272 == int32(0) {
		goto L76
	} else {
		goto L95
	}
L95:
	;
	v280 = F_get_attstatsslot(m, v18+int32(88), v272, int32(6), int32(0), int32(1))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	if v280 == int32(0) {
		goto L76
	} else {
		goto L97
	}
L97:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	if int32(2) <= v285 {
		goto L92
	} else {
		goto L98
	}
L98:
	;
	v421 = float64(-1)
	goto L91
L99:
	;
	switch v73 - int32(3884) {
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
	v418 = F_calc_hist_selectivity_scalar(m, v116, v18+int32(80), v217, v211, int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L5
	} else {
		goto L134
	}
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L5
	} else {
		goto L131
	}
L102:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+84)))
	if v376 == int32(1) {
		goto L122
	} else {
		goto L123
	}
L103:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v18)+100))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	v374 = F_calc_hist_selectivity_contains(m, v116, v18+int32(80), v18+int32(72), v217, v211, v372, v373)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L5
	} else {
		goto L121
	}
L104:
	;
	v357 = F_calc_hist_selectivity_scalar(m, v116, v18+int32(80), v219, v211, int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L5
	} else {
		goto L119
	}
L105:
	;
	v351 = F_calc_hist_selectivity_scalar(m, v116, v18+int32(72), v219, v211, int32(1))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L5
	} else {
		goto L118
	}
L106:
	;
	v345 = F_calc_hist_selectivity_scalar(m, v116, v18+int32(80), v217, v211, int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L5
	} else {
		goto L117
	}
L107:
	;
	v338 = F_calc_hist_selectivity_scalar(m, v116, v18+int32(72), v217, v211, int32(1))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L5
	} else {
		goto L116
	}
L108:
	;
	v332 = F_calc_hist_selectivity_scalar(m, v116, v18+int32(80), v219, v211, int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L5
	} else {
		goto L115
	}
L109:
	;
	v326 = F_calc_hist_selectivity_scalar(m, v116, v18+int32(80), v217, v211, int32(1))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L5
	} else {
		goto L114
	}
L110:
	;
	v319 = F_calc_hist_selectivity_scalar(m, v116, v18+int32(80), v217, v211, int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L5
	} else {
		goto L113
	}
L111:
	;
	v313 = F_calc_hist_selectivity_scalar(m, v116, v18+int32(80), v217, v211, int32(1))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	v421 = v313
	goto L91
L113:
	;
	v421 = base.F64_sub(float64(1), v319)
	goto L91
L114:
	;
	v421 = base.F64_sub(float64(1), v326)
	goto L91
L115:
	;
	v421 = v332
	goto L91
L116:
	;
	v421 = base.F64_sub(float64(1), v338)
	goto L91
L117:
	;
	v421 = base.F64_sub(float64(1), v345)
	goto L91
L118:
	;
	v421 = v351
	goto L91
L119:
	;
	v363 = F_calc_hist_selectivity_scalar(m, v116, v18+int32(72), v217, v211, int32(1))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L5
	} else {
		goto L120
	}
L120:
	;
	v421 = base.F64_sub(float64(1), base.F64_add(v357, base.F64_sub(float64(1), v363)))
	goto L91
L121:
	;
	v421 = v374
	goto L91
L122:
	;
	v382 = F_calc_hist_selectivity_scalar(m, v116, v18+int32(72), v219, v211, int32(1))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L5
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+76)))
	if v384 == int32(1) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v421 = v382
	goto L91
L126:
	;
	v391 = F_calc_hist_selectivity_scalar(m, v116, v18+int32(80), v217, v211, int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L5
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v18)+100))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	v400 = F_calc_hist_selectivity_contained(m, v116, v18+int32(80), v18+int32(72), v217, v211, v398, v399)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L5
	} else {
		goto L130
	}
L129:
	;
	v421 = base.F64_sub(float64(1), v391)
	goto L91
L130:
	;
	v421 = v400
	goto L91
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v73
	F_errmsg_internal(m, int32(40993), v18)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L5
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(471911), int32(579), int32(8825))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
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
	v421 = v418
	goto L91
L135:
	;
	F_free_attstatsslot(m, v18+int32(124))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L5
	} else {
		goto L136
	}
L136:
	;
	if base.F64_lt(v421, float64(0)) != 0 {
		goto L65
	} else {
		goto L137
	}
L137:
	;
	v485 = v421
	goto L64
L138:
	;
	goto L65
L139:
	;
	v475 = *(*float64)(unsafe.Add(mBase, uint32(v468<<(uint(int32(3))%32))+uint32(_consts[1110])))
	v485 = v475
	goto L64
L140:
	;
	v507 = base.F64_add(base.F64_mul(v492, v485), v152)
	goto L52
L141:
	;
	goto L142
L142:
	;
	v507 = base.F64_mul(v492, v485)
	goto L52
L143:
	;
	if base.F64_gt(v516, float64(1)) == int32(0) {
		v546 = v516
		goto L30
	} else {
		goto L144
	}
L144:
	;
	v546 = float64(1)
	goto L30
L145:
	;
	v535 = *(*float64)(unsafe.Add(mBase, uint32(v528<<(uint(int32(3))%32))+uint32(_consts[1111])))
	v546 = v535
	goto L30
L146:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	m.T0[v552].(func(*base.Module, int32))(m, v551)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L5
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v555 = float64(1)
	if base.F64_gt(v546, v555) != 0 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	goto L148
L150:
	;
	v558 = v555
	goto L152
L151:
	;
	v558 = v546
	goto L152
L152:
	;
	v579 = v558
	goto L2
L153:
	;
	m.G0 = v18 + int32(160)
	return v584
L154:
	;
	F_errmsg_internal(m, int32(467937), int32(0))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L5
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(471911), int32(257), int32(292265))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L5
	} else {
		goto L156
	}
L156:
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
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
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
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
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
	var v297 int32
	_ = v297
	v10 = m.G0
	v12 = v10 - int32(8208)
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
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[1431]))) = l0
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
	F_errstart_cold(m, int32(21), int32(525467))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L63
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(525467))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L58
	}
L5:
	;
	F_plpgsql_yyerror(m, l4, int32(0), l5, int32(201674))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L57
	}
L6:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	F_cword_is_not_variable(m, l3, v241, l5)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L56
	}
L7:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	F_word_is_not_variable(m, l3, v238, l5)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L55
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
	*(*int32)(unsafe.Add(mBase, uint32(v63+(v12+int32(4112))))) = v61
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
	v92 = F_palloc0(m, int32(40))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92)+8)) = int32(629437)
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = int32(1)
	v99 = int32(0)
	if l2 < v99 {
		v142 = v99
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92)+28)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v92)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = v142
	v150 = v80 << (uint(int32(2)) % 32)
	v151 = F_palloc(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L43
	}
L30:
	;
	goto L29
L31:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+60))
	if v105 == int32(0) {
		v142 = v99
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v108 = l2 + v105
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v104)+188))
	if base.Ui32(v109) <= base.Ui32(v108) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v104)+196))
	if v118 == int32(0) {
		v142 = v119
		goto L30
	} else {
		goto L37
	}
L34:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v104)+192))
	v118 = v111
	goto L33
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v104)+188)) = v105
	v116 = F_strchr(m, v105, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v104)+192)) = v116
	v118 = v116
	goto L33
L37:
	;
	if base.Ui32(v108) <= base.Ui32(v118) {
		v142 = v119
		goto L30
	} else {
		goto L38
	}
L38:
	;
	v123 = v118
	v125 = v119
	goto L39
L39:
	;
	v128 = int32(1)
	v129 = v125 + v128
	*(*int32)(unsafe.Add(mBase, uint32(v104)+196)) = v129
	v132 = v123 + v128
	*(*int32)(unsafe.Add(mBase, uint32(v104)+188)) = v132
	v135 = F_strchr(m, v132, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v104)+192)) = v135
	if v135 == int32(0) {
		v142 = v129
		goto L30
	} else {
		goto L41
	}
L40:
	;
	v142 = v129
	goto L30
L41:
	;
	if base.Ui32(v135) < base.Ui32(v108) {
		v123 = v135
		v125 = v129
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92)+32)) = v151
	v154 = F_palloc(m, v150)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92)+36)) = v154
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v92)+32))
	if v80&int32(1) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v161 = v80 - int32(1)
	v163 = v161 << (uint(int32(2)) % 32)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(4112)+v163)))
	*(*int32)(unsafe.Add(mBase, uint32(v157+v163))) = v168
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(16)+v163)))
	*(*int32)(unsafe.Add(mBase, uint32(v163+v154))) = v174
	v177 = v161
	goto L47
L46:
	;
	v177 = v80
	goto L47
L47:
	;
	if v80 != int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v183 = v177
	goto L51
L49:
	;
	goto L50
L50:
	;
	F_plpgsql_adddatum(m, v92)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L54
	}
L51:
	;
	v189 = int32(2)
	v192 = v183<<(uint(v189)%32) - int32(4)
	v195 = v12 + int32(4112)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v195+v192)))
	*(*int32)(unsafe.Add(mBase, uint32(v157+v192))) = v197
	v201 = v12 + int32(16)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v201+v192)))
	*(*int32)(unsafe.Add(mBase, uint32(v192+v154))) = v203
	v206 = v183 - v189
	v208 = v206 << (uint(v189) % 32)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v195+v208)))
	*(*int32)(unsafe.Add(mBase, uint32(v157+v208))) = v213
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v201+v208)))
	*(*int32)(unsafe.Add(mBase, uint32(v208+v154))) = v219
	if v189 < v183 {
		v183 = v206
		goto L51
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	goto L52
L54:
	;
	m.G0 = v12 + int32(8208)
	return v92
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
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(436550), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v263 = F_plpgsql_scanner_errposition(m, v262, l5)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(25112), int32(3674), int32(70559))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v281 = F_NameOfDatum(m, l3)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v281
	F_errmsg(m, int32(377733), v12)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v289 = F_plpgsql_scanner_errposition(m, v288, l5)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(25112), int32(3687), int32(70559))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
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
func F_rebuild_eclass_attr_needed(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v8 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v11 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = v2
	goto L4
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v16<<(uint(int32(2))%32))))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v26 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	v67 = v16 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v67 < v68 {
		v16 = v67
		goto L4
	} else {
		goto L17
	}
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v29 < int32(2) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+40)))
	if v32 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v38 = int32(0)
	goto L10
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v38<<(uint(int32(2))%32))))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v48 = F_pull_var_clause(m, v46, int32(26))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L6
L12:
	;
	return
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	F_add_vars_to_attr_needed(m, l0, v48, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_list_free(m, v48)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v56 = v38 + int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v56 < v57 {
		v38 = v56
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L11
L17:
	;
	goto L5
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v97 int64
	_ = v97
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int64
	_ = v117
	var v121 int64
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v13 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v56 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(0)
	goto L1
L3:
	;
	v16 = int32(4442992)
	v17 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v19
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v24 = m.T0[v23].(func(*base.Module, int32, int32, int32) int32)(m, v13, v12, v10+int32(15))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v17
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
	if v28 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v29
	if int64(0) <= v29 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	F_errcode(m, int32(671350914))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	F_errmsg(m, int32(326717), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(470411), int32(373), int32(115602))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L12:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v126 = v123
	goto L26
L13:
	;
	v106 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v106
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+136)) = uint8(v106)
	v112 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v114 == v112 {
		v121 = int64(-1)
		goto L12
	} else {
		goto L23
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = int32(0)
	v97 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = v97
	v99 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+136)) = uint8(v99)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v99
	v121 = int64(-1)
	goto L12
L15:
	;
	v59 = int32(4442992)
	v60 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v62
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	v67 = m.T0[v66].(func(*base.Module, int32, int32, int32) int32)(m, v56, v12, v10+int32(15))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v60
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
	if v71 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v67)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v72
	if int64(0) <= v72 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(654573698))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	F_errmsg(m, int32(326690), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(470411), int32(399), int32(115602))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v117 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v121 = v117 + v72
	goto L12
L24:
	;
	m.G0 = v10 + int32(16)
	return
L25:
	;
	goto L24
L26:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	switch v128 - int32(394) {
	case 0:
		goto L32
	default:
		goto L25
	case 3:
		goto L35
	case 4:
		goto L36
	case 17:
		goto L31
	case 32:
		goto L34
	case 33:
		goto L33
	case 38, 39:
		goto L30
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v126)+120)) = v121
	v187 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v126)+117)) = uint8(v187)
	goto L25
L28:
	;
	goto L27
L29:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v126+v183)))
	v126 = v185
	goto L26
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v126)+112)) = v121
	v183 = int32(36)
	goto L29
L31:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v126)+32))
	if v178 == int32(0) {
		v183 = int32(116)
		goto L29
	} else {
		goto L50
	}
L32:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v126)+36))
	if v176 != 0 {
		v126 = v176
		goto L26
	} else {
		goto L49
	}
L33:
	;
	if v121 < int64(0) {
		goto L46
	} else {
		goto L47
	}
L34:
	;
	if int64(0) <= v121 {
		goto L28
	} else {
		goto L45
	}
L35:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v126)+108))
	if v148 <= int32(0) {
		goto L25
	} else {
		goto L41
	}
L36:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v126)+108))
	if v131 <= int32(0) {
		goto L25
	} else {
		goto L37
	}
L37:
	;
	v137 = int32(0)
	goto L38
L38:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v126)+104))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v138+v137<<(uint(int32(2))%32))))
	F_ExecSetTupleBound(m, v121, v142)
	mBase = m.M
	v145 = v137 + int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v126)+108))
	if v145 < v146 {
		v137 = v145
		goto L38
	} else {
		goto L40
	}
L39:
	;
	goto L25
L40:
	;
	goto L39
L41:
	;
	v154 = int32(0)
	goto L42
L42:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v126)+104))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155+v154<<(uint(int32(2))%32))))
	F_ExecSetTupleBound(m, v121, v159)
	mBase = m.M
	v162 = v154 + int32(1)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v126)+108))
	if v162 < v163 {
		v154 = v162
		goto L42
	} else {
		goto L44
	}
L43:
	;
	goto L25
L44:
	;
	goto L43
L45:
	;
	v167 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v126)+117)) = uint8(v167)
	goto L24
L46:
	;
	v171 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v126)+116)) = uint8(v171)
	goto L24
L47:
	;
	goto L48
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v126)+120)) = v121
	v174 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v126)+116)) = uint8(v174)
	goto L24
L49:
	;
	goto L25
L50:
	;
	goto L25
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v88 int32
	_ = v88
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
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
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
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
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
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
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
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
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
	F_op_input_types(m, v24, v18+int32(12), v18+int32(8))
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
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v63 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L4:
	;
	v36 = v18 + int32(8)
	v37 = int32(48)
	if v34 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v47 = v18 + int32(12)
	v48 = int32(44)
	if v34 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v58 = v36
	v60 = v37
	v61 = v4
	v62 = v4
	goto L3
L8:
	;
	goto L9
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v42 < int32(2) {
		v58 = v36
		v60 = v37
		v61 = v41
		v62 = v4
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v58 = v36
	v60 = v37
	v61 = v41
	v62 = v45
	goto L3
L11:
	;
	v58 = v47
	v60 = v48
	v61 = v4
	v62 = v4
	goto L3
L12:
	;
	goto L13
L13:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if int32(2) <= v52 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v56 = v55
	goto L16
L15:
	;
	v56 = v4
	goto L16
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v58 = v47
	v60 = v48
	v61 = v56
	v62 = v57
	goto L3
L17:
	;
	m.G0 = v18 + int32(16)
	return v390 & int32(1)
L18:
	;
	v390 = v4
	goto L17
L19:
	;
	goto L20
L20:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v66 <= int32(0) {
		v390 = v4
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v60+v21)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v74 = int32(0)
	goto L22
L22:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v74<<(uint(int32(2))%32))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+40)))
	if v93 != int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v390 = v4
	goto L17
L24:
	;
	v378 = v74 + int32(1)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v378 < v379 {
		v74 = v378
		goto L22
	} else {
		goto L91
	}
L25:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+41)))
	if v96 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
	if v23 != v97 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v21)+96))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v101 = F_equal(m, v99, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v101 == int32(0) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
	if v105 == int32(0) {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v108 = int32(0)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v109 <= v108 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	v116 = v108
	goto L32
L32:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127+v116<<(uint(int32(2))%32))))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v133 = F_equal(m, v61, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
	if v141 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	if v133 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v138 = v116 + int32(1)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v138 < v139 {
		v116 = v138
		goto L32
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	goto L33
L38:
	;
	goto L24
L39:
	;
	v390 = v4
	goto L17
L40:
	;
	goto L41
L41:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if v144 <= int32(0) {
		v390 = v4
		goto L17
	} else {
		goto L42
	}
L42:
	;
	v153 = int32(0)
	v157 = v4
	goto L43
L43:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v163+v153<<(uint(int32(2))%32))))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+12)))
	if v168 != int32(1) {
		v324 = v157
		goto L46
	} else {
		goto L47
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L88
	}
L45:
	;
	goto L44
L46:
	;
	v331 = v153 + int32(1)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if v331 < v332 {
		v153 = v331
		v157 = v324
		goto L43
	} else {
		goto L87
	}
L47:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v171 == int32(0) {
		v324 = v157
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	if v174 <= int32(0) {
		v324 = v157
		goto L46
	} else {
		goto L49
	}
L49:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v180 = int32(0)
	goto L50
L50:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v194+v180<<(uint(int32(2))%32))))
	v200 = F_get_opfamily_member_for_cmptype(m, v198, v71, v177, int32(3))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v215 = F_bms_copy(m, v70)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L62
	}
L52:
	;
	goto L51
L53:
	;
	if v200 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v92)+52))
	if v202 == int32(0) {
		goto L52
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v210 = v180 + int32(1)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	if v210 < v211 {
		v180 = v210
		goto L50
	} else {
		goto L61
	}
L57:
	;
	v205 = F_get_opcode(m, v200)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v207 = F_get_func_leakproof(m, v205)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	if v207 != 0 {
		goto L52
	} else {
		goto L60
	}
L60:
	;
	goto L56
L61:
	;
	v324 = v157
	goto L46
L62:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v92)+48))
	v218 = F_build_implied_join_equality(m, l0, v200, v213, v62, v214, v215, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v218
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v221 == int32(0) {
		goto L45
	} else {
		goto L64
	}
L64:
	;
	v224 = int32(0)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if v225 <= v224 {
		goto L45
	} else {
		goto L65
	}
L65:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v230 = v224
	goto L66
L66:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v244+v230<<(uint(int32(2))%32))))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	v250 = int32(0)
	if v249 == v250 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v312 = F_process_equivalence(m, l0, v18+int32(4), v248)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L86
	}
L68:
	;
	if v303 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L69:
	;
	v303 = int32(1)
	goto L68
L70:
	;
	goto L71
L71:
	;
	if v228 == int32(0) {
		v294 = v250
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v303 = v294
	goto L68
L73:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	if v260 < v259 {
		v294 = v250
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v262 = int32(1)
	if v259 <= v262 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v265 = v262
	goto L77
L76:
	;
	v265 = v259
	goto L77
L77:
	;
	v266 = int32(8)
	v271 = int32(0)
	goto L78
L78:
	;
	v278 = v271 << (uint(int32(2)) % 32)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v249+v266+v278)))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v278+(v228+v266))))
	v285 = v280 & (v282 ^ int32(-1))
	v287 = base.B2i32(v285 == int32(0))
	if v285 != 0 {
		v294 = v287
		goto L72
	} else {
		goto L80
	}
L79:
	;
	v294 = v287
	goto L72
L80:
	;
	v289 = v271 + int32(1)
	if v289 != v265 {
		v271 = v289
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v307 = v230 + int32(1)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if v307 < v308 {
		v230 = v307
		goto L66
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	goto L67
L85:
	;
	goto L45
L86:
	;
	v324 = v312 | v157
	goto L46
L87:
	;
	v390 = v324
	goto L17
L88:
	;
	F_errmsg_internal(m, int32(264571), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(470894), int32(2628), int32(264190))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
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
	goto L23
}
func F_record_image_gt(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_image_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(int32(0) < v2)
	}
}
func F_record_image_le(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_image_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v2 <= int32(0))
	}
}
func F_record_image_ne(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_image_eq(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2 ^ int32(1)
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
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	F_pq_startmsgread(m)
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
	v13 = F_pq_getbyte(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L49
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L45
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L41
	}
L6:
	;
	m.G0 = v7 + int32(32)
	return v108
L7:
	;
	if v13 == int32(-1) {
		v108 = v1
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if v13 != int32(112) {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	F_initStringInfo(m, v7+int32(16))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v26 = F_pq_getmessage(m, v7+int32(16), int32(65535))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v26 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_pfree(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v28&int32(3) == int32(0) {
		v54 = v28
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v108 = v1
	goto L6
L16:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	if v87+int32(1) != v90 {
		goto L4
	} else {
		goto L33
	}
L17:
	;
	v87 = v79 - v28
	goto L16
L18:
	;
	v58 = v54
	goto L27
L19:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v38 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v87 = int32(0)
	goto L16
L21:
	;
	goto L22
L22:
	;
	v43 = v28
	goto L23
L23:
	;
	v47 = v43 + int32(1)
	if v47&int32(3) == int32(0) {
		v54 = v47
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v79 = v47
	goto L17
L25:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v52 != 0 {
		v43 = v47
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v67 = int32(-2139062144)
	if (int32(16843008)-v64|v64)&v67 == v67 {
		v58 = v58 + int32(4)
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v73 = v58
	goto L30
L29:
	;
	goto L28
L30:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v77 != 0 {
		v73 = v73 + int32(1)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v79 = v73
	goto L17
L32:
	;
	goto L31
L33:
	;
	if v87 == int32(0) {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v96 = F_errstart(m, int32(10), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v96 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	F_errmsg_internal(m, int32(100386), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v108 = v107
	goto L6
L39:
	;
	F_errfinish(m, int32(474729), int32(768), int32(100141))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v13
	F_errmsg(m, int32(454317), v7)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(474729), int32(727), int32(100141))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(324664), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(474729), int32(747), int32(100141))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
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
	F_errcode(m, int32(16908802))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(90550), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(474729), int32(765), int32(100141))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regexeqsel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v10 float64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = F_patternsel_common(m, v2, v3, v4, v5, v6, v7, int32(2), v4)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_Float8GetDatum(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v14
		}
	}
}
func F_regnamespaceout(m *base.Module, l0 int32) int32 {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v8 == int32(0) {
		v12 = F_pstrdup(m, int32(628667))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v30 = v12
			m.G0 = v6 + int32(16)
			return v30
		}
	} else {
		v16 = F_get_namespace_name(m, v8)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v16 != 0 {
				v18 = F_quote_identifier(m, v16)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = F_pstrdup(m, v18)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v30 = v20
						m.G0 = v6 + int32(16)
						return v30
					}
				}
			} else {
				v23 = F_palloc(m, int32(64))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
					v28 = F_pg_snprintf(m, v23, int32(64), int32(56934), v6)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = v23
						m.G0 = v6 + int32(16)
						return v30
					}
				}
			}
		}
	}
}
func F_regtypein(m *base.Module, l0 int32) int32 {
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
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v11 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L48
	} else {
		goto L52
	}
L2:
	;
	m.G0 = v7 + int32(16)
	return v186
L3:
	;
	v173 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v173
	v176 = *(*int32)(unsafe.Add(mBase, _consts[231]))
	if v176 == v173 {
		goto L1
	} else {
		goto L50
	}
L4:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v14 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v11-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v186 = int32(0)
	goto L2
L8:
	;
	v22 = int32(523654)
	v26 = m.G0
	v28 = v26 - int32(32)
	v29 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+24)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v28)+16)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = v29
	v37 = int32(*(*uint8)(unsafe.Add(mBase, _consts[180])))
	if v37 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v10&int32(3) == int32(0) {
		v129 = v10
		goto L32
	} else {
		goto L33
	}
L10:
	;
	v105 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, _consts[181])))
	if v41 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v45 = v10
	goto L16
L14:
	;
	goto L15
L15:
	;
	v55 = v22
	v56 = v37
	goto L19
L16:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v51 == v37 {
		v45 = v45 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v105 = v45 - v10
	goto L9
L18:
	;
	goto L17
L19:
	;
	v63 = v28 + int32(base.Ui32(v56)>>(uint(int32(3))%32))&int32(28)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v65 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v64 | v65<<(uint(v56)%32)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v69 != 0 {
		v55 = v55 + v65
		v56 = v69
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v72 == int32(0) {
		v97 = v10
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v105 = v97 - v10
	goto L9
L23:
	;
	v76 = v10
	v77 = v72
	goto L24
L24:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(base.Ui32(v77)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v85)>>(uint(v77)%32))&int32(1) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v97 = v93
	goto L22
L26:
	;
	v97 = v76
	goto L22
L27:
	;
	goto L28
L28:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	v93 = v76 + int32(1)
	if v91 != 0 {
		v76 = v93
		v77 = v91
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	if v105 != v162 {
		goto L3
	} else {
		goto L47
	}
L31:
	;
	v162 = v154 - v10
	goto L30
L32:
	;
	v133 = v129
	goto L41
L33:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v113 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v162 = int32(0)
	goto L30
L35:
	;
	goto L36
L36:
	;
	v118 = v10
	goto L37
L37:
	;
	v122 = v118 + int32(1)
	if v122&int32(3) == int32(0) {
		v129 = v122
		goto L32
	} else {
		goto L39
	}
L38:
	;
	v154 = v122
	goto L31
L39:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	if v127 != 0 {
		v118 = v122
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v142 = int32(-2139062144)
	if (int32(16843008)-v139|v139)&v142 == v142 {
		v133 = v133 + int32(4)
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v148 = v133
	goto L44
L43:
	;
	goto L42
L44:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v152 != 0 {
		v148 = v148 + int32(1)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v154 = v148
	goto L31
L46:
	;
	goto L45
L47:
	;
	v168 = F_DirectInputFunctionCallSafe(m, int32(547), v10, int32(-1), v9, v7+int32(12))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	return int32(0)
L49:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v186 = v172
	goto L2
L50:
	;
	v183 = F_parseTypeString(m, v10, v7+int32(8), v7+int32(4), v9)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v186 = v185
	goto L2
L52:
	;
	F_errmsg_internal(m, int32(392616), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L48
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(476730), int32(1191), int32(263550))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L48
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_relmap_desc(m *base.Module, l0 int32, l1 int32) {
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
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+48)))
	if base.Ui32(v10) <= base.Ui32(int32(15)) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v15
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = v14
		F_appendStringInfo(m, l0, int32(452855), v7)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	} else {
		m.G0 = v7 + int32(16)
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
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
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
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
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
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
	if v8 != int32(319) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v159 = F_expression_tree_mutator_impl(m, l0, int32(1052), l1)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L12
	} else {
		goto L52
	}
L5:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v145 + int32(1)
	v151 = F_query_tree_mutator_impl(m, l0, int32(1052), l1, int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L12
	} else {
		goto L51
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
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v81 != v82 {
		goto L4
	} else {
		goto L32
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
	if v24 == v26 {
		v67 = v26
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v67 == int32(0) {
		goto L4
	} else {
		goto L29
	}
L16:
	;
	goto L15
L17:
	;
	if v25 == int32(0) {
		v67 = v26
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v35 < v36 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v38 = v35
	goto L21
L20:
	;
	v38 = v36
	goto L21
L21:
	;
	if v38 <= int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v41 = int32(1)
	goto L24
L23:
	;
	v41 = v38
	goto L24
L24:
	;
	v42 = int32(8)
	v47 = int32(0)
	goto L25
L25:
	;
	v54 = v47 << (uint(int32(2)) % 32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v25+v42+v54)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54+(v24+v42))))
	v59 = v56 & v58
	v61 = base.B2i32(v59 != int32(0))
	if v59 != 0 {
		v67 = v61
		goto L16
	} else {
		goto L27
	}
L26:
	;
	v67 = v61
	goto L16
L27:
	;
	v63 = v47 + int32(1)
	if v63 != v41 {
		v47 = v63
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v73 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L12
	} else {
		goto L30
	}
L30:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+24))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v77 = F_bms_difference(m, v75, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = v77
	return v73
L32:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v86 = int32(0)
	if v84 == v86 {
		v127 = v86
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v127 != 0 {
		goto L4
	} else {
		goto L47
	}
L34:
	;
	goto L33
L35:
	;
	if v85 == int32(0) {
		v127 = v86
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v95 < v96 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v98 = v95
	goto L39
L38:
	;
	v98 = v96
	goto L39
L39:
	;
	if v98 <= int32(1) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v101 = int32(1)
	goto L42
L41:
	;
	v101 = v98
	goto L42
L42:
	;
	v102 = int32(8)
	v107 = int32(0)
	goto L43
L43:
	;
	v114 = v107 << (uint(int32(2)) % 32)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v85+v102+v114)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114+(v84+v102))))
	v119 = v116 & v118
	v121 = base.B2i32(v119 != int32(0))
	if v119 != 0 {
		v127 = v121
		goto L34
	} else {
		goto L45
	}
L44:
	;
	v127 = v121
	goto L34
L45:
	;
	v123 = v107 + int32(1)
	if v123 != v101 {
		v107 = v123
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v132 = F_expression_tree_mutator_impl(m, l0, int32(1052), l1)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v136 = F_bms_difference(m, v134, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132)+12)) = v136
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v141 = F_bms_difference(m, v139, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = v141
	return v132
L51:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v153 - int32(1)
	return v151
L52:
	;
	return v159
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
	var v41 int32
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
	v25 = l1 + v19<<(uint(int32(4))%32)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v28 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v26 - v28
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v27 + v28
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v37 = int32(*(*int16)(unsafe.Add(mBase, uint32(v36)+10)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v41 = F_heap_getattr_1(m, v11+int32(12), v37, v38, v25+v28)
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
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v41
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
	var v23 int32
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
	v16 = l1 + v12<<(uint(int32(4))%32)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v23 = F_index_getattr_2(m, v17, int32(1), v20, v16+int32(8))
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v23
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
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
				v24 = int32(4)
				v26 = l1 + v17<<(uint(v24)%32)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v28
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v31
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v34
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v37
				v40 = v17 + v24
				v42 = v22 + v24
				if v42 != l2&int32(2147483644) {
					v17 = v40
					v22 = v42
					continue
				} else {
					break
				}
				break
			}
			v44 = v40
		} else {
			v44 = v12
		}
		if v11 == int32(0) {
		} else {
			v53 = v44
			v57 = v4
			for {
				v62 = l1 + v53<<(uint(int32(4))%32)
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v64
				v66 = int32(1)
				v69 = v57 + v66
				if v69 != v11 {
					v53 = v53 + v66
					v57 = v69
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
	var v41 int32
	_ = v41
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
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	return v51
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
	return (v26 ^ int32(1)) & int32(255)
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
	v41 = v18 << (uint(int32(2)) % 32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v16+v41)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v14)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+196))
	v49 = v46 + v18*int32(36)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	v51 = m.T0[v50].(func(*base.Module, int32, int32, int32) int32)(m, v43, v45, v49)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	if v51 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v56 = v18 + int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)+136))
	if v56 < v57 {
		v18 = v56
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
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
	var v62 int32
	_ = v62
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
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
	var v99 int32
	_ = v99
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
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
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
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
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
	var v162 int32
	_ = v162
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
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
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
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
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
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
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
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 float64
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
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
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v491 int32
	_ = v491
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v10 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v491
L2:
	;
	v491 = l1
	goto L1
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+228))
	v15 = int32(0)
	if v13 == v15 {
		v56 = v15
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v56 == int32(0) {
		goto L2
	} else {
		goto L18
	}
L5:
	;
	goto L4
L6:
	;
	if v14 == int32(0) {
		v56 = v15
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v24 < v25 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v27 = v24
	goto L10
L9:
	;
	v27 = v25
	goto L10
L10:
	;
	if v27 <= int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v30 = int32(1)
	goto L13
L12:
	;
	v30 = v27
	goto L13
L13:
	;
	v31 = int32(8)
	v36 = int32(0)
	goto L14
L14:
	;
	v43 = v36 << (uint(int32(2)) % 32)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v14+v31+v43)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+(v13+v31))))
	v48 = v45 & v47
	v50 = base.B2i32(v48 != int32(0))
	if v48 != 0 {
		v56 = v50
		goto L5
	} else {
		goto L16
	}
L15:
	;
	v56 = v50
	goto L5
L16:
	;
	v52 = v36 + int32(1)
	if v52 != v30 {
		v36 = v52
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v62 = int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v63 - int32(279) {
	case 0:
		goto L33
	case 1:
		goto L32
	default:
		v491 = v62
		goto L1
	case 3:
		goto L31
	case 4:
		goto L30
	case 5:
		goto L29
	case 9:
		goto L28
	case 10:
		goto L27
	case 11:
		goto L23
	case 14:
		goto L22
	case 15:
		goto L21
	case 17:
		goto L20
	case 19:
		goto L26
	case 20:
		goto L25
	case 21:
		goto L24
	}
L19:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	v284 = F_adjust_child_relids_multilevel(m, l0, v282, l2, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L34
	} else {
		goto L106
	}
L20:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v274 = F_reparameterize_path_by_child(m, l0, v273, l2)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L34
	} else {
		goto L103
	}
L21:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v263 = F_reparameterize_path_by_child(m, l0, v262, l2)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L34
	} else {
		goto L100
	}
L22:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v259 = F_reparameterize_path_by_child(m, l0, v258, l2)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L34
	} else {
		goto L98
	}
L23:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v252 == int32(0) {
		goto L19
	} else {
		goto L95
	}
L24:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v231 = F_reparameterize_path_by_child(m, l0, v230, l2)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L34
	} else {
		goto L89
	}
L25:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v209 = F_reparameterize_path_by_child(m, l0, v208, l2)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L34
	} else {
		goto L83
	}
L26:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v192 = F_reparameterize_path_by_child(m, l0, v191, l2)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L34
	} else {
		goto L78
	}
L27:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+184))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	v165 = F_adjust_appendrel_attrs_multilevel(m, l0, v163, l2, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L34
	} else {
		goto L65
	}
L28:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+184))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	v137 = F_adjust_appendrel_attrs_multilevel(m, l0, v135, l2, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L34
	} else {
		goto L53
	}
L29:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v128 == int32(0) {
		goto L19
	} else {
		goto L50
	}
L30:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v122 == int32(0) {
		goto L19
	} else {
		goto L47
	}
L31:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+184))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	v114 = F_adjust_appendrel_attrs_multilevel(m, l0, v112, l2, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L34
	} else {
		goto L44
	}
L32:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+96))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	v102 = F_adjust_appendrel_attrs_multilevel(m, l0, v100, l2, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L34
	} else {
		goto L42
	}
L33:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+184))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	v69 = F_adjust_appendrel_attrs_multilevel(m, l0, v67, l2, v68)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	return int32(0)
L35:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+184)) = v69
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v75 != int32(340) {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+68))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v80 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+32))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	v96 = F_adjust_appendrel_attrs_multilevel(m, l0, v94, l2, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L34
	} else {
		goto L41
	}
L38:
	;
	v92 = v80 + v79<<(uint(int32(2))%32)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+52))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v92 = v86 + v79<<(uint(int32(2))%32) - int32(4)
	goto L37
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+32)) = v96
	goto L19
L42:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v104)+96)) = v102
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	v108 = F_adjust_appendrel_attrs_multilevel(m, l0, v106, l2, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L34
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v108
	goto L19
L44:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+184)) = v114
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v119 = F_reparameterize_path_by_child(m, l0, v118, l2)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L34
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v119
	if v119 != 0 {
		goto L19
	} else {
		goto L46
	}
L46:
	;
	v491 = v62
	goto L1
L47:
	;
	v125 = F_reparameterize_pathlist_by_child(m, l0, v122, l2)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L34
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v125
	if v125 != 0 {
		goto L19
	} else {
		goto L49
	}
L49:
	;
	v491 = v62
	goto L1
L50:
	;
	v131 = F_reparameterize_pathlist_by_child(m, l0, v128, l2)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L34
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v131
	if v131 != 0 {
		goto L19
	} else {
		goto L52
	}
L52:
	;
	v491 = v62
	goto L1
L53:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v139)+184)) = v137
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v141 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v142 = F_reparameterize_path_by_child(m, l0, v141, l2)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L34
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v148 != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v142
	if v142 == int32(0) {
		v491 = v62
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	v150 = F_adjust_appendrel_attrs_multilevel(m, l0, v148, l2, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L34
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+168))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+164))
	if v155 == int32(0) {
		goto L19
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v150
	goto L61
L63:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v159 = m.T0[v155].(func(*base.Module, int32, int32, int32) int32)(m, l0, v158, l2)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L34
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v159
	goto L19
L65:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v167)+184)) = v165
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v169 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v170 = F_reparameterize_pathlist_by_child(m, l0, v169, l2)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L34
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v176 != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v170
	if v170 == int32(0) {
		v491 = v62
		goto L1
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	v178 = F_adjust_appendrel_attrs_multilevel(m, l0, v176, l2, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L34
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	if v181 == int32(0) {
		goto L19
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v178
	goto L73
L75:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v181)+8))
	if v184 == int32(0) {
		goto L19
	} else {
		goto L76
	}
L76:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v188 = m.T0[v184].(func(*base.Module, int32, int32, int32) int32)(m, l0, v187, l2)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L34
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+84)) = v188
	goto L19
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v192
	if v192 == int32(0) {
		v491 = v62
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v198 = F_reparameterize_path_by_child(m, l0, v197, l2)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L34
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+84)) = v198
	if v198 == int32(0) {
		v491 = v62
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	v205 = F_adjust_appendrel_attrs_multilevel(m, l0, v203, l2, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L34
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v205
	goto L19
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v209
	if v209 == int32(0) {
		v491 = v62
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v215 = F_reparameterize_path_by_child(m, l0, v214, l2)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L34
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+84)) = v215
	if v215 == int32(0) {
		v491 = v62
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	v222 = F_adjust_appendrel_attrs_multilevel(m, l0, v220, l2, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L34
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v222
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	v227 = F_adjust_appendrel_attrs_multilevel(m, l0, v225, l2, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L34
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v227
	goto L19
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v231
	if v231 == int32(0) {
		v491 = v62
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v237 = F_reparameterize_path_by_child(m, l0, v236, l2)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L34
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+84)) = v237
	if v237 == int32(0) {
		v491 = v62
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	v244 = F_adjust_appendrel_attrs_multilevel(m, l0, v242, l2, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L34
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v244
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	v249 = F_adjust_appendrel_attrs_multilevel(m, l0, v247, l2, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L34
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v249
	goto L19
L95:
	;
	v255 = F_reparameterize_pathlist_by_child(m, l0, v252, l2)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L34
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v255
	if v255 != 0 {
		goto L19
	} else {
		goto L97
	}
L97:
	;
	v491 = v62
	goto L1
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v259
	if v259 != 0 {
		goto L19
	} else {
		goto L99
	}
L99:
	;
	v491 = v62
	goto L1
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v263
	if v263 == int32(0) {
		v491 = v62
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	v270 = F_adjust_appendrel_attrs_multilevel(m, l0, v268, l2, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L34
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v270
	goto L19
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v274
	if v274 == int32(0) {
		v491 = v62
		goto L1
	} else {
		goto L104
	}
L104:
	;
	goto L19
L105:
	;
	if v375 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L106:
	;
	v286 = int32(0)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+36))
	if v288 == v286 {
		v375 = v286
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v291 = int32(0)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	if v292 <= v291 {
		v375 = v291
		goto L105
	} else {
		goto L108
	}
L108:
	;
	v300 = int32(0)
	goto L110
L109:
	;
	v375 = v309
	goto L105
L110:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v288)+12))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v305+v300<<(uint(int32(2))%32))))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	v311 = int32(0)
	v318 = base.B2i32(v310|v284 == v311)
	if v310 == v311 {
		v357 = v318
		goto L113
	} else {
		goto L114
	}
L111:
	;
	v375 = int32(0)
	goto L105
L112:
	;
	if v357 != 0 {
		goto L109
	} else {
		goto L124
	}
L113:
	;
	goto L112
L114:
	;
	if v284 == int32(0) {
		v357 = v318
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v310)+4))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	if v324 != v325 {
		v357 = int32(0)
		goto L113
	} else {
		goto L116
	}
L116:
	;
	v327 = int32(1)
	if v324 <= v327 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v330 = v327
	goto L119
L118:
	;
	v330 = v324
	goto L119
L119:
	;
	v331 = int32(8)
	v336 = int32(0)
	goto L120
L120:
	;
	v344 = v336 << (uint(int32(2)) % 32)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v310+v331+v344)))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v344+(v284+v331))))
	v349 = base.B2i32(v346 == v348)
	if v348 != v346 {
		v357 = v349
		goto L113
	} else {
		goto L122
	}
L121:
	;
	v357 = v349
	goto L113
L122:
	;
	v352 = v336 + int32(1)
	if v352 != v330 {
		v336 = v352
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	v362 = v300 + int32(1)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	if v362 < v363 {
		v300 = v362
		goto L110
	} else {
		goto L125
	}
L125:
	;
	goto L111
L126:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v379 = F_GetMemoryChunkContext(m, v378)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L34
	} else {
		goto L129
	}
L127:
	;
	v411 = v375
	goto L128
L128:
	;
	F_bms_free(m, v284)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L34
	} else {
		goto L135
	}
L129:
	;
	v381 = int32(4442992)
	v382 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v379
	v386 = F_palloc0(m, int32(24))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L34
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386))) = int32(278)
	v390 = F_bms_copy(m, v284)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L34
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+4)) = v390
	v393 = *(*float64)(unsafe.Add(mBase, uint32(v281)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v386)+8)) = v393
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v281)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v386)+16)) = v395
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	v398 = F_adjust_appendrel_attrs_multilevel(m, l0, v395, l2, v397)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L34
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+16)) = v398
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v281)+20))
	v402 = F_bms_copy(m, v401)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L34
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+20)) = v402
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v378)+36))
	v406 = F_lappend(m, v405, v386)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L34
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v378)+36)) = v406
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v382
	v411 = v386
	goto L128
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v411
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)+64))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l2)+228))
	v421 = int32(0)
	if v419 == v421 {
		v462 = v421
		goto L137
	} else {
		goto L138
	}
L136:
	;
	if v462 == int32(0) {
		goto L2
	} else {
		goto L150
	}
L137:
	;
	goto L136
L138:
	;
	if v420 == int32(0) {
		v462 = v421
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v419)+4))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	if v430 < v431 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v433 = v430
	goto L142
L141:
	;
	v433 = v431
	goto L142
L142:
	;
	if v433 <= int32(1) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v436 = int32(1)
	goto L145
L144:
	;
	v436 = v433
	goto L145
L145:
	;
	v437 = int32(8)
	v442 = int32(0)
	goto L146
L146:
	;
	v449 = v442 << (uint(int32(2)) % 32)
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v420+v437+v449)))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v449+(v419+v437))))
	v454 = v451 & v453
	v456 = base.B2i32(v454 != int32(0))
	if v454 != 0 {
		v462 = v456
		goto L137
	} else {
		goto L148
	}
L147:
	;
	v462 = v456
	goto L137
L148:
	;
	v458 = v442 + int32(1)
	if v458 != v436 {
		v442 = v458
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v469 = F_copy_pathtarget(m, v468)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L34
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v469
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v469)+4))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	v474 = F_adjust_appendrel_attrs_multilevel(m, l0, v472, l2, v473)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L34
	} else {
		goto L152
	}
L152:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v476)+4)) = v474
	goto L2
}
func F_reparameterize_pathlist_by_child(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v4 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v7 <= v4 {
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
	v15 = v4
	v16 = v4
	goto L4
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v16<<(uint(int32(2))%32))))
	v23 = F_reparameterize_path_by_child(m, l0, v22, l2)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return v33
L6:
	;
	return int32(0)
L7:
	;
	if v23 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_list_free(m, v15)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v33 = F_lappend(m, v15, v23)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v36 = v16 + int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v36 < v37 {
		v15 = v33
		v16 = v36
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
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
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
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
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
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
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
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
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
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v486 int32
	_ = v486
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	v10 = l2
	v11 = l3
	v12 = l4
	goto L2
L1:
	;
	return
L2:
	;
	v16 = int32(2)
	if v16 <= v12 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v577 != 0 {
		goto L266
	} else {
		goto L267
	}
L4:
	;
	v19 = v16
	goto L6
L5:
	;
	v19 = v12
	goto L6
L6:
	;
	if v12 == int32(256) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v22 = int32(3)
	goto L9
L8:
	;
	v22 = v19
	goto L9
L9:
	;
	v25 = v10
	v26 = v11
	goto L10
L10:
	;
	v31 = int32(2)
	if v31 <= v26 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L3
L12:
	;
	goto L11
L13:
	;
	v34 = v31
	goto L15
L14:
	;
	v34 = v26
	goto L15
L15:
	;
	if v26 == int32(256) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v39 = int32(12)
	goto L18
L17:
	;
	v39 = v34 << (uint(int32(2)) % 32)
	goto L18
L18:
	;
	v40 = v39 + v22
	if v40 != int32(11) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	switch v40 {
	case 0:
		goto L28
	case 1:
		goto L27
	case 2:
		goto L26
	case 3:
		goto L25
	default:
		goto L12
	case 5:
		goto L1
	case 6:
		goto L24
	case 7:
		goto L23
	case 10:
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v563 = F_newstate(m, v562)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L29
	} else {
		goto L261
	}
L22:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v548 = F_newstate(m, v547)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L29
	} else {
		goto L256
	}
L23:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v362 = F_newstate(m, v361)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L29
	} else {
		goto L174
	}
L24:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v342 = F_newstate(m, v341)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L29
	} else {
		goto L166
	}
L25:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v219 = F_newstate(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L29
	} else {
		goto L111
	}
L26:
	;
	F_repeat_1(m, l0, l1, v25, int32(1), v12)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L29
	} else {
		goto L84
	}
L27:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v110 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v110 != 0 {
		goto L59
	} else {
		goto L60
	}
L28:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v25
	F_deltraverse(m, v43, l1)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return
L30:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+76))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if v48 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v51
	goto L33
L32:
	;
	goto L33
L33:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v57 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v57 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L29
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v60 <= v61 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L36
L38:
	;
	F_createarc(m, v55, int32(110), int32(0), l1, v25)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L29
	} else {
		goto L58
	}
L39:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v63 == int32(0) {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v80 == int32(0) {
		goto L38
	} else {
		goto L50
	}
L42:
	;
	v69 = v63
	goto L43
L43:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	if v73 != v25 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L38
L45:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	if v79 != 0 {
		v69 = v79
		goto L43
	} else {
		goto L49
	}
L46:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
	if v75 != 0 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v76 == int32(110) {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	goto L44
L50:
	;
	v86 = v80
	goto L51
L51:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	if v90 != l1 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L38
L53:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v86)+24))
	if v96 != 0 {
		v86 = v96
		goto L51
	} else {
		goto L57
	}
L54:
	;
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+4)))
	if v92 != 0 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if v93 == int32(110) {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	goto L52
L58:
	;
	return
L59:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L29
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v113 <= v114 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L61
L63:
	;
	F_createarc(m, v108, int32(110), int32(0), l1, v25)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L29
	} else {
		goto L83
	}
L64:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v116 == int32(0) {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v133 == int32(0) {
		goto L63
	} else {
		goto L75
	}
L67:
	;
	v122 = v116
	goto L68
L68:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	if v126 != v25 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L63
L70:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	if v132 != 0 {
		v122 = v132
		goto L68
	} else {
		goto L74
	}
L71:
	;
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+4)))
	if v128 != 0 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v129 == int32(110) {
		goto L1
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
	v139 = v133
	goto L76
L76:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139)+8))
	if v143 != l1 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L63
L78:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v139)+24))
	if v149 != 0 {
		v139 = v149
		goto L76
	} else {
		goto L82
	}
L79:
	;
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+4)))
	if v145 != 0 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if v146 == int32(110) {
		goto L1
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
	return
L84:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v164 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v167 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v167 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L29
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v170 <= v171 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L88
L90:
	;
	F_createarc(m, v165, int32(110), int32(0), l1, v25)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L29
	} else {
		goto L110
	}
L91:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v173 == int32(0) {
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v190 == int32(0) {
		goto L90
	} else {
		goto L102
	}
L94:
	;
	v179 = v173
	goto L95
L95:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
	if v183 != v25 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L90
L97:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
	if v189 != 0 {
		v179 = v189
		goto L95
	} else {
		goto L101
	}
L98:
	;
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v179)+4)))
	if v185 != 0 {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	if v186 == int32(110) {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	goto L97
L101:
	;
	goto L96
L102:
	;
	v196 = v190
	goto L103
L103:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	if v200 != l1 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	goto L90
L105:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v196)+24))
	if v206 != 0 {
		v196 = v206
		goto L103
	} else {
		goto L109
	}
L106:
	;
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v196)+4)))
	if v202 != 0 {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	if v203 == int32(110) {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	goto L105
L109:
	;
	goto L104
L110:
	;
	return
L111:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v221 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_moveouts(m, v222, l1, v219)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L29
	} else {
		goto L113
	}
L113:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_moveins(m, v225, v25, v219)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L29
	} else {
		goto L114
	}
L114:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v230 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v230 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L29
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v219)+8))
	if v233 <= v234 {
		goto L121
	} else {
		goto L122
	}
L118:
	;
	goto L117
L119:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v290 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v290 != 0 {
		goto L141
	} else {
		goto L142
	}
L120:
	;
	F_createarc(m, v228, int32(110), int32(0), l1, v219)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L29
	} else {
		goto L140
	}
L121:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v236 == int32(0) {
		goto L120
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	if v253 == int32(0) {
		goto L120
	} else {
		goto L132
	}
L124:
	;
	v242 = v236
	goto L125
L125:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
	if v246 != v219 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	goto L120
L127:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v242)+16))
	if v252 != 0 {
		v242 = v252
		goto L125
	} else {
		goto L131
	}
L128:
	;
	v248 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v242)+4)))
	if v248 != 0 {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	if v249 == int32(110) {
		goto L119
	} else {
		goto L130
	}
L130:
	;
	goto L127
L131:
	;
	goto L126
L132:
	;
	v259 = v253
	goto L133
L133:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v259)+8))
	if v263 != l1 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	goto L120
L135:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v259)+24))
	if v269 != 0 {
		v259 = v269
		goto L133
	} else {
		goto L139
	}
L136:
	;
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v259)+4)))
	if v265 != 0 {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	if v266 == int32(110) {
		goto L119
	} else {
		goto L138
	}
L138:
	;
	goto L135
L139:
	;
	goto L134
L140:
	;
	goto L119
L141:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L29
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v293 <= v294 {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	goto L143
L145:
	;
	F_createarc(m, v288, int32(110), int32(0), v219, v25)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L29
	} else {
		goto L165
	}
L146:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
	if v296 == int32(0) {
		goto L145
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v313 == int32(0) {
		goto L145
	} else {
		goto L157
	}
L149:
	;
	v302 = v296
	goto L150
L150:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v302)+12))
	if v306 != v25 {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	goto L145
L152:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v302)+16))
	if v312 != 0 {
		v302 = v312
		goto L150
	} else {
		goto L156
	}
L153:
	;
	v308 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v302)+4)))
	if v308 != 0 {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	if v309 == int32(110) {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	goto L152
L156:
	;
	goto L151
L157:
	;
	v319 = v313
	goto L158
L158:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v319)+8))
	if v323 != v219 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	goto L145
L160:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v319)+24))
	if v329 != 0 {
		v319 = v329
		goto L158
	} else {
		goto L164
	}
L161:
	;
	v325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v319)+4)))
	if v325 != 0 {
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	if v326 == int32(110) {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	goto L160
L164:
	;
	goto L159
L165:
	;
	return
L166:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v344 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_moveouts(m, v345, l1, v342)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L29
	} else {
		goto L168
	}
L168:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_dupnfa(m, v348, v342, v25, l1, v342)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L29
	} else {
		goto L169
	}
L169:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v351 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	v352 = int32(1)
	F_repeat_1(m, l0, l1, v342, v352, v12-v352)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L29
	} else {
		goto L171
	}
L171:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v357 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_newarc(m, v358, l1, v342)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L29
	} else {
		goto L173
	}
L173:
	;
	return
L174:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v365 = F_newstate(m, v364)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L29
	} else {
		goto L175
	}
L175:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v367 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_moveouts(m, v368, l1, v362)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L29
	} else {
		goto L177
	}
L177:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_moveins(m, v371, v25, v365)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L29
	} else {
		goto L178
	}
L178:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v376 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v376 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L29
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v362)+8))
	if v379 <= v380 {
		goto L185
	} else {
		goto L186
	}
L182:
	;
	goto L181
L183:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v436 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v436 != 0 {
		goto L205
	} else {
		goto L206
	}
L184:
	;
	F_createarc(m, v374, int32(110), int32(0), l1, v362)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L29
	} else {
		goto L204
	}
L185:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v382 == int32(0) {
		goto L184
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v362)+16))
	if v399 == int32(0) {
		goto L184
	} else {
		goto L196
	}
L188:
	;
	v388 = v382
	goto L189
L189:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v388)+12))
	if v392 != v362 {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	goto L184
L191:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v388)+16))
	if v398 != 0 {
		v388 = v398
		goto L189
	} else {
		goto L195
	}
L192:
	;
	v394 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v388)+4)))
	if v394 != 0 {
		goto L191
	} else {
		goto L193
	}
L193:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	if v395 == int32(110) {
		goto L183
	} else {
		goto L194
	}
L194:
	;
	goto L191
L195:
	;
	goto L190
L196:
	;
	v405 = v399
	goto L197
L197:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v405)+8))
	if v409 != l1 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	goto L184
L199:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v405)+24))
	if v415 != 0 {
		v405 = v415
		goto L197
	} else {
		goto L203
	}
L200:
	;
	v411 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v405)+4)))
	if v411 != 0 {
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	if v412 == int32(110) {
		goto L183
	} else {
		goto L202
	}
L202:
	;
	goto L199
L203:
	;
	goto L198
L204:
	;
	goto L183
L205:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L29
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v365)+12))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v439 <= v440 {
		goto L211
	} else {
		goto L212
	}
L208:
	;
	goto L207
L209:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v496 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v496 != 0 {
		goto L231
	} else {
		goto L232
	}
L210:
	;
	F_createarc(m, v434, int32(110), int32(0), v365, v25)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L29
	} else {
		goto L230
	}
L211:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v365)+20))
	if v442 == int32(0) {
		goto L210
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v459 == int32(0) {
		goto L210
	} else {
		goto L222
	}
L214:
	;
	v448 = v442
	goto L215
L215:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v448)+12))
	if v452 != v25 {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	goto L210
L217:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v448)+16))
	if v458 != 0 {
		v448 = v458
		goto L215
	} else {
		goto L221
	}
L218:
	;
	v454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v448)+4)))
	if v454 != 0 {
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	if v455 == int32(110) {
		goto L209
	} else {
		goto L220
	}
L220:
	;
	goto L217
L221:
	;
	goto L216
L222:
	;
	v465 = v459
	goto L223
L223:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v465)+8))
	if v469 != v365 {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	goto L210
L225:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v465)+24))
	if v475 != 0 {
		v465 = v475
		goto L223
	} else {
		goto L229
	}
L226:
	;
	v471 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v465)+4)))
	if v471 != 0 {
		goto L225
	} else {
		goto L227
	}
L227:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v465)))
	if v472 == int32(110) {
		goto L209
	} else {
		goto L228
	}
L228:
	;
	goto L225
L229:
	;
	goto L224
L230:
	;
	goto L209
L231:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L29
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v365)+12))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v362)+8))
	if v499 <= v500 {
		goto L236
	} else {
		goto L237
	}
L234:
	;
	goto L233
L235:
	;
	F_createarc(m, v494, int32(110), int32(0), v365, v362)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L29
	} else {
		goto L255
	}
L236:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v365)+20))
	if v502 == int32(0) {
		goto L235
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v362)+16))
	if v519 == int32(0) {
		goto L235
	} else {
		goto L247
	}
L239:
	;
	v508 = v502
	goto L240
L240:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v508)+12))
	if v512 != v362 {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	goto L235
L242:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v508)+16))
	if v518 != 0 {
		v508 = v518
		goto L240
	} else {
		goto L246
	}
L243:
	;
	v514 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v508)+4)))
	if v514 != 0 {
		goto L242
	} else {
		goto L244
	}
L244:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v508)))
	if v515 == int32(110) {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	goto L242
L246:
	;
	goto L241
L247:
	;
	v525 = v519
	goto L248
L248:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v525)+8))
	if v529 != v365 {
		goto L250
	} else {
		goto L251
	}
L249:
	;
	goto L235
L250:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v525)+24))
	if v535 != 0 {
		v525 = v535
		goto L248
	} else {
		goto L254
	}
L251:
	;
	v531 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v525)+4)))
	if v531 != 0 {
		goto L250
	} else {
		goto L252
	}
L252:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v525)))
	if v532 == int32(110) {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	goto L250
L254:
	;
	goto L249
L255:
	;
	return
L256:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v550 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_moveouts(m, v551, l1, v548)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L29
	} else {
		goto L258
	}
L258:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_dupnfa(m, v554, v548, v25, l1, v548)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L29
	} else {
		goto L259
	}
L259:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v557 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	v558 = int32(1)
	v10 = v548
	v11 = v26 - v558
	v12 = v12 - v558
	goto L2
L261:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v565 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_moveouts(m, v566, l1, v563)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L29
	} else {
		goto L263
	}
L263:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_dupnfa(m, v569, v563, v25, l1, v563)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L29
	} else {
		goto L264
	}
L264:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v572 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	v25 = v563
	v26 = v26 - int32(1)
	goto L10
L266:
	;
	v579 = v577
	goto L268
L267:
	;
	v579 = int32(15)
	goto L268
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v579
	goto L1
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
				v21 = F_makeAlias(m, int32(628696), int32(0))
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
	var v27 int32
	_ = v27
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
	var v57 int32
	_ = v57
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
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
	v27 = l0
	goto L3
L3:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
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
	v27 = v127 + int32(1)
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
	v40 = v27 + int32(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
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
	v57 = int32(0)
	v58 = l2
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
		v76 = v57
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
		v76 = v57
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
	v78 = v58 + int32(1)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v79 != 0 {
		v54 = v79
		v57 = v76
		v58 = v78
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
	F_errmsg(m, int32(685962), v11+int32(-48))
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
	F_errdetail(m, int32(627843), v13)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(474254), int32(125), int32(126286))
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
	v127 = v27
	goto L7
L35:
	;
	v27 = v27 + int32(2)
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
	F_errmsg(m, int32(685962), v11+int32(-32))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errdetail(m, int32(628589), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(474254), int32(86), int32(126286))
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
func F_replace_relid_callback(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
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
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v321 int32
	_ = v321
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = int32(1)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v16 == int32(63) {
		v321 = v15
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v321
L2:
	;
	if v16 == int32(318) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(-1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v24 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	v321 = int32(0)
	goto L1
L6:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v74 = F_bms_is_member(m, v72, v73)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L7:
	;
	v71 = int32(0)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v33 = int32(1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v34 <= v33 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v37 = v33
	goto L12
L11:
	;
	v37 = v34
	goto L12
L12:
	;
	v40 = int32(0)
	v42 = v40
	v43 = v40
	goto L13
L13:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(8)+v42<<(uint(int32(2))%32))))
	if v51 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v71 = v64
	goto L6
L15:
	;
	goto L14
L16:
	;
	v52 = int32(2)
	if v43 != 0 {
		v64 = v52
		goto L15
	} else {
		goto L19
	}
L17:
	;
	v57 = v43
	goto L18
L18:
	;
	v60 = v42 + int32(1)
	if v60 != v37 {
		v42 = v60
		v43 = v57
		goto L13
	} else {
		goto L21
	}
L19:
	;
	v53 = int32(1)
	if base.Ui32(v53) < base.Ui32(base.I32_popcnt(v51)) {
		v64 = v52
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v57 = v53
	goto L18
L21:
	;
	v64 = v57
	goto L15
L22:
	;
	if v23 == v24 {
		goto L62
	} else {
		goto L63
	}
L23:
	;
	return int32(0)
L24:
	;
	if v74 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v82 = F_bms_is_member(m, v80, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ChangeVarNodesWalkExpression(m, v86, l1)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L23
	} else {
		goto L30
	}
L28:
	;
	if v82 == int32(0) {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_ChangeVarNodesWalkExpression(m, v89, l1)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v95 = F_adjust_relid_set(m, v92, v93, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L23
	} else {
		goto L32
	}
L32:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v98 = int32(0)
	if v97 == v98 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v134 = int32(0)
	if v95 == v134 {
		goto L47
	} else {
		goto L48
	}
L34:
	;
	v133 = int32(0)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v105 = int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if v106 <= v105 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v109 = v105
	goto L39
L38:
	;
	v109 = v106
	goto L39
L39:
	;
	v113 = int32(0)
	v115 = v98
	goto L40
L40:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v97+int32(8)+v113<<(uint(int32(2))%32))))
	if v121 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v133 = v124
	goto L33
L42:
	;
	v124 = v115 + base.I32_popcnt(v121)
	goto L44
L43:
	;
	v124 = v115
	goto L44
L44:
	;
	v126 = v113 + int32(1)
	if v126 != v109 {
		v113 = v126
		v115 = v124
		goto L40
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v95
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v171 + (v169 - v133)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v178 = F_adjust_relid_set(m, v175, v176, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L23
	} else {
		goto L59
	}
L47:
	;
	v169 = int32(0)
	goto L46
L48:
	;
	goto L49
L49:
	;
	v141 = int32(1)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v142 <= v141 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v145 = v141
	goto L52
L51:
	;
	v145 = v142
	goto L52
L52:
	;
	v149 = int32(0)
	v151 = v134
	goto L53
L53:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v95+int32(8)+v149<<(uint(int32(2))%32))))
	if v157 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v169 = v160
	goto L46
L55:
	;
	v160 = v151 + base.I32_popcnt(v157)
	goto L57
L56:
	;
	v160 = v151
	goto L57
L57:
	;
	v162 = v149 + int32(1)
	if v162 != v145 {
		v149 = v162
		v151 = v160
		goto L53
	} else {
		goto L58
	}
L58:
	;
	goto L54
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v178
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v184 = F_adjust_relid_set(m, v181, v182, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L23
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v184
	goto L22
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v197
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v202 = F_adjust_relid_set(m, v199, v200, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L23
	} else {
		goto L66
	}
L62:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v197 = v191
	goto L61
L63:
	;
	goto L64
L64:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v195 = F_adjust_relid_set(m, v192, v193, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L23
	} else {
		goto L65
	}
L65:
	;
	v197 = v195
	goto L61
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v202
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v208 = F_adjust_relid_set(m, v205, v206, v207)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L23
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v208
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v211 == int32(0) {
		v321 = v15
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v217 = int32(0)
	if v214 == v217 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v270 == int32(0) {
		v321 = v15
		goto L1
	} else {
		goto L85
	}
L70:
	;
	v270 = int32(0)
	goto L69
L71:
	;
	goto L72
L72:
	;
	v225 = int32(1)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v226 <= v225 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v229 = v225
	goto L75
L74:
	;
	v229 = v226
	goto L75
L75:
	;
	v234 = int32(0)
	v237 = int32(-1)
	goto L77
L76:
	;
	v270 = v262
	goto L69
L77:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v214+int32(8)+v234<<(uint(int32(2))%32))))
	if v244 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(12)))) = v254
	v262 = int32(1)
	goto L76
L79:
	;
	if int32(0) <= v237 {
		v262 = v217
		goto L76
	} else {
		goto L82
	}
L80:
	;
	v254 = v237
	goto L81
L81:
	;
	v256 = v234 + int32(1)
	if v256 != v229 {
		v234 = v256
		v237 = v254
		goto L77
	} else {
		goto L84
	}
L82:
	;
	if base.Ui32(int32(1)) < base.Ui32(base.I32_popcnt(v244)) {
		v262 = v217
		goto L76
	} else {
		goto L83
	}
L83:
	;
	v254 = base.I32_ctz(v244) | v234<<(uint(int32(5))%32)
	goto L81
L84:
	;
	goto L78
L85:
	;
	if v71 != int32(2) {
		v321 = v15
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v275 != v276 {
		v321 = v15
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	if v279 != int32(17) {
		v321 = v15
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v278)+28))
	if v282 == int32(0) {
		v321 = v15
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v282)+12))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	if int32(2) <= v288 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v285)+4))
	v292 = v291
	goto L92
L91:
	;
	v292 = int32(0)
	goto L92
L92:
	;
	if v286 == int32(0) {
		v321 = v15
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v295 = F_equal(m, v286, v292)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L23
	} else {
		goto L94
	}
L94:
	;
	if v295 == int32(0) {
		v321 = v15
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v300 = F_palloc0(m, int32(20))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L23
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v300)+16)) = int32(-1)
	v304 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v300)+12)) = uint8(v304)
	*(*int32)(unsafe.Add(mBase, uint32(v300)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v300)+4)) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v300))) = int32(52)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+108)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v300
	v321 = v15
	goto L1
}
func F_resolve_aggregate_transtype(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 <= int32(3830) {
		switch l1 - int32(2277) {
		case 0, 6:
			F_get_func_signature(m, l0, v7+int32(12), v7+int32(8))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
				v38 = F_enforce_generic_type_consistency(m, l2, v35, v36, l1, int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					F_pfree(m, v40)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v43 = v38
						m.G0 = v7 + int32(16)
						return v43
					}
				}
			}
		case 1, 2, 3, 4, 5:
			v43 = l1
			m.G0 = v7 + int32(16)
			return v43
		default:
			if l1 == int32(2776) {
				F_get_func_signature(m, l0, v7+int32(12), v7+int32(8))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
					v38 = F_enforce_generic_type_consistency(m, l2, v35, v36, l1, int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						F_pfree(m, v40)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v43 = v38
							m.G0 = v7 + int32(16)
							return v43
						}
					}
				}
			} else {
				if l1 == int32(3500) {
					F_get_func_signature(m, l0, v7+int32(12), v7+int32(8))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
						v38 = F_enforce_generic_type_consistency(m, l2, v35, v36, l1, int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
							F_pfree(m, v40)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = v38
								m.G0 = v7 + int32(16)
								return v43
							}
						}
					}
				} else {
					v43 = l1
					m.G0 = v7 + int32(16)
					return v43
				}
			}
		}
	} else {
		if base.Ui32(l1-int32(5077)) < base.Ui32(int32(4)) {
			F_get_func_signature(m, l0, v7+int32(12), v7+int32(8))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
				v38 = F_enforce_generic_type_consistency(m, l2, v35, v36, l1, int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					F_pfree(m, v40)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v43 = v38
						m.G0 = v7 + int32(16)
						return v43
					}
				}
			}
		} else {
			if base.Ui32(l1-int32(4537)) < base.Ui32(int32(2)) {
				F_get_func_signature(m, l0, v7+int32(12), v7+int32(8))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
					v38 = F_enforce_generic_type_consistency(m, l2, v35, v36, l1, int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						F_pfree(m, v40)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v43 = v38
							m.G0 = v7 + int32(16)
							return v43
						}
					}
				}
			} else {
				if l1 != int32(3831) {
					v43 = l1
					m.G0 = v7 + int32(16)
					return v43
				} else {
					F_get_func_signature(m, l0, v7+int32(12), v7+int32(8))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
						v38 = F_enforce_generic_type_consistency(m, l2, v35, v36, l1, int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
							F_pfree(m, v40)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = v38
								m.G0 = v7 + int32(16)
								return v43
							}
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v3 = int32(0)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)))
	if v6 != 0 {
		v78 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v78
L2:
	;
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v7 != 0 {
		v78 = v3
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
		v78 = v3
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	goto L21
L7:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)))
	if v13 != 0 {
		v78 = v3
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v15 != int32(6) {
		v78 = v3
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v18 != 0 {
		v78 = v3
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+8)))
	if v19 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(1)
L12:
	;
	goto L13
L13:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v25 = F_find_base_rel(m, l0, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+8)))
	if int32(0) < v30 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
	v34 = F_bms_is_member(m, v30, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	return int32(0)
L19:
	;
	if v34 != 0 {
		v78 = int32(1)
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	if base.B2i32(v38 != int32(0)) == int32(0) {
		v78 = v3
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v43 = int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	if v45 == int32(0) {
		v78 = v43
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v48 = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v49 <= v48 {
		v78 = v43
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v53 = v48
	goto L25
L25:
	;
	v57 = int32(0)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v53<<(uint(int32(2))%32))))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v63 != int32(318) {
		v78 = v57
		goto L1
	} else {
		goto L27
	}
L26:
	;
	v78 = v70
	goto L1
L27:
	;
	v66 = F_restriction_is_always_false(m, l0, v62)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L14
	} else {
		goto L28
	}
L28:
	;
	if v66 == int32(0) {
		v78 = v57
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v70 = int32(1)
	v72 = v53 + v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v72 < v73 {
		v53 = v72
		goto L25
	} else {
		goto L30
	}
L30:
	;
	goto L26
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
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
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
					v32 = F_lookup_type_cache(m, v30, int32(4352))
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
									F_errstart_cold(m, int32(21), int32(525467))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										F_errcode(m, int32(151027844))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
											v70 = F_format_type_be(m, v69)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v70
												F_errmsg(m, int32(332358), v8)
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return
												} else {
													F_errfinish(m, int32(476889), int32(7088), int32(415389))
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
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
								F_errstart_cold(m, int32(21), int32(525467))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									F_errcode(m, int32(151027844))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
										v70 = F_format_type_be(m, v69)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v70
											F_errmsg(m, int32(332358), v8)
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return
											} else {
												F_errfinish(m, int32(476889), int32(7088), int32(415389))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
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
				v32 = F_lookup_type_cache(m, v30, int32(4352))
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
								F_errstart_cold(m, int32(21), int32(525467))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									F_errcode(m, int32(151027844))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
										v70 = F_format_type_be(m, v69)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v70
											F_errmsg(m, int32(332358), v8)
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return
											} else {
												F_errfinish(m, int32(476889), int32(7088), int32(415389))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
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
							F_errstart_cold(m, int32(21), int32(525467))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								F_errcode(m, int32(151027844))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
									v70 = F_format_type_be(m, v69)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v70
										F_errmsg(m, int32(332358), v8)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return
										} else {
											F_errfinish(m, int32(476889), int32(7088), int32(415389))
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
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
						v32 = F_lookup_type_cache(m, v30, int32(4352))
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
										F_errstart_cold(m, int32(21), int32(525467))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return
										} else {
											F_errcode(m, int32(151027844))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
												v70 = F_format_type_be(m, v69)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8))) = v70
													F_errmsg(m, int32(332358), v8)
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return
													} else {
														F_errfinish(m, int32(476889), int32(7088), int32(415389))
														mBase = m.M
														v83 = m.ExcPending
														if v83 != 0 {
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
									F_errstart_cold(m, int32(21), int32(525467))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										F_errcode(m, int32(151027844))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
											v70 = F_format_type_be(m, v69)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v70
												F_errmsg(m, int32(332358), v8)
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return
												} else {
													F_errfinish(m, int32(476889), int32(7088), int32(415389))
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
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
					v32 = F_lookup_type_cache(m, v30, int32(4352))
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
									F_errstart_cold(m, int32(21), int32(525467))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										F_errcode(m, int32(151027844))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
											v70 = F_format_type_be(m, v69)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v70
												F_errmsg(m, int32(332358), v8)
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return
												} else {
													F_errfinish(m, int32(476889), int32(7088), int32(415389))
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
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
								F_errstart_cold(m, int32(21), int32(525467))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									F_errcode(m, int32(151027844))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
										v70 = F_format_type_be(m, v69)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v70
											F_errmsg(m, int32(332358), v8)
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return
											} else {
												F_errfinish(m, int32(476889), int32(7088), int32(415389))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
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
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
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
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
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
	return v269
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[108]))
	if v41 != 0 {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1033])))
	if v30 != l0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v30 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32))+uint32(_consts[1034])))
	v269 = v38
	goto L4
L9:
	;
	v242 = int32(4442992)
	v243 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v246 = *(*int32)(unsafe.Add(mBase, _consts[146]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v246
	v248 = F_list_copy(m, v232)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L15
	} else {
		goto L56
	}
L10:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	if v220 == int32(0) {
		v232 = v210
		goto L9
	} else {
		goto L54
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L15
	} else {
		goto L51
	}
L12:
	;
	v43 = F_SearchSysCache1(m, int32(21), v41)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v57 = v5
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = l0
	v63 = F_list_make1_impl(m, int32(472), v20+int32(12))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L15
	} else {
		goto L19
	}
L15:
	;
	return int32(0)
L16:
	;
	if v43 == int32(0) {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+22)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49+v50)+68))
	F_ReleaseCatCache(m, v43)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v57 = v52
	goto L14
L19:
	;
	if v63 == int32(0) {
		v232 = v5
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v67 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v210 = v63
	goto L10
L22:
	;
	goto L23
L23:
	;
	v79 = v63
	v84 = v5
	goto L24
L24:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91+v84<<(uint(int32(2))%32))))
	v96 = int32(0)
	v98 = F_SearchSysCacheList(m, int32(8), int32(1), v95, v96, v96)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L15
	} else {
		goto L26
	}
L25:
	;
	v210 = v181
	goto L10
L26:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v98)+40))
	if int32(0) < v100 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v110 = int32(0)
	v113 = v79
	goto L30
L28:
	;
	v161 = v79
	goto L29
L29:
	;
	F_ReleaseCatCacheList(m, v98)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L15
	} else {
		goto L45
	}
L30:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v98+int32(48)+v110<<(uint(int32(2))%32))))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+56))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+22)))
	v129 = v127 + v128
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v130 != l2 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v161 = v149
	goto L29
L32:
	;
	switch l1 - int32(1) {
	case 0:
		goto L40
	case 1:
		goto L39
	default:
		goto L38
	}
L33:
	;
	if l2 == int32(0) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+16)))
	if v134&int32(1) == int32(0) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v139 != 0 {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v95
	goto L32
L37:
	;
	v151 = v110 + int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v98)+40))
	if v151 < v152 {
		v110 = v151
		v113 = v149
		goto L30
	} else {
		goto L44
	}
L38:
	;
	v147 = F_roles_list_append(m, v113, v20+int32(28), v130)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L15
	} else {
		goto L43
	}
L39:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+18)))
	if v142 != int32(1) {
		v149 = v113
		goto L37
	} else {
		goto L42
	}
L40:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+17)))
	if v141 != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v149 = v113
	goto L37
L42:
	;
	goto L38
L43:
	;
	v149 = v147
	goto L37
L44:
	;
	goto L31
L45:
	;
	if v57 == int32(0) {
		v181 = v161
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v183 = v84 + int32(1)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v183 < v184 {
		v79 = v181
		v84 = v183
		goto L24
	} else {
		goto L50
	}
L47:
	;
	if v57 != v95 {
		v181 = v161
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v179 = F_roles_list_append(m, v161, v20+int32(28), int32(6171))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L15
	} else {
		goto L49
	}
L49:
	;
	v181 = v179
	goto L46
L50:
	;
	goto L25
L51:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _consts[108]))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v191
	F_errmsg_internal(m, int32(46947), v20+int32(16))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L15
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(474450), int32(5185), int32(321990))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L15
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	F_pfree(m, v220)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L15
	} else {
		goto L55
	}
L55:
	;
	v232 = v210
	goto L9
L56:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v243
	F_list_free(m, v232)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L15
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1033]))) = int32(0)
	v257 = l1 << (uint(int32(2)) % 32)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v257)+uint32(_consts[1034])))
	F_list_free(m, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L15
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257)+uint32(_consts[1034]))) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1033]))) = l0
	v269 = v248
	goto L4
}
func F_rtree_internal_consistent(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	switch l2 - int32(1) {
	case 0:
		v88 = F_DirectFunctionCall2Coll(m, int32(104), int32(0), l0, l1)
		mBase = m.M
		v89 = m.ExcPending
		if v89 != 0 {
			return int32(0)
		} else {
			v92 = base.B2i32(v88 == int32(0))
			m.G0 = v7 + int32(16)
			return v92
		}
	case 1:
		v13 = F_DirectFunctionCall2Coll(m, int32(95), int32(0), l0, l1)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v92 = base.B2i32(v13 == int32(0))
			m.G0 = v7 + int32(16)
			return v92
		}
	case 2:
		v21 = F_DirectFunctionCall2Coll(m, int32(96), int32(0), l0, l1)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v92 = base.B2i32(v21 != int32(0))
			m.G0 = v7 + int32(16)
			return v92
		}
	case 3:
		v27 = F_DirectFunctionCall2Coll(m, int32(97), int32(0), l0, l1)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v92 = base.B2i32(v27 == int32(0))
			m.G0 = v7 + int32(16)
			return v92
		}
	case 4:
		v33 = F_DirectFunctionCall2Coll(m, int32(98), int32(0), l0, l1)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			v92 = base.B2i32(v33 == int32(0))
			m.G0 = v7 + int32(16)
			return v92
		}
	case 5, 6:
		v39 = F_DirectFunctionCall2Coll(m, int32(99), int32(0), l0, l1)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			v92 = base.B2i32(v39 != int32(0))
			m.G0 = v7 + int32(16)
			return v92
		}
	case 7:
		v45 = F_DirectFunctionCall2Coll(m, int32(96), int32(0), l0, l1)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			v92 = base.B2i32(v45 != int32(0))
			m.G0 = v7 + int32(16)
			return v92
		}
	case 8:
		v51 = F_DirectFunctionCall2Coll(m, int32(100), int32(0), l0, l1)
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			v92 = base.B2i32(v51 == int32(0))
			m.G0 = v7 + int32(16)
			return v92
		}
	case 9:
		v57 = F_DirectFunctionCall2Coll(m, int32(101), int32(0), l0, l1)
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return int32(0)
		} else {
			v92 = base.B2i32(v57 == int32(0))
			m.G0 = v7 + int32(16)
			return v92
		}
	case 10:
		v63 = F_DirectFunctionCall2Coll(m, int32(102), int32(0), l0, l1)
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return int32(0)
		} else {
			v92 = base.B2i32(v63 == int32(0))
			m.G0 = v7 + int32(16)
			return v92
		}
	case 11:
		v69 = F_DirectFunctionCall2Coll(m, int32(103), int32(0), l0, l1)
		mBase = m.M
		v70 = m.ExcPending
		if v70 != 0 {
			return int32(0)
		} else {
			v92 = base.B2i32(v69 == int32(0))
			m.G0 = v7 + int32(16)
			return v92
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v76 = m.ExcPending
		if v76 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l2
			F_errmsg_internal(m, int32(460098), v7)
			mBase = m.M
			v80 = m.ExcPending
			if v80 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(476719), int32(1020), int32(87429))
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
	}
}
