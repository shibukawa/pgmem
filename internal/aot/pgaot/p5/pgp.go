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
										F_px_debug(m, int32(_a_F__pgp_read_public_key_0), v9)
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
	var v13 int32
	_ = v13
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v7 != 0 {
		v8 = int32(_a_F_pgp_cfb_encrypt_0)
	} else {
		v8 = int32(_a_F_pgp_cfb_encrypt_1)
	}
	F_cfb_process(m, l0, l1, l2, l3, v8, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
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
	v5 = F_pushf_create(m, l0, int32(_a_F_pgp_compress_filter_0), l1, l2)
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
	var v36 int32
	_ = v36
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
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
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
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v280 int32
	_ = v280
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
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
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
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
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
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
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
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
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
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	if v515 != 0 {
		goto L150
	} else {
		goto L151
	}
L2:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v499 == int32(0) {
		v505 = v489
		v511 = v495
		goto L1
	} else {
		goto L148
	}
L3:
	;
	return int32(0)
L4:
	;
	if v22 < int32(0) {
		v489 = v22
		v495 = v4
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v31 = l0 + int32(136)
	v36 = v4
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
	v505 = v471
	v511 = v477
	goto L1
L8:
	;
	if v49 <= int32(0) {
		v489 = v49
		v495 = v39
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
	v63 = F_pullf_create(m, v14+int32(24), int32(_a_F_pgp_decrypt_0), v56, v53)
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
	v489 = v63
	v495 = v39
	goto L2
L16:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	F_pullf_free(m, v481)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L3
	} else {
		goto L146
	}
L17:
	;
	v467 = int32(1)
	v471 = int32(-100)
	v474 = v467
	v477 = v467
	goto L16
L18:
	;
	v465 = int32(1)
	v471 = v463
	v474 = v465
	v477 = v465
	goto L16
L19:
	;
	v471 = int32(-100)
	v474 = v461
	v477 = v39
	goto L16
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v69
	F_px_debug(m, int32(_a_F_pgp_decrypt_1), v14)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L3
	} else {
		goto L145
	}
L21:
	;
	if v36 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L22:
	;
	if v36 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L23:
	;
	if v36 != 0 {
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
	v87 = F_pullf_read(m, v72, int32(_a_F_pgp_decrypt_2), v14+int32(32))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L3
	} else {
		goto L28
	}
L27:
	;
	v471 = v87
	v474 = v36
	v477 = v39
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
	v471 = v93
	v474 = int32(1)
	v477 = v39
	goto L16
L31:
	;
	F_px_debug(m, int32(_a_F_pgp_decrypt_3), int32(0))
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
	v103 = v14 + int32(32)
	v104 = F_pullf_read_fixed(m, v100, int32(1), v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L36
	}
L34:
	;
	v461 = int32(1)
	goto L19
L35:
	;
	v471 = v286
	v474 = int32(1)
	v477 = v39
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
	v110 = F_pullf_read_fixed(m, v100, int32(1), v103)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	if v110 < int32(0) {
		v286 = v110
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+32)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v114
	if v108 != int32(4) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_px_debug(m, int32(_a_F_pgp_decrypt_4), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L3
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v124 = F_pgp_s2k_read(m, v100, l0)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L44
	}
L43:
	;
	v471 = int32(-100)
	v474 = int32(1)
	v477 = v39
	goto L16
L44:
	;
	if v124 < int32(0) {
		v286 = v124
		goto L35
	} else {
		goto L45
	}
L45:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v128
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v130
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = (v132&int32(15) | int32(16)) << (uint(int32(base.Ui32(v132)>>(uint(int32(4))%32))+int32(6)) % 32)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v146 = F_pgp_s2k_process(m, l0, v143, v144, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	if v146 < int32(0) {
		v286 = v146
		goto L35
	} else {
		goto L47
	}
L47:
	;
	v155 = F_pullf_read_max(m, v100, int32(34), v14+int32(76), v14+int32(32))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	if v155 < int32(0) {
		v286 = v155
		goto L35
	} else {
		goto L49
	}
L49:
	;
	if v155 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L78
L51:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)))
	if v161 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	if base.Ui32(v155-int32(34)) <= base.Ui32(int32(-18)) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	base.MemoryCopy(m, v31, l0+int32(11), v161)
	goto L56
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v161
	v164 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v166
	v280 = v164
	goto L50
L57:
	;
	F_px_debug(m, int32(_a_F_pgp_decrypt_5), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L3
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = int32(1)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	v182 = m.G0
	v184 = v182 - int32(32)
	m.G0 = v184
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)))
	v192 = int32(0)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v195 = F_pgp_cfb_create(m, v184+int32(24), v188, l0+int32(11), v191, v192, v192, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L3
	} else {
		goto L62
	}
L60:
	;
	v471 = int32(-100)
	v474 = int32(1)
	v477 = v39
	goto L16
L61:
	;
	m.G0 = v184 + int32(32)
	v280 = v271
	goto L50
L62:
	;
	if v195 < int32(0) {
		v271 = v195
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v184)+24))
	v203 = F_pgp_cfb_decrypt(m, v199, v181, int32(1), v184+int32(31))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L3
	} else {
		goto L64
	}
L64:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v184)+24))
	v206 = int32(1)
	v209 = v155 - v206
	v212 = F_pgp_cfb_decrypt(m, v205, v181+v206, v209, l0+int32(136))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v184)+24))
	F_pgp_cfb_free(m, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L3
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v209
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+31)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v218
	v220 = int32(0)
	v223 = v218 - int32(2)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v223))|base.B2i32(int32(base.Ui32(int32(487))>>(uint(v223)%32))&int32(1) == v220) != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v240 == v209 {
		v271 = v220
		goto L61
	} else {
		goto L71
	}
L68:
	;
	v240 = int32(0)
	goto L70
L69:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v223<<(uint(int32(2))%32))+uint32(_c_F_pgp_decrypt[0])))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v240 = v239
	goto L70
L70:
	;
	goto L67
L71:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+31)))
	v245 = v242 - int32(2)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v245))|base.B2i32(int32(base.Ui32(int32(487))>>(uint(v245)%32))&int32(1) == int32(0)) != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v184)+8)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v184)+4)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = v242
	F_px_debug(m, int32(_a_F_pgp_decrypt_6), v184)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L3
	} else {
		goto L76
	}
L73:
	;
	v262 = int32(0)
	goto L75
L74:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v245<<(uint(int32(2))%32))+uint32(_c_F_pgp_decrypt[0])))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
	v262 = v261
	goto L75
L75:
	;
	goto L72
L76:
	;
	v271 = int32(-100)
	goto L61
L77:
	;
	v286 = v280
	goto L35
L78:
	;
	base.MemoryFill(m, v14+int32(32), int32(0), int32(34))
	goto L80
L80:
	;
	goto L77
L81:
	;
	v293 = int32(0)
	F_px_debug(m, int32(_a_F_pgp_decrypt_7), v293)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L3
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	if v39 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v461 = v293
	goto L19
L85:
	;
	F_px_debug(m, int32(_a_F_pgp_decrypt_8), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L3
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v302 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v302
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v305 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v305
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v305
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v305
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v318 = F_pgp_cfb_create(m, v14+int32(32), v313, v31, v314, v302, v305, v317)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L3
	} else {
		goto L90
	}
L88:
	;
	goto L17
L89:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	if v352 != 0 {
		goto L101
	} else {
		goto L102
	}
L90:
	;
	if v318 < int32(0) {
		v350 = v318
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v326 = F_pullf_create(m, v14+int32(76), int32(_a_F_pgp_decrypt_9), v325, v304)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L3
	} else {
		goto L92
	}
L92:
	;
	if v326 < int32(0) {
		v350 = v326
		goto L89
	} else {
		goto L93
	}
L93:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	v334 = F_pullf_create(m, v14+int32(72), int32(_a_F_pgp_decrypt_10), l0, v333)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L3
	} else {
		goto L94
	}
L94:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	if int32(0) <= v334 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v341 = F_process_data_packets(m, l0, l2, v336, int32(1), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L3
	} else {
		goto L98
	}
L96:
	;
	v344 = v334
	v345 = v336
	goto L97
L97:
	;
	if v345 == int32(0) {
		v350 = v344
		goto L89
	} else {
		goto L99
	}
L98:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v344 = v341
	v345 = v343
	goto L97
L99:
	;
	F_pullf_free(m, v345)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L3
	} else {
		goto L100
	}
L100:
	;
	v350 = v344
	goto L89
L101:
	;
	F_pullf_free(m, v352)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L3
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	if v355 == int32(0) {
		v463 = v350
		goto L18
	} else {
		goto L105
	}
L104:
	;
	goto L103
L105:
	;
	F_pgp_cfb_free(m, v355)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L3
	} else {
		goto L106
	}
L106:
	;
	v463 = v350
	goto L18
L107:
	;
	v362 = int32(0)
	F_px_debug(m, int32(_a_F_pgp_decrypt_7), v362)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L3
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	if v39 != 0 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v461 = v362
	goto L19
L111:
	;
	F_px_debug(m, int32(_a_F_pgp_decrypt_11), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L3
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v371 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v371
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v371
	v385 = F_pullf_read_fixed(m, v373, int32(1), v14+int32(67))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L3
	} else {
		goto L115
	}
L114:
	;
	goto L17
L115:
	;
	if v385 < int32(0) {
		v463 = v385
		goto L18
	} else {
		goto L116
	}
L116:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+67)))
	if v389 != int32(1) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	F_px_debug(m, int32(_a_F_pgp_decrypt_12), int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L3
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v401 = int32(0)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v404 = F_pgp_cfb_create(m, v14+int32(32), v399, v31, v400, v401, v401, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L3
	} else {
		goto L122
	}
L120:
	;
	v463 = int32(-100)
	goto L18
L121:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	if v446 != 0 {
		goto L135
	} else {
		goto L136
	}
L122:
	;
	if v404 < int32(0) {
		v444 = v404
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v412 = F_pullf_create(m, v14+int32(76), int32(_a_F_pgp_decrypt_9), v411, v373)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L3
	} else {
		goto L124
	}
L124:
	;
	if v412 < int32(0) {
		v444 = v412
		goto L121
	} else {
		goto L125
	}
L125:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	v420 = F_pullf_create(m, v14+int32(68), int32(_a_F_pgp_decrypt_13), l0, v419)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L3
	} else {
		goto L126
	}
L126:
	;
	if v420 < int32(0) {
		v444 = v420
		goto L121
	} else {
		goto L127
	}
L127:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	v428 = F_pullf_create(m, v14+int32(72), int32(_a_F_pgp_decrypt_10), l0, v427)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L3
	} else {
		goto L128
	}
L128:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	if int32(0) <= v428 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v433 = int32(1)
	v435 = F_process_data_packets(m, l0, l2, v430, v433, v433)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L3
	} else {
		goto L132
	}
L130:
	;
	v438 = v428
	v439 = v430
	goto L131
L131:
	;
	if v439 == int32(0) {
		v444 = v438
		goto L121
	} else {
		goto L133
	}
L132:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v438 = v435
	v439 = v437
	goto L131
L133:
	;
	F_pullf_free(m, v439)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L3
	} else {
		goto L134
	}
L134:
	;
	v444 = v438
	goto L121
L135:
	;
	F_pullf_free(m, v446)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L3
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	if v449 != 0 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	goto L137
L139:
	;
	F_pullf_free(m, v449)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L3
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	if v452 == int32(0) {
		v463 = v444
		goto L18
	} else {
		goto L143
	}
L142:
	;
	goto L141
L143:
	;
	F_pgp_cfb_free(m, v452)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L3
	} else {
		goto L144
	}
L144:
	;
	v463 = v444
	goto L18
L145:
	;
	v461 = v36
	goto L19
L146:
	;
	v484 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v484
	if v484 <= v471 {
		v36 = v474
		v39 = v477
		goto L6
	} else {
		goto L147
	}
L147:
	;
	goto L7
L148:
	;
	F_pullf_free(m, v499)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L3
	} else {
		goto L149
	}
L149:
	;
	v505 = v489
	v511 = v495
	goto L1
L150:
	;
	F_pullf_free(m, v515)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L3
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	if v505 < int32(0) {
		v530 = v505
		goto L154
	} else {
		goto L155
	}
L153:
	;
	goto L152
L154:
	;
	m.G0 = v14 + int32(80)
	return v530
L155:
	;
	v520 = int32(-100)
	if v511 == int32(0) {
		v530 = v520
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v523 != 0 {
		v530 = v520
		goto L154
	} else {
		goto L157
	}
L157:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v525 != 0 {
		v530 = int32(-102)
		goto L154
	} else {
		goto L158
	}
L158:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v528 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v529 = int32(-106)
	goto L161
L160:
	;
	v529 = int32(0)
	goto L161
L161:
	;
	v530 = v529
	goto L154
}
func F_pgp_key_id_w(m *base.Module, l0 int32) int64 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
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
	var v72 int32
	_ = v72
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		if v11 == int32(1) {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
			if v17 == int32(18) {
				v20 = int32(16)
			} else {
				v20 = int32(0)
			}
			if base.Ui32((v17-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v27 = int32(4)
			} else {
				v27 = v20
			}
			v40 = v27
		} else {
			v28 = int32(1)
			if v11&v28 != 0 {
				v40 = int32(base.Ui32(v11)>>(uint(v28)%32)) - v28
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v40 = int32(base.Ui32(v34)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v41 = int32(1)
		if v11&v41 != 0 {
			v45 = v41
		} else {
			v45 = int32(4)
		}
		v47 = F_mbuf_create_from_data(m, v7+v45, v40)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int64(0)
		} else {
			v50 = F_palloc(m, int32(21))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int64(0)
			} else {
				v54 = F_pgp_get_keyid(m, v47, v50+int32(4))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int64(0)
				} else {
					v56 = F_mbuf_free(m, v47)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int64(0)
					} else {
						if int32(0) <= v54 {
							*(*int32)(unsafe.Add(mBase, uint32(v50))) = v54<<(uint(int32(2))%32) + int32(16)
							v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v65 != v7 {
								F_pfree(m, v7)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int64(0)
								} else {
									return base.I64_extend_i32_u(v50)
								}
							} else {
								return base.I64_extend_i32_u(v50)
							}
						} else {
							F_px_THROW_ERROR(m, v54)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int64(0)
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
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	if l0 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v5 = v3 + int32(12)
		if v5 != 0 {
			base.MemoryFill(m, l0, int32(0), v5)
		} else {
		}
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
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
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
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
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v312 int32
	_ = v312
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	v3 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v19 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v17 + int32(48)
	return v405
L2:
	;
	F_px_debug(m, int32(_a_F_pgp_parse_pubenc_sesskey_0), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v32 = F_pullf_read_fixed(m, l1, int32(1), v17+int32(36))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v405 = int32(-12)
	goto L1
L7:
	;
	if v32 < int32(0) {
		v405 = v32
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+36)))
	if v36 != int32(3) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v36
	F_px_debug(m, int32(_a_F_pgp_parse_pubenc_sesskey_1), v17+int32(16))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v49 = F_pullf_read_fixed(m, l1, int32(8), v17+int32(36))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v405 = int32(-100)
	goto L1
L13:
	;
	if v49 < int32(0) {
		v405 = v49
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v17)+36))
	if v53 == int64(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v67 = F_pullf_read_fixed(m, l1, int32(1), v17+int32(44))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L5
	} else {
		goto L19
	}
L16:
	;
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v17)+36))
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v19)+40))
	if v56 == v57 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_px_debug(m, int32(_a_F_pgp_parse_pubenc_sesskey_2), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v405 = int32(-113)
	goto L1
L19:
	;
	if v67 < int32(0) {
		v405 = v67
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+44)))
	switch v72 - int32(1) {
	case 0, 1:
		goto L22
	default:
		v405 = int32(-112)
		goto L1
	case 15:
		goto L23
	}
L21:
	;
	if v139 < int32(0) {
		v405 = v139
		goto L1
	} else {
		goto L40
	}
L22:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+5)))
	v119 = int32(1)
	if base.Ui32(v119) < base.Ui32((v118-v119)&int32(255)) {
		v139 = int32(-113)
		goto L21
	} else {
		goto L35
	}
L23:
	;
	v75 = m.G0
	v76 = int32(16)
	v77 = v75 - v76
	m.G0 = v77
	v79 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = v79
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+5)))
	if v84 == v76 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v91 = F_pgp_mpi_read(m, l1, v77+int32(12))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L28
	}
L25:
	;
	v112 = int32(-113)
	goto L26
L26:
	;
	m.G0 = v77 + int32(16)
	v139 = v112
	goto L21
L27:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v107 = F_pgp_mpi_free(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L33
	}
L28:
	;
	if v91 < int32(0) {
		v105 = v91
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v97 = F_pgp_mpi_read(m, l1, v77+int32(8))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	if v97 < int32(0) {
		v105 = v97
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v103 = F_pgp_elgamal_decrypt(m, v19, v101, v102, v17+int32(32))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v105 = v103
	goto L27
L33:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v110 = F_pgp_mpi_free(m, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v112 = v105
	goto L26
L35:
	;
	v127 = F_pgp_mpi_read(m, l1, v17+int32(44))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	if v127 < int32(0) {
		v139 = v127
		goto L21
	} else {
		goto L37
	}
L37:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	v134 = F_pgp_rsa_decrypt(m, v19, v131, v17+int32(32))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	v137 = F_pgp_mpi_free(m, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	v139 = v134
	goto L21
L40:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	if v145 < int32(10) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v365
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v390
	if v365 != 0 {
		goto L81
	} else {
		goto L82
	}
L42:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v388 = F_pgp_mpi_free(m, v387)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L5
	} else {
		goto L80
	}
L43:
	;
	v216 = v184 + int32(2)
	v219 = v148 - v216 + v145
	if v219 < int32(3) {
		v350 = int32(-113)
		goto L59
	} else {
		goto L60
	}
L44:
	;
	F_px_debug(m, int32(_a_F_pgp_parse_pubenc_sesskey_3), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L5
	} else {
		goto L58
	}
L45:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v149 != int32(2) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v152 = v148 + v145
	v154 = v148 + int32(1)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
	if v155 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v180 == v152 {
		goto L44
	} else {
		goto L55
	}
L48:
	;
	v180 = v154
	v184 = v148
	v186 = v3
	goto L47
L49:
	;
	goto L50
L50:
	;
	v160 = v154
	v166 = v3
	goto L51
L51:
	;
	v172 = int32(1)
	v173 = v166 + v172
	v175 = v160 + v172
	if base.Ui32(v152) <= base.Ui32(v175) {
		v180 = v175
		v184 = v160
		v186 = v173
		goto L47
	} else {
		goto L53
	}
L52:
	;
	v180 = v175
	v184 = v160
	v186 = v173
	goto L47
L53:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v177 != 0 {
		v160 = v175
		v166 = v173
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v193 != 0 {
		goto L44
	} else {
		goto L56
	}
L56:
	;
	if base.Ui32(int32(7)) < base.Ui32(v186) {
		goto L43
	} else {
		goto L57
	}
L57:
	;
	goto L44
L58:
	;
	v375 = int32(-113)
	goto L42
L59:
	;
	if v350 < int32(0) {
		v375 = v350
		goto L42
	} else {
		goto L77
	}
L60:
	;
	if v219 != int32(3) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v225 = int32(3)
	v226 = v219 - v225
	v228 = v226 & v225
	v229 = int32(1)
	if base.Ui32(v225) <= base.Ui32(v219-int32(4)) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v325 = v3
	goto L63
L63:
	;
	v332 = v219 + v216
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332-int32(2)))))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332-int32(1)))))
	if v325 == v335<<(uint(int32(8))%32)|v340 {
		v350 = int32(0)
		goto L59
	} else {
		goto L75
	}
L64:
	;
	v325 = v312 & int32(_a_F_pgp_parse_pubenc_sesskey_4)
	goto L63
L65:
	;
	v243 = v229
	v246 = v3
	v248 = v3
	goto L68
L66:
	;
	v273 = v229
	v276 = v3
	goto L67
L67:
	;
	v287 = v273
	v290 = v276
	v293 = v3
	goto L72
L68:
	;
	v250 = v243 + v216
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+1)))
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+2)))
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+3)))
	v258 = v246 + v251 + v253 + v255 + v257
	v259 = int32(4)
	v260 = v243 + v259
	v262 = v248 + v259
	if v262 != v226&int32(-4) {
		v243 = v260
		v246 = v258
		v248 = v262
		goto L68
	} else {
		goto L70
	}
L69:
	;
	if v228 == int32(0) {
		v312 = v258
		goto L64
	} else {
		goto L71
	}
L70:
	;
	goto L69
L71:
	;
	v273 = v260
	v276 = v258
	goto L67
L72:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287+v216))))
	v296 = v290 + v295
	v297 = int32(1)
	v300 = v293 + v297
	if v300 != v228 {
		v287 = v287 + v297
		v290 = v296
		v293 = v300
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v312 = v296
	goto L64
L74:
	;
	goto L73
L75:
	;
	F_px_debug(m, int32(_a_F_pgp_parse_pubenc_sesskey_5), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	v350 = int32(-113)
	goto L59
L77:
	;
	v365 = v219 - int32(3)
	if base.Ui32(v365) < base.Ui32(int32(33)) {
		goto L41
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v365
	F_px_debug(m, int32(_a_F_pgp_parse_pubenc_sesskey_6), v17)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	v375 = int32(-111)
	goto L42
L80:
	;
	v405 = v375
	goto L1
L81:
	;
	base.MemoryCopy(m, l0+int32(136), v184+int32(3), v365)
	goto L83
L82:
	;
	goto L83
L83:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v399 = F_pgp_mpi_free(m, v398)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	v401 = F_pgp_expect_packet_end(m, l1)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	v405 = v401
	goto L1
}
func F_pgp_set_s2k_mode(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	if base.B2i32(l1 == int32(2))|base.B2i32(base.Ui32(int32(3)) < base.Ui32(l1)) != 0 {
		v13 = int32(-13)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l1
		v13 = int32(0)
	}
	return v13
}
func F_pgp_sym_decrypt_bytea(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v17 int32
	_ = v17
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
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v17 {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v21 = F_pg_detoast_datum_packed(m, v20)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					v23 = v21
					v24 = F_decrypt_internal(m, v2, v2, v9, v14, int32(0), v23)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int64(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v26 != v9 {
							F_pfree(m, v9)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int64(0)
							} else {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v30 != v14 {
									F_pfree(m, v14)
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return int64(0)
									} else {
										v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v34 < int32(3) {
											return base.I64_extend_i32_u(v24)
										} else {
											v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
											if v23 == v37 {
												return base.I64_extend_i32_u(v24)
											} else {
												F_pfree(m, v23)
												mBase = m.M
												v40 = m.ExcPending
												if v40 != 0 {
													return int64(0)
												} else {
													return base.I64_extend_i32_u(v24)
												}
											}
										}
									}
								} else {
									v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v34 < int32(3) {
										return base.I64_extend_i32_u(v24)
									} else {
										v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if v23 == v37 {
											return base.I64_extend_i32_u(v24)
										} else {
											F_pfree(m, v23)
											mBase = m.M
											v40 = m.ExcPending
											if v40 != 0 {
												return int64(0)
											} else {
												return base.I64_extend_i32_u(v24)
											}
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v30 != v14 {
								F_pfree(m, v14)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int64(0)
								} else {
									v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v34 < int32(3) {
										return base.I64_extend_i32_u(v24)
									} else {
										v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if v23 == v37 {
											return base.I64_extend_i32_u(v24)
										} else {
											F_pfree(m, v23)
											mBase = m.M
											v40 = m.ExcPending
											if v40 != 0 {
												return int64(0)
											} else {
												return base.I64_extend_i32_u(v24)
											}
										}
									}
								}
							} else {
								v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v34 < int32(3) {
									return base.I64_extend_i32_u(v24)
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									if v23 == v37 {
										return base.I64_extend_i32_u(v24)
									} else {
										F_pfree(m, v23)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return int64(0)
										} else {
											return base.I64_extend_i32_u(v24)
										}
									}
								}
							}
						}
					}
				}
			} else {
				v23 = v2
				v24 = F_decrypt_internal(m, v2, v2, v9, v14, int32(0), v23)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v26 != v9 {
						F_pfree(m, v9)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int64(0)
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v30 != v14 {
								F_pfree(m, v14)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int64(0)
								} else {
									v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v34 < int32(3) {
										return base.I64_extend_i32_u(v24)
									} else {
										v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if v23 == v37 {
											return base.I64_extend_i32_u(v24)
										} else {
											F_pfree(m, v23)
											mBase = m.M
											v40 = m.ExcPending
											if v40 != 0 {
												return int64(0)
											} else {
												return base.I64_extend_i32_u(v24)
											}
										}
									}
								}
							} else {
								v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v34 < int32(3) {
									return base.I64_extend_i32_u(v24)
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									if v23 == v37 {
										return base.I64_extend_i32_u(v24)
									} else {
										F_pfree(m, v23)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return int64(0)
										} else {
											return base.I64_extend_i32_u(v24)
										}
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v30 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int64(0)
							} else {
								v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v34 < int32(3) {
									return base.I64_extend_i32_u(v24)
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									if v23 == v37 {
										return base.I64_extend_i32_u(v24)
									} else {
										F_pfree(m, v23)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return int64(0)
										} else {
											return base.I64_extend_i32_u(v24)
										}
									}
								}
							}
						} else {
							v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
							if v34 < int32(3) {
								return base.I64_extend_i32_u(v24)
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
								if v23 == v37 {
									return base.I64_extend_i32_u(v24)
								} else {
									F_pfree(m, v23)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int64(0)
									} else {
										return base.I64_extend_i32_u(v24)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pgp_sym_encrypt_text(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14352(m, l0, int32(1), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
