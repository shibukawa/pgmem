package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__pgp_read_public_key(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_palloc0(m, int32(52))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v19 = F_pullf_read_fixed(m, l0, int32(1), v9+int32(15))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v19 < int32(0) {
				v128 = v19
				m.G0 = v9 + int32(16)
				return v128
			} else {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
				*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v23)
				if v23 != int32(4) {
					v116 = int32(-117)
					F_pgp_key_free(m, v12)
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						v128 = v116
						m.G0 = v9 + int32(16)
						return v128
					}
				} else {
					v31 = F_pullf_read_fixed(m, l0, int32(4), v12+int32(1))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if v31 < int32(0) {
							v116 = v31
							F_pgp_key_free(m, v12)
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								v128 = v116
								m.G0 = v9 + int32(16)
								return v128
							}
						} else {
							v38 = F_pullf_read_fixed(m, l0, int32(1), v9+int32(14))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								if v38 < int32(0) {
									v128 = v38
									m.G0 = v9 + int32(16)
									return v128
								} else {
									v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)))
									*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)) = uint8(v42)
									switch v42 - int32(1) {
									case 0, 1, 2:
										v74 = F_pgp_mpi_read(m, l0, v12+int32(8))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
										} else {
											if v74 < int32(0) {
												v116 = v74
												F_pgp_key_free(m, v12)
												mBase = m.M
												v118 = m.ExcPending
												if v118 != 0 {
													return int32(0)
												} else {
													v128 = v116
													m.G0 = v9 + int32(16)
													return v128
												}
											} else {
												v80 = F_pgp_mpi_read(m, l0, v12+int32(12))
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													if v80 < int32(0) {
														v116 = v80
														F_pgp_key_free(m, v12)
														mBase = m.M
														v118 = m.ExcPending
														if v118 != 0 {
															return int32(0)
														} else {
															v128 = v116
															m.G0 = v9 + int32(16)
															return v128
														}
													} else {
														v84 = F_calc_key_id(m, v12)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return int32(0)
														} else {
															v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)))
															if v86 == int32(3) {
																v113 = v84
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(1)
																v113 = v84
															}
															if int32(0) <= v113 {
																*(*int32)(unsafe.Add(mBase, uint32(l1))) = v12
																v128 = int32(0)
																m.G0 = v9 + int32(16)
																return v128
															} else {
																v116 = v113
																F_pgp_key_free(m, v12)
																mBase = m.M
																v118 = m.ExcPending
																if v118 != 0 {
																	return int32(0)
																} else {
																	v128 = v116
																	m.G0 = v9 + int32(16)
																	return v128
																}
															}
														}
													}
												}
											}
										}
									default:
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v42
										F_px_debug(m, int32(477954), v9)
										mBase = m.M
										v124 = m.ExcPending
										if v124 != 0 {
											return int32(0)
										} else {
											F_pgp_key_free(m, v12)
											mBase = m.M
											v127 = m.ExcPending
											if v127 != 0 {
												return int32(0)
											} else {
												v128 = int32(-112)
												m.G0 = v9 + int32(16)
												return v128
											}
										}
									case 15:
										v93 = F_pgp_mpi_read(m, l0, v12+int32(8))
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return int32(0)
										} else {
											if v93 < int32(0) {
												v116 = v93
												F_pgp_key_free(m, v12)
												mBase = m.M
												v118 = m.ExcPending
												if v118 != 0 {
													return int32(0)
												} else {
													v128 = v116
													m.G0 = v9 + int32(16)
													return v128
												}
											} else {
												v99 = F_pgp_mpi_read(m, l0, v12+int32(12))
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return int32(0)
												} else {
													if v99 < int32(0) {
														v116 = v99
														F_pgp_key_free(m, v12)
														mBase = m.M
														v118 = m.ExcPending
														if v118 != 0 {
															return int32(0)
														} else {
															v128 = v116
															m.G0 = v9 + int32(16)
															return v128
														}
													} else {
														v105 = F_pgp_mpi_read(m, l0, v12+int32(16))
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return int32(0)
														} else {
															if v105 < int32(0) {
																v116 = v105
																F_pgp_key_free(m, v12)
																mBase = m.M
																v118 = m.ExcPending
																if v118 != 0 {
																	return int32(0)
																} else {
																	v128 = v116
																	m.G0 = v9 + int32(16)
																	return v128
																}
															} else {
																v109 = F_calc_key_id(m, v12)
																mBase = m.M
																v110 = m.ExcPending
																if v110 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(1)
																	v113 = v109
																	if int32(0) <= v113 {
																		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v12
																		v128 = int32(0)
																		m.G0 = v9 + int32(16)
																		return v128
																	} else {
																		v116 = v113
																		F_pgp_key_free(m, v12)
																		mBase = m.M
																		v118 = m.ExcPending
																		if v118 != 0 {
																			return int32(0)
																		} else {
																			v128 = v116
																			m.G0 = v9 + int32(16)
																			return v128
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									case 16:
										v48 = F_pgp_mpi_read(m, l0, v12+int32(8))
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return int32(0)
										} else {
											if v48 < int32(0) {
												v116 = v48
												F_pgp_key_free(m, v12)
												mBase = m.M
												v118 = m.ExcPending
												if v118 != 0 {
													return int32(0)
												} else {
													v128 = v116
													m.G0 = v9 + int32(16)
													return v128
												}
											} else {
												v54 = F_pgp_mpi_read(m, l0, v12+int32(12))
												mBase = m.M
												v55 = m.ExcPending
												if v55 != 0 {
													return int32(0)
												} else {
													if v54 < int32(0) {
														v116 = v54
														F_pgp_key_free(m, v12)
														mBase = m.M
														v118 = m.ExcPending
														if v118 != 0 {
															return int32(0)
														} else {
															v128 = v116
															m.G0 = v9 + int32(16)
															return v128
														}
													} else {
														v60 = F_pgp_mpi_read(m, l0, v12+int32(16))
														mBase = m.M
														v61 = m.ExcPending
														if v61 != 0 {
															return int32(0)
														} else {
															if v60 < int32(0) {
																v116 = v60
																F_pgp_key_free(m, v12)
																mBase = m.M
																v118 = m.ExcPending
																if v118 != 0 {
																	return int32(0)
																} else {
																	v128 = v116
																	m.G0 = v9 + int32(16)
																	return v128
																}
															} else {
																v66 = F_pgp_mpi_read(m, l0, v12+int32(20))
																mBase = m.M
																v67 = m.ExcPending
																if v67 != 0 {
																	return int32(0)
																} else {
																	if v66 < int32(0) {
																		v116 = v66
																		F_pgp_key_free(m, v12)
																		mBase = m.M
																		v118 = m.ExcPending
																		if v118 != 0 {
																			return int32(0)
																		} else {
																			v128 = v116
																			m.G0 = v9 + int32(16)
																			return v128
																		}
																	} else {
																		v70 = F_calc_key_id(m, v12)
																		mBase = m.M
																		v71 = m.ExcPending
																		if v71 != 0 {
																			return int32(0)
																		} else {
																			v113 = v70
																			if int32(0) <= v113 {
																				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v12
																				v128 = int32(0)
																				m.G0 = v9 + int32(16)
																				return v128
																			} else {
																				v116 = v113
																				F_pgp_key_free(m, v12)
																				mBase = m.M
																				v118 = m.ExcPending
																				if v118 != 0 {
																					return int32(0)
																				} else {
																					v128 = v116
																					m.G0 = v9 + int32(16)
																					return v128
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pgp_cfb_encrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v7 != 0 {
		v8 = int32(6563)
	} else {
		v8 = int32(6564)
	}
	F_cfb_process(m, l0, l1, l2, l3, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_pgp_compress_filter(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_pushf_create(m, l0, int32(4368528), l1, l2)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_pgp_decrypt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
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
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
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
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
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
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v4
	v22 = F_pullf_create_mbuf_reader(m, v14+int32(28), l1)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	if v513 != 0 {
		goto L148
	} else {
		goto L149
	}
L2:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v497 == int32(0) {
		v503 = v487
		v509 = v493
		goto L1
	} else {
		goto L146
	}
L3:
	;
	return int32(0)
L4:
	;
	if v22 < int32(0) {
		v487 = v22
		v493 = v4
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v31 = l0 + int32(132)
	v37 = v4
	v39 = v4
	goto L6
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v49 = F_pgp_parse_pkt_hdr(m, v43, v14+int32(23), v14+int32(16), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v503 = v469
	v509 = v475
	goto L1
L8:
	;
	if v49 <= int32(0) {
		v487 = v49
		v493 = v39
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v56 = F_palloc(m, int32(8))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v49
	v63 = F_pullf_create(m, v14+int32(24), int32(4368576), v56, v53)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	if v63 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_pfree(m, v56)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L3
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+23)))
	switch v69 - int32(1) {
	case 0:
		goto L24
	default:
		goto L20
	case 2:
		goto L23
	case 8:
		goto L22
	case 9:
		goto L25
	case 17:
		goto L21
	}
L15:
	;
	v487 = v63
	v493 = v39
	goto L2
L16:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	F_pullf_free(m, v479)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L3
	} else {
		goto L144
	}
L17:
	;
	v465 = int32(1)
	v469 = int32(-100)
	v473 = v465
	v475 = v465
	goto L16
L18:
	;
	v463 = int32(1)
	v469 = v461
	v473 = v463
	v475 = v463
	goto L16
L19:
	;
	v469 = int32(-100)
	v473 = v459
	v475 = v39
	goto L16
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v69
	F_px_debug(m, int32(29465), v14)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L3
	} else {
		goto L143
	}
L21:
	;
	if v37 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L22:
	;
	if v37 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L23:
	;
	if v37 != 0 {
		goto L31
	} else {
		goto L32
	}
L24:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v93 = F_pgp_parse_pubenc_sesskey(m, l0, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L3
	} else {
		goto L30
	}
L25:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	goto L26
L26:
	;
	v87 = F_pullf_read(m, v72, int32(32768), v14+int32(32))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L3
	} else {
		goto L28
	}
L27:
	;
	v469 = v87
	v473 = v37
	v475 = v39
	goto L16
L28:
	;
	if int32(0) < v87 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v469 = v93
	v473 = int32(1)
	v475 = v39
	goto L16
L31:
	;
	F_px_debug(m, int32(111997), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L3
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v104 = F_pullf_read_fixed(m, v100, int32(1), v14+int32(32))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L36
	}
L34:
	;
	v459 = int32(1)
	goto L19
L35:
	;
	v469 = v286
	v473 = int32(1)
	v475 = v39
	goto L16
L36:
	;
	if v104 < int32(0) {
		v286 = v104
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+32)))
	v112 = F_pullf_read_fixed(m, v100, int32(1), v14+int32(32))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	if v112 < int32(0) {
		v286 = v112
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+32)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v116
	if v108&int32(255) != int32(4) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_px_debug(m, int32(213434), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v128 = F_pgp_s2k_read(m, v100, l0)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L3
	} else {
		goto L44
	}
L43:
	;
	v469 = int32(-100)
	v473 = int32(1)
	v475 = v39
	goto L16
L44:
	;
	if v128 < int32(0) {
		v286 = v128
		goto L35
	} else {
		goto L45
	}
L45:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v132
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v134
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = (v136&int32(15) | int32(16)) << (uint(int32(base.Ui32(v136)>>(uint(int32(4))%32))+int32(6)) % 32)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v150 = F_pgp_s2k_process(m, l0, v147, v148, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	if v150 < int32(0) {
		v286 = v150
		goto L35
	} else {
		goto L47
	}
L47:
	;
	v159 = F_pullf_read_max(m, v100, int32(34), v14+int32(76), v14+int32(32))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	if v159 < int32(0) {
		v286 = v159
		goto L35
	} else {
		goto L49
	}
L49:
	;
	if v159 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v285 = F___memset(m, v14+int32(32), int32(0), int32(34))
	mBase = m.M
	goto L78
L51:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)))
	if v165 != 0 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	goto L53
L53:
	;
	if base.Ui32(v159-int32(34)) <= base.Ui32(int32(-18)) {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v165
	v169 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v169
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v171
	v280 = v169
	goto L50
L55:
	;
	v166 = F__emscripten_memcpy_bulkmem(m, v31, l0+int32(11), v165)
	mBase = m.M
	goto L57
L56:
	;
	goto L57
L57:
	;
	goto L54
L58:
	;
	F_px_debug(m, int32(499957), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L3
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = int32(1)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	v187 = m.G0
	v189 = v187 - int32(32)
	m.G0 = v189
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)))
	v197 = int32(0)
	v199 = F_pgp_cfb_create(m, v189+int32(24), v193, l0+int32(11), v196, v197, v197)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L3
	} else {
		goto L63
	}
L61:
	;
	v469 = int32(-100)
	v473 = int32(1)
	v475 = v39
	goto L16
L62:
	;
	m.G0 = v189 + int32(32)
	v280 = v272
	goto L50
L63:
	;
	if v199 < int32(0) {
		v272 = v199
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v189)+24))
	v207 = F_pgp_cfb_decrypt(m, v203, v186, int32(1), v189+int32(31))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v189)+24))
	v210 = int32(1)
	v213 = v159 - v210
	v216 = F_pgp_cfb_decrypt(m, v209, v186+v210, v213, l0+int32(132))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L3
	} else {
		goto L66
	}
L66:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v189)+24))
	F_pgp_cfb_free(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v213
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+31)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v222
	v224 = int32(0)
	v227 = v222 - int32(2)
	if base.Ui32(int32(8)) < base.Ui32(v227) {
		v242 = v224
		goto L69
	} else {
		goto L70
	}
L68:
	;
	if v242 == v213 {
		v272 = v224
		goto L62
	} else {
		goto L72
	}
L69:
	;
	goto L68
L70:
	;
	if int32(base.Ui32(int32(487))>>(uint(v227)%32))&int32(1) == int32(0) {
		v242 = v224
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v227<<(uint(int32(2))%32))+uint32(_consts[1453])))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+12))
	v242 = v241
	goto L69
L72:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+31)))
	v245 = int32(0)
	v247 = v244 - int32(2)
	if base.Ui32(int32(8)) < base.Ui32(v247) {
		v262 = v245
		goto L74
	} else {
		goto L75
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189)+8)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v189)+4)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v244
	F_px_debug(m, int32(461560), v189)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L3
	} else {
		goto L77
	}
L74:
	;
	goto L73
L75:
	;
	if int32(base.Ui32(int32(487))>>(uint(v247)%32))&int32(1) == int32(0) {
		v262 = v245
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v247<<(uint(int32(2))%32))+uint32(_consts[1453])))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
	v262 = v261
	goto L74
L77:
	;
	v272 = int32(-100)
	goto L62
L78:
	;
	v286 = v280
	goto L35
L79:
	;
	v293 = int32(0)
	F_px_debug(m, int32(21558), v293)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L3
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	if v39 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v459 = v293
	goto L19
L83:
	;
	F_px_debug(m, int32(106932), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L3
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v302 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v302
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v305 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v305
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v305
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v305
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v317 = F_pgp_cfb_create(m, v14+int32(32), v313, v31, v314, v302, v305)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L3
	} else {
		goto L88
	}
L86:
	;
	goto L17
L87:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	if v351 != 0 {
		goto L99
	} else {
		goto L100
	}
L88:
	;
	if v317 < int32(0) {
		v349 = v317
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v325 = F_pullf_create(m, v14+int32(76), int32(4368588), v324, v304)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L3
	} else {
		goto L90
	}
L90:
	;
	if v325 < int32(0) {
		v349 = v325
		goto L87
	} else {
		goto L91
	}
L91:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	v333 = F_pullf_create(m, v14+int32(72), int32(4368600), l0, v332)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L3
	} else {
		goto L92
	}
L92:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	if int32(0) <= v333 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v340 = F_process_data_packets(m, l0, l2, v335, int32(1), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L3
	} else {
		goto L96
	}
L94:
	;
	v343 = v333
	v344 = v335
	goto L95
L95:
	;
	if v344 == int32(0) {
		v349 = v343
		goto L87
	} else {
		goto L97
	}
L96:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v343 = v340
	v344 = v342
	goto L95
L97:
	;
	F_pullf_free(m, v344)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L3
	} else {
		goto L98
	}
L98:
	;
	v349 = v343
	goto L87
L99:
	;
	F_pullf_free(m, v351)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L3
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	if v354 == int32(0) {
		v461 = v349
		goto L18
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	F_pgp_cfb_free(m, v354)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L3
	} else {
		goto L104
	}
L104:
	;
	v461 = v349
	goto L18
L105:
	;
	v361 = int32(0)
	F_px_debug(m, int32(21558), v361)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L3
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if v39 != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v459 = v361
	goto L19
L109:
	;
	F_px_debug(m, int32(437375), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L3
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v370 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v370
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v370
	v384 = F_pullf_read_fixed(m, v372, int32(1), v14+int32(67))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L3
	} else {
		goto L113
	}
L112:
	;
	goto L17
L113:
	;
	if v384 < int32(0) {
		v461 = v384
		goto L18
	} else {
		goto L114
	}
L114:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+67)))
	if v388 != int32(1) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	F_px_debug(m, int32(549138), int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L3
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v400 = int32(0)
	v402 = F_pgp_cfb_create(m, v14+int32(32), v398, v31, v399, v400, v400)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L3
	} else {
		goto L120
	}
L118:
	;
	v461 = int32(-100)
	goto L18
L119:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	if v444 != 0 {
		goto L133
	} else {
		goto L134
	}
L120:
	;
	if v402 < int32(0) {
		v442 = v402
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v410 = F_pullf_create(m, v14+int32(76), int32(4368588), v409, v372)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L3
	} else {
		goto L122
	}
L122:
	;
	if v410 < int32(0) {
		v442 = v410
		goto L119
	} else {
		goto L123
	}
L123:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	v418 = F_pullf_create(m, v14+int32(68), int32(4368624), l0, v417)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L3
	} else {
		goto L124
	}
L124:
	;
	if v418 < int32(0) {
		v442 = v418
		goto L119
	} else {
		goto L125
	}
L125:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	v426 = F_pullf_create(m, v14+int32(72), int32(4368600), l0, v425)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L3
	} else {
		goto L126
	}
L126:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	if int32(0) <= v426 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v431 = int32(1)
	v433 = F_process_data_packets(m, l0, l2, v428, v431, v431)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L3
	} else {
		goto L130
	}
L128:
	;
	v436 = v426
	v437 = v428
	goto L129
L129:
	;
	if v437 == int32(0) {
		v442 = v436
		goto L119
	} else {
		goto L131
	}
L130:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v436 = v433
	v437 = v435
	goto L129
L131:
	;
	F_pullf_free(m, v437)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L3
	} else {
		goto L132
	}
L132:
	;
	v442 = v436
	goto L119
L133:
	;
	F_pullf_free(m, v444)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L3
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	if v447 != 0 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	goto L135
L137:
	;
	F_pullf_free(m, v447)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L3
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	if v450 == int32(0) {
		v461 = v442
		goto L18
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	F_pgp_cfb_free(m, v450)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L3
	} else {
		goto L142
	}
L142:
	;
	v461 = v442
	goto L18
L143:
	;
	v459 = v37
	goto L19
L144:
	;
	v482 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v482
	if v482 <= v469 {
		v37 = v473
		v39 = v475
		goto L6
	} else {
		goto L145
	}
L145:
	;
	goto L7
L146:
	;
	F_pullf_free(m, v497)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L3
	} else {
		goto L147
	}
L147:
	;
	v503 = v487
	v509 = v493
	goto L1
L148:
	;
	F_pullf_free(m, v513)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L3
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	if v503 < int32(0) {
		v528 = v503
		goto L152
	} else {
		goto L153
	}
L151:
	;
	goto L150
L152:
	;
	m.G0 = v14 + int32(80)
	return v528
L153:
	;
	v518 = int32(-100)
	if v509 == int32(0) {
		v528 = v518
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v521 != 0 {
		v528 = v518
		goto L152
	} else {
		goto L155
	}
L155:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v523 != 0 {
		v528 = int32(-102)
		goto L152
	} else {
		goto L156
	}
L156:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v526 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v527 = int32(-106)
	goto L159
L158:
	;
	v527 = int32(0)
	goto L159
L159:
	;
	v528 = v527
	goto L152
}
func F_pgp_key_id_w(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(1)
		v12 = v7 + v11
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		v17 = v15 & v11
		if v17 != 0 {
			v18 = v12
		} else {
			v18 = v7 + int32(4)
		}
		if v15 == int32(1) {
			v21 = int32(4)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v23&int32(254) == int32(2) {
				v32 = v21
			} else {
				v32 = base.B2i32(v23 == int32(18)) << (uint(v21) % 32)
			}
			if v23 == int32(1) {
				v35 = v21
			} else {
				v35 = v32
			}
			v46 = v35
		} else {
			v36 = int32(1)
			if v17 != 0 {
				v46 = int32(base.Ui32(v15)>>(uint(v36)%32)) - v36
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v47 = F_mbuf_create_from_data(m, v18, v46)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int32(0)
		} else {
			v50 = F_palloc(m, int32(21))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				v54 = F_pgp_get_keyid(m, v47, v50+int32(4))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					v56 = F_mbuf_free(m, v47)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v54 {
							*(*int32)(unsafe.Add(mBase, uint32(v50))) = v54<<(uint(int32(2))%32) + int32(16)
							v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v65 != v7 {
								F_pfree(m, v7)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									return v50
								}
							} else {
								return v50
							}
						} else {
							F_px_THROW_ERROR(m, v54)
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
			}
		}
	}
}
func F_pgp_mpi_free(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	if l0 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v6 = F___memset(m, l0, int32(0), v3+int32(12))
		mBase = m.M
		F_pfree(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		return int32(0)
	}
}
func F_pgp_parse_pubenc_sesskey(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
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
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
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
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
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
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
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
	var v301 int32
	_ = v301
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v359 int32
	_ = v359
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v396 int32
	_ = v396
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	v3 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v21 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v19 + int32(48)
	return v448
L2:
	;
	F_px_debug(m, int32(540135), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v34 = F_pullf_read_fixed(m, l1, int32(1), v19+int32(36))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v448 = int32(-12)
	goto L1
L7:
	;
	if v34 < int32(0) {
		v448 = v34
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+36)))
	if v38 != int32(3) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v38
	F_px_debug(m, int32(461715), v19+int32(16))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v51 = F_pullf_read_fixed(m, l1, int32(8), v19+int32(36))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v448 = int32(-100)
	goto L1
L13:
	;
	if v51 < int32(0) {
		v448 = v51
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v19)+36))
	if v55 == int64(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v69 = F_pullf_read_fixed(m, l1, int32(1), v19+int32(44))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L19
	}
L16:
	;
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v19)+36))
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v21)+40))
	if v58 == v59 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_px_debug(m, int32(322262), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v448 = int32(-113)
	goto L1
L19:
	;
	if v69 < int32(0) {
		v448 = v69
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+44)))
	switch v74 - int32(1) {
	case 0, 1:
		goto L22
	default:
		v448 = int32(-112)
		goto L1
	case 15:
		goto L23
	}
L21:
	;
	if v140 < int32(0) {
		v448 = v140
		goto L1
	} else {
		goto L40
	}
L22:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+5)))
	v120 = int32(1)
	if base.Ui32(v120) < base.Ui32((v119-v120)&int32(255)) {
		v140 = int32(-113)
		goto L21
	} else {
		goto L35
	}
L23:
	;
	v79 = m.G0
	v80 = int32(16)
	v81 = v79 - v80
	m.G0 = v81
	v83 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+12)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = v83
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+5)))
	if v88 == v80 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v93 = F_pgp_mpi_read(m, l1, v81+int32(12))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L28
	}
L25:
	;
	v114 = int32(-113)
	goto L26
L26:
	;
	m.G0 = v81 + int32(16)
	v140 = v114
	goto L21
L27:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v109 = F_pgp_mpi_free(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L33
	}
L28:
	;
	if v93 < int32(0) {
		v107 = v93
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v99 = F_pgp_mpi_read(m, l1, v81+int32(8))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	if v99 < int32(0) {
		v107 = v99
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v105 = F_pgp_elgamal_decrypt(m, v21, v103, v104, v19+int32(32))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v107 = v105
	goto L27
L33:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v112 = F_pgp_mpi_free(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v114 = v107
	goto L26
L35:
	;
	v128 = F_pgp_mpi_read(m, l1, v19+int32(44))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	if v128 < int32(0) {
		v140 = v128
		goto L21
	} else {
		goto L37
	}
L37:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v135 = F_pgp_rsa_decrypt(m, v21, v132, v19+int32(32))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v138 = F_pgp_mpi_free(m, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	v140 = v135
	goto L21
L40:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
	if v146 < int32(10) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v432
	if v405 != 0 {
		goto L84
	} else {
		goto L85
	}
L42:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v430 = F_pgp_mpi_free(m, v429)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L5
	} else {
		goto L82
	}
L43:
	;
	v233 = v192 + int32(2)
	v237 = v149 - v233 + v146
	if v237 < int32(3) {
		v396 = int32(-113)
		goto L59
	} else {
		goto L60
	}
L44:
	;
	F_px_debug(m, int32(449972), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L5
	} else {
		goto L58
	}
L45:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v150 != int32(2) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v153 = v149 + v146
	v154 = int32(1)
	v156 = v149 + v154
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
	if v157 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v188 == v153 {
		goto L44
	} else {
		goto L56
	}
L48:
	;
	v188 = v156
	v190 = v154
	v192 = v149
	goto L47
L49:
	;
	goto L50
L50:
	;
	v163 = v156
	v165 = int32(0)
	goto L51
L51:
	;
	v178 = v163 + int32(1)
	if base.Ui32(v153) <= base.Ui32(v178) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v188 = v178
	v190 = base.B2i32(base.Ui32(v165) < base.Ui32(int32(7)))
	v192 = v163
	goto L47
L53:
	;
	goto L52
L54:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if v182 != 0 {
		v163 = v178
		v165 = v165 + int32(1)
		goto L51
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	v204 = int32(0)
	if (base.B2i32(v203 != v204)|v190)&int32(1) == v204 {
		goto L43
	} else {
		goto L57
	}
L57:
	;
	goto L44
L58:
	;
	v415 = int32(-113)
	goto L42
L59:
	;
	if v396 < int32(0) {
		v415 = v396
		goto L42
	} else {
		goto L79
	}
L60:
	;
	if v237 != int32(3) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v243 = int32(3)
	v244 = v237 - v243
	v246 = v244 & v243
	if base.Ui32(v237-int32(4)) < base.Ui32(v243) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v359 = int32(0)
	goto L63
L63:
	;
	v370 = v237 + v233
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370-int32(2)))))
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370-int32(1)))))
	if v359 == v373<<(uint(int32(8))%32)|v378 {
		v396 = int32(0)
		goto L59
	} else {
		goto L77
	}
L64:
	;
	if v246 != 0 {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	v298 = int32(0)
	v301 = int32(1)
	goto L64
L66:
	;
	goto L67
L67:
	;
	v265 = int32(0)
	v268 = int32(1)
	v275 = v3
	goto L68
L68:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268+(v192+int32(5))))))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268+(v192+int32(4))))))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268+(v192+int32(3))))))
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268+v233))))
	v290 = v280 + (v282 + (v284 + (v265 + v286)))
	v291 = int32(4)
	v292 = v268 + v291
	v294 = v275 + v291
	if v294 != v244&int32(-4) {
		v265 = v290
		v268 = v292
		v275 = v294
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v298 = v290
	v301 = v292
	goto L64
L70:
	;
	goto L69
L71:
	;
	v314 = v298
	v317 = v301
	v323 = v3
	goto L74
L72:
	;
	v338 = v298
	goto L73
L73:
	;
	v359 = v338 & int32(65535)
	goto L63
L74:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317+v233))))
	v330 = v314 + v329
	v331 = int32(1)
	v334 = v323 + v331
	if v334 != v246 {
		v314 = v330
		v317 = v317 + v331
		v323 = v334
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v338 = v330
	goto L73
L76:
	;
	goto L75
L77:
	;
	F_px_debug(m, int32(449820), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	v396 = int32(-113)
	goto L59
L79:
	;
	v405 = v237 - int32(3)
	if base.Ui32(v405) < base.Ui32(int32(33)) {
		goto L41
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v405
	F_px_debug(m, int32(38436), v19)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	v415 = int32(-111)
	goto L42
L82:
	;
	v448 = v415
	goto L1
L83:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v442 = F_pgp_mpi_free(m, v441)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L5
	} else {
		goto L87
	}
L84:
	;
	v439 = F__emscripten_memcpy_bulkmem(m, l0+int32(132), v192+int32(3), v405)
	mBase = m.M
	goto L86
L85:
	;
	goto L86
L86:
	;
	goto L83
L87:
	;
	v444 = F_pgp_expect_packet_end(m, l1)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	v448 = v444
	goto L1
}
func F_pgp_set_s2k_mode(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	v4 = int32(-13)
	if base.Ui32(int32(3)) < base.Ui32(l1) {
		v11 = v4
	} else {
		if l1 == int32(2) {
			v11 = v4
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l1
			v11 = int32(0)
		}
	}
	return v11
}
func F_pgp_sym_decrypt_bytea(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_pg_detoast_datum_packed(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v14 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v18 = F_pg_detoast_datum_packed(m, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = v18
					v21 = int32(0)
					v24 = F_decrypt_internal(m, v21, v21, v7, v12, v21, v20)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v26 != v7 {
							F_pfree(m, v7)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v30 != v12 {
									F_pfree(m, v12)
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return int32(0)
									} else {
										v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v34 < int32(3) {
											return v24
										} else {
											v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v20 == v37 {
												return v24
											} else {
												F_pfree(m, v20)
												mBase = m.M
												v40 = m.ExcPending
												if v40 != 0 {
													return int32(0)
												} else {
													return v24
												}
											}
										}
									}
								} else {
									v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v34 < int32(3) {
										return v24
									} else {
										v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v20 == v37 {
											return v24
										} else {
											F_pfree(m, v20)
											mBase = m.M
											v40 = m.ExcPending
											if v40 != 0 {
												return int32(0)
											} else {
												return v24
											}
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v30 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int32(0)
								} else {
									v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v34 < int32(3) {
										return v24
									} else {
										v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v20 == v37 {
											return v24
										} else {
											F_pfree(m, v20)
											mBase = m.M
											v40 = m.ExcPending
											if v40 != 0 {
												return int32(0)
											} else {
												return v24
											}
										}
									}
								}
							} else {
								v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v34 < int32(3) {
									return v24
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v20 == v37 {
										return v24
									} else {
										F_pfree(m, v20)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return int32(0)
										} else {
											return v24
										}
									}
								}
							}
						}
					}
				}
			} else {
				v20 = int32(0)
				v21 = int32(0)
				v24 = F_decrypt_internal(m, v21, v21, v7, v12, v21, v20)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v26 != v7 {
						F_pfree(m, v7)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v30 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int32(0)
								} else {
									v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v34 < int32(3) {
										return v24
									} else {
										v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v20 == v37 {
											return v24
										} else {
											F_pfree(m, v20)
											mBase = m.M
											v40 = m.ExcPending
											if v40 != 0 {
												return int32(0)
											} else {
												return v24
											}
										}
									}
								}
							} else {
								v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v34 < int32(3) {
									return v24
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v20 == v37 {
										return v24
									} else {
										F_pfree(m, v20)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return int32(0)
										} else {
											return v24
										}
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v30 != v12 {
							F_pfree(m, v12)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v34 < int32(3) {
									return v24
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v20 == v37 {
										return v24
									} else {
										F_pfree(m, v20)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return int32(0)
										} else {
											return v24
										}
									}
								}
							}
						} else {
							v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
							if v34 < int32(3) {
								return v24
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v20 == v37 {
									return v24
								} else {
									F_pfree(m, v20)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										return v24
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pgp_sym_encrypt_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_pg_detoast_datum_packed(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v14 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v18 = F_pg_detoast_datum_packed(m, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = v18
					v23 = F_encrypt_internal(m, int32(0), int32(1), v7, v12, v20)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v25 != v7 {
							F_pfree(m, v7)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v29 != v12 {
									F_pfree(m, v12)
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return int32(0)
									} else {
										v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v33 < int32(3) {
											return v23
										} else {
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v20 == v36 {
												return v23
											} else {
												F_pfree(m, v20)
												mBase = m.M
												v39 = m.ExcPending
												if v39 != 0 {
													return int32(0)
												} else {
													return v23
												}
											}
										}
									}
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v20 == v36 {
											return v23
										} else {
											F_pfree(m, v20)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							}
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v29 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v20 == v36 {
											return v23
										} else {
											F_pfree(m, v20)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v20 == v36 {
										return v23
									} else {
										F_pfree(m, v20)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						}
					}
				}
			} else {
				v20 = int32(0)
				v23 = F_encrypt_internal(m, int32(0), int32(1), v7, v12, v20)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v25 != v7 {
						F_pfree(m, v7)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v29 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v20 == v36 {
											return v23
										} else {
											F_pfree(m, v20)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v20 == v36 {
										return v23
									} else {
										F_pfree(m, v20)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v29 != v12 {
							F_pfree(m, v12)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v20 == v36 {
										return v23
									} else {
										F_pfree(m, v20)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						} else {
							v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
							if v33 < int32(3) {
								return v23
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v20 == v36 {
									return v23
								} else {
									F_pfree(m, v20)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										return v23
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
